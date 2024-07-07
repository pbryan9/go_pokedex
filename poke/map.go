package poke

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func cmdMapClosure() (cmdMap, cmdMapb callback) {
	page := 0
	cmdMap = func(config *Config) error {
		if config.Next == "" {
			config.Next = BaseAPI + "/location-area"
		}
		page++
		areas := getMapPage(config.Next)
		config.Next = areas.Next
		config.Previous = areas.Previous
		for _, area := range areas.Results {
			fmt.Println(area.Name)
		}
		return nil
	}

	cmdMapb = func(config *Config) error {
		if config.Previous == "" {
			return errors.New("cannot decrement page: already at first page")
		}
		page--
		areas := getMapPage(config.Previous)
		config.Next = areas.Next
		config.Previous = areas.Previous
		for _, area := range areas.Results {
			fmt.Println(area.Name)
		}
		return nil
	}

	return cmdMap, cmdMapb
}

func getMapPage(endpoint string) AreaList {
	areas := AreaList{}

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("error contacting endpoint")
		return areas
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code %d and\nbody: %s\n", res.StatusCode, body)
		return areas
	}

	json.Unmarshal(body, &areas)

	return areas
}

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
