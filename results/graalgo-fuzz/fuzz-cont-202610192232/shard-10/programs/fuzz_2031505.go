package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int32 }
type GenS2 struct { F0 int16; F1 int64; F2 uint32 }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 int32; F1 bool; F2 bool }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 string; F1 string }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint64; F1 string; F2 int }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int16; F1 uint; F2 int8 }
type GenS12 struct { F0 uint; F1 uint; F2 int32 }
type GenS13 struct { F0 uint8; F1 int32 }
type GenS14 struct { F0 uint8; F1 uint8 }

func main() {
    var v1 string = ("" + "hello")
    var v2 uint16 = 255
    v3 := make(map[int]bool)
    var v4 int8 = 127
    for fi5 := 0; fi5 < 15; fi5++ {
        for fi6 := 0; fi6 < 13; fi6++ {
            v2 -= v2
            switch v1 {
            case "a":
                var v7 int = fi5
                var v8 int64 = 0
                _ = v7
                _ = v8
            case "hello":
                var v9 uint = 255
                _ = v9
            }
            _ = fi6
        }
        for fi10 := 0; fi10 < 12; fi10++ {
            if (v4 >= v4) {
                if (!(!(!false))) {
                    switch ("foo bar" + "foo bar") {
                    case "foo bar":
                        var v11 uint16 = v2
                        _ = v11
                    case "a":
                        v1 = v1
                    case "\n\t":
                        var v12 int64 = -9223372036854775808
                        _ = v12
                    }
                    v3[(fi10 ^ (fi5 & (fi10 & -783)))] = (1 != 1)
                } else {
                    fmt.Println(v4)
                    var v13 float64 = 3.14159
                    v1 = v1
                    _ = v13
                }
                if false {
                    fmt.Println(fi5, v1, v1)
                    v2 += v2
                    v4 = 51
                }
            } else {
                var v14 uint32 = 0
                if true {
                    var v15 uint16 = v2
                    v14 = (v14 % (((v14 ^ 2)) | 1))
                    v16 := GenS1{F0: 32767, F1: -2147483648, F2: 2147483647}
                    _ = v15
                    _ = v16
                } else {
                    var v17 bool = false
                    if (-1 != 235) {
                        var v18 uint16 = 0
                        if true {
                            var v19 float64 = ((-2.25 - 1e10) / float64((((1e10 + (0.0 / 1.0)) + 3.14159) / 1.5)))
                            switch -639 {
                            case 127:
                                var v20 uint16 = (v2 << ((255) & 63))
                                _ = v20
                            default:
                                v2 = v18
                            }
                            switch v1 {
                            case "hello":
                                v21 := GenS2{F0: -1, F1: 1, F2: v14}
                                _ = v21
                            case "":
                                if v17 {
                                    var v22 string = v1
                                    var v23 bool = (v17 && (v4 > (v4 + (v4 + v4))))
                                    _ = v22
                                    _ = v23
                                }
                            }
                            v2 += (^(v18))
                            _ = v19
                        } else {
                            v24 := []int{}
                            v25 := GenS1{F0: (^(v2)), F1: 2147483647, F2: -2147483648}
                            var v26 uint8 = 255
                            if false {
                                v1 = v1
                                _ = len(v24)
                            } else {
                                var v27 int = (fi5 << ((1) & 63))
                                _ = v27
                            }
                            if (!(!v17)) {
                                v28 := GenS3{F0: "\n\t", F1: fi10}
                                _ = len(v24)
                                v29 := GenS3{F0: v1, F1: 128}
                                _ = v28
                                _ = v29
                            }
                            var v30 uint = 2
                            var v31 float32 = 3.14159
                            _ = v24
                            _ = v25
                            _ = v26
                            _ = v30
                            _ = v31
                        }
                        if v17 {
                            v4 += (v4 ^ v4)
                            v2 += (^(v2))
                        } else {
                            var v32 uint32 = v14
                            v3[-57] = (65535 <= 1)
                            switch v1 {
                            case "\n\t":
                                var v33 int8 = v4
                                _ = v33
                            case "hello":
                                switch 42 {
                                case 4294967295:
                                    switch v1 {
                                    case "":
                                        fmt.Println(v2)
                                    case "\n\t":
                                        if v17 {
                                            var v34 float64 = 3.14159
                                            _ = len(v3)
                                            _ = v34
                                        }
                                    case "a":
                                        var v35 uint8 = 42
                                        _ = v35
                                    }
                                    var v36 uint8 = 1
                                    v3[fi10] = ((v4 ^ 42) != v4)
                                    var v37 uint8 = v36
                                    _ = v36
                                    _ = v37
                                case 9223372036854775807:
                                    v38 := GenS4{F0: 2, F1: -1}
                                    _ = v38
                                default:
                                    switch "a" {
                                    case "\n\t":
                                        fmt.Println(v1)
                                    case "foo bar":
                                        v18 = (v2 >> ((42) & 63))
                                    case "":
                                        if v17 {
                                            v2 += 42
                                            var v39 float64 = 1e10
                                            v40 := GenS3{F0: (v1 + v1), F1: (fi5 >> ((42) & 63))}
                                            _ = v39
                                            _ = v40
                                        } else {
                                            v41 := []int{0}
                                            if v17 {
                                                v3[(-(fi10))] = true
                                                if ((fi10 <= fi5) || ((0 >= 18446744073709551615) && ((1e10 - 0.1) != float32((0.1 - 1e10))))) {
                                                    var v42 int64 = 14
                                                    _ = v42
                                                }
                                                switch v1 {
                                                case "foo bar":
                                                    fmt.Println(v17)
                                                case "":
                                                    var v43 int16 = 0
                                                    _ = v43
                                                case "a":
                                                    _ = len(v41)
                                                }
                                            }
                                            if v17 {
                                                var v44 int16 = -402
                                                _ = v44
                                            } else {
                                                v2 -= (v18 >> ((18446744073709551615) & 63))
                                                fmt.Println(fi5)
                                                var v45 uint = 431
                                                _ = v45
                                            }
                                            var v46 int16 = -772
                                            _ = len(v3)
                                            _ = v41
                                            _ = v46
                                        }
                                    default:
                                        if false {
                                            var v47 int = 127
                                            v48 := GenS4{F0: 2, F1: 336}
                                            _ = v47
                                            _ = v48
                                        } else {
                                            fmt.Println(fi10)
                                            if (1e10 > float64((1e10 * (0.0 + 1.5)))) {
                                                var v49 uint8 = 1
                                                switch v1 {
                                                case "foo bar":
                                                    switch "hello" {
                                                    case "a":
                                                        var v50 int = fi10
                                                        _ = v50
                                                    case "foo bar":
                                                        v1 = ("\n\t" + (v1 + v1))
                                                    case "hello":
                                                        var v51 int32 = 42
                                                        _ = v51
                                                    }
                                                    var v52 bool = v17
                                                    switch v1 {
                                                    case "foo bar":
                                                        v4 = (v4 << ((2) & 63))
                                                    case "\n\t":
                                                        v4 = (v4 & (v4 * v4))
                                                    case "hello":
                                                        _ = len(v3)
                                                    }
                                                    _ = v52
                                                }
                                                var v53 uint32 = (v32 + (v32 | 2147483647))
                                                _ = v49
                                                _ = v53
                                            } else {
                                                var v54 uint64 = 2
                                                v55 := []int{fi10, 32767, fi10}
                                                v32 += v32
                                                _ = v54
                                                _ = v55
                                            }
                                            var v56 int32 = -411
                                            _ = v56
                                        }
                                    }
                                    v57 := GenS5{F0: -2147483648, F1: v17, F2: ((v17 && v17) || ((^(v18)) > 2))}
                                    var v58 int64 = 42
                                    var v59 uint8 = 202
                                    switch 1 {
                                    case 9223372036854775807:
                                        v60 := GenS6{F0: v4}
                                        _ = v60
                                    case 0:
                                        switch v1 {
                                        case "foo bar":
                                            fmt.Println(v14, v58)
                                            v61 := []int{(fi5 | 32767), (fi5 ^ (fi5 * fi10))}
                                            _ = v61
                                        case "":
                                            var v62 uint = 42
                                            _ = v62
                                        case "a":
                                            v17 = ((v14 % ((v14) | 1)) <= uint32(v2))
                                        }
                                    }
                                    _ = v57
                                    _ = v58
                                    _ = v59
                                }
                            case "":
                                v63 := []int{int(v18), (-(fi10)), (fi10 << ((674) & 63)), 9223372036854775807}
                                _ = v63
                            }
                            _ = v32
                        }
                        _ = v18
                    }
                    v64 := []int{32768, (fi5 * fi10), int(v14)}
                    _ = v17
                    _ = v64
                }
                _ = v14
            }
            _ = fi10
        }
        if true {
            fmt.Println(v4, v2, v4)
            switch (fi5 + (fi5 + (fi5 >> ((582) & 63)))) {
            case -54:
                v65 := []int{(fi5 >> ((956) & 63))}
                _ = v65
            case 9223372036854775807:
                var v66 string = v1
                _ = v66
            case 0:
                var v67 int32 = -2147483648
                _ = v67
            }
            var v68 int64 = 9223372036854775807
            _ = v68
        }
        _ = fi5
    }
    v1 = v1
    var v69 int = 2147483647
    v69 = v69
    if (!(!true)) {
        for fi70 := 0; fi70 < 2; fi70++ {
            v71 := GenS7{F0: ("hello" + v1), F1: "\n\t"}
            fmt.Println(v1)
            v69 += (fi70 & fi70)
            _ = fi70
            _ = v71
        }
    } else {
        fmt.Println(v2, v2, v4)
        for fi72 := 0; fi72 < 7; fi72++ {
            var v73 int64 = -913
            for fi74 := 0; fi74 < 18; fi74++ {
                if true {
                    v75 := []int{}
                    _ = v75
                }
                switch 0 {
                case 1:
                    var v76 int = fi74
                    _ = v76
                default:
                    v77 := []int{}
                    _ = v77
                }
                v1 = (((v1 + ("foo bar" + ("\n\t" + v1))) + "") + ("" + v1))
                _ = fi74
            }
            var v78 float64 = -2.25
            _ = fi72
            _ = v73
            _ = v78
        }
    }
    var v79 int = v69
    if false {
        v80 := GenS4{F0: 128, F1: 42}
        var v81 uint16 = v2
        _ = v80
        _ = v81
    } else {
        var v82 int = -93
        switch (v82 + 127) {
        case -1:
            var v83 float64 = (0.1 - 0.1)
            _ = v83
        case 145:
            if false {
                switch v69 {
                case 944:
                    var v84 int64 = 1
                    _ = v84
                case 42:
                    v4 = v4
                case 128:
                    v69 = v82
                }
                if (false && ((v1 <= "foo bar") && (1 > 834))) {
                    for fi85 := 0; fi85 < 18; fi85++ {
                        v4 += 1
                        switch v1 {
                        case "":
                            if (!(false || (((0 == 42) || (9223372036854775807 == -1)) && ((!true) || (false || false))))) {
                                var v86 uint32 = 42
                                _ = v86
                            } else {
                                v69 += (v82 ^ (fi85 % ((fi85) | 1)))
                            }
                        case "a":
                            if true {
                                var v87 uint32 = 577
                                _ = v87
                            } else {
                                fmt.Println(v79, v69, v2)
                                v79 += 885
                                var v88 uint = 0
                                _ = v88
                            }
                        case "hello":
                            v79 = v69
                        default:
                            switch "foo bar" {
                            case "":
                                for fi89 := 0; fi89 < 9; fi89++ {
                                    var v90 uint = 290
                                    _ = fi89
                                    _ = v90
                                }
                            case "foo bar":
                                for fi91 := 0; fi91 < 5; fi91++ {
                                    v92 := []int{(v79 << ((255) & 63)), (-(v82)), 994, -389}
                                    v4 = 1
                                    _ = fi91
                                    _ = v92
                                }
                            case "hello":
                                v3[v69] = (fi85 == 32768)
                            default:
                                for fi93 := 0; fi93 < 9; fi93++ {
                                    switch -1 {
                                    case 9223372036854775807:
                                        fmt.Println(v79, fi85)
                                    default:
                                        v82 = fi85
                                    }
                                    if (65535 > 351) {
                                        fmt.Println(v4, v1)
                                    } else {
                                        v94 := GenS8{F0: 0}
                                        if (2 == 200) {
                                            v95 := GenS9{F0: 254, F1: "\n\t", F2: (v82 << ((42) & 63))}
                                            var v96 string = v1
                                            _ = v95
                                            _ = v96
                                        } else {
                                            v97 := GenS2{F0: 1, F1: 9223372036854775807, F2: 217}
                                            var v98 int = (v79 / ((195) | 1))
                                            if true {
                                                v99 := GenS10{F0: ((v1 + v1) + "hello")}
                                                v69 -= 2147483647
                                                switch 127 {
                                                case 32767:
                                                    v100 := GenS5{F0: -271, F1: (!false), F2: (fi85 == 32767)}
                                                    _ = v100
                                                case -268:
                                                    fmt.Println(fi93, v79)
                                                }
                                                v82 -= fi93
                                                var v101 bool = (1 >= 32767)
                                                v102 := v3[(v82 + 9223372036854775807)]
                                                _ = v101
                                                _ = v102
                                                _ = v99
                                            }
                                            _ = v97
                                            _ = v98
                                        }
                                        v103 := GenS8{F0: 255}
                                        _ = v103
                                        _ = v94
                                    }
                                    _ = fi93
                                }
                            }
                        }
                        _ = fi85
                    }
                    if false {
                        var v104 int16 = 981
                        _ = v104
                    }
                    _ = len(v3)
                }
                var v105 uint64 = 18446744073709551615
                _ = v105
            } else {
                if (!(!false)) {
                    if true {
                        for fi106 := 0; fi106 < 14; fi106++ {
                            v4 = 0
                            _ = fi106
                        }
                        for fi107 := 0; fi107 < 12; fi107++ {
                            for fi108 := 0; fi108 < 15; fi108++ {
                                fmt.Println(v79, v82)
                                if false {
                                    v109 := GenS8{F0: 231}
                                    delete(v3, fi107)
                                    v1 = v1
                                    _ = v109
                                } else {
                                    v110 := GenS3{F0: v1, F1: v79}
                                    var v111 int16 = -495
                                    _ = v110
                                    _ = v111
                                }
                                _ = fi108
                            }
                            v79 += (v69 & (v69 & -2147483648))
                            for fi112 := 0; fi112 < 15; fi112++ {
                                fmt.Println(v1, v2, v79)
                                switch (v82 + 42) {
                                case 673:
                                    fmt.Println(fi112, fi112, fi107)
                                case 42:
                                    v113 := GenS9{F0: 42, F1: ("" + "\n\t"), F2: (^(fi112))}
                                    _ = v113
                                case 127:
                                    fmt.Println(fi112, fi112)
                                default:
                                    var v114 uint64 = 18446744073709551615
                                    _ = v114
                                }
                                _ = len(v3)
                                _ = fi112
                            }
                            _ = fi107
                        }
                        var v115 string = (("a" + v1) + (("foo bar" + ("\n\t" + "foo bar")) + ((((v1 + (v1 + ("a" + (("a" + (v1 + (v1 + v1))) + "hello")))) + "") + ("foo bar" + "hello")) + (v1 + "hello"))))
                        _ = v115
                    } else {
                        if (false || (!false)) {
                            var v116 int32 = 0
                            v1 = ""
                            _ = v116
                        } else {
                            v117 := GenS9{F0: 0, F1: ("a" + (v1 + (v1 + ""))), F2: 42}
                            for fi118 := 0; fi118 < 17; fi118++ {
                                if true {
                                    v2 = 1
                                    v82 += (fi118 / ((v82) | 1))
                                }
                                _ = fi118
                            }
                            _ = v117
                        }
                        for fi119 := 0; fi119 < 17; fi119++ {
                            if (!false) {
                                v120 := []int{v69, (v82 % ((32768) | 1))}
                                v82 = (v69 * 9223372036854775807)
                                _ = v120
                            }
                            var v121 float32 = 1.5
                            _ = fi119
                            _ = v121
                        }
                        v122 := GenS8{F0: 1}
                        _ = v122
                    }
                    v79 = (v69 >> ((0) & 63))
                    if false {
                        fmt.Println(v79, v4)
                        if (true || (0 > 42)) {
                            var v123 int8 = v4
                            for fi124 := 0; fi124 < 9; fi124++ {
                                v69 += 741
                                v125 := GenS11{F0: -1, F1: 0, F2: (v4 << ((18446744073709551615) & 63))}
                                _ = fi124
                                _ = v125
                            }
                            _ = v123
                        }
                    }
                } else {
                    v126 := GenS12{F0: 255, F1: 4294967295, F2: 42}
                    _ = v126
                }
            }
        default:
            var v127 uint = 328
            _ = v127
        }
        switch "hello" {
        case "a":
            for fi128 := 0; fi128 < 11; fi128++ {
                fmt.Println(v69)
                if false {
                    v1 = v1
                    fmt.Println(v82, v2)
                    var v129 uint16 = v2
                    _ = v129
                } else {
                    fmt.Println(v1)
                }
                _ = fi128
            }
        }
        _ = v82
    }
    v130 := GenS12{F0: 717, F1: 2, F2: 97}
    var v131 float64 = 0.1
    _ = len(v3)
    v4 -= v4
    var v132 int16 = -1
    for fi133 := 0; fi133 < 17; fi133++ {
        var v134 float32 = float32(((v131 * (1.5 - (0.0 * (v131 - (0.1 + 1.5))))) * v131))
        if false {
            if (1 == 9223372036854775807) {
                if (628 <= 9223372036854775807) {
                    switch v1 {
                    case "hello":
                        var v135 float64 = (v131 * 1e10)
                        _ = v135
                    default:
                        var v136 uint64 = 4294967295
                        _ = v136
                    }
                    for fi137 := 0; fi137 < 10; fi137++ {
                        v138 := []int{(v69 | (v69 - fi137)), (fi133 & fi133), (fi137 * (^(fi133)))}
                        var v139 uint16 = v2
                        _ = fi137
                        _ = v138
                        _ = v139
                    }
                }
                for fi140 := 0; fi140 < 20; fi140++ {
                    if true {
                        var v141 uint32 = 42
                        _ = v141
                    } else {
                        if true {
                            v142 := GenS4{F0: 128, F1: -2147483648}
                            v1 = v1
                            if (1 > (fi140 * v69)) {
                                if true {
                                    var v143 int8 = 127
                                    v144 := GenS7{F0: (("\n\t" + v1) + "\n\t"), F1: (v1 + v1)}
                                    _ = v143
                                    _ = v144
                                } else {
                                    var v145 string = v1
                                    _ = v145
                                }
                                fmt.Println(v4, v131, v132)
                                if false {
                                    if false {
                                        v79 += (^(v79))
                                        if true {
                                            v3[v79] = (((!((0.1 < 3.14159) && false)) && true) || (false && ((v79 >= 128) || ((false || ((v131 <= v131) && true)) || false))))
                                            v146 := []int{v69, (fi133 + v69)}
                                            var v147 int8 = v4
                                            _ = v146
                                            _ = v147
                                        }
                                        var v148 uint = 857
                                        _ = v148
                                    } else {
                                        if (65535 != 255) {
                                            var v149 float32 = 0.1
                                            _ = v149
                                        }
                                        var v150 int = 708
                                        if (v134 > 1e10) {
                                            v131 += v131
                                        }
                                        _ = v150
                                    }
                                    var v151 int = 1
                                    _ = v151
                                }
                            }
                            _ = v142
                        } else {
                            v152 := GenS13{F0: 42, F1: -2147483648}
                            delete(v3, (fi140 | (fi133 % ((v69) | 1))))
                            _ = v152
                        }
                    }
                    _ = fi140
                }
            }
            v2 = 0
        } else {
            v153 := GenS14{F0: 1, F1: 127}
            switch (v1 + ("" + v1)) {
            case "a":
                for fi154 := 0; fi154 < 11; fi154++ {
                    v79 += (-(fi154))
                    _ = fi154
                }
            case "":
                v132 = (v132 >> ((42) & 63))
            case "hello":
                fmt.Println(v132)
            }
            _ = v153
        }
        _ = fi133
        _ = v134
    }
    var v155 uint16 = v2
    var v156 int64 = 42
    var v157 int = (v79 | (v69 | v79))
    fmt.Println("v1:", v1)
    fmt.Println("v131:", v131)
    fmt.Println("v132:", v132)
    fmt.Println("v155:", v155)
    fmt.Println("v156:", v156)
    fmt.Println("v157:", v157)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v69:", v69)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v130
    _ = v131
    _ = v132
    _ = v155
    _ = v156
    _ = v157
    _ = v2
    _ = v3
    _ = v4
    _ = v69
    _ = v79
}
