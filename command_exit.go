package main

import (
	"fmt"
	"os"
)

func commandExit([]string) error{
    fmt.Println("Exiting the pokedex...")
    os.Exit(0)
    return nil
}