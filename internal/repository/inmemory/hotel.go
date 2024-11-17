package inmemory

import (
	"tonible14012002/ascenda-test-cli/core/domain"
	"tonible14012002/ascenda-test-cli/core/port"
)

type HotelRepository struct {
	hotels []domain.Hotel
}

func NewHotelRepository() port.HotelRepository {
	return &HotelRepository{
		hotels: make([]domain.Hotel, 0),
	}
}

func (s *HotelRepository) Save(hotels []domain.Hotel) *domain.Error {
	s.hotels = hotels
	return nil
}

func (s *HotelRepository) List(q domain.HotelsQuery) []domain.Hotel {
	filtered := make([]domain.Hotel, 0)

	// If no hotelIDs and destinationIDs are provided, return all hotels
	if len(q.HotelIDs) == 0 && len(q.DestinationIDs) == 0 {
		return s.hotels
	}

	// // Filter hotels
	for _, hotel := range s.hotels {
		if len(q.HotelIDs) > 0 && !contains(q.HotelIDs, hotel.Id) {
			continue
		}
		if len(q.DestinationIDs) > 0 && !contains(q.DestinationIDs, hotel.DestinationId) {
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
