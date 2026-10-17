package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int16 }
type GenS2 struct { F0 int64; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16; F1 bool }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int; F1 uint }
type GenS8 struct { F0 bool; F1 int; F2 int }
type GenS9 struct { F0 int32; F1 uint8; F2 uint }
func genFn0(p0 int, p1 uint) string {
    v5 := []int{p0, p0, (p0 >> ((1) & 63))}
    switch "\n\t" {
    case "a":
        var v6 string = "hello"
        _ = v6
    case "":
        fmt.Println(p0, p1)
    case "foo bar":
        if (1 > 42) {
            if len(v5) > 0 {
                v5[0] = (p0 - p0)
            }
            var v7 float32 = 3.14159
            for fi8 := 0; fi8 < 8; fi8++ {
                var v9 int = fi8
                _ = fi8
                _ = v9
            }
            _ = v7
        } else {
            if true {
                if (true && (true || true)) {
                    p0 = 32767
                    switch (p0 - p0) {
                    case -9223372036854775808:
                        for fi10 := 0; fi10 < 14; fi10++ {
                            _ = len(v5)
                            p0 = (^(fi10))
                            p0 += -9223372036854775808
                            _ = fi10
                        }
                    }
                }
                if false {
                    p0 += (p0 * 127)
                    v11 := GenS1{F0: 32767, F1: 858, F2: 1000}
                    _ = v11
                }
                if len(v5) > 0 {
                    v5[0] = (^(p0))
                }
            } else {
                var v12 int16 = -32768
                _ = len(v5)
                _ = v12
            }
        }
    default:
        v13 := GenS1{F0: 833, F1: 2, F2: 1000}
        _ = v13
    }
    if false {
        v5 = append(v5, p0)
    }
    switch p0 {
    case -385:
        if true {
            for fi14 := 0; fi14 < 15; fi14++ {
                var v15 float32 = 1.5
                for fi16 := 0; fi16 < 8; fi16++ {
                    v17 := GenS1{F0: 42, F1: 1, F2: -32768}
                    v18 := GenS2{F0: 1, F1: -717, F2: 109}
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                _ = fi14
                _ = v15
            }
            switch p0 {
            case 32767:
                var v19 uint16 = 2
                _ = v19
            default:
                for fi20 := 0; fi20 < 14; fi20++ {
                    for fi21 := 0; fi21 < 19; fi21++ {
                        if ((!false) && true) {
                            var v22 int = p0
                            _ = v22
                        }
                        switch fi20 {
                        case -159:
                            v23 := GenS2{F0: -9223372036854775808, F1: 1000, F2: 127}
                            _ = v23
                        case -86:
                            var v24 int16 = 0
                            _ = len(v5)
                            _ = v24
                        default:
                            v25 := GenS3{F0: 1, F1: 998}
                            _ = v25
                        }
                        _ = fi21
                    }
                    fmt.Println(p0, fi20)
                    if false {
                        switch "foo bar" {
                        case "\n\t":
                            if true {
                                fmt.Println(p0)
                            } else {
                                switch p0 {
                                case -363:
                                    for fi26 := 0; fi26 < 17; fi26++ {
                                        fmt.Println(fi20)
                                        if (1 <= 127) {
                                            var v27 int16 = -978
                                            fmt.Println(v27, fi20, fi20)
                                            v28 := GenS2{F0: 9223372036854775807, F1: v27, F2: 128}
                                            if false {
                                                v27 = (v27 & 0)
                                                if false {
                                                    v27 -= 1
                                                    var v29 int8 = 1
                                                    _ = v29
                                                } else {
                                                    v5 = append(v5, (fi20 >> ((p1) & 63)))
                                                    v30 := GenS4{F0: 1}
                                                    _ = v30
                                                }
                                            } else {
                                                if false {
                                                    var v31 uint32 = 2147483648
                                                    var v32 int = fi20
                                                    _ = v31
                                                    _ = v32
                                                }
                                                if len(v5) > 0 {
                                                    v5[0] = (fi26 & (fi26 * 42))
                                                }
                                            }
                                            v33 := GenS3{F0: 148, F1: 205}
                                            _ = v27
                                            _ = v28
                                            _ = v33
                                        } else {
                                            if ((2 >= 430) && false) {
                                                if false {
                                                    p1 -= p1
                                                    switch fi20 {
                                                    case -265:
                                                        v34 := GenS5{F0: 517, F1: ((float32(((((((((3.14159 * 0.1) - 1e10) - (float64(1e10) * float64(1e10))) / ((1.5 + (3.14159 * -2.25)) - 3.14159)) - 0.0) / (1e10 - (float64((float32((1.5 / (-2.25 + 3.14159))) * -2.25)) - (1.5 * (((3.14159 + ((((3.14159 + (1e10 * (0.1 * 0.1))) * float64(0.1)) * float64(0.0)) * 0.1)) - 0.0) * (float64(1e10) * (float64((-2.25 / 3.14159)) * (((3.14159 * 1e10) / -2.25) - ((1e10 / (0.0 - 0.1)) + (0.0 + 1.5)))))))))) - (3.14159 - -2.25)) + 1e10)) * (float32((1e10 + 0.1)) / (1e10 - 0.1))) < 0.0)}
                                                        _ = v34
                                                    default:
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    }
                                                    fmt.Println(p1, fi20, fi20)
                                                } else {
                                                    fmt.Println(fi20)
                                                    v5 = append(v5, 0)
                                                    var v35 int32 = 0
                                                    _ = v35
                                                }
                                                switch (fi26 - fi26) {
                                                case 32768:
                                                    v36 := GenS1{F0: 32767, F1: 710, F2: 1000}
                                                    _ = v36
                                                case 168:
                                                    var v37 uint8 = 42
                                                    _ = v37
                                                case 4294967295:
                                                    p0 = p0
                                                default:
                                                    p0 += 1
                                                }
                                                fmt.Println(p0)
                                            } else {
                                                _ = len(v5)
                                                var v38 int64 = 42
                                                _ = v38
                                            }
                                            v39 := GenS2{F0: -317, F1: 32767, F2: 128}
                                            var v40 float64 = ((0.0 * float64(-2.25)) * (-2.25 * 1.5))
                                            _ = v39
                                            _ = v40
                                        }
                                        var v41 uint16 = 944
                                        v42 := GenS5{F0: (v41 % (((^(v41))) | 1)), F1: true}
                                        v41 = (v41 % (((v41 >> ((p1) & 63))) | 1))
                                        _ = len(v5)
                                        var v43 float64 = (1.5 / 0.1)
                                        _ = fi26
                                        _ = v41
                                        _ = v42
                                        _ = v43
                                    }
                                case 32767:
                                    if (!false) {
                                        fmt.Println(fi20, p0)
                                    }
                                case -2147483648:
                                    p0 = fi20
                                }
                                _ = len(v5)
                            }
                        case "foo bar":
                            p0 += (p0 - (-(fi20)))
                        default:
                            for fi44 := 0; fi44 < 2; fi44++ {
                                if (((!true) && true) || (2 <= 18446744073709551615)) {
                                    if (!(true && ((false || true) || (((fi20 | fi44) < (p0 | (p0 - p0))) || false)))) {
                                        p1 += 84
                                        p1 = (p1 ^ p1)
                                        v45 := GenS6{F0: 0}
                                        _ = v45
                                    }
                                    var v46 uint16 = 473
                                    _ = v46
                                }
                                v47 := GenS3{F0: 4294967295, F1: 987}
                                if true {
                                    fmt.Println(p0, p0, fi20)
                                }
                                _ = fi44
                                _ = v47
                            }
                        }
                        if (2 <= 2147483648) {
                            p0 -= fi20
                        }
                    } else {
                        for fi48 := 0; fi48 < 18; fi48++ {
                            p0 += fi20
                            switch (^(fi20)) {
                            case 1:
                                switch p0 {
                                case 32768:
                                    p0 = 1
                                case -2147483648:
                                    var v49 uint16 = 255
                                    _ = v49
                                }
                            default:
                                v5 = append(v5, (fi20 % ((128) | 1)))
                            }
                            _ = fi48
                        }
                        switch (-(p0)) {
                        case -806:
                            var v50 bool = (1e10 == 1e10)
                            _ = v50
                        default:
                            var v51 uint16 = uint16(p1)
                            _ = v51
                        }
                    }
                    _ = fi20
                }
            }
            if (true && true) {
                fmt.Println(p0)
                switch "" {
                case "a":
                    var v52 uint32 = 42
                    _ = v52
                default:
                    var v53 int64 = 1
                    _ = v53
                }
                switch "\n\t" {
                case "\n\t":
                    v54 := GenS6{F0: 0}
                    _ = v54
                }
            } else {
                v55 := GenS2{F0: 741, F1: -1, F2: 128}
                _ = v55
            }
        }
    case 127:
        if (p0 >= (p0 | (p0 % ((p0) | 1)))) {
            p0 = p0
            var v56 int16 = 1
            if (!true) {
                var v57 uint32 = 2147483648
                for fi58 := 0; fi58 < 9; fi58++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi58
                }
                _ = v57
            }
            _ = v56
        }
    }
    p0 = 315
    _ = p0
    _ = p1
    _ = v5
    return ("\n\t" + "a")
}
func genFn1(p0 uint16, p1 int16) bool {
    fmt.Println(p0, p0)
    if false {
        p0 += p0
        if false {
            if false {
                fmt.Println(p0, p0)
            } else {
                v59 := []int{458, -883, -1}
                if len(v59) > 0 {
                    v59[0] = 4294967295
                }
                _ = v59
            }
            switch -2147483648 {
            case 0:
                v60 := []int{128}
                _ = v60
            default:
                for fi61 := 0; fi61 < 1; fi61++ {
                    switch 605 {
                    case -612:
                        p1 = p1
                    case -1:
                        for fi62 := 0; fi62 < 8; fi62++ {
                            v63 := []int{(fi61 * (fi61 >> ((65535) & 63))), (fi62 & 1), (^(fi62))}
                            _ = fi62
                            _ = v63
                        }
                    case 42:
                        var v64 float32 = 0.1
                        _ = v64
                    default:
                        if (!true) {
                            var v65 bool = (true || (true && (1 > 127)))
                            _ = v65
                        } else {
                            for fi66 := 0; fi66 < 13; fi66++ {
                                p1 = (p1 & p1)
                                _ = fi66
                            }
                            switch (("foo bar" + "hello") + "\n\t") {
                            case "":
                                switch 42 {
                                case -911:
                                    v67 := []int{-1, 2147483647, fi61, fi61}
                                    _ = v67
                                }
                            }
                        }
                    }
                    if true {
                        v68 := GenS7{F0: fi61, F1: 0}
                        var v69 uint8 = 2
                        _ = v68
                        _ = v69
                    }
                    _ = fi61
                }
            }
            switch 4294967295 {
            case 407:
                var v70 uint64 = 42
                _ = v70
            }
        } else {
            var v71 uint = 612
            _ = v71
        }
        v72 := []int{0, 32768, 4294967295, 32767}
        _ = v72
    } else {
        fmt.Println(p0, p1)
        fmt.Println(p1)
        p0 = 42
    }
    var v73 float32 = (0.0 + 1e10)
    p1 = p1
    _ = p0
    _ = p1
    _ = v73
    return (!true)
}

