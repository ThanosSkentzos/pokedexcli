package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func getMAP() (MapJSON, error) {

	fmt.Printf("Requesting map data...")
	mapJSON := MapJSON{}
	URL := "https://pokeapi.co/api/v2/location/"
	res, err := http.Get(URL)
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
	fmt.Println("Done.")
	return mapJSON, nil
}
