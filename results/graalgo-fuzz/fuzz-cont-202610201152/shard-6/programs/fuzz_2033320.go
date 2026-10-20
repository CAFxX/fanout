package main

import (
    "fmt"
)

func genFn0(p0 uint32, p1 int8, p2 uint16) int16 {
    var v6 int64 = 982
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -93
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = 33
    v3 := []string{}
    v4 := make(map[int]uint32)
    v5 := make(map[int]uint32)
    fmt.Println(v2)
    var v8 uint64 = 42
    var v9 int8 = 1
    v2 -= v2
    var v10 int = -157
    fmt.Println(v2, v8, v10)
    v10 = v10
    delete(v1, 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
