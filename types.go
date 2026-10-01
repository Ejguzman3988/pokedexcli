package main

import (
	. "github.com/ejguzman3988/pokedexcli/internal/httpClient"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	commands   map[string]cliCommand
	httpClient Client
}
