package port

import "tonible14012002/ascenda-test-cli/core/domain"

type HotelMerger interface {
	MergeHotels(hotels []domain.Hotel) []domain.Hotel
}
