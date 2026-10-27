package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint16; F1 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 42
    var v4 uint = 2
    v5 := []uint32{2, 456, 1, 2147483647}
    fmt.Println(v4, v4)
    var v6 int8 = 92
    v7 := GenS1{F0: v6, F1: "foo bar"}
    fmt.Println(v4)
    fmt.Println(v3)
    v8 := GenS2{F0: 42, F1: 112}
    var v9 bool = true
    v10 := GenS2{F0: 255, F1: uint8(v3)}
    var v11 string = (("" + "\n\t") + "\n\t")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
