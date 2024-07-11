package poke

import (
	"errors"
	"fmt"

	"github.com/pbryan9/go_pokedex/api"
)

func cmdExplore(c *Config, args ...string) error {
	if len(args) == 0 {
		return errors.New("please provide area name (example: `explore snowpoint-temple-1f`)")
	}
	areaName := args[0]
	fmt.Printf("exploring %s...\n", areaName)

	areaPage := api.GetExplorePage(areaName, &c.Cache)
	pokeList := getPokeListFromArea(areaPage)
	for _, entry := range pokeList {
		fmt.Println(entry.Name)
	}
	return nil
}

func getPokeListFromArea(areaPage api.AreaPage) []api.Pokemon {
	pokeList := make([]api.Pokemon, 0)
	for _, encounter := range areaPage.PokemonEncounters {
		pokeList = append(pokeList, encounter.Pokemon)
	}
	return pokeList
}
