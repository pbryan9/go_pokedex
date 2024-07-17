package poke

import (
	"fmt"

	"github.com/pbryan9/go_pokedex/api"
)

type Pokedex struct {
	Pokedex map[string]api.PokemonPage
}

func (p *Pokedex) Add(pokemon api.PokemonPage) {
	if _, ok := p.Pokedex[pokemon.Name]; !ok {
		p.Pokedex[pokemon.Name] = pokemon
		fmt.Printf("Added %s to pokedex\n", pokemon.Name)
	} else {
		fmt.Printf("%s is already in pokedex!\n", pokemon.Name)
	}
}

func (p Pokedex) List() {
	for poke := range p.Pokedex {
		fmt.Println(poke)
	}
}

func NewPokedex() *Pokedex {
	dex := Pokedex{
		Pokedex: make(map[string]api.PokemonPage),
	}
	return &dex
}
