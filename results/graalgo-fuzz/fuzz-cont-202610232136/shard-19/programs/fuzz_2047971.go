package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int64; F1 int8; F2 uint16 }
func genFn0(p0 string, p1 uint16, p2 uint) int8 {
    fmt.Println(p1, p0, p0)
    v5 := []int{-1, 4294967295, 2147483647, 306}
    var v6 int = 9223372036854775807
    _ = len(v5)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -128
}

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 464
    var v3 bool = false
    var v4 uint = 4294967295
    var v7 int32 = 2147483647
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := GenS1{F0: (v4 - (v4 + v4)), F1: -1}
        var v10 int32 = (^(v7))
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: v4, F1: -1}
    var v12 float32 = 0.0
    v13 := v1["\n\t"]
    v14 := GenS2{F0: 1, F1: 42, F2: (v2 + (v2 / (((v2 ^ 42)) | 1)))}
    var v15 uint32 = 2147483648
    v15 -= v15
    v16 := GenS2{F0: 540, F1: genFn0(v13, v2, (v4 >> ((uint(v15)) & 63))), F2: (v2 | (v2 / (((v2 | v2)) | 1)))}
    fmt.Println(v13)
    switch "\n\t" {
    case "\n\t":
        for fi17 := 0; fi17 < 13; fi17++ {
            v12 += v12
            _ = fi17
        }
    case "":
        var v18 uint16 = (^(v2))
        _ = v18
    }
    v19 := GenS2{F0: -9223372036854775808, F1: 0, F2: v2}
    if ((^(v2)) > v2) {
        for fi20 := 0; fi20 < 15; fi20++ {
            v13 = v13
            var v21 uint8 = 42
            _ = fi20
            _ = v21
        }
        v3 = (v2 != (^(v2)))
        v4 += (v4 | (v4 % ((v4) | 1)))
    }
    v22 := []int{-9223372036854775808, -486}
    v3 = (18446744073709551615 < 0)
    var v23 uint32 = v15
    v2 += (v2 - (v2 & 65535))
    switch 128 {
    case 1:
        var v24 float64 = (float64(((v12 * -2.25) * v12)) - ((1.5 - (0.1 - -2.25)) - (float64(v12) / (1.5 - 3.14159))))
        _ = v24
    default:
        v7 = -152
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v7
}
