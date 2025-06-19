package poller

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/Tyler-Arciniaga/SWESniper/internal/models"
	"github.com/Tyler-Arciniaga/SWESniper/internal/services"
)

type Poller struct {
	UrlService       services.URLService
	ChangeLogService services.ChangeLogService
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
	if time.Since(r.LastCheckedAt) >= time.Duration(r.CheckInterval) {
		resp, err := http.Get(r.URL)

		if err != nil {
			fmt.Printf("Get request to %q failed", r.URL)
			return
		}

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("Failed to read body of Get request response")
		}

		newHashedBody, e := p.FetchHash(string(body))

		if e != nil {
			fmt.Print(e)
		}

		if newHashedBody != r.LastKnownHash {
			log.Printf("Detected change in URL %q", r.URL)
			newChangeLog := models.ChangeRecord{URL: r.URL, Timestamp: time.Now(), DiffSummary: "changed"}
			p.ChangeLogService.PersistChangeRecord(&newChangeLog)
		}

		r.LastCheckedAt = time.Now()
		r.LastKnownHash = newHashedBody

		_ = p.UrlService.UpdateURL(r)
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
