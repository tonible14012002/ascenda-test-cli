package acme

import (
	"encoding/json"
	"fmt"
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
		fmt.Println("Error reading response body:", err)
		return nil, domain.NewErr("Error decoding Json", http.StatusInternalServerError)
	}
	hotels := make([]domain.Hotel, 0, len(acmeHotels))
	for _, h := range acmeHotels {
		hotels = append(hotels, acmeHotelToDomain(h))
	}

	return hotels, nil
}

func acmeHotelToDomain(h AcmeHotel) domain.Hotel {
	return domain.Hotel{
		Id:            h.ID,
		Name:          h.Name,
		DestinationId: h.DestinationID,
		Description:   h.Description,
		Location: domain.Location{
			Long:    h.Longitude,
			Lat:     h.Latitude,
			Address: h.Address,
			City:    h.City,
			Country: h.Country,
		},
		Amenities: domain.Amenities{
			General: make([]string, 0),
			Room:    make([]string, 0),
		},
		Images: domain.Images{
			Site:      make([]domain.Image, 0),
			Rooms:     make([]domain.Image, 0),
			Amenities: make([]domain.Image, 0),
		},
		Condition: make([]string, 0),
	}
}
