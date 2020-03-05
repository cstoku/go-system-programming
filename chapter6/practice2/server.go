package main

import (
    "io"
    "io/ioutil"
    "os"
    "net"
    "fmt"
    "time"
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
            defer fmt.Println("Connection Closed.")
            defer conn.Close()

            for {
                conn.SetReadDeadline(time.Now().Add(5 * time.Second))
                request, err := http.ReadRequest(bufio.NewReader(conn))
                if err != nil {
                    neterr, ok := err.(net.Error)
                    if ok && neterr.Timeout() {
                        break
                    } else if err == io.EOF {
                        break
                    }
                    panic(err)
                }

                dump, err := httputil.DumpRequest(request, true)
                if err != nil {
                    panic(err)
                }
                os.Stdout.Write(dump)
                os.Stdout.Write([]byte("\n==========\n"))

                body := "Hello World!!\n"
                responce := http.Response{
                    StatusCode: 200,
                    ProtoMajor: 1,
                    ProtoMinor: 1,
                    ContentLength: int64(len(body)),
                    Body: ioutil.NopCloser(strings.NewReader(body)),
                }
                responce.Write(conn)
            }
        }()
    }
}

