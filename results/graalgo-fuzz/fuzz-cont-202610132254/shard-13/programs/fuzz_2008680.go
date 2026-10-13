package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 int; F1 uint64 }

func main() {
    var v1 int32 = -1
    var v2 uint = 1
    v3 := make(map[int]int)
    for fi4 := 0; fi4 < 1; fi4++ {
        var v5 uint = 664
        v5 += v5
        var v6 int16 = 975
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 int32 = (v1 / ((v1) | 1))
    v8 := GenS1{F0: -30, F1: -2147483648}
    fmt.Println(v7, v7, v2)
    v2 -= 0
    var v9 int32 = 262
    v10 := GenS2{F0: -9223372036854775808, F1: 288}
    v11 := []int{674, -2147483648, 564, 813}
    var v12 string = "hello"
    fmt.Println(v7, v1)
    var v13 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
