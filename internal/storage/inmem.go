package storage

import (
	"errors"
	"maps"
	"slices"

	"github.com/Tyler-Arciniaga/SWESniper/internal/models"
)

type InMemStore struct {
	URLTable       map[string]models.URLRecord
	ChangeLogTable map[string][]models.ChangeLog
}

// URLStore interface methods
func (s *InMemStore) SaveURL(r models.URLRecord) error {
	_, exists := s.URLTable[r.URL]
	if exists {
		return errors.New("URL already existed in DB")
	}

	s.URLTable[r.URL] = r
	return nil

}

func (s *InMemStore) UpdateURLInfo(r models.URLRecord) error {
	s.URLTable[r.URL] = r
	return nil
}

func (s *InMemStore) URL_GetAll() ([]models.URLRecord, error) {
	dataAsSlice := slices.Collect(maps.Values(s.URLTable))
	if dataAsSlice == nil {
		return nil, errors.New("currently have no URLS in database")
	}
	return dataAsSlice, nil
}

// ChangeLogStore interface methods

func (s *InMemStore) LogURLChange(l models.ChangeLog) error {
	urlLog := s.ChangeLogTable[l.URL]
	urlLog = append(urlLog, l)
	s.ChangeLogTable[l.URL] = urlLog
	return nil
}
