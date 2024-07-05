package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {
	const prompt string = "pokedex > "

	s := bufio.NewScanner(os.Stdin)
	for {
		cmd := getInput(prompt, s)
		if cmd == "exit" {
			fmt.Println("ok byeeee!")
			return
		}
		err := runCommand(cmd)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cmdHelp() error {
	return nil
}

func getInput(prompt string, s *bufio.Scanner) string {
	fmt.Print(prompt)
	s.Scan()
	return s.Text()
}

func runCommand(c string) error {
	if _, ok := commands[c]; !ok {
		return errors.New(fmt.Sprintf("unrecognized command: %v", c))
	}
	commands[c].cb()
	return nil
}

type command struct {
	name        string
	description string
	cb          func() error
}

var commands map[string]command = map[string]command{
	"help": {
		name:        "help",
		description: "Displays this help message.",
		cb:          cmdHelp,
	},
}
