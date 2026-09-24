package pokeapi

import (
	"encoding/json"
	"math/rand"
	"time"
)

// name, height, weight, stats {hp, attack, defense, etc.}, types
type Pokemon struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	BaseExperience int     `json:"base_experience"`
	Height         int     `json:"height"`
	Weight         int     `json:"weight"`
	Stats          []Stats `json:"stats"`
	Types          []Types `json:"types"`
}

type Stats struct {
	BaseStat int `json:"base_stat"`
	Stat Location `json:"stat"`
}

type Types struct {
	Slot int `json:"slot"`
	Type Location `json:"type"`
}

func GetPokemon(url string) (Pokemon, error) {
	// HTTP GET
	body, err := httpGet(url)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon, err := decodePokemon(body)
	if err != nil {
		return Pokemon{}, err
	}

	if ok := catchPokemon(pokemon.BaseExperience); !ok {
		return Pokemon{}, nil
	}

	return pokemon, nil
}

func decodePokemon(b []byte) (Pokemon, error) {
	var data Pokemon
	err := json.Unmarshal(b, &data)
	if err != nil {
		return Pokemon{}, err
	}

	return data, nil
}

func catchPokemon(exp int) (bool) {
	// Random int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	chance := r.Intn(exp)

	if chance < int(exp / 3) {
		return true
	}

	return false
}
