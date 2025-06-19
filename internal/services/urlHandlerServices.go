package services

import (
	"fmt"
	"time"

	"github.com/Tyler-Arciniaga/SWESniper/internal/models"
)

type URLService struct {
	URLStore URLStore
}

type URLStore interface {
	SaveURL(r models.URLRecord) error
	UpdateURLInfo(r models.URLRecord) error
	URL_GetAll() ([]models.URLRecord, error)
}

// validates add URL POST request
func (s *URLService) ValidateURLPost(r *models.AddURLRequest) error {
	if r.CheckInterval < 60 || r.CheckInterval > 86400 {
		return fmt.Errorf("check interval must be between 1 minute and 1 day")
	} //reject check intervals that are greater than a day or less than 1 minute

	return nil
}

func (s *URLService) StoreURL(r *models.AddURLRequest) error {
	c := time.Now()
	urlRecord := models.URLRecord{
		URL:           r.URL,
		Description:   r.Description,
		CheckInterval: r.CheckInterval,
		LastCheckedAt: c,
		LastKnownHash: "",
		Created_at:    c,
	}

	e := s.URLStore.SaveURL(urlRecord)
	if e != nil {
		return e
	}
	return nil
}

func (s *URLService) GetAllURLs() ([]models.URLRecord, error) {
	data, e := s.URLStore.URL_GetAll()

	if e != nil {
		return nil, e
	}

	return data, nil
}

func (s *URLService) UpdateURL(r *models.URLRecord) error {
	e := s.URLStore.UpdateURLInfo(*r)
	if e != nil {
		return e
	}
	return nil
}
