package poller

import (
	"crypto/sha256"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Tyler-Arciniaga/SWESniper/internal/models"
	"github.com/Tyler-Arciniaga/SWESniper/internal/notifier"
	"github.com/Tyler-Arciniaga/SWESniper/internal/services"
)

type Poller struct {
	UrlService       services.URLService
	ChangeLogService services.ChangeLogService
	ScraperService   services.ScraperService
	DiffCheckService services.DiffService
	Notifier         notifier.Notifier
}

// current plan: think it's probably best to run poller every minute since that is the lowest
// check interval possible (production), in dev i'll probably test with 15 seconds for QOL

/* Poller Start Flow
Start()
├── Every 10s:
│   ├── Get all URLs from store
│   ├── For each URL:
│   │   ├── If time.Since(LastCheckedAt) > CheckInterval:
│   │   │   ├── Fetch URL
│   │   │   ├── Compare to previous
│   │   │   ├── Log if changed
│   │   │   ├── Update LastCheckedAt
*/

func (p *Poller) StartPoller() {
	fmt.Print("Starting poller...\n")
	//DEV: change back to 10 * time.Second after testing non poller funcitonality
	ticker := time.NewTicker(10 * time.Second)

	for range ticker.C {

		var wg sync.WaitGroup

		url_list, _ := p.UrlService.GetAllURLsGlobally()

		for _, value := range url_list {
			wg.Add(1)

			go func(v models.URLRecord) {
				defer wg.Done()
				p.CheckURL(&v)
			}(value)
		}

		wg.Wait()
	}
}

/*
	High level workflow of checkURL():

1. Make an HTTP GET request to the URL
2. Read the response body
3. Hash the body (e.g., SHA256)
4. Compare the hash to the previously stored hash
5. If different:
  - Mark as changed (log, notify, update)

6. Always:
  - Update LastCheckedAt
  - Store the new hash
*/
func (p *Poller) CheckURL(r *models.URLRecord) {
	if time.Since(r.LastCheckedAt) >= time.Duration(r.CheckInterval)*time.Second {
		scrappedContent_RawString, scrappedContent_Formatted, e := p.ScraperService.ExtractURLContent(r.URL)

		if scrappedContent_RawString == "" && scrappedContent_Formatted == nil {
			log.Printf("No content was extracted for %q", r.URL)
		}

		if e != nil {
			r.LastKnownHash = "INVALID_COULD_NOT_EXTRACT_WEB_CONTENT"
			e = p.UrlService.UpdateURL(r)
			if e != nil {
				log.Fatalf("Error invalidating url: %v", e)
			}
			return
		}

		newHash, e := p.FetchHash(string(scrappedContent_RawString))

		if e != nil {
			log.Printf("Error creating hash value for new extracted url content: %v\n", e)
			return
		}

		if newHash != r.LastKnownHash {
			log.Printf("Detected change in URL %q\n", r.URL)
			/*
				diffRes := p.DiffCheckService.DiffCheckContentsFormatted(r.LastKnownContent, scrappedContent_Formatted)
				newChangeLog := models.ChangeRecord{URL: r.URL, Timestamp: time.Now(), Added: diffRes.Added, DiffSummary: diffRes.Summary}

				p.ChangeLogService.PersistChangeRecord(&newChangeLog)

				tmp := r.Description
				e = p.Notifier.SendNotification(newChangeLog, tmp)
				if e != nil {
					log.Printf("error sending notification for: %q\n", r.URL)
				}
			*/

			if scrappedContent_Formatted != nil {
				r.LastKnownContent = scrappedContent_Formatted
			} else {
				r.LastKnownContent = nil
			}

		}
		r.LastCheckedAt = time.Now()
		r.LastKnownHash = newHash

		e = p.UrlService.UpdateURL(r)
		if e != nil {
			log.Fatalf("Error: %v", e)
		}
	}
}

func (p *Poller) FetchHash(body string) (string, error) {
	hasher := sha256.New()

	_, e := hasher.Write([]byte(body))

	if e != nil {
		return "", e
	}

	hashed_body := hasher.Sum(nil)

	return fmt.Sprintf("%x", hashed_body), nil
}
