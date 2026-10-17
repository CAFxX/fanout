package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int8 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 int; F2 int16 }
type GenS5 struct { F0 int16; F1 int32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 string; F1 uint8; F2 uint8 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int64; F1 int64 }
type GenS10 struct { F0 uint16; F1 int }
type GenS11 struct { F0 bool; F1 uint; F2 uint8 }
type GenS12 struct { F0 int32; F1 uint64 }
func genFn0(p0 uint64, p1 uint8, p2 string) int64 {
    p1 += p1
    for fi4 := 0; fi4 < 13; fi4++ {
        var v5 string = "hello"
        for fi6 := 0; fi6 < 4; fi6++ {
            if (!true) {
                var v7 int64 = 42
                switch (fi4 % ((fi4) | 1)) {
                case -9223372036854775808:
                    fmt.Println(p0, fi4)
                case 127:
                    var v8 uint8 = (p1 << ((65535) & 63))
                    _ = v8
                case 128:
                    if false {
                        var v9 uint64 = (p0 ^ (^(p0)))
                        v10 := []int{-1, fi6, (fi6 - fi4)}
                        _ = v10
                        _ = v9
                    } else {
                        var v11 bool = true
                        v12 := GenS1{F0: p2, F1: v7, F2: 100}
                        fmt.Println(p2)
                        v13 := []int{}
                        v14 := GenS1{F0: "hello", F1: v7, F2: 42}
                        _ = v11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                }
                _ = v7
            } else {
                switch (fi4 / ((-1) | 1)) {
                case 32767:
                    var v15 int8 = 122
                    _ = v15
                }
            }
            _ = fi6
        }
        v16 := GenS1{F0: (p2 + ""), F1: 1000, F2: 1}
        _ = fi4
        _ = v16
        _ = v5
    }
    if (uint16(p0) >= 0) {
        if true {
            v17 := GenS2{F0: -108, F1: 1}
            _ = v17
        } else {
            if false {
                p2 = "foo bar"
                var v18 uint8 = 213
                for fi19 := 0; fi19 < 5; fi19++ {
                    p0 += 0
                    var v20 float32 = ((1.5 * -2.25) / (1.5 - (-2.25 * -2.25)))
                    _ = fi19
                    _ = v20
                }
                _ = v18
            }
        }
        p2 = ""
    } else {
        for fi21 := 0; fi21 < 20; fi21++ {
            if (3.14159 != 1e10) {
                v22 := GenS1{F0: p2, F1: -524, F2: 0}
                _ = v22
            }
            _ = fi21
        }
        p0 = 42
    }
    v23 := GenS3{F0: 516}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return -379
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 127
    v3 := []uint32{0, 2}
    if false {
        var v24 string = "hello"
        _ = v24
    }
    var v25 float32 = (0.0 / 1e10)
    switch 128 {
    case -835:
        v2 = (v2 + (v2 & v2))
    case -9223372036854775808:
        var v26 uint32 = 353
        _ = v26
    }
    var v27 uint8 = (^(v2))
    v28 := GenS2{F0: 42, F1: 42}
    if ((true || (0 > 4294967295)) && false) {
        switch -2147483648 {
        case -9223372036854775808:
            for fi29 := 0; fi29 < 2; fi29++ {
                if true {
                    v3 = append(v3, 4294967295)
                    var v30 int8 = -114
                    var v31 uint8 = (v27 * (v27 << ((646) & 63)))
                    _ = v30
                    _ = v31
                }
                if (42 < 1) {
                    var v32 int16 = 0
                    switch (("a" + "\n\t") + "") {
                    case "":
                        v27 -= (v27 >> ((0) & 63))
                    case "hello":
                        switch ("foo bar" + ("hello" + ("\n\t" + "a"))) {
                        case "a":
                            var v33 float64 = (3.14159 - (3.14159 / 1.5))
                            _ = v33
                        default:
                            var v34 int8 = 127
                            _ = v34
                        }
                    case "foo bar":
                        for fi35 := 0; fi35 < 14; fi35++ {
                            v36 := v1[42]
                            if (true && true) {
                                v36 = fi35
                                switch (v36 - fi35) {
                                case 238:
                                    var v37 uint64 = 551
                                    _ = v37
                                case -834:
                                    switch "" {
                                    case "foo bar":
                                        if false {
                                            v27 = 1
                                            delete(v1, 779)
                                            switch (fi29 + 210) {
                                            case 42:
                                                if false {
                                                    v38 := GenS3{F0: 0}
                                                    _ = v38
                                                } else {
                                                    var v39 uint64 = 29
                                                    _ = v39
                                                }
                                            case -1:
                                                v40 := GenS3{F0: 4294967295}
                                                _ = v40
                                            case 571:
                                                if true {
                                                    switch (v36 & fi29) {
                                                    case 4294967295:
                                                        v32 += v32
                                                    }
                                                    v41 := GenS4{F0: 1, F1: -152, F2: v32}
                                                    if (4294967295 != 4294967295) {
                                                        v42 := GenS1{F0: "foo bar", F1: genFn0(834, (^(v27)), ("" + ("foo bar" + ("a" + (("hello" + ("hello" + ("a" + ""))) + "hello"))))), F2: 28}
                                                        v32 = (v32 | v32)
                                                        _ = v42
                                                    }
                                                    _ = v41
                                                }
                                            default:
                                                var v43 float32 = (((v25 - ((0.0 / v25) / (3.14159 / 3.14159))) - float32((float64(-2.25) - (-2.25 * 1e10)))) + (v25 / (v25 + 0.1)))
                                                _ = v43
                                            }
                                        }
                                    }
                                default:
                                    if false {
                                        v36 = v36
                                        v32 = -356
                                        var v44 uint64 = 0
                                        _ = v44
                                    }
                                }
                            }
                            v45 := GenS5{F0: (v32 / ((0) | 1)), F1: 1}
                            v27 = v2
                            _ = fi35
                            _ = v36
                            _ = v45
                        }
                    }
                    switch "" {
                    case "hello":
                        if (!(false || false)) {
                            var v46 float64 = 0.0
                            _ = v46
                        } else {
                            var v47 int8 = 127
                            v3 = append(v3, 4294967295)
                            _ = v47
                        }
                    case "a":
                        v3 = append(v3, 0)
                    case "":
                        v1[2147483648] = fi29
                    }
                    _ = v32
                }
                v2 = v2
                _ = fi29
            }
        case 135:
            if (127 <= 2147483647) {
                _ = len(v1)
                v48 := GenS1{F0: "", F1: -258, F2: -11}
                _ = v48
            } else {
                for fi49 := 0; fi49 < 5; fi49++ {
                    if false {
                        if false {
                            var v50 uint32 = 4294967295
                            _ = v50
                        }
                        _ = len(v1)
                        if true {
                            var v51 uint32 = 1
                            for fi52 := 0; fi52 < 3; fi52++ {
                                if (441 >= 1) {
                                    if (127 != 1) {
                                        var v53 int32 = 302
                                        var v54 int64 = genFn0(652, v27, "")
                                        fmt.Println(v51)
                                        _ = v53
                                        _ = v54
                                    } else {
                                        v55 := GenS4{F0: -128, F1: (-(fi49)), F2: -939}
                                        _ = v55
                                    }
                                    v27 = (v2 - 42)
                                    _ = len(v3)
                                }
                                _ = fi52
                            }
                            _ = v51
                        }
                    } else {
                        delete(v1, 152)
                        v25 += v25
                    }
                    delete(v1, 702)
                    if false {
                        for fi56 := 0; fi56 < 7; fi56++ {
                            switch (fi56 ^ -9223372036854775808) {
                            case -169:
                                if false {
                                    var v57 bool = true
                                    if ((1 <= -1) || false) {
                                        if false {
                                            if (-544 >= 9223372036854775807) {
                                                if true {
                                                    v58 := GenS2{F0: int8(v2), F1: 42}
                                                    fmt.Println(v27, v25, v25)
                                                    _ = v58
                                                } else {
                                                    fmt.Println(v2, v57)
                                                    var v59 string = "\n\t"
                                                    _ = v59
                                                }
                                            }
                                            if (fi56 < fi49) {
                                                var v60 uint32 = 65535
                                                v61 := GenS4{F0: -58, F1: (fi56 % ((-631) | 1)), F2: -32768}
                                                if (v57 && (false || (fi56 <= fi56))) {
                                                    v60 = (v60 & (v60 << ((42) & 63)))
                                                    _ = len(v1)
                                                } else {
                                                    v27 -= (v27 % ((42) | 1))
                                                }
                                                _ = v60
                                                _ = v61
                                            } else {
                                                var v62 int = fi56
                                                _ = v62
                                            }
                                            v63 := GenS3{F0: 2147483648}
                                            var v64 uint64 = uint64(v27)
                                            var v65 uint16 = 65535
                                            fmt.Println(fi49)
                                            if (v57 || ((v64 % ((2) | 1)) < (^(v64)))) {
                                                var v66 int = 316
                                                fmt.Println(v64)
                                                _ = v66
                                            }
                                            _ = v63
                                            _ = v64
                                            _ = v65
                                        } else {
                                            var v67 int = fi49
                                            if true {
                                                v68 := v1[613]
                                                var v69 uint16 = 72
                                                _ = v68
                                                _ = v69
                                            } else {
                                                v70 := GenS6{F0: 942}
                                                _ = v70
                                            }
                                            _ = v67
                                        }
                                    } else {
                                        var v71 uint = 1
                                        v2 -= 226
                                        _ = v71
                                    }
                                    _ = v57
                                } else {
                                    var v72 uint16 = 32767
                                    v27 = (v2 ^ 42)
                                    _ = v72
                                }
                            case 4294967295:
                                v2 = (v2 | (v2 | (^(v27))))
                                var v73 uint16 = 2
                                _ = v73
                            default:
                                v25 = 1.5
                            }
                            var v74 float32 = v25
                            _ = fi56
                            _ = v74
                        }
                        if false {
                            v75 := GenS7{F0: "\n\t", F1: v2, F2: (v27 / ((2) | 1))}
                            switch (-(fi49)) {
                            case -340:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            default:
                                if (0.0 <= 1e10) {
                                    var v76 uint32 = uint32(v2)
                                    _ = v76
                                }
                            }
                            fmt.Println(fi49, v27)
                            _ = v75
                        }
                        var v77 uint8 = 1
                        _ = v77
                    } else {
                        for fi78 := 0; fi78 < 4; fi78++ {
                            if true {
                                v79 := GenS8{F0: 2}
                                if false {
                                    switch "a" {
                                    case "\n\t":
                                        if false {
                                            switch (fi78 ^ (fi78 * 42)) {
                                            case -2147483648:
                                                v3 = append(v3, 2147483648)
                                            case 32767:
                                                v27 -= 157
                                            case 2147483647:
                                                if len(v3) > 0 {
                                                    v3[0] = 2147483647
                                                }
                                            default:
                                                var v80 int8 = 42
                                                _ = v80
                                            }
                                        }
                                        fmt.Println(v25)
                                        var v81 float32 = v25
                                        _ = v81
                                    case "foo bar":
                                        if (!(9223372036854775807 <= -1)) {
                                            v82 := GenS9{F0: -1, F1: 118}
                                            _ = v82
                                        }
                                    case "hello":
                                        var v83 uint32 = 4294967295
                                        _ = v83
                                    default:
                                        v25 = ((v25 + v25) * 0.1)
                                    }
                                    v84 := GenS10{F0: 32767, F1: fi78}
                                    switch (("\n\t" + "") + ("hello" + "foo bar")) {
                                    case "a":
                                        fmt.Println(v2, v25, v2)
                                    case "foo bar":
                                        v85 := GenS2{F0: 109, F1: 0}
                                        _ = v85
                                    }
                                    _ = v84
                                }
                                _ = v79
                            }
                            _ = fi78
                        }
                        v3 = append(v3, 2147483648)
                        v25 = ((v25 + 1e10) - 0.0)
                    }
                    _ = fi49
                }
            }
        default:
            v25 += v25
        }
        v3 = append(v3, 0)
    }
    switch "foo bar" {
    case "a":
        for fi86 := 0; fi86 < 20; fi86++ {
            fmt.Println(v2)
            _ = fi86
        }
    case "":
        var v87 uint32 = 4294967295
        _ = v87
    case "hello":
        if true {
            for fi88 := 0; fi88 < 14; fi88++ {
                v2 += v27
                v89 := GenS1{F0: ("a" + "foo bar"), F1: -730, F2: int8(fi88)}
                _ = fi88
                _ = v89
            }
            switch 4294967295 {
            case -9223372036854775808:
                var v90 float32 = ((-2.25 + v25) * v25)
                _ = v90
            case -1:
                var v91 int8 = 0
                _ = v91
            }
        } else {
            if ((42 == 87) && (0 > -32768)) {
                for fi92 := 0; fi92 < 4; fi92++ {
                    if ((42 == 0) || (0 >= 0)) {
                        switch 32768 {
                        case 42:
                            var v93 int64 = 0
                            _ = v93
                        case -1:
                            v1[1] = fi92
                        default:
                            v25 = v25
                        }
                        v2 += 96
                        v94 := GenS10{F0: 255, F1: 562}
                        _ = v94
                    }
                    for fi95 := 0; fi95 < 19; fi95++ {
                        if false {
                            if false {
                                switch (("hello" + "foo bar") + "a") {
                                case "hello":
                                    var v96 uint32 = 2147483648
                                    _ = v96
                                case "":
                                    v2 += 114
                                case "foo bar":
                                    var v97 uint = 1
                                    _ = v97
                                }
                                fmt.Println(v27, v25)
                                _ = len(v3)
                                v98 := GenS8{F0: 18446744073709551615}
                                v99 := GenS2{F0: 1, F1: 42}
                                _ = v98
                                _ = v99
                            }
                            v2 -= v2
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v100 bool = false
                            var v101 int16 = -1
                            _ = v100
                            _ = v101
                        } else {
                            fmt.Println(v25)
                        }
                        _ = fi95
                    }
                    v3 = append(v3, 2)
                    _ = fi92
                }
                var v102 int32 = 2147483647
                if false {
                    delete(v1, 1)
                    v27 += (v27 % ((v27) | 1))
                    v103 := GenS2{F0: 0, F1: 18446744073709551615}
                    _ = v103
                }
                _ = v102
            }
        }
    }
    v2 = (v2 / (((^(v27))) | 1))
    var v104 bool = ((true && (!(0.0 <= 0.0))) && true)
    _ = len(v1)
    if (v25 > (v25 / v25)) {
        v27 = (v2 % ((2) | 1))
        v27 = v2
        fmt.Println(v27, v104, v2)
    }
    v104 = (338 > 683)
    var v105 int = -1
    for fi106 := 0; fi106 < 17; fi106++ {
        var v107 uint = 4294967295
        _ = fi106
        _ = v107
    }
    fmt.Println(v105, v2)
    var v108 bool = (!(((0.1 - float64(float32(0.0))) - float64((v25 * -2.25))) == 1e10))
    if v104 {
        switch "\n\t" {
        case "":
            var v109 float32 = (1.5 + ((float32((1e10 * 1.5)) + 1e10) / v25))
            _ = v109
        case "foo bar":
            for fi110 := 0; fi110 < 7; fi110++ {
                if ((v108 || (!(-482 != 1))) && (-291 == -949)) {
                    for fi111 := 0; fi111 < 14; fi111++ {
                        switch (v105 | 2147483647) {
                        case 758:
                            var v112 int8 = -128
                            _ = v112
                        case -49:
                            fmt.Println(v27, v108)
                        }
                        if ((^(v2)) > v27) {
                            switch (fi111 ^ fi111) {
                            case 12:
                                _ = len(v3)
                            case 42:
                                var v113 int16 = -32768
                                _ = v113
                            }
                            var v114 uint = 1
                            _ = v114
                        } else {
                            v115 := GenS8{F0: 2}
                            fmt.Println(v105, v105, fi110)
                            _ = v115
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi111
                    }
                    if true {
                        _ = len(v1)
                        for fi116 := 0; fi116 < 18; fi116++ {
                            if (v104 && (((0.0 / 3.14159) / ((float64(v25) * 3.14159) - 3.14159)) <= 3.14159)) {
                                var v117 int64 = 79
                                switch "" {
                                case "hello":
                                    fmt.Println(v108, fi110, v2)
                                    v118 := GenS10{F0: 255, F1: v105}
                                    _ = v118
                                case "foo bar":
                                    v119 := GenS9{F0: v117, F1: (v117 | (v117 - (v117 * (-(v117)))))}
                                    _ = v119
                                case "\n\t":
                                    var v120 uint64 = 672
                                    _ = v120
                                }
                                switch (-(v105)) {
                                case -1:
                                    var v121 uint8 = 255
                                    _ = v121
                                case 42:
                                    var v122 uint64 = 19
                                    _ = v122
                                }
                                _ = v117
                            }
                            if v108 {
                                v27 -= (v2 + (v2 & v27))
                            } else {
                                v1[65535] = (fi110 | 42)
                                var v123 uint16 = 1
                                v124 := GenS11{F0: v104, F1: 42, F2: (^(v27))}
                                _ = v123
                                _ = v124
                            }
                            _ = fi116
                        }
                        v125 := GenS11{F0: false, F1: 18446744073709551615, F2: (v2 ^ v2)}
                        _ = v125
                    } else {
                        var v126 uint64 = 1
                        if ("\n\t" >= "\n\t") {
                            var v127 float32 = 1e10
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v127 = v25
                            _ = v127
                        }
                        var v128 int8 = 89
                        _ = v126
                        _ = v128
                    }
                } else {
                    v129 := GenS5{F0: -189, F1: 2147483647}
                    _ = v129
                }
                for fi130 := 0; fi130 < 2; fi130++ {
                    v104 = v104
                    switch 127 {
                    case 475:
                        v131 := GenS11{F0: (32767 == 0), F1: 169, F2: v2}
                        _ = v131
                    case 128:
                        if v108 {
                            v27 = 208
                            fmt.Println(fi110, fi130)
                        }
                    case 1:
                        if true {
                            switch "" {
                            case "":
                                var v132 uint32 = 0
                                _ = v132
                            }
                            var v133 uint16 = 972
                            v105 += (v105 | (fi110 % ((v105) | 1)))
                            _ = v133
                        }
                        v134 := GenS1{F0: "a", F1: -1, F2: 65}
                        _ = v134
                    default:
                        var v135 bool = (155 <= 628)
                        _ = v135
                    }
                    _ = fi130
                }
                _ = fi110
            }
        case "\n\t":
            if v108 {
                v2 = v27
            }
        default:
            var v136 int = 32767
            _ = v136
        }
    } else {
        var v137 uint16 = 868
        var v138 int8 = 100
        _ = v137
        _ = v138
    }
    if (v108 && (!v108)) {
        v139 := GenS12{F0: -2147483648, F1: 0}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v139
    } else {
        v2 = (^(v2))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("v108:", v108)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v104
    _ = v105
    _ = v108
    _ = v2
    _ = v25
    _ = v27
    _ = v28
    _ = v3
}
