package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func commandBMap() error{
    currentMapPage = result.Previous
    resp, err := http.Get(currentMapPage)
    if err != nil{
        return err
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return errors.New("Error reading response body")
    }
    err = json.Unmarshal(body, &result)
    mapPrinter(&result)
    return nil
}