package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

func commandExit(con *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(con *config, args []string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range con.commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func getPokemon(con *config, args []string) error {
	baseUrl := "https://pokeapi.co/api/v2"
	baseParams := ""
	endPoint := "/pokemon"

	if len(args) == 0 {
		errMsg := fmt.Errorf("No Args Passed, try again")
		fmt.Println(errMsg)
		return errMsg

	}

	pokemonArg := fmt.Sprintf("/%v", args[0])
	url := baseUrl + endPoint + pokemonArg + baseParams
	p, exists := con.pokedex[pokemonArg[1:]]

	if exists {
		fmt.Printf("%v is in the pokedex. :) \n", p.Name)
		return nil
	}

	pokemon, err := con.httpClient.GetPokemon(url)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)

	expFloat := float64(pokemon.BaseExperience)

	const maxExp = 340.0
	chance := 1 - math.Min(expFloat/maxExp, 1.0)
	chance += (rand.Float64() - 0.5) * 0.2 // ±10%
	chance = math.Max(0.05, math.Min(0.95, chance))

	if rand.Float64() < chance {
		fmt.Printf("%v was caught!\n", pokemon.Name)
		con.pokedex[pokemon.Name] = pokemon

	} else {
		fmt.Printf("%v escaped!\n", pokemon.Name)
	}

	return nil
}

func findPokemon(con *config, args []string) error {
	pokemon, exists := con.pokedex[args[0]]
	if !exists {
		fmt.Println("you have not caught that pokemon")
	} else {
		fmt.Printf("Name: %v\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%v: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf("  - %v\n", t.Type.Name)
		}
	}

	return nil
}

func getPokedex(con *config, args []string) error {
	if len(con.pokedex) < 1 {
		errMsg := fmt.Errorf("You have not caught any pokemon")
		fmt.Println(errMsg)
		return errMsg
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range con.pokedex {
		fmt.Printf("- %v\n", pokemon.Name)
	}

	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon World.",
			callback:    mapCmd("listNext"),
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of 20 location areas in the Pokemon World.",
			callback:    mapCmd("listPrev"),
		},
		"explore": {
			name:        "explore",
			description: "Add location name argument to the explore cmd to list pokemon in that area",
			callback:    mapCmd("explore"),
		},
		"catch": {
			name:        "catch",
			description: "Add pokemon name argument to the explore cmd to list pokemon in that area",
			callback:    getPokemon,
		},
		"inspect": {
			name:        "inspect",
			description: "If caught view information on pokemon",
			callback:    findPokemon,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all pokemon caught and in your pokedex",
			callback:    getPokedex,
		},
	}
}
