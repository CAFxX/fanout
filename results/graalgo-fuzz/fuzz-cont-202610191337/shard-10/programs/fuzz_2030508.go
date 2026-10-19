package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint16 }

func main() {
    var v1 float64 = (3.14159 * 0.0)
    var v2 float32 = 1.5
    v3 := make(map[int]uint32)
    delete(v3, -2147483648)
    v1 -= (3.14159 + 1.5)
    v1 = 0.0
    var v4 uint16 = 65535
    v1 = v1
    v3[-9223372036854775808] = 2147483647
    for fi5 := 0; fi5 < 17; fi5++ {
        v1 += v1
        _ = fi5
    }
    switch "a" {
    case "a":
        if false {
            var v6 uint8 = 139
            fmt.Println(v4, v6, v2)
            fmt.Println(v4, v6, v2)
            _ = v6
        }
    case "\n\t":
        var v7 int = 0
        _ = v7
    }
    v4 = (^(v4))
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS1{F0: 1000, F1: 42, F2: v4}
        v2 += ((v2 / v2) * v2)
        var v10 int64 = -9223372036854775808
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 int8 = 38
    var v12 int64 = -130
    var v13 float32 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
