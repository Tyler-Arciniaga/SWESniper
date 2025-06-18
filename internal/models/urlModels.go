package models

type AddURLRequest struct {
	URL           string `json:"url"`
	Description   string `json:"description"`
	CheckInterval int    `json:"checkInterval"`
}

type URLRecord struct {
	URL           string `json:"url"`
	Description   string `json:"description"`
	CheckInterval int    `json:"checkInterval"`
	Created_at    int64  `json:"created_at"`
}
