package models

import "time"

type AddURLRequest struct {
	URL           string `json:"url"`
	Description   string `json:"description"`
	CheckInterval int    `json:"checkInterval"`
}

type URLRecord struct {
	ID               int          `json:"id"`
	URL              string       `json:"url"`
	Description      string       `json:"description"`
	CheckInterval    int          `json:"checkInterval"`
	LastCheckedAt    time.Time    `json:"lastCheckAt"`
	LastKnownHash    string       `json:"lastKnownHash"`
	LastKnownContent []JobListing `json:"lastKnownContent"`
	Created_at       time.Time    `json:"created_at"`
}
