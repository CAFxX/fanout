package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint8 }

func main() {
    var v1 bool = false
    v2 := make(map[string]bool)
    v3 := make(map[string]uint32)
    var v4 uint8 = 58
    var v5 uint16 = 35
    v4 = (v4 * v4)
    v6 := GenS1{F0: 127}
    var v7 bool = false
    v8 := GenS1{F0: 1}
    if false {
        v9 := GenS1{F0: 62}
        _ = v9
    }
    delete(v2, "")
    v5 -= 1
    v10 := GenS2{F0: 2147483647, F1: 42, F2: v4}
    var v11 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
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
}
