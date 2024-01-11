package pokeapi

import (
	"fmt"
)


func PrintMap(response *LocationResponse){
    for _,area := range(response.Locations){
        fmt.Printf("%s\n", area.Name)
    }
}

func (c *Client) MoveNextPageMap(resp LocationResponse){
	c.currentMapPage = resp.Next
}

func (c *Client) MovePreviousPageMap(resp LocationResponse){
	c.currentMapPage = resp.Previous
}
