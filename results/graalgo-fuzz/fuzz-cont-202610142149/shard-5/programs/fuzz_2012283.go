package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int32 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]bool)
    var v3 uint64 = 0
    var v4 int64 = 42
    v5 := GenS1{F0: 32767, F1: v4, F2: -2147483648}
    v3 -= (v3 >> ((961) & 63))
    v6 := []int{2147483647, -9223372036854775808, -882, 786}
    switch "foo bar" {
    case "a":
        v4 -= (v4 >> ((165) & 63))
    case "foo bar":
        v3 -= (v3 * v3)
    }
    v7 := GenS1{F0: 1000, F1: v4, F2: 1}
    var v8 int32 = 2147483647
    var v9 string = "foo bar"
    fmt.Println(v9, v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
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
