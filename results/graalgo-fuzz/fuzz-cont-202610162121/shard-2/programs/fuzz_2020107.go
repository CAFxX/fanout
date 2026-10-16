package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int64 }

func main() {
    var v1 string = "hello"
    v2 := []uint32{0, 187, 2, 409}
    v3 := make(map[int]uint32)
    var v4 bool = ((false || false) && false)
    v5 := make(map[int]bool)
    var v6 uint = 65535
    switch 32768 {
    case 127:
        var v7 uint16 = 881
        _ = v7
    case 128:
        var v8 int16 = -274
        _ = v8
    }
    var v9 uint8 = 0
    v10 := GenS1{F0: 1, F1: (0 <= (v6 / ((746) | 1)))}
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := GenS2{F0: -732}
        _ = fi11
        _ = v12
    }
    var v13 uint = (^(v6))
    v1 = "a"
    for fi14 := 0; fi14 < 14; fi14++ {
        v15 := GenS2{F0: -306}
        v13 = v6
        _ = fi14
        _ = v15
    }
    v16 := GenS2{F0: -1}
    v9 = (v9 ^ v9)
    var v17 string = "\n\t"
    var v18 uint8 = (v9 + v9)
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
