package main

import (
    "fmt"
    "time"
)

func main() {
    d := time.Duration(5)
    fmt.Printf("Start: %v\n", time.Now())
    c := time.After(d * time.Second)
    t := <-c 
    fmt.Printf("End: %v\n", t)
}

