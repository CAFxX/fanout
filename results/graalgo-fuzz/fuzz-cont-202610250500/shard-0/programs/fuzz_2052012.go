package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 int16; F2 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16; F1 int16; F2 uint64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int8; F1 uint8; F2 int64 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint32; F1 bool }
type GenS11 struct { F0 uint32 }

func main() {
    var v1 bool = ((65535 == 134) || false)
    v2 := make(map[uint32]int)
    v3 := []int{-63, 128}
    v4 := make(map[int]string)
    v1 = v1
    var v5 int32 = 432
    if v1 {
        var v6 int32 = v5
        var v7 int8 = 42
        _ = v6
        _ = v7
    }
    if false {
        if v1 {
            var v8 int8 = -1
            _ = v8
        }
        v3 = append(v3, 822)
        fmt.Println(v5)
    } else {
        var v9 uint16 = 2
        v10 := GenS1{F0: 1, F1: -9223372036854775808}
        _ = v10
        _ = v9
    }
    var v11 bool = v1
    var v12 uint8 = 190
    v13 := v4[127]
    v1 = v1
    v14 := GenS1{F0: 47, F1: 0}
    var v15 uint8 = (v12 % ((v12) | 1))
    if (v11 && (-2.25 <= 1.5)) {
        if v11 {
            for fi16 := 0; fi16 < 13; fi16++ {
                var v17 int32 = (v5 & v5)
                _ = len(v2)
                if (694 <= v5) {
                    v18 := GenS2{F0: (v17 << ((18446744073709551615) & 63))}
                    var v19 uint16 = 255
                    v19 = (v19 - (v19 >> ((42) & 63)))
                    _ = v18
                    _ = v19
                } else {
                    if (!v1) {
                        if v11 {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if false {
                                if v11 {
                                    v20 := GenS2{F0: 0}
                                    _ = v20
                                } else {
                                    var v21 uint8 = v12
                                    if (109 > 417) {
                                        for fi22 := 0; fi22 < 11; fi22++ {
                                            v1 = false
                                            fmt.Println(fi22, v15)
                                            _ = fi22
                                        }
                                        for fi23 := 0; fi23 < 16; fi23++ {
                                            if (42 != 112) {
                                                var v24 uint8 = (v21 | v12)
                                                v21 = v24
                                                var v25 uint8 = 42
                                                _ = v24
                                                _ = v25
                                            } else {
                                                var v26 uint32 = 376
                                                _ = v26
                                            }
                                            var v27 uint16 = 1
                                            var v28 uint8 = v12
                                            _ = fi23
                                            _ = v27
                                            _ = v28
                                        }
                                    }
                                    _ = v21
                                }
                            }
                        }
                        var v29 float64 = 3.14159
                        _ = v29
                    }
                }
                _ = fi16
                _ = v17
            }
            switch "\n\t" {
            case "hello":
                for fi30 := 0; fi30 < 2; fi30++ {
                    if ((0 == 0) && v1) {
                        for fi31 := 0; fi31 < 4; fi31++ {
                            v32 := GenS2{F0: (^(v5))}
                            _ = fi31
                            _ = v32
                        }
                        for fi33 := 0; fi33 < 4; fi33++ {
                            if true {
                                fmt.Println(v5, v12)
                                fmt.Println(v5)
                                if true {
                                    if (v11 && (false || v11)) {
                                        v13 = ""
                                        if true {
                                            v1 = ((fi33 - fi30) != (fi30 * fi30))
                                        }
                                    } else {
                                        v34 := GenS2{F0: v5}
                                        _ = v34
                                    }
                                    var v35 int16 = 1000
                                    var v36 string = (((v13 + (v13 + "hello")) + v13) + v13)
                                    v37 := GenS3{F0: 564, F1: 42, F2: (^(v12))}
                                    v12 = (^(v12))
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                } else {
                                    v5 -= (-(v5))
                                    v38 := GenS4{F0: 0}
                                    if (!v11) {
                                        v4[9223372036854775807] = v13
                                        v39 := GenS4{F0: 0}
                                        _ = len(v4)
                                        _ = v39
                                    }
                                    v40 := GenS5{F0: v11}
                                    _ = v38
                                    _ = v40
                                }
                            } else {
                                delete(v2, 42)
                                switch fi33 {
                                case 896:
                                    if v1 {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if (220 == 4) {
                                            v13 = "foo bar"
                                            if (true && ((0.1 >= 1.5) && (0 != 1))) {
                                                switch ((v13 + v13) + v13) {
                                                case "\n\t":
                                                    v13 = v13
                                                case "hello":
                                                    v12 = 128
                                                case "":
                                                    if v1 {
                                                        v1 = false
                                                        v11 = v1
                                                        if len(v3) > 0 {
                                                            v3[0] = -75
                                                        }
                                                    }
                                                    v41 := GenS2{F0: (-(v5))}
                                                    var v42 int16 = int16(v15)
                                                    _ = v41
                                                    _ = v42
                                                default:
                                                    if v1 {
                                                        fmt.Println(v12, v1, fi33)
                                                        v43 := GenS3{F0: 65535, F1: -1, F2: (^(v15))}
                                                        _ = v43
                                                    } else {
                                                        var v44 uint8 = (v15 >> ((18446744073709551615) & 63))
                                                        _ = v44
                                                    }
                                                }
                                            } else {
                                                v45 := GenS5{F0: (!v1)}
                                                _ = v45
                                            }
                                        } else {
                                            if v11 {
                                                if v1 {
                                                    v46 := GenS3{F0: 0, F1: 1, F2: 0}
                                                    var v47 float64 = 3.14159
                                                    _ = v46
                                                    _ = v47
                                                }
                                            } else {
                                                v13 = (("hello" + "") + v13)
                                            }
                                            v13 = v13
                                            switch (v13 + ("a" + v13)) {
                                            case "\n\t":
                                                switch fi33 {
                                                case 433:
                                                    var v48 bool = (((0.0 / 1.0) < (0.0 - 1e10)) || true)
                                                    _ = v48
                                                case -532:
                                                    var v49 float64 = -2.25
                                                    _ = v49
                                                case -9223372036854775808:
                                                    if len(v3) > 0 {
                                                        v3[0] = (fi30 | -2147483648)
                                                    }
                                                    var v50 string = ((v13 + v13) + v13)
                                                    _ = v50
                                                default:
                                                    v12 = v12
                                                }
                                            }
                                        }
                                        var v51 float32 = 1.5
                                        _ = v51
                                    }
                                case 379:
                                    if (!(1 < 1)) {
                                        v52 := GenS1{F0: 91, F1: -609}
                                        v12 -= (^(v15))
                                        _ = v52
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = (fi30 / ((0) | 1))
                                    }
                                    v5 += v5
                                    switch fi30 {
                                    case 32767:
                                        if v11 {
                                            if (11 > 50) {
                                                var v53 int16 = 1000
                                                if (!true) {
                                                    v54 := GenS6{F0: v53, F1: 757, F2: 2}
                                                    var v55 int64 = 96
                                                    var v56 int8 = 100
                                                    fmt.Println(v11)
                                                    v55 -= v55
                                                    _ = v54
                                                    _ = v55
                                                    _ = v56
                                                } else {
                                                    if (2147483648 > 2147483647) {
                                                        var v57 uint = 106
                                                        var v58 int8 = 127
                                                        v59 := GenS3{F0: v57, F1: -695, F2: (^(v12))}
                                                        _ = v57
                                                        _ = v58
                                                        _ = v59
                                                    }
                                                    v1 = (937 <= 600)
                                                    var v60 int32 = (v5 / ((2147483647) | 1))
                                                    _ = v60
                                                }
                                                _ = v53
                                            } else {
                                                fmt.Println(v13, v13)
                                                _ = len(v3)
                                                if len(v3) > 0 {
                                                    v3[0] = fi33
                                                }
                                            }
                                            var v61 int16 = 1000
                                            _ = v61
                                        }
                                    case 1:
                                        if (!false) {
                                            fmt.Println(v15)
                                            v13 = "hello"
                                        }
                                    }
                                }
                            }
                            switch "\n\t" {
                            case "foo bar":
                                var v62 uint = 255
                                _ = v62
                            case "a":
                                if (!v11) {
                                    var v63 uint32 = 760
                                    switch fi33 {
                                    case 128:
                                        var v64 uint64 = 18446744073709551615
                                        _ = v64
                                    case -2147483648:
                                        var v65 int = fi30
                                        _ = v65
                                    }
                                    fmt.Println(v13)
                                    _ = v63
                                }
                            default:
                                var v66 float64 = float64((0.1 + -2.25))
                                _ = v66
                            }
                            _ = fi33
                        }
                    } else {
                        fmt.Println(v5, v13)
                        switch (fi30 % ((fi30) | 1)) {
                        case 1:
                            if true {
                                switch 32768 {
                                case 1:
                                    if (0 <= 724) {
                                        fmt.Println(v12, v13)
                                        for fi67 := 0; fi67 < 1; fi67++ {
                                            if (1 == 100) {
                                                fmt.Println(fi67, v5)
                                            }
                                            _ = fi67
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                case 128:
                                    v68 := GenS4{F0: 255}
                                    _ = v68
                                }
                            }
                        default:
                            _ = len(v3)
                        }
                    }
                    v12 = v12
                    _ = fi30
                }
            }
            v15 -= v15
        } else {
            var v69 uint16 = 851
            _ = v69
        }
        for fi70 := 0; fi70 < 8; fi70++ {
            fmt.Println(v13, v13, v5)
            var v71 uint = 65535
            fmt.Println(v15, v71, v15)
            _ = fi70
            _ = v71
        }
    }
    if true {
        v72 := GenS6{F0: -269, F1: 920, F2: 421}
        _ = v72
    }
    if v1 {
        var v73 bool = v1
        if (!(!true)) {
            var v74 uint64 = 4294967295
            v75 := GenS1{F0: (v74 >> ((0) & 63)), F1: 1}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v74
            _ = v75
        }
        var v76 uint16 = 42
        _ = v73
        _ = v76
    }
    switch 9223372036854775807 {
    case 9223372036854775807:
        switch "\n\t" {
        case "":
            v13 = v13
        default:
            if (65535 >= 4294967295) {
                if v1 {
                    if v11 {
                        v77 := GenS1{F0: 4294967295, F1: 0}
                        _ = v77
                    } else {
                        v11 = v11
                    }
                    v11 = (436 != 2147483648)
                } else {
                    if v1 {
                        delete(v4, 639)
                        switch "hello" {
                        case "":
                            v78 := GenS1{F0: 42, F1: 0}
                            _ = v78
                        case "\n\t":
                            v12 = 42
                        case "foo bar":
                            if v11 {
                                if v1 {
                                    var v79 float32 = 1e10
                                    var v80 int = 2147483647
                                    _ = v79
                                    _ = v80
                                } else {
                                    if ((0 > 2) && (v1 && (32767 <= 633))) {
                                        if v1 {
                                            var v81 uint32 = 2
                                            if ((!(-1 < -643)) && (((((0.0 * (3.14159 * float32((((((((float64(1e10) + (((3.14159 + 3.14159) * 1e10) / 1.5)) - 3.14159) + -2.25) - 1e10) * 1e10) - 3.14159) - 1e10)))) + 1e10) * 0.0) - (3.14159 - -2.25)) >= 0.1)) {
                                                for fi82 := 0; fi82 < 6; fi82++ {
                                                    var v83 uint32 = v81
                                                    v5 -= v5
                                                    _ = fi82
                                                    _ = v83
                                                }
                                                if (v81 > v81) {
                                                    var v84 int = int(v15)
                                                    if (v5 < (v5 >> ((4294967295) & 63))) {
                                                        for fi85 := 0; fi85 < 9; fi85++ {
                                                            var v86 float32 = 0.0
                                                            v87 := GenS4{F0: 65535}
                                                            _ = fi85
                                                            _ = v86
                                                            _ = v87
                                                        }
                                                        for fi88 := 0; fi88 < 5; fi88++ {
                                                            var v89 uint16 = 32767
                                                            v90 := GenS7{F0: (^(v81))}
                                                            _ = fi88
                                                            _ = v89
                                                            _ = v90
                                                        }
                                                    } else {
                                                        if (32767 != 459) {
                                                            var v91 int16 = 801
                                                            var v92 uint32 = (v81 << ((4294967295) & 63))
                                                            _ = v91
                                                            _ = v92
                                                        }
                                                        v5 = v5
                                                    }
                                                    _ = v84
                                                } else {
                                                    var v93 string = (("" + v13) + (("hello" + "") + v13))
                                                    _ = v93
                                                }
                                            }
                                            switch v13 {
                                            case "":
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            case "foo bar":
                                                if (19 != (v15 / (((v15 << ((0) & 63))) | 1))) {
                                                    switch "hello" {
                                                    case "a":
                                                        var v94 int = -137
                                                        _ = v94
                                                    default:
                                                        v81 += (^(v81))
                                                    }
                                                }
                                            case "\n\t":
                                                var v95 int64 = 0
                                                _ = v95
                                            }
                                            _ = v81
                                        }
                                        if ((!true) && (!v1)) {
                                            for fi96 := 0; fi96 < 8; fi96++ {
                                                for fi97 := 0; fi97 < 2; fi97++ {
                                                    var v98 int16 = 1000
                                                    _ = fi97
                                                    _ = v98
                                                }
                                                v12 = v15
                                                _ = fi96
                                            }
                                        } else {
                                            var v99 uint32 = 2147483648
                                            v100 := GenS1{F0: 2, F1: 900}
                                            _ = v100
                                            _ = v99
                                        }
                                        for fi101 := 0; fi101 < 9; fi101++ {
                                            if v1 {
                                                switch fi101 {
                                                case 4294967295:
                                                    v102 := GenS5{F0: v11}
                                                    _ = v102
                                                case -2147483648:
                                                    v13 = v13
                                                default:
                                                    if v11 {
                                                        var v103 uint = 2
                                                        _ = v103
                                                    } else {
                                                        if ((v5 + (v5 * v5)) <= (v5 - v5)) {
                                                            for fi104 := 0; fi104 < 13; fi104++ {
                                                                var v105 int32 = (v5 << ((0) & 63))
                                                                var v106 int32 = 1
                                                                if v11 {
                                                                    if (628 >= 210) {
                                                                        v106 = (v105 ^ (v105 % ((v105) | 1)))
                                                                        var v107 int8 = 85
                                                                        _ = v107
                                                                    } else {
                                                                        if v1 {
                                                                            fmt.Println(v106)
                                                                            var v108 float32 = 0.0
                                                                            _ = v108
                                                                        }
                                                                        v109 := GenS8{F0: -94, F1: v15, F2: 9223372036854775807}
                                                                        _ = v109
                                                                    }
                                                                    switch (fi104 ^ 42) {
                                                                    case 32767:
                                                                        fmt.Println(v105, v5, v106)
                                                                        v110 := GenS5{F0: v1}
                                                                        _ = v110
                                                                    case -1:
                                                                        v111 := GenS3{F0: 18446744073709551615, F1: 0, F2: v15}
                                                                        _ = v111
                                                                    }
                                                                } else {
                                                                    v12 = (^(v15))
                                                                }
                                                                var v112 string = v13
                                                                var v113 uint = 1
                                                                _ = fi104
                                                                _ = v105
                                                                _ = v106
                                                                _ = v112
                                                                _ = v113
                                                            }
                                                            var v114 uint = 1
                                                            _ = v114
                                                        }
                                                    }
                                                }
                                                switch (fi101 << ((0) & 63)) {
                                                case 42:
                                                    switch (fi101 - fi101) {
                                                    case 770:
                                                        var v115 uint8 = 2
                                                        _ = v115
                                                    case 1:
                                                        v116 := GenS7{F0: 42}
                                                        _ = v116
                                                    case 32768:
                                                        var v117 int64 = 1
                                                        _ = v117
                                                    default:
                                                        v4[(-(fi101))] = v13
                                                    }
                                                case -9223372036854775808:
                                                    var v118 float32 = 1.5
                                                    _ = v118
                                                case 414:
                                                    v119 := GenS8{F0: -9, F1: v15, F2: 42}
                                                    _ = v119
                                                default:
                                                    switch v13 {
                                                    case "hello":
                                                        _ = len(v4)
                                                    }
                                                }
                                            }
                                            var v120 uint32 = 328
                                            if v1 {
                                                v13 = ("" + v13)
                                                for fi121 := 0; fi121 < 2; fi121++ {
                                                    var v122 int64 = -9223372036854775808
                                                    v123 := GenS7{F0: (v120 * 923)}
                                                    _ = fi121
                                                    _ = v122
                                                    _ = v123
                                                }
                                            } else {
                                                for fi124 := 0; fi124 < 4; fi124++ {
                                                    var v125 string = "foo bar"
                                                    v126 := GenS5{F0: v11}
                                                    _ = fi124
                                                    _ = v125
                                                    _ = v126
                                                }
                                                delete(v2, v120)
                                                v127 := GenS1{F0: 18446744073709551615, F1: -389}
                                                _ = v127
                                            }
                                            _ = fi101
                                            _ = v120
                                        }
                                    } else {
                                        for fi128 := 0; fi128 < 20; fi128++ {
                                            if (!(v13 >= v13)) {
                                                for fi129 := 0; fi129 < 19; fi129++ {
                                                    if v11 {
                                                        v2[542] = fi128
                                                    }
                                                    switch v13 {
                                                    case "hello":
                                                        v11 = v1
                                                    case "\n\t":
                                                        fmt.Println(v13)
                                                    }
                                                    _ = fi129
                                                }
                                                var v130 int8 = 0
                                                var v131 int = fi128
                                                _ = v130
                                                _ = v131
                                            } else {
                                                var v132 uint8 = 2
                                                _ = v132
                                            }
                                            _ = fi128
                                        }
                                        v15 = v12
                                    }
                                    switch v13 {
                                    case "foo bar":
                                        for fi133 := 0; fi133 < 7; fi133++ {
                                            _ = len(v2)
                                            var v134 int64 = -9223372036854775808
                                            _ = fi133
                                            _ = v134
                                        }
                                    }
                                    v1 = false
                                }
                                delete(v2, 629)
                            }
                        }
                    } else {
                        fmt.Println(v15, v15, v12)
                        for fi135 := 0; fi135 < 4; fi135++ {
                            for fi136 := 0; fi136 < 18; fi136++ {
                                if v1 {
                                    v1 = (!(v11 && false))
                                    v137 := GenS4{F0: 1}
                                    var v138 float64 = (((1.5 / 0.1) - 1.5) / (3.14159 * ((0.1 - 1.5) * float64(1e10))))
                                    _ = v137
                                    _ = v138
                                } else {
                                    delete(v2, 0)
                                }
                                _ = fi136
                            }
                            v13 = v13
                            for fi139 := 0; fi139 < 16; fi139++ {
                                if (!v1) {
                                    v1 = v11
                                } else {
                                    fmt.Println(v15)
                                }
                                switch v13 {
                                case "hello":
                                    v140 := GenS9{F0: -32}
                                    _ = v140
                                case "a":
                                    v141 := GenS4{F0: 458}
                                    _ = v141
                                default:
                                    v11 = (v11 || v1)
                                }
                                _ = fi139
                            }
                            _ = fi135
                        }
                    }
                }
                for fi142 := 0; fi142 < 3; fi142++ {
                    if false {
                        switch fi142 {
                        case 32767:
                            var v143 uint = 315
                            _ = v143
                        default:
                            v12 = 127
                        }
                        var v144 int16 = -1
                        var v145 float64 = 3.14159
                        _ = v144
                        _ = v145
                    } else {
                        v146 := GenS10{F0: 2, F1: (!((-9223372036854775808 != -9223372036854775808) && (!false)))}
                        for fi147 := 0; fi147 < 3; fi147++ {
                            v3 = append(v3, fi142)
                            v148 := GenS10{F0: 135, F1: false}
                            _ = fi147
                            _ = v148
                        }
                        _ = v146
                    }
                    var v149 int8 = 0
                    _ = fi142
                    _ = v149
                }
                v150 := GenS11{F0: 345}
                _ = v150
            } else {
                if v11 {
                    v11 = v11
                    v15 -= v15
                }
                var v151 string = v13
                _ = v151
            }
        }
    case 0:
        _ = len(v4)
    case -9223372036854775808:
        v11 = v1
    default:
        var v152 uint64 = 18446744073709551615
        _ = v152
    }
    v1 = true
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
