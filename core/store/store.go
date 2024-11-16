package store

import (
	"tonible14012002/ascenda-test-cli/core/domain"
	"tonible14012002/ascenda-test-cli/core/port"
)

type InmemoryStore struct {
	hotels []domain.Hotel
}

func New() port.Store {
	return &InmemoryStore{
		hotels: make([]domain.Hotel, 0),
	}
}

func (s *InmemoryStore) Save(hotels []domain.Hotel) *domain.Error {
	s.hotels = hotels
	return nil
}

func (s *InmemoryStore) Find(ID string, DesID int) ([]domain.Hotel, *domain.Error) {
	return s.hotels, nil
}
