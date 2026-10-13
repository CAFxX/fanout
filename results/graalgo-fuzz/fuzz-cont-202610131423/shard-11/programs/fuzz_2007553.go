package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int8 }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []int{-170, -473}
    var v3 float32 = 1e10
    var v4 uint64 = 18446744073709551615
    var v5 uint = 65535
    if true {
        switch "\n\t" {
        case "\n\t":
            v5 -= (^(v5))
        default:
            _ = len(v2)
        }
    }
    if true {
        v1 = (v1 / ((v1) | 1))
    }
    for fi6 := 0; fi6 < 17; fi6++ {
        _ = len(v2)
        _ = fi6
    }
    switch 2147483647 {
    case -954:
        v7 := GenS1{F0: 32767}
        _ = v7
    case -9223372036854775808:
        if ((v4 - v4) == v4) {
            v8 := GenS1{F0: 128}
            _ = v8
        }
    case 42:
        var v9 float64 = 0.1
        _ = v9
    default:
        fmt.Println(v4, v4)
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 float64 = float64((0.1 - v3))
        if true {
            var v12 uint32 = 122
            fmt.Println(v5, fi10, v3)
            var v13 uint64 = uint64(v12)
            _ = v12
            _ = v13
        } else {
            if (!false) {
                var v14 int32 = 1
                _ = v14
            } else {
                v3 -= (1e10 + ((1e10 + 0.1) + (v3 - ((v3 + v3) - ((v3 * (float32((float64(((v3 / float32(v11)) * 1.5)) - 3.14159)) - 3.14159)) + (((float32(((v11 + float64(float32(v11))) / v11)) * (0.0 * v3)) / 1.0) - -2.25))))))
                switch "" {
                case "\n\t":
                    fmt.Println(v3, v4, v11)
                default:
                    v1 += v1
                }
                if true {
                    var v15 uint16 = 42
                    _ = v15
                } else {
                    for fi16 := 0; fi16 < 14; fi16++ {
                        v11 = 1e10
                        _ = fi16
                    }
                    v17 := GenS2{F0: "\n\t", F1: 0}
                    _ = v17
                }
            }
        }
        var v18 float32 = float32(v11)
        _ = fi10
        _ = v11
        _ = v18
    }
    fmt.Println(v5, v3, v1)
    var v19 int = 522
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) > 0 {
        v2[0] = (v19 + 127)
    }
    fmt.Println(v4)
    var v20 int32 = 0
    var v21 int16 = 42
    if len(v2) > 0 {
        v2[0] = 115
    }
    var v22 int = v19
    v23 := GenS2{F0: ("hello" + "hello"), F1: -18}
    v24 := GenS1{F0: v19}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
