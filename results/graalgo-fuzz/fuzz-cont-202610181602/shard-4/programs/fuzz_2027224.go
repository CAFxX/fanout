package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 bool; F1 bool }
func genFn0(p0 int8, p1 int64, p2 uint64) bool {
    p1 = p1
    var v5 int16 = -351
    p0 += 36
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return true
}
func genFn1(p0 uint64) int16 {
    p0 = (p0 * p0)
    p0 = p0
    var v6 int64 = 0
    if true {
        switch ("hello" + "\n\t") {
        case "\n\t":
            v7 := GenS1{F0: 65535, F1: 2147483648}
            _ = v7
        }
    }
    _ = p0
    _ = v6
    return 1000
}

func main() {
    var v1 int8 = 1
    var v2 int32 = -1
    v3 := make(map[int]uint32)
    var v4 int64 = 92
    if (612 >= 4294967295) {
        fmt.Println(v2, v4)
        v8 := GenS2{F0: ((-(v2)) < (^(v2))), F1: false}
        v4 -= v4
        _ = v8
    }
    var v9 uint64 = 2
    v10 := v3[0]
    for fi11 := 0; fi11 < 1; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            fmt.Println(v2, v9)
            delete(v3, fi12)
            _ = fi12
        }
        _ = fi11
    }
    if (-275 >= 1) {
        if (v10 <= (v10 >> ((18446744073709551615) & 63))) {
            v10 += 2147483647
        }
    } else {
        for fi13 := 0; fi13 < 13; fi13++ {
            _ = len(v3)
            fmt.Println(v9, v2)
            _ = fi13
        }
    }
    switch -79 {
    case -9223372036854775808:
        v4 -= -9223372036854775808
    case 2147483647:
        v14 := []int{127, -9223372036854775808, -1, 961}
        _ = v14
    default:
        v10 = v10
    }
    var v15 uint32 = v10
    v16 := GenS2{F0: false, F1: true}
    v15 -= v15
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
