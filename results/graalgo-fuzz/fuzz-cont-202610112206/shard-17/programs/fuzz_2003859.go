package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint16 }

func main() {
    var v1 int32 = -2147483648
    var v2 uint8 = 193
    var v3 bool = false
    v4 := make(map[int]string)
    var v5 int32 = (v1 ^ v1)
    v6 := make(map[uint32]string)
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 int8 = 112
        _ = fi7
        _ = v8
    }
    v9 := []int{-535, 4294967295, -988}
    if len(v9) > 0 {
        v9[0] = 9223372036854775807
    }
    v1 -= (-(v5))
    switch "a" {
    case "foo bar":
        v10 := GenS1{F0: v1}
        _ = v10
    case "\n\t":
        var v11 int16 = 1000
        _ = v11
    case "hello":
        v12 := GenS2{F0: 42, F1: 10, F2: 32767}
        _ = v12
    default:
        var v13 int = -115
        _ = v13
    }
    if (!false) {
        fmt.Println(v5)
        var v14 float32 = 1e10
        _ = v14
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        v1 -= (^(v5))
        for fi16 := 0; fi16 < 3; fi16++ {
            v4[fi16] = ("" + ("\n\t" + ("hello" + "foo bar")))
            fmt.Println(v1, v1, v1)
            fmt.Println(v5, v3, v1)
            _ = fi16
        }
        _ = fi15
    }
    fmt.Println(v2, v3)
    delete(v6, 79)
    if len(v9) > 0 {
        v9[0] = 92
    }
    var v17 int = -174
    v18 := v4[v17]
    v9 = append(v9, (v17 * -449))
    var v19 bool = (true || ((0.0 <= 1e10) || ((3.14159 + -2.25) != 0.0)))
    for fi20 := 0; fi20 < 9; fi20++ {
        _ = len(v6)
        _ = fi20
    }
    if len(v9) > 0 {
        v9[0] = (v17 & 1)
    }
    var v21 int32 = (v5 % (((-(v5))) | 1))
    v5 += (v1 % ((v1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
