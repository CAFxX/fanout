package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 int8; F2 uint64 }

func main() {
    v1 := []uint32{2, 1, 0}
    var v2 float32 = 3.14159
    var v3 int = -1
    var v4 string = "a"
    var v5 float32 = (v2 + v2)
    fmt.Println(v2, v4)
    v6 := GenS1{F0: -115}
    v7 := GenS2{F0: -128, F1: 1, F2: 631}
    fmt.Println(v2)
    v2 += v5
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v8 := GenS1{F0: -7}
    v9 := GenS2{F0: -31, F1: 127, F2: 42}
    v2 += float32(0.1)
    v5 += (v2 - v5)
    if false {
        v2 = (v2 / (0.0 - v5))
    }
    var v10 uint8 = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
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
