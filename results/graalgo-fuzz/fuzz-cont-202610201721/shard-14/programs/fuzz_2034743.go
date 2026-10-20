package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool; F1 int8; F2 uint16 }
type GenS6 struct { F0 int64; F1 int16; F2 uint32 }
type GenS7 struct { F0 int16; F1 uint; F2 int8 }
type GenS8 struct { F0 uint8; F1 int16 }
type GenS9 struct { F0 int16; F1 uint32; F2 int64 }
type GenS10 struct { F0 int; F1 uint8; F2 string }
type GenS11 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS12 struct { F0 int32; F1 uint32 }
type GenS13 struct { F0 int; F1 int }
type GenS14 struct { F0 string }
type GenS15 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS16 struct { F0 bool; F1 int }
type GenS17 struct { F0 int64 }
type GenS18 struct { F0 uint8; F1 int32; F2 string }
type GenS19 struct { F0 uint8 }
type GenS20 struct { F0 uint; F1 bool; F2 int32 }
type GenS21 struct { F0 int32 }
type GenS22 struct { F0 string }
type GenS23 struct { F0 int64 }
type GenS24 struct { F0 string; F1 int8 }
type GenS25 struct { F0 uint16 }
type GenS26 struct { F0 int16; F1 uint64 }
type GenS27 struct { F0 int32; F1 int16 }
type GenS28 struct { F0 int64; F1 int32 }
type GenS29 struct { F0 int8; F1 uint; F2 int8 }
type GenS30 struct { F0 int8 }
type GenS31 struct { F0 bool }
type GenS32 struct { F0 int32; F1 uint8; F2 uint16 }
type GenS33 struct { F0 uint64; F1 uint64; F2 uint }
type GenS34 struct { F0 int8 }
type GenS35 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS36 struct { F0 uint16 }
type GenS37 struct { F0 bool; F1 int16; F2 uint16 }
type GenS38 struct { F0 int32 }
func genFn0(p0 int32) int8 {
    if ((0 == -567) || false) {
        v5 := GenS1{F0: 100, F1: "\n\t"}
        for fi6 := 0; fi6 < 14; fi6++ {
            fmt.Println(p0)
            _ = fi6
        }
        _ = v5
    } else {
        v7 := []int{}
        _ = v7
    }
    fmt.Println(p0)
    _ = p0
    return -126
}
func genFn1(p0 int8, p1 string, p2 bool) int32 {
    if p2 {
        switch "foo bar" {
        case "\n\t":
            if (169 >= 65535) {
                var v8 int = 1
                p0 = genFn0(-2147483648)
                _ = v8
            }
        case "":
            var v9 bool = (177 >= 339)
            _ = v9
        case "a":
            if ((p1 <= (("\n\t" + p1) + "\n\t")) && p2) {
                fmt.Println(p2, p0, p1)
            } else {
                var v10 int = -1
                var v11 int64 = 1
                switch (v10 - 1) {
                case 9223372036854775807:
                    switch "hello" {
                    case "":
                        v12 := []int{v10, (v10 / (((v10 * (v10 + (-(v10))))) | 1)), (^(v10))}
                        _ = v12
                    }
                default:
                    if (0 >= -1) {
                        v10 += (^(v10))
                        v13 := []int{}
                        _ = v13
                    }
                }
                _ = v10
                _ = v11
            }
        }
        var v14 int = 32767
        _ = v14
    } else {
        var v15 uint16 = 65535
        p2 = p2
        _ = v15
    }
    p2 = (42 > 137)
    v16 := []int{32767, 32767}
    if (0 <= 9223372036854775807) {
        var v17 uint64 = 0
        switch 229 {
        case 42:
            var v18 int = 0
            _ = v18
        case 2147483647:
            for fi19 := 0; fi19 < 9; fi19++ {
                switch 32768 {
                case 157:
                    if ((false && p2) || p2) {
                        var v20 float32 = 0.1
                        for fi21 := 0; fi21 < 20; fi21++ {
                            v20 = v20
                            v20 -= v20
                            _ = fi21
                        }
                        _ = len(v16)
                        _ = v20
                    }
                default:
                    v22 := GenS2{F0: p0, F1: 127}
                    _ = v22
                }
                p0 = (p0 << ((65535) & 63))
                var v23 float64 = 0.0
                _ = fi19
                _ = v23
            }
        }
        _ = v17
    } else {
        switch 0 {
        case 0:
            p0 = (p0 / (((p0 & p0)) | 1))
        case -853:
            v16 = append(v16, 42)
        default:
            fmt.Println(p2)
        }
        if false {
            if false {
                fmt.Println(p0, p1)
                switch 0 {
                case -173:
                    var v24 string = p1
                    _ = v24
                }
            }
            switch p1 {
            case "foo bar":
                v25 := GenS3{F0: p0, F1: p2}
                _ = v25
            case "hello":
                switch 128 {
                case 42:
                    var v26 uint8 = 177
                    _ = v26
                case 2147483647:
                    fmt.Println(p1)
                case -440:
                    for fi27 := 0; fi27 < 17; fi27++ {
                        p1 = p1
                        _ = fi27
                    }
                }
            }
        } else {
            p2 = p2
        }
    }
    p0 -= genFn0(42)
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return -2147483648
}

