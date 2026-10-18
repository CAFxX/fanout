package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint16; F1 int64; F2 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32; F1 uint16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 bool; F1 int8; F2 uint8 }
type GenS8 struct { F0 int16; F1 uint16 }
type GenS9 struct { F0 int64; F1 uint64 }
func genFn0() int32 {
    var v7 int64 = 0
    fmt.Println(v7)
    _ = v7
    return 1
}

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 42
    v3 := make(map[string]uint32)
    v4 := make(map[string]bool)
    v5 := make(map[uint32]uint32)
    var v6 uint64 = 1
    for fi8 := 0; fi8 < 4; fi8++ {
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 int16 = 42
            var v11 uint8 = 5
            v12 := GenS1{F0: fi8, F1: 32767}
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        var v13 int64 = -9223372036854775808
        _ = fi8
        _ = v13
    }
    var v14 bool = true
    v14 = true
    v14 = v14
    if false {
        v2 = (v2 & uint32(v6))
        if true {
            v6 = v6
            if v14 {
                switch "hello" {
                case "":
                    for fi15 := 0; fi15 < 14; fi15++ {
                        var v16 int64 = 644
                        v2 = (^(v2))
                        _ = fi15
                        _ = v16
                    }
                }
            } else {
                fmt.Println(v2)
                var v17 float64 = (float64((0.1 * float32((0.0 + (float64((1e10 - -2.25)) - (((0.0 / 0.1) - (3.14159 * (0.0 + 3.14159))) / (-2.25 / float64(1e10)))))))) / 3.14159)
                _ = v17
            }
        } else {
            var v18 uint32 = (v2 + v2)
            v19 := GenS2{F0: 0, F1: 0, F2: v18}
            _ = v18
            _ = v19
        }
        var v20 int64 = 1
        _ = v20
    } else {
        v2 = v2
        switch -287 {
        case -885:
            if (true || false) {
                v21 := GenS3{F0: 127}
                _ = v21
            } else {
                switch int(v2) {
                case -414:
                    fmt.Println(v6, v14, v6)
                case -1:
                    var v22 int16 = 32767
                    _ = v22
                }
                v2 += (v2 ^ 2)
                var v23 float32 = 1e10
                _ = v23
            }
        default:
            fmt.Println(v2)
        }
        if false {
            switch 42 {
            case -9223372036854775808:
                switch "foo bar" {
                case "":
                    switch (("" + "") + "foo bar") {
                    case "\n\t":
                        switch (("hello" + "\n\t") + "\n\t") {
                        case "foo bar":
                            if v14 {
                                var v24 float32 = -2.25
                                switch 185 {
                                case 32767:
                                    if v14 {
                                        var v25 bool = ((0 <= v6) || (!false))
                                        if false {
                                            v26 := GenS2{F0: 65535, F1: 1, F2: v2}
                                            switch 32767 {
                                            case 1:
                                                v27 := GenS4{F0: -673}
                                                _ = v27
                                            case -9223372036854775808:
                                                fmt.Println(v14, v2)
                                            }
                                            _ = v26
                                        } else {
                                            var v28 bool = v25
                                            _ = v28
                                        }
                                        _ = v25
                                    }
                                default:
                                    var v29 uint = 2
                                    _ = v29
                                }
                                _ = len(v4)
                                _ = v24
                            } else {
                                if false {
                                    for fi30 := 0; fi30 < 7; fi30++ {
                                        for fi31 := 0; fi31 < 9; fi31++ {
                                            v6 = 669
                                            v2 -= (v2 * (v2 >> ((18446744073709551615) & 63)))
                                            _ = fi31
                                        }
                                        var v32 uint = 0
                                        var v33 float64 = float64((3.14159 - 3.14159))
                                        _ = fi30
                                        _ = v32
                                        _ = v33
                                    }
                                    switch "\n\t" {
                                    case "hello":
                                        if (124 <= 100) {
                                            var v34 uint = 1
                                            v34 = 0
                                            _ = v34
                                        }
                                    case "":
                                        var v35 int16 = -114
                                        _ = v35
                                    case "\n\t":
                                        var v36 uint16 = 0
                                        _ = v36
                                    default:
                                        if v14 {
                                            if (0.1 == 0.0) {
                                                fmt.Println(v14)
                                                for fi37 := 0; fi37 < 17; fi37++ {
                                                    for fi38 := 0; fi38 < 2; fi38++ {
                                                        switch "a" {
                                                        case "hello":
                                                            v14 = ((!v14) && v14)
                                                            if (2 != 943) {
                                                                var v39 int64 = 0
                                                                if (uint8(v2) < 42) {
                                                                    var v40 int8 = -9
                                                                    if v14 {
                                                                        v14 = (!v14)
                                                                    }
                                                                    var v41 float32 = 1.5
                                                                    _ = v40
                                                                    _ = v41
                                                                } else {
                                                                    v42 := GenS1{F0: fi37, F1: 42}
                                                                    if v14 {
                                                                        fmt.Println(fi37, fi37, v2)
                                                                    }
                                                                    _ = v42
                                                                }
                                                                _ = v39
                                                            }
                                                        case "":
                                                            var v43 int64 = -9223372036854775808
                                                            _ = v43
                                                        case "a":
                                                            fmt.Println(v2)
                                                        default:
                                                            switch "foo bar" {
                                                            case "hello":
                                                                var v44 int32 = genFn0()
                                                                _ = v44
                                                            case "a":
                                                                fmt.Println(v2, fi38)
                                                            case "foo bar":
                                                                v45 := []int{fi37}
                                                                _ = v45
                                                            default:
                                                                var v46 float32 = 1e10
                                                                _ = v46
                                                            }
                                                            var v47 string = ""
                                                            var v48 int64 = -9223372036854775808
                                                            var v49 uint = 2
                                                            if (!v14) {
                                                                if v14 {
                                                                    switch v47 {
                                                                    case "":
                                                                        v50 := GenS2{F0: 1, F1: 0, F2: (v2 % ((578) | 1))}
                                                                        _ = v50
                                                                    case "a":
                                                                        if v14 {
                                                                            var v51 uint8 = 2
                                                                            _ = v51
                                                                        }
                                                                    default:
                                                                        if v14 {
                                                                            v6 = (v6 & v6)
                                                                        } else {
                                                                            v49 += 290
                                                                            fmt.Println(v47)
                                                                            fmt.Println(v14)
                                                                            var v52 uint64 = v6
                                                                            var v53 float64 = 1e10
                                                                            _ = v52
                                                                            _ = v53
                                                                        }
                                                                    }
                                                                }
                                                                switch (^(fi38)) {
                                                                case 0:
                                                                    var v54 uint16 = 65535
                                                                    _ = v54
                                                                case 127:
                                                                    var v55 uint32 = uint32(v49)
                                                                    _ = v55
                                                                default:
                                                                    var v56 float64 = (1.5 + 0.1)
                                                                    _ = v56
                                                                }
                                                            }
                                                            v14 = ((!((fi38 % (((fi38 | fi38)) | 1)) <= (fi38 << (((^(v49))) & 63)))) || (true || v14))
                                                            _ = v47
                                                            _ = v48
                                                            _ = v49
                                                        }
                                                        v57 := []int{(^(fi37)), 127, (fi38 + 463)}
                                                        v58 := GenS5{F0: -2147483648, F1: 42}
                                                        _ = fi38
                                                        _ = v57
                                                        _ = v58
                                                    }
                                                    switch 153 {
                                                    case 42:
                                                        for fi59 := 0; fi59 < 6; fi59++ {
                                                            fmt.Println(v6)
                                                            if v14 {
                                                                _ = len(v5)
                                                                var v60 int64 = 813
                                                                if true {
                                                                    v61 := GenS4{F0: 42}
                                                                    if v14 {
                                                                        var v62 float32 = 3.14159
                                                                        v60 -= (v60 * (v60 << ((162) & 63)))
                                                                        _ = v62
                                                                    }
                                                                    if (fi59 >= 82) {
                                                                        var v63 uint8 = 128
                                                                        _ = v63
                                                                    }
                                                                    _ = v61
                                                                } else {
                                                                    fmt.Println(v6, fi37, v14)
                                                                    var v64 uint16 = 32767
                                                                    v6 += 79
                                                                    _ = v64
                                                                }
                                                                switch 4294967295 {
                                                                case 128:
                                                                    v14 = ((1.5 * 3.14159) == 3.14159)
                                                                case -2147483648:
                                                                    v6 = v6
                                                                default:
                                                                    v14 = v14
                                                                }
                                                                _ = v60
                                                            } else {
                                                                var v65 uint16 = 42
                                                                _ = v65
                                                            }
                                                            _ = fi59
                                                        }
                                                    case 2147483647:
                                                        fmt.Println(v2)
                                                    case 711:
                                                        v66 := []int{-206, fi37, fi37}
                                                        _ = v66
                                                    default:
                                                        v6 = 1
                                                    }
                                                    _ = fi37
                                                }
                                                v14 = v14
                                            } else {
                                                if (924 != 65535) {
                                                    v6 = v6
                                                    v2 = 4294967295
                                                }
                                                if false {
                                                    if (!v14) {
                                                        v67 := GenS5{F0: int32(v2), F1: 42}
                                                        _ = v67
                                                    } else {
                                                        var v68 uint16 = 107
                                                        if v14 {
                                                            var v69 int = 4294967295
                                                            _ = len(v3)
                                                            _ = v69
                                                        } else {
                                                            var v70 int32 = 0
                                                            v2 = 2147483647
                                                            _ = v70
                                                        }
                                                        if (!(!false)) {
                                                            switch ((("hello" + ("hello" + (("hello" + "\n\t") + ("foo bar" + "\n\t")))) + ("a" + "a")) + "foo bar") {
                                                            case "":
                                                                var v71 int64 = 0
                                                                _ = v71
                                                            case "\n\t":
                                                                switch "" {
                                                                case "hello":
                                                                    v14 = (v2 < (^(v2)))
                                                                }
                                                            case "hello":
                                                                for fi72 := 0; fi72 < 8; fi72++ {
                                                                    var v73 int = (fi72 * int(v68))
                                                                    v74 := []int{(v73 - (fi72 | 4294967295)), (fi72 + (fi72 * (fi72 << ((4294967295) & 63))))}
                                                                    v68 = 255
                                                                    _ = fi72
                                                                    _ = v73
                                                                    _ = v74
                                                                }
                                                            default:
                                                                v2 += (v2 >> ((18446744073709551615) & 63))
                                                            }
                                                            var v75 bool = (!v14)
                                                            _ = v75
                                                        }
                                                        _ = v68
                                                    }
                                                }
                                                var v76 int32 = 1
                                                _ = v76
                                            }
                                        } else {
                                            v14 = (float32(((1.5 - 1.5) - 0.1)) == (0.0 * 0.0))
                                        }
                                    }
                                    v2 += v2
                                } else {
                                    var v77 uint = 1
                                    v6 = (v6 * 1)
                                    _ = v77
                                }
                            }
                        default:
                            fmt.Println(v14, v14)
                        }
                    case "":
                        if v14 {
                            if false {
                                var v78 uint16 = 141
                                v14 = false
                                var v79 int16 = 0
                                _ = v78
                                _ = v79
                            } else {
                                v14 = (0.0 != 3.14159)
                                switch "\n\t" {
                                case "\n\t":
                                    for fi80 := 0; fi80 < 17; fi80++ {
                                        var v81 string = "a"
                                        var v82 uint32 = (v2 + v2)
                                        _ = fi80
                                        _ = v81
                                        _ = v82
                                    }
                                case "":
                                    fmt.Println(v2, v2, v6)
                                case "a":
                                    v83 := GenS2{F0: 255, F1: 0, F2: (^(v2))}
                                    _ = v83
                                }
                                v84 := []int{1}
                                _ = v84
                            }
                        } else {
                            v6 -= (^(v6))
                            v85 := []int{32768, 128, int(v2)}
                            for fi86 := 0; fi86 < 18; fi86++ {
                                v14 = false
                                v14 = v14
                                v87 := GenS4{F0: -32768}
                                _ = fi86
                                _ = v87
                            }
                            _ = v85
                        }
                    }
                default:
                    for fi88 := 0; fi88 < 2; fi88++ {
                        fmt.Println(v6, v14)
                        if (((^(v6)) < 1) && v14) {
                            v89 := []int{(fi88 - (fi88 | (-(fi88)))), (fi88 >> ((18446744073709551615) & 63)), 9223372036854775807, fi88}
                            v90 := GenS4{F0: -639}
                            _ = v89
                            _ = v90
                        } else {
                            fmt.Println(v14, v2)
                            v91 := GenS6{F0: -9223372036854775808}
                            _ = v91
                        }
                        _ = fi88
                    }
                }
            case -372:
                var v92 bool = v14
                _ = v92
            case -2147483648:
                var v93 int16 = 0
                _ = v93
            }
            switch "hello" {
            case "a":
                switch 359 {
                case -1:
                    var v94 string = "foo bar"
                    _ = v94
                case -579:
                    fmt.Println(v2, v2, v6)
                }
            case "\n\t":
                v14 = (false || v14)
            case "foo bar":
                v95 := []int{}
                _ = v95
            default:
                fmt.Println(v6, v2)
            }
            var v96 float64 = (float64((float32((0.0 / 1.0)) - 1e10)) + -2.25)
            _ = v96
        } else {
            v2 = (v2 << ((771) & 63))
        }
    }
    if v14 {
        for fi97 := 0; fi97 < 17; fi97++ {
            v6 -= v6
            for fi98 := 0; fi98 < 4; fi98++ {
                if false {
                    var v99 int = (^(fi98))
                    v99 += (fi97 - 1)
                    var v100 int8 = int8(fi98)
                    v101 := GenS3{F0: v100}
                    _ = v100
                    _ = v101
                    _ = v99
                } else {
                    var v102 uint = 0
                    _ = v102
                }
                _ = fi98
            }
            _ = fi97
        }
        if v14 {
            if v14 {
                delete(v1, 668)
                var v103 int16 = 3
                switch -2147483648 {
                case -1:
                    v103 = -32768
                case 451:
                    if v14 {
                        v104 := GenS7{F0: (!((^(v2)) != v2)), F1: 100, F2: 255}
                        for fi105 := 0; fi105 < 19; fi105++ {
                            v106 := GenS8{F0: (v103 * v103), F1: 156}
                            switch fi105 {
                            case 32767:
                                if false {
                                    v107 := GenS5{F0: genFn0(), F1: 1}
                                    _ = v107
                                } else {
                                    fmt.Println(v2)
                                    v108 := GenS5{F0: 0, F1: 65535}
                                    _ = v108
                                }
                            case -1:
                                if ((^(v2)) > (v2 % ((0) | 1))) {
                                    var v109 int64 = 1
                                    _ = v109
                                }
                            case 1:
                                if (2 > 88) {
                                    v6 = v6
                                    var v110 uint64 = 882
                                    _ = v110
                                }
                            }
                            _ = fi105
                            _ = v106
                        }
                        _ = v104
                    }
                }
                _ = v103
            } else {
                fmt.Println(v6)
                var v111 uint = uint(v2)
                if false {
                    fmt.Println(v6, v111, v14)
                    v112 := GenS1{F0: -793, F1: -32768}
                    _ = v112
                } else {
                    for fi113 := 0; fi113 < 13; fi113++ {
                        var v114 uint16 = 892
                        v115 := GenS8{F0: 0, F1: uint16(v6)}
                        if (v14 && v14) {
                            if v14 {
                                switch (("foo bar" + "") + "") {
                                case "hello":
                                    v2 = 65535
                                case "":
                                    var v116 uint8 = 128
                                    _ = v116
                                }
                                var v117 uint = v111
                                var v118 int16 = 541
                                _ = v117
                                _ = v118
                            }
                        }
                        _ = fi113
                        _ = v114
                        _ = v115
                    }
                    fmt.Println(v2)
                }
                _ = v111
            }
            var v119 uint = 42
            if (!(255 > 372)) {
                fmt.Println(v2, v6, v119)
            } else {
                fmt.Println(v14, v14)
                switch 32768 {
                case -448:
                    if (((0.1 - (0.0 * (3.14159 + float32((0.0 + 0.0))))) - 3.14159) > ((1e10 + -2.25) / 1.0)) {
                        var v120 uint = v119
                        v121 := GenS7{F0: (false && (true || (649 >= 42))), F1: -34, F2: 1}
                        _ = v120
                        _ = v121
                    } else {
                        v122 := []int{}
                        if true {
                            switch -3 {
                            case 42:
                                if ("" <= "hello") {
                                    var v123 int64 = 42
                                    fmt.Println(v14, v6, v2)
                                    v124 := GenS9{F0: -9223372036854775808, F1: v6}
                                    _ = v123
                                    _ = v124
                                } else {
                                    _ = len(v122)
                                    fmt.Println(v14, v119, v2)
                                }
                            case 1:
                                if len(v122) >= 2 {
                                    v122 = v122[1:len(v122)]
                                }
                            default:
                                fmt.Println(v119)
                            }
                        }
                        v119 = (v119 & (v119 % (((v119 | 1)) | 1)))
                        _ = v122
                    }
                case -439:
                    v125 := GenS7{F0: (-283 > 2147483647), F1: 42, F2: 2}
                    _ = v125
                default:
                    if (v14 && (("foo bar" + "\n\t") != ("" + "a"))) {
                        if (true || (true && true)) {
                            v126 := v4[(("" + ("\n\t" + "\n\t")) + ("" + "hello"))]
                            fmt.Println(v2)
                            v6 = (v6 + v6)
                            _ = v126
                        }
                    }
                }
                for fi127 := 0; fi127 < 5; fi127++ {
                    if (886 >= -327) {
                        if v14 {
                            var v128 int32 = genFn0()
                            v6 -= 540
                            switch 127 {
                            case 4294967295:
                                var v129 uint8 = 191
                                _ = v129
                            case -497:
                                var v130 int64 = 884
                                _ = v130
                            }
                            _ = v128
                        }
                        var v131 float64 = 0.1
                        v132 := GenS9{F0: 42, F1: (v6 >> ((v119) & 63))}
                        _ = v131
                        _ = v132
                    } else {
                        if (!(1e10 > 0.1)) {
                            if ((v119 + 0) >= v119) {
                                v119 += 326
                                switch ("hello" + ("" + "")) {
                                case "a":
                                    var v133 uint64 = (v6 >> (((v119 & (v119 ^ (v119 % (((v119 + (v119 % ((v119) | 1)))) | 1))))) & 63))
                                    _ = v133
                                default:
                                    var v134 int8 = 127
                                    _ = v134
                                }
                            }
                            v135 := GenS4{F0: 973}
                            for fi136 := 0; fi136 < 12; fi136++ {
                                if v14 {
                                    if (0 <= 0) {
                                        v14 = v14
                                        var v137 int = 32767
                                        v138 := GenS8{F0: 0, F1: 65535}
                                        _ = v137
                                        _ = v138
                                    }
                                    fmt.Println(v6, fi136, fi136)
                                }
                                v119 += v119
                                switch ("foo bar" + ("foo bar" + "")) {
                                case "hello":
                                    var v139 bool = false
                                    _ = v139
                                case "a":
                                    if v14 {
                                        var v140 uint64 = v6
                                        var v141 int32 = genFn0()
                                        var v142 int32 = v141
                                        v140 = v6
                                        _ = v140
                                        _ = v141
                                        _ = v142
                                    } else {
                                        v3["\n\t"] = v2
                                        v143 := GenS2{F0: 260, F1: 591, F2: v2}
                                        v6 -= v6
                                        _ = v143
                                    }
                                default:
                                    var v144 int = (fi127 ^ fi127)
                                    _ = v144
                                }
                                _ = fi136
                            }
                            _ = v135
                        } else {
                            fmt.Println(v119, v14)
                            var v145 int64 = -339
                            _ = v145
                        }
                        for fi146 := 0; fi146 < 9; fi146++ {
                            v147 := GenS2{F0: 969, F1: 387, F2: (^(v2))}
                            if (65535 == (v119 >> ((v119) & 63))) {
                                if true {
                                    v148 := []int{fi127}
                                    _ = v148
                                } else {
                                    var v149 uint8 = uint8(v6)
                                    if v14 {
                                        fmt.Println(v14)
                                        var v150 int8 = 127
                                        var v151 float64 = 3.14159
                                        _ = v150
                                        _ = v151
                                    }
                                    var v152 int64 = -160
                                    _ = v149
                                    _ = v152
                                }
                                var v153 int16 = 42
                                var v154 int64 = -9223372036854775808
                                _ = v153
                                _ = v154
                            } else {
                                v119 += v119
                                var v155 uint64 = 269
                                var v156 int32 = genFn0()
                                if (v119 <= 0) {
                                    var v157 bool = (v119 <= v119)
                                    var v158 uint32 = 711
                                    v6 = v6
                                    _ = v157
                                    _ = v158
                                }
                                var v159 string = "a"
                                v2 -= uint32(v155)
                                var v160 float32 = (1.5 + (0.1 * (((-2.25 * 3.14159) / 0.1) / -2.25)))
                                _ = v155
                                _ = v156
                                _ = v159
                                _ = v160
                            }
                            fmt.Println(v6, fi146, v119)
                            _ = fi146
                            _ = v147
                        }
                        for fi161 := 0; fi161 < 13; fi161++ {
                            var v162 uint8 = 0
                            v14 = (9223372036854775807 >= 0)
                            v162 = 30
                            _ = fi161
                            _ = v162
                        }
                    }
                    _ = fi127
                }
            }
            _ = v119
        }
    }
    fmt.Println(v14, v14, v6)
    v163 := GenS2{F0: 255, F1: 0, F2: v2}
    if false {
        if false {
            v2 = (v2 + v2)
        } else {
            fmt.Println(v14, v14, v6)
        }
        fmt.Println(v14, v2)
        v164 := GenS4{F0: 32767}
        _ = v164
    } else {
        for fi165 := 0; fi165 < 2; fi165++ {
            var v166 int = 0
            var v167 int64 = 331
            for fi168 := 0; fi168 < 11; fi168++ {
                if v14 {
                    var v169 uint8 = 255
                    v170 := v3["hello"]
                    v171 := []int{}
                    _ = v169
                    _ = v170
                    _ = v171
                }
                _ = fi168
            }
            _ = fi165
            _ = v166
            _ = v167
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v163
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
