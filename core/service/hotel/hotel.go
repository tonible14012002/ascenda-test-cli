package hotel

import (
	"fmt"
	"tonible14012002/ascenda-test-cli/core/domain"
	"tonible14012002/ascenda-test-cli/core/port"
)

type Service struct {
	supliers []port.Suplier
}

type ServiceParams struct {
	Supliers []port.Suplier
}

func New(params ServiceParams) Service {
	return Service{
		supliers: params.Supliers,
	}
}

func (s Service) Get(hotelId string, destinationId string) ([]domain.Hotel, *domain.Error) {
	hotels := make([]domain.Hotel, 0, 10)
	for _, s := range s.supliers {
		hs, err := s.GetHotels()
		if err != nil {
			fmt.Print(err)
		}
		hotels = append(hotels, hs...)
	}
	return hotels, nil
}
