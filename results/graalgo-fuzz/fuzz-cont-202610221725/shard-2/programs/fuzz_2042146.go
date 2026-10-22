package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    var v4 int16 = -32768
    _ = len(v2)
    v5 := GenS1{F0: -792, F1: -107}
    if false {
        fmt.Println(v4)
        var v6 uint = 233
        _ = v6
    }
    v4 = v4
    if true {
        v4 -= (v4 + v4)
    } else {
        v4 -= (v4 / ((v4) | 1))
    }
    if false {
        fmt.Println(v4)
        if ((0.1 / float64((float32(0.1) / 0.1))) >= 3.14159) {
            v4 = v4
        }
        v7 := GenS1{F0: (v4 & (v4 / (((v4 & (^(v4)))) | 1))), F1: 1}
        _ = v7
    } else {
        fmt.Println(v4)
        if false {
            v4 = 42
            v4 += 1000
        }
        v4 = 1
    }
    v4 -= v4
    fmt.Println(v4)
    var v8 uint = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
