package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int32; F1 bool; F2 bool }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int8; F1 int64; F2 string }
type GenS12 struct { F0 int8; F1 int; F2 uint }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 int16; F1 uint32; F2 int16 }
type GenS15 struct { F0 string; F1 string }
type GenS16 struct { F0 string }
type GenS17 struct { F0 uint; F1 string; F2 uint }
type GenS18 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS19 struct { F0 int; F1 string; F2 int16 }
type GenS20 struct { F0 int }
type GenS21 struct { F0 uint32 }
type GenS22 struct { F0 int32 }
func genFn0(p0 int64) int8 {
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := GenS1{F0: 0, F1: 4294967295}
        var v8 int16 = -32768
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 1, F1: 42}
    p0 += (-(p0))
    for fi10 := 0; fi10 < 11; fi10++ {
        for fi11 := 0; fi11 < 14; fi11++ {
            switch (-(fi10)) {
            case 128:
                v12 := []int{fi11, 127, (^(fi10))}
                _ = v12
            case 641:
                var v13 string = ("" + ("foo bar" + "\n\t"))
                _ = v13
            }
            _ = fi11
        }
        var v14 string = ("hello" + "")
        _ = fi10
        _ = v14
    }
    v15 := GenS1{F0: 18446744073709551615, F1: 1}
    _ = p0
    _ = v15
    _ = v9
    return 100
}
func genFn1() int16 {
    switch "hello" {
    case "\n\t":
        v16 := []int{0}
        _ = v16
    case "hello":
        if false {
            v17 := GenS2{F0: 127}
            if true {
                if (110 <= -9223372036854775808) {
                    v18 := GenS3{F0: 1}
                    _ = v18
                }
                var v19 float64 = (((3.14159 - (0.0 - 0.1)) / 1.0) * ((0.1 / 1.0) + (-2.25 / 1.0)))
                _ = v19
            }
            var v20 uint32 = 2147483647
            _ = v17
            _ = v20
        }
    case "foo bar":
        v21 := []int{}
        _ = v21
    }
    var v22 float32 = 0.1
    v22 = (v22 - (((0.0 - v22) - v22) / float32(-2.25)))
    var v23 int64 = 350
    if (((32767 > 823) || true) || (127 >= 1)) {
        for fi24 := 0; fi24 < 3; fi24++ {
            for fi25 := 0; fi25 < 11; fi25++ {
                var v26 string = "a"
                if false {
                    switch v26 {
                    case "a":
                        if ((!(!(true && false))) && false) {
                            var v27 int16 = 1000
                            var v28 uint = 83
                            var v29 uint64 = 4294967295
                            _ = v27
                            _ = v28
                            _ = v29
                        } else {
                            if (((3.14159 / (-2.25 * ((-2.25 * (((0.0 / 3.14159) * 1.5) / 3.14159)) * (-2.25 / float64(v22))))) < float64(0.0)) || false) {
                                var v30 float32 = 0.0
                                v23 = -9223372036854775808
                                v31 := GenS4{F0: -251, F1: fi25}
                                switch "\n\t" {
                                case "hello":
                                    v32 := GenS3{F0: -32768}
                                    _ = v32
                                }
                                if false {
                                    v26 = v26
                                    v33 := []int{fi24, fi24, (fi25 - 127), fi25}
                                    _ = v33
                                } else {
                                    var v34 uint64 = 1
                                    var v35 float64 = 0.1
                                    var v36 uint = 298
                                    _ = v34
                                    _ = v35
                                    _ = v36
                                }
                                _ = v30
                                _ = v31
                            } else {
                                var v37 int16 = 0
                                v22 -= 3.14159
                                var v38 uint16 = 2
                                _ = v37
                                _ = v38
                            }
                            if true {
                                var v39 bool = (!false)
                                v40 := GenS5{F0: -2147483648}
                                v41 := GenS5{F0: 1}
                                _ = v39
                                _ = v40
                                _ = v41
                            } else {
                                v42 := GenS6{F0: 255}
                                _ = v42
                            }
                        }
                    default:
                        switch "foo bar" {
                        case "foo bar":
                            v22 = v22
                        case "\n\t":
                            v43 := GenS1{F0: 942, F1: 1}
                            _ = v43
                        case "hello":
                            v44 := []int{(fi25 ^ (-(fi25))), (fi24 + fi24)}
                            _ = v44
                        default:
                            fmt.Println(v23, v22, v26)
                            switch int(v23) {
                            case -2147483648:
                                v45 := []int{42, (fi24 & fi25)}
                                v45 = append(v45, fi25)
                                var v46 int32 = 933
                                _ = v45
                                _ = v46
                            case 9223372036854775807:
                                v47 := []int{-2147483648, fi25, -9223372036854775808, fi25}
                                _ = v47
                            default:
                                if ((!false) && false) {
                                    if false {
                                        var v48 int64 = 42
                                        v23 = 989
                                        switch v26 {
                                        case "a":
                                            var v49 int32 = 100
                                            _ = v49
                                        default:
                                            v50 := []int{128}
                                            _ = v50
                                        }
                                        _ = v48
                                    } else {
                                        var v51 bool = (!(false && (!(!true))))
                                        var v52 uint32 = 42
                                        _ = v51
                                        _ = v52
                                    }
                                    v53 := []int{(^(fi24))}
                                    if (1.5 >= 3.14159) {
                                        if (!false) {
                                            var v54 float32 = (float32(0.1) + v22)
                                            v55 := GenS4{F0: 42, F1: (fi25 / ((32767) | 1))}
                                            _ = v54
                                            _ = v55
                                        }
                                        if len(v53) > 0 {
                                            v53[0] = fi25
                                        }
                                        switch (("\n\t" + v26) + "hello") {
                                        case "foo bar":
                                            fmt.Println(v23, v22)
                                        default:
                                            if false {
                                                v56 := GenS6{F0: 742}
                                                var v57 int = fi24
                                                v57 += fi24
                                                _ = v56
                                                _ = v57
                                            }
                                        }
                                    }
                                    switch v26 {
                                    case "a":
                                        v58 := GenS4{F0: 0, F1: fi24}
                                        _ = v58
                                    default:
                                        if false {
                                            var v59 float32 = (v22 + ((float32(0.0) / ((0.1 + v22) / 1.0)) * v22))
                                            var v60 float64 = (1e10 * float64((float32(1.5) - -2.25)))
                                            v23 -= (v23 + v23)
                                            _ = v59
                                            _ = v60
                                        }
                                        switch (v26 + v26) {
                                        case "":
                                            if false {
                                                v61 := GenS7{F0: "hello"}
                                                var v62 int16 = -32768
                                                var v63 int16 = (v62 & v62)
                                                _ = v61
                                                _ = v62
                                                _ = v63
                                            } else {
                                                var v64 uint32 = 65535
                                                _ = v64
                                            }
                                        }
                                        v65 := GenS8{F0: -1}
                                        _ = v65
                                    }
                                    var v66 float64 = 1e10
                                    if (!true) {
                                        fmt.Println(v23, fi25, fi24)
                                        var v67 uint32 = 955
                                        _ = v67
                                    } else {
                                        var v68 uint64 = 18446744073709551615
                                        fmt.Println(v23, v23, fi24)
                                        var v69 float32 = v22
                                        switch fi25 {
                                        case -646:
                                            var v70 float32 = ((v69 * 1e10) - float32(float64(0.0)))
                                            _ = v70
                                        }
                                        _ = v68
                                        _ = v69
                                    }
                                    if len(v53) > 0 {
                                        v53[0] = (fi24 + 32767)
                                    }
                                    _ = v53
                                    _ = v66
                                } else {
                                    if true {
                                        v22 = 1.5
                                        v71 := []int{}
                                        v72 := GenS9{F0: 42, F1: true, F2: (4294967295 <= 0)}
                                        v22 += 3.14159
                                        var v73 int64 = (v23 << ((80) & 63))
                                        _ = v71
                                        _ = v72
                                        _ = v73
                                    }
                                }
                                if (359 <= 2) {
                                    switch fi24 {
                                    case 128:
                                        switch v26 {
                                        case "hello":
                                            v74 := GenS10{F0: 100}
                                            _ = v74
                                        case "a":
                                            v75 := []int{(fi24 << ((4294967295) & 63)), (fi24 / (((fi24 * (-(fi25)))) | 1)), (fi25 / (((fi24 / ((9223372036854775807) | 1))) | 1))}
                                            var v76 uint16 = 507
                                            v77 := GenS5{F0: -1}
                                            if len(v75) >= 2 {
                                                v75 = v75[1:len(v75)]
                                            }
                                            _ = v75
                                            _ = v76
                                            _ = v77
                                        }
                                    case 4294967295:
                                        var v78 uint64 = 1
                                        _ = v78
                                    }
                                    v79 := []int{fi25, -1}
                                    _ = v79
                                } else {
                                    fmt.Println(v23, fi24, fi24)
                                    if (0.1 != 3.14159) {
                                        var v80 float64 = 3.14159
                                        _ = v80
                                    }
                                    v22 += 1e10
                                }
                                var v81 string = ("hello" + "foo bar")
                                _ = v81
                            }
                            if ((42 == 1) && (-1 >= 2147483647)) {
                                v22 += (1e10 - (0.0 / 1e10))
                            } else {
                                switch 0 {
                                case -215:
                                    var v82 int = 0
                                    _ = v82
                                case -491:
                                    v83 := []int{fi24}
                                    _ = v83
                                default:
                                    switch fi24 {
                                    case 9223372036854775807:
                                        v84 := []int{198}
                                        _ = v84
                                    case 127:
                                        switch "hello" {
                                        case "":
                                            v85 := GenS6{F0: 18446744073709551615}
                                            _ = v85
                                        }
                                    case 2147483647:
                                        if false {
                                            v86 := GenS3{F0: -899}
                                            fmt.Println(v26)
                                            v87 := []int{(^(fi24)), fi25, fi25}
                                            _ = v86
                                            _ = v87
                                        }
                                        switch v26 {
                                        case "":
                                            if true {
                                                v88 := []int{fi24, -266, (fi25 | (fi24 + fi24)), fi24}
                                                v26 = (((v26 + v26) + (((v26 + (("a" + v26) + (v26 + v26))) + (v26 + ("hello" + (("\n\t" + v26) + "")))) + "a")) + ("a" + v26))
                                                switch "\n\t" {
                                                case "foo bar":
                                                    if len(v88) > 0 {
                                                        v88[0] = (fi25 & 127)
                                                    }
                                                case "\n\t":
                                                    v88 = append(v88, fi25)
                                                    var v89 uint32 = 2
                                                    _ = v89
                                                case "":
                                                    v23 += (^(v23))
                                                default:
                                                    if (false || ((0.1 * (float32(3.14159) / v22)) <= 0.0)) {
                                                        var v90 float32 = float32(1e10)
                                                        switch fi25 {
                                                        case -1:
                                                            if len(v88) >= 2 {
                                                                v88 = v88[1:len(v88)]
                                                            }
                                                        case 9223372036854775807:
                                                            if (0 != 255) {
                                                                var v91 float64 = (0.1 / 1.5)
                                                                v91 = (v91 / float64(1e10))
                                                                _ = v91
                                                            }
                                                        default:
                                                            v92 := GenS11{F0: 100, F1: (^(v23)), F2: "foo bar"}
                                                            _ = v92
                                                        }
                                                        v23 = v23
                                                        if len(v88) >= 2 {
                                                            v88 = v88[1:len(v88)]
                                                        }
                                                        v93 := GenS5{F0: -1}
                                                        v94 := GenS10{F0: 0}
                                                        _ = v90
                                                        _ = v93
                                                        _ = v94
                                                    }
                                                    var v95 int8 = 42
                                                    if true {
                                                        if len(v88) > 0 {
                                                            v88[0] = 2147483647
                                                        }
                                                        v88 = append(v88, (fi25 >> ((247) & 63)))
                                                        var v96 int8 = 127
                                                        _ = v96
                                                    } else {
                                                        if len(v88) >= 2 {
                                                            v88 = v88[1:len(v88)]
                                                        }
                                                        switch ("a" + ("hello" + (v26 + "foo bar"))) {
                                                        case "a":
                                                            v22 -= ((v22 + (v22 * ((3.14159 + v22) - v22))) + 0.0)
                                                        case "\n\t":
                                                            v23 += v23
                                                        case "foo bar":
                                                            v95 = (-(v95))
                                                        }
                                                        if (830 < -2147483648) {
                                                            var v97 bool = true
                                                            _ = v97
                                                        } else {
                                                            switch (fi24 - fi25) {
                                                            case 4294967295:
                                                                v95 = v95
                                                            default:
                                                                var v98 bool = (((-2.25 / (1e10 + ((3.14159 - -2.25) - float64(3.14159)))) + 1.5) < -2.25)
                                                                _ = v98
                                                            }
                                                            if (!(!(!false))) {
                                                                var v99 uint = 795
                                                                v100 := GenS12{F0: v95, F1: fi25, F2: 4294967295}
                                                                v101 := GenS13{F0: (v23 & 1)}
                                                                _ = v100
                                                                _ = v101
                                                                _ = v99
                                                            } else {
                                                                var v102 uint32 = 2147483647
                                                                _ = v102
                                                            }
                                                        }
                                                        v23 = 186
                                                        v95 = genFn0(v23)
                                                    }
                                                    fmt.Println(v22, v95, fi25)
                                                    v95 -= genFn0(v23)
                                                    _ = v95
                                                }
                                                _ = v88
                                            }
                                        default:
                                            var v103 string = ((v26 + ("" + v26)) + v26)
                                            if (v26 == (((v26 + "hello") + "") + "")) {
                                                v104 := GenS11{F0: 127, F1: -178, F2: "\n\t"}
                                                var v105 string = v103
                                                _ = v104
                                                _ = v105
                                            } else {
                                                v106 := []int{fi24, 249}
                                                _ = v106
                                            }
                                            switch "a" {
                                            case "a":
                                                fmt.Println(v23, v26)
                                            }
                                            _ = v103
                                        }
                                        v107 := GenS14{F0: 32767, F1: 1, F2: -32768}
                                        _ = v107
                                    }
                                }
                                var v108 uint64 = 155
                                v109 := GenS8{F0: -32768}
                                fmt.Println(v22, v23)
                                v110 := []int{}
                                _ = v108
                                _ = v109
                                _ = v110
                            }
                            fmt.Println(v26, fi24)
                            switch (fi25 >> ((972) & 63)) {
                            case 378:
                                v23 = (-(v23))
                            case 4294967295:
                                switch ("foo bar" + ((v26 + v26) + "foo bar")) {
                                case "a":
                                    fmt.Println(fi25, v22)
                                case "hello":
                                    v23 -= (v23 * v23)
                                case "\n\t":
                                    if false {
                                        var v111 uint8 = 127
                                        var v112 uint = 650
                                        switch v26 {
                                        case "":
                                            v113 := GenS15{F0: v26, F1: v26}
                                            _ = v113
                                        default:
                                            fmt.Println(v23, v22, v111)
                                        }
                                        _ = v111
                                        _ = v112
                                    }
                                }
                            default:
                                if (1 != fi25) {
                                    var v114 uint = 255
                                    v114 -= 398
                                    _ = v114
                                } else {
                                    var v115 uint64 = 885
                                    v116 := []int{(fi25 >> ((4294967295) & 63))}
                                    _ = v115
                                    _ = v116
                                }
                            }
                            var v117 int16 = 1
                            _ = v117
                        }
                    }
                }
                fmt.Println(fi24, v22, fi24)
                fmt.Println(v26, v26, v26)
                fmt.Println(v23, fi24, fi24)
                v26 = (v26 + "")
                v118 := []int{fi25, 42, (fi25 >> ((976) & 63))}
                _ = fi25
                _ = v118
                _ = v26
            }
            if ((!true) && ((fi24 | fi24) >= (fi24 & 32767))) {
                var v119 float32 = ((v22 * v22) / 1e10)
                if (!(0 != 2147483647)) {
                    v120 := GenS16{F0: ""}
                    _ = v120
                } else {
                    fmt.Println(v23)
                }
                _ = v119
            } else {
                switch fi24 {
                case -9223372036854775808:
                    var v121 bool = true
                    _ = v121
                case 383:
                    fmt.Println(fi24)
                case -677:
                    for fi122 := 0; fi122 < 6; fi122++ {
                        switch ("hello" + ("hello" + ("\n\t" + ("" + "hello")))) {
                        case "hello":
                            v123 := []int{fi24, 1}
                            _ = v123
                        default:
                            v23 = 0
                        }
                        v23 -= v23
                        if (18446744073709551615 <= 560) {
                            v23 = (v23 / (((v23 ^ v23)) | 1))
                            v124 := GenS17{F0: 2, F1: "foo bar", F2: 812}
                            v125 := GenS18{F0: 754, F1: 2, F2: 2147483648}
                            _ = v124
                            _ = v125
                        } else {
                            v126 := GenS19{F0: fi122, F1: "hello", F2: int16(v23)}
                            if false {
                                switch ("a" + "\n\t") {
                                case "\n\t":
                                    var v127 float32 = -2.25
                                    _ = v127
                                }
                                fmt.Println(v23)
                                v128 := []int{}
                                _ = v128
                            } else {
                                var v129 bool = (4294967295 <= 1)
                                switch 4294967295 {
                                case 35:
                                    if v129 {
                                        v130 := []int{-887, (fi122 ^ 128), fi24, (fi24 ^ (-(fi122)))}
                                        _ = v130
                                    }
                                case 91:
                                    if (-606 < 2147483647) {
                                        var v131 float32 = float32((-2.25 / 1.0))
                                        v132 := []int{(^(fi122))}
                                        var v133 float64 = (-2.25 * (3.14159 + 1.5))
                                        _ = v131
                                        _ = v132
                                        _ = v133
                                    }
                                case -317:
                                    v23 -= (v23 ^ v23)
                                default:
                                    v134 := GenS20{F0: (fi122 << ((255) & 63))}
                                    _ = v134
                                }
                                if v129 {
                                    if v129 {
                                        v135 := []int{}
                                        v23 = -9223372036854775808
                                        v129 = (v129 || v129)
                                        _ = v135
                                    }
                                    v136 := []int{fi24}
                                    switch "\n\t" {
                                    case "":
                                        v22 = v22
                                        if v129 {
                                            v129 = (32767 >= 1)
                                        }
                                        if (v129 && v129) {
                                            fmt.Println(v129)
                                            var v137 int8 = 1
                                            _ = v137
                                        } else {
                                            var v138 bool = v129
                                            v136 = append(v136, fi122)
                                            fmt.Println(fi122, fi122)
                                            _ = v138
                                        }
                                        var v139 uint = 65535
                                        _ = v139
                                    }
                                    _ = v136
                                }
                                _ = v129
                            }
                            v140 := []int{fi24, (fi24 & 128), fi122}
                            _ = v126
                            _ = v140
                        }
                        _ = fi122
                    }
                }
                switch (fi24 >> ((2) & 63)) {
                case -2147483648:
                    var v141 int8 = 1
                    _ = v141
                case 32768:
                    var v142 int64 = (v23 / ((v23) | 1))
                    _ = v142
                default:
                    v143 := GenS5{F0: 688}
                    _ = v143
                }
            }
            v144 := []int{-292, fi24, 42}
            _ = fi24
            _ = v144
        }
        v23 = (v23 % (((^(v23))) | 1))
    }
    _ = v22
    _ = v23
    return -32768
}

