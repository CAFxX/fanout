package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 int; F2 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint8; F1 int16; F2 int32 }
type GenS8 struct { F0 bool; F1 int64 }
type GenS9 struct { F0 int64; F1 bool }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int16; F1 uint64 }
func genFn0(p0 uint64, p1 bool, p2 int16) uint8 {
    v7 := []int{32768, -29}
    var v8 uint = 34
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 170
}
func genFn1(p0 uint8, p1 int) uint {
    fmt.Println(p0, p0)
    switch 127 {
    case 2147483647:
        p1 = p1
    }
    _ = p0
    _ = p1
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 0
    v3 := make(map[int]bool)
    v4 := make(map[int]bool)
    var v5 int32 = 2147483647
    var v6 uint32 = v2
    var v9 int = 128
    v6 = 404
    switch -2147483648 {
    case 969:
        if (!true) {
            v10 := GenS1{F0: genFn0(1, (1 <= -1), -1)}
            _ = v10
        } else {
            fmt.Println(v9, v9, v9)
            var v11 string = ("a" + "")
            _ = v11
        }
    case 32767:
        switch ("\n\t" + "a") {
        case "hello":
            switch (v9 << ((0) & 63)) {
            case 127:
                var v12 uint8 = 239
                _ = v12
            default:
                v13 := GenS1{F0: 93}
                _ = v13
            }
        default:
            v6 -= 2
        }
    case 1:
        if (3.14159 == (1.5 / ((3.14159 - 0.1) / (1e10 + 1e10)))) {
            for fi14 := 0; fi14 < 17; fi14++ {
                v5 -= v5
                var v15 int = (-(v9))
                var v16 uint32 = v6
                _ = fi14
                _ = v15
                _ = v16
            }
            fmt.Println(v6, v5, v9)
        } else {
            v17 := GenS2{F0: 790}
            var v18 uint8 = 21
            _ = v17
            _ = v18
        }
    default:
        if true {
            var v19 uint64 = 18446744073709551615
            _ = v19
        }
    }
    var v20 uint16 = 185
    for fi21 := 0; fi21 < 7; fi21++ {
        var v22 uint32 = (v2 * v2)
        switch "\n\t" {
        case "hello":
            if (v20 == v20) {
                var v23 int8 = -112
                _ = v23
            }
        case "foo bar":
            if true {
                if false {
                    switch "" {
                    case "":
                        if false {
                            var v24 uint64 = 2
                            switch fi21 {
                            case 42:
                                if (!(uint8(v6) != 26)) {
                                    var v25 uint = 65535
                                    v26 := GenS1{F0: 212}
                                    _ = v25
                                    _ = v26
                                } else {
                                    switch "hello" {
                                    case "\n\t":
                                        v27 := []int{int(v24)}
                                        _ = v27
                                    case "foo bar":
                                        fmt.Println(v5, v20, v20)
                                    default:
                                        fmt.Println(v22, v22)
                                    }
                                    var v28 uint = 908
                                    if ("a" >= "\n\t") {
                                        var v29 bool = true
                                        var v30 uint8 = 210
                                        _ = v29
                                        _ = v30
                                    } else {
                                        v24 -= (v24 ^ v24)
                                        v31 := v4[(-(fi21))]
                                        _ = v31
                                    }
                                    _ = v28
                                }
                            case 0:
                                if true {
                                    var v32 int16 = 945
                                    v33 := []int{v9, fi21, -614}
                                    _ = v32
                                    _ = v33
                                }
                            }
                            v34 := []int{-29, fi21, -1}
                            _ = v24
                            _ = v34
                        }
                    case "\n\t":
                        for fi35 := 0; fi35 < 5; fi35++ {
                            if (((127 >= 66) && false) || false) {
                                switch "hello" {
                                case "hello":
                                    fmt.Println(fi21)
                                case "a":
                                    var v36 float32 = (1e10 * 0.0)
                                    _ = v36
                                case "\n\t":
                                    var v37 string = "hello"
                                    _ = v37
                                default:
                                    v38 := []int{fi35, (v9 / (((fi35 % ((v9) | 1))) | 1))}
                                    _ = v38
                                }
                                delete(v4, (fi21 >> ((971) & 63)))
                                v39 := GenS3{F0: -327, F1: fi35, F2: 127}
                                _ = v39
                            } else {
                                var v40 int16 = -132
                                _ = v40
                            }
                            _ = fi35
                        }
                    }
                    fmt.Println(v20, v22)
                    for fi41 := 0; fi41 < 11; fi41++ {
                        v6 = (v22 / ((65535) | 1))
                        v42 := []int{fi21}
                        _ = fi41
                        _ = v42
                    }
                }
                var v43 uint32 = (v2 >> ((uint(v22)) & 63))
                var v44 int32 = (v5 | (v5 >> ((255) & 63)))
                _ = v43
                _ = v44
            }
        default:
            switch (v9 - fi21) {
            case 32768:
                v45 := GenS3{F0: 32767, F1: 42, F2: 2}
                _ = v45
            case 98:
                fmt.Println(v22, v9, v6)
            default:
                if true {
                    fmt.Println(fi21, v5, fi21)
                    v6 = (v6 % ((v6) | 1))
                }
            }
        }
        if (106 != 1) {
            fmt.Println(fi21)
            v46 := []int{}
            _ = len(v46)
            _ = v46
        } else {
            _ = len(v1)
            if (42 < -78) {
                for fi47 := 0; fi47 < 15; fi47++ {
                    v2 = (v2 * (v22 >> ((207) & 63)))
                    switch "" {
                    case "\n\t":
                        if (fi47 != 128) {
                            var v48 uint64 = 485
                            _ = v48
                        }
                    case "a":
                        fmt.Println(v5)
                    case "foo bar":
                        if (false || true) {
                            switch -9223372036854775808 {
                            case 0:
                                var v49 uint = 1
                                _ = v49
                            default:
                                v50 := GenS2{F0: -32768}
                                _ = v50
                            }
                            var v51 bool = false
                            if v51 {
                                var v52 uint16 = v20
                                v4[(^(fi47))] = true
                                if v51 {
                                    switch (-(v9)) {
                                    case -221:
                                        v51 = (0.1 == 1e10)
                                    default:
                                        var v53 int32 = -1
                                        _ = v53
                                    }
                                    v54 := GenS3{F0: 974, F1: -1, F2: 255}
                                    switch "foo bar" {
                                    case "a":
                                        fmt.Println(v6, v2)
                                    case "hello":
                                        var v55 uint16 = 484
                                        _ = v55
                                    default:
                                        var v56 float64 = 3.14159
                                        _ = v56
                                    }
                                    if (float64(((0.0 * -2.25) - (0.0 / 0.1))) == 0.1) {
                                        v57 := GenS4{F0: v5}
                                        fmt.Println(fi21, v51)
                                        if (!true) {
                                            var v58 int64 = 1
                                            var v59 int64 = -9223372036854775808
                                            var v60 int32 = v5
                                            switch "a" {
                                            case "a":
                                                var v61 uint8 = genFn0(2, v51, -1)
                                                _ = v61
                                            default:
                                                if v51 {
                                                    fmt.Println(v2, v20, v22)
                                                    if v51 {
                                                        v62 := v1[536]
                                                        delete(v4, fi21)
                                                        _ = v62
                                                    }
                                                    var v63 float64 = 1e10
                                                    v5 -= (v60 << ((129) & 63))
                                                    _ = v63
                                                } else {
                                                    v59 += v58
                                                    switch fi47 {
                                                    case 32768:
                                                        var v64 int16 = -32768
                                                        _ = v64
                                                    case 2147483647:
                                                        fmt.Println(v22, fi21, v20)
                                                    default:
                                                        v65 := GenS5{F0: v58, F1: 42}
                                                        _ = v65
                                                    }
                                                }
                                            }
                                            v66 := GenS1{F0: 221}
                                            _ = v58
                                            _ = v59
                                            _ = v60
                                            _ = v66
                                        }
                                        _ = v57
                                    } else {
                                        switch ("" + "hello") {
                                        case "foo bar":
                                            var v67 float64 = 3.14159
                                            _ = v67
                                        case "hello":
                                            if false {
                                                if ((v2 == 4294967295) || (v9 >= int(v5))) {
                                                    v68 := GenS5{F0: -1, F1: 4294967295}
                                                    v69 := []int{(^(v9)), fi47, int(v5), fi47}
                                                    var v70 int = fi47
                                                    _ = v68
                                                    _ = v69
                                                    _ = v70
                                                }
                                                v52 += v20
                                                var v71 int8 = 127
                                                v5 = (^(v5))
                                                var v72 uint8 = 1
                                                var v73 bool = false
                                                _ = v71
                                                _ = v72
                                                _ = v73
                                            }
                                        }
                                        var v74 int32 = v5
                                        _ = v74
                                    }
                                    v75 := []int{(fi21 / (((^(v9))) | 1))}
                                    _ = v54
                                    _ = v75
                                } else {
                                    if (!v51) {
                                        var v76 int16 = 0
                                        fmt.Println(fi47, v5)
                                        _ = v76
                                    } else {
                                        v77 := GenS5{F0: -9223372036854775808, F1: 1}
                                        var v78 uint32 = v22
                                        _ = v77
                                        _ = v78
                                    }
                                    fmt.Println(v22)
                                    var v79 uint32 = v2
                                    _ = v79
                                }
                                _ = v52
                            } else {
                                v22 += v6
                                v80 := GenS1{F0: 215}
                                var v81 string = ""
                                _ = v80
                                _ = v81
                            }
                            if (true && v51) {
                                if v51 {
                                    switch "" {
                                    case "\n\t":
                                        if (float64((float32(-2.25) + 1e10)) >= ((((0.1 / 1e10) * 1e10) / -2.25) - -2.25)) {
                                            fmt.Println(v51, v20, v2)
                                        }
                                    case "hello":
                                        fmt.Println(v9, v5)
                                    case "":
                                        var v82 string = "hello"
                                        _ = v82
                                    default:
                                        v51 = (-2147483648 == fi21)
                                    }
                                    v83 := GenS3{F0: -32768, F1: fi21, F2: 128}
                                    v84 := GenS4{F0: v5}
                                    if false {
                                        var v85 int8 = 21
                                        var v86 float32 = (0.0 / float32((1.5 + (1.5 * 0.1))))
                                        var v87 uint = genFn1(1, (-(fi47)))
                                        _ = v85
                                        _ = v86
                                        _ = v87
                                    }
                                    switch ((("a" + ("\n\t" + "foo bar")) + "\n\t") + ("foo bar" + (("foo bar" + (("foo bar" + ("foo bar" + "foo bar")) + "")) + ("hello" + "")))) {
                                    case "a":
                                        v6 -= (v6 | (v2 ^ v22))
                                    case "":
                                        var v88 uint = 2
                                        _ = v88
                                    case "\n\t":
                                        v89 := GenS2{F0: 42}
                                        _ = v89
                                    default:
                                        fmt.Println(v51, v2)
                                    }
                                    _ = v83
                                    _ = v84
                                } else {
                                    if ((-(v5)) <= (-(v5))) {
                                        if (("" + "foo bar") <= ((("a" + "foo bar") + "a") + ("hello" + "a"))) {
                                            var v90 float64 = 0.1
                                            var v91 uint32 = 4294967295
                                            var v92 float64 = v90
                                            _ = v90
                                            _ = v91
                                            _ = v92
                                        }
                                        if (v51 && (2 <= genFn0(0, v51, -142))) {
                                            var v93 int32 = 647
                                            fmt.Println(fi21)
                                            var v94 float64 = 3.14159
                                            _ = v93
                                            _ = v94
                                        } else {
                                            var v95 uint64 = 0
                                            _ = v95
                                        }
                                    } else {
                                        fmt.Println(v20)
                                        v96 := []int{32768, fi21, fi47}
                                        _ = v96
                                    }
                                    v97 := []int{}
                                    switch "\n\t" {
                                    case "hello":
                                        v9 = (^(fi21))
                                    case "foo bar":
                                        v5 -= 0
                                    }
                                    _ = v97
                                }
                            }
                            v9 += (fi47 - (v9 >> ((genFn1(genFn0(18446744073709551615, (genFn0(490, v51, 32767) > 128), int16(v5)), fi47)) & 63)))
                            _ = v51
                        }
                    }
                    _ = fi47
                }
            } else {
                var v98 uint = 4294967295
                _ = v98
            }
            switch (v9 >> ((4294967295) & 63)) {
            case 195:
                v99 := []int{(v9 & fi21), fi21, fi21}
                _ = v99
            }
        }
        _ = fi21
        _ = v22
    }
    if true {
        for fi100 := 0; fi100 < 9; fi100++ {
            var v101 bool = false
            for fi102 := 0; fi102 < 11; fi102++ {
                v6 -= (v2 - (v2 - (v2 - v6)))
                _ = fi102
            }
            if v101 {
                var v103 float64 = 0.0
                v104 := GenS6{F0: 18446744073709551615}
                v2 = (v6 + v6)
                _ = v103
                _ = v104
            }
            _ = fi100
            _ = v101
        }
        if true {
            for fi105 := 0; fi105 < 9; fi105++ {
                var v106 int32 = v5
                var v107 int = 127
                var v108 int32 = 0
                _ = fi105
                _ = v106
                _ = v107
                _ = v108
            }
        } else {
            if true {
                for fi109 := 0; fi109 < 17; fi109++ {
                    v2 = (^(v2))
                    fmt.Println(v5)
                    v3[0] = ((true && false) && (0 <= 125))
                    _ = fi109
                }
            }
            v2 -= v6
        }
        if (4294967295 != 42) {
            if (1.5 <= (0.0 - (1.5 - float32(0.0)))) {
                var v110 string = ""
                for fi111 := 0; fi111 < 16; fi111++ {
                    for fi112 := 0; fi112 < 15; fi112++ {
                        var v113 uint = 18446744073709551615
                        v114 := GenS3{F0: 1, F1: fi111, F2: 2}
                        switch "" {
                        case "foo bar":
                            var v115 int8 = -23
                            _ = v115
                        case "hello":
                            v9 = 9223372036854775807
                        case "a":
                            fmt.Println(fi111)
                        default:
                            var v116 int64 = 0
                            v117 := v4[fi112]
                            _ = v116
                            _ = v117
                        }
                        _ = fi112
                        _ = v113
                        _ = v114
                    }
                    _ = fi111
                }
                for fi118 := 0; fi118 < 12; fi118++ {
                    var v119 uint32 = v6
                    _ = fi118
                    _ = v119
                }
                _ = v110
            }
            if false {
                if (-884 >= 1) {
                    var v120 bool = (!(3.14159 < 1.5))
                    for fi121 := 0; fi121 < 14; fi121++ {
                        v122 := []int{v9}
                        for fi123 := 0; fi123 < 15; fi123++ {
                            v9 = -2147483648
                            var v124 int16 = 1000
                            _ = fi123
                            _ = v124
                        }
                        if v120 {
                            if (v120 || (v120 && true)) {
                                var v125 int16 = 1
                                delete(v3, 2147483647)
                                _ = v125
                            }
                            fmt.Println(v2, v6)
                            switch "foo bar" {
                            case "hello":
                                if (9223372036854775807 == 1) {
                                    var v126 uint64 = 540
                                    _ = v126
                                } else {
                                    for fi127 := 0; fi127 < 5; fi127++ {
                                        v128 := GenS7{F0: 128, F1: 681, F2: (v5 << ((2) & 63))}
                                        fmt.Println(fi127, fi127)
                                        _ = fi127
                                        _ = v128
                                    }
                                    for fi129 := 0; fi129 < 9; fi129++ {
                                        switch (("\n\t" + "a") + "foo bar") {
                                        case "a":
                                            v120 = v120
                                        case "hello":
                                            _ = len(v1)
                                        default:
                                            fmt.Println(fi129, fi121, v6)
                                            fmt.Println(v2)
                                            delete(v1, v2)
                                        }
                                        var v130 uint32 = (v2 << ((792) & 63))
                                        var v131 uint = 255
                                        _ = fi129
                                        _ = v130
                                        _ = v131
                                    }
                                }
                            case "":
                                var v132 int16 = -32768
                                _ = v132
                            case "a":
                                var v133 uint16 = v20
                                _ = v133
                            }
                        }
                        _ = fi121
                        _ = v122
                    }
                    if v120 {
                        var v134 uint16 = (v20 ^ (v20 & v20))
                        fmt.Println(v20)
                        _ = v134
                    }
                    _ = v120
                } else {
                    v135 := GenS8{F0: (true || (true && false)), F1: 9223372036854775807}
                    if true {
                        v136 := GenS9{F0: 1, F1: (!true)}
                        v137 := GenS6{F0: 18446744073709551615}
                        switch 9223372036854775807 {
                        case 127:
                            v2 = (v6 << ((uint(v20)) & 63))
                        case 2147483647:
                            fmt.Println(v2, v20, v20)
                        }
                        _ = v136
                        _ = v137
                    }
                    _ = v135
                }
                switch int(v20) {
                case 42:
                    v6 -= (v6 + v6)
                default:
                    switch 1 {
                    case -34:
                        v6 = v6
                    case 32767:
                        for fi138 := 0; fi138 < 14; fi138++ {
                            v2 = 68
                            fmt.Println(v20, v6)
                            var v139 int32 = v5
                            _ = fi138
                            _ = v139
                        }
                    }
                }
                fmt.Println(v5, v2, v6)
            } else {
                switch "a" {
                case "a":
                    v5 = (v5 | 0)
                default:
                    switch ("a" + "\n\t") {
                    case "hello":
                        v9 += (^(v9))
                    case "":
                        v140 := []int{-255, -2147483648}
                        _ = v140
                    case "a":
                        switch "foo bar" {
                        case "hello":
                            v3[v9] = (4294967295 <= 1)
                        }
                    }
                }
            }
            v6 += (v2 >> ((genFn1(92, -2147483648)) & 63))
        }
    } else {
        var v141 int8 = -127
        _ = v141
    }
    v142 := GenS10{F0: 42}
    var v143 int = v9
    var v144 uint32 = v2
    fmt.Println(v20, v144, v6)
    fmt.Println(v143, v5)
    v145 := GenS11{F0: -299, F1: 820}
    for fi146 := 0; fi146 < 2; fi146++ {
        v147 := GenS8{F0: true, F1: -1}
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v6, v144)
        case "\n\t":
            if (!(42 < 1000)) {
                switch ("hello" + "\n\t") {
                case "hello":
                    var v148 uint32 = v2
                    _ = v148
                }
                v149 := []int{882, fi146, 127, -9223372036854775808}
                var v150 string = "hello"
                _ = v149
                _ = v150
            }
        case "":
            var v151 float32 = 0.0
            _ = v151
        default:
            _ = len(v3)
        }
        for fi152 := 0; fi152 < 12; fi152++ {
            delete(v1, (v6 - 836))
            _ = fi152
        }
        _ = fi146
        _ = v147
    }
    var v153 uint = 18446744073709551615
    fmt.Println(v153, v5, v5)
    v144 = (v2 % ((65535) | 1))
    v154 := GenS5{F0: 884, F1: 0}
    var v155 int16 = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v143:", v143)
    fmt.Println("v144:", v144)
    fmt.Println("v153:", v153)
    fmt.Println("v155:", v155)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v142
    _ = v143
    _ = v144
    _ = v145
    _ = v153
    _ = v154
    _ = v155
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
