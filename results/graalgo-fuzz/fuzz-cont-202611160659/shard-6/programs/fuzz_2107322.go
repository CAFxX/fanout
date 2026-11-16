package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int16 }
func genFn0(p0 uint32, p1 uint8, p2 int8) int64 {
    var v6 int = -989
    v7 := []int{(v6 % (((v6 + -118)) | 1)), (v6 - v6), (v6 | int(p1)), (v6 | 2147483647)}
    var v8 int32 = 2147483647
    var v9 uint8 = 128
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return -342
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 0
    var v3 int64 = -328
    var v4 int8 = 105
    v5 := make(map[uint32]uint32)
    if false {
        var v10 uint8 = 120
        _ = v10
    }
    switch 42 {
    case 4294967295:
        v11 := []int{0, 127, -808, 224}
        _ = v11
    default:
        var v12 int64 = (v3 * v3)
        _ = v12
    }
    v2 += (^(v2))
    v2 = 42
    v4 += (v4 ^ (v4 | (v4 / ((v4) | 1))))
    var v13 uint8 = 167
    var v14 uint16 = 136
    var v15 string = ""
    v16 := GenS1{F0: 42, F1: "", F2: -111}
    fmt.Println(v14, v2, v14)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
