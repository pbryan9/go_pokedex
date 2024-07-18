package poke

import (
	"fmt"

	"github.com/pbryan9/go_pokedex/api"
)

type Pokedex map[string]api.PokemonPage

func (p Pokedex) Add(pokemon api.PokemonPage) {
	if _, ok := p[pokemon.Name]; !ok {
		p[pokemon.Name] = pokemon
		fmt.Printf("Added %s to pokedex\n", pokemon.Name)
	} else {
		fmt.Printf("%s is already in pokedex!\n", pokemon.Name)
	}
}

func (p Pokedex) List() {
	for poke := range p {
		fmt.Println(poke)
	}
}

func (p Pokedex) Check(name string) bool {
	_, ok := p[name]
	return ok
}

func NewPokedex() Pokedex {
	dex := make(map[string]api.PokemonPage)
	return dex
}
