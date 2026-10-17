package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int32 }

func main() {
    var v1 int8 = 1
    v2 := make(map[int]string)
    var v3 uint8 = 127
    var v4 int32 = -130
    v5 := GenS1{F0: -214, F1: 493, F2: v4}
    var v6 int32 = 0
    if false {
        v7 := v2[0]
        var v8 float32 = (0.0 - 0.1)
        _ = v7
        _ = v8
    }
    v4 = (v6 << ((574) & 63))
    if false {
        v9 := GenS1{F0: -633, F1: 208, F2: v4}
        var v10 uint64 = 18446744073709551615
        _ = v10
        _ = v9
    }
    var v11 bool = false
    var v12 uint8 = 123
    if v11 {
        var v13 uint16 = 65535
        v14 := []int{32767, -693, 32768}
        _ = v13
        _ = v14
    } else {
        v15 := GenS1{F0: 1, F1: 603, F2: (v6 - (v4 >> ((227) & 63)))}
        var v16 uint16 = 793
        _ = v15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
