package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }

func main() {
    var v1 bool = false
    var v2 int = 42
    var v3 uint32 = 2147483647
    v4 := make(map[string]bool)
    var v5 int64 = -1
    v6 := make(map[int]int)
    var v7 float32 = ((1e10 * 0.1) / float32(3.14159))
    for fi8 := 0; fi8 < 9; fi8++ {
        v7 -= 1e10
        if v1 {
            v2 = fi8
        }
        _ = fi8
    }
    var v9 int = -9223372036854775808
    fmt.Println(v9)
    var v10 int = -556
    var v11 uint16 = 1
    v12 := GenS1{F0: 2147483647, F1: v1}
    for fi13 := 0; fi13 < 18; fi13++ {
        for fi14 := 0; fi14 < 11; fi14++ {
            if (v1 || (124 > 255)) {
                fmt.Println(v1, v5)
                v9 = (v2 | (v2 * fi14))
                v5 = v5
            } else {
                var v15 float32 = v7
                v11 = (v11 * (v11 % ((0) | 1)))
                _ = v15
            }
            _ = fi14
        }
        _ = fi13
    }
    v16 := GenS1{F0: 928, F1: (248 > 65535)}
    switch (v9 ^ v2) {
    case 127:
        v7 = v7
    default:
        var v17 string = ("foo bar" + "\n\t")
        _ = v17
    }
    var v18 float64 = (((1e10 * (3.14159 * 3.14159)) + (1e10 / 0.1)) * 1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
