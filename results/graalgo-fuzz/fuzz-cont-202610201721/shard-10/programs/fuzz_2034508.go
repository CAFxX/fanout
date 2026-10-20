package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS6 struct { F0 bool; F1 int }
type GenS7 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS8 struct { F0 uint32; F1 uint }
type GenS9 struct { F0 uint64; F1 string }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int16; F1 int; F2 uint32 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint32; F1 uint; F2 int64 }
type GenS14 struct { F0 int }
type GenS15 struct { F0 uint; F1 bool; F2 uint16 }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 bool; F1 uint16; F2 int16 }
type GenS18 struct { F0 uint32; F1 uint64; F2 uint }
type GenS19 struct { F0 int; F1 string }
type GenS20 struct { F0 uint; F1 uint32; F2 bool }
func genFn0(p0 int, p1 int64, p2 string) int8 {
    p0 -= p0
    p1 = (p1 << ((4294967295) & 63))
    var v5 float32 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 100
}
func genFn1() uint {
    switch -9223372036854775808 {
    case -1:
        var v6 string = "foo bar"
        _ = v6
    case 398:
        if true {
            for fi7 := 0; fi7 < 8; fi7++ {
                var v8 float32 = 0.1
                v9 := GenS1{F0: 18, F1: 255, F2: ""}
                _ = fi7
                _ = v8
                _ = v9
            }
            if (631 <= 1) {
                if ((-1 <= -332) || (((false && ((22 == 0) && (115 <= 65535))) || false) || (539 <= 42))) {
                    var v10 int = -2147483648
                    _ = v10
                } else {
                    v11 := GenS1{F0: 1, F1: 4294967295, F2: "\n\t"}
                    v12 := GenS2{F0: 210}
                    var v13 uint8 = 205
                    _ = v11
                    _ = v12
                    _ = v13
                }
                for fi14 := 0; fi14 < 3; fi14++ {
                    var v15 bool = (1 != 0)
                    fmt.Println(v15, fi14)
                    for fi16 := 0; fi16 < 5; fi16++ {
                        var v17 bool = (v15 && ((!(42 < 1)) && v15))
                        var v18 int = 536
                        v19 := GenS2{F0: 1}
                        _ = fi16
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    _ = fi14
                    _ = v15
                }
                if false {
                    v20 := GenS3{F0: 1, F1: true}
                    v21 := []int{}
                    _ = v20
                    _ = v21
                } else {
                    v22 := []int{32767, 2147483647, 2147483647, 4294967295}
                    v23 := GenS3{F0: 127, F1: (255 < 32767)}
                    _ = v22
                    _ = v23
                }
            }
            if (42 > 173) {
                for fi24 := 0; fi24 < 10; fi24++ {
                    for fi25 := 0; fi25 < 11; fi25++ {
                        v26 := []int{(fi24 ^ (fi25 + 127)), 32768, 546}
                        if ((false && true) || ("foo bar" == "a")) {
                            if false {
                                var v27 bool = true
                                var v28 int64 = 94
                                v29 := GenS3{F0: -1, F1: true}
                                _ = v27
                                _ = v28
                                _ = v29
                            }
                            v30 := GenS2{F0: 1}
                            _ = v30
                        } else {
                            var v31 uint = 44
                            v26 = append(v26, (fi25 | (fi25 << ((18446744073709551615) & 63))))
                            var v32 float32 = 1e10
                            _ = v31
                            _ = v32
                        }
                        v33 := GenS4{F0: 2}
                        _ = fi25
                        _ = v26
                        _ = v33
                    }
                    _ = fi24
                }
                v34 := []int{-745, 127}
                _ = v34
            }
        }
    case -9223372036854775808:
        if (!true) {
            if (350 != 207) {
                switch -2147483648 {
                case 785:
                    switch 0 {
                    case 4294967295:
                        if (!false) {
                            if true {
                                var v35 int16 = 32767
                                if (true && (true || (-887 != 9223372036854775807))) {
                                    var v36 int16 = -879
                                    if false {
                                        for fi37 := 0; fi37 < 5; fi37++ {
                                            if false {
                                                fmt.Println(fi37)
                                            }
                                            if ((v35 < (v36 % ((0) | 1))) || true) {
                                                v38 := GenS4{F0: 42}
                                                v39 := GenS5{F0: 65535, F1: 0, F2: 0}
                                                _ = v38
                                                _ = v39
                                            } else {
                                                if (!false) {
                                                    fmt.Println(v35, v36)
                                                    switch 4294967295 {
                                                    case -540:
                                                        v36 = (v36 - -32768)
                                                    case 32768:
                                                        v40 := []int{fi37, fi37, 128, fi37}
                                                        _ = v40
                                                    default:
                                                        for fi41 := 0; fi41 < 14; fi41++ {
                                                            v35 = (v35 - v35)
                                                            v42 := []int{-131}
                                                            switch (fi41 % (((fi37 / (((^(fi41))) | 1))) | 1)) {
                                                            case 4294967295:
                                                                v42 = append(v42, (fi41 + fi41))
                                                            case 2147483647:
                                                                if len(v42) >= 2 {
                                                                    v42 = v42[1:len(v42)]
                                                                }
                                                            }
                                                            fmt.Println(fi41, v35, v35)
                                                            var v43 uint = 42
                                                            v44 := GenS6{F0: true, F1: (fi37 * (fi41 & 9223372036854775807))}
                                                            _ = fi41
                                                            _ = v42
                                                            _ = v43
                                                            _ = v44
                                                        }
                                                    }
                                                    for fi45 := 0; fi45 < 13; fi45++ {
                                                        v46 := GenS3{F0: -128, F1: true}
                                                        var v47 string = "hello"
                                                        _ = fi45
                                                        _ = v46
                                                        _ = v47
                                                    }
                                                } else {
                                                    for fi48 := 0; fi48 < 5; fi48++ {
                                                        var v49 bool = (false || true)
                                                        v50 := GenS7{F0: 42, F1: (!(v35 != 328)), F2: 736}
                                                        var v51 int8 = -96
                                                        _ = fi48
                                                        _ = v49
                                                        _ = v50
                                                        _ = v51
                                                    }
                                                }
                                                v52 := GenS6{F0: (42 >= 46), F1: (fi37 ^ fi37)}
                                                var v53 int8 = 100
                                                _ = v52
                                                _ = v53
                                            }
                                            if true {
                                                v54 := GenS8{F0: 2147483648, F1: 42}
                                                switch "hello" {
                                                case "":
                                                    var v55 int64 = 42
                                                    _ = v55
                                                }
                                                var v56 uint8 = 128
                                                _ = v54
                                                _ = v56
                                            } else {
                                                var v57 int8 = -1
                                                switch "hello" {
                                                case "hello":
                                                    v58 := GenS9{F0: 0, F1: "\n\t"}
                                                    _ = v58
                                                default:
                                                    switch fi37 {
                                                    case -273:
                                                        var v59 bool = (489 < 42)
                                                        _ = v59
                                                    case 2147483647:
                                                        switch 1 {
                                                        case -1:
                                                            fmt.Println(fi37)
                                                        case 32768:
                                                            v60 := GenS2{F0: 128}
                                                            _ = v60
                                                        case 747:
                                                            v61 := GenS2{F0: 0}
                                                            _ = v61
                                                        }
                                                    case 10:
                                                        if (65535 < 575) {
                                                            if false {
                                                                v62 := []int{127, -283, fi37, (fi37 + (fi37 & fi37))}
                                                                _ = v62
                                                            } else {
                                                                for fi63 := 0; fi63 < 12; fi63++ {
                                                                    var v64 uint16 = 1
                                                                    var v65 uint8 = 2
                                                                    _ = fi63
                                                                    _ = v64
                                                                    _ = v65
                                                                }
                                                                v66 := []int{(fi37 / ((32767) | 1)), (-(fi37))}
                                                                _ = v66
                                                            }
                                                            var v67 uint8 = 6
                                                            v68 := GenS10{F0: ("foo bar" + (("hello" + "\n\t") + "\n\t"))}
                                                            _ = v67
                                                            _ = v68
                                                        }
                                                    default:
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            switch ("foo bar" + "a") {
                                                            case "\n\t":
                                                                for fi69 := 0; fi69 < 10; fi69++ {
                                                                    switch (("foo bar" + "") + "a") {
                                                                    case "\n\t":
                                                                        var v70 uint8 = 1
                                                                        _ = v70
                                                                    case "foo bar":
                                                                        v36 = (v35 | (v35 * v35))
                                                                    case "":
                                                                        if true {
                                                                            v35 += v36
                                                                            if (!true) {
                                                                                fmt.Println(fi37, fi69, fi69)
                                                                            } else {
                                                                                v71 := GenS9{F0: 868, F1: "a"}
                                                                                _ = v71
                                                                            }
                                                                            v72 := GenS1{F0: (^(v57)), F1: 607, F2: "\n\t"}
                                                                            _ = v72
                                                                        } else {
                                                                            var v73 int = fi69
                                                                            var v74 uint8 = 192
                                                                            _ = v73
                                                                            _ = v74
                                                                        }
                                                                    default:
                                                                        if true {
                                                                            v36 -= 32767
                                                                            v75 := GenS11{F0: v36, F1: fi69, F2: 4294967295}
                                                                            _ = v75
                                                                        } else {
                                                                            var v76 bool = (!false)
                                                                            v77 := []int{127, int(v35), (fi37 & (fi69 | 141))}
                                                                            v36 = (-(v36))
                                                                            _ = len(v77)
                                                                            if (v76 && (-66 != 42)) {
                                                                                if (105 != 32767) {
                                                                                    var v78 int16 = v35
                                                                                    _ = v78
                                                                                }
                                                                                if (false || (832 > 18446744073709551615)) {
                                                                                    fmt.Println(v76)
                                                                                } else {
                                                                                    if true {
                                                                                        var v79 uint64 = 176
                                                                                        _ = v79
                                                                                    } else {
                                                                                        v80 := GenS11{F0: v35, F1: (fi37 % (((fi69 + fi37)) | 1)), F2: 815}
                                                                                        if len(v77) > 0 {
                                                                                            v77[0] = (fi37 * fi69)
                                                                                        }
                                                                                        v81 := GenS12{F0: 466}
                                                                                        _ = v80
                                                                                        _ = v81
                                                                                    }
                                                                                    v82 := GenS10{F0: "foo bar"}
                                                                                    _ = v82
                                                                                }
                                                                            }
                                                                            if len(v77) > 0 {
                                                                                v77[0] = (fi37 / ((662) | 1))
                                                                            }
                                                                            _ = v76
                                                                            _ = v77
                                                                        }
                                                                    }
                                                                    v83 := []int{fi37, fi69, 0}
                                                                    _ = fi69
                                                                    _ = v83
                                                                }
                                                            }
                                                        case "a":
                                                            v35 += (^(v35))
                                                        case "":
                                                            fmt.Println(v57, v57, v35)
                                                        }
                                                    }
                                                }
                                                _ = v57
                                            }
                                            _ = fi37
                                        }
                                        v84 := []int{int(v35), 128, -820, -2147483648}
                                        _ = v84
                                    }
                                    _ = v36
                                }
                                _ = v35
                            } else {
                                switch ("foo bar" + "") {
                                case "a":
                                    var v85 string = "foo bar"
                                    _ = v85
                                case "hello":
                                    for fi86 := 0; fi86 < 15; fi86++ {
                                        var v87 string = "a"
                                        _ = fi86
                                        _ = v87
                                    }
                                default:
                                    var v88 int16 = -400
                                    _ = v88
                                }
                            }
                            var v89 int64 = -339
                            switch "\n\t" {
                            case "a":
                                var v90 uint64 = 18446744073709551615
                                _ = v90
                            }
                            _ = v89
                        }
                    }
                case -9223372036854775808:
                    var v91 string = "a"
                    _ = v91
                case 127:
                    switch 534 {
                    case -2147483648:
                        if (65535 != 65535) {
                            if false {
                                v92 := []int{42, 0, -621}
                                var v93 uint = 87
                                _ = v92
                                _ = v93
                            } else {
                                for fi94 := 0; fi94 < 11; fi94++ {
                                    v95 := GenS8{F0: 2147483648, F1: 42}
                                    for fi96 := 0; fi96 < 9; fi96++ {
                                        var v97 uint32 = 0
                                        if (255 <= 42) {
                                            if false {
                                                var v98 int16 = 0
                                                _ = v98
                                            }
                                        }
                                        switch "foo bar" {
                                        case "foo bar":
                                            if ((((42 != 695) && (140 <= 1)) && (9223372036854775807 > 0)) || true) {
                                                switch (fi94 | (fi94 / (((fi94 ^ 2147483647)) | 1))) {
                                                case 9223372036854775807:
                                                    v99 := []int{fi96, fi96, (fi96 ^ fi94)}
                                                    _ = v99
                                                case 32767:
                                                    v97 = (v97 / ((v97) | 1))
                                                }
                                                v100 := []int{}
                                                _ = v100
                                            } else {
                                                if false {
                                                    v97 = (v97 | v97)
                                                    var v101 int32 = 0
                                                    var v102 int32 = v101
                                                    _ = v101
                                                    _ = v102
                                                } else {
                                                    v103 := []int{(fi96 % ((865) | 1)), fi96, (fi94 * fi94)}
                                                    v97 += v97
                                                    v104 := GenS12{F0: 2}
                                                    _ = v103
                                                    _ = v104
                                                }
                                            }
                                        }
                                        _ = fi96
                                        _ = v97
                                    }
                                    _ = fi94
                                    _ = v95
                                }
                            }
                            v105 := []int{32767, 32768}
                            var v106 int = -437
                            _ = v105
                            _ = v106
                        }
                    default:
                        v107 := GenS10{F0: (("hello" + ("hello" + "\n\t")) + "\n\t")}
                        _ = v107
                    }
                }
                if (2 <= 254) {
                    var v108 uint = 18446744073709551615
                    var v109 int32 = 42
                    _ = v108
                    _ = v109
                } else {
                    for fi110 := 0; fi110 < 18; fi110++ {
                        v111 := GenS13{F0: 257, F1: 339, F2: 9223372036854775807}
                        var v112 int8 = 0
                        if false {
                            v112 = genFn0(int(v112), -1, (("hello" + "hello") + "\n\t"))
                            v112 = v112
                        } else {
                            switch fi110 {
                            case 855:
                                v112 += -49
                            case -1:
                                if ((!((("hello" + "\n\t") + ((("" + "foo bar") + "\n\t") + "a")) <= ("foo bar" + ("foo bar" + "")))) || false) {
                                    var v113 int16 = -32768
                                    v112 = (^(v112))
                                    _ = v113
                                }
                            }
                        }
                        _ = fi110
                        _ = v111
                        _ = v112
                    }
                }
            } else {
                v114 := GenS14{F0: 556}
                _ = v114
            }
            switch "" {
            case "":
                v115 := GenS14{F0: -910}
                _ = v115
            case "a":
                for fi116 := 0; fi116 < 11; fi116++ {
                    var v117 int8 = -1
                    if true {
                        var v118 int8 = (v117 - (v117 - (v117 + genFn0(fi116, 42, "a"))))
                        if true {
                            v118 = (v117 & genFn0(fi116, 9223372036854775807, "hello"))
                            var v119 bool = true
                            switch "hello" {
                            case "foo bar":
                                if (127 > 42) {
                                    v119 = (!(v118 >= (v118 & (v118 % (((v118 * 91)) | 1)))))
                                }
                            }
                            _ = v119
                        }
                        _ = v118
                    } else {
                        v120 := []int{}
                        for fi121 := 0; fi121 < 13; fi121++ {
                            var v122 int8 = (^(v117))
                            _ = fi121
                            _ = v122
                        }
                        switch (("foo bar" + "hello") + "") {
                        case "foo bar":
                            var v123 int32 = -2147483648
                            _ = v123
                        case "\n\t":
                            switch ((("a" + ("a" + "foo bar")) + "a") + "hello") {
                            case "\n\t":
                                v124 := GenS12{F0: 825}
                                _ = v124
                            case "":
                                v125 := GenS7{F0: 2, F1: false, F2: 42}
                                _ = v125
                            }
                        case "":
                            switch -2147483648 {
                            case 21:
                                v126 := GenS15{F0: 2, F1: true, F2: 2}
                                _ = v126
                            case -1:
                                v127 := GenS11{F0: -64, F1: 1, F2: 42}
                                _ = v127
                            default:
                                for fi128 := 0; fi128 < 11; fi128++ {
                                    var v129 int16 = -214
                                    _ = fi128
                                    _ = v129
                                }
                            }
                        default:
                            v120 = append(v120, 9223372036854775807)
                        }
                        _ = v120
                    }
                    var v130 int = fi116
                    _ = fi116
                    _ = v117
                    _ = v130
                }
            default:
                var v131 uint16 = 42
                _ = v131
            }
            for fi132 := 0; fi132 < 13; fi132++ {
                v133 := GenS8{F0: 42, F1: 0}
                _ = fi132
                _ = v133
            }
        }
    }
    if (true || false) {
        v134 := GenS16{F0: 100}
        var v135 int16 = -1
        _ = v134
        _ = v135
    } else {
        var v136 uint32 = 4294967295
        _ = v136
    }
    for fi137 := 0; fi137 < 12; fi137++ {
        var v138 string = "a"
        switch v138 {
        case "\n\t":
            var v139 uint64 = 18446744073709551615
            _ = v139
        case "a":
            v140 := []int{127, fi137, (fi137 % (((^(fi137))) | 1)), fi137}
            _ = v140
        case "":
            if ((1e10 / 0.1) == 0.1) {
                switch 778 {
                case 832:
                    var v141 float32 = -2.25
                    _ = v141
                }
                v138 = v138
                v138 = v138
            }
        default:
            v138 = v138
        }
        if false {
            fmt.Println(v138, v138)
        }
        _ = fi137
        _ = v138
    }
    var v142 int8 = 0
    _ = v142
    return 1
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[string]int)
    var v3 bool = ("hello" >= ("foo bar" + ("a" + "")))
    var v4 uint64 = 337
    if v3 {
        var v143 uint32 = 583
        switch "\n\t" {
        case "":
            for fi144 := 0; fi144 < 7; fi144++ {
                var v145 uint32 = v143
                v146 := []int{4294967295, fi144, fi144, fi144}
                for fi147 := 0; fi147 < 8; fi147++ {
                    if v3 {
                        var v148 uint16 = (v1 ^ uint16(v4))
                        delete(v2, "hello")
                        if v3 {
                            v143 -= (v145 + (v145 | (^(v145))))
                            var v149 int = fi144
                            _ = v149
                        } else {
                            v150 := GenS17{F0: false, F1: 2, F2: 1000}
                            _ = v150
                        }
                        _ = v148
                    } else {
                        var v151 int8 = 0
                        if ((247 > -404) && true) {
                            var v152 int8 = genFn0((fi147 ^ (-(fi147))), 9223372036854775807, "a")
                            var v153 string = ""
                            if (-922 != 9223372036854775807) {
                                if (42 < 97) {
                                    switch v153 {
                                    case "foo bar":
                                        var v154 string = (v153 + ("hello" + v153))
                                        _ = v154
                                    }
                                } else {
                                    if ((v143 | (v143 * v143)) == v143) {
                                        if len(v146) >= 2 {
                                            v146 = v146[1:len(v146)]
                                        }
                                        var v155 int64 = 42
                                        if len(v146) > 0 {
                                            v146[0] = (fi147 - (fi144 % ((int(v4)) | 1)))
                                        }
                                        _ = v155
                                    } else {
                                        if (607 == 487) {
                                            v156 := GenS9{F0: (^(v4)), F1: v153}
                                            _ = v156
                                        }
                                    }
                                }
                            }
                            _ = v152
                            _ = v153
                        }
                        v145 += v145
                        _ = v151
                    }
                    if v3 {
                        v145 = (^(v143))
                        switch fi144 {
                        case 2147483647:
                            if len(v146) >= 2 {
                                v146 = v146[1:len(v146)]
                            }
                        }
                        v4 -= 895
                    }
                    switch (("a" + "foo bar") + "foo bar") {
                    case "a":
                        if false {
                            var v157 int32 = -792
                            fmt.Println(fi144)
                            v158 := GenS18{F0: (v145 + 2147483647), F1: uint64(v143), F2: 4294967295}
                            _ = v157
                            _ = v158
                        } else {
                            v159 := GenS17{F0: false, F1: v1, F2: 0}
                            _ = v159
                        }
                    default:
                        var v160 int8 = -128
                        if len(v146) >= 2 {
                            v146 = v146[1:len(v146)]
                        }
                        _ = v160
                    }
                    _ = fi147
                }
                _ = fi144
                _ = v145
                _ = v146
            }
        case "foo bar":
            fmt.Println(v143, v143, v143)
        }
        _ = v143
    } else {
        v161 := []int{-9223372036854775808, 127, -2147483648, 772}
        fmt.Println(v1)
        v162 := GenS10{F0: (("hello" + ("a" + "a")) + ("a" + ""))}
        _ = v161
        _ = v162
    }
    v163 := GenS10{F0: "hello"}
    var v164 uint16 = (v1 + v1)
    if true {
        var v165 int = -2147483648
        if (!true) {
            v166 := []int{(^(v165)), v165, 927, 128}
            v167 := GenS19{F0: v165, F1: "\n\t"}
            _ = v166
            _ = v167
        } else {
            var v168 uint8 = 255
            _ = v168
        }
        v4 += v4
        _ = v165
    } else {
        var v169 int16 = 32767
        if (!v3) {
            v170 := []int{-1, 4294967295, 9223372036854775807, 128}
            v171 := GenS20{F0: 706, F1: 588, F2: ((v4 + 18446744073709551615) < v4)}
            v172 := GenS12{F0: genFn1()}
            _ = v170
            _ = v171
            _ = v172
        } else {
            var v173 uint64 = v4
            _ = v173
        }
        _ = v169
    }
    v174 := GenS11{F0: -1, F1: 127, F2: 0}
    for fi175 := 0; fi175 < 9; fi175++ {
        v3 = v3
        for fi176 := 0; fi176 < 8; fi176++ {
            v3 = false
            v177 := GenS11{F0: -779, F1: fi175, F2: 42}
            var v178 float64 = 0.0
            _ = fi176
            _ = v177
            _ = v178
        }
        if (!v3) {
            v179 := []int{fi175}
            fmt.Println(v1, v3)
            _ = v179
        } else {
            var v180 int32 = 43
            for fi181 := 0; fi181 < 10; fi181++ {
                fmt.Println(v164, fi181)
                var v182 string = ""
                _ = fi181
                _ = v182
            }
            _ = v180
        }
        _ = fi175
    }
    for fi183 := 0; fi183 < 16; fi183++ {
        v184 := GenS11{F0: 32767, F1: (fi183 | fi183), F2: 65535}
        _ = fi183
        _ = v184
    }
    for fi185 := 0; fi185 < 15; fi185++ {
        v186 := GenS16{F0: 127}
        _ = fi185
        _ = v186
    }
    fmt.Println(v3, v3, v3)
    if v3 {
        v187 := GenS12{F0: 592}
        _ = v187
    } else {
        v164 -= (v164 / ((v164) | 1))
        fmt.Println(v164, v1)
    }
    var v188 uint64 = v4
    v189 := GenS19{F0: 0, F1: ("hello" + "a")}
    fmt.Println("v1:", v1)
    fmt.Println("v164:", v164)
    fmt.Println("v188:", v188)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v163
    _ = v164
    _ = v174
    _ = v188
    _ = v189
    _ = v2
    _ = v3
    _ = v4
}
