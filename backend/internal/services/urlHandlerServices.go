package services

import (
	"fmt"
	"strconv"
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
	URL_GetOne(urlID int) (models.URLRecord, error)
	URL_Delete(urlID int) error
}

// validates add URL POST request
func (s *URLService) ValidateURLPost(r *models.AddURLRequest) error {
	if r.CheckInterval < 60 || r.CheckInterval > 86400 {
		return fmt.Errorf("check interval must be between 1 minute and 1 day")
	} //reject check intervals that are greater than a day or less than 1 minute

	return nil
}

func (s *URLService) StoreURL(r *models.AddURLRequest, u *models.User) error {
	c := time.Now()
	urlRecord := models.URLRecord{
		URL:           r.URL,
		User_id:       u.Id,
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

func (s *URLService) GetURLById(urlID string) (models.URLRecord, error) {
	int_id, err := strconv.Atoi(urlID)
	if err != nil {
		return models.URLRecord{}, fmt.Errorf("invalid url id parameter")
	}

	data, e := s.URLStore.URL_GetOne(int_id)

	if e != nil {
		return models.URLRecord{}, fmt.Errorf("error getting url by id in postgres: %v", e)
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

func (s *URLService) DeleteURL(urlID string) error {
	int_id, err := strconv.Atoi(urlID)

	if err != nil {
		return fmt.Errorf("invalid url id parameter")
	}
	e := s.URLStore.URL_Delete(int_id)

	if e != nil {
		return e
	}

	return nil
}
