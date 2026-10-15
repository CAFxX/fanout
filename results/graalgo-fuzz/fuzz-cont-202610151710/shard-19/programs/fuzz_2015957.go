package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 uint; F1 uint16 }
type GenS6 struct { F0 string; F1 string; F2 int }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 bool; F1 int64 }
type GenS11 struct { F0 uint16; F1 uint; F2 bool }
type GenS12 struct { F0 int16; F1 uint16; F2 int64 }
type GenS13 struct { F0 int64; F1 string }
type GenS14 struct { F0 int64; F1 bool; F2 uint16 }
type GenS15 struct { F0 int; F1 uint64 }
type GenS16 struct { F0 string }
type GenS17 struct { F0 int8 }
func genFn0(p0 string, p1 uint8, p2 int16) int64 {
    if true {
        v7 := GenS1{F0: -9223372036854775808, F1: 42}
        p0 = "a"
        _ = v7
    } else {
        var v8 string = p0
        var v9 int16 = (^(p2))
        _ = v8
        _ = v9
    }
    var v10 float32 = -2.25
    p0 = ((p0 + "foo bar") + p0)
    switch p0 {
    case "\n\t":
        fmt.Println(p1)
    case "a":
        p1 += (p1 / ((p1) | 1))
    case "hello":
        p2 = p2
    default:
        if (!false) {
            for fi11 := 0; fi11 < 19; fi11++ {
                v12 := []int{}
                for fi13 := 0; fi13 < 18; fi13++ {
                    var v14 uint = 2
                    switch fi13 {
                    case 32768:
                        var v15 uint32 = 0
                        _ = v15
                    default:
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                    }
                    if (766 < 42) {
                        if (1.5 == 3.14159) {
                            if len(v12) > 0 {
                                v12[0] = (fi11 & -2147483648)
                            }
                            p2 = -297
                            v16 := GenS2{F0: 42, F1: 392}
                            _ = v16
                        }
                        if (p0 != p0) {
                            if (int64(p2) != int64(v14)) {
                                switch ("a" + p0) {
                                case "a":
                                    _ = len(v12)
                                case "hello":
                                    fmt.Println(fi13, fi11, p2)
                                    v17 := GenS2{F0: -1, F1: 318}
                                    var v18 uint8 = p1
                                    var v19 float64 = 0.1
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                default:
                                    var v20 uint8 = (p1 | 128)
                                    _ = v20
                                }
                                v21 := GenS1{F0: -1, F1: 756}
                                p0 = p0
                                _ = v21
                            }
                        } else {
                            switch p0 {
                            case "hello":
                                p1 += 197
                            case "foo bar":
                                v22 := GenS3{F0: 791}
                                _ = v22
                            default:
                                var v23 uint16 = 179
                                _ = v23
                            }
                        }
                        var v24 uint32 = 791
                        var v25 int8 = 100
                        if false {
                            var v26 int64 = 42
                            if true {
                                var v27 uint = v14
                                v28 := GenS4{F0: (^(p1)), F1: fi11}
                                _ = v27
                                _ = v28
                            }
                            var v29 float32 = ((v10 + 3.14159) / 1.5)
                            _ = v26
                            _ = v29
                        } else {
                            var v30 int64 = -9223372036854775808
                            _ = v30
                        }
                        switch (-(fi11)) {
                        case 4294967295:
                            v24 = (v24 >> ((v14) & 63))
                        case 32767:
                            switch fi11 {
                            case -9223372036854775808:
                                var v31 uint64 = 347
                                var v32 float32 = v10
                                _ = len(v12)
                                _ = v31
                                _ = v32
                            case 42:
                                var v33 uint64 = 2
                                _ = v33
                            case 128:
                                v34 := GenS5{F0: (v14 / ((0) | 1)), F1: 193}
                                _ = v34
                            default:
                                if (-121 != int8(p1)) {
                                    var v35 uint16 = 143
                                    v36 := GenS4{F0: (p1 | (p1 >> (((^(v14))) & 63))), F1: 1}
                                    _ = len(v12)
                                    _ = v35
                                    _ = v36
                                } else {
                                    if false {
                                        v24 += v24
                                    }
                                    if true {
                                        if (2 < 264) {
                                            if (v14 <= (^(v14))) {
                                                var v37 int = 128
                                                _ = v37
                                            }
                                            fmt.Println(v14, fi13)
                                        } else {
                                            fmt.Println(v14, p0)
                                        }
                                        if true {
                                            p1 = (p1 << (((v14 - v14)) & 63))
                                        }
                                    }
                                }
                            }
                        case -2147483648:
                            p1 = p1
                        default:
                            if (575 >= (^(v14))) {
                                v38 := GenS3{F0: 4294967295}
                                fmt.Println(v10, v25)
                                switch fi13 {
                                case 704:
                                    v39 := GenS2{F0: 42, F1: uint32(p1)}
                                    _ = v39
                                }
                                v14 = uint(p1)
                                p2 += int16(fi11)
                                v24 = 295
                                switch (p0 + p0) {
                                case "foo bar":
                                    fmt.Println(v14, p0, v10)
                                case "":
                                    v40 := GenS5{F0: v14, F1: 0}
                                    _ = v40
                                case "\n\t":
                                    var v41 int64 = 9223372036854775807
                                    _ = v41
                                }
                                _ = v38
                            } else {
                                p2 -= int16(p1)
                                switch (fi11 & 2147483647) {
                                case 1:
                                    v24 = v24
                                case -316:
                                    v24 = 843
                                case 4294967295:
                                    v14 = (^(v14))
                                default:
                                    v42 := GenS4{F0: 194, F1: (^(fi11))}
                                    _ = v42
                                }
                            }
                        }
                        switch (fi11 / ((-2147483648) | 1)) {
                        case 32767:
                            if ((fi11 / ((fi13) | 1)) != 289) {
                                v10 += 1.5
                                if (true || false) {
                                    v25 += v25
                                }
                                if false {
                                    if len(v12) >= 2 {
                                        v12 = v12[1:len(v12)]
                                    }
                                    var v43 float64 = (3.14159 - 1e10)
                                    var v44 uint64 = 18446744073709551615
                                    fmt.Println(v44)
                                    _ = v43
                                    _ = v44
                                } else {
                                    v24 -= (v24 * v24)
                                    v12 = append(v12, fi11)
                                    if false {
                                        v10 -= v10
                                        if false {
                                            if false {
                                                v45 := GenS1{F0: 9223372036854775807, F1: -1}
                                                _ = v45
                                            } else {
                                                if len(v12) > 0 {
                                                    v12[0] = fi13
                                                }
                                                var v46 string = (p0 + p0)
                                                var v47 int32 = 42
                                                _ = v46
                                                _ = v47
                                            }
                                            v10 = ((v10 + v10) * 0.1)
                                        } else {
                                            var v48 int64 = -809
                                            _ = v48
                                        }
                                        if (!false) {
                                            var v49 int64 = -1
                                            _ = v49
                                        } else {
                                            var v50 int64 = -111
                                            _ = v50
                                        }
                                    } else {
                                        switch 9223372036854775807 {
                                        case 127:
                                            fmt.Println(fi13, p1, v14)
                                        case 534:
                                            var v51 string = (p0 + p0)
                                            _ = v51
                                        case 174:
                                            switch fi11 {
                                            case -925:
                                                if false {
                                                    if true {
                                                        v52 := GenS4{F0: (p1 * 128), F1: int(v14)}
                                                        if false {
                                                            if len(v12) >= 2 {
                                                                v12 = v12[1:len(v12)]
                                                            }
                                                            if true {
                                                                v24 -= v24
                                                                p1 = (p1 | p1)
                                                                if (((v24 - v24) >= v24) || true) {
                                                                    if (((p0 + ("a" + "\n\t")) + p0) < ("" + ("" + p0))) {
                                                                        v53 := GenS6{F0: "", F1: p0, F2: (fi11 + 0)}
                                                                        switch p0 {
                                                                        case "a":
                                                                            v54 := GenS7{F0: v25}
                                                                            _ = v54
                                                                        case "hello":
                                                                            if false {
                                                                                var v55 uint8 = p1
                                                                                var v56 int16 = p2
                                                                                _ = v55
                                                                                _ = v56
                                                                            } else {
                                                                                var v57 int = fi11
                                                                                v24 = v24
                                                                                _ = v57
                                                                            }
                                                                        case "":
                                                                            fmt.Println(v24)
                                                                        default:
                                                                            var v58 int = (fi13 * (fi13 * (fi11 + (fi13 - -9223372036854775808))))
                                                                            v58 += v58
                                                                            var v59 int = 32767
                                                                            var v60 int16 = -1
                                                                            switch (v58 << (((v14 - 471)) & 63)) {
                                                                            case 32768:
                                                                                var v61 int32 = -2147483648
                                                                                _ = v61
                                                                            case 1:
                                                                                if len(v12) > 0 {
                                                                                    v12[0] = 397
                                                                                }
                                                                            default:
                                                                                switch fi13 {
                                                                                case 9223372036854775807:
                                                                                    v12 = append(v12, (v58 << (((v14 - (^(v14)))) & 63)))
                                                                                case 42:
                                                                                    var v62 uint16 = 979
                                                                                    v25 = 1
                                                                                    fmt.Println(v62)
                                                                                    _ = v62
                                                                                }
                                                                            }
                                                                            _ = v58
                                                                            _ = v59
                                                                            _ = v60
                                                                        }
                                                                        var v63 int16 = (p2 << (((v14 / (((v14 / (((v14 & (v14 | 991))) | 1))) | 1))) & 63))
                                                                        _ = len(v12)
                                                                        if len(v12) > 0 {
                                                                            v12[0] = (^(fi13))
                                                                        }
                                                                        var v64 float32 = float32(0.1)
                                                                        _ = v53
                                                                        _ = v63
                                                                        _ = v64
                                                                    }
                                                                    var v65 string = p0
                                                                    _ = v65
                                                                }
                                                            } else {
                                                                v66 := GenS8{F0: 0, F1: (v24 + v24), F2: 874}
                                                                v10 = 1.5
                                                                v67 := GenS8{F0: 2, F1: v24, F2: 2}
                                                                _ = v66
                                                                _ = v67
                                                            }
                                                            var v68 int16 = 46
                                                            v69 := GenS9{F0: uint(v24)}
                                                            p1 += 42
                                                            v10 = (float32(3.14159) / float32(((float64((v10 - ((v10 + 0.1) * ((v10 - v10) * (((1e10 - (-2.25 - float32(0.1))) * v10) / v10))))) + 1.5) * (0.0 / 3.14159))))
                                                            _ = len(v12)
                                                            if true {
                                                                if false {
                                                                    v70 := GenS3{F0: 3}
                                                                    _ = v70
                                                                }
                                                                if (false || ((fi13 - 128) > (fi13 * (fi13 & fi11)))) {
                                                                    if len(v12) > 0 {
                                                                        v12[0] = fi11
                                                                    }
                                                                    v71 := GenS10{F0: true, F1: -540}
                                                                    v24 -= (v24 % (((v24 % (((v24 << ((v14) & 63))) | 1))) | 1))
                                                                    v14 += (v14 * 42)
                                                                    _ = v71
                                                                }
                                                            }
                                                            v72 := GenS7{F0: 42}
                                                            _ = v68
                                                            _ = v69
                                                            _ = v72
                                                        }
                                                        fmt.Println(v14)
                                                        var v73 string = "foo bar"
                                                        v74 := GenS1{F0: 228, F1: -448}
                                                        _ = v52
                                                        _ = v73
                                                        _ = v74
                                                    } else {
                                                        switch fi13 {
                                                        case 32767:
                                                            switch "foo bar" {
                                                            case "hello":
                                                                var v75 int8 = -128
                                                                _ = v75
                                                            case "":
                                                                if len(v12) > 0 {
                                                                    v12[0] = fi11
                                                                }
                                                            default:
                                                                var v76 uint32 = 2147483647
                                                                _ = v76
                                                            }
                                                        case 127:
                                                            if (false && (!(true || (!(((v14 / ((42) | 1)) <= v14) || (((((p0 + ("\n\t" + p0)) + p0) + p0) + ("\n\t" + "\n\t")) >= p0)))))) {
                                                                p0 = p0
                                                            }
                                                        }
                                                        var v77 uint64 = 42
                                                        _ = v77
                                                    }
                                                    if ((v25 >> (((^(v14))) & 63)) > 42) {
                                                        v78 := GenS7{F0: (v25 ^ (^(v25)))}
                                                        _ = v78
                                                    } else {
                                                        p1 = p1
                                                        if false {
                                                            p2 = p2
                                                        }
                                                        fmt.Println(v14, fi11, v14)
                                                    }
                                                }
                                            case 4294967295:
                                                if false {
                                                    v79 := GenS11{F0: 171, F1: (^(v14)), F2: (!(false && true))}
                                                    var v80 uint = (^(v14))
                                                    _ = v79
                                                    _ = v80
                                                } else {
                                                    var v81 uint16 = 2
                                                    fmt.Println(p1, p0, p1)
                                                    v24 += v24
                                                    _ = v81
                                                }
                                            case 127:
                                                v82 := GenS3{F0: 18446744073709551615}
                                                _ = v82
                                            default:
                                                var v83 int16 = p2
                                                _ = v83
                                            }
                                        }
                                    }
                                    var v84 float64 = ((0.0 - -2.25) - (float64((-2.25 + float32(-2.25))) * 0.0))
                                    _ = v84
                                }
                            } else {
                                var v85 int64 = 1
                                var v86 int = (^(fi11))
                                v12 = append(v12, (^(fi11)))
                                _ = v85
                                _ = v86
                            }
                        }
                        _ = v24
                        _ = v25
                    }
                    _ = fi13
                    _ = v14
                }
                for fi87 := 0; fi87 < 17; fi87++ {
                    var v88 float32 = v10
                    var v89 uint = 4294967295
                    v90 := GenS8{F0: 32767, F1: 2147483648, F2: 0}
                    v91 := GenS2{F0: 0, F1: 4294967295}
                    _ = fi87
                    _ = v88
                    _ = v89
                    _ = v90
                    _ = v91
                }
                _ = fi11
                _ = v12
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return -341
}
func genFn1(p0 uint16) uint32 {
    v92 := GenS3{F0: 775}
    if true {
        for fi93 := 0; fi93 < 18; fi93++ {
            var v94 int8 = 112
            if true {
                var v95 int32 = 2147483647
                for fi96 := 0; fi96 < 2; fi96++ {
                    if false {
                        v97 := GenS12{F0: -578, F1: (p0 % ((p0) | 1)), F2: -9223372036854775808}
                        switch fi93 {
                        case 2:
                            v95 -= v95
                        default:
                            if ((1.5 + 1e10) == ((3.14159 * 3.14159) - 0.0)) {
                                v95 += -2147483648
                                v98 := GenS9{F0: 255}
                                _ = v98
                            }
                        }
                        if false {
                            if (255 < 4294967295) {
                                v94 = v94
                            }
                        }
                        _ = v97
                    }
                    v99 := []int{}
                    v99 = append(v99, (fi96 ^ (fi96 & fi93)))
                    v94 = v94
                    fmt.Println(v94, fi93)
                    _ = fi96
                    _ = v99
                }
                _ = v95
            }
            if false {
                v100 := []int{}
                _ = v100
            } else {
                switch "a" {
                case "a":
                    v101 := GenS4{F0: 23, F1: fi93}
                    _ = v101
                default:
                    for fi102 := 0; fi102 < 20; fi102++ {
                        p0 -= 2
                        if false {
                            fmt.Println(fi102, p0, v94)
                        } else {
                            var v103 uint8 = 255
                            if ((^(p0)) < 739) {
                                var v104 float64 = (0.1 + (-2.25 * 1e10))
                                var v105 float64 = float64(3.14159)
                                _ = v104
                                _ = v105
                            } else {
                                v106 := []int{4294967295, (fi102 & fi93), (fi102 ^ -9223372036854775808), fi102}
                                var v107 bool = true
                                _ = v106
                                _ = v107
                            }
                            v103 -= v103
                            _ = v103
                        }
                        _ = fi102
                    }
                }
                if false {
                    v108 := []int{}
                    _ = v108
                } else {
                    v94 -= v94
                    v109 := GenS8{F0: (p0 << ((65535) & 63)), F1: 355, F2: p0}
                    _ = v109
                }
            }
            _ = fi93
            _ = v94
        }
        var v110 uint8 = 93
        _ = v110
    }
    _ = p0
    _ = v92
    return 0
}

func main() {
    v1 := []int{-9223372036854775808, 42, 490}
    var v2 uint32 = 65535
    var v3 int8 = 42
    var v4 uint64 = 1
    var v5 float32 = (float32((3.14159 / 0.1)) * 0.1)
    var v6 uint16 = 2
    if ("a" < "\n\t") {
        var v111 float64 = 1e10
        _ = v111
    } else {
        v5 = -2.25
        var v112 uint = 255
        if true {
            for fi113 := 0; fi113 < 8; fi113++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                for fi114 := 0; fi114 < 3; fi114++ {
                    if (false || (127 > 0)) {
                        var v115 string = (("a" + "\n\t") + ("\n\t" + "a"))
                        switch v115 {
                        case "hello":
                            v116 := GenS8{F0: 255, F1: v2, F2: v6}
                            _ = v116
                        case "foo bar":
                            v117 := GenS12{F0: -1, F1: (v6 | 32767), F2: -1}
                            v2 += 2147483648
                            var v118 uint32 = (^(v2))
                            _ = v117
                            _ = v118
                        }
                        _ = v115
                    }
                    _ = fi114
                }
                _ = fi113
            }
            var v119 bool = (float32(0.1) >= float32(0.1))
            _ = v119
        } else {
            for fi120 := 0; fi120 < 10; fi120++ {
                fmt.Println(v4)
                switch fi120 {
                case 32768:
                    v3 = v3
                case 0:
                    for fi121 := 0; fi121 < 11; fi121++ {
                        if (((^(v4)) != 1) || (!true)) {
                            if true {
                                v4 = (^(v4))
                                var v122 int8 = -120
                                v123 := GenS8{F0: v6, F1: 1, F2: v6}
                                _ = v122
                                _ = v123
                            }
                            var v124 float64 = 0.1
                            _ = v124
                        } else {
                            v3 -= -128
                            var v125 int16 = -1
                            if (fi121 == fi120) {
                                if false {
                                    fmt.Println(v3)
                                    v126 := GenS6{F0: "", F1: "\n\t", F2: (fi121 >> ((v112) & 63))}
                                    _ = v126
                                }
                                var v127 int16 = 744
                                v128 := GenS6{F0: "", F1: (("foo bar" + "hello") + "\n\t"), F2: (fi121 / ((32768) | 1))}
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v1 = append(v1, (fi121 | (fi120 * (fi121 & fi121))))
                                _ = v127
                                _ = v128
                            } else {
                                var v129 int = fi121
                                switch 262 {
                                case 769:
                                    v129 += v129
                                }
                                v4 -= (v4 + v4)
                                _ = v129
                            }
                            _ = v125
                        }
                        switch (^(fi120)) {
                        case 32768:
                            v130 := GenS13{F0: 1, F1: (("" + ("foo bar" + "\n\t")) + "")}
                            _ = v130
                        case 505:
                            var v131 int64 = 535
                            _ = v131
                        case 580:
                            if ((!(((1 < -522) && (0 == 1000)) || true)) || false) {
                                fmt.Println(v2)
                            } else {
                                v1 = append(v1, (fi120 / ((fi120) | 1)))
                                v5 += ((float32(0.1) / v5) / v5)
                                v3 = (v3 | v3)
                            }
                        }
                        var v132 float64 = 0.0
                        _ = fi121
                        _ = v132
                    }
                }
                _ = fi120
            }
            v3 = 39
        }
        _ = v112
    }
    if (!((v2 | genFn1(uint16(v4))) < 306)) {
        switch 1 {
        case 4294967295:
            var v133 float64 = -2.25
            _ = v133
        default:
            for fi134 := 0; fi134 < 17; fi134++ {
                switch fi134 {
                case 1:
                    var v135 bool = ((v5 * float32((3.14159 - (float64(-2.25) - (1e10 * 0.1))))) < (float32((-2.25 - -2.25)) / 0.1))
                    _ = v135
                case -739:
                    v3 = v3
                default:
                    var v136 float64 = -2.25
                    _ = v136
                }
                var v137 string = "a"
                _ = fi134
                _ = v137
            }
        }
        v1 = append(v1, 42)
    }
    var v138 uint32 = 2
    var v139 int = 128
    var v140 string = "\n\t"
    var v141 uint32 = (v138 + (v2 ^ (^(v138))))
    if false {
        fmt.Println(v139)
        v142 := GenS10{F0: true, F1: 0}
        var v143 uint16 = v6
        _ = v142
        _ = v143
    }
    fmt.Println(v4, v138)
    if false {
        v1 = append(v1, v139)
    }
    for fi144 := 0; fi144 < 19; fi144++ {
        if len(v1) > 0 {
            v1[0] = (fi144 & 1)
        }
        _ = fi144
    }
    var v145 int32 = 2147483647
    switch v140 {
    case "\n\t":
        v146 := GenS14{F0: -663, F1: true, F2: 0}
        _ = v146
    case "":
        v2 = v141
    case "hello":
        for fi147 := 0; fi147 < 18; fi147++ {
            if (("foo bar" != ("hello" + v140)) || ((v3 * -1) <= 42)) {
                if true {
                    if false {
                        var v148 uint64 = (v4 % (((^(v4))) | 1))
                        switch (^(v139)) {
                        case 128:
                            switch v139 {
                            case 1:
                                for fi149 := 0; fi149 < 15; fi149++ {
                                    if false {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        fmt.Println(v138)
                                        fmt.Println(v6, v6, v5)
                                    } else {
                                        switch v139 {
                                        case 32767:
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        case -305:
                                            switch (v140 + "a") {
                                            case "a":
                                                fmt.Println(v141, v4, v4)
                                            default:
                                                v145 = 607
                                            }
                                        }
                                    }
                                    _ = fi149
                                }
                            default:
                                if len(v1) > 0 {
                                    v1[0] = fi147
                                }
                            }
                        case 902:
                            var v150 float32 = 3.14159
                            _ = v150
                        case 42:
                            _ = len(v1)
                        }
                        _ = v148
                    }
                    v151 := GenS7{F0: (v3 << ((uint(v6)) & 63))}
                    switch int(v3) {
                    case 32768:
                        if (!(true && (!false))) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    case 235:
                        v152 := GenS13{F0: 824, F1: ""}
                        _ = v152
                    case -2147483648:
                        if true {
                            var v153 float32 = 1.5
                            for fi154 := 0; fi154 < 15; fi154++ {
                                v155 := GenS2{F0: -1, F1: uint32(v4)}
                                _ = fi154
                                _ = v155
                            }
                            v3 -= 127
                            _ = v153
                        } else {
                            var v156 uint8 = 240
                            switch 0 {
                            case 6:
                                v6 = (v6 + (v6 << ((1) & 63)))
                            case 4294967295:
                                for fi157 := 0; fi157 < 16; fi157++ {
                                    var v158 int8 = v3
                                    if len(v1) > 0 {
                                        v1[0] = (-(fi157))
                                    }
                                    _ = fi157
                                    _ = v158
                                }
                            case 311:
                                fmt.Println(v156)
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v156
                        }
                    }
                    _ = v151
                }
                v159 := GenS15{F0: v139, F1: (^(v4))}
                _ = v159
            } else {
                var v160 uint16 = (v6 - uint16(v138))
                if (v4 >= v4) {
                    v138 -= v138
                    v1 = append(v1, v139)
                    if (0 >= -756) {
                        v140 = ("" + (v140 + ((((v140 + "hello") + v140) + (((v140 + v140) + v140) + v140)) + (v140 + (v140 + ((v140 + ("a" + "a")) + (v140 + "hello")))))))
                    } else {
                        var v161 string = ("foo bar" + ((("a" + ((v140 + v140) + ((v140 + (v140 + (v140 + ((("foo bar" + "") + "\n\t") + v140)))) + (v140 + v140)))) + "\n\t") + v140))
                        _ = v161
                    }
                } else {
                    var v162 float64 = (1.5 + 0.1)
                    v138 += (v141 ^ v2)
                    switch (("\n\t" + (v140 + ("hello" + (v140 + "a")))) + v140) {
                    case "foo bar":
                        v163 := GenS16{F0: ("" + "hello")}
                        _ = v163
                    case "":
                        v164 := GenS3{F0: v4}
                        _ = v164
                    default:
                        if (495 >= 32767) {
                            for fi165 := 0; fi165 < 2; fi165++ {
                                fmt.Println(v162)
                                _ = fi165
                            }
                        } else {
                            var v166 int64 = -1
                            _ = v166
                        }
                    }
                    _ = v162
                }
                if false {
                    switch ((v140 + (("" + (v140 + (v140 + (v140 + (v140 + v140))))) + v140)) + (v140 + "foo bar")) {
                    case "a":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                } else {
                    var v167 int16 = -32768
                    if len(v1) > 0 {
                        v1[0] = v139
                    }
                    fmt.Println(v2, v140)
                    _ = v167
                }
                _ = v160
            }
            switch 127 {
            case -1:
                var v168 int = 2147483647
                _ = v168
            case 0:
                v2 += (v138 | v138)
            case -2147483648:
                var v169 int = 32767
                _ = v169
            }
            for fi170 := 0; fi170 < 14; fi170++ {
                var v171 int = (fi170 | v139)
                v139 += (fi147 * 2147483647)
                var v172 float64 = ((-2.25 / -2.25) * 0.1)
                _ = fi170
                _ = v171
                _ = v172
            }
            _ = fi147
        }
    default:
        if ((v6 / ((0) | 1)) > (^(v6))) {
            var v173 uint8 = 152
            _ = v173
        }
    }
    fmt.Println(v139, v4)
    switch "a" {
    case "":
        fmt.Println(v3)
    case "\n\t":
        for fi174 := 0; fi174 < 16; fi174++ {
            fmt.Println(v6)
            _ = fi174
        }
    case "a":
        v6 = (v6 + v6)
    }
    for fi175 := 0; fi175 < 12; fi175++ {
        if len(v1) > 0 {
            v1[0] = v139
        }
        _ = fi175
    }
    v176 := GenS17{F0: v3}
    if len(v1) > 0 {
        v1[0] = v139
    }
    for fi177 := 0; fi177 < 13; fi177++ {
        for fi178 := 0; fi178 < 14; fi178++ {
            var v179 uint32 = (^(v2))
            var v180 int8 = v3
            var v181 uint64 = v4
            var v182 uint32 = v179
            _ = fi178
            _ = v179
            _ = v180
            _ = v181
            _ = v182
        }
        _ = fi177
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v138:", v138)
    fmt.Println("v139:", v139)
    fmt.Println("v140:", v140)
    fmt.Println("v141:", v141)
    fmt.Println("v145:", v145)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v138
    _ = v139
    _ = v140
    _ = v141
    _ = v145
    _ = v176
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
