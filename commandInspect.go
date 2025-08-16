package main

import "fmt"

func commandInspect(args []string, cfg *Config) error {
	if len(args) < 2 {
		return fmt.Errorf("too few arguments")
	}
	pokemon, exists := cfg.pokedex[args[1]]
	if exists {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Name: %v\n", pokemon.Weight)
		fmt.Printf("Stats: \n")
		for i := range pokemon.Stats {
			fmt.Printf("  - %s: %v\n", pokemon.Stats[i].Stat.Name, pokemon.Stats[i].BaseStat)
		}
		fmt.Print("Types: \n")
		for i := range pokemon.Types {
			fmt.Printf("  - %s\n", pokemon.Types[i].Type.Name)
		}
	} else {
		return fmt.Errorf("you haven't caught that pokemon yet")
	}
	return nil
}
