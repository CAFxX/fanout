package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }

func main() {
    var v1 uint64 = 982
    var v2 uint16 = 2
    var v3 uint64 = (v1 << ((4294967295) & 63))
    var v4 int = 128
    v5 := make(map[uint32]bool)
    v6 := make(map[uint32]uint32)
    var v7 uint64 = v3
    var v8 float64 = 0.1
    v3 = (v7 % (((v7 - (v3 & (^(v7))))) | 1))
    v9 := GenS1{F0: 255, F1: (true && (!true))}
    var v10 uint8 = 255
    v4 = 32767
    var v11 float32 = (0.0 - 0.0)
    fmt.Println(v3, v7)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
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
