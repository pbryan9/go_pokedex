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
	pokemon, err := api.Catch(name, &c.Cache)
	if err != nil {
		return err
	}
	// if catch failed, pokemon name will be empty
	if pokemon.Name != "" {
		c.Pokedex.Add(pokemon)
	}
	return nil
}
