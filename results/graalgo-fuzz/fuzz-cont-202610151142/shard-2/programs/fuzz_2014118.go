package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int }
func genFn0(p0 int16, p1 int16) int64 {
    v7 := GenS1{F0: "a", F1: false, F2: -815}
    var v8 bool = true
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]int)
    v3 := make(map[string]int)
    v4 := make(map[int]int)
    var v5 int8 = 122
    var v6 uint8 = 94
    v5 = v5
    v5 = 42
    v9 := []int{-1, 127}
    v6 = 191
    v10 := GenS1{F0: "a", F1: (534 != 1), F2: 32767}
    var v11 int64 = 9223372036854775807
    v11 = v11
    v6 = v6
    v6 = (v6 ^ 2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
