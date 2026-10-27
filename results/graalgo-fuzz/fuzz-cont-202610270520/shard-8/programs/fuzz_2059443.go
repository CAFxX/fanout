package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 string }

func main() {
    v1 := []int{4294967295, 42}
    var v2 float32 = 0.1
    var v3 uint32 = 1
    v4 := make(map[string]string)
    v5 := make(map[uint32]int)
    switch -233 {
    case 9223372036854775807:
        var v6 uint16 = 32767
        _ = v6
    }
    switch 876 {
    case 0:
        var v7 float32 = float32(float64(1e10))
        _ = v7
    case 9223372036854775807:
        for fi8 := 0; fi8 < 4; fi8++ {
            for fi9 := 0; fi9 < 9; fi9++ {
                var v10 int64 = 0
                var v11 int16 = -32768
                fmt.Println(v11, fi9)
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi8
        }
    default:
        v3 -= (v3 | (v3 % (((v3 << ((1) & 63))) | 1)))
    }
    if (!false) {
        if len(v1) > 0 {
            v1[0] = 9223372036854775807
        }
        var v12 uint32 = v3
        _ = len(v1)
        _ = v12
    } else {
        v13 := GenS1{F0: 32767, F1: true, F2: ""}
        var v14 uint = 1
        fmt.Println(v3, v3, v14)
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: 449, F1: (-2.25 < (0.0 * 1.5)), F2: ""}
    fmt.Println(v3, v3)
    _ = len(v1)
    v3 -= v3
    var v16 string = "hello"
    v1 = append(v1, -1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
