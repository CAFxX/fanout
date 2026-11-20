package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 uint16 }
func genFn0(p0 uint16, p1 int16, p2 string) uint {
    p1 -= (^(p1))
    var v5 uint32 = 0
    if (2147483647 >= 1) {
        var v6 string = ("a" + p2)
        p1 = -32768
        _ = v6
    } else {
        var v7 int = 788
        for fi8 := 0; fi8 < 15; fi8++ {
            if true {
                var v9 uint8 = 42
                v10 := GenS1{F0: v5}
                v11 := []int{-133, (v7 + (v7 % (((v7 - v7)) | 1))), v7, 42}
                _ = v10
                _ = v11
                _ = v9
            } else {
                p2 = ("" + (p2 + p2))
                v12 := []int{32768, (v7 | (v7 + (fi8 % (((fi8 & 2147483647)) | 1))))}
                _ = v12
            }
            _ = fi8
        }
        _ = v7
    }
    v13 := GenS2{F0: (-(p1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v5
    return 656
}
func genFn1(p0 int32) uint8 {
    var v14 uint32 = 42
    v15 := GenS3{F0: true, F1: 281}
    _ = p0
    _ = v14
    _ = v15
    return 255
}

func main() {
    v1 := []uint32{42}
    v2 := make(map[uint32]bool)
    var v3 uint16 = 32767
    var v4 float64 = 3.14159
    v4 -= (3.14159 / 1e10)
    var v16 uint32 = 1
    v3 -= v3
    var v17 float32 = float32(3.14159)
    var v18 uint = 0
    var v19 string = "foo bar"
    for fi20 := 0; fi20 < 7; fi20++ {
        var v21 uint8 = 232
        var v22 int64 = 845
        _ = fi20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 3; fi23++ {
        var v24 float64 = v4
        _ = fi23
        _ = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
