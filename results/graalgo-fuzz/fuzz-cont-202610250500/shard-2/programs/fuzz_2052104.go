package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint64 }
func genFn0(p0 uint8, p1 uint64, p2 uint32) uint32 {
    p1 = p1
    v4 := GenS1{F0: -218, F1: 100, F2: (p1 & p1)}
    var v5 uint = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p2
}

func main() {
    var v1 int8 = 49
    var v2 uint8 = 38
    v3 := make(map[int]int)
    v6 := GenS1{F0: 4294967295, F1: (v1 + v1), F2: 2}
    var v7 int = 42
    v7 += v7
    if (-820 >= 148) {
        v8 := []int{(v7 << ((82) & 63)), 4294967295, (v7 % (((-(v7))) | 1)), v7}
        _ = v8
    }
    var v9 uint32 = 2147483647
    var v10 float64 = (((1.5 * 1e10) * (-2.25 - ((1.5 / 3.14159) / 1.0))) + 3.14159)
    fmt.Println(v1)
    v11 := GenS1{F0: -816, F1: v1, F2: 1}
    v2 += 1
    v12 := []int{(v7 << ((586) & 63)), (^(v7))}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v9
}
