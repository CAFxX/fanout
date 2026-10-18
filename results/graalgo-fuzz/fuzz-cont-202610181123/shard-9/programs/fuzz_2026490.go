package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 int16; F2 uint32 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 uint; F1 string }
type GenS6 struct { F0 uint; F1 bool }
type GenS7 struct { F0 uint64; F1 uint32 }
type GenS8 struct { F0 uint16; F1 int8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint64; F1 int8; F2 int32 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int32; F1 int8 }
type GenS13 struct { F0 int; F1 uint16 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 int8; F1 int16; F2 int }
type GenS16 struct { F0 int16; F1 uint32; F2 int }
type GenS17 struct { F0 bool }
type GenS18 struct { F0 int64; F1 int32 }
func genFn0() uint {
    if true {
        var v7 int8 = 127
        v7 = -75
        _ = v7
    } else {
        var v8 int16 = 42
        v9 := GenS1{F0: 42, F1: (^(v8)), F2: (v8 >> ((0) & 63))}
        for fi10 := 0; fi10 < 4; fi10++ {
            for fi11 := 0; fi11 < 11; fi11++ {
                v12 := []int{1}
                v8 = (-(v8))
                fmt.Println(v8, v8, fi11)
                _ = fi11
                _ = v12
            }
            fmt.Println(v8, v8)
            v13 := []int{-2147483648}
            _ = fi10
            _ = v13
        }
        _ = v8
        _ = v9
    }
    if false {
        var v14 uint32 = 2147483648
        var v15 float64 = 0.0
        v15 = (v15 - (v15 * float64(3.14159)))
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 float64 = 0.1
        _ = fi16
        _ = v17
    }
    switch 4294967295 {
    case 4294967295:
        for fi18 := 0; fi18 < 7; fi18++ {
            if ((true || ((0.1 + 0.1) != float32(1.5))) || (!(!false))) {
                var v19 string = "\n\t"
                _ = v19
            } else {
                if (32767 != -1) {
                    var v20 int32 = 1
                    _ = v20
                } else {
                    var v21 string = "\n\t"
                    v22 := GenS2{F0: 42}
                    v23 := GenS3{F0: (fi18 ^ (fi18 | (fi18 + (fi18 | (fi18 % (((fi18 - fi18)) | 1)))))), F1: -1, F2: 2}
                    _ = v21
                    _ = v22
                    _ = v23
                }
                fmt.Println(fi18)
            }
            switch -9223372036854775808 {
            case -2147483648:
                v24 := GenS4{F0: (-(fi18)), F1: -128}
                _ = v24
            case 1:
                switch "foo bar" {
                case "a":
                    for fi25 := 0; fi25 < 9; fi25++ {
                        switch "hello" {
                        case "a":
                            var v26 float32 = 0.0
                            _ = v26
                        case "foo bar":
                            var v27 uint = 255
                            _ = v27
                        case "hello":
                            if true {
                                var v28 float64 = (0.1 / float64((1e10 / 3.14159)))
                                _ = v28
                            } else {
                                var v29 int = 0
                                v30 := GenS1{F0: 255, F1: -5, F2: 32767}
                                var v31 string = "hello"
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                        default:
                            var v32 uint8 = 128
                            _ = v32
                        }
                        var v33 uint8 = 72
                        _ = fi25
                        _ = v33
                    }
                case "\n\t":
                    if (!false) {
                        if false {
                            v34 := []int{fi18}
                            fmt.Println(fi18)
                            _ = v34
                        }
                    }
                case "":
                    var v35 int = fi18
                    _ = v35
                }
            case -986:
                var v36 uint16 = 42
                _ = v36
            default:
                v37 := GenS4{F0: fi18, F1: 0}
                _ = v37
            }
            var v38 uint64 = 334
            _ = fi18
            _ = v38
        }
    case 141:
        var v39 int16 = 32767
        _ = v39
    case -544:
        switch 32767 {
        case 244:
            if true {
                var v40 int8 = 77
                switch 4294967295 {
                case -9223372036854775808:
                    fmt.Println(v40)
                default:
                    switch ("foo bar" + "\n\t") {
                    case "hello":
                        v41 := []int{388, 612, 513}
                        _ = v41
                    case "":
                        var v42 uint16 = 700
                        _ = v42
                    default:
                        switch 682 {
                        case 4294967295:
                            if false {
                                if false {
                                    v40 = v40
                                }
                                switch ("" + ("hello" + "a")) {
                                case "foo bar":
                                    var v43 int64 = -120
                                    _ = v43
                                }
                            }
                        case 32767:
                            v40 = -16
                        }
                    }
                }
                _ = v40
            } else {
                v44 := []int{32768, 0, 4294967295}
                if len(v44) >= 2 {
                    v44 = v44[1:len(v44)]
                }
                _ = v44
            }
        }
    default:
        v45 := GenS5{F0: 1, F1: "foo bar"}
        _ = v45
    }
    return 0
}
func genFn1(p0 int32, p1 string, p2 int16) uint {
    var v46 int = -2147483648
    if (!false) {
        for fi47 := 0; fi47 < 11; fi47++ {
            var v48 float32 = (float32(0.0) + 1e10)
            v49 := []int{32767, 9223372036854775807, (fi47 ^ (v46 | (-(fi47)))), (v46 % ((4294967295) | 1))}
            var v50 uint = 2
            _ = fi47
            _ = v48
            _ = v49
            _ = v50
        }
    } else {
        p0 -= (p0 % ((p0) | 1))
        p0 -= p0
        switch v46 {
        case 1:
            v51 := []int{v46, v46, v46, v46}
            _ = v51
        default:
            var v52 float32 = -2.25
            _ = v52
        }
    }
    v46 += 1
    switch v46 {
    case 4294967295:
        v53 := []int{(v46 * 127), v46, v46, v46}
        _ = v53
    case 9223372036854775807:
        var v54 float32 = 0.0
        _ = v54
    case -9223372036854775808:
        v55 := GenS1{F0: 121, F1: p2, F2: (-(p2))}
        _ = v55
    default:
        for fi56 := 0; fi56 < 20; fi56++ {
            v57 := GenS4{F0: (-(v46)), F1: -1}
            switch "" {
            case "":
                if false {
                    p2 += p2
                }
            case "a":
                for fi58 := 0; fi58 < 5; fi58++ {
                    v59 := GenS6{F0: 1, F1: false}
                    v60 := GenS2{F0: genFn0()}
                    var v61 float32 = 1.5
                    fmt.Println(p2, p1)
                    _ = fi58
                    _ = v59
                    _ = v60
                    _ = v61
                }
            case "\n\t":
                if ((1.5 * 1e10) == 3.14159) {
                    v62 := []int{127, fi56}
                    p0 = (p0 % ((p0) | 1))
                    _ = v62
                } else {
                    v63 := GenS4{F0: 1, F1: -128}
                    for fi64 := 0; fi64 < 8; fi64++ {
                        v65 := []int{}
                        p1 = p1
                        var v66 uint = 0
                        switch v46 {
                        case 263:
                            if false {
                                p1 = p1
                            } else {
                                if (100 == 1) {
                                    if len(v65) > 0 {
                                        v65[0] = fi64
                                    }
                                    var v67 uint64 = 18446744073709551615
                                    if len(v65) > 0 {
                                        v65[0] = (v46 * 1)
                                    }
                                    var v68 float32 = (0.1 - 1.5)
                                    var v69 float64 = -2.25
                                    _ = v67
                                    _ = v68
                                    _ = v69
                                } else {
                                    v70 := GenS7{F0: 2, F1: 2147483648}
                                    if false {
                                        if len(v65) > 0 {
                                            v65[0] = fi64
                                        }
                                    }
                                    switch p1 {
                                    case "":
                                        var v71 int = (fi56 + (fi56 << ((967) & 63)))
                                        _ = v71
                                    default:
                                        p1 = "hello"
                                    }
                                    _ = v70
                                }
                            }
                        case 1:
                            switch fi64 {
                            case 47:
                                v46 = -353
                            case -2147483648:
                                switch fi64 {
                                case 118:
                                    if (p1 <= p1) {
                                        if false {
                                            var v72 int = fi64
                                            switch p1 {
                                            case "a":
                                                var v73 uint = (^(v66))
                                                _ = v73
                                            }
                                            fmt.Println(v72, v66, fi64)
                                            _ = v72
                                        } else {
                                            var v74 int32 = p0
                                            switch 4294967295 {
                                            case 32768:
                                                var v75 int32 = (v74 % ((-41) | 1))
                                                _ = v75
                                            case 4294967295:
                                                p1 = (p1 + (p1 + "a"))
                                            }
                                            p2 = p2
                                            fmt.Println(p2, fi56)
                                            var v76 string = "foo bar"
                                            _ = v74
                                            _ = v76
                                        }
                                        var v77 uint = (v66 ^ v66)
                                        p0 = p0
                                        _ = v77
                                    }
                                case -2147483648:
                                    var v78 uint = genFn0()
                                    _ = v78
                                case 42:
                                    if (32768 >= (-(fi64))) {
                                        if ((1e10 != 1.5) && true) {
                                            v46 += (fi56 * fi64)
                                            if (true && true) {
                                                v65 = append(v65, (v46 << ((327) & 63)))
                                                if (1 != 2) {
                                                    var v79 int32 = p0
                                                    if false {
                                                        v66 = v66
                                                        if true {
                                                            switch (fi56 + -830) {
                                                            case -1:
                                                                v80 := GenS1{F0: 255, F1: (^(p2)), F2: (p2 * (-(p2)))}
                                                                _ = v80
                                                            default:
                                                                v81 := GenS1{F0: 255, F1: (p2 | p2), F2: (p2 | p2)}
                                                                _ = v81
                                                            }
                                                            if (!((!false) || (false && (179 == 1)))) {
                                                                fmt.Println(v79, p0)
                                                            } else {
                                                                var v82 float32 = 0.0
                                                                _ = v82
                                                            }
                                                        }
                                                    }
                                                    v83 := GenS8{F0: 65535, F1: 1}
                                                    _ = v79
                                                    _ = v83
                                                }
                                                var v84 bool = true
                                                _ = v84
                                            }
                                            var v85 uint32 = 0
                                            fmt.Println(v66)
                                            switch fi64 {
                                            case -2147483648:
                                                v65 = append(v65, (fi64 / (((fi56 - -2147483648)) | 1)))
                                            default:
                                                var v86 string = p1
                                                _ = v86
                                            }
                                            _ = len(v65)
                                            _ = v85
                                        }
                                        var v87 float64 = (0.0 * (0.0 - 0.1))
                                        var v88 string = p1
                                        _ = v87
                                        _ = v88
                                    } else {
                                        v46 = -9223372036854775808
                                    }
                                    p1 = ("hello" + "hello")
                                    var v89 int8 = -128
                                    _ = v89
                                default:
                                    var v90 int64 = -139
                                    _ = v90
                                }
                            default:
                                p2 -= 0
                            }
                            _ = len(v65)
                            var v91 float32 = (1e10 - 1.5)
                            v92 := GenS9{F0: v66}
                            _ = v91
                            _ = v92
                        }
                        var v93 uint64 = 2
                        v94 := GenS10{F0: v93, F1: int8(v93), F2: p0}
                        _ = fi64
                        _ = v65
                        _ = v66
                        _ = v93
                        _ = v94
                    }
                    v46 += v46
                    _ = v63
                }
            default:
                if ((^(p0)) <= (^(p0))) {
                    var v95 int64 = 0
                    var v96 int32 = p0
                    _ = v95
                    _ = v96
                }
            }
            _ = fi56
            _ = v57
        }
    }
    if false {
        var v97 int16 = p2
        v98 := GenS7{F0: 2, F1: 65535}
        _ = v97
        _ = v98
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v46
    return 495
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{763, 0, 1, 2}
    v3 := make(map[uint32]uint32)
    v4 := []string{""}
    var v5 int16 = 0
    var v6 uint64 = 627
    v6 = v6
    fmt.Println(v6, v6)
    fmt.Println(v5)
    if false {
        v5 = v5
        fmt.Println(v6, v6)
        var v99 int8 = -128
        _ = v99
    } else {
        if (true && true) {
            v6 += 964
            switch "foo bar" {
            case "a":
                var v100 bool = (42 <= 1)
                _ = v100
            }
        } else {
            v5 -= -894
            fmt.Println(v5)
            if true {
                v6 -= (^(v6))
                fmt.Println(v5)
            }
        }
        for fi101 := 0; fi101 < 15; fi101++ {
            var v102 int64 = 42
            _ = fi101
            _ = v102
        }
        if ((("foo bar" + "foo bar") + "") < "") {
            v5 -= 42
        }
    }
    var v103 int32 = -2147483648
    v6 = v6
    switch ("" + "\n\t") {
    case "a":
        v104 := GenS2{F0: 417}
        _ = v104
    case "hello":
        var v105 bool = false
        _ = v105
    default:
        for fi106 := 0; fi106 < 14; fi106++ {
            var v107 uint = 555
            _ = fi106
            _ = v107
        }
    }
    if (0 > -9223372036854775808) {
        for fi108 := 0; fi108 < 4; fi108++ {
            for fi109 := 0; fi109 < 3; fi109++ {
                var v110 uint8 = 214
                if true {
                    fmt.Println(fi109, v103)
                }
                _ = fi109
                _ = v110
            }
            var v111 float64 = ((0.0 - ((1e10 / 1e10) / (float64(-2.25) / (1e10 / 0.1)))) / 1.0)
            _ = fi108
            _ = v111
        }
        switch "\n\t" {
        case "hello":
            switch -2147483648 {
            case -739:
                v5 = -32768
            case 0:
                if ((v6 | (v6 * 175)) == (v6 & (v6 << ((105) & 63)))) {
                    v112 := GenS11{F0: 687}
                    var v113 int16 = (v5 * v5)
                    v114 := v1["a"]
                    _ = v112
                    _ = v113
                    _ = v114
                }
            case -2147483648:
                v115 := GenS1{F0: 29, F1: (v5 >> ((42) & 63)), F2: (v5 << ((genFn0()) & 63))}
                _ = v115
            }
        case "foo bar":
            switch "" {
            case "":
                v116 := v1[("" + "a")]
                _ = v116
            case "foo bar":
                var v117 uint64 = 256
                _ = v117
            case "hello":
                var v118 uint64 = 18446744073709551615
                _ = v118
            default:
                for fi119 := 0; fi119 < 9; fi119++ {
                    if (((1e10 / 1.0) - ((-2.25 * 1.5) + 1.5)) < float32((0.0 + 3.14159))) {
                        var v120 uint = 0
                        _ = v120
                    } else {
                        v5 += -173
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    _ = fi119
                }
            }
        case "\n\t":
            fmt.Println(v5, v103)
        }
    } else {
        if len(v2) > 0 {
            v2[0] = 42
        }
    }
    if false {
        for fi121 := 0; fi121 < 19; fi121++ {
            for fi122 := 0; fi122 < 18; fi122++ {
                switch (fi122 / (((^(fi122))) | 1)) {
                case -9223372036854775808:
                    v6 -= v6
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                if ((1 == -115) || (255 >= 32767)) {
                    v3[2147483647] = 1
                    var v123 float32 = 1e10
                    _ = v123
                }
                fmt.Println(v5, v103)
                _ = fi122
            }
            fmt.Println(fi121)
            v124 := GenS1{F0: 127, F1: v5, F2: v5}
            _ = fi121
            _ = v124
        }
        if false {
            v6 -= v6
            v125 := GenS8{F0: 65535, F1: 42}
            _ = v125
        } else {
            switch "" {
            case "hello":
                var v126 uint32 = 524
                _ = v126
            case "\n\t":
                v5 = v5
            }
            v6 = v6
        }
        v127 := GenS6{F0: genFn1(v103, ("hello" + "foo bar"), (v5 | v5)), F1: false}
        _ = v127
    }
    if (42 != 0) {
        if (3.14159 > 0.1) {
            for fi128 := 0; fi128 < 9; fi128++ {
                var v129 uint = uint(v6)
                _ = fi128
                _ = v129
            }
            var v130 uint32 = 2
            _ = v130
        }
        v131 := GenS12{F0: -520, F1: 11}
        if len(v4) > 0 {
            v4[0] = "a"
        }
        _ = v131
    } else {
        fmt.Println(v5, v103, v5)
        v5 -= 1000
        v132 := GenS12{F0: 1, F1: 0}
        _ = v132
    }
    var v133 string = "hello"
    v3[uint32(v6)] = 597
    v4 = append(v4, "foo bar")
    if (!(true || ((2147483647 != -530) && (true && true)))) {
        var v134 float32 = 1.5
        for fi135 := 0; fi135 < 3; fi135++ {
            var v136 int32 = v103
            v133 = "\n\t"
            v134 = v134
            _ = fi135
            _ = v136
        }
        if false {
            if (!false) {
                if true {
                    var v137 float64 = (-2.25 - 3.14159)
                    if (("\n\t" < ((((v133 + "hello") + v133) + v133) + v133)) || true) {
                        v138 := GenS1{F0: 127, F1: -326, F2: (v5 | v5)}
                        _ = v138
                    }
                    switch 967 {
                    case 360:
                        switch -9223372036854775808 {
                        case 210:
                            v137 -= float64(0.0)
                        case -89:
                            v103 += v103
                        case -150:
                            if false {
                                for fi139 := 0; fi139 < 3; fi139++ {
                                    var v140 int32 = (v103 + -716)
                                    _ = fi139
                                    _ = v140
                                }
                            } else {
                                v141 := GenS3{F0: 4294967295, F1: v5, F2: 2147483648}
                                if true {
                                    if true {
                                        switch 32767 {
                                        case -1:
                                            fmt.Println(v133)
                                        case 1:
                                            var v142 string = ((v133 + (v133 + "a")) + (v133 + v133))
                                            _ = v142
                                        default:
                                            var v143 bool = (!true)
                                            _ = v143
                                        }
                                        for fi144 := 0; fi144 < 16; fi144++ {
                                            switch v133 {
                                            case "":
                                                if true {
                                                    for fi145 := 0; fi145 < 16; fi145++ {
                                                        v146 := v1[v133]
                                                        _ = fi145
                                                        _ = v146
                                                    }
                                                    if (true && (1.5 < 1e10)) {
                                                        if false {
                                                            for fi147 := 0; fi147 < 8; fi147++ {
                                                                if ((v5 != int16(v6)) || true) {
                                                                    v148 := GenS7{F0: (v6 >> ((genFn0()) & 63)), F1: 654}
                                                                    v5 -= v5
                                                                    v149 := GenS13{F0: -2147483648, F1: 607}
                                                                    _ = v148
                                                                    _ = v149
                                                                } else {
                                                                    if len(v4) > 0 {
                                                                        v4[0] = ""
                                                                    }
                                                                }
                                                                _ = fi147
                                                            }
                                                        } else {
                                                            for fi150 := 0; fi150 < 8; fi150++ {
                                                                v151 := GenS5{F0: 65535, F1: v133}
                                                                v152 := GenS10{F0: (^(v6)), F1: 3, F2: v103}
                                                                v153 := GenS14{F0: 920}
                                                                _ = fi150
                                                                _ = v151
                                                                _ = v152
                                                                _ = v153
                                                            }
                                                            v154 := GenS3{F0: fi144, F1: v5, F2: 4294967295}
                                                            _ = v154
                                                        }
                                                        v134 = v134
                                                        if true {
                                                            var v155 string = "a"
                                                            for fi156 := 0; fi156 < 7; fi156++ {
                                                                v5 -= v5
                                                                if false {
                                                                    v137 = (v137 - ((((v137 - float64(v134)) * (v137 - ((v137 + v137) + 1.5))) / float64(((v134 * v134) * ((-2.25 / 1.0) * v134)))) + float64(3.14159)))
                                                                    if true {
                                                                        var v157 int64 = -269
                                                                        var v158 uint64 = v6
                                                                        _ = v157
                                                                        _ = v158
                                                                    } else {
                                                                        var v159 int16 = 42
                                                                        var v160 uint16 = 805
                                                                        v161 := GenS7{F0: (v6 % (((v6 & v6)) | 1)), F1: 0}
                                                                        v103 = (-(v103))
                                                                        var v162 int8 = 127
                                                                        if true {
                                                                            fmt.Println(v160, v134)
                                                                            fmt.Println(fi156)
                                                                        }
                                                                        _ = v159
                                                                        _ = v160
                                                                        _ = v161
                                                                        _ = v162
                                                                    }
                                                                }
                                                                switch v133 {
                                                                case "foo bar":
                                                                    var v163 uint64 = (v6 >> ((57) & 63))
                                                                    v5 = -32768
                                                                    v164 := GenS15{F0: 100, F1: v5, F2: 931}
                                                                    _ = v163
                                                                    _ = v164
                                                                case "a":
                                                                    if (!true) {
                                                                        v165 := v3[2147483647]
                                                                        v103 -= (v103 >> ((4294967295) & 63))
                                                                        var v166 uint8 = 2
                                                                        if true {
                                                                            v167 := GenS4{F0: fi144, F1: 8}
                                                                            switch v155 {
                                                                            case "a":
                                                                                if false {
                                                                                    v5 -= -1
                                                                                    fmt.Println(v5)
                                                                                    fmt.Println(v133)
                                                                                }
                                                                            }
                                                                            if true {
                                                                                var v168 bool = true
                                                                                switch (fi144 >> ((2) & 63)) {
                                                                                case -89:
                                                                                    if true {
                                                                                        var v169 uint32 = 644
                                                                                        v166 = v166
                                                                                        _ = v169
                                                                                    }
                                                                                }
                                                                                v170 := GenS4{F0: fi144, F1: -126}
                                                                                _ = v168
                                                                                _ = v170
                                                                            }
                                                                            _ = v167
                                                                        } else {
                                                                            var v171 uint = 798
                                                                            fmt.Println(v155, v133, v155)
                                                                            v165 -= (^(v165))
                                                                            switch (v155 + (v155 + (v155 + v155))) {
                                                                            case "hello":
                                                                                v172 := GenS15{F0: int8(v166), F1: 550, F2: 9223372036854775807}
                                                                                _ = v172
                                                                            }
                                                                            _ = v171
                                                                        }
                                                                        _ = v165
                                                                        _ = v166
                                                                    } else {
                                                                        if true {
                                                                            switch fi144 {
                                                                            case 537:
                                                                                if (true && true) {
                                                                                    if true {
                                                                                        delete(v1, (v155 + (v155 + v155)))
                                                                                        if false {
                                                                                            if false {
                                                                                                var v173 uint32 = 379
                                                                                                if false {
                                                                                                    if len(v4) >= 2 {
                                                                                                        v4 = v4[1:len(v4)]
                                                                                                    }
                                                                                                    v134 = v134
                                                                                                    if ((!false) && true) {
                                                                                                        v155 = ((((v133 + v133) + v133) + "hello") + ((((v155 + ("" + v133)) + v155) + v155) + "a"))
                                                                                                        if (1 >= 1) {
                                                                                                            v103 = v103
                                                                                                            v133 = v155
                                                                                                        }
                                                                                                        v174 := GenS16{F0: v5, F1: v173, F2: fi156}
                                                                                                        _ = v174
                                                                                                    } else {
                                                                                                        v175 := GenS17{F0: false}
                                                                                                        var v176 int8 = 39
                                                                                                        v134 = ((v134 / (v134 + (v134 + (v134 * float32(v137))))) * (v134 - 1e10))
                                                                                                        if false {
                                                                                                            _ = len(v4)
                                                                                                        }
                                                                                                        _ = v175
                                                                                                        _ = v176
                                                                                                    }
                                                                                                }
                                                                                                _ = v173
                                                                                            } else {
                                                                                                v177 := GenS14{F0: 1}
                                                                                                if true {
                                                                                                    if ((v103 << ((2) & 63)) == v103) {
                                                                                                        v178 := GenS9{F0: genFn1((v103 ^ (v103 + (v103 - v103))), v133, 42)}
                                                                                                        v179 := GenS18{F0: 1, F1: v103}
                                                                                                        _ = v178
                                                                                                        _ = v179
                                                                                                    }
                                                                                                    fmt.Println(v5)
                                                                                                } else {
                                                                                                    v180 := GenS18{F0: -519, F1: v103}
                                                                                                    var v181 int32 = v103
                                                                                                    _ = v180
                                                                                                    _ = v181
                                                                                                }
                                                                                                if (v6 >= (v6 + v6)) {
                                                                                                    var v182 int8 = 42
                                                                                                    switch "" {
                                                                                                    case "a":
                                                                                                        _ = len(v3)
                                                                                                    case "":
                                                                                                        var v183 int8 = (v182 % (((v182 | (v182 | (v182 >> ((213) & 63))))) | 1))
                                                                                                        _ = v183
                                                                                                    case "hello":
                                                                                                        var v184 float32 = 3.14159
                                                                                                        _ = v184
                                                                                                    }
                                                                                                    switch "foo bar" {
                                                                                                    case "a":
                                                                                                        if true {
                                                                                                            fmt.Println(fi144, v155)
                                                                                                            v137 += (v137 - v137)
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v182
                                                                                                } else {
                                                                                                    v133 = ("\n\t" + v155)
                                                                                                }
                                                                                                _ = v177
                                                                                            }
                                                                                            v185 := GenS2{F0: 255}
                                                                                            v186 := GenS7{F0: 561, F1: 4294967295}
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                            v187 := GenS18{F0: 7, F1: v103}
                                                                                            _ = v185
                                                                                            _ = v186
                                                                                            _ = v187
                                                                                        } else {
                                                                                            v137 += -2.25
                                                                                            if false {
                                                                                                v134 = (v134 - v134)
                                                                                            }
                                                                                            if (true && false) {
                                                                                                v5 -= -476
                                                                                                var v188 bool = false
                                                                                                _ = v188
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    delete(v3, 2)
                                                                                    if false {
                                                                                        v2 = append(v2, 1)
                                                                                    }
                                                                                } else {
                                                                                    v103 = 113
                                                                                }
                                                                            case 1:
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = 4294967295
                                                                                }
                                                                            }
                                                                            _ = len(v3)
                                                                        } else {
                                                                            switch "a" {
                                                                            case "\n\t":
                                                                                switch "a" {
                                                                                case "a":
                                                                                    fmt.Println(v155, v6)
                                                                                default:
                                                                                    if true {
                                                                                        fmt.Println(fi144)
                                                                                        if (122 < 114) {
                                                                                            if ((fi156 ^ fi144) <= 2147483647) {
                                                                                                if (!false) {
                                                                                                    var v189 int = 777
                                                                                                    _ = v189
                                                                                                }
                                                                                            }
                                                                                            var v190 uint64 = v6
                                                                                            _ = v190
                                                                                        } else {
                                                                                            switch int(v5) {
                                                                                            case 1:
                                                                                                var v191 uint = 2
                                                                                                _ = v191
                                                                                            case -9223372036854775808:
                                                                                                fmt.Println(fi156, v137)
                                                                                            case -291:
                                                                                                var v192 uint = 766
                                                                                                _ = v192
                                                                                            }
                                                                                            var v193 int64 = -9223372036854775808
                                                                                            _ = v193
                                                                                        }
                                                                                    } else {
                                                                                        v3[2147483648] = 278
                                                                                    }
                                                                                }
                                                                            case "":
                                                                                v194 := GenS2{F0: genFn1(2147483647, ("" + "a"), v5)}
                                                                                _ = v194
                                                                            default:
                                                                                v195 := GenS11{F0: 42}
                                                                                _ = v195
                                                                            }
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                        }
                                                                    }
                                                                default:
                                                                    if (float64(v134) <= float64((v134 + -2.25))) {
                                                                        fmt.Println(v134, v6, v103)
                                                                    }
                                                                }
                                                                _ = fi156
                                                            }
                                                            _ = v155
                                                        } else {
                                                            v6 -= (v6 & (v6 ^ (v6 + (v6 - v6))))
                                                            v196 := GenS3{F0: fi144, F1: (^(v5)), F2: 4294967295}
                                                            _ = v196
                                                        }
                                                    }
                                                }
                                            }
                                            for fi197 := 0; fi197 < 1; fi197++ {
                                                v6 = (v6 + v6)
                                                v137 = (1.5 + 3.14159)
                                                switch (v133 + v133) {
                                                case "foo bar":
                                                    fmt.Println(fi144)
                                                case "a":
                                                    v5 = (v5 >> ((1) & 63))
                                                case "\n\t":
                                                    v6 = 725
                                                }
                                                _ = fi197
                                            }
                                            for fi198 := 0; fi198 < 14; fi198++ {
                                                switch v133 {
                                                case "":
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                                v103 = 133
                                                _ = fi198
                                            }
                                            _ = fi144
                                        }
                                        v4 = append(v4, v133)
                                    }
                                    v103 += v103
                                } else {
                                    if false {
                                        fmt.Println(v6, v6)
                                        fmt.Println(v133)
                                    } else {
                                        if ((v5 << ((genFn0()) & 63)) <= -1) {
                                            for fi199 := 0; fi199 < 9; fi199++ {
                                                if true {
                                                    v134 = 0.1
                                                    v200 := GenS5{F0: 0, F1: v133}
                                                    _ = v200
                                                }
                                                var v201 int16 = (v5 ^ (v5 % ((-70) | 1)))
                                                _ = fi199
                                                _ = v201
                                            }
                                        }
                                    }
                                    switch v133 {
                                    case "a":
                                        v2 = append(v2, 0)
                                    }
                                }
                                _ = v141
                            }
                        }
                    case 32767:
                        var v202 uint32 = 685
                        _ = v202
                    }
                    _ = v137
                }
            } else {
                fmt.Println(v133)
            }
        }
        _ = v134
    } else {
        v103 = 151
        v103 = int32(v5)
        if false {
            v2 = append(v2, 65535)
            for fi203 := 0; fi203 < 18; fi203++ {
                fmt.Println(v6)
                var v204 bool = (v133 <= "a")
                var v205 float32 = 1.5
                _ = fi203
                _ = v204
                _ = v205
            }
            v206 := GenS10{F0: 0, F1: -3, F2: (-(v103))}
            _ = v206
        }
    }
    switch -130 {
    case 1:
        if (!(1 <= -9223372036854775808)) {
            for fi207 := 0; fi207 < 15; fi207++ {
                v208 := GenS13{F0: (fi207 | (fi207 & fi207)), F1: 1}
                _ = fi207
                _ = v208
            }
            v5 += (v5 & (v5 - (v5 + (v5 >> ((65535) & 63)))))
        }
    }
    var v209 uint16 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v103:", v103)
    fmt.Println("v133:", v133)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v209:", v209)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v103
    _ = v133
    _ = v2
    _ = v209
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
