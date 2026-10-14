package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int16; F2 int32 }
type GenS3 struct { F0 int32; F1 uint; F2 uint }
type GenS4 struct { F0 int; F1 uint64; F2 uint8 }
func genFn0(p0 uint64, p1 int8) int16 {
    var v6 uint64 = p0
    if (false && true) {
        fmt.Println(p1, v6, v6)
        switch "foo bar" {
        case "":
            for fi7 := 0; fi7 < 13; fi7++ {
                var v8 uint8 = 2
                _ = fi7
                _ = v8
            }
        case "hello":
            if (!(1 != 65)) {
                p0 += (p0 | (p0 ^ 578))
            } else {
                fmt.Println(v6)
            }
        case "\n\t":
            v9 := GenS1{F0: 9223372036854775807}
            _ = v9
        default:
            switch ("" + ("" + "foo bar")) {
            case "hello":
                var v10 int16 = 428
                _ = v10
            case "":
                v11 := []int{844, 565, -69, 128}
                _ = v11
            default:
                v12 := []int{1, 0, 32767}
                _ = v12
            }
        }
        v13 := GenS2{F0: 65535, F1: 1, F2: 1}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    v1 := []uint32{}
    var v2 uint8 = 119
    v3 := []uint32{}
    v4 := []int{-2147483648}
    var v5 int = -9223372036854775808
    for fi14 := 0; fi14 < 7; fi14++ {
        _ = len(v1)
        _ = fi14
    }
    var v15 int8 = 127
    var v16 float64 = (1e10 + 0.0)
    if (42 != 1000) {
        fmt.Println(v2)
        switch "foo bar" {
        case "foo bar":
            v16 -= (1.5 / 1e10)
        default:
            var v17 int16 = 1000
            _ = v17
        }
        switch "foo bar" {
        case "\n\t":
            v18 := GenS1{F0: 894}
            _ = v18
        case "hello":
            v4 = append(v4, (v5 / ((v5) | 1)))
        default:
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
        }
    }
    v19 := GenS3{F0: -2147483648, F1: 0, F2: 1}
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    v20 := GenS4{F0: (v5 | v5), F1: 1, F2: (v2 ^ v2)}
    for fi21 := 0; fi21 < 18; fi21++ {
        if false {
            var v22 uint32 = 412
            for fi23 := 0; fi23 < 15; fi23++ {
                if false {
                    fmt.Println(v15)
                }
                fmt.Println(fi23)
                var v24 uint16 = 1
                _ = fi23
                _ = v24
            }
            _ = v22
        } else {
            if false {
                var v25 uint32 = 0
                var v26 bool = (int64(v5) >= -9223372036854775808)
                fmt.Println(v16)
                _ = v25
                _ = v26
            } else {
                v5 = (fi21 | fi21)
            }
            if false {
                v27 := GenS4{F0: (fi21 - v5), F1: 2, F2: v2}
                for fi28 := 0; fi28 < 19; fi28++ {
                    _ = len(v3)
                    v16 += (((v16 / ((v16 / (v16 + -2.25)) * v16)) - v16) + v16)
                    _ = fi28
                }
                v15 += (v15 / ((127) | 1))
                _ = v27
            }
            for fi29 := 0; fi29 < 4; fi29++ {
                var v30 bool = (967 != 0)
                var v31 float64 = 1e10
                var v32 bool = (((18446744073709551615 < 221) && ((v2 & (v2 % ((v2) | 1))) < 255)) || true)
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
            }
        }
        v15 += (v15 % ((int8(v2)) | 1))
        _ = fi21
    }
    v3 = append(v3, 2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
