package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 uint8; F2 int32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint8; F1 int16 }
type GenS7 struct { F0 uint8; F1 uint32 }
type GenS8 struct { F0 int32; F1 uint16; F2 uint }
type GenS9 struct { F0 uint16; F1 int32 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int64; F1 int; F2 uint16 }
type GenS12 struct { F0 string; F1 bool; F2 uint32 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 int64; F1 int }
type GenS16 struct { F0 uint16; F1 uint8; F2 uint }
type GenS17 struct { F0 uint64 }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 uint }
func genFn0(p0 uint8) uint32 {
    fmt.Println(p0)
    p0 += 255
    v7 := GenS1{F0: 42, F1: 405, F2: 65535}
    _ = p0
    _ = v7
    return 564
}
func genFn1(p0 uint8, p1 string) uint64 {
    p0 += p0
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := []int{fi8, -2147483648, (-(fi8)), (fi8 | fi8)}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 string = "hello"
        _ = fi10
        _ = v11
    }
    switch 134 {
    case -9223372036854775808:
        var v12 int = 0
        _ = v12
    case -612:
        fmt.Println(p0)
    default:
        var v13 float64 = 0.0
        _ = v13
    }
    _ = p0
    _ = p1
    return 743
}

func main() {
    var v1 int = 128
    var v2 uint64 = 774
    var v3 int64 = 9223372036854775807
    var v4 int64 = (v3 - 29)
    var v5 float32 = 0.0
    var v6 uint8 = 42
    var v14 int32 = -2147483648
    if true {
        switch v1 {
        case 128:
            v15 := GenS1{F0: 167, F1: 0, F2: 834}
            _ = v15
        case 42:
            if (v4 < v4) {
                var v16 int32 = (v14 - (^(v14)))
                v6 = v6
                if false {
                    v5 += -2.25
                    if ((v16 <= v14) && (!false)) {
                        switch ("a" + "a") {
                        case "foo bar":
                            var v17 uint16 = 384
                            _ = v17
                        case "hello":
                            v5 += 1e10
                        case "a":
                            if false {
                                if (!(1 != 32767)) {
                                    switch ("\n\t" + "hello") {
                                    case "foo bar":
                                        v16 = (v16 * (-(v14)))
                                    case "":
                                        v18 := []int{}
                                        _ = v18
                                    case "a":
                                        if ((true || (0.1 < 1.5)) && true) {
                                            v4 -= v3
                                        }
                                    default:
                                        switch "\n\t" {
                                        case "\n\t":
                                            v19 := GenS1{F0: 0, F1: 65535, F2: uint(v2)}
                                            _ = v19
                                        case "a":
                                            v20 := []int{v1, (-(v1))}
                                            _ = v20
                                        case "":
                                            v21 := GenS1{F0: 252, F1: 65535, F2: 443}
                                            _ = v21
                                        default:
                                            var v22 bool = (v14 >= v16)
                                            _ = v22
                                        }
                                    }
                                    fmt.Println(v5, v4, v5)
                                    if true {
                                        switch "a" {
                                        case "foo bar":
                                            v2 = v2
                                        default:
                                            v6 = (^(v6))
                                        }
                                        for fi23 := 0; fi23 < 8; fi23++ {
                                            v24 := GenS2{F0: 1}
                                            _ = fi23
                                            _ = v24
                                        }
                                        for fi25 := 0; fi25 < 13; fi25++ {
                                            v16 = (v14 | v14)
                                            v26 := GenS3{F0: v14}
                                            _ = fi25
                                            _ = v26
                                        }
                                    } else {
                                        v27 := GenS2{F0: 42}
                                        if (((3.14159 / ((1e10 * (float64((v5 - 0.1)) - -2.25)) - 1e10)) > float64(float32(float64(float32(3.14159))))) && false) {
                                            v28 := []int{v1}
                                            _ = v28
                                        } else {
                                            var v29 int16 = 1
                                            fmt.Println(v29, v16)
                                            _ = v29
                                        }
                                        _ = v27
                                    }
                                } else {
                                    var v30 string = ""
                                    v31 := GenS4{F0: (!true), F1: v6, F2: 554}
                                    _ = v30
                                    _ = v31
                                }
                                var v32 uint64 = v2
                                switch "foo bar" {
                                case "a":
                                    var v33 int8 = 123
                                    _ = v33
                                }
                                _ = v32
                            } else {
                                v4 = v3
                                v34 := []int{32768, int(v14)}
                                if ((!(((v4 >> ((42) & 63)) < v3) || false)) || true) {
                                    if (!true) {
                                        v35 := GenS5{F0: v6}
                                        _ = v35
                                    } else {
                                        if (65535 <= 255) {
                                            v5 = 1e10
                                        } else {
                                            v2 += (v2 & (v2 >> ((122) & 63)))
                                            v36 := GenS2{F0: 106}
                                            var v37 float64 = -2.25
                                            _ = v36
                                            _ = v37
                                        }
                                        v38 := GenS6{F0: v6, F1: 32767}
                                        v34 = append(v34, (v1 & (v1 & int(v14))))
                                        _ = v38
                                    }
                                    var v39 uint64 = (v2 % ((genFn1(v6, ("\n\t" + ("hello" + "")))) | 1))
                                    _ = v39
                                }
                                _ = v34
                            }
                        }
                        v40 := GenS7{F0: (^(v6)), F1: 533}
                        v41 := GenS3{F0: v16}
                        _ = v40
                        _ = v41
                    }
                } else {
                    if true {
                        v1 = (v1 << ((948) & 63))
                        if (!((1.5 - -2.25) == 0.1)) {
                            var v42 int8 = -128
                            _ = v42
                        } else {
                            if true {
                                var v43 uint16 = 0
                                _ = v43
                            }
                            v44 := []int{(v1 ^ v1)}
                            var v45 uint64 = 4294967295
                            _ = v44
                            _ = v45
                        }
                        v3 = v3
                    } else {
                        var v46 int32 = (v16 & (v14 % ((v16) | 1)))
                        v47 := []int{(-(v1))}
                        for fi48 := 0; fi48 < 11; fi48++ {
                            switch int(v14) {
                            case -472:
                                var v49 uint8 = 253
                                _ = v49
                            }
                            _ = fi48
                        }
                        _ = v46
                        _ = v47
                    }
                    for fi50 := 0; fi50 < 6; fi50++ {
                        var v51 int32 = (v16 % (((-(v16))) | 1))
                        v52 := []int{fi50, (v1 | fi50), (^(fi50)), fi50}
                        for fi53 := 0; fi53 < 13; fi53++ {
                            var v54 int16 = 374
                            _ = fi53
                            _ = v54
                        }
                        _ = fi50
                        _ = v51
                        _ = v52
                    }
                }
                _ = v16
            }
        case 9223372036854775807:
            if true {
                v5 += v5
                v14 += v14
                v6 = 101
            } else {
                if (((false && true) || true) || false) {
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(v2, v2, v6)
                    case "a":
                        v14 = 965
                    }
                } else {
                    for fi55 := 0; fi55 < 1; fi55++ {
                        v56 := GenS8{F0: v14, F1: 255, F2: 297}
                        v57 := GenS7{F0: 137, F1: 4294967295}
                        v5 -= 3.14159
                        _ = fi55
                        _ = v56
                        _ = v57
                    }
                    v58 := []int{2147483647, v1, (v1 % ((v1) | 1)), -1}
                    _ = v58
                }
                var v59 uint32 = genFn0(127)
                _ = v59
            }
        default:
            fmt.Println(v1, v1)
        }
        var v60 uint64 = 16
        _ = v60
    }
    var v61 int16 = -1
    if (2 > 42) {
        var v62 uint32 = 2
        _ = v62
    }
    v63 := GenS7{F0: v6, F1: genFn0((v6 / ((v6) | 1)))}
    var v64 string = "a"
    var v65 uint32 = 0
    if false {
        v2 -= (v2 & 344)
    } else {
        for fi66 := 0; fi66 < 14; fi66++ {
            v1 -= v1
            _ = fi66
        }
        v4 -= v3
        if ((!true) && (false || true)) {
            v67 := GenS6{F0: (v6 % ((1) | 1)), F1: (v61 - v61)}
            _ = v67
        }
    }
    for fi68 := 0; fi68 < 1; fi68++ {
        for fi69 := 0; fi69 < 16; fi69++ {
            if (!true) {
                var v70 uint8 = (^(v6))
                v14 = 1
                fmt.Println(fi68, v64)
                _ = v70
            } else {
                v71 := GenS2{F0: 0}
                _ = v71
            }
            if true {
                fmt.Println(fi69)
                var v72 int = (fi69 >> ((65535) & 63))
                _ = v72
            } else {
                switch (v64 + v64) {
                case "foo bar":
                    v6 += v6
                case "\n\t":
                    v73 := GenS9{F0: 65535, F1: (v14 | -686)}
                    _ = v73
                }
                switch (v64 + v64) {
                case "a":
                    v74 := []int{-127, (^(v1))}
                    switch "\n\t" {
                    case "hello":
                        v2 += v2
                    case "a":
                        if len(v74) >= 2 {
                            v74 = v74[1:len(v74)]
                        }
                    default:
                        if true {
                            v61 = v61
                            v75 := GenS2{F0: -128}
                            switch 9223372036854775807 {
                            case 0:
                                _ = len(v74)
                                v3 = int64(v2)
                            case 142:
                                var v76 uint32 = v65
                                _ = v76
                            case -1:
                                var v77 int32 = v14
                                _ = v77
                            default:
                                var v78 int32 = v14
                                _ = v78
                            }
                            _ = v75
                        } else {
                            var v79 uint16 = 179
                            if false {
                                switch fi68 {
                                case 127:
                                    if (!(!((v6 | (v6 + 42)) != (v6 - v6)))) {
                                        v80 := GenS10{F0: 2}
                                        var v81 int32 = 2147483647
                                        v82 := GenS7{F0: v6, F1: (v65 / (((v65 << ((149) & 63))) | 1))}
                                        var v83 int64 = 507
                                        if (!((false && (!false)) && true)) {
                                            _ = len(v74)
                                            var v84 uint = 65535
                                            v14 = int32(v2)
                                            _ = v84
                                        } else {
                                            if (((100 > 0) && true) || (true || true)) {
                                                var v85 int16 = (-(v61))
                                                _ = v85
                                            } else {
                                                v86 := GenS6{F0: 127, F1: v61}
                                                var v87 float64 = 1e10
                                                fmt.Println(v79)
                                                _ = v86
                                                _ = v87
                                            }
                                            fmt.Println(v1)
                                        }
                                        _ = v80
                                        _ = v81
                                        _ = v82
                                        _ = v83
                                    } else {
                                        v6 = v6
                                    }
                                }
                            } else {
                                v65 -= (v65 << ((uint(v6)) & 63))
                            }
                            _ = v79
                        }
                    }
                    v65 = 333
                    _ = v74
                }
                var v88 int64 = (^(v3))
                switch (fi69 + (fi68 << ((18446744073709551615) & 63))) {
                case 649:
                    var v89 uint16 = 32767
                    _ = v89
                case -1:
                    var v90 float64 = ((3.14159 - 0.0) - 0.0)
                    _ = v90
                case -335:
                    if false {
                        if true {
                            v5 -= (1e10 + (((v5 + 1e10) / v5) * (v5 / v5)))
                        }
                    }
                default:
                    fmt.Println(v88, v3, v3)
                }
                _ = v88
            }
            var v91 float64 = 0.1
            _ = fi69
            _ = v91
        }
        _ = fi68
    }
    var v92 uint8 = 255
    for fi93 := 0; fi93 < 7; fi93++ {
        v3 -= int64(v14)
        v94 := []int{}
        if false {
            v6 = v92
            if (v65 == 2147483648) {
                switch (v64 + ("foo bar" + "")) {
                case "foo bar":
                    if true {
                        switch "hello" {
                        case "foo bar":
                            v92 = v92
                        case "":
                            if false {
                                _ = len(v94)
                                var v95 uint64 = uint64(v65)
                                _ = v95
                            }
                        }
                    }
                case "\n\t":
                    v96 := GenS3{F0: int32(v1)}
                    _ = v96
                }
            }
            fmt.Println(v61, v92, v2)
        }
        _ = fi93
        _ = v94
    }
    if false {
        var v97 int8 = 100
        switch v1 {
        case 127:
            v2 = 2
        case 9223372036854775807:
            if true {
                var v98 uint = 1
                _ = v98
            }
        case 4294967295:
            if true {
                v61 += (v61 * -32768)
            }
        }
        v99 := []int{(v1 / ((v1) | 1)), (v1 & 9223372036854775807)}
        _ = v97
        _ = v99
    } else {
        for fi100 := 0; fi100 < 7; fi100++ {
            v61 = 522
            fmt.Println(v4, v64, v2)
            var v101 float32 = v5
            _ = fi100
            _ = v101
        }
    }
    var v102 float64 = (1.5 - 0.0)
    v5 = float32(v102)
    if false {
        switch (v1 / ((v1) | 1)) {
        case -629:
            if false {
                if ((v2 << ((18446744073709551615) & 63)) > genFn1(v92, v64)) {
                    v103 := []int{v1, (v1 % ((127) | 1)), v1, (^(v1))}
                    v3 = v4
                    _ = v103
                } else {
                    switch -2147483648 {
                    case 632:
                        var v104 bool = true
                        _ = v104
                    case 4294967295:
                        v65 = 113
                    }
                    if (v1 < (v1 << ((255) & 63))) {
                        if true {
                            if true {
                                var v105 int16 = (v61 ^ v61)
                                fmt.Println(v2)
                                v61 -= (v61 + v105)
                                _ = v105
                            }
                            v6 += v6
                        }
                        fmt.Println(v61, v6, v61)
                    }
                    v106 := []int{v1}
                    _ = v106
                }
                var v107 uint32 = uint32(v92)
                if true {
                    v108 := GenS10{F0: 2}
                    switch v1 {
                    case 107:
                        v1 -= 9223372036854775807
                    case 1:
                        var v109 uint16 = 42
                        _ = v109
                    }
                    if (!false) {
                        v110 := GenS11{F0: v3, F1: (v1 / ((-858) | 1)), F2: 32767}
                        var v111 uint16 = 596
                        if (-1 > -68) {
                            switch v64 {
                            case "foo bar":
                                var v112 int16 = -1
                                _ = v112
                            case "":
                                v113 := []int{1, -9223372036854775808, v1}
                                _ = v113
                            case "hello":
                                fmt.Println(v92, v4, v6)
                            default:
                                if (v61 != (^(v61))) {
                                    var v114 uint = 629
                                    _ = v114
                                }
                            }
                            v5 = (v5 / 1.0)
                            v2 -= v2
                        }
                        _ = v110
                        _ = v111
                    } else {
                        for fi115 := 0; fi115 < 15; fi115++ {
                            if (v61 > v61) {
                                switch v64 {
                                case "hello":
                                    for fi116 := 0; fi116 < 3; fi116++ {
                                        switch ("a" + (((v64 + (v64 + "\n\t")) + "") + v64)) {
                                        case "a":
                                            switch fi116 {
                                            case 128:
                                                v117 := []int{}
                                                _ = v117
                                            case 175:
                                                var v118 int16 = 1
                                                _ = v118
                                            case 817:
                                                fmt.Println(v4, v4, v65)
                                            }
                                            v65 = (^(v107))
                                            var v119 int8 = 127
                                            _ = v119
                                        case "\n\t":
                                            v107 = 856
                                        case "foo bar":
                                            if true {
                                                if true {
                                                    v120 := []int{v1, (v1 ^ (^(v1))), (-(fi115)), (fi115 + (v1 << ((0) & 63)))}
                                                    fmt.Println(v2)
                                                    v121 := GenS12{F0: ((v64 + (v64 + v64)) + ""), F1: false, F2: v65}
                                                    v4 -= v3
                                                    v2 += genFn1((^(v92)), v64)
                                                    _ = v120
                                                    _ = v121
                                                }
                                                var v122 int64 = v3
                                                v6 -= (v6 >> ((42) & 63))
                                                v2 = 615
                                                _ = v122
                                            } else {
                                                if ((0.1 - v5) > v5) {
                                                    v102 += (0.1 + 1e10)
                                                    v123 := []int{(v1 / ((v1) | 1)), (fi116 >> ((2) & 63))}
                                                    switch v1 {
                                                    case 4294967295:
                                                        fmt.Println(fi116, v107)
                                                        if false {
                                                            var v124 float32 = 1.5
                                                            var v125 bool = true
                                                            v5 -= v124
                                                            _ = len(v123)
                                                            _ = v124
                                                            _ = v125
                                                        }
                                                    case 0:
                                                        v123 = append(v123, (fi115 | v1))
                                                    default:
                                                        v61 += v61
                                                    }
                                                    _ = v123
                                                }
                                                v102 = v102
                                            }
                                        default:
                                            v126 := []int{int(v2)}
                                            _ = v126
                                        }
                                        v4 = v3
                                        _ = fi116
                                    }
                                case "\n\t":
                                    var v127 int16 = 42
                                    _ = v127
                                case "":
                                    var v128 uint32 = (v65 - (^(v65)))
                                    _ = v128
                                }
                            } else {
                                if false {
                                    fmt.Println(v61, v64, v64)
                                }
                                if true {
                                    var v129 int16 = v61
                                    v107 -= v65
                                    switch "hello" {
                                    case "foo bar":
                                        var v130 uint64 = (v2 * 0)
                                        _ = v130
                                    }
                                    _ = v129
                                }
                            }
                            fmt.Println(v61)
                            var v131 uint32 = v65
                            _ = fi115
                            _ = v131
                        }
                        v132 := []int{}
                        for fi133 := 0; fi133 < 10; fi133++ {
                            var v134 uint64 = v2
                            switch v64 {
                            case "hello":
                                v135 := GenS5{F0: 128}
                                _ = v135
                            case "":
                                v4 -= (-(v3))
                            default:
                                if (582 == 42) {
                                    var v136 int64 = (v4 * v3)
                                    _ = v136
                                } else {
                                    fmt.Println(v14, v2, v134)
                                }
                            }
                            _ = fi133
                            _ = v134
                        }
                        _ = v132
                    }
                    _ = v108
                }
                _ = v107
            } else {
                v137 := GenS13{F0: ("foo bar" + v64)}
                var v138 float64 = (0.1 * ((v102 + v102) - (3.14159 * v102)))
                _ = v137
                _ = v138
            }
        case 2147483647:
            for fi139 := 0; fi139 < 12; fi139++ {
                v2 = genFn1((v92 ^ v92), v64)
                _ = fi139
            }
        case -1:
            v140 := GenS7{F0: (v6 + 0), F1: v65}
            _ = v140
        }
        var v141 bool = (v61 <= -212)
        _ = v141
    } else {
        switch "foo bar" {
        case "foo bar":
            switch (v1 & (v1 | (v1 | (v1 * int(v92))))) {
            case -9223372036854775808:
                v142 := []int{220, v1, (v1 >> ((402) & 63))}
                _ = v142
            case 1:
                v143 := GenS4{F0: true, F1: (v6 % ((28) | 1)), F2: -173}
                _ = v143
            case 128:
                v144 := []int{1, v1}
                _ = v144
            }
        case "hello":
            fmt.Println(v92, v1, v61)
        default:
            v145 := []int{(v1 / (((v1 + v1)) | 1)), v1}
            _ = v145
        }
    }
    for fi146 := 0; fi146 < 5; fi146++ {
        if false {
            switch (v1 * 1) {
            case 127:
                if true {
                    var v147 int32 = v14
                    switch v1 {
                    case 473:
                        fmt.Println(v65)
                    default:
                        var v148 uint8 = (v92 >> ((108) & 63))
                        _ = v148
                    }
                    if ((!true) && true) {
                        if (true || true) {
                            var v149 float64 = 3.14159
                            _ = v149
                        }
                    } else {
                        for fi150 := 0; fi150 < 15; fi150++ {
                            fmt.Println(v102)
                            fmt.Println(v1, v3, v102)
                            var v151 uint16 = 958
                            v152 := []int{fi150}
                            if true {
                                fmt.Println(v102, v64, v3)
                                switch v64 {
                                case "a":
                                    v153 := GenS14{F0: (v92 * v92)}
                                    _ = v153
                                case "\n\t":
                                    fmt.Println(v61, fi150, v147)
                                    var v154 uint16 = 874
                                    fmt.Println(v5, v4, v2)
                                    _ = v154
                                }
                            } else {
                                fmt.Println(v102)
                                if (genFn1((^(v92)), (v64 + ((((v64 + ("a" + v64)) + v64) + v64) + v64))) != uint64(v6)) {
                                    var v155 uint32 = (^(v65))
                                    var v156 float64 = float64(v5)
                                    v157 := GenS15{F0: (v3 ^ (v3 % (((v3 * (v3 | (v4 / ((int64(v2)) | 1))))) | 1))), F1: (v1 % (((^(fi146))) | 1))}
                                    _ = v155
                                    _ = v156
                                    _ = v157
                                }
                            }
                            _ = fi150
                            _ = v151
                            _ = v152
                        }
                    }
                    _ = v147
                } else {
                    for fi158 := 0; fi158 < 19; fi158++ {
                        v159 := GenS16{F0: 2, F1: v6, F2: 65535}
                        var v160 uint8 = 1
                        switch (("hello" + v64) + v64) {
                        case "":
                            fmt.Println(v4)
                        case "foo bar":
                            v161 := GenS17{F0: 82}
                            if ((v65 >= (v65 ^ 737)) || (!(!(2 >= 66)))) {
                                var v162 uint = 212
                                v163 := GenS18{F0: true}
                                v102 = 0.1
                                _ = v162
                                _ = v163
                            }
                            v164 := GenS11{F0: v3, F1: (v1 & v1), F2: 542}
                            var v165 int32 = -210
                            _ = v161
                            _ = v164
                            _ = v165
                        case "\n\t":
                            var v166 float32 = (((0.1 - float32(v102)) * (-2.25 - (0.1 / v5))) + (v5 * v5))
                            _ = v166
                        default:
                            if (845 != 65535) {
                                var v167 float32 = v5
                                if (v102 != (-2.25 / (-2.25 + float64(v5)))) {
                                    v168 := []int{(v1 * (fi146 >> ((2) & 63)))}
                                    if (!true) {
                                        v169 := GenS4{F0: false, F1: v6, F2: v14}
                                        _ = v169
                                    }
                                    _ = v168
                                }
                                _ = v167
                            } else {
                                v4 = int64(v2)
                                if (-1 <= 42) {
                                    v102 += (1.5 * (v102 / ((v102 / ((v102 - float64(float32(v102))) * v102)) - v102)))
                                    var v170 int32 = (-(v14))
                                    _ = v170
                                }
                                v171 := GenS11{F0: v4, F1: -9223372036854775808, F2: 170}
                                _ = v171
                            }
                        }
                        _ = fi158
                        _ = v159
                        _ = v160
                    }
                    fmt.Println(v6, v3)
                    for fi172 := 0; fi172 < 13; fi172++ {
                        if true {
                            fmt.Println(v2, v2)
                        } else {
                            v4 -= v3
                            v173 := GenS14{F0: 42}
                            fmt.Println(v5)
                            _ = v173
                        }
                        var v174 uint32 = (v65 / (((v65 | (v65 ^ v65))) | 1))
                        _ = fi172
                        _ = v174
                    }
                }
            case -416:
                v175 := GenS19{F0: 386}
                _ = v175
            default:
                fmt.Println(v2, v14)
            }
        } else {
            var v176 int64 = v4
            switch (v1 | (^(fi146))) {
            case -9223372036854775808:
                for fi177 := 0; fi177 < 18; fi177++ {
                    switch (fi146 + (-(v1))) {
                    case 4294967295:
                        fmt.Println(fi146, v5)
                    default:
                        v178 := []int{(fi146 | v1), v1, v1}
                        _ = v178
                    }
                    v179 := []int{v1, (fi177 % ((128) | 1)), fi177, v1}
                    _ = fi177
                    _ = v179
                }
            case 2147483647:
                v102 = v102
            }
            _ = v176
        }
        v180 := GenS3{F0: v14}
        var v181 float32 = 0.1
        _ = fi146
        _ = v180
        _ = v181
    }
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v102
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v63
    _ = v64
    _ = v65
    _ = v92
}
