package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint32 }

func main() {
    var v1 bool = (false && false)
    var v2 uint64 = 42
    v3 := []uint32{325}
    v4 := make(map[string]int)
    var v5 uint8 = 63
    v6 := []uint32{2, 1, 0}
    v1 = (694 <= -667)
    switch 1 {
    case 0:
        v1 = (815 >= 65535)
    default:
        v7 := GenS1{F0: v5, F1: 53, F2: 681}
        _ = v7
    }
    v4["a"] = -146
    var v8 uint = 42
    var v9 uint32 = 980
    v9 = (v9 << (((v8 & v8)) & 63))
    var v10 float64 = (0.0 * (1e10 + (3.14159 * 0.1)))
    var v11 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
