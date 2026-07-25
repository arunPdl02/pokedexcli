package main

import (
	"fmt"
	"math/rand"
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

	NORMALIZER := 1200.00
	base_experience := float64(pokemon.BaseExperience)
	normalized_base_experience := (NORMALIZER - base_experience) / NORMALIZER

	caught_chance := normalized_base_experience * float64(rand.Intn(20))
	// fmt.Printf("Base Experience of %s is %d, and it's normalized version is %f\n", args[0], pokemon.BaseExperience, normalized_base_experience)
	// fmt.Printf("Catch change is %f\n", caught_chance)

	if caught_chance < 10 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemon.Name)
	cfg.caughtPokemon[pokemon.Name] = pokemon
	return nil
}
