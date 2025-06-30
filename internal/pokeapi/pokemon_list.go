package pokeapi

import (
	"encoding/json"
	"net/http"
	"github.com/nunseik/pokedexcli/internal/pokecache"
	"io"
	"errors"
)

// ListLocations -
func (c *Client) GetLocationPokemon(cache *pokecache.Cache, locationArg string) (RespShallowPokemons, error) {
	
	if locationArg == "" {
		return RespShallowPokemons{}, errors.New("location not provided")
	}

	url := baseURL + "/location-area/" + locationArg

	if data, ok := cache.Get(url); ok {
		var pokemons RespShallowPokemons
		err := json.Unmarshal(data, &pokemons)
		if err != nil {
			return RespShallowPokemons{}, err
		}
		return pokemons, nil
	}
	
	// Not cached, fetch from API:
	res, err := http.Get(url)
	if err != nil {
		return RespShallowPokemons{}, err
	}
	defer res.Body.Close()
	

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowPokemons{}, err
	}
	cache.Add(url, bodyBytes)

	var pokemons RespShallowPokemons
	err = json.Unmarshal(bodyBytes, &pokemons)
	if err != nil {
		return RespShallowPokemons{}, err
	}

	return pokemons, nil
}