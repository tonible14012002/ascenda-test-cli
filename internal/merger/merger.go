package merger

import (
	"tonible14012002/ascenda-test-cli/core/domain"
	"tonible14012002/ascenda-test-cli/core/port"
	"tonible14012002/ascenda-test-cli/core/ultils/mergeutils"
)

type HotelSimpleMerger struct {
	formatProvider *port.FormatProvider
}

type HotelMergerParams struct {
	*port.FormatProvider
}

func NewHotelSimpleMerger(params HotelMergerParams) *HotelSimpleMerger {
	return &HotelSimpleMerger{
		formatProvider: params.FormatProvider,
	}
}

func (m *HotelSimpleMerger) MergeHotels(pipe []domain.Hotel) []domain.Hotel {
	hotelMapper := make(map[string]domain.Hotel)

	// Merge matching hotels
	for _, h := range pipe {
		if _, existed := hotelMapper[h.Id]; existed {
			hotelMapper[h.Id] = m.Merge(hotelMapper[h.Id], m.formatProvider.Format(h))
		} else {
			hotelMapper[h.Id] = m.formatProvider.Format(h)
		}
	}

	// Convert to slice
	merged := make([]domain.Hotel, 0, len(hotelMapper))
	for _, h := range hotelMapper {
		merged = append(merged, h)
	}

	return merged
}

func (m *HotelSimpleMerger) Merge(h domain.Hotel, d domain.Hotel) domain.Hotel {
	h.Name = mergeutils.PickLongerStr(h.Name, d.Name)
	h.Address = mergeutils.PickLongerStr(h.Address, d.Address)
	h.City = mergeutils.PickLongerStr(h.City, d.City)
	h.Description = mergeutils.PickLongerStr(h.Description, d.Description)

	h.Location.Lat = mergeutils.FirstNonNil(h.Location.Lat, d.Location.Lat)
	h.Location.Long = mergeutils.FirstNonNil(h.Location.Long, d.Location.Long)
	h.Location.Address = mergeutils.PickLongerStr(h.Location.Address, d.Location.Address)
	h.Location.City = mergeutils.PickLongerStr(h.Location.City, d.Location.City)
	h.Location.Country = mergeutils.PickLongerStr(h.Location.Country, d.Location.Country)

	h.Condition = mergeutils.PickLongerSlice(h.Condition, d.Condition)

	h.Amenities.General = mergeutils.PickLongerSlice(h.Amenities.General, d.Amenities.General)
	h.Amenities.Room = mergeutils.PickLongerSlice(h.Amenities.Room, d.Amenities.Room)

	h.Images.Rooms = append(h.Images.Rooms, d.Images.Rooms...)
	h.Images.Site = append(h.Images.Site, d.Images.Site...)
	h.Images.Amenities = append(h.Images.Amenities, d.Images.Amenities...)

	return h
}
