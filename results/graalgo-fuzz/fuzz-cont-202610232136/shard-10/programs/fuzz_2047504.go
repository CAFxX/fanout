package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint8 }
func genFn0() uint {
    var v4 uint32 = 869
    v5 := GenS1{F0: -13, F1: 779}
    v6 := []int{1, 50}
    var v7 int32 = 0
    v4 -= 469
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 701
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = 0
    var v3 uint8 = 248
    v8 := []int{-2147483648, -289, 0, 663}
    v2 = 0
    for fi9 := 0; fi9 < 2; fi9++ {
        fmt.Println(v2, v3)
        _ = fi9
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        if (false || false) {
            v2 -= -9223372036854775808
            v2 -= (v2 & 440)
            v8 = append(v8, fi10)
        } else {
            v11 := GenS1{F0: 127, F1: 439}
            _ = v11
        }
        for fi12 := 0; fi12 < 8; fi12++ {
            switch (fi12 + fi12) {
            case -2147483648:
                v2 -= (v2 / ((v2) | 1))
            case -9223372036854775808:
                if true {
                    var v13 int16 = -32768
                    var v14 uint = 43
                    _ = v13
                    _ = v14
                }
            default:
                v2 = (v2 - v2)
            }
            _ = fi12
        }
        _ = fi10
    }
    if (!false) {
        v8 = append(v8, -70)
    } else {
        if len(v8) > 0 {
            v8[0] = 32768
        }
        v3 = (v3 >> ((812) & 63))
    }
    var v15 int32 = 42
    var v16 int8 = 127
    if (1 >= 897) {
        v17 := GenS1{F0: v16, F1: 1}
        var v18 int8 = (v16 * (v16 | 88))
        _ = v17
        _ = v18
    }
    v16 += (v16 - (v16 | (v16 >> ((0) & 63))))
    v19 := GenS2{F0: 2, F1: (v3 >> ((65535) & 63))}
    v8 = append(v8, 128)
    fmt.Println(v2, v16)
    var v20 uint16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v8
}
