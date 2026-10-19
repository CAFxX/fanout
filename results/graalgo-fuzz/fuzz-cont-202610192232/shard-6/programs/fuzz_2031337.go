package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 int64; F1 int8; F2 uint }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int32; F1 uint8 }
type GenS8 struct { F0 uint8; F1 int64 }
type GenS9 struct { F0 string; F1 int; F2 int64 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 int; F1 uint16 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 int; F1 int16 }
type GenS16 struct { F0 uint64; F1 bool; F2 uint32 }
func genFn0(p0 int32, p1 uint8) int64 {
    fmt.Println(p0)
    var v7 int8 = -28
    if true {
        if false {
            if (!(0 < 9223372036854775807)) {
                if false {
                    v8 := GenS1{F0: 2}
                    switch 0 {
                    case 1:
                        v7 += (v7 | (^(v7)))
                    case 0:
                        switch (("" + "\n\t") + "foo bar") {
                        case "\n\t":
                            var v9 bool = (42 > 32767)
                            _ = v9
                        case "a":
                            v10 := []int{2147483647, -17, 9223372036854775807}
                            _ = v10
                        default:
                            var v11 int = 109
                            _ = v11
                        }
                    case -901:
                        var v12 float32 = 0.0
                        _ = v12
                    default:
                        v13 := []int{128, 600, 2147483647}
                        _ = v13
                    }
                    _ = v8
                }
                var v14 int32 = 0
                for fi15 := 0; fi15 < 14; fi15++ {
                    if true {
                        v16 := GenS2{F0: fi15}
                        _ = v16
                    } else {
                        v7 = v7
                        fmt.Println(fi15)
                        v17 := []int{(fi15 ^ (fi15 * fi15)), 431, (fi15 ^ 42), (fi15 + 128)}
                        _ = v17
                    }
                    if (!true) {
                        switch "a" {
                        case "\n\t":
                            fmt.Println(fi15)
                        case "foo bar":
                            p1 = (p1 >> ((816) & 63))
                        default:
                            for fi18 := 0; fi18 < 11; fi18++ {
                                v19 := GenS3{F0: 100, F1: 1000}
                                _ = fi18
                                _ = v19
                            }
                        }
                        p0 += (^(v14))
                    }
                    v20 := GenS4{F0: 4294967295, F1: p1}
                    _ = fi15
                    _ = v20
                }
                _ = v14
            } else {
                v21 := GenS2{F0: 435}
                var v22 uint8 = (p1 * 255)
                for fi23 := 0; fi23 < 10; fi23++ {
                    var v24 int32 = p0
                    if false {
                        var v25 int16 = 42
                        _ = v25
                    } else {
                        switch "\n\t" {
                        case "\n\t":
                            for fi26 := 0; fi26 < 18; fi26++ {
                                fmt.Println(p1, v22)
                                v27 := []int{-9223372036854775808}
                                if (("foo bar" <= "hello") && true) {
                                    if true {
                                        if len(v27) >= 2 {
                                            v27 = v27[1:len(v27)]
                                        }
                                        if len(v27) > 0 {
                                            v27[0] = fi26
                                        }
                                        _ = len(v27)
                                    }
                                    var v28 float64 = 1.5
                                    _ = v28
                                } else {
                                    switch "\n\t" {
                                    case "hello":
                                        v29 := GenS2{F0: fi26}
                                        if false {
                                            v27 = append(v27, int(v7))
                                        } else {
                                            v30 := GenS4{F0: 0, F1: 217}
                                            _ = v30
                                        }
                                        _ = v29
                                    case "foo bar":
                                        var v31 int8 = (v7 / ((127) | 1))
                                        _ = v31
                                    }
                                    fmt.Println(fi26, v24)
                                    v32 := GenS2{F0: 2147483647}
                                    var v33 bool = false
                                    _ = v32
                                    _ = v33
                                }
                                _ = fi26
                                _ = v27
                            }
                        case "hello":
                            var v34 uint8 = p1
                            _ = v34
                        case "":
                            fmt.Println(p0)
                        default:
                            var v35 int = 0
                            _ = v35
                        }
                        p1 = (p1 << ((4294967295) & 63))
                    }
                    _ = fi23
                    _ = v24
                }
                _ = v21
                _ = v22
            }
            var v36 uint16 = 0
            _ = v36
        } else {
            var v37 uint16 = 1
            for fi38 := 0; fi38 < 16; fi38++ {
                if true {
                    v39 := GenS2{F0: fi38}
                    fmt.Println(fi38, p1, v37)
                    _ = v39
                } else {
                    var v40 int = 32768
                    var v41 uint8 = 2
                    var v42 int16 = 977
                    _ = v40
                    _ = v41
                    _ = v42
                }
                _ = fi38
            }
            _ = v37
        }
    }
    fmt.Println(v7)
    _ = p0
    _ = p1
    _ = v7
    return -614
}
func genFn1(p0 int32) int16 {
    p0 = (p0 / ((508) | 1))
    for fi43 := 0; fi43 < 10; fi43++ {
        v44 := []int{fi43, (fi43 & 2147483647), -2147483648, -9223372036854775808}
        for fi45 := 0; fi45 < 3; fi45++ {
            p0 += p0
            _ = fi45
        }
        var v46 int16 = 42
        _ = fi43
        _ = v44
        _ = v46
    }
    v47 := GenS2{F0: 9223372036854775807}
    v48 := []int{1, -9223372036854775808}
    if ("hello" > "") {
        if true {
            p0 = 0
        } else {
            if len(v48) >= 2 {
                v48 = v48[1:len(v48)]
            }
            switch ("foo bar" + "hello") {
            case "hello":
                fmt.Println(p0)
            }
            fmt.Println(p0)
        }
        p0 += p0
        if true {
            if (false || true) {
                switch ((("hello" + ("hello" + ("a" + (("hello" + "foo bar") + ("\n\t" + ("\n\t" + "")))))) + "\n\t") + "hello") {
                case "hello":
                    if len(v48) >= 2 {
                        v48 = v48[1:len(v48)]
                    }
                case "foo bar":
                    if (false || false) {
                        if (32767 < 484) {
                            if true {
                                v49 := GenS5{F0: 9223372036854775807, F1: 42, F2: 0}
                                p0 = (p0 & (p0 / (((-(p0))) | 1)))
                                _ = v49
                            } else {
                                if len(v48) >= 2 {
                                    v48 = v48[1:len(v48)]
                                }
                                var v50 uint = 0
                                for fi51 := 0; fi51 < 16; fi51++ {
                                    v52 := GenS3{F0: -11, F1: 608}
                                    _ = len(v48)
                                    v53 := GenS6{F0: 648}
                                    _ = fi51
                                    _ = v52
                                    _ = v53
                                }
                                _ = v50
                            }
                            for fi54 := 0; fi54 < 7; fi54++ {
                                var v55 uint = 1
                                _ = fi54
                                _ = v55
                            }
                            for fi56 := 0; fi56 < 8; fi56++ {
                                p0 = (p0 / (((-(p0))) | 1))
                                _ = fi56
                            }
                        } else {
                            p0 = (^(p0))
                        }
                        for fi57 := 0; fi57 < 1; fi57++ {
                            var v58 string = "a"
                            _ = fi57
                            _ = v58
                        }
                        var v59 uint16 = 32767
                        _ = v59
                    }
                case "a":
                    var v60 int64 = genFn0(p0, 42)
                    _ = v60
                default:
                    v61 := GenS2{F0: 383}
                    _ = v61
                }
                var v62 uint64 = 4294967295
                _ = v62
            } else {
                var v63 int16 = -469
                for fi64 := 0; fi64 < 12; fi64++ {
                    v65 := GenS1{F0: 1}
                    switch "foo bar" {
                    case "":
                        v48 = append(v48, -848)
                    case "\n\t":
                        var v66 uint8 = 1
                        _ = v66
                    case "foo bar":
                        if false {
                            v67 := GenS7{F0: p0, F1: 255}
                            var v68 float64 = 0.0
                            _ = v67
                            _ = v68
                        } else {
                            if len(v48) >= 2 {
                                v48 = v48[1:len(v48)]
                            }
                            if len(v48) >= 2 {
                                v48 = v48[1:len(v48)]
                            }
                            switch -1 {
                            case -853:
                                switch (fi64 | fi64) {
                                case 106:
                                    if (!true) {
                                        v69 := GenS7{F0: (p0 >> ((1) & 63)), F1: 2}
                                        var v70 uint = 0
                                        fmt.Println(v70, v63)
                                        _ = v69
                                        _ = v70
                                    } else {
                                        if (2 >= 255) {
                                            fmt.Println(fi64)
                                        }
                                        var v71 float64 = ((3.14159 + 3.14159) * 3.14159)
                                        if (true || (188 <= 128)) {
                                            var v72 string = (("hello" + "foo bar") + ("a" + "a"))
                                            if true {
                                                fmt.Println(v71)
                                                var v73 int = 954
                                                fmt.Println(p0)
                                                _ = v73
                                            }
                                            _ = v72
                                        } else {
                                            v74 := GenS3{F0: 0, F1: int16(fi64)}
                                            v75 := GenS8{F0: 1, F1: -1}
                                            _ = v74
                                            _ = v75
                                        }
                                        _ = v71
                                    }
                                }
                            case 4294967295:
                                if false {
                                    for fi76 := 0; fi76 < 20; fi76++ {
                                        fmt.Println(v63, fi76)
                                        v77 := GenS6{F0: 65535}
                                        _ = fi76
                                        _ = v77
                                    }
                                    switch 2147483647 {
                                    case 128:
                                        p0 = p0
                                    }
                                    v78 := GenS7{F0: 0, F1: 165}
                                    _ = v78
                                }
                            }
                        }
                    default:
                        v63 += 69
                    }
                    var v79 string = ("foo bar" + "a")
                    _ = fi64
                    _ = v65
                    _ = v79
                }
                _ = v63
            }
            if false {
                for fi80 := 0; fi80 < 2; fi80++ {
                    for fi81 := 0; fi81 < 16; fi81++ {
                        var v82 uint8 = 0
                        fmt.Println(v82, fi81, fi80)
                        _ = fi81
                        _ = v82
                    }
                    _ = len(v48)
                    _ = fi80
                }
            }
            for fi83 := 0; fi83 < 9; fi83++ {
                if (1e10 < float64(1e10)) {
                    var v84 int8 = -106
                    _ = v84
                } else {
                    v85 := GenS9{F0: "a", F1: fi83, F2: -9223372036854775808}
                    _ = v85
                }
                if (77 == 2) {
                    var v86 float32 = float32(float64(-2.25))
                    v87 := GenS10{F0: 40}
                    _ = v86
                    _ = v87
                } else {
                    p0 = p0
                    p0 = p0
                    if false {
                        if len(v48) > 0 {
                            v48[0] = (fi83 | fi83)
                        }
                    }
                }
                fmt.Println(p0)
                _ = fi83
            }
        }
    } else {
        var v88 int = 9223372036854775807
        v89 := GenS6{F0: 634}
        _ = v88
        _ = v89
    }
    _ = p0
    _ = v47
    _ = v48
    return -32768
}

