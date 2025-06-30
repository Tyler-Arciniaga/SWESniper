package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type ScraperService struct{}

// Todo: make this a smart scraper to have different parsing logic depending on domain of url (i.e. github vs linkedin vs glassdoor, etc)
func (s *ScraperService) ExtractURLContent(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", fmt.Errorf("get request to %q failed", url)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("did not recieve 200 status code from GET request, instead: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to parse url content")
	}

	//NOTE: as of now for the MVP the scraper only works with Github public repositories.
	github_readme := doc.Find("article.markdown-body")

	if github_readme.Length() < 1 {
		return "", fmt.Errorf("no github readme file detected")
	}

	accessiblity_table := github_readme.Find("markdown-accessiblity-table")

	if accessiblity_table.Length() == 0 {
		log.Print("checking for a-table under strong tag...")
		accessiblity_table = github_readme.Find("strong markdown-accessiblity-table")
	}

	table := accessiblity_table.Find("table")

	if table.Length() == 0 {
		log.Printf("no table detected under: %v", accessiblity_table)
	}

	if table.Length() >= 1 {
		return table.Text(), nil
	}

	return github_readme.Text(), nil
}
