package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }

func main() {
    var v1 float64 = 0.0
    v2 := []string{("a" + "a"), "", "", "\n\t"}
    v3 := []string{}
    v4 := make(map[int]int)
    var v5 int32 = -784
    _ = len(v4)
    var v6 uint64 = 18446744073709551615
    var v7 float64 = ((v1 / 1.5) - (1.5 * (1.5 - v1)))
    var v8 int64 = 721
    v9 := GenS1{F0: 128, F1: 1}
    v7 += 1.5
    var v10 int32 = 0
    switch ("a" + "a") {
    case "a":
        var v11 uint8 = 128
        _ = v11
    case "foo bar":
        delete(v4, 4294967295)
    case "hello":
        v5 = 0
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
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
