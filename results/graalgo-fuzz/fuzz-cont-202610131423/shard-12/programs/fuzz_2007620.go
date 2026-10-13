package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 int16 = -1
    var v2 float32 = 1e10
    var v3 int8 = -1
    v4 := []string{"foo bar", "a"}
    var v5 float64 = 0.0
    v6 := make(map[int]bool)
    fmt.Println(v3, v5, v5)
    fmt.Println(v5, v1)
    var v7 int8 = (v3 - v3)
    _ = len(v6)
    if ((v1 ^ (v1 | (v1 << ((42) & 63)))) <= v1) {
        v8 := GenS1{F0: 42}
        fmt.Println(v7, v5)
        _ = v8
    } else {
        v9 := GenS1{F0: 42}
        _ = v9
    }
    fmt.Println(v5, v1)
    fmt.Println(v7, v3, v3)
    var v10 uint64 = 0
    var v11 float64 = (v5 / -2.25)
    v6[-2147483648] = true
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
