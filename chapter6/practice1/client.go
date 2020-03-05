package main

import (
    "os"
    "net"
    "net/http"
    "net/http/httputil"
    "bufio"
)

func main() {
    conn, err := net.Dial("tcp", "127.0.0.1:8080")
    if err != nil {
        panic(err)
    }
    request, err := http.NewRequest("GET", "http://127.0.0.1:8080", nil)
    request.Write(conn)
    response, err := http.ReadResponse(bufio.NewReader(conn), request)
    if err != nil {
        panic(err)
    }
    dump, err := httputil.DumpResponse(response, true)
    if err != nil {
        panic(err)
    }
    os.Stdout.Write(dump)
}
