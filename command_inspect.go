package main

import (
	"errors"
	"fmt"

	"github.com/nunseik/pokedexcli/internal/pokecache"
)

func commandInspect(cache *pokecache.Cache, cfg *config, pokemon string) error {
	if pokemonStruct, ok := cfg.pokedex[pokemon]; ok {
		fmt.Printf("Name: %s\n", pokemonStruct.Name)
		fmt.Printf("Height: %v\n", pokemonStruct.Height)
		fmt.Printf("Weight: %v\n", pokemonStruct.Weight)
		return nil
	} 
	return errors.New("pokemon not in pokedex")
}