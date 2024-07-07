package poke

import (
	"bufio"
	"fmt"
	"os"
)

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

	cmdMap, cmdMapb := cmdMapClosure()

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

	cmds.AddCommand(
		"find",
		"search for a pokemon by name",
		findByName,
	)

	return cmds
}

func cmdExit() error {
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

func cmdHelp() error {
	commands := getCommands()
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
