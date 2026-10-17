package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 int32; F1 int8; F2 int64 }
type GenS5 struct { F0 int; F1 int32; F2 string }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint32; F1 uint32 }
type GenS9 struct { F0 uint16; F1 uint64 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint32; F1 string }
type GenS12 struct { F0 int16; F1 uint64 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 uint32; F1 int16 }
type GenS15 struct { F0 uint8; F1 uint16; F2 int8 }
func genFn0(p0 string, p1 uint32, p2 bool) int16 {
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := GenS1{F0: -1, F1: -836, F2: -525}
        for fi7 := 0; fi7 < 6; fi7++ {
            p0 = (p0 + p0)
            var v8 int16 = 157
            v9 := GenS2{F0: 805}
            var v10 int32 = -176
            p0 = "a"
            if (!(!p2)) {
                v8 += 42
            }
            switch fi5 {
            case 32768:
                v11 := []int{}
                _ = v11
            case 32767:
                p2 = p2
            }
            v12 := []int{128}
            if p2 {
                v13 := GenS3{F0: 2147483647, F1: 1}
                switch "hello" {
                case "foo bar":
                    fmt.Println(v10, fi7, v8)
                default:
                    var v14 int64 = 0
                    _ = v14
                }
                fmt.Println(fi5)
                _ = v13
            }
            _ = fi7
            _ = v10
            _ = v12
            _ = v8
            _ = v9
        }
        for fi15 := 0; fi15 < 2; fi15++ {
            switch p0 {
            case "a":
                var v16 int8 = -128
                _ = v16
            case "":
                fmt.Println(fi5, fi5)
            case "foo bar":
                if p2 {
                    var v17 uint16 = 42
                    fmt.Println(fi5, p1, fi15)
                    v18 := []int{fi5}
                    _ = v17
                    _ = v18
                } else {
                    v19 := GenS1{F0: 42, F1: 104, F2: 764}
                    var v20 float64 = 1e10
                    switch "foo bar" {
                    case "hello":
                        fmt.Println(p1, fi15)
                    }
                    _ = v19
                    _ = v20
                }
            }
            v21 := GenS2{F0: -2147483648}
            _ = fi15
            _ = v21
        }
        _ = fi5
        _ = v6
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        var v23 int32 = 207
        _ = fi22
        _ = v23
    }
    if false {
        if ((0.1 * float32(0.1)) < ((-2.25 / float32(1e10)) + 0.0)) {
            v24 := []int{-993, 0, -716}
            _ = v24
        } else {
            v25 := []int{626}
            _ = v25
        }
        if p2 {
            fmt.Println(p1, p0, p1)
            if p2 {
                p1 -= 2147483647
                if p2 {
                    for fi26 := 0; fi26 < 18; fi26++ {
                        v27 := []int{(fi26 ^ (fi26 + fi26)), (^(fi26))}
                        v28 := GenS4{F0: 226, F1: 37, F2: -9223372036854775808}
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                }
                v29 := GenS3{F0: -621, F1: 100}
                _ = v29
            } else {
                if false {
                    var v30 uint16 = 2
                    var v31 uint8 = 0
                    fmt.Println(p0, p1, p2)
                    _ = v30
                    _ = v31
                } else {
                    for fi32 := 0; fi32 < 9; fi32++ {
                        if (p0 > "") {
                            var v33 float64 = 3.14159
                            for fi34 := 0; fi34 < 15; fi34++ {
                                var v35 int32 = -2147483648
                                var v36 int = (fi34 % ((262) | 1))
                                switch (("a" + p0) + p0) {
                                case "\n\t":
                                    v36 = fi34
                                default:
                                    p1 -= p1
                                }
                                _ = fi34
                                _ = v35
                                _ = v36
                            }
                            _ = v33
                        } else {
                            if (145 <= 65535) {
                                var v37 uint = 4294967295
                                _ = v37
                            }
                            p2 = ((42 != -1) && p2)
                        }
                        _ = fi32
                    }
                    switch 4294967295 {
                    case 2147483647:
                        p2 = (-2147483648 > 344)
                    case 0:
                        v38 := GenS2{F0: -2147483648}
                        _ = v38
                    default:
                        for fi39 := 0; fi39 < 6; fi39++ {
                            for fi40 := 0; fi40 < 9; fi40++ {
                                p2 = true
                                p0 = ((((p0 + (p0 + "foo bar")) + (p0 + "")) + p0) + p0)
                                _ = fi40
                            }
                            if true {
                                p1 += (^(p1))
                            } else {
                                v41 := []int{fi39}
                                _ = v41
                            }
                            _ = fi39
                        }
                    }
                }
                v42 := []int{}
                _ = v42
            }
        } else {
            p1 = p1
            p0 = (p0 + ("a" + "foo bar"))
            var v43 bool = p2
            _ = v43
        }
    } else {
        fmt.Println(p0)
    }
    switch p0 {
    case "\n\t":
        for fi44 := 0; fi44 < 3; fi44++ {
            p2 = (!false)
            p1 = p1
            p2 = (!p2)
            _ = fi44
        }
    default:
        if true {
            v45 := GenS4{F0: -1, F1: 0, F2: -290}
            _ = v45
        } else {
            p1 += p1
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1() string {
    if (true || (true || (2 > 255))) {
        var v46 uint32 = 1
        if ("" >= ((("a" + "\n\t") + "\n\t") + "a")) {
            v46 = v46
        }
        v46 = (v46 << ((4294967295) & 63))
        _ = v46
    }
    for fi47 := 0; fi47 < 14; fi47++ {
        v48 := []int{fi47, -413, 2147483647, (^(fi47))}
        v49 := GenS4{F0: 2147483647, F1: -1, F2: 1}
        var v50 uint64 = 0
        _ = fi47
        _ = v48
        _ = v49
        _ = v50
    }
    return ("" + "a")
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{-681, -1, 9223372036854775807, 4294967295}
    var v3 string = ("foo bar" + "a")
    var v4 uint16 = 32767
    for fi51 := 0; fi51 < 14; fi51++ {
        var v52 bool = (!false)
        delete(v1, 2)
        if v52 {
            v1[65535] = 1
            switch fi51 {
            case 0:
                v4 -= 255
            case 127:
                fmt.Println(v52)
            case 42:
                v52 = true
            }
            fmt.Println(fi51, fi51)
        } else {
            v4 = (v4 & (v4 | 255))
            delete(v1, 403)
        }
        _ = fi51
        _ = v52
    }
    v4 += (v4 >> ((42) & 63))
    var v53 uint32 = 0
    v54 := GenS4{F0: 1, F1: -128, F2: 9223372036854775807}
    v4 += 658
    var v55 uint8 = 255
    var v56 string = (v3 + (("hello" + (("a" + v3) + "foo bar")) + v3))
    for fi57 := 0; fi57 < 13; fi57++ {
        v2 = append(v2, fi57)
        if true {
            var v58 int = (fi57 & (fi57 << ((226) & 63)))
            if (true || false) {
                var v59 int = 218
                v56 = v56
                switch v58 {
                case 594:
                    fmt.Println(v53)
                case -1:
                    v58 += (fi57 - (fi57 ^ v58))
                case 128:
                    if true {
                        for fi60 := 0; fi60 < 4; fi60++ {
                            if true {
                                if (!(2 != 4294967295)) {
                                    fmt.Println(v4, v53, v56)
                                    if false {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                    v3 = v3
                                    var v61 int64 = -87
                                    _ = v61
                                }
                            }
                            var v62 int64 = 9223372036854775807
                            v4 += (^(v4))
                            switch fi60 {
                            case 0:
                                v63 := GenS5{F0: (fi60 | (v58 | 2147483647)), F1: -2147483648, F2: v56}
                                _ = v63
                            case -417:
                                v55 = (^(v55))
                                _ = len(v2)
                            case 9223372036854775807:
                                if true {
                                    v2 = append(v2, v58)
                                    fmt.Println(fi57, v55, v55)
                                } else {
                                    if false {
                                        v55 -= v55
                                        if true {
                                            if (!(-2.25 == 1e10)) {
                                                var v64 uint = 661
                                                _ = v64
                                            } else {
                                                var v65 uint64 = uint64(v55)
                                                fmt.Println(v4, v55, v58)
                                                if (true || false) {
                                                    var v66 uint = 2
                                                    if (uint32(v55) >= (v53 << (((v66 + v66)) & 63))) {
                                                        switch v3 {
                                                        case "a":
                                                            var v67 float32 = 1e10
                                                            _ = v67
                                                        case "\n\t":
                                                            var v68 int64 = v62
                                                            _ = v68
                                                        default:
                                                            if (!true) {
                                                                v2 = append(v2, 9223372036854775807)
                                                                var v69 int = 1
                                                                var v70 string = v3
                                                                _ = v69
                                                                _ = v70
                                                            }
                                                        }
                                                        fmt.Println(v62)
                                                        var v71 uint32 = 0
                                                        _ = v71
                                                    } else {
                                                        v65 = (v65 | uint64(v53))
                                                    }
                                                    v56 = ((v3 + v3) + "\n\t")
                                                    _ = v66
                                                } else {
                                                    switch v56 {
                                                    case "foo bar":
                                                        v3 = v3
                                                    case "\n\t":
                                                        switch v58 {
                                                        case -491:
                                                            var v72 int32 = 0
                                                            _ = v72
                                                        case -1:
                                                            v65 += v65
                                                        }
                                                    case "a":
                                                        var v73 uint64 = (^(v65))
                                                        _ = v73
                                                    default:
                                                        _ = len(v1)
                                                    }
                                                    switch v3 {
                                                    case "\n\t":
                                                        delete(v1, v53)
                                                    case "":
                                                        v74 := GenS6{F0: 42}
                                                        _ = v74
                                                    }
                                                    var v75 uint16 = (^(v4))
                                                    var v76 int8 = int8(v53)
                                                    v77 := v1[v53]
                                                    _ = v75
                                                    _ = v76
                                                    _ = v77
                                                }
                                                v78 := GenS7{F0: (-(v62))}
                                                _ = v65
                                                _ = v78
                                            }
                                            var v79 int8 = 0
                                            if true {
                                                fmt.Println(v4, v79, v58)
                                                var v80 int8 = v79
                                                v2 = append(v2, int(v55))
                                                v81 := GenS1{F0: 42, F1: 978, F2: 39}
                                                v4 = v4
                                                _ = v80
                                                _ = v81
                                            }
                                            if false {
                                                v55 -= 243
                                                _ = len(v2)
                                            } else {
                                                fmt.Println(v58, v4)
                                            }
                                            _ = v79
                                        } else {
                                            var v82 string = v3
                                            switch v59 {
                                            case 657:
                                                var v83 uint64 = 227
                                                _ = v83
                                            case 1:
                                                v1[v53] = v53
                                            default:
                                                var v84 string = ("hello" + v56)
                                                _ = v84
                                            }
                                            if len(v2) > 0 {
                                                v2[0] = (v58 - fi57)
                                            }
                                            _ = len(v2)
                                            _ = v82
                                        }
                                        if (!true) {
                                            if ((v55 % ((v55) | 1)) < (v55 | (v55 & uint8(v4)))) {
                                                v62 += 0
                                            } else {
                                                switch ((((v3 + "a") + "\n\t") + v56) + ("" + v3)) {
                                                case "\n\t":
                                                    var v85 float64 = 0.0
                                                    v62 = (v62 ^ (-(v62)))
                                                    v86 := GenS4{F0: -2147483648, F1: 127, F2: -9223372036854775808}
                                                    _ = v85
                                                    _ = v86
                                                case "":
                                                    v58 = -9223372036854775808
                                                default:
                                                    switch fi60 {
                                                    case 2147483647:
                                                        if (9223372036854775807 != 0) {
                                                            var v87 int = v59
                                                            _ = v87
                                                        }
                                                    default:
                                                        v55 = (^(v55))
                                                    }
                                                }
                                                v88 := GenS6{F0: 42}
                                                _ = v88
                                            }
                                            switch "a" {
                                            case "foo bar":
                                                v89 := v1[v53]
                                                _ = v89
                                            case "\n\t":
                                                if (0 < -1) {
                                                    if (0.1 < float32(3.14159)) {
                                                        v2 = append(v2, 32768)
                                                        v90 := GenS7{F0: (v62 * (v62 - v62))}
                                                        var v91 uint8 = 28
                                                        delete(v1, (v53 << ((89) & 63)))
                                                        _ = v90
                                                        _ = v91
                                                    }
                                                    var v92 int64 = v62
                                                    var v93 uint8 = uint8(v53)
                                                    _ = v92
                                                    _ = v93
                                                }
                                            default:
                                                if true {
                                                    var v94 string = "\n\t"
                                                    v1[v53] = 704
                                                    _ = v94
                                                } else {
                                                    v95 := GenS8{F0: v53, F1: v53}
                                                    v62 += (v62 - (v62 * 1))
                                                    v59 += (^(v59))
                                                    _ = v95
                                                }
                                            }
                                            v53 -= 2147483647
                                        }
                                        if false {
                                            switch (v56 + (v56 + "hello")) {
                                            case "":
                                                var v96 uint8 = v55
                                                _ = v96
                                            case "\n\t":
                                                switch -599 {
                                                case 42:
                                                    var v97 int32 = -118
                                                    _ = v97
                                                }
                                            case "a":
                                                if true {
                                                    var v98 uint8 = 128
                                                    _ = len(v2)
                                                    switch ("hello" + ("\n\t" + (v56 + v56))) {
                                                    case "":
                                                        var v99 uint32 = (^(v53))
                                                        _ = v99
                                                    case "foo bar":
                                                        if true {
                                                            var v100 uint16 = (v4 * 134)
                                                            _ = v100
                                                        } else {
                                                            v101 := GenS5{F0: 4294967295, F1: 42, F2: ((((v3 + v3) + "hello") + v56) + ("foo bar" + ("" + (v56 + ""))))}
                                                            _ = v101
                                                        }
                                                    case "hello":
                                                        var v102 int32 = 42
                                                        _ = v102
                                                    default:
                                                        switch v56 {
                                                        case "foo bar":
                                                            if false {
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    var v103 float32 = ((0.1 - 1e10) / ((0.1 - 1.5) + 0.1))
                                                                    _ = v103
                                                                case "":
                                                                    if (898 == 42) {
                                                                        var v104 string = ("\n\t" + v3)
                                                                        v105 := GenS9{F0: v4, F1: 4294967295}
                                                                        v3 = v56
                                                                        var v106 float64 = 0.1
                                                                        var v107 float64 = 0.0
                                                                        v107 += float64(((0.1 - float32((v106 / 3.14159))) / (3.14159 * (float32(v106) - -2.25))))
                                                                        _ = v104
                                                                        _ = v105
                                                                        _ = v106
                                                                        _ = v107
                                                                    } else {
                                                                        switch int(v55) {
                                                                        case -493:
                                                                            v55 += 0
                                                                        default:
                                                                            var v108 float32 = 1.5
                                                                            _ = v108
                                                                        }
                                                                        delete(v1, 65535)
                                                                    }
                                                                }
                                                                var v109 int16 = 42
                                                                if (((0.1 + (1e10 / (1e10 / 1.5))) > float32(1.5)) && true) {
                                                                    v2 = append(v2, v58)
                                                                    var v110 int = (v58 | 0)
                                                                    v56 = v56
                                                                    _ = v110
                                                                } else {
                                                                    fmt.Println(fi60)
                                                                }
                                                                v111 := v1[v53]
                                                                var v112 int16 = 1
                                                                _ = v109
                                                                _ = v111
                                                                _ = v112
                                                            }
                                                        case "":
                                                            v4 -= v4
                                                        case "a":
                                                            var v113 uint16 = (v4 ^ v4)
                                                            _ = v113
                                                        }
                                                        fmt.Println(v98)
                                                    }
                                                    v62 = (v62 - -213)
                                                    _ = v98
                                                } else {
                                                    fmt.Println(v56, v56, fi60)
                                                    v4 = 2
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            default:
                                                if (v53 > (v53 << ((4294967295) & 63))) {
                                                    if (!((-2.25 * (0.1 / (1.5 / (1.5 / -2.25)))) == (0.1 / ((1.5 / float32(3.14159)) / 1e10)))) {
                                                        var v114 uint = 980
                                                        v115 := GenS10{F0: v56}
                                                        v59 = 4294967295
                                                        _ = v114
                                                        _ = v115
                                                    }
                                                    var v116 int64 = 9223372036854775807
                                                    v59 += 42
                                                    _ = v116
                                                } else {
                                                    v55 = (v55 / ((v55) | 1))
                                                    var v117 int8 = 127
                                                    _ = v117
                                                }
                                            }
                                            v58 = fi60
                                            var v118 int16 = -274
                                            _ = v118
                                        }
                                        v53 = 2147483647
                                    }
                                    var v119 int64 = 42
                                    _ = v119
                                }
                            }
                            var v120 string = "\n\t"
                            _ = fi60
                            _ = v120
                            _ = v62
                        }
                    }
                }
                _ = v59
            }
            v121 := GenS11{F0: (v53 & 2147483648), F1: v3}
            _ = v121
            _ = v58
        } else {
            for fi122 := 0; fi122 < 6; fi122++ {
                v1[(v53 ^ (v53 * v53))] = v53
                switch fi122 {
                case 128:
                    fmt.Println(fi122, fi57, v55)
                case 2147483647:
                    v123 := GenS7{F0: 847}
                    _ = v123
                case 32767:
                    var v124 int32 = 764
                    _ = v124
                }
                fmt.Println(v56)
                v125 := GenS11{F0: (v53 & 7), F1: v3}
                _ = fi122
                _ = v125
            }
            v1[(v53 | v53)] = v53
            v56 = v56
        }
        v126 := GenS6{F0: 2147483647}
        _ = fi57
        _ = v126
    }
    switch v3 {
    case "foo bar":
        var v127 uint8 = (v55 << ((18446744073709551615) & 63))
        _ = v127
    case "a":
        var v128 uint64 = 0
        _ = v128
    }
    if false {
        v129 := GenS12{F0: -32768, F1: uint64(v53)}
        _ = v129
    }
    var v130 uint32 = (^(v53))
    v131 := GenS13{F0: -97}
    for fi132 := 0; fi132 < 17; fi132++ {
        switch fi132 {
        case -2147483648:
            if (-9223372036854775808 > 243) {
                var v133 uint64 = 5
                _ = v133
            } else {
                v3 = v3
                v134 := GenS14{F0: (v53 & v53), F1: -32768}
                _ = v134
            }
        default:
            for fi135 := 0; fi135 < 15; fi135++ {
                var v136 int8 = -128
                var v137 uint8 = 42
                v3 = "foo bar"
                _ = fi135
                _ = v136
                _ = v137
            }
        }
        _ = fi132
    }
    for fi138 := 0; fi138 < 5; fi138++ {
        v139 := GenS15{F0: 183, F1: v4, F2: -97}
        if false {
            for fi140 := 0; fi140 < 17; fi140++ {
                if (!true) {
                    var v141 int32 = -1
                    v3 = v56
                    v142 := GenS5{F0: (fi138 % ((fi140) | 1)), F1: (v141 * (v141 ^ v141)), F2: v56}
                    v2 = append(v2, fi140)
                    _ = v141
                    _ = v142
                } else {
                    if len(v2) > 0 {
                        v2[0] = 32767
                    }
                }
                var v143 int = fi140
                v144 := v1[4294967295]
                if (!(-9223372036854775808 != -9223372036854775808)) {
                    if (0 <= 1) {
                        if false {
                            v145 := GenS1{F0: genFn0((("" + ("" + "a")) + "foo bar"), v130, false), F1: -1, F2: 42}
                            if true {
                                fmt.Println(v130, fi138)
                                if false {
                                    fmt.Println(v130)
                                }
                            }
                            var v146 int = fi140
                            _ = v145
                            _ = v146
                        } else {
                            fmt.Println(v4, v3)
                            switch ("a" + v3) {
                            case "hello":
                                v147 := GenS3{F0: 1, F1: -1}
                                _ = v147
                            }
                        }
                    } else {
                        v143 += fi138
                        v55 -= v55
                        var v148 int16 = -503
                        var v149 int64 = -469
                        var v150 uint = 0
                        _ = v148
                        _ = v149
                        _ = v150
                    }
                    if (-128 > int8(v130)) {
                        fmt.Println(v53, v144, v3)
                        var v151 uint8 = v55
                        _ = v151
                    }
                    var v152 uint32 = v144
                    _ = v152
                }
                _ = fi140
                _ = v143
                _ = v144
            }
            v56 = "hello"
            var v153 float64 = 0.1
            _ = v153
        }
        switch 127 {
        case 32767:
            v154 := GenS4{F0: -518, F1: -1, F2: 1}
            _ = v154
        case -729:
            var v155 int64 = 9223372036854775807
            _ = v155
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = fi138
        _ = v139
    }
    v130 = (v130 ^ v130)
    switch 1 {
    case 0:
        v55 = (v55 ^ 255)
    case 250:
        if true {
            v156 := GenS11{F0: v130, F1: (v3 + v3)}
            var v157 int64 = -9223372036854775808
            for fi158 := 0; fi158 < 18; fi158++ {
                v159 := GenS5{F0: fi158, F1: -2147483648, F2: ("a" + "foo bar")}
                _ = len(v2)
                _ = fi158
                _ = v159
            }
            _ = v156
            _ = v157
        } else {
            v160 := GenS10{F0: "hello"}
            _ = v160
        }
    case 1:
        fmt.Println(v4)
    }
    v130 += v53
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v130:", v130)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v53:", v53)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v130
    _ = v131
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v54
    _ = v55
    _ = v56
}
