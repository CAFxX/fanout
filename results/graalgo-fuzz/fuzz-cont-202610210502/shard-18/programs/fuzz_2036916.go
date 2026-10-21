package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 uint8; F1 int8 }
func genFn0(p0 int64) uint32 {
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 uint8 = 122
        var v7 int32 = 730
        var v8 int64 = (p0 | p0)
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    p0 = (p0 | (-(p0)))
    _ = p0
    return 0
}
func genFn1(p0 bool) uint8 {
    v9 := []int{-85}
    p0 = p0
    v9 = append(v9, 42)
    if true {
        fmt.Println(p0)
        switch -9223372036854775808 {
        case -2147483648:
            if true {
                var v10 uint16 = 670
                _ = v10
            } else {
                p0 = (!p0)
            }
        default:
            var v11 uint = 0
            _ = v11
        }
    } else {
        var v12 float64 = 1.5
        var v13 int8 = -5
        _ = v12
        _ = v13
    }
    _ = p0
    _ = v9
    return 1
}

func main() {
    v1 := []string{"foo bar", "hello", "\n\t", ("hello" + "foo bar")}
    var v2 uint = 1
    var v3 uint16 = 752
    v4 := make(map[string]int)
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 uint = 42
        for fi16 := 0; fi16 < 7; fi16++ {
            var v17 string = "a"
            _ = fi16
            _ = v17
        }
        _ = fi14
        _ = v15
    }
    v2 -= v2
    v3 -= (v3 & v3)
    v18 := GenS1{F0: 32767, F1: -556}
    switch "foo bar" {
    case "a":
        v19 := GenS2{F0: 0, F1: (v3 / ((v3) | 1))}
        _ = v19
    case "\n\t":
        switch 9223372036854775807 {
        case -54:
            var v20 uint = (v2 * (v2 - 18446744073709551615))
            _ = v20
        default:
            if (0 <= 0) {
                if (uint32(v3) == 0) {
                    switch "a" {
                    case "foo bar":
                        var v21 int8 = 100
                        _ = v21
                    }
                }
            }
        }
    case "hello":
        v3 = v3
    default:
        v22 := GenS2{F0: 2147483647, F1: v3}
        _ = v22
    }
    switch "\n\t" {
    case "":
        var v23 int8 = 127
        _ = v23
    case "a":
        v3 -= (v3 << (((v2 / ((v2) | 1))) & 63))
    default:
        if ((42 != 32767) || (100 == 108)) {
            v24 := GenS3{F0: 128, F1: 100}
            _ = v24
        }
    }
    if false {
        v3 += (v3 * (v3 >> ((1) & 63)))
        fmt.Println(v3)
        v3 -= (v3 >> (((v2 - 1)) & 63))
    }
    fmt.Println(v3, v2)
    fmt.Println(v2, v2)
    for fi25 := 0; fi25 < 5; fi25++ {
        v2 = v2
        var v26 int16 = 1
        v26 = -467
        _ = fi25
        _ = v26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
