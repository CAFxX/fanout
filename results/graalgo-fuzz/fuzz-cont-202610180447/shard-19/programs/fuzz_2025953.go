package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 bool }
func genFn0() int8 {
    v4 := []int{}
    var v5 float32 = ((((3.14159 - 0.0) * ((((3.14159 - (-2.25 / float32(float64(-2.25)))) - (float32(1.5) / 1e10)) + (0.1 + 0.0)) + 0.0)) + float32((1.5 / 0.1))) + 0.1)
    _ = v4
    _ = v5
    return 85
}
func genFn1(p0 uint16, p1 int16) int16 {
    if (false && ((3.14159 * 0.0) > -2.25)) {
        switch 0 {
        case 4294967295:
            switch "a" {
            case "a":
                v6 := GenS1{F0: -852}
                _ = v6
            default:
                v7 := []int{117, 42}
                _ = v7
            }
        }
        var v8 int = 32768
        _ = v8
    }
    var v9 uint64 = 18446744073709551615
    v10 := []int{32768, -59, -368}
    switch -9223372036854775808 {
    case -9223372036854775808:
        var v11 uint = 65535
        _ = v11
    default:
        if (470 < 4294967295) {
            switch "\n\t" {
            case "hello":
                p1 += (p1 * p1)
            case "a":
                p0 -= p0
            case "":
                var v12 uint8 = 0
                _ = v12
            default:
                v10 = append(v10, 32768)
            }
        } else {
            p0 = (p0 | 173)
        }
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return (p1 >> ((1) & 63))
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[string]bool)
    v3 := []int{127, 32768, -231}
    var v13 uint16 = 1
    v14 := GenS2{F0: 128, F1: true}
    _ = len(v2)
    switch "a" {
    case "foo bar":
        fmt.Println(v13, v1)
    case "":
        var v15 int64 = 42
        _ = v15
    }
    var v16 int64 = -9223372036854775808
    _ = len(v2)
    v16 -= v16
    var v17 int = 694
    fmt.Println(v16, v16)
    v1 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
