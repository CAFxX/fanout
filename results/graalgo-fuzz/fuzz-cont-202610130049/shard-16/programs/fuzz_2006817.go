package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64; F1 int16 }
type GenS7 struct { F0 uint8; F1 uint8; F2 string }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint64; F1 uint64 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int16; F1 uint64; F2 int8 }
type GenS13 struct { F0 int32; F1 uint32 }
type GenS14 struct { F0 uint; F1 uint32; F2 int8 }
type GenS15 struct { F0 int; F1 uint8 }
type GenS16 struct { F0 int32; F1 int; F2 int64 }
func genFn0(p0 uint8, p1 uint, p2 int16) int {
    fmt.Println(p1, p0)
    var v5 uint32 = uint32(p1)
    p1 = (p1 + uint(p0))
    v6 := GenS1{F0: 2, F1: 1, F2: -1}
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 371
}
func genFn1(p0 int32) int64 {
    for fi8 := 0; fi8 < 17; fi8++ {
        switch "foo bar" {
        case "":
            v9 := []int{(fi8 ^ (^(fi8))), fi8, (fi8 & -842), fi8}
            _ = v9
        case "hello":
            if false {
                if true {
                    if (!(42 != 747)) {
                        var v10 uint32 = 1
                        _ = v10
                    } else {
                        v11 := GenS1{F0: 255, F1: 2147483647, F2: 1}
                        _ = v11
                    }
                    switch 1 {
                    case -1:
                        if false {
                            switch ("hello" + "foo bar") {
                            case "":
                                fmt.Println(fi8, fi8)
                            }
                            var v12 uint64 = 91
                            _ = v12
                        } else {
                            fmt.Println(p0, p0)
                            for fi13 := 0; fi13 < 1; fi13++ {
                                v14 := []int{fi13, fi8, genFn0(48, 758, 1000)}
                                var v15 uint8 = 128
                                if false {
                                    v16 := GenS2{F0: v15}
                                    _ = v16
                                }
                                if (true && false) {
                                    fmt.Println(p0, fi8)
                                    if (-907 > int16(fi8)) {
                                        var v17 uint16 = 743
                                        _ = v17
                                    } else {
                                        p0 = int32(v15)
                                    }
                                } else {
                                    v14 = append(v14, -1)
                                    var v18 uint16 = 812
                                    switch fi13 {
                                    case 42:
                                        if len(v14) >= 2 {
                                            v14 = v14[1:len(v14)]
                                        }
                                    case 0:
                                        v19 := GenS1{F0: (v18 << ((65535) & 63)), F1: 2147483648, F2: -1}
                                        _ = v19
                                    default:
                                        fmt.Println(fi8, v15, fi8)
                                    }
                                    _ = v18
                                }
                                _ = fi13
                                _ = v14
                                _ = v15
                            }
                            p0 = 1
                        }
                    case -544:
                        v20 := GenS3{F0: 0, F1: (("a" + (("foo bar" + (((("hello" + (("" + "foo bar") + ("" + "foo bar"))) + "foo bar") + "") + ("foo bar" + ("hello" + ("hello" + "\n\t"))))) + "foo bar")) + "a")}
                        _ = v20
                    case 127:
                        p0 = (^(p0))
                    }
                    if (!false) {
                        p0 = p0
                    } else {
                        v21 := GenS2{F0: 116}
                        _ = v21
                    }
                } else {
                    p0 += 686
                    for fi22 := 0; fi22 < 3; fi22++ {
                        fmt.Println(fi8, fi8)
                        _ = fi22
                    }
                    p0 -= (^(p0))
                }
            } else {
                var v23 uint32 = 2147483648
                fmt.Println(p0)
                _ = v23
            }
        case "a":
            switch 127 {
            case 32767:
                if (!true) {
                    v24 := []int{-336, fi8}
                    _ = v24
                } else {
                    switch ("a" + ("hello" + "a")) {
                    case "":
                        fmt.Println(p0, fi8)
                    default:
                        p0 -= p0
                    }
                    if (!(((3.14159 / (0.1 - 0.0)) * 1.5) <= -2.25)) {
                        var v25 uint8 = 0
                        for fi26 := 0; fi26 < 5; fi26++ {
                            if (int(p0) < fi8) {
                                v25 = (v25 | v25)
                                switch "" {
                                case "\n\t":
                                    p0 -= (-(p0))
                                }
                            } else {
                                v27 := GenS4{F0: (fi26 + fi8), F1: (p0 & p0)}
                                p0 = -1
                                if true {
                                    switch "a" {
                                    case "foo bar":
                                        v28 := []int{genFn0((v25 << ((18446744073709551615) & 63)), 1, 32767)}
                                        _ = v28
                                    case "a":
                                        p0 += p0
                                    case "":
                                        switch "a" {
                                        case "a":
                                            v29 := []int{(fi8 & fi26), -579, (fi26 ^ fi8)}
                                            _ = v29
                                        case "":
                                            if true {
                                                p0 += 77
                                                if false {
                                                    v30 := []int{genFn0((v25 & (v25 | v25)), 0, -877), (fi8 * genFn0((v25 / (((v25 + (^(v25)))) | 1)), 1, 0))}
                                                    if false {
                                                        v31 := GenS5{F0: (p0 % ((-746) | 1))}
                                                        var v32 uint64 = 493
                                                        if false {
                                                            if (2147483647 == 999) {
                                                                var v33 string = "hello"
                                                                var v34 uint = 255
                                                                if (1e10 <= (0.0 - 1e10)) {
                                                                    var v35 string = (v33 + (("foo bar" + v33) + v33))
                                                                    switch (v33 + (v35 + v33)) {
                                                                    case "\n\t":
                                                                        var v36 uint8 = (^(v25))
                                                                        _ = len(v30)
                                                                        _ = v36
                                                                    case "foo bar":
                                                                        v37 := GenS1{F0: 0, F1: 2, F2: 1}
                                                                        _ = v37
                                                                    case "":
                                                                        fmt.Println(fi26, v33)
                                                                    default:
                                                                        if len(v30) >= 2 {
                                                                            v30 = v30[1:len(v30)]
                                                                        }
                                                                    }
                                                                    _ = v35
                                                                } else {
                                                                    if (1 < 548) {
                                                                        if (fi26 != (fi8 / ((int(p0)) | 1))) {
                                                                            if (0.1 > 3.14159) {
                                                                                v33 = "hello"
                                                                                if true {
                                                                                    v38 := GenS6{F0: 591, F1: -1}
                                                                                    _ = v38
                                                                                } else {
                                                                                    v39 := GenS7{F0: v25, F1: v25, F2: "hello"}
                                                                                    v40 := GenS3{F0: 178, F1: (v33 + v33)}
                                                                                    switch fi26 {
                                                                                    case 938:
                                                                                        var v41 int = (fi26 & fi8)
                                                                                        _ = v41
                                                                                    default:
                                                                                        v30 = append(v30, fi26)
                                                                                        _ = len(v30)
                                                                                        var v42 uint64 = uint64(v34)
                                                                                        _ = v42
                                                                                    }
                                                                                    _ = v39
                                                                                    _ = v40
                                                                                }
                                                                                v43 := GenS8{F0: ""}
                                                                                v33 = "\n\t"
                                                                                _ = v43
                                                                            } else {
                                                                                fmt.Println(v32, fi26)
                                                                                switch ((v33 + "hello") + (v33 + ("" + (v33 + "a")))) {
                                                                                case "foo bar":
                                                                                    var v44 int = (fi26 << ((v34) & 63))
                                                                                    if (fi8 < (^(fi26))) {
                                                                                        v45 := GenS6{F0: 42, F1: 154}
                                                                                        v32 = v32
                                                                                        var v46 string = v33
                                                                                        _ = v45
                                                                                        _ = v46
                                                                                    }
                                                                                    _ = v44
                                                                                case "hello":
                                                                                    var v47 bool = (fi8 >= fi8)
                                                                                    _ = v47
                                                                                }
                                                                                if true {
                                                                                    var v48 uint16 = 573
                                                                                    var v49 int16 = -32768
                                                                                    _ = v48
                                                                                    _ = v49
                                                                                } else {
                                                                                    var v50 uint = (v34 / ((21) | 1))
                                                                                    var v51 uint16 = 65535
                                                                                    _ = v50
                                                                                    _ = v51
                                                                                }
                                                                            }
                                                                            if true {
                                                                                var v52 int8 = 127
                                                                                _ = v52
                                                                            } else {
                                                                                fmt.Println(fi26, v25, p0)
                                                                                var v53 int64 = 9223372036854775807
                                                                                v34 -= v34
                                                                                _ = v53
                                                                            }
                                                                            v34 = v34
                                                                        }
                                                                        switch genFn0((v25 | 127), v34, 1000) {
                                                                        case -834:
                                                                            p0 = p0
                                                                        default:
                                                                            var v54 float32 = 0.0
                                                                            _ = v54
                                                                        }
                                                                        switch "" {
                                                                        case "":
                                                                            if len(v30) >= 2 {
                                                                                v30 = v30[1:len(v30)]
                                                                            }
                                                                        }
                                                                    } else {
                                                                        if true {
                                                                            v55 := GenS9{F0: v32}
                                                                            _ = v55
                                                                        }
                                                                        _ = len(v30)
                                                                    }
                                                                    var v56 uint = (v34 ^ (v34 >> ((2) & 63)))
                                                                    _ = v56
                                                                }
                                                                _ = v33
                                                                _ = v34
                                                            }
                                                            var v57 int = -886
                                                            _ = v57
                                                        }
                                                        _ = v31
                                                        _ = v32
                                                    } else {
                                                        if len(v30) >= 2 {
                                                            v30 = v30[1:len(v30)]
                                                        }
                                                        v58 := GenS9{F0: 1}
                                                        switch "hello" {
                                                        case "a":
                                                            if true {
                                                                fmt.Println(fi26)
                                                                if (("foo bar" + "hello") <= ((("" + "hello") + "a") + "")) {
                                                                    if true {
                                                                        fmt.Println(fi26, p0, p0)
                                                                        fmt.Println(fi26, fi26, v25)
                                                                        var v59 float64 = (1e10 * 0.1)
                                                                        _ = v59
                                                                    }
                                                                    var v60 int32 = 1
                                                                    _ = v60
                                                                } else {
                                                                    v25 = 26
                                                                    var v61 uint32 = 2147483647
                                                                    _ = v61
                                                                }
                                                                if ((!false) && false) {
                                                                    v62 := GenS3{F0: 32767, F1: (("a" + ("hello" + "hello")) + "a")}
                                                                    _ = v62
                                                                } else {
                                                                    if len(v30) >= 2 {
                                                                        v30 = v30[1:len(v30)]
                                                                    }
                                                                    var v63 int64 = -9223372036854775808
                                                                    if len(v30) > 0 {
                                                                        v30[0] = (fi26 % (((fi26 + fi8)) | 1))
                                                                    }
                                                                    _ = v63
                                                                }
                                                            }
                                                        case "foo bar":
                                                            switch "" {
                                                            case "hello":
                                                                if true {
                                                                    v64 := GenS9{F0: 18446744073709551615}
                                                                    _ = v64
                                                                }
                                                            case "\n\t":
                                                                if len(v30) >= 2 {
                                                                    v30 = v30[1:len(v30)]
                                                                }
                                                            case "foo bar":
                                                                if false {
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        fmt.Println(v25)
                                                                        var v65 int8 = -9
                                                                        v65 = (-(v65))
                                                                        _ = v65
                                                                    case "":
                                                                        p0 = p0
                                                                    default:
                                                                        if len(v30) > 0 {
                                                                            v30[0] = (fi8 & fi8)
                                                                        }
                                                                    }
                                                                    var v66 bool = true
                                                                    var v67 uint16 = 255
                                                                    _ = v66
                                                                    _ = v67
                                                                }
                                                            }
                                                        case "hello":
                                                            v68 := GenS8{F0: "foo bar"}
                                                            _ = v68
                                                        default:
                                                            var v69 int32 = p0
                                                            if len(v30) >= 2 {
                                                                v30 = v30[1:len(v30)]
                                                            }
                                                            v70 := GenS4{F0: genFn0((v25 + v25), 773, 1000), F1: v69}
                                                            v69 = v69
                                                            _ = v69
                                                            _ = v70
                                                        }
                                                        _ = v58
                                                    }
                                                    _ = v30
                                                } else {
                                                    v71 := []int{}
                                                    p0 += p0
                                                    if len(v71) > 0 {
                                                        v71[0] = 127
                                                    }
                                                    _ = v71
                                                }
                                            }
                                            switch ("\n\t" + "foo bar") {
                                            case "":
                                                if true {
                                                    var v72 float64 = 1.5
                                                    if (false || false) {
                                                        var v73 int = genFn0((v25 << ((4294967295) & 63)), 255, 42)
                                                        if (false && false) {
                                                            if ((!false) && (427 < 2147483648)) {
                                                                switch ("a" + ("a" + "foo bar")) {
                                                                case "":
                                                                    v74 := []int{int(p0), fi8, 45}
                                                                    _ = v74
                                                                case "\n\t":
                                                                    v75 := GenS4{F0: 1, F1: p0}
                                                                    _ = v75
                                                                case "hello":
                                                                    v76 := GenS4{F0: (v73 | (fi8 | (v73 & -885))), F1: p0}
                                                                    _ = v76
                                                                default:
                                                                    var v77 float64 = (0.0 + 0.1)
                                                                    _ = v77
                                                                }
                                                                switch (("" + ("" + ("a" + "hello"))) + "foo bar") {
                                                                case "":
                                                                    v72 += v72
                                                                case "a":
                                                                    var v78 int = genFn0(v25, 553, 0)
                                                                    _ = v78
                                                                }
                                                                v72 = v72
                                                                if ((357 <= 822) || (0 != 42)) {
                                                                    v79 := []int{fi26}
                                                                    v80 := GenS10{F0: 862, F1: 969}
                                                                    _ = v79
                                                                    _ = v80
                                                                }
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    switch ("foo bar" + "hello") {
                                                                    case "":
                                                                        v25 += v25
                                                                    default:
                                                                        var v81 uint = 255
                                                                        _ = v81
                                                                    }
                                                                    switch fi26 {
                                                                    case -736:
                                                                        var v82 uint = 593
                                                                        _ = v82
                                                                    case 1:
                                                                        var v83 uint16 = 32767
                                                                        _ = v83
                                                                    case 32768:
                                                                        var v84 uint = 2
                                                                        _ = v84
                                                                    }
                                                                    v85 := []int{(-(fi26)), genFn0((v25 | v25), 0, -32768), -9223372036854775808, genFn0(v25, 4294967295, 900)}
                                                                    _ = v85
                                                                default:
                                                                    var v86 uint32 = 0
                                                                    _ = v86
                                                                }
                                                                fmt.Println(v25, v73)
                                                            } else {
                                                                var v87 uint8 = 1
                                                                fmt.Println(v25)
                                                                _ = v87
                                                            }
                                                            v73 = 913
                                                        }
                                                        v73 = fi8
                                                        _ = v73
                                                    } else {
                                                        fmt.Println(fi26, v25, fi8)
                                                        var v88 int = (fi8 & fi8)
                                                        var v89 int32 = p0
                                                        fmt.Println(v72, v89)
                                                        var v90 int = fi26
                                                        _ = v88
                                                        _ = v89
                                                        _ = v90
                                                    }
                                                    _ = v72
                                                } else {
                                                    v91 := GenS11{F0: -370}
                                                    fmt.Println(fi8, v25, v25)
                                                    fmt.Println(p0)
                                                    _ = v91
                                                }
                                            }
                                        case "\n\t":
                                            v92 := GenS6{F0: -9223372036854775808, F1: -483}
                                            _ = v92
                                        }
                                    default:
                                        fmt.Println(fi26, p0)
                                    }
                                    if (false || false) {
                                        var v93 float64 = 0.0
                                        switch (fi8 ^ (fi8 ^ (fi26 | -9223372036854775808))) {
                                        case 32767:
                                            v93 = v93
                                        case 0:
                                            var v94 float32 = (3.14159 + float32(float64(0.1)))
                                            _ = v94
                                        case 128:
                                            switch "a" {
                                            case "a":
                                                v95 := GenS5{F0: 42}
                                                _ = v95
                                            default:
                                                switch "hello" {
                                                case "a":
                                                    if true {
                                                        p0 -= p0
                                                    }
                                                case "\n\t":
                                                    switch (fi26 / (((fi8 | 160)) | 1)) {
                                                    case 4294967295:
                                                        v96 := []int{(fi26 ^ (-(fi8)))}
                                                        _ = v96
                                                    case 32768:
                                                        var v97 int16 = 55
                                                        _ = v97
                                                    case -989:
                                                        v25 += v25
                                                    default:
                                                        if (229 >= 2) {
                                                            v98 := []int{fi8, genFn0(235, 42, -32768)}
                                                            fmt.Println(v93, fi26)
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                v25 += (v25 / ((128) | 1))
                                                            case "\n\t":
                                                                var v99 int16 = 1
                                                                _ = v99
                                                            case "":
                                                                v100 := GenS6{F0: -9223372036854775808, F1: 32767}
                                                                _ = v100
                                                            }
                                                            if false {
                                                                fmt.Println(fi8)
                                                                var v101 bool = false
                                                                if true {
                                                                    var v102 int16 = 1000
                                                                    p0 = p0
                                                                    _ = v102
                                                                } else {
                                                                    fmt.Println(v101, fi26)
                                                                    if true {
                                                                        if ((101 > -9223372036854775808) && (!v101)) {
                                                                            v103 := GenS9{F0: 819}
                                                                            v25 = 255
                                                                            _ = v103
                                                                        }
                                                                    } else {
                                                                        p0 = p0
                                                                        v104 := GenS6{F0: 0, F1: 42}
                                                                        _ = v104
                                                                    }
                                                                    if false {
                                                                        v105 := GenS1{F0: 42, F1: 65535, F2: 0}
                                                                        _ = v105
                                                                    } else {
                                                                        var v106 string = ""
                                                                        v106 = v106
                                                                        _ = v106
                                                                    }
                                                                }
                                                                _ = v101
                                                            }
                                                            if len(v98) >= 2 {
                                                                v98 = v98[1:len(v98)]
                                                            }
                                                            var v107 uint = 255
                                                            var v108 float32 = 1.5
                                                            _ = v107
                                                            _ = v108
                                                            _ = v98
                                                        }
                                                    }
                                                }
                                            }
                                        default:
                                            var v109 float64 = ((1.5 * float64(0.0)) + (((v93 + float64(((3.14159 + 0.1) * (0.0 * 1.5)))) / v93) - v93))
                                            _ = v109
                                        }
                                        _ = v93
                                    } else {
                                        v110 := []int{471}
                                        var v111 int8 = 0
                                        v112 := GenS3{F0: 2, F1: "foo bar"}
                                        _ = v110
                                        _ = v111
                                        _ = v112
                                    }
                                    p0 += (p0 * 2147483647)
                                } else {
                                    v113 := GenS12{F0: 1, F1: 2, F2: -128}
                                    var v114 int16 = -1
                                    if (true && false) {
                                        var v115 uint8 = (v25 + v25)
                                        if ("foo bar" < ("a" + ((("" + "hello") + (("" + "a") + "a")) + "foo bar"))) {
                                            v115 = (v115 & v25)
                                            var v116 float64 = (0.1 / (-2.25 / 1.0))
                                            _ = v116
                                        } else {
                                            p0 += (p0 / ((p0) | 1))
                                            v117 := []int{(^(fi8)), genFn0(v115, 1, 32767), -574, (-(fi8))}
                                            _ = v117
                                        }
                                        var v118 int16 = -486
                                        v114 -= (-(v118))
                                        if (!("a" != ("\n\t" + ""))) {
                                            fmt.Println(v25, fi8, v114)
                                            var v119 uint16 = 1
                                            p0 = p0
                                            _ = v119
                                        }
                                        _ = v115
                                        _ = v118
                                    }
                                    _ = v113
                                    _ = v114
                                }
                                _ = v27
                            }
                            _ = fi26
                        }
                        _ = v25
                    }
                    switch "a" {
                    case "a":
                        if false {
                            for fi120 := 0; fi120 < 12; fi120++ {
                                v121 := []int{1, (fi8 / ((genFn0(180, 4294967295, 1)) | 1)), fi120}
                                if (127 > 42) {
                                    if true {
                                        var v122 float32 = (1e10 + -2.25)
                                        if (!true) {
                                            v123 := GenS4{F0: (^(fi120)), F1: p0}
                                            v124 := GenS8{F0: ""}
                                            _ = v123
                                            _ = v124
                                        }
                                        v122 += float32(1.5)
                                        p0 = p0
                                        v125 := GenS13{F0: p0, F1: 527}
                                        _ = v122
                                        _ = v125
                                    }
                                    fmt.Println(fi120)
                                    p0 -= 1
                                }
                                v121 = append(v121, 4294967295)
                                if ((-2.25 - -2.25) > (1.5 - ((1e10 / 1.0) / float32(float64(1.5))))) {
                                    switch fi120 {
                                    case -578:
                                        var v126 uint8 = 1
                                        _ = v126
                                    case 32767:
                                        if false {
                                            var v127 int64 = 9223372036854775807
                                            if false {
                                                fmt.Println(fi120, fi8)
                                                if false {
                                                    var v128 uint8 = 127
                                                    _ = v128
                                                } else {
                                                    p0 -= (p0 * 1)
                                                }
                                            } else {
                                                var v129 int16 = -311
                                                if (!true) {
                                                    v129 = int16(p0)
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        var v130 uint8 = 255
                                                        _ = v130
                                                    case "hello":
                                                        v129 = -618
                                                    case "":
                                                        if true {
                                                            var v131 float64 = float64(0.1)
                                                            _ = v131
                                                        } else {
                                                            var v132 float32 = float32(1e10)
                                                            var v133 uint32 = 984
                                                            if (4294967295 == 1000) {
                                                                v127 = v127
                                                                if len(v121) > 0 {
                                                                    v121[0] = fi120
                                                                }
                                                                switch fi120 {
                                                                case 9223372036854775807:
                                                                    if true {
                                                                        v121 = append(v121, 4294967295)
                                                                        var v134 uint64 = 4294967295
                                                                        var v135 uint32 = v133
                                                                        v134 += v134
                                                                        fmt.Println(fi120)
                                                                        _ = v134
                                                                        _ = v135
                                                                    }
                                                                case 42:
                                                                    p0 -= (p0 & (p0 % (((p0 | 0)) | 1)))
                                                                    var v136 string = (("" + "") + ("" + ("" + "foo bar")))
                                                                    _ = v136
                                                                case 32768:
                                                                    var v137 uint64 = 18446744073709551615
                                                                    _ = v137
                                                                }
                                                                v132 = float32(1.5)
                                                            } else {
                                                                if len(v121) > 0 {
                                                                    v121[0] = fi120
                                                                }
                                                            }
                                                            _ = v132
                                                            _ = v133
                                                        }
                                                        var v138 int32 = (p0 * p0)
                                                        if (v129 > -32768) {
                                                            var v139 int64 = v127
                                                            _ = v139
                                                        }
                                                        _ = v138
                                                    }
                                                } else {
                                                    var v140 uint16 = 1
                                                    _ = v140
                                                }
                                                var v141 int16 = 32767
                                                fmt.Println(p0, p0)
                                                _ = v129
                                                _ = v141
                                            }
                                            v142 := GenS9{F0: 18446744073709551615}
                                            _ = v127
                                            _ = v142
                                        } else {
                                            _ = len(v121)
                                            fmt.Println(p0, fi120)
                                            if len(v121) > 0 {
                                                v121[0] = fi8
                                            }
                                        }
                                        v121 = append(v121, fi120)
                                        var v143 float64 = (0.0 - 1.5)
                                        _ = v143
                                    }
                                    v144 := GenS14{F0: 65535, F1: 285, F2: 127}
                                    var v145 uint = 705
                                    _ = v144
                                    _ = v145
                                } else {
                                    var v146 uint16 = 65535
                                    var v147 int64 = 340
                                    if false {
                                        v148 := GenS15{F0: -235, F1: 128}
                                        fmt.Println(p0, v146, v147)
                                        _ = v148
                                    } else {
                                        var v149 uint64 = 1
                                        _ = v149
                                    }
                                    v150 := GenS16{F0: p0, F1: int(v146), F2: (^(v147))}
                                    _ = v146
                                    _ = v147
                                    _ = v150
                                }
                                p0 = p0
                                _ = len(v121)
                                _ = fi120
                                _ = v121
                            }
                            p0 = -961
                            var v151 uint = 808
                            _ = v151
                        } else {
                            v152 := []int{-21, fi8, (fi8 & -30)}
                            _ = v152
                        }
                    case "hello":
                        if false {
                            p0 -= p0
                            var v153 int8 = int8(fi8)
                            _ = v153
                        }
                    case "":
                        var v154 uint64 = 42
                        _ = v154
                    default:
                        p0 = p0
                    }
                }
            case 4294967295:
                v155 := GenS4{F0: (fi8 >> ((255) & 63)), F1: p0}
                _ = v155
            default:
                switch fi8 {
                case 32767:
                    fmt.Println(fi8, fi8)
                }
            }
        default:
            var v156 bool = false
            _ = v156
        }
        p0 = (p0 ^ p0)
        var v157 int32 = p0
        _ = fi8
        _ = v157
    }
    for fi158 := 0; fi158 < 19; fi158++ {
        var v159 int8 = 0
        fmt.Println(p0, p0)
        v160 := GenS2{F0: 255}
        _ = fi158
        _ = v159
        _ = v160
    }
    fmt.Println(p0)
    _ = p0
    return -1
}

func main() {
    var v1 uint8 = 127
    var v2 uint64 = 0
    var v3 float64 = 0.0
    var v4 int32 = 63
    fmt.Println(v2)
    v2 = 2
    fmt.Println(v3, v3)
    var v161 float64 = ((v3 * v3) + (v3 + v3))
    v2 += 0
    v162 := []int{-271}
    var v163 uint64 = v2
    switch ("foo bar" + ("" + (("" + "\n\t") + "foo bar"))) {
    case "a":
        for fi164 := 0; fi164 < 9; fi164++ {
            fmt.Println(v4, v4, v3)
            _ = fi164
        }
    }
    v161 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v161:", v161)
    fmt.Println("slicelen v162:", len(v162), v162)
    fmt.Println("v163:", v163)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v161
    _ = v162
    _ = v163
    _ = v2
    _ = v3
    _ = v4
}
