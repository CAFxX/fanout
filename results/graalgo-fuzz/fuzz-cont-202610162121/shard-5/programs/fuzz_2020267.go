package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int64 }

func main() {
    v1 := []int{}
    v2 := []uint32{2, 4294967295}
    v3 := make(map[int]bool)
    v4 := make(map[uint32]bool)
    v5 := make(map[int]bool)
    var v6 float32 = 3.14159
    v6 = ((v6 - (float32(1e10) + float32(((-2.25 + 1.5) / (0.0 - 1e10))))) - v6)
    v6 = -2.25
    if (0 > 827) {
        v6 = 1.5
    } else {
        delete(v5, -325)
    }
    fmt.Println(v6)
    var v7 float32 = v6
    v6 -= (1.5 - (v6 * v6))
    v6 -= float32(0.1)
    var v8 uint32 = 1
    v7 += 0.1
    var v9 int = -127
    v7 -= v6
    v8 += v8
    v10 := GenS1{F0: v9, F1: 370, F2: -9223372036854775808}
    var v11 uint16 = 0
    var v12 int16 = -249
    switch (v9 * v9) {
    case 530:
        v13 := GenS1{F0: v9, F1: 2, F2: -635}
        _ = v13
    case 2147483647:
        if (!(true && true)) {
            v6 = -2.25
            var v14 uint64 = 0
            v7 -= v7
            _ = v14
        }
    case 409:
        switch "foo bar" {
        case "\n\t":
            var v15 uint8 = 128
            _ = v15
        case "foo bar":
            v8 += (v8 % ((2147483648) | 1))
        case "":
            var v16 uint8 = 127
            _ = v16
        }
    default:
        v3[(v9 + -176)] = ((("" + "") + "") <= "\n\t")
    }
    _ = len(v2)
    var v17 int16 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
