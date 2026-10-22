package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16; F1 int16; F2 int16 }
type GenS6 struct { F0 uint64; F1 int16 }
type GenS7 struct { F0 int64; F1 int8; F2 bool }
type GenS8 struct { F0 string; F1 int8; F2 int32 }
type GenS9 struct { F0 bool; F1 uint32; F2 int16 }
type GenS10 struct { F0 int8; F1 bool; F2 uint }
type GenS11 struct { F0 int16; F1 uint8 }
type GenS12 struct { F0 int64; F1 uint8; F2 uint16 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 int16; F1 uint8 }
type GenS15 struct { F0 bool }

func main() {
    v1 := make(map[string]int)
    var v2 float64 = 0.0
    var v3 bool = false
    var v4 string = "\n\t"
    v5 := make(map[uint32]bool)
    if v3 {
        if (true && ((!true) || (4294967295 > 1))) {
            if (634 <= 32767) {
                var v6 int32 = -2147483648
                var v7 int8 = -58
                for fi8 := 0; fi8 < 10; fi8++ {
                    var v9 float64 = (1.5 + v2)
                    for fi10 := 0; fi10 < 14; fi10++ {
                        v3 = v3
                        if (v3 || (v3 && v3)) {
                            switch v4 {
                            case "\n\t":
                                v2 = v2
                            case "foo bar":
                                v7 = 25
                            case "":
                                v11 := []int{(fi10 % ((32767) | 1)), 9223372036854775807}
                                _ = v11
                            default:
                                v9 -= ((float64(1e10) / v2) * 1e10)
                            }
                            if true {
                                var v12 float32 = ((1e10 / -2.25) - (3.14159 - 3.14159))
                                if (0 < 42) {
                                    if v3 {
                                        fmt.Println(v7, v2)
                                    } else {
                                        v13 := GenS1{F0: (true && (fi8 <= fi10)), F1: 2}
                                        fmt.Println(fi8, v9, v9)
                                        var v14 int = 544
                                        _ = v13
                                        _ = v14
                                    }
                                    v15 := GenS2{F0: 65535}
                                    _ = v15
                                }
                                v2 = 3.14159
                                if ((0 == 2) && v3) {
                                    v16 := GenS2{F0: 133}
                                    switch ("a" + "a") {
                                    case "hello":
                                        if (1 <= 0) {
                                            fmt.Println(fi10)
                                            v2 -= v9
                                            v17 := []int{fi8, 9223372036854775807, 128}
                                            _ = v17
                                        } else {
                                            if (v3 && (0 != 1)) {
                                                v18 := GenS3{F0: 1, F1: 377}
                                                v3 = v3
                                                v19 := GenS3{F0: 2, F1: 1}
                                                _ = v18
                                                _ = v19
                                            }
                                            switch v4 {
                                            case "hello":
                                                v6 = 876
                                            case "foo bar":
                                                v12 += v12
                                            }
                                        }
                                    case "":
                                        if v3 {
                                            v3 = v3
                                        } else {
                                            var v20 uint16 = 2
                                            _ = v20
                                        }
                                    case "\n\t":
                                        fmt.Println(v2, v2, v3)
                                    }
                                    if v3 {
                                        v2 = v2
                                        if (true && (!v3)) {
                                            var v21 float32 = (v12 * float32((((v2 / (1e10 * v2)) + 1.5) * v2)))
                                            fmt.Println(v12, fi8, v4)
                                            if (!(42 != 0)) {
                                                v22 := GenS2{F0: 0}
                                                if v3 {
                                                    switch -9223372036854775808 {
                                                    case 128:
                                                        var v23 int64 = -820
                                                        _ = v23
                                                    }
                                                    v24 := GenS3{F0: 1, F1: 1}
                                                    if (!(2 <= 558)) {
                                                        fmt.Println(v6, v4, fi8)
                                                        v12 += (v21 + 1.5)
                                                        v3 = v3
                                                    } else {
                                                        fmt.Println(fi8, v21)
                                                        v25 := []int{(fi8 * fi8)}
                                                        _ = v25
                                                    }
                                                    _ = v24
                                                }
                                                v26 := GenS2{F0: 65535}
                                                _ = v22
                                                _ = v26
                                            } else {
                                                v4 = "\n\t"
                                            }
                                            _ = v21
                                        } else {
                                            v27 := GenS2{F0: 255}
                                            v2 = float64(1.5)
                                            v2 += v9
                                            _ = v27
                                        }
                                    }
                                    _ = v16
                                } else {
                                    v2 = (v2 - 0.1)
                                    fmt.Println(v12, fi8, v7)
                                }
                                v28 := GenS2{F0: 2}
                                _ = v12
                                _ = v28
                            } else {
                                fmt.Println(v7, v6)
                            }
                            switch -848 {
                            case 9223372036854775807:
                                v29 := v5[1]
                                _ = v29
                            case 2147483647:
                                v30 := GenS3{F0: 18446744073709551615, F1: 248}
                                _ = v30
                            default:
                                if false {
                                    var v31 int32 = (v6 | (v6 - (v6 | v6)))
                                    var v32 uint64 = 489
                                    v33 := GenS4{F0: 185}
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                } else {
                                    v6 -= (v6 + (v6 / ((int32(fi10)) | 1)))
                                    switch (v4 + v4) {
                                    case "":
                                        if (((2 < 0) || (false || v3)) && (v4 >= v4)) {
                                            var v34 float64 = v2
                                            _ = v34
                                        } else {
                                            v2 -= 1e10
                                            if v3 {
                                                switch "\n\t" {
                                                case "\n\t":
                                                    v35 := GenS4{F0: 42}
                                                    _ = v35
                                                }
                                                v6 += (v6 << ((65535) & 63))
                                            } else {
                                                var v36 bool = (false || v3)
                                                var v37 uint64 = 2
                                                v38 := []int{}
                                                _ = v36
                                                _ = v37
                                                _ = v38
                                            }
                                            v39 := []int{fi8, fi8, (fi10 & (fi10 + (fi8 >> ((255) & 63)))), 634}
                                            _ = v39
                                        }
                                    }
                                    v40 := []int{626, fi10}
                                    _ = v40
                                }
                            }
                            v41 := GenS5{F0: 42, F1: 0, F2: -32768}
                            _ = v41
                        }
                        if (390 <= 1) {
                            var v42 int32 = 962
                            var v43 uint64 = 1
                            _ = v42
                            _ = v43
                        }
                        v6 = -129
                        v44 := []int{}
                        v6 = v6
                        _ = fi10
                        _ = v44
                    }
                    if false {
                        v7 = 100
                    } else {
                        v1[(("" + (v4 + ("foo bar" + (((v4 + v4) + v4) + (((("" + v4) + "\n\t") + "a") + "a"))))) + "")] = 0
                        v9 += v2
                    }
                    _ = fi8
                    _ = v9
                }
                _ = v6
                _ = v7
            } else {
                var v45 uint = 42
                v46 := GenS2{F0: 1}
                var v47 uint16 = 2
                _ = v45
                _ = v46
                _ = v47
            }
            for fi48 := 0; fi48 < 15; fi48++ {
                for fi49 := 0; fi49 < 7; fi49++ {
                    v3 = true
                    fmt.Println(fi48, v2, fi48)
                    _ = fi49
                }
                var v50 int64 = 398
                _ = fi48
                _ = v50
            }
        }
        if v3 {
            v3 = (42 < -2147483648)
            if (9223372036854775807 != -9223372036854775808) {
                for fi51 := 0; fi51 < 8; fi51++ {
                    v4 = (v4 + "\n\t")
                    for fi52 := 0; fi52 < 12; fi52++ {
                        v53 := v1[v4]
                        _ = fi52
                        _ = v53
                    }
                    _ = fi51
                }
                v4 = v4
                var v54 uint16 = 65535
                _ = v54
            }
            v3 = v3
        } else {
            var v55 float64 = (v2 / v2)
            var v56 float32 = 3.14159
            _ = v55
            _ = v56
        }
        for fi57 := 0; fi57 < 19; fi57++ {
            fmt.Println(v2)
            var v58 int32 = -2147483648
            _ = fi57
            _ = v58
        }
    }
    if true {
        v3 = false
    } else {
        v59 := GenS6{F0: 2, F1: 42}
        if (!true) {
            v2 -= v2
            v60 := []int{1, 42}
            if (!(32767 > 2)) {
                if v3 {
                    if len(v60) > 0 {
                        v60[0] = 174
                    }
                    var v61 float32 = 1.5
                    for fi62 := 0; fi62 < 17; fi62++ {
                        v63 := GenS4{F0: 2}
                        _ = fi62
                        _ = v63
                    }
                    _ = v61
                } else {
                    var v64 int8 = 100
                    _ = v64
                }
                v60 = append(v60, 42)
                switch "" {
                case "foo bar":
                    _ = len(v60)
                }
            }
            _ = v60
        } else {
            for fi65 := 0; fi65 < 17; fi65++ {
                switch (fi65 & (fi65 << ((65535) & 63))) {
                case 128:
                    v66 := GenS7{F0: 9223372036854775807, F1: -128, F2: false}
                    _ = v66
                case 4294967295:
                    var v67 int = fi65
                    _ = v67
                case 32768:
                    fmt.Println(v3)
                default:
                    if v3 {
                        delete(v1, "\n\t")
                        v2 = (3.14159 - (((float64(3.14159) + v2) / float64((((1.5 * (-2.25 * (1e10 + 1.5))) - (1.5 * 0.1)) - 1e10))) + 0.1))
                    } else {
                        for fi68 := 0; fi68 < 15; fi68++ {
                            if false {
                                var v69 uint = 199
                                var v70 int64 = -9223372036854775808
                                v4 = v4
                                _ = v69
                                _ = v70
                            } else {
                                var v71 string = v4
                                v71 = ("foo bar" + ("a" + ""))
                                _ = v71
                            }
                            if (42 != 1) {
                                if v3 {
                                    v72 := GenS2{F0: 255}
                                    var v73 int64 = 0
                                    if (1 > 1) {
                                        v3 = false
                                    }
                                    _ = len(v1)
                                    _ = v72
                                    _ = v73
                                }
                                _ = len(v1)
                                v74 := GenS8{F0: (v4 + v4), F1: 127, F2: 2147483647}
                                _ = v74
                            } else {
                                var v75 int = (fi65 & 4294967295)
                                _ = v75
                            }
                            _ = fi68
                        }
                    }
                }
                fmt.Println(v3)
                for fi76 := 0; fi76 < 17; fi76++ {
                    var v77 uint32 = 284
                    _ = fi76
                    _ = v77
                }
                _ = fi65
            }
            v78 := GenS8{F0: (v4 + "a"), F1: 1, F2: 0}
            _ = v78
        }
        for fi79 := 0; fi79 < 15; fi79++ {
            var v80 int16 = -754
            _ = fi79
            _ = v80
        }
        _ = v59
    }
    for fi81 := 0; fi81 < 14; fi81++ {
        v3 = ((v2 - v2) < 0.1)
        _ = fi81
    }
    v3 = true
    var v82 int = 351
    v2 = (float64((0.0 * float32(v2))) * ((v2 * (float64(0.1) - v2)) * v2))
    if ((!v3) || ((v3 || (!(true || (904 == 203)))) && v3)) {
        if false {
            var v83 bool = v3
            switch (v82 ^ v82) {
            case 32768:
                fmt.Println(v3)
            case 32767:
                v82 -= (v82 << ((0) & 63))
            case -9223372036854775808:
                if v83 {
                    var v84 uint = 42
                    var v85 float32 = 0.1
                    v86 := GenS9{F0: true, F1: 1, F2: 1000}
                    _ = v84
                    _ = v85
                    _ = v86
                }
            default:
                for fi87 := 0; fi87 < 16; fi87++ {
                    for fi88 := 0; fi88 < 9; fi88++ {
                        var v89 float32 = 1.5
                        var v90 int64 = -1
                        _ = fi88
                        _ = v89
                        _ = v90
                    }
                    switch v82 {
                    case 515:
                        switch ((v4 + v4) + "foo bar") {
                        case "foo bar":
                            switch fi87 {
                            case 32768:
                                var v91 uint16 = 65535
                                _ = v91
                            default:
                                v92 := GenS10{F0: 0, F1: ((fi87 % ((v82) | 1)) < 560), F2: 1}
                                _ = v92
                            }
                        }
                    default:
                        v82 = v82
                    }
                    var v93 float64 = (1.5 + v2)
                    _ = fi87
                    _ = v93
                }
            }
            v1[v4] = v82
            _ = v83
        } else {
            var v94 uint32 = 665
            _ = v94
        }
    } else {
        if (v2 > (((1e10 + 1e10) - 3.14159) + v2)) {
            v95 := GenS7{F0: 9223372036854775807, F1: -1, F2: true}
            switch "" {
            case "\n\t":
                for fi96 := 0; fi96 < 20; fi96++ {
                    for fi97 := 0; fi97 < 2; fi97++ {
                        switch 4294967295 {
                        case 32768:
                            switch v4 {
                            case "":
                                if (2 < 4) {
                                    if (42 > 489) {
                                        switch v4 {
                                        case "\n\t":
                                            v5[2] = (31 != 42)
                                        case "":
                                            v82 -= fi96
                                        case "hello":
                                            var v98 uint8 = 78
                                            _ = v98
                                        }
                                        v99 := []int{(fi97 | (v82 / (((fi96 >> ((0) & 63))) | 1))), 1, fi97, fi97}
                                        fmt.Println(v2, fi97, fi96)
                                        switch (("foo bar" + "a") + v4) {
                                        case "foo bar":
                                            var v100 float64 = 1.5
                                            _ = v100
                                        case "a":
                                            _ = len(v1)
                                        case "\n\t":
                                            v4 = v4
                                        default:
                                            v2 = (0.1 * v2)
                                        }
                                        var v101 int32 = 2147483647
                                        _ = v101
                                        _ = v99
                                    }
                                    if v3 {
                                        if true {
                                            v102 := GenS9{F0: v3, F1: 1, F2: -32768}
                                            switch fi97 {
                                            case 4294967295:
                                                var v103 int = 32768
                                                _ = v103
                                            case -1:
                                                v104 := GenS9{F0: (!(42 != 4294967295)), F1: 4294967295, F2: 1000}
                                                _ = v104
                                            case -821:
                                                switch (v82 << ((528) & 63)) {
                                                case 4294967295:
                                                    v105 := []int{(v82 << ((1) & 63)), 32767, v82}
                                                    _ = v105
                                                case 1:
                                                    v3 = v3
                                                }
                                            }
                                            v3 = v3
                                            _ = v102
                                        } else {
                                            if v3 {
                                                delete(v5, 915)
                                                if true {
                                                    var v106 int16 = 20
                                                    switch fi96 {
                                                    case 1:
                                                        var v107 int = (fi96 % ((fi96) | 1))
                                                        v4 = ""
                                                        v4 = "\n\t"
                                                        v108 := GenS1{F0: (!v3), F1: 14}
                                                        var v109 int8 = 105
                                                        _ = v107
                                                        _ = v108
                                                        _ = v109
                                                    }
                                                    var v110 bool = (!((!((((v4 + (v4 + "a")) + v4) + "") > "\n\t")) || (false || (!v3))))
                                                    _ = v106
                                                    _ = v110
                                                } else {
                                                    v2 -= ((1.5 / (v2 + v2)) / 1.0)
                                                }
                                                if (9223372036854775807 >= 939) {
                                                    v111 := []int{v82, (fi97 >> ((18446744073709551615) & 63))}
                                                    _ = v111
                                                }
                                            }
                                        }
                                    }
                                    var v112 float32 = (1e10 - float32(-2.25))
                                    fmt.Println(fi97, v3, fi97)
                                    fmt.Println(v4)
                                    v113 := GenS11{F0: -1, F1: 1}
                                    _ = v112
                                    _ = v113
                                }
                                v114 := GenS8{F0: "foo bar", F1: 0, F2: 277}
                                var v115 string = "\n\t"
                                _ = v114
                                _ = v115
                            default:
                                var v116 bool = (66 > 65535)
                                _ = v116
                            }
                        case 908:
                            var v117 int = 42
                            _ = v117
                        case -1:
                            v82 += (fi97 >> ((56) & 63))
                        }
                        if false {
                            v4 = ("\n\t" + v4)
                        } else {
                            v118 := GenS11{F0: 812, F1: 0}
                            _ = v118
                        }
                        var v119 bool = (32767 >= 65535)
                        v2 -= v2
                        var v120 int16 = 1
                        _ = len(v1)
                        _ = fi97
                        _ = v119
                        _ = v120
                    }
                    v121 := GenS1{F0: (516 != 42), F1: 242}
                    var v122 uint8 = 128
                    _ = fi96
                    _ = v121
                    _ = v122
                }
            case "foo bar":
                v4 = ((v4 + (v4 + v4)) + "hello")
            }
            _ = v95
        }
        for fi123 := 0; fi123 < 7; fi123++ {
            v3 = v3
            var v124 string = ("hello" + v4)
            _ = fi123
            _ = v124
        }
        if (true || (-851 > 0)) {
            v125 := []int{v82, v82, -447}
            if (375 != 42) {
                if v3 {
                    v82 += 9223372036854775807
                    v4 = v4
                }
                fmt.Println(v4, v82, v4)
                v4 = v4
            } else {
                v4 = (v4 + v4)
                var v126 bool = (!true)
                _ = v126
            }
            fmt.Println(v3)
            _ = v125
        } else {
            for fi127 := 0; fi127 < 14; fi127++ {
                switch (v82 >> ((447) & 63)) {
                case 0:
                    v128 := GenS12{F0: -9223372036854775808, F1: 0, F2: 0}
                    _ = v128
                case -2147483648:
                    v3 = (3.14159 >= 0.1)
                case 1:
                    var v129 int16 = 0
                    _ = v129
                default:
                    v130 := GenS9{F0: (0.1 <= -2.25), F1: 65535, F2: 1000}
                    _ = v130
                }
                v131 := []int{v82, (v82 << ((4294967295) & 63)), 0}
                if (!v3) {
                    var v132 int = (v82 / (((v82 ^ fi127)) | 1))
                    v132 = 2147483647
                    _ = len(v5)
                    _ = v132
                } else {
                    v133 := GenS10{F0: -128, F1: (!v3), F2: 0}
                    _ = v133
                }
                _ = fi127
                _ = v131
            }
        }
    }
    var v134 uint64 = 42
    if v3 {
        fmt.Println(v82, v4)
        for fi135 := 0; fi135 < 15; fi135++ {
            _ = len(v5)
            if (357 <= 42) {
                fmt.Println(fi135)
            }
            _ = fi135
        }
        var v136 int64 = -9223372036854775808
        _ = v136
    } else {
        switch (v82 >> ((194) & 63)) {
        case 53:
            var v137 uint16 = 0
            _ = v137
        case 9223372036854775807:
            var v138 int = 111
            _ = v138
        case -9223372036854775808:
            switch (v4 + v4) {
            case "":
                var v139 int = (v82 & v82)
                _ = v139
            case "hello":
                v140 := GenS13{F0: 2147483648}
                _ = v140
            default:
                v141 := GenS10{F0: -128, F1: (!v3), F2: 158}
                _ = v141
            }
        }
    }
    v4 = "foo bar"
    if v3 {
        for fi142 := 0; fi142 < 7; fi142++ {
            for fi143 := 0; fi143 < 1; fi143++ {
                v134 -= v134
                if (4294967295 != 2) {
                    fmt.Println(v3)
                    v82 += v82
                    v144 := GenS14{F0: -839, F1: 59}
                    _ = v144
                }
                v145 := []int{0, (fi142 / (((fi143 ^ (^(fi142)))) | 1))}
                v3 = (100 != -1)
                _ = fi143
                _ = v145
            }
            _ = fi142
        }
        if (-341 != 0) {
            v146 := GenS15{F0: (464 == 42)}
            if true {
                v3 = true
            }
            fmt.Println(v4, v134)
            _ = v146
        } else {
            v82 = (v82 | v82)
            v82 = v82
        }
        v4 = (v4 + ("\n\t" + v4))
    } else {
        for fi147 := 0; fi147 < 6; fi147++ {
            var v148 uint32 = 907
            _ = len(v5)
            _ = fi147
            _ = v148
        }
        var v149 bool = v3
        _ = v149
    }
    for fi150 := 0; fi150 < 18; fi150++ {
        if false {
            fmt.Println(v2, v134, fi150)
            fmt.Println(v3)
            v151 := []int{v82}
            _ = v151
        }
        for fi152 := 0; fi152 < 13; fi152++ {
            v3 = v3
            _ = fi152
        }
        _ = fi150
    }
    v3 = (1 != 2147483647)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v134:", v134)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v82:", v82)
    _ = v1
    _ = v134
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v82
}
