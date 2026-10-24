package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }

func main() {
    var v1 uint64 = 996
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]string)
    v4 := make(map[string]bool)
    v2[65535] = 383
    v1 += (v1 % ((v1) | 1))
    fmt.Println(v1)
    fmt.Println(v1)
    var v5 uint8 = 0
    var v6 uint32 = 2
    v7 := GenS1{F0: v6, F1: int(v5)}
    v8 := GenS1{F0: (^(v6)), F1: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
