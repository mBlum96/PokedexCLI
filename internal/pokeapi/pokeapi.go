package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"pokedexcli/internal/pokecache"
)


type Client struct {
	baseUrl string
	cache   *pokecache.Cache
	currentMapPage string
}

func NewClient() *Client {
	return &Client{
		baseUrl: "https://pokeapi.co/api/v2/",
		cache: pokecache.NewCache(5),
	}
}

func (c *Client) FetchLocation() (*LocationResponse, error) {
	resp, err := http.Get(c.baseUrl + c.currentMapPage)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Error reading response body")
	}
	var result LocationResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.New("Error unmarshalling response body")
	}
	return &result, nil
}
