package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }

func main() {
    var v1 int64 = 42
    var v2 int32 = 0
    var v3 int16 = 1
    var v4 int8 = -1
    var v5 int32 = v2
    var v6 int32 = 2147483647
    v7 := []int{32767, -523}
    v7 = append(v7, -1)
    switch int(v6) {
    case 0:
        if ((-(v2)) <= 979) {
            var v8 uint = 255
            _ = len(v7)
            _ = v8
        }
    case 9223372036854775807:
        fmt.Println(v6)
    }
    var v9 uint8 = 255
    v4 -= (v4 << ((2) & 63))
    for fi10 := 0; fi10 < 11; fi10++ {
        v1 = 1
        _ = fi10
    }
    switch 89 {
    case 1:
        var v11 uint32 = 2147483648
        _ = v11
    case 101:
        v3 = (^(v3))
    case 32767:
        if len(v7) > 0 {
            v7[0] = 786
        }
    default:
        v5 -= 2147483647
    }
    v7 = append(v7, 9223372036854775807)
    if len(v7) > 0 {
        v7[0] = -307
    }
    if (1 >= 0) {
        var v12 float64 = ((0.0 - 1.5) * (float64((((0.0 - 1.5) - float32(-2.25)) * float32(0.0))) - (((1.5 + -2.25) / 1.0) + 0.1)))
        v13 := GenS1{F0: 2147483647, F1: 65535}
        _ = v12
        _ = v13
    }
    fmt.Println(v5, v4)
    v1 += (v1 | v1)
    for fi14 := 0; fi14 < 19; fi14++ {
        v5 -= -1
        switch "a" {
        case "foo bar":
            v5 = 1
        case "\n\t":
            fmt.Println(v3, v1, v9)
        case "a":
            var v15 uint8 = 26
            _ = v15
        default:
            for fi16 := 0; fi16 < 13; fi16++ {
                fmt.Println(fi14, v3)
                fmt.Println(fi16)
                var v17 float64 = 1.5
                _ = fi16
                _ = v17
            }
        }
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
