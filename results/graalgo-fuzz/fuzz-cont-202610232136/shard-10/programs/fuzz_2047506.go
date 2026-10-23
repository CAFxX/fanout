package main

import (
    "fmt"
)


func main() {
    var v1 int8 = -1
    v2 := make(map[uint32]int)
    var v3 int64 = 42
    var v4 int = -9223372036854775808
    v5 := []int{-2147483648, v4, v4}
    var v6 float64 = float64(0.0)
    v3 -= (v3 / ((42) | 1))
    v7 := v2[0]
    v7 -= 32767
    v2[34] = (-(v7))
    fmt.Println(v7, v4, v6)
    var v8 float64 = (v6 - (3.14159 - (1.5 + 1.5)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
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
}
