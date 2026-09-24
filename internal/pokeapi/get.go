package pokeapi

import (
	"io"
	"net/http"
)

func httpGet(url string) ([]byte, error) {
	// do http GET
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	
	// convert response Body to []bytes
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

