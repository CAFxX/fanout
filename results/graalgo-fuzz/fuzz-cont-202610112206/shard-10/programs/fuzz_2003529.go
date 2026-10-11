package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 string }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int8; F1 string; F2 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16; F1 uint16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint32; F1 int }
type GenS10 struct { F0 int32; F1 int }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 string; F1 int8 }
type GenS15 struct { F0 bool; F1 uint }
type GenS16 struct { F0 int64; F1 int }
type GenS17 struct { F0 int32; F1 uint32 }
type GenS18 struct { F0 int32; F1 int16 }

func main() {
    var v1 uint = 453
    v2 := make(map[uint32]int)
    var v3 float64 = (3.14159 / 0.1)
    var v4 uint16 = 42
    v5 := []int{-9223372036854775808}
    if (42 > 42) {
        switch 469 {
        case 4294967295:
            var v6 int16 = 713
            _ = v6
        case -1:
            var v7 string = "hello"
            _ = v7
        default:
            switch 1 {
            case 2147483647:
                delete(v2, 4294967295)
            case 429:
                v8 := GenS1{F0: 157, F1: 1, F2: ("\n\t" + "foo bar")}
                _ = v8
            }
        }
        v4 = v4
        var v9 float64 = (0.1 / v3)
        _ = v9
    } else {
        if (-32768 != 681) {
            v10 := GenS2{F0: 4294967295, F1: "hello"}
            if false {
                var v11 uint32 = 0
                fmt.Println(v1)
                _ = v11
            } else {
                var v12 float32 = (((1.5 * 0.1) - 0.1) * 3.14159)
                switch "a" {
                case "hello":
                    switch "" {
                    case "":
                        v13 := GenS3{F0: 100, F1: "\n\t", F2: 42}
                        _ = v13
                    default:
                        v2[586] = 237
                    }
                default:
                    var v14 bool = false
                    _ = v14
                }
                _ = v12
            }
            for fi15 := 0; fi15 < 10; fi15++ {
                fmt.Println(v4)
                _ = fi15
            }
            _ = v10
        } else {
            switch 4294967295 {
            case 709:
                v16 := GenS1{F0: 2, F1: 27, F2: ""}
                _ = v16
            case -1:
                fmt.Println(v4)
            default:
                var v17 string = ("\n\t" + "hello")
                _ = v17
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        if (-129 > -586) {
            for fi18 := 0; fi18 < 1; fi18++ {
                fmt.Println(v3, v1)
                _ = fi18
            }
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    var v19 string = "a"
    var v20 float64 = v3
    switch -1 {
    case 4294967295:
        var v21 int8 = 13
        _ = v21
    case -2147483648:
        var v22 uint16 = (v4 - 0)
        _ = v22
    default:
        var v23 float64 = 0.0
        _ = v23
    }
    if (((true || false) || ((-173 < -135) || false)) || (!(-325 >= 5))) {
        fmt.Println(v19)
        v24 := GenS4{F0: 252}
        _ = v24
    }
    fmt.Println(v3, v1, v19)
    if true {
        switch v19 {
        case "":
            for fi25 := 0; fi25 < 16; fi25++ {
                fmt.Println(fi25, v20, v19)
                fmt.Println(v20)
                if len(v5) > 0 {
                    v5[0] = fi25
                }
                _ = fi25
            }
        case "hello":
            v3 -= 0.1
        case "a":
            for fi26 := 0; fi26 < 10; fi26++ {
                if (fi26 == fi26) {
                    for fi27 := 0; fi27 < 12; fi27++ {
                        var v28 uint8 = 1
                        _ = fi27
                        _ = v28
                    }
                } else {
                    _ = len(v2)
                    if (19 <= v4) {
                        if true {
                            v3 = v3
                            var v29 uint16 = (^(v4))
                            for fi30 := 0; fi30 < 5; fi30++ {
                                var v31 uint64 = 831
                                v1 += 255
                                _ = fi30
                                _ = v31
                            }
                            _ = v29
                        } else {
                            fmt.Println(v3)
                            v32 := GenS5{F0: 32767, F1: 42}
                            var v33 string = (v19 + "hello")
                            _ = v32
                            _ = v33
                        }
                        if true {
                            fmt.Println(v3, v3, v3)
                        }
                        switch v19 {
                        case "":
                            for fi34 := 0; fi34 < 18; fi34++ {
                                if (false && ("a" == (v19 + ("hello" + v19)))) {
                                    if false {
                                        if (0.0 > float32(v3)) {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        }
                                    } else {
                                        if (false && ((32767 < -737) && (!(993 >= 2147483648)))) {
                                            v35 := GenS5{F0: 1000, F1: 250}
                                            _ = v35
                                        }
                                        var v36 int = fi34
                                        var v37 int16 = -258
                                        _ = v36
                                        _ = v37
                                    }
                                    v3 -= -2.25
                                    v38 := GenS3{F0: 100, F1: "\n\t", F2: 229}
                                    _ = v38
                                }
                                _ = fi34
                            }
                        case "foo bar":
                            var v39 float64 = 1.5
                            _ = v39
                        case "a":
                            var v40 uint = (v1 << (((v1 * 1)) & 63))
                            _ = v40
                        }
                    } else {
                        var v41 uint16 = v4
                        for fi42 := 0; fi42 < 16; fi42++ {
                            switch "a" {
                            case "hello":
                                v43 := GenS5{F0: 32767, F1: 255}
                                _ = v43
                            case "":
                                v41 = v4
                            case "\n\t":
                                switch fi42 {
                                case -736:
                                    v19 = "hello"
                                    switch fi26 {
                                    case 32768:
                                        v1 = 1
                                    case -849:
                                        switch (v19 + ("a" + v19)) {
                                        case "\n\t":
                                            fmt.Println(v3)
                                        case "foo bar":
                                            switch 612 {
                                            case 530:
                                                v4 = v4
                                            case 533:
                                                v4 = (v41 | v41)
                                            default:
                                                if len(v5) > 0 {
                                                    v5[0] = (fi26 / ((2147483647) | 1))
                                                }
                                            }
                                        default:
                                            v41 += 2
                                            var v44 uint16 = 0
                                            var v45 uint = uint(v44)
                                            _ = v44
                                            _ = v45
                                        }
                                    case 42:
                                        if (false || false) {
                                            v4 += v41
                                            if true {
                                                v41 -= v4
                                                if (!true) {
                                                    v3 = v3
                                                }
                                                var v46 int32 = -1
                                                _ = v46
                                            }
                                            v41 -= (v4 ^ v4)
                                        } else {
                                            v47 := v2[993]
                                            v48 := v2[2147483647]
                                            _ = v47
                                            _ = v48
                                        }
                                    default:
                                        if (v1 > 65535) {
                                            v49 := GenS1{F0: 255, F1: 1, F2: v19}
                                            v50 := GenS5{F0: -1, F1: (v4 ^ (v4 - (v41 << (((v1 - (v1 * v1))) & 63))))}
                                            switch 4294967295 {
                                            case -1:
                                                if false {
                                                    v51 := GenS1{F0: 1, F1: int8(v1), F2: v19}
                                                    var v52 uint8 = 0
                                                    v3 = ((v3 * v3) + float64(0.0))
                                                    _ = v51
                                                    _ = v52
                                                }
                                                var v53 float64 = v20
                                                _ = v53
                                            case 127:
                                                if false {
                                                    var v54 bool = (true || (false && false))
                                                    v4 = (^(v4))
                                                    var v55 uint8 = 42
                                                    _ = v54
                                                    _ = v55
                                                }
                                            case -885:
                                                var v56 uint16 = (v41 ^ (v41 ^ v4))
                                                _ = v56
                                            default:
                                                if (-262 != 832) {
                                                    if (true || (0 > -1)) {
                                                        v19 = "\n\t"
                                                        var v57 float32 = 1e10
                                                        _ = v57
                                                    } else {
                                                        var v58 int8 = -1
                                                        _ = v58
                                                    }
                                                }
                                            }
                                            _ = v49
                                            _ = v50
                                        }
                                    }
                                }
                            default:
                                switch fi42 {
                                case 32768:
                                    if (-2.25 != (((0.1 - (float64(-2.25) * v20)) - v20) + 1e10)) {
                                        var v59 float32 = 1.5
                                        switch (fi26 << ((255) & 63)) {
                                        case 127:
                                            var v60 uint8 = 255
                                            _ = v60
                                        }
                                        _ = v59
                                    }
                                case -767:
                                    v61 := GenS6{F0: true}
                                    _ = v61
                                }
                            }
                            v62 := GenS5{F0: -1, F1: (v41 >> ((2) & 63))}
                            var v63 uint32 = 1
                            _ = fi42
                            _ = v62
                            _ = v63
                        }
                        _ = v41
                    }
                    _ = len(v2)
                }
                var v64 int16 = -1
                if false {
                    for fi65 := 0; fi65 < 2; fi65++ {
                        var v66 int64 = -1
                        _ = fi65
                        _ = v66
                    }
                    for fi67 := 0; fi67 < 10; fi67++ {
                        if true {
                            var v68 string = v19
                            v64 += (v64 / ((v64) | 1))
                            v5 = append(v5, fi26)
                            var v69 uint8 = 42
                            _ = v68
                            _ = v69
                        } else {
                            var v70 uint64 = 2
                            if ((!(1e10 <= 0.0)) && (false || false)) {
                                switch fi26 {
                                case 734:
                                    var v71 bool = true
                                    _ = v71
                                case -9223372036854775808:
                                    if false {
                                        v64 -= (v64 + (v64 << ((v1) & 63)))
                                        switch (^(fi26)) {
                                        case -9223372036854775808:
                                            switch -9223372036854775808 {
                                            case 32767:
                                                v72 := GenS7{F0: fi67}
                                                _ = v72
                                            }
                                        case 979:
                                            v73 := GenS7{F0: int(v1)}
                                            _ = v73
                                        case 42:
                                            v64 = (-(v64))
                                        }
                                        v74 := GenS4{F0: 2}
                                        v75 := GenS8{F0: 13}
                                        var v76 int16 = 1000
                                        _ = v74
                                        _ = v75
                                        _ = v76
                                    }
                                    var v77 float32 = 1e10
                                    switch (fi26 ^ 1) {
                                    case 4294967295:
                                        v3 += -2.25
                                    default:
                                        if ((!(1 <= -977)) && false) {
                                            var v78 int64 = -9223372036854775808
                                            if true {
                                                if true {
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                } else {
                                                    if (65535 != 1) {
                                                        fmt.Println(v70, fi67, v19)
                                                        v64 += (v64 << ((639) & 63))
                                                    } else {
                                                        v79 := GenS1{F0: 42, F1: 99, F2: (v19 + (v19 + ""))}
                                                        _ = v79
                                                    }
                                                    v78 -= v78
                                                    v80 := GenS9{F0: 2, F1: fi67}
                                                    _ = v80
                                                }
                                            }
                                            var v81 float32 = float32(float64(((v77 - 0.0) - (float32(v3) / v77))))
                                            _ = v78
                                            _ = v81
                                        } else {
                                            v82 := GenS7{F0: (-(fi26))}
                                            _ = v82
                                        }
                                    }
                                    _ = v77
                                case 42:
                                    var v83 int8 = 127
                                    _ = v83
                                }
                            } else {
                                switch (fi67 / ((fi26) | 1)) {
                                case 128:
                                    v64 -= v64
                                default:
                                    v84 := GenS6{F0: true}
                                    v85 := GenS3{F0: 0, F1: (v19 + ""), F2: (^(v70))}
                                    _ = v84
                                    _ = v85
                                }
                                var v86 float64 = (float64(float32(v20)) + (v3 / 1.0))
                                _ = v86
                            }
                            _ = v70
                        }
                        fmt.Println(fi26)
                        switch (fi67 >> (((v1 * v1)) & 63)) {
                        case 2147483647:
                            fmt.Println(v64, v20)
                        case -9223372036854775808:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        case 752:
                            switch (((("a" + v19) + (("\n\t" + (v19 + v19)) + "foo bar")) + "") + v19) {
                            case "foo bar":
                                if len(v5) > 0 {
                                    v5[0] = (^(fi26))
                                }
                            case "":
                                v87 := GenS7{F0: 548}
                                _ = v87
                            case "a":
                                var v88 string = v19
                                _ = v88
                            }
                        }
                        v89 := GenS10{F0: 42, F1: (fi67 % ((fi26) | 1))}
                        v5 = append(v5, fi67)
                        var v90 int8 = 42
                        _ = fi67
                        _ = v89
                        _ = v90
                    }
                    var v91 int8 = 42
                    _ = v91
                } else {
                    _ = len(v2)
                }
                _ = fi26
                _ = v64
            }
        }
        v19 = (v19 + v19)
        v20 += v3
    } else {
        if true {
            v92 := GenS11{F0: 0}
            if false {
                if true {
                    var v93 int16 = 771
                    _ = v93
                } else {
                    for fi94 := 0; fi94 < 1; fi94++ {
                        var v95 int16 = 1000
                        _ = len(v2)
                        _ = fi94
                        _ = v95
                    }
                    var v96 int = int(v4)
                    _ = v96
                }
                if true {
                    v97 := GenS12{F0: 759}
                    for fi98 := 0; fi98 < 9; fi98++ {
                        for fi99 := 0; fi99 < 8; fi99++ {
                            if false {
                                fmt.Println(v19)
                                var v100 uint32 = uint32(v4)
                                _ = v100
                            } else {
                                v101 := GenS9{F0: 465, F1: (^(fi98))}
                                v2[719] = fi99
                                _ = v101
                            }
                            var v102 uint = v1
                            var v103 bool = (82 >= 2)
                            _ = fi99
                            _ = v102
                            _ = v103
                        }
                        for fi104 := 0; fi104 < 14; fi104++ {
                            v20 = v20
                            v105 := GenS4{F0: 127}
                            _ = fi104
                            _ = v105
                        }
                        _ = fi98
                    }
                    var v106 int16 = 0
                    _ = v106
                    _ = v97
                }
                if true {
                    if true {
                        var v107 bool = false
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v107
                    } else {
                        var v108 uint = v1
                        switch -902 {
                        case 2147483647:
                            v109 := GenS13{F0: 1}
                            _ = v109
                        default:
                            v20 = 0.1
                        }
                        if (!((!false) && false)) {
                            v2[65535] = -2147483648
                        } else {
                            if len(v5) > 0 {
                                v5[0] = 32768
                            }
                            _ = len(v5)
                            if true {
                                switch 1 {
                                case 32768:
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                case 127:
                                    if false {
                                        v110 := GenS3{F0: 100, F1: (("foo bar" + ("hello" + ("hello" + v19))) + v19), F2: 2}
                                        v19 = v19
                                        for fi111 := 0; fi111 < 16; fi111++ {
                                            for fi112 := 0; fi112 < 17; fi112++ {
                                                switch ((("foo bar" + ((v19 + ("a" + v19)) + v19)) + v19) + "hello") {
                                                case "":
                                                    v4 = 913
                                                case "\n\t":
                                                    delete(v2, 65535)
                                                default:
                                                    v113 := GenS14{F0: "", F1: 100}
                                                    _ = v113
                                                }
                                                _ = fi112
                                            }
                                            var v114 int64 = -996
                                            if ((^(v1)) != (v108 * (v1 + v108))) {
                                                var v115 uint8 = 42
                                                _ = v115
                                            } else {
                                                v114 = v114
                                            }
                                            _ = fi111
                                            _ = v114
                                        }
                                        _ = v110
                                    } else {
                                        v20 -= 1.5
                                        var v116 string = (v19 + "foo bar")
                                        _ = v116
                                    }
                                }
                                v108 = 18446744073709551615
                            } else {
                                switch 268 {
                                case 42:
                                    for fi117 := 0; fi117 < 17; fi117++ {
                                        fmt.Println(v20, v19)
                                        for fi118 := 0; fi118 < 10; fi118++ {
                                            fmt.Println(v3, v4)
                                            _ = fi118
                                        }
                                        var v119 string = "\n\t"
                                        _ = fi117
                                        _ = v119
                                    }
                                case 851:
                                    v1 = 255
                                case 1:
                                    var v120 float32 = (1e10 + (0.0 / ((1e10 + 1.5) * -2.25)))
                                    _ = v120
                                default:
                                    var v121 int32 = 1
                                    _ = v121
                                }
                                if len(v5) > 0 {
                                    v5[0] = int(v1)
                                }
                            }
                        }
                        _ = v108
                    }
                    var v122 int8 = -1
                    for fi123 := 0; fi123 < 6; fi123++ {
                        if (286 == 1) {
                            var v124 string = v19
                            if false {
                                v19 = v124
                                for fi125 := 0; fi125 < 3; fi125++ {
                                    var v126 uint = 723
                                    v20 += v3
                                    var v127 int8 = v122
                                    _ = fi125
                                    _ = v126
                                    _ = v127
                                }
                                v4 += (^(v4))
                            }
                            v20 = ((float64(float32(v3)) * v3) * (v3 / v20))
                            _ = v124
                        } else {
                            v128 := v2[uint32(v4)]
                            var v129 uint16 = v4
                            _ = v128
                            _ = v129
                        }
                        _ = fi123
                    }
                    _ = v122
                } else {
                    v130 := GenS7{F0: -2147483648}
                    _ = v130
                }
            } else {
                if false {
                    var v131 string = v19
                    _ = v131
                } else {
                    v20 += ((0.1 / float64(1.5)) - 3.14159)
                    v4 = v4
                    v2[1] = 259
                }
                v3 += v3
            }
            if ((!false) || (v4 >= (^(v4)))) {
                v132 := GenS15{F0: ((-32768 != 384) && (-32768 < 1)), F1: 1}
                for fi133 := 0; fi133 < 6; fi133++ {
                    v19 = v19
                    _ = fi133
                }
                _ = v132
            }
            _ = v92
        }
        v134 := GenS9{F0: 65535, F1: 9223372036854775807}
        for fi135 := 0; fi135 < 12; fi135++ {
            v19 = v19
            if (false && false) {
                v136 := GenS16{F0: 0, F1: (fi135 >> ((255) & 63))}
                v3 += (((((v20 / 0.1) + v3) * 0.1) - (v3 + ((0.1 * v3) + 0.0))) / (1.5 - ((v20 * 0.0) + (v20 / v3))))
                _ = v136
            }
            v137 := GenS17{F0: 1, F1: 65535}
            _ = fi135
            _ = v137
        }
        _ = v134
    }
    v4 += v4
    for fi138 := 0; fi138 < 12; fi138++ {
        for fi139 := 0; fi139 < 19; fi139++ {
            if true {
                var v140 uint32 = 2147483648
                var v141 int16 = 42
                v141 -= (v141 * v141)
                _ = v140
                _ = v141
            } else {
                switch ("" + (v19 + v19)) {
                case "a":
                    v20 += (v20 - (v20 + (float64((0.1 - float32(float64(-2.25)))) / ((float64(3.14159) + (v20 * ((1.5 + (0.1 / (v20 - v20))) - (v20 - (v20 + v3))))) - 1.5))))
                case "foo bar":
                    switch -762 {
                    case 127:
                        v1 = (v1 | uint(v4))
                    case 304:
                        var v142 uint16 = v4
                        _ = v142
                    case -9223372036854775808:
                        var v143 bool = ((!(((-2147483648 >= 0) && false) || true)) && ((!false) && true))
                        _ = v143
                    default:
                        v4 -= (v4 * (v4 - (v4 << ((65535) & 63))))
                    }
                default:
                    var v144 uint32 = 2147483647
                    _ = v144
                }
                var v145 uint8 = 255
                _ = v145
            }
            var v146 int64 = -1
            _ = len(v2)
            _ = fi139
            _ = v146
        }
        var v147 int32 = -683
        if (true && (true || (uint(v4) != (v1 * (v1 & (^(v1))))))) {
            v20 += v20
            if (!(false && (!(100 == 42)))) {
                for fi148 := 0; fi148 < 4; fi148++ {
                    switch (v19 + (("hello" + v19) + "\n\t")) {
                    case "foo bar":
                        fmt.Println(fi138)
                    default:
                        switch "" {
                        case "hello":
                            v149 := GenS5{F0: -1, F1: v4}
                            _ = v149
                        case "":
                            if false {
                                v19 = v19
                                if ((^(fi148)) == fi148) {
                                    var v150 int64 = 9223372036854775807
                                    var v151 uint16 = 474
                                    _ = v150
                                    _ = v151
                                }
                                var v152 int16 = 519
                                _ = v152
                            }
                        case "foo bar":
                            v153 := GenS8{F0: 1}
                            _ = v153
                        }
                    }
                    switch fi148 {
                    case 0:
                        v3 += ((0.1 + ((((v20 - -2.25) - v3) + (v20 + v20)) - (float64((3.14159 - (0.1 + float32(1e10)))) + (v3 - (float64(1.5) / (((v20 / ((v20 + (1.5 - v20)) / v3)) * (v20 * v20)) + (v3 - v3))))))) + v20)
                        v4 = (v4 << ((2) & 63))
                    case -140:
                        v2[1] = (fi148 - (-(fi138)))
                    }
                    v154 := GenS15{F0: true, F1: 4294967295}
                    _ = fi148
                    _ = v154
                }
                var v155 uint16 = v4
                v156 := GenS18{F0: (^(v147)), F1: -980}
                _ = v155
                _ = v156
            }
        } else {
            fmt.Println(v147)
        }
        _ = fi138
        _ = v147
    }
    if false {
        fmt.Println(v3)
        if len(v5) > 0 {
            v5[0] = -1
        }
        for fi157 := 0; fi157 < 8; fi157++ {
            var v158 bool = true
            if v158 {
                delete(v2, 2147483647)
                switch fi157 {
                case -1:
                    var v159 int32 = 2147483647
                    _ = v159
                default:
                    for fi160 := 0; fi160 < 4; fi160++ {
                        var v161 uint32 = 2147483647
                        v4 += v4
                        v158 = true
                        _ = fi160
                        _ = v161
                    }
                }
            }
            _ = fi157
            _ = v158
        }
    } else {
        if (235 > 42) {
            var v162 int64 = -9223372036854775808
            _ = v162
        } else {
            v20 = (-2.25 + 0.0)
        }
        v20 += -2.25
    }
    switch 0 {
    case 730:
        v20 = v20
    }
    var v163 string = "foo bar"
    v5 = append(v5, -1)
    if false {
        v2[0] = 2147483647
    } else {
        if (2147483647 >= 209) {
            v4 = (v4 >> ((v1) & 63))
        } else {
            switch -535 {
            case 42:
                var v164 uint32 = 65535
                _ = v164
            case -1:
                for fi165 := 0; fi165 < 19; fi165++ {
                    var v166 uint32 = 0
                    v163 = v163
                    _ = fi165
                    _ = v166
                }
            case -897:
                v163 = (v163 + (v19 + "foo bar"))
            }
            if true {
                var v167 int64 = 42
                for fi168 := 0; fi168 < 2; fi168++ {
                    v19 = (v19 + (v19 + v163))
                    _ = fi168
                }
                _ = v167
            }
            v169 := GenS6{F0: (true && true)}
            _ = v169
        }
    }
    var v170 string = (("foo bar" + "foo bar") + v163)
    for fi171 := 0; fi171 < 12; fi171++ {
        delete(v2, 2147483647)
        switch fi171 {
        case -2147483648:
            if len(v5) > 0 {
                v5[0] = fi171
            }
        }
        var v172 bool = (!(int64(v4) <= int64(v1)))
        _ = fi171
        _ = v172
    }
    var v173 bool = false
    v174 := GenS1{F0: 1, F1: 100, F2: v19}
    fmt.Println("v1:", v1)
    fmt.Println("v163:", v163)
    fmt.Println("v170:", v170)
    fmt.Println("v173:", v173)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v163
    _ = v170
    _ = v173
    _ = v174
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
