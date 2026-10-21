package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint32 }

func main() {
    var v1 uint64 = 452
    v2 := make(map[int]int)
    var v3 uint64 = 280
    var v4 uint64 = (v1 | (^(v3)))
    v5 := make(map[string]uint32)
    var v6 uint16 = 558
    var v7 int8 = -1
    var v8 int64 = -711
    var v9 int32 = 463
    v3 = v3
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 uint8 = 42
        v12 := GenS1{F0: v7, F1: v6, F2: 420}
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: v7, F1: v6, F2: 65535}
    switch 9223372036854775807 {
    case 32767:
        var v14 int8 = -115
        _ = v14
    case -2147483648:
        fmt.Println(v8, v3, v8)
    }
    if true {
        for fi15 := 0; fi15 < 4; fi15++ {
            var v16 uint64 = (v3 / (((v3 % ((v1) | 1))) | 1))
            _ = fi15
            _ = v16
        }
        v7 -= (v7 | v7)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
