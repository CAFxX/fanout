package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 bool; F2 int32 }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 uint8; F1 int64; F2 string }
type GenS7 struct { F0 uint; F1 bool; F2 int }
type GenS8 struct { F0 int; F1 int8; F2 uint32 }
type GenS9 struct { F0 bool; F1 int16 }
type GenS10 struct { F0 int8; F1 bool; F2 uint32 }
type GenS11 struct { F0 uint32; F1 uint16; F2 int }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int64; F1 bool; F2 int }
type GenS14 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS15 struct { F0 int; F1 string; F2 uint32 }
type GenS16 struct { F0 string; F1 int8; F2 bool }
type GenS17 struct { F0 uint32; F1 bool; F2 int8 }
type GenS18 struct { F0 uint; F1 uint8 }
type GenS19 struct { F0 bool }
type GenS20 struct { F0 int16 }
type GenS21 struct { F0 int8; F1 uint32 }
type GenS22 struct { F0 int16; F1 int32 }
type GenS23 struct { F0 uint; F1 int64; F2 uint16 }
func genFn0() int16 {
    for fi5 := 0; fi5 < 4; fi5++ {
        for fi6 := 0; fi6 < 1; fi6++ {
            v7 := GenS1{F0: 127, F1: 0, F2: 0}
            fmt.Println(fi6, fi6)
            v8 := GenS2{F0: 42}
            _ = fi6
            _ = v7
            _ = v8
        }
        switch "" {
        case "":
            for fi9 := 0; fi9 < 7; fi9++ {
                fmt.Println(fi9)
                if true {
                    var v10 int16 = 0
                    var v11 int16 = v10
                    if true {
                        v11 += v11
                        if false {
                            v12 := []int{fi9}
                            v12 = append(v12, fi5)
                            var v13 int32 = -2147483648
                            fmt.Println(v11)
                            _ = v12
                            _ = v13
                        }
                        var v14 uint32 = 2
                        _ = v14
                    }
                    _ = v10
                    _ = v11
                } else {
                    var v15 float64 = 1e10
                    v16 := []int{}
                    _ = v15
                    _ = v16
                }
                _ = fi9
            }
        case "hello":
            v17 := GenS2{F0: 0}
            _ = v17
        default:
            for fi18 := 0; fi18 < 7; fi18++ {
                switch (fi18 >> ((674) & 63)) {
                case 9223372036854775807:
                    fmt.Println(fi5, fi18)
                case -1:
                    var v19 uint = 255
                    _ = v19
                default:
                    if (2147483647 < 2) {
                        var v20 float64 = float64(1.5)
                        var v21 uint8 = 161
                        var v22 float64 = 1.5
                        _ = v20
                        _ = v21
                        _ = v22
                    } else {
                        if false {
                            fmt.Println(fi18, fi5)
                        } else {
                            if false {
                                var v23 int = fi18
                                var v24 uint32 = 1
                                switch "a" {
                                case "a":
                                    switch fi5 {
                                    case 127:
                                        switch (v23 & fi18) {
                                        case 32767:
                                            if (1.5 >= -2.25) {
                                                fmt.Println(fi5)
                                            } else {
                                                v25 := []int{(fi18 << ((uint(v24)) & 63)), (fi5 % (((fi18 << ((18446744073709551615) & 63))) | 1))}
                                                v26 := GenS1{F0: 0, F1: 1, F2: 255}
                                                v24 -= (v24 << ((65535) & 63))
                                                var v27 bool = true
                                                _ = v25
                                                _ = v26
                                                _ = v27
                                            }
                                        default:
                                            v28 := []int{fi18, fi18, 2147483647}
                                            if (0 <= 1) {
                                                v29 := GenS3{F0: (1.5 <= (-2.25 * 1e10))}
                                                if true {
                                                    v24 = 4294967295
                                                    v24 += 1
                                                    v24 = v24
                                                } else {
                                                    var v30 int16 = -1
                                                    if (false || (("\n\t" + "") >= (("\n\t" + "hello") + ""))) {
                                                        switch ("foo bar" + "") {
                                                        case "foo bar":
                                                            switch "" {
                                                            case "\n\t":
                                                                if len(v28) > 0 {
                                                                    v28[0] = 32767
                                                                }
                                                            default:
                                                                var v31 string = "\n\t"
                                                                _ = v31
                                                            }
                                                        default:
                                                            if (!true) {
                                                                fmt.Println(fi18)
                                                            }
                                                        }
                                                    } else {
                                                        if true {
                                                            v30 = (v30 & v30)
                                                            switch ("a" + "hello") {
                                                            case "":
                                                                v32 := GenS1{F0: 12, F1: 711, F2: 112}
                                                                _ = v32
                                                            case "hello":
                                                                var v33 int64 = 9223372036854775807
                                                                _ = v33
                                                            case "foo bar":
                                                                if ((true || (true || (v30 == (v30 & -32768)))) && false) {
                                                                    if len(v28) >= 2 {
                                                                        v28 = v28[1:len(v28)]
                                                                    }
                                                                    fmt.Println(fi18, v30, v30)
                                                                    fmt.Println(v24, v30)
                                                                    _ = len(v28)
                                                                } else {
                                                                    if true {
                                                                        if true {
                                                                            v23 = fi18
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v24, fi5)
                                                                        _ = len(v28)
                                                                        v34 := GenS4{F0: "", F1: (1e10 > float64((((0.1 * (float32(((-2.25 + 1.5) + 1.5)) - 3.14159)) / 1e10) * (float32(3.14159) + float32((float64(3.14159) - 3.14159)))))), F2: 42}
                                                                        _ = v34
                                                                    }
                                                                    fmt.Println(v24)
                                                                    v35 := GenS5{F0: 2, F1: v30}
                                                                    v28 = append(v28, (-(fi5)))
                                                                    _ = v35
                                                                }
                                                            }
                                                        } else {
                                                            v23 = 4294967295
                                                            v23 += (v23 * v23)
                                                        }
                                                        v28 = append(v28, fi18)
                                                        fmt.Println(fi5, v23)
                                                        if (true || true) {
                                                            var v36 int16 = v30
                                                            _ = v36
                                                        } else {
                                                            if true {
                                                                var v37 uint64 = 349
                                                                v37 -= 42
                                                                var v38 int16 = v30
                                                                _ = v37
                                                                _ = v38
                                                            }
                                                        }
                                                        var v39 float32 = (1.5 / 1.5)
                                                        _ = v39
                                                    }
                                                    _ = v30
                                                }
                                                v40 := GenS6{F0: 16, F1: 1, F2: "foo bar"}
                                                if true {
                                                    fmt.Println(fi5)
                                                    var v41 uint64 = 42
                                                    v23 = fi5
                                                    _ = v41
                                                }
                                                _ = v29
                                                _ = v40
                                            }
                                            _ = v28
                                        }
                                    case 1:
                                        v23 -= (v23 ^ v23)
                                    }
                                case "foo bar":
                                    var v42 string = "foo bar"
                                    _ = v42
                                case "\n\t":
                                    var v43 int8 = 30
                                    _ = v43
                                }
                                _ = v23
                                _ = v24
                            }
                        }
                        v44 := GenS6{F0: 208, F1: -9223372036854775808, F2: (("\n\t" + "hello") + "")}
                        _ = v44
                    }
                }
                fmt.Println(fi5, fi18)
                _ = fi18
            }
        }
        _ = fi5
    }
    v45 := GenS2{F0: 1000}
    for fi46 := 0; fi46 < 13; fi46++ {
        if true {
            v47 := GenS7{F0: 4294967295, F1: (false || ((!true) || (fi46 > fi46))), F2: fi46}
            switch fi46 {
            case -9223372036854775808:
                for fi48 := 0; fi48 < 9; fi48++ {
                    var v49 float64 = 1.5
                    v49 += (v49 * v49)
                    v50 := GenS8{F0: 42, F1: 100, F2: 832}
                    _ = fi48
                    _ = v49
                    _ = v50
                }
            default:
                v51 := []int{522, (fi46 << ((255) & 63)), 4294967295, (-(fi46))}
                _ = v51
            }
            _ = v47
        } else {
            v52 := GenS9{F0: (!true), F1: -1}
            v53 := GenS6{F0: 19, F1: -237, F2: "foo bar"}
            fmt.Println(fi46)
            _ = v52
            _ = v53
        }
        _ = fi46
    }
    for fi54 := 0; fi54 < 3; fi54++ {
        var v55 uint = 0
        var v56 int8 = 1
        if (true || (-2.25 == 0.0)) {
            if (!false) {
                fmt.Println(fi54, fi54, v55)
                v55 -= (v55 * (v55 ^ 0))
            } else {
                var v57 int32 = -2147483648
                v56 += 42
                _ = v57
            }
            var v58 int32 = -1
            if false {
                v59 := []int{}
                _ = v59
            }
            _ = v58
        } else {
            var v60 uint64 = 2
            for fi61 := 0; fi61 < 15; fi61++ {
                var v62 int32 = 696
                var v63 int32 = v62
                _ = fi61
                _ = v62
                _ = v63
            }
            _ = v60
        }
        _ = fi54
        _ = v55
        _ = v56
    }
    _ = v45
    return -882
}

