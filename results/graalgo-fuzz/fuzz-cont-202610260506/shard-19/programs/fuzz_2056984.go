package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int32 }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint8 = 0
    var v3 int64 = 1
    v4 := []int{566, 0}
    var v5 int32 = -2147483648
    var v6 uint = 1
    var v7 int = -9223372036854775808
    v8 := GenS1{F0: 2, F1: -281, F2: v5}
    v1 = (-(v1))
    _ = len(v4)
    fmt.Println(v2, v1, v6)
    var v9 uint32 = 2
    var v10 uint8 = 42
    fmt.Println(v7, v5)
    var v11 int = (v7 - int(v5))
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
