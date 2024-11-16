package port

import "tonible14012002/ascenda-test-cli/core/domain"

type Suplier interface {
	GetHotels() ([]domain.Hotel, *domain.Error)
	GetSourceName() string
}
