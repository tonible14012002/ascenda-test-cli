package hotel

import (
	"tonible14012002/ascenda-test-cli/core/domain"
	"tonible14012002/ascenda-test-cli/core/port"
	"tonible14012002/ascenda-test-cli/core/ultils/mergeutils"
)

type Service struct {
	supliers []port.Suplier
	store    port.Store
}

type ServiceParams struct {
	Supliers []port.Suplier
	Store    port.Store
}

func New(params ServiceParams) Service {
	return Service{
		supliers: params.Supliers,
		store:    params.Store,
	}
}

func (s Service) Get(hotelId string, destinationId int) ([]domain.Hotel, *domain.Error) {
	unmergedHotels, err := s.fetchUnmerged()

	hotels := s.mergeHotelsByID(unmergedHotels)

	if err != nil {
		return nil, err
	}

	s.store.Save(hotels)
	return s.store.Find(hotelId, destinationId)
}

func (s Service) fetchUnmerged() ([]domain.Hotel, *domain.Error) {
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

func (s Service) mergeHotelsByID(unmergedHotels []domain.Hotel) []domain.Hotel {
	hotelMapper := make(map[string]domain.Hotel)

	// Merge matching hotels
	for _, h := range unmergedHotels {
		if _, ok := hotelMapper[h.Id]; ok {
			hotelMapper[h.Id] = s.merge(hotelMapper[h.Id], h)
		} else {
			hotelMapper[h.Id] = h
		}
	}

	// Convert to slice
	merged := make([]domain.Hotel, 0, len(hotelMapper))
	for _, h := range hotelMapper {
		merged = append(merged, h)
	}

	return merged
}

func (s Service) merge(base domain.Hotel, h domain.Hotel) domain.Hotel {
	// Ignore DestinationId
	capitalize := mergeutils.CapitalizeFirstLetters

	base.Name = capitalize(mergeutils.PickLongerStr(base.Name, h.Name))
	base.Address = capitalize(mergeutils.PickLongerStr(base.Address, h.Address))
	base.City = capitalize(mergeutils.PickLongerStr(base.City, h.City))
	base.Description = capitalize(mergeutils.PickLongerStr(base.Description, h.Description))

	base.Location.Lat = mergeutils.FirstNonNil(base.Location.Lat, h.Location.Lat)
	base.Location.Long = mergeutils.FirstNonNil(base.Location.Long, h.Location.Long)
	base.Location.Address = capitalize(mergeutils.PickLongerStr(base.Location.Address, h.Location.Address))
	base.Location.City = capitalize(mergeutils.PickLongerStr(base.Location.City, h.Location.City))
	base.Location.Country = capitalize(mergeutils.PickLongerStr(base.Location.Country, h.Location.Country))

	base.Condition = mergeutils.PickLongerSlice(base.Condition, h.Condition)
	for i := 0; i < len(base.Condition); i++ {
		h.Condition[i] = mergeutils.RemoveRedundantSpaces(h.Condition[i])
	}

	base.Amenities.General = mergeutils.PickLongerSlice(base.Amenities.General, h.Amenities.General)

	for i := 0; i < len(base.Amenities.General); i++ {
		base.Amenities.General[i] = mergeutils.RemoveRedundantSpaces(base.Amenities.General[i])
	}

	base.Amenities.Room = mergeutils.PickLongerSlice(base.Amenities.Room, h.Amenities.Room)

	for i := 0; i < len(base.Amenities.Room); i++ {
		base.Amenities.Room[i] = mergeutils.RemoveRedundantSpaces(base.Amenities.Room[i])
	}

	base.Images.Rooms = append(base.Images.Rooms, h.Images.Rooms...)
	base.Images.Site = append(base.Images.Site, h.Images.Site...)
	base.Images.Amenities = append(base.Images.Amenities, h.Images.Amenities...)

	return base
}
