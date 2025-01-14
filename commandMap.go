package main

import (
	"fmt"
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

var mapJSON MapJSON

var counter int

func commandMap() error {
	var err error
	if mapJSON.Count == 0 {
		mapJSON, err = getMAP()
		if err != nil {
			fmt.Println("Error while getting map:", err)
			return err
		}
	}
	names := getNames(mapJSON)
	start, end := getLimits(names, counter)
	for i, name := range names[start:end] {
		fmt.Printf("%d. %s\n", start+i+1, name)
	}
	counter += 1
	return nil
}

func getNames(mapJSON MapJSON) []string {
	var names []string
	for _, result := range mapJSON.Results {
		names = append(names, result.Name)
	}
	return names
}

func getLimits(array []string, counter int) (int, int) {
	N := 20
	start := counter * N
	end := (counter + 1) * N
	if end > len(array) {
		end = len(array)
	}
	if start >= len(array) {
		counter = 0
		start = 0
		end = N
	}
	return start, end
}
