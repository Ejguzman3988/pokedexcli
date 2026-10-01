package main

import (
	. "github.com/ejguzman3988/pokedexcli/internal/httpClient"
	"time"
)

func main() {

	con := &config{
		commands:   getCommands(),
		httpClient: NewClient(5*time.Second, 5*time.Minute),
	}
	repl(con)
}
