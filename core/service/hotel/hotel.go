package hotel

import (
	"tonible14012002/ascenda-test-cli/core/domain"
	"tonible14012002/ascenda-test-cli/core/formater"
	"tonible14012002/ascenda-test-cli/core/port"
)

type Service struct {
	supliers       []port.Suplier
	store          port.Store
	formatProvider *formater.FormatProvider
}

type ServiceParams struct {
	Supliers       []port.Suplier
	Store          port.Store
	FormatProvider *formater.FormatProvider
}

func New(params ServiceParams) *Service {
	return &Service{
		supliers:       params.Supliers,
		store:          params.Store,
		formatProvider: params.FormatProvider,
	}
}

func (s *Service) Filter(hotelIds []string, destinationIds []int) ([]domain.Hotel, *domain.Error) {
	unmergedHotels, err := s.fetchUnmerged()

	hotels := s.mergeHotelsByID(unmergedHotels)

	if err != nil {
		return nil, err
	}

	s.store.Save(hotels)
	return s.store.List(hotelIds, destinationIds), nil
}

func (s *Service) fetchUnmerged() ([]domain.Hotel, *domain.Error) {
	hotels := make([]domain.Hotel, 0)
	for _, s := range s.supliers {
		hs, err := s.GetHotels()
		if err != nil {
			continue
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
