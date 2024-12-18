package patagonia

import (
	"encoding/json"
	"net/http"
	"tonible14012002/ascenda-test-cli/core/domain"
	"tonible14012002/ascenda-test-cli/core/port"
	"tonible14012002/ascenda-test-cli/core/ultils/fetchutils"
)

const source = "patagonia"

type PatagoniaSuplier struct {
	url string
}

type NewPatagoniaSuplierParams struct {
	Url string
}

func New(params NewPatagoniaSuplierParams) port.Suplier {
	return &PatagoniaSuplier{
		url: params.Url,
	}
}

func (s *PatagoniaSuplier) GetHotels() ([]domain.Hotel, *domain.Error) {
	body, ferr := fetchutils.FetchJSON(s.url)
	if ferr != nil {
		return nil, ferr
	}

	var patagoniaHotels []PatagoniaHotel
	if err := json.Unmarshal(body, &patagoniaHotels); err != nil {
		return nil, domain.NewErr("Error decoding Json", http.StatusInternalServerError)
	}

	hotels := make([]domain.Hotel, 0, len(patagoniaHotels))

	for _, h := range patagoniaHotels {
		domainHotel := h.ToDomain()
		domainHotel.SetSource(source)
		hotels = append(hotels, domainHotel)
	}

	return hotels, nil
}

func (s *PatagoniaSuplier) GetSourceName() string {
	return source
}
