package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 bool; F1 int64; F2 bool }
type GenS6 struct { F0 int32; F1 int64; F2 uint32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int16; F1 int32 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint64; F1 uint; F2 int8 }
type GenS12 struct { F0 uint32; F1 string; F2 uint64 }

func main() {
    var v1 int = 4294967295
    var v2 int = v1
    var v3 bool = (579 >= 126)
    var v4 uint32 = 42
    v4 -= v4
    var v5 uint8 = 128
    for fi6 := 0; fi6 < 19; fi6++ {
        for fi7 := 0; fi7 < 12; fi7++ {
            var v8 int8 = 127
            v9 := GenS1{F0: 1000}
            if (65535 > 18446744073709551615) {
                v10 := []int{(-(fi7)), (fi7 % ((fi6) | 1))}
                _ = v10
            }
            var v11 uint8 = v5
            var v12 string = ""
            _ = fi7
            _ = v11
            _ = v12
            _ = v8
            _ = v9
        }
        v13 := []int{32767, v1}
        switch "" {
        case "\n\t":
            for fi14 := 0; fi14 < 17; fi14++ {
                if v3 {
                    switch ((((("\n\t" + "hello") + (((("" + ((("a" + "foo bar") + (("foo bar" + "foo bar") + "")) + "")) + "hello") + ("foo bar" + "hello")) + "")) + (("hello" + (("foo bar" + "hello") + "\n\t")) + "")) + "") + "") {
                    case "":
                        var v15 int = (^(fi14))
                        _ = v15
                    default:
                        switch 127 {
                        case 128:
                            v13 = append(v13, fi14)
                        case 32767:
                            if len(v13) > 0 {
                                v13[0] = v1
                            }
                        case 478:
                            switch "foo bar" {
                            case "\n\t":
                                var v16 int64 = -1
                                _ = v16
                            case "foo bar":
                                var v17 int16 = 1000
                                _ = v17
                            }
                        }
                    }
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                }
                switch ("hello" + "foo bar") {
                case "":
                    v1 = fi14
                case "foo bar":
                    switch v2 {
                    case 0:
                        v18 := GenS1{F0: -103}
                        _ = v18
                    case 1:
                        var v19 float32 = 1.5
                        _ = v19
                    case 128:
                        v3 = v3
                    }
                    v13 = append(v13, 128)
                case "\n\t":
                    v13 = append(v13, 232)
                default:
                    _ = len(v13)
                }
                _ = fi14
            }
        default:
            if (1 < 635) {
                if (32768 != 32767) {
                    var v20 int8 = -12
                    var v21 float32 = 1.5
                    switch "hello" {
                    case "\n\t":
                        if (v3 || v3) {
                            if v3 {
                                for fi22 := 0; fi22 < 17; fi22++ {
                                    var v23 int32 = 1
                                    _ = fi22
                                    _ = v23
                                }
                                v1 -= fi6
                            }
                            switch v1 {
                            case -49:
                                v1 = fi6
                            default:
                                var v24 float32 = v21
                                _ = v24
                            }
                        } else {
                            if false {
                                v25 := GenS1{F0: 903}
                                _ = v25
                            }
                        }
                    case "foo bar":
                        var v26 uint32 = (v4 >> ((255) & 63))
                        _ = v26
                    case "a":
                        var v27 uint = 4294967295
                        _ = v27
                    }
                    _ = v20
                    _ = v21
                }
            } else {
                var v28 float32 = 0.0
                _ = v28
            }
        }
        _ = fi6
        _ = v13
    }
    for fi29 := 0; fi29 < 9; fi29++ {
        if (!((float32(0.1) * 1e10) > (float32(3.14159) / 1e10))) {
            switch (v2 ^ fi29) {
            case 0:
                switch fi29 {
                case 1:
                    v30 := GenS2{F0: 413, F1: -63}
                    _ = v30
                case -2147483648:
                    if (!v3) {
                        v4 = v4
                    } else {
                        var v31 bool = (127 >= -1)
                        for fi32 := 0; fi32 < 2; fi32++ {
                            var v33 uint32 = v4
                            switch "hello" {
                            case "hello":
                                var v34 float64 = 0.0
                                _ = v34
                            default:
                                v35 := GenS1{F0: -662}
                                _ = v35
                            }
                            switch "foo bar" {
                            case "\n\t":
                                if (v3 && true) {
                                    v36 := GenS3{F0: 145, F1: -9223372036854775808}
                                    v1 = (fi29 / ((-9223372036854775808) | 1))
                                    _ = v36
                                } else {
                                    v37 := []int{int(v5)}
                                    v2 = fi32
                                    _ = v37
                                }
                            case "foo bar":
                                v38 := GenS2{F0: 255, F1: 49}
                                if v3 {
                                    switch (v1 + (v1 >> ((854) & 63))) {
                                    case 127:
                                        fmt.Println(v5, fi29, v3)
                                    }
                                } else {
                                    v39 := []int{9223372036854775807, -9223372036854775808}
                                    _ = v39
                                }
                                fmt.Println(v1)
                                _ = v38
                            default:
                                var v40 int64 = -974
                                _ = v40
                            }
                            _ = fi32
                            _ = v33
                        }
                        if (1 >= 0) {
                            v5 = (v5 - uint8(v4))
                            if v3 {
                                v2 -= v1
                                fmt.Println(v5, v5, v31)
                            } else {
                                fmt.Println(v4, v31, v3)
                                v4 = 2
                                switch ("foo bar" + "\n\t") {
                                case "":
                                    if v3 {
                                        v5 += v5
                                        var v41 int = (v2 ^ fi29)
                                        _ = v41
                                    } else {
                                        var v42 int64 = -9223372036854775808
                                        v43 := GenS1{F0: -32768}
                                        if v3 {
                                            var v44 int16 = 42
                                            var v45 int64 = v42
                                            for fi46 := 0; fi46 < 8; fi46++ {
                                                switch ("hello" + ("hello" + "")) {
                                                case "foo bar":
                                                    fmt.Println(v2, fi29, fi46)
                                                case "hello":
                                                    var v47 uint32 = v4
                                                    _ = v47
                                                case "a":
                                                    v45 += v45
                                                }
                                                fmt.Println(v45, v2)
                                                _ = fi46
                                            }
                                            _ = v44
                                            _ = v45
                                        } else {
                                            var v48 uint64 = 1
                                            v49 := GenS2{F0: 2, F1: 42}
                                            v31 = (65535 > 961)
                                            _ = v48
                                            _ = v49
                                        }
                                        _ = v42
                                        _ = v43
                                    }
                                case "foo bar":
                                    fmt.Println(v2, v4, v3)
                                default:
                                    v50 := GenS1{F0: -246}
                                    _ = v50
                                }
                            }
                        } else {
                            var v51 string = (("hello" + "\n\t") + "hello")
                            var v52 int8 = 100
                            for fi53 := 0; fi53 < 17; fi53++ {
                                v54 := []int{}
                                _ = fi53
                                _ = v54
                            }
                            _ = v51
                            _ = v52
                        }
                        _ = v31
                    }
                case 32767:
                    if (255 > 58) {
                        fmt.Println(v1)
                        v55 := GenS1{F0: 0}
                        _ = v55
                    } else {
                        switch (((((("hello" + "hello") + "a") + "a") + "\n\t") + ((("hello" + (("\n\t" + "") + "foo bar")) + ((((("hello" + "foo bar") + "foo bar") + "a") + ("a" + ("foo bar" + ("" + "\n\t")))) + "a")) + ("a" + "a"))) + "\n\t") {
                        case "hello":
                            var v56 uint = 42
                            _ = v56
                        case "\n\t":
                            v3 = v3
                        case "":
                            v3 = false
                        }
                        v57 := []int{32767, v1, 32767, (fi29 * v2)}
                        _ = v57
                    }
                }
            case 41:
                v58 := GenS3{F0: (v4 / ((v4) | 1)), F1: -9223372036854775808}
                _ = v58
            case 32767:
                for fi59 := 0; fi59 < 17; fi59++ {
                    if (v1 > (v2 + fi59)) {
                        if v3 {
                            v60 := []int{}
                            v61 := GenS3{F0: (v4 - v4), F1: -1}
                            var v62 string = ""
                            _ = v60
                            _ = v61
                            _ = v62
                        }
                        v63 := GenS4{F0: -378, F1: (v5 * (v5 | (v5 << ((42) & 63))))}
                        if (!(239 < 18446744073709551615)) {
                            if (v3 && v3) {
                                v2 = v1
                                var v64 uint16 = 380
                                _ = v64
                            } else {
                                var v65 uint64 = uint64(v4)
                                _ = v65
                            }
                            var v66 uint32 = 0
                            v67 := GenS2{F0: 1, F1: 100}
                            v68 := GenS5{F0: ((((-2.25 * 3.14159) + -2.25) + float32((1e10 + 3.14159))) >= (1e10 * 0.0)), F1: 1, F2: v3}
                            var v69 string = "hello"
                            _ = v66
                            _ = v67
                            _ = v68
                            _ = v69
                        } else {
                            var v70 uint16 = 2
                            var v71 float32 = float32((-2.25 / ((1.5 - 0.1) + -2.25)))
                            _ = v70
                            _ = v71
                        }
                        _ = v63
                    } else {
                        var v72 uint16 = 0
                        _ = v72
                    }
                    _ = fi59
                }
            default:
                v73 := GenS3{F0: (v4 >> ((914) & 63)), F1: -779}
                _ = v73
            }
        } else {
            if ((v4 >= 1) || false) {
                v1 = -9223372036854775808
                var v74 bool = true
                _ = v74
            } else {
                for fi75 := 0; fi75 < 4; fi75++ {
                    v3 = v3
                    if v3 {
                        if v3 {
                            if (!(!(!true))) {
                                v76 := []int{(v2 % ((v1) | 1)), v2, (-(fi29))}
                                v77 := GenS4{F0: int16(v4), F1: v5}
                                if (false || false) {
                                    if ("hello" < "hello") {
                                        v4 -= (v4 & v4)
                                        if v3 {
                                            if (((((65535 != 1) && (fi29 < (v2 + 32768))) && v3) && v3) || ((^(v4)) >= v4)) {
                                                if (!(263 < 42)) {
                                                    v78 := GenS2{F0: 32767, F1: 32}
                                                    _ = v78
                                                }
                                                var v79 int8 = 70
                                                v3 = false
                                                var v80 int8 = (v79 ^ (^(v79)))
                                                v76 = append(v76, v1)
                                                _ = v79
                                                _ = v80
                                            } else {
                                                v2 = fi29
                                                v4 = v4
                                                if (!false) {
                                                    if v3 {
                                                        if (255 != 65535) {
                                                            v5 = uint8(v4)
                                                            var v81 uint8 = v5
                                                            var v82 uint64 = 375
                                                            _ = v81
                                                            _ = v82
                                                        } else {
                                                            if v3 {
                                                                var v83 uint32 = v4
                                                                if v3 {
                                                                    _ = len(v76)
                                                                    v84 := GenS1{F0: -32768}
                                                                    fmt.Println(fi29, v3)
                                                                    _ = v84
                                                                }
                                                                _ = v83
                                                            } else {
                                                                if len(v76) >= 2 {
                                                                    v76 = v76[1:len(v76)]
                                                                }
                                                                if v3 {
                                                                    var v85 int64 = 787
                                                                    switch "a" {
                                                                    case "hello":
                                                                        v4 += v4
                                                                    case "foo bar":
                                                                        v1 = 32768
                                                                    }
                                                                    _ = v85
                                                                } else {
                                                                    switch (^(fi75)) {
                                                                    case -206:
                                                                        fmt.Println(v2, v1, v5)
                                                                    case -978:
                                                                        v2 = -9223372036854775808
                                                                    case 2147483647:
                                                                        v2 += (v2 * (^(fi29)))
                                                                    }
                                                                    if len(v76) > 0 {
                                                                        v76[0] = (v2 - (fi75 << ((2) & 63)))
                                                                    }
                                                                }
                                                                if len(v76) >= 2 {
                                                                    v76 = v76[1:len(v76)]
                                                                }
                                                            }
                                                        }
                                                    }
                                                    switch "foo bar" {
                                                    case "a":
                                                        if true {
                                                            if (v3 || ((!(-25 != 0)) && v3)) {
                                                                if v3 {
                                                                    v2 += fi75
                                                                    v5 += (^(v5))
                                                                } else {
                                                                    v2 = (v1 % (((fi29 + 644)) | 1))
                                                                    v4 = 42
                                                                    if len(v76) > 0 {
                                                                        v76[0] = fi29
                                                                    }
                                                                }
                                                            } else {
                                                                if len(v76) >= 2 {
                                                                    v76 = v76[1:len(v76)]
                                                                }
                                                                v5 -= v5
                                                                v4 += 1
                                                            }
                                                            v86 := GenS2{F0: 65535, F1: 114}
                                                            v87 := GenS1{F0: 1}
                                                            _ = v86
                                                            _ = v87
                                                        }
                                                    case "":
                                                        switch "" {
                                                        case "\n\t":
                                                            _ = len(v76)
                                                        case "":
                                                            v76 = append(v76, (v1 + v1))
                                                        case "hello":
                                                            v4 -= v4
                                                        default:
                                                            if (((v3 || v3) && (1.5 == (3.14159 + -2.25))) || (("foo bar" < "hello") && v3)) {
                                                                if (!(!(!((v3 || (!(-32768 < 0))) || ((42 >= 9223372036854775807) || v3))))) {
                                                                    v76 = append(v76, (v1 % (((v1 ^ (^(v1)))) | 1)))
                                                                } else {
                                                                    var v88 int32 = -2147483648
                                                                    if len(v76) > 0 {
                                                                        v76[0] = (fi29 & fi75)
                                                                    }
                                                                    v89 := GenS3{F0: 42, F1: -9223372036854775808}
                                                                    _ = v88
                                                                    _ = v89
                                                                }
                                                                if v3 {
                                                                    v1 = v2
                                                                    fmt.Println(v1, v5)
                                                                    v90 := GenS6{F0: 42, F1: 0, F2: (^(v4))}
                                                                    _ = v90
                                                                } else {
                                                                    var v91 int64 = -9223372036854775808
                                                                    _ = v91
                                                                }
                                                                switch "a" {
                                                                case "":
                                                                    if v3 {
                                                                        switch (^(v1)) {
                                                                        case -498:
                                                                            v92 := GenS7{F0: 1}
                                                                            _ = v92
                                                                        }
                                                                        if ("a" == "") {
                                                                            if (v3 || (true || true)) {
                                                                                if len(v76) >= 2 {
                                                                                    v76 = v76[1:len(v76)]
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                v93 := GenS1{F0: -32768}
                                                                _ = v93
                                                            } else {
                                                                if len(v76) >= 2 {
                                                                    v76 = v76[1:len(v76)]
                                                                }
                                                                _ = len(v76)
                                                                fmt.Println(v3, v3, v1)
                                                            }
                                                        }
                                                    }
                                                    v2 = v1
                                                    var v94 int32 = -193
                                                    v2 -= 128
                                                    _ = v94
                                                }
                                                _ = len(v76)
                                            }
                                            var v95 float32 = float32(1e10)
                                            var v96 int32 = 525
                                            _ = v95
                                            _ = v96
                                        } else {
                                            _ = len(v76)
                                        }
                                        v97 := GenS8{F0: (-30 <= 817)}
                                        _ = v97
                                    }
                                }
                                _ = v76
                                _ = v77
                            }
                        }
                        if (!v3) {
                            v98 := GenS8{F0: true}
                            var v99 uint32 = v4
                            var v100 int16 = 32767
                            _ = v100
                            _ = v98
                            _ = v99
                        }
                    }
                    v5 = 2
                    switch (v2 | (fi29 >> ((65535) & 63))) {
                    case 308:
                        var v101 uint8 = v5
                        _ = v101
                    default:
                        var v102 uint8 = v5
                        _ = v102
                    }
                    v103 := []int{}
                    _ = fi75
                    _ = v103
                }
                for fi104 := 0; fi104 < 13; fi104++ {
                    if (255 <= 686) {
                        var v105 uint8 = v5
                        _ = v105
                    }
                    if (-699 != 1) {
                        if (!(2 < 18446744073709551615)) {
                            var v106 int16 = 536
                            _ = v106
                        }
                        switch "hello" {
                        case "\n\t":
                            v107 := []int{(fi29 | v2), (fi29 << ((2) & 63))}
                            _ = v107
                        case "":
                            switch (^(fi29)) {
                            case 127:
                                var v108 string = "hello"
                                _ = v108
                            case 603:
                                v1 += (v1 - v1)
                            default:
                                fmt.Println(v5, v5, v3)
                            }
                        case "foo bar":
                            v109 := []int{(v1 & fi104), v1, (fi29 << ((1) & 63))}
                            v110 := GenS9{F0: 292, F1: 650}
                            if (v3 || (!(true || (0 >= -549)))) {
                                fmt.Println(v5)
                            }
                            _ = v109
                            _ = v110
                        }
                        v111 := []int{(v1 % (((fi104 << ((1) & 63))) | 1)), v2, (v1 >> ((1) & 63))}
                        _ = v111
                    } else {
                        v5 += v5
                        switch (-(fi104)) {
                        case 4294967295:
                            if (42 != 949) {
                                if (false || (716 > -2147483648)) {
                                    var v112 float32 = float32(1.5)
                                    v113 := []int{}
                                    switch -1 {
                                    case -27:
                                        switch (fi104 + fi104) {
                                        case 0:
                                            var v114 uint16 = 1
                                            _ = v114
                                        case -1:
                                            if false {
                                                v4 += (v4 % (((v4 * v4)) | 1))
                                                if v3 {
                                                    var v115 float32 = ((1e10 - 1.5) * v112)
                                                    _ = v115
                                                }
                                            }
                                        case -9223372036854775808:
                                            var v116 uint32 = uint32(v5)
                                            _ = v116
                                        }
                                    case 9223372036854775807:
                                        if len(v113) > 0 {
                                            v113[0] = fi29
                                        }
                                        v3 = true
                                    case 0:
                                        v3 = false
                                        _ = len(v113)
                                    }
                                    _ = v112
                                    _ = v113
                                } else {
                                    v117 := []int{(fi104 * (v1 / (((fi29 ^ (v2 / (((v1 ^ v1)) | 1)))) | 1))), (-(v1)), 32767, (fi104 << ((4294967295) & 63))}
                                    _ = v117
                                }
                            } else {
                                fmt.Println(v4, fi29)
                                var v118 int16 = 42
                                _ = v118
                            }
                        default:
                            v3 = false
                        }
                        var v119 int8 = 1
                        _ = v119
                    }
                    v5 -= (v5 / (((v5 >> ((65535) & 63))) | 1))
                    _ = fi104
                }
                var v120 string = ("foo bar" + "hello")
                _ = v120
            }
        }
        v121 := GenS10{F0: (v5 >> ((4294967295) & 63))}
        _ = fi29
        _ = v121
    }
    if (597 <= 42) {
        v122 := GenS11{F0: 4294967295, F1: 1, F2: 0}
        _ = v122
    }
    v1 -= (v1 ^ 1)
    fmt.Println(v1, v3)
    for fi123 := 0; fi123 < 13; fi123++ {
        v5 = (v5 | v5)
        v124 := GenS5{F0: v3, F1: -745, F2: (-1 != -9223372036854775808)}
        for fi125 := 0; fi125 < 19; fi125++ {
            v126 := GenS5{F0: v3, F1: -1, F2: v3}
            if false {
                v3 = v3
                var v127 uint16 = uint16(v5)
                _ = v127
            }
            var v128 uint8 = v5
            _ = fi125
            _ = v126
            _ = v128
        }
        _ = fi123
        _ = v124
    }
    var v129 string = (("hello" + "hello") + "hello")
    v3 = (1 < 670)
    for fi130 := 0; fi130 < 7; fi130++ {
        fmt.Println(v129)
        if (((float32(3.14159) + ((1e10 - ((1.5 - (1e10 - 0.1)) / 1.5)) * 1e10)) > 1.5) || false) {
            if v3 {
                if v3 {
                    v131 := []int{-2147483648, 4294967295}
                    fmt.Println(v3)
                    var v132 int64 = 432
                    _ = v131
                    _ = v132
                } else {
                    if (v3 || (float32(((float64(-2.25) * 0.1) / 1.0)) >= (1e10 + 1.5))) {
                        if ((!(1 >= 127)) || (!(!v3))) {
                            var v133 int32 = 2147483647
                            for fi134 := 0; fi134 < 2; fi134++ {
                                fmt.Println(v2, v1, fi134)
                                var v135 int8 = int8(v133)
                                v1 -= 127
                                _ = fi134
                                _ = v135
                            }
                            _ = v133
                        }
                        var v136 int64 = -417
                        _ = v136
                    }
                    v137 := GenS6{F0: -830, F1: 42, F2: v4}
                    _ = v137
                }
            }
            var v138 string = (v129 + "foo bar")
            _ = v138
        }
        _ = fi130
    }
    v139 := GenS12{F0: 753, F1: v129, F2: 0}
    if v3 {
        v1 += v1
        if (!(!(!v3))) {
            v129 = (v129 + "")
            switch v129 {
            case "":
                var v140 uint32 = v4
                _ = v140
            }
            var v141 uint16 = 304
            _ = v141
        } else {
            v142 := []int{(v2 << ((1) & 63)), (^(v1))}
            switch v2 {
            case 128:
                v142 = append(v142, (v2 & -9223372036854775808))
            }
            fmt.Println(v5)
            _ = v142
        }
    }
    v2 += 42
    v143 := GenS1{F0: -32768}
    var v144 uint32 = (v4 % ((65535) | 1))
    if (!v3) {
        v145 := GenS3{F0: (v4 % ((v4) | 1)), F1: -9223372036854775808}
        switch v2 {
        case 32767:
            var v146 int16 = 1
            _ = v146
        case 2147483647:
            v147 := GenS1{F0: 42}
            _ = v147
        case 9223372036854775807:
            for fi148 := 0; fi148 < 14; fi148++ {
                switch fi148 {
                case 4294967295:
                    fmt.Println(v129, v2, v1)
                default:
                    var v149 uint = 2
                    _ = v149
                }
                for fi150 := 0; fi150 < 15; fi150++ {
                    if v3 {
                        switch (-(v1)) {
                        case 32767:
                            var v151 int8 = 127
                            _ = v151
                        case 32768:
                            v3 = v3
                        case -2147483648:
                            if (((v129 + v129) + v129) > v129) {
                                if (42 == 564) {
                                    var v152 uint64 = 0
                                    if ((((0.1 * (((-2.25 + 0.1) * (3.14159 - 1e10)) - 3.14159)) * 0.0) - 3.14159) < ((1.5 * -2.25) * (float64((-2.25 + -2.25)) - 0.0))) {
                                        fmt.Println(v2, v4)
                                    } else {
                                        if v3 {
                                            v153 := []int{(v2 >> ((255) & 63)), (^(fi148))}
                                            if (!(-542 < -925)) {
                                                if v3 {
                                                    if v3 {
                                                        var v154 uint64 = uint64(v4)
                                                        _ = v154
                                                    }
                                                    var v155 bool = true
                                                    var v156 int16 = 237
                                                    _ = v155
                                                    _ = v156
                                                }
                                                if v3 {
                                                    var v157 int16 = 140
                                                    if (116 <= 111) {
                                                        var v158 int16 = 0
                                                        _ = v158
                                                    } else {
                                                        fmt.Println(v3, v1, v5)
                                                    }
                                                    _ = v157
                                                }
                                            }
                                            _ = v153
                                        } else {
                                            var v159 uint32 = (v4 | v144)
                                            v160 := GenS2{F0: 255, F1: 1}
                                            _ = v159
                                            _ = v160
                                        }
                                    }
                                    _ = v152
                                }
                            } else {
                                if v3 {
                                    v129 = "a"
                                    v161 := GenS11{F0: 2, F1: 4294967295, F2: 42}
                                    v162 := []int{fi148, fi150, v1, v1}
                                    _ = v161
                                    _ = v162
                                }
                                v129 = (v129 + v129)
                                var v163 string = v129
                                var v164 uint32 = 4294967295
                                v165 := GenS2{F0: 0, F1: 0}
                                _ = v163
                                _ = v164
                                _ = v165
                            }
                        default:
                            var v166 uint16 = 0
                            _ = v166
                        }
                    } else {
                        var v167 uint16 = 2
                        var v168 uint32 = (v144 & uint32(v5))
                        if (-37 <= -2147483648) {
                            v169 := []int{(v2 & (v1 & 9223372036854775807)), (fi150 ^ -2147483648)}
                            if ((v5 / (((v5 << ((0) & 63))) | 1)) == 2) {
                                if true {
                                    v170 := GenS7{F0: 18446744073709551615}
                                    if len(v169) >= 2 {
                                        v169 = v169[1:len(v169)]
                                    }
                                    var v171 bool = ((v4 >> ((18446744073709551615) & 63)) == 309)
                                    _ = v170
                                    _ = v171
                                } else {
                                    fmt.Println(v129)
                                    var v172 int64 = -1
                                    _ = v172
                                }
                                var v173 int16 = 1000
                                _ = v173
                            }
                            var v174 uint64 = 42
                            _ = v169
                            _ = v174
                        }
                        _ = v167
                        _ = v168
                    }
                    v4 = 456
                    _ = fi150
                }
                switch v129 {
                case "a":
                    v175 := GenS2{F0: 255, F1: 1}
                    _ = v175
                case "foo bar":
                    fmt.Println(v144)
                }
                _ = fi148
            }
        default:
            for fi176 := 0; fi176 < 3; fi176++ {
                var v177 string = (v129 + "")
                _ = fi176
                _ = v177
            }
        }
        var v178 float64 = 0.1
        _ = v145
        _ = v178
    } else {
        switch v2 {
        case 32768:
            for fi179 := 0; fi179 < 13; fi179++ {
                fmt.Println(v5, v5, v3)
                v5 = (v5 ^ uint8(v144))
                v180 := []int{fi179, v2, (^(v1))}
                _ = fi179
                _ = v180
            }
        case -1:
            var v181 float32 = 3.14159
            _ = v181
        default:
            v182 := []int{52, -262, (v2 | v2), 2147483647}
            _ = v182
        }
        for fi183 := 0; fi183 < 13; fi183++ {
            v184 := []int{(fi183 + 1), (v2 >> ((0) & 63))}
            v5 -= v5
            var v185 uint32 = (v144 - v144)
            _ = fi183
            _ = v184
            _ = v185
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v129:", v129)
    fmt.Println("v144:", v144)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v129
    _ = v139
    _ = v143
    _ = v144
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
