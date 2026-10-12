package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool; F1 int32 }
type GenS7 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int16; F1 int16 }
type GenS10 struct { F0 uint32; F1 int16; F2 int8 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 int64; F1 bool; F2 uint8 }
type GenS13 struct { F0 string; F1 int; F2 int32 }
type GenS14 struct { F0 int16; F1 bool }
func genFn0() uint8 {
    if true {
        if (42 < 151) {
            v6 := []int{0, 128}
            var v7 uint8 = 130
            _ = v6
            _ = v7
        } else {
            var v8 int32 = -774
            v8 = 2147483647
            v9 := GenS1{F0: 784}
            _ = v8
            _ = v9
        }
        if true {
            var v10 bool = true
            _ = v10
        } else {
            if (!true) {
                var v11 int64 = -9223372036854775808
                for fi12 := 0; fi12 < 15; fi12++ {
                    if (2 > 496) {
                        v13 := GenS2{F0: (fi12 - (fi12 & -9223372036854775808)), F1: true}
                        _ = v13
                    }
                    switch "" {
                    case "hello":
                        v11 += (v11 * int64(fi12))
                    case "foo bar":
                        for fi14 := 0; fi14 < 20; fi14++ {
                            fmt.Println(fi12)
                            v11 = -1
                            _ = fi14
                        }
                    default:
                        v15 := GenS3{F0: 207, F1: 1}
                        _ = v15
                    }
                    _ = fi12
                }
                for fi16 := 0; fi16 < 7; fi16++ {
                    var v17 int = (fi16 >> ((0) & 63))
                    v18 := []int{-542}
                    for fi19 := 0; fi19 < 20; fi19++ {
                        v17 = fi19
                        if ((-2.25 / (0.1 * 3.14159)) != 0.0) {
                            v17 -= v17
                            v18 = append(v18, (v17 + 128))
                            v20 := GenS4{F0: fi19}
                            _ = v20
                        }
                        _ = fi19
                    }
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                _ = v11
            } else {
                switch -901 {
                case 32767:
                    if true {
                        var v21 uint8 = 128
                        v22 := []int{32767, -2147483648}
                        var v23 uint = 631
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                case -989:
                    v24 := GenS5{F0: ("" + ("foo bar" + ""))}
                    _ = v24
                }
            }
            var v25 int8 = 100
            _ = v25
        }
    } else {
        var v26 float64 = (1e10 - 3.14159)
        v26 += (1e10 * 1e10)
        _ = v26
    }
    for fi27 := 0; fi27 < 20; fi27++ {
        v28 := []int{(fi27 * fi27), fi27}
        var v29 uint32 = 392
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 float32 = 1e10
    if ((1000 < 32767) && false) {
        var v31 int8 = 42
        if true {
            v31 = 1
            v32 := []int{-2147483648}
            _ = v32
        } else {
            for fi33 := 0; fi33 < 16; fi33++ {
                v30 = float32(-2.25)
                var v34 uint32 = 72
                _ = fi33
                _ = v34
            }
            if (42 >= 2147483648) {
                if (true || false) {
                    v31 = 119
                    if true {
                        v30 += v30
                        switch "\n\t" {
                        case "a":
                            fmt.Println(v30, v31)
                        case "hello":
                            if (false && (true && true)) {
                                if true {
                                    fmt.Println(v30, v31)
                                    v31 -= (^(v31))
                                    if false {
                                        v30 -= (v30 - v30)
                                        for fi35 := 0; fi35 < 13; fi35++ {
                                            v36 := []int{0}
                                            _ = fi35
                                            _ = v36
                                        }
                                        for fi37 := 0; fi37 < 4; fi37++ {
                                            var v38 int = (fi37 ^ 837)
                                            _ = fi37
                                            _ = v38
                                        }
                                    } else {
                                        v39 := []int{-9223372036854775808, 32768, 9223372036854775807, 327}
                                        for fi40 := 0; fi40 < 11; fi40++ {
                                            if true {
                                                switch "foo bar" {
                                                case "foo bar":
                                                    if true {
                                                        if len(v39) > 0 {
                                                            v39[0] = (fi40 * (-(fi40)))
                                                        }
                                                        v30 = v30
                                                    } else {
                                                        for fi41 := 0; fi41 < 20; fi41++ {
                                                            var v42 float64 = 3.14159
                                                            _ = fi41
                                                            _ = v42
                                                        }
                                                        v43 := GenS2{F0: (^(fi40)), F1: (true || false)}
                                                        _ = len(v39)
                                                        _ = v43
                                                    }
                                                case "":
                                                    var v44 int32 = 2147483647
                                                    _ = v44
                                                default:
                                                    var v45 int8 = v31
                                                    _ = v45
                                                }
                                                for fi46 := 0; fi46 < 19; fi46++ {
                                                    v30 = v30
                                                    switch "a" {
                                                    case "foo bar":
                                                        fmt.Println(fi46)
                                                    case "\n\t":
                                                        v30 = v30
                                                    default:
                                                        var v47 uint32 = 81
                                                        _ = v47
                                                    }
                                                    var v48 int32 = 2147483647
                                                    _ = fi46
                                                    _ = v48
                                                }
                                                for fi49 := 0; fi49 < 7; fi49++ {
                                                    var v50 uint16 = 2
                                                    switch ("a" + ("" + (("" + "a") + "\n\t"))) {
                                                    case "foo bar":
                                                        _ = len(v39)
                                                    }
                                                    if len(v39) >= 2 {
                                                        v39 = v39[1:len(v39)]
                                                    }
                                                    _ = fi49
                                                    _ = v50
                                                }
                                            } else {
                                                fmt.Println(fi40)
                                                for fi51 := 0; fi51 < 13; fi51++ {
                                                    switch (fi40 ^ 1) {
                                                    case -47:
                                                        if true {
                                                            if ((!(!((!false) && (42 >= 18446744073709551615)))) || (-435 <= 0)) {
                                                                switch (fi40 + 4294967295) {
                                                                case 128:
                                                                    v30 = v30
                                                                }
                                                                var v52 int32 = -1
                                                                v52 += v52
                                                                _ = v52
                                                            }
                                                            switch "\n\t" {
                                                            case "":
                                                                fmt.Println(fi51, v31)
                                                            case "hello":
                                                                if (!false) {
                                                                    v53 := GenS3{F0: 18446744073709551615, F1: 4294967295}
                                                                    _ = v53
                                                                } else {
                                                                    v30 = v30
                                                                    if (683 < 384) {
                                                                        v30 = (v30 - v30)
                                                                        switch 2147483647 {
                                                                        case 9223372036854775807:
                                                                            v30 -= v30
                                                                        case 32768:
                                                                            if ((-32768 < 1) || true) {
                                                                                var v54 uint32 = 437
                                                                                _ = v54
                                                                            }
                                                                        case 127:
                                                                            if len(v39) >= 2 {
                                                                                v39 = v39[1:len(v39)]
                                                                            }
                                                                        }
                                                                        if true {
                                                                            switch 4294967295 {
                                                                            case -88:
                                                                                fmt.Println(fi51, v31, v30)
                                                                            case -1:
                                                                                fmt.Println(v31)
                                                                            default:
                                                                                if len(v39) > 0 {
                                                                                    v39[0] = 1
                                                                                }
                                                                            }
                                                                        }
                                                                    } else {
                                                                        fmt.Println(fi40, v31)
                                                                    }
                                                                    if (false || false) {
                                                                        v31 -= 1
                                                                        v39 = append(v39, 128)
                                                                    } else {
                                                                        v55 := GenS3{F0: 2, F1: 2147483648}
                                                                        _ = v55
                                                                    }
                                                                }
                                                            case "a":
                                                                v56 := GenS1{F0: 4294967295}
                                                                v30 += v30
                                                                var v57 int32 = 1
                                                                _ = v56
                                                                _ = v57
                                                            }
                                                            v58 := GenS2{F0: (fi51 | (fi51 - fi40)), F1: false}
                                                            _ = v58
                                                        }
                                                    default:
                                                        v39 = append(v39, fi51)
                                                    }
                                                    _ = fi51
                                                }
                                                for fi59 := 0; fi59 < 17; fi59++ {
                                                    var v60 int = (fi40 * (fi40 * fi59))
                                                    _ = fi59
                                                    _ = v60
                                                }
                                            }
                                            switch ("a" + "foo bar") {
                                            case "foo bar":
                                                v31 -= 100
                                            case "":
                                                fmt.Println(v30)
                                            case "hello":
                                                if len(v39) >= 2 {
                                                    v39 = v39[1:len(v39)]
                                                }
                                            default:
                                                for fi61 := 0; fi61 < 5; fi61++ {
                                                    v62 := GenS6{F0: (!(65535 > 932)), F1: -2147483648}
                                                    v63 := GenS2{F0: (fi40 / ((fi40) | 1)), F1: (v30 == -2.25)}
                                                    _ = fi61
                                                    _ = v62
                                                    _ = v63
                                                }
                                            }
                                            var v64 int8 = v31
                                            _ = fi40
                                            _ = v64
                                        }
                                        _ = v39
                                    }
                                }
                            }
                        case "":
                            v30 += 1e10
                        default:
                            switch 577 {
                            case -970:
                                if false {
                                    if true {
                                        var v65 uint8 = 140
                                        var v66 uint64 = 4294967295
                                        for fi67 := 0; fi67 < 11; fi67++ {
                                            if false {
                                                if (100 == (v31 & (v31 & 1))) {
                                                    for fi68 := 0; fi68 < 1; fi68++ {
                                                        var v69 float32 = v30
                                                        v69 = v69
                                                        v70 := []int{-398, fi67, fi68, (fi67 - -1)}
                                                        v71 := GenS1{F0: 4294967295}
                                                        v69 = v69
                                                        fmt.Println(v31)
                                                        _ = fi68
                                                        _ = v69
                                                        _ = v70
                                                        _ = v71
                                                    }
                                                    fmt.Println(fi67, v30, v66)
                                                } else {
                                                    v30 = ((((v30 * (float32(1e10) / v30)) / v30) / v30) * v30)
                                                }
                                                v30 = float32(1e10)
                                                v72 := GenS4{F0: int(v65)}
                                                _ = v72
                                            }
                                            _ = fi67
                                        }
                                        _ = v65
                                        _ = v66
                                    } else {
                                        for fi73 := 0; fi73 < 20; fi73++ {
                                            if (!false) {
                                                v74 := GenS3{F0: 229, F1: 0}
                                                _ = v74
                                            } else {
                                                v75 := GenS4{F0: 9223372036854775807}
                                                _ = v75
                                            }
                                            _ = fi73
                                        }
                                        fmt.Println(v31, v30)
                                        var v76 int64 = -1
                                        _ = v76
                                    }
                                } else {
                                    if (1 <= 127) {
                                        if (!false) {
                                            v30 = ((v30 / v30) + 0.1)
                                        }
                                        var v77 int32 = 1
                                        _ = v77
                                    }
                                    var v78 float32 = 1e10
                                    _ = v78
                                }
                            case 32767:
                                v30 -= float32((-2.25 / 0.1))
                            case 310:
                                v79 := GenS3{F0: 18446744073709551615, F1: 42}
                                _ = v79
                            default:
                                switch 42 {
                                case -9223372036854775808:
                                    for fi80 := 0; fi80 < 16; fi80++ {
                                        switch "foo bar" {
                                        case "\n\t":
                                            fmt.Println(v30, v30, v31)
                                        case "a":
                                            v30 -= (v30 / ((v30 * (float32((0.0 / float64(v30))) * ((1e10 + 3.14159) - float32(0.0)))) / 1e10))
                                        case "foo bar":
                                            switch fi80 {
                                            case -320:
                                                for fi81 := 0; fi81 < 11; fi81++ {
                                                    v82 := []int{(^(fi81)), 2147483647, fi81}
                                                    if true {
                                                        v83 := GenS7{F0: (v31 | (v31 / (((v31 << ((0) & 63))) | 1))), F1: 128, F2: 128}
                                                        var v84 uint32 = 42
                                                        v82 = append(v82, 32767)
                                                        _ = v83
                                                        _ = v84
                                                    }
                                                    v82 = append(v82, fi81)
                                                    if (true && (-9223372036854775808 == 9223372036854775807)) {
                                                        v30 += v30
                                                        v85 := GenS6{F0: false, F1: 42}
                                                        v31 = -128
                                                        _ = v85
                                                    } else {
                                                        if false {
                                                            if len(v82) >= 2 {
                                                                v82 = v82[1:len(v82)]
                                                            }
                                                        } else {
                                                            v31 -= v31
                                                            if len(v82) > 0 {
                                                                v82[0] = fi81
                                                            }
                                                            var v86 uint32 = 2147483648
                                                            _ = v86
                                                        }
                                                    }
                                                    _ = fi81
                                                    _ = v82
                                                }
                                            case 1:
                                                v87 := []int{(-(fi80)), fi80}
                                                _ = v87
                                            case -425:
                                                var v88 uint = 18446744073709551615
                                                _ = v88
                                            default:
                                                var v89 int16 = 32767
                                                _ = v89
                                            }
                                        }
                                        _ = fi80
                                    }
                                case -1:
                                    v30 -= v30
                                default:
                                    v30 = v30
                                }
                            }
                        }
                        if true {
                            fmt.Println(v30)
                            if ((176 > 65535) && false) {
                                for fi90 := 0; fi90 < 18; fi90++ {
                                    var v91 uint16 = 256
                                    v92 := []int{fi90, fi90, 127}
                                    _ = fi90
                                    _ = v91
                                    _ = v92
                                }
                                if true {
                                    if false {
                                        var v93 uint32 = 42
                                        _ = v93
                                    } else {
                                        if (!true) {
                                            v94 := []int{}
                                            if true {
                                                if (82 > 0) {
                                                    fmt.Println(v31, v31)
                                                    for fi95 := 0; fi95 < 13; fi95++ {
                                                        if ((9223372036854775807 != -912) && (v31 != v31)) {
                                                            switch "" {
                                                            case "":
                                                                if len(v94) >= 2 {
                                                                    v94 = v94[1:len(v94)]
                                                                }
                                                            case "\n\t":
                                                                v96 := GenS1{F0: 42}
                                                                _ = v96
                                                            }
                                                        } else {
                                                            var v97 int8 = (-(v31))
                                                            v97 -= (v97 % (((v97 ^ v31)) | 1))
                                                            v97 -= v31
                                                            _ = v97
                                                        }
                                                        for fi98 := 0; fi98 < 1; fi98++ {
                                                            if len(v94) >= 2 {
                                                                v94 = v94[1:len(v94)]
                                                            }
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                fmt.Println(v31, fi95, v31)
                                                            case "hello":
                                                                v31 += 127
                                                            case "foo bar":
                                                                if ((32 > 1) && false) {
                                                                    var v99 uint8 = 1
                                                                    var v100 uint8 = v99
                                                                    _ = v100
                                                                    _ = v99
                                                                } else {
                                                                    v31 = (-(v31))
                                                                    v30 += float32(3.14159)
                                                                }
                                                            }
                                                            _ = fi98
                                                        }
                                                        _ = fi95
                                                    }
                                                    v101 := GenS5{F0: ""}
                                                    _ = v101
                                                } else {
                                                    var v102 string = "a"
                                                    if (!true) {
                                                        for fi103 := 0; fi103 < 20; fi103++ {
                                                            v31 = (v31 ^ v31)
                                                            v104 := GenS4{F0: 127}
                                                            _ = fi103
                                                            _ = v104
                                                        }
                                                        for fi105 := 0; fi105 < 19; fi105++ {
                                                            switch 9223372036854775807 {
                                                            case 42:
                                                                var v106 int = (^(fi105))
                                                                _ = v106
                                                            default:
                                                                fmt.Println(v30)
                                                            }
                                                            v102 = "hello"
                                                            _ = fi105
                                                        }
                                                    }
                                                    _ = v102
                                                }
                                            }
                                            _ = v94
                                        }
                                        var v107 int16 = 194
                                        _ = v107
                                    }
                                }
                            }
                        } else {
                            if (960 <= 15) {
                                v31 = -128
                                for fi108 := 0; fi108 < 20; fi108++ {
                                    if (("a" < "hello") || (fi108 > fi108)) {
                                        switch -2147483648 {
                                        case 2147483647:
                                            var v109 string = ((("\n\t" + ("hello" + ("hello" + ("" + "foo bar")))) + "") + "")
                                            _ = v109
                                        case 918:
                                            v30 += 1e10
                                        default:
                                            v31 -= (v31 >> ((255) & 63))
                                        }
                                        for fi110 := 0; fi110 < 8; fi110++ {
                                            var v111 uint64 = 4294967295
                                            var v112 uint32 = 956
                                            _ = fi110
                                            _ = v111
                                            _ = v112
                                        }
                                        v30 += (float32((-2.25 + float64(float32(3.14159)))) - v30)
                                    }
                                    _ = fi108
                                }
                            }
                            fmt.Println(v30, v31)
                        }
                    }
                    v31 = v31
                } else {
                    for fi113 := 0; fi113 < 14; fi113++ {
                        if true {
                            if (false || false) {
                                if (!true) {
                                    for fi114 := 0; fi114 < 20; fi114++ {
                                        v115 := []int{fi114, fi113, 885}
                                        v30 += v30
                                        _ = fi114
                                        _ = v115
                                    }
                                } else {
                                    switch -2147483648 {
                                    case 4294967295:
                                        fmt.Println(v31, fi113)
                                    case 1:
                                        if true {
                                            v116 := GenS3{F0: 42, F1: 65535}
                                            v117 := []int{(fi113 & fi113), fi113, fi113, fi113}
                                            _ = v116
                                            _ = v117
                                        } else {
                                            if ((float32(1.5) == v30) || (179 <= 746)) {
                                                if ((1 == 42) && false) {
                                                    var v118 uint16 = 878
                                                    fmt.Println(v31, v31, fi113)
                                                    _ = v118
                                                } else {
                                                    for fi119 := 0; fi119 < 9; fi119++ {
                                                        v30 = (v30 + v30)
                                                        _ = fi119
                                                    }
                                                    for fi120 := 0; fi120 < 4; fi120++ {
                                                        switch fi120 {
                                                        case 0:
                                                            v121 := []int{(-(fi120)), fi113}
                                                            _ = v121
                                                        case 128:
                                                            var v122 float64 = (-2.25 * 3.14159)
                                                            v123 := []int{(fi113 << ((751) & 63)), fi120, fi113, (fi120 * (fi120 << ((1) & 63)))}
                                                            var v124 string = ("a" + "\n\t")
                                                            _ = v122
                                                            _ = v123
                                                            _ = v124
                                                        default:
                                                            fmt.Println(v31, fi120, fi120)
                                                        }
                                                        _ = fi120
                                                    }
                                                    for fi125 := 0; fi125 < 17; fi125++ {
                                                        var v126 uint8 = 255
                                                        if ((1 >= -688) || (v30 < ((1e10 / float32(3.14159)) * 1.5))) {
                                                            if (18446744073709551615 <= 399) {
                                                                if ((v31 / (((v31 * v31)) | 1)) >= (v31 & 0)) {
                                                                    fmt.Println(v30, fi125, fi125)
                                                                    if ((fi125 - -433) >= fi125) {
                                                                        var v127 int32 = -1
                                                                        v126 += 140
                                                                        _ = v127
                                                                    }
                                                                } else {
                                                                    var v128 uint32 = 4294967295
                                                                    v129 := []int{}
                                                                    v129 = append(v129, (fi113 ^ (-(fi113))))
                                                                    _ = v128
                                                                    _ = v129
                                                                }
                                                            } else {
                                                                switch "" {
                                                                case "":
                                                                    v31 -= (v31 & v31)
                                                                    if (false || (-410 < -572)) {
                                                                        v130 := []int{4294967295}
                                                                        if len(v130) >= 2 {
                                                                            v130 = v130[1:len(v130)]
                                                                        }
                                                                        var v131 int32 = -1
                                                                        v132 := GenS6{F0: true, F1: 2147483647}
                                                                        _ = v130
                                                                        _ = v131
                                                                        _ = v132
                                                                    }
                                                                case "hello":
                                                                    v30 = v30
                                                                case "foo bar":
                                                                    var v133 int64 = 42
                                                                    _ = v133
                                                                }
                                                            }
                                                            v31 -= v31
                                                            fmt.Println(fi125)
                                                        } else {
                                                            switch (^(fi125)) {
                                                            case -9223372036854775808:
                                                                v31 += (-(v31))
                                                            case 232:
                                                                var v134 int16 = 0
                                                                _ = v134
                                                            }
                                                            v31 = (v31 % (((v31 - v31)) | 1))
                                                        }
                                                        switch (fi125 - (^(fi113))) {
                                                        case 42:
                                                            var v135 uint = 4294967295
                                                            _ = v135
                                                        default:
                                                            v136 := []int{(fi113 + 2147483647), fi113, fi125}
                                                            _ = v136
                                                        }
                                                        _ = fi125
                                                        _ = v126
                                                    }
                                                }
                                                for fi137 := 0; fi137 < 8; fi137++ {
                                                    fmt.Println(v30, v31, v30)
                                                    var v138 int64 = 0
                                                    var v139 int16 = -1
                                                    _ = fi137
                                                    _ = v138
                                                    _ = v139
                                                }
                                                if false {
                                                    for fi140 := 0; fi140 < 14; fi140++ {
                                                        var v141 uint64 = 2
                                                        v141 += (v141 ^ v141)
                                                        _ = fi140
                                                        _ = v141
                                                    }
                                                    v30 -= float32(((1.5 / -2.25) + 0.1))
                                                    fmt.Println(fi113)
                                                } else {
                                                    v142 := GenS3{F0: 1, F1: 65535}
                                                    fmt.Println(v31, v31)
                                                    var v143 float32 = v30
                                                    _ = v142
                                                    _ = v143
                                                }
                                            }
                                        }
                                    case 759:
                                        if (v31 >= v31) {
                                            fmt.Println(v31, fi113)
                                            v30 -= -2.25
                                        } else {
                                            v144 := []int{}
                                            _ = v144
                                        }
                                    }
                                }
                                for fi145 := 0; fi145 < 12; fi145++ {
                                    fmt.Println(fi113)
                                    _ = fi145
                                }
                                var v146 float32 = v30
                                _ = v146
                            } else {
                                if ((false && (!((("\n\t" + "a") < "") || (true || false)))) || (!(2147483647 > -501))) {
                                    var v147 int16 = 1000
                                    _ = v147
                                } else {
                                    if false {
                                        v148 := GenS4{F0: fi113}
                                        v149 := []int{fi113}
                                        _ = v148
                                        _ = v149
                                    }
                                    v150 := GenS1{F0: 215}
                                    _ = v150
                                }
                                var v151 float32 = 0.0
                                var v152 int64 = 9223372036854775807
                                _ = v151
                                _ = v152
                            }
                            v31 += (v31 & v31)
                        } else {
                            v31 = v31
                        }
                        if ((false || true) || false) {
                            switch "" {
                            case "":
                                switch "a" {
                                case "hello":
                                    fmt.Println(v30, v31, v31)
                                }
                            case "foo bar":
                                if (true || (1 == 2147483647)) {
                                    v153 := GenS2{F0: 32768, F1: (1 <= 65535)}
                                    v154 := []int{fi113, (fi113 << ((18446744073709551615) & 63))}
                                    _ = v153
                                    _ = v154
                                } else {
                                    if ((!false) && true) {
                                        var v155 int8 = 0
                                        _ = v155
                                    }
                                    v31 = v31
                                    if (("a" < "hello") && false) {
                                        var v156 int64 = -949
                                        var v157 int = 4294967295
                                        _ = v156
                                        _ = v157
                                    } else {
                                        if false {
                                            if false {
                                                v30 -= 1.5
                                                for fi158 := 0; fi158 < 17; fi158++ {
                                                    fmt.Println(fi158, v31, fi113)
                                                    v159 := []int{fi158}
                                                    _ = fi158
                                                    _ = v159
                                                }
                                            }
                                            var v160 uint32 = 42
                                            _ = v160
                                        }
                                    }
                                }
                            case "a":
                                var v161 uint16 = 344
                                _ = v161
                            }
                            v31 = v31
                            if ((9223372036854775807 <= 0) || false) {
                                v162 := []int{127}
                                switch "foo bar" {
                                case "\n\t":
                                    var v163 float32 = (v30 - (v30 + v30))
                                    _ = v163
                                case "":
                                    if (((155 > fi113) || ((("hello" >= "hello") || true) || (v31 > v31))) && (!(243 <= 42))) {
                                        var v164 string = ("\n\t" + "\n\t")
                                        var v165 uint64 = 675
                                        v166 := GenS4{F0: 128}
                                        _ = v164
                                        _ = v165
                                        _ = v166
                                    }
                                default:
                                    v167 := GenS2{F0: fi113, F1: ((-2.25 + float32((float64(v30) / float64(v30)))) < -2.25)}
                                    _ = v167
                                }
                                v162 = append(v162, (fi113 ^ 1))
                                _ = v162
                            }
                        } else {
                            if false {
                                v168 := []int{fi113, (fi113 * (fi113 ^ -1))}
                                _ = v168
                            }
                            for fi169 := 0; fi169 < 12; fi169++ {
                                if (true || (4294967295 == 2147483648)) {
                                    var v170 bool = ((0 > 0) || true)
                                    _ = v170
                                }
                                fmt.Println(fi169, fi169)
                                v171 := GenS6{F0: false, F1: 1}
                                _ = fi169
                                _ = v171
                            }
                        }
                        _ = fi113
                    }
                    var v172 bool = false
                    _ = v172
                }
            }
        }
        _ = v31
    } else {
        v30 -= v30
        var v173 float32 = (0.1 / v30)
        v173 -= (v30 * (1e10 + ((v30 - (1.5 + v30)) + (-2.25 + v173))))
        _ = v173
    }
    _ = v30
    return 128
}

func main() {
    v1 := []uint32{2147483647}
    var v2 float64 = (0.1 / 3.14159)
    var v3 bool = true
    v4 := make(map[int]bool)
    var v5 uint64 = 2
    for fi174 := 0; fi174 < 1; fi174++ {
        v175 := GenS1{F0: 2}
        _ = fi174
        _ = v175
    }
    v176 := GenS1{F0: 371}
    v5 = v5
    v177 := GenS1{F0: 18446744073709551615}
    if true {
        var v178 float64 = float64((0.1 / float32(v2)))
        var v179 int = 770
        _ = v178
        _ = v179
    } else {
        if v3 {
            for fi180 := 0; fi180 < 9; fi180++ {
                for fi181 := 0; fi181 < 1; fi181++ {
                    if (false || false) {
                        fmt.Println(fi181)
                    }
                    v5 -= 525
                    _ = fi181
                }
                v3 = v3
                var v182 float64 = v2
                _ = fi180
                _ = v182
            }
            for fi183 := 0; fi183 < 11; fi183++ {
                for fi184 := 0; fi184 < 4; fi184++ {
                    if v3 {
                        var v185 int64 = 9223372036854775807
                        _ = v185
                    } else {
                        switch -9223372036854775808 {
                        case 4294967295:
                            switch fi183 {
                            case -93:
                                v186 := GenS8{F0: -1}
                                _ = v186
                            case 1:
                                v3 = v3
                            case 2147483647:
                                v187 := v4[(fi184 + fi183)]
                                _ = v187
                            default:
                                switch (fi183 % ((128) | 1)) {
                                case 128:
                                    switch fi184 {
                                    case 32768:
                                        if (255 == 42) {
                                            var v188 string = "\n\t"
                                            if len(v1) > 0 {
                                                v1[0] = 2
                                            }
                                            var v189 float64 = v2
                                            _ = v188
                                            _ = v189
                                        }
                                        var v190 float64 = v2
                                        v191 := GenS9{F0: 646, F1: 1000}
                                        _ = v190
                                        _ = v191
                                    default:
                                        if (fi183 != fi183) {
                                            v3 = (v5 > (v5 / (((v5 % ((v5) | 1))) | 1)))
                                        } else {
                                            if (-120 == 0) {
                                                _ = len(v4)
                                                if false {
                                                    var v192 uint32 = 0
                                                    delete(v4, fi184)
                                                    switch 713 {
                                                    case 32767:
                                                        v5 = 0
                                                    default:
                                                        v5 += v5
                                                    }
                                                    _ = v192
                                                }
                                            }
                                            var v193 uint32 = 65535
                                            _ = v193
                                        }
                                    }
                                case -2147483648:
                                    var v194 int = int(v5)
                                    _ = v194
                                }
                                if (!v3) {
                                    v195 := GenS9{F0: 42, F1: -409}
                                    fmt.Println(fi183, v2, v2)
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = v195
                                } else {
                                    var v196 int = -932
                                    _ = len(v4)
                                    _ = v196
                                }
                            }
                        default:
                            if (!v3) {
                                v197 := GenS4{F0: 32767}
                                _ = v197
                            }
                        }
                    }
                    _ = fi184
                }
                fmt.Println(fi183, fi183, v2)
                _ = fi183
            }
            v198 := GenS10{F0: 4294967295, F1: 1, F2: -1}
            _ = v198
        } else {
            v199 := GenS6{F0: v3, F1: 1}
            for fi200 := 0; fi200 < 2; fi200++ {
                v5 += 42
                _ = fi200
            }
            _ = v199
        }
        v5 = 42
        for fi201 := 0; fi201 < 16; fi201++ {
            var v202 uint32 = 4294967295
            fmt.Println(v3, fi201)
            _ = fi201
            _ = v202
        }
    }
    for fi203 := 0; fi203 < 18; fi203++ {
        var v204 uint = 255
        for fi205 := 0; fi205 < 9; fi205++ {
            v206 := GenS11{F0: (uint16(v204) == 666)}
            _ = fi205
            _ = v206
        }
        var v207 string = ("" + (("a" + "") + "a"))
        _ = fi203
        _ = v204
        _ = v207
    }
    var v208 int16 = -842
    if (!v3) {
        for fi209 := 0; fi209 < 3; fi209++ {
            v2 = (0.0 / v2)
            v210 := GenS12{F0: 0, F1: true, F2: 1}
            _ = fi209
            _ = v210
        }
    } else {
        for fi211 := 0; fi211 < 6; fi211++ {
            var v212 int8 = 127
            _ = fi211
            _ = v212
        }
        if v3 {
            if (!(1 != 1)) {
                var v213 uint = 0
                if (v5 < v5) {
                    var v214 int32 = 1
                    v2 += v2
                    _ = v214
                } else {
                    v2 = float64(0.1)
                    var v215 uint16 = 42
                    fmt.Println(v208, v2)
                    _ = v215
                }
                _ = v213
            } else {
                if ("a" <= "") {
                    var v216 uint8 = 1
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v2 -= (v2 / float64(float32(v2)))
                    _ = v216
                } else {
                    fmt.Println(v5, v3, v2)
                    for fi217 := 0; fi217 < 13; fi217++ {
                        v3 = v3
                        _ = fi217
                    }
                    v218 := GenS13{F0: "\n\t", F1: 42, F2: -950}
                    _ = v218
                }
                for fi219 := 0; fi219 < 10; fi219++ {
                    for fi220 := 0; fi220 < 7; fi220++ {
                        var v221 bool = (0 != 717)
                        _ = fi220
                        _ = v221
                    }
                    _ = fi219
                }
                v222 := GenS12{F0: -1, F1: v3, F2: 39}
                _ = v222
            }
            var v223 int32 = 0
            fmt.Println(v208)
            _ = v223
        }
        switch 32767 {
        case 4294967295:
            if ((!(128 != 1)) && v3) {
                for fi224 := 0; fi224 < 19; fi224++ {
                    if (-128 < 100) {
                        var v225 float64 = v2
                        _ = v225
                    }
                    v226 := GenS5{F0: (("hello" + "a") + (("foo bar" + ("a" + "")) + ("foo bar" + "hello")))}
                    _ = fi224
                    _ = v226
                }
                var v227 float64 = (float64((float32(v2) / (1.5 + 1.5))) + -2.25)
                var v228 uint64 = (v5 & v5)
                _ = v227
                _ = v228
            }
        case 928:
            for fi229 := 0; fi229 < 7; fi229++ {
                v2 = (-2.25 * (((1.5 * 1.5) + (v2 / float64(3.14159))) - (float64(0.0) / ((v2 - 1.5) / (1e10 / 3.14159)))))
                fmt.Println(v2)
                v230 := GenS14{F0: (v208 * 32767), F1: true}
                _ = fi229
                _ = v230
            }
        case -807:
            var v231 int32 = -2147483648
            _ = v231
        default:
            var v232 uint16 = 1
            _ = v232
        }
    }
    fmt.Println(v5)
    v2 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v208:", v208)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v176
    _ = v177
    _ = v2
    _ = v208
    _ = v3
    _ = v4
    _ = v5
}
