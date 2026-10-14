package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint }
type GenS2 struct { F0 int64; F1 string; F2 string }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint16 = 32767
    var v3 uint8 = 0
    var v4 float32 = -2.25
    var v5 float64 = 1e10
    v4 -= v4
    fmt.Println(v5)
    v6 := []int{-693, 9223372036854775807, 260, 1}
    if (42 != 65535) {
        _ = len(v6)
        v5 += v5
        var v7 int8 = -112
        _ = v7
    }
    _ = len(v6)
    var v8 int = 2147483647
    v4 = float32((v5 + 3.14159))
    _ = len(v6)
    for fi9 := 0; fi9 < 3; fi9++ {
        v1 = (v1 & (v1 * -9223372036854775808))
        var v10 uint64 = 2
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: v1, F1: 4294967295, F2: 152}
    var v12 int8 = 1
    var v13 string = ""
    for fi14 := 0; fi14 < 5; fi14++ {
        v15 := GenS2{F0: v1, F1: (v13 + v13), F2: (v13 + v13)}
        for fi16 := 0; fi16 < 17; fi16++ {
            v2 += 32767
            v12 -= 1
            _ = fi16
        }
        v8 += (fi14 & (fi14 % ((v8) | 1)))
        _ = fi14
        _ = v15
    }
    var v17 int8 = (v12 | (v12 ^ (^(v12))))
    var v18 float32 = (v4 * (v4 - v4))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
