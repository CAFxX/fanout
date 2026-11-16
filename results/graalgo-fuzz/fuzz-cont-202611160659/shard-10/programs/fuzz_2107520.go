package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int8 }
type GenS2 struct { F0 int32; F1 uint16; F2 uint8 }
func genFn0() uint8 {
    v5 := GenS1{F0: 127, F1: -2147483648, F2: 1}
    var v6 uint32 = 65535
    v6 -= (v6 - v6)
    v6 -= v6
    _ = v5
    _ = v6
    return 228
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{-9223372036854775808, 1}
    var v3 bool = true
    v4 := make(map[uint32]bool)
    if false {
        v3 = (1 <= 998)
        var v7 uint8 = 248
        _ = v7
    }
    v8 := GenS1{F0: genFn0(), F1: -579, F2: 127}
    v9 := GenS1{F0: 128, F1: -2147483648, F2: -26}
    if v3 {
        fmt.Println(v3)
        v10 := GenS2{F0: -930, F1: 32767, F2: 161}
        _ = v10
    }
    var v11 uint16 = 42
    var v12 int = 1
    var v13 float32 = (float32(1e10) / 3.14159)
    v14 := GenS1{F0: 127, F1: -2147483648, F2: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
