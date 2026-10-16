package main

import (
    "fmt"
)


func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int32 = 2147483647
    var v3 int = 1
    v2 -= v2
    if false {
        v3 = v3
        fmt.Println(v2, v2, v2)
    }
    fmt.Println(v1, v1, v1)
    var v4 int32 = 0
    var v5 uint8 = 0
    v1 = (v1 << ((42) & 63))
    v6 := []int{-954}
    v1 += 0
    var v7 uint8 = 128
    v4 -= v2
    v7 += (v5 >> ((209) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
