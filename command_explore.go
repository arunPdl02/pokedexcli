package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if args == nil || len(args) != 1 {
		return fmt.Errorf("Not enough argument passed to explore command")
	}
	area := args[0]
	locationPokemonResp, err := cfg.pokeAPIClient.ListPokemonsInLocation(&area)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", locationPokemonResp.Name)
	fmt.Println("Found Pokemon: ")
	for _, encounter := range locationPokemonResp.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
