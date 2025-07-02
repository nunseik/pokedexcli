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

func commandPokedex (cache *pokecache.Cache, cfg *config, arg string) error {
	if len(cfg.pokedex) == 0 {
		return errors.New("no pokemons in your pokedex, go catch them")
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf("- %s\n", pokemon.Name)
	}
	return nil
}