package main

import (
	"errors"
	"fmt"
	"github.com/nunseik/pokedexcli/internal/pokecache"
)

func commandMapf(cache *pokecache.Cache, cfg *config, locationArg string) (error) {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cache, cfg.nextLocationsURL)

	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb (cache *pokecache.Cache, cfg *config, locationArg string) (error) {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}
	locationsResp, err := cfg.pokeapiClient.ListLocations(cache, cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

