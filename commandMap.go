package main

import (
	"fmt"

	"github.com/virean196/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *Config) error {
	var urlToFetch string

	if cfg.nextLocationsURL == nil {
		urlToFetch = pokeapi.GetBaseUrl() + "location-area/"
	} else {
		urlToFetch = *cfg.nextLocationsURL
	}

	locations, err := cfg.pokeapiClient.ListLocations(urlToFetch)
	if err != nil {
		return fmt.Errorf("error listing locations")
	}

	cfg.nextLocationsURL = &locations.Next
	cfg.prevLocationsURL = &locations.Previous

	for i := range locations.Results {
		fmt.Printf("%s\n", locations.Results[i].Name)
	}
	return nil
}
