package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint16; F1 string; F2 uint32 }
type GenS3 struct { F0 int64; F1 uint; F2 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8; F1 uint32; F2 bool }
type GenS6 struct { F0 uint32; F1 uint8; F2 uint8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int32; F1 bool; F2 bool }
type GenS9 struct { F0 int16; F1 int8; F2 int32 }
type GenS10 struct { F0 uint64; F1 int }
type GenS11 struct { F0 int32; F1 int }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 string }
type GenS14 struct { F0 uint16; F1 int8 }
type GenS15 struct { F0 int8; F1 uint64; F2 int }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 bool }
func genFn0() uint {
    if ((!true) || true) {
        switch ("foo bar" + "hello") {
        case "":
            switch "" {
            case "hello":
                switch 0 {
                case -998:
                    v7 := []int{9223372036854775807, 4294967295}
                    _ = v7
                case -1:
                    v8 := GenS1{F0: 4294967295, F1: 105}
                    _ = v8
                case 32767:
                    if (128 != 0) {
                        for fi9 := 0; fi9 < 14; fi9++ {
                            v10 := GenS2{F0: 474, F1: "hello", F2: 2147483647}
                            v11 := GenS1{F0: 4294967295, F1: 130}
                            _ = fi9
                            _ = v10
                            _ = v11
                        }
                        var v12 int16 = 1
                        _ = v12
                    } else {
                        var v13 string = "hello"
                        v14 := []int{875, 1}
                        _ = v13
                        _ = v14
                    }
                }
            case "foo bar":
                var v15 float64 = 1e10
                _ = v15
            case "\n\t":
                var v16 bool = false
                _ = v16
            }
        default:
            v17 := []int{577, 1}
            _ = v17
        }
        for fi18 := 0; fi18 < 3; fi18++ {
            var v19 string = ("foo bar" + "hello")
            _ = fi18
            _ = v19
        }
        if (!true) {
            var v20 uint32 = 65535
            var v21 string = "a"
            v22 := GenS2{F0: 32767, F1: v21, F2: (^(v20))}
            _ = v20
            _ = v21
            _ = v22
        }
    }
    if (794 > 42) {
        v23 := []int{-2147483648, -2147483648, -873}
        var v24 uint16 = 255
        _ = v23
        _ = v24
    } else {
        if (false && true) {
            switch "foo bar" {
            case "hello":
                if (-2.25 != 0.1) {
                    v25 := GenS2{F0: 32767, F1: ("foo bar" + ("a" + "hello")), F2: 1}
                    switch ("a" + "hello") {
                    case "hello":
                        v26 := []int{-2147483648, 32767, 1}
                        _ = v26
                    case "\n\t":
                        if false {
                            var v27 int8 = 100
                            var v28 string = "foo bar"
                            var v29 uint8 = 1
                            _ = v27
                            _ = v28
                            _ = v29
                        } else {
                            var v30 int8 = 29
                            _ = v30
                        }
                    default:
                        var v31 uint32 = 167
                        _ = v31
                    }
                    var v32 int8 = 127
                    _ = v25
                    _ = v32
                }
            }
            if (!true) {
                var v33 float32 = -2.25
                _ = v33
            }
        } else {
            for fi34 := 0; fi34 < 1; fi34++ {
                var v35 float64 = ((0.0 * (3.14159 + 1e10)) + -2.25)
                v35 -= ((v35 / float64((1.5 + float32(1e10)))) + ((0.0 * ((v35 + 3.14159) * (-2.25 * (v35 * float64(float32(1.5)))))) + 3.14159))
                if (fi34 > (fi34 % ((32768) | 1))) {
                    fmt.Println(v35, v35)
                }
                _ = fi34
                _ = v35
            }
            var v36 uint16 = 1
            _ = v36
        }
    }
    var v37 int64 = -438
    _ = v37
    return 0
}

