package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 524
    var v3 uint16 = 32767
    var v4 int64 = (v2 / ((v2) | 1))
    var v5 float32 = float32(((0.1 + 1.5) + 1e10))
    var v6 uint32 = 1
    v7 := GenS1{F0: 0, F1: (true && (14 > 1))}
    fmt.Println(v2)
    for fi8 := 0; fi8 < 10; fi8++ {
        fmt.Println(v6, v6, fi8)
        for fi9 := 0; fi9 < 13; fi9++ {
            delete(v1, fi9)
            _ = fi9
        }
        var v10 float64 = -2.25
        _ = fi8
        _ = v10
    }
    v4 -= (v2 / ((-9223372036854775808) | 1))
    v5 += 1e10
    var v11 bool = (!(50 == 1))
    var v12 int8 = 47
    v3 = v3
    fmt.Println(v6, v5, v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
