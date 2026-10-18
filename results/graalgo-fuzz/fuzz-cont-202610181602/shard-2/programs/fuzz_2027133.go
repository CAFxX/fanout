package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 bool }

func main() {
    var v1 uint16 = 2
    v2 := make(map[uint32]int)
    v3 := []uint32{0, 1}
    v4 := make(map[int]int)
    v5 := make(map[string]int)
    v6 := v4[-9223372036854775808]
    fmt.Println(v1, v1)
    fmt.Println(v1, v1)
    v1 -= (v1 | (^(v1)))
    if true {
        fmt.Println(v6, v1)
        v1 += 42
        v1 = (v1 + (v1 | 42))
    } else {
        v7 := GenS1{F0: 0, F1: 18446744073709551615, F2: false}
        v1 = (v1 / ((v1) | 1))
        _ = v7
    }
    if (false && (32767 <= v6)) {
        v6 = v6
    }
    v3 = append(v3, 113)
    _ = len(v5)
    var v8 int16 = 1000
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := GenS1{F0: 0, F1: 1, F2: ((fi9 >> ((18446744073709551615) & 63)) >= (v6 & (v6 * int(v8))))}
        fmt.Println(fi9)
        v5["a"] = (fi9 << ((862) & 63))
        _ = fi9
        _ = v10
    }
    var v11 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
