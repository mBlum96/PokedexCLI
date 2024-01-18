package main

import "pokedexcli/internal/pokeapi"

func commandExplore(c *pokeapi.Client, targetLocation string) error{
	exploreAreaAdress := targetLocation + "-area"
	locationEndpoint := pokeapi.DefaultBaseUrl + "location-area/"
	pokemon, err := c.FetchPokemonEncountered(locationEndpoint + exploreAreaAdress)
	if(err!=nil){
		return err
	}
	pokeapi.PrintPokemon(pokemon)
	return nil
}
