package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int8 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 int; F2 uint8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint32; F1 int64; F2 int64 }
type GenS8 struct { F0 int8; F1 uint16 }
type GenS9 struct { F0 int64; F1 int64 }
type GenS10 struct { F0 int16; F1 uint32 }
type GenS11 struct { F0 uint16; F1 bool }
type GenS12 struct { F0 uint; F1 uint32; F2 int64 }
type GenS13 struct { F0 int8; F1 uint32; F2 bool }
type GenS14 struct { F0 bool; F1 uint16 }
type GenS15 struct { F0 string; F1 uint64 }
type GenS16 struct { F0 int; F1 int32 }
type GenS17 struct { F0 int }
type GenS18 struct { F0 uint64; F1 string }
func genFn0(p0 uint16, p1 uint16) string {
    fmt.Println(p0)
    for fi6 := 0; fi6 < 7; fi6++ {
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 uint32 = 2147483648
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 uint64 = 4294967295
            _ = fi9
            _ = v10
        }
        _ = fi6
    }
    v11 := GenS1{F0: -1, F1: 0, F2: 100}
    p1 -= 1
    var v12 bool = false
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return ""
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    var v4 string = "foo bar"
    v5 := make(map[uint32]int)
    switch v4 {
    case "":
        v4 = v4
    default:
        fmt.Println(v4)
    }
    v4 = ""
    var v13 uint16 = 1
    v14 := GenS2{F0: 1, F1: -2147483648}
    v15 := GenS3{F0: 9223372036854775807}
    var v16 uint16 = (^(v13))
    for fi17 := 0; fi17 < 10; fi17++ {
        if (true || false) {
            for fi18 := 0; fi18 < 13; fi18++ {
                v13 -= v13
                if false {
                    if true {
                        v19 := GenS1{F0: 0, F1: 65535, F2: -21}
                        fmt.Println(v16, fi18)
                        v20 := []int{fi18, (fi17 & (fi17 & 0))}
                        _ = v19
                        _ = v20
                    } else {
                        _ = len(v2)
                        var v21 int32 = 42
                        switch v4 {
                        case "":
                            if (false && false) {
                                if true {
                                    v22 := GenS3{F0: (fi18 | fi17)}
                                    v23 := GenS2{F0: 1, F1: (v21 | (v21 - (v21 & 2147483647)))}
                                    if true {
                                        v24 := []int{}
                                        if true {
                                            var v25 bool = (1000 != -967)
                                            _ = v25
                                        } else {
                                            var v26 float64 = 0.0
                                            v1[42] = v4
                                            if true {
                                                v21 -= v21
                                            }
                                            _ = v26
                                        }
                                        var v27 uint16 = (v16 >> ((4294967295) & 63))
                                        _ = v24
                                        _ = v27
                                    } else {
                                        var v28 int32 = 0
                                        if ((-9223372036854775808 != 42) || (-128 >= -128)) {
                                            v28 -= (v28 / ((0) | 1))
                                            v28 = (v21 - v28)
                                            var v29 uint = 1
                                            v28 -= -2147483648
                                            if (!(true && true)) {
                                                var v30 float32 = -2.25
                                                v31 := GenS2{F0: 42, F1: v28}
                                                switch ("a" + v4) {
                                                case "":
                                                    if false {
                                                        var v32 int64 = 0
                                                        v33 := GenS2{F0: -33, F1: v28}
                                                        if true {
                                                            v29 = (v29 / ((v29) | 1))
                                                            switch fi18 {
                                                            case 4294967295:
                                                                var v34 int = 801
                                                                _ = v34
                                                            case 402:
                                                                var v35 int32 = v28
                                                                _ = v35
                                                            case -2147483648:
                                                                v36 := GenS4{F0: 42, F1: 127, F2: 42}
                                                                _ = v36
                                                            }
                                                        } else {
                                                            var v37 uint = (v29 / ((42) | 1))
                                                            v38 := GenS4{F0: int(v32), F1: fi18, F2: 42}
                                                            _ = v37
                                                            _ = v38
                                                        }
                                                        _ = v32
                                                        _ = v33
                                                    } else {
                                                        fmt.Println(v16, v4, v4)
                                                        var v39 uint = (v29 + 0)
                                                        _ = v39
                                                    }
                                                    if (((v30 + 3.14159) + v30) > float32(1.5)) {
                                                        v13 = (v13 & v13)
                                                        var v40 int8 = -1
                                                        var v41 float64 = (float64(v30) * 0.1)
                                                        if true {
                                                            var v42 int64 = -73
                                                            v43 := []int{}
                                                            _ = len(v5)
                                                            var v44 uint8 = 89
                                                            _ = v42
                                                            _ = v43
                                                            _ = v44
                                                        } else {
                                                            v45 := GenS4{F0: fi18, F1: int(v13), F2: 42}
                                                            var v46 uint8 = uint8(v29)
                                                            _ = v45
                                                            _ = v46
                                                        }
                                                        _ = v40
                                                        _ = v41
                                                    }
                                                case "hello":
                                                    v13 = 65535
                                                case "\n\t":
                                                    var v47 uint = v29
                                                    fmt.Println(v28, fi17, v13)
                                                    v48 := GenS2{F0: 1, F1: v28}
                                                    _ = v47
                                                    _ = v48
                                                }
                                                _ = v30
                                                _ = v31
                                            }
                                            v21 = v21
                                            if (-128 > 0) {
                                                v16 -= v16
                                                fmt.Println(v13)
                                                switch fi18 {
                                                case -1:
                                                    fmt.Println(v29, v4)
                                                case -9223372036854775808:
                                                    var v49 int8 = 31
                                                    _ = v49
                                                default:
                                                    fmt.Println(v16, v4)
                                                }
                                            }
                                            v50 := GenS5{F0: v21}
                                            _ = v29
                                            _ = v50
                                        } else {
                                            switch -739 {
                                            case 32767:
                                                v28 -= (v28 | v28)
                                            case 42:
                                                var v51 int = (fi18 >> ((421) & 63))
                                                _ = v51
                                            }
                                            v13 = v16
                                            if false {
                                                v28 = (v28 | (v21 ^ (v28 ^ (v21 << ((1) & 63)))))
                                                v4 = "foo bar"
                                            } else {
                                                switch ("\n\t" + "") {
                                                case "foo bar":
                                                    delete(v5, 42)
                                                case "hello":
                                                    switch ("a" + ((v4 + "") + v4)) {
                                                    case "foo bar":
                                                        var v52 int64 = -1
                                                        _ = v52
                                                    default:
                                                        v16 = 32767
                                                        fmt.Println(v28, v4)
                                                        if (243 == 782) {
                                                            if true {
                                                                fmt.Println(v4)
                                                                v53 := GenS5{F0: 1}
                                                                var v54 string = ""
                                                                if true {
                                                                    var v55 int16 = 746
                                                                    fmt.Println(v13, fi18)
                                                                    _ = v55
                                                                }
                                                                var v56 int16 = 1
                                                                _ = v53
                                                                _ = v54
                                                                _ = v56
                                                            } else {
                                                                v57 := []int{}
                                                                v21 += -256
                                                                _ = v57
                                                            }
                                                            v4 = "hello"
                                                            v58 := GenS3{F0: fi18}
                                                            fmt.Println(fi18, v13, v16)
                                                            v21 = -1
                                                            delete(v3, 986)
                                                            v59 := GenS6{F0: (v28 / ((v21) | 1))}
                                                            var v60 float32 = (-2.25 - 3.14159)
                                                            _ = v58
                                                            _ = v59
                                                            _ = v60
                                                        } else {
                                                            v13 -= v16
                                                            v28 += v28
                                                        }
                                                    }
                                                case "\n\t":
                                                    var v61 int16 = 0
                                                    _ = v61
                                                }
                                            }
                                        }
                                        var v62 uint8 = 0
                                        _ = v28
                                        _ = v62
                                    }
                                    _ = v22
                                    _ = v23
                                } else {
                                    v5[175] = fi17
                                    v63 := []int{(fi17 << ((65535) & 63)), (fi17 | 128)}
                                    fmt.Println(v13, v21, v4)
                                    v4 = ""
                                    var v64 uint = 255
                                    _ = v63
                                    _ = v64
                                }
                            } else {
                                var v65 string = v4
                                v66 := GenS3{F0: 9223372036854775807}
                                if ((9223372036854775807 < -551) && true) {
                                    if false {
                                        v67 := GenS1{F0: 303, F1: 65535, F2: 100}
                                        if (!(128 >= 128)) {
                                            v68 := []int{}
                                            delete(v3, 2)
                                            v4 = v4
                                            var v69 int = 2147483647
                                            _ = v68
                                            _ = v69
                                        }
                                        if (42 != 1) {
                                            switch v65 {
                                            case "foo bar":
                                                var v70 float32 = (-2.25 - -2.25)
                                                _ = v70
                                            }
                                            fmt.Println(fi18, v65)
                                            v71 := []int{fi17, (fi18 | fi18)}
                                            _ = v71
                                        } else {
                                            v72 := GenS1{F0: 42, F1: 0, F2: 104}
                                            _ = v72
                                        }
                                        _ = v67
                                    }
                                    v73 := GenS3{F0: -2147483648}
                                    v74 := []int{(fi17 & fi17)}
                                    _ = v73
                                    _ = v74
                                } else {
                                    if false {
                                        if false {
                                            v13 = 1
                                        }
                                    }
                                    switch ((((v65 + v65) + (v65 + "foo bar")) + v65) + ("" + "a")) {
                                    case "\n\t":
                                        fmt.Println(v13, v21)
                                    case "foo bar":
                                        var v75 int32 = v21
                                        _ = v75
                                    default:
                                        v76 := []int{(fi17 << ((42) & 63)), fi17, (fi18 ^ (fi17 % (((^(fi18))) | 1))), (^(fi17))}
                                        _ = v76
                                    }
                                }
                                var v77 int16 = 0
                                var v78 int32 = (v21 ^ 2147483647)
                                var v79 uint32 = 65535
                                v77 -= v77
                                _ = v65
                                _ = v66
                                _ = v77
                                _ = v78
                                _ = v79
                            }
                        case "\n\t":
                            var v80 int = (fi18 << ((65535) & 63))
                            _ = v80
                        default:
                            v4 = "foo bar"
                        }
                        _ = v21
                    }
                    var v81 bool = true
                    _ = v81
                } else {
                    v82 := []int{-9223372036854775808, (fi18 - fi17)}
                    if (!true) {
                        switch (("a" + ("hello" + v4)) + v4) {
                        case "a":
                            v83 := GenS2{F0: 127, F1: 183}
                            _ = v83
                        case "\n\t":
                            var v84 float64 = 0.1
                            _ = v84
                        }
                        _ = len(v82)
                    }
                    _ = v82
                }
                var v85 int64 = 42
                _ = fi18
                _ = v85
            }
            if (1000 >= 0) {
                v86 := GenS7{F0: 4294967295, F1: 0, F2: -706}
                if (42 > 18446744073709551615) {
                    if false {
                        v87 := GenS8{F0: 2, F1: 475}
                        v88 := []int{}
                        _ = v87
                        _ = v88
                    } else {
                        v89 := GenS3{F0: -192}
                        var v90 uint64 = 4294967295
                        var v91 int32 = -2147483648
                        _ = v89
                        _ = v90
                        _ = v91
                    }
                    v4 = v4
                } else {
                    v16 += (v16 >> ((249) & 63))
                }
                var v92 uint8 = 138
                _ = v86
                _ = v92
            }
            v16 += (^(v13))
        } else {
            switch "a" {
            case "hello":
                v93 := GenS7{F0: 65535, F1: -1, F2: 1}
                _ = v93
            case "foo bar":
                switch fi17 {
                case 0:
                    for fi94 := 0; fi94 < 4; fi94++ {
                        var v95 uint16 = 32767
                        v13 = 255
                        var v96 int = fi94
                        if (false || (0 < 585)) {
                            var v97 float64 = ((1.5 * (-2.25 / 1e10)) + 1e10)
                            switch 42 {
                            case 0:
                                v96 += (fi17 - fi17)
                            case -516:
                                fmt.Println(fi17)
                            case 915:
                                fmt.Println(v95)
                            }
                            _ = v97
                        } else {
                            v95 = v16
                            v98 := []int{fi94, (fi17 << ((42) & 63)), int(v16), 2147483647}
                            v13 = (v16 * (v13 * 255))
                            _ = v98
                        }
                        _ = fi94
                        _ = v95
                        _ = v96
                    }
                case 9223372036854775807:
                    for fi99 := 0; fi99 < 15; fi99++ {
                        v100 := []int{fi99, (fi99 & (fi17 * (fi17 ^ (fi17 ^ (fi17 & (fi17 - (^(fi17)))))))), fi99}
                        switch fi17 {
                        case -2147483648:
                            var v101 float32 = -2.25
                            _ = v101
                        case -1:
                            fmt.Println(fi99, fi17)
                        case -264:
                            if (42 >= -990) {
                                var v102 int64 = 538
                                _ = v102
                            }
                            var v103 float64 = (0.1 / (((1.5 * float64(0.1)) - (1e10 - (0.0 - float64(0.1)))) - (float64(3.14159) / (0.1 / 1e10))))
                            _ = v103
                        }
                        _ = fi99
                        _ = v100
                    }
                case 42:
                    v16 -= v16
                }
            default:
                v104 := GenS7{F0: 2, F1: -639, F2: 421}
                _ = v104
            }
        }
        _ = fi17
    }
    var v105 uint64 = 933
    for fi106 := 0; fi106 < 16; fi106++ {
        fmt.Println(v105, v105, v105)
        _ = fi106
    }
    for fi107 := 0; fi107 < 10; fi107++ {
        for fi108 := 0; fi108 < 4; fi108++ {
            v109 := []int{fi108}
            if false {
                var v110 uint = 0
                var v111 int16 = 1000
                _ = v110
                _ = v111
            } else {
                if len(v109) > 0 {
                    v109[0] = fi107
                }
                switch v4 {
                case "a":
                    var v112 uint32 = 4294967295
                    _ = v112
                case "":
                    v109 = append(v109, 2147483647)
                default:
                    if true {
                        switch 372 {
                        case -9223372036854775808:
                            fmt.Println(v16, v4, v105)
                        case 188:
                            v113 := GenS3{F0: -2147483648}
                            _ = v113
                        default:
                            var v114 int = fi108
                            _ = v114
                        }
                    }
                }
            }
            fmt.Println(fi108, v16, fi108)
            _ = fi108
            _ = v109
        }
        switch "a" {
        case "foo bar":
            if (1.5 < 0.1) {
                _ = len(v2)
                v115 := []int{}
                _ = v115
            } else {
                v13 -= 255
                var v116 uint64 = (^(v105))
                for fi117 := 0; fi117 < 7; fi117++ {
                    v118 := v5[746]
                    _ = fi117
                    _ = v118
                }
                _ = v116
            }
        case "a":
            v105 += (v105 >> ((685) & 63))
        default:
            if (-1 != 444) {
                if false {
                    if false {
                        var v119 uint64 = v105
                        if true {
                            v16 -= 42
                            v13 += 255
                        }
                        fmt.Println(v105)
                        _ = v119
                    }
                } else {
                    v120 := v3[0]
                    v120 = v120
                    _ = v120
                }
                var v121 int32 = -817
                if (!(42 >= 561)) {
                    v121 -= v121
                    switch (fi107 >> ((282) & 63)) {
                    case 608:
                        v105 = v105
                    }
                } else {
                    var v122 int = (fi107 << ((2) & 63))
                    var v123 int32 = 609
                    var v124 float64 = 3.14159
                    _ = v122
                    _ = v123
                    _ = v124
                }
                _ = v121
            }
        }
        _ = fi107
    }
    v4 = (v4 + v4)
    v4 = ""
    if true {
        delete(v2, (v4 + (v4 + v4)))
        v125 := GenS9{F0: 9223372036854775807, F1: 577}
        if false {
            for fi126 := 0; fi126 < 11; fi126++ {
                v127 := GenS5{F0: -2147483648}
                _ = fi126
                _ = v127
            }
            v4 = v4
            if false {
                v4 = v4
                var v128 uint = 255
                _ = v128
            } else {
                if true {
                    for fi129 := 0; fi129 < 17; fi129++ {
                        var v130 float64 = 0.0
                        v131 := []int{}
                        _ = fi129
                        _ = v130
                        _ = v131
                    }
                } else {
                    var v132 float32 = (0.1 / 0.1)
                    _ = v132
                }
                for fi133 := 0; fi133 < 6; fi133++ {
                    for fi134 := 0; fi134 < 15; fi134++ {
                        var v135 int8 = 1
                        v135 = -103
                        _ = fi134
                        _ = v135
                    }
                    v13 += (v16 ^ v13)
                    if (float32(1.5) > ((3.14159 + 0.0) - 1.5)) {
                        var v136 bool = false
                        var v137 bool = v136
                        if v137 {
                            for fi138 := 0; fi138 < 11; fi138++ {
                                fmt.Println(v13)
                                _ = fi138
                            }
                            v137 = (-2147483648 >= 1)
                            var v139 int64 = 0
                            _ = v139
                        } else {
                            if (18446744073709551615 >= (v105 << ((2) & 63))) {
                                var v140 int64 = 42
                                _ = v140
                            }
                        }
                        _ = v136
                        _ = v137
                    } else {
                        var v141 uint64 = 1
                        if false {
                            v105 += (v105 & (v105 | (v141 ^ v105)))
                            var v142 int8 = -1
                            for fi143 := 0; fi143 < 6; fi143++ {
                                v144 := []int{}
                                v13 -= (v13 & (^(v16)))
                                _ = fi143
                                _ = v144
                            }
                            _ = v142
                        }
                        var v145 bool = (226 < 255)
                        _ = v141
                        _ = v145
                    }
                    _ = fi133
                }
                var v146 uint32 = 195
                _ = v146
            }
        } else {
            if true {
                v147 := GenS10{F0: 32767, F1: 874}
                for fi148 := 0; fi148 < 10; fi148++ {
                    if (!true) {
                        var v149 float64 = 1e10
                        _ = v149
                    } else {
                        switch v4 {
                        case "\n\t":
                            fmt.Println(v105, v105, v16)
                        default:
                            for fi150 := 0; fi150 < 20; fi150++ {
                                var v151 uint32 = 1
                                v152 := []int{}
                                _ = fi150
                                _ = v151
                                _ = v152
                            }
                        }
                    }
                    var v153 int64 = -9223372036854775808
                    v105 = (^(v105))
                    _ = fi148
                    _ = v153
                }
                _ = v147
            } else {
                if ((float64(1.5) - float64(3.14159)) >= -2.25) {
                    if (95 >= 127) {
                        switch 32768 {
                        case 9223372036854775807:
                            if (false && (65535 < 65535)) {
                                var v154 int64 = -875
                                v155 := v5[65535]
                                v16 += (v16 * 0)
                                _ = v154
                                _ = v155
                            } else {
                                fmt.Println(v13)
                                var v156 int16 = -96
                                v157 := GenS5{F0: 1}
                                _ = v156
                                _ = v157
                            }
                        default:
                            if false {
                                v4 = (v4 + ((v4 + (v4 + "hello")) + "a"))
                            } else {
                                var v158 float32 = -2.25
                                var v159 uint64 = v105
                                var v160 uint = 2
                                _ = v158
                                _ = v159
                                _ = v160
                            }
                        }
                        for fi161 := 0; fi161 < 17; fi161++ {
                            var v162 bool = false
                            switch fi161 {
                            case 127:
                                var v163 uint8 = 1
                                _ = v163
                            case 52:
                                switch 313 {
                                case 4294967295:
                                    v164 := GenS6{F0: 42}
                                    _ = v164
                                case 127:
                                    v3[1] = v4
                                default:
                                    if v162 {
                                        for fi165 := 0; fi165 < 15; fi165++ {
                                            v105 = 294
                                            if ("foo bar" <= v4) {
                                                if ((fi161 ^ -1) < (-(fi161))) {
                                                    v166 := []int{fi165, -620}
                                                    _ = v166
                                                }
                                                if (97 != 127) {
                                                    switch (fi161 & 0) {
                                                    case -397:
                                                        if ((0.0 - (1.5 + 1.5)) == 0.0) {
                                                            v13 += 2
                                                        }
                                                    case 4294967295:
                                                        v162 = (v4 == "")
                                                    }
                                                    var v167 int32 = 451
                                                    if v162 {
                                                        v168 := []int{fi161, 32768, fi161, (fi165 - (^(fi165)))}
                                                        _ = v168
                                                    } else {
                                                        switch (fi161 + fi161) {
                                                        case 4294967295:
                                                            var v169 uint8 = 207
                                                            _ = v169
                                                        case 127:
                                                            v170 := GenS9{F0: 0, F1: -9223372036854775808}
                                                            _ = v170
                                                        case -2147483648:
                                                            v167 = v167
                                                        default:
                                                            v171 := []int{(fi161 / (((fi161 ^ 1)) | 1)), 32767}
                                                            _ = v171
                                                        }
                                                    }
                                                    _ = v167
                                                }
                                            }
                                            _ = fi165
                                        }
                                        v172 := GenS11{F0: (^(v16)), F1: (fi161 == (fi161 ^ (fi161 ^ fi161)))}
                                        _ = v172
                                    } else {
                                        fmt.Println(v105)
                                        delete(v3, 0)
                                        if v162 {
                                            if v162 {
                                                v173 := GenS7{F0: 601, F1: 122, F2: 9223372036854775807}
                                                _ = v173
                                            } else {
                                                fmt.Println(v13, v105)
                                            }
                                            switch int(v13) {
                                            case 0:
                                                if (42 > 255) {
                                                    var v174 int8 = 121
                                                    _ = v174
                                                }
                                            case 2147483647:
                                                if (!v162) {
                                                    if v162 {
                                                        var v175 int32 = 42
                                                        v176 := []int{(^(fi161))}
                                                        _ = v175
                                                        _ = v176
                                                    }
                                                }
                                            case 299:
                                                for fi177 := 0; fi177 < 16; fi177++ {
                                                    v178 := GenS12{F0: 4294967295, F1: uint32(v105), F2: 42}
                                                    _ = fi177
                                                    _ = v178
                                                }
                                            default:
                                                v4 = v4
                                            }
                                        } else {
                                            v105 -= 1
                                            var v179 uint16 = v13
                                            _ = v179
                                        }
                                    }
                                }
                            case 2147483647:
                                for fi180 := 0; fi180 < 16; fi180++ {
                                    v105 = (^(v105))
                                    var v181 int64 = 42
                                    var v182 int32 = -2147483648
                                    _ = fi180
                                    _ = v181
                                    _ = v182
                                }
                            }
                            _ = fi161
                            _ = v162
                        }
                    } else {
                        if true {
                            if (815 != 21) {
                                switch 0 {
                                case 731:
                                    var v183 int8 = 0
                                    _ = v183
                                case 9223372036854775807:
                                    switch v4 {
                                    case "foo bar":
                                        v1[4294967295] = "a"
                                    }
                                default:
                                    for fi184 := 0; fi184 < 16; fi184++ {
                                        if false {
                                            for fi185 := 0; fi185 < 10; fi185++ {
                                                v186 := GenS13{F0: 0, F1: 65535, F2: false}
                                                _ = fi185
                                                _ = v186
                                            }
                                            fmt.Println(v16)
                                            var v187 int8 = 127
                                            _ = v187
                                        } else {
                                            for fi188 := 0; fi188 < 7; fi188++ {
                                                fmt.Println(v13, fi184)
                                                fmt.Println(fi184)
                                                if true {
                                                    var v189 uint8 = 1
                                                    if (((-2.25 - 1.5) > ((1e10 * 0.0) + float32(0.0))) && false) {
                                                        _ = len(v2)
                                                    }
                                                    _ = v189
                                                }
                                                _ = fi188
                                            }
                                            _ = len(v2)
                                        }
                                        v4 = ((v4 + "foo bar") + ((v4 + (v4 + (((v4 + v4) + v4) + ("\n\t" + (v4 + (v4 + "foo bar")))))) + v4))
                                        _ = fi184
                                    }
                                }
                                var v190 float32 = -2.25
                                switch (v4 + (((v4 + v4) + "") + v4)) {
                                case "\n\t":
                                    v191 := GenS5{F0: -360}
                                    _ = v191
                                default:
                                    var v192 float32 = (0.0 + (float32(1.5) - v190))
                                    _ = v192
                                }
                                _ = v190
                            }
                        }
                        for fi193 := 0; fi193 < 16; fi193++ {
                            var v194 int8 = 0
                            v4 = ""
                            _ = fi193
                            _ = v194
                        }
                    }
                    v13 = 1
                }
                var v195 int16 = -1
                _ = v195
            }
            for fi196 := 0; fi196 < 14; fi196++ {
                if true {
                    for fi197 := 0; fi197 < 18; fi197++ {
                        switch fi197 {
                        case 4294967295:
                            v105 = 995
                        case 32768:
                            v198 := GenS2{F0: 19, F1: -2147483648}
                            _ = v198
                        }
                        v16 -= 2
                        v13 -= (v16 - 112)
                        var v199 bool = (false || (588 <= 82))
                        switch v4 {
                        case "\n\t":
                            var v200 bool = (v199 || v199)
                            _ = v200
                        case "":
                            v16 = (v16 << ((273) & 63))
                        case "a":
                            v201 := []int{(fi196 | fi197), (fi196 / (((fi197 << ((347) & 63))) | 1)), fi197, fi196}
                            _ = v201
                        }
                        _ = fi197
                        _ = v199
                    }
                } else {
                    v202 := []int{0, 477}
                    if true {
                        v16 = 156
                        var v203 float64 = -2.25
                        v204 := GenS11{F0: (v16 << ((18446744073709551615) & 63)), F1: true}
                        _ = v203
                        _ = v204
                    }
                    var v205 string = "hello"
                    _ = v202
                    _ = v205
                }
                _ = fi196
            }
        }
        _ = v125
    } else {
        if true {
            var v206 string = v4
            _ = v206
        } else {
            for fi207 := 0; fi207 < 1; fi207++ {
                var v208 int32 = 0
                if true {
                    v209 := GenS3{F0: 128}
                    _ = v209
                } else {
                    v210 := GenS5{F0: (v208 * v208)}
                    for fi211 := 0; fi211 < 2; fi211++ {
                        v4 = (v4 + "foo bar")
                        if ((float64(((float32(1.5) + (float32((0.1 - ((0.1 * 0.1) * 1e10))) / float32((((((1.5 + (0.1 - (0.1 + float64(((1e10 * (3.14159 + ((float32(1e10) / 3.14159) + -2.25))) / 3.14159))))) * 1.5) / 1.0) - 0.1) * 1e10)))) * 1e10)) < (float64(1e10) * (-2.25 / 1.0))) && false) {
                            var v212 bool = ((589 <= 159) && true)
                            if v212 {
                                var v213 int8 = -121
                                fmt.Println(v4, v212, v4)
                                v16 += 117
                                _ = v213
                            }
                            _ = v212
                        } else {
                            v214 := GenS14{F0: false, F1: v16}
                            _ = v214
                        }
                        _ = fi211
                    }
                    _ = v210
                }
                _ = fi207
                _ = v208
            }
        }
    }
    for fi215 := 0; fi215 < 18; fi215++ {
        v216 := GenS15{F0: "\n\t", F1: (v105 ^ (v105 | (v105 & v105)))}
        v16 = (v13 << ((251) & 63))
        v217 := GenS15{F0: v4, F1: v105}
        _ = fi215
        _ = v216
        _ = v217
    }
    switch v4 {
    case "foo bar":
        if (65535 == (v16 / (((v16 + (^(v13)))) | 1))) {
            switch 9223372036854775807 {
            case 678:
                var v218 uint8 = 91
                _ = v218
            case 42:
                switch 9223372036854775807 {
                case -970:
                    switch -403 {
                    case 0:
                        switch 32768 {
                        case -884:
                            fmt.Println(v105, v16)
                        case -1:
                            v219 := []int{int(v13), 2147483647}
                            _ = v219
                        default:
                            v105 = (v105 << ((uint(v13)) & 63))
                        }
                    case -645:
                        if true {
                            switch "hello" {
                            case "a":
                                v220 := GenS16{F0: -9223372036854775808, F1: -2147483648}
                                _ = v220
                            case "foo bar":
                                var v221 uint8 = 92
                                _ = v221
                            case "hello":
                                fmt.Println(v105)
                            }
                            var v222 uint32 = 42
                            var v223 float64 = 0.1
                            _ = v222
                            _ = v223
                        }
                    }
                }
            default:
                v16 -= v13
            }
            fmt.Println(v16, v16)
        }
    default:
        if false {
            switch 2147483647 {
            case 4294967295:
                for fi224 := 0; fi224 < 14; fi224++ {
                    v4 = v4
                    v225 := []int{(fi224 - 128), fi224, 32767, (fi224 - (^(fi224)))}
                    v226 := GenS17{F0: (fi224 << ((821) & 63))}
                    _ = fi224
                    _ = v225
                    _ = v226
                }
            case 0:
                fmt.Println(v105, v4, v4)
            case 32767:
                v227 := GenS18{F0: (v105 + v105), F1: v4}
                _ = v227
            default:
                v5[4294967295] = 128
            }
        } else {
            if true {
                v4 = v4
                var v228 uint32 = 42
                _ = v228
            } else {
                v229 := GenS7{F0: 537, F1: -1, F2: 0}
                _ = v229
            }
            fmt.Println(v16, v105, v105)
            var v230 float32 = -2.25
            _ = v230
        }
    }
    v105 -= (^(v105))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v105:", v105)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v105
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
