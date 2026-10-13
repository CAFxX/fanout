package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 int; F2 int32 }
type GenS5 struct { F0 uint; F1 int16; F2 int }
type GenS6 struct { F0 uint8; F1 bool; F2 int32 }
type GenS7 struct { F0 int16; F1 int8; F2 string }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int64; F1 uint16 }
type GenS10 struct { F0 int32; F1 uint }
type GenS11 struct { F0 int32; F1 int32; F2 uint64 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 uint; F1 int16; F2 int64 }
type GenS14 struct { F0 uint16; F1 uint }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 int16; F1 uint8 }

func main() {
    var v1 string = ""
    v2 := make(map[string]string)
    v3 := []uint32{2147483648}
    v4 := make(map[string]int)
    var v5 bool = (-9223372036854775808 > 266)
    var v6 int = 244
    fmt.Println(v6, v6, v6)
    v5 = v5
    v1 = (v1 + ((((v1 + v1) + v1) + v1) + v1))
    fmt.Println(v1, v1, v1)
    if v5 {
        v6 = (v6 % ((32768) | 1))
    } else {
        switch v1 {
        case "":
            var v7 int32 = 761
            _ = v7
        }
        var v8 float32 = 0.0
        switch v6 {
        case 127:
            var v9 uint64 = 422
            _ = v9
        case 2147483647:
            v10 := GenS1{F0: 0, F1: 120}
            _ = v10
        default:
            v1 = (v1 + ((v1 + v1) + v1))
        }
        _ = v8
    }
    switch v6 {
    case -248:
        delete(v2, v1)
    case -579:
        var v11 int32 = -2147483648
        _ = v11
    case 835:
        fmt.Println(v6, v6)
    }
    var v12 int32 = 0
    for fi13 := 0; fi13 < 18; fi13++ {
        if false {
            if v5 {
                v1 = ""
                var v14 int32 = (v12 | v12)
                v15 := GenS1{F0: 32767, F1: 2}
                _ = v14
                _ = v15
            }
            v5 = true
        }
        if len(v3) > 0 {
            v3[0] = 802
        }
        _ = fi13
    }
    var v16 uint16 = 2
    fmt.Println(v12, v5)
    var v17 float64 = 1.5
    fmt.Println(v17, v12, v5)
    for fi18 := 0; fi18 < 10; fi18++ {
        switch (v6 & v6) {
        case 42:
            var v19 bool = v5
            _ = v19
        default:
            for fi20 := 0; fi20 < 7; fi20++ {
                v5 = true
                v6 = fi18
                var v21 string = v1
                _ = fi20
                _ = v21
            }
        }
        for fi22 := 0; fi22 < 19; fi22++ {
            fmt.Println(fi18)
            var v23 float32 = (((3.14159 / 0.1) * 0.0) + (float32(v17) - 1e10))
            _ = fi22
            _ = v23
        }
        if v5 {
            for fi24 := 0; fi24 < 10; fi24++ {
                v16 = (v16 << ((81) & 63))
                var v25 bool = v5
                _ = fi24
                _ = v25
            }
            switch v1 {
            case "hello":
                fmt.Println(v12)
            case "a":
                v16 -= (v16 - v16)
            case "":
                for fi26 := 0; fi26 < 16; fi26++ {
                    switch (-(fi18)) {
                    case 185:
                        var v27 uint8 = 141
                        _ = v27
                    case 127:
                        v5 = (-1 > v12)
                    case -730:
                        _ = len(v2)
                        switch fi18 {
                        case 32767:
                            switch 0 {
                            case 0:
                                var v28 int16 = 32767
                                _ = v28
                            default:
                                var v29 string = v1
                                var v30 bool = (false && (495 <= 1))
                                _ = v29
                                _ = v30
                            }
                        default:
                            if v5 {
                                var v31 int16 = 320
                                switch fi18 {
                                case -442:
                                    if len(v3) > 0 {
                                        v3[0] = 1
                                    }
                                case 32768:
                                    v32 := GenS1{F0: (v31 * -775), F1: 143}
                                    _ = v32
                                case 32767:
                                    var v33 uint = 42
                                    _ = v33
                                default:
                                    var v34 int32 = v12
                                    _ = v34
                                }
                                v3 = append(v3, 675)
                                _ = v31
                            }
                        }
                        v3 = append(v3, 425)
                    default:
                        v35 := GenS2{F0: 220, F1: 0}
                        _ = v35
                    }
                    var v36 int32 = -1
                    _ = fi26
                    _ = v36
                }
            }
            if v5 {
                var v37 int = fi18
                switch ("hello" + (v1 + ("a" + "hello"))) {
                case "foo bar":
                    v3 = append(v3, 839)
                }
                _ = v37
            } else {
                if (!(!false)) {
                    if true {
                        var v38 int8 = 127
                        var v39 int32 = v12
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v38
                        _ = v39
                    }
                } else {
                    var v40 uint8 = 0
                    _ = v40
                }
                if false {
                    v6 += (-(v6))
                    v1 = v1
                    if false {
                        var v41 int = (fi18 >> ((1) & 63))
                        v5 = (4294967295 > 18446744073709551615)
                        v42 := GenS3{F0: 0}
                        _ = v41
                        _ = v42
                    } else {
                        if true {
                            if (!v5) {
                                v43 := GenS4{F0: 625, F1: (v6 - (fi18 % (((fi18 * (fi18 * v6))) | 1))), F2: v12}
                                _ = v43
                            } else {
                                var v44 int32 = v12
                                for fi45 := 0; fi45 < 5; fi45++ {
                                    var v46 float32 = 0.0
                                    v44 = 328
                                    _ = fi45
                                    _ = v46
                                }
                                if ((v16 >> ((65535) & 63)) < 302) {
                                    v47 := GenS5{F0: 273, F1: 1, F2: 437}
                                    _ = v47
                                }
                                _ = v44
                            }
                            if ((float32((v17 * v17)) + (0.1 / 1.0)) != (-2.25 + -2.25)) {
                                for fi48 := 0; fi48 < 2; fi48++ {
                                    if false {
                                        v49 := GenS2{F0: uint32(v16), F1: v16}
                                        var v50 string = (v1 + (v1 + v1))
                                        fmt.Println(v5, v1)
                                        var v51 int32 = 1
                                        var v52 uint32 = 1
                                        var v53 int = (^(fi48))
                                        _ = v49
                                        _ = v50
                                        _ = v51
                                        _ = v52
                                        _ = v53
                                    } else {
                                        switch v1 {
                                        case "foo bar":
                                            v54 := GenS2{F0: 42, F1: v16}
                                            _ = v54
                                        case "a":
                                            fmt.Println(v6, fi18, fi48)
                                        case "":
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        fmt.Println(fi48, v16)
                                        var v55 float64 = 1.5
                                        _ = v55
                                    }
                                    _ = fi48
                                }
                            } else {
                                _ = len(v2)
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        if false {
                            v17 += float64(0.1)
                        } else {
                            v6 = (fi18 % ((fi18) | 1))
                            for fi56 := 0; fi56 < 3; fi56++ {
                                switch v1 {
                                case "\n\t":
                                    var v57 uint64 = 2
                                    _ = v57
                                case "":
                                    v1 = "hello"
                                case "a":
                                    v58 := GenS4{F0: 2, F1: fi18, F2: v12}
                                    _ = v58
                                }
                                _ = fi56
                            }
                            var v59 int32 = v12
                            _ = v59
                        }
                        fmt.Println(v17)
                    }
                } else {
                    var v60 int8 = 7
                    for fi61 := 0; fi61 < 14; fi61++ {
                        var v62 uint8 = 255
                        if false {
                            v16 = v16
                            v63 := GenS6{F0: 0, F1: false, F2: v12}
                            _ = len(v4)
                            v64 := GenS5{F0: 969, F1: 32767, F2: (fi18 << ((152) & 63))}
                            if true {
                                var v65 uint16 = (v16 * v16)
                                if ((v60 != -72) || v5) {
                                    if (v65 <= v65) {
                                        var v66 int32 = (v12 & v12)
                                        var v67 bool = (2 < uint(v65))
                                        _ = v66
                                        _ = v67
                                    }
                                    v16 -= 720
                                } else {
                                    v3 = append(v3, 196)
                                    fmt.Println(fi18)
                                    fmt.Println(v1, v5)
                                }
                                _ = v65
                            } else {
                                var v68 int8 = -10
                                _ = len(v3)
                                _ = v68
                            }
                            _ = v63
                            _ = v64
                        } else {
                            switch "\n\t" {
                            case "a":
                                _ = len(v4)
                            case "hello":
                                v69 := GenS2{F0: 2147483647, F1: (v16 | (v16 % (((v16 & v16)) | 1)))}
                                switch (v1 + (v1 + v1)) {
                                case "hello":
                                    v70 := GenS2{F0: 2, F1: 1}
                                    _ = v70
                                case "foo bar":
                                    switch v1 {
                                    case "foo bar":
                                        var v71 int32 = -1
                                        _ = v71
                                    case "":
                                        v72 := GenS1{F0: 1000, F1: v62}
                                        _ = v72
                                    default:
                                        if v5 {
                                            var v73 int32 = (-(v12))
                                            _ = v73
                                        } else {
                                            v74 := GenS5{F0: 26, F1: -1, F2: (fi61 * fi18)}
                                            var v75 float64 = -2.25
                                            if (!true) {
                                                _ = len(v3)
                                            } else {
                                                switch "a" {
                                                case "a":
                                                    var v76 int16 = -32768
                                                    _ = v76
                                                case "\n\t":
                                                    var v77 int16 = -32768
                                                    _ = v77
                                                default:
                                                    v5 = (((1e10 / 1.0) * ((3.14159 * -2.25) * 1e10)) > 3.14159)
                                                }
                                                var v78 uint16 = v16
                                                if (false || v5) {
                                                    v17 += 1.5
                                                    var v79 int = (^(v6))
                                                    _ = v79
                                                } else {
                                                    var v80 uint32 = 291
                                                    var v81 int64 = -1
                                                    _ = len(v4)
                                                    _ = v80
                                                    _ = v81
                                                }
                                                _ = v78
                                            }
                                            _ = v74
                                            _ = v75
                                        }
                                    }
                                }
                                _ = v69
                            default:
                                var v82 int32 = 555
                                v1 = v1
                                v83 := GenS1{F0: 0, F1: (^(v62))}
                                _ = v82
                                _ = v83
                            }
                            var v84 int32 = v12
                            if (!((!v5) || v5)) {
                                if (-2147483648 == v12) {
                                    v60 += v60
                                    var v85 float32 = 1e10
                                    v86 := GenS3{F0: (v62 | 249)}
                                    _ = v85
                                    _ = v86
                                } else {
                                    var v87 uint64 = 1
                                    var v88 float32 = 0.1
                                    _ = v87
                                    _ = v88
                                }
                                var v89 int = -2147483648
                                v6 += (v89 + (fi18 >> ((616) & 63)))
                                _ = v89
                            }
                            _ = v84
                        }
                        var v90 int32 = v12
                        _ = fi61
                        _ = v62
                        _ = v90
                    }
                    if (!v5) {
                        v12 = v12
                    } else {
                        fmt.Println(v1, v5, v12)
                    }
                    _ = v60
                }
                v91 := GenS2{F0: 2147483647, F1: v16}
                _ = v91
            }
        } else {
            if (float64(0.1) < v17) {
                var v92 uint16 = (v16 * (^(v16)))
                fmt.Println(v17, v1, v5)
                if false {
                    if ("hello" < (v1 + v1)) {
                        for fi93 := 0; fi93 < 10; fi93++ {
                            v5 = ((v17 >= v17) && (float32((3.14159 + v17)) == -2.25))
                            _ = fi93
                        }
                        v94 := v2[((v1 + (v1 + v1)) + "a")]
                        _ = v94
                    } else {
                        if (v16 > (^(v16))) {
                            switch ("a" + v1) {
                            case "\n\t":
                                var v95 string = v1
                                _ = v95
                            case "foo bar":
                                var v96 uint8 = 128
                                _ = v96
                            case "":
                                v17 = 1.5
                            }
                            for fi97 := 0; fi97 < 5; fi97++ {
                                fmt.Println(v17)
                                switch v1 {
                                case "\n\t":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v2[(("\n\t" + "hello") + v1)] = "a"
                                case "":
                                    _ = len(v3)
                                default:
                                    if len(v3) > 0 {
                                        v3[0] = 4294967295
                                    }
                                }
                                _ = fi97
                            }
                        }
                    }
                    if v5 {
                        for fi98 := 0; fi98 < 13; fi98++ {
                            _ = len(v3)
                            var v99 string = ""
                            _ = fi98
                            _ = v99
                        }
                    } else {
                        if (18446744073709551615 < 805) {
                            switch v1 {
                            case "a":
                                v3 = append(v3, 2147483647)
                            case "hello":
                                var v100 float32 = 0.0
                                _ = v100
                            }
                            var v101 uint32 = 1
                            v102 := GenS5{F0: 995, F1: 1, F2: (fi18 / ((4294967295) | 1))}
                            _ = v101
                            _ = v102
                        }
                        v92 += v16
                        for fi103 := 0; fi103 < 8; fi103++ {
                            v104 := GenS7{F0: 0, F1: -128, F2: v1}
                            _ = fi103
                            _ = v104
                        }
                    }
                } else {
                    for fi105 := 0; fi105 < 2; fi105++ {
                        v12 = (v12 * -67)
                        v17 = v17
                        _ = fi105
                    }
                    for fi106 := 0; fi106 < 12; fi106++ {
                        v107 := GenS3{F0: 20}
                        switch -9223372036854775808 {
                        case 986:
                            v17 = v17
                        default:
                            switch v1 {
                            case "\n\t":
                                if false {
                                    if (((fi106 >> ((4294967295) & 63)) != 1) && v5) {
                                        switch v1 {
                                        case "\n\t":
                                            var v108 uint = 1
                                            _ = v108
                                        }
                                        v109 := GenS3{F0: 1}
                                        _ = v109
                                    } else {
                                        delete(v2, v1)
                                    }
                                    _ = len(v2)
                                    var v110 float32 = ((1e10 + (-2.25 / (((-2.25 - 3.14159) + float32((((v17 + v17) / float64(3.14159)) + 1.5))) - 1.5))) - float32(float64(1.5)))
                                    _ = v110
                                } else {
                                    v3 = append(v3, 2147483648)
                                    if v5 {
                                        v111 := GenS3{F0: 128}
                                        var v112 float64 = 1e10
                                        var v113 uint64 = 268
                                        _ = v111
                                        _ = v112
                                        _ = v113
                                    } else {
                                        _ = len(v3)
                                        v17 = -2.25
                                        v17 -= (v17 / float64((0.0 + 3.14159)))
                                    }
                                    fmt.Println(v6)
                                    if v5 {
                                        fmt.Println(v92, fi106)
                                    } else {
                                        fmt.Println(v6)
                                        v92 += v92
                                        if false {
                                            v114 := GenS8{F0: 789}
                                            if (4294967295 > 104) {
                                                var v115 int8 = 105
                                                if (int32(fi106) == v12) {
                                                    v12 = 0
                                                } else {
                                                    _ = len(v2)
                                                    v1 = ("" + (v1 + ((((v1 + (("hello" + v1) + v1)) + (v1 + v1)) + "") + v1)))
                                                    switch ((v1 + "\n\t") + v1) {
                                                    case "\n\t":
                                                        var v116 int16 = int16(fi18)
                                                        _ = v116
                                                    default:
                                                        var v117 int64 = -134
                                                        _ = v117
                                                    }
                                                }
                                                v118 := GenS9{F0: 0, F1: 567}
                                                _ = v115
                                                _ = v118
                                            }
                                            v119 := GenS10{F0: v12, F1: 65535}
                                            if false {
                                                switch v1 {
                                                case "foo bar":
                                                    v17 = float64(3.14159)
                                                case "\n\t":
                                                    var v120 uint = 18446744073709551615
                                                    _ = v120
                                                case "hello":
                                                    fmt.Println(fi18, v5, v5)
                                                }
                                            } else {
                                                v121 := GenS11{F0: (v12 - -634), F1: (v12 | (v12 % ((v12) | 1))), F2: 0}
                                                v122 := GenS12{F0: -7}
                                                _ = v121
                                                _ = v122
                                            }
                                            var v123 uint = 42
                                            _ = v114
                                            _ = v119
                                            _ = v123
                                        }
                                    }
                                    switch "" {
                                    case "foo bar":
                                        if false {
                                            var v124 int = (^(fi106))
                                            v125 := GenS13{F0: 777, F1: -1, F2: 0}
                                            _ = v124
                                            _ = v125
                                        } else {
                                            _ = len(v2)
                                        }
                                    default:
                                        var v126 bool = v5
                                        _ = v126
                                    }
                                }
                            default:
                                var v127 uint = 4294967295
                                _ = v127
                            }
                        }
                        _ = fi106
                        _ = v107
                    }
                    switch (v6 * (v6 & 2147483647)) {
                    case 785:
                        switch (v1 + (("a" + v1) + v1)) {
                        case "hello":
                            if v5 {
                                if true {
                                    var v128 int8 = 79
                                    v17 = v17
                                    var v129 uint64 = 0
                                    _ = v128
                                    _ = v129
                                }
                                fmt.Println(v12, v16)
                                var v130 uint64 = 288
                                _ = v130
                            } else {
                                if (!v5) {
                                    for fi131 := 0; fi131 < 15; fi131++ {
                                        fmt.Println(fi18, v1)
                                        if v5 {
                                            _ = len(v4)
                                        }
                                        v132 := GenS6{F0: 2, F1: ((v12 & -608) >= -1), F2: (v12 >> ((255) & 63))}
                                        if (!(1e10 < float32((v17 * (v17 + v17))))) {
                                            if (!(!v5)) {
                                                v5 = ((-32768 > -32768) && (v5 && false))
                                                var v133 uint16 = 0
                                                v134 := GenS13{F0: 486, F1: -402, F2: -1}
                                                _ = v133
                                                _ = v134
                                            } else {
                                                switch ("foo bar" + v1) {
                                                case "a":
                                                    if (-32768 > 1) {
                                                        var v135 int = (fi18 ^ (^(v6)))
                                                        _ = v135
                                                    }
                                                case "":
                                                    if false {
                                                        v2["foo bar"] = "\n\t"
                                                        switch fi131 {
                                                        case 874:
                                                            v5 = v5
                                                        case 9223372036854775807:
                                                            v17 = v17
                                                        case 32768:
                                                            switch 0 {
                                                            case 128:
                                                                var v136 int = fi131
                                                                _ = v136
                                                            case 1:
                                                                v137 := GenS9{F0: -1, F1: v92}
                                                                _ = v137
                                                            default:
                                                                fmt.Println(v1, v12, fi131)
                                                            }
                                                        }
                                                        if (float32((v17 / ((v17 / v17) * v17))) == ((float32(v17) - (float32(v17) + 1e10)) * 0.1)) {
                                                            v1 = v1
                                                            fmt.Println(v12, v92)
                                                        } else {
                                                            v5 = (675 < 0)
                                                            v1 = "hello"
                                                        }
                                                        v138 := GenS9{F0: 9223372036854775807, F1: (v92 * 65535)}
                                                        if (32767 != v16) {
                                                            switch v1 {
                                                            case "":
                                                                var v139 uint64 = 18446744073709551615
                                                                _ = v139
                                                            }
                                                            var v140 int = -596
                                                            var v141 int = v6
                                                            var v142 uint16 = (v92 >> ((4294967295) & 63))
                                                            _ = v140
                                                            _ = v141
                                                            _ = v142
                                                        } else {
                                                            fmt.Println(v5)
                                                            v92 = 42
                                                            fmt.Println(fi131, v92, v16)
                                                        }
                                                        _ = v138
                                                    } else {
                                                        switch fi18 {
                                                        case 128:
                                                            v6 = 32767
                                                        case 943:
                                                            v1 = ("a" + "")
                                                        case -700:
                                                            var v143 int64 = -1
                                                            _ = v143
                                                        default:
                                                            var v144 int = (v6 + (fi131 % ((-437) | 1)))
                                                            _ = v144
                                                        }
                                                    }
                                                case "hello":
                                                    var v145 int = (fi18 + -1)
                                                    _ = v145
                                                default:
                                                    fmt.Println(v6, v1, v92)
                                                }
                                                if false {
                                                    delete(v2, "hello")
                                                    v146 := GenS13{F0: 2, F1: 827, F2: -310}
                                                    _ = v146
                                                } else {
                                                    if (fi131 < v6) {
                                                        var v147 uint8 = 127
                                                        _ = len(v3)
                                                        var v148 uint64 = 18446744073709551615
                                                        _ = v147
                                                        _ = v148
                                                    }
                                                    v149 := GenS11{F0: v12, F1: v12, F2: 880}
                                                    v150 := GenS12{F0: -32768}
                                                    v151 := GenS14{F0: v92, F1: 518}
                                                    _ = v149
                                                    _ = v150
                                                    _ = v151
                                                }
                                                if v5 {
                                                    v152 := GenS2{F0: 964, F1: 200}
                                                    _ = v152
                                                } else {
                                                    delete(v4, ("a" + (v1 + "\n\t")))
                                                    v6 -= (v6 & 0)
                                                }
                                            }
                                            fmt.Println(v1, v12, v92)
                                        } else {
                                            if len(v3) > 0 {
                                                v3[0] = 119
                                            }
                                            v153 := GenS5{F0: 5, F1: 0, F2: (fi131 << ((209) & 63))}
                                            if v5 {
                                                if ((v17 / v17) == v17) {
                                                    if false {
                                                        v92 += v92
                                                    }
                                                    fmt.Println(v17)
                                                    switch (fi18 >> ((255) & 63)) {
                                                    case 2147483647:
                                                        if (!(0 < 0)) {
                                                            if (v5 || v5) {
                                                                v4[v1] = fi131
                                                                if v5 {
                                                                    var v154 int16 = 1
                                                                    _ = v154
                                                                }
                                                                v12 = (v12 & (-(v12)))
                                                            } else {
                                                                fmt.Println(v5, v92, fi18)
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    v16 += (v16 - 286)
                                                                default:
                                                                    var v155 uint = 1
                                                                    _ = v155
                                                                }
                                                                v156 := GenS7{F0: 1000, F1: -86, F2: "\n\t"}
                                                                switch (v6 >> ((4294967295) & 63)) {
                                                                case 128:
                                                                    v12 = v12
                                                                    var v157 uint8 = 127
                                                                    _ = v157
                                                                default:
                                                                    v158 := GenS14{F0: (v16 ^ 662), F1: 0}
                                                                    _ = v158
                                                                }
                                                                v92 += v92
                                                                _ = v156
                                                            }
                                                            var v159 uint = 65535
                                                            _ = len(v3)
                                                            var v160 uint8 = 127
                                                            if false {
                                                                _ = len(v4)
                                                                v161 := GenS2{F0: 2, F1: (v92 & v92)}
                                                                var v162 float64 = (float64(((((0.1 + 1.5) * 3.14159) - ((1e10 * ((((((-2.25 + 1.5) / -2.25) * 3.14159) * 1.5) / 0.1) * 0.1)) - (1.5 * (1e10 + float32((v17 + v17)))))) - 0.0)) / (((1e10 / v17) + -2.25) * (float64(1.5) / v17)))
                                                                _ = v161
                                                                _ = v162
                                                            }
                                                            var v163 int8 = 1
                                                            v164 := GenS9{F0: 534, F1: v16}
                                                            _ = v159
                                                            _ = v160
                                                            _ = v163
                                                            _ = v164
                                                        } else {
                                                            switch -522 {
                                                            case 127:
                                                                v165 := GenS4{F0: 1, F1: fi18, F2: (v12 * v12)}
                                                                _ = v165
                                                            case 1:
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                            default:
                                                                v166 := GenS2{F0: 2147483647, F1: 434}
                                                                _ = v166
                                                            }
                                                            fmt.Println(v17, v12, v1)
                                                        }
                                                    case 9223372036854775807:
                                                        v17 -= v17
                                                    case -296:
                                                        fmt.Println(v92)
                                                        if (1 == 633) {
                                                            switch "hello" {
                                                            case "foo bar":
                                                                v6 += (v6 % (((fi18 >> ((18446744073709551615) & 63))) | 1))
                                                            case "":
                                                                v12 += v12
                                                                if len(v3) > 0 {
                                                                    v3[0] = 4294967295
                                                                }
                                                            case "\n\t":
                                                                fmt.Println(v16)
                                                            }
                                                            if v5 {
                                                                switch fi131 {
                                                                case 427:
                                                                    v12 += (v12 ^ v12)
                                                                case 2147483647:
                                                                    fmt.Println(v6)
                                                                }
                                                            }
                                                        }
                                                        v5 = ((float32(v17) * 1.5) <= 1.5)
                                                        v167 := GenS4{F0: 65535, F1: (fi131 + (^(fi131))), F2: (v12 & 0)}
                                                        switch (v1 + v1) {
                                                        case "a":
                                                            v168 := GenS8{F0: 0}
                                                            _ = len(v2)
                                                            var v169 uint16 = (v92 & v92)
                                                            _ = v168
                                                            _ = v169
                                                        case "\n\t":
                                                            fmt.Println(v6, v5)
                                                        case "hello":
                                                            var v170 uint = 1
                                                            _ = v170
                                                        default:
                                                            if true {
                                                                v92 += (v16 ^ 586)
                                                                var v171 float32 = 0.0
                                                                _ = v171
                                                            }
                                                        }
                                                        v92 -= v92
                                                        _ = v167
                                                    default:
                                                        v172 := GenS12{F0: -32768}
                                                        _ = v172
                                                    }
                                                } else {
                                                    var v173 float64 = 3.14159
                                                    switch v1 {
                                                    case "foo bar":
                                                        v4[(("" + "hello") + "hello")] = -2147483648
                                                    case "\n\t":
                                                        fmt.Println(v17, fi18, v1)
                                                    }
                                                    if v5 {
                                                        v174 := GenS15{F0: 814}
                                                        _ = v174
                                                    }
                                                    _ = v173
                                                }
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            } else {
                                                var v175 uint16 = v92
                                                var v176 uint16 = 0
                                                fmt.Println(v16, v16, v175)
                                                _ = v175
                                                _ = v176
                                            }
                                            _ = v153
                                        }
                                        fmt.Println(v5, fi131)
                                        var v177 uint8 = 128
                                        fmt.Println(v17)
                                        _ = fi131
                                        _ = v132
                                        _ = v177
                                    }
                                } else {
                                    var v178 uint = 42
                                    switch 127 {
                                    case 1:
                                        v179 := GenS10{F0: (v12 << ((v178) & 63)), F1: (v178 << ((v178) & 63))}
                                        _ = v179
                                    case 32767:
                                        v6 = 457
                                    }
                                    _ = v178
                                }
                                v180 := GenS11{F0: v12, F1: v12, F2: 42}
                                _ = v180
                            }
                        }
                    case 96:
                        if true {
                            for fi181 := 0; fi181 < 12; fi181++ {
                                switch (^(fi18)) {
                                case 32768:
                                    switch 4294967295 {
                                    case -1:
                                        if ((53 <= 128) || true) {
                                            v182 := GenS16{F0: 32767, F1: 19}
                                            _ = v182
                                        }
                                    default:
                                        var v183 string = v1
                                        _ = v183
                                    }
                                case 9223372036854775807:
                                    var v184 float32 = 1e10
                                    _ = v184
                                }
                                v17 -= v17
                                _ = fi181
                            }
                            switch "" {
                            case "\n\t":
                                switch v6 {
                                case -2147483648:
                                    v185 := v2[v1]
                                    _ = v185
                                }
                            case "a":
                                delete(v2, v1)
                            }
                        }
                    }
                }
                _ = v92
            }
        }
        _ = fi18
    }
    for fi186 := 0; fi186 < 19; fi186++ {
        v12 = (v12 | v12)
        v6 -= fi186
        _ = fi186
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
