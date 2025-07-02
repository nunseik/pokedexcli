package main

import (
	"fmt"
	"os"

	"github.com/nunseik/pokedexcli/internal/pokecache"
)

func commandExit(cache *pokecache.Cache,cfg *config,arg string) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}