package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 int32; F1 uint }
type GenS5 struct { F0 uint8; F1 uint32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint; F1 uint64 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint16; F1 int16; F2 uint }
type GenS10 struct { F0 int32; F1 uint32 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int32; F1 int8; F2 int32 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 int64; F1 int16 }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS17 struct { F0 uint; F1 uint; F2 uint }
type GenS18 struct { F0 uint8 }
func genFn0(p0 uint32) uint32 {
    switch -996 {
    case 2147483647:
        v5 := []int{152}
        _ = v5
    }
    switch "foo bar" {
    case "a":
        fmt.Println(p0)
    case "\n\t":
        var v6 int32 = 2147483647
        _ = v6
    default:
        var v7 uint8 = 2
        _ = v7
    }
    _ = p0
    return 96
}
func genFn1(p0 uint32, p1 bool, p2 int64) string {
    var v8 int8 = 1
    v9 := GenS1{F0: (v8 << ((642) & 63)), F1: (v8 ^ v8)}
    var v10 int8 = (-(v8))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return "foo bar"
}

func main() {
    v1 := make(map[string]string)
    var v2 float64 = ((0.1 - (-2.25 + float64(3.14159))) + ((3.14159 - 1e10) * 1.5))
    var v3 uint16 = 42
    var v4 uint16 = v3
    if false {
        v11 := GenS1{F0: 48, F1: -1}
        switch -886 {
        case 32767:
            v12 := []int{9223372036854775807, -2147483648}
            _ = v12
        default:
            v13 := GenS2{F0: 0, F1: 1000, F2: 1}
            _ = v13
        }
        _ = v11
    } else {
        if false {
            for fi14 := 0; fi14 < 6; fi14++ {
                var v15 float32 = -2.25
                var v16 int32 = 989
                v17 := []int{}
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
        } else {
            v2 = (3.14159 - -2.25)
        }
        v4 += (v3 << ((0) & 63))
        v18 := GenS1{F0: -8, F1: -66}
        _ = v18
    }
    v19 := GenS2{F0: 732, F1: -875, F2: 2}
    if (-746 != 0) {
        v20 := GenS3{F0: 255, F1: 127}
        _ = v20
    }
    v21 := GenS2{F0: uint32(v3), F1: 66, F2: 255}
    fmt.Println(v4, v3)
    fmt.Println(v4)
    if (1 < 1) {
        var v22 uint32 = 678
        v23 := GenS1{F0: 0, F1: -1}
        var v24 int32 = 1
        _ = v22
        _ = v23
        _ = v24
    } else {
        v25 := GenS1{F0: 0, F1: 0}
        _ = v25
    }
    fmt.Println(v4)
    for fi26 := 0; fi26 < 16; fi26++ {
        for fi27 := 0; fi27 < 20; fi27++ {
            var v28 uint32 = genFn0(genFn0(751))
            if (("foo bar" + "foo bar") >= "hello") {
                v29 := GenS2{F0: v28, F1: 0, F2: 255}
                if (!true) {
                    v30 := GenS1{F0: 0, F1: int8(v3)}
                    var v31 int64 = 1
                    fmt.Println(v3)
                    _ = v30
                    _ = v31
                } else {
                    if (!false) {
                        v3 += v3
                    } else {
                        if false {
                            v28 += (v28 ^ v28)
                            var v32 float32 = 0.1
                            _ = v32
                        } else {
                            var v33 uint32 = 2147483648
                            if true {
                                if true {
                                    var v34 string = "hello"
                                    _ = v34
                                } else {
                                    v2 -= ((v2 * -2.25) / v2)
                                    var v35 int16 = 1000
                                    _ = v35
                                }
                                v36 := GenS2{F0: 65535, F1: 0, F2: 1}
                                fmt.Println(v28)
                                _ = v36
                            } else {
                                v37 := []int{(fi26 ^ (fi26 % ((fi26) | 1))), int(v33)}
                                _ = v37
                            }
                            v33 = genFn0((v33 & (v28 & 4294967295)))
                            _ = v33
                        }
                        v38 := GenS4{F0: 10, F1: 255}
                        _ = v38
                    }
                    var v39 int64 = -283
                    fmt.Println(v39, fi26)
                    _ = v39
                }
                switch 393 {
                case 988:
                    v4 = v4
                case 472:
                    if true {
                        v3 += uint16(v28)
                        v40 := GenS5{F0: 255, F1: genFn0(v28)}
                        switch "\n\t" {
                        case "\n\t":
                            v28 += v28
                        case "hello":
                            if (!true) {
                                v2 = (float64(1.5) / float64((0.0 - 1.5)))
                                v41 := GenS6{F0: (v3 & v3)}
                                _ = v41
                            } else {
                                v2 = (v2 * (v2 / v2))
                                fmt.Println(v28, v4, v4)
                                if (-773 <= 0) {
                                    v42 := GenS3{F0: 65535, F1: 1}
                                    var v43 uint = 117
                                    _ = v42
                                    _ = v43
                                }
                            }
                        default:
                            var v44 int8 = 42
                            _ = v44
                        }
                        _ = v40
                    } else {
                        if true {
                            var v45 uint8 = 0
                            if true {
                                v46 := GenS4{F0: 1, F1: 1}
                                v2 -= v2
                                var v47 uint8 = 127
                                var v48 uint64 = 0
                                fmt.Println(v4)
                                v49 := GenS1{F0: 100, F1: 42}
                                switch "hello" {
                                case "a":
                                    v47 += (v45 ^ (v47 >> ((950) & 63)))
                                case "foo bar":
                                    switch ("a" + ("foo bar" + "")) {
                                    case "hello":
                                        var v50 int16 = 42
                                        _ = v50
                                    case "":
                                        if false {
                                            var v51 uint32 = (v28 % ((genFn0(42)) | 1))
                                            _ = v51
                                        } else {
                                            fmt.Println(v48, v28)
                                            v52 := GenS6{F0: 32767}
                                            var v53 int = (fi26 & fi27)
                                            _ = v52
                                            _ = v53
                                        }
                                    case "a":
                                        var v54 int8 = -1
                                        _ = v54
                                    default:
                                        v4 = v4
                                    }
                                }
                                v55 := v1[((("\n\t" + "a") + "\n\t") + "")]
                                v56 := []int{(fi26 ^ 2147483647)}
                                _ = v46
                                _ = v47
                                _ = v48
                                _ = v49
                                _ = v55
                                _ = v56
                            } else {
                                if true {
                                    var v57 uint64 = 1
                                    v58 := []int{32767, -970}
                                    _ = v57
                                    _ = v58
                                } else {
                                    fmt.Println(fi27)
                                    var v59 int32 = -983
                                    _ = v59
                                }
                            }
                            var v60 int32 = 42
                            _ = v45
                            _ = v60
                        }
                        if (!true) {
                            switch fi26 {
                            case 875:
                                fmt.Println(v3)
                            case 127:
                                var v61 float32 = 0.0
                                _ = v61
                            case 0:
                                v4 = (v3 - 32767)
                            default:
                                if false {
                                    v1["\n\t"] = "foo bar"
                                    var v62 uint8 = 41
                                    switch ("foo bar" + "") {
                                    case "a":
                                        switch ((("hello" + "") + "foo bar") + ("" + (("foo bar" + ("foo bar" + "\n\t")) + ""))) {
                                        case "a":
                                            v63 := GenS7{F0: 870, F1: 4294967295}
                                            _ = v63
                                        case "":
                                            var v64 int16 = 0
                                            _ = v64
                                        case "foo bar":
                                            if (true && ((^(fi27)) <= (fi27 ^ 127))) {
                                                fmt.Println(v28, fi26)
                                                v65 := GenS8{F0: ("foo bar" + (("\n\t" + (("foo bar" + "a") + ("a" + ""))) + ""))}
                                                v28 = (v28 ^ (v28 + 2147483647))
                                                v66 := GenS9{F0: v3, F1: -431, F2: 18446744073709551615}
                                                if (false && false) {
                                                    var v67 int32 = 42
                                                    v4 = (v3 + v3)
                                                    _ = v67
                                                } else {
                                                    fmt.Println(v4, v3)
                                                    v2 += v2
                                                    v28 = 2147483648
                                                }
                                                _ = v65
                                                _ = v66
                                            }
                                        }
                                    case "hello":
                                        var v68 int = (fi27 >> ((0) & 63))
                                        _ = v68
                                    default:
                                        switch (fi27 * fi27) {
                                        case 32768:
                                            switch "foo bar" {
                                            case "foo bar":
                                                if (false && false) {
                                                    var v69 int8 = -128
                                                    var v70 uint16 = v4
                                                    if ((^(v62)) != v62) {
                                                        v71 := GenS10{F0: -2147483648, F1: (v28 | v28)}
                                                        _ = v71
                                                    } else {
                                                        v72 := GenS11{F0: 18446744073709551615}
                                                        _ = v72
                                                    }
                                                    var v73 float64 = float64(-2.25)
                                                    v28 += v28
                                                    if ((0.0 < 3.14159) && true) {
                                                        v2 -= v2
                                                    } else {
                                                        v69 = (v69 - 127)
                                                    }
                                                    _ = v69
                                                    _ = v70
                                                    _ = v73
                                                } else {
                                                    v4 = v4
                                                }
                                                v62 = (v62 / ((v62) | 1))
                                            case "a":
                                                var v74 uint16 = v3
                                                _ = v74
                                            }
                                        case 291:
                                            if (132 <= (v62 ^ uint8(v28))) {
                                                v28 = (v28 * 856)
                                                var v75 uint = 18446744073709551615
                                                _ = v75
                                            } else {
                                                fmt.Println(fi26)
                                                v76 := []int{fi27, -9223372036854775808, (fi26 % ((649) | 1))}
                                                if false {
                                                    var v77 uint = 18446744073709551615
                                                    v2 = ((v2 + 3.14159) - v2)
                                                    _ = v77
                                                } else {
                                                    v78 := GenS3{F0: 65535, F1: v62}
                                                    if (505 <= 258) {
                                                        v79 := GenS5{F0: (v62 | (v62 - (v62 >> ((648) & 63)))), F1: v28}
                                                        v3 = (v3 / ((v3) | 1))
                                                        v4 -= v3
                                                        _ = v79
                                                    }
                                                    _ = v78
                                                }
                                                _ = v76
                                            }
                                        case 630:
                                            v80 := GenS12{F0: -647, F1: 100, F2: 1}
                                            _ = v80
                                        }
                                    }
                                    v28 += genFn0(genFn0(v28))
                                    fmt.Println(v62)
                                    _ = v62
                                } else {
                                    var v81 uint8 = 255
                                    var v82 bool = ((2 < 3) || true)
                                    v83 := v1[("" + "\n\t")]
                                    _ = v81
                                    _ = v82
                                    _ = v83
                                }
                            }
                            var v84 string = "\n\t"
                            var v85 uint8 = 1
                            _ = v84
                            _ = v85
                        }
                        switch ("hello" + "hello") {
                        case "\n\t":
                            switch "" {
                            case "foo bar":
                                v3 = v3
                            case "hello":
                                var v86 int64 = -576
                                _ = v86
                            }
                        default:
                            var v87 uint16 = v3
                            _ = v87
                        }
                    }
                case 32768:
                    if (-1 <= int64(v4)) {
                        var v88 string = "hello"
                        v3 = v3
                        _ = v88
                    } else {
                        if ((-(fi26)) <= (fi26 & 2147483647)) {
                            switch (((("" + ("foo bar" + "")) + ((("a" + "\n\t") + ("\n\t" + "foo bar")) + "\n\t")) + "foo bar") + ("hello" + "")) {
                            case "hello":
                                var v89 uint = 214
                                _ = v89
                            default:
                                v90 := v1[("hello" + (((("" + ((("a" + "hello") + (("\n\t" + ("\n\t" + "")) + "hello")) + "\n\t")) + "a") + ("\n\t" + "hello")) + "hello"))]
                                _ = v90
                            }
                        } else {
                            v2 = v2
                            _ = len(v1)
                        }
                        switch (fi27 | fi27) {
                        case 644:
                            v91 := []int{913}
                            _ = v91
                        }
                    }
                    var v92 int8 = -1
                    _ = v92
                default:
                    if (921 == -795) {
                        switch ("hello" + ("\n\t" + (("hello" + (("foo bar" + "\n\t") + "foo bar")) + "foo bar"))) {
                        case "hello":
                            v3 += v3
                        case "foo bar":
                            v93 := GenS5{F0: 255, F1: 65535}
                            _ = v93
                        }
                    }
                }
                _ = v29
            } else {
                var v94 int32 = -270
                v94 -= (v94 ^ int32(v4))
                var v95 int8 = -48
                var v96 uint = 2
                _ = v94
                _ = v95
                _ = v96
            }
            var v97 int16 = 32767
            _ = fi27
            _ = v28
            _ = v97
        }
        for fi98 := 0; fi98 < 4; fi98++ {
            switch 32767 {
            case 32767:
                v2 -= 0.1
            }
            if false {
                switch (fi98 / (((fi98 ^ 4294967295)) | 1)) {
                case 4294967295:
                    fmt.Println(v4, fi98)
                    v99 := []int{(fi26 * 404), fi98}
                    _ = v99
                case 301:
                    var v100 int32 = -2147483648
                    _ = v100
                case 0:
                    v4 = (^(v4))
                    fmt.Println(v4)
                }
            } else {
                v101 := GenS2{F0: 2, F1: 42, F2: 2}
                if false {
                    fmt.Println(v3)
                } else {
                    switch "foo bar" {
                    case "foo bar":
                        var v102 float32 = (-2.25 - 1.5)
                        _ = v102
                    case "":
                        v4 = (v3 / ((v3) | 1))
                    case "hello":
                        v103 := v1["foo bar"]
                        _ = v103
                    }
                    v3 -= (^(v4))
                    v104 := GenS10{F0: 1, F1: 42}
                    _ = v104
                }
                fmt.Println(v3, v4, v2)
                _ = v101
            }
            switch "foo bar" {
            case "":
                switch ("" + "hello") {
                case "":
                    var v105 int32 = -896
                    _ = v105
                case "foo bar":
                    var v106 float32 = ((((-2.25 / 0.1) + (((float32(v2) - (((1.5 / 1.0) / 3.14159) - -2.25)) / -2.25) * 1e10)) * 1e10) * (0.0 + 0.1))
                    _ = v106
                case "hello":
                    fmt.Println(v4, v4, fi26)
                default:
                    v2 -= ((v2 + -2.25) * (1e10 / 1.0))
                }
            case "hello":
                v3 += v3
            default:
                var v107 int = fi26
                _ = v107
            }
            v3 = 42
            _ = fi98
        }
        _ = fi26
    }
    switch (("foo bar" + "") + "foo bar") {
    case "\n\t":
        for fi108 := 0; fi108 < 3; fi108++ {
            switch "\n\t" {
            case "\n\t":
                v109 := []int{}
                _ = v109
            case "hello":
                if (true && false) {
                    for fi110 := 0; fi110 < 9; fi110++ {
                        v4 = (^(v4))
                        _ = fi110
                    }
                }
            default:
                v111 := GenS13{F0: 0}
                _ = v111
            }
            if false {
                for fi112 := 0; fi112 < 2; fi112++ {
                    v3 = (v4 * 42)
                    switch fi112 {
                    case -9223372036854775808:
                        if ("foo bar" >= (("hello" + ("a" + "a")) + "")) {
                            v113 := GenS14{F0: 9223372036854775807, F1: 1000}
                            v2 -= v2
                            _ = v113
                        }
                    case 9223372036854775807:
                        var v114 bool = false
                        _ = v114
                    case 128:
                        switch "foo bar" {
                        case "":
                            var v115 int16 = -945
                            _ = v115
                        case "\n\t":
                            switch fi108 {
                            case 4294967295:
                                v116 := GenS5{F0: 255, F1: 4294967295}
                                _ = v116
                            }
                        }
                    default:
                        v117 := []int{}
                        _ = v117
                    }
                    _ = fi112
                }
                if ((!(uint8(v3) < 227)) && true) {
                    for fi118 := 0; fi118 < 12; fi118++ {
                        fmt.Println(fi108, v2)
                        v4 -= (^(v4))
                        if (((!false) || false) || (fi108 <= fi118)) {
                            v119 := GenS13{F0: 1}
                            var v120 int32 = 1
                            _ = v119
                            _ = v120
                        } else {
                            if (3.14159 <= 3.14159) {
                                var v121 int32 = int32(v4)
                                if true {
                                    var v122 string = "a"
                                    _ = v122
                                } else {
                                    v2 = 1e10
                                    var v123 uint = 65535
                                    v4 += 2
                                    _ = v123
                                }
                                v124 := GenS6{F0: v4}
                                _ = v121
                                _ = v124
                            }
                            var v125 float32 = 3.14159
                            _ = v125
                        }
                        _ = fi118
                    }
                    v4 += (^(v3))
                } else {
                    for fi126 := 0; fi126 < 9; fi126++ {
                        v127 := GenS2{F0: 4294967295, F1: 32767, F2: 0}
                        _ = fi126
                        _ = v127
                    }
                }
            } else {
                v128 := GenS2{F0: 0, F1: 1, F2: 42}
                _ = v128
            }
            _ = fi108
        }
    case "hello":
        v129 := []int{-562}
        _ = v129
    case "a":
        var v130 int = -2147483648
        _ = v130
    default:
        var v131 uint8 = 0
        _ = v131
    }
    var v132 uint16 = v3
    if ((!true) && false) {
        v2 = 3.14159
        var v133 int64 = 42
        _ = v133
    } else {
        switch -948 {
        case 4294967295:
            var v134 float64 = (v2 / v2)
            _ = v134
        default:
            v135 := GenS15{F0: 1}
            _ = v135
        }
        v136 := v1[("" + "")]
        _ = v136
    }
    switch 42 {
    case 0:
        v3 -= 32767
    case -1:
        v3 += (v4 << ((65535) & 63))
    case -2147483648:
        fmt.Println(v2)
    }
    for fi137 := 0; fi137 < 14; fi137++ {
        fmt.Println(v2)
        if ((false && (true || (!(uint8(v132) != 2)))) && (!true)) {
            v138 := GenS16{F0: 127, F1: 127, F2: 42}
            v139 := GenS17{F0: 910, F1: 935, F2: 1}
            _ = v138
            _ = v139
        } else {
            switch "hello" {
            case "foo bar":
                v2 = v2
            default:
                v132 = (v4 | 0)
            }
            for fi140 := 0; fi140 < 20; fi140++ {
                if false {
                    v141 := GenS1{F0: 34, F1: 42}
                    v1[("hello" + (("hello" + "a") + ""))] = ("" + "\n\t")
                    var v142 int16 = 1000
                    _ = v141
                    _ = v142
                }
                _ = fi140
            }
            switch 879 {
            case -471:
                switch "hello" {
                case "":
                    v143 := GenS18{F0: 2}
                    _ = v143
                case "a":
                    if true {
                        if false {
                            v3 = v4
                        }
                        if false {
                            var v144 int16 = -32768
                            for fi145 := 0; fi145 < 15; fi145++ {
                                var v146 uint8 = 255
                                fmt.Println(fi145, v132)
                                _ = fi145
                                _ = v146
                            }
                            v132 = v3
                            _ = v144
                        }
                    }
                case "\n\t":
                    for fi147 := 0; fi147 < 20; fi147++ {
                        switch (fi137 / (((fi147 + (fi147 - fi137))) | 1)) {
                        case 9223372036854775807:
                            v148 := GenS4{F0: -52, F1: 255}
                            v4 = v3
                            _ = v148
                        default:
                            var v149 uint16 = v4
                            _ = v149
                        }
                        var v150 int64 = 386
                        v1["a"] = ""
                        _ = fi147
                        _ = v150
                    }
                default:
                    for fi151 := 0; fi151 < 12; fi151++ {
                        if false {
                            if ((42 <= 936) || (false || ((!true) || false))) {
                                fmt.Println(fi137, fi151, v132)
                            } else {
                                v152 := GenS15{F0: 2147483647}
                                _ = v152
                            }
                        } else {
                            v132 = v3
                        }
                        switch fi137 {
                        case 2147483647:
                            v3 -= v4
                        case 42:
                            var v153 string = "\n\t"
                            _ = v153
                        }
                        _ = fi151
                    }
                }
            case 0:
                v3 = v132
            case 32767:
                var v154 bool = (!true)
                _ = v154
            }
        }
        switch (^(fi137)) {
        case 0:
            var v155 float32 = (0.1 / 0.1)
            _ = v155
        case -9223372036854775808:
            v2 = 1e10
        }
        _ = fi137
    }
    var v156 float32 = 1.5
    for fi157 := 0; fi157 < 5; fi157++ {
        var v158 float32 = float32((v2 * v2))
        _ = fi157
        _ = v158
    }
    var v159 uint = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v132:", v132)
    fmt.Println("v156:", v156)
    fmt.Println("v159:", v159)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v132
    _ = v156
    _ = v159
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
