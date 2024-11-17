package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tonible14012002/ascenda-test-cli/core/domain"
	"tonible14012002/ascenda-test-cli/core/formater"
	"tonible14012002/ascenda-test-cli/core/port"
	"tonible14012002/ascenda-test-cli/core/service/hotel"
	"tonible14012002/ascenda-test-cli/internal/repository/inmemory"
	"tonible14012002/ascenda-test-cli/internal/suplier/acme"
	"tonible14012002/ascenda-test-cli/internal/suplier/paperfiles"
	"tonible14012002/ascenda-test-cli/internal/suplier/patagonia"
	"tonible14012002/ascenda-test-cli/logger"
)

type Params struct {
	HotelIDs       []string
	DestinationIDs []int
}

func parseParams(args []string) (Params, error) {
	if len(args) < 2 {
		return Params{}, errors.New("insufficient arguments, requires hotel_ids and destination_ids")
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

	log := logger.New()

	if len(os.Args) < 3 {
		log.Info("Usage: my_hotel_merger <hotel_ids:[]string> <destination_ids:int[]>")
		return
	}

	params, pErr := parseParams(os.Args[1:])

	if pErr != nil {
		log.Error(pErr.Error())
		return
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
	hotelRepository := inmemory.NewHotelRepository()

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
		HotelRepository: hotelRepository,
		Supliers: []port.Suplier{
			acmeSuplier,
			patagoniaSuplier,
			paperfliesSuplier,
		},
		FormatProvider: formatProvider,
	})

	hotels, err := hotelService.Filter(domain.HotelsQuery{
		HotelIDs:       params.HotelIDs,
		DestinationIDs: params.DestinationIDs,
	})

	if err != nil {
		log.Error(err.Message)
		return
	}

	jsonResults, mError := json.MarshalIndent(hotels, "", "  ")
	if mError != nil {
		log.Error(mError.Error())
		return
	}

	log.Print(string(jsonResults))
}
