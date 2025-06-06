package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Location struct {
    Next   string `json:"next"`
    Previous    string `json:"previous"`
    Location string `json:"results"`
}

func commandMap() error {
	fmt.Println("map starting...")
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var locations []Location

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locations); err != nil {
		return err
	}	

	 for _, location := range locations {
        fmt.Print(location.Location)
    }
	return nil
} 