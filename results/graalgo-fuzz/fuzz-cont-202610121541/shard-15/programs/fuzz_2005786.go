package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int8; F1 bool; F2 uint16 }
type GenS3 struct { F0 int8; F1 uint; F2 uint }
type GenS4 struct { F0 uint8; F1 uint; F2 uint8 }
func genFn0(p0 bool, p1 uint, p2 string) int16 {
    var v7 uint8 = 245
    var v8 int = 131
    v9 := []int{v8, v8, (v8 & v8), (-(v8))}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1() int {
    v10 := GenS1{F0: genFn0((503 <= 2), 433, "\n\t"), F1: (float32(1e10) == (0.0 * (float32((float64((float32(3.14159) + 1.5)) * 3.14159)) - 0.1)))}
    if true {
        if true {
            var v11 uint8 = 2
            v11 = v11
            if true {
                v11 -= v11
            } else {
                fmt.Println(v11)
            }
            _ = v11
        }
    }
    var v12 float64 = float64(((((0.0 / 1.5) - (-2.25 + float32((((((1e10 * 1.5) - 1e10) + (0.1 * float64((1.5 - 1.5)))) + 1e10) * 1e10)))) * float32(3.14159)) / float32(-2.25)))
    _ = v10
    _ = v12
    return 127
}

func main() {
    v1 := []int{32767}
    var v2 int8 = 127
    var v3 uint16 = 0
    v4 := []int{0, -1, 837, 1}
    var v5 uint8 = 255
    v6 := make(map[string]string)
    var v13 uint8 = v5
    var v14 int64 = -9223372036854775808
    v15 := GenS2{F0: v2, F1: (2147483648 == 2147483648), F2: uint16(v5)}
    v16 := GenS2{F0: v2, F1: false, F2: v3}
    switch "\n\t" {
    case "foo bar":
        v17 := GenS3{F0: v2, F1: uint(v3), F2: 255}
        _ = v17
    case "":
        fmt.Println(v3, v3, v2)
    default:
        fmt.Println(v3)
    }
    var v18 int = 31
    v2 += (v2 - (v2 / ((v2) | 1)))
    v2 -= v2
    v13 -= v13
    v18 = (v18 - (^(v18)))
    if false {
        var v19 uint = 509
        _ = v19
    }
    var v20 int8 = 127
    for fi21 := 0; fi21 < 20; fi21++ {
        switch (fi21 ^ fi21) {
        case 0:
            v22 := GenS4{F0: 2, F1: 120, F2: v5}
            _ = v22
        default:
            var v23 uint = 870
            _ = v23
        }
        if len(v1) > 0 {
            v1[0] = (fi21 / ((genFn1()) | 1))
        }
        _ = fi21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
