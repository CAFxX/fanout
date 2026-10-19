package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int; F1 string }
func genFn0(p0 int64, p1 int32) uint32 {
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS1{F0: 0, F1: 42}
        if (((1.5 <= 1e10) && (0.0 < 0.1)) && ((!(true || true)) && (false || (!(!false))))) {
            if (2 > 255) {
                p0 = (-(p0))
            }
            fmt.Println(p0, p0)
        } else {
            p0 += -1
        }
        switch ("" + "") {
        case "foo bar":
            p1 += (p1 - -105)
        case "":
            for fi9 := 0; fi9 < 12; fi9++ {
                v10 := []int{fi9, -888, fi7, (-(fi7))}
                _ = fi9
                _ = v10
            }
        case "a":
            if true {
                v11 := GenS1{F0: 934, F1: -261}
                p0 -= (p0 + (p0 % ((p0) | 1)))
                p1 += -910
                _ = v11
            } else {
                var v12 float32 = -2.25
                _ = v12
            }
        default:
            var v13 int32 = p1
            _ = v13
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(p0, p0)
    var v14 uint64 = 841
    switch -9223372036854775808 {
    case -1:
        var v15 uint = 1
        _ = v15
    case 32768:
        for fi16 := 0; fi16 < 8; fi16++ {
            v17 := []int{fi16, 32767}
            var v18 uint32 = 567
            _ = fi16
            _ = v17
            _ = v18
        }
    case 42:
        var v19 uint32 = 2147483647
        _ = v19
    default:
        for fi20 := 0; fi20 < 17; fi20++ {
            p0 = p0
            var v21 float64 = 3.14159
            _ = fi20
            _ = v21
        }
    }
    p1 = (-(p1))
    _ = p0
    _ = p1
    _ = v14
    return uint32(v14)
}
func genFn1(p0 uint32) string {
    v22 := []int{128, 195, -9223372036854775808, -1}
    p0 -= genFn0(-916, 580)
    _ = p0
    _ = v22
    return "foo bar"
}

func main() {
    v1 := []int{4294967295, 1, -1}
    var v2 float32 = (3.14159 / -2.25)
    var v3 int16 = 501
    var v4 uint = 255
    v5 := make(map[string]uint32)
    v6 := make(map[uint32]bool)
    v23 := GenS1{F0: 2, F1: v3}
    if true {
        _ = len(v1)
        var v24 uint = (v4 + (v4 >> (((v4 ^ v4)) & 63)))
        v4 += 42
        _ = v24
    } else {
        v4 = (v4 % (((v4 / (((v4 | (^(v4)))) | 1))) | 1))
    }
    var v25 bool = true
    var v26 uint64 = 853
    v26 -= v26
    fmt.Println(v3, v3)
    v2 = (float32((float64((float32(1.5) + v2)) * 1.5)) + 0.0)
    for fi27 := 0; fi27 < 19; fi27++ {
        v26 += v26
        if false {
            v25 = v25
            if true {
                fmt.Println(v2, v2)
                v4 = (v4 ^ (v4 / ((v4) | 1)))
            }
            v28 := GenS2{F0: 9223372036854775807, F1: "a"}
            _ = v28
        }
        _ = fi27
    }
    v29 := GenS2{F0: -284, F1: ("foo bar" + "foo bar")}
    v1 = append(v1, 2147483647)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
