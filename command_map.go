package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"errors"
)

type Results struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

type Config struct {
	Next string
	Previous string
}

func commandMap() (error) {
	url := ""
	if config.Next != ""{
		url = config.Next
	} else {
		url = "https://pokeapi.co/api/v2/location-area/"
	}
	
	locations, err := getData(url)
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	config = Config{
		Next: locations.Next,
		Previous: locations.Previous,
	}

	return nil
}

func commandMapb () (error) {
	if config.Previous == "" {
		return errors.New("you're on the first page")
	}
	locations, err := getData(config.Previous)
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	config = Config{
		Next: locations.Next,
		Previous: locations.Previous,
	}

	return nil
}

func getData(url string) (Results, error){
	res, err := http.Get(url)
	if err != nil {
		return Results{}, err
	}
	defer res.Body.Close()

	var locations Results

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locations)
	if err != nil {
		return Results{},err
	}

	return locations, nil
}