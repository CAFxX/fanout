package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 uint; F1 int16 }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int }
type GenS9 struct { F0 bool; F1 int64; F2 bool }
type GenS10 struct { F0 int16; F1 uint64 }
type GenS11 struct { F0 int64; F1 string; F2 int }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 uint; F1 int16 }
type GenS14 struct { F0 uint64 }
func genFn0() int16 {
    var v6 float64 = 1.5
    var v7 int = 1
    switch 127 {
    case 631:
        v8 := GenS1{F0: 2, F1: 65535}
        _ = v8
    default:
        fmt.Println(v7)
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        switch "foo bar" {
        case "":
            v10 := []int{v7, v7, (fi9 ^ (v7 + v7))}
            _ = v10
        case "\n\t":
            switch fi9 {
            case 611:
                var v11 uint = 718
                _ = v11
            case -347:
                var v12 int8 = int8(v7)
                _ = v12
            case 1:
                v7 = 32768
            default:
                v7 = (v7 - v7)
            }
        default:
            v13 := GenS2{F0: 788}
            _ = v13
        }
        var v14 int32 = -616
        switch "foo bar" {
        case "a":
            v7 = (v7 - v7)
        case "":
            v15 := []int{127}
            _ = v15
        default:
            var v16 int32 = v14
            _ = v16
        }
        _ = fi9
        _ = v14
    }
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 uint8, p1 int16) uint16 {
    v17 := GenS2{F0: -580}
    v18 := []int{int(p0)}
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return 204
}

