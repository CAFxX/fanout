package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int32 }
type GenS2 struct { F0 uint16; F1 int16; F2 uint8 }
func genFn0(p0 int8, p1 bool, p2 uint8) uint32 {
    if p1 {
        p1 = (p1 || p1)
    }
    v4 := GenS1{F0: 498, F1: (p0 ^ (p0 ^ (p0 & (p0 | (p0 >> ((1) & 63)))))), F2: 909}
    fmt.Println(p0, p2, p1)
    v5 := []int{42}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 uint16 = 430
    var v2 int32 = -1
    v3 := make(map[int]int)
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 int64 = 1
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: -380, F1: -1, F2: v2}
    fmt.Println(v2)
    var v9 int = 32767
    switch (-(v9)) {
    case -1:
        fmt.Println(v2, v9, v1)
    }
    var v10 float64 = 0.0
    fmt.Println(v2, v10)
    fmt.Println(v10)
    fmt.Println(v1)
    v1 -= v1
    v11 := GenS2{F0: (v1 << ((65535) & 63)), F1: 32767, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
