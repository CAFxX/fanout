package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint }
type GenS2 struct { F0 int16; F1 string }

func main() {
    var v1 uint64 = 894
    v2 := make(map[int]bool)
    var v3 bool = false
    var v4 int16 = 1
    v5 := []uint32{42, 42, 563}
    v4 = v4
    var v6 bool = (-2.25 >= ((-2.25 * (0.1 + -2.25)) / 1.0))
    var v7 int32 = 42
    v8 := GenS1{F0: 538, F1: 224, F2: 18446744073709551615}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v6, v3)
    if false {
        var v9 int32 = v7
        if (481 >= 4294967295) {
            v4 -= v4
            for fi10 := 0; fi10 < 13; fi10++ {
                v11 := GenS1{F0: 1, F1: 65535, F2: 2}
                v12 := GenS2{F0: v4, F1: "a"}
                _ = fi10
                _ = v11
                _ = v12
            }
        }
        v4 += (v4 ^ v4)
        _ = v9
    }
    v4 = v4
    v3 = false
    switch "" {
    case "":
        _ = len(v5)
    case "a":
        var v13 int = 32768
        _ = v13
    default:
        for fi14 := 0; fi14 < 9; fi14++ {
            for fi15 := 0; fi15 < 3; fi15++ {
                v1 += v1
                var v16 bool = (!(928 > 2))
                _ = fi15
                _ = v16
            }
            v4 += v4
            _ = fi14
        }
    }
    v3 = v6
    if (v4 <= 0) {
        var v17 bool = (!(v3 || v6))
        delete(v2, 9223372036854775807)
        delete(v2, 1)
        _ = v17
    }
    v18 := v2[182]
    v19 := GenS1{F0: 9223372036854775807, F1: 255, F2: 42}
    if v18 {
        var v20 int32 = v7
        _ = v20
    } else {
        v18 = (v7 < v7)
        v6 = v3
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
