package pokeapi

import (
	"encoding/json"
	"fmt"
//	"io"
	"net/http"
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

func MapGet(url string) (MapAreaList, error) {
	// Check if url is "null"
	if url == "null" || url == ""{
		fmt.Println("You're on the first page\n")
		return MapAreaList{}, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return MapAreaList{}, err
	}
	defer res.Body.Close()
	
	//var map Map
	var data MapAreaList
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&data)
	if err != nil {
		return MapAreaList{}, err
	}

	mapPrint(&data)

	return data, nil
}

func mapPrint(data *MapAreaList) error {
	//fmt.Printf("count: %d\n", data.Count)
	//fmt.Printf("next: %s\n", data.Next)
	//fmt.Printf("previous: %s\n", data.Previous)
	for _, area := range data.Results{
		fmt.Printf("%s\n", area.Name)
	}
	fmt.Println()

	return nil
}
