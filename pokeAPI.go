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

type AreaJSON struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func getMap(url string, cache *pokecache.Cache) (MapJSON, error) {

	// fmt.Printf("Requesting map data...")
	mapJSON := MapJSON{}
	data, exists := cache.Get(url)
	if exists {
		json.Unmarshal(data, &mapJSON)
		return mapJSON, nil
	}

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
	// fmt.Println("Done.")
	return mapJSON, nil
}

func getArea(url string, cache *pokecache.Cache) (AreaJSON, error) {
	areaJSON := AreaJSON{}
	data, exists := cache.Get(url)
	if exists {
		json.Unmarshal(data, &areaJSON)
		return areaJSON,nil
	}
	res,err := http.Get(url)
	if err!=nil{
		log.Fatal(err)
		return areaJSON,err
	}
	body,err:= io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code %d and \nbody: %s\n", res.StatusCode, body)
		return areaJSON, err
	}
	if err != nil {
		log.Fatal(err)
		return areaJSON, err
	}
	json.Unmarshal(body, &areaJSON)
	cache.Add(url, body)
	// fmt.Println("Done.")
	return areaJSON, nil
}
