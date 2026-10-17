package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint }

func main() {
    v1 := []uint32{}
    var v2 bool = true
    var v3 uint = 4294967295
    var v4 uint8 = 254
    var v5 float64 = (-2.25 - 3.14159)
    var v6 bool = (55 > 42)
    v5 = (0.1 + 0.0)
    v3 -= v3
    switch 128 {
    case -2147483648:
        fmt.Println(v6, v3, v2)
    case 42:
        var v7 uint16 = 65535
        _ = v7
    case -370:
        if v2 {
            v8 := GenS1{F0: (v3 + (v3 << (((^(v3))) & 63)))}
            _ = v8
        } else {
            v3 = (^(v3))
            v4 = 127
            v9 := GenS1{F0: (v3 ^ (v3 % ((4294967295) | 1)))}
            _ = v9
        }
    default:
        fmt.Println(v2, v4)
    }
    v10 := GenS1{F0: (v3 * v3)}
    v5 = ((v5 - ((v5 / v5) * v5)) * v5)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v11 := GenS1{F0: (v3 ^ 101)}
    var v12 int = 317
    if false {
        if (true && (!(!v6))) {
            for fi13 := 0; fi13 < 1; fi13++ {
                v14 := GenS2{F0: -128, F1: 2}
                _ = fi13
                _ = v14
            }
            var v15 uint32 = 289
            _ = v15
        }
        switch (v12 | (-(v12))) {
        case 2147483647:
            v4 += (v4 - (v4 ^ (^(v4))))
        case -9223372036854775808:
            var v16 uint64 = 1
            _ = v16
        default:
            for fi17 := 0; fi17 < 7; fi17++ {
                var v18 float32 = -2.25
                v2 = true
                _ = fi17
                _ = v18
            }
        }
    }
    var v19 int64 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
