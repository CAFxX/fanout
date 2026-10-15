package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 string }

func main() {
    var v1 float64 = (1.5 - float64(((((0.0 * ((float32(float64(3.14159)) + 1e10) / ((1e10 / 1e10) * ((-2.25 * (1e10 * (0.1 / (-2.25 / (-2.25 * (-2.25 * -2.25)))))) - (float32(1e10) + ((-2.25 * (3.14159 - 1.5)) + (float32(0.1) - 1.5))))))) * (0.0 - -2.25)) * ((0.0 * float32((0.1 / 0.1))) - (-2.25 - float32(float64(0.0))))) / (((3.14159 + 1.5) * ((1.5 * (3.14159 * 1.5)) - (-2.25 - (1.5 / 1.5)))) * 1.5))))
    v2 := make(map[int]uint32)
    var v3 string = "foo bar"
    var v4 int64 = -9223372036854775808
    var v5 int64 = (v4 % ((v4) | 1))
    v6 := []int{128, 127}
    if (false || true) {
        v7 := GenS1{F0: v4}
        _ = v7
    }
    switch v3 {
    case "":
        for fi8 := 0; fi8 < 16; fi8++ {
            v1 = (v1 * float64(1.5))
            _ = fi8
        }
    case "\n\t":
        var v9 bool = false
        _ = v9
    default:
        fmt.Println(v5, v1, v5)
    }
    var v10 uint16 = 65535
    if (v1 == (1.5 + 1.5)) {
        v11 := GenS2{F0: 0, F1: 255}
        fmt.Println(v5)
        _ = v11
    } else {
        v10 += v10
    }
    var v12 uint32 = 4294967295
    v13 := GenS3{F0: -445}
    for fi14 := 0; fi14 < 5; fi14++ {
        v15 := GenS4{F0: 42, F1: "a"}
        v16 := GenS2{F0: 1, F1: 65535}
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 int32 = 1
    var v18 uint8 = 128
    var v19 bool = (uint16(v12) >= v10)
    v1 += 0.0
    if len(v6) > 0 {
        v6[0] = -1
    }
    if v19 {
        v5 -= (v4 * (v4 & (^(v5))))
        for fi20 := 0; fi20 < 4; fi20++ {
            v18 = (v18 | v18)
            _ = fi20
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
