package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 uint32; F1 int16; F2 int64 }
type GenS5 struct { F0 int; F1 string }
type GenS6 struct { F0 uint16; F1 int32 }
type GenS7 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int; F1 string }
type GenS11 struct { F0 int; F1 int64; F2 uint64 }
type GenS12 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS13 struct { F0 uint16; F1 uint64 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 uint16; F1 bool }
type GenS16 struct { F0 int64; F1 int; F2 int64 }
func genFn0(p0 int, p1 uint64, p2 uint16) uint64 {
    if true {
        switch (("foo bar" + "a") + "hello") {
        case "foo bar":
            p0 += p0
        }
        var v5 uint = 1
        _ = v5
    }
    v6 := []int{p0}
    if (true || true) {
        var v7 int = -236
        v7 -= 32767
        _ = v7
    } else {
        var v8 int8 = 100
        for fi9 := 0; fi9 < 5; fi9++ {
            fmt.Println(p0, p2)
            var v10 int = (^(fi9))
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    var v11 bool = (false || (true && (((1.5 + 0.1) != (1e10 / 1.5)) || (911 < 791))))
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v6
    return p1
}
func genFn1(p0 int64) int64 {
    switch 1 {
    case 32767:
        if (true && (42 >= 470)) {
            switch 128 {
            case 9223372036854775807:
                p0 -= (p0 / ((p0) | 1))
            case 0:
                for fi12 := 0; fi12 < 19; fi12++ {
                    if true {
                        v13 := []int{fi12}
                        _ = v13
                    } else {
                        v14 := []int{(fi12 & fi12), fi12, (fi12 >> ((247) & 63)), (^(fi12))}
                        v15 := GenS1{F0: 11}
                        if (-128 != -128) {
                            p0 += (p0 << ((255) & 63))
                        }
                        _ = v14
                        _ = v15
                    }
                    if false {
                        p0 = (p0 | -1)
                    } else {
                        fmt.Println(p0, fi12)
                        p0 += 9223372036854775807
                    }
                    _ = fi12
                }
            case 2147483647:
                if true {
                    switch (("a" + "hello") + ("" + "\n\t")) {
                    case "hello":
                        p0 += p0
                    case "\n\t":
                        fmt.Println(p0)
                    case "foo bar":
                        fmt.Println(p0)
                    }
                    v16 := GenS2{F0: 1, F1: 2}
                    fmt.Println(p0)
                    _ = v16
                } else {
                    if true {
                        v17 := GenS3{F0: 2, F1: 128}
                        switch ("foo bar" + "") {
                        case "hello":
                            fmt.Println(p0)
                        }
                        fmt.Println(p0)
                        _ = v17
                    }
                    fmt.Println(p0)
                    for fi18 := 0; fi18 < 4; fi18++ {
                        for fi19 := 0; fi19 < 18; fi19++ {
                            p0 += 1
                            if false {
                                if false {
                                    fmt.Println(fi19, p0)
                                }
                                var v20 int64 = -6
                                _ = v20
                            }
                            _ = fi19
                        }
                        switch "hello" {
                        case "\n\t":
                            if true {
                                switch "a" {
                                case "":
                                    var v21 int32 = -81
                                    _ = v21
                                case "a":
                                    var v22 uint = 485
                                    _ = v22
                                }
                                if true {
                                    var v23 int32 = 1
                                    var v24 int16 = 42
                                    _ = v23
                                    _ = v24
                                } else {
                                    for fi25 := 0; fi25 < 14; fi25++ {
                                        var v26 float32 = (float32(float64(0.1)) * 3.14159)
                                        fmt.Println(fi18, fi25)
                                        v27 := []int{(fi25 / ((9223372036854775807) | 1)), fi25, 9223372036854775807}
                                        _ = fi25
                                        _ = v26
                                        _ = v27
                                    }
                                    if (true || (float32((1e10 / (1.5 - 3.14159))) > (1e10 + ((0.0 - 1e10) - 1.5)))) {
                                        p0 = (p0 >> ((124) & 63))
                                        if true {
                                            if (-1 != (fi18 << ((345) & 63))) {
                                                p0 = (p0 + p0)
                                                v28 := []int{9223372036854775807}
                                                fmt.Println(p0)
                                                _ = v28
                                            }
                                            var v29 string = (("" + "a") + "foo bar")
                                            fmt.Println(fi18, v29, p0)
                                            _ = v29
                                        }
                                    }
                                }
                                fmt.Println(fi18, p0)
                            } else {
                                switch 635 {
                                case -1:
                                    v30 := GenS3{F0: 2, F1: 255}
                                    _ = v30
                                case 32767:
                                    if false {
                                        var v31 uint = 255
                                        _ = v31
                                    }
                                case 497:
                                    v32 := GenS4{F0: 2147483647, F1: 42, F2: 78}
                                    _ = v32
                                default:
                                    switch -60 {
                                    case 9223372036854775807:
                                        for fi33 := 0; fi33 < 12; fi33++ {
                                            v34 := []int{}
                                            if ((p0 * p0) > p0) {
                                                p0 -= -621
                                            }
                                            _ = fi33
                                            _ = v34
                                        }
                                    case 32768:
                                        p0 += (p0 + (p0 / (((p0 ^ -707)) | 1)))
                                    default:
                                        for fi35 := 0; fi35 < 8; fi35++ {
                                            p0 -= -621
                                            v36 := GenS2{F0: 100, F1: 0}
                                            var v37 string = "hello"
                                            _ = fi35
                                            _ = v36
                                            _ = v37
                                        }
                                    }
                                }
                                var v38 uint8 = 2
                                _ = v38
                            }
                        case "hello":
                            for fi39 := 0; fi39 < 3; fi39++ {
                                fmt.Println(p0)
                                p0 = 0
                                _ = fi39
                            }
                        case "a":
                            if true {
                                v40 := []int{(fi18 * (fi18 + 166)), fi18, 9223372036854775807}
                                var v41 uint = 4294967295
                                p0 += (-(p0))
                                _ = v40
                                _ = v41
                            }
                        default:
                            var v42 uint8 = 12
                            _ = v42
                        }
                        _ = fi18
                    }
                }
            }
            if false {
                var v43 float64 = 0.1
                var v44 uint64 = 1
                _ = v43
                _ = v44
            }
            v45 := GenS5{F0: 127, F1: "a"}
            _ = v45
        } else {
            switch 2147483647 {
            case -1:
                if false {
                    switch "a" {
                    case "hello":
                        p0 -= (p0 << ((266) & 63))
                    case "\n\t":
                        var v46 bool = false
                        _ = v46
                    default:
                        var v47 bool = true
                        _ = v47
                    }
                    for fi48 := 0; fi48 < 19; fi48++ {
                        var v49 uint = 18446744073709551615
                        v49 = (v49 >> ((v49) & 63))
                        for fi50 := 0; fi50 < 7; fi50++ {
                            var v51 uint64 = 1
                            v52 := GenS3{F0: 1, F1: 203}
                            v49 = v49
                            v53 := GenS4{F0: 42, F1: 1, F2: -731}
                            _ = fi50
                            _ = v51
                            _ = v52
                            _ = v53
                        }
                        _ = fi48
                        _ = v49
                    }
                    v54 := GenS4{F0: 2147483647, F1: 0, F2: (p0 + (p0 * -260))}
                    _ = v54
                } else {
                    p0 += (p0 - p0)
                    var v55 bool = (42 < 32767)
                    if v55 {
                        for fi56 := 0; fi56 < 13; fi56++ {
                            if (false && false) {
                                if ("" > "\n\t") {
                                    for fi57 := 0; fi57 < 18; fi57++ {
                                        if true {
                                            v58 := []int{9223372036854775807, fi57}
                                            fmt.Println(fi57)
                                            _ = v58
                                        }
                                        fmt.Println(fi57)
                                        _ = fi57
                                    }
                                    p0 -= (p0 | 9223372036854775807)
                                } else {
                                    if (int8(p0) == -37) {
                                        for fi59 := 0; fi59 < 19; fi59++ {
                                            var v60 int64 = (p0 >> ((4294967295) & 63))
                                            _ = fi59
                                            _ = v60
                                        }
                                    }
                                    v61 := []int{}
                                    if v55 {
                                        var v62 uint32 = 896
                                        _ = v62
                                    } else {
                                        fmt.Println(p0)
                                    }
                                    _ = v61
                                }
                            } else {
                                v63 := []int{}
                                v64 := GenS3{F0: 0, F1: 255}
                                _ = v63
                                _ = v64
                            }
                            v65 := GenS6{F0: 731, F1: 42}
                            _ = fi56
                            _ = v65
                        }
                    } else {
                        p0 = p0
                    }
                    _ = v55
                }
            case 2147483647:
                var v66 uint64 = 4294967295
                _ = v66
            default:
                v67 := GenS7{F0: -1, F1: 294, F2: 65535}
                _ = v67
            }
            p0 += (p0 ^ (^(p0)))
            p0 = (p0 << ((18446744073709551615) & 63))
        }
    case -265:
        v68 := GenS5{F0: 32767, F1: "foo bar"}
        _ = v68
    case 520:
        p0 = 0
    default:
        p0 -= (p0 ^ -1)
    }
    var v69 int8 = 1
    v70 := GenS3{F0: 128, F1: 0}
    for fi71 := 0; fi71 < 11; fi71++ {
        fmt.Println(fi71)
        var v72 int32 = 42
        v72 -= (v72 % (((^(v72))) | 1))
        _ = fi71
        _ = v72
    }
    _ = p0
    _ = v69
    _ = v70
    return p0
}

func main() {
    var v1 int64 = 0
    var v2 float64 = float64((1.5 + 0.1))
    v3 := make(map[int]int)
    v4 := make(map[string]int)
    v73 := GenS8{F0: 32767}
    var v74 uint16 = 255
    for fi75 := 0; fi75 < 14; fi75++ {
        var v76 int8 = -101
        v2 += 0.1
        _ = fi75
        _ = v76
    }
    var v77 int16 = int16(v1)
    v74 = 65535
    v74 += (^(v74))
    var v78 int32 = 1
    if (85 > 18446744073709551615) {
        switch -21 {
        case 4294967295:
            if ((("\n\t" + "a") + (("hello" + "hello") + ("\n\t" + "foo bar"))) >= "") {
                var v79 int32 = v78
                if true {
                    for fi80 := 0; fi80 < 11; fi80++ {
                        v81 := GenS9{F0: fi80}
                        _ = fi80
                        _ = v81
                    }
                    fmt.Println(v1, v2)
                }
                for fi82 := 0; fi82 < 9; fi82++ {
                    if ((v74 * v74) >= 2) {
                        v83 := GenS6{F0: (^(v74)), F1: 42}
                        _ = v83
                    } else {
                        v84 := []int{}
                        v77 += (v77 + 1000)
                        v79 += (v78 / (((v78 >> ((18446744073709551615) & 63))) | 1))
                        _ = v84
                    }
                    switch "a" {
                    case "a":
                        v78 = int32(v74)
                    case "foo bar":
                        v85 := GenS10{F0: (fi82 * int(v74)), F1: "a"}
                        _ = v85
                    default:
                        switch "\n\t" {
                        case "a":
                            v86 := GenS10{F0: (fi82 % (((fi82 + fi82)) | 1)), F1: ""}
                            _ = v86
                        case "hello":
                            if (true && false) {
                                var v87 bool = ((genFn1(v1) < genFn1(v1)) || (!(!((genFn1(v1) < genFn1((v1 + 0))) && true))))
                                for fi88 := 0; fi88 < 14; fi88++ {
                                    var v89 uint64 = 0
                                    _ = fi88
                                    _ = v89
                                }
                                _ = v87
                            } else {
                                v90 := GenS6{F0: (v74 << ((255) & 63)), F1: v78}
                                v74 -= 479
                                switch (("\n\t" + "foo bar") + ("a" + "foo bar")) {
                                case "\n\t":
                                    var v91 uint16 = v74
                                    _ = v91
                                case "foo bar":
                                    var v92 int16 = v77
                                    _ = v92
                                case "a":
                                    v93 := GenS11{F0: 42, F1: -9223372036854775808, F2: 18446744073709551615}
                                    _ = v93
                                default:
                                    var v94 uint = 42
                                    _ = v94
                                }
                                _ = v90
                            }
                        case "foo bar":
                            for fi95 := 0; fi95 < 2; fi95++ {
                                switch fi82 {
                                case -393:
                                    v96 := GenS1{F0: 0}
                                    _ = v96
                                case 9223372036854775807:
                                    _ = len(v4)
                                default:
                                    if false {
                                        switch fi95 {
                                        case 127:
                                            v1 = int64(v74)
                                        case -9223372036854775808:
                                            v97 := []int{}
                                            _ = v97
                                        case 9223372036854775807:
                                            v77 = (v77 / ((v77) | 1))
                                        default:
                                            switch 0 {
                                            case 127:
                                                v98 := []int{2147483647}
                                                _ = v98
                                            case 1:
                                                var v99 uint32 = 127
                                                _ = v99
                                            default:
                                                v1 = v1
                                            }
                                        }
                                        fmt.Println(v78, v74, v79)
                                        v100 := GenS9{F0: fi95}
                                        _ = v100
                                    } else {
                                        if true {
                                            v2 += (1.5 - v2)
                                            var v101 int64 = 0
                                            fmt.Println(fi95)
                                            _ = v101
                                        } else {
                                            var v102 int16 = v77
                                            v103 := []int{fi95}
                                            _ = v102
                                            _ = v103
                                        }
                                        fmt.Println(v77, v2, fi82)
                                        v77 -= (v77 / ((v77) | 1))
                                    }
                                }
                                v77 = v77
                                if false {
                                    v104 := GenS9{F0: 128}
                                    switch fi95 {
                                    case 32768:
                                        var v105 float64 = ((v2 * v2) - 0.1)
                                        _ = v105
                                    case 2147483647:
                                        switch "hello" {
                                        case "foo bar":
                                            var v106 int16 = (v77 - int16(fi82))
                                            _ = v106
                                        case "hello":
                                            var v107 uint16 = v74
                                            _ = v107
                                        }
                                        switch (fi95 * 42) {
                                        case -9223372036854775808:
                                            var v108 uint8 = 128
                                            _ = v108
                                        }
                                        v79 -= (v79 / (((v78 - 1)) | 1))
                                    }
                                    v79 = v79
                                    _ = v104
                                }
                                _ = fi95
                            }
                        default:
                            fmt.Println(v77, v78)
                        }
                    }
                    switch (fi82 ^ 1) {
                    case 1:
                        v109 := []int{-957}
                        _ = v109
                    case -9223372036854775808:
                        if (genFn0((fi82 << ((342) & 63)), 4294967295, (v74 - (^(v74)))) < 35) {
                            v110 := GenS12{F0: 2, F1: v1, F2: genFn0(fi82, 0, (^(v74)))}
                            v111 := GenS6{F0: (^(v74)), F1: int32(v77)}
                            _ = v110
                            _ = v111
                        }
                    case 127:
                        for fi112 := 0; fi112 < 8; fi112++ {
                            var v113 int8 = -109
                            var v114 uint8 = 55
                            var v115 uint = 255
                            var v116 int16 = v77
                            _ = fi112
                            _ = v113
                            _ = v114
                            _ = v115
                            _ = v116
                        }
                    default:
                        if true {
                            var v117 int64 = 42
                            var v118 int8 = 100
                            _ = v117
                            _ = v118
                        }
                    }
                    _ = fi82
                }
                _ = v79
            } else {
                var v119 bool = (!false)
                if v119 {
                    for fi120 := 0; fi120 < 5; fi120++ {
                        var v121 int8 = 127
                        _ = fi120
                        _ = v121
                    }
                } else {
                    fmt.Println(v2)
                    var v122 bool = (0 == 42)
                    _ = v122
                }
                _ = v119
            }
        case -9223372036854775808:
            fmt.Println(v1, v74)
        case 809:
            for fi123 := 0; fi123 < 12; fi123++ {
                var v124 int = -2147483648
                v124 = fi123
                _ = fi123
                _ = v124
            }
        }
    } else {
        for fi125 := 0; fi125 < 4; fi125++ {
            var v126 int = -9223372036854775808
            _ = fi125
            _ = v126
        }
    }
    v127 := v4["hello"]
    fmt.Println(v74)
    if ("foo bar" == "a") {
        v127 = 9223372036854775807
    } else {
        if (-1 < (v78 % ((541) | 1))) {
            v128 := []int{-639, (v127 << ((1) & 63))}
            var v129 float64 = 0.0
            for fi130 := 0; fi130 < 11; fi130++ {
                v78 = int32(v1)
                _ = fi130
            }
            _ = v128
            _ = v129
        }
        var v131 int8 = -128
        _ = v131
    }
    if ((!false) || (v78 < v78)) {
        v127 += 2147483647
    }
    if false {
        if (!false) {
            var v132 int = v127
            _ = v132
        } else {
            if true {
                if true {
                    for fi133 := 0; fi133 < 11; fi133++ {
                        if false {
                            fmt.Println(v78, fi133)
                            var v134 int64 = -9223372036854775808
                            v127 = v127
                            _ = v134
                        } else {
                            v78 += 1
                        }
                        var v135 uint64 = genFn0(int(v77), 190, v74)
                        fmt.Println(fi133, fi133, v78)
                        _ = fi133
                        _ = v135
                    }
                }
            } else {
                v136 := GenS13{F0: (v74 & 2), F1: 459}
                v77 -= (^(v77))
                v137 := GenS14{F0: 2147483647}
                _ = v136
                _ = v137
            }
            switch 503 {
            case 9223372036854775807:
                v138 := []int{(v127 & (v127 * v127)), 223}
                _ = v138
            case -2147483648:
                v139 := []int{-1, int(v78)}
                _ = v139
            case 659:
                v140 := []int{532, v127, v127, v127}
                _ = v140
            }
        }
        switch "a" {
        case "a":
            v78 = v78
        case "foo bar":
            for fi141 := 0; fi141 < 8; fi141++ {
                v127 += (v127 % ((fi141) | 1))
                _ = fi141
            }
        case "\n\t":
            fmt.Println(v2, v78)
        default:
            var v142 float64 = v2
            _ = v142
        }
        _ = len(v3)
    }
    var v143 uint64 = 0
    if (-673 < int(v143)) {
        v144 := GenS4{F0: 99, F1: (-(v77)), F2: genFn1(v1)}
        _ = v144
    } else {
        switch (v127 << ((386) & 63)) {
        case 9223372036854775807:
            var v145 uint64 = 18446744073709551615
            _ = v145
        case 2147483647:
            v78 = 0
        default:
            var v146 bool = (v77 != v77)
            _ = v146
        }
    }
    var v147 bool = (-1 < -77)
    switch 32767 {
    case 128:
        fmt.Println(v2, v1, v1)
    default:
        v2 -= v2
    }
    if v147 {
        if (0 < 2) {
            for fi148 := 0; fi148 < 4; fi148++ {
                for fi149 := 0; fi149 < 16; fi149++ {
                    v150 := []int{fi149, (fi148 * v127)}
                    _ = fi149
                    _ = v150
                }
                v151 := []int{}
                v151 = append(v151, (v127 * 302))
                _ = fi148
                _ = v151
            }
        } else {
            v127 = (v127 << ((4294967295) & 63))
        }
    } else {
        v152 := GenS10{F0: v127, F1: "\n\t"}
        switch (("hello" + ("hello" + ((("a" + "a") + ("foo bar" + ("foo bar" + "a"))) + ("\n\t" + "foo bar")))) + "\n\t") {
        case "a":
            switch v127 {
            case 874:
                v153 := []int{}
                _ = v153
            case 9223372036854775807:
                if (!(42 > v143)) {
                    if (true && v147) {
                        switch v127 {
                        case 32767:
                            fmt.Println(v143)
                        case 690:
                            var v154 uint8 = 124
                            _ = v154
                        }
                        for fi155 := 0; fi155 < 15; fi155++ {
                            switch (v127 | fi155) {
                            case 1:
                                switch "hello" {
                                case "a":
                                    if true {
                                        v78 = (v78 ^ v78)
                                    } else {
                                        fmt.Println(v1, fi155)
                                        var v156 int16 = int16(v74)
                                        _ = v156
                                    }
                                case "":
                                    var v157 bool = (0 <= 42)
                                    _ = v157
                                case "foo bar":
                                    var v158 int64 = 281
                                    _ = v158
                                default:
                                    for fi159 := 0; fi159 < 1; fi159++ {
                                        v160 := GenS5{F0: v127, F1: "foo bar"}
                                        v161 := GenS1{F0: v77}
                                        _ = fi159
                                        _ = v160
                                        _ = v161
                                    }
                                }
                            case 725:
                                switch 128 {
                                case 0:
                                    v162 := v4["a"]
                                    _ = v162
                                case 32768:
                                    fmt.Println(v78, v147, v1)
                                }
                            default:
                                v163 := GenS13{F0: v74, F1: (v143 & uint64(v74))}
                                _ = v163
                            }
                            v164 := []int{(fi155 * (v127 % ((int(v74)) | 1))), v127, (v127 ^ 1), 1}
                            _ = fi155
                            _ = v164
                        }
                    } else {
                        fmt.Println(v127, v147, v143)
                    }
                    for fi165 := 0; fi165 < 13; fi165++ {
                        v166 := GenS2{F0: -117, F1: 42}
                        v78 -= 463
                        if true {
                            var v167 uint16 = (v74 - 81)
                            for fi168 := 0; fi168 < 10; fi168++ {
                                v169 := v4[("hello" + "foo bar")]
                                v167 += (v74 - (v74 % ((v74) | 1)))
                                _ = fi168
                                _ = v169
                            }
                            _ = v167
                        }
                        _ = fi165
                        _ = v166
                    }
                    var v170 int16 = v77
                    _ = v170
                } else {
                    for fi171 := 0; fi171 < 18; fi171++ {
                        if (!(true || v147)) {
                            v172 := []int{(^(fi171)), (fi171 >> ((604) & 63))}
                            v173 := GenS15{F0: v74, F1: v147}
                            _ = v172
                            _ = v173
                        }
                        var v174 int = 32767
                        for fi175 := 0; fi175 < 4; fi175++ {
                            _ = len(v3)
                            if v147 {
                                v143 -= (v143 + v143)
                                v147 = v147
                                var v176 int8 = int8(v78)
                                _ = v176
                            } else {
                                var v177 bool = false
                                fmt.Println(v77, v177)
                                _ = v177
                            }
                            v178 := GenS9{F0: 618}
                            if v147 {
                                var v179 bool = v147
                                fmt.Println(v77, v78, fi171)
                                var v180 int32 = v78
                                var v181 string = "hello"
                                _ = v179
                                _ = v180
                                _ = v181
                            } else {
                                fmt.Println(v143, v2)
                                if true {
                                    delete(v3, v174)
                                    v147 = (((v2 + float64(((0.1 / 3.14159) * 0.0))) + (((v2 + ((1.5 - (float64(0.1) / 0.1)) + float64(((0.1 * 1.5) - 0.0)))) * v2) * (v2 / float64(1e10)))) != -2.25)
                                } else {
                                    fmt.Println(v147)
                                    fmt.Println(v78)
                                }
                                fmt.Println(fi175, v77, v77)
                            }
                            _ = fi175
                            _ = v178
                        }
                        _ = fi171
                        _ = v174
                    }
                    switch v127 {
                    case 128:
                        for fi182 := 0; fi182 < 13; fi182++ {
                            for fi183 := 0; fi183 < 1; fi183++ {
                                var v184 uint8 = 255
                                v185 := []int{fi183, int(v143)}
                                _ = fi183
                                _ = v184
                                _ = v185
                            }
                            for fi186 := 0; fi186 < 1; fi186++ {
                                v187 := v4["foo bar"]
                                _ = fi186
                                _ = v187
                            }
                            var v188 float32 = 1.5
                            _ = fi182
                            _ = v188
                        }
                    case 982:
                        var v189 uint8 = 72
                        _ = v189
                    default:
                        switch int(v143) {
                        case 4294967295:
                            v190 := GenS16{F0: genFn1(v1), F1: (v127 << ((4294967295) & 63)), F2: 9223372036854775807}
                            _ = v190
                        }
                    }
                }
            default:
                for fi191 := 0; fi191 < 3; fi191++ {
                    v192 := []int{}
                    _ = fi191
                    _ = v192
                }
            }
        case "foo bar":
            v193 := GenS13{F0: 65535, F1: 285}
            _ = v193
        }
        if false {
            v194 := GenS10{F0: (v127 & 32767), F1: ("\n\t" + ("foo bar" + (("foo bar" + ("\n\t" + "\n\t")) + "")))}
            v127 -= (v127 % (((^(v127))) | 1))
            v195 := []int{}
            _ = v194
            _ = v195
        } else {
            for fi196 := 0; fi196 < 7; fi196++ {
                v197 := GenS16{F0: (v1 - 0), F1: (fi196 & 0), F2: (v1 | (v1 - v1))}
                var v198 string = "hello"
                v143 = 260
                _ = fi196
                _ = v197
                _ = v198
            }
            var v199 uint8 = 169
            _ = v199
        }
        _ = v152
    }
    fmt.Println("v1:", v1)
    fmt.Println("v127:", v127)
    fmt.Println("v143:", v143)
    fmt.Println("v147:", v147)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v74:", v74)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v127
    _ = v143
    _ = v147
    _ = v2
    _ = v3
    _ = v4
    _ = v73
    _ = v74
    _ = v77
    _ = v78
}
