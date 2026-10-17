package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 uint; F1 uint8; F2 uint }
type GenS6 struct { F0 int64; F1 int16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int16; F1 int8 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 bool; F1 int8; F2 string }
type GenS12 struct { F0 uint8; F1 uint; F2 bool }
type GenS13 struct { F0 int64; F1 uint32; F2 uint8 }
type GenS14 struct { F0 uint32; F1 uint64 }
type GenS15 struct { F0 bool }

func main() {
    var v1 string = "foo bar"
    var v2 uint = 1
    var v3 int = -152
    v4 := make(map[uint32]bool)
    v5 := []int{(-(v3)), (v3 % ((v3) | 1))}
    v3 -= (v3 + 703)
    _ = len(v4)
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 string = v1
        var v8 uint64 = 1
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint64 = 670
    if true {
        for fi10 := 0; fi10 < 17; fi10++ {
            v11 := GenS1{F0: v9}
            for fi12 := 0; fi12 < 18; fi12++ {
                v2 += (v2 * (v2 % (((v2 | 4294967295)) | 1)))
                _ = fi12
            }
            var v13 string = "hello"
            _ = fi10
            _ = v11
            _ = v13
        }
        if true {
            fmt.Println(v1, v3)
            v3 = v3
            var v14 string = (v1 + ("" + v1))
            _ = v14
        }
    } else {
        var v15 int32 = 0
        fmt.Println(v1, v1)
        for fi16 := 0; fi16 < 20; fi16++ {
            v2 = v2
            switch "a" {
            case "\n\t":
                var v17 int = fi16
                _ = v17
            case "foo bar":
                switch (v1 + (v1 + (v1 + v1))) {
                case "foo bar":
                    if true {
                        switch v1 {
                        case "foo bar":
                            switch v3 {
                            case -550:
                                v1 = v1
                            case 1:
                                v18 := GenS1{F0: (v9 ^ (v9 + (v9 / ((4294967295) | 1))))}
                                _ = v18
                            case 603:
                                var v19 int8 = 127
                                _ = v19
                            }
                        case "":
                            for fi20 := 0; fi20 < 16; fi20++ {
                                fmt.Println(v2)
                                _ = fi20
                            }
                        default:
                            switch (v3 - (v3 - (v3 & (fi16 * 971)))) {
                            case 902:
                                switch (v1 + "a") {
                                case "hello":
                                    v21 := GenS2{F0: 204, F1: fi16}
                                    _ = v21
                                case "a":
                                    if true {
                                        v22 := v4[1]
                                        _ = v22
                                    } else {
                                        v23 := GenS1{F0: v9}
                                        if (!false) {
                                            fmt.Println(v1)
                                            fmt.Println(v1, v1)
                                        }
                                        _ = v23
                                    }
                                case "\n\t":
                                    if ((861 == 42) && (100 > -128)) {
                                        v24 := GenS3{F0: 0}
                                        for fi25 := 0; fi25 < 1; fi25++ {
                                            var v26 string = "hello"
                                            v15 = (v15 | v15)
                                            v9 = (v9 >> (((v2 % (((v2 - (v2 / ((v2) | 1)))) | 1))) & 63))
                                            _ = fi25
                                            _ = v26
                                        }
                                        _ = v24
                                    } else {
                                        v27 := GenS4{F0: -9223372036854775808, F1: 73}
                                        v28 := GenS2{F0: 2, F1: (fi16 << ((v2) & 63))}
                                        _ = v27
                                        _ = v28
                                    }
                                default:
                                    var v29 uint16 = 1
                                    _ = v29
                                }
                            case 32768:
                                switch v1 {
                                case "hello":
                                    var v30 uint64 = 739
                                    _ = v30
                                default:
                                    if (false || (!(!true))) {
                                        for fi31 := 0; fi31 < 14; fi31++ {
                                            v2 -= v2
                                            _ = fi31
                                        }
                                        var v32 float64 = ((0.0 * -2.25) - ((-2.25 / 1.0) + -2.25))
                                        _ = v32
                                    }
                                }
                            case -624:
                                fmt.Println(fi16)
                            }
                        }
                    }
                default:
                    if (true && (!false)) {
                        for fi33 := 0; fi33 < 6; fi33++ {
                            v34 := GenS4{F0: 9223372036854775807, F1: 255}
                            v35 := v4[uint32(v9)]
                            fmt.Println(fi33)
                            _ = fi33
                            _ = v34
                            _ = v35
                        }
                        var v36 uint8 = 255
                        _ = v36
                    } else {
                        var v37 uint64 = v9
                        _ = v37
                    }
                }
            case "hello":
                for fi38 := 0; fi38 < 14; fi38++ {
                    switch v1 {
                    case "hello":
                        var v39 int = (fi16 | (fi38 - v3))
                        _ = v39
                    default:
                        var v40 int64 = 1
                        _ = v40
                    }
                    fmt.Println(fi16, v3, fi16)
                    v41 := v4[2147483648]
                    _ = len(v4)
                    _ = fi38
                    _ = v41
                }
            default:
                v2 = (v2 | 42)
            }
            _ = fi16
        }
        _ = v15
    }
    var v42 uint = uint(v9)
    v9 -= uint64(v42)
    if (0.1 != 0.1) {
        for fi43 := 0; fi43 < 9; fi43++ {
            if ((v2 | (v2 % (((^(v2))) | 1))) < (^(v42))) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                for fi44 := 0; fi44 < 4; fi44++ {
                    fmt.Println(fi43)
                    v3 = (^(v3))
                    _ = fi44
                }
                switch 32767 {
                case -347:
                    var v45 bool = (!false)
                    _ = v45
                case -67:
                    if false {
                        if (((2 == 1) || ((v1 + (((((v1 + (((v1 + (v1 + (v1 + "\n\t"))) + "hello") + ((v1 + v1) + "hello"))) + "hello") + "foo bar") + "\n\t") + ("" + "a"))) > v1)) && true) {
                            v46 := GenS1{F0: v9}
                            _ = v46
                        } else {
                            var v47 float64 = -2.25
                            v3 = (^(v3))
                            _ = v47
                        }
                    }
                default:
                    for fi48 := 0; fi48 < 11; fi48++ {
                        v4[722] = (true && (1 != 1))
                        if (42 >= 127) {
                            if true {
                                var v49 float64 = float64(1e10)
                                var v50 int32 = 0
                                v42 += 65535
                                _ = v49
                                _ = v50
                            }
                        } else {
                            v2 -= v2
                            v1 = v1
                            v9 = v9
                        }
                        v51 := GenS5{F0: uint(v9), F1: 42, F2: v42}
                        v3 -= fi43
                        _ = fi48
                        _ = v51
                    }
                }
            } else {
                var v52 int32 = 1
                _ = v52
            }
            _ = fi43
        }
    } else {
        for fi53 := 0; fi53 < 12; fi53++ {
            if ((v1 + v1) != v1) {
                var v54 uint16 = 156
                var v55 int32 = -1
                _ = v54
                _ = v55
            }
            var v56 uint64 = (v9 ^ (v9 | v9))
            v57 := GenS6{F0: 1, F1: 32767}
            _ = fi53
            _ = v56
            _ = v57
        }
        v58 := GenS2{F0: uint16(v9), F1: v3}
        fmt.Println(v3, v42)
        _ = v58
    }
    if false {
        fmt.Println(v3)
        var v59 string = "a"
        _ = v59
    }
    var v60 int = (v3 & v3)
    v42 -= 233
    var v61 float32 = 0.0
    switch (^(v3)) {
    case 1:
        if true {
            _ = len(v4)
            fmt.Println(v61)
            if true {
                for fi62 := 0; fi62 < 12; fi62++ {
                    for fi63 := 0; fi63 < 8; fi63++ {
                        if true {
                            var v64 float32 = (float32((0.1 + 0.1)) * v61)
                            if false {
                                v4[2] = (0 >= 128)
                            } else {
                                v2 += v2
                            }
                            fmt.Println(v2, v9)
                            _ = v64
                        }
                        fmt.Println(v9)
                        _ = fi63
                    }
                    for fi65 := 0; fi65 < 16; fi65++ {
                        v9 -= (v9 ^ v9)
                        var v66 float32 = v61
                        if (0.1 < v66) {
                            switch ("\n\t" + "a") {
                            case "":
                                if (-1 == 875) {
                                    if false {
                                        if true {
                                            fmt.Println(v61, fi62, v9)
                                            v67 := GenS6{F0: -1, F1: 42}
                                            switch v1 {
                                            case "a":
                                                v9 = (v9 >> ((v2) & 63))
                                            default:
                                                v68 := v4[0]
                                                _ = v68
                                            }
                                            _ = v67
                                        } else {
                                            var v69 bool = true
                                            var v70 uint64 = v9
                                            _ = v69
                                            _ = v70
                                        }
                                        if (!(false && (1 == 42))) {
                                            if (!false) {
                                                fmt.Println(v2)
                                            } else {
                                                var v71 bool = true
                                                switch v1 {
                                                case "hello":
                                                    v42 -= 503
                                                default:
                                                    var v72 int8 = 1
                                                    _ = v72
                                                }
                                                _ = v71
                                            }
                                            fmt.Println(v1, v1, v1)
                                            v73 := v4[2]
                                            _ = v73
                                        }
                                        var v74 float32 = (float32(float64(v66)) + v61)
                                        _ = v74
                                    } else {
                                        var v75 bool = (!false)
                                        _ = v75
                                    }
                                }
                            }
                        }
                        _ = fi65
                        _ = v66
                    }
                    fmt.Println(fi62)
                    _ = fi62
                }
                v76 := GenS7{F0: 1}
                v77 := GenS5{F0: v2, F1: 2, F2: v42}
                _ = v76
                _ = v77
            }
        } else {
            if (true || (!true)) {
                var v78 int16 = -548
                fmt.Println(v61, v61)
                if (!true) {
                    if false {
                        var v79 int16 = -32768
                        v80 := GenS7{F0: 1000}
                        _ = len(v5)
                        _ = v79
                        _ = v80
                    } else {
                        var v81 int8 = 127
                        _ = v81
                    }
                    for fi82 := 0; fi82 < 20; fi82++ {
                        v60 = fi82
                        _ = fi82
                    }
                    v83 := v4[2147483647]
                    _ = v83
                } else {
                    if false {
                        if (v9 != v9) {
                            var v84 float64 = 1e10
                            _ = len(v5)
                            _ = v84
                        } else {
                            var v85 uint = (v2 | v42)
                            var v86 uint = 471
                            var v87 uint = 891
                            _ = v85
                            _ = v86
                            _ = v87
                        }
                        fmt.Println(v60)
                        switch v3 {
                        case -444:
                            var v88 uint16 = 0
                            _ = v88
                        default:
                            var v89 int64 = int64(v2)
                            _ = v89
                        }
                    }
                    if true {
                        delete(v4, 2147483647)
                    } else {
                        v1 = "hello"
                        fmt.Println(v2, v60, v60)
                    }
                    if (false && false) {
                        v90 := GenS5{F0: 5, F1: 62, F2: v42}
                        if false {
                            fmt.Println(v78, v3)
                        }
                        var v91 uint32 = 52
                        _ = v90
                        _ = v91
                    } else {
                        v3 += v60
                        var v92 uint16 = 42
                        fmt.Println(v78, v42, v42)
                        _ = v92
                    }
                }
                _ = v78
            } else {
                switch "hello" {
                case "":
                    v60 += (v3 | (v3 / ((int(v42)) | 1)))
                case "foo bar":
                    v93 := GenS8{F0: -1}
                    _ = v93
                default:
                    var v94 int32 = 136
                    _ = v94
                }
                for fi95 := 0; fi95 < 14; fi95++ {
                    fmt.Println(v60, v60, v1)
                    _ = fi95
                }
                if (!((v9 > 2) || (!false))) {
                    for fi96 := 0; fi96 < 18; fi96++ {
                        fmt.Println(v1, v2, v61)
                        switch 0 {
                        case 32767:
                            if false {
                                var v97 uint32 = 145
                                _ = v97
                            }
                        case 32768:
                            v98 := GenS1{F0: 4294967295}
                            _ = v98
                        }
                        _ = fi96
                    }
                    fmt.Println(v9, v3)
                    var v99 int32 = -124
                    _ = v99
                } else {
                    var v100 uint8 = 128
                    for fi101 := 0; fi101 < 7; fi101++ {
                        v100 = (v100 & (v100 | 154))
                        _ = fi101
                    }
                    var v102 int = v3
                    _ = v100
                    _ = v102
                }
            }
            v1 = ("foo bar" + v1)
        }
    case 32767:
        for fi103 := 0; fi103 < 20; fi103++ {
            switch v60 {
            case 42:
                for fi104 := 0; fi104 < 9; fi104++ {
                    if (true || (false && (!false))) {
                        v105 := GenS9{F0: 95, F1: 100}
                        var v106 uint8 = 229
                        var v107 uint8 = 251
                        _ = v105
                        _ = v106
                        _ = v107
                    } else {
                        switch "a" {
                        case "\n\t":
                            var v108 uint = (^(v42))
                            _ = v108
                        case "a":
                            v109 := GenS3{F0: 0}
                            _ = v109
                        case "hello":
                            var v110 string = ("a" + "\n\t")
                            _ = v110
                        default:
                            v111 := GenS4{F0: 1, F1: uint8(v42)}
                            _ = v111
                        }
                        if true {
                            if true {
                                v1 = v1
                                v3 += (v3 + v3)
                            } else {
                                switch ("foo bar" + ((v1 + (v1 + (v1 + ("\n\t" + v1)))) + v1)) {
                                case "a":
                                    var v112 int16 = 0
                                    _ = v112
                                }
                                fmt.Println(v60)
                                switch v1 {
                                case "a":
                                    v61 += 1.5
                                case "":
                                    v60 = v3
                                case "hello":
                                    switch v1 {
                                    case "hello":
                                        var v113 uint8 = 255
                                        _ = v113
                                    case "":
                                        fmt.Println(v3, v1, v42)
                                    case "a":
                                        if true {
                                            if (1e10 <= ((v61 * 1.5) + v61)) {
                                                v9 -= (^(v9))
                                                v3 += (fi104 | 32768)
                                            } else {
                                                var v114 uint16 = 42
                                                v115 := GenS1{F0: v9}
                                                v60 = v3
                                                var v116 float64 = 1e10
                                                _ = v114
                                                _ = v115
                                                _ = v116
                                            }
                                        }
                                    default:
                                        switch (fi103 >> ((690) & 63)) {
                                        case -1:
                                            var v117 uint64 = v9
                                            _ = v117
                                        default:
                                            if true {
                                                switch v3 {
                                                case -39:
                                                    v118 := GenS10{F0: v3}
                                                    _ = v118
                                                case 128:
                                                    if (uint(v9) == (v42 * (v2 | (v2 << ((255) & 63))))) {
                                                        switch (v3 << (((v42 >> (((^(v2))) & 63))) & 63)) {
                                                        case 32767:
                                                            switch int(v9) {
                                                            case -919:
                                                                v119 := GenS11{F0: (509 == fi103), F1: 33, F2: v1}
                                                                _ = v119
                                                            case 4294967295:
                                                                v120 := GenS12{F0: 147, F1: v42, F2: true}
                                                                _ = v120
                                                            case 2147483647:
                                                                if false {
                                                                    if (v2 == v42) {
                                                                        switch fi103 {
                                                                        case 128:
                                                                            v61 = v61
                                                                        case -2147483648:
                                                                            if (655 > 65535) {
                                                                                var v121 int8 = 7
                                                                                v42 -= uint(v9)
                                                                                if (42 != 1) {
                                                                                    if (226 != 587) {
                                                                                        v121 += (v121 >> (((v2 - 0)) & 63))
                                                                                        fmt.Println(v9)
                                                                                        v122 := GenS12{F0: 2, F1: (v42 + v42), F2: true}
                                                                                        _ = v122
                                                                                    } else {
                                                                                        if false {
                                                                                            v121 -= -1
                                                                                        }
                                                                                        _ = len(v5)
                                                                                    }
                                                                                } else {
                                                                                    var v123 int = (^(v60))
                                                                                    _ = v123
                                                                                }
                                                                                if ((v2 < v42) || ((-(v121)) == (v121 >> ((v42) & 63)))) {
                                                                                    var v124 float64 = 0.1
                                                                                    _ = v124
                                                                                }
                                                                                _ = v121
                                                                            }
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v1)
                                                                        if (false || ((254 >= 42) && ((v3 & fi103) != fi103))) {
                                                                            fmt.Println(v1)
                                                                        }
                                                                    }
                                                                } else {
                                                                    var v125 int16 = 42
                                                                    if true {
                                                                        v126 := v4[594]
                                                                        _ = v126
                                                                    }
                                                                    v127 := GenS13{F0: -9223372036854775808, F1: 2, F2: 42}
                                                                    _ = v125
                                                                    _ = v127
                                                                }
                                                            default:
                                                                var v128 uint16 = 65535
                                                                _ = v128
                                                            }
                                                        }
                                                        var v129 uint = (v2 | (v2 * 0))
                                                        _ = v129
                                                    } else {
                                                        v130 := v4[65535]
                                                        _ = v130
                                                    }
                                                case -501:
                                                    var v131 uint8 = 2
                                                    _ = v131
                                                default:
                                                    var v132 int8 = 42
                                                    _ = v132
                                                }
                                                var v133 int = 32767
                                                var v134 float32 = (((v61 / (v61 * ((v61 * v61) + -2.25))) + 0.0) * 0.0)
                                                var v135 float32 = v134
                                                if false {
                                                    if true {
                                                        fmt.Println(v133, v9)
                                                        fmt.Println(v60, v1)
                                                        if true {
                                                            if false {
                                                                if (true || (!false)) {
                                                                    var v136 uint64 = v9
                                                                    var v137 float64 = 0.0
                                                                    fmt.Println(v135, v2, v134)
                                                                    _ = v136
                                                                    _ = v137
                                                                }
                                                                if (0.0 <= 1e10) {
                                                                    switch "foo bar" {
                                                                    case "foo bar":
                                                                        v61 += float32(1.5)
                                                                    case "a":
                                                                        if true {
                                                                            v138 := GenS13{F0: -21, F1: 4294967295, F2: uint8(v9)}
                                                                            v2 = (v2 % (((v2 - v42)) | 1))
                                                                            _ = v138
                                                                        }
                                                                        var v139 float32 = ((v134 - (v134 / v134)) - v134)
                                                                        _ = v139
                                                                    case "hello":
                                                                        delete(v4, 2)
                                                                    }
                                                                    v134 = v134
                                                                    var v140 int32 = -857
                                                                    _ = v140
                                                                }
                                                                v134 = v61
                                                            } else {
                                                                v141 := GenS5{F0: (v2 - v42), F1: 255, F2: 411}
                                                                _ = len(v5)
                                                                v42 += (^(v42))
                                                                _ = v141
                                                            }
                                                            var v142 uint64 = v9
                                                            switch (v1 + (((v1 + "\n\t") + (v1 + "foo bar")) + "\n\t")) {
                                                            case "":
                                                                v1 = v1
                                                            case "foo bar":
                                                                var v143 int16 = 32767
                                                                _ = v143
                                                            case "a":
                                                                if len(v5) > 0 {
                                                                    v5[0] = v3
                                                                }
                                                            }
                                                            _ = v142
                                                        } else {
                                                            switch v3 {
                                                            case 302:
                                                                if (1.5 == 0.1) {
                                                                    v144 := v4[71]
                                                                    if v144 {
                                                                        v9 += v9
                                                                        v5 = append(v5, -2147483648)
                                                                        if false {
                                                                            delete(v4, 343)
                                                                            v61 = (v135 + v134)
                                                                        }
                                                                    } else {
                                                                        v133 = fi103
                                                                        if v144 {
                                                                            var v145 float64 = 3.14159
                                                                            switch v1 {
                                                                            case "a":
                                                                                v146 := GenS3{F0: 4294967295}
                                                                                _ = v146
                                                                            case "foo bar":
                                                                                if len(v5) >= 2 {
                                                                                    v5 = v5[1:len(v5)]
                                                                                }
                                                                            default:
                                                                                switch v1 {
                                                                                case "a":
                                                                                    var v147 uint16 = 32767
                                                                                    _ = v147
                                                                                default:
                                                                                    switch -501 {
                                                                                    case 76:
                                                                                        _ = len(v4)
                                                                                    }
                                                                                }
                                                                            }
                                                                            _ = v145
                                                                        } else {
                                                                            var v148 int32 = -2147483648
                                                                            v149 := GenS5{F0: uint(v9), F1: 42, F2: 18446744073709551615}
                                                                            v135 = 1e10
                                                                            _ = v148
                                                                            _ = v149
                                                                        }
                                                                        v150 := GenS5{F0: v42, F1: 125, F2: v2}
                                                                        _ = v150
                                                                    }
                                                                    v5 = append(v5, 2147483647)
                                                                    _ = v144
                                                                }
                                                            default:
                                                                fmt.Println(v1, v3, v135)
                                                            }
                                                        }
                                                        var v151 uint64 = (v9 % ((v9) | 1))
                                                        var v152 string = ((v1 + "") + "")
                                                        _ = v151
                                                        _ = v152
                                                    } else {
                                                        v9 = 18446744073709551615
                                                    }
                                                    var v153 int64 = 1
                                                    _ = v153
                                                } else {
                                                    var v154 uint8 = 2
                                                    _ = v154
                                                }
                                                var v155 int16 = -32768
                                                _ = v133
                                                _ = v134
                                                _ = v135
                                                _ = v155
                                            } else {
                                                v5 = append(v5, v3)
                                                var v156 int16 = 42
                                                _ = v156
                                            }
                                        }
                                    }
                                }
                                delete(v4, 2147483648)
                                fmt.Println(v61, v61)
                            }
                            v157 := GenS14{F0: 487, F1: v9}
                            switch ((((v1 + (v1 + (("hello" + ("foo bar" + v1)) + (((v1 + "hello") + ("\n\t" + (v1 + v1))) + (v1 + (v1 + "\n\t")))))) + "\n\t") + (("" + ((v1 + "foo bar") + (((v1 + v1) + v1) + v1))) + (v1 + v1))) + v1) {
                            case "a":
                                fmt.Println(v60, v61, v3)
                            case "foo bar":
                                if false {
                                    var v158 int16 = 1
                                    var v159 int16 = v158
                                    _ = v158
                                    _ = v159
                                } else {
                                    v2 = 2
                                    v60 -= -2147483648
                                    var v160 float32 = float32((3.14159 - (((0.1 * 3.14159) + -2.25) - 0.1)))
                                    _ = v160
                                }
                            case "":
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                            _ = v157
                        }
                    }
                    v161 := GenS15{F0: ((v42 - 981) >= 255)}
                    v1 = "\n\t"
                    _ = fi104
                    _ = v161
                }
            default:
                if false {
                    switch (v1 + (v1 + ("" + v1))) {
                    case "\n\t":
                        v5 = append(v5, v3)
                    case "foo bar":
                        v162 := GenS3{F0: 820}
                        _ = v162
                    }
                    v5 = append(v5, v3)
                    var v163 uint64 = (v9 / ((v9) | 1))
                    _ = v163
                } else {
                    v164 := GenS3{F0: 2147483648}
                    var v165 bool = false
                    _ = v164
                    _ = v165
                }
            }
            v9 -= (v9 & (v9 << (((v42 << (((^(v42))) & 63))) & 63)))
            var v166 bool = (!(42 <= 1))
            _ = fi103
            _ = v166
        }
    case 128:
        v1 = v1
    }
    if (153 != 1) {
        switch "\n\t" {
        case "a":
            var v167 string = (v1 + (((v1 + "\n\t") + v1) + v1))
            _ = v167
        case "\n\t":
            switch v1 {
            case "\n\t":
                var v168 uint32 = 42
                _ = v168
            case "hello":
                switch v1 {
                case "a":
                    _ = len(v4)
                case "foo bar":
                    var v169 int64 = 0
                    _ = v169
                }
            default:
                _ = len(v5)
            }
        case "":
            var v170 uint8 = 128
            _ = v170
        }
    } else {
        if false {
            v9 += v9
        }
    }
    for fi171 := 0; fi171 < 10; fi171++ {
        var v172 int16 = 1000
        _ = fi171
        _ = v172
    }
    if true {
        if ((v2 - (^(v42))) >= 219) {
            v9 -= v9
        } else {
            var v173 int8 = -74
            if true {
                if (1e10 == (0.0 * (3.14159 * -2.25))) {
                    for fi174 := 0; fi174 < 6; fi174++ {
                        v175 := GenS7{F0: 42}
                        _ = fi174
                        _ = v175
                    }
                }
                var v176 uint8 = 127
                _ = v176
            }
            _ = v173
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v60
    _ = v61
    _ = v9
}
