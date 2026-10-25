package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
func genFn0(p0 uint64, p1 uint8, p2 int8) int16 {
    p2 = p2
    v6 := []int{297, 780, 76}
    p1 = p1
    v7 := GenS1{F0: 2147483647, F1: ((!(uint64(p1) < p0)) || true)}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int8 = -1
    v3 := make(map[int]int)
    v4 := []uint32{2147483648, 744, 4294967295}
    var v5 float64 = 0.0
    var v8 int16 = 42
    v8 = (v8 >> ((42) & 63))
    v9 := GenS1{F0: 2147483647, F1: (-1 != v2)}
    fmt.Println(v5, v8, v2)
    v8 -= v8
    if len(v4) > 0 {
        v4[0] = 865
    }
    v8 -= 0
    var v10 int8 = (v2 % ((v2) | 1))
    var v11 int32 = 602
    var v12 int64 = (v1 * v1)
    v11 = v11
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
