package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 int8; F1 int; F2 uint8 }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint8; F1 int8; F2 uint64 }
type GenS8 struct { F0 uint16; F1 uint8; F2 bool }
type GenS9 struct { F0 int }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 bool; F1 uint8 }
type GenS14 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS15 struct { F0 int8; F1 bool; F2 string }
type GenS16 struct { F0 int8; F1 uint16 }
func genFn0() int16 {
    var v5 int8 = 0
    if true {
        v6 := GenS1{F0: ((0.0 + 1e10) <= (0.1 + 1.5)), F1: 1, F2: 32767}
        for fi7 := 0; fi7 < 18; fi7++ {
            v8 := []int{2147483647, 1}
            for fi9 := 0; fi9 < 6; fi9++ {
                v5 = (v5 * (v5 >> ((255) & 63)))
                if len(v8) > 0 {
                    v8[0] = (fi9 ^ (fi9 << ((65535) & 63)))
                }
                _ = fi9
            }
            _ = fi7
            _ = v8
        }
        _ = v6
    } else {
        v10 := []int{-2147483648}
        _ = v10
    }
    _ = v5
    return -32768
}

func main() {
    var v1 int64 = 42
    v2 := make(map[int]int)
    var v3 uint64 = 2
    var v4 float64 = 0.0
    var v11 int = 42
    switch v11 {
    case 32768:
        var v12 uint = 1
        _ = v12
    case -1:
        v11 -= (-(v11))
    }
    var v13 uint8 = 128
    if false {
        for fi14 := 0; fi14 < 10; fi14++ {
            v13 = v13
            for fi15 := 0; fi15 < 3; fi15++ {
                var v16 bool = false
                v3 = v3
                v17 := GenS2{F0: v13, F1: genFn0()}
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = fi14
        }
        switch ("hello" + "hello") {
        case "a":
            if ((true || true) || true) {
                var v18 string = "foo bar"
                fmt.Println(v11, v11)
                _ = v18
            } else {
                var v19 int = v11
                _ = v19
            }
        default:
            var v20 int = 770
            _ = v20
        }
    } else {
        var v21 int8 = 1
        _ = v21
    }
    switch 2147483647 {
    case -1:
        fmt.Println(v13, v1)
    }
    switch v11 {
    case 765:
        fmt.Println(v3)
    case -2147483648:
        v13 -= (v13 | 128)
    case 0:
        var v22 int32 = 480
        _ = v22
    }
    v23 := GenS3{F0: 1, F1: (-(v11)), F2: 127}
    if (2147483647 <= 4294967295) {
        if true {
            for fi24 := 0; fi24 < 11; fi24++ {
                for fi25 := 0; fi25 < 2; fi25++ {
                    var v26 int8 = 90
                    v27 := GenS2{F0: v13, F1: -623}
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                v4 = 0.1
                switch v11 {
                case 0:
                    v4 = (float64((1e10 - -2.25)) / (3.14159 - (0.0 + v4)))
                case 1:
                    var v28 uint = 255
                    _ = v28
                }
                _ = fi24
            }
            for fi29 := 0; fi29 < 12; fi29++ {
                switch (("hello" + "\n\t") + "") {
                case "a":
                    v30 := GenS2{F0: v13, F1: -1}
                    _ = v30
                case "":
                    v2[int(v3)] = fi29
                }
                _ = fi29
            }
        } else {
            switch v11 {
            case -1:
                fmt.Println(v13, v13, v4)
            default:
                if true {
                    v13 = (v13 / ((v13) | 1))
                }
            }
            var v31 int = -912
            var v32 int8 = -90
            _ = v31
            _ = v32
        }
        var v33 float64 = (v4 * v4)
        _ = v33
    }
    v3 -= (v3 ^ v3)
    v34 := []int{(-(v11))}
    v35 := GenS4{F0: (v11 << ((65535) & 63)), F1: 0}
    v36 := GenS5{F0: 44}
    switch (^(v11)) {
    case -577:
        v37 := GenS6{F0: "hello"}
        _ = v37
    case 1:
        var v38 string = ("\n\t" + "hello")
        _ = v38
    case 128:
        switch "hello" {
        case "foo bar":
            v39 := GenS6{F0: ("foo bar" + "foo bar")}
            _ = v39
        case "":
            v34 = append(v34, (v11 * v11))
        }
    default:
        for fi40 := 0; fi40 < 18; fi40++ {
            delete(v2, (^(v11)))
            _ = fi40
        }
    }
    if false {
        if len(v34) >= 2 {
            v34 = v34[1:len(v34)]
        }
        if false {
            v41 := v2[v11]
            var v42 int8 = 90
            _ = v41
            _ = v42
        }
        switch v11 {
        case 127:
            v43 := GenS2{F0: v13, F1: 1}
            _ = v43
        case -2147483648:
            v1 = v1
        case -1:
            for fi44 := 0; fi44 < 6; fi44++ {
                v4 = v4
                _ = fi44
            }
        default:
            for fi45 := 0; fi45 < 8; fi45++ {
                v4 = float64(-2.25)
                if true {
                    fmt.Println(v1, v3)
                } else {
                    for fi46 := 0; fi46 < 16; fi46++ {
                        var v47 uint8 = v13
                        if len(v34) >= 2 {
                            v34 = v34[1:len(v34)]
                        }
                        _ = fi46
                        _ = v47
                    }
                    for fi48 := 0; fi48 < 1; fi48++ {
                        delete(v2, (fi45 << ((132) & 63)))
                        if false {
                            var v49 int64 = (v1 & -1)
                            if false {
                                switch ("" + "hello") {
                                case "\n\t":
                                    switch (("" + "foo bar") + "a") {
                                    case "hello":
                                        if true {
                                            fmt.Println(fi48, v1)
                                            v49 -= 9223372036854775807
                                            var v50 float32 = 1.5
                                            _ = v50
                                        } else {
                                            switch (("foo bar" + "a") + "\n\t") {
                                            case "\n\t":
                                                if len(v34) >= 2 {
                                                    v34 = v34[1:len(v34)]
                                                }
                                            case "":
                                                var v51 uint32 = 669
                                                _ = v51
                                            case "a":
                                                v1 = -1
                                            default:
                                                fmt.Println(v13, v11)
                                            }
                                            v52 := GenS7{F0: 9, F1: 1, F2: 18446744073709551615}
                                            _ = v52
                                        }
                                        var v53 string = (("\n\t" + "hello") + "")
                                        _ = v53
                                    case "":
                                        if len(v34) > 0 {
                                            v34[0] = int(v3)
                                        }
                                    }
                                case "a":
                                    if (v13 > 192) {
                                        v3 -= (v3 & v3)
                                        v11 = (fi48 ^ (fi45 ^ -649))
                                    } else {
                                        v3 -= 18446744073709551615
                                        fmt.Println(fi45)
                                        var v54 int32 = 0
                                        _ = v54
                                    }
                                    switch "foo bar" {
                                    case "hello":
                                        if (!true) {
                                            fmt.Println(v3, fi48)
                                            v1 = (-(v49))
                                            if true {
                                                var v55 float64 = float64(3.14159)
                                                _ = v55
                                            } else {
                                                var v56 int32 = 358
                                                _ = v56
                                            }
                                            v3 += 42
                                        }
                                    case "\n\t":
                                        if ("hello" > "hello") {
                                            fmt.Println(v4, fi45)
                                            v11 += -9223372036854775808
                                        }
                                    case "foo bar":
                                        v57 := GenS7{F0: v13, F1: 0, F2: (v3 & v3)}
                                        _ = v57
                                    }
                                    v1 += v1
                                    if false {
                                        fmt.Println(v1, v4)
                                        fmt.Println(v49, fi45)
                                    }
                                    v58 := GenS4{F0: (-(v11)), F1: int32(v13)}
                                    _ = v58
                                case "hello":
                                    if (true && (((!false) && false) && true)) {
                                        v11 = fi45
                                        v11 = 32767
                                    }
                                default:
                                    fmt.Println(v3, v11)
                                    var v59 uint32 = 985
                                    _ = v59
                                }
                            }
                            _ = v49
                        } else {
                            if false {
                                var v60 float64 = v4
                                _ = v60
                            } else {
                                if (42 != 1) {
                                    if true {
                                        if true {
                                            v11 += 128
                                            switch "" {
                                            case "hello":
                                                if ((v1 + v1) > (v1 | v1)) {
                                                    var v61 int32 = -2147483648
                                                    v13 += (v13 * (^(v13)))
                                                    var v62 bool = (uint16(v13) == 856)
                                                    _ = v61
                                                    _ = v62
                                                } else {
                                                    var v63 uint64 = v3
                                                    var v64 uint8 = (v13 ^ v13)
                                                    _ = v63
                                                    _ = v64
                                                }
                                            case "a":
                                                v3 -= 1
                                                v65 := GenS1{F0: false, F1: 0, F2: 255}
                                                var v66 uint32 = 42
                                                _ = v65
                                                _ = v66
                                            case "":
                                                var v67 uint32 = 2
                                                _ = v67
                                            default:
                                                switch ("" + "foo bar") {
                                                case "hello":
                                                    if (0 >= 2) {
                                                        v34 = append(v34, (fi45 << ((1) & 63)))
                                                    } else {
                                                        if (-2147483648 >= 2147483647) {
                                                            if ("" < "foo bar") {
                                                                v13 = 255
                                                                v68 := GenS5{F0: 4294967295}
                                                                v4 = (0.0 * (v4 - v4))
                                                                _ = v68
                                                            }
                                                            if (320 >= 42) {
                                                                v69 := GenS4{F0: v11, F1: 0}
                                                                _ = v69
                                                            }
                                                        } else {
                                                            if len(v34) > 0 {
                                                                v34[0] = (fi48 << ((0) & 63))
                                                            }
                                                            if false {
                                                                v70 := v2[fi48]
                                                                switch "" {
                                                                case "\n\t":
                                                                    fmt.Println(v4)
                                                                case "":
                                                                    v71 := GenS5{F0: 42}
                                                                    _ = v71
                                                                default:
                                                                    var v72 int16 = 42
                                                                    if ((511 <= 4294967295) || ((v13 + v13) > 2)) {
                                                                        v73 := v2[v70]
                                                                        var v74 int16 = v72
                                                                        var v75 float64 = (((v4 + 0.1) * 1e10) / 1.5)
                                                                        v1 = v1
                                                                        switch v11 {
                                                                        case -439:
                                                                            switch ("foo bar" + "\n\t") {
                                                                            case "\n\t":
                                                                                var v76 int64 = v1
                                                                                _ = v76
                                                                            case "foo bar":
                                                                                var v77 uint8 = v13
                                                                                _ = v77
                                                                            case "a":
                                                                                var v78 int64 = v1
                                                                                _ = v78
                                                                            default:
                                                                                v74 = 433
                                                                            }
                                                                        }
                                                                        _ = v73
                                                                        _ = v74
                                                                        _ = v75
                                                                    }
                                                                    _ = v72
                                                                }
                                                                fmt.Println(v13)
                                                                var v79 int = (-(v70))
                                                                v70 = fi48
                                                                _ = v70
                                                                _ = v79
                                                            } else {
                                                                _ = len(v34)
                                                                if (false || ((65535 == 65535) && true)) {
                                                                    var v80 bool = ((v13 + (v13 - (v13 * (v13 - v13)))) >= v13)
                                                                    if false {
                                                                        if (0.1 > 3.14159) {
                                                                            var v81 int = (^(fi45))
                                                                            switch ("hello" + "foo bar") {
                                                                            case "":
                                                                                v82 := GenS4{F0: v81, F1: int32(v1)}
                                                                                _ = v82
                                                                            case "\n\t":
                                                                                v3 = (^(v3))
                                                                            }
                                                                            _ = v81
                                                                        } else {
                                                                            fmt.Println(v80, v80)
                                                                        }
                                                                        switch fi45 {
                                                                        case 902:
                                                                            switch "a" {
                                                                            case "a":
                                                                                if v80 {
                                                                                    fmt.Println(v4)
                                                                                    v1 += (-(v1))
                                                                                } else {
                                                                                    var v83 uint64 = v3
                                                                                    switch fi48 {
                                                                                    case 42:
                                                                                        if v80 {
                                                                                            switch 71 {
                                                                                            case -284:
                                                                                                fmt.Println(v80, v11, v83)
                                                                                            case -389:
                                                                                                v3 = 354
                                                                                                var v84 int32 = 0
                                                                                                var v85 int64 = (v1 & (^(v1)))
                                                                                                var v86 int64 = (v85 << ((694) & 63))
                                                                                                _ = v84
                                                                                                _ = v85
                                                                                                _ = v86
                                                                                            case 1:
                                                                                                v11 = (v11 >> ((405) & 63))
                                                                                            default:
                                                                                                fmt.Println(v80, v3, v80)
                                                                                            }
                                                                                            var v87 int = (v11 | fi48)
                                                                                            v4 = v4
                                                                                            _ = v87
                                                                                        } else {
                                                                                            fmt.Println(fi45)
                                                                                            var v88 float32 = ((-2.25 / 1.5) / (1.5 + float32(0.0)))
                                                                                            _ = v88
                                                                                        }
                                                                                    case 127:
                                                                                        var v89 uint8 = 255
                                                                                        _ = v89
                                                                                    }
                                                                                    _ = v83
                                                                                }
                                                                            default:
                                                                                delete(v2, 127)
                                                                            }
                                                                        case 664:
                                                                            v11 -= (fi48 * (v11 & v11))
                                                                        case 460:
                                                                            v90 := GenS7{F0: (v13 | v13), F1: 42, F2: v3}
                                                                            _ = v90
                                                                        }
                                                                        if len(v34) >= 2 {
                                                                            v34 = v34[1:len(v34)]
                                                                        }
                                                                    }
                                                                    if ((0.1 + 0.1) > 3.14159) {
                                                                        switch v11 {
                                                                        case 32768:
                                                                            v91 := GenS2{F0: uint8(v3), F1: 42}
                                                                            _ = v91
                                                                        case 2147483647:
                                                                            _ = len(v34)
                                                                            v1 = v1
                                                                        }
                                                                        v13 += 255
                                                                    }
                                                                    _ = v80
                                                                }
                                                            }
                                                        }
                                                    }
                                                case "\n\t":
                                                    var v92 uint64 = (v3 + uint64(v13))
                                                    _ = v92
                                                }
                                            }
                                        } else {
                                            switch fi48 {
                                            case 127:
                                                if (127 > 0) {
                                                    switch fi48 {
                                                    case -1:
                                                        v93 := GenS2{F0: 127, F1: -338}
                                                        _ = v93
                                                    case -9223372036854775808:
                                                        switch (((("\n\t" + ("hello" + ("a" + ("a" + "a")))) + ("\n\t" + "a")) + "a") + ("hello" + "")) {
                                                        case "":
                                                            v94 := GenS7{F0: v13, F1: 0, F2: 90}
                                                            _ = v94
                                                        }
                                                    case 4294967295:
                                                        v34 = append(v34, -9223372036854775808)
                                                    }
                                                    if len(v34) >= 2 {
                                                        v34 = v34[1:len(v34)]
                                                    }
                                                    var v95 uint64 = (^(v3))
                                                    v1 -= v1
                                                    _ = len(v2)
                                                    var v96 int32 = 1
                                                    _ = v95
                                                    _ = v96
                                                } else {
                                                    var v97 int32 = -1
                                                    _ = v97
                                                }
                                                if (4294967295 <= 18446744073709551615) {
                                                    v98 := GenS8{F0: 181, F1: v13, F2: (-1 < -128)}
                                                    _ = v98
                                                }
                                            default:
                                                v2[(fi48 * (fi48 >> ((741) & 63)))] = v11
                                            }
                                            switch "hello" {
                                            case "foo bar":
                                                var v99 bool = false
                                                _ = v99
                                            default:
                                                var v100 int64 = (-(v1))
                                                _ = v100
                                            }
                                        }
                                        var v101 int8 = int8(v1)
                                        _ = v101
                                    }
                                } else {
                                    switch (^(v11)) {
                                    case -678:
                                        var v102 int8 = 66
                                        _ = v102
                                    case 135:
                                        var v103 float64 = v4
                                        _ = v103
                                    case 127:
                                        switch int(v3) {
                                        case 42:
                                            fmt.Println(fi45, v11, v11)
                                        }
                                    default:
                                        var v104 bool = (-1 < genFn0())
                                        v105 := GenS9{F0: (^(v11))}
                                        _ = v104
                                        _ = v105
                                    }
                                    fmt.Println(fi48, v1, v1)
                                    if ((v1 - v1) <= -1) {
                                        if false {
                                            switch ("hello" + ("hello" + "\n\t")) {
                                            case "\n\t":
                                                var v106 int32 = 42
                                                _ = v106
                                            }
                                            v13 = 127
                                            var v107 bool = true
                                            _ = v107
                                        } else {
                                            var v108 bool = (!false)
                                            var v109 uint8 = v13
                                            _ = v108
                                            _ = v109
                                        }
                                        var v110 uint = 65535
                                        var v111 int8 = -97
                                        var v112 int64 = v1
                                        var v113 uint8 = 1
                                        _ = v110
                                        _ = v111
                                        _ = v112
                                        _ = v113
                                    } else {
                                        fmt.Println(v11)
                                        var v114 uint8 = (v13 << ((0) & 63))
                                        switch (fi45 << ((65535) & 63)) {
                                        case -2147483648:
                                            v115 := GenS1{F0: false, F1: v114, F2: 311}
                                            _ = v115
                                        case -9223372036854775808:
                                            var v116 int32 = 2147483647
                                            _ = v116
                                        }
                                        _ = v114
                                    }
                                }
                                var v117 uint64 = (v3 - v3)
                                v118 := GenS10{F0: (!(v4 >= float64((1e10 / -2.25))))}
                                _ = v117
                                _ = v118
                            }
                            v13 = v13
                            v119 := GenS11{F0: fi48}
                            if len(v34) >= 2 {
                                v34 = v34[1:len(v34)]
                            }
                            if (true && ((v1 >> ((4294967295) & 63)) > 1)) {
                                v1 = (^(v1))
                                v120 := GenS12{F0: v13}
                                v11 = 1
                                _ = v120
                            } else {
                                v121 := GenS13{F0: (!true), F1: v13}
                                var v122 float64 = float64(0.1)
                                _ = v121
                                _ = v122
                            }
                            delete(v2, (^(fi48)))
                            _ = v119
                        }
                        v123 := v2[127]
                        switch "" {
                        case "":
                            switch (-(v123)) {
                            case -474:
                                var v124 float32 = 1e10
                                _ = v124
                            }
                        }
                        _ = fi48
                        _ = v123
                    }
                    v4 -= (v4 - v4)
                }
                if len(v34) > 0 {
                    v34[0] = (fi45 ^ (fi45 - (fi45 + fi45)))
                }
                _ = fi45
            }
        }
    }
    var v125 int = v11
    if (true || true) {
        v11 += v125
    } else {
        for fi126 := 0; fi126 < 11; fi126++ {
            v4 += ((v4 - v4) / v4)
            _ = fi126
        }
        fmt.Println(v3, v4)
        if (("foo bar" + "foo bar") == "\n\t") {
            switch (v11 / ((v11) | 1)) {
            case -358:
                v34 = append(v34, v11)
            case 307:
                v34 = append(v34, -750)
            case -2147483648:
                v11 += (v125 << ((416) & 63))
            default:
                if len(v34) > 0 {
                    v34[0] = (v11 ^ (v125 & (-(v11))))
                }
            }
            switch "" {
            case "":
                v4 = 1e10
            }
        } else {
            var v127 int16 = 1000
            if ((true && ((0 > 4294967295) || false)) || false) {
                v128 := v2[(v11 + (v125 & 544))]
                v129 := GenS6{F0: "foo bar"}
                _ = v128
                _ = v129
            } else {
                switch v125 {
                case 1:
                    v4 = v4
                case 9223372036854775807:
                    if ((0 != 4294967295) && false) {
                        var v130 uint32 = uint32(v3)
                        for fi131 := 0; fi131 < 14; fi131++ {
                            if (-76 == (v11 << ((0) & 63))) {
                                if (false || (-491 != -462)) {
                                    if (!((!false) || false)) {
                                        var v132 int8 = 0
                                        _ = v132
                                    }
                                    v4 = v4
                                    switch ("\n\t" + "hello") {
                                    case "\n\t":
                                        if (!true) {
                                            if (-843 < 42) {
                                                v133 := GenS11{F0: (v125 + (v11 ^ v125))}
                                                fmt.Println(v127)
                                                v134 := v2[4294967295]
                                                _ = v133
                                                _ = v134
                                            } else {
                                                _ = len(v2)
                                                var v135 string = ("\n\t" + "foo bar")
                                                _ = v135
                                            }
                                        }
                                    case "a":
                                        fmt.Println(v13, v3, v125)
                                    case "":
                                        var v136 int = v11
                                        _ = v136
                                    }
                                } else {
                                    var v137 int16 = genFn0()
                                    fmt.Println(v127, v127, v3)
                                    _ = v137
                                }
                                var v138 int8 = 0
                                _ = v138
                            } else {
                                var v139 string = "foo bar"
                                var v140 uint64 = (v3 ^ 304)
                                if len(v34) >= 2 {
                                    v34 = v34[1:len(v34)]
                                }
                                _ = v139
                                _ = v140
                            }
                            _ = fi131
                        }
                        _ = v130
                    } else {
                        if false {
                            fmt.Println(v4)
                        } else {
                            if (("" + "hello") < (("foo bar" + "a") + (("hello" + "foo bar") + "a"))) {
                                v3 += (v3 << ((58) & 63))
                            } else {
                                switch v11 {
                                case 2147483647:
                                    switch v125 {
                                    case 1:
                                        var v141 string = "hello"
                                        _ = v141
                                    case 127:
                                        var v142 bool = true
                                        _ = v142
                                    }
                                }
                                switch 2147483647 {
                                case -1:
                                    if (-6 >= 42) {
                                        var v143 string = "foo bar"
                                        _ = v143
                                    } else {
                                        v13 -= v13
                                        if true {
                                            v13 = 231
                                        } else {
                                            v125 += (v125 / ((v125) | 1))
                                        }
                                        fmt.Println(v127, v11)
                                    }
                                case -528:
                                    var v144 int8 = -128
                                    _ = v144
                                case 32768:
                                    v125 += v125
                                default:
                                    for fi145 := 0; fi145 < 9; fi145++ {
                                        switch 9223372036854775807 {
                                        case 502:
                                            v3 = v3
                                        case 9223372036854775807:
                                            switch (^(fi145)) {
                                            case 128:
                                                switch ("" + "foo bar") {
                                                case "":
                                                    v1 += v1
                                                case "\n\t":
                                                    v146 := GenS14{F0: (^(v3)), F1: (v4 != v4), F2: 0}
                                                    _ = v146
                                                }
                                            case -2147483648:
                                                switch "\n\t" {
                                                case "":
                                                    fmt.Println(fi145)
                                                case "a":
                                                    var v147 float32 = (1e10 + 1.5)
                                                    _ = v147
                                                default:
                                                    var v148 int32 = 0
                                                    _ = v148
                                                }
                                            }
                                        case 32767:
                                            v149 := GenS15{F0: 113, F1: true, F2: ("hello" + "hello")}
                                            _ = v149
                                        default:
                                            if (32767 > 42) {
                                                for fi150 := 0; fi150 < 3; fi150++ {
                                                    v13 += (v13 * (v13 ^ v13))
                                                    switch ((("a" + "") + "") + "a") {
                                                    case "\n\t":
                                                        if len(v34) > 0 {
                                                            v34[0] = v125
                                                        }
                                                    }
                                                    var v151 bool = true
                                                    _ = fi150
                                                    _ = v151
                                                }
                                                v152 := GenS4{F0: -473, F1: 38}
                                                if (true && false) {
                                                    v127 = v127
                                                    v3 += (^(v3))
                                                } else {
                                                    v125 += 0
                                                    v153 := GenS7{F0: v13, F1: 4, F2: 782}
                                                    fmt.Println(fi145)
                                                    _ = v153
                                                }
                                                _ = v152
                                            }
                                        }
                                        v154 := GenS9{F0: fi145}
                                        _ = fi145
                                        _ = v154
                                    }
                                }
                                fmt.Println(v3, v125, v13)
                            }
                            var v155 uint32 = 885
                            var v156 uint64 = v3
                            _ = v155
                            _ = v156
                        }
                        var v157 string = "\n\t"
                        _ = v157
                    }
                case -9223372036854775808:
                    var v158 float64 = 1e10
                    _ = v158
                }
                if false {
                    fmt.Println(v11)
                    if (!true) {
                        if (false && true) {
                            for fi159 := 0; fi159 < 18; fi159++ {
                                var v160 uint64 = (v3 << ((42) & 63))
                                v161 := v2[-9223372036854775808]
                                _ = fi159
                                _ = v160
                                _ = v161
                            }
                            var v162 uint64 = 18446744073709551615
                            if false {
                                var v163 string = ("" + "foo bar")
                                v1 -= v1
                                if true {
                                    v164 := GenS5{F0: 2}
                                    for fi165 := 0; fi165 < 11; fi165++ {
                                        v11 = 42
                                        _ = fi165
                                    }
                                    _ = v164
                                } else {
                                    v166 := GenS14{F0: v162, F1: (65535 >= 32767), F2: 717}
                                    fmt.Println(v163)
                                    _ = v166
                                }
                                _ = v163
                            }
                            _ = v162
                        }
                        for fi167 := 0; fi167 < 14; fi167++ {
                            var v168 int = v125
                            var v169 int8 = -128
                            if len(v34) > 0 {
                                v34[0] = (-(v11))
                            }
                            _ = fi167
                            _ = v168
                            _ = v169
                        }
                        v11 += 1
                    } else {
                        for fi170 := 0; fi170 < 7; fi170++ {
                            switch ("foo bar" + ("\n\t" + "hello")) {
                            case "foo bar":
                                var v171 string = "hello"
                                _ = v171
                            case "hello":
                                var v172 uint16 = 0
                                _ = v172
                            case "":
                                v125 = -677
                            }
                            var v173 bool = false
                            if v173 {
                                for fi174 := 0; fi174 < 8; fi174++ {
                                    switch 9223372036854775807 {
                                    case -1:
                                        if false {
                                            v175 := GenS16{F0: 127, F1: 42}
                                            var v176 int = -212
                                            v4 = -2.25
                                            _ = v175
                                            _ = v176
                                        }
                                    case 4294967295:
                                        fmt.Println(v4, v11, v4)
                                    case 0:
                                        _ = len(v2)
                                    default:
                                        _ = len(v2)
                                    }
                                    v11 = -2147483648
                                    _ = fi174
                                }
                            }
                            _ = fi170
                            _ = v173
                        }
                    }
                    var v177 int16 = (v127 & v127)
                    _ = v177
                }
            }
            var v178 int16 = (v127 + genFn0())
            _ = v127
            _ = v178
        }
    }
    v13 = v13
    var v179 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v125:", v125)
    fmt.Println("v13:", v13)
    fmt.Println("v179:", v179)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v125
    _ = v13
    _ = v179
    _ = v2
    _ = v23
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
