package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint8; F2 bool }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 bool; F1 string; F2 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint64; F1 int64; F2 uint }
type GenS7 struct { F0 uint32; F1 uint8; F2 string }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint64; F1 int32 }
type GenS10 struct { F0 int; F1 bool; F2 uint64 }
type GenS11 struct { F0 uint; F1 int }
type GenS12 struct { F0 uint8; F1 uint64 }
type GenS13 struct { F0 int16; F1 string }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 uint }
type GenS17 struct { F0 uint }
type GenS18 struct { F0 uint; F1 uint16 }
func genFn0(p0 int32) int {
    v6 := GenS1{F0: 42, F1: 65535, F2: 21}
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 int32 = -2147483648
            v10 := []int{(fi8 * 433), fi8, fi8, (-(fi7))}
            var v11 int16 = -32768
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        p0 += p0
        p0 += (p0 * p0)
        _ = fi7
    }
    var v12 uint8 = 52
    p0 = p0
    fmt.Println(v12, v12)
    _ = p0
    _ = v12
    _ = v6
    return 72
}

func main() {
    var v1 float32 = float32(0.1)
    var v2 bool = (!(1e10 > (1e10 + 0.0)))
    var v3 float64 = 3.14159
    v4 := []string{("hello" + (("\n\t" + "") + ("hello" + "foo bar"))), "\n\t", "hello"}
    var v5 uint8 = 157
    var v13 bool = v2
    for fi14 := 0; fi14 < 7; fi14++ {
        switch "hello" {
        case "foo bar":
            v15 := GenS1{F0: (v5 & 1), F1: 2147483648, F2: (v5 & (v5 % ((v5) | 1)))}
            _ = v15
        case "a":
            v1 = (v1 - (v1 + (v1 - ((v1 / v1) - 3.14159))))
        case "\n\t":
            for fi16 := 0; fi16 < 2; fi16++ {
                fmt.Println(v3)
                if (v3 < v3) {
                    switch (fi14 >> ((628) & 63)) {
                    case 4294967295:
                        var v17 uint64 = 964
                        _ = v17
                    case 2147483647:
                        switch (fi16 - (fi16 - (fi16 + 127))) {
                        case 42:
                            v3 = (1.5 / v3)
                        case -9223372036854775808:
                            if (65535 < 1) {
                                switch ("a" + "foo bar") {
                                case "":
                                    v3 = ((v3 / (v3 - v3)) - v3)
                                default:
                                    if (v5 >= (^(v5))) {
                                        v13 = v2
                                        fmt.Println(fi14)
                                        fmt.Println(v2, v1, v5)
                                    }
                                }
                                var v18 uint32 = 466
                                v18 = v18
                                _ = v18
                            } else {
                                v1 = -2.25
                                if len(v4) > 0 {
                                    v4[0] = "\n\t"
                                }
                                v3 = (-2.25 / (0.0 + ((-2.25 / 1.0) * 1.5)))
                            }
                        }
                    case 637:
                        v13 = (33 != 1)
                    }
                } else {
                    var v19 uint32 = 0
                    fmt.Println(v19, v13, v19)
                    _ = v19
                }
                _ = fi16
            }
        default:
            for fi20 := 0; fi20 < 4; fi20++ {
                if v13 {
                    v2 = true
                    if true {
                        var v21 uint16 = 2
                        _ = v21
                    } else {
                        var v22 uint32 = 444
                        var v23 uint = 18446744073709551615
                        _ = v22
                        _ = v23
                    }
                    v4 = append(v4, ("" + "foo bar"))
                }
                switch (fi20 - genFn0(2147483647)) {
                case 9223372036854775807:
                    v13 = (v13 && (210 > -1))
                    v24 := GenS2{F0: 523, F1: v5, F2: (2 >= 18446744073709551615)}
                    var v25 uint = 203
                    _ = v24
                    _ = v25
                case -378:
                    var v26 uint16 = 255
                    v27 := GenS2{F0: v26, F1: v5, F2: v13}
                    _ = v26
                    _ = v27
                case -253:
                    v3 = v3
                }
                _ = fi20
            }
        }
        var v28 uint64 = 370
        for fi29 := 0; fi29 < 14; fi29++ {
            fmt.Println(v1)
            var v30 int64 = 1
            _ = fi29
            _ = v30
        }
        _ = fi14
        _ = v28
    }
    for fi31 := 0; fi31 < 15; fi31++ {
        v3 = (0.1 * v3)
        switch "" {
        case "a":
            var v32 float32 = (float32(v3) * float32(1.5))
            _ = v32
        case "":
            switch fi31 {
            case 4294967295:
                if (674 != 54) {
                    for fi33 := 0; fi33 < 19; fi33++ {
                        v3 = 1e10
                        var v34 float64 = v3
                        _ = fi33
                        _ = v34
                    }
                    if false {
                        if (true && ((!(42 == 0)) && ((v1 + (0.0 * 1e10)) >= (float32(-2.25) - float32((((float64(v1) + v3) - (float64((float32(float64(v1)) + v1)) + float64((v1 / 3.14159)))) * (3.14159 - v3))))))) {
                            for fi35 := 0; fi35 < 4; fi35++ {
                                v36 := GenS2{F0: 65535, F1: 128, F2: ((65535 != 42) && true)}
                                fmt.Println(v3, v5)
                                _ = fi35
                                _ = v36
                            }
                            if v2 {
                                var v37 uint32 = 454
                                _ = v37
                            }
                        }
                        switch (fi31 * 2147483647) {
                        case 624:
                            for fi38 := 0; fi38 < 7; fi38++ {
                                if len(v4) > 0 {
                                    v4[0] = ("hello" + ("\n\t" + ((("foo bar" + "foo bar") + ("" + "\n\t")) + ("" + ""))))
                                }
                                if len(v4) > 0 {
                                    v4[0] = ""
                                }
                                _ = fi38
                            }
                        default:
                            v1 = v1
                        }
                        var v39 float64 = float64(1e10)
                        _ = v39
                    }
                    switch genFn0(-867) {
                    case -1:
                        switch fi31 {
                        case -2147483648:
                            var v40 int64 = 998
                            _ = v40
                        }
                    case 32768:
                        var v41 int64 = 9223372036854775807
                        _ = v41
                    }
                } else {
                    for fi42 := 0; fi42 < 9; fi42++ {
                        if (v5 < v5) {
                            if v2 {
                                if (42 != 118) {
                                    switch "foo bar" {
                                    case "\n\t":
                                        v3 -= (v3 * 1.5)
                                    default:
                                        v5 = v5
                                    }
                                } else {
                                    var v43 float64 = ((((v3 * v3) / v3) + (v3 - 3.14159)) * (-2.25 + v3))
                                    var v44 uint64 = 3
                                    switch "foo bar" {
                                    case "":
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    default:
                                        switch (fi31 * genFn0(-1)) {
                                        case 9223372036854775807:
                                            if v13 {
                                                v4 = append(v4, ((("hello" + ("" + "")) + "") + ((("a" + ("" + ("a" + ""))) + "foo bar") + "a")))
                                                v44 = (v44 | v44)
                                                switch "\n\t" {
                                                case "foo bar":
                                                    if (v13 && (!(1 > 518))) {
                                                        switch genFn0(-1) {
                                                        case -496:
                                                            _ = len(v4)
                                                        case -29:
                                                            fmt.Println(fi31, fi42)
                                                        }
                                                        switch "a" {
                                                        case "a":
                                                            v45 := GenS2{F0: 42, F1: (v5 * 128), F2: v13}
                                                            _ = v45
                                                        case "hello":
                                                            switch fi42 {
                                                            case -373:
                                                                if len(v4) > 0 {
                                                                    v4[0] = ("" + "foo bar")
                                                                }
                                                            default:
                                                                var v46 uint64 = (^(v44))
                                                                _ = v46
                                                            }
                                                        case "":
                                                            v44 = v44
                                                        default:
                                                            var v47 string = "a"
                                                            _ = v47
                                                        }
                                                    } else {
                                                        v48 := GenS2{F0: 1, F1: (v5 & (v5 ^ 1)), F2: (v13 && false)}
                                                        if true {
                                                            v5 = 30
                                                            v49 := GenS3{F0: -128, F1: 4294967295}
                                                            _ = v49
                                                        }
                                                        _ = v48
                                                    }
                                                case "hello":
                                                    if false {
                                                        fmt.Println(v43, v2)
                                                    }
                                                case "a":
                                                    fmt.Println(v3, fi31)
                                                }
                                                v44 = uint64(v5)
                                                v44 -= (v44 | (v44 | v44))
                                                switch -2147483648 {
                                                case 4294967295:
                                                    if ((v5 % ((uint8(v44)) | 1)) >= v5) {
                                                        if false {
                                                            _ = len(v4)
                                                            var v50 float64 = (v3 - 1.5)
                                                            var v51 int = -381
                                                            _ = v50
                                                            _ = v51
                                                        } else {
                                                            fmt.Println(v2, v5, v13)
                                                            v5 -= (v5 | v5)
                                                        }
                                                        if false {
                                                            var v52 float64 = ((0.0 + -2.25) / (0.1 + (((v43 + 0.0) / (float64(((v1 + (0.1 + v1)) * v1)) + (v3 / ((v3 / v43) - 1e10)))) / ((((1e10 + -2.25) / v43) + (v3 * v43)) + v43))))
                                                            if v2 {
                                                                fmt.Println(v52, v1, v3)
                                                            }
                                                            _ = v52
                                                        } else {
                                                            var v53 int8 = 107
                                                            _ = v53
                                                        }
                                                        var v54 uint16 = uint16(v44)
                                                        _ = v54
                                                    } else {
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            fmt.Println(fi42, v2)
                                                        default:
                                                            var v55 int16 = 286
                                                            _ = v55
                                                        }
                                                        v1 = (float32(((((-2.25 / (float64((v1 / 1.0)) - v43)) * (float64(v1) / v3)) * v3) - v3)) - 0.1)
                                                    }
                                                case 414:
                                                    var v56 int32 = 2147483647
                                                    _ = v56
                                                case -9223372036854775808:
                                                    v44 -= (^(v44))
                                                default:
                                                    var v57 int16 = -474
                                                    _ = v57
                                                }
                                            } else {
                                                switch "" {
                                                case "":
                                                    if (!false) {
                                                        v58 := GenS4{F0: v2, F1: "foo bar", F2: 210}
                                                        v4 = append(v4, "\n\t")
                                                        var v59 uint32 = 2
                                                        fmt.Println(v3, v5, v59)
                                                        _ = v58
                                                        _ = v59
                                                    } else {
                                                        v60 := GenS2{F0: 32767, F1: (v5 * (v5 ^ v5)), F2: true}
                                                        var v61 int64 = -9223372036854775808
                                                        if len(v4) > 0 {
                                                            v4[0] = "\n\t"
                                                        }
                                                        _ = v60
                                                        _ = v61
                                                    }
                                                case "foo bar":
                                                    v13 = (!(false || v13))
                                                default:
                                                    fmt.Println(v13, fi31, v43)
                                                }
                                                _ = len(v4)
                                                v4 = append(v4, ("" + "\n\t"))
                                            }
                                        case 2147483647:
                                            _ = len(v4)
                                        default:
                                            v1 = v1
                                        }
                                    }
                                    _ = v43
                                    _ = v44
                                }
                            }
                            _ = len(v4)
                            if (fi31 <= (fi42 ^ (fi42 ^ genFn0(-2147483648)))) {
                                v13 = v2
                                if v2 {
                                    v2 = v13
                                    v62 := GenS1{F0: 128, F1: 962, F2: 1}
                                    switch 159 {
                                    case -1:
                                        fmt.Println(v13)
                                    case -510:
                                        if v2 {
                                            if len(v4) > 0 {
                                                v4[0] = "\n\t"
                                            }
                                            var v63 float32 = ((3.14159 / 1.5) / v1)
                                            var v64 uint8 = (v5 >> ((23) & 63))
                                            _ = len(v4)
                                            v65 := GenS4{F0: (0 <= 42), F1: "hello", F2: v5}
                                            _ = v63
                                            _ = v64
                                            _ = v65
                                        } else {
                                            switch fi31 {
                                            case 2147483647:
                                                switch (^(fi42)) {
                                                case 69:
                                                    v66 := GenS5{F0: 187}
                                                    _ = v66
                                                case 1:
                                                    fmt.Println(v1, v1, v1)
                                                case 518:
                                                    v67 := GenS6{F0: 2, F1: 0, F2: 255}
                                                    _ = v67
                                                default:
                                                    if (0 > 316) {
                                                        switch "a" {
                                                        case "hello":
                                                            if (v3 == v3) {
                                                                if v13 {
                                                                    v68 := GenS7{F0: 847, F1: (^(v5)), F2: ""}
                                                                    var v69 uint16 = 0
                                                                    _ = v68
                                                                    _ = v69
                                                                }
                                                            } else {
                                                                v5 -= v5
                                                            }
                                                        case "":
                                                            v1 = (-2.25 - v1)
                                                        case "foo bar":
                                                            var v70 int64 = 9223372036854775807
                                                            _ = v70
                                                        default:
                                                            switch ("a" + "hello") {
                                                            case "":
                                                                if (306 <= 0) {
                                                                    if (fi31 != (fi31 ^ (fi42 - fi31))) {
                                                                        v13 = (!v13)
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                    } else {
                                                                        v71 := GenS5{F0: v5}
                                                                        _ = v71
                                                                    }
                                                                    if ((1.5 < v3) || (0 > 4294967295)) {
                                                                        v1 += v1
                                                                        var v72 int16 = 32767
                                                                        v73 := GenS8{F0: "a"}
                                                                        _ = v72
                                                                        _ = v73
                                                                    }
                                                                    v74 := GenS9{F0: 1, F1: -1}
                                                                    _ = v74
                                                                } else {
                                                                    v2 = v13
                                                                    v1 -= 1e10
                                                                    var v75 uint = 510
                                                                    _ = v75
                                                                }
                                                            case "\n\t":
                                                                var v76 uint64 = 1
                                                                _ = v76
                                                            case "hello":
                                                                v3 = 1e10
                                                            }
                                                        }
                                                        if (415 != 199) {
                                                            fmt.Println(fi42, v3, fi31)
                                                            var v77 float64 = (((v3 / v3) / (((0.1 + (v3 + (v3 - v3))) / float64(3.14159)) + ((v3 - float64((v1 / v1))) + (v3 * 0.1)))) * 1.5)
                                                            fmt.Println(v5, fi31)
                                                            _ = v77
                                                        }
                                                    } else {
                                                        var v78 int8 = -128
                                                        v5 -= (v5 + (v5 * (^(v5))))
                                                        var v79 float32 = 0.1
                                                        _ = v78
                                                        _ = v79
                                                    }
                                                }
                                            case -358:
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            case 774:
                                                if true {
                                                    var v80 uint16 = 844
                                                    _ = v80
                                                } else {
                                                    v1 = (v1 * v1)
                                                }
                                            default:
                                                fmt.Println(v5, fi31)
                                            }
                                            if v13 {
                                                switch genFn0(0) {
                                                case 867:
                                                    if (genFn0(1) != (^(fi31))) {
                                                        var v81 float64 = float64(v1)
                                                        _ = v81
                                                    } else {
                                                        fmt.Println(v3, v1, fi42)
                                                        var v82 uint32 = 1
                                                        v83 := GenS6{F0: 1, F1: 500, F2: 1}
                                                        _ = v82
                                                        _ = v83
                                                    }
                                                }
                                                v84 := GenS4{F0: true, F1: "", F2: v5}
                                                _ = v84
                                            } else {
                                                v85 := GenS4{F0: v13, F1: ("hello" + ("" + "hello")), F2: 42}
                                                _ = v85
                                            }
                                            if (-1 >= -1) {
                                                v13 = v13
                                                v86 := GenS6{F0: 390, F1: 0, F2: 894}
                                                fmt.Println(v2, v5)
                                                _ = v86
                                            }
                                        }
                                    case 2147483647:
                                        v5 -= 42
                                        v87 := GenS5{F0: v5}
                                        var v88 bool = (!true)
                                        v3 -= 3.14159
                                        _ = v87
                                        _ = v88
                                    default:
                                        fmt.Println(v13, v2, v5)
                                        switch "a" {
                                        case "":
                                            v13 = false
                                            if (0 == 255) {
                                                if v13 {
                                                    switch fi31 {
                                                    case 42:
                                                        v89 := GenS10{F0: (fi31 * (fi42 + 32768)), F1: (4294967295 == 65535), F2: 18446744073709551615}
                                                        _ = v89
                                                    default:
                                                        v2 = false
                                                    }
                                                }
                                                var v90 uint8 = v5
                                                if (908 == 0) {
                                                    if (559 == 42) {
                                                        if v13 {
                                                            v91 := GenS11{F0: 2, F1: (^(fi31))}
                                                            _ = v91
                                                        } else {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                        if v13 {
                                                            v3 = v3
                                                            v90 = 255
                                                            fmt.Println(v1, v5)
                                                            v92 := GenS6{F0: 472, F1: 1, F2: 281}
                                                            _ = v92
                                                        } else {
                                                            fmt.Println(fi31, v3, fi42)
                                                            v93 := GenS5{F0: v90}
                                                            v90 = v90
                                                            _ = v93
                                                        }
                                                    }
                                                }
                                                _ = v90
                                            }
                                        case "a":
                                            if (!v2) {
                                                v94 := GenS3{F0: 42, F1: 721}
                                                var v95 int32 = 738
                                                _ = v94
                                                _ = v95
                                            } else {
                                                if len(v4) > 0 {
                                                    v4[0] = "a"
                                                }
                                                v96 := GenS12{F0: (^(v5)), F1: 18446744073709551615}
                                                _ = v96
                                            }
                                        case "\n\t":
                                            v2 = true
                                        }
                                        fmt.Println(v13, v5)
                                    }
                                    _ = v62
                                } else {
                                    v13 = (fi42 > fi42)
                                    _ = len(v4)
                                    var v97 uint32 = 4294967295
                                    _ = v97
                                }
                            }
                        }
                        var v98 int = fi31
                        _ = fi42
                        _ = v98
                    }
                    v1 = -2.25
                }
            case 42:
                for fi99 := 0; fi99 < 2; fi99++ {
                    v100 := GenS7{F0: 726, F1: 0, F2: "foo bar"}
                    var v101 int32 = 974
                    _ = fi99
                    _ = v100
                    _ = v101
                }
            case 32767:
                if true {
                    var v102 int32 = -1
                    v13 = ((!true) && (157 == 127))
                    fmt.Println(v5, fi31)
                    _ = v102
                } else {
                    if (fi31 <= (^(fi31))) {
                        switch fi31 {
                        case -1:
                            for fi103 := 0; fi103 < 2; fi103++ {
                                v2 = (fi31 <= (^(fi103)))
                                _ = fi103
                            }
                        case -2147483648:
                            v13 = (42 != v5)
                        }
                    } else {
                        switch ("" + "hello") {
                        case "hello":
                            v104 := GenS4{F0: (2 < 42), F1: "a", F2: 128}
                            _ = v104
                        case "a":
                            var v105 int16 = -32768
                            _ = v105
                        }
                        v106 := GenS11{F0: 65535, F1: genFn0(0)}
                        v2 = v13
                        _ = v106
                    }
                }
            default:
                fmt.Println(v5, v2)
            }
        case "\n\t":
            fmt.Println(v5, v3)
        default:
            for fi107 := 0; fi107 < 9; fi107++ {
                if (1 != 0) {
                    v108 := GenS13{F0: 0, F1: ("hello" + "foo bar")}
                    v5 += v5
                    _ = v108
                }
                _ = fi107
            }
        }
        v109 := GenS6{F0: 4294967295, F1: -1, F2: 2}
        _ = fi31
        _ = v109
    }
    switch "a" {
    case "\n\t":
        v110 := GenS3{F0: -128, F1: 65535}
        _ = v110
    }
    if ((!v2) && v13) {
        fmt.Println(v2, v13, v5)
        switch (("foo bar" + "foo bar") + "foo bar") {
        case "":
            switch 2147483647 {
            case 2147483647:
                if (65535 != 255) {
                    fmt.Println(v5, v1, v5)
                    if v13 {
                        v111 := GenS14{F0: -1}
                        _ = v111
                    }
                }
            case 32768:
                v112 := GenS8{F0: ("" + "foo bar")}
                _ = v112
            default:
                for fi113 := 0; fi113 < 16; fi113++ {
                    for fi114 := 0; fi114 < 4; fi114++ {
                        switch ("hello" + (("foo bar" + ("\n\t" + (((("\n\t" + "\n\t") + "\n\t") + "\n\t") + "hello"))) + "hello")) {
                        case "hello":
                            fmt.Println(v3)
                        default:
                            var v115 string = ""
                            _ = v115
                        }
                        var v116 float64 = v3
                        _ = fi114
                        _ = v116
                    }
                    v5 = v5
                    v117 := GenS15{F0: v5}
                    _ = fi113
                    _ = v117
                }
            }
        default:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        if (9223372036854775807 <= 9223372036854775807) {
            v13 = ((2 == 2) || (v2 && v2))
        }
    } else {
        if ((32767 < 65535) || v13) {
            var v118 int8 = 96
            for fi119 := 0; fi119 < 1; fi119++ {
                if (!(0 > 707)) {
                    v2 = v13
                } else {
                    if (false && v13) {
                        var v120 uint32 = 227
                        _ = v120
                    }
                    if (0 <= 42) {
                        var v121 float64 = (((((v3 + (1.5 / float64(0.1))) * (3.14159 + (-2.25 + v3))) / 1.0) - float64(v1)) / v3)
                        fmt.Println(fi119, fi119, fi119)
                        v122 := GenS16{F0: 218}
                        _ = v121
                        _ = v122
                    }
                    if len(v4) > 0 {
                        v4[0] = ""
                    }
                }
                var v123 string = "hello"
                var v124 bool = (!(uint(v5) > 609))
                _ = fi119
                _ = v123
                _ = v124
            }
            _ = v118
        } else {
            var v125 int64 = 425
            if v13 {
                if (-2147483648 >= 2147483647) {
                    v13 = (v2 || true)
                    v126 := GenS5{F0: 62}
                    if v2 {
                        var v127 float32 = 1.5
                        _ = v127
                    } else {
                        for fi128 := 0; fi128 < 18; fi128++ {
                            v13 = (371 < -119)
                            var v129 bool = true
                            _ = fi128
                            _ = v129
                        }
                        if len(v4) > 0 {
                            v4[0] = ""
                        }
                        v130 := GenS16{F0: 483}
                        _ = v130
                    }
                    _ = v126
                }
            }
            _ = v125
        }
        fmt.Println(v3, v2)
    }
    if (18446744073709551615 <= 255) {
        var v131 uint8 = v5
        var v132 bool = (2147483647 > 227)
        _ = v131
        _ = v132
    } else {
        fmt.Println(v13, v5, v5)
        var v133 int8 = -128
        if false {
            v134 := GenS3{F0: -128, F1: 42}
            if false {
                if (!v13) {
                    v3 = v3
                }
            }
            _ = v134
        }
        _ = v133
    }
    for fi135 := 0; fi135 < 15; fi135++ {
        if true {
            switch (^(fi135)) {
            case -9223372036854775808:
                fmt.Println(v5, v3)
            case 572:
                var v136 float64 = float64(((v1 - 0.0) / 1.5))
                _ = v136
            case 74:
                if v13 {
                    for fi137 := 0; fi137 < 7; fi137++ {
                        var v138 string = "a"
                        _ = fi137
                        _ = v138
                    }
                    for fi139 := 0; fi139 < 4; fi139++ {
                        switch genFn0(0) {
                        case 26:
                            v3 = (float64(0.1) + 0.0)
                        case 9223372036854775807:
                            v2 = (65535 <= 92)
                        default:
                            var v140 uint8 = (v5 >> ((18446744073709551615) & 63))
                            _ = v140
                        }
                        switch "hello" {
                        case "\n\t":
                            switch ("a" + ("" + ("foo bar" + "a"))) {
                            case "\n\t":
                                if v2 {
                                    fmt.Println(v1)
                                    v3 = v3
                                    _ = len(v4)
                                }
                            case "a":
                                var v141 uint16 = 42
                                _ = v141
                            case "":
                                v1 += ((v1 / (v1 - ((3.14159 + v1) + (v1 + ((float32(((float64(0.0) - v3) + float64((v1 * v1)))) - v1) + v1))))) * ((v1 / (v1 + ((0.0 - v1) - ((v1 - v1) / (1e10 / 0.1))))) + (3.14159 - (v1 * v1))))
                            default:
                                v5 = v5
                            }
                        case "foo bar":
                            v3 += v3
                        }
                        _ = fi139
                    }
                    var v142 uint64 = 169
                    _ = v142
                } else {
                    v3 -= (float64(3.14159) + v3)
                    fmt.Println(v5)
                    if (!(v2 && ((^(v5)) != (v5 * (v5 % ((42) | 1)))))) {
                        v2 = (0 >= -113)
                        if (v2 && ((0.1 < v3) || (42 != 1))) {
                            fmt.Println(v13)
                            v3 = v3
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        v143 := GenS6{F0: 519, F1: 1, F2: 1}
                        _ = v143
                    } else {
                        v3 -= 0.0
                        fmt.Println(v13, fi135, fi135)
                    }
                }
            default:
                var v144 int8 = 100
                _ = v144
            }
        }
        fmt.Println(v13, v13)
        _ = fi135
    }
    _ = len(v4)
    if len(v4) > 0 {
        v4[0] = (("a" + ("\n\t" + "")) + ((("a" + "") + (("a" + ("\n\t" + "hello")) + "")) + (((("a" + "a") + "") + "foo bar") + "\n\t")))
    }
    if false {
        if ((("foo bar" + "a") + ("" + "\n\t")) >= "") {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v145 uint64 = 0
            v145 += 42
            _ = v145
        } else {
            var v146 uint = 385
            switch 9223372036854775807 {
            case 890:
                v2 = ((!(-262 != -9223372036854775808)) && v13)
            case 9223372036854775807:
                v146 -= 1
            case 32768:
                var v147 int16 = 42
                _ = v147
            default:
                var v148 uint16 = 853
                _ = v148
            }
            switch "foo bar" {
            case "":
                v146 = (^(v146))
            }
            _ = v146
        }
        if false {
            v2 = (1 <= 1)
            for fi149 := 0; fi149 < 8; fi149++ {
                var v150 int64 = 9223372036854775807
                if (1 != -207) {
                    v2 = (421 != 509)
                } else {
                    v151 := GenS17{F0: 0}
                    _ = v151
                }
                var v152 int8 = 127
                _ = fi149
                _ = v150
                _ = v152
            }
            for fi153 := 0; fi153 < 4; fi153++ {
                var v154 uint = 1
                var v155 uint16 = 1
                _ = fi153
                _ = v154
                _ = v155
            }
        } else {
            var v156 int16 = -401
            fmt.Println(v5, v5)
            v157 := GenS18{F0: 2, F1: 798}
            _ = v156
            _ = v157
        }
        v3 = (v3 - v3)
    } else {
        for fi158 := 0; fi158 < 6; fi158++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch fi158 {
            case -9223372036854775808:
                var v159 int32 = -2147483648
                _ = v159
            case 2147483647:
                var v160 uint16 = 255
                _ = v160
            }
            _ = fi158
        }
        for fi161 := 0; fi161 < 20; fi161++ {
            switch "hello" {
            case "hello":
                var v162 int = fi161
                _ = v162
            case "\n\t":
                var v163 uint64 = 2
                _ = v163
            case "a":
                var v164 uint = 71
                _ = v164
            default:
                if v2 {
                    if false {
                        if true {
                            var v165 int8 = 127
                            for fi166 := 0; fi166 < 9; fi166++ {
                                v5 = (v5 >> ((uint(v5)) & 63))
                                v4 = append(v4, "hello")
                                _ = fi166
                            }
                            _ = v165
                        }
                        v2 = ((^(v5)) != v5)
                    }
                    fmt.Println(fi161, v13)
                }
            }
            v1 = v1
            _ = fi161
        }
        for fi167 := 0; fi167 < 14; fi167++ {
            if (7 > (v5 + 42)) {
                for fi168 := 0; fi168 < 10; fi168++ {
                    v169 := GenS18{F0: 42, F1: 635}
                    if (!true) {
                        v4 = append(v4, (("" + ("" + "hello")) + ((("\n\t" + "foo bar") + ("" + "hello")) + ("foo bar" + (((("hello" + ("hello" + ("foo bar" + ("" + "a")))) + "") + ("hello" + "\n\t")) + "")))))
                        v170 := GenS17{F0: 42}
                        v13 = (!(42 != 65535))
                        if v13 {
                            var v171 uint = 42
                            _ = len(v4)
                            _ = v171
                        } else {
                            var v172 string = ("a" + ("" + ("a" + (("a" + "") + "foo bar"))))
                            _ = v172
                        }
                        v173 := GenS12{F0: 121, F1: 446}
                        _ = v170
                        _ = v173
                    }
                    fmt.Println(fi168, v5, fi168)
                    _ = fi168
                    _ = v169
                }
                if (false && false) {
                    var v174 uint = 966
                    v175 := GenS15{F0: v5}
                    v5 = 16
                    _ = v174
                    _ = v175
                } else {
                    for fi176 := 0; fi176 < 2; fi176++ {
                        v2 = v13
                        v177 := GenS16{F0: 365}
                        v5 -= v5
                        _ = fi176
                        _ = v177
                    }
                }
            }
            if (1000 >= 42) {
                v178 := GenS17{F0: 65535}
                _ = len(v4)
                switch "foo bar" {
                case "foo bar":
                    v2 = ((-(fi167)) != genFn0(2147483647))
                case "\n\t":
                    var v179 int8 = 0
                    _ = v179
                case "":
                    var v180 uint8 = (v5 / ((v5) | 1))
                    _ = v180
                default:
                    switch fi167 {
                    case 32768:
                        for fi181 := 0; fi181 < 8; fi181++ {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            fmt.Println(fi167)
                            v2 = v13
                            _ = fi181
                        }
                    case -1:
                        if (128 >= genFn0(864)) {
                            switch 976 {
                            case 314:
                                if false {
                                    if len(v4) > 0 {
                                        v4[0] = "hello"
                                    }
                                }
                            case -2147483648:
                                var v182 int = genFn0(0)
                                _ = v182
                            case 128:
                                v3 = v3
                            }
                        }
                    case -9223372036854775808:
                        var v183 uint16 = 65535
                        _ = v183
                    default:
                        fmt.Println(v3, v2, v13)
                    }
                }
                _ = v178
            } else {
                v4 = append(v4, "hello")
            }
            _ = fi167
        }
    }
    v184 := GenS13{F0: 1000, F1: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v184
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
