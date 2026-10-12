package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 int32; F2 int64 }
type GenS5 struct { F0 int32; F1 string; F2 int16 }
type GenS6 struct { F0 int64; F1 int16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint32; F1 int8; F2 int }
type GenS9 struct { F0 int32; F1 uint }
type GenS10 struct { F0 uint; F1 uint64; F2 uint }
type GenS11 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS12 struct { F0 uint32; F1 int8; F2 int }
type GenS13 struct { F0 int32; F1 int; F2 int }
type GenS14 struct { F0 uint16; F1 uint8 }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 uint8; F1 int32 }
func genFn0(p0 int8, p1 uint8, p2 int) uint32 {
    if true {
        var v7 int64 = 0
        for fi8 := 0; fi8 < 12; fi8++ {
            v9 := []int{(p2 << ((1) & 63))}
            v10 := GenS1{F0: 4294967295}
            v11 := GenS1{F0: 42}
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        _ = v7
    }
    var v12 int16 = 1000
    v13 := GenS1{F0: 2}
    var v14 float32 = (1e10 / (1.5 - (3.14159 + 0.1)))
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v14
    return 506
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{}
    var v3 int16 = 1
    var v4 int8 = 100
    var v5 int64 = -9223372036854775808
    var v6 int8 = v4
    v15 := GenS2{F0: 237}
    for fi16 := 0; fi16 < 20; fi16++ {
        fmt.Println(v5)
        switch 127 {
        case 9223372036854775807:
            var v17 float64 = 1e10
            _ = v17
        case -270:
            if true {
                switch -2147483648 {
                case 32768:
                    v18 := GenS2{F0: 4294967295}
                    _ = v18
                default:
                    v5 = v5
                }
                v3 = 1
                v3 += (^(v3))
            }
        case 128:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = fi16
    }
    v19 := GenS3{F0: 424}
    if (!(1.5 >= -2.25)) {
        var v20 uint8 = 200
        v5 = -9223372036854775808
        if (!(-496 <= 94)) {
            v21 := GenS4{F0: 4294967295, F1: 0, F2: v5}
            _ = v21
        } else {
            for fi22 := 0; fi22 < 5; fi22++ {
                _ = len(v2)
                v23 := GenS1{F0: 255}
                v2 = append(v2, ("a" + ("\n\t" + ("a" + "\n\t"))))
                _ = fi22
                _ = v23
            }
            if (-1 != -914) {
                if false {
                    switch ("\n\t" + "\n\t") {
                    case "hello":
                        v24 := v1[2147483647]
                        _ = v24
                    case "foo bar":
                        switch -2147483648 {
                        case 9223372036854775807:
                            fmt.Println(v5, v5)
                        case -2147483648:
                            if (!(0 == 32767)) {
                                v3 = (v3 / ((int16(v6)) | 1))
                                switch "" {
                                case "\n\t":
                                    switch 32768 {
                                    case 127:
                                        var v25 string = ("a" + "a")
                                        _ = v25
                                    case -323:
                                        var v26 uint = 940
                                        _ = v26
                                    case 128:
                                        if false {
                                            v20 = 127
                                            v3 += (v3 ^ 1000)
                                        }
                                    default:
                                        for fi27 := 0; fi27 < 17; fi27++ {
                                            for fi28 := 0; fi28 < 19; fi28++ {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    delete(v1, genFn0(1, v20, 42))
                                                case "a":
                                                    if (v5 != (v5 / (((-(v5))) | 1))) {
                                                        var v29 uint16 = 0
                                                        _ = v29
                                                    } else {
                                                        fmt.Println(v4, v5)
                                                        fmt.Println(fi27)
                                                        if (0.1 == ((float32(-2.25) + float32(0.1)) + 1e10)) {
                                                            fmt.Println(v6, v5)
                                                            if false {
                                                                if true {
                                                                    var v30 int = (fi28 >> ((42) & 63))
                                                                    _ = v30
                                                                } else {
                                                                    _ = len(v2)
                                                                    v31 := GenS5{F0: 42, F1: "a", F2: (v3 & v3)}
                                                                    v32 := GenS6{F0: v5, F1: v3}
                                                                    _ = v31
                                                                    _ = v32
                                                                }
                                                                var v33 uint8 = v20
                                                                switch ("" + "a") {
                                                                case "a":
                                                                    fmt.Println(fi27)
                                                                case "hello":
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        if (true || (42 <= 42)) {
                                                                            v1[2147483647] = ("foo bar" + "hello")
                                                                            v34 := GenS4{F0: 0, F1: -2147483648, F2: (v5 & (v5 | (^(v5))))}
                                                                            var v35 bool = ((v6 * (v6 - (v6 + v4))) <= v6)
                                                                            _ = v34
                                                                            _ = v35
                                                                        } else {
                                                                            v33 = 2
                                                                            var v36 int32 = 2147483647
                                                                            _ = v36
                                                                        }
                                                                    case "\n\t":
                                                                        v4 = (v6 / (((v6 - (v6 * (v4 % ((v4) | 1))))) | 1))
                                                                        fmt.Println(v20, fi27, v4)
                                                                        v6 = 127
                                                                        v5 = v5
                                                                        var v37 uint8 = (v33 ^ v33)
                                                                        v38 := v1[4294967295]
                                                                        _ = v37
                                                                        _ = v38
                                                                    }
                                                                default:
                                                                    fmt.Println(v20, v5, fi27)
                                                                }
                                                                fmt.Println(v33)
                                                                _ = v33
                                                            } else {
                                                                fmt.Println(v6)
                                                            }
                                                            if true {
                                                                var v39 float32 = ((1.5 / 0.1) - 0.1)
                                                                if ((true || (676 == 255)) && false) {
                                                                    var v40 string = ""
                                                                    fmt.Println(v3)
                                                                    if false {
                                                                        v39 += v39
                                                                        v41 := GenS2{F0: 18446744073709551615}
                                                                        v5 = v5
                                                                        _ = v41
                                                                    }
                                                                    _ = v40
                                                                } else {
                                                                    fmt.Println(fi28)
                                                                    var v42 bool = false
                                                                    if len(v2) > 0 {
                                                                        v2[0] = "hello"
                                                                    }
                                                                    _ = v42
                                                                }
                                                                v43 := GenS2{F0: 828}
                                                                _ = v39
                                                                _ = v43
                                                            } else {
                                                                switch (fi28 + fi28) {
                                                                case 1:
                                                                    v44 := GenS3{F0: 831}
                                                                    _ = v44
                                                                case -37:
                                                                    v45 := GenS6{F0: 42, F1: 1000}
                                                                    _ = v45
                                                                case 927:
                                                                    v46 := GenS7{F0: fi28}
                                                                    _ = v46
                                                                default:
                                                                    v47 := GenS8{F0: 4294967295, F1: v4, F2: (fi28 >> ((4294967295) & 63))}
                                                                    _ = v47
                                                                }
                                                                var v48 int32 = -2147483648
                                                                fmt.Println(v48, v4, v3)
                                                                var v49 uint = 255
                                                                switch 4294967295 {
                                                                case 0:
                                                                    v20 -= v20
                                                                default:
                                                                    v4 = -1
                                                                }
                                                                if true {
                                                                    switch 127 {
                                                                    case -855:
                                                                        v50 := GenS4{F0: 779, F1: (v48 * (v48 & 0)), F2: int64(v6)}
                                                                        _ = v50
                                                                    case 128:
                                                                        v51 := GenS9{F0: v48, F1: v49}
                                                                        _ = v51
                                                                    }
                                                                    v3 -= v3
                                                                    v49 = (v49 / ((445) | 1))
                                                                } else {
                                                                    fmt.Println(fi27)
                                                                    var v52 int32 = 42
                                                                    switch 834 {
                                                                    case 0:
                                                                        fmt.Println(v5, v3)
                                                                    case 127:
                                                                        switch fi28 {
                                                                        case 978:
                                                                            var v53 int32 = v48
                                                                            _ = v53
                                                                        default:
                                                                            switch fi28 {
                                                                            case 0:
                                                                                switch (fi27 & (fi28 << ((42) & 63))) {
                                                                                case -53:
                                                                                    _ = len(v1)
                                                                                default:
                                                                                    var v54 int = (fi28 & fi27)
                                                                                    _ = v54
                                                                                }
                                                                            default:
                                                                                v55 := v1[65535]
                                                                                _ = v55
                                                                            }
                                                                        }
                                                                        v56 := GenS10{F0: (v49 & v49), F1: 2, F2: v49}
                                                                        _ = v56
                                                                    case 1:
                                                                        if (true || false) {
                                                                            var v57 float64 = 3.14159
                                                                            v6 += v6
                                                                            fmt.Println(v57, v3)
                                                                            _ = v57
                                                                        }
                                                                    }
                                                                    _ = v52
                                                                }
                                                                _ = v48
                                                                _ = v49
                                                            }
                                                            v58 := GenS11{F0: 1, F1: 32767, F2: v20}
                                                            var v59 int64 = (v5 | v5)
                                                            _ = v58
                                                            _ = v59
                                                        }
                                                    }
                                                default:
                                                    var v60 float32 = 3.14159
                                                    _ = v60
                                                }
                                                _ = fi28
                                            }
                                            v3 = -1
                                            fmt.Println(v3, v3, v4)
                                            _ = fi27
                                        }
                                    }
                                case "a":
                                    for fi61 := 0; fi61 < 10; fi61++ {
                                        switch ("hello" + "a") {
                                        case "a":
                                            fmt.Println(v5, fi61)
                                        case "foo bar":
                                            for fi62 := 0; fi62 < 19; fi62++ {
                                                _ = len(v2)
                                                _ = fi62
                                            }
                                        }
                                        switch (-(fi61)) {
                                        case 9223372036854775807:
                                            for fi63 := 0; fi63 < 7; fi63++ {
                                                v64 := GenS12{F0: 1, F1: (v6 | v6), F2: fi63}
                                                _ = fi63
                                                _ = v64
                                            }
                                        case 0:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        case -873:
                                            v65 := GenS13{F0: -1, F1: (fi61 / (((fi61 * 127)) | 1)), F2: (fi61 >> ((0) & 63))}
                                            _ = v65
                                        }
                                        if ((1.5 > 1.5) && false) {
                                            switch fi61 {
                                            case 4294967295:
                                                switch (fi61 * (fi61 | fi61)) {
                                                case 2147483647:
                                                    v3 += v3
                                                case 42:
                                                    var v66 int = fi61
                                                    _ = v66
                                                case 1:
                                                    if true {
                                                        if (-1 == -2147483648) {
                                                            switch ("a" + (("foo bar" + "hello") + ("a" + ""))) {
                                                            case "foo bar":
                                                                v67 := GenS8{F0: 542, F1: (^(v6)), F2: fi61}
                                                                _ = v67
                                                            }
                                                        } else {
                                                            v68 := GenS14{F0: 42, F1: v20}
                                                            var v69 int16 = 234
                                                            _ = v68
                                                            _ = v69
                                                        }
                                                    }
                                                default:
                                                    _ = len(v1)
                                                }
                                            }
                                            v70 := GenS3{F0: 246}
                                            _ = v70
                                        }
                                        _ = fi61
                                    }
                                case "hello":
                                    for fi71 := 0; fi71 < 2; fi71++ {
                                        var v72 float32 = 3.14159
                                        var v73 int64 = (v5 + v5)
                                        v4 = (v6 + v4)
                                        _ = fi71
                                        _ = v72
                                        _ = v73
                                    }
                                default:
                                    var v74 int = 128
                                    _ = v74
                                }
                            }
                        default:
                            for fi75 := 0; fi75 < 20; fi75++ {
                                if true {
                                    for fi76 := 0; fi76 < 1; fi76++ {
                                        _ = len(v2)
                                        switch fi75 {
                                        case -2147483648:
                                            fmt.Println(v5)
                                        case 32768:
                                            fmt.Println(v4, fi75, v6)
                                        case 127:
                                            var v77 uint16 = 255
                                            switch ("hello" + ("a" + "\n\t")) {
                                            case "a":
                                                v77 -= (^(v77))
                                            default:
                                                if (1 <= 231) {
                                                    fmt.Println(v6, v6, v3)
                                                } else {
                                                    v78 := GenS13{F0: 0, F1: fi76, F2: -1}
                                                    _ = v78
                                                }
                                                var v79 int = (fi75 ^ (-(fi76)))
                                                _ = v79
                                            }
                                            _ = v77
                                        default:
                                            var v80 int32 = 0
                                            _ = v80
                                        }
                                        var v81 uint64 = 819
                                        _ = fi76
                                        _ = v81
                                    }
                                    v82 := GenS11{F0: 65535, F1: 42, F2: v20}
                                    _ = v82
                                } else {
                                    var v83 float32 = (3.14159 - 1e10)
                                    var v84 uint32 = 2
                                    v85 := GenS6{F0: (v5 / (((v5 & 1)) | 1)), F1: -649}
                                    _ = v83
                                    _ = v84
                                    _ = v85
                                }
                                _ = fi75
                            }
                        }
                    case "":
                        if (!false) {
                            var v86 int = 268
                            _ = v86
                        } else {
                            switch 752 {
                            case 32767:
                                v87 := GenS11{F0: 42, F1: 618, F2: (v20 + 127)}
                                _ = v87
                            case 4294967295:
                                if false {
                                    v88 := GenS14{F0: 51, F1: v20}
                                    v3 -= v3
                                    if true {
                                        fmt.Println(v5, v5)
                                    } else {
                                        v5 = 379
                                        var v89 uint8 = (v20 << ((4294967295) & 63))
                                        v4 = (v4 ^ (v6 * 36))
                                        _ = v89
                                    }
                                    _ = v88
                                }
                            }
                            v90 := GenS15{F0: genFn0(-1, v20, 42)}
                            _ = v90
                        }
                    }
                }
                fmt.Println(v20, v4, v3)
            } else {
                v91 := GenS15{F0: 0}
                v6 = v4
                if (true || (-20 > (-(v4)))) {
                    v5 = v5
                } else {
                    v5 += v5
                    if (true || false) {
                        for fi92 := 0; fi92 < 19; fi92++ {
                            for fi93 := 0; fi93 < 19; fi93++ {
                                if true {
                                    _ = len(v2)
                                    if (v6 == v4) {
                                        if false {
                                            switch fi93 {
                                            case 32767:
                                                v94 := GenS2{F0: 18446744073709551615}
                                                _ = v94
                                            case 32768:
                                                v95 := GenS6{F0: 42, F1: 42}
                                                _ = v95
                                            case -351:
                                                var v96 float32 = -2.25
                                                _ = v96
                                            }
                                            var v97 int32 = 42
                                            v98 := v1[65535]
                                            delete(v1, 566)
                                            _ = v97
                                            _ = v98
                                        }
                                    } else {
                                        v4 += v6
                                        var v99 float64 = 0.1
                                        _ = v99
                                    }
                                }
                                _ = fi93
                            }
                            for fi100 := 0; fi100 < 12; fi100++ {
                                v101 := GenS8{F0: 2147483647, F1: v6, F2: fi92}
                                _ = fi100
                                _ = v101
                            }
                            _ = fi92
                        }
                        if ((4294967295 != 0) && (!true)) {
                            for fi102 := 0; fi102 < 4; fi102++ {
                                if false {
                                    for fi103 := 0; fi103 < 12; fi103++ {
                                        switch "\n\t" {
                                        case "foo bar":
                                            v4 = v6
                                        case "":
                                            v20 = v20
                                        default:
                                            v104 := GenS6{F0: (v5 << ((2) & 63)), F1: 1}
                                            _ = v104
                                        }
                                        _ = fi103
                                    }
                                }
                                _ = fi102
                            }
                            var v105 int64 = 42
                            _ = v105
                        } else {
                            var v106 int = 9223372036854775807
                            if len(v2) > 0 {
                                v2[0] = (((("\n\t" + "hello") + "a") + ((("a" + "hello") + ("a" + "\n\t")) + "\n\t")) + "\n\t")
                            }
                            var v107 string = "\n\t"
                            _ = v106
                            _ = v107
                        }
                    } else {
                        for fi108 := 0; fi108 < 9; fi108++ {
                            v20 = v20
                            switch "a" {
                            case "\n\t":
                                var v109 int32 = int32(v20)
                                _ = v109
                            case "":
                                if false {
                                    if (true || false) {
                                        var v110 int = fi108
                                        _ = len(v1)
                                        _ = v110
                                    }
                                    if (false && false) {
                                        var v111 int16 = -138
                                        if false {
                                            if (477 == 2) {
                                                switch 9223372036854775807 {
                                                case 383:
                                                    for fi112 := 0; fi112 < 6; fi112++ {
                                                        v5 -= v5
                                                        fmt.Println(fi112)
                                                        _ = fi112
                                                    }
                                                case 42:
                                                    v5 += (v5 - 42)
                                                case -2147483648:
                                                    var v113 uint16 = 42
                                                    _ = v113
                                                }
                                                var v114 int16 = v111
                                                var v115 uint8 = 127
                                                _ = v114
                                                _ = v115
                                            } else {
                                                fmt.Println(v3, fi108)
                                            }
                                        } else {
                                            v116 := GenS11{F0: 983, F1: 1, F2: v20}
                                            _ = v116
                                        }
                                        _ = v111
                                    }
                                    var v117 string = ""
                                    _ = v117
                                } else {
                                    var v118 int32 = 2147483647
                                    _ = v118
                                }
                            default:
                                if (1 < (v5 + (v5 * (v5 % ((v5) | 1))))) {
                                    var v119 int16 = v3
                                    var v120 string = "hello"
                                    v121 := GenS15{F0: 159}
                                    _ = v119
                                    _ = v120
                                    _ = v121
                                } else {
                                    if len(v2) > 0 {
                                        v2[0] = ""
                                    }
                                    v20 -= 141
                                    var v122 float32 = 3.14159
                                    _ = v122
                                }
                            }
                            _ = fi108
                        }
                    }
                }
                _ = v91
            }
        }
        _ = v20
    } else {
        v2 = append(v2, "")
        switch "hello" {
        case "hello":
            var v123 int = 908
            _ = v123
        default:
            for fi124 := 0; fi124 < 17; fi124++ {
                if false {
                    if true {
                        for fi125 := 0; fi125 < 1; fi125++ {
                            var v126 float64 = 1e10
                            _ = fi125
                            _ = v126
                        }
                    }
                    v3 = (v3 << ((0) & 63))
                    for fi127 := 0; fi127 < 20; fi127++ {
                        v128 := GenS16{F0: 969}
                        _ = fi127
                        _ = v128
                    }
                } else {
                    _ = len(v2)
                    var v129 uint8 = 235
                    _ = v129
                }
                var v130 string = ""
                _ = fi124
                _ = v130
            }
        }
    }
    if true {
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
        switch "\n\t" {
        case "\n\t":
            for fi131 := 0; fi131 < 18; fi131++ {
                for fi132 := 0; fi132 < 4; fi132++ {
                    _ = len(v1)
                    var v133 float32 = (1e10 + (((0.1 - -2.25) / -2.25) * (0.1 / 1e10)))
                    _ = fi132
                    _ = v133
                }
                _ = fi131
            }
        case "hello":
            switch ("\n\t" + "foo bar") {
            case "hello":
                var v134 float64 = 0.0
                _ = v134
            }
        case "foo bar":
            switch 2147483647 {
            case 42:
                var v135 uint16 = 0
                _ = v135
            default:
                fmt.Println(v5)
            }
        }
    }
    for fi136 := 0; fi136 < 12; fi136++ {
        switch "hello" {
        case "\n\t":
            fmt.Println(v6)
        case "":
            v4 += 0
        case "foo bar":
            var v137 uint32 = 946
            _ = v137
        }
        v3 -= 32767
        _ = fi136
    }
    for fi138 := 0; fi138 < 17; fi138++ {
        if ((((("hello" + "hello") + "\n\t") + "a") >= "hello") || false) {
            var v139 int32 = int32(v6)
            _ = v139
        }
        _ = fi138
    }
    switch 2147483647 {
    case 32768:
        v140 := GenS17{F0: 0, F1: 1}
        _ = v140
    case 32767:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case -116:
        v1[2] = ""
    default:
        var v141 int8 = v4
        _ = v141
    }
    v5 = 1
    var v142 int16 = (v3 + v3)
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v142:", v142)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v142
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
