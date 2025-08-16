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

func GetBaseUrl() string {
	return baseURL
}
