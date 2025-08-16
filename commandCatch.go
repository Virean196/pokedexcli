package main

import (
	"fmt"
	"math/rand"

	"github.com/virean196/pokedexcli/internal/pokeapi"
)

func commandCatch(args []string, cfg *Config) error {
	baseURL := pokeapi.GetBaseUrl()
	urlToFetch := baseURL + "pokemon/" + args[1]
	if len(args) < 2 {
		return fmt.Errorf("no pokemon selected")
	}
	pokemonInfo, err := cfg.pokeapiClient.GetPokemonInfo(urlToFetch)
	if err != nil {
		return fmt.Errorf("invalid pokemon name")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonInfo.Name)
	if rand.Intn(pokemonInfo.BaseExperience) < pokemonInfo.BaseExperience/2 {
		fmt.Printf("%s was caught!\n", pokemonInfo.Name)
		cfg.pokedex[pokemonInfo.Name] = pokemonInfo
	} else {
		fmt.Printf("%s escaped\n", pokemonInfo.Name)
	}
	return nil
}
