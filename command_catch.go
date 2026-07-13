package main

import (
	"fmt"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Not enough argument passed to catch command")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	pokemon, err := cfg.pokeAPIClient.GetPokemon(&args[0])
	if err != nil {
		return err
	}

	// TODO: add catch logic using base experience, assume max base experience is 635 for chansey
	fmt.Printf("Base Experience of %s is %d\n", args[0], pokemon.BaseExperience)

	return nil
}
