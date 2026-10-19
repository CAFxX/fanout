package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS4 struct { F0 uint16; F1 uint8; F2 string }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]string)
    v3 := make(map[int]bool)
    var v4 uint8 = 42
    v5 := []int{70}
    if true {
        var v6 float64 = 0.0
        _ = v6
    }
    for fi7 := 0; fi7 < 19; fi7++ {
        v4 = v4
        _ = fi7
    }
    v4 = 187
    var v8 uint16 = 32767
    if len(v5) > 0 {
        v5[0] = 9223372036854775807
    }
    if len(v5) > 0 {
        v5[0] = -9223372036854775808
    }
    v5 = append(v5, 788)
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 uint32 = 42
        switch fi9 {
        case 861:
            v4 = (^(v4))
        case 862:
            var v11 int = (^(fi9))
            _ = v11
        case -9223372036854775808:
            var v12 uint64 = 219
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    var v13 uint8 = 0
    if len(v5) > 0 {
        v5[0] = 1
    }
    if len(v5) > 0 {
        v5[0] = 896
    }
    switch ("hello" + ("foo bar" + "a")) {
    case "foo bar":
        fmt.Println(v8, v8)
    default:
        switch ("a" + (("foo bar" + "a") + "hello")) {
        case "foo bar":
            v14 := GenS1{F0: v4, F1: -2147483648, F2: 18446744073709551615}
            _ = v14
        case "":
            var v15 uint64 = 42
            _ = v15
        case "\n\t":
            if false {
                for fi16 := 0; fi16 < 14; fi16++ {
                    if false {
                        if (18446744073709551615 > 4294967295) {
                            switch (^(fi16)) {
                            case 469:
                                var v17 bool = false
                                _ = v17
                            case 32767:
                                var v18 int8 = 0
                                _ = v18
                            case 9223372036854775807:
                                v19 := GenS1{F0: 255, F1: 74, F2: 42}
                                _ = v19
                            default:
                                if true {
                                    for fi20 := 0; fi20 < 3; fi20++ {
                                        var v21 bool = false
                                        _ = fi20
                                        _ = v21
                                    }
                                    var v22 bool = true
                                    _ = v22
                                }
                            }
                            v3[(fi16 - fi16)] = true
                            v23 := GenS1{F0: (v13 ^ 187), F1: -374, F2: 18446744073709551615}
                            _ = v23
                        }
                    }
                    v13 -= (v13 >> ((uint(v8)) & 63))
                    for fi24 := 0; fi24 < 17; fi24++ {
                        _ = len(v3)
                        v25 := GenS1{F0: v13, F1: -2147483648, F2: 1}
                        v13 = 128
                        _ = fi24
                        _ = v25
                    }
                    _ = fi16
                }
            }
        default:
            if (-222 < -9223372036854775808) {
                v13 += (v13 % ((128) | 1))
                if (!true) {
                    for fi26 := 0; fi26 < 10; fi26++ {
                        v1[4294967295] = true
                        _ = fi26
                    }
                    for fi27 := 0; fi27 < 16; fi27++ {
                        v13 -= v4
                        if false {
                            v13 = v13
                        } else {
                            v8 += v8
                        }
                        switch "a" {
                        case "\n\t":
                            var v28 uint64 = 2
                            _ = v28
                        case "a":
                            fmt.Println(v4)
                        default:
                            for fi29 := 0; fi29 < 2; fi29++ {
                                v30 := GenS2{F0: 791, F1: ("a" + ("hello" + (((("foo bar" + (("" + "a") + "hello")) + "a") + "hello") + (("" + ("a" + "a")) + "a"))))}
                                v8 -= (v8 | v8)
                                _ = fi29
                                _ = v30
                            }
                        }
                        _ = fi27
                    }
                } else {
                    if (!false) {
                        fmt.Println(v4, v13)
                    } else {
                        fmt.Println(v8)
                        for fi31 := 0; fi31 < 18; fi31++ {
                            var v32 int64 = 1
                            _ = fi31
                            _ = v32
                        }
                        v4 -= (v4 / (((v4 ^ (v13 | v13))) | 1))
                    }
                    for fi33 := 0; fi33 < 3; fi33++ {
                        switch 127 {
                        case 4294967295:
                            v34 := GenS2{F0: 42, F1: "foo bar"}
                            _ = v34
                        case 9223372036854775807:
                            if true {
                                switch fi33 {
                                case 42:
                                    v4 = (v13 - (v13 & v4))
                                case -686:
                                    for fi35 := 0; fi35 < 19; fi35++ {
                                        switch -554 {
                                        case 9223372036854775807:
                                            if ((false || (true || true)) && (255 < (v8 ^ (v8 - 0)))) {
                                                switch (("hello" + "") + (("a" + "foo bar") + "\n\t")) {
                                                case "hello":
                                                    var v36 int = 543
                                                    _ = v36
                                                case "\n\t":
                                                    v8 = (^(v8))
                                                default:
                                                    var v37 uint64 = 912
                                                    _ = v37
                                                }
                                                v8 = (v8 | (v8 - v8))
                                                fmt.Println(v4, v8, fi33)
                                                v38 := GenS2{F0: 0, F1: ""}
                                                _ = v38
                                            }
                                        case 2147483647:
                                            if (!false) {
                                                var v39 int8 = 84
                                                v40 := GenS3{F0: 0, F1: 65535, F2: 4294967295}
                                                _ = len(v5)
                                                _ = v39
                                                _ = v40
                                            }
                                        case 12:
                                            var v41 uint = 65535
                                            _ = v41
                                        default:
                                            v8 -= v8
                                        }
                                        _ = fi35
                                    }
                                default:
                                    if false {
                                        v4 -= 255
                                        fmt.Println(v13, v13)
                                        v8 -= (^(v8))
                                    } else {
                                        switch "foo bar" {
                                        case "hello":
                                            for fi42 := 0; fi42 < 3; fi42++ {
                                                var v43 int = (fi33 << ((18446744073709551615) & 63))
                                                v8 += 42
                                                _ = fi42
                                                _ = v43
                                            }
                                        default:
                                            var v44 int16 = 1
                                            _ = v44
                                        }
                                    }
                                }
                            }
                        case -858:
                            var v45 float64 = (1e10 - 0.1)
                            _ = v45
                        default:
                            v4 -= (v13 % (((v4 << ((255) & 63))) | 1))
                        }
                        _ = fi33
                    }
                    for fi46 := 0; fi46 < 14; fi46++ {
                        switch fi46 {
                        case 598:
                            var v47 uint16 = (v8 ^ (v8 >> ((66) & 63)))
                            _ = v47
                        case 42:
                            for fi48 := 0; fi48 < 4; fi48++ {
                                var v49 int = (fi46 << ((65535) & 63))
                                fmt.Println(fi46)
                                v8 = v8
                                _ = fi48
                                _ = v49
                            }
                        case 4294967295:
                            var v50 float32 = ((1e10 + (3.14159 / 1e10)) + ((1e10 - 1.5) - (float32(3.14159) / ((1e10 * 1e10) + 1.5))))
                            _ = v50
                        default:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                        v4 = 1
                        if (!true) {
                            v4 -= v4
                            fmt.Println(v13, v13)
                        }
                        _ = fi46
                    }
                }
                v4 += 244
            } else {
                for fi51 := 0; fi51 < 4; fi51++ {
                    switch fi51 {
                    case 128:
                        var v52 uint16 = v8
                        _ = v52
                    }
                    v53 := GenS4{F0: v8, F1: v4, F2: "a"}
                    _ = fi51
                    _ = v53
                }
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
