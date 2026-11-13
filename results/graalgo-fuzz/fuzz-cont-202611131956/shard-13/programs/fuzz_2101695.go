package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint16 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]int)
    v3 := []int{783, 9223372036854775807, 42, -9223372036854775808}
    var v4 int64 = 470
    v5 := []int{}
    var v6 uint8 = 127
    var v7 uint = 42
    fmt.Println(v4, v4, v6)
    fmt.Println(v4, v6, v7)
    v2[2147483647] = -9223372036854775808
    v8 := GenS1{F0: v4, F1: -128, F2: 65535}
    var v9 uint8 = v6
    v7 += (v7 / ((2) | 1))
    v7 += (v7 ^ (v7 / ((18446744073709551615) | 1)))
    v4 += v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
