package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 int; F2 int32 }
type GenS3 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 int32; F1 int8; F2 int16 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 810
    v3 := []int{-1, -552, -2147483648, 127}
    var v4 float64 = (3.14159 * 1.5)
    var v5 int16 = -343
    v6 := GenS1{F0: 1}
    if (9223372036854775807 != 0) {
        var v7 int = 768
        v4 = (1e10 + (1.5 / v4))
        _ = v7
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        if ((94 < 127) && false) {
            for fi9 := 0; fi9 < 12; fi9++ {
                switch 2147483647 {
                case 9223372036854775807:
                    v5 += (v5 + v5)
                    var v10 string = "\n\t"
                    v4 -= v4
                    _ = v10
                default:
                    v2 += 292
                }
                switch "foo bar" {
                case "\n\t":
                    var v11 int32 = -2147483648
                    _ = v11
                default:
                    var v12 bool = true
                    _ = v12
                }
                var v13 uint16 = 1
                _ = fi9
                _ = v13
            }
        } else {
            v2 -= (v2 & 48)
            var v14 uint = 2
            _ = v14
        }
        _ = fi8
    }
    v15 := GenS2{F0: true, F1: 127, F2: 0}
    for fi16 := 0; fi16 < 2; fi16++ {
        v17 := GenS3{F0: 819, F1: 267, F2: 273}
        var v18 uint32 = 867
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 float32 = 1e10
    switch 42 {
    case 2147483647:
        v2 = 42
    case -271:
        v20 := GenS4{F0: 42, F1: (v5 % (((v5 - v5)) | 1))}
        _ = v20
    case 1:
        v19 += ((0.0 + (v19 * float32(float64(3.14159)))) + 1.5)
    }
    var v21 string = "hello"
    v22 := GenS5{F0: 629, F1: (v5 ^ v5)}
    v23 := GenS6{F0: -2147483648, F1: 42, F2: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
