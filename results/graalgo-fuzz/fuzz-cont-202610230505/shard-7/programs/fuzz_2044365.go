package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 bool }
func genFn0(p0 bool, p1 uint) int64 {
    p1 = p1
    p1 -= p1
    fmt.Println(p1)
    v6 := GenS1{F0: -9223372036854775808, F1: 32768}
    _ = p0
    _ = p1
    _ = v6
    return 0
}
func genFn1() int {
    var v7 float64 = 3.14159
    for fi8 := 0; fi8 < 20; fi8++ {
        if ((3.14159 >= (v7 - 1.5)) || (!(!true))) {
            v7 = v7
        }
        _ = fi8
    }
    v7 -= float64(0.0)
    _ = v7
    return 128
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 678
    var v3 uint32 = 0
    v4 := []int{-1, 9223372036854775807}
    var v5 int16 = -32768
    v9 := v1[940]
    v4 = append(v4, -9223372036854775808)
    v3 = (v3 & (v3 / ((uint32(v2)) | 1)))
    var v10 int16 = (v5 >> ((1) & 63))
    v11 := v1[32768]
    fmt.Println(v2, v5)
    v12 := GenS1{F0: -1, F1: 0}
    fmt.Println(v5, v2, v3)
    for fi13 := 0; fi13 < 4; fi13++ {
        v5 += 32767
        _ = fi13
    }
    v14 := GenS2{F0: true}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
