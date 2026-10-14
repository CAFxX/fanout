package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = 0
    var v3 float32 = -2.25
    var v4 int16 = 42
    v3 = 3.14159
    var v5 int8 = (v2 + -110)
    var v6 int16 = v4
    v7 := GenS1{F0: 32767, F1: 255}
    v5 = 100
    v2 -= -128
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 uint64 = 908
        fmt.Println(fi8, v2, v3)
        _ = fi8
        _ = v9
    }
    v6 += (v6 >> ((550) & 63))
    if (0.0 < float32(0.0)) {
        switch ("hello" + "foo bar") {
        case "hello":
            var v10 float64 = 0.0
            _ = v10
        default:
            v1[-513] = 42
        }
        if (!(v4 > (v6 / ((1000) | 1)))) {
            var v11 int16 = v6
            _ = v11
        }
    }
    var v12 uint8 = 1
    fmt.Println(v5, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
