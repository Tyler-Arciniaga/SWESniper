package models

type AddURLRequest struct {
	URL           string `json:"url"`
	Description   string `json:"description"`
	CheckInterval int    `json:"checkInterval"`
}
