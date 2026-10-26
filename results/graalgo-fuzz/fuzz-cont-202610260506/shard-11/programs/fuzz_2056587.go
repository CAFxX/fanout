package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
func genFn0(p0 string, p1 uint32, p2 string) bool {
    var v4 uint8 = 2
    v5 := GenS1{F0: 32768, F1: 32767}
    var v6 int64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return false
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = true
    var v3 uint32 = 42
    var v7 int64 = -1
    v3 += 976
    var v8 int16 = 42
    v2 = v2
    var v9 int = 127
    var v10 uint = 1
    v11 := []int{(v9 ^ (v9 * 1)), v9, v9}
    if v2 {
        var v12 int8 = 100
        fmt.Println(v8, v7, v3)
        _ = v12
    }
    v13 := v1[(v3 & (v3 - 520))]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
