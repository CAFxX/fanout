package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int32; F1 int; F2 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS5 struct { F0 int32; F1 uint8 }
type GenS6 struct { F0 uint8; F1 uint8 }
type GenS7 struct { F0 int8; F1 int64; F2 int }
type GenS8 struct { F0 int8; F1 int; F2 string }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 bool; F1 int32 }
type GenS11 struct { F0 uint8; F1 int32 }
func genFn0(p0 int8, p1 int16) int8 {
    if (9223372036854775807 != 42) {
        if false {
            if false {
                var v7 uint = 4294967295
                _ = v7
            }
            if (true || true) {
                v8 := GenS1{F0: 0, F1: 1}
                switch ("foo bar" + "a") {
                case "\n\t":
                    for fi9 := 0; fi9 < 3; fi9++ {
                        var v10 uint8 = 128
                        var v11 uint8 = v10
                        if false {
                            var v12 int = fi9
                            _ = v12
                        } else {
                            for fi13 := 0; fi13 < 15; fi13++ {
                                fmt.Println(fi13)
                                fmt.Println(fi9, p1)
                                v14 := GenS1{F0: 1, F1: 0}
                                _ = fi13
                                _ = v14
                            }
                            for fi15 := 0; fi15 < 20; fi15++ {
                                v11 = v11
                                v16 := GenS1{F0: 65535, F1: 2147483647}
                                _ = fi15
                                _ = v16
                            }
                            var v17 string = "a"
                            _ = v17
                        }
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                case "hello":
                    if false {
                        for fi18 := 0; fi18 < 3; fi18++ {
                            p0 = p0
                            v19 := GenS1{F0: 42, F1: 451}
                            _ = fi18
                            _ = v19
                        }
                    }
                }
                _ = v8
            }
        }
        var v20 int64 = 42
        for fi21 := 0; fi21 < 19; fi21++ {
            var v22 string = ("a" + "foo bar")
            if (false || true) {
                for fi23 := 0; fi23 < 16; fi23++ {
                    fmt.Println(p0)
                    var v24 int16 = (p1 / ((int16(fi23)) | 1))
                    _ = fi23
                    _ = v24
                }
                switch (fi21 / (((-(fi21))) | 1)) {
                case 127:
                    fmt.Println(fi21)
                default:
                    v22 = v22
                }
                switch "a" {
                case "a":
                    for fi25 := 0; fi25 < 11; fi25++ {
                        switch v22 {
                        case "foo bar":
                            if (!true) {
                                if ((2 == 1) && (false || true)) {
                                    fmt.Println(p1, p0)
                                    var v26 uint16 = 2
                                    if false {
                                        v22 = v22
                                    }
                                    _ = v26
                                } else {
                                    if (("\n\t" + v22) > v22) {
                                        v27 := []int{fi25}
                                        if false {
                                            v28 := GenS1{F0: 4294967295, F1: 64}
                                            _ = v28
                                        } else {
                                            if len(v27) > 0 {
                                                v27[0] = fi21
                                            }
                                            fmt.Println(fi25, v20, v20)
                                        }
                                        var v29 uint8 = 1
                                        _ = v27
                                        _ = v29
                                    } else {
                                        var v30 uint32 = 641
                                        _ = v30
                                    }
                                }
                                v22 = "a"
                                fmt.Println(p1)
                            } else {
                                v31 := GenS1{F0: 255, F1: 22}
                                v32 := []int{(-(fi21)), fi21}
                                _ = v31
                                _ = v32
                            }
                            switch (fi21 - 623) {
                            case 975:
                                fmt.Println(v22)
                                var v33 int64 = v20
                                var v34 float32 = (((1.5 + -2.25) * 1e10) - float32((0.1 + 1e10)))
                                _ = v33
                                _ = v34
                            default:
                                if false {
                                    p1 += (^(p1))
                                    p0 += p0
                                    v35 := GenS2{F0: 589, F1: (fi21 + fi21), F2: "a"}
                                    _ = v35
                                }
                            }
                        case "hello":
                            var v36 uint16 = 255
                            _ = v36
                        case "":
                            switch "" {
                            case "foo bar":
                                if false {
                                    if (!(!true)) {
                                        v37 := GenS2{F0: 849, F1: (fi21 & (fi25 + fi25)), F2: v22}
                                        var v38 bool = false
                                        if ((1e10 - -2.25) < 0.1) {
                                            p0 = p0
                                        }
                                        _ = v37
                                        _ = v38
                                    }
                                    var v39 int64 = (v20 - v20)
                                    v20 = 9223372036854775807
                                    _ = v39
                                }
                            default:
                                p0 = (p0 / ((p0) | 1))
                            }
                        default:
                            var v40 uint8 = 128
                            var v41 int16 = (-(p1))
                            v42 := []int{(fi21 / ((478) | 1)), fi25, (fi21 >> ((42) & 63)), (-(fi21))}
                            var v43 int64 = (^(v20))
                            if true {
                                var v44 int32 = -576
                                var v45 int = fi25
                                v42 = append(v42, (v45 + -209))
                                var v46 int16 = 32767
                                _ = v44
                                _ = v45
                                _ = v46
                            } else {
                                fmt.Println(fi21, v40, v20)
                            }
                            _ = v40
                            _ = v41
                            _ = v42
                            _ = v43
                        }
                        if false {
                            var v47 string = ("\n\t" + "hello")
                            v48 := GenS3{F0: -480}
                            p0 -= (p0 << ((349) & 63))
                            p1 = (p1 / ((p1) | 1))
                            _ = v47
                            _ = v48
                        } else {
                            v49 := []int{fi25}
                            v50 := GenS1{F0: 2, F1: 1}
                            var v51 int32 = 1
                            _ = v49
                            _ = v50
                            _ = v51
                        }
                        _ = fi25
                    }
                case "\n\t":
                    for fi52 := 0; fi52 < 16; fi52++ {
                        var v53 float64 = (float64(-2.25) - (float64(0.0) / 1e10))
                        if false {
                            v54 := []int{(fi52 ^ (fi52 << ((988) & 63))), -1, (fi52 % ((fi52) | 1))}
                            if false {
                                if false {
                                    _ = len(v54)
                                } else {
                                    fmt.Println(fi52, v53, v53)
                                }
                                var v55 uint64 = 4294967295
                                _ = v55
                            } else {
                                if len(v54) > 0 {
                                    v54[0] = fi21
                                }
                            }
                            _ = v54
                        }
                        _ = fi52
                        _ = v53
                    }
                default:
                    for fi56 := 0; fi56 < 1; fi56++ {
                        v57 := []int{fi56, 42, 425}
                        if (true && true) {
                            if len(v57) >= 2 {
                                v57 = v57[1:len(v57)]
                            }
                        } else {
                            var v58 int8 = 42
                            if true {
                                p1 = int16(fi21)
                                v22 = v22
                                var v59 uint8 = 2
                                var v60 bool = false
                                v61 := GenS4{F0: 214, F1: (-(p1)), F2: v59}
                                _ = v59
                                _ = v60
                                _ = v61
                            }
                            v62 := GenS1{F0: 65535, F1: 2}
                            v58 -= v58
                            var v63 int32 = 893
                            _ = v58
                            _ = v62
                            _ = v63
                        }
                        _ = fi56
                        _ = v57
                    }
                }
            } else {
                switch fi21 {
                case 42:
                    v64 := GenS5{F0: -767, F1: 42}
                    _ = v64
                case 860:
                    if ((-694 >= fi21) && true) {
                        switch fi21 {
                        case -2147483648:
                            fmt.Println(p1, p0, p1)
                        case -690:
                            v65 := []int{fi21, fi21, -2147483648}
                            _ = v65
                        case 128:
                            for fi66 := 0; fi66 < 4; fi66++ {
                                p1 -= 92
                                var v67 uint32 = 0
                                _ = fi66
                                _ = v67
                            }
                        }
                        if (4294967295 >= 42) {
                            v68 := []int{128, fi21, 1}
                            var v69 int32 = 2147483647
                            var v70 uint64 = 4294967295
                            _ = v68
                            _ = v69
                            _ = v70
                        } else {
                            p1 = (p1 & p1)
                        }
                        if true {
                            var v71 int32 = 276
                            v72 := []int{(-(fi21)), (fi21 | 9223372036854775807), (fi21 & fi21)}
                            _ = v71
                            _ = v72
                        } else {
                            if false {
                                v73 := GenS6{F0: 2, F1: 21}
                                v20 += (v20 * (v20 << ((130) & 63)))
                                _ = v73
                            } else {
                                v74 := GenS2{F0: -1, F1: (fi21 >> ((522) & 63)), F2: v22}
                                _ = v74
                            }
                            v75 := []int{(fi21 >> ((65535) & 63))}
                            if (true || true) {
                                var v76 uint64 = 18446744073709551615
                                fmt.Println(v22, p1, v22)
                                if false {
                                    var v77 int32 = -268
                                    if true {
                                        if true {
                                            if len(v75) >= 2 {
                                                v75 = v75[1:len(v75)]
                                            }
                                            if true {
                                                var v78 float32 = 0.1
                                                for fi79 := 0; fi79 < 18; fi79++ {
                                                    if true {
                                                        if (140 != 42) {
                                                            switch fi21 {
                                                            case 909:
                                                                fmt.Println(v77, v20)
                                                            case 0:
                                                                fmt.Println(v22, v78)
                                                            }
                                                            var v80 int32 = 2147483647
                                                            v77 += (v77 | v80)
                                                            _ = v80
                                                        }
                                                        if true {
                                                            if false {
                                                                switch ("foo bar" + "foo bar") {
                                                                case "a":
                                                                    var v81 uint16 = 42
                                                                    _ = v81
                                                                }
                                                            }
                                                        } else {
                                                            v75 = append(v75, (-(fi79)))
                                                            v22 = v22
                                                        }
                                                        p1 = p1
                                                    } else {
                                                        v82 := GenS7{F0: p0, F1: (v20 & (-(v20))), F2: fi21}
                                                        _ = v82
                                                    }
                                                    _ = fi79
                                                }
                                                v20 = v20
                                                _ = v78
                                            }
                                        }
                                    }
                                    _ = v77
                                } else {
                                    if true {
                                        v75 = append(v75, fi21)
                                        var v83 uint16 = 1
                                        _ = v83
                                    }
                                    switch (-(fi21)) {
                                    case 2147483647:
                                        if false {
                                            v84 := GenS2{F0: 720, F1: fi21, F2: ""}
                                            var v85 float64 = (((((0.1 / 1.0) + (0.1 - 0.1)) + 0.0) / (1e10 / float64(float32((1e10 * (float64(((-2.25 / 1.0) * 1e10)) / 1.5)))))) * (1.5 - (0.1 + 1.5)))
                                            v75 = append(v75, -2147483648)
                                            _ = v84
                                            _ = v85
                                        }
                                    default:
                                        p0 = (p0 & (p0 * p0))
                                    }
                                    for fi86 := 0; fi86 < 11; fi86++ {
                                        v22 = v22
                                        var v87 uint64 = 0
                                        _ = fi86
                                        _ = v87
                                    }
                                }
                                _ = v76
                            } else {
                                if true {
                                    fmt.Println(p0, v20)
                                }
                                for fi88 := 0; fi88 < 13; fi88++ {
                                    if false {
                                        v89 := GenS1{F0: 4294967295, F1: 42}
                                        switch -585 {
                                        case -78:
                                            switch "\n\t" {
                                            case "":
                                                fmt.Println(v20)
                                            }
                                        case -9223372036854775808:
                                            v90 := GenS4{F0: 1, F1: 0, F2: 42}
                                            _ = v90
                                        default:
                                            if (-1 != 93) {
                                                _ = len(v75)
                                                var v91 uint = 77
                                                _ = v91
                                            }
                                        }
                                        fmt.Println(v22)
                                        _ = v89
                                    } else {
                                        switch 128 {
                                        case 4294967295:
                                            if true {
                                                if len(v75) > 0 {
                                                    v75[0] = fi88
                                                }
                                            }
                                        case 128:
                                            var v92 string = v22
                                            _ = v92
                                        case 465:
                                            switch (fi21 - fi88) {
                                            case 0:
                                                _ = len(v75)
                                            default:
                                                var v93 int32 = -1
                                                p0 -= (p0 | (p0 >> ((4294967295) & 63)))
                                                v22 = (v22 + "foo bar")
                                                _ = v93
                                            }
                                        }
                                        v20 += v20
                                        var v94 uint32 = 421
                                        v75 = append(v75, fi88)
                                        _ = v94
                                    }
                                    _ = fi88
                                }
                                if (v22 <= "hello") {
                                    switch fi21 {
                                    case 663:
                                        v20 += int64(p1)
                                    }
                                    var v95 bool = false
                                    _ = v95
                                }
                            }
                            _ = v75
                        }
                    }
                }
                v96 := GenS6{F0: 176, F1: 2}
                p1 -= p1
                _ = v96
            }
            v97 := []int{(fi21 & fi21), -2147483648, (^(fi21)), (fi21 & (fi21 << ((699) & 63)))}
            _ = fi21
            _ = v22
            _ = v97
        }
        _ = v20
    }
    for fi98 := 0; fi98 < 6; fi98++ {
        var v99 int64 = 42
        switch 257 {
        case -1:
            var v100 int8 = int8(fi98)
            _ = v100
        case 42:
            if ((-(v99)) == 42) {
                v101 := GenS3{F0: int32(v99)}
                p1 -= 1000
                _ = v101
            }
        case 0:
            for fi102 := 0; fi102 < 2; fi102++ {
                var v103 float64 = 0.0
                if false {
                    if true {
                        switch fi102 {
                        case 42:
                            var v104 float32 = (((1.5 + 1.5) * -2.25) * 1e10)
                            _ = v104
                        case -375:
                            v105 := GenS6{F0: 249, F1: 124}
                            _ = v105
                        }
                    }
                    v103 += v103
                }
                if ((65535 <= 696) && ((!(!(v99 != (v99 | v99)))) && true)) {
                    v103 += ((v103 + (0.0 * v103)) / float64(1e10))
                }
                _ = fi102
                _ = v103
            }
        default:
            p0 = (p0 + 42)
        }
        fmt.Println(p0, p0, p1)
        _ = fi98
        _ = v99
    }
    for fi106 := 0; fi106 < 10; fi106++ {
        if ((fi106 + 315) != (fi106 | fi106)) {
            switch "" {
            case "":
                p1 -= (p1 & 1000)
            case "a":
                var v107 int16 = (p1 + 32767)
                _ = v107
            case "foo bar":
                for fi108 := 0; fi108 < 9; fi108++ {
                    var v109 int16 = (^(p1))
                    var v110 uint32 = 42
                    _ = fi108
                    _ = v109
                    _ = v110
                }
            default:
                p0 = (p0 * p0)
            }
            p0 = p0
            for fi111 := 0; fi111 < 2; fi111++ {
                if (p0 < (p0 >> ((1) & 63))) {
                    v112 := []int{(^(fi106)), (fi111 / (((fi111 + fi106)) | 1))}
                    v113 := GenS8{F0: (p0 ^ 93), F1: (^(fi111)), F2: "hello"}
                    if len(v112) > 0 {
                        v112[0] = fi106
                    }
                    _ = v112
                    _ = v113
                } else {
                    var v114 string = ""
                    _ = v114
                }
                fmt.Println(p1, fi106)
                fmt.Println(p0)
                _ = fi111
            }
        }
        p1 = (p1 & 1000)
        _ = fi106
    }
    switch 9223372036854775807 {
    case 510:
        for fi115 := 0; fi115 < 2; fi115++ {
            var v116 uint8 = 128
            var v117 int8 = (p0 * p0)
            _ = fi115
            _ = v116
            _ = v117
        }
    case 0:
        fmt.Println(p1, p1)
    }
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 uint8) int8 {
    if (65535 >= 2) {
        switch "foo bar" {
        case "foo bar":
            fmt.Println(p0)
        case "":
            var v118 uint64 = 348
            _ = v118
        default:
            if false {
                if true {
                    p0 = (^(p0))
                    p0 += (p0 & p0)
                }
                fmt.Println(p0)
            } else {
                p0 -= p0
            }
        }
        v119 := []int{-2147483648}
        fmt.Println(p0)
        _ = v119
    }
    v120 := GenS6{F0: p0, F1: (p0 / ((128) | 1))}
    v121 := []int{-1}
    _ = p0
    _ = v120
    _ = v121
    return 42
}

func main() {
    v1 := []uint32{1, 0, 4294967295}
    var v2 int64 = 0
    v3 := []uint32{65535, 1, 65535, 0}
    var v4 uint64 = 843
    var v5 string = ""
    var v6 int64 = (v2 / ((v2) | 1))
    v122 := GenS3{F0: -2147483648}
    if (42 != 1) {
        var v123 int = 0
        for fi124 := 0; fi124 < 18; fi124++ {
            v123 = fi124
            var v125 bool = true
            _ = fi124
            _ = v125
        }
        _ = v123
    }
    fmt.Println(v6, v2, v4)
    if (false && ("\n\t" < v5)) {
        v5 = (("\n\t" + v5) + ((v5 + ((((v5 + v5) + "\n\t") + v5) + "")) + (v5 + v5)))
    } else {
        v2 = -9223372036854775808
        var v126 bool = (42 != 255)
        v127 := GenS4{F0: 119, F1: 651, F2: 1}
        _ = v126
        _ = v127
    }
    var v128 int16 = 23
    if len(v3) > 0 {
        v3[0] = 932
    }
    switch 2147483647 {
    case 127:
        v4 += v4
    case 2147483647:
        var v129 uint16 = 255
        _ = v129
    case 1:
        v130 := GenS7{F0: 42, F1: (v6 / ((v2) | 1)), F2: 550}
        _ = v130
    }
    var v131 int = 4294967295
    v132 := GenS5{F0: 2147483647, F1: 1}
    v133 := GenS3{F0: -567}
    for fi134 := 0; fi134 < 9; fi134++ {
        var v135 int16 = 1000
        fmt.Println(v128, v131)
        for fi136 := 0; fi136 < 1; fi136++ {
            if true {
                v135 = 42
            } else {
                if ((1e10 == (0.0 - 3.14159)) && false) {
                    fmt.Println(fi134, v6)
                    var v137 int = fi136
                    _ = len(v3)
                    _ = v137
                } else {
                    v138 := GenS4{F0: 127, F1: (v135 >> ((65535) & 63)), F2: 128}
                    if ((42 > 241) && (false && true)) {
                        v6 = 42
                        var v139 uint64 = (v4 / ((95) | 1))
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v139
                    } else {
                        v140 := GenS7{F0: 76, F1: (v6 << ((210) & 63)), F2: v131}
                        switch fi134 {
                        case -2147483648:
                            v5 = ((v5 + (v5 + "\n\t")) + (((v5 + "") + v5) + "a"))
                        default:
                            switch v131 {
                            case 195:
                                switch v5 {
                                case "a":
                                    v4 += v4
                                case "\n\t":
                                    switch 32767 {
                                    case -1:
                                        if (!true) {
                                            v6 = (v6 * -9223372036854775808)
                                        } else {
                                            var v141 uint32 = 955
                                            v128 += 42
                                            var v142 uint16 = 351
                                            _ = v141
                                            _ = v142
                                        }
                                    default:
                                        var v143 uint = 4294967295
                                        _ = v143
                                    }
                                case "":
                                    switch v131 {
                                    case 458:
                                        switch (fi136 << ((2) & 63)) {
                                        case 726:
                                            v4 = 42
                                        case 471:
                                            var v144 bool = true
                                            _ = v144
                                        default:
                                            switch "" {
                                            case "hello":
                                                var v145 string = v5
                                                _ = v145
                                            case "":
                                                if len(v1) > 0 {
                                                    v1[0] = 65535
                                                }
                                            case "a":
                                                switch fi136 {
                                                case -896:
                                                    v131 += -1
                                                case 127:
                                                    v146 := GenS4{F0: 155, F1: v128, F2: 1}
                                                    _ = v146
                                                case 32767:
                                                    if true {
                                                        var v147 bool = false
                                                        if (false || v147) {
                                                            var v148 string = v5
                                                            var v149 uint = 0
                                                            v150 := GenS9{F0: -1}
                                                            _ = v148
                                                            _ = v149
                                                            _ = v150
                                                        }
                                                        switch fi136 {
                                                        case -647:
                                                            fmt.Println(v135, v147, v135)
                                                        case -615:
                                                            v151 := GenS1{F0: 255, F1: 42}
                                                            _ = v151
                                                        default:
                                                            var v152 int = (fi136 * v131)
                                                            _ = v152
                                                        }
                                                        _ = v147
                                                    }
                                                default:
                                                    var v153 float32 = 1e10
                                                    _ = v153
                                                }
                                                v2 = (v2 >> ((984) & 63))
                                                var v154 float32 = float32(((3.14159 - 1.5) * -2.25))
                                                var v155 uint64 = (v4 / ((v4) | 1))
                                                if (false && true) {
                                                    v156 := GenS3{F0: -1}
                                                    if false {
                                                        switch (v5 + v5) {
                                                        case "\n\t":
                                                            v135 += 42
                                                        default:
                                                            var v157 uint = 42
                                                            _ = v157
                                                        }
                                                    } else {
                                                        v158 := GenS2{F0: -1, F1: (v131 % ((fi136) | 1)), F2: "\n\t"}
                                                        if (!true) {
                                                            switch (v5 + (((v5 + "foo bar") + v5) + "a")) {
                                                            case "hello":
                                                                if ((^(v4)) > (v155 ^ (v155 >> ((255) & 63)))) {
                                                                    var v159 int = (fi134 * fi136)
                                                                    var v160 uint16 = 255
                                                                    var v161 uint64 = (v155 / ((0) | 1))
                                                                    _ = v159
                                                                    _ = v160
                                                                    _ = v161
                                                                } else {
                                                                    var v162 string = v5
                                                                    _ = v162
                                                                }
                                                            case "foo bar":
                                                                fmt.Println(v155)
                                                            case "\n\t":
                                                                var v163 int64 = (v6 >> ((313) & 63))
                                                                _ = v163
                                                            default:
                                                                v164 := GenS10{F0: (1.5 == 0.0), F1: 0}
                                                                _ = v164
                                                            }
                                                        } else {
                                                            if true {
                                                                v135 += v128
                                                                switch v5 {
                                                                case "a":
                                                                    var v165 int = 728
                                                                    _ = v165
                                                                case "":
                                                                    v5 = ((v5 + v5) + "\n\t")
                                                                default:
                                                                    _ = len(v3)
                                                                }
                                                                var v166 string = ((v5 + v5) + "hello")
                                                                _ = v166
                                                            } else {
                                                                v155 = (^(v4))
                                                            }
                                                        }
                                                        switch v5 {
                                                        case "a":
                                                            switch (("\n\t" + v5) + v5) {
                                                            case "":
                                                                var v167 int = fi136
                                                                _ = v167
                                                            case "a":
                                                                var v168 int32 = -511
                                                                _ = v168
                                                            case "\n\t":
                                                                if len(v3) > 0 {
                                                                    v3[0] = 715
                                                                }
                                                                v3 = append(v3, 8)
                                                                if (32767 != uint16(v155)) {
                                                                    v135 -= 1
                                                                    v5 = (v5 + (v5 + v5))
                                                                    v2 -= (v2 | (^(v6)))
                                                                    var v169 uint16 = 65535
                                                                    _ = v169
                                                                }
                                                            default:
                                                                fmt.Println(v6)
                                                            }
                                                        default:
                                                            fmt.Println(fi136, fi136)
                                                        }
                                                        switch v131 {
                                                        case 1:
                                                            v170 := GenS11{F0: 42, F1: 2147483647}
                                                            _ = v170
                                                        default:
                                                            var v171 uint32 = 4294967295
                                                            _ = v171
                                                        }
                                                        _ = v158
                                                    }
                                                    _ = v156
                                                }
                                                _ = v154
                                                _ = v155
                                            default:
                                                v2 = v2
                                            }
                                        }
                                    }
                                }
                            case -21:
                                var v172 uint32 = 65535
                                _ = v172
                            case 128:
                                fmt.Println(fi136, v128)
                            default:
                                v4 -= (v4 | v4)
                            }
                        }
                        _ = v140
                    }
                    _ = v138
                }
                if (!(520 < 4294967295)) {
                    v131 = (v131 >> ((183) & 63))
                    if true {
                        var v173 uint64 = (v4 - 22)
                        var v174 int8 = genFn1(128)
                        _ = v173
                        _ = v174
                    } else {
                        fmt.Println(v2)
                    }
                    if ((!false) || (4294967295 == 0)) {
                        fmt.Println(v5, v5)
                    }
                } else {
                    if false {
                        _ = len(v3)
                    } else {
                        _ = len(v1)
                        if len(v1) > 0 {
                            v1[0] = 2
                        }
                    }
                    if false {
                        var v175 uint = 526
                        var v176 uint8 = 255
                        _ = v175
                        _ = v176
                    }
                    v5 = v5
                }
            }
            _ = fi136
        }
        _ = fi134
        _ = v135
    }
    for fi177 := 0; fi177 < 15; fi177++ {
        for fi178 := 0; fi178 < 16; fi178++ {
            _ = len(v3)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v5 = v5
            var v179 string = "\n\t"
            var v180 int16 = 106
            _ = fi178
            _ = v179
            _ = v180
        }
        _ = len(v1)
        _ = fi177
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v128:", v128)
    fmt.Println("v131:", v131)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v122
    _ = v128
    _ = v131
    _ = v132
    _ = v133
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
