package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int16 }
type GenS2 struct { F0 int64; F1 bool }

func main() {
    var v1 int32 = 1
    var v2 uint64 = 0
    var v3 int32 = v1
    var v4 int32 = (v3 - 1)
    switch 4294967295 {
    case 561:
        for fi5 := 0; fi5 < 2; fi5++ {
            fmt.Println(fi5, v2)
            _ = fi5
        }
    case 127:
        v6 := GenS1{F0: 100, F1: 1000, F2: 1000}
        _ = v6
    case -644:
        if false {
            v3 = (v1 * v4)
        }
    }
    v1 = v3
    var v7 float64 = 0.1
    if false {
        fmt.Println(v7, v7, v4)
        v2 -= v2
        var v8 int16 = -407
        _ = v8
    } else {
        if false {
            v9 := GenS1{F0: 123, F1: -32768, F2: 0}
            v4 -= (v1 % (((v4 ^ (v1 >> ((0) & 63)))) | 1))
            var v10 int32 = v3
            _ = v10
            _ = v9
        } else {
            var v11 uint = 633
            _ = v11
        }
        var v12 int64 = 0
        v3 = v1
        _ = v12
    }
    v1 -= int32(v2)
    v13 := []int{}
    var v14 int64 = 42
    v4 += v4
    v15 := GenS1{F0: 105, F1: int16(v3), F2: -32768}
    v16 := GenS2{F0: v14, F1: true}
    if false {
        switch "" {
        case "":
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
        case "foo bar":
            var v17 float32 = 1.5
            _ = v17
        case "hello":
            v3 += v1
        default:
            _ = len(v13)
        }
        var v18 string = "\n\t"
        _ = v18
    } else {
        v3 = v4
        fmt.Println(v3, v7, v2)
        var v19 float64 = 1.5
        _ = v19
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        var v21 int32 = 1
        _ = fi20
        _ = v21
    }
    var v22 float64 = v7
    v23 := GenS2{F0: v14, F1: true}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v7
}
