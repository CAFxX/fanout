package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int64 }

func main() {
    var v1 int32 = 1
    var v2 int32 = v1
    var v3 uint32 = 92
    var v4 uint16 = 703
    v5 := GenS1{F0: true, F1: (^(v2)), F2: 0}
    v1 += (v2 * v1)
    v6 := []int{-147, -1, 4294967295, 127}
    var v7 int16 = 32767
    v8 := GenS1{F0: false, F1: (^(v2)), F2: 1}
    var v9 float64 = 0.1
    var v10 uint8 = uint8(v3)
    for fi11 := 0; fi11 < 3; fi11++ {
        fmt.Println(fi11, v1)
        _ = fi11
    }
    var v12 float32 = 1e10
    v2 -= (v2 | v2)
    var v13 int64 = 0
    var v14 int64 = v13
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
