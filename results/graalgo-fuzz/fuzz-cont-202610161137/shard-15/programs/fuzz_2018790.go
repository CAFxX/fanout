package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS8 struct { F0 uint64; F1 uint64; F2 int64 }
type GenS9 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int8; F1 int16; F2 int32 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 int8; F1 uint16 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int8; F1 uint }
type GenS16 struct { F0 uint32 }
type GenS17 struct { F0 int32; F1 int; F2 int16 }
func genFn0(p0 int32) uint64 {
    switch "\n\t" {
    case "hello":
        for fi4 := 0; fi4 < 10; fi4++ {
            v5 := []int{32767, -383}
            for fi6 := 0; fi6 < 5; fi6++ {
                var v7 uint64 = 0
                var v8 float64 = float64(0.0)
                _ = fi6
                _ = v7
                _ = v8
            }
            _ = fi4
            _ = v5
        }
    case "\n\t":
        var v9 uint8 = 255
        _ = v9
    case "foo bar":
        for fi10 := 0; fi10 < 6; fi10++ {
            for fi11 := 0; fi11 < 8; fi11++ {
                p0 = p0
                if (0 > 32767) {
                    v12 := GenS1{F0: 69}
                    v13 := []int{4294967295, (fi10 << ((18446744073709551615) & 63))}
                    _ = v12
                    _ = v13
                } else {
                    fmt.Println(p0, fi11, fi10)
                    fmt.Println(fi11)
                    p0 = p0
                }
                if true {
                    var v14 int32 = -2147483648
                    v15 := []int{(fi10 | (fi11 << ((4294967295) & 63))), (fi10 | (fi11 >> ((0) & 63))), -9223372036854775808, fi11}
                    _ = v14
                    _ = v15
                } else {
                    var v16 int16 = 1
                    p0 += 2147483647
                    _ = v16
                }
                _ = fi11
            }
            v17 := GenS1{F0: 0}
            _ = fi10
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        for fi19 := 0; fi19 < 9; fi19++ {
            var v20 uint32 = 0
            v21 := GenS2{F0: 771, F1: 32767}
            switch "" {
            case "a":
                var v22 uint64 = 142
                _ = v22
            case "":
                p0 = p0
            default:
                fmt.Println(v20, fi18)
            }
            _ = fi19
            _ = v20
            _ = v21
        }
        v23 := GenS1{F0: 255}
        if true {
            for fi24 := 0; fi24 < 7; fi24++ {
                var v25 int32 = (p0 & p0)
                _ = fi24
                _ = v25
            }
        } else {
            p0 -= p0
            p0 -= (p0 ^ (p0 + -252))
        }
        _ = fi18
        _ = v23
    }
    var v26 int16 = 1
    p0 += p0
    _ = p0
    _ = v26
    return 1
}

func main() {
    var v1 int8 = 0
    v2 := []uint32{42, 891, 0, 65535}
    var v3 uint8 = 128
    if ((true && true) || true) {
        v3 += (v3 % (((v3 << ((2) & 63))) | 1))
    }
    switch "a" {
    case "a":
        if (((0 <= 244) || true) && false) {
            v3 -= 42
            switch 2147483647 {
            case -739:
                switch (("hello" + "hello") + "\n\t") {
                case "":
                    v1 = (v1 | v1)
                case "foo bar":
                    var v27 bool = true
                    _ = v27
                default:
                    if len(v2) > 0 {
                        v2[0] = 2
                    }
                }
            case 2147483647:
                if false {
                    v2 = append(v2, 2147483647)
                    if (v3 > 255) {
                        for fi28 := 0; fi28 < 17; fi28++ {
                            v1 = -1
                            if false {
                                fmt.Println(v3, fi28, fi28)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v3 = v3
                            }
                            var v29 uint8 = v3
                            _ = fi28
                            _ = v29
                        }
                        v3 = v3
                        v1 -= (^(v1))
                    }
                    var v30 uint64 = genFn0(-2147483648)
                    _ = v30
                }
            case 9223372036854775807:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        }
    case "foo bar":
        _ = len(v2)
    case "":
        var v31 string = ""
        _ = v31
    default:
        if true {
            if true {
                for fi32 := 0; fi32 < 5; fi32++ {
                    fmt.Println(fi32, fi32)
                    v33 := GenS2{F0: 0, F1: 32767}
                    for fi34 := 0; fi34 < 3; fi34++ {
                        v35 := GenS1{F0: v3}
                        v1 -= 100
                        _ = fi34
                        _ = v35
                    }
                    _ = fi32
                    _ = v33
                }
            } else {
                var v36 string = (("a" + "\n\t") + ("" + ("a" + "hello")))
                var v37 uint16 = 2
                var v38 bool = ((v1 ^ -27) >= 31)
                _ = v36
                _ = v37
                _ = v38
            }
            if true {
                var v39 int16 = 0
                _ = v39
            } else {
                v3 = v3
                v40 := GenS3{F0: 2147483647, F1: -32768}
                _ = v40
            }
            if len(v2) > 0 {
                v2[0] = 1
            }
        } else {
            var v41 uint = 255
            var v42 uint64 = 18446744073709551615
            v1 += (v1 | v1)
            _ = v41
            _ = v42
        }
    }
    if false {
        if false {
            if len(v2) > 0 {
                v2[0] = 0
            }
            v43 := GenS4{F0: ""}
            v3 = v3
            _ = v43
        }
        fmt.Println(v3)
    } else {
        if (!((2147483647 > 1) && (9223372036854775807 >= int64(v1)))) {
            var v44 int64 = -1
            switch -9223372036854775808 {
            case 4294967295:
                if (1 != 0) {
                    if false {
                        v3 = 128
                        if false {
                            v45 := GenS5{F0: true}
                            _ = v45
                        }
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            v3 += 214
                        }
                        v46 := GenS6{F0: "hello"}
                        v2 = append(v2, 0)
                        _ = v46
                    }
                }
            default:
                v47 := GenS6{F0: "a"}
                _ = v47
            }
            v44 = -9223372036854775808
            _ = v44
        } else {
            if true {
                fmt.Println(v1)
                v1 = (^(v1))
                v3 -= v3
            }
            if len(v2) > 0 {
                v2[0] = 1
            }
        }
        for fi48 := 0; fi48 < 17; fi48++ {
            for fi49 := 0; fi49 < 19; fi49++ {
                fmt.Println(fi48, v1)
                _ = fi49
            }
            v1 -= v1
            for fi50 := 0; fi50 < 8; fi50++ {
                var v51 uint32 = 2
                if true {
                    if (false && false) {
                        var v52 float64 = float64(float32(((3.14159 - -2.25) - 0.0)))
                        _ = v52
                    } else {
                        if (false || (32767 > 32767)) {
                            var v53 uint32 = v51
                            switch "" {
                            case "foo bar":
                                var v54 uint = 4294967295
                                _ = v54
                            case "\n\t":
                                var v55 int32 = 2147483647
                                _ = v55
                            default:
                                fmt.Println(v53)
                            }
                            v53 -= 65535
                            _ = v53
                        } else {
                            _ = len(v2)
                            if len(v2) > 0 {
                                v2[0] = 0
                            }
                        }
                    }
                    _ = len(v2)
                }
                _ = fi50
                _ = v51
            }
            _ = fi48
        }
        if (0 > 1) {
            switch 9223372036854775807 {
            case 42:
                for fi56 := 0; fi56 < 18; fi56++ {
                    v1 = v1
                    var v57 int64 = 1
                    v58 := GenS7{F0: (v1 | (v1 & v1)), F1: 332, F2: 32767}
                    _ = fi56
                    _ = v57
                    _ = v58
                }
            case -781:
                switch "hello" {
                case "":
                    v3 = (v3 & (v3 % ((223) | 1)))
                default:
                    switch 127 {
                    case 32768:
                        if false {
                            for fi59 := 0; fi59 < 2; fi59++ {
                                if (!(-2.25 == 1.5)) {
                                    var v60 string = ""
                                    _ = v60
                                } else {
                                    v3 = v3
                                }
                                fmt.Println(v3, v3)
                                _ = fi59
                            }
                            v3 -= v3
                            if true {
                                v3 += 112
                            } else {
                                var v61 uint64 = 18446744073709551615
                                _ = v61
                            }
                        }
                    default:
                        if (2 > 4294967295) {
                            fmt.Println(v3)
                            if len(v2) > 0 {
                                v2[0] = 4294967295
                            }
                            v2 = append(v2, 0)
                        }
                    }
                }
            case -9223372036854775808:
                v62 := GenS6{F0: ("" + "hello")}
                _ = v62
            default:
                fmt.Println(v3)
            }
        } else {
            for fi63 := 0; fi63 < 19; fi63++ {
                for fi64 := 0; fi64 < 7; fi64++ {
                    v3 -= v3
                    v65 := GenS3{F0: 4294967295, F1: -32768}
                    _ = fi64
                    _ = v65
                }
                var v66 int8 = v1
                for fi67 := 0; fi67 < 11; fi67++ {
                    v68 := GenS5{F0: false}
                    if (255 == 18446744073709551615) {
                        v69 := GenS6{F0: "\n\t"}
                        if (true && (496 < 65535)) {
                            var v70 uint8 = (v3 / ((v3) | 1))
                            v70 = (v3 + 128)
                            if (-976 >= 32767) {
                                var v71 uint32 = uint32(v3)
                                v1 += 42
                                _ = v71
                            } else {
                                var v72 uint32 = 4294967295
                                fmt.Println(v70, fi67, fi63)
                                var v73 uint8 = 116
                                _ = v72
                                _ = v73
                            }
                            var v74 string = ("" + "foo bar")
                            var v75 uint = 42
                            _ = v70
                            _ = v74
                            _ = v75
                        }
                        if (-1 <= -128) {
                            v1 -= (v66 * (v1 / ((100) | 1)))
                            var v76 int = fi63
                            _ = v76
                        }
                        _ = v69
                    } else {
                        if true {
                            if true {
                                var v77 int64 = -1
                                fmt.Println(fi63)
                                v78 := GenS1{F0: 1}
                                _ = v77
                                _ = v78
                            }
                            if len(v2) > 0 {
                                v2[0] = 2
                            }
                            if true {
                                fmt.Println(v1)
                                var v79 int16 = 1
                                var v80 uint = 65535
                                v81 := GenS8{F0: 4294967295, F1: 4294967295, F2: 1}
                                _ = v79
                                _ = v80
                                _ = v81
                            }
                            if ((true && (false && false)) && false) {
                                var v82 int64 = -1
                                switch (fi63 & fi67) {
                                case 1:
                                    v82 = v82
                                case -9223372036854775808:
                                    var v83 int8 = -54
                                    if len(v2) > 0 {
                                        v2[0] = 2147483648
                                    }
                                    switch ("" + "a") {
                                    case "foo bar":
                                        if false {
                                            v84 := GenS4{F0: ("hello" + "hello")}
                                            var v85 uint16 = 0
                                            var v86 uint32 = uint32(v3)
                                            if false {
                                                v87 := GenS9{F0: 0, F1: 42, F2: (v85 / ((1) | 1))}
                                                _ = v87
                                            }
                                            _ = v84
                                            _ = v85
                                            _ = v86
                                        } else {
                                            v83 = v83
                                            var v88 uint64 = 0
                                            fmt.Println(v83, v66)
                                            _ = v88
                                        }
                                    case "a":
                                        switch int(v66) {
                                        case -471:
                                            switch -2147483648 {
                                            case 0:
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            default:
                                                v66 = v1
                                            }
                                        case -764:
                                            _ = len(v2)
                                        default:
                                            var v89 uint8 = 255
                                            _ = v89
                                        }
                                    default:
                                        var v90 uint16 = 0
                                        _ = v90
                                    }
                                    v83 -= v83
                                    _ = v83
                                default:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                v66 -= (v66 & (v1 / ((42) | 1)))
                                _ = v82
                            }
                            if ((!(0.0 <= float32(3.14159))) && (1 <= 0)) {
                                _ = len(v2)
                                if (!true) {
                                    var v91 int64 = -1
                                    var v92 int16 = 42
                                    var v93 uint16 = 401
                                    _ = v91
                                    _ = v92
                                    _ = v93
                                } else {
                                    var v94 uint16 = 1
                                    if true {
                                        if false {
                                            var v95 int32 = -774
                                            _ = v95
                                        } else {
                                            var v96 int8 = v1
                                            var v97 int16 = 0
                                            _ = v96
                                            _ = v97
                                        }
                                        var v98 int32 = -2147483648
                                        _ = v98
                                    }
                                    v99 := GenS10{F0: -2147483648}
                                    _ = v94
                                    _ = v99
                                }
                                var v100 float64 = 1e10
                                _ = v100
                            }
                            v66 -= v66
                            if (!(true || true)) {
                                v3 += v3
                                if true {
                                    switch (fi63 % (((-(fi63))) | 1)) {
                                    case 184:
                                        var v101 float64 = (0.1 + (((3.14159 / -2.25) * -2.25) / 1.0))
                                        _ = v101
                                    case 127:
                                        var v102 uint32 = 42
                                        _ = v102
                                    case 128:
                                        v3 -= 1
                                    default:
                                        var v103 uint16 = 32767
                                        _ = v103
                                    }
                                } else {
                                    v1 -= v1
                                    var v104 int64 = 9223372036854775807
                                    switch (-(fi67)) {
                                    case 4294967295:
                                        switch 127 {
                                        case 426:
                                            v2 = append(v2, 0)
                                        case 2147483647:
                                            if (!(42 > -1)) {
                                                v3 = (v3 >> ((65535) & 63))
                                                switch (fi63 * (fi63 << ((331) & 63))) {
                                                case -661:
                                                    var v105 uint8 = (^(v3))
                                                    _ = v105
                                                case 127:
                                                    var v106 float32 = (1e10 / float32((3.14159 + 0.0)))
                                                    _ = v106
                                                case 128:
                                                    if len(v2) > 0 {
                                                        v2[0] = 4294967295
                                                    }
                                                default:
                                                    var v107 int32 = 42
                                                    _ = v107
                                                }
                                            } else {
                                                var v108 int32 = 1
                                                _ = v108
                                            }
                                        case 501:
                                            switch "\n\t" {
                                            case "\n\t":
                                                if ((2147483648 == 2) || ((32767 != fi63) && (float32(float64(-2.25)) >= (0.0 + 0.0)))) {
                                                    v109 := GenS6{F0: ((("foo bar" + "foo bar") + "\n\t") + "hello")}
                                                    _ = v109
                                                }
                                            case "a":
                                                var v110 uint16 = 0
                                                _ = v110
                                            default:
                                                fmt.Println(fi67, fi63)
                                            }
                                        }
                                    case 32768:
                                        v3 += (v3 >> ((10) & 63))
                                    case 0:
                                        if true {
                                            fmt.Println(fi63)
                                            fmt.Println(v66, fi63, fi67)
                                            var v111 float64 = (0.0 / 3.14159)
                                            _ = v111
                                        }
                                    default:
                                        if len(v2) > 0 {
                                            v2[0] = 2147483648
                                        }
                                        if ((v3 - 77) >= v3) {
                                            if ((fi67 << ((65535) & 63)) > (fi67 * 861)) {
                                                if len(v2) > 0 {
                                                    v2[0] = 53
                                                }
                                                v112 := GenS11{F0: (v1 >> ((4294967295) & 63)), F1: 190, F2: -1}
                                                v113 := GenS12{F0: 134}
                                                _ = v112
                                                _ = v113
                                            }
                                            if false {
                                                var v114 int64 = v104
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                var v115 int64 = v104
                                                _ = v114
                                                _ = v115
                                            }
                                            var v116 uint64 = 42
                                            _ = v116
                                        }
                                        var v117 uint32 = 2147483647
                                        var v118 string = ("\n\t" + "foo bar")
                                        _ = v117
                                        _ = v118
                                    }
                                    var v119 uint64 = 0
                                    var v120 int32 = 42
                                    _ = v104
                                    _ = v119
                                    _ = v120
                                }
                                v121 := GenS9{F0: 32767, F1: genFn0(561), F2: 130}
                                _ = v121
                            } else {
                                var v122 uint = 679
                                _ = v122
                            }
                        } else {
                            v66 += 72
                            v3 = (v3 * (^(v3)))
                        }
                        if (1.5 != 0.0) {
                            v1 = v66
                        } else {
                            if len(v2) > 0 {
                                v2[0] = 65535
                            }
                        }
                    }
                    var v123 int64 = -9223372036854775808
                    _ = fi67
                    _ = v123
                    _ = v68
                }
                _ = fi63
                _ = v66
            }
            var v124 string = "foo bar"
            _ = v124
        }
    }
    var v125 uint8 = (v3 - v3)
    switch 128 {
    case 127:
        var v126 uint8 = 51
        _ = v126
    case 2147483647:
        v3 = v125
    default:
        v127 := GenS8{F0: 185, F1: 42, F2: -1}
        _ = v127
    }
    if (!(-1 == -125)) {
        var v128 string = ("" + "")
        _ = v128
    } else {
        v125 = 0
        var v129 bool = (genFn0(450) < genFn0(180))
        _ = v129
    }
    for fi130 := 0; fi130 < 10; fi130++ {
        var v131 uint = 478
        var v132 int = fi130
        _ = fi130
        _ = v131
        _ = v132
    }
    v1 += 100
    v125 -= 0
    if false {
        if true {
            if (272 == 42) {
                var v133 int64 = 1
                v134 := GenS13{F0: (v1 | (v1 ^ v1)), F1: 14}
                if (genFn0(714) >= 1) {
                    v1 = (v1 | (v1 % (((v1 / ((v1) | 1))) | 1)))
                    var v135 uint8 = v125
                    for fi136 := 0; fi136 < 8; fi136++ {
                        var v137 int32 = -163
                        _ = fi136
                        _ = v137
                    }
                    _ = v135
                }
                _ = v133
                _ = v134
            } else {
                var v138 int = 42
                if ((v1 & (^(v1))) != (v1 - v1)) {
                    for fi139 := 0; fi139 < 1; fi139++ {
                        if false {
                            v138 -= v138
                            switch ("" + "\n\t") {
                            case "":
                                _ = len(v2)
                            case "foo bar":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            default:
                                v140 := GenS14{F0: "foo bar"}
                                _ = v140
                            }
                            for fi141 := 0; fi141 < 4; fi141++ {
                                if true {
                                    v138 += (v138 / ((-1) | 1))
                                    v125 = (v125 >> ((272) & 63))
                                    v142 := GenS15{F0: (v1 / ((v1) | 1)), F1: 0}
                                    var v143 float64 = 1e10
                                    _ = v142
                                    _ = v143
                                } else {
                                    v138 += fi139
                                    var v144 uint = 882
                                    v145 := GenS1{F0: (v125 % ((v125) | 1))}
                                    _ = v144
                                    _ = v145
                                }
                                v146 := GenS4{F0: "\n\t"}
                                _ = fi141
                                _ = v146
                            }
                        }
                        for fi147 := 0; fi147 < 3; fi147++ {
                            var v148 uint = 51
                            _ = fi147
                            _ = v148
                        }
                        var v149 float32 = (((0.0 * 0.1) * 0.0) - (0.0 / 1.5))
                        _ = fi139
                        _ = v149
                    }
                }
                for fi150 := 0; fi150 < 11; fi150++ {
                    if len(v2) > 0 {
                        v2[0] = 42
                    }
                    v125 -= (v3 + v125)
                    _ = fi150
                }
                _ = v138
            }
        } else {
            for fi151 := 0; fi151 < 9; fi151++ {
                v2 = append(v2, 654)
                for fi152 := 0; fi152 < 3; fi152++ {
                    v153 := GenS16{F0: 2147483647}
                    var v154 uint8 = (^(v3))
                    if true {
                        switch (^(fi152)) {
                        case 1:
                            var v155 float32 = 3.14159
                            _ = v155
                        case 824:
                            var v156 bool = true
                            _ = v156
                        case -330:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        v157 := GenS17{F0: -1, F1: 127, F2: -766}
                        v125 += (v3 * v154)
                        _ = v157
                    }
                    _ = fi152
                    _ = v153
                    _ = v154
                }
                _ = fi151
            }
            for fi158 := 0; fi158 < 10; fi158++ {
                if false {
                    for fi159 := 0; fi159 < 7; fi159++ {
                        switch fi159 {
                        case 0:
                            v2 = append(v2, 0)
                        }
                        switch (("\n\t" + "hello") + ("" + "foo bar")) {
                        case "\n\t":
                            v3 = (v3 * (^(v3)))
                        default:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        var v160 uint16 = 386
                        _ = fi159
                        _ = v160
                    }
                    var v161 int = fi158
                    _ = v161
                }
                var v162 int16 = -32768
                _ = fi158
                _ = v162
            }
            var v163 int64 = -1
            _ = v163
        }
    } else {
        if false {
            v164 := GenS17{F0: 1, F1: 465, F2: -1}
            v3 = 0
            if len(v2) > 0 {
                v2[0] = 1
            }
            _ = v164
        }
    }
    var v165 string = "a"
    v3 = (^(v3))
    v3 += 2
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v125:", v125)
    fmt.Println("v165:", v165)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v125
    _ = v165
    _ = v2
    _ = v3
}
