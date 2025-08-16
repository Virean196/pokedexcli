package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2/"

type Client struct {
	client http.Client
}

func NewClient() *Client {
	return &Client{
		client: http.Client{Timeout: 5 * time.Second},
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
	res, err := c.client.Get(req_url)
	if err != nil {
		return Location{}, fmt.Errorf("error listing locations")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, fmt.Errorf("error reading body")
	}
	var locations Location
	if err := json.Unmarshal(body, &locations); err != nil {
		return Location{}, fmt.Errorf("error unmarshelling locations")
	}
	return locations, nil
}

func GetBaseUrl() string {
	return baseURL
}
