package domain

import "tonible14012002/ascenda-test-cli/core/ultils/mergeutils"

type Image struct {
	Link string `json:"link"`
	Desc string `json:"description"`
}

type Images struct {
	Rooms     []Image `json:"rooms"`
	Site      []Image `json:"site"`
	Amenities []Image `json:"amenities"`
}

type Location struct {
	Lat     *float64 `json:"lat"`
	Long    *float64 `json:"long"`
	Address string   `json:"address"`
	City    string   `json:"city"`
	Country string   `json:"country"`
}

type Amenities struct {
	General []string `json:"general"`
	Room    []string `json:"room"`
}

type Hotel struct {
	source        string   `json:"-"`
	Id            string   `json:"id"`
	DestinationId int      `json:"destination_id"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Condition     []string `json:"booking_condition"`
	Location      `json:"location"`
	Amenities     `json:"amenities"`
	Images        `json:"images"`
}

func (h Hotel) GetSource() string {
	return h.source
}
func (h *Hotel) SetSource(source string) {
	h.source = source
}

func (h Hotel) Merge(d Hotel) Hotel {
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
