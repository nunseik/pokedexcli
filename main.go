package main

import (
	"time"
	"github.com/nunseik/pokedexcli/internal/pokeapi"
	"github.com/nunseik/pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cache := pokecache.NewCache(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex: make(map[string]pokeapi.Pokemon),


	}

	startRepl(cache, cfg)
}