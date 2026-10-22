package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 int32; F2 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int64; F1 uint }
type GenS11 struct { F0 uint; F1 int64 }
type GenS12 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS13 struct { F0 uint32; F1 uint16; F2 uint16 }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 bool; F1 int16; F2 int16 }
type GenS16 struct { F0 uint32 }
type GenS17 struct { F0 uint16; F1 string }
type GenS18 struct { F0 uint; F1 int16; F2 int16 }
type GenS19 struct { F0 int64; F1 uint16 }
type GenS20 struct { F0 uint16; F1 int32 }
type GenS21 struct { F0 string; F1 int }
type GenS22 struct { F0 uint16 }
type GenS23 struct { F0 uint32; F1 string }
type GenS24 struct { F0 uint64 }
type GenS25 struct { F0 int8 }
type GenS26 struct { F0 int32; F1 uint64; F2 int64 }
func genFn0() uint16 {
    v5 := []int{}
    var v6 uint16 = 1
    if (317 < 792) {
        for fi7 := 0; fi7 < 9; fi7++ {
            for fi8 := 0; fi8 < 4; fi8++ {
                v9 := GenS1{F0: 0}
                _ = fi8
                _ = v9
            }
            _ = len(v5)
            _ = fi7
        }
        v6 -= v6
    }
    v6 -= v6
    _ = v5
    _ = v6
    return v6
}

