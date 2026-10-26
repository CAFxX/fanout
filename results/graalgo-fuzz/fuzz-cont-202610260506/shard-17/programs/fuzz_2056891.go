package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 int64; F2 bool }
type GenS3 struct { F0 uint8; F1 uint32; F2 int16 }
func genFn0(p0 uint16, p1 uint8, p2 int) uint8 {
    var v5 float32 = (-2.25 / 1.5)
    switch p2 {
    case 127:
        for fi6 := 0; fi6 < 8; fi6++ {
            v7 := []int{p2, int(p0)}
            _ = fi6
            _ = v7
        }
    }
    var v8 uint = 65535
    v9 := []int{(p2 ^ 2147483647), p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v8
    _ = v9
    return p1
}
func genFn1(p0 int8) int {
    if false {
        p0 = 1
    }
    var v10 uint8 = genFn0(30, 127, 1)
    _ = p0
    _ = v10
    return -1
}

func main() {
    var v1 int16 = 32767
    v2 := []string{"a", ("hello" + "a"), ("a" + "\n\t"), "a"}
    var v3 int64 = 185
    var v4 float32 = (1.5 * 1e10)
    for fi11 := 0; fi11 < 19; fi11++ {
        v1 = v1
        v1 += (v1 * v1)
        _ = fi11
    }
    switch -548 {
    case -817:
        for fi12 := 0; fi12 < 13; fi12++ {
            v13 := GenS1{F0: 0}
            _ = fi12
            _ = v13
        }
    case 780:
        v14 := GenS2{F0: false, F1: v3, F2: false}
        _ = v14
    case 9223372036854775807:
        for fi15 := 0; fi15 < 14; fi15++ {
            fmt.Println(fi15, v1)
            _ = fi15
        }
    default:
        var v16 uint16 = 42
        _ = v16
    }
    if len(v2) > 0 {
        v2[0] = ("foo bar" + "hello")
    }
    var v17 int8 = 0
    if true {
        fmt.Println(v3, v3)
        v4 -= float32((float64(v4) * 1e10))
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 uint32 = 668
        v19 = (v19 / (((v19 / (((^(v19))) | 1))) | 1))
        _ = fi18
        _ = v19
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v4)
    v20 := GenS3{F0: 127, F1: 2147483648, F2: (-(v1))}
    var v21 uint = 989
    for fi22 := 0; fi22 < 19; fi22++ {
        var v23 uint32 = 0
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
