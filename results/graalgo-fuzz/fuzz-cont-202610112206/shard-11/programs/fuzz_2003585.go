package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 bool }
func genFn0(p0 int64, p1 int32) uint64 {
    fmt.Println(p0, p0)
    v6 := GenS1{F0: 2147483648}
    var v7 float64 = 0.0
    v8 := []int{9223372036854775807}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1() int {
    var v9 uint = 527
    for fi10 := 0; fi10 < 16; fi10++ {
        switch fi10 {
        case -9223372036854775808:
            var v11 int64 = 1
            _ = v11
        case -962:
            v9 = v9
        default:
            var v12 int64 = 0
            _ = v12
        }
        fmt.Println(v9)
        v13 := GenS1{F0: 4294967295}
        _ = fi10
        _ = v13
    }
    _ = v9
    return 128
}

func main() {
    var v1 int32 = -1
    var v2 float32 = (((0.0 * ((1e10 - 1.5) * -2.25)) + (3.14159 * (1e10 / ((3.14159 / 0.1) / 1.0)))) * (float32(0.1) + ((-2.25 - -2.25) + (0.1 * -2.25))))
    var v3 int = 4294967295
    var v4 uint16 = 164
    var v5 uint8 = 128
    var v14 float32 = v2
    v15 := []int{v3}
    var v16 uint = 42
    switch "foo bar" {
    case "\n\t":
        v17 := GenS1{F0: 2147483647}
        _ = v17
    case "a":
        var v18 int16 = -1
        _ = v18
    case "foo bar":
        fmt.Println(v3, v1, v3)
    default:
        if (0.0 <= -2.25) {
            var v19 float64 = 0.1
            var v20 uint32 = 1
            switch "foo bar" {
            case "a":
                fmt.Println(v14, v16, v4)
            }
            _ = v19
            _ = v20
        }
    }
    var v21 float64 = ((-2.25 + -2.25) - 0.0)
    fmt.Println(v5, v4)
    if true {
        v22 := GenS2{F0: -91, F1: (v1 <= (v1 & (v1 / (((v1 >> (((^(v16))) & 63))) | 1))))}
        fmt.Println(v5)
        _ = v22
    }
    fmt.Println(v16, v21, v4)
    var v23 uint64 = 876
    v14 -= float32(-2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
