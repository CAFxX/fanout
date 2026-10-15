package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 float64 = 1.5
    v2 := []uint32{813, 175, 42, 907}
    var v3 uint32 = 42
    var v4 int16 = 32767
    v5 := make(map[string]bool)
    v3 = (v3 & (^(v3)))
    v6 := GenS1{F0: 167}
    v1 = v1
    var v7 int8 = -37
    var v8 float32 = 0.0
    v3 = (^(v3))
    switch "hello" {
    case "hello":
        if len(v2) > 0 {
            v2[0] = v3
        }
    }
    if (true && ((v1 * ((0.1 * float64(v8)) * 1e10)) <= v1)) {
        var v9 int64 = -9223372036854775808
        if false {
            v4 -= v4
            var v10 float64 = v1
            _ = v10
        }
        v7 = 127
        _ = v9
    } else {
        v8 = v8
        if true {
            var v11 float32 = (-2.25 - v8)
            if ((v7 >> ((18446744073709551615) & 63)) < (v7 | v7)) {
                _ = len(v2)
                v12 := GenS1{F0: 128}
                v7 += (-(v7))
                _ = v12
            }
            _ = v11
        } else {
            switch 32768 {
            case -295:
                v4 += v4
            case -61:
                var v13 uint16 = 2
                _ = v13
            case 2147483647:
                for fi14 := 0; fi14 < 17; fi14++ {
                    v8 = v8
                    _ = fi14
                }
            default:
                fmt.Println(v8, v8, v7)
            }
        }
    }
    var v15 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
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
