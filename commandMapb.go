package main

import "fmt"

func commandMapb(cfg *Config) error {
	if cfg.prevLocationsURL == nil {
		return fmt.Errorf("you're on the first page")
	} else {
		locations, err := cfg.pokeapiClient.ListLocations(*cfg.prevLocationsURL)
		if err != nil {
			return fmt.Errorf("error listing locations")
		}

		cfg.nextLocationsURL = &locations.Next
		cfg.prevLocationsURL = &locations.Previous

		for i := range locations.Results {
			fmt.Printf("%s\n", locations.Results[i].Name)
		}
	}
	return nil
}
