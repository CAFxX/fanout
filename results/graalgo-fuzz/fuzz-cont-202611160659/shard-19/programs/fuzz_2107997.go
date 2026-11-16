package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
func genFn0(p0 uint64) int32 {
    for fi4 := 0; fi4 < 13; fi4++ {
        p0 += 1
        fmt.Println(fi4, fi4)
        _ = fi4
    }
    p0 += 0
    p0 -= (^(p0))
    switch -734 {
    case 860:
        var v5 float64 = ((1.5 - 1.5) - float64(1e10))
        _ = v5
    default:
        v6 := GenS1{F0: 869, F1: 211}
        _ = v6
    }
    p0 = p0
    _ = p0
    return -1
}
func genFn1(p0 bool) int {
    v7 := []int{32767, 4294967295, 42, 32767}
    v8 := GenS1{F0: 18446744073709551615, F1: 18446744073709551615}
    _ = p0
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 uint64 = 0
    var v2 bool = true
    var v3 float32 = 1.5
    v9 := GenS1{F0: 4294967295, F1: 0}
    var v10 uint = 1
    var v11 uint = v10
    v12 := []int{-9223372036854775808, -2147483648, 2147483647}
    var v13 int = genFn1(v2)
    for fi14 := 0; fi14 < 13; fi14++ {
        fmt.Println(v13, v2)
        if v2 {
            v13 -= (-(fi14))
        }
        _ = fi14
    }
    v1 -= v1
    var v15 int64 = 0
    var v16 int8 = -1
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v9
}
