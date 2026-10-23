package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 bool }

func main() {
    var v1 float64 = (float64(1.5) + (0.0 / ((float64(0.1) + -2.25) - 1.5)))
    var v2 int32 = 0
    var v3 float64 = (1.5 * v1)
    v4 := make(map[int]uint32)
    var v5 bool = true
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 int8 = 42
        _ = fi6
        _ = v7
    }
    v8 := []int{}
    v9 := GenS1{F0: false, F1: "foo bar"}
    var v10 int64 = 1
    switch "hello" {
    case "foo bar":
        delete(v4, -312)
    }
    v1 = v3
    v11 := GenS2{F0: 576, F1: (v2 % (((v2 % ((v2) | 1))) | 1))}
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 uint8 = 222
        v4[fi12] = 610
        var v14 float64 = (float64((float32(1e10) - (float32(float64(float32(v3))) - 3.14159))) - (v1 * v3))
        _ = fi12
        _ = v13
        _ = v14
    }
    if ((2 == 255) && v5) {
        if true {
            v2 += (v2 & (-(v2)))
            fmt.Println(v3, v3)
            if false {
                v2 -= 352
            } else {
                fmt.Println(v3)
                if (-345 == -124) {
                    v3 = -2.25
                }
                if false {
                    for fi15 := 0; fi15 < 14; fi15++ {
                        v3 += v1
                        fmt.Println(v3, fi15, v2)
                        _ = fi15
                    }
                    fmt.Println(v10, v10, v5)
                    if (!false) {
                        v3 -= v3
                    }
                }
            }
        }
    }
    var v16 int8 = 100
    delete(v4, 2147483647)
    if v5 {
        v3 += v3
    } else {
        fmt.Println(v1, v16)
        fmt.Println(v1, v2)
        v8 = append(v8, 42)
    }
    v17 := GenS3{F0: v5}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
