package api

import (
	"encoding/json"
	"fmt"
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

func GetMapPage(endpoint string) AreaList {
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
