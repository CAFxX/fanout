package main

import (
    "fmt"
)


func main() {
    var v1 int32 = 42
    v2 := make(map[int]bool)
    v3 := make(map[uint32]int)
    var v4 uint32 = 42
    var v5 string = "a"
    fmt.Println(v4, v1, v4)
    v4 = (v4 / ((v4) | 1))
    var v6 uint16 = 65535
    v5 = "a"
    _ = len(v2)
    var v7 int8 = 1
    var v8 int = -84
    v9 := []int{v8, (v8 ^ v8), v8, (^(v8))}
    switch (v5 + v5) {
    case "":
        v4 -= v4
    default:
        var v10 uint32 = 0
        _ = v10
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
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
