package main

import (
	"encoding/json"
	"fmt"
	"tonible14012002/ascenda-test-cli/core/port"
	"tonible14012002/ascenda-test-cli/core/service/hotel"
	"tonible14012002/ascenda-test-cli/core/suplier/acme"
	"tonible14012002/ascenda-test-cli/core/suplier/patagonia"
)

func main() {

	acmeSuplier := acme.NewSuplier(
		acme.NewAcmeSuplierParams{
			Url: "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/acme",
		},
	)

	patagoniaSuplier := patagonia.NewSuplier(
		patagonia.NewPatagoniaSuplierParams{
			Url: "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/patagonia",
		},
	)

	hotelService := hotel.New(hotel.ServiceParams{
		Supliers: []port.Suplier{
			acmeSuplier,
			patagoniaSuplier,
		},
	})

	hotels, err := hotelService.Get("hotelId", "destinationId")
	if err != nil {
		panic(err)
	}

	jsonResults, _ := json.MarshalIndent(hotels, "", "  ")
	fmt.Println(string(jsonResults))
}
