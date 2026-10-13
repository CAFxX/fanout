package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
func genFn0(p0 int64, p1 int8) int8 {
    var v6 uint64 = 4294967295
    if ((false || true) && (247 <= 127)) {
        var v7 int = 530
        _ = v7
    } else {
        fmt.Println(v6)
        if true {
            var v8 uint = 93
            p0 += 1
            _ = v8
        } else {
            var v9 int = 2147483647
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 uint16 = 1
        v6 = (^(v6))
        for fi12 := 0; fi12 < 8; fi12++ {
            p0 -= (p0 % ((p0) | 1))
            var v13 uint = 42
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v6
    return (^(p1))
}

func main() {
    v1 := []uint32{843}
    v2 := []string{}
    var v3 uint32 = 367
    var v4 int64 = -9223372036854775808
    v5 := []int{}
    if true {
        var v14 float64 = (1.5 + ((0.0 / 1.0) - 1e10))
        switch -9223372036854775808 {
        case 566:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 127:
            v4 = 757
        }
        var v15 string = ""
        _ = v14
        _ = v15
    } else {
        v3 -= (v3 << ((4294967295) & 63))
        v16 := GenS1{F0: 1, F1: 1}
        _ = v16
    }
    if (false && ((uint64(v3) == 0) && true)) {
        v1 = append(v1, v3)
    }
    var v17 uint = 255
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v18 int = 90
    v17 = v17
    if false {
        var v19 int32 = 0
        _ = v19
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
