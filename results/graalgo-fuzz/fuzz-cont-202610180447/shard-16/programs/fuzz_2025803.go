package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int8; F1 uint8 }

func main() {
    var v1 bool = false
    v2 := make(map[string]int)
    var v3 uint32 = 65535
    var v4 int16 = 1000
    v3 = (^(v3))
    v5 := GenS1{F0: 0, F1: -1}
    var v6 float32 = 0.0
    if v1 {
        v4 -= -789
    } else {
        var v7 int64 = 1
        _ = v7
    }
    v8 := GenS2{F0: 26, F1: 127}
    var v9 uint8 = 128
    for fi10 := 0; fi10 < 18; fi10++ {
        v4 -= (v4 % ((v4) | 1))
        _ = fi10
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
