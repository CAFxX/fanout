package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 int32 }
type GenS6 struct { F0 int }

func main() {
    var v1 uint8 = 255
    v2 := make(map[string]int)
    v3 := make(map[string]string)
    var v4 uint = 18446744073709551615
    var v5 uint16 = 65535
    v6 := GenS1{F0: false, F1: -6, F2: 0}
    for fi7 := 0; fi7 < 12; fi7++ {
        v1 -= v1
        _ = fi7
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        v3[""] = "foo bar"
        _ = fi8
    }
    var v9 bool = ((^(v5)) > 255)
    v10 := GenS1{F0: v9, F1: -1, F2: 1000}
    v11 := GenS2{F0: -9223372036854775808}
    var v12 uint64 = uint64(v5)
    if true {
        var v13 uint16 = uint16(v12)
        _ = len(v3)
        var v14 int16 = -266
        _ = v13
        _ = v14
    }
    v15 := GenS3{F0: 107, F1: 0}
    v16 := []int{}
    v17 := GenS4{F0: 255}
    if (-1 == 21) {
        if v9 {
            v18 := GenS4{F0: 32767}
            _ = v18
        } else {
            fmt.Println(v12)
            v19 := GenS5{F0: -87, F1: -1}
            if (130 < 613) {
                v4 += (v4 & (v4 / ((1) | 1)))
            }
            _ = v19
        }
        for fi20 := 0; fi20 < 2; fi20++ {
            v12 -= (^(v12))
            v21 := GenS6{F0: fi20}
            v22 := GenS1{F0: true, F1: 332, F2: -1}
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = len(v16)
    } else {
        v23 := GenS2{F0: -859}
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
