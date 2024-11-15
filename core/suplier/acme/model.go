package acme

import (
	"tonible14012002/ascenda-test-cli/core/domain"
)

type AcmeHotel struct {
	ID            string           `json:"Id"`
	DestinationID int              `json:"DestinationId"`
	Name          *string          `json:"Name,omitempty"`
	Latitude      domain.JsonFloat `json:"Latitude"`
	Longitude     domain.JsonFloat `json:"Longitude"`
	Address       *string          `json:"Address,omitempty"`
	City          *string          `json:"City,omitempty"`
	Country       *string          `json:"Country,omitempty"`
	PostalCode    *string          `json:"PostalCode,omitempty"`
	Description   *string          `json:"Description,omitempty"`
	Facilities    []string         `json:"Facilities,omitempty"`
}

func (h AcmeHotel) ToDomainType() (dh domain.Hotel) {
	// sanitized field
	dh.Id = h.ID
	dh.DestinationId = h.DestinationID

	// nullable
	if h.Name != nil {
		dh.Name = *h.Name
	}
	if h.Description != nil {
		dh.Description = *h.Description
	}
	if h.Latitude.IsValid {
		dh.Location.Lat = &h.Latitude.Value
	}
	if h.Longitude.IsValid {
		dh.Location.Long = &h.Longitude.Value
	}
	if h.Address != nil {
		dh.Location.Address = *h.Address
	}
	if h.City != nil {
		dh.Location.City = *h.City
	}
	if h.Country != nil {
		dh.Location.Country = *h.Country
	}
	dh.DestinationId = h.DestinationID
	dh.Amenities.General = h.Facilities
	return dh
}
