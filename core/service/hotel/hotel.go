package hotel

import (
	"tonible14012002/ascenda-test-cli/core/domain"
	"tonible14012002/ascenda-test-cli/core/port"
)

type Service struct {
	supliers        []port.Suplier
	hotelRepository port.HotelRepository
	formatProvider  *port.FormatProvider
	hotelMerger     port.HotelMerger
}

type ServiceParams struct {
	Supliers []port.Suplier
	port.HotelRepository
	port.HotelMerger
}

func New(params ServiceParams) *Service {
	return &Service{
		supliers:        params.Supliers,
		hotelRepository: params.HotelRepository,
		hotelMerger:     params.HotelMerger,
	}
}

func (s *Service) Filter(q domain.HotelsQuery) ([]domain.Hotel, *domain.Error) {
	unmergedHotels, err := s.fetchUnmerged()

	hotels := s.hotelMerger.MergeHotels(unmergedHotels)

	if err != nil {
		return nil, err
	}

	s.hotelRepository.Save(hotels)
	return s.hotelRepository.List(q), nil
}

func (s *Service) fetchUnmerged() ([]domain.Hotel, *domain.Error) {
	hotels := make([]domain.Hotel, 0)
	for _, s := range s.supliers {
		hs, err := s.GetHotels()
		if err != nil {
			return nil, err
		}
		hotels = append(hotels, hs...)
	}
	return hotels, nil
}
