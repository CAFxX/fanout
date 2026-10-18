package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 string }

func main() {
    var v1 uint16 = 1
    var v2 bool = (!(1 != 42))
    v3 := make(map[uint32]int)
    v4 := []int{-1, 4294967295, -505}
    v5 := GenS1{F0: 4294967295, F1: -1, F2: (((("a" + "") + "foo bar") + "\n\t") + "foo bar")}
    fmt.Println(v1, v1)
    v4 = append(v4, -182)
    fmt.Println(v2, v1)
    delete(v3, 2)
    _ = len(v4)
    if (20 < 0) {
        fmt.Println(v1)
    } else {
        _ = len(v4)
        var v6 int16 = 42
        v6 += (v6 / (((-(v6))) | 1))
        _ = v6
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
