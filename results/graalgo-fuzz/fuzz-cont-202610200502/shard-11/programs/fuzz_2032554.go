package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 float32 = 0.1
    var v2 uint = 2
    var v3 int32 = 42
    var v4 uint64 = 1
    var v5 int16 = -1
    var v6 uint32 = 2147483647
    v7 := GenS1{F0: 746, F1: -9223372036854775808}
    v8 := []int{32767, 127, -2147483648, -669}
    if (65535 != (^(v6))) {
        fmt.Println(v4, v1, v1)
        v5 -= -32768
        v1 = v1
    }
    v2 -= 219
    if true {
        v9 := GenS2{F0: v5}
        _ = v9
    }
    v10 := GenS1{F0: -141, F1: -158}
    v1 = v1
    v2 -= uint(v4)
    var v11 uint16 = 31
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
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
