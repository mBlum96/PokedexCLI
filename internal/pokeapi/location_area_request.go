package pokeapi

import (
	"fmt"
	"strings"
)


func PrintMap(response *LocationResponse){
    for _,area := range(response.Locations){
        fmt.Printf("%s\n", area.Name)
    }
}

func (c *Client) MoveNextPageMap(resp LocationResponse){
	resp.Next = strings.Replace(resp.Next, c.baseUrl,"",-1)
	c.currentMapPage = resp.Next
}

func (c *Client) MovePreviousPageMap(resp LocationResponse){
	resp.Previous = strings.Replace(resp.Previous, c.baseUrl,"",-1)
	c.currentMapPage = resp.Previous
}
