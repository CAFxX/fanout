package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int16 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS6 struct { F0 uint32; F1 bool; F2 int16 }
type GenS7 struct { F0 int32; F1 uint }
type GenS8 struct { F0 int64; F1 int }
type GenS9 struct { F0 int64; F1 int16; F2 uint8 }
type GenS10 struct { F0 int; F1 bool; F2 int32 }
type GenS11 struct { F0 uint8; F1 uint; F2 string }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 int32; F1 int64; F2 int }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 int8; F1 uint }
type GenS17 struct { F0 uint64; F1 uint; F2 string }
type GenS18 struct { F0 uint8 }
type GenS19 struct { F0 uint32; F1 uint32; F2 bool }
type GenS20 struct { F0 bool; F1 string; F2 uint32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = 979
    var v3 int8 = -70
    var v4 uint = 1
    var v5 float64 = 1e10
    if (!false) {
        v4 += (v4 << ((v4) & 63))
    } else {
        if true {
            v6 := GenS1{F0: 1, F1: 42, F2: (v2 << ((v4) & 63))}
            _ = v6
        } else {
            v7 := v1["hello"]
            _ = len(v1)
            _ = v7
        }
    }
    fmt.Println(v4)
    switch "\n\t" {
    case "foo bar":
        if (3.14159 == v5) {
            for fi8 := 0; fi8 < 2; fi8++ {
                v9 := GenS2{F0: "a", F1: 1}
                v5 = float64(0.1)
                _ = fi8
                _ = v9
            }
        } else {
            var v10 uint8 = 128
            fmt.Println(v5)
            _ = v10
        }
    case "":
        fmt.Println(v5, v3, v2)
    }
    fmt.Println(v3, v2, v4)
    for fi11 := 0; fi11 < 4; fi11++ {
        for fi12 := 0; fi12 < 3; fi12++ {
            fmt.Println(fi11, fi12)
            v13 := []int{(fi11 % ((-908) | 1)), fi12, -971, 128}
            v3 -= (v3 | (-(v3)))
            _ = fi12
            _ = v13
        }
        if (!(987 >= 1)) {
            fmt.Println(v5)
            v2 = 1
            v14 := GenS3{F0: 4294967295, F1: 42}
            _ = v14
        } else {
            switch "" {
            case "":
                var v15 int32 = -2147483648
                _ = v15
            default:
                v4 = 2
            }
            var v16 bool = true
            _ = v16
        }
        _ = fi11
    }
    delete(v1, "foo bar")
    if false {
        v17 := []int{127, -729}
        if true {
            var v18 int = 302
            for fi19 := 0; fi19 < 17; fi19++ {
                v20 := GenS4{F0: (-(v18)), F1: v3}
                v2 = (v2 + v2)
                _ = fi19
                _ = v20
            }
            for fi21 := 0; fi21 < 18; fi21++ {
                var v22 string = ((("hello" + ("" + "\n\t")) + "") + ("" + ("foo bar" + "hello")))
                var v23 uint64 = 812
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = v18
        } else {
            var v24 bool = (42 > 2)
            _ = v24
        }
        _ = v17
    } else {
        switch "a" {
        case "foo bar":
            switch "foo bar" {
            case "\n\t":
                if (!(9223372036854775807 <= -1)) {
                    var v25 float32 = ((0.1 / 1.0) * 3.14159)
                    if true {
                        for fi26 := 0; fi26 < 4; fi26++ {
                            v5 = float64(v25)
                            v2 -= v2
                            _ = fi26
                        }
                        for fi27 := 0; fi27 < 5; fi27++ {
                            for fi28 := 0; fi28 < 1; fi28++ {
                                _ = len(v1)
                                var v29 uint = v4
                                switch int(v3) {
                                case 1:
                                    v30 := v1[("foo bar" + ("a" + "hello"))]
                                    _ = v30
                                }
                                _ = fi28
                                _ = v29
                            }
                            for fi31 := 0; fi31 < 4; fi31++ {
                                v2 = v2
                                v32 := v1["\n\t"]
                                if true {
                                    v33 := v1["hello"]
                                    if (v25 < ((0.1 - (v25 / v25)) - 3.14159)) {
                                        var v34 uint16 = 255
                                        var v35 bool = (-1 >= 428)
                                        v2 = v2
                                        _ = v34
                                        _ = v35
                                    }
                                    _ = v33
                                } else {
                                    v36 := []int{}
                                    if (false && true) {
                                        if true {
                                            v4 -= v4
                                        }
                                        v2 += (v2 - (v2 / ((1) | 1)))
                                        fmt.Println(v2, v2, v32)
                                        if false {
                                            fmt.Println(v25)
                                        } else {
                                            v5 -= ((((v5 / (3.14159 + (v5 - -2.25))) + 0.0) * v5) + -2.25)
                                            if true {
                                                _ = len(v1)
                                                v37 := GenS2{F0: ((("foo bar" + "") + "foo bar") + "a"), F1: 895}
                                                var v38 bool = false
                                                _ = v37
                                                _ = v38
                                            }
                                            delete(v1, "hello")
                                            var v39 uint8 = 127
                                            v40 := GenS5{F0: 2, F1: (v4 | (^(v4))), F2: 0}
                                            switch fi27 {
                                            case 128:
                                                if true {
                                                    v32 -= v32
                                                    fmt.Println(v39, v4)
                                                }
                                                if len(v36) >= 2 {
                                                    v36 = v36[1:len(v36)]
                                                }
                                            case 127:
                                                if false {
                                                    v41 := GenS3{F0: uint64(v32), F1: (^(v39))}
                                                    if (-943 == 80) {
                                                        v42 := GenS3{F0: 1, F1: v39}
                                                        fmt.Println(v4, v4, v2)
                                                        _ = v42
                                                    } else {
                                                        var v43 uint8 = v39
                                                        _ = v43
                                                    }
                                                    v39 -= v39
                                                    var v44 int32 = 0
                                                    _ = v41
                                                    _ = v44
                                                } else {
                                                    if ((fi27 - (fi27 / ((fi27) | 1))) <= (fi31 + 772)) {
                                                        switch fi31 {
                                                        case 1:
                                                            if true {
                                                                if true {
                                                                    switch -9223372036854775808 {
                                                                    case 0:
                                                                        var v45 float32 = v25
                                                                        _ = v45
                                                                    case 2147483647:
                                                                        if ((v4 * (v4 >> ((993) & 63))) != (v4 ^ 42)) {
                                                                            var v46 int = 32767
                                                                            if false {
                                                                                v46 -= -449
                                                                                v47 := GenS6{F0: (v32 + v32), F1: false, F2: v2}
                                                                                _ = v47
                                                                            }
                                                                            if len(v36) > 0 {
                                                                                v36[0] = -2147483648
                                                                            }
                                                                            _ = v46
                                                                        } else {
                                                                            delete(v1, ("foo bar" + "foo bar"))
                                                                        }
                                                                    default:
                                                                        var v48 float64 = v5
                                                                        _ = v48
                                                                    }
                                                                    if false {
                                                                        switch fi27 {
                                                                        case -859:
                                                                            v3 -= 127
                                                                        case 2147483647:
                                                                            var v49 uint8 = 42
                                                                            _ = v49
                                                                        }
                                                                    }
                                                                    if len(v36) >= 2 {
                                                                        v36 = v36[1:len(v36)]
                                                                    }
                                                                }
                                                            } else {
                                                                var v50 int32 = -2147483648
                                                                _ = v50
                                                            }
                                                        case -2147483648:
                                                            fmt.Println(v3, v5, fi31)
                                                        case 0:
                                                            v32 -= v32
                                                        }
                                                        v51 := GenS7{F0: -2147483648, F1: 72}
                                                        var v52 int32 = 1
                                                        if true {
                                                            v1["foo bar"] = v32
                                                        }
                                                        switch fi31 {
                                                        case 606:
                                                            if (0.1 > v25) {
                                                                delete(v1, "\n\t")
                                                                var v53 bool = true
                                                                _ = v53
                                                            } else {
                                                                v54 := GenS3{F0: 42, F1: (v39 / (((v39 * 1)) | 1))}
                                                                fmt.Println(v2, v25, v25)
                                                                fmt.Println(v3, fi27)
                                                                v55 := GenS1{F0: -160, F1: int(v4), F2: v2}
                                                                _ = v54
                                                                _ = v55
                                                            }
                                                        default:
                                                            v56 := GenS7{F0: int32(v2), F1: v4}
                                                            _ = v56
                                                        }
                                                        var v57 bool = false
                                                        v36 = append(v36, (fi27 & fi27))
                                                        _ = v51
                                                        _ = v52
                                                        _ = v57
                                                    } else {
                                                        v3 += v3
                                                        v58 := GenS5{F0: 255, F1: (v4 ^ v4), F2: 18446744073709551615}
                                                        _ = v58
                                                    }
                                                    v59 := GenS8{F0: -9223372036854775808, F1: fi27}
                                                    v4 -= uint(v39)
                                                    _ = v59
                                                }
                                            default:
                                                if true {
                                                    v60 := GenS9{F0: 9223372036854775807, F1: -32768, F2: v39}
                                                    _ = v60
                                                } else {
                                                    fmt.Println(v39, fi31, v39)
                                                    v3 -= v3
                                                }
                                            }
                                            if true {
                                                var v61 string = "foo bar"
                                                _ = v61
                                            } else {
                                                v3 -= (v3 >> ((v4) & 63))
                                            }
                                            _ = v39
                                            _ = v40
                                        }
                                        v25 = v25
                                        if len(v36) > 0 {
                                            v36[0] = fi27
                                        }
                                        if len(v36) >= 2 {
                                            v36 = v36[1:len(v36)]
                                        }
                                    } else {
                                        v62 := v1[""]
                                        v25 = 3.14159
                                        _ = v62
                                    }
                                    if ((("\n\t" + ("a" + "hello")) + ("a" + "\n\t")) > ((("" + ("\n\t" + "foo bar")) + ("foo bar" + ("hello" + "a"))) + "\n\t")) {
                                        if len(v36) >= 2 {
                                            v36 = v36[1:len(v36)]
                                        }
                                        var v63 int16 = (v2 - v2)
                                        v4 += 107
                                        _ = v63
                                    }
                                    _ = v36
                                }
                                _ = fi31
                                _ = v32
                            }
                            _ = fi27
                        }
                        for fi64 := 0; fi64 < 3; fi64++ {
                            v65 := GenS10{F0: (-(fi64)), F1: (!false), F2: 0}
                            _ = fi64
                            _ = v65
                        }
                    } else {
                        for fi66 := 0; fi66 < 14; fi66++ {
                            v67 := []int{fi66, fi66, fi66}
                            _ = fi66
                            _ = v67
                        }
                    }
                    for fi68 := 0; fi68 < 6; fi68++ {
                        v25 += (v25 / float32((v5 * 3.14159)))
                        v4 -= v4
                        v3 -= v3
                        _ = fi68
                    }
                    _ = v25
                } else {
                    v69 := GenS11{F0: 137, F1: v4, F2: "hello"}
                    fmt.Println(v2)
                    v70 := GenS1{F0: -821, F1: 42, F2: (v2 & 42)}
                    _ = v69
                    _ = v70
                }
            case "":
                if true {
                    for fi71 := 0; fi71 < 8; fi71++ {
                        v72 := GenS1{F0: 0, F1: fi71, F2: v2}
                        _ = fi71
                        _ = v72
                    }
                }
            case "hello":
                for fi73 := 0; fi73 < 19; fi73++ {
                    v74 := GenS10{F0: (fi73 | 127), F1: ((957 >= 161) && (42 == 1)), F2: -479}
                    v2 = v2
                    _ = fi73
                    _ = v74
                }
            }
        case "\n\t":
            v4 = 0
        case "":
            switch "hello" {
            case "hello":
                fmt.Println(v2, v2, v5)
            case "":
                switch ("" + "a") {
                case "\n\t":
                    var v75 int = -9223372036854775808
                    _ = v75
                default:
                    var v76 uint16 = 42
                    _ = v76
                }
            case "a":
                v77 := GenS6{F0: 698, F1: (false && (2147483647 != -1)), F2: v2}
                _ = v77
            default:
                v78 := []int{-2147483648, -644}
                _ = v78
            }
        default:
            if ((38 == uint16(v4)) && ("hello" != "hello")) {
                var v79 uint32 = 2
                var v80 uint64 = 42
                _ = v79
                _ = v80
            } else {
                for fi81 := 0; fi81 < 16; fi81++ {
                    if (2 <= 4294967295) {
                        for fi82 := 0; fi82 < 13; fi82++ {
                            v5 -= 1.5
                            v2 += (v2 * (v2 + (v2 ^ v2)))
                            var v83 bool = ("a" <= "\n\t")
                            _ = fi82
                            _ = v83
                        }
                        if (2147483647 != 1) {
                            fmt.Println(v2, v5, v3)
                            v84 := []int{fi81}
                            _ = v84
                        }
                    }
                    var v85 uint64 = 0
                    var v86 string = "hello"
                    _ = fi81
                    _ = v85
                    _ = v86
                }
                var v87 float64 = 1.5
                if (!true) {
                    v2 += (v2 % (((v2 & -212)) | 1))
                } else {
                    v88 := GenS10{F0: 9223372036854775807, F1: ((true && (-9223372036854775808 >= 1)) && ((42 > -1) && false)), F2: -1}
                    _ = v88
                }
                _ = v87
            }
        }
        v2 = v2
        if (2147483647 >= -859) {
            var v89 int = 4294967295
            switch 348 {
            case 32767:
                var v90 uint = v4
                _ = v90
            }
            v91 := GenS9{F0: 0, F1: v2, F2: 255}
            _ = v89
            _ = v91
        }
    }
    v2 = 32767
    for fi92 := 0; fi92 < 2; fi92++ {
        for fi93 := 0; fi93 < 14; fi93++ {
            if ("\n\t" != ("\n\t" + ("foo bar" + ((("\n\t" + (("" + "foo bar") + (("a" + "\n\t") + "a"))) + "foo bar") + "foo bar")))) {
                if false {
                    switch (-(fi92)) {
                    case -113:
                        if true {
                            v94 := GenS8{F0: -9223372036854775808, F1: fi93}
                            v95 := GenS4{F0: 128, F1: (v3 + -1)}
                            v5 += ((v5 / v5) / v5)
                            _ = v94
                            _ = v95
                        } else {
                            if true {
                                fmt.Println(v3)
                                v96 := v1[("foo bar" + "hello")]
                                fmt.Println(v96, fi93, fi92)
                                switch 1 {
                                case 9223372036854775807:
                                    switch "a" {
                                    case "a":
                                        v97 := []int{}
                                        _ = v97
                                    case "hello":
                                        v98 := GenS11{F0: 236, F1: v4, F2: ""}
                                        _ = v98
                                    }
                                case 2147483647:
                                    v99 := []int{int(v96), fi92, (fi93 >> ((v4) & 63))}
                                    _ = v99
                                }
                                _ = v96
                            }
                        }
                        v5 = 3.14159
                        v5 -= v5
                    case 32767:
                        var v100 uint = 4294967295
                        _ = v100
                    case 42:
                        if false {
                            v2 += v2
                            v101 := v1["hello"]
                            v102 := []int{fi92}
                            fmt.Println(fi92, v101, v2)
                            _ = v101
                            _ = v102
                        } else {
                            _ = len(v1)
                            switch 32768 {
                            case 4294967295:
                                switch ("foo bar" + "foo bar") {
                                case "\n\t":
                                    switch "\n\t" {
                                    case "":
                                        v5 = v5
                                    }
                                case "hello":
                                    _ = len(v1)
                                default:
                                    v103 := GenS12{F0: 1}
                                    var v104 int = fi92
                                    if true {
                                        if (false || false) {
                                            var v105 float64 = v5
                                            var v106 bool = true
                                            v2 -= (^(v2))
                                            var v107 uint16 = 42
                                            _ = v105
                                            _ = v106
                                            _ = v107
                                        } else {
                                            switch fi92 {
                                            case 127:
                                                v4 = v4
                                            case 1:
                                                v108 := []int{}
                                                _ = v108
                                            case 270:
                                                fmt.Println(v5, fi92)
                                            }
                                        }
                                    }
                                    _ = v103
                                    _ = v104
                                }
                            }
                            if true {
                                delete(v1, "a")
                                v109 := v1[""]
                                _ = v109
                            }
                            v2 = 1000
                        }
                    default:
                        if false {
                            var v110 uint8 = 42
                            switch fi92 {
                            case -372:
                                v111 := GenS1{F0: 42, F1: (fi92 << ((v4) & 63)), F2: v2}
                                _ = v111
                            }
                            _ = v110
                        } else {
                            fmt.Println(fi92)
                            switch fi92 {
                            case 1:
                                v112 := v1["\n\t"]
                                _ = v112
                            case 42:
                                switch (fi92 ^ (fi92 % ((fi92) | 1))) {
                                case 42:
                                    var v113 bool = (!false)
                                    _ = v113
                                case 714:
                                    if true {
                                        v114 := GenS9{F0: 0, F1: v2, F2: 255}
                                        v115 := GenS13{F0: 887, F1: 753, F2: fi93}
                                        v4 -= (v4 / (((^(v4))) | 1))
                                        _ = v114
                                        _ = v115
                                    } else {
                                        v4 = v4
                                    }
                                default:
                                    var v116 bool = ((1 < 2) && true)
                                    _ = v116
                                }
                            case -707:
                                if ((3.14159 - (-2.25 + (0.0 / 1.0))) > ((0.0 + 1e10) / (0.0 - float32(float64(0.1))))) {
                                    fmt.Println(v4)
                                } else {
                                    v117 := GenS1{F0: -1, F1: (fi93 - (fi92 & fi93)), F2: (v2 * (v2 % ((v2) | 1)))}
                                    switch (fi92 << ((1) & 63)) {
                                    case 128:
                                        fmt.Println(v2)
                                    case 690:
                                        var v118 float64 = v5
                                        _ = v118
                                    default:
                                        _ = len(v1)
                                    }
                                    v3 = v3
                                    if false {
                                        fmt.Println(v4, v4)
                                        var v119 bool = (true && ((v3 * (v3 - (v3 | v3))) == (v3 ^ v3)))
                                        if true {
                                            fmt.Println(v3, fi93, fi92)
                                            v2 += v2
                                        }
                                        _ = v119
                                    }
                                    if true {
                                        if ((false || (!true)) && false) {
                                            v3 = 127
                                            v120 := GenS14{F0: 0}
                                            _ = v120
                                        }
                                        var v121 uint32 = 42
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v122 int8 = v3
                                            _ = v122
                                        case "foo bar":
                                            if true {
                                                if ((1 > 0) || true) {
                                                    v5 += float64((1e10 - 1e10))
                                                    if false {
                                                        if (1.5 >= 3.14159) {
                                                            var v123 uint32 = 0
                                                            v124 := []int{fi93, 4294967295, 2147483647}
                                                            v5 = v5
                                                            _ = v123
                                                            _ = v124
                                                        }
                                                        if (float64(-2.25) <= v5) {
                                                            v4 -= 949
                                                        } else {
                                                            var v125 int64 = -857
                                                            var v126 bool = ((true || true) || true)
                                                            var v127 int32 = 430
                                                            _ = v125
                                                            _ = v126
                                                            _ = v127
                                                        }
                                                        v2 = 0
                                                        delete(v1, (("foo bar" + "") + ""))
                                                    } else {
                                                        if false {
                                                            fmt.Println(v5, v2)
                                                            if (1 <= 0) {
                                                                v128 := GenS8{F0: 0, F1: (fi92 * -1)}
                                                                _ = v128
                                                            }
                                                            fmt.Println(v3, fi93, v3)
                                                        } else {
                                                            v5 -= v5
                                                        }
                                                        if (!false) {
                                                            var v129 uint32 = 2
                                                            switch (^(fi93)) {
                                                            case 127:
                                                                var v130 int32 = 2147483647
                                                                _ = v130
                                                            case 4294967295:
                                                                if false {
                                                                    if ((^(v4)) < (v4 >> ((v4) & 63))) {
                                                                        v131 := []int{0}
                                                                        _ = v131
                                                                    }
                                                                    v132 := GenS15{F0: -371}
                                                                    v133 := GenS6{F0: v121, F1: (!((v4 & v4) < v4)), F2: v2}
                                                                    _ = v132
                                                                    _ = v133
                                                                } else {
                                                                    var v134 uint32 = v129
                                                                    v135 := GenS15{F0: -1}
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        v136 := GenS16{F0: (v3 & (v3 | v3)), F1: (v4 * v4)}
                                                                        _ = v136
                                                                    default:
                                                                        var v137 uint32 = 546
                                                                        _ = v137
                                                                    }
                                                                    _ = v134
                                                                    _ = v135
                                                                }
                                                            case -478:
                                                                fmt.Println(v121, v5, fi92)
                                                            }
                                                            var v138 int32 = -2147483648
                                                            _ = v129
                                                            _ = v138
                                                        } else {
                                                            v139 := GenS17{F0: 1, F1: v4, F2: "foo bar"}
                                                            _ = v139
                                                        }
                                                        var v140 int64 = -1
                                                        _ = v140
                                                    }
                                                    v141 := GenS18{F0: 42}
                                                    var v142 uint8 = 2
                                                    var v143 int8 = (v3 << ((v4) & 63))
                                                    _ = v141
                                                    _ = v142
                                                    _ = v143
                                                } else {
                                                    if (true || false) {
                                                        v3 = (v3 - (v3 | v3))
                                                    } else {
                                                        v144 := GenS7{F0: 42, F1: v4}
                                                        var v145 float32 = 1.5
                                                        _ = v144
                                                        _ = v145
                                                    }
                                                    v146 := GenS6{F0: (^(v121)), F1: true, F2: v2}
                                                    var v147 int16 = v2
                                                    _ = v146
                                                    _ = v147
                                                }
                                                var v148 bool = true
                                                _ = v148
                                            } else {
                                                v5 += (v5 + ((((-2.25 * 3.14159) * 3.14159) + v5) - v5))
                                            }
                                        }
                                        _ = v121
                                    } else {
                                        switch (fi93 % ((-395) | 1)) {
                                        case 1:
                                            fmt.Println(v5)
                                        default:
                                            fmt.Println(v4, v2, v3)
                                        }
                                        v4 = 0
                                    }
                                    _ = v117
                                }
                            }
                            v149 := GenS19{F0: 0, F1: 732, F2: true}
                            _ = v149
                        }
                    }
                    v4 += v4
                } else {
                    var v150 uint16 = 0
                    v151 := []int{128, (fi92 ^ 32768)}
                    _ = v150
                    _ = v151
                }
            }
            fmt.Println(fi93)
            _ = fi93
        }
        _ = fi92
    }
    switch ("a" + "") {
    case "foo bar":
        v152 := GenS5{F0: 65535, F1: 834, F2: 42}
        _ = v152
    case "":
        v153 := v1[("foo bar" + "\n\t")]
        _ = v153
    case "a":
        var v154 int16 = v2
        _ = v154
    }
    for fi155 := 0; fi155 < 5; fi155++ {
        var v156 uint32 = 269
        _ = fi155
        _ = v156
    }
    for fi157 := 0; fi157 < 17; fi157++ {
        switch 555 {
        case -1:
            for fi158 := 0; fi158 < 1; fi158++ {
                if ((true && false) || ((true && true) && (4294967295 <= 4294967295))) {
                    v159 := GenS18{F0: 242}
                    v5 = (v5 - (v5 - 1.5))
                    v2 = 0
                    _ = v159
                }
                v4 = (v4 % ((v4) | 1))
                _ = fi158
            }
        case 4294967295:
            for fi160 := 0; fi160 < 12; fi160++ {
                v161 := GenS20{F0: (4294967295 > 1), F1: (("hello" + "hello") + "hello"), F2: 0}
                _ = fi160
                _ = v161
            }
        case 2147483647:
            v3 -= v3
        }
        switch fi157 {
        case -1:
            var v162 uint = v4
            _ = v162
        case 32768:
            var v163 bool = true
            _ = v163
        case 0:
            var v164 float32 = 1.5
            _ = v164
        }
        _ = fi157
    }
    for fi165 := 0; fi165 < 6; fi165++ {
        for fi166 := 0; fi166 < 9; fi166++ {
            var v167 int32 = 1000
            _ = fi166
            _ = v167
        }
        _ = fi165
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
