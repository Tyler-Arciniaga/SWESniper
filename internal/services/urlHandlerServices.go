package services

import "github.com/Tyler-Arciniaga/SWESniper/internal/models"

// validates add URL POST request
func ValidateURLPost(r *models.AddURLRequest) bool {
	if r.CheckInterval < 60 || r.CheckInterval > 86400 {
		return false
	} //reject check intervals that are greater than a day or less than 1 minute

	return true
}
