package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS5 struct { F0 uint64; F1 bool; F2 bool }
func genFn0(p0 int) int8 {
    p0 -= (p0 | p0)
    var v4 int = p0
    _ = p0
    _ = v4
    return -128
}

func main() {
    v1 := make(map[string]string)
    var v2 int64 = 755
    var v3 string = "hello"
    var v5 int16 = 42
    switch ((v3 + v3) + v3) {
    case "\n\t":
        var v6 uint8 = 84
        _ = v6
    default:
        var v7 int32 = 2147483647
        _ = v7
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        for fi9 := 0; fi9 < 6; fi9++ {
            fmt.Println(v3)
            v3 = ""
            v10 := GenS1{F0: 1}
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    switch 9223372036854775807 {
    case 2147483647:
        var v11 uint64 = 276
        _ = v11
    case 4294967295:
        v12 := v1[(v3 + "\n\t")]
        _ = v12
    case 0:
        if true {
            v13 := GenS2{F0: 65535, F1: 64}
            v14 := GenS3{F0: 127, F1: (v5 ^ v5)}
            fmt.Println(v5)
            _ = v13
            _ = v14
        } else {
            v2 -= (v2 | (v2 | v2))
        }
    }
    v5 -= v5
    v15 := GenS4{F0: 994, F1: 854, F2: -2147483648}
    delete(v1, v3)
    v16 := GenS5{F0: 42, F1: ((!true) && ((151 <= 2) || true)), F2: true}
    var v17 uint32 = 42
    if true {
        var v18 float64 = 1.5
        var v19 float64 = v18
        _ = v18
        _ = v19
    } else {
        var v20 uint16 = 110
        _ = v20
    }
    var v21 int = -668
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v5
}
