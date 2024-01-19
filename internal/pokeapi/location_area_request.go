package pokeapi

import (
	"fmt"
)


func PrintMap(response *LocationResponse){
    for _,area := range(response.Locations){
        fmt.Printf("%s\n", area.Name)
    }
}

func PrintPokemon(response *pokemonEncountered){
    println("Found Pokemon:")
    for _,pokemon := range(response.PokemonEncounters){
        fmt.Printf("%s\n", pokemon.Pokemon.Name)
    }
}

func PrintPokemonInfo(response *PokemonInfo){
    fmt.Printf("Name: %s\n", response.Name)
    fmt.Printf("Height: %d\n", response.Height)
    fmt.Printf("Weight: %d\n", response.Weight)
    fmt.Println("Stats:")
    for _,s := range(response.Stats){
        fmt.Printf("\t%s: %d\n", s.Stat.Name, s.BaseStat)
    }
    fmt.Println("Types:")
    for _,t := range(response.Types){
        fmt.Printf("\t%s\n", t.Type.Name)
    }
}