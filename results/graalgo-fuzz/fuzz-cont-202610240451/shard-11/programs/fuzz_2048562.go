package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint8 }
type GenS2 struct { F0 bool; F1 uint32; F2 uint16 }
func genFn0(p0 uint, p1 int64) uint {
    if (true && (p1 <= p1)) {
        var v4 float64 = 0.0
        v5 := GenS1{F0: p1, F1: -547, F2: 1}
        if true {
            var v6 uint64 = 456
            fmt.Println(v6, p1)
            v7 := []int{}
            _ = v6
            _ = v7
        } else {
            v8 := GenS1{F0: (p1 & int64(p0)), F1: -790, F2: 50}
            if true {
                p0 = (p0 & (p0 - p0))
                for fi9 := 0; fi9 < 6; fi9++ {
                    var v10 float32 = 0.1
                    _ = fi9
                    _ = v10
                }
            }
            var v11 uint32 = 1
            _ = v11
            _ = v8
        }
        _ = v4
        _ = v5
    }
    switch -936 {
    case 32768:
        v12 := []int{}
        _ = v12
    default:
        v13 := []int{-1, 127, 857}
        _ = v13
    }
    v14 := []int{}
    _ = p0
    _ = p1
    _ = v14
    return (p0 ^ p0)
}
func genFn1(p0 uint8, p1 int8, p2 uint8) uint8 {
    var v15 int32 = 42
    p1 -= (p1 + p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return p0
}

func main() {
    v1 := []uint32{}
    v2 := make(map[int]uint32)
    var v3 int = 4294967295
    if (("\n\t" + "\n\t") == (("foo bar" + "foo bar") + "hello")) {
        var v16 int64 = 42
        v3 -= 32767
        _ = v16
    }
    v17 := v2[1]
    if true {
        if false {
            v3 = v3
            fmt.Println(v17)
            var v18 int64 = 9223372036854775807
            _ = v18
        } else {
            switch "foo bar" {
            case "hello":
                fmt.Println(v17, v17)
            default:
                v3 = (v3 % ((199) | 1))
            }
        }
        if (!true) {
            delete(v2, (v3 / ((v3) | 1)))
        } else {
            v19 := GenS1{F0: 9223372036854775807, F1: (v3 & v3), F2: 2}
            for fi20 := 0; fi20 < 19; fi20++ {
                v3 -= v3
                v21 := GenS2{F0: (!false), F1: (v17 << ((18446744073709551615) & 63)), F2: 1}
                _ = fi20
                _ = v21
            }
            _ = v19
        }
    }
    var v22 float64 = 1.5
    _ = len(v2)
    var v23 bool = false
    for fi24 := 0; fi24 < 13; fi24++ {
        switch v3 {
        case 32768:
            fmt.Println(v17)
        case 2147483647:
            fmt.Println(fi24)
        case -456:
            _ = len(v2)
        default:
            v22 -= 1e10
        }
        if v23 {
            v25 := GenS1{F0: 42, F1: 0, F2: 255}
            v3 = (v3 - 2147483647)
            var v26 uint = 255
            _ = v25
            _ = v26
        }
        _ = fi24
    }
    _ = len(v1)
    _ = len(v2)
    fmt.Println(v22, v23)
    var v27 uint32 = 0
    v23 = (0.0 < (-2.25 * v22))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v3
}
