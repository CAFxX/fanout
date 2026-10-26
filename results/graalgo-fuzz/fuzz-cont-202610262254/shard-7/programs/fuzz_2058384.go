package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint32; F1 uint32; F2 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8; F1 int8; F2 uint8 }
func genFn0(p0 int16) uint32 {
    var v7 int8 = -90
    switch -1 {
    case 381:
        v8 := GenS1{F0: 0, F1: -991, F2: 2}
        _ = v8
    case 9223372036854775807:
        p0 += p0
    case 0:
        v9 := GenS1{F0: 665, F1: 2147483647, F2: 4294967295}
        _ = v9
    }
    switch "foo bar" {
    case "\n\t":
        var v10 uint64 = 733
        _ = v10
    case "a":
        var v11 float64 = (float64(((((1e10 / 1.5) * 1e10) / (0.1 + 0.0)) / 1e10)) * (1e10 + 1.5))
        _ = v11
    case "hello":
        var v12 uint = 0
        _ = v12
    default:
        if (23 >= 91) {
            for fi13 := 0; fi13 < 12; fi13++ {
                if (1 > -1) {
                    v7 = v7
                    var v14 float32 = -2.25
                    v7 = v7
                    _ = v14
                }
                _ = fi13
            }
        }
    }
    var v15 uint8 = 128
    _ = p0
    _ = v15
    _ = v7
    return 0
}

func main() {
    var v1 uint64 = 4294967295
    var v2 float64 = 3.14159
    var v3 uint32 = 42
    var v4 int32 = 2147483647
    var v5 string = ""
    var v6 uint16 = 2
    v16 := GenS1{F0: v4, F1: v4, F2: (v1 % (((v1 << ((uint(v1)) & 63))) | 1))}
    v3 -= (v3 * 0)
    for fi17 := 0; fi17 < 9; fi17++ {
        v18 := GenS2{F0: v1, F1: (v3 | (v3 >> ((403) & 63)))}
        _ = fi17
        _ = v18
    }
    v1 = (^(v1))
    fmt.Println(v5, v5)
    var v19 uint8 = 8
    fmt.Println(v6)
    if false {
        for fi20 := 0; fi20 < 11; fi20++ {
            var v21 uint64 = (v1 / ((v1) | 1))
            for fi22 := 0; fi22 < 14; fi22++ {
                var v23 int64 = 42
                _ = fi22
                _ = v23
            }
            if false {
                v24 := GenS3{F0: 4294967295, F1: uint32(v19), F2: 4294967295}
                if false {
                    v21 = uint64(v3)
                    var v25 bool = false
                    _ = v25
                } else {
                    v26 := GenS1{F0: v4, F1: (v4 + (v4 + v4)), F2: v1}
                    v27 := []int{}
                    v19 -= v19
                    _ = v26
                    _ = v27
                }
                v28 := GenS4{F0: (v6 << ((18446744073709551615) & 63))}
                _ = v24
                _ = v28
            } else {
                v29 := GenS5{F0: 0, F1: 127, F2: (^(v19))}
                v21 = v21
                _ = v29
            }
            _ = fi20
            _ = v21
        }
        var v30 uint64 = 42
        fmt.Println(v3)
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
