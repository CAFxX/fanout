package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 uint8; F1 string }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int8; F1 string }
func genFn0() int64 {
    for fi4 := 0; fi4 < 10; fi4++ {
        if true {
            switch ("foo bar" + "\n\t") {
            case "hello":
                var v5 int8 = -128
                _ = v5
            case "foo bar":
                for fi6 := 0; fi6 < 6; fi6++ {
                    var v7 bool = true
                    _ = fi6
                    _ = v7
                }
            default:
                var v8 uint64 = 0
                _ = v8
            }
        }
        if false {
            fmt.Println(fi4)
            switch 1 {
            case -1:
                var v9 uint = 643
                _ = v9
            case 9223372036854775807:
                var v10 int16 = 32767
                _ = v10
            }
        } else {
            var v11 uint8 = 127
            var v12 int64 = 1
            _ = v11
            _ = v12
        }
        _ = fi4
    }
    var v13 bool = false
    _ = v13
    return 1
}
func genFn1() int8 {
    var v14 uint8 = 49
    v14 = (v14 ^ (v14 / (((v14 ^ v14)) | 1)))
    _ = v14
    return 18
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 941
    v3 := make(map[uint32]int)
    fmt.Println(v2)
    if false {
        v2 += (v2 % ((v2) | 1))
        v2 = 1
    } else {
        var v15 bool = (((0.1 <= (float64(0.1) + (((1e10 + float64(1.5)) - float64((1e10 * 1e10))) + 1.5))) || ((325 == -2147483648) && false)) && true)
        _ = v15
    }
    switch 57 {
    case 127:
        var v16 int8 = -66
        _ = v16
    case -9223372036854775808:
        switch "hello" {
        case "":
            switch -474 {
            case 32767:
                fmt.Println(v2)
            default:
                v2 += 643
            }
        case "a":
            v2 -= 0
        case "foo bar":
            fmt.Println(v2)
        }
    case 9223372036854775807:
        fmt.Println(v2)
    default:
        var v17 uint16 = 1
        _ = v17
    }
    if (18446744073709551615 < 65535) {
        var v18 int16 = 42
        for fi19 := 0; fi19 < 10; fi19++ {
            var v20 int64 = genFn0()
            fmt.Println(v2, v18)
            switch 32767 {
            case 0:
                var v21 uint32 = v2
                _ = v21
            case 128:
                v20 = v20
            case 127:
                v22 := GenS1{F0: 636}
                _ = v22
            default:
                var v23 uint8 = 106
                _ = v23
            }
            _ = fi19
            _ = v20
        }
        _ = v18
    } else {
        v2 = 0
        var v24 uint = 199
        _ = v24
    }
    if false {
        var v25 uint32 = (v2 / ((v2) | 1))
        _ = v25
    }
    var v26 float32 = 1.5
    v27 := []int{-239, 2147483647, -38, 2147483647}
    fmt.Println(v26)
    for fi28 := 0; fi28 < 8; fi28++ {
        delete(v3, (^(v2)))
        for fi29 := 0; fi29 < 9; fi29++ {
            var v30 int16 = int16(v2)
            v26 = v26
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    for fi31 := 0; fi31 < 6; fi31++ {
        switch fi31 {
        case -9223372036854775808:
            if (v26 != 0.0) {
                for fi32 := 0; fi32 < 17; fi32++ {
                    fmt.Println(fi32, v2)
                    var v33 bool = false
                    _ = fi32
                    _ = v33
                }
                for fi34 := 0; fi34 < 15; fi34++ {
                    v35 := GenS1{F0: 32767}
                    delete(v3, v2)
                    _ = fi34
                    _ = v35
                }
            } else {
                var v36 uint32 = v2
                v37 := GenS1{F0: 42}
                _ = v36
                _ = v37
            }
        case 0:
            if false {
                var v38 int = (fi31 * fi31)
                _ = v38
            } else {
                for fi39 := 0; fi39 < 6; fi39++ {
                    v40 := GenS2{F0: 255, F1: 32767}
                    switch fi39 {
                    case 127:
                        fmt.Println(v2, fi39)
                    case -2147483648:
                        var v41 int16 = 0
                        _ = v41
                    case 485:
                        var v42 uint8 = 77
                        _ = v42
                    default:
                        var v43 int32 = -753
                        _ = v43
                    }
                    switch (fi39 | (fi31 / ((-1) | 1))) {
                    case -9223372036854775808:
                        _ = len(v3)
                    case 4294967295:
                        if false {
                            var v44 uint8 = 2
                            _ = v44
                        }
                        _ = len(v27)
                        v2 -= v2
                    }
                    if true {
                        v2 = (v2 * (^(v2)))
                        var v45 int = (fi31 + fi31)
                        _ = v45
                    } else {
                        v26 += v26
                        if false {
                            switch ("a" + ("a" + "hello")) {
                            case "foo bar":
                                v2 = v2
                            }
                            var v46 uint64 = 18446744073709551615
                            v47 := GenS3{F0: genFn1()}
                            v48 := GenS4{F0: -2147483648}
                            _ = v46
                            _ = v47
                            _ = v48
                        }
                        v26 += v26
                    }
                    v2 = (v2 >> ((18446744073709551615) & 63))
                    _ = fi39
                    _ = v40
                }
                switch (fi31 | (fi31 - (-(fi31)))) {
                case 9223372036854775807:
                    for fi49 := 0; fi49 < 18; fi49++ {
                        if (9223372036854775807 != 0) {
                            var v50 int = 1
                            _ = len(v3)
                            if true {
                                if ((!false) || true) {
                                    v2 = 42
                                    v51 := GenS1{F0: 32767}
                                    _ = v51
                                } else {
                                    v52 := GenS1{F0: 85}
                                    _ = v52
                                }
                                switch "hello" {
                                case "\n\t":
                                    v53 := v3[2147483647]
                                    _ = v53
                                case "":
                                    v27 = append(v27, 9223372036854775807)
                                default:
                                    if true {
                                        v54 := GenS1{F0: 2}
                                        if ((v26 / v26) > 0.1) {
                                            switch ("\n\t" + ("hello" + ("hello" + "\n\t"))) {
                                            case "a":
                                                v55 := GenS5{F0: 65535, F1: 1}
                                                _ = v55
                                            case "foo bar":
                                                fmt.Println(fi49, v50)
                                            default:
                                                var v56 int32 = 144
                                                _ = v56
                                            }
                                        } else {
                                            if ((genFn0() > 42) && ("" > "\n\t")) {
                                                if true {
                                                    switch fi49 {
                                                    case -9223372036854775808:
                                                        switch "a" {
                                                        case "a":
                                                            if true {
                                                                v57 := GenS6{F0: 128, F1: ("\n\t" + (("foo bar" + "foo bar") + ""))}
                                                                _ = v57
                                                            } else {
                                                                v50 = fi49
                                                                fmt.Println(v2)
                                                            }
                                                        default:
                                                            var v58 float32 = v26
                                                            _ = v58
                                                        }
                                                    case 2147483647:
                                                        if false {
                                                            var v59 uint32 = 4294967295
                                                            fmt.Println(fi31, v59)
                                                            _ = v59
                                                        }
                                                    }
                                                    if len(v27) > 0 {
                                                        v27[0] = (^(fi49))
                                                    }
                                                }
                                            } else {
                                                fmt.Println(v2)
                                            }
                                            var v60 float64 = 1.5
                                            _ = v60
                                        }
                                        var v61 bool = true
                                        _ = v54
                                        _ = v61
                                    } else {
                                        v62 := GenS6{F0: 42, F1: "foo bar"}
                                        _ = v62
                                    }
                                }
                                var v63 int = 596
                                _ = v63
                            }
                            _ = v50
                        } else {
                            fmt.Println(fi49, v2)
                            var v64 int = 42
                            var v65 int16 = -32768
                            _ = v64
                            _ = v65
                        }
                        delete(v1, v2)
                        var v66 int = -338
                        _ = fi49
                        _ = v66
                    }
                default:
                    var v67 int16 = -32768
                    _ = v67
                }
            }
        default:
            fmt.Println(fi31)
        }
        var v68 uint16 = 65535
        if false {
            if true {
                fmt.Println(v26, v26)
                _ = len(v27)
                v2 += 2147483647
            } else {
                _ = len(v3)
                v2 = v2
            }
        } else {
            v2 = v2
        }
        _ = fi31
        _ = v68
    }
    var v69 uint = 42
    var v70 float32 = (float32(0.0) + (v26 * (((v26 + ((v26 * 0.0) - v26)) / ((v26 + (3.14159 / -2.25)) + ((v26 * 1e10) + v26))) * v26)))
    v71 := GenS4{F0: -1}
    var v72 uint64 = 365
    if true {
        switch 32767 {
        case -846:
            var v73 int64 = 42
            _ = v73
        case -2147483648:
            switch "" {
            case "":
                if ("" > "\n\t") {
                    for fi74 := 0; fi74 < 2; fi74++ {
                        switch (("" + (("" + "") + "foo bar")) + "a") {
                        case "hello":
                            switch 9223372036854775807 {
                            case 127:
                                if false {
                                    v69 += 42
                                    var v75 string = "\n\t"
                                    switch fi74 {
                                    case -1:
                                        var v76 uint16 = 42
                                        _ = v76
                                    case -2147483648:
                                        var v77 int16 = 32767
                                        _ = v77
                                    case 42:
                                        v26 -= v26
                                    }
                                    _ = v75
                                } else {
                                    var v78 int64 = -212
                                    if true {
                                        v72 = (v72 >> (((^(v69))) & 63))
                                        v70 += 1e10
                                    } else {
                                        if false {
                                            for fi79 := 0; fi79 < 9; fi79++ {
                                                v80 := GenS2{F0: uint8(v72), F1: (fi79 ^ 32767)}
                                                switch "\n\t" {
                                                case "foo bar":
                                                    v81 := GenS7{F0: fi74}
                                                    _ = v81
                                                case "\n\t":
                                                    v26 = ((v26 / ((float32((((1.5 / 1.0) / float64(1e10)) + ((0.1 * (1.5 / 0.1)) / 0.1))) + -2.25) * v70)) + (((0.1 / (v70 / v70)) * (v70 - v70)) / float32(-2.25)))
                                                default:
                                                    var v82 bool = (!false)
                                                    _ = v82
                                                }
                                                _ = fi79
                                                _ = v80
                                            }
                                            if len(v27) >= 2 {
                                                v27 = v27[1:len(v27)]
                                            }
                                            v78 = v78
                                        }
                                        v78 += -1
                                        var v83 string = "foo bar"
                                        _ = v83
                                    }
                                    v78 -= v78
                                    _ = v78
                                }
                            case -97:
                                v2 -= 207
                            case 851:
                                var v84 uint = v69
                                _ = v84
                            default:
                                for fi85 := 0; fi85 < 18; fi85++ {
                                    v27 = append(v27, 32768)
                                    var v86 bool = (0 <= 2147483647)
                                    v87 := v3[(v2 << (((v69 ^ (v69 % ((v69) | 1)))) & 63))]
                                    _ = fi85
                                    _ = v86
                                    _ = v87
                                }
                            }
                        case "a":
                            delete(v3, v2)
                        case "":
                            var v88 int16 = 32767
                            _ = v88
                        default:
                            if (!false) {
                                var v89 uint8 = 2
                                var v90 int = int(v72)
                                fmt.Println(v2, v69)
                                _ = v89
                                _ = v90
                            }
                        }
                        _ = fi74
                    }
                    switch 128 {
                    case 507:
                        var v91 uint32 = 88
                        _ = v91
                    case 931:
                        v92 := GenS4{F0: -2147483648}
                        _ = v92
                    }
                    var v93 int32 = 2147483647
                    _ = v93
                } else {
                    switch "\n\t" {
                    case "":
                        switch (("\n\t" + "foo bar") + (("" + "a") + "foo bar")) {
                        case "\n\t":
                            var v94 int8 = 127
                            _ = v94
                        case "a":
                            v69 += (v69 - (v69 % (((v69 * (v69 / ((v69) | 1)))) | 1)))
                        case "hello":
                            v27 = append(v27, 74)
                        default:
                            _ = len(v27)
                        }
                    }
                    _ = len(v3)
                }
            case "\n\t":
                for fi95 := 0; fi95 < 11; fi95++ {
                    if (v69 > (v69 >> ((42) & 63))) {
                        v2 = v2
                        for fi96 := 0; fi96 < 5; fi96++ {
                            v97 := GenS2{F0: 0, F1: (fi95 << (((v69 - (^(v69)))) & 63))}
                            _ = fi96
                            _ = v97
                        }
                        var v98 uint16 = 2
                        _ = v98
                    } else {
                        v99 := GenS4{F0: -2147483648}
                        v72 = (v72 + 0)
                        _ = v99
                    }
                    v2 += (v2 * 0)
                    switch "\n\t" {
                    case "foo bar":
                        v70 = 1.5
                    default:
                        for fi100 := 0; fi100 < 5; fi100++ {
                            v2 = 632
                            if true {
                                fmt.Println(fi95)
                                v101 := GenS4{F0: 926}
                                _ = v101
                            } else {
                                fmt.Println(fi100, v26)
                            }
                            fmt.Println(v70)
                            delete(v1, 65535)
                            fmt.Println(v26)
                            _ = fi100
                        }
                    }
                    _ = fi95
                }
            case "a":
                if false {
                    v27 = append(v27, -932)
                    if true {
                        v70 -= v70
                    } else {
                        switch -2147483648 {
                        case 413:
                            v102 := GenS8{F0: v69}
                            _ = v102
                        case -551:
                            v103 := GenS8{F0: v69}
                            _ = v103
                        case 128:
                            for fi104 := 0; fi104 < 14; fi104++ {
                                var v105 uint32 = (v2 % ((2147483647) | 1))
                                if (42 < 42) {
                                    switch "foo bar" {
                                    case "":
                                        for fi106 := 0; fi106 < 2; fi106++ {
                                            v105 = v105
                                            _ = fi106
                                        }
                                    case "foo bar":
                                        v69 = 255
                                    case "\n\t":
                                        for fi107 := 0; fi107 < 3; fi107++ {
                                            v108 := GenS7{F0: fi107}
                                            _ = fi107
                                            _ = v108
                                        }
                                    default:
                                        fmt.Println(v69, v70, v26)
                                    }
                                    v109 := GenS4{F0: 336}
                                    _ = v109
                                }
                                v26 -= v26
                                _ = fi104
                                _ = v105
                            }
                        default:
                            var v110 int = 128
                            _ = v110
                        }
                        var v111 float32 = (v70 / float32(0.1))
                        var v112 int64 = genFn0()
                        _ = v111
                        _ = v112
                    }
                    if true {
                        var v113 string = ("\n\t" + ("" + "a"))
                        if (true && ((!false) && ((-286 < -553) && (0.0 < v70)))) {
                            v2 += (v2 + 65535)
                            for fi114 := 0; fi114 < 6; fi114++ {
                                switch 0 {
                                case -2147483648:
                                    if len(v27) >= 2 {
                                        v27 = v27[1:len(v27)]
                                    }
                                }
                                v113 = ""
                                _ = fi114
                            }
                        } else {
                            v115 := GenS9{F0: 42}
                            switch "" {
                            case "a":
                                var v116 uint8 = 1
                                _ = v116
                            case "foo bar":
                                if (2 < 1) {
                                    _ = len(v27)
                                    fmt.Println(v72, v2)
                                    fmt.Println(v113, v113)
                                } else {
                                    v1[2147483647] = 32768
                                    v27 = append(v27, 9223372036854775807)
                                    for fi117 := 0; fi117 < 20; fi117++ {
                                        if (-1 != 84) {
                                            for fi118 := 0; fi118 < 1; fi118++ {
                                                switch (fi118 << (((^(v69))) & 63)) {
                                                case 42:
                                                    if true {
                                                        switch 127 {
                                                        case -1:
                                                            v2 += 2
                                                        }
                                                    } else {
                                                        var v119 string = v113
                                                        v120 := v1[4294967295]
                                                        if ((fi117 % ((fi117) | 1)) >= (fi118 & fi118)) {
                                                            v120 = 32768
                                                            var v121 float32 = (v26 + 0.0)
                                                            v26 += v121
                                                            _ = v121
                                                        } else {
                                                            var v122 int = fi117
                                                            v3[v2] = fi118
                                                            var v123 int = fi118
                                                            _ = v122
                                                            _ = v123
                                                        }
                                                        _ = v119
                                                        _ = v120
                                                    }
                                                default:
                                                    v69 = (v69 >> ((65535) & 63))
                                                }
                                                _ = fi118
                                            }
                                        } else {
                                            var v124 string = (("\n\t" + v113) + ("hello" + "foo bar"))
                                            var v125 bool = (!(int64(fi117) != -1))
                                            _ = v124
                                            _ = v125
                                        }
                                        for fi126 := 0; fi126 < 10; fi126++ {
                                            v69 += (v69 - (v69 % (((v69 * (v69 / ((65535) | 1)))) | 1)))
                                            _ = fi126
                                        }
                                        _ = fi117
                                    }
                                }
                            default:
                                for fi127 := 0; fi127 < 15; fi127++ {
                                    v69 += 42
                                    var v128 float32 = 1e10
                                    if false {
                                        var v129 bool = (false || (!(100 <= 0)))
                                        switch (v113 + v113) {
                                        case "\n\t":
                                            fmt.Println(v129)
                                        }
                                        v128 -= ((v70 / (v128 / float32((1.5 - (-2.25 * (float64(1e10) / ((((((0.1 - -2.25) * -2.25) * 1.5) * 0.0) / 1.0) - (float64((1.5 * (v70 - v128))) / 1.5)))))))) / (v70 / 1.0))
                                        _ = v129
                                    } else {
                                        var v130 int32 = 626
                                        var v131 int = (fi127 * (-(fi127)))
                                        v132 := GenS10{F0: 15, F1: v113}
                                        _ = v130
                                        _ = v131
                                        _ = v132
                                    }
                                    _ = fi127
                                    _ = v128
                                }
                            }
                            v2 += v2
                            _ = v115
                        }
                        if true {
                            if true {
                                v2 = (v2 ^ v2)
                                v133 := GenS7{F0: -768}
                                _ = v133
                            } else {
                                var v134 uint16 = 255
                                fmt.Println(v69, v69, v2)
                                fmt.Println(v26, v113, v69)
                                _ = v134
                            }
                        } else {
                            v2 += (v2 * v2)
                            var v135 int16 = 32767
                            var v136 uint8 = 0
                            _ = v135
                            _ = v136
                        }
                        _ = v113
                    }
                } else {
                    v70 = v26
                    v137 := GenS9{F0: 100}
                    _ = v137
                }
            }
        case 964:
            if false {
                v70 += (0.0 + (3.14159 / v70))
            } else {
                fmt.Println(v2, v70)
                if false {
                    for fi138 := 0; fi138 < 14; fi138++ {
                        fmt.Println(v2)
                        fmt.Println(fi138, v72)
                        _ = fi138
                    }
                }
                delete(v1, v2)
            }
        }
        if false {
            var v139 bool = true
            var v140 int16 = 1
            _ = v139
            _ = v140
        } else {
            var v141 uint16 = 499
            _ = v141
        }
        v72 -= v72
    }
    v2 = (^(v2))
    v142 := GenS3{F0: genFn1()}
    for fi143 := 0; fi143 < 9; fi143++ {
        if (false && true) {
            v144 := GenS6{F0: 2, F1: ""}
            _ = v144
        }
        var v145 uint8 = 68
        _ = fi143
        _ = v145
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v142
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v69
    _ = v70
    _ = v71
    _ = v72
}
