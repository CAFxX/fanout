package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint32; F1 string; F2 bool }
type GenS8 struct { F0 uint; F1 bool }
type GenS9 struct { F0 string; F1 uint16; F2 uint16 }
type GenS10 struct { F0 string; F1 bool; F2 int8 }
type GenS11 struct { F0 uint; F1 int64 }
type GenS12 struct { F0 uint; F1 bool; F2 uint64 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 uint16; F1 uint16 }
func genFn0(p0 string) uint {
    var v6 bool = (3.14159 > 1.5)
    switch 9223372036854775807 {
    case 128:
        switch ((p0 + "\n\t") + "\n\t") {
        case "a":
            var v7 int64 = -9223372036854775808
            _ = v7
        case "hello":
            var v8 int8 = 55
            _ = v8
        case "\n\t":
            if v6 {
                var v9 string = (p0 + p0)
                _ = v9
            } else {
                v10 := GenS1{F0: 4294967295, F1: -9223372036854775808}
                if v6 {
                    fmt.Println(p0, p0)
                    v11 := GenS2{F0: 0, F1: 2, F2: 255}
                    v12 := GenS1{F0: 1, F1: 42}
                    _ = v11
                    _ = v12
                } else {
                    if v6 {
                        switch (p0 + "hello") {
                        case "":
                            var v13 int = -1
                            _ = v13
                        case "hello":
                            v14 := []int{32768, -1, 0}
                            _ = v14
                        case "foo bar":
                            var v15 uint32 = 42
                            _ = v15
                        default:
                            var v16 int = 127
                            _ = v16
                        }
                    } else {
                        p0 = p0
                        p0 = ""
                        for fi17 := 0; fi17 < 4; fi17++ {
                            var v18 int8 = 42
                            var v19 bool = (153 >= 2)
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                    }
                    switch 330 {
                    case 0:
                        var v20 uint8 = 241
                        _ = v20
                    case -100:
                        for fi21 := 0; fi21 < 10; fi21++ {
                            var v22 int8 = 1
                            _ = fi21
                            _ = v22
                        }
                    }
                    if true {
                        switch ("foo bar" + "\n\t") {
                        case "a":
                            for fi23 := 0; fi23 < 15; fi23++ {
                                v24 := GenS2{F0: -128, F1: 4294967295, F2: 9}
                                v6 = v6
                                if v6 {
                                    for fi25 := 0; fi25 < 2; fi25++ {
                                        v6 = v6
                                        v26 := []int{-1}
                                        _ = fi25
                                        _ = v26
                                    }
                                    var v27 int8 = 127
                                    _ = v27
                                } else {
                                    var v28 uint16 = 1
                                    v29 := []int{(fi23 >> ((106) & 63)), fi23}
                                    _ = v28
                                    _ = v29
                                }
                                _ = fi23
                                _ = v24
                            }
                        default:
                            for fi30 := 0; fi30 < 14; fi30++ {
                                var v31 bool = (-2.25 > 3.14159)
                                var v32 int = -9223372036854775808
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                        }
                        if v6 {
                            var v33 float64 = ((((0.1 - 1e10) + 3.14159) - (float64((-2.25 * 1.5)) - 3.14159)) + 1.5)
                            _ = v33
                        } else {
                            if v6 {
                                for fi34 := 0; fi34 < 14; fi34++ {
                                    var v35 uint16 = 32767
                                    _ = fi34
                                    _ = v35
                                }
                            }
                        }
                    }
                }
                switch (p0 + "hello") {
                case "a":
                    var v36 uint64 = 2
                    _ = v36
                case "foo bar":
                    var v37 int = 32767
                    _ = v37
                case "":
                    if (-238 <= -2147483648) {
                        fmt.Println(p0, p0)
                        v38 := []int{42, 179, 944}
                        v39 := GenS2{F0: 42, F1: 2, F2: 176}
                        _ = v38
                        _ = v39
                    } else {
                        var v40 int8 = -1
                        _ = v40
                    }
                default:
                    fmt.Println(v6, p0)
                }
                _ = v10
            }
        default:
            var v41 uint32 = 2147483648
            _ = v41
        }
    case 32767:
        v6 = v6
    case 127:
        switch "foo bar" {
        case "foo bar":
            var v42 int64 = 42
            _ = v42
        case "":
            v43 := GenS2{F0: 99, F1: 0, F2: 1}
            _ = v43
        }
    default:
        if (true && ((!(1e10 >= (0.1 + -2.25))) && ((1 == -850) || v6))) {
            var v44 uint64 = 18446744073709551615
            v45 := GenS1{F0: 723, F1: 326}
            _ = v44
            _ = v45
        } else {
            v6 = v6
            v6 = (1 >= 32767)
            for fi46 := 0; fi46 < 7; fi46++ {
                var v47 uint = 42
                _ = fi46
                _ = v47
            }
        }
    }
    _ = p0
    _ = v6
    return 255
}
func genFn1() uint32 {
    if false {
        var v48 bool = true
        _ = v48
    }
    var v49 float32 = -2.25
    v49 += v49
    v50 := GenS3{F0: 42, F1: 65535}
    if true {
        var v51 uint32 = 304
        var v52 float64 = -2.25
        _ = v51
        _ = v52
    } else {
        for fi53 := 0; fi53 < 14; fi53++ {
            for fi54 := 0; fi54 < 2; fi54++ {
                v55 := GenS4{F0: 34}
                _ = fi54
                _ = v55
            }
            var v56 uint = 196
            switch 9223372036854775807 {
            case 2147483647:
                switch "foo bar" {
                case "\n\t":
                    if (803 <= -761) {
                        if (((0 <= (^(v56))) && false) || ((true && (v49 > (1e10 * 3.14159))) || true)) {
                            v56 = v56
                        } else {
                            var v57 float64 = -2.25
                            fmt.Println(v56)
                            _ = v57
                        }
                    }
                case "foo bar":
                    var v58 uint32 = 1
                    _ = v58
                case "":
                    v59 := []int{fi53, (fi53 >> ((65535) & 63)), (-(fi53))}
                    _ = v59
                default:
                    var v60 string = ("\n\t" + (("foo bar" + "hello") + "foo bar"))
                    _ = v60
                }
            case 32768:
                var v61 uint64 = 1
                _ = v61
            case 9223372036854775807:
                switch 128 {
                case -2147483648:
                    v62 := []int{(^(fi53)), fi53}
                    _ = v62
                case 2147483647:
                    for fi63 := 0; fi63 < 5; fi63++ {
                        var v64 int8 = 121
                        _ = fi63
                        _ = v64
                    }
                }
            default:
                switch 2147483647 {
                case -52:
                    var v65 int32 = 778
                    _ = v65
                case -2147483648:
                    for fi66 := 0; fi66 < 9; fi66++ {
                        if false {
                            var v67 bool = (403 < 0)
                            if (float64(float32(1.5)) < 0.0) {
                                v68 := []int{(fi66 / ((fi53) | 1)), 224}
                                _ = v68
                            }
                            _ = v67
                        } else {
                            var v69 string = ("a" + "\n\t")
                            var v70 int = 1
                            var v71 int8 = 127
                            _ = v69
                            _ = v70
                            _ = v71
                        }
                        v72 := GenS1{F0: (^(v56)), F1: -1}
                        if (-78 == 42) {
                            var v73 bool = true
                            v74 := []int{42, 127}
                            if v73 {
                                if v73 {
                                    v75 := GenS1{F0: (v56 - 2), F1: 9223372036854775807}
                                    if true {
                                        var v76 int32 = 2147483647
                                        if false {
                                            v76 = (v76 * v76)
                                        } else {
                                            if (127 == (fi66 >> ((v56) & 63))) {
                                                var v77 int16 = -32768
                                                switch fi66 {
                                                case 32768:
                                                    if len(v74) >= 2 {
                                                        v74 = v74[1:len(v74)]
                                                    }
                                                case 42:
                                                    if ((255 == 64) && (v73 && (!(!v73)))) {
                                                        v76 -= v76
                                                    } else {
                                                        if v73 {
                                                            fmt.Println(fi66)
                                                            fmt.Println(v77)
                                                            if len(v74) > 0 {
                                                                v74[0] = (fi53 >> ((v56) & 63))
                                                            }
                                                            v78 := GenS3{F0: (^(v56)), F1: v56}
                                                            _ = v78
                                                        } else {
                                                            v77 = (v77 & -437)
                                                        }
                                                    }
                                                default:
                                                    if len(v74) >= 2 {
                                                        v74 = v74[1:len(v74)]
                                                    }
                                                }
                                                _ = v77
                                            }
                                            switch fi53 {
                                            case 32767:
                                                switch (fi66 >> (((v56 - (v56 / (((v56 << ((255) & 63))) | 1)))) & 63)) {
                                                case 127:
                                                    fmt.Println(fi53)
                                                case 504:
                                                    _ = len(v74)
                                                case -2147483648:
                                                    if len(v74) > 0 {
                                                        v74[0] = fi66
                                                    }
                                                default:
                                                    var v79 int8 = -19
                                                    _ = v79
                                                }
                                            }
                                            if ((!v73) && false) {
                                                if ("foo bar" > "a") {
                                                    v49 = float32(((0.1 / -2.25) + 0.1))
                                                }
                                                var v80 int8 = 42
                                                if len(v74) > 0 {
                                                    v74[0] = (fi66 % ((fi53) | 1))
                                                }
                                                _ = v80
                                            }
                                        }
                                        var v81 int = fi53
                                        v56 = (v56 << ((18446744073709551615) & 63))
                                        _ = v76
                                        _ = v81
                                    }
                                    fmt.Println(fi53, v49, v73)
                                    _ = v75
                                } else {
                                    if v73 {
                                        var v82 bool = v73
                                        var v83 float32 = 0.0
                                        fmt.Println(v49)
                                        switch (("\n\t" + "") + "\n\t") {
                                        case "foo bar":
                                            fmt.Println(v82)
                                        case "hello":
                                            v84 := GenS5{F0: 192}
                                            _ = v84
                                        case "a":
                                            if v73 {
                                                var v85 int32 = 2147483647
                                                _ = v85
                                            } else {
                                                v73 = v82
                                                v82 = (v73 || (!v73))
                                            }
                                        }
                                        v49 -= v83
                                        _ = v82
                                        _ = v83
                                    }
                                    var v86 int32 = 349
                                    v73 = false
                                    _ = v86
                                }
                                v73 = v73
                                var v87 int = fi53
                                _ = v87
                            } else {
                                _ = len(v74)
                                switch -1 {
                                case 32767:
                                    v88 := GenS6{F0: ("" + ("hello" + "a"))}
                                    _ = v88
                                case 1:
                                    v49 += (1.5 * v49)
                                default:
                                    fmt.Println(v56, v73, v56)
                                }
                                if (110 <= 128) {
                                    if v73 {
                                        if (-447 != -1) {
                                            if len(v74) >= 2 {
                                                v74 = v74[1:len(v74)]
                                            }
                                        } else {
                                            var v89 uint = 255
                                            if (231 > -1) {
                                                if len(v74) >= 2 {
                                                    v74 = v74[1:len(v74)]
                                                }
                                                if len(v74) >= 2 {
                                                    v74 = v74[1:len(v74)]
                                                }
                                                v89 += 1
                                                if v73 {
                                                    if (((2147483647 == -2147483648) || v73) && (28 > genFn0("foo bar"))) {
                                                        var v90 int8 = -75
                                                        _ = v90
                                                    }
                                                    v56 += genFn0(("" + ""))
                                                    v74 = append(v74, fi66)
                                                }
                                            } else {
                                                v56 = (v56 >> ((331) & 63))
                                                if v73 {
                                                    fmt.Println(v49, v49)
                                                    v91 := GenS1{F0: v56, F1: -9223372036854775808}
                                                    var v92 float64 = 0.0
                                                    _ = v91
                                                    _ = v92
                                                } else {
                                                    var v93 int = fi66
                                                    _ = v93
                                                }
                                                var v94 uint = v56
                                                v95 := GenS4{F0: 41}
                                                v74 = append(v74, fi66)
                                                var v96 string = ""
                                                _ = v94
                                                _ = v95
                                                _ = v96
                                            }
                                            _ = v89
                                        }
                                        if true {
                                            var v97 uint = (^(v56))
                                            fmt.Println(v73)
                                            _ = v97
                                        } else {
                                            v73 = v73
                                        }
                                        var v98 int64 = -9223372036854775808
                                        _ = v98
                                    } else {
                                        v74 = append(v74, fi53)
                                        var v99 float32 = (v49 / (v49 / ((v49 * v49) * v49)))
                                        switch (("hello" + "hello") + (("foo bar" + ("foo bar" + "hello")) + ("" + "foo bar"))) {
                                        case "a":
                                            var v100 uint8 = 128
                                            _ = v100
                                        default:
                                            if v73 {
                                                v49 += v49
                                            }
                                        }
                                        var v101 string = "foo bar"
                                        if false {
                                            v99 = v49
                                            v101 = (v101 + v101)
                                            v102 := GenS4{F0: 2}
                                            _ = v102
                                        }
                                        if v73 {
                                            if v73 {
                                                v49 = float32(0.1)
                                                if (true && v73) {
                                                    v101 = (v101 + v101)
                                                    var v103 uint = 18446744073709551615
                                                    _ = v103
                                                } else {
                                                    if len(v74) >= 2 {
                                                        v74 = v74[1:len(v74)]
                                                    }
                                                    v104 := GenS3{F0: (v56 * 351), F1: (^(v56))}
                                                    _ = v104
                                                }
                                            }
                                            v101 = v101
                                        } else {
                                            if len(v74) > 0 {
                                                v74[0] = fi53
                                            }
                                            var v105 uint16 = 42
                                            _ = v105
                                        }
                                        _ = v101
                                        _ = v99
                                    }
                                    v56 = 42
                                } else {
                                    v56 = v56
                                    if (!v73) {
                                        v73 = (true && (-1 < 0))
                                        if (!v73) {
                                            if v73 {
                                                var v106 uint = 2
                                                v74 = append(v74, (fi53 / ((127) | 1)))
                                                _ = v106
                                            } else {
                                                v74 = append(v74, (fi53 ^ fi53))
                                                v107 := GenS5{F0: 244}
                                                _ = v107
                                            }
                                            v56 = (^(v56))
                                        }
                                    }
                                }
                            }
                            _ = v73
                            _ = v74
                        } else {
                            var v108 uint32 = 412
                            _ = v108
                        }
                        _ = fi66
                        _ = v72
                    }
                default:
                    var v109 uint = 42
                    _ = v109
                }
            }
            _ = fi53
            _ = v56
        }
        var v110 uint = 4294967295
        _ = v110
    }
    _ = v49
    _ = v50
    return 2147483648
}

func main() {
    var v1 int8 = 1
    var v2 uint8 = 1
    v3 := make(map[int]bool)
    var v4 uint64 = 18446744073709551615
    var v5 uint64 = v4
    for fi111 := 0; fi111 < 10; fi111++ {
        v112 := GenS7{F0: 38, F1: ("\n\t" + ""), F2: (!false)}
        if (1.5 >= 0.1) {
            if true {
                v4 = v4
                switch fi111 {
                case -565:
                    fmt.Println(v1, v5)
                case 0:
                    fmt.Println(v2, v2, v5)
                case 32768:
                    for fi113 := 0; fi113 < 2; fi113++ {
                        if ((-131 >= -1) && ((!(600 == 0)) || true)) {
                            var v114 int8 = 42
                            v5 = 278
                            _ = v114
                        }
                        v115 := []int{-2147483648, fi113, (fi113 & 0), fi111}
                        _ = fi113
                        _ = v115
                    }
                }
                switch "\n\t" {
                case "":
                    switch 32767 {
                    case 2147483647:
                        switch ("foo bar" + "\n\t") {
                        case "foo bar":
                            v116 := GenS7{F0: 65535, F1: (("a" + ("foo bar" + "hello")) + "foo bar"), F2: false}
                            _ = v116
                        }
                    case -9223372036854775808:
                        v1 = (v1 / ((-128) | 1))
                    default:
                        _ = len(v3)
                    }
                case "a":
                    var v117 uint16 = 683
                    _ = v117
                default:
                    v118 := GenS8{F0: uint(v5), F1: (true && (-2.25 == (((float64(3.14159) + (((1e10 / (((1e10 - ((0.1 / -2.25) / -2.25)) - ((1.5 / (-2.25 * (3.14159 + ((0.1 / (-2.25 - ((0.1 / (1e10 + -2.25)) + -2.25))) * ((float64(1.5) * ((0.1 * 3.14159) - -2.25)) * 1e10))))) * 3.14159)) * 0.1)) + 1e10) / 3.14159)) * (1.5 / 1e10)) / 1e10)))}
                    _ = v118
                }
            }
        }
        switch 32768 {
        case 382:
            if true {
                var v119 int64 = 576
                switch "" {
                case "foo bar":
                    for fi120 := 0; fi120 < 17; fi120++ {
                        v121 := GenS4{F0: 1}
                        switch (fi120 | (fi120 >> ((4294967295) & 63))) {
                        case -1:
                            var v122 int16 = 1
                            _ = v122
                        case 801:
                            v123 := []int{-9223372036854775808, (^(fi120)), 1}
                            _ = v123
                        }
                        v124 := GenS7{F0: 65535, F1: ((("foo bar" + "\n\t") + "a") + (("foo bar" + ("foo bar" + ("foo bar" + (("foo bar" + ("foo bar" + "\n\t")) + "")))) + ("" + (("" + "foo bar") + "hello")))), F2: (!(false && (true && true)))}
                        _ = fi120
                        _ = v121
                        _ = v124
                    }
                case "a":
                    switch "" {
                    case "":
                        if false {
                            switch "hello" {
                            case "hello":
                                v125 := GenS9{F0: "hello", F1: 49, F2: 32767}
                                _ = v125
                            }
                        } else {
                            v126 := []int{}
                            var v127 uint64 = (v5 + 999)
                            switch ("hello" + ("foo bar" + "")) {
                            case "\n\t":
                                if (-128 == (v1 - 100)) {
                                    v3[fi111] = (true && true)
                                    fmt.Println(v5, v119)
                                    var v128 int16 = 32767
                                    _ = v128
                                }
                            default:
                                var v129 int32 = 42
                                _ = v129
                            }
                            _ = v126
                            _ = v127
                        }
                    }
                default:
                    var v130 int64 = 1
                    _ = v130
                }
                var v131 int = 4294967295
                _ = v119
                _ = v131
            } else {
                fmt.Println(fi111)
            }
        case 396:
            switch "a" {
            case "foo bar":
                var v132 int32 = 2147483647
                _ = v132
            case "\n\t":
                switch (-(fi111)) {
                case 42:
                    var v133 string = ("foo bar" + ("" + "foo bar"))
                    _ = v133
                default:
                    v2 = v2
                }
            case "":
                for fi134 := 0; fi134 < 11; fi134++ {
                    if (!(false && (false && (0.0 > 0.1)))) {
                        var v135 int32 = 42
                        _ = v135
                    } else {
                        if (2147483647 >= 97) {
                            var v136 float64 = 0.0
                            _ = v136
                        }
                        v1 += v1
                    }
                    _ = fi134
                }
            }
        case -598:
            var v137 bool = (989 < 65535)
            _ = v137
        }
        _ = fi111
        _ = v112
    }
    var v138 uint16 = 65535
    if ((((1e10 / 1e10) * 3.14159) + 1.5) >= ((1.5 * 1.5) - 3.14159)) {
        for fi139 := 0; fi139 < 12; fi139++ {
            if ((^(v5)) <= 809) {
                v138 = v138
                switch (fi139 - (fi139 + fi139)) {
                case -2147483648:
                    if (!(false || (true && (fi139 != (^(fi139)))))) {
                        v2 = (v2 % ((v2) | 1))
                        if (362 == 2) {
                            if false {
                                switch fi139 {
                                case -914:
                                    var v140 uint32 = 2
                                    _ = v140
                                case 127:
                                    v141 := []int{int(v4)}
                                    _ = v141
                                default:
                                    fmt.Println(fi139, v4, v5)
                                }
                                switch fi139 {
                                case 378:
                                    var v142 uint64 = v4
                                    _ = v142
                                case -658:
                                    if (9223372036854775807 > -1) {
                                        v143 := v3[(fi139 >> ((2) & 63))]
                                        for fi144 := 0; fi144 < 5; fi144++ {
                                            var v145 int = fi144
                                            _ = fi144
                                            _ = v145
                                        }
                                        _ = v143
                                    }
                                default:
                                    var v146 float32 = (3.14159 + 3.14159)
                                    _ = v146
                                }
                            }
                            v147 := GenS7{F0: uint32(v5), F1: "foo bar", F2: ((v1 != (v1 | (^(v1)))) || false)}
                            for fi148 := 0; fi148 < 17; fi148++ {
                                if true {
                                    v1 = v1
                                } else {
                                    v2 = (^(v2))
                                }
                                v5 = v5
                                _ = fi148
                            }
                            _ = v147
                        }
                    }
                case 42:
                    fmt.Println(fi139, v5)
                case -333:
                    v138 = 32767
                }
            }
            if false {
                v4 += 4294967295
                for fi149 := 0; fi149 < 12; fi149++ {
                    switch fi149 {
                    case -9223372036854775808:
                        v2 -= v2
                    case 42:
                        v150 := GenS10{F0: "a", F1: (1 >= 0), F2: 100}
                        _ = v150
                    }
                    _ = fi149
                }
            } else {
                var v151 int32 = -400
                for fi152 := 0; fi152 < 2; fi152++ {
                    if (((0.1 * 0.1) >= 3.14159) || true) {
                        v5 += 854
                        v151 = (v151 / ((0) | 1))
                        var v153 int64 = int64(v151)
                        _ = v153
                    } else {
                        v151 = v151
                        var v154 int = (-(fi139))
                        var v155 uint64 = v5
                        _ = v154
                        _ = v155
                    }
                    if (false || (((0.0 / 1.0) - 1.5) > float32(1.5))) {
                        v156 := GenS11{F0: 255, F1: -184}
                        v1 = v1
                        if (!true) {
                            v151 = v151
                            fmt.Println(fi152)
                        } else {
                            v157 := GenS7{F0: 452, F1: (("\n\t" + "a") + ("" + "foo bar")), F2: false}
                            v2 += v2
                            if true {
                                v1 = (v1 % ((0) | 1))
                            }
                            _ = v157
                        }
                        if true {
                            if (1.5 != (0.0 + 0.1)) {
                                if (-2.25 < (float32(0.1) - (0.1 / 0.1))) {
                                    fmt.Println(v5, v2)
                                    var v158 float32 = 1e10
                                    _ = v158
                                }
                            } else {
                                v5 -= 660
                            }
                        }
                        _ = v156
                    } else {
                        var v159 string = ""
                        if (((v1 * v1) < (v1 ^ -128)) || false) {
                            v138 = (v138 + 255)
                            var v160 int64 = 9223372036854775807
                            _ = v160
                        }
                        _ = v159
                    }
                    _ = fi152
                }
                v161 := v3[fi139]
                _ = v151
                _ = v161
            }
            _ = fi139
        }
    }
    v162 := GenS3{F0: 255, F1: 18446744073709551615}
    if false {
        switch "a" {
        case "":
            v2 -= 63
        case "hello":
            var v163 float64 = 1e10
            _ = v163
        default:
            fmt.Println(v4)
        }
        var v164 uint16 = (v138 >> ((0) & 63))
        for fi165 := 0; fi165 < 14; fi165++ {
            v138 = (v164 + v138)
            _ = fi165
        }
        _ = v164
    } else {
        for fi166 := 0; fi166 < 8; fi166++ {
            v2 += 164
            v167 := GenS12{F0: 535, F1: true, F2: (v4 % ((6) | 1))}
            v5 = (v5 - v4)
            _ = fi166
            _ = v167
        }
        v138 = (^(v138))
    }
    var v168 int32 = 0
    switch 128 {
    case 0:
        var v169 float32 = 3.14159
        _ = v169
    case -666:
        v138 = (v138 ^ v138)
    case 2147483647:
        switch -750 {
        case 42:
            v170 := []int{9223372036854775807}
            _ = v170
        default:
            v171 := GenS13{F0: (v168 | (v168 << ((18446744073709551615) & 63)))}
            _ = v171
        }
    }
    switch (("hello" + "a") + "\n\t") {
    case "foo bar":
        v172 := GenS14{F0: v138, F1: 2}
        _ = v172
    case "a":
        fmt.Println(v4, v138, v5)
    }
    var v173 uint32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v138:", v138)
    fmt.Println("v168:", v168)
    fmt.Println("v173:", v173)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v138
    _ = v162
    _ = v168
    _ = v173
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
