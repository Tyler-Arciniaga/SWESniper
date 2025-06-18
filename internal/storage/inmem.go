package storage

import (
	"errors"
	"maps"
	"slices"

	"github.com/Tyler-Arciniaga/SWESniper/internal/models"
)

type InMemStore struct {
	Data map[string]models.URLRecord
}

func (s *InMemStore) SaveURL(r models.URLRecord) error {
	_, exists := s.Data[r.URL]
	if exists {
		return errors.New("URL already existed in DB")
	}

	s.Data[r.URL] = r
	return nil

}

func (s *InMemStore) GetAll() ([]models.URLRecord, error) {
	dataAsSlice := slices.Collect(maps.Values(s.Data))
	if dataAsSlice == nil {
		return nil, errors.New("currently have no URLS in database")
	}
	return dataAsSlice, nil
}
