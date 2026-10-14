package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 uint16; F2 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint; F1 int16 }
func genFn0(p0 uint16, p1 int, p2 int64) uint {
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := GenS1{F0: p2}
        fmt.Println(p0, p0)
        _ = fi7
        _ = v8
    }
    v9 := GenS2{F0: ((!(563 == 0)) && true)}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 4294967295
}
func genFn1(p0 bool, p1 int64, p2 bool) int32 {
    v10 := GenS2{F0: false}
    v11 := []int{-108, -1, 9223372036854775807, -375}
    var v12 bool = (-1 != -1)
    switch int(p1) {
    case -9223372036854775808:
        v13 := GenS3{F0: -1, F1: 217, F2: 670}
        _ = v13
    case 128:
        switch 0 {
        case 9223372036854775807:
            var v14 uint16 = 668
            _ = v14
        default:
            if (("a" + ("" + "")) < "") {
                p1 = p1
            }
        }
    default:
        var v15 float32 = ((3.14159 + (0.0 * (0.1 - 3.14159))) + (0.0 * float32(1.5)))
        _ = v15
    }
    switch ("hello" + "foo bar") {
    case "\n\t":
        switch 1 {
        case 42:
            for fi16 := 0; fi16 < 19; fi16++ {
                fmt.Println(fi16, p2, p0)
                for fi17 := 0; fi17 < 4; fi17++ {
                    var v18 float32 = (float32((-2.25 + -2.25)) * 1.5)
                    var v19 int = (fi16 | 128)
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = fi16
            }
        }
    case "foo bar":
        v11 = append(v11, -9223372036854775808)
    default:
        v11 = append(v11, 127)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return -2147483648
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]bool)
    var v3 int = 128
    var v4 int = -9223372036854775808
    var v5 uint32 = 0
    var v6 string = "\n\t"
    v20 := []int{(^(v4))}
    v4 += v4
    v4 = v3
    var v21 int8 = -1
    v22 := GenS4{F0: v3}
    if len(v20) > 0 {
        v20[0] = v4
    }
    v4 -= (v4 / ((v4) | 1))
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        for fi24 := 0; fi24 < 5; fi24++ {
            v6 = "a"
            v25 := GenS5{F0: genFn0(228, v3, 42), F1: 489}
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    var v26 uint8 = 2
    v5 = v5
    v27 := GenS2{F0: (65535 != (v5 << ((10) & 63)))}
    var v28 uint64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
