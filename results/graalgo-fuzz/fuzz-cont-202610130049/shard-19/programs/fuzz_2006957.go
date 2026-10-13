package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint }

func main() {
    var v1 uint = 2
    var v2 int16 = 0
    var v3 int32 = -2147483648
    var v4 uint16 = 0
    var v5 int64 = -701
    v4 -= (v4 << ((v1) & 63))
    var v6 uint = v1
    var v7 float32 = 1e10
    fmt.Println(v4, v3, v5)
    v8 := GenS1{F0: (^(v4)), F1: 255, F2: v6}
    v9 := []int{1, 323, 2147483647}
    v6 -= (v1 - 42)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
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
