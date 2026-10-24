package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint32 }
func genFn0() uint {
    var v5 int32 = 0
    v5 = v5
    switch -730 {
    case 1:
        v6 := []int{-37}
        _ = v6
    case 477:
        v7 := GenS1{F0: -9223372036854775808, F1: 1, F2: 2}
        _ = v7
    case 32768:
        v5 = (v5 / (((v5 << ((380) & 63))) | 1))
    default:
        switch "a" {
        case "\n\t":
            fmt.Println(v5)
        case "hello":
            v5 = (v5 ^ (v5 & v5))
        }
    }
    v5 -= 2147483647
    for fi8 := 0; fi8 < 4; fi8++ {
        fmt.Println(fi8, fi8)
        var v9 uint32 = 0
        v5 += (-(v5))
        _ = fi8
        _ = v9
    }
    _ = v5
    return 4294967295
}

func main() {
    var v1 float32 = 0.1
    v2 := []string{"foo bar"}
    v3 := make(map[string]bool)
    var v4 uint8 = 1
    fmt.Println(v4, v1)
    v4 += v4
    var v10 bool = false
    var v11 int64 = 42
    v11 -= int64(v4)
    v12 := GenS1{F0: (^(v11)), F1: 42, F2: 250}
    _ = len(v2)
    delete(v3, "hello")
    v13 := GenS1{F0: -722, F1: -9223372036854775808, F2: 178}
    if len(v2) > 0 {
        v2[0] = "a"
    }
    v14 := GenS1{F0: (-(v11)), F1: int(v4), F2: 42}
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 float64 = 3.14159
        if (!v10) {
            switch 0 {
            case -498:
                var v17 float32 = v1
                _ = v17
            }
            var v18 float64 = float64((v1 + ((v1 - 1e10) / 1.5)))
            _ = v18
        } else {
            v10 = true
        }
        _ = fi15
        _ = v16
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
