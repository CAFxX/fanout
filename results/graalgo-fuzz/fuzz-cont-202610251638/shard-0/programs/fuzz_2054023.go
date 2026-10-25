package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 uint32; F2 string }
type GenS4 struct { F0 uint8; F1 uint16; F2 uint }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 uint64; F1 uint; F2 uint32 }
type GenS7 struct { F0 int8; F1 uint }
type GenS8 struct { F0 bool; F1 uint32 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint16; F1 int16 }
type GenS11 struct { F0 int16; F1 uint64; F2 int16 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 int8; F1 int16; F2 uint64 }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 uint16; F1 bool; F2 int64 }
type GenS17 struct { F0 int; F1 bool }
type GenS18 struct { F0 int32 }
type GenS19 struct { F0 uint }
type GenS20 struct { F0 uint8; F1 bool }
type GenS21 struct { F0 uint16; F1 uint }
func genFn0(p0 int64, p1 uint8, p2 uint) int32 {
    p1 -= p1
    v4 := GenS1{F0: 0, F1: -120}
    var v5 int16 = -32768
    var v6 int8 = -103
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 1
    var v3 int = 1
    if (1 >= -1) {
        switch v3 {
        case -2147483648:
            switch "\n\t" {
            case "hello":
                v2 = v2
            case "foo bar":
                for fi8 := 0; fi8 < 10; fi8++ {
                    var v9 string = ""
                    v3 = 1
                    for fi10 := 0; fi10 < 20; fi10++ {
                        delete(v1, 42)
                        _ = fi10
                    }
                    _ = fi8
                    _ = v9
                }
            case "":
                var v11 int32 = 415
                _ = v11
            default:
                if (!(4294967295 < 42)) {
                    switch "foo bar" {
                    case "foo bar":
                        var v12 uint32 = 581
                        _ = v12
                    case "\n\t":
                        for fi13 := 0; fi13 < 6; fi13++ {
                            v14 := GenS1{F0: 1, F1: 1}
                            delete(v1, 805)
                            _ = fi13
                            _ = v14
                        }
                    default:
                        v15 := GenS1{F0: 2147483647, F1: 0}
                        _ = v15
                    }
                } else {
                    var v16 float64 = (((1.5 + (0.1 / (0.0 + float64(1.5)))) * 1.5) / 1e10)
                    _ = v16
                }
            }
        default:
            for fi17 := 0; fi17 < 15; fi17++ {
                var v18 float64 = ((0.1 * -2.25) / 1.0)
                for fi19 := 0; fi19 < 13; fi19++ {
                    var v20 int8 = 1
                    v21 := []int{}
                    switch (fi19 - (fi17 % (((v3 + fi19)) | 1))) {
                    case 412:
                        switch ("" + "\n\t") {
                        case "":
                            if ((v20 * 100) != v20) {
                                switch fi19 {
                                case -974:
                                    var v22 uint64 = 988
                                    _ = v22
                                case 32768:
                                    if len(v21) >= 2 {
                                        v21 = v21[1:len(v21)]
                                    }
                                    v23 := GenS1{F0: 273, F1: -128}
                                    _ = v23
                                default:
                                    var v24 uint16 = v2
                                    _ = v24
                                }
                                v25 := GenS2{F0: (v20 & (v20 - (v20 >> ((1) & 63))))}
                                var v26 uint16 = (^(v2))
                                _ = v25
                                _ = v26
                            }
                        case "a":
                            v18 = (((-2.25 - v18) * (3.14159 - float64(-2.25))) + 3.14159)
                        case "\n\t":
                            if (("\n\t" + ("\n\t" + ("a" + "hello"))) >= ((("a" + "a") + ((("foo bar" + (("a" + "a") + "")) + "") + "\n\t")) + ("" + "foo bar"))) {
                                var v27 bool = (355 > 32767)
                                switch (("foo bar" + "foo bar") + "\n\t") {
                                case "hello":
                                    v18 = v18
                                case "foo bar":
                                    v28 := GenS1{F0: 2, F1: (v20 / ((127) | 1))}
                                    _ = v28
                                case "":
                                    v20 = v20
                                }
                                v3 -= fi19
                                var v29 int8 = (v20 / ((100) | 1))
                                if len(v21) > 0 {
                                    v21[0] = (-(fi19))
                                }
                                _ = v27
                                _ = v29
                            }
                        }
                    case -1:
                        v20 = (v20 << ((0) & 63))
                    default:
                        fmt.Println(v3)
                    }
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                _ = fi17
                _ = v18
            }
        }
    } else {
        v3 -= (-(v3))
        switch (v3 << ((4294967295) & 63)) {
        case 79:
            fmt.Println(v3, v3)
        case 630:
            fmt.Println(v2, v3)
        case -645:
            v30 := GenS3{F0: ("a" + "hello"), F1: 2, F2: (("hello" + "\n\t") + "foo bar")}
            _ = v30
        default:
            for fi31 := 0; fi31 < 7; fi31++ {
                for fi32 := 0; fi32 < 9; fi32++ {
                    v2 = v2
                    if ("\n\t" <= "a") {
                        if ((305 <= (v2 + v2)) || (-1 < 0)) {
                            v33 := GenS4{F0: 215, F1: v2, F2: 1}
                            var v34 bool = true
                            _ = v33
                            _ = v34
                        }
                        v3 = (v3 * v3)
                    }
                    v35 := GenS5{F0: -1, F1: 42}
                    _ = fi32
                    _ = v35
                }
                for fi36 := 0; fi36 < 17; fi36++ {
                    if false {
                        if (65535 > 1) {
                            if false {
                                v3 += 972
                            } else {
                                v37 := []int{v3, (v3 & -9223372036854775808), v3}
                                var v38 int = fi36
                                _ = v37
                                _ = v38
                            }
                        } else {
                            var v39 int16 = 1
                            if false {
                                var v40 uint32 = 2
                                if false {
                                    if (!false) {
                                        v39 += 32767
                                        var v41 int32 = -1
                                        _ = v41
                                    }
                                    v42 := v1[v40]
                                    _ = v42
                                }
                                v43 := GenS6{F0: 42, F1: 846, F2: (^(v40))}
                                _ = v40
                                _ = v43
                            } else {
                                v44 := v1[2]
                                v3 += 4294967295
                                v45 := GenS5{F0: int64(v3), F1: 0}
                                fmt.Println(fi36, v3)
                                _ = v44
                                _ = v45
                            }
                            _ = v39
                        }
                        var v46 int = -2147483648
                        var v47 int64 = 42
                        v47 = v47
                        v48 := GenS3{F0: "foo bar", F1: 4294967295, F2: "foo bar"}
                        if (!false) {
                            switch "" {
                            case "":
                                v49 := []int{fi36, v3, v46}
                                _ = v49
                            default:
                                var v50 uint32 = 858
                                _ = v50
                            }
                        } else {
                            v51 := GenS4{F0: 1, F1: 21, F2: 243}
                            _ = v51
                        }
                        v52 := GenS1{F0: 801, F1: 127}
                        _ = v46
                        _ = v47
                        _ = v48
                        _ = v52
                    }
                    var v53 int8 = 127
                    v54 := GenS6{F0: 0, F1: 2, F2: 42}
                    _ = fi36
                    _ = v53
                    _ = v54
                }
                _ = fi31
            }
        }
    }
    if (18446744073709551615 >= 268) {
        if ((!(true || (v3 == v3))) && (uint64(v2) <= 461)) {
            switch v3 {
            case 3:
                v55 := GenS7{F0: -58, F1: 18446744073709551615}
                _ = v55
            case -573:
                v3 -= (v3 + (v3 ^ v3))
            case -2147483648:
                var v56 float32 = (3.14159 - 0.1)
                _ = v56
            default:
                switch "foo bar" {
                case "foo bar":
                    if true {
                        for fi57 := 0; fi57 < 13; fi57++ {
                            var v58 uint8 = 42
                            _ = fi57
                            _ = v58
                        }
                        v2 += v2
                    } else {
                        var v59 string = "hello"
                        switch ("\n\t" + v59) {
                        case "\n\t":
                            v59 = ((v59 + v59) + (((("" + ("foo bar" + v59)) + v59) + v59) + "\n\t"))
                        case "foo bar":
                            v59 = ("a" + v59)
                        case "":
                            v2 += (v2 & (^(v2)))
                        default:
                            var v60 bool = (false || false)
                            _ = v60
                        }
                        _ = v59
                    }
                case "hello":
                    switch "a" {
                    case "":
                        var v61 uint32 = 2147483647
                        _ = v61
                    default:
                        var v62 float32 = 0.1
                        _ = v62
                    }
                case "a":
                    var v63 float32 = 1e10
                    _ = v63
                }
            }
            v64 := GenS2{F0: 126}
            for fi65 := 0; fi65 < 20; fi65++ {
                fmt.Println(v3, v3)
                fmt.Println(fi65, v3)
                _ = fi65
            }
            _ = v64
        } else {
            var v66 int8 = 127
            var v67 uint = 919
            if (398 <= 0) {
                var v68 uint64 = 995
                fmt.Println(v67, v3, v66)
                var v69 uint8 = 127
                _ = v68
                _ = v69
            }
            _ = v66
            _ = v67
        }
        if false {
            var v70 int64 = -121
            v71 := []int{}
            if false {
                v72 := GenS5{F0: -9223372036854775808, F1: uint64(v2)}
                var v73 int32 = genFn0(1, 0, 2)
                _ = v72
                _ = v73
            } else {
                var v74 string = ("" + "hello")
                v2 -= (^(v2))
                if false {
                    fmt.Println(v2)
                    var v75 uint64 = 18446744073709551615
                    for fi76 := 0; fi76 < 10; fi76++ {
                        v77 := GenS3{F0: "hello", F1: 2147483648, F2: "hello"}
                        if false {
                            if true {
                                switch v74 {
                                case "a":
                                    var v78 int8 = -1
                                    _ = v78
                                case "":
                                    for fi79 := 0; fi79 < 1; fi79++ {
                                        v80 := GenS6{F0: (v75 ^ 0), F1: 4294967295, F2: 227}
                                        _ = fi79
                                        _ = v80
                                    }
                                case "foo bar":
                                    var v81 uint32 = 2147483647
                                    _ = v81
                                default:
                                    v82 := GenS5{F0: (v70 >> ((275) & 63)), F1: uint64(v2)}
                                    _ = v82
                                }
                                if false {
                                    switch "\n\t" {
                                    case "\n\t":
                                        for fi83 := 0; fi83 < 13; fi83++ {
                                            var v84 float64 = 1.5
                                            fmt.Println(fi76, v84)
                                            _ = fi83
                                            _ = v84
                                        }
                                    }
                                    for fi85 := 0; fi85 < 9; fi85++ {
                                        if false {
                                            var v86 float64 = (-2.25 - 0.1)
                                            _ = v86
                                        } else {
                                            if true {
                                                _ = len(v1)
                                            } else {
                                                if ((2 != 147) || (false && (32767 == -32768))) {
                                                    v75 += v75
                                                    v87 := GenS8{F0: true, F1: 65535}
                                                    _ = v87
                                                } else {
                                                    var v88 int16 = 1000
                                                    if (!false) {
                                                        v3 = fi85
                                                    }
                                                    var v89 float32 = 3.14159
                                                    _ = v88
                                                    _ = v89
                                                }
                                            }
                                            if (-878 == -1) {
                                                var v90 int8 = 100
                                                _ = v90
                                            }
                                            switch 32767 {
                                            case 2147483647:
                                                v91 := GenS9{F0: 2}
                                                _ = v91
                                            }
                                            if false {
                                                v92 := GenS10{F0: (^(v2)), F1: 1000}
                                                v70 += int64(fi85)
                                                switch ("hello" + (v74 + v74)) {
                                                case "foo bar":
                                                    switch ("a" + ("a" + v74)) {
                                                    case "a":
                                                        var v93 int8 = 100
                                                        _ = v93
                                                    case "":
                                                        v75 = (v75 << ((42) & 63))
                                                    default:
                                                        var v94 float64 = 1.5
                                                        _ = v94
                                                    }
                                                default:
                                                    switch v74 {
                                                    case "":
                                                        fmt.Println(fi85, fi76)
                                                    case "hello":
                                                        var v95 uint32 = 65535
                                                        _ = v95
                                                    default:
                                                        if (false && false) {
                                                            v1[0] = (fi85 & (fi76 / ((fi76) | 1)))
                                                        } else {
                                                            fmt.Println(fi85, v74)
                                                            v2 = v2
                                                            if false {
                                                                v74 = v74
                                                            }
                                                        }
                                                    }
                                                }
                                                _ = v92
                                            }
                                        }
                                        delete(v1, 461)
                                        v70 = (v70 ^ v70)
                                        v70 += (v70 & v70)
                                        var v96 uint8 = 42
                                        v97 := GenS11{F0: -32768, F1: (v75 & (v75 % (((v75 << ((908) & 63))) | 1))), F2: 1}
                                        _ = fi85
                                        _ = v96
                                        _ = v97
                                    }
                                } else {
                                    for fi98 := 0; fi98 < 15; fi98++ {
                                        var v99 string = v74
                                        v75 = (v75 % ((v75) | 1))
                                        _ = fi98
                                        _ = v99
                                    }
                                    v74 = (((("" + (v74 + ("foo bar" + (v74 + v74)))) + "hello") + "foo bar") + "foo bar")
                                    if len(v71) >= 2 {
                                        v71 = v71[1:len(v71)]
                                    }
                                }
                                delete(v1, 0)
                            }
                        } else {
                            v3 = (fi76 / (((-(fi76))) | 1))
                            v75 += (v75 & v75)
                        }
                        v74 = "hello"
                        _ = fi76
                        _ = v77
                    }
                    _ = v75
                }
                _ = v74
            }
            _ = v70
            _ = v71
        } else {
            var v100 float32 = 3.14159
            switch (v3 << ((502) & 63)) {
            case 128:
                if (false || (!((("a" + (("" + "a") + ("a" + ""))) + "a") < ""))) {
                    var v101 uint = 979
                    _ = v101
                } else {
                    v102 := GenS12{F0: v2}
                    fmt.Println(v100)
                    v3 -= (v3 + (v3 / (((^(v3))) | 1)))
                    _ = v102
                }
            case 1:
                v103 := []int{9223372036854775807, v3, v3, v3}
                _ = v103
            case 32767:
                var v104 bool = true
                _ = v104
            }
            _ = v100
        }
        switch v3 {
        case 9223372036854775807:
            v105 := v1[399]
            _ = v105
        default:
            if (("hello" + "a") <= "\n\t") {
                fmt.Println(v2, v2)
                v1[0] = (v3 | 127)
                v106 := GenS7{F0: 109, F1: 18446744073709551615}
                _ = v106
            }
        }
    }
    fmt.Println(v2)
    var v107 string = "a"
    if (!true) {
        var v108 int = v3
        _ = v108
    } else {
        if true {
            if (true || true) {
                for fi109 := 0; fi109 < 6; fi109++ {
                    for fi110 := 0; fi110 < 20; fi110++ {
                        v111 := GenS3{F0: (v107 + "hello"), F1: 2, F2: v107}
                        _ = fi110
                        _ = v111
                    }
                    v112 := []int{(fi109 - v3), fi109, 32767, (v3 ^ fi109)}
                    _ = fi109
                    _ = v112
                }
            } else {
                v3 = int(v2)
            }
        }
    }
    var v113 string = ("foo bar" + v107)
    if (-2.25 > 1e10) {
        v114 := GenS11{F0: 675, F1: 355, F2: 0}
        _ = v114
    }
    fmt.Println(v113, v3)
    switch ("" + "\n\t") {
    case "a":
        for fi115 := 0; fi115 < 14; fi115++ {
            if true {
                var v116 bool = false
                if true {
                    var v117 uint16 = (^(v2))
                    var v118 int = fi115
                    v119 := GenS13{F0: 2}
                    _ = v117
                    _ = v118
                    _ = v119
                } else {
                    v120 := GenS1{F0: 272, F1: 100}
                    var v121 uint64 = 18446744073709551615
                    _ = v120
                    _ = v121
                }
                _ = v116
            }
            v122 := GenS14{F0: -128, F1: 455, F2: 397}
            v123 := []int{9223372036854775807, 42, (fi115 | (v3 / ((1) | 1)))}
            _ = fi115
            _ = v122
            _ = v123
        }
    case "foo bar":
        v124 := []int{v3, (v3 & (v3 << ((65535) & 63)))}
        _ = v124
    default:
        switch v113 {
        case "":
            var v125 int64 = 355
            _ = v125
        }
    }
    v126 := GenS15{F0: -32768}
    switch "foo bar" {
    case "\n\t":
        for fi127 := 0; fi127 < 4; fi127++ {
            var v128 int = -400
            if (1 <= 18446744073709551615) {
                v3 += -2147483648
                v128 -= (-(v3))
            } else {
                if true {
                    if (2 < uint64(v2)) {
                        delete(v1, 65535)
                        switch fi127 {
                        case 32767:
                            var v129 uint = 42
                            _ = v129
                        default:
                            v1[4294967295] = (v3 << ((0) & 63))
                        }
                    }
                    v3 = v128
                    var v130 int64 = -9223372036854775808
                    _ = v130
                }
                v131 := v1[2147483647]
                _ = v131
            }
            switch ("hello" + v113) {
            case "":
                v113 = v113
            case "foo bar":
                fmt.Println(v3, v2, v2)
            }
            _ = fi127
            _ = v128
        }
    case "foo bar":
        v132 := GenS5{F0: 42, F1: 4294967295}
        _ = v132
    case "a":
        for fi133 := 0; fi133 < 20; fi133++ {
            fmt.Println(v107, v3, v2)
            for fi134 := 0; fi134 < 8; fi134++ {
                v135 := GenS4{F0: 42, F1: 2, F2: 4294967295}
                v136 := GenS16{F0: (v2 / (((v2 >> ((65535) & 63))) | 1)), F1: true, F2: -1}
                switch 984 {
                case 127:
                    var v137 int = (-(v3))
                    _ = v137
                }
                switch fi133 {
                case 32768:
                    switch (-(fi134)) {
                    case 316:
                        v138 := v1[1]
                        _ = v138
                    case 127:
                        v139 := GenS17{F0: fi134, F1: (!false)}
                        _ = v139
                    case 2147483647:
                        if false {
                            v113 = v107
                            switch -937 {
                            case 12:
                                var v140 bool = false
                                _ = v140
                            case -1:
                                var v141 int8 = 0
                                v113 = ("" + (v107 + v107))
                                _ = v141
                            case 1:
                                fmt.Println(fi134)
                            default:
                                v2 = (v2 ^ v2)
                            }
                            if true {
                                if (false || false) {
                                    var v142 float64 = (0.0 / 3.14159)
                                    if false {
                                        v143 := GenS2{F0: 127}
                                        v144 := []int{-793, 32768, 51}
                                        _ = v143
                                        _ = v144
                                    } else {
                                        v145 := []int{(v3 % (((^(v3))) | 1)), 32768}
                                        _ = v145
                                    }
                                    var v146 int32 = genFn0(0, 2, 0)
                                    var v147 int = v3
                                    _ = v142
                                    _ = v146
                                    _ = v147
                                }
                                v148 := GenS5{F0: 0, F1: 962}
                                _ = v148
                            }
                            if true {
                                v113 = "hello"
                                v149 := GenS13{F0: 255}
                                var v150 uint = 639
                                _ = v149
                                _ = v150
                            }
                            var v151 uint8 = 211
                            v113 = v113
                            _ = v151
                        }
                    }
                }
                var v152 float32 = 1e10
                _ = fi134
                _ = v135
                _ = v136
                _ = v152
            }
            switch v107 {
            case "foo bar":
                v3 = 231
            case "\n\t":
                var v153 uint32 = 334
                _ = v153
            case "a":
                for fi154 := 0; fi154 < 8; fi154++ {
                    if true {
                        if (true && false) {
                            if ((v3 ^ -9223372036854775808) != v3) {
                                var v155 uint64 = 1
                                if true {
                                    v2 = (v2 * v2)
                                    v156 := GenS17{F0: 4294967295, F1: (42 > 252)}
                                    v157 := GenS13{F0: 65535}
                                    _ = v156
                                    _ = v157
                                } else {
                                    var v158 string = (v113 + v107)
                                    switch v107 {
                                    case "a":
                                        if (-128 == -1) {
                                            v155 = (v155 >> ((2) & 63))
                                            switch (fi133 & (fi133 + (fi154 ^ -666))) {
                                            case 0:
                                                v159 := GenS4{F0: 67, F1: (^(v2)), F2: 0}
                                                _ = v159
                                            case -470:
                                                switch v107 {
                                                case "a":
                                                    if (true || (0.1 == 1.5)) {
                                                        if true {
                                                            switch (v3 & fi133) {
                                                            case 359:
                                                                if false {
                                                                    v3 = (fi133 & fi133)
                                                                }
                                                            case 128:
                                                                v155 -= (v155 * v155)
                                                            case 2147483647:
                                                                v160 := GenS18{F0: -1}
                                                                _ = v160
                                                            }
                                                            v3 += (^(v3))
                                                            var v161 int32 = int32(v155)
                                                            v162 := GenS13{F0: 823}
                                                            v107 = v113
                                                            var v163 uint8 = 1
                                                            var v164 float64 = 3.14159
                                                            v165 := GenS17{F0: 128, F1: true}
                                                            v166 := GenS5{F0: 42, F1: (v155 - 75)}
                                                            var v167 int32 = genFn0(-9223372036854775808, (v163 & (v163 << ((552) & 63))), 46)
                                                            _ = v161
                                                            _ = v162
                                                            _ = v163
                                                            _ = v164
                                                            _ = v165
                                                            _ = v166
                                                            _ = v167
                                                        }
                                                        _ = len(v1)
                                                    }
                                                case "":
                                                    v158 = v158
                                                default:
                                                    fmt.Println(v107)
                                                }
                                                v168 := []int{}
                                                _ = v168
                                            case -1:
                                                v169 := GenS10{F0: v2, F1: 1000}
                                                _ = v169
                                            default:
                                                v107 = ("hello" + "\n\t")
                                            }
                                            v170 := GenS1{F0: 65535, F1: 18}
                                            _ = v170
                                        } else {
                                            v158 = v158
                                            fmt.Println(v107, fi133)
                                            v113 = ("foo bar" + v107)
                                        }
                                    default:
                                        var v171 bool = false
                                        _ = v171
                                    }
                                    _ = v158
                                }
                                if (2147483648 >= 1) {
                                    v172 := []int{v3, (v3 & -2147483648)}
                                    _ = v172
                                } else {
                                    v107 = "\n\t"
                                    if ((true && (false || (!(!(fi133 != fi133))))) && ("a" <= v107)) {
                                        if false {
                                            v173 := []int{(fi154 << ((18446744073709551615) & 63)), fi154, (v3 & v3)}
                                            var v174 uint16 = (v2 & 32767)
                                            switch (v113 + v107) {
                                            case "\n\t":
                                                if true {
                                                    delete(v1, 4294967295)
                                                    v155 += (v155 >> ((255) & 63))
                                                    if (false || true) {
                                                        v175 := GenS19{F0: 4294967295}
                                                        _ = v175
                                                    } else {
                                                        var v176 uint8 = 2
                                                        var v177 int8 = 0
                                                        var v178 float64 = 1e10
                                                        _ = v176
                                                        _ = v177
                                                        _ = v178
                                                    }
                                                }
                                                v179 := v1[0]
                                                _ = v179
                                            case "a":
                                                v174 = (^(v174))
                                            case "hello":
                                                if (!(!true)) {
                                                    var v180 int32 = genFn0(0, 128, 42)
                                                    if len(v173) > 0 {
                                                        v173[0] = v3
                                                    }
                                                    var v181 uint8 = uint8(v2)
                                                    _ = v180
                                                    _ = v181
                                                } else {
                                                    var v182 float32 = (1.5 - 0.1)
                                                    _ = v182
                                                }
                                            default:
                                                fmt.Println(fi133, v155)
                                            }
                                            _ = v173
                                            _ = v174
                                        } else {
                                            switch -1 {
                                            case -9223372036854775808:
                                                v2 = 255
                                            case 32767:
                                                var v183 uint16 = v2
                                                _ = v183
                                            case -60:
                                                if false {
                                                    v2 += v2
                                                    v3 = 4294967295
                                                }
                                            default:
                                                var v184 uint8 = uint8(v155)
                                                _ = v184
                                            }
                                        }
                                        if (200 <= -398) {
                                            v1[0] = (v3 & 1)
                                            switch fi133 {
                                            case 32768:
                                                fmt.Println(fi154, v107, v113)
                                            case 9223372036854775807:
                                                v185 := []int{v3, (v3 >> ((1) & 63)), (-(fi133))}
                                                _ = v185
                                            case -2147483648:
                                                switch fi133 {
                                                case 2147483647:
                                                    v155 = 1
                                                case 128:
                                                    delete(v1, 2147483647)
                                                case -22:
                                                    fmt.Println(fi154)
                                                    v2 = 255
                                                default:
                                                    v113 = v107
                                                }
                                            }
                                            if false {
                                                if false {
                                                    v3 = (v3 << ((18446744073709551615) & 63))
                                                    var v186 string = "a"
                                                    _ = v186
                                                } else {
                                                    v113 = "a"
                                                }
                                            } else {
                                                v3 -= v3
                                                var v187 int32 = 42
                                                _ = v187
                                            }
                                        } else {
                                            v2 = uint16(v155)
                                            delete(v1, 450)
                                            v2 = v2
                                        }
                                    }
                                    v188 := GenS20{F0: 177, F1: false}
                                    _ = v188
                                }
                                if (false && true) {
                                    var v189 uint64 = v155
                                    var v190 int8 = -101
                                    _ = v189
                                    _ = v190
                                }
                                _ = v155
                            } else {
                                var v191 int64 = 9223372036854775807
                                if false {
                                    v113 = v107
                                    var v192 uint32 = 2
                                    var v193 int = fi154
                                    _ = v192
                                    _ = v193
                                } else {
                                    fmt.Println(v113, fi154, v191)
                                    switch v113 {
                                    case "hello":
                                        v194 := []int{}
                                        _ = v194
                                    case "a":
                                        v195 := GenS13{F0: 294}
                                        _ = v195
                                    default:
                                        switch (^(fi133)) {
                                        case 128:
                                            var v196 uint8 = 255
                                            _ = v196
                                        }
                                    }
                                    if (!true) {
                                        if (127 == 42) {
                                            var v197 float32 = 0.0
                                            v198 := v1[0]
                                            _ = v197
                                            _ = v198
                                        } else {
                                            v113 = (v113 + ((v107 + ("" + "a")) + ""))
                                            v107 = ("foo bar" + v113)
                                        }
                                    }
                                }
                                _ = v191
                            }
                        } else {
                            v199 := v1[877]
                            v113 = v113
                            _ = v199
                        }
                        v2 -= v2
                        switch (^(v3)) {
                        case 545:
                            v113 = (v107 + v107)
                            v113 = "\n\t"
                        case 1:
                            fmt.Println(v3, fi133, v2)
                        case 128:
                            var v200 uint = 588
                            _ = v200
                        default:
                            var v201 uint = 2
                            _ = v201
                        }
                        var v202 float32 = 0.1
                        v203 := GenS19{F0: 255}
                        if false {
                            var v204 int16 = 1
                            switch v3 {
                            case -2147483648:
                                fmt.Println(v3, v3)
                            default:
                                var v205 uint = 601
                                _ = v205
                            }
                            var v206 uint = 18446744073709551615
                            _ = v204
                            _ = v206
                        }
                        fmt.Println(v3)
                        _ = v202
                        _ = v203
                    } else {
                        var v207 float64 = 0.1
                        _ = v207
                    }
                    _ = fi154
                }
            default:
                var v208 int16 = 54
                _ = v208
            }
            _ = fi133
        }
    default:
        if ((-9223372036854775808 <= -1) || true) {
            v209 := GenS6{F0: 621, F1: 67, F2: 0}
            var v210 int16 = 1
            _ = v209
            _ = v210
        } else {
            switch v113 {
            case "hello":
                v3 -= (-(v3))
            case "foo bar":
                switch 677 {
                case 1:
                    for fi211 := 0; fi211 < 6; fi211++ {
                        v2 += v2
                        v113 = ""
                        _ = fi211
                    }
                case 0:
                    v107 = v113
                default:
                    v113 = v113
                }
            case "a":
                v212 := []int{127}
                _ = v212
            }
            if (true && false) {
                if (2147483647 >= int32(v3)) {
                    var v213 uint = 255
                    _ = v213
                }
                var v214 int8 = -128
                v3 += (v3 % (((v3 % ((2147483647) | 1))) | 1))
                _ = v214
            } else {
                var v215 int32 = 42
                _ = v215
            }
        }
    }
    v3 = 9223372036854775807
    v216 := GenS21{F0: (^(v2)), F1: 213}
    var v217 uint = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v107:", v107)
    fmt.Println("v113:", v113)
    fmt.Println("v2:", v2)
    fmt.Println("v217:", v217)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v107
    _ = v113
    _ = v126
    _ = v2
    _ = v216
    _ = v217
    _ = v3
}
