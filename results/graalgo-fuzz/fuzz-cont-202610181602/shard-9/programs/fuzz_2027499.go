package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int16 }
type GenS2 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 uint32; F1 uint8 }
type GenS7 struct { F0 int32; F1 uint32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int16; F1 int8; F2 int }
type GenS10 struct { F0 int; F1 uint }
type GenS11 struct { F0 uint64; F1 string }
type GenS12 struct { F0 uint8; F1 int64 }
func genFn0(p0 bool, p1 int) int64 {
    var v6 uint16 = 870
    if p0 {
        if (float64((3.14159 / 1e10)) != 3.14159) {
            v7 := []int{9223372036854775807}
            v6 += v6
            _ = v7
        } else {
            switch "foo bar" {
            case "a":
                var v8 int8 = 0
                _ = v8
            case "":
                fmt.Println(p0)
            case "foo bar":
                fmt.Println(p1)
            }
            if true {
                if (("hello" + "\n\t") == "hello") {
                    var v9 uint16 = v6
                    _ = v9
                }
            }
        }
        v6 = v6
    } else {
        var v10 uint = 65535
        if (1e10 <= float64(1e10)) {
            p0 = p0
            v11 := GenS1{F0: -1, F1: p0, F2: 32767}
            var v12 int8 = -128
            _ = v11
            _ = v12
        }
        _ = v10
    }
    var v13 float32 = 1.5
    _ = p0
    _ = p1
    _ = v13
    _ = v6
    return 429
}
func genFn1(p0 int16, p1 int16) bool {
    for fi14 := 0; fi14 < 19; fi14++ {
        for fi15 := 0; fi15 < 9; fi15++ {
            var v16 uint64 = 4294967295
            p1 += (p1 ^ p1)
            if true {
                var v17 uint64 = (v16 / (((v16 - v16)) | 1))
                v18 := GenS1{F0: 9223372036854775807, F1: ((((0 == 127) && ((false && true) && true)) || (fi14 >= fi15)) || true), F2: (p0 | p0)}
                fmt.Println(p0, fi14)
                _ = v17
                _ = v18
            } else {
                var v19 uint32 = 4294967295
                switch 2147483647 {
                case 661:
                    switch "a" {
                    case "\n\t":
                        var v20 uint32 = 42
                        _ = v20
                    default:
                        var v21 uint16 = 255
                        _ = v21
                    }
                case 815:
                    var v22 int = fi15
                    _ = v22
                case 639:
                    var v23 int8 = -1
                    _ = v23
                }
                _ = v19
            }
            v24 := []int{fi15, fi15}
            var v25 int8 = 0
            _ = fi15
            _ = v16
            _ = v24
            _ = v25
        }
        v26 := []int{(fi14 / ((fi14) | 1))}
        p0 = -1
        _ = fi14
        _ = v26
    }
    v27 := GenS1{F0: 0, F1: (42 >= 119), F2: 188}
    _ = p0
    _ = p1
    _ = v27
    return false
}

