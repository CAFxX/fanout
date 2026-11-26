package main

import (
    "fmt"
)


func main() {
    var v1 int64 = 0
    v2 := make(map[string]uint32)
    var v3 int64 = (v1 % ((v1) | 1))
    var v4 float32 = 1.5
    var v5 uint8 = 42
    var v6 uint16 = 372
    var v7 int8 = 100
    var v8 int8 = (v7 + v7)
    v3 -= v1
    v5 = v5
    var v9 uint64 = 4294967295
    v10 := []int{}
    v9 = (v9 >> ((18446744073709551615) & 63))
    fmt.Println(v3, v8)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
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
