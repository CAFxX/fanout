package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 int; F1 uint8; F2 int8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 bool; F1 int32 }
type GenS8 struct { F0 uint32; F1 uint8 }
type GenS9 struct { F0 bool; F1 int32 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int; F1 uint8 }
type GenS12 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS13 struct { F0 uint8; F1 int64 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 int64; F1 int16; F2 uint }
func genFn0(p0 uint16) uint32 {
    for fi7 := 0; fi7 < 6; fi7++ {
        if true {
            var v8 uint64 = 1
            v9 := []int{fi7, (fi7 / ((342) | 1))}
            for fi10 := 0; fi10 < 11; fi10++ {
                var v11 float32 = -2.25
                var v12 float32 = (v11 + 0.1)
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = v8
            _ = v9
        }
        var v13 uint8 = 255
        _ = fi7
        _ = v13
    }
    p0 = (p0 * 2)
    p0 -= 72
    var v14 bool = false
    _ = p0
    _ = v14
    return 2
}
func genFn1(p0 int, p1 uint16) int64 {
    v15 := GenS1{F0: "hello", F1: 42}
    p1 -= 255
    switch p0 {
    case 1:
        for fi16 := 0; fi16 < 10; fi16++ {
            p1 = (p1 >> ((42) & 63))
            p0 += fi16
            if false {
                var v17 uint8 = 42
                var v18 int16 = 0
                _ = v17
                _ = v18
            } else {
                p1 = p1
                v19 := []int{(p0 * p0), (^(p0)), -1}
                _ = v19
            }
            _ = fi16
        }
    default:
        p1 = (p1 % (((p1 * p1)) | 1))
    }
    switch p0 {
    case 233:
        fmt.Println(p0, p0)
    case 4294967295:
        for fi20 := 0; fi20 < 6; fi20++ {
            p1 = (^(p1))
            switch -1 {
            case -21:
                if true {
                    switch ("hello" + "hello") {
                    case "\n\t":
                        for fi21 := 0; fi21 < 4; fi21++ {
                            var v22 uint64 = 42
                            v23 := GenS1{F0: "\n\t", F1: -1}
                            _ = fi21
                            _ = v22
                            _ = v23
                        }
                    }
                } else {
                    v24 := GenS2{F0: (p1 % (((^(p1))) | 1))}
                    if false {
                        var v25 int64 = -1
                        var v26 int32 = int32(p1)
                        var v27 int64 = v25
                        _ = v25
                        _ = v26
                        _ = v27
                    } else {
                        if (31 >= int32(p0)) {
                            p0 += -961
                            var v28 uint16 = (p1 - (p1 >> ((786) & 63)))
                            v28 += v28
                            _ = v28
                        } else {
                            fmt.Println(fi20, fi20)
                            for fi29 := 0; fi29 < 3; fi29++ {
                                var v30 int32 = int32(p0)
                                switch p0 {
                                case -2147483648:
                                    p0 += -9223372036854775808
                                default:
                                    var v31 uint32 = 599
                                    _ = v31
                                }
                                _ = fi29
                                _ = v30
                            }
                            if true {
                                v32 := GenS1{F0: ("\n\t" + "foo bar"), F1: 9223372036854775807}
                                var v33 float64 = ((1.5 / (-2.25 + 0.0)) / (0.1 - 1.5))
                                p0 = 756
                                _ = v32
                                _ = v33
                            }
                        }
                        var v34 int32 = 1
                        if (((p1 > p1) && false) && false) {
                            var v35 float32 = -2.25
                            _ = v35
                        } else {
                            if true {
                                switch "hello" {
                                case "a":
                                    v34 = (-(v34))
                                default:
                                    v34 = -2147483648
                                }
                                fmt.Println(v34)
                                fmt.Println(v34)
                            }
                            var v36 string = "a"
                            _ = v36
                        }
                        _ = v34
                    }
                    _ = v24
                }
            case 4294967295:
                v37 := []int{(^(p0)), (fi20 + (^(p0))), (p0 & fi20)}
                _ = v37
            default:
                for fi38 := 0; fi38 < 6; fi38++ {
                    p0 = (^(fi38))
                    _ = fi38
                }
            }
            var v39 uint16 = 443
            _ = fi20
            _ = v39
        }
    case 128:
        p1 = (p1 / ((211) | 1))
    default:
        var v40 uint32 = 127
        _ = v40
    }
    if ((((0.0 / 1.0) < 1e10) && (false || (!true))) && false) {
        for fi41 := 0; fi41 < 9; fi41++ {
            p1 = p1
            v42 := []int{p0, (fi41 + fi41)}
            _ = len(v42)
            _ = fi41
            _ = v42
        }
        var v43 int = p0
        _ = v43
    }
    _ = p0
    _ = p1
    _ = v15
    return -9223372036854775808
}

func main() {
    var v1 uint8 = 42
    var v2 uint16 = 75
    var v3 int = int(v1)
    var v4 int8 = int8(v3)
    v5 := []uint32{42, 2147483648}
    var v6 uint8 = v1
    for fi44 := 0; fi44 < 13; fi44++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi44
    }
    if (false && true) {
        v1 = (v6 | 1)
        var v45 int32 = 42
        if false {
            if (2 >= 1) {
                var v46 uint16 = (v2 % ((2) | 1))
                _ = v46
            }
            for fi47 := 0; fi47 < 16; fi47++ {
                if false {
                    if true {
                        for fi48 := 0; fi48 < 18; fi48++ {
                            fmt.Println(v2, v1)
                            _ = fi48
                        }
                        switch (^(fi47)) {
                        case 4294967295:
                            v49 := GenS3{F0: 2147483648}
                            _ = v49
                        }
                    } else {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        if true {
                            v50 := GenS3{F0: 2147483647}
                            v3 = v3
                            _ = v50
                        } else {
                            v1 = (v1 & v6)
                            var v51 bool = (true && true)
                            _ = v51
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                    if (v3 < -267) {
                        v5 = append(v5, 670)
                    }
                    for fi52 := 0; fi52 < 4; fi52++ {
                        fmt.Println(v45, v45, v6)
                        _ = len(v5)
                        _ = fi52
                    }
                } else {
                    var v53 uint32 = 2
                    var v54 bool = (!true)
                    if v54 {
                        if len(v5) > 0 {
                            v5[0] = (v53 | (v53 << ((4294967295) & 63)))
                        }
                        var v55 uint = 760
                        _ = v55
                    } else {
                        if v54 {
                            for fi56 := 0; fi56 < 11; fi56++ {
                                if false {
                                    if v54 {
                                        v57 := GenS1{F0: ("foo bar" + "foo bar"), F1: genFn1(9223372036854775807, v2)}
                                        _ = v57
                                    } else {
                                        v58 := GenS3{F0: (^(v53))}
                                        v59 := GenS3{F0: (v53 | (^(v53)))}
                                        _ = v58
                                        _ = v59
                                    }
                                    var v60 int16 = 466
                                    v61 := GenS1{F0: "a", F1: -683}
                                    _ = v60
                                    _ = v61
                                } else {
                                    var v62 uint = 65535
                                    v6 = v1
                                    _ = len(v5)
                                    _ = v62
                                }
                                var v63 uint32 = (v53 % ((v53) | 1))
                                v63 = v63
                                _ = fi56
                                _ = v63
                            }
                        }
                        var v64 int16 = 0
                        var v65 int32 = v45
                        _ = v64
                        _ = v65
                    }
                    _ = v53
                    _ = v54
                }
                _ = fi47
            }
        }
        _ = v45
    }
    v3 -= (v3 % ((v3) | 1))
    for fi66 := 0; fi66 < 4; fi66++ {
        for fi67 := 0; fi67 < 19; fi67++ {
            switch "" {
            case "hello":
                var v68 int32 = 0
                _ = v68
            case "\n\t":
                var v69 float64 = 1e10
                _ = v69
            default:
                var v70 uint16 = 2
                _ = v70
            }
            _ = fi67
        }
        _ = fi66
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi71 := 0; fi71 < 2; fi71++ {
        v72 := GenS4{F0: 42, F1: 32767}
        switch ("hello" + "") {
        case "":
            v2 += (v2 & (v2 + 255))
        case "a":
            var v73 uint32 = 871
            _ = v73
        case "foo bar":
            fmt.Println(fi71)
        default:
            switch "hello" {
            case "":
                for fi74 := 0; fi74 < 14; fi74++ {
                    if (1 != 934) {
                        var v75 float32 = 0.0
                        if (0.0 == 3.14159) {
                            fmt.Println(v3, v75, v3)
                            v76 := GenS3{F0: 0}
                            switch fi71 {
                            case 479:
                                v2 = (v2 % ((65535) | 1))
                            case 4294967295:
                                switch "a" {
                                case "":
                                    v5 = append(v5, 2)
                                case "a":
                                    switch "\n\t" {
                                    case "":
                                        if (-2.25 > 1e10) {
                                            fmt.Println(fi71)
                                            var v77 float64 = 3.14159
                                            _ = v77
                                        }
                                        var v78 int32 = 0
                                        _ = len(v5)
                                        _ = v78
                                    }
                                case "\n\t":
                                    _ = len(v5)
                                }
                            case 42:
                                var v79 string = ("foo bar" + "\n\t")
                                switch 193 {
                                case -9223372036854775808:
                                    var v80 uint8 = (^(v1))
                                    v81 := GenS4{F0: 2, F1: (v2 / ((v2) | 1))}
                                    _ = v80
                                    _ = v81
                                case 9223372036854775807:
                                    v4 = v4
                                case -2147483648:
                                    v79 = "a"
                                }
                                fmt.Println(v79, v6)
                                v1 -= v6
                                _ = len(v5)
                                _ = v79
                            }
                            _ = v76
                        }
                        _ = v75
                    }
                    _ = fi74
                }
            case "hello":
                v2 = (v2 - v2)
            case "a":
                for fi82 := 0; fi82 < 12; fi82++ {
                    if true {
                        if ((-32768 == 32767) || (32767 != v2)) {
                            fmt.Println(v2, fi82, v1)
                        } else {
                            switch (("" + "a") + ("\n\t" + ("hello" + "a"))) {
                            case "foo bar":
                                v83 := GenS4{F0: 1, F1: (v2 & 60)}
                                if true {
                                    if false {
                                        if false {
                                            fmt.Println(v3, v1)
                                        } else {
                                            v84 := GenS5{F0: (fi71 + 32768), F1: (v6 >> ((65535) & 63)), F2: v4}
                                            fmt.Println(v6, v4, fi82)
                                            if true {
                                                if true {
                                                    v3 = fi82
                                                    v3 = -1
                                                    var v85 string = ""
                                                    _ = v85
                                                } else {
                                                    var v86 float64 = 1e10
                                                    _ = v86
                                                }
                                                v87 := GenS4{F0: 79, F1: v2}
                                                fmt.Println(fi82, fi82)
                                                switch 63 {
                                                case 2147483647:
                                                    v2 -= (^(v2))
                                                case 0:
                                                    v88 := GenS1{F0: "a", F1: 0}
                                                    if true {
                                                        var v89 int64 = 42
                                                        _ = len(v5)
                                                        v4 = 127
                                                        switch -190 {
                                                        case 127:
                                                            _ = len(v5)
                                                        case 939:
                                                            v6 += v6
                                                        case 756:
                                                            var v90 uint64 = 0
                                                            _ = v90
                                                        default:
                                                            v5 = append(v5, 1)
                                                        }
                                                        _ = v89
                                                    }
                                                    var v91 int64 = 0
                                                    _ = v88
                                                    _ = v91
                                                case 9223372036854775807:
                                                    var v92 uint32 = 193
                                                    _ = v92
                                                default:
                                                    switch v3 {
                                                    case -881:
                                                        switch "hello" {
                                                        case "\n\t":
                                                            if true {
                                                                v93 := GenS6{F0: v2}
                                                                if false {
                                                                    if (134 < 244) {
                                                                        var v94 uint32 = 0
                                                                        var v95 int64 = -9223372036854775808
                                                                        v5 = append(v5, 377)
                                                                        _ = v94
                                                                        _ = v95
                                                                    } else {
                                                                        var v96 uint8 = (v6 + v1)
                                                                        _ = v96
                                                                    }
                                                                    v1 = v1
                                                                } else {
                                                                    fmt.Println(fi71)
                                                                    var v97 uint16 = 0
                                                                    if ((-299 <= fi82) || (true || (1000 <= -1))) {
                                                                        v6 = v6
                                                                    } else {
                                                                        var v98 int64 = -1
                                                                        switch ("foo bar" + "foo bar") {
                                                                        case "hello":
                                                                            v6 -= v6
                                                                        case "":
                                                                            var v99 uint32 = 65535
                                                                            _ = v99
                                                                        default:
                                                                            v100 := GenS7{F0: (false && (genFn0((v97 / (((v97 - 65535)) | 1))) == 2)), F1: 2147483647}
                                                                            switch "a" {
                                                                            case "foo bar":
                                                                                var v101 uint16 = (v2 ^ v97)
                                                                                _ = v101
                                                                            }
                                                                            fmt.Println(v3, v2, v1)
                                                                            _ = v100
                                                                        }
                                                                        _ = v98
                                                                    }
                                                                    _ = v97
                                                                }
                                                                v1 = 246
                                                                _ = v93
                                                            }
                                                        case "hello":
                                                            if false {
                                                                v2 = uint16(v1)
                                                                v1 = 42
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                                if (18446744073709551615 < 371) {
                                                                    v4 = v4
                                                                }
                                                            } else {
                                                                if true {
                                                                    _ = len(v5)
                                                                }
                                                                switch (fi82 % ((9223372036854775807) | 1)) {
                                                                case 4294967295:
                                                                    v102 := GenS8{F0: 1, F1: (v6 - (^(v6)))}
                                                                    _ = v102
                                                                case 127:
                                                                    if len(v5) > 0 {
                                                                        v5[0] = genFn0(32767)
                                                                    }
                                                                case 0:
                                                                    if false {
                                                                        v103 := GenS9{F0: (18446744073709551615 != 42), F1: -2147483648}
                                                                        _ = v103
                                                                    }
                                                                default:
                                                                    v6 = (^(v6))
                                                                    var v104 bool = (!false)
                                                                    _ = v104
                                                                }
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    v105 := GenS10{F0: -29}
                                                                    _ = v105
                                                                case "a":
                                                                    var v106 uint = 912
                                                                    _ = v106
                                                                case "hello":
                                                                    if false {
                                                                        fmt.Println(v1, v4, fi82)
                                                                    }
                                                                default:
                                                                    var v107 int = 32768
                                                                    _ = v107
                                                                }
                                                            }
                                                        default:
                                                            var v108 int8 = 1
                                                            _ = v108
                                                        }
                                                    case -2147483648:
                                                        v1 += (^(v1))
                                                    case 149:
                                                        v4 = 42
                                                        var v109 uint32 = 65535
                                                        _ = v109
                                                    }
                                                }
                                                v110 := GenS2{F0: (^(v2))}
                                                _ = v110
                                                _ = v87
                                            }
                                            _ = v84
                                        }
                                        v2 = v2
                                        _ = len(v5)
                                    } else {
                                        v4 -= (v4 & (v4 | (v4 + (v4 << ((101) & 63)))))
                                        v111 := GenS1{F0: ("hello" + "hello"), F1: 0}
                                        v4 += v4
                                        _ = v111
                                    }
                                    if (v4 == v4) {
                                        v1 = 2
                                    }
                                    if true {
                                        v112 := GenS10{F0: genFn1(fi82, v2)}
                                        _ = v112
                                    }
                                    if false {
                                        v113 := GenS9{F0: (!false), F1: 2147483647}
                                        _ = v113
                                    }
                                    switch (v3 | (fi71 & fi71)) {
                                    case -68:
                                        fmt.Println(v4, v6, v3)
                                    case 0:
                                        switch (fi71 / (((fi71 + fi82)) | 1)) {
                                        case 4294967295:
                                            v114 := GenS3{F0: 1}
                                            var v115 float32 = 0.1
                                            _ = v114
                                            _ = v115
                                        case 127:
                                            switch -2147483648 {
                                            case 478:
                                                switch (^(fi71)) {
                                                case 9223372036854775807:
                                                    v6 = v1
                                                case 541:
                                                    var v116 uint16 = v2
                                                    _ = v116
                                                case 128:
                                                    switch 32768 {
                                                    case -595:
                                                        switch (v3 - fi71) {
                                                        case 128:
                                                            switch ("" + (("" + "") + "foo bar")) {
                                                            case "hello":
                                                                var v117 bool = false
                                                                _ = len(v5)
                                                                _ = v117
                                                            case "\n\t":
                                                                switch "" {
                                                                case "":
                                                                    var v118 uint8 = (v1 % (((v6 * (v6 >> ((4294967295) & 63)))) | 1))
                                                                    _ = v118
                                                                case "hello":
                                                                    v2 = (v2 + (v2 % ((1) | 1)))
                                                                case "a":
                                                                    switch (^(v3)) {
                                                                    case 62:
                                                                        var v119 float32 = (((0.1 + 1e10) + ((1e10 + 0.0) + 0.1)) + float32(1e10))
                                                                        _ = v119
                                                                    default:
                                                                        v120 := GenS9{F0: false, F1: 42}
                                                                        _ = v120
                                                                    }
                                                                default:
                                                                    fmt.Println(fi82)
                                                                }
                                                            }
                                                            var v121 int = (^(fi71))
                                                            v3 -= (fi71 ^ fi71)
                                                            _ = v121
                                                        case -1:
                                                            v2 = v2
                                                            v122 := GenS6{F0: (^(v2))}
                                                            _ = v122
                                                        case 32767:
                                                            var v123 int64 = -265
                                                            if false {
                                                                switch ("a" + "") {
                                                                case "a":
                                                                    v124 := GenS9{F0: true, F1: -726}
                                                                    _ = v124
                                                                case "foo bar":
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                case "\n\t":
                                                                    var v125 int32 = 2147483647
                                                                    _ = v125
                                                                default:
                                                                    fmt.Println(v3, fi71)
                                                                }
                                                                v4 += (v4 / ((v4) | 1))
                                                                if false {
                                                                    switch v3 {
                                                                    case 42:
                                                                        switch "a" {
                                                                        case "foo bar":
                                                                            v126 := GenS11{F0: fi82, F1: (v6 / ((127) | 1))}
                                                                            _ = v126
                                                                        default:
                                                                            _ = len(v5)
                                                                        }
                                                                    case -5:
                                                                        fmt.Println(fi82, v6, v4)
                                                                    case 32767:
                                                                        if false {
                                                                            var v127 uint64 = 18446744073709551615
                                                                            if len(v5) > 0 {
                                                                                v5[0] = 920
                                                                            }
                                                                            _ = v127
                                                                        }
                                                                    default:
                                                                        var v128 int64 = (v123 + (v123 << ((375) & 63)))
                                                                        _ = v128
                                                                    }
                                                                }
                                                            } else {
                                                                v4 = 1
                                                                fmt.Println(v123, v4, v123)
                                                                v2 -= (v2 << ((4294967295) & 63))
                                                            }
                                                            v6 = (v6 * 58)
                                                            _ = v123
                                                        default:
                                                            fmt.Println(v1, v4)
                                                        }
                                                    case -110:
                                                        v129 := GenS12{F0: 2147483648, F1: ((-(fi82)) <= 9223372036854775807), F2: v1}
                                                        _ = v129
                                                    case -370:
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    }
                                                default:
                                                    var v130 uint8 = (v1 ^ v6)
                                                    _ = v130
                                                }
                                            case -2147483648:
                                                switch "foo bar" {
                                                case "hello":
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                case "":
                                                    var v131 uint8 = (^(v6))
                                                    _ = v131
                                                default:
                                                    v1 += v6
                                                }
                                            case 42:
                                                v5 = append(v5, 0)
                                            }
                                        case 128:
                                            fmt.Println(v6)
                                        }
                                        if (!true) {
                                            if true {
                                                switch (fi71 & 128) {
                                                case 945:
                                                    v132 := GenS6{F0: 932}
                                                    _ = v132
                                                }
                                            }
                                        } else {
                                            v5 = append(v5, genFn0(v2))
                                            switch (fi82 << ((1) & 63)) {
                                            case 127:
                                                var v133 int16 = 32767
                                                _ = v133
                                            case 9223372036854775807:
                                                fmt.Println(v6, v4)
                                            case -1:
                                                fmt.Println(v6)
                                            }
                                            v6 = (v6 << ((0) & 63))
                                        }
                                        var v134 uint64 = uint64(v2)
                                        _ = v134
                                    }
                                }
                                v2 = (v2 / ((1) | 1))
                                var v135 int8 = v4
                                _ = v135
                                _ = v83
                            }
                        }
                        if true {
                            v136 := GenS8{F0: 70, F1: v1}
                            switch ("a" + "foo bar") {
                            case "foo bar":
                                var v137 float32 = (1.5 / (1e10 + 0.0))
                                _ = v137
                            case "a":
                                v4 = (v4 * (v4 - (v4 << ((2) & 63))))
                            case "\n\t":
                                v138 := GenS13{F0: v1, F1: 177}
                                _ = v138
                            }
                            _ = v136
                        } else {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                    } else {
                        if len(v5) > 0 {
                            v5[0] = 1
                        }
                        switch "hello" {
                        case "":
                            var v139 int64 = 323
                            _ = v139
                        case "hello":
                            var v140 int8 = -128
                            _ = v140
                        case "a":
                            fmt.Println(v4, v6, v1)
                        default:
                            v141 := GenS14{F0: (v1 - 128)}
                            _ = v141
                        }
                        v2 = v2
                    }
                    fmt.Println(v3)
                    switch (fi82 | fi82) {
                    case -1:
                        if (true && false) {
                            var v142 int8 = v4
                            _ = v142
                        }
                    case -848:
                        _ = len(v5)
                    case 446:
                        v143 := GenS8{F0: 4294967295, F1: 128}
                        var v144 float64 = (-2.25 + ((3.14159 - (-2.25 / (3.14159 * 3.14159))) - 3.14159))
                        switch fi71 {
                        case 42:
                            v2 += v2
                        default:
                            var v145 uint32 = 2147483648
                            _ = v145
                        }
                        _ = v143
                        _ = v144
                    }
                    v4 = 0
                    v2 -= (^(v2))
                    _ = fi82
                }
            default:
                v2 += (^(v2))
            }
        }
        _ = fi71
        _ = v72
    }
    var v146 int = v3
    switch v3 {
    case -228:
        for fi147 := 0; fi147 < 11; fi147++ {
            v148 := GenS8{F0: 42, F1: 222}
            _ = fi147
            _ = v148
        }
    case 32767:
        v146 += (v146 & v3)
    case -2147483648:
        var v149 int64 = 0
        _ = v149
    default:
        var v150 float64 = 0.1
        _ = v150
    }
    for fi151 := 0; fi151 < 19; fi151++ {
        switch ("hello" + "a") {
        case "hello":
            for fi152 := 0; fi152 < 8; fi152++ {
                v146 += int(v6)
                v6 = v1
                _ = fi152
            }
        case "a":
            var v153 uint = 65535
            _ = v153
        }
        switch 4294967295 {
        case 128:
            v154 := GenS15{F0: 9223372036854775807, F1: -32768, F2: 2}
            _ = v154
        case 240:
            v155 := GenS2{F0: (v2 ^ 2)}
            _ = v155
        default:
            fmt.Println(v146, v2)
        }
        _ = fi151
    }
    if false {
        for fi156 := 0; fi156 < 3; fi156++ {
            v5 = append(v5, 0)
            v5 = append(v5, uint32(v2))
            switch (v146 * (v3 | v3)) {
            case 32767:
                v1 = v6
            }
            _ = fi156
        }
    } else {
        v1 = v6
        var v157 float32 = 0.0
        v146 = v3
        _ = v157
    }
    if true {
        v158 := GenS4{F0: 18446744073709551615, F1: (^(v2))}
        if true {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        } else {
            switch ("a" + "foo bar") {
            case "\n\t":
                if (false || false) {
                    v6 = v1
                    var v159 int8 = (v4 % (((v4 << ((918) & 63))) | 1))
                    v1 = 128
                    _ = v159
                } else {
                    v5 = append(v5, 42)
                }
            case "hello":
                fmt.Println(v6, v6, v1)
            case "a":
                var v160 uint8 = (v1 ^ v6)
                _ = v160
            }
            var v161 int32 = 296
            _ = v161
        }
        _ = v158
    } else {
        fmt.Println(v2)
        if false {
            if true {
                var v162 float64 = 3.14159
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if len(v5) > 0 {
                    v5[0] = genFn0((v2 / (((v2 % ((v2) | 1))) | 1)))
                }
                _ = v162
            }
            var v163 string = "hello"
            _ = v163
        }
    }
    v146 += v3
    if len(v5) > 0 {
        v5[0] = uint32(v2)
    }
    var v164 int = v3
    v165 := GenS8{F0: 0, F1: 124}
    fmt.Println("v1:", v1)
    fmt.Println("v146:", v146)
    fmt.Println("v164:", v164)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v146
    _ = v164
    _ = v165
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
