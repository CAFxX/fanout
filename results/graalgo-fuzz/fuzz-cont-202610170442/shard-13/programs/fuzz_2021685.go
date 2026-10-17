package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint; F1 int16; F2 bool }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 int32; F1 uint; F2 uint64 }
type GenS7 struct { F0 uint; F1 uint; F2 uint16 }
func genFn0(p0 int64, p1 int32) int8 {
    v5 := GenS1{F0: 4294967295, F1: (^(p0))}
    switch "foo bar" {
    case "foo bar":
        p0 = p0
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v5
    return 90
}
func genFn1(p0 uint64, p1 uint32, p2 uint) uint8 {
    v6 := GenS2{F0: p2, F1: 32767, F2: false}
    if ((^(p1)) > 4294967295) {
        switch 381 {
        case 0:
            fmt.Println(p2, p0, p2)
        default:
            if ((519 != 229) && (false || ((float32(3.14159) * 3.14159) <= 3.14159))) {
                v7 := GenS3{F0: -1, F1: 0}
                _ = v7
            }
        }
        switch "foo bar" {
        case "\n\t":
            v8 := GenS4{F0: p1, F1: p1}
            _ = v8
        case "hello":
            fmt.Println(p2, p0, p0)
        case "a":
            switch -9223372036854775808 {
            case 127:
                v9 := []int{}
                _ = v9
            default:
                p2 = p2
            }
        default:
            if (1.5 < (1e10 * 3.14159)) {
                for fi10 := 0; fi10 < 10; fi10++ {
                    p2 = (^(p2))
                    for fi11 := 0; fi11 < 16; fi11++ {
                        switch fi10 {
                        case 419:
                            switch (fi10 >> (((p2 / (((p2 | (p2 | uint(p0)))) | 1))) & 63)) {
                            case 4294967295:
                                fmt.Println(fi10, p2)
                            case 127:
                                p1 += p1
                            case 0:
                                v12 := GenS5{F0: true, F1: 0}
                                _ = v12
                            default:
                                if true {
                                    p2 -= p2
                                } else {
                                    v13 := GenS5{F0: (false && true), F1: -9223372036854775808}
                                    _ = v13
                                }
                            }
                        case -9223372036854775808:
                            var v14 int32 = 54
                            _ = v14
                        case 776:
                            v15 := GenS2{F0: p2, F1: -32768, F2: (fi11 < fi10)}
                            _ = v15
                        }
                        p0 += (p0 % ((2) | 1))
                        p1 = (p1 % ((2) | 1))
                        fmt.Println(p1, p0)
                        var v16 int32 = -933
                        v17 := []int{}
                        var v18 float64 = 0.0
                        _ = len(v17)
                        _ = fi11
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    fmt.Println(fi10, p1, p2)
                    _ = fi10
                }
            }
        }
    } else {
        var v19 int16 = -32768
        for fi20 := 0; fi20 < 7; fi20++ {
            var v21 int8 = 127
            _ = fi20
            _ = v21
        }
        var v22 string = ("" + ("" + "foo bar"))
        _ = v19
        _ = v22
    }
    v23 := []int{}
    for fi24 := 0; fi24 < 13; fi24++ {
        v25 := GenS5{F0: (-87 != -128), F1: 288}
        _ = fi24
        _ = v25
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    _ = v6
    return uint8(p2)
}

func main() {
    var v1 int8 = -101
    var v2 float64 = (0.1 * ((1.5 + -2.25) * (3.14159 - 0.0)))
    var v3 bool = (-1 > -2147483648)
    var v4 uint = 2
    if (128 > 128) {
        v1 = (v1 >> (((v4 & (v4 >> ((v4) & 63)))) & 63))
        fmt.Println(v3, v3, v4)
        var v26 uint8 = 1
        _ = v26
    }
    v27 := []int{-9223372036854775808, 128}
    var v28 uint64 = 146
    v2 += v2
    var v29 bool = (32767 >= -2147483648)
    v30 := GenS6{F0: -311, F1: (v4 / (((v4 / (((v4 << ((v4) & 63))) | 1))) | 1)), F2: (v28 >> ((uint(v28)) & 63))}
    switch 760 {
    case 534:
        var v31 uint = (v4 + v4)
        _ = v31
    case -587:
        v29 = v3
    }
    fmt.Println(v1, v2)
    var v32 float64 = float64(0.0)
    var v33 uint16 = 1
    var v34 uint = (v4 >> ((uint(v28)) & 63))
    var v35 bool = (!(v29 && false))
    fmt.Println(v29, v33, v33)
    for fi36 := 0; fi36 < 4; fi36++ {
        switch fi36 {
        case 891:
            var v37 int32 = 2147483647
            _ = v37
        case 0:
            var v38 uint32 = 2147483648
            _ = v38
        case 4294967295:
            var v39 int64 = -9223372036854775808
            _ = v39
        }
        _ = fi36
    }
    var v40 uint64 = v28
    for fi41 := 0; fi41 < 6; fi41++ {
        var v42 bool = (v35 && (v29 && v3))
        v43 := GenS7{F0: (v4 << ((42) & 63)), F1: v4, F2: v33}
        _ = fi41
        _ = v42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v40
}
