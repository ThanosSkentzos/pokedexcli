package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/ThanosSkentzos/pokedexcli/internal/pokecache"
)

type MapJSON struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getMAP(url string, cache *pokecache.Cache) (MapJSON, error) {

	// fmt.Printf("Requesting map data...")
	mapJSON := MapJSON{}
	data, exists := cache.Get(url)
	if exists {
		json.Unmarshal(data, &mapJSON)
	} else {

		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
			return mapJSON, err
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code %d and \nbody: %s\n", res.StatusCode, body)
			return mapJSON, err
		}
		if err != nil {
			log.Fatal(err)
			return mapJSON, err
		}
		json.Unmarshal(body, &mapJSON)
		cache.Add(url, body)
	}
	// fmt.Println("Done.")
	return mapJSON, nil
}
