package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint16; F1 int16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint8; F1 int64; F2 uint }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 string; F1 uint16; F2 uint32 }
func genFn0(p0 uint8, p1 int16, p2 uint64) int8 {
    fmt.Println(p0)
    var v6 float64 = 1e10
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 uint16 = 700
        v9 := GenS1{F0: p0}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: (^(p0))}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    return -128
}
func genFn1(p0 int8, p1 int8) uint64 {
    v11 := []int{32768, -1, 130}
    for fi12 := 0; fi12 < 14; fi12++ {
        if false {
            var v13 int32 = 0
            _ = v13
        }
        var v14 string = "hello"
        switch (-(fi12)) {
        case 4294967295:
            v14 = v14
        case -876:
            if ((!false) && true) {
                v15 := GenS2{F0: true, F1: -585}
                if true {
                    v11 = append(v11, (fi12 ^ (fi12 % (((fi12 << ((4294967295) & 63))) | 1))))
                    for fi16 := 0; fi16 < 4; fi16++ {
                        p0 -= p1
                        var v17 uint32 = 508
                        fmt.Println(v17, v17, p0)
                        v14 = "foo bar"
                        if (!(-718 != 42)) {
                            if false {
                                fmt.Println(fi16, p1)
                                if false {
                                    var v18 int8 = genFn0(157, 42, 4294967295)
                                    var v19 int32 = -257
                                    _ = v18
                                    _ = v19
                                }
                            }
                        } else {
                            var v20 int16 = -952
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                            fmt.Println(p0, p0, v14)
                            _ = v20
                        }
                        var v21 float64 = -2.25
                        v11 = append(v11, 935)
                        _ = fi16
                        _ = v17
                        _ = v21
                    }
                    fmt.Println(fi12, p0)
                }
                _ = v15
            } else {
                v22 := GenS2{F0: true, F1: 392}
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                var v23 int32 = -1
                _ = v22
                _ = v23
            }
        }
        _ = fi12
        _ = v14
    }
    var v24 bool = (1.5 == (3.14159 * 0.0))
    _ = p0
    _ = p1
    _ = v11
    _ = v24
    return 61
}

