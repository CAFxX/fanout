package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 int8 = -77
    var v2 float32 = 3.14159
    v3 := make(map[string]int)
    var v4 int64 = -9223372036854775808
    var v5 int8 = (v1 ^ (v1 >> ((710) & 63)))
    var v6 float32 = v2
    fmt.Println(v1)
    var v7 uint8 = 255
    var v8 float64 = -2.25
    switch 32768 {
    case 0:
        v9 := GenS1{F0: 32767, F1: 2, F2: 175}
        _ = v9
    case 173:
        v10 := GenS2{F0: v4}
        _ = v10
    default:
        v11 := GenS1{F0: 95, F1: uint64(v7), F2: 2}
        _ = v11
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 int32 = 0
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        v7 -= (v7 << ((1) & 63))
        v8 = 0.1
        _ = fi14
    }
    var v15 bool = (65535 > 32767)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