func main() {
    var v1 uint16 = 42
    var v2 int16 = -1
    var v3 float32 = ((0.1 + 1e10) / 3.14159)
    v4 := make(map[int]string)
    var v5 string = ("\n\t" + ("a" + "foo bar"))
    if true {
        v19 := GenS2{F0: (v2 * (v2 / ((int16(v1)) | 1)))}
        _ = v19
    } else {
        switch ((("hello" + "foo bar") + v5) + (v5 + "a")) {
        case "":
            v5 = (v5 + "\n\t")
        case "a":
            for fi20 := 0; fi20 < 16; fi20++ {
                if false {
                    switch (v5 + (("a" + v5) + (v5 + (((((v5 + v5) + "") + (((v5 + v5) + "a") + v5)) + v5) + (v5 + (v5 + v5)))))) {
                    case "hello":
                        v21 := []int{fi20}
                        _ = v21
                    case "\n\t":
                        v5 = v5
                    }
                }
                _ = fi20
            }
        case "foo bar":
            if (42 < 0) {
                v22 := GenS1{F0: 239, F1: 408}
                _ = v22
            } else {
                if (!(1 == -306)) {
                    fmt.Println(v3)
                    v23 := GenS2{F0: -385}
                    _ = v23
                } else {
                    v24 := GenS1{F0: 127, F1: 2}
                    v25 := GenS2{F0: v2}
                    var v26 string = ((("foo bar" + "hello") + v5) + "")
                    _ = v24
                    _ = v25
                    _ = v26
                }
                fmt.Println(v1, v1, v1)
                v27 := GenS2{F0: (v2 * v2)}
                _ = v27
            }
        default:
            for fi28 := 0; fi28 < 1; fi28++ {
                var v29 uint8 = 128
                fmt.Println(v29, fi28, fi28)
                if (0 != 0) {
                    v30 := v4[(fi28 % ((1) | 1))]
                    if (!((!false) && true)) {
                        v3 -= -2.25
                        fmt.Println(v29)
                        for fi31 := 0; fi31 < 3; fi31++ {
                            var v32 float64 = (((((0.1 - 0.1) + float64((v3 / v3))) / 1e10) / float64(1e10)) - float64(v3))
                            _ = fi31
                            _ = v32
                        }
                    } else {
                        delete(v4, fi28)
                        fmt.Println(fi28, v30, v30)
                        fmt.Println(v29)
                    }
                    v29 = (v29 | (^(v29)))
                    _ = v30
                } else {
                    v29 = (v29 & (^(v29)))
                    switch v5 {
                    case "foo bar":
                        var v33 int64 = 0
                        _ = v33
                    case "hello":
                        fmt.Println(v2, v1)
                    default:
                        if true {
                            fmt.Println(v2, fi28)
                            v34 := GenS3{F0: v29, F1: ((!(42 == 42)) && (genFn0() > (v2 * v2)))}
                            _ = v34
                        }
                    }
                    var v35 int16 = (-(v2))
                    _ = v35
                }
                _ = fi28
                _ = v29
            }
        }
        v36 := GenS2{F0: (^(v2))}
        _ = v36
    }
    if ((v5 == v5) || false) {
        v1 += v1
        v5 = "hello"
        if true {
            var v37 int16 = 1000
            _ = v37
        } else {
            v38 := v4[2147483647]
            _ = v38
        }
    } else {
        fmt.Println(v3, v5)
        v2 = (v2 / ((v2) | 1))
        for fi39 := 0; fi39 < 14; fi39++ {
            for fi40 := 0; fi40 < 3; fi40++ {
                if (55 != 100) {
                    var v41 float64 = (1e10 * 1.5)
                    if true {
                        v5 = v5
                    }
                    fmt.Println(v5, v3, v2)
                    v3 = ((v3 + 3.14159) * v3)
                    _ = v41
                } else {
                    if false {
                        var v42 string = v5
                        var v43 float32 = ((v3 / 3.14159) - float32(((0.1 / 1.0) * (((3.14159 + 1.5) + 3.14159) / 1.0))))
                        _ = v42
                        _ = v43
                    } else {
                        if (true && true) {
                            v1 = genFn1(0, (v2 / (((v2 | 0)) | 1)))
                        } else {
                            v1 += v1
                            v1 = (v1 | 626)
                        }
                    }
                    v3 = float32(0.1)
                    fmt.Println(v3, v1, fi39)
                    fmt.Println(fi40)
                }
                if false {
                    var v44 uint16 = v1
                    if true {
                        v5 = ("" + v5)
                    }
                    _ = v44
                }
                if (v5 <= (v5 + ((((v5 + "") + (v5 + v5)) + "hello") + ""))) {
                    v3 = (v3 * v3)
                }
                _ = fi40
            }
            var v45 uint64 = 274
            v1 = v1
            _ = fi39
            _ = v45
        }
    }
    var v46 uint16 = 42
    fmt.Println(v3)
    switch -901 {
    case -518:
        switch 32767 {
        case 2147483647:
            switch -162 {
            case 42:
                if (v2 < 1) {
                    v3 = v3
                } else {
                    v1 -= (^(v1))
                    var v47 float32 = v3
                    if (true || ((!(v46 == genFn1(uint8(v46), (v2 ^ 1000)))) && false)) {
                        for fi48 := 0; fi48 < 4; fi48++ {
                            v47 -= float32(float64(float32(1e10)))
                            v49 := GenS3{F0: 0, F1: false}
                            for fi50 := 0; fi50 < 19; fi50++ {
                                var v51 float32 = (((v47 / v47) * v47) / 3.14159)
                                _ = fi50
                                _ = v51
                            }
                            _ = fi48
                            _ = v49
                        }
                    }
                    _ = v47
                }
            case -866:
                v2 = (v2 & v2)
            }
        }
    case -9223372036854775808:
        v5 = "foo bar"
    default:
        var v52 float32 = v3
        _ = v52
    }
    switch "a" {
    case "a":
        if false {
            if (665 < 2) {
                delete(v4, int(v2))
                var v53 int = -9223372036854775808
                _ = v53
            } else {
                if true {
                    v54 := []int{-2147483648, -373, -94, -1}
                    if (!true) {
                        switch "\n\t" {
                        case "foo bar":
                            for fi55 := 0; fi55 < 11; fi55++ {
                                v56 := GenS4{F0: uint(v1), F1: (v2 << ((1) & 63))}
                                var v57 string = "a"
                                _ = fi55
                                _ = v56
                                _ = v57
                            }
                        case "":
                            if false {
                                switch -209 {
                                case 42:
                                    fmt.Println(v1, v1)
                                case 4294967295:
                                    v3 = v3
                                }
                                fmt.Println(v1, v46, v2)
                                var v58 int16 = v2
                                _ = v58
                            }
                        }
                    }
                    _ = v54
                } else {
                    var v59 uint8 = 2
                    _ = v59
                }
                var v60 uint32 = 0
                _ = v60
            }
            if false {
                var v61 int32 = -2147483648
                fmt.Println(v1, v46, v46)
                _ = v61
            }
        } else {
            v3 = v3
            v62 := GenS5{F0: 455, F1: 18446744073709551615}
            v5 = v5
            _ = v62
        }
    case "":
        if false {
            v63 := GenS6{F0: 42, F1: -430, F2: 65535}
            var v64 int32 = 251
            _ = v63
            _ = v64
        } else {
            v65 := []int{133, 9223372036854775807, int(v2), 1}
            _ = v65
        }
    default:
        switch 2147483647 {
        case 32767:
            for fi66 := 0; fi66 < 18; fi66++ {
                v5 = "hello"
                if (0 == 42) {
                    fmt.Println(v5, fi66)
                }
                for fi67 := 0; fi67 < 18; fi67++ {
                    var v68 int32 = 42
                    _ = fi67
                    _ = v68
                }
                _ = fi66
            }
        default:
            if (true && (v5 == v5)) {
                switch -9223372036854775808 {
                case 856:
                    v5 = v5
                }
            } else {
                for fi69 := 0; fi69 < 6; fi69++ {
                    switch -1 {
                    case -55:
                        fmt.Println(v3, v46, v46)
                    case 127:
                        var v70 string = v5
                        _ = v70
                    default:
                        v5 = (v5 + "\n\t")
                    }
                    v71 := GenS7{F0: v5}
                    _ = fi69
                    _ = v71
                }
                if (-2.25 == v3) {
                    switch ("foo bar" + "hello") {
                    case "":
                        var v72 int = 32767
                        _ = v72
                    case "a":
                        if false {
                            var v73 int8 = -1
                            if false {
                                switch -1 {
                                case 816:
                                    var v74 uint = 390
                                    _ = v74
                                case 0:
                                    v2 += (-(v2))
                                case -2147483648:
                                    fmt.Println(v3, v5, v46)
                                }
                            } else {
                                fmt.Println(v1, v5)
                                v1 = (v46 & (v46 ^ v1))
                                for fi75 := 0; fi75 < 4; fi75++ {
                                    v4[(fi75 << ((42) & 63))] = v5
                                    for fi76 := 0; fi76 < 19; fi76++ {
                                        v77 := GenS8{F0: -387}
                                        if (145 > 42) {
                                            if false {
                                                v78 := GenS9{F0: (v73 != (v73 % ((v73) | 1))), F1: 42, F2: (((!(27 >= 42)) || false) && false)}
                                                _ = v78
                                            } else {
                                                var v79 int = fi75
                                                fmt.Println(v46, v46, v79)
                                                _ = v79
                                            }
                                            switch fi76 {
                                            case 127:
                                                v1 = 255
                                            case 1:
                                                fmt.Println(fi75)
                                            default:
                                                fmt.Println(v2, v1, fi75)
                                            }
                                            v80 := []int{32767, (fi75 % (((fi76 - 167)) | 1)), 0}
                                            _ = v80
                                        }
                                        v81 := []int{fi76, (^(fi76))}
                                        _ = fi76
                                        _ = v77
                                        _ = v81
                                    }
                                    switch v5 {
                                    case "a":
                                        v5 = ("" + "")
                                    case "\n\t":
                                        fmt.Println(v1, v46)
                                    }
                                    _ = fi75
                                }
                            }
                            _ = v73
                        } else {
                            if (int64(v2) >= 42) {
                                var v82 int = 42
                                v3 -= v3
                                v83 := []int{v82, 629, v82}
                                _ = v82
                                _ = v83
                            }
                            if true {
                                v46 -= (^(v46))
                                for fi84 := 0; fi84 < 9; fi84++ {
                                    var v85 uint64 = 0
                                    v2 += genFn0()
                                    _ = fi84
                                    _ = v85
                                }
                                if (!true) {
                                    if true {
                                        v5 = "hello"
                                        var v86 int16 = -600
                                        v87 := GenS4{F0: 18446744073709551615, F1: (v86 | (^(v2)))}
                                        _ = v86
                                        _ = v87
                                    } else {
                                        v2 = v2
                                        if true {
                                            var v88 int32 = -2147483648
                                            _ = v88
                                        } else {
                                            v89 := GenS10{F0: 1, F1: uint64(v1)}
                                            switch v5 {
                                            case "foo bar":
                                                v3 -= 3.14159
                                            }
                                            _ = v89
                                        }
                                    }
                                    fmt.Println(v2)
                                    var v90 string = (v5 + ("foo bar" + v5))
                                    _ = v90
                                } else {
                                    for fi91 := 0; fi91 < 15; fi91++ {
                                        var v92 uint = 4294967295
                                        var v93 uint8 = 211
                                        _ = fi91
                                        _ = v92
                                        _ = v93
                                    }
                                    v4[9223372036854775807] = (v5 + v5)
                                }
                            }
                        }
                    }
                    var v94 uint = 0
                    v95 := GenS5{F0: -1, F1: 18446744073709551615}
                    _ = v94
                    _ = v95
                } else {
                    if true {
                        for fi96 := 0; fi96 < 5; fi96++ {
                            for fi97 := 0; fi97 < 18; fi97++ {
                                var v98 uint64 = 708
                                v1 = v1
                                _ = fi97
                                _ = v98
                            }
                            _ = fi96
                        }
                        v5 = v5
                    }
                }
                var v99 uint = 893
                _ = v99
            }
        }
    }
    switch v5 {
    case "a":
        delete(v4, 9223372036854775807)
    case "\n\t":
        v3 = v3
    case "":
        if (!(18446744073709551615 < 4294967295)) {
            var v100 int64 = 1
            _ = v100
        } else {
            var v101 uint8 = 2
            _ = v101
        }
    default:
        v102 := GenS9{F0: (-651 > 1), F1: 9223372036854775807, F2: (!(!false))}
        _ = v102
    }
    if ((!false) || true) {
        for fi103 := 0; fi103 < 8; fi103++ {
            if (18446744073709551615 > 4294967295) {
                switch (fi103 % (((fi103 | fi103)) | 1)) {
                case 127:
                    for fi104 := 0; fi104 < 15; fi104++ {
                        if (!((!(!(false && false))) || false)) {
                            v105 := GenS10{F0: v2, F1: 2}
                            v2 -= (-(v2))
                            var v106 int64 = 0
                            _ = v105
                            _ = v106
                        } else {
                            if false {
                                v2 -= (v2 >> ((65535) & 63))
                                v107 := v4[fi103]
                                _ = v107
                            }
                            v1 = 32767
                        }
                        switch (fi103 & fi104) {
                        case -2147483648:
                            v3 = 1e10
                            if (false && true) {
                                if false {
                                    var v108 float64 = 3.14159
                                    v109 := GenS6{F0: 169, F1: -1, F2: 4294967295}
                                    _ = v108
                                    _ = v109
                                }
                                v110 := GenS6{F0: 168, F1: -1, F2: 2147483647}
                                _ = v110
                            }
                        case 266:
                            v1 += 2
                        case 1:
                            v111 := GenS7{F0: ("" + "a")}
                            _ = v111
                        }
                        v5 = v5
                        _ = fi104
                    }
                default:
                    if false {
                        v112 := GenS11{F0: -229, F1: v5, F2: fi103}
                        _ = v112
                    } else {
                        v3 -= (1e10 - (v3 + 0.1))
                        for fi113 := 0; fi113 < 8; fi113++ {
                            fmt.Println(v46, fi103)
                            switch (v5 + "a") {
                            case "foo bar":
                                _ = len(v4)
                                if (true || true) {
                                    var v114 uint = 11
                                    v115 := GenS12{F0: (v2 / ((v2) | 1))}
                                    v114 = (v114 | 4294967295)
                                    if true {
                                        switch "foo bar" {
                                        case "\n\t":
                                            if (!false) {
                                                v116 := []int{(fi113 % ((fi113) | 1)), 2147483647}
                                                v5 = (("" + ("" + (v5 + ""))) + (v5 + ("" + "a")))
                                                _ = v116
                                            }
                                        case "a":
                                            var v117 uint32 = 1
                                            _ = v117
                                        case "":
                                            if ((79 != uint32(v114)) && true) {
                                                fmt.Println(v5)
                                                var v118 int = fi113
                                                _ = v118
                                            }
                                        }
                                        if true {
                                            v119 := GenS6{F0: 2, F1: -2147483648, F2: 65535}
                                            _ = v119
                                        } else {
                                            v120 := []int{(fi113 / ((32768) | 1)), (fi113 << ((1) & 63)), fi113}
                                            _ = v120
                                        }
                                        switch fi103 {
                                        case 127:
                                            v121 := GenS5{F0: -2147483648, F1: 661}
                                            _ = v121
                                        }
                                    } else {
                                        fmt.Println(v114, fi103, fi103)
                                        var v122 float64 = 0.1
                                        _ = v122
                                    }
                                    _ = v114
                                    _ = v115
                                } else {
                                    delete(v4, fi113)
                                }
                                var v123 uint64 = 360
                                _ = v123
                            }
                            _ = fi113
                        }
                    }
                }
                for fi124 := 0; fi124 < 17; fi124++ {
                    v125 := GenS6{F0: 255, F1: -120, F2: 65535}
                    if false {
                        fmt.Println(fi103, fi124)
                    }
                    _ = fi124
                    _ = v125
                }
                v5 = v5
            } else {
                v126 := []int{(fi103 + 2147483647)}
                if len(v126) >= 2 {
                    v126 = v126[1:len(v126)]
                }
                switch (fi103 & fi103) {
                case -734:
                    delete(v4, 501)
                default:
                    v3 += v3
                }
                _ = v126
            }
            for fi127 := 0; fi127 < 5; fi127++ {
                switch ((v5 + v5) + v5) {
                case "hello":
                    delete(v4, fi127)
                default:
                    var v128 float64 = 0.1
                    _ = v128
                }
                v3 = (v3 + 3.14159)
                _ = fi127
            }
            _ = fi103
        }
    }
    if (80 <= 0) {
        if false {
            delete(v4, 1)
        } else {
            var v129 int64 = 42
            for fi130 := 0; fi130 < 9; fi130++ {
                for fi131 := 0; fi131 < 17; fi131++ {
                    v3 += (0.1 - 1e10)
                    var v132 uint32 = 65535
                    fmt.Println(fi131, v2)
                    _ = fi131
                    _ = v132
                }
                _ = fi130
            }
            _ = v129
        }
        if true {
            fmt.Println(v5)
            var v133 uint64 = 18446744073709551615
            v134 := []int{}
            _ = v133
            _ = v134
        } else {
            var v135 uint32 = 1
            v136 := GenS12{F0: -542}
            _ = v135
            _ = v136
        }
        fmt.Println(v5, v2)
    }
    var v137 float64 = 3.14159
    if (v5 <= "") {
        v138 := []int{-633}
        _ = v138
    }
    for fi139 := 0; fi139 < 9; fi139++ {
        v140 := GenS13{F0: 816, F1: v2}
        var v141 float32 = float32(0.1)
        if ((v3 + float32(float64(3.14159))) < (-2.25 - (v3 * 0.1))) {
            if (42 < 255) {
                v142 := []int{}
                fmt.Println(v137, v141, v1)
                v1 += 2
                _ = v142
            } else {
                var v143 uint = 2
                if true {
                    fmt.Println(v46, fi139)
                    v4[(fi139 | fi139)] = "a"
                    fmt.Println(v5)
                }
                switch 0 {
                case 127:
                    for fi144 := 0; fi144 < 5; fi144++ {
                        fmt.Println(v1, fi144)
                        if false {
                            v145 := GenS4{F0: v143, F1: (^(v2))}
                            _ = v145
                        } else {
                            v137 = ((((((v137 + -2.25) * 1e10) + v137) / (v137 / ((float64(-2.25) + v137) + (v137 - float64((v141 * (v3 / 1.0))))))) + (float64((((v141 + v3) / 0.1) / float32(v137))) - v137)) - v137)
                            v146 := GenS14{F0: 2}
                            _ = v146
                        }
                        v1 = (v46 << ((v143) & 63))
                        _ = fi144
                    }
                case 740:
                    v4[(fi139 >> ((v143) & 63))] = (v5 + v5)
                case 403:
                    var v147 uint = 4294967295
                    _ = v147
                }
                _ = v143
            }
            v148 := []int{fi139, (fi139 * fi139)}
            _ = v148
        }
        _ = fi139
        _ = v140
        _ = v141
    }
    for fi149 := 0; fi149 < 13; fi149++ {
        if (genFn0() != (-(v2))) {
            v46 += genFn1(1, -32768)
            v5 = "a"
            v137 -= (float64(3.14159) / ((1e10 * v137) + 3.14159))
        }
        var v150 int64 = 746
        for fi151 := 0; fi151 < 17; fi151++ {
            v5 = ("" + v5)
            v1 = (v46 | v46)
            if false {
                v152 := []int{547, int(v1), (fi149 ^ fi151), -9223372036854775808}
                var v153 uint32 = 821
                _ = v152
                _ = v153
            }
            _ = fi151
        }
        _ = fi149
        _ = v150
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v137:", v137)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v137
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
}
