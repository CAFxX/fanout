package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 int16; F1 int16 }
func genFn0(p0 uint, p1 int, p2 int8) bool {
    for fi6 := 0; fi6 < 14; fi6++ {
        if (p0 <= 42) {
            v7 := GenS1{F0: 491, F1: 32767}
            if (false && ((p2 - (p2 | (p2 ^ p2))) == (p2 / ((p2) | 1)))) {
                if ((-2.25 / 1e10) == ((1.5 + ((3.14159 * (1.5 - 1e10)) * 0.1)) / 1e10)) {
                    v8 := GenS1{F0: 2, F1: 1}
                    _ = v8
                } else {
                    var v9 bool = ((p0 / ((566) | 1)) != 42)
                    _ = v9
                }
            }
            switch (fi6 / ((-1) | 1)) {
            case -251:
                fmt.Println(p2, p0, p1)
            case 127:
                var v10 uint64 = 1
                _ = v10
            default:
                p2 -= -45
            }
            _ = v7
        }
        _ = fi6
    }
    if ((p1 ^ p1) <= (p1 ^ (p1 * (p1 * (p1 * (p1 | p1)))))) {
        var v11 uint64 = 799
        p1 -= p1
        _ = v11
    }
    v12 := []int{(p1 % ((p1) | 1)), (p1 * p1), p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return (false && false)
}

func main() {
    v1 := []uint32{2}
    var v2 uint16 = 1
    var v3 float32 = 1.5
    var v4 float32 = 1e10
    v5 := make(map[uint32]bool)
    v1 = append(v1, 4294967295)
    fmt.Println(v3, v3, v4)
    v2 += (v2 | v2)
    if (true || true) {
        fmt.Println(v2)
        v4 = v4
        v13 := GenS2{F0: 0, F1: -2147483648}
        _ = v13
    }
    v14 := GenS3{F0: 42, F1: 660}
    if true {
        v3 = v3
        v3 = v4
        v15 := GenS4{F0: 42, F1: 32}
        _ = v15
    }
    v3 += v3
    var v16 int32 = 526
    v5[4294967295] = ((0 < 18446744073709551615) && false)
    fmt.Println(v3)
    fmt.Println(v4, v2)
    fmt.Println(v16)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
