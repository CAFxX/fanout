package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int8 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint32) int {
    v4 := GenS1{F0: 229, F1: -9223372036854775808, F2: 42}
    p0 = p0
    v5 := GenS1{F0: 1, F1: 209, F2: -37}
    for fi6 := 0; fi6 < 17; fi6++ {
        for fi7 := 0; fi7 < 11; fi7++ {
            v8 := GenS1{F0: 0, F1: -643, F2: 118}
            _ = fi7
            _ = v8
        }
        var v9 int16 = 249
        _ = fi6
        _ = v9
    }
    v10 := []int{-962, -9223372036854775808, -2147483648, 42}
    _ = p0
    _ = v10
    _ = v4
    _ = v5
    return 2147483647
}
func genFn1(p0 bool) uint {
    v11 := GenS1{F0: 127, F1: 9223372036854775807, F2: 79}
    v12 := []int{83, 519, 9223372036854775807, 435}
    v12 = append(v12, 773)
    var v13 int16 = 28
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return 59
}

func main() {
    var v1 uint64 = 2
    var v2 uint16 = 2
    v3 := make(map[uint32]uint32)
    v2 = 2
    fmt.Println(v2)
    for fi14 := 0; fi14 < 5; fi14++ {
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 uint32 = 993
            fmt.Println(fi14)
            switch genFn0(0) {
            case -363:
                v17 := GenS2{F0: v2}
                _ = v17
            case 128:
                switch genFn0((v16 & (v16 | v16))) {
                case 128:
                    var v18 float64 = float64((float32(1e10) + 1e10))
                    _ = v18
                default:
                    v16 += (^(v16))
                }
            case 758:
                fmt.Println(v2, v16)
            }
            v19 := GenS2{F0: (v2 | v2)}
            var v20 float32 = 1e10
            _ = fi15
            _ = v16
            _ = v19
            _ = v20
        }
        v1 += (v1 / (((v1 * v1)) | 1))
        _ = fi14
    }
    v2 = (v2 % (((v2 >> ((genFn1(true)) & 63))) | 1))
    v21 := []int{-9223372036854775808}
    v2 = (^(v2))
    var v22 int16 = 1000
    fmt.Println(v2, v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
