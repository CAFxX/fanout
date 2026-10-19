package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 uint }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint; F1 bool; F2 uint }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint64; F1 uint32; F2 string }
type GenS9 struct { F0 bool; F1 uint16; F2 int32 }
type GenS10 struct { F0 int32; F1 uint8 }
type GenS11 struct { F0 string; F1 int16; F2 int64 }
type GenS12 struct { F0 int8; F1 int16 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 uint; F1 bool }
type GenS15 struct { F0 int32; F1 bool; F2 int8 }
type GenS16 struct { F0 int16; F1 int16 }
type GenS17 struct { F0 uint64; F1 uint32; F2 bool }
type GenS18 struct { F0 uint16; F1 int32; F2 uint16 }
type GenS19 struct { F0 int8; F1 uint64 }
type GenS20 struct { F0 int32; F1 string; F2 uint16 }
type GenS21 struct { F0 int8 }

func main() {
    var v1 float64 = 3.14159
    var v2 int64 = -9223372036854775808
    var v3 int8 = 42
    var v4 float32 = 0.0
    var v5 int = -343
    var v6 int64 = -462
    v7 := GenS1{F0: true}
    switch "hello" {
    case "":
        var v8 uint8 = 255
        _ = v8
    default:
        v9 := []int{801, 1, (v5 % ((v5) | 1)), 4294967295}
        _ = v9
    }
    v5 = v5
    var v10 int = v5
    v11 := GenS1{F0: (!true)}
    fmt.Println(v2)
    var v12 uint8 = 56
    if false {
        if ((v10 % (((-(v10))) | 1)) > v5) {
            v13 := GenS1{F0: false}
            _ = v13
        } else {
            for fi14 := 0; fi14 < 2; fi14++ {
                fmt.Println(v3)
                var v15 int8 = (v3 | (v3 | 0))
                if false {
                    switch ("hello" + "") {
                    case "a":
                        v5 += 42
                    default:
                        v16 := GenS1{F0: true}
                        _ = v16
                    }
                    v17 := GenS2{F0: v12}
                    _ = v17
                } else {
                    v18 := []int{}
                    v19 := GenS2{F0: (v12 & (v12 * (v12 + (v12 / (((^(v12))) | 1)))))}
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                    _ = v18
                    _ = v19
                }
                _ = fi14
                _ = v15
            }
            var v20 int16 = 1
            if false {
                if true {
                    v21 := GenS3{F0: (v12 >> ((1) & 63))}
                    _ = v21
                }
                v22 := GenS4{F0: 2147483648, F1: 42}
                _ = v22
            } else {
                v23 := GenS2{F0: v12}
                _ = v23
            }
            _ = v20
        }
    } else {
        v24 := GenS2{F0: 82}
        _ = v24
    }
    v25 := []int{}
    switch "\n\t" {
    case "":
        for fi26 := 0; fi26 < 6; fi26++ {
            v12 += 42
            var v27 uint64 = 4294967295
            _ = fi26
            _ = v27
        }
    case "a":
        for fi28 := 0; fi28 < 3; fi28++ {
            v10 = 9223372036854775807
            var v29 int8 = v3
            v30 := GenS1{F0: false}
            _ = fi28
            _ = v29
            _ = v30
        }
    }
    for fi31 := 0; fi31 < 18; fi31++ {
        if (0 >= -1) {
            var v32 float32 = (v4 + ((v4 + 3.14159) + 0.0))
            for fi33 := 0; fi33 < 17; fi33++ {
                v2 = v6
                if false {
                    if (2147483648 >= 2147483648) {
                        v6 = v6
                        var v34 uint64 = 0
                        var v35 uint = 65535
                        _ = v34
                        _ = v35
                    } else {
                        var v36 uint = 339
                        var v37 uint64 = 1
                        _ = v36
                        _ = v37
                    }
                    v4 += ((float32(1.5) - v32) * float32(3.14159))
                    v38 := GenS5{F0: 246}
                    _ = v38
                }
                v39 := GenS1{F0: false}
                if ((v5 / ((127) | 1)) < (fi33 ^ 4294967295)) {
                    var v40 uint8 = (^(v12))
                    v41 := GenS6{F0: 1, F1: (0 < 0), F2: 118}
                    var v42 uint32 = 0
                    _ = v40
                    _ = v41
                    _ = v42
                } else {
                    var v43 bool = (4294967295 != 42)
                    var v44 int64 = v6
                    v10 = 505
                    _ = v43
                    _ = v44
                }
                switch fi31 {
                case -9223372036854775808:
                    var v45 uint16 = 1
                    _ = v45
                case 365:
                    if (!(!true)) {
                        switch "foo bar" {
                        case "hello":
                            var v46 bool = true
                            _ = v46
                        case "":
                            switch "" {
                            case "":
                                fmt.Println(v12, v5)
                            default:
                                v4 -= (0.1 + (3.14159 - float32((0.1 * 1e10))))
                                v47 := GenS7{F0: 1}
                                var v48 uint16 = 830
                                fmt.Println(v2)
                                var v49 int32 = 0
                                _ = v47
                                _ = v48
                                _ = v49
                            }
                        default:
                            v50 := GenS8{F0: 4294967295, F1: 135, F2: ("a" + "\n\t")}
                            _ = v50
                        }
                        if (42 == -1) {
                            var v51 int8 = v3
                            _ = v51
                        } else {
                            if ((false && (42 != -1)) && true) {
                                v1 -= 1.5
                                v1 += v1
                            } else {
                                v52 := GenS4{F0: 4294967295, F1: 65535}
                                v32 = ((v4 - float32(((v1 + (float64(-2.25) / (0.1 / v1))) / 1.5))) * (v4 - v32))
                                _ = v52
                            }
                        }
                        v10 -= fi31
                    } else {
                        fmt.Println(v2, v32)
                        if true {
                            var v53 int16 = -455
                            _ = len(v25)
                            v54 := GenS9{F0: true, F1: 1, F2: 42}
                            _ = v53
                            _ = v54
                        } else {
                            switch -9223372036854775808 {
                            case 42:
                                if (!false) {
                                    v55 := GenS9{F0: (42 < -2147483648), F1: 255, F2: -1}
                                    _ = v55
                                } else {
                                    var v56 int64 = v2
                                    _ = v56
                                }
                            default:
                                var v57 int8 = v3
                                _ = v57
                            }
                            v58 := GenS10{F0: -1, F1: v12}
                            if (32767 >= 1) {
                                fmt.Println(fi31, fi31, v4)
                                var v59 int = (v10 | fi33)
                                if (0 < 1) {
                                    var v60 int8 = 8
                                    _ = v60
                                }
                                _ = v59
                            }
                            if len(v25) >= 2 {
                                v25 = v25[1:len(v25)]
                            }
                            _ = v58
                        }
                        if len(v25) >= 2 {
                            v25 = v25[1:len(v25)]
                        }
                    }
                }
                _ = fi33
                _ = v39
            }
            switch (v10 << ((42) & 63)) {
            case 339:
                for fi61 := 0; fi61 < 12; fi61++ {
                    var v62 uint64 = 0
                    v1 += 0.1
                    v6 = -9223372036854775808
                    if true {
                        v63 := GenS7{F0: v62}
                        if len(v25) > 0 {
                            v25[0] = int(v62)
                        }
                        if ((v12 ^ (^(v12))) < (^(v12))) {
                            v64 := GenS1{F0: true}
                            v4 -= v4
                            _ = v64
                        }
                        switch v10 {
                        case 1:
                            switch v10 {
                            case 81:
                                v5 -= fi61
                            default:
                                var v65 int16 = 42
                                _ = v65
                            }
                        case 32768:
                            fmt.Println(fi61, v12)
                        case -888:
                            var v66 int64 = 0
                            _ = v66
                        }
                        _ = v63
                    } else {
                        var v67 int = 32767
                        switch (v67 >> ((42) & 63)) {
                        case 128:
                            fmt.Println(v62, v32)
                        default:
                            v25 = append(v25, v10)
                        }
                        _ = v67
                    }
                    _ = fi61
                    _ = v62
                }
            case 42:
                var v68 float64 = 0.0
                _ = v68
            default:
                for fi69 := 0; fi69 < 11; fi69++ {
                    v25 = append(v25, v5)
                    if ("foo bar" < "\n\t") {
                        if true {
                            var v70 float64 = v1
                            if len(v25) > 0 {
                                v25[0] = v5
                            }
                            v5 += (v5 >> ((411) & 63))
                            _ = v70
                        }
                    } else {
                        var v71 float64 = (v1 + (v1 / (float64((v32 - v4)) / v1)))
                        if (false || (((!(0 <= 986)) || (39 != (^(v12)))) || true)) {
                            v25 = append(v25, int(v2))
                            v72 := GenS11{F0: "a", F1: 42, F2: (v6 ^ v6)}
                            if len(v25) > 0 {
                                v25[0] = (fi69 >> ((42) & 63))
                            }
                            _ = v72
                        } else {
                            if (1 < 1) {
                                fmt.Println(fi31)
                                v6 = (v6 * (-(v6)))
                                v6 += v2
                            } else {
                                v2 = 0
                                var v73 int32 = 2147483647
                                switch int(v3) {
                                case -342:
                                    fmt.Println(v73, v32)
                                case 4294967295:
                                    if true {
                                        fmt.Println(v4, v3, v12)
                                        v6 += (v2 | (v2 % ((v6) | 1)))
                                        var v74 int64 = v6
                                        _ = v74
                                    }
                                case -47:
                                    v4 += (v4 * -2.25)
                                }
                                _ = v73
                            }
                        }
                        _ = v71
                    }
                    _ = fi69
                }
            }
            _ = v32
        }
        v2 = v2
        _ = fi31
    }
    for fi75 := 0; fi75 < 19; fi75++ {
        v4 += v4
        if true {
            if true {
                var v76 int = v5
                switch v76 {
                case 0:
                    if len(v25) >= 2 {
                        v25 = v25[1:len(v25)]
                    }
                default:
                    for fi77 := 0; fi77 < 1; fi77++ {
                        v2 = v6
                        switch "foo bar" {
                        case "":
                            var v78 int16 = 42
                            _ = v78
                        case "foo bar":
                            var v79 uint8 = (v12 * v12)
                            _ = v79
                        }
                        fmt.Println(v4, v76)
                        _ = fi77
                    }
                }
                _ = v76
            } else {
                _ = len(v25)
                switch v5 {
                case 1:
                    var v80 int16 = 472
                    _ = v80
                case 707:
                    v81 := GenS12{F0: v3, F1: 0}
                    _ = v81
                case 0:
                    var v82 int = (v10 / ((-614) | 1))
                    _ = v82
                default:
                    var v83 uint32 = 2147483648
                    _ = v83
                }
                if (9223372036854775807 != 1) {
                    var v84 uint = 18446744073709551615
                    var v85 uint8 = 121
                    switch "foo bar" {
                    case "a":
                        v4 = (v4 * (0.0 - 1.5))
                    case "foo bar":
                        switch (v5 >> ((v84) & 63)) {
                        case 42:
                            for fi86 := 0; fi86 < 8; fi86++ {
                                if true {
                                    if (-2147483648 > 313) {
                                        switch "hello" {
                                        case "a":
                                            v87 := GenS8{F0: 736, F1: 933, F2: "a"}
                                            _ = v87
                                        case "hello":
                                            v88 := GenS8{F0: 272, F1: 44, F2: "foo bar"}
                                            _ = v88
                                        case "\n\t":
                                            switch ("hello" + "hello") {
                                            case "foo bar":
                                                _ = len(v25)
                                                switch int(v3) {
                                                case 747:
                                                    v6 = v6
                                                case -197:
                                                    if ((false || (uint(v12) <= 1)) || (!true)) {
                                                        if true {
                                                            if len(v25) >= 2 {
                                                                v25 = v25[1:len(v25)]
                                                            }
                                                            if true {
                                                                v12 -= v85
                                                                if false {
                                                                    var v89 bool = (!false)
                                                                    if v89 {
                                                                        v25 = append(v25, (fi75 * v10))
                                                                    } else {
                                                                        switch fi75 {
                                                                        case -2147483648:
                                                                            if (true || (v84 >= 65535)) {
                                                                                var v90 uint8 = 2
                                                                                if len(v25) > 0 {
                                                                                    v25[0] = fi86
                                                                                }
                                                                                _ = v90
                                                                            } else {
                                                                                if false {
                                                                                    var v91 int16 = int16(v3)
                                                                                    var v92 int64 = v6
                                                                                    _ = v91
                                                                                    _ = v92
                                                                                } else {
                                                                                    v3 = v3
                                                                                    v4 = ((((float32(v1) * (1e10 / 1.5)) - (v4 / v4)) * -2.25) * v4)
                                                                                    if (!(-2147483648 == -319)) {
                                                                                        var v93 uint8 = (v85 / ((v85) | 1))
                                                                                        v85 = v93
                                                                                        fmt.Println(v10, v1)
                                                                                        switch (^(v10)) {
                                                                                        case 9223372036854775807:
                                                                                            if (v1 > v1) {
                                                                                                v1 -= v1
                                                                                                var v94 uint = (v84 % ((v84) | 1))
                                                                                                v95 := GenS7{F0: 4294967295}
                                                                                                _ = v94
                                                                                                _ = v95
                                                                                            } else {
                                                                                                v25 = append(v25, (v5 | v5))
                                                                                                fmt.Println(v6, fi75, v84)
                                                                                            }
                                                                                        case 1:
                                                                                            fmt.Println(v10)
                                                                                        }
                                                                                        var v96 string = ""
                                                                                        _ = v93
                                                                                        _ = v96
                                                                                    }
                                                                                }
                                                                                fmt.Println(fi86, v85)
                                                                                var v97 float64 = (0.1 - v1)
                                                                                v98 := GenS2{F0: (v12 * v12)}
                                                                                _ = v97
                                                                                _ = v98
                                                                            }
                                                                        case -887:
                                                                            if len(v25) >= 2 {
                                                                                v25 = v25[1:len(v25)]
                                                                            }
                                                                        case 0:
                                                                            fmt.Println(fi75, v5, v3)
                                                                        }
                                                                        var v99 float64 = v1
                                                                        v12 -= (v12 / (((v85 / ((v12) | 1))) | 1))
                                                                        if true {
                                                                            switch (fi86 * v10) {
                                                                            case 0:
                                                                                _ = len(v25)
                                                                            default:
                                                                                _ = len(v25)
                                                                            }
                                                                            var v100 uint32 = 1
                                                                            v6 -= 441
                                                                            _ = v100
                                                                        }
                                                                        var v101 uint64 = 4294967295
                                                                        v5 = (v10 + -1)
                                                                        _ = v101
                                                                        _ = v99
                                                                    }
                                                                    _ = v89
                                                                } else {
                                                                    v10 += v5
                                                                    var v102 uint8 = (v85 % ((127) | 1))
                                                                    _ = v102
                                                                }
                                                                var v103 int8 = 0
                                                                _ = v103
                                                            } else {
                                                                var v104 uint64 = 0
                                                                fmt.Println(v1, v85)
                                                                _ = v104
                                                            }
                                                            v5 += -668
                                                        } else {
                                                            fmt.Println(v2, v5)
                                                            v25 = append(v25, v10)
                                                            _ = len(v25)
                                                            var v105 uint8 = 42
                                                            if false {
                                                                if false {
                                                                    v106 := GenS1{F0: (!true)}
                                                                    _ = v106
                                                                }
                                                            } else {
                                                                fmt.Println(v6, v84)
                                                                var v107 int32 = -1
                                                                _ = v107
                                                            }
                                                            var v108 int16 = 231
                                                            _ = v105
                                                            _ = v108
                                                        }
                                                        if ((((v4 * (v4 / v4)) + v4) - v4) < (float32((1.5 + 1.5)) - v4)) {
                                                            if false {
                                                                switch fi75 {
                                                                case 32:
                                                                    v3 = v3
                                                                    switch (v5 / ((128) | 1)) {
                                                                    case 4294967295:
                                                                        var v109 bool = false
                                                                        _ = v109
                                                                    case 618:
                                                                        v110 := GenS5{F0: 964}
                                                                        _ = v110
                                                                    case -1:
                                                                        _ = len(v25)
                                                                    default:
                                                                        v111 := GenS13{F0: v6}
                                                                        _ = v111
                                                                    }
                                                                case 436:
                                                                    v112 := GenS13{F0: int64(v84)}
                                                                    _ = v112
                                                                }
                                                                v2 -= (v2 & 9223372036854775807)
                                                                var v113 uint32 = 441
                                                                _ = v113
                                                            }
                                                            var v114 uint32 = 65535
                                                            _ = v114
                                                        } else {
                                                            var v115 bool = true
                                                            switch "hello" {
                                                            case "hello":
                                                                var v116 float32 = v4
                                                                _ = v116
                                                            case "a":
                                                                v117 := GenS1{F0: v115}
                                                                v10 += 2147483647
                                                                var v118 int64 = v2
                                                                _ = v117
                                                                _ = v118
                                                            default:
                                                                v10 += (v5 * 1)
                                                            }
                                                            v10 = (v10 | -9223372036854775808)
                                                            _ = v115
                                                        }
                                                    }
                                                default:
                                                    if true {
                                                        var v119 float64 = v1
                                                        switch fi86 {
                                                        case -382:
                                                            fmt.Println(v119)
                                                            if len(v25) >= 2 {
                                                                v25 = v25[1:len(v25)]
                                                            }
                                                        }
                                                        var v120 uint16 = 895
                                                        _ = v119
                                                        _ = v120
                                                    } else {
                                                        _ = len(v25)
                                                    }
                                                }
                                                var v121 uint16 = 255
                                                var v122 int64 = v6
                                                if false {
                                                    v123 := GenS3{F0: v12}
                                                    switch (("hello" + "a") + "") {
                                                    case "foo bar":
                                                        fmt.Println(v5, v84, fi75)
                                                    case "\n\t":
                                                        switch "" {
                                                        case "hello":
                                                            switch v5 {
                                                            case -188:
                                                                var v124 uint8 = uint8(v84)
                                                                _ = v124
                                                            }
                                                        case "":
                                                            var v125 int32 = 1
                                                            var v126 uint = v84
                                                            v125 -= v125
                                                            v3 = 42
                                                            _ = v125
                                                            _ = v126
                                                        case "a":
                                                            var v127 string = "\n\t"
                                                            var v128 string = ("" + "a")
                                                            v5 -= fi75
                                                            if false {
                                                                if ((true && (-1 != -2147483648)) || ((0 <= 42) && (!true))) {
                                                                    var v129 bool = false
                                                                    var v130 int16 = 1000
                                                                    _ = v129
                                                                    _ = v130
                                                                }
                                                                if false {
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        var v131 bool = true
                                                                        _ = v131
                                                                    case "a":
                                                                        var v132 uint16 = 42
                                                                        _ = v132
                                                                    default:
                                                                        if false {
                                                                            if len(v25) >= 2 {
                                                                                v25 = v25[1:len(v25)]
                                                                            }
                                                                            fmt.Println(v2, v5)
                                                                            var v133 uint = 65535
                                                                            var v134 float64 = (v1 * v1)
                                                                            _ = v133
                                                                            _ = v134
                                                                        } else {
                                                                            switch fi86 {
                                                                            case -916:
                                                                                v5 += v10
                                                                            case -9223372036854775808:
                                                                                var v135 uint = (v84 | (v84 + 78))
                                                                                _ = v135
                                                                            case 757:
                                                                                v136 := GenS14{F0: 963, F1: (42 >= -207)}
                                                                                _ = v136
                                                                            }
                                                                            switch 42 {
                                                                            case 4294967295:
                                                                                v122 = 902
                                                                                var v137 int8 = (v3 + v3)
                                                                                v121 -= v121
                                                                                _ = v137
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                switch (v5 >> ((uint(v121)) & 63)) {
                                                                case 0:
                                                                    v121 += (^(v121))
                                                                default:
                                                                    fmt.Println(v5, v2, v2)
                                                                    v10 = (fi75 | (v5 + fi75))
                                                                }
                                                                var v138 float64 = float64(v4)
                                                                v139 := GenS8{F0: 1, F1: 1, F2: v127}
                                                                var v140 uint16 = 1
                                                                var v141 int = -393
                                                                v85 -= (v85 - (v12 ^ 79))
                                                                v127 = v128
                                                                _ = v138
                                                                _ = v139
                                                                _ = v140
                                                                _ = v141
                                                            }
                                                            _ = v127
                                                            _ = v128
                                                        default:
                                                            fmt.Println(v1)
                                                        }
                                                    }
                                                    fmt.Println(fi75)
                                                    var v142 bool = true
                                                    v122 = 0
                                                    if (-2.25 == (float32(((-2.25 / (1.5 + 0.0)) - v1)) - 3.14159)) {
                                                        v143 := GenS9{F0: (((!v142) || (0.1 >= v1)) || v142), F1: v121, F2: 866}
                                                        v2 -= (v2 | (v2 * (^(v2))))
                                                        _ = v143
                                                    } else {
                                                        v5 -= (^(fi75))
                                                        var v144 int64 = v122
                                                        var v145 int = v5
                                                        _ = v144
                                                        _ = v145
                                                    }
                                                    if len(v25) >= 2 {
                                                        v25 = v25[1:len(v25)]
                                                    }
                                                    v10 = fi86
                                                    _ = v123
                                                    _ = v142
                                                } else {
                                                    v146 := GenS15{F0: 42, F1: (1 >= 1), F2: int8(v84)}
                                                    _ = v146
                                                }
                                                if false {
                                                    var v147 uint64 = 0
                                                    fmt.Println(fi86)
                                                    var v148 int64 = -1
                                                    _ = v147
                                                    _ = v148
                                                } else {
                                                    if len(v25) >= 2 {
                                                        v25 = v25[1:len(v25)]
                                                    }
                                                }
                                                _ = v121
                                                _ = v122
                                            case "a":
                                                switch 230 {
                                                case 1:
                                                    v85 = 2
                                                    var v149 uint16 = 149
                                                    _ = v149
                                                case 32768:
                                                    if ("" >= "foo bar") {
                                                        switch -9223372036854775808 {
                                                        case 2147483647:
                                                            if false {
                                                                v150 := GenS7{F0: 474}
                                                                var v151 uint16 = uint16(v84)
                                                                _ = v150
                                                                _ = v151
                                                            }
                                                        default:
                                                            v152 := GenS16{F0: -32768, F1: -32768}
                                                            fmt.Println(v10, v12, v84)
                                                            _ = v152
                                                        }
                                                    } else {
                                                        var v153 int64 = (v6 << ((v84) & 63))
                                                        var v154 uint32 = 42
                                                        var v155 uint = v84
                                                        fmt.Println(fi75, v85)
                                                        _ = v153
                                                        _ = v154
                                                        _ = v155
                                                    }
                                                case -515:
                                                    fmt.Println(v6, v3, v85)
                                                }
                                            case "":
                                                var v156 string = ("\n\t" + ("\n\t" + ("" + "hello")))
                                                _ = v156
                                            default:
                                                switch v5 {
                                                case 127:
                                                    v157 := GenS17{F0: 42, F1: uint32(v85), F2: false}
                                                    _ = v157
                                                default:
                                                    if true {
                                                        v3 -= v3
                                                    }
                                                }
                                            }
                                        default:
                                            var v158 uint32 = 105
                                            _ = v158
                                        }
                                    } else {
                                        if false {
                                            v2 = 655
                                        } else {
                                            if (v2 != (v6 - v2)) {
                                                switch (((("foo bar" + "") + ("\n\t" + ("hello" + ("" + "")))) + "hello") + "hello") {
                                                case "a":
                                                    if len(v25) >= 2 {
                                                        v25 = v25[1:len(v25)]
                                                    }
                                                }
                                                switch "hello" {
                                                case "\n\t":
                                                    if (-770 >= -2147483648) {
                                                        if (v12 >= 127) {
                                                            if false {
                                                                v4 += v4
                                                                var v159 int8 = 42
                                                                if ((v10 != -390) || true) {
                                                                    fmt.Println(v85, v12, v1)
                                                                    if len(v25) >= 2 {
                                                                        v25 = v25[1:len(v25)]
                                                                    }
                                                                } else {
                                                                    var v160 uint16 = 831
                                                                    var v161 float64 = v1
                                                                    _ = v160
                                                                    _ = v161
                                                                }
                                                                if (false || false) {
                                                                    var v162 uint16 = 255
                                                                    _ = v162
                                                                }
                                                                _ = v159
                                                            } else {
                                                                if len(v25) > 0 {
                                                                    v25[0] = 2147483647
                                                                }
                                                                switch "hello" {
                                                                case "\n\t":
                                                                    if true {
                                                                        v1 += (-2.25 + (float64(v4) * v1))
                                                                    }
                                                                case "a":
                                                                    v85 -= 176
                                                                default:
                                                                    v163 := GenS7{F0: 42}
                                                                    _ = v163
                                                                }
                                                                _ = len(v25)
                                                            }
                                                            fmt.Println(v85, fi86, v10)
                                                        } else {
                                                            if len(v25) > 0 {
                                                                v25[0] = v10
                                                            }
                                                            fmt.Println(v6, v84, v10)
                                                        }
                                                        if true {
                                                            v164 := GenS14{F0: (v84 ^ (v84 >> ((v84) & 63))), F1: false}
                                                            if ((v10 | v5) < (v5 / (((fi75 % (((fi86 + -2147483648)) | 1))) | 1))) {
                                                                v165 := GenS18{F0: 1000, F1: -735, F2: 281}
                                                                if len(v25) > 0 {
                                                                    v25[0] = (-(fi86))
                                                                }
                                                                fmt.Println(v84, v10, v84)
                                                                _ = v165
                                                            } else {
                                                                v6 -= 202
                                                            }
                                                            _ = v164
                                                        } else {
                                                            if (!("a" != ("a" + (("" + ((("foo bar" + "foo bar") + "foo bar") + (("\n\t" + "\n\t") + ((("hello" + "hello") + "foo bar") + "foo bar")))) + "\n\t")))) {
                                                                if true {
                                                                    var v166 bool = (true || (true || (128 <= (v12 ^ (^(v85))))))
                                                                    v10 = fi75
                                                                    v167 := GenS12{F0: 42, F1: 1000}
                                                                    _ = v166
                                                                    _ = v167
                                                                } else {
                                                                    if len(v25) > 0 {
                                                                        v25[0] = 32768
                                                                    }
                                                                    v10 -= (fi75 - (v5 ^ (v5 << (((v84 ^ (v84 / ((v84) | 1)))) & 63))))
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        var v168 uint32 = 511
                                                                        _ = v168
                                                                    case "foo bar":
                                                                        v3 = v3
                                                                    case "hello":
                                                                        if (true && (false || (7 < 140))) {
                                                                            if false {
                                                                                if (!false) {
                                                                                    var v169 uint8 = v85
                                                                                    _ = v169
                                                                                } else {
                                                                                    v3 = v3
                                                                                    if len(v25) > 0 {
                                                                                        v25[0] = v10
                                                                                    }
                                                                                    v25 = append(v25, (v10 | 4294967295))
                                                                                    v25 = append(v25, fi86)
                                                                                }
                                                                                if len(v25) >= 2 {
                                                                                    v25 = v25[1:len(v25)]
                                                                                }
                                                                                if true {
                                                                                    var v170 uint32 = 0
                                                                                    v171 := GenS1{F0: ((1 == 2147483647) && true)}
                                                                                    _ = v170
                                                                                    _ = v171
                                                                                } else {
                                                                                    v84 = v84
                                                                                    v12 -= (v85 >> ((v84) & 63))
                                                                                }
                                                                            } else {
                                                                                v10 += -1
                                                                                if (924 >= (^(v84))) {
                                                                                    var v172 float64 = (v1 + 0.0)
                                                                                    _ = v172
                                                                                }
                                                                            }
                                                                        }
                                                                    default:
                                                                        v173 := GenS19{F0: v3, F1: 611}
                                                                        _ = v173
                                                                    }
                                                                }
                                                            } else {
                                                                var v174 int64 = v2
                                                                _ = v174
                                                            }
                                                        }
                                                        v25 = append(v25, (fi86 % ((-970) | 1)))
                                                    }
                                                }
                                                v25 = append(v25, 9223372036854775807)
                                            } else {
                                                v25 = append(v25, fi86)
                                                v175 := GenS20{F0: 42, F1: "foo bar", F2: 56}
                                                v176 := GenS20{F0: -94, F1: "foo bar", F2: 0}
                                                _ = v175
                                                _ = v176
                                            }
                                            _ = len(v25)
                                            var v177 int16 = 32767
                                            var v178 float32 = v4
                                            v10 = (fi75 / ((v5) | 1))
                                            _ = v177
                                            _ = v178
                                        }
                                        var v179 uint32 = 327
                                        _ = v179
                                    }
                                }
                                if false {
                                    v180 := GenS17{F0: 42, F1: 4294967295, F2: true}
                                    var v181 int = (v10 / ((879) | 1))
                                    v85 = (^(v85))
                                    _ = v180
                                    _ = v181
                                }
                                _ = fi86
                            }
                        case 32768:
                            for fi182 := 0; fi182 < 15; fi182++ {
                                var v183 uint32 = 42
                                if len(v25) > 0 {
                                    v25[0] = fi75
                                }
                                _ = fi182
                                _ = v183
                            }
                        default:
                            if (true && true) {
                                v2 += -9223372036854775808
                            }
                        }
                    default:
                        var v184 int64 = 9223372036854775807
                        _ = v184
                    }
                    _ = v84
                    _ = v85
                } else {
                    _ = len(v25)
                }
            }
            if ("" == "\n\t") {
                for fi185 := 0; fi185 < 8; fi185++ {
                    if false {
                        if false {
                            var v186 int8 = 100
                            var v187 uint32 = 277
                            switch "\n\t" {
                            case "foo bar":
                                v12 -= v12
                            }
                            _ = v186
                            _ = v187
                        }
                    }
                    _ = fi185
                }
                switch ("a" + "foo bar") {
                case "\n\t":
                    if false {
                        for fi188 := 0; fi188 < 10; fi188++ {
                            v25 = append(v25, fi188)
                            var v189 uint8 = 255
                            var v190 uint = 149
                            _ = fi188
                            _ = v189
                            _ = v190
                        }
                        if ((1000 != 42) && (true && (v4 < ((v4 + v4) * v4)))) {
                            var v191 uint = 1
                            _ = v191
                        } else {
                            for fi192 := 0; fi192 < 12; fi192++ {
                                if false {
                                    v2 += (v6 * 0)
                                }
                                _ = fi192
                            }
                            v193 := GenS1{F0: ((v3 < (v3 >> ((65535) & 63))) && false)}
                            _ = v193
                        }
                        if len(v25) >= 2 {
                            v25 = v25[1:len(v25)]
                        }
                    }
                case "":
                    fmt.Println(v1, v5, v6)
                }
            } else {
                var v194 int = (fi75 ^ fi75)
                fmt.Println(v10, v12, fi75)
                _ = v194
            }
        }
        _ = fi75
    }
    _ = len(v25)
    if true {
        v195 := GenS12{F0: v3, F1: -1}
        _ = v195
    } else {
        v5 -= 32768
        v196 := GenS21{F0: v3}
        _ = v196
    }
    for fi197 := 0; fi197 < 8; fi197++ {
        for fi198 := 0; fi198 < 2; fi198++ {
            _ = len(v25)
            _ = fi198
        }
        switch (v10 / (((v10 + (v5 ^ (^(fi197))))) | 1)) {
        case 128:
            switch "\n\t" {
            case "":
                v199 := GenS17{F0: 18446744073709551615, F1: 445, F2: true}
                _ = v199
            }
        case -671:
            if (2147483647 < -2147483648) {
                v200 := GenS9{F0: false, F1: 2, F2: -683}
                _ = len(v25)
                var v201 string = ("" + "foo bar")
                _ = v200
                _ = v201
            } else {
                if false {
                    if false {
                        v202 := GenS3{F0: v12}
                        fmt.Println(v6)
                        for fi203 := 0; fi203 < 17; fi203++ {
                            v204 := GenS5{F0: 42}
                            var v205 float64 = v1
                            var v206 int8 = (v3 >> ((291) & 63))
                            v207 := GenS8{F0: 2, F1: 832, F2: "\n\t"}
                            var v208 int16 = 383
                            if len(v25) >= 2 {
                                v25 = v25[1:len(v25)]
                            }
                            _ = fi203
                            _ = v204
                            _ = v205
                            _ = v206
                            _ = v207
                            _ = v208
                        }
                        _ = v202
                    } else {
                        var v209 float32 = v4
                        v4 += 1.5
                        v25 = append(v25, fi197)
                        _ = v209
                    }
                }
                for fi210 := 0; fi210 < 11; fi210++ {
                    var v211 float32 = 1e10
                    var v212 int16 = 32767
                    var v213 float32 = 3.14159
                    _ = fi210
                    _ = v211
                    _ = v212
                    _ = v213
                }
                fmt.Println(v2, v6, v5)
            }
        }
        _ = fi197
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
