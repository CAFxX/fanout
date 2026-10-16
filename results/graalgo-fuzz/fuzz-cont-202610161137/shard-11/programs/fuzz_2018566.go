package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 int16; F1 uint16; F2 bool }
type GenS5 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS6 struct { F0 uint; F1 uint; F2 int }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS9 struct { F0 bool; F1 uint64; F2 uint64 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 bool; F1 int }
type GenS12 struct { F0 uint8; F1 int64 }
type GenS13 struct { F0 uint16; F1 int8; F2 bool }
type GenS14 struct { F0 int32; F1 int64 }
type GenS15 struct { F0 int; F1 uint64 }
type GenS16 struct { F0 uint8; F1 int8; F2 int32 }
type GenS17 struct { F0 uint }

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = ((3.14159 * -2.25) - 3.14159)
    var v3 int64 = 42
    v4 := GenS1{F0: 65535, F1: -1}
    fmt.Println(v2, v2)
    var v5 int = -9223372036854775808
    v6 := []int{(v5 << ((860) & 63)), v5, (^(v5))}
    _ = len(v1)
    v7 := GenS1{F0: 4294967295, F1: 42}
    switch "foo bar" {
    case "hello":
        if false {
            var v8 uint64 = 128
            fmt.Println(v2)
            _ = v8
        } else {
            fmt.Println(v5, v3)
            switch (("" + "a") + ((("" + "a") + "hello") + "a")) {
            case "":
                if false {
                    fmt.Println(v2, v5, v5)
                } else {
                    var v9 uint8 = 2
                    _ = v9
                }
            case "\n\t":
                var v10 uint8 = 1
                _ = v10
            default:
                if false {
                    v2 = v2
                }
            }
            v1[(v5 ^ 457)] = 42
        }
    case "foo bar":
        for fi11 := 0; fi11 < 10; fi11++ {
            v3 += v3
            _ = fi11
        }
    case "a":
        fmt.Println(v3, v3)
    default:
        switch "hello" {
        case "\n\t":
            if false {
                v1[-562] = 4294967295
                v3 = v3
                switch ("\n\t" + "\n\t") {
                case "a":
                    v12 := GenS2{F0: 1, F1: -2147483648, F2: 2}
                    _ = v12
                case "\n\t":
                    for fi13 := 0; fi13 < 2; fi13++ {
                        for fi14 := 0; fi14 < 4; fi14++ {
                            switch (fi14 * 128) {
                            case -138:
                                v15 := GenS3{F0: 2, F1: 243}
                                _ = v15
                            case -976:
                                v2 = (v2 - v2)
                            case 1:
                                v3 = v3
                                v2 += float32(1.5)
                                if false {
                                    var v16 int64 = (v3 + v3)
                                    delete(v1, -2147483648)
                                    _ = v16
                                } else {
                                    var v17 string = "foo bar"
                                    switch v17 {
                                    case "\n\t":
                                        _ = len(v1)
                                    case "":
                                        v18 := GenS2{F0: 2, F1: -1, F2: 2}
                                        _ = v18
                                    }
                                    v17 = (((v17 + v17) + ((v17 + v17) + ((v17 + v17) + ("\n\t" + "a")))) + v17)
                                    _ = v17
                                }
                            }
                            if false {
                                fmt.Println(fi13)
                                switch 32768 {
                                case -9223372036854775808:
                                    v19 := GenS3{F0: 4294967295, F1: (^(v3))}
                                    _ = v19
                                case 128:
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v20 int16 = 1
                                        _ = v20
                                    case "foo bar":
                                        v21 := GenS2{F0: 18446744073709551615, F1: 42, F2: 0}
                                        _ = v21
                                    case "hello":
                                        var v22 int32 = 2147483647
                                        _ = v22
                                    }
                                case -987:
                                    if (v5 > 9223372036854775807) {
                                        v23 := GenS1{F0: 4, F1: -1}
                                        fmt.Println(v2, v2, v5)
                                        if false {
                                            if ((1e10 >= 0.0) || (true || true)) {
                                                v1[(fi13 * fi14)] = 2147483648
                                                switch (fi14 / ((fi13) | 1)) {
                                                case 32767:
                                                    var v24 bool = (true && (127 > 1))
                                                    _ = v24
                                                case 11:
                                                    var v25 int16 = 1
                                                    _ = v25
                                                case 128:
                                                    v5 -= 32767
                                                default:
                                                    v2 = v2
                                                }
                                                v2 = (1e10 / -2.25)
                                            } else {
                                                var v26 int64 = -9223372036854775808
                                                if (!(!false)) {
                                                    switch "hello" {
                                                    case "":
                                                        fmt.Println(v5, v26)
                                                    case "hello":
                                                        if (42 < 908) {
                                                            if (false && (!true)) {
                                                                v5 += (-(fi13))
                                                                if (2 != 189) {
                                                                    fmt.Println(v5)
                                                                }
                                                            } else {
                                                                v5 = fi13
                                                                var v27 float64 = 3.14159
                                                                v27 -= (v27 * ((v27 + 0.1) * 3.14159))
                                                                _ = v27
                                                            }
                                                        } else {
                                                            var v28 float64 = 0.0
                                                            var v29 int64 = v26
                                                            switch ("" + "a") {
                                                            case "hello":
                                                                fmt.Println(v3, v3, v5)
                                                                var v30 int64 = v26
                                                                if ((v5 >> ((65535) & 63)) == (v5 >> ((56) & 63))) {
                                                                    var v31 float64 = (1e10 / ((v28 * v28) - 0.0))
                                                                    _ = v31
                                                                } else {
                                                                    fmt.Println(v29, v28)
                                                                    v6 = append(v6, v5)
                                                                }
                                                                _ = v30
                                                            case "a":
                                                                switch ("hello" + "") {
                                                                case "":
                                                                    var v32 uint32 = 42
                                                                    _ = v32
                                                                case "\n\t":
                                                                    if true {
                                                                        if false {
                                                                            if ((true && (84 != 1)) && true) {
                                                                                v28 = 1e10
                                                                                var v33 float64 = 1.5
                                                                                switch "" {
                                                                                case "foo bar":
                                                                                    fmt.Println(v29, v2)
                                                                                default:
                                                                                    fmt.Println(v26)
                                                                                }
                                                                                _ = v33
                                                                            }
                                                                        } else {
                                                                            var v34 uint32 = 2
                                                                            var v35 int16 = -1
                                                                            delete(v1, -1)
                                                                            _ = v34
                                                                            _ = v35
                                                                        }
                                                                        v5 = 32768
                                                                    } else {
                                                                        switch "hello" {
                                                                        case "":
                                                                            v36 := v1[fi14]
                                                                            _ = v36
                                                                        }
                                                                        var v37 uint32 = 42
                                                                        _ = v37
                                                                    }
                                                                }
                                                            }
                                                            _ = v28
                                                            _ = v29
                                                        }
                                                        var v38 int32 = 42
                                                        _ = v38
                                                    case "foo bar":
                                                        var v39 int = fi14
                                                        _ = v39
                                                    }
                                                    var v40 int64 = v3
                                                    _ = v40
                                                } else {
                                                    v41 := GenS4{F0: 42, F1: 1, F2: ((!(0.1 == 0.0)) && false)}
                                                    v5 -= fi13
                                                    if ("" != "a") {
                                                        v42 := GenS5{F0: 42, F1: 4294967295, F2: 0}
                                                        _ = v42
                                                    } else {
                                                        switch "hello" {
                                                        case "":
                                                            var v43 float64 = -2.25
                                                            _ = v43
                                                        case "foo bar":
                                                            if false {
                                                                v26 = (-(v3))
                                                                v44 := GenS6{F0: 4294967295, F1: 723, F2: fi13}
                                                                switch ("" + "a") {
                                                                case "hello":
                                                                    switch v5 {
                                                                    case 128:
                                                                        switch v5 {
                                                                        case 632:
                                                                            fmt.Println(v5, v26, fi13)
                                                                        case 903:
                                                                            v45 := GenS1{F0: 42, F1: 1000}
                                                                            _ = v45
                                                                        case 128:
                                                                            var v46 int8 = 8
                                                                            _ = v46
                                                                        default:
                                                                            var v47 bool = false
                                                                            _ = v47
                                                                        }
                                                                    default:
                                                                        var v48 uint8 = 2
                                                                        _ = v48
                                                                    }
                                                                case "":
                                                                    if false {
                                                                        if false {
                                                                            switch fi13 {
                                                                            case -2147483648:
                                                                                switch "a" {
                                                                                case "":
                                                                                    fmt.Println(fi13, v5)
                                                                                }
                                                                            case 72:
                                                                                if true {
                                                                                    if (true && false) {
                                                                                        v3 -= v26
                                                                                    } else {
                                                                                        v49 := GenS7{F0: 65535}
                                                                                        var v50 uint = 4294967295
                                                                                        _ = v49
                                                                                        _ = v50
                                                                                    }
                                                                                    switch ("\n\t" + ("hello" + "a")) {
                                                                                    case "\n\t":
                                                                                        var v51 uint = 1
                                                                                        if false {
                                                                                            var v52 uint = v51
                                                                                            v3 += (^(v26))
                                                                                            _ = v52
                                                                                        }
                                                                                        _ = v51
                                                                                    }
                                                                                }
                                                                            }
                                                                            var v53 int8 = 0
                                                                            _ = v53
                                                                        }
                                                                        v26 = -9223372036854775808
                                                                    }
                                                                default:
                                                                    fmt.Println(fi14, v3)
                                                                    v5 = v5
                                                                    _ = len(v6)
                                                                    fmt.Println(v3)
                                                                    v54 := GenS8{F0: 1, F1: 2, F2: -1}
                                                                    v2 = (1.5 * ((-2.25 - (float32(1e10) * v2)) - v2))
                                                                    _ = v54
                                                                }
                                                                var v55 uint32 = 1
                                                                if ((!true) || (42 < uint16(v55))) {
                                                                    var v56 string = ""
                                                                    fmt.Println(fi13)
                                                                    fmt.Println(v2, v26)
                                                                    _ = v56
                                                                } else {
                                                                    v57 := GenS4{F0: 1, F1: 1, F2: ((!false) || false)}
                                                                    v26 = (v26 ^ -599)
                                                                    _ = v57
                                                                }
                                                                _ = v44
                                                                _ = v55
                                                            }
                                                        }
                                                    }
                                                    _ = v41
                                                }
                                                var v58 uint64 = 4294967295
                                                _ = v26
                                                _ = v58
                                            }
                                            switch "a" {
                                            case "a":
                                                switch -2147483648 {
                                                case -28:
                                                    switch ("" + "") {
                                                    case "a":
                                                        v59 := GenS9{F0: ("a" <= "foo bar"), F1: 18446744073709551615, F2: 244}
                                                        _ = v59
                                                    default:
                                                        v1[(fi14 * fi13)] = 4294967295
                                                    }
                                                }
                                            case "":
                                                if false {
                                                    v1[(fi14 / (((fi13 & (v5 ^ fi13))) | 1))] = 611
                                                    var v60 float32 = v2
                                                    v61 := GenS10{F0: v3}
                                                    v3 -= (v3 / ((v3) | 1))
                                                    _ = v60
                                                    _ = v61
                                                }
                                            case "\n\t":
                                                fmt.Println(v5, fi14)
                                            default:
                                                v3 += 9223372036854775807
                                            }
                                            var v62 uint64 = 4294967295
                                            _ = v62
                                        } else {
                                            if ("\n\t" <= ("a" + "hello")) {
                                                v5 = (-(fi14))
                                                var v63 uint = 1
                                                v3 += v3
                                                _ = v63
                                            }
                                            var v64 int64 = 911
                                            _ = v64
                                        }
                                        v65 := GenS11{F0: (true || ((v3 % ((v3) | 1)) <= 9223372036854775807)), F1: (fi14 * v5)}
                                        v3 -= (^(v3))
                                        _ = v23
                                        _ = v65
                                    } else {
                                        switch ("hello" + "hello") {
                                        case "foo bar":
                                            v66 := GenS9{F0: false, F1: 18446744073709551615, F2: 509}
                                            _ = v66
                                        case "\n\t":
                                            fmt.Println(v5, v5)
                                        default:
                                            v67 := GenS3{F0: 65535, F1: (v3 | v3)}
                                            _ = v67
                                        }
                                        switch fi14 {
                                        case 4294967295:
                                            var v68 int8 = 1
                                            _ = v68
                                        case 714:
                                            v5 += v5
                                            v69 := GenS1{F0: 0, F1: -32768}
                                            _ = len(v6)
                                            _ = v69
                                        }
                                        var v70 uint8 = 255
                                        _ = v70
                                    }
                                }
                            } else {
                                if true {
                                    v71 := GenS12{F0: 226, F1: (v3 ^ 0)}
                                    _ = v71
                                } else {
                                    var v72 int8 = -106
                                    switch (fi13 | (^(v5))) {
                                    case -2147483648:
                                        if false {
                                            v73 := GenS3{F0: 65535, F1: int64(v72)}
                                            var v74 float64 = 1e10
                                            var v75 float64 = v74
                                            delete(v1, 32767)
                                            _ = v73
                                            _ = v74
                                            _ = v75
                                        }
                                    }
                                    _ = v72
                                }
                            }
                            _ = fi14
                        }
                        v1[-430] = 2147483648
                        _ = fi13
                    }
                case "hello":
                    var v76 uint16 = 910
                    _ = v76
                }
            } else {
                var v77 uint32 = 0
                v78 := GenS13{F0: uint16(v77), F1: 127, F2: false}
                if true {
                    v5 += (-(v5))
                    var v79 uint = 4294967295
                    var v80 bool = false
                    _ = v79
                    _ = v80
                } else {
                    v81 := GenS6{F0: 0, F1: 4294967295, F2: (v5 ^ int(v77))}
                    var v82 float64 = 1e10
                    _ = v81
                    _ = v82
                }
                _ = v77
                _ = v78
            }
        case "":
            for fi83 := 0; fi83 < 14; fi83++ {
                var v84 float32 = v2
                switch fi83 {
                case -1:
                    var v85 float32 = 0.1
                    _ = v85
                case -583:
                    var v86 uint32 = 2147483648
                    _ = v86
                case -794:
                    fmt.Println(v2, v3, v3)
                default:
                    switch "hello" {
                    case "a":
                        var v87 int32 = -2147483648
                        _ = v87
                    }
                }
                _ = fi83
                _ = v84
            }
        }
    }
    v2 = (3.14159 / 1.5)
    for fi88 := 0; fi88 < 19; fi88++ {
        v89 := GenS14{F0: -725, F1: v3}
        var v90 int = (^(fi88))
        _ = fi88
        _ = v89
        _ = v90
    }
    v5 = 518
    for fi91 := 0; fi91 < 6; fi91++ {
        for fi92 := 0; fi92 < 3; fi92++ {
            v2 -= float32((0.0 / 1e10))
            _ = fi92
        }
        if (true || false) {
            var v93 uint32 = 160
            _ = v93
        } else {
            v94 := GenS4{F0: 1000, F1: 1, F2: true}
            fmt.Println(v2, fi91, v5)
            _ = v94
        }
        _ = fi91
    }
    switch 32767 {
    case 476:
        fmt.Println(v2, v3, v5)
    case 4294967295:
        if true {
            v2 = v2
            for fi95 := 0; fi95 < 10; fi95++ {
                var v96 uint8 = 42
                var v97 int8 = 127
                v98 := GenS15{F0: (v5 / (((^(fi95))) | 1)), F1: 611}
                _ = fi95
                _ = v96
                _ = v97
                _ = v98
            }
            var v99 int64 = v3
            _ = v99
        }
    case 9223372036854775807:
        _ = len(v6)
    default:
        for fi100 := 0; fi100 < 14; fi100++ {
            _ = len(v6)
            v3 += (v3 + v3)
            var v101 bool = false
            _ = fi100
            _ = v101
        }
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi102 := 0; fi102 < 5; fi102++ {
        for fi103 := 0; fi103 < 2; fi103++ {
            switch fi102 {
            case 9223372036854775807:
                v2 -= 0.1
            case 127:
                if (true || false) {
                    v3 = 996
                    v2 += ((v2 - v2) * v2)
                    fmt.Println(fi103, fi102)
                }
            case -436:
                var v104 int = fi103
                _ = v104
            default:
                if false {
                    v2 += (0.1 + 0.0)
                    v3 = 9223372036854775807
                    if false {
                        v2 = (v2 + float32((1.5 * 1e10)))
                    } else {
                        var v105 uint64 = 4294967295
                        v106 := GenS16{F0: 128, F1: 0, F2: 1}
                        var v107 uint16 = 0
                        _ = v105
                        _ = v106
                        _ = v107
                    }
                    v108 := GenS12{F0: 127, F1: 682}
                    if true {
                        fmt.Println(v5, fi103)
                        v2 -= v2
                        if false {
                            v109 := GenS5{F0: 4294967295, F1: 2147483647, F2: 1}
                            v3 = (^(v3))
                            _ = v109
                        } else {
                            if true {
                                if len(v6) > 0 {
                                    v6[0] = (-(fi102))
                                }
                            } else {
                                if (250 < 275) {
                                    switch "" {
                                    case "\n\t":
                                        v5 = fi103
                                    case "":
                                        var v110 uint32 = 2147483648
                                        _ = v110
                                    case "a":
                                        var v111 uint64 = 894
                                        _ = v111
                                    }
                                } else {
                                    v5 = (-(v5))
                                    var v112 int8 = 4
                                    _ = v112
                                }
                                v3 = 0
                                v113 := GenS9{F0: (42 > 0), F1: 0, F2: 687}
                                _ = v113
                            }
                            var v114 uint64 = 0
                            fmt.Println(fi102)
                            if ((false || (v3 < v3)) || false) {
                                var v115 bool = (!(!(true || ((false || ((51 <= 2) && (42 == 548))) || false))))
                                fmt.Println(v5, v115)
                                _ = v115
                            }
                            _ = v114
                        }
                        var v116 float64 = (1.5 * 0.0)
                        _ = v116
                    } else {
                        v117 := GenS7{F0: 32767}
                        v118 := GenS15{F0: -955, F1: 4294967295}
                        switch (-(fi103)) {
                        case -128:
                            switch ("foo bar" + (("foo bar" + "hello") + ("" + ("hello" + "\n\t")))) {
                            case "":
                                if false {
                                    var v119 uint = 1
                                    v5 -= fi102
                                    _ = v119
                                }
                            case "hello":
                                v5 = (fi103 << ((708) & 63))
                            }
                        case 0:
                            var v120 uint16 = 958
                            _ = v120
                        case -582:
                            if (42 == 1) {
                                v2 = 0.0
                            }
                        }
                        _ = v117
                        _ = v118
                    }
                    v121 := GenS17{F0: 4294967295}
                    _ = v108
                    _ = v121
                }
            }
            if false {
                var v122 float32 = (0.1 + float32(1e10))
                var v123 uint8 = 1
                v123 += v123
                _ = v122
                _ = v123
            }
            _ = fi103
        }
        _ = fi102
    }
    switch "a" {
    case "":
        fmt.Println(v3, v5, v3)
    case "hello":
        for fi124 := 0; fi124 < 4; fi124++ {
            v5 -= v5
            _ = fi124
        }
    }
    var v125 int16 = 32767
    switch -990 {
    case 4294967295:
        var v126 int64 = (v3 >> ((65535) & 63))
        _ = v126
    case 32767:
        for fi127 := 0; fi127 < 7; fi127++ {
            v128 := GenS8{F0: 154, F1: 586, F2: (^(v125))}
            fmt.Println(v125, v2, fi127)
            _ = fi127
            _ = v128
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v125:", v125)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v125
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
