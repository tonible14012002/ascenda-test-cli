package acme

import (
	"encoding/json"
	"io"
	"net/http"
	"tonible14012002/ascenda-test-cli/core/domain"
	"tonible14012002/ascenda-test-cli/core/port"
)

type AcmeSuplier struct {
	url string
}

type NewAcmeSuplierParams struct {
	Url string
}

func NewSuplier(params NewAcmeSuplierParams) port.Suplier {
	return &AcmeSuplier{
		url: params.Url,
	}
}

func (s *AcmeSuplier) GetHotels() ([]domain.Hotel, *domain.Error) {
	resp, err := http.Get(s.url)
	if err != nil {
		return nil, domain.NewErr(err.Error(), http.StatusInternalServerError)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, domain.NewErr("Bad request", http.StatusBadRequest)
	}

	body, err := io.ReadAll(resp.Body) // Read the response body
	if err != nil {
		return nil, domain.NewErr("Error reading response body", http.StatusInternalServerError)
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
