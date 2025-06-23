package services

import "github.com/Tyler-Arciniaga/SWESniper/internal/models"

type ChangeLogService struct {
	ChangeRepository ChangeLogStore
}

type ChangeLogStore interface {
	LogURLChange(l models.ChangeRecord) error
	ChangeLog_GetAll() ([]models.ChangeRecord, error)
}

func (s *ChangeLogService) PersistChangeRecord(r *models.ChangeRecord) error {
	e := s.ChangeRepository.LogURLChange(*r)
	if e != nil {
		return e
	}
	return nil
}

func (s *ChangeLogService) GetAllChangeRecords() ([]models.ChangeRecord, error) {
	data, e := s.ChangeRepository.ChangeLog_GetAll()
	if e != nil {
		return nil, e
	}

	return data, nil
}
