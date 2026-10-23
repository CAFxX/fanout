package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 string }

func main() {
    v1 := []int{}
    v2 := make(map[uint32]uint32)
    var v3 int16 = 42
    var v4 uint32 = 2147483647
    var v5 int = 42
    v2[v4] = (v4 * (v4 - v4))
    var v6 bool = (!true)
    v7 := GenS1{F0: v5, F1: -9223372036854775808, F2: ("foo bar" + "foo bar")}
    v3 = 1
    v2[1] = v4
    if (-2147483648 >= 1) {
        var v8 uint = 2
        _ = v8
    }
    v6 = (255 <= 128)
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
