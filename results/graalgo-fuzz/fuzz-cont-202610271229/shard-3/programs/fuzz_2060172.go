package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS5 struct { F0 int; F1 int64; F2 int }
type GenS6 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS7 struct { F0 uint8; F1 uint32 }
type GenS8 struct { F0 uint16; F1 int32; F2 int8 }
type GenS9 struct { F0 string; F1 uint; F2 uint32 }
type GenS10 struct { F0 string; F1 uint8; F2 uint64 }
type GenS11 struct { F0 bool; F1 int64; F2 int }
type GenS12 struct { F0 int64; F1 uint32; F2 uint8 }
func genFn0(p0 int, p1 uint32, p2 uint8) uint32 {
    var v7 bool = true
    p1 = p1
    var v8 int16 = -32768
    for fi9 := 0; fi9 < 15; fi9++ {
        if ("\n\t" == ("hello" + "")) {
            p0 = 32767
            switch "\n\t" {
            case "a":
                var v10 int32 = -119
                _ = v10
            }
        } else {
            for fi11 := 0; fi11 < 14; fi11++ {
                if (("hello" >= "") && true) {
                    switch (("a" + "foo bar") + ("\n\t" + ("\n\t" + "a"))) {
                    case "":
                        switch (fi11 | (p0 / ((fi9) | 1))) {
                        case 127:
                            var v12 int8 = 127
                            v13 := GenS1{F0: v8}
                            fmt.Println(p0, fi11)
                            _ = v12
                            _ = v13
                        case 0:
                            fmt.Println(fi11, fi11, p0)
                        case -405:
                            switch p0 {
                            case -453:
                                p0 -= fi9
                            case -861:
                                var v14 uint32 = (p1 ^ 4294967295)
                                _ = v14
                            case -9223372036854775808:
                                var v15 uint8 = 0
                                _ = v15
                            default:
                                switch fi11 {
                                case 42:
                                    v16 := []int{128, 32767, fi9}
                                    _ = v16
                                default:
                                    p0 = fi9
                                }
                            }
                        default:
                            v8 -= (^(v8))
                        }
                    default:
                        v17 := []int{-796, 1, (p0 + 32768)}
                        _ = v17
                    }
                    var v18 int8 = -127
                    _ = v18
                } else {
                    var v19 uint32 = p1
                    _ = v19
                }
                v7 = v7
                _ = fi11
            }
            v8 += 359
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p1
}

func main() {
    v1 := []string{(("" + ("\n\t" + ("foo bar" + ""))) + ("foo bar" + "a")), "foo bar"}
    var v2 uint = 0
    var v3 int8 = 1
    var v4 uint = v2
    v5 := make(map[int]uint32)
    var v6 uint64 = 117
    v20 := GenS1{F0: 32767}
    if (65535 < 2147483647) {
        if true {
            v3 += (v3 << ((668) & 63))
            fmt.Println(v3, v4, v4)
        } else {
            v21 := GenS2{F0: 118}
            _ = v21
        }
        v4 = uint(v6)
    } else {
        v22 := GenS3{F0: (!true)}
        v23 := GenS1{F0: 1000}
        _ = v22
        _ = v23
    }
    v24 := GenS4{F0: 65535, F1: 703, F2: (v6 / (((v6 - 161)) | 1))}
    for fi25 := 0; fi25 < 10; fi25++ {
        v26 := GenS5{F0: 0, F1: 0, F2: (fi25 % ((fi25) | 1))}
        var v27 string = "\n\t"
        _ = fi25
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 11; fi28++ {
        for fi29 := 0; fi29 < 3; fi29++ {
            var v30 float32 = 1.5
            v4 -= 69
            _ = fi29
            _ = v30
        }
        if (662 != 42) {
            var v31 int32 = -2147483648
            _ = v31
        } else {
            v32 := GenS5{F0: fi28, F1: 0, F2: (fi28 ^ fi28)}
            _ = len(v5)
            _ = v32
        }
        _ = fi28
    }
    if false {
        switch -9223372036854775808 {
        case 967:
            fmt.Println(v2, v4)
        default:
            v6 = 2
        }
        var v33 int16 = 686
        var v34 int8 = v3
        _ = v33
        _ = v34
    }
    switch 32768 {
    case -683:
        v4 += v4
    default:
        _ = len(v5)
    }
    var v35 float32 = -2.25
    var v36 float64 = ((-2.25 + 1.5) + float64(((0.1 + v35) / (v35 - v35))))
    if (0.0 <= float32(float64(0.1))) {
        for fi37 := 0; fi37 < 10; fi37++ {
            var v38 float64 = (((1e10 / float64(v35)) * (float64(v35) / v36)) - (v36 + v36))
            if false {
                var v39 uint32 = genFn0(fi37, uint32(v4), 118)
                delete(v5, int(v2))
                var v40 float32 = v35
                _ = v39
                _ = v40
            }
            _ = fi37
            _ = v38
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if (!(v6 != v6)) {
            fmt.Println(v2, v6, v6)
            fmt.Println(v35, v2, v3)
        }
    }
    if false {
        for fi41 := 0; fi41 < 13; fi41++ {
            for fi42 := 0; fi42 < 17; fi42++ {
                var v43 int64 = 42
                var v44 int = (^(fi42))
                var v45 uint64 = 616
                var v46 float32 = 0.0
                v47 := GenS5{F0: (v44 ^ fi42), F1: (v43 + (v43 << (((v2 * (v2 % (((^(v4))) | 1)))) & 63))), F2: -9223372036854775808}
                _ = fi42
                _ = v43
                _ = v44
                _ = v45
                _ = v46
                _ = v47
            }
            _ = fi41
        }
        fmt.Println(v6, v36)
    }
    for fi48 := 0; fi48 < 4; fi48++ {
        var v49 uint = v4
        var v50 uint8 = 128
        _ = fi48
        _ = v49
        _ = v50
    }
    for fi51 := 0; fi51 < 7; fi51++ {
        switch -1 {
        case 4294967295:
            if (!((false && false) && true)) {
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
            } else {
                var v52 uint64 = v6
                v5[-396] = 2147483647
                if (42 <= 1000) {
                    for fi53 := 0; fi53 < 8; fi53++ {
                        if (false && true) {
                            switch "foo bar" {
                            case "\n\t":
                                if (0 <= 1) {
                                    if len(v1) > 0 {
                                        v1[0] = "a"
                                    }
                                } else {
                                    switch (-(fi53)) {
                                    case 625:
                                        var v54 uint16 = 42
                                        _ = v54
                                    default:
                                        v55 := GenS5{F0: fi51, F1: 383, F2: fi51}
                                        _ = v55
                                    }
                                    var v56 float32 = 0.0
                                    var v57 int32 = 21
                                    _ = v56
                                    _ = v57
                                }
                                var v58 uint32 = 1
                                _ = v58
                            case "":
                                v36 -= v36
                            case "a":
                                v6 = v52
                                switch fi53 {
                                case -185:
                                    if false {
                                        v59 := GenS6{F0: (v3 % ((v3) | 1)), F1: (^(v6)), F2: v6}
                                        _ = v59
                                    }
                                case -779:
                                    v60 := GenS7{F0: 1, F1: 4294967295}
                                    _ = v60
                                case 127:
                                    fmt.Println(fi51, v36)
                                default:
                                    fmt.Println(v3, v52, v2)
                                }
                                v61 := GenS6{F0: 42, F1: (v6 >> ((42) & 63)), F2: (v52 - v52)}
                                _ = v61
                            default:
                                v62 := GenS8{F0: 32767, F1: -570, F2: 0}
                                _ = v62
                            }
                            switch 2147483647 {
                            case -787:
                                v6 -= (v52 - 106)
                            case 42:
                                fmt.Println(v6, v2)
                            case 32767:
                                if (v4 >= v2) {
                                    var v63 int64 = 42
                                    var v64 uint = 4294967295
                                    _ = v63
                                    _ = v64
                                } else {
                                    v6 = v52
                                }
                            }
                        } else {
                            switch "\n\t" {
                            case "foo bar":
                                var v65 uint16 = 864
                                _ = v65
                            case "\n\t":
                                v36 += v36
                            case "hello":
                                var v66 uint8 = 244
                                _ = v66
                            }
                            var v67 uint32 = 2
                            _ = v67
                        }
                        _ = fi53
                    }
                    if ((v2 - (v2 * v2)) == (v2 ^ (v2 % ((uint(v6)) | 1)))) {
                        fmt.Println(fi51, fi51)
                        for fi68 := 0; fi68 < 10; fi68++ {
                            v69 := GenS7{F0: 94, F1: 4294967295}
                            v70 := GenS6{F0: v3, F1: 2, F2: (v6 & 0)}
                            _ = fi68
                            _ = v69
                            _ = v70
                        }
                    } else {
                        var v71 uint64 = v52
                        _ = v71
                    }
                } else {
                    if false {
                        v5[(fi51 << ((v2) & 63))] = 4294967295
                        v2 = v2
                    }
                    var v72 uint64 = 18446744073709551615
                    _ = v72
                }
                _ = v52
            }
        case 179:
            var v73 uint = uint(v6)
            _ = v73
        case 128:
            v2 += (v4 * (v2 % ((v2) | 1)))
        default:
            var v74 float64 = v36
            _ = v74
        }
        _ = fi51
    }
    if false {
        switch "\n\t" {
        case "a":
            v35 = float32(v36)
        case "":
            for fi75 := 0; fi75 < 3; fi75++ {
                v76 := GenS6{F0: v3, F1: 42, F2: v6}
                var v77 uint8 = 179
                _ = fi75
                _ = v76
                _ = v77
            }
        }
        for fi78 := 0; fi78 < 19; fi78++ {
            var v79 string = "a"
            _ = fi78
            _ = v79
        }
    }
    v6 -= v6
    if true {
        for fi80 := 0; fi80 < 2; fi80++ {
            switch ("\n\t" + "a") {
            case "hello":
                switch (fi80 - 2147483647) {
                case 42:
                    v81 := GenS1{F0: -940}
                    _ = v81
                }
            case "foo bar":
                v82 := GenS9{F0: ((("hello" + "hello") + "") + ("a" + ((("hello" + "\n\t") + "\n\t") + ""))), F1: v2, F2: 0}
                _ = v82
            case "a":
                var v83 int16 = 0
                _ = v83
            }
            v84 := GenS8{F0: 302, F1: -868, F2: 1}
            var v85 uint16 = 65535
            _ = fi80
            _ = v84
            _ = v85
        }
        if (false || ((^(v6)) >= (v6 % ((42) | 1)))) {
            if true {
                var v86 uint16 = 65535
                var v87 string = ""
                for fi88 := 0; fi88 < 7; fi88++ {
                    if false {
                        var v89 bool = ((!true) && (!true))
                        if v89 {
                            if v89 {
                                if v89 {
                                    v86 += (v86 / ((v86) | 1))
                                } else {
                                    for fi90 := 0; fi90 < 11; fi90++ {
                                        v91 := GenS8{F0: v86, F1: -435, F2: 66}
                                        var v92 int64 = -103
                                        v3 = (v3 | (v3 ^ (v3 & (v3 & -128))))
                                        _ = fi90
                                        _ = v91
                                        _ = v92
                                    }
                                    fmt.Println(v87)
                                }
                                fmt.Println(v87)
                            } else {
                                if false {
                                    v93 := GenS2{F0: 127}
                                    if ((!true) && (!false)) {
                                        for fi94 := 0; fi94 < 20; fi94++ {
                                            if ((^(v3)) != v3) {
                                                fmt.Println(v87, v2, v36)
                                                v95 := GenS3{F0: ((fi94 & (fi94 + (-(fi88)))) == fi94)}
                                                v3 = (v3 % ((v3) | 1))
                                                var v96 int8 = v3
                                                if ((v87 + v87) >= (v87 + "hello")) {
                                                    var v97 uint32 = 65535
                                                    if (0 > v6) {
                                                        var v98 uint16 = (v86 / (((^(v86))) | 1))
                                                        _ = v98
                                                    }
                                                    switch v87 {
                                                    case "hello":
                                                        switch "foo bar" {
                                                        case "":
                                                            if ((v97 >> (((v4 + (v2 * (v4 & v2)))) & 63)) > genFn0(int(v3), v97, 101)) {
                                                                if (v4 <= uint(v86)) {
                                                                    v6 -= (v6 + v6)
                                                                    if v89 {
                                                                        if (32767 > fi88) {
                                                                            switch (fi88 | fi94) {
                                                                            case 4294967295:
                                                                                fmt.Println(v3)
                                                                            case 994:
                                                                                switch v87 {
                                                                                case "\n\t":
                                                                                    var v99 uint = (v4 ^ (^(v4)))
                                                                                    _ = v99
                                                                                case "foo bar":
                                                                                    var v100 int16 = 1000
                                                                                    _ = v100
                                                                                case "hello":
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = ("foo bar" + v87)
                                                                                    }
                                                                                }
                                                                                fmt.Println(fi88, v2)
                                                                                v101 := GenS5{F0: int(v86), F1: 133, F2: (fi94 / (((^(fi94))) | 1))}
                                                                                _ = v101
                                                                            case 32768:
                                                                                var v102 int16 = -103
                                                                                _ = v102
                                                                            default:
                                                                                v35 = 0.0
                                                                            }
                                                                            if v89 {
                                                                                v87 = ("\n\t" + (((v87 + "hello") + v87) + v87))
                                                                                var v103 uint8 = 255
                                                                                switch "\n\t" {
                                                                                case "hello":
                                                                                    switch fi88 {
                                                                                    case 1:
                                                                                        var v104 int8 = v96
                                                                                        _ = v104
                                                                                    case -1:
                                                                                        var v105 int8 = 127
                                                                                        _ = v105
                                                                                    default:
                                                                                        v86 = 32767
                                                                                    }
                                                                                case "\n\t":
                                                                                    fmt.Println(v6)
                                                                                case "foo bar":
                                                                                    v106 := v5[(fi88 & 42)]
                                                                                    _ = v106
                                                                                }
                                                                                _ = v103
                                                                            } else {
                                                                                switch int(v86) {
                                                                                case 811:
                                                                                    v3 = -110
                                                                                }
                                                                                v89 = (-32768 <= 1000)
                                                                            }
                                                                            v5[fi94] = genFn0(fi88, uint32(v2), 128)
                                                                            if (v89 || v89) {
                                                                                v3 = v96
                                                                                var v107 float64 = -2.25
                                                                                var v108 float32 = 3.14159
                                                                                _ = v107
                                                                                _ = v108
                                                                            } else {
                                                                                v109 := GenS10{F0: "a", F1: 34, F2: (v6 << ((255) & 63))}
                                                                                if (!v89) {
                                                                                    switch 32768 {
                                                                                    case 42:
                                                                                        if false {
                                                                                            _ = len(v1)
                                                                                        } else {
                                                                                            v110 := GenS1{F0: 1000}
                                                                                            v35 = 0.1
                                                                                            var v111 float64 = 1.5
                                                                                            _ = v110
                                                                                            _ = v111
                                                                                        }
                                                                                    }
                                                                                    v87 = v87
                                                                                }
                                                                                v1 = append(v1, ((("hello" + v87) + "a") + "hello"))
                                                                                if (!((v86 >> ((v2) & 63)) < v86)) {
                                                                                    if false {
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = "\n\t"
                                                                                        }
                                                                                        var v112 bool = ("foo bar" < ("\n\t" + (v87 + v87)))
                                                                                        if v112 {
                                                                                            _ = len(v5)
                                                                                            fmt.Println(v36, fi94, v36)
                                                                                        } else {
                                                                                            var v113 uint16 = 0
                                                                                            fmt.Println(v6, v6)
                                                                                            var v114 int32 = 42
                                                                                            var v115 int16 = 0
                                                                                            v86 = v86
                                                                                            _ = v113
                                                                                            _ = v114
                                                                                            _ = v115
                                                                                        }
                                                                                        _ = v112
                                                                                    } else {
                                                                                        var v116 int = fi94
                                                                                        switch v87 {
                                                                                        case "hello":
                                                                                            switch v87 {
                                                                                            case "\n\t":
                                                                                                fmt.Println(v97, v87)
                                                                                            case "foo bar":
                                                                                                fmt.Println(v87, fi88)
                                                                                            case "hello":
                                                                                                v116 -= v116
                                                                                            default:
                                                                                                var v117 float32 = v35
                                                                                                _ = v117
                                                                                            }
                                                                                        case "a":
                                                                                            if false {
                                                                                                if v89 {
                                                                                                    v118 := GenS9{F0: v87, F1: 169, F2: genFn0((^(fi94)), uint32(v86), 0)}
                                                                                                    v116 += (fi88 >> (((v4 / ((v4) | 1))) & 63))
                                                                                                    _ = v118
                                                                                                }
                                                                                            }
                                                                                        case "foo bar":
                                                                                            v119 := GenS1{F0: 42}
                                                                                            _ = v119
                                                                                        default:
                                                                                            v120 := GenS10{F0: (v87 + v87), F1: 255, F2: (v6 & v6)}
                                                                                            fmt.Println(fi94, v36, fi88)
                                                                                            fmt.Println(v35, v116, v35)
                                                                                            _ = v120
                                                                                        }
                                                                                        _ = v116
                                                                                    }
                                                                                    var v121 int8 = (-(v3))
                                                                                    fmt.Println(v121, v87, v87)
                                                                                    if false {
                                                                                        v36 += 1.5
                                                                                        if true {
                                                                                            v122 := GenS10{F0: v87, F1: 42, F2: v6}
                                                                                            v121 = (^(v96))
                                                                                            _ = v122
                                                                                        } else {
                                                                                            v97 = 425
                                                                                            var v123 uint64 = v6
                                                                                            _ = v123
                                                                                        }
                                                                                        v3 += v96
                                                                                    }
                                                                                    _ = v121
                                                                                }
                                                                                _ = v109
                                                                            }
                                                                        } else {
                                                                            v124 := GenS11{F0: (!(128 == 0)), F1: 0, F2: fi88}
                                                                            _ = v124
                                                                        }
                                                                    }
                                                                } else {
                                                                    var v125 uint = (v2 + v2)
                                                                    var v126 float32 = 1.5
                                                                    if true {
                                                                        var v127 float64 = float64(v35)
                                                                        _ = v127
                                                                    }
                                                                    _ = v125
                                                                    _ = v126
                                                                }
                                                            }
                                                        case "\n\t":
                                                            switch -791 {
                                                            case 638:
                                                                if false {
                                                                    v89 = (fi94 > fi88)
                                                                } else {
                                                                    delete(v5, (-(fi88)))
                                                                    fmt.Println(v35)
                                                                    v96 -= 0
                                                                }
                                                            case 127:
                                                                var v128 int = fi88
                                                                fmt.Println(v6)
                                                                var v129 uint32 = v97
                                                                _ = v128
                                                                _ = v129
                                                            case -2147483648:
                                                                fmt.Println(v35)
                                                            default:
                                                                delete(v5, fi94)
                                                            }
                                                        default:
                                                            var v130 float64 = v36
                                                            _ = v130
                                                        }
                                                    case "a":
                                                        v36 = v36
                                                    case "":
                                                        _ = len(v5)
                                                    }
                                                    _ = v97
                                                } else {
                                                    if (!(!(!false))) {
                                                        if (2 >= (v6 & v6)) {
                                                            var v131 uint = v2
                                                            var v132 bool = false
                                                            v133 := GenS4{F0: genFn0(32767, 2, 197), F1: (v86 % (((v86 | (v86 << ((v4) & 63)))) | 1)), F2: v6}
                                                            _ = v131
                                                            _ = v132
                                                            _ = v133
                                                        }
                                                        v134 := GenS3{F0: (!(934 < 0))}
                                                        if (v87 != "\n\t") {
                                                            var v135 uint8 = 0
                                                            v136 := v5[fi94]
                                                            v87 = v87
                                                            _ = v135
                                                            _ = v136
                                                        }
                                                        _ = v134
                                                    } else {
                                                        _ = len(v5)
                                                        if (-1 == -291) {
                                                            fmt.Println(v87)
                                                        } else {
                                                            v96 = v3
                                                            v89 = (v36 == (0.1 - v36))
                                                            var v137 int8 = v96
                                                            _ = v137
                                                        }
                                                        var v138 uint = (v2 % (((v4 + v2)) | 1))
                                                        _ = v138
                                                    }
                                                    fmt.Println(v86, v2)
                                                }
                                                _ = v95
                                                _ = v96
                                            }
                                            var v139 string = (v87 + v87)
                                            _ = fi94
                                            _ = v139
                                        }
                                        v5[fi88] = 65535
                                    }
                                    v89 = v89
                                    _ = v93
                                }
                                v2 = 2
                                v140 := GenS12{F0: 1, F1: 0, F2: 42}
                                _ = v140
                            }
                            v141 := GenS9{F0: v87, F1: (v4 + (v4 - v4)), F2: genFn0(32768, 4294967295, 89)}
                            _ = v141
                        }
                        v4 = uint(v86)
                        _ = v89
                    }
                    _ = fi88
                }
                _ = v86
                _ = v87
            }
            var v142 int16 = -68
            _ = v142
        }
    }
    v5[-73] = 2
    var v143 int8 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v143:", v143)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v143
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
