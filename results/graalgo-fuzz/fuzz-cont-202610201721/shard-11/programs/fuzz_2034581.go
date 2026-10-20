package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint64 }

func main() {
    var v1 int8 = 0
    v2 := make(map[string]int)
    var v3 uint16 = 133
    var v4 int32 = -2147483648
    var v5 int32 = v4
    fmt.Println(v5, v1)
    if (810 >= 1) {
        v6 := GenS1{F0: 0, F1: -9223372036854775808}
        _ = v6
    } else {
        fmt.Println(v4, v1, v5)
        v7 := []int{-859, 32767, 0, 32768}
        _ = v7
    }
    fmt.Println(v5, v4, v4)
    v8 := GenS1{F0: 889, F1: 562}
    switch "foo bar" {
    case "":
        v5 -= (v5 << ((4294967295) & 63))
    case "\n\t":
        v5 -= (v4 & v5)
    case "foo bar":
        v9 := GenS2{F0: 505, F1: 163, F2: 4294967295}
        _ = v9
    }
    var v10 int64 = 1
    var v11 float32 = 0.0
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
