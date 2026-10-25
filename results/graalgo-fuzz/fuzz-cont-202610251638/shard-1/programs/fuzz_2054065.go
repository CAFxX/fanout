package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 int8; F2 int8 }
type GenS4 struct { F0 int64; F1 int }
func genFn0(p0 uint64) int32 {
    v5 := []int{-1, 0, 128}
    p0 = p0
    p0 -= p0
    v6 := GenS1{F0: 1, F1: 255, F2: "\n\t"}
    v7 := GenS1{F0: -1, F1: 32767, F2: "a"}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 int8) int64 {
    if false {
        var v8 float64 = -2.25
        var v9 int64 = 0
        fmt.Println(v9, p0)
        _ = v8
        _ = v9
    } else {
        p0 -= (p0 >> ((1) & 63))
    }
    v10 := GenS2{F0: 1000}
    _ = p0
    _ = v10
    return 0
}

func main() {
    var v1 int8 = 42
    var v2 uint32 = 42
    var v3 int32 = 263
    var v4 float32 = -2.25
    v11 := []int{32768}
    fmt.Println(v3, v4, v1)
    if (false || (false || ((v3 % ((v3) | 1)) <= 42))) {
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 uint64 = 4294967295
            v1 = (v1 - v1)
            _ = fi12
            _ = v13
        }
        var v14 int8 = v1
        _ = v14
    }
    var v15 uint8 = 255
    var v16 uint64 = 467
    fmt.Println(v16, v3)
    v17 := GenS3{F0: 1, F1: -1, F2: -128}
    fmt.Println(v1, v15, v1)
    switch 0 {
    case 486:
        switch -437 {
        case 998:
            fmt.Println(v2, v3)
        }
    case -207:
        fmt.Println(v16, v16)
    case -9223372036854775808:
        _ = len(v11)
    }
    fmt.Println(v3)
    for fi18 := 0; fi18 < 17; fi18++ {
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        _ = len(v11)
        _ = fi18
    }
    var v19 int16 = 485
    var v20 int32 = (v3 & (v3 + 798))
    for fi21 := 0; fi21 < 18; fi21++ {
        if false {
            switch (fi21 * fi21) {
            case 2147483647:
                v22 := GenS4{F0: -9223372036854775808, F1: -9223372036854775808}
                _ = v22
            default:
                if len(v11) > 0 {
                    v11[0] = fi21
                }
            }
            v16 -= 918
        }
        _ = fi21
    }
    fmt.Println(v4)
    _ = len(v11)
    if true {
        v15 = (v15 * (v15 ^ v15))
    }
    v16 = v16
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
