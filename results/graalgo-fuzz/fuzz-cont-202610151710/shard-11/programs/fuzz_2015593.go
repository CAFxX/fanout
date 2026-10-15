package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 int8; F2 uint32 }
type GenS4 struct { F0 int8; F1 uint32 }
type GenS5 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int32; F1 uint8; F2 int }
type GenS8 struct { F0 uint; F1 int32; F2 int64 }
type GenS9 struct { F0 string; F1 uint16 }
type GenS10 struct { F0 bool; F1 uint16; F2 int16 }
type GenS11 struct { F0 int16; F1 uint8 }
type GenS12 struct { F0 uint64 }

func main() {
    v1 := make(map[string]int)
    var v2 string = ""
    var v3 int64 = -1
    v4 := []uint32{1, 1}
    var v5 int8 = -1
    var v6 float32 = (1e10 * float32(float64(3.14159)))
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 float32 = -2.25
        _ = fi7
        _ = v8
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch 127 {
    case 127:
        fmt.Println(v3, v6)
    case -864:
        v5 = (v5 ^ (v5 / ((v5) | 1)))
    case -848:
        v9 := GenS1{F0: 65535}
        _ = v9
    }
    var v10 float64 = 0.0
    for fi11 := 0; fi11 < 9; fi11++ {
        if ((32767 >= -32768) || true) {
            if false {
                var v12 uint8 = 127
                _ = v12
            } else {
                var v13 int32 = -456
                delete(v1, v2)
                _ = v13
            }
            for fi14 := 0; fi14 < 7; fi14++ {
                v6 -= ((v6 * float32((0.1 * 1.5))) + v6)
                var v15 uint8 = 127
                var v16 int = fi14
                _ = fi14
                _ = v15
                _ = v16
            }
            v3 += 9223372036854775807
        } else {
            v6 += v6
            var v17 int8 = (^(v5))
            _ = v17
        }
        v6 = (1e10 * float32(v10))
        switch fi11 {
        case 0:
            var v18 int8 = -128
            _ = v18
        case 2147483647:
            var v19 uint16 = 0
            _ = v19
        }
        _ = fi11
    }
    v6 = (1.5 * 0.1)
    if (v3 >= v3) {
        switch 4294967295 {
        case -314:
            v10 -= 0.0
        case 32767:
            if (v2 >= (v2 + v2)) {
                if true {
                    v3 = v3
                    v20 := GenS2{F0: (true || (!false))}
                    _ = v20
                } else {
                    var v21 uint = 255
                    var v22 int64 = (v3 - v3)
                    for fi23 := 0; fi23 < 17; fi23++ {
                        fmt.Println(v6)
                        _ = fi23
                    }
                    _ = v21
                    _ = v22
                }
                for fi24 := 0; fi24 < 4; fi24++ {
                    for fi25 := 0; fi25 < 13; fi25++ {
                        if ((^(v3)) < (v3 - v3)) {
                            switch fi25 {
                            case 0:
                                delete(v1, (v2 + (v2 + v2)))
                            case 789:
                                v3 = -922
                            default:
                                fmt.Println(v2, v5)
                            }
                            v26 := GenS2{F0: ((!true) || true)}
                            _ = len(v4)
                            _ = v26
                        }
                        var v27 float64 = float64(3.14159)
                        if (-1 <= -2147483648) {
                            var v28 uint16 = 0
                            if (2 >= 88) {
                                if false {
                                    if true {
                                        v4 = append(v4, 42)
                                        if true {
                                            var v29 bool = true
                                            if v29 {
                                                var v30 int8 = v5
                                                _ = v30
                                            } else {
                                                if (v29 || (1 == 42)) {
                                                    var v31 uint8 = 93
                                                    v32 := GenS2{F0: (!v29)}
                                                    v33 := GenS2{F0: v29}
                                                    _ = v31
                                                    _ = v32
                                                    _ = v33
                                                } else {
                                                    switch (fi24 << ((uint(v28)) & 63)) {
                                                    case 32768:
                                                        _ = len(v4)
                                                    case 0:
                                                        v34 := v1["hello"]
                                                        _ = v34
                                                    }
                                                }
                                            }
                                            _ = v29
                                        }
                                    }
                                }
                                v4 = append(v4, 4294967295)
                                v35 := GenS2{F0: true}
                                _ = v35
                            }
                            _ = len(v4)
                            _ = v28
                        }
                        if false {
                            if (false || (!true)) {
                                if true {
                                    var v36 int = fi24
                                    var v37 int = (v36 ^ (-(fi25)))
                                    _ = v36
                                    _ = v37
                                } else {
                                    v5 = v5
                                    v2 = v2
                                    if (true && true) {
                                        var v38 uint16 = 32767
                                        _ = v38
                                    } else {
                                        if (false && (!true)) {
                                            switch fi24 {
                                            case 1:
                                                v39 := GenS1{F0: 255}
                                                _ = v39
                                            case -2147483648:
                                                v4 = append(v4, 76)
                                            case 3:
                                                if false {
                                                    if (!((v5 ^ v5) > (v5 << ((2) & 63)))) {
                                                        v5 += 127
                                                        switch fi24 {
                                                        case -2147483648:
                                                            if false {
                                                                var v40 float64 = 0.1
                                                                v41 := GenS3{F0: "\n\t", F1: 0, F2: 0}
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                                var v42 int16 = 32767
                                                                var v43 float32 = (v6 - -2.25)
                                                                _ = v40
                                                                _ = v41
                                                                _ = v42
                                                                _ = v43
                                                            }
                                                        default:
                                                            if ((v5 << ((2) & 63)) > 100) {
                                                                v6 = v6
                                                            }
                                                            if false {
                                                                v44 := GenS1{F0: 2}
                                                                v5 = v5
                                                                _ = v44
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    v45 := GenS2{F0: (1 != 42)}
                                                    v4 = append(v4, 672)
                                                    v10 = (((3.14159 + (v27 * -2.25)) / 1.0) - ((float64((v6 * (3.14159 - (v6 / 1.0)))) * v10) / 3.14159))
                                                    _ = v45
                                                }
                                            }
                                            if false {
                                                if false {
                                                    var v46 int32 = 1
                                                    fmt.Println(v10, v27, v5)
                                                    _ = v46
                                                }
                                                v6 = float32((v27 + (v10 / v27)))
                                                delete(v1, v2)
                                                v5 -= -1
                                            } else {
                                                v47 := GenS1{F0: 413}
                                                v27 = 0.1
                                                var v48 float64 = (v10 - (3.14159 - v27))
                                                switch (((v2 + v2) + ((("a" + ((((v2 + ("foo bar" + "a")) + ("a" + v2)) + ((v2 + v2) + ("hello" + ""))) + "foo bar")) + (v2 + v2)) + v2)) + "") {
                                                case "a":
                                                    fmt.Println(v48, v48)
                                                case "hello":
                                                    fmt.Println(v5)
                                                default:
                                                    if (0.1 >= ((v48 - 1.5) + (v10 / (float64(1.5) / ((-2.25 / 1.0) * (v10 * 0.1)))))) {
                                                        v48 = ((((-2.25 * v48) / float64((0.0 + v6))) * 0.1) / v10)
                                                        var v49 int8 = (v5 % ((v5) | 1))
                                                        fmt.Println(v48, fi25)
                                                        _ = v49
                                                    }
                                                    var v50 string = v2
                                                    var v51 uint16 = 393
                                                    _ = v50
                                                    _ = v51
                                                }
                                                if (!true) {
                                                    fmt.Println(v2, v10)
                                                    if true {
                                                        var v52 uint16 = 877
                                                        if (fi24 != fi25) {
                                                            if true {
                                                                var v53 int = (fi24 << ((2) & 63))
                                                                var v54 float32 = ((0.1 * 1.5) + (0.1 / (3.14159 + 1e10)))
                                                                _ = v53
                                                                _ = v54
                                                            } else {
                                                                v55 := GenS4{F0: 127, F1: 689}
                                                                _ = v55
                                                            }
                                                        }
                                                        var v56 uint16 = (v52 / ((v52) | 1))
                                                        _ = v52
                                                        _ = v56
                                                    }
                                                }
                                                fmt.Println(fi25, v48, v10)
                                                _ = v47
                                                _ = v48
                                            }
                                        }
                                    }
                                }
                                var v57 float32 = (v6 - float32(v10))
                                switch (fi25 % (((fi24 * fi24)) | 1)) {
                                case 128:
                                    v1[v2] = 32767
                                case -831:
                                    fmt.Println(v27, fi24)
                                case 4294967295:
                                    if len(v4) > 0 {
                                        v4[0] = 65535
                                    }
                                default:
                                    v58 := GenS5{F0: -85, F1: 185, F2: 0}
                                    _ = v58
                                }
                                if (false && true) {
                                    _ = len(v1)
                                }
                                _ = v57
                            }
                            if (1 == 739) {
                                v2 = ("hello" + v2)
                                switch 42 {
                                case -9223372036854775808:
                                    var v59 int = (fi25 | 349)
                                    _ = v59
                                }
                                v6 = (1e10 + v6)
                            } else {
                                if false {
                                    if false {
                                        if len(v4) > 0 {
                                            v4[0] = 2
                                        }
                                        v60 := GenS1{F0: 42}
                                        _ = v60
                                    }
                                    fmt.Println(v10, v6)
                                } else {
                                    switch -9223372036854775808 {
                                    case -999:
                                        fmt.Println(v27, v10, v6)
                                    case 32768:
                                        var v61 int64 = (v3 >> ((327) & 63))
                                        fmt.Println(v27, v61, fi25)
                                        var v62 int64 = 64
                                        _ = v61
                                        _ = v62
                                    case 597:
                                        switch (fi25 - fi25) {
                                        case 2147483647:
                                            fmt.Println(v27, v3)
                                        case -529:
                                            if (v3 <= v3) {
                                                fmt.Println(v27)
                                                if (!(int16(v3) < -1)) {
                                                    delete(v1, "foo bar")
                                                    delete(v1, v2)
                                                    switch (fi25 % (((fi24 - -2147483648)) | 1)) {
                                                    case -2147483648:
                                                        v27 -= -2.25
                                                    case -211:
                                                        v63 := v1[v2]
                                                        _ = v63
                                                    }
                                                }
                                                if false {
                                                    var v64 uint64 = 0
                                                    _ = v64
                                                } else {
                                                    v65 := GenS1{F0: 4294967295}
                                                    _ = v65
                                                }
                                            }
                                        default:
                                            v66 := GenS1{F0: 2}
                                            _ = v66
                                        }
                                    default:
                                        if false {
                                            var v67 int64 = v3
                                            var v68 float32 = (v6 * 1e10)
                                            _ = len(v1)
                                            var v69 int8 = v5
                                            _ = v67
                                            _ = v68
                                            _ = v69
                                        }
                                    }
                                }
                                if false {
                                    var v70 uint16 = 1
                                    if false {
                                        v71 := GenS6{F0: 712}
                                        var v72 int32 = -2147483648
                                        var v73 int16 = -103
                                        v1[""] = -2147483648
                                        switch (fi25 - (fi24 & (^(fi25)))) {
                                        case 9223372036854775807:
                                            switch (fi24 | (fi25 + fi24)) {
                                            case 2147483647:
                                                v74 := GenS6{F0: (^(v70))}
                                                _ = v74
                                            case 976:
                                                var v75 int = (fi25 * (fi25 % ((-2147483648) | 1)))
                                                _ = v75
                                            case 0:
                                                v76 := GenS3{F0: "a", F1: (v5 << ((4294967295) & 63)), F2: 42}
                                                _ = v76
                                            default:
                                                v1[v2] = (-(fi25))
                                            }
                                        case 403:
                                            switch 127 {
                                            case 32768:
                                                var v77 int8 = 44
                                                _ = v77
                                            case 128:
                                                if false {
                                                    if true {
                                                        var v78 bool = false
                                                        var v79 bool = ("foo bar" <= v2)
                                                        _ = v78
                                                        _ = v79
                                                    }
                                                } else {
                                                    switch ("a" + v2) {
                                                    case "\n\t":
                                                        v80 := GenS6{F0: 32767}
                                                        _ = v80
                                                    case "":
                                                        _ = len(v1)
                                                    default:
                                                        v81 := GenS6{F0: 42}
                                                        _ = v81
                                                    }
                                                    fmt.Println(v2, v3)
                                                }
                                            case 1:
                                                var v82 string = v2
                                                _ = v82
                                            default:
                                                if true {
                                                    var v83 float32 = (-2.25 + v6)
                                                    v84 := GenS6{F0: (v70 % ((858) | 1))}
                                                    _ = v83
                                                    _ = v84
                                                } else {
                                                    if (!true) {
                                                        var v85 float64 = v10
                                                        _ = v85
                                                    } else {
                                                        var v86 uint16 = v70
                                                        v2 = v2
                                                        _ = v86
                                                    }
                                                    fmt.Println(fi24, v10)
                                                    v27 = float64(v6)
                                                }
                                            }
                                        case -920:
                                            var v87 uint8 = 128
                                            _ = v87
                                        }
                                        fmt.Println(v70, v72, v10)
                                        if ((v3 >> ((255) & 63)) >= v3) {
                                            if false {
                                                switch "foo bar" {
                                                case "a":
                                                    switch (v2 + v2) {
                                                    case "\n\t":
                                                        v6 = v6
                                                    }
                                                case "hello":
                                                    _ = len(v1)
                                                default:
                                                    var v88 uint = 2
                                                    _ = v88
                                                }
                                                var v89 float64 = 0.1
                                                _ = v89
                                            }
                                            v10 = ((v10 - v10) * -2.25)
                                            var v90 int64 = (v3 - (^(v3)))
                                            _ = v90
                                        }
                                        var v91 uint8 = 18
                                        _ = v71
                                        _ = v72
                                        _ = v73
                                        _ = v91
                                    }
                                    _ = v70
                                }
                                switch (fi24 >> ((665) & 63)) {
                                case 127:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                case 2147483647:
                                    v3 = (v3 & v3)
                                case 32767:
                                    var v92 int = (fi24 ^ fi24)
                                    _ = v92
                                }
                                v5 -= (v5 + 124)
                            }
                            v10 = float64(((((((1e10 + v6) + v6) - (v6 * (((float32(((v27 * v27) + v27)) / (((float32(v27) / v6) - float32(0.1)) / (v6 * float32(v27)))) + (v6 - v6)) + (v6 / 3.14159)))) - 1.5) - v6) * v6))
                            var v93 string = (v2 + v2)
                            v27 -= (v27 / v10)
                            _ = v93
                        }
                        _ = fi25
                        _ = v27
                    }
                    switch -1 {
                    case -9223372036854775808:
                        v2 = "\n\t"
                    default:
                        var v94 uint32 = 687
                        _ = v94
                    }
                    var v95 int32 = 558
                    _ = fi24
                    _ = v95
                }
                v96 := GenS7{F0: -1, F1: 29, F2: 0}
                _ = v96
            }
        case 533:
            if ((-2.25 - (float32(0.0) * 1.5)) != v6) {
                for fi97 := 0; fi97 < 3; fi97++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v98 uint64 = 0
                    _ = fi97
                    _ = v98
                }
            } else {
                for fi99 := 0; fi99 < 6; fi99++ {
                    if true {
                        for fi100 := 0; fi100 < 20; fi100++ {
                            var v101 int8 = (v5 - (^(v5)))
                            var v102 string = (((v2 + v2) + v2) + v2)
                            _ = len(v1)
                            _ = fi100
                            _ = v101
                            _ = v102
                        }
                    }
                    var v103 int32 = -309
                    _ = fi99
                    _ = v103
                }
                var v104 int = -543
                _ = v104
            }
        default:
            for fi105 := 0; fi105 < 3; fi105++ {
                v2 = ""
                v2 = v2
                _ = fi105
            }
        }
    }
    if (v6 < 1e10) {
        var v106 uint8 = 219
        v107 := GenS4{F0: (v5 % ((127) | 1)), F1: 65535}
        var v108 float64 = (((v10 - ((v10 + 1e10) / -2.25)) * ((float64(float32(v10)) * v10) - float64((1.5 + v6)))) * v10)
        _ = v106
        _ = v107
        _ = v108
    } else {
        delete(v1, v2)
    }
    var v109 float64 = float64(v6)
    v109 -= float64(((v6 / float32(3.14159)) - ((((0.1 - (1.5 / float32(1e10))) * (float32(v109) + v6)) + v6) * 3.14159)))
    switch 459 {
    case 127:
        var v110 uint16 = 2
        _ = v110
    case -1:
        var v111 int16 = 788
        _ = v111
    case -452:
        v112 := v1["foo bar"]
        _ = v112
    }
    switch ("" + (v2 + v2)) {
    case "foo bar":
        var v113 uint = 205
        _ = v113
    default:
        v109 = 0.1
    }
    for fi114 := 0; fi114 < 11; fi114++ {
        v109 += v10
        for fi115 := 0; fi115 < 20; fi115++ {
            v10 = float64(((v6 + v6) + -2.25))
            switch v2 {
            case "foo bar":
                fmt.Println(v2)
            case "\n\t":
                if (0 < -603) {
                    v116 := GenS8{F0: 1, F1: 42, F2: (v3 | (v3 / ((v3) | 1)))}
                    v2 = (("\n\t" + "") + (v2 + ("a" + v2)))
                    v117 := GenS7{F0: -526, F1: 152, F2: 42}
                    _ = v116
                    _ = v117
                } else {
                    fmt.Println(v6)
                    v118 := GenS3{F0: "\n\t", F1: v5, F2: 0}
                    switch "hello" {
                    case "":
                        var v119 int16 = -975
                        _ = v119
                    case "a":
                        fmt.Println(v3)
                    }
                    var v120 string = "hello"
                    _ = v118
                    _ = v120
                }
            default:
                if false {
                    v6 = v6
                    if true {
                        var v121 int32 = 42
                        v10 = (v109 / ((1e10 + v10) + v10))
                        v3 = int64(fi114)
                        var v122 bool = (!true)
                        switch v2 {
                        case "a":
                            v123 := GenS9{F0: (v2 + (v2 + ((v2 + "foo bar") + v2))), F1: 32767}
                            _ = v123
                        case "\n\t":
                            v2 = v2
                        case "hello":
                            if false {
                                var v124 uint64 = 2
                                switch (v2 + ("" + ("\n\t" + v2))) {
                                case "foo bar":
                                    v2 = ("hello" + "foo bar")
                                case "":
                                    if v122 {
                                        v125 := GenS2{F0: v122}
                                        v126 := GenS9{F0: (("hello" + v2) + v2), F1: 32767}
                                        _ = v125
                                        _ = v126
                                    } else {
                                        if (v124 <= v124) {
                                            v124 = (v124 * v124)
                                            var v127 int32 = v121
                                            _ = v127
                                        }
                                    }
                                case "a":
                                    var v128 uint32 = 374
                                    _ = v128
                                }
                                var v129 int16 = 0
                                _ = v124
                                _ = v129
                            }
                        }
                        _ = v121
                        _ = v122
                    } else {
                        v10 = 1e10
                        switch v2 {
                        case "hello":
                            v6 -= float32(v10)
                            if (42 == 2147483647) {
                                if true {
                                    fmt.Println(v109, v109)
                                    if true {
                                        if false {
                                            v1[v2] = (^(fi115))
                                        }
                                        var v130 int16 = 42
                                        _ = v130
                                    } else {
                                        fmt.Println(v3, v109)
                                        if true {
                                            v2 = v2
                                            v4 = append(v4, 311)
                                        }
                                        if (649 < 2) {
                                            var v131 uint16 = 255
                                            _ = v131
                                        }
                                    }
                                    switch "\n\t" {
                                    case "":
                                        switch v2 {
                                        case "a":
                                            v132 := GenS10{F0: true, F1: 721, F2: 42}
                                            _ = v132
                                        default:
                                            if ("hello" <= "\n\t") {
                                                v1[v2] = (^(fi115))
                                                if (v3 <= -383) {
                                                    if ((0.0 < v109) && ((fi115 - fi114) == (fi114 + fi114))) {
                                                        switch -9223372036854775808 {
                                                        case -344:
                                                            v109 = v10
                                                        }
                                                    } else {
                                                        switch -114 {
                                                        case -1:
                                                            v133 := GenS11{F0: 32767, F1: 255}
                                                            _ = v133
                                                        case 43:
                                                            var v134 uint32 = 225
                                                            _ = v134
                                                        default:
                                                            v135 := GenS12{F0: 622}
                                                            _ = v135
                                                        }
                                                    }
                                                }
                                                if (v109 < -2.25) {
                                                    if (4294967295 != 65535) {
                                                        v136 := GenS11{F0: 709, F1: 8}
                                                        _ = v136
                                                    } else {
                                                        var v137 int16 = 1000
                                                        _ = v137
                                                    }
                                                } else {
                                                    v138 := GenS7{F0: 2147483647, F1: 96, F2: (fi115 % ((32767) | 1))}
                                                    var v139 bool = (1 == -871)
                                                    var v140 uint16 = 984
                                                    _ = v138
                                                    _ = v139
                                                    _ = v140
                                                }
                                            }
                                        }
                                    }
                                } else {
                                    v141 := v1[((v2 + "hello") + v2)]
                                    fmt.Println(v6, v3, fi114)
                                    var v142 uint16 = 42
                                    _ = v141
                                    _ = v142
                                }
                                var v143 float64 = v109
                                _ = v143
                            } else {
                                if len(v4) > 0 {
                                    v4[0] = 1
                                }
                            }
                        }
                    }
                    _ = len(v4)
                }
            }
            _ = fi115
        }
        v144 := GenS2{F0: ((60 >= 128) || (!true))}
        _ = fi114
        _ = v144
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v109:", v109)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v109
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