func main() {
    var v1 float32 = 3.14159
    v2 := []uint32{0, 65535, 0}
    v3 := make(map[string]int)
    var v4 uint = 4294967295
    v1 -= 0.0
    for fi74 := 0; fi74 < 9; fi74++ {
        for fi75 := 0; fi75 < 18; fi75++ {
            var v76 uint64 = 2
            _ = fi75
            _ = v76
        }
        var v77 int32 = 418
        var v78 uint64 = 4294967295
        _ = fi74
        _ = v77
        _ = v78
    }
    fmt.Println(v1, v1)
    fmt.Println(v4, v1)
    for fi79 := 0; fi79 < 17; fi79++ {
        var v80 int32 = 42
        _ = fi79
        _ = v80
    }
    var v81 int8 = 9
    v82 := GenS8{F0: false, F1: -1, F2: 1}
    if (1 != 0) {
        switch "" {
        case "":
            switch "a" {
            case "a":
                switch (("\n\t" + "hello") + ((("\n\t" + "") + "foo bar") + "a")) {
                case "foo bar":
                    var v83 uint64 = 1
                    _ = v83
                default:
                    v4 = (^(v4))
                }
            case "foo bar":
                fmt.Println(v4)
            case "hello":
                if ((-(v81)) == (v81 * v81)) {
                    for fi84 := 0; fi84 < 18; fi84++ {
                        var v85 uint = (^(v4))
                        var v86 int64 = 1
                        _ = fi84
                        _ = v85
                        _ = v86
                    }
                }
            }
        }
    }
    for fi87 := 0; fi87 < 4; fi87++ {
        v88 := GenS9{F0: 551, F1: 177, F2: v4}
        _ = fi87
        _ = v88
    }
    v3["\n\t"] = -1
    for fi89 := 0; fi89 < 1; fi89++ {
        var v90 int64 = -9223372036854775808
        for fi91 := 0; fi91 < 11; fi91++ {
            var v92 uint = 18446744073709551615
            v4 = v92
            _ = fi91
            _ = v92
        }
        _ = fi89
        _ = v90
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v81
    _ = v82
}
