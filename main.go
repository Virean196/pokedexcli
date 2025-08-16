package main

import (
	"github.com/virean196/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := &Config{
		pokeapiClient: *pokeapi.NewClient(),
	}
	startRepl(cfg)
}
