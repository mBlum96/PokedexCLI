package main

import "pokedexcli/internal/pokeapi"

	
type pokemonEncountered struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}
	}
}

func commandExplore(c *pokeapi.Client, targetLocation string) error{
	exploreAreaAdress := targetLocation + "-area"
	locationAdress := pokeapi.DefaultBaseUrl + "location-area/"
	locationArea, err := c.FetchPokemonEncountered(locationAdress + exploreAreaAdress)
	
	return nil
}
