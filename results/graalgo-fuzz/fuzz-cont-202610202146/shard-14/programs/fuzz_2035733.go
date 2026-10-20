package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint32) uint32 {
    fmt.Println(p0)
    if (true && false) {
        p0 = 42
        var v5 uint8 = 207
        v5 -= v5
        _ = v5
    }
    switch "" {
    case "\n\t":
        if (32767 <= 220) {
            if false {
                v6 := GenS1{F0: p0}
                p0 = (p0 * 994)
                _ = v6
            } else {
                p0 = p0
            }
            var v7 int64 = 1
            _ = v7
        }
    default:
        fmt.Println(p0)
    }
    var v8 uint64 = 173
    for fi9 := 0; fi9 < 12; fi9++ {
        fmt.Println(v8, fi9, v8)
        fmt.Println(p0, fi9, p0)
        v8 += (v8 ^ (^(v8)))
        _ = fi9
    }
    _ = p0
    _ = v8
    return (p0 ^ (^(p0)))
}
func genFn1(p0 uint32) bool {
    v10 := GenS2{F0: 65535}
    var v11 float32 = 1.5
    v12 := GenS1{F0: p0}
    p0 -= p0
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    return true
}

func main() {
    var v1 int32 = 1
    var v2 uint8 = 1
    var v3 int16 = 32767
    v4 := make(map[string]int)
    var v13 float32 = (0.0 + (0.1 / 1.5))
    for fi14 := 0; fi14 < 18; fi14++ {
        v3 = (v3 % (((^(v3))) | 1))
        _ = fi14
    }
    v2 -= (v2 + (v2 << ((65535) & 63)))
    var v15 uint8 = 136
    v16 := GenS2{F0: 32767}
    fmt.Println(v13, v13)
    v2 += v2
    var v17 uint16 = 862
    v18 := GenS1{F0: 105}
    v17 = (v17 ^ v17)
    var v19 uint32 = 325
    v20 := []int{132}
    var v21 int32 = v1
    for fi22 := 0; fi22 < 3; fi22++ {
        var v23 uint16 = v17
        switch -1 {
        case -1:
            var v24 float64 = float64(float32(-2.25))
            _ = v24
        case -9223372036854775808:
            var v25 uint8 = 154
            _ = v25
        case 9223372036854775807:
            fmt.Println(v1, v19)
        default:
            v26 := GenS2{F0: (v17 | v23)}
            _ = v26
        }
        _ = fi22
        _ = v23
    }
    var v27 string = "foo bar"
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
    _ = v4
}
