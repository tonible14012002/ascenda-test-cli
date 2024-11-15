package patagonia

import (
	"net/http"
	"tonible14012002/ascenda-test-cli/core/domain"
	"tonible14012002/ascenda-test-cli/core/port"
)

type PatagoniaSuplier struct {
	url string
}

type NewPatagoniaSuplierParams struct {
	Url string
}

func NewSuplier(params NewPatagoniaSuplierParams) port.Suplier {
	return &PatagoniaSuplier{
		url: params.Url,
	}
}

func (s *PatagoniaSuplier) GetHotels() ([]domain.Hotel, *domain.Error) {
	_, err := http.Get(s.url)
	if err != nil {
		return nil, domain.NewErr("Error getting Patagonia hotels", http.StatusInternalServerError)
	}
	return nil, nil
}