func main() {
    var v1 uint64 = 2
    var v2 int16 = 1000
    var v3 float64 = float64((1.5 * (((0.1 + 0.0) * 3.14159) - float32(0.0))))
    v4 := make(map[uint32]uint32)
    if true {
        if false {
            switch 422 {
            case 127:
                if ((v3 != v3) && (!false)) {
                    if false {
                        switch "hello" {
                        case "a":
                            if (18446744073709551615 < 1) {
                                for fi10 := 0; fi10 < 20; fi10++ {
                                    fmt.Println(fi10, v1, fi10)
                                    if (!false) {
                                        v1 = v1
                                    } else {
                                        var v11 int = 9223372036854775807
                                        v12 := GenS1{F0: 904}
                                        _ = v11
                                        _ = v12
                                    }
                                    v13 := []int{}
                                    _ = fi10
                                    _ = v13
                                }
                                v14 := []int{1, -478, 128, 2147483647}
                                v2 = 109
                                _ = v14
                            }
                        case "hello":
                            if false {
                                _ = len(v4)
                            }
                        case "foo bar":
                            delete(v4, 281)
                        }
                    } else {
                        var v15 int8 = -128
                        if (127 == v15) {
                            v16 := []int{}
                            _ = v16
                        }
                        _ = v15
                    }
                    v3 += 0.0
                    var v17 string = "hello"
                    _ = v17
                } else {
                    var v18 int32 = -334
                    for fi19 := 0; fi19 < 3; fi19++ {
                        for fi20 := 0; fi20 < 16; fi20++ {
                            fmt.Println(v2, v2, fi20)
                            switch (("hello" + "\n\t") + ("\n\t" + "\n\t")) {
                            case "foo bar":
                                _ = len(v4)
                            case "\n\t":
                                var v21 int64 = 42
                                _ = v21
                            case "":
                                if ((!false) && false) {
                                    if true {
                                        v22 := GenS2{F0: (!false), F1: v18, F2: "hello"}
                                        var v23 int16 = (^(v2))
                                        var v24 int8 = 42
                                        _ = v22
                                        _ = v23
                                        _ = v24
                                    } else {
                                        v25 := GenS3{F0: 1}
                                        v26 := []int{fi19, 304, -862, (fi20 % ((128) | 1))}
                                        var v27 int8 = 101
                                        _ = v25
                                        _ = v26
                                        _ = v27
                                    }
                                    v2 = (v2 | -742)
                                    if (2 >= 921) {
                                        if false {
                                            switch fi20 {
                                            case 542:
                                                var v28 uint8 = 128
                                                _ = v28
                                            case -516:
                                                var v29 int16 = 32767
                                                _ = v29
                                            default:
                                                if (int32(v1) >= v18) {
                                                    v30 := []int{fi19, (fi20 << ((255) & 63))}
                                                    var v31 float64 = 3.14159
                                                    _ = v30
                                                    _ = v31
                                                }
                                            }
                                            var v32 uint8 = 127
                                            _ = v32
                                        }
                                        if (!true) {
                                            if false {
                                                v33 := []int{}
                                                _ = v33
                                            }
                                            var v34 uint = 708
                                            _ = v34
                                        }
                                    } else {
                                        var v35 uint = 995
                                        if true {
                                            fmt.Println(fi20, v3)
                                            fmt.Println(v18, v2)
                                            switch fi19 {
                                            case 405:
                                                fmt.Println(v1, v35, fi19)
                                            case 32768:
                                                var v36 uint8 = 42
                                                _ = v36
                                            case -9223372036854775808:
                                                var v37 string = "hello"
                                                _ = v37
                                            }
                                        }
                                        if true {
                                            if false {
                                                var v38 bool = (((v35 / ((2) | 1)) > 104) && (true || ((float64(0.0) / float64((float32((0.1 - 1.5)) / 1.5))) >= 1e10)))
                                                _ = v38
                                            }
                                            v39 := GenS4{F0: (float32(1e10) < ((1e10 + 3.14159) - ((0.0 / float32(-2.25)) - (3.14159 / 3.14159)))), F1: ""}
                                            var v40 string = "foo bar"
                                            _ = v39
                                            _ = v40
                                        } else {
                                            var v41 int32 = (v18 ^ (v18 ^ v18))
                                            v1 = v1
                                            if (!true) {
                                                v42 := []int{fi20, 9223372036854775807, (fi19 * (fi20 >> (((^(v35))) & 63)))}
                                                _ = v42
                                            }
                                            var v43 bool = (false && false)
                                            _ = v41
                                            _ = v43
                                        }
                                        v44 := GenS5{F0: (v1 - (^(v1))), F1: (^(v2))}
                                        if (1e10 == ((float32(3.14159) / -2.25) - (1e10 + 1e10))) {
                                            fmt.Println(v35, v18, v2)
                                            v45 := GenS1{F0: (v35 ^ v35)}
                                            _ = v45
                                        } else {
                                            v35 += v35
                                            if false {
                                                fmt.Println(fi20, fi20, v2)
                                            } else {
                                                var v46 int64 = 0
                                                fmt.Println(v35)
                                                _ = v46
                                            }
                                        }
                                        _ = v35
                                        _ = v44
                                    }
                                } else {
                                    var v47 uint = 42
                                    v18 -= 0
                                    if false {
                                        var v48 float64 = 0.1
                                        if false {
                                            if (true || (v1 < (v1 / (((v1 - v1)) | 1)))) {
                                                var v49 float64 = 3.14159
                                                v3 = 1.5
                                                _ = v49
                                            } else {
                                                if (0.1 <= (-2.25 - (1e10 - -2.25))) {
                                                    switch fi20 {
                                                    case 0:
                                                        switch (fi19 | fi19) {
                                                        case -9223372036854775808:
                                                            if true {
                                                                var v50 int64 = -691
                                                                v18 += (v18 * 42)
                                                                _ = v50
                                                            }
                                                        case -353:
                                                            var v51 int8 = 87
                                                            _ = v51
                                                        default:
                                                            v52 := []int{}
                                                            _ = v52
                                                        }
                                                    }
                                                } else {
                                                    v2 += -32768
                                                    v53 := GenS3{F0: 166}
                                                    if false {
                                                        v48 = ((v48 - -2.25) - float64(0.1))
                                                    } else {
                                                        v54 := GenS6{F0: 1}
                                                        delete(v4, 42)
                                                        _ = v54
                                                    }
                                                    v55 := v4[325]
                                                    switch -194 {
                                                    case -582:
                                                        v56 := GenS7{F0: (v55 >> ((v47) & 63))}
                                                        _ = v56
                                                    case -2147483648:
                                                        if false {
                                                            v57 := GenS3{F0: 1}
                                                            _ = v57
                                                        }
                                                    case 4294967295:
                                                        var v58 uint32 = (^(v55))
                                                        _ = v58
                                                    }
                                                    _ = v53
                                                    _ = v55
                                                }
                                                v1 += (v1 - 4294967295)
                                            }
                                            if true {
                                                v47 = 18446744073709551615
                                                v2 -= 32767
                                                v48 = (3.14159 / v48)
                                            }
                                        }
                                        if (-1 != int8(fi20)) {
                                            v3 = (1e10 * v48)
                                            switch (fi20 + (fi19 >> ((v47) & 63))) {
                                            case -1:
                                                var v59 int16 = (v2 >> (((v47 & (v47 >> ((v47) & 63)))) & 63))
                                                _ = v59
                                            case 2147483647:
                                                v3 -= (((float64(3.14159) + v3) / 1e10) / v3)
                                            default:
                                                if true {
                                                    v60 := GenS1{F0: 608}
                                                    v61 := []int{}
                                                    _ = v60
                                                    _ = v61
                                                }
                                            }
                                            v62 := GenS8{F0: (fi20 ^ -9223372036854775808)}
                                            _ = v62
                                        } else {
                                            fmt.Println(v47)
                                            v63 := []int{(fi20 << ((800) & 63))}
                                            _ = v63
                                        }
                                        _ = v48
                                    } else {
                                        v2 += v2
                                    }
                                    switch fi19 {
                                    case 1:
                                        if (0 == 23) {
                                            v2 = v2
                                            var v64 uint32 = 1
                                            _ = v64
                                        } else {
                                            v1 += v1
                                        }
                                    case -536:
                                        _ = len(v4)
                                    }
                                    switch "a" {
                                    case "foo bar":
                                        fmt.Println(v47)
                                    case "\n\t":
                                        fmt.Println(v18, v18)
                                    default:
                                        v65 := GenS9{F0: v47}
                                        _ = v65
                                    }
                                    fmt.Println(v1, fi19)
                                    v66 := GenS2{F0: false, F1: (v18 - -355), F2: ("foo bar" + "foo bar")}
                                    _ = v47
                                    _ = v66
                                }
                            }
                            _ = fi20
                        }
                        for fi67 := 0; fi67 < 20; fi67++ {
                            v2 += v2
                            if (127 >= -49) {
                                var v68 uint8 = 45
                                var v69 int32 = v18
                                v18 = v18
                                _ = v68
                                _ = v69
                            } else {
                                v70 := GenS10{F0: 1, F1: 2}
                                switch "foo bar" {
                                case "a":
                                    v1 = v1
                                case "hello":
                                    switch (fi19 << ((42) & 63)) {
                                    case 32768:
                                        if false {
                                            v71 := GenS7{F0: 4294967295}
                                            _ = v71
                                        } else {
                                            var v72 string = ""
                                            v18 -= (v18 + (v18 & (v18 / ((v18) | 1))))
                                            v3 = ((v3 - 0.1) / 3.14159)
                                            var v73 float32 = (float32(v3) + 0.0)
                                            v73 += v73
                                            switch 4294967295 {
                                            case -9223372036854775808:
                                                v74 := v4[4294967295]
                                                _ = v74
                                            }
                                            v75 := GenS11{F0: 2, F1: -9223372036854775808}
                                            v3 = v3
                                            var v76 int32 = v18
                                            _ = v72
                                            _ = v73
                                            _ = v75
                                            _ = v76
                                        }
                                    case 619:
                                        if false {
                                            v77 := []int{fi19}
                                            var v78 float32 = 1e10
                                            if true {
                                                switch "" {
                                                case "a":
                                                    var v79 uint8 = 2
                                                    _ = v79
                                                case "hello":
                                                    fmt.Println(v78, v2, fi67)
                                                }
                                            } else {
                                                var v80 uint64 = v1
                                                if (2147483648 == 919) {
                                                    v80 = 42
                                                    v81 := GenS1{F0: 65535}
                                                    _ = v81
                                                }
                                                _ = v80
                                            }
                                            _ = v77
                                            _ = v78
                                        } else {
                                            v1 += v1
                                            var v82 uint16 = genFn0()
                                            v83 := GenS2{F0: true, F1: 0, F2: "\n\t"}
                                            _ = v82
                                            _ = v83
                                        }
                                    case -1:
                                        v84 := v4[0]
                                        fmt.Println(v2)
                                        _ = v84
                                    default:
                                        if true {
                                            v85 := GenS12{F0: v1, F1: 4294967295, F2: 190}
                                            _ = v85
                                        }
                                    }
                                }
                                var v86 int8 = 1
                                _ = v70
                                _ = v86
                            }
                            v1 = v1
                            v87 := GenS13{F0: 956, F1: 0, F2: 0}
                            v4[0] = 42
                            _ = fi67
                            _ = v87
                        }
                        _ = fi19
                    }
                    _ = v18
                }
            default:
                v88 := v4[0]
                _ = v88
            }
        } else {
            if true {
                v1 -= v1
                v89 := GenS3{F0: 42}
                v1 -= 18446744073709551615
                _ = v89
            } else {
                var v90 int16 = v2
                fmt.Println(v2)
                _ = v90
            }
        }
        var v91 bool = (!(true && (3.14159 > 1e10)))
        _ = v91
    }
    fmt.Println(v2, v3, v3)
    for fi92 := 0; fi92 < 9; fi92++ {
        var v93 float32 = -2.25
        v2 = v2
        v2 = (v2 >> ((4294967295) & 63))
        _ = fi92
        _ = v93
    }
    if (1 != 42) {
        v94 := v4[238]
        _ = v94
    } else {
        switch "\n\t" {
        case "hello":
            var v95 uint64 = (v1 + 18446744073709551615)
            _ = v95
        case "\n\t":
            var v96 float32 = 3.14159
            _ = v96
        case "":
            v97 := GenS14{F0: 454}
            _ = v97
        default:
            v98 := GenS15{F0: true, F1: v2, F2: (v2 ^ 1)}
            _ = v98
        }
        v3 += (v3 - -2.25)
        if true {
            v3 = v3
        } else {
            v99 := GenS3{F0: 128}
            _ = v99
        }
    }
    v4[65535] = 2
    var v100 uint = 1
    var v101 float64 = v3
    for fi102 := 0; fi102 < 6; fi102++ {
        v103 := GenS16{F0: 930}
        var v104 float32 = 0.0
        _ = fi102
        _ = v103
        _ = v104
    }
    switch "foo bar" {
    case "foo bar":
        var v105 int8 = -10
        _ = v105
    case "a":
        if true {
            for fi106 := 0; fi106 < 9; fi106++ {
                v107 := GenS17{F0: genFn0(), F1: "a"}
                _ = fi106
                _ = v107
            }
        }
    default:
        if (-1 != 0) {
            for fi108 := 0; fi108 < 12; fi108++ {
                v109 := []int{}
                fmt.Println(fi108, v3, v100)
                for fi110 := 0; fi110 < 4; fi110++ {
                    v111 := GenS7{F0: 2}
                    var v112 uint16 = 32767
                    _ = fi110
                    _ = v111
                    _ = v112
                }
                _ = fi108
                _ = v109
            }
            if (255 < 1) {
                var v113 uint64 = v1
                _ = v113
            } else {
                if false {
                    switch -487 {
                    case 9223372036854775807:
                        v2 = v2
                    case 128:
                        delete(v4, 42)
                    default:
                        fmt.Println(v100, v101, v101)
                    }
                }
                v114 := GenS8{F0: -9223372036854775808}
                if true {
                    if (-1 != -9223372036854775808) {
                        if (3.14159 <= 0.0) {
                            v2 = v2
                            var v115 uint32 = 42
                            v116 := GenS18{F0: v100, F1: 42, F2: v2}
                            _ = v115
                            _ = v116
                        }
                        v117 := []int{int(v100)}
                        v118 := GenS19{F0: 9223372036854775807, F1: 1}
                        _ = v117
                        _ = v118
                    } else {
                        if false {
                            fmt.Println(v2, v2, v100)
                        }
                        switch ("" + ("" + "")) {
                        case "hello":
                            var v119 uint32 = 1
                            _ = v119
                        case "a":
                            _ = len(v4)
                        case "foo bar":
                            v1 = v1
                        }
                    }
                }
                _ = v114
            }
        } else {
            v2 += v2
            var v120 uint64 = v1
            _ = v120
        }
    }
    var v121 int8 = -12
    v1 = (v1 * v1)
    v1 += (v1 ^ v1)
    if (v1 < (v1 * 4294967295)) {
        var v122 uint64 = (v1 << ((v100) & 63))
        _ = v122
    } else {
        var v123 uint16 = 2
        _ = v123
    }
    switch "" {
    case "a":
        switch ("" + "a") {
        case "foo bar":
            _ = len(v4)
        case "a":
            v2 = -797
        case "":
            var v124 int64 = -1
            _ = v124
        }
    case "":
        fmt.Println(v100, v2, v2)
    case "foo bar":
        var v125 int8 = -8
        _ = v125
    default:
        var v126 int16 = int16(v1)
        _ = v126
    }
    if ((v100 == 855) || false) {
        if true {
            if (true || true) {
                v127 := GenS20{F0: 324, F1: 42}
                var v128 int64 = 0
                v1 = v1
                _ = v127
                _ = v128
            } else {
                switch 127 {
                case -735:
                    fmt.Println(v3, v100, v2)
                default:
                    switch "" {
                    case "a":
                        v129 := GenS21{F0: "", F1: int(v121)}
                        _ = v129
                    case "hello":
                        var v130 int = -398
                        _ = v130
                    case "":
                        delete(v4, 0)
                    default:
                        v1 = (v1 ^ v1)
                    }
                }
            }
            v100 -= (^(v100))
            if ((1.5 > float64((float32(v101) / 1.0))) && true) {
                for fi131 := 0; fi131 < 6; fi131++ {
                    var v132 uint64 = (^(v1))
                    delete(v4, 0)
                    switch fi131 {
                    case 683:
                        for fi133 := 0; fi133 < 6; fi133++ {
                            v101 += ((((v101 / float64(float32(v3))) - v101) / ((-2.25 - float64(0.0)) * v3)) / 3.14159)
                            var v134 int8 = (v121 / ((53) | 1))
                            switch fi133 {
                            case 0:
                                var v135 uint64 = 42
                                _ = v135
                            }
                            _ = fi133
                            _ = v134
                        }
                    case 42:
                        v136 := GenS8{F0: fi131}
                        _ = v136
                    }
                    _ = fi131
                    _ = v132
                }
                v137 := GenS22{F0: genFn0()}
                if false {
                    if (((0.1 * (1e10 + 0.1)) >= float32(((float64((3.14159 / 0.1)) / v3) + v3))) && true) {
                        v100 = (v100 | (v100 | v100))
                        if (("foo bar" + "a") > "") {
                            v138 := v4[1]
                            if (42 != 42) {
                                v138 -= v138
                                var v139 int64 = 42
                                var v140 int = 127
                                _ = v139
                                _ = v140
                            } else {
                                var v141 uint = (^(v100))
                                switch 4294967295 {
                                case -981:
                                    v121 = (v121 << (((v141 ^ (v100 | (v141 << ((552) & 63))))) & 63))
                                case 250:
                                    var v142 int32 = 42
                                    _ = v142
                                case -333:
                                    v143 := GenS7{F0: (v138 >> ((v141) & 63))}
                                    _ = v143
                                default:
                                    var v144 string = ""
                                    _ = v144
                                }
                                _ = v141
                            }
                            _ = v138
                        } else {
                            if false {
                                var v145 int32 = -349
                                if false {
                                    for fi146 := 0; fi146 < 8; fi146++ {
                                        var v147 int = fi146
                                        fmt.Println(v1, v101)
                                        v148 := GenS20{F0: 42, F1: (v145 | (v145 ^ 519))}
                                        _ = fi146
                                        _ = v147
                                        _ = v148
                                    }
                                    switch 1 {
                                    case 0:
                                        v149 := GenS23{F0: 0, F1: "hello"}
                                        _ = v149
                                    case 9223372036854775807:
                                        v150 := v4[uint32(v100)]
                                        _ = v150
                                    case 606:
                                        v151 := []int{2147483647, -346, 4294967295, 127}
                                        _ = v151
                                    default:
                                        v121 -= (v121 + 1)
                                    }
                                }
                                _ = v145
                            }
                        }
                        if ((3.14159 + (float32((float64(1e10) - (v3 * ((v3 * v3) / 0.1)))) - (((((((1.5 + (3.14159 - 1.5)) + 1.5) * 1e10) / (-2.25 * (-2.25 - (0.0 - float32(((-2.25 * v101) * 1.5)))))) / 3.14159) + 1e10) - ((0.1 - 0.0) / 1.5)))) > 1.5) {
                            for fi152 := 0; fi152 < 8; fi152++ {
                                for fi153 := 0; fi153 < 13; fi153++ {
                                    if ((!((255 <= 0) || ("\n\t" > "hello"))) && (-139 > -2147483648)) {
                                        v3 = v101
                                        v154 := GenS24{F0: 42}
                                        _ = v154
                                    }
                                    switch "a" {
                                    case "":
                                        if (!false) {
                                            v3 -= (v3 - 1e10)
                                        }
                                        v155 := []int{}
                                        _ = v155
                                    default:
                                        v156 := GenS16{F0: 1}
                                        _ = v156
                                    }
                                    if ((3.14159 - -2.25) > (1e10 - -2.25)) {
                                        v157 := []int{}
                                        _ = v157
                                    }
                                    _ = fi153
                                }
                                var v158 uint32 = 2
                                v100 = (v100 | uint(v1))
                                _ = fi152
                                _ = v158
                            }
                            if ((1e10 <= -2.25) || ((v2 >> ((0) & 63)) == (v2 / (((v2 ^ (v2 << (((v100 + v100)) & 63)))) | 1)))) {
                                if true {
                                    v100 += (v100 / (((^(v100))) | 1))
                                    var v159 uint8 = 1
                                    _ = v159
                                } else {
                                    v101 = v101
                                    var v160 float64 = v3
                                    _ = v160
                                }
                                v3 = v3
                                switch ((("a" + "a") + "\n\t") + "") {
                                case "\n\t":
                                    var v161 bool = true
                                    _ = v161
                                }
                            }
                        }
                    }
                    for fi162 := 0; fi162 < 16; fi162++ {
                        switch int(v1) {
                        case 32768:
                            for fi163 := 0; fi163 < 7; fi163++ {
                                if true {
                                    v101 = float64(float32((v101 - (v3 - 0.1))))
                                } else {
                                    var v164 uint8 = 1
                                    v164 = v164
                                    var v165 string = "hello"
                                    _ = v164
                                    _ = v165
                                }
                                switch "foo bar" {
                                case "hello":
                                    v2 = v2
                                case "a":
                                    fmt.Println(fi163, v101, v2)
                                case "foo bar":
                                    v166 := []int{}
                                    _ = v166
                                }
                                _ = fi163
                            }
                        case 32767:
                            if (false && (-2.25 == ((v3 - 0.1) + 1e10))) {
                                for fi167 := 0; fi167 < 17; fi167++ {
                                    var v168 int = (fi167 / ((int(v1)) | 1))
                                    v3 = (float64(((float32((v101 / v101)) - ((float32(v101) - 0.1) * 1e10)) * 0.0)) * (v101 * v101))
                                    _ = fi167
                                    _ = v168
                                }
                                var v169 uint8 = 42
                                _ = v169
                            } else {
                                if true {
                                    var v170 bool = (!(!false))
                                    v171 := GenS25{F0: v121}
                                    _ = v170
                                    _ = v171
                                } else {
                                    v172 := GenS14{F0: -1}
                                    _ = v172
                                }
                                v100 -= (v100 & (^(v100)))
                                v173 := GenS26{F0: 2147483647, F1: v1, F2: 449}
                                _ = v173
                            }
                        case 9223372036854775807:
                            v174 := []int{128, fi162, fi162, 96}
                            _ = v174
                        default:
                            if (true && (65535 <= 1)) {
                                v1 -= v1
                                switch "" {
                                case "\n\t":
                                    v3 -= v3
                                }
                                for fi175 := 0; fi175 < 1; fi175++ {
                                    v176 := GenS16{F0: 2147483647}
                                    _ = len(v4)
                                    v177 := []int{fi162, (fi175 >> ((v100) & 63)), (-(fi162))}
                                    _ = fi175
                                    _ = v176
                                    _ = v177
                                }
                            }
                        }
                        var v178 float64 = 1.5
                        _ = fi162
                        _ = v178
                    }
                }
                _ = v137
            }
        }
        v2 -= v2
    } else {
        var v179 int8 = 84
        v4[601] = 0
        _ = v179
    }
    switch 9223372036854775807 {
    case 9223372036854775807:
        var v180 uint64 = (^(v1))
        _ = v180
    case -9223372036854775808:
        if (!(true && (float32(-2.25) > 3.14159))) {
            var v181 bool = ((^(v2)) < -32768)
            _ = v181
        }
    default:
        v182 := GenS21{F0: ("a" + (("\n\t" + ("" + ("\n\t" + "hello"))) + ("hello" + ("hello" + "foo bar")))), F1: 9223372036854775807}
        _ = v182
    }
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v101:", v101)
    fmt.Println("v121:", v121)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v100
    _ = v101
    _ = v121
    _ = v2
    _ = v3
    _ = v4
}
