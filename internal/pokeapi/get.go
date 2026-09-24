package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gminton07/pokedex/internal/pokecache"
)

type MapAreaList struct {
	Count        int `json:"count"`
	Next         string `json:"next"`
	Previous     string `json:"previous,omitempty"`
	Results      []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
}

func MapGet(url string, cache *pokecache.Cache) (MapAreaList, error) {
	// Check if url is "null"
	if url == "null" || url == "" {
		fmt.Println("You're on the first page\n")
		return MapAreaList{}, nil
	}

	// check if url in cache
	if i, ok := cache.Get(url); ok {
		data, err := mapDecode(i)
		if err != nil {
			return MapAreaList{}, err
		}
		////fmt.Println("*** CACHED ***")
		mapPrint(&data)
		
		return data, nil
	}

	// do http GET
	res, err := http.Get(url)
	if err != nil {
		return MapAreaList{}, err
	}
	defer res.Body.Close()
	
	// convert response Body to []bytes
	body, err := io.ReadAll(res.Body)

	// save to cache
	////fmt.Printf("Added %s to cache\n", url)
	cache.Add(url, body)
	
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
	fmt.Println()

	return nil
}
