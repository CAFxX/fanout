package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 42
    var v3 uint16 = (^(v2))
    switch "\n\t" {
    case "":
        v2 += 42
    case "foo bar":
        var v4 uint16 = 1
        _ = v4
    default:
        fmt.Println(v2)
    }
    v2 = v2
    var v5 uint = 25
    v6 := GenS1{F0: (v5 & v5), F1: 1}
    var v7 int8 = -128
    v8 := GenS1{F0: v5, F1: 4294967295}
    var v9 uint64 = 42
    fmt.Println(v5, v3, v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
