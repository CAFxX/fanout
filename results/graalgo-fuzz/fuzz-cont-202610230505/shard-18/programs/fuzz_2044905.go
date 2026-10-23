package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint16 }

func main() {
    var v1 int8 = -128
    v2 := []string{("" + "a"), "", ""}
    var v3 bool = (715 != 289)
    var v4 uint16 = 441
    v4 = v4
    var v5 float64 = (((-2.25 - (((0.1 - (1.5 - float64((1.5 * 0.1)))) + (0.0 + -2.25)) * 1e10)) - (0.0 + 1.5)) + float64(-2.25))
    v4 -= (v4 << ((42) & 63))
    fmt.Println(v1, v1)
    if false {
        var v6 int8 = (^(v1))
        v4 = (v4 / ((v4) | 1))
        _ = v6
    }
    if v3 {
        var v7 uint32 = 65535
        var v8 float32 = (0.0 * 0.1)
        _ = v7
        _ = v8
    }
    if v3 {
        var v9 uint8 = 42
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v9
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if len(v2) > 0 {
            v2[0] = ""
        }
        if (v3 && (!(!v3))) {
            if v3 {
                if true {
                    fmt.Println(v4, v1, v3)
                    var v10 int = 4294967295
                    _ = v10
                }
                var v11 float32 = 3.14159
                _ = v11
            } else {
                v12 := GenS1{F0: 32767, F1: v4}
                var v13 bool = v3
                _ = v12
                _ = v13
            }
            var v14 float64 = 0.0
            _ = v14
        } else {
            if len(v2) > 0 {
                v2[0] = ("" + "a")
            }
            if (int16(v1) != int16(v4)) {
                var v15 float64 = v5
                var v16 int16 = -545
                var v17 uint = 630
                _ = v15
                _ = v16
                _ = v17
            }
        }
    }
    switch 245 {
    case 715:
        if (42 > 2) {
            switch 737 {
            case 1:
                var v18 bool = v3
                _ = v18
            }
            v19 := GenS1{F0: 42, F1: (v4 & 255)}
            v20 := GenS1{F0: 42, F1: (v4 / ((336) | 1))}
            _ = v19
            _ = v20
        }
    default:
        fmt.Println(v5, v1)
    }
    v21 := GenS2{F0: (v4 | 255), F1: v1, F2: v4}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
