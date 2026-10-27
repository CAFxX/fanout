package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 uint8; F2 bool }
type GenS3 struct { F0 int64; F1 int8; F2 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32; F1 int64; F2 int32 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint; F1 uint64; F2 uint64 }
type GenS9 struct { F0 string; F1 uint }
type GenS10 struct { F0 int64; F1 string; F2 int8 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint; F1 string; F2 int16 }
type GenS13 struct { F0 string; F1 uint32 }
func genFn0(p0 int64) int {
    v4 := []int{-2147483648, 0, -9223372036854775808}
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := GenS1{F0: 2}
        var v7 int64 = 0
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS2{F0: 127, F1: 0, F2: (p0 <= 503)}
    var v9 int32 = 42
    for fi10 := 0; fi10 < 14; fi10++ {
        v11 := GenS3{F0: -9223372036854775808, F1: 127, F2: ""}
        if true {
            v9 = (v9 ^ (v9 ^ v9))
            if ((42 < (p0 / (((p0 % ((p0) | 1))) | 1))) && true) {
                var v12 int32 = (v9 % (((v9 % (((^(v9))) | 1))) | 1))
                _ = v12
            } else {
                p0 = (-(p0))
                var v13 bool = true
                v13 = true
                _ = v13
            }
            switch 933 {
            case -988:
                if false {
                    _ = len(v4)
                    v9 += -2147483648
                    switch 42 {
                    case -354:
                        var v14 uint64 = 18446744073709551615
                        _ = v14
                    case -410:
                        v15 := GenS4{F0: (-(p0))}
                        _ = v15
                    default:
                        for fi16 := 0; fi16 < 16; fi16++ {
                            _ = len(v4)
                            switch 4294967295 {
                            case 42:
                                if true {
                                    v17 := GenS4{F0: (p0 >> ((554) & 63))}
                                    var v18 uint16 = 42
                                    var v19 int = fi10
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                } else {
                                    fmt.Println(p0, fi10)
                                    var v20 uint32 = 2
                                    var v21 uint16 = 32767
                                    _ = v20
                                    _ = v21
                                }
                            default:
                                var v22 string = "a"
                                _ = v22
                            }
                            var v23 uint8 = 1
                            _ = fi16
                            _ = v23
                        }
                    }
                } else {
                    v24 := GenS5{F0: (-(p0))}
                    for fi25 := 0; fi25 < 8; fi25++ {
                        v9 -= -1
                        fmt.Println(p0)
                        _ = fi25
                    }
                    _ = v24
                }
            case 32767:
                for fi26 := 0; fi26 < 13; fi26++ {
                    var v27 float64 = (0.0 - (3.14159 / 3.14159))
                    var v28 int64 = (^(p0))
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            default:
                switch "a" {
                case "\n\t":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case "":
                    switch (fi10 & (fi10 ^ (fi10 * fi10))) {
                    case -441:
                        v9 -= 1
                    case -182:
                        for fi29 := 0; fi29 < 3; fi29++ {
                            var v30 int32 = (-(v9))
                            _ = fi29
                            _ = v30
                        }
                    case 714:
                        fmt.Println(p0)
                    default:
                        fmt.Println(fi10, fi10, fi10)
                    }
                case "a":
                    v31 := GenS6{F0: v9, F1: (p0 * p0), F2: (v9 >> ((2) & 63))}
                    _ = v31
                }
            }
        } else {
            var v32 string = "hello"
            if true {
                switch "hello" {
                case "a":
                    v33 := GenS7{F0: 127}
                    _ = v33
                case "\n\t":
                    switch "" {
                    case "":
                        for fi34 := 0; fi34 < 16; fi34++ {
                            v35 := GenS3{F0: (p0 - (p0 + (p0 | p0))), F1: -1, F2: "hello"}
                            var v36 string = (v32 + (("" + v32) + ("" + (v32 + v32))))
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                    default:
                        if (42 != 441) {
                            for fi37 := 0; fi37 < 9; fi37++ {
                                switch (^(fi10)) {
                                case 2147483647:
                                    v38 := GenS6{F0: v9, F1: (p0 % ((p0) | 1)), F2: (v9 / (((v9 % (((v9 << ((18446744073709551615) & 63))) | 1))) | 1))}
                                    _ = v38
                                case 190:
                                    switch "hello" {
                                    case "foo bar":
                                        var v39 uint16 = 1
                                        _ = v39
                                    case "":
                                        if len(v4) > 0 {
                                            v4[0] = fi10
                                        }
                                    default:
                                        var v40 string = (v32 + v32)
                                        _ = v40
                                    }
                                case -666:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                default:
                                    fmt.Println(fi10, fi10, p0)
                                }
                                p0 = (-(p0))
                                _ = fi37
                            }
                        }
                    }
                }
            } else {
                for fi41 := 0; fi41 < 11; fi41++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v9 += v9
                    v42 := GenS7{F0: 127}
                    v9 = v9
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi41
                    _ = v42
                }
                var v43 bool = true
                _ = v43
            }
            var v44 string = (v32 + ((v32 + (v32 + (v32 + "hello"))) + v32))
            _ = v32
            _ = v44
        }
        _ = len(v4)
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = v4
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := []uint32{0, 2147483647}
    v2 := make(map[int]int)
    v3 := []uint32{}
    v2[42] = 0
    v45 := GenS7{F0: 171}
    v46 := GenS6{F0: 42, F1: 42, F2: 2147483647}
    switch 128 {
    case 1:
        if (!(false && false)) {
            switch "a" {
            case "a":
                var v47 uint = 0
                _ = v47
            case "\n\t":
                var v48 bool = ((3.14159 / 0.1) >= 0.1)
                _ = v48
            case "hello":
                for fi49 := 0; fi49 < 15; fi49++ {
                    if (694 > 1) {
                        switch (-(fi49)) {
                        case 127:
                            v50 := GenS7{F0: 42}
                            _ = v50
                        case 1:
                            fmt.Println(fi49)
                        }
                    }
                    var v51 uint = 0
                    v51 -= (v51 * 42)
                    _ = fi49
                    _ = v51
                }
            default:
                var v52 uint = 42
                _ = v52
            }
            var v53 int32 = 2147483647
            _ = len(v2)
            _ = v53
        }
    }
    if ((679 >= 711) || (-1 < -1)) {
        v54 := GenS1{F0: 268}
        if true {
            var v55 bool = false
            _ = v55
        } else {
            for fi56 := 0; fi56 < 11; fi56++ {
                switch fi56 {
                case 1:
                    var v57 uint8 = 128
                    _ = v57
                default:
                    var v58 int8 = 114
                    _ = v58
                }
                if false {
                    if (("a" + ("" + "foo bar")) <= "foo bar") {
                        var v59 uint16 = 324
                        if len(v3) > 0 {
                            v3[0] = 65535
                        }
                        _ = v59
                    } else {
                        fmt.Println(fi56)
                    }
                    var v60 int = (fi56 & fi56)
                    _ = v60
                }
                var v61 int = fi56
                _ = fi56
                _ = v61
            }
            if false {
                _ = len(v2)
                var v62 string = "foo bar"
                _ = v62
            } else {
                v63 := GenS1{F0: 2}
                for fi64 := 0; fi64 < 3; fi64++ {
                    var v65 string = (("foo bar" + "") + ("\n\t" + "\n\t"))
                    _ = fi64
                    _ = v65
                }
                _ = v63
            }
        }
        v66 := v2[2147483647]
        _ = v54
        _ = v66
    } else {
        var v67 float64 = 0.0
        _ = v67
    }
    var v68 uint8 = 127
    if (!(0 == 42)) {
        v68 += 196
        var v69 float32 = 0.1
        switch "hello" {
        case "hello":
            switch "hello" {
            case "\n\t":
                v70 := v2[9223372036854775807]
                _ = v70
            case "":
                var v71 int16 = 1
                _ = v71
            case "a":
                fmt.Println(v69, v69)
            default:
                var v72 int8 = -128
                _ = v72
            }
        }
        _ = v69
    }
    var v73 float64 = float64((-2.25 + (((0.1 / 0.1) + (float32(1e10) * (((-2.25 + -2.25) / (1.5 / 0.1)) / ((0.0 + 1.5) * ((0.1 / (float32(1e10) / float32((3.14159 + (1e10 * 0.0))))) / (1.5 + -2.25)))))) * 3.14159)))
    v68 = 1
    var v74 string = ("\n\t" + "hello")
    if true {
        v75 := GenS7{F0: (v68 / ((v68) | 1))}
        var v76 float32 = ((0.1 / -2.25) + 0.0)
        _ = v75
        _ = v76
    } else {
        var v77 int64 = 242
        v78 := GenS3{F0: (v77 >> ((255) & 63)), F1: -1, F2: "a"}
        _ = v77
        _ = v78
    }
    fmt.Println(v73)
    v79 := GenS4{F0: 529}
    v74 = "hello"
    for fi80 := 0; fi80 < 11; fi80++ {
        var v81 uint8 = (v68 % ((v68) | 1))
        fmt.Println(v73, v74)
        _ = fi80
        _ = v81
    }
    for fi82 := 0; fi82 < 7; fi82++ {
        v2[fi82] = 611
        fmt.Println(v74, v73, v73)
        switch v74 {
        case "hello":
            if (2 > 742) {
                v3 = append(v3, 42)
                switch (fi82 / ((genFn0(-9223372036854775808)) | 1)) {
                case 43:
                    switch (fi82 ^ genFn0(0)) {
                    case 4294967295:
                        fmt.Println(fi82)
                    case 32768:
                        var v83 uint8 = (v68 >> ((uint(v68)) & 63))
                        _ = v83
                    default:
                        for fi84 := 0; fi84 < 2; fi84++ {
                            if (1 > -1) {
                                v85 := GenS3{F0: -175, F1: -1, F2: (((v74 + (v74 + "")) + v74) + v74)}
                                _ = v85
                            }
                            var v86 int8 = -1
                            v86 = v86
                            _ = fi84
                            _ = v86
                        }
                    }
                case 258:
                    var v87 float32 = 1.5
                    _ = v87
                }
                switch genFn0(367) {
                case 0:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case -624:
                    var v88 int16 = -32768
                    _ = v88
                case -158:
                    if (true || (true && (!false))) {
                        switch "\n\t" {
                        case "":
                            var v89 uint16 = 42
                            _ = v89
                        case "a":
                            v68 = v68
                        case "hello":
                            fmt.Println(fi82, v73, fi82)
                        default:
                            v68 -= (^(v68))
                        }
                    } else {
                        if true {
                            var v90 bool = true
                            var v91 uint64 = 0
                            _ = v90
                            _ = v91
                        } else {
                            var v92 float64 = 0.0
                            _ = v92
                        }
                        var v93 int16 = 32767
                        switch v74 {
                        case "foo bar":
                            for fi94 := 0; fi94 < 11; fi94++ {
                                v68 += (v68 >> ((863) & 63))
                                if false {
                                    if true {
                                        v68 += v68
                                    }
                                    if (!((!false) && true)) {
                                        v95 := GenS7{F0: v68}
                                        var v96 int = (fi82 % (((fi94 + (fi82 % (((fi94 << ((2) & 63))) | 1)))) | 1))
                                        switch v74 {
                                        case "foo bar":
                                            v68 = v68
                                        case "hello":
                                            var v97 uint8 = 127
                                            _ = v97
                                        case "\n\t":
                                            var v98 uint = 255
                                            _ = v98
                                        }
                                        _ = v95
                                        _ = v96
                                    } else {
                                        var v99 int16 = v93
                                        _ = v99
                                    }
                                    switch 9223372036854775807 {
                                    case 42:
                                        v100 := GenS8{F0: 1, F1: 42, F2: 42}
                                        _ = v100
                                    case -751:
                                        var v101 uint = 42
                                        _ = v101
                                    }
                                    var v102 float64 = (((3.14159 - v73) / 0.1) / 0.1)
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    if false {
                                        var v103 uint64 = 42
                                        switch "hello" {
                                        case "":
                                            fmt.Println(v103, fi94)
                                        default:
                                            v73 = float64(((float32(v73) * (3.14159 / ((3.14159 - 0.1) / 3.14159))) + 0.1))
                                        }
                                        _ = v103
                                    } else {
                                        if false {
                                            v104 := GenS4{F0: 1}
                                            v74 = "a"
                                            var v105 uint16 = 32767
                                            _ = v104
                                            _ = v105
                                        } else {
                                            _ = len(v1)
                                        }
                                        fmt.Println(v68)
                                        v102 = v73
                                        fmt.Println(v74)
                                        v102 = (v102 - v102)
                                        if len(v3) > 0 {
                                            v3[0] = 1
                                        }
                                        _ = len(v1)
                                        var v106 float64 = v73
                                        delete(v2, genFn0(42))
                                        fmt.Println(v74)
                                        v102 -= (float64(1e10) - v102)
                                        switch (("hello" + v74) + ("foo bar" + ("hello" + v74))) {
                                        case "a":
                                            switch v74 {
                                            case "foo bar":
                                                v107 := GenS1{F0: 42}
                                                _ = v107
                                            default:
                                                var v108 int = (fi94 - (fi82 | (-(fi94))))
                                                if true {
                                                    var v109 int64 = 42
                                                    if false {
                                                        v110 := GenS9{F0: ("a" + (("a" + "foo bar") + "")), F1: 42}
                                                        var v111 int16 = 42
                                                        _ = v110
                                                        _ = v111
                                                    }
                                                    v102 += (-2.25 + (float64(0.0) + 1.5))
                                                    _ = v109
                                                }
                                                v112 := GenS10{F0: -1, F1: v74, F2: 69}
                                                _ = v108
                                                _ = v112
                                            }
                                        default:
                                            v68 -= 2
                                        }
                                        _ = v106
                                    }
                                    _ = v102
                                } else {
                                    switch v74 {
                                    case "a":
                                        var v113 int8 = -1
                                        _ = v113
                                    case "\n\t":
                                        fmt.Println(fi82, v68)
                                    default:
                                        v73 = (float64(-2.25) - (float64(1.5) * v73))
                                    }
                                }
                                var v114 bool = (2 > 1)
                                if (v73 == v73) {
                                    var v115 int8 = -128
                                    v116 := GenS11{F0: 275}
                                    _ = v115
                                    _ = v116
                                } else {
                                    switch "\n\t" {
                                    case "\n\t":
                                        v73 = 3.14159
                                    case "foo bar":
                                        switch "foo bar" {
                                        case "foo bar":
                                            switch (fi94 << ((176) & 63)) {
                                            case 0:
                                                if true {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    var v117 float64 = -2.25
                                                    switch fi94 {
                                                    case 32767:
                                                        switch (v74 + "hello") {
                                                        case "\n\t":
                                                            var v118 string = ("\n\t" + v74)
                                                            _ = v118
                                                        default:
                                                            v119 := GenS9{F0: v74, F1: 361}
                                                            _ = v119
                                                        }
                                                        var v120 uint32 = 4294967295
                                                        var v121 uint16 = 1
                                                        v122 := GenS5{F0: 165}
                                                        _ = v120
                                                        _ = v121
                                                        _ = v122
                                                    }
                                                    _ = v117
                                                }
                                            default:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                        case "\n\t":
                                            v73 += float64(1.5)
                                        case "a":
                                            var v123 uint16 = 367
                                            _ = v123
                                        }
                                    case "a":
                                        var v124 uint = 0
                                        _ = v124
                                    }
                                }
                                _ = fi94
                                _ = v114
                            }
                        case "a":
                            v125 := v2[(fi82 / ((genFn0(1)) | 1))]
                            _ = v125
                        }
                        _ = v93
                    }
                default:
                    var v126 int32 = 733
                    _ = v126
                }
            } else {
                var v127 int8 = -106
                switch (fi82 + fi82) {
                case 42:
                    if len(v1) > 0 {
                        v1[0] = 1
                    }
                default:
                    if ((0 < 1) || false) {
                        v128 := GenS9{F0: v74, F1: 0}
                        _ = v128
                    } else {
                        switch fi82 {
                        case 0:
                            if false {
                                _ = len(v1)
                            }
                        }
                        v129 := v2[(fi82 % ((2147483647) | 1))]
                        if true {
                            v130 := GenS12{F0: 646, F1: "a", F2: -922}
                            fmt.Println(v73)
                            v131 := GenS13{F0: v74, F1: 65535}
                            _ = v130
                            _ = v131
                        } else {
                            switch v129 {
                            case -2147483648:
                                v73 -= (float64(float32((0.0 + v73))) + (1.5 / v73))
                            case -437:
                                var v132 int8 = v127
                                _ = v132
                            case 42:
                                v133 := GenS5{F0: 42}
                                _ = v133
                            default:
                                v74 = ""
                            }
                        }
                        _ = v129
                    }
                }
                _ = v127
            }
        case "foo bar":
            fmt.Println(fi82)
        case "":
            var v134 int = -198
            _ = v134
        default:
            var v135 uint16 = 255
            _ = v135
        }
        _ = fi82
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v68:", v68)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v46
    _ = v68
    _ = v73
    _ = v74
    _ = v79
}
