package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 int; F1 int32; F2 int16 }
type GenS6 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int64; F1 uint }
type GenS10 struct { F0 int64; F1 int32 }
type GenS11 struct { F0 int; F1 uint32; F2 int }
type GenS12 struct { F0 bool; F1 uint; F2 uint32 }
type GenS13 struct { F0 int8; F1 string }
type GenS14 struct { F0 string; F1 int32; F2 uint8 }
type GenS15 struct { F0 bool; F1 string }
type GenS16 struct { F0 uint64; F1 int64 }
type GenS17 struct { F0 int8 }
type GenS18 struct { F0 uint16 }
type GenS19 struct { F0 uint; F1 int32 }
type GenS20 struct { F0 int16; F1 uint8; F2 int }
type GenS21 struct { F0 int8 }
type GenS22 struct { F0 uint64 }
type GenS23 struct { F0 string; F1 uint64 }
type GenS24 struct { F0 bool; F1 uint32; F2 string }
type GenS25 struct { F0 uint8; F1 uint8; F2 string }
type GenS26 struct { F0 int32; F1 uint }
func genFn0(p0 uint, p1 bool) uint32 {
    v5 := GenS1{F0: 255}
    switch 35 {
    case 32767:
        var v6 string = ""
        _ = v6
    }
    if p1 {
        fmt.Println(p0, p1)
        v7 := []int{4294967295, 128, int(p0)}
        var v8 int64 = 1
        _ = v7
        _ = v8
    } else {
        var v9 uint32 = 2147483647
        fmt.Println(v9)
        for fi10 := 0; fi10 < 11; fi10++ {
            if p1 {
                switch -318 {
                case -9223372036854775808:
                    var v11 uint8 = 127
                    _ = v11
                case 741:
                    switch "foo bar" {
                    case "hello":
                        var v12 string = "foo bar"
                        _ = v12
                    case "a":
                        fmt.Println(v9, fi10)
                    default:
                        v13 := GenS1{F0: 1}
                        _ = v13
                    }
                default:
                    p1 = ((0 >= 255) || false)
                }
            }
            _ = fi10
        }
        _ = v9
    }
    if p1 {
        switch 667 {
        case 0:
            var v14 int = -827
            _ = v14
        case 1:
            for fi15 := 0; fi15 < 17; fi15++ {
                var v16 string = ("\n\t" + ("a" + ("" + ((("a" + "\n\t") + "a") + ("a" + (("" + "a") + ""))))))
                switch ("foo bar" + v16) {
                case "\n\t":
                    if true {
                        var v17 int64 = -1
                        v18 := []int{127, -2147483648, 127, 283}
                        var v19 int64 = (v17 / (((v17 % ((v17) | 1))) | 1))
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                default:
                    for fi20 := 0; fi20 < 19; fi20++ {
                        if p1 {
                            if ((((((1e10 + (((1e10 - -2.25) - float32((1.5 + -2.25))) - ((1e10 - float32(3.14159)) / float32(((0.1 - 0.0) + ((0.1 - float64((1.5 - -2.25))) / 1e10)))))) * -2.25) < (1e10 + 1.5)) && true) || (647 != -1)) || p1) {
                                v21 := GenS2{F0: -309, F1: 42}
                                fmt.Println(fi15, v16, p0)
                                _ = v21
                            } else {
                                if (2147483647 >= 65535) {
                                    var v22 int8 = 56
                                    var v23 uint64 = 4294967295
                                    _ = v22
                                    _ = v23
                                }
                                var v24 int = (fi20 + -463)
                                v24 -= 4294967295
                                _ = v24
                            }
                        } else {
                            v25 := GenS3{F0: 393}
                            if true {
                                if p1 {
                                    var v26 uint = 65535
                                    if p1 {
                                        if (!p1) {
                                            fmt.Println(fi20, v16)
                                            switch (v16 + (("foo bar" + "a") + v16)) {
                                            case "":
                                                if false {
                                                    v26 = 404
                                                }
                                            case "a":
                                                fmt.Println(p1)
                                            case "hello":
                                                v16 = v16
                                            default:
                                                fmt.Println(v16)
                                            }
                                            fmt.Println(p0, v26, v26)
                                            p1 = p1
                                            v27 := GenS4{F0: ((("" + (("\n\t" + (("\n\t" + ((v16 + "a") + v16)) + ((v16 + ("foo bar" + ("" + "a"))) + v16))) + (v16 + ((v16 + v16) + (("hello" + v16) + v16))))) + "a") + (v16 + v16)), F1: 42}
                                            _ = v27
                                        }
                                        var v28 float32 = -2.25
                                        p1 = (128 > 29)
                                        _ = v28
                                    } else {
                                        var v29 float32 = 0.1
                                        v30 := GenS5{F0: fi15, F1: 0, F2: int16(p0)}
                                        _ = v29
                                        _ = v30
                                    }
                                    _ = v26
                                } else {
                                    v31 := GenS1{F0: 32767}
                                    v32 := []int{}
                                    _ = v31
                                    _ = v32
                                }
                            } else {
                                p0 += (p0 & 243)
                                p0 = 174
                                switch ((v16 + "hello") + "a") {
                                case "a":
                                    v33 := GenS6{F0: 0, F1: 42, F2: 756}
                                    _ = v33
                                case "hello":
                                    var v34 int = 963
                                    _ = v34
                                case "foo bar":
                                    var v35 string = ((v16 + ("hello" + v16)) + "hello")
                                    _ = v35
                                default:
                                    var v36 float32 = float32(((float64(float32(0.0)) / 1.5) * (-2.25 * 0.0)))
                                    _ = v36
                                }
                                v37 := []int{}
                                var v38 int64 = -1
                                if (true || (204 < 42)) {
                                    v38 = v38
                                }
                                var v39 string = v16
                                _ = v37
                                _ = v38
                                _ = v39
                            }
                            v40 := []int{127, (fi20 ^ (fi15 % ((fi15) | 1))), 0}
                            v41 := GenS4{F0: ((v16 + "\n\t") + v16), F1: 281}
                            var v42 int16 = 1
                            _ = v25
                            _ = v40
                            _ = v41
                            _ = v42
                        }
                        _ = fi20
                    }
                }
                v43 := GenS4{F0: "\n\t", F1: 0}
                _ = fi15
                _ = v16
                _ = v43
            }
        }
        if ((!false) && false) {
            p1 = (!((((p1 && (231 <= 98)) || false) && false) && (65535 <= 191)))
            switch 964 {
            case 9223372036854775807:
                if (2147483647 <= 42) {
                    p0 -= (p0 & (p0 * 3))
                    v44 := GenS7{F0: 858}
                    _ = v44
                }
            case 128:
                switch "\n\t" {
                case "hello":
                    v45 := GenS7{F0: -2147483648}
                    _ = v45
                }
            }
        } else {
            v46 := []int{32767, -9223372036854775808, 291, 9223372036854775807}
            v46 = append(v46, 32767)
            if (32767 <= 0) {
                p0 = p0
                fmt.Println(p0)
                p0 = (p0 / ((849) | 1))
            }
            _ = v46
        }
    } else {
        switch ("\n\t" + "hello") {
        case "":
            p1 = (36 < 127)
        default:
            var v47 int8 = 0
            _ = v47
        }
        for fi48 := 0; fi48 < 4; fi48++ {
            fmt.Println(p0, p1)
            v49 := []int{(fi48 - (fi48 % (((fi48 << ((p0) & 63))) | 1))), fi48, (fi48 + 32767), fi48}
            _ = fi48
            _ = v49
        }
        v50 := []int{382, 574}
        _ = v50
    }
    v51 := []int{42, -1, 127, 310}
    _ = p0
    _ = p1
    _ = v5
    _ = v51
    return 556
}
func genFn1(p0 int8) int {
    for fi52 := 0; fi52 < 5; fi52++ {
        if true {
            var v53 uint = 4294967295
            _ = v53
        } else {
            var v54 int64 = -1
            v55 := GenS7{F0: 2147483647}
            if ((v54 % ((v54) | 1)) == (v54 * v54)) {
                v56 := GenS6{F0: 1, F1: 3, F2: 0}
                var v57 int64 = (v54 ^ -1)
                v58 := GenS8{F0: 65535}
                _ = v56
                _ = v57
                _ = v58
            }
            _ = v54
            _ = v55
        }
        for fi59 := 0; fi59 < 9; fi59++ {
            switch (fi52 - (fi52 ^ 9223372036854775807)) {
            case 1:
                var v60 int8 = p0
                _ = v60
            case -660:
                v61 := GenS9{F0: 677, F1: 0}
                _ = v61
            case 42:
                p0 += (p0 & (-(p0)))
            default:
                switch fi52 {
                case 9223372036854775807:
                    v62 := []int{int(p0), fi59}
                    var v63 bool = true
                    if (36 != 0) {
                        v64 := GenS10{F0: 9223372036854775807, F1: -1}
                        var v65 uint8 = 2
                        p0 = p0
                        var v66 uint32 = 2147483647
                        _ = v64
                        _ = v65
                        _ = v66
                    } else {
                        v63 = (74 < 1)
                    }
                    _ = v62
                    _ = v63
                case 330:
                    p0 = 74
                }
            }
            var v67 float32 = 1.5
            if true {
                if (157 >= 187) {
                    var v68 uint32 = genFn0(65535, true)
                    v68 = (v68 * v68)
                    _ = v68
                } else {
                    if false {
                        if (((("\n\t" + "") + "a") + "hello") == ("foo bar" + "hello")) {
                            if (p0 == 9) {
                                v69 := GenS2{F0: 1, F1: 343}
                                _ = v69
                            } else {
                                if false {
                                    v67 = v67
                                    switch "foo bar" {
                                    case "":
                                        v70 := []int{(fi52 & (-(fi52))), fi59}
                                        p0 = -95
                                        v71 := GenS9{F0: 790, F1: 625}
                                        _ = v70
                                        _ = v71
                                    default:
                                        v72 := []int{}
                                        _ = v72
                                    }
                                }
                                if false {
                                    v73 := GenS11{F0: 42, F1: 2, F2: 9223372036854775807}
                                    if (981 == 679) {
                                        v74 := []int{}
                                        _ = v74
                                    }
                                    v75 := []int{fi52, (fi52 & fi59), fi59, (^(fi59))}
                                    _ = v73
                                    _ = v75
                                }
                                var v76 uint8 = 120
                                _ = v76
                            }
                        } else {
                            if (!false) {
                                v67 += (v67 + v67)
                                v77 := GenS7{F0: -1}
                                switch fi52 {
                                case -9223372036854775808:
                                    if (!(!true)) {
                                        var v78 uint64 = 1
                                        v79 := GenS9{F0: 42, F1: 657}
                                        _ = v78
                                        _ = v79
                                    } else {
                                        switch "\n\t" {
                                        case "a":
                                            v67 -= ((0.0 - 1e10) + float32(1e10))
                                        case "\n\t":
                                            v67 += v67
                                        }
                                    }
                                case -160:
                                    switch fi52 {
                                    case 4294967295:
                                        switch fi59 {
                                        case -9223372036854775808:
                                            var v80 uint64 = 1
                                            _ = v80
                                        case 661:
                                            var v81 string = "a"
                                            v81 = "a"
                                            _ = v81
                                        default:
                                            switch "a" {
                                            case "\n\t":
                                                v82 := []int{32768, fi59, fi52, fi59}
                                                var v83 uint = 169
                                                _ = v82
                                                _ = v83
                                            case "a":
                                                fmt.Println(v67, fi52)
                                            default:
                                                v84 := []int{fi59}
                                                _ = v84
                                            }
                                        }
                                    case -680:
                                        var v85 float32 = float32((0.0 + 1.5))
                                        _ = v85
                                    default:
                                        var v86 float64 = ((1.5 * 0.0) * 3.14159)
                                        _ = v86
                                    }
                                case -264:
                                    var v87 float64 = (0.0 + (1e10 + float64(float32(0.0))))
                                    _ = v87
                                }
                                v88 := GenS6{F0: 84, F1: 2, F2: 160}
                                v89 := []int{fi52, fi52}
                                if (!true) {
                                    var v90 string = ""
                                    fmt.Println(v90)
                                    switch v90 {
                                    case "":
                                        _ = len(v89)
                                    case "hello":
                                        if false {
                                            switch (v90 + "a") {
                                            case "\n\t":
                                                v89 = append(v89, fi52)
                                            case "":
                                                var v91 string = (v90 + ("" + (v90 + v90)))
                                                _ = v91
                                            case "a":
                                                fmt.Println(fi52, fi52, v90)
                                            }
                                            var v92 int16 = 1
                                            v93 := GenS12{F0: (-276 != -787), F1: 1, F2: 373}
                                            if (v90 <= ((v90 + v90) + v90)) {
                                                v90 = v90
                                                if (!false) {
                                                    v94 := GenS7{F0: -1}
                                                    _ = v94
                                                } else {
                                                    switch v90 {
                                                    case "hello":
                                                        fmt.Println(v90)
                                                    case "":
                                                        switch v90 {
                                                        case "":
                                                            _ = len(v89)
                                                        case "\n\t":
                                                            v92 = 1
                                                        default:
                                                            var v95 int64 = 70
                                                            _ = v95
                                                        }
                                                    case "\n\t":
                                                        switch fi59 {
                                                        case -969:
                                                            switch -9223372036854775808 {
                                                            case -920:
                                                                if true {
                                                                    if false {
                                                                        var v96 string = (("foo bar" + "\n\t") + (v90 + "hello"))
                                                                        _ = v96
                                                                    } else {
                                                                        var v97 uint64 = 1
                                                                        v92 += (v92 % (((v92 << ((159) & 63))) | 1))
                                                                        var v98 uint32 = 42
                                                                        _ = v97
                                                                        _ = v98
                                                                    }
                                                                }
                                                            case 2147483647:
                                                                if false {
                                                                    p0 += -105
                                                                    fmt.Println(fi52, p0, v67)
                                                                    fmt.Println(v90, v67, v92)
                                                                }
                                                            case 446:
                                                                if false {
                                                                    v99 := GenS13{F0: 100, F1: ((v90 + (v90 + v90)) + ((v90 + (((("" + "\n\t") + "a") + "\n\t") + "a")) + "hello"))}
                                                                    switch 1 {
                                                                    case 530:
                                                                        fmt.Println(v90, fi59, v92)
                                                                    case -1:
                                                                        var v100 uint = 4294967295
                                                                        _ = v100
                                                                    }
                                                                    var v101 uint = 845
                                                                    _ = v101
                                                                    _ = v99
                                                                }
                                                            }
                                                        case 32768:
                                                            _ = len(v89)
                                                        case -707:
                                                            if len(v89) >= 2 {
                                                                v89 = v89[1:len(v89)]
                                                            }
                                                        }
                                                    }
                                                    fmt.Println(p0)
                                                    fmt.Println(v92, v92)
                                                }
                                            }
                                            if len(v89) >= 2 {
                                                v89 = v89[1:len(v89)]
                                            }
                                            _ = v92
                                            _ = v93
                                        }
                                    }
                                    _ = v90
                                }
                                p0 += p0
                                var v102 bool = false
                                if len(v89) >= 2 {
                                    v89 = v89[1:len(v89)]
                                }
                                _ = v102
                                _ = v77
                                _ = v88
                                _ = v89
                            }
                            v67 = (3.14159 / (v67 / 0.1))
                            v103 := []int{(^(fi52)), 32768, (fi59 % ((0) | 1))}
                            v104 := GenS9{F0: -1, F1: 998}
                            var v105 uint64 = 18446744073709551615
                            var v106 uint64 = 0
                            _ = v103
                            _ = v104
                            _ = v105
                            _ = v106
                        }
                        var v107 int64 = 0
                        fmt.Println(fi52)
                        _ = v107
                    } else {
                        var v108 int = (^(fi52))
                        v109 := GenS14{F0: "a", F1: 42, F2: 0}
                        _ = v108
                        _ = v109
                    }
                    var v110 float64 = 1e10
                    _ = v110
                }
            } else {
                v111 := GenS15{F0: false, F1: "hello"}
                fmt.Println(v67)
                var v112 int8 = 100
                _ = v111
                _ = v112
            }
            _ = fi59
            _ = v67
        }
        var v113 int64 = -9223372036854775808
        _ = fi52
        _ = v113
    }
    p0 = (p0 << ((18446744073709551615) & 63))
    for fi114 := 0; fi114 < 5; fi114++ {
        var v115 uint32 = 65535
        _ = fi114
        _ = v115
    }
    v116 := GenS15{F0: ((0.0 + 3.14159) < 1.5), F1: ""}
    v117 := GenS14{F0: "hello", F1: 42, F2: 42}
    _ = p0
    _ = v116
    _ = v117
    return 32767
}

func main() {
    var v1 int = 4294967295
    v2 := make(map[string]bool)
    var v3 int32 = -2147483648
    var v4 int32 = (v3 >> ((65535) & 63))
    var v118 int32 = 2147483647
    v119 := GenS13{F0: 0, F1: "hello"}
    v120 := v2["hello"]
    if ((false && false) && (533 == 56)) {
        v121 := GenS3{F0: 490}
        if v120 {
            switch genFn1(50) {
            case -1:
                for fi122 := 0; fi122 < 17; fi122++ {
                    fmt.Println(fi122, v120, v4)
                    switch ("\n\t" + "hello") {
                    case "hello":
                        if v120 {
                            for fi123 := 0; fi123 < 14; fi123++ {
                                var v124 int64 = 1
                                var v125 uint64 = 93
                                var v126 uint32 = 2147483648
                                v118 += (v3 % (((-(v4))) | 1))
                                _ = fi123
                                _ = v124
                                _ = v125
                                _ = v126
                            }
                            v127 := []int{32768}
                            var v128 int32 = 2147483647
                            _ = v127
                            _ = v128
                        } else {
                            if ("foo bar" != "hello") {
                                if (float32(((float64((0.0 - (1.5 * (-2.25 / (1e10 + 1e10))))) / 0.1) - 1.5)) == ((float32(3.14159) * float32(float64((1e10 - -2.25)))) + 1e10)) {
                                    for fi129 := 0; fi129 < 10; fi129++ {
                                        if v120 {
                                            v130 := GenS16{F0: 913, F1: int64(v3)}
                                            v131 := GenS12{F0: true, F1: 42, F2: 570}
                                            fmt.Println(v1)
                                            _ = v130
                                            _ = v131
                                        } else {
                                            v132 := GenS17{F0: -128}
                                            _ = v132
                                        }
                                        _ = fi129
                                    }
                                    if v120 {
                                        if v120 {
                                            var v133 uint8 = 128
                                            if v120 {
                                                var v134 uint = 255
                                                var v135 int64 = 9223372036854775807
                                                if v120 {
                                                    fmt.Println(v3)
                                                    switch "foo bar" {
                                                    case "":
                                                        var v136 uint8 = (v133 - v133)
                                                        _ = v136
                                                    }
                                                    for fi137 := 0; fi137 < 20; fi137++ {
                                                        v133 += v133
                                                        v138 := GenS4{F0: "", F1: 0}
                                                        fmt.Println(fi122, v118, v120)
                                                        _ = fi137
                                                        _ = v138
                                                    }
                                                }
                                                _ = v134
                                                _ = v135
                                            }
                                            _ = v133
                                        }
                                        v139 := []int{fi122, 42, genFn1(-49)}
                                        _ = v139
                                    }
                                    fmt.Println(v118, v3)
                                } else {
                                    var v140 uint64 = 42
                                    v140 -= 2
                                    _ = v140
                                }
                            } else {
                                var v141 uint32 = 2147483647
                                if (127 <= -1) {
                                    v142 := GenS18{F0: uint16(v141)}
                                    for fi143 := 0; fi143 < 5; fi143++ {
                                        v1 = (fi143 & (fi143 | fi143))
                                        _ = fi143
                                    }
                                    fmt.Println(v118, v1, fi122)
                                    _ = v142
                                } else {
                                    switch "hello" {
                                    case "":
                                        v141 = (v141 / ((genFn0(65535, v120)) | 1))
                                    case "\n\t":
                                        var v144 float64 = (3.14159 + 0.1)
                                        _ = v144
                                    case "foo bar":
                                        if (int8(v141) != 0) {
                                            v3 = 0
                                        } else {
                                            for fi145 := 0; fi145 < 5; fi145++ {
                                                var v146 float64 = 0.0
                                                switch "foo bar" {
                                                case "hello":
                                                    var v147 bool = v120
                                                    _ = v147
                                                case "\n\t":
                                                    var v148 float64 = (3.14159 + (((v146 * (float64(1e10) + 3.14159)) - v146) - v146))
                                                    _ = v148
                                                case "a":
                                                    if (false && false) {
                                                        v1 = v1
                                                        v149 := GenS1{F0: 255}
                                                        v150 := GenS19{F0: 564, F1: (^(v4))}
                                                        fmt.Println(v120, v118, v120)
                                                        _ = v149
                                                        _ = v150
                                                    } else {
                                                        v141 += 2147483648
                                                    }
                                                default:
                                                    v141 = (v141 >> ((2) & 63))
                                                }
                                                _ = fi145
                                                _ = v146
                                            }
                                            if v120 {
                                                for fi151 := 0; fi151 < 7; fi151++ {
                                                    v152 := GenS8{F0: 65535}
                                                    v120 = ((false && true) && v120)
                                                    _ = fi151
                                                    _ = v152
                                                }
                                                var v153 float32 = 1e10
                                                _ = v153
                                            } else {
                                                for fi154 := 0; fi154 < 7; fi154++ {
                                                    fmt.Println(v120)
                                                    var v155 int64 = -1
                                                    _ = fi154
                                                    _ = v155
                                                }
                                            }
                                            v156 := v2["hello"]
                                            _ = v156
                                        }
                                    }
                                    for fi157 := 0; fi157 < 1; fi157++ {
                                        v118 += v118
                                        v158 := GenS13{F0: 0, F1: ("a" + ("" + "\n\t"))}
                                        v120 = true
                                        _ = fi157
                                        _ = v158
                                    }
                                    v118 = (v3 & (v118 % (((v3 - v118)) | 1)))
                                }
                                for fi159 := 0; fi159 < 11; fi159++ {
                                    var v160 int32 = (v3 + -652)
                                    var v161 uint32 = (^(v141))
                                    var v162 string = "hello"
                                    _ = fi159
                                    _ = v160
                                    _ = v161
                                    _ = v162
                                }
                                _ = v141
                            }
                            for fi163 := 0; fi163 < 12; fi163++ {
                                _ = len(v2)
                                switch 1 {
                                case 42:
                                    var v164 float32 = 3.14159
                                    _ = v164
                                case 1:
                                    if (!true) {
                                        v3 = (v4 - (v3 & 937))
                                        if false {
                                            var v165 string = ""
                                            _ = v165
                                        }
                                        var v166 uint32 = 2
                                        _ = v166
                                    } else {
                                        v167 := GenS3{F0: 0}
                                        v168 := []int{(fi163 >> ((255) & 63))}
                                        _ = v167
                                        _ = v168
                                    }
                                default:
                                    v169 := GenS20{F0: 1, F1: 127, F2: (^(fi122))}
                                    _ = v169
                                }
                                fmt.Println(v1, v3, fi122)
                                _ = fi163
                            }
                            v170 := GenS21{F0: 111}
                            _ = v170
                        }
                    case "a":
                        var v171 uint8 = 2
                        _ = v171
                    }
                    if false {
                        var v172 float32 = 1.5
                        switch 128 {
                        case 1:
                            v172 = (-2.25 * 1.5)
                        case -549:
                            switch "\n\t" {
                            case "hello":
                                var v173 int64 = -737
                                _ = v173
                            case "\n\t":
                                switch (^(fi122)) {
                                case 558:
                                    if v120 {
                                        if v120 {
                                            fmt.Println(v3, v118, v172)
                                            v4 = (^(v118))
                                        } else {
                                            v174 := []int{}
                                            _ = v174
                                        }
                                    } else {
                                        if (int8(v4) <= 100) {
                                            v175 := GenS22{F0: 42}
                                            v176 := []int{0, (fi122 + fi122)}
                                            _ = v175
                                            _ = v176
                                        }
                                        for fi177 := 0; fi177 < 2; fi177++ {
                                            var v178 uint16 = 42
                                            _ = fi177
                                            _ = v178
                                        }
                                    }
                                case 127:
                                    switch ((("foo bar" + "\n\t") + "a") + "") {
                                    case "foo bar":
                                        fmt.Println(v4, v3, v172)
                                    case "a":
                                        v179 := GenS23{F0: "\n\t", F1: 89}
                                        _ = v179
                                    }
                                case -9223372036854775808:
                                    if (v120 || v120) {
                                        v120 = ("a" == "a")
                                        switch "hello" {
                                        case "hello":
                                            if false {
                                                if (1 >= -9223372036854775808) {
                                                    switch "hello" {
                                                    case "":
                                                        var v180 float64 = 3.14159
                                                        _ = v180
                                                    }
                                                }
                                                if (0 > 18446744073709551615) {
                                                    v181 := []int{1, (-(fi122)), fi122, (fi122 / (((fi122 + (v1 << ((0) & 63)))) | 1))}
                                                    v172 -= (v172 * (1e10 - v172))
                                                    if (v3 < (-(v4))) {
                                                        var v182 uint8 = 203
                                                        v183 := v2["\n\t"]
                                                        _ = v182
                                                        _ = v183
                                                    } else {
                                                        v3 += v3
                                                        var v184 string = ""
                                                        _ = v184
                                                    }
                                                    _ = v181
                                                }
                                                if (42 <= 115) {
                                                    for fi185 := 0; fi185 < 7; fi185++ {
                                                        if (619 > 42) {
                                                            if (v120 && true) {
                                                                var v186 int16 = 42
                                                                fmt.Println(v1)
                                                                _ = v186
                                                            }
                                                        }
                                                        v4 += -718
                                                        _ = fi185
                                                    }
                                                    for fi187 := 0; fi187 < 7; fi187++ {
                                                        if v120 {
                                                            fmt.Println(v1)
                                                            var v188 uint16 = 1
                                                            v189 := GenS2{F0: 32767, F1: 219}
                                                            _ = v188
                                                            _ = v189
                                                        } else {
                                                            var v190 int8 = 127
                                                            v118 = (v118 / (((v118 & 0)) | 1))
                                                            var v191 float32 = (0.0 + v172)
                                                            _ = v190
                                                            _ = v191
                                                        }
                                                        fmt.Println(v1, v4, v120)
                                                        var v192 uint = 966
                                                        _ = fi187
                                                        _ = v192
                                                    }
                                                } else {
                                                    v120 = v120
                                                }
                                            } else {
                                                var v193 int = (fi122 + -702)
                                                for fi194 := 0; fi194 < 18; fi194++ {
                                                    v195 := []int{(v193 << ((255) & 63))}
                                                    _ = fi194
                                                    _ = v195
                                                }
                                                if false {
                                                    v196 := []int{(-(v1))}
                                                    fmt.Println(v3, v4, v118)
                                                    _ = len(v196)
                                                    _ = v196
                                                }
                                                _ = v193
                                            }
                                        case "":
                                            v197 := GenS2{F0: 0, F1: 4294967295}
                                            _ = v197
                                        default:
                                            v198 := GenS24{F0: (128 != 128), F1: 65535, F2: ("\n\t" + ("" + "foo bar"))}
                                            _ = v198
                                        }
                                    } else {
                                        v199 := v2["hello"]
                                        fmt.Println(v1)
                                        fmt.Println(v120, v172, v172)
                                        _ = v199
                                    }
                                default:
                                    fmt.Println(fi122, v3)
                                }
                            case "":
                                v200 := GenS10{F0: -1, F1: 0}
                                _ = v200
                            }
                        }
                        var v201 float64 = (1e10 + -2.25)
                        _ = v172
                        _ = v201
                    } else {
                        for fi202 := 0; fi202 < 16; fi202++ {
                            v2[""] = v120
                            v1 = 32767
                            var v203 string = "a"
                            var v204 string = ((v203 + v203) + "\n\t")
                            switch 32767 {
                            case 480:
                                v204 = "a"
                            case -9223372036854775808:
                                var v205 bool = (0.0 < ((3.14159 - 0.1) - ((((float64(3.14159) - -2.25) / 3.14159) * ((0.0 * (3.14159 - 1.5)) * (0.0 * 0.0))) * (3.14159 - 1e10))))
                                _ = v205
                            }
                            _ = fi202
                            _ = v203
                            _ = v204
                        }
                        var v206 uint64 = 486
                        v3 = (v3 ^ (v4 | v3))
                        _ = v206
                    }
                    _ = fi122
                }
            case 769:
                var v207 uint = 65535
                _ = v207
            }
        } else {
            for fi208 := 0; fi208 < 10; fi208++ {
                var v209 uint16 = 808
                for fi210 := 0; fi210 < 6; fi210++ {
                    if false {
                        var v211 int16 = -32768
                        if (v120 || v120) {
                            var v212 uint16 = (v209 + (v209 - v209))
                            _ = v212
                        }
                        _ = v211
                    } else {
                        switch (-(fi210)) {
                        case 968:
                            if v120 {
                                v4 -= (v4 * (v118 * (^(v3))))
                            } else {
                                if (385 >= 138) {
                                    var v213 int16 = -1
                                    _ = v213
                                }
                                v1 = -1
                                v214 := GenS25{F0: 2, F1: 0, F2: "foo bar"}
                                _ = v214
                            }
                        case -9223372036854775808:
                            v215 := GenS18{F0: (v209 / ((42) | 1))}
                            _ = v215
                        case 32767:
                            v216 := GenS9{F0: 9223372036854775807, F1: 849}
                            _ = v216
                        }
                        var v217 int32 = v4
                        v218 := GenS8{F0: 2}
                        _ = v217
                        _ = v218
                    }
                    v1 -= fi210
                    _ = fi210
                }
                _ = fi208
                _ = v209
            }
        }
        _ = v121
    }
    var v219 float64 = -2.25
    if v120 {
        for fi220 := 0; fi220 < 1; fi220++ {
            switch fi220 {
            case 301:
                var v221 int8 = 127
                _ = v221
            }
            _ = fi220
        }
        var v222 float64 = 1e10
        var v223 string = "\n\t"
        _ = v222
        _ = v223
    }
    v3 = -1
    v4 -= (v3 % (((^(v4))) | 1))
    fmt.Println(v120, v219, v219)
    v2["a"] = (v120 || (!v120))
    var v224 uint64 = 0
    var v225 int64 = 895
    v226 := GenS25{F0: 127, F1: 181, F2: "foo bar"}
    switch "\n\t" {
    case "":
        v227 := []int{v1, 243}
        _ = v227
    case "a":
        if false {
            var v228 int8 = 100
            _ = v228
        }
    case "\n\t":
        var v229 int = v1
        _ = v229
    default:
        v230 := GenS26{F0: (^(v118)), F1: 986}
        _ = v230
    }
    fmt.Println("v1:", v1)
    fmt.Println("v118:", v118)
    fmt.Println("v120:", v120)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v219:", v219)
    fmt.Println("v224:", v224)
    fmt.Println("v225:", v225)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v118
    _ = v119
    _ = v120
    _ = v2
    _ = v219
    _ = v224
    _ = v225
    _ = v226
    _ = v3
    _ = v4
}
