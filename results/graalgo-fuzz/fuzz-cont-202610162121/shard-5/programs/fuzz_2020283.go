package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint }

func main() {
    v1 := []uint32{537, 0, 2, 2147483648}
    v2 := make(map[string]bool)
    v3 := make(map[int]string)
    var v4 int64 = -9223372036854775808
    var v5 int8 = 42
    var v6 int32 = 2147483647
    if false {
        v4 += v4
    }
    var v7 int64 = 787
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    v5 -= v5
    var v8 uint = 1
    v1 = append(v1, 0)
    fmt.Println(v7, v4, v8)
    v9 := GenS1{F0: 127, F1: 0, F2: v8}
    switch -963 {
    case -770:
        v7 = v4
    case 32768:
        v7 += int64(v5)
    }
    v5 = v5
    var v10 uint8 = 0
    switch 42 {
    case 959:
        v6 = v6
    }
    var v11 float32 = 1.5
    var v12 float64 = 0.0
    v13 := v3[0]
    v5 = -48
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
