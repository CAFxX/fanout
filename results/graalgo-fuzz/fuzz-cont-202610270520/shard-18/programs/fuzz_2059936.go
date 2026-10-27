package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }

func main() {
    var v1 uint16 = 42
    var v2 int16 = 1
    var v3 bool = (v1 >= v1)
    v2 = (v2 / (((v2 / ((-1) | 1))) | 1))
    v1 = v1
    v4 := []int{42, 9223372036854775807, 4294967295, 9223372036854775807}
    var v5 uint16 = (v1 | v1)
    fmt.Println(v3)
    switch "\n\t" {
    case "\n\t":
        v6 := GenS1{F0: 133, F1: 296}
        _ = v6
    case "hello":
        var v7 int8 = 39
        _ = v7
    }
    var v8 float32 = float32(0.0)
    v4 = append(v4, -130)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
