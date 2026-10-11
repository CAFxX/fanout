package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int; F1 bool; F2 int16 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 int32 = 42
    v2 := []int{-684}
    var v3 float32 = (-2.25 - 0.1)
    var v4 uint16 = 255
    v5 := []int{-683, 728, -9223372036854775808, -1}
    var v6 string = "hello"
    switch 4294967295 {
    case 1:
        var v7 uint16 = v4
        _ = v7
    default:
        if len(v2) > 0 {
            v2[0] = 0
        }
    }
    var v8 int16 = 503
    v5 = append(v5, -1)
    fmt.Println(v3)
    if true {
        for fi9 := 0; fi9 < 10; fi9++ {
            v6 = "hello"
            _ = fi9
        }
        fmt.Println(v8, v3, v6)
    } else {
        var v10 uint8 = 156
        var v11 int64 = -777
        _ = v10
        _ = v11
    }
    v4 = (v4 % (((v4 * (v4 * v4))) | 1))
    v12 := GenS1{F0: 0, F1: (^(v1))}
    v1 = (v1 | (v1 ^ (v1 & 2147483647)))
    v13 := GenS2{F0: 32768, F1: ("\n\t" < "hello"), F2: v8}
    fmt.Println(v1, v6)
    for fi14 := 0; fi14 < 12; fi14++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v3 += (-2.25 * (v3 / v3))
        _ = fi14
    }
    var v15 float64 = 0.1
    v16 := GenS3{F0: (v8 >> ((755) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
