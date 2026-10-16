package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int }
type GenS2 struct { F0 int64 }
func genFn0() uint32 {
    var v7 uint16 = 0
    v8 := GenS1{F0: false, F1: 42, F2: 1}
    if (true && (16 > -1)) {
        v7 += v7
        fmt.Println(v7)
    } else {
        v9 := GenS1{F0: false, F1: 79, F2: -1}
        var v10 uint32 = uint32(v7)
        v10 -= v10
        _ = v10
        _ = v9
    }
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = 42
    var v3 uint32 = 2147483648
    var v4 int32 = (v2 | -450)
    v5 := make(map[int]bool)
    v6 := make(map[int]bool)
    v2 -= 338
    v11 := []int{-473, 32768, 9223372036854775807, 32768}
    v2 = (v4 / (((^(v2))) | 1))
    v12 := GenS2{F0: -9223372036854775808}
    _ = len(v1)
    v3 -= genFn0()
    v4 = v4
    v2 -= v4
    v2 = v4
    var v13 int8 = 127
    v4 = (v2 ^ v4)
    if (false && (754 == 513)) {
        if ((-2.25 / ((((3.14159 - 1.5) * (((0.0 / (0.0 + 1e10)) / 1.5) - -2.25)) * 3.14159) * 1e10)) != 1e10) {
            var v14 int16 = -1
            var v15 uint32 = v3
            _ = v14
            _ = v15
        }
    }
    v16 := GenS2{F0: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
