package services

import (
	"fmt"

	"github.com/Tyler-Arciniaga/SWESniper/internal/models"
)

type DiffCheckResult struct {
	Added   []models.JobListing
	Summary string
}

type DiffService struct{}

func (s *DiffService) DiffCheckContentsFormatted(old, new []models.JobListing) DiffCheckResult {
	/*
		dmp := diffmatchpatch.DiffMatchPatch{}
		diffs := dmp.DiffMain(old, new, false)
		diffs = dmp.DiffCleanupSemantic(diffs)
	*/
	var added []models.JobListing

	oldMap := make(map[string]models.JobListing)
	for _, listing := range old {
		oldMap[listing.Hash()] = listing
	}

	for _, listing := range new {
		if _, exists := oldMap[listing.Hash()]; !exists {
			added = append(added, listing)
		}
	}

	summary := s.GenerateDiffSummary(added)
	r := DiffCheckResult{Added: added, Summary: summary}

	return r
}

// Note: right now generate diff summary will almost always report one detected change
// (since most times the insert is done in a bulk of text), thus need to look into better parsing logic
func (s *DiffService) GenerateDiffSummary(added []models.JobListing) string {
	return fmt.Sprintf("Number of detected changes: %d", len(added))
}

//FUTURE: add parsing logic to make viewing changes in text content more human readable
