package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
func genFn0(p0 int64, p1 int, p2 uint8) uint64 {
    fmt.Println(p0, p1)
    p1 += (-(p1))
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}

func main() {
    var v1 uint64 = 452
    var v2 uint32 = 771
    v3 := make(map[int]string)
    var v4 float32 = 1.5
    switch -2147483648 {
    case 127:
        v5 := []int{276, 32767}
        _ = v5
    }
    v6 := GenS1{F0: 93, F1: true}
    var v7 float64 = (((1e10 + 1.5) + (float64(v4) + ((-2.25 / 1.5) * 1e10))) / 1.0)
    if true {
        v8 := v3[32768]
        _ = v8
    }
    v9 := []int{}
    var v10 uint = 2
    v7 -= float64(float32(v7))
    v11 := GenS1{F0: 2, F1: true}
    var v12 int32 = int32(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v9
}
