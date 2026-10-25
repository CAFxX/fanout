package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint16; F1 uint8 }
func genFn0() string {
    switch "\n\t" {
    case "\n\t":
        var v4 uint32 = 0
        _ = v4
    }
    var v5 uint = 786
    _ = v5
    return ""
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 0
    var v3 uint64 = 18446744073709551615
    switch -493 {
    case 2147483647:
        if (128 > 42) {
            switch "\n\t" {
            case "\n\t":
                v2 += v2
            case "a":
                v6 := []int{}
                _ = v6
            case "hello":
                var v7 float64 = 0.0
                _ = v7
            }
            if (0.0 < 1e10) {
                v3 -= v3
                switch 942 {
                case 9223372036854775807:
                    for fi8 := 0; fi8 < 12; fi8++ {
                        var v9 uint32 = 2147483647
                        _ = fi8
                        _ = v9
                    }
                case -998:
                    var v10 bool = true
                    _ = v10
                }
            } else {
                v11 := GenS1{F0: 1, F1: 522, F2: 69}
                _ = v11
            }
        } else {
            fmt.Println(v3)
        }
    default:
        v12 := GenS1{F0: 1, F1: v3, F2: 1}
        _ = v12
    }
    v3 = v3
    v3 = 710
    v13 := []int{4294967295, -1, -535, 32767}
    v2 += -979
    v14 := GenS1{F0: 127, F1: v3, F2: -5}
    v3 -= (^(v3))
    v15 := GenS1{F0: 1, F1: (v3 / ((18446744073709551615) | 1)), F2: -108}
    v16 := GenS1{F0: 0, F1: (v3 % (((v3 % ((475) | 1))) | 1)), F2: 1}
    var v17 int64 = 0
    if true {
        _ = len(v13)
    } else {
        v18 := GenS1{F0: 42, F1: v3, F2: -107}
        var v19 bool = (!false)
        v13 = append(v13, 127)
        _ = v18
        _ = v19
    }
    v20 := GenS2{F0: 40, F1: 137}
    var v21 int64 = v17
    v21 = int64(v2)
    if (!true) {
        for fi22 := 0; fi22 < 17; fi22++ {
            var v23 uint64 = v3
            _ = fi22
            _ = v23
        }
        if (true || (true || (1.5 >= ((float64((((3.14159 * float32(0.0)) / float32(1e10)) - float32(float64((3.14159 / 1e10))))) - (1.5 + (((3.14159 * 0.0) * -2.25) + 3.14159))) / (((float64((0.0 - 1e10)) + (0.1 + 0.1)) * (3.14159 * 0.0)) + 3.14159))))) {
            v17 += 767
            var v24 int32 = v2
            var v25 uint64 = 1
            _ = v24
            _ = v25
        } else {
            v17 = -860
            fmt.Println(v2, v21)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
