package patagonia

import "tonible14012002/ascenda-test-cli/core/domain"

type PatagoniaImage struct {
	Url         string `json:"url"`
	Description string `json:"description"`
}

type Images struct {
	Rooms     []PatagoniaImage `json:"rooms"`
	Amenities []PatagoniaImage `json:"amenities"`
}

type PatagoniaHotel struct {
	ID            string           `json:"id"`
	DestinationID int              `json:"destination"`
	Name          *string          `json:"name,omitempty"`
	Lat           domain.JsonFloat `json:"lat"`
	Long          domain.JsonFloat `json:"long"`
	Address       *string          `json:"address,omitempty"`
	Info          *string          `json:"info,omitempty"`
	Aminities     []string         `json:"aminities,omitempty"`
	Images        *Images          `json:"images,omitempty"`
}

func (p *PatagoniaHotel) ToDomain() (dh domain.Hotel) {
	dh.Id = p.ID
	dh.DestinationId = p.DestinationID
	if p.Name != nil {
		dh.Name = *p.Name
	}
	if p.Info != nil {
		dh.Description = *p.Info
	}
	if p.Lat.IsValid {
		dh.Location.Lat = &p.Lat.Value
	}
	if p.Long.IsValid {
		dh.Location.Long = &p.Long.Value
	}
	if p.Address != nil {
		dh.Location.Address = *p.Address
	}

	dh.Amenities.Room = p.Aminities

	if p.Images != nil {
		for _, room := range p.Images.Rooms {
			dh.Images.Rooms = append(dh.Images.Rooms, domain.Image{
				Link: room.Url,
				Desc: room.Description,
			})
		}
		for _, amenity := range p.Images.Amenities {
			dh.Images.Amenities = append(dh.Images.Amenities, domain.Image{
				Link: amenity.Url,
				Desc: amenity.Description,
			})
		}
	}
	return dh
}
