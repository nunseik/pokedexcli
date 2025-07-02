package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nunseik/pokedexcli/internal/pokeapi"
	"github.com/nunseik/pokedexcli/internal/pokecache"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex          map[string]pokeapi.Pokemon
}

func startRepl(cache *pokecache.Cache, cfg *config) {

	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())

		commandName := words[0]
		arg := ""
		if len(words) > 1 {
			arg = words[1]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cache, cfg, arg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*pokecache.Cache, *config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Show 20 locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Show the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "See a list of all the Pokémon located in an area",
			callback:    commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Catch a pokemon, use with a 'catch pokemon name'",
			callback: commandCatch,
		},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
