package main

import "fmt"

func commandMapf(cfg *config, _ []string) error {
	locationResp, err := cfg.pokeAPIClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}
	cfg.prevLocationsURL = locationResp.Previous
	cfg.nextLocationsURL = locationResp.Next

	for _, location := range locationResp.Locations {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config, _ []string) error {

	if cfg.prevLocationsURL == nil {
		return fmt.Errorf("you're on the first page\n")
	}

	locationResp, err := cfg.pokeAPIClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.prevLocationsURL = locationResp.Previous
	cfg.nextLocationsURL = locationResp.Next

	for _, location := range locationResp.Locations {
		fmt.Println(location.Name)
	}

	return nil
}
