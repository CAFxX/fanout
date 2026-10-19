package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 string }

func main() {
    var v1 bool = (65535 < 110)
    v2 := make(map[int]string)
    var v3 float32 = 1.5
    _ = len(v2)
    var v4 uint32 = 0
    fmt.Println(v3, v3, v4)
    v5 := GenS1{F0: 1, F1: 0, F2: "hello"}
    var v6 uint8 = 128
    v6 -= (v6 ^ v6)
    var v7 int8 = -1
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
