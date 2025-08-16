package main

import (
	"fmt"

	"github.com/virean196/pokedexcli/internal/pokeapi"
)

func commandExplore(location []string, cfg *Config) error {
	if len(location) < 2 {
		return fmt.Errorf("no location selected")
	}
	fmt.Printf("Exporing %s...\nFound pokemon: \n", location[1])
	urlToFetch := pokeapi.GetBaseUrl() + "location-area/" + location[1]
	nameList, err := cfg.pokeapiClient.PokemonFromLocation(urlToFetch)
	for i := range nameList {
		fmt.Printf("- %s\n", nameList[i])
	}
	return err
}
