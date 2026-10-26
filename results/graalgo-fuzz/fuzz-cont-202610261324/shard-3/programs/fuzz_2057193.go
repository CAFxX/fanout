package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint }
type GenS2 struct { F0 uint; F1 int8; F2 string }

func main() {
    var v1 int16 = -32768
    var v2 int8 = 100
    v3 := make(map[uint32]bool)
    v4 := make(map[int]int)
    var v5 float64 = 1e10
    var v6 float32 = 0.1
    fmt.Println(v6)
    v7 := GenS1{F0: 42, F1: 1, F2: 42}
    var v8 int = 822
    v9 := []int{(v8 >> ((18446744073709551615) & 63)), 590, v8}
    if false {
        v10 := GenS2{F0: 255, F1: v2, F2: "a"}
        _ = len(v9)
        _ = v10
    } else {
        if len(v9) > 0 {
            v9[0] = v8
        }
        var v11 int64 = -9223372036854775808
        _ = v11
    }
    var v12 int64 = 0
    var v13 float32 = v6
    var v14 int32 = 0
    _ = len(v9)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
