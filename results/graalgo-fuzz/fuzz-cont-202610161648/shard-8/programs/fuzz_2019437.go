package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint8 }
type GenS2 struct { F0 int; F1 string; F2 uint }

func main() {
    var v1 uint = 427
    v2 := make(map[int]int)
    v3 := []uint32{4294967295, 1, 65535}
    var v4 uint = v1
    switch -482 {
    case 128:
        var v5 int32 = 331
        _ = v5
    }
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 uint = (v1 >> (((^(v4))) & 63))
        v8 := GenS1{F0: -9223372036854775808, F1: 521, F2: 94}
        if (0 < 1) {
            var v9 uint8 = 127
            v4 = v4
            _ = v9
        } else {
            v10 := GenS1{F0: 42, F1: 383, F2: 127}
            _ = v10
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    if false {
        v11 := GenS1{F0: -1, F1: 0, F2: 128}
        switch 837 {
        case 9223372036854775807:
            fmt.Println(v1, v4)
        case 775:
            var v12 bool = ((0.0 != 0.0) || false)
            _ = v12
        case -599:
            if false {
                v4 += (v1 + v4)
                var v13 int32 = 1
                _ = v13
            } else {
                fmt.Println(v4)
                v1 -= (v4 | (v1 | (v1 / ((1) | 1))))
            }
        default:
            switch "\n\t" {
            case "hello":
                var v14 int32 = -2147483648
                _ = v14
            default:
                _ = len(v2)
            }
        }
        _ = v11
    }
    fmt.Println(v4, v4)
    v15 := GenS1{F0: 42, F1: 42, F2: 1}
    var v16 int8 = 42
    var v17 bool = false
    fmt.Println(v17, v1, v1)
    v18 := GenS2{F0: 9223372036854775807, F1: (("\n\t" + "foo bar") + (("\n\t" + "a") + ("hello" + "a"))), F2: (^(v1))}
    fmt.Println(v4, v17, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
