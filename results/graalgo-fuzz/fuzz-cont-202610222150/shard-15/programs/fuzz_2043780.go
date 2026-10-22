package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 uint32 }
type GenS6 struct { F0 uint8; F1 int64 }
type GenS7 struct { F0 int32; F1 string; F2 bool }
type GenS8 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS9 struct { F0 int8; F1 int32; F2 uint }
type GenS10 struct { F0 uint8; F1 int32; F2 int8 }
type GenS11 struct { F0 int8; F1 int8; F2 int64 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 uint64 }
func genFn0() int16 {
    v4 := []int{0, 32768, 1}
    var v5 string = "a"
    _ = v4
    _ = v5
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{"foo bar", "foo bar"}
    var v3 float64 = 0.0
    if true {
        for fi6 := 0; fi6 < 13; fi6++ {
            var v7 uint16 = 42
            _ = fi6
            _ = v7
        }
        switch 32767 {
        case -1:
            var v8 int = -9223372036854775808
            _ = v8
        case 2147483647:
            if true {
                if false {
                    if true {
                        var v9 uint64 = 42
                        switch -9223372036854775808 {
                        case 128:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case 0:
                            var v10 string = ("hello" + ("" + (("\n\t" + "\n\t") + ("\n\t" + "hello"))))
                            _ = v10
                        case -2147483648:
                            v11 := GenS1{F0: 29, F1: 2147483648, F2: true}
                            _ = v11
                        default:
                            var v12 uint8 = 115
                            _ = v12
                        }
                        var v13 bool = (-73 >= 84)
                        _ = v13
                        _ = v9
                    }
                    switch 999 {
                    case 4294967295:
                        if (0 > -480) {
                            var v14 int16 = 1
                            _ = v14
                        } else {
                            v1[704] = 4294967295
                            v15 := GenS2{F0: -32768}
                            if ((228 < 9223372036854775807) && (false && true)) {
                                var v16 float64 = v3
                                var v17 int16 = 32767
                                _ = v16
                                _ = v17
                            }
                            _ = v15
                        }
                    default:
                        fmt.Println(v3)
                    }
                    v2 = append(v2, ("foo bar" + "hello"))
                }
                var v18 int32 = 0
                _ = v18
            } else {
                var v19 uint16 = 19
                var v20 string = "\n\t"
                _ = v19
                _ = v20
            }
        }
        for fi21 := 0; fi21 < 15; fi21++ {
            var v22 string = "\n\t"
            _ = fi21
            _ = v22
        }
    }
    _ = len(v1)
    v23 := GenS3{F0: 494, F1: 0}
    if true {
        if (false && (26 <= 0)) {
            if (false && false) {
                v3 = (v3 * float64(float32(0.0)))
            } else {
                fmt.Println(v3)
                v1[300] = 1
                if (977 > 0) {
                    if (true || (0 > 0)) {
                        switch 708 {
                        case 32768:
                            v24 := GenS1{F0: 71, F1: 2, F2: (false || ((1 > 2147483647) && false))}
                            _ = v24
                        }
                        _ = len(v1)
                    }
                }
            }
            if len(v2) > 0 {
                v2[0] = ("\n\t" + ("\n\t" + ""))
            }
        } else {
            if (255 != 255) {
                v3 = ((0.0 - (3.14159 / v3)) - 1.5)
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
            }
            if false {
                for fi25 := 0; fi25 < 12; fi25++ {
                    v3 -= v3
                    fmt.Println(fi25)
                    var v26 uint16 = 255
                    _ = fi25
                    _ = v26
                }
            }
            v27 := GenS4{F0: 42}
            _ = v27
        }
    }
    v3 += 0.1
    delete(v1, 2)
    if true {
        var v28 uint32 = 2147483647
        if true {
            if (18 == 224) {
                for fi29 := 0; fi29 < 4; fi29++ {
                    fmt.Println(v28, v28)
                    _ = len(v1)
                    v3 = v3
                    _ = fi29
                }
                if (-623 > -1) {
                    for fi30 := 0; fi30 < 9; fi30++ {
                        for fi31 := 0; fi31 < 4; fi31++ {
                            v32 := GenS2{F0: genFn0()}
                            v33 := GenS5{F0: 1, F1: 42}
                            var v34 bool = (0 == 617)
                            _ = fi31
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                        _ = fi30
                    }
                } else {
                    v35 := GenS4{F0: -446}
                    var v36 uint64 = 4294967295
                    _ = v35
                    _ = v36
                }
                for fi37 := 0; fi37 < 14; fi37++ {
                    fmt.Println(v28)
                    _ = fi37
                }
            } else {
                if false {
                    v28 = (v28 + 2147483647)
                    var v38 int32 = 42
                    var v39 int32 = v38
                    _ = v38
                    _ = v39
                }
                v40 := GenS2{F0: genFn0()}
                if (1 >= 799) {
                    v28 += 769
                    for fi41 := 0; fi41 < 2; fi41++ {
                        v42 := GenS1{F0: -51, F1: (v28 ^ 979), F2: false}
                        var v43 string = "a"
                        if true {
                            var v44 int64 = -1
                            switch fi41 {
                            case -838:
                                switch fi41 {
                                case -1:
                                    fmt.Println(v43, v28, v3)
                                case 4294967295:
                                    v28 += v28
                                case -512:
                                    if true {
                                        var v45 int16 = -32768
                                        fmt.Println(v28, v3, v45)
                                        var v46 bool = true
                                        _ = v45
                                        _ = v46
                                    }
                                }
                            default:
                                var v47 bool = ((3.14159 / v3) == 1.5)
                                _ = v47
                            }
                            _ = v44
                        } else {
                            for fi48 := 0; fi48 < 14; fi48++ {
                                if false {
                                    var v49 int32 = -14
                                    if false {
                                        v1[2147483648] = fi41
                                    }
                                    fmt.Println(v28, fi48, v28)
                                    _ = v49
                                } else {
                                    v1[1] = 32768
                                    fmt.Println(fi41, v43, fi41)
                                }
                                if false {
                                    var v50 int32 = 1
                                    if (0 < 255) {
                                        v28 = (v28 + v28)
                                        var v51 uint16 = 2
                                        _ = v51
                                    } else {
                                        _ = len(v1)
                                    }
                                    delete(v1, 737)
                                    _ = v50
                                } else {
                                    var v52 float32 = ((0.1 * 1.5) / 3.14159)
                                    var v53 int = -2147483648
                                    v53 += -572
                                    var v54 float32 = v52
                                    _ = v52
                                    _ = v53
                                    _ = v54
                                }
                                v55 := GenS2{F0: 478}
                                var v56 string = "a"
                                _ = fi48
                                _ = v55
                                _ = v56
                            }
                        }
                        _ = fi41
                        _ = v42
                        _ = v43
                    }
                    v57 := GenS6{F0: 1, F1: -620}
                    _ = v57
                } else {
                    _ = len(v1)
                    fmt.Println(v28, v28)
                    switch "foo bar" {
                    case "":
                        if (v3 <= (float64((3.14159 * ((3.14159 * (-2.25 + (1.5 * -2.25))) + -2.25))) + 0.0)) {
                            v3 = v3
                            v28 += (v28 - 4294967295)
                        } else {
                            var v58 uint16 = 0
                            if true {
                                for fi59 := 0; fi59 < 18; fi59++ {
                                    for fi60 := 0; fi60 < 18; fi60++ {
                                        var v61 float64 = v3
                                        _ = fi60
                                        _ = v61
                                    }
                                    switch 4294967295 {
                                    case 267:
                                        v62 := GenS7{F0: 2147483647, F1: "hello", F2: false}
                                        _ = v62
                                    case 32768:
                                        var v63 bool = (true && (-2.25 != (v3 * v3)))
                                        _ = v63
                                    default:
                                        fmt.Println(v28, v28, v58)
                                    }
                                    _ = fi59
                                }
                                v58 = v58
                            } else {
                                switch ("foo bar" + "\n\t") {
                                case "\n\t":
                                    _ = len(v2)
                                }
                                v3 = 0.1
                                if false {
                                    var v64 int8 = -58
                                    fmt.Println(v3, v3, v3)
                                    for fi65 := 0; fi65 < 8; fi65++ {
                                        v64 = (v64 >> ((2) & 63))
                                        v66 := GenS8{F0: 42, F1: v58, F2: 634}
                                        _ = fi65
                                        _ = v66
                                    }
                                    _ = v64
                                } else {
                                    fmt.Println(v58, v28)
                                    var v67 int16 = 1000
                                    _ = v67
                                }
                            }
                            if true {
                                if false {
                                    for fi68 := 0; fi68 < 5; fi68++ {
                                        v58 = (v58 ^ v58)
                                        _ = fi68
                                    }
                                    if true {
                                        if false {
                                            if (!false) {
                                                switch -1 {
                                                case 42:
                                                    if true {
                                                        v28 = v28
                                                    } else {
                                                        v69 := GenS8{F0: 0, F1: v58, F2: 2}
                                                        fmt.Println(v3, v3, v28)
                                                        _ = len(v1)
                                                        _ = v69
                                                    }
                                                case -42:
                                                    if (0.1 == ((float64(float32(v3)) - float64(3.14159)) - 1e10)) {
                                                        for fi70 := 0; fi70 < 6; fi70++ {
                                                            if (42 > 300) {
                                                                v3 = v3
                                                            }
                                                            _ = fi70
                                                        }
                                                    }
                                                default:
                                                    var v71 uint64 = 0
                                                    _ = v71
                                                }
                                                fmt.Println(v3)
                                                for fi72 := 0; fi72 < 8; fi72++ {
                                                    v3 -= 0.0
                                                    for fi73 := 0; fi73 < 8; fi73++ {
                                                        switch "hello" {
                                                        case "":
                                                            v74 := GenS3{F0: 550, F1: 1}
                                                            _ = v74
                                                        case "hello":
                                                            v1[v28] = (fi72 * (-(fi73)))
                                                        case "a":
                                                            v75 := GenS9{F0: 100, F1: 0, F2: 168}
                                                            _ = v75
                                                        }
                                                        _ = fi73
                                                    }
                                                    if (!true) {
                                                        for fi76 := 0; fi76 < 9; fi76++ {
                                                            v28 += (v28 | v28)
                                                            if (438 == 2147483647) {
                                                                v2 = append(v2, "hello")
                                                                v28 += v28
                                                            } else {
                                                                var v77 uint = 0
                                                                _ = v77
                                                            }
                                                            switch fi72 {
                                                            case 588:
                                                                v78 := GenS9{F0: -1, F1: 2147483647, F2: 42}
                                                                _ = v78
                                                            case 32768:
                                                                switch "hello" {
                                                                case "a":
                                                                    v79 := GenS1{F0: 0, F1: (^(v28)), F2: (255 < 925)}
                                                                    _ = v79
                                                                case "":
                                                                    var v80 uint16 = (v58 + (^(v58)))
                                                                    _ = v80
                                                                default:
                                                                    if true {
                                                                        if true {
                                                                            if len(v2) > 0 {
                                                                                v2[0] = "foo bar"
                                                                            }
                                                                            v81 := GenS9{F0: 100, F1: -2147483648, F2: 42}
                                                                            var v82 bool = true
                                                                            _ = v81
                                                                            _ = v82
                                                                        }
                                                                        fmt.Println(v58, fi76, v3)
                                                                    }
                                                                    if (2147483647 != -2147483648) {
                                                                        if true {
                                                                            var v83 bool = (v28 <= (v28 << ((42) & 63)))
                                                                            var v84 int64 = 0
                                                                            _ = v83
                                                                            _ = v84
                                                                        } else {
                                                                            v85 := GenS9{F0: 42, F1: -1, F2: 1}
                                                                            switch (fi72 & 0) {
                                                                            case 32767:
                                                                                v58 = (v58 | v58)
                                                                            case 9223372036854775807:
                                                                                v86 := GenS3{F0: -32768, F1: 4294967295}
                                                                                _ = v86
                                                                            }
                                                                            switch ("foo bar" + "\n\t") {
                                                                            case "":
                                                                                v3 = v3
                                                                            }
                                                                            if true {
                                                                                switch fi76 {
                                                                                case 32768:
                                                                                    fmt.Println(v28, fi76)
                                                                                case 4294967295:
                                                                                    if ("\n\t" < "foo bar") {
                                                                                        _ = len(v2)
                                                                                    } else {
                                                                                        fmt.Println(v58)
                                                                                        if false {
                                                                                            v28 += v28
                                                                                            v28 = v28
                                                                                            v3 = 1e10
                                                                                            fmt.Println(v3, v58)
                                                                                            v1[v28] = fi72
                                                                                            v28 += (^(v28))
                                                                                            if (1e10 == 1e10) {
                                                                                                fmt.Println(v3)
                                                                                            } else {
                                                                                                var v87 uint64 = 60
                                                                                                _ = v87
                                                                                            }
                                                                                            var v88 float32 = 0.0
                                                                                            fmt.Println(fi76, v58)
                                                                                            if (255 >= 1) {
                                                                                                var v89 int64 = -530
                                                                                                _ = v89
                                                                                            } else {
                                                                                                v90 := GenS4{F0: -856}
                                                                                                v91 := GenS2{F0: genFn0()}
                                                                                                _ = v90
                                                                                                _ = v91
                                                                                            }
                                                                                            _ = v88
                                                                                        }
                                                                                        var v92 uint16 = 255
                                                                                        _ = v92
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                fmt.Println(fi76)
                                                                                var v93 uint64 = 4294967295
                                                                                if false {
                                                                                    var v94 uint64 = v93
                                                                                    v95 := GenS10{F0: 128, F1: 2147483647, F2: 1}
                                                                                    _ = v94
                                                                                    _ = v95
                                                                                } else {
                                                                                    var v96 uint32 = v28
                                                                                    _ = v96
                                                                                }
                                                                                _ = v93
                                                                            }
                                                                            fmt.Println(v3)
                                                                            _ = v85
                                                                        }
                                                                    } else {
                                                                        v97 := GenS4{F0: -795}
                                                                        v28 = (v28 << ((977) & 63))
                                                                        _ = v97
                                                                    }
                                                                }
                                                            case 6:
                                                                var v98 float64 = 1.5
                                                                _ = v98
                                                            default:
                                                                switch -787 {
                                                                case -9223372036854775808:
                                                                    delete(v1, v28)
                                                                case 545:
                                                                    switch fi76 {
                                                                    case 434:
                                                                        v58 = (v58 - v58)
                                                                    case -318:
                                                                        var v99 float64 = v3
                                                                        _ = v99
                                                                    case 2147483647:
                                                                        fmt.Println(v3, fi72)
                                                                    default:
                                                                        v1[v28] = -43
                                                                    }
                                                                    v3 = v3
                                                                    v3 = v3
                                                                }
                                                            }
                                                            var v100 int32 = 1
                                                            _ = fi76
                                                            _ = v100
                                                        }
                                                        v58 = v58
                                                        switch "hello" {
                                                        case "\n\t":
                                                            fmt.Println(v28, fi72)
                                                        default:
                                                            v3 = (v3 / (v3 + 1.5))
                                                        }
                                                    }
                                                    _ = fi72
                                                }
                                            }
                                            if true {
                                                var v101 uint64 = 841
                                                _ = v101
                                            }
                                        }
                                        var v102 float64 = 1.5
                                        _ = v102
                                    }
                                    for fi103 := 0; fi103 < 12; fi103++ {
                                        v58 += (^(v58))
                                        _ = fi103
                                    }
                                } else {
                                    for fi104 := 0; fi104 < 14; fi104++ {
                                        var v105 uint = 18446744073709551615
                                        _ = fi104
                                        _ = v105
                                    }
                                    v28 = v28
                                    v2 = append(v2, "")
                                }
                            } else {
                                var v106 float64 = v3
                                var v107 int64 = 9223372036854775807
                                var v108 int = 32768
                                _ = v106
                                _ = v107
                                _ = v108
                            }
                            _ = v58
                        }
                    default:
                        fmt.Println(v3, v28)
                    }
                }
                _ = v40
            }
        } else {
            v28 = 2147483647
            for fi109 := 0; fi109 < 13; fi109++ {
                var v110 float64 = (((v3 + ((v3 * v3) - v3)) - (v3 - ((v3 + v3) * 1e10))) - 0.1)
                _ = fi109
                _ = v110
            }
        }
        var v111 uint = 0
        _ = v111
        _ = v28
    }
    if true {
        v3 -= 0.1
        v3 -= v3
    }
    fmt.Println(v3)
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    v3 = v3
    if ((9223372036854775807 >= 1) && (0.1 > 1.5)) {
        var v112 uint8 = 90
        fmt.Println(v112)
        v112 += 42
        _ = v112
    } else {
        for fi113 := 0; fi113 < 12; fi113++ {
            for fi114 := 0; fi114 < 19; fi114++ {
                var v115 float32 = (1.5 / (1.5 / -2.25))
                _ = fi114
                _ = v115
            }
            if false {
                v116 := GenS6{F0: 127, F1: -1}
                fmt.Println(v3)
                var v117 uint16 = 42
                _ = v116
                _ = v117
            }
            switch (("foo bar" + "") + "a") {
            case "":
                var v118 int64 = -35
                _ = v118
            case "foo bar":
                v1[857] = 9223372036854775807
            default:
                var v119 int16 = genFn0()
                _ = v119
            }
            _ = fi113
        }
        for fi120 := 0; fi120 < 2; fi120++ {
            switch "hello" {
            case "foo bar":
                for fi121 := 0; fi121 < 15; fi121++ {
                    var v122 uint16 = 65535
                    _ = fi121
                    _ = v122
                }
            case "":
                v3 = v3
            }
            _ = fi120
        }
    }
    if (!false) {
        for fi123 := 0; fi123 < 11; fi123++ {
            switch 32768 {
            case -441:
                v3 = 1e10
            default:
                switch "a" {
                case "a":
                    var v124 bool = ((!false) && (true || true))
                    _ = v124
                }
            }
            var v125 uint = 4294967295
            if false {
                for fi126 := 0; fi126 < 5; fi126++ {
                    fmt.Println(fi126, v3, v125)
                    v1[661] = (fi126 % ((fi123) | 1))
                    v3 += (v3 * v3)
                    _ = fi126
                }
                if false {
                    var v127 int64 = 1
                    v125 = 4294967295
                    _ = v127
                }
                var v128 float64 = float64(1.5)
                _ = v128
            } else {
                if ((fi123 / ((2147483647) | 1)) != fi123) {
                    if (true && true) {
                        _ = len(v2)
                    }
                    switch 583 {
                    case -9223372036854775808:
                        if false {
                            var v129 int = 42
                            _ = v129
                        }
                    default:
                        v130 := GenS11{F0: 100, F1: -103, F2: -9223372036854775808}
                        _ = v130
                    }
                    var v131 float64 = 0.0
                    _ = v131
                } else {
                    v132 := GenS1{F0: 42, F1: 2147483648, F2: false}
                    switch "foo bar" {
                    case "foo bar":
                        var v133 string = ""
                        _ = v133
                    case "\n\t":
                        if true {
                            _ = len(v2)
                        } else {
                            var v134 uint32 = 2147483648
                            _ = len(v1)
                            v135 := GenS2{F0: 164}
                            _ = v134
                            _ = v135
                        }
                    case "":
                        fmt.Println(v3, v125)
                    default:
                        v136 := GenS9{F0: 100, F1: -960, F2: v125}
                        _ = v136
                    }
                    _ = v132
                }
                for fi137 := 0; fi137 < 4; fi137++ {
                    _ = len(v1)
                    v3 = 1.5
                    if true {
                        v3 += (1.5 + float64((0.0 - (float32((1e10 - (v3 + 3.14159))) + 0.0))))
                    } else {
                        switch "\n\t" {
                        case "a":
                            fmt.Println(v125)
                        }
                        if (true && (127 >= 127)) {
                            var v138 string = "\n\t"
                            if true {
                                var v139 float32 = 3.14159
                                _ = v139
                            }
                            var v140 int = (-(fi123))
                            _ = v138
                            _ = v140
                        } else {
                            switch fi137 {
                            case 42:
                                if (!false) {
                                    var v141 float32 = 0.1
                                    v142 := GenS12{F0: "\n\t"}
                                    v3 = v3
                                    _ = v141
                                    _ = v142
                                } else {
                                    v143 := GenS13{F0: 4294967295}
                                    switch "" {
                                    case "foo bar":
                                        var v144 uint8 = 1
                                        _ = v144
                                    }
                                    _ = v143
                                }
                                switch -675 {
                                case -1:
                                    v3 += v3
                                default:
                                    _ = len(v1)
                                }
                            case -679:
                                v145 := GenS10{F0: 141, F1: -2147483648, F2: -1}
                                _ = v145
                            }
                            var v146 float64 = float64(0.0)
                            v3 = (-2.25 / 0.1)
                            _ = v146
                        }
                        var v147 int16 = 32767
                        _ = v147
                    }
                    var v148 uint8 = 255
                    _ = fi137
                    _ = v148
                }
            }
            _ = fi123
            _ = v125
        }
        v3 = (1.5 / v3)
        switch "a" {
        case "\n\t":
            for fi149 := 0; fi149 < 10; fi149++ {
                v150 := GenS3{F0: 1, F1: 0}
                _ = fi149
                _ = v150
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
}
