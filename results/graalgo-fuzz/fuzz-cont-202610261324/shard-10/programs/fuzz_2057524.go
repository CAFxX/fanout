package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 bool; F2 int32 }
func genFn0(p0 int64, p1 uint8) uint8 {
    p0 = 42
    p1 = (p1 / ((p1) | 1))
    p0 = (p0 >> ((1) & 63))
    p1 -= p1
    _ = p0
    _ = p1
    return p1
}

func main() {
    v1 := []uint32{4294967295, 42}
    v2 := make(map[int]int)
    v3 := make(map[uint32]bool)
    var v4 int = 127
    v4 -= 0
    v5 := GenS1{F0: -903}
    if (595 == 42) {
        fmt.Println(v4)
        v2[-9223372036854775808] = v4
    }
    for fi6 := 0; fi6 < 1; fi6++ {
        fmt.Println(fi6)
        v7 := GenS1{F0: 759}
        _ = fi6
        _ = v7
    }
    v8 := GenS2{F0: false, F1: true, F2: -2147483648}
    if len(v1) > 0 {
        v1[0] = 2
    }
    v4 -= v4
    v4 += (v4 | (^(v4)))
    fmt.Println(v4)
    fmt.Println(v4)
    v9 := GenS1{F0: -848}
    v4 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
