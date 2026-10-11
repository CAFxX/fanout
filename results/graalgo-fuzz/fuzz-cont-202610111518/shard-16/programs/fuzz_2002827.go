package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint8; F2 int16 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 uint32; F1 int16; F2 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int16; F1 string }
type GenS9 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS10 struct { F0 string; F1 uint16 }
type GenS11 struct { F0 int; F1 int64 }
type GenS12 struct { F0 int32; F1 uint64 }
type GenS13 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 uint8; F1 uint64 }
type GenS17 struct { F0 uint64; F1 int16 }
type GenS18 struct { F0 int8; F1 int; F2 int16 }
type GenS19 struct { F0 int; F1 int }
type GenS20 struct { F0 int; F1 int8 }
type GenS21 struct { F0 uint; F1 int16 }
type GenS22 struct { F0 int16; F1 string }
type GenS23 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS24 struct { F0 int64; F1 uint64; F2 int64 }
type GenS25 struct { F0 uint; F1 int16; F2 uint64 }
type GenS26 struct { F0 uint; F1 int8 }
type GenS27 struct { F0 bool; F1 bool; F2 uint }
type GenS28 struct { F0 uint16; F1 uint32; F2 bool }
type GenS29 struct { F0 uint8; F1 int16 }
type GenS30 struct { F0 string; F1 int8 }
type GenS31 struct { F0 uint; F1 uint64 }
type GenS32 struct { F0 int16; F1 uint32; F2 uint16 }
type GenS33 struct { F0 int8 }
type GenS34 struct { F0 uint; F1 string; F2 uint }
type GenS35 struct { F0 bool; F1 uint8; F2 uint16 }
type GenS36 struct { F0 uint16; F1 int; F2 int16 }
type GenS37 struct { F0 uint16; F1 string }
type GenS38 struct { F0 int64; F1 int; F2 bool }
func genFn0(p0 int) bool {
    if (0.1 <= (1e10 - 3.14159)) {
        var v7 int16 = 1000
        _ = v7
    }
    fmt.Println(p0)
    _ = p0
    return false
}
func genFn1() uint64 {
    if false {
        var v8 string = "foo bar"
        v8 = "a"
        if true {
            switch (v8 + v8) {
            case "a":
                v9 := GenS1{F0: (v8 + "\n\t"), F1: 1}
                _ = v9
            }
            v10 := GenS2{F0: 42, F1: 56, F2: 830}
            switch 32768 {
            case 42:
                var v11 uint64 = 724
                _ = v11
            case -2147483648:
                v12 := GenS2{F0: 4294967295, F1: 58, F2: -605}
                _ = v12
            case 9223372036854775807:
                v8 = v8
            default:
                v13 := GenS3{F0: 518, F1: 705}
                _ = v13
            }
            _ = v10
        } else {
            var v14 bool = (400 < -443)
            v15 := GenS1{F0: v8, F1: 32767}
            switch 128 {
            case -1:
                switch v8 {
                case "":
                    var v16 uint16 = 2
                    _ = v16
                default:
                    if (!(100 == 0)) {
                        if (!((1 > 18446744073709551615) || false)) {
                            for fi17 := 0; fi17 < 11; fi17++ {
                                fmt.Println(fi17, v14, v14)
                                var v18 uint16 = 1
                                v19 := GenS1{F0: ("" + (v8 + ("foo bar" + v8))), F1: v18}
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                            switch -792 {
                            case 2147483647:
                                var v20 int = 346
                                _ = v20
                            case 128:
                                var v21 uint64 = 221
                                _ = v21
                            }
                        }
                        for fi22 := 0; fi22 < 12; fi22++ {
                            switch (("" + "\n\t") + v8) {
                            case "":
                                v23 := GenS3{F0: 65535, F1: 2147483647}
                                _ = v23
                            default:
                                v24 := GenS1{F0: "foo bar", F1: 42}
                                _ = v24
                            }
                            if false {
                                v25 := GenS3{F0: 4294967295, F1: 1}
                                v14 = v14
                                if false {
                                    v26 := []int{9223372036854775807, fi22, fi22}
                                    v27 := GenS4{F0: 1, F1: 361, F2: 2147483647}
                                    if v14 {
                                        if true {
                                            if len(v26) > 0 {
                                                v26[0] = (fi22 ^ (fi22 | 127))
                                            }
                                            var v28 uint16 = 971
                                            _ = v28
                                        } else {
                                            switch ((v8 + v8) + v8) {
                                            case "":
                                                v8 = v8
                                            case "foo bar":
                                                if len(v26) >= 2 {
                                                    v26 = v26[1:len(v26)]
                                                }
                                            case "hello":
                                                if (0 <= 128) {
                                                    switch (fi22 >> ((361) & 63)) {
                                                    case -151:
                                                        var v29 uint = 398
                                                        _ = v29
                                                    case -505:
                                                        var v30 uint = 2
                                                        _ = v30
                                                    }
                                                } else {
                                                    v31 := GenS4{F0: 697, F1: -32768, F2: -2147483648}
                                                    if true {
                                                        for fi32 := 0; fi32 < 6; fi32++ {
                                                            var v33 float64 = float64(1e10)
                                                            _ = fi32
                                                            _ = v33
                                                        }
                                                    } else {
                                                        v34 := GenS5{F0: 849}
                                                        _ = v34
                                                    }
                                                    _ = v31
                                                }
                                            }
                                            var v35 int8 = -112
                                            _ = v35
                                        }
                                        v14 = v14
                                    }
                                    _ = v26
                                    _ = v27
                                }
                                _ = v25
                            } else {
                                v36 := []int{fi22}
                                v14 = ("foo bar" < "")
                                _ = v36
                            }
                            _ = fi22
                        }
                    } else {
                        for fi37 := 0; fi37 < 7; fi37++ {
                            switch v8 {
                            case "":
                                v8 = v8
                            default:
                                switch (fi37 ^ fi37) {
                                case 127:
                                    var v38 int64 = -1
                                    _ = v38
                                case 681:
                                    v39 := GenS6{F0: 42}
                                    _ = v39
                                case 128:
                                    for fi40 := 0; fi40 < 1; fi40++ {
                                        if true {
                                            v41 := GenS7{F0: 127}
                                            switch v8 {
                                            case "hello":
                                                var v42 uint8 = 0
                                                _ = v42
                                            case "":
                                                v43 := GenS8{F0: 0, F1: v8}
                                                _ = v43
                                            }
                                            if (!("a" > (v8 + v8))) {
                                                v44 := []int{-93, fi37}
                                                fmt.Println(v14, v8, fi40)
                                                _ = v44
                                            } else {
                                                switch v8 {
                                                case "hello":
                                                    fmt.Println(v8, v14, fi37)
                                                case "\n\t":
                                                    v14 = false
                                                default:
                                                    var v45 int8 = 42
                                                    _ = v45
                                                }
                                                switch fi37 {
                                                case 0:
                                                    var v46 uint16 = 42
                                                    _ = v46
                                                }
                                                var v47 float64 = ((float64(-2.25) + 1.5) * -2.25)
                                                _ = v47
                                            }
                                            _ = v41
                                        }
                                        _ = fi40
                                    }
                                }
                            }
                            for fi48 := 0; fi48 < 13; fi48++ {
                                switch (fi37 | fi48) {
                                case 105:
                                    if true {
                                        var v49 float64 = (((1e10 * ((0.0 * (-2.25 + -2.25)) * -2.25)) + (0.0 * 3.14159)) * 3.14159)
                                        var v50 uint64 = 42
                                        var v51 uint32 = 0
                                        _ = v49
                                        _ = v50
                                        _ = v51
                                    } else {
                                        v14 = v14
                                        v52 := GenS9{F0: 2, F1: 38, F2: 32767}
                                        _ = v52
                                    }
                                default:
                                    v8 = v8
                                }
                                _ = fi48
                            }
                            _ = fi37
                        }
                    }
                }
            case 128:
                v14 = (!false)
            }
            _ = v14
            _ = v15
        }
        _ = v8
    }
    for fi53 := 0; fi53 < 17; fi53++ {
        for fi54 := 0; fi54 < 1; fi54++ {
            v55 := []int{}
            _ = fi54
            _ = v55
        }
        for fi56 := 0; fi56 < 8; fi56++ {
            v57 := []int{}
            v58 := GenS10{F0: "hello", F1: 520}
            _ = fi56
            _ = v57
            _ = v58
        }
        switch "\n\t" {
        case "hello":
            v59 := []int{fi53, 714, fi53, (fi53 % (((fi53 | fi53)) | 1))}
            _ = v59
        case "":
            if true {
                var v60 float64 = 1.5
                if false {
                    for fi61 := 0; fi61 < 7; fi61++ {
                        fmt.Println(fi61)
                        _ = fi61
                    }
                    switch (fi53 - 42) {
                    case 0:
                        v62 := GenS11{F0: (fi53 - (fi53 | 127)), F1: 0}
                        _ = v62
                    case -2147483648:
                        var v63 int = fi53
                        _ = v63
                    default:
                        if (false && false) {
                            switch "a" {
                            case "hello":
                                v64 := GenS4{F0: 2, F1: 1, F2: 1}
                                _ = v64
                            case "":
                                fmt.Println(fi53, fi53)
                            case "\n\t":
                                fmt.Println(fi53, fi53)
                            }
                            v60 -= -2.25
                            if true {
                                v65 := GenS12{F0: 2147483647, F1: 18446744073709551615}
                                switch -547 {
                                case -1:
                                    if (true || (true && true)) {
                                        for fi66 := 0; fi66 < 14; fi66++ {
                                            var v67 uint = 0
                                            if true {
                                                v60 = v60
                                                if (false && ("\n\t" <= ("a" + "a"))) {
                                                    var v68 int64 = 0
                                                    v69 := []int{(fi66 << ((18446744073709551615) & 63)), (fi53 & 42), fi66, (fi66 / (((fi66 % (((^(fi53))) | 1))) | 1))}
                                                    v68 = -991
                                                    _ = v68
                                                    _ = v69
                                                } else {
                                                    v67 -= 65535
                                                }
                                                if true {
                                                    var v70 uint16 = 1
                                                    _ = v70
                                                } else {
                                                    var v71 float32 = float32((1.5 + float64(0.0)))
                                                    var v72 uint16 = 32767
                                                    switch fi53 {
                                                    case 128:
                                                        v73 := GenS13{F0: 42, F1: 135, F2: 100}
                                                        _ = v73
                                                    case -131:
                                                        switch "\n\t" {
                                                        case "hello":
                                                            var v74 int8 = 105
                                                            _ = v74
                                                        default:
                                                            v71 -= v71
                                                            switch fi53 {
                                                            case -1:
                                                                v75 := []int{2147483647, fi53, (fi66 ^ (fi53 / (((-(fi66))) | 1))), (fi53 + 32767)}
                                                                _ = v75
                                                            case 9223372036854775807:
                                                                v76 := []int{(fi66 & fi66), 128, (fi66 % ((2147483647) | 1))}
                                                                _ = v76
                                                            case 128:
                                                                if false {
                                                                    if false {
                                                                        var v77 uint64 = 4294967295
                                                                        switch fi53 {
                                                                        case -9223372036854775808:
                                                                            var v78 int8 = -89
                                                                            _ = v78
                                                                        case 42:
                                                                            if true {
                                                                                var v79 int64 = 42
                                                                                v80 := []int{(^(fi53)), (fi53 - (fi66 * fi66)), fi53, fi53}
                                                                                var v81 uint = 255
                                                                                _ = v79
                                                                                _ = v80
                                                                                _ = v81
                                                                            }
                                                                        }
                                                                        _ = v77
                                                                    }
                                                                } else {
                                                                    switch fi53 {
                                                                    case -753:
                                                                        var v82 uint32 = 65535
                                                                        _ = v82
                                                                    }
                                                                }
                                                                if false {
                                                                    v83 := GenS14{F0: false}
                                                                    if (32767 > 42) {
                                                                        var v84 uint = 304
                                                                        v85 := []int{fi53, fi66}
                                                                        var v86 int16 = 884
                                                                        _ = v84
                                                                        _ = v85
                                                                        _ = v86
                                                                    }
                                                                    _ = v83
                                                                }
                                                                switch (("a" + "a") + "a") {
                                                                case "a":
                                                                    v67 -= (v67 & v67)
                                                                case "\n\t":
                                                                    fmt.Println(fi53, v72)
                                                                case "foo bar":
                                                                    var v87 int = 1
                                                                    _ = v87
                                                                default:
                                                                    v60 = v60
                                                                }
                                                            default:
                                                                v67 = (v67 | (v67 + v67))
                                                            }
                                                            v88 := GenS15{F0: 0}
                                                            _ = v88
                                                        }
                                                    }
                                                    v89 := GenS16{F0: 42, F1: 1}
                                                    v90 := GenS17{F0: 18446744073709551615, F1: 32767}
                                                    if ((!true) || false) {
                                                        if (127 < 0) {
                                                            v91 := []int{(fi53 ^ (^(fi53))), fi66, (-(fi53))}
                                                            _ = v91
                                                        } else {
                                                            var v92 float32 = v71
                                                            _ = v92
                                                        }
                                                        var v93 int32 = 522
                                                        switch ("" + ("\n\t" + ("hello" + ("" + "foo bar")))) {
                                                        case "hello":
                                                            v94 := []int{(fi53 ^ fi66), fi53}
                                                            switch ("" + "") {
                                                            case "\n\t":
                                                                var v95 uint8 = 42
                                                                _ = v95
                                                            default:
                                                                if len(v94) >= 2 {
                                                                    v94 = v94[1:len(v94)]
                                                                }
                                                                var v96 uint64 = 18446744073709551615
                                                                _ = v96
                                                            }
                                                            fmt.Println(v93, v72, fi66)
                                                            _ = v94
                                                        case "foo bar":
                                                            var v97 uint32 = 327
                                                            _ = v97
                                                        }
                                                        _ = v93
                                                    } else {
                                                        if false {
                                                            if ((^(v67)) >= (v67 << ((v67) & 63))) {
                                                                v98 := GenS16{F0: 1, F1: 1}
                                                                if false {
                                                                    var v99 bool = false
                                                                    v72 -= 1
                                                                    fmt.Println(v71, v67, v71)
                                                                    _ = v99
                                                                } else {
                                                                    var v100 uint = 42
                                                                    _ = v100
                                                                }
                                                                switch 32767 {
                                                                case 825:
                                                                    var v101 int = (fi53 % ((fi66) | 1))
                                                                    _ = v101
                                                                case -385:
                                                                    var v102 uint = (v67 ^ (v67 + v67))
                                                                    _ = v102
                                                                case 0:
                                                                    if false {
                                                                        switch "hello" {
                                                                        case "":
                                                                            v103 := []int{-9223372036854775808}
                                                                            _ = v103
                                                                        case "foo bar":
                                                                            v104 := GenS17{F0: 42, F1: 0}
                                                                            _ = v104
                                                                        default:
                                                                            v105 := []int{fi53, (fi53 * fi66), (fi53 / (((^(fi53))) | 1))}
                                                                            _ = v105
                                                                        }
                                                                        v106 := []int{(fi66 << ((v67) & 63)), (fi66 >> ((v67) & 63)), (fi53 & 2147483647), 488}
                                                                        _ = v106
                                                                    }
                                                                default:
                                                                    if (true || true) {
                                                                        if (((fi53 & 542) < 9223372036854775807) || true) {
                                                                            fmt.Println(v71, v72)
                                                                        } else {
                                                                            var v107 int64 = 42
                                                                            _ = v107
                                                                        }
                                                                        switch ("\n\t" + ("foo bar" + "hello")) {
                                                                        case "foo bar":
                                                                            v108 := GenS18{F0: 127, F1: 155, F2: 32767}
                                                                            _ = v108
                                                                        }
                                                                        fmt.Println(v60, v60, v72)
                                                                        v109 := GenS19{F0: -1, F1: fi66}
                                                                        _ = v109
                                                                    } else {
                                                                        v67 = v67
                                                                    }
                                                                }
                                                                _ = v98
                                                            } else {
                                                                var v110 uint32 = 42
                                                                v72 = (v72 | (v72 ^ 2))
                                                                var v111 int = 0
                                                                switch (fi53 / ((v111) | 1)) {
                                                                case 32768:
                                                                    v71 = (v71 + 0.1)
                                                                case 284:
                                                                    if (-1 == 958) {
                                                                        v112 := GenS20{F0: -9223372036854775808, F1: 127}
                                                                        v110 = v110
                                                                        v111 += (fi66 / ((42) | 1))
                                                                        _ = v112
                                                                    } else {
                                                                        var v113 int8 = -1
                                                                        v111 = v111
                                                                        v67 = (v67 & (^(v67)))
                                                                        _ = v113
                                                                    }
                                                                default:
                                                                    fmt.Println(v111)
                                                                }
                                                                v111 += fi53
                                                                if false {
                                                                    v72 += v72
                                                                    switch "foo bar" {
                                                                    case "foo bar":
                                                                        var v114 float64 = 0.1
                                                                        _ = v114
                                                                    case "a":
                                                                        var v115 uint64 = 139
                                                                        switch (fi66 % ((fi66) | 1)) {
                                                                        case 32768:
                                                                            if (v71 >= 0.1) {
                                                                                var v116 uint = uint(v72)
                                                                                _ = v116
                                                                            } else {
                                                                                v110 = v110
                                                                                v117 := []int{}
                                                                                _ = v117
                                                                            }
                                                                        }
                                                                        var v118 int64 = -1
                                                                        v110 += (^(v110))
                                                                        var v119 bool = false
                                                                        _ = v115
                                                                        _ = v118
                                                                        _ = v119
                                                                    case "\n\t":
                                                                        v110 = v110
                                                                    }
                                                                    var v120 uint32 = 511
                                                                    fmt.Println(v67, fi66, v120)
                                                                    v121 := GenS19{F0: int(v72), F1: 523}
                                                                    _ = v120
                                                                    _ = v121
                                                                } else {
                                                                    v122 := GenS7{F0: 71}
                                                                    v67 += v67
                                                                    if true {
                                                                        v72 = (^(v72))
                                                                        v123 := []int{}
                                                                        var v124 float64 = v60
                                                                        var v125 bool = ((fi53 / ((32768) | 1)) >= (fi66 / ((fi66) | 1)))
                                                                        v126 := GenS5{F0: 32767}
                                                                        switch (^(v111)) {
                                                                        case 9223372036854775807:
                                                                            v60 -= (0.0 * ((((v124 + 0.0) - v124) + 1.5) * 0.1))
                                                                        default:
                                                                            v125 = false
                                                                        }
                                                                        _ = v123
                                                                        _ = v124
                                                                        _ = v125
                                                                        _ = v126
                                                                    }
                                                                    _ = v122
                                                                }
                                                                var v127 uint = 4294967295
                                                                _ = v110
                                                                _ = v111
                                                                _ = v127
                                                            }
                                                            v128 := GenS9{F0: v72, F1: 84, F2: (v72 >> ((v67) & 63))}
                                                            _ = v128
                                                        }
                                                        v129 := GenS12{F0: 42, F1: 1}
                                                        _ = v129
                                                    }
                                                    v130 := GenS12{F0: 2147483647, F1: 0}
                                                    _ = v130
                                                    _ = v71
                                                    _ = v72
                                                    _ = v89
                                                    _ = v90
                                                }
                                            }
                                            switch "foo bar" {
                                            case "":
                                                v131 := GenS21{F0: v67, F1: -429}
                                                _ = v131
                                            case "\n\t":
                                                switch (fi53 - fi53) {
                                                case 32768:
                                                    v60 -= v60
                                                case 127:
                                                    v60 = float64((1.5 * float32(v60)))
                                                case 9223372036854775807:
                                                    v132 := []int{9223372036854775807, fi66, (fi53 ^ 127), (-(fi53))}
                                                    _ = v132
                                                default:
                                                    v60 = v60
                                                }
                                            case "hello":
                                                v133 := GenS8{F0: -1, F1: "\n\t"}
                                                _ = v133
                                            default:
                                                switch ("foo bar" + "a") {
                                                case "\n\t":
                                                    switch "\n\t" {
                                                    case "hello":
                                                        var v134 uint8 = 0
                                                        _ = v134
                                                    default:
                                                        v135 := []int{(fi53 >> (((v67 | v67)) & 63)), fi53, (fi66 / ((fi66) | 1)), (fi66 << ((v67) & 63))}
                                                        _ = v135
                                                    }
                                                case "":
                                                    v60 = (float64((float32(v60) - 3.14159)) - 1.5)
                                                default:
                                                    if true {
                                                        var v136 bool = false
                                                        if v136 {
                                                            var v137 uint64 = 2
                                                            _ = v137
                                                        }
                                                        _ = v136
                                                    }
                                                }
                                                switch (fi66 | 1) {
                                                case 0:
                                                    var v138 int64 = 9223372036854775807
                                                    _ = v138
                                                case 32768:
                                                    var v139 uint = v67
                                                    _ = v139
                                                case 127:
                                                    v60 = (float64((0.1 * 0.1)) * float64((3.14159 / 1e10)))
                                                default:
                                                    if false {
                                                        var v140 int16 = 42
                                                        v141 := GenS9{F0: 32767, F1: 1, F2: 130}
                                                        if (((0.1 / 1.0) + float32(v60)) != 1.5) {
                                                            var v142 string = "\n\t"
                                                            fmt.Println(fi66, v140)
                                                            fmt.Println(v67, v60)
                                                            _ = v142
                                                        } else {
                                                            v60 += float64(3.14159)
                                                            fmt.Println(v60, v140, fi53)
                                                        }
                                                        _ = v140
                                                        _ = v141
                                                    }
                                                    switch ("" + "hello") {
                                                    case "\n\t":
                                                        v143 := GenS22{F0: 166, F1: ("foo bar" + "\n\t")}
                                                        _ = v143
                                                    case "foo bar":
                                                        v67 = v67
                                                    case "hello":
                                                        if false {
                                                            v144 := GenS8{F0: -32768, F1: "hello"}
                                                            if (42 < 0) {
                                                                v67 = (^(v67))
                                                                var v145 uint8 = 127
                                                                _ = v145
                                                            } else {
                                                                if (42 < 42) {
                                                                    fmt.Println(fi66, fi66)
                                                                    var v146 uint8 = 255
                                                                    _ = v146
                                                                } else {
                                                                    v147 := GenS23{F0: 2, F1: 211, F2: 432}
                                                                    var v148 uint16 = 1
                                                                    v149 := GenS9{F0: v148, F1: 15, F2: v148}
                                                                    v150 := []int{}
                                                                    var v151 float32 = 1e10
                                                                    _ = v147
                                                                    _ = v148
                                                                    _ = v149
                                                                    _ = v150
                                                                    _ = v151
                                                                }
                                                                switch (fi53 + (fi66 - -2147483648)) {
                                                                case 42:
                                                                    v60 = v60
                                                                case 2147483647:
                                                                    var v152 int32 = -1
                                                                    _ = v152
                                                                default:
                                                                    v60 -= v60
                                                                }
                                                            }
                                                            _ = v144
                                                        } else {
                                                            fmt.Println(v67)
                                                            switch "\n\t" {
                                                            case "hello":
                                                                v153 := []int{(fi53 * (-(fi53))), (fi53 * (fi53 << (((v67 / (((^(v67))) | 1))) & 63)))}
                                                                _ = v153
                                                            default:
                                                                var v154 int8 = 100
                                                                _ = v154
                                                            }
                                                        }
                                                    default:
                                                        v155 := GenS24{F0: -9223372036854775808, F1: 18446744073709551615, F2: 9223372036854775807}
                                                        _ = v155
                                                    }
                                                }
                                            }
                                            var v156 int32 = 1
                                            _ = fi66
                                            _ = v156
                                            _ = v67
                                        }
                                    } else {
                                        v60 = ((v60 * 3.14159) * (v60 + v60))
                                        fmt.Println(fi53, v60)
                                        for fi157 := 0; fi157 < 5; fi157++ {
                                            var v158 int16 = 0
                                            switch "hello" {
                                            case "hello":
                                                var v159 bool = true
                                                _ = v159
                                            case "":
                                                fmt.Println(fi53, v60)
                                            default:
                                                var v160 uint8 = 128
                                                _ = v160
                                            }
                                            v60 -= (v60 + float64((-2.25 * (3.14159 + (-2.25 - -2.25)))))
                                            _ = fi157
                                            _ = v158
                                        }
                                    }
                                default:
                                    v161 := []int{fi53, (fi53 ^ fi53), fi53}
                                    _ = v161
                                }
                                v162 := GenS25{F0: 567, F1: 589, F2: 590}
                                _ = v162
                                _ = v65
                            }
                        } else {
                            v163 := GenS20{F0: (fi53 / ((344) | 1)), F1: 1}
                            fmt.Println(fi53)
                            _ = v163
                        }
                    }
                } else {
                    fmt.Println(fi53, v60)
                    v60 += v60
                    if (!false) {
                        v60 += 1e10
                    } else {
                        if (!false) {
                            v60 = v60
                            for fi164 := 0; fi164 < 16; fi164++ {
                                switch (fi164 - (fi53 << ((389) & 63))) {
                                case 32768:
                                    switch (-(fi164)) {
                                    case 1:
                                        fmt.Println(v60, fi164, v60)
                                        v165 := GenS15{F0: 1}
                                        v60 -= ((3.14159 / v60) + v60)
                                        _ = v165
                                    case -2147483648:
                                        var v166 int16 = 32767
                                        if false {
                                            v167 := GenS26{F0: 65535, F1: -20}
                                            var v168 uint64 = 594
                                            switch "" {
                                            case "foo bar":
                                                v60 -= (v60 - v60)
                                            case "\n\t":
                                                if false {
                                                    v60 -= 1.5
                                                } else {
                                                    var v169 int64 = 52
                                                    fmt.Println(v166, fi53, fi164)
                                                    _ = v169
                                                }
                                            default:
                                                fmt.Println(fi53)
                                            }
                                            _ = v167
                                            _ = v168
                                        } else {
                                            var v170 uint = 2
                                            _ = v170
                                        }
                                        _ = v166
                                    case 9223372036854775807:
                                        v171 := GenS25{F0: 42, F1: 1000, F2: 42}
                                        _ = v171
                                    default:
                                        var v172 uint32 = 42
                                        _ = v172
                                    }
                                default:
                                    v60 -= (((0.0 / 1.0) / ((v60 / 1.0) - ((float64(0.1) + (0.1 * v60)) + v60))) - (v60 * (-2.25 - float64(-2.25))))
                                }
                                fmt.Println(fi164, fi53)
                                var v173 int16 = -391
                                if false {
                                    if false {
                                        v173 -= 0
                                    } else {
                                        v174 := GenS27{F0: false, F1: (42 != 0), F2: 4294967295}
                                        _ = v174
                                    }
                                    if true {
                                        var v175 int16 = (-(v173))
                                        v173 -= v175
                                        if true {
                                            var v176 int64 = 9223372036854775807
                                            _ = v176
                                        }
                                        _ = v175
                                    } else {
                                        v177 := []int{fi53, fi53}
                                        var v178 uint8 = 0
                                        if false {
                                            var v179 int8 = 28
                                            v178 = (v178 * (v178 / (((v178 >> ((65535) & 63))) | 1)))
                                            if true {
                                                switch "" {
                                                case "\n\t":
                                                    if true {
                                                        v177 = append(v177, fi164)
                                                        v180 := GenS25{F0: 42, F1: v173, F2: 1}
                                                        _ = v180
                                                    }
                                                case "hello":
                                                    v181 := GenS28{F0: 1, F1: 386, F2: true}
                                                    _ = v181
                                                }
                                                v178 -= v178
                                            }
                                            _ = v179
                                        }
                                        _ = v177
                                        _ = v178
                                    }
                                    v173 -= -32768
                                }
                                var v182 uint32 = 2147483647
                                _ = fi164
                                _ = v173
                                _ = v182
                            }
                            v183 := GenS21{F0: 4294967295, F1: -1}
                            _ = v183
                        }
                        v184 := []int{fi53, fi53}
                        v185 := GenS22{F0: 42, F1: (((("hello" + ("foo bar" + "\n\t")) + "\n\t") + (("foo bar" + "\n\t") + "foo bar")) + "a")}
                        _ = v184
                        _ = v185
                    }
                }
                v186 := GenS29{F0: 42, F1: -32768}
                _ = v186
                _ = v60
            } else {
                v187 := GenS30{F0: "foo bar", F1: 50}
                _ = v187
            }
        }
        _ = fi53
    }
    return 0
}

func main() {
    var v1 string = ((("a" + ("foo bar" + "\n\t")) + (("\n\t" + ("hello" + "\n\t")) + "\n\t")) + "foo bar")
    v2 := make(map[string]uint32)
    v3 := make(map[int]uint32)
    v4 := []int{-744, 136}
    var v5 uint32 = 42
    var v6 string = v1
    for fi188 := 0; fi188 < 14; fi188++ {
        v6 = "a"
        var v189 int32 = -2147483648
        v6 = v1
        _ = fi188
        _ = v189
    }
    v190 := GenS31{F0: 0, F1: uint64(v5)}
    if ((1 < 42) && false) {
        for fi191 := 0; fi191 < 3; fi191++ {
            switch -722 {
            case -435:
                for fi192 := 0; fi192 < 12; fi192++ {
                    var v193 uint64 = 904
                    if (v1 < ("" + (v1 + ("\n\t" + "hello")))) {
                        if (1 >= 1) {
                            v5 = v5
                            v6 = "foo bar"
                        } else {
                            v193 -= v193
                        }
                        var v194 uint64 = (v193 | 42)
                        var v195 string = (v1 + v6)
                        _ = v194
                        _ = v195
                    } else {
                        fmt.Println(v6)
                    }
                    _ = fi192
                    _ = v193
                }
            case 0:
                fmt.Println(v1, v5, v6)
            case -2147483648:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            default:
                v5 = v5
            }
            v1 = "\n\t"
            _ = fi191
        }
    } else {
        var v196 uint16 = 255
        _ = v196
    }
    fmt.Println(v1, v5)
    v197 := GenS18{F0: -1, F1: -1, F2: -99}
    switch "hello" {
    case "a":
        _ = len(v2)
    case "hello":
        fmt.Println(v1, v1, v1)
    }
    if false {
        if ((128 == 125) || false) {
            v2[v6] = v5
        }
        _ = len(v3)
        v198 := GenS4{F0: (v5 * 42), F1: 0, F2: -2147483648}
        _ = v198
    }
    delete(v3, -1)
    for fi199 := 0; fi199 < 18; fi199++ {
        for fi200 := 0; fi200 < 9; fi200++ {
            if (!(122 != 1)) {
                if false {
                    if (!true) {
                        var v201 uint64 = genFn1()
                        _ = v201
                    } else {
                        switch 32768 {
                        case 127:
                            if true {
                                if false {
                                    v1 = ("foo bar" + v1)
                                    switch int(v5) {
                                    case -2147483648:
                                        v5 = v5
                                    default:
                                        var v202 int8 = 70
                                        _ = v202
                                    }
                                    if false {
                                        v203 := v3[9223372036854775807]
                                        var v204 int32 = -217
                                        var v205 string = (v1 + ("\n\t" + ""))
                                        _ = v203
                                        _ = v204
                                        _ = v205
                                    } else {
                                        var v206 uint16 = 32767
                                        if ((1.5 * ((1.5 * -2.25) + -2.25)) != 0.0) {
                                            if true {
                                                var v207 string = ("a" + v6)
                                                v208 := GenS32{F0: int16(fi199), F1: (^(v5)), F2: (v206 * 796)}
                                                if (v1 != v207) {
                                                    v209 := GenS33{F0: -1}
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    _ = v209
                                                } else {
                                                    var v210 uint8 = 1
                                                    _ = v210
                                                }
                                                _ = v207
                                                _ = v208
                                            } else {
                                                var v211 int8 = 1
                                                var v212 uint32 = 207
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                _ = v211
                                                _ = v212
                                            }
                                            v6 = ((v1 + "a") + (v1 + (v1 + "")))
                                            var v213 uint = 18446744073709551615
                                            var v214 float64 = 3.14159
                                            _ = v213
                                            _ = v214
                                        } else {
                                            v215 := GenS23{F0: (v5 << ((uint(v5)) & 63)), F1: 168, F2: 0}
                                            var v216 float32 = ((1.5 - (((((0.1 - 1e10) / float32((0.1 - 0.0))) + 0.1) / 1e10) + -2.25)) * 1.5)
                                            _ = v215
                                            _ = v216
                                        }
                                        fmt.Println(v5, v1, fi199)
                                        fmt.Println(v5, v6)
                                        _ = v206
                                    }
                                }
                                fmt.Println(v5)
                                v4 = append(v4, 128)
                            } else {
                                v6 = v1
                                v217 := v3[128]
                                switch ((v1 + "") + v6) {
                                case "foo bar":
                                    v4 = append(v4, 127)
                                case "a":
                                    var v218 int8 = 0
                                    _ = v218
                                case "":
                                    var v219 uint16 = 1
                                    _ = v219
                                default:
                                    var v220 bool = false
                                    _ = v220
                                }
                                _ = v217
                            }
                        case 42:
                            v2["hello"] = (v5 + 4294967295)
                        }
                        v221 := GenS17{F0: genFn1(), F1: 1}
                        if ((!true) && (1 == -795)) {
                            var v222 int = -1
                            _ = v222
                        } else {
                            var v223 int = (fi199 + (fi199 | -165))
                            if (42 > 2147483647) {
                                if false {
                                    fmt.Println(v5, v5, v5)
                                    fmt.Println(v223, fi200)
                                    if (!(!(127 == fi199))) {
                                        v5 -= 4294967295
                                        switch v6 {
                                        case "hello":
                                            v5 = (v5 ^ v5)
                                        case "a":
                                            if true {
                                                if (!true) {
                                                    fmt.Println(v6)
                                                    v1 = v6
                                                } else {
                                                    fmt.Println(fi200, fi199)
                                                }
                                                var v224 uint32 = 0
                                                fmt.Println(fi200, fi199, v224)
                                                v5 += (v5 + v5)
                                                fmt.Println(v224)
                                                fmt.Println(v224, v6, fi199)
                                                var v225 float32 = 3.14159
                                                var v226 uint64 = 908
                                                var v227 int32 = int32(v5)
                                                _ = v224
                                                _ = v225
                                                _ = v226
                                                _ = v227
                                            } else {
                                                switch (v223 % (((fi199 ^ fi199)) | 1)) {
                                                case -1:
                                                    fmt.Println(v1, fi200, fi199)
                                                case -2147483648:
                                                    if false {
                                                        v1 = v6
                                                        switch 32768 {
                                                        case 128:
                                                            delete(v2, (((("hello" + "hello") + "") + "") + ("a" + ((v1 + (((v1 + (v6 + (v1 + (((((("\n\t" + v6) + v6) + v6) + "a") + "") + v6)))) + v1) + v1)) + ("hello" + "hello")))))
                                                        }
                                                        if (true || true) {
                                                            if (2 > 1) {
                                                                v6 = ("foo bar" + v1)
                                                                switch (v6 + ((v1 + "\n\t") + v6)) {
                                                                case "a":
                                                                    fmt.Println(v5, v5, v1)
                                                                case "foo bar":
                                                                    if ((0.1 - 3.14159) < 1e10) {
                                                                        if (-32768 == 1) {
                                                                            var v228 uint = 18446744073709551615
                                                                            _ = v228
                                                                        }
                                                                    }
                                                                case "hello":
                                                                    switch "a" {
                                                                    case "hello":
                                                                        delete(v2, ((v1 + "foo bar") + (v6 + v1)))
                                                                    case "a":
                                                                        fmt.Println(v5)
                                                                    case "foo bar":
                                                                        switch v6 {
                                                                        case "a":
                                                                            if true {
                                                                                v6 = ""
                                                                            } else {
                                                                                var v229 int = (v223 % ((489) | 1))
                                                                                v1 = v6
                                                                                _ = v229
                                                                            }
                                                                        }
                                                                    default:
                                                                        v1 = "foo bar"
                                                                    }
                                                                default:
                                                                    if ((false && true) || false) {
                                                                        var v230 string = v6
                                                                        _ = v230
                                                                    }
                                                                }
                                                            } else {
                                                                v231 := v2[("hello" + v6)]
                                                                _ = v231
                                                            }
                                                            v232 := GenS34{F0: 907, F1: "\n\t", F2: 255}
                                                            v2["a"] = (v5 >> ((531) & 63))
                                                            v233 := GenS35{F0: false, F1: 42, F2: 349}
                                                            _ = v232
                                                            _ = v233
                                                        } else {
                                                            delete(v3, (fi200 ^ (fi199 ^ (^(v223)))))
                                                            var v234 int8 = 113
                                                            _ = v234
                                                        }
                                                    } else {
                                                        v4 = append(v4, (fi199 / ((v223) | 1)))
                                                    }
                                                case -359:
                                                    var v235 int16 = 0
                                                    _ = v235
                                                default:
                                                    v1 = v6
                                                }
                                                fmt.Println(v6, v6)
                                            }
                                            if true {
                                                _ = len(v4)
                                                v5 = 0
                                                fmt.Println(v6, v1, v5)
                                            }
                                            fmt.Println(v6, fi199, v6)
                                        default:
                                            var v236 uint = 4294967295
                                            _ = v236
                                        }
                                        var v237 float32 = 3.14159
                                        _ = v237
                                    } else {
                                        _ = len(v3)
                                    }
                                }
                            }
                            _ = v223
                        }
                        if (true && ((!(true || (!false))) || (!(!false)))) {
                            v5 += (v5 << ((65535) & 63))
                            fmt.Println(v6, v1, fi199)
                        }
                        _ = v221
                    }
                    switch fi200 {
                    case -2147483648:
                        if (-128 >= -1) {
                            var v238 int8 = 0
                            _ = v238
                        }
                    case 4294967295:
                        if false {
                            v1 = ""
                            v6 = "foo bar"
                        }
                    }
                }
                v1 = v1
            }
            v6 = (v1 + v1)
            v1 = "\n\t"
            fmt.Println(v1, v5, fi200)
            var v239 uint16 = uint16(v5)
            var v240 int64 = 9223372036854775807
            _ = fi200
            _ = v239
            _ = v240
        }
        var v241 uint32 = 2147483648
        _ = fi199
        _ = v241
    }
    switch "\n\t" {
    case "\n\t":
        switch 0 {
        case 32768:
            var v242 int64 = -1
            _ = v242
        case 1:
            for fi243 := 0; fi243 < 12; fi243++ {
                fmt.Println(v6, v6, v6)
                fmt.Println(v5, v1, v1)
                _ = fi243
            }
        case -756:
            var v244 uint = 42
            _ = v244
        }
    }
    v6 = v1
    var v245 int16 = 42
    for fi246 := 0; fi246 < 8; fi246++ {
        _ = len(v4)
        var v247 uint8 = 1
        for fi248 := 0; fi248 < 5; fi248++ {
            v247 -= v247
            v249 := GenS22{F0: v245, F1: ("a" + "a")}
            switch (v6 + (("\n\t" + (v6 + "hello")) + "")) {
            case "a":
                var v250 uint8 = 255
                _ = v250
            case "":
                switch (fi246 % ((fi248) | 1)) {
                case 32767:
                    var v251 float32 = 0.1
                    v1 = ((v1 + v1) + v6)
                    v252 := GenS3{F0: 2, F1: -1}
                    _ = v251
                    _ = v252
                case -159:
                    var v253 float64 = ((1e10 * 0.1) / float64((1.5 + (0.0 / float32(1e10)))))
                    _ = v253
                }
            }
            _ = fi248
            _ = v249
        }
        _ = fi246
        _ = v247
    }
    v254 := GenS36{F0: 1, F1: 32767, F2: (v245 * v245)}
    if ((32767 < (v245 + -32768)) || true) {
        switch 128 {
        case 32767:
            v5 = v5
        case 127:
            switch 9223372036854775807 {
            case 32767:
                v245 += v245
            case 32768:
                if len(v4) > 0 {
                    v4[0] = -863
                }
            default:
                var v255 uint32 = v5
                _ = v255
            }
        default:
            v245 -= (v245 * (v245 << ((1) & 63)))
        }
        if (1 != 245) {
            fmt.Println(v245)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if true {
                fmt.Println(v5)
                v245 = v245
                v245 = v245
            } else {
                switch ("" + ("" + "foo bar")) {
                case "\n\t":
                    v245 = (v245 << ((65535) & 63))
                case "hello":
                    v6 = "a"
                case "":
                    for fi256 := 0; fi256 < 10; fi256++ {
                        v257 := GenS37{F0: 32767, F1: v6}
                        v1 = (v6 + "\n\t")
                        _ = fi256
                        _ = v257
                    }
                default:
                    v1 = ("a" + ((v1 + v1) + ""))
                }
                switch 42 {
                case 1:
                    v245 += (v245 << ((uint(v5)) & 63))
                case 127:
                    v1 = ""
                case 951:
                    fmt.Println(v245)
                default:
                    v258 := GenS38{F0: -9223372036854775808, F1: int(v245), F2: true}
                    _ = v258
                }
                v4 = append(v4, 982)
            }
        }
    }
    for fi259 := 0; fi259 < 15; fi259++ {
        v4 = append(v4, 9223372036854775807)
        v1 = v6
        _ = fi259
    }
    for fi260 := 0; fi260 < 17; fi260++ {
        var v261 uint = 1
        _ = fi260
        _ = v261
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v245:", v245)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v190
    _ = v197
    _ = v2
    _ = v245
    _ = v254
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
