package main

import "fmt"

func commandHelp(_ []string, cfg *Config) error {
	command_list := getCommands()
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	if len(command_list) <= 0 {
		return fmt.Errorf("empty command list")
	}
	for command := range command_list {
		fmt.Printf("%s: %s\n", command_list[command].name, command_list[command].description)
	}
	return nil
}
