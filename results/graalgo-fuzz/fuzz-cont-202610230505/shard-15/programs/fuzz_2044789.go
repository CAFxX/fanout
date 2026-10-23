package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int16 }
type GenS2 struct { F0 string; F1 uint64 }

func main() {
    var v1 uint16 = 65535
    var v2 float32 = 0.1
    v3 := make(map[string]string)
    var v4 uint32 = 134
    var v5 uint8 = 226
    var v6 bool = (0 <= 32768)
    if (-63 < 1000) {
        v1 = v1
    }
    var v7 uint8 = v5
    var v8 uint16 = 0
    if (32768 > 514) {
        v9 := GenS1{F0: 4294967295, F1: 9223372036854775807, F2: 1000}
        _ = v9
    }
    v5 = (v5 | v5)
    v2 = -2.25
    if (!false) {
        var v10 int32 = 7
        _ = v10
    }
    switch "foo bar" {
    case "hello":
        fmt.Println(v2, v8)
    default:
        delete(v3, "a")
    }
    v2 += ((1e10 * (((1.5 - v2) * (v2 * v2)) / v2)) * float32((((((-2.25 - 0.0) * 0.1) - 0.0) + -2.25) - 3.14159)))
    v11 := GenS2{F0: "", F1: 876}
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 float32 = float32(0.1)
        var v14 int16 = 0
        var v15 int32 = -726
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    switch 128 {
    case 9223372036854775807:
        var v16 uint8 = 0
        _ = v16
    default:
        if false {
            fmt.Println(v7, v8)
        } else {
            v17 := v3[""]
            for fi18 := 0; fi18 < 6; fi18++ {
                switch v17 {
                case "":
                    if v6 {
                        switch fi18 {
                        case 32767:
                            v4 += (v4 | (v4 ^ v4))
                        default:
                            v2 -= v2
                        }
                        v8 += 2
                        var v19 uint8 = v7
                        _ = v19
                    } else {
                        v6 = (v2 <= 0.0)
                    }
                default:
                    for fi20 := 0; fi20 < 15; fi20++ {
                        var v21 int32 = 2147483647
                        _ = fi20
                        _ = v21
                    }
                }
                _ = fi18
            }
            if ((263 == 18446744073709551615) && (2 > 4294967295)) {
                fmt.Println(v6, v2)
            } else {
                fmt.Println(v2)
            }
            _ = v17
        }
    }
    var v22 float64 = 0.0
    fmt.Println(v8, v22, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
