package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 bool; F2 int16 }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 uint8; F1 string; F2 uint32 }
type GenS6 struct { F0 int; F1 bool; F2 int64 }
type GenS7 struct { F0 int32; F1 string }
type GenS8 struct { F0 int16; F1 int32 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int64; F1 uint32 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 string; F1 uint16; F2 string }
type GenS14 struct { F0 uint32; F1 int32 }
type GenS15 struct { F0 int16; F1 uint8 }
type GenS16 struct { F0 string; F1 int8 }
type GenS17 struct { F0 int8; F1 int8; F2 string }
type GenS18 struct { F0 uint8 }
type GenS19 struct { F0 uint32; F1 int8 }
type GenS20 struct { F0 int; F1 uint64; F2 int64 }
type GenS21 struct { F0 uint16; F1 uint32 }
type GenS22 struct { F0 uint8; F1 int8 }
type GenS23 struct { F0 int32 }
type GenS24 struct { F0 bool }
func genFn0(p0 int64, p1 uint32) uint8 {
    var v7 int64 = (p0 >> ((1) & 63))
    v8 := GenS1{F0: 1, F1: 0}
    var v9 float32 = 1e10
    for fi10 := 0; fi10 < 13; fi10++ {
        p1 -= p1
        v11 := []int{32768, (fi10 * (^(fi10))), (fi10 + (fi10 + 127))}
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 255
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = -2147483648
    var v3 int64 = 0
    var v4 float32 = ((((0.1 - -2.25) - 0.0) + 0.1) + 3.14159)
    var v5 bool = (!false)
    v6 := []string{(("a" + "hello") + "a")}
    switch "a" {
    case "a":
        switch "hello" {
        case "\n\t":
            if true {
                switch "a" {
                case "":
                    if v5 {
                        v2 += (v2 ^ 2147483647)
                    }
                case "foo bar":
                    v12 := GenS1{F0: 668, F1: 1000}
                    _ = v12
                }
            } else {
                for fi13 := 0; fi13 < 17; fi13++ {
                    switch fi13 {
                    case -1:
                        for fi14 := 0; fi14 < 19; fi14++ {
                            switch fi13 {
                            case 32767:
                                if len(v6) > 0 {
                                    v6[0] = (("\n\t" + "a") + (("" + "a") + ""))
                                }
                            case -2147483648:
                                switch 127 {
                                case -593:
                                    v4 -= (v4 - v4)
                                case 4294967295:
                                    delete(v1, 4294967295)
                                case 2147483647:
                                    if (70 < -128) {
                                        var v15 int64 = (-(v3))
                                        fmt.Println(v3, fi14)
                                        var v16 int8 = -87
                                        if (!v5) {
                                            var v17 uint = 0
                                            v18 := GenS2{F0: 2147483647}
                                            switch "a" {
                                            case "hello":
                                                var v19 int8 = (^(v16))
                                                var v20 bool = true
                                                v21 := GenS3{F0: fi13, F1: false, F2: -1}
                                                _ = v19
                                                _ = v20
                                                _ = v21
                                            case "":
                                                v22 := GenS4{F0: 42, F1: (fi14 - int(v15))}
                                                _ = v22
                                            case "a":
                                                v23 := GenS2{F0: 1}
                                                _ = v23
                                            default:
                                                switch (fi14 * (^(fi13))) {
                                                case 32767:
                                                    switch "hello" {
                                                    case "\n\t":
                                                        v24 := v1[4294967295]
                                                        _ = v24
                                                    case "a":
                                                        v6 = append(v6, ("a" + "a"))
                                                    }
                                                case 42:
                                                    var v25 int = -63
                                                    _ = v25
                                                case 790:
                                                    v26 := GenS5{F0: 1, F1: ("\n\t" + ("foo bar" + "a")), F2: 2147483647}
                                                    _ = v26
                                                default:
                                                    if (!(1 <= uint64(v17))) {
                                                        v27 := GenS3{F0: 9223372036854775807, F1: false, F2: 42}
                                                        _ = v27
                                                    }
                                                }
                                            }
                                            _ = v17
                                            _ = v18
                                        } else {
                                            v28 := GenS1{F0: 2, F1: -487}
                                            v29 := GenS6{F0: fi14, F1: (fi14 < -2147483648), F2: v15}
                                            _ = v28
                                            _ = v29
                                        }
                                        _ = v15
                                        _ = v16
                                    } else {
                                        var v30 uint8 = 102
                                        var v31 bool = (4294967295 < 1)
                                        v2 = (^(v2))
                                        v32 := v1[946]
                                        fmt.Println(fi14)
                                        _ = v30
                                        _ = v31
                                        _ = v32
                                    }
                                }
                                v1[187] = -2147483648
                            case 32768:
                                v4 -= v4
                            }
                            if (!((fi14 << ((0) & 63)) <= fi13)) {
                                var v33 uint = 255
                                v4 -= (-2.25 + v4)
                                v2 = v2
                                _ = v33
                            } else {
                                _ = len(v1)
                                var v34 int32 = (v2 | (v2 >> ((1) & 63)))
                                switch 242 {
                                case 99:
                                    v35 := GenS2{F0: 0}
                                    _ = len(v6)
                                    _ = v35
                                default:
                                    v3 = v3
                                }
                                if (int8(fi14) < 42) {
                                    var v36 float32 = 1e10
                                    var v37 float64 = 3.14159
                                    _ = v36
                                    _ = v37
                                } else {
                                    var v38 uint64 = 2
                                    _ = v38
                                }
                                fmt.Println(v3, v2, v3)
                                var v39 uint16 = 42
                                v40 := GenS7{F0: 834, F1: "hello"}
                                _ = v34
                                _ = v39
                                _ = v40
                            }
                            var v41 string = "foo bar"
                            delete(v1, 2147483647)
                            _ = fi14
                            _ = v41
                        }
                    case -640:
                        var v42 int8 = 0
                        _ = v42
                    case 2147483647:
                        var v43 uint64 = 1
                        _ = v43
                    }
                    _ = fi13
                }
                for fi44 := 0; fi44 < 8; fi44++ {
                    fmt.Println(v4)
                    v6 = append(v6, ("\n\t" + "a"))
                    var v45 uint32 = 2147483648
                    _ = fi44
                    _ = v45
                }
                v46 := GenS8{F0: -255, F1: v2}
                _ = v46
            }
        case "hello":
            var v47 uint32 = 2
            _ = v47
        case "foo bar":
            v48 := GenS4{F0: -1, F1: 1}
            _ = v48
        default:
            v49 := GenS1{F0: 288, F1: 0}
            _ = v49
        }
    case "foo bar":
        if len(v6) > 0 {
            v6[0] = "\n\t"
        }
    default:
        v4 = (((v4 - v4) + v4) / float32((1.5 - 0.0)))
    }
    v50 := GenS9{F0: "\n\t"}
    fmt.Println(v2, v3, v5)
    for fi51 := 0; fi51 < 3; fi51++ {
        v4 = -2.25
        _ = fi51
    }
    if v5 {
        var v52 uint8 = 255
        v53 := GenS8{F0: 42, F1: (v2 / ((v2) | 1))}
        if (1e10 == -2.25) {
            v4 -= v4
            var v54 bool = true
            _ = v54
        } else {
            v6 = append(v6, "a")
            for fi55 := 0; fi55 < 6; fi55++ {
                for fi56 := 0; fi56 < 9; fi56++ {
                    if (v5 && v5) {
                        var v57 int64 = 9223372036854775807
                        _ = v57
                    } else {
                        v58 := GenS10{F0: (v3 & -9223372036854775808), F1: 2147483648}
                        v2 = v2
                        _ = v58
                    }
                    v59 := GenS7{F0: v2, F1: "a"}
                    _ = fi56
                    _ = v59
                }
                v60 := GenS3{F0: fi55, F1: v5, F2: -32768}
                _ = fi55
                _ = v60
            }
            for fi61 := 0; fi61 < 11; fi61++ {
                var v62 int32 = (v2 * (v2 & 0))
                if (-391 > 42) {
                    var v63 int32 = 1
                    switch 135 {
                    case 2147483647:
                        v62 = (-(v62))
                    default:
                        if v5 {
                            var v64 uint16 = 1
                            if (v3 <= (v3 & (v3 - (v3 + (v3 % (((v3 / (((v3 << ((42) & 63))) | 1))) | 1)))))) {
                                fmt.Println(v62, v4, v63)
                            } else {
                                fmt.Println(fi61)
                            }
                            _ = v64
                        }
                    }
                    _ = v63
                } else {
                    v5 = ((v52 >> ((255) & 63)) <= genFn0(v3, 0))
                    switch (fi61 / ((128) | 1)) {
                    case 32767:
                        v4 = v4
                    case 0:
                        if (!(682 == 18446744073709551615)) {
                            if v5 {
                                v65 := GenS9{F0: ("foo bar" + "a")}
                                for fi66 := 0; fi66 < 9; fi66++ {
                                    var v67 uint32 = 2147483648
                                    fmt.Println(v4)
                                    v2 += (^(v2))
                                    if ((((v4 * (((v4 * (v4 / v4)) + (-2.25 - (v4 * v4))) - (0.1 * 1.5))) * (v4 + v4)) * v4) <= v4) {
                                        v52 = (v52 % (((v52 - 2)) | 1))
                                    } else {
                                        v2 = (v62 + -1)
                                        if (2 == 277) {
                                            v52 = v52
                                            var v68 uint = 1
                                            _ = v68
                                        } else {
                                            var v69 int32 = v62
                                            v70 := GenS11{F0: 332}
                                            v52 = (^(v52))
                                            _ = v69
                                            _ = v70
                                        }
                                    }
                                    fmt.Println(v5, v2)
                                    fmt.Println(fi61, v62, v67)
                                    v2 = (v62 * (v2 ^ (v62 * (v62 >> ((1) & 63)))))
                                    v71 := GenS4{F0: 32767, F1: (-(fi66))}
                                    _ = len(v6)
                                    if (v5 || v5) {
                                        if v5 {
                                            _ = len(v6)
                                            fmt.Println(v4, fi61)
                                            switch (fi61 * (fi66 | -2147483648)) {
                                            case -2147483648:
                                                v62 += v2
                                            case 128:
                                                if (!v5) {
                                                    v52 += v52
                                                } else {
                                                    v67 += v67
                                                    var v72 int = fi61
                                                    switch (fi66 | -1) {
                                                    case 4294967295:
                                                        var v73 uint64 = 42
                                                        _ = v73
                                                    case 38:
                                                        var v74 uint32 = v67
                                                        _ = v74
                                                    }
                                                    _ = v72
                                                }
                                            default:
                                                fmt.Println(v2)
                                            }
                                        } else {
                                            v75 := GenS12{F0: 18446744073709551615}
                                            v76 := GenS6{F0: int(v3), F1: v5, F2: -9223372036854775808}
                                            _ = v75
                                            _ = v76
                                        }
                                        fmt.Println(v4)
                                        v77 := GenS13{F0: "a", F1: 42, F2: "foo bar"}
                                        var v78 uint64 = 2
                                        _ = v77
                                        _ = v78
                                    } else {
                                        if ((genFn0((v3 - v3), (^(v67))) == v52) || (0 != v62)) {
                                            v79 := v1[1]
                                            var v80 bool = v5
                                            var v81 int = -1
                                            v3 -= (-(v3))
                                            var v82 int32 = 0
                                            _ = v79
                                            _ = v80
                                            _ = v81
                                            _ = v82
                                        } else {
                                            var v83 uint8 = 153
                                            switch "\n\t" {
                                            case "a":
                                                fmt.Println(v67, v5, v3)
                                            case "\n\t":
                                                if (true && false) {
                                                    delete(v1, (v67 - v67))
                                                    switch ("foo bar" + "") {
                                                    case "foo bar":
                                                        v84 := GenS2{F0: (^(v67))}
                                                        _ = v84
                                                    case "hello":
                                                        v85 := GenS14{F0: 2147483647, F1: (^(v2))}
                                                        _ = v85
                                                    default:
                                                        v86 := GenS15{F0: -520, F1: v83}
                                                        _ = v86
                                                    }
                                                }
                                            default:
                                                var v87 uint16 = 5
                                                _ = v87
                                            }
                                            var v88 int64 = (^(v3))
                                            v52 = 2
                                            _ = v83
                                            _ = v88
                                        }
                                        switch "" {
                                        case "hello":
                                            v3 = int64(v62)
                                        default:
                                            fmt.Println(v4, v5, v5)
                                            if ("\n\t" >= ("\n\t" + (("a" + ("foo bar" + ("a" + ("\n\t" + "a")))) + ("foo bar" + ("\n\t" + "a"))))) {
                                                v89 := GenS16{F0: "a", F1: -1}
                                                v1[2147483647] = -26
                                                var v90 uint16 = 255
                                                _ = v89
                                                _ = v90
                                            } else {
                                                if v5 {
                                                    if ((1000 == -935) && v5) {
                                                        if v5 {
                                                            fmt.Println(v5)
                                                            _ = len(v6)
                                                            _ = len(v6)
                                                            v91 := GenS8{F0: 1, F1: 2147483647}
                                                            var v92 string = "\n\t"
                                                            _ = v91
                                                            _ = v92
                                                        } else {
                                                            v62 = int32(v52)
                                                        }
                                                        v3 -= v3
                                                    } else {
                                                        v93 := GenS17{F0: 1, F1: -68, F2: ""}
                                                        v4 = ((1.5 / 1.0) * (v4 - ((v4 / 1.0) * 3.14159)))
                                                        _ = v93
                                                    }
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                    switch ("a" + ("" + ("\n\t" + ""))) {
                                                    case "a":
                                                        switch ((((("hello" + "\n\t") + "hello") + "foo bar") + "hello") + "foo bar") {
                                                        case "hello":
                                                            fmt.Println(fi66, v62)
                                                        case "a":
                                                            var v94 uint16 = 1
                                                            _ = v94
                                                        case "":
                                                            var v95 uint32 = 65535
                                                            _ = v95
                                                        }
                                                    case "hello":
                                                        var v96 bool = false
                                                        _ = v96
                                                    }
                                                }
                                                var v97 int64 = v3
                                                _ = len(v1)
                                                _ = v97
                                            }
                                        }
                                        v67 = (v67 * 578)
                                    }
                                    fmt.Println(fi66, v4, v5)
                                    if (v62 < (v2 ^ int32(v3))) {
                                        switch "" {
                                        case "":
                                            v4 += float32(float64(1.5))
                                        case "hello":
                                            var v98 uint64 = 670
                                            _ = v98
                                        case "a":
                                            if v5 {
                                                if true {
                                                    var v99 float32 = float32(1.5)
                                                    fmt.Println(v4, fi61)
                                                    v100 := GenS18{F0: (v52 * 56)}
                                                    _ = v100
                                                    _ = v99
                                                }
                                                if true {
                                                    v5 = true
                                                }
                                                fmt.Println(v5)
                                            } else {
                                                v67 -= v67
                                                var v101 float64 = 3.14159
                                                _ = v101
                                            }
                                        }
                                        switch (-(fi61)) {
                                        case 36:
                                            var v102 string = ("foo bar" + "foo bar")
                                            _ = v102
                                        case 1:
                                            v62 = (v2 >> ((uint(v52)) & 63))
                                        case 4294967295:
                                            if (v5 && true) {
                                                fmt.Println(fi66, fi61)
                                            } else {
                                                var v103 uint16 = 42
                                                v104 := GenS19{F0: v67, F1: 127}
                                                _ = v103
                                                _ = v104
                                            }
                                        default:
                                            v62 += (v62 % ((-1) | 1))
                                        }
                                        switch (fi66 ^ fi66) {
                                        case 9223372036854775807:
                                            switch (("a" + ("foo bar" + "\n\t")) + "") {
                                            case "a":
                                                if (711 == 642) {
                                                    if len(v6) > 0 {
                                                        v6[0] = "\n\t"
                                                    }
                                                }
                                            }
                                        case -677:
                                            v105 := GenS20{F0: fi66, F1: 1, F2: v3}
                                            _ = v105
                                        default:
                                            var v106 int32 = int32(v3)
                                            _ = v106
                                        }
                                    }
                                    var v107 uint16 = 102
                                    _ = fi66
                                    _ = v107
                                    _ = v67
                                    _ = v71
                                }
                                _ = v65
                            } else {
                                var v108 uint64 = 271
                                _ = v108
                            }
                        } else {
                            v4 = (-2.25 + v4)
                            v4 += -2.25
                            v3 += (v3 | -9223372036854775808)
                        }
                    case -756:
                        fmt.Println(v2, v3, v52)
                    }
                    for fi109 := 0; fi109 < 4; fi109++ {
                        v2 = (v62 * (v2 % (((^(v2))) | 1)))
                        _ = fi109
                    }
                }
                _ = fi61
                _ = v62
            }
        }
        _ = v52
        _ = v53
    } else {
        if ((!(9223372036854775807 != -244)) || true) {
            for fi110 := 0; fi110 < 18; fi110++ {
                v111 := GenS6{F0: fi110, F1: v5, F2: v3}
                _ = fi110
                _ = v111
            }
            v5 = (!v5)
        }
        if (53 == -74) {
            fmt.Println(v3, v4)
            fmt.Println(v3)
        }
    }
    v4 += v4
    v6 = append(v6, "foo bar")
    if (-2147483648 > (v2 >> ((4294967295) & 63))) {
        for fi112 := 0; fi112 < 17; fi112++ {
            if false {
                switch "" {
                case "a":
                    switch (fi112 << ((446) & 63)) {
                    case 4294967295:
                        for fi113 := 0; fi113 < 1; fi113++ {
                            var v114 string = "hello"
                            fmt.Println(v4)
                            var v115 int16 = -839
                            switch v114 {
                            case "":
                                v116 := GenS6{F0: fi113, F1: v5, F2: (v3 << ((4294967295) & 63))}
                                _ = v116
                            case "hello":
                                var v117 uint16 = 42
                                _ = v117
                            }
                            switch (fi112 % (((fi113 >> ((4294967295) & 63))) | 1)) {
                            case 127:
                                if v5 {
                                    var v118 uint64 = 1
                                    var v119 float64 = 1.5
                                    var v120 uint16 = 2
                                    _ = v118
                                    _ = v119
                                    _ = v120
                                } else {
                                    v1[148] = fi112
                                    fmt.Println(v3, v114, v115)
                                }
                            case 32768:
                                v121 := GenS13{F0: v114, F1: 32767, F2: ""}
                                _ = v121
                            default:
                                if ((42 < 103) || ((v115 % ((1) | 1)) < (v115 % (((v115 * (v115 % (((v115 / ((32767) | 1))) | 1)))) | 1)))) {
                                    var v122 float32 = float32(-2.25)
                                    _ = v122
                                }
                            }
                            _ = fi113
                            _ = v114
                            _ = v115
                        }
                    }
                case "hello":
                    if v5 {
                        fmt.Println(v3, fi112, v5)
                        v4 = -2.25
                    } else {
                        var v123 int8 = 0
                        fmt.Println(v2)
                        if len(v6) > 0 {
                            v6[0] = ("foo bar" + ("\n\t" + "foo bar"))
                        }
                        _ = v123
                    }
                case "\n\t":
                    fmt.Println(v3)
                default:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                v2 = (-(v2))
                switch "a" {
                case "a":
                    v124 := GenS21{F0: 924, F1: 65535}
                    _ = v124
                }
            } else {
                for fi125 := 0; fi125 < 8; fi125++ {
                    switch fi125 {
                    case 4294967295:
                        var v126 uint16 = 2
                        _ = v126
                    case 292:
                        if (851 <= (v2 / (((v2 % (((v2 % ((v2) | 1))) | 1))) | 1))) {
                            var v127 int = fi125
                            _ = v127
                        } else {
                            if (0.0 != ((1e10 - (0.0 - -2.25)) - 1e10)) {
                                v128 := GenS10{F0: (^(v3)), F1: 540}
                                var v129 int8 = 71
                                v130 := GenS22{F0: 200, F1: (v129 ^ v129)}
                                v131 := GenS16{F0: ("hello" + ("a" + "hello")), F1: 100}
                                var v132 int = fi112
                                var v133 bool = false
                                switch "hello" {
                                case "\n\t":
                                    var v134 int16 = int16(v129)
                                    _ = v134
                                case "foo bar":
                                    v6 = append(v6, "\n\t")
                                case "hello":
                                    var v135 string = "hello"
                                    _ = v135
                                }
                                var v136 int = (v132 - (fi125 | (^(fi125))))
                                _ = v128
                                _ = v129
                                _ = v130
                                _ = v131
                                _ = v132
                                _ = v133
                                _ = v136
                            }
                            var v137 int32 = v2
                            if (v4 > v4) {
                                fmt.Println(v4)
                                var v138 uint64 = 2
                                _ = v138
                            } else {
                                v139 := GenS9{F0: ""}
                                var v140 uint = 517
                                v141 := GenS23{F0: (v137 >> (((^(v140))) & 63))}
                                _ = v139
                                _ = v140
                                _ = v141
                            }
                            _ = v137
                        }
                    case 41:
                        if (4294967295 != 2) {
                            v4 -= 0.1
                        }
                    }
                    var v142 uint32 = 456
                    _ = fi125
                    _ = v142
                }
                v2 = (v2 * (v2 & (^(v2))))
            }
            _ = fi112
        }
    }
    v5 = (42 != 434)
    if true {
        var v143 int16 = -272
        if v5 {
            _ = len(v6)
        } else {
            v144 := GenS24{F0: (280 != 924)}
            v5 = v5
            v145 := v1[790]
            _ = v144
            _ = v145
        }
        v146 := v1[2]
        _ = v143
        _ = v146
    } else {
        v3 = v3
        var v147 int64 = (-(v3))
        _ = v147
    }
    v4 += v4
    if true {
        v5 = v5
        if (!(!true)) {
            v148 := GenS24{F0: true}
            var v149 uint16 = 1
            _ = v148
            _ = v149
        }
        if v5 {
            switch 322 {
            case 521:
                v4 += ((v4 + v4) * float32((0.1 * ((3.14159 * 3.14159) / 1e10))))
            case 32768:
                var v150 string = "\n\t"
                _ = v150
            case 638:
                fmt.Println(v5, v4)
            default:
                v1[2147483648] = 32767
            }
        } else {
            v151 := GenS10{F0: (-(v3)), F1: 0}
            v1[2147483648] = -183
            fmt.Println(v2, v4)
            _ = v151
        }
    } else {
        var v152 bool = v5
        switch 32767 {
        case -9223372036854775808:
            _ = len(v1)
        case 42:
            v3 -= v3
        default:
            v5 = (v2 >= 201)
        }
        _ = v152
    }
    v153 := GenS13{F0: (("\n\t" + "hello") + ""), F1: 42, F2: (("" + ("foo bar" + "hello")) + "a")}
    var v154 string = "hello"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v154:", v154)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v153
    _ = v154
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v6
}
