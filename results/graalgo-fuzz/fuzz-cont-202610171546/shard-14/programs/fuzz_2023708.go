package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 int16; F2 uint16 }
func genFn0(p0 uint32, p1 uint32) uint8 {
    p0 -= p1
    var v7 uint8 = 55
    for fi8 := 0; fi8 < 8; fi8++ {
        fmt.Println(v7, v7)
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v7
    return 128
}

func main() {
    var v1 uint32 = 362
    var v2 float32 = 0.1
    var v3 int = 4294967295
    var v4 int64 = 0
    v5 := []uint32{(v1 & v1)}
    v6 := make(map[uint32]string)
    var v9 bool = true
    var v10 float64 = 3.14159
    fmt.Println(v3, v2, v3)
    _ = len(v5)
    if ((!(v9 || ((v9 && false) && v9))) && v9) {
        var v11 float32 = v2
        _ = len(v5)
        _ = v11
    }
    var v12 uint = 739
    v9 = false
    v4 = 0
    v13 := GenS1{F0: uint16(v12)}
    v4 = (v4 / ((v4) | 1))
    if (v10 <= (float64((float32(float64(v2)) * v2)) * v10)) {
        v14 := v6[v1]
        v15 := GenS1{F0: uint16(v12)}
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: 32767}
    v17 := GenS2{F0: 1, F1: -1, F2: 1}
    fmt.Println(v9)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
