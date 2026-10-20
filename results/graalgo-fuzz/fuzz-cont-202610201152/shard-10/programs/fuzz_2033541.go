package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint8 }

func main() {
    v1 := make(map[int]string)
    var v2 float64 = float64(float32(1.5))
    var v3 int = 9223372036854775807
    v4 := make(map[string]int)
    var v5 uint = 42
    var v6 uint64 = 2
    v3 = v3
    v6 -= (v6 & 303)
    v7 := GenS1{F0: 237, F1: 60, F2: 2}
    var v8 float32 = 1.5
    delete(v4, "hello")
    v6 = 951
    fmt.Println(v6, v6)
    var v9 uint64 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
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
