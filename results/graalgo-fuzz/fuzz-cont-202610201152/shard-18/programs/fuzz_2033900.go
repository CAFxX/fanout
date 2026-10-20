package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 uint8; F1 uint8 }
func genFn0(p0 int8) uint16 {
    var v5 int64 = -9223372036854775808
    v6 := []int{}
    v5 -= v5
    for fi7 := 0; fi7 < 6; fi7++ {
        switch ("hello" + ("foo bar" + "hello")) {
        case "":
            if true {
                v6 = append(v6, (fi7 >> ((4294967295) & 63)))
                var v8 uint32 = 4294967295
                _ = v8
            }
        case "foo bar":
            p0 = (p0 << ((0) & 63))
        default:
            for fi9 := 0; fi9 < 6; fi9++ {
                switch "hello" {
                case "a":
                    fmt.Println(fi9, v5, v5)
                case "hello":
                    if len(v6) > 0 {
                        v6[0] = (-(fi9))
                    }
                case "\n\t":
                    if (128 == 127) {
                        v10 := GenS1{F0: 42, F1: 228, F2: "hello"}
                        p0 += (p0 % (((p0 & (-(p0)))) | 1))
                        switch fi9 {
                        case -2147483648:
                            var v11 string = "a"
                            _ = v11
                        case 9223372036854775807:
                            p0 = (p0 & (p0 / ((p0) | 1)))
                        default:
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                        p0 -= p0
                        v12 := GenS1{F0: 715, F1: 18446744073709551615, F2: ""}
                        _ = v10
                        _ = v12
                    }
                default:
                    var v13 int16 = 965
                    _ = v13
                }
                if ((-32768 > 199) || false) {
                    switch 2147483647 {
                    case 455:
                        _ = len(v6)
                    case -565:
                        v6 = append(v6, (^(fi9)))
                    case -2147483648:
                        if (65535 == 429) {
                            p0 = (p0 & p0)
                            fmt.Println(v5, p0)
                            if true {
                                var v14 uint32 = 472
                                switch 32768 {
                                case 2147483647:
                                    switch (fi7 ^ (fi7 / ((fi9) | 1))) {
                                    case 42:
                                        v5 -= int64(v14)
                                    case 867:
                                        if true {
                                            if true {
                                                if false {
                                                    if false {
                                                        v15 := GenS1{F0: 0, F1: 18446744073709551615, F2: ("\n\t" + "a")}
                                                        var v16 float64 = 0.0
                                                        _ = v15
                                                        _ = v16
                                                    }
                                                    v17 := GenS2{F0: 1}
                                                    v6 = append(v6, (fi9 * fi7))
                                                    _ = v17
                                                } else {
                                                    p0 = p0
                                                    _ = len(v6)
                                                    if (382 == 1) {
                                                        p0 = p0
                                                        v5 = v5
                                                    }
                                                    fmt.Println(p0)
                                                }
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                v6 = append(v6, fi7)
                                            } else {
                                                var v18 uint = 65535
                                                v5 += (-(v5))
                                                fmt.Println(fi9, v18, v14)
                                                fmt.Println(fi9, v18, v5)
                                                _ = v18
                                            }
                                            var v19 int = fi9
                                            _ = v19
                                        }
                                    }
                                    if (-2.25 < -2.25) {
                                        v14 = (^(v14))
                                        fmt.Println(v14)
                                    }
                                case 32767:
                                    v5 -= (v5 + v5)
                                default:
                                    v14 = 65535
                                }
                                _ = v14
                            }
                        }
                    default:
                        v5 = (-(v5))
                    }
                    var v20 float64 = float64(-2.25)
                    v5 = (v5 % (((v5 / ((42) | 1))) | 1))
                    _ = v20
                }
                switch 353 {
                case 39:
                    var v21 uint64 = 18446744073709551615
                    _ = v21
                case 32768:
                    v22 := GenS2{F0: 18446744073709551615}
                    _ = v22
                case 2147483647:
                    var v23 int16 = 1
                    _ = v23
                default:
                    var v24 uint8 = 86
                    _ = v24
                }
                _ = fi9
            }
        }
        _ = fi7
    }
    v6 = append(v6, -2147483648)
    _ = p0
    _ = v5
    _ = v6
    return 221
}

func main() {
    var v1 uint16 = 1
    var v2 int32 = 2147483647
    var v3 uint8 = 127
    v4 := []uint32{}
    v25 := GenS3{F0: 467, F1: 171, F2: v3}
    v4 = append(v4, 738)
    switch "foo bar" {
    case "foo bar":
        if true {
            var v26 uint64 = 4294967295
            var v27 bool = true
            _ = v26
            _ = v27
        }
    default:
        var v28 uint = 65535
        _ = v28
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v29 uint16 = genFn0(100)
    v29 = 255
    v30 := GenS4{F0: -478, F1: (v2 << ((2) & 63))}
    _ = len(v4)
    var v31 int = 65
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v32 uint8 = (v3 << ((0) & 63))
    v31 = v31
    if len(v4) > 0 {
        v4[0] = 699
    }
    if false {
        var v33 int8 = int8(v2)
        _ = v33
    } else {
        v29 = 624
        var v34 int32 = (v2 ^ v2)
        _ = v34
    }
    for fi35 := 0; fi35 < 8; fi35++ {
        var v36 int16 = int16(v29)
        switch "hello" {
        case "a":
            if len(v4) > 0 {
                v4[0] = 69
            }
        }
        _ = fi35
        _ = v36
    }
    if true {
        v37 := GenS5{F0: (v3 << ((758) & 63)), F1: v32}
        v29 = 255
        _ = v37
    } else {
        v32 = v3
    }
    var v38 uint64 = 155
    v39 := GenS4{F0: 128, F1: (v2 * 775)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v38
    _ = v39
    _ = v4
}