func main() {
    var v1 float32 = 3.14159
    var v2 int64 = -1
    var v3 uint32 = 2147483648
    v4 := make(map[uint32]int)
    v5 := make(map[int]bool)
    if ("hello" < "\n\t") {
        switch ("\n\t" + "a") {
        case "hello":
            if (v2 > (v2 >> ((95) & 63))) {
                if (!false) {
                    v3 = (v3 << ((4294967295) & 63))
                    var v28 int8 = 42
                    _ = v28
                }
                v29 := GenS2{F0: true, F1: 2, F2: v3}
                _ = v29
            }
        case "a":
            v1 = -2.25
        default:
            for fi30 := 0; fi30 < 18; fi30++ {
                var v31 float32 = (((0.1 * v1) - 1.5) + v1)
                v31 = 0.0
                delete(v4, v3)
                _ = fi30
                _ = v31
            }
        }
    } else {
        if true {
            switch -2147483648 {
            case -1:
                var v32 uint32 = (^(v3))
                _ = v32
            case -9223372036854775808:
                for fi33 := 0; fi33 < 9; fi33++ {
                    switch (fi33 | fi33) {
                    case 1:
                        for fi34 := 0; fi34 < 15; fi34++ {
                            var v35 int = fi33
                            _ = fi34
                            _ = v35
                        }
                    default:
                        v1 = v1
                    }
                    v2 = (v2 >> ((874) & 63))
                    _ = fi33
                }
            default:
                v2 = v2
            }
            var v36 uint16 = 65535
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 15; fi37++ {
        switch -259 {
        case 32767:
            var v38 float32 = float32(float64(v1))
            _ = v38
        case -1:
            v39 := GenS1{F0: v2, F1: (550 >= 1), F2: 42}
            _ = v39
        default:
            v40 := []int{}
            _ = v40
        }
        var v41 float64 = 1.5
        var v42 uint16 = 255
        _ = fi37
        _ = v41
        _ = v42
    }
    v43 := v5[-1]
    v43 = (false && (-14 <= 1))
    var v44 int = 888
    v1 = v1
    fmt.Println(v43, v2)
    v3 -= (v3 % ((2147483647) | 1))
    for fi45 := 0; fi45 < 4; fi45++ {
        for fi46 := 0; fi46 < 11; fi46++ {
            var v47 string = (("foo bar" + "\n\t") + ((("\n\t" + "a") + ("foo bar" + "a")) + ""))
            _ = fi46
            _ = v47
        }
        _ = fi45
    }
    if (!v43) {
        v48 := GenS2{F0: (((v1 != -2.25) && v43) && ((-(v44)) != (v44 << ((614) & 63)))), F1: 974, F2: (^(v3))}
        _ = v48
    } else {
        var v49 uint16 = 0
        v50 := GenS2{F0: (!((1e10 / float32((1.5 / 3.14159))) <= v1)), F1: 76, F2: (v3 | v3)}
        _ = v49
        _ = v50
    }
    for fi51 := 0; fi51 < 9; fi51++ {
        if v43 {
            if (3.14159 >= (1.5 / ((0.0 * float64((v1 - ((((float32(3.14159) / v1) * v1) / 3.14159) * ((v1 / (v1 / (v1 + (float32(3.14159) / 1.0)))) / (0.0 * v1)))))) / 3.14159))) {
                for fi52 := 0; fi52 < 15; fi52++ {
                    var v53 uint8 = 248
                    var v54 int8 = -1
                    if v43 {
                        v55 := []int{32768, 128, (fi51 | 391)}
                        switch (fi52 | 127) {
                        case -1:
                            switch ("\n\t" + "\n\t") {
                            case "\n\t":
                                var v56 uint16 = 32767
                                _ = v56
                            case "hello":
                                var v57 uint32 = 102
                                _ = v57
                            case "foo bar":
                                v3 -= (v3 - (v3 / (((^(v3))) | 1)))
                            default:
                                var v58 int32 = 1
                                fmt.Println(v54, v58, v44)
                                _ = v58
                            }
                        case 1:
                            delete(v4, 529)
                        default:
                            if (false && true) {
                                if (-2.25 > (1.5 * (3.14159 - 0.0))) {
                                    v2 += v2
                                } else {
                                    var v59 bool = (2 > v53)
                                    fmt.Println(v54, v3)
                                    _ = v59
                                }
                                v60 := GenS3{F0: 2}
                                if v43 {
                                    v44 -= fi52
                                }
                                v2 = (-(v2))
                                fmt.Println(v1, fi52, v54)
                                _ = v60
                            } else {
                                var v61 bool = v43
                                switch v44 {
                                case 970:
                                    v53 -= (^(v53))
                                case -2147483648:
                                    var v62 string = ("foo bar" + "foo bar")
                                    _ = v62
                                case 0:
                                    var v63 float64 = 3.14159
                                    _ = v63
                                }
                                var v64 int64 = v2
                                v2 += 262
                                fmt.Println(v64)
                                _ = v61
                                _ = v64
                            }
                        }
                        var v65 uint = uint(v53)
                        _ = v55
                        _ = v65
                    } else {
                        if (2147483647 < 42) {
                            var v66 int32 = 1
                            switch 4294967295 {
                            case 4294967295:
                                v67 := GenS2{F0: true, F1: 4294967295, F2: 0}
                                _ = v67
                            default:
                                if v43 {
                                    _ = len(v5)
                                } else {
                                    v54 += v54
                                    v43 = false
                                    v68 := GenS3{F0: 2}
                                    switch fi52 {
                                    case 32768:
                                        var v69 int64 = (-(v2))
                                        _ = v69
                                    case 945:
                                        v2 += (-(v2))
                                    }
                                    fmt.Println(v1)
                                    _ = v68
                                }
                                var v70 uint64 = 2
                                _ = v70
                            }
                            _ = v66
                        } else {
                            if (v43 || v43) {
                                var v71 int32 = -2147483648
                                var v72 int32 = v71
                                _ = v71
                                _ = v72
                            }
                        }
                    }
                    _ = fi52
                    _ = v53
                    _ = v54
                }
            } else {
                v2 = v2
            }
            if true {
                var v73 uint16 = 0
                for fi74 := 0; fi74 < 3; fi74++ {
                    v75 := []int{}
                    v73 = v73
                    _ = fi74
                    _ = v75
                }
                v76 := []int{}
                _ = v73
                _ = v76
            }
        } else {
            v77 := GenS3{F0: 860}
            _ = v77
        }
        var v78 int16 = 1
        if (false || ((false && (42 == 2)) && v43)) {
            v1 = -2.25
        } else {
            delete(v4, (v3 >> ((18446744073709551615) & 63)))
            v5[(fi51 + v44)] = v43
        }
        _ = fi51
        _ = v78
    }
    v79 := []int{}
    v5[560] = (false || (18446744073709551615 > 1))
    switch (v44 + v44) {
    case 42:
        switch 884 {
        case 128:
            delete(v5, v44)
        case 23:
            if ((v2 / ((v2) | 1)) != v2) {
                switch "hello" {
                case "":
                    v44 -= v44
                case "hello":
                    v2 += (v2 / ((v2) | 1))
                default:
                    v80 := v4[(v3 >> ((65535) & 63))]
                    _ = v80
                }
                for fi81 := 0; fi81 < 3; fi81++ {
                    if (1 < -2147483648) {
                        var v82 uint64 = 42
                        v83 := GenS4{F0: (42 == 1), F1: -960}
                        _ = v82
                        _ = v83
                    }
                    _ = len(v79)
                    fmt.Println(v44, v2)
                    _ = fi81
                }
                v44 = v44
            }
        case 4294967295:
            if true {
                fmt.Println(v2, v1, v3)
                if (42 != 121) {
                    switch 32768 {
                    case 380:
                        fmt.Println(v3, v3)
                    case 42:
                        switch "foo bar" {
                        case "":
                            v84 := GenS5{F0: 255, F1: 2}
                            _ = v84
                        case "a":
                            v43 = v43
                        default:
                            v44 = v44
                        }
                    case 32768:
                        v2 = v2
                    }
                    for fi85 := 0; fi85 < 16; fi85++ {
                        if false {
                            var v86 uint32 = 2
                            switch 32767 {
                            case 127:
                                var v87 uint32 = v86
                                _ = v87
                            default:
                                for fi88 := 0; fi88 < 8; fi88++ {
                                    switch int(v2) {
                                    case -2147483648:
                                        fmt.Println(v1, v86)
                                    }
                                    _ = fi88
                                }
                            }
                            delete(v5, fi85)
                            _ = v86
                        } else {
                            var v89 int32 = 1
                            _ = v89
                        }
                        v43 = v43
                        _ = fi85
                    }
                } else {
                    for fi90 := 0; fi90 < 16; fi90++ {
                        var v91 int32 = 1
                        if len(v79) >= 2 {
                            v79 = v79[1:len(v79)]
                        }
                        _ = fi90
                        _ = v91
                    }
                    var v92 float32 = (((float32(-2.25) - 3.14159) - float32(0.1)) - 0.0)
                    _ = v92
                }
                var v93 int8 = 30
                _ = v93
            } else {
                for fi94 := 0; fi94 < 6; fi94++ {
                    switch "a" {
                    case "hello":
                        v5[4294967295] = (!(((-(v2)) != v2) || false))
                    case "foo bar":
                        v2 = v2
                    default:
                        v3 = 854
                    }
                    v95 := GenS3{F0: 4294967295}
                    _ = fi94
                    _ = v95
                }
                if (2 >= 169) {
                    v79 = append(v79, (v44 * (-(v44))))
                    delete(v5, (v44 ^ v44))
                }
                if true {
                    delete(v4, v3)
                    fmt.Println(v43, v2, v43)
                    switch 327 {
                    case 127:
                        if v43 {
                            for fi96 := 0; fi96 < 4; fi96++ {
                                for fi97 := 0; fi97 < 10; fi97++ {
                                    switch (-(v44)) {
                                    case 240:
                                        var v98 uint8 = 42
                                        _ = v98
                                    case 32767:
                                        v2 -= v2
                                    case -868:
                                        if v43 {
                                            var v99 float32 = 0.0
                                            _ = v99
                                        } else {
                                            v1 = float32(((1e10 - float64(((((3.14159 / (float32((((float64(((1e10 - v1) + v1)) - (((3.14159 / 1.5) / 3.14159) - 0.1)) / 1.5) + (0.0 - 1.5))) - (float32(((0.0 + 0.0) / float64((v1 - v1)))) - 1.5))) * ((v1 - 0.0) / ((v1 * v1) - (-2.25 / v1)))) / v1) + (v1 - v1)))) / 1.5))
                                            fmt.Println(v44, v3, v43)
                                        }
                                    default:
                                        fmt.Println(v43)
                                    }
                                    _ = fi97
                                }
                                if false {
                                    if ((((0.1 - (1e10 * (3.14159 - 0.1))) + ((((-2.25 - 3.14159) * -2.25) - 0.1) - 1.5)) * float64(1.5)) >= 3.14159) {
                                        switch "a" {
                                        case "a":
                                            if (((v1 - 3.14159) / 1.5) == float32(1e10)) {
                                                var v100 uint64 = 0
                                                if false {
                                                    fmt.Println(v100, v1)
                                                    for fi101 := 0; fi101 < 16; fi101++ {
                                                        v102 := GenS1{F0: v2, F1: true, F2: int16(v3)}
                                                        var v103 int = (fi101 | 32767)
                                                        _ = fi101
                                                        _ = v102
                                                        _ = v103
                                                    }
                                                    var v104 uint32 = 991
                                                    _ = v104
                                                }
                                                for fi105 := 0; fi105 < 10; fi105++ {
                                                    if true {
                                                        if v43 {
                                                            var v106 int8 = 52
                                                            var v107 int16 = 32767
                                                            _ = v106
                                                            _ = v107
                                                        } else {
                                                            _ = len(v79)
                                                        }
                                                        var v108 uint8 = 255
                                                        _ = v108
                                                    } else {
                                                        v3 = (v3 & v3)
                                                        switch "foo bar" {
                                                        case "a":
                                                            var v109 bool = v43
                                                            _ = v109
                                                        }
                                                        v2 = genFn0(((("a" + "\n\t") + "") < "foo bar"), fi96)
                                                        var v110 string = "hello"
                                                        v111 := GenS3{F0: v100}
                                                        _ = v110
                                                        _ = v111
                                                    }
                                                    switch ("" + "hello") {
                                                    case "":
                                                        if (741 != 18446744073709551615) {
                                                            v112 := GenS4{F0: true, F1: 337}
                                                            _ = v112
                                                        }
                                                    case "foo bar":
                                                        if false {
                                                            v113 := GenS6{F0: (^(v3)), F1: 41}
                                                            fmt.Println(v3, fi96, v3)
                                                            _ = v113
                                                        } else {
                                                            if len(v79) >= 2 {
                                                                v79 = v79[1:len(v79)]
                                                            }
                                                            v100 -= (v100 & (^(v100)))
                                                        }
                                                    default:
                                                        switch (fi105 & (-(fi105))) {
                                                        case 128:
                                                            fmt.Println(fi105, v43)
                                                        default:
                                                            v79 = append(v79, (v44 / (((v44 + fi105)) | 1)))
                                                        }
                                                    }
                                                    _ = fi105
                                                }
                                                _ = v100
                                            }
                                        }
                                        for fi114 := 0; fi114 < 19; fi114++ {
                                            var v115 int64 = (^(v2))
                                            switch (fi96 / ((fi114) | 1)) {
                                            case -9223372036854775808:
                                                var v116 int = -1
                                                _ = v116
                                            default:
                                                fmt.Println(v115, v115)
                                            }
                                            var v117 int32 = int32(v2)
                                            _ = fi114
                                            _ = v115
                                            _ = v117
                                        }
                                        if (621 < 32767) {
                                            v118 := GenS7{F0: int32(v2), F1: (v3 % ((251) | 1))}
                                            for fi119 := 0; fi119 < 19; fi119++ {
                                                if (false && v43) {
                                                    if ((127 <= 0) && (18446744073709551615 == 929)) {
                                                        v44 -= fi96
                                                        v120 := GenS7{F0: 1, F1: 236}
                                                        _ = v120
                                                    }
                                                }
                                                _ = fi119
                                            }
                                            switch (("" + "hello") + "\n\t") {
                                            case "":
                                                v121 := GenS8{F0: "foo bar"}
                                                _ = v121
                                            case "\n\t":
                                                for fi122 := 0; fi122 < 16; fi122++ {
                                                    v1 += 0.0
                                                    fmt.Println(v43, fi96, fi96)
                                                    if (true && ((^(v44)) == fi122)) {
                                                        if true {
                                                            v5[(fi122 >> ((280) & 63))] = (false && (255 > 2))
                                                            v43 = (int16(v2) > 1000)
                                                            var v123 string = ("hello" + ("hello" + "hello"))
                                                            _ = v123
                                                        } else {
                                                            v43 = (!v43)
                                                        }
                                                        fmt.Println(v43, fi122)
                                                        v2 += (v2 | 9223372036854775807)
                                                    }
                                                    _ = fi122
                                                }
                                            case "foo bar":
                                                v3 = 65535
                                            default:
                                                v124 := v5[(fi96 >> ((579) & 63))]
                                                _ = v124
                                            }
                                            _ = v118
                                        } else {
                                            v5[fi96] = true
                                            var v125 uint8 = 255
                                            fmt.Println(v2, v2, v43)
                                            _ = v125
                                        }
                                    } else {
                                        v126 := GenS9{F0: -1, F1: -126, F2: (^(fi96))}
                                        if v43 {
                                            v127 := GenS10{F0: 4294967295, F1: 42}
                                            v128 := v4[v3]
                                            var v129 uint16 = 1
                                            _ = v127
                                            _ = v128
                                            _ = v129
                                        }
                                        for fi130 := 0; fi130 < 3; fi130++ {
                                            v1 = 1.5
                                            var v131 int32 = -859
                                            _ = fi130
                                            _ = v131
                                        }
                                        _ = v126
                                    }
                                } else {
                                    if true {
                                        fmt.Println(v43)
                                        _ = len(v5)
                                    } else {
                                        var v132 int8 = 1
                                        var v133 float64 = -2.25
                                        _ = v132
                                        _ = v133
                                    }
                                }
                                for fi134 := 0; fi134 < 14; fi134++ {
                                    if v43 {
                                        var v135 uint8 = 53
                                        v136 := GenS11{F0: 18446744073709551615, F1: "foo bar"}
                                        v4[(v3 & (v3 / ((v3) | 1)))] = v44
                                        _ = v135
                                        _ = v136
                                    } else {
                                        var v137 int32 = 0
                                        var v138 int64 = 9223372036854775807
                                        v1 = v1
                                        fmt.Println(fi134, fi134, v43)
                                        _ = v137
                                        _ = v138
                                    }
                                    switch (v44 << ((42) & 63)) {
                                    case 335:
                                        if (!v43) {
                                            v139 := v4[(^(v3))]
                                            _ = len(v79)
                                            v44 += fi134
                                            switch (v44 / ((fi96) | 1)) {
                                            case 4294967295:
                                                var v140 int = v139
                                                _ = v140
                                            case -2147483648:
                                                v3 += (v3 | v3)
                                            }
                                            fmt.Println(v1, v139)
                                            if (4294967295 <= 117) {
                                                var v141 int16 = 1
                                                if v43 {
                                                    v142 := GenS1{F0: 9223372036854775807, F1: (v3 > v3), F2: 1}
                                                    v3 -= v3
                                                    _ = v142
                                                }
                                                switch ("a" + ("\n\t" + ("\n\t" + ("\n\t" + "hello")))) {
                                                case "hello":
                                                    if true {
                                                        v141 -= v141
                                                    }
                                                    v143 := v4[v3]
                                                    v144 := GenS9{F0: 1, F1: 1, F2: fi96}
                                                    _ = v143
                                                    _ = v144
                                                case "":
                                                    var v145 uint8 = 243
                                                    _ = v145
                                                }
                                                _ = v141
                                            } else {
                                                if v43 {
                                                    fmt.Println(v43)
                                                }
                                                var v146 int8 = -128
                                                _ = v146
                                            }
                                            _ = v139
                                        } else {
                                            var v147 uint8 = 128
                                            _ = v147
                                        }
                                    }
                                    _ = fi134
                                }
                                _ = fi96
                            }
                            v79 = append(v79, v44)
                            v148 := GenS7{F0: 2147483647, F1: 641}
                            _ = v148
                        }
                    case 720:
                        for fi149 := 0; fi149 < 17; fi149++ {
                            v150 := GenS12{F0: uint8(v3), F1: genFn0(v43, (v44 % (((fi149 * v44)) | 1)))}
                            _ = fi149
                            _ = v150
                        }
                    }
                }
            }
        default:
            if len(v79) > 0 {
                v79[0] = v44
            }
        }
    case 127:
        v1 = (1e10 * ((((v1 - -2.25) * (v1 * 0.1)) * v1) - ((((((((float32(0.1) / (v1 + float32(((float64(v1) - 1.5) - 3.14159)))) * v1) / 1.0) * v1) - float32(1e10)) - 0.1) + (v1 / v1)) - v1)))
    case -331:
        var v151 uint = 0
        _ = v151
    default:
        v79 = append(v79, (-(v44)))
    }
    var v152 string = ("foo bar" + "")
    switch v44 {
    case -9223372036854775808:
        var v153 string = (v152 + "")
        _ = v153
    case 9223372036854775807:
        for fi154 := 0; fi154 < 9; fi154++ {
            fmt.Println(v2)
            _ = fi154
        }
    default:
        v1 += 1e10
    }
    if v43 {
        fmt.Println(v3)
        v1 -= (1e10 + (v1 * v1))
    }
    var v155 uint = 255
    fmt.Println("v1:", v1)
    fmt.Println("v152:", v152)
    fmt.Println("v155:", v155)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v79:", len(v79), v79)
    _ = v1
    _ = v152
    _ = v155
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v79
}
