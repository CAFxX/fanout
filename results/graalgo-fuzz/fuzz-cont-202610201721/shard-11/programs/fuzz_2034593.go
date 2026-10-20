package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 int; F2 string }
type GenS5 struct { F0 string; F1 uint16; F2 uint }
type GenS6 struct { F0 uint8; F1 uint32; F2 bool }
type GenS7 struct { F0 int16; F1 int32; F2 int8 }
type GenS8 struct { F0 int16; F1 uint8; F2 int32 }
type GenS9 struct { F0 string; F1 uint32 }
type GenS10 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS11 struct { F0 uint; F1 uint16 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 int; F1 bool }
type GenS16 struct { F0 int64; F1 int }
type GenS17 struct { F0 bool; F1 uint64 }
type GenS18 struct { F0 uint16 }
func genFn0() uint64 {
    var v7 uint8 = 2
    if false {
        fmt.Println(v7)
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 int = -1
            for fi10 := 0; fi10 < 5; fi10++ {
                var v11 int8 = -86
                _ = fi10
                _ = v11
            }
            var v12 uint16 = 42
            _ = fi8
            _ = v12
            _ = v9
        }
        v7 -= 0
    } else {
        v13 := GenS1{F0: 0}
        var v14 uint64 = 103
        for fi15 := 0; fi15 < 5; fi15++ {
            for fi16 := 0; fi16 < 15; fi16++ {
                v14 = v14
                _ = fi16
            }
            v17 := []int{4294967295, fi15, (fi15 % (((^(fi15))) | 1))}
            _ = fi15
            _ = v17
        }
        _ = v13
        _ = v14
    }
    fmt.Println(v7)
    fmt.Println(v7)
    v18 := GenS1{F0: 9223372036854775807}
    _ = v18
    _ = v7
    return 42
}
func genFn1(p0 uint8, p1 int32) int8 {
    if true {
        switch "" {
        case "hello":
            switch ("foo bar" + "\n\t") {
            case "hello":
                p1 = p1
            case "a":
                for fi19 := 0; fi19 < 12; fi19++ {
                    var v20 uint32 = 2147483647
                    v21 := []int{(fi19 + fi19), fi19}
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            case "":
                p1 = p1
            default:
                v22 := GenS2{F0: 1}
                _ = v22
            }
        case "\n\t":
            p0 = p0
        case "":
            fmt.Println(p0, p1)
        default:
            v23 := GenS3{F0: p1}
            _ = v23
        }
        for fi24 := 0; fi24 < 20; fi24++ {
            if false {
                switch (fi24 & 9223372036854775807) {
                case 1:
                    if false {
                        v25 := GenS4{F0: p0, F1: (fi24 >> ((255) & 63)), F2: "foo bar"}
                        if false {
                            var v26 uint32 = 1
                            if (0 <= 100) {
                                var v27 string = ("\n\t" + "\n\t")
                                var v28 int = int(v26)
                                v29 := []int{}
                                _ = v27
                                _ = v28
                                _ = v29
                            }
                            _ = v26
                        } else {
                            for fi30 := 0; fi30 < 4; fi30++ {
                                if false {
                                    var v31 uint32 = 4294967295
                                    switch ("" + "foo bar") {
                                    case "\n\t":
                                        var v32 int8 = 127
                                        _ = v32
                                    case "hello":
                                        v33 := []int{-2147483648, (fi30 & 127), (^(fi30)), 42}
                                        _ = v33
                                    case "":
                                        v34 := GenS5{F0: "", F1: 65535, F2: 465}
                                        _ = v34
                                    default:
                                        p1 = p1
                                    }
                                    if true {
                                        v35 := GenS4{F0: (p0 >> ((uint(p0)) & 63)), F1: -259, F2: ("hello" + "")}
                                        _ = v35
                                    } else {
                                        var v36 uint32 = v31
                                        v37 := []int{-1, (^(fi30))}
                                        _ = v36
                                        _ = v37
                                    }
                                    _ = v31
                                }
                                _ = fi30
                            }
                            switch ("\n\t" + ("\n\t" + ("a" + "\n\t"))) {
                            case "":
                                fmt.Println(p0)
                            case "foo bar":
                                var v38 float64 = 1e10
                                _ = v38
                            case "hello":
                                fmt.Println(p0, fi24)
                            default:
                                for fi39 := 0; fi39 < 2; fi39++ {
                                    v40 := GenS2{F0: 961}
                                    _ = fi39
                                    _ = v40
                                }
                            }
                            if ((((0.1 + 1.5) == float32((((0.0 / 1e10) / float64(1.5)) - 0.1))) && false) || (502 >= 255)) {
                                if ((-1 < 42) && (!true)) {
                                    fmt.Println(p1, p0)
                                    var v41 int32 = (p1 & -2147483648)
                                    p1 = 1
                                    _ = v41
                                } else {
                                    if false {
                                        p0 -= (p0 - p0)
                                        p1 -= -1
                                    } else {
                                        if (279 != 4294967295) {
                                            fmt.Println(p1, fi24, fi24)
                                            var v42 int32 = 1
                                            _ = v42
                                        }
                                        for fi43 := 0; fi43 < 5; fi43++ {
                                            switch ("hello" + "\n\t") {
                                            case "":
                                                v44 := GenS1{F0: 9223372036854775807}
                                                _ = v44
                                            case "\n\t":
                                                var v45 uint64 = 18446744073709551615
                                                _ = v45
                                            case "foo bar":
                                                var v46 int16 = 42
                                                _ = v46
                                            }
                                            _ = fi43
                                        }
                                        for fi47 := 0; fi47 < 17; fi47++ {
                                            if (!(!true)) {
                                                p0 = (p0 << ((338) & 63))
                                            }
                                            switch "hello" {
                                            case "hello":
                                                p1 -= 2147483647
                                            default:
                                                var v48 int8 = 127
                                                _ = v48
                                            }
                                            v49 := []int{fi24}
                                            _ = fi47
                                            _ = v49
                                        }
                                    }
                                    v50 := []int{(fi24 << ((4294967295) & 63)), fi24}
                                    fmt.Println(p0, p1, p0)
                                    _ = v50
                                }
                                fmt.Println(fi24)
                                switch (("" + ("foo bar" + "hello")) + ("\n\t" + (("foo bar" + "") + ("foo bar" + ((("a" + ((("" + "foo bar") + "\n\t") + "")) + "foo bar") + ("a" + "foo bar")))))) {
                                case "":
                                    p1 -= (p1 | p1)
                                }
                            } else {
                                for fi51 := 0; fi51 < 11; fi51++ {
                                    if (((-2.25 + 0.0) - (3.14159 / float64(3.14159))) <= 3.14159) {
                                        fmt.Println(fi24, p1, p0)
                                    }
                                    var v52 uint64 = 347
                                    v53 := GenS4{F0: p0, F1: fi51, F2: "foo bar"}
                                    v52 += (v52 * (^(v52)))
                                    switch "" {
                                    case "foo bar":
                                        if ((p1 != p1) && true) {
                                            if (((!(fi51 >= (fi51 - 0))) || (65535 > 105)) && (true && ((true || (1 >= -569)) && true))) {
                                                p0 = (p0 + p0)
                                                fmt.Println(fi51)
                                                v54 := GenS3{F0: (^(p1))}
                                                var v55 bool = (1e10 < 3.14159)
                                                var v56 uint32 = 4294967295
                                                _ = v54
                                                _ = v55
                                                _ = v56
                                            }
                                        }
                                    case "":
                                        switch "\n\t" {
                                        case "foo bar":
                                            v57 := []int{fi51, fi24, fi51, (fi51 / ((-474) | 1))}
                                            _ = v57
                                        }
                                    case "\n\t":
                                        fmt.Println(v52, v52)
                                    }
                                    v58 := []int{}
                                    var v59 uint = 2
                                    _ = fi51
                                    _ = v52
                                    _ = v53
                                    _ = v58
                                    _ = v59
                                }
                                p0 = 2
                                v60 := GenS5{F0: "", F1: 42, F2: 1}
                                _ = v60
                            }
                        }
                        _ = v25
                    }
                case 0:
                    switch 127 {
                    case 728:
                        v61 := GenS1{F0: 0}
                        _ = v61
                    }
                default:
                    if false {
                        fmt.Println(p0, p0, fi24)
                    } else {
                        if true {
                            switch (fi24 | fi24) {
                            case 128:
                                fmt.Println(p1, p1, p0)
                            case -476:
                                p0 += p0
                            case -758:
                                switch (fi24 >> ((18446744073709551615) & 63)) {
                                case 32767:
                                    v62 := []int{}
                                    _ = v62
                                default:
                                    var v63 float32 = 1.5
                                    _ = v63
                                }
                            }
                            switch "foo bar" {
                            case "":
                                switch ((("hello" + "a") + ((("hello" + "a") + "") + ("" + "foo bar"))) + ((("a" + ("hello" + (("" + "a") + "a"))) + ("a" + "\n\t")) + "\n\t")) {
                                case "":
                                    switch (fi24 ^ (fi24 ^ (^(fi24)))) {
                                    case 871:
                                        if true {
                                            fmt.Println(p0)
                                        } else {
                                            var v64 float32 = 1.5
                                            var v65 float64 = 0.0
                                            v64 += 1.5
                                            _ = v64
                                            _ = v65
                                        }
                                    case 1:
                                        p1 = (p1 >> ((4294967295) & 63))
                                    case 42:
                                        fmt.Println(p0)
                                    default:
                                        switch (fi24 >> ((2) & 63)) {
                                        case 909:
                                            var v66 uint8 = (p0 & 215)
                                            _ = v66
                                        case -809:
                                            fmt.Println(fi24, p0, p0)
                                        default:
                                            v67 := GenS2{F0: 101}
                                            _ = v67
                                        }
                                    }
                                case "a":
                                    v68 := []int{(fi24 << ((2) & 63))}
                                    _ = v68
                                case "\n\t":
                                    if ((!true) || false) {
                                        if (!(false && (!(!true)))) {
                                            fmt.Println(p1, fi24)
                                            v69 := GenS6{F0: p0, F1: 540, F2: false}
                                            _ = v69
                                        }
                                        p1 -= p1
                                        for fi70 := 0; fi70 < 6; fi70++ {
                                            v71 := []int{(fi70 + (^(fi24))), (fi24 - fi24), fi70, (fi24 - fi24)}
                                            _ = fi70
                                            _ = v71
                                        }
                                    } else {
                                        switch fi24 {
                                        case 42:
                                            if (!false) {
                                                p1 += p1
                                                var v72 uint8 = p0
                                                for fi73 := 0; fi73 < 11; fi73++ {
                                                    p1 = -202
                                                    v74 := GenS4{F0: v72, F1: fi73, F2: (((("\n\t" + ("\n\t" + "a")) + "foo bar") + "hello") + "\n\t")}
                                                    var v75 bool = ((1.5 * 0.0) != -2.25)
                                                    _ = fi73
                                                    _ = v74
                                                    _ = v75
                                                }
                                                _ = v72
                                            }
                                        case -1:
                                            var v76 float32 = (-2.25 - 0.1)
                                            _ = v76
                                        }
                                        var v77 float64 = ((3.14159 + (0.1 * (3.14159 - 3.14159))) / 3.14159)
                                        for fi78 := 0; fi78 < 3; fi78++ {
                                            var v79 uint8 = (p0 - (p0 % ((p0) | 1)))
                                            if false {
                                                v79 = p0
                                                p1 -= p1
                                            }
                                            fmt.Println(fi24)
                                            _ = fi78
                                            _ = v79
                                        }
                                        _ = v77
                                    }
                                }
                            case "\n\t":
                                for fi80 := 0; fi80 < 15; fi80++ {
                                    switch (-(fi80)) {
                                    case -122:
                                        v81 := GenS1{F0: 42}
                                        _ = v81
                                    }
                                    p0 = 2
                                    if (("hello" > "\n\t") || false) {
                                        switch 836 {
                                        case 32767:
                                            p0 = p0
                                        case 32768:
                                            var v82 int16 = 360
                                            _ = v82
                                        case 128:
                                            p1 = (p1 | p1)
                                        }
                                        v83 := []int{fi80, 127, fi80}
                                        _ = v83
                                    }
                                    _ = fi80
                                }
                            default:
                                switch "hello" {
                                case "":
                                    var v84 uint32 = 830
                                    _ = v84
                                case "a":
                                    if (42 >= 157) {
                                        var v85 int64 = 9223372036854775807
                                        for fi86 := 0; fi86 < 5; fi86++ {
                                            var v87 uint8 = 42
                                            _ = fi86
                                            _ = v87
                                        }
                                        switch (^(fi24)) {
                                        case 9223372036854775807:
                                            var v88 float32 = 0.0
                                            _ = v88
                                        case 32767:
                                            v89 := GenS1{F0: (v85 ^ 0)}
                                            _ = v89
                                        }
                                        _ = v85
                                    }
                                case "hello":
                                    switch "" {
                                    case "":
                                        var v90 int16 = -1
                                        _ = v90
                                    case "foo bar":
                                        if (!((!(false || (3.14159 == (1e10 * float32(0.0))))) && (!(true || (0 > 4294967295))))) {
                                            var v91 int32 = p1
                                            fmt.Println(p1, p1, fi24)
                                            _ = v91
                                        } else {
                                            var v92 uint64 = 4294967295
                                            if (0 <= 42) {
                                                if false {
                                                    if false {
                                                        var v93 float64 = ((1.5 / 0.1) / 1.5)
                                                        var v94 int64 = 9223372036854775807
                                                        _ = v93
                                                        _ = v94
                                                    } else {
                                                        fmt.Println(v92, p0)
                                                        for fi95 := 0; fi95 < 8; fi95++ {
                                                            var v96 int8 = 51
                                                            _ = fi95
                                                            _ = v96
                                                        }
                                                    }
                                                    var v97 uint8 = p0
                                                    _ = v97
                                                }
                                            } else {
                                                v92 = (v92 / ((18446744073709551615) | 1))
                                                var v98 int16 = int16(fi24)
                                                _ = v98
                                            }
                                            _ = v92
                                        }
                                    case "hello":
                                        fmt.Println(p0, p0, fi24)
                                    }
                                default:
                                    var v99 uint = 33
                                    _ = v99
                                }
                            }
                            v100 := GenS3{F0: p1}
                            _ = v100
                        }
                        var v101 int8 = -114
                        v102 := []int{}
                        _ = v101
                        _ = v102
                    }
                }
                var v103 float64 = 1.5
                _ = v103
            } else {
                v104 := []int{fi24}
                p0 = (p0 ^ 128)
                for fi105 := 0; fi105 < 19; fi105++ {
                    v106 := GenS7{F0: 0, F1: (-(p1)), F2: 1}
                    v104 = append(v104, (fi24 * 922))
                    v107 := GenS5{F0: "hello", F1: 255, F2: 1}
                    _ = fi105
                    _ = v106
                    _ = v107
                }
                _ = v104
            }
            p1 += 257
            for fi108 := 0; fi108 < 10; fi108++ {
                if true {
                    v109 := GenS6{F0: p0, F1: 818, F2: ((p1 - 1) >= p1)}
                    _ = v109
                } else {
                    var v110 float64 = 1.5
                    var v111 int16 = -32768
                    if (!(-677 > 42)) {
                        var v112 string = "a"
                        _ = v112
                    } else {
                        fmt.Println(fi108, p1, p0)
                        var v113 int32 = 0
                        _ = v113
                    }
                    _ = v110
                    _ = v111
                }
                v114 := []int{42, fi24, 2147483647}
                _ = fi108
                _ = v114
            }
            _ = fi24
        }
        var v115 float32 = 0.0
        _ = v115
    } else {
        var v116 uint = 255
        _ = v116
    }
    v117 := GenS7{F0: int16(p0), F1: int32(p0), F2: -52}
    p0 += 1
    p0 -= p0
    var v118 bool = (!true)
    _ = p0
    _ = p1
    _ = v117
    _ = v118
    return 100
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"foo bar"}
    var v3 int16 = -1
    var v4 float32 = 0.1
    v5 := make(map[int]int)
    var v6 string = ""
    for fi119 := 0; fi119 < 8; fi119++ {
        for fi120 := 0; fi120 < 7; fi120++ {
            v4 -= 1e10
            _ = fi120
        }
        _ = fi119
    }
    if (248 < 42) {
        var v121 int16 = (v3 & v3)
        if (false || ((!false) && (!false))) {
            var v122 int64 = 272
            if (360 >= 18446744073709551615) {
                v123 := GenS8{F0: (v121 - v3), F1: 255, F2: -2147483648}
                _ = v123
            }
            var v124 string = v6
            _ = v122
            _ = v124
        } else {
            var v125 uint32 = 65535
            delete(v1, 128)
            _ = v125
        }
        _ = v121
    } else {
        v6 = (((v6 + v6) + (v6 + "")) + v6)
        delete(v1, 952)
        v4 = 0.0
    }
    v4 = v4
    var v126 float32 = 0.1
    var v127 uint32 = 331
    v128 := GenS1{F0: 604}
    var v129 bool = (65535 == 42)
    v126 = v126
    v6 = "\n\t"
    for fi130 := 0; fi130 < 10; fi130++ {
        for fi131 := 0; fi131 < 14; fi131++ {
            v129 = v129
            switch (fi130 >> ((1) & 63)) {
            case 128:
                v132 := GenS2{F0: 9223372036854775807}
                _ = v132
            case 127:
                v5[-1] = 32768
            case 141:
                v129 = (!false)
            default:
                v133 := GenS9{F0: ("a" + "a"), F1: v127}
                _ = v133
            }
            _ = fi131
        }
        _ = fi130
    }
    v134 := GenS5{F0: (v6 + "\n\t"), F1: 65535, F2: 18446744073709551615}
    v3 -= v3
    v126 = v4
    if v129 {
        if (42 <= 128) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v135 int = 9223372036854775807
            fmt.Println(v126, v126, v127)
            _ = v135
        } else {
            for fi136 := 0; fi136 < 15; fi136++ {
                switch (fi136 + (fi136 * (fi136 * (fi136 - fi136)))) {
                case 2147483647:
                    v137 := GenS10{F0: 4294967295, F1: 4294967295, F2: 2147483647}
                    _ = v137
                default:
                    if (128 >= 1) {
                        if v129 {
                            var v138 uint16 = 32767
                            var v139 int16 = int16(v127)
                            var v140 string = v6
                            _ = v138
                            _ = v139
                            _ = v140
                        }
                        v2 = append(v2, "")
                        v141 := GenS7{F0: -1, F1: -1, F2: 95}
                        _ = v141
                    } else {
                        var v142 int = 239
                        for fi143 := 0; fi143 < 5; fi143++ {
                            switch fi136 {
                            case 128:
                                var v144 bool = v129
                                _ = v144
                            case -9223372036854775808:
                                var v145 int = (fi136 << ((42) & 63))
                                _ = v145
                            case -191:
                                var v146 uint8 = 172
                                if (v129 || v129) {
                                    v147 := GenS11{F0: 1, F1: 642}
                                    _ = v147
                                } else {
                                    switch ((v6 + v6) + "\n\t") {
                                    case "\n\t":
                                        switch "a" {
                                        case "\n\t":
                                            v148 := GenS2{F0: 0}
                                            _ = v148
                                        default:
                                            fmt.Println(fi143, v129)
                                            if (v129 || (v127 != (^(v127)))) {
                                                var v149 float32 = (float32(1e10) / (v126 - v126))
                                                v1[-94] = v129
                                                var v150 uint32 = (^(v127))
                                                _ = v149
                                                _ = v150
                                            } else {
                                                v3 -= v3
                                                v2 = append(v2, (v6 + "a"))
                                                fmt.Println(v142, v127, v126)
                                            }
                                        }
                                    case "hello":
                                        fmt.Println(v129, v3, v142)
                                    case "foo bar":
                                        v127 -= (v127 + (v127 + 2147483647))
                                    }
                                    _ = len(v1)
                                    var v151 uint = 0
                                    var v152 float32 = (((v4 + 1.5) * 0.1) + ((1e10 / float32(-2.25)) - (float32((1.5 + (0.0 - 0.0))) * v4)))
                                    v142 = fi143
                                    var v153 float64 = 0.1
                                    if len(v2) > 0 {
                                        v2[0] = "a"
                                    }
                                    switch v6 {
                                    case "":
                                        v2 = append(v2, v6)
                                    case "foo bar":
                                        var v154 string = v6
                                        _ = v154
                                    default:
                                        v153 -= float64(0.0)
                                    }
                                    v155 := GenS6{F0: 255, F1: v127, F2: (!v129)}
                                    _ = v151
                                    _ = v152
                                    _ = v153
                                    _ = v155
                                }
                                v4 -= v4
                                _ = v146
                            default:
                                v156 := GenS2{F0: 9223372036854775807}
                                _ = v156
                            }
                            v126 = float32((0.0 * (-2.25 * 1e10)))
                            var v157 int32 = 650
                            _ = fi143
                            _ = v157
                        }
                        _ = v142
                    }
                }
                _ = len(v5)
                _ = fi136
            }
            fmt.Println(v127, v129)
        }
        var v158 string = "hello"
        for fi159 := 0; fi159 < 5; fi159++ {
            switch (fi159 | fi159) {
            case -362:
                var v160 uint16 = 32767
                _ = v160
            default:
                var v161 uint32 = v127
                _ = v161
            }
            _ = len(v2)
            fmt.Println(v3, v127)
            _ = fi159
        }
        _ = v158
    }
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    switch (("" + "hello") + v6) {
    case "\n\t":
        v162 := GenS12{F0: -2147483648}
        _ = v162
    case "hello":
        v163 := GenS13{F0: v3}
        _ = v163
    case "":
        for fi164 := 0; fi164 < 7; fi164++ {
            if v129 {
                var v165 int16 = 1000
                v166 := GenS14{F0: v165}
                var v167 float32 = 0.0
                _ = v165
                _ = v166
                _ = v167
            } else {
                if v129 {
                    switch (fi164 % ((fi164) | 1)) {
                    case 235:
                        for fi168 := 0; fi168 < 20; fi168++ {
                            if v129 {
                                var v169 uint64 = 1
                                fmt.Println(v3)
                                _ = v169
                            } else {
                                v170 := GenS10{F0: genFn0(), F1: 18446744073709551615, F2: 0}
                                var v171 float32 = (v126 + (v126 + -2.25))
                                v127 = 2
                                var v172 float64 = (-2.25 - (3.14159 / float64((((float32(1e10) - (v171 - (3.14159 + v126))) * (((((0.1 / v126) + v126) + 1.5) - float32((-2.25 - 3.14159))) * float32(0.1))) - float32(-2.25)))))
                                if (!v129) {
                                    v172 = (0.1 / v172)
                                }
                                switch v6 {
                                case "foo bar":
                                    v173 := GenS15{F0: -277, F1: v129}
                                    if true {
                                        switch v6 {
                                        case "hello":
                                            var v174 uint16 = 32767
                                            _ = v174
                                        default:
                                            if (1 > 127) {
                                                var v175 bool = v129
                                                if false {
                                                    v176 := GenS16{F0: 42, F1: (fi168 * (fi168 - fi164))}
                                                    v127 -= v127
                                                    _ = v176
                                                } else {
                                                    v177 := GenS17{F0: v129, F1: 724}
                                                    v178 := GenS13{F0: 42}
                                                    var v179 uint32 = (v127 + (^(v127)))
                                                    _ = v177
                                                    _ = v178
                                                    _ = v179
                                                }
                                                _ = v175
                                            }
                                        }
                                        v172 -= (float64(v126) / float64(1e10))
                                        v126 = v171
                                    } else {
                                        var v180 uint16 = 0
                                        var v181 uint64 = 42
                                        v171 = (3.14159 + 0.1)
                                        _ = v180
                                        _ = v181
                                    }
                                    _ = v173
                                case "\n\t":
                                    var v182 int64 = -129
                                    _ = v182
                                }
                                _ = v170
                                _ = v171
                                _ = v172
                            }
                            fmt.Println(v127)
                            v129 = (v6 == (v6 + (v6 + ("\n\t" + "a"))))
                            _ = fi168
                        }
                    default:
                        v127 = v127
                    }
                } else {
                    fmt.Println(v3)
                    _ = len(v2)
                    for fi183 := 0; fi183 < 13; fi183++ {
                        fmt.Println(fi164)
                        _ = fi183
                    }
                }
            }
            v4 += ((((0.0 + 1.5) + v126) / v4) * v126)
            v184 := GenS18{F0: 1}
            _ = fi164
            _ = v184
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v126:", v126)
    fmt.Println("v127:", v127)
    fmt.Println("v129:", v129)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v126
    _ = v127
    _ = v128
    _ = v129
    _ = v134
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
