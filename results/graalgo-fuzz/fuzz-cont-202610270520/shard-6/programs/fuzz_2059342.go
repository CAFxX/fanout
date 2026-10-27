package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint64 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = -1
    v3 := make(map[uint32]string)
    v2 += v2
    v4 := []int{-2147483648, -9223372036854775808, 32767, -403}
    switch 32768 {
    case -706:
        v2 += -1
    case 127:
        fmt.Println(v2)
    }
    var v5 int = 157
    fmt.Println(v5, v2)
    switch "a" {
    case "a":
        var v6 bool = (true || true)
        _ = v6
    default:
        var v7 uint = 65535
        _ = v7
    }
    _ = len(v4)
    var v8 int16 = (v2 ^ (v2 + v2))
    v4 = append(v4, 42)
    v8 += v2
    v9 := GenS1{F0: 42, F1: 793, F2: 6}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
