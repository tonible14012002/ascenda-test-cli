package acme

type AcmeHotel struct {
	ID            string   `json:"Id"`
	DestinationID int      `json:"DestinationId"`
	Name          string   `json:"Name"`
	Latitude      float64  `json:"Latitude,string"`
	Longitude     float64  `json:"Longitude,string"`
	Address       string   `json:"Address"`
	City          string   `json:"City"`
	Country       string   `json:"Country"`
	PostalCode    string   `json:"PostalCode"`
	Description   string   `json:"Description"`
	Facilities    []string `json:"Facilities"`
}
