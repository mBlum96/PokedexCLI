package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"strings"
	"sync"
)

func handleCommand(c *pokeapi.Client ,input string) error{
    words := strings.Fields(input)
    if(len(words) == 0){
        return errors.New("No command entered")
    }
    command := words[0]
    params := words[1:]
    cmd,ok := commands[command]
    if(!ok){
        var err_str string = fmt.Sprintf("command '%s' is unknown", command)
        return errors.New(err_str)
    }
    err := cmd.callback(c,params)
	if err != nil{
		return err
	}
    return nil
}


func repl(c *pokeapi.Client) {
	reader := bufio.NewReader(os.Stdin)
    for{
        fmt.Print("pokedex>")
        input, err:= reader.ReadString('\n')
        if err != nil{
            fmt.Println("Error reading input:", err)
            continue
        }
        input = strings.TrimSpace(input)
		commandLock := &sync.Mutex{}
        commandLock.Lock()
        err = handleCommand(c,input)
        commandLock.Unlock()
        if(err != nil){
            fmt.Println("Error executing command:", err)
            continue
        }
    }
}