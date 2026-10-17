package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
func genFn0() int32 {
    if (32767 == 255) {
        var v6 bool = false
        v6 = (255 != 131)
        _ = v6
    }
    var v7 float32 = ((1e10 - float32(float64((float32(3.14159) * (0.1 * 0.0))))) + 1e10)
    fmt.Println(v7)
    _ = v7
    return -2147483648
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 823
    v3 := make(map[int]int)
    v4 := make(map[int]uint32)
    var v5 int16 = 0
    fmt.Println(v5, v2)
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 int64 = int64(fi8)
        _ = fi8
        _ = v9
    }
    var v10 uint64 = uint64(v2)
    var v11 int64 = -1
    v2 += 311
    fmt.Println(v10, v11)
    v12 := []int{38, 32767, 0}
    var v13 uint = (^(v2))
    switch 42 {
    case 559:
        v13 = v13
    case -9223372036854775808:
        v5 += v5
    case -1:
        fmt.Println(v11)
    }
    v13 = (v13 & v13)
    if ((true || false) || true) {
        _ = len(v4)
        if len(v12) > 0 {
            v12[0] = 127
        }
    }
    v14 := GenS1{F0: (v2 % (((v2 ^ (v2 << ((0) & 63)))) | 1)), F1: 32768}
    if false {
        v13 = v13
        v5 -= (v5 / ((-1) | 1))
    }
    switch -693 {
    case 9223372036854775807:
        v5 = (v5 / (((v5 ^ (v5 << (((v13 % ((v13) | 1))) & 63)))) | 1))
    default:
        for fi15 := 0; fi15 < 17; fi15++ {
            delete(v3, -9223372036854775808)
            var v16 int16 = (^(v5))
            fmt.Println(v5)
            _ = fi15
            _ = v16
        }
    }
    var v17 float64 = (3.14159 + 1e10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
