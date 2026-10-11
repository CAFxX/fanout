package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 uint32 }
func genFn0(p0 bool, p1 uint8) uint8 {
    v6 := GenS1{F0: -128}
    if false {
        var v7 int = 42
        _ = v7
    } else {
        v8 := []int{2147483647, 0}
        if len(v8) > 0 {
            v8[0] = 32767
        }
        for fi9 := 0; fi9 < 11; fi9++ {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            fmt.Println(p1, p0)
            _ = fi9
        }
        _ = v8
    }
    if (p0 || true) {
        var v10 int32 = 2147483647
        switch 9223372036854775807 {
        case -9223372036854775808:
            var v11 int64 = 0
            _ = v11
        case -313:
            v12 := []int{0, -920}
            _ = v12
        default:
            p1 += p1
        }
        v13 := GenS1{F0: 1}
        _ = v10
        _ = v13
    } else {
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = v6
    return p1
}
func genFn1(p0 int16, p1 string, p2 uint64) int8 {
    p2 -= (p2 & p2)
    var v14 uint64 = (^(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 127
}

func main() {
    v1 := []string{"\n\t"}
    var v2 int32 = -161
    v3 := make(map[string]int)
    var v4 bool = ((-128 != 42) || false)
    v5 := make(map[string]bool)
    var v15 uint = 65535
    v16 := GenS1{F0: genFn1(-418, "\n\t", 18446744073709551615)}
    if v4 {
        if v4 {
            v15 = (v15 - (v15 | v15))
            for fi17 := 0; fi17 < 4; fi17++ {
                v18 := GenS1{F0: 99}
                _ = fi17
                _ = v18
            }
        } else {
            v2 -= -126
        }
    }
    v1 = append(v1, ("\n\t" + "hello"))
    v19 := GenS1{F0: -128}
    fmt.Println(v4, v15, v15)
    var v20 uint8 = 35
    _ = len(v5)
    v20 -= genFn0((!v4), v20)
    var v21 int64 = -921
    var v22 uint16 = 32767
    switch -31 {
    case 32767:
        v23 := GenS2{F0: v22, F1: 4294967295}
        _ = v23
    default:
        _ = len(v1)
    }
    v20 = v20
    v24 := GenS1{F0: int8(v20)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
