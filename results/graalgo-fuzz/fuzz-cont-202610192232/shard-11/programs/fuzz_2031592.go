package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int64; F1 uint64; F2 uint }
type GenS3 struct { F0 int; F1 int32; F2 int64 }
func genFn0(p0 uint, p1 uint32, p2 int64) uint8 {
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 float64 = (((3.14159 * 0.0) - float64(float32(3.14159))) + float64(0.0))
        for fi8 := 0; fi8 < 19; fi8++ {
            p0 += p0
            v7 = (v7 * v7)
            _ = fi8
        }
        fmt.Println(p0)
        _ = fi6
        _ = v7
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 uint8 = 225
        for fi11 := 0; fi11 < 11; fi11++ {
            v10 = uint8(p1)
            v12 := []int{(fi9 >> ((2) & 63)), -715, fi11}
            v13 := GenS1{F0: 127, F1: 1}
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = -1
    var v3 uint16 = 1
    var v4 int64 = v2
    var v5 uint64 = 753
    var v14 bool = true
    v3 += (v3 ^ 255)
    for fi15 := 0; fi15 < 16; fi15++ {
        for fi16 := 0; fi16 < 9; fi16++ {
            v17 := GenS2{F0: v2, F1: v5, F2: 622}
            v18 := []int{fi15}
            _ = fi16
            _ = v17
            _ = v18
        }
        var v19 uint16 = (v3 + 2)
        _ = fi15
        _ = v19
    }
    fmt.Println(v14, v5)
    v20 := []int{-1}
    if ((v3 * (v3 ^ v3)) < v3) {
        var v21 uint16 = (v3 * v3)
        _ = v21
    }
    fmt.Println(v4, v14, v3)
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    var v22 uint8 = 128
    switch 977 {
    case 4294967295:
        var v23 int32 = -1
        _ = v23
    case 127:
        fmt.Println(v14, v5)
    default:
        switch 136 {
        case 128:
            _ = len(v20)
        }
    }
    v3 += (v3 & 255)
    fmt.Println(v3, v14, v14)
    switch "hello" {
    case "foo bar":
        fmt.Println(v22)
    case "":
        v24 := GenS3{F0: 4294967295, F1: int32(v22), F2: int64(v5)}
        _ = v24
    default:
        v5 = 1
    }
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    v22 = v22
    v5 = (v5 / ((v5) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
