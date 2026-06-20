package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.http_client.Do(req)
	if err != nil {
		return RespShallowLocations{}, fmt.Errorf("Error while making request to PokeAPI location-area: %v", err)
	}

	defer resp.Body.Close()
	var locations RespShallowLocations
	decoder := json.NewDecoder(resp.Body)

	if err := decoder.Decode(&locations); err != nil {
		return RespShallowLocations{}, fmt.Errorf("Error while unmarshalling PokeAPI location-area response: %v", err)
	}

	return locations, nil
}
