package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int64; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint8; F1 int8; F2 int16 }
func genFn0(p0 uint16) uint64 {
    v4 := []int{975, 2147483647}
    v5 := GenS1{F0: 597, F1: 32767}
    _ = p0
    _ = v4
    _ = v5
    return 825
}
func genFn1(p0 int64, p1 uint16) int8 {
    switch "hello" {
    case "\n\t":
        v6 := GenS2{F0: p0, F1: 127, F2: 153}
        _ = v6
    case "":
        v7 := []int{}
        _ = v7
    default:
        var v8 uint64 = 2
        _ = v8
    }
    var v9 float64 = (float64((1e10 - 3.14159)) * 3.14159)
    v10 := GenS2{F0: 1, F1: 213, F2: 2}
    if (0 < 0) {
        v11 := GenS1{F0: -928, F1: 1000}
        v12 := GenS1{F0: -1, F1: 1}
        _ = v11
        _ = v12
    } else {
        var v13 int = 127
        for fi14 := 0; fi14 < 13; fi14++ {
            v15 := GenS3{F0: 1, F1: 60, F2: 409}
            _ = fi14
            _ = v15
        }
        var v16 bool = true
        _ = v13
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 0
}

func main() {
    v1 := []int{-9223372036854775808}
    var v2 uint64 = 4294967295
    var v3 uint64 = (v2 * (v2 % ((v2) | 1)))
    fmt.Println(v3)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v17 uint64 = v3
    v1 = append(v1, 796)
    v18 := GenS1{F0: -1, F1: 150}
    _ = len(v1)
    if len(v1) > 0 {
        v1[0] = -206
    }
    var v19 int16 = -32768
    fmt.Println(v19, v19)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
