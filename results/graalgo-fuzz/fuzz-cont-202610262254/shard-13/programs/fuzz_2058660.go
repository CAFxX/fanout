package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]bool)
    var v2 int32 = -466
    var v3 uint64 = 284
    v2 = 937
    var v4 float32 = (float32(1.5) + -2.25)
    var v5 uint64 = (v3 ^ 2)
    v6 := []int{}
    v3 = v5
    v4 += v4
    if (true && (!true)) {
        var v7 uint32 = 925
        _ = v7
    }
    fmt.Println(v4)
    var v8 int8 = -128
    fmt.Println(v4, v3, v3)
    v3 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
