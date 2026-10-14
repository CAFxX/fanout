package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]string)
    var v3 float32 = (0.1 * 1e10)
    v4 := []uint32{}
    v5 := make(map[int]bool)
    var v6 uint8 = 33
    v5[911] = false
    v3 += v3
    v5[-1] = false
    var v7 float32 = (float32(1.5) + 0.0)
    fmt.Println(v7)
    v6 = 88
    _ = len(v4)
    v8 := v5[42]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
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
