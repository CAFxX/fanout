package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 string; F1 uint16; F2 int8 }
func genFn0(p0 string, p1 uint, p2 string) int16 {
    v7 := []int{-9223372036854775808}
    for fi8 := 0; fi8 < 11; fi8++ {
        p2 = (p0 + (("hello" + p0) + ("" + p2)))
        v7 = append(v7, (^(fi8)))
        var v9 int32 = 1
        _ = fi8
        _ = v9
    }
    if len(v7) > 0 {
        v7[0] = 128
    }
    switch 128 {
    case 9223372036854775807:
        if true {
            var v10 uint16 = 1
            if (((false || true) || (true || (!(127 > 76)))) || ((1.5 != -2.25) || (true && false))) {
                _ = len(v7)
                var v11 int32 = 2147483647
                var v12 uint16 = v10
                _ = v11
                _ = v12
            }
            _ = v10
        }
    case 128:
        fmt.Println(p2, p0, p0)
    case -545:
        v13 := GenS1{F0: -1, F1: -275, F2: 0}
        _ = v13
    }
    var v14 float64 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v7
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = ""
    var v3 uint64 = 545
    var v4 string = (v2 + "hello")
    var v5 uint8 = 118
    v6 := []uint32{975, 681}
    v3 -= 2
    v2 = "foo bar"
    var v15 int32 = 0
    v2 = (v4 + (v2 + v2))
    if false {
        var v16 int = 32767
        _ = v16
    } else {
        v5 = (v5 >> ((18446744073709551615) & 63))
        v17 := GenS2{F0: (v5 * 158)}
        for fi18 := 0; fi18 < 18; fi18++ {
            v19 := GenS1{F0: -28, F1: (v15 & -1), F2: v5}
            v20 := GenS3{F0: (v2 + "a"), F1: -1}
            fmt.Println(v5, v5, v5)
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = v17
    }
    switch "foo bar" {
    case "":
        v15 -= v15
    case "a":
        _ = len(v6)
    case "foo bar":
        v21 := GenS4{F0: v2, F1: uint16(v3), F2: -1}
        _ = v21
    }
    var v22 int64 = -1
    for fi23 := 0; fi23 < 4; fi23++ {
        switch v2 {
        case "hello":
            for fi24 := 0; fi24 < 14; fi24++ {
                v25 := GenS2{F0: v5}
                _ = fi24
                _ = v25
            }
        case "a":
            switch 936 {
            case -2147483648:
                var v26 float32 = -2.25
                _ = v26
            case -407:
                v5 = (v5 + (v5 - v5))
            case 42:
                var v27 uint8 = v5
                _ = v27
            }
        case "foo bar":
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        default:
            fmt.Println(v22, v15)
        }
        _ = fi23
    }
    v22 -= (-(v22))
    v4 = "\n\t"
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v28 := v1[-897]
    v3 += uint64(v28)
    v29 := GenS1{F0: 127, F1: v15, F2: (v5 | (^(v5)))}
    fmt.Println(v2)
    fmt.Println(v3)
    v4 = (("foo bar" + v2) + v4)
    v2 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
