package main

import (
	"fmt"
	"errors"
	"github.com/nunseik/pokedexcli/internal/pokecache"
)

func commandExplore(cache *pokecache.Cache, cfg *config, locationArg string) error {
	if locationArg == ""{
		return errors.New("no location provided")
	}
	pokemonsResp, err := cfg.pokeapiClient.GetLocationPokemon(cache, locationArg)

	if err != nil {
		return err
	}
	fmt.Println("Exploring ", locationArg, "...")
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonsResp.PokemonEncounters {
		fmt.Println(" - ", pokemon.Pokemon.Name)
	}
	return nil
}
