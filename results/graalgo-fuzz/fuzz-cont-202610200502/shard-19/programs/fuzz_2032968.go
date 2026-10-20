package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }

func main() {
    var v1 int = 2147483647
    var v2 int8 = -106
    var v3 uint64 = 0
    var v4 int16 = 1000
    v5 := []string{}
    v6 := []string{"foo bar"}
    v2 = (^(v2))
    var v7 float64 = 0.1
    v8 := GenS1{F0: 359, F1: "foo bar"}
    _ = len(v6)
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        _ = len(v6)
    }
    v2 += 1
    fmt.Println(v1, v7, v3)
    var v9 int16 = (v4 << ((18446744073709551615) & 63))
    v10 := GenS1{F0: -489, F1: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
