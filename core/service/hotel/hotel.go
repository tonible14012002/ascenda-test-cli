package hotel

import (
	"tonible14012002/ascenda-test-cli/core/domain"
	"tonible14012002/ascenda-test-cli/core/formater"
	"tonible14012002/ascenda-test-cli/core/port"
)

type Service struct {
	supliers        []port.Suplier
	hotelRepository port.HotelRepository
	formatProvider  *formater.FormatProvider
}

type ServiceParams struct {
	Supliers []port.Suplier
	port.HotelRepository
	*formater.FormatProvider
}

func New(params ServiceParams) *Service {
	return &Service{
		supliers:        params.Supliers,
		hotelRepository: params.HotelRepository,
		formatProvider:  params.FormatProvider,
	}
}

func (s *Service) Filter(q domain.HotelsQuery) ([]domain.Hotel, *domain.Error) {
	unmergedHotels, err := s.fetchUnmerged()

	hotels := s.mergeHotelsByID(unmergedHotels)

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

func (s *Service) mergeHotelsByID(unmergedHotels []domain.Hotel) []domain.Hotel {
	hotelMapper := make(map[string]domain.Hotel)

	// Merge matching hotels
	for _, h := range unmergedHotels {
		if _, existed := hotelMapper[h.Id]; existed {
			hotelMapper[h.Id] = hotelMapper[h.Id].Merge(s.formatProvider.Format(h))
		} else {
			hotelMapper[h.Id] = s.formatProvider.Format(h)
		}
	}

	// Convert to slice
	merged := make([]domain.Hotel, 0, len(hotelMapper))
	for _, h := range hotelMapper {
		merged = append(merged, h)
	}

	return merged
}
