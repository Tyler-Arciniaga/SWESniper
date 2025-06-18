package storage

import (
	"errors"

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
