package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 int8; F1 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int64; F1 int64; F2 uint }
type GenS8 struct { F0 int; F1 uint32 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int8; F1 string }
type GenS11 struct { F0 uint32; F1 int }
type GenS12 struct { F0 int16; F1 int; F2 int }
type GenS13 struct { F0 int; F1 string }
type GenS14 struct { F0 uint32; F1 uint }
type GenS15 struct { F0 uint64; F1 string; F2 uint }
type GenS16 struct { F0 int; F1 uint64; F2 int }
type GenS17 struct { F0 bool }
type GenS18 struct { F0 bool; F1 uint32 }
type GenS19 struct { F0 int; F1 uint }
type GenS20 struct { F0 int64 }
type GenS21 struct { F0 int8 }
type GenS22 struct { F0 uint; F1 string }
type GenS23 struct { F0 int64; F1 bool }
type GenS24 struct { F0 uint; F1 bool; F2 int }
type GenS25 struct { F0 int16; F1 int16 }
type GenS26 struct { F0 uint; F1 int64 }
type GenS27 struct { F0 uint16; F1 int64 }
type GenS28 struct { F0 int8; F1 uint32; F2 int32 }
type GenS29 struct { F0 int }
type GenS30 struct { F0 uint64; F1 string; F2 uint8 }
type GenS31 struct { F0 uint64 }
func genFn0(p0 int32, p1 uint8) string {
    p0 = p0
    v6 := []int{0, -142, 32768, -224}
    for fi7 := 0; fi7 < 17; fi7++ {
        switch fi7 {
        case 905:
            if (!(!true)) {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                if (false && true) {
                    for fi8 := 0; fi8 < 9; fi8++ {
                        p0 = 1
                        v9 := GenS1{F0: -1}
                        _ = fi8
                        _ = v9
                    }
                    if true {
                        p1 = 42
                    } else {
                        for fi10 := 0; fi10 < 14; fi10++ {
                            p0 += (p0 % ((-190) | 1))
                            fmt.Println(p1, p1)
                            p0 += (p0 << ((151) & 63))
                            _ = fi10
                        }
                        if (111 == 302) {
                            if true {
                                if false {
                                    var v11 float64 = 1e10
                                    if false {
                                        p1 = (p1 / ((p1) | 1))
                                        v11 = ((v11 - v11) / 1.5)
                                        var v12 int32 = (p0 / ((1) | 1))
                                        _ = v12
                                    } else {
                                        fmt.Println(v11)
                                    }
                                    var v13 uint = 65535
                                    _ = v11
                                    _ = v13
                                }
                            }
                            var v14 uint = 446
                            _ = v14
                        } else {
                            p1 -= p1
                        }
                        if (fi7 < (fi7 % (((fi7 ^ (fi7 << ((0) & 63)))) | 1))) {
                            v6 = append(v6, 713)
                            switch fi7 {
                            case -958:
                                v15 := GenS2{F0: 1}
                                _ = v15
                            case 32768:
                                var v16 int = 914
                                _ = v16
                            case 730:
                                p0 = (p0 & p0)
                            default:
                                var v17 uint = 1
                                _ = v17
                            }
                        }
                    }
                    fmt.Println(fi7, p1)
                } else {
                    p1 += 1
                }
                if len(v6) > 0 {
                    v6[0] = fi7
                }
            }
        case 32767:
            if ((p1 + 62) != p1) {
                fmt.Println(fi7)
                var v18 int32 = p0
                _ = v18
            }
        case 1:
            p0 -= (p0 << ((42) & 63))
        default:
            if false {
                var v19 uint16 = 1
                if (120 < 127) {
                    var v20 uint16 = (v19 | (^(v19)))
                    var v21 uint16 = 1
                    var v22 uint8 = p1
                    _ = v20
                    _ = v21
                    _ = v22
                }
                _ = v19
            }
        }
        _ = fi7
    }
    switch "" {
    case "a":
        _ = len(v6)
    case "foo bar":
        var v23 int = 969
        _ = v23
    default:
        if true {
            if true {
                switch 2147483647 {
                case 42:
                    var v24 int8 = 89
                    _ = v24
                case 32767:
                    if false {
                        if ((-2.25 != 0.1) && (580 > 0)) {
                            if false {
                                p1 += (p1 + p1)
                                if (!(!false)) {
                                    for fi25 := 0; fi25 < 2; fi25++ {
                                        if true {
                                            switch 42 {
                                            case 9223372036854775807:
                                                var v26 int8 = 0
                                                _ = v26
                                            case 32768:
                                                if (fi25 != fi25) {
                                                    v6 = append(v6, (-(fi25)))
                                                    p1 += 161
                                                } else {
                                                    p1 -= p1
                                                    v27 := GenS3{F0: fi25, F1: 0}
                                                    if false {
                                                        p1 = 1
                                                        for fi28 := 0; fi28 < 2; fi28++ {
                                                            v29 := GenS4{F0: (p1 == p1), F1: int64(p0)}
                                                            var v30 int = fi28
                                                            v30 -= 32767
                                                            _ = fi28
                                                            _ = v29
                                                            _ = v30
                                                        }
                                                    }
                                                    _ = v27
                                                }
                                            case -569:
                                                if true {
                                                    var v31 int16 = -133
                                                    fmt.Println(p1)
                                                    _ = v31
                                                } else {
                                                    p0 = (p0 & p0)
                                                }
                                            default:
                                                p1 -= 128
                                            }
                                            for fi32 := 0; fi32 < 3; fi32++ {
                                                var v33 int16 = -32768
                                                p1 = (^(p1))
                                                v6 = append(v6, fi32)
                                                if (!(!(!true))) {
                                                    v34 := GenS5{F0: 1, F1: p1}
                                                    var v35 uint8 = (^(p1))
                                                    switch ("" + "") {
                                                    case "\n\t":
                                                        if true {
                                                            v6 = append(v6, (fi32 - (fi25 & fi25)))
                                                        }
                                                    }
                                                    _ = v34
                                                    _ = v35
                                                } else {
                                                    var v36 string = "foo bar"
                                                    if true {
                                                        v37 := GenS6{F0: -1}
                                                        _ = v37
                                                    }
                                                    _ = v36
                                                }
                                                _ = fi32
                                                _ = v33
                                            }
                                        }
                                        _ = fi25
                                    }
                                }
                            }
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                        var v38 uint32 = 2147483647
                        _ = v38
                    } else {
                        v39 := GenS1{F0: 0}
                        _ = v39
                    }
                default:
                    if len(v6) > 0 {
                        v6[0] = 127
                    }
                }
                var v40 int64 = -9223372036854775808
                for fi41 := 0; fi41 < 8; fi41++ {
                    if (42 >= v40) {
                        v6 = append(v6, 2147483647)
                        for fi42 := 0; fi42 < 10; fi42++ {
                            if (255 > 255) {
                                var v43 uint8 = p1
                                _ = v43
                            }
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            var v44 uint16 = 1
                            _ = fi42
                            _ = v44
                        }
                        switch "hello" {
                        case "hello":
                            v45 := GenS6{F0: -1}
                            _ = v45
                        case "":
                            if true {
                                var v46 float64 = 3.14159
                                var v47 int = fi41
                                v40 = v40
                                _ = v46
                                _ = v47
                            }
                        default:
                            for fi48 := 0; fi48 < 6; fi48++ {
                                if (-124 != 0) {
                                    if (-1 <= p0) {
                                        v49 := GenS7{F0: v40, F1: (v40 | v40), F2: 65535}
                                        p1 += p1
                                        switch 4294967295 {
                                        case -2147483648:
                                            fmt.Println(v40)
                                        case 4294967295:
                                            fmt.Println(p0, v40)
                                        default:
                                            var v50 int32 = p0
                                            _ = v50
                                        }
                                        v51 := GenS8{F0: 2147483647, F1: 4294967295}
                                        if (42 > 1) {
                                            if true {
                                                var v52 uint32 = 0
                                                if (false || false) {
                                                    v40 = (v40 * int64(p0))
                                                } else {
                                                    fmt.Println(fi41, p1)
                                                }
                                                v6 = append(v6, (fi48 >> ((897) & 63)))
                                                _ = v52
                                            }
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        }
                                        switch "foo bar" {
                                        case "foo bar":
                                            v53 := GenS9{F0: 18446744073709551615}
                                            _ = v53
                                        case "a":
                                            switch ("a" + "hello") {
                                            case "hello":
                                                var v54 uint8 = p1
                                                _ = v54
                                            case "foo bar":
                                                if true {
                                                    var v55 uint64 = 829
                                                    _ = v55
                                                } else {
                                                    var v56 float32 = (((-2.25 + 1e10) - 0.1) - (0.0 * 3.14159))
                                                    _ = v56
                                                }
                                            case "\n\t":
                                                var v57 uint = 65535
                                                _ = v57
                                            }
                                        default:
                                            switch "hello" {
                                            case "hello":
                                                var v58 float64 = float64(0.1)
                                                _ = v58
                                            case "\n\t":
                                                if true {
                                                    if ((((-(v40)) >= (^(v40))) || true) && false) {
                                                        var v59 float32 = 1.5
                                                        _ = v59
                                                    } else {
                                                        var v60 int64 = (v40 & 42)
                                                        _ = v60
                                                    }
                                                    v61 := GenS4{F0: ((v40 & v40) == (v40 & v40)), F1: v40}
                                                    if false {
                                                        if len(v6) > 0 {
                                                            v6[0] = fi41
                                                        }
                                                        var v62 int64 = (v40 << ((17) & 63))
                                                        fmt.Println(v40)
                                                        var v63 bool = (!false)
                                                        v63 = v63
                                                        _ = v62
                                                        _ = v63
                                                    } else {
                                                        if (0.1 == float32(-2.25)) {
                                                            switch (fi48 / ((32768) | 1)) {
                                                            case 0:
                                                                v64 := GenS2{F0: 0}
                                                                _ = v64
                                                            case 2147483647:
                                                                p1 = (p1 / ((p1) | 1))
                                                            }
                                                            v65 := GenS5{F0: -10, F1: p1}
                                                            var v66 uint16 = 1
                                                            if (true || true) {
                                                                var v67 int = fi41
                                                                var v68 uint8 = 0
                                                                _ = v67
                                                                _ = v68
                                                            } else {
                                                                if len(v6) > 0 {
                                                                    v6[0] = int(p0)
                                                                }
                                                                if ((v66 / (((v66 ^ 65535)) | 1)) <= v66) {
                                                                    v69 := GenS3{F0: fi41, F1: 42}
                                                                    _ = v69
                                                                } else {
                                                                    if false {
                                                                        if len(v6) >= 2 {
                                                                            v6 = v6[1:len(v6)]
                                                                        }
                                                                        v6 = append(v6, 42)
                                                                        p0 = (p0 + 0)
                                                                    }
                                                                    switch (fi41 + (fi41 | fi41)) {
                                                                    case 4294967295:
                                                                        v70 := GenS10{F0: -85, F1: "a"}
                                                                        _ = v70
                                                                    }
                                                                    if (false && (42 != -32768)) {
                                                                        if false {
                                                                            if false {
                                                                                var v71 int16 = 1
                                                                                _ = v71
                                                                            }
                                                                        } else {
                                                                            var v72 int64 = v40
                                                                            if false {
                                                                                v73 := GenS11{F0: 2147483647, F1: (fi41 << ((218) & 63))}
                                                                                if true {
                                                                                    v74 := GenS5{F0: 100, F1: 127}
                                                                                    if false {
                                                                                        var v75 uint64 = 42
                                                                                        v6 = append(v6, (^(fi48)))
                                                                                        _ = v75
                                                                                    }
                                                                                    fmt.Println(fi48, v66, fi41)
                                                                                    v40 = (^(v72))
                                                                                    _ = v74
                                                                                }
                                                                                switch fi41 {
                                                                                case 127:
                                                                                    v40 = v72
                                                                                }
                                                                                _ = v73
                                                                            }
                                                                            v72 -= (v40 & -9223372036854775808)
                                                                            _ = v72
                                                                        }
                                                                        var v76 int8 = -1
                                                                        if ((((v40 + v40) <= (v40 | (v40 - v40))) && (182 >= 2147483647)) || (121 != p1)) {
                                                                            var v77 uint64 = uint64(p1)
                                                                            v40 = v40
                                                                            _ = v77
                                                                        }
                                                                        _ = v76
                                                                    } else {
                                                                        _ = len(v6)
                                                                        fmt.Println(p0)
                                                                    }
                                                                }
                                                            }
                                                            _ = v65
                                                            _ = v66
                                                        }
                                                        fmt.Println(p0, p0)
                                                        v78 := GenS10{F0: 1, F1: "foo bar"}
                                                        fmt.Println(p0, v40, fi48)
                                                        _ = v78
                                                    }
                                                    p1 -= p1
                                                    _ = v61
                                                } else {
                                                    var v79 uint32 = 0
                                                    _ = v79
                                                }
                                            default:
                                                p0 -= -388
                                            }
                                        }
                                        _ = v49
                                        _ = v51
                                    }
                                    p0 += 0
                                } else {
                                    fmt.Println(fi48)
                                    if ((1e10 - 1.5) >= (float32((-2.25 / (-2.25 + (0.1 - -2.25)))) - float32(((1.5 + 0.0) * -2.25)))) {
                                        switch 4294967295 {
                                        case 32768:
                                            var v80 uint32 = 4294967295
                                            _ = v80
                                        case -853:
                                            var v81 bool = false
                                            _ = v81
                                        }
                                        var v82 uint16 = 0
                                        _ = v82
                                    }
                                    p0 += (p0 % ((0) | 1))
                                }
                                _ = fi48
                            }
                        }
                    }
                    _ = fi41
                }
                _ = v40
            }
        }
    }
    p1 += 181
    _ = p0
    _ = p1
    _ = v6
    return "a"
}
func genFn1(p0 uint16, p1 int16, p2 uint8) int {
    v83 := GenS12{F0: (-(p1)), F1: 32768, F2: 128}
    var v84 float32 = float32((1e10 / -2.25))
    v85 := GenS13{F0: -382, F1: "a"}
    var v86 int32 = 2147483647
    v87 := []int{32768, 42, 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v83
    _ = v84
    _ = v85
    _ = v86
    _ = v87
    return 42
}

func main() {
    v1 := []uint32{0, 871, 516}
    var v2 uint16 = 856
    v3 := []int{-293}
    var v4 int16 = 42
    var v5 uint8 = 127
    var v88 float64 = 0.0
    fmt.Println(v4, v88)
    v2 = (v2 & (v2 / ((v2) | 1)))
    switch "" {
    case "foo bar":
        if true {
            for fi89 := 0; fi89 < 18; fi89++ {
                v2 = 553
                _ = fi89
            }
        }
    }
    fmt.Println(v88, v2)
    if false {
        v90 := GenS14{F0: 2147483647, F1: 42}
        v3 = append(v3, 128)
        _ = v90
    } else {
        switch 2147483647 {
        case 2147483647:
            v91 := GenS15{F0: 843, F1: "\n\t", F2: 42}
            _ = v91
        }
        v92 := GenS16{F0: 322, F1: 0, F2: 127}
        v93 := GenS17{F0: (true && (true && true))}
        _ = v92
        _ = v93
    }
    var v94 int = genFn1((v2 & (v2 - v2)), (v4 - 0), (v5 ^ v5))
    if false {
        var v95 int32 = -125
        _ = v95
    } else {
        for fi96 := 0; fi96 < 19; fi96++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi96
        }
        for fi97 := 0; fi97 < 18; fi97++ {
            var v98 float64 = v88
            var v99 float32 = float32((0.1 / 0.1))
            v88 = v88
            _ = fi97
            _ = v98
            _ = v99
        }
        v2 += 32767
    }
    v2 = (v2 + v2)
    var v100 uint = 4294967295
    v88 = v88
    switch (v94 * (v94 | (v94 << ((v100) & 63)))) {
    case -92:
        v88 = (3.14159 + ((3.14159 + 1.5) * v88))
    default:
        for fi101 := 0; fi101 < 3; fi101++ {
            var v102 int8 = -1
            if false {
                var v103 uint8 = v5
                for fi104 := 0; fi104 < 11; fi104++ {
                    switch v94 {
                    case -480:
                        if (!(genFn1(255, (v4 | v4), 0) < int(v100))) {
                            if ((362 > 968) && (uint(v2) < (v100 / (((^(v100))) | 1)))) {
                                fmt.Println(fi104, v2)
                                v2 -= v2
                                v5 = 155
                                v4 = v4
                                var v105 int = int(v2)
                                _ = v105
                            } else {
                                v106 := GenS13{F0: 42, F1: "\n\t"}
                                v4 += (v4 & (v4 | (v4 & (v4 * (v4 >> ((v100) & 63))))))
                                _ = v106
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if (((42 > 42) && (!false)) && true) {
                                var v107 int64 = -1
                                _ = v107
                            } else {
                                _ = len(v1)
                                if false {
                                    switch 993 {
                                    case 9223372036854775807:
                                        if (48 < 185) {
                                            if false {
                                                if false {
                                                    if (18446744073709551615 == 0) {
                                                        v108 := GenS16{F0: genFn1((v2 % (((v2 % ((v2) | 1))) | 1)), 1000, v103), F1: 42, F2: (^(v94))}
                                                        if (v103 != (^(v5))) {
                                                            var v109 bool = false
                                                            var v110 int8 = 1
                                                            fmt.Println(v100)
                                                            _ = v109
                                                            _ = v110
                                                        }
                                                        _ = v108
                                                    } else {
                                                        var v111 uint = (v100 >> ((v100) & 63))
                                                        v102 += v102
                                                        _ = len(v3)
                                                        _ = v111
                                                    }
                                                }
                                                var v112 int64 = -178
                                                _ = len(v1)
                                                _ = v112
                                            } else {
                                                fmt.Println(v88, v100)
                                            }
                                            v2 += v2
                                            v94 = -2147483648
                                        } else {
                                            switch genFn1(521, 32767, 127) {
                                            case 1:
                                                switch (v94 | 39) {
                                                case -9223372036854775808:
                                                    switch 32767 {
                                                    case 328:
                                                        v2 = (^(v2))
                                                    case 1:
                                                        v113 := GenS9{F0: uint64(v2)}
                                                        _ = v113
                                                    }
                                                case 128:
                                                    var v114 int = fi101
                                                    _ = v114
                                                case 977:
                                                    if len(v1) > 0 {
                                                        v1[0] = 18
                                                    }
                                                }
                                            case -2147483648:
                                                switch (-(v94)) {
                                                case 1:
                                                    v5 -= v103
                                                default:
                                                    v115 := GenS1{F0: int16(fi101)}
                                                    _ = v115
                                                }
                                                switch genFn1((v2 << ((v100) & 63)), 840, v5) {
                                                case 4294967295:
                                                    if len(v1) > 0 {
                                                        v1[0] = 2147483647
                                                    }
                                                default:
                                                    v100 = 0
                                                }
                                                if len(v3) > 0 {
                                                    v3[0] = genFn1(v2, (v4 ^ int16(v5)), (v5 | v103))
                                                }
                                            }
                                        }
                                    case -801:
                                        v116 := GenS18{F0: false, F1: 42}
                                        _ = v116
                                    }
                                    var v117 float64 = v88
                                    _ = v117
                                }
                                v118 := GenS19{F0: (fi104 - genFn1(32767, int16(v100), 2)), F1: 995}
                                _ = v118
                            }
                        } else {
                            v5 -= v5
                            var v119 int64 = 42
                            v88 += (float64((-2.25 + 3.14159)) * (1e10 / 1.0))
                            _ = v119
                        }
                    case 1:
                        if true {
                            fmt.Println(v103)
                            var v120 uint16 = 42
                            _ = v120
                        }
                    case -1:
                        if true {
                            if len(v1) > 0 {
                                v1[0] = 2
                            }
                            v94 += fi101
                            v121 := GenS20{F0: -1}
                            _ = v121
                        } else {
                            v88 = (v88 + (v88 * 1.5))
                        }
                    default:
                        var v122 int = (fi101 >> (((v100 - (^(v100)))) & 63))
                        _ = v122
                    }
                    if (!((!(v4 != (v4 & 1000))) && (false && false))) {
                        v4 = (v4 ^ -32768)
                    }
                    _ = fi104
                }
                if ((!true) || ((^(fi101)) >= 2147483647)) {
                    _ = len(v3)
                    for fi123 := 0; fi123 < 19; fi123++ {
                        v4 += (-(v4))
                        if true {
                            var v124 int64 = 0
                            if (335 == 493) {
                                var v125 float64 = 0.1
                                var v126 string = "foo bar"
                                switch ((("hello" + v126) + v126) + (v126 + v126)) {
                                case "hello":
                                    if true {
                                        if (0 > 859) {
                                            v102 -= 85
                                        }
                                    }
                                }
                                _ = v125
                                _ = v126
                            } else {
                                if true {
                                    if (true || (fi123 == (v94 & int(v4)))) {
                                        v102 -= 0
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v1 = append(v1, 42)
                                    }
                                } else {
                                    if (88 >= 2) {
                                        if false {
                                            fmt.Println(v94, v102, v103)
                                        } else {
                                            v94 -= 1
                                            switch ("\n\t" + "hello") {
                                            case "foo bar":
                                                v4 -= -32768
                                            }
                                            switch "hello" {
                                            case "hello":
                                                v2 += (v2 >> ((42) & 63))
                                            case "\n\t":
                                                v127 := GenS21{F0: (v102 / ((int8(v4)) | 1))}
                                                _ = v127
                                            case "":
                                                if true {
                                                    var v128 int32 = 42
                                                    if true {
                                                        var v129 int8 = v102
                                                        switch fi101 {
                                                        case -2147483648:
                                                            v129 += (v129 >> (((v100 * 749)) & 63))
                                                        default:
                                                            if true {
                                                                v130 := GenS22{F0: (v100 & v100), F1: ("\n\t" + "foo bar")}
                                                                var v131 string = "foo bar"
                                                                _ = v130
                                                                _ = v131
                                                            } else {
                                                                if len(v3) > 0 {
                                                                    v3[0] = (v94 % ((int(v129)) | 1))
                                                                }
                                                                v103 = v5
                                                            }
                                                            var v132 uint16 = 32767
                                                            var v133 float32 = 1.5
                                                            var v134 uint = (v100 ^ (v100 & (v100 % ((v100) | 1))))
                                                            _ = v132
                                                            _ = v133
                                                            _ = v134
                                                        }
                                                        v100 = v100
                                                        _ = v129
                                                    } else {
                                                        var v135 int32 = (-(v128))
                                                        _ = v135
                                                    }
                                                    var v136 float64 = (((float64(1.5) / (1.5 / v88)) - v88) / 1e10)
                                                    _ = v128
                                                    _ = v136
                                                } else {
                                                    if len(v3) > 0 {
                                                        v3[0] = genFn1(v2, v4, uint8(v100))
                                                    }
                                                    fmt.Println(v94, v2, v103)
                                                }
                                            }
                                            switch ("foo bar" + "hello") {
                                            case "\n\t":
                                                v88 += (v88 * v88)
                                            case "foo bar":
                                                var v137 uint8 = (v103 - v103)
                                                _ = v137
                                            case "":
                                                var v138 float32 = (-2.25 + 1.5)
                                                if (-1 == 42) {
                                                    switch genFn1(v2, v4, v103) {
                                                    case 4294967295:
                                                        var v139 uint16 = (v2 - 0)
                                                        _ = v139
                                                    case 128:
                                                        var v140 uint16 = v2
                                                        _ = v140
                                                    case -1:
                                                        v1 = append(v1, 311)
                                                        v88 += (0.0 / float64(v138))
                                                    default:
                                                        var v141 float64 = v88
                                                        _ = v141
                                                    }
                                                    var v142 uint64 = 2
                                                    var v143 float64 = (v88 / v88)
                                                    _ = v142
                                                    _ = v143
                                                } else {
                                                    v94 = genFn1((v2 * 65535), int16(v2), v5)
                                                    v4 = (^(v4))
                                                }
                                                var v144 uint = 18446744073709551615
                                                _ = v138
                                                _ = v144
                                            default:
                                                var v145 int32 = int32(v4)
                                                _ = v145
                                            }
                                            v5 = (v103 << ((v100) & 63))
                                        }
                                        v1 = append(v1, 465)
                                        if (int16(v94) >= (v4 & (v4 ^ (v4 - (v4 / (((v4 + 0)) | 1)))))) {
                                            var v146 uint32 = 51
                                            _ = v146
                                        } else {
                                            fmt.Println(v124)
                                            switch v94 {
                                            case -108:
                                                v147 := GenS23{F0: (v124 >> ((4294967295) & 63)), F1: true}
                                                _ = v147
                                            case 32767:
                                                if false {
                                                    if true {
                                                        fmt.Println(v4, v88, v102)
                                                        fmt.Println(v2)
                                                        if false {
                                                            var v148 int64 = v124
                                                            if true {
                                                                v88 = 0.1
                                                                v102 = (v102 + int8(v2))
                                                                v149 := GenS24{F0: (v100 / ((v100) | 1)), F1: true, F2: (fi123 >> ((4294967295) & 63))}
                                                                var v150 uint32 = 65535
                                                                _ = v149
                                                                _ = v150
                                                            }
                                                            switch "foo bar" {
                                                            case "a":
                                                                var v151 bool = ((!true) || (-2.25 > ((0.1 + 1.5) * (((((1.5 / float32(v88)) + 3.14159) + ((-2.25 / (0.1 / (0.1 * ((-2.25 - ((float32(float64((float32(v88) / 3.14159))) * (1e10 / 1.5)) / -2.25)) - 1e10)))) / float32((float64(-2.25) / (v88 / 1.0))))) + (3.14159 * float32(0.1))) - ((((3.14159 * float32(v88)) + (0.1 * -2.25)) * (-2.25 * 1.5)) - 0.1)))))
                                                                _ = v151
                                                            case "":
                                                                var v152 int64 = -1
                                                                _ = v152
                                                            }
                                                            _ = v148
                                                        }
                                                        v4 = int16(v102)
                                                        var v153 uint32 = 65535
                                                        fmt.Println(v100, v103)
                                                        v102 -= v102
                                                        _ = v153
                                                    } else {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    }
                                                    var v154 uint8 = (v5 & v103)
                                                    if (true && (false || false)) {
                                                        if (true && (4294967295 != uint64(v2))) {
                                                            v88 += v88
                                                            var v155 bool = true
                                                            var v156 uint16 = (^(v2))
                                                            _ = v155
                                                            _ = v156
                                                        } else {
                                                            var v157 int16 = (v4 - int16(v154))
                                                            if (v2 < v2) {
                                                                v4 = (v4 | -32768)
                                                            }
                                                            _ = v157
                                                        }
                                                        if (v94 > 1) {
                                                            if false {
                                                                fmt.Println(v2, v5, v88)
                                                                v100 += 18446744073709551615
                                                            } else {
                                                                v158 := GenS25{F0: v4, F1: int16(v5)}
                                                                _ = v158
                                                            }
                                                        }
                                                    } else {
                                                        switch "hello" {
                                                        case "foo bar":
                                                            v159 := GenS13{F0: (fi123 + (-(v94))), F1: "\n\t"}
                                                            _ = v159
                                                        case "":
                                                            v160 := GenS24{F0: 42, F1: true, F2: (fi101 * 2147483647)}
                                                            _ = v160
                                                        default:
                                                            if true {
                                                                fmt.Println(v2)
                                                                v88 -= v88
                                                                v161 := GenS14{F0: 716, F1: v100}
                                                                _ = v161
                                                            }
                                                        }
                                                    }
                                                    _ = v154
                                                } else {
                                                    v162 := GenS21{F0: -34}
                                                    var v163 int = 42
                                                    v164 := GenS26{F0: v100, F1: (v124 - v124)}
                                                    _ = v162
                                                    _ = v163
                                                    _ = v164
                                                }
                                            case 2147483647:
                                                v165 := GenS27{F0: v2, F1: v124}
                                                _ = v165
                                            }
                                            if true {
                                                fmt.Println(v103)
                                                switch "hello" {
                                                case "a":
                                                    if (fi123 < 9223372036854775807) {
                                                        var v166 int64 = -9223372036854775808
                                                        switch (fi101 - genFn1(v2, (v4 - (v4 + (v4 % ((v4) | 1)))), 2)) {
                                                        case 4294967295:
                                                            var v167 uint8 = v5
                                                            _ = v167
                                                        default:
                                                            var v168 int8 = (v102 << ((341) & 63))
                                                            _ = v168
                                                        }
                                                        _ = v166
                                                    }
                                                }
                                                v2 = (v2 - 65535)
                                                switch ("" + (("hello" + "") + ("\n\t" + "foo bar"))) {
                                                case "a":
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                default:
                                                    v100 = (v100 + (v100 + (v100 % (((v100 / ((v100) | 1))) | 1))))
                                                    if true {
                                                        _ = len(v1)
                                                    } else {
                                                        v169 := GenS20{F0: -1}
                                                        var v170 uint32 = 164
                                                        _ = v169
                                                        _ = v170
                                                    }
                                                    var v171 int64 = (v124 + -1)
                                                    _ = v171
                                                }
                                                v100 = v100
                                            }
                                        }
                                    } else {
                                        fmt.Println(v124, v4, v88)
                                        if (!true) {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                    }
                                }
                                v94 = 9223372036854775807
                            }
                            _ = v124
                        } else {
                            v4 += (v4 % ((v4) | 1))
                        }
                        _ = fi123
                    }
                    v5 = 42
                }
                _ = v103
            }
            v1 = append(v1, 2147483648)
            _ = fi101
            _ = v102
        }
    }
    v172 := GenS12{F0: (v4 % ((v4) | 1)), F1: v94, F2: (v94 * (v94 | (v94 << ((42) & 63))))}
    switch (v94 % (((v94 * genFn1((v2 / ((32767) | 1)), (-(v4)), 127))) | 1)) {
    case -9223372036854775808:
        v88 = (0.1 - v88)
    case -74:
        if len(v1) > 0 {
            v1[0] = 42
        }
    default:
        for fi173 := 0; fi173 < 3; fi173++ {
            for fi174 := 0; fi174 < 1; fi174++ {
                fmt.Println(v2, v4, fi174)
                v100 -= (v100 * uint(v5))
                _ = fi174
            }
            var v175 string = "\n\t"
            _ = fi173
            _ = v175
        }
    }
    for fi176 := 0; fi176 < 6; fi176++ {
        v177 := GenS28{F0: 22, F1: 39, F2: -1}
        var v178 int = 9223372036854775807
        v179 := GenS21{F0: 0}
        _ = fi176
        _ = v177
        _ = v178
        _ = v179
    }
    if (float32(1.5) == ((0.0 + (-2.25 * 0.1)) * 3.14159)) {
        if true {
            if ((1e10 - (0.1 / float64(-2.25))) < 1.5) {
                var v180 int16 = v4
                var v181 float64 = v88
                _ = v180
                _ = v181
            } else {
                v94 = v94
                for fi182 := 0; fi182 < 3; fi182++ {
                    v183 := GenS29{F0: v94}
                    _ = len(v3)
                    switch "" {
                    case "\n\t":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case "foo bar":
                        for fi184 := 0; fi184 < 6; fi184++ {
                            v1 = append(v1, 65535)
                            _ = fi184
                        }
                    case "hello":
                        switch "a" {
                        case "":
                            if (1 < 4294967295) {
                                for fi185 := 0; fi185 < 2; fi185++ {
                                    var v186 int32 = 42
                                    var v187 uint16 = 2
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = fi185
                                    _ = v186
                                    _ = v187
                                }
                                v94 += 1
                            } else {
                                v100 = (v100 & 18446744073709551615)
                                v100 = v100
                                if true {
                                    switch "" {
                                    case "\n\t":
                                        v4 -= v4
                                    default:
                                        switch "hello" {
                                        case "foo bar":
                                            for fi188 := 0; fi188 < 13; fi188++ {
                                                v3 = append(v3, (v94 & v94))
                                                var v189 uint = 255
                                                var v190 float32 = float32(v88)
                                                _ = fi188
                                                _ = v189
                                                _ = v190
                                            }
                                        case "":
                                            var v191 float32 = 1e10
                                            _ = v191
                                        default:
                                            switch (-(fi182)) {
                                            case -187:
                                                fmt.Println(v88, v88)
                                            default:
                                                v192 := GenS30{F0: 0, F1: "hello", F2: v5}
                                                _ = v192
                                            }
                                        }
                                    }
                                    var v193 string = ("" + "hello")
                                    _ = v193
                                } else {
                                    if true {
                                        if (1 <= 0) {
                                            v88 = v88
                                        }
                                    }
                                    var v194 uint = 23
                                    _ = v194
                                }
                            }
                        case "hello":
                            v100 = v100
                        case "a":
                            if (-499 >= 386) {
                                v100 -= v100
                                switch ("\n\t" + ((("foo bar" + "") + "") + "foo bar")) {
                                case "\n\t":
                                    for fi195 := 0; fi195 < 12; fi195++ {
                                        v196 := GenS12{F0: (v4 + int16(v5)), F1: 229, F2: fi195}
                                        _ = fi195
                                        _ = v196
                                    }
                                case "foo bar":
                                    var v197 bool = false
                                    _ = v197
                                }
                            }
                        default:
                            var v198 int64 = int64(v4)
                            _ = v198
                        }
                    default:
                        v199 := GenS31{F0: 42}
                        _ = v199
                    }
                    _ = fi182
                    _ = v183
                }
                var v200 int32 = -2147483648
                _ = v200
            }
        }
    }
    if len(v1) > 0 {
        v1[0] = 482
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v88:", v88)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v100
    _ = v172
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v88
    _ = v94
}
