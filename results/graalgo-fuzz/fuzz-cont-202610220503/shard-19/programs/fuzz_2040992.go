package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int8 }

func main() {
    var v1 int16 = -32768
    var v2 uint8 = 127
    var v3 int8 = -1
    var v4 float32 = 3.14159
    v5 := make(map[int]uint32)
    var v6 int8 = (^(v3))
    if true {
        v1 = (v1 + 42)
    } else {
        var v7 int64 = 42
        v8 := []int{127}
        _ = v7
        _ = v8
    }
    var v9 uint64 = 18446744073709551615
    var v10 int = 1
    v11 := []int{(-(v10)), (v10 << ((0) & 63))}
    for fi12 := 0; fi12 < 17; fi12++ {
        fmt.Println(v9)
        var v13 uint64 = v9
        _ = fi12
        _ = v13
    }
    _ = len(v5)
    v1 = v1
    v9 = 294
    for fi14 := 0; fi14 < 13; fi14++ {
        v15 := GenS1{F0: 65535, F1: 42, F2: (v6 / (((-(v3))) | 1))}
        var v16 int64 = -9223372036854775808
        _ = fi14
        _ = v15
        _ = v16
    }
    if len(v11) > 0 {
        v11[0] = (v10 << ((uint(v2)) & 63))
    }
    v2 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
