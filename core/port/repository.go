package port

import "tonible14012002/ascenda-test-cli/core/domain"

type HotelRepository interface {
	Save([]domain.Hotel) *domain.Error
	List(listQuery domain.HotelsQuery) []domain.Hotel
}
