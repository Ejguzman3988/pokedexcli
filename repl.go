package main

import (
	"bufio"
	"fmt"
	"os"
)

func repl(con *config) {

	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()

		command, ok := con.commands[text]
		if ok {
			command.callback(con)
		} else {
			fmt.Println("Unknown command")
		}
		// 	fmt.Println("Your command was: %v\n", text)
	}
}
