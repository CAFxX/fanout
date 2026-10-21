package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 bool }

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = -102
    var v3 int64 = -9223372036854775808
    var v4 uint = 2
    v4 = (v4 << (((^(v4))) & 63))
    if true {
        fmt.Println(v3)
        var v5 float32 = 3.14159
        var v6 int64 = (v3 << ((v4) & 63))
        _ = v5
        _ = v6
    } else {
        var v7 float32 = (0.1 / 3.14159)
        _ = v7
    }
    v3 = v2
    v4 -= v4
    fmt.Println(v3, v2)
    v8 := GenS1{F0: 431, F1: (127 <= 42)}
    var v9 uint16 = 255
    v10 := GenS1{F0: 756, F1: true}
    var v11 int16 = -32768
    v12 := GenS2{F0: true}
    v13 := GenS1{F0: 2, F1: (v9 <= (v9 << ((v4) & 63)))}
    var v14 uint16 = v9
    v9 -= (v9 + (v14 / (((^(v14))) | 1)))
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
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
