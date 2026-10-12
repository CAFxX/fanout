package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 bool }

func main() {
    v1 := []int{-1, 9223372036854775807, -23}
    var v2 float64 = float64(0.1)
    v3 := make(map[uint32]bool)
    var v4 int64 = 0
    var v5 string = ("foo bar" + "")
    v4 = (v4 / ((42) | 1))
    fmt.Println(v5, v4)
    _ = len(v3)
    delete(v3, 4294967295)
    v6 := GenS1{F0: -300, F1: 127, F2: ("a" >= "hello")}
    fmt.Println(v4, v5, v4)
    var v7 string = v5
    if (!false) {
        v8 := GenS1{F0: 42, F1: 0, F2: true}
        _ = v8
    }
    var v9 int32 = -1
    v7 = v5
    if true {
        v2 -= 1.5
    } else {
        v2 -= ((v2 / -2.25) - v2)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
