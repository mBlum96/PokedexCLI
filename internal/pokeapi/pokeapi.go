package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"pokedexcli/internal/pokecache"
	"time"
)

const(
	DefaultLocationPage = "location/"
	DefaultInterval = 5 * time.Minute
	DefaultBaseUrl = "https://pokeapi.co/api/v2/"
	NEXT = "next"
	PREV = "previous"
)

type Client struct {
	baseUrl string
	cache   *pokecache.Cache
	addresses LocationAddresses
}

type LocationFetcher interface{
	FetchLocation(direction string)
}

type PokemonEncounterFetcher interface{
	FetchPokemonEncountered(location string)
}

func (c *Client) GetPreviousAddress() string{
	return c.addresses.Previous
}


func NewClient() *Client {
	return &Client{
		baseUrl: DefaultBaseUrl,
		cache: pokecache.NewCache(DefaultInterval),
		addresses: LocationAddresses{
			Current: DefaultBaseUrl + DefaultLocationPage,
			Next: "",
			Previous: "",
		},
	}
}

func (c *Client) FetchLocation(direction string) (*LocationResponse, error) {
	var result LocationResponse
	var requestAddress string
	if(direction == NEXT){
		if(c.addresses.Next == ""){
			requestAddress = c.addresses.Current
		}else{
			requestAddress = c.addresses.Next
		}
	}else{
		if(c.addresses.Previous == ""){
			return nil, errors.New("No previous location")
		}
		requestAddress = c.addresses.Previous
	}
	cacheResponse, exists := c.cache.Get(requestAddress)
	var locationFetchResponse interface{}
	var fetchErr error
	if exists {
		locationFetchResponse, fetchErr = fetchFromCache(cacheResponse, &result, c)
	}else{
		locationFetchResponse, fetchErr = fetchFromServer(requestAddress, &result,c)
	}
	if fetchErr != nil{
		return nil, fetchErr
	}
	assertedLocationFetchResponse, ok := locationFetchResponse.(*LocationResponse)
	if(ok == false){
		return nil, errors.New("Error casting result")
	}
	c.addresses.Next = assertedLocationFetchResponse.Next
	c.addresses.Previous = assertedLocationFetchResponse.Previous
	return assertedLocationFetchResponse, fetchErr
}

func fetchFromCache(cacheResponse []byte, result *LocationResponse,c *Client) (interface{}, error){
	err := json.Unmarshal(cacheResponse, &result)
	if err != nil {
		return nil, errors.New("Error unmarshalling response body")
	}
	return result, nil
}

func fetchFromServer(requestAddress string, result interface{},c *Client)(interface{}, error){
	resp, err := http.Get(requestAddress)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Error reading response body")
	}
	err = json.Unmarshal(body, &result)
	c.cache.Add(requestAddress, body)
	if err != nil {
		return nil, errors.New("Error unmarshalling response body")
	}
	return result, nil
}

func FetchPokemonEncountered(location string) (*PokemonEncounterResponse, error){
	return nil, nil
}