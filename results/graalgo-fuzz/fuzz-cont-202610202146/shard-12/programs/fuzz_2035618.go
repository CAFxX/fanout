package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint16 }
type GenS2 struct { F0 int; F1 int8; F2 int8 }
type GenS3 struct { F0 int32; F1 uint8; F2 int16 }
func genFn0() uint {
    switch 1 {
    case 32767:
        var v5 uint8 = 82
        _ = v5
    case 128:
        var v6 uint64 = 4294967295
        _ = v6
    case -2147483648:
        v7 := []int{1, 936}
        _ = v7
    default:
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 string = "hello"
            v9 = v9
            var v10 int8 = 42
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    var v11 bool = false
    fmt.Println(v11)
    v12 := []int{-11, -1}
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    _ = v11
    _ = v12
    return 304
}
func genFn1(p0 uint16, p1 int) uint32 {
    p0 = p0
    v13 := GenS1{F0: 1000, F1: 127, F2: (^(p0))}
    _ = p0
    _ = p1
    _ = v13
    return 2147483647
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 4294967295
    var v3 uint = 14
    var v4 int16 = 100
    v14 := GenS1{F0: v4, F1: 0, F2: 177}
    var v15 uint16 = 255
    var v16 int8 = -15
    var v17 float64 = (((1e10 / (0.0 - 3.14159)) / 0.1) + float64(float32(((0.1 - 3.14159) - 3.14159))))
    for fi18 := 0; fi18 < 1; fi18++ {
        fmt.Println(v15)
        fmt.Println(v2, v15)
        v4 -= v4
        _ = fi18
    }
    var v19 uint8 = 128
    v4 = (v4 * 1000)
    v19 -= 100
    var v20 int8 = v16
    v17 += v17
    fmt.Println(v2)
    var v21 int = 1
    v22 := GenS2{F0: (v21 & (v21 ^ -2147483648)), F1: v20, F2: v20}
    for fi23 := 0; fi23 < 20; fi23++ {
        if false {
            v24 := GenS3{F0: 2147483647, F1: v19, F2: v4}
            var v25 int64 = 897
            _ = v24
            _ = v25
        }
        v15 += (^(v15))
        _ = fi23
    }
    v26 := []int{715}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v4
}
