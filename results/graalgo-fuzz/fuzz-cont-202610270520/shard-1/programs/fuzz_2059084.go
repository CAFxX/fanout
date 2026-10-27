package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }

func main() {
    v1 := make(map[string]int)
    v2 := []string{}
    var v3 float64 = (0.1 / (1e10 - (-2.25 - float64(1.5))))
    var v4 uint16 = 914
    var v5 uint = 2
    var v6 int64 = -1
    var v7 int8 = 100
    var v8 uint32 = 4294967295
    fmt.Println(v4, v6, v6)
    var v9 float32 = 3.14159
    var v10 int64 = -9223372036854775808
    v11 := GenS1{F0: 18446744073709551615, F1: 293}
    fmt.Println(v4, v7)
    var v12 int16 = 468
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
