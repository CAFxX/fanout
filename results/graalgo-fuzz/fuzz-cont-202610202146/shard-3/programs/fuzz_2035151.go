package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 bool }
type GenS2 struct { F0 int64; F1 int8; F2 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 uint16; F1 uint16 }
type GenS7 struct { F0 int8; F1 int32; F2 uint }
type GenS8 struct { F0 int8; F1 uint32 }
type GenS9 struct { F0 uint8; F1 uint32 }
type GenS10 struct { F0 int8; F1 uint64 }
type GenS11 struct { F0 uint; F1 int8; F2 int32 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 uint32 }
func genFn0() uint8 {
    v5 := GenS1{F0: 127, F1: 255, F2: true}
    v6 := GenS1{F0: 177, F1: 42, F2: true}
    _ = v5
    _ = v6
    return 128
}

func main() {
    var v1 float32 = 0.1
    v2 := make(map[string]uint32)
    v3 := []uint32{}
    var v4 int64 = -1
    var v7 uint32 = 4294967295
    fmt.Println(v1, v4, v1)
    if false {
        v7 = v7
        fmt.Println(v4, v7, v7)
        if false {
            v4 += v4
            v8 := GenS2{F0: (-(v4)), F1: -1, F2: 127}
            _ = v8
        } else {
            switch 128 {
            case 127:
                var v9 int64 = v4
                _ = v9
            case -2147483648:
                var v10 int64 = -9223372036854775808
                _ = v10
            default:
                if ((("\n\t" + ("" + "foo bar")) + "a") <= "") {
                    fmt.Println(v7, v4, v7)
                    v3 = append(v3, (v7 + v7))
                    switch "a" {
                    case "foo bar":
                        var v11 int16 = -1
                        _ = v11
                    case "":
                        v12 := GenS1{F0: 128, F1: 44, F2: true}
                        _ = v12
                    default:
                        for fi13 := 0; fi13 < 18; fi13++ {
                            v14 := GenS2{F0: v4, F1: -61, F2: -1}
                            _ = fi13
                            _ = v14
                        }
                    }
                }
            }
            v15 := GenS2{F0: v4, F1: 54, F2: 100}
            var v16 int32 = 592
            _ = v15
            _ = v16
        }
    }
    _ = len(v2)
    for fi17 := 0; fi17 < 7; fi17++ {
        for fi18 := 0; fi18 < 2; fi18++ {
            switch "hello" {
            case "":
                switch fi17 {
                case 2147483647:
                    var v19 int8 = 42
                    _ = v19
                default:
                    var v20 int8 = 127
                    _ = v20
                }
                if true {
                    v3 = append(v3, v7)
                } else {
                    if (((false || (!((fi17 + fi18) != fi18))) && (false && ((!false) && (2 > 42)))) || (!(false && (0.1 == 3.14159)))) {
                        v21 := GenS2{F0: v4, F1: 1, F2: 42}
                        v4 += (v4 << ((42) & 63))
                        v4 -= -1
                        switch 42 {
                        case 42:
                            if false {
                                var v22 uint64 = 970
                                if len(v3) > 0 {
                                    v3[0] = (v7 & v7)
                                }
                                _ = v22
                            }
                        case 2147483647:
                            var v23 uint16 = 32767
                            _ = v23
                        case 904:
                            var v24 uint8 = 144
                            _ = v24
                        }
                        _ = v21
                    } else {
                        if (((v4 * (v4 % ((v4) | 1))) != (-(v4))) || (true && ((v4 | (v4 & (v4 / ((v4) | 1)))) > v4))) {
                            if (((fi18 * -1) >= fi18) && (!false)) {
                                if len(v3) > 0 {
                                    v3[0] = (v7 >> ((313) & 63))
                                }
                                if (!(true && false)) {
                                    v2["\n\t"] = v7
                                    var v25 uint32 = v7
                                    if (-2.25 > 1.5) {
                                        fmt.Println(fi17)
                                        var v26 uint16 = 2
                                        fmt.Println(fi17, v7, v4)
                                        _ = v26
                                    }
                                    var v27 float64 = 0.0
                                    var v28 int = fi17
                                    _ = v25
                                    _ = v27
                                    _ = v28
                                } else {
                                    if (-2147483648 > -411) {
                                        var v29 int32 = -1
                                        var v30 uint8 = 1
                                        switch "foo bar" {
                                        case "a":
                                            v31 := GenS2{F0: v4, F1: 100, F2: 0}
                                            _ = v31
                                        case "hello":
                                            v32 := GenS3{F0: v4}
                                            _ = v32
                                        default:
                                            var v33 int64 = (v4 >> ((268) & 63))
                                            _ = v33
                                        }
                                        _ = v29
                                        _ = v30
                                    }
                                }
                            }
                            switch fi17 {
                            case 4294967295:
                                var v34 int32 = -2147483648
                                _ = v34
                            case 9223372036854775807:
                                var v35 float64 = (0.0 * 1e10)
                                _ = v35
                            default:
                                v36 := GenS4{F0: 19}
                                _ = v36
                            }
                            if true {
                                var v37 int16 = 32767
                                v4 += (v4 - (v4 ^ (-(v4))))
                                _ = v37
                            }
                        } else {
                            if true {
                                v38 := GenS5{F0: 127, F1: 42}
                                if false {
                                    fmt.Println(fi17)
                                    _ = len(v3)
                                    switch -421 {
                                    case 48:
                                        if true {
                                            v1 = (float32((3.14159 - -2.25)) * v1)
                                            v2["\n\t"] = 42
                                            fmt.Println(v4, v7)
                                        } else {
                                            v1 = v1
                                            fmt.Println(fi18, fi18)
                                            v39 := GenS2{F0: (v4 % (((v4 / (((v4 / (((v4 / ((v4) | 1))) | 1))) | 1))) | 1)), F1: -75, F2: 100}
                                            fmt.Println(fi18)
                                            _ = v39
                                        }
                                    default:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                } else {
                                    if (65535 != 18446744073709551615) {
                                        var v40 uint16 = 421
                                        _ = v40
                                    } else {
                                        v4 = -628
                                    }
                                    var v41 uint32 = v7
                                    var v42 uint32 = v41
                                    _ = v41
                                    _ = v42
                                }
                                _ = v38
                            } else {
                                var v43 int32 = 2147483647
                                _ = v43
                            }
                            if (true && true) {
                                if (1 != (v4 | v4)) {
                                    var v44 uint16 = 70
                                    _ = len(v2)
                                    _ = v44
                                } else {
                                    var v45 string = "foo bar"
                                    fmt.Println(fi18)
                                    _ = v45
                                }
                            } else {
                                var v46 int8 = 81
                                v7 -= v7
                                switch fi18 {
                                case -176:
                                    var v47 uint64 = 0
                                    _ = v47
                                case -166:
                                    v46 += v46
                                case 4294967295:
                                    switch ("foo bar" + "foo bar") {
                                    case "foo bar":
                                        delete(v2, (("" + ("" + "\n\t")) + "a"))
                                    case "\n\t":
                                        v7 += 523
                                    default:
                                        v48 := v2[""]
                                        _ = v48
                                    }
                                default:
                                    if true {
                                        v4 = (v4 | v4)
                                        v49 := GenS6{F0: 2, F1: 1}
                                        _ = v49
                                    } else {
                                        if ((false && (v1 == 0.1)) && true) {
                                            var v50 uint8 = genFn0()
                                            var v51 int32 = 2147483647
                                            v50 += 158
                                            var v52 float64 = ((1e10 * 0.1) / 1e10)
                                            v46 = (v46 ^ 0)
                                            _ = v50
                                            _ = v51
                                            _ = v52
                                        } else {
                                            v1 += ((((((v1 / (1e10 - (-2.25 * (v1 - v1)))) / (v1 / 1e10)) / ((v1 + v1) / float32((1.5 + ((float64(v1) + (float64(1.5) - (0.1 / ((1.5 + 0.1) - 1e10)))) - 3.14159))))) * (0.0 / v1)) / -2.25) + (v1 * v1))
                                            v46 += v46
                                            var v53 int64 = v4
                                            var v54 uint16 = 255
                                            _ = v53
                                            _ = v54
                                        }
                                        var v55 bool = false
                                        var v56 uint64 = 42
                                        _ = v55
                                        _ = v56
                                    }
                                }
                                _ = v46
                            }
                            v57 := v2[(("foo bar" + "\n\t") + "")]
                            _ = v57
                        }
                        v58 := GenS6{F0: 1, F1: 42}
                        switch fi18 {
                        case 32767:
                            var v59 uint16 = 41
                            _ = v59
                        case -1:
                            v60 := GenS1{F0: 42, F1: 199, F2: true}
                            _ = v60
                        default:
                            switch (fi18 & fi18) {
                            case 2147483647:
                                var v61 bool = ((!false) && false)
                                _ = v61
                            default:
                                switch fi18 {
                                case 127:
                                    v1 = v1
                                default:
                                    switch (fi17 % (((fi18 + fi17)) | 1)) {
                                    case 127:
                                        v62 := GenS2{F0: 42, F1: int8(fi17), F2: 100}
                                        _ = v62
                                    case 0:
                                        if false {
                                            switch "a" {
                                            case "":
                                                switch fi18 {
                                                case 127:
                                                    switch "" {
                                                    case "foo bar":
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    case "":
                                                        var v63 uint8 = 127
                                                        _ = v63
                                                    case "a":
                                                        v1 += (v1 + v1)
                                                    }
                                                case -192:
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                            case "a":
                                                fmt.Println(fi17, v4, v4)
                                                var v64 uint64 = 21
                                                _ = v64
                                            default:
                                                _ = len(v3)
                                            }
                                        }
                                    case -345:
                                        var v65 uint = 255
                                        _ = v65
                                    default:
                                        var v66 uint64 = 1
                                        _ = v66
                                    }
                                }
                            }
                        }
                        _ = v58
                    }
                    if (!(true && false)) {
                        v1 = float32(0.0)
                    } else {
                        v7 = v7
                        var v67 float32 = 1e10
                        v68 := GenS5{F0: 1, F1: 255}
                        v1 = v1
                        if (2 <= 128) {
                            v69 := GenS7{F0: -21, F1: -2147483648, F2: 1}
                            v70 := GenS1{F0: 42, F1: 42, F2: true}
                            _ = v69
                            _ = v70
                        } else {
                            v1 = 3.14159
                            v71 := GenS4{F0: 862}
                            if (32767 >= fi17) {
                                v72 := v2[""]
                                if true {
                                    v7 = (v7 ^ (v72 ^ v7))
                                    v4 -= (v4 << ((18446744073709551615) & 63))
                                } else {
                                    _ = len(v3)
                                    v4 = 1
                                    v73 := GenS6{F0: 489, F1: 496}
                                    _ = v73
                                }
                                fmt.Println(v4, fi18)
                                _ = v72
                            }
                            _ = v71
                        }
                        _ = v67
                        _ = v68
                    }
                }
            case "hello":
                v74 := GenS1{F0: 2, F1: 0, F2: true}
                _ = v74
            case "\n\t":
                v4 += -377
            default:
                var v75 int8 = 1
                _ = v75
            }
            _ = fi18
        }
        var v76 string = "a"
        v1 += float32(1.5)
        _ = fi17
        _ = v76
    }
    _ = len(v3)
    if ((-(v4)) >= v4) {
        fmt.Println(v7, v1, v4)
        v77 := v2[("a" + "foo bar")]
        _ = v77
    } else {
        for fi78 := 0; fi78 < 4; fi78++ {
            if (32767 >= -1) {
                if true {
                    v7 += v7
                    v79 := GenS7{F0: 47, F1: int32(fi78), F2: 441}
                    _ = v79
                }
            }
            _ = fi78
        }
        switch -9223372036854775808 {
        case 128:
            var v80 int32 = 863
            _ = v80
        case 2147483647:
            var v81 int8 = -1
            _ = v81
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (18446744073709551615 == 4294967295) {
        switch int(v7) {
        case 1:
            for fi82 := 0; fi82 < 2; fi82++ {
                fmt.Println(v1, v7)
                _ = fi82
            }
        case -147:
            var v83 uint8 = 128
            _ = v83
        case 32768:
            var v84 bool = true
            _ = v84
        }
    } else {
        for fi85 := 0; fi85 < 14; fi85++ {
            if (42 != 426) {
                var v86 uint64 = 42
                v4 = 862
                _ = v86
            }
            if (!true) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                fmt.Println(v4, v4, fi85)
                for fi87 := 0; fi87 < 6; fi87++ {
                    switch 42 {
                    case 127:
                        var v88 float64 = float64(((v1 + v1) / v1))
                        _ = v88
                    case -327:
                        fmt.Println(v7)
                    case -2147483648:
                        fmt.Println(fi87, v1, fi87)
                    }
                    fmt.Println(v1, fi85)
                    if (v4 <= 9223372036854775807) {
                        v89 := GenS8{F0: 100, F1: v7}
                        _ = v89
                    }
                    _ = fi87
                }
            } else {
                if false {
                    switch (fi85 - fi85) {
                    case 9223372036854775807:
                        for fi90 := 0; fi90 < 1; fi90++ {
                            var v91 int8 = 127
                            switch "hello" {
                            case "a":
                                v92 := GenS4{F0: 42}
                                _ = v92
                            }
                            _ = fi90
                            _ = v91
                        }
                    case -494:
                        var v93 int32 = -1
                        _ = v93
                    case -1:
                        fmt.Println(v4, v1, fi85)
                    }
                    var v94 string = "hello"
                    for fi95 := 0; fi95 < 9; fi95++ {
                        delete(v2, "hello")
                        var v96 int = (fi95 * 9223372036854775807)
                        switch (v96 - -230) {
                        case 4294967295:
                            var v97 int32 = -1
                            _ = v97
                        default:
                            var v98 int32 = 695
                            var v99 uint8 = 207
                            switch -106 {
                            case 756:
                                switch v94 {
                                case "":
                                    var v100 float64 = 1e10
                                    _ = v100
                                default:
                                    if true {
                                        var v101 bool = (2 != 42)
                                        v102 := GenS1{F0: genFn0(), F1: v99, F2: v101}
                                        var v103 uint16 = 32767
                                        switch "foo bar" {
                                        case "hello":
                                            v7 = (^(v7))
                                        }
                                        if ((v101 || v101) && ((v101 && v101) && ((v4 & v4) <= v4))) {
                                            delete(v2, (("" + "hello") + v94))
                                        }
                                        _ = v101
                                        _ = v102
                                        _ = v103
                                    }
                                }
                            case 2147483647:
                                var v104 uint8 = (^(v99))
                                _ = v104
                            case 32768:
                                v105 := GenS2{F0: v4, F1: 1, F2: 127}
                                _ = v105
                            default:
                                var v106 float64 = 3.14159
                                _ = v106
                            }
                            fmt.Println(v96)
                            fmt.Println(v94, v7, v1)
                            v1 -= -2.25
                            var v107 int16 = -317
                            var v108 uint32 = v7
                            _ = v107
                            _ = v108
                            _ = v98
                            _ = v99
                        }
                        _ = fi95
                        _ = v96
                    }
                    _ = v94
                } else {
                    var v109 int = fi85
                    v110 := GenS8{F0: -58, F1: v7}
                    for fi111 := 0; fi111 < 7; fi111++ {
                        var v112 uint8 = 132
                        _ = fi111
                        _ = v112
                    }
                    _ = v109
                    _ = v110
                }
                v1 -= v1
                v113 := GenS9{F0: 1, F1: (^(v7))}
                _ = v113
            }
            var v114 string = "\n\t"
            _ = fi85
            _ = v114
        }
        _ = len(v2)
        v7 = v7
    }
    switch 127 {
    case -1:
        if (!(32767 > -391)) {
            for fi115 := 0; fi115 < 16; fi115++ {
                var v116 uint = 42
                v3 = append(v3, (^(v7)))
                for fi117 := 0; fi117 < 3; fi117++ {
                    switch fi115 {
                    case 32768:
                        var v118 uint16 = 42
                        _ = v118
                    default:
                        if (((1e10 / 1.0) - (-2.25 - 1.5)) == (3.14159 * -2.25)) {
                            v119 := GenS10{F0: 14, F1: 42}
                            _ = v119
                        }
                    }
                    _ = len(v2)
                    _ = fi117
                }
                _ = fi115
                _ = v116
            }
            v1 += (v1 - v1)
            fmt.Println(v7, v1, v7)
        } else {
            if (726 != v7) {
                v1 = ((-2.25 + v1) / 0.1)
                v4 += v4
            } else {
                var v120 uint64 = 4294967295
                if (177 > 1) {
                    fmt.Println(v7, v1)
                    v121 := GenS10{F0: 49, F1: (v120 ^ (v120 << ((65535) & 63)))}
                    _ = v121
                }
                _ = v120
            }
            var v122 int = -9223372036854775808
            for fi123 := 0; fi123 < 7; fi123++ {
                if (1 <= 0) {
                    v122 = (v122 >> ((4294967295) & 63))
                    var v124 uint64 = 0
                    _ = v124
                } else {
                    var v125 bool = false
                    if false {
                        v4 -= 0
                        v126 := v2["foo bar"]
                        _ = v126
                    }
                    var v127 uint32 = v7
                    _ = v125
                    _ = v127
                }
                v128 := v2[("\n\t" + "\n\t")]
                var v129 float64 = 1e10
                _ = fi123
                _ = v128
                _ = v129
            }
            _ = v122
        }
    case 363:
        for fi130 := 0; fi130 < 16; fi130++ {
            v131 := GenS11{F0: 4294967295, F1: -43, F2: 42}
            var v132 float64 = 0.1
            _ = fi130
            _ = v131
            _ = v132
        }
    }
    for fi133 := 0; fi133 < 6; fi133++ {
        for fi134 := 0; fi134 < 14; fi134++ {
            if (!true) {
                v7 -= 65535
                var v135 bool = false
                _ = v135
            } else {
                if false {
                    v136 := GenS12{F0: v4}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if true {
                        switch "foo bar" {
                        case "":
                            if false {
                                var v137 int16 = -1
                                _ = v137
                            }
                        }
                        v7 = (v7 % (((^(v7))) | 1))
                    } else {
                        v4 += v4
                        fmt.Println(fi133, v1)
                        fmt.Println(v4)
                        if (!true) {
                            fmt.Println(fi134, v1, v1)
                        }
                        v138 := GenS11{F0: 0, F1: -88, F2: 0}
                        _ = v138
                    }
                    _ = v136
                } else {
                    v139 := GenS10{F0: 0, F1: 42}
                    v140 := GenS4{F0: 255}
                    v7 = v7
                    _ = v139
                    _ = v140
                }
                v4 += 9223372036854775807
                if (2 == 0) {
                    v7 = 2147483648
                } else {
                    v141 := GenS4{F0: 65535}
                    var v142 string = "hello"
                    _ = v141
                    _ = v142
                }
                v1 = ((v1 / float32(1e10)) / 3.14159)
                var v143 uint32 = 42
                _ = v143
            }
            _ = fi134
        }
        for fi144 := 0; fi144 < 10; fi144++ {
            switch 32768 {
            case 128:
                var v145 int8 = 0
                _ = v145
            case 1:
                var v146 uint8 = genFn0()
                _ = v146
            default:
                switch ("\n\t" + ("foo bar" + (("" + "") + "\n\t"))) {
                case "hello":
                    v7 += v7
                case "a":
                    if false {
                        v147 := GenS13{F0: 1}
                        var v148 uint = 18446744073709551615
                        _ = v147
                        _ = v148
                    } else {
                        v149 := GenS14{F0: v7}
                        _ = v149
                    }
                }
            }
            _ = fi144
        }
        _ = fi133
    }
    if (false || (255 <= 859)) {
        v1 += v1
        var v150 int8 = 0
        _ = v150
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
