package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint; F1 uint; F2 uint32 }

func main() {
    v1 := []uint32{2147483648, 4294967295, 42}
    v2 := []int{-2147483648, -1, 0}
    var v3 int16 = -909
    var v4 uint64 = 1
    fmt.Println(v4, v4)
    var v5 int64 = -111
    var v6 int64 = 0
    v7 := GenS1{F0: (v3 - v3), F1: 65535, F2: -253}
    v8 := GenS2{F0: 18446744073709551615, F1: 2, F2: 809}
    fmt.Println(v6)
    var v9 bool = false
    if len(v1) > 0 {
        v1[0] = 65535
    }
    fmt.Println(v6, v5, v4)
    v9 = true
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
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
