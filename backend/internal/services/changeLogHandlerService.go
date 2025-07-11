package services

import (
	"fmt"
	"strconv"

	"github.com/Tyler-Arciniaga/SWESniper/internal/models"
)

type ChangeLogService struct {
	ChangeRepository ChangeLogStore
}

type ChangeLogStore interface {
	LogURLChange(l models.ChangeRecord) error
	ChangeLog_GetAll(u models.User) ([]models.ChangeRecord, error)
	ChangeLog_GetOneUrl(u models.User, urlID int) ([]models.ChangeRecord, error)
}

func (s *ChangeLogService) PersistChangeRecord(r *models.ChangeRecord) error {
	e := s.ChangeRepository.LogURLChange(*r)
	if e != nil {
		return e
	}
	return nil
}

func (s *ChangeLogService) GetAllChangeRecords(u models.User) ([]models.ChangeRecord, error) {
	data, e := s.ChangeRepository.ChangeLog_GetAll(u)
	if e != nil {
		return nil, e
	}

	return data, nil
}

func (s *ChangeLogService) GetOneUrlChangeRecord(u models.User, urlID string) ([]models.ChangeRecord, error) {
	int_id, err := strconv.Atoi(urlID)
	if err != nil {
		return nil, fmt.Errorf("invalid url id parameter")
	}

	ChangeRecords, e := s.ChangeRepository.ChangeLog_GetOneUrl(u, int_id)
	if e != nil {
		return nil, fmt.Errorf("error getting change record from postgres DB: %v", e)
	}

	return ChangeRecords, nil
}
