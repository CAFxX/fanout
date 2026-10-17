package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 uint64; F2 string }
type GenS4 struct { F0 uint64; F1 int8; F2 int64 }
type GenS5 struct { F0 int16; F1 uint32; F2 bool }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint16; F1 uint }
type GenS9 struct { F0 bool; F1 uint32; F2 int }
type GenS10 struct { F0 uint64; F1 string }
type GenS11 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS12 struct { F0 int; F1 string; F2 uint }
type GenS13 struct { F0 int32; F1 int; F2 uint8 }
type GenS14 struct { F0 bool; F1 int16 }
type GenS15 struct { F0 int16; F1 int64; F2 int8 }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 int32; F1 bool }
type GenS18 struct { F0 int16 }
func genFn0() string {
    v7 := []int{1}
    var v8 int64 = 9223372036854775807
    if false {
        if len(v7) > 0 {
            v7[0] = -834
        }
        if false {
            for fi9 := 0; fi9 < 4; fi9++ {
                v10 := GenS1{F0: false}
                _ = fi9
                _ = v10
            }
            if (-32768 != 32767) {
                if false {
                    if len(v7) > 0 {
                        v7[0] = 42
                    }
                    v8 += (v8 * v8)
                } else {
                    v11 := GenS1{F0: (546 < 42)}
                    for fi12 := 0; fi12 < 16; fi12++ {
                        v8 -= (v8 << ((1) & 63))
                        var v13 uint = 283
                        var v14 uint8 = 199
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    _ = v11
                }
            } else {
                v7 = append(v7, 617)
                fmt.Println(v8)
                var v15 uint = 19
                _ = v15
            }
            v8 += (v8 << ((913) & 63))
        }
        if len(v7) > 0 {
            v7[0] = -2147483648
        }
    }
    _ = v7
    _ = v8
    return "hello"
}
func genFn1(p0 uint32, p1 string) int64 {
    v16 := GenS2{F0: 929}
    var v17 uint16 = 549
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return -467
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 0
    v3 := []uint32{65535, 735, 2147483647, 2147483648}
    v4 := []int{-1, 32768, 2147483647, 42}
    v5 := make(map[int]bool)
    v6 := []uint32{857, 4294967295, 319, 0}
    switch ("hello" + ("foo bar" + "\n\t")) {
    case "":
        fmt.Println(v2)
    case "\n\t":
        for fi18 := 0; fi18 < 14; fi18++ {
            v2 -= (v2 / ((-64) | 1))
            _ = fi18
        }
    case "foo bar":
        v2 = (v2 | 100)
    default:
        if true {
            v19 := GenS3{F0: -633, F1: 4294967295, F2: ""}
            if ((!false) && false) {
                v2 += (v2 << ((255) & 63))
                var v20 uint8 = 66
                _ = v20
            } else {
                if (true || (128 != 1)) {
                    fmt.Println(v2)
                    switch "\n\t" {
                    case "a":
                        for fi21 := 0; fi21 < 1; fi21++ {
                            v22 := GenS2{F0: (fi21 + fi21)}
                            _ = fi21
                            _ = v22
                        }
                    case "foo bar":
                        if true {
                            v2 += v2
                            for fi23 := 0; fi23 < 2; fi23++ {
                                var v24 uint64 = 662
                                _ = fi23
                                _ = v24
                            }
                        }
                    case "":
                        if (!(!(2 == 0))) {
                            var v25 int = 844
                            var v26 uint16 = 0
                            _ = v25
                            _ = v26
                        } else {
                            if (2147483647 >= 42) {
                                v27 := GenS1{F0: ((true || (!((874 <= 1) || false))) && (true && (0 > 2147483647)))}
                                var v28 uint32 = 629
                                _ = v27
                                _ = v28
                            }
                        }
                    }
                    var v29 uint = 18446744073709551615
                    _ = v29
                }
                for fi30 := 0; fi30 < 1; fi30++ {
                    fmt.Println(v2, v2)
                    var v31 string = (("" + "foo bar") + "\n\t")
                    for fi32 := 0; fi32 < 2; fi32++ {
                        var v33 float32 = 0.0
                        v2 = int8(fi30)
                        v33 = ((1.5 * (-2.25 / v33)) / (-2.25 - 3.14159))
                        _ = fi32
                        _ = v33
                    }
                    _ = fi30
                    _ = v31
                }
            }
            var v34 uint8 = 158
            _ = v19
            _ = v34
        } else {
            for fi35 := 0; fi35 < 2; fi35++ {
                fmt.Println(fi35, v2)
                _ = fi35
            }
            v2 = (v2 - (v2 / (((v2 & (v2 << ((255) & 63)))) | 1)))
        }
    }
    v2 = (-(v2))
    v2 += (^(v2))
    v2 += v2
    v36 := GenS2{F0: -859}
    var v37 uint32 = 1
    var v38 uint32 = v37
    v39 := GenS4{F0: 430, F1: v2, F2: 1}
    v37 -= (^(v38))
    v40 := GenS5{F0: 1000, F1: (v38 - 42), F2: (v2 != (v2 | -57))}
    if true {
        _ = len(v4)
        for fi41 := 0; fi41 < 10; fi41++ {
            var v42 float32 = 3.14159
            if (-9223372036854775808 > -132) {
                if false {
                    delete(v5, fi41)
                    if (0.1 == ((3.14159 * -2.25) - (-2.25 * (0.0 - 1.5)))) {
                        var v43 float64 = ((float64(float32(((1.5 / 1.0) * (3.14159 / 0.1)))) + (1.5 - (float64(v42) * 3.14159))) - 3.14159)
                        v37 = (v37 / (((v37 - v38)) | 1))
                        if (-32768 == 1) {
                            var v44 uint8 = 127
                            var v45 float64 = float64((0.0 - v42))
                            _ = v44
                            _ = v45
                        }
                        _ = v43
                    } else {
                        v46 := GenS6{F0: 0}
                        _ = v46
                    }
                }
                if ("hello" == "foo bar") {
                    v47 := GenS7{F0: v2}
                    fmt.Println(v38)
                    v42 = 1e10
                    _ = v47
                } else {
                    if len(v3) > 0 {
                        v3[0] = (v37 * 42)
                    }
                }
            } else {
                if false {
                    v37 = v38
                } else {
                    v37 -= (v38 << ((255) & 63))
                    if (int32(v38) != 2147483647) {
                        if false {
                            if len(v6) > 0 {
                                v6[0] = v38
                            }
                            for fi48 := 0; fi48 < 14; fi48++ {
                                var v49 uint64 = 518
                                v50 := v5[(fi41 % ((fi41) | 1))]
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                _ = fi48
                                _ = v49
                                _ = v50
                            }
                            v3 = append(v3, 65535)
                        } else {
                            if true {
                                switch "foo bar" {
                                case "hello":
                                    v5[fi41] = false
                                case "":
                                    _ = len(v6)
                                default:
                                    fmt.Println(v37)
                                }
                            }
                            var v51 int8 = v2
                            if true {
                                var v52 int = 21
                                fmt.Println(v42, v42, fi41)
                                var v53 uint16 = 762
                                _ = v52
                                _ = v53
                            } else {
                                var v54 uint16 = 65535
                                v55 := v5[fi41]
                                v2 -= -52
                                _ = v54
                                _ = v55
                            }
                            _ = v51
                        }
                        if len(v3) > 0 {
                            v3[0] = v38
                        }
                    }
                    switch "a" {
                    case "\n\t":
                        if true {
                            v2 = 1
                            _ = len(v1)
                            var v56 uint8 = 1
                            _ = v56
                        } else {
                            for fi57 := 0; fi57 < 5; fi57++ {
                                switch "\n\t" {
                                case "hello":
                                    v2 += (v2 & v2)
                                    var v58 string = "foo bar"
                                    _ = v58
                                case "":
                                    var v59 float64 = -2.25
                                    _ = v59
                                default:
                                    switch (fi57 ^ fi57) {
                                    case -2147483648:
                                        if (!true) {
                                            var v60 int16 = 1
                                            if false {
                                                v61 := GenS4{F0: 89, F1: 1, F2: -842}
                                                _ = v61
                                            }
                                            _ = v60
                                        } else {
                                            if ("\n\t" != ((("a" + ("hello" + ("hello" + ""))) + "") + "foo bar")) {
                                                var v62 int = (-(fi41))
                                                _ = v62
                                            } else {
                                                fmt.Println(v2, fi41)
                                                if (4294967295 < 18446744073709551615) {
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                    v63 := GenS8{F0: 42, F1: 18446744073709551615}
                                                    if (1 < int64(v38)) {
                                                        var v64 uint32 = (^(v38))
                                                        _ = v64
                                                    }
                                                    _ = v63
                                                }
                                                v42 -= (v42 / ((v42 - v42) * v42))
                                            }
                                            v2 += v2
                                            fmt.Println(v2, v2, v37)
                                            switch "foo bar" {
                                            case "a":
                                                switch fi57 {
                                                case 811:
                                                    if false {
                                                        var v65 int32 = 0
                                                        var v66 uint = 0
                                                        v67 := v1[v38]
                                                        _ = v65
                                                        _ = v66
                                                        _ = v67
                                                    } else {
                                                        delete(v5, (fi41 & fi57))
                                                    }
                                                default:
                                                    var v68 int64 = int64(v2)
                                                    _ = v68
                                                }
                                            }
                                        }
                                    case -1:
                                        switch 1 {
                                        case -2147483648:
                                            if true {
                                                v37 -= v38
                                                _ = len(v6)
                                            }
                                        }
                                    case -780:
                                        fmt.Println(v2, v2)
                                    }
                                }
                                _ = fi57
                            }
                            v37 = v38
                            var v69 int = (fi41 >> ((796) & 63))
                            _ = v69
                        }
                    case "a":
                        v5[(fi41 - 2147483647)] = (!((v2 >> ((1) & 63)) <= v2))
                    case "foo bar":
                        if true {
                            var v70 int64 = 1
                            v42 = float32((1e10 / (1.5 - float64(v42))))
                            _ = v70
                        } else {
                            v37 = v38
                        }
                    default:
                        var v71 uint64 = 0
                        _ = v71
                    }
                }
                var v72 uint64 = 2
                _ = v72
            }
            if true {
                switch "a" {
                case "a":
                    if (1 != 825) {
                        var v73 uint32 = v38
                        _ = v73
                    }
                }
                v37 = 2
            } else {
                if (true || (1 != 42)) {
                    v74 := GenS6{F0: -762}
                    v75 := GenS9{F0: (false && false), F1: 2147483647, F2: fi41}
                    v42 += v42
                    _ = v74
                    _ = v75
                }
                var v76 int8 = (^(v2))
                _ = v76
            }
            _ = fi41
            _ = v42
        }
        v77 := GenS10{F0: 250, F1: "foo bar"}
        _ = v77
    } else {
        switch 0 {
        case 32768:
            if true {
                v38 -= (v37 ^ (v37 << ((1) & 63)))
                switch ("foo bar" + "a") {
                case "\n\t":
                    for fi78 := 0; fi78 < 15; fi78++ {
                        var v79 uint32 = 0
                        if false {
                            fmt.Println(v79, fi78, v2)
                        }
                        if true {
                            _ = len(v1)
                            v80 := GenS1{F0: true}
                            _ = v80
                        } else {
                            if true {
                                var v81 int8 = (v2 - v2)
                                delete(v5, (^(fi78)))
                                _ = v81
                            } else {
                                v79 = (^(v79))
                                for fi82 := 0; fi82 < 7; fi82++ {
                                    var v83 float64 = 0.0
                                    _ = fi82
                                    _ = v83
                                }
                                v2 = (v2 >> ((2) & 63))
                            }
                            _ = len(v4)
                        }
                        _ = fi78
                        _ = v79
                    }
                case "hello":
                    if true {
                        var v84 bool = true
                        v84 = true
                        fmt.Println(v38)
                        _ = v84
                    } else {
                        if ((v2 ^ v2) > (v2 | 0)) {
                            switch "" {
                            case "\n\t":
                                delete(v1, 1)
                            }
                            for fi85 := 0; fi85 < 20; fi85++ {
                                v86 := GenS11{F0: 0, F1: 0, F2: (v38 + v38)}
                                if (!false) {
                                    var v87 int16 = 1000
                                    var v88 int64 = -9223372036854775808
                                    _ = v87
                                    _ = v88
                                }
                                _ = fi85
                                _ = v86
                            }
                            var v89 bool = true
                            _ = v89
                        } else {
                            switch "a" {
                            case "\n\t":
                                if (false || (32767 >= 32767)) {
                                    _ = len(v6)
                                    var v90 string = ("a" + "")
                                    var v91 int = 0
                                    _ = v90
                                    _ = v91
                                }
                            case "":
                                var v92 int16 = -227
                                _ = v92
                            }
                            v5[-2147483648] = (!((-9223372036854775808 <= genFn1(v37, "foo bar")) && false))
                            v2 = v2
                        }
                        switch ("a" + ("\n\t" + "")) {
                        case "\n\t":
                            switch 127 {
                            case -9223372036854775808:
                                if true {
                                    v38 += 2
                                    v38 = (v37 - (v38 | 53))
                                    for fi93 := 0; fi93 < 7; fi93++ {
                                        switch ("foo bar" + "hello") {
                                        case "":
                                            v37 += v38
                                        case "foo bar":
                                            v2 -= v2
                                        case "\n\t":
                                            v37 = v38
                                        }
                                        var v94 int16 = 798
                                        var v95 int32 = -1
                                        _ = fi93
                                        _ = v94
                                        _ = v95
                                    }
                                } else {
                                    v96 := GenS9{F0: false, F1: 836, F2: -817}
                                    v37 = v37
                                    _ = v96
                                }
                            default:
                                v97 := v1[4294967295]
                                _ = v97
                            }
                        default:
                            if false {
                                v4 = append(v4, 2147483647)
                                for fi98 := 0; fi98 < 19; fi98++ {
                                    v99 := GenS6{F0: 32767}
                                    _ = fi98
                                    _ = v99
                                }
                            }
                        }
                        var v100 int8 = 42
                        _ = v100
                    }
                default:
                    v101 := GenS2{F0: -1}
                    _ = v101
                }
                for fi102 := 0; fi102 < 9; fi102++ {
                    var v103 int8 = v2
                    if (239 != 0) {
                        for fi104 := 0; fi104 < 17; fi104++ {
                            if false {
                                v105 := GenS12{F0: 0, F1: "hello", F2: 617}
                                _ = v105
                            } else {
                                v38 += (v38 & (^(v38)))
                            }
                            var v106 uint8 = 0
                            var v107 uint = 65535
                            fmt.Println(v37)
                            switch fi104 {
                            case -912:
                                var v108 int32 = 946
                                fmt.Println(v108, v106)
                                switch fi102 {
                                case -394:
                                    if (fi104 >= fi102) {
                                        switch int(v107) {
                                        case 32768:
                                            switch "hello" {
                                            case "":
                                                var v109 uint32 = v37
                                                var v110 int = (fi104 % ((127) | 1))
                                                v107 += v107
                                                fmt.Println(v38)
                                                v110 -= (v110 - fi102)
                                                var v111 uint8 = 49
                                                switch "foo bar" {
                                                case "a":
                                                    if len(v4) > 0 {
                                                        v4[0] = -882
                                                    }
                                                }
                                                v5[-230] = true
                                                _ = v109
                                                _ = v110
                                                _ = v111
                                            case "a":
                                                var v112 uint8 = v106
                                                _ = v112
                                            case "\n\t":
                                                v113 := GenS1{F0: true}
                                                _ = v113
                                            default:
                                                v38 = (v37 + 4294967295)
                                            }
                                        case 321:
                                            var v114 bool = false
                                            _ = v114
                                        case -318:
                                            v37 += (v38 << ((18446744073709551615) & 63))
                                        }
                                        v103 = v2
                                        v38 = 596
                                        v107 = v107
                                    } else {
                                        v115 := GenS5{F0: 42, F1: 4294967295, F2: true}
                                        _ = v115
                                    }
                                case 32767:
                                    _ = len(v4)
                                default:
                                    if true {
                                        if true {
                                            var v116 bool = (2 > 255)
                                            if (v2 < -25) {
                                                v117 := GenS7{F0: -27}
                                                fmt.Println(v103, fi104)
                                                var v118 uint8 = v106
                                                v106 -= (v106 << ((v107) & 63))
                                                var v119 bool = ((2 != uint64(v37)) && v116)
                                                fmt.Println(v103, fi102)
                                                v37 = v37
                                                _ = v117
                                                _ = v118
                                                _ = v119
                                            } else {
                                                delete(v5, (fi104 / ((fi104) | 1)))
                                                var v120 uint = v107
                                                _ = v120
                                            }
                                            var v121 uint16 = 2
                                            _ = v116
                                            _ = v121
                                        }
                                        if (false || false) {
                                            v103 += v2
                                            if true {
                                                v5[0] = (255 == 128)
                                                v122 := GenS8{F0: 255, F1: v107}
                                                _ = v122
                                            }
                                        } else {
                                            var v123 int = int(v38)
                                            if (uint64(v106) < uint64(v107)) {
                                                v106 += 2
                                            } else {
                                                v2 += (v2 - -58)
                                                v2 += 1
                                                v107 += (v107 >> (((v107 | 2)) & 63))
                                                v124 := GenS13{F0: v108, F1: (fi104 + 272), F2: (v106 << ((802) & 63))}
                                                _ = v124
                                            }
                                            _ = v123
                                        }
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        var v125 uint8 = 2
                                        var v126 float32 = (1e10 / 1.0)
                                        var v127 int16 = 42
                                        fmt.Println(v2)
                                        _ = v125
                                        _ = v126
                                        _ = v127
                                    }
                                }
                                if len(v6) > 0 {
                                    v6[0] = uint32(v107)
                                }
                                if ((v108 << (((v107 >> (((v107 + 780)) & 63))) & 63)) != (v108 / ((-195) | 1))) {
                                    if true {
                                        v128 := GenS14{F0: (!((v107 + v107) == v107)), F1: -1}
                                        _ = v128
                                    } else {
                                        v38 = v37
                                        if (((((((("hello" + (("a" + "a") + "hello")) + "") + "foo bar") + "") + ("a" + "foo bar")) + "") == "") && true) {
                                            v129 := GenS9{F0: (128 <= fi102), F1: v38, F2: fi104}
                                            _ = v129
                                        }
                                        fmt.Println(v37)
                                    }
                                    v38 += v38
                                    v106 += (^(v106))
                                }
                                _ = v108
                            default:
                                v130 := v5[847]
                                _ = v130
                            }
                            _ = fi104
                            _ = v106
                            _ = v107
                        }
                        v131 := GenS3{F0: 0, F1: 4294967295, F2: "\n\t"}
                        var v132 int8 = (v2 / (((v103 << ((4294967295) & 63))) | 1))
                        _ = v131
                        _ = v132
                    }
                    if (3.14159 == 0.1) {
                        v133 := GenS15{F0: -1, F1: -808, F2: -75}
                        var v134 float32 = 3.14159
                        v135 := GenS16{F0: ((0 != 18446744073709551615) && (false && ((v2 / (((v103 * v2)) | 1)) != v103)))}
                        _ = v133
                        _ = v134
                        _ = v135
                    }
                    _ = fi102
                    _ = v103
                }
            }
        default:
            if (false && true) {
                var v136 int = 737
                var v137 float64 = -2.25
                if (true && (-128 < (v2 * v2))) {
                    if true {
                        var v138 int32 = 0
                        var v139 int8 = -31
                        var v140 int16 = -949
                        _ = v138
                        _ = v139
                        _ = v140
                    }
                    var v141 uint32 = (v38 / (((v37 & 997)) | 1))
                    _ = v141
                } else {
                    for fi142 := 0; fi142 < 16; fi142++ {
                        if false {
                            v143 := GenS17{F0: 0, F1: ((v37 << ((814) & 63)) <= (v37 << ((4294967295) & 63)))}
                            if true {
                                var v144 int64 = 36
                                v137 += (v137 / 1e10)
                                _ = v144
                            } else {
                                v145 := GenS18{F0: 785}
                                v137 -= (v137 / float64(1.5))
                                _ = v145
                            }
                            _ = v143
                        } else {
                            var v146 float64 = ((v137 - v137) * (v137 * ((v137 * v137) / float64(-2.25))))
                            fmt.Println(v146, v136)
                            _ = v146
                        }
                        _ = fi142
                    }
                }
                _ = v136
                _ = v137
            } else {
                var v147 int = 9223372036854775807
                _ = v147
            }
        }
    }
    var v148 uint64 = 0
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v149 int = 9223372036854775807
    var v150 int64 = 1
    v1[v38] = "a"
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v148:", v148)
    fmt.Println("v149:", v149)
    fmt.Println("v150:", v150)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v148
    _ = v149
    _ = v150
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
