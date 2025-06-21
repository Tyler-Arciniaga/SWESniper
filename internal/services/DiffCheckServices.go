package services

import (
	"fmt"

	"github.com/sergi/go-diff/diffmatchpatch"
)

type DiffCheckResult struct {
	Added   []string
	Summary string
}

type DiffService struct{}

func (s *DiffService) DiffCheckContents(old, new string) DiffCheckResult {
	dmp := diffmatchpatch.DiffMatchPatch{}

	diffs := dmp.DiffMain(old, new, false)

	diffs = dmp.DiffCleanupSemantic(diffs)

	var added []string

	for _, change := range diffs {
		if change.Type.String() == "Insert" {
			added = append(added, change.Text)
		}
	}

	summary := s.GenerateDiffSummary(added)
	r := DiffCheckResult{Added: added, Summary: summary}

	return r
}

// Note: right now generate diff summary will almost always report one detected change
// (since most times the insert is done in a bulk of text), thus need to look into better parsing logic
func (s *DiffService) GenerateDiffSummary(added []string) string {
	return fmt.Sprintf("Number of detected changes to URL text content: %d", len(added))
}

//FUTURE: add parsing logic to make viewing changes in text content more human readable
