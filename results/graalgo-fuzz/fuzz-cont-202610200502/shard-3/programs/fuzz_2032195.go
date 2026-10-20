package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint64; F2 bool }
func genFn0(p0 uint, p1 int, p2 uint) int {
    switch (^(p1)) {
    case 213:
        v4 := []int{}
        _ = v4
    case 0:
        v5 := []int{(p1 / (((p1 << (((p2 * 0)) & 63))) | 1)), (p1 - p1), int(p2)}
        _ = v5
    case 32767:
        for fi6 := 0; fi6 < 7; fi6++ {
            v7 := GenS1{F0: 4, F1: ("hello" + "foo bar"), F2: 32767}
            for fi8 := 0; fi8 < 19; fi8++ {
                var v9 uint64 = 264
                p0 = uint(v9)
                var v10 int32 = 2147483647
                _ = fi8
                _ = v10
                _ = v9
            }
            v11 := GenS1{F0: 42, F1: "", F2: 42}
            _ = fi6
            _ = v11
            _ = v7
        }
    default:
        p0 -= p2
    }
    p0 = (p0 ^ 1)
    p1 = (p1 | 992)
    switch p1 {
    case -888:
        switch (p1 ^ (^(p1))) {
        case 0:
            v12 := []int{p1, p1}
            _ = v12
        default:
            v13 := []int{p1, (p1 >> ((p2) & 63)), (p1 | 437)}
            _ = v13
        }
    default:
        var v14 uint32 = 2
        _ = v14
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    return (p1 ^ (p1 | int(p2)))
}

func main() {
    var v1 uint8 = 114
    v2 := []uint32{}
    var v3 uint64 = 42
    var v15 bool = true
    var v16 uint = 255
    if v15 {
        for fi17 := 0; fi17 < 19; fi17++ {
            switch int(v16) {
            case -597:
                fmt.Println(fi17, v15)
            case 32768:
                v16 = (^(v16))
            }
            _ = fi17
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println(v15, v1, v15)
    var v18 int16 = 42
    v18 += v18
    v18 -= v18
    fmt.Println(v15, v3, v18)
    _ = len(v2)
    if false {
        var v19 int64 = -9223372036854775808
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v1, v16)
        _ = v19
    }
    var v20 int64 = 1
    fmt.Println(v3)
    var v21 int64 = v20
    var v22 float64 = 0.1
    v23 := GenS2{F0: 1, F1: v3, F2: v15}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
