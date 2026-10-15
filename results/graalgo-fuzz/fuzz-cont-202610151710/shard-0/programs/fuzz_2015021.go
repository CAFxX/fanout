package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = (false && (2 <= 255))
    var v3 uint8 = 42
    var v4 string = "a"
    var v5 uint = 0
    v6 := []uint32{277, 4294967295, 247}
    var v7 float32 = 0.0
    var v8 float32 = v7
    if (0 >= 65535) {
        fmt.Println(v4, v8)
    }
    v2 = v2
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v9 := GenS1{F0: 65535, F1: 948, F2: 246}
    var v10 int64 = 42
    var v11 uint = (v5 + 2)
    v10 = v10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
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
