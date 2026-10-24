package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 string }
func genFn0(p0 uint8, p1 uint, p2 string) int16 {
    v4 := GenS1{F0: p0, F1: 1, F2: p2}
    p2 = p2
    v5 := []int{2147483647, 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 32767
}

func main() {
    var v1 int8 = 0
    v2 := make(map[int]string)
    var v3 int16 = 545
    v6 := []int{-1}
    v3 = -32768
    v7 := v2[-9223372036854775808]
    v3 += v3
    fmt.Println(v1, v1, v3)
    var v8 uint16 = 42
    v7 = (v7 + "")
    fmt.Println(v7)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
