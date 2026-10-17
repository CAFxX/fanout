package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 bool, p1 uint64, p2 int8) uint16 {
    p1 = (p1 | p1)
    p2 = (p2 % ((p2) | 1))
    p1 -= 956
    var v4 int8 = (p2 >> ((255) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 554
}
func genFn1(p0 uint32) int32 {
    p0 = (p0 ^ (p0 ^ p0))
    for fi5 := 0; fi5 < 3; fi5++ {
        v6 := []int{}
        p0 -= p0
        _ = fi5
        _ = v6
    }
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return -1
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    v7 := GenS1{F0: uint32(v1)}
    var v8 float32 = 3.14159
    delete(v2, -9223372036854775808)
    if (false && (1e10 != -2.25)) {
        v1 = v1
    }
    _ = len(v3)
    var v9 uint64 = (v1 * (v1 | v1))
    switch -201 {
    case 9223372036854775807:
        v8 = ((1e10 - v8) / (v8 * v8))
    }
    v10 := []int{}
    v8 += 0.1
    var v11 uint32 = 2147483647
    var v12 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
