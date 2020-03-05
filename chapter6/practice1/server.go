package main

import (
    "io/ioutil"
    "os"
    "net"
    "bufio"
    "strings"
    "net/http"
    "net/http/httputil"
)

func main() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        panic(err)
    }
    defer listener.Close()
    for {
        conn, err := listener.Accept()
        if err != nil {
            panic(err)
        }
        go func() {
            request, err := http.ReadRequest(bufio.NewReader(conn))
            if err != nil {
                panic(err)
            }

            dump, err := httputil.DumpRequest(request, true)
            if err != nil {
                panic(err)
            }
            os.Stdout.Write(dump)

            body := "Hello World!!"
            responce := http.Response{
                StatusCode: 200,
                ProtoMajor: 1,
                ProtoMinor: 0,
                Body: ioutil.NopCloser(strings.NewReader(body)),
            }
            responce.Write(conn)
            conn.Close()
        }()
    }
}

