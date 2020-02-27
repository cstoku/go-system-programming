package main

import (
    "strings"
    "io"
    "os"
)

var (
    computer    = strings.NewReader("COMPUTER")
    system      = strings.NewReader("SYSTEM")
    programming = strings.NewReader("PROGRAMMING")
)

func main() {
    var stream io.Reader

    // Write code here.
    a := io.NewSectionReader(programming, 5, 1)
    s := io.LimitReader(system, 1)
    c := io.LimitReader(computer, 1)
    i := io.NewSectionReader(programming, 8, 1)

    r, w := io.Pipe()
    mw := io.MultiWriter(w, w)
    go io.Copy(mw, i)
    defer w.Close()
    ii := io.LimitReader(r, 2)

    stream = io.MultiReader(a, s, c, ii)

    // Output "ASCII"
    io.Copy(os.Stdout, stream)
}
