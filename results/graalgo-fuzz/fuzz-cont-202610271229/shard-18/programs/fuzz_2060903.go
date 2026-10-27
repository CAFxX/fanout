package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 int64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 string = "hello"
    var v3 uint64 = 42
    v4 := []int{9223372036854775807, 0}
    v5 := GenS1{F0: false, F1: 0, F2: 1}
    if true {
        var v6 int64 = 1
        v7 := GenS1{F0: (1 >= 100), F1: -54, F2: (v6 >> ((65535) & 63))}
        _ = v6
        _ = v7
    }
    v3 -= v3
    var v8 int64 = 42
    fmt.Println(v8, v8)
    fmt.Println(v2)
    v2 = v2
    var v9 uint64 = v3
    var v10 uint = 601
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
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
