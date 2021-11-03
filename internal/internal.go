package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
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
    jsonPayload,err := json.Marshal(payload)
    if err != nil {
        fmt.Printf("Error %s",err)
        return err 
    }
    fmt.Println(string(jsonPayload))
    jsonBuffer := bytes.NewBuffer(jsonPayload)
    resp , err := c.Post(route,"application/json",jsonBuffer)
    if err != nil {
        fmt.Printf("Error %s ",err)
        return err 
    }
    defer resp.Body.Close()
    dec := json.NewDecoder(resp.Body);
    var output interface{}
    err = dec.Decode(&output)
    payload = output
    if err != nil {
        fmt.Printf("Error %s ", err)
    }
    out,err := json.Marshal(output)
    if err != nil {
        fmt.Printf("Error %s",err)
        return err 
    }
    fmt.Println(string(out))
    return nil
}

func GetResult(route string, output interface{}) error {
    c := http.Client{
        Timeout:time.Duration(1) * time.Second,
    }
    resp, err := c.Get(route)
    if resp.StatusCode == 404 {
        return errors.New("Not Found")
    }
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

func PostTask(route string, payload interface{}) (error,string){
    c := http.Client{
        Timeout:time.Duration(1) * time.Second,
    }
    jsonPayload,err := json.Marshal(payload)
    if err != nil {
        log.Println("A")
        fmt.Printf("Error %s",err)
        return err,""
    }
    jsonBuffer := bytes.NewBuffer(jsonPayload)
    resp , err := c.Post(route,"application/json",jsonBuffer)
    if resp.StatusCode == 404 {
        return errors.New("Not Found"),""
    }
    if err != nil {
        log.Println("B")
        fmt.Printf("Error %s ",err)
        return err,"" 
    }
    defer resp.Body.Close()
    dec := json.NewDecoder(resp.Body);
    var output struct{
        TaskId string `json:"id"`
    }
    err = dec.Decode(&output)
    payload = output
    if err != nil {
        log.Println("C")
        fmt.Printf("Error %s ", err)
    }
    return nil,output.TaskId
}
