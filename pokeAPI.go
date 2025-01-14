package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
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

func getMAP(url string) (MapJSON, error) {

	// fmt.Printf("Requesting map data...")
	mapJSON := MapJSON{}
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
	// fmt.Println("Done.")
	return mapJSON, nil
}
