package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int64 }
func genFn0(p0 uint8, p1 uint32) int {
    fmt.Println(p1, p1)
    switch "" {
    case "hello":
        p1 = p1
    case "\n\t":
        p1 = (p1 & p1)
    }
    _ = p0
    _ = p1
    return -1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]int)
    v3 := []uint32{}
    var v4 int = -9223372036854775808
    var v5 uint16 = 32767
    v6 := make(map[uint32]int)
    fmt.Println(v4)
    fmt.Println(v5, v4)
    v4 = v4
    v4 = 9223372036854775807
    var v7 int = v4
    var v8 bool = true
    v8 = v8
    v5 -= v5
    switch genFn0(127, 351) {
    case 1:
        v9 := GenS1{F0: v4, F1: v5, F2: -9223372036854775808}
        _ = v9
    case -443:
        fmt.Println(v8, v4)
    case -959:
        var v10 int32 = 611
        _ = v10
    }
    v7 = 42
    var v11 int32 = -110
    fmt.Println(v4, v11)
    var v12 int = (v7 + v4)
    var v13 uint8 = 42
    fmt.Println(v8, v12)
    if (v8 || v8) {
        var v14 uint64 = 42
        v13 = v13
        _ = v14
    } else {
        _ = len(v3)
        var v15 int8 = 100
        _ = v15
    }
    v12 = v4
    fmt.Println(v13, v13, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
