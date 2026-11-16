package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 string; F1 uint64; F2 int32 }
type GenS6 struct { F0 uint64; F1 uint32 }
type GenS7 struct { F0 uint8; F1 int16; F2 int64 }
type GenS8 struct { F0 int64; F1 bool; F2 int8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS11 struct { F0 int16; F1 int64 }
type GenS12 struct { F0 int; F1 uint8; F2 uint8 }
type GenS13 struct { F0 int8; F1 uint64; F2 uint }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int16; F1 uint8 }
type GenS16 struct { F0 int64; F1 uint8 }
type GenS17 struct { F0 int64; F1 uint; F2 int32 }
type GenS18 struct { F0 uint16; F1 bool }
func genFn0(p0 uint64, p1 uint) uint64 {
    var v6 uint64 = 0
    if false {
        v7 := GenS1{F0: 0}
        v8 := []int{-11, -9223372036854775808, 4294967295}
        _ = v7
        _ = v8
    }
    var v9 int8 = -128
    if (491 != 242) {
        v6 += (p0 + (p0 << ((p1) & 63)))
        v10 := []int{-331, 4294967295, 2147483647, 9223372036854775807}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v9
    return p0
}
func genFn1(p0 int8) int32 {
    for fi11 := 0; fi11 < 8; fi11++ {
        v12 := GenS1{F0: 42}
        if true {
            fmt.Println(fi11)
            for fi13 := 0; fi13 < 2; fi13++ {
                v14 := GenS1{F0: 42}
                fmt.Println(p0, p0, fi11)
                _ = fi13
                _ = v14
            }
        } else {
            var v15 int8 = p0
            fmt.Println(fi11)
            switch "foo bar" {
            case "":
                if ((!false) && ((((float32(1.5) + (float32(-2.25) / 1.0)) - float32((3.14159 + 0.1))) < 1.5) || ("a" == ""))) {
                    for fi16 := 0; fi16 < 16; fi16++ {
                        v15 = v15
                        switch (^(fi16)) {
                        case 65:
                            if (!(-2147483648 < int32(fi16))) {
                                switch (fi16 / (((^(fi16))) | 1)) {
                                case 580:
                                    fmt.Println(p0, fi16, fi11)
                                case 32768:
                                    v17 := []int{}
                                    _ = v17
                                }
                            }
                        case 4294967295:
                            switch 32767 {
                            case 32768:
                                if false {
                                    p0 = v15
                                    v18 := GenS2{F0: "foo bar", F1: 0}
                                    v19 := []int{(fi11 * fi11), 32768, -9223372036854775808}
                                    _ = v18
                                    _ = v19
                                }
                            case 4294967295:
                                v20 := GenS3{F0: false}
                                _ = v20
                            case 32:
                                switch "foo bar" {
                                case "foo bar":
                                    var v21 float64 = 1.5
                                    _ = v21
                                case "\n\t":
                                    v22 := GenS3{F0: ((false || (true && true)) || true)}
                                    _ = v22
                                default:
                                    v23 := GenS3{F0: true}
                                    _ = v23
                                }
                                p0 = v15
                            default:
                                v15 = v15
                            }
                        case 42:
                            var v24 float32 = (-2.25 + -2.25)
                            var v25 bool = (!false)
                            _ = v24
                            _ = v25
                        default:
                            v26 := []int{(fi11 / ((fi16) | 1)), (fi11 | (fi16 & -9223372036854775808)), fi16, (fi16 << ((255) & 63))}
                            _ = v26
                        }
                        _ = fi16
                    }
                } else {
                    v15 = v15
                    if (621 == 916) {
                        var v27 bool = true
                        if (0.0 < -2.25) {
                            if v27 {
                                var v28 int16 = 470
                                v29 := GenS4{F0: -1, F1: fi11}
                                v30 := GenS4{F0: -32768, F1: 127}
                                _ = v28
                                _ = v29
                                _ = v30
                            }
                            for fi31 := 0; fi31 < 19; fi31++ {
                                p0 = 100
                                if (!false) {
                                    fmt.Println(fi11, fi11)
                                } else {
                                    switch ("foo bar" + ("\n\t" + "\n\t")) {
                                    case "\n\t":
                                        v15 += (-(v15))
                                    case "foo bar":
                                        v32 := []int{}
                                        _ = v32
                                    default:
                                        switch (((("foo bar" + "a") + "hello") + "a") + "a") {
                                        case "hello":
                                            var v33 uint = 109
                                            _ = v33
                                        case "\n\t":
                                            fmt.Println(fi31, fi31, p0)
                                            if false {
                                                switch ("" + ((("a" + "\n\t") + "hello") + "foo bar")) {
                                                case "foo bar":
                                                    switch "a" {
                                                    case "a":
                                                        var v34 uint8 = 128
                                                        _ = v34
                                                    case "":
                                                        v35 := []int{(fi31 | 127)}
                                                        _ = v35
                                                    case "hello":
                                                        var v36 uint64 = 42
                                                        _ = v36
                                                    default:
                                                        var v37 int64 = 180
                                                        _ = v37
                                                    }
                                                case "\n\t":
                                                    v38 := []int{(fi31 << ((18446744073709551615) & 63)), fi11, 9223372036854775807, fi11}
                                                    _ = v38
                                                case "":
                                                    p0 += (v15 / (((v15 << ((255) & 63))) | 1))
                                                }
                                                v39 := GenS3{F0: v27}
                                                _ = v39
                                            }
                                            v40 := []int{(-(fi11)), 32768, (fi11 / ((fi31) | 1))}
                                            _ = v40
                                        case "a":
                                            v41 := []int{}
                                            _ = v41
                                        default:
                                            var v42 int16 = -796
                                            _ = v42
                                        }
                                        v27 = false
                                        v43 := []int{}
                                        _ = v43
                                    }
                                }
                                _ = fi31
                            }
                            switch -708 {
                            case -991:
                                var v44 int64 = 0
                                _ = v44
                            case 127:
                                for fi45 := 0; fi45 < 1; fi45++ {
                                    v46 := GenS4{F0: -449, F1: (fi45 + (fi11 / (((fi11 * fi11)) | 1)))}
                                    var v47 uint32 = 1
                                    _ = fi45
                                    _ = v46
                                    _ = v47
                                }
                            }
                        }
                        _ = v27
                    } else {
                        if true {
                            if (42 == 371) {
                                var v48 uint32 = 2147483647
                                switch fi11 {
                                case 79:
                                    v48 -= v48
                                }
                                v49 := GenS4{F0: -497, F1: (^(fi11))}
                                _ = v48
                                _ = v49
                            } else {
                                for fi50 := 0; fi50 < 4; fi50++ {
                                    var v51 uint8 = 255
                                    _ = fi50
                                    _ = v51
                                }
                            }
                            fmt.Println(v15, p0)
                            if (!false) {
                                fmt.Println(fi11, v15)
                                var v52 int32 = 715
                                _ = v52
                            } else {
                                v53 := GenS5{F0: ("" + "foo bar"), F1: 2, F2: -991}
                                _ = v53
                            }
                        }
                        v54 := GenS6{F0: 770, F1: 4294967295}
                        v55 := GenS7{F0: 118, F1: 1, F2: 1}
                        _ = v54
                        _ = v55
                    }
                    var v56 int8 = (p0 - (p0 + p0))
                    _ = v56
                }
            default:
                switch -443 {
                case -1:
                    for fi57 := 0; fi57 < 16; fi57++ {
                        switch "\n\t" {
                        case "foo bar":
                            var v58 uint16 = 32767
                            _ = v58
                        default:
                            switch "a" {
                            case "\n\t":
                                fmt.Println(fi11)
                            case "foo bar":
                                var v59 int64 = 42
                                _ = v59
                            default:
                                if (fi57 < fi11) {
                                    var v60 uint16 = 32767
                                    switch fi57 {
                                    case 961:
                                        v61 := GenS3{F0: (!(fi11 <= 127))}
                                        _ = v61
                                    }
                                    _ = v60
                                } else {
                                    p0 = (v15 - (-(v15)))
                                    var v62 int16 = -937
                                    fmt.Println(v62, fi11, v15)
                                    _ = v62
                                }
                            }
                            v63 := GenS8{F0: 9223372036854775807, F1: false, F2: (v15 ^ p0)}
                            _ = v63
                        }
                        if (4294967295 <= 0) {
                            switch (("\n\t" + "a") + ("foo bar" + "\n\t")) {
                            case "":
                                var v64 float64 = 1.5
                                if false {
                                    var v65 float32 = (0.0 + 0.1)
                                    _ = v65
                                } else {
                                    p0 = (^(v15))
                                    if ((42 == 32767) || (false || (("hello" + (("\n\t" + ("hello" + (("a" + "a") + ("" + "foo bar")))) + "")) < "a"))) {
                                        var v66 int8 = v15
                                        v66 = 76
                                        _ = v66
                                    } else {
                                        if true {
                                            v15 += (v15 + (^(v15)))
                                        }
                                        v64 += (((-2.25 - (float64(0.1) * v64)) * 3.14159) * float64((((0.1 / 1.5) * 0.1) * (((3.14159 * ((-2.25 - 0.1) / 1.0)) + float32(float64(1.5))) + 0.0))))
                                        var v67 int32 = 900
                                        v67 -= v67
                                        _ = v67
                                    }
                                    fmt.Println(fi11, fi57, v64)
                                }
                                _ = v64
                            default:
                                switch 32768 {
                                case 235:
                                    v68 := []int{fi57, -245}
                                    _ = v68
                                }
                            }
                            if true {
                                var v69 uint32 = 2147483647
                                _ = v69
                            }
                            p0 -= (p0 << ((42) & 63))
                        } else {
                            fmt.Println(fi11, fi57, v15)
                            v70 := GenS8{F0: 0, F1: (-2.25 != float64(3.14159)), F2: p0}
                            var v71 int = fi11
                            _ = v70
                            _ = v71
                        }
                        _ = fi57
                    }
                case 0:
                    v15 = 127
                }
            }
            _ = v15
        }
        var v72 uint32 = 65535
        _ = fi11
        _ = v12
        _ = v72
    }
    v73 := GenS8{F0: 0, F1: (p0 >= p0), F2: p0}
    if ((1 <= 2) && true) {
        var v74 uint16 = 693
        v75 := GenS6{F0: 4294967295, F1: 2147483647}
        _ = v74
        _ = v75
    } else {
        var v76 bool = false
        v76 = false
        _ = v76
    }
    var v77 float64 = 3.14159
    v78 := []int{32768, 4294967295}
    _ = p0
    _ = v73
    _ = v77
    _ = v78
    return -33
}

func main() {
    var v1 bool = true
    v2 := []int{}
    var v3 int8 = 91
    var v4 float64 = 0.0
    var v5 float64 = ((float64(3.14159) / (float64(0.0) * v4)) + 1e10)
    for fi79 := 0; fi79 < 19; fi79++ {
        if ((((("foo bar" + (("" + "") + "hello")) + "a") + "\n\t") < ("\n\t" + "\n\t")) || (v1 || v1)) {
            if false {
                v4 -= v5
                switch (fi79 ^ fi79) {
                case 127:
                    var v80 uint64 = 864
                    _ = v80
                case 0:
                    for fi81 := 0; fi81 < 7; fi81++ {
                        var v82 uint = 65535
                        var v83 int16 = 1
                        _ = fi81
                        _ = v82
                        _ = v83
                    }
                default:
                    fmt.Println(v4)
                }
                v3 = (v3 * (^(v3)))
            } else {
                var v84 uint16 = 204
                _ = v84
            }
            for fi85 := 0; fi85 < 15; fi85++ {
                if v1 {
                    switch "foo bar" {
                    case "foo bar":
                        if false {
                            var v86 float32 = 1.5
                            _ = v86
                        } else {
                            fmt.Println(v3, v3)
                        }
                    case "\n\t":
                        _ = len(v2)
                    case "a":
                        var v87 float32 = 1.5
                        if len(v2) > 0 {
                            v2[0] = fi79
                        }
                        v4 -= (((v5 + v4) / 1.0) / (1e10 / v4))
                        var v88 bool = (993 != 277)
                        var v89 uint32 = 516
                        _ = v87
                        _ = v88
                        _ = v89
                    }
                    var v90 int16 = 32767
                    v91 := GenS9{F0: 2}
                    _ = v90
                    _ = v91
                } else {
                    v92 := GenS10{F0: 1, F1: 592, F2: 715}
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v92
                }
                _ = fi85
            }
            var v93 int64 = 918
            _ = v93
        } else {
            if v1 {
                v5 = 1e10
                if false {
                    v94 := GenS3{F0: true}
                    _ = v94
                }
                var v95 int8 = v3
                _ = v95
            } else {
                v96 := GenS10{F0: 128, F1: 32767, F2: 159}
                _ = v96
            }
            var v97 int16 = 623
            _ = v97
        }
        for fi98 := 0; fi98 < 17; fi98++ {
            if len(v2) > 0 {
                v2[0] = (fi98 << ((0) & 63))
            }
            v99 := GenS5{F0: "a", F1: 4294967295, F2: 869}
            if (!false) {
                switch "a" {
                case "":
                    if (v1 && (!false)) {
                        fmt.Println(fi79, fi79)
                    } else {
                        var v100 uint8 = 168
                        _ = v100
                    }
                case "hello":
                    var v101 uint = 83
                    _ = v101
                default:
                    if ((-(v3)) <= v3) {
                        if len(v2) > 0 {
                            v2[0] = 42
                        }
                    } else {
                        switch "hello" {
                        case "foo bar":
                            switch fi79 {
                            case 1:
                                if (v1 && true) {
                                    var v102 int16 = 1000
                                    v1 = (32767 == v102)
                                    switch "\n\t" {
                                    case "a":
                                        v103 := GenS8{F0: 1, F1: v1, F2: 91}
                                        _ = v103
                                    case "hello":
                                        v1 = v1
                                    case "\n\t":
                                        v104 := GenS5{F0: (((("hello" + "\n\t") + "hello") + "foo bar") + "a"), F1: 1, F2: genFn1((v3 << ((42) & 63)))}
                                        _ = v104
                                    }
                                    v4 += (v5 * (((v4 - 3.14159) + (v5 + v5)) - (v4 / v4)))
                                    _ = v102
                                }
                            case -1:
                                switch (fi79 - (fi98 | (-(fi79)))) {
                                case -25:
                                    v3 = (v3 & (v3 ^ v3))
                                case -571:
                                    fmt.Println(v5, v3)
                                default:
                                    var v105 int8 = 100
                                    _ = v105
                                }
                            case 4294967295:
                                if (790 == 42) {
                                    var v106 int = 128
                                    var v107 int = fi79
                                    _ = v106
                                    _ = v107
                                } else {
                                    switch ("" + "") {
                                    case "a":
                                        _ = len(v2)
                                    default:
                                        v2 = append(v2, (fi79 + (^(fi79))))
                                    }
                                    if v1 {
                                        if false {
                                            v3 += (v3 - v3)
                                            var v108 uint32 = 2147483647
                                            var v109 int32 = 2147483647
                                            _ = v108
                                            _ = v109
                                        }
                                        var v110 uint8 = 128
                                        fmt.Println(v5, fi98, v5)
                                        _ = v110
                                    } else {
                                        v4 -= 1.5
                                        v111 := GenS11{F0: 42, F1: 42}
                                        v5 = v4
                                        _ = v111
                                    }
                                    _ = len(v2)
                                }
                            }
                            fmt.Println(v5, fi98)
                            var v112 int32 = 0
                            _ = v112
                        case "\n\t":
                            v3 = v3
                        case "":
                            v113 := GenS2{F0: "hello", F1: (v3 & v3)}
                            _ = v113
                        default:
                            var v114 int16 = -32768
                            _ = v114
                        }
                        if len(v2) > 0 {
                            v2[0] = 1
                        }
                        v115 := GenS12{F0: (fi98 & 0), F1: 0, F2: 42}
                        v116 := GenS13{F0: (v3 * v3), F1: 2, F2: 42}
                        _ = v115
                        _ = v116
                    }
                }
            }
            if true {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if v1 {
                    v3 -= (v3 % ((48) | 1))
                    v5 = (float64(1e10) / (v5 - (v4 - v4)))
                    v4 = ((v4 / (v5 - 0.1)) / float64((0.1 + 1e10)))
                }
                fmt.Println(v1)
                v3 += v3
            } else {
                switch (fi79 ^ (fi79 / ((-9223372036854775808) | 1))) {
                case -9223372036854775808:
                    v5 += v5
                case 183:
                    var v117 uint64 = 29
                    _ = v117
                case 458:
                    var v118 int8 = v3
                    _ = v118
                default:
                    v2 = append(v2, fi79)
                }
                if true {
                    var v119 uint8 = 42
                    v4 = v4
                    var v120 float32 = float32(v5)
                    _ = v119
                    _ = v120
                } else {
                    if len(v2) > 0 {
                        v2[0] = 32768
                    }
                }
                v121 := GenS11{F0: 42, F1: -733}
                v1 = v1
                v5 -= ((1.5 / ((v4 / v4) / (0.1 + ((((v4 / (((v5 + float64((1e10 * (1e10 + 1e10)))) / v4) / float64(float32((0.1 - v4))))) + v5) * v5) / v5)))) + v4)
                _ = v121
            }
            var v122 int = fi79
            var v123 bool = v1
            _ = fi98
            _ = v122
            _ = v123
            _ = v99
        }
        fmt.Println(v4, v3)
        _ = fi79
    }
    fmt.Println(v5)
    if v1 {
        var v124 int64 = -143
        _ = v124
    } else {
        v125 := GenS4{F0: 234, F1: -9223372036854775808}
        var v126 uint16 = 264
        _ = v125
        _ = v126
    }
    for fi127 := 0; fi127 < 19; fi127++ {
        fmt.Println(v4, v1)
        fmt.Println(v4, v1)
        for fi128 := 0; fi128 < 16; fi128++ {
            v129 := GenS8{F0: -1, F1: v1, F2: (v3 * 42)}
            switch ("foo bar" + "hello") {
            case "":
                v3 += v3
                var v130 int = (fi128 & fi128)
                if false {
                    var v131 uint16 = 991
                    v131 += v131
                    var v132 float64 = ((-2.25 + ((v4 + v5) * v4)) - v4)
                    _ = v131
                    _ = v132
                } else {
                    if (-850 != 0) {
                        var v133 float64 = (1.5 + v5)
                        fmt.Println(fi128, v130)
                        v3 = -1
                        v133 -= float64(-2.25)
                        _ = v133
                    } else {
                        if v1 {
                            v134 := GenS9{F0: 18446744073709551615}
                            _ = v134
                        }
                        v3 = (v3 | (v3 / ((42) | 1)))
                        fmt.Println(v130, fi127)
                    }
                }
                _ = v130
            default:
                if v1 {
                    if ("\n\t" > ("foo bar" + "a")) {
                        if len(v2) > 0 {
                            v2[0] = fi128
                        }
                        if (!v1) {
                            fmt.Println(v4)
                        }
                        var v135 int32 = -2147483648
                        _ = v135
                    }
                    var v136 float64 = v4
                    _ = v136
                } else {
                    v4 = v4
                }
            }
            switch "hello" {
            case "foo bar":
                fmt.Println(v1, v1, v3)
            }
            _ = fi128
            _ = v129
        }
        _ = fi127
    }
    for fi137 := 0; fi137 < 19; fi137++ {
        if v1 {
            v4 -= ((3.14159 - v5) * (v4 - (float64(float32((3.14159 + -2.25))) + (float64(((float32(1.5) / (3.14159 * (((1.5 * 3.14159) + 1e10) + float32(v4)))) + (-2.25 - (0.1 + 3.14159)))) * v4))))
            v2 = append(v2, fi137)
            v2 = append(v2, 42)
        } else {
            switch (fi137 << ((1) & 63)) {
            case 32767:
                v138 := GenS6{F0: 4294967295, F1: 1}
                _ = v138
            default:
                switch "hello" {
                case "foo bar":
                    for fi139 := 0; fi139 < 15; fi139++ {
                        if v1 {
                            switch "foo bar" {
                            case "":
                                _ = len(v2)
                            case "a":
                                if true {
                                    v4 = (v5 * (v4 * v5))
                                    if (false && v1) {
                                        v140 := GenS10{F0: 6, F1: -1, F2: 754}
                                        var v141 int16 = 0
                                        _ = v140
                                        _ = v141
                                    } else {
                                        switch "hello" {
                                        case "hello":
                                            switch -764 {
                                            case 2147483647:
                                                var v142 float64 = ((v5 / float64(3.14159)) / 1.5)
                                                _ = v142
                                            case -306:
                                                v2 = append(v2, (^(fi137)))
                                                var v143 int32 = -25
                                                var v144 uint32 = 65535
                                                _ = v143
                                                _ = v144
                                            case -211:
                                                var v145 int32 = -427
                                                _ = v145
                                            }
                                        case "":
                                            var v146 int8 = (v3 | v3)
                                            _ = v146
                                        default:
                                            v4 = ((v4 - v5) + v4)
                                            v147 := GenS13{F0: v3, F1: 475, F2: 4294967295}
                                            v2 = append(v2, fi137)
                                            var v148 int = fi137
                                            _ = v147
                                            _ = v148
                                        }
                                    }
                                    v2 = append(v2, (fi139 - (fi139 / ((fi139) | 1))))
                                } else {
                                    fmt.Println(v5, v1, fi137)
                                    v2 = append(v2, fi139)
                                }
                            default:
                                var v149 int32 = 1
                                _ = v149
                            }
                            var v150 bool = (412 <= 2)
                            _ = len(v2)
                            _ = v150
                        } else {
                            v1 = false
                        }
                        v5 = v5
                        _ = fi139
                    }
                case "a":
                    var v151 float32 = 3.14159
                    _ = v151
                default:
                    switch "" {
                    case "\n\t":
                        for fi152 := 0; fi152 < 9; fi152++ {
                            var v153 uint16 = 591
                            v3 -= v3
                            if ((-2.25 / float32((float64((0.0 * ((1.5 / 3.14159) - -2.25))) / (0.1 * (v5 / 3.14159))))) >= float32((v4 / v5))) {
                                if v1 {
                                    var v154 int32 = -1
                                    v155 := GenS5{F0: (("" + (((("hello" + "hello") + ("hello" + "a")) + "") + (("a" + "\n\t") + ""))) + ""), F1: 74, F2: int32(v3)}
                                    v156 := GenS7{F0: 128, F1: 1000, F2: -900}
                                    _ = v154
                                    _ = v155
                                    _ = v156
                                } else {
                                    v157 := GenS14{F0: ("foo bar" + "a")}
                                    var v158 float32 = (1e10 / 0.1)
                                    v2 = append(v2, (^(fi152)))
                                    v158 = 1.5
                                    if v1 {
                                        v159 := GenS15{F0: -590, F1: 183}
                                        if v1 {
                                            fmt.Println(v4, v3)
                                        }
                                        _ = v159
                                    }
                                    var v160 int16 = 26
                                    v5 = v4
                                    _ = v157
                                    _ = v158
                                    _ = v160
                                }
                                var v161 string = "\n\t"
                                var v162 int = (fi137 << ((18446744073709551615) & 63))
                                _ = v161
                                _ = v162
                            }
                            _ = fi152
                            _ = v153
                        }
                    case "a":
                        v2 = append(v2, -796)
                    default:
                        switch ("a" + "a") {
                        case "foo bar":
                            v163 := GenS16{F0: 0, F1: 128}
                            _ = v163
                        case "\n\t":
                            for fi164 := 0; fi164 < 3; fi164++ {
                                var v165 float64 = ((v4 - ((0.1 / (v4 - (v5 / v5))) * float64(0.1))) / v5)
                                fmt.Println(fi164, v4, v5)
                                fmt.Println(v1, v165, fi164)
                                _ = fi164
                                _ = v165
                            }
                        }
                    }
                }
            }
            if false {
                for fi166 := 0; fi166 < 2; fi166++ {
                    var v167 int = -9223372036854775808
                    var v168 int64 = -735
                    v169 := GenS12{F0: 0, F1: 0, F2: 25}
                    var v170 int8 = 68
                    v171 := GenS17{F0: (v168 - (v168 & (^(v168)))), F1: 255, F2: genFn1(100)}
                    _ = fi166
                    _ = v167
                    _ = v168
                    _ = v169
                    _ = v170
                    _ = v171
                }
                var v172 int8 = v3
                _ = v172
            } else {
                v173 := GenS12{F0: fi137, F1: 127, F2: 232}
                _ = v173
            }
            for fi174 := 0; fi174 < 18; fi174++ {
                if len(v2) > 0 {
                    v2[0] = -974
                }
                _ = fi174
            }
        }
        if (2 != 255) {
            v175 := GenS7{F0: 42, F1: -634, F2: 9223372036854775807}
            _ = v175
        }
        _ = fi137
    }
    fmt.Println(v1, v1)
    v176 := GenS3{F0: true}
    var v177 int32 = 0
    v177 -= v177
    v178 := GenS7{F0: 74, F1: -929, F2: -73}
    _ = len(v2)
    if (0.0 < ((0.0 - 1.5) + -2.25)) {
        v4 = float64(3.14159)
    } else {
        if (18446744073709551615 > 1) {
            v177 = v177
        }
        v179 := GenS18{F0: 373, F1: (-2147483648 >= -9223372036854775808)}
        _ = v179
    }
    v1 = (65535 > 158)
    fmt.Println("v1:", v1)
    fmt.Println("v177:", v177)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v176
    _ = v177
    _ = v178
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
