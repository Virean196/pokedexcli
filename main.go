package main

import (
	"time"

	"github.com/virean196/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := &Config{
		pokeapiClient: *pokeapi.NewClient(5*time.Second, 5*time.Minute),
	}

	startRepl(cfg)
}
