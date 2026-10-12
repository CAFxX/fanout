package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 bool }
type GenS2 struct { F0 int32; F1 int8; F2 int }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS5 struct { F0 bool; F1 uint16; F2 int8 }
type GenS6 struct { F0 int16; F1 int32; F2 uint64 }
type GenS7 struct { F0 string; F1 int64; F2 uint8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint64; F1 int32 }
type GenS10 struct { F0 int64; F1 string; F2 int32 }
type GenS11 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 int64; F1 int }
type GenS16 struct { F0 int8; F1 uint32; F2 int32 }
type GenS17 struct { F0 int8; F1 int16 }
func genFn0() uint {
    if false {
        var v7 float32 = 1.5
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 int32 = 1
            v7 = (v7 * v7)
            _ = fi8
            _ = v9
        }
        var v10 uint8 = 2
        _ = v10
        _ = v7
    } else {
        for fi11 := 0; fi11 < 19; fi11++ {
            for fi12 := 0; fi12 < 19; fi12++ {
                var v13 int = (fi11 >> ((330) & 63))
                if true {
                    v13 += fi12
                    var v14 uint = 1
                    _ = v14
                }
                v13 += fi12
                _ = fi12
                _ = v13
            }
            switch fi11 {
            case 32767:
                var v15 uint8 = 0
                _ = v15
            case 127:
                v16 := GenS1{F0: 1000, F1: "hello", F2: true}
                _ = v16
            case 32768:
                if false {
                    if false {
                        if (18 > -1) {
                            var v17 float32 = (((1.5 * 1e10) - (float32(-2.25) / float32((-2.25 - 3.14159)))) - 0.1)
                            for fi18 := 0; fi18 < 2; fi18++ {
                                v19 := GenS1{F0: -874, F1: "hello", F2: (0 < 477)}
                                v17 -= (1e10 - (v17 * v17))
                                _ = fi18
                                _ = v19
                            }
                            _ = v17
                        }
                        if (!true) {
                            for fi20 := 0; fi20 < 18; fi20++ {
                                if (!((852 >= 42) || false)) {
                                    var v21 int8 = -98
                                    if true {
                                        fmt.Println(fi20, fi20)
                                        v21 = 100
                                        if ((true && true) || false) {
                                            var v22 int16 = -18
                                            _ = v22
                                        } else {
                                            v23 := []int{(fi11 / ((fi20) | 1)), (fi20 & (fi11 % ((32767) | 1))), (-(fi11)), -2147483648}
                                            if len(v23) >= 2 {
                                                v23 = v23[1:len(v23)]
                                            }
                                            v21 = 42
                                            _ = v23
                                        }
                                    } else {
                                        var v24 int16 = 1000
                                        if (false || (2147483647 < 1)) {
                                            v25 := GenS2{F0: 555, F1: 50, F2: fi20}
                                            v26 := []int{128, 72, fi11}
                                            if (1.5 == (1.5 / (1e10 - 1.5))) {
                                                v26 = append(v26, (fi20 + (fi20 << ((65535) & 63))))
                                                if false {
                                                    v26 = append(v26, fi11)
                                                }
                                                v21 = v21
                                                switch "" {
                                                case "foo bar":
                                                    v27 := GenS3{F0: v21, F1: 0}
                                                    var v28 uint8 = 42
                                                    fmt.Println(v28)
                                                    _ = v27
                                                    _ = v28
                                                case "hello":
                                                    v29 := GenS3{F0: v21, F1: -366}
                                                    if len(v26) > 0 {
                                                        v26[0] = fi20
                                                    }
                                                    var v30 uint32 = 1
                                                    if false {
                                                        var v31 uint16 = 2
                                                        _ = v31
                                                    }
                                                    _ = v29
                                                    _ = v30
                                                case "\n\t":
                                                    if len(v26) >= 2 {
                                                        v26 = v26[1:len(v26)]
                                                    }
                                                default:
                                                    fmt.Println(fi11)
                                                }
                                                if (4294967295 < 1) {
                                                    if ((fi20 / ((42) | 1)) >= fi20) {
                                                        if false {
                                                            v21 = v21
                                                            if false {
                                                                v21 += v21
                                                            } else {
                                                                v32 := GenS1{F0: (-(v24)), F1: "a", F2: (828 > 18446744073709551615)}
                                                                _ = v32
                                                            }
                                                        } else {
                                                            if true {
                                                                var v33 uint64 = 1
                                                                var v34 uint32 = 4294967295
                                                                var v35 uint32 = v34
                                                                _ = v33
                                                                _ = v34
                                                                _ = v35
                                                            } else {
                                                                if ((0.0 * 3.14159) == (0.1 + (0.1 * 1e10))) {
                                                                    fmt.Println(fi11)
                                                                }
                                                            }
                                                            var v36 float64 = 3.14159
                                                            _ = v36
                                                        }
                                                        var v37 float32 = 0.1
                                                        fmt.Println(v24)
                                                        switch fi11 {
                                                        case -622:
                                                            v24 = (v24 ^ (v24 * -1))
                                                        case -2147483648:
                                                            fmt.Println(fi20, v37, fi11)
                                                        default:
                                                            v38 := GenS3{F0: 127, F1: 1}
                                                            _ = v38
                                                        }
                                                        var v39 int8 = 127
                                                        if false {
                                                            var v40 uint16 = 255
                                                            fmt.Println(v39, v40)
                                                            _ = v40
                                                        }
                                                        _ = v37
                                                        _ = v39
                                                    } else {
                                                        v41 := GenS2{F0: -2147483648, F1: v21, F2: (fi11 | 32768)}
                                                        _ = v41
                                                    }
                                                    v21 = (v21 & 56)
                                                }
                                            }
                                            _ = v25
                                            _ = v26
                                        }
                                        _ = v24
                                    }
                                    var v42 float64 = 0.0
                                    _ = v21
                                    _ = v42
                                } else {
                                    v43 := []int{fi20, (fi11 | -746), 9223372036854775807, (^(fi20))}
                                    var v44 int32 = 42
                                    _ = v43
                                    _ = v44
                                }
                                var v45 uint32 = 820
                                v46 := []int{(fi20 | 128)}
                                _ = fi20
                                _ = v45
                                _ = v46
                            }
                            var v47 bool = (!false)
                            _ = v47
                        } else {
                            var v48 uint8 = 2
                            v48 = v48
                            v49 := []int{fi11, fi11, fi11}
                            _ = v48
                            _ = v49
                        }
                        v50 := GenS4{F0: 0, F1: 837, F2: 18446744073709551615}
                        _ = v50
                    } else {
                        if true {
                            v51 := []int{fi11}
                            var v52 float64 = 3.14159
                            _ = v51
                            _ = v52
                        }
                        switch (fi11 | fi11) {
                        case -711:
                            var v53 int16 = 1
                            _ = v53
                        }
                    }
                    var v54 uint32 = 1
                    _ = v54
                }
            default:
                for fi55 := 0; fi55 < 8; fi55++ {
                    switch (fi11 | (-(fi55))) {
                    case 766:
                        if (true && (((-2.25 == 0.1) && true) || (false && false))) {
                            if true {
                                switch (fi11 ^ (-(fi55))) {
                                case 42:
                                    var v56 uint8 = 255
                                    var v57 int32 = -1
                                    _ = v56
                                    _ = v57
                                }
                                var v58 int16 = -123
                                if true {
                                    v59 := []int{fi11}
                                    v58 += (v58 % ((v58) | 1))
                                    v60 := GenS5{F0: (false || true), F1: 726, F2: 100}
                                    _ = v59
                                    _ = v60
                                } else {
                                    var v61 int = (fi11 >> ((278) & 63))
                                    _ = v61
                                }
                                var v62 int32 = 2147483647
                                v63 := []int{}
                                switch (^(fi55)) {
                                case -598:
                                    v64 := GenS6{F0: 1000, F1: v62, F2: 0}
                                    _ = v64
                                case 958:
                                    if true {
                                        fmt.Println(fi11, fi11)
                                    } else {
                                        switch fi55 {
                                        case -9223372036854775808:
                                            v63 = append(v63, (fi55 % ((1) | 1)))
                                        case 9223372036854775807:
                                            if (0 >= 42) {
                                                v65 := GenS7{F0: (("\n\t" + "a") + (("a" + "foo bar") + "")), F1: 0, F2: 79}
                                                _ = v65
                                            }
                                        case 1:
                                            var v66 bool = true
                                            _ = v66
                                        }
                                        v62 -= -275
                                        v67 := GenS5{F0: (!true), F1: 592, F2: 73}
                                        _ = v67
                                    }
                                case 2147483647:
                                    fmt.Println(fi55)
                                default:
                                    if len(v63) >= 2 {
                                        v63 = v63[1:len(v63)]
                                    }
                                }
                                _ = v58
                                _ = v62
                                _ = v63
                            } else {
                                var v68 string = ""
                                var v69 uint32 = 530
                                v70 := []int{(-(fi55)), 90, fi11, 128}
                                _ = v68
                                _ = v69
                                _ = v70
                            }
                            switch (fi55 >> ((0) & 63)) {
                            case 1:
                                v71 := []int{(^(fi11)), 4294967295, 42, 32767}
                                _ = v71
                            default:
                                switch (fi55 | -2147483648) {
                                case 127:
                                    if (812 > 1) {
                                        v72 := GenS8{F0: "a"}
                                        switch "\n\t" {
                                        case "a":
                                            if true {
                                                var v73 int8 = -1
                                                fmt.Println(fi55, fi55)
                                                _ = v73
                                            } else {
                                                if false {
                                                    var v74 float64 = float64(0.0)
                                                    v74 = v74
                                                    v74 = ((3.14159 + -2.25) / v74)
                                                    switch ("" + "\n\t") {
                                                    case "foo bar":
                                                        var v75 int32 = 588
                                                        _ = v75
                                                    case "\n\t":
                                                        v74 = v74
                                                    case "a":
                                                        v76 := GenS9{F0: 1, F1: -2147483648}
                                                        _ = v76
                                                    default:
                                                        v74 = 0.1
                                                    }
                                                    v74 -= (1.5 * ((v74 - (v74 * float64(1e10))) * v74))
                                                    switch fi11 {
                                                    case 9223372036854775807:
                                                        v74 += v74
                                                        v77 := []int{63, fi55}
                                                        var v78 float64 = v74
                                                        _ = v77
                                                        _ = v78
                                                    default:
                                                        v74 -= v74
                                                    }
                                                    var v79 uint32 = 2
                                                    if false {
                                                        if true {
                                                            var v80 uint = 42
                                                            v79 -= 831
                                                            v81 := GenS1{F0: -32768, F1: "hello", F2: (fi11 >= fi11)}
                                                            _ = v80
                                                            _ = v81
                                                        }
                                                        v82 := []int{}
                                                        _ = v82
                                                    } else {
                                                        v79 = v79
                                                        var v83 int16 = 0
                                                        var v84 int32 = 402
                                                        _ = v83
                                                        _ = v84
                                                    }
                                                    if false {
                                                        v79 = 791
                                                        v85 := []int{(fi55 / (((fi11 & 42)) | 1))}
                                                        switch "a" {
                                                        case "foo bar":
                                                            v79 += 2147483647
                                                            var v86 int64 = -478
                                                            _ = v86
                                                        default:
                                                            var v87 uint16 = 32767
                                                            _ = v87
                                                        }
                                                        _ = v85
                                                    }
                                                    if false {
                                                        var v88 uint8 = 128
                                                        if (true && (935 != 1)) {
                                                            fmt.Println(v74, v79)
                                                            var v89 int64 = 1
                                                            if ((^(v88)) != 10) {
                                                                v79 += 0
                                                                fmt.Println(fi55)
                                                            }
                                                            _ = v89
                                                        } else {
                                                            v90 := GenS4{F0: 283, F1: v79, F2: 1}
                                                            _ = v90
                                                        }
                                                        _ = v88
                                                    } else {
                                                        v91 := GenS5{F0: (2147483647 <= 42), F1: 32767, F2: -128}
                                                        var v92 int32 = 42
                                                        v79 += 65535
                                                        _ = v91
                                                        _ = v92
                                                    }
                                                    _ = v74
                                                    _ = v79
                                                }
                                                if true {
                                                    v93 := GenS8{F0: ("hello" + (("" + ("foo bar" + "a")) + "hello"))}
                                                    v94 := GenS3{F0: -98, F1: 1}
                                                    _ = v93
                                                    _ = v94
                                                }
                                            }
                                        case "foo bar":
                                            var v95 string = ""
                                            _ = v95
                                        case "":
                                            v96 := GenS4{F0: 2, F1: 1, F2: 42}
                                            _ = v96
                                        }
                                        _ = v72
                                    }
                                case 2147483647:
                                    switch (-(fi55)) {
                                    case 32767:
                                        switch "" {
                                        case "foo bar":
                                            v97 := GenS7{F0: ("hello" + "\n\t"), F1: 572, F2: 127}
                                            _ = v97
                                        case "hello":
                                            var v98 float32 = (0.1 * (float32(0.0) + (3.14159 + float32((1.5 / 3.14159)))))
                                            _ = v98
                                        default:
                                            var v99 int64 = -761
                                            _ = v99
                                        }
                                    case 651:
                                        if ((790 >= 2) && (-128 == -1)) {
                                            v100 := GenS10{F0: 9223372036854775807, F1: "", F2: 932}
                                            var v101 int = (-(fi11))
                                            _ = v100
                                            _ = v101
                                        }
                                        if true {
                                            v102 := GenS7{F0: "a", F1: int64(fi55), F2: 77}
                                            var v103 uint64 = 2
                                            _ = v102
                                            _ = v103
                                        } else {
                                            var v104 int32 = -319
                                            _ = v104
                                        }
                                        v105 := []int{42, fi55}
                                        _ = v105
                                    case 0:
                                        switch "" {
                                        case "":
                                            var v106 float32 = 1e10
                                            _ = v106
                                        }
                                    default:
                                        var v107 uint = 33
                                        _ = v107
                                    }
                                case -2147483648:
                                    var v108 uint16 = 0
                                    _ = v108
                                }
                            }
                        } else {
                            if false {
                                switch (fi55 >> ((65535) & 63)) {
                                case 42:
                                    v109 := []int{(-(fi11)), 0, (fi55 + (fi11 >> ((42) & 63))), 127}
                                    _ = v109
                                case 128:
                                    var v110 uint32 = 613
                                    var v111 int32 = 651
                                    v111 -= v111
                                    _ = v110
                                    _ = v111
                                case 1:
                                    var v112 int64 = 235
                                    _ = v112
                                }
                                var v113 float32 = 1.5
                                _ = v113
                            } else {
                                switch ((("a" + ("hello" + "foo bar")) + "") + "") {
                                case "hello":
                                    var v114 int64 = 42
                                    _ = v114
                                }
                            }
                        }
                    default:
                        if false {
                            var v115 int16 = -1
                            v116 := []int{0, (fi11 | fi55), fi55}
                            v115 = -1
                            if len(v116) > 0 {
                                v116[0] = (-(fi55))
                            }
                            _ = v115
                            _ = v116
                        }
                    }
                    v117 := []int{fi11, fi11, fi55}
                    switch "\n\t" {
                    case "hello":
                        switch "foo bar" {
                        case "foo bar":
                            var v118 bool = true
                            _ = v118
                        case "hello":
                            if (-2147483648 <= 0) {
                                if true {
                                    var v119 uint32 = 1
                                    fmt.Println(fi11, fi55)
                                    switch (^(fi55)) {
                                    case 9223372036854775807:
                                        if ((-(fi55)) <= 1) {
                                            var v120 int64 = 1
                                            var v121 int16 = 1
                                            v120 += v120
                                            _ = v120
                                            _ = v121
                                        } else {
                                            v119 -= v119
                                        }
                                    }
                                    v119 -= (v119 - v119)
                                    var v122 uint32 = v119
                                    if len(v117) > 0 {
                                        v117[0] = (fi11 % ((fi55) | 1))
                                    }
                                    if len(v117) > 0 {
                                        v117[0] = fi11
                                    }
                                    _ = v119
                                    _ = v122
                                }
                                if false {
                                    switch 32767 {
                                    case 127:
                                        if true {
                                            fmt.Println(fi55, fi11)
                                        } else {
                                            switch ("" + "") {
                                            case "":
                                                if true {
                                                    var v123 uint64 = 826
                                                    v123 += v123
                                                    _ = v123
                                                }
                                            case "hello":
                                                if true {
                                                    fmt.Println(fi55, fi11)
                                                    var v124 int64 = -949
                                                    _ = v124
                                                }
                                            default:
                                                var v125 string = (("a" + "") + "\n\t")
                                                _ = v125
                                            }
                                        }
                                    }
                                } else {
                                    if len(v117) > 0 {
                                        v117[0] = 128
                                    }
                                }
                            }
                            var v126 int32 = 695
                            _ = v126
                        }
                    case "foo bar":
                        fmt.Println(fi11)
                    }
                    _ = fi55
                    _ = v117
                }
            }
            _ = fi11
        }
        v127 := []int{}
        _ = v127
    }
    var v128 uint8 = 219
    if true {
        if (-292 == 32767) {
            fmt.Println(v128)
            switch "foo bar" {
            case "\n\t":
                v129 := GenS10{F0: 9223372036854775807, F1: ("" + "\n\t"), F2: 464}
                _ = v129
            default:
                v128 = 255
            }
        } else {
            v130 := GenS5{F0: (1 == 42), F1: 1, F2: 93}
            switch "foo bar" {
            case "a":
                v131 := GenS5{F0: (336 < -1), F1: 42, F2: 127}
                _ = v131
            }
            _ = v130
        }
        v132 := GenS10{F0: -879, F1: "\n\t", F2: -629}
        if true {
            v128 = (v128 ^ 128)
            v128 = (v128 >> ((42) & 63))
            v133 := GenS6{F0: -1, F1: -707, F2: 1}
            _ = v133
        } else {
            v134 := GenS9{F0: 0, F1: 42}
            for fi135 := 0; fi135 < 5; fi135++ {
                v128 = (v128 | (v128 % ((v128) | 1)))
                switch fi135 {
                case -2147483648:
                    var v136 uint8 = (^(v128))
                    _ = v136
                case 32768:
                    v128 += v128
                case 120:
                    v128 = v128
                default:
                    v137 := []int{fi135, (fi135 >> ((42) & 63))}
                    _ = v137
                }
                if true {
                    v138 := []int{(fi135 / ((1) | 1)), 9223372036854775807, 127}
                    _ = v138
                }
                _ = fi135
            }
            var v139 int32 = 299
            _ = v134
            _ = v139
        }
        _ = v132
    }
    if false {
        v128 += 122
        var v140 int32 = 42
        for fi141 := 0; fi141 < 19; fi141++ {
            v142 := []int{fi141}
            v143 := GenS11{F0: 2, F1: (!false), F2: 42}
            if true {
                var v144 float64 = ((float64(float32((0.0 + ((float64(((-2.25 / ((0.0 / 3.14159) - -2.25)) - (0.0 / -2.25))) - 0.0) + (1e10 / 1.0))))) * 3.14159) - (((((3.14159 / 1.5) - 3.14159) + -2.25) / ((0.0 * 0.0) + (1e10 / 1.0))) * 3.14159))
                v144 += (((v144 * float64(0.1)) + v144) / ((v144 + (v144 / v144)) * v144))
                if true {
                    if len(v142) >= 2 {
                        v142 = v142[1:len(v142)]
                    }
                    if (true && (-9223372036854775808 >= -566)) {
                        if false {
                            v145 := GenS3{F0: -99, F1: 1}
                            _ = v145
                        } else {
                            v144 += v144
                        }
                        v144 -= v144
                    }
                    switch fi141 {
                    case 4294967295:
                        for fi146 := 0; fi146 < 16; fi146++ {
                            v144 = ((v144 - -2.25) / (v144 / ((-2.25 / 3.14159) - v144)))
                            var v147 uint32 = 1
                            _ = fi146
                            _ = v147
                        }
                    }
                } else {
                    var v148 int = (fi141 | 0)
                    fmt.Println(v140, v128, v140)
                    _ = v148
                }
                _ = v144
            } else {
                v140 = v140
            }
            _ = fi141
            _ = v142
            _ = v143
        }
        _ = v140
    } else {
        switch ("foo bar" + "hello") {
        case "":
            v128 -= 128
        case "foo bar":
            var v149 bool = (18446744073709551615 <= 784)
            _ = v149
        case "hello":
            v150 := GenS2{F0: 994, F1: 100, F2: 32767}
            _ = v150
        }
        var v151 int8 = -1
        _ = v151
    }
    _ = v128
    return 18446744073709551615
}
func genFn1(p0 uint64, p1 bool) int {
    var v152 float32 = (((3.14159 - ((0.0 - 1.5) / 0.1)) + 3.14159) - (1e10 - 1e10))
    v153 := []int{-9223372036854775808, 309, 1, -9223372036854775808}
    v153 = append(v153, 42)
    var v154 int16 = -1
    switch -620 {
    case -1:
        fmt.Println(v154, p1, v152)
    case 60:
        var v155 int64 = -785
        _ = v155
    case 32768:
        var v156 uint16 = 0
        _ = v156
    }
    _ = p0
    _ = p1
    _ = v152
    _ = v153
    _ = v154
    return -2147483648
}

