package pokeapi

import (
	"encoding/json"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	res, err := http.Get(url)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer res.Body.Close()

	var locations RespShallowLocations

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locations)
	if err != nil {
		return RespShallowLocations{},err
	}

	return locations, nil
}