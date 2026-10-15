package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 string; F2 string }
type GenS3 struct { F0 uint8; F1 uint32; F2 int }
type GenS4 struct { F0 int8; F1 int64; F2 uint32 }
func genFn0(p0 string) int32 {
    var v4 uint16 = 42
    v4 = 2
    v5 := []int{}
    var v6 int32 = -361
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return -781
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 2
    v3 := []uint32{}
    if ((2 == 882) || false) {
        for fi7 := 0; fi7 < 20; fi7++ {
            var v8 uint8 = 2
            _ = fi7
            _ = v8
        }
    } else {
        v9 := GenS1{F0: ((v2 / (((v2 + (v2 & v2))) | 1)) <= (v2 / (((v2 << ((2) & 63))) | 1)))}
        _ = v9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 uint8 = 128
        fmt.Println(v2)
        fmt.Println(v2)
        _ = fi10
        _ = v11
    }
    v2 += v2
    switch 103 {
    case 358:
        var v12 uint8 = 2
        _ = v12
    case -9223372036854775808:
        for fi13 := 0; fi13 < 19; fi13++ {
            v2 = 65535
            _ = fi13
        }
    case 459:
        _ = len(v3)
    }
    if (true || (755 >= 227)) {
        for fi14 := 0; fi14 < 3; fi14++ {
            var v15 int = (fi14 & fi14)
            var v16 float32 = 1e10
            v16 += 1.5
            _ = fi14
            _ = v15
            _ = v16
        }
        var v17 int64 = -910
        v18 := GenS2{F0: -74, F1: "hello", F2: "a"}
        _ = v17
        _ = v18
    }
    fmt.Println(v2)
    fmt.Println(v2)
    var v19 float64 = 3.14159
    _ = len(v1)
    v19 -= v19
    var v20 uint64 = 2
    delete(v1, 158)
    switch "" {
    case "hello":
        v21 := GenS3{F0: 23, F1: 317, F2: 2147483647}
        _ = v21
    default:
        if false {
            fmt.Println(v20)
            if (167 != 9223372036854775807) {
                v22 := GenS3{F0: 214, F1: 2, F2: 127}
                v2 = (v2 & v2)
                var v23 int = 1
                _ = v22
                _ = v23
            }
        } else {
            v20 += v20
            var v24 float32 = 1e10
            var v25 uint8 = 27
            _ = v24
            _ = v25
        }
    }
    v26 := GenS4{F0: 119, F1: 9223372036854775807, F2: uint32(v20)}
    for fi27 := 0; fi27 < 4; fi27++ {
        var v28 bool = true
        _ = fi27
        _ = v28
    }
    fmt.Println(v2, v2, v19)
    fmt.Println(v19, v20)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v3
}
