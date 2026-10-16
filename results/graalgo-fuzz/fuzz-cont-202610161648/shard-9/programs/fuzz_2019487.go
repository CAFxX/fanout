package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 int64; F2 uint64 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 579
    var v3 float64 = float64(0.1)
    var v4 uint8 = 2
    switch "" {
    case "hello":
        for fi5 := 0; fi5 < 10; fi5++ {
            v6 := GenS1{F0: 1}
            var v7 int8 = 0
            v8 := []int{-37, fi5, fi5, (fi5 * fi5)}
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
    case "foo bar":
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 int32 = -400
            v11 := GenS1{F0: 32767}
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    switch "a" {
    case "hello":
        switch "hello" {
        case "hello":
            v12 := GenS2{F0: 18446744073709551615, F1: 0, F2: v2}
            _ = v12
        case "\n\t":
            var v13 float32 = (((0.0 - float32(float64(((1.5 / ((float32(v3) + 1e10) * float32(v3))) + 0.0)))) + float32(0.1)) - (-2.25 / -2.25))
            _ = v13
        default:
            for fi14 := 0; fi14 < 12; fi14++ {
                v2 = (^(v2))
                _ = fi14
            }
        }
    case "":
        fmt.Println(v3, v2)
    default:
        var v15 float32 = -2.25
        _ = v15
    }
    v4 -= v4
    v4 = 1
    v3 += (float64(0.1) / 1.0)
    v16 := GenS1{F0: -244}
    v4 -= (v4 ^ 0)
    switch ("hello" + "hello") {
    case "\n\t":
        var v17 int16 = -594
        _ = v17
    }
    v4 = v4
    var v18 int8 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
