package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tonible14012002/ascenda-test-cli/core/formater"
	"tonible14012002/ascenda-test-cli/core/port"
	"tonible14012002/ascenda-test-cli/core/service/hotel"
	"tonible14012002/ascenda-test-cli/core/store"
	"tonible14012002/ascenda-test-cli/core/suplier/acme"
	"tonible14012002/ascenda-test-cli/core/suplier/paperfiles"
	"tonible14012002/ascenda-test-cli/core/suplier/patagonia"
)

type Params struct {
	HotelIDs       []string
	DestinationIDs []int
}

func parseParams(args []string) (Params, error) {
	if len(args) < 2 {
		return Params{}, fmt.Errorf("insufficient arguments: requires hotel_ids and destination_ids")
	}
	var hotelIDs []string
	if args[0] == "none" {
		hotelIDs = nil
	} else {
		hotelIDs = parseList(args[0])
	}
	var destinationIDs []string
	if len(args[1]) == 0 {
		destinationIDs = nil
	} else {
		destinationIDs = parseList(args[1])
	}
	desIDs := make([]int, 0, len(destinationIDs))

	for _, idStr := range destinationIDs {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return Params{}, fmt.Errorf("invalid destination id: %s", idStr)
		}
		desIDs = append(desIDs, id)
	}

	return Params{
		HotelIDs:       hotelIDs,
		DestinationIDs: desIDs,
	}, nil
}

func parseList(input string) []string {
	if input == "none" {
		return nil
	}
	return strings.Split(input, ",")
}

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: my_hotel_merger <hotel_ids:[]string> <destination_ids:int[]>")
		return
	}

	params, pErr := parseParams(os.Args[1:])
	if pErr != nil {
		fmt.Println("Error:", pErr)
		os.Exit(1)
	}

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

	formatProvider := formater.NewFormatProvider([]port.Formatter{
		formater.NewCapitalizeInfoFormatter([]string{
			acmeSuplier.GetSourceName(),
			patagoniaSuplier.GetSourceName(),
			paperfliesSuplier.GetSourceName(),
		}),
		formater.NewPascalToSentenceFormatter([]string{
			acmeSuplier.GetSourceName(),
			patagoniaSuplier.GetSourceName(),
		}),
		formater.NewDescFormatter([]string{
			acmeSuplier.GetSourceName(),
			patagoniaSuplier.GetSourceName(),
			paperfliesSuplier.GetSourceName(),
		}),
	})

	// service
	hotelService := hotel.New(hotel.ServiceParams{
		Store: store,
		Supliers: []port.Suplier{
			acmeSuplier,
			patagoniaSuplier,
			paperfliesSuplier,
		},
		FormatProvider: formatProvider,
	})

	hotels, err := hotelService.Filter(params.HotelIDs, params.DestinationIDs)
	if err != nil {
		panic(err)
	}

	jsonResults, _ := json.MarshalIndent(hotels, "", "  ")
	fmt.Println(string(jsonResults))
}
