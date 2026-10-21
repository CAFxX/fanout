package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
func genFn0(p0 int, p1 int) int {
    var v6 int8 = -128
    if true {
        p0 -= (p1 / ((p1) | 1))
        if false {
            v7 := []int{}
            _ = v7
        }
        for fi8 := 0; fi8 < 3; fi8++ {
            v6 = (v6 | v6)
            _ = fi8
        }
    } else {
        v9 := GenS1{F0: 1, F1: -818}
        var v10 int32 = 2147483647
        p0 = (p1 ^ p0)
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    return (p0 >> ((255) & 63))
}

func main() {
    var v1 uint = 4294967295
    var v2 int32 = -117
    var v3 bool = false
    v4 := make(map[uint32]int)
    var v5 bool = v3
    v1 += (^(v1))
    v11 := v4[812]
    if (1 != 1) {
        switch v11 {
        case 32768:
            v12 := []int{9223372036854775807, (v11 / (((-(v11))) | 1))}
            _ = v12
        case -2147483648:
            fmt.Println(v3)
        default:
            var v13 float64 = float64(1.5)
            _ = v13
        }
    } else {
        v2 = v2
        v3 = false
    }
    v14 := []int{(v11 << ((42) & 63))}
    var v15 int32 = v2
    var v16 uint32 = 0
    v15 += 0
    var v17 int16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
