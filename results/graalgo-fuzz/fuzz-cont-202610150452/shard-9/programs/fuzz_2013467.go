package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int32 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 int; F2 uint64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint32; F1 uint32 }
type GenS7 struct { F0 int16; F1 uint32; F2 int8 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int64; F1 string; F2 bool }
type GenS10 struct { F0 uint32; F1 uint64; F2 int }
type GenS11 struct { F0 uint64; F1 int }
type GenS12 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 int16; F1 string }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 uint16 }
func genFn0() int8 {
    var v5 uint32 = 187
    v5 += v5
    if true {
        for fi6 := 0; fi6 < 11; fi6++ {
            v5 = 741
            var v7 int8 = 23
            var v8 int64 = 9223372036854775807
            _ = fi6
            _ = v7
            _ = v8
        }
        v9 := []int{223, -37, 0}
        v10 := GenS1{F0: -1, F1: 354, F2: 290}
        _ = v10
        _ = v9
    } else {
        var v11 uint8 = 215
        if ((0 == 1) || true) {
            if false {
                for fi12 := 0; fi12 < 4; fi12++ {
                    v13 := []int{9223372036854775807, (fi12 / ((9223372036854775807) | 1)), -2147483648, fi12}
                    _ = fi12
                    _ = v13
                }
            }
            v11 = (^(v11))
        } else {
            v14 := GenS1{F0: 1, F1: -9223372036854775808, F2: 562}
            v15 := GenS1{F0: 1, F1: -1, F2: 887}
            _ = v14
            _ = v15
        }
        _ = v11
    }
    if (false || (false || (!false))) {
        if true {
            v5 -= v5
            for fi16 := 0; fi16 < 5; fi16++ {
                var v17 bool = true
                _ = fi16
                _ = v17
            }
            v5 = (v5 / ((v5) | 1))
        }
    } else {
        for fi18 := 0; fi18 < 1; fi18++ {
            v19 := GenS2{F0: 42, F1: 128}
            v20 := GenS3{F0: 0}
            fmt.Println(fi18, v5)
            _ = fi18
            _ = v19
            _ = v20
        }
        v5 += v5
    }
    fmt.Println(v5)
    _ = v5
    return -2
}
func genFn1() int {
    var v21 int32 = 0
    for fi22 := 0; fi22 < 20; fi22++ {
        var v23 float32 = 0.0
        v24 := GenS1{F0: (fi22 ^ fi22), F1: -489, F2: -1}
        if false {
            var v25 int = 128
            _ = v25
        } else {
            var v26 uint64 = 42
            fmt.Println(v26, v23)
            _ = v26
        }
        _ = fi22
        _ = v23
        _ = v24
    }
    _ = v21
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 1
    v3 := []int{32767, 32768, 32767}
    v4 := []uint32{}
    if (0 <= 556) {
        v2 -= v2
        var v27 float32 = 3.14159
        if len(v4) > 0 {
            v4[0] = 321
        }
        _ = v27
    }
    switch 128 {
    case 32767:
        var v28 int = 9223372036854775807
        _ = v28
    case 1:
        v29 := GenS4{F0: 9223372036854775807, F1: 32768, F2: 4294967295}
        _ = v29
    case 9223372036854775807:
        v1[""] = 2147483647
    default:
        for fi30 := 0; fi30 < 19; fi30++ {
            var v31 int64 = 275
            v2 = (v2 / ((v2) | 1))
            var v32 uint8 = 127
            _ = fi30
            _ = v31
            _ = v32
        }
    }
    for fi33 := 0; fi33 < 11; fi33++ {
        v2 -= 65535
        _ = fi33
    }
    for fi34 := 0; fi34 < 20; fi34++ {
        if true {
            if (242 < 85) {
                var v35 int = genFn1()
                v35 = v35
                _ = v35
            } else {
                if (2147483647 < 65535) {
                    if true {
                        if true {
                            var v36 string = ("a" + "")
                            _ = v36
                        }
                        v37 := GenS5{F0: 1}
                        for fi38 := 0; fi38 < 3; fi38++ {
                            var v39 int16 = -448
                            _ = fi38
                            _ = v39
                        }
                        _ = v37
                    } else {
                        var v40 uint16 = 2
                        _ = v40
                    }
                    switch "a" {
                    case "":
                        for fi41 := 0; fi41 < 20; fi41++ {
                            _ = len(v1)
                            var v42 float32 = (0.0 + (3.14159 + -2.25))
                            _ = fi41
                            _ = v42
                        }
                    case "a":
                        fmt.Println(v2)
                    case "foo bar":
                        v43 := v1[("foo bar" + ((("" + ("foo bar" + (("foo bar" + "\n\t") + ("hello" + "\n\t")))) + "hello") + "a"))]
                        _ = v43
                    }
                    var v44 bool = (false || ((((v2 == v2) && (!(false && (!(3.14159 <= 1e10))))) || (120 >= 482)) || ((67 <= 1) || (2147483648 <= 4294967295))))
                    _ = v44
                } else {
                    var v45 uint = 1
                    v45 = v45
                    _ = v45
                }
                fmt.Println(v2, v2)
            }
        } else {
            if false {
                v46 := GenS6{F0: 2147483648, F1: 2}
                _ = v46
            } else {
                if (540 <= 32767) {
                    switch (^(fi34)) {
                    case 1:
                        var v47 uint8 = 255
                        _ = v47
                    }
                }
            }
        }
        v2 = v2
        _ = fi34
    }
    if true {
        if true {
            var v48 int32 = -727
            for fi49 := 0; fi49 < 3; fi49++ {
                v50 := GenS5{F0: 0}
                var v51 uint16 = 32767
                if false {
                    if true {
                        switch fi49 {
                        case 4294967295:
                            var v52 float64 = -2.25
                            _ = v52
                        case -1:
                            if true {
                                if false {
                                    v1["\n\t"] = 356
                                }
                                if (811 <= 0) {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                var v53 uint = (v2 / ((42) | 1))
                                _ = v53
                            }
                        case -2147483648:
                            var v54 int = genFn1()
                            _ = v54
                        }
                        for fi55 := 0; fi55 < 14; fi55++ {
                            _ = len(v4)
                            _ = fi55
                        }
                    } else {
                        fmt.Println(v48)
                        v2 += (v2 - v2)
                        if false {
                            var v56 float64 = (1e10 - (3.14159 - -2.25))
                            v57 := GenS7{F0: 42, F1: 4294967295, F2: 72}
                            switch "foo bar" {
                            case "hello":
                                fmt.Println(v51)
                            }
                            _ = v56
                            _ = v57
                        }
                    }
                    fmt.Println(v51, v2)
                } else {
                    for fi58 := 0; fi58 < 14; fi58++ {
                        var v59 float64 = (-2.25 / ((1e10 * -2.25) / 1.0))
                        if true {
                            if (1 < 1000) {
                                v51 = 238
                                v2 = v2
                                if (128 <= 2) {
                                    v48 = 0
                                    v59 = v59
                                    switch ("" + "foo bar") {
                                    case "a":
                                        fmt.Println(v59, fi58)
                                    case "foo bar":
                                        fmt.Println(v59, v59, v51)
                                    case "hello":
                                        v60 := GenS8{F0: 0}
                                        _ = len(v1)
                                        delete(v1, ((("a" + ("foo bar" + "")) + ((("hello" + "hello") + "a") + (("\n\t" + "foo bar") + ""))) + ""))
                                        _ = v60
                                    default:
                                        switch (fi58 * (fi49 + fi49)) {
                                        case -2147483648:
                                            v61 := GenS9{F0: 184, F1: ("" + "\n\t"), F2: false}
                                            _ = v61
                                        }
                                    }
                                } else {
                                    v62 := GenS5{F0: 255}
                                    _ = v62
                                }
                            }
                            switch ("" + ("\n\t" + (("a" + "foo bar") + "hello"))) {
                            case "a":
                                switch fi58 {
                                case -774:
                                    var v63 float32 = 3.14159
                                    var v64 uint64 = 1
                                    v2 += 42
                                    _ = v63
                                    _ = v64
                                case -9223372036854775808:
                                    var v65 string = (("hello" + "hello") + "")
                                    _ = v65
                                default:
                                    v51 = v51
                                }
                            }
                            switch "hello" {
                            case "foo bar":
                                if (42 == -32768) {
                                    switch int(v48) {
                                    case 32768:
                                        var v66 int16 = 32767
                                        _ = v66
                                    default:
                                        var v67 string = ("hello" + ("hello" + "hello"))
                                        _ = v67
                                    }
                                }
                            case "hello":
                                v68 := GenS5{F0: 42}
                                var v69 int8 = -10
                                if true {
                                    v1[(("" + ("\n\t" + "\n\t")) + "hello")] = 2147483648
                                    var v70 int8 = genFn0()
                                    _ = v70
                                }
                                v71 := GenS1{F0: (fi58 | (^(fi49))), F1: 184, F2: v48}
                                _ = v68
                                _ = v69
                                _ = v71
                            default:
                                v72 := GenS10{F0: 31, F1: 0, F2: int(v51)}
                                _ = v72
                            }
                        }
                        _ = fi58
                        _ = v59
                    }
                    for fi73 := 0; fi73 < 11; fi73++ {
                        v2 -= v2
                        switch int(v2) {
                        case 32767:
                            if ((-879 > (^(fi73))) && false) {
                                v74 := GenS2{F0: 2147483647, F1: 127}
                                var v75 int8 = genFn0()
                                _ = v74
                                _ = v75
                            } else {
                                var v76 uint16 = (v51 % ((v51) | 1))
                                _ = len(v4)
                                var v77 uint64 = 2
                                _ = v76
                                _ = v77
                            }
                        case 42:
                            v2 = v2
                        }
                        _ = fi73
                    }
                }
                _ = fi49
                _ = v50
                _ = v51
            }
            var v78 float64 = -2.25
            _ = v48
            _ = v78
        }
        v79 := GenS10{F0: 2147483648, F1: 42, F2: 1}
        _ = v79
    } else {
        if true {
            v2 = v2
            v2 += (^(v2))
            if ((((("foo bar" + "") + (("\n\t" + "a") + ("a" + ""))) < "foo bar") && (false || false)) || true) {
                v2 = v2
                var v80 uint32 = 764
                v1["\n\t"] = (v80 * v80)
                _ = v80
            } else {
                v81 := GenS11{F0: 1, F1: -2147483648}
                if len(v3) > 0 {
                    v3[0] = -2147483648
                }
                _ = v81
            }
        }
        if false {
            v82 := GenS10{F0: 2147483648, F1: 42, F2: 2147483647}
            _ = v82
        }
        v2 = v2
    }
    v83 := GenS5{F0: 0}
    var v84 uint = (v2 % (((^(v2))) | 1))
    for fi85 := 0; fi85 < 4; fi85++ {
        var v86 int8 = 27
        _ = fi85
        _ = v86
    }
    var v87 int = 571
    switch "" {
    case "\n\t":
        var v88 string = "a"
        _ = v88
    case "hello":
        var v89 uint32 = 353
        _ = v89
    case "foo bar":
        if false {
            switch "a" {
            case "\n\t":
                var v90 bool = (1 == 1)
                _ = v90
            case "foo bar":
                var v91 uint16 = 42
                _ = v91
            default:
                for fi92 := 0; fi92 < 1; fi92++ {
                    fmt.Println(fi92, v84)
                    for fi93 := 0; fi93 < 8; fi93++ {
                        if (1.5 == -2.25) {
                            v2 = v2
                            _ = len(v1)
                            v94 := GenS11{F0: 4294967295, F1: (v87 | (fi93 + (v87 * (v87 % (((fi93 - (fi93 & fi93))) | 1)))))}
                            _ = v94
                        }
                        _ = fi93
                    }
                    var v95 bool = true
                    _ = fi92
                    _ = v95
                }
            }
            v2 -= 65535
        } else {
            if ((v2 << (((^(v84))) & 63)) != v2) {
                for fi96 := 0; fi96 < 13; fi96++ {
                    if (v84 > 0) {
                        v84 = (v84 >> ((v84) & 63))
                    } else {
                        var v97 float32 = 0.0
                        v1[""] = 4294967295
                        _ = v97
                    }
                    if (-262 > int16(v84)) {
                        var v98 int = v87
                        _ = v98
                    }
                    v99 := GenS12{F0: 65535, F1: -1, F2: 0}
                    _ = fi96
                    _ = v99
                }
            } else {
                var v100 int = (v87 | -810)
                var v101 int32 = 2147483647
                _ = v100
                _ = v101
            }
            if true {
                if true {
                    var v102 float32 = float32(0.0)
                    _ = v102
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        v2 = 2
                    default:
                        if false {
                            for fi103 := 0; fi103 < 18; fi103++ {
                                v87 -= fi103
                                if false {
                                    for fi104 := 0; fi104 < 2; fi104++ {
                                        fmt.Println(fi104, v2)
                                        var v105 int64 = 0
                                        if false {
                                            var v106 int32 = 42
                                            _ = v106
                                        }
                                        _ = fi104
                                        _ = v105
                                    }
                                    v87 = fi103
                                } else {
                                    for fi107 := 0; fi107 < 17; fi107++ {
                                        var v108 uint16 = 32767
                                        if false {
                                            var v109 float32 = (3.14159 * (1.5 * 1.5))
                                            _ = v109
                                        } else {
                                            var v110 int32 = -1
                                            _ = len(v1)
                                            var v111 uint16 = 53
                                            _ = v110
                                            _ = v111
                                        }
                                        switch (-(fi103)) {
                                        case -2147483648:
                                            if false {
                                                if true {
                                                    _ = len(v1)
                                                    v2 -= 833
                                                    if len(v4) > 0 {
                                                        v4[0] = 3
                                                    }
                                                }
                                                if false {
                                                    var v112 string = "\n\t"
                                                    var v113 bool = (2 < 0)
                                                    _ = v112
                                                    _ = v113
                                                }
                                            } else {
                                                var v114 int16 = 1000
                                                switch fi107 {
                                                case 544:
                                                    if (18446744073709551615 >= 18446744073709551615) {
                                                        v115 := GenS4{F0: 0, F1: genFn1(), F2: 18446744073709551615}
                                                        var v116 float32 = ((3.14159 - -2.25) + 3.14159)
                                                        v117 := GenS13{F0: 127}
                                                        switch "" {
                                                        case "foo bar":
                                                            v116 = 1e10
                                                        default:
                                                            v116 += (0.1 + v116)
                                                        }
                                                        _ = v115
                                                        _ = v116
                                                        _ = v117
                                                    } else {
                                                        if true {
                                                            v118 := GenS14{F0: -74}
                                                            _ = v118
                                                        }
                                                        var v119 int = -122
                                                        if true {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            var v120 uint = v84
                                                            v84 = 255
                                                            _ = v120
                                                        }
                                                        _ = v119
                                                    }
                                                default:
                                                    var v121 bool = true
                                                    _ = v121
                                                }
                                                var v122 bool = true
                                                _ = v114
                                                _ = v122
                                            }
                                        case 42:
                                            v123 := GenS12{F0: (v108 * v108), F1: 1, F2: v108}
                                            _ = v123
                                        default:
                                            if (2147483647 >= 0) {
                                                if true {
                                                    switch "\n\t" {
                                                    case "hello":
                                                        fmt.Println(fi103, v108)
                                                    default:
                                                        var v124 uint16 = 274
                                                        _ = v124
                                                    }
                                                }
                                                var v125 uint64 = 18446744073709551615
                                                switch genFn1() {
                                                case 416:
                                                    v87 = 9223372036854775807
                                                case 128:
                                                    _ = len(v4)
                                                case 613:
                                                    var v126 string = ("" + "foo bar")
                                                    _ = v126
                                                }
                                                fmt.Println(v84)
                                                _ = v125
                                            } else {
                                                if false {
                                                    delete(v1, (("foo bar" + ("a" + "foo bar")) + "a"))
                                                } else {
                                                    delete(v1, "")
                                                }
                                                switch "" {
                                                case "hello":
                                                    var v127 int16 = 945
                                                    _ = v127
                                                case "foo bar":
                                                    fmt.Println(v2, fi103)
                                                default:
                                                    v128 := GenS3{F0: 42}
                                                    _ = v128
                                                }
                                            }
                                        }
                                        _ = fi107
                                        _ = v108
                                    }
                                }
                                _ = fi103
                            }
                        } else {
                            v84 = 18446744073709551615
                        }
                    }
                    var v129 float32 = (1.5 / 1e10)
                    for fi130 := 0; fi130 < 19; fi130++ {
                        switch (fi130 >> (((v84 ^ v84)) & 63)) {
                        case -638:
                            v129 += (v129 / v129)
                        case 38:
                            if (1 >= -961) {
                                var v131 uint32 = 65535
                                var v132 int16 = 1000
                                for fi133 := 0; fi133 < 8; fi133++ {
                                    var v134 bool = (true && false)
                                    _ = fi133
                                    _ = v134
                                }
                                _ = v131
                                _ = v132
                            }
                        default:
                            v135 := GenS10{F0: 1, F1: 0, F2: (fi130 + 127)}
                            _ = v135
                        }
                        v2 += v84
                        _ = fi130
                    }
                    _ = v129
                }
                switch "a" {
                case "":
                    var v136 uint16 = 387
                    _ = v136
                case "\n\t":
                    for fi137 := 0; fi137 < 7; fi137++ {
                        var v138 int16 = 1000
                        v139 := GenS14{F0: -2147483648}
                        var v140 int16 = v138
                        _ = fi137
                        _ = v138
                        _ = v139
                        _ = v140
                    }
                default:
                    switch ("\n\t" + "foo bar") {
                    case "foo bar":
                        v84 -= v2
                    case "a":
                        for fi141 := 0; fi141 < 6; fi141++ {
                            v142 := GenS3{F0: 42}
                            _ = fi141
                            _ = v142
                        }
                    default:
                        if (("hello" > "") || ("a" > ("" + "a"))) {
                            for fi143 := 0; fi143 < 9; fi143++ {
                                v87 = (fi143 & -838)
                                var v144 float32 = -2.25
                                _ = fi143
                                _ = v144
                            }
                            v145 := GenS15{F0: -32768, F1: "\n\t"}
                            var v146 string = "a"
                            _ = v145
                            _ = v146
                        }
                    }
                }
                fmt.Println(v2, v2)
            }
        }
    default:
        v147 := GenS4{F0: 42, F1: genFn1(), F2: 762}
        _ = v147
    }
    v84 += (^(v84))
    v148 := GenS2{F0: 1, F1: 2}
    if true {
        var v149 bool = (590 >= 42)
        v87 += v87
        _ = v149
    } else {
        switch ("hello" + "") {
        case "foo bar":
            if len(v3) > 0 {
                v3[0] = (v87 & (v87 % (((v87 - v87)) | 1)))
            }
        case "":
            for fi150 := 0; fi150 < 8; fi150++ {
                _ = len(v4)
                for fi151 := 0; fi151 < 8; fi151++ {
                    var v152 uint8 = 2
                    v84 = 255
                    if (0.1 != float32(1.5)) {
                        switch "foo bar" {
                        case "\n\t":
                            if true {
                                var v153 uint32 = 65535
                                _ = v153
                            }
                        case "foo bar":
                            switch "hello" {
                            case "a":
                                if len(v4) > 0 {
                                    v4[0] = 2
                                }
                                var v154 float32 = 3.14159
                                var v155 uint32 = 719
                                _ = v154
                                _ = v155
                            case "foo bar":
                                if false {
                                    v156 := GenS10{F0: 1, F1: 0, F2: (^(fi150))}
                                    _ = v156
                                } else {
                                    fmt.Println(v2, v87)
                                    var v157 bool = (0 >= 42)
                                    _ = v157
                                }
                            case "\n\t":
                                var v158 uint32 = 765
                                _ = v158
                            default:
                                var v159 bool = true
                                switch (("a" + "\n\t") + ("" + ("" + ("" + "")))) {
                                case "a":
                                    if v159 {
                                        var v160 uint = (v2 % ((42) | 1))
                                        delete(v1, "\n\t")
                                        _ = v160
                                    }
                                case "foo bar":
                                    var v161 uint32 = 2147483648
                                    _ = v161
                                case "":
                                    fmt.Println(v87, v87)
                                }
                                _ = v159
                            }
                        case "a":
                            fmt.Println(v84, v2, v87)
                        }
                        if true {
                            switch "foo bar" {
                            case "":
                                v152 -= v152
                                v162 := GenS2{F0: 0, F1: (^(v152))}
                                _ = v162
                            case "\n\t":
                                v87 -= fi151
                            }
                            v84 += 255
                        }
                        var v163 int64 = 1
                        _ = v163
                    }
                    if true {
                        if len(v4) > 0 {
                            v4[0] = 2147483648
                        }
                        v2 -= v2
                    } else {
                        v164 := v1["foo bar"]
                        _ = v164
                    }
                    switch fi150 {
                    case -2147483648:
                        fmt.Println(v152, v152, fi150)
                    case 4294967295:
                        v84 += 2
                    case 32768:
                        var v165 int32 = 1
                        _ = v165
                    }
                    v4 = append(v4, 628)
                    var v166 uint32 = 2147483647
                    if (!false) {
                        var v167 float64 = 1e10
                        v2 = uint(v166)
                        _ = v167
                    }
                    v152 += v152
                    switch 128 {
                    case 128:
                        v2 -= 687
                    }
                    _ = fi151
                    _ = v152
                    _ = v166
                }
                for fi168 := 0; fi168 < 15; fi168++ {
                    v169 := GenS16{F0: 9223372036854775807}
                    v2 -= (v2 & v2)
                    var v170 uint = (v84 >> ((10) & 63))
                    _ = fi168
                    _ = v169
                    _ = v170
                }
                _ = fi150
            }
        case "hello":
            v2 = 515
        default:
            v171 := GenS17{F0: 0}
            _ = v171
        }
        var v172 uint32 = 216
        _ = v172
    }
    var v173 uint16 = 65535
    var v174 int16 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v173:", v173)
    fmt.Println("v174:", v174)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v84:", v84)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v148
    _ = v173
    _ = v174
    _ = v2
    _ = v3
    _ = v4
    _ = v83
    _ = v84
    _ = v87
}
