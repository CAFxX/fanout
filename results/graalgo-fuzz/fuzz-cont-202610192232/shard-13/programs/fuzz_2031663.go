package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 uint8; F2 uint64 }
func genFn0(p0 uint64, p1 uint32) int32 {
    v5 := []int{128, -2147483648}
    switch 654 {
    case 778:
        var v6 int32 = -860
        _ = v6
    }
    p1 = (p1 - 2)
    v5 = append(v5, -9223372036854775808)
    var v7 int = 0
    _ = p0
    _ = p1
    _ = v5
    _ = v7
    return -2147483648
}

func main() {
    var v1 int32 = 191
    v2 := make(map[int]bool)
    var v3 int8 = 106
    v4 := []uint32{42, 0, 4294967295, 154}
    v1 = (-(v1))
    for fi8 := 0; fi8 < 20; fi8++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := GenS1{F0: 32767}
            v1 = (v1 ^ (v1 & 42))
            _ = fi9
            _ = v10
        }
        if true {
            var v11 int8 = v3
            v1 = genFn0(4294967295, 389)
            v12 := GenS1{F0: 1}
            _ = v11
            _ = v12
        }
        _ = fi8
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        v1 = v1
        var v13 int8 = int8(v1)
        var v14 int = 32768
        _ = v13
        _ = v14
    }
    if false {
        v3 += v3
    }
    var v15 int = 1
    var v16 int8 = v3
    v17 := GenS2{F0: 9223372036854775807, F1: 194, F2: 18446744073709551615}
    v15 += (v15 | (v15 ^ v15))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
