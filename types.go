package main

import (
	. "github.com/ejguzman3988/pokedexcli/internal/httpClient"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	pokedex    map[string]PokemonResData
	commands   map[string]cliCommand
	httpClient Client
}
