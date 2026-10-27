package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }

func main() {
    var v1 uint64 = 1
    v2 := []string{}
    v3 := make(map[uint32]int)
    v4 := make(map[string]int)
    var v5 float64 = 1e10
    var v6 int64 = -1
    fmt.Println(v1, v1, v1)
    v6 -= v6
    switch 32767 {
    case -9223372036854775808:
        var v7 int64 = (v6 | -9223372036854775808)
        _ = v7
    case -767:
        v6 -= v6
    default:
        fmt.Println(v6, v1, v5)
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 int = (-(fi8))
        _ = fi8
        _ = v9
    }
    fmt.Println(v6, v6, v5)
    var v10 uint = 18
    var v11 int64 = v6
    for fi12 := 0; fi12 < 17; fi12++ {
        _ = len(v4)
        var v13 uint64 = v1
        var v14 int16 = -32768
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 uint64 = v1
    v16 := GenS1{F0: int64(v1), F1: v10}
    var v17 int = 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
