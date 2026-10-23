package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := []string{"\n\t", "a", "hello", "hello"}
    var v2 string = ("a" + "\n\t")
    var v3 float64 = (3.14159 / -2.25)
    v4 := GenS1{F0: 32768}
    v1 = append(v1, (((v2 + v2) + "foo bar") + v2))
    var v5 int32 = -540
    v5 = v5
    switch v2 {
    case "\n\t":
        v3 -= v3
    }
    v2 = v2
    v2 = "foo bar"
    var v6 float64 = 1e10
    var v7 uint = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
