package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 int16; F2 uint8 }
type GenS4 struct { F0 uint32; F1 int64; F2 uint8 }
type GenS5 struct { F0 uint; F1 uint64 }
type GenS6 struct { F0 uint8; F1 uint32 }
type GenS7 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS8 struct { F0 int; F1 int }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint16; F1 string }
type GenS11 struct { F0 uint32; F1 string; F2 int8 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 bool; F1 uint16; F2 string }
type GenS14 struct { F0 int; F1 bool; F2 uint }
type GenS15 struct { F0 int; F1 int64; F2 uint32 }
type GenS16 struct { F0 string; F1 string }
type GenS17 struct { F0 int }
type GenS18 struct { F0 uint32; F1 int32 }
func genFn0() uint64 {
    var v4 int16 = 1
    v4 -= (v4 | 288)
    var v5 uint64 = 42
    _ = v4
    _ = v5
    return 895
}
func genFn1() uint16 {
    if (0 != 206) {
        var v6 int16 = 892
        if ((v6 >= (-(v6))) || (59 == 0)) {
            for fi7 := 0; fi7 < 17; fi7++ {
                var v8 uint16 = 1
                var v9 uint8 = 0
                _ = fi7
                _ = v8
                _ = v9
            }
            var v10 int8 = 17
            v6 += 42
            _ = v10
        }
        _ = v6
    } else {
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 uint16 = 2
            fmt.Println(v12, fi11)
            _ = fi11
            _ = v12
        }
        v13 := GenS1{F0: 32767, F1: 0}
        _ = v13
    }
    var v14 int64 = 1
    var v15 uint = 4294967295
    _ = v14
    _ = v15
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 242
    var v3 float64 = -2.25
    fmt.Println(v3, v3)
    for fi16 := 0; fi16 < 17; fi16++ {
        v17 := GenS1{F0: -1, F1: 2147483647}
        fmt.Println(fi16)
        switch ((("" + "a") + ("\n\t" + "")) + "") {
        case "foo bar":
            fmt.Println(fi16, v2)
        case "\n\t":
            v2 = (v2 + 37)
        }
        _ = fi16
        _ = v17
    }
    switch 4294967295 {
    case 127:
        fmt.Println(v3)
    case -2147483648:
        v3 -= (v3 - v3)
    case -1:
        var v18 bool = false
        _ = v18
    default:
        v19 := []int{127, 0}
        _ = v19
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        for fi21 := 0; fi21 < 13; fi21++ {
            var v22 uint16 = 65535
            var v23 int16 = 1
            var v24 int8 = -1
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
        v25 := GenS1{F0: 1, F1: -449}
        _ = fi20
        _ = v25
    }
    var v26 int64 = 1
    if false {
        for fi27 := 0; fi27 < 4; fi27++ {
            switch fi27 {
            case 9223372036854775807:
                for fi28 := 0; fi28 < 19; fi28++ {
                    v2 = 1
                    v29 := GenS2{F0: -836}
                    var v30 bool = true
                    var v31 uint16 = 630
                    _ = fi28
                    _ = v29
                    _ = v30
                    _ = v31
                }
            }
            switch "" {
            case "hello":
                if false {
                    var v32 uint32 = 422
                    fmt.Println(v26, fi27)
                    _ = v32
                } else {
                    if false {
                        v33 := GenS3{F0: "\n\t", F1: -32768, F2: v2}
                        _ = v33
                    } else {
                        switch "\n\t" {
                        case "hello":
                            v34 := []int{(fi27 + fi27)}
                            _ = v34
                        case "a":
                            for fi35 := 0; fi35 < 16; fi35++ {
                                switch fi35 {
                                case 127:
                                    var v36 uint32 = 65535
                                    _ = v36
                                case 4294967295:
                                    switch 2147483647 {
                                    case 779:
                                        var v37 int32 = 42
                                        _ = v37
                                    case 0:
                                        v26 -= 42
                                        v26 += (v26 ^ v26)
                                    case -433:
                                        var v38 float32 = float32(((v3 * v3) / ((v3 * float64(0.1)) + v3)))
                                        _ = v38
                                    }
                                case -632:
                                    v3 += (1e10 - v3)
                                default:
                                    var v39 int16 = 1000
                                    if true {
                                        v40 := []int{fi27}
                                        fmt.Println(v26, v39, v26)
                                        v40 = append(v40, -1)
                                        _ = v40
                                    } else {
                                        var v41 uint64 = 2
                                        if true {
                                            var v42 uint64 = (v41 << ((648) & 63))
                                            var v43 float64 = ((3.14159 * 3.14159) - (v3 / v3))
                                            _ = v42
                                            _ = v43
                                        }
                                        v2 -= v2
                                        v44 := GenS4{F0: 42, F1: v26, F2: 2}
                                        _ = v41
                                        _ = v44
                                    }
                                    v39 -= 1000
                                    _ = v39
                                }
                                v45 := []int{(fi27 << ((776) & 63))}
                                _ = fi35
                                _ = v45
                            }
                        default:
                            v46 := []int{fi27, (-(fi27)), fi27, fi27}
                            _ = v46
                        }
                        v3 += float64(0.1)
                    }
                    if true {
                        v2 += (v2 | 2)
                    } else {
                        if (false && (v3 <= ((float64(-2.25) + float64(3.14159)) - 0.0))) {
                            for fi47 := 0; fi47 < 1; fi47++ {
                                if (!false) {
                                    delete(v1, "")
                                    v48 := GenS1{F0: 32767, F1: 977}
                                    _ = v48
                                }
                                if false {
                                    v3 -= v3
                                    v49 := GenS5{F0: 255, F1: 70}
                                    _ = v49
                                } else {
                                    var v50 uint64 = genFn0()
                                    _ = v50
                                }
                                v51 := []int{fi27}
                                var v52 int16 = 0
                                _ = fi47
                                _ = v51
                                _ = v52
                            }
                            var v53 int = (-(fi27))
                            _ = v53
                        }
                        v54 := []int{0}
                        _ = v54
                    }
                }
            case "a":
                v2 = 128
            case "\n\t":
                for fi55 := 0; fi55 < 16; fi55++ {
                    var v56 int8 = 1
                    v57 := GenS5{F0: 272, F1: 654}
                    var v58 float32 = (1e10 + 0.1)
                    _ = fi55
                    _ = v56
                    _ = v57
                    _ = v58
                }
            }
            if (((v3 + 0.1) * float64(float32(v3))) > 0.0) {
                v59 := GenS6{F0: v2, F1: 58}
                v60 := []int{}
                switch fi27 {
                case 9223372036854775807:
                    for fi61 := 0; fi61 < 6; fi61++ {
                        if (!false) {
                            switch 315 {
                            case 0:
                                var v62 int8 = 0
                                _ = v62
                            case -9223372036854775808:
                                if false {
                                    if len(v60) > 0 {
                                        v60[0] = (fi61 * fi61)
                                    }
                                } else {
                                    switch fi27 {
                                    case 2147483647:
                                        v2 -= v2
                                    default:
                                        if (false || false) {
                                            if true {
                                                if len(v60) >= 2 {
                                                    v60 = v60[1:len(v60)]
                                                }
                                            }
                                            if true {
                                                var v63 uint16 = 32767
                                                if true {
                                                    if ((v3 >= 1e10) && true) {
                                                        if false {
                                                            v64 := GenS4{F0: 287, F1: (v26 - (v26 - (v26 % (((v26 >> ((298) & 63))) | 1)))), F2: 42}
                                                            _ = len(v60)
                                                            v65 := GenS3{F0: "hello", F1: -739, F2: 127}
                                                            _ = v64
                                                            _ = v65
                                                        } else {
                                                            fmt.Println(fi27)
                                                        }
                                                        switch (fi61 % ((fi27) | 1)) {
                                                        case 415:
                                                            var v66 int = -871
                                                            _ = v66
                                                        case 0:
                                                            if len(v60) >= 2 {
                                                                v60 = v60[1:len(v60)]
                                                            }
                                                        }
                                                    } else {
                                                        v63 = genFn1()
                                                        v67 := v1[""]
                                                        v67 = uint32(v2)
                                                        _ = v67
                                                    }
                                                    var v68 uint16 = (v63 * v63)
                                                    _ = v68
                                                }
                                                var v69 uint16 = (^(v63))
                                                if (!false) {
                                                    v70 := GenS4{F0: 716, F1: (v26 & v26), F2: 128}
                                                    v63 -= v69
                                                    _ = v70
                                                } else {
                                                    switch ("foo bar" + "hello") {
                                                    case "a":
                                                        switch fi27 {
                                                        case 262:
                                                            if true {
                                                                v2 = 193
                                                            } else {
                                                                v71 := GenS5{F0: 630, F1: 42}
                                                                var v72 int64 = (v26 * (v26 & -886))
                                                                var v73 uint64 = 524
                                                                _ = v71
                                                                _ = v72
                                                                _ = v73
                                                            }
                                                        case -113:
                                                            if (-1 <= 966) {
                                                                v74 := v1["hello"]
                                                                v69 = v69
                                                                v63 += uint16(v2)
                                                                _ = v74
                                                            }
                                                            fmt.Println(v69)
                                                            switch fi61 {
                                                            case 1:
                                                                var v75 int8 = -50
                                                                _ = v75
                                                            case 42:
                                                                if ((true || (!true)) && false) {
                                                                    var v76 float64 = v3
                                                                    v3 += 1e10
                                                                    v63 += (v69 >> ((131) & 63))
                                                                    _ = v76
                                                                } else {
                                                                    v26 += -9223372036854775808
                                                                    if true {
                                                                        if (!true) {
                                                                            v26 += -9223372036854775808
                                                                            switch int(v2) {
                                                                            case 869:
                                                                                if (1 > 42) {
                                                                                    var v77 uint8 = v2
                                                                                    var v78 string = ("\n\t" + "foo bar")
                                                                                    _ = v77
                                                                                    _ = v78
                                                                                }
                                                                            case 4294967295:
                                                                                if false {
                                                                                    var v79 int32 = 0
                                                                                    _ = v79
                                                                                } else {
                                                                                    v1[("hello" + ("" + "hello"))] = 2147483648
                                                                                    var v80 bool = false
                                                                                    _ = v80
                                                                                }
                                                                            case 1:
                                                                                v2 = (v2 + (v2 ^ v2))
                                                                            }
                                                                            fmt.Println(v69, fi27)
                                                                            var v81 int16 = 42
                                                                            var v82 uint8 = 2
                                                                            _ = v81
                                                                            _ = v82
                                                                        }
                                                                        switch "foo bar" {
                                                                        case "a":
                                                                            v69 += uint16(v2)
                                                                            v63 = v69
                                                                            v83 := GenS7{F0: 4294967295, F1: v2, F2: 2147483647}
                                                                            _ = v83
                                                                        case "":
                                                                            v84 := GenS4{F0: 0, F1: (v26 << ((872) & 63)), F2: 2}
                                                                            _ = v84
                                                                        case "\n\t":
                                                                            var v85 bool = false
                                                                            _ = v85
                                                                        default:
                                                                            if len(v60) > 0 {
                                                                                v60[0] = (fi61 + (^(fi27)))
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            case -730:
                                                                v63 += 1
                                                            }
                                                        case 9223372036854775807:
                                                            fmt.Println(v26, v63)
                                                        }
                                                    }
                                                    v2 = (^(v2))
                                                    var v86 string = ""
                                                    _ = v86
                                                }
                                                if false {
                                                    var v87 float64 = 3.14159
                                                    var v88 float64 = (-2.25 - v87)
                                                    _ = v87
                                                    _ = v88
                                                }
                                                v69 = genFn1()
                                                v2 = 255
                                                _ = v63
                                                _ = v69
                                            }
                                        }
                                    }
                                    fmt.Println(v3, v3)
                                }
                            case -2147483648:
                                v89 := GenS8{F0: fi27, F1: (fi27 | fi61)}
                                _ = v89
                            default:
                                fmt.Println(fi61, fi27, v2)
                            }
                            v3 -= v3
                            var v90 uint32 = 0
                            fmt.Println(v26, v90)
                            var v91 int32 = 273
                            fmt.Println(v90, v3, v91)
                            _ = v90
                            _ = v91
                        } else {
                            var v92 string = ((("foo bar" + "foo bar") + "") + ("hello" + ("a" + ("\n\t" + "a"))))
                            if len(v60) >= 2 {
                                v60 = v60[1:len(v60)]
                            }
                            _ = v92
                        }
                        v93 := v1[""]
                        switch fi61 {
                        case 127:
                            switch (-(fi61)) {
                            case 1:
                                v94 := GenS9{F0: 521}
                                _ = v94
                            case -1:
                                v95 := GenS9{F0: 42}
                                _ = v95
                            }
                        }
                        var v96 int64 = v26
                        if true {
                            switch ("\n\t" + "") {
                            case "":
                                switch (fi27 * fi61) {
                                case 32768:
                                    v96 -= (v96 - v96)
                                case -9223372036854775808:
                                    var v97 int8 = 42
                                    _ = v97
                                case 32767:
                                    if false {
                                        var v98 int64 = 1
                                        var v99 uint16 = 1
                                        _ = v98
                                        _ = v99
                                    } else {
                                        v100 := v1["foo bar"]
                                        switch -918 {
                                        case 1:
                                            var v101 uint32 = 648
                                            _ = v101
                                        case 662:
                                            v1[("\n\t" + "\n\t")] = (v100 ^ 42)
                                        case -2147483648:
                                            if len(v60) >= 2 {
                                                v60 = v60[1:len(v60)]
                                            }
                                        }
                                        if ((v96 | (^(v26))) == (v26 >> ((1) & 63))) {
                                            v26 = v26
                                        } else {
                                            var v102 int8 = -26
                                            v103 := GenS1{F0: -1, F1: -939}
                                            fmt.Println(fi27, v96)
                                            var v104 uint = 1
                                            v1["foo bar"] = v93
                                            _ = v102
                                            _ = v103
                                            _ = v104
                                        }
                                        _ = v100
                                    }
                                default:
                                    if false {
                                        fmt.Println(fi61, v26, v96)
                                    } else {
                                        var v105 uint32 = (^(v93))
                                        _ = v105
                                    }
                                    switch (fi27 ^ (fi27 | 127)) {
                                    case -861:
                                        if len(v60) >= 2 {
                                            v60 = v60[1:len(v60)]
                                        }
                                    }
                                    v106 := GenS6{F0: 214, F1: 42}
                                    _ = v106
                                }
                            case "a":
                                if (321 != -1) {
                                    if (!(!false)) {
                                        var v107 float64 = v3
                                        var v108 int8 = 1
                                        var v109 uint = 166
                                        _ = v107
                                        _ = v108
                                        _ = v109
                                    } else {
                                        var v110 float32 = 3.14159
                                        if (false || (!false)) {
                                            v111 := GenS10{F0: genFn1(), F1: "foo bar"}
                                            v112 := GenS11{F0: (v93 * v93), F1: "foo bar", F2: -68}
                                            _ = v111
                                            _ = v112
                                        } else {
                                            switch ((("a" + "foo bar") + ("a" + "foo bar")) + "") {
                                            case "hello":
                                                switch (fi61 * (fi27 + fi27)) {
                                                case 0:
                                                    if true {
                                                        var v113 uint16 = uint16(v93)
                                                        _ = v113
                                                    }
                                                case 42:
                                                    var v114 uint64 = 4294967295
                                                    _ = v114
                                                default:
                                                    var v115 int32 = 1
                                                    _ = v115
                                                }
                                            case "":
                                                v60 = append(v60, fi61)
                                            default:
                                                var v116 float32 = (float32(0.0) / (v110 / (v110 * v110)))
                                                _ = v116
                                            }
                                            fmt.Println(v96)
                                        }
                                        fmt.Println(v3)
                                        if false {
                                            var v117 uint64 = 653
                                            v110 += (v110 / 1.0)
                                            var v118 uint8 = (v2 % (((v2 * 1)) | 1))
                                            _ = v117
                                            _ = v118
                                        } else {
                                            var v119 uint = 216
                                            _ = v119
                                        }
                                        v26 = v26
                                        _ = v110
                                    }
                                } else {
                                    switch fi27 {
                                    case 127:
                                        if (1 > 0) {
                                            var v120 uint = 18446744073709551615
                                            _ = v120
                                        } else {
                                            if true {
                                                fmt.Println(v26, v2)
                                                v1["\n\t"] = v93
                                                fmt.Println(v26, v96)
                                            }
                                            if (!("foo bar" != "")) {
                                                var v121 int16 = -894
                                                var v122 uint16 = 312
                                                _ = v121
                                                _ = v122
                                            }
                                            if ((^(v26)) > 0) {
                                                if false {
                                                    var v123 int = fi27
                                                    v60 = append(v60, fi27)
                                                    var v124 int = 2147483647
                                                    _ = v123
                                                    _ = v124
                                                } else {
                                                    v125 := GenS12{F0: v26}
                                                    v126 := GenS2{F0: (fi27 | (fi27 ^ (fi61 << ((18446744073709551615) & 63))))}
                                                    v127 := GenS13{F0: true, F1: 1, F2: "foo bar"}
                                                    _ = v125
                                                    _ = v126
                                                    _ = v127
                                                }
                                            } else {
                                                v93 += uint32(v2)
                                                if true {
                                                    fmt.Println(fi61)
                                                    var v128 float32 = (0.0 - 3.14159)
                                                    _ = v128
                                                }
                                            }
                                        }
                                    case -9223372036854775808:
                                        var v129 int16 = -1
                                        _ = v129
                                    case 1:
                                        var v130 int32 = 1
                                        switch "a" {
                                        case "foo bar":
                                            _ = len(v60)
                                        }
                                        _ = v130
                                    default:
                                        if (false || (!false)) {
                                            v131 := GenS5{F0: 65535, F1: 18446744073709551615}
                                            v132 := GenS14{F0: (fi27 * 4294967295), F1: false, F2: 335}
                                            _ = v131
                                            _ = v132
                                        }
                                    }
                                    v2 += 42
                                    if (2 >= 4294967295) {
                                        if false {
                                            switch ("foo bar" + "hello") {
                                            case "a":
                                                var v133 float32 = 0.1
                                                _ = v133
                                            case "hello":
                                                switch (^(fi61)) {
                                                case 928:
                                                    fmt.Println(fi61, v26, v2)
                                                case 9223372036854775807:
                                                    var v134 bool = false
                                                    _ = v134
                                                case 2147483647:
                                                    v26 -= 155
                                                default:
                                                    var v135 int = 276
                                                    _ = v135
                                                }
                                            }
                                        }
                                    } else {
                                        fmt.Println(v3, v26, v2)
                                        fmt.Println(v2, fi27, v93)
                                        if false {
                                            v3 -= ((-2.25 / (float64(-2.25) * float64(-2.25))) + (1.5 + v3))
                                        }
                                        if (!(!false)) {
                                            v136 := GenS1{F0: -997, F1: 1}
                                            if false {
                                                var v137 string = "\n\t"
                                                v3 += v3
                                                v138 := GenS10{F0: 444, F1: v137}
                                                _ = v137
                                                _ = v138
                                            } else {
                                                var v139 string = (((("" + "\n\t") + "a") + "\n\t") + "hello")
                                                switch "\n\t" {
                                                case "hello":
                                                    v93 = (^(v93))
                                                case "":
                                                    if (-2.25 != -2.25) {
                                                        _ = len(v60)
                                                        v96 -= 42
                                                        var v140 uint8 = (v2 >> ((255) & 63))
                                                        _ = v140
                                                    }
                                                default:
                                                    switch "" {
                                                    case "":
                                                        switch 780 {
                                                        case -9223372036854775808:
                                                            v93 += (^(v93))
                                                        }
                                                    case "a":
                                                        var v141 float64 = (v3 / -2.25)
                                                        _ = v141
                                                    case "\n\t":
                                                        var v142 int8 = 0
                                                        _ = v142
                                                    default:
                                                        switch 0 {
                                                        case 32767:
                                                            v60 = append(v60, (fi27 << ((25) & 63)))
                                                        case 518:
                                                            v143 := GenS15{F0: 4294967295, F1: v96, F2: v93}
                                                            _ = v143
                                                        }
                                                    }
                                                }
                                                var v144 uint16 = 2
                                                fmt.Println(v26, v93, v96)
                                                _ = v139
                                                _ = v144
                                            }
                                            if false {
                                                if false {
                                                    v145 := GenS5{F0: 753, F1: 4294967295}
                                                    switch "hello" {
                                                    case "":
                                                        _ = len(v60)
                                                    case "a":
                                                        v93 += (v93 >> ((255) & 63))
                                                        switch ((("a" + "hello") + "\n\t") + "foo bar") {
                                                        case "foo bar":
                                                            switch 9223372036854775807 {
                                                            case -48:
                                                                fmt.Println(v96, v93, fi61)
                                                            case -714:
                                                                var v146 uint = 18446744073709551615
                                                                _ = v146
                                                            case -956:
                                                                v147 := GenS9{F0: 826}
                                                                _ = v147
                                                            }
                                                        default:
                                                            fmt.Println(fi27, v3)
                                                        }
                                                        var v148 float32 = float32(0.1)
                                                        _ = v148
                                                    }
                                                    _ = v145
                                                }
                                                switch "hello" {
                                                case "":
                                                    if len(v60) > 0 {
                                                        v60[0] = (fi27 ^ (^(fi61)))
                                                    }
                                                    var v149 float32 = (1e10 / 1.5)
                                                    _ = v149
                                                case "hello":
                                                    if len(v60) >= 2 {
                                                        v60 = v60[1:len(v60)]
                                                    }
                                                default:
                                                    var v150 bool = (true || (2147483647 >= 1))
                                                    _ = v150
                                                }
                                                fmt.Println(fi61, fi61, fi27)
                                            }
                                            _ = v136
                                        } else {
                                            v151 := GenS16{F0: "\n\t", F1: ("" + "")}
                                            fmt.Println(fi27, v93, v3)
                                            v152 := GenS5{F0: 18446744073709551615, F1: 271}
                                            _ = v151
                                            _ = v152
                                        }
                                        _ = len(v60)
                                    }
                                }
                            case "\n\t":
                                if (2 != 65535) {
                                    var v153 uint16 = 42
                                    v154 := GenS10{F0: v153, F1: "\n\t"}
                                    var v155 int64 = v96
                                    _ = v153
                                    _ = v154
                                    _ = v155
                                }
                                var v156 uint32 = v93
                                _ = v156
                            default:
                                if (3.14159 != (1e10 - v3)) {
                                    v3 -= float64(0.1)
                                }
                            }
                            if false {
                                switch fi61 {
                                case 0:
                                    v157 := GenS16{F0: "hello", F1: "\n\t"}
                                    _ = v157
                                case 393:
                                    var v158 bool = true
                                    _ = v158
                                default:
                                    v96 = (v26 ^ (v96 * (v26 | -9223372036854775808)))
                                }
                            }
                            v93 = (v93 % ((v93) | 1))
                        }
                        if len(v60) >= 2 {
                            v60 = v60[1:len(v60)]
                        }
                        v3 = v3
                        _ = fi61
                        _ = v93
                        _ = v96
                    }
                case -1:
                    var v159 uint = 0
                    _ = v159
                case 127:
                    v160 := GenS17{F0: fi27}
                    _ = v160
                default:
                    if (("\n\t" + "\n\t") > "a") {
                        if false {
                            var v161 int = fi27
                            v161 -= fi27
                            _ = v161
                        } else {
                            switch ("hello" + "") {
                            case "hello":
                                for fi162 := 0; fi162 < 15; fi162++ {
                                    v163 := GenS18{F0: 1, F1: 2147483647}
                                    v3 += (v3 * 1.5)
                                    _ = fi162
                                    _ = v163
                                }
                            default:
                                var v164 uint32 = 42
                                _ = v164
                            }
                            if (3.14159 <= (0.1 * 3.14159)) {
                                var v165 bool = (228 <= 227)
                                switch "" {
                                case "\n\t":
                                    if len(v60) >= 2 {
                                        v60 = v60[1:len(v60)]
                                    }
                                case "":
                                    v165 = (float64(0.1) <= 1e10)
                                case "a":
                                    for fi166 := 0; fi166 < 15; fi166++ {
                                        v165 = ((!v165) || (!(false && v165)))
                                        _ = fi166
                                    }
                                default:
                                    var v167 int8 = -12
                                    _ = v167
                                }
                                _ = v165
                            } else {
                                v168 := GenS17{F0: (fi27 ^ (fi27 ^ 0))}
                                for fi169 := 0; fi169 < 20; fi169++ {
                                    var v170 uint16 = 65535
                                    fmt.Println(fi27, v26)
                                    _ = fi169
                                    _ = v170
                                }
                                _ = v168
                            }
                        }
                    } else {
                        var v171 float64 = (float64(3.14159) + ((v3 * ((float64((-2.25 * (-2.25 - ((3.14159 / -2.25) * -2.25)))) + v3) - (v3 + ((v3 - v3) * (float64(-2.25) + (v3 * (((v3 + (v3 - v3)) / 1.0) - (v3 + ((((v3 * v3) - 1e10) * v3) + float64(0.0)))))))))) + v3))
                        v172 := GenS14{F0: fi27, F1: true, F2: 255}
                        _ = v171
                        _ = v172
                    }
                }
                _ = v59
                _ = v60
            }
            _ = fi27
        }
    } else {
        v26 = v26
        var v173 uint8 = v2
        _ = v173
    }
    for fi174 := 0; fi174 < 14; fi174++ {
        var v175 int16 = int16(v26)
        _ = fi174
        _ = v175
    }
    var v176 float32 = (0.0 * (1e10 / float32(v3)))
    switch ("hello" + "\n\t") {
    case "":
        if ((((v2 > 0) && true) || (false || ((^(v2)) == 127))) && (!(65535 != 2))) {
            var v177 uint = 535
            _ = v177
        }
    case "\n\t":
        v176 += (v176 / v176)
    }
    var v178 uint64 = 42
    var v179 uint32 = 909
    var v180 float64 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v176:", v176)
    fmt.Println("v178:", v178)
    fmt.Println("v179:", v179)
    fmt.Println("v180:", v180)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v176
    _ = v178
    _ = v179
    _ = v180
    _ = v2
    _ = v26
    _ = v3
}
