package poke

import (
	"fmt"
	"os"
)

type command struct {
	prompt      string
	description string
	callback    callback
}

type Config struct {
	Next     string
	Previous string
}

type commands map[string]command

type callback func(*Config) error

func getCommands() commands {
	cmds := make(commands, 0)
	cmds.AddCommand(
		"help",
		"List of available commands and their descriptions",
		cmdHelp,
	)

	cmds.AddCommand(
		"exit",
		"Exit the program",
		cmdExit,
	)

	cmds.AddCommand(
		"map",
		"display the next 20 locations",
		cmdMap,
	)
	cmds.AddCommand(
		"mapb",
		"display the previous 20 locations",
		cmdMapb,
	)

	return cmds
}

func cmdExit(_ *Config) error {
	fmt.Println("goodbye!")
	os.Exit(0)
	return nil
}

func (cmds commands) AddCommand(prompt, desc string, cb callback) {
	cmds[prompt] = command{
		prompt:      prompt,
		description: desc,
		callback:    cb,
	}
}

func cmdHelp(_ *Config) error {
	commands := getCommands()
	for _, c := range commands {
		fmt.Println(c.prompt)
		fmt.Println(c.description)
		fmt.Print("\n\n")
	}
	return nil
}
