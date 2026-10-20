package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint16; F1 string; F2 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 uint }

func main() {
    var v1 uint32 = 1
    v2 := make(map[string]string)
    v3 := make(map[uint32]int)
    var v4 float64 = (float64(1.5) - float64(0.0))
    v5 := []uint32{0, (v1 / (((v1 & 2)) | 1)), (v1 - v1)}
    switch 52 {
    case 603:
        if (227 > 156) {
            _ = len(v5)
            var v6 float64 = v4
            v6 -= 3.14159
            _ = v6
        }
    case 0:
        _ = len(v3)
    case 863:
        if false {
            if true {
                fmt.Println(v4)
                for fi7 := 0; fi7 < 9; fi7++ {
                    v8 := GenS1{F0: uint64(v1), F1: 42, F2: 1}
                    var v9 float32 = 0.0
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                var v10 uint8 = 2
                _ = v10
            } else {
                var v11 int = 0
                v12 := GenS2{F0: 204, F1: "a", F2: v11}
                _ = len(v5)
                _ = v11
                _ = v12
            }
            _ = len(v5)
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    default:
        if true {
            v4 = 0.0
            v4 = v4
        }
    }
    var v13 uint32 = (v1 % ((v1) | 1))
    if false {
        v14 := GenS3{F0: (0 <= 0)}
        if ((0.1 == 0.0) || true) {
            v5 = append(v5, (v1 - v1))
            var v15 float32 = 0.0
            v5 = append(v5, (v13 / ((42) | 1)))
            _ = v15
        } else {
            var v16 uint64 = 42
            _ = v16
        }
        switch 32767 {
        case 1:
            var v17 string = "foo bar"
            _ = v17
        }
        _ = v14
    } else {
        fmt.Println(v1, v13, v13)
        v13 = v1
    }
    var v18 uint64 = 2
    var v19 int64 = 9223372036854775807
    var v20 bool = true
    for fi21 := 0; fi21 < 12; fi21++ {
        var v22 uint32 = v1
        fmt.Println(v19)
        _ = fi21
        _ = v22
    }
    if (v20 && v20) {
        var v23 float64 = 0.0
        if v20 {
            var v24 bool = (100 == 127)
            v20 = (false && (!v24))
            _ = len(v2)
            _ = v24
        } else {
            fmt.Println(v13, v20, v19)
        }
        _ = v23
    }
    v25 := GenS4{F0: 940, F1: uint(v13)}
    fmt.Println(v18, v20, v13)
    fmt.Println(v20, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
