package main

import (
    "os"
    "io"
)

func main() {
    n, _ := os.OpenFile("new.txt", os.O_RDWR|os.O_CREATE, 0644)
    defer n.Close()
    o, _ := os.Open("old.txt")
    defer o.Close()
    io.Copy(n, o)
}

