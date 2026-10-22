package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 uint64; F2 int64 }
type GenS4 struct { F0 bool; F1 uint; F2 string }
type GenS5 struct { F0 uint8; F1 int32 }
func genFn0(p0 uint8, p1 uint, p2 int) uint16 {
    switch 9223372036854775807 {
    case 32768:
        if (false && false) {
            v4 := GenS1{F0: -2147483648, F1: 65535, F2: ((p2 * -395) < (p2 & (p2 >> ((uint(p0)) & 63))))}
            for fi5 := 0; fi5 < 4; fi5++ {
                p2 += 2147483647
                var v6 int32 = 1
                _ = fi5
                _ = v6
            }
            var v7 uint16 = 42
            _ = v4
            _ = v7
        }
    case 7:
        for fi8 := 0; fi8 < 9; fi8++ {
            if false {
                for fi9 := 0; fi9 < 14; fi9++ {
                    if (true && true) {
                        fmt.Println(p1)
                        v10 := []int{p2, (fi8 - p2)}
                        _ = v10
                    }
                    p1 = p1
                    _ = fi9
                }
                v11 := []int{fi8, (-(p2)), p2, 42}
                _ = v11
            } else {
                var v12 int8 = 42
                switch "" {
                case "":
                    switch p2 {
                    case -9223372036854775808:
                        p2 -= (fi8 + 127)
                    case -2147483648:
                        p0 -= uint8(p1)
                    case -924:
                        var v13 int8 = (v12 - 100)
                        _ = v13
                    }
                }
                _ = v12
            }
            p0 -= p0
            _ = fi8
        }
    }
    p2 -= (p2 * 9223372036854775807)
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 string, p1 uint8) int8 {
    p0 = p0
    p0 = p0
    var v14 int8 = -128
    v15 := GenS2{F0: 0}
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return v14
}

func main() {
    v1 := []uint32{795, 4294967295, 0}
    var v2 float32 = (1e10 * 1.5)
    v3 := []uint32{332, 322, 4294967295, 0}
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 uint8 = 198
        var v18 float64 = 1.5
        _ = fi16
        _ = v17
        _ = v18
    }
    switch 182 {
    case 663:
        if false {
            v19 := GenS1{F0: -343, F1: 42, F2: true}
            _ = v19
        }
    default:
        v2 -= 1.5
    }
    v20 := GenS1{F0: 0, F1: genFn0(164, 0, -528), F2: false}
    fmt.Println(v2)
    var v21 uint16 = 2
    var v22 uint8 = 127
    if true {
        switch ("foo bar" + "hello") {
        case "foo bar":
            var v23 float64 = ((0.0 - -2.25) / 1e10)
            _ = v23
        case "\n\t":
            if ((!true) && false) {
                v22 = v22
                for fi24 := 0; fi24 < 19; fi24++ {
                    if len(v1) > 0 {
                        v1[0] = 65535
                    }
                    _ = fi24
                }
            }
        }
    }
    v25 := GenS2{F0: v22}
    fmt.Println(v21)
    v26 := GenS2{F0: 2}
    v27 := GenS3{F0: -262, F1: 18446744073709551615, F2: 42}
    var v28 bool = (!false)
    var v29 int64 = 9223372036854775807
    for fi30 := 0; fi30 < 5; fi30++ {
        if (!(!((32767 >= 1) || v28))) {
            switch 47 {
            case 1:
                v31 := GenS4{F0: ((v29 << ((42) & 63)) < -9223372036854775808), F1: 886, F2: "\n\t"}
                _ = v31
            default:
                var v32 int16 = int16(v22)
                _ = v32
            }
        }
        switch fi30 {
        case 25:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case -532:
            v21 = 215
        case -442:
            fmt.Println(v28, v29)
        default:
            v33 := GenS2{F0: (v22 + (v22 ^ (v22 | (v22 << ((18446744073709551615) & 63)))))}
            _ = v33
        }
        _ = fi30
    }
    v34 := GenS5{F0: v22, F1: -48}
    fmt.Println(v2, v21)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v34
}
