package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 bool; F2 int8 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16; F1 int8 }
type GenS7 struct { F0 uint64; F1 int32; F2 int64 }
type GenS8 struct { F0 uint8; F1 int16 }
type GenS9 struct { F0 uint32; F1 uint16 }
type GenS10 struct { F0 string; F1 int8; F2 int8 }
func genFn0(p0 string, p1 int) int16 {
    v5 := []int{(p1 * 1)}
    switch p0 {
    case "hello":
        p1 += p1
    case "":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    fmt.Println(p0, p1)
    for fi6 := 0; fi6 < 15; fi6++ {
        if false {
            var v7 float32 = (1e10 - (-2.25 + 3.14159))
            _ = v7
        } else {
            for fi8 := 0; fi8 < 4; fi8++ {
                var v9 int = fi6
                switch p0 {
                case "a":
                    var v10 int32 = -128
                    _ = v10
                case "hello":
                    p0 = (p0 + "\n\t")
                case "":
                    p0 = p0
                }
                _ = fi8
                _ = v9
            }
            var v11 int = -2147483648
            _ = v11
        }
        v5 = append(v5, fi6)
        var v12 string = ""
        _ = fi6
        _ = v12
    }
    switch 32768 {
    case 345:
        for fi13 := 0; fi13 < 11; fi13++ {
            fmt.Println(p1, p0, p0)
            for fi14 := 0; fi14 < 14; fi14++ {
                if (false || true) {
                    if ((true && (!((3.14159 + 3.14159) >= (-2.25 / (1.5 - float32((1e10 / (1.5 - (0.0 - 0.0))))))))) || true) {
                        if (!(false && (false || (0.1 < 3.14159)))) {
                            if ((((3.14159 - 1e10) / (1.5 + 1e10)) + 0.1) != 3.14159) {
                                v15 := GenS1{F0: 2}
                                p1 = 4294967295
                                if (42 >= 2) {
                                    var v16 float32 = 0.1
                                    var v17 uint16 = 42
                                    var v18 int = (^(p1))
                                    _ = v16
                                    _ = v17
                                    _ = v18
                                } else {
                                    v19 := GenS1{F0: 230}
                                    var v20 int64 = -9223372036854775808
                                    var v21 int32 = -526
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                }
                                _ = v15
                            }
                            p1 += (fi14 - 32767)
                            v22 := GenS2{F0: "hello", F1: true, F2: 127}
                            _ = v22
                        } else {
                            _ = len(v5)
                            v5 = append(v5, fi13)
                        }
                        var v23 bool = true
                        _ = v23
                    }
                    var v24 int8 = -128
                    var v25 uint32 = 65535
                    _ = v24
                    _ = v25
                }
                v5 = append(v5, (p1 & 127))
                _ = fi14
            }
            switch (fi13 ^ p1) {
            case 2147483647:
                if len(v5) > 0 {
                    v5[0] = (fi13 << ((995) & 63))
                }
            default:
                var v26 uint16 = 42
                _ = v26
            }
            _ = fi13
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return -1
}
func genFn1(p0 int32, p1 int16, p2 uint16) uint16 {
    var v27 int32 = 334
    v27 -= v27
    v28 := []int{0}
    switch "a" {
    case "":
        fmt.Println(p0, p1, p0)
    case "foo bar":
        var v29 float64 = 0.0
        _ = v29
    case "\n\t":
        if false {
            v30 := GenS3{F0: -1, F1: (p2 + p2)}
            if ((p0 ^ p0) != v27) {
                var v31 string = "\n\t"
                _ = v31
            }
            _ = v30
        }
    default:
        var v32 float32 = 1.5
        _ = v32
    }
    v33 := GenS2{F0: ("hello" + "\n\t"), F1: (256 == 278), F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    _ = v28
    _ = v33
    return p2
}

func main() {
    var v1 uint32 = 366
    v2 := make(map[uint32]uint32)
    var v3 string = ((("\n\t" + "") + "") + "a")
    v4 := []int{128, 283}
    fmt.Println(v1, v3)
    v34 := GenS1{F0: 115}
    switch 128 {
    case -984:
        switch (v3 + v3) {
        case "\n\t":
            v1 += v1
        default:
            if (648 == 655) {
                var v35 int64 = 42
                v35 += (v35 | v35)
                var v36 uint8 = 2
                _ = v35
                _ = v36
            }
        }
    default:
        v4 = append(v4, 32768)
    }
    switch v3 {
    case "\n\t":
        if true {
            v37 := GenS3{F0: -128, F1: genFn1(-1, -207, 2)}
            _ = v37
        }
    case "foo bar":
        v3 = ("\n\t" + "\n\t")
    case "a":
        for fi38 := 0; fi38 < 14; fi38++ {
            var v39 uint16 = 684
            for fi40 := 0; fi40 < 11; fi40++ {
                if len(v4) > 0 {
                    v4[0] = (fi40 >> ((695) & 63))
                }
                v41 := GenS4{F0: 127, F1: genFn0(v3, (^(fi40)))}
                if (-2.25 <= 3.14159) {
                    switch (("a" + v3) + v3) {
                    case "foo bar":
                        if true {
                            v3 = v3
                            v39 = uint16(v1)
                            var v42 int64 = 140
                            var v43 int32 = 2147483647
                            switch v3 {
                            case "":
                                v43 = (v43 & -1)
                            case "hello":
                                v2[(v1 - uint32(v39))] = (v1 << ((1) & 63))
                            case "foo bar":
                                v1 += v1
                                if true {
                                    if true {
                                        v44 := GenS3{F0: -128, F1: v39}
                                        v45 := GenS1{F0: uint8(v39)}
                                        _ = v44
                                        _ = v45
                                    } else {
                                        fmt.Println(v43, v43, fi38)
                                        var v46 int = fi38
                                        _ = v46
                                    }
                                } else {
                                    var v47 uint = 2
                                    switch "hello" {
                                    case "foo bar":
                                        v43 -= (v43 | v43)
                                    case "":
                                        if true {
                                            var v48 int64 = (v42 % ((v42) | 1))
                                            fmt.Println(v43)
                                            v48 += v42
                                            v49 := GenS4{F0: -128, F1: 1000}
                                            _ = v48
                                            _ = v49
                                        } else {
                                            var v50 uint16 = uint16(v1)
                                            var v51 uint64 = 823
                                            v52 := GenS5{F0: (v50 % ((78) | 1))}
                                            _ = v50
                                            _ = v51
                                            _ = v52
                                        }
                                    case "\n\t":
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    default:
                                        if (!(255 <= 42)) {
                                            if (4 <= 1) {
                                                fmt.Println(v47, v43, v3)
                                            }
                                            v3 = "\n\t"
                                        }
                                    }
                                    var v53 uint = 817
                                    _ = v47
                                    _ = v53
                                }
                                v1 = (v1 + 0)
                            default:
                                v4 = append(v4, 9223372036854775807)
                            }
                            _ = v42
                            _ = v43
                        }
                    case "":
                        _ = len(v4)
                    case "a":
                        if true {
                            v1 = (v1 + v1)
                            v1 -= (v1 - (v1 & v1))
                            var v54 float64 = (0.0 * 1e10)
                            _ = v54
                        } else {
                            var v55 float64 = -2.25
                            if true {
                                if (65 >= 60) {
                                    v2[(v1 ^ (v1 + (v1 / ((v1) | 1))))] = (v1 + v1)
                                    delete(v2, (v1 ^ (v1 % ((v1) | 1))))
                                    v3 = v3
                                } else {
                                    v55 = v55
                                    v39 -= (v39 % (((v39 & (v39 | (v39 | (v39 << ((318) & 63)))))) | 1))
                                }
                                switch v3 {
                                case "foo bar":
                                    if (false || true) {
                                        if (genFn0(("" + v3), -839) != 595) {
                                            switch ("hello" + v3) {
                                            case "\n\t":
                                                if false {
                                                    v1 = 42
                                                    var v56 int64 = 0
                                                    _ = v56
                                                } else {
                                                    switch (fi38 << ((1) & 63)) {
                                                    case 927:
                                                        v57 := GenS5{F0: v39}
                                                        _ = v57
                                                    case 9223372036854775807:
                                                        v39 -= v39
                                                    case 300:
                                                        v1 = (^(v1))
                                                    default:
                                                        var v58 uint8 = 255
                                                        _ = v58
                                                    }
                                                }
                                                var v59 uint64 = 639
                                                v55 = v55
                                                _ = v59
                                            default:
                                                if false {
                                                    var v60 uint8 = 128
                                                    _ = v60
                                                }
                                            }
                                        }
                                        if (0.1 <= -2.25) {
                                            var v61 uint32 = (v1 & v1)
                                            var v62 uint8 = 0
                                            fmt.Println(v55)
                                            _ = v61
                                            _ = v62
                                        } else {
                                            v1 += (v1 - (^(v1)))
                                            v4 = append(v4, (fi40 * (^(fi38))))
                                        }
                                        switch 9223372036854775807 {
                                        case 127:
                                            fmt.Println(fi38)
                                        case 0:
                                            v63 := GenS1{F0: 157}
                                            _ = v63
                                        case -819:
                                            v2[(v1 - v1)] = v1
                                        }
                                    }
                                }
                            } else {
                                var v64 int16 = 32767
                                if false {
                                    v3 = (v3 + (v3 + "a"))
                                    if len(v4) > 0 {
                                        v4[0] = (fi40 + fi40)
                                    }
                                }
                                v64 -= 32767
                                _ = v64
                            }
                            v3 = "hello"
                            var v65 uint64 = 334
                            _ = v55
                            _ = v65
                        }
                    default:
                        var v66 uint64 = uint64(v1)
                        _ = v66
                    }
                }
                _ = fi40
                _ = v41
            }
            _ = fi38
            _ = v39
        }
    default:
        v4 = append(v4, -9223372036854775808)
    }
    fmt.Println(v3, v1)
    if true {
        v3 = v3
        for fi67 := 0; fi67 < 8; fi67++ {
            for fi68 := 0; fi68 < 12; fi68++ {
                var v69 int64 = 42
                delete(v2, v1)
                var v70 int = (fi68 & (^(fi67)))
                if (2 <= 1) {
                    switch -448 {
                    case -607:
                        v71 := v2[(v1 << ((289) & 63))]
                        _ = v71
                    }
                    if (!true) {
                        fmt.Println(v3)
                        if ((true || (760 >= 4294967295)) && (1e10 > float64(float32(0.0)))) {
                            fmt.Println(v70, v1)
                            if false {
                                var v72 uint64 = 401
                                v70 -= 4294967295
                                v72 = (v72 / (((^(v72))) | 1))
                                if (0 >= 547) {
                                    if (((-2.25 <= ((-2.25 / 1.0) * 1.5)) && false) || (!(42 >= 0))) {
                                        v72 = (v72 & (v72 - (v72 % (((v72 - (v72 - (v72 << ((18446744073709551615) & 63))))) | 1))))
                                        var v73 int = (fi67 / ((v70) | 1))
                                        _ = v73
                                    }
                                    v3 = ""
                                    v74 := GenS5{F0: 255}
                                    _ = v74
                                }
                                _ = v72
                            } else {
                                v3 = v3
                                v75 := GenS6{F0: -1, F1: 42}
                                if false {
                                    if ((false && (!(v3 != v3))) && true) {
                                        v1 -= (v1 * v1)
                                    }
                                } else {
                                    v3 = v3
                                    if (false && false) {
                                        var v76 int8 = 0
                                        var v77 float32 = (((1e10 * 0.1) + 0.0) + 1.5)
                                        var v78 int8 = (v76 << ((2) & 63))
                                        _ = v76
                                        _ = v77
                                        _ = v78
                                    } else {
                                        v79 := GenS2{F0: v3, F1: true, F2: 66}
                                        v4 = append(v4, fi68)
                                        v69 -= v69
                                        var v80 int64 = v69
                                        _ = v79
                                        _ = v80
                                    }
                                }
                                _ = v75
                            }
                            var v81 uint = 255
                            _ = v81
                        }
                    } else {
                        v82 := GenS5{F0: 1}
                        _ = v82
                    }
                    delete(v2, v1)
                }
                v83 := GenS7{F0: 42, F1: 511, F2: 235}
                _ = fi68
                _ = v69
                _ = v70
                _ = v83
            }
            var v84 string = v3
            _ = fi67
            _ = v84
        }
    } else {
        for fi85 := 0; fi85 < 14; fi85++ {
            v86 := v2[v1]
            if len(v4) > 0 {
                v4[0] = fi85
            }
            for fi87 := 0; fi87 < 16; fi87++ {
                v1 = v1
                _ = fi87
            }
            _ = fi85
            _ = v86
        }
        for fi88 := 0; fi88 < 14; fi88++ {
            var v89 int32 = 116
            for fi90 := 0; fi90 < 7; fi90++ {
                v1 -= 4294967295
                v1 = v1
                v1 = v1
                _ = fi90
            }
            if false {
                for fi91 := 0; fi91 < 14; fi91++ {
                    _ = len(v4)
                    fmt.Println(fi91)
                    _ = fi91
                }
                v1 -= 2147483648
            }
            _ = fi88
            _ = v89
        }
        fmt.Println(v1, v3)
    }
    v1 = 65535
    if false {
        v92 := GenS8{F0: 255, F1: 32767}
        v4 = append(v4, -2147483648)
        _ = v92
    } else {
        if false {
            for fi93 := 0; fi93 < 17; fi93++ {
                var v94 int = (fi93 / (((fi93 >> ((404) & 63))) | 1))
                if false {
                    switch (v3 + v3) {
                    case "a":
                        var v95 uint = 1
                        _ = v95
                    case "foo bar":
                        var v96 uint8 = 42
                        _ = v96
                    default:
                        var v97 int8 = 24
                        _ = v97
                    }
                    for fi98 := 0; fi98 < 6; fi98++ {
                        if (2 != 255) {
                            if (!(0 == 469)) {
                                if true {
                                    var v99 int = v94
                                    if (9223372036854775807 != 797) {
                                        v100 := GenS5{F0: 65535}
                                        _ = v100
                                    }
                                    if false {
                                        fmt.Println(v3)
                                        _ = len(v4)
                                    } else {
                                        switch "foo bar" {
                                        case "":
                                            fmt.Println(fi93, v1, fi98)
                                        case "hello":
                                            fmt.Println(v1, fi93)
                                        case "a":
                                            v99 += (fi93 / ((9223372036854775807) | 1))
                                        default:
                                            var v101 float64 = float64(-2.25)
                                            v102 := GenS9{F0: (v1 ^ 1), F1: 0}
                                            fmt.Println(fi93, fi98)
                                            _ = v101
                                            _ = v102
                                        }
                                        v103 := GenS1{F0: 163}
                                        if (42 <= 922) {
                                            if len(v4) > 0 {
                                                v4[0] = v94
                                            }
                                            v104 := GenS5{F0: 32767}
                                            _ = v104
                                        } else {
                                            if ((702 >= 42) || true) {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    v105 := GenS1{F0: 1}
                                                    _ = v105
                                                default:
                                                    v1 = (v1 - v1)
                                                }
                                                v106 := GenS3{F0: 1, F1: 443}
                                                _ = v106
                                            } else {
                                                switch v3 {
                                                case "foo bar":
                                                    var v107 uint8 = 152
                                                    _ = v107
                                                case "a":
                                                    var v108 uint16 = 2
                                                    _ = v108
                                                default:
                                                    var v109 string = v3
                                                    _ = v109
                                                }
                                            }
                                            delete(v2, v1)
                                            v99 -= v99
                                            _ = len(v2)
                                        }
                                        _ = v103
                                    }
                                    var v110 uint32 = v1
                                    var v111 string = v3
                                    _ = v110
                                    _ = v111
                                    _ = v99
                                } else {
                                    if len(v4) > 0 {
                                        v4[0] = fi93
                                    }
                                    v3 = "\n\t"
                                    delete(v2, 2)
                                    v2[2] = (^(v1))
                                    fmt.Println(v94, v94, v3)
                                    fmt.Println(fi93)
                                }
                                if true {
                                    v94 = v94
                                    fmt.Println(v94, fi93, v1)
                                    v1 = 973
                                    v1 = (v1 ^ (v1 << ((255) & 63)))
                                    var v112 int8 = 0
                                    v113 := GenS2{F0: v3, F1: true, F2: 127}
                                    _ = v112
                                    _ = v113
                                }
                                if (false || true) {
                                    v114 := GenS10{F0: ((v3 + v3) + v3), F1: -27, F2: -1}
                                    _ = v114
                                } else {
                                    var v115 float64 = 1.5
                                    v1 = 42
                                    if len(v4) > 0 {
                                        v4[0] = v94
                                    }
                                    if false {
                                        v1 += (^(v1))
                                        if (!true) {
                                            var v116 uint8 = 0
                                            _ = v116
                                        } else {
                                            if len(v4) > 0 {
                                                v4[0] = 128
                                            }
                                            v4 = append(v4, v94)
                                        }
                                        v117 := GenS1{F0: 2}
                                        delete(v2, 4294967295)
                                        var v118 bool = true
                                        var v119 string = v3
                                        _ = v117
                                        _ = v118
                                        _ = v119
                                    }
                                    delete(v2, v1)
                                    v3 = (v3 + (v3 + v3))
                                    _ = v115
                                }
                            } else {
                                var v120 string = v3
                                _ = v120
                            }
                            switch (v94 - -659) {
                            case -814:
                                v1 = (v1 >> ((1) & 63))
                            case 1:
                                switch fi93 {
                                case 32767:
                                    v121 := v2[v1]
                                    _ = v121
                                default:
                                    v94 = -2147483648
                                }
                                if false {
                                    fmt.Println(fi93, fi93, fi98)
                                    var v122 uint8 = 42
                                    v123 := GenS3{F0: 100, F1: 32767}
                                    _ = v122
                                    _ = v123
                                }
                            }
                            v124 := GenS9{F0: (v1 % ((1) | 1)), F1: 65535}
                            if (!false) {
                                if false {
                                    v1 -= v1
                                } else {
                                    v1 = v1
                                    switch fi93 {
                                    case 888:
                                        v1 -= v1
                                    case -1:
                                        var v125 string = ""
                                        _ = len(v2)
                                        var v126 uint16 = 2
                                        _ = v125
                                        _ = v126
                                    }
                                    if true {
                                        v127 := GenS3{F0: 42, F1: genFn1(-2147483648, 628, 42)}
                                        fmt.Println(v94)
                                        _ = v127
                                    }
                                }
                                v128 := GenS6{F0: 944, F1: -128}
                                var v129 string = "hello"
                                _ = v128
                                _ = v129
                            } else {
                                var v130 float64 = 0.1
                                v131 := GenS4{F0: 15, F1: -32768}
                                _ = v130
                                _ = v131
                            }
                            _ = v124
                        }
                        var v132 float64 = 1.5
                        var v133 float64 = v132
                        _ = fi98
                        _ = v132
                        _ = v133
                    }
                }
                _ = fi93
                _ = v94
            }
        }
        v134 := GenS8{F0: 127, F1: 0}
        v135 := GenS9{F0: (v1 | 2), F1: 255}
        _ = v134
        _ = v135
    }
    v136 := GenS4{F0: 42, F1: 284}
    v3 = v3
    var v137 uint32 = (v1 % ((2147483647) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v137:", v137)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v136
    _ = v137
    _ = v2
    _ = v3
    _ = v34
    _ = v4
}
