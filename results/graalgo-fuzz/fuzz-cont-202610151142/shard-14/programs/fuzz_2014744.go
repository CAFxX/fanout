package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS4 struct { F0 uint; F1 int16; F2 uint64 }
type GenS5 struct { F0 bool; F1 uint; F2 uint64 }
type GenS6 struct { F0 uint32; F1 uint }
type GenS7 struct { F0 uint64; F1 uint16 }
type GenS8 struct { F0 int32; F1 string }
type GenS9 struct { F0 int32; F1 uint8 }
type GenS10 struct { F0 string; F1 uint8 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 int8; F1 uint64 }
type GenS13 struct { F0 uint64; F1 int64; F2 int64 }
type GenS14 struct { F0 int32 }
func genFn0(p0 int, p1 uint16, p2 uint16) bool {
    var v7 uint64 = 4294967295
    var v8 uint = 4294967295
    for fi9 := 0; fi9 < 6; fi9++ {
        if (-32768 <= -539) {
            v10 := []int{(p0 + 361), 1, fi9, -917}
            _ = v10
        } else {
            v11 := GenS1{F0: 2147483647}
            v12 := []int{-242, 283}
            _ = v11
            _ = v12
        }
        fmt.Println(p1, p1, fi9)
        _ = fi9
    }
    if (false && false) {
        for fi13 := 0; fi13 < 9; fi13++ {
            var v14 float64 = float64(0.1)
            var v15 int8 = 100
            _ = fi13
            _ = v14
            _ = v15
        }
    } else {
        if true {
            if true {
                switch "a" {
                case "hello":
                    var v16 int = 2147483647
                    _ = v16
                default:
                    var v17 uint32 = 2147483648
                    _ = v17
                }
            }
            v18 := GenS1{F0: -1}
            v7 -= v7
            _ = v18
        } else {
            for fi19 := 0; fi19 < 13; fi19++ {
                fmt.Println(p2, v8, fi19)
                for fi20 := 0; fi20 < 14; fi20++ {
                    v21 := []int{p0, (-(fi19))}
                    var v22 uint8 = 2
                    var v23 float32 = 0.0
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
                _ = fi19
            }
            var v24 int64 = 1
            _ = v24
        }
    }
    var v25 float32 = (3.14159 / float32(((float64((-2.25 - float32(3.14159))) * 1e10) / ((-2.25 / -2.25) - 1.5))))
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    _ = v7
    _ = v8
    return false
}
func genFn1(p0 uint16, p1 uint32, p2 uint16) uint64 {
    v26 := []int{}
    p1 = 65535
    for fi27 := 0; fi27 < 3; fi27++ {
        fmt.Println(p1, p0)
        if (false || (32767 < int16(p1))) {
            switch fi27 {
            case 744:
                var v28 bool = (96 != 42)
                _ = v28
            case -2147483648:
                v29 := GenS1{F0: 42}
                _ = v29
            }
            if ("foo bar" != "hello") {
                v30 := GenS1{F0: 1}
                _ = v30
            }
        } else {
            p1 -= p1
            fmt.Println(fi27, p2)
            for fi31 := 0; fi31 < 2; fi31++ {
                var v32 uint64 = 815
                if (127 >= 0) {
                    var v33 int = 4294967295
                    if true {
                        var v34 float64 = float64(0.0)
                        var v35 string = "hello"
                        p1 = (p1 * 65535)
                        _ = v34
                        _ = v35
                    }
                    _ = v33
                }
                _ = fi31
                _ = v32
            }
        }
        p0 += p2
        _ = fi27
    }
    switch 42 {
    case 42:
        switch "foo bar" {
        case "":
            var v36 uint16 = 529
            _ = v36
        case "a":
            var v37 int64 = -1
            _ = v37
        case "\n\t":
            if (4294967295 <= uint64(p2)) {
                p1 += p1
                if len(v26) > 0 {
                    v26[0] = -462
                }
                v38 := GenS1{F0: 1}
                _ = v38
            } else {
                for fi39 := 0; fi39 < 13; fi39++ {
                    switch fi39 {
                    case -341:
                        switch (fi39 * -827) {
                        case -9223372036854775808:
                            p0 -= 255
                        default:
                            if false {
                                v40 := GenS1{F0: 1}
                                _ = v40
                            }
                        }
                    case 759:
                        p1 += p1
                    case 802:
                        v41 := GenS2{F0: 0, F1: 18446744073709551615}
                        _ = v41
                    }
                    switch "hello" {
                    case "hello":
                        p0 = p0
                    case "foo bar":
                        if (!false) {
                            _ = len(v26)
                            v42 := GenS3{F0: 1, F1: 579, F2: p1}
                            p1 = uint32(p0)
                            _ = v42
                        }
                    default:
                        v43 := GenS3{F0: 11, F1: (p2 + (^(p2))), F2: p1}
                        _ = v43
                    }
                    _ = fi39
                }
                v44 := GenS2{F0: -1, F1: 4294967295}
                _ = v44
            }
        default:
            if len(v26) > 0 {
                v26[0] = 676
            }
        }
    default:
        for fi45 := 0; fi45 < 13; fi45++ {
            var v46 int16 = 1
            _ = fi45
            _ = v46
        }
    }
    var v47 int64 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v47
    return 0
}

func main() {
    v1 := make(map[int]int)
    var v2 string = "foo bar"
    v3 := make(map[string]bool)
    v4 := make(map[int]bool)
    v5 := make(map[uint32]bool)
    v6 := make(map[int]string)
    v2 = "foo bar"
    switch v2 {
    case "":
        v2 = (v2 + v2)
    case "hello":
        for fi48 := 0; fi48 < 17; fi48++ {
            switch v2 {
            case "foo bar":
                var v49 float64 = ((0.1 - 3.14159) / -2.25)
                _ = v49
            case "a":
                fmt.Println(fi48)
            default:
                v2 = v2
            }
            _ = fi48
        }
    case "\n\t":
        if (!(255 >= 101)) {
            var v50 string = v2
            _ = v50
        }
    }
    for fi51 := 0; fi51 < 11; fi51++ {
        v52 := GenS3{F0: 127, F1: 945, F2: 42}
        v2 = v2
        v2 = ("" + (v2 + ("hello" + v2)))
        _ = fi51
        _ = v52
    }
    for fi53 := 0; fi53 < 6; fi53++ {
        if (false && true) {
            var v54 bool = true
            _ = v54
        }
        if (!true) {
            fmt.Println(v2, v2)
            v2 = ""
            v55 := GenS2{F0: -1, F1: 249}
            _ = v55
        } else {
            for fi56 := 0; fi56 < 11; fi56++ {
                var v57 string = v2
                v57 = ((("hello" + (v2 + (("hello" + "a") + ("a" + "foo bar")))) + v2) + ("" + v57))
                if false {
                    fmt.Println(fi53, fi53, fi56)
                } else {
                    if true {
                        v58 := GenS4{F0: 255, F1: 42, F2: 4294967295}
                        _ = v58
                    } else {
                        var v59 int8 = 97
                        _ = v59
                    }
                    var v60 int8 = 93
                    var v61 bool = false
                    _ = v60
                    _ = v61
                }
                _ = fi56
                _ = v57
            }
            v2 = "hello"
            v2 = "foo bar"
        }
        if (2 >= 1) {
            if (-1 != 0) {
                v2 = "a"
            }
            v5[65535] = false
            for fi62 := 0; fi62 < 3; fi62++ {
                var v63 int32 = 0
                v2 = "a"
                _ = fi62
                _ = v63
            }
        }
        _ = fi53
    }
    if false {
        var v64 int64 = -705
        _ = v64
    }
    v65 := GenS5{F0: (-9223372036854775808 < 9223372036854775807), F1: 4294967295, F2: genFn1(255, 860, 32767)}
    delete(v4, -870)
    var v66 uint64 = 18446744073709551615
    switch v2 {
    case "\n\t":
        switch (((("\n\t" + v2) + v2) + (((("a" + v2) + v2) + ("hello" + (v2 + v2))) + v2)) + v2) {
        case "foo bar":
            var v67 int = -2147483648
            _ = v67
        case "hello":
            switch v2 {
            case "foo bar":
                delete(v1, 4294967295)
            case "":
                var v68 bool = true
                _ = v68
            case "a":
                v66 = 0
            default:
                v66 = (v66 << ((4294967295) & 63))
            }
        default:
            v69 := GenS6{F0: 2147483648, F1: 327}
            _ = v69
        }
    case "foo bar":
        fmt.Println(v2, v66)
    default:
        var v70 int8 = 1
        _ = v70
    }
    v71 := []int{-1, 32767, -625}
    if false {
        for fi72 := 0; fi72 < 11; fi72++ {
            if (false && true) {
                v66 = v66
                v73 := GenS4{F0: 2, F1: -1, F2: 113}
                _ = v73
            }
            if (2147483647 < -420) {
                var v74 int8 = int8(v66)
                switch v2 {
                case "hello":
                    if false {
                        for fi75 := 0; fi75 < 14; fi75++ {
                            var v76 float32 = (3.14159 + (0.0 / (1e10 + 0.0)))
                            _ = fi75
                            _ = v76
                        }
                        fmt.Println(v66, v66)
                        switch 127 {
                        case 2147483647:
                            var v77 int = 4294967295
                            _ = v77
                        }
                    } else {
                        for fi78 := 0; fi78 < 1; fi78++ {
                            v79 := GenS2{F0: -829, F1: v66}
                            var v80 int32 = -1
                            switch v2 {
                            case "\n\t":
                                v81 := GenS7{F0: (v66 - 583), F1: 0}
                                _ = v81
                            case "":
                                var v82 bool = false
                                var v83 uint8 = 108
                                _ = v82
                                _ = v83
                            case "a":
                                var v84 float64 = -2.25
                                _ = v84
                            default:
                                var v85 uint = 788
                                v86 := GenS8{F0: (v80 - -2147483648), F1: v2}
                                if ((true && (!(((v74 - v74) < v74) || false))) || (1e10 >= (((1.5 - ((1e10 / 0.1) - 0.0)) + (3.14159 - (float64(((1.5 - (((-2.25 - 0.1) * (float32(0.1) + 1.5)) - 0.0)) - -2.25)) * (((1.5 - float64(-2.25)) - (0.1 / 3.14159)) / (0.0 - -2.25))))) + float64(0.0)))) {
                                    var v87 int16 = -1
                                    var v88 int32 = (v80 & 0)
                                    _ = v87
                                    _ = v88
                                }
                                _ = v85
                                _ = v86
                            }
                            _ = fi78
                            _ = v79
                            _ = v80
                        }
                    }
                case "a":
                    switch ((v2 + v2) + v2) {
                    case "":
                        if true {
                            var v89 int = (-(fi72))
                            _ = v89
                        }
                    case "\n\t":
                        if false {
                            for fi90 := 0; fi90 < 9; fi90++ {
                                var v91 int16 = 42
                                _ = fi90
                                _ = v91
                            }
                        } else {
                            var v92 int64 = 1
                            for fi93 := 0; fi93 < 14; fi93++ {
                                if (((3.14159 - float32((1e10 * 0.1))) * (3.14159 * 1e10)) < 1.5) {
                                    if (619 == (v92 * (v92 >> ((2) & 63)))) {
                                        v2 = "a"
                                        if ((593 >= 4294967295) || (!(!true))) {
                                            var v94 uint = 42
                                            if false {
                                                if len(v71) >= 2 {
                                                    v71 = v71[1:len(v71)]
                                                }
                                                v71 = append(v71, -695)
                                            } else {
                                                v95 := GenS9{F0: 2147483647, F1: 2}
                                                if (v74 == (v74 + v74)) {
                                                    _ = len(v71)
                                                    var v96 int = fi72
                                                    _ = v96
                                                }
                                                v66 -= uint64(v94)
                                                _ = v95
                                            }
                                            _ = v94
                                        }
                                        switch fi93 {
                                        case -2147483648:
                                            var v97 uint32 = 2147483647
                                            _ = v97
                                        case 9223372036854775807:
                                            v74 = v74
                                        case 128:
                                            v74 = v74
                                        default:
                                            v98 := GenS5{F0: false, F1: 822, F2: (v66 & v66)}
                                            _ = v98
                                        }
                                        v2 = "\n\t"
                                    }
                                    if false {
                                        v99 := v3[v2]
                                        _ = v99
                                    } else {
                                        delete(v4, 1)
                                        fmt.Println(v92)
                                    }
                                    switch (v2 + "\n\t") {
                                    case "a":
                                        v2 = (("a" + v2) + ((v2 + "hello") + ((("\n\t" + "foo bar") + v2) + ((("" + (((v2 + v2) + v2) + "a")) + (v2 + ((v2 + v2) + v2))) + v2))))
                                    case "":
                                        switch -575 {
                                        case -716:
                                            v2 = ("foo bar" + "foo bar")
                                        case 39:
                                            if false {
                                                if len(v71) > 0 {
                                                    v71[0] = fi93
                                                }
                                                var v100 uint16 = 2
                                                switch ("foo bar" + v2) {
                                                case "":
                                                    if false {
                                                        v101 := GenS7{F0: genFn1(v100, 65535, (v100 | 514)), F1: (v100 ^ v100)}
                                                        if true {
                                                            v102 := GenS8{F0: 514, F1: (v2 + v2)}
                                                            var v103 int64 = v92
                                                            var v104 int8 = v74
                                                            _ = v102
                                                            _ = v103
                                                            _ = v104
                                                        } else {
                                                            fmt.Println(v2, v92)
                                                            v105 := GenS8{F0: 0, F1: "hello"}
                                                            _ = v105
                                                        }
                                                        v106 := GenS10{F0: v2, F1: 0}
                                                        var v107 string = v2
                                                        _ = v101
                                                        _ = v106
                                                        _ = v107
                                                    } else {
                                                        var v108 int16 = -620
                                                        var v109 float64 = 0.0
                                                        _ = v108
                                                        _ = v109
                                                    }
                                                case "a":
                                                    var v110 float64 = (1.5 * (0.1 / (((1e10 - 3.14159) / (-2.25 * (0.1 + -2.25))) / ((1e10 / 1.0) * 1.5))))
                                                    _ = v110
                                                }
                                                _ = len(v71)
                                                _ = v100
                                            } else {
                                                var v111 int16 = -851
                                                fmt.Println(v92, v2, v66)
                                                _ = v111
                                            }
                                        case -26:
                                            v112 := GenS6{F0: 2147483648, F1: 0}
                                            _ = v112
                                        }
                                    case "\n\t":
                                        var v113 uint8 = 255
                                        _ = v113
                                    }
                                } else {
                                    v2 = "foo bar"
                                    switch "" {
                                    case "hello":
                                        v66 = genFn1(0, 1, 32767)
                                    }
                                }
                                if true {
                                    switch 127 {
                                    case 127:
                                        v6[(fi93 + (fi72 & (^(fi72))))] = v2
                                    default:
                                        _ = len(v3)
                                    }
                                }
                                _ = fi93
                            }
                            _ = v92
                        }
                    case "hello":
                        for fi114 := 0; fi114 < 5; fi114++ {
                            if (false && ((3.14159 == 1.5) && false)) {
                                var v115 int8 = v74
                                v66 += v66
                                _ = v115
                            } else {
                                if (!true) {
                                    if false {
                                        var v116 uint64 = (v66 >> ((1) & 63))
                                        _ = v116
                                    } else {
                                        if (true || ((false || true) || (42 > 276))) {
                                            if ((fi114 * (fi72 ^ (fi114 & fi114))) != (^(fi72))) {
                                                v117 := GenS3{F0: 127, F1: 859, F2: 4294967295}
                                                _ = v117
                                            }
                                            _ = len(v1)
                                        } else {
                                            var v118 uint32 = 2147483648
                                            v66 = genFn1(1, (v118 + v118), 65535)
                                            if true {
                                                var v119 string = ((("\n\t" + "foo bar") + ("hello" + ("\n\t" + (("foo bar" + v2) + "")))) + v2)
                                                _ = v119
                                            }
                                            v6[fi114] = v2
                                            fmt.Println(v2, v2)
                                            v71 = append(v71, fi72)
                                            _ = v118
                                        }
                                        var v120 int = fi114
                                        _ = v120
                                    }
                                    fmt.Println(v74)
                                    v66 += (v66 ^ genFn1(255, 2, 403))
                                }
                            }
                            fmt.Println(v2)
                            _ = fi114
                        }
                    default:
                        switch fi72 {
                        case 4294967295:
                            for fi121 := 0; fi121 < 8; fi121++ {
                                v122 := GenS10{F0: (v2 + (v2 + "\n\t")), F1: 255}
                                _ = fi121
                                _ = v122
                            }
                        case -2147483648:
                            switch 4294967295 {
                            case 2147483647:
                                if (127 >= 255) {
                                    delete(v4, fi72)
                                }
                            case -2147483648:
                                if (-128 < v74) {
                                    for fi123 := 0; fi123 < 13; fi123++ {
                                        _ = len(v71)
                                        switch ("hello" + ("a" + v2)) {
                                        case "foo bar":
                                            v124 := v1[fi123]
                                            _ = v124
                                        default:
                                            if (42 == 9223372036854775807) {
                                                v125 := GenS7{F0: (v66 | (v66 / (((v66 << ((65535) & 63))) | 1))), F1: 255}
                                                _ = v125
                                            } else {
                                                delete(v6, fi123)
                                                switch fi72 {
                                                case 687:
                                                    fmt.Println(fi72, fi72)
                                                case -607:
                                                    v126 := GenS11{F0: true}
                                                    var v127 int16 = 397
                                                    v2 = "foo bar"
                                                    _ = v126
                                                    _ = v127
                                                case 128:
                                                    var v128 int8 = 127
                                                    _ = v128
                                                default:
                                                    _ = len(v5)
                                                }
                                            }
                                            v74 -= (v74 / (((^(v74))) | 1))
                                            var v129 int16 = 1
                                            _ = v129
                                        }
                                        v2 = (v2 + v2)
                                        _ = fi123
                                    }
                                } else {
                                    for fi130 := 0; fi130 < 5; fi130++ {
                                        if false {
                                            var v131 uint64 = (v66 & (^(v66)))
                                            v66 += v131
                                            var v132 bool = true
                                            _ = v131
                                            _ = v132
                                        } else {
                                            v133 := GenS9{F0: 0, F1: 128}
                                            v134 := GenS9{F0: 350, F1: 127}
                                            var v135 int16 = 0
                                            _ = v133
                                            _ = v134
                                            _ = v135
                                        }
                                        if (-367 < 155) {
                                            v66 += (v66 + v66)
                                        }
                                        _ = fi130
                                    }
                                    if false {
                                        for fi136 := 0; fi136 < 17; fi136++ {
                                            v66 = (v66 / ((genFn1(0, 2147483647, 1)) | 1))
                                            var v137 int = -741
                                            switch v137 {
                                            case 42:
                                                var v138 uint32 = 2147483647
                                                _ = v138
                                            case 1:
                                                v5[120] = true
                                            default:
                                                v139 := GenS1{F0: -2147483648}
                                                _ = v139
                                            }
                                            _ = fi136
                                            _ = v137
                                        }
                                        for fi140 := 0; fi140 < 18; fi140++ {
                                            v141 := GenS8{F0: -1, F1: "foo bar"}
                                            fmt.Println(v74)
                                            fmt.Println(v66, fi140)
                                            _ = fi140
                                            _ = v141
                                        }
                                    } else {
                                        v142 := GenS12{F0: (^(v74)), F1: 8}
                                        for fi143 := 0; fi143 < 3; fi143++ {
                                            if (true && true) {
                                                if (!false) {
                                                    v2 = v2
                                                }
                                                v144 := GenS5{F0: ((v74 & (^(v74))) >= v74), F1: 0, F2: 826}
                                                _ = v144
                                            }
                                            v66 -= 422
                                            var v145 uint16 = 2
                                            v71 = append(v71, 127)
                                            switch fi143 {
                                            case -854:
                                                v74 -= v74
                                            case 32767:
                                                fmt.Println(v2)
                                            case 42:
                                                switch 128 {
                                                case -9223372036854775808:
                                                    switch 671 {
                                                    case 32767:
                                                        fmt.Println(fi72)
                                                    case 32768:
                                                        var v146 bool = true
                                                        _ = v146
                                                    case 0:
                                                        var v147 float32 = 1e10
                                                        _ = v147
                                                    }
                                                case -759:
                                                    v145 = 65535
                                                default:
                                                    if true {
                                                        v74 = v74
                                                    }
                                                    if len(v71) > 0 {
                                                        v71[0] = (fi143 % ((32767) | 1))
                                                    }
                                                    _ = len(v5)
                                                }
                                            }
                                            if true {
                                                switch (fi143 % (((fi143 ^ 32767)) | 1)) {
                                                case 42:
                                                    if true {
                                                        var v148 float64 = 1e10
                                                        if (((false || (!((fi143 & fi72) > fi143))) && (false && (false && false))) || (-1 > -663)) {
                                                            if len(v71) > 0 {
                                                                v71[0] = fi72
                                                            }
                                                            v149 := GenS3{F0: 49, F1: 65535, F2: 2147483648}
                                                            _ = v149
                                                        }
                                                        _ = v148
                                                    }
                                                }
                                                v66 = genFn1(v145, 0, v145)
                                                v74 -= v74
                                            }
                                            _ = fi143
                                            _ = v145
                                        }
                                        _ = v142
                                    }
                                    for fi150 := 0; fi150 < 6; fi150++ {
                                        if true {
                                            switch (fi150 / (((-(fi72))) | 1)) {
                                            case 32767:
                                                delete(v3, "foo bar")
                                            }
                                            v2 = (v2 + v2)
                                            if false {
                                                v71 = append(v71, (^(fi72)))
                                                if false {
                                                    v2 = v2
                                                    if len(v71) > 0 {
                                                        v71[0] = fi150
                                                    }
                                                    v151 := GenS4{F0: 1, F1: 1, F2: v66}
                                                    _ = v151
                                                } else {
                                                    var v152 int64 = 1
                                                    fmt.Println(fi150, v74)
                                                    if true {
                                                        v153 := GenS2{F0: 42, F1: (v66 / ((v66) | 1))}
                                                        switch 1 {
                                                        case -225:
                                                            switch int(v152) {
                                                            case -330:
                                                                var v154 int32 = 0
                                                                _ = v154
                                                            case 128:
                                                                var v155 uint = 0
                                                                _ = v155
                                                            default:
                                                                v156 := GenS1{F0: 42}
                                                                _ = v156
                                                            }
                                                        default:
                                                            var v157 int32 = -1
                                                            _ = v157
                                                        }
                                                        var v158 int32 = 42
                                                        _ = v153
                                                        _ = v158
                                                    }
                                                    v71 = append(v71, fi150)
                                                    switch (^(fi72)) {
                                                    case 838:
                                                        var v159 bool = true
                                                        _ = v159
                                                    }
                                                    var v160 int32 = -2147483648
                                                    if false {
                                                        fmt.Println(v74)
                                                        var v161 uint64 = (^(v66))
                                                        v71 = append(v71, (fi150 >> ((4294967295) & 63)))
                                                        fmt.Println(v161, v152)
                                                        var v162 float64 = 0.1
                                                        _ = v161
                                                        _ = v162
                                                    }
                                                    _ = v152
                                                    _ = v160
                                                }
                                                switch (fi72 / (((fi72 % (((fi150 / ((fi72) | 1))) | 1))) | 1)) {
                                                case 32767:
                                                    v163 := GenS13{F0: 4294967295, F1: 9223372036854775807, F2: 998}
                                                    _ = v163
                                                }
                                            }
                                        } else {
                                            var v164 string = "\n\t"
                                            _ = v164
                                        }
                                        _ = fi150
                                    }
                                }
                            default:
                                var v165 int32 = -1
                                _ = v165
                            }
                        case 955:
                            if len(v71) >= 2 {
                                v71 = v71[1:len(v71)]
                            }
                        default:
                            var v166 float64 = ((3.14159 - 1.5) - 1e10)
                            _ = v166
                        }
                    }
                default:
                    for fi167 := 0; fi167 < 10; fi167++ {
                        v168 := GenS14{F0: 0}
                        v5[125] = true
                        v66 -= 1
                        _ = fi167
                        _ = v168
                    }
                }
                _ = v74
            }
            _ = fi72
        }
    } else {
        fmt.Println(v2)
        v66 += (^(v66))
        v71 = append(v71, 555)
    }
    _ = len(v71)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v66:", v66)
    fmt.Println("slicelen v71:", len(v71), v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v66
    _ = v71
}
