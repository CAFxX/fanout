package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 int8; F1 uint32; F2 int64 }
type GenS5 struct { F0 uint64; F1 int; F2 uint8 }
type GenS6 struct { F0 uint16; F1 uint32; F2 bool }
type GenS7 struct { F0 int8; F1 uint64; F2 string }
type GenS8 struct { F0 uint64; F1 int }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int64; F1 uint32 }
type GenS11 struct { F0 int16; F1 bool }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 uint16; F1 uint8; F2 bool }
type GenS16 struct { F0 int8; F1 uint64 }
type GenS17 struct { F0 bool; F1 uint32; F2 string }
type GenS18 struct { F0 uint }
type GenS19 struct { F0 uint64; F1 uint64 }
type GenS20 struct { F0 uint8; F1 int }
type GenS21 struct { F0 int }
func genFn0(p0 string) int64 {
    var v6 int8 = 118
    fmt.Println(p0)
    switch p0 {
    case "foo bar":
        p0 = "a"
    case "":
        v6 = v6
    case "\n\t":
        v6 += (v6 * -128)
    }
    if true {
        if false {
            var v7 uint8 = 0
            for fi8 := 0; fi8 < 7; fi8++ {
                for fi9 := 0; fi9 < 10; fi9++ {
                    fmt.Println(fi8, fi8, v6)
                    _ = fi9
                }
                for fi10 := 0; fi10 < 19; fi10++ {
                    fmt.Println(v6, fi10)
                    v7 = 2
                    if (3.14159 != -2.25) {
                        var v11 uint = 795
                        v12 := GenS1{F0: 617, F1: (v11 + 4294967295), F2: 763}
                        v13 := GenS1{F0: 255, F1: (v11 * (^(v11))), F2: 793}
                        _ = v11
                        _ = v12
                        _ = v13
                    } else {
                        var v14 uint = 18446744073709551615
                        var v15 uint16 = 65535
                        if true {
                            var v16 uint64 = 133
                            var v17 int64 = 0
                            fmt.Println(v6)
                            _ = v16
                            _ = v17
                        } else {
                            if true {
                                var v18 bool = true
                                v18 = v18
                                var v19 bool = ("hello" >= p0)
                                if (42 != 768) {
                                    v20 := GenS1{F0: v15, F1: v14, F2: 717}
                                    v21 := GenS1{F0: v15, F1: v14, F2: 18446744073709551615}
                                    v15 = (v15 << ((v14) & 63))
                                    _ = v20
                                    _ = v21
                                }
                                var v22 float32 = (1e10 * 3.14159)
                                var v23 uint16 = (v15 / (((v15 ^ 0)) | 1))
                                fmt.Println(v22, v19, v18)
                                var v24 int32 = 0
                                if v18 {
                                    v19 = v19
                                    v25 := []int{fi10}
                                    var v26 uint8 = (^(v7))
                                    _ = v25
                                    _ = v26
                                }
                                _ = v18
                                _ = v19
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                            if false {
                                if true {
                                    v27 := GenS2{F0: 0}
                                    switch int(v7) {
                                    case 128:
                                        v28 := GenS1{F0: (v15 / ((2) | 1)), F1: 255, F2: 130}
                                        _ = v28
                                    case 42:
                                        var v29 int64 = 401
                                        _ = v29
                                    }
                                    _ = v27
                                }
                                var v30 uint64 = 1
                                if (!(!true)) {
                                    v15 += 984
                                    p0 = (p0 + (p0 + p0))
                                    fmt.Println(v14, v14, v6)
                                }
                                fmt.Println(fi10, v7, p0)
                                var v31 int8 = 54
                                var v32 uint = v14
                                _ = v30
                                _ = v31
                                _ = v32
                            } else {
                                if false {
                                    var v33 int16 = -406
                                    _ = v33
                                }
                                var v34 float32 = (0.0 + 0.0)
                                v14 -= v14
                                _ = v34
                            }
                            if ((0.0 - float32(((float64(-2.25) + ((-2.25 - -2.25) + (0.1 * (0.0 / 1e10)))) + 1e10))) <= 1.5) {
                                v35 := []int{(^(fi10)), fi8, fi8}
                                if (true || false) {
                                    v14 = 2
                                    switch fi8 {
                                    case -35:
                                        if len(v35) > 0 {
                                            v35[0] = -9223372036854775808
                                        }
                                    case 575:
                                        var v36 uint64 = 1
                                        _ = v36
                                    case 127:
                                        p0 = "\n\t"
                                    }
                                    var v37 int = 473
                                    _ = v37
                                } else {
                                    fmt.Println(v15)
                                    if len(v35) >= 2 {
                                        v35 = v35[1:len(v35)]
                                    }
                                    var v38 uint64 = 42
                                    _ = v38
                                }
                                v6 = (-(v6))
                                v39 := GenS3{F0: (v14 ^ (v14 >> ((v14) & 63))), F1: v15}
                                var v40 uint32 = 2147483648
                                _ = v35
                                _ = v39
                                _ = v40
                            }
                        }
                        _ = v14
                        _ = v15
                    }
                    _ = fi10
                }
                _ = fi8
            }
            _ = v7
        } else {
            for fi41 := 0; fi41 < 10; fi41++ {
                v6 += v6
                v42 := GenS3{F0: 2, F1: 627}
                v6 = v6
                _ = fi41
                _ = v42
            }
        }
        var v43 int16 = -32768
        for fi44 := 0; fi44 < 1; fi44++ {
            v6 = 100
            if true {
                v45 := []int{-716, (^(fi44)), fi44}
                switch (p0 + p0) {
                case "foo bar":
                    for fi46 := 0; fi46 < 4; fi46++ {
                        fmt.Println(fi44, fi44)
                        switch fi44 {
                        case -285:
                            if (-2147483648 <= 503) {
                                var v47 uint32 = 4294967295
                                if (1.5 != 3.14159) {
                                    switch p0 {
                                    case "":
                                        switch "\n\t" {
                                        case "hello":
                                            if (!true) {
                                                switch int(v47) {
                                                case 4294967295:
                                                    v45 = append(v45, (fi44 << ((65535) & 63)))
                                                }
                                                v48 := GenS3{F0: 42, F1: 226}
                                                if true {
                                                    v43 = (v43 - 42)
                                                    v49 := GenS1{F0: 1, F1: 18446744073709551615, F2: 4294967295}
                                                    var v50 int64 = int64(v43)
                                                    _ = len(v45)
                                                    if len(v45) > 0 {
                                                        v45[0] = (fi44 & fi44)
                                                    }
                                                    _ = v49
                                                    _ = v50
                                                }
                                                _ = v48
                                            }
                                        default:
                                            var v51 uint32 = v47
                                            _ = v51
                                        }
                                    }
                                    fmt.Println(v6, fi46)
                                    fmt.Println(p0)
                                }
                                _ = v47
                            } else {
                                var v52 float32 = ((-2.25 * float32(1e10)) + ((-2.25 - 0.0) * -2.25))
                                _ = v52
                            }
                            switch p0 {
                            case "foo bar":
                                fmt.Println(v6)
                            case "\n\t":
                                var v53 uint32 = 5
                                _ = v53
                            case "hello":
                                var v54 int16 = -32768
                                _ = v54
                            }
                        case 794:
                            p0 = (("a" + p0) + p0)
                        }
                        var v55 uint64 = 4294967295
                        _ = fi46
                        _ = v55
                    }
                case "a":
                    if (v43 < (^(v43))) {
                        _ = len(v45)
                        switch "a" {
                        case "":
                            v56 := GenS4{F0: v6, F1: 324, F2: 959}
                            _ = v56
                        }
                    } else {
                        fmt.Println(v6, v6, fi44)
                        var v57 uint32 = 1
                        _ = v57
                    }
                case "\n\t":
                    if (!true) {
                        fmt.Println(v6)
                        var v58 int64 = -1
                        var v59 int8 = (-(v6))
                        _ = v58
                        _ = v59
                    }
                }
                _ = v45
            }
            _ = fi44
        }
        _ = v43
    } else {
        p0 = p0
        var v60 string = ("hello" + p0)
        _ = v60
    }
    for fi61 := 0; fi61 < 6; fi61++ {
        var v62 string = p0
        v63 := []int{(fi61 >> ((2) & 63))}
        _ = fi61
        _ = v62
        _ = v63
    }
    _ = p0
    _ = v6
    return -820
}
func genFn1(p0 uint32) uint16 {
    switch 4294967295 {
    case -9223372036854775808:
        p0 -= p0
    case 128:
        p0 = (p0 & (p0 >> ((255) & 63)))
    case 32767:
        var v64 uint32 = (^(p0))
        _ = v64
    default:
        for fi65 := 0; fi65 < 6; fi65++ {
            p0 -= p0
            var v66 uint32 = (^(p0))
            _ = fi65
            _ = v66
        }
    }
    p0 -= (p0 | (p0 & 568))
    var v67 float64 = -2.25
    if false {
        for fi68 := 0; fi68 < 2; fi68++ {
            for fi69 := 0; fi69 < 16; fi69++ {
                v70 := GenS2{F0: -172}
                v67 -= v67
                v71 := GenS3{F0: 65535, F1: 1}
                _ = fi69
                _ = v70
                _ = v71
            }
            _ = fi68
        }
    } else {
        v72 := []int{4294967295, 32767}
        if false {
            if (("" < "hello") && (false || true)) {
                for fi73 := 0; fi73 < 9; fi73++ {
                    v67 -= float64(1.5)
                    _ = fi73
                }
                fmt.Println(p0, v67)
            } else {
                fmt.Println(v67)
                fmt.Println(p0)
            }
        }
        v67 = (1.5 - v67)
        _ = v72
    }
    v67 -= v67
    _ = p0
    _ = v67
    return 32767
}

func main() {
    var v1 int32 = 42
    v2 := make(map[string]uint32)
    var v3 string = ("" + "foo bar")
    v4 := make(map[int]string)
    v5 := []int{}
    var v74 string = "hello"
    v74 = ((v3 + v74) + ((("" + (v3 + (((v3 + ((v74 + v3) + ("foo bar" + ("" + "foo bar")))) + v3) + v3))) + v74) + v3))
    if ((!true) && (true && (4294967295 != 42))) {
        _ = len(v4)
        if (!(9223372036854775807 <= -1)) {
            for fi75 := 0; fi75 < 9; fi75++ {
                for fi76 := 0; fi76 < 4; fi76++ {
                    switch v74 {
                    case "\n\t":
                        v77 := GenS4{F0: -128, F1: 817, F2: 319}
                        _ = v77
                    case "hello":
                        var v78 int32 = v1
                        _ = v78
                    case "a":
                        if false {
                            if (255 < 0) {
                                if (!false) {
                                    if true {
                                        _ = len(v5)
                                        var v79 int8 = -1
                                        _ = v79
                                    }
                                    var v80 string = v3
                                    v5 = append(v5, 4294967295)
                                    _ = v80
                                }
                                var v81 uint16 = 65535
                                v2[v74] = 2147483647
                                _ = v81
                            }
                            v82 := GenS3{F0: 1, F1: 42}
                            v74 = v3
                            _ = v82
                        } else {
                            _ = len(v2)
                            switch ((("foo bar" + (v3 + ("hello" + v74))) + v3) + ((("" + "hello") + "hello") + (v3 + v74))) {
                            case "":
                                var v83 bool = (!false)
                                switch v3 {
                                case "foo bar":
                                    if ((v1 + v1) == v1) {
                                        if (!v83) {
                                            var v84 int = fi75
                                            v3 = ""
                                            v3 = "a"
                                            v1 = v1
                                            var v85 bool = true
                                            _ = v84
                                            _ = v85
                                        } else {
                                            var v86 int64 = -89
                                            v5 = append(v5, int(v86))
                                            fmt.Println(v74, fi75, fi76)
                                            _ = v86
                                        }
                                        v87 := v4[fi76]
                                        _ = v87
                                    }
                                default:
                                    fmt.Println(v1, v1)
                                }
                                if v83 {
                                    fmt.Println(v74, fi76)
                                    switch (v74 + "") {
                                    case "foo bar":
                                        v1 += (^(v1))
                                    case "":
                                        var v88 int64 = 0
                                        _ = v88
                                    case "a":
                                        v1 -= v1
                                    }
                                } else {
                                    var v89 int = 128
                                    v89 += -503
                                    var v90 int32 = -798
                                    _ = v89
                                    _ = v90
                                }
                                _ = v83
                            case "foo bar":
                                var v91 uint16 = 0
                                _ = v91
                            case "hello":
                                var v92 int = fi76
                                _ = v92
                            }
                        }
                    default:
                        v93 := GenS5{F0: 18446744073709551615, F1: fi76, F2: 1}
                        _ = v93
                    }
                    var v94 uint = 2
                    _ = fi76
                    _ = v94
                }
                var v95 float64 = -2.25
                switch (fi75 * fi75) {
                case -835:
                    if (4294967295 == 42) {
                        v96 := GenS4{F0: -39, F1: 1, F2: 0}
                        if true {
                            fmt.Println(v1, v1, fi75)
                            switch v74 {
                            case "a":
                                var v97 float32 = 0.0
                                _ = v97
                            default:
                                if true {
                                    v95 += v95
                                    if len(v5) > 0 {
                                        v5[0] = fi75
                                    }
                                    fmt.Println(fi75, v1)
                                }
                            }
                        } else {
                            var v98 float64 = v95
                            _ = v98
                        }
                        _ = v96
                    } else {
                        v99 := GenS5{F0: 2, F1: fi75, F2: 85}
                        _ = v99
                    }
                }
                _ = fi75
                _ = v95
            }
            for fi100 := 0; fi100 < 13; fi100++ {
                if (2147483648 >= 2147483648) {
                    var v101 int16 = -486
                    switch "" {
                    case "a":
                        v3 = v74
                    default:
                        if false {
                            v102 := GenS6{F0: 255, F1: 65535, F2: true}
                            _ = len(v5)
                            _ = v102
                        }
                    }
                    if (508 <= 358) {
                        fmt.Println(v74)
                        var v103 int16 = (-(v101))
                        _ = v103
                    }
                    _ = v101
                } else {
                    v104 := v4[fi100]
                    _ = v104
                }
                fmt.Println(v3, v1)
                var v105 uint16 = 65535
                _ = fi100
                _ = v105
            }
        }
        for fi106 := 0; fi106 < 3; fi106++ {
            var v107 bool = ((v1 < v1) && true)
            var v108 uint64 = 31
            _ = fi106
            _ = v107
            _ = v108
        }
    } else {
        v109 := GenS3{F0: 866, F1: 138}
        if true {
            if (2147483647 != 903) {
                for fi110 := 0; fi110 < 1; fi110++ {
                    v111 := GenS6{F0: genFn1(2), F1: 543, F2: ((v74 + "foo bar") == v3)}
                    _ = fi110
                    _ = v111
                }
                if true {
                    switch v3 {
                    case "foo bar":
                        var v112 int = 128
                        _ = v112
                    default:
                        for fi113 := 0; fi113 < 13; fi113++ {
                            for fi114 := 0; fi114 < 7; fi114++ {
                                if (!(-1 < 42)) {
                                    if true {
                                        var v115 string = "foo bar"
                                        var v116 string = v3
                                        _ = v115
                                        _ = v116
                                    }
                                } else {
                                    if (false && true) {
                                        v117 := GenS7{F0: 73, F1: 4294967295, F2: (v3 + ((v3 + (v74 + v3)) + ""))}
                                        switch (v74 + ("foo bar" + v74)) {
                                        case "\n\t":
                                            fmt.Println(v1)
                                        case "foo bar":
                                            if false {
                                                v3 = (v3 + (("hello" + (("foo bar" + "hello") + (v3 + v74))) + ""))
                                                v118 := GenS8{F0: 1, F1: fi114}
                                                _ = v118
                                            }
                                        case "hello":
                                            var v119 int8 = 100
                                            _ = v119
                                        default:
                                            v120 := GenS4{F0: -67, F1: 4294967295, F2: genFn0(v74)}
                                            _ = v120
                                        }
                                        v1 -= -289
                                        v74 = "hello"
                                        v1 = (v1 | v1)
                                        delete(v4, (^(fi114)))
                                        if len(v5) > 0 {
                                            v5[0] = (fi113 ^ (fi114 % (((fi114 - fi113)) | 1)))
                                        }
                                        if (0.1 == 1.5) {
                                            if len(v5) > 0 {
                                                v5[0] = fi113
                                            }
                                        }
                                        v1 = 0
                                        _ = v117
                                    } else {
                                        v1 = v1
                                    }
                                    v121 := GenS9{F0: 4294967295}
                                    fmt.Println(v74, v1)
                                    v122 := GenS10{F0: 0, F1: 42}
                                    v74 = v74
                                    _ = v121
                                    _ = v122
                                }
                                var v123 uint16 = 1
                                if true {
                                    if false {
                                        if (genFn0(v3) != 42) {
                                            v124 := GenS5{F0: 2, F1: 4294967295, F2: 2}
                                            v1 += -2147483648
                                            switch (v3 + "") {
                                            case "\n\t":
                                                fmt.Println(fi113, fi114, v1)
                                            }
                                            _ = v124
                                        } else {
                                            if (false || ((18446744073709551615 > 1) && (genFn0(v74) >= 0))) {
                                                _ = len(v4)
                                            } else {
                                                v1 += v1
                                                v125 := GenS4{F0: 127, F1: 2147483648, F2: -9223372036854775808}
                                                if (false || true) {
                                                    var v126 float64 = ((float64(0.0) * 0.0) * 1e10)
                                                    _ = v126
                                                }
                                                if true {
                                                    v127 := v2[v74]
                                                    _ = v127
                                                } else {
                                                    if (!true) {
                                                        v74 = v74
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        var v128 string = v3
                                                        v3 = (("" + ("\n\t" + "foo bar")) + "hello")
                                                        _ = v128
                                                    }
                                                    v3 = ("a" + "a")
                                                    fmt.Println(v1, v123)
                                                }
                                                v74 = ("hello" + v3)
                                                _ = v125
                                            }
                                            switch 1 {
                                            case 32767:
                                                if true {
                                                    switch v74 {
                                                    case "":
                                                        var v129 float64 = 0.0
                                                        _ = v129
                                                    case "foo bar":
                                                        v130 := GenS6{F0: (^(v123)), F1: 42, F2: ((540 == 2) || ((((float32(1.5) + 1.5) / -2.25) / 1.5) == 0.1))}
                                                        _ = v130
                                                    case "hello":
                                                        _ = len(v5)
                                                        fmt.Println(v1)
                                                        fmt.Println(v3, fi114)
                                                    }
                                                    switch ("foo bar" + (("hello" + ((("foo bar" + v3) + v3) + "\n\t")) + ("a" + v74))) {
                                                    case "":
                                                        v4[fi113] = (v3 + "")
                                                    }
                                                    v3 = "a"
                                                    var v131 int64 = 9223372036854775807
                                                    _ = v131
                                                } else {
                                                    var v132 uint8 = 203
                                                    if (v74 == (v3 + "hello")) {
                                                        if true {
                                                            switch v74 {
                                                            case "hello":
                                                                v123 = (v123 | genFn1(2147483647))
                                                            case "a":
                                                                v133 := GenS7{F0: -1, F1: 324, F2: v3}
                                                                var v134 uint8 = v132
                                                                fmt.Println(v123, v74)
                                                                _ = v133
                                                                _ = v134
                                                            }
                                                            if (true && false) {
                                                                switch fi113 {
                                                                case -717:
                                                                    var v135 float32 = (0.0 * float32((-2.25 - 0.0)))
                                                                    v136 := GenS3{F0: 2, F1: (v123 << ((4294967295) & 63))}
                                                                    if len(v5) > 0 {
                                                                        v5[0] = fi114
                                                                    }
                                                                    v135 = v135
                                                                    _ = v135
                                                                    _ = v136
                                                                case 4294967295:
                                                                    fmt.Println(v3)
                                                                default:
                                                                    if true {
                                                                        if false {
                                                                            fmt.Println(v74, fi113, v1)
                                                                            v137 := v4[-9223372036854775808]
                                                                            if true {
                                                                                switch fi113 {
                                                                                case -95:
                                                                                    v137 = ("a" + (((v74 + v137) + (((v137 + ((v137 + v3) + v137)) + "\n\t") + ("" + ((((v74 + ("a" + "a")) + v3) + ((v137 + "\n\t") + "hello")) + "hello")))) + (("" + ("" + ((v137 + ((v74 + "hello") + (v137 + "hello"))) + v137))) + v74)))
                                                                                case -320:
                                                                                    if (!true) {
                                                                                        v123 = (v123 * v123)
                                                                                        v138 := GenS10{F0: 0, F1: 216}
                                                                                        _ = v138
                                                                                    }
                                                                                    if false {
                                                                                        var v139 uint16 = 0
                                                                                        delete(v2, v137)
                                                                                        if false {
                                                                                            v132 = (v132 | v132)
                                                                                            v140 := GenS11{F0: 1, F1: false}
                                                                                            v141 := GenS12{F0: -128}
                                                                                            fmt.Println(v3, v132, v1)
                                                                                            fmt.Println(v74)
                                                                                            _ = v140
                                                                                            _ = v141
                                                                                        } else {
                                                                                            if (-1 < -128) {
                                                                                                v74 = v3
                                                                                                var v142 int8 = -128
                                                                                                _ = v142
                                                                                            } else {
                                                                                                v143 := v2[(("\n\t" + "a") + ("" + v74))]
                                                                                                if true {
                                                                                                    v132 += (v132 / (((v132 >> ((65535) & 63))) | 1))
                                                                                                    v144 := GenS10{F0: 42, F1: v143}
                                                                                                    _ = v144
                                                                                                } else {
                                                                                                    var v145 string = v3
                                                                                                    switch v145 {
                                                                                                    case "":
                                                                                                        v1 += (v1 | (v1 ^ (v1 ^ -336)))
                                                                                                    }
                                                                                                    var v146 int8 = -128
                                                                                                    _ = v145
                                                                                                    _ = v146
                                                                                                }
                                                                                                _ = v143
                                                                                            }
                                                                                            v147 := GenS9{F0: 2147483648}
                                                                                            if ((float32(((-2.25 - 1.5) + 1e10)) * (((3.14159 * 0.0) / 1e10) * 0.0)) < 1.5) {
                                                                                                v148 := GenS5{F0: 68, F1: fi114, F2: v132}
                                                                                                v139 -= uint16(v132)
                                                                                                v149 := GenS13{F0: 1}
                                                                                                _ = v148
                                                                                                _ = v149
                                                                                            } else {
                                                                                                var v150 bool = false
                                                                                                v151 := GenS4{F0: 127, F1: 65535, F2: 692}
                                                                                                v2[v74] = 2147483647
                                                                                                _ = v150
                                                                                                _ = v151
                                                                                            }
                                                                                            _ = v147
                                                                                        }
                                                                                        _ = v139
                                                                                    }
                                                                                    delete(v4, (fi113 ^ fi113))
                                                                                default:
                                                                                    var v152 int32 = (-(v1))
                                                                                    _ = v152
                                                                                }
                                                                                v153 := GenS3{F0: 96, F1: 255}
                                                                                _ = v153
                                                                            }
                                                                            _ = v137
                                                                        }
                                                                        if (false && (!false)) {
                                                                            v154 := GenS2{F0: 1}
                                                                            switch (-(fi114)) {
                                                                            case 344:
                                                                                fmt.Println(v74, fi114, fi113)
                                                                            case -620:
                                                                                v3 = v3
                                                                            case -306:
                                                                                v3 = ("\n\t" + v74)
                                                                            default:
                                                                                if true {
                                                                                    v3 = "hello"
                                                                                    fmt.Println(fi114)
                                                                                }
                                                                            }
                                                                            _ = v154
                                                                        }
                                                                        fmt.Println(v123, v1, v132)
                                                                        if (((^(v132)) >= 125) || false) {
                                                                            var v155 uint16 = (v123 + (v123 / ((uint16(v132)) | 1)))
                                                                            if true {
                                                                                var v156 int64 = 42
                                                                                var v157 bool = (0.1 >= 0.0)
                                                                                if (!v157) {
                                                                                    _ = len(v5)
                                                                                }
                                                                                _ = v156
                                                                                _ = v157
                                                                            }
                                                                            _ = v155
                                                                        }
                                                                        var v158 bool = ((3.14159 == (1e10 - -2.25)) && false)
                                                                        _ = v158
                                                                    } else {
                                                                        v159 := GenS14{F0: 1000}
                                                                        _ = len(v2)
                                                                        v3 = ("a" + ("foo bar" + ""))
                                                                        var v160 int16 = -890
                                                                        if true {
                                                                            if (false || false) {
                                                                                v5 = append(v5, fi114)
                                                                            }
                                                                            switch (("foo bar" + v74) + "foo bar") {
                                                                            case "":
                                                                                fmt.Println(fi114, v74)
                                                                            case "foo bar":
                                                                                var v161 float32 = -2.25
                                                                                _ = v161
                                                                            case "\n\t":
                                                                                v3 = ((v74 + ("\n\t" + "a")) + (v74 + ("\n\t" + ((((v74 + v74) + v74) + ((((v3 + v74) + ("a" + ("foo bar" + ((v3 + v3) + v74)))) + v3) + v3)) + ("\n\t" + ("a" + ("a" + "a")))))))
                                                                            }
                                                                            fmt.Println(fi113)
                                                                            switch "a" {
                                                                            case "\n\t":
                                                                                v123 += v123
                                                                            default:
                                                                                v74 = v3
                                                                            }
                                                                        } else {
                                                                            v1 -= -2147483648
                                                                            var v162 int32 = -1
                                                                            v163 := GenS15{F0: genFn1(2), F1: v132, F2: false}
                                                                            _ = v162
                                                                            _ = v163
                                                                        }
                                                                        _ = v159
                                                                        _ = v160
                                                                    }
                                                                }
                                                                if true {
                                                                    v164 := GenS8{F0: 805, F1: (fi113 - fi114)}
                                                                    fmt.Println(v123)
                                                                    var v165 uint = 2
                                                                    _ = v164
                                                                    _ = v165
                                                                }
                                                                fmt.Println(v1)
                                                            }
                                                            var v166 int16 = 355
                                                            v167 := GenS16{F0: 42, F1: 18446744073709551615}
                                                            v168 := GenS8{F0: 42, F1: fi113}
                                                            if false {
                                                                v1 -= 527
                                                                v123 -= v123
                                                                var v169 bool = (((false || (((2147483648 == 928) || (127 > v132)) && false)) && (4294967295 < 497)) || ((1 <= 737) || (42 >= -1)))
                                                                _ = v169
                                                            } else {
                                                                var v170 uint16 = genFn1(2147483647)
                                                                v171 := GenS13{F0: 42}
                                                                _ = v170
                                                                _ = v171
                                                            }
                                                            _ = len(v5)
                                                            _ = v166
                                                            _ = v167
                                                            _ = v168
                                                        }
                                                    }
                                                    v172 := GenS16{F0: 100, F1: 18446744073709551615}
                                                    var v173 float64 = (0.1 - ((((((1e10 + 1.5) * -2.25) + (3.14159 - 1.5)) * 0.1) * 1e10) / 1.5))
                                                    v123 += uint16(v132)
                                                    v174 := GenS12{F0: -128}
                                                    _ = v132
                                                    _ = v172
                                                    _ = v173
                                                    _ = v174
                                                }
                                            case -733:
                                                var v175 uint = 18446744073709551615
                                                _ = v175
                                            case -929:
                                                v74 = (v3 + ((v74 + (v74 + "hello")) + "\n\t"))
                                            default:
                                                switch (-(fi113)) {
                                                case 1:
                                                    switch (fi114 * (fi113 >> ((453) & 63))) {
                                                    case 9223372036854775807:
                                                        if (2147483647 <= 65535) {
                                                            if false {
                                                                v176 := v4[(fi113 << ((42) & 63))]
                                                                _ = v176
                                                            }
                                                        }
                                                    case 128:
                                                        v74 = ""
                                                        if ("" < "hello") {
                                                            switch v3 {
                                                            case "":
                                                                fmt.Println(v123, v123)
                                                            case "\n\t":
                                                                switch -623 {
                                                                case 32767:
                                                                    v123 += v123
                                                                case 32768:
                                                                    v177 := GenS1{F0: 255, F1: 667, F2: 683}
                                                                    _ = v177
                                                                default:
                                                                    v3 = "\n\t"
                                                                }
                                                            }
                                                        }
                                                        _ = len(v2)
                                                    default:
                                                        var v178 int16 = -1
                                                        _ = v178
                                                    }
                                                case 9223372036854775807:
                                                    v2[("foo bar" + ("" + "a"))] = 42
                                                }
                                            }
                                            var v179 int16 = 1000
                                            _ = v179
                                        }
                                    }
                                    switch v3 {
                                    case "hello":
                                        var v180 uint32 = 2147483647
                                        _ = v180
                                    case "a":
                                        if (v74 <= ("\n\t" + "hello")) {
                                            var v181 float32 = (0.0 - ((float32((0.1 / 3.14159)) + float32((1e10 + ((float64((0.0 + (3.14159 - 1.5))) - (0.0 + 1e10)) - (0.1 / (((-2.25 + 1e10) * (-2.25 / 0.1)) - -2.25)))))) / ((0.1 + 1e10) / 3.14159)))
                                            _ = v181
                                        }
                                        v123 += (v123 - (v123 * 332))
                                        fmt.Println(v123)
                                        v1 = v1
                                        delete(v2, v74)
                                        v182 := GenS17{F0: true, F1: 806, F2: "a"}
                                        _ = v182
                                    case "foo bar":
                                        v183 := GenS1{F0: 42, F1: 2, F2: 4294967295}
                                        _ = v183
                                    }
                                    fmt.Println(v3)
                                } else {
                                    if (!true) {
                                        v3 = (v3 + v3)
                                    } else {
                                        switch v74 {
                                        case "foo bar":
                                            v1 = v1
                                        case "hello":
                                            v3 = v74
                                        default:
                                            v123 = (v123 - 255)
                                        }
                                    }
                                    var v184 float32 = 0.0
                                    if (v1 > v1) {
                                        v185 := GenS7{F0: 100, F1: 595, F2: v3}
                                        v5 = append(v5, -954)
                                        var v186 uint32 = 0
                                        _ = v185
                                        _ = v186
                                    }
                                    _ = v184
                                }
                                _ = fi114
                                _ = v123
                            }
                            fmt.Println(v74)
                            _ = fi113
                        }
                    }
                } else {
                    var v187 uint64 = 42
                    fmt.Println(v3)
                    v2[(((("hello" + "foo bar") + v3) + v74) + v74)] = 132
                    _ = v187
                }
            } else {
                v3 = ((("foo bar" + v74) + "\n\t") + v74)
            }
            v74 = v74
            fmt.Println(v3)
        }
        var v188 float64 = float64((1e10 + ((float32(1.5) - (0.1 + 3.14159)) / (1e10 - 3.14159))))
        _ = v109
        _ = v188
    }
    switch -2147483648 {
    case 4294967295:
        v1 -= 2147483647
    default:
        if true {
            for fi189 := 0; fi189 < 12; fi189++ {
                v3 = ("hello" + v3)
                v190 := GenS18{F0: 211}
                var v191 int64 = -1
                _ = fi189
                _ = v190
                _ = v191
            }
            v192 := GenS19{F0: 2, F1: 811}
            _ = v192
        }
    }
    v74 = "hello"
    switch 2147483647 {
    case 288:
        switch "hello" {
        case "\n\t":
            for fi193 := 0; fi193 < 11; fi193++ {
                var v194 float32 = 0.0
                _ = fi193
                _ = v194
            }
        case "foo bar":
            switch -255 {
            case -1:
                for fi195 := 0; fi195 < 9; fi195++ {
                    v74 = "a"
                    _ = fi195
                }
            case 42:
                fmt.Println(v1, v3)
            case -537:
                v3 = v3
            }
        }
    case 158:
        v196 := GenS20{F0: 0, F1: 127}
        _ = v196
    }
    for fi197 := 0; fi197 < 18; fi197++ {
        if true {
            for fi198 := 0; fi198 < 14; fi198++ {
                fmt.Println(v3, v74)
                _ = fi198
            }
            var v199 uint16 = 2
            var v200 uint8 = 255
            _ = v199
            _ = v200
        } else {
            if (-9223372036854775808 == 42) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        }
        _ = fi197
    }
    for fi201 := 0; fi201 < 10; fi201++ {
        fmt.Println(v3, v1)
        v1 = (-(v1))
        if true {
            if true {
                fmt.Println(v1, fi201, v3)
                v74 = ""
            } else {
                switch "\n\t" {
                case "hello":
                    v202 := GenS21{F0: (^(fi201))}
                    _ = v202
                }
            }
        } else {
            switch v3 {
            case "a":
                if true {
                    v74 = (v74 + ((("" + v74) + "a") + ((("\n\t" + v3) + v3) + "\n\t")))
                    if (0.1 != 0.1) {
                        v203 := GenS18{F0: 1}
                        v1 = 1
                        v1 -= (v1 - (v1 & v1))
                        _ = v203
                    } else {
                        fmt.Println(v1)
                    }
                } else {
                    var v204 uint16 = 0
                    switch fi201 {
                    case 393:
                        v74 = ("" + v3)
                    case -9223372036854775808:
                        if ((((0.1 / 3.14159) * 1e10) - ((float64(3.14159) / (1e10 - (0.0 * (0.1 / 0.1)))) + 0.1)) != 0.1) {
                            v74 = ""
                            var v205 string = (("a" + "foo bar") + v74)
                            var v206 uint8 = 127
                            _ = v205
                            _ = v206
                        } else {
                            if true {
                                _ = len(v2)
                                v204 = v204
                                v207 := GenS18{F0: 0}
                                _ = v207
                            } else {
                                if true {
                                    v3 = ""
                                }
                            }
                            for fi208 := 0; fi208 < 18; fi208++ {
                                var v209 int = (fi201 & fi201)
                                fmt.Println(v74)
                                if len(v5) > 0 {
                                    v5[0] = 32767
                                }
                                _ = fi208
                                _ = v209
                            }
                        }
                    }
                    if (753 <= v1) {
                        if true {
                            delete(v4, (fi201 * fi201))
                            v3 = v74
                        } else {
                            fmt.Println(v74, v204, v74)
                            fmt.Println(v74)
                        }
                        for fi210 := 0; fi210 < 6; fi210++ {
                            var v211 string = ((v3 + v3) + v3)
                            fmt.Println(fi210)
                            switch 127 {
                            case -1:
                                fmt.Println(fi210)
                            case 127:
                                if (-180 <= 42) {
                                    _ = len(v5)
                                    if (237 < 42) {
                                        v212 := GenS10{F0: -9223372036854775808, F1: 65535}
                                        delete(v2, "foo bar")
                                        _ = v212
                                    } else {
                                        delete(v2, "foo bar")
                                    }
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                } else {
                                    fmt.Println(fi210, fi210)
                                }
                            case 2147483647:
                                switch "a" {
                                case "foo bar":
                                    var v213 int = fi201
                                    _ = v213
                                case "\n\t":
                                    if len(v5) > 0 {
                                        v5[0] = fi210
                                    }
                                case "a":
                                    _ = len(v5)
                                }
                            }
                            _ = fi210
                            _ = v211
                        }
                        var v214 uint16 = v204
                        _ = v214
                    }
                    _ = v204
                }
            case "foo bar":
                for fi215 := 0; fi215 < 9; fi215++ {
                    if true {
                        var v216 int64 = -1
                        v216 = (v216 | 42)
                        _ = v216
                    }
                    v217 := GenS4{F0: 1, F1: 2, F2: -9223372036854775808}
                    v74 = v3
                    _ = fi215
                    _ = v217
                }
            case "hello":
                for fi218 := 0; fi218 < 19; fi218++ {
                    v74 = ""
                    fmt.Println(fi201)
                    _ = fi218
                }
            }
            fmt.Println(v1, v74)
            v219 := GenS1{F0: 785, F1: 2, F2: 2}
            _ = v219
        }
        _ = fi201
    }
    for fi220 := 0; fi220 < 14; fi220++ {
        v221 := GenS10{F0: 854, F1: 65535}
        _ = fi220
        _ = v221
    }
    var v222 uint = 79
    if (9 > 255) {
        var v223 int = 128
        v224 := GenS10{F0: -9223372036854775808, F1: 0}
        _ = v223
        _ = v224
    } else {
        v222 += (v222 - (v222 % ((760) | 1)))
        fmt.Println(v3)
        fmt.Println(v222, v74, v3)
    }
    v222 += (v222 - v222)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v222:", v222)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v222
    _ = v3
    _ = v4
    _ = v5
    _ = v74
}
