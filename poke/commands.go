package poke

import (
	"bufio"
	"fmt"
	"os"
)

const (
	ExitError = Error("exit")
)

type Error string

func (e Error) Error() string {
	return string(e)
}

type command struct {
	prompt      string
	description string
	callback    func() error
}

type commands map[string]command

type callback func() error

func getCommands() commands {
	cmds := make(commands, 0)
	cmds.AddCommand(
		"exit",
		"Exit the program",
		cmdExit,
	)
	cmds.AddCommand(
		"help",
		"List of available commands and their descriptions",
		cmdHelp,
	)

	cmds.AddCommand(
		"find",
		"search for a pokemon by name",
		findByName,
	)

	return cmds
}

func cmdExit() error {
	return ExitError
}

func (cmds commands) AddCommand(prompt, desc string, cb callback) {
	cmds[prompt] = command{
		prompt:      prompt,
		description: desc,
		callback:    cb,
	}
}

func cmdHelp() error {
	commands := getCommands()
	fmt.Println("help")
	for _, c := range commands {
		fmt.Println(c.prompt)
		fmt.Println(c.description)
		fmt.Print("\n\n")
	}
	return nil
}

func findByName() error {
	fmt.Println("find pokemon by name")
	prompt := "enter name to search (back to go back) > "
	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)
		s.Scan()
		name := s.Text()
		if name == "back" {
			return nil
		}
		fmt.Printf("finding %s...\n", name)
	}
}
