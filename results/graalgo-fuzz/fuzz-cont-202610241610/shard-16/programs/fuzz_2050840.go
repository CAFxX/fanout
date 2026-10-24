package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64 }
func genFn0(p0 bool, p1 uint64, p2 uint) int64 {
    p0 = p0
    v5 := []int{42, 42, 127}
    v5 = append(v5, 2147483647)
    _ = len(v5)
    var v6 uint8 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 9223372036854775807
}
func genFn1(p0 uint8) string {
    p0 = p0
    if true {
        if false {
            if (1 >= -66) {
                fmt.Println(p0)
            }
            switch -394 {
            case 32768:
                p0 += p0
            default:
                p0 -= 254
            }
            v7 := []int{}
            _ = v7
        }
        p0 += 127
    }
    _ = p0
    return ""
}

func main() {
    var v1 float64 = (((0.1 + (float64(((1e10 - (1.5 - 1e10)) - ((1e10 / 0.1) * (-2.25 / (3.14159 * 1.5))))) / 3.14159)) + (0.0 / -2.25)) + (-2.25 + 0.1))
    v2 := make(map[string]string)
    var v3 uint16 = 670
    v4 := []uint32{65535, 4294967295}
    for fi8 := 0; fi8 < 20; fi8++ {
        fmt.Println(v1, fi8)
        _ = fi8
    }
    v3 -= v3
    for fi9 := 0; fi9 < 3; fi9++ {
        v3 += v3
        _ = fi9
    }
    fmt.Println(v1, v3)
    v10 := GenS1{F0: 127, F1: 113}
    v1 = (0.1 / -2.25)
    v1 -= v1
    fmt.Println(v1)
    switch 127 {
    case -9223372036854775808:
        var v11 float32 = (0.1 - ((0.0 + (0.0 / -2.25)) / (1e10 / 1e10)))
        _ = v11
    case 127:
        if (!true) {
            for fi12 := 0; fi12 < 1; fi12++ {
                v3 -= v3
                v13 := GenS1{F0: 127, F1: 0}
                v1 += 0.0
                _ = fi12
                _ = v13
            }
            for fi14 := 0; fi14 < 4; fi14++ {
                var v15 int64 = 9223372036854775807
                _ = fi14
                _ = v15
            }
        }
    case 9223372036854775807:
        if false {
            if len(v4) > 0 {
                v4[0] = 1
            }
        }
    default:
        v1 -= v1
    }
    if (false && true) {
        v16 := GenS2{F0: 9223372036854775807}
        _ = v16
    }
    var v17 uint16 = v3
    v18 := GenS3{F0: ""}
    v3 = (v17 ^ 2)
    v1 += v1
    for fi19 := 0; fi19 < 18; fi19++ {
        v20 := GenS4{F0: -78}
        v21 := GenS5{F0: -1}
        _ = fi19
        _ = v20
        _ = v21
    }
    v3 = v3
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
