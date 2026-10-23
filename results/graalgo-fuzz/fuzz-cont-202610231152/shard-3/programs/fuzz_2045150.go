package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint; F1 int16 }
func genFn0(p0 bool, p1 uint64, p2 int64) uint64 {
    var v5 uint32 = 42
    var v6 int8 = 127
    v7 := []int{-819, -945, int(v5), 2147483647}
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = 0.1
    v3 := []int{}
    v4 := []int{9223372036854775807, 9223372036854775807, -305}
    v2 -= (float64(3.14159) / float64(1e10))
    delete(v1, 179)
    v8 := GenS1{F0: -128}
    var v9 uint32 = 0
    v2 += v2
    fmt.Println(v9)
    for fi10 := 0; fi10 < 10; fi10++ {
        v2 = 3.14159
        switch "\n\t" {
        case "":
            v9 = (v9 / ((42) | 1))
        case "\n\t":
            v11 := GenS2{F0: 255, F1: 472}
            _ = v11
        default:
            v9 -= v9
        }
        if true {
            _ = len(v4)
            v2 += (1e10 * ((v2 * v2) - v2))
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v9 = v9
        }
        _ = fi10
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        switch fi12 {
        case -2147483648:
            var v13 uint8 = 20
            _ = v13
        case 0:
            v1[(v9 << ((664) & 63))] = (v9 % (((v9 * v9)) | 1))
        default:
            var v14 uint32 = v9
            _ = v14
        }
        fmt.Println(fi12)
        _ = fi12
    }
    var v15 uint8 = 1
    switch 505 {
    case -9223372036854775808:
        for fi16 := 0; fi16 < 17; fi16++ {
            v17 := GenS1{F0: -128}
            v18 := GenS3{F0: uint(v9), F1: 203}
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
