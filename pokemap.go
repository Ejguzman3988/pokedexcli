package main

import (
	"fmt"
)

func mapExplore(con *config) error {
	return nil
}

func mapCmd(direction string) func(con *config) error {

	callback := func(con *config) error {
		url := "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
		if direction == "next" && con.httpClient.Locations.Next != "" {
			url = con.httpClient.Locations.Next
		}
		if direction == "prev" && con.httpClient.Locations.Prev != "" {
			url = con.httpClient.Locations.Prev
		}

		locations, err := con.httpClient.GetLocations(url)

		if err != nil {
			return err
		}

		for _, lo := range locations.Results {
			fmt.Println(lo.Name)
		}

		return nil

	}

	return callback
}
