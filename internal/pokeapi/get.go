package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)

func mapGet(url string) (any, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		err := fmt.Errorf("Response failed with status code: %d and\nbody: %s", res.StatusCode, body)
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return "Success", nil
}
