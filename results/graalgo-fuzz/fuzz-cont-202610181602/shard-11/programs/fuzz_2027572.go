package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 uint32 }
func genFn0(p0 int64, p1 int, p2 int8) uint8 {
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := GenS1{F0: p1, F1: 18446744073709551615, F2: (p2 - (p2 << ((806) & 63)))}
        v8 := []int{fi6, 360, int(p2), (p1 ^ p1)}
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := []int{(p1 / ((32768) | 1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 1
}
func genFn1(p0 int16) int16 {
    switch 1 {
    case 1:
        p0 += p0
    case 127:
        var v10 int16 = (p0 & (p0 % ((p0) | 1)))
        _ = v10
    case -9223372036854775808:
        fmt.Println(p0)
    default:
        var v11 bool = true
        _ = v11
    }
    v12 := []int{493, 128}
    for fi13 := 0; fi13 < 7; fi13++ {
        v14 := GenS2{F0: 197, F1: "\n\t"}
        if len(v12) > 0 {
            v12[0] = fi13
        }
        p0 = (p0 | p0)
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = v12
    return p0
}

func main() {
    v1 := []int{0, -2147483648, 591}
    v2 := make(map[uint32]string)
    v3 := make(map[int]uint32)
    var v4 uint8 = 1
    v5 := make(map[uint32]bool)
    var v15 int16 = -980
    var v16 float32 = -2.25
    var v17 uint16 = 65535
    var v18 string = "foo bar"
    fmt.Println(v16, v4)
    v19 := GenS3{F0: -35}
    v20 := GenS4{F0: 963, F1: 2147483647}
    for fi21 := 0; fi21 < 4; fi21++ {
        _ = len(v5)
        var v22 uint64 = 311
        _ = fi21
        _ = v22
    }
    if true {
        v23 := GenS3{F0: 0}
        v15 = (-(v15))
        _ = v23
    }
    for fi24 := 0; fi24 < 2; fi24++ {
        var v25 uint64 = 4294967295
        if true {
            if (13 == -1) {
                var v26 int32 = int32(fi24)
                _ = v26
            }
            v4 = 128
            for fi27 := 0; fi27 < 18; fi27++ {
                v25 -= (v25 * (v25 & 1))
                fmt.Println(v4, v25)
                v15 = v15
                _ = fi27
            }
        } else {
            fmt.Println(v18)
        }
        _ = fi24
        _ = v25
    }
    v28 := GenS3{F0: -20}
    v1 = append(v1, -769)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
