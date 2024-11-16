package acme

import (
	"encoding/json"
	"net/http"
	"tonible14012002/ascenda-test-cli/core/domain"
	"tonible14012002/ascenda-test-cli/core/port"
	"tonible14012002/ascenda-test-cli/core/ultils/fetchutils"
)

type AcmeSuplier struct {
	url string
}

type NewAcmeSuplierParams struct {
	Url string
}

func New(params NewAcmeSuplierParams) port.Suplier {
	return &AcmeSuplier{
		url: params.Url,
	}
}

func (s *AcmeSuplier) GetHotels() ([]domain.Hotel, *domain.Error) {
	body, ferr := fetchutils.FetchJSON(s.url)
	if ferr != nil {
		return nil, ferr
	}

	var acmeHotels []AcmeHotel
	if err := json.Unmarshal(body, &acmeHotels); err != nil {
		return nil, domain.NewErr("Error decoding Json", http.StatusInternalServerError)
	}

	hotels := make([]domain.Hotel, 0, len(acmeHotels))

	for _, h := range acmeHotels {
		hotels = append(hotels, h.ToDomainType())
	}

	return hotels, nil
}
