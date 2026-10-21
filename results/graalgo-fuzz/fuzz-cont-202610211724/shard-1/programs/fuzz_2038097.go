package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }

func main() {
    var v1 int32 = -2147483648
    var v2 uint8 = 2
    var v3 uint64 = 228
    var v4 float32 = 0.1
    var v5 int16 = 32767
    var v6 uint8 = v2
    v7 := []int{42, 471, 864, 9223372036854775807}
    v2 = (^(v6))
    var v8 uint32 = 864
    v6 -= v2
    v1 -= -1
    v9 := GenS1{F0: uint64(v8), F1: true}
    var v10 string = "foo bar"
    v4 -= v4
    v4 -= -2.25
    v11 := GenS1{F0: v3, F1: true}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    fmt.Println(v6, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
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
    _ = v9
}
