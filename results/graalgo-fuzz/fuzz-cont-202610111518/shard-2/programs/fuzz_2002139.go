package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 uint32; F2 string }
type GenS3 struct { F0 int32; F1 uint8; F2 string }

func main() {
    var v1 int8 = 71
    var v2 uint = 921
    var v3 int32 = 0
    var v4 uint = 255
    var v5 int64 = -752
    v6 := []int{9223372036854775807, 631, 2147483647, 1}
    var v7 uint32 = 995
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    switch 127 {
    case 751:
        if ((((1e10 * ((-2.25 - 1e10) / -2.25)) / (1.5 - 0.0)) - -2.25) > 3.14159) {
            for fi8 := 0; fi8 < 11; fi8++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                fmt.Println(v7, v5, v3)
                _ = fi8
            }
            v6 = append(v6, -9223372036854775808)
            var v9 float64 = ((float64(float32(1e10)) - (float64(1.5) + (-2.25 + 0.1))) - 1.5)
            _ = v9
        } else {
            switch -255 {
            case -1:
                fmt.Println(v2, v3)
            case 128:
                v7 = v7
            default:
                for fi10 := 0; fi10 < 7; fi10++ {
                    if false {
                        fmt.Println(v4, v1, fi10)
                        v11 := GenS1{F0: 15}
                        _ = v11
                    } else {
                        switch fi10 {
                        case 4294967295:
                            v2 += (^(v4))
                        default:
                            var v12 float32 = (1.5 + (1.5 + ((0.0 / -2.25) + -2.25)))
                            _ = v12
                        }
                    }
                    var v13 int16 = 42
                    _ = fi10
                    _ = v13
                }
            }
            v4 = (v2 >> ((v2) & 63))
        }
    }
    v3 = (v3 % ((v3) | 1))
    for fi14 := 0; fi14 < 16; fi14++ {
        _ = len(v6)
        v3 = v3
        _ = fi14
    }
    v4 -= (v2 >> (((v4 ^ v2)) & 63))
    fmt.Println(v1, v4)
    fmt.Println(v7, v3, v4)
    var v15 uint64 = 551
    var v16 uint32 = (v7 / ((280) | 1))
    v17 := GenS1{F0: 0}
    v18 := GenS2{F0: 202, F1: v7, F2: "hello"}
    v19 := GenS2{F0: uint16(v2), F1: v7, F2: "\n\t"}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v20 := GenS3{F0: v3, F1: 2, F2: ""}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
