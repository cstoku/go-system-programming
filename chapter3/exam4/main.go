package main

import (
    "fmt"
    "net/http"
    "archive/zip"
)

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/zip")
    w.Header().Set("Content-Disposition", "attachment; filename=ascii_sample.zip")
    zipwriter := zip.NewWriter(w)
    defer zipwriter.Close()
    f, _ := zipwriter.Create("tmp.txt")
    fmt.Fprintf(f, "hogehoge\n")
    
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
