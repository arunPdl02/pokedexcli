package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name *string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + *name

	pokemon := RespPokemon{}

	if data, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(data, &pokemon); err != nil {
			return RespPokemon{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	resp, err := c.http_client.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return RespPokemon{}, fmt.Errorf("Request to pokeAPI /pokemon failed with status %v\n", resp.StatusCode)
	}

	defer resp.Body.Close()

	raw_data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	if err := json.Unmarshal(raw_data, &pokemon); err != nil {
		return RespPokemon{}, err
	}

	c.cache.Add(url, raw_data)

	return pokemon, nil
}
