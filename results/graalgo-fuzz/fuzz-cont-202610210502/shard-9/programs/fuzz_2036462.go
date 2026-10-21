package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint64 }
func genFn0() int32 {
    var v7 uint16 = 0
    v7 = (v7 - (v7 + 65535))
    _ = v7
    return 1
}
func genFn1(p0 int8, p1 uint, p2 string) uint16 {
    p0 += -128
    fmt.Println(p2)
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 bool = true
        v10 := GenS1{F0: "a", F1: fi8, F2: 2}
        p1 -= p1
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 string = (p2 + "a")
    p1 = (p1 ^ (^(p1)))
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 211
}

func main() {
    var v1 int16 = 0
    v2 := make(map[int]int)
    v3 := make(map[uint32]string)
    var v4 bool = true
    v5 := make(map[uint32]uint32)
    var v6 uint32 = 960
    var v12 uint64 = 746
    if v4 {
        var v13 uint8 = 42
        var v14 uint32 = 2
        v15 := []int{127, 9223372036854775807, 127}
        _ = v13
        _ = v14
        _ = v15
    } else {
        var v16 uint = 65535
        var v17 int64 = 1
        _ = v16
        _ = v17
    }
    v4 = (!false)
    v6 = (v6 | (v6 / ((v6) | 1)))
    fmt.Println(v12)
    v1 += 32767
    v18 := []int{128, -9223372036854775808, 2147483647}
    var v19 bool = v4
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
