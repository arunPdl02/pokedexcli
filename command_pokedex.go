package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for key, _ := range cfg.caughtPokemon {
		fmt.Println("- " + key)
	}
	return nil
}
