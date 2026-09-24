package pokeapi

import (
	"encoding/json"
	"fmt"

	"github.com/gminton07/pokedex/internal/pokecache"
)

type MapArea struct {
	ID                   int                    `json:"id"`
	Name                 string                 `json:"name"`
	GameIndex            int                    `json:"game_index"`
	EncounterMethodRates []EncounterMethodRates `json:"encounter_method_rates"`
	Location             Location               `json:"location"`
	Names                []Names                `json:"names"`
	PokemonEncounters    []PokemonEncounters    `json:"pokemon_encounters"`
}

type EncounterMethodRates struct {
	EncounterMethod Location         `json:"encounter_method"`
	VersionDetails  []struct {
		Rate 	int 	 `json:"rate"`
		Version Location `json:"version"`
	}
}

type Names struct {
	Name     string   `json:"name"`
	Language Location `json:"language"`
}

type EncounterDetails struct {
	MinLevel        int      `json:"min_level"`
	MaxLevel        int      `json:"max_level"`
	Chance          int      `json:"chance"`
	Method          Location `json:"method"`
	ConditionValues []any  `json:"condition_values"`
	PokemonDetails  any    `json:"pokemon_details"`
}

type VersionDetails struct {
	Version          Location           `json:"version"`
	MaxChance        int                `json:"max_chance"`
	EncounterDetails []EncounterDetails `json:"encounter_details"`
}

type PokemonEncounters struct {
	Pokemon        Location         `json:"pokemon"`
	VersionDetails []VersionDetails `json:"version_details"`
}

// Functions
func GetAreaPokemon(url string, cache *pokecache.Cache) (error) {
	// Check if url is cached
	if i, ok := cache.Get(url); ok {
		data, err := decodeMapArea(i)
		if err != nil {
			return fmt.Errorf("error pokeapi.decodeMapArea: %w", err)
		}
		printAreaPokemon(&data)
		return nil
	}

	// http GET
	body, err := httpGet(url)
	if err != nil {
		return fmt.Errorf("error pokeapi.httpGet: %w", err)
	}
	
	// save cache
	cache.Add(url, body)

	// Decode data
	data, err := decodeMapArea(body)
	if err != nil {
		return fmt.Errorf("error pokeapi.decodeMapArea: %w", err)
	}

	// output
	printAreaPokemon(&data)

	return nil
}

func decodeMapArea(b []byte) (MapArea, error) {
	var data MapArea
	err := json.Unmarshal(b, &data)
	if err != nil {
		return MapArea{}, err
	}
	return data, nil
}

func printAreaPokemon(d *MapArea) error {
	// Loop through pokemon names
	for _, v := range d.PokemonEncounters {
		pkmnName := v.Pokemon.Name
		fmt.Printf("%s\n", pkmnName)
	}

	return nil

}
