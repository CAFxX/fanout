package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint32; F2 int }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 646
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 112
    v5 := []int{-2147483648, 4294967295}
    v2 = 18446744073709551615
    v1[2] = 32767
    var v6 uint64 = 4294967295
    v4 -= v4
    var v7 uint8 = (v4 >> ((uint(v6)) & 63))
    v3[0] = 4294967295
    switch ("hello" + "") {
    case "foo bar":
        var v8 uint32 = 1
        _ = v8
    case "\n\t":
        v4 = 128
    case "":
        var v9 int64 = 9223372036854775807
        _ = v9
    default:
        v10 := GenS1{F0: 65535}
        _ = v10
    }
    v11 := GenS1{F0: v2}
    fmt.Println(v4, v6)
    v12 := GenS2{F0: 0, F1: 65535, F2: 42}
    var v13 int32 = -701
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
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
}
