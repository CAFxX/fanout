package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 int; F1 int; F2 int }
type GenS5 struct { F0 uint16; F1 int64; F2 uint }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 uint32; F1 uint }
type GenS8 struct { F0 string; F1 bool }
type GenS9 struct { F0 string; F1 int32 }
type GenS10 struct { F0 int32; F1 string }
type GenS11 struct { F0 int; F1 int8 }
type GenS12 struct { F0 int16; F1 bool; F2 int64 }
type GenS13 struct { F0 int; F1 int32; F2 uint8 }
type GenS14 struct { F0 uint64; F1 int16 }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 int64; F1 int32 }
type GenS17 struct { F0 uint64 }
type GenS18 struct { F0 uint32 }
type GenS19 struct { F0 int8; F1 int16 }
func genFn0() int {
    if ((false || true) || (true || (4294967295 == 1))) {
        switch (("hello" + "") + "hello") {
        case "foo bar":
            for fi7 := 0; fi7 < 11; fi7++ {
                var v8 float64 = 1e10
                if true {
                    for fi9 := 0; fi9 < 12; fi9++ {
                        var v10 int8 = 100
                        v8 -= v8
                        _ = fi9
                        _ = v10
                    }
                } else {
                    if (!(-9223372036854775808 < -9223372036854775808)) {
                        v8 = 0.0
                        fmt.Println(fi7, v8)
                        v8 = (v8 + v8)
                    } else {
                        var v11 float32 = 3.14159
                        _ = v11
                    }
                }
                v12 := GenS1{F0: 130, F1: 94}
                _ = fi7
                _ = v12
                _ = v8
            }
        case "hello":
            var v13 bool = (-2.25 <= (float64(0.1) * 1.5))
            _ = v13
        default:
            v14 := []int{0, -838, -436}
            _ = v14
        }
    } else {
        v15 := []int{-9223372036854775808}
        v16 := GenS1{F0: 255, F1: -1}
        _ = v15
        _ = v16
    }
    var v17 uint8 = 1
    switch "foo bar" {
    case "foo bar":
        if (!true) {
            v18 := []int{}
            var v19 uint = 244
            v19 = v19
            _ = v18
            _ = v19
        }
    default:
        if false {
            var v20 uint8 = v17
            if true {
                v21 := GenS1{F0: v20, F1: 28}
                switch 32768 {
                case 128:
                    if (693 > uint16(v20)) {
                        if ((362 > 2) || (-1 == -9223372036854775808)) {
                            var v22 uint8 = (v17 << ((0) & 63))
                            _ = v22
                        }
                    } else {
                        for fi23 := 0; fi23 < 8; fi23++ {
                            v24 := GenS1{F0: (v17 << ((2) & 63)), F1: -128}
                            v25 := []int{(fi23 & fi23), (fi23 * (fi23 ^ fi23)), (fi23 & (fi23 / (((fi23 * 2147483647)) | 1))), 9223372036854775807}
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                        v26 := GenS1{F0: 42, F1: 100}
                        _ = v26
                    }
                case 330:
                    var v27 uint8 = 42
                    _ = v27
                default:
                    if ((!(true && (255 < 16))) && false) {
                        switch ("a" + "hello") {
                        case "hello":
                            if (true && false) {
                                v20 += (v20 >> ((18446744073709551615) & 63))
                            } else {
                                v28 := GenS1{F0: (v17 << ((4294967295) & 63)), F1: 0}
                                for fi29 := 0; fi29 < 16; fi29++ {
                                    if (-626 >= 9223372036854775807) {
                                        var v30 uint = 348
                                        v17 -= (^(v17))
                                        _ = v30
                                    } else {
                                        v31 := GenS1{F0: 42, F1: 0}
                                        _ = v31
                                    }
                                    var v32 int64 = 0
                                    if false {
                                        v33 := []int{9223372036854775807, fi29, 32767}
                                        var v34 string = ("hello" + "")
                                        _ = v33
                                        _ = v34
                                    }
                                    _ = fi29
                                    _ = v32
                                }
                                v35 := GenS2{F0: 330, F1: uint(v17)}
                                _ = v28
                                _ = v35
                            }
                        case "a":
                            v17 = 128
                        default:
                            v36 := []int{}
                            _ = v36
                        }
                    }
                }
                _ = v21
            }
            v37 := []int{128, 715, 42}
            _ = v20
            _ = v37
        } else {
            v17 += (v17 ^ v17)
            for fi38 := 0; fi38 < 19; fi38++ {
                v39 := []int{fi38, -2147483648, (-(fi38)), fi38}
                var v40 int16 = 42
                var v41 int = fi38
                _ = fi38
                _ = v39
                _ = v40
                _ = v41
            }
        }
    }
    var v42 int8 = 122
    for fi43 := 0; fi43 < 9; fi43++ {
        if ((v42 - -1) <= v42) {
            for fi44 := 0; fi44 < 10; fi44++ {
                v45 := []int{fi43, fi44, (^(fi43))}
                _ = fi44
                _ = v45
            }
        }
        _ = fi43
    }
    _ = v17
    _ = v42
    return 42
}

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = 1.5
    var v3 uint64 = 869
    var v4 uint32 = 493
    var v5 uint64 = v3
    v6 := []uint32{(v4 | (v4 >> ((244) & 63))), (v4 & 926), 2147483648, (v4 - v4)}
    fmt.Println(v3)
    v6 = append(v6, 908)
    for fi46 := 0; fi46 < 1; fi46++ {
        for fi47 := 0; fi47 < 15; fi47++ {
            switch "foo bar" {
            case "\n\t":
                if false {
                    delete(v1, fi47)
                    v48 := GenS3{F0: fi47, F1: genFn0()}
                    _ = v48
                } else {
                    v49 := GenS2{F0: 32767, F1: 2}
                    switch fi47 {
                    case -880:
                        delete(v1, -1)
                    }
                    v5 -= 0
                    _ = v49
                }
                if (!false) {
                    fmt.Println(v4)
                    v5 -= (v3 / ((v3) | 1))
                    var v50 float32 = 1e10
                    _ = v50
                } else {
                    v4 -= (^(v4))
                    v51 := GenS3{F0: genFn0(), F1: (fi47 * -747)}
                    var v52 int32 = 42
                    v53 := GenS4{F0: (fi47 | (fi47 - fi46)), F1: fi46, F2: genFn0()}
                    var v54 float64 = ((0.0 - 0.1) - v2)
                    var v55 int64 = 9223372036854775807
                    v56 := GenS5{F0: 42, F1: (v55 / ((v55) | 1)), F2: 58}
                    switch fi46 {
                    case -9223372036854775808:
                        switch ("" + "foo bar") {
                        case "":
                            var v57 string = ("hello" + "\n\t")
                            _ = v57
                        case "a":
                            if false {
                                if (!((2 < 187) && ("\n\t" >= "foo bar"))) {
                                    v58 := GenS2{F0: 42, F1: 18446744073709551615}
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    _ = v58
                                }
                                if (42 == uint(v5)) {
                                    var v59 float64 = 3.14159
                                    v6 = append(v6, 0)
                                    v5 -= v3
                                    _ = v59
                                } else {
                                    v2 += v2
                                }
                                v60 := GenS4{F0: (-(fi47)), F1: 4294967295, F2: fi47}
                                _ = v60
                            } else {
                                var v61 int = (fi46 + fi46)
                                _ = v61
                            }
                        case "foo bar":
                            var v62 int = int(v52)
                            _ = v62
                        default:
                            if (2 != 2) {
                                v54 = v2
                                v55 -= 1
                                if false {
                                    switch "hello" {
                                    case "":
                                        var v63 float32 = -2.25
                                        _ = v63
                                    }
                                } else {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    if (1.5 != 3.14159) {
                                        v64 := GenS3{F0: (fi47 ^ (^(fi47))), F1: fi47}
                                        switch (fi47 | fi46) {
                                        case 127:
                                            var v65 float32 = -2.25
                                            _ = v65
                                        case 42:
                                            var v66 uint16 = 292
                                            _ = v66
                                        }
                                        _ = v64
                                    }
                                    v2 -= (float64(1.5) - float64(0.0))
                                    switch "" {
                                    case "":
                                        if (2 <= 650) {
                                            v55 -= 9223372036854775807
                                            var v67 uint8 = 255
                                            v54 = v2
                                            var v68 float32 = 0.0
                                            v52 = 1
                                            var v69 int8 = 127
                                            fmt.Println(fi47)
                                            _ = v67
                                            _ = v68
                                            _ = v69
                                        } else {
                                            var v70 uint32 = v4
                                            var v71 uint32 = 127
                                            delete(v1, fi46)
                                            _ = v70
                                            _ = v71
                                        }
                                    }
                                    if (!(v5 != v5)) {
                                        switch ("foo bar" + "") {
                                        case "a":
                                            var v72 uint16 = 473
                                            _ = v72
                                        case "hello":
                                            if ((v3 | (^(v3))) >= (v3 ^ (v5 >> ((533) & 63)))) {
                                                switch "hello" {
                                                case "hello":
                                                    if false {
                                                        v73 := GenS6{F0: fi47, F1: 0}
                                                        _ = v73
                                                    } else {
                                                        var v74 uint = 388
                                                        v5 += (v5 - v5)
                                                        _ = v74
                                                    }
                                                case "a":
                                                    v5 = v5
                                                    fmt.Println(v55)
                                                case "":
                                                    var v75 uint = 941
                                                    _ = v75
                                                default:
                                                    if (2 != 2) {
                                                        v76 := GenS7{F0: v4, F1: 4294967295}
                                                        switch "foo bar" {
                                                        case "hello":
                                                            var v77 uint = 369
                                                            _ = v77
                                                        case "foo bar":
                                                            fmt.Println(v4)
                                                        default:
                                                            if true {
                                                                v78 := GenS3{F0: (fi47 >> ((455) & 63)), F1: fi46}
                                                                v79 := GenS8{F0: "foo bar", F1: ((1 != 420) || (v5 <= v3))}
                                                                var v80 uint32 = (v4 >> ((2) & 63))
                                                                v81 := GenS2{F0: 598, F1: 629}
                                                                if (!(false && false)) {
                                                                    var v82 uint16 = 767
                                                                    v52 += v52
                                                                    _ = v82
                                                                } else {
                                                                    switch "a" {
                                                                    case "":
                                                                        switch "a" {
                                                                        case "foo bar":
                                                                            v6 = append(v6, (v4 ^ (^(v4))))
                                                                        case "\n\t":
                                                                            v2 -= v2
                                                                        default:
                                                                            var v83 uint32 = 65535
                                                                            if false {
                                                                                fmt.Println(fi46)
                                                                                v55 += v55
                                                                            } else {
                                                                                switch "a" {
                                                                                case "foo bar":
                                                                                    v84 := GenS5{F0: 42, F1: v55, F2: 27}
                                                                                    _ = v84
                                                                                default:
                                                                                    var v85 int = 127
                                                                                    _ = v85
                                                                                }
                                                                                var v86 uint16 = 156
                                                                                _ = v86
                                                                            }
                                                                            v83 = (v83 % ((v83) | 1))
                                                                            fmt.Println(v3, v83, v83)
                                                                            if len(v6) >= 2 {
                                                                                v6 = v6[1:len(v6)]
                                                                            }
                                                                            _ = v83
                                                                        }
                                                                    default:
                                                                        var v87 string = ("a" + ("foo bar" + ((("\n\t" + "foo bar") + "a") + "\n\t")))
                                                                        _ = v87
                                                                    }
                                                                }
                                                                v1[fi46] = true
                                                                _ = v78
                                                                _ = v79
                                                                _ = v80
                                                                _ = v81
                                                            }
                                                        }
                                                        _ = v76
                                                    } else {
                                                        v88 := GenS9{F0: ("hello" + "foo bar"), F1: (v52 | -345)}
                                                        v89 := GenS4{F0: 876, F1: (fi46 + (fi46 << ((42) & 63))), F2: fi46}
                                                        v90 := v1[(fi47 * (^(fi47)))]
                                                        if ((true || false) && (9223372036854775807 != v55)) {
                                                            switch genFn0() {
                                                            case 4294967295:
                                                                if v90 {
                                                                    switch 542 {
                                                                    case 9223372036854775807:
                                                                        var v91 uint16 = 2
                                                                        _ = v91
                                                                    case 230:
                                                                        var v92 int8 = 35
                                                                        _ = v92
                                                                    case 2147483647:
                                                                        v5 -= v3
                                                                    default:
                                                                        delete(v1, -777)
                                                                    }
                                                                    v1[(fi46 >> ((42) & 63))] = v90
                                                                }
                                                            case 11:
                                                                var v93 uint64 = (v3 - v3)
                                                                _ = v93
                                                            default:
                                                                if len(v6) > 0 {
                                                                    v6[0] = (v4 & 35)
                                                                }
                                                            }
                                                            if true {
                                                                v94 := GenS7{F0: 0, F1: 42}
                                                                v95 := GenS3{F0: fi46, F1: (fi46 * (fi46 / ((genFn0()) | 1)))}
                                                                _ = v94
                                                                _ = v95
                                                            }
                                                            v96 := GenS1{F0: 149, F1: 1}
                                                            _ = v96
                                                        } else {
                                                            var v97 uint8 = 2
                                                            v97 = (^(v97))
                                                            _ = v97
                                                        }
                                                        if ((v52 & v52) != v52) {
                                                            if v90 {
                                                                var v98 float32 = -2.25
                                                                _ = v98
                                                            } else {
                                                                if false {
                                                                    var v99 int = (^(fi47))
                                                                    _ = v99
                                                                }
                                                            }
                                                            if v90 {
                                                                v5 = 781
                                                                var v100 uint8 = 1
                                                                var v101 int64 = v55
                                                                _ = v100
                                                                _ = v101
                                                            }
                                                            var v102 int64 = v55
                                                            _ = v102
                                                        } else {
                                                            v5 -= (v3 >> ((1) & 63))
                                                            v103 := v1[genFn0()]
                                                            var v104 float32 = 0.1
                                                            _ = v103
                                                            _ = v104
                                                        }
                                                        var v105 bool = true
                                                        _ = v105
                                                        _ = v88
                                                        _ = v89
                                                        _ = v90
                                                    }
                                                }
                                                switch "foo bar" {
                                                case "":
                                                    fmt.Println(v3, v54)
                                                default:
                                                    var v106 int8 = -1
                                                    _ = v106
                                                }
                                                v4 += (v4 - (v4 | 0))
                                            }
                                            fmt.Println(v4, v5)
                                        case "":
                                            fmt.Println(v2, v55, v2)
                                        default:
                                            v55 = (v55 + v55)
                                        }
                                        v107 := GenS9{F0: "foo bar", F1: v52}
                                        switch "" {
                                        case "":
                                            v55 = v55
                                        }
                                        _ = v107
                                    }
                                    var v108 int = (fi47 % (((fi47 << ((2) & 63))) | 1))
                                    v109 := GenS5{F0: uint16(v5), F1: 23, F2: 14}
                                    _ = v108
                                    _ = v109
                                }
                            }
                        }
                    case -506:
                        v1[(fi46 + 42)] = (822 == 255)
                    case 773:
                        v4 = (v4 >> ((0) & 63))
                        if (!true) {
                            if (!(2147483647 < v52)) {
                                v2 = 0.0
                                if len(v6) > 0 {
                                    v6[0] = 2147483648
                                }
                            } else {
                                if (!(482 <= 138)) {
                                    v3 += (v5 * (v3 / ((4294967295) | 1)))
                                    switch "" {
                                    case "":
                                        fmt.Println(v4, v4)
                                    }
                                    var v110 int64 = -291
                                    fmt.Println(v2, v3, v55)
                                    _ = v110
                                } else {
                                    v54 += float64(1e10)
                                    if false {
                                        v55 += 9223372036854775807
                                        v111 := GenS9{F0: "a", F1: v52}
                                        if false {
                                            switch "hello" {
                                            case "foo bar":
                                                v2 = (v54 / (v2 / (v2 - v54)))
                                            case "hello":
                                                v112 := GenS10{F0: int32(v4), F1: "\n\t"}
                                                var v113 uint8 = 51
                                                v114 := v1[fi47]
                                                _ = v112
                                                _ = v113
                                                _ = v114
                                            case "\n\t":
                                                v4 += v4
                                                var v115 int16 = 3
                                                _ = v115
                                            }
                                            var v116 float64 = ((0.0 + (v2 + 3.14159)) + (v2 + ((v2 / (v54 / v2)) + 1e10)))
                                            var v117 int64 = (v55 ^ v55)
                                            _ = v116
                                            _ = v117
                                        }
                                        _ = v111
                                    }
                                    var v118 int16 = 880
                                    if true {
                                        var v119 int8 = 127
                                        _ = v119
                                    }
                                    _ = v118
                                }
                                if len(v6) > 0 {
                                    v6[0] = (v4 * v4)
                                }
                            }
                            v4 -= v4
                        } else {
                            v55 = (v55 - (v55 * (v55 << ((1) & 63))))
                            fmt.Println(v3, fi47, v4)
                            var v120 bool = true
                            _ = v120
                        }
                        if false {
                            var v121 bool = (-2.25 >= ((1e10 * (0.0 + (1.5 * 0.0))) - 3.14159))
                            _ = v121
                        } else {
                            if true {
                                v4 = (v4 * (v4 - (v4 ^ 371)))
                                var v122 uint16 = 65535
                                var v123 float32 = (0.1 * 1.5)
                                _ = v122
                                _ = v123
                            } else {
                                fmt.Println(v54)
                                switch (fi47 * fi47) {
                                case 1:
                                    v124 := GenS11{F0: (fi47 + genFn0()), F1: -113}
                                    _ = v124
                                case 42:
                                    if false {
                                        v5 -= (v3 << ((487) & 63))
                                    } else {
                                        var v125 string = ""
                                        _ = v125
                                    }
                                case 32768:
                                    switch genFn0() {
                                    case 2147483647:
                                        switch "" {
                                        case "foo bar":
                                            v52 = v52
                                        case "":
                                            _ = len(v1)
                                        default:
                                            v126 := GenS3{F0: (fi46 % ((genFn0()) | 1)), F1: (^(fi46))}
                                            _ = v126
                                        }
                                    default:
                                        var v127 uint16 = 65535
                                        _ = v127
                                    }
                                }
                                v2 = v54
                            }
                            var v128 string = "a"
                            fmt.Println(v2, v2)
                            _ = v128
                        }
                    default:
                        v129 := GenS11{F0: (fi47 >> ((4294967295) & 63)), F1: 114}
                        _ = v129
                    }
                    if (false || (v54 == (v2 * (float64(1e10) - v54)))) {
                        var v130 uint16 = 203
                        delete(v1, 0)
                        _ = v130
                    } else {
                        v131 := GenS4{F0: (fi47 | genFn0()), F1: fi47, F2: (fi46 ^ genFn0())}
                        var v132 int16 = 32767
                        _ = v131
                        _ = v132
                    }
                    if false {
                        var v133 uint = 42
                        if (v3 == v3) {
                            v6 = append(v6, (v4 >> ((65535) & 63)))
                            v5 += v5
                            switch "" {
                            case "hello":
                                v6 = append(v6, (v4 & v4))
                            default:
                                switch (fi46 << ((v133) & 63)) {
                                case -2147483648:
                                    v1[-1] = false
                                }
                            }
                            _ = len(v6)
                        }
                        _ = v133
                    }
                    _ = v51
                    _ = v52
                    _ = v53
                    _ = v54
                    _ = v55
                    _ = v56
                }
            case "hello":
                var v134 float64 = -2.25
                _ = v134
            case "foo bar":
                v4 += 0
            default:
                switch (fi46 & (fi47 << ((715) & 63))) {
                case 2147483647:
                    v4 = 173
                case 264:
                    var v135 uint32 = v4
                    _ = v135
                case 128:
                    if ((v4 | (v4 >> ((1) & 63))) < v4) {
                        delete(v1, fi46)
                    } else {
                        v136 := GenS12{F0: 1, F1: true, F2: 1}
                        switch genFn0() {
                        case 42:
                            var v137 uint32 = 65535
                            _ = v137
                        case 127:
                            v138 := GenS4{F0: fi47, F1: fi46, F2: (^(fi47))}
                            _ = v138
                        }
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        if true {
                            var v139 string = "a"
                            v140 := GenS9{F0: v139, F1: -2147483648}
                            _ = v139
                            _ = v140
                        }
                        _ = v136
                    }
                default:
                    v1[genFn0()] = false
                }
            }
            var v141 int64 = 1
            switch (fi47 & (fi46 ^ genFn0())) {
            case 432:
                v3 -= 2
            case 701:
                switch (fi47 & 4294967295) {
                case 127:
                    v5 -= v5
                default:
                    var v142 float64 = v2
                    _ = v142
                }
                v143 := v1[(^(fi46))]
                _ = len(v1)
                if (v143 || v143) {
                    v144 := GenS10{F0: 2147483647, F1: "\n\t"}
                    v145 := GenS10{F0: -2147483648, F1: ("" + "\n\t")}
                    v146 := GenS13{F0: (fi46 >> ((561) & 63)), F1: 42, F2: 1}
                    var v147 float32 = 1e10
                    _ = v144
                    _ = v145
                    _ = v146
                    _ = v147
                } else {
                    var v148 float32 = ((3.14159 + 3.14159) - (0.1 - (0.1 / 1.0)))
                    if (1 != 255) {
                        var v149 uint8 = 0
                        v149 -= (v149 + (v149 >> ((0) & 63)))
                        _ = v149
                    } else {
                        v148 = v148
                        v4 -= 450
                        fmt.Println(v148, v148)
                    }
                    _ = v148
                }
                _ = v143
            case 32767:
                var v150 uint8 = uint8(v4)
                _ = v150
            default:
                v151 := GenS14{F0: (v3 & (v3 << ((65535) & 63))), F1: 0}
                _ = v151
            }
            _ = fi47
            _ = v141
        }
        for fi152 := 0; fi152 < 11; fi152++ {
            var v153 uint32 = 65535
            v1[0] = (-702 >= -466)
            v154 := GenS10{F0: 42, F1: "hello"}
            _ = fi152
            _ = v153
            _ = v154
        }
        _ = fi46
    }
    var v155 float64 = v2
    if (!true) {
        if true {
            if true {
                if false {
                    v156 := GenS15{F0: 65535}
                    fmt.Println(v155)
                    _ = v156
                }
            } else {
                v3 = 42
                switch (("foo bar" + "\n\t") + (("hello" + (("foo bar" + "hello") + "")) + "")) {
                case "\n\t":
                    for fi157 := 0; fi157 < 12; fi157++ {
                        var v158 uint32 = (^(v4))
                        var v159 uint64 = (v5 % ((v3) | 1))
                        if (!true) {
                            var v160 uint16 = 42
                            _ = v160
                        } else {
                            if true {
                                fmt.Println(fi157, v2)
                                switch fi157 {
                                case -9223372036854775808:
                                    v4 -= v158
                                default:
                                    switch "hello" {
                                    case "hello":
                                        var v161 int = (fi157 + fi157)
                                        _ = v161
                                    default:
                                        v162 := GenS16{F0: 42, F1: 853}
                                        _ = v162
                                    }
                                }
                            }
                            v6 = append(v6, v158)
                            v163 := GenS1{F0: 42, F1: 60}
                            _ = v163
                        }
                        _ = fi157
                        _ = v158
                        _ = v159
                    }
                default:
                    var v164 int8 = 118
                    _ = v164
                }
            }
        } else {
            if len(v6) > 0 {
                v6[0] = v4
            }
            var v165 string = ("hello" + "hello")
            _ = v165
        }
        v1[-2147483648] = (18446744073709551615 != 18446744073709551615)
    } else {
        if (1 >= 42) {
            if true {
                switch 21 {
                case 2147483647:
                    var v166 int32 = 1
                    _ = v166
                case 32768:
                    v167 := GenS2{F0: 32767, F1: 255}
                    _ = v167
                case -272:
                    for fi168 := 0; fi168 < 7; fi168++ {
                        v169 := GenS2{F0: 32767, F1: 4294967295}
                        _ = fi168
                        _ = v169
                    }
                default:
                    var v170 string = ("" + "\n\t")
                    _ = v170
                }
                var v171 uint8 = 80
                _ = v171
            } else {
                fmt.Println(v4)
                var v172 string = ""
                _ = v172
            }
            if (-645 == -103) {
                var v173 uint32 = 2147483647
                v4 += (^(v173))
                var v174 float32 = (0.1 + 3.14159)
                _ = v173
                _ = v174
            }
            delete(v1, 0)
        } else {
            for fi175 := 0; fi175 < 4; fi175++ {
                var v176 int16 = -361
                v177 := GenS4{F0: (fi175 & 2147483647), F1: fi175, F2: 9223372036854775807}
                _ = fi175
                _ = v176
                _ = v177
            }
        }
        v155 += v2
    }
    for fi178 := 0; fi178 < 9; fi178++ {
        v5 -= 2
        v155 += v2
        switch (("hello" + "foo bar") + "") {
        case "hello":
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        case "\n\t":
            v179 := GenS17{F0: v5}
            _ = v179
        case "":
            if (999 > (v5 % ((688) | 1))) {
                var v180 int32 = 42
                v4 -= (v4 ^ 338)
                v3 += (^(v3))
                _ = v180
            }
        }
        _ = fi178
    }
    fmt.Println(v3, v155)
    var v181 int = -9223372036854775808
    var v182 bool = (42 != (v3 << ((502) & 63)))
    for fi183 := 0; fi183 < 19; fi183++ {
        v181 -= 128
        var v184 bool = false
        for fi185 := 0; fi185 < 20; fi185++ {
            v186 := GenS12{F0: -32768, F1: (2 < 127), F2: -1}
            v182 = ((v5 | (v5 / (((v3 % ((v3) | 1))) | 1))) != v5)
            _ = fi185
            _ = v186
        }
        _ = fi183
        _ = v184
    }
    var v187 int = (v181 - (-(v181)))
    if true {
        var v188 uint = 42
        _ = v188
    } else {
        _ = len(v1)
        v1[genFn0()] = false
        if v182 {
            v189 := GenS18{F0: 299}
            v190 := GenS19{F0: 95, F1: -227}
            _ = v189
            _ = v190
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v155:", v155)
    fmt.Println("v181:", v181)
    fmt.Println("v182:", v182)
    fmt.Println("v187:", v187)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v155
    _ = v181
    _ = v182
    _ = v187
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
