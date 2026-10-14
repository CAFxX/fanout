package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint8; F2 int }

func main() {
    var v1 uint8 = 1
    var v2 int16 = -32768
    var v3 int = -42
    v4 := []uint32{4294967295, 0, 65535}
    v5 := []uint32{65535}
    var v6 float32 = 0.0
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v6 += ((1.5 - 0.0) + (v6 - float32(0.1)))
    var v7 float32 = (v6 * (v6 / 1.0))
    var v8 uint16 = 0
    var v9 int8 = 102
    for fi10 := 0; fi10 < 12; fi10++ {
        if true {
            v11 := GenS1{F0: v9}
            v8 += (v8 % ((255) | 1))
            _ = v11
        } else {
            v6 = v7
            var v12 uint64 = 1
            v12 = (v12 << ((4294967295) & 63))
            _ = v12
        }
        v13 := GenS1{F0: v9}
        _ = fi10
        _ = v13
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        switch fi14 {
        case 4294967295:
            v9 = (-(v9))
        default:
            for fi15 := 0; fi15 < 15; fi15++ {
                v2 = v2
                _ = fi15
            }
        }
        var v16 int32 = -674
        fmt.Println(v7)
        _ = fi14
        _ = v16
    }
    if len(v4) > 0 {
        v4[0] = 112
    }
    if (0 >= 2) {
        var v17 float32 = float32(0.0)
        var v18 int16 = -727
        _ = v17
        _ = v18
    }
    v9 -= (v9 + 51)
    var v19 uint32 = 356
    _ = len(v5)
    var v20 bool = true
    var v21 int = -1
    switch 32768 {
    case 127:
        v6 = (v7 / v7)
    case 864:
        for fi22 := 0; fi22 < 8; fi22++ {
            v9 = v9
            var v23 float64 = 1.5
            fmt.Println(v7, v21, v6)
            _ = fi22
            _ = v23
        }
    }
    v24 := GenS2{F0: ("a" + ("" + ("" + ""))), F1: (v1 | v1), F2: (v21 / (((v3 - (v21 | v3))) | 1))}
    if (v20 && ((v19 << ((2) & 63)) < uint32(v8))) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
