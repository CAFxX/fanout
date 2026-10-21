package main

import (
    "fmt"
)

func genFn0(p0 int64, p1 int16) int16 {
    fmt.Println(p1, p0)
    v6 := []int{-9, -1, 1, -230}
    _ = p0
    _ = p1
    _ = v6
    return (p1 * p1)
}

func main() {
    var v1 string = "a"
    var v2 bool = true
    v3 := make(map[int]int)
    var v4 uint16 = 2
    v5 := make(map[uint32]int)
    v7 := []int{}
    if v2 {
        var v8 int64 = 0
        _ = v8
    } else {
        _ = len(v5)
    }
    var v9 float32 = ((((0.0 + 0.0) * (1.5 + 1.5)) * 3.14159) / 1.5)
    v5[65535] = 128
    v2 = false
    v10 := v5[480]
    v2 = false
    if true {
        var v11 string = v1
        _ = v11
    }
    var v12 int8 = 50
    var v13 uint16 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
