package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListPokemonsInLocation(location *string) (RespLocationPokemons, error) {
	url := baseURL + "/location-area/" + *location

	if data, ok := c.cache.Get(url); ok {
		var pokemons RespLocationPokemons
		if err := json.Unmarshal(data, &pokemons); err != nil {
			return RespLocationPokemons{}, fmt.Errorf("Error unmarshalling the cache: %v\n", err)
		}
		return pokemons, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationPokemons{}, err
	}

	resp, err := c.http_client.Do(req)
	if err != nil {
		return RespLocationPokemons{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return RespLocationPokemons{}, fmt.Errorf("Request to pokeAPI failed with satus %v\n", resp.StatusCode)
	}

	defer resp.Body.Close()

	var pokemons RespLocationPokemons
	raw_data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationPokemons{}, fmt.Errorf("Error converting response to bytes: %v\n", err)
	}

	if err := json.Unmarshal(raw_data, &pokemons); err != nil {
		return RespLocationPokemons{}, fmt.Errorf("Error while unmarshalling PokeAPI location pokemon response: %v\n", err)
	}

	c.cache.Add(url, raw_data)

	return pokemons, nil
}
