package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS9 struct { F0 int; F1 string }
type GenS10 struct { F0 int8; F1 uint64 }
type GenS11 struct { F0 bool; F1 uint32 }
type GenS12 struct { F0 int64; F1 int; F2 uint }
type GenS13 struct { F0 uint64; F1 int32 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 uint16; F1 int8; F2 int32 }
type GenS16 struct { F0 uint64 }
type GenS17 struct { F0 int64 }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 uint16 }
type GenS20 struct { F0 uint8 }
type GenS21 struct { F0 string; F1 string }
type GenS22 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS23 struct { F0 int64; F1 bool; F2 uint16 }
type GenS24 struct { F0 string; F1 string; F2 uint64 }
type GenS25 struct { F0 uint8; F1 string; F2 uint32 }
type GenS26 struct { F0 uint }
type GenS27 struct { F0 string; F1 int16 }
type GenS28 struct { F0 bool; F1 uint32; F2 uint16 }
func genFn0(p0 uint32, p1 uint16, p2 int64) int64 {
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: p2}
        _ = fi6
        _ = v7
    }
    p1 = p1
    if false {
        var v8 uint8 = 2
        _ = v8
    } else {
        if false {
            v9 := GenS2{F0: 42}
            v10 := []int{}
            v11 := GenS3{F0: (^(p0))}
            _ = v10
            _ = v11
            _ = v9
        }
        p1 -= 210
        v12 := GenS4{F0: 32767, F1: "a"}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    v1 := make(map[int]string)
    v2 := []string{"", "hello", "", ""}
    var v3 int = 128
    var v4 string = "hello"
    v5 := []int{128}
    for fi13 := 0; fi13 < 20; fi13++ {
        _ = len(v1)
        switch (v4 + "") {
        case "hello":
            for fi14 := 0; fi14 < 10; fi14++ {
                v15 := v1[fi13]
                delete(v1, (v3 - v3))
                var v16 int8 = 0
                _ = fi14
                _ = v15
                _ = v16
            }
        default:
            v3 -= (v3 / ((fi13) | 1))
        }
        v1[(^(v3))] = v4
        _ = fi13
    }
    fmt.Println(v4, v4)
    v4 = v4
    v17 := GenS4{F0: 0, F1: "\n\t"}
    switch v3 {
    case 408:
        for fi18 := 0; fi18 < 7; fi18++ {
            switch fi18 {
            case -9223372036854775808:
                v3 += (fi18 + 32767)
            }
            v4 = v4
            switch v4 {
            case "foo bar":
                v19 := GenS5{F0: 1}
                _ = v19
            default:
                v4 = "a"
            }
            _ = fi18
        }
    case -9223372036854775808:
        v20 := GenS6{F0: 2}
        _ = v20
    default:
        switch v4 {
        case "foo bar":
            if false {
                var v21 int32 = 1
                _ = v21
            } else {
                _ = len(v2)
                if true {
                    var v22 int16 = 0
                    _ = v22
                } else {
                    v23 := GenS7{F0: (42 == 9223372036854775807)}
                    v24 := GenS5{F0: 2147483647}
                    v3 -= v3
                    _ = v23
                    _ = v24
                }
                _ = len(v1)
            }
        default:
            v1[732] = ((("\n\t" + v4) + (("\n\t" + ("" + v4)) + (v4 + v4))) + v4)
        }
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        for fi26 := 0; fi26 < 14; fi26++ {
            if false {
                var v27 bool = false
                v4 = (((v4 + v4) + "a") + (((((v4 + "") + v4) + "a") + "hello") + v4))
                if false {
                    v27 = (int8(v3) <= 0)
                    var v28 int = (v3 % ((fi26) | 1))
                    _ = v28
                } else {
                    v3 -= fi26
                }
                if ((fi26 % ((42) | 1)) > -1) {
                    var v29 int64 = 9223372036854775807
                    var v30 uint = 65535
                    v4 = "\n\t"
                    _ = v29
                    _ = v30
                } else {
                    var v31 string = ((v4 + "foo bar") + v4)
                    _ = len(v2)
                    if true {
                        v3 -= (v3 & (^(fi26)))
                        v3 -= v3
                    } else {
                        switch 128 {
                        case -334:
                            v3 -= (fi26 & (v3 + (^(fi26))))
                        case -975:
                            var v32 float32 = (0.1 + 3.14159)
                            _ = v32
                        case 4294967295:
                            switch (v3 << ((4294967295) & 63)) {
                            case -1:
                                _ = len(v1)
                            }
                        default:
                            v33 := GenS8{F0: 127, F1: 2147483648, F2: 112}
                            _ = v33
                        }
                        v4 = "foo bar"
                        var v34 uint8 = 2
                        _ = v34
                    }
                    _ = v31
                }
                _ = v27
            }
            delete(v1, 715)
            v4 = "a"
            v35 := GenS2{F0: 127}
            _ = fi26
            _ = v35
        }
        switch v4 {
        case "\n\t":
            var v36 uint8 = 2
            _ = v36
        case "a":
            fmt.Println(v3, v3, v3)
        default:
            var v37 float64 = -2.25
            _ = v37
        }
        if (!false) {
            v3 = -416
        } else {
            if true {
                for fi38 := 0; fi38 < 14; fi38++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    fmt.Println(fi25, fi25, v3)
                    if ((true || true) || false) {
                        if false {
                            if true {
                                v4 = v4
                                v39 := GenS8{F0: 0, F1: 4294967295, F2: -1}
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v40 string = "hello"
                                v40 = "hello"
                                _ = v39
                                _ = v40
                            }
                            if (-2.25 <= (0.0 - 1.5)) {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                var v41 int64 = 1
                                _ = v41
                            }
                            v42 := v1[fi38]
                            v43 := GenS9{F0: (fi38 + (^(v3))), F1: "foo bar"}
                            v3 += 0
                            v4 = (v4 + v4)
                            if (-60 == -35) {
                                var v44 uint8 = 127
                                v4 = v4
                                var v45 string = v4
                                _ = v44
                                _ = v45
                            }
                            v2 = append(v2, v4)
                            if (v4 == (v42 + v4)) {
                                switch (-(v3)) {
                                case 32768:
                                    var v46 uint8 = 127
                                    if true {
                                        var v47 int64 = -72
                                        fmt.Println(v42, fi25, v3)
                                        v46 -= 128
                                        var v48 int = 128
                                        var v49 float64 = float64(1e10)
                                        _ = v47
                                        _ = v48
                                        _ = v49
                                    }
                                    _ = v46
                                case 9223372036854775807:
                                    var v50 uint = 794
                                    _ = v50
                                case 4294967295:
                                    switch fi25 {
                                    case 650:
                                        v2 = append(v2, ((((v4 + "a") + ((v4 + "hello") + v4)) + (v42 + "foo bar")) + "hello"))
                                    case -9223372036854775808:
                                        v4 = ("a" + "\n\t")
                                    }
                                default:
                                    if (128 != 0) {
                                        switch ("foo bar" + (v4 + (v4 + "hello"))) {
                                        case "foo bar":
                                            v3 = (fi38 % ((fi25) | 1))
                                            if true {
                                                fmt.Println(fi38)
                                            }
                                        }
                                    } else {
                                        switch 128 {
                                        case -458:
                                            v51 := GenS10{F0: 0, F1: 0}
                                            _ = v51
                                        case 42:
                                            v52 := GenS8{F0: 107, F1: 65535, F2: -128}
                                            _ = v52
                                        case 128:
                                            v53 := GenS11{F0: (127 < 42), F1: 42}
                                            _ = v53
                                        default:
                                            v1[fi25] = v42
                                        }
                                        fmt.Println(v3)
                                        switch -398 {
                                        case 2147483647:
                                            var v54 uint64 = 0
                                            _ = v54
                                        }
                                    }
                                }
                            }
                            if len(v5) > 0 {
                                v5[0] = (fi38 << ((1) & 63))
                            }
                            _ = v42
                            _ = v43
                        }
                        if false {
                            fmt.Println(v3, fi38, v3)
                            delete(v1, fi38)
                        }
                    } else {
                        var v55 int = (v3 - v3)
                        var v56 float32 = 3.14159
                        switch (^(fi38)) {
                        case 280:
                            v57 := v1[fi38]
                            _ = v57
                        }
                        if false {
                            fmt.Println(fi38)
                            v55 = (^(v55))
                            v1[fi38] = v4
                        }
                        _ = v55
                        _ = v56
                    }
                    _ = fi38
                }
            } else {
                v58 := GenS12{F0: 9223372036854775807, F1: (v3 | (v3 * 39)), F2: 1}
                if true {
                    if false {
                        for fi59 := 0; fi59 < 4; fi59++ {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = fi59
                        }
                        delete(v1, v3)
                        var v60 uint32 = 42
                        _ = v60
                    }
                    var v61 uint = 255
                    v62 := GenS13{F0: 380, F1: -2147483648}
                    _ = v61
                    _ = v62
                } else {
                    v3 = 0
                }
                _ = v58
            }
            v63 := v1[32767]
            if (-1 != -1) {
                if false {
                    for fi64 := 0; fi64 < 7; fi64++ {
                        v3 = 81
                        var v65 uint32 = 305
                        fmt.Println(v63, v63, v4)
                        v66 := GenS8{F0: 63, F1: v65, F2: -22}
                        _ = fi64
                        _ = v65
                        _ = v66
                    }
                    switch v4 {
                    case "":
                        var v67 float64 = (0.0 - 3.14159)
                        _ = v67
                    }
                    if (!true) {
                        v63 = ""
                        var v68 uint16 = 65535
                        v3 += 429
                        _ = v68
                    } else {
                        for fi69 := 0; fi69 < 19; fi69++ {
                            fmt.Println(fi69)
                            v3 = (fi69 % ((42) | 1))
                            var v70 bool = (false && true)
                            fmt.Println(v63, v63, v70)
                            var v71 int = 127
                            if (66 != 72) {
                                v3 += fi69
                                var v72 float64 = 1.5
                                var v73 int64 = 0
                                _ = v72
                                _ = v73
                            } else {
                                var v74 bool = (v4 <= v63)
                                var v75 uint64 = 1
                                var v76 uint8 = 124
                                _ = v74
                                _ = v75
                                _ = v76
                            }
                            _ = fi69
                            _ = v70
                            _ = v71
                        }
                        v63 = (v4 + "hello")
                    }
                }
                if ((1 > 2147483647) || false) {
                    for fi77 := 0; fi77 < 15; fi77++ {
                        var v78 int16 = 373
                        if true {
                            fmt.Println(fi25, fi77)
                            var v79 string = "\n\t"
                            if true {
                                v63 = ("\n\t" + v4)
                                if false {
                                    var v80 int64 = genFn0(2147483648, 0, 604)
                                    _ = v80
                                } else {
                                    var v81 uint64 = 1
                                    v63 = v79
                                    _ = v81
                                }
                                v63 = v63
                            }
                            _ = v79
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = len(v1)
                        }
                        if (1 == genFn0(885, 255, -1)) {
                            v82 := GenS14{F0: 0}
                            _ = v82
                        } else {
                            var v83 string = "a"
                            var v84 uint16 = 759
                            _ = v83
                            _ = v84
                        }
                        if (!(127 > 0)) {
                            if true {
                                var v85 int = fi25
                                switch (-(fi77)) {
                                case 128:
                                    v85 = (fi77 << ((1) & 63))
                                case -618:
                                    v4 = ("a" + "")
                                case -9223372036854775808:
                                    var v86 uint = 2
                                    _ = v86
                                }
                                v87 := GenS8{F0: 1, F1: 382, F2: 100}
                                _ = v85
                                _ = v87
                            } else {
                                v88 := GenS13{F0: 686, F1: 454}
                                _ = v88
                            }
                            var v89 int16 = 1000
                            _ = v89
                        } else {
                            if (false || false) {
                                v5 = append(v5, (v3 | fi25))
                                v90 := GenS15{F0: 0, F1: 0, F2: 908}
                                _ = v90
                            } else {
                                v91 := GenS16{F0: 1}
                                if (!true) {
                                    _ = len(v1)
                                    v92 := GenS14{F0: 4294967295}
                                    if len(v2) > 0 {
                                        v2[0] = (v63 + "hello")
                                    }
                                    var v93 int32 = -2147483648
                                    _ = v92
                                    _ = v93
                                } else {
                                    v63 = v4
                                }
                                _ = v91
                            }
                            v94 := v1[(fi77 & (fi77 * fi25))]
                            _ = v94
                        }
                        v3 -= fi77
                        _ = fi77
                        _ = v78
                    }
                    switch v63 {
                    case "":
                        if false {
                            var v95 uint32 = 930
                            v4 = (("foo bar" + (("\n\t" + (("foo bar" + ((((v4 + v4) + "foo bar") + (v63 + "hello")) + v63)) + v4)) + (((("" + "foo bar") + (v63 + v4)) + "\n\t") + v63))) + "a")
                            _ = v95
                        } else {
                            v3 = 981
                        }
                    case "a":
                        if true {
                            var v96 bool = false
                            switch v4 {
                            case "":
                                for fi97 := 0; fi97 < 5; fi97++ {
                                    v3 = -543
                                    _ = fi97
                                }
                            }
                            _ = v96
                        } else {
                            var v98 string = v4
                            for fi99 := 0; fi99 < 11; fi99++ {
                                switch 127 {
                                case 4294967295:
                                    v4 = "foo bar"
                                    var v100 uint32 = 2
                                    _ = v100
                                }
                                v98 = ("a" + "")
                                if (!false) {
                                    fmt.Println(v63, v4)
                                    if (!false) {
                                        v101 := GenS17{F0: -1}
                                        _ = v101
                                    } else {
                                        v63 = "foo bar"
                                    }
                                    v3 = fi99
                                } else {
                                    fmt.Println(v3)
                                }
                                _ = fi99
                            }
                            _ = v98
                        }
                    default:
                        for fi102 := 0; fi102 < 14; fi102++ {
                            var v103 uint32 = 1
                            if true {
                                if (v103 <= (v103 % ((v103) | 1))) {
                                    v63 = (((v63 + ("" + "a")) + v63) + v4)
                                }
                                var v104 int32 = 0
                                _ = v104
                            }
                            v103 = (v103 % ((42) | 1))
                            v105 := GenS5{F0: -1}
                            if ((42 < 2) && (!(!(!false)))) {
                                var v106 int16 = -1
                                if len(v2) > 0 {
                                    v2[0] = "a"
                                }
                                v107 := GenS9{F0: -1, F1: v63}
                                _ = v106
                                _ = v107
                            }
                            _ = fi102
                            _ = v103
                            _ = v105
                        }
                    }
                }
                v108 := GenS18{F0: true}
                _ = v108
            }
            _ = v63
        }
        _ = fi25
    }
    if len(v5) > 0 {
        v5[0] = 32768
    }
    var v109 uint32 = 455
    if true {
        fmt.Println(v109, v109)
    }
    switch ((v4 + (v4 + v4)) + v4) {
    case "\n\t":
        for fi110 := 0; fi110 < 19; fi110++ {
            if false {
                var v111 string = v4
                _ = v111
            } else {
                switch v4 {
                case "hello":
                    v112 := GenS19{F0: 588}
                    _ = v112
                case "a":
                    fmt.Println(v109)
                case "foo bar":
                    for fi113 := 0; fi113 < 9; fi113++ {
                        delete(v1, (fi110 / (((fi113 & -595)) | 1)))
                        switch "foo bar" {
                        case "hello":
                            v4 = v4
                        case "":
                            v109 += v109
                        case "a":
                            v114 := GenS8{F0: 1, F1: (^(v109)), F2: 42}
                            if (v4 >= ("a" + "\n\t")) {
                                var v115 int32 = 42
                                var v116 int8 = 1
                                v5 = append(v5, (fi110 ^ 2147483647))
                                _ = v115
                                _ = v116
                            } else {
                                if (42 > 4294967295) {
                                    if (false || false) {
                                        v109 += (v109 & v109)
                                    }
                                    var v117 int = 2147483647
                                    _ = v117
                                }
                                v109 += (v109 / (((v109 / ((v109) | 1))) | 1))
                            }
                            _ = v114
                        }
                        if false {
                            var v118 uint = 65535
                            _ = v118
                        }
                        _ = fi113
                    }
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                var v119 float32 = 1.5
                var v120 int8 = 0
                _ = v119
                _ = v120
            }
            v121 := v1[fi110]
            _ = fi110
            _ = v121
        }
    case "a":
        var v122 uint32 = 519
        _ = v122
    case "":
        if (42 < genFn0(v109, 2, genFn0(42, 418, 753))) {
            var v123 float64 = (float64(-2.25) - 3.14159)
            for fi124 := 0; fi124 < 12; fi124++ {
                var v125 uint16 = 42
                for fi126 := 0; fi126 < 10; fi126++ {
                    if ((!(2 != uint64(v109))) || (-456 > 0)) {
                        v127 := GenS11{F0: (true || ((v123 >= v123) && true)), F1: v109}
                        _ = v127
                    }
                    v128 := GenS20{F0: 2}
                    fmt.Println(fi124)
                    _ = fi126
                    _ = v128
                }
                for fi129 := 0; fi129 < 3; fi129++ {
                    v109 -= v109
                    v125 += (^(v125))
                    switch "foo bar" {
                    case "foo bar":
                        fmt.Println(v4, fi129, v123)
                    case "a":
                        if (!(0.0 > -2.25)) {
                            if (!(1 == 2)) {
                                var v130 float64 = 0.0
                                v1[fi129] = v4
                                var v131 uint16 = v125
                                _ = v130
                                _ = v131
                            } else {
                                var v132 uint8 = 2
                                var v133 int = (fi124 >> ((255) & 63))
                                _ = v132
                                _ = v133
                            }
                            var v134 uint32 = 0
                            if false {
                                switch ((v4 + "foo bar") + "hello") {
                                case "foo bar":
                                    var v135 uint64 = 463
                                    _ = v135
                                case "":
                                    v109 = v109
                                }
                                switch 4294967295 {
                                case 4294967295:
                                    v109 += (v134 & (v134 + v134))
                                }
                            } else {
                                switch -9223372036854775808 {
                                case -304:
                                    v136 := GenS21{F0: v4, F1: v4}
                                    _ = v136
                                case -450:
                                    switch fi129 {
                                    case 252:
                                        var v137 uint64 = 0
                                        _ = v137
                                    }
                                default:
                                    var v138 int64 = -1
                                    _ = v138
                                }
                            }
                            _ = v134
                        } else {
                            fmt.Println(v4, fi129, fi129)
                            fmt.Println(fi124, v4, v125)
                            v139 := GenS22{F0: (^(v125)), F1: 2, F2: 4294967295}
                            _ = v139
                        }
                    }
                    v140 := GenS19{F0: 2}
                    _ = fi129
                    _ = v140
                }
                _ = fi124
                _ = v125
            }
            _ = v123
        } else {
            for fi141 := 0; fi141 < 4; fi141++ {
                v4 = ""
                var v142 int = (v3 & -2147483648)
                _ = fi141
                _ = v142
            }
            if len(v5) > 0 {
                v5[0] = (v3 << ((4294967295) & 63))
            }
            switch v3 {
            case -9223372036854775808:
                v3 = -920
            case 127:
                var v143 bool = false
                _ = v143
            case 42:
                var v144 uint32 = (^(v109))
                _ = v144
            }
        }
    default:
        fmt.Println(v3, v3)
    }
    if (-734 > 32767) {
        for fi145 := 0; fi145 < 4; fi145++ {
            v146 := GenS13{F0: 1, F1: 1}
            v147 := GenS23{F0: -90, F1: (genFn0(v109, 358, -9223372036854775808) > -9223372036854775808), F2: 255}
            for fi148 := 0; fi148 < 15; fi148++ {
                switch "foo bar" {
                case "foo bar":
                    if false {
                        v149 := GenS11{F0: ((2 == 255) || true), F1: 872}
                        v150 := GenS24{F0: (v4 + (("foo bar" + (("a" + ("foo bar" + "\n\t")) + v4)) + (v4 + v4))), F1: v4, F2: 182}
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v149
                        _ = v150
                    } else {
                        v3 = fi148
                        var v151 int = -457
                        if (-650 <= -2147483648) {
                            fmt.Println(v4)
                            var v152 uint = 0
                            var v153 int32 = -125
                            _ = v152
                            _ = v153
                        }
                        var v154 int32 = 1
                        v4 = "\n\t"
                        _ = v151
                        _ = v154
                    }
                case "a":
                    v155 := GenS6{F0: 127}
                    _ = v155
                default:
                    var v156 bool = (true && (798 <= v109))
                    _ = v156
                }
                if true {
                    fmt.Println(v4, fi148)
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v157 int = 128
                    _ = v157
                } else {
                    var v158 float64 = -2.25
                    var v159 int16 = 922
                    _ = v158
                    _ = v159
                }
                _ = len(v2)
                _ = fi148
            }
            _ = fi145
            _ = v146
            _ = v147
        }
        var v160 uint8 = 46
        _ = v160
    } else {
        switch (v3 + (v3 / ((v3) | 1))) {
        case -2147483648:
            switch (v3 * -9223372036854775808) {
            case 42:
                v161 := GenS11{F0: (1000 > 1), F1: (v109 % ((65535) | 1))}
                _ = v161
            case -977:
                var v162 uint32 = v109
                _ = v162
            default:
                for fi163 := 0; fi163 < 14; fi163++ {
                    v164 := GenS25{F0: 42, F1: v4, F2: (v109 >> ((215) & 63))}
                    _ = fi163
                    _ = v164
                }
            }
        case 42:
            if (!(!(42 == 0))) {
                v3 = 42
                for fi165 := 0; fi165 < 2; fi165++ {
                    var v166 int8 = 127
                    var v167 uint32 = (v109 % ((1) | 1))
                    var v168 int = fi165
                    _ = fi165
                    _ = v166
                    _ = v167
                    _ = v168
                }
                switch "a" {
                case "foo bar":
                    switch (v3 << ((2) & 63)) {
                    case -2147483648:
                        var v169 uint64 = 18446744073709551615
                        _ = v169
                    case 476:
                        fmt.Println(v3, v3)
                    default:
                        fmt.Println(v3)
                    }
                }
            }
        case 4294967295:
            fmt.Println(v3)
        }
        if false {
            fmt.Println(v109)
            for fi170 := 0; fi170 < 16; fi170++ {
                v171 := GenS26{F0: 999}
                v4 = v4
                _ = fi170
                _ = v171
            }
            if (!(-49 >= 100)) {
                v3 -= v3
                v109 -= (v109 & v109)
                for fi172 := 0; fi172 < 4; fi172++ {
                    delete(v1, (^(fi172)))
                    v109 = v109
                    for fi173 := 0; fi173 < 20; fi173++ {
                        var v174 int64 = 263
                        var v175 float64 = ((-2.25 + (1.5 * float64(0.0))) * 1.5)
                        _ = fi173
                        _ = v174
                        _ = v175
                    }
                    _ = fi172
                }
            }
        } else {
            if (65535 > 19) {
                var v176 string = v4
                var v177 int64 = 910
                for fi178 := 0; fi178 < 18; fi178++ {
                    v179 := GenS8{F0: 2, F1: (^(v109)), F2: 0}
                    _ = fi178
                    _ = v179
                }
                _ = v176
                _ = v177
            } else {
                v180 := GenS27{F0: v4, F1: 1}
                _ = v180
            }
            if true {
                var v181 uint16 = 65535
                if (false && (100 == -128)) {
                    var v182 bool = true
                    var v183 uint32 = v109
                    v182 = (!((^(v3)) > 130))
                    _ = v182
                    _ = v183
                }
                _ = v181
            }
        }
        v4 = "hello"
    }
    v3 = (v3 / ((v3) | 1))
    var v184 uint = 65535
    var v185 uint32 = (^(v109))
    _ = len(v2)
    switch (v3 % (((v3 ^ -474)) | 1)) {
    case -296:
        fmt.Println(v3, v3, v184)
    case 394:
        v185 -= (v109 / (((v109 / ((v109) | 1))) | 1))
    }
    v186 := GenS28{F0: (float32(float64(float32(((((-2.25 + ((1.5 / 1e10) * ((1.5 + 0.1) - (float64((((0.0 * 1.5) / 1.0) - 0.0)) * ((-2.25 * 3.14159) * (1.5 * ((((3.14159 * 1e10) + (float64(0.0) / -2.25)) * 1e10) * float64(0.0)))))))) / 1e10) - 1.5) * 3.14159)))) == 3.14159), F1: 0, F2: 65535}
    for fi187 := 0; fi187 < 7; fi187++ {
        fmt.Println(v184)
        var v188 uint8 = 89
        _ = fi187
        _ = v188
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v109:", v109)
    fmt.Println("v184:", v184)
    fmt.Println("v185:", v185)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v109
    _ = v17
    _ = v184
    _ = v185
    _ = v186
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
