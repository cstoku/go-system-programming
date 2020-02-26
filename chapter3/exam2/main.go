package main

import (
    "os"
    "crypto/rand"
)

var (
    MaxBufferSize = 200
    ReadSize = 1500
)

func main() {
    f, _ := os.OpenFile("rand.txt", os.O_RDWR|os.O_CREATE, 0644)
    defer f.Close()

    buf := make([]byte, MaxBufferSize)

    for cnt := 0; cnt < ReadSize; cnt += MaxBufferSize {
        var readSize = 0
        if ReadSize - cnt - MaxBufferSize > 0 {
            readSize = MaxBufferSize
        } else {
            readSize = ReadSize - cnt
        }
        c, _ := rand.Read(buf[:readSize])
        f.Write(buf[:c])
    }
}

