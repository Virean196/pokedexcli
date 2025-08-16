package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *Config) error {
	fmt.Printf("\nClosing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
