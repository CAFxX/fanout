package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
func genFn0(p0 uint8, p1 int) uint32 {
    v5 := []int{-920, (p1 | (p1 & p1)), p1}
    p1 = -888
    p1 = p1
    _ = p0
    _ = p1
    _ = v5
    return 821
}

func main() {
    var v1 int32 = 2147483647
    var v2 int32 = v1
    var v3 int32 = (^(v1))
    v4 := make(map[string]int)
    if (true || false) {
        fmt.Println(v2, v1, v1)
    }
    switch "foo bar" {
    case "foo bar":
        var v6 int64 = -1
        _ = v6
    }
    switch int(v3) {
    case -795:
        v7 := GenS1{F0: 0, F1: 42}
        _ = v7
    default:
        var v8 uint32 = 65535
        _ = v8
    }
    v9 := []int{32767, 418, 0}
    var v10 int32 = v3
    v11 := GenS1{F0: 94, F1: 42}
    var v12 int8 = 42
    var v13 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
