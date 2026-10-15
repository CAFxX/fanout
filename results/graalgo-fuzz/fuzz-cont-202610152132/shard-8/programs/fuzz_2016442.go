package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int; F1 string; F2 uint8 }
type GenS7 struct { F0 uint8; F1 int; F2 string }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint8; F1 uint8 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 int16; F1 string }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 int64; F1 uint32 }
func genFn0() uint8 {
    var v7 int16 = -32768
    for fi8 := 0; fi8 < 10; fi8++ {
        v9 := GenS1{F0: (fi8 ^ fi8), F1: -2}
        _ = fi8
        _ = v9
    }
    if (1 > 2) {
        switch 423 {
        case 9223372036854775807:
            for fi10 := 0; fi10 < 5; fi10++ {
                v11 := GenS2{F0: ((-1 != 2147483647) && (!true))}
                v12 := GenS2{F0: (false && (false && true))}
                fmt.Println(fi10)
                _ = fi10
                _ = v11
                _ = v12
            }
        case 127:
            for fi13 := 0; fi13 < 5; fi13++ {
                for fi14 := 0; fi14 < 18; fi14++ {
                    if ((!true) || true) {
                        v15 := GenS1{F0: (fi14 * fi13), F1: 42}
                        var v16 string = "hello"
                        v16 = "a"
                        if (42 < 2) {
                            v7 += 42
                            if (true && true) {
                                var v17 uint64 = 4294967295
                                v7 = v7
                                _ = v17
                            } else {
                                var v18 uint32 = 1
                                var v19 int64 = -753
                                if true {
                                    switch v16 {
                                    case "hello":
                                        v20 := GenS3{F0: 2147483648, F1: v18}
                                        _ = v20
                                    default:
                                        v21 := GenS1{F0: fi13, F1: 42}
                                        _ = v21
                                    }
                                    if false {
                                        var v22 float32 = 1.5
                                        var v23 uint8 = 2
                                        _ = v22
                                        _ = v23
                                    }
                                } else {
                                    var v24 int32 = -1
                                    _ = v24
                                }
                                _ = v18
                                _ = v19
                            }
                            v16 = v16
                        } else {
                            var v25 int = fi14
                            v26 := GenS2{F0: (true || (42 >= 1))}
                            _ = v25
                            _ = v26
                        }
                        var v27 uint64 = 265
                        switch fi13 {
                        case 2147483647:
                            v27 += v27
                        case -441:
                            switch ("foo bar" + "foo bar") {
                            case "foo bar":
                                var v28 float64 = 1.5
                                _ = v28
                            case "\n\t":
                                var v29 bool = ("foo bar" != (v16 + ("\n\t" + v16)))
                                _ = v29
                            default:
                                if false {
                                    v7 = (^(v7))
                                    var v30 int32 = 109
                                    if true {
                                        v31 := []int{}
                                        _ = v31
                                    } else {
                                        v32 := GenS2{F0: (float64(-2.25) != (0.1 + 3.14159))}
                                        switch (fi14 ^ (fi13 % (((fi14 & fi14)) | 1))) {
                                        case -9223372036854775808:
                                            v27 = 42
                                        case 270:
                                            var v33 string = (v16 + ("hello" + "\n\t"))
                                            _ = v33
                                        case 1:
                                            v34 := GenS4{F0: 1, F1: 162}
                                            _ = v34
                                        default:
                                            var v35 uint = uint(v27)
                                            _ = v35
                                        }
                                        v36 := []int{2147483647, fi13, 32768}
                                        v37 := GenS2{F0: false}
                                        _ = v32
                                        _ = v36
                                        _ = v37
                                    }
                                    _ = v30
                                } else {
                                    var v38 uint64 = (v27 | (v27 ^ v27))
                                    var v39 uint16 = 595
                                    switch v16 {
                                    case "a":
                                        if ((0.0 == float32(1e10)) && true) {
                                            v40 := []int{}
                                            v41 := GenS3{F0: 4294967295, F1: 2}
                                            _ = v40
                                            _ = v41
                                        } else {
                                            v42 := GenS5{F0: false}
                                            _ = v42
                                        }
                                    }
                                    _ = v38
                                    _ = v39
                                }
                            }
                        case 1:
                            var v43 uint64 = v27
                            _ = v43
                        }
                        v27 += v27
                        if (v27 >= (v27 >> ((662) & 63))) {
                            v44 := []int{(fi14 << ((197) & 63)), (-(fi14)), (fi13 ^ fi13), fi13}
                            switch (fi14 / (((-(fi13))) | 1)) {
                            case 9223372036854775807:
                                switch (fi14 & fi14) {
                                case 11:
                                    _ = len(v44)
                                case 9223372036854775807:
                                    if ((-2.25 - float32(float64((3.14159 / 1e10)))) < 0.0) {
                                        var v45 uint8 = 128
                                        v46 := GenS2{F0: true}
                                        fmt.Println(fi14, v45)
                                        _ = v45
                                        _ = v46
                                    } else {
                                        if len(v44) > 0 {
                                            v44[0] = (fi13 + 0)
                                        }
                                        var v47 string = v16
                                        _ = v47
                                    }
                                }
                            case 850:
                                v48 := GenS1{F0: fi13, F1: -100}
                                _ = v48
                            }
                            v49 := GenS3{F0: 315, F1: 0}
                            if ((v16 + v16) == "\n\t") {
                                switch v16 {
                                case "a":
                                    if (((3.14159 - (((((((3.14159 / (((0.1 / -2.25) + 1e10) - (float64(0.1) / (0.1 - float64((3.14159 - 1e10)))))) - (3.14159 / 3.14159)) + 3.14159) / (1.5 + -2.25)) - (1e10 / 0.1)) / 1e10) / 1.0)) - 1e10) < (0.1 - (0.1 * 1.5))) {
                                        if len(v44) > 0 {
                                            v44[0] = (^(fi13))
                                        }
                                    }
                                    if (0 <= 42) {
                                        fmt.Println(v27)
                                    }
                                case "foo bar":
                                    if (!false) {
                                        if len(v44) >= 2 {
                                            v44 = v44[1:len(v44)]
                                        }
                                        _ = len(v44)
                                        var v50 int = fi14
                                        _ = v50
                                    }
                                case "hello":
                                    v51 := GenS3{F0: 42, F1: 2147483648}
                                    if len(v44) > 0 {
                                        v44[0] = fi14
                                    }
                                    _ = v51
                                default:
                                    switch (fi14 - (fi14 + (^(fi13)))) {
                                    case 42:
                                        var v52 uint32 = 42
                                        _ = v52
                                    case -2147483648:
                                        if true {
                                            if false {
                                                var v53 int64 = 411
                                                _ = v53
                                            } else {
                                                _ = len(v44)
                                                var v54 float64 = (3.14159 / 0.1)
                                                _ = v54
                                            }
                                        } else {
                                            _ = len(v44)
                                            if false {
                                                _ = len(v44)
                                            } else {
                                                fmt.Println(v7, v27)
                                            }
                                        }
                                    case -315:
                                        if (false || false) {
                                            if len(v44) >= 2 {
                                                v44 = v44[1:len(v44)]
                                            }
                                        } else {
                                            var v55 int64 = -9223372036854775808
                                            _ = v55
                                        }
                                    default:
                                        v16 = v16
                                    }
                                }
                                var v56 int = (fi13 * fi13)
                                _ = v56
                            } else {
                                if len(v44) >= 2 {
                                    v44 = v44[1:len(v44)]
                                }
                                v27 -= 809
                                v57 := GenS3{F0: 2, F1: 42}
                                _ = v57
                            }
                            _ = v44
                            _ = v49
                        } else {
                            var v58 uint = 0
                            v59 := GenS6{F0: (-(fi14)), F1: v16, F2: 42}
                            _ = v58
                            _ = v59
                        }
                        var v60 int = 0
                        switch v60 {
                        case 2147483647:
                            switch (^(fi13)) {
                            case 42:
                                v61 := []int{}
                                _ = v61
                            default:
                                v62 := []int{}
                                _ = v62
                            }
                        case -9223372036854775808:
                            var v63 int8 = 43
                            _ = v63
                        default:
                            v60 += fi13
                        }
                        _ = v15
                        _ = v16
                        _ = v27
                        _ = v60
                    } else {
                        v64 := []int{fi13, fi13, (fi14 - fi14), fi14}
                        v7 = 1000
                        v64 = append(v64, 219)
                        _ = v64
                    }
                    v65 := GenS5{F0: (65535 >= 531)}
                    v7 -= (^(v7))
                    _ = fi14
                    _ = v65
                }
                fmt.Println(v7, v7)
                switch (fi13 ^ (^(fi13))) {
                case 4294967295:
                    v66 := GenS5{F0: ((false && (296 < 386)) && false)}
                    _ = v66
                case -706:
                    v7 = v7
                case 127:
                    v67 := []int{-2147483648, fi13, (fi13 + (fi13 * fi13))}
                    _ = v67
                }
                _ = fi13
            }
        default:
            v68 := GenS3{F0: 0, F1: 42}
            _ = v68
        }
    }
    v69 := GenS1{F0: -142, F1: 100}
    _ = v69
    _ = v7
    return 127
}

