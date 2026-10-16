package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int }
type GenS2 struct { F0 uint64; F1 bool; F2 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 bool; F2 uint32 }
func genFn0(p0 uint32) bool {
    var v7 int8 = 0
    v8 := GenS1{F0: 93, F1: (-1 > int32(v7)), F2: 9223372036854775807}
    _ = p0
    _ = v7
    _ = v8
    return (1 < 1)
}
func genFn1(p0 uint32, p1 uint8) bool {
    var v9 int32 = 414
    v10 := GenS1{F0: (p1 - uint8(p0)), F1: false, F2: 128}
    v9 = v9
    v11 := GenS2{F0: uint64(p1), F1: (p0 >= 2147483647), F2: 4294967295}
    var v12 uint16 = 1
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return ((false || false) && false)
}

func main() {
    var v1 int = 209
    v2 := []int{v1, v1, 2147483647}
    v3 := make(map[int]int)
    v4 := []int{951, (v1 + v1)}
    v5 := []int{v1}
    var v6 uint64 = 477
    var v13 bool = true
    if false {
        var v14 float32 = float32(float64((((1e10 * (((1.5 + (1e10 + float32((((3.14159 * 1.5) * 1e10) / 1.0)))) / (1e10 - 0.0)) / 1.5)) * (0.1 * 0.1)) - -2.25)))
        var v15 uint64 = (v6 | 4294967295)
        _ = v14
        _ = v15
    }
    var v16 uint16 = 32767
    switch -2147483648 {
    case 128:
        v13 = true
    default:
        var v17 float64 = -2.25
        _ = v17
    }
    if (42 >= 51) {
        v13 = (!(3.14159 >= 0.1))
    }
    v18 := GenS3{F0: 1}
    _ = len(v3)
    var v19 int32 = 2147483647
    v20 := GenS1{F0: 255, F1: v13, F2: (^(v1))}
    v21 := GenS2{F0: v6, F1: v13, F2: 230}
    if (3.14159 <= 0.0) {
        for fi22 := 0; fi22 < 12; fi22++ {
            v23 := GenS4{F0: -1, F1: v13, F2: 448}
            _ = fi22
            _ = v23
        }
    }
    _ = len(v3)
    var v24 int16 = 116
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
