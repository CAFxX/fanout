package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int; F1 uint64; F2 int8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int64; F1 uint32 }
type GenS9 struct { F0 int8; F1 int8; F2 int }
type GenS10 struct { F0 int32; F1 uint }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 bool; F1 uint16; F2 uint64 }
type GenS13 struct { F0 uint; F1 int16; F2 string }
type GenS14 struct { F0 string; F1 uint8; F2 int8 }
type GenS15 struct { F0 int32; F1 uint16 }

func main() {
    var v1 uint32 = 1
    var v2 float32 = (3.14159 - (1.5 / 0.1))
    v3 := make(map[uint32]string)
    v4 := make(map[int]int)
    var v5 int32 = -1
    var v6 uint64 = 4294967295
    v7 := GenS1{F0: (false || (481 < 805))}
    var v8 int8 = 107
    var v9 uint32 = v1
    var v10 bool = false
    if false {
        _ = len(v4)
        fmt.Println(v8)
    }
    v6 += 0
    v9 = v9
    fmt.Println(v9, v5)
    var v11 uint = 838
    v12 := GenS1{F0: false}
    var v13 int32 = -2147483648
    if (float64((((v2 - 0.1) - v2) - v2)) != (float64(1.5) - (0.0 + float64(float32((0.1 - 1.5)))))) {
        switch "a" {
        case "a":
            var v14 uint8 = 171
            _ = v14
        case "\n\t":
            _ = len(v4)
        }
        var v15 uint = (v11 & 2)
        if v10 {
            var v16 float64 = 1.5
            var v17 uint8 = 76
            v16 = v16
            _ = v16
            _ = v17
        }
        _ = v15
    }
    if v10 {
        switch ("hello" + ("foo bar" + "hello")) {
        case "\n\t":
            if v10 {
                for fi18 := 0; fi18 < 12; fi18++ {
                    if true {
                        for fi19 := 0; fi19 < 14; fi19++ {
                            v10 = (653 >= fi19)
                            v6 = v6
                            if true {
                                v20 := []int{fi19, 42, 834, fi19}
                                v5 = (v13 % (((v5 + (v5 + (v13 | 0)))) | 1))
                                _ = v20
                            }
                            if true {
                                v21 := GenS1{F0: v10}
                                if ((v1 << (((v11 << (((v11 ^ (v11 * 237))) & 63))) & 63)) != v9) {
                                    var v22 int16 = -32768
                                    v23 := GenS2{F0: fi18, F1: 1}
                                    var v24 float32 = 0.1
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                } else {
                                    var v25 string = "\n\t"
                                    _ = v25
                                }
                                _ = v21
                            } else {
                                v26 := GenS2{F0: (fi18 / ((fi19) | 1)), F1: 241}
                                var v27 uint = v11
                                var v28 uint8 = 255
                                var v29 uint32 = v1
                                _ = v26
                                _ = v27
                                _ = v28
                                _ = v29
                            }
                            var v30 uint64 = v6
                            var v31 int = (-(fi19))
                            _ = fi19
                            _ = v30
                            _ = v31
                        }
                        _ = len(v3)
                        if v10 {
                            var v32 float32 = v2
                            fmt.Println(v1, v2, v32)
                            switch "a" {
                            case "a":
                                var v33 uint32 = v9
                                _ = v33
                            case "hello":
                                v8 -= (v8 ^ (v8 * 127))
                            case "":
                                v11 = v11
                            }
                            _ = v32
                        }
                    } else {
                        v34 := []int{740, fi18}
                        v34 = append(v34, (fi18 * fi18))
                        _ = v34
                    }
                    v35 := v4[1]
                    _ = fi18
                    _ = v35
                }
                fmt.Println(v10, v8, v10)
            } else {
                var v36 int16 = 1000
                v1 = v9
                _ = v36
            }
        case "":
            var v37 int32 = (v13 / (((v13 + (^(v5)))) | 1))
            _ = v37
        case "hello":
            var v38 uint32 = 837
            _ = v38
        default:
            for fi39 := 0; fi39 < 5; fi39++ {
                _ = len(v3)
                v40 := GenS3{F0: (v6 ^ v6), F1: int(v13)}
                for fi41 := 0; fi41 < 20; fi41++ {
                    if (int64(fi41) >= 1) {
                        v4[32768] = (^(fi41))
                        if true {
                            switch ("a" + "\n\t") {
                            case "\n\t":
                                var v42 uint8 = 128
                                _ = v42
                            }
                            v43 := GenS1{F0: (fi39 >= fi41)}
                            switch fi41 {
                            case -9223372036854775808:
                                v9 = v9
                            default:
                                fmt.Println(v11, v10)
                            }
                            _ = v43
                        } else {
                            switch (("a" + "a") + "a") {
                            case "\n\t":
                                var v44 int8 = (v8 * v8)
                                _ = v44
                            case "hello":
                                v6 = v6
                            default:
                                var v45 float32 = (v2 * 1e10)
                                _ = v45
                            }
                        }
                    } else {
                        switch ("\n\t" + "a") {
                        case "hello":
                            v46 := v3[uint32(v6)]
                            _ = v46
                        }
                        delete(v3, (v9 / (((v9 % ((1) | 1))) | 1)))
                    }
                    var v47 string = "\n\t"
                    var v48 uint = uint(v1)
                    _ = fi41
                    _ = v47
                    _ = v48
                }
                _ = fi39
                _ = v40
            }
        }
        var v49 uint32 = (v1 & 1)
        _ = v49
    } else {
        if ((!v10) || (v1 > v1)) {
            v1 = 840
        } else {
            fmt.Println(v6, v8, v2)
            if ((!(4294967295 < (v1 << ((v11) & 63)))) || v10) {
                var v50 uint = (v11 - 2)
                v51 := GenS3{F0: (^(v6)), F1: 30}
                v1 = v1
                _ = v50
                _ = v51
            } else {
                if v10 {
                    v52 := GenS1{F0: (("foo bar" + "") < ("" + "hello"))}
                    var v53 string = (("foo bar" + "") + "\n\t")
                    for fi54 := 0; fi54 < 18; fi54++ {
                        switch (v53 + "foo bar") {
                        case "\n\t":
                            v8 = v8
                        case "a":
                            v5 += (v13 ^ 1)
                        }
                        if v10 {
                            var v55 int32 = v13
                            _ = v55
                        }
                        _ = fi54
                    }
                    _ = v52
                    _ = v53
                } else {
                    var v56 int8 = v8
                    var v57 float64 = (-2.25 * 0.1)
                    _ = v56
                    _ = v57
                }
                for fi58 := 0; fi58 < 5; fi58++ {
                    for fi59 := 0; fi59 < 7; fi59++ {
                        if ((!((v1 + (v1 | v9)) == (^(v9)))) && ((fi59 & (fi59 * (fi59 >> (((v11 * (^(v11)))) & 63)))) < (fi59 & -1))) {
                            v60 := v3[v9]
                            v10 = v10
                            v2 = v2
                            if false {
                                v61 := []int{}
                                v62 := GenS3{F0: 18446744073709551615, F1: (fi58 >> (((v11 | (v11 % (((v11 ^ (v11 | v11))) | 1)))) & 63))}
                                delete(v4, fi58)
                                _ = v61
                                _ = v62
                            } else {
                                v6 = v6
                            }
                            v60 = v60
                            v5 = v13
                            _ = v60
                        } else {
                            var v63 string = "hello"
                            var v64 int16 = 32767
                            var v65 int32 = v13
                            fmt.Println(v65, fi58)
                            _ = v63
                            _ = v64
                            _ = v65
                        }
                        _ = fi59
                    }
                    _ = fi58
                }
                if true {
                    v66 := []int{-2147483648, -9223372036854775808, 32768, 4294967295}
                    for fi67 := 0; fi67 < 20; fi67++ {
                        if v10 {
                            switch (fi67 * fi67) {
                            case -608:
                                for fi68 := 0; fi68 < 14; fi68++ {
                                    var v69 int16 = int16(v6)
                                    _ = len(v3)
                                    _ = fi68
                                    _ = v69
                                }
                            case 2147483647:
                                for fi70 := 0; fi70 < 12; fi70++ {
                                    var v71 int64 = 9223372036854775807
                                    _ = fi70
                                    _ = v71
                                }
                            default:
                                fmt.Println(fi67)
                            }
                            for fi72 := 0; fi72 < 17; fi72++ {
                                v66 = append(v66, 47)
                                _ = fi72
                            }
                            fmt.Println(v13)
                        }
                        var v73 bool = (v10 || true)
                        switch (((("hello" + ("" + ("hello" + "a"))) + "\n\t") + "\n\t") + "foo bar") {
                        case "":
                            v74 := GenS2{F0: 645, F1: 56}
                            _ = v74
                        case "\n\t":
                            v75 := GenS3{F0: (v6 << ((v11) & 63)), F1: 32767}
                            _ = v75
                        }
                        _ = fi67
                        _ = v73
                    }
                    _ = v66
                } else {
                    if (!true) {
                        v76 := GenS4{F0: v10}
                        var v77 float32 = (v2 / v2)
                        v78 := GenS5{F0: 9223372036854775807}
                        _ = v76
                        _ = v77
                        _ = v78
                    } else {
                        var v79 uint64 = (v6 ^ v6)
                        _ = v79
                    }
                    if ((2147483647 < 2147483647) && ((-2.25 + 1.5) < 1.5)) {
                        fmt.Println(v1)
                    } else {
                        fmt.Println(v11, v10)
                        switch "a" {
                        case "":
                            v1 = (^(v1))
                        case "a":
                            switch "\n\t" {
                            case "foo bar":
                                switch "hello" {
                                case "":
                                    if v10 {
                                        for fi80 := 0; fi80 < 12; fi80++ {
                                            var v81 uint64 = (v6 ^ (v6 << ((v11) & 63)))
                                            for fi82 := 0; fi82 < 6; fi82++ {
                                                if (!v10) {
                                                    switch ("" + "foo bar") {
                                                    case "\n\t":
                                                        fmt.Println(v6, v9)
                                                    default:
                                                        if v10 {
                                                            v83 := GenS6{F0: (fi82 + fi82), F1: 908, F2: v8}
                                                            var v84 float64 = (float64((v2 / float32(-2.25))) - 1.5)
                                                            v85 := []int{}
                                                            _ = v83
                                                            _ = v84
                                                            _ = v85
                                                        }
                                                    }
                                                    var v86 int64 = -642
                                                    fmt.Println(v86, v1, v10)
                                                    _ = v86
                                                } else {
                                                    var v87 int32 = (v13 / ((2147483647) | 1))
                                                    fmt.Println(fi82)
                                                    _ = v87
                                                }
                                                v9 += (v9 >> ((v11) & 63))
                                                fmt.Println(v9, v2)
                                                _ = fi82
                                            }
                                            _ = fi80
                                            _ = v81
                                        }
                                    } else {
                                        switch 0 {
                                        case 32767:
                                            v88 := []int{-9223372036854775808, 132, 1, 32768}
                                            _ = v88
                                        case 9223372036854775807:
                                            var v89 int = -871
                                            _ = v89
                                        case 42:
                                            var v90 float32 = v2
                                            _ = v90
                                        }
                                        v91 := []int{588, -9223372036854775808, 42, 42}
                                        var v92 int64 = -1
                                        _ = v91
                                        _ = v92
                                    }
                                case "hello":
                                    switch (("" + ("foo bar" + "\n\t")) + "hello") {
                                    case "":
                                        v93 := GenS2{F0: -977, F1: uint8(v11)}
                                        _ = v93
                                    default:
                                        if (v9 != 343) {
                                            v4[2147483647] = -428
                                            if true {
                                                v10 = false
                                                var v94 int8 = 42
                                                _ = v94
                                            }
                                        }
                                    }
                                case "a":
                                    var v95 bool = (371 >= 9223372036854775807)
                                    _ = v95
                                }
                            case "hello":
                                for fi96 := 0; fi96 < 8; fi96++ {
                                    v9 += 121
                                    switch ((("hello" + "foo bar") + "hello") + "a") {
                                    case "":
                                        v6 = v6
                                    default:
                                        var v97 string = "\n\t"
                                        _ = v97
                                    }
                                    _ = fi96
                                }
                            case "\n\t":
                                v98 := v3[2147483648]
                                _ = v98
                            default:
                                var v99 float32 = (1e10 + 0.0)
                                _ = v99
                            }
                        default:
                            for fi100 := 0; fi100 < 20; fi100++ {
                                if (1 < 935) {
                                    v101 := GenS3{F0: (v6 << (((v11 + uint(v9))) & 63)), F1: (fi100 + fi100)}
                                    _ = v101
                                } else {
                                    switch "\n\t" {
                                    case "a":
                                        v102 := GenS7{F0: v11}
                                        _ = v102
                                    case "foo bar":
                                        if (1 > 103) {
                                            switch fi100 {
                                            case 128:
                                                v103 := GenS1{F0: ((65535 != (v11 + 65)) || v10)}
                                                _ = v103
                                            default:
                                                switch ("\n\t" + ("" + "")) {
                                                case "hello":
                                                    fmt.Println(v11)
                                                case "\n\t":
                                                    v11 = v11
                                                case "a":
                                                    var v104 uint32 = (v9 * 0)
                                                    _ = v104
                                                }
                                            }
                                            var v105 int8 = -128
                                            var v106 uint32 = 273
                                            _ = v105
                                            _ = v106
                                        }
                                    }
                                    for fi107 := 0; fi107 < 11; fi107++ {
                                        v108 := GenS8{F0: 0, F1: (v9 | 4294967295)}
                                        _ = fi107
                                        _ = v108
                                    }
                                }
                                var v109 uint = (^(v11))
                                _ = fi100
                                _ = v109
                            }
                        }
                        fmt.Println(v2, v9)
                    }
                    v6 = (v6 ^ v6)
                }
            }
        }
        v11 -= (v11 * (^(v11)))
    }
    if true {
        switch 4294967295 {
        case 1:
            for fi110 := 0; fi110 < 1; fi110++ {
                delete(v3, (v9 & (v1 ^ v9)))
                _ = fi110
            }
        default:
            var v111 int8 = -128
            _ = v111
        }
        if (-9223372036854775808 <= 42) {
            switch 470 {
            case 4294967295:
                v112 := GenS4{F0: (!((v11 << ((v11) & 63)) > 694))}
                _ = v112
            case 42:
                var v113 uint8 = 255
                _ = v113
            }
            var v114 bool = (1 <= v6)
            _ = len(v4)
            _ = v114
        } else {
            if v10 {
                fmt.Println(v2)
            } else {
                v115 := GenS2{F0: -2147483648, F1: 255}
                for fi116 := 0; fi116 < 1; fi116++ {
                    var v117 float32 = float32(0.0)
                    switch 9223372036854775807 {
                    case 291:
                        for fi118 := 0; fi118 < 3; fi118++ {
                            v11 += v11
                            var v119 float32 = v2
                            var v120 float32 = v117
                            _ = fi118
                            _ = v119
                            _ = v120
                        }
                    case 9223372036854775807:
                        for fi121 := 0; fi121 < 10; fi121++ {
                            v10 = v10
                            _ = fi121
                        }
                    case 128:
                        v122 := GenS8{F0: 535, F1: v9}
                        _ = v122
                    default:
                        if v10 {
                            var v123 uint = 0
                            var v124 int64 = 9223372036854775807
                            if (!v10) {
                                var v125 int8 = 51
                                if (!(849 >= (v123 >> ((v11) & 63)))) {
                                    if true {
                                        v2 -= (v2 / (0.0 * v2))
                                        if (-1 < (v5 * (v13 | (v5 << ((v11) & 63))))) {
                                            v126 := GenS9{F0: v8, F1: (^(v8)), F2: 2147483647}
                                            for fi127 := 0; fi127 < 14; fi127++ {
                                                v128 := []int{fi127, fi127}
                                                switch fi127 {
                                                case 0:
                                                    v2 = 0.1
                                                default:
                                                    v125 -= 127
                                                }
                                                switch ("hello" + "a") {
                                                case "hello":
                                                    if (!false) {
                                                        var v129 int32 = int32(v6)
                                                        _ = v129
                                                    }
                                                case "":
                                                    var v130 float32 = v2
                                                    _ = v130
                                                }
                                                _ = fi127
                                                _ = v128
                                            }
                                            var v131 bool = ((v2 * v2) < 1.5)
                                            _ = v126
                                            _ = v131
                                        } else {
                                            fmt.Println(v125, v13, v6)
                                            switch (-(fi116)) {
                                            case -508:
                                                v9 -= (v1 ^ 65535)
                                            }
                                        }
                                        var v132 uint32 = v9
                                        _ = v132
                                    }
                                    var v133 int = 42
                                    if ((!(65535 == (^(v123)))) && false) {
                                        switch "" {
                                        case "foo bar":
                                            fmt.Println(v133, v5)
                                        case "hello":
                                            v134 := GenS10{F0: v5, F1: (v123 | 316)}
                                            _ = v134
                                        case "a":
                                            for fi135 := 0; fi135 < 5; fi135++ {
                                                if v10 {
                                                    v136 := []int{2147483647, -2147483648}
                                                    fmt.Println(v9, v8, v9)
                                                    v8 = 42
                                                    if v10 {
                                                        var v137 string = "\n\t"
                                                        _ = v137
                                                    } else {
                                                        if (!v10) {
                                                            var v138 uint = v123
                                                            v8 = 1
                                                            _ = v138
                                                        }
                                                        if v10 {
                                                            v10 = ((v133 * 4294967295) > (v133 / (((fi116 % (((fi135 % ((v133) | 1))) | 1))) | 1)))
                                                            v139 := GenS11{F0: v6}
                                                            v133 = int(v6)
                                                            _ = v139
                                                        } else {
                                                            v13 = (v5 >> (((v11 & v11)) & 63))
                                                        }
                                                        v140 := GenS12{F0: (!v10), F1: 40, F2: 766}
                                                        _ = v140
                                                    }
                                                    var v141 uint32 = 0
                                                    _ = v136
                                                    _ = v141
                                                }
                                                _ = fi135
                                            }
                                        }
                                        delete(v3, v1)
                                        if ((v11 >> (((v123 >> (((v123 / (((v123 >> ((650) & 63))) | 1))) & 63))) & 63)) < uint(v9)) {
                                            v3[v1] = ("\n\t" + "hello")
                                            var v142 uint32 = v9
                                            _ = v142
                                        } else {
                                            v6 = 42
                                            v1 -= v9
                                        }
                                    } else {
                                        v143 := GenS13{F0: v123, F1: -1, F2: "hello"}
                                        _ = v143
                                    }
                                    _ = v133
                                }
                                _ = v125
                            }
                            _ = v123
                            _ = v124
                        }
                    }
                    var v144 uint16 = 32767
                    _ = fi116
                    _ = v117
                    _ = v144
                }
                var v145 uint = 624
                _ = v115
                _ = v145
            }
        }
    }
    for fi146 := 0; fi146 < 7; fi146++ {
        v147 := GenS12{F0: false, F1: 42, F2: (^(v6))}
        _ = fi146
        _ = v147
    }
    if (((v11 | (v11 >> ((65535) & 63))) > v11) || (0.0 != 0.1)) {
        switch "" {
        case "hello":
            switch 652 {
            case 128:
                v148 := []int{42}
                _ = v148
            case 2147483647:
                for fi149 := 0; fi149 < 18; fi149++ {
                    _ = len(v3)
                    var v150 int16 = -32768
                    for fi151 := 0; fi151 < 6; fi151++ {
                        var v152 int8 = (-(v8))
                        switch "hello" {
                        case "a":
                            var v153 uint16 = 132
                            _ = v153
                        case "hello":
                            fmt.Println(v10)
                        case "":
                            v154 := GenS14{F0: ("hello" + "\n\t"), F1: 1, F2: -1}
                            _ = v154
                        }
                        var v155 int16 = 1000
                        _ = fi151
                        _ = v152
                        _ = v155
                    }
                    _ = fi149
                    _ = v150
                }
            case -1:
                v156 := GenS10{F0: 70, F1: (v11 % ((v11) | 1))}
                _ = v156
            }
        case "":
            v157 := GenS9{F0: -1, F1: v8, F2: 0}
            _ = v157
        case "foo bar":
            v158 := GenS15{F0: 238, F1: 1}
            _ = v158
        default:
            for fi159 := 0; fi159 < 18; fi159++ {
                delete(v3, v9)
                fmt.Println(v8)
                v1 -= v1
                _ = fi159
            }
        }
        v8 -= (^(v8))
        v160 := v3[v1]
        _ = v160
    }
    v11 += v11
    var v161 int16 = -597
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v161:", v161)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v161
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
