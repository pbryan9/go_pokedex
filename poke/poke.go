package poke

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartRepl() {
	prompt := "pokedex > "
	commands := getCommands()
	b := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)
		b.Scan()
		input := ParseInput(b.Text())
		if len(input) == 0 {
			continue
		}

		cmd := input[0]

		c, ok := commands[cmd]
		if !ok {
			fmt.Printf("invalid command: %s\nplease try again\n", input)
			continue
		}

		err := c.callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func ParseInput(s string) []string {
	s = strings.ToLower(s)
	words := strings.Fields(s)
	return words
}
