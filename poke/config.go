package poke

import "github.com/pbryan9/go_pokedex/internal/pokecache"

type Config struct {
	Next     string
	Previous string
	Cache    pokecache.PokeCache
	Pokedex  Pokedex
}
