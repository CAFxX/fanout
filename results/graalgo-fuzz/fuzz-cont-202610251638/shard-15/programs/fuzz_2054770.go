package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 255
    var v3 int64 = 42
    var v4 float32 = ((0.0 - (3.14159 * -2.25)) / 1.0)
    var v5 float32 = -2.25
    var v6 int16 = 1
    var v7 int8 = 42
    v8 := []int{-2147483648, 632, 2147483647}
    var v9 uint64 = 814
    v3 += (v3 ^ int64(v9))
    v7 -= v7
    var v10 uint16 = 648
    fmt.Println(v2, v9, v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
