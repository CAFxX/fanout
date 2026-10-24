package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }

func main() {
    var v1 int32 = -207
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]uint32)
    switch 0 {
    case 42:
        v4 := GenS1{F0: 9223372036854775807, F1: 65535}
        _ = v4
    case -1:
        v5 := GenS1{F0: 1, F1: 32767}
        _ = v5
    case -379:
        v6 := v3[2147483648]
        _ = v6
    }
    delete(v2, 338)
    var v7 int32 = (^(v1))
    var v8 float64 = 1e10
    switch "hello" {
    case "a":
        v2[-2147483648] = 2147483648
    }
    fmt.Println(v7, v1, v1)
    v8 -= v8
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 int16 = -643
        _ = fi9
        _ = v10
    }
    fmt.Println(v7, v8, v8)
    var v11 int64 = 1
    for fi12 := 0; fi12 < 12; fi12++ {
        v1 = (v7 / ((v1) | 1))
        fmt.Println(v7, v8)
        v11 += (v11 & (v11 + v11))
        _ = fi12
    }
    v7 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
