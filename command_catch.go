package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/nunseik/pokedexcli/internal/pokecache"
)

func commandCatch(cache *pokecache.Cache, cfg *config, pokemon string) error {
	if pokemon == "" {
		return errors.New("no location provided")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	pokemonsResp, err := cfg.pokeapiClient.GetPokemon(cache, pokemon)

	if err != nil {
		return err
	}

	randomInt := rand.Intn(500)

	if randomInt >= pokemonsResp.BaseExperience {
		fmt.Printf("%s was caught\n", pokemon)
		cfg.pokedex[pokemonsResp.Name] = pokemonsResp
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}
	return nil
}
