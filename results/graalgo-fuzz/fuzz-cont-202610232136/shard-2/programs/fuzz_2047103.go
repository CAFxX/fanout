package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 float64 = float64((float32(0.0) + -2.25))
    v2 := []string{"\n\t", ""}
    var v3 uint64 = 0
    v3 = v3
    v3 -= (^(v3))
    fmt.Println(v3)
    v4 := GenS1{F0: 32768}
    var v5 float64 = (0.0 - 0.0)
    _ = len(v2)
    var v6 float32 = float32(v5)
    var v7 int = 2147483647
    var v8 uint16 = 42
    v9 := GenS1{F0: v7}
    var v10 uint = 2
    fmt.Println(v8, v5, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
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
    _ = v9
}