func main() {
    var v1 float64 = 1.5
    var v2 int32 = -2147483648
    var v3 float32 = (1e10 - 1.5)
    var v4 int = 128
    v5 := make(map[int]int)
    var v6 bool = (0 != 701)
    var v70 int = 9223372036854775807
    v1 += v1
    var v71 float32 = float32((0.0 * 3.14159))
    v2 = (-(v2))
    var v72 int32 = v2
    for fi73 := 0; fi73 < 1; fi73++ {
        v71 += float32((float64(1.5) + (float64((3.14159 - v3)) * float64(v3))))
        for fi74 := 0; fi74 < 16; fi74++ {
            v4 = -9223372036854775808
            v75 := GenS7{F0: 255, F1: 979, F2: ((("\n\t" + (("foo bar" + ("" + ("" + ""))) + "hello")) + "") + ((("hello" + "hello") + ("" + "")) + ("" + "\n\t")))}
            fmt.Println(v70)
            _ = fi74
            _ = v75
        }
        fmt.Println(v2, v6, v3)
        _ = fi73
    }
    v2 -= v72
    var v76 float64 = v1
    fmt.Println(v1, v3, v3)
    if v6 {
        for fi77 := 0; fi77 < 14; fi77++ {
            if v6 {
                switch v70 {
                case 9223372036854775807:
                    for fi78 := 0; fi78 < 17; fi78++ {
                        fmt.Println(fi78)
                        var v79 uint64 = 18446744073709551615
                        v80 := []int{(fi77 % ((128) | 1)), fi78, -15, fi78}
                        _ = fi78
                        _ = v79
                        _ = v80
                    }
                case -508:
                    v81 := GenS6{F0: v70, F1: ("\n\t" + ("\n\t" + ("a" + "hello"))), F2: genFn0()}
                    _ = v81
                case 2147483647:
                    var v82 bool = true
                    _ = v82
                default:
                    v71 += 3.14159
                }
            }
            if v6 {
                if ((v6 || (true || false)) && false) {
                    v83 := GenS7{F0: 0, F1: (v4 / ((fi77) | 1)), F2: ("hello" + "\n\t")}
                    _ = v83
                } else {
                    var v84 int16 = -142
                    v3 -= v71
                    var v85 int16 = (v84 / (((v84 * (^(v84)))) | 1))
                    _ = v84
                    _ = v85
                }
            }
            _ = fi77
        }
    } else {
        if (!v6) {
            v86 := []int{(v4 ^ 682)}
            var v87 bool = true
            for fi88 := 0; fi88 < 17; fi88++ {
                if (2 != 92) {
                    fmt.Println(v1, v6, v71)
                    v6 = (4294967295 >= (v4 + (v4 & 32767)))
                    switch "foo bar" {
                    case "foo bar":
                        v5[fi88] = fi88
                    case "":
                        v89 := GenS3{F0: 563, F1: 2}
                        _ = v89
                    case "hello":
                        fmt.Println(fi88)
                    }
                }
                if v87 {
                    var v90 int32 = (v2 >> ((18446744073709551615) & 63))
                    switch "\n\t" {
                    case "a":
                        if v87 {
                            if true {
                                v91 := GenS7{F0: 255, F1: (fi88 ^ (v4 - fi88)), F2: "a"}
                                if (!(253 == -602)) {
                                    fmt.Println(v3, v3)
                                } else {
                                    v92 := GenS3{F0: 2147483648, F1: 65535}
                                    delete(v5, fi88)
                                    _ = v92
                                }
                                if (!false) {
                                    var v93 int8 = -128
                                    _ = v93
                                }
                                _ = v91
                            }
                            v94 := GenS8{F0: 166}
                            _ = v94
                        } else {
                            fmt.Println(v3)
                        }
                    case "\n\t":
                        for fi95 := 0; fi95 < 2; fi95++ {
                            if len(v86) >= 2 {
                                v86 = v86[1:len(v86)]
                            }
                            _ = fi95
                        }
                    case "foo bar":
                        v96 := GenS9{F0: -757}
                        _ = v96
                    default:
                        switch (v4 - v70) {
                        case 127:
                            var v97 int64 = 42
                            _ = v97
                        case 225:
                            if (1 == 829) {
                                var v98 float64 = ((3.14159 * v1) + (v1 + 0.0))
                                v70 += (v4 % ((32767) | 1))
                                v99 := GenS9{F0: 1}
                                _ = v98
                                _ = v99
                            } else {
                                v86 = append(v86, v4)
                                if len(v86) > 0 {
                                    v86[0] = v4
                                }
                                for fi100 := 0; fi100 < 13; fi100++ {
                                    v101 := v5[(fi100 + 786)]
                                    _ = fi100
                                    _ = v101
                                }
                            }
                        case -1:
                            if ((v72 << ((647) & 63)) > (v2 / ((v72) | 1))) {
                                _ = len(v86)
                                switch (^(v4)) {
                                case -864:
                                    v102 := GenS10{F0: 616}
                                    _ = v102
                                case -9223372036854775808:
                                    v3 -= ((0.1 + v71) / 0.1)
                                case 0:
                                    var v103 int = v70
                                    _ = v103
                                }
                            }
                        default:
                            if (0 < -945) {
                                v3 += -2.25
                                switch ("hello" + ((("a" + ("foo bar" + "foo bar")) + "foo bar") + "foo bar")) {
                                case "hello":
                                    fmt.Println(v6)
                                case "a":
                                    var v104 uint64 = 18446744073709551615
                                    _ = v104
                                case "\n\t":
                                    for fi105 := 0; fi105 < 1; fi105++ {
                                        v86 = append(v86, v4)
                                        _ = fi105
                                    }
                                default:
                                    v106 := GenS8{F0: 65535}
                                    _ = v106
                                }
                                fmt.Println(v1, v72, v87)
                            }
                        }
                    }
                    for fi107 := 0; fi107 < 12; fi107++ {
                        v71 = ((1e10 / 1.5) / 1e10)
                        var v108 uint32 = 808
                        _ = fi107
                        _ = v108
                    }
                    _ = v90
                } else {
                    if v87 {
                        var v109 int64 = 119
                        _ = v109
                    }
                    if v6 {
                        switch (v70 >> ((978) & 63)) {
                        case 4294967295:
                            switch (("\n\t" + "a") + "hello") {
                            case "foo bar":
                                v2 -= v2
                            case "\n\t":
                                var v110 bool = (-1 >= 59)
                                _ = v110
                            }
                        case -66:
                            if v6 {
                                for fi111 := 0; fi111 < 14; fi111++ {
                                    switch (("hello" + "\n\t") + "a") {
                                    case "\n\t":
                                        delete(v5, (^(fi111)))
                                    case "":
                                        if v6 {
                                            v112 := GenS11{F0: 180}
                                            v2 += 1
                                            _ = v112
                                        }
                                    default:
                                        switch -9223372036854775808 {
                                        case -2147483648:
                                            var v113 int16 = 1
                                            v114 := GenS1{F0: (fi88 % ((2147483647) | 1)), F1: 105}
                                            _ = len(v86)
                                            var v115 uint8 = 42
                                            v70 = (fi111 % ((v70) | 1))
                                            _ = v113
                                            _ = v114
                                            _ = v115
                                        case 2147483647:
                                            v72 = -87
                                            fmt.Println(v4, v70, v72)
                                            var v116 uint32 = 4294967295
                                            _ = v116
                                        case 128:
                                            v1 = (v76 * v76)
                                        default:
                                            v117 := GenS11{F0: 42}
                                            _ = v117
                                        }
                                    }
                                    if len(v86) >= 2 {
                                        v86 = v86[1:len(v86)]
                                    }
                                    v5[(fi111 >> ((18446744073709551615) & 63))] = 32767
                                    _ = fi111
                                }
                                if ("a" != "foo bar") {
                                    var v118 float32 = (float32(float64((v3 / float32(v1)))) - (v71 * ((3.14159 - v71) * float32((v1 * float64(v71))))))
                                    v6 = true
                                    var v119 string = (("hello" + ("" + "foo bar")) + "hello")
                                    _ = v118
                                    _ = v119
                                }
                                switch 2147483647 {
                                case 32767:
                                    fmt.Println(v71, v70, v76)
                                case 2147483647:
                                    for fi120 := 0; fi120 < 9; fi120++ {
                                        var v121 uint8 = 1
                                        v122 := GenS11{F0: 20}
                                        _ = fi120
                                        _ = v121
                                        _ = v122
                                    }
                                case 4294967295:
                                    for fi123 := 0; fi123 < 20; fi123++ {
                                        _ = len(v5)
                                        v3 = 0.0
                                        _ = fi123
                                    }
                                }
                            }
                        default:
                            if (false || v6) {
                                switch (fi88 ^ v4) {
                                case -1:
                                    switch "hello" {
                                    case "foo bar":
                                        v2 -= (v72 & (^(v72)))
                                    case "\n\t":
                                        fmt.Println(v87, v6)
                                    default:
                                        var v124 string = ("a" + "\n\t")
                                        _ = v124
                                    }
                                case 32767:
                                    if v6 {
                                        if (!(0.0 > (v76 - v76))) {
                                            if true {
                                                for fi125 := 0; fi125 < 11; fi125++ {
                                                    v2 = 0
                                                    var v126 uint8 = genFn0()
                                                    if (!((v4 + (v70 >> ((65535) & 63))) != fi125)) {
                                                        var v127 uint32 = 2
                                                        if ((42 > 0) && v87) {
                                                            fmt.Println(v3)
                                                        } else {
                                                            if (42 != v72) {
                                                                switch "\n\t" {
                                                                case "":
                                                                    var v128 uint8 = genFn0()
                                                                    _ = v128
                                                                }
                                                                v129 := GenS9{F0: -1}
                                                                _ = v129
                                                            } else {
                                                                fmt.Println(v4, v70)
                                                                var v130 uint32 = (v127 * 0)
                                                                if ("hello" < "") {
                                                                    var v131 int64 = int64(v130)
                                                                    v2 = (v2 * int32(v130))
                                                                    if v6 {
                                                                        var v132 float64 = v76
                                                                        _ = v132
                                                                    } else {
                                                                        v2 -= 19
                                                                        var v133 bool = v87
                                                                        _ = v133
                                                                    }
                                                                    _ = v131
                                                                } else {
                                                                    var v134 uint64 = 0
                                                                    v76 -= (3.14159 * v1)
                                                                    var v135 uint32 = (v130 & 2147483648)
                                                                    _ = v134
                                                                    _ = v135
                                                                }
                                                                _ = v130
                                                            }
                                                        }
                                                        var v136 int = (fi125 >> ((18446744073709551615) & 63))
                                                        _ = v127
                                                        _ = v136
                                                    }
                                                    _ = fi125
                                                    _ = v126
                                                }
                                                var v137 string = "\n\t"
                                                _ = v137
                                            } else {
                                                v138 := GenS12{F0: 255, F1: 42}
                                                for fi139 := 0; fi139 < 7; fi139++ {
                                                    delete(v5, (fi88 - fi88))
                                                    if len(v86) >= 2 {
                                                        v86 = v86[1:len(v86)]
                                                    }
                                                    _ = fi139
                                                }
                                                var v140 int = (v70 & (v70 / (((-(v4))) | 1)))
                                                _ = v138
                                                _ = v140
                                            }
                                            switch -9223372036854775808 {
                                            case 42:
                                                if (1 < -930) {
                                                    if false {
                                                        v141 := GenS12{F0: 46, F1: 255}
                                                        var v142 uint8 = 224
                                                        _ = v141
                                                        _ = v142
                                                    }
                                                    for fi143 := 0; fi143 < 17; fi143++ {
                                                        v144 := GenS11{F0: 917}
                                                        _ = fi143
                                                        _ = v144
                                                    }
                                                    if (v6 && (v3 < v71)) {
                                                        var v145 int8 = 1
                                                        _ = v145
                                                    }
                                                } else {
                                                    v86 = append(v86, 128)
                                                    v146 := GenS13{F0: (!false)}
                                                    if v6 {
                                                        v147 := GenS9{F0: 1}
                                                        if v6 {
                                                            v148 := GenS2{F0: (2147483648 > 65535)}
                                                            _ = v148
                                                        }
                                                        _ = v147
                                                    }
                                                    _ = v146
                                                }
                                            default:
                                                for fi149 := 0; fi149 < 3; fi149++ {
                                                    var v150 string = ""
                                                    var v151 int64 = 731
                                                    var v152 int16 = 0
                                                    _ = fi149
                                                    _ = v150
                                                    _ = v151
                                                    _ = v152
                                                }
                                            }
                                            fmt.Println(v3)
                                        } else {
                                            fmt.Println(v4)
                                            var v153 uint8 = 128
                                            v2 = v2
                                            _ = v153
                                        }
                                        v154 := GenS14{F0: -1, F1: ("" + ("foo bar" + "a"))}
                                        v86 = append(v86, v4)
                                        _ = v154
                                    }
                                default:
                                    fmt.Println(fi88, v6)
                                }
                            } else {
                                var v155 bool = (!true)
                                switch v70 {
                                case 1:
                                    var v156 uint64 = 0
                                    _ = v156
                                case -9223372036854775808:
                                    v6 = v87
                                case 32768:
                                    v157 := GenS7{F0: genFn0(), F1: v70, F2: "hello"}
                                    _ = v157
                                }
                                var v158 uint = 2
                                _ = v155
                                _ = v158
                            }
                        }
                    }
                    var v159 int = v70
                    _ = v159
                }
                _ = fi88
            }
            _ = v86
            _ = v87
        }
        for fi160 := 0; fi160 < 7; fi160++ {
            if true {
                v161 := GenS15{F0: -32768}
                if v6 {
                    for fi162 := 0; fi162 < 6; fi162++ {
                        var v163 int16 = 0
                        v3 += (v3 + -2.25)
                        _ = fi162
                        _ = v163
                    }
                } else {
                    switch "\n\t" {
                    case "":
                        v164 := v5[v4]
                        _ = v164
                    case "foo bar":
                        v165 := GenS6{F0: (fi160 + (-(fi160))), F1: "foo bar", F2: 2}
                        _ = v165
                    }
                    fmt.Println(v72, v6, v70)
                    v6 = (v6 || (100 == -128))
                }
                _ = v161
            } else {
                v2 = -1
                var v166 string = "\n\t"
                var v167 int64 = 160
                _ = v166
                _ = v167
            }
            switch v70 {
            case 32767:
                v1 = float64(1.5)
            case 4294967295:
                var v168 string = "foo bar"
                _ = v168
            }
            fmt.Println(v1)
            _ = fi160
        }
        if v6 {
            var v169 string = ""
            _ = v169
        } else {
            for fi170 := 0; fi170 < 19; fi170++ {
                v4 = -2147483648
                var v171 uint64 = 18446744073709551615
                _ = fi170
                _ = v171
            }
        }
    }
    if v6 {
        if ((1 != -9223372036854775808) || (!v6)) {
            var v172 uint16 = 76
            _ = v172
        } else {
            var v173 uint8 = 213
            if (100 < 1) {
                if true {
                    v174 := GenS16{F0: 1, F1: 4294967295}
                    if (701 != 32767) {
                        var v175 float32 = ((1e10 - float32(((((1.5 * (0.1 - v1)) * v1) / (v1 * (v76 / v76))) + v1))) / v71)
                        switch (v4 + 128) {
                        case 42:
                            v176 := []int{v4}
                            _ = v176
                        case -2147483648:
                            var v177 uint16 = 32767
                            _ = v177
                        case 127:
                            for fi178 := 0; fi178 < 1; fi178++ {
                                v179 := []int{(v4 / ((fi178) | 1)), (v70 | fi178), 647}
                                v175 -= v3
                                v180 := GenS12{F0: v173, F1: (v173 + v173)}
                                _ = fi178
                                _ = v179
                                _ = v180
                            }
                        }
                        _ = v175
                    }
                    _ = v174
                }
                var v181 string = ("foo bar" + (("foo bar" + "a") + "foo bar"))
                for fi182 := 0; fi182 < 15; fi182++ {
                    for fi183 := 0; fi183 < 19; fi183++ {
                        v184 := GenS5{F0: (!(255 == 290))}
                        _ = fi183
                        _ = v184
                    }
                    _ = fi182
                }
                _ = v181
            } else {
                for fi185 := 0; fi185 < 16; fi185++ {
                    switch fi185 {
                    case -7:
                        var v186 int16 = 0
                        _ = v186
                    case -1:
                        switch 32768 {
                        case 9223372036854775807:
                            var v187 uint32 = 2
                            _ = v187
                        case 4294967295:
                            if v6 {
                                for fi188 := 0; fi188 < 11; fi188++ {
                                    var v189 uint64 = 969
                                    _ = fi188
                                    _ = v189
                                }
                            } else {
                                for fi190 := 0; fi190 < 15; fi190++ {
                                    _ = len(v5)
                                    var v191 int64 = -753
                                    v192 := GenS6{F0: v4, F1: ("a" + "hello"), F2: (^(v173))}
                                    _ = fi190
                                    _ = v191
                                    _ = v192
                                }
                                for fi193 := 0; fi193 < 1; fi193++ {
                                    var v194 uint32 = 42
                                    _ = fi193
                                    _ = v194
                                }
                            }
                        case -1:
                            if true {
                                for fi195 := 0; fi195 < 9; fi195++ {
                                    v71 = v3
                                    v6 = (v76 == float64(float32(v76)))
                                    var v196 uint16 = 32767
                                    _ = fi195
                                    _ = v196
                                }
                            } else {
                                fmt.Println(v70)
                            }
                        default:
                            switch (^(v4)) {
                            case 32768:
                                switch (("foo bar" + ("hello" + "")) + ("\n\t" + "\n\t")) {
                                case "foo bar":
                                    v72 = v72
                                case "\n\t":
                                    v6 = v6
                                case "a":
                                    v197 := GenS9{F0: 150}
                                    _ = v197
                                }
                            case 0:
                                fmt.Println(v1, v1)
                            }
                        }
                    }
                    v198 := GenS8{F0: 45}
                    _ = fi185
                    _ = v198
                }
                v6 = (948 <= 42)
                v199 := GenS2{F0: ((v4 - (v4 ^ (v4 ^ (v4 >> ((728) & 63))))) >= (v70 << ((351) & 63)))}
                _ = v199
            }
            v200 := []int{int(v173)}
            _ = v173
            _ = v200
        }
        if v6 {
            v71 += 0.1
            if false {
                for fi201 := 0; fi201 < 4; fi201++ {
                    v202 := GenS4{F0: -1, F1: 2}
                    for fi203 := 0; fi203 < 19; fi203++ {
                        var v204 bool = ((-2.25 * (v76 + float64((v71 + (v71 - ((v3 * -2.25) - 1.5)))))) != v1)
                        _ = fi203
                        _ = v204
                    }
                    for fi205 := 0; fi205 < 10; fi205++ {
                        if true {
                            fmt.Println(v71)
                        }
                        _ = fi205
                    }
                    _ = fi201
                    _ = v202
                }
                for fi206 := 0; fi206 < 1; fi206++ {
                    var v207 uint8 = 141
                    v207 -= v207
                    v3 = v3
                    _ = fi206
                    _ = v207
                }
            }
        } else {
            v72 = (v72 + v2)
            fmt.Println(v76, v71)
        }
    }
    v4 = v70
    var v208 float64 = v1
    v71 += float32(v1)
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v208:", v208)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v208
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v70
    _ = v71
    _ = v72
    _ = v76
}
