package main

import (
    "fmt"
    "os"
)

func main() {
    f, _ := os.OpenFile("msg.txt", os.O_RDWR|os.O_CREATE, 0644)
    defer f.Close()
    fmt.Fprintf(f, "%d\n", 100)
}

