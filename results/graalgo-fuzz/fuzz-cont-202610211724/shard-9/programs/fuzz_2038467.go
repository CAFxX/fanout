package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 255
    var v3 bool = false
    var v4 uint32 = 2147483647
    fmt.Println(v4)
    switch "foo bar" {
    case "foo bar":
        v2 = (^(v2))
    default:
        v4 = (v4 | v4)
    }
    v5 := []int{-1, 782, 42}
    var v6 int = 202
    v7 := GenS1{F0: -480, F1: -1}
    var v8 int64 = 0
    fmt.Println(v2)
    var v9 int8 = -128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
