package api

import (
	"encoding/json"
	"fmt"
	"github.com/pbryan9/go_pokedex/internal/pokecache"
	"io"
	"net/http"
)

type AreaList struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func GetMapPage(endpoint string, cache *pokecache.PokeCache) AreaList {
	val, hit := cache.Get(endpoint)
	if !hit {
		fmt.Println("cache miss")
		val = fetchMapPageFromAPI(endpoint)
		cache.Add(endpoint, val)
	} else {
		fmt.Println("cache hit!!!!!!")
	}

	areas := AreaList{}
	err := json.Unmarshal(val, &areas)
	if err != nil {
		fmt.Printf("error unmarshaling json\n")
		fmt.Print(val, "\n")
	}
	return areas
}

func fetchMapPageFromAPI(endpoint string) []byte {
	fmt.Println("~~~fetching data from PokeAPI~~~")
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("error contacting endpoint\n")
		return []byte{}
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code %d and\nbody: %s\n", res.StatusCode, body)
		return body
	}

	return body
}
