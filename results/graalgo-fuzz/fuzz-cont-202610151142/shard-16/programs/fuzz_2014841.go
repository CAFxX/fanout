package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]bool)
    var v3 bool = true
    var v4 float64 = 0.0
    v5 := make(map[int]uint32)
    v3 = v3
    switch "hello" {
    case "":
        var v6 uint32 = 4294967295
        _ = v6
    case "\n\t":
        var v7 uint = 148
        _ = v7
    case "foo bar":
        v3 = v3
    default:
        v3 = true
    }
    v8 := []int{163, -9223372036854775808, 288, 0}
    v3 = (-2147483648 > 2147483647)
    var v9 float64 = ((v4 * (((float64(0.0) - (float64((1e10 * -2.25)) / 1e10)) + -2.25) - v4)) + v4)
    for fi10 := 0; fi10 < 5; fi10++ {
        switch 4294967295 {
        case 128:
            v11 := GenS1{F0: 694, F1: 42}
            _ = v11
        case 442:
            var v12 uint16 = 65535
            _ = v12
        default:
            var v13 int16 = 0
            _ = v13
        }
        v14 := GenS2{F0: ((0.0 >= (1e10 / 1.5)) || (4294967295 > 1))}
        var v15 float64 = (((v4 / -2.25) / v4) / 1e10)
        _ = fi10
        _ = v14
        _ = v15
    }
    if (0 > 1) {
        if len(v8) > 0 {
            v8[0] = -426
        }
        v16 := GenS3{F0: 1}
        v17 := GenS2{F0: (42 <= 42)}
        _ = v16
        _ = v17
    }
    var v18 int8 = 16
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
