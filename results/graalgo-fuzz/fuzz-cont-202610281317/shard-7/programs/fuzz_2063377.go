package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 int; F2 uint16 }
type GenS4 struct { F0 uint64; F1 int }
func genFn0() int8 {
    v5 := GenS1{F0: -1}
    switch ("foo bar" + "") {
    case "a":
        var v6 int64 = 0
        _ = v6
    case "hello":
        v7 := []int{}
        _ = v7
    case "\n\t":
        v8 := []int{}
        _ = v8
    }
    _ = v5
    return 100
}
func genFn1(p0 int16) int16 {
    fmt.Println(p0)
    fmt.Println(p0)
    p0 = p0
    var v9 uint = 65535
    switch -2147483648 {
    case 0:
        p0 -= 0
    case 32767:
        fmt.Println(v9, v9)
    }
    _ = p0
    _ = v9
    return -443
}

func main() {
    v1 := []string{}
    v2 := []uint32{2147483647, 118, 1, 2147483648}
    v3 := make(map[int]bool)
    var v4 int64 = 1
    for fi10 := 0; fi10 < 14; fi10++ {
        v11 := GenS1{F0: 42}
        v4 -= v4
        _ = fi10
        _ = v11
    }
    var v12 int64 = v4
    switch "foo bar" {
    case "hello":
        v13 := GenS2{F0: -1}
        _ = v13
    case "\n\t":
        var v14 float64 = ((float64((-2.25 - (((1.5 * float32(0.0)) - 0.0) / -2.25))) - (-2.25 + -2.25)) + (float64(-2.25) + -2.25))
        _ = v14
    default:
        var v15 uint16 = 0
        _ = v15
    }
    v16 := GenS2{F0: -2147483648}
    v2 = append(v2, 2147483648)
    for fi17 := 0; fi17 < 8; fi17++ {
        _ = len(v3)
        for fi18 := 0; fi18 < 17; fi18++ {
            var v19 int = fi17
            if ((!(255 != 1)) && (42 >= 1)) {
                v12 = (^(v12))
                var v20 int64 = 292
                _ = v20
            }
            _ = fi18
            _ = v19
        }
        switch "\n\t" {
        case "foo bar":
            var v21 float32 = 0.1
            _ = v21
        default:
            v4 = -16
        }
        _ = fi17
    }
    if false {
        fmt.Println(v12, v12)
        v22 := v3[42]
        _ = v22
    } else {
        v12 -= v12
        if false {
            v12 -= v12
            var v23 int = 1
            v12 = (v12 >> ((1) & 63))
            _ = v23
        } else {
            if true {
                v24 := GenS3{F0: 1, F1: -917, F2: 2}
                _ = v24
            } else {
                v12 = v4
                var v25 int32 = 2147483647
                _ = v25
            }
        }
    }
    for fi26 := 0; fi26 < 18; fi26++ {
        var v27 int16 = -124
        if false {
            fmt.Println(fi26, fi26, v27)
            v28 := GenS1{F0: (v27 & 1)}
            v29 := GenS4{F0: 2, F1: (fi26 / (((fi26 << ((18446744073709551615) & 63))) | 1))}
            _ = v28
            _ = v29
        }
        _ = fi26
        _ = v27
    }
    switch -2147483648 {
    case -1:
        v30 := v3[708]
        _ = v30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
