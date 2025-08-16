package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/virean196/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2/"

// TODO - Move cache into proper cache package, rn im too dumb for that
type Client struct {
	client http.Client
	cache  *pokecache.Cache
}

func NewClient(timeout, interval time.Duration) *Client {
	return &Client{
		client: http.Client{Timeout: timeout},
		cache:  pokecache.NewCache(interval),
	}
}

type Location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationData struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) ListLocations(req_url string) (Location, error) {
	cachedData, exists := c.cache.Get(req_url)
	if exists {
		var locations Location
		if err := json.Unmarshal(cachedData, &locations); err != nil {
			return Location{}, fmt.Errorf("error unmarshelling locations")
		}
		return locations, nil
	}

	res, err := c.client.Get(req_url)
	if err != nil {
		return Location{}, fmt.Errorf("error listing locations")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, fmt.Errorf("error reading body")
	}

	c.cache.Add(req_url, body)

	var locations Location
	if err := json.Unmarshal(body, &locations); err != nil {
		return Location{}, fmt.Errorf("error unmarshelling locations")
	}
	return locations, nil
}

func (c *Client) getLocationData(req_url string) (LocationData, error) {
	cachedData, exists := c.cache.Get(req_url)
	if exists {
		var locationData LocationData
		if err := json.Unmarshal(cachedData, &locationData); err != nil {
			return LocationData{}, fmt.Errorf("error unmarshelling locations")
		}
		return locationData, nil
	}
	res, err := c.client.Get(req_url)
	if err != nil {
		return LocationData{}, fmt.Errorf("error listing locations")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationData{}, fmt.Errorf("error reading body")
	}

	c.cache.Add(req_url, body)

	var locationData LocationData
	if err := json.Unmarshal(body, &locationData); err != nil {
		return LocationData{}, fmt.Errorf("error unmarshelling locations")
	}
	return locationData, nil
}

func (c *Client) PokemonFromLocation(req_url string) ([]string, error) {
	locationData, err := c.getLocationData(req_url)
	var locationPokemonNames []string
	for i := range locationData.PokemonEncounters {
		locationPokemonNames = append(locationPokemonNames, locationData.PokemonEncounters[i].Pokemon.Name)
	}
	return locationPokemonNames, err
}

func GetBaseUrl() string {
	return baseURL
}
