package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int64 }
type GenS2 struct { F0 int8; F1 bool }

func main() {
    var v1 string = "foo bar"
    v2 := []uint32{65535, 407, 65535}
    var v3 uint32 = 65535
    v4 := make(map[string]int)
    v5 := make(map[string]bool)
    var v6 uint8 = 42
    var v7 float32 = 0.0
    v6 = uint8(v3)
    var v8 int64 = 9223372036854775807
    fmt.Println(v1, v7)
    fmt.Println(v6, v8)
    if true {
        v9 := GenS1{F0: uint32(v6), F1: 1, F2: v8}
        v4["hello"] = 128
        v10 := GenS2{F0: 112, F1: false}
        _ = v10
        _ = v9
    }
    v8 -= v8
    var v11 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
