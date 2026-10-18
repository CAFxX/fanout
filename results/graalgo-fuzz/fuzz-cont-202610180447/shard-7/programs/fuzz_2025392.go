package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 bool; F1 int32; F2 int32 }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8; F1 uint16 }
type GenS7 struct { F0 uint64; F1 int16; F2 int32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS11 struct { F0 bool; F1 uint; F2 string }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 int32; F1 int8 }
type GenS14 struct { F0 int16; F1 uint32; F2 int }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 uint64; F1 int64; F2 uint16 }
type GenS17 struct { F0 uint8; F1 uint64 }
type GenS18 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS19 struct { F0 int32; F1 string }
type GenS20 struct { F0 uint; F1 int }
type GenS21 struct { F0 int8; F1 int64; F2 int32 }
func genFn0() uint8 {
    var v5 float64 = 0.0
    if true {
        v6 := []int{}
        _ = v6
    } else {
        for fi7 := 0; fi7 < 20; fi7++ {
            v8 := GenS1{F0: true, F1: 0, F2: 2}
            var v9 float64 = (((v5 / v5) - (v5 - (v5 / (0.1 / v5)))) * 0.1)
            _ = fi7
            _ = v8
            _ = v9
        }
        v5 = v5
        v10 := GenS2{F0: ("" + "foo bar"), F1: 1}
        _ = v10
    }
    _ = v5
    return 128
}

