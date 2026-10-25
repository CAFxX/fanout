package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 uint32 = 0
    v2 := []uint32{}
    var v3 uint32 = 2147483648
    var v4 string = "foo bar"
    v5 := make(map[uint32]bool)
    v6 := GenS1{F0: 9223372036854775807, F1: -68}
    switch (v4 + v4) {
    case "":
        var v7 uint16 = 0
        _ = v7
    default:
        fmt.Println(v4, v1, v3)
    }
    v8 := GenS1{F0: 32768, F1: 32768}
    v3 -= 125
    switch "" {
    case "foo bar":
        v9 := GenS1{F0: 0, F1: 9223372036854775807}
        _ = v9
    default:
        var v10 int8 = -128
        _ = v10
    }
    switch 32768 {
    case -1:
        v11 := GenS2{F0: 2}
        _ = v11
    case 1:
        if false {
            v12 := GenS1{F0: 267, F1: -2147483648}
            var v13 float32 = (-2.25 * float32(1.5))
            _ = v12
            _ = v13
        } else {
            v3 -= (^(v3))
            for fi14 := 0; fi14 < 11; fi14++ {
                fmt.Println(v1)
                v15 := GenS1{F0: (fi14 % ((fi14) | 1)), F1: 9223372036854775807}
                _ = fi14
                _ = v15
            }
        }
    case 0:
        v1 = v3
    }
    v16 := GenS2{F0: 2}
    v4 = v4
    fmt.Println(v4, v4, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
