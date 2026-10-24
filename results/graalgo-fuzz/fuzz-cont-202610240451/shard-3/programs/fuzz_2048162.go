package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }

func main() {
    var v1 int16 = 614
    v2 := make(map[string]uint32)
    var v3 int64 = 434
    var v4 uint8 = 38
    var v5 bool = true
    var v6 uint32 = 645
    fmt.Println(v5)
    v5 = ((true || ((!v5) || v5)) || true)
    fmt.Println(v5)
    v7 := GenS1{F0: (v4 | (v4 ^ v4)), F1: -9223372036854775808}
    var v8 int8 = 0
    if v5 {
        v9 := GenS1{F0: (v4 - v4), F1: 9223372036854775807}
        _ = v9
    } else {
        var v10 bool = (!(1e10 >= (1.5 + 1.5)))
        _ = v10
    }
    var v11 uint64 = 272
    var v12 uint64 = v11
    v13 := []int{}
    _ = len(v13)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
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
