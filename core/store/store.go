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

func (s *InmemoryStore) List(ids []string, desIds []int) []domain.Hotel {
	var filtered []domain.Hotel

	// If no hotelIDs and destinationIDs are provided, return all hotels
	if len(ids) == 0 && len(desIds) == 0 {
		return s.hotels
	}

	// // Filter hotels
	for _, hotel := range s.hotels {
		if len(ids) > 0 && !contains(ids, hotel.Id) {
			continue
		}
		if len(desIds) > 0 && !contains(desIds, hotel.DestinationId) {
			continue
		}
		filtered = append(filtered, hotel)
	}

	return filtered
}

func contains[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}
