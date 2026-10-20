package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 string }

func main() {
    var v1 uint16 = 1
    v2 := make(map[int]string)
    var v3 float32 = 0.0
    var v4 uint64 = 881
    var v5 uint8 = 23
    var v6 uint32 = 499
    v5 = (^(v5))
    v7 := []int{}
    v4 = (v4 / (((v4 * v4)) | 1))
    v8 := GenS1{F0: 42}
    switch 2147483647 {
    case 107:
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 float32 = (-2.25 * (v3 - float32(1.5)))
            v11 := GenS2{F0: 4294967295, F1: v4}
            _ = fi9
            _ = v10
            _ = v11
        }
    case 4294967295:
        fmt.Println(v5)
    case -765:
        v12 := GenS3{F0: "foo bar"}
        _ = v12
    }
    switch 127 {
    case 128:
        v1 += (v1 * 255)
    }
    var v13 int16 = 0
    v14 := GenS2{F0: (v6 % ((v6) | 1)), F1: (v4 >> ((255) & 63))}
    var v15 int64 = 42
    switch 32767 {
    case 6:
        v15 -= (v15 ^ v15)
    default:
        v6 -= (v6 & (v6 % ((65535) | 1)))
    }
    if (0.1 > 1e10) {
        var v16 int64 = (^(v15))
        _ = v16
    }
    var v17 bool = false
    var v18 int64 = -9223372036854775808
    var v19 uint16 = v1
    fmt.Println(v6, v1)
    v17 = ("" <= "foo bar")
    switch 1 {
    case -9223372036854775808:
        v18 += (v18 / (((v18 & 42)) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v17
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
