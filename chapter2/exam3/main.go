package main

import (
    "compress/gzip"
    "encoding/json"
    "io"
    "net/http"
    "os"
    "fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Log: %s %s %s\n", r.Method, r.URL, r.Proto)
    w.Header().Set("Content-Encoding", "gzip")
    w.Header().Set("Content-Type", "application/json")
    source := map[string]string{
        "Hello": "World",
    }
    g := gzip.NewWriter(w)
    m := io.MultiWriter(os.Stdout, g)
    e := json.NewEncoder(m)
    e.Encode(source)
    g.Flush()
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
