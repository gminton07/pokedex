package pokeapi

import (
	"encoding/json"
	"fmt"

	"github.com/gminton07/pokedex/internal/pokecache"
)

type MapAreaList struct {
	Count        int        `json:"count"`
	Next         string     `json:"next"`
	Previous     string     `json:"previous,omitempty"`
	Results      []Location `json:"Results`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Functions
func MapGet(url string, cache *pokecache.Cache) (MapAreaList, error) {
	// Check if url is "null"
	if url == "null" || url == "" {
		fmt.Println("You're on the first page")
		fmt.Println()
		return MapAreaList{}, nil
	}

	// Check if url is in cache
	if i, ok := cache.Get(url); ok {
		data, err := mapDecode(i)
		if err != nil {
			return MapAreaList{}, err
		}
		mapPrint(&data)

		return data, nil
	}

	// Http GET
	body, err := httpGet(url)
	if err != nil {
		return MapAreaList{}, err
	}

	// Save to cache
	cache.Add(url, body)

	// Decode & print data
	data, err := mapDecode(body)
	if err != nil {
		return MapAreaList{}, err
	}
	mapPrint(&data)

	return data, nil
}

func mapDecode(bytes []byte) (MapAreaList, error) {
	var data MapAreaList
	err := json.Unmarshal(bytes, &data)
	if err != nil {
		return MapAreaList{}, err
	}

	return data, nil
}

func mapPrint(data *MapAreaList) error {
	for _, area := range data.Results{
		fmt.Printf("%s\n", area.Name)
	}

	return nil
}
