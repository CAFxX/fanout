package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 bool }

func main() {
    v1 := make(map[int]string)
    var v2 string = "foo bar"
    var v3 int8 = 0
    v2 = v2
    var v4 uint64 = 18446744073709551615
    _ = len(v1)
    var v5 float32 = -2.25
    var v6 int8 = (v3 | (v3 - 1))
    v7 := GenS1{F0: 0, F1: 202, F2: (0.0 >= 3.14159)}
    v8 := []int{-1, 469}
    for fi9 := 0; fi9 < 20; fi9++ {
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 int64 = -9223372036854775808
            var v12 float32 = v5
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
