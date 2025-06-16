package pokeapi

import (
	"encoding/json"
	"net/http"
	"github.com/nunseik/pokedexcli/internal/pokecache"
	"io"
)

// ListLocations -
func (c *Client) ListLocations(cache *pokecache.Cache, pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if data, ok := cache.Get(url); ok {
		var locations RespShallowLocations
		err := json.Unmarshal(data, &locations)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locations, nil
	}
	
	// Not cached, fetch from API:
	res, err := http.Get(url)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer res.Body.Close()
	

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}
	cache.Add(url, bodyBytes)

	var locations RespShallowLocations
	err = json.Unmarshal(bodyBytes, &locations)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locations, nil
}