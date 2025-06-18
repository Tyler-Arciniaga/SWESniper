package services

import (
	"fmt"

	"github.com/Tyler-Arciniaga/SWESniper/internal/models"
)

type URLService struct {
	Store URLStore
}

type URLStore interface {
	SaveURL(r models.URLRecord) error
}

// validates add URL POST request
func (s *URLService) ValidateURLPost(r *models.AddURLRequest) error {
	if r.CheckInterval < 60 || r.CheckInterval > 86400 {
		return fmt.Errorf("check interval must be between 1 minute and 1 day")
	} //reject check intervals that are greater than a day or less than 1 minute

	return nil
}
