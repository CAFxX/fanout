package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 uint16 = 65535
    v2 := make(map[string]bool)
    var v3 float64 = 0.0
    var v4 int16 = 1000
    v5 := make(map[int]bool)
    v6 := GenS1{F0: 0}
    switch "" {
    case "a":
        var v7 int32 = 42
        _ = v7
    case "\n\t":
        v8 := []int{}
        _ = v8
    default:
        v3 = float64(1.5)
    }
    v9 := GenS1{F0: -2147483648}
    v1 = 65535
    fmt.Println(v4, v1, v1)
    v10 := []int{-833, 32767, -777}
    var v11 uint = 408
    v4 = (v4 + -1)
    v12 := GenS1{F0: int32(v11)}
    fmt.Println(v1, v11, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
