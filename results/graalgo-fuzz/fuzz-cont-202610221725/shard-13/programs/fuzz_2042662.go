package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]bool)
    var v3 uint16 = 255
    v4 := []int{-824, int(v3), -1, 2147483647}
    var v5 int64 = 0
    v6 := GenS1{F0: (v3 & v3), F1: 42, F2: 0}
    fmt.Println(v3)
    v5 += (v5 << ((18446744073709551615) & 63))
    v3 += v3
    var v7 int8 = 100
    v5 += (v5 >> ((645) & 63))
    var v8 string = "\n\t"
    v9 := GenS1{F0: v3, F1: (-(v7)), F2: 32767}
    var v10 bool = false
    var v11 uint64 = 317
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
