package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int64, p1 int) uint16 {
    fmt.Println(p1, p0)
    p0 -= (-(p0))
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint16 = 118
    var v2 uint16 = (^(v1))
    var v3 uint32 = 2
    v4 := make(map[int]uint32)
    var v5 int8 = 70
    v6 := GenS1{F0: v5}
    v3 = v3
    var v7 int = -9223372036854775808
    for fi8 := 0; fi8 < 6; fi8++ {
        _ = len(v4)
        fmt.Println(v5)
        var v9 uint16 = (^(v1))
        _ = fi8
        _ = v9
    }
    v5 = v5
    v5 = (v5 ^ v5)
    var v10 uint = 255
    var v11 int8 = (v5 | (v5 / ((0) | 1)))
    v7 = (v7 ^ (v7 >> ((uint(v3)) & 63)))
    v5 = v5
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
