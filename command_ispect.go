package main

import (
	"errors"
	"pokedexcli/internal/pokeapi"
)

func commandInspect(c *pokeapi.Client, pokemonName string) error{
	pokedex := c.GetPokeDex()
	pokemon,ok := pokedex[pokemonName]
	if !ok{
		return errors.New("You haven't caught this pokemon yet")
	}
	pokeapi.PrintPokemonInfo(&pokemon)
	return nil
}