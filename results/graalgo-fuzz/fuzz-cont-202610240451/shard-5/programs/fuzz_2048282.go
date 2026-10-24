package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int16; F1 int64; F2 int }
func genFn0(p0 int, p1 uint32, p2 uint32) uint32 {
    var v4 float32 = 1.5
    var v5 uint64 = 42
    v6 := []int{p0}
    var v7 int64 = 0
    if true {
        var v8 int = int(p2)
        p2 -= p1
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return (p2 - 745)
}
func genFn1(p0 int64) int64 {
    fmt.Println(p0)
    var v9 int = 2147483647
    _ = p0
    _ = v9
    return (p0 & -9223372036854775808)
}

func main() {
    var v1 uint32 = 720
    var v2 float64 = (0.0 / float64(-2.25))
    v3 := make(map[int]int)
    v10 := GenS1{F0: -614, F1: (v1 + v1)}
    v2 -= v2
    var v11 uint64 = 4294967295
    switch "a" {
    case "foo bar":
        var v12 int8 = -65
        _ = v12
    case "a":
        var v13 uint16 = 2
        _ = v13
    case "\n\t":
        v14 := []int{76, 1}
        _ = v14
    default:
        v15 := GenS2{F0: 2147483647, F1: 9223372036854775807}
        _ = v15
    }
    fmt.Println(v1)
    v11 = v11
    for fi16 := 0; fi16 < 14; fi16++ {
        if (676 > 42) {
            v17 := []int{786, (fi16 - (fi16 + (fi16 & 980))), 32767, fi16}
            _ = v17
        }
        if (true || false) {
            v11 = 1
        }
        _ = fi16
    }
    switch 857 {
    case 754:
        var v18 uint64 = v11
        _ = v18
    }
    if true {
        if true {
            var v19 uint64 = v11
            v3[127] = 79
            v11 = (v19 % (((^(v11))) | 1))
            _ = v19
        }
    }
    fmt.Println(v1)
    switch 32767 {
    case -614:
        v11 = v11
    default:
        var v20 uint8 = 1
        _ = v20
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        for fi22 := 0; fi22 < 2; fi22++ {
            switch "a" {
            case "a":
                v1 = (v1 & 2147483648)
            default:
                fmt.Println(v11)
            }
            v11 = 1
            if (123 != 128) {
                v11 += (^(v11))
                v23 := GenS3{F0: -1, F1: genFn1(-1), F2: -1}
                v1 -= v1
                _ = v23
            } else {
                if (false || ((123 != 94) && true)) {
                    v1 = v1
                } else {
                    if false {
                        var v24 float32 = float32(float64((-2.25 - 1.5)))
                        v24 += 0.0
                        v25 := []int{}
                        _ = v24
                        _ = v25
                    }
                    v3[-879] = 42
                }
            }
            _ = fi22
        }
        _ = fi21
    }
    fmt.Println(v11, v2, v1)
    v26 := GenS2{F0: 420, F1: 2147483647}
    var v27 int8 = -72
    v28 := GenS2{F0: -2147483648, F1: -26}
    for fi29 := 0; fi29 < 17; fi29++ {
        v3[(fi29 & (fi29 + 2147483647))] = -1
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
