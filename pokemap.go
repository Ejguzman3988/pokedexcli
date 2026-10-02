package main

import (
	"fmt"
)

func mapCmd(command string) func(con *config, args []string) error {

	callback := func(con *config, args []string) error {
		baseUrl := "https://pokeapi.co/api/v2"
		baseParams := "?offset=0&limit=20"
		url := baseUrl + "/location-area" + baseParams

		if command == "listNext" && con.httpClient.Locations.Next != "" {
			url = con.httpClient.Locations.Next
		}

		if command == "listPrev" && con.httpClient.Locations.Prev != "" {
			url = con.httpClient.Locations.Prev
		}

		if command == "explore" {
			if len(args) == 0 {
				errMsg := fmt.Sprintf("No Args Passed, try again")
				fmt.Println(errMsg)
				return fmt.Errorf(errMsg)

			}
			area := args[0]
			url = baseUrl + "/location-area" + fmt.Sprintf("/%v", area) + baseParams

			pokemons, err := con.httpClient.GetEncounters(url)

			if len(pokemons.Results) == 0 {
				errMsg := fmt.Errorf("No results found for: \n%v\n", url)
				fmt.Println(errMsg)
				return errMsg
			}

			if err != nil {
				return err
			}

			fmt.Printf("Exploring %v...\n", area)
			if len(pokemons.Results) == 0 {
				errMsg := fmt.Errorf("No results found for: \n%v\n", url)
				fmt.Println(errMsg)
				return errMsg
			}
			fmt.Println("Found Pokemon:")
			for _, poke := range pokemons.Results {
				fmt.Println("- ", poke.Pokemon.Name)
			}

			return nil
		}

		locations, err := con.httpClient.GetLocations(url)

		if len(locations.Results) == 0 {
			errMsg := fmt.Errorf("No results found for: \n%v\n", url)
			fmt.Println(errMsg)
			return errMsg
		}

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
