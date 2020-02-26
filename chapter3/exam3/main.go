package main

import (
    "os"
    "io"
    "strings"
    "archive/zip"
)

func main() {
    f, _ := os.Create("tmp.zip")
    zipwriter := zip.NewWriter(f)
    defer zipwriter.Close()

    w, _ := zipwriter.Create("tmp.txt")
    r := strings.NewReader("hogehoge")
    io.Copy(w, r) 
}
