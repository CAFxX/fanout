package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int8 }

func main() {
    var v1 uint32 = 643
    var v2 int64 = 0
    var v3 bool = false
    v4 := make(map[string]int)
    var v5 uint32 = v1
    var v6 string = "foo bar"
    v1 -= 1
    v3 = v3
    fmt.Println(v5)
    fmt.Println(v2, v3, v6)
    var v7 int16 = 1
    v7 = v7
    fmt.Println(v5)
    v8 := []int{-893, 128, 127, -9223372036854775808}
    v9 := GenS1{F0: 138, F1: v6, F2: 127}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
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
