package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 bool }
type GenS2 struct { F0 uint64 }
func genFn0() uint8 {
    var v4 int64 = -9223372036854775808
    v5 := []int{128}
    _ = v4
    _ = v5
    return 213
}
func genFn1(p0 uint16) int64 {
    p0 -= p0
    v6 := []int{32768, 128}
    var v7 uint32 = 2
    fmt.Println(v7, p0)
    _ = p0
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 uint8 = 1
    var v2 bool = (!(32768 != -248))
    v3 := make(map[int]int)
    var v8 uint = 0
    v9 := GenS1{F0: v8, F1: 0, F2: (0.1 <= 1e10)}
    if v2 {
        fmt.Println(v2, v2)
        var v10 float32 = (0.0 + 1.5)
        fmt.Println(v8)
        _ = v10
    } else {
        v11 := GenS1{F0: (^(v8)), F1: 1, F2: v2}
        _ = v11
    }
    v12 := []int{-515, 32768, 2147483647, 3}
    _ = len(v12)
    switch "" {
    case "":
        var v13 uint16 = 149
        _ = v13
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 uint8 = v1
        v15 -= genFn0()
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        v8 += 2
        v17 := v3[(fi16 - (fi16 / (((fi16 << ((0) & 63))) | 1)))]
        _ = fi16
        _ = v17
    }
    v8 = (v8 << (((v8 & v8)) & 63))
    if (genFn0() != v1) {
        fmt.Println(v8)
    } else {
        v18 := v3[42]
        _ = v18
    }
    var v19 uint = 18446744073709551615
    fmt.Println(v1, v2)
    if true {
        v20 := GenS2{F0: 18446744073709551615}
        if v2 {
            var v21 float32 = 3.14159
            _ = v21
        }
        v22 := GenS1{F0: v19, F1: 1, F2: false}
        _ = v20
        _ = v22
    } else {
        v23 := GenS2{F0: 42}
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
