package poke

import (
	"fmt"
	"github.com/pbryan9/go_pokedex/api"
)

func cmdCatch(c *Config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("please specify a pokemon to catch")
		return nil
	}

	name := args[0]
	api.Catch(name, &c.Cache)
	return nil
}