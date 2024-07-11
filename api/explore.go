package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pbryan9/go_pokedex/internal/pokecache"
)

func GetExplorePage(areaName string, cache *pokecache.PokeCache) AreaPage {
	url := fmt.Sprintf("%s/location-area/%s", BaseURL, areaName)
	cacheData, hit := cache.Get(url)
	if hit {
		return convertExploreJson(cacheData)
	}

	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("error contacting endpoint %s: %s", url, err)
		return AreaPage{}
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error reading response at endpoing %s", url)
		return AreaPage{}
	}
	return convertExploreJson(body)
}

func convertExploreJson(data []byte) AreaPage {
	explorePage := AreaPage{}
	err := json.Unmarshal(data, &explorePage)
	if err != nil {
		fmt.Printf("error unmashaling explore page json: %s", err)
	}
	return explorePage
}
