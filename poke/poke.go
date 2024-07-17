package poke

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pbryan9/go_pokedex/internal/pokecache"
)

const (
	BaseAPI       = "https://pokeapi.co/api/v2"
	CacheInterval = 5 * time.Minute
)

func StartRepl() {
	prompt := "pokedex > "
	commands := getCommands()
	b := bufio.NewScanner(os.Stdin)

	config := Config{
		Next:     "",
		Previous: "",
		Cache:    *pokecache.NewCache(CacheInterval),
		Pokedex:  *NewPokedex(),
	}

	for {
		fmt.Print(prompt)
		b.Scan()
		input := ParseInput(b.Text())
		if len(input) == 0 {
			continue
		}

		cmd := input[0]
		args := input[1:]

		c, ok := commands[cmd]
		if !ok {
			fmt.Printf("invalid command: %s\nplease try again\n", input)
			continue
		}

		err := c.callback(&config, args...)
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
