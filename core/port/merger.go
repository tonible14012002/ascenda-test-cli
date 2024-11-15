package port

import "tonible14012002/ascenda-test-cli/core/domain"

type Merger interface {
	Merge(hotels []domain.Hotel) []domain.Hotel
}
