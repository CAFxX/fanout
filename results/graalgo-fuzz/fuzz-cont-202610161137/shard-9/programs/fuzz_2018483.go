package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint8 = 42
    v3 := make(map[string]bool)
    v4 := make(map[uint32]uint32)
    v5 := []int{}
    v2 -= (v2 - (v2 | (v2 | (v2 << ((4294967295) & 63)))))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v6 := GenS1{F0: -2147483648, F1: 0}
    v2 += (v2 << ((65535) & 63))
    v1 += 1
    fmt.Println(v2, v2)
    v1 += v1
    switch "foo bar" {
    case "foo bar":
        var v7 int64 = v1
        _ = v7
    }
    v2 = (v2 >> ((1) & 63))
    var v8 int8 = 1
    var v9 int32 = 339
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
