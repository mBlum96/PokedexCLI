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
	exploreAreaAdress string
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
		exploreAreaAdress: "",
	}
}

func fetchData (requestAddress string, result interface{}, c *Client) (error){
	cacheResponse, exists := c.cache.Get(requestAddress)
	var fetchErr error
	if exists {
		fetchErr = fetchFromCache(cacheResponse, &result, c)
	}else{
		fetchErr = fetchFromServer(requestAddress, &result,c)
	}
	if fetchErr != nil{
		return fetchErr
	}
	return nil
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
	fetchErr := fetchData(requestAddress, &result, c)
	if(fetchErr!=nil){
		return nil,fetchErr
	}
	c.addresses.Next = result.Next
	c.addresses.Previous = result.Previous
	return &result, nil
}

func fetchFromCache(cacheResponse []byte, result interface{},c *Client) (error){
	err := json.Unmarshal(cacheResponse, &result)
	if err != nil {
		return errors.New("Error unmarshalling response body")
	}
	return nil
}

func fetchFromServer(requestAddress string, result interface{},c *Client)(error){
	resp, err := http.Get(requestAddress)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New("Error reading response body")
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return errors.New("Error unmarshalling response body")
	}
	c.cache.Add(requestAddress, body)
	return nil
}

func (c *Client) FetchPokemonEncountered(location string) (*pokemonEncountered, error){
	var result pokemonEncountered
	if(location == ""){
		return nil, errors.New("No location provided")
	}
	fetchErr := fetchData(location, &result, c)
	if(fetchErr!=nil){
		if(fetchErr.Error() == "Error unmarshalling response body"){
			return nil, errors.New("No pokemon encountered, try a different location or check the spelling")
		}
		return nil,fetchErr
	}
	return &result, nil
}