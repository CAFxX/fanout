package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 uint32; F1 int64; F2 uint8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int16; F1 uint; F2 int8 }
type GenS7 struct { F0 string; F1 int32 }
type GenS8 struct { F0 int; F1 int }
type GenS9 struct { F0 uint; F1 uint32; F2 int8 }
type GenS10 struct { F0 uint16; F1 int64; F2 int64 }
type GenS11 struct { F0 int64; F1 string; F2 uint16 }
type GenS12 struct { F0 int16; F1 int32; F2 string }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 int16; F1 int32; F2 uint32 }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 uint16; F1 int32 }
type GenS17 struct { F0 bool; F1 bool; F2 uint16 }
type GenS18 struct { F0 int }
type GenS19 struct { F0 int; F1 bool }
type GenS20 struct { F0 bool; F1 int32 }
type GenS21 struct { F0 uint8 }
type GenS22 struct { F0 int32 }
type GenS23 struct { F0 uint }
type GenS24 struct { F0 uint64; F1 bool }
func genFn0() int {
    switch "hello" {
    case "hello":
        if (1.5 != (1e10 / (3.14159 / (((float32((1.5 - (0.0 / 0.1))) * (((3.14159 + 0.0) - (1.5 - 0.0)) * (0.1 + ((3.14159 - -2.25) * 0.1)))) / -2.25) - 0.1)))) {
            if (1.5 >= 1.5) {
                var v4 uint8 = 11
                _ = v4
            } else {
                var v5 int = -25
                _ = v5
            }
            var v6 uint64 = 18446744073709551615
            _ = v6
        } else {
            var v7 int64 = 0
            _ = v7
        }
    case "":
        for fi8 := 0; fi8 < 8; fi8++ {
            switch "hello" {
            case "a":
                for fi9 := 0; fi9 < 14; fi9++ {
                    var v10 uint16 = 65535
                    v11 := GenS1{F0: 42, F1: 2147483647}
                    if true {
                        v10 = 255
                        v12 := GenS2{F0: 109}
                        _ = v12
                    } else {
                        var v13 int64 = 9223372036854775807
                        _ = v13
                    }
                    fmt.Println(fi9)
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            case "":
                for fi14 := 0; fi14 < 2; fi14++ {
                    fmt.Println(fi8)
                    v15 := GenS3{F0: 32767, F1: -22}
                    var v16 int64 = -598
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            case "foo bar":
                v17 := GenS4{F0: 4294967295, F1: -1, F2: 128}
                _ = v17
            default:
                for fi18 := 0; fi18 < 15; fi18++ {
                    if ((1 == 868) && (!(0 < 0))) {
                        switch "\n\t" {
                        case "":
                            if (42 <= 1) {
                                var v19 uint32 = 4294967295
                                var v20 int16 = -582
                                _ = v19
                                _ = v20
                            } else {
                                fmt.Println(fi18)
                                switch "a" {
                                case "\n\t":
                                    v21 := []int{(fi8 + fi18)}
                                    if true {
                                        fmt.Println(fi18, fi8)
                                        var v22 uint64 = 0
                                        _ = v22
                                    } else {
                                        var v23 uint = 255
                                        if true {
                                            if ((fi18 * fi8) == 4294967295) {
                                                v23 = v23
                                                v23 += v23
                                                fmt.Println(fi18, v23, v23)
                                            }
                                            v23 += 42
                                            v23 = v23
                                        }
                                        _ = v23
                                    }
                                    if len(v21) >= 2 {
                                        v21 = v21[1:len(v21)]
                                    }
                                    v24 := GenS5{F0: 255}
                                    var v25 uint64 = 4294967295
                                    switch (^(fi18)) {
                                    case -772:
                                        var v26 float64 = (((-2.25 - (3.14159 * float64(3.14159))) * (1e10 / 0.1)) * (((0.0 * (0.1 + (1e10 + (1e10 - ((3.14159 + ((1e10 - 3.14159) + 3.14159)) - 0.0))))) + 1.5) + -2.25))
                                        _ = v26
                                    default:
                                        switch fi18 {
                                        case 4294967295:
                                            switch "foo bar" {
                                            case "hello":
                                                v27 := GenS6{F0: -762, F1: 2, F2: -1}
                                                _ = v27
                                            case "foo bar":
                                                switch int(v25) {
                                                case 32768:
                                                    fmt.Println(v25, fi8)
                                                case 0:
                                                    v28 := GenS5{F0: 255}
                                                    _ = v28
                                                default:
                                                    if (false || true) {
                                                        if ("foo bar" <= "hello") {
                                                            var v29 string = ("\n\t" + "foo bar")
                                                            if (95 >= -1) {
                                                                var v30 uint64 = v25
                                                                v29 = v29
                                                                _ = v30
                                                            } else {
                                                                if (!(127 != -128)) {
                                                                    _ = len(v21)
                                                                } else {
                                                                    switch fi8 {
                                                                    case -2147483648:
                                                                        switch 127 {
                                                                        case 1:
                                                                            var v31 uint64 = (v25 - (v25 ^ (v25 - v25)))
                                                                            _ = v31
                                                                        case 9223372036854775807:
                                                                            v25 += v25
                                                                        case -2147483648:
                                                                            var v32 float64 = 0.1
                                                                            _ = v32
                                                                        }
                                                                    case 367:
                                                                        var v33 int = -86
                                                                        _ = v33
                                                                    case 940:
                                                                        v25 += v25
                                                                    default:
                                                                        v25 -= (^(v25))
                                                                    }
                                                                    switch (fi18 % (((-(fi8))) | 1)) {
                                                                    case -294:
                                                                        if len(v21) >= 2 {
                                                                            v21 = v21[1:len(v21)]
                                                                        }
                                                                        fmt.Println(fi8)
                                                                        v34 := GenS7{F0: v29, F1: 0}
                                                                        _ = v34
                                                                    }
                                                                    if true {
                                                                        fmt.Println(v25, fi8, fi18)
                                                                        if len(v21) > 0 {
                                                                            v21[0] = (fi8 << ((2) & 63))
                                                                        }
                                                                    } else {
                                                                        if (v25 == 398) {
                                                                            if true {
                                                                                fmt.Println(fi8, fi8, v29)
                                                                                v25 += (v25 ^ v25)
                                                                                var v35 int64 = 0
                                                                                _ = v35
                                                                            }
                                                                            if (94 < 127) {
                                                                                var v36 bool = (1e10 < (1e10 / (0.1 * float32(3.14159))))
                                                                                if len(v21) > 0 {
                                                                                    v21[0] = fi8
                                                                                }
                                                                                var v37 float64 = 0.1
                                                                                _ = v36
                                                                                _ = v37
                                                                            } else {
                                                                                var v38 uint = 65535
                                                                                var v39 int = 172
                                                                                if len(v21) >= 2 {
                                                                                    v21 = v21[1:len(v21)]
                                                                                }
                                                                                _ = v38
                                                                                _ = v39
                                                                            }
                                                                            var v40 uint = 42
                                                                            _ = v40
                                                                        } else {
                                                                            if true {
                                                                                switch v29 {
                                                                                case "foo bar":
                                                                                    var v41 int = (fi18 / ((-809) | 1))
                                                                                    _ = v41
                                                                                case "\n\t":
                                                                                    var v42 int32 = -2147483648
                                                                                    _ = v42
                                                                                default:
                                                                                    v29 = v29
                                                                                }
                                                                                switch 998 {
                                                                                case -2147483648:
                                                                                    fmt.Println(v29, v29, v29)
                                                                                case 4294967295:
                                                                                    if len(v21) > 0 {
                                                                                        v21[0] = (fi8 | fi18)
                                                                                    }
                                                                                    v29 = (v29 + (((v29 + ("a" + "hello")) + ((v29 + v29) + "hello")) + (v29 + v29)))
                                                                                    fmt.Println(fi8, v29, v29)
                                                                                }
                                                                            } else {
                                                                                v43 := GenS8{F0: (fi18 * (fi8 << ((65535) & 63))), F1: -233}
                                                                                _ = v43
                                                                            }
                                                                            v44 := GenS9{F0: 42, F1: 280, F2: 0}
                                                                            _ = v44
                                                                        }
                                                                        v29 = (v29 + ("a" + (("foo bar" + ("foo bar" + ("foo bar" + v29))) + (v29 + v29))))
                                                                        v45 := GenS8{F0: -972, F1: (fi18 ^ (fi8 | 4294967295))}
                                                                        _ = v45
                                                                    }
                                                                }
                                                                var v46 uint8 = 0
                                                                _ = v46
                                                            }
                                                            v25 = v25
                                                            _ = v29
                                                        }
                                                        if false {
                                                            switch ("" + ((("\n\t" + (("foo bar" + "a") + ("hello" + "a"))) + "foo bar") + "\n\t")) {
                                                            case "foo bar":
                                                                if len(v21) >= 2 {
                                                                    v21 = v21[1:len(v21)]
                                                                }
                                                            default:
                                                                var v47 int8 = 42
                                                                _ = v47
                                                            }
                                                            v25 = 717
                                                            v25 -= 1
                                                            if (false || (!false)) {
                                                                var v48 int32 = 0
                                                                var v49 int8 = 19
                                                                _ = v48
                                                                _ = v49
                                                            }
                                                            if len(v21) >= 2 {
                                                                v21 = v21[1:len(v21)]
                                                            }
                                                            if false {
                                                                _ = len(v21)
                                                            } else {
                                                                v50 := GenS10{F0: 255, F1: 42, F2: 9223372036854775807}
                                                                v25 += (^(v25))
                                                                _ = v50
                                                            }
                                                        } else {
                                                            var v51 int8 = 64
                                                            _ = v51
                                                        }
                                                        v25 -= v25
                                                    } else {
                                                        var v52 uint = 255
                                                        _ = v52
                                                    }
                                                }
                                            case "a":
                                                switch ((("foo bar" + "foo bar") + "a") + ("foo bar" + ("a" + "hello"))) {
                                                case "hello":
                                                    var v53 int = fi18
                                                    _ = v53
                                                case "":
                                                    var v54 int16 = -1
                                                    _ = v54
                                                case "\n\t":
                                                    v25 = (v25 >> ((283) & 63))
                                                }
                                            default:
                                                if true {
                                                    v55 := GenS11{F0: 42, F1: ("\n\t" + "\n\t"), F2: 255}
                                                    v25 -= v25
                                                    _ = v55
                                                } else {
                                                    v56 := GenS2{F0: 255}
                                                    var v57 int = (fi18 & (fi18 << ((4294967295) & 63)))
                                                    _ = v56
                                                    _ = v57
                                                }
                                            }
                                        case 959:
                                            v58 := GenS12{F0: 1, F1: 1, F2: ("hello" + "")}
                                            _ = v58
                                        case 2147483647:
                                            v25 = (v25 | 929)
                                        }
                                        if (650 >= 32767) {
                                            var v59 bool = true
                                            _ = v59
                                        }
                                    }
                                    _ = v21
                                    _ = v24
                                    _ = v25
                                }
                                var v60 int8 = 100
                                v60 = v60
                                if (((!false) && false) && false) {
                                    v61 := []int{fi8, fi18, 32768, (fi18 | (fi8 * (fi8 | fi8)))}
                                    v60 = v60
                                    v60 += v60
                                    _ = v61
                                }
                                var v62 int64 = 9223372036854775807
                                if ((v62 % ((623) | 1)) == (v62 >> ((42) & 63))) {
                                    var v63 int64 = (v62 >> ((207) & 63))
                                    _ = v63
                                }
                                _ = v60
                                _ = v62
                            }
                        case "\n\t":
                            v64 := []int{fi18, -373, fi8}
                            _ = v64
                        }
                        var v65 int = fi18
                        fmt.Println(fi8, fi18)
                        v66 := []int{(^(v65)), fi18, fi18}
                        var v67 int8 = -128
                        _ = v65
                        _ = v66
                        _ = v67
                    } else {
                        v68 := []int{(fi8 ^ 1), (fi18 % ((-214) | 1)), fi8}
                        _ = v68
                    }
                    _ = fi18
                }
            }
            var v69 int = (-(fi8))
            _ = fi8
            _ = v69
        }
    }
    for fi70 := 0; fi70 < 8; fi70++ {
        if (212 <= 697) {
            switch "\n\t" {
            case "foo bar":
                var v71 int32 = 549
                _ = v71
            default:
                v72 := []int{(-(fi70)), (fi70 + (fi70 << ((18446744073709551615) & 63))), -912, (fi70 ^ fi70)}
                _ = v72
            }
            if ((0.1 != float64(0.1)) || false) {
                var v73 float64 = -2.25
                for fi74 := 0; fi74 < 2; fi74++ {
                    var v75 string = ""
                    _ = fi74
                    _ = v75
                }
                v76 := []int{fi70, (fi70 >> ((0) & 63)), (fi70 << ((150) & 63))}
                _ = v73
                _ = v76
            } else {
                fmt.Println(fi70)
                v77 := GenS1{F0: 0, F1: 42}
                _ = v77
            }
            v78 := GenS6{F0: -32768, F1: 564, F2: 42}
            _ = v78
        }
        var v79 uint32 = 42
        _ = fi70
        _ = v79
    }
    if true {
        if true {
            var v80 uint16 = 255
            _ = v80
        } else {
            if (!true) {
                v81 := GenS11{F0: 820, F1: "\n\t", F2: 198}
                var v82 uint8 = 1
                _ = v81
                _ = v82
            } else {
                if false {
                    if (626 == 4294967295) {
                        if true {
                            var v83 uint = 255
                            var v84 int = -15
                            if true {
                                var v85 uint8 = 2
                                v86 := []int{2147483647, (v84 << ((4294967295) & 63)), (-(v84)), -2147483648}
                                v87 := GenS12{F0: -1, F1: 42, F2: "foo bar"}
                                _ = v85
                                _ = v86
                                _ = v87
                            }
                            _ = v83
                            _ = v84
                        } else {
                            var v88 uint8 = 1
                            fmt.Println(v88)
                            fmt.Println(v88)
                            _ = v88
                        }
                    } else {
                        for fi89 := 0; fi89 < 11; fi89++ {
                            v90 := GenS6{F0: -139, F1: 2, F2: 42}
                            v91 := GenS4{F0: 993, F1: -9223372036854775808, F2: 127}
                            _ = fi89
                            _ = v90
                            _ = v91
                        }
                    }
                    v92 := GenS13{F0: 42}
                    _ = v92
                } else {
                    var v93 float64 = (0.0 - -2.25)
                    _ = v93
                }
                v94 := []int{-559, 373}
                if len(v94) >= 2 {
                    v94 = v94[1:len(v94)]
                }
                _ = v94
            }
        }
    }
    return 42
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 898
    for fi95 := 0; fi95 < 1; fi95++ {
        var v96 int32 = 1
        var v97 int16 = -336
        v97 = int16(v96)
        _ = fi95
        _ = v96
        _ = v97
    }
    for fi98 := 0; fi98 < 11; fi98++ {
        fmt.Println(fi98, v3)
        v3 = (v3 % (((v3 ^ (v3 ^ v3))) | 1))
        _ = fi98
    }
    v3 -= v3
    v99 := []int{32767}
    v100 := GenS4{F0: (v3 / ((v3) | 1)), F1: -1, F2: 128}
    var v101 uint32 = 4294967295
    if false {
        if len(v99) > 0 {
            v99[0] = -1
        }
        var v102 float32 = (1.5 - 1.5)
        v103 := GenS14{F0: 858, F1: -687, F2: 0}
        _ = v102
        _ = v103
    }
    switch 706 {
    case 0:
        if (!false) {
            v3 = v101
        } else {
            for fi104 := 0; fi104 < 18; fi104++ {
                var v105 uint8 = 221
                _ = fi104
                _ = v105
            }
            fmt.Println(v3, v3)
            if false {
                switch 273 {
                case 0:
                    fmt.Println(v101)
                }
            } else {
                if (((1e10 + 3.14159) <= 0.0) && true) {
                    v101 -= v3
                    if (((0 > -32768) && (310 != -486)) || true) {
                        for fi106 := 0; fi106 < 13; fi106++ {
                            v2[(v3 | v101)] = (v101 << ((42) & 63))
                            var v107 uint8 = 1
                            _ = fi106
                            _ = v107
                        }
                        switch "foo bar" {
                        case "a":
                            v101 += v101
                        default:
                            if false {
                                var v108 uint = 929
                                for fi109 := 0; fi109 < 11; fi109++ {
                                    for fi110 := 0; fi110 < 7; fi110++ {
                                        switch genFn0() {
                                        case 32768:
                                            if true {
                                                if false {
                                                    var v111 uint64 = 0
                                                    if true {
                                                        v2[v3] = v3
                                                        fmt.Println(v101)
                                                    } else {
                                                        var v112 uint32 = 173
                                                        v113 := GenS15{F0: 248}
                                                        _ = v112
                                                        _ = v113
                                                    }
                                                    _ = v111
                                                } else {
                                                    v114 := GenS16{F0: 2, F1: 0}
                                                    switch "a" {
                                                    case "a":
                                                        if (0 <= -2147483648) {
                                                            var v115 uint8 = 128
                                                            var v116 int16 = 42
                                                            _ = len(v2)
                                                            _ = v115
                                                            _ = v116
                                                        } else {
                                                            var v117 int = fi109
                                                            _ = v117
                                                        }
                                                    default:
                                                        switch fi109 {
                                                        case -2147483648:
                                                            fmt.Println(fi110)
                                                        case -156:
                                                            var v118 uint64 = 0
                                                            _ = v118
                                                        case -773:
                                                            var v119 int = genFn0()
                                                            v120 := GenS11{F0: 933, F1: "foo bar", F2: 7}
                                                            v119 = -2147483648
                                                            _ = v119
                                                            _ = v120
                                                        }
                                                    }
                                                    fmt.Println(fi110, fi109)
                                                    v121 := GenS17{F0: (-1 > 0), F1: (fi109 == fi110), F2: 255}
                                                    var v122 uint32 = (v3 / ((2147483648) | 1))
                                                    _ = v114
                                                    _ = v121
                                                    _ = v122
                                                }
                                                v123 := GenS18{F0: genFn0()}
                                                _ = v123
                                            }
                                        default:
                                            if len(v99) >= 2 {
                                                v99 = v99[1:len(v99)]
                                            }
                                        }
                                        v3 += v3
                                        _ = fi110
                                    }
                                    _ = fi109
                                }
                                _ = v108
                            } else {
                                if len(v99) >= 2 {
                                    v99 = v99[1:len(v99)]
                                }
                                if (42 < 0) {
                                    v99 = append(v99, 812)
                                    var v124 int16 = 1
                                    _ = v124
                                } else {
                                    if len(v99) >= 2 {
                                        v99 = v99[1:len(v99)]
                                    }
                                }
                                fmt.Println(v3, v101)
                            }
                        }
                    } else {
                        var v125 uint8 = 42
                        var v126 uint32 = 42
                        var v127 uint = 470
                        _ = v125
                        _ = v126
                        _ = v127
                    }
                    fmt.Println(v3, v101)
                } else {
                    fmt.Println(v3, v101)
                    var v128 bool = true
                    v3 = v101
                    _ = v128
                }
                if ((0.0 * 0.1) <= 0.0) {
                    switch 127 {
                    case -2147483648:
                        v3 = 1
                    case 32767:
                        if true {
                            var v129 uint8 = 97
                            _ = v129
                        } else {
                            var v130 float64 = (0.1 * 1.5)
                            for fi131 := 0; fi131 < 6; fi131++ {
                                if (-666 >= -815) {
                                    var v132 float64 = ((v130 + float64(1e10)) / v130)
                                    v133 := GenS17{F0: (!(false || ((0.0 * v132) == 3.14159))), F1: true, F2: 842}
                                    _ = v132
                                    _ = v133
                                } else {
                                    var v134 uint64 = 762
                                    _ = v134
                                }
                                _ = fi131
                            }
                            _ = v130
                        }
                    default:
                        switch "a" {
                        case "a":
                            _ = len(v99)
                        }
                    }
                    switch "hello" {
                    case "a":
                        fmt.Println(v101)
                    default:
                        switch "foo bar" {
                        case "hello":
                            var v135 uint = 65535
                            _ = v135
                        case "foo bar":
                            if len(v99) > 0 {
                                v99[0] = genFn0()
                            }
                        }
                    }
                } else {
                    v101 -= 2147483648
                    switch "" {
                    case "\n\t":
                        v3 = (v3 << ((0) & 63))
                    case "a":
                        switch 9223372036854775807 {
                        case 42:
                            if false {
                                v101 = v101
                                _ = len(v2)
                                v101 = v3
                            }
                        default:
                            var v136 bool = ("" <= ("hello" + "\n\t"))
                            _ = v136
                        }
                    default:
                        for fi137 := 0; fi137 < 2; fi137++ {
                            v2[v101] = v3
                            if (42 != 178) {
                                var v138 uint8 = 184
                                _ = v138
                            } else {
                                if ((((("a" + "foo bar") + "") + "\n\t") != "\n\t") && false) {
                                    v3 += v3
                                    v3 += v101
                                    v99 = append(v99, fi137)
                                } else {
                                    var v139 string = "foo bar"
                                    _ = v139
                                }
                                if ((0.0 + 3.14159) != -2.25) {
                                    var v140 int8 = 55
                                    _ = v140
                                }
                            }
                            v101 -= v101
                            _ = fi137
                        }
                    }
                }
            }
        }
    case 846:
        if false {
            if false {
                var v141 int64 = 42
                _ = v141
            } else {
                switch ("foo bar" + "foo bar") {
                case "a":
                    var v142 int8 = 0
                    _ = v142
                }
            }
            for fi143 := 0; fi143 < 8; fi143++ {
                if ((1000 > 1) && true) {
                    v3 = v3
                    v3 += 414
                    v3 += (v3 & v3)
                } else {
                    fmt.Println(v3)
                    _ = len(v99)
                    v3 = (v101 | (v3 + 2))
                }
                var v144 int = genFn0()
                var v145 uint64 = 42
                _ = fi143
                _ = v144
                _ = v145
            }
        }
    default:
        var v146 string = ""
        _ = v146
    }
    for fi147 := 0; fi147 < 17; fi147++ {
        if (("\n\t" + "hello") == ((("hello" + "\n\t") + "") + ("\n\t" + ""))) {
            var v148 float32 = ((0.0 * ((0.0 + 3.14159) * 0.1)) - 0.1)
            v149 := GenS19{F0: fi147, F1: false}
            var v150 uint32 = (v101 * (^(v3)))
            _ = v148
            _ = v149
            _ = v150
        }
        fmt.Println(v101)
        v151 := GenS15{F0: 9223372036854775807}
        _ = fi147
        _ = v151
    }
    for fi152 := 0; fi152 < 19; fi152++ {
        var v153 float32 = ((3.14159 + -2.25) / 3.14159)
        var v154 int = (fi152 & fi152)
        _ = fi152
        _ = v153
        _ = v154
    }
    if (false && false) {
        for fi155 := 0; fi155 < 13; fi155++ {
            v101 = 2147483647
            for fi156 := 0; fi156 < 1; fi156++ {
                var v157 int = genFn0()
                _ = fi156
                _ = v157
            }
            switch "hello" {
            case "foo bar":
                if false {
                    switch (fi155 & 1) {
                    case 0:
                        var v158 uint = 4294967295
                        _ = v158
                    case 128:
                        var v159 uint16 = 960
                        _ = v159
                    case -9223372036854775808:
                        switch 9223372036854775807 {
                        case -9223372036854775808:
                            for fi160 := 0; fi160 < 10; fi160++ {
                                switch (fi155 % ((fi155) | 1)) {
                                case 128:
                                    var v161 uint = 4294967295
                                    _ = v161
                                case 42:
                                    var v162 uint32 = (v101 | v101)
                                    _ = v162
                                default:
                                    fmt.Println(v101, fi155, fi160)
                                }
                                var v163 int32 = 2147483647
                                fmt.Println(v101, v3, v101)
                                _ = fi160
                                _ = v163
                            }
                        case 2147483647:
                            v101 += v3
                        case 4294967295:
                            v101 = 0
                        }
                    }
                    switch "hello" {
                    case "\n\t":
                        var v164 uint64 = 1
                        _ = v164
                    case "hello":
                        if false {
                            var v165 bool = true
                            delete(v2, v3)
                            _ = v165
                        }
                    case "foo bar":
                        var v166 uint = 0
                        _ = v166
                    default:
                        switch genFn0() {
                        case 981:
                            var v167 int = genFn0()
                            _ = v167
                        default:
                            v3 -= 589
                        }
                    }
                    v168 := GenS20{F0: (4294967295 < 495), F1: -2147483648}
                    _ = v168
                }
            case "\n\t":
                if (v3 != v3) {
                    v101 -= v3
                    if len(v99) > 0 {
                        v99[0] = (fi155 + (fi155 / ((-627) | 1)))
                    }
                    switch fi155 {
                    case 1:
                        for fi169 := 0; fi169 < 16; fi169++ {
                            v170 := GenS21{F0: 42}
                            _ = fi169
                            _ = v170
                        }
                    case -748:
                        if (9223372036854775807 <= -1) {
                            v171 := GenS14{F0: 1, F1: 0, F2: 2147483648}
                            _ = v171
                        }
                    case 772:
                        v3 += 782
                    }
                } else {
                    v99 = append(v99, fi155)
                }
            case "":
                var v172 int = fi155
                _ = v172
            }
            _ = fi155
        }
        if true {
            if false {
                for fi173 := 0; fi173 < 20; fi173++ {
                    for fi174 := 0; fi174 < 11; fi174++ {
                        var v175 int8 = 42
                        _ = fi174
                        _ = v175
                    }
                    for fi176 := 0; fi176 < 20; fi176++ {
                        switch fi173 {
                        case 9223372036854775807:
                            var v177 float32 = 1e10
                            _ = v177
                        default:
                            if true {
                                var v178 int8 = -128
                                fmt.Println(fi176, fi173)
                                _ = v178
                            }
                        }
                        v3 = 765
                        if false {
                            v179 := GenS16{F0: 2, F1: -9}
                            fmt.Println(fi176, v3, v3)
                            var v180 uint16 = 255
                            v181 := GenS20{F0: false, F1: 2147483647}
                            _ = v179
                            _ = v180
                            _ = v181
                        }
                        v101 = v3
                        _ = fi176
                    }
                    v182 := GenS12{F0: 77, F1: -1, F2: "\n\t"}
                    _ = fi173
                    _ = v182
                }
            }
            v3 += 2147483648
        }
        switch ("hello" + "a") {
        case "\n\t":
            var v183 uint = 1
            _ = v183
        case "foo bar":
            for fi184 := 0; fi184 < 4; fi184++ {
                switch "\n\t" {
                case "a":
                    v1[(fi184 & genFn0())] = ("\n\t" + ("\n\t" + "\n\t"))
                case "":
                    for fi185 := 0; fi185 < 15; fi185++ {
                        var v186 uint64 = 67
                        _ = fi185
                        _ = v186
                    }
                default:
                    var v187 int16 = 42
                    _ = v187
                }
                _ = fi184
            }
        case "a":
            for fi188 := 0; fi188 < 5; fi188++ {
                v101 -= v101
                for fi189 := 0; fi189 < 2; fi189++ {
                    var v190 uint32 = 2147483647
                    if false {
                        v3 = v101
                        v191 := GenS22{F0: -2147483648}
                        _ = v191
                    }
                    _ = fi189
                    _ = v190
                }
                _ = fi188
            }
        }
    } else {
        v192 := GenS17{F0: false, F1: false, F2: 255}
        fmt.Println(v101, v3)
        v3 += 2147483648
        _ = v192
    }
    if (-32768 < 903) {
        if (!((1 >= 559) && (4294967295 >= -76))) {
            switch 128 {
            case 9223372036854775807:
                if true {
                    v193 := GenS21{F0: 24}
                    for fi194 := 0; fi194 < 13; fi194++ {
                        var v195 float64 = -2.25
                        _ = fi194
                        _ = v195
                    }
                    fmt.Println(v101, v101)
                    _ = v193
                } else {
                    var v196 string = ("hello" + "")
                    if (4294967295 == 421) {
                        for fi197 := 0; fi197 < 18; fi197++ {
                            delete(v2, 4294967295)
                            var v198 int64 = 0
                            _ = fi197
                            _ = v198
                        }
                        var v199 uint16 = 1
                        for fi200 := 0; fi200 < 20; fi200++ {
                            fmt.Println(fi200, fi200, v3)
                            var v201 uint8 = 0
                            _ = fi200
                            _ = v201
                        }
                        _ = v199
                    }
                    _ = v196
                }
            case -734:
                v3 += v3
            case 264:
                v3 -= 2147483647
            default:
                v202 := v2[v3]
                _ = v202
            }
            var v203 bool = false
            if len(v99) > 0 {
                v99[0] = 845
            }
            _ = v203
        } else {
            v204 := GenS23{F0: 4294967295}
            delete(v2, (v3 >> ((1) & 63)))
            _ = v204
        }
        var v205 int = -472
        var v206 uint64 = 18446744073709551615
        _ = v205
        _ = v206
    } else {
        var v207 int32 = -959
        if false {
            if len(v99) > 0 {
                v99[0] = -1
            }
            switch -312 {
            case 886:
                if len(v99) > 0 {
                    v99[0] = genFn0()
                }
            case 2147483647:
                v101 = v3
            }
            _ = len(v99)
        } else {
            v208 := GenS24{F0: 646, F1: (!true)}
            v209 := GenS18{F0: genFn0()}
            if (421 >= 378) {
                v207 += 1
                var v210 int32 = -860
                _ = v210
            }
            _ = v208
            _ = v209
        }
        _ = v207
    }
    var v211 int32 = 106
    for fi212 := 0; fi212 < 8; fi212++ {
        var v213 float64 = 0.1
        _ = fi212
        _ = v213
    }
    fmt.Println(v101, v101, v211)
    v211 = v211
    fmt.Println(v3, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v101:", v101)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v211:", v211)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v99:", len(v99), v99)
    _ = v1
    _ = v100
    _ = v101
    _ = v2
    _ = v211
    _ = v3
    _ = v99
}
