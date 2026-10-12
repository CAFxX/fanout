package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int; F1 int16 }
type GenS8 struct { F0 int; F1 uint; F2 uint8 }
type GenS9 struct { F0 int32; F1 string }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint64; F1 uint64; F2 bool }
type GenS12 struct { F0 uint; F1 int64; F2 bool }
func genFn0() int32 {
    v7 := GenS1{F0: -816}
    if (!true) {
        v8 := GenS1{F0: 303}
        _ = v8
    } else {
        var v9 int8 = -128
        v9 = 123
        _ = v9
    }
    v10 := GenS2{F0: (0 != -377), F1: false}
    _ = v10
    _ = v7
    return -578
}

func main() {
    var v1 float32 = 1.5
    v2 := []uint32{}
    var v3 uint32 = 2
    var v4 int32 = 2147483647
    v5 := []string{"a", "hello", (("\n\t" + "") + "hello"), "a"}
    v6 := []int{4294967295, 4294967295, -310}
    if (-997 != 32767) {
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 uint8 = 78
            switch int(v3) {
            case -1:
                switch "\n\t" {
                case "\n\t":
                    var v13 int16 = -32768
                    _ = v13
                case "a":
                    v2 = append(v2, v3)
                }
            default:
                var v14 float64 = 1.5
                _ = v14
            }
            _ = fi11
            _ = v12
        }
        v15 := GenS3{F0: 1, F1: int64(v3)}
        if false {
            v16 := GenS4{F0: (-608 < -894)}
            _ = v16
        }
        _ = v15
    } else {
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            var v17 float64 = 3.14159
            for fi18 := 0; fi18 < 13; fi18++ {
                if len(v5) > 0 {
                    v5[0] = "hello"
                }
                var v19 int8 = 124
                _ = fi18
                _ = v19
            }
            var v20 float64 = 3.14159
            _ = v17
            _ = v20
        }
        for fi21 := 0; fi21 < 12; fi21++ {
            var v22 int16 = -32768
            _ = fi21
            _ = v22
        }
        for fi23 := 0; fi23 < 12; fi23++ {
            switch (^(fi23)) {
            case -2147483648:
                switch 32767 {
                case 1:
                    if true {
                        var v24 float64 = (-2.25 - (((0.0 * 0.1) * 1e10) - (0.1 * (0.0 * 1e10))))
                        _ = v24
                    } else {
                        var v25 int32 = (v4 | v4)
                        if false {
                            var v26 int = 128
                            var v27 float32 = (1.5 * (v1 - v1))
                            v28 := GenS2{F0: ((v25 * 1) <= (v4 / ((v25) | 1))), F1: ((-(v4)) < v25)}
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                        _ = v25
                    }
                case -9223372036854775808:
                    fmt.Println(v1, v4, v1)
                case 128:
                    for fi29 := 0; fi29 < 8; fi29++ {
                        var v30 uint32 = v3
                        switch "" {
                        case "foo bar":
                            var v31 uint8 = 128
                            _ = v31
                        case "hello":
                            if true {
                                fmt.Println(fi29, v4)
                                fmt.Println(v4, fi29)
                                if true {
                                    v32 := GenS1{F0: -9223372036854775808}
                                    switch (fi23 << ((532) & 63)) {
                                    case 42:
                                        v1 -= v1
                                    case 32767:
                                        if false {
                                            v33 := GenS4{F0: ((0 >= -128) && (false || (42 == 2)))}
                                            var v34 uint64 = 4294967295
                                            v35 := GenS5{F0: v4}
                                            _ = v33
                                            _ = v34
                                            _ = v35
                                        }
                                        v3 -= v30
                                        v4 += v4
                                    case 127:
                                        if false {
                                            var v36 uint8 = 108
                                            v2 = append(v2, (v3 - v3))
                                            _ = v36
                                        } else {
                                            fmt.Println(v3, v1, v1)
                                        }
                                    default:
                                        if len(v5) > 0 {
                                            v5[0] = "hello"
                                        }
                                    }
                                    if true {
                                        v37 := GenS6{F0: 42}
                                        v3 -= 380
                                        v4 -= v4
                                        _ = v37
                                    }
                                    _ = v32
                                } else {
                                    var v38 bool = false
                                    if (v38 || false) {
                                        switch fi29 {
                                        case 32768:
                                            _ = len(v6)
                                        case 42:
                                            var v39 uint64 = 0
                                            _ = v39
                                        case -959:
                                            v30 -= v30
                                        }
                                    } else {
                                        v5 = append(v5, ("\n\t" + ""))
                                    }
                                    _ = v38
                                }
                                var v40 uint = 1
                                _ = v40
                            } else {
                                v41 := GenS6{F0: -1}
                                if false {
                                    v42 := GenS1{F0: 0}
                                    _ = v42
                                }
                                var v43 uint32 = (v30 | v3)
                                var v44 int16 = -1
                                if false {
                                    var v45 string = "foo bar"
                                    _ = v45
                                }
                                _ = v41
                                _ = v43
                                _ = v44
                            }
                        }
                        _ = fi29
                        _ = v30
                    }
                }
            case 32768:
                v3 += 1
            case -707:
                fmt.Println(v3, v1, v1)
            default:
                if (!((v4 - v4) >= (-(v4)))) {
                    var v46 uint8 = 15
                    for fi47 := 0; fi47 < 9; fi47++ {
                        v3 = 4294967295
                        var v48 uint = 416
                        _ = fi47
                        _ = v48
                    }
                    v1 += v1
                    _ = v46
                }
            }
            v6 = append(v6, (fi23 * 42))
            _ = fi23
        }
    }
    v49 := GenS5{F0: (v4 / ((v4) | 1))}
    if false {
        fmt.Println(v4)
        if true {
            if (!true) {
                for fi50 := 0; fi50 < 6; fi50++ {
                    v4 = (^(v4))
                    v4 = v4
                    _ = fi50
                }
                v51 := GenS2{F0: true, F1: (875 < 1)}
                _ = v51
            } else {
                var v52 uint32 = (v3 | (v3 >> ((18446744073709551615) & 63)))
                if true {
                    if (42 <= -730) {
                        if true {
                            fmt.Println(v3, v3)
                            for fi53 := 0; fi53 < 7; fi53++ {
                                for fi54 := 0; fi54 < 5; fi54++ {
                                    if (!(1 > 1)) {
                                        fmt.Println(v52)
                                        var v55 string = "a"
                                        _ = v55
                                    }
                                    fmt.Println(fi53, fi54, v4)
                                    _ = fi54
                                }
                                v52 -= (v3 ^ (^(v52)))
                                v4 += genFn0()
                                _ = fi53
                            }
                        } else {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            if len(v6) > 0 {
                                v6[0] = 32767
                            }
                            for fi56 := 0; fi56 < 10; fi56++ {
                                var v57 uint16 = 42
                                _ = fi56
                                _ = v57
                            }
                        }
                        v4 = genFn0()
                    }
                }
                v52 -= v3
                _ = v52
            }
            if false {
                var v58 uint64 = 4294967295
                if false {
                    v59 := GenS1{F0: -9223372036854775808}
                    if true {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        v60 := GenS5{F0: (v4 % ((v4) | 1))}
                        _ = v60
                    } else {
                        if true {
                            v4 = 1
                            if true {
                                for fi61 := 0; fi61 < 19; fi61++ {
                                    if true {
                                        v58 += v58
                                        for fi62 := 0; fi62 < 18; fi62++ {
                                            var v63 uint = 42
                                            _ = fi62
                                            _ = v63
                                        }
                                    } else {
                                        v1 += v1
                                        var v64 int16 = 32767
                                        var v65 bool = (0 >= 868)
                                        _ = v64
                                        _ = v65
                                    }
                                    switch (^(fi61)) {
                                    case -2147483648:
                                        fmt.Println(v4, v4)
                                    case 4294967295:
                                        if len(v2) > 0 {
                                            v2[0] = v3
                                        }
                                    }
                                    _ = fi61
                                }
                                for fi66 := 0; fi66 < 19; fi66++ {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    switch (("a" + ("hello" + "")) + ("" + (("a" + "hello") + "\n\t"))) {
                                    case "a":
                                        fmt.Println(v58, v1, v1)
                                    }
                                    _ = fi66
                                }
                            }
                            v67 := GenS7{F0: 153, F1: 0}
                            _ = v67
                        }
                        for fi68 := 0; fi68 < 10; fi68++ {
                            var v69 string = "\n\t"
                            fmt.Println(v69)
                            _ = fi68
                            _ = v69
                        }
                    }
                    _ = v59
                } else {
                    v1 = v1
                    v3 -= v3
                }
                for fi70 := 0; fi70 < 10; fi70++ {
                    for fi71 := 0; fi71 < 1; fi71++ {
                        v72 := GenS8{F0: (fi70 >> ((0) & 63)), F1: 0, F2: 0}
                        switch "\n\t" {
                        case "a":
                            v3 = (v3 | 2)
                            v58 = v58
                            if true {
                                switch "a" {
                                case "a":
                                    if (!(false && ((true || ((!((true && false) && (true || false))) && true)) || (42 >= 1000)))) {
                                        var v73 int32 = genFn0()
                                        v58 = v58
                                        _ = v73
                                    } else {
                                        v4 = 767
                                        var v74 int8 = 127
                                        _ = v74
                                    }
                                case "\n\t":
                                    var v75 string = ("\n\t" + ("\n\t" + "hello"))
                                    _ = v75
                                case "":
                                    v1 = (((v1 / (-2.25 / 1.5)) + v1) / v1)
                                }
                                var v76 float64 = 3.14159
                                _ = v76
                            } else {
                                switch ("a" + "a") {
                                case "":
                                    if true {
                                        v3 += 2147483647
                                        v77 := GenS6{F0: 42}
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        _ = v77
                                    } else {
                                        if (!((v4 - (v4 ^ (v4 << ((0) & 63)))) <= -764)) {
                                            fmt.Println(fi70)
                                            fmt.Println(v3)
                                        }
                                        v3 = (^(v3))
                                        var v78 string = "\n\t"
                                        _ = v78
                                    }
                                case "\n\t":
                                    var v79 int16 = -1
                                    _ = v79
                                case "foo bar":
                                    v80 := GenS3{F0: 42, F1: 1}
                                    _ = v80
                                }
                                fmt.Println(v1, v58)
                                var v81 float32 = 1e10
                                if false {
                                    v82 := GenS8{F0: fi70, F1: 0, F2: 0}
                                    v83 := GenS9{F0: genFn0(), F1: "foo bar"}
                                    _ = v82
                                    _ = v83
                                } else {
                                    var v84 float64 = -2.25
                                    _ = v84
                                }
                                if false {
                                    switch fi70 {
                                    case -1:
                                        var v85 bool = (false || ((!true) || true))
                                        _ = v85
                                    case 42:
                                        if (!false) {
                                            fmt.Println(v4, v4)
                                            switch (fi70 * fi70) {
                                            case 9223372036854775807:
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            default:
                                                if (false && ((-9223372036854775808 != int64(v58)) || true)) {
                                                    var v86 int16 = 1
                                                    v87 := GenS2{F0: (true || false), F1: true}
                                                    switch "" {
                                                    case "\n\t":
                                                        v6 = append(v6, -2147483648)
                                                    case "a":
                                                        var v88 int16 = v86
                                                        _ = v88
                                                    case "hello":
                                                        switch fi71 {
                                                        case 32768:
                                                            if true {
                                                                v1 = v1
                                                                v58 = (v58 | 843)
                                                                v3 -= v3
                                                            }
                                                        case -9223372036854775808:
                                                            fmt.Println(v3, v58)
                                                        case 42:
                                                            var v89 uint64 = 1
                                                            _ = v89
                                                        }
                                                    default:
                                                        v90 := GenS10{F0: "hello"}
                                                        _ = v90
                                                    }
                                                    _ = v86
                                                    _ = v87
                                                } else {
                                                    var v91 uint8 = 127
                                                    v3 -= (^(v3))
                                                    v1 -= ((-2.25 / ((v1 / (v1 + (-2.25 - v1))) * (v81 - v1))) * 1e10)
                                                    _ = v91
                                                }
                                            }
                                            switch fi70 {
                                            case -9223372036854775808:
                                                v92 := GenS6{F0: 0}
                                                _ = v92
                                            default:
                                                v81 -= v1
                                                var v93 uint8 = 127
                                                var v94 string = ("foo bar" + ("hello" + (("foo bar" + "a") + "hello")))
                                                _ = v93
                                                _ = v94
                                            }
                                            var v95 uint16 = 1
                                            _ = v95
                                        } else {
                                            if false {
                                                v58 += (v58 & (v58 / ((272) | 1)))
                                            } else {
                                                v96 := GenS9{F0: (v4 >> ((65535) & 63)), F1: "hello"}
                                                _ = v96
                                            }
                                            var v97 int32 = 579
                                            v81 = 3.14159
                                            _ = v97
                                        }
                                        var v98 float32 = float32((-2.25 + 0.1))
                                        v99 := GenS11{F0: (v58 - (v58 + v58)), F1: 715, F2: ((true || true) || false)}
                                        _ = v98
                                        _ = v99
                                    case 875:
                                        v100 := GenS10{F0: "hello"}
                                        _ = v100
                                    }
                                    if false {
                                        if (false && (((v58 >= 42) || true) || (-1 > -128))) {
                                            if len(v6) > 0 {
                                                v6[0] = 1
                                            }
                                        }
                                        var v101 uint = 281
                                        _ = v101
                                    }
                                    v81 += float32(3.14159)
                                } else {
                                    var v102 uint8 = 255
                                    _ = v102
                                }
                                v103 := GenS11{F0: (v58 * v58), F1: v58, F2: (569 != 42)}
                                _ = v103
                                _ = v81
                            }
                        case "hello":
                            _ = len(v6)
                        default:
                            var v104 uint8 = 0
                            _ = v104
                        }
                        if true {
                            var v105 float64 = (1.5 + (((((0.0 * (1.5 + 0.1)) / ((0.0 + 1.5) / ((float64(3.14159) - (3.14159 / 3.14159)) * -2.25))) + -2.25) - 3.14159) / 1.0))
                            if true {
                                if (!(!false)) {
                                    var v106 int64 = -9223372036854775808
                                    var v107 float64 = (v105 - float64(v1))
                                    if (42 < 0) {
                                        var v108 uint32 = (v3 / (((v3 ^ (v3 | (v3 ^ 339)))) | 1))
                                        switch fi71 {
                                        case 32767:
                                            if len(v6) > 0 {
                                                v6[0] = (fi70 * 4294967295)
                                            }
                                            var v109 float32 = 1.5
                                            if true {
                                                var v110 bool = true
                                                _ = v110
                                            } else {
                                                switch 880 {
                                                case 2147483647:
                                                    if (!(!((false || (32767 <= 182)) && false))) {
                                                        fmt.Println(v109, v105)
                                                    }
                                                }
                                                var v111 int16 = 42
                                                var v112 bool = true
                                                _ = v111
                                                _ = v112
                                            }
                                            v3 = v108
                                            switch "foo bar" {
                                            case "\n\t":
                                                var v113 uint = 65535
                                                _ = v113
                                            case "foo bar":
                                                switch (("\n\t" + "") + (("a" + "") + "")) {
                                                case "foo bar":
                                                    var v114 int64 = (-(v106))
                                                    _ = v114
                                                case "a":
                                                    v108 -= v3
                                                }
                                            case "":
                                                switch ("a" + "\n\t") {
                                                case "hello":
                                                    v3 = v108
                                                }
                                            default:
                                                if false {
                                                    switch fi71 {
                                                    case 4294967295:
                                                        var v115 uint64 = 4294967295
                                                        _ = v115
                                                    case 137:
                                                        var v116 uint = 0
                                                        _ = v116
                                                    }
                                                }
                                            }
                                            var v117 bool = true
                                            _ = v109
                                            _ = v117
                                        case 42:
                                            var v118 uint32 = v3
                                            _ = v118
                                        }
                                        _ = v108
                                    } else {
                                        v107 -= 1e10
                                        v107 = v105
                                        _ = len(v5)
                                    }
                                    _ = v106
                                    _ = v107
                                } else {
                                    if (255 < 648) {
                                        switch "hello" {
                                        case "hello":
                                            fmt.Println(v58, fi70, v58)
                                        default:
                                            var v119 uint = 4294967295
                                            _ = v119
                                        }
                                        if false {
                                            if (42 <= int64(v4)) {
                                                var v120 int = (fi71 + (fi71 * fi70))
                                                if (-109 == 296) {
                                                    v58 += (v58 & (v58 + v58))
                                                    var v121 int = (^(fi70))
                                                    _ = v121
                                                }
                                                _ = v120
                                            } else {
                                                var v122 uint16 = 560
                                                if false {
                                                    v58 -= (v58 >> ((uint(v122)) & 63))
                                                }
                                                switch (fi70 * fi70) {
                                                case 70:
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                                _ = v122
                                            }
                                            v123 := GenS3{F0: 1, F1: -1}
                                            switch (fi70 >> ((18446744073709551615) & 63)) {
                                            case -1:
                                                var v124 uint = 2
                                                _ = v124
                                            case -82:
                                                if true {
                                                    v125 := GenS2{F0: false, F1: false}
                                                    _ = v125
                                                } else {
                                                    if (true || false) {
                                                        v58 += v58
                                                    }
                                                    if (!(!false)) {
                                                        v6 = append(v6, fi71)
                                                        var v126 int = (fi71 >> ((4294967295) & 63))
                                                        v4 = (v4 * v4)
                                                        _ = v126
                                                    }
                                                }
                                            }
                                            if (100 == 42) {
                                                _ = len(v2)
                                            } else {
                                                switch 1 {
                                                case 4294967295:
                                                    if ((-711 == 0) || true) {
                                                        var v127 float64 = v105
                                                        _ = v127
                                                    }
                                                default:
                                                    var v128 uint32 = v3
                                                    _ = v128
                                                }
                                                fmt.Println(v4)
                                                v3 = (^(v3))
                                            }
                                            if true {
                                                v129 := GenS8{F0: -967, F1: 18446744073709551615, F2: 94}
                                                var v130 int = (^(fi71))
                                                v131 := GenS10{F0: ((("hello" + "foo bar") + ("hello" + (((("a" + "") + "") + "") + "a"))) + "hello")}
                                                _ = v129
                                                _ = v130
                                                _ = v131
                                            } else {
                                                fmt.Println(v4, v1)
                                                switch "a" {
                                                case "foo bar":
                                                    v132 := GenS1{F0: -9223372036854775808}
                                                    _ = v132
                                                }
                                            }
                                            _ = v123
                                        } else {
                                            var v133 uint32 = v3
                                            var v134 bool = (!((fi70 / ((fi71) | 1)) != fi70))
                                            _ = v133
                                            _ = v134
                                        }
                                        var v135 int16 = -1
                                        _ = v135
                                    } else {
                                        var v136 bool = (!(((v105 == v105) || false) || (-642 == -1)))
                                        fmt.Println(v136, v1)
                                        fmt.Println(v136)
                                        _ = v136
                                    }
                                    var v137 string = "foo bar"
                                    if (32767 >= 1) {
                                        v138 := GenS6{F0: 0}
                                        _ = v138
                                    }
                                    _ = v137
                                }
                                fmt.Println(fi70)
                            } else {
                                v2 = append(v2, (v3 & v3))
                            }
                            _ = v105
                        }
                        _ = fi71
                        _ = v72
                    }
                    fmt.Println(v3, v1, v58)
                    _ = fi70
                }
                _ = v58
            } else {
                fmt.Println(v3)
            }
        }
        for fi139 := 0; fi139 < 17; fi139++ {
            for fi140 := 0; fi140 < 8; fi140++ {
                v3 -= v3
                _ = fi140
            }
            switch fi139 {
            case 4294967295:
                var v141 uint16 = 642
                _ = v141
            case 936:
                v1 = v1
            case -9223372036854775808:
                v3 = (v3 ^ (v3 % ((v3) | 1)))
            }
            _ = fi139
        }
    }
    var v142 uint64 = 816
    v143 := GenS12{F0: 0, F1: 1, F2: (926 >= uint16(v142))}
    var v144 int8 = 1
    var v145 uint8 = 127
    for fi146 := 0; fi146 < 5; fi146++ {
        v147 := GenS2{F0: (!(0 != -744)), F1: (v3 > (v3 % ((v3) | 1)))}
        v148 := GenS2{F0: (!true), F1: (-19 != int(v145))}
        v3 = (v3 * 284)
        _ = fi146
        _ = v147
        _ = v148
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (!true) {
        var v149 uint8 = 0
        switch "foo bar" {
        case "\n\t":
            var v150 uint = 1
            _ = v150
        case "hello":
            v149 = v145
        default:
            switch -151 {
            case 32768:
                var v151 int8 = -128
                _ = v151
            }
        }
        _ = v149
    }
    var v152 uint64 = v142
    for fi153 := 0; fi153 < 18; fi153++ {
        v5 = append(v5, "")
        var v154 int32 = (v4 * v4)
        _ = fi153
        _ = v154
    }
    fmt.Println("v1:", v1)
    fmt.Println("v142:", v142)
    fmt.Println("v144:", v144)
    fmt.Println("v145:", v145)
    fmt.Println("v152:", v152)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v142
    _ = v143
    _ = v144
    _ = v145
    _ = v152
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}
