package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint64 }

func main() {
    var v1 int16 = -773
    var v2 int8 = 42
    v3 := []uint32{42, 0}
    v4 := make(map[uint32]string)
    var v5 uint16 = 42
    var v6 int16 = 32767
    fmt.Println(v1, v6)
    v7 := GenS1{F0: v2, F1: true, F2: 2}
    v6 -= (v6 % (((v1 / (((-(v6))) | 1))) | 1))
    var v8 int64 = 1
    delete(v4, 1)
    var v9 bool = true
    fmt.Println(v8, v6)
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
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
