package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }

func main() {
    v1 := make(map[int]int)
    v2 := []int{-235, -1, -2147483648}
    var v3 uint64 = 42
    var v4 int64 = 9223372036854775807
    var v5 uint64 = (v3 >> ((4294967295) & 63))
    v5 = v5
    if ((v3 < v5) && ((^(v3)) != 425)) {
        v6 := GenS1{F0: (^(v3)), F1: v4}
        _ = v6
    } else {
        var v7 int64 = 9223372036854775807
        _ = v7
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch 187 {
    case 32768:
        for fi8 := 0; fi8 < 9; fi8++ {
            v5 = (v3 % ((v5) | 1))
            v2 = append(v2, fi8)
            _ = fi8
        }
    }
    var v9 int64 = (v4 + 9223372036854775807)
    v3 = v3
    fmt.Println(v5)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v10 := GenS1{F0: v5, F1: (v4 - v9)}
    v9 = (^(v4))
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 uint32 = 922
        var v13 int8 = -1
        _ = fi11
        _ = v12
        _ = v13
    }
    fmt.Println(v3, v4)
    var v14 int16 = 1000
    if true {
        var v15 int32 = 922
        fmt.Println(v15, v9, v15)
        _ = v15
    }
    v16 := v1[-1]
    var v17 int64 = (-(v9))
    var v18 int16 = (v14 * (v14 & 42))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
