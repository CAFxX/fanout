package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint32 }

func main() {
    var v1 float32 = (float32(((1e10 + -2.25) * ((1e10 + float64(0.0)) + 0.1))) - (((1e10 - float32((((1e10 / (3.14159 * (((-2.25 / float64(0.1)) / -2.25) / (0.1 + -2.25)))) / ((1e10 / 1.5) + float64(float32(-2.25)))) / ((1.5 + -2.25) + float64((float32(float64(3.14159)) - 3.14159)))))) / -2.25) - 1.5))
    v2 := []string{}
    var v3 int = 0
    var v4 uint32 = 2147483647
    fmt.Println(v1, v1)
    _ = len(v2)
    fmt.Println(v4)
    if true {
        v1 = v1
        fmt.Println(v1)
        var v5 int32 = 42
        _ = v5
    } else {
        var v6 float64 = (0.0 - (0.0 / float64(v1)))
        _ = v6
    }
    if false {
        var v7 float64 = float64(v1)
        v4 = 1
        _ = v7
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        _ = len(v2)
        switch "a" {
        case "foo bar":
            fmt.Println(v4)
        default:
            var v9 uint64 = 42
            _ = v9
        }
        v10 := GenS1{F0: 0, F1: 923, F2: v4}
        _ = fi8
        _ = v10
    }
    var v11 uint32 = v4
    v12 := GenS1{F0: 9223372036854775807, F1: 793, F2: v4}
    v2 = append(v2, "a")
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
