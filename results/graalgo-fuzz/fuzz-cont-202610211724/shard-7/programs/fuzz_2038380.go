package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int8 }

func main() {
    v1 := []string{}
    var v2 uint = 109
    v3 := make(map[uint32]bool)
    var v4 string = "foo bar"
    switch -2147483648 {
    case 32768:
        v4 = "foo bar"
    default:
        v3[2] = (float32(0.0) != ((1e10 * (-2.25 - 1.5)) + (3.14159 / 0.1)))
    }
    v4 = v4
    fmt.Println(v4)
    var v5 uint32 = 42
    v6 := GenS1{F0: 42, F1: (v5 * v5)}
    v7 := v3[v5]
    var v8 bool = false
    v5 = v5
    switch 42 {
    case 1:
        v9 := GenS2{F0: 42}
        _ = v9
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
