package poke

import (
	"bufio"
	"fmt"
	"os"
)

func StartRepl() {
	prompt := "pokedex > "
	commands := getCommands()
	b := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)
		b.Scan()
		input := b.Text()

		c, ok := commands[input]
		if !ok {
			fmt.Printf("invalid command: %s\nplease try again\n", input)
			continue
		}

		err := c.callback()
		if err != nil {
			switch err {
			case ExitError:
				fmt.Println("goodbye!")
				return
			}
		}
	}
}
