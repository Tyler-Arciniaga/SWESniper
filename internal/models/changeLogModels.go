package models

import "time"

// DiffSummary will carry more advanced information later, for now just "changed"
type ChangeRecord struct {
	URL         string    `json:"url"`
	Timestamp   time.Time `json:"timestamp"`
	DiffSummary string    `json:"diffsummary"`
}
