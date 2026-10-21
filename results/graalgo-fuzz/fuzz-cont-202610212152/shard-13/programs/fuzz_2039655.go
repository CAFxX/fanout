package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 bool }

func main() {
    var v1 float64 = (0.1 - -2.25)
    var v2 uint = 2
    v3 := []string{}
    var v4 uint32 = 0
    v5 := make(map[int]int)
    v5[9223372036854775807] = -695
    var v6 int8 = 0
    v6 = -1
    var v7 string = "foo bar"
    fmt.Println(v6, v6)
    v2 += v2
    v6 = 100
    var v8 int8 = -1
    v4 -= (^(v4))
    v9 := GenS1{F0: v2, F1: true, F2: false}
    fmt.Println(v6)
    v2 -= (v2 % ((v2) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
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
