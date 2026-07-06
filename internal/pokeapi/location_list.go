package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	if data, ok := c.cache.Get(url); ok {
		var locations RespShallowLocations
		if err := json.Unmarshal(data, &locations); err != nil {
			return RespShallowLocations{}, fmt.Errorf("Error unmarshalling the cache: %v\n", err)
		}
		return locations, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.http_client.Do(req)
	if err != nil {
		return RespShallowLocations{}, fmt.Errorf("Error while making request to PokeAPI location-area: %v\n", err)
	}

	defer resp.Body.Close()
	var locations RespShallowLocations
	raw_data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, fmt.Errorf("Error converting response to bytes: %v\n", err)
	}

	if err := json.Unmarshal(raw_data, &locations); err != nil {
		return RespShallowLocations{}, fmt.Errorf("Error while unmarshalling PokeAPI location-area response: %v\n", err)
	}

	c.cache.Add(url, raw_data)

	return locations, nil
}
