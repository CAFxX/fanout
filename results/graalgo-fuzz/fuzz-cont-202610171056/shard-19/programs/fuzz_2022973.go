package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int8 }
func genFn0(p0 int32, p1 uint64) string {
    p1 = (p1 + p1)
    v4 := GenS1{F0: 1, F1: (!false), F2: 42}
    p1 += (p1 % (((^(p1))) | 1))
    var v5 uint32 = 0
    v6 := []int{128}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return "hello"
}

func main() {
    v1 := []int{1, 32767, 1}
    var v2 int32 = 2147483647
    v3 := []int{0, 127, 32768, 32768}
    fmt.Println(v2)
    v2 = v2
    v2 -= (v2 >> ((2) & 63))
    var v7 int = -442
    var v8 float32 = float32((float64(1e10) / 1.5))
    var v9 float32 = 3.14159
    fmt.Println(v7, v8)
    var v10 string = ""
    if len(v1) > 0 {
        v1[0] = 32768
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
