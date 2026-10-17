package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 int; F2 string }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 int; F1 int16 }
type GenS6 struct { F0 int8; F1 string; F2 bool }
type GenS7 struct { F0 int8; F1 int32; F2 int8 }
type GenS8 struct { F0 uint32; F1 uint8 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 string; F1 uint16 }
type GenS11 struct { F0 uint32; F1 int16 }
type GenS12 struct { F0 uint32; F1 uint16 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 int32; F1 string; F2 int32 }
type GenS15 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS16 struct { F0 uint; F1 uint; F2 uint }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 int8 }

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{}
    var v3 uint16 = 1
    var v4 uint32 = 301
    v5 := []int{2147483647, 32768, 670, 2147483647}
    var v6 float32 = float32(1.5)
    for fi7 := 0; fi7 < 4; fi7++ {
        switch ("hello" + "hello") {
        case "foo bar":
            var v8 int32 = 521
            _ = v8
        }
        var v9 int16 = 1000
        _ = fi7
        _ = v9
    }
    var v10 float64 = float64((v6 * (v6 + float32(-2.25))))
    if (v3 < (^(v3))) {
        for fi11 := 0; fi11 < 16; fi11++ {
            if false {
                if true {
                    var v12 int64 = 0
                    _ = v12
                } else {
                    _ = len(v2)
                    v13 := v1[32767]
                    var v14 bool = (v13 || v13)
                    _ = v13
                    _ = v14
                }
            } else {
                if false {
                    switch fi11 {
                    case 273:
                        var v15 int32 = 862
                        _ = v15
                    case 128:
                        for fi16 := 0; fi16 < 14; fi16++ {
                            if true {
                                var v17 int32 = 1
                                if false {
                                    switch (^(fi16)) {
                                    case 832:
                                        v18 := GenS1{F0: (v3 + 2), F1: 255, F2: fi11}
                                        _ = v18
                                    default:
                                        var v19 float64 = ((v10 / (v10 - (float64((v6 - (v6 / 1.0))) + ((v10 - v10) - v10)))) * 0.1)
                                        _ = v19
                                    }
                                    var v20 float32 = v6
                                    _ = len(v5)
                                    _ = v20
                                }
                                fmt.Println(v4, fi11)
                                v21 := v1[fi11]
                                v3 -= (v3 + v3)
                                var v22 uint64 = uint64(v4)
                                if true {
                                    if (!v21) {
                                        if (128 <= 139) {
                                            fmt.Println(v4, v3)
                                            if v21 {
                                                _ = len(v1)
                                                fmt.Println(v22, v21, v21)
                                            }
                                            fmt.Println(v21, v17, v17)
                                        } else {
                                            if v21 {
                                                v23 := GenS2{F0: -9223372036854775808}
                                                if len(v2) > 0 {
                                                    v2[0] = 0
                                                }
                                                switch ("foo bar" + "a") {
                                                case "\n\t":
                                                    fmt.Println(fi11, v4, v10)
                                                case "a":
                                                    if ((false || v21) && v21) {
                                                        switch "foo bar" {
                                                        case "hello":
                                                            v24 := v1[int(v17)]
                                                            _ = v24
                                                        }
                                                        v21 = v21
                                                    }
                                                case "hello":
                                                    var v25 uint = uint(v22)
                                                    _ = v25
                                                }
                                                _ = v23
                                            } else {
                                                v22 += (v22 % (((v22 >> ((999) & 63))) | 1))
                                                var v26 int = (fi16 | fi16)
                                                _ = v26
                                            }
                                            var v27 uint64 = (^(v22))
                                            switch "" {
                                            case "\n\t":
                                                _ = len(v1)
                                            case "a":
                                                v28 := GenS1{F0: v3, F1: 42, F2: (fi11 * -177)}
                                                _ = v28
                                            default:
                                                fmt.Println(v6, v10)
                                                var v29 uint16 = v3
                                                switch (fi11 + (fi16 / (((fi16 & fi16)) | 1))) {
                                                case 32767:
                                                    var v30 int8 = 42
                                                    _ = v30
                                                case 2147483647:
                                                    v31 := GenS1{F0: v29, F1: 127, F2: (fi11 % ((128) | 1))}
                                                    _ = v31
                                                }
                                                switch (fi11 + (-(fi11))) {
                                                case -732:
                                                    var v32 uint32 = v4
                                                    _ = v32
                                                case -2147483648:
                                                    switch (("hello" + "hello") + ("foo bar" + ("foo bar" + ("\n\t" + "")))) {
                                                    case "hello":
                                                        v6 += v6
                                                    case "a":
                                                        var v33 uint64 = v27
                                                        _ = v33
                                                    }
                                                default:
                                                    v34 := GenS2{F0: -9223372036854775808}
                                                    _ = v34
                                                }
                                                _ = v29
                                            }
                                            fmt.Println(v22, v22, v3)
                                            v4 = v4
                                            _ = v27
                                        }
                                        var v35 int = 9223372036854775807
                                        _ = v35
                                    }
                                    v4 = 4294967295
                                } else {
                                    fmt.Println(fi16, v6, v10)
                                }
                                v36 := GenS2{F0: -1}
                                v10 = (0.1 - 0.0)
                                fmt.Println(fi11)
                                _ = v17
                                _ = v21
                                _ = v22
                                _ = v36
                            }
                            fmt.Println(v10, v6, v3)
                            _ = fi16
                        }
                    default:
                        var v37 uint32 = (v4 ^ (^(v4)))
                        _ = v37
                    }
                    for fi38 := 0; fi38 < 1; fi38++ {
                        var v39 uint8 = 2
                        var v40 int = fi11
                        v41 := GenS2{F0: -754}
                        if false {
                            if (false && true) {
                                var v42 string = "\n\t"
                                fmt.Println(v6, v3)
                                fmt.Println(v4, fi11, v40)
                                _ = v42
                            } else {
                                if true {
                                    v43 := GenS3{F0: 4294967295, F1: (fi11 * fi38), F2: "\n\t"}
                                    _ = len(v1)
                                    _ = v43
                                } else {
                                    v1[(^(v40))] = false
                                    var v44 uint16 = (^(v3))
                                    var v45 float64 = 3.14159
                                    _ = v44
                                    _ = v45
                                }
                                v39 = (^(v39))
                                var v46 int32 = -2147483648
                                _ = v46
                            }
                            _ = len(v1)
                            if (float32(1.5) > v6) {
                                var v47 float32 = v6
                                delete(v1, (fi11 ^ v40))
                                var v48 bool = (false || (!((false || (v40 > (fi38 / ((-1) | 1)))) && (!true))))
                                _ = v47
                                _ = v48
                            }
                        }
                        if true {
                            var v49 uint = 0
                            v6 += v6
                            var v50 int16 = 42
                            _ = v49
                            _ = v50
                        }
                        _ = fi38
                        _ = v39
                        _ = v40
                        _ = v41
                    }
                    var v51 string = "\n\t"
                    _ = v51
                }
                v1[fi11] = ((0 >= 2147483647) || (42 == 745))
            }
            if (100 >= 100) {
                if (1 <= -2147483648) {
                    for fi52 := 0; fi52 < 10; fi52++ {
                        v53 := GenS4{F0: 42, F1: -1}
                        if (1 < 0) {
                            delete(v1, fi11)
                            var v54 int16 = -32768
                            var v55 uint16 = 181
                            _ = v54
                            _ = v55
                        } else {
                            var v56 uint16 = (v3 + 32767)
                            var v57 uint32 = (v4 / ((v4) | 1))
                            if false {
                                var v58 bool = (fi52 < 784)
                                _ = len(v5)
                                _ = v58
                            }
                            fmt.Println(v10, fi11)
                            switch "\n\t" {
                            case "a":
                                v59 := GenS5{F0: (fi11 << ((0) & 63)), F1: 32767}
                                _ = v59
                            case "\n\t":
                                v60 := GenS5{F0: fi52, F1: 32767}
                                _ = v60
                            case "hello":
                                if true {
                                    var v61 int = 619
                                    v62 := GenS6{F0: 73, F1: "a", F2: false}
                                    delete(v1, (fi52 * v61))
                                    _ = v61
                                    _ = v62
                                } else {
                                    v63 := GenS5{F0: 42, F1: int16(fi52)}
                                    v6 = ((v6 + -2.25) - (((v6 - (((v6 * v6) / v6) * ((0.1 - 1.5) - float32(v10)))) + float32(v10)) / v6))
                                    _ = v63
                                }
                            default:
                                v64 := GenS1{F0: 255, F1: 162, F2: 32767}
                                _ = v64
                            }
                            _ = v56
                            _ = v57
                        }
                        switch (^(fi52)) {
                        case 4294967295:
                            fmt.Println(fi11, fi52)
                        case 42:
                            v5 = append(v5, (-(fi52)))
                            if false {
                                _ = len(v2)
                            } else {
                                var v65 uint16 = (v3 | (^(v3)))
                                fmt.Println(v3, fi11, v10)
                                _ = v65
                            }
                        case 32767:
                            fmt.Println(v10, fi11, v4)
                        }
                        v66 := GenS1{F0: 32767, F1: 255, F2: 128}
                        _ = fi52
                        _ = v53
                        _ = v66
                    }
                } else {
                    for fi67 := 0; fi67 < 15; fi67++ {
                        v68 := GenS7{F0: 127, F1: -1, F2: -17}
                        if ("foo bar" == ("\n\t" + (("" + "hello") + (("hello" + "\n\t") + "hello")))) {
                            if true {
                                var v69 uint = 4294967295
                                _ = v69
                            } else {
                                v70 := GenS6{F0: -128, F1: "a", F2: true}
                                _ = v70
                            }
                        }
                        if false {
                            var v71 uint32 = v4
                            fmt.Println(v10, fi67)
                            switch "hello" {
                            case "a":
                                v72 := GenS8{F0: 1, F1: 2}
                                _ = v72
                            case "\n\t":
                                if (-1 != -115) {
                                    fmt.Println(v71, v4)
                                    v6 += v6
                                } else {
                                    if (false || (true && (false && false))) {
                                        var v73 uint16 = (v3 | v3)
                                        if false {
                                            var v74 int = (fi11 ^ -1)
                                            switch "\n\t" {
                                            case "":
                                                v1[fi11] = ((^(v3)) == v3)
                                            }
                                            switch (^(fi67)) {
                                            case 368:
                                                v73 -= 167
                                            case 1:
                                                var v75 float64 = v10
                                                _ = v75
                                            default:
                                                fmt.Println(v4)
                                            }
                                            var v76 int64 = 1
                                            v4 = (v71 >> ((2) & 63))
                                            _ = v74
                                            _ = v76
                                        }
                                        switch fi11 {
                                        case 0:
                                            var v77 int64 = 1
                                            _ = v77
                                        default:
                                            switch "hello" {
                                            case "\n\t":
                                                switch 2147483647 {
                                                case -1:
                                                    if true {
                                                        if len(v5) > 0 {
                                                            v5[0] = fi67
                                                        }
                                                        var v78 string = "hello"
                                                        var v79 int32 = int32(v73)
                                                        _ = v78
                                                        _ = v79
                                                    } else {
                                                        if true {
                                                            v3 = 356
                                                            fmt.Println(v3, v73, fi11)
                                                            var v80 bool = (!true)
                                                            _ = v80
                                                        } else {
                                                            var v81 int16 = 1000
                                                            _ = v81
                                                        }
                                                    }
                                                case 966:
                                                    var v82 uint64 = 4294967295
                                                    _ = v82
                                                default:
                                                    v83 := GenS7{F0: 42, F1: -2147483648, F2: int8(v4)}
                                                    _ = v83
                                                }
                                            case "":
                                                v84 := v1[(fi67 * 269)]
                                                _ = v84
                                            case "a":
                                                v73 += (v73 << ((4294967295) & 63))
                                            default:
                                                v85 := GenS7{F0: -1, F1: 283, F2: -95}
                                                _ = v85
                                            }
                                        }
                                        _ = v73
                                    }
                                    if (!((false && (true || (2147483647 <= 0))) && (float64((v6 - v6)) != v10))) {
                                        var v86 uint32 = (v71 * 42)
                                        v87 := GenS7{F0: 100, F1: -472, F2: 83}
                                        v86 += (^(v71))
                                        _ = v86
                                        _ = v87
                                    }
                                    v88 := GenS9{F0: 9223372036854775807}
                                    v6 = 1.5
                                    _ = v88
                                }
                            }
                            _ = v71
                        } else {
                            fmt.Println(fi11, fi11)
                            if (-9223372036854775808 == 0) {
                                v6 = float32((v10 * v10))
                            } else {
                                v3 = v3
                            }
                        }
                        _ = fi67
                        _ = v68
                    }
                    var v89 float32 = 1.5
                    if true {
                        fmt.Println(fi11)
                        v90 := GenS2{F0: 537}
                        var v91 bool = (799 != 2)
                        _ = v90
                        _ = v91
                    }
                    _ = v89
                }
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi11
        }
        v92 := GenS5{F0: -704, F1: 1}
        for fi93 := 0; fi93 < 20; fi93++ {
            var v94 int = 2147483647
            v4 = 356
            _ = fi93
            _ = v94
        }
        _ = v92
    } else {
        for fi95 := 0; fi95 < 20; fi95++ {
            v6 = ((1e10 + v6) * (0.1 * v6))
            delete(v1, fi95)
            _ = fi95
        }
        for fi96 := 0; fi96 < 11; fi96++ {
            if true {
                for fi97 := 0; fi97 < 2; fi97++ {
                    v10 = v10
                    _ = fi97
                }
                var v98 uint32 = (v4 | (v4 * (v4 >> ((2) & 63))))
                v5 = append(v5, fi96)
                _ = v98
            } else {
                v10 += v10
                var v99 uint8 = 0
                switch fi96 {
                case 9223372036854775807:
                    delete(v1, (-(fi96)))
                case -633:
                    v100 := GenS10{F0: "", F1: 1}
                    _ = v100
                case 0:
                    var v101 float32 = v6
                    _ = v101
                }
                _ = v99
            }
            _ = len(v1)
            if true {
                v6 = 3.14159
                var v102 float32 = float32((v10 - v10))
                _ = v102
            }
            _ = fi96
        }
    }
    for fi103 := 0; fi103 < 17; fi103++ {
        if false {
            delete(v1, (fi103 ^ fi103))
            for fi104 := 0; fi104 < 7; fi104++ {
                if false {
                    v6 = v6
                } else {
                    var v105 uint16 = (v3 + v3)
                    _ = v105
                }
                v106 := GenS5{F0: (fi104 * fi103), F1: 0}
                v107 := GenS8{F0: v4, F1: 128}
                _ = fi104
                _ = v106
                _ = v107
            }
            var v108 int16 = 32767
            _ = v108
        } else {
            if false {
                v6 = (1.5 + v6)
                switch (("a" + "hello") + "") {
                case "a":
                    fmt.Println(v10)
                case "\n\t":
                    for fi109 := 0; fi109 < 13; fi109++ {
                        v110 := GenS4{F0: 42, F1: -1}
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi109
                        _ = v110
                    }
                default:
                    if true {
                        if true {
                            v6 = v6
                        }
                        switch "" {
                        case "":
                            v3 = (v3 | v3)
                        }
                    }
                }
                fmt.Println(v4)
            }
            if len(v5) > 0 {
                v5[0] = (fi103 - (-(fi103)))
            }
        }
        _ = fi103
    }
    var v111 uint16 = 32767
    v4 = 73
    fmt.Println(v10)
    for fi112 := 0; fi112 < 7; fi112++ {
        switch "\n\t" {
        case "a":
            var v113 int16 = 32767
            _ = v113
        case "hello":
            for fi114 := 0; fi114 < 14; fi114++ {
                var v115 uint64 = 74
                _ = fi114
                _ = v115
            }
        case "":
            if true {
                var v116 uint32 = v4
                _ = v116
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v117 int64 = 0
                _ = v117
            }
        default:
            v118 := GenS6{F0: 127, F1: "hello", F2: true}
            _ = v118
        }
        for fi119 := 0; fi119 < 11; fi119++ {
            switch -887 {
            case -785:
                if (v3 <= v3) {
                    v10 = v10
                }
            case 32768:
                var v120 float64 = 3.14159
                _ = v120
            default:
                fmt.Println(fi112, v3)
            }
            switch (fi112 * fi112) {
            case -407:
                if (false || true) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if (-484 <= 0) {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        var v121 int64 = -1
                        fmt.Println(v4, v111)
                        _ = v121
                    } else {
                        v10 += 0.1
                        fmt.Println(v10, v111)
                    }
                }
            case -9223372036854775808:
                if ((v4 - (v4 >> ((935) & 63))) < (v4 << ((255) & 63))) {
                    v4 += (v4 + (v4 - v4))
                }
            case 4294967295:
                fmt.Println(v3, fi119, v10)
            default:
                _ = len(v1)
            }
            v122 := GenS11{F0: v4, F1: 1000}
            v123 := GenS12{F0: v4, F1: v3}
            v124 := GenS1{F0: (v111 * (v111 | 2)), F1: 255, F2: fi112}
            switch (fi112 & fi119) {
            case -9223372036854775808:
                var v125 int8 = -50
                _ = v125
            case 0:
                v4 = v4
            }
            _ = fi119
            _ = v122
            _ = v123
            _ = v124
        }
        _ = fi112
    }
    v126 := GenS13{F0: 7}
    v127 := GenS14{F0: 0, F1: ("hello" + "\n\t"), F2: 1}
    var v128 uint32 = 63
    v4 -= v128
    if (445 == 981) {
        v111 -= (v111 % ((1) | 1))
        fmt.Println(v128, v6)
    }
    var v129 uint32 = 207
    for fi130 := 0; fi130 < 15; fi130++ {
        if false {
            v111 = 908
            switch (fi130 << ((65535) & 63)) {
            case 32768:
                if true {
                    v131 := GenS15{F0: 128, F1: 1, F2: 127}
                    var v132 uint = 154
                    _ = v131
                    _ = v132
                } else {
                    v1[fi130] = true
                    v133 := GenS16{F0: 415, F1: 0, F2: 831}
                    v111 -= (^(v111))
                    _ = v133
                }
            case -1:
                _ = len(v5)
            case 4294967295:
                v134 := v1[-369]
                _ = v134
            }
            v135 := GenS17{F0: v3}
            _ = v135
        }
        v136 := GenS18{F0: -72}
        v2 = append(v2, v128)
        _ = fi130
        _ = v136
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v111:", v111)
    fmt.Println("v128:", v128)
    fmt.Println("v129:", v129)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v111
    _ = v126
    _ = v127
    _ = v128
    _ = v129
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
