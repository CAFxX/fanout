package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 uint; F2 int64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint32; F1 uint32 }
type GenS7 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 bool; F1 uint; F2 int64 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint32; F1 uint16 }
type GenS12 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int }

func main() {
    var v1 uint = 65535
    v2 := []uint32{65535, 65535, 2147483648}
    var v3 uint = v1
    switch ("" + "hello") {
    case "a":
        var v4 uint16 = 939
        _ = v4
    case "foo bar":
        for fi5 := 0; fi5 < 15; fi5++ {
            v6 := GenS1{F0: 2}
            if (0 < 0) {
                if true {
                    if false {
                        if (4294967295 != 0) {
                            var v7 int16 = 512
                            var v8 uint16 = 255
                            _ = v7
                            _ = v8
                        } else {
                            if (!(!((207 > fi5) && false))) {
                                var v9 float32 = (float32((3.14159 * -2.25)) - -2.25)
                                _ = v9
                            }
                        }
                        for fi10 := 0; fi10 < 19; fi10++ {
                            if true {
                                v1 = v3
                            } else {
                                var v11 int64 = 42
                                if false {
                                    v3 -= (v3 * (v1 / (((v1 + v1)) | 1)))
                                    if len(v2) > 0 {
                                        v2[0] = 42
                                    }
                                    fmt.Println(fi5, fi10)
                                } else {
                                    v12 := GenS2{F0: 802, F1: 1, F2: 0}
                                    v2 = append(v2, 272)
                                    _ = v12
                                }
                                v11 = -9223372036854775808
                                _ = v11
                            }
                            v3 -= 1
                            if false {
                                if true {
                                    v13 := GenS3{F0: 42}
                                    var v14 int8 = 122
                                    _ = v13
                                    _ = v14
                                }
                            }
                            switch fi5 {
                            case -222:
                                if (!false) {
                                    fmt.Println(v3, fi5, v3)
                                } else {
                                    var v15 uint64 = 2
                                    var v16 uint8 = 154
                                    _ = v15
                                    _ = v16
                                }
                            case 42:
                                switch "\n\t" {
                                case "\n\t":
                                    v17 := GenS4{F0: 42, F1: (v3 % ((v3) | 1)), F2: 0}
                                    _ = v17
                                case "a":
                                    var v18 float32 = float32((1.5 * (((1.5 / -2.25) / -2.25) + 0.0)))
                                    _ = v18
                                case "":
                                    v19 := GenS1{F0: 18446744073709551615}
                                    _ = v19
                                }
                            default:
                                v2 = append(v2, 0)
                            }
                            if ((1e10 <= 3.14159) || (-32768 != 1000)) {
                                var v20 uint32 = 482
                                fmt.Println(v1)
                                _ = v20
                            }
                            _ = fi10
                        }
                        switch fi5 {
                        case -580:
                            switch fi5 {
                            case -9223372036854775808:
                                switch (((("foo bar" + ("a" + "hello")) + (("\n\t" + "foo bar") + "")) + ("" + "\n\t")) + (("\n\t" + ("foo bar" + ("" + "hello"))) + "\n\t")) {
                                case "":
                                    for fi21 := 0; fi21 < 3; fi21++ {
                                        var v22 int32 = 1
                                        if (1 != 1) {
                                            var v23 uint64 = 18446744073709551615
                                            if (513 == v23) {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                switch ("hello" + "a") {
                                                case "":
                                                    var v24 int64 = 0
                                                    _ = v24
                                                }
                                                var v25 bool = true
                                                _ = v25
                                            }
                                            if (!(!false)) {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                switch (^(fi21)) {
                                                case 559:
                                                    _ = len(v2)
                                                }
                                            } else {
                                                v26 := GenS5{F0: 32767}
                                                var v27 float64 = (3.14159 * (0.1 + 0.1))
                                                _ = v26
                                                _ = v27
                                            }
                                            _ = v23
                                        }
                                        var v28 bool = (0 <= 0)
                                        _ = fi21
                                        _ = v22
                                        _ = v28
                                    }
                                default:
                                    var v29 float32 = -2.25
                                    _ = v29
                                }
                            case 303:
                                var v30 uint16 = 42
                                _ = v30
                            case 128:
                                for fi31 := 0; fi31 < 6; fi31++ {
                                    fmt.Println(fi31)
                                    if false {
                                        var v32 float32 = 0.0
                                        v33 := GenS5{F0: 615}
                                        _ = v32
                                        _ = v33
                                    } else {
                                        var v34 string = "hello"
                                        v34 = "foo bar"
                                        v35 := GenS4{F0: 628, F1: 2, F2: -1}
                                        v3 = v1
                                        _ = v34
                                        _ = v35
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = fi31
                                }
                            default:
                                if (-63 != -128) {
                                    var v36 uint16 = 0
                                    _ = v36
                                }
                            }
                        case -109:
                            if false {
                                var v37 uint8 = 255
                                if ("hello" >= "\n\t") {
                                    v38 := GenS4{F0: 1, F1: v3, F2: 308}
                                    fmt.Println(v3, v3)
                                    var v39 string = "foo bar"
                                    _ = v38
                                    _ = v39
                                } else {
                                    if true {
                                        for fi40 := 0; fi40 < 4; fi40++ {
                                            if len(v2) > 0 {
                                                v2[0] = 2
                                            }
                                            v41 := GenS5{F0: 65535}
                                            var v42 uint64 = 0
                                            _ = fi40
                                            _ = v41
                                            _ = v42
                                        }
                                    }
                                }
                                v3 = v3
                                _ = v37
                            }
                        default:
                            if (1 < -9223372036854775808) {
                                v43 := GenS3{F0: 1}
                                v44 := GenS6{F0: 0, F1: 787}
                                for fi45 := 0; fi45 < 9; fi45++ {
                                    switch (fi45 ^ (fi45 / ((0) | 1))) {
                                    case -580:
                                        fmt.Println(v1, fi5)
                                    case 2147483647:
                                        if false {
                                            fmt.Println(fi45, v1, v3)
                                        } else {
                                            v3 -= (v3 * 255)
                                            var v46 uint64 = 1
                                            switch (("foo bar" + "\n\t") + "foo bar") {
                                            case "hello":
                                                v2 = append(v2, 65535)
                                            case "\n\t":
                                                if true {
                                                    fmt.Println(v1, v1)
                                                    if (true || (false && ((v46 - uint64(v1)) >= v46))) {
                                                        if true {
                                                            _ = len(v2)
                                                        } else {
                                                            var v47 bool = (!false)
                                                            _ = v47
                                                        }
                                                    }
                                                    fmt.Println(fi5, fi5)
                                                    switch (fi45 % (((fi5 / ((fi45) | 1))) | 1)) {
                                                    case 9223372036854775807:
                                                        if true {
                                                            if false {
                                                                switch (-(fi5)) {
                                                                case 2147483647:
                                                                    var v48 int = (fi5 ^ -1)
                                                                    _ = v48
                                                                case 560:
                                                                    if len(v2) > 0 {
                                                                        v2[0] = 461
                                                                    }
                                                                }
                                                            }
                                                            if true {
                                                                var v49 string = ("" + "")
                                                                _ = v49
                                                            } else {
                                                                v1 -= 2
                                                            }
                                                        } else {
                                                            var v50 float32 = (3.14159 / 1.0)
                                                            _ = v50
                                                        }
                                                    case 32767:
                                                        v51 := GenS7{F0: 614, F1: 9223372036854775807, F2: 2}
                                                        _ = v51
                                                    case 42:
                                                        v1 -= v1
                                                    default:
                                                        v52 := GenS8{F0: "\n\t"}
                                                        _ = v52
                                                    }
                                                    var v53 uint64 = v46
                                                    v54 := GenS4{F0: (^(v53)), F1: (^(v3)), F2: 9223372036854775807}
                                                    _ = v53
                                                    _ = v54
                                                }
                                                var v55 float64 = (1.5 + 3.14159)
                                                v2 = append(v2, 973)
                                                _ = v55
                                            default:
                                                var v56 uint16 = 408
                                                _ = v56
                                            }
                                            _ = v46
                                        }
                                    case 32768:
                                        fmt.Println(fi45)
                                    default:
                                        v1 = (v3 | v1)
                                    }
                                    v3 = (v1 * (v3 & 675))
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = fi45
                                }
                                _ = v43
                                _ = v44
                            } else {
                                for fi57 := 0; fi57 < 10; fi57++ {
                                    var v58 uint = 65535
                                    if len(v2) > 0 {
                                        v2[0] = 2
                                    }
                                    if true {
                                        var v59 int = fi57
                                        _ = v59
                                    }
                                    _ = fi57
                                    _ = v58
                                }
                                if false {
                                    _ = len(v2)
                                    v60 := GenS5{F0: 0}
                                    var v61 uint32 = 0
                                    _ = v60
                                    _ = v61
                                }
                                for fi62 := 0; fi62 < 17; fi62++ {
                                    switch -2147483648 {
                                    case 319:
                                        v63 := GenS9{F0: true, F1: (v1 | (v3 ^ (v3 & v1))), F2: 1}
                                        _ = v63
                                    case -185:
                                        v1 = v1
                                    }
                                    _ = fi62
                                }
                            }
                        }
                    } else {
                        fmt.Println(v3, v1)
                        var v64 float32 = (((float32((1.5 / 1e10)) - 3.14159) / -2.25) + -2.25)
                        _ = v64
                    }
                } else {
                    var v65 int8 = -128
                    _ = v65
                }
            }
            v66 := GenS6{F0: 1, F1: 1}
            _ = fi5
            _ = v6
            _ = v66
        }
    case "":
        if true {
            v1 += (v1 % (((v1 + (v3 & v1))) | 1))
            for fi67 := 0; fi67 < 13; fi67++ {
                switch (fi67 | (fi67 * fi67)) {
                case 236:
                    v1 -= v3
                case 128:
                    var v68 uint16 = 2
                    _ = v68
                }
                if true {
                    fmt.Println(fi67)
                    switch "hello" {
                    case "a":
                        var v69 int16 = 42
                        _ = v69
                    }
                    v70 := GenS10{F0: 18446744073709551615}
                    _ = v70
                } else {
                    fmt.Println(v3, v3)
                    switch "" {
                    case "a":
                        var v71 int16 = 1
                        _ = v71
                    default:
                        if len(v2) > 0 {
                            v2[0] = 0
                        }
                    }
                    fmt.Println(fi67)
                }
                _ = fi67
            }
        } else {
            switch 32767 {
            case -221:
                for fi72 := 0; fi72 < 18; fi72++ {
                    var v73 int32 = 45
                    _ = len(v2)
                    _ = fi72
                    _ = v73
                }
            }
            if (2 < 42) {
                switch 830 {
                case 9223372036854775807:
                    var v74 float64 = 0.0
                    _ = v74
                case 32767:
                    if true {
                        for fi75 := 0; fi75 < 2; fi75++ {
                            var v76 uint8 = 0
                            var v77 string = "a"
                            _ = fi75
                            _ = v76
                            _ = v77
                        }
                        var v78 uint16 = 32767
                        _ = v78
                    } else {
                        var v79 int64 = 1
                        fmt.Println(v79)
                        if (true || false) {
                            v3 -= 255
                            var v80 float64 = 1e10
                            _ = v80
                        }
                        _ = v79
                    }
                case 32768:
                    var v81 float32 = 3.14159
                    _ = v81
                }
            }
            if false {
                var v82 uint16 = 32767
                var v83 int8 = 0
                _ = v82
                _ = v83
            } else {
                if (!(false && (0 < -236))) {
                    for fi84 := 0; fi84 < 19; fi84++ {
                        switch fi84 {
                        case 42:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        for fi85 := 0; fi85 < 12; fi85++ {
                            fmt.Println(v3, v1)
                            v1 += v1
                            _ = fi85
                        }
                        _ = fi84
                    }
                    if (true && ((4294967295 > 775) && true)) {
                        v3 = v1
                    }
                } else {
                    var v86 int64 = 896
                    _ = v86
                }
                v3 += v1
            }
        }
    }
    var v87 bool = false
    if false {
        for fi88 := 0; fi88 < 13; fi88++ {
            if len(v2) > 0 {
                v2[0] = 2147483647
            }
            if (!true) {
                _ = len(v2)
                v89 := GenS10{F0: 4294967295}
                if ((float32((0.0 * (float64(3.14159) / (0.0 - float64(float32(float64(0.1))))))) * 1.5) >= float32((3.14159 * 0.0))) {
                    var v90 uint32 = 42
                    for fi91 := 0; fi91 < 6; fi91++ {
                        var v92 string = "a"
                        _ = fi91
                        _ = v92
                    }
                    _ = v90
                }
                _ = v89
            }
            _ = fi88
        }
        v1 += 65535
        v87 = v87
    } else {
        var v93 int16 = -32768
        _ = v93
    }
    var v94 uint64 = 18446744073709551615
    v1 = v3
    if false {
        v95 := GenS11{F0: 2147483647, F1: 32767}
        v3 = v3
        var v96 bool = (uint(v94) > (v1 / (((v1 - 2)) | 1)))
        _ = v95
        _ = v96
    } else {
        if (!(!true)) {
            fmt.Println(v94, v1)
        } else {
            if len(v2) > 0 {
                v2[0] = 0
            }
            if (!(v3 <= 65535)) {
                v97 := GenS12{F0: 1, F1: (v87 && v87), F2: 42}
                _ = v97
            } else {
                _ = len(v2)
            }
            if false {
                if (0.0 >= 3.14159) {
                    v98 := GenS8{F0: "hello"}
                    var v99 uint32 = 0
                    v1 = 863
                    _ = v98
                    _ = v99
                }
                var v100 float64 = (((0.1 - (-2.25 / -2.25)) - 1e10) + (1.5 - -2.25))
                _ = v100
            } else {
                v94 = 0
                v101 := GenS2{F0: 607, F1: -128, F2: 58}
                v1 += 0
                _ = v101
            }
        }
        v1 = 18446744073709551615
        if ((!(1000 != 42)) || v87) {
            for fi102 := 0; fi102 < 4; fi102++ {
                fmt.Println(v1, fi102)
                _ = fi102
            }
            v94 += 587
        }
    }
    v3 += v3
    for fi103 := 0; fi103 < 19; fi103++ {
        fmt.Println(v94, v94)
        _ = fi103
    }
    v104 := GenS1{F0: v94}
    v94 = (v94 * v94)
    fmt.Println(v87, v87)
    v2 = append(v2, 42)
    if false {
        v105 := GenS6{F0: 2147483647, F1: 304}
        _ = v105
    }
    if v87 {
        v94 -= (v94 | 0)
        var v106 uint = v1
        fmt.Println(v106, v94)
        _ = v106
    } else {
        var v107 int = -9223372036854775808
        _ = v107
    }
    for fi108 := 0; fi108 < 15; fi108++ {
        for fi109 := 0; fi109 < 14; fi109++ {
            if (v87 && v87) {
                if false {
                    fmt.Println(v94, v94)
                    v1 += v3
                    _ = len(v2)
                } else {
                    switch (fi109 % (((fi108 & (fi108 - fi109))) | 1)) {
                    case 1:
                        switch ("foo bar" + "hello") {
                        case "foo bar":
                            v110 := GenS13{F0: 2}
                            _ = v110
                        case "a":
                            v87 = (255 == 2)
                        }
                    case 0:
                        v94 = v94
                    case 894:
                        fmt.Println(fi109, fi108, v87)
                    }
                    v2 = append(v2, 0)
                }
                fmt.Println(v1, v3, v87)
            } else {
                var v111 bool = false
                if false {
                    if true {
                        if false {
                            v3 = v1
                        }
                        if (!((fi109 << ((255) & 63)) >= (^(fi108)))) {
                            if (!(4294967295 > 460)) {
                                fmt.Println(v111, fi108, v111)
                                v3 -= 4294967295
                            }
                            v87 = v87
                            v94 = (v94 | (^(v94)))
                        }
                        fmt.Println(fi109, fi108)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v112 uint8 = 0
                        fmt.Println(fi108, v3, v3)
                        v113 := GenS8{F0: ""}
                        _ = v112
                        _ = v113
                    } else {
                        v114 := GenS3{F0: (v94 >> ((65535) & 63))}
                        _ = v114
                    }
                    var v115 bool = false
                    var v116 uint8 = 2
                    _ = v115
                    _ = v116
                }
                v94 -= v94
                _ = v111
            }
            if (-128 < -9) {
                v1 = (v1 ^ v3)
                var v117 int16 = 1
                if (317 > 42) {
                    if ((v1 >> ((uint(v94)) & 63)) >= (v1 & v3)) {
                        fmt.Println(fi109, fi109, v117)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v118 uint16 = 65535
                        _ = v118
                    } else {
                        v117 = 42
                    }
                    switch ((("" + "foo bar") + "\n\t") + "\n\t") {
                    case "a":
                        if ((v87 && true) || false) {
                            var v119 int32 = 1
                            _ = v119
                        }
                    case "\n\t":
                        _ = len(v2)
                    case "hello":
                        if ((!(v87 && v87)) && v87) {
                            v94 += 18446744073709551615
                            var v120 uint16 = 1
                            v121 := GenS10{F0: uint64(v120)}
                            _ = v120
                            _ = v121
                        } else {
                            var v122 float32 = -2.25
                            _ = v122
                        }
                    }
                    fmt.Println(fi108, v1)
                    v94 -= (v94 + v94)
                }
                _ = v117
            } else {
                var v123 uint = v1
                switch ("hello" + "a") {
                case "hello":
                    if false {
                        v94 -= 4294967295
                    } else {
                        var v124 uint = v1
                        var v125 uint64 = (^(v94))
                        fmt.Println(v125, v123, fi109)
                        _ = v124
                        _ = v125
                    }
                    if v87 {
                        v126 := GenS3{F0: v94}
                        switch (fi109 & -2147483648) {
                        case 387:
                            if (true || true) {
                                if true {
                                    var v127 bool = (2147483647 > fi109)
                                    v128 := GenS3{F0: 0}
                                    _ = v127
                                    _ = v128
                                }
                            } else {
                                var v129 uint = 0
                                fmt.Println(v87, v3, fi109)
                                _ = v129
                            }
                        case -9223372036854775808:
                            v3 += (v3 >> (((^(v1))) & 63))
                        case 127:
                            var v130 uint8 = 127
                            _ = v130
                        default:
                            var v131 float64 = float64(3.14159)
                            _ = v131
                        }
                        _ = v126
                    }
                    switch -9223372036854775808 {
                    case -24:
                        if true {
                            v132 := GenS5{F0: 1}
                            _ = v132
                        } else {
                            fmt.Println(v3, v123)
                        }
                    case 705:
                        if (true && true) {
                            v94 += (v94 - v94)
                            v87 = (((2147483648 != 2147483647) && (517 != 100)) && v87)
                            switch (fi108 ^ 2147483647) {
                            case -1:
                                var v133 int64 = 508
                                _ = v133
                            case -9223372036854775808:
                                if ((v1 & (v123 / ((v1) | 1))) > 2) {
                                    v134 := GenS1{F0: (v94 - (^(v94)))}
                                    switch -9223372036854775808 {
                                    case 146:
                                        v135 := GenS1{F0: v94}
                                        _ = v135
                                    case -585:
                                        if (-1 > 0) {
                                            switch "hello" {
                                            case "":
                                                switch fi109 {
                                                case 0:
                                                    var v136 int16 = 1000
                                                    _ = v136
                                                }
                                            case "\n\t":
                                                fmt.Println(v87, v87, v123)
                                            case "a":
                                                if true {
                                                    switch (fi109 - -833) {
                                                    case 42:
                                                        v3 = 516
                                                    case 9223372036854775807:
                                                        var v137 int = (fi109 & (fi108 ^ (fi108 + int(v123))))
                                                        _ = v137
                                                    case -9223372036854775808:
                                                        fmt.Println(fi108, v94, v94)
                                                    }
                                                    v138 := GenS5{F0: 65535}
                                                    _ = v138
                                                } else {
                                                    fmt.Println(v87, v94, v3)
                                                    v139 := GenS6{F0: 567, F1: 990}
                                                    var v140 string = "\n\t"
                                                    _ = v139
                                                    _ = v140
                                                }
                                            }
                                            if len(v2) > 0 {
                                                v2[0] = 0
                                            }
                                            fmt.Println(v123, v3, v1)
                                        } else {
                                            switch "" {
                                            case "foo bar":
                                                var v141 uint32 = 164
                                                _ = v141
                                            case "hello":
                                                if true {
                                                    fmt.Println(v87)
                                                    if len(v2) > 0 {
                                                        v2[0] = uint32(v3)
                                                    }
                                                } else {
                                                    v94 += v94
                                                    v2 = append(v2, 335)
                                                }
                                            case "a":
                                                var v142 bool = v87
                                                _ = v142
                                            }
                                            var v143 string = ("" + ((("foo bar" + "\n\t") + "a") + "foo bar"))
                                            var v144 int = (fi109 - (^(fi109)))
                                            _ = v143
                                            _ = v144
                                        }
                                    }
                                    _ = v134
                                } else {
                                    v2 = append(v2, 4294967295)
                                    v123 = (v1 & (^(v123)))
                                    fmt.Println(v3, v123, v123)
                                }
                            case 42:
                                switch fi109 {
                                case -793:
                                    var v145 int64 = 9223372036854775807
                                    _ = v145
                                case -539:
                                    var v146 int8 = 76
                                    _ = v146
                                }
                            }
                        } else {
                            var v147 uint32 = 65535
                            if true {
                                _ = len(v2)
                                if (true || (!v87)) {
                                    switch fi109 {
                                    case 128:
                                        fmt.Println(fi108)
                                    case 32768:
                                        _ = len(v2)
                                    }
                                    v94 = 314
                                }
                            } else {
                                v94 -= (v94 + v94)
                                var v148 uint = (v3 - (v1 | (v123 % (((v1 ^ (v3 + (v3 | (v1 / ((1) | 1)))))) | 1))))
                                var v149 uint64 = (v94 >> (((v148 % ((v1) | 1))) & 63))
                                _ = v148
                                _ = v149
                            }
                            _ = v147
                        }
                    }
                case "\n\t":
                    v3 = (v3 | 18446744073709551615)
                }
                var v150 int32 = -1
                _ = v123
                _ = v150
            }
            v3 = 0
            _ = fi109
        }
        if v87 {
            if v87 {
                var v151 int64 = -899
                switch (fi108 | (fi108 - (fi108 * 1))) {
                case 1:
                    switch 4294967295 {
                    case -2147483648:
                        switch fi108 {
                        case 4294967295:
                            v3 = (v1 * v3)
                        default:
                            fmt.Println(v1)
                        }
                    case 378:
                        v151 += (-(v151))
                    case 2147483647:
                        for fi152 := 0; fi152 < 17; fi152++ {
                            var v153 uint64 = v94
                            v87 = (!((4 <= 0) && (3.14159 <= float32((0.1 - ((1.5 - 0.0) + 3.14159))))))
                            _ = fi152
                            _ = v153
                        }
                    }
                case 568:
                    _ = len(v2)
                default:
                    fmt.Println(fi108)
                }
                _ = v151
            } else {
                v154 := GenS14{F0: (fi108 ^ 32768)}
                _ = v154
            }
            var v155 float64 = 3.14159
            _ = v155
        }
        _ = fi108
    }
    if v87 {
        fmt.Println(v87)
        v3 += 0
    }
    fmt.Println(v94)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v87:", v87)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v104
    _ = v2
    _ = v3
    _ = v87
    _ = v94
}
