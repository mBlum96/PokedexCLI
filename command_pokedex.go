package main

import (
	"errors"
	"pokedexcli/internal/pokeapi"
)

func commandPokedex(c *pokeapi.Client) error{
	pokedex := c.GetPokeDex()
	if(len(pokedex)==0){
		return errors.New("You haven't caught any pokemon yet")
	}
	pokeapi.PrintPokedex(pokedex)
	return nil
}