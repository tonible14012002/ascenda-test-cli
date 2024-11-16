package paperfiles

import (
	"encoding/json"
	"net/http"
	"tonible14012002/ascenda-test-cli/core/domain"
	"tonible14012002/ascenda-test-cli/core/port"
	"tonible14012002/ascenda-test-cli/core/ultils/fetchutils"
)

const source = "paperflies"

type PaperFliesHotelSuplier struct {
	url string
}

type NewPaperFliesSuplierParams struct {
	Url string
}

func New(params NewPaperFliesSuplierParams) port.Suplier {
	return &PaperFliesHotelSuplier{
		url: params.Url,
	}
}

func (s *PaperFliesHotelSuplier) GetHotels() ([]domain.Hotel, *domain.Error) {
	body, ferr := fetchutils.FetchJSON(s.url)
	if ferr != nil {
		return nil, ferr
	}

	var paperfilesHotels []PaperFliesHotel
	if err := json.Unmarshal(body, &paperfilesHotels); err != nil {
		return nil, domain.NewErr("Error decoding Json", http.StatusInternalServerError)
	}

	hotels := make([]domain.Hotel, 0, len(paperfilesHotels))

	for _, h := range paperfilesHotels {
		domainHotel := h.ToDomain()
		domainHotel.SetSource(source)
		hotels = append(hotels, domainHotel)
	}

	return hotels, nil
}

func (s *PaperFliesHotelSuplier) GetSourceName() string {
	return source
}
