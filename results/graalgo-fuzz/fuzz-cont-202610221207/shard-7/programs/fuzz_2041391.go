package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint }

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 139
    var v3 bool = true
    v4 := make(map[uint32]string)
    v5 := GenS1{F0: 1, F1: false, F2: 4294967295}
    v2 -= v2
    v2 = 0
    var v6 int16 = 1000
    var v7 int32 = int32(v2)
    var v8 int8 = -128
    var v9 int8 = (v8 << ((2) & 63))
    v2 = v2
    for fi10 := 0; fi10 < 6; fi10++ {
        for fi11 := 0; fi11 < 3; fi11++ {
            fmt.Println(v2)
            v2 += (v2 | 0)
            v7 -= (v7 % ((-1) | 1))
            _ = fi11
        }
        _ = fi10
    }
    v9 -= (-(v8))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
