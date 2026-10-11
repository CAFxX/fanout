package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 int8; F1 int; F2 uint }
func genFn0(p0 int32, p1 uint16, p2 uint16) int {
    var v5 uint = 255
    v6 := GenS1{F0: 1, F1: p1, F2: (p2 << ((4294967295) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 127
}
func genFn1(p0 int32, p1 uint) uint8 {
    v7 := GenS2{F0: 0, F1: -582}
    v8 := GenS2{F0: 42, F1: p0}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 82
}

func main() {
    var v1 int16 = 1000
    var v2 bool = true
    v3 := make(map[uint32]bool)
    var v4 float64 = (1e10 * 0.1)
    fmt.Println(v1)
    var v9 bool = (!(true || false))
    switch genFn0(-789, 255, 160) {
    case 32767:
        _ = len(v3)
    case -2147483648:
        fmt.Println(v1, v2)
    default:
        v10 := GenS3{F0: (v1 / ((v1) | 1)), F1: 1}
        _ = v10
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        if v9 {
            for fi12 := 0; fi12 < 16; fi12++ {
                if ((1e10 * 3.14159) < (((1e10 / 1.0) + 3.14159) * (1e10 - -2.25))) {
                    fmt.Println(fi12, v4, v2)
                    var v13 int32 = 768
                    delete(v3, 0)
                    _ = v13
                }
                delete(v3, 16)
                _ = fi12
            }
        }
        _ = fi11
    }
    v2 = (!v9)
    v14 := GenS4{F0: 123, F1: 128, F2: 2}
    if v2 {
        v4 = (v4 * (v4 / ((v4 - v4) * v4)))
        v1 -= (v1 >> ((255) & 63))
        v3[65535] = v2
    } else {
        v4 -= -2.25
        v9 = v2
    }
    v4 -= (v4 - (v4 * v4))
    if (146 >= (v1 >> ((581) & 63))) {
        fmt.Println(v9)
    }
    v15 := GenS2{F0: (v1 << ((42) & 63)), F1: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
