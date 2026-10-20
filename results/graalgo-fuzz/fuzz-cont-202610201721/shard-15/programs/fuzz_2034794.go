package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int64; F1 bool }
func genFn0() uint8 {
    var v6 int = 1
    var v7 int16 = -546
    v8 := GenS1{F0: 42, F1: -97}
    v9 := GenS2{F0: -181, F1: true}
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 128
}

func main() {
    var v1 float64 = float64((1e10 - 0.1))
    var v2 float64 = 1.5
    var v3 float64 = 0.0
    v4 := []uint32{4294967295, 65535, 145, 804}
    var v5 uint = 65535
    switch -1 {
    case -1:
        var v10 int16 = -32768
        _ = v10
    case 0:
        var v11 int16 = 0
        _ = v11
    }
    switch 9223372036854775807 {
    case 42:
        var v12 uint8 = 127
        _ = v12
    case 32768:
        fmt.Println(v5)
    case -832:
        v5 = 65535
    default:
        if len(v4) > 0 {
            v4[0] = 824
        }
    }
    var v13 int32 = -2147483648
    if true {
        var v14 int64 = 786
        if (genFn0() < 255) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        } else {
            var v15 int16 = -32768
            _ = v15
        }
        var v16 int8 = -128
        _ = v14
        _ = v16
    } else {
        v17 := GenS1{F0: 42, F1: 127}
        for fi18 := 0; fi18 < 16; fi18++ {
            fmt.Println(v3)
            _ = fi18
        }
        _ = v17
    }
    v3 -= (1e10 * v2)
    switch "a" {
    case "hello":
        v3 -= 1.5
    case "a":
        _ = len(v4)
    default:
        v5 = (^(v5))
    }
    if len(v4) > 0 {
        v4[0] = 0
    }
    v19 := GenS2{F0: -654, F1: (false || false)}
    fmt.Println(v13)
    v2 += 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
