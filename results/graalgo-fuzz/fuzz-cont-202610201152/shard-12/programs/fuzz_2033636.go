package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int16 }
type GenS2 struct { F0 bool }

func main() {
    var v1 uint64 = 160
    var v2 int8 = 0
    var v3 uint = 1
    v4 := GenS1{F0: v2, F1: 2, F2: 42}
    v3 = v3
    var v5 uint8 = 0
    v2 -= (v2 * 1)
    v6 := GenS1{F0: v2, F1: (^(v5)), F2: -270}
    var v7 uint = 486
    fmt.Println(v5)
    if false {
        v8 := GenS2{F0: (v1 > (v1 | 1))}
        v1 = 976
        v2 -= (v2 >> ((65535) & 63))
        _ = v8
    }
    v7 -= (v3 - (^(v7)))
    if (!((v2 % ((v2) | 1)) != 42)) {
        var v9 uint = v3
        _ = v9
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        v11 := GenS2{F0: (!(!(-2.25 != (((0.1 * float64(float32(0.1))) - 1.5) - -2.25))))}
        var v12 float32 = float32(1e10)
        v13 := GenS2{F0: false}
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
