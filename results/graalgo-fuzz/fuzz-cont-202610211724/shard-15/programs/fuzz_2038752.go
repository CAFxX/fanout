package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint32 }
type GenS2 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int }
func genFn0(p0 uint16, p1 uint64) int {
    var v4 uint64 = (p1 - p1)
    v5 := GenS1{F0: 54, F1: 821, F2: 2147483647}
    var v6 float32 = 1.5
    fmt.Println(p0, v4, v4)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return -1
}

func main() {
    var v1 int8 = 3
    v2 := []uint32{417, 423, 2147483647, 2147483647}
    v3 := []int{}
    _ = len(v3)
    fmt.Println(v1)
    _ = len(v2)
    if (true && (30 != 972)) {
        v7 := GenS1{F0: 252, F1: -321, F2: 4294967295}
        fmt.Println(v1)
        _ = v7
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 uint64 = 1
        fmt.Println(v1)
        v9 -= v9
        _ = fi8
        _ = v9
    }
    v1 = v1
    v1 = (v1 / (((v1 ^ v1)) | 1))
    v10 := GenS2{F0: 0, F1: 42, F2: 1}
    fmt.Println(v1)
    if (false || false) {
        var v11 float64 = (-2.25 * 0.1)
        v11 = v11
        _ = v11
    } else {
        if false {
            v1 = (v1 ^ v1)
            v12 := GenS2{F0: 973, F1: 641, F2: 18446744073709551615}
            _ = v12
        }
        _ = len(v3)
        v13 := GenS2{F0: -785, F1: 32767, F2: 4294967295}
        _ = v13
    }
    if true {
        fmt.Println(v1)
        switch 42 {
        case 32768:
            v1 += v1
        case -1:
            var v14 float32 = 0.1
            _ = v14
        default:
            for fi15 := 0; fi15 < 9; fi15++ {
                v16 := GenS3{F0: genFn0(32767, 2)}
                _ = fi15
                _ = v16
            }
        }
    }
    v1 = 76
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
}
