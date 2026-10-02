package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl(con *config) {

	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		fullText := scanner.Text()

		textSlice := strings.Fields(fullText)

		// Define Args -- allow for no args
		var args []string
		if len(textSlice) > 1 {
			args = textSlice[1:]
		}

		var command string
		if len(textSlice) > 0 {
			command = textSlice[0]
		}

		c, ok := con.commands[command]
		if ok {
			c.callback(con, args)
		} else {
			fmt.Println("Unknown command")
		}
		// 	fmt.Println("Your command was: %v\n", text)
	}
}
