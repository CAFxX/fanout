package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 string }
func genFn0() uint32 {
    var v5 int8 = 97
    for fi6 := 0; fi6 < 12; fi6++ {
        fmt.Println(v5, fi6)
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 string = ("" + ("hello" + "a"))
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    _ = v5
    return 42
}
func genFn1(p0 int16, p1 bool, p2 uint64) uint16 {
    v9 := []int{-2147483648, -109, 0, 903}
    var v10 float64 = ((3.14159 / 1e10) - float64(((1e10 * -2.25) / 0.1)))
    var v11 float32 = float32(0.0)
    v12 := GenS1{F0: (p0 ^ (p0 & (p0 | -157))), F1: (^(p2)), F2: -142}
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return 0
}

func main() {
    var v1 float64 = (1.5 / (0.1 + (1.5 - 3.14159)))
    v2 := []uint32{1, 309, 869}
    var v3 bool = (893 >= 2)
    var v4 float32 = (((1e10 * (0.0 - 0.1)) * -2.25) - (float32(((1e10 / (v1 * (v1 - v1))) / v1)) + float32(-2.25)))
    switch 266 {
    case -1:
        fmt.Println(v4, v1, v4)
    }
    v13 := GenS1{F0: 32767, F1: 768, F2: 1}
    if (2147483647 <= 156) {
        if v3 {
            v14 := GenS1{F0: -1, F1: 2, F2: 1}
            fmt.Println(v3)
            if (-1 < 42) {
                v15 := GenS1{F0: 1, F1: 825, F2: 2147483647}
                _ = v15
            }
            _ = v14
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v4 -= (-2.25 - v4)
    if (1e10 < ((0.1 - v1) + v1)) {
        v3 = ((v3 || (!false)) && v3)
        v16 := GenS1{F0: 307, F1: 810, F2: 0}
        for fi17 := 0; fi17 < 10; fi17++ {
            for fi18 := 0; fi18 < 10; fi18++ {
                v19 := GenS2{F0: 1, F1: 0}
                var v20 string = "a"
                switch fi18 {
                case 128:
                    _ = len(v2)
                case 32768:
                    v21 := GenS3{F0: (v20 + v20)}
                    _ = v21
                case 127:
                    v22 := GenS1{F0: -418, F1: 606, F2: -451}
                    _ = v22
                default:
                    var v23 int64 = 42
                    _ = v23
                }
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = fi17
        }
        _ = v16
    }
    if len(v2) > 0 {
        v2[0] = 330
    }
    var v24 int16 = -1
    var v25 int = 4294967295
    v2 = append(v2, 0)
    v24 -= v24
    var v26 float64 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
