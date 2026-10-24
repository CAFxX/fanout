package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }

func main() {
    var v1 uint = 248
    v2 := []uint32{924, 2147483647, 2147483648, 962}
    var v3 int8 = 1
    v4 := []int{127, 4294967295}
    var v5 int16 = 983
    v6 := make(map[uint32]bool)
    v1 += (^(v1))
    var v7 uint64 = 18446744073709551615
    var v8 int32 = -1
    v9 := GenS1{F0: 0, F1: 4294967295}
    switch "foo bar" {
    case "\n\t":
        v1 = uint(v7)
    }
    v10 := v6[0]
    var v11 uint = 673
    var v12 uint = 42
    v1 -= v11
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 uint64 = (v7 - (v7 * 2))
        _ = fi13
        _ = v14
    }
    fmt.Println(v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
