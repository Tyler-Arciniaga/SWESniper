package services

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestLinkedinParser(t *testing.T) {
	scraperService := ScraperService{}

	t.Run("test linkedin website parsing", func(t *testing.T) {
		url := "https://www.linkedin.com/jobs/search-results/?currentJobId=4132249048&eBP=NON_CHARGEABLE_CHANNEL&f_TPR=r86400&keywords=software%20engineer%20intern&origin=JOB_COLLECTION_PAGE_SEARCH_BUTTON&refId=Klag2C0ELs7HKd%2FBIu3weg%3D%3D&trackingId=nlLG1ibJF4dW25o2g5I%2FGg%3D%3D"
		_, e := scraperService.ExtractURLContent(url)

		resp, _ := http.Get(url)

		doc, _ := goquery.NewDocumentFromReader(resp.Body)

		fmt.Print(doc.Text())

		if e != nil {
			t.Error(e.Error())
		}
	})
}
