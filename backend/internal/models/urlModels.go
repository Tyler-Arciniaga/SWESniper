package models

import "time"

type AddURLRequest struct {
	URL           string `json:"url"`
	Description   string `json:"description"`
	CheckInterval int    `json:"checkInterval"`
}

type URLRecord struct {
	ID               int          `json:"id"`
	User_id          string       `json:"user_id"` //currently representing as string, consider changing to using go UUID package
	URL              string       `json:"url"`
	Description      string       `json:"description"`
	CheckInterval    int          `json:"checkInterval"`
	LastCheckedAt    time.Time    `json:"lastCheckAt"`
	LastKnownHash    string       `json:"lastKnownHash"`
	LastKnownContent []JobListing `json:"lastKnownContent"`
	Created_at       time.Time    `json:"created_at"`
}
