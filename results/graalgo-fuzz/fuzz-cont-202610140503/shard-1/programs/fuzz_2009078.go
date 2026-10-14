package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 uint32; F2 int }
type GenS3 struct { F0 int8; F1 string; F2 bool }
type GenS4 struct { F0 int32; F1 uint16 }
func genFn0() bool {
    if true {
        if false {
            var v6 int64 = 0
            _ = v6
        }
    } else {
        var v7 uint8 = 18
        for fi8 := 0; fi8 < 16; fi8++ {
            v9 := GenS1{F0: 1}
            _ = fi8
            _ = v9
        }
        v10 := GenS2{F0: 959, F1: 2147483648, F2: -2147483648}
        _ = v10
        _ = v7
    }
    if (2147483648 > 42) {
        v11 := GenS3{F0: -1, F1: "\n\t", F2: false}
        var v12 uint = 42
        var v13 float64 = 1e10
        _ = v11
        _ = v12
        _ = v13
    } else {
        var v14 int64 = -9223372036854775808
        _ = v14
    }
    return ((1.5 + 0.0) >= float32(0.1))
}

func main() {
    var v1 int = 128
    var v2 uint = 318
    v3 := []uint32{187}
    var v4 uint64 = 18446744073709551615
    v5 := make(map[uint32]bool)
    var v15 bool = true
    v4 = (v4 ^ (v4 / ((0) | 1)))
    v16 := v5[158]
    v3 = append(v3, 11)
    var v17 int8 = 40
    if (746 >= uint32(v2)) {
        var v18 uint16 = 1
        delete(v5, 0)
        _ = v18
    } else {
        v1 += 564
    }
    var v19 uint16 = 543
    v17 = v17
    v5[77] = (384 == 1)
    v20 := GenS4{F0: 2147483647, F1: (v19 ^ v19)}
    var v21 uint32 = 430
    fmt.Println(v19, v2)
    var v22 int16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
