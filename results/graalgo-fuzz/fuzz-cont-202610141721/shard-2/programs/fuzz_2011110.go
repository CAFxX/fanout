package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 int32 }
type GenS5 struct { F0 int32; F1 int64 }
func genFn0() uint32 {
    var v4 int = 2147483647
    v4 -= -1
    if (42 == 127) {
        var v5 int16 = -32768
        _ = v5
    }
    _ = v4
    return 2
}
func genFn1(p0 uint16) int8 {
    v6 := GenS1{F0: 403, F1: 1}
    v7 := GenS2{F0: 317, F1: 42}
    _ = p0
    _ = v6
    _ = v7
    return 79
}

func main() {
    var v1 int32 = 42
    v2 := []int{-9223372036854775808, 0, -9223372036854775808, -1}
    var v3 int16 = 32767
    v8 := GenS3{F0: 0}
    var v9 int16 = (v3 * (v3 % ((1000) | 1)))
    v3 = v9
    v10 := GenS4{F0: 42, F1: (v1 & -318)}
    if false {
        fmt.Println(v3, v9)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v11 uint32 = 4294967295
    v3 = 32767
    switch 32767 {
    case 128:
        for fi12 := 0; fi12 < 8; fi12++ {
            if (fi12 <= fi12) {
                for fi13 := 0; fi13 < 6; fi13++ {
                    var v14 float32 = 0.0
                    _ = len(v2)
                    v11 += v11
                    _ = fi13
                    _ = v14
                }
            } else {
                var v15 int16 = (v9 * (v3 & v3))
                _ = v15
            }
            _ = fi12
        }
    case -2147483648:
        var v16 bool = false
        _ = v16
    case 584:
        var v17 int32 = v1
        _ = v17
    default:
        for fi18 := 0; fi18 < 9; fi18++ {
            v19 := GenS5{F0: v1, F1: -272}
            _ = fi18
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 11; fi20++ {
        var v21 int = fi20
        v1 = (^(v1))
        _ = fi20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
