package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16; F1 bool }
type GenS7 struct { F0 int8; F1 uint64; F2 int32 }
type GenS8 struct { F0 int64; F1 uint; F2 int64 }
type GenS9 struct { F0 int; F1 string; F2 int8 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int64; F1 int32; F2 string }
type GenS12 struct { F0 uint; F1 int; F2 int32 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 bool; F1 bool }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 uint8; F1 uint64 }
func genFn0(p0 int16, p1 uint) uint16 {
    fmt.Println(p0, p0)
    if (0 == 42) {
        var v7 int64 = -744
        v8 := GenS1{F0: 42, F1: true}
        var v9 int64 = -657
        _ = v7
        _ = v8
        _ = v9
    } else {
        p1 -= (p1 >> (((p1 | (p1 % (((p1 % ((p1) | 1))) | 1)))) & 63))
        var v10 int = 42
        _ = v10
    }
    if false {
        for fi11 := 0; fi11 < 11; fi11++ {
            p0 = p0
            var v12 uint = (p1 ^ (^(p1)))
            _ = fi11
            _ = v12
        }
        fmt.Println(p1)
        var v13 uint8 = 127
        _ = v13
    }
    _ = p0
    _ = p1
    return 2
}

func main() {
    var v1 bool = false
    v2 := []uint32{4294967295}
    var v3 uint64 = 424
    var v4 uint64 = v3
    v5 := []uint32{}
    v6 := []int{731, 1, 32767, 1}
    var v14 int16 = -740
    if v1 {
        switch "foo bar" {
        case "":
            var v15 float32 = 0.0
            _ = v15
        case "a":
            switch 4294967295 {
            case 537:
                for fi16 := 0; fi16 < 11; fi16++ {
                    for fi17 := 0; fi17 < 15; fi17++ {
                        switch (("foo bar" + ("" + ("a" + ((("hello" + ("\n\t" + ((("foo bar" + "a") + "hello") + "hello"))) + (((("" + "") + (("\n\t" + "\n\t") + ("" + "a"))) + "a") + "")) + (((("" + "") + (("hello" + ("hello" + "")) + "")) + ("hello" + "hello")) + "\n\t"))))) + "") {
                        case "":
                            if true {
                                if false {
                                    v3 = v3
                                    var v18 string = "a"
                                    v19 := GenS1{F0: (-(v14)), F1: (808 == 65535)}
                                    _ = v18
                                    _ = v19
                                }
                                v3 += v4
                                var v20 int32 = -699
                                _ = v20
                            }
                        case "a":
                            var v21 uint64 = 18446744073709551615
                            v22 := GenS1{F0: (v14 + (v14 + 1)), F1: v1}
                            _ = v21
                            _ = v22
                        case "hello":
                            v3 += v3
                        }
                        v1 = false
                        _ = fi17
                    }
                    if (v4 != (v3 + (^(v4)))) {
                        if len(v2) > 0 {
                            v2[0] = 590
                        }
                        v23 := GenS2{F0: fi16, F1: 49}
                        _ = v23
                    }
                    var v24 string = "\n\t"
                    _ = fi16
                    _ = v24
                }
            }
        case "\n\t":
            if (v1 || (964 > 194)) {
                v25 := GenS2{F0: 2147483647, F1: (v3 ^ (v4 / ((v4) | 1)))}
                v14 = v14
                v14 += (^(v14))
                _ = v25
            }
        }
        if ((v4 * 42) > (v3 / ((v4) | 1))) {
            if v1 {
                var v26 float32 = ((0.0 / float32((3.14159 * -2.25))) - -2.25)
                for fi27 := 0; fi27 < 14; fi27++ {
                    if v1 {
                        var v28 uint16 = 2
                        v28 -= v28
                        if (1 < int32(v28)) {
                            v29 := GenS3{F0: (v28 % ((v28) | 1)), F1: -128}
                            var v30 uint8 = 0
                            _ = v29
                            _ = v30
                        }
                        _ = v28
                    } else {
                        switch (fi27 | 128) {
                        case -268:
                            switch ("\n\t" + ("\n\t" + "")) {
                            case "a":
                                for fi31 := 0; fi31 < 15; fi31++ {
                                    if v1 {
                                        if (4294967295 < 2) {
                                            _ = len(v6)
                                            if v1 {
                                                if v1 {
                                                    var v32 uint8 = 79
                                                    v32 = v32
                                                    v32 -= 1
                                                    _ = v32
                                                }
                                                if len(v2) > 0 {
                                                    v2[0] = 921
                                                }
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            }
                                        } else {
                                            v6 = append(v6, (^(fi31)))
                                        }
                                    }
                                    fmt.Println(v3)
                                    if v1 {
                                        _ = len(v6)
                                    }
                                    _ = fi31
                                }
                            case "hello":
                                _ = len(v6)
                            }
                        default:
                            for fi33 := 0; fi33 < 5; fi33++ {
                                if (v14 <= (v14 & -32768)) {
                                    if (!false) {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        v34 := GenS4{F0: 514, F1: (v4 & (^(v4)))}
                                        var v35 uint64 = (v3 | (v4 * 4294967295))
                                        _ = v34
                                        _ = v35
                                    } else {
                                        switch "\n\t" {
                                        case "foo bar":
                                            v1 = false
                                            if v1 {
                                                var v36 uint8 = 127
                                                _ = v36
                                            } else {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v37 int16 = -101
                                                    _ = v37
                                                case "hello":
                                                    var v38 uint32 = 2147483647
                                                    _ = v38
                                                default:
                                                    var v39 int8 = 42
                                                    _ = v39
                                                }
                                            }
                                            var v40 int32 = 2147483647
                                            var v41 uint64 = (^(v4))
                                            switch "foo bar" {
                                            case "foo bar":
                                                if v1 {
                                                    fmt.Println(v14, fi33, v41)
                                                } else {
                                                    switch "foo bar" {
                                                    case "":
                                                        v42 := GenS2{F0: -563, F1: (v41 / ((1) | 1))}
                                                        _ = v42
                                                    }
                                                    _ = len(v2)
                                                }
                                            case "a":
                                                if v1 {
                                                    v1 = v1
                                                    if (0 >= 130) {
                                                        if false {
                                                            if ((1 < 42) || true) {
                                                                v2 = append(v2, 2)
                                                                fmt.Println(v14, v41)
                                                                v3 += (v3 * 18446744073709551615)
                                                            } else {
                                                                if len(v6) > 0 {
                                                                    v6[0] = fi33
                                                                }
                                                                if (-128 != 127) {
                                                                    v43 := GenS1{F0: (v14 * (v14 | 42)), F1: (959 <= -9223372036854775808)}
                                                                    _ = v43
                                                                }
                                                            }
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                            var v44 uint8 = 1
                                                            _ = v44
                                                        } else {
                                                            v45 := GenS3{F0: 163, F1: 42}
                                                            _ = v45
                                                        }
                                                    } else {
                                                        v46 := GenS4{F0: v4, F1: 823}
                                                        _ = v46
                                                    }
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        switch "" {
                                                        case "\n\t":
                                                            switch "" {
                                                            case "":
                                                                fmt.Println(v3, v1)
                                                            case "hello":
                                                                if ("foo bar" == ("a" + "hello")) {
                                                                    if len(v6) > 0 {
                                                                        v6[0] = fi33
                                                                    }
                                                                    v47 := GenS3{F0: 0, F1: 1}
                                                                    v48 := GenS4{F0: (v41 + (v3 & v4)), F1: 1}
                                                                    fmt.Println(v26, v26, v41)
                                                                    _ = v47
                                                                    _ = v48
                                                                } else {
                                                                    v49 := GenS4{F0: (v4 >> ((42) & 63)), F1: v3}
                                                                    v50 := GenS2{F0: (fi33 - (-(fi27))), F1: v41}
                                                                    _ = v49
                                                                    _ = v50
                                                                }
                                                            default:
                                                                v4 -= (^(v41))
                                                            }
                                                        case "hello":
                                                            if ((42 != 42) && v1) {
                                                                var v51 int16 = 0
                                                                _ = v51
                                                            }
                                                        }
                                                    case "a":
                                                        v1 = (fi27 > -1)
                                                    case "foo bar":
                                                        v14 += (v14 % (((v14 << ((260) & 63))) | 1))
                                                    default:
                                                        if (133 < 2) {
                                                            if (-128 < 0) {
                                                                if v1 {
                                                                    v52 := GenS5{F0: (!(v1 || (108 >= 42)))}
                                                                    v53 := GenS4{F0: (v41 >> ((255) & 63)), F1: v3}
                                                                    var v54 int = fi27
                                                                    _ = v52
                                                                    _ = v53
                                                                    _ = v54
                                                                } else {
                                                                    var v55 bool = (!true)
                                                                    v56 := GenS5{F0: (!(false || v55))}
                                                                    _ = v55
                                                                    _ = v56
                                                                }
                                                                v57 := GenS1{F0: v14, F1: ((967 != 18446744073709551615) || v1)}
                                                                _ = v57
                                                            } else {
                                                                if (1 != 136) {
                                                                    fmt.Println(v4, fi27)
                                                                }
                                                                switch "hello" {
                                                                case "a":
                                                                    v5 = append(v5, 2)
                                                                case "\n\t":
                                                                    if true {
                                                                        switch "foo bar" {
                                                                        case "foo bar":
                                                                            v41 = (v3 + (v4 + v4))
                                                                        case "hello":
                                                                            if len(v6) > 0 {
                                                                                v6[0] = fi33
                                                                            }
                                                                        case "a":
                                                                            switch "foo bar" {
                                                                            case "\n\t":
                                                                                if (v1 && (!(!(false || v1)))) {
                                                                                    if (271 < 32767) {
                                                                                        fmt.Println(v3)
                                                                                        var v58 float32 = 1.5
                                                                                        fmt.Println(v4)
                                                                                        if ("\n\t" >= ("a" + "foo bar")) {
                                                                                            v1 = ((v14 * (v14 | (v14 % ((v14) | 1)))) <= 42)
                                                                                            if v1 {
                                                                                                if (0.1 < 3.14159) {
                                                                                                    if v1 {
                                                                                                        if len(v6) > 0 {
                                                                                                            v6[0] = (fi33 & 127)
                                                                                                        }
                                                                                                        fmt.Println(fi33, fi27)
                                                                                                        switch fi33 {
                                                                                                        case -757:
                                                                                                            v40 = v40
                                                                                                        }
                                                                                                    } else {
                                                                                                        var v59 int8 = 127
                                                                                                        if false {
                                                                                                            if v1 {
                                                                                                                switch "hello" {
                                                                                                                case "a":
                                                                                                                    v58 = v26
                                                                                                                case "\n\t":
                                                                                                                    v4 = (^(v4))
                                                                                                                    switch ("\n\t" + "") {
                                                                                                                    case "hello":
                                                                                                                        if (!(1 > 4294967295)) {
                                                                                                                            if (0.1 != 3.14159) {
                                                                                                                                if (float32((-2.25 / ((((0.1 - (0.0 - 0.0)) * 0.0) / float64(v58)) + 0.1))) < (v26 - (((v26 / v58) + v26) * 3.14159))) {
                                                                                                                                    v58 = (float32((3.14159 * 3.14159)) / (-2.25 / v26))
                                                                                                                                } else {
                                                                                                                                    var v60 bool = (!v1)
                                                                                                                                    switch ("hello" + "hello") {
                                                                                                                                    case "a":
                                                                                                                                        v41 = 351
                                                                                                                                    }
                                                                                                                                    var v61 bool = ((v40 >> ((42) & 63)) < (v40 - (v40 / ((v40) | 1))))
                                                                                                                                    _ = len(v6)
                                                                                                                                    _ = v60
                                                                                                                                    _ = v61
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                if v1 {
                                                                                                                                    var v62 int8 = (-(v59))
                                                                                                                                    var v63 float64 = (1e10 + 1.5)
                                                                                                                                    v3 = (v3 % ((2) | 1))
                                                                                                                                    var v64 int16 = v14
                                                                                                                                    if len(v6) >= 2 {
                                                                                                                                        v6 = v6[1:len(v6)]
                                                                                                                                    }
                                                                                                                                    var v65 uint16 = 255
                                                                                                                                    _ = v62
                                                                                                                                    _ = v63
                                                                                                                                    _ = v64
                                                                                                                                    _ = v65
                                                                                                                                } else {
                                                                                                                                    fmt.Println(v26)
                                                                                                                                    v1 = (((v40 & (v40 * int32(fi33))) < (v40 * -410)) && (-2.25 > 3.14159))
                                                                                                                                    var v66 int = (fi33 % ((fi33) | 1))
                                                                                                                                    _ = v66
                                                                                                                                }
                                                                                                                                v59 -= v59
                                                                                                                                v4 -= (v3 + v4)
                                                                                                                            }
                                                                                                                            v67 := GenS1{F0: (v14 | (-(v14))), F1: (2147483647 != v40)}
                                                                                                                            v68 := GenS5{F0: v1}
                                                                                                                            _ = v67
                                                                                                                            _ = v68
                                                                                                                        } else {
                                                                                                                            _ = len(v5)
                                                                                                                            fmt.Println(v59, v1, v58)
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        if false {
                                                                                                                            v3 += (v3 - v3)
                                                                                                                            switch ("foo bar" + "") {
                                                                                                                            case "foo bar":
                                                                                                                                v26 = (v58 / 1.5)
                                                                                                                            case "":
                                                                                                                                v69 := GenS4{F0: v41, F1: v4}
                                                                                                                                _ = v69
                                                                                                                            }
                                                                                                                            if len(v2) > 0 {
                                                                                                                                v2[0] = 42
                                                                                                                            }
                                                                                                                            v70 := GenS6{F0: v14, F1: (false && (uint(v3) >= 0))}
                                                                                                                            fmt.Println(v14, fi33)
                                                                                                                            _ = v70
                                                                                                                        } else {
                                                                                                                            v2 = append(v2, 2147483648)
                                                                                                                            switch "\n\t" {
                                                                                                                            case "":
                                                                                                                                if (false || ((!((!(true && (0.1 >= 1e10))) && v1)) && (!((v59 >> ((324) & 63)) == v59)))) {
                                                                                                                                    v71 := GenS5{F0: (255 <= 255)}
                                                                                                                                    if (2 >= 1) {
                                                                                                                                        v72 := GenS7{F0: (v59 * 100), F1: 773, F2: -2147483648}
                                                                                                                                        v73 := GenS2{F0: 9223372036854775807, F1: (v41 >> ((2) & 63))}
                                                                                                                                        _ = len(v2)
                                                                                                                                        v5 = append(v5, 992)
                                                                                                                                        _ = v72
                                                                                                                                        _ = v73
                                                                                                                                    } else {
                                                                                                                                        v4 = v3
                                                                                                                                    }
                                                                                                                                    fmt.Println(v59, fi27, v41)
                                                                                                                                    _ = v71
                                                                                                                                }
                                                                                                                            }
                                                                                                                            switch int(v59) {
                                                                                                                            case 32767:
                                                                                                                                switch ("foo bar" + "\n\t") {
                                                                                                                                case "\n\t":
                                                                                                                                    switch (((("a" + "") + "a") + ("hello" + "\n\t")) + ("a" + ((("foo bar" + "") + "a") + (("\n\t" + "foo bar") + (("hello" + "") + ("hello" + "")))))) {
                                                                                                                                    case "\n\t":
                                                                                                                                        _ = len(v2)
                                                                                                                                    case "hello":
                                                                                                                                        v3 -= v3
                                                                                                                                    }
                                                                                                                                case "foo bar":
                                                                                                                                    v74 := GenS2{F0: 42, F1: 367}
                                                                                                                                    _ = v74
                                                                                                                                case "a":
                                                                                                                                    if v1 {
                                                                                                                                        v14 += -76
                                                                                                                                        var v75 int16 = v14
                                                                                                                                        _ = v75
                                                                                                                                    } else {
                                                                                                                                        if len(v5) >= 2 {
                                                                                                                                            v5 = v5[1:len(v5)]
                                                                                                                                        }
                                                                                                                                        if false {
                                                                                                                                            if false {
                                                                                                                                                v3 = (^(v3))
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                default:
                                                                                                                                    v76 := GenS4{F0: v3, F1: v4}
                                                                                                                                    _ = v76
                                                                                                                                }
                                                                                                                            case 127:
                                                                                                                                v77 := GenS7{F0: (^(v59)), F1: 18446744073709551615, F2: v40}
                                                                                                                                _ = v77
                                                                                                                            case 128:
                                                                                                                                _ = len(v5)
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                case "hello":
                                                                                                                    var v78 uint = 18446744073709551615
                                                                                                                    _ = v78
                                                                                                                default:
                                                                                                                    v79 := GenS6{F0: v14, F1: v1}
                                                                                                                    _ = v79
                                                                                                                }
                                                                                                                v26 -= v58
                                                                                                                v59 -= (v59 | (v59 | v59))
                                                                                                            } else {
                                                                                                                v59 += (-(v59))
                                                                                                            }
                                                                                                            fmt.Println(v40, v41)
                                                                                                        }
                                                                                                        fmt.Println(v40, v41, v40)
                                                                                                        _ = v59
                                                                                                    }
                                                                                                    switch (("foo bar" + "\n\t") + "\n\t") {
                                                                                                    case "hello":
                                                                                                        var v80 float32 = v26
                                                                                                        _ = v80
                                                                                                    default:
                                                                                                        var v81 uint32 = 0
                                                                                                        _ = v81
                                                                                                    }
                                                                                                    v82 := GenS7{F0: -128, F1: (v4 * 773), F2: v40}
                                                                                                    v6 = append(v6, (fi27 * fi27))
                                                                                                    var v83 float32 = v26
                                                                                                    _ = v82
                                                                                                    _ = v83
                                                                                                }
                                                                                            } else {
                                                                                                switch fi33 {
                                                                                                case 4294967295:
                                                                                                    fmt.Println(v26, v14)
                                                                                                case 128:
                                                                                                    v84 := GenS7{F0: -47, F1: 4294967295, F2: v40}
                                                                                                    _ = v84
                                                                                                case -137:
                                                                                                    if v1 {
                                                                                                        var v85 bool = ((-2.25 >= float32(((1.5 + 3.14159) + (float64(0.1) + ((-2.25 * 0.1) - float64(v26)))))) && ((-103 == -9223372036854775808) && v1))
                                                                                                        _ = v85
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            if ((!true) || (true && true)) {
                                                                                                v3 -= v3
                                                                                            } else {
                                                                                                if ((v14 & (v14 << ((917) & 63))) < v14) {
                                                                                                    var v86 uint32 = 65535
                                                                                                    if v1 {
                                                                                                        fmt.Println(v14, v40, v58)
                                                                                                    } else {
                                                                                                        var v87 uint64 = 4294967295
                                                                                                        _ = v87
                                                                                                    }
                                                                                                    fmt.Println(v40)
                                                                                                    _ = v86
                                                                                                } else {
                                                                                                    fmt.Println(v1)
                                                                                                }
                                                                                            }
                                                                                            if v1 {
                                                                                                v88 := GenS6{F0: v14, F1: (!v1)}
                                                                                                _ = v88
                                                                                            }
                                                                                            v26 = float32(0.0)
                                                                                        }
                                                                                        if (1.5 != ((0.1 / (-2.25 / -2.25)) + (1e10 / 1e10))) {
                                                                                            v26 -= (((((v58 * (v58 - v26)) + float32(1.5)) + (1e10 * v58)) - v58) + (v58 * 3.14159))
                                                                                            if (v1 || v1) {
                                                                                                switch int(v40) {
                                                                                                case 42:
                                                                                                    if len(v5) > 0 {
                                                                                                        v5[0] = 703
                                                                                                    }
                                                                                                case 529:
                                                                                                    v40 = v40
                                                                                                }
                                                                                                v89 := GenS8{F0: 9223372036854775807, F1: 42, F2: 9223372036854775807}
                                                                                                var v90 uint = 896
                                                                                                _ = v89
                                                                                                _ = v90
                                                                                            }
                                                                                            switch "\n\t" {
                                                                                            case "foo bar":
                                                                                                v26 -= 0.1
                                                                                            case "\n\t":
                                                                                                fmt.Println(v41, v4)
                                                                                            }
                                                                                        }
                                                                                        _ = v58
                                                                                    }
                                                                                    var v91 float32 = v26
                                                                                    _ = v91
                                                                                }
                                                                                var v92 int64 = -9223372036854775808
                                                                                v3 = v4
                                                                                _ = v92
                                                                            case "a":
                                                                                var v93 float32 = (v26 - v26)
                                                                                _ = v93
                                                                            default:
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                            }
                                                                        }
                                                                        var v94 uint32 = 2
                                                                        switch (^(fi27)) {
                                                                        case 2147483647:
                                                                            fmt.Println(fi27, v41, v94)
                                                                            switch -903 {
                                                                            case 4294967295:
                                                                                if true {
                                                                                    var v95 bool = v1
                                                                                    var v96 int32 = v40
                                                                                    _ = v95
                                                                                    _ = v96
                                                                                } else {
                                                                                    if (v94 < (v94 >> ((770) & 63))) {
                                                                                        v14 -= (^(v14))
                                                                                        switch "" {
                                                                                        case "\n\t":
                                                                                            var v97 float64 = -2.25
                                                                                            _ = v97
                                                                                        case "foo bar":
                                                                                            if (0 > 2) {
                                                                                                v98 := GenS5{F0: v1}
                                                                                                _ = v98
                                                                                            }
                                                                                        default:
                                                                                            var v99 uint8 = 255
                                                                                            _ = v99
                                                                                        }
                                                                                        v41 -= (v3 * (v3 / ((v3) | 1)))
                                                                                    }
                                                                                    switch "hello" {
                                                                                    case "foo bar":
                                                                                        switch (fi33 >> ((880) & 63)) {
                                                                                        case 1:
                                                                                            v100 := GenS1{F0: (v14 >> ((18446744073709551615) & 63)), F1: (fi27 <= fi27)}
                                                                                            _ = v100
                                                                                        case 9223372036854775807:
                                                                                            if v1 {
                                                                                                if true {
                                                                                                    var v101 int8 = -1
                                                                                                    var v102 uint8 = 0
                                                                                                    _ = v101
                                                                                                    _ = v102
                                                                                                }
                                                                                                fmt.Println(v26, v3)
                                                                                                v94 = v94
                                                                                            } else {
                                                                                                v5 = append(v5, 65535)
                                                                                                v5 = append(v5, v94)
                                                                                            }
                                                                                        default:
                                                                                            var v103 int16 = 1000
                                                                                            _ = v103
                                                                                        }
                                                                                    case "\n\t":
                                                                                        v104 := GenS6{F0: (v14 << ((500) & 63)), F1: (!v1)}
                                                                                        _ = v104
                                                                                    case "a":
                                                                                        var v105 uint32 = 0
                                                                                        _ = v105
                                                                                    default:
                                                                                        _ = len(v2)
                                                                                    }
                                                                                    var v106 string = "hello"
                                                                                    v107 := GenS2{F0: fi33, F1: (^(v4))}
                                                                                    _ = v106
                                                                                    _ = v107
                                                                                }
                                                                            case -685:
                                                                                v108 := GenS9{F0: fi27, F1: (("a" + ("foo bar" + "")) + "a"), F2: 100}
                                                                                _ = v108
                                                                            case -189:
                                                                                if false {
                                                                                    if ((!((v41 > v41) && (2 < 183))) || ((-94 < 42) && ((0 == 7) && true))) {
                                                                                        v40 -= int32(fi33)
                                                                                    } else {
                                                                                        v109 := GenS3{F0: genFn0(v14, 42), F1: 42}
                                                                                        if (!((!true) && ((true || v1) || ((false || ((^(v94)) >= (v94 / ((0) | 1)))) || v1)))) {
                                                                                            var v110 float64 = ((float64(3.14159) / 1e10) + float64(1.5))
                                                                                            _ = v110
                                                                                        }
                                                                                        v3 += (v4 >> ((18446744073709551615) & 63))
                                                                                        _ = v109
                                                                                    }
                                                                                    var v111 string = "\n\t"
                                                                                    var v112 int16 = v14
                                                                                    _ = v111
                                                                                    _ = v112
                                                                                }
                                                                            }
                                                                            v113 := GenS8{F0: 905, F1: 65535, F2: 9223372036854775807}
                                                                            _ = v113
                                                                        }
                                                                        _ = v94
                                                                    }
                                                                case "hello":
                                                                    switch fi27 {
                                                                    case -1:
                                                                        v3 = v41
                                                                    }
                                                                default:
                                                                    var v114 uint32 = 853
                                                                    _ = v114
                                                                }
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                                var v115 int8 = 0
                                                                _ = v115
                                                            }
                                                        } else {
                                                            switch 560 {
                                                            case 32767:
                                                                fmt.Println(v41)
                                                            case 42:
                                                                v116 := GenS9{F0: (fi27 << ((0) & 63)), F1: ("foo bar" + "foo bar"), F2: -1}
                                                                _ = v116
                                                            default:
                                                                var v117 bool = v1
                                                                _ = v117
                                                            }
                                                        }
                                                    }
                                                }
                                            case "\n\t":
                                                if len(v5) > 0 {
                                                    v5[0] = 1
                                                }
                                            }
                                            _ = v40
                                            _ = v41
                                        }
                                    }
                                    var v118 string = "\n\t"
                                    _ = v118
                                }
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                v6 = append(v6, fi27)
                                _ = fi33
                            }
                        }
                    }
                    _ = fi27
                }
                _ = v26
            }
        } else {
            for fi119 := 0; fi119 < 16; fi119++ {
                for fi120 := 0; fi120 < 20; fi120++ {
                    var v121 float32 = 1.5
                    if (127 != 127) {
                        var v122 float32 = float32((1.5 + 3.14159))
                        _ = len(v6)
                        var v123 float32 = v121
                        _ = v122
                        _ = v123
                    } else {
                        var v124 int8 = 100
                        switch ("" + "") {
                        case "":
                            var v125 bool = (v1 || (18446744073709551615 != 203))
                            _ = v125
                        case "foo bar":
                            fmt.Println(fi119, v3, v124)
                        case "\n\t":
                            switch ("" + "\n\t") {
                            case "\n\t":
                                if (!(!((v121 + (-2.25 / (v121 - (v121 + v121)))) != (3.14159 + v121)))) {
                                    v3 = (v3 & v4)
                                    _ = len(v2)
                                    var v126 int8 = v124
                                    _ = v126
                                }
                            case "foo bar":
                                v5 = append(v5, 0)
                            case "":
                                switch "foo bar" {
                                case "\n\t":
                                    v3 += 1
                                case "hello":
                                    _ = len(v6)
                                }
                            }
                        default:
                            var v127 uint64 = v3
                            _ = v127
                        }
                        var v128 int = fi119
                        _ = v124
                        _ = v128
                    }
                    _ = fi120
                    _ = v121
                }
                v3 = v4
                var v129 int16 = (-(v14))
                _ = fi119
                _ = v129
            }
            v130 := GenS4{F0: 189, F1: 2}
            _ = v130
        }
        fmt.Println(v3, v3)
    } else {
        for fi131 := 0; fi131 < 2; fi131++ {
            _ = len(v2)
            if (true || true) {
                var v132 uint8 = 1
                _ = v132
            }
            _ = fi131
        }
        if (v1 || (("\n\t" >= "\n\t") && ((!v1) || v1))) {
            if (42 < 587) {
                var v133 uint = 0
                if (!(v3 == v3)) {
                    if v1 {
                        v134 := GenS6{F0: (v14 - 1000), F1: (325 <= 42)}
                        _ = v134
                    } else {
                        var v135 int = -274
                        _ = v135
                    }
                    v4 = v4
                    for fi136 := 0; fi136 < 20; fi136++ {
                        v5 = append(v5, 65535)
                        switch fi136 {
                        case 128:
                            v3 -= v4
                        case 1:
                            var v137 bool = (false || v1)
                            _ = v137
                        case -2147483648:
                            v138 := GenS1{F0: v14, F1: ((v4 / (((v3 >> ((18446744073709551615) & 63))) | 1)) < (v3 - 2))}
                            _ = v138
                        }
                        if true {
                            var v139 bool = ((-(fi136)) >= fi136)
                            _ = v139
                        } else {
                            for fi140 := 0; fi140 < 15; fi140++ {
                                var v141 int8 = 100
                                var v142 uint32 = 42
                                _ = fi140
                                _ = v141
                                _ = v142
                            }
                            v14 -= (v14 + (v14 >> ((v133) & 63)))
                        }
                        _ = fi136
                    }
                } else {
                    v4 = (v4 | v3)
                    if (0 < 42) {
                        v3 = v4
                        switch 130 {
                        case 437:
                            var v143 uint32 = 948
                            _ = v143
                        case 42:
                            if v1 {
                                v144 := GenS1{F0: -855, F1: ((!true) && (-128 <= int8(v3)))}
                                _ = v144
                            }
                        default:
                            switch ("\n\t" + "a") {
                            case "":
                                for fi145 := 0; fi145 < 6; fi145++ {
                                    v146 := GenS10{F0: 2}
                                    v147 := GenS11{F0: 42, F1: 927, F2: ""}
                                    var v148 int32 = -436
                                    _ = fi145
                                    _ = v146
                                    _ = v147
                                    _ = v148
                                }
                            default:
                                v149 := GenS12{F0: v133, F1: 2147483647, F2: 0}
                                _ = v149
                            }
                        }
                        if (199 < 765) {
                            switch (("" + "a") + (("" + "\n\t") + "")) {
                            case "a":
                                var v150 string = ""
                                _ = v150
                            case "\n\t":
                                if ("foo bar" >= ("" + ("" + "foo bar"))) {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    if (v3 > (^(v4))) {
                                        var v151 int64 = 9223372036854775807
                                        if (v1 || (151 < 0)) {
                                            for fi152 := 0; fi152 < 4; fi152++ {
                                                var v153 uint = (v133 & v133)
                                                if v1 {
                                                    for fi154 := 0; fi154 < 13; fi154++ {
                                                        switch (fi154 * 9223372036854775807) {
                                                        case 42:
                                                            v6 = append(v6, 32768)
                                                        case -154:
                                                            if false {
                                                                if false {
                                                                    v153 = (^(v153))
                                                                    if len(v2) > 0 {
                                                                        v2[0] = 2
                                                                    }
                                                                    v6 = append(v6, fi152)
                                                                }
                                                                if (0 == int32(fi154)) {
                                                                    var v155 uint32 = 2147483648
                                                                    v156 := GenS9{F0: 128, F1: ("" + "\n\t"), F2: 42}
                                                                    if v1 {
                                                                        var v157 uint16 = 32767
                                                                        _ = v157
                                                                    } else {
                                                                        var v158 int32 = -842
                                                                        fmt.Println(v1)
                                                                        var v159 float64 = (1e10 / ((0.1 / 1e10) - float64(1.5)))
                                                                        _ = v158
                                                                        _ = v159
                                                                    }
                                                                    var v160 uint = (v133 + (v133 ^ uint(v155)))
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                    _ = v155
                                                                    _ = v156
                                                                    _ = v160
                                                                }
                                                            }
                                                        default:
                                                            if len(v2) > 0 {
                                                                v2[0] = 2147483647
                                                            }
                                                        }
                                                        _ = fi154
                                                    }
                                                    var v161 float64 = 0.0
                                                    for fi162 := 0; fi162 < 11; fi162++ {
                                                        v133 = (v133 - 272)
                                                        _ = fi162
                                                    }
                                                    _ = v161
                                                } else {
                                                    for fi163 := 0; fi163 < 10; fi163++ {
                                                        var v164 float64 = -2.25
                                                        v165 := GenS4{F0: v4, F1: v3}
                                                        _ = fi163
                                                        _ = v164
                                                        _ = v165
                                                    }
                                                }
                                                var v166 int64 = v151
                                                _ = fi152
                                                _ = v153
                                                _ = v166
                                            }
                                            if (!v1) {
                                                for fi167 := 0; fi167 < 16; fi167++ {
                                                    if (-673 == (v151 - (v151 << (((v133 & (v133 * (v133 / ((v133) | 1))))) & 63)))) {
                                                        switch ("a" + "\n\t") {
                                                        case "a":
                                                            v5 = append(v5, 4294967295)
                                                        case "":
                                                            if len(v2) > 0 {
                                                                v2[0] = 2147483648
                                                            }
                                                        case "\n\t":
                                                            fmt.Println(v3)
                                                        default:
                                                            for fi168 := 0; fi168 < 1; fi168++ {
                                                                if (!v1) {
                                                                    var v169 uint32 = 42
                                                                    fmt.Println(v14)
                                                                    _ = v169
                                                                }
                                                                _ = len(v2)
                                                                _ = fi168
                                                            }
                                                        }
                                                    }
                                                    var v170 int = (fi167 ^ (-(fi167)))
                                                    var v171 int32 = 2147483647
                                                    _ = fi167
                                                    _ = v170
                                                    _ = v171
                                                }
                                                var v172 int = 0
                                                _ = v172
                                            }
                                            if len(v2) > 0 {
                                                v2[0] = 42
                                            }
                                        } else {
                                            if v1 {
                                                v14 = v14
                                                v3 = v3
                                                v14 -= v14
                                            } else {
                                                v133 = (^(v133))
                                            }
                                            switch ("" + "\n\t") {
                                            case "foo bar":
                                                v133 -= v133
                                            case "a":
                                                v173 := GenS5{F0: (("a" != ("hello" + "foo bar")) || (v1 && v1))}
                                                _ = v173
                                            case "\n\t":
                                                if (18446744073709551615 < 195) {
                                                    fmt.Println(v133)
                                                    for fi174 := 0; fi174 < 8; fi174++ {
                                                        if (!(!(!v1))) {
                                                            var v175 int8 = int8(v14)
                                                            _ = v175
                                                        } else {
                                                            for fi176 := 0; fi176 < 3; fi176++ {
                                                                var v177 uint8 = 38
                                                                _ = fi176
                                                                _ = v177
                                                            }
                                                        }
                                                        if (v14 >= (v14 / ((v14) | 1))) {
                                                            v3 = 42
                                                        }
                                                        _ = fi174
                                                    }
                                                    var v178 uint32 = 65535
                                                    _ = v178
                                                }
                                            }
                                            if ((65535 > 2) && v1) {
                                                v179 := GenS9{F0: -104, F1: "hello", F2: 42}
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                var v180 uint64 = v4
                                                _ = v179
                                                _ = v180
                                            } else {
                                                v1 = (v1 || true)
                                                if v1 {
                                                    v181 := GenS13{F0: 255}
                                                    v182 := GenS11{F0: (v151 << ((524) & 63)), F1: 605, F2: "\n\t"}
                                                    _ = v181
                                                    _ = v182
                                                } else {
                                                    var v183 bool = (v4 <= v3)
                                                    if false {
                                                        v14 -= (v14 - int16(v4))
                                                        _ = len(v2)
                                                    }
                                                    fmt.Println(v183, v1)
                                                    _ = v183
                                                }
                                                if v1 {
                                                    if v1 {
                                                        var v184 string = "hello"
                                                        for fi185 := 0; fi185 < 11; fi185++ {
                                                            var v186 uint8 = 1
                                                            _ = fi185
                                                            _ = v186
                                                        }
                                                        v151 = (^(v151))
                                                        _ = v184
                                                    }
                                                    v187 := GenS7{F0: -62, F1: (v3 << ((v133) & 63)), F2: 0}
                                                    for fi188 := 0; fi188 < 20; fi188++ {
                                                        v14 = v14
                                                        _ = fi188
                                                    }
                                                    _ = v187
                                                }
                                            }
                                        }
                                        if ((0.0 / 1.0) >= ((0.0 / 1.5) + 1e10)) {
                                            var v189 string = "\n\t"
                                            for fi190 := 0; fi190 < 9; fi190++ {
                                                var v191 int8 = 100
                                                switch -109 {
                                                case 128:
                                                    var v192 float64 = -2.25
                                                    _ = v192
                                                case -9223372036854775808:
                                                    fmt.Println(v133, v191, v189)
                                                case 4294967295:
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                default:
                                                    if (v14 <= (v14 / ((v14) | 1))) {
                                                        v193 := GenS10{F0: 65535}
                                                        switch v189 {
                                                        case "foo bar":
                                                            v14 = (-(v14))
                                                        }
                                                        _ = v193
                                                    }
                                                }
                                                _ = fi190
                                                _ = v191
                                            }
                                            var v194 int8 = 127
                                            _ = v189
                                            _ = v194
                                        } else {
                                            v151 -= (v151 % ((v151) | 1))
                                            v151 += v151
                                        }
                                        _ = v151
                                    } else {
                                        var v195 int16 = 85
                                        if len(v2) > 0 {
                                            v2[0] = 2
                                        }
                                        v3 -= 474
                                        _ = v195
                                    }
                                    v196 := GenS11{F0: 0, F1: -931, F2: ("a" + "foo bar")}
                                    _ = v196
                                }
                            }
                            v133 -= (v133 | 4294967295)
                        } else {
                            var v197 uint64 = 1
                            if false {
                                if true {
                                    v14 = (v14 | 1000)
                                    fmt.Println(v4)
                                    v197 -= (^(v3))
                                }
                                if v1 {
                                    switch ("" + ("a" + ("hello" + ("" + "a")))) {
                                    case "foo bar":
                                        _ = len(v6)
                                    case "\n\t":
                                        if (float32((-2.25 + 0.0)) == ((-2.25 - 0.0) - -2.25)) {
                                            v1 = v1
                                            var v198 int = 56
                                            _ = v198
                                        }
                                    default:
                                        if v1 {
                                            if (!(1e10 != (0.1 * (0.1 / 3.14159)))) {
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                var v199 int8 = -6
                                                var v200 int16 = (v14 << ((2) & 63))
                                                _ = v199
                                                _ = v200
                                            } else {
                                                for fi201 := 0; fi201 < 15; fi201++ {
                                                    fmt.Println(v4)
                                                    var v202 int32 = -909
                                                    if v1 {
                                                        if v1 {
                                                            fmt.Println(v133, fi201)
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        } else {
                                                            for fi203 := 0; fi203 < 1; fi203++ {
                                                                v197 -= v4
                                                                _ = fi203
                                                            }
                                                            var v204 string = "hello"
                                                            var v205 uint = (v133 / ((v133) | 1))
                                                            _ = v204
                                                            _ = v205
                                                        }
                                                        v202 += v202
                                                    }
                                                    _ = fi201
                                                    _ = v202
                                                }
                                                v206 := GenS4{F0: (v3 * (v197 % (((v4 - v197)) | 1))), F1: (^(v4))}
                                                _ = v206
                                            }
                                            for fi207 := 0; fi207 < 16; fi207++ {
                                                for fi208 := 0; fi208 < 1; fi208++ {
                                                    if (!(true || (-1 < 889))) {
                                                        var v209 uint16 = 255
                                                        _ = v209
                                                    }
                                                    var v210 int64 = 9223372036854775807
                                                    var v211 int = fi207
                                                    var v212 int16 = -660
                                                    if false {
                                                        if v1 {
                                                            switch (v211 >> (((v133 % ((65535) | 1))) & 63)) {
                                                            case 32768:
                                                                v4 = v3
                                                                if false {
                                                                    var v213 uint32 = 65535
                                                                    _ = v213
                                                                }
                                                            case 32767:
                                                                v6 = append(v6, fi208)
                                                            default:
                                                                switch (fi208 + fi207) {
                                                                case 9223372036854775807:
                                                                    if len(v6) > 0 {
                                                                        v6[0] = fi207
                                                                    }
                                                                case -1:
                                                                    var v214 string = "foo bar"
                                                                    _ = v214
                                                                default:
                                                                    var v215 int16 = v212
                                                                    _ = v215
                                                                }
                                                            }
                                                            if v1 {
                                                                if (v1 && (true && v1)) {
                                                                    v6 = append(v6, 32768)
                                                                    var v216 int8 = 127
                                                                    fmt.Println(v211, v212)
                                                                    var v217 float64 = 0.1
                                                                    _ = v216
                                                                    _ = v217
                                                                }
                                                            } else {
                                                                var v218 uint = (v133 % (((v133 | v133)) | 1))
                                                                v219 := GenS14{F0: false, F1: v1}
                                                                switch (-(v211)) {
                                                                case 128:
                                                                    var v220 string = "a"
                                                                    _ = v220
                                                                case -517:
                                                                    switch (("a" + "foo bar") + "a") {
                                                                    case "hello":
                                                                        var v221 float32 = -2.25
                                                                        _ = v221
                                                                    case "\n\t":
                                                                        fmt.Println(v218, v4)
                                                                    case "foo bar":
                                                                        v222 := GenS11{F0: v210, F1: 2147483647, F2: "\n\t"}
                                                                        if v1 {
                                                                            _ = len(v2)
                                                                        } else {
                                                                            v14 = v212
                                                                        }
                                                                        fmt.Println(v210)
                                                                        _ = v222
                                                                    }
                                                                default:
                                                                    switch "" {
                                                                    case "":
                                                                        if true {
                                                                            if ("" < "foo bar") {
                                                                                v211 = (fi207 ^ (fi208 >> ((1) & 63)))
                                                                                var v223 int8 = 127
                                                                                var v224 uint = v218
                                                                                _ = v223
                                                                                _ = v224
                                                                            } else {
                                                                                if (v1 && false) {
                                                                                    if (v1 || v1) {
                                                                                        v211 = fi208
                                                                                        if true {
                                                                                            var v225 int = fi207
                                                                                            _ = len(v2)
                                                                                            _ = v225
                                                                                        }
                                                                                        v211 -= (v211 - v211)
                                                                                    } else {
                                                                                        v3 = v197
                                                                                        fmt.Println(v14, v211)
                                                                                        if false {
                                                                                            var v226 uint32 = 117
                                                                                            var v227 uint32 = v226
                                                                                            _ = v226
                                                                                            _ = v227
                                                                                        } else {
                                                                                            if len(v6) > 0 {
                                                                                                v6[0] = -9223372036854775808
                                                                                            }
                                                                                            if len(v6) >= 2 {
                                                                                                v6 = v6[1:len(v6)]
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    if (2147483647 >= -2147483648) {
                                                                                        v218 = (v133 + 1)
                                                                                    }
                                                                                    var v228 float64 = -2.25
                                                                                    if (v218 != v218) {
                                                                                        var v229 int64 = (v210 - v210)
                                                                                        if (false && false) {
                                                                                            if v1 {
                                                                                                var v230 int = (fi207 << ((v218) & 63))
                                                                                                _ = v230
                                                                                            }
                                                                                        } else {
                                                                                            v210 -= (-(v229))
                                                                                            var v231 uint32 = 458
                                                                                            _ = v231
                                                                                        }
                                                                                        _ = v229
                                                                                    } else {
                                                                                        v211 = int(v14)
                                                                                    }
                                                                                    fmt.Println(v210, v1, v4)
                                                                                    _ = v228
                                                                                }
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = 47
                                                                                }
                                                                            }
                                                                            v14 = (v14 & (v14 - (v14 / (((v212 - v212)) | 1))))
                                                                        }
                                                                    case "\n\t":
                                                                        var v232 int32 = 2147483647
                                                                        _ = v232
                                                                    case "a":
                                                                        v1 = v1
                                                                    }
                                                                }
                                                                _ = v218
                                                                _ = v219
                                                            }
                                                        } else {
                                                            fmt.Println(v3, v14)
                                                        }
                                                    } else {
                                                        var v233 uint64 = (v3 / ((1) | 1))
                                                        if (4294967295 == v3) {
                                                            v234 := GenS15{F0: 2147483647}
                                                            _ = v234
                                                        }
                                                        var v235 int8 = 42
                                                        _ = v233
                                                        _ = v235
                                                    }
                                                    _ = fi208
                                                    _ = v210
                                                    _ = v211
                                                    _ = v212
                                                }
                                                var v236 float32 = (0.0 / 1.5)
                                                _ = fi207
                                                _ = v236
                                            }
                                        } else {
                                            for fi237 := 0; fi237 < 9; fi237++ {
                                                for fi238 := 0; fi238 < 12; fi238++ {
                                                    _ = len(v2)
                                                    fmt.Println(fi237, v197)
                                                    v2 = append(v2, 0)
                                                    v3 = v3
                                                    fmt.Println(fi237, v133)
                                                    v6 = append(v6, fi237)
                                                    _ = fi238
                                                }
                                                _ = fi237
                                            }
                                            v4 -= 42
                                        }
                                    }
                                    fmt.Println(v197)
                                    v239 := GenS8{F0: 0, F1: v133, F2: -550}
                                    _ = v239
                                } else {
                                    v240 := GenS16{F0: 2, F1: v3}
                                    switch "" {
                                    case "\n\t":
                                        fmt.Println(v197)
                                    }
                                    var v241 uint16 = 338
                                    _ = v240
                                    _ = v241
                                }
                            }
                            _ = v197
                        }
                    } else {
                        switch "" {
                        case "foo bar":
                            var v242 int64 = -815
                            _ = v242
                        default:
                            var v243 uint8 = 128
                            _ = v243
                        }
                        if v1 {
                            var v244 uint8 = 157
                            _ = v244
                        }
                        var v245 uint = v133
                        _ = v245
                    }
                }
                v246 := GenS13{F0: v133}
                _ = v133
                _ = v246
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            if len(v5) > 0 {
                v5[0] = 156
            }
            var v247 float64 = (3.14159 / (-2.25 - 3.14159))
            _ = v247
        }
    }
    v14 -= v14
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v248 int8 = 0
    var v249 float64 = ((1.5 * 1.5) / 1.0)
    var v250 int64 = 9223372036854775807
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi251 := 0; fi251 < 18; fi251++ {
        _ = len(v5)
        _ = fi251
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v248:", v248)
    fmt.Println("v249:", v249)
    fmt.Println("v250:", v250)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v248
    _ = v249
    _ = v250
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
