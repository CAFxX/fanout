package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
func genFn0(p0 uint8, p1 uint8) uint32 {
    var v4 int64 = 227
    if true {
        var v5 uint16 = 0
        var v6 int = 42
        fmt.Println(v5, v6)
        _ = v5
        _ = v6
    } else {
        p1 = p0
        if (false && false) {
            var v7 float32 = float32((0.0 / ((0.1 - 0.1) + 1e10)))
            _ = v7
        }
    }
    _ = p0
    _ = p1
    _ = v4
    return 1
}

func main() {
    var v1 int16 = 824
    v2 := make(map[int]uint32)
    v3 := []int{-9223372036854775808}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 -= (-(v1))
    v8 := GenS1{F0: 42, F1: 1}
    var v9 uint32 = 2147483647
    var v10 int32 = 221
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 uint32 = (^(v9))
        _ = fi11
        _ = v12
    }
    var v13 uint16 = 32767
    if true {
        v9 += (^(v9))
    } else {
        if true {
            var v14 int16 = v1
            _ = v14
        }
        var v15 int = -2147483648
        switch (v15 * (v15 % ((v15) | 1))) {
        case 1:
            var v16 uint = 42
            _ = v16
        default:
            var v17 uint8 = 195
            _ = v17
        }
        _ = v15
    }
    var v18 int64 = 1
    v3 = append(v3, 9223372036854775807)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
