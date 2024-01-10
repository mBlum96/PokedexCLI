package main

import "fmt"

func commandHelp() error{
    fmt.Println("Available commands:")
    for _,cmd := range(commands){
        //print command name and description
        fmt.Printf("%v: %v\n",cmd.name,cmd.description)
    }
    return nil
}