func main() {
    var v1 float64 = 1e10
    var v2 int16 = 42
    var v3 string = "a"
    v4 := make(map[int]uint32)
    var v5 string = v3
    var v6 bool = false
    fmt.Println(v2, v1, v2)
    v38 := GenS3{F0: -295, F1: 65535, F2: 42}
    for fi39 := 0; fi39 < 4; fi39++ {
        for fi40 := 0; fi40 < 1; fi40++ {
            var v41 uint = 1
            v42 := GenS2{F0: 104, F1: "hello", F2: 65535}
            _ = fi40
            _ = v41
            _ = v42
        }
        _ = fi39
    }
    var v43 int64 = 0
    for fi44 := 0; fi44 < 13; fi44++ {
        if (v1 == (((1.5 - v1) + 0.0) * -2.25)) {
            var v45 bool = false
            _ = v45
        }
        v46 := GenS2{F0: 1, F1: ((((v3 + v3) + (v5 + v3)) + "a") + ((v5 + v3) + ((v3 + "\n\t") + v3))), F2: 2}
        if v6 {
            fmt.Println(fi44, v6)
        } else {
            if (((v43 | 1) <= -544) || true) {
                if (1 <= -1) {
                    for fi47 := 0; fi47 < 8; fi47++ {
                        delete(v4, fi44)
                        v5 = v3
                        fmt.Println(v2)
                        _ = fi47
                    }
                    for fi48 := 0; fi48 < 12; fi48++ {
                        fmt.Println(v6, v6, v2)
                        _ = fi48
                    }
                    v3 = "a"
                }
                if (0 > (v43 * (^(v43)))) {
                    v49 := GenS1{F0: 2147483647, F1: 127}
                    v50 := []int{-931}
                    _ = v49
                    _ = v50
                } else {
                    if true {
                        if ((fi44 * 86) < (fi44 * (-(fi44)))) {
                            v51 := []int{fi44}
                            _ = v51
                        } else {
                            if (!(v1 == v1)) {
                                var v52 uint = 0
                                _ = v52
                            }
                        }
                        v6 = false
                        fmt.Println(v5, v2, v2)
                    }
                }
            } else {
                for fi53 := 0; fi53 < 13; fi53++ {
                    switch fi44 {
                    case 128:
                        if true {
                            if (true || (42 > -1)) {
                                fmt.Println(fi53, v6, fi53)
                                v54 := GenS3{F0: v43, F1: 42, F2: 40}
                                if true {
                                    fmt.Println(v3, v3, v6)
                                    v5 = v5
                                    switch v5 {
                                    case "hello":
                                        switch "hello" {
                                        case "foo bar":
                                            switch fi44 {
                                            case 32767:
                                                if (v6 && v6) {
                                                    v43 = -1
                                                    if (!((4294967295 <= 255) || false)) {
                                                        fmt.Println(v3, fi53)
                                                    }
                                                    v6 = false
                                                } else {
                                                    if (!((v6 && false) && (492 == 4294967295))) {
                                                        var v55 string = (v5 + ((((v5 + v5) + (v5 + (v5 + ((v3 + ((((v3 + v3) + "\n\t") + "a") + v5)) + v3)))) + ("\n\t" + "foo bar")) + "foo bar"))
                                                        _ = v55
                                                    } else {
                                                        var v56 string = "foo bar"
                                                        _ = v56
                                                    }
                                                    var v57 int = fi44
                                                    _ = v57
                                                }
                                            }
                                        }
                                    case "foo bar":
                                        fmt.Println(v2, v3, fi44)
                                    case "\n\t":
                                        v58 := []int{-2147483648, 32767, fi53, fi44}
                                        _ = v58
                                    }
                                    var v59 string = "hello"
                                    _ = v59
                                }
                                _ = v54
                            } else {
                                if (!v6) {
                                    if (!v6) {
                                        if (v6 || false) {
                                            if (genFn0() == 255) {
                                                var v60 bool = true
                                                _ = v60
                                            }
                                            v61 := GenS1{F0: 950, F1: 128}
                                            v62 := GenS2{F0: 2, F1: v3, F2: 0}
                                            if true {
                                                var v63 uint = 0
                                                var v64 float32 = (1e10 / 3.14159)
                                                var v65 int32 = 2147483647
                                                _ = v63
                                                _ = v64
                                                _ = v65
                                            } else {
                                                fmt.Println(v5, fi44, v5)
                                            }
                                            var v66 int16 = v2
                                            _ = v61
                                            _ = v62
                                            _ = v66
                                        }
                                    }
                                }
                                var v67 float64 = v1
                                if v6 {
                                    if false {
                                        if (994 < 433) {
                                            fmt.Println(v1, v43, v5)
                                        }
                                        var v68 uint = 42
                                        if true {
                                            switch fi53 {
                                            case 786:
                                                var v69 int16 = v2
                                                _ = v69
                                            }
                                        } else {
                                            var v70 int = fi44
                                            _ = v70
                                        }
                                        _ = v68
                                    } else {
                                        switch v5 {
                                        case "\n\t":
                                            var v71 int = -1
                                            if ((146 >= -1) && (0.0 < 3.14159)) {
                                                if (v71 == (-(fi44))) {
                                                    if v6 {
                                                        var v72 uint64 = 42
                                                        fmt.Println(v5)
                                                        _ = v72
                                                    } else {
                                                        if v6 {
                                                            v3 = "hello"
                                                            v67 = float64(0.0)
                                                            var v73 int16 = (v2 << ((1) & 63))
                                                            v2 = v73
                                                            _ = v73
                                                        } else {
                                                            v67 = 0.0
                                                            if (0 < 0) {
                                                                switch "" {
                                                                case "foo bar":
                                                                    var v74 int64 = 0
                                                                    _ = v74
                                                                case "hello":
                                                                    if ((1 != 42) || v6) {
                                                                        var v75 int8 = -128
                                                                        _ = v75
                                                                    }
                                                                }
                                                                var v76 bool = v6
                                                                _ = v76
                                                            }
                                                            switch v5 {
                                                            case "":
                                                                v77 := GenS4{F0: -128}
                                                                _ = v77
                                                            case "foo bar":
                                                                v6 = false
                                                            case "\n\t":
                                                                if ((v43 / (((v43 - (v43 * v43))) | 1)) <= v43) {
                                                                    if (2147483647 != -467) {
                                                                        var v78 uint64 = 1
                                                                        if v6 {
                                                                            v43 = (-(v43))
                                                                            var v79 string = v3
                                                                            _ = v79
                                                                        }
                                                                        _ = v78
                                                                    }
                                                                    var v80 uint8 = 128
                                                                    var v81 uint32 = 0
                                                                    _ = v80
                                                                    _ = v81
                                                                } else {
                                                                    v82 := GenS5{F0: 97, F1: 65535, F2: (42 >= 1)}
                                                                    v83 := GenS6{F0: 1, F1: 42, F2: 42}
                                                                    fmt.Println(fi53, v6)
                                                                    _ = v82
                                                                    _ = v83
                                                                }
                                                            default:
                                                                var v84 int8 = 42
                                                                _ = v84
                                                            }
                                                            if v6 {
                                                                v6 = (int32(v71) >= 43)
                                                                v85 := GenS7{F0: 1}
                                                                if (!v6) {
                                                                    switch "" {
                                                                    case "":
                                                                        var v86 float64 = v1
                                                                        _ = v86
                                                                    case "foo bar":
                                                                        if false {
                                                                            if v6 {
                                                                                v5 = v3
                                                                                var v87 int64 = (v43 ^ (v43 ^ (v43 % (((v43 / ((-1) | 1))) | 1))))
                                                                                _ = v87
                                                                            } else {
                                                                                v88 := []int{(fi44 ^ fi53)}
                                                                                fmt.Println(v67, v6, v6)
                                                                                if v6 {
                                                                                    if len(v88) > 0 {
                                                                                        v88[0] = v71
                                                                                    }
                                                                                    v67 = v67
                                                                                } else {
                                                                                    v89 := GenS3{F0: int64(v2), F1: genFn0(), F2: 1}
                                                                                    v1 += -2.25
                                                                                    _ = v89
                                                                                }
                                                                                _ = v88
                                                                            }
                                                                        }
                                                                    case "hello":
                                                                        v43 = (^(v43))
                                                                        var v90 uint32 = 42
                                                                        _ = v90
                                                                    default:
                                                                        switch (fi53 * fi53) {
                                                                        case 128:
                                                                            var v91 uint = 18446744073709551615
                                                                            _ = v91
                                                                        case 4294967295:
                                                                            v2 = 1
                                                                        case 42:
                                                                            if true {
                                                                                v3 = ("foo bar" + v3)
                                                                                var v92 uint32 = 785
                                                                                switch 804 {
                                                                                case 571:
                                                                                    v6 = true
                                                                                    v93 := []int{}
                                                                                    _ = v93
                                                                                default:
                                                                                    v67 = float64((0.1 / (1e10 / (((0.0 / 0.1) - ((3.14159 - 0.1) - (float32(float64((0.1 - ((((((3.14159 / -2.25) - ((((3.14159 * 0.1) + -2.25) + -2.25) / 1.0)) - (3.14159 + 0.0)) / (0.0 + 1.5)) - 1e10) - (float32(v1) * 1e10))))) / 1e10))) - 1e10))))
                                                                                }
                                                                                _ = v92
                                                                            } else {
                                                                                var v94 uint32 = 2147483648
                                                                                switch ("" + ("" + ("hello" + "\n\t"))) {
                                                                                case "a":
                                                                                    var v95 uint64 = 754
                                                                                    _ = v95
                                                                                case "":
                                                                                    v5 = "hello"
                                                                                case "foo bar":
                                                                                    switch v5 {
                                                                                    case "foo bar":
                                                                                        if (v3 < "a") {
                                                                                            var v96 int8 = 127
                                                                                            v97 := []int{fi44, (^(v71)), int(v2)}
                                                                                            _ = v96
                                                                                            _ = v97
                                                                                        } else {
                                                                                            v98 := []int{fi44, (fi44 + v71), (fi44 << ((42) & 63)), fi53}
                                                                                            _ = v98
                                                                                        }
                                                                                    case "a":
                                                                                        if v6 {
                                                                                            switch (v71 | (fi53 - (v71 << ((42) & 63)))) {
                                                                                            case 127:
                                                                                                var v99 float32 = (-2.25 + 3.14159)
                                                                                                _ = v99
                                                                                            default:
                                                                                                v67 = v1
                                                                                            }
                                                                                            fmt.Println(v3, v6, fi53)
                                                                                            v100 := GenS8{F0: 801, F1: ((!(((v5 + v5) + ((((v3 + v3) + (("hello" + v3) + (v3 + (v3 + v3)))) + v3) + v3)) == ((v5 + (v3 + v3)) + v5))) || v6), F2: (!(120 >= 77))}
                                                                                            _ = v100
                                                                                        }
                                                                                    case "":
                                                                                        switch 9223372036854775807 {
                                                                                        case 4294967295:
                                                                                            fmt.Println(v3, v2)
                                                                                        case 639:
                                                                                            var v101 uint64 = 42
                                                                                            var v102 int16 = v2
                                                                                            _ = v101
                                                                                            _ = v102
                                                                                        case -624:
                                                                                            var v103 uint16 = 1
                                                                                            _ = v103
                                                                                        default:
                                                                                            if false {
                                                                                                v104 := GenS9{F0: v2, F1: -128, F2: 0}
                                                                                                _ = v104
                                                                                            } else {
                                                                                                switch (v5 + "\n\t") {
                                                                                                case "hello":
                                                                                                    v6 = false
                                                                                                case "":
                                                                                                    v3 = "hello"
                                                                                                default:
                                                                                                    var v105 uint16 = 480
                                                                                                    _ = v105
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        switch "hello" {
                                                                                        case "foo bar":
                                                                                            if (v6 || v6) {
                                                                                                if (255 > 65535) {
                                                                                                    if (-24 < 2147483647) {
                                                                                                        v106 := GenS10{F0: 765, F1: (fi53 * int(v2))}
                                                                                                        var v107 bool = (v6 || v6)
                                                                                                        var v108 bool = (!(-5 <= -39))
                                                                                                        _ = v106
                                                                                                        _ = v107
                                                                                                        _ = v108
                                                                                                    } else {
                                                                                                        var v109 int64 = 0
                                                                                                        _ = v109
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                fmt.Println(fi53, v94, v1)
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                                var v110 int32 = 665
                                                                                _ = v110
                                                                                _ = v94
                                                                            }
                                                                        default:
                                                                            v2 += v2
                                                                        }
                                                                    }
                                                                } else {
                                                                    if v6 {
                                                                        var v111 uint32 = 42
                                                                        v6 = v6
                                                                        fmt.Println(v67, v2)
                                                                        _ = v111
                                                                    }
                                                                    var v112 int32 = -44
                                                                    if ((42 <= (v43 | -1)) || v6) {
                                                                        if (1.5 > 3.14159) {
                                                                            var v113 uint16 = 301
                                                                            _ = v113
                                                                        }
                                                                    } else {
                                                                        v1 = (1.5 * 1e10)
                                                                        v114 := GenS1{F0: 42, F1: 141}
                                                                        if (!(v112 > 898)) {
                                                                            fmt.Println(v71)
                                                                            v3 = v3
                                                                            var v115 int16 = v2
                                                                            var v116 int = fi53
                                                                            _ = v115
                                                                            _ = v116
                                                                        } else {
                                                                            var v117 int = (fi44 & (^(v71)))
                                                                            v118 := GenS4{F0: -1}
                                                                            var v119 int32 = 1
                                                                            _ = v117
                                                                            _ = v118
                                                                            _ = v119
                                                                        }
                                                                        _ = v114
                                                                    }
                                                                    v120 := []int{v71, (fi44 / ((fi53) | 1)), fi53}
                                                                    _ = v112
                                                                    _ = v120
                                                                }
                                                                _ = v85
                                                            }
                                                        }
                                                    }
                                                    v5 = "\n\t"
                                                    if v6 {
                                                        var v121 uint16 = 65535
                                                        var v122 bool = v6
                                                        v123 := GenS11{F0: 2147483647, F1: 4294967295}
                                                        _ = v121
                                                        _ = v122
                                                        _ = v123
                                                    } else {
                                                        var v124 uint64 = 18446744073709551615
                                                        fmt.Println(fi44, v6)
                                                        var v125 int16 = 824
                                                        v71 = (-(v71))
                                                        _ = v124
                                                        _ = v125
                                                    }
                                                    _ = len(v4)
                                                    if (v43 >= (v43 - (v43 / ((v43) | 1)))) {
                                                        v2 += 773
                                                    }
                                                }
                                                switch int(v43) {
                                                case 127:
                                                    switch (fi53 + (v71 ^ (fi44 | (fi53 & (fi44 & 128))))) {
                                                    case 326:
                                                        var v126 uint64 = 2
                                                        _ = v126
                                                    case 32768:
                                                        var v127 int32 = 635
                                                        _ = v127
                                                    default:
                                                        switch fi53 {
                                                        case 730:
                                                            if (((v43 << ((2) & 63)) >= 1) && v6) {
                                                                var v128 float64 = v1
                                                                var v129 uint16 = 42
                                                                var v130 uint64 = 4294967295
                                                                _ = v128
                                                                _ = v129
                                                                _ = v130
                                                            }
                                                        case -2147483648:
                                                            var v131 int64 = -9223372036854775808
                                                            _ = v131
                                                        default:
                                                            v6 = false
                                                        }
                                                    }
                                                case 42:
                                                    v132 := GenS12{F0: 42}
                                                    _ = v132
                                                }
                                                if true {
                                                    if (!(769 <= 165)) {
                                                        v133 := GenS2{F0: 32767, F1: "\n\t", F2: 916}
                                                        if (((1e10 / 1.0) * 0.0) <= -2.25) {
                                                            switch v71 {
                                                            case -458:
                                                                var v134 int32 = -1
                                                                _ = v134
                                                            case -937:
                                                                v2 -= (v2 | (v2 >> ((75) & 63)))
                                                            case 0:
                                                                switch (fi53 >> ((genFn0()) & 63)) {
                                                                case -9223372036854775808:
                                                                    var v135 int64 = v43
                                                                    _ = v135
                                                                case -2147483648:
                                                                    var v136 uint32 = 0
                                                                    _ = v136
                                                                case 137:
                                                                    v5 = v3
                                                                }
                                                            default:
                                                                v71 = 4294967295
                                                            }
                                                            fmt.Println(v71)
                                                            v1 += 1.5
                                                        }
                                                        v137 := GenS13{F0: v5}
                                                        _ = v133
                                                        _ = v137
                                                    }
                                                    v2 = (v2 + v2)
                                                }
                                            } else {
                                                v138 := []int{42, (fi44 * fi44)}
                                                _ = v138
                                            }
                                            if (!false) {
                                                switch fi44 {
                                                case 128:
                                                    var v139 uint16 = 42
                                                    _ = v139
                                                case 58:
                                                    fmt.Println(v1)
                                                default:
                                                    var v140 uint8 = 158
                                                    _ = v140
                                                }
                                                switch fi53 {
                                                case -2147483648:
                                                    switch (v71 ^ (fi53 + 764)) {
                                                    case -9223372036854775808:
                                                        var v141 float64 = v1
                                                        _ = v141
                                                    default:
                                                        var v142 float32 = 1.5
                                                        _ = v142
                                                    }
                                                case 385:
                                                    if (191 >= 2) {
                                                        fmt.Println(v1, fi53, fi44)
                                                        var v143 int8 = -1
                                                        _ = v143
                                                    }
                                                case 9223372036854775807:
                                                    var v144 uint64 = 77
                                                    _ = v144
                                                default:
                                                    var v145 uint8 = 120
                                                    _ = v145
                                                }
                                                if false {
                                                    fmt.Println(v43, v1, v1)
                                                    switch (fi53 - fi53) {
                                                    case 9223372036854775807:
                                                        v1 += -2.25
                                                    case -9223372036854775808:
                                                        delete(v4, 2147483647)
                                                        switch fi53 {
                                                        case 128:
                                                            v146 := []int{}
                                                            _ = v146
                                                        case 585:
                                                            if true {
                                                                v147 := GenS1{F0: 2147483647, F1: 101}
                                                                fmt.Println(v43, fi53, v1)
                                                                var v148 float32 = 0.0
                                                                _ = v147
                                                                _ = v148
                                                            } else {
                                                                v149 := []int{fi53, (fi53 - fi44)}
                                                                fmt.Println(fi44)
                                                                v67 += float64(1.5)
                                                                fmt.Println(v2)
                                                                _ = v149
                                                            }
                                                        case 32768:
                                                            var v150 bool = true
                                                            _ = v150
                                                        }
                                                        v3 = v5
                                                        if (1 > 0) {
                                                            if ("" > v3) {
                                                                v43 = 1
                                                                v151 := []int{(fi44 << ((792) & 63)), fi53, fi44}
                                                                _ = v151
                                                            }
                                                            if v6 {
                                                                var v152 uint16 = 2
                                                                _ = v152
                                                            } else {
                                                                switch fi53 {
                                                                case 32768:
                                                                    if (true || (956 > 2)) {
                                                                        v1 = (float64(float32(v67)) / (1.5 * (-2.25 / float64(-2.25))))
                                                                        var v153 int8 = 0
                                                                        fmt.Println(fi44, v3)
                                                                        _ = v153
                                                                    } else {
                                                                        var v154 uint32 = 42
                                                                        _ = v154
                                                                    }
                                                                    v155 := GenS14{F0: 2, F1: 100}
                                                                    _ = v155
                                                                }
                                                            }
                                                        } else {
                                                            if ((v6 || (2 != 18446744073709551615)) && (v67 < v67)) {
                                                                fmt.Println(v2, fi53, v67)
                                                                v156 := GenS11{F0: int32(v71), F1: (v71 >> ((genFn0()) & 63))}
                                                                var v157 int = (v71 * (v71 * (fi53 >> ((332) & 63))))
                                                                _ = v156
                                                                _ = v157
                                                            }
                                                            fmt.Println(v5, v6, v2)
                                                            if (!v6) {
                                                                v158 := []int{(^(v71))}
                                                                _ = v158
                                                            }
                                                        }
                                                        switch v3 {
                                                        case "foo bar":
                                                            v159 := GenS4{F0: -128}
                                                            _ = v159
                                                        case "hello":
                                                            v160 := GenS6{F0: 228, F1: 200, F2: 1}
                                                            _ = v160
                                                        }
                                                        if (v5 != "\n\t") {
                                                            v161 := []int{-229}
                                                            switch v5 {
                                                            case "":
                                                                var v162 uint = genFn0()
                                                                _ = v162
                                                            default:
                                                                if v6 {
                                                                    v163 := v4[(fi53 ^ fi53)]
                                                                    v71 = (fi53 / ((fi53) | 1))
                                                                    _ = v163
                                                                }
                                                            }
                                                            v4[(v71 | (fi53 >> ((255) & 63)))] = 2147483648
                                                            var v164 uint16 = 134
                                                            _ = v161
                                                            _ = v164
                                                        } else {
                                                            var v165 bool = false
                                                            _ = v165
                                                        }
                                                    }
                                                }
                                                v1 = float64(1e10)
                                                v6 = (-91 == 42)
                                                v43 -= v43
                                                switch (^(fi53)) {
                                                case -9223372036854775808:
                                                    var v166 float32 = (0.0 * (((0.1 - 3.14159) / float32(1e10)) + (1.5 - float32(((1e10 + (float64(float32((((((((v1 * v1) - float64(3.14159)) + v1) + ((0.0 / (v67 - v67)) + v67)) - 1e10) / float64((0.1 / -2.25))) / 1.0))) * float64(float32(v67)))) * (float64(float32(float64(0.0))) + (1e10 * float64(1.5))))))))
                                                    _ = v166
                                                case 2147483647:
                                                    v2 = 75
                                                case 848:
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        switch 32767 {
                                                        case 0:
                                                            v167 := []int{(fi44 - fi44), fi53, 99}
                                                            _ = v167
                                                        case 128:
                                                            var v168 int64 = 0
                                                            _ = v168
                                                        case 2147483647:
                                                            v2 -= (-(v2))
                                                        default:
                                                            var v169 int16 = v2
                                                            _ = v169
                                                        }
                                                    default:
                                                        if v6 {
                                                            var v170 uint32 = 0
                                                            _ = v170
                                                        }
                                                    }
                                                default:
                                                    if v6 {
                                                        v5 = v3
                                                    }
                                                    delete(v4, v71)
                                                    v171 := []int{230}
                                                    var v172 uint32 = 0
                                                    var v173 int8 = -1
                                                    _ = v171
                                                    _ = v172
                                                    _ = v173
                                                }
                                                switch (v5 + (v5 + v5)) {
                                                case "hello":
                                                    if v6 {
                                                        v174 := []int{v71, 341, (fi44 >> ((255) & 63))}
                                                        _ = v174
                                                    } else {
                                                        v175 := []int{(v71 / ((fi53) | 1)), v71, fi44, -2147483648}
                                                        _ = v175
                                                    }
                                                case "":
                                                    var v176 int = v71
                                                    _ = v176
                                                case "foo bar":
                                                    if ((((((-1 <= 2147483647) && v6) && v6) || v6) || v6) || ((v43 >= (^(v43))) && v6)) {
                                                        v177 := GenS15{F0: -1, F1: 4294967295, F2: fi53}
                                                        v71 -= (fi53 ^ (v71 | 128))
                                                        _ = v177
                                                    }
                                                }
                                                v178 := []int{528, -936, (fi44 ^ (-(v71))), fi44}
                                                var v179 bool = (!v6)
                                                v180 := GenS16{F0: (v2 ^ (^(v2)))}
                                                v5 = v3
                                                _ = v178
                                                _ = v179
                                                _ = v180
                                            }
                                            _ = v71
                                        }
                                    }
                                    v6 = v6
                                }
                                if false {
                                    fmt.Println(v67, v6, v2)
                                    v181 := []int{188, 32767}
                                    var v182 int8 = int8(fi44)
                                    _ = v181
                                    _ = v182
                                } else {
                                    if (0.1 < 0.1) {
                                        switch (fi53 - fi44) {
                                        case 4294967295:
                                            switch (fi44 & (fi53 ^ fi44)) {
                                            case 0:
                                                v43 = (v43 >> ((4294967295) & 63))
                                            case -765:
                                                fmt.Println(v1, v67)
                                            default:
                                                var v183 uint = 0
                                                _ = v183
                                            }
                                        case 0:
                                            if (!(4294967295 >= 1)) {
                                                switch v3 {
                                                case "a":
                                                    var v184 uint32 = 805
                                                    _ = v184
                                                case "hello":
                                                    var v185 int32 = 0
                                                    _ = v185
                                                case "foo bar":
                                                    switch (fi44 ^ 967) {
                                                    case 0:
                                                        v43 += (^(v43))
                                                        if true {
                                                            v4[-512] = 42
                                                            var v186 int8 = 100
                                                            _ = len(v4)
                                                            _ = v186
                                                        } else {
                                                            var v187 uint64 = 343
                                                            var v188 float64 = 3.14159
                                                            _ = v187
                                                            _ = v188
                                                        }
                                                    case 4294967295:
                                                        switch v5 {
                                                        case "a":
                                                            v43 = v43
                                                        default:
                                                            fmt.Println(v3, fi53)
                                                        }
                                                    case -9223372036854775808:
                                                        v43 += v43
                                                    }
                                                }
                                                v189 := []int{fi44, int(v2)}
                                                _ = v189
                                            } else {
                                                var v190 uint64 = 18446744073709551615
                                                _ = v190
                                            }
                                        }
                                        v43 -= 0
                                    }
                                    if (0 != 479) {
                                        if (((v5 + "foo bar") + v3) >= v5) {
                                            fmt.Println(fi44, v6, v5)
                                            v67 = -2.25
                                        } else {
                                            v4[(fi53 * 707)] = 736
                                        }
                                        v191 := GenS15{F0: -85, F1: 2, F2: 128}
                                        switch v3 {
                                        case "foo bar":
                                            if (!(v5 >= v3)) {
                                                if ((genFn0() == genFn0()) && (false && (!v6))) {
                                                    fmt.Println(v5)
                                                } else {
                                                    if v6 {
                                                        v1 += v1
                                                    }
                                                    if v6 {
                                                        var v192 int64 = (v43 * (v43 + (-(v43))))
                                                        v193 := GenS17{F0: v6}
                                                        v67 += v1
                                                        _ = v192
                                                        _ = v193
                                                    } else {
                                                        var v194 bool = true
                                                        _ = v194
                                                    }
                                                }
                                                v3 = v3
                                                switch fi53 {
                                                case 32767:
                                                    if (-2147483648 <= 334) {
                                                        var v195 int = fi44
                                                        if true {
                                                            var v196 int32 = 405
                                                            v6 = (v6 || (true && (!true)))
                                                            _ = v196
                                                        }
                                                        _ = v195
                                                    }
                                                case 2147483647:
                                                    v197 := GenS10{F0: 2, F1: 0}
                                                    _ = v197
                                                case -79:
                                                    fmt.Println(fi44, v5)
                                                    v198 := []int{}
                                                    v1 = 1.5
                                                    delete(v4, fi53)
                                                    _ = v198
                                                }
                                                var v199 int = fi53
                                                v200 := GenS5{F0: 255, F1: 677, F2: false}
                                                _ = v199
                                                _ = v200
                                            } else {
                                                if v6 {
                                                    v3 = (("hello" + (v5 + v5)) + (((v3 + ("hello" + "hello")) + (v5 + "hello")) + v3))
                                                    fmt.Println(v6)
                                                }
                                            }
                                        }
                                        if (v6 || v6) {
                                            var v201 uint32 = 1
                                            _ = v201
                                        }
                                        if (!v6) {
                                            var v202 int16 = v2
                                            if ((v43 - (v43 / (((v43 % (((^(v43))) | 1))) | 1))) == 9223372036854775807) {
                                                switch v3 {
                                                case "hello":
                                                    v203 := GenS12{F0: 0}
                                                    _ = v203
                                                case "\n\t":
                                                    v6 = v6
                                                case "a":
                                                    v2 = (v202 & v202)
                                                default:
                                                    v204 := GenS15{F0: -93, F1: 2, F2: (fi53 + fi44)}
                                                    var v205 uint = 2
                                                    v5 = v5
                                                    _ = v204
                                                    _ = v205
                                                }
                                                var v206 int64 = v43
                                                v207 := []int{fi53, (fi44 * (fi53 >> ((42) & 63))), fi53, (fi44 / (((fi53 & fi53)) | 1))}
                                                var v208 string = (v5 + v3)
                                                _ = v206
                                                _ = v207
                                                _ = v208
                                            }
                                            v1 -= v1
                                            _ = v202
                                        }
                                        var v209 string = v5
                                        _ = v191
                                        _ = v209
                                    }
                                }
                                _ = v67
                            }
                        } else {
                            var v210 float32 = 1e10
                            _ = v210
                        }
                    }
                    var v211 uint16 = 347
                    _ = fi53
                    _ = v211
                }
            }
        }
        _ = fi44
        _ = v46
    }
    v212 := []int{32767}
    fmt.Println(v2, v3, v2)
    switch v3 {
    case "foo bar":
        var v213 int32 = 417
        _ = v213
    default:
        switch v3 {
        case "":
            for fi214 := 0; fi214 < 4; fi214++ {
                if v6 {
                    v6 = ((-(fi214)) < fi214)
                    if len(v212) >= 2 {
                        v212 = v212[1:len(v212)]
                    }
                }
                switch fi214 {
                case 791:
                    var v215 string = v3
                    _ = v215
                default:
                    var v216 uint64 = 110
                    _ = v216
                }
                v3 = v3
                _ = fi214
            }
        case "\n\t":
            for fi217 := 0; fi217 < 17; fi217++ {
                if (v6 && (2 <= 372)) {
                    var v218 int = fi217
                    v218 = fi217
                    _ = v218
                }
                _ = fi217
            }
        case "a":
            v219 := GenS8{F0: 975, F1: (true || v6), F2: v6}
            _ = v219
        default:
            var v220 bool = v6
            _ = v220
        }
    }
    v3 = "foo bar"
    _ = len(v212)
    v221 := GenS10{F0: 42, F1: 32767}
    for fi222 := 0; fi222 < 1; fi222++ {
        if len(v212) > 0 {
            v212[0] = fi222
        }
        _ = fi222
    }
    _ = len(v4)
    var v223 int8 = 0
    for fi224 := 0; fi224 < 1; fi224++ {
        var v225 int16 = v2
        var v226 float64 = ((v1 * (0.1 * -2.25)) - float64(1e10))
        _ = fi224
        _ = v225
        _ = v226
    }
    var v227 uint16 = 65535
    for fi228 := 0; fi228 < 18; fi228++ {
        if false {
            for fi229 := 0; fi229 < 20; fi229++ {
                v43 = 1
                _ = len(v212)
                var v230 int = 128
                _ = fi229
                _ = v230
            }
            v2 = v2
        } else {
            switch ("" + v5) {
            case "hello":
                v212 = append(v212, (fi228 % (((fi228 >> ((65535) & 63))) | 1)))
            case "a":
                for fi231 := 0; fi231 < 2; fi231++ {
                    if len(v212) > 0 {
                        v212[0] = (fi228 % (((fi228 % ((fi231) | 1))) | 1))
                    }
                    v223 = v223
                    if false {
                        fmt.Println(v227)
                        _ = len(v212)
                        var v232 int8 = (v223 | v223)
                        _ = v232
                    } else {
                        fmt.Println(v6, v43, fi231)
                        v6 = ((fi231 + fi228) != fi231)
                    }
                    _ = fi231
                }
            default:
                v212 = append(v212, (fi228 ^ fi228))
            }
            if (-128 <= v223) {
                if v6 {
                    var v233 int32 = -2147483648
                    _ = v233
                }
                if false {
                    v234 := GenS11{F0: -187, F1: 32768}
                    _ = v234
                } else {
                    var v235 float32 = -2.25
                    _ = v235
                }
            }
            var v236 uint32 = 2
            _ = v236
        }
        _ = fi228
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v212:", len(v212), v212)
    fmt.Println("v223:", v223)
    fmt.Println("v227:", v227)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v212
    _ = v221
    _ = v223
    _ = v227
    _ = v3
    _ = v38
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
