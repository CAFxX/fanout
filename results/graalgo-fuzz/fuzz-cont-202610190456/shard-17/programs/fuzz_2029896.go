package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int64 }
type GenS2 struct { F0 string }

func main() {
    var v1 float64 = (-2.25 + 0.0)
    v2 := make(map[uint32]int)
    v3 := []int{0, 2147483647, 92}
    v1 += v1
    for fi4 := 0; fi4 < 17; fi4++ {
        var v5 float32 = -2.25
        fmt.Println(v1, v1)
        v6 := GenS1{F0: 4294967295, F1: -128, F2: -586}
        _ = fi4
        _ = v5
        _ = v6
    }
    v1 -= float64(float32(3.14159))
    var v7 uint = 42
    var v8 uint64 = 680
    var v9 float64 = float64(1e10)
    fmt.Println(v1, v9, v9)
    if (!(0 != -32768)) {
        v8 = (v8 & (v8 | v8))
    } else {
        for fi10 := 0; fi10 < 14; fi10++ {
            v7 = 1
            switch (fi10 >> ((0) & 63)) {
            case 385:
                v11 := GenS2{F0: "foo bar"}
                _ = v11
            default:
                _ = len(v2)
            }
            _ = fi10
        }
        switch -2147483648 {
        case -34:
            switch "foo bar" {
            case "foo bar":
                v8 -= (^(v8))
            }
        default:
            v12 := v2[4294967295]
            _ = v12
        }
    }
    var v13 int16 = 492
    switch 4294967295 {
    case 42:
        var v14 uint = (v7 + v7)
        _ = v14
    case -739:
        switch 0 {
        case 662:
            fmt.Println(v8, v13)
        case 149:
            v9 += (v1 + v9)
        case 681:
            v15 := v2[4294967295]
            _ = v15
        }
    }
    var v16 int8 = -1
    var v17 float64 = (-2.25 / 0.1)
    var v18 int16 = (v13 | (v13 % (((-(v13))) | 1)))
    v1 = (-2.25 + 1e10)
    var v19 uint32 = 580
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
