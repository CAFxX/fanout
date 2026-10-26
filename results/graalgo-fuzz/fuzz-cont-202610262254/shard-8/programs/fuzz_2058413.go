package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int64 }

func main() {
    var v1 float64 = 1e10
    var v2 uint32 = 96
    v3 := []uint32{(v2 - (v2 ^ 42)), (^(v2))}
    var v4 int8 = 42
    for fi5 := 0; fi5 < 19; fi5++ {
        _ = len(v3)
        _ = fi5
    }
    v4 -= -128
    var v6 float32 = float32(1.5)
    for fi7 := 0; fi7 < 12; fi7++ {
        _ = len(v3)
        _ = fi7
    }
    v8 := GenS1{F0: v2, F1: 127, F2: -622}
    var v9 uint16 = 433
    var v10 float32 = float32(v1)
    var v11 int64 = -1
    fmt.Println(v4, v9, v2)
    var v12 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v8
    _ = v9
}
