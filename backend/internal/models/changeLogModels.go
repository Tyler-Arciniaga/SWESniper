package models

import "time"

type ChangeRecord struct {
	ID          int          `json:"id"`
	URL_id      int          `json:"url_id"`
	URL         string       `json:"url"`
	Timestamp   time.Time    `json:"timestamp"`
	Added       []JobListing `json:"added"`
	DiffSummary string       `json:"diffsummary"`
}
