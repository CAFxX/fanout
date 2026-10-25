package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = "a"
    v3 := []uint32{33, 314, 491, 972}
    var v4 float32 = 0.1
    var v5 bool = true
    if (65535 <= 645) {
        _ = len(v3)
        fmt.Println(v4)
        var v6 int8 = 9
        _ = v6
    } else {
        if true {
            v5 = false
        } else {
            fmt.Println(v2)
        }
    }
    fmt.Println(v4, v5)
    v2 = "\n\t"
    var v7 int64 = 42
    v3 = append(v3, 4294967295)
    v4 = v4
    v8 := GenS1{F0: 42, F1: "foo bar", F2: -644}
    var v9 int16 = -255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
