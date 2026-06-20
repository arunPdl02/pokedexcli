package main

import (
	"time"

	"github.com/arunPdl02/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := &config{
		pokeAPIClient: pokeapi.NewClient(5 * time.Second),
	}
	startRepl(cfg)
}
