package main

import "os"

func main() {
    f, _ := os.Create("test.txt")
    f.Write([]byte("hoge"))
    f.Close()
}