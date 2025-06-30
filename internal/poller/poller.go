package poller

import (
	"crypto/sha256"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Tyler-Arciniaga/SWESniper/internal/models"
	"github.com/Tyler-Arciniaga/SWESniper/internal/services"
)

type Poller struct {
	UrlService       services.URLService
	ChangeLogService services.ChangeLogService
	ScraperService   services.ScraperService
	DiffCheckService services.DiffService
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
	ticker := time.NewTicker(10 * time.Second)

	for range ticker.C {

		var wg sync.WaitGroup

		url_list, _ := p.UrlService.GetAllURLs()

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
	//DONT FORGET TO CHANGE TO time.duration(r.CheckInterval).seconds()
	if time.Since(r.LastCheckedAt) >= time.Duration(r.CheckInterval) {
		scrappedContent_RawString, scrappedContent_Formatted, e := p.ScraperService.ExtractURLContent(r.URL)

		if scrappedContent_RawString == "" && scrappedContent_Formatted == nil {
			log.Printf("No content was extracted for %q", r.URL)
		}

		if e != nil {
			log.Printf("Failed to extract main content from url %q, recieved err %q\n", r.URL, e)
			return
		}

		//DELETE ME! (used to generate testData from real URLs)
		//os.WriteFile("../testdata/repo2.txt", []byte(scrappedContent), 0644)

		newHash, e := p.FetchHash(string(scrappedContent_RawString))

		if e != nil {
			log.Printf("Error creating hash value for new extracted url content: %v\n", e)
			return
		}

		if newHash != r.LastKnownHash {
			log.Printf("Detected change in URL %q\n", r.URL)
			//diffRes := p.DiffCheckService.DiffCheckContents(strings.Join(r.LastKnownContent, ""), scrappedContent_RawString)
			//newChangeLog := models.ChangeRecord{URL: r.URL, Timestamp: time.Now(), DiffSummary: diffRes.Summary}
			//p.ChangeLogService.PersistChangeRecord(&newChangeLog)
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
