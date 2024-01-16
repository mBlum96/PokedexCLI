package main

import "pokedexcli/internal/pokeapi"

func commandExplore(c *pokeapi.Client, targetLocation string) error{
	exploreAreaAdress := targetLocation + "-area"
	locationAdress := pokeapi.DefaultBaseUrl + "location-area/"
	pokemon, err := c.FetchPokemonEncountered(locationAdress + exploreAreaAdress)
	
	return nil
}
