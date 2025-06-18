package poller

import (
	"fmt"
	"sync"
	"time"

	"github.com/Tyler-Arciniaga/SWESniper/internal/models"
	"github.com/Tyler-Arciniaga/SWESniper/internal/services"
)

type Poller struct {
	Store services.URLStore //interface
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

		url_list, _ := p.Store.GetAll()

		for _, value := range url_list {
			wg.Add(1)

			go func(v models.URLRecord) {
				defer wg.Done()
				p.CheckURL(v)
			}(value)
		}

		wg.Wait()
	}
}

func (p *Poller) CheckURL(r models.URLRecord) {
	fmt.Println(r)
}