func main() {
    v1 := []int{}
    var v2 uint8 = 255
    var v3 string = ("" + "hello")
    v4 := make(map[int]int)
    var v5 uint = 15
    fmt.Println(v5, v2, v2)
    var v145 int32 = 2147483647
    switch "a" {
    case "foo bar":
        var v146 uint64 = 4294967295
        _ = v146
    case "\n\t":
        switch 9223372036854775807 {
        case -2147483648:
            if (!(307 < 1)) {
                for fi147 := 0; fi147 < 20; fi147++ {
                    var v148 bool = ((-1 == 32767) || ((((-2.25 / 1.0) - 1e10) - (1.5 / ((float32(1e10) - ((0.0 + float32(((0.1 / (1.5 + 0.1)) + 1.5))) + -2.25)) * float32(-2.25)))) <= 1e10))
                    for fi149 := 0; fi149 < 15; fi149++ {
                        v145 += (v145 ^ v145)
                        var v150 float32 = float32((1.5 * float64(-2.25)))
                        v151 := GenS8{F0: 1}
                        _ = fi149
                        _ = v150
                        _ = v151
                    }
                    if ((!true) && v148) {
                        var v152 int16 = 0
                        v3 = ((v3 + ("foo bar" + v3)) + (v3 + ""))
                        var v153 int32 = -762
                        _ = v152
                        _ = v153
                    } else {
                        _ = len(v4)
                        for fi154 := 0; fi154 < 11; fi154++ {
                            v148 = (((v145 << ((42) & 63)) != -1) || v148)
                            _ = fi154
                        }
                        var v155 int16 = 0
                        _ = v155
                    }
                    _ = fi147
                    _ = v148
                }
            } else {
                _ = len(v4)
                if ((0 <= 825) && false) {
                    if (!false) {
                        v156 := GenS9{F0: 1, F1: false, F2: (v145 < 705)}
                        var v157 float64 = 1e10
                        var v158 int64 = 0
                        _ = v156
                        _ = v157
                        _ = v158
                    }
                    v159 := GenS21{F0: 0}
                    if ("foo bar" < v3) {
                        v5 = v5
                        for fi160 := 0; fi160 < 17; fi160++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            for fi161 := 0; fi161 < 9; fi161++ {
                                switch "\n\t" {
                                case "hello":
                                    var v162 int16 = -930
                                    _ = v162
                                case "":
                                    fmt.Println(v2, fi161, v3)
                                default:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                v3 = v3
                                v5 = (v5 << ((2) & 63))
                                _ = fi161
                            }
                            v163 := GenS12{F0: -29, F1: (fi160 ^ (fi160 | (^(fi160)))), F2: v5}
                            _ = fi160
                            _ = v163
                        }
                        var v164 int8 = 100
                        _ = v164
                    }
                    _ = v159
                } else {
                    v1 = append(v1, -1)
                }
                if len(v1) > 0 {
                    v1[0] = -808
                }
            }
        default:
            v4[-1] = -1
        }
    }
    for fi165 := 0; fi165 < 15; fi165++ {
        v3 = v3
        var v166 uint64 = 18446744073709551615
        _ = fi165
        _ = v166
    }
    v5 -= v5
    v167 := v4[-1]
    v167 = v167
    fmt.Println(v3, v145, v5)
    v5 = uint(v2)
    delete(v4, int(v5))
    v4[(v167 - 42)] = v167
    if false {
        if true {
            var v168 uint64 = 41
            var v169 int = v167
            _ = v168
            _ = v169
        }
    } else {
        for fi170 := 0; fi170 < 1; fi170++ {
            if (!(uint32(v2) >= 2147483648)) {
                v5 = (v5 & 65535)
            }
            _ = fi170
        }
        v171 := GenS22{F0: -827}
        _ = v171
    }
    fmt.Println(v145, v145)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v145:", v145)
    fmt.Println("v167:", v167)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v145
    _ = v167
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
