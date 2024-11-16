package main

import (
	"encoding/json"
	"fmt"
	"tonible14012002/ascenda-test-cli/core/port"
	"tonible14012002/ascenda-test-cli/core/service/hotel"
	"tonible14012002/ascenda-test-cli/core/store"
	"tonible14012002/ascenda-test-cli/core/suplier/acme"
	"tonible14012002/ascenda-test-cli/core/suplier/paperfiles"
	"tonible14012002/ascenda-test-cli/core/suplier/patagonia"
)

func main() {

	// init supliers
	acmeSuplier := acme.New(
		acme.NewAcmeSuplierParams{
			Url: "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/acme",
		},
	)

	patagoniaSuplier := patagonia.New(
		patagonia.NewPatagoniaSuplierParams{
			Url: "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/patagonia",
		},
	)

	paperfliesSuplier := paperfiles.New(
		paperfiles.NewPaperFliesSuplierParams{
			Url: "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/paperflies",
		},
	)

	// init store
	store := store.New()

	// service
	hotelService := hotel.New(hotel.ServiceParams{
		Store: store,
		Supliers: []port.Suplier{
			acmeSuplier,
			patagoniaSuplier,
			paperfliesSuplier,
			// add more supliers here
		},
	})

	hotels, err := hotelService.Get("hotelId", 1231)
	if err != nil {
		panic(err)
	}

	jsonResults, _ := json.MarshalIndent(hotels, "", "  ")
	fmt.Println(string(jsonResults))
}
