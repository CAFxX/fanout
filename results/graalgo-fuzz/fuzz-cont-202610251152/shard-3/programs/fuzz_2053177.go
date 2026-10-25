package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int8; F1 int; F2 uint }

func main() {
    var v1 float32 = 3.14159
    var v2 int16 = 423
    v3 := make(map[int]int)
    var v4 uint = 1
    var v5 uint8 = 128
    v6 := []int{32767, 42, 0, -2147483648}
    if true {
        v7 := GenS1{F0: -2147483648, F1: 2, F2: v5}
        _ = v7
    } else {
        var v8 float32 = v1
        _ = v8
    }
    fmt.Println(v2, v1, v1)
    v5 = v5
    switch 9223372036854775807 {
    case 127:
        v9 := GenS1{F0: -2147483648, F1: uint16(v5), F2: v5}
        _ = v9
    case 32768:
        if true {
            delete(v3, int(v2))
        } else {
            if (true || true) {
                fmt.Println(v5)
                v10 := GenS2{F0: -127, F1: -2147483648, F2: 832}
                _ = v10
            } else {
                v6 = append(v6, 2147483647)
                for fi11 := 0; fi11 < 15; fi11++ {
                    var v12 float64 = 1e10
                    v2 = -32768
                    v1 -= ((v1 - float32(v12)) - 3.14159)
                    _ = fi11
                    _ = v12
                }
                v13 := GenS2{F0: -1, F1: 32768, F2: (v4 & v4)}
                _ = v13
            }
        }
    case 718:
        v14 := GenS1{F0: -692, F1: 42, F2: v5}
        _ = v14
    default:
        delete(v3, -9223372036854775808)
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        v4 = v4
        v16 := v3[fi15]
        var v17 uint16 = 32767
        _ = fi15
        _ = v16
        _ = v17
    }
    switch 32767 {
    case 9223372036854775807:
        for fi18 := 0; fi18 < 8; fi18++ {
            var v19 uint64 = 2
            var v20 uint8 = (v5 & (v5 & 2))
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    var v21 int8 = 90
    v21 -= (v21 ^ (v21 | v21))
    for fi22 := 0; fi22 < 19; fi22++ {
        switch "hello" {
        case "foo bar":
            fmt.Println(fi22, fi22, v2)
        default:
            fmt.Println(v2, v2)
        }
        if (-313 > -1) {
            v23 := GenS2{F0: 1, F1: fi22, F2: v4}
            fmt.Println(v2)
            _ = v23
        }
        _ = fi22
    }
    v24 := GenS2{F0: v21, F1: 32768, F2: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
