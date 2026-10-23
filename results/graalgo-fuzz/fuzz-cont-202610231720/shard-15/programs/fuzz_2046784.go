package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = 2147483647
    v3 := make(map[string]bool)
    var v4 uint8 = 0
    v5 := GenS1{F0: 9223372036854775807, F1: (1 <= 4294967295)}
    fmt.Println(v2)
    v4 -= (^(v4))
    var v6 int16 = 42
    v7 := []int{-1}
    v4 = v4
    var v8 float32 = 0.0
    v2 = v2
    if (1 < 42) {
        var v9 string = "foo bar"
        _ = v9
    }
    v4 = v4
    fmt.Println(v2, v2, v4)
    var v10 bool = (1 <= -1)
    v6 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
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
}