func main() {
    var v1 uint = 255
    var v2 uint32 = 0
    v3 := []int{-2147483648, -103}
    var v4 string = "a"
    var v5 int64 = 1
    var v6 float32 = 1e10
    v90 := GenS11{F0: 683}
    var v91 uint = (v1 * v1)
    var v92 uint32 = (v2 / (((v2 + 2147483647)) | 1))
    v93 := GenS9{F0: (("" + "foo bar") + ((v4 + ("foo bar" + v4)) + (("a" + "\n\t") + v4))), F1: 974, F2: (v5 + 9223372036854775807)}
    fmt.Println(v6)
    if false {
        v4 = v4
        v92 = (^(v92))
    } else {
        if true {
            v94 := GenS7{F0: -940, F1: 42}
            if (4294967295 == 32767) {
                if (float64((float32(1.5) / float32((1.5 / 1e10)))) > 0.1) {
                    if (!true) {
                        var v95 uint32 = v92
                        _ = v95
                    } else {
                        v4 = v4
                    }
                    var v96 int = 1
                    var v97 int8 = 0
                    _ = v96
                    _ = v97
                } else {
                    for fi98 := 0; fi98 < 19; fi98++ {
                        if (!false) {
                            if len(v3) > 0 {
                                v3[0] = fi98
                            }
                            fmt.Println(fi98, v4, v92)
                        }
                        v1 += 4294967295
                        v99 := GenS12{F0: (fi98 & 256)}
                        _ = fi98
                        _ = v99
                    }
                    if true {
                        for fi100 := 0; fi100 < 13; fi100++ {
                            var v101 uint = v1
                            var v102 bool = false
                            v5 += 42
                            _ = fi100
                            _ = v101
                            _ = v102
                        }
                        v91 += (v91 & (v1 / ((v1) | 1)))
                        _ = len(v3)
                    }
                }
            }
            _ = v94
        } else {
            v6 = v6
            fmt.Println(v5)
            var v103 uint8 = 0
            _ = v103
        }
        if len(v3) > 0 {
            v3[0] = 451
        }
        v5 -= genFn0(42, 67)
    }
    if true {
        var v104 uint32 = v92
        _ = v104
    }
    v5 -= genFn0(0, 255)
    if false {
        for fi105 := 0; fi105 < 1; fi105++ {
            var v106 int32 = -251
            v4 = v4
            _ = fi105
            _ = v106
        }
        v5 = -9223372036854775808
        if true {
            for fi107 := 0; fi107 < 15; fi107++ {
                v6 = v6
                var v108 int = (^(fi107))
                var v109 int32 = 1
                _ = fi107
                _ = v108
                _ = v109
            }
            for fi110 := 0; fi110 < 20; fi110++ {
                var v111 int32 = -849
                v112 := GenS7{F0: (v111 - 1), F1: 2}
                v3 = append(v3, fi110)
                _ = fi110
                _ = v111
                _ = v112
            }
        }
    } else {
        v113 := GenS1{F0: 1}
        _ = v113
    }
    switch ((v4 + ((v4 + "") + "")) + v4) {
    case "":
        v4 = v4
    case "\n\t":
        for fi114 := 0; fi114 < 8; fi114++ {
            v2 += v92
            v6 -= 0.1
            _ = fi114
        }
    case "a":
        var v115 uint32 = 2147483647
        _ = v115
    }
    fmt.Println(v2, v6)
    var v116 float32 = 3.14159
    for fi117 := 0; fi117 < 15; fi117++ {
        if ((false && (!(!true))) || (32767 < fi117)) {
            if (!false) {
                for fi118 := 0; fi118 < 20; fi118++ {
                    if false {
                        v92 += v2
                        var v119 uint = 0
                        _ = v119
                    } else {
                        v91 -= 21
                        v6 = v116
                    }
                    fmt.Println(fi117, v1)
                    _ = fi118
                }
                var v120 uint32 = 577
                _ = v120
            } else {
                for fi121 := 0; fi121 < 10; fi121++ {
                    if true {
                        var v122 float64 = 0.1
                        var v123 float32 = (v6 + (v116 - ((v116 / (-2.25 * (v6 / 0.1))) / (((float32(v122) - float32(v122)) / (v6 - (v6 - (((1.5 / -2.25) + v116) / ((v6 - (v116 + 0.1)) + v116))))) - v116))))
                        var v124 string = ("foo bar" + (v4 + "hello"))
                        v122 -= v122
                        fmt.Println(fi117, v116, fi117)
                        if (!(false && ((v5 >> ((583) & 63)) == (^(v5))))) {
                            fmt.Println(fi117, fi121)
                        } else {
                            fmt.Println(v91, v116)
                        }
                        v125 := GenS9{F0: (v4 + ""), F1: 42, F2: v5}
                        if (255 < 1) {
                            fmt.Println(v124)
                            if false {
                                var v126 uint = (v1 + 255)
                                _ = v126
                            } else {
                                switch fi117 {
                                case -643:
                                    v2 -= (^(v92))
                                case 127:
                                    v116 = v123
                                case 2147483647:
                                    v127 := GenS3{F0: -128, F1: -226}
                                    _ = v127
                                default:
                                    v122 += float64((((3.14159 - (v6 * (0.0 * float32((v122 * v122))))) - float32((v122 / 0.1))) * v116))
                                    if false {
                                        v5 = (v5 + (^(v5)))
                                        v128 := GenS9{F0: "\n\t", F1: (fi117 + (fi121 | fi121)), F2: (v5 & v5)}
                                        _ = v128
                                    }
                                    if false {
                                        var v129 uint64 = 812
                                        var v130 int64 = v5
                                        if true {
                                            fmt.Println(v4, v122)
                                            v131 := GenS5{F0: genFn0(1, 128), F1: -98, F2: 790}
                                            _ = v131
                                        }
                                        fmt.Println(v2)
                                        fmt.Println(v116, fi117, v124)
                                        _ = v129
                                        _ = v130
                                    }
                                }
                                fmt.Println(v91, v5)
                                switch (v124 + ("foo bar" + v124)) {
                                case "\n\t":
                                    v3 = append(v3, (fi117 + (-(fi117))))
                                }
                            }
                        } else {
                            v122 += float64((v6 - -2.25))
                            switch fi121 {
                            case -1:
                                var v132 uint = v1
                                _ = v132
                            case -34:
                                if true {
                                    if (true || (genFn1(1) != -817)) {
                                        v6 = (((1e10 / 1e10) + v123) / 1.0)
                                        if false {
                                            fmt.Println(v116)
                                            v133 := GenS11{F0: 98}
                                            v134 := GenS13{F0: fi117, F1: 42}
                                            _ = v133
                                            _ = v134
                                        }
                                    }
                                } else {
                                    v116 -= v123
                                }
                            case -2147483648:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            fmt.Println(v1, v1)
                        }
                        var v135 uint = v1
                        var v136 float64 = float64(v123)
                        switch fi117 {
                        case 9223372036854775807:
                            var v137 int64 = v5
                            _ = v137
                        case 32767:
                            switch v4 {
                            case "a":
                                switch v124 {
                                case "":
                                    var v138 uint = uint(v92)
                                    fmt.Println(v5, v124, v138)
                                    var v139 int32 = 2147483647
                                    _ = v138
                                    _ = v139
                                case "hello":
                                    switch 460 {
                                    case 127:
                                        fmt.Println(v123)
                                    case 9223372036854775807:
                                        v116 = ((v123 - v123) / ((((3.14159 / 1e10) / ((0.0 + v116) + ((1.5 + (((0.1 * float32(1e10)) + -2.25) * v6)) + float32(1.5)))) - 0.0) + (1.5 * (v123 / v116))))
                                    case -9223372036854775808:
                                        switch v124 {
                                        case "\n\t":
                                            var v140 int16 = 42
                                            _ = v140
                                        default:
                                            _ = len(v3)
                                        }
                                    default:
                                        var v141 uint = (v91 + (v135 | 4294967295))
                                        _ = v141
                                    }
                                }
                            case "":
                                fmt.Println(fi121)
                            default:
                                var v142 float64 = (0.1 - 0.1)
                                _ = v142
                            }
                        case 956:
                            switch (fi117 >> ((v135) & 63)) {
                            case 15:
                                var v143 int = -735
                                _ = v143
                            }
                        default:
                            if true {
                                if true {
                                    fmt.Println(v6)
                                }
                                var v144 bool = (genFn0(0, 0) > v5)
                                _ = len(v3)
                                _ = v144
                            } else {
                                if len(v3) > 0 {
                                    v3[0] = (fi121 % ((fi121) | 1))
                                }
                            }
                        }
                        fmt.Println(v123, v6)
                        v1 = 42
                        _ = v122
                        _ = v123
                        _ = v124
                        _ = v125
                        _ = v135
                        _ = v136
                    } else {
                        if (!true) {
                            _ = len(v3)
                            if ((^(v91)) >= v1) {
                                var v145 uint = 571
                                var v146 int = fi117
                                var v147 uint64 = 42
                                _ = v145
                                _ = v146
                                _ = v147
                            }
                            var v148 uint16 = uint16(v1)
                            _ = v148
                        } else {
                            if false {
                                switch int(v5) {
                                case 4294967295:
                                    fmt.Println(v4, v2)
                                case 1:
                                    v3 = append(v3, -2147483648)
                                default:
                                    v149 := GenS14{F0: 0}
                                    _ = v149
                                }
                                switch v4 {
                                case "hello":
                                    var v150 uint8 = 255
                                    _ = v150
                                case "":
                                    var v151 uint16 = 255
                                    _ = v151
                                }
                                var v152 uint8 = 42
                                _ = v152
                            } else {
                                v1 -= v1
                                v3 = append(v3, fi117)
                            }
                            if true {
                                v153 := GenS15{F0: -2147483648, F1: -421}
                                v154 := GenS5{F0: (-(v5)), F1: -1, F2: v1}
                                switch "foo bar" {
                                case "hello":
                                    var v155 int32 = -2147483648
                                    switch -407 {
                                    case 9223372036854775807:
                                        var v156 string = (v4 + (((("" + v4) + "foo bar") + ((v4 + v4) + v4)) + ("a" + v4)))
                                        _ = v156
                                    }
                                    _ = v155
                                case "a":
                                    v157 := GenS16{F0: 2, F1: false, F2: v92}
                                    _ = v157
                                case "":
                                    v116 = 1.5
                                }
                                if false {
                                    _ = len(v3)
                                } else {
                                    v1 = (v91 << (((v1 | (v1 * 475))) & 63))
                                    v158 := GenS4{F0: 45, F1: 122}
                                    _ = v158
                                }
                                _ = v153
                                _ = v154
                            } else {
                                v159 := GenS14{F0: 1}
                                fmt.Println(v1, fi117)
                                _ = v159
                            }
                        }
                    }
                    var v160 int = fi117
                    if len(v3) > 0 {
                        v3[0] = (v160 << ((v1) & 63))
                    }
                    _ = fi121
                    _ = v160
                }
            }
        } else {
            for fi161 := 0; fi161 < 13; fi161++ {
                switch v4 {
                case "":
                    var v162 uint32 = v2
                    v163 := GenS4{F0: 1, F1: 2}
                    _ = v162
                    _ = v163
                case "hello":
                    if len(v3) > 0 {
                        v3[0] = (fi161 - (fi161 << ((18446744073709551615) & 63)))
                    }
                case "a":
                    var v164 bool = true
                    _ = v164
                }
                _ = fi161
            }
            var v165 uint64 = 593
            _ = v165
        }
        _ = fi117
    }
    v92 = 65535
    var v166 uint8 = 167
    fmt.Println("v1:", v1)
    fmt.Println("v116:", v116)
    fmt.Println("v166:", v166)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v116
    _ = v166
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v90
    _ = v91
    _ = v92
    _ = v93
}