func main() {
    var v1 float64 = 1e10
    var v2 int64 = -9223372036854775808
    var v3 int32 = -1
    v4 := []int{32767}
    v5 := make(map[uint32]uint32)
    if (!true) {
        switch 2147483647 {
        case 226:
            _ = len(v5)
        case 587:
            if ((!(644 < (v2 >> ((65535) & 63)))) || (255 == 650)) {
                for fi25 := 0; fi25 < 10; fi25++ {
                    var v26 int16 = -32768
                    v2 -= v2
                    for fi27 := 0; fi27 < 17; fi27++ {
                        v26 = (v26 & v26)
                        var v28 float32 = 1e10
                        v29 := v5[42]
                        var v30 uint64 = genFn1(1, 0)
                        fmt.Println(v2, v26)
                        _ = fi27
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                    _ = fi25
                    _ = v26
                }
                if (!(!(v1 < (v1 / (v1 - float64(3.14159)))))) {
                    var v31 float64 = ((((float64((float32(0.0) - (float32(float64(0.0)) / -2.25))) * ((0.1 + ((1.5 * v1) / v1)) - (v1 * (1.5 / v1)))) / (v1 + 1.5)) * ((v1 + ((v1 + v1) + ((v1 - v1) + (v1 * 1e10)))) + v1)) / v1)
                    if false {
                        v3 = (v3 / (((v3 % (((v3 & (v3 / ((v3) | 1)))) | 1))) | 1))
                        if false {
                            var v32 uint8 = 128
                            _ = v32
                        } else {
                            fmt.Println(v1, v31)
                            _ = len(v4)
                            fmt.Println(v3, v3, v31)
                        }
                        fmt.Println(v3)
                    } else {
                        v33 := GenS1{F0: 147}
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 7; fi34++ {
                        v2 = v2
                        if (((float64(1.5) * (v1 + v1)) / (float64(1.5) * v31)) != v1) {
                            v31 -= v1
                            for fi35 := 0; fi35 < 3; fi35++ {
                                v36 := GenS3{F0: v2, F1: 505}
                                fmt.Println(fi35)
                                var v37 bool = (true || ((!true) && (-1 >= 0)))
                                if v37 {
                                    v37 = (-128 <= genFn0(42, 1000, 2))
                                    var v38 uint16 = 68
                                    v39 := v5[356]
                                    _ = v38
                                    _ = v39
                                } else {
                                    var v40 uint = 18446744073709551615
                                    v41 := v5[934]
                                    _ = v40
                                    _ = v41
                                }
                                _ = fi35
                                _ = v36
                                _ = v37
                            }
                            var v42 string = ("a" + "hello")
                            _ = v42
                        }
                        for fi43 := 0; fi43 < 6; fi43++ {
                            _ = len(v4)
                            v3 += v3
                            _ = fi43
                        }
                        _ = fi34
                    }
                    _ = v31
                } else {
                    fmt.Println(v1, v1, v2)
                }
            } else {
                for fi44 := 0; fi44 < 18; fi44++ {
                    v1 -= (v1 * (v1 / v1))
                    switch (fi44 / ((127) | 1)) {
                    case -190:
                        fmt.Println(v3, fi44, v1)
                    case -2147483648:
                        v4 = append(v4, -9223372036854775808)
                    }
                    for fi45 := 0; fi45 < 6; fi45++ {
                        if true {
                            if true {
                                var v46 int32 = -576
                                switch "" {
                                case "\n\t":
                                    v4 = append(v4, (fi44 | fi45))
                                default:
                                    switch (fi45 * (-(fi44))) {
                                    case 1:
                                        v47 := GenS4{F0: 1, F1: 0}
                                        _ = v47
                                    }
                                }
                                _ = len(v4)
                                var v48 string = ("hello" + ("hello" + "a"))
                                if true {
                                    if true {
                                        v49 := GenS5{F0: 99}
                                        _ = v49
                                    }
                                    if true {
                                        _ = len(v5)
                                    }
                                    v50 := GenS6{F0: 32767, F1: -1}
                                    _ = v50
                                }
                                var v51 int16 = 1
                                v52 := GenS7{F0: v48}
                                v2 = (-(v2))
                                _ = v46
                                _ = v48
                                _ = v51
                                _ = v52
                            } else {
                                v53 := GenS3{F0: int64(v3), F1: 66}
                                v1 = (v1 - float64(3.14159))
                                switch "" {
                                case "":
                                    v54 := GenS4{F0: 0, F1: 0}
                                    _ = v54
                                }
                                _ = v53
                            }
                            v5[732] = 42
                        }
                        v4 = append(v4, fi44)
                        _ = fi45
                    }
                    _ = fi44
                }
                var v55 float32 = 1e10
                _ = v55
            }
        default:
            if (1 >= -284) {
                var v56 int32 = v3
                var v57 uint = 2
                _ = v56
                _ = v57
            } else {
                var v58 uint16 = 42
                var v59 uint16 = 431
                _ = v58
                _ = v59
            }
        }
    } else {
        var v60 uint64 = 425
        v61 := GenS1{F0: 212}
        _ = v60
        _ = v61
    }
    v62 := GenS8{F0: 100, F1: (v2 - (v2 << ((65535) & 63))), F2: 767}
    for fi63 := 0; fi63 < 3; fi63++ {
        v3 -= (v3 & -459)
        fmt.Println(fi63, v1, fi63)
        if (false && true) {
            var v64 uint = 42
            v5[0] = 270
            _ = v64
        } else {
            if (-128 > 1) {
                var v65 uint8 = 255
                _ = v65
            }
            for fi66 := 0; fi66 < 8; fi66++ {
                var v67 bool = false
                var v68 uint32 = 65535
                v1 -= v1
                if ((127 != 128) && v67) {
                    v2 -= (v2 * 9223372036854775807)
                    if (v67 || v67) {
                        if (2 == 582) {
                            var v69 uint64 = 0
                            v3 += v3
                            _ = v69
                        } else {
                            if v67 {
                                if v67 {
                                    v2 = (v2 << ((799) & 63))
                                    var v70 bool = (32767 == 32767)
                                    v67 = (0 <= 826)
                                    _ = v70
                                } else {
                                    var v71 uint8 = 42
                                    _ = v71
                                }
                                v72 := GenS3{F0: v2, F1: 42}
                                _ = v72
                            }
                            var v73 bool = (!false)
                            fmt.Println(fi63)
                            _ = v73
                        }
                        if v67 {
                            var v74 float64 = v1
                            var v75 int32 = (v3 & v3)
                            _ = v74
                            _ = v75
                        }
                    } else {
                        if (2 <= 1) {
                            v3 = v3
                        }
                    }
                } else {
                    var v76 uint16 = 553
                    switch fi63 {
                    case -954:
                        v77 := GenS4{F0: -1, F1: 42}
                        _ = v77
                    default:
                        fmt.Println(v2, v1)
                    }
                    _ = v76
                }
                v78 := GenS1{F0: 2}
                _ = fi66
                _ = v67
                _ = v68
                _ = v78
            }
            v79 := GenS3{F0: (v2 & -173), F1: 631}
            _ = v79
        }
        _ = fi63
    }
    switch 32768 {
    case -9223372036854775808:
        v3 = v3
    case 975:
        for fi80 := 0; fi80 < 15; fi80++ {
            v2 = (-(v2))
            v3 = 700
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi80
        }
    default:
        var v81 uint64 = 2
        _ = v81
    }
    v3 = v3
    fmt.Println(v3, v2, v1)
    if (!(!(!(!false)))) {
        v82 := v5[2147483648]
        var v83 uint64 = 741
        _ = v82
        _ = v83
    }
    v1 += v1
    var v84 uint64 = 18446744073709551615
    var v85 uint8 = 128
    if false {
        if (v3 == v3) {
            if false {
                var v86 int = 2147483647
                v87 := GenS3{F0: (v2 << ((918) & 63)), F1: 65535}
                for fi88 := 0; fi88 < 6; fi88++ {
                    v89 := GenS9{F0: (v2 ^ (^(v2)))}
                    switch ("a" + "\n\t") {
                    case "a":
                        fmt.Println(v86, v2, fi88)
                    case "":
                        var v90 int32 = (v3 + (v3 >> ((445) & 63)))
                        _ = v90
                    case "\n\t":
                        var v91 int64 = (-(v2))
                        _ = v91
                    }
                    v3 -= 0
                    _ = fi88
                    _ = v89
                }
                _ = v86
                _ = v87
            }
            v2 -= 1
            var v92 uint8 = v85
            _ = v92
        }
        if (!((false || (0.1 <= 0.1)) && (!false))) {
            for fi93 := 0; fi93 < 18; fi93++ {
                if true {
                    delete(v5, 4294967295)
                    for fi94 := 0; fi94 < 2; fi94++ {
                        fmt.Println(v84, v85, v3)
                        var v95 int = fi94
                        if len(v4) > 0 {
                            v4[0] = fi94
                        }
                        v85 = 42
                        _ = fi94
                        _ = v95
                    }
                    if len(v4) > 0 {
                        v4[0] = fi93
                    }
                }
                v96 := GenS3{F0: (v2 ^ -9223372036854775808), F1: 65535}
                var v97 float64 = (v1 + 0.1)
                _ = fi93
                _ = v96
                _ = v97
            }
            if (327 > 708) {
                v5[65535] = 2
                switch "foo bar" {
                case "a":
                    if len(v4) > 0 {
                        v4[0] = -593
                    }
                case "":
                    v2 = (v2 / (((v2 / ((v2) | 1))) | 1))
                default:
                    if (-1 == 100) {
                        for fi98 := 0; fi98 < 20; fi98++ {
                            v99 := GenS4{F0: 32767, F1: 42}
                            for fi100 := 0; fi100 < 3; fi100++ {
                                if false {
                                    v2 += int64(fi98)
                                }
                                fmt.Println(fi98, fi100, fi98)
                                _ = fi100
                            }
                            _ = fi98
                            _ = v99
                        }
                        switch -1 {
                        case -846:
                            switch 127 {
                            case 127:
                                var v101 float32 = 1.5
                                _ = v101
                            case -519:
                                v102 := GenS3{F0: 42, F1: 670}
                                _ = v102
                            case 878:
                                for fi103 := 0; fi103 < 5; fi103++ {
                                    fmt.Println(v1, v85, fi103)
                                    _ = fi103
                                }
                            default:
                                for fi104 := 0; fi104 < 19; fi104++ {
                                    var v105 float32 = 1.5
                                    var v106 uint16 = 43
                                    v1 -= (0.0 * 3.14159)
                                    _ = fi104
                                    _ = v105
                                    _ = v106
                                }
                            }
                        case 127:
                            for fi107 := 0; fi107 < 2; fi107++ {
                                var v108 float32 = 0.0
                                v3 = (v3 - (v3 * v3))
                                _ = fi107
                                _ = v108
                            }
                        default:
                            for fi109 := 0; fi109 < 6; fi109++ {
                                v110 := GenS10{F0: "", F1: 32767, F2: 42}
                                v111 := GenS6{F0: 444, F1: -32768}
                                _ = fi109
                                _ = v110
                                _ = v111
                            }
                        }
                    } else {
                        if false {
                            switch -475 {
                            case 1:
                                fmt.Println(v2, v3, v1)
                            case -9223372036854775808:
                                v1 += 0.0
                            default:
                                for fi112 := 0; fi112 < 12; fi112++ {
                                    var v113 bool = false
                                    switch 9223372036854775807 {
                                    case -9223372036854775808:
                                        var v114 string = ("\n\t" + "a")
                                        _ = v114
                                    }
                                    _ = fi112
                                    _ = v113
                                }
                            }
                            if true {
                                switch int(v2) {
                                case 288:
                                    if false {
                                        for fi115 := 0; fi115 < 15; fi115++ {
                                            v116 := GenS4{F0: int16(v84), F1: 42}
                                            v84 = genFn1(-78, -1)
                                            _ = fi115
                                            _ = v116
                                        }
                                        v84 = (^(v84))
                                        switch 4294967295 {
                                        case 2147483647:
                                            for fi117 := 0; fi117 < 5; fi117++ {
                                                switch 9223372036854775807 {
                                                case -207:
                                                    v118 := GenS7{F0: "\n\t"}
                                                    _ = v118
                                                case -564:
                                                    fmt.Println(v2)
                                                default:
                                                    if ("a" == ("hello" + (("a" + "\n\t") + ("hello" + "a")))) {
                                                        if (448 <= 65535) {
                                                            var v119 uint64 = uint64(v85)
                                                            fmt.Println(v1, fi117, v119)
                                                            _ = v119
                                                        } else {
                                                            var v120 uint32 = 2147483647
                                                            v121 := v5[v120]
                                                            switch -9223372036854775808 {
                                                            case 32768:
                                                                _ = len(v5)
                                                            case -1:
                                                                for fi122 := 0; fi122 < 14; fi122++ {
                                                                    v123 := GenS3{F0: v2, F1: 0}
                                                                    v124 := GenS2{F0: false, F1: 1}
                                                                    if true {
                                                                        v84 = (v84 << ((357) & 63))
                                                                    }
                                                                    _ = fi122
                                                                    _ = v123
                                                                    _ = v124
                                                                }
                                                            case 551:
                                                                _ = len(v4)
                                                            }
                                                            _ = v120
                                                            _ = v121
                                                        }
                                                        v125 := GenS7{F0: (("a" + "") + "")}
                                                        v2 += -461
                                                        _ = v125
                                                    }
                                                }
                                                _ = fi117
                                            }
                                        }
                                    } else {
                                        fmt.Println(v84)
                                        for fi126 := 0; fi126 < 6; fi126++ {
                                            if (42 <= 2) {
                                                fmt.Println(v85, v84)
                                                v3 = -1
                                                v84 -= (v84 & v84)
                                            }
                                            v4 = append(v4, (^(fi126)))
                                            _ = fi126
                                        }
                                        v85 -= (^(v85))
                                    }
                                case 839:
                                    v5[42] = 2
                                }
                                switch (("a" + "a") + "a") {
                                case "a":
                                    v84 -= (v84 >> ((65535) & 63))
                                case "\n\t":
                                    if len(v4) > 0 {
                                        v4[0] = 1
                                    }
                                default:
                                    v2 -= v2
                                }
                            }
                            var v127 int16 = 1
                            _ = v127
                        } else {
                            v84 = v84
                            v2 = int64(v85)
                            if (!(true && (-1 <= genFn0(0, 42, (v84 ^ 1))))) {
                                switch 42 {
                                case 32767:
                                    fmt.Println(v85, v1, v84)
                                default:
                                    v128 := GenS5{F0: 0}
                                    _ = v128
                                }
                                var v129 uint64 = (^(v84))
                                v130 := GenS9{F0: 410}
                                _ = v129
                                _ = v130
                            }
                        }
                    }
                }
            }
            var v131 float64 = (3.14159 / float64((-2.25 * (-2.25 - (-2.25 - 1e10)))))
            _ = v131
        }
        v85 += 145
    }
    for fi132 := 0; fi132 < 14; fi132++ {
        fmt.Println(v85, v3)
        if (1 < -32768) {
            v3 = v3
            if false {
                switch "hello" {
                case "":
                    if false {
                        v84 -= (v84 * 4294967295)
                        var v133 uint64 = v84
                        fmt.Println(v3)
                        _ = v133
                    }
                case "\n\t":
                    delete(v5, 2147483648)
                case "a":
                    v134 := GenS5{F0: -128}
                    _ = v134
                default:
                    var v135 bool = (true || true)
                    _ = v135
                }
                for fi136 := 0; fi136 < 14; fi136++ {
                    var v137 uint8 = 128
                    v85 = (v137 >> ((1) & 63))
                    fmt.Println(v2, v3, v137)
                    _ = fi136
                    _ = v137
                }
            }
            switch (fi132 << ((0) & 63)) {
            case 128:
                var v138 uint = 0
                _ = v138
            case 4294967295:
                if len(v4) > 0 {
                    v4[0] = (fi132 / ((32767) | 1))
                }
            case 32767:
                switch (("" + "") + ("\n\t" + "\n\t")) {
                case "":
                    if true {
                        fmt.Println(v85, v85, v85)
                    }
                case "a":
                    v5[740] = 675
                }
            }
        }
        _ = fi132
    }
    if (false && true) {
        if (1 <= 0) {
            v1 += float64(0.0)
        }
        fmt.Println(v85, v3)
    } else {
        if len(v4) > 0 {
            v4[0] = -133
        }
        var v139 int64 = (v2 - 720)
        for fi140 := 0; fi140 < 19; fi140++ {
            var v141 float32 = float32(float64(1.5))
            _ = fi140
            _ = v141
        }
        _ = v139
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v62
    _ = v84
    _ = v85
}
