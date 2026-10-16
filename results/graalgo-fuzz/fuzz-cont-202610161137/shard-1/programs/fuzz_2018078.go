package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 int16; F2 uint }
type GenS4 struct { F0 uint64; F1 int; F2 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32; F1 bool; F2 uint32 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint32; F1 uint16 }
type GenS13 struct { F0 int32; F1 uint8; F2 int16 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 uint64 }
type GenS17 struct { F0 int32 }
type GenS18 struct { F0 int16; F1 uint32; F2 uint64 }
func genFn0(p0 int64, p1 bool) int8 {
    v4 := []int{}
    if p1 {
        for fi5 := 0; fi5 < 17; fi5++ {
            if (916 < 42) {
                for fi6 := 0; fi6 < 2; fi6++ {
                    v7 := GenS1{F0: -32768, F1: 171}
                    if (42 != -1) {
                        var v8 uint32 = 975
                        _ = v8
                    }
                    _ = fi6
                    _ = v7
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                switch "\n\t" {
                case "hello":
                    if ((-2147483648 < -2147483648) && (p1 || p1)) {
                        _ = len(v4)
                        p0 = 42
                        var v9 float32 = (3.14159 - 1e10)
                        _ = v9
                    } else {
                        if false {
                            p0 = (p0 / (((p0 >> ((4294967295) & 63))) | 1))
                        }
                        var v10 int16 = 1000
                        _ = v10
                    }
                default:
                    if (-2.25 > -2.25) {
                        var v11 bool = (0 > 255)
                        var v12 int8 = -57
                        var v13 uint = 456
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                }
                for fi14 := 0; fi14 < 2; fi14++ {
                    var v15 uint16 = 2
                    v4 = append(v4, fi14)
                    _ = fi14
                    _ = v15
                }
            }
            for fi16 := 0; fi16 < 13; fi16++ {
                switch "a" {
                case "\n\t":
                    var v17 float64 = (-2.25 * 1e10)
                    _ = v17
                default:
                    fmt.Println(p1)
                }
                var v18 int64 = p0
                if (42 > 594) {
                    var v19 string = ""
                    fmt.Println(fi16, v18)
                    _ = v19
                }
                _ = fi16
                _ = v18
            }
            _ = fi5
        }
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v20 int16 = 42
        _ = v20
    }
    if true {
        if (!p1) {
            v21 := GenS2{F0: p1}
            var v22 uint32 = 321
            _ = v21
            _ = v22
        } else {
            v23 := GenS3{F0: 753, F1: -671, F2: 390}
            _ = v23
        }
    } else {
        fmt.Println(p0, p0)
    }
    switch (("hello" + "") + "foo bar") {
    case "foo bar":
        var v24 bool = (939 < -2147483648)
        _ = v24
    case "hello":
        var v25 int8 = 42
        _ = v25
    case "a":
        p1 = p1
    default:
        if p1 {
            var v26 int64 = 359
            _ = v26
        } else {
            if p1 {
                var v27 int32 = 0
                _ = v27
            } else {
                var v28 string = "hello"
                var v29 int8 = 81
                if len(v4) > 0 {
                    v4[0] = 383
                }
                _ = v28
                _ = v29
            }
        }
    }
    for fi30 := 0; fi30 < 3; fi30++ {
        fmt.Println(p1)
        if true {
            p1 = p1
            if len(v4) > 0 {
                v4[0] = (^(fi30))
            }
        }
        var v31 uint32 = 2
        _ = fi30
        _ = v31
    }
    _ = p0
    _ = p1
    _ = v4
    return 100
}
func genFn1(p0 uint8, p1 int64) uint {
    fmt.Println(p0)
    if (("" + "a") != "") {
        var v32 int8 = -108
        v33 := []int{0, 551, 49}
        v34 := GenS3{F0: 65535, F1: 837, F2: 100}
        _ = v32
        _ = v33
        _ = v34
    } else {
        if ((0.1 - (-2.25 * 1.5)) > 3.14159) {
            p1 -= p1
            fmt.Println(p0)
        } else {
            v35 := GenS4{F0: 374, F1: 2147483647, F2: 74}
            _ = v35
        }
    }
    _ = p0
    _ = p1
    return 18446744073709551615
}

func main() {
    var v1 uint8 = 1
    var v2 float64 = 0.1
    v3 := make(map[string]bool)
    v1 += v1
    var v36 int8 = 42
    if true {
        if true {
            if (true || (0 <= 0)) {
                switch (("" + "hello") + "a") {
                case "a":
                    if true {
                        if (4294967295 < 4294967295) {
                            var v37 int16 = 0
                            if (0 <= 1) {
                                var v38 uint64 = 359
                                v38 = (v38 << ((genFn1((v1 << ((1) & 63)), -1)) & 63))
                                _ = v38
                            }
                            _ = v37
                        }
                        delete(v3, (("" + ("foo bar" + "hello")) + ("a" + "foo bar")))
                    } else {
                        switch -1 {
                        case -698:
                            if false {
                                for fi39 := 0; fi39 < 5; fi39++ {
                                    if false {
                                        for fi40 := 0; fi40 < 2; fi40++ {
                                            var v41 uint = 501
                                            v1 -= (v1 ^ 0)
                                            _ = fi40
                                            _ = v41
                                        }
                                    } else {
                                        if (!false) {
                                            v42 := v3["a"]
                                            fmt.Println(v2, v42)
                                            for fi43 := 0; fi43 < 19; fi43++ {
                                                v44 := v3[("hello" + "")]
                                                _ = fi43
                                                _ = v44
                                            }
                                            _ = v42
                                        }
                                        if false {
                                            if true {
                                                for fi45 := 0; fi45 < 10; fi45++ {
                                                    var v46 uint32 = 2
                                                    switch "" {
                                                    case "foo bar":
                                                        v47 := GenS5{F0: v46}
                                                        _ = v47
                                                    case "\n\t":
                                                        v36 = (v36 | (v36 >> ((1) & 63)))
                                                    case "a":
                                                        var v48 int = fi39
                                                        _ = v48
                                                    }
                                                    _ = fi45
                                                    _ = v46
                                                }
                                                v36 -= -128
                                                for fi49 := 0; fi49 < 16; fi49++ {
                                                    v1 += (v1 - 0)
                                                    _ = fi49
                                                }
                                            }
                                            v36 += -37
                                        } else {
                                            v36 = genFn0(243, false)
                                        }
                                    }
                                    switch (fi39 + (fi39 << ((genFn1((v1 / ((v1) | 1)), -9223372036854775808)) & 63))) {
                                    case 2147483647:
                                        switch (fi39 - (fi39 * 17)) {
                                        case -1:
                                            if false {
                                                switch "\n\t" {
                                                case "hello":
                                                    for fi50 := 0; fi50 < 1; fi50++ {
                                                        switch (fi39 % ((9223372036854775807) | 1)) {
                                                        case -9223372036854775808:
                                                            v51 := GenS6{F0: 42, F1: (0 == 0), F2: 98}
                                                            _ = v51
                                                        case 4294967295:
                                                            var v52 bool = false
                                                            _ = v52
                                                        default:
                                                            v53 := []int{(fi50 - 32768)}
                                                            _ = v53
                                                        }
                                                        var v54 int32 = 2147483647
                                                        _ = fi50
                                                        _ = v54
                                                    }
                                                }
                                                v36 -= 127
                                            }
                                        case 0:
                                            v55 := v3[("a" + ("\n\t" + "a"))]
                                            _ = v55
                                        case 32767:
                                            v56 := GenS4{F0: 724, F1: -9223372036854775808, F2: 0}
                                            _ = v56
                                        default:
                                            var v57 int = 128
                                            _ = v57
                                        }
                                    case 1:
                                        v58 := GenS7{F0: (v1 % ((v1) | 1))}
                                        _ = v58
                                    }
                                    _ = fi39
                                }
                            } else {
                                var v59 float64 = (3.14159 - v2)
                                var v60 int16 = 32767
                                _ = v59
                                _ = v60
                            }
                        case 32768:
                            for fi61 := 0; fi61 < 16; fi61++ {
                                for fi62 := 0; fi62 < 5; fi62++ {
                                    var v63 int32 = 934
                                    fmt.Println(v2, v63)
                                    _ = fi62
                                    _ = v63
                                }
                                _ = fi61
                            }
                        case 0:
                            switch ("hello" + "") {
                            case "hello":
                                if (!(!false)) {
                                    var v64 int32 = 2147483647
                                    _ = v64
                                }
                            case "foo bar":
                                v2 -= -2.25
                            case "a":
                                for fi65 := 0; fi65 < 16; fi65++ {
                                    var v66 uint = 1
                                    switch "" {
                                    case "hello":
                                        if true {
                                            for fi67 := 0; fi67 < 7; fi67++ {
                                                delete(v3, "a")
                                                _ = fi67
                                            }
                                            v68 := []int{(fi65 ^ (fi65 & 128)), fi65, (fi65 >> ((v66) & 63))}
                                            v69 := GenS8{F0: -1}
                                            _ = v68
                                            _ = v69
                                        }
                                    case "foo bar":
                                        var v70 int64 = 42
                                        _ = v70
                                    default:
                                        v71 := v3[("a" + "hello")]
                                        _ = v71
                                    }
                                    _ = fi65
                                    _ = v66
                                }
                            default:
                                var v72 uint32 = 65535
                                _ = v72
                            }
                        default:
                            if ((v2 - (v2 / 1.0)) > v2) {
                                var v73 float64 = v2
                                var v74 float64 = v73
                                _ = v73
                                _ = v74
                            }
                        }
                    }
                case "":
                    v75 := GenS9{F0: 65535, F1: 42, F2: 978}
                    _ = v75
                case "\n\t":
                    for fi76 := 0; fi76 < 9; fi76++ {
                        v2 += (0.1 / v2)
                        v2 = v2
                        for fi77 := 0; fi77 < 7; fi77++ {
                            v78 := []int{int(v1), 566}
                            v79 := GenS10{F0: 1}
                            _ = fi77
                            _ = v78
                            _ = v79
                        }
                        _ = fi76
                    }
                default:
                    fmt.Println(v1, v2)
                }
                var v80 int = -2147483648
                fmt.Println(v36, v2, v36)
                _ = v80
            } else {
                var v81 int32 = 1
                var v82 int32 = v81
                _ = v81
                _ = v82
            }
        } else {
            v83 := GenS11{F0: -1}
            v2 -= v2
            var v84 float32 = 1e10
            _ = v83
            _ = v84
        }
        var v85 int64 = 0
        fmt.Println(v2)
        _ = v85
    } else {
        v36 = int8(v1)
    }
    var v86 uint64 = 2
    var v87 uint8 = (v1 << ((42) & 63))
    var v88 int = 32767
    v88 += 128
    for fi89 := 0; fi89 < 3; fi89++ {
        for fi90 := 0; fi90 < 11; fi90++ {
            delete(v3, "a")
            v2 = v2
            if (2147483647 != -841) {
                var v91 float32 = (3.14159 / -2.25)
                v92 := []int{(fi89 % ((v88) | 1))}
                if true {
                    if true {
                        v91 += (((((-2.25 / v91) * 3.14159) * v91) * (v91 + (float32(v2) + ((1.5 / 0.1) / ((v91 + (v91 + float32(-2.25))) + (-2.25 / 1e10)))))) * 0.0)
                        fmt.Println(v88, v1)
                        v93 := GenS12{F0: 2147483647, F1: 32767}
                        _ = v93
                    } else {
                        switch "foo bar" {
                        case "":
                            var v94 uint64 = v86
                            _ = v94
                        case "a":
                            if (v36 > genFn0(int64(fi89), true)) {
                                _ = len(v3)
                            }
                        case "\n\t":
                            v95 := GenS13{F0: 72, F1: 42, F2: 1}
                            _ = v95
                        default:
                            if false {
                                switch "" {
                                case "":
                                    v96 := GenS9{F0: 542, F1: 65535, F2: 1}
                                    _ = v96
                                case "hello":
                                    v91 += ((float32(((((((((1.5 - (0.1 + (0.1 - ((0.0 - v2) * (v2 * v2))))) / v2) + v2) - (v2 + v2)) * (1e10 / float64(v91))) / -2.25) - ((((3.14159 * 3.14159) * -2.25) * (v2 - (float64(v91) / (v2 - 3.14159)))) + float64((3.14159 - v91)))) - v2)) + 1.5) - 1e10)
                                case "\n\t":
                                    v91 -= v91
                                default:
                                    v88 -= v88
                                }
                                var v97 uint16 = 2
                                fmt.Println(v91, v36)
                                _ = v97
                            } else {
                                v98 := GenS12{F0: 1, F1: 255}
                                fmt.Println(v91, v2, v1)
                                _ = v98
                            }
                        }
                        if false {
                            var v99 string = "\n\t"
                            _ = v99
                        }
                    }
                    var v100 int = -560
                    v101 := v3["\n\t"]
                    _ = v100
                    _ = v101
                }
                _ = v91
                _ = v92
            }
            switch "a" {
            case "\n\t":
                v1 += (v87 / ((v1) | 1))
            case "foo bar":
                var v102 int8 = (^(v36))
                _ = v102
            case "":
                var v103 bool = (fi89 == (v88 & fi90))
                _ = v103
            default:
                v104 := GenS12{F0: 943, F1: 948}
                _ = v104
            }
            fmt.Println(fi90, fi89)
            var v105 uint32 = 2
            fmt.Println(fi89, v2, v87)
            _ = fi90
            _ = v105
        }
        var v106 uint64 = v86
        _ = fi89
        _ = v106
    }
    var v107 uint32 = 0
    if (112 >= 274) {
        if ((v86 - (^(v86))) >= 4294967295) {
            v108 := GenS14{F0: 127}
            _ = v108
        }
        switch 32768 {
        case -9223372036854775808:
            switch v88 {
            case 685:
                if true {
                    switch (("" + "a") + "foo bar") {
                    case "\n\t":
                        if true {
                            switch (v88 - (-(v88))) {
                            case 32768:
                                fmt.Println(v86)
                            case 2147483647:
                                v109 := GenS11{F0: 431}
                                _ = v109
                            case 9223372036854775807:
                                v36 = (v36 >> ((15) & 63))
                            default:
                                v3["hello"] = (false && (false || ((-9223372036854775808 < (v88 + (v88 / (((v88 / ((v88) | 1))) | 1)))) || (true && true))))
                            }
                            var v110 int16 = 1
                            switch "foo bar" {
                            case "foo bar":
                                fmt.Println(v107, v87)
                            case "\n\t":
                                for fi111 := 0; fi111 < 15; fi111++ {
                                    var v112 int16 = (v110 * 1)
                                    var v113 uint8 = 128
                                    var v114 int16 = (v112 << ((genFn1(v1, 42)) & 63))
                                    _ = fi111
                                    _ = v112
                                    _ = v113
                                    _ = v114
                                }
                            case "a":
                                for fi115 := 0; fi115 < 1; fi115++ {
                                    switch "a" {
                                    case "\n\t":
                                        v110 -= (v110 << ((1) & 63))
                                    case "":
                                        var v116 int64 = -9223372036854775808
                                        _ = v116
                                    case "hello":
                                        switch "foo bar" {
                                        case "hello":
                                            var v117 string = "\n\t"
                                            _ = v117
                                        default:
                                            v86 = v86
                                        }
                                    }
                                    for fi118 := 0; fi118 < 13; fi118++ {
                                        var v119 int8 = v36
                                        v120 := []int{(fi115 / ((v88) | 1)), (fi115 / ((v88) | 1))}
                                        _ = fi118
                                        _ = v119
                                        _ = v120
                                    }
                                    _ = fi115
                                }
                            default:
                                v87 -= 128
                            }
                            _ = v110
                        } else {
                            v121 := GenS15{F0: 1}
                            _ = v121
                        }
                    }
                    v122 := GenS16{F0: uint64(v1)}
                    _ = v122
                } else {
                    delete(v3, "a")
                    v123 := GenS17{F0: 1}
                    switch 128 {
                    case 257:
                        var v124 int32 = 0
                        _ = v124
                    }
                    _ = v123
                }
            default:
                fmt.Println(v87)
            }
        case 32768:
            for fi125 := 0; fi125 < 17; fi125++ {
                delete(v3, "\n\t")
                _ = fi125
            }
        case 4294967295:
            v126 := v3["hello"]
            _ = v126
        default:
            v127 := GenS1{F0: 973, F1: 2}
            _ = v127
        }
    } else {
        v107 -= uint32(v86)
        _ = len(v3)
    }
    if false {
        v128 := []int{v88}
        _ = v128
    }
    if (!(!(((v2 + v2) <= float64(((0.1 / (-2.25 / 1.5)) - 0.0))) || true))) {
        v86 += v86
        delete(v3, ("foo bar" + ("foo bar" + ("\n\t" + ("\n\t" + "hello")))))
        for fi129 := 0; fi129 < 11; fi129++ {
            var v130 int64 = 910
            v2 = v2
            v36 += -40
            _ = fi129
            _ = v130
        }
    } else {
        fmt.Println(v2, v107, v88)
        v131 := GenS9{F0: v107, F1: 32767, F2: v107}
        if false {
            for fi132 := 0; fi132 < 13; fi132++ {
                _ = len(v3)
                fmt.Println(fi132, v86)
                _ = fi132
            }
            v36 += (v36 | 0)
            var v133 int = (v88 % (((^(v88))) | 1))
            _ = v133
        } else {
            if (!true) {
                v88 += 131
                delete(v3, "hello")
                fmt.Println(v107, v107)
            } else {
                v87 = uint8(v107)
                if false {
                    v86 -= (v86 ^ (v86 | (v86 & 42)))
                    fmt.Println(v2)
                } else {
                    v36 = (-(v36))
                    var v134 uint = 421
                    if (v107 != 441) {
                        v86 = (^(v86))
                        if true {
                            if false {
                                for fi135 := 0; fi135 < 17; fi135++ {
                                    v136 := v3["foo bar"]
                                    v137 := []int{int(v87)}
                                    _ = fi135
                                    _ = v136
                                    _ = v137
                                }
                                var v138 int32 = -228
                                v86 = v86
                                _ = v138
                            }
                            fmt.Println(v1)
                        }
                        v87 += 128
                    }
                    _ = v134
                }
                v139 := v3[(("a" + "hello") + "a")]
                _ = v139
            }
            for fi140 := 0; fi140 < 10; fi140++ {
                v141 := GenS18{F0: 0, F1: v107, F2: (^(v86))}
                _ = fi140
                _ = v141
            }
            var v142 float64 = (float64(((-2.25 - 0.0) / 0.1)) + v2)
            _ = v142
        }
        _ = v131
    }
    fmt.Println(v107, v88)
    v143 := []int{(v88 | 832), 4294967295, (v88 | v88), v88}
    v88 = v88
    v88 = v88
    var v144 float32 = (float32(v2) + 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v107:", v107)
    fmt.Println("slicelen v143:", len(v143), v143)
    fmt.Println("v144:", v144)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v107
    _ = v143
    _ = v144
    _ = v2
    _ = v3
    _ = v36
    _ = v86
    _ = v87
    _ = v88
}
