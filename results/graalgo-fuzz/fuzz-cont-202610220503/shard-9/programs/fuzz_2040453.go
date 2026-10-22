package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := []uint32{2147483647}
    v2 := []int{795, 0, 9223372036854775807, 2147483647}
    var v3 bool = (251 != 1)
    var v4 uint16 = 1
    v5 := make(map[int]uint32)
    v6 := []int{-521, 9223372036854775807}
    _ = len(v5)
    var v7 bool = (1 < 1)
    var v8 int16 = 42
    if (150 <= 128) {
        var v9 int8 = 73
        _ = v9
    }
    v10 := GenS1{F0: (v8 << ((38) & 63))}
    _ = len(v5)
    v11 := GenS1{F0: (v8 ^ 1000)}
    v7 = (((!true) || v7) && false)
    switch -650 {
    case 9223372036854775807:
        var v12 int8 = -15
        _ = v12
    case 32767:
        var v13 uint16 = (v4 >> ((42) & 63))
        _ = v13
    case -2147483648:
        var v14 uint8 = 71
        _ = v14
    }
    var v15 int8 = -128
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
