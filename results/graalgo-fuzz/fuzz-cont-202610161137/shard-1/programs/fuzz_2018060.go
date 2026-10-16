package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 int64; F1 int16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 string; F1 uint }
type GenS7 struct { F0 int; F1 uint8; F2 int8 }
type GenS8 struct { F0 string; F1 int32 }
type GenS9 struct { F0 bool; F1 uint8; F2 int }
type GenS10 struct { F0 uint32; F1 bool }
func genFn0(p0 int) int64 {
    p0 -= -9223372036854775808
    v6 := []int{p0, p0, (p0 ^ (p0 ^ p0))}
    _ = p0
    _ = v6
    return 42
}
func genFn1(p0 uint64) int8 {
    p0 -= (^(p0))
    fmt.Println(p0)
    var v7 uint = 4294967295
    fmt.Println(p0)
    _ = p0
    _ = v7
    return 70
}

func main() {
    var v1 uint8 = 185
    var v2 uint = 23
    var v3 bool = (false && true)
    var v4 int16 = 32767
    v5 := make(map[string]string)
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 int64 = -1
        v9 = (v9 % (((v9 ^ (v9 * (^(v9))))) | 1))
        v10 := v5[""]
        _ = fi8
        _ = v10
        _ = v9
    }
    v5[(("\n\t" + ("hello" + "")) + "a")] = ((("" + ("\n\t" + "hello")) + "") + ("a" + "hello"))
    v4 = v4
    fmt.Println(v1, v2)
    var v11 int = 2147483647
    if true {
        v11 += (v11 * (v11 | (v11 ^ v11)))
        var v12 string = "\n\t"
        _ = v12
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        switch "hello" {
        case "foo bar":
            for fi14 := 0; fi14 < 12; fi14++ {
                v3 = v3
                v15 := GenS1{F0: v4, F1: ("" + "a")}
                _ = fi14
                _ = v15
            }
        case "hello":
            v4 -= (v4 - v4)
        }
        var v16 uint32 = 143
        switch "a" {
        case "":
            v17 := GenS1{F0: (v4 / ((v4) | 1)), F1: "foo bar"}
            _ = v17
        default:
            var v18 uint16 = 0
            _ = v18
        }
        _ = fi13
        _ = v16
    }
    if v3 {
        v1 -= 0
        v19 := v5["a"]
        var v20 int16 = (v4 << ((18446744073709551615) & 63))
        _ = v19
        _ = v20
    } else {
        v4 += v4
        v21 := GenS1{F0: v4, F1: ((("a" + ("a" + "")) + ("\n\t" + (("a" + "a") + "foo bar"))) + ("\n\t" + "\n\t"))}
        v11 -= int(v4)
        _ = v21
    }
    var v22 int = 1
    v11 += (v11 & v22)
    for fi23 := 0; fi23 < 4; fi23++ {
        var v24 uint32 = 707
        if true {
            v11 = 4294967295
            fmt.Println(v3, fi23, v24)
            v25 := GenS1{F0: 0, F1: "\n\t"}
            _ = v25
        }
        _ = fi23
        _ = v24
    }
    if v3 {
        var v26 int64 = genFn0(int(v4))
        switch v22 {
        case -1:
            delete(v5, (("a" + ("" + "hello")) + "\n\t"))
        default:
            if (v3 && (true || v3)) {
                delete(v5, "hello")
                for fi27 := 0; fi27 < 16; fi27++ {
                    v28 := GenS1{F0: v4, F1: ""}
                    _ = fi27
                    _ = v28
                }
                v2 -= 1
            }
        }
        for fi29 := 0; fi29 < 14; fi29++ {
            var v30 float64 = -2.25
            _ = fi29
            _ = v30
        }
        _ = v26
    } else {
        switch -498 {
        case 32768:
            for fi31 := 0; fi31 < 11; fi31++ {
                if true {
                    v4 = (v4 >> ((255) & 63))
                    delete(v5, (("" + "") + "\n\t"))
                }
                _ = fi31
            }
        case -405:
            var v32 string = "\n\t"
            _ = v32
        }
    }
    fmt.Println(v22)
    if v3 {
        if (false || ((v4 | 32767) <= -1)) {
            switch v22 {
            case 2147483647:
                fmt.Println(v2)
            case 9223372036854775807:
                v33 := GenS2{F0: false, F1: v2}
                _ = v33
            }
        } else {
            if (v4 > int16(v1)) {
                if true {
                    _ = len(v5)
                    switch v22 {
                    case 128:
                        for fi34 := 0; fi34 < 12; fi34++ {
                            if (true || (!(v2 >= (v2 >> ((v2) & 63))))) {
                                var v35 int8 = -99
                                fmt.Println(v2, v11)
                                v36 := []int{}
                                _ = v35
                                _ = v36
                            } else {
                                fmt.Println(v4, v11)
                                var v37 bool = (v1 == v1)
                                _ = v37
                            }
                            for fi38 := 0; fi38 < 1; fi38++ {
                                var v39 uint32 = 1
                                _ = fi38
                                _ = v39
                            }
                            _ = fi34
                        }
                    default:
                        var v40 uint16 = 42
                        _ = v40
                    }
                    switch "hello" {
                    case "foo bar":
                        v4 -= int16(v1)
                    case "a":
                        v2 = (v2 ^ v2)
                    default:
                        _ = len(v5)
                    }
                }
                v41 := []int{}
                v11 += v11
                _ = v41
            }
        }
        var v42 uint32 = 1
        _ = v42
    }
    if (false || false) {
        switch -345 {
        case -834:
            v3 = (!true)
        case 42:
            v43 := []int{v22, (v11 & 1), v11, (v22 % (((v11 - (^(v11)))) | 1))}
            _ = v43
        }
        if v3 {
            v1 = 128
            switch v22 {
            case 9223372036854775807:
                v44 := []int{}
                _ = v44
            case 128:
                switch v11 {
                case 42:
                    v22 += (v22 >> ((v2) & 63))
                case -708:
                    switch 482 {
                    case -2147483648:
                        switch "foo bar" {
                        case "":
                            var v45 uint = 394
                            _ = v45
                        case "foo bar":
                            fmt.Println(v11)
                        case "\n\t":
                            if (((^(v2)) > v2) && (v1 != v1)) {
                                v46 := []int{v11}
                                var v47 uint8 = (v1 - v1)
                                _ = v46
                                _ = v47
                            }
                        default:
                            v2 = v2
                        }
                    case 9223372036854775807:
                        v48 := GenS2{F0: true, F1: 2}
                        _ = v48
                    }
                }
            case 350:
                v49 := GenS1{F0: v4, F1: "\n\t"}
                _ = v49
            }
            var v50 float64 = ((3.14159 + 3.14159) * 3.14159)
            _ = v50
        }
    } else {
        if (v3 || v3) {
            var v51 int = (v11 & v22)
            _ = v51
        } else {
            fmt.Println(v2, v4, v3)
            if (!false) {
                v52 := GenS2{F0: v3, F1: (v2 + v2)}
                fmt.Println(v22, v2, v11)
                v4 -= (v4 + v4)
                _ = v52
            }
        }
        v53 := GenS3{F0: 4294967295, F1: v1}
        var v54 float64 = float64(((1.5 - 1e10) / 1.5))
        _ = v53
        _ = v54
    }
    if v3 {
        if ((-(v11)) < (v11 / ((-8) | 1))) {
            switch ("\n\t" + "foo bar") {
            case "":
                fmt.Println(v1, v1, v2)
            case "hello":
                v55 := v5[(("" + "hello") + "hello")]
                _ = v55
            default:
                v56 := []int{v11, (v11 & 32768), (v22 * 9223372036854775807), (v22 << (((^(v2))) & 63))}
                _ = v56
            }
            if (!v3) {
                v2 += v2
            } else {
                for fi57 := 0; fi57 < 6; fi57++ {
                    switch "a" {
                    case "a":
                        fmt.Println(fi57)
                    case "":
                        switch v11 {
                        case 128:
                            var v58 uint64 = 42
                            _ = v58
                        case 32767:
                            var v59 string = "\n\t"
                            _ = v59
                        case 220:
                            v1 = (v1 % ((255) | 1))
                        }
                    case "\n\t":
                        var v60 int64 = genFn0((-(v22)))
                        _ = v60
                    }
                    _ = fi57
                }
                for fi61 := 0; fi61 < 14; fi61++ {
                    var v62 int32 = 1
                    _ = fi61
                    _ = v62
                }
            }
            var v63 uint64 = 4294967295
            _ = v63
        }
        for fi64 := 0; fi64 < 10; fi64++ {
            v2 = 0
            v65 := GenS1{F0: (v4 & -1), F1: ("hello" + (((("hello" + "hello") + "") + ("a" + "")) + "hello"))}
            v11 += 32768
            _ = fi64
            _ = v65
        }
    } else {
        v66 := []int{(-(v11))}
        if (((!v3) && v3) || v3) {
            for fi67 := 0; fi67 < 2; fi67++ {
                var v68 uint32 = 42
                v69 := v5[(("foo bar" + "foo bar") + "hello")]
                _ = fi67
                _ = v68
                _ = v69
            }
            _ = len(v66)
            for fi70 := 0; fi70 < 15; fi70++ {
                var v71 string = ""
                v11 = v11
                if (367 >= 2) {
                    var v72 int8 = 0
                    if (255 > uint8(v2)) {
                        var v73 uint32 = 1
                        v4 = 83
                        var v74 uint32 = 1
                        _ = v73
                        _ = v74
                    }
                    _ = v72
                } else {
                    fmt.Println(v71)
                    for fi75 := 0; fi75 < 8; fi75++ {
                        var v76 int32 = 624
                        if (genFn0((v11 >> ((v2) & 63))) != int64(v1)) {
                            if (!v3) {
                                v3 = (v2 != (^(v2)))
                                var v77 float32 = float32(1e10)
                                v78 := GenS4{F0: 42, F1: -208}
                                v1 += (v1 - (v1 / ((v1) | 1)))
                                var v79 bool = (697 == genFn0(v22))
                                _ = v77
                                _ = v78
                                _ = v79
                            }
                            if v3 {
                                if true {
                                    v2 = 2
                                    v4 += v4
                                }
                                v80 := GenS4{F0: 0, F1: -751}
                                var v81 int32 = -2147483648
                                v1 = v1
                                v1 = (^(v1))
                                var v82 float32 = ((1e10 / (((-2.25 / 1e10) * -2.25) + (1e10 + (-2.25 + (3.14159 * (0.0 + ((0.1 / 3.14159) + float32(1e10)))))))) - float32(1e10))
                                _ = v80
                                _ = v81
                                _ = v82
                            } else {
                                v66 = append(v66, fi70)
                            }
                        } else {
                            var v83 uint = v2
                            var v84 uint32 = 397
                            _ = v83
                            _ = v84
                        }
                        v1 += (v1 - (v1 + (v1 % ((v1) | 1))))
                        _ = fi75
                        _ = v76
                    }
                    var v85 float32 = -2.25
                    _ = v85
                }
                _ = fi70
                _ = v71
            }
        } else {
            var v86 uint32 = 499
            for fi87 := 0; fi87 < 12; fi87++ {
                _ = len(v66)
                for fi88 := 0; fi88 < 9; fi88++ {
                    v86 = (v86 ^ v86)
                    v1 = (v1 * uint8(v2))
                    _ = fi88
                }
                _ = fi87
            }
            if true {
                if v3 {
                    if false {
                        if v3 {
                            if false {
                                if v3 {
                                    v89 := GenS5{F0: v4}
                                    var v90 int64 = 627
                                    _ = v89
                                    _ = v90
                                }
                                var v91 float64 = 0.1
                                _ = v91
                            } else {
                                switch "foo bar" {
                                case "a":
                                    v22 += v11
                                case "\n\t":
                                    v86 = v86
                                default:
                                    switch v11 {
                                    case 42:
                                        var v92 uint = (v2 << ((v2) & 63))
                                        _ = v92
                                    case 0:
                                        if (!v3) {
                                            var v93 int64 = 9223372036854775807
                                            var v94 int16 = 42
                                            v2 -= 277
                                            _ = v93
                                            _ = v94
                                        }
                                    case 128:
                                        fmt.Println(v2, v22)
                                    default:
                                        v4 = 1000
                                    }
                                }
                                if ((v2 << ((109) & 63)) > 65535) {
                                    if false {
                                        switch (v11 + v22) {
                                        case 42:
                                            fmt.Println(v11, v11)
                                        }
                                        switch (v22 - 2147483647) {
                                        case 552:
                                            v66 = append(v66, (v11 | v11))
                                        default:
                                            if (37 <= 0) {
                                                v3 = (!((v4 + v4) != (v4 + v4)))
                                            } else {
                                                var v95 int8 = 1
                                                _ = v95
                                            }
                                        }
                                        v11 -= -480
                                    } else {
                                        switch "foo bar" {
                                        case "\n\t":
                                            for fi96 := 0; fi96 < 17; fi96++ {
                                                if ((v1 + (v1 - 205)) <= v1) {
                                                    var v97 int = v11
                                                    _ = v97
                                                }
                                                if (!v3) {
                                                    var v98 uint32 = v86
                                                    var v99 uint64 = 743
                                                    _ = v98
                                                    _ = v99
                                                } else {
                                                    if ((!v3) || v3) {
                                                        fmt.Println(v22, v3, v11)
                                                    } else {
                                                        if v3 {
                                                            for fi100 := 0; fi100 < 18; fi100++ {
                                                                if false {
                                                                    var v101 uint64 = 1
                                                                    var v102 uint16 = 1000
                                                                    _ = v101
                                                                    _ = v102
                                                                } else {
                                                                    var v103 int64 = 1
                                                                    var v104 float32 = (float32(1e10) / 1.0)
                                                                    v1 = (^(v1))
                                                                    _ = v103
                                                                    _ = v104
                                                                }
                                                                v86 += (v86 - (v86 ^ v86))
                                                                _ = fi100
                                                            }
                                                            _ = len(v5)
                                                        } else {
                                                            v105 := v5[(("" + ("a" + "\n\t")) + ((("hello" + "foo bar") + (("foo bar" + "hello") + "")) + "foo bar"))]
                                                            fmt.Println(v105, v4)
                                                            _ = v105
                                                        }
                                                    }
                                                    v106 := GenS6{F0: "\n\t", F1: 42}
                                                    _ = v106
                                                }
                                                fmt.Println(v22, v22)
                                                _ = fi96
                                            }
                                        case "":
                                            v1 -= v1
                                        default:
                                            for fi107 := 0; fi107 < 10; fi107++ {
                                                for fi108 := 0; fi108 < 11; fi108++ {
                                                    v109 := GenS5{F0: (v4 & (v4 & v4))}
                                                    _ = fi108
                                                    _ = v109
                                                }
                                                _ = fi107
                                            }
                                        }
                                        var v110 uint64 = 4294967295
                                        _ = v110
                                    }
                                    v111 := GenS6{F0: "", F1: uint(v1)}
                                    var v112 uint16 = 2
                                    _ = v111
                                    _ = v112
                                }
                                v2 = (v2 / ((v2) | 1))
                            }
                            v113 := GenS7{F0: (-(v22)), F1: 97, F2: 0}
                            if (false || false) {
                                fmt.Println(v4, v86)
                                switch (("hello" + "hello") + "") {
                                case "hello":
                                    v22 -= -1
                                case "a":
                                    switch v22 {
                                    case 127:
                                        if ((v4 << (((^(v2))) & 63)) < v4) {
                                            if true {
                                                fmt.Println(v22)
                                                var v114 bool = v3
                                                if (!((2 >= 4294967295) || v3)) {
                                                    delete(v5, "a")
                                                }
                                                _ = v114
                                            } else {
                                                v11 = v11
                                                var v115 int16 = v4
                                                _ = v115
                                            }
                                        }
                                    }
                                case "\n\t":
                                    if ("a" == ("foo bar" + "\n\t")) {
                                        var v116 uint = (^(v2))
                                        var v117 string = ("a" + "")
                                        fmt.Println(v116, v3, v86)
                                        _ = v116
                                        _ = v117
                                    } else {
                                        fmt.Println(v86)
                                    }
                                }
                                v118 := v5[("hello" + "")]
                                _ = v118
                            } else {
                                switch v11 {
                                case 777:
                                    if v3 {
                                        v2 = (v2 % (((v2 + v2)) | 1))
                                        var v119 int = (v11 & -225)
                                        v22 = 498
                                        _ = v119
                                    }
                                }
                            }
                            _ = v113
                        } else {
                            if len(v66) >= 2 {
                                v66 = v66[1:len(v66)]
                            }
                            fmt.Println(v1, v2, v3)
                        }
                        var v120 int = (v22 ^ v22)
                        var v121 uint16 = 1
                        _ = v120
                        _ = v121
                    }
                    v3 = (v3 && v3)
                    if v3 {
                        switch "hello" {
                        case "hello":
                            v122 := GenS1{F0: v4, F1: "hello"}
                            _ = v122
                        default:
                            if v3 {
                                switch v11 {
                                case 4294967295:
                                    switch ("hello" + ("a" + ("foo bar" + "a"))) {
                                    case "hello":
                                        var v123 uint16 = 0
                                        _ = v123
                                    case "\n\t":
                                        v124 := GenS2{F0: ((!(1e10 < 1.5)) && (2147483648 != 176)), F1: (^(v2))}
                                        _ = v124
                                    case "a":
                                        for fi125 := 0; fi125 < 3; fi125++ {
                                            v11 -= int(v86)
                                            switch v11 {
                                            case 1:
                                                v126 := GenS1{F0: v4, F1: "a"}
                                                _ = v126
                                            case -9223372036854775808:
                                                v127 := GenS7{F0: (v11 / ((128) | 1)), F1: v1, F2: -128}
                                                _ = v127
                                            case 0:
                                                switch v11 {
                                                case 0:
                                                    for fi128 := 0; fi128 < 4; fi128++ {
                                                        v66 = append(v66, v22)
                                                        var v129 uint = v2
                                                        _ = fi128
                                                        _ = v129
                                                    }
                                                }
                                            }
                                            fmt.Println(fi125, v22)
                                            _ = fi125
                                        }
                                    default:
                                        v11 = (v11 - v11)
                                    }
                                case 1:
                                    var v130 int = int(v1)
                                    _ = v130
                                case 392:
                                    for fi131 := 0; fi131 < 5; fi131++ {
                                        var v132 float64 = (-2.25 / 1.0)
                                        if (295 > v86) {
                                            v11 += v11
                                            _ = len(v66)
                                        }
                                        _ = fi131
                                        _ = v132
                                    }
                                default:
                                    v133 := GenS8{F0: ("\n\t" + "a"), F1: 433}
                                    _ = v133
                                }
                                var v134 uint8 = v1
                                var v135 int32 = -1
                                _ = v134
                                _ = v135
                            } else {
                                delete(v5, ("" + "foo bar"))
                                for fi136 := 0; fi136 < 8; fi136++ {
                                    v137 := GenS9{F0: true, F1: 42, F2: (v22 - 892)}
                                    fmt.Println(v4, v22, v11)
                                    var v138 float32 = 0.0
                                    _ = fi136
                                    _ = v137
                                    _ = v138
                                }
                            }
                        }
                        v11 += (-(v11))
                    }
                } else {
                    if v3 {
                        var v139 string = ("hello" + "foo bar")
                        v140 := GenS6{F0: (v139 + v139), F1: (^(v2))}
                        var v141 uint32 = 0
                        _ = v139
                        _ = v140
                        _ = v141
                    }
                }
                switch (-(v22)) {
                case 4294967295:
                    for fi142 := 0; fi142 < 10; fi142++ {
                        if (true && (!(42 >= 399))) {
                            var v143 uint16 = 846
                            _ = v143
                        } else {
                            if v3 {
                                v144 := GenS10{F0: 65535, F1: (!((v4 ^ 1000) < (v4 % ((v4) | 1))))}
                                _ = len(v66)
                                if (307 <= 1) {
                                    v145 := GenS1{F0: v4, F1: "foo bar"}
                                    if false {
                                        switch "a" {
                                        case "hello":
                                            v5["hello"] = ""
                                        case "foo bar":
                                            var v146 float32 = 1e10
                                            _ = v146
                                        }
                                        var v147 int16 = v4
                                        _ = v147
                                    }
                                    switch (fi142 * v22) {
                                    case 610:
                                        if (-1 < -2147483648) {
                                            v148 := GenS3{F0: (v2 * v2), F1: 0}
                                            _ = v148
                                        }
                                    case 32767:
                                        if v3 {
                                            v22 = v22
                                            var v149 bool = (v4 <= 0)
                                            _ = v149
                                        }
                                    default:
                                        v1 -= 42
                                    }
                                    _ = v145
                                }
                                _ = v144
                            }
                        }
                        var v150 uint16 = 1
                        _ = len(v5)
                        _ = fi142
                        _ = v150
                    }
                default:
                    switch v22 {
                    case -816:
                        switch -2147483648 {
                        case 127:
                            v3 = false
                        case -9223372036854775808:
                            fmt.Println(v22, v3)
                        case 277:
                            var v151 uint32 = (v86 | v86)
                            _ = v151
                        default:
                            v152 := v5[("" + "")]
                            _ = v152
                        }
                    default:
                        fmt.Println(v2, v4, v86)
                    }
                }
                v22 += v22
            }
            _ = v86
        }
        _ = v66
    }
    for fi153 := 0; fi153 < 18; fi153++ {
        v3 = v3
        v154 := v5[(((("foo bar" + "a") + "") + "") + ("a" + ("hello" + "foo bar")))]
        _ = fi153
        _ = v154
    }
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
