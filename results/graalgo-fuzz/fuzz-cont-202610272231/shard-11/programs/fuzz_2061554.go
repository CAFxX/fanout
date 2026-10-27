package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int64 }
type GenS2 struct { F0 uint8; F1 uint16; F2 bool }

func main() {
    var v1 int64 = 1
    var v2 float32 = 1e10
    v3 := make(map[int]bool)
    var v4 uint64 = 96
    v5 := make(map[string]string)
    var v6 int16 = 42
    v7 := []int{32768, -162, -1}
    v1 += v1
    v6 = (v6 + (v6 & 32767))
    var v8 int = -9223372036854775808
    if false {
        var v9 uint64 = v4
        _ = v9
    }
    var v10 string = ""
    fmt.Println(v8)
    var v11 float32 = v2
    for fi12 := 0; fi12 < 9; fi12++ {
        v1 = -1
        fmt.Println(v4)
        if ((!((v2 / v2) >= -2.25)) && false) {
            v6 = 28
        }
        _ = fi12
    }
    if false {
        for fi13 := 0; fi13 < 6; fi13++ {
            v6 = (v6 << ((652) & 63))
            for fi14 := 0; fi14 < 18; fi14++ {
                v15 := GenS1{F0: fi13, F1: -2147483648, F2: (v1 >> ((18446744073709551615) & 63))}
                _ = fi14
                _ = v15
            }
            fmt.Println(v10, v1, v6)
            _ = fi13
        }
        v1 = (-(v1))
        v16 := GenS1{F0: v8, F1: (v8 / ((467) | 1)), F2: (v1 % ((v1) | 1))}
        _ = v16
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        for fi18 := 0; fi18 < 8; fi18++ {
            v2 -= ((v11 / float32(0.1)) / (1.5 * (v2 * float32((((-2.25 * 0.0) + 0.0) - (0.1 - (float64(0.0) / (-2.25 - float64(v2)))))))))
            if len(v7) > 0 {
                v7[0] = (fi17 ^ (fi18 + fi17))
            }
            var v19 int16 = -32768
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        var v21 bool = true
        var v22 float64 = 1e10
        _ = fi20
        _ = v21
        _ = v22
    }
    v3[(v8 | v8)] = true
    v23 := GenS2{F0: 1, F1: uint16(v4), F2: (!false)}
    var v24 int = v8
    var v25 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
