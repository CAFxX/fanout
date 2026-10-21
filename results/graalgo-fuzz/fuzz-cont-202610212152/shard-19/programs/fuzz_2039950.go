package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 string; F2 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 bool; F1 bool }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint; F1 string; F2 uint8 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint; F1 int64 }
type GenS12 struct { F0 uint8; F1 int }
type GenS13 struct { F0 uint; F1 string; F2 uint16 }
type GenS14 struct { F0 bool; F1 bool }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]int)
    var v3 bool = (false && true)
    var v4 int64 = 861
    v5 := make(map[string]bool)
    if v3 {
        var v6 float64 = -2.25
        if false {
            v6 += v6
        }
        _ = v6
    } else {
        var v7 int64 = v4
        var v8 uint16 = 42
        _ = v7
        _ = v8
    }
    var v9 int32 = 2147483647
    v10 := v1[804]
    v4 -= v4
    v11 := GenS1{F0: 18446744073709551615, F1: ("\n\t" + ""), F2: 472}
    v12 := GenS1{F0: 2, F1: "\n\t", F2: v9}
    var v13 uint16 = 65535
    v14 := GenS1{F0: 148, F1: "hello", F2: -1}
    switch ("a" + "hello") {
    case "a":
        if (1 != 4294967295) {
            if (v10 && (false && (v13 < (v13 | v13)))) {
                var v15 int = -9223372036854775808
                _ = v15
            } else {
                fmt.Println(v4, v4)
                if true {
                    v16 := v2[""]
                    var v17 uint64 = 42
                    _ = v16
                    _ = v17
                }
            }
            v10 = (0 >= 4294967295)
        } else {
            v13 = 42
            if (913 > 42) {
                v18 := v5["a"]
                if false {
                    var v19 uint16 = (v13 << ((2) & 63))
                    v4 = (v4 * v4)
                    _ = v19
                }
                _ = v18
            } else {
                if (!v10) {
                    if v10 {
                        switch 746 {
                        case -715:
                            var v20 int8 = -1
                            _ = v20
                        case -1:
                            v21 := GenS2{F0: 714}
                            _ = v21
                        case 1:
                            for fi22 := 0; fi22 < 3; fi22++ {
                                v1[1] = v10
                                var v23 float32 = (1.5 * 1.5)
                                _ = fi22
                                _ = v23
                            }
                        default:
                            v9 = -2147483648
                        }
                        delete(v1, 42)
                        v10 = (-128 != 0)
                    }
                } else {
                    v24 := GenS2{F0: 506}
                    v25 := []int{599, 2147483647, 0, 32768}
                    fmt.Println(v10, v4, v9)
                    _ = v24
                    _ = v25
                }
                if (269 <= 175) {
                    switch "\n\t" {
                    case "hello":
                        v3 = false
                    case "foo bar":
                        if false {
                            var v26 uint = 42
                            _ = v26
                        } else {
                            var v27 int = -415
                            fmt.Println(v13, v9, v4)
                            _ = v27
                        }
                    case "a":
                        var v28 int = 608
                        _ = v28
                    default:
                        for fi29 := 0; fi29 < 3; fi29++ {
                            if v10 {
                                v9 -= (v9 * (v9 * v9))
                                for fi30 := 0; fi30 < 15; fi30++ {
                                    v31 := GenS1{F0: 1, F1: "\n\t", F2: (v9 >> ((860) & 63))}
                                    v32 := v1[2]
                                    if ((v4 % ((-1) | 1)) > v4) {
                                        v32 = (true || v3)
                                        switch "a" {
                                        case "":
                                            var v33 uint = 255
                                            _ = v33
                                        case "\n\t":
                                            var v34 string = "\n\t"
                                            _ = v34
                                        case "foo bar":
                                            fmt.Println(fi29, v13)
                                        }
                                        var v35 uint64 = 4294967295
                                        _ = v35
                                    } else {
                                        switch 0 {
                                        case 9223372036854775807:
                                            if true {
                                                if v32 {
                                                    v36 := []int{(fi29 << ((uint(v13)) & 63)), (fi30 ^ 128), (fi29 % (((fi30 & (fi29 ^ fi30))) | 1)), fi30}
                                                    _ = v36
                                                } else {
                                                    switch "hello" {
                                                    case "":
                                                        switch "foo bar" {
                                                        case "a":
                                                            if (18446744073709551615 >= 1) {
                                                                v13 = 255
                                                                v37 := GenS3{F0: ((2147483648 == 150) && (1 > 1)), F1: ("foo bar" + "foo bar"), F2: "foo bar"}
                                                                var v38 int = int(v13)
                                                                _ = v37
                                                                _ = v38
                                                            } else {
                                                                var v39 float32 = (float32(1.5) - 0.1)
                                                                var v40 int8 = 1
                                                                if v32 {
                                                                    fmt.Println(fi29, v4)
                                                                    v9 = v9
                                                                    if (42 <= 42) {
                                                                        v41 := GenS2{F0: 2147483647}
                                                                        _ = v41
                                                                    }
                                                                    switch (fi30 ^ fi29) {
                                                                    case -285:
                                                                        switch fi29 {
                                                                        case 128:
                                                                            v1[2147483648] = (0 <= 2)
                                                                        case 9223372036854775807:
                                                                            var v42 string = ("hello" + "hello")
                                                                            _ = v42
                                                                        default:
                                                                            var v43 int16 = 42
                                                                            _ = v43
                                                                        }
                                                                    case -9223372036854775808:
                                                                        switch "foo bar" {
                                                                        case "\n\t":
                                                                            v39 = float32(0.1)
                                                                            var v44 int16 = -497
                                                                            if ((v40 >> ((493) & 63)) > int8(v4)) {
                                                                                var v45 uint64 = 577
                                                                                v46 := GenS4{F0: v44}
                                                                                _ = v45
                                                                                _ = v46
                                                                            }
                                                                            _ = v44
                                                                        case "hello":
                                                                            var v47 uint8 = 128
                                                                            _ = v47
                                                                        default:
                                                                            switch fi29 {
                                                                            case -919:
                                                                                fmt.Println(v13, fi30, v13)
                                                                            case 9223372036854775807:
                                                                                v48 := GenS4{F0: int16(v9)}
                                                                                _ = v48
                                                                            case 32767:
                                                                                if (!v3) {
                                                                                    var v49 uint32 = 2
                                                                                    v49 -= (v49 | (v49 & (v49 + 80)))
                                                                                    _ = v49
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    var v50 float64 = 1.5
                                                                    _ = v50
                                                                } else {
                                                                    if ("\n\t" >= "") {
                                                                        switch fi29 {
                                                                        case -1:
                                                                            switch "a" {
                                                                            case "":
                                                                                if (((fi30 - fi30) < fi30) && (v13 >= (v13 + v13))) {
                                                                                    v13 = v13
                                                                                    v51 := GenS1{F0: 0, F1: ("\n\t" + (("a" + "hello") + ("foo bar" + "hello"))), F2: (v9 & v9)}
                                                                                    _ = v51
                                                                                } else {
                                                                                    v52 := GenS5{F0: v10, F1: v3}
                                                                                    _ = v52
                                                                                }
                                                                                _ = len(v2)
                                                                                if (!(!((fi29 << ((4294967295) & 63)) != fi29))) {
                                                                                    v10 = (v39 == ((v39 + (((((v39 * (1.5 / (v39 * v39))) - (3.14159 * -2.25)) - ((0.0 / (0.0 + ((float32(float64((v39 + 3.14159))) / (((0.0 + v39) - 1.5) * (v39 + 1e10))) - float32(1e10)))) * (v39 + 0.0))) * ((float32(1e10) - 1e10) - v39)) - 1.5)) * (((((v39 * (float32((1e10 / (-2.25 / (0.0 - 1.5)))) - ((v39 - v39) * (v39 + float32(0.1))))) / v39) / (float32(((3.14159 / float64(v39)) - 1e10)) * (v39 / float32((0.0 + 0.1))))) / v39) - ((((-2.25 + (v39 / v39)) * (v39 * float32(float64((-2.25 / float32(float64((v39 + (0.1 * 1e10))))))))) - (v39 * ((0.1 * ((v39 * v39) + v39)) / 1.0))) * 0.1))))
                                                                                    var v53 uint16 = (v13 >> ((2) & 63))
                                                                                    var v54 uint16 = (^(v53))
                                                                                    _ = v53
                                                                                    _ = v54
                                                                                }
                                                                            case "a":
                                                                                var v55 int8 = (v40 << ((2) & 63))
                                                                                _ = v55
                                                                            }
                                                                            v39 = float32(-2.25)
                                                                            v13 += v13
                                                                        default:
                                                                            v10 = v3
                                                                        }
                                                                        v39 += (1.5 / (v39 + 0.1))
                                                                        v56 := []int{-815, 32768}
                                                                        var v57 int64 = (^(v4))
                                                                        _ = v56
                                                                        _ = v57
                                                                    } else {
                                                                        var v58 int8 = 1
                                                                        _ = v58
                                                                    }
                                                                }
                                                                _ = v39
                                                                _ = v40
                                                            }
                                                        }
                                                    case "\n\t":
                                                        v59 := []int{-1, (fi30 + fi30), 676, -1}
                                                        _ = v59
                                                    }
                                                }
                                            }
                                        case 645:
                                            fmt.Println(v9)
                                        default:
                                            var v60 uint8 = 255
                                            _ = v60
                                        }
                                        var v61 float64 = (-2.25 / 1.5)
                                        if false {
                                            var v62 int16 = 32767
                                            v63 := GenS6{F0: (v62 - (v62 - -32768))}
                                            v2[""] = (fi29 + 646)
                                            if true {
                                                v64 := []int{fi29, fi30}
                                                _ = v64
                                            } else {
                                                fmt.Println(v3, v32, v61)
                                                v65 := []int{}
                                                var v66 float32 = 1e10
                                                fmt.Println(v66, v66)
                                                _ = v65
                                                _ = v66
                                            }
                                            _ = v62
                                            _ = v63
                                        }
                                        _ = v61
                                    }
                                    _ = fi30
                                    _ = v31
                                    _ = v32
                                }
                            }
                            _ = fi29
                        }
                    }
                }
                v67 := []int{}
                _ = v67
            }
            v13 = (v13 ^ (v13 + 255))
        }
    case "\n\t":
        var v68 float32 = float32(-2.25)
        _ = v68
    case "":
        v13 -= 255
    default:
        if (29 < int16(v9)) {
            if v3 {
                if (v3 || true) {
                    fmt.Println(v10, v13)
                    if (true || ((v3 && (168 == 0)) && true)) {
                        var v69 float32 = (0.1 + 1.5)
                        _ = v69
                    } else {
                        switch -21 {
                        case 9223372036854775807:
                            var v70 uint = 1
                            _ = v70
                        case 639:
                            v9 = 0
                        }
                        var v71 uint = 135
                        for fi72 := 0; fi72 < 16; fi72++ {
                            v13 = (v13 >> ((18446744073709551615) & 63))
                            if (!(!((v71 * v71) >= (v71 / ((v71) | 1))))) {
                                for fi73 := 0; fi73 < 20; fi73++ {
                                    v3 = v3
                                    _ = fi73
                                }
                                var v74 string = "a"
                                _ = v74
                            }
                            _ = fi72
                        }
                        _ = v71
                    }
                } else {
                    switch "" {
                    case "foo bar":
                        if v3 {
                            var v75 uint8 = 255
                            if false {
                                var v76 float64 = 1e10
                                _ = v76
                            } else {
                                if v3 {
                                    var v77 float64 = float64(0.1)
                                    _ = v77
                                }
                            }
                            _ = v75
                        }
                    case "\n\t":
                        switch "foo bar" {
                        case "a":
                            for fi78 := 0; fi78 < 17; fi78++ {
                                fmt.Println(v10)
                                switch (("" + "") + "hello") {
                                case "hello":
                                    if (v3 && (2 > 1)) {
                                        for fi79 := 0; fi79 < 12; fi79++ {
                                            fmt.Println(v3, v3, fi79)
                                            var v80 int32 = v9
                                            var v81 uint8 = 127
                                            _ = fi79
                                            _ = v80
                                            _ = v81
                                        }
                                    } else {
                                        var v82 string = "a"
                                        for fi83 := 0; fi83 < 15; fi83++ {
                                            v82 = ((v82 + ((v82 + v82) + v82)) + "")
                                            var v84 float32 = float32(float64((((-2.25 / 0.1) / -2.25) + -2.25)))
                                            v85 := []int{fi83, fi78, fi83}
                                            if true {
                                                var v86 int = 4294967295
                                                _ = v86
                                            } else {
                                                switch fi83 {
                                                case 1:
                                                    var v87 uint64 = 1
                                                    _ = v87
                                                }
                                            }
                                            _ = fi83
                                            _ = v84
                                            _ = v85
                                        }
                                        _ = v82
                                    }
                                case "":
                                    v88 := GenS7{F0: 2}
                                    _ = v88
                                case "\n\t":
                                    var v89 float64 = 0.1
                                    _ = v89
                                default:
                                    v90 := []int{1, fi78, fi78}
                                    _ = v90
                                }
                                _ = fi78
                            }
                        case "":
                            var v91 int32 = (v9 | v9)
                            _ = v91
                        case "hello":
                            v92 := GenS8{F0: -15}
                            _ = v92
                        default:
                            v10 = (!(42 != 42))
                        }
                    default:
                        var v93 uint8 = 87
                        _ = v93
                    }
                    v94 := GenS2{F0: 2147483647}
                    v95 := []int{562}
                    _ = v94
                    _ = v95
                }
            }
            v13 = (v13 << ((206) & 63))
        } else {
            if (1 < (v9 ^ v9)) {
                if v3 {
                    v10 = (!v3)
                    v96 := []int{4294967295, 0, 32768, 404}
                    v2["foo bar"] = 624
                    _ = v96
                } else {
                    for fi97 := 0; fi97 < 14; fi97++ {
                        switch fi97 {
                        case 127:
                            v4 += (^(v4))
                        case -9223372036854775808:
                            if (v3 && false) {
                                if true {
                                    v98 := GenS9{F0: 0, F1: "a", F2: 128}
                                    var v99 uint32 = 65535
                                    fmt.Println(v9)
                                    _ = v98
                                    _ = v99
                                }
                                v9 += v9
                            } else {
                                v3 = (!false)
                                v100 := []int{(-(fi97)), (fi97 >> ((4294967295) & 63)), fi97, fi97}
                                var v101 uint32 = 245
                                _ = v100
                                _ = v101
                            }
                        case 913:
                            switch fi97 {
                            case -99:
                                if v10 {
                                    v102 := []int{(fi97 >> ((4294967295) & 63)), (-(fi97)), (fi97 << ((2) & 63))}
                                    for fi103 := 0; fi103 < 3; fi103++ {
                                        var v104 bool = (!(-481 >= v9))
                                        _ = fi103
                                        _ = v104
                                    }
                                    _ = v102
                                } else {
                                    var v105 int64 = v4
                                    v10 = (2 >= 2147483647)
                                    _ = v105
                                }
                            case -9223372036854775808:
                                for fi106 := 0; fi106 < 16; fi106++ {
                                    v107 := []int{(^(fi97)), -2147483648, -125, fi97}
                                    v4 = v4
                                    var v108 uint32 = 2147483647
                                    _ = fi106
                                    _ = v107
                                    _ = v108
                                }
                            case 4294967295:
                                for fi109 := 0; fi109 < 1; fi109++ {
                                    v110 := GenS2{F0: 925}
                                    _ = fi109
                                    _ = v110
                                }
                            }
                        default:
                            fmt.Println(v3, v3, v10)
                        }
                        switch "a" {
                        case "foo bar":
                            v13 = v13
                        default:
                            _ = len(v5)
                        }
                        _ = fi97
                    }
                    switch 42 {
                    case 32767:
                        for fi111 := 0; fi111 < 1; fi111++ {
                            v112 := GenS5{F0: (0.0 > (0.1 + 0.0)), F1: false}
                            _ = fi111
                            _ = v112
                        }
                    case -102:
                        v9 += v9
                    }
                    for fi113 := 0; fi113 < 13; fi113++ {
                        v114 := GenS10{F0: (fi113 << ((18446744073709551615) & 63))}
                        var v115 string = "\n\t"
                        _ = fi113
                        _ = v114
                        _ = v115
                    }
                }
                var v116 float64 = 1.5
                v116 = 1.5
                _ = v116
            } else {
                for fi117 := 0; fi117 < 14; fi117++ {
                    v3 = v3
                    fmt.Println(v10, v4, v9)
                    var v118 string = ""
                    _ = fi117
                    _ = v118
                }
                fmt.Println(v10, v10)
                var v119 string = "hello"
                _ = v119
            }
            v120 := []int{-230, -1}
            var v121 uint16 = (^(v13))
            _ = v120
            _ = v121
        }
    }
    var v122 uint8 = 199
    var v123 int64 = v4
    var v124 uint32 = 42
    switch -66 {
    case 32767:
        fmt.Println(v13)
    case -2147483648:
        if v10 {
            if (4294967295 <= 42) {
                if (471 > -9223372036854775808) {
                    v10 = (!((0.0 / (1e10 / 0.1)) > 3.14159))
                    var v125 int32 = 42
                    _ = v125
                }
            } else {
                for fi126 := 0; fi126 < 3; fi126++ {
                    for fi127 := 0; fi127 < 20; fi127++ {
                        _ = len(v5)
                        if v3 {
                            var v128 int8 = 42
                            if (v9 == -2147483648) {
                                v129 := GenS4{F0: 1000}
                                _ = v129
                            }
                            _ = v128
                        }
                        var v130 uint = 984
                        fmt.Println(v130, v3)
                        if true {
                            v131 := GenS11{F0: (v130 / (((v130 - v130)) | 1)), F1: v123}
                            if true {
                                fmt.Println(v123, fi126)
                                if v3 {
                                    switch 418 {
                                    case -1:
                                        _ = len(v1)
                                    default:
                                        fmt.Println(v4, fi127)
                                    }
                                }
                            }
                            v132 := []int{}
                            _ = v131
                            _ = v132
                        } else {
                            var v133 float64 = 1.5
                            _ = v133
                        }
                        v9 += 785
                        _ = fi127
                        _ = v130
                    }
                    for fi134 := 0; fi134 < 18; fi134++ {
                        v135 := []int{fi126, -1, (fi126 - (fi134 / (((fi134 << ((940) & 63))) | 1)))}
                        if (v3 && v10) {
                            switch ((("foo bar" + "a") + "a") + "") {
                            case "foo bar":
                                v2[("hello" + "")] = fi134
                            case "a":
                                if (int8(v4) >= -54) {
                                    _ = len(v5)
                                    v136 := GenS11{F0: 958, F1: (v4 + (^(v123)))}
                                    _ = v136
                                }
                            }
                            fmt.Println(v122, v123)
                            v135 = append(v135, (^(fi134)))
                            var v137 int8 = -31
                            switch "hello" {
                            case "foo bar":
                                v9 = v9
                            case "a":
                                fmt.Println(v124, v123)
                            case "\n\t":
                                v138 := v2[((("foo bar" + "\n\t") + "\n\t") + "foo bar")]
                                _ = v138
                            }
                            _ = v137
                        }
                        var v139 string = "foo bar"
                        _ = fi134
                        _ = v135
                        _ = v139
                    }
                    v123 -= v123
                    _ = fi126
                }
                for fi140 := 0; fi140 < 6; fi140++ {
                    v141 := GenS12{F0: v122, F1: fi140}
                    _ = fi140
                    _ = v141
                }
            }
            switch ("hello" + "a") {
            case "hello":
                for fi142 := 0; fi142 < 14; fi142++ {
                    fmt.Println(v4)
                    var v143 uint8 = 2
                    fmt.Println(v13)
                    _ = fi142
                    _ = v143
                }
            case "a":
                v144 := v1[673]
                _ = v144
            default:
                for fi145 := 0; fi145 < 8; fi145++ {
                    fmt.Println(v122, v124, v124)
                    v146 := GenS9{F0: 218, F1: ("" + "a"), F2: 0}
                    for fi147 := 0; fi147 < 2; fi147++ {
                        var v148 int16 = int16(v4)
                        _ = fi147
                        _ = v148
                    }
                    _ = fi145
                    _ = v146
                }
            }
            for fi149 := 0; fi149 < 8; fi149++ {
                switch "\n\t" {
                case "":
                    v13 = v13
                case "a":
                    v123 -= (v4 >> ((224) & 63))
                case "foo bar":
                    switch (fi149 ^ fi149) {
                    case 0:
                        var v150 int64 = v4
                        _ = v150
                    default:
                        if (2 > 1) {
                            switch "hello" {
                            case "":
                                switch "a" {
                                case "\n\t":
                                    v151 := GenS9{F0: 2, F1: "hello", F2: (^(v122))}
                                    _ = v151
                                case "foo bar":
                                    fmt.Println(v122, v122)
                                }
                            case "a":
                                switch -287 {
                                case 328:
                                    v122 = v122
                                case -329:
                                    if v10 {
                                        v152 := []int{(^(fi149)), -674}
                                        _ = v152
                                    }
                                case 786:
                                    for fi153 := 0; fi153 < 14; fi153++ {
                                        v154 := []int{}
                                        _ = fi153
                                        _ = v154
                                    }
                                }
                            }
                            fmt.Println(fi149)
                            if (911 != v9) {
                                for fi155 := 0; fi155 < 4; fi155++ {
                                    var v156 bool = false
                                    _ = fi155
                                    _ = v156
                                }
                                if (48 >= -128) {
                                    var v157 float32 = float32((1e10 * (0.1 / -2.25)))
                                    var v158 uint16 = v13
                                    v2["foo bar"] = fi149
                                    _ = v157
                                    _ = v158
                                } else {
                                    fmt.Println(v9, v3, v122)
                                }
                            } else {
                                for fi159 := 0; fi159 < 4; fi159++ {
                                    v9 -= v9
                                    fmt.Println(v4, v4)
                                    v2["\n\t"] = -9223372036854775808
                                    _ = fi159
                                }
                                var v160 int = (fi149 & 32767)
                                for fi161 := 0; fi161 < 4; fi161++ {
                                    v162 := GenS13{F0: 891, F1: ("\n\t" + "foo bar"), F2: (v13 >> ((uint(v124)) & 63))}
                                    var v163 uint32 = (v124 * v124)
                                    fmt.Println(fi149, v10, v160)
                                    v124 -= 881
                                    _ = fi161
                                    _ = v162
                                    _ = v163
                                }
                                _ = v160
                            }
                        }
                    }
                default:
                    switch ("\n\t" + "") {
                    case "":
                        if (516 != fi149) {
                            v164 := GenS5{F0: v3, F1: false}
                            _ = v164
                        } else {
                            for fi165 := 0; fi165 < 20; fi165++ {
                                if false {
                                    switch fi165 {
                                    case -1:
                                        v3 = (-1 == 0)
                                    case -9223372036854775808:
                                        var v166 uint16 = 2
                                        _ = v166
                                    }
                                } else {
                                    v123 -= (v4 * v123)
                                    var v167 float32 = (1.5 * -2.25)
                                    _ = v167
                                }
                                fmt.Println(v122, v13, fi165)
                                _ = fi165
                            }
                            var v168 int16 = 958
                            _ = v168
                        }
                    case "\n\t":
                        if v10 {
                            for fi169 := 0; fi169 < 14; fi169++ {
                                v170 := GenS14{F0: false, F1: (69 <= 42)}
                                _ = fi169
                                _ = v170
                            }
                            fmt.Println(v124)
                            if true {
                                v122 = 128
                            }
                        } else {
                            v171 := GenS7{F0: 4294967295}
                            var v172 int32 = -2147483648
                            _ = v171
                            _ = v172
                        }
                    default:
                        for fi173 := 0; fi173 < 5; fi173++ {
                            if v3 {
                                v174 := []int{(fi173 | (fi149 / ((fi149) | 1))), 127, fi173, fi149}
                                _ = len(v174)
                                var v175 float32 = (1.5 / -2.25)
                                _ = v174
                                _ = v175
                            } else {
                                v124 = (v124 << ((1) & 63))
                                var v176 uint = 65535
                                v9 = v9
                                _ = v176
                            }
                            fmt.Println(v9, fi173)
                            _ = fi173
                        }
                    }
                }
                _ = fi149
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v122:", v122)
    fmt.Println("v123:", v123)
    fmt.Println("v124:", v124)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v122
    _ = v123
    _ = v124
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
