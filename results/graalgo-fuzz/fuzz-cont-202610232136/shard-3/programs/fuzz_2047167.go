package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []int{687, 1, 2147483647}
    var v3 uint = 541
    v4 := []uint32{65535, 624}
    v5 := make(map[uint32]int)
    var v6 float64 = ((0.0 / (((float64(0.1) / 1.0) - -2.25) / (3.14159 - 0.1))) - -2.25)
    switch "\n\t" {
    case "hello":
        var v7 int64 = v1
        _ = v7
    case "foo bar":
        var v8 int = 128
        _ = v8
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        v6 = float64(3.14159)
        _ = fi9
    }
    v1 = v1
    var v10 uint8 = 42
    fmt.Println(v3, v10)
    v5[54] = 4294967295
    var v11 int = 42
    for fi12 := 0; fi12 < 14; fi12++ {
        v3 -= v3
        _ = fi12
    }
    v10 += (v10 / ((v10) | 1))
    var v13 float32 = (-2.25 * -2.25)
    switch (v11 * 840) {
    case -798:
        fmt.Println(v10, v11)
    default:
        v14 := GenS1{F0: true, F1: 18446744073709551615}
        _ = v14
    }
    fmt.Println(v10)
    v4 = append(v4, 2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