func main() {
    var v1 int = 1
    var v2 int16 = -32768
    var v3 int64 = -1
    var v4 float32 = 0.0
    v28 := []int{}
    for fi29 := 0; fi29 < 13; fi29++ {
        v3 = (v3 * (v3 & (v3 >> ((609) & 63))))
        var v30 uint32 = 0
        var v31 float64 = 1.5
        _ = fi29
        _ = v30
        _ = v31
    }
    _ = len(v28)
    fmt.Println(v1)
    if ((v3 == int64(v2)) || ((1 == genFn1(127, "foo bar", (65535 != 32767))) || false)) {
        if len(v28) >= 2 {
            v28 = v28[1:len(v28)]
        }
        for fi32 := 0; fi32 < 15; fi32++ {
            for fi33 := 0; fi33 < 11; fi33++ {
                v1 = (fi33 * int(v2))
                var v34 int8 = -78
                var v35 uint8 = 127
                _ = fi33
                _ = v34
                _ = v35
            }
            _ = fi32
        }
    }
    var v36 int8 = 42
    if len(v28) > 0 {
        v28[0] = v1
    }
    v37 := GenS3{F0: v36, F1: (!true)}
    var v38 uint32 = 4294967295
    fmt.Println(v2)
    var v39 int8 = (v36 + -116)
    v39 = (v36 * (v39 * (v36 - (v39 * genFn0(42)))))
    if ((v1 * (^(v1))) > (v1 >> ((0) & 63))) {
        v1 = v1
    } else {
        var v40 uint64 = 93
        v2 -= (v2 & v2)
        _ = v40
    }
    switch v1 {
    case -9223372036854775808:
        if false {
            var v41 float64 = ((0.1 - (((0.0 * 0.1) + -2.25) + (0.0 - (1.5 * 0.0)))) * 1e10)
            _ = v41
        }
    case -296:
        v3 = v3
    }
    for fi42 := 0; fi42 < 6; fi42++ {
        if false {
            if (!(fi42 != fi42)) {
                if ((!false) && false) {
                    fmt.Println(v39)
                    v43 := GenS1{F0: v39, F1: ""}
                    _ = v43
                }
                v1 = 128
                var v44 float32 = (float32(0.0) / (v4 * (-2.25 * -2.25)))
                _ = v44
            }
        } else {
            fmt.Println(v4)
            for fi45 := 0; fi45 < 14; fi45++ {
                if (float32(0.1) < 3.14159) {
                    v3 = v3
                    var v46 uint32 = (^(v38))
                    var v47 float32 = 0.1
                    _ = v46
                    _ = v47
                } else {
                    fmt.Println(v39, v36, v3)
                    switch (("" + "") + ("foo bar" + (("hello" + (((("a" + "foo bar") + "\n\t") + "a") + "foo bar")) + "foo bar"))) {
                    case "foo bar":
                        switch "a" {
                        case "":
                            var v48 uint32 = v38
                            _ = v48
                        default:
                            v49 := GenS2{F0: (v39 % ((v36) | 1)), F1: 226}
                            _ = v49
                        }
                    case "":
                        var v50 uint8 = 218
                        _ = v50
                    case "a":
                        var v51 int8 = v36
                        v39 = v51
                        if true {
                            switch ("foo bar" + ("hello" + ("a" + ("" + ("\n\t" + "\n\t"))))) {
                            case "\n\t":
                                v51 = v39
                            case "hello":
                                switch ("hello" + "\n\t") {
                                case "a":
                                    switch "" {
                                    case "a":
                                        var v52 uint = 42
                                        _ = v52
                                    case "":
                                        if (128 == 255) {
                                            var v53 string = "a"
                                            _ = v53
                                        }
                                    default:
                                        fmt.Println(v36, v38)
                                    }
                                case "foo bar":
                                    _ = len(v28)
                                }
                            case "a":
                                if true {
                                    switch (fi42 / (((v1 * fi42)) | 1)) {
                                    case 32767:
                                        v4 = 0.1
                                    default:
                                        var v54 uint = 255
                                        _ = v54
                                    }
                                    fmt.Println(v3, fi45, v36)
                                    v55 := GenS3{F0: genFn0(0), F1: false}
                                    _ = v55
                                } else {
                                    if (!false) {
                                        var v56 float64 = 3.14159
                                        _ = v56
                                    }
                                    if true {
                                        if true {
                                            var v57 uint16 = 2
                                            _ = v57
                                        } else {
                                            if (!((-(v3)) < v3)) {
                                                if true {
                                                    var v58 int = fi42
                                                    var v59 int16 = v2
                                                    switch (fi45 >> ((1) & 63)) {
                                                    case 42:
                                                        var v60 uint16 = 32767
                                                        _ = v60
                                                    case -81:
                                                        if (0 != 138) {
                                                            v38 = v38
                                                        } else {
                                                            switch ("" + ("" + "")) {
                                                            case "a":
                                                                if (!false) {
                                                                    var v61 int = (v1 & (fi45 + -306))
                                                                    v58 -= 563
                                                                    _ = v61
                                                                } else {
                                                                    v62 := GenS4{F0: 2}
                                                                    _ = v62
                                                                }
                                                                v28 = append(v28, int(v38))
                                                            case "\n\t":
                                                                var v63 int16 = (v59 << ((349) & 63))
                                                                _ = v63
                                                            case "":
                                                                v64 := GenS2{F0: 42, F1: 128}
                                                                _ = v64
                                                            default:
                                                                var v65 int8 = v36
                                                                v66 := GenS3{F0: int8(v3), F1: (false || false)}
                                                                fmt.Println(v4, v38, fi45)
                                                                _ = v65
                                                                _ = v66
                                                            }
                                                        }
                                                    case 32767:
                                                        v59 -= (v2 & v59)
                                                    }
                                                    _ = v58
                                                    _ = v59
                                                }
                                                var v67 uint8 = 52
                                                switch ("" + "a") {
                                                case "a":
                                                    if false {
                                                        if true {
                                                            if ((true && (false && true)) || (((1.5 - float64(((((((v4 / v4) / float32(0.1)) + (1e10 / ((-2.25 * float32(3.14159)) - (3.14159 - (float32(1.5) + (0.1 / 1e10)))))) - v4) * v4) + (((0.1 * float32(1.5)) - v4) - float32(3.14159))))) <= (0.1 + 1e10)) && ((!false) && (1 == 636)))) {
                                                                v68 := GenS4{F0: v67}
                                                                var v69 float32 = (3.14159 + (((float32(-2.25) + v4) / (v4 * 0.1)) - 0.1))
                                                                switch "" {
                                                                case "foo bar":
                                                                    if ((!true) || (v67 <= (v67 | v67))) {
                                                                        v1 += fi42
                                                                        v3 -= int64(v67)
                                                                        switch (v1 << ((161) & 63)) {
                                                                        case 32767:
                                                                            fmt.Println(v3, v1)
                                                                        case 1:
                                                                            switch fi45 {
                                                                            case -853:
                                                                                v4 += -2.25
                                                                            }
                                                                            var v70 bool = (42 == 950)
                                                                            _ = v70
                                                                        }
                                                                    } else {
                                                                        v2 = -398
                                                                        if (!false) {
                                                                            v38 = 4294967295
                                                                            if len(v28) > 0 {
                                                                                v28[0] = fi42
                                                                            }
                                                                        } else {
                                                                            v39 = genFn0(genFn1(v39, "hello", (v4 == float32(1.5))))
                                                                            v71 := GenS5{F0: (v3 == v3), F1: v51, F2: 766}
                                                                            var v72 uint8 = v67
                                                                            if true {
                                                                                if (true && (((!((float64((float32(-2.25) / (v4 + v69))) - 3.14159) >= (1.5 / 3.14159))) && (v3 < (-(v3)))) || (true || (uint32(v67) <= 1000)))) {
                                                                                    v73 := GenS2{F0: 42, F1: 127}
                                                                                    v28 = append(v28, v1)
                                                                                    _ = v73
                                                                                }
                                                                                fmt.Println(v3, v51, v4)
                                                                                if false {
                                                                                    if len(v28) > 0 {
                                                                                        v28[0] = 32768
                                                                                    }
                                                                                    if (32767 != 818) {
                                                                                        if ((!false) || false) {
                                                                                            if (1e10 >= 3.14159) {
                                                                                                v74 := GenS3{F0: 100, F1: (v4 > ((v4 + 0.0) + float32(3.14159)))}
                                                                                                _ = v74
                                                                                            } else {
                                                                                                v2 = (v2 ^ v2)
                                                                                            }
                                                                                            v75 := GenS6{F0: 9223372036854775807, F1: v2, F2: (v38 << ((42) & 63))}
                                                                                            var v76 string = ("foo bar" + "foo bar")
                                                                                            _ = v75
                                                                                            _ = v76
                                                                                        }
                                                                                        v28 = append(v28, v1)
                                                                                        if (true || (!(!(!(((!false) && true) && true))))) {
                                                                                            if (v2 <= v2) {
                                                                                                fmt.Println(v36)
                                                                                            } else {
                                                                                                if len(v28) > 0 {
                                                                                                    v28[0] = (-(fi45))
                                                                                                }
                                                                                                switch -778 {
                                                                                                case -540:
                                                                                                    fmt.Println(fi45, v3, v39)
                                                                                                case 723:
                                                                                                    v28 = append(v28, fi45)
                                                                                                }
                                                                                            }
                                                                                            var v77 bool = (!(false && (-1 < 42)))
                                                                                            if v77 {
                                                                                                switch (-(fi42)) {
                                                                                                case 9223372036854775807:
                                                                                                    fmt.Println(v4, v39)
                                                                                                case 818:
                                                                                                    fmt.Println(v67)
                                                                                                }
                                                                                                v36 += -43
                                                                                                fmt.Println(v72, v77, v3)
                                                                                                var v78 float64 = 0.0
                                                                                                var v79 uint16 = 65535
                                                                                                _ = v78
                                                                                                _ = v79
                                                                                            } else {
                                                                                                switch (v1 / (((fi45 + v1)) | 1)) {
                                                                                                case -384:
                                                                                                    if (!(true && false)) {
                                                                                                        v80 := GenS6{F0: -557, F1: (v2 | 42), F2: 4294967295}
                                                                                                        _ = v80
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            if len(v28) >= 2 {
                                                                                                v28 = v28[1:len(v28)]
                                                                                            }
                                                                                            _ = v77
                                                                                        } else {
                                                                                            var v81 int64 = -9223372036854775808
                                                                                            if (!(v38 < (v38 >> ((42) & 63)))) {
                                                                                                switch fi45 {
                                                                                                case 0:
                                                                                                    var v82 float64 = -2.25
                                                                                                    _ = v82
                                                                                                case 32767:
                                                                                                    v1 += -979
                                                                                                case -1:
                                                                                                    switch "\n\t" {
                                                                                                    case "foo bar":
                                                                                                        v2 -= int16(v38)
                                                                                                    default:
                                                                                                        switch "" {
                                                                                                        case "foo bar":
                                                                                                            v38 = v38
                                                                                                        }
                                                                                                    }
                                                                                                default:
                                                                                                    var v83 uint32 = (v38 ^ 4294967295)
                                                                                                    _ = v83
                                                                                                }
                                                                                                v84 := GenS7{F0: (v2 & v2), F1: 65535, F2: v51}
                                                                                                if len(v28) >= 2 {
                                                                                                    v28 = v28[1:len(v28)]
                                                                                                }
                                                                                                _ = v84
                                                                                            } else {
                                                                                                if true {
                                                                                                    v38 = v38
                                                                                                    var v85 int32 = 0
                                                                                                    v86 := GenS2{F0: -1, F1: (v72 % ((247) | 1))}
                                                                                                    _ = v85
                                                                                                    _ = v86
                                                                                                } else {
                                                                                                    if len(v28) >= 2 {
                                                                                                        v28 = v28[1:len(v28)]
                                                                                                    }
                                                                                                    v81 = -9223372036854775808
                                                                                                    switch (fi45 * 1) {
                                                                                                    case 127:
                                                                                                        v36 = v39
                                                                                                    case 1:
                                                                                                        switch fi42 {
                                                                                                        case 32768:
                                                                                                            if (0 == 313) {
                                                                                                                fmt.Println(v51)
                                                                                                                fmt.Println(v1, v2)
                                                                                                                v39 -= -127
                                                                                                                v2 -= -1
                                                                                                                v38 += (v38 * (v38 | v38))
                                                                                                            } else {
                                                                                                                if false {
                                                                                                                    if true {
                                                                                                                        var v87 int16 = v2
                                                                                                                        _ = v87
                                                                                                                    }
                                                                                                                    v36 = v36
                                                                                                                    fmt.Println(fi42, v38, v39)
                                                                                                                } else {
                                                                                                                    if len(v28) >= 2 {
                                                                                                                        v28 = v28[1:len(v28)]
                                                                                                                    }
                                                                                                                    v88 := GenS5{F0: (-765 < -791), F1: (v36 << ((885) & 63)), F2: 256}
                                                                                                                    v72 = 128
                                                                                                                    v3 = v81
                                                                                                                    _ = v88
                                                                                                                }
                                                                                                                switch fi42 {
                                                                                                                case 2147483647:
                                                                                                                    v89 := GenS6{F0: (v81 ^ v3), F1: v2, F2: (v38 / ((v38) | 1))}
                                                                                                                    _ = v89
                                                                                                                }
                                                                                                                if len(v28) >= 2 {
                                                                                                                    v28 = v28[1:len(v28)]
                                                                                                                }
                                                                                                            }
                                                                                                        case 761:
                                                                                                            switch (^(v1)) {
                                                                                                            case 9223372036854775807:
                                                                                                                if ((v36 >= 0) && (1 <= v81)) {
                                                                                                                    switch 0 {
                                                                                                                    case -958:
                                                                                                                        var v90 int = v1
                                                                                                                        var v91 string = "a"
                                                                                                                        _ = v90
                                                                                                                        _ = v91
                                                                                                                    case 1:
                                                                                                                        var v92 int64 = v3
                                                                                                                        var v93 uint = 2
                                                                                                                        _ = v92
                                                                                                                        _ = v93
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    v38 = 211
                                                                                                                    v94 := GenS2{F0: v39, F1: 220}
                                                                                                                    _ = v94
                                                                                                                }
                                                                                                                if false {
                                                                                                                    if (genFn1(-128, ("\n\t" + "a"), (-905 <= 42)) != 42) {
                                                                                                                        var v95 uint = 18446744073709551615
                                                                                                                        _ = v95
                                                                                                                    } else {
                                                                                                                        v67 += v72
                                                                                                                        v67 = (v72 << ((933) & 63))
                                                                                                                        if (v72 >= 0) {
                                                                                                                            if len(v28) > 0 {
                                                                                                                                v28[0] = fi45
                                                                                                                            }
                                                                                                                            if true {
                                                                                                                                var v96 uint64 = 1
                                                                                                                                _ = v96
                                                                                                                            } else {
                                                                                                                                switch "\n\t" {
                                                                                                                                case "":
                                                                                                                                    var v97 int32 = 1
                                                                                                                                    _ = v97
                                                                                                                                case "foo bar":
                                                                                                                                    var v98 int16 = (^(v2))
                                                                                                                                    _ = v98
                                                                                                                                default:
                                                                                                                                    v51 = (v36 >> ((811) & 63))
                                                                                                                                }
                                                                                                                                v39 = v51
                                                                                                                                var v99 int16 = (v2 + v2)
                                                                                                                                _ = v99
                                                                                                                            }
                                                                                                                            if len(v28) >= 2 {
                                                                                                                                v28 = v28[1:len(v28)]
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            if len(v28) >= 2 {
                                                                                                                                v28 = v28[1:len(v28)]
                                                                                                                            }
                                                                                                                            v69 = (v4 / (v69 - (v4 / 3.14159)))
                                                                                                                            fmt.Println(v3, v39)
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                                v100 := GenS5{F0: false, F1: -11, F2: 66}
                                                                                                                _ = v100
                                                                                                            case 32768:
                                                                                                                v101 := GenS8{F0: (v72 | 2), F1: v2}
                                                                                                                _ = v101
                                                                                                            }
                                                                                                        case 127:
                                                                                                            _ = len(v28)
                                                                                                        default:
                                                                                                            v102 := GenS8{F0: 127, F1: 1}
                                                                                                            _ = v102
                                                                                                        }
                                                                                                    default:
                                                                                                        if true {
                                                                                                            _ = len(v28)
                                                                                                        } else {
                                                                                                            v103 := GenS9{F0: (v2 % ((v2) | 1)), F1: (v38 / ((v38) | 1)), F2: v3}
                                                                                                            if (1.5 >= v69) {
                                                                                                                if false {
                                                                                                                    if false {
                                                                                                                        var v104 string = "\n\t"
                                                                                                                        if (42 >= (v38 << ((4294967295) & 63))) {
                                                                                                                            var v105 int64 = 42
                                                                                                                            _ = v105
                                                                                                                        } else {
                                                                                                                            var v106 uint64 = 752
                                                                                                                            _ = v106
                                                                                                                        }
                                                                                                                        _ = v104
                                                                                                                    } else {
                                                                                                                        _ = len(v28)
                                                                                                                        switch 9223372036854775807 {
                                                                                                                        case 958:
                                                                                                                            var v107 int8 = 95
                                                                                                                            _ = v107
                                                                                                                        default:
                                                                                                                            if (0 > 42) {
                                                                                                                                if len(v28) >= 2 {
                                                                                                                                    v28 = v28[1:len(v28)]
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }
                                                                                                                        v69 = 1e10
                                                                                                                        if (653 >= 32767) {
                                                                                                                            if (((4294967295 > 0) || (!(1e10 != 1e10))) && (genFn1(int8(v2), "a", (!("\n\t" <= ("" + ("hello" + "foo bar"))))) >= 388)) {
                                                                                                                                fmt.Println(v69)
                                                                                                                                var v108 int32 = -15
                                                                                                                                if (!(true && (!false))) {
                                                                                                                                    if len(v28) > 0 {
                                                                                                                                        v28[0] = (v1 * fi42)
                                                                                                                                    }
                                                                                                                                    if len(v28) >= 2 {
                                                                                                                                        v28 = v28[1:len(v28)]
                                                                                                                                    }
                                                                                                                                    v109 := GenS10{F0: fi42, F1: v72, F2: "a"}
                                                                                                                                    _ = v109
                                                                                                                                }
                                                                                                                                v110 := GenS11{F0: 1, F1: 0, F2: 1}
                                                                                                                                fmt.Println(v39, v69)
                                                                                                                                _ = v108
                                                                                                                                _ = v110
                                                                                                                            } else {
                                                                                                                                var v111 uint = 4294967295
                                                                                                                                var v112 int8 = (v36 & v39)
                                                                                                                                v113 := GenS12{F0: -1, F1: 635}
                                                                                                                                _ = v111
                                                                                                                                _ = v112
                                                                                                                                _ = v113
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            v2 = v2
                                                                                                                            var v114 uint8 = 128
                                                                                                                            _ = v114
                                                                                                                        }
                                                                                                                        v72 = v72
                                                                                                                    }
                                                                                                                    var v115 uint64 = 10
                                                                                                                    if false {
                                                                                                                        switch 4294967295 {
                                                                                                                        case -9223372036854775808:
                                                                                                                            fmt.Println(v2, v36)
                                                                                                                        case 0:
                                                                                                                            v116 := GenS13{F0: (-(fi42)), F1: 32768}
                                                                                                                            _ = v116
                                                                                                                        default:
                                                                                                                            _ = len(v28)
                                                                                                                        }
                                                                                                                        v28 = append(v28, (v1 << ((42) & 63)))
                                                                                                                        v2 -= int16(v38)
                                                                                                                    } else {
                                                                                                                        switch fi42 {
                                                                                                                        case 243:
                                                                                                                            v117 := GenS6{F0: v81, F1: (v2 & v2), F2: (v38 >> ((42) & 63))}
                                                                                                                            _ = v117
                                                                                                                        case -9223372036854775808:
                                                                                                                            var v118 int32 = genFn1(v36, "a", false)
                                                                                                                            _ = v118
                                                                                                                        default:
                                                                                                                            fmt.Println(v36, v38, v81)
                                                                                                                            var v119 uint8 = (v72 << ((0) & 63))
                                                                                                                            _ = v119
                                                                                                                        }
                                                                                                                        fmt.Println(v3, v2)
                                                                                                                    }
                                                                                                                    if true {
                                                                                                                        v120 := GenS8{F0: v72, F1: (^(v2))}
                                                                                                                        fmt.Println(v1, v4, v72)
                                                                                                                        v115 = v115
                                                                                                                        var v121 int8 = v39
                                                                                                                        v81 += (v81 | (v81 / ((9223372036854775807) | 1)))
                                                                                                                        v38 = 4294967295
                                                                                                                        _ = v120
                                                                                                                        _ = v121
                                                                                                                    } else {
                                                                                                                        var v122 string = "hello"
                                                                                                                        switch "hello" {
                                                                                                                        case "":
                                                                                                                            fmt.Println(fi45)
                                                                                                                        case "a":
                                                                                                                            switch v122 {
                                                                                                                            case "":
                                                                                                                                var v123 uint16 = 1
                                                                                                                                _ = v123
                                                                                                                            case "foo bar":
                                                                                                                                if ((v38 < (^(v38))) || true) {
                                                                                                                                    var v124 float32 = (float32(-2.25) + 1.5)
                                                                                                                                    switch v122 {
                                                                                                                                    case "\n\t":
                                                                                                                                        if true {
                                                                                                                                            v125 := GenS2{F0: int8(v3), F1: v72}
                                                                                                                                            if false {
                                                                                                                                                v39 += (^(v51))
                                                                                                                                                var v126 int16 = v2
                                                                                                                                                _ = v126
                                                                                                                                            } else {
                                                                                                                                                var v127 int64 = v81
                                                                                                                                                if ("hello" >= v122) {
                                                                                                                                                    v81 = (v81 >> ((1) & 63))
                                                                                                                                                    v124 -= v4
                                                                                                                                                    switch v1 {
                                                                                                                                                    case -2147483648:
                                                                                                                                                        var v128 float64 = 3.14159
                                                                                                                                                        _ = v128
                                                                                                                                                    case 32768:
                                                                                                                                                        var v129 int32 = genFn1((v51 * (v36 * 42)), (v122 + "hello"), (v72 >= v67))
                                                                                                                                                        _ = v129
                                                                                                                                                    case -9223372036854775808:
                                                                                                                                                        v130 := GenS14{F0: v122}
                                                                                                                                                        _ = v130
                                                                                                                                                    }
                                                                                                                                                } else {
                                                                                                                                                    var v131 uint32 = v38
                                                                                                                                                    var v132 string = "\n\t"
                                                                                                                                                    if false {
                                                                                                                                                        _ = len(v28)
                                                                                                                                                    }
                                                                                                                                                    v4 -= (v69 / v69)
                                                                                                                                                    if (0.1 >= (3.14159 / v124)) {
                                                                                                                                                        switch 1 {
                                                                                                                                                        case 0:
                                                                                                                                                            if true {
                                                                                                                                                                fmt.Println(v2, v131)
                                                                                                                                                            }
                                                                                                                                                        default:
                                                                                                                                                            v133 := GenS3{F0: v39, F1: false}
                                                                                                                                                            _ = v133
                                                                                                                                                        }
                                                                                                                                                    } else {
                                                                                                                                                        v38 += 2
                                                                                                                                                        v134 := GenS15{F0: (v67 ^ v72), F1: v38, F2: (v38 % ((v131) | 1))}
                                                                                                                                                        v131 = v38
                                                                                                                                                        _ = v134
                                                                                                                                                    }
                                                                                                                                                    _ = v131
                                                                                                                                                    _ = v132
                                                                                                                                                }
                                                                                                                                                fmt.Println(v51, v2, v1)
                                                                                                                                                _ = v127
                                                                                                                                            }
                                                                                                                                            fmt.Println(v2)
                                                                                                                                            _ = v125
                                                                                                                                        } else {
                                                                                                                                            switch v122 {
                                                                                                                                            case "a":
                                                                                                                                                v28 = append(v28, (fi45 - (v1 % ((128) | 1))))
                                                                                                                                            case "\n\t":
                                                                                                                                                v124 = 3.14159
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    _ = v124
                                                                                                                                } else {
                                                                                                                                    _ = len(v28)
                                                                                                                                }
                                                                                                                            case "hello":
                                                                                                                                v135 := GenS6{F0: v81, F1: (v2 / ((-222) | 1)), F2: v38}
                                                                                                                                _ = v135
                                                                                                                            default:
                                                                                                                                if true {
                                                                                                                                    v136 := GenS2{F0: (v51 - v36), F1: 108}
                                                                                                                                    _ = len(v28)
                                                                                                                                    if false {
                                                                                                                                        switch v122 {
                                                                                                                                        case "hello":
                                                                                                                                            if true {
                                                                                                                                                var v137 uint = 855
                                                                                                                                                v2 = v2
                                                                                                                                                fmt.Println(v115)
                                                                                                                                                _ = v137
                                                                                                                                            }
                                                                                                                                        case "":
                                                                                                                                            switch 127 {
                                                                                                                                            case 1:
                                                                                                                                                var v138 uint16 = 1
                                                                                                                                                _ = v138
                                                                                                                                            case -1:
                                                                                                                                                if (false || false) {
                                                                                                                                                    if false {
                                                                                                                                                        if false {
                                                                                                                                                            v139 := GenS7{F0: (v2 - 32767), F1: 18446744073709551615, F2: (v51 << ((2) & 63))}
                                                                                                                                                            var v140 bool = true
                                                                                                                                                            _ = v139
                                                                                                                                                            _ = v140
                                                                                                                                                        } else {
                                                                                                                                                            if (32767 >= 2) {
                                                                                                                                                                var v141 int64 = v81
                                                                                                                                                                var v142 float32 = (3.14159 / v69)
                                                                                                                                                                switch v122 {
                                                                                                                                                                case "foo bar":
                                                                                                                                                                    var v143 uint64 = v115
                                                                                                                                                                    _ = v143
                                                                                                                                                                case "\n\t":
                                                                                                                                                                    v3 -= -1
                                                                                                                                                                case "hello":
                                                                                                                                                                    fmt.Println(v4, v72, v1)
                                                                                                                                                                    var v144 uint = 255
                                                                                                                                                                    var v145 int16 = (v2 + (^(v2)))
                                                                                                                                                                    _ = v144
                                                                                                                                                                    _ = v145
                                                                                                                                                                default:
                                                                                                                                                                    var v146 string = "hello"
                                                                                                                                                                    _ = v146
                                                                                                                                                                }
                                                                                                                                                                var v147 uint32 = (^(v38))
                                                                                                                                                                _ = v141
                                                                                                                                                                _ = v142
                                                                                                                                                                _ = v147
                                                                                                                                                            }
                                                                                                                                                            var v148 int = (fi45 & 127)
                                                                                                                                                            var v149 uint8 = v67
                                                                                                                                                            _ = v148
                                                                                                                                                            _ = v149
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                    var v150 string = v122
                                                                                                                                                    switch v122 {
                                                                                                                                                    case "":
                                                                                                                                                        if false {
                                                                                                                                                            var v151 int32 = 1
                                                                                                                                                            _ = v151
                                                                                                                                                        } else {
                                                                                                                                                            var v152 string = v122
                                                                                                                                                            v36 = 0
                                                                                                                                                            _ = v152
                                                                                                                                                        }
                                                                                                                                                    case "foo bar":
                                                                                                                                                        var v153 uint8 = uint8(v38)
                                                                                                                                                        _ = v153
                                                                                                                                                    case "\n\t":
                                                                                                                                                        v3 -= -9223372036854775808
                                                                                                                                                    default:
                                                                                                                                                        if false {
                                                                                                                                                            var v154 string = (("" + v150) + v150)
                                                                                                                                                            v155 := GenS16{F0: false, F1: 0}
                                                                                                                                                            _ = v154
                                                                                                                                                            _ = v155
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                    var v156 int32 = 2147483647
                                                                                                                                                    var v157 int32 = v156
                                                                                                                                                    _ = v150
                                                                                                                                                    _ = v156
                                                                                                                                                    _ = v157
                                                                                                                                                }
                                                                                                                                            default:
                                                                                                                                                switch v122 {
                                                                                                                                                case "foo bar":
                                                                                                                                                    v67 = v67
                                                                                                                                                    switch fi45 {
                                                                                                                                                    case 4294967295:
                                                                                                                                                        v158 := GenS17{F0: (v3 ^ (v3 * v81))}
                                                                                                                                                        _ = v158
                                                                                                                                                    }
                                                                                                                                                case "\n\t":
                                                                                                                                                    switch 9223372036854775807 {
                                                                                                                                                    case 903:
                                                                                                                                                        var v159 string = (v122 + v122)
                                                                                                                                                        _ = v159
                                                                                                                                                    case 128:
                                                                                                                                                        switch fi42 {
                                                                                                                                                        case -672:
                                                                                                                                                            switch (fi42 / ((v1) | 1)) {
                                                                                                                                                            case 32767:
                                                                                                                                                                switch (v122 + ("hello" + "foo bar")) {
                                                                                                                                                                case "\n\t":
                                                                                                                                                                    v122 = v122
                                                                                                                                                                case "hello":
                                                                                                                                                                    fmt.Println(v36, v36)
                                                                                                                                                                }
                                                                                                                                                            case 1:
                                                                                                                                                                if len(v28) >= 2 {
                                                                                                                                                                    v28 = v28[1:len(v28)]
                                                                                                                                                                }
                                                                                                                                                            }
                                                                                                                                                        default:
                                                                                                                                                            v72 = (^(v67))
                                                                                                                                                        }
                                                                                                                                                    default:
                                                                                                                                                        switch v1 {
                                                                                                                                                        case 826:
                                                                                                                                                            fmt.Println(v69, v39, v81)
                                                                                                                                                            switch v122 {
                                                                                                                                                            case "hello":
                                                                                                                                                                v67 += 127
                                                                                                                                                                v1 += (fi42 >> ((1) & 63))
                                                                                                                                                            case "foo bar":
                                                                                                                                                                var v160 uint8 = v67
                                                                                                                                                                _ = v160
                                                                                                                                                            case "":
                                                                                                                                                                fmt.Println(fi45)
                                                                                                                                                                if len(v28) > 0 {
                                                                                                                                                                    v28[0] = v1
                                                                                                                                                                }
                                                                                                                                                                var v161 int32 = genFn1((^(v51)), v122, (((true || (((-(v3)) >= v3) && false)) && (!(!true))) && ((4294967295 <= v115) && false)))
                                                                                                                                                                _ = v161
                                                                                                                                                            }
                                                                                                                                                            v4 = 0.0
                                                                                                                                                        default:
                                                                                                                                                            v39 = v36
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                case "a":
                                                                                                                                                    if len(v28) > 0 {
                                                                                                                                                        v28[0] = 0
                                                                                                                                                    }
                                                                                                                                                    switch "hello" {
                                                                                                                                                    case "foo bar":
                                                                                                                                                        fmt.Println(v38)
                                                                                                                                                    case "a":
                                                                                                                                                        var v162 uint16 = 0
                                                                                                                                                        _ = v162
                                                                                                                                                    case "\n\t":
                                                                                                                                                        var v163 int = -9223372036854775808
                                                                                                                                                        _ = v163
                                                                                                                                                    }
                                                                                                                                                    if true {
                                                                                                                                                        v72 -= v72
                                                                                                                                                        v164 := GenS18{F0: 43, F1: genFn1(v51, v122, true), F2: ((v122 + "foo bar") + "foo bar")}
                                                                                                                                                        _ = v164
                                                                                                                                                    } else {
                                                                                                                                                        switch (v122 + ("foo bar" + ("a" + v122))) {
                                                                                                                                                        case "a":
                                                                                                                                                            fmt.Println(v2, v122)
                                                                                                                                                        case "foo bar":
                                                                                                                                                            v165 := GenS19{F0: uint8(v115)}
                                                                                                                                                            fmt.Println(v51)
                                                                                                                                                            _ = v165
                                                                                                                                                        default:
                                                                                                                                                            if (4294967295 != 1) {
                                                                                                                                                                if (!true) {
                                                                                                                                                                    v166 := GenS20{F0: 718, F1: (65535 > 65535), F2: -2147483648}
                                                                                                                                                                    v81 -= (v3 + v81)
                                                                                                                                                                    var v167 int8 = int8(v115)
                                                                                                                                                                    _ = v166
                                                                                                                                                                    _ = v167
                                                                                                                                                                }
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    _ = v136
                                                                                                                                }
                                                                                                                            }
                                                                                                                        case "\n\t":
                                                                                                                            v168 := GenS21{F0: -2147483648}
                                                                                                                            _ = v168
                                                                                                                        }
                                                                                                                        if ((v4 / 1.5) >= v69) {
                                                                                                                            v122 = v122
                                                                                                                        }
                                                                                                                        switch fi42 {
                                                                                                                        case -68:
                                                                                                                            if true {
                                                                                                                                var v169 uint64 = (v115 ^ (v115 ^ v115))
                                                                                                                                _ = v169
                                                                                                                            }
                                                                                                                        case 128:
                                                                                                                            v170 := GenS8{F0: 127, F1: (v2 + v2)}
                                                                                                                            _ = v170
                                                                                                                        case 65:
                                                                                                                            fmt.Println(v38)
                                                                                                                        }
                                                                                                                        v122 = v122
                                                                                                                        _ = v122
                                                                                                                    }
                                                                                                                    _ = v115
                                                                                                                }
                                                                                                                fmt.Println(v81, v67, fi45)
                                                                                                            }
                                                                                                            if true {
                                                                                                                fmt.Println(v3, v1, v4)
                                                                                                                var v171 uint16 = 255
                                                                                                                v67 += 237
                                                                                                                _ = v171
                                                                                                            }
                                                                                                            v81 = v81
                                                                                                            _ = v103
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            _ = len(v28)
                                                                                            _ = v81
                                                                                        }
                                                                                    }
                                                                                    v172 := GenS22{F0: "\n\t"}
                                                                                    _ = v172
                                                                                }
                                                                            }
                                                                            var v173 uint32 = (v38 | 4294967295)
                                                                            _ = v173
                                                                            _ = v71
                                                                            _ = v72
                                                                        }
                                                                        v36 += v39
                                                                    }
                                                                case "":
                                                                    v174 := GenS18{F0: 136, F1: 954, F2: "foo bar"}
                                                                    _ = v174
                                                                default:
                                                                    fmt.Println(v36, v3)
                                                                }
                                                                _ = v68
                                                                _ = v69
                                                            } else {
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    var v175 uint32 = (v38 << ((477) & 63))
                                                                    _ = v175
                                                                }
                                                                switch (fi42 >> ((517) & 63)) {
                                                                case -1:
                                                                    fmt.Println(v67)
                                                                default:
                                                                    _ = len(v28)
                                                                }
                                                                if true {
                                                                    switch "" {
                                                                    case "hello":
                                                                        _ = len(v28)
                                                                    case "a":
                                                                        var v176 uint32 = (v38 * v38)
                                                                        _ = v176
                                                                    case "\n\t":
                                                                        var v177 string = "a"
                                                                        _ = v177
                                                                    default:
                                                                        fmt.Println(v36)
                                                                    }
                                                                    switch fi42 {
                                                                    case 9223372036854775807:
                                                                        v178 := GenS23{F0: v3}
                                                                        _ = v178
                                                                    case 1:
                                                                        switch (fi42 - v1) {
                                                                        case 127:
                                                                            v39 = v51
                                                                        case 151:
                                                                            switch (fi42 % (((fi42 % ((0) | 1))) | 1)) {
                                                                            case 1:
                                                                                var v179 int32 = 811
                                                                                _ = v179
                                                                            case 193:
                                                                                var v180 int32 = genFn1(v51, "a", true)
                                                                                _ = v180
                                                                            case 826:
                                                                                if ((false && false) || false) {
                                                                                    v181 := GenS24{F0: ("foo bar" + "hello"), F1: (^(v36))}
                                                                                    _ = v181
                                                                                }
                                                                            default:
                                                                                if (v36 == 42) {
                                                                                    if true {
                                                                                        var v182 int8 = v39
                                                                                        v67 = 1
                                                                                        _ = v182
                                                                                    }
                                                                                    v39 = genFn0(-2147483648)
                                                                                    if len(v28) >= 2 {
                                                                                        v28 = v28[1:len(v28)]
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                    case 271:
                                                                        v183 := GenS7{F0: v2, F1: 42, F2: (v51 % ((1) | 1))}
                                                                        _ = v183
                                                                    }
                                                                    _ = len(v28)
                                                                    fmt.Println(v2, v67)
                                                                    if (true || false) {
                                                                        if len(v28) > 0 {
                                                                            v28[0] = v1
                                                                        }
                                                                        switch fi45 {
                                                                        case -997:
                                                                            var v184 int32 = -905
                                                                            _ = v184
                                                                        case 42:
                                                                            v51 = 105
                                                                        case 1:
                                                                            v185 := GenS17{F0: (v3 - v3)}
                                                                            _ = v185
                                                                        default:
                                                                            switch (fi42 + 912) {
                                                                            case 800:
                                                                                if true {
                                                                                    if ((3.14159 - (-2.25 * 1.5)) >= 1.5) {
                                                                                        v39 += (v39 | v36)
                                                                                        _ = len(v28)
                                                                                    } else {
                                                                                        if ((v38 | (v38 << ((1) & 63))) >= (v38 * (^(v38)))) {
                                                                                            var v186 uint32 = v38
                                                                                            var v187 uint = 0
                                                                                            v188 := GenS2{F0: v36, F1: 65}
                                                                                            if true {
                                                                                                v189 := GenS22{F0: "hello"}
                                                                                                _ = v189
                                                                                            }
                                                                                            var v190 int32 = -421
                                                                                            _ = v186
                                                                                            _ = v187
                                                                                            _ = v188
                                                                                            _ = v190
                                                                                        }
                                                                                        v191 := GenS25{F0: 255}
                                                                                        _ = v191
                                                                                    }
                                                                                    v192 := GenS4{F0: v67}
                                                                                    fmt.Println(v38, v3, v38)
                                                                                    if ((v1 * fi42) <= (v1 / (((v1 - -559)) | 1))) {
                                                                                        var v193 int64 = (v3 | (v3 + v3))
                                                                                        switch (fi45 % (((v1 | fi45)) | 1)) {
                                                                                        case -2147483648:
                                                                                            fmt.Println(fi42)
                                                                                        case -145:
                                                                                            var v194 int = -352
                                                                                            if false {
                                                                                                v195 := GenS17{F0: 9223372036854775807}
                                                                                                _ = len(v28)
                                                                                                _ = v195
                                                                                            }
                                                                                            _ = v194
                                                                                        case 128:
                                                                                            if (((("hello" + "") == "hello") && false) && ((2 <= 0) && (!true))) {
                                                                                                if (4294967295 > 65535) {
                                                                                                    if ((false && false) && false) {
                                                                                                        if len(v28) > 0 {
                                                                                                            v28[0] = -5
                                                                                                        }
                                                                                                        if (v2 == -32768) {
                                                                                                            var v196 uint8 = v67
                                                                                                            v28 = append(v28, v1)
                                                                                                            if false {
                                                                                                                v4 = v4
                                                                                                                if len(v28) > 0 {
                                                                                                                    v28[0] = (fi42 - 0)
                                                                                                                }
                                                                                                                var v197 uint32 = (v38 % (((v38 / (((^(v38))) | 1))) | 1))
                                                                                                                _ = v197
                                                                                                            } else {
                                                                                                                switch (((("a" + "a") + "a") + "\n\t") + "") {
                                                                                                                case "foo bar":
                                                                                                                    if len(v28) >= 2 {
                                                                                                                        v28 = v28[1:len(v28)]
                                                                                                                    }
                                                                                                                case "":
                                                                                                                    v196 = v196
                                                                                                                case "a":
                                                                                                                    var v198 bool = (2147483647 > 520)
                                                                                                                    _ = v198
                                                                                                                default:
                                                                                                                    var v199 uint64 = 168
                                                                                                                    var v200 uint32 = (v38 * v38)
                                                                                                                    var v201 float64 = (0.1 / float64(v4))
                                                                                                                    _ = v199
                                                                                                                    _ = v200
                                                                                                                    _ = v201
                                                                                                                }
                                                                                                                if false {
                                                                                                                    fmt.Println(v39)
                                                                                                                    fmt.Println(v193, fi42, v2)
                                                                                                                    var v202 uint64 = 140
                                                                                                                    _ = v202
                                                                                                                }
                                                                                                                switch v1 {
                                                                                                                case 0:
                                                                                                                    var v203 int32 = -93
                                                                                                                    _ = v203
                                                                                                                case 1:
                                                                                                                    var v204 int64 = v3
                                                                                                                    _ = v204
                                                                                                                case -9223372036854775808:
                                                                                                                    var v205 int32 = 2147483647
                                                                                                                    _ = v205
                                                                                                                }
                                                                                                            }
                                                                                                            if len(v28) > 0 {
                                                                                                                v28[0] = -996
                                                                                                            }
                                                                                                            _ = v196
                                                                                                        } else {
                                                                                                            v206 := GenS6{F0: (-(v3)), F1: (v2 * v2), F2: (v38 / ((2147483647) | 1))}
                                                                                                            v51 = v39
                                                                                                            _ = v206
                                                                                                        }
                                                                                                        v38 += (v38 << ((255) & 63))
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            switch fi45 {
                                                                                            case 128:
                                                                                                switch "a" {
                                                                                                case "a":
                                                                                                    var v207 uint16 = uint16(v38)
                                                                                                    _ = v207
                                                                                                case "hello":
                                                                                                    if true {
                                                                                                        var v208 uint32 = 199
                                                                                                        switch (^(fi45)) {
                                                                                                        case 42:
                                                                                                            if ((v193 << ((0) & 63)) > (-(v3))) {
                                                                                                                switch fi42 {
                                                                                                                case -9223372036854775808:
                                                                                                                    var v209 uint64 = 804
                                                                                                                    _ = v209
                                                                                                                case 32767:
                                                                                                                    if false {
                                                                                                                        var v210 uint16 = 0
                                                                                                                        _ = v210
                                                                                                                    } else {
                                                                                                                        var v211 uint32 = (v208 << ((264) & 63))
                                                                                                                        v193 -= 42
                                                                                                                        if (!true) {
                                                                                                                            var v212 string = ("foo bar" + ("a" + "\n\t"))
                                                                                                                            if false {
                                                                                                                                fmt.Println(v212, v193, v51)
                                                                                                                                fmt.Println(v38)
                                                                                                                                switch v212 {
                                                                                                                                case "":
                                                                                                                                    v213 := GenS26{F0: (v2 >> ((1) & 63)), F1: 384}
                                                                                                                                    _ = v213
                                                                                                                                case "hello":
                                                                                                                                    var v214 int = (fi42 >> ((211) & 63))
                                                                                                                                    _ = v214
                                                                                                                                case "\n\t":
                                                                                                                                    var v215 int = (fi42 << ((255) & 63))
                                                                                                                                    _ = v215
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                v38 = v211
                                                                                                                                v51 = int8(fi45)
                                                                                                                            }
                                                                                                                            if (v2 >= (v2 & v2)) {
                                                                                                                                v36 = v39
                                                                                                                            } else {
                                                                                                                                v193 = -386
                                                                                                                                fmt.Println(v67, v39)
                                                                                                                            }
                                                                                                                            _ = v212
                                                                                                                        } else {
                                                                                                                            switch (fi42 >> ((2) & 63)) {
                                                                                                                            case 42:
                                                                                                                                v4 -= v4
                                                                                                                            case 510:
                                                                                                                                v193 = (v193 - (v3 << ((895) & 63)))
                                                                                                                            default:
                                                                                                                                v51 = -87
                                                                                                                            }
                                                                                                                            v216 := GenS27{F0: 1, F1: 1}
                                                                                                                            _ = v216
                                                                                                                        }
                                                                                                                        var v217 float64 = 3.14159
                                                                                                                        _ = v211
                                                                                                                        _ = v217
                                                                                                                    }
                                                                                                                default:
                                                                                                                    if false {
                                                                                                                        fmt.Println(v1)
                                                                                                                        switch 32768 {
                                                                                                                        case -9223372036854775808:
                                                                                                                            fmt.Println(v36, v67)
                                                                                                                        case -1:
                                                                                                                            var v218 uint = 694
                                                                                                                            _ = v218
                                                                                                                        case 0:
                                                                                                                            var v219 uint64 = 4294967295
                                                                                                                            if (true || true) {
                                                                                                                                v220 := GenS1{F0: v39, F1: ("foo bar" + ("hello" + "a"))}
                                                                                                                                var v221 int = fi45
                                                                                                                                v222 := GenS19{F0: v67}
                                                                                                                                _ = v220
                                                                                                                                _ = v221
                                                                                                                                _ = v222
                                                                                                                            } else {
                                                                                                                                v2 -= 42
                                                                                                                                var v223 int64 = v193
                                                                                                                                _ = v223
                                                                                                                            }
                                                                                                                            if len(v28) >= 2 {
                                                                                                                                v28 = v28[1:len(v28)]
                                                                                                                            }
                                                                                                                            _ = v219
                                                                                                                        }
                                                                                                                        var v224 int16 = v2
                                                                                                                        _ = v224
                                                                                                                    } else {
                                                                                                                        switch "hello" {
                                                                                                                        case "a":
                                                                                                                            var v225 uint = 65535
                                                                                                                            _ = v225
                                                                                                                        case "":
                                                                                                                            switch 0 {
                                                                                                                            case -2147483648:
                                                                                                                                var v226 int16 = (v2 & (v2 % ((-1) | 1)))
                                                                                                                                _ = v226
                                                                                                                            }
                                                                                                                        case "hello":
                                                                                                                            switch (("hello" + "foo bar") + ("\n\t" + "hello")) {
                                                                                                                            case "a":
                                                                                                                                v38 -= 243
                                                                                                                                v36 += 1
                                                                                                                                v51 = (v39 % (((v51 + v51)) | 1))
                                                                                                                            case "hello":
                                                                                                                                var v227 int = fi45
                                                                                                                                _ = v227
                                                                                                                            case "":
                                                                                                                                var v228 int64 = -219
                                                                                                                                _ = v228
                                                                                                                            default:
                                                                                                                                v39 = v51
                                                                                                                            }
                                                                                                                        }
                                                                                                                        if (1 == 1) {
                                                                                                                            v229 := GenS28{F0: (v193 % ((v3) | 1)), F1: -2147483648}
                                                                                                                            if true {
                                                                                                                                fmt.Println(fi45)
                                                                                                                                var v230 float32 = (float32(3.14159) - float32(float64((v4 + (((((3.14159 / v4) / ((1e10 - v4) + float32(((1e10 * (float64(0.1) / ((-2.25 + -2.25) * -2.25))) + ((((-2.25 / 1.0) - 1e10) - (0.1 - 1e10)) / 1.5))))) * 1.5) + (v4 + v4)) + float32(0.0))))))
                                                                                                                                if true {
                                                                                                                                    v51 += v39
                                                                                                                                    if len(v28) > 0 {
                                                                                                                                        v28[0] = v1
                                                                                                                                    }
                                                                                                                                    var v231 uint8 = v67
                                                                                                                                    _ = v231
                                                                                                                                } else {
                                                                                                                                    if len(v28) >= 2 {
                                                                                                                                        v28 = v28[1:len(v28)]
                                                                                                                                    }
                                                                                                                                    fmt.Println(v230)
                                                                                                                                    switch ((("\n\t" + "foo bar") + ("\n\t" + "a")) + "") {
                                                                                                                                    case "hello":
                                                                                                                                        if (605 == 255) {
                                                                                                                                            v67 -= v67
                                                                                                                                            fmt.Println(fi42, v208, v208)
                                                                                                                                        }
                                                                                                                                    case "\n\t":
                                                                                                                                        _ = len(v28)
                                                                                                                                    case "":
                                                                                                                                        v232 := GenS15{F0: v67, F1: 995, F2: (v38 ^ v38)}
                                                                                                                                        _ = v232
                                                                                                                                    default:
                                                                                                                                        if false {
                                                                                                                                            if (1 >= 42) {
                                                                                                                                                var v233 int32 = int32(v2)
                                                                                                                                                if (1.5 < 1e10) {
                                                                                                                                                    v234 := GenS14{F0: ("foo bar" + "\n\t")}
                                                                                                                                                    switch "a" {
                                                                                                                                                    case "a":
                                                                                                                                                        if (-605 < (v3 >> ((255) & 63))) {
                                                                                                                                                            var v235 float32 = 1.5
                                                                                                                                                            var v236 int16 = 768
                                                                                                                                                            _ = v235
                                                                                                                                                            _ = v236
                                                                                                                                                        }
                                                                                                                                                    case "hello":
                                                                                                                                                        if len(v28) >= 2 {
                                                                                                                                                            v28 = v28[1:len(v28)]
                                                                                                                                                        }
                                                                                                                                                    default:
                                                                                                                                                        if ((v2 / (((-(v2))) | 1)) < v2) {
                                                                                                                                                            fmt.Println(fi42, v230, v1)
                                                                                                                                                            v237 := GenS14{F0: ("a" + "a")}
                                                                                                                                                            _ = v237
                                                                                                                                                        } else {
                                                                                                                                                            if len(v28) > 0 {
                                                                                                                                                                v28[0] = (fi42 * (^(fi42)))
                                                                                                                                                            }
                                                                                                                                                            switch "\n\t" {
                                                                                                                                                            case "a":
                                                                                                                                                                v2 += v2
                                                                                                                                                            }
                                                                                                                                                            v67 = v67
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                    v238 := GenS29{F0: v39, F1: 65535, F2: 0}
                                                                                                                                                    _ = v234
                                                                                                                                                    _ = v238
                                                                                                                                                } else {
                                                                                                                                                    v239 := GenS14{F0: "a"}
                                                                                                                                                    v240 := GenS20{F0: 18446744073709551615, F1: false, F2: (v233 << ((18446744073709551615) & 63))}
                                                                                                                                                    _ = v239
                                                                                                                                                    _ = v240
                                                                                                                                                }
                                                                                                                                                switch ("a" + "") {
                                                                                                                                                case "":
                                                                                                                                                    if false {
                                                                                                                                                        var v241 int8 = (v51 << ((65535) & 63))
                                                                                                                                                        _ = v241
                                                                                                                                                    }
                                                                                                                                                case "a":
                                                                                                                                                    var v242 uint = uint(v208)
                                                                                                                                                    _ = v242
                                                                                                                                                }
                                                                                                                                                _ = v233
                                                                                                                                            } else {
                                                                                                                                                v243 := GenS30{F0: genFn0(1)}
                                                                                                                                                _ = v243
                                                                                                                                            }
                                                                                                                                            v244 := GenS26{F0: (v2 >> ((0) & 63)), F1: 622}
                                                                                                                                            var v245 string = ((("a" + ("" + "a")) + "hello") + (((("" + "") + "") + (("foo bar" + ("a" + ("foo bar" + "foo bar"))) + "\n\t")) + "a"))
                                                                                                                                            _ = v244
                                                                                                                                            _ = v245
                                                                                                                                        } else {
                                                                                                                                            fmt.Println(v193, v193)
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                _ = v230
                                                                                                                            } else {
                                                                                                                                _ = len(v28)
                                                                                                                            }
                                                                                                                            if false {
                                                                                                                                var v246 uint32 = 0
                                                                                                                                v67 -= 2
                                                                                                                                var v247 uint = 969
                                                                                                                                _ = v246
                                                                                                                                _ = v247
                                                                                                                            }
                                                                                                                            v248 := GenS3{F0: 32, F1: (!(32767 >= 1000))}
                                                                                                                            _ = v229
                                                                                                                            _ = v248
                                                                                                                        } else {
                                                                                                                            v249 := GenS19{F0: v67}
                                                                                                                            _ = v249
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                        case -504:
                                                                                                            if true {
                                                                                                                v193 -= v193
                                                                                                                switch 9223372036854775807 {
                                                                                                                case 790:
                                                                                                                    v38 += v208
                                                                                                                default:
                                                                                                                    var v250 int64 = 1
                                                                                                                    _ = v250
                                                                                                                }
                                                                                                            } else {
                                                                                                                v36 -= 42
                                                                                                                fmt.Println(v193, fi42, v3)
                                                                                                                v251 := GenS19{F0: v67}
                                                                                                                if true {
                                                                                                                    v1 = (v1 & fi42)
                                                                                                                    v252 := GenS31{F0: ((v39 - 100) != v51)}
                                                                                                                    var v253 int64 = v193
                                                                                                                    _ = v252
                                                                                                                    _ = v253
                                                                                                                }
                                                                                                                var v254 float32 = v4
                                                                                                                _ = v251
                                                                                                                _ = v254
                                                                                                            }
                                                                                                        case 4294967295:
                                                                                                            v255 := GenS32{F0: -466, F1: (^(v67)), F2: 42}
                                                                                                            if false {
                                                                                                                if ((v193 >= 214) || true) {
                                                                                                                    var v256 uint64 = 0
                                                                                                                    _ = v256
                                                                                                                } else {
                                                                                                                    v67 += 128
                                                                                                                    fmt.Println(v36)
                                                                                                                    v193 = int64(v67)
                                                                                                                }
                                                                                                                switch "\n\t" {
                                                                                                                case "\n\t":
                                                                                                                    v257 := GenS16{F0: true, F1: int(v3)}
                                                                                                                    _ = v257
                                                                                                                case "a":
                                                                                                                    var v258 bool = ("foo bar" == ("hello" + ("" + ((("" + "foo bar") + "a") + ("a" + "foo bar")))))
                                                                                                                    _ = v258
                                                                                                                case "hello":
                                                                                                                    if true {
                                                                                                                        v4 += ((v4 / 0.1) / 1.0)
                                                                                                                        var v259 float64 = 3.14159
                                                                                                                        v260 := GenS16{F0: (744 == 912), F1: v1}
                                                                                                                        if (1 == fi42) {
                                                                                                                            v193 -= v3
                                                                                                                            fmt.Println(v4)
                                                                                                                            v261 := GenS33{F0: 18446744073709551615, F1: 988, F2: 47}
                                                                                                                            _ = v261
                                                                                                                        }
                                                                                                                        _ = v259
                                                                                                                        _ = v260
                                                                                                                    }
                                                                                                                }
                                                                                                                switch 127 {
                                                                                                                case -1:
                                                                                                                    if (1e10 < (0.1 * 1.5)) {
                                                                                                                        v262 := GenS12{F0: 1, F1: v38}
                                                                                                                        if false {
                                                                                                                            if true {
                                                                                                                                var v263 uint32 = v208
                                                                                                                                v28 = append(v28, v1)
                                                                                                                                _ = v263
                                                                                                                            } else {
                                                                                                                                switch "\n\t" {
                                                                                                                                case "foo bar":
                                                                                                                                    if (!false) {
                                                                                                                                        v193 -= v3
                                                                                                                                        if len(v28) > 0 {
                                                                                                                                            v28[0] = -9223372036854775808
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                case "a":
                                                                                                                                    var v264 uint64 = 42
                                                                                                                                    _ = v264
                                                                                                                                case "":
                                                                                                                                    var v265 uint32 = v38
                                                                                                                                    _ = v265
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }
                                                                                                                        var v266 int32 = genFn1(v51, "hello", false)
                                                                                                                        var v267 uint64 = 993
                                                                                                                        _ = v262
                                                                                                                        _ = v266
                                                                                                                        _ = v267
                                                                                                                    }
                                                                                                                    v193 += v193
                                                                                                                    if false {
                                                                                                                        if len(v28) >= 2 {
                                                                                                                            v28 = v28[1:len(v28)]
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        _ = len(v28)
                                                                                                                        switch "a" {
                                                                                                                        case "hello":
                                                                                                                            switch v1 {
                                                                                                                            case 894:
                                                                                                                                v208 = 371
                                                                                                                            case -1:
                                                                                                                                if (true || (0 != 42)) {
                                                                                                                                    v268 := GenS34{F0: v39}
                                                                                                                                    _ = len(v28)
                                                                                                                                    _ = v268
                                                                                                                                } else {
                                                                                                                                    if len(v28) >= 2 {
                                                                                                                                        v28 = v28[1:len(v28)]
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            case 127:
                                                                                                                                var v269 int = 42
                                                                                                                                _ = v269
                                                                                                                            default:
                                                                                                                                var v270 uint32 = 2147483647
                                                                                                                                _ = v270
                                                                                                                            }
                                                                                                                        default:
                                                                                                                            v271 := GenS30{F0: (^(v36))}
                                                                                                                            _ = v271
                                                                                                                        }
                                                                                                                    }
                                                                                                                case 0:
                                                                                                                    var v272 int16 = v2
                                                                                                                    _ = v272
                                                                                                                default:
                                                                                                                    v3 = (v193 % (((^(v193))) | 1))
                                                                                                                }
                                                                                                            }
                                                                                                            v39 = genFn0(2147483647)
                                                                                                            _ = v255
                                                                                                        }
                                                                                                        if (-2147483648 != -1) {
                                                                                                            fmt.Println(v67, fi45, v4)
                                                                                                            if false {
                                                                                                                if len(v28) > 0 {
                                                                                                                    v28[0] = v1
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v208
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        _ = v193
                                                                                    } else {
                                                                                        switch fi42 {
                                                                                        case 748:
                                                                                            if false {
                                                                                                if (42 <= 1) {
                                                                                                    var v273 bool = (!false)
                                                                                                    _ = v273
                                                                                                } else {
                                                                                                    fmt.Println(v4)
                                                                                                    if len(v28) > 0 {
                                                                                                        v28[0] = 0
                                                                                                    }
                                                                                                    v39 = v51
                                                                                                }
                                                                                                if false {
                                                                                                    if (true || (-2147483648 < (v1 | fi45))) {
                                                                                                        if len(v28) >= 2 {
                                                                                                            v28 = v28[1:len(v28)]
                                                                                                        }
                                                                                                    } else {
                                                                                                        v36 -= v39
                                                                                                    }
                                                                                                } else {
                                                                                                    switch fi42 {
                                                                                                    case 9223372036854775807:
                                                                                                        switch "foo bar" {
                                                                                                        case "\n\t":
                                                                                                            v274 := GenS13{F0: (v1 * int(v2)), F1: (v1 | int(v3))}
                                                                                                            _ = v274
                                                                                                        case "hello":
                                                                                                            v39 += (v39 | -1)
                                                                                                        default:
                                                                                                            v38 = uint32(v67)
                                                                                                        }
                                                                                                    case -1:
                                                                                                        if ((false && true) && false) {
                                                                                                            if len(v28) >= 2 {
                                                                                                                v28 = v28[1:len(v28)]
                                                                                                            }
                                                                                                        } else {
                                                                                                            var v275 float32 = (((v4 * v4) + 0.1) + v4)
                                                                                                            v67 = 255
                                                                                                            v38 -= uint32(v67)
                                                                                                            _ = v275
                                                                                                        }
                                                                                                    case 1:
                                                                                                        v276 := GenS8{F0: (v67 >> ((148) & 63)), F1: (v2 % ((v2) | 1))}
                                                                                                        _ = v276
                                                                                                    }
                                                                                                    v39 = -52
                                                                                                }
                                                                                                fmt.Println(v67)
                                                                                            }
                                                                                        case 498:
                                                                                            var v277 float64 = 1.5
                                                                                            _ = v277
                                                                                        case -452:
                                                                                            v278 := GenS24{F0: (("\n\t" + "\n\t") + ("hello" + "hello")), F1: int8(v67)}
                                                                                            _ = v278
                                                                                        }
                                                                                        v36 += v51
                                                                                        if (!(!((v3 << ((218) & 63)) != v3))) {
                                                                                            v279 := GenS34{F0: (v39 >> ((1) & 63))}
                                                                                            _ = v279
                                                                                        } else {
                                                                                            var v280 uint32 = (v38 >> ((1) & 63))
                                                                                            _ = v280
                                                                                        }
                                                                                        var v281 float64 = -2.25
                                                                                        var v282 int = (fi45 & v1)
                                                                                        _ = v281
                                                                                        _ = v282
                                                                                    }
                                                                                    switch (((("a" + "\n\t") + (("foo bar" + "a") + ("hello" + "hello"))) + "foo bar") + ("hello" + "\n\t")) {
                                                                                    case "foo bar":
                                                                                        var v283 float32 = v4
                                                                                        _ = v283
                                                                                    case "hello":
                                                                                        v28 = append(v28, int(v3))
                                                                                    case "\n\t":
                                                                                        fmt.Println(v67)
                                                                                    }
                                                                                    _ = v192
                                                                                } else {
                                                                                    if len(v28) >= 2 {
                                                                                        v28 = v28[1:len(v28)]
                                                                                    }
                                                                                }
                                                                            case 32768:
                                                                                v2 = -208
                                                                            default:
                                                                                if (((v3 & v3) == int64(v67)) && ((v39 << ((430) & 63)) > 84)) {
                                                                                    v1 += (fi42 * 0)
                                                                                    var v284 float32 = v4
                                                                                    v285 := GenS35{F0: 255, F1: (^(v3)), F2: 42}
                                                                                    fmt.Println(v36, v284)
                                                                                    v51 = genFn0(1)
                                                                                    v3 += v3
                                                                                    var v286 float64 = -2.25
                                                                                    var v287 int32 = -469
                                                                                    if len(v28) > 0 {
                                                                                        v28[0] = fi45
                                                                                    }
                                                                                    _ = v284
                                                                                    _ = v285
                                                                                    _ = v286
                                                                                    _ = v287
                                                                                }
                                                                            }
                                                                        }
                                                                        v51 = (v36 % ((int8(v67)) | 1))
                                                                    } else {
                                                                        if false {
                                                                            v3 -= -9223372036854775808
                                                                            if len(v28) > 0 {
                                                                                v28[0] = fi42
                                                                            }
                                                                        }
                                                                        fmt.Println(v67)
                                                                        fmt.Println(v4, v39, v2)
                                                                    }
                                                                    v28 = append(v28, v1)
                                                                    switch (fi45 % (((fi45 + (^(v1)))) | 1)) {
                                                                    case 42:
                                                                        if (0 >= v67) {
                                                                            if len(v28) >= 2 {
                                                                                v28 = v28[1:len(v28)]
                                                                            }
                                                                            var v288 float32 = v4
                                                                            v1 = 148
                                                                            _ = v288
                                                                        }
                                                                    case 844:
                                                                        var v289 int64 = v3
                                                                        _ = v289
                                                                    default:
                                                                        var v290 int8 = v39
                                                                        _ = v290
                                                                    }
                                                                    fmt.Println(v3, fi42)
                                                                }
                                                            }
                                                            v291 := GenS36{F0: 2}
                                                            _ = v291
                                                        } else {
                                                            if (((fi45 | v1) < (v1 << ((42) & 63))) || true) {
                                                                if (!false) {
                                                                    var v292 int16 = v2
                                                                    _ = v292
                                                                }
                                                            }
                                                            v293 := GenS16{F0: true, F1: (fi45 & (fi45 | v1))}
                                                            fmt.Println(v3)
                                                            fmt.Println(fi42, fi45)
                                                            _ = v293
                                                        }
                                                        v67 = uint8(v38)
                                                    } else {
                                                        v2 -= int16(v36)
                                                        if false {
                                                            fmt.Println(v1, v1, v3)
                                                            v39 = v39
                                                            var v294 int64 = -9223372036854775808
                                                            _ = v294
                                                        } else {
                                                            if false {
                                                                if len(v28) > 0 {
                                                                    v28[0] = (fi45 + fi42)
                                                                }
                                                                v38 = v38
                                                                fmt.Println(v67, v3)
                                                            }
                                                            v36 -= 0
                                                        }
                                                        var v295 uint16 = 768
                                                        _ = v295
                                                    }
                                                }
                                                _ = v67
                                            }
                                        }
                                    }
                                    switch fi45 {
                                    case -1:
                                        switch (("hello" + ("hello" + "\n\t")) + "a") {
                                        case "hello":
                                            v296 := GenS32{F0: 2147483647, F1: 1, F2: 95}
                                            _ = v296
                                        case "":
                                            v28 = append(v28, (^(fi42)))
                                        case "foo bar":
                                            _ = len(v28)
                                        }
                                    case -166:
                                        if true {
                                            _ = len(v28)
                                            v297 := GenS37{F0: true, F1: (v2 >> ((18446744073709551615) & 63)), F2: 255}
                                            _ = v297
                                        }
                                    case 42:
                                        if true {
                                            if (!true) {
                                                v39 += v39
                                                v2 += (v2 % ((-663) | 1))
                                                _ = len(v28)
                                                v298 := GenS36{F0: 147}
                                                v28 = append(v28, (fi45 ^ 127))
                                                _ = v298
                                            } else {
                                                v299 := GenS38{F0: int32(v51)}
                                                v39 = (v51 * genFn0(2147483647))
                                                _ = v299
                                            }
                                            if ((fi42 != 4294967295) && (v4 < float32(3.14159))) {
                                                v3 += v3
                                            } else {
                                                var v300 uint = 1
                                                _ = v300
                                            }
                                        }
                                    }
                                }
                            }
                            var v301 int64 = 226
                            _ = v301
                        } else {
                            v39 = v36
                        }
                        _ = v51
                    default:
                        var v302 int8 = (v36 * genFn0(475))
                        _ = v302
                    }
                    if len(v28) >= 2 {
                        v28 = v28[1:len(v28)]
                    }
                    v303 := GenS5{F0: (-905 == -1), F1: (v36 ^ v36), F2: 1}
                    _ = v303
                }
                v1 = 42
                if false {
                    v304 := GenS2{F0: (^(v39)), F1: 248}
                    if true {
                        fmt.Println(v39, v3)
                        var v305 uint64 = 18446744073709551615
                        v305 += (v305 >> ((4294967295) & 63))
                        _ = v305
                    }
                    _ = v304
                }
                _ = fi45
            }
            v36 = v39
        }
        var v306 uint16 = 1
        fmt.Println(v1, v38, v36)
        _ = fi42
        _ = v306
    }
    v3 += v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
