package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint32 }
func genFn0(p0 uint8) int8 {
    p0 += 42
    var v4 int64 = -1
    var v5 int64 = (v4 | v4)
    var v6 uint64 = 18446744073709551615
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return -94
}

func main() {
    v1 := []uint32{42, 4294967295}
    var v2 float64 = (-2.25 / -2.25)
    v3 := make(map[string]string)
    var v7 uint8 = 2
    fmt.Println(v2, v2)
    v1 = append(v1, 647)
    v2 = (v2 + 0.0)
    fmt.Println(v7, v7)
    v8 := GenS1{F0: -5, F1: -2147483648, F2: 1}
    switch 1 {
    case 1:
        v7 += (v7 / (((v7 ^ v7)) | 1))
    }
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    if false {
        fmt.Println(v7)
        _ = len(v1)
    }
    v2 += -2.25
    v2 -= v2
    if (1 != 0) {
        delete(v3, "a")
    } else {
        switch 280 {
        case 893:
            var v9 int32 = 244
            _ = v9
        default:
            for fi10 := 0; fi10 < 6; fi10++ {
                v11 := GenS1{F0: fi10, F1: (fi10 / ((-2147483648) | 1)), F2: 0}
                _ = fi10
                _ = v11
            }
        }
    }
    v7 += 255
    fmt.Println(v7)
    _ = len(v1)
    if len(v1) > 0 {
        v1[0] = 65535
    }
    v2 += ((((v2 - (((float64((-2.25 * 0.0)) * v2) / v2) + v2)) * (float64(3.14159) * 0.1)) / v2) * float64(3.14159))
    v2 -= v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
