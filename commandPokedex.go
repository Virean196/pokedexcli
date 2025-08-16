package main

import "fmt"

func commandPokedex(_ []string, cfg *Config) error {
	fmt.Printf("Your Pokedex: \n")
	if len(cfg.pokedex) == 0 {
		return fmt.Errorf("you haven't caught any pokemon yet")
	}
	for i := range cfg.pokedex {
		fmt.Printf("  - %s\n", cfg.pokedex[i].Name)
	}
	return nil
}
