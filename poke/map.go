package poke

import (
	"errors"
	"fmt"
	"github.com/pbryan9/go_pokedex/api"
)

func cmdMap(config *Config) error {
	if config.Next == "" {
		config.Next = BaseAPI + "/location-area?offset=0&limit=20"
	}
	areas := api.GetMapPage(config.Next, &config.Cache)
	config.Next = areas.Next
	config.Previous = areas.Previous
	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func cmdMapb(config *Config) error {
	if config.Previous == "" {
		return errors.New("cannot decrement page: already at first page")
	}
	areas := api.GetMapPage(config.Previous, &config.Cache)
	config.Next = areas.Next
	config.Previous = areas.Previous
	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func cmdMapClosure() (cmdMap, cmdMapb callback) {
	cmdMap = func(config *Config) error {
		if config.Next == "" {
			config.Next = BaseAPI + "/location-area"
		}
		areas := api.GetMapPage(config.Next, &config.Cache)
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
		areas := api.GetMapPage(config.Previous, &config.Cache)
		config.Next = areas.Next
		config.Previous = areas.Previous
		for _, area := range areas.Results {
			fmt.Println(area.Name)
		}
		return nil
	}

	return cmdMap, cmdMapb
}
