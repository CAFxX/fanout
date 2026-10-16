package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int64 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string; F1 uint16; F2 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int64; F1 string; F2 string }
type GenS8 struct { F0 uint64; F1 bool; F2 int8 }
type GenS9 struct { F0 int16; F1 uint64 }
type GenS10 struct { F0 uint64; F1 int8; F2 int8 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int; F1 string }

func main() {
    v1 := make(map[int]string)
    var v2 float64 = 1e10
    var v3 uint8 = 42
    var v4 uint = 871
    var v5 float32 = (0.0 / 3.14159)
    for fi6 := 0; fi6 < 17; fi6++ {
        for fi7 := 0; fi7 < 4; fi7++ {
            if false {
                v3 -= (v3 * (v3 * (v3 << ((v4) & 63))))
            }
            if true {
                v4 -= (v4 + (v4 * v4))
                if (!(v3 >= (^(v3)))) {
                    v3 += v3
                }
                if true {
                    if ((948 >= -9223372036854775808) || (!false)) {
                        v4 = v4
                        v2 = v2
                    }
                    if (v5 != (v5 - v5)) {
                        v1[fi7] = "hello"
                        v2 = 0.0
                        var v8 float32 = v5
                        v5 = (((v8 * v5) * (1.5 * v8)) + v8)
                        fmt.Println(v2, v3, v4)
                        _ = v8
                    }
                    v9 := []int{}
                    _ = v9
                } else {
                    v5 += (v5 * v5)
                    v2 = v2
                    switch "hello" {
                    case "\n\t":
                        if (false && (false || true)) {
                            v10 := GenS1{F0: 42, F1: (fi7 ^ fi7), F2: -1}
                            switch "hello" {
                            case "":
                                switch "foo bar" {
                                case "\n\t":
                                    v11 := GenS1{F0: 100, F1: fi7, F2: 647}
                                    _ = v11
                                case "hello":
                                    fmt.Println(v3, v5)
                                case "foo bar":
                                    var v12 uint64 = 18446744073709551615
                                    _ = v12
                                default:
                                    switch ((("hello" + ("" + ("a" + "foo bar"))) + "hello") + "foo bar") {
                                    case "a":
                                        var v13 float32 = (-2.25 / float32(v2))
                                        _ = v13
                                    default:
                                        v14 := v1[32767]
                                        _ = v14
                                    }
                                }
                            case "hello":
                                v5 = v5
                            }
                            _ = v10
                        }
                    }
                    v4 = 18446744073709551615
                }
                var v15 float64 = v2
                _ = v15
            } else {
                if (((!((int8(v3) > 121) && ((2 < 2) || false))) && false) && (((!true) && true) && true)) {
                    if (1 == 1) {
                        v16 := []int{}
                        _ = v16
                    } else {
                        v3 += (v3 << ((382) & 63))
                        if false {
                            switch "\n\t" {
                            case "a":
                                var v17 uint8 = (v3 >> (((v4 % (((v4 * v4)) | 1))) & 63))
                                _ = v17
                            case "\n\t":
                                var v18 string = "\n\t"
                                _ = v18
                            default:
                                switch "\n\t" {
                                case "":
                                    v19 := []int{}
                                    _ = v19
                                default:
                                    v20 := []int{(fi7 | fi6), fi6, (fi7 - (fi6 << ((42) & 63))), (^(fi7))}
                                    _ = v20
                                }
                            }
                            var v21 uint16 = 0
                            var v22 float32 = float32(v2)
                            var v23 int8 = 80
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                        fmt.Println(fi6, v3, v4)
                    }
                    fmt.Println(v5, fi7)
                }
                if false {
                    v2 -= v2
                    delete(v1, fi6)
                }
            }
            _ = fi7
        }
        v24 := []int{(fi6 % (((fi6 * 1)) | 1)), fi6, (fi6 ^ (fi6 | 2147483647))}
        _ = fi6
        _ = v24
    }
    fmt.Println(v2, v3)
    switch -350 {
    case 42:
        switch "" {
        case "foo bar":
            var v25 uint64 = 300
            _ = v25
        case "":
            switch "foo bar" {
            case "a":
                v26 := GenS2{F0: 0, F1: -32768}
                _ = v26
            case "hello":
                v27 := GenS1{F0: 0, F1: -994, F2: 0}
                _ = v27
            case "":
                for fi28 := 0; fi28 < 10; fi28++ {
                    v29 := GenS1{F0: -1, F1: (fi28 - fi28), F2: 1}
                    _ = fi28
                    _ = v29
                }
            }
        case "a":
            var v30 uint16 = 42
            _ = v30
        }
    case 0:
        switch ("foo bar" + (("foo bar" + (("hello" + "") + "a")) + "hello")) {
        case "a":
            for fi31 := 0; fi31 < 8; fi31++ {
                var v32 uint16 = 42
                if true {
                    var v33 int64 = -9223372036854775808
                    v34 := GenS3{F0: 1}
                    if (!true) {
                        _ = len(v1)
                    } else {
                        v35 := GenS2{F0: v33, F1: 1000}
                        v36 := GenS3{F0: 0}
                        _ = v35
                        _ = v36
                    }
                    _ = v33
                    _ = v34
                }
                for fi37 := 0; fi37 < 9; fi37++ {
                    var v38 float32 = v5
                    v38 += v5
                    _ = fi37
                    _ = v38
                }
                _ = fi31
                _ = v32
            }
        }
    }
    var v39 uint32 = 467
    var v40 int16 = -32768
    switch 2147483647 {
    case 32768:
        for fi41 := 0; fi41 < 9; fi41++ {
            v2 -= (v2 + (((v2 + v2) + (float64(v5) + 1e10)) - (3.14159 - ((((((0.1 / v2) / ((1.5 - 0.1) / (v2 * float64(v5)))) / 0.1) + float64((v5 * 0.1))) / ((v2 / 1.0) + v2)) / v2))))
            fmt.Println(v4)
            v5 = v5
            _ = fi41
        }
    case 1:
        if (127 < 4294967295) {
            for fi42 := 0; fi42 < 1; fi42++ {
                switch fi42 {
                case 767:
                    fmt.Println(fi42, v3)
                case 4294967295:
                    v5 = (3.14159 / 1.0)
                case 32767:
                    v39 += v39
                }
                _ = fi42
            }
            v3 = v3
            fmt.Println(v2, v40)
        } else {
            var v43 uint32 = uint32(v4)
            if (766 == 72) {
                if true {
                    fmt.Println(v43, v39)
                }
                v44 := []int{int(v3), -482}
                v45 := v1[32768]
                _ = v44
                _ = v45
            } else {
                if (true && false) {
                    var v46 float32 = (((v5 - v5) - v5) * v5)
                    _ = v46
                } else {
                    v47 := GenS3{F0: 42}
                    _ = v47
                }
            }
            v3 -= 0
            _ = v43
        }
    default:
        for fi48 := 0; fi48 < 12; fi48++ {
            var v49 float32 = (v5 * v5)
            for fi50 := 0; fi50 < 2; fi50++ {
                switch 128 {
                case 2147483647:
                    var v51 float32 = v5
                    _ = v51
                default:
                    v1[42] = "\n\t"
                }
                fmt.Println(fi50, v49, v39)
                if true {
                    v52 := GenS4{F0: "hello", F1: 2, F2: (fi48 / ((fi50) | 1))}
                    if (true && false) {
                        if (false || ((^(v39)) >= v39)) {
                            v53 := GenS5{F0: 236}
                            _ = v53
                        }
                        if (v3 != 210) {
                            v5 = (v5 + v5)
                            v54 := GenS3{F0: -659}
                            if true {
                                if false {
                                    if (127 < 100) {
                                        v55 := v1[(fi48 >> ((v4) & 63))]
                                        _ = v55
                                    }
                                    fmt.Println(v39, v4, v5)
                                    fmt.Println(v39, v39, v3)
                                    v49 += v49
                                    if false {
                                        if false {
                                            if ("" > ("\n\t" + "foo bar")) {
                                                var v56 uint32 = (v39 << (((v4 + 0)) & 63))
                                                fmt.Println(v39, v2)
                                                fmt.Println(v5, fi50, v40)
                                                _ = v56
                                            } else {
                                                if true {
                                                    var v57 uint8 = (v3 ^ uint8(v4))
                                                    fmt.Println(v4, v49)
                                                    v58 := []int{fi48, (fi48 + (^(fi48))), fi50, fi50}
                                                    _ = v57
                                                    _ = v58
                                                } else {
                                                    if true {
                                                        if false {
                                                            v59 := GenS4{F0: "hello", F1: 961, F2: 128}
                                                            var v60 bool = ("" == "\n\t")
                                                            _ = v59
                                                            _ = v60
                                                        } else {
                                                            fmt.Println(v2, fi48, v40)
                                                        }
                                                        var v61 int16 = v40
                                                        switch ("" + ("foo bar" + "a")) {
                                                        case "hello":
                                                            v2 -= (v2 + ((v2 + v2) / (v2 / (float64((float32(1e10) + 1.5)) / (((v2 + (v2 * v2)) + v2) / v2)))))
                                                        default:
                                                            v2 += (v2 + (v2 / (0.1 + v2)))
                                                        }
                                                        v3 = v3
                                                        switch 39 {
                                                        case -1:
                                                            if (false || (1e10 >= (v5 * v5))) {
                                                                v40 = v61
                                                                var v62 int32 = 1
                                                                _ = v62
                                                            }
                                                        default:
                                                            if false {
                                                                v63 := []int{}
                                                                v64 := GenS6{F0: ("hello" + "a")}
                                                                v3 = (v3 - 0)
                                                                v65 := GenS6{F0: ("\n\t" + "a")}
                                                                switch ((("a" + (("hello" + "") + "")) + (("hello" + ("\n\t" + "")) + ("foo bar" + ("hello" + (("\n\t" + "") + "a"))))) + ("\n\t" + "hello")) {
                                                                case "":
                                                                    v66 := GenS7{F0: 0, F1: "\n\t", F2: "a"}
                                                                    _ = v66
                                                                case "hello":
                                                                    fmt.Println(fi48)
                                                                case "a":
                                                                    var v67 bool = ((fi50 != fi48) && ((590 != 42) && true))
                                                                    _ = v67
                                                                }
                                                                _ = v63
                                                                _ = v64
                                                                _ = v65
                                                            } else {
                                                                v68 := []int{127, 4294967295, fi48, (fi48 ^ -1)}
                                                                var v69 int32 = 1
                                                                v1[-94] = ("hello" + "\n\t")
                                                                if ((!(!(375 < 32767))) || (818 <= uint64(v39))) {
                                                                    v40 += (v61 - v61)
                                                                    var v70 int32 = v69
                                                                    v71 := v1[(^(fi50))]
                                                                    delete(v1, fi48)
                                                                    switch (fi50 + int(v69)) {
                                                                    case -828:
                                                                        v68 = append(v68, 128)
                                                                    case -431:
                                                                        var v72 uint64 = 2
                                                                        _ = v72
                                                                    case 24:
                                                                        if len(v68) >= 2 {
                                                                            v68 = v68[1:len(v68)]
                                                                        }
                                                                        var v73 uint64 = 42
                                                                        if (!false) {
                                                                            fmt.Println(v2)
                                                                        } else {
                                                                            switch ((v71 + v71) + (v71 + (v71 + (v71 + "")))) {
                                                                            case "foo bar":
                                                                                var v74 uint = v4
                                                                                _ = v74
                                                                            case "a":
                                                                                v73 += (v73 & (v73 * v73))
                                                                            default:
                                                                                v73 = (v73 & (v73 % (((v73 & 18446744073709551615)) | 1)))
                                                                            }
                                                                            if false {
                                                                                var v75 uint16 = 65535
                                                                                var v76 uint = (v4 * v4)
                                                                                _ = v75
                                                                                _ = v76
                                                                            } else {
                                                                                v73 = uint64(v4)
                                                                                var v77 uint8 = 173
                                                                                _ = v77
                                                                            }
                                                                        }
                                                                        _ = v73
                                                                    }
                                                                    _ = v70
                                                                    _ = v71
                                                                } else {
                                                                    fmt.Println(fi48)
                                                                    var v78 int16 = v61
                                                                    v40 -= (v40 & int16(v4))
                                                                    _ = v78
                                                                }
                                                                _ = v68
                                                                _ = v69
                                                            }
                                                        }
                                                        _ = v61
                                                    }
                                                    v2 -= 0.1
                                                }
                                            }
                                        }
                                        if true {
                                            if false {
                                                var v79 uint = 171
                                                _ = v79
                                            } else {
                                                fmt.Println(v4, fi48, fi50)
                                                v4 = (v4 | (v4 % ((970) | 1)))
                                                switch "a" {
                                                case "hello":
                                                    var v80 uint8 = (v3 & 127)
                                                    v5 += ((((v5 + -2.25) * v5) + ((v5 - -2.25) - (3.14159 / -2.25))) * (v5 * (((-2.25 - v5) - (float32(0.1) / 0.1)) * (v5 / (v49 + ((v5 - ((v5 / ((((1.5 + (0.1 * 1e10)) / v5) / 1e10) - v5)) / v5)) * v49))))))
                                                    _ = v80
                                                case "foo bar":
                                                    v81 := GenS3{F0: 0}
                                                    v82 := []int{(fi48 >> ((v4) & 63))}
                                                    var v83 int = 9223372036854775807
                                                    _ = v81
                                                    _ = v82
                                                    _ = v83
                                                case "":
                                                    var v84 string = "a"
                                                    _ = v84
                                                }
                                            }
                                        } else {
                                            var v85 float64 = 0.0
                                            var v86 uint = v4
                                            _ = v85
                                            _ = v86
                                        }
                                    }
                                } else {
                                    _ = len(v1)
                                    fmt.Println(v39)
                                }
                                if ((42 < 2147483647) && false) {
                                    switch 122 {
                                    case 9223372036854775807:
                                        v87 := v1[(-(fi48))]
                                        switch "" {
                                        case "":
                                            var v88 string = "a"
                                            _ = v88
                                        case "\n\t":
                                            switch 1 {
                                            case -593:
                                                fmt.Println(fi50, v39, v3)
                                            case -480:
                                                var v89 uint64 = 2
                                                _ = v89
                                            }
                                        default:
                                            if (!true) {
                                                v90 := GenS8{F0: 0, F1: (137 < 42), F2: 87}
                                                v49 = v5
                                                v91 := GenS9{F0: -1, F1: 0}
                                                _ = v90
                                                _ = v91
                                            }
                                        }
                                        _ = v87
                                    case 127:
                                        var v92 uint16 = 537
                                        _ = v92
                                    default:
                                        switch (fi50 - (fi48 / ((565) | 1))) {
                                        case -628:
                                            switch "foo bar" {
                                            case "\n\t":
                                                v1[127] = "a"
                                            case "":
                                                v49 -= ((1e10 * (((v49 * -2.25) * v49) - v49)) * v5)
                                            case "foo bar":
                                                v93 := GenS10{F0: 2, F1: -75, F2: int8(v40)}
                                                _ = v93
                                            default:
                                                v4 = (v4 >> (((v4 ^ v4)) & 63))
                                            }
                                        case 697:
                                            v39 += v39
                                        }
                                    }
                                    var v94 uint64 = 154
                                    fmt.Println(v2, v3, v94)
                                    _ = v94
                                } else {
                                    v95 := []int{-576}
                                    _ = v95
                                }
                            } else {
                                v49 = v5
                            }
                            _ = v54
                        } else {
                            if (!true) {
                                if (2 <= 18446744073709551615) {
                                    v40 = v40
                                    if true {
                                        v96 := GenS7{F0: int64(v3), F1: "", F2: "hello"}
                                        var v97 int32 = -2147483648
                                        _ = v96
                                        _ = v97
                                    } else {
                                        switch (^(fi50)) {
                                        case 4294967295:
                                            var v98 float64 = 3.14159
                                            if ((v40 | (v40 >> ((255) & 63))) < int16(fi50)) {
                                                var v99 int32 = 2147483647
                                                v100 := GenS4{F0: ("foo bar" + "foo bar"), F1: 226, F2: (fi50 % ((fi48) | 1))}
                                                _ = v100
                                                _ = v99
                                            } else {
                                                switch fi48 {
                                                case 9223372036854775807:
                                                    fmt.Println(fi48, v5)
                                                case 534:
                                                    var v101 uint = v4
                                                    _ = v101
                                                default:
                                                    var v102 uint32 = v39
                                                    _ = v102
                                                }
                                            }
                                            _ = v98
                                        case 9223372036854775807:
                                            v103 := GenS5{F0: 0}
                                            _ = v103
                                        default:
                                            var v104 int8 = -49
                                            _ = v104
                                        }
                                    }
                                } else {
                                    switch fi50 {
                                    case 128:
                                        if true {
                                            var v105 int8 = 1
                                            v5 += v5
                                            v2 -= v2
                                            _ = v105
                                        } else {
                                            v106 := GenS11{F0: uint32(v3)}
                                            _ = v106
                                        }
                                    case 42:
                                        var v107 float32 = 0.1
                                        _ = v107
                                    }
                                    v108 := GenS11{F0: 4294967295}
                                    delete(v1, fi50)
                                    _ = v108
                                }
                                var v109 float32 = (float32(-2.25) - -2.25)
                                if false {
                                    switch "foo bar" {
                                    case "hello":
                                        fmt.Println(v39)
                                    }
                                }
                                _ = v109
                            }
                            v110 := GenS3{F0: -11}
                            switch ("foo bar" + ("" + ("\n\t" + (("hello" + "foo bar") + (("a" + "") + ("\n\t" + "foo bar")))))) {
                            case "":
                                if (v4 >= v4) {
                                    if (!false) {
                                        v111 := GenS7{F0: -1, F1: (("a" + "\n\t") + ("hello" + "a")), F2: "foo bar"}
                                        _ = v111
                                    } else {
                                        switch (-(fi50)) {
                                        case 2147483647:
                                            v49 -= 3.14159
                                        case 0:
                                            v112 := GenS10{F0: 42, F1: 102, F2: 127}
                                            _ = v112
                                        }
                                        v2 = (v2 + v2)
                                        fmt.Println(fi48, fi48)
                                        var v113 uint = 2
                                        _ = v113
                                    }
                                } else {
                                    v114 := []int{}
                                    if len(v114) > 0 {
                                        v114[0] = (fi48 | (fi50 >> ((uint(v3)) & 63)))
                                    }
                                    var v115 float64 = v2
                                    _ = v114
                                    _ = v115
                                }
                            case "hello":
                                v116 := GenS8{F0: 18446744073709551615, F1: false, F2: -128}
                                _ = v116
                            }
                            _ = v110
                        }
                    } else {
                        fmt.Println(v49, fi50)
                        fmt.Println(fi48)
                    }
                    v39 -= v39
                    _ = v52
                }
                _ = fi50
            }
            var v117 uint32 = v39
            _ = fi48
            _ = v117
            _ = v49
        }
    }
    var v118 bool = ((!(743 < -1)) || ("a" < ((("hello" + ("" + "hello")) + ("\n\t" + "")) + "")))
    var v119 bool = v118
    var v120 int16 = (v40 * 1000)
    for fi121 := 0; fi121 < 5; fi121++ {
        switch "hello" {
        case "a":
            if v118 {
                if (!((fi121 - (fi121 % ((9223372036854775807) | 1))) <= fi121)) {
                    switch "foo bar" {
                    case "\n\t":
                        for fi122 := 0; fi122 < 18; fi122++ {
                            var v123 int32 = 42
                            fmt.Println(v2)
                            _ = fi122
                            _ = v123
                        }
                    case "hello":
                        if v119 {
                            if true {
                                for fi124 := 0; fi124 < 20; fi124++ {
                                    var v125 float64 = ((((v2 - (v2 * (((float64(-2.25) + v2) * 3.14159) + v2))) / (v2 / float64(1e10))) * float64(1.5)) / v2)
                                    _ = fi124
                                    _ = v125
                                }
                                v4 += (v4 | (^(v4)))
                            }
                            v126 := GenS5{F0: 0}
                            if v118 {
                                var v127 int = 32767
                                _ = v127
                            } else {
                                for fi128 := 0; fi128 < 8; fi128++ {
                                    var v129 uint32 = (v39 ^ (v39 & 263))
                                    _ = fi128
                                    _ = v129
                                }
                                v39 -= v39
                            }
                            _ = v126
                        }
                    default:
                        v2 = v2
                    }
                    v119 = (false || v119)
                } else {
                    v130 := GenS9{F0: v40, F1: 120}
                    for fi131 := 0; fi131 < 9; fi131++ {
                        var v132 float32 = (1e10 - (((v5 * ((v5 + ((v5 + ((float32(float64((v5 - 3.14159))) * v5) - v5)) * (1.5 - v5))) / (1e10 - v5))) + 0.1) / (v5 * v5)))
                        _ = len(v1)
                        _ = fi131
                        _ = v132
                    }
                    for fi133 := 0; fi133 < 16; fi133++ {
                        switch (fi121 + fi133) {
                        case 32768:
                            v134 := []int{fi121, (fi133 << (((v4 - v4)) & 63)), 42, (fi133 % (((fi121 & -175)) | 1))}
                            _ = v134
                        }
                        if v119 {
                            v1[fi121] = "\n\t"
                            v120 = v40
                            if true {
                                var v135 uint16 = 65535
                                _ = v135
                            }
                        } else {
                            v39 -= 2
                            switch fi133 {
                            case 42:
                                v136 := GenS12{F0: fi121, F1: ((("hello" + ("\n\t" + ((("a" + "hello") + "") + "a"))) + "") + "")}
                                _ = v136
                            case 9223372036854775807:
                                v2 = v2
                            case 32767:
                                v5 = v5
                            }
                        }
                        var v137 int8 = -128
                        v138 := GenS1{F0: v137, F1: fi133, F2: -577}
                        _ = fi133
                        _ = v137
                        _ = v138
                    }
                    _ = v130
                }
                var v139 bool = v118
                fmt.Println(v2)
                _ = v139
            }
        case "\n\t":
            v40 += (v120 / (((v120 / (((v40 / ((v120) | 1))) | 1))) | 1))
        default:
            for fi140 := 0; fi140 < 9; fi140++ {
                v3 = (v3 + v3)
                v4 += v4
                v118 = (!(v3 == v3))
                var v141 int32 = -2147483648
                var v142 float64 = 3.14159
                _ = fi140
                _ = v141
                _ = v142
            }
        }
        _ = fi121
    }
    v5 = v5
    switch 2147483647 {
    case 32767:
        v2 += v2
    case 4294967295:
        fmt.Println(v39, v5, v3)
    default:
        for fi143 := 0; fi143 < 11; fi143++ {
            var v144 int8 = 125
            var v145 int64 = 42
            _ = fi143
            _ = v144
            _ = v145
        }
    }
    if (-32768 == 393) {
        for fi146 := 0; fi146 < 1; fi146++ {
            var v147 string = ("hello" + "")
            v148 := []int{32767, 32767}
            _ = fi146
            _ = v147
            _ = v148
        }
        v119 = false
    } else {
        v120 -= (v120 + v40)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v118:", v118)
    fmt.Println("v119:", v119)
    fmt.Println("v120:", v120)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v118
    _ = v119
    _ = v120
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
