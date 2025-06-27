package services

import (
	"os"
	"testing"
)

func TestDiffService(t *testing.T) {
	diffService := DiffService{}
	t.Run("inspect return of diff check with multiple differences", func(t *testing.T) {
		repo1_byte, _ := os.ReadFile("../../testdata/repo1_short.txt")
		repo2_byte, _ := os.ReadFile("../../testdata/repo1_changed.txt")

		diffService.DiffCheckContents(string(repo1_byte), string(repo2_byte))

		//fmt.Println(r.Added)
		//fmt.Println(r.Summary)
	})
}