func main() {
    v1 := []string{}
    var v2 int32 = 2147483647
    v3 := []int{1, 193, 42, 4294967295}
    v4 := make(map[uint32]uint32)
    v2 = v2
    v2 += -771
    v64 := GenS3{F0: (1 != 967)}
    for fi65 := 0; fi65 < 17; fi65++ {
        v66 := v4[1]
        for fi67 := 0; fi67 < 15; fi67++ {
            if false {
                var v68 float32 = ((3.14159 / (3.14159 / -2.25)) - (0.0 - (1.5 * (1.5 + 3.14159))))
                if false {
                    var v69 int32 = v2
                    if true {
                        if (!((false && (292 > -9223372036854775808)) || true)) {
                            _ = len(v3)
                            var v70 float64 = 1.5
                            v66 += v66
                            _ = v70
                        } else {
                            v71 := GenS9{F0: (!("\n\t" <= ("hello" + "hello"))), F1: genFn0()}
                            if true {
                                v68 = float32((1.5 * (((3.14159 - 1.5) + 0.1) * (0.1 / 3.14159))))
                                v66 = (v66 - (v66 * (v66 - (v66 - (v66 | (v66 + 2147483648))))))
                            } else {
                                v72 := GenS10{F0: -111, F1: true, F2: 1}
                                v73 := v4[65535]
                                switch fi67 {
                                case 106:
                                    switch ("hello" + "") {
                                    case "a":
                                        switch "a" {
                                        case "hello":
                                            v74 := GenS5{F0: 2, F1: 32767}
                                            _ = v74
                                        case "\n\t":
                                            var v75 string = "hello"
                                            var v76 int8 = 42
                                            _ = v75
                                            _ = v76
                                        case "":
                                            if ((92 <= 100) || true) {
                                                var v77 uint32 = (v73 ^ 20)
                                                v66 = (^(v77))
                                                _ = v77
                                            } else {
                                                var v78 int = (-(fi65))
                                                v73 += (v66 & (v73 * v73))
                                                _ = v78
                                            }
                                        }
                                    case "foo bar":
                                        v2 = (v69 >> ((255) & 63))
                                        v79 := GenS3{F0: (true && (625 != 968))}
                                        _ = v79
                                    case "hello":
                                        v80 := GenS11{F0: v66, F1: 947, F2: fi65}
                                        _ = v80
                                    }
                                case 696:
                                    delete(v4, v66)
                                default:
                                    fmt.Println(v2)
                                }
                                _ = v72
                                _ = v73
                            }
                            _ = v71
                        }
                        var v81 uint8 = 0
                        v68 = v68
                        _ = v81
                    } else {
                        if (732 == 2) {
                            v82 := GenS9{F0: (-100 < 1), F1: 439}
                            _ = v82
                        } else {
                            switch (fi67 - (fi67 * fi67)) {
                            case 1:
                                var v83 int8 = -16
                                _ = v83
                            case 908:
                                v4[(v66 & v66)] = (v66 >> ((2) & 63))
                            case -447:
                                if (15 != 185) {
                                    switch "foo bar" {
                                    case "a":
                                        switch "" {
                                        case "":
                                            switch (("a" + "foo bar") + ("\n\t" + "a")) {
                                            case "a":
                                                if ((!(false && (false && (!false)))) || true) {
                                                    v68 = v68
                                                    var v84 uint64 = 0
                                                    _ = v84
                                                }
                                            }
                                        case "hello":
                                            v2 -= (v2 >> ((255) & 63))
                                        }
                                    case "":
                                        switch (fi65 >> ((59) & 63)) {
                                        case 128:
                                            switch "\n\t" {
                                            case "hello":
                                                v2 -= (v69 << ((65535) & 63))
                                            case "\n\t":
                                                if false {
                                                    v85 := GenS12{F0: "\n\t"}
                                                    _ = v85
                                                }
                                            case "":
                                                var v86 uint = 183
                                                _ = v86
                                            default:
                                                v87 := GenS13{F0: 0, F1: ((42 < 2147483647) || (9223372036854775807 != 0)), F2: 9223372036854775807}
                                                _ = v87
                                            }
                                        case -685:
                                            switch (fi67 - (fi65 / ((fi65) | 1))) {
                                            case -1:
                                                fmt.Println(fi65, v69)
                                            }
                                        case 42:
                                            v88 := GenS6{F0: 128, F1: 1, F2: ("foo bar" + ("\n\t" + ("hello" + "")))}
                                            _ = v88
                                        default:
                                            if len(v1) > 0 {
                                                v1[0] = ""
                                            }
                                        }
                                    case "foo bar":
                                        if true {
                                            fmt.Println(fi67, fi67, v2)
                                            switch (fi65 - fi67) {
                                            case 9223372036854775807:
                                                v89 := GenS14{F0: 42, F1: 0, F2: 42}
                                                _ = v89
                                            }
                                            v68 = (v68 - v68)
                                        }
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    fmt.Println(fi67)
                                } else {
                                    var v90 int16 = 88
                                    _ = v90
                                }
                            default:
                                var v91 uint = 206
                                _ = v91
                            }
                        }
                        switch ("a" + "") {
                        case "\n\t":
                            var v92 uint16 = 42
                            _ = v92
                        case "":
                            var v93 int64 = 339
                            _ = v93
                        default:
                            v2 += v2
                        }
                    }
                    var v94 uint64 = 410
                    if (int8(fi67) < 52) {
                        fmt.Println(v94, v2, fi65)
                        var v95 uint = 4294967295
                        _ = v95
                    } else {
                        v66 = 881
                    }
                    _ = v69
                    _ = v94
                }
                v96 := v4[1]
                _ = v68
                _ = v96
            }
            var v97 int16 = genFn0()
            if (uint64(v66) < 682) {
                v98 := GenS15{F0: (fi67 << ((42) & 63)), F1: "\n\t", F2: (v66 ^ (v66 * (^(v66))))}
                _ = v98
            }
            switch fi65 {
            case -188:
                delete(v4, (v66 << ((6) & 63)))
            case 1:
                _ = len(v3)
            case 5:
                v66 += (v66 ^ (v66 >> ((717) & 63)))
            }
            fmt.Println(v97)
            _ = fi67
            _ = v97
        }
        _ = fi65
        _ = v66
    }
    v2 -= v2
    for fi99 := 0; fi99 < 5; fi99++ {
        var v100 bool = (100 >= -1)
        fmt.Println(v2, v2, v100)
        _ = fi99
        _ = v100
    }
    switch -1 {
    case 4294967295:
        if (!false) {
            v2 = v2
        } else {
            fmt.Println(v2)
            fmt.Println(v2)
        }
    default:
        if false {
            for fi101 := 0; fi101 < 3; fi101++ {
                var v102 int16 = -465
                for fi103 := 0; fi103 < 6; fi103++ {
                    var v104 int = (fi103 / ((fi101) | 1))
                    if true {
                        v102 = -780
                        v102 = int16(v104)
                    } else {
                        var v105 int = (v104 | v104)
                        v106 := GenS14{F0: -2147483648, F1: 4294967295, F2: 128}
                        _ = v105
                        _ = v106
                    }
                    _ = fi103
                    _ = v104
                }
                _ = fi101
                _ = v102
            }
            var v107 int8 = -1
            var v108 uint64 = 129
            _ = v107
            _ = v108
        }
    }
    var v109 uint = 2
    var v110 bool = false
    v109 -= (^(v109))
    var v111 int = 4294967295
    v112 := GenS16{F0: "hello", F1: int8(v109), F2: v110}
    v4[2147483647] = 0
    if len(v3) > 0 {
        v3[0] = v111
    }
    if v110 {
        if v110 {
            if ("\n\t" == ("a" + "\n\t")) {
                if true {
                    fmt.Println(v2)
                    if v110 {
                        for fi113 := 0; fi113 < 3; fi113++ {
                            for fi114 := 0; fi114 < 3; fi114++ {
                                var v115 uint32 = 2147483647
                                _ = fi114
                                _ = v115
                            }
                            v109 -= (v109 ^ v109)
                            _ = fi113
                        }
                    }
                    fmt.Println(v2)
                }
                for fi116 := 0; fi116 < 11; fi116++ {
                    if v110 {
                        delete(v4, 42)
                    }
                    _ = fi116
                }
            }
        } else {
            v109 = v109
            var v117 int = (^(v111))
            switch "a" {
            case "a":
                v118 := GenS8{F0: (v111 - (v111 | (v111 % (((v111 & (v117 & v117))) | 1)))), F1: -124, F2: 1}
                _ = v118
            case "foo bar":
                _ = len(v3)
            case "":
                v119 := GenS6{F0: 136, F1: 867, F2: "foo bar"}
                _ = v119
            }
            _ = v117
        }
    } else {
        for fi120 := 0; fi120 < 13; fi120++ {
            for fi121 := 0; fi121 < 15; fi121++ {
                v109 = (v109 & (^(v109)))
                var v122 int16 = -359
                v1 = append(v1, "hello")
                _ = fi121
                _ = v122
            }
            for fi123 := 0; fi123 < 15; fi123++ {
                var v124 uint8 = 1
                delete(v4, 0)
                var v125 int8 = 42
                v125 = -95
                v109 = (v109 & 65535)
                _ = fi123
                _ = v124
                _ = v125
            }
            v110 = true
            _ = fi120
        }
        for fi126 := 0; fi126 < 19; fi126++ {
            fmt.Println(v109, v111, v109)
            if (!(2 == 445)) {
                v111 += int(v109)
                if (v110 || v110) {
                    _ = len(v4)
                    switch "\n\t" {
                    case "":
                        v111 = 481
                    default:
                        switch "foo bar" {
                        case "a":
                            if ((4294967295 != 4294967295) && false) {
                                for fi127 := 0; fi127 < 4; fi127++ {
                                    var v128 uint16 = 1
                                    switch "\n\t" {
                                    case "":
                                        v110 = true
                                        _ = len(v3)
                                        switch ("\n\t" + ("hello" + "\n\t")) {
                                        case "a":
                                            var v129 string = ""
                                            _ = v129
                                        case "foo bar":
                                            fmt.Println(v110, fi127)
                                        default:
                                            switch fi126 {
                                            case -9223372036854775808:
                                                switch "hello" {
                                                case "foo bar":
                                                    v110 = true
                                                case "\n\t":
                                                    fmt.Println(v109)
                                                default:
                                                    var v130 uint32 = 42
                                                    _ = v130
                                                }
                                            case 32767:
                                                v131 := GenS17{F0: 2147483648, F1: ((-999 != 42) && ((!v110) && (100 <= 127))), F2: 42}
                                                _ = v131
                                            case 1:
                                                if (((32767 >= (v111 & -489)) || false) || v110) {
                                                    v132 := GenS16{F0: ("" + ""), F1: -54, F2: v110}
                                                    v111 = -616
                                                    _ = v132
                                                }
                                            }
                                        }
                                    case "foo bar":
                                        if (((3.14159 + float64((float32(0.0) / 3.14159))) > 0.0) || v110) {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        } else {
                                            if false {
                                                if v110 {
                                                    fmt.Println(v2, fi126, v109)
                                                    var v133 float64 = 3.14159
                                                    v111 = -49
                                                    fmt.Println(v111, v128, fi127)
                                                    _ = v133
                                                }
                                                switch "hello" {
                                                case "\n\t":
                                                    if v110 {
                                                        fmt.Println(fi126)
                                                    }
                                                case "a":
                                                    var v134 uint16 = v128
                                                    _ = v134
                                                case "hello":
                                                    var v135 uint8 = 67
                                                    _ = v135
                                                }
                                                v111 = int(v109)
                                                if true {
                                                    if ((^(fi127)) != v111) {
                                                        _ = len(v1)
                                                    } else {
                                                        switch (fi126 / (((^(fi126))) | 1)) {
                                                        case -869:
                                                            var v136 uint = v109
                                                            _ = v136
                                                        }
                                                        if len(v1) > 0 {
                                                            v1[0] = "a"
                                                        }
                                                    }
                                                    v2 -= v2
                                                    v111 = v111
                                                    switch int(v109) {
                                                    case -375:
                                                        if (!(-1 >= v2)) {
                                                            var v137 int = (-(v111))
                                                            v138 := GenS2{F0: -796}
                                                            v139 := GenS6{F0: 127, F1: -9223372036854775808, F2: "a"}
                                                            _ = v137
                                                            _ = v138
                                                            _ = v139
                                                        }
                                                    case 453:
                                                        fmt.Println(v109, v109, fi127)
                                                    }
                                                    var v140 uint32 = 4294967295
                                                    switch 127 {
                                                    case -2147483648:
                                                        if len(v3) > 0 {
                                                            v3[0] = (^(v111))
                                                        }
                                                    case 128:
                                                        if v110 {
                                                            v110 = false
                                                            var v141 int32 = -1
                                                            _ = v141
                                                        } else {
                                                            v142 := GenS18{F0: (v109 % (((^(v109))) | 1)), F1: 255}
                                                            v3 = append(v3, fi127)
                                                            v109 -= 1
                                                            _ = v142
                                                        }
                                                    case 32768:
                                                        if (!v110) {
                                                            var v143 uint = (v109 % ((1) | 1))
                                                            if ((^(v140)) <= (v140 ^ (v140 ^ (v140 & v140)))) {
                                                                v143 -= 42
                                                                var v144 float32 = 1.5
                                                                var v145 string = ("" + ("\n\t" + "a"))
                                                                _ = v144
                                                                _ = v145
                                                            } else {
                                                                v109 += (v109 / ((v109) | 1))
                                                                switch "" {
                                                                case "a":
                                                                    fmt.Println(v110)
                                                                default:
                                                                    v146 := v4[(v140 - (v140 - v140))]
                                                                    _ = v146
                                                                }
                                                                var v147 uint16 = v128
                                                                _ = v147
                                                            }
                                                            var v148 uint = (v143 & (v109 - v143))
                                                            _ = v143
                                                            _ = v148
                                                        } else {
                                                            fmt.Println(v110)
                                                            v4[(v140 * v140)] = v140
                                                        }
                                                    }
                                                    _ = v140
                                                }
                                                if (2147483648 < 42) {
                                                    v149 := GenS19{F0: v110}
                                                    _ = v149
                                                } else {
                                                    delete(v4, 2147483647)
                                                    var v150 float64 = 0.1
                                                    var v151 uint16 = v128
                                                    _ = v150
                                                    _ = v151
                                                }
                                                var v152 int = fi127
                                                var v153 uint32 = 65535
                                                _ = v152
                                                _ = v153
                                            } else {
                                                fmt.Println(fi126, v111, v109)
                                                v154 := GenS3{F0: false}
                                                var v155 uint8 = 39
                                                _ = v154
                                                _ = v155
                                            }
                                            v156 := GenS20{F0: 42}
                                            v128 -= v128
                                            _ = v156
                                        }
                                    default:
                                        fmt.Println(v109, v128, v109)
                                    }
                                    v157 := GenS13{F0: -968, F1: v110, F2: fi126}
                                    _ = fi127
                                    _ = v128
                                    _ = v157
                                }
                                var v158 uint32 = 870
                                _ = len(v3)
                                _ = v158
                            } else {
                                if (-1 > -1) {
                                    var v159 int32 = (v2 - (-(v2)))
                                    var v160 bool = (127 > 128)
                                    for fi161 := 0; fi161 < 12; fi161++ {
                                        var v162 int = (^(fi126))
                                        _ = fi161
                                        _ = v162
                                    }
                                    _ = v159
                                    _ = v160
                                } else {
                                    if (!v110) {
                                        v163 := GenS19{F0: false}
                                        _ = len(v4)
                                        var v164 int64 = 42
                                        _ = v163
                                        _ = v164
                                    }
                                    var v165 int32 = 695
                                    _ = v165
                                }
                                delete(v4, 213)
                                var v166 uint = v109
                                _ = v166
                            }
                        case "\n\t":
                            var v167 int32 = (^(v2))
                            _ = v167
                        case "foo bar":
                            var v168 uint64 = 1
                            _ = v168
                        default:
                            var v169 int = fi126
                            _ = v169
                        }
                    }
                    var v170 int16 = 1000
                    _ = v170
                } else {
                    v109 = (v109 ^ 42)
                    var v171 uint64 = 375
                    _ = v171
                }
                v172 := GenS21{F0: 100, F1: 1}
                _ = v172
            } else {
                _ = len(v4)
                v2 = v2
                if (!false) {
                    var v173 int = 128
                    var v174 int = (v111 | (v111 & v173))
                    v175 := GenS12{F0: "a"}
                    _ = v173
                    _ = v174
                    _ = v175
                }
            }
            var v176 int32 = (v2 / ((-551) | 1))
            _ = fi126
            _ = v176
        }
        if v110 {
            v177 := GenS22{F0: genFn0(), F1: 42}
            var v178 uint32 = 4294967295
            for fi179 := 0; fi179 < 12; fi179++ {
                v180 := GenS20{F0: -1}
                fmt.Println(v2)
                if v110 {
                    v181 := GenS16{F0: "hello", F1: 1, F2: true}
                    for fi182 := 0; fi182 < 4; fi182++ {
                        switch "hello" {
                        case "foo bar":
                            v110 = false
                        }
                        _ = fi182
                    }
                    _ = v181
                } else {
                    var v183 string = "foo bar"
                    for fi184 := 0; fi184 < 8; fi184++ {
                        if false {
                            v183 = v183
                        }
                        v110 = ((float64((0.1 * 3.14159)) + 1e10) < -2.25)
                        if false {
                            v185 := GenS23{F0: v109, F1: 636, F2: 307}
                            v109 -= v109
                            v3 = append(v3, (fi179 + (^(fi179))))
                            _ = v185
                        }
                        _ = fi184
                    }
                    var v186 float32 = 1.5
                    _ = v183
                    _ = v186
                }
                _ = fi179
                _ = v180
            }
            _ = v177
            _ = v178
        }
    }
    var v187 uint16 = 0
    v4[2147483648] = 829
    fmt.Println(v111, v110, v109)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v109:", v109)
    fmt.Println("v110:", v110)
    fmt.Println("v111:", v111)
    fmt.Println("v187:", v187)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v109
    _ = v110
    _ = v111
    _ = v112
    _ = v187
    _ = v2
    _ = v3
    _ = v4
    _ = v64
}
