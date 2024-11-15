package paperfiles

import "tonible14012002/ascenda-test-cli/core/domain"

type Location struct {
	Address string `json:"address"`
	Country string `json:"country"`
}

type Amenities struct {
	General []string `json:"general"`
	Room    []string `json:"room"`
}

type Images struct {
	Rooms []Image `json:"rooms"`
	Site  []Image `json:"site"`
}

type Image struct {
	Link    string `json:"link"`
	Caption string `json:"caption"`
}

type PaperFliesHotel struct {
	HotelID           string    `json:"hotel_id"`
	DestinationID     int       `json:"destination_id"`
	HotelName         *string   `json:"hotel_name"`
	Location          Location  `json:"location"`
	Details           string    `json:"details"`
	Amenities         Amenities `json:"amenities"`
	Images            Images    `json:"images"`
	BookingConditions []string  `json:"booking_conditions"`
}

func (p PaperFliesHotel) ToDomainType() (dh domain.Hotel) {
	dh.Id = p.HotelID
	dh.DestinationId = p.DestinationID
	if p.HotelName != nil {
		dh.Name = *p.HotelName
	}
	dh.Location.Address = p.Location.Address
	dh.Location.Country = p.Location.Country
	dh.Description = p.Details
	dh.Condition = p.BookingConditions
	dh.Amenities.General = p.Amenities.General
	dh.Amenities.Room = p.Amenities.Room
	dh.Images.Rooms = make([]domain.Image, 0, len(p.Images.Rooms))
	for _, r := range p.Images.Rooms {
		dh.Images.Rooms = append(dh.Images.Rooms, domain.Image{
			Link: r.Link,
			Desc: r.Caption,
		})
	}
	dh.Images.Site = make([]domain.Image, 0, len(p.Images.Site))
	for _, s := range p.Images.Site {
		dh.Images.Site = append(dh.Images.Site, domain.Image{
			Link: s.Link,
			Desc: s.Caption,
		})
	}

	return dh
}
