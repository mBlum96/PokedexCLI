package pokeapi

import (
	"fmt"
)


func PrintMap(response *LocationResponse){
    for _,area := range(response.Locations){
        fmt.Printf("%s\n", area.Name)
    }
}