func main() {
    var v1 uint8 = 163
    var v2 uint16 = 967
    var v3 int8 = 127
    var v4 uint32 = 0
    switch 337 {
    case 42:
        var v11 int8 = 127
        _ = v11
    case -472:
        for fi12 := 0; fi12 < 10; fi12++ {
            fmt.Println(v2, v4, v4)
            _ = fi12
        }
    default:
        var v13 string = ("" + "foo bar")
        _ = v13
    }
    var v14 uint16 = (^(v2))
    if (-1 != 32767) {
        switch "hello" {
        case "":
            v15 := GenS2{F0: (("\n\t" + "a") + "hello"), F1: 307}
            _ = v15
        case "hello":
            v16 := []int{-399}
            _ = v16
        case "a":
            if (false && true) {
                v2 = v14
                if (true && true) {
                    if (true && (-478 >= -1)) {
                        if (307 < 18446744073709551615) {
                            var v17 int = 128
                            if ((v4 / ((v4) | 1)) != v4) {
                                switch v17 {
                                case 59:
                                    v18 := GenS1{F0: true, F1: 18446744073709551615, F2: 42}
                                    _ = v18
                                }
                            } else {
                                switch (^(v17)) {
                                case 1:
                                    var v19 float64 = 3.14159
                                    _ = v19
                                default:
                                    v20 := GenS3{F0: (128 <= 61), F1: -2147483648, F2: 456}
                                    _ = v20
                                }
                                v17 -= -1
                                v2 = (^(v2))
                            }
                            if true {
                                v4 = v4
                            }
                            _ = v17
                        }
                        var v21 bool = false
                        for fi22 := 0; fi22 < 19; fi22++ {
                            v23 := []int{}
                            v2 = v2
                            fmt.Println(v2, v21, v4)
                            _ = fi22
                            _ = v23
                        }
                        _ = v21
                    }
                    for fi24 := 0; fi24 < 19; fi24++ {
                        v14 = 1
                        v1 -= v1
                        _ = fi24
                    }
                    for fi25 := 0; fi25 < 5; fi25++ {
                        for fi26 := 0; fi26 < 19; fi26++ {
                            if false {
                                if true {
                                    v1 += uint8(v4)
                                } else {
                                    v27 := GenS1{F0: (true && (fi26 != fi26)), F1: 0, F2: uint(v1)}
                                    v28 := []int{(^(fi26)), fi26}
                                    var v29 uint = 2
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                }
                                var v30 uint16 = v14
                                _ = v30
                            } else {
                                if true {
                                    v14 -= 0
                                }
                                v31 := GenS4{F0: (v14 - (v2 - (^(v14)))), F1: v3}
                                if false {
                                    v14 = v2
                                    var v32 uint64 = 0
                                    v33 := []int{32768, 522}
                                    switch 9223372036854775807 {
                                    case 32768:
                                        var v34 int32 = 42
                                        _ = v34
                                    case 0:
                                        if len(v33) > 0 {
                                            v33[0] = (fi25 | fi25)
                                        }
                                    case 32767:
                                        fmt.Println(fi26, v32)
                                    }
                                    v35 := GenS5{F0: 602}
                                    _ = v32
                                    _ = v33
                                    _ = v35
                                } else {
                                    v36 := GenS5{F0: -2147483648}
                                    switch ("a" + "hello") {
                                    case "":
                                        var v37 int = fi25
                                        _ = v37
                                    case "\n\t":
                                        var v38 uint = 4294967295
                                        _ = v38
                                    default:
                                        v39 := GenS1{F0: ("hello" <= "a"), F1: 819, F2: 2}
                                        _ = v39
                                    }
                                    _ = v36
                                }
                                _ = v31
                            }
                            _ = fi26
                        }
                        var v40 int = fi25
                        _ = fi25
                        _ = v40
                    }
                }
                if (true && false) {
                    v41 := []int{127, 2147483647, 194, 4294967295}
                    _ = v41
                } else {
                    var v42 uint = 4294967295
                    switch int(v1) {
                    case 42:
                        if ((v14 | (v14 / ((v2) | 1))) >= (v14 ^ uint16(v4))) {
                            v43 := []int{}
                            v44 := GenS6{F0: v3, F1: v14}
                            _ = v43
                            _ = v44
                        }
                    case 58:
                        fmt.Println(v3)
                    case -188:
                        var v45 int8 = 76
                        _ = v45
                    }
                    _ = v42
                }
            } else {
                v3 = v3
                fmt.Println(v1, v14, v3)
            }
        }
        fmt.Println(v4, v3, v4)
    } else {
        if ((v3 * 0) <= (v3 | (^(v3)))) {
            if true {
                switch "" {
                case "":
                    v46 := []int{32768}
                    _ = v46
                case "a":
                    v47 := []int{2147483647, -455, -2147483648}
                    _ = v47
                case "foo bar":
                    var v48 int16 = 164
                    _ = v48
                }
                switch 0 {
                case -2147483648:
                    v49 := []int{127, 1}
                    _ = v49
                }
                switch "a" {
                case "foo bar":
                    for fi50 := 0; fi50 < 17; fi50++ {
                        v51 := GenS7{F0: 391, F1: -1, F2: -622}
                        var v52 int16 = -865
                        if ((fi50 & (^(fi50))) >= (fi50 ^ 0)) {
                            switch int(v3) {
                            case 9223372036854775807:
                                var v53 string = "hello"
                                _ = v53
                            }
                            var v54 uint64 = 42
                            _ = v54
                        }
                        _ = fi50
                        _ = v51
                        _ = v52
                    }
                }
            } else {
                switch "a" {
                case "\n\t":
                    switch 2147483647 {
                    case 900:
                        for fi55 := 0; fi55 < 11; fi55++ {
                            fmt.Println(v3, fi55, v1)
                            v56 := []int{(fi55 % ((128) | 1)), -2147483648, (fi55 & fi55)}
                            _ = fi55
                            _ = v56
                        }
                    default:
                        v57 := []int{-1, 61, 4294967295}
                        _ = v57
                    }
                default:
                    v1 = 128
                }
                v58 := []int{-438}
                v59 := GenS4{F0: v14, F1: v3}
                _ = v58
                _ = v59
            }
        }
        var v60 int16 = -1
        var v61 float32 = ((((0.1 * float32(((float64(1.5) + 1.5) / 1.5))) - 1.5) * -2.25) / (float32(0.1) + -2.25))
        _ = v60
        _ = v61
    }
    switch -9223372036854775808 {
    case 0:
        v62 := []int{-1, -474, 58, 127}
        _ = v62
    case 2147483647:
        v2 -= (v2 % ((v14) | 1))
    case 860:
        v63 := GenS7{F0: 42, F1: -32768, F2: -2147483648}
        _ = v63
    default:
        var v64 int16 = 1000
        _ = v64
    }
    v65 := GenS5{F0: -2147483648}
    var v66 uint8 = v1
    if (18446744073709551615 < 461) {
        for fi67 := 0; fi67 < 8; fi67++ {
            v4 = v4
            if true {
                for fi68 := 0; fi68 < 12; fi68++ {
                    if (779 != 4294967295) {
                        var v69 int64 = -9223372036854775808
                        v70 := []int{(fi67 & (-(fi68))), 32768, 9223372036854775807, 2147483647}
                        _ = v69
                        _ = v70
                    } else {
                        v71 := GenS6{F0: (v3 & v3), F1: (^(v14))}
                        v72 := []int{fi68, fi68}
                        var v73 int8 = (v3 >> ((18446744073709551615) & 63))
                        _ = v71
                        _ = v72
                        _ = v73
                    }
                    if false {
                        var v74 uint32 = v4
                        _ = v74
                    }
                    _ = fi68
                }
                var v75 uint32 = 65535
                if (!(!false)) {
                    v76 := GenS8{F0: 32767}
                    switch (fi67 + (^(fi67))) {
                    case 801:
                        var v77 int8 = v3
                        _ = v77
                    default:
                        v78 := GenS9{F0: 0}
                        _ = v78
                    }
                    _ = v76
                } else {
                    switch ("hello" + (("foo bar" + ("a" + "foo bar")) + ("foo bar" + ("\n\t" + "foo bar")))) {
                    case "":
                        v4 = (v4 & uint32(v14))
                    default:
                        fmt.Println(v75)
                    }
                    if (!(-9223372036854775808 == fi67)) {
                        if ((1 > v66) || ((^(v75)) > v75)) {
                            if true {
                                if ("foo bar" >= (("" + "foo bar") + ("foo bar" + "a"))) {
                                    switch (fi67 * fi67) {
                                    case 946:
                                        v79 := []int{}
                                        _ = v79
                                    default:
                                        switch ((("\n\t" + "hello") + "") + ("\n\t" + "")) {
                                        case "":
                                            if false {
                                                switch "foo bar" {
                                                case "foo bar":
                                                    var v80 uint8 = v66
                                                    _ = v80
                                                }
                                                fmt.Println(v66, fi67, fi67)
                                            }
                                        case "foo bar":
                                            var v81 int16 = -131
                                            _ = v81
                                        }
                                    }
                                    for fi82 := 0; fi82 < 5; fi82++ {
                                        var v83 int32 = -591
                                        _ = fi82
                                        _ = v83
                                    }
                                }
                                v84 := []int{(fi67 << ((4294967295) & 63))}
                                fmt.Println(v14, v66)
                                _ = v84
                            } else {
                                v75 = uint32(v66)
                                switch "foo bar" {
                                case "\n\t":
                                    if ((3.14159 * float32(float64(((float32((0.0 - 1e10)) + 0.1) * (-2.25 * 3.14159))))) >= float32((0.1 / -2.25))) {
                                        v85 := GenS5{F0: 2147483647}
                                        _ = v85
                                    }
                                }
                            }
                            v86 := GenS10{F0: 382, F1: v75, F2: 542}
                            switch "a" {
                            case "foo bar":
                                fmt.Println(fi67, v2)
                            case "\n\t":
                                switch (-(fi67)) {
                                case 0:
                                    v87 := []int{fi67, -1, fi67, fi67}
                                    _ = v87
                                case -1:
                                    for fi88 := 0; fi88 < 2; fi88++ {
                                        v89 := GenS4{F0: v14, F1: (v3 << ((676) & 63))}
                                        v3 -= (v3 * v3)
                                        _ = fi88
                                        _ = v89
                                    }
                                default:
                                    v4 -= 598
                                }
                            case "a":
                                switch fi67 {
                                case 127:
                                    v90 := GenS11{F0: false, F1: 417, F2: ""}
                                    _ = v90
                                case 32767:
                                    var v91 int64 = 9223372036854775807
                                    _ = v91
                                case -1:
                                    var v92 uint64 = 2
                                    _ = v92
                                }
                            }
                            _ = v86
                        } else {
                            for fi93 := 0; fi93 < 9; fi93++ {
                                var v94 uint8 = v1
                                v4 += v75
                                _ = fi93
                                _ = v94
                            }
                        }
                        v2 = v2
                        v75 -= (v4 % ((4294967295) | 1))
                    } else {
                        v95 := GenS12{F0: -2147483648}
                        var v96 int32 = -1
                        v97 := []int{-355, (fi67 & 4294967295)}
                        _ = v95
                        _ = v96
                        _ = v97
                    }
                }
                _ = v75
            } else {
                for fi98 := 0; fi98 < 3; fi98++ {
                    var v99 int64 = 0
                    v100 := []int{(-(fi98)), (-(fi98)), fi67, (fi98 << ((2) & 63))}
                    _ = fi98
                    _ = v100
                    _ = v99
                }
            }
            if true {
                var v101 bool = true
                _ = v101
            } else {
                v2 = 65535
                fmt.Println(v3, v14, v4)
            }
            _ = fi67
        }
        var v102 uint = 65535
        _ = v102
    } else {
        var v103 bool = (455 < 933)
        if v103 {
            if (((1e10 != 0.1) || false) || (1 == 768)) {
                if v103 {
                    var v104 uint64 = 4294967295
                    for fi105 := 0; fi105 < 15; fi105++ {
                        var v106 uint16 = uint16(v104)
                        switch fi105 {
                        case -706:
                            for fi107 := 0; fi107 < 19; fi107++ {
                                var v108 int8 = v3
                                v109 := GenS13{F0: int32(v4), F1: (v3 & (v3 & 100))}
                                _ = fi107
                                _ = v108
                                _ = v109
                            }
                        default:
                            var v110 uint32 = (v4 / (((v4 | (v4 / ((2) | 1)))) | 1))
                            _ = v110
                        }
                        v3 += (v3 >> ((65) & 63))
                        _ = fi105
                        _ = v106
                    }
                    for fi111 := 0; fi111 < 13; fi111++ {
                        for fi112 := 0; fi112 < 17; fi112++ {
                            var v113 float32 = 3.14159
                            var v114 int8 = v3
                            _ = fi112
                            _ = v113
                            _ = v114
                        }
                        switch "foo bar" {
                        case "":
                            var v115 string = ("a" + "hello")
                            _ = v115
                        }
                        v66 -= 0
                        _ = fi111
                    }
                    _ = v104
                }
            } else {
                var v116 int16 = -32768
                v66 += v1
                _ = v116
            }
        }
        switch "hello" {
        case "a":
            for fi117 := 0; fi117 < 5; fi117++ {
                v3 -= (-(v3))
                _ = fi117
            }
        case "\n\t":
            switch "\n\t" {
            case "\n\t":
                var v118 string = "a"
                _ = v118
            case "":
                if (v103 && (!v103)) {
                    if (!false) {
                        if false {
                            switch "a" {
                            case "foo bar":
                                for fi119 := 0; fi119 < 5; fi119++ {
                                    v2 = v2
                                    v66 = v66
                                    _ = fi119
                                }
                            case "":
                                fmt.Println(v1, v4)
                            case "hello":
                                v120 := GenS14{F0: 269, F1: 200, F2: int(v1)}
                                _ = v120
                            default:
                                v103 = v103
                            }
                        } else {
                            for fi121 := 0; fi121 < 14; fi121++ {
                                for fi122 := 0; fi122 < 13; fi122++ {
                                    var v123 uint16 = 467
                                    v66 = (^(v1))
                                    _ = fi122
                                    _ = v123
                                }
                                switch ("\n\t" + ("a" + "\n\t")) {
                                case "":
                                    v3 -= v3
                                case "\n\t":
                                    if (32767 >= 1) {
                                        fmt.Println(v1)
                                    }
                                case "foo bar":
                                    fmt.Println(v3, v1, fi121)
                                default:
                                    v124 := GenS15{F0: 1}
                                    _ = v124
                                }
                                _ = fi121
                            }
                            var v125 int32 = 1
                            var v126 float64 = ((-2.25 * (((1e10 + 0.0) / -2.25) / ((0.1 + float64(1e10)) / 1.0))) * 0.0)
                            _ = v125
                            _ = v126
                        }
                        if false {
                            if v103 {
                                v103 = (!((!v103) && v103))
                                var v127 uint16 = v2
                                _ = v127
                            } else {
                                var v128 uint8 = genFn0()
                                v129 := []int{}
                                _ = v128
                                _ = v129
                            }
                            for fi130 := 0; fi130 < 10; fi130++ {
                                if v103 {
                                    if false {
                                        var v131 string = ""
                                        _ = v131
                                    }
                                    fmt.Println(v2)
                                    v132 := GenS16{F0: 2, F1: -1, F2: (v14 + 65535)}
                                    _ = v132
                                } else {
                                    v133 := GenS17{F0: v66, F1: 281}
                                    v103 = v103
                                    _ = v133
                                }
                                _ = fi130
                            }
                            v134 := []int{1, -42}
                            _ = v134
                        } else {
                            v1 -= (v66 >> ((2) & 63))
                            if v103 {
                                v66 = v1
                            } else {
                                if false {
                                    v2 += (v14 - v14)
                                    var v135 float64 = (3.14159 + (1e10 + 0.1))
                                    _ = v135
                                } else {
                                    var v136 string = "\n\t"
                                    v1 = (v1 + 174)
                                    var v137 uint16 = v2
                                    _ = v136
                                    _ = v137
                                }
                            }
                        }
                    } else {
                        var v138 uint32 = v4
                        var v139 uint16 = v2
                        switch 32767 {
                        case 32767:
                            if ((("" + "hello") + "hello") >= "") {
                                v140 := []int{}
                                var v141 uint64 = uint64(v138)
                                var v142 float32 = (0.0 / (3.14159 - 0.1))
                                _ = v140
                                _ = v141
                                _ = v142
                            } else {
                                fmt.Println(v14, v2, v2)
                                for fi143 := 0; fi143 < 2; fi143++ {
                                    var v144 uint8 = 140
                                    v145 := GenS11{F0: v103, F1: 2, F2: (("hello" + "\n\t") + "hello")}
                                    var v146 uint16 = v14
                                    _ = fi143
                                    _ = v144
                                    _ = v145
                                    _ = v146
                                }
                                v147 := GenS17{F0: v1, F1: 2}
                                _ = v147
                            }
                        case 392:
                            v148 := GenS11{F0: (("" + ("hello" + (("foo bar" + ("a" + (("a" + "foo bar") + ("" + ("hello" + ("\n\t" + "\n\t")))))) + "a"))) != ("\n\t" + (("hello" + "hello") + "a"))), F1: 4294967295, F2: "hello"}
                            _ = v148
                        default:
                            var v149 int = 185
                            _ = v149
                        }
                        _ = v138
                        _ = v139
                    }
                    if ("foo bar" != "") {
                        if (2 == 144) {
                            if ((float64(1e10) <= 0.0) || (v103 || v103)) {
                                var v150 uint32 = v4
                                var v151 uint16 = v2
                                for fi152 := 0; fi152 < 3; fi152++ {
                                    v153 := []int{fi152, int(v66), fi152}
                                    v154 := GenS17{F0: (v1 << ((31) & 63)), F1: 2}
                                    v155 := GenS10{F0: -9223372036854775808, F1: (v150 % ((v150) | 1)), F2: 0}
                                    _ = fi152
                                    _ = v153
                                    _ = v154
                                    _ = v155
                                }
                                _ = v150
                                _ = v151
                            } else {
                                if (42 > 339) {
                                    switch "" {
                                    case "hello":
                                        v2 = (v14 + (^(v2)))
                                    default:
                                        var v156 string = ("hello" + (("\n\t" + ("foo bar" + "a")) + ("hello" + (("foo bar" + "\n\t") + ("foo bar" + "a")))))
                                        _ = v156
                                    }
                                    v2 -= (v14 ^ (v2 * (v2 ^ v2)))
                                    fmt.Println(v103)
                                }
                            }
                            for fi157 := 0; fi157 < 19; fi157++ {
                                v158 := []int{}
                                var v159 uint = 22
                                if (32767 > int16(v14)) {
                                    v158 = append(v158, int(v3))
                                    if len(v158) >= 2 {
                                        v158 = v158[1:len(v158)]
                                    }
                                } else {
                                    v1 = 145
                                    switch "" {
                                    case "":
                                        var v160 uint8 = (v1 << (((v159 & (v159 ^ (v159 / ((v159) | 1))))) & 63))
                                        _ = v160
                                    case "foo bar":
                                        var v161 string = ""
                                        _ = v161
                                    case "\n\t":
                                        if (-1 >= 554) {
                                            v66 = (v66 / ((v1) | 1))
                                            for fi162 := 0; fi162 < 3; fi162++ {
                                                switch 9223372036854775807 {
                                                case 539:
                                                    _ = len(v158)
                                                default:
                                                    switch 2147483647 {
                                                    case 32768:
                                                        if v103 {
                                                            var v163 uint8 = uint8(v4)
                                                            _ = v163
                                                        }
                                                    }
                                                }
                                                if len(v158) > 0 {
                                                    v158[0] = (fi162 & 32768)
                                                }
                                                _ = fi162
                                            }
                                        }
                                    default:
                                        var v164 float64 = (((-2.25 - (1e10 + 1.5)) + 0.0) + 1.5)
                                        _ = v164
                                    }
                                    _ = len(v158)
                                }
                                _ = fi157
                                _ = v158
                                _ = v159
                            }
                            v165 := GenS18{F0: genFn0(), F1: 2, F2: v66}
                            _ = v165
                        } else {
                            v3 = v3
                            switch 844 {
                            case 32768:
                                v1 -= v1
                            }
                            if true {
                                fmt.Println(v3, v103)
                                fmt.Println(v4, v1)
                                switch "a" {
                                case "\n\t":
                                    if (v3 < -28) {
                                        if v103 {
                                            for fi166 := 0; fi166 < 14; fi166++ {
                                                for fi167 := 0; fi167 < 9; fi167++ {
                                                    if v103 {
                                                        v168 := []int{fi166}
                                                        var v169 uint8 = (v1 + (v1 % ((v66) | 1)))
                                                        _ = v168
                                                        _ = v169
                                                    } else {
                                                        switch fi167 {
                                                        case 0:
                                                            v170 := GenS19{F0: 1, F1: "a"}
                                                            _ = v170
                                                        case 42:
                                                            var v171 int = fi166
                                                            _ = v171
                                                        default:
                                                            if false {
                                                                fmt.Println(v1)
                                                                var v172 bool = ((42 > 9223372036854775807) && (v2 < 411))
                                                                v173 := GenS20{F0: 336, F1: fi167}
                                                                _ = v172
                                                                _ = v173
                                                            } else {
                                                                v14 = v14
                                                                v2 = v2
                                                                v174 := GenS15{F0: -737}
                                                                _ = v174
                                                            }
                                                        }
                                                    }
                                                    var v175 uint64 = 14
                                                    _ = fi167
                                                    _ = v175
                                                }
                                                fmt.Println(v2)
                                                _ = fi166
                                            }
                                            switch ("\n\t" + "\n\t") {
                                            case "":
                                                v103 = v103
                                            case "hello":
                                                fmt.Println(v3, v14, v1)
                                            default:
                                                for fi176 := 0; fi176 < 20; fi176++ {
                                                    fmt.Println(fi176)
                                                    _ = fi176
                                                }
                                            }
                                            switch "a" {
                                            case "a":
                                                v177 := GenS21{F0: (v3 / (((-(v3))) | 1)), F1: -646, F2: -1}
                                                _ = v177
                                            case "hello":
                                                v66 += v66
                                            }
                                        }
                                    }
                                case "hello":
                                    switch -739 {
                                    case -1:
                                        if (v103 && v103) {
                                            var v178 uint32 = (v4 + (v4 | (v4 & v4)))
                                            _ = v178
                                        }
                                    }
                                default:
                                    fmt.Println(v4, v1, v2)
                                }
                            } else {
                                if false {
                                    if (!(true && (87 != (^(v3))))) {
                                        fmt.Println(v14, v4)
                                        var v179 float32 = (-2.25 - 0.1)
                                        _ = v179
                                    }
                                    for fi180 := 0; fi180 < 13; fi180++ {
                                        v181 := []int{}
                                        if len(v181) >= 2 {
                                            v181 = v181[1:len(v181)]
                                        }
                                        _ = fi180
                                        _ = v181
                                    }
                                    v3 = v3
                                } else {
                                    fmt.Println(v4)
                                    if false {
                                        var v182 uint64 = 418
                                        var v183 int16 = 357
                                        var v184 int64 = int64(v182)
                                        _ = v182
                                        _ = v183
                                        _ = v184
                                    } else {
                                        v1 = (v66 - genFn0())
                                    }
                                    fmt.Println(v66, v2, v66)
                                }
                                var v185 float32 = 1.5
                                _ = v185
                            }
                        }
                        v186 := []int{-506, 32767, 352, 2147483647}
                        for fi187 := 0; fi187 < 9; fi187++ {
                            fmt.Println(v3, v66, v14)
                            v14 -= v2
                            _ = fi187
                        }
                        _ = v186
                    }
                }
            }
        case "":
            if (-951 <= -688) {
                for fi188 := 0; fi188 < 16; fi188++ {
                    fmt.Println(v4, v14)
                    if v103 {
                        fmt.Println(v2, v4, v14)
                        v1 = (^(v66))
                        v1 -= v1
                    }
                    _ = fi188
                }
            }
        }
        _ = v103
    }
    v14 += 255
    if true {
        if (-1 <= 196) {
            v189 := []int{483, 1, 0}
            v66 += genFn0()
            _ = v189
        } else {
            var v190 uint16 = v14
            _ = v190
        }
        for fi191 := 0; fi191 < 14; fi191++ {
            v3 = v3
            fmt.Println(v66)
            _ = fi191
        }
    } else {
        v192 := []int{-9223372036854775808, 1, 2147483647}
        var v193 string = ""
        _ = v192
        _ = v193
    }
    v194 := []int{-2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v194:", len(v194), v194)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v14
    _ = v194
    _ = v2
    _ = v3
    _ = v4
    _ = v65
    _ = v66
}
