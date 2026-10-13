package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 string; F2 uint }
type GenS5 struct { F0 uint16; F1 uint64; F2 string }
type GenS6 struct { F0 string; F1 int16; F2 uint64 }
type GenS7 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS8 struct { F0 uint8; F1 int8 }
type GenS9 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS10 struct { F0 int64; F1 int }
type GenS11 struct { F0 uint32; F1 int8 }
type GenS12 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS13 struct { F0 int; F1 uint32 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 uint64; F1 string; F2 uint }
func genFn0(p0 uint32, p1 int) int16 {
    fmt.Println(p0, p0)
    for fi4 := 0; fi4 < 15; fi4++ {
        fmt.Println(fi4)
        if false {
            var v5 int16 = 32767
            var v6 float32 = 0.1
            _ = v5
            _ = v6
        } else {
            var v7 int64 = 134
            p0 -= 182
            switch ("hello" + "a") {
            case "":
                v8 := []int{(p1 + 128), (p1 % (((fi4 * 4294967295)) | 1)), fi4, (p1 + p1)}
                _ = v8
            }
            _ = v7
        }
        v9 := GenS1{F0: 255}
        _ = fi4
        _ = v9
    }
    v10 := []int{p1, -1}
    for fi11 := 0; fi11 < 1; fi11++ {
        for fi12 := 0; fi12 < 16; fi12++ {
            if false {
                if (((false && (((255 < 587) || (!false)) || (!(false || false)))) && ((fi12 < fi12) || false)) || true) {
                    switch "foo bar" {
                    case "\n\t":
                        fmt.Println(fi11, fi11, p0)
                    case "":
                        var v13 bool = false
                        _ = v13
                    case "foo bar":
                        if (42 < 9223372036854775807) {
                            var v14 float32 = 1e10
                            var v15 uint32 = (p0 - p0)
                            _ = v14
                            _ = v15
                        }
                    default:
                        p0 = (^(p0))
                    }
                    v16 := GenS1{F0: 0}
                    if (2 <= 4294967295) {
                        var v17 uint16 = 0
                        var v18 float32 = 0.1
                        v19 := GenS2{F0: 1, F1: 1}
                        var v20 bool = true
                        v10 = append(v10, fi12)
                        _ = v17
                        _ = v18
                        _ = v19
                        _ = v20
                    } else {
                        var v21 int = p1
                        v21 -= 42
                        _ = v21
                    }
                    _ = v16
                }
                var v22 int = (^(fi12))
                _ = v22
            }
            _ = fi12
        }
        if ((0.0 - 0.0) < 1e10) {
            p1 = p1
            v23 := GenS2{F0: 0, F1: -824}
            _ = v23
        } else {
            v10 = append(v10, 32768)
            for fi24 := 0; fi24 < 3; fi24++ {
                switch (p1 | fi11) {
                case 252:
                    if len(v10) > 0 {
                        v10[0] = fi11
                    }
                }
                _ = fi24
            }
        }
        _ = fi11
    }
    if (true && (9223372036854775807 < 1)) {
        p0 = 42
        p1 = 0
    } else {
        var v25 int64 = -671
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        var v26 uint = 4294967295
        _ = v25
        _ = v26
    }
    _ = p0
    _ = p1
    _ = v10
    return 0
}

