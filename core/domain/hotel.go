package domain

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

type HotelsQuery struct {
	HotelIDs       []string
	DestinationIDs []int
}

func (h Hotel) GetSource() string {
	return h.source
}
func (h *Hotel) SetSource(source string) {
	h.source = source
}
