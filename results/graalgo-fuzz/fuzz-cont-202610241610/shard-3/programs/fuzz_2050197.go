package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int64; F1 bool; F2 int32 }

func main() {
    var v1 float64 = -2.25
    var v2 uint = 255
    v3 := []string{}
    v4 := make(map[int]string)
    v5 := []int{32768, 127, 128}
    v6 := make(map[string]string)
    var v7 int16 = -1
    fmt.Println(v2, v7)
    v7 = 1
    switch -2147483648 {
    case -2147483648:
        v8 := GenS1{F0: 800, F1: -749}
        _ = v8
    default:
        v2 += (^(v2))
    }
    v4[42] = ("a" + "a")
    v2 = 42
    v2 = (v2 | 18446744073709551615)
    _ = len(v3)
    v9 := GenS2{F0: 42, F1: false, F2: -151}
    fmt.Println(v7)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
