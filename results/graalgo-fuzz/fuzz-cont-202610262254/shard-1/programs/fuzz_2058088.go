package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 float64 = 1.5
    var v3 int8 = 100
    var v4 float64 = float64(3.14159)
    v5 := make(map[string]uint32)
    v6 := GenS1{F0: 0, F1: -32768}
    v7 := GenS1{F0: 127, F1: 153}
    var v8 uint64 = 136
    if (true || (true || ((((^(v8)) == (v8 - 2)) && ((938 != 1) && true)) || (92 >= 127)))) {
        var v9 int16 = -683
        v3 -= v3
        v3 = 127
        _ = v9
    } else {
        var v10 int = 520
        fmt.Println(v8, v10)
        _ = v10
    }
    var v11 int = -9223372036854775808
    v11 -= (v11 >> ((65535) & 63))
    v12 := GenS1{F0: 2, F1: -32768}
    if (244 <= 255) {
        for fi13 := 0; fi13 < 16; fi13++ {
            v4 = v2
            _ = fi13
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
