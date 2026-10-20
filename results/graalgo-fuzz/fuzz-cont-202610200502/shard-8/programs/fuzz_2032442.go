package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint16; F1 uint16; F2 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int64; F1 uint; F2 int32 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 uint64; F1 int; F2 int16 }
type GenS15 struct { F0 int8; F1 int16; F2 int }

func main() {
    var v1 int64 = 42
    var v2 string = (((("foo bar" + "\n\t") + ("a" + "\n\t")) + "a") + ("foo bar" + ""))
    var v3 int8 = 81
    v4 := make(map[int]bool)
    var v5 int8 = -1
    var v6 float64 = (-2.25 / float64(((-2.25 + 1e10) / 3.14159)))
    v1 += (v1 ^ v1)
    v1 -= v1
    if true {
        fmt.Println(v3, v3)
    } else {
        for fi7 := 0; fi7 < 19; fi7++ {
            if (!true) {
                var v8 uint8 = 2
                var v9 uint8 = 127
                v10 := GenS1{F0: ("foo bar" + ("" + ((v2 + (v2 + (v2 + "hello"))) + v2))), F1: 255}
                _ = v10
                _ = v8
                _ = v9
            } else {
                var v11 bool = true
                var v12 uint = 345
                v13 := []int{(fi7 << (((v12 + 18446744073709551615)) & 63))}
                _ = v11
                _ = v12
                _ = v13
            }
            var v14 int = fi7
            _ = fi7
            _ = v14
        }
        if ((!((v1 & v1) <= (v1 * v1))) && true) {
            for fi15 := 0; fi15 < 18; fi15++ {
                if (18446744073709551615 == 18446744073709551615) {
                    if false {
                        v16 := GenS1{F0: (v2 + v2), F1: 42}
                        fmt.Println(v1)
                        v17 := []int{fi15, (fi15 ^ 32767)}
                        _ = v16
                        _ = v17
                    }
                    var v18 float64 = float64(1e10)
                    v19 := []int{(fi15 % (((-(fi15))) | 1)), 128, fi15, fi15}
                    _ = v18
                    _ = v19
                } else {
                    switch fi15 {
                    case 127:
                        switch fi15 {
                        case 4294967295:
                            v3 = 100
                        case 9223372036854775807:
                            if ((!false) && ((!true) || (32767 < 0))) {
                                for fi20 := 0; fi20 < 2; fi20++ {
                                    v21 := GenS2{F0: 0, F1: 42, F2: (v1 ^ v1)}
                                    if (v5 != v3) {
                                        delete(v4, fi20)
                                        fmt.Println(v2, fi15)
                                        v22 := GenS1{F0: v2, F1: 32767}
                                        _ = v22
                                    } else {
                                        var v23 uint64 = 2
                                        _ = v23
                                    }
                                    var v24 int32 = 651
                                    _ = fi20
                                    _ = v21
                                    _ = v24
                                }
                                if false {
                                    fmt.Println(v1, v2, v3)
                                    v25 := v4[(^(fi15))]
                                    _ = v25
                                }
                                var v26 int8 = v3
                                _ = v26
                            }
                        }
                    case 32767:
                        var v27 uint16 = 32767
                        _ = v27
                    }
                    delete(v4, fi15)
                    v28 := GenS3{F0: 65535}
                    _ = v28
                }
                v29 := []int{(fi15 % ((42) | 1)), fi15}
                fmt.Println(v2, v3)
                _ = fi15
                _ = v29
            }
            for fi30 := 0; fi30 < 17; fi30++ {
                if false {
                    var v31 int64 = v1
                    _ = v31
                } else {
                    if true {
                        switch (fi30 << ((4294967295) & 63)) {
                        case -455:
                            v1 = 9223372036854775807
                        case 42:
                            v32 := []int{201, (fi30 + (fi30 / ((fi30) | 1))), 127, fi30}
                            _ = v32
                        }
                        v33 := GenS1{F0: v2, F1: 32767}
                        for fi34 := 0; fi34 < 11; fi34++ {
                            v2 = (v2 + "")
                            _ = fi34
                        }
                        _ = v33
                    }
                }
                var v35 string = (v2 + (v2 + v2))
                v1 = v1
                _ = fi30
                _ = v35
            }
            var v36 int64 = v1
            _ = v36
        } else {
            var v37 uint16 = 255
            v3 += -6
            fmt.Println(v3, v2)
            _ = v37
        }
    }
    delete(v4, 128)
    fmt.Println(v3, v5, v3)
    if (0 >= 1) {
        v1 -= (v1 - v1)
        _ = len(v4)
        v3 -= 42
    } else {
        v4[248] = true
        var v38 int64 = (v1 << ((2) & 63))
        _ = v38
    }
    if (((v6 * float64((-2.25 / -2.25))) - v6) < v6) {
        if false {
            v39 := GenS4{F0: (v2 + v2)}
            _ = v39
        } else {
            for fi40 := 0; fi40 < 3; fi40++ {
                v41 := GenS2{F0: 65535, F1: 42, F2: -9223372036854775808}
                v6 = v6
                _ = fi40
                _ = v41
            }
            switch 4294967295 {
            case 42:
                var v42 float64 = (v6 + ((v6 - ((v6 / (float64(-2.25) / 3.14159)) - float64((3.14159 / 1.0)))) / float64(float32(v6))))
                _ = v42
            }
            if ((!true) && ((-(v3)) >= (v3 << ((758) & 63)))) {
                var v43 string = (v2 + v2)
                v44 := GenS1{F0: v43, F1: 255}
                _ = v43
                _ = v44
            } else {
                var v45 string = v2
                _ = v45
            }
        }
        v4[42] = true
    }
    var v46 bool = ((float64((((-2.25 - (1.5 - 3.14159)) - -2.25) / float32((v6 - 1e10)))) - (v6 / ((((float64(3.14159) - float64(3.14159)) + float64(-2.25)) * 1e10) / v6))) < 0.0)
    var v47 float64 = v6
    for fi48 := 0; fi48 < 19; fi48++ {
        fmt.Println(v3, v5, v47)
        fmt.Println(v1, fi48, v2)
        _ = fi48
    }
    if false {
        var v49 uint64 = 63
        v3 = (v5 / ((v5) | 1))
        v50 := GenS1{F0: ("a" + v2), F1: 430}
        _ = v49
        _ = v50
    } else {
        switch (v2 + v2) {
        case "":
            if ((v1 >= -1) || (0 != 36)) {
                v6 -= 0.1
            }
        case "a":
            switch v2 {
            case "\n\t":
                if v46 {
                    v51 := GenS2{F0: 1, F1: 754, F2: (v1 & (v1 - (v1 * -9223372036854775808)))}
                    var v52 string = ((v2 + "hello") + (v2 + "a"))
                    v53 := GenS5{F0: 1, F1: 377}
                    _ = v51
                    _ = v52
                    _ = v53
                }
            case "a":
                if v46 {
                    for fi54 := 0; fi54 < 9; fi54++ {
                        for fi55 := 0; fi55 < 3; fi55++ {
                            v56 := GenS1{F0: ("\n\t" + "a"), F1: 65535}
                            if (42 <= 32767) {
                                var v57 int8 = (v5 | v5)
                                v5 = (^(v3))
                                switch (fi54 | -948) {
                                case 32768:
                                    if (v46 && (!true)) {
                                        v58 := v4[fi55]
                                        switch fi54 {
                                        case -2147483648:
                                            if v46 {
                                                v6 += 1e10
                                                v59 := []int{1, fi55, fi55, fi55}
                                                _ = v59
                                            } else {
                                                v60 := GenS6{F0: 32767}
                                                var v61 uint64 = 526
                                                _ = v60
                                                _ = v61
                                            }
                                        case -1:
                                            v2 = v2
                                        case 1:
                                            if ((1e10 / 3.14159) < 1.5) {
                                                if v58 {
                                                    fmt.Println(v1, v46, v1)
                                                    v2 = v2
                                                } else {
                                                    if v46 {
                                                        if ((0.1 < 1e10) || v58) {
                                                            switch (^(fi54)) {
                                                            case 32768:
                                                                v3 = v57
                                                            default:
                                                                if (!(0.1 >= 0.0)) {
                                                                    v62 := GenS5{F0: 255, F1: (^(fi55))}
                                                                    var v63 int64 = (v1 >> ((18446744073709551615) & 63))
                                                                    _ = v62
                                                                    _ = v63
                                                                }
                                                            }
                                                            if true {
                                                                var v64 uint = 2
                                                                _ = v64
                                                            }
                                                        } else {
                                                            var v65 float64 = v6
                                                            if true {
                                                                v66 := GenS6{F0: 291}
                                                                if (v46 && (!v58)) {
                                                                    var v67 uint8 = 127
                                                                    var v68 int8 = v57
                                                                    _ = v67
                                                                    _ = v68
                                                                } else {
                                                                    v46 = v46
                                                                    var v69 int64 = v1
                                                                    v70 := GenS1{F0: v2, F1: 2}
                                                                    _ = v69
                                                                    _ = v70
                                                                }
                                                                fmt.Println(v46, v1)
                                                                _ = v66
                                                            } else {
                                                                v71 := GenS7{F0: 37, F1: 65535, F2: 2}
                                                                v3 = 42
                                                                _ = v71
                                                            }
                                                            _ = v65
                                                        }
                                                        var v72 float64 = 0.0
                                                        _ = v72
                                                    } else {
                                                        var v73 int32 = 407
                                                        v74 := GenS2{F0: 2, F1: 65535, F2: 0}
                                                        var v75 float64 = (v47 * -2.25)
                                                        _ = v73
                                                        _ = v74
                                                        _ = v75
                                                    }
                                                    v2 = v2
                                                }
                                            } else {
                                                v76 := GenS2{F0: 0, F1: 2, F2: v1}
                                                v47 = float64((-2.25 - 1.5))
                                                if false {
                                                    _ = len(v4)
                                                    if (v46 && v46) {
                                                        v77 := GenS5{F0: 272, F1: 4294967295}
                                                        _ = v77
                                                    }
                                                }
                                                v6 = v6
                                                _ = v76
                                            }
                                        default:
                                            var v78 int64 = (v1 << ((255) & 63))
                                            _ = v78
                                        }
                                        _ = v58
                                    }
                                default:
                                    v5 += (v57 * v57)
                                }
                                _ = v57
                            } else {
                                var v79 float64 = v47
                                switch 312 {
                                case -510:
                                    v47 += v47
                                case 127:
                                    var v80 uint64 = 0
                                    _ = v80
                                case -322:
                                    if (1 > 465) {
                                        switch (fi55 ^ (^(fi54))) {
                                        case -198:
                                            var v81 float64 = v47
                                            _ = v81
                                        }
                                        if v46 {
                                            if false {
                                                switch fi55 {
                                                case 32767:
                                                    v6 += v79
                                                case -9223372036854775808:
                                                    if false {
                                                        v82 := []int{310, fi55, 127, 345}
                                                        _ = v82
                                                    }
                                                case 9223372036854775807:
                                                    var v83 int16 = 183
                                                    _ = v83
                                                }
                                                var v84 int = (fi55 | fi54)
                                                v85 := GenS3{F0: 42}
                                                _ = v84
                                                _ = v85
                                            } else {
                                                if v46 {
                                                    if (!(2147483648 > 4294967295)) {
                                                        v86 := GenS8{F0: 693}
                                                        _ = v86
                                                    }
                                                    switch v2 {
                                                    case "\n\t":
                                                        fmt.Println(v5)
                                                    case "":
                                                        var v87 bool = v46
                                                        _ = v87
                                                    case "foo bar":
                                                        v88 := GenS1{F0: v2, F1: 42}
                                                        _ = v88
                                                    }
                                                } else {
                                                    var v89 int8 = (-(v5))
                                                    if v46 {
                                                        if (1 >= fi54) {
                                                            v5 = (v89 | (v89 << ((4294967295) & 63)))
                                                            switch ("" + "") {
                                                            case "a":
                                                                if (-75 == v3) {
                                                                    if (v46 && true) {
                                                                        v90 := []int{fi55}
                                                                        v91 := GenS9{F0: (v89 - (v89 - 0))}
                                                                        switch (fi55 - (^(fi55))) {
                                                                        case 127:
                                                                            v3 = (v3 ^ v5)
                                                                        case -2147483648:
                                                                            var v92 uint8 = 255
                                                                            _ = v92
                                                                        }
                                                                        _ = v90
                                                                        _ = v91
                                                                    } else {
                                                                        fmt.Println(v47)
                                                                        fmt.Println(v89, v5)
                                                                        v6 += v79
                                                                        delete(v4, -1)
                                                                    }
                                                                    fmt.Println(v46, v46)
                                                                    if v46 {
                                                                        v46 = v46
                                                                        switch v2 {
                                                                        case "\n\t":
                                                                            v93 := GenS9{F0: (v3 % (((v3 / ((v3) | 1))) | 1))}
                                                                            fmt.Println(v6, fi55, v47)
                                                                            _ = v93
                                                                        case "foo bar":
                                                                            switch (fi54 - -2147483648) {
                                                                            case -117:
                                                                                switch (fi54 >> ((28) & 63)) {
                                                                                case 503:
                                                                                    var v94 int = (fi54 | -202)
                                                                                    _ = v94
                                                                                case -2147483648:
                                                                                    switch (fi55 - (fi55 & (fi55 << ((483) & 63)))) {
                                                                                    case -1:
                                                                                        v3 = v89
                                                                                    case 127:
                                                                                        var v95 uint16 = 32767
                                                                                        _ = v95
                                                                                    case 845:
                                                                                        if ((0.0 >= 0.1) || true) {
                                                                                            switch (fi54 * fi54) {
                                                                                            case 4294967295:
                                                                                                switch (fi54 ^ (fi54 >> ((1) & 63))) {
                                                                                                case 0:
                                                                                                    var v96 int8 = (v3 + int8(fi55))
                                                                                                    _ = v96
                                                                                                case -219:
                                                                                                    v79 += v6
                                                                                                case 385:
                                                                                                    if (!(!v46)) {
                                                                                                        var v97 int64 = v1
                                                                                                        v98 := GenS4{F0: "\n\t"}
                                                                                                        _ = v97
                                                                                                        _ = v98
                                                                                                    } else {
                                                                                                        var v99 int8 = v3
                                                                                                        var v100 bool = (42 != -2147483648)
                                                                                                        _ = v100
                                                                                                        _ = v99
                                                                                                    }
                                                                                                default:
                                                                                                    v101 := GenS4{F0: "\n\t"}
                                                                                                    _ = v101
                                                                                                }
                                                                                            case 32768:
                                                                                                if false {
                                                                                                    v102 := GenS10{F0: v1, F1: 1, F2: -1}
                                                                                                    v3 += v5
                                                                                                    if v46 {
                                                                                                        v103 := []int{fi55, (fi55 | fi54)}
                                                                                                        var v104 uint = 42
                                                                                                        switch (fi54 + (fi54 & (fi54 % (((fi54 << ((2) & 63))) | 1)))) {
                                                                                                        case 656:
                                                                                                            v105 := GenS10{F0: 0, F1: v104, F2: -1}
                                                                                                            _ = v105
                                                                                                        default:
                                                                                                            var v106 string = v2
                                                                                                            _ = v106
                                                                                                        }
                                                                                                        _ = v103
                                                                                                        _ = v104
                                                                                                    }
                                                                                                    fmt.Println(v46, v89, v89)
                                                                                                    v107 := []int{fi55, (fi55 / ((fi54) | 1))}
                                                                                                    _ = v102
                                                                                                    _ = v107
                                                                                                }
                                                                                            default:
                                                                                                var v108 int16 = 588
                                                                                                _ = v108
                                                                                            }
                                                                                            if (int8(v1) >= 1) {
                                                                                                var v109 uint32 = 475
                                                                                                fmt.Println(v109, v3)
                                                                                                _ = v109
                                                                                            } else {
                                                                                                v1 = (^(v1))
                                                                                                fmt.Println(v3, v46, v2)
                                                                                                v3 = (-(v3))
                                                                                            }
                                                                                            if false {
                                                                                                v1 += v1
                                                                                                v110 := []int{fi55, 4294967295}
                                                                                                v111 := GenS8{F0: 18446744073709551615}
                                                                                                _ = v110
                                                                                                _ = v111
                                                                                            }
                                                                                        } else {
                                                                                            var v112 uint64 = 18446744073709551615
                                                                                            if v46 {
                                                                                                v3 = v5
                                                                                                v4[-2147483648] = (1.5 == ((float64(((float32(3.14159) - float32(1e10)) / 1.0)) * 1e10) + (-2.25 + v6)))
                                                                                                v2 = v2
                                                                                            }
                                                                                            v113 := []int{(fi54 ^ 32768), fi54}
                                                                                            _ = v112
                                                                                            _ = v113
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    v114 := GenS9{F0: 100}
                                                                                    _ = v114
                                                                                }
                                                                            case 127:
                                                                                v6 = 1.5
                                                                            case 602:
                                                                                v115 := GenS5{F0: 32767, F1: (fi55 / ((-9223372036854775808) | 1))}
                                                                                _ = v115
                                                                            }
                                                                        case "hello":
                                                                            var v116 int = fi54
                                                                            _ = v116
                                                                        }
                                                                    }
                                                                } else {
                                                                    v117 := GenS11{F0: int64(v3)}
                                                                    var v118 int32 = 2147483647
                                                                    v2 = v2
                                                                    _ = v117
                                                                    _ = v118
                                                                }
                                                            case "foo bar":
                                                                var v119 string = ("foo bar" + v2)
                                                                _ = v119
                                                            }
                                                        }
                                                        v89 = -1
                                                    }
                                                    v120 := GenS12{F0: ("\n\t" > (("a" + v2) + v2))}
                                                    _ = v120
                                                    _ = v89
                                                }
                                                fmt.Println(fi54, v3, fi55)
                                                if false {
                                                    v121 := GenS13{F0: 8}
                                                    v122 := GenS1{F0: "\n\t", F1: 810}
                                                    _ = v121
                                                    _ = v122
                                                }
                                                v79 = 1.5
                                                fmt.Println(fi54, v3, v79)
                                                var v123 int32 = -2147483648
                                                var v124 uint16 = 920
                                                _ = v123
                                                _ = v124
                                            }
                                        } else {
                                            v125 := GenS7{F0: 2, F1: 145, F2: 0}
                                            _ = v125
                                        }
                                    } else {
                                        var v126 uint64 = 4294967295
                                        _ = v126
                                    }
                                default:
                                    if v46 {
                                        v4[42] = ((fi54 ^ (fi55 >> ((823) & 63))) < (fi54 - fi55))
                                        v3 = v5
                                        v5 -= (^(v3))
                                    }
                                }
                                _ = v79
                            }
                            v127 := GenS8{F0: 4294967295}
                            _ = fi55
                            _ = v127
                            _ = v56
                        }
                        if v46 {
                            v47 -= (float64(((1.5 / 1.0) / (1.5 * ((-2.25 - (float32(v6) + 0.1)) * 1.5)))) - v47)
                            fmt.Println(v5, v47)
                            if true {
                                for fi128 := 0; fi128 < 9; fi128++ {
                                    v129 := GenS9{F0: (v3 * v3)}
                                    v130 := GenS3{F0: 389}
                                    if (2 == 0) {
                                        v5 = v5
                                    } else {
                                        switch fi54 {
                                        case 9223372036854775807:
                                            delete(v4, (fi54 ^ 785))
                                        }
                                    }
                                    _ = fi128
                                    _ = v129
                                    _ = v130
                                }
                                if v46 {
                                    v131 := []int{fi54, fi54, fi54, fi54}
                                    _ = v131
                                }
                                v3 -= v3
                            } else {
                                v4[42] = (4294967295 < 2)
                                var v132 uint = 65535
                                v133 := []int{fi54}
                                _ = v132
                                _ = v133
                            }
                        }
                        v46 = (18446744073709551615 >= 4294967295)
                        _ = fi54
                    }
                    if (v46 && (2 < 1)) {
                        v2 = v2
                        for fi134 := 0; fi134 < 15; fi134++ {
                            for fi135 := 0; fi135 < 11; fi135++ {
                                if v46 {
                                    v136 := GenS14{F0: 2, F1: (fi134 | (fi134 / ((fi134) | 1))), F2: -32768}
                                    if (158 > 108) {
                                        v137 := GenS13{F0: 589}
                                        fmt.Println(v6, fi135)
                                        v3 = v3
                                        _ = v137
                                    } else {
                                        v1 -= (v1 - v1)
                                        v138 := GenS11{F0: 42}
                                        _ = v138
                                    }
                                    _ = v136
                                } else {
                                    v139 := GenS15{F0: v3, F1: 589, F2: fi134}
                                    v46 = (((0.1 / -2.25) == float64((3.14159 * (3.14159 / 0.1)))) || (0 >= -981))
                                    _ = v139
                                }
                                var v140 uint32 = 2
                                if ((!v46) && v46) {
                                    var v141 int8 = v5
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v142 float64 = v47
                                        _ = v142
                                    case "":
                                        fmt.Println(fi135)
                                    }
                                    v5 -= v141
                                    _ = v141
                                } else {
                                    v46 = (2 < 617)
                                    v3 -= (v3 + (v3 + v5))
                                    v2 = "foo bar"
                                }
                                _ = fi135
                                _ = v140
                            }
                            v143 := GenS4{F0: "a"}
                            if v46 {
                                for fi144 := 0; fi144 < 5; fi144++ {
                                    _ = len(v4)
                                    fmt.Println(v2, v2)
                                    _ = fi144
                                }
                                var v145 int = fi134
                                var v146 bool = (!(v46 && v46))
                                _ = v145
                                _ = v146
                            }
                            _ = fi134
                            _ = v143
                        }
                        v147 := []int{}
                        _ = v147
                    }
                    v148 := []int{330, -9223372036854775808}
                    _ = v148
                }
            }
        case "foo bar":
            v47 -= (float64((-2.25 - float32((float64(((1e10 + ((float32(1.5) + (1e10 / float32(v47))) - 1.5)) + (1.5 - ((0.1 + 0.1) * 1.5)))) * float64((float32((-2.25 / 1.0)) * (((1e10 - -2.25) * 1e10) * 3.14159))))))) * v6)
        default:
            v3 += (v3 * v5)
        }
    }
    for fi149 := 0; fi149 < 15; fi149++ {
        for fi150 := 0; fi150 < 6; fi150++ {
            var v151 float32 = 1.5
            var v152 int8 = -4
            _ = fi150
            _ = v151
            _ = v152
        }
        _ = fi149
    }
    if false {
        v153 := []int{-2147483648, 127}
        _ = v153
    } else {
        var v154 uint = 42
        v155 := []int{-405, -840, -885}
        for fi156 := 0; fi156 < 10; fi156++ {
            delete(v4, 0)
            fmt.Println(v46)
            _ = fi156
        }
        _ = v154
        _ = v155
    }
    v1 -= (v1 | (v1 ^ (v1 * -1)))
    v6 += ((v6 * (v6 - float64(1.5))) - v6)
    fmt.Println(v47)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
