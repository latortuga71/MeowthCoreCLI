package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// create functions for post requests
// create functions for get requests
// to make it a simple function call to get the data back before converting to struct


func StringIn (a string,list []string) bool {
    for _,b := range list {
        if b == a {
            return true
        }
    }
    return false
}


func Get(route string, output interface{}) error {
    c := http.Client{
        Timeout:time.Duration(1) * time.Second,
    }
    resp, err := c.Get(route)
    if err != nil {
        fmt.Printf("Error %s ",err)
        return err
    }
    defer resp.Body.Close()
    dec := json.NewDecoder(resp.Body);
    err = dec.Decode(&output)
    if err != nil {
        fmt.Printf("Error %s ",err)
        return err
    }
    return nil
}


func Post(route string,payload interface{}) error {
    c := http.Client{
        Timeout:time.Duration(1) * time.Second,
    }
    json,err := json.Marshal(payload)
    if err != nil {
        fmt.Printf("Error %s",err)
        return err 
    }
    fmt.Println(string(json))
    jsonBuffer := bytes.NewBuffer(json)
    resp , err := c.Post(route,"application/json",jsonBuffer)
    if err != nil {
        fmt.Printf("Error %s ",err)
        return err 
    }
    defer resp.Body.Close()
    if err != nil {
        fmt.Printf("Error %s ", err)
    }
    return nil

}
