package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }

func main() {
    var v1 float64 = 0.1
    var v2 uint8 = 255
    v3 := make(map[string]string)
    v4 := make(map[uint32]int)
    v5 := make(map[string]string)
    var v6 bool = false
    var v7 uint16 = 42
    switch -9223372036854775808 {
    case -2147483648:
        delete(v4, 848)
    }
    var v8 uint16 = (^(v7))
    var v9 int16 = -603
    v6 = (v7 <= v8)
    var v10 int64 = -941
    v11 := GenS1{F0: (v9 & v9), F1: -554}
    var v12 int64 = (v10 | v10)
    v13 := GenS1{F0: v9, F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
