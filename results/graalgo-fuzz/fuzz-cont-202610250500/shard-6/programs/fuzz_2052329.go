package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint64 }

func main() {
    var v1 uint32 = 2
    v2 := []int{}
    var v3 float32 = (-2.25 * ((3.14159 + float32(float64(0.0))) * (-2.25 * 0.1)))
    v4 := []int{-249}
    var v5 bool = (0.0 <= (float64((-2.25 / v3)) + (0.0 + 0.1)))
    var v6 float64 = 1e10
    fmt.Println(v6, v6)
    if len(v4) > 0 {
        v4[0] = -1
    }
    v6 = v6
    v5 = v5
    var v7 int8 = 78
    if v5 {
        for fi8 := 0; fi8 < 15; fi8++ {
            fmt.Println(v7, v1)
            _ = fi8
        }
    } else {
        fmt.Println(v1)
        var v9 float64 = (0.1 * (v6 * (v6 * (v6 / v6))))
        fmt.Println(v1, v1, v6)
        _ = v9
    }
    v5 = true
    switch "\n\t" {
    case "foo bar":
        for fi10 := 0; fi10 < 6; fi10++ {
            v11 := GenS1{F0: 322, F1: -9223372036854775808, F2: 4294967295}
            _ = fi10
            _ = v11
        }
    case "hello":
        var v12 int8 = -62
        _ = v12
    default:
        var v13 uint = 42
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
