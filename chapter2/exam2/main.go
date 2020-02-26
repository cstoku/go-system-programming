package main

import (
    "os"
    "encoding/csv"
)

func main() {
    f, _ := os.OpenFile("test.csv", os.O_RDWR|os.O_CREATE, 0644)
    defer f.Close()
    c := csv.NewWriter(f)
    c.Write([]string{"a", "b", "c", "d"})
    c.Flush()
}

