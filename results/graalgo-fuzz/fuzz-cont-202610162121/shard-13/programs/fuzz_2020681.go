package main

import (
    "fmt"
)


func main() {
    var v1 uint16 = 255
    var v2 int64 = -269
    var v3 uint64 = 0
    v4 := make(map[int]bool)
    var v5 uint64 = 1
    fmt.Println(v2)
    fmt.Println(v5, v2, v3)
    fmt.Println(v2, v2)
    v6 := []int{}
    var v7 int16 = -1
    v5 += v3
    if len(v6) > 0 {
        v6[0] = 4294967295
    }
    v8 := v4[-2147483648]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
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
