package poke

import (
	"errors"
	"fmt"
)

func cmdInspect(c *Config, args ...string) error {
	if len(args) < 1 {
		return errors.New("please specify pokemon to inspect")
	}

	name := args[0]

	pokemon, ok := c.Pokedex[name]
	if !ok {
		return errors.New(fmt.Sprintf("%s is not in pokedex", name))
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokeType := range pokemon.Types {
		fmt.Printf("  - %s\n", pokeType.Type.Name)
	}
	return nil
}
