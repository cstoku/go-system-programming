package main

import (
    "io"
    "fmt"
    "errors"
    "strings"
    "bytes"
)

func CopyN(dst io.Writer, src io.Reader, n int) (int, error) {
    BUFSIZE := 64
    buf := make([]byte, BUFSIZE)
    var written int

    if dst == nil || src == nil {
        return 0, errors.New("src or dst is nil...")
    }

    for {
        cnt, err := src.Read(buf)
        switch {
            case err == io.EOF:
                return written, nil
            case err != nil:
                return 0, err
        }

        if cnt < n {
            n -= cnt
        } else {
            cnt = n
        }
        w, err := dst.Write(buf[:cnt])
        written += w
    }
}

func main() {
    s := strings.NewReader("hogehoge")
    var b bytes.Buffer
    n, _ := CopyN(&b, s, 5)
    fmt.Printf("%v: %v\n", n, b.String())
}
