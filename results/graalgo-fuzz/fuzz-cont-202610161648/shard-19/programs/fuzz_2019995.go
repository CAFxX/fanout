package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int32; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int; F1 uint16; F2 uint }
func genFn0(p0 int8, p1 int16) int16 {
    v4 := []int{128, 9223372036854775807, 32768}
    var v5 float32 = ((1.5 * 1e10) + (1.5 * 1.5))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (p1 * (p1 & 1))
}

func main() {
    var v1 float64 = 3.14159
    var v2 uint8 = 128
    v3 := []int{-9223372036854775808, 32768, 4294967295, 844}
    v6 := GenS1{F0: (^(v2)), F1: 0}
    v7 := GenS2{F0: -1, F1: 97, F2: 18446744073709551615}
    v2 -= (v2 + v2)
    if ("a" != "foo bar") {
        fmt.Println(v2, v1)
    } else {
        for fi8 := 0; fi8 < 14; fi8++ {
            switch (^(fi8)) {
            case -9223372036854775808:
                v1 = (v1 / 1.0)
            case 216:
                if true {
                    for fi9 := 0; fi9 < 8; fi9++ {
                        v2 = (v2 % ((2) | 1))
                        _ = fi9
                    }
                    var v10 uint16 = 65535
                    switch fi8 {
                    case -2147483648:
                        v11 := GenS2{F0: 0, F1: 18446744073709551615, F2: 456}
                        _ = v11
                    case 32767:
                        v1 += (v1 * float64((3.14159 - float32(v1))))
                    case -737:
                        v1 -= v1
                    default:
                        v12 := GenS3{F0: fi8, F1: v10, F2: 610}
                        _ = v12
                    }
                    _ = v10
                } else {
                    fmt.Println(v1, v1)
                }
            default:
                var v13 int32 = -2147483648
                _ = v13
            }
            _ = fi8
        }
    }
    var v14 uint16 = 255
    if len(v3) > 0 {
        v3[0] = 32768
    }
    v1 = float64(((0.1 - -2.25) / -2.25))
    var v15 float64 = v1
    fmt.Println(v15, v2, v14)
    v16 := GenS2{F0: 2147483647, F1: 42, F2: 0}
    var v17 int32 = -2147483648
    if false {
        var v18 string = ""
        v17 = 42
        v2 = v2
        _ = v18
    }
    var v19 float32 = 1e10
    v1 += (v1 / (0.0 + v1))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
