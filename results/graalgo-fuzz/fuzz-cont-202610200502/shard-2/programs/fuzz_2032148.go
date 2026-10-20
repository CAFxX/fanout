package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32 }

func main() {
    v1 := []string{}
    var v2 int8 = 127
    var v3 int16 = -32768
    var v4 uint16 = 42
    v5 := []int{9223372036854775807, 232, 32767, 0}
    fmt.Println(v3, v2, v2)
    v6 := GenS1{F0: (v2 * v2), F1: -1}
    v4 = 478
    var v7 float64 = 0.0
    if true {
        v3 = -550
    } else {
        var v8 uint8 = 69
        v9 := GenS1{F0: (v2 - v2), F1: 42}
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        switch "\n\t" {
        case "\n\t":
            var v11 int8 = 15
            _ = v11
        case "hello":
            v7 -= (0.0 + v7)
        default:
            v1 = append(v1, (("" + "foo bar") + "a"))
        }
        v2 -= (v2 / (((v2 / ((100) | 1))) | 1))
        fmt.Println(v7, fi10)
        _ = fi10
    }
    var v12 int = -2147483648
    for fi13 := 0; fi13 < 8; fi13++ {
        v14 := GenS2{F0: v3}
        _ = fi13
        _ = v14
    }
    switch v12 {
    case 315:
        var v15 uint64 = 728
        _ = v15
    case -249:
        if (42 < 1) {
            if len(v5) > 0 {
                v5[0] = (v12 + int(v3))
            }
        } else {
            switch "a" {
            case "":
                if true {
                    v2 = (v2 * (v2 ^ -128))
                    var v16 float32 = float32(v7)
                    v17 := GenS2{F0: (v3 << ((65535) & 63))}
                    _ = v16
                    _ = v17
                }
            }
            if (v2 == v2) {
                v18 := GenS3{F0: 2147483647}
                _ = v18
            }
        }
    case 9223372036854775807:
        v4 = (v4 % ((v4) | 1))
    default:
        v19 := GenS1{F0: v2, F1: -626}
        _ = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
