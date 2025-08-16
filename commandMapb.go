package main

import "fmt"

func commandMapb(cfg *Config) error {
	if cfg.prevLocationsURL == nil {
		return fmt.Errorf("no previous locations to print, use 'map'")
	} else {
		locations, err := cfg.pokeapiClient.ListLocations(*cfg.prevLocationsURL)
		if err != nil {
			return fmt.Errorf("error listing locations")
		}
		for i := range locations.Results {
			fmt.Printf("%s\n", locations.Results[i].Name)
		}
	}
	return nil
}
