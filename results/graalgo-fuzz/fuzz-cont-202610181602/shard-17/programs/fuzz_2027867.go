package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }

func main() {
    var v1 int8 = 121
    var v2 uint8 = 128
    var v3 float32 = ((0.1 - 0.1) / 3.14159)
    v4 := []string{"hello", "", ""}
    v5 := []int{1}
    v6 := []int{-385}
    fmt.Println(v2, v2, v1)
    v7 := GenS1{F0: true, F1: -150}
    if (48 >= 2) {
        for fi8 := 0; fi8 < 13; fi8++ {
            fmt.Println(fi8, v3, fi8)
            v1 += (v1 - v1)
            _ = fi8
        }
    }
    var v9 int32 = 42
    v2 = (^(v2))
    v2 += (^(v2))
    fmt.Println(v1)
    var v10 int64 = -9223372036854775808
    var v11 uint32 = 2
    var v12 uint = 802
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