func main() {
    var v1 int = 696
    v2 := []uint32{0}
    v3 := []int{(v1 ^ v1)}
    v4 := make(map[int]string)
    var v5 float64 = 0.0
    var v6 uint = 18446744073709551615
    var v157 uint64 = 4294967295
    var v158 int = genFn1((v157 << ((65535) & 63)), ((9223372036854775807 >= 2147483647) && (42 <= -1)))
    v6 += uint(v157)
    fmt.Println(v157, v158, v158)
    var v159 int64 = int64(v157)
    if (!true) {
        var v160 float64 = v5
        _ = len(v4)
        if (v159 < (^(v159))) {
            v159 -= -1
            var v161 int16 = 0
            var v162 uint = v6
            _ = v161
            _ = v162
        }
        _ = v160
    } else {
        var v163 int64 = 1
        v159 = -1
        v164 := GenS12{F0: 0}
        _ = v163
        _ = v164
    }
    switch "" {
    case "a":
        fmt.Println(v158)
    case "foo bar":
        var v165 int32 = 1
        _ = v165
    default:
        v159 -= (v159 ^ v159)
    }
    fmt.Println(v1, v1, v159)
    _ = len(v2)
    v1 += genFn1((v157 - v157), true)
    delete(v4, 32768)
    switch v1 {
    case -295:
        fmt.Println(v5, v158)
    case 127:
        var v166 int32 = 2147483647
        _ = v166
    case -1:
        for fi167 := 0; fi167 < 4; fi167++ {
            switch "a" {
            case "\n\t":
                _ = len(v4)
            }
            v158 -= (v1 % ((v158) | 1))
            v168 := GenS13{F0: true}
            _ = fi167
            _ = v168
        }
    default:
        delete(v4, (v158 * (v1 * 9223372036854775807)))
    }
    var v169 int = 32767
    switch v169 {
    case 378:
        switch v158 {
        case 0:
            for fi170 := 0; fi170 < 12; fi170++ {
                if false {
                    switch ("a" + "\n\t") {
                    case "foo bar":
                        switch v169 {
                        case 4294967295:
                            fmt.Println(v1)
                        case 0:
                            for fi171 := 0; fi171 < 13; fi171++ {
                                v172 := GenS2{F0: 219, F1: -62, F2: fi170}
                                fmt.Println(v158)
                                if true {
                                    switch "foo bar" {
                                    case "hello":
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    case "a":
                                        delete(v4, v169)
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = v158
                                    }
                                }
                                _ = fi171
                                _ = v172
                            }
                        }
                    default:
                        v158 -= (^(v158))
                    }
                    for fi173 := 0; fi173 < 5; fi173++ {
                        v174 := GenS4{F0: v157, F1: 2147483647, F2: 18446744073709551615}
                        if false {
                            v1 = fi173
                        } else {
                            if (false || true) {
                                var v175 string = "a"
                                fmt.Println(v1, v157)
                                _ = v175
                            } else {
                                if (true && false) {
                                    var v176 int32 = 0
                                    _ = v176
                                } else {
                                    switch (fi173 | (^(fi173))) {
                                    case -985:
                                        switch "foo bar" {
                                        case "hello":
                                            fmt.Println(v159, fi173)
                                            delete(v4, 572)
                                            var v177 string = "foo bar"
                                            var v178 bool = (((-32768 < 0) || ((true && ((float32(1.5) * 1.5) == 0.1)) && false)) && false)
                                            v179 := GenS6{F0: 1, F1: 819, F2: (v157 | v157)}
                                            var v180 bool = ((v6 | 773) != v6)
                                            fmt.Println(v169)
                                            _ = v177
                                            _ = v178
                                            _ = v179
                                            _ = v180
                                        }
                                    default:
                                        var v181 int64 = 250
                                        _ = v181
                                    }
                                    v5 = -2.25
                                }
                                _ = len(v4)
                            }
                            var v182 int64 = v159
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v182
                        }
                        v183 := v4[fi173]
                        _ = fi173
                        _ = v174
                        _ = v183
                    }
                }
                v184 := GenS14{F0: int32(v6)}
                _ = fi170
                _ = v184
            }
        default:
            v157 -= (^(v157))
        }
    case 695:
        for fi185 := 0; fi185 < 14; fi185++ {
            _ = len(v4)
            if (-81 > 32767) {
                if false {
                    for fi186 := 0; fi186 < 15; fi186++ {
                        v187 := GenS15{F0: 9223372036854775807, F1: v169}
                        var v188 string = ("a" + "foo bar")
                        _ = fi186
                        _ = v187
                        _ = v188
                    }
                    v6 = 255
                    if true {
                        var v189 uint8 = 255
                        _ = v189
                    }
                } else {
                    v158 = (v1 + (v158 << ((1) & 63)))
                    for fi190 := 0; fi190 < 14; fi190++ {
                        var v191 uint64 = v157
                        switch "" {
                        case "\n\t":
                            v192 := GenS5{F0: true, F1: 545, F2: 100}
                            _ = v192
                        case "":
                            if (((-(v158)) < fi185) && false) {
                                var v193 float32 = (float32((1e10 + float64(0.0))) / ((0.0 - float32(1.5)) / float32((float64(0.0) / v5))))
                                _ = v193
                            } else {
                                v159 = (v159 >> ((v6) & 63))
                                fmt.Println(v157, v157)
                            }
                        case "a":
                            if true {
                                var v194 int = genFn1(v191, false)
                                if true {
                                    v191 = v191
                                    v158 = v194
                                    fmt.Println(v1)
                                } else {
                                    if false {
                                        if (true || false) {
                                            var v195 float64 = v5
                                            var v196 uint32 = uint32(v191)
                                            _ = v195
                                            _ = v196
                                        } else {
                                            fmt.Println(v6, fi190, v159)
                                            if (0.0 > (3.14159 - -2.25)) {
                                                switch (fi185 & 9223372036854775807) {
                                                case 0:
                                                    if (false && (v159 >= int64(v191))) {
                                                        v197 := GenS15{F0: int64(v194), F1: fi185}
                                                        fmt.Println(v194)
                                                        v198 := GenS16{F0: 127, F1: 2147483647, F2: 0}
                                                        _ = v197
                                                        _ = v198
                                                    } else {
                                                        fmt.Println(v6, v159, v1)
                                                        v199 := GenS8{F0: "\n\t"}
                                                        fmt.Println(v159, v6)
                                                        _ = v199
                                                    }
                                                case -1:
                                                    v200 := GenS17{F0: -94, F1: -958}
                                                    _ = v200
                                                }
                                                fmt.Println(fi190, v159, fi190)
                                                v3 = append(v3, (v158 * v1))
                                            }
                                        }
                                        _ = len(v3)
                                        v158 -= 632
                                    } else {
                                        v194 += int(v6)
                                        if true {
                                            fmt.Println(v159)
                                            var v201 uint16 = 65535
                                            _ = v201
                                        }
                                    }
                                }
                                fmt.Println(fi185, fi190, fi190)
                                _ = v194
                            }
                        default:
                            var v202 string = "foo bar"
                            _ = v202
                        }
                        _ = fi190
                        _ = v191
                    }
                    _ = len(v4)
                }
                v169 += int(v159)
            } else {
                v157 -= (v157 ^ (v157 | v157))
                v159 += v159
                switch ("hello" + ("foo bar" + (("hello" + "") + (("\n\t" + "") + "\n\t")))) {
                case "":
                    v6 += genFn0()
                case "a":
                    var v203 int64 = v159
                    _ = v203
                }
            }
            var v204 int8 = 100
            _ = fi185
            _ = v204
        }
    default:
        for fi205 := 0; fi205 < 1; fi205++ {
            v206 := GenS17{F0: 0, F1: int16(v159)}
            _ = fi205
            _ = v206
        }
    }
    for fi207 := 0; fi207 < 13; fi207++ {
        v2 = append(v2, 1)
        _ = fi207
    }
    fmt.Println(v1, v157)
    fmt.Println("v1:", v1)
    fmt.Println("v157:", v157)
    fmt.Println("v158:", v158)
    fmt.Println("v159:", v159)
    fmt.Println("v169:", v169)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v157
    _ = v158
    _ = v159
    _ = v169
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