func main() {
    var v1 uint16 = 0
    v2 := make(map[int]int)
    var v3 int = -2147483648
    if false {
        v1 += (v1 / ((65535) | 1))
        switch (("" + "foo bar") + "a") {
        case "\n\t":
            switch "\n\t" {
            case "foo bar":
                for fi27 := 0; fi27 < 9; fi27++ {
                    var v28 float64 = ((((1e10 - 3.14159) + (float64(1.5) + 1e10)) * (1.5 / 1.5)) / (1.5 * 1e10))
                    _ = fi27
                    _ = v28
                }
            case "hello":
                for fi29 := 0; fi29 < 18; fi29++ {
                    var v30 int32 = 0
                    _ = fi29
                    _ = v30
                }
            }
        case "a":
            delete(v2, (v3 << ((2) & 63)))
        default:
            if (0 >= 133) {
                v31 := GenS3{F0: -1}
                fmt.Println(v1, v3)
                _ = v31
            }
        }
    } else {
        fmt.Println(v1, v3)
        for fi32 := 0; fi32 < 12; fi32++ {
            v3 = 814
            _ = fi32
        }
    }
    v33 := GenS1{F0: uint(v1)}
    v1 -= (v1 >> ((42) & 63))
    v1 += (v1 / (((v1 % (((v1 + 503)) | 1))) | 1))
    for fi34 := 0; fi34 < 6; fi34++ {
        if ((v1 & 0) == v1) {
            var v35 int64 = 1
            v35 += (^(v35))
            _ = v35
        }
        var v36 int32 = 42
        if (32 >= 18446744073709551615) {
            v1 -= v1
            var v37 bool = true
            _ = v37
        }
        _ = fi34
        _ = v36
    }
    switch 32768 {
    case -2147483648:
        if (0 == 2) {
            v38 := GenS2{F0: 128, F1: genFn0(42, 103)}
            _ = v38
        }
    case -885:
        var v39 int32 = 694
        _ = v39
    default:
        if false {
            if true {
                v1 += 0
                var v40 float64 = 0.1
                for fi41 := 0; fi41 < 4; fi41++ {
                    var v42 uint64 = 741
                    v43 := GenS3{F0: 127}
                    if ((("\n\t" + "") > "a") && false) {
                        var v44 uint16 = v1
                        for fi45 := 0; fi45 < 5; fi45++ {
                            if false {
                                if (v3 >= (fi41 | 2147483647)) {
                                    fmt.Println(v44)
                                }
                            } else {
                                v46 := GenS3{F0: 100}
                                if ("" >= "hello") {
                                    var v47 string = ((("foo bar" + "foo bar") + "") + ("foo bar" + "hello"))
                                    var v48 float32 = (((1.5 - -2.25) - 0.0) * ((1e10 / (((0.0 * 0.1) * (0.1 - (3.14159 * ((1e10 / float32(v40)) / -2.25)))) - (-2.25 + float32(v40)))) + 0.1))
                                    _ = v47
                                    _ = v48
                                } else {
                                    v49 := GenS4{F0: false, F1: ("\n\t" + "a"), F2: 65535}
                                    if true {
                                        v3 += -2147483648
                                        var v50 uint64 = 2
                                        if (608 >= -62) {
                                            v51 := GenS5{F0: (v1 >> ((42) & 63)), F1: v50, F2: "foo bar"}
                                            v52 := GenS4{F0: false, F1: "a", F2: 4294967295}
                                            var v53 int32 = 42
                                            _ = v51
                                            _ = v52
                                            _ = v53
                                        }
                                        _ = v50
                                    }
                                    v44 += v1
                                    _ = v49
                                }
                                if false {
                                    fmt.Println(v3, v42, fi41)
                                    var v54 uint = 822
                                    _ = v54
                                }
                                _ = v46
                            }
                            if ((("" + "hello") + "hello") == ("hello" + "")) {
                                v55 := GenS2{F0: 1, F1: -32768}
                                v40 = v40
                                var v56 float32 = (((-2.25 + 1.5) + 0.0) / float32((0.1 + v40)))
                                _ = v55
                                _ = v56
                            } else {
                                if (!(0 < -2147483648)) {
                                    v57 := GenS6{F0: ("" + "foo bar"), F1: -32768, F2: (v42 << ((0) & 63))}
                                    var v58 bool = ((1e10 <= -2.25) && false)
                                    _ = v57
                                    _ = v58
                                }
                            }
                            v1 = (v44 << ((2) & 63))
                            _ = fi45
                        }
                        for fi59 := 0; fi59 < 8; fi59++ {
                            v3 += v3
                            _ = fi59
                        }
                        _ = v44
                    } else {
                        for fi60 := 0; fi60 < 2; fi60++ {
                            if (0 <= 255) {
                                var v61 int64 = -9223372036854775808
                                v2[32768] = v3
                                _ = v61
                            } else {
                                fmt.Println(fi60, v1, v40)
                                v1 = 377
                            }
                            if true {
                                fmt.Println(fi41, v40, v3)
                                v62 := []int{v3, v3, fi41, int(v42)}
                                _ = v62
                            } else {
                                if false {
                                    if true {
                                        v40 += v40
                                        v63 := GenS7{F0: 732, F1: (v1 ^ 32767), F2: 65535}
                                        v64 := []int{(^(fi41)), (-(v3)), v3, (fi60 + 0)}
                                        _ = v63
                                        _ = v64
                                    }
                                    var v65 uint32 = 0
                                    _ = v65
                                }
                            }
                            _ = fi60
                        }
                        v66 := []int{-9223372036854775808, fi41}
                        _ = v66
                    }
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                _ = v40
            }
            var v67 uint8 = 86
            v3 += 4294967295
            _ = v67
        } else {
            for fi68 := 0; fi68 < 16; fi68++ {
                var v69 uint64 = 4294967295
                var v70 float32 = 1e10
                v71 := []int{(fi68 ^ (fi68 + (fi68 / (((fi68 ^ -201)) | 1)))), int(v1), -733, fi68}
                _ = fi68
                _ = v69
                _ = v70
                _ = v71
            }
            fmt.Println(v1, v3)
            v3 = -2147483648
        }
    }
    v72 := []int{(v3 ^ (v3 + 2147483647)), v3, (v3 + v3), (v3 + v3)}
    var v73 bool = (940 == 65535)
    if (18446744073709551615 == 2) {
        if false {
            v73 = v73
        } else {
            v73 = (128 > 0)
            for fi74 := 0; fi74 < 8; fi74++ {
                v75 := GenS8{F0: 0, F1: -1}
                for fi76 := 0; fi76 < 3; fi76++ {
                    v73 = (v1 > v1)
                    if v73 {
                        v77 := v2[(-(fi74))]
                        if v73 {
                            var v78 int64 = 206
                            _ = v78
                        } else {
                            if true {
                                var v79 bool = true
                                var v80 int64 = -870
                                v77 = (v77 ^ (v3 / ((fi74) | 1)))
                                _ = v79
                                _ = v80
                            }
                        }
                        _ = v77
                    } else {
                        if false {
                            v73 = true
                        }
                        v81 := v2[32768]
                        _ = v81
                    }
                    switch v3 {
                    case 4294967295:
                        v3 = (fi76 * -9223372036854775808)
                    case -2147483648:
                        v82 := GenS9{F0: 433, F1: 42, F2: 42}
                        _ = v82
                    case 127:
                        if false {
                            v3 = 127
                        } else {
                            var v83 float64 = 0.1
                            v1 -= (v1 * (v1 | (v1 & 42)))
                            if (281 >= 2) {
                                switch (("" + "foo bar") + "hello") {
                                case "a":
                                    var v84 int8 = -1
                                    _ = v84
                                case "hello":
                                    if len(v72) >= 2 {
                                        v72 = v72[1:len(v72)]
                                    }
                                case "":
                                    v3 = (fi76 << ((18446744073709551615) & 63))
                                default:
                                    switch (fi74 % ((fi74) | 1)) {
                                    case 32767:
                                        var v85 float32 = (3.14159 - 1e10)
                                        switch ("a" + "\n\t") {
                                        case "hello":
                                            fmt.Println(v83, v83)
                                        case "foo bar":
                                            switch 9223372036854775807 {
                                            case 753:
                                                fmt.Println(v3, fi74, v1)
                                            case 9223372036854775807:
                                                if len(v72) >= 2 {
                                                    v72 = v72[1:len(v72)]
                                                }
                                                var v86 int = -62
                                                _ = v86
                                            case 4294967295:
                                                if (128 >= 252) {
                                                    var v87 uint64 = 4294967295
                                                    _ = v87
                                                } else {
                                                    v73 = true
                                                    v88 := GenS10{F0: 9223372036854775807, F1: v3}
                                                    if len(v72) > 0 {
                                                        v72[0] = v3
                                                    }
                                                    _ = v88
                                                }
                                            default:
                                                if len(v72) > 0 {
                                                    v72[0] = (v3 | (fi74 / ((fi76) | 1)))
                                                }
                                            }
                                        case "a":
                                            var v89 uint = 480
                                            _ = v89
                                        }
                                        v90 := GenS8{F0: 255, F1: 1}
                                        _ = v85
                                        _ = v90
                                    case 0:
                                        if len(v72) > 0 {
                                            v72[0] = v3
                                        }
                                    case 2147483647:
                                        switch int(v1) {
                                        case 54:
                                            var v91 uint64 = 2
                                            _ = v91
                                        case 32768:
                                            if v73 {
                                                switch (fi76 & (v3 ^ (^(fi74)))) {
                                                case -9223372036854775808:
                                                    v92 := GenS3{F0: 0}
                                                    _ = v92
                                                case 42:
                                                    v2[(fi74 & fi76)] = 9223372036854775807
                                                case 127:
                                                    fmt.Println(v83, v73, fi74)
                                                default:
                                                    var v93 uint16 = v1
                                                    _ = v93
                                                }
                                                if len(v72) > 0 {
                                                    v72[0] = (v3 + 32767)
                                                }
                                                if v73 {
                                                    v94 := GenS9{F0: 361, F1: 4294967295, F2: 127}
                                                    if v73 {
                                                        switch (fi76 / ((fi76) | 1)) {
                                                        case -1:
                                                            if (3.14159 == 3.14159) {
                                                                var v95 float64 = float64((0.0 + 1.5))
                                                                _ = v95
                                                            } else {
                                                                if len(v72) > 0 {
                                                                    v72[0] = v3
                                                                }
                                                                v3 = fi74
                                                            }
                                                        }
                                                        var v96 int16 = 1
                                                        _ = v96
                                                    } else {
                                                        var v97 float64 = 1.5
                                                        if true {
                                                            v98 := GenS11{F0: 2147483648, F1: 86}
                                                            _ = v98
                                                        } else {
                                                            v83 += (v83 - 0.0)
                                                            v83 -= ((v83 / 0.1) / v83)
                                                            v97 = 3.14159
                                                        }
                                                        _ = v97
                                                    }
                                                    if len(v72) >= 2 {
                                                        v72 = v72[1:len(v72)]
                                                    }
                                                    switch "\n\t" {
                                                    case "hello":
                                                        v73 = v73
                                                    case "\n\t":
                                                        v83 = (3.14159 - v83)
                                                    case "a":
                                                        v99 := GenS9{F0: 0, F1: 1, F2: 42}
                                                        _ = v99
                                                    default:
                                                        var v100 string = "\n\t"
                                                        _ = v100
                                                    }
                                                    _ = v94
                                                } else {
                                                    fmt.Println(v83)
                                                    var v101 uint = 4294967295
                                                    var v102 int64 = 42
                                                    _ = v101
                                                    _ = v102
                                                }
                                            }
                                        case -2147483648:
                                            if (3.14159 >= ((float32(v83) - (1e10 + 1e10)) - ((0.1 + -2.25) + 1.5))) {
                                                v3 = (fi76 % ((v3) | 1))
                                            }
                                        }
                                        v103 := GenS2{F0: 31, F1: 1}
                                        _ = v103
                                    }
                                }
                                fmt.Println(v1)
                                var v104 float32 = -2.25
                                _ = v104
                            }
                            _ = v83
                        }
                    default:
                        switch fi74 {
                        case 4294967295:
                            if v73 {
                                if true {
                                    v73 = (11 >= 42)
                                    var v105 int32 = 42
                                    _ = v105
                                } else {
                                    switch fi74 {
                                    case 2147483647:
                                        v106 := GenS4{F0: (!false), F1: "\n\t", F2: 0}
                                        _ = v106
                                    case 32768:
                                        fmt.Println(v3)
                                    default:
                                        var v107 uint = 71
                                        _ = v107
                                    }
                                    var v108 int = fi76
                                    _ = v108
                                }
                                switch ("a" + (("a" + "hello") + (("" + "a") + "foo bar"))) {
                                case "a":
                                    switch ("" + "a") {
                                    case "hello":
                                        switch "hello" {
                                        case "a":
                                            v109 := GenS5{F0: v1, F1: 42, F2: ("foo bar" + "hello")}
                                            _ = v109
                                        case "\n\t":
                                            var v110 uint = 548
                                            _ = v110
                                        default:
                                            switch v3 {
                                            case -1:
                                                if (!(519 != 0)) {
                                                    var v111 bool = false
                                                    v72 = append(v72, (-(v3)))
                                                    fmt.Println(fi74)
                                                    _ = v111
                                                } else {
                                                    switch "hello" {
                                                    case "\n\t":
                                                        if false {
                                                            v1 = 2
                                                        } else {
                                                            if (!false) {
                                                                v112 := GenS11{F0: 457, F1: int8(v1)}
                                                                v113 := v2[(fi74 / ((fi76) | 1))]
                                                                var v114 int = -551
                                                                _ = v112
                                                                _ = v113
                                                                _ = v114
                                                            }
                                                            switch ("foo bar" + "") {
                                                            case "a":
                                                                if v73 {
                                                                    v115 := GenS2{F0: 255, F1: 1000}
                                                                    if (!v73) {
                                                                        if v73 {
                                                                            _ = len(v72)
                                                                        }
                                                                        var v116 uint = 2
                                                                        switch "a" {
                                                                        case "a":
                                                                            v73 = (v73 && (65535 > 2))
                                                                        default:
                                                                            if v73 {
                                                                                var v117 string = "a"
                                                                                _ = v117
                                                                            }
                                                                        }
                                                                        _ = v116
                                                                    } else {
                                                                        var v118 int16 = 119
                                                                        v1 -= v1
                                                                        _ = v118
                                                                    }
                                                                    _ = v115
                                                                }
                                                            case "\n\t":
                                                                v119 := v2[32767]
                                                                _ = v119
                                                            case "hello":
                                                                var v120 int = fi76
                                                                _ = v120
                                                            default:
                                                                if len(v72) > 0 {
                                                                    v72[0] = (fi76 % (((fi74 & fi74)) | 1))
                                                                }
                                                            }
                                                        }
                                                    }
                                                    fmt.Println(v3, v3)
                                                    var v121 uint32 = 42
                                                    _ = v121
                                                }
                                            case 127:
                                                if true {
                                                    var v122 int32 = -148
                                                    v122 -= -99
                                                    _ = len(v72)
                                                    var v123 uint = 1
                                                    _ = v122
                                                    _ = v123
                                                } else {
                                                    if (1000 > -32768) {
                                                        _ = len(v2)
                                                        v124 := GenS12{F0: 0, F1: 18446744073709551615, F2: 2147483647}
                                                        _ = v124
                                                    } else {
                                                        if (!v73) {
                                                            var v125 int = (fi76 | fi76)
                                                            _ = v125
                                                        } else {
                                                            v2[fi74] = fi76
                                                            v73 = (1e10 < (1e10 * (float32((0.0 * (1e10 + (3.14159 / (1.5 / 1.5))))) + 1e10)))
                                                        }
                                                    }
                                                }
                                            default:
                                                v3 += 127
                                            }
                                        }
                                    case "a":
                                        v126 := GenS13{F0: (v3 >> ((65535) & 63)), F1: 949}
                                        _ = v126
                                    case "foo bar":
                                        _ = len(v2)
                                    }
                                }
                            }
                        case -1:
                            if v73 {
                                fmt.Println(v1)
                            }
                        case 32767:
                            switch fi76 {
                            case 0:
                                v1 += (^(v1))
                            case -9223372036854775808:
                                if v73 {
                                    v73 = v73
                                    var v127 uint64 = 533
                                    var v128 bool = (!(!v73))
                                    _ = v127
                                    _ = v128
                                }
                                v72 = append(v72, fi76)
                            default:
                                v129 := GenS14{F0: 1}
                                _ = v129
                            }
                        default:
                            v1 = v1
                        }
                    }
                    var v130 uint64 = 110
                    _ = fi76
                    _ = v130
                }
                for fi131 := 0; fi131 < 20; fi131++ {
                    switch -37 {
                    case 42:
                        _ = len(v72)
                        var v132 uint16 = 1
                        var v133 int32 = 2147483647
                        _ = v132
                        _ = v133
                    case 1:
                        var v134 int32 = 2147483647
                        _ = v134
                    case 139:
                        v1 = 2
                    }
                    v73 = v73
                    _ = fi131
                }
                _ = fi74
                _ = v75
            }
        }
    } else {
        var v135 int8 = 100
        fmt.Println(v3)
        switch "\n\t" {
        case "":
            var v136 float32 = float32(-2.25)
            _ = v136
        case "foo bar":
            v3 = 466
        case "a":
            fmt.Println(v3, v3)
        }
        _ = v135
    }
    _ = len(v72)
    switch v3 {
    case 127:
        fmt.Println(v1)
    case 0:
        for fi137 := 0; fi137 < 1; fi137++ {
            fmt.Println(v1, v1, v1)
            switch "hello" {
            case "":
                var v138 float32 = float32(0.1)
                _ = v138
            case "foo bar":
                v73 = v73
            }
            v139 := GenS15{F0: 948, F1: "foo bar", F2: 2}
            _ = fi137
            _ = v139
        }
    case 3:
        v1 -= v1
    }
    var v140 int16 = 26
    fmt.Println("v1:", v1)
    fmt.Println("v140:", v140)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v72:", len(v72), v72)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v140
    _ = v2
    _ = v3
    _ = v33
    _ = v72
    _ = v73
}
