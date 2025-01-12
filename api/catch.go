package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/pbryan9/go_pokedex/internal/pokecache"
)

func Catch(name string, cache *pokecache.PokeCache) (PokemonPage, error) {
	pokePage := fetchPokemonPage(name, cache)
	if pokePage.Name == "" {
		return PokemonPage{}, fmt.Errorf("could not get data for pokemon: %s", name)
	}
	fmt.Printf("Throwing pokeball at %s...\n", pokePage.Name)
	target := pokePage.BaseExperience
	fmt.Printf("%s's base experience is %v\n", pokePage.Name, pokePage.BaseExperience)
	roll := rand.Intn(500)
	if roll >= target {
		fmt.Printf("You caught a %s!\n", pokePage.Name)
		return pokePage, nil
	} else {
		fmt.Printf("Whoops, your pokeball failed to capture %s...\n", pokePage.Name)
	}
	return PokemonPage{}, nil
}

func fetchPokemonPage(name string, cache *pokecache.PokeCache) PokemonPage {
	url := fmt.Sprintf("%s/pokemon/%s/", BaseURL, name)
	page := make([]byte, 0)
	if _, ok := cache.Get(url); !ok {
		page, err := getPokemonPageFromAPI(url)
		if err != nil {
			fmt.Printf("error fetching page %s: %s\n", url, err)
			return PokemonPage{}
		}
		cachePokemonPage(url, page, cache)
	}

	page, _ = cache.Get(url)
	pokePage := PokemonPage{}
	err := json.Unmarshal(page, &pokePage)
	if err != nil {
		fmt.Printf("error unmarshaling json for page %s\n", url)
		fmt.Println(err)
	}

	return pokePage
}

func getPokemonPageFromAPI(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error reading body")
		return []byte{}, err
	}
	pokePage := PokemonPage{}
	err = json.Unmarshal(body, &pokePage)
	if err != nil {
		fmt.Println("not sure why that didn't work...")
		return []byte{}, errors.New("didn't work")
	}

	return body, nil
}

func cachePokemonPage(url string, rawPage []byte, cache *pokecache.PokeCache) {
	cache.Add(url, rawPage)
}
