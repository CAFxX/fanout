package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }

func main() {
    v1 := []string{}
    var v2 string = "\n\t"
    v3 := make(map[int]bool)
    var v4 bool = (1 < 18446744073709551615)
    var v5 int = 4294967295
    var v6 bool = (!v4)
    switch "hello" {
    case "hello":
        var v7 int64 = 425
        _ = v7
    case "a":
        fmt.Println(v6)
    case "\n\t":
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := GenS1{F0: 217, F1: v2}
            _ = fi8
            _ = v9
        }
    default:
        v1 = append(v1, v2)
    }
    var v10 float64 = 1e10
    _ = len(v1)
    fmt.Println(v5)
    var v11 bool = v6
    _ = len(v3)
    v5 = (v5 + v5)
    v12 := GenS1{F0: 42, F1: v2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
