package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 int32; F1 int64; F2 int }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 string; F1 int64 }
type GenS7 struct { F0 string; F1 bool }
type GenS8 struct { F0 uint; F1 int32; F2 uint }
type GenS9 struct { F0 uint8; F1 int64; F2 int64 }
type GenS10 struct { F0 int32; F1 uint32 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 string; F1 bool; F2 uint8 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 uint }
func genFn0(p0 uint64) uint32 {
    var v6 uint = 18446744073709551615
    v6 = uint(p0)
    _ = p0
    _ = v6
    return 2
}
func genFn1(p0 int64) uint64 {
    var v7 int = 574
    v8 := []int{v7, v7}
    _ = p0
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 string = ("" + ("a" + "a"))
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]uint32)
    var v4 int64 = 0
    v5 := []string{v1, "foo bar", "hello", (v1 + v1)}
    v9 := v3[4294967295]
    if true {
        switch "a" {
        case "":
            v1 = v1
        case "hello":
            if true {
                var v10 uint = 2
                if false {
                    v1 = "\n\t"
                    fmt.Println(v4)
                    v1 = v1
                } else {
                    fmt.Println(v9)
                    v9 = v9
                    v11 := GenS1{F0: 417, F1: uint16(v10), F2: 42}
                    _ = v11
                }
                _ = v10
            }
        case "\n\t":
            fmt.Println(v1)
        }
    } else {
        if (!((3.14159 * (1.5 - -2.25)) == -2.25)) {
            v12 := GenS2{F0: 18446744073709551615}
            var v13 bool = false
            var v14 int8 = 1
            _ = v12
            _ = v13
            _ = v14
        } else {
            if (!((!true) || true)) {
                fmt.Println(v9, v1)
                switch "hello" {
                case "":
                    var v15 string = v1
                    _ = v15
                case "\n\t":
                    v16 := v3[(v9 >> ((759) & 63))]
                    _ = v16
                default:
                    v17 := GenS2{F0: 0}
                    _ = v17
                }
            }
            if (-1 <= -4) {
                switch v1 {
                case "foo bar":
                    for fi18 := 0; fi18 < 8; fi18++ {
                        fmt.Println(fi18)
                        if false {
                            v4 += -878
                        }
                        _ = fi18
                    }
                default:
                    v19 := GenS3{F0: 42, F1: false}
                    _ = v19
                }
            }
        }
        fmt.Println(v9, v4, v1)
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        if (v4 != (v4 ^ v4)) {
            v21 := GenS1{F0: 42, F1: 349, F2: 1}
            _ = v21
        } else {
            if ("hello" >= v1) {
                for fi22 := 0; fi22 < 4; fi22++ {
                    v23 := GenS2{F0: uint64(v9)}
                    _ = fi22
                    _ = v23
                }
                var v24 uint8 = 195
                for fi25 := 0; fi25 < 17; fi25++ {
                    v26 := v2[1]
                    if (!false) {
                        v4 = (v4 >> ((0) & 63))
                        _ = len(v2)
                        var v27 int16 = 621
                        v28 := GenS3{F0: 604, F1: true}
                        _ = v27
                        _ = v28
                    }
                    _ = fi25
                    _ = v26
                }
                _ = v24
            }
            v29 := GenS1{F0: -32768, F1: 1, F2: -128}
            _ = v29
        }
        v9 = 4294967295
        for fi30 := 0; fi30 < 14; fi30++ {
            var v31 uint64 = genFn1(381)
            if (-605 == 0) {
                if false {
                    v32 := GenS1{F0: -1, F1: 988, F2: -55}
                    var v33 uint32 = 207
                    if ((true && true) && true) {
                        fmt.Println(v1)
                        var v34 uint16 = 0
                        _ = v34
                    }
                    _ = v32
                    _ = v33
                } else {
                    v35 := GenS2{F0: genFn1((v4 << ((2) & 63)))}
                    if false {
                        v31 -= v31
                        var v36 int64 = int64(fi20)
                        fmt.Println(v36, v4)
                        _ = v36
                    }
                    v3[(v9 ^ (v9 / ((42) | 1)))] = 2
                    _ = v35
                }
                switch (fi30 - fi20) {
                case 32768:
                    v4 += (v4 - v4)
                }
                if (1 > 42) {
                    var v37 uint32 = (v9 + genFn0((v31 - genFn1((^(v4))))))
                    v38 := GenS2{F0: (v31 / ((2) | 1))}
                    var v39 bool = false
                    _ = v37
                    _ = v38
                    _ = v39
                }
            }
            _ = fi30
            _ = v31
        }
        _ = fi20
    }
    fmt.Println(v4, v9, v9)
    for fi40 := 0; fi40 < 20; fi40++ {
        v1 = (v1 + v1)
        for fi41 := 0; fi41 < 11; fi41++ {
            if (true || (!(true || true))) {
                if (!true) {
                    switch (fi41 / (((fi41 >> ((18446744073709551615) & 63))) | 1)) {
                    case -2147483648:
                        if (0 < 0) {
                            if true {
                                switch v1 {
                                case "foo bar":
                                    if false {
                                        var v42 bool = true
                                        _ = v42
                                    } else {
                                        _ = len(v3)
                                        delete(v3, v9)
                                        v9 += (v9 | (v9 | (v9 | v9)))
                                    }
                                case "\n\t":
                                    v43 := GenS2{F0: genFn1(9223372036854775807)}
                                    _ = v43
                                case "":
                                    fmt.Println(v1, fi41, v9)
                                }
                                var v44 uint32 = 42
                                v45 := GenS1{F0: -232, F1: 42, F2: 100}
                                v9 = 211
                                var v46 uint = 42
                                v44 = v9
                                v44 = v9
                                _ = v44
                                _ = v45
                                _ = v46
                            }
                        } else {
                            v4 -= v4
                        }
                    case 42:
                        var v47 int64 = v4
                        fmt.Println(v1, v9, v1)
                        switch (fi41 / ((0) | 1)) {
                        case -1:
                            v1 = ""
                        case -9223372036854775808:
                            var v48 uint32 = genFn0(881)
                            _ = v48
                        }
                        _ = v47
                    }
                    if (-1 != -128) {
                        var v49 string = v1
                        var v50 float32 = 3.14159
                        _ = v49
                        _ = v50
                    }
                }
                if (32767 < uint16(v9)) {
                    if (425 > 1) {
                        if (127 >= 243) {
                            v51 := GenS4{F0: -971, F1: -1, F2: -421}
                            v1 = v1
                            var v52 float32 = (1.5 - (3.14159 * 1.5))
                            _ = v51
                            _ = v52
                        }
                        if (!true) {
                            _ = len(v2)
                            delete(v2, 217)
                        } else {
                            v9 = (v9 | genFn0(4294967295))
                            v4 -= v4
                            if (genFn0(2) < 4294967295) {
                                v4 = (v4 % ((v4) | 1))
                            } else {
                                var v53 float64 = 0.1
                                var v54 string = v1
                                switch (v54 + (((v1 + "\n\t") + v1) + "\n\t")) {
                                case "hello":
                                    if false {
                                        if false {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        } else {
                                            var v55 float32 = ((1e10 + 1e10) / 0.1)
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            var v56 int32 = -1
                                            _ = v55
                                            _ = v56
                                        }
                                    }
                                case "\n\t":
                                    if false {
                                        v57 := GenS1{F0: 879, F1: 42, F2: 42}
                                        var v58 uint32 = (v9 % ((v9) | 1))
                                        v59 := GenS5{F0: 369, F1: 0}
                                        _ = v57
                                        _ = v58
                                        _ = v59
                                    } else {
                                        var v60 int32 = -760
                                        v5 = append(v5, v54)
                                        v61 := GenS5{F0: -1, F1: (v60 ^ v60)}
                                        _ = v60
                                        _ = v61
                                    }
                                case "foo bar":
                                    v62 := GenS1{F0: 195, F1: 42, F2: -1}
                                    _ = v62
                                }
                                var v63 bool = false
                                fmt.Println(v63)
                                _ = v53
                                _ = v54
                                _ = v63
                            }
                        }
                    }
                    v64 := GenS3{F0: 1, F1: (!false)}
                    _ = v64
                } else {
                    switch fi41 {
                    case 32767:
                        var v65 int = (fi40 / (((fi41 % ((fi40) | 1))) | 1))
                        _ = v65
                    }
                    if (true || true) {
                        v66 := GenS1{F0: 0, F1: 371, F2: int8(v4)}
                        v4 = v4
                        fmt.Println(fi40)
                        _ = v66
                    }
                }
            } else {
                if false {
                    fmt.Println(v4, fi40)
                    var v67 int16 = 42
                    var v68 uint32 = (v9 >> ((452) & 63))
                    var v69 string = (v1 + (v1 + "hello"))
                    var v70 bool = (false && false)
                    _ = v67
                    _ = v68
                    _ = v69
                    _ = v70
                } else {
                    var v71 uint16 = 2
                    var v72 uint16 = (v71 % ((v71) | 1))
                    if false {
                        var v73 uint64 = 900
                        if true {
                            v74 := GenS2{F0: (v73 % (((v73 - v73)) | 1))}
                            v75 := GenS6{F0: (v1 + (v1 + "hello")), F1: (-(v4))}
                            v9 += (v9 | genFn0(v73))
                            var v76 int16 = 1000
                            _ = v74
                            _ = v75
                            _ = v76
                        } else {
                            v71 += v72
                            fmt.Println(v1, fi40, v73)
                        }
                        _ = v73
                    }
                    fmt.Println(v71, v4)
                    _ = v71
                    _ = v72
                }
                v2[679] = (127 == 91)
                v1 = "\n\t"
            }
            _ = fi41
        }
        switch fi40 {
        case 0:
            for fi77 := 0; fi77 < 8; fi77++ {
                var v78 bool = (-96 > -1)
                if (0 > 42) {
                    v9 = genFn0(4294967295)
                }
                _ = fi77
                _ = v78
            }
        }
        _ = fi40
    }
    for fi79 := 0; fi79 < 14; fi79++ {
        var v80 bool = (1e10 > float64((0.1 * (3.14159 + -2.25))))
        _ = fi79
        _ = v80
    }
    var v81 int16 = -116
    v5 = append(v5, v1)
    var v82 int32 = 42
    var v83 uint8 = 1
    switch 42 {
    case 128:
        for fi84 := 0; fi84 < 19; fi84++ {
            v1 = v1
            _ = fi84
        }
    case 71:
        var v85 string = v1
        _ = v85
    default:
        for fi86 := 0; fi86 < 19; fi86++ {
            if (!false) {
                for fi87 := 0; fi87 < 2; fi87++ {
                    v88 := GenS3{F0: 0, F1: false}
                    if (uint16(v9) <= 255) {
                        var v89 int8 = -59
                        var v90 uint8 = (v83 | v83)
                        v89 -= 99
                        _ = v89
                        _ = v90
                    } else {
                        var v91 bool = ((3.14159 + 1e10) <= float64(-2.25))
                        var v92 bool = (1 > 1)
                        switch (-(fi86)) {
                        case 32767:
                            if ((0.1 + (0.0 - float64(((0.1 * -2.25) - (1.5 - 1e10))))) == 3.14159) {
                                switch ("a" + "\n\t") {
                                case "foo bar":
                                    var v93 uint64 = 0
                                    _ = v93
                                case "":
                                    switch (fi86 - 1) {
                                    case -692:
                                        v92 = false
                                    case 862:
                                        v92 = (!v91)
                                    case 32768:
                                        v4 += -906
                                    }
                                case "a":
                                    v1 = "hello"
                                }
                                if false {
                                    v94 := GenS6{F0: "a", F1: (v4 << ((1) & 63))}
                                    v81 -= v81
                                    delete(v2, uint32(v83))
                                    _ = v94
                                }
                                v5 = append(v5, v1)
                                switch ("" + (v1 + (((((v1 + v1) + (v1 + v1)) + ((("hello" + v1) + "") + (v1 + v1))) + v1) + v1))) {
                                case "a":
                                    v95 := GenS5{F0: (-(v4)), F1: 927}
                                    _ = v95
                                default:
                                    v96 := GenS2{F0: 48}
                                    _ = v96
                                }
                            } else {
                                v83 = v83
                            }
                        }
                        _ = v91
                        _ = v92
                    }
                    switch v1 {
                    case "\n\t":
                        if true {
                            _ = len(v3)
                            if (fi87 == 786) {
                                var v97 float64 = -2.25
                                v98 := GenS3{F0: 65535, F1: (2 < 303)}
                                v9 = genFn0(415)
                                _ = v97
                                _ = v98
                            }
                            var v99 string = ((v1 + v1) + v1)
                            _ = v99
                        } else {
                            if false {
                                v83 = v83
                                v100 := v3[2147483648]
                                _ = v100
                            } else {
                                var v101 bool = true
                                _ = v101
                            }
                            switch ((v1 + v1) + "a") {
                            case "\n\t":
                                v102 := GenS1{F0: v81, F1: 0, F2: int8(v83)}
                                if false {
                                    v103 := GenS4{F0: -2147483648, F1: 42, F2: fi87}
                                    v104 := GenS5{F0: (v4 - (v4 / ((v4) | 1))), F1: -883}
                                    v5 = append(v5, "a")
                                    delete(v2, genFn0(863))
                                    v105 := GenS5{F0: v4, F1: (-(v82))}
                                    _ = v103
                                    _ = v104
                                    _ = v105
                                }
                                var v106 uint = 517
                                _ = v102
                                _ = v106
                            case "a":
                                var v107 bool = true
                                _ = v107
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                    case "a":
                        v1 = v1
                    case "foo bar":
                        if (v1 != v1) {
                            switch (v1 + ((v1 + (("hello" + v1) + "\n\t")) + v1)) {
                            case "a":
                                var v108 string = v1
                                _ = v108
                            case "hello":
                                fmt.Println(v1)
                            default:
                                var v109 int = fi86
                                _ = v109
                            }
                            var v110 string = (v1 + v1)
                            _ = v110
                        } else {
                            delete(v2, v9)
                            var v111 float32 = float32(3.14159)
                            _ = v111
                        }
                    }
                    _ = fi87
                    _ = v88
                }
                if false {
                    for fi112 := 0; fi112 < 14; fi112++ {
                        v113 := v3[genFn0(88)]
                        switch fi86 {
                        case 0:
                            v114 := GenS2{F0: genFn1(42)}
                            _ = v114
                        default:
                            v115 := GenS2{F0: uint64(v83)}
                            _ = v115
                        }
                        fmt.Println(fi86)
                        _ = fi112
                        _ = v113
                    }
                }
            } else {
                v116 := v2[(v9 ^ (v9 - 2))]
                _ = v116
            }
            var v117 uint16 = 801
            v118 := GenS4{F0: (v82 << ((668) & 63)), F1: (v4 ^ v4), F2: 0}
            _ = fi86
            _ = v117
            _ = v118
        }
    }
    switch (v1 + "") {
    case "foo bar":
        for fi119 := 0; fi119 < 5; fi119++ {
            v81 = (v81 | v81)
            for fi120 := 0; fi120 < 2; fi120++ {
                v121 := GenS4{F0: v82, F1: v4, F2: fi120}
                if (false || (true && (0.1 != -2.25))) {
                    if (619 == uint16(v83)) {
                        var v122 bool = false
                        _ = v122
                    }
                    switch (fi119 * 732) {
                    case 596:
                        v123 := GenS3{F0: 1, F1: ("hello" == ((v1 + v1) + v1))}
                        _ = v123
                    default:
                        if (77 != 1) {
                            v124 := GenS7{F0: (("a" + "") + ((v1 + ((v1 + v1) + "foo bar")) + ((v1 + v1) + (((v1 + ((v1 + v1) + v1)) + "") + "\n\t")))), F1: (true && (1 >= 320))}
                            delete(v2, 0)
                            _ = v124
                        } else {
                            var v125 string = "a"
                            if true {
                                v125 = "\n\t"
                            } else {
                                var v126 float32 = 0.1
                                switch fi119 {
                                case 2147483647:
                                    var v127 string = v1
                                    _ = v127
                                case 0:
                                    var v128 int32 = (v82 | v82)
                                    _ = v128
                                case 32767:
                                    var v129 string = (v1 + "a")
                                    _ = v129
                                default:
                                    if ((-2.25 + 0.1) < -2.25) {
                                        switch fi120 {
                                        case -529:
                                            switch (fi119 >> ((1) & 63)) {
                                            case -1:
                                                fmt.Println(fi120, v83)
                                            }
                                            switch (fi120 / (((fi120 ^ (-(fi119)))) | 1)) {
                                            case 957:
                                                var v130 int32 = int32(v81)
                                                _ = v130
                                            case 855:
                                                v126 = v126
                                            default:
                                                v126 += (v126 * -2.25)
                                            }
                                            if (!(false && (0.0 != float64(0.1)))) {
                                                var v131 uint64 = 0
                                                if false {
                                                    delete(v3, genFn0(v131))
                                                    var v132 uint = 42
                                                    v133 := GenS4{F0: int32(v4), F1: v4, F2: (fi120 * fi120)}
                                                    var v134 bool = true
                                                    _ = v132
                                                    _ = v133
                                                    _ = v134
                                                } else {
                                                    v135 := GenS3{F0: 42, F1: true}
                                                    var v136 string = "foo bar"
                                                    _ = v135
                                                    _ = v136
                                                }
                                                _ = v131
                                            }
                                        case -1:
                                            if false {
                                                var v137 uint64 = genFn1((-(v4)))
                                                v4 += 138
                                                _ = v137
                                            }
                                        case -114:
                                            v5 = append(v5, v1)
                                        default:
                                            if len(v5) > 0 {
                                                v5[0] = v125
                                            }
                                        }
                                        v1 = ((v1 + "hello") + v125)
                                        if (((v126 - 1.5) + (v126 - 1e10)) <= float32(3.14159)) {
                                            v126 -= v126
                                        } else {
                                            var v138 int8 = 0
                                            var v139 string = (((v1 + v125) + "hello") + ((v125 + "") + v125))
                                            _ = v138
                                            _ = v139
                                        }
                                    } else {
                                        switch v1 {
                                        case "foo bar":
                                            var v140 uint8 = (v83 * 246)
                                            _ = v140
                                        }
                                    }
                                }
                                _ = v126
                            }
                            switch (fi120 << ((599) & 63)) {
                            case 42:
                                if false {
                                    var v141 uint8 = v83
                                    _ = v141
                                } else {
                                    fmt.Println(v82)
                                    if ((false && (v81 != 1000)) && (!true)) {
                                        v9 += 2147483648
                                        var v142 uint64 = 0
                                        v81 += v81
                                        v1 = ((v1 + ("\n\t" + "hello")) + ("\n\t" + v125))
                                        var v143 int32 = (v82 & (v82 % ((v82) | 1)))
                                        _ = v142
                                        _ = v143
                                    }
                                    delete(v2, (v9 - 0))
                                }
                            case 32768:
                                switch fi119 {
                                case -2147483648:
                                    v144 := GenS2{F0: genFn1(v4)}
                                    _ = v144
                                case 1:
                                    switch fi120 {
                                    case 128:
                                        v81 = 1000
                                    case 0:
                                        var v145 uint32 = v9
                                        _ = v145
                                    default:
                                        var v146 float32 = (float32(1.5) * 1e10)
                                        _ = v146
                                    }
                                default:
                                    v147 := GenS8{F0: 4294967295, F1: -755, F2: 273}
                                    _ = v147
                                }
                            default:
                                var v148 int32 = v82
                                _ = v148
                            }
                            _ = v125
                        }
                    }
                } else {
                    v5 = append(v5, v1)
                    switch "" {
                    case "\n\t":
                        if (true && (!false)) {
                            v81 += v81
                        } else {
                            if true {
                                _ = len(v5)
                                switch (v1 + (((v1 + (v1 + v1)) + "a") + (("hello" + ("" + v1)) + v1))) {
                                case "a":
                                    v1 = v1
                                default:
                                    if true {
                                        v83 += v83
                                    } else {
                                        if (float64(1e10) == ((3.14159 / (-2.25 / (3.14159 / -2.25))) / 1.0)) {
                                            switch ("" + v1) {
                                            case "":
                                                v81 = (v81 & v81)
                                            case "a":
                                                if (!true) {
                                                    if true {
                                                        if (220 == 32767) {
                                                            if false {
                                                                v149 := GenS9{F0: v83, F1: 768, F2: (^(v4))}
                                                                fmt.Println(v1, v82, v81)
                                                                _ = v149
                                                            } else {
                                                                switch (fi120 >> ((2) & 63)) {
                                                                case 128:
                                                                    v9 += (v9 & 2147483647)
                                                                case 4294967295:
                                                                    v150 := GenS1{F0: v81, F1: 32767, F2: 100}
                                                                    _ = v150
                                                                default:
                                                                    var v151 uint64 = 2
                                                                    _ = v151
                                                                }
                                                                v152 := GenS5{F0: 9223372036854775807, F1: (v82 % (((v82 * v82)) | 1))}
                                                                v1 = v1
                                                                _ = v152
                                                            }
                                                            var v153 int32 = (v82 | v82)
                                                            var v154 int64 = v4
                                                            _ = v153
                                                            _ = v154
                                                        } else {
                                                            fmt.Println(v4)
                                                            v82 += 296
                                                            v81 += 32767
                                                            switch (v1 + "\n\t") {
                                                            case "":
                                                                switch "" {
                                                                case "a":
                                                                    var v155 uint32 = v9
                                                                    _ = v155
                                                                case "":
                                                                    v1 = v1
                                                                case "\n\t":
                                                                    v2[(v9 << ((65535) & 63))] = false
                                                                    switch ("foo bar" + "hello") {
                                                                    case "hello":
                                                                        v4 = (v4 * (v4 | (v4 % ((v4) | 1))))
                                                                    }
                                                                }
                                                            case "hello":
                                                                v82 = v82
                                                                fmt.Println(fi120, v1)
                                                                switch (fi119 * fi119) {
                                                                case -9223372036854775808:
                                                                    v83 = v83
                                                                case -323:
                                                                    v82 -= (v82 ^ (v82 + (v82 & v82)))
                                                                default:
                                                                    v82 -= v82
                                                                }
                                                                v9 += (v9 - genFn0(0))
                                                            case "foo bar":
                                                                v156 := GenS10{F0: v82, F1: v9}
                                                                _ = v156
                                                            }
                                                        }
                                                    } else {
                                                        switch ((v1 + v1) + ("foo bar" + "")) {
                                                        case "":
                                                            if ((0.1 + -2.25) == (0.1 / 3.14159)) {
                                                                v1 = v1
                                                            }
                                                        case "\n\t":
                                                            v157 := GenS9{F0: v83, F1: (v4 / (((v4 % ((-994) | 1))) | 1)), F2: v4}
                                                            _ = v157
                                                        }
                                                        v158 := v3[v9]
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        if true {
                                                            var v159 int16 = v81
                                                            _ = v159
                                                        }
                                                        var v160 float64 = -2.25
                                                        _ = v158
                                                        _ = v160
                                                    }
                                                    fmt.Println(v81, v81)
                                                }
                                                var v161 float64 = (0.0 - 1e10)
                                                _ = v161
                                            default:
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                if ((!false) || false) {
                                                    fmt.Println(fi120, v4, v4)
                                                    fmt.Println(v81)
                                                    v162 := GenS11{F0: (v81 / (((v81 | v81)) | 1))}
                                                    _ = v162
                                                } else {
                                                    var v163 uint8 = 255
                                                    fmt.Println(v81)
                                                    _ = v163
                                                }
                                            }
                                            var v164 float64 = 3.14159
                                            if false {
                                                v82 = v82
                                                var v165 bool = true
                                                var v166 float32 = (((1e10 + float32(0.0)) + 0.1) * 1.5)
                                                var v167 int64 = -157
                                                var v168 float32 = 3.14159
                                                var v169 uint = 171
                                                var v170 int64 = -387
                                                _ = v165
                                                _ = v166
                                                _ = v167
                                                _ = v168
                                                _ = v169
                                                _ = v170
                                            } else {
                                                v83 += v83
                                                if len(v5) > 0 {
                                                    v5[0] = "\n\t"
                                                }
                                                var v171 int16 = v81
                                                var v172 uint = 30
                                                _ = v171
                                                _ = v172
                                            }
                                            _ = v164
                                        }
                                        if false {
                                            v173 := GenS9{F0: v83, F1: (v4 + 0), F2: (v4 << ((0) & 63))}
                                            if ((1.5 - 3.14159) > 0.0) {
                                                fmt.Println(v83, v83)
                                                v3[v9] = v9
                                                if false {
                                                    switch (-(fi120)) {
                                                    case -167:
                                                        var v174 int16 = -1
                                                        _ = v174
                                                    case -1:
                                                        v175 := GenS5{F0: 1, F1: int32(v9)}
                                                        _ = v175
                                                    case 229:
                                                        switch fi119 {
                                                        case 9223372036854775807:
                                                            fmt.Println(fi119, v83, fi119)
                                                            var v176 uint8 = uint8(v9)
                                                            _ = v176
                                                        case 42:
                                                            v177 := GenS2{F0: 270}
                                                            _ = v177
                                                        case 21:
                                                            v83 += v83
                                                            var v178 float32 = -2.25
                                                            v81 = (^(v81))
                                                            if (v1 != v1) {
                                                                fmt.Println(v81, v83)
                                                                delete(v3, (v9 / (((v9 / (((v9 - (v9 | v9))) | 1))) | 1)))
                                                            } else {
                                                                fmt.Println(v178)
                                                            }
                                                            v82 = (v82 - v82)
                                                            v179 := v2[genFn0(226)]
                                                            switch fi119 {
                                                            case 9223372036854775807:
                                                                v81 += v81
                                                            case -941:
                                                                v180 := GenS12{F0: ("foo bar" + (("foo bar" + ((v1 + v1) + v1)) + v1)), F1: v179, F2: 255}
                                                                _ = v180
                                                            default:
                                                                if v179 {
                                                                    var v181 uint16 = 42
                                                                    _ = v181
                                                                } else {
                                                                    var v182 uint16 = 0
                                                                    fmt.Println(v178, v4)
                                                                    v183 := GenS8{F0: 255, F1: (v82 + (v82 >> ((4294967295) & 63))), F2: 18446744073709551615}
                                                                    _ = v182
                                                                    _ = v183
                                                                }
                                                            }
                                                            if (2 == 1) {
                                                                if v179 {
                                                                    var v184 uint8 = v83
                                                                    _ = v184
                                                                }
                                                            }
                                                            var v185 string = ("hello" + v1)
                                                            _ = v178
                                                            _ = v179
                                                            _ = v185
                                                        }
                                                    }
                                                    var v186 string = v1
                                                    switch v1 {
                                                    case "hello":
                                                        v187 := GenS13{F0: (!(122 >= -1))}
                                                        _ = v187
                                                    case "foo bar":
                                                        v81 += v81
                                                    case "a":
                                                        if true {
                                                            var v188 int8 = 1
                                                            var v189 uint32 = genFn0(4294967295)
                                                            _ = v188
                                                            _ = v189
                                                        } else {
                                                            v4 -= v4
                                                            var v190 int = (fi120 >> ((65535) & 63))
                                                            if (42 >= 65535) {
                                                                var v191 uint32 = (v9 << ((uint(v9)) & 63))
                                                                v186 = v1
                                                                delete(v2, (v191 * 456))
                                                                _ = v191
                                                            }
                                                            _ = v190
                                                        }
                                                    }
                                                    _ = v186
                                                } else {
                                                    fmt.Println(v81, v9, v9)
                                                    if true {
                                                        v192 := GenS14{F0: uint(v9)}
                                                        var v193 bool = ((0.1 / 1.5) != (0.0 * ((3.14159 / 1.0) - -2.25)))
                                                        switch fi120 {
                                                        case -2147483648:
                                                            switch v1 {
                                                            case "hello":
                                                                fmt.Println(v81, v4, fi119)
                                                                _ = len(v5)
                                                                v194 := GenS7{F0: "", F1: v193}
                                                                _ = v194
                                                            case "\n\t":
                                                                if len(v5) > 0 {
                                                                    v5[0] = v1
                                                                }
                                                            case "":
                                                                v82 = 2147483647
                                                            }
                                                        default:
                                                            v195 := GenS13{F0: (!((((0.1 * (1.5 + (float64((1.5 / 0.1)) * 1e10))) + 1e10) * 3.14159) >= 3.14159))}
                                                            _ = v195
                                                        }
                                                        _ = v192
                                                        _ = v193
                                                    }
                                                    switch (fi119 * 2147483647) {
                                                    case -1:
                                                        var v196 int64 = 1
                                                        _ = v196
                                                    case -47:
                                                        v197 := GenS14{F0: 65535}
                                                        _ = v197
                                                    }
                                                    var v198 float64 = 0.1
                                                    switch 0 {
                                                    case 2147483647:
                                                        var v199 uint = 223
                                                        _ = v199
                                                    case -9223372036854775808:
                                                        fmt.Println(v82)
                                                    case 172:
                                                        fmt.Println(v9, v82, v9)
                                                    default:
                                                        v83 += (v83 << ((689) & 63))
                                                    }
                                                    v200 := GenS2{F0: genFn1(v4)}
                                                    v81 += (^(v81))
                                                    _ = v198
                                                    _ = v200
                                                }
                                            } else {
                                                switch "a" {
                                                case "foo bar":
                                                    fmt.Println(v4)
                                                case "\n\t":
                                                    var v201 int32 = 8
                                                    _ = v201
                                                case "hello":
                                                    switch v1 {
                                                    case "":
                                                        var v202 uint64 = 0
                                                        _ = v202
                                                    case "hello":
                                                        var v203 int = int(v82)
                                                        _ = v203
                                                    default:
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    }
                                                default:
                                                    var v204 uint16 = 0
                                                    _ = v204
                                                }
                                            }
                                            var v205 int = (fi120 & fi119)
                                            if (-2.25 <= -2.25) {
                                                var v206 int8 = -1
                                                switch "foo bar" {
                                                case "\n\t":
                                                    var v207 string = v1
                                                    _ = v207
                                                default:
                                                    fmt.Println(v205)
                                                }
                                                if len(v5) > 0 {
                                                    v5[0] = v1
                                                }
                                                var v208 int8 = (v206 - v206)
                                                if (false || (false && false)) {
                                                    fmt.Println(v83)
                                                }
                                                _ = v206
                                                _ = v208
                                            }
                                            _ = v173
                                            _ = v205
                                        } else {
                                            v209 := GenS3{F0: 1, F1: (2 > 1)}
                                            _ = v209
                                        }
                                        var v210 uint8 = (v83 | 42)
                                        _ = v210
                                    }
                                }
                                fmt.Println(v81)
                            } else {
                                if false {
                                    switch (fi120 ^ (fi119 % ((9223372036854775807) | 1))) {
                                    case 665:
                                        v211 := GenS11{F0: (v81 / ((v81) | 1))}
                                        _ = v211
                                    default:
                                        v9 += uint32(v83)
                                    }
                                    v212 := GenS10{F0: 0, F1: v9}
                                    _ = v212
                                } else {
                                    v213 := v2[(v9 | (v9 & v9))]
                                    var v214 bool = (!true)
                                    delete(v2, genFn0(1))
                                    _ = v213
                                    _ = v214
                                }
                            }
                            v9 = (^(v9))
                            var v215 uint8 = (^(v83))
                            _ = v215
                        }
                    default:
                        v4 = v4
                    }
                    if false {
                        v216 := GenS1{F0: (v81 % (((v81 ^ v81)) | 1)), F1: 32767, F2: -1}
                        _ = v216
                    } else {
                        fmt.Println(v4, v1, fi119)
                        _ = len(v3)
                        v217 := GenS5{F0: 9223372036854775807, F1: (v82 % ((v82) | 1))}
                        var v218 string = "a"
                        _ = v217
                        _ = v218
                    }
                }
                _ = fi120
                _ = v121
            }
            _ = fi119
        }
    case "hello":
        if true {
            v83 = (v83 * (v83 / (((v83 / ((v83) | 1))) | 1)))
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    case "\n\t":
        v4 = v4
    }
    v219 := v3[(^(v9))]
    fmt.Println(v1)
    var v220 int16 = (v81 ^ int16(v83))
    v5 = append(v5, "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v219:", v219)
    fmt.Println("v220:", v220)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v219
    _ = v220
    _ = v3
    _ = v4
    _ = v5
    _ = v81
    _ = v82
    _ = v83
    _ = v9
}
