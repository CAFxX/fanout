package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 uint32 }

func main() {
    var v1 int64 = 1
    var v2 int16 = -32768
    var v3 int64 = v1
    var v4 uint8 = 2
    v5 := make(map[string]uint32)
    var v6 float64 = 0.0
    v7 := []int{-2147483648, 127, -325}
    var v8 uint8 = 179
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 uint8 = v4
        _ = fi9
        _ = v10
    }
    v6 += (3.14159 + (v6 * v6))
    switch "a" {
    case "":
        v1 += (v3 >> ((0) & 63))
    }
    v11 := GenS1{F0: (v3 - -9223372036854775808)}
    var v12 int64 = v3
    var v13 uint64 = 761
    var v14 uint = 305
    var v15 int64 = int64(v2)
    switch "foo bar" {
    case "\n\t":
        v16 := GenS2{F0: (((^(v1)) > v1) || false), F1: 2}
        _ = v16
    default:
        v13 = uint64(v14)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
