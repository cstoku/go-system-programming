package main

import (
    "os"
    "fmt"
    "net"
    "net/http"
    "net/http/httputil"
    "bufio"
    "strings"
)

var data = []string{
    "foo",
    "bar",
    "baz",
}

func main() {
    var (
        conn net.Conn
        err  error
        cnt  int
    )

    for {
        if conn == nil {
            conn, err = net.Dial("tcp", "127.0.0.1:8080")
            if err != nil {
                panic(err)
            }
            fmt.Println("Connected.")
        }
        request, err := http.NewRequest("POST", "http://127.0.0.1:8080", strings.NewReader(data[cnt]))
        if err != nil {
            panic(err)
        }

        err = request.Write(conn)
        if err != nil {
            panic(err)
        }

        response, err := http.ReadResponse(bufio.NewReader(conn), request)
        if err != nil {
            fmt.Println("Retry...")
            conn = nil
            continue
        }
        dump, err := httputil.DumpResponse(response, true)
        if err != nil {
            panic(err)
        }
        os.Stdout.Write([]byte("\n==========\n"))
        os.Stdout.Write(dump)
        if cnt++; cnt >= len(data) {
            break
        }
    }
    conn.Close()
}
