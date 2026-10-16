package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 int32; F1 int8; F2 uint32 }
type GenS4 struct { F0 int8; F1 int; F2 int16 }
type GenS5 struct { F0 int16; F1 uint8 }
type GenS6 struct { F0 bool; F1 uint }
type GenS7 struct { F0 int32; F1 int8 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 string; F1 uint8; F2 uint16 }
type GenS10 struct { F0 uint16; F1 uint8 }
type GenS11 struct { F0 uint; F1 int16 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint32; F1 uint; F2 uint32 }
type GenS14 struct { F0 int32; F1 int8; F2 uint32 }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 string; F1 uint64; F2 int16 }
type GenS18 struct { F0 int; F1 string; F2 uint16 }
type GenS19 struct { F0 uint8 }
func genFn0(p0 uint32, p1 string, p2 uint8) int32 {
    v4 := []int{}
    if (!(-2147483648 >= 42)) {
        v4 = append(v4, 114)
        _ = len(v4)
    } else {
        for fi5 := 0; fi5 < 12; fi5++ {
            if (!false) {
                fmt.Println(p1)
            }
            fmt.Println(p0, p2, fi5)
            _ = fi5
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 42
}
func genFn1(p0 uint16, p1 uint16) uint8 {
    var v6 bool = true
    if (((^(p0)) > p0) && v6) {
        for fi7 := 0; fi7 < 15; fi7++ {
            switch (fi7 << ((255) & 63)) {
            case 32767:
                v8 := []int{fi7}
                _ = v8
            case 0:
                for fi9 := 0; fi9 < 13; fi9++ {
                    switch "" {
                    case "hello":
                        var v10 int16 = 0
                        _ = v10
                    case "":
                        p1 = 65535
                    case "\n\t":
                        if false {
                            p1 -= 0
                            var v11 uint32 = 0
                            _ = v11
                        }
                    default:
                        if (fi9 <= (fi9 + 228)) {
                            if (-2147483648 <= -2147483648) {
                                var v12 uint64 = 100
                                p0 -= p0
                                _ = v12
                            } else {
                                v13 := GenS1{F0: -128, F1: 4294967295}
                                v14 := GenS1{F0: 127, F1: 4294967295}
                                _ = v13
                                _ = v14
                            }
                            switch -681 {
                            case 705:
                                switch fi7 {
                                case 1:
                                    fmt.Println(fi9, v6)
                                default:
                                    fmt.Println(fi7, p0, fi7)
                                }
                            default:
                                v15 := []int{-157, fi9, -556}
                                _ = v15
                            }
                        }
                    }
                    v6 = true
                    _ = fi9
                }
            default:
                if v6 {
                    if v6 {
                        p1 += p1
                        p0 = 32767
                        p0 = 1
                    }
                    var v16 int8 = -128
                    var v17 uint8 = 2
                    _ = v16
                    _ = v17
                } else {
                    for fi18 := 0; fi18 < 15; fi18++ {
                        if true {
                            v19 := []int{(fi18 >> ((255) & 63)), fi18, 2147483647, fi7}
                            fmt.Println(v6, p1, fi7)
                            _ = v19
                        } else {
                            v20 := []int{(fi18 % ((fi18) | 1)), fi18, fi18, fi7}
                            v21 := GenS1{F0: 42, F1: 2}
                            _ = v20
                            _ = v21
                        }
                        p1 = (p0 * (p1 ^ 0))
                        _ = fi18
                    }
                    v22 := GenS1{F0: -128, F1: 205}
                    _ = v22
                }
            }
            v23 := []int{}
            if ((p0 + (^(p1))) > (p1 & (p0 * p0))) {
                p1 = (^(p0))
                for fi24 := 0; fi24 < 3; fi24++ {
                    fmt.Println(fi24)
                    switch (((("" + "a") + ("foo bar" + ("a" + ""))) + "foo bar") + "a") {
                    case "hello":
                        p0 -= (p1 >> ((1) & 63))
                    }
                    v6 = false
                    _ = fi24
                }
                v25 := GenS2{F0: 0, F1: "a"}
                _ = v25
            } else {
                p1 -= p1
                if len(v23) >= 2 {
                    v23 = v23[1:len(v23)]
                }
                v26 := GenS3{F0: genFn0(2147483647, "hello", 2), F1: -1, F2: 4294967295}
                _ = v26
            }
            _ = fi7
            _ = v23
        }
    }
    v27 := GenS1{F0: -128, F1: 255}
    v6 = true
    v28 := GenS3{F0: 983, F1: 127, F2: 42}
    _ = p0
    _ = p1
    _ = v27
    _ = v28
    _ = v6
    return 64
}

func main() {
    var v1 uint = 156
    var v2 int = int(v1)
    var v3 bool = (18446744073709551615 == 4294967295)
    if v3 {
        if false {
            if false {
                for fi29 := 0; fi29 < 16; fi29++ {
                    v30 := []int{}
                    _ = fi29
                    _ = v30
                }
                switch 127 {
                case 1:
                    for fi31 := 0; fi31 < 14; fi31++ {
                        for fi32 := 0; fi32 < 10; fi32++ {
                            if false {
                                var v33 int = (-(fi31))
                                if false {
                                    v2 -= fi31
                                    fmt.Println(v33, v2)
                                    v2 += fi31
                                } else {
                                    v34 := []int{fi31, (v2 % ((v33) | 1)), v33, fi32}
                                    if true {
                                        v35 := GenS4{F0: 127, F1: (fi31 - v33), F2: 1000}
                                        if len(v34) >= 2 {
                                            v34 = v34[1:len(v34)]
                                        }
                                        switch ("\n\t" + "hello") {
                                        case "hello":
                                            _ = len(v34)
                                        case "":
                                            var v36 int8 = 127
                                            _ = v36
                                        default:
                                            switch 0 {
                                            case 1:
                                                switch (fi32 * fi31) {
                                                case 9223372036854775807:
                                                    if len(v34) >= 2 {
                                                        v34 = v34[1:len(v34)]
                                                    }
                                                    var v37 int32 = 2147483647
                                                    _ = v37
                                                case 572:
                                                    var v38 int = (-(fi32))
                                                    _ = v38
                                                case 773:
                                                    v39 := GenS3{F0: genFn0(0, "foo bar", 2), F1: -1, F2: 996}
                                                    _ = v39
                                                }
                                            }
                                        }
                                        var v40 int8 = -76
                                        v41 := GenS5{F0: 1, F1: genFn1(255, 1)}
                                        _ = v35
                                        _ = v40
                                        _ = v41
                                    }
                                    if true {
                                        switch 219 {
                                        case 127:
                                            if v3 {
                                                fmt.Println(v2, fi32, fi31)
                                                if (1 > -2147483648) {
                                                    v2 = v33
                                                    fmt.Println(v33)
                                                }
                                                var v42 float64 = 1e10
                                                _ = v42
                                            } else {
                                                switch ("" + "hello") {
                                                case "a":
                                                    var v43 uint = v1
                                                    _ = v43
                                                case "\n\t":
                                                    v2 = v33
                                                default:
                                                    fmt.Println(v33)
                                                }
                                            }
                                        case 42:
                                            if v3 {
                                                v1 += v1
                                                fmt.Println(fi32)
                                            } else {
                                                var v44 uint32 = 42
                                                fmt.Println(v3, fi31)
                                                var v45 bool = (42 == 9223372036854775807)
                                                _ = v44
                                                _ = v45
                                            }
                                        case 4294967295:
                                            var v46 int32 = -23
                                            _ = v46
                                        default:
                                            switch "foo bar" {
                                            case "foo bar":
                                                switch "\n\t" {
                                                case "foo bar":
                                                    _ = len(v34)
                                                case "\n\t":
                                                    var v47 uint16 = 1
                                                    v48 := GenS1{F0: 84, F1: v1}
                                                    v1 += v1
                                                    _ = v47
                                                    _ = v48
                                                case "a":
                                                    if (v3 || false) {
                                                        if v3 {
                                                            if len(v34) > 0 {
                                                                v34[0] = 2147483647
                                                            }
                                                            var v49 int = (fi31 | 1)
                                                            _ = v49
                                                        }
                                                    }
                                                default:
                                                    fmt.Println(fi31, v1, fi31)
                                                }
                                            }
                                        }
                                    }
                                    _ = v34
                                }
                                _ = v33
                            } else {
                                var v50 int16 = 401
                                var v51 uint16 = 214
                                fmt.Println(v1, fi31)
                                _ = v50
                                _ = v51
                            }
                            var v52 bool = true
                            _ = fi32
                            _ = v52
                        }
                        var v53 int = v2
                        fmt.Println(fi31)
                        _ = fi31
                        _ = v53
                    }
                default:
                    var v54 int64 = -1
                    _ = v54
                }
            } else {
                for fi55 := 0; fi55 < 4; fi55++ {
                    if v3 {
                        for fi56 := 0; fi56 < 2; fi56++ {
                            if true {
                                v2 += v2
                                var v57 uint = v1
                                _ = v57
                            } else {
                                v3 = false
                                v58 := GenS6{F0: v3, F1: 300}
                                v59 := GenS7{F0: genFn0(2147483648, ("" + ""), 4), F1: -78}
                                _ = v58
                                _ = v59
                            }
                            fmt.Println(fi56, v2, fi55)
                            _ = fi56
                        }
                        v60 := []int{}
                        _ = v60
                    } else {
                        v61 := GenS5{F0: 1000, F1: 150}
                        _ = v61
                    }
                    v2 -= v2
                    v62 := GenS8{F0: 42}
                    _ = fi55
                    _ = v62
                }
                switch "hello" {
                case "a":
                    if v3 {
                        v63 := GenS4{F0: -21, F1: (v2 & (-(v2))), F2: -332}
                        _ = v63
                    }
                case "hello":
                    if v3 {
                        v64 := []int{v2, v2, (v2 >> (((^(v1))) & 63))}
                        if len(v64) >= 2 {
                            v64 = v64[1:len(v64)]
                        }
                        _ = v64
                    }
                case "":
                    switch -348 {
                    case 127:
                        switch ("hello" + ("foo bar" + "")) {
                        case "\n\t":
                            v65 := []int{(v2 & int(v1)), v2, 520, (v2 << ((42) & 63))}
                            _ = v65
                        }
                    }
                }
                fmt.Println(v3)
            }
            for fi66 := 0; fi66 < 7; fi66++ {
                if (546 == 2) {
                    v67 := []int{}
                    v68 := GenS2{F0: 0, F1: (("" + "foo bar") + "a")}
                    if true {
                        v2 = v2
                        var v69 int8 = 1
                        for fi70 := 0; fi70 < 4; fi70++ {
                            switch 127 {
                            case -219:
                                var v71 uint8 = 178
                                _ = v71
                            }
                            var v72 uint16 = 981
                            v72 += (v72 + (v72 * v72))
                            _ = fi70
                            _ = v72
                        }
                        _ = v69
                    }
                    _ = v67
                    _ = v68
                } else {
                    if v3 {
                        if (-89 <= -112) {
                            for fi73 := 0; fi73 < 2; fi73++ {
                                v74 := GenS4{F0: 127, F1: (fi66 & fi73), F2: -32768}
                                if false {
                                    fmt.Println(fi66, fi66, fi73)
                                    v1 = v1
                                    if (!(float64(0.0) != -2.25)) {
                                        v75 := GenS1{F0: 22, F1: (v1 * (v1 * v1))}
                                        _ = v75
                                    } else {
                                        var v76 string = ((((("a" + "hello") + "\n\t") + ("a" + "")) + ("" + ("\n\t" + "\n\t"))) + "a")
                                        _ = v76
                                    }
                                }
                                switch "" {
                                case "a":
                                    var v77 int8 = -9
                                    _ = v77
                                case "foo bar":
                                    var v78 uint32 = 2147483648
                                    _ = v78
                                case "\n\t":
                                    switch ("" + "hello") {
                                    case "hello":
                                        v79 := []int{(fi73 ^ fi73), 32767, v2, fi73}
                                        _ = v79
                                    case "foo bar":
                                        v80 := GenS9{F0: "hello", F1: genFn1(672, 124), F2: 65535}
                                        _ = v80
                                    case "\n\t":
                                        v3 = false
                                        switch int(v1) {
                                        case -208:
                                            v1 = v1
                                            v3 = (!v3)
                                            v2 -= (v2 * (fi66 << ((v1) & 63)))
                                        }
                                    default:
                                        var v81 float64 = 1.5
                                        _ = v81
                                    }
                                    v82 := GenS1{F0: 1, F1: v1}
                                    _ = v82
                                }
                                _ = fi73
                                _ = v74
                            }
                            v2 = (-(v2))
                            for fi83 := 0; fi83 < 14; fi83++ {
                                switch 656 {
                                case 998:
                                    v84 := []int{(v2 & fi66)}
                                    _ = v84
                                case 42:
                                    v2 -= fi66
                                    v2 = (fi66 >> (((v1 >> ((v1) & 63))) & 63))
                                    fmt.Println(v2, v3)
                                default:
                                    if true {
                                        var v85 uint32 = 0
                                        v86 := []int{-2147483648, (fi83 % (((fi83 * (v2 * fi66))) | 1)), fi66}
                                        if v3 {
                                            var v87 uint16 = 65535
                                            if v3 {
                                                var v88 uint64 = 458
                                                v2 = (^(fi66))
                                                var v89 float32 = 1e10
                                                switch ("foo bar" + ("\n\t" + "a")) {
                                                case "\n\t":
                                                    var v90 bool = v3
                                                    _ = v90
                                                default:
                                                    var v91 uint8 = 127
                                                    _ = v91
                                                }
                                                var v92 int32 = 42
                                                _ = v88
                                                _ = v89
                                                _ = v92
                                            }
                                            var v93 bool = (!true)
                                            _ = v87
                                            _ = v93
                                        } else {
                                            if false {
                                                var v94 int32 = 2147483647
                                                var v95 uint64 = 367
                                                var v96 uint8 = genFn1(65535, 782)
                                                _ = v94
                                                _ = v95
                                                _ = v96
                                            } else {
                                                v2 = fi83
                                                if (42 < 1000) {
                                                    var v97 uint32 = (v85 % ((501) | 1))
                                                    _ = v97
                                                } else {
                                                    fmt.Println(v3)
                                                    fmt.Println(fi66, fi83)
                                                    _ = len(v86)
                                                }
                                                v98 := GenS8{F0: -32768}
                                                _ = len(v86)
                                                var v99 int8 = 42
                                                v99 = v99
                                                switch v2 {
                                                case 32767:
                                                    v85 -= v85
                                                }
                                                v85 += (v85 >> (((^(v1))) & 63))
                                                v99 = (v99 >> ((255) & 63))
                                                var v100 bool = v3
                                                if len(v86) >= 2 {
                                                    v86 = v86[1:len(v86)]
                                                }
                                                v100 = v3
                                                if false {
                                                    v85 += v85
                                                }
                                                var v101 bool = false
                                                v1 = 18446744073709551615
                                                switch ("a" + "") {
                                                case "a":
                                                    var v102 int16 = int16(v1)
                                                    _ = v102
                                                case "hello":
                                                    v99 = -1
                                                case "foo bar":
                                                    v103 := GenS8{F0: 42}
                                                    _ = v103
                                                }
                                                _ = v100
                                                _ = v101
                                                _ = v98
                                                _ = v99
                                            }
                                            var v104 bool = v3
                                            _ = v104
                                        }
                                        _ = v85
                                        _ = v86
                                    } else {
                                        v1 = v1
                                        switch "\n\t" {
                                        case "hello":
                                            switch "hello" {
                                            case "":
                                                var v105 uint64 = 248
                                                _ = v105
                                            case "foo bar":
                                                v1 -= v1
                                            case "hello":
                                                v106 := []int{}
                                                _ = v106
                                            default:
                                                if v3 {
                                                    if (!v3) {
                                                        fmt.Println(fi66, v2)
                                                        v3 = (0.0 <= 1.5)
                                                    }
                                                    if v3 {
                                                        var v107 uint = v1
                                                        v108 := GenS10{F0: 2, F1: 42}
                                                        _ = v107
                                                        _ = v108
                                                    } else {
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            v109 := []int{(v2 & -42), fi83, fi66}
                                                            _ = v109
                                                        }
                                                        if v3 {
                                                            if ((!(v3 || false)) && v3) {
                                                                var v110 uint8 = 2
                                                                _ = v110
                                                            }
                                                            v111 := GenS11{F0: (^(v1)), F1: 42}
                                                            v1 = (v1 & (v1 ^ 97))
                                                            _ = v111
                                                        }
                                                    }
                                                    v112 := GenS6{F0: v3, F1: v1}
                                                    v113 := []int{(fi83 - v2)}
                                                    if (((0.0 * 3.14159) - ((1.5 - (0.0 / 0.1)) * 3.14159)) == (1.5 - (0.1 * (0.0 - 3.14159)))) {
                                                        var v114 int32 = 202
                                                        v113 = append(v113, 128)
                                                        _ = v114
                                                    } else {
                                                        var v115 int64 = -402
                                                        _ = v115
                                                    }
                                                    _ = v112
                                                    _ = v113
                                                } else {
                                                    v1 = v1
                                                    var v116 int = 42
                                                    var v117 float32 = -2.25
                                                    var v118 int = -9223372036854775808
                                                    _ = v116
                                                    _ = v117
                                                    _ = v118
                                                }
                                            }
                                        case "a":
                                            v119 := []int{fi83}
                                            _ = v119
                                        case "foo bar":
                                            var v120 int16 = -107
                                            var v121 int8 = 89
                                            var v122 bool = v3
                                            _ = v120
                                            _ = v121
                                            _ = v122
                                        }
                                    }
                                }
                                _ = fi83
                            }
                        }
                    }
                }
                _ = fi66
            }
        }
        v2 = 9223372036854775807
        v2 += v2
    }
    for fi123 := 0; fi123 < 20; fi123++ {
        if (0.1 <= (0.0 - 1.5)) {
            fmt.Println(v3, fi123, fi123)
        } else {
            v124 := GenS11{F0: v1, F1: 32767}
            _ = v124
        }
        v125 := GenS12{F0: 4294967295}
        _ = fi123
        _ = v125
    }
    v126 := GenS10{F0: 638, F1: 59}
    v127 := GenS6{F0: (42 >= 255), F1: (v1 % ((v1) | 1))}
    if v3 {
        for fi128 := 0; fi128 < 5; fi128++ {
            if v3 {
                var v129 int8 = -128
                var v130 uint32 = 2
                if true {
                    switch ("hello" + "a") {
                    case "":
                        switch "hello" {
                        case "a":
                            var v131 float64 = (1e10 + 0.1)
                            _ = v131
                        case "\n\t":
                            v129 += -38
                        default:
                            if v3 {
                                v132 := GenS8{F0: 0}
                                var v133 int = (fi128 + -107)
                                _ = v132
                                _ = v133
                            }
                        }
                    default:
                        switch v2 {
                        case -2147483648:
                            var v134 uint64 = 405
                            _ = v134
                        case 4294967295:
                            if (3.14159 < (1e10 * 1.5)) {
                                if (!(!v3)) {
                                    for fi135 := 0; fi135 < 4; fi135++ {
                                        v136 := []int{v2, (fi128 >> (((v1 - 4294967295)) & 63)), -911}
                                        switch (fi128 * (fi135 | (v2 << ((107) & 63)))) {
                                        case 1:
                                            var v137 int64 = -711
                                            _ = v137
                                        case 128:
                                            v2 = (^(fi128))
                                        }
                                        _ = fi135
                                        _ = v136
                                    }
                                    v1 += v1
                                    v138 := []int{128, 32767}
                                    _ = v138
                                } else {
                                    var v139 int = v2
                                    v140 := GenS8{F0: -1}
                                    if false {
                                        switch "\n\t" {
                                        case "hello":
                                            v141 := []int{}
                                            _ = v141
                                        case "":
                                            if (368 <= v1) {
                                                var v142 int16 = -111
                                                v143 := GenS13{F0: v130, F1: (^(v1)), F2: (v130 >> (((^(v1))) & 63))}
                                                _ = v142
                                                _ = v143
                                            }
                                        }
                                        v139 = v139
                                        v144 := GenS13{F0: v130, F1: v1, F2: v130}
                                        _ = v144
                                    } else {
                                        if v3 {
                                            if (1e10 >= 1.5) {
                                                fmt.Println(v3)
                                                var v145 string = ("hello" + "hello")
                                                var v146 float64 = -2.25
                                                _ = v145
                                                _ = v146
                                            } else {
                                                v147 := []int{}
                                                var v148 float64 = 0.0
                                                _ = v147
                                                _ = v148
                                            }
                                            v149 := []int{v139, 42, (v2 | (fi128 & (v2 / (((-(v2))) | 1)))), (v139 / ((4294967295) | 1))}
                                            _ = v149
                                        } else {
                                            switch "a" {
                                            case "\n\t":
                                                fmt.Println(v139, v2)
                                            case "a":
                                                var v150 int8 = 127
                                                _ = v150
                                            }
                                            fmt.Println(v139)
                                            fmt.Println(v1, v130)
                                        }
                                        var v151 int16 = 0
                                        _ = v151
                                    }
                                    _ = v139
                                    _ = v140
                                }
                                switch (("a" + ((("" + "hello") + "a") + "\n\t")) + "hello") {
                                case "hello":
                                    v152 := GenS13{F0: 65535, F1: 4294967295, F2: v130}
                                    _ = v152
                                case "foo bar":
                                    for fi153 := 0; fi153 < 15; fi153++ {
                                        switch -847 {
                                        case 619:
                                            v3 = (v3 || true)
                                        case 2147483647:
                                            if v3 {
                                                switch fi153 {
                                                case 4294967295:
                                                    v154 := GenS12{F0: (v130 << (((v1 * (v1 * v1))) & 63))}
                                                    _ = v154
                                                case 0:
                                                    fmt.Println(v2, v2, v1)
                                                case -912:
                                                    fmt.Println(fi153, v1, fi153)
                                                }
                                                v3 = v3
                                                var v155 int64 = -1
                                                if false {
                                                    var v156 int = v2
                                                    v3 = v3
                                                    _ = v156
                                                } else {
                                                    v3 = true
                                                    v155 -= v155
                                                    switch fi153 {
                                                    case -551:
                                                        v155 = (^(v155))
                                                    case 127:
                                                        switch (fi128 >> ((v1) & 63)) {
                                                        case -2147483648:
                                                            switch "a" {
                                                            case "":
                                                                fmt.Println(fi153, v130, v129)
                                                            case "hello":
                                                                v129 = (v129 / (((^(v129))) | 1))
                                                            case "a":
                                                                fmt.Println(v2, v2, v129)
                                                            default:
                                                                v157 := []int{-2147483648, 506, fi153}
                                                                _ = v157
                                                            }
                                                        }
                                                    default:
                                                        v158 := []int{v2}
                                                        _ = v158
                                                    }
                                                    var v159 int8 = (v129 - int8(v155))
                                                    if false {
                                                        v160 := []int{2147483647}
                                                        _ = v160
                                                    }
                                                    _ = v159
                                                }
                                                _ = v155
                                            } else {
                                                var v161 int16 = -456
                                                var v162 int16 = v161
                                                if true {
                                                    if v3 {
                                                        var v163 float64 = float64((0.1 / -2.25))
                                                        _ = v163
                                                    } else {
                                                        v164 := GenS14{F0: 2147483647, F1: int8(v130), F2: v130}
                                                        v165 := []int{}
                                                        _ = v164
                                                        _ = v165
                                                    }
                                                    fmt.Println(v2, v129, v3)
                                                    v166 := []int{fi128, 4294967295, (fi128 * fi153)}
                                                    _ = v166
                                                }
                                                _ = v161
                                                _ = v162
                                            }
                                        default:
                                            var v167 uint8 = 127
                                            _ = v167
                                        }
                                        var v168 uint = 1
                                        v130 = uint32(v168)
                                        _ = fi153
                                        _ = v168
                                    }
                                case "":
                                    switch -937 {
                                    case 9223372036854775807:
                                        fmt.Println(v2, v3)
                                    }
                                default:
                                    if v3 {
                                        if false {
                                            for fi169 := 0; fi169 < 9; fi169++ {
                                                if false {
                                                    v170 := GenS15{F0: genFn0(v130, "", 236)}
                                                    v129 = (v129 - (v129 ^ (v129 - v129)))
                                                    _ = v170
                                                } else {
                                                    var v171 float32 = -2.25
                                                    _ = v171
                                                }
                                                _ = fi169
                                            }
                                            v172 := []int{(fi128 | 32767), (fi128 ^ (^(v2))), (fi128 >> ((4294967295) & 63)), 1}
                                            _ = v172
                                        }
                                    }
                                }
                            }
                        default:
                            if v3 {
                                if v3 {
                                    if (133 < 170) {
                                        var v173 int32 = 42
                                        for fi174 := 0; fi174 < 2; fi174++ {
                                            switch (v2 - (fi174 % ((fi174) | 1))) {
                                            case 2147483647:
                                                var v175 uint64 = 18446744073709551615
                                                _ = v175
                                            case 4294967295:
                                                fmt.Println(fi128, v173)
                                            case 988:
                                                v1 = 255
                                            }
                                            switch "a" {
                                            case "":
                                                fmt.Println(fi128, v3)
                                            }
                                            _ = fi174
                                        }
                                        _ = v173
                                    }
                                    v130 = v130
                                    var v176 uint = v1
                                    _ = v176
                                } else {
                                    var v177 float32 = 1.5
                                    _ = v177
                                }
                            }
                        }
                    }
                    if (v1 < (v1 << ((v1) & 63))) {
                        v1 = (v1 | 65535)
                    } else {
                        v178 := []int{984, fi128, 32768}
                        _ = v178
                    }
                }
                _ = v129
                _ = v130
            }
            v179 := GenS1{F0: -128, F1: (v1 ^ v1)}
            _ = fi128
            _ = v179
        }
        var v180 bool = v3
        for fi181 := 0; fi181 < 14; fi181++ {
            v180 = v180
            _ = fi181
        }
        _ = v180
    } else {
        for fi182 := 0; fi182 < 6; fi182++ {
            v183 := GenS16{F0: -74}
            var v184 int32 = 1
            fmt.Println(fi182, v184, v1)
            _ = fi182
            _ = v183
            _ = v184
        }
        v2 += (v2 & 32768)
    }
    v185 := []int{v2, 1, -1}
    v1 += v1
    v185 = append(v185, -363)
    for fi186 := 0; fi186 < 6; fi186++ {
        if true {
            switch (v2 ^ -984) {
            case 127:
                var v187 int8 = 42
                _ = v187
            case 42:
                for fi188 := 0; fi188 < 10; fi188++ {
                    v3 = (v1 < v1)
                    v1 -= 0
                    _ = fi188
                }
            case 32768:
                if (("hello" <= ("foo bar" + "a")) || v3) {
                    v189 := GenS2{F0: 127, F1: "a"}
                    var v190 bool = (1.5 <= float32((1.5 * (0.1 + ((-2.25 + (1.5 + 1e10)) + (1e10 - (float64((3.14159 * 3.14159)) * ((0.1 - (1.5 * 0.0)) * (((3.14159 - -2.25) - (0.0 + float64(1e10))) - (((3.14159 * (1.5 / (-2.25 * 1e10))) / float64(1e10)) / 1e10))))))))))
                    for fi191 := 0; fi191 < 12; fi191++ {
                        switch (^(fi191)) {
                        case 541:
                            var v192 float32 = float32((0.1 + 0.0))
                            _ = v192
                        case 9223372036854775807:
                            fmt.Println(v3)
                            v193 := GenS12{F0: 23}
                            var v194 float32 = (0.1 * float32((3.14159 * (float64((1e10 * 1.5)) / 1.5))))
                            _ = v193
                            _ = v194
                        default:
                            v190 = v3
                            switch (fi191 * fi191) {
                            case -80:
                                v195 := GenS17{F0: ("a" + "a"), F1: 18446744073709551615, F2: 1000}
                                _ = v195
                            case -187:
                                v196 := GenS18{F0: fi186, F1: "", F2: 65535}
                                fmt.Println(fi186)
                                switch 339 {
                                case 773:
                                    fmt.Println(fi186, v190)
                                }
                                _ = v196
                            case 2147483647:
                                v3 = (!true)
                            default:
                                if v3 {
                                    v190 = (0.1 == 0.1)
                                }
                            }
                        }
                        _ = fi191
                    }
                    _ = v189
                    _ = v190
                } else {
                    var v197 int32 = 523
                    switch "hello" {
                    case "\n\t":
                        v197 -= -2147483648
                    }
                    _ = v197
                }
            }
        } else {
            for fi198 := 0; fi198 < 16; fi198++ {
                var v199 float32 = 0.1
                _ = fi198
                _ = v199
            }
            fmt.Println(v1, v3, v2)
            var v200 int64 = 42
            _ = v200
        }
        if (!v3) {
            if true {
                v3 = v3
                _ = len(v185)
            } else {
                var v201 uint8 = genFn1(350, 65535)
                v202 := GenS1{F0: 1, F1: v1}
                for fi203 := 0; fi203 < 8; fi203++ {
                    v185 = append(v185, -321)
                    fmt.Println(fi186)
                    fmt.Println(fi203)
                    _ = fi203
                }
                _ = v201
                _ = v202
            }
            if (-156 == 33) {
                _ = len(v185)
                v3 = v3
                v204 := GenS4{F0: -1, F1: fi186, F2: 42}
                _ = v204
            } else {
                v3 = v3
                v205 := GenS19{F0: 2}
                _ = v205
            }
            var v206 bool = v3
            _ = v206
        } else {
            v207 := GenS18{F0: (fi186 | -497), F1: ((("foo bar" + "a") + "\n\t") + "\n\t"), F2: 32767}
            switch "foo bar" {
            case "a":
                switch ("\n\t" + "foo bar") {
                case "\n\t":
                    var v208 uint16 = 32767
                    _ = v208
                default:
                    if len(v185) > 0 {
                        v185[0] = -1
                    }
                }
            }
            v1 = (v1 << ((255) & 63))
            _ = v207
        }
        v2 += fi186
        _ = fi186
    }
    if false {
        fmt.Println(v2)
    }
    v3 = (v3 || v3)
    var v209 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v185:", len(v185), v185)
    fmt.Println("v2:", v2)
    fmt.Println("v209:", v209)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v126
    _ = v127
    _ = v185
    _ = v2
    _ = v209
    _ = v3
}
