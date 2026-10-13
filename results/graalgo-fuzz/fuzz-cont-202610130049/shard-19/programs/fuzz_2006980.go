package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int8 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 string) int64 {
    p0 = p0
    switch 1 {
    case 127:
        p0 = p0
    case 490:
        p0 = p0
    default:
        var v7 uint32 = 119
        _ = v7
    }
    _ = p0
    return 42
}

func main() {
    var v1 float32 = (3.14159 / -2.25)
    var v2 int = -1
    var v3 bool = false
    var v4 float32 = float32(1.5)
    var v5 uint64 = 0
    v6 := make(map[uint32]bool)
    if (1 != 127) {
        for fi8 := 0; fi8 < 9; fi8++ {
            fmt.Println(fi8)
            v9 := GenS1{F0: 206, F1: fi8, F2: 0}
            v1 = v1
            _ = fi8
            _ = v9
        }
    }
    fmt.Println(v3)
    v4 += 0.0
    var v10 uint = 42
    for fi11 := 0; fi11 < 15; fi11++ {
        switch v2 {
        case 9223372036854775807:
            v3 = true
        default:
            v12 := GenS1{F0: 2, F1: (v2 << ((v10) & 63)), F2: -128}
            _ = v12
        }
        v13 := GenS1{F0: 0, F1: fi11, F2: 64}
        var v14 int64 = 0
        _ = fi11
        _ = v13
        _ = v14
    }
    var v15 int16 = 42
    v16 := GenS2{F0: 2147483648}
    v1 = v1
    var v17 uint64 = (v5 | 4294967295)
    v17 = 883
    v17 -= v17
    v18 := []int{-2147483648, int(v17), (v2 >> (((v10 / ((v10) | 1))) & 63)), 171}
    if v3 {
        if (!(!(v15 < 1))) {
            fmt.Println(v4, v3)
            v15 = -1
        } else {
            var v19 int16 = v15
            var v20 int64 = 330
            _ = v19
            _ = v20
        }
    }
    var v21 float32 = float32(0.0)
    switch (v2 << (((v10 | 18446744073709551615)) & 63)) {
    case -550:
        fmt.Println(v21, v3, v5)
    case 2147483647:
        var v22 string = "foo bar"
        _ = v22
    case 9223372036854775807:
        v3 = v3
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
