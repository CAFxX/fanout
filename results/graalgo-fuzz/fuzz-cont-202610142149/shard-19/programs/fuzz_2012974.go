package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int16 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int64; F1 int16 }
type GenS7 struct { F0 int8; F1 int8 }
type GenS8 struct { F0 bool; F1 bool; F2 string }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 65535
    var v3 string = "a"
    var v4 uint8 = 42
    v5 := v1[1]
    var v6 int32 = 538
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 uint16 = 1
        v9 := []int{(^(fi7)), 32768, (fi7 - 213)}
        if len(v9) > 0 {
            v9[0] = (-(fi7))
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    v5 = (339 >= 18446744073709551615)
    v10 := GenS1{F0: -128, F1: 112, F2: 32767}
    v2 -= v2
    var v11 uint8 = (v4 >> ((v2) & 63))
    v12 := GenS1{F0: 127, F1: 128, F2: 32767}
    if v5 {
        for fi13 := 0; fi13 < 5; fi13++ {
            for fi14 := 0; fi14 < 17; fi14++ {
                delete(v1, 875)
                _ = fi14
            }
            v6 += 1
            for fi15 := 0; fi15 < 1; fi15++ {
                var v16 float32 = -2.25
                v17 := GenS1{F0: int8(v4), F1: int(v4), F2: int16(v11)}
                if v5 {
                    if v5 {
                        var v18 float64 = 0.0
                        _ = v18
                    } else {
                        var v19 uint64 = 4294967295
                        if ("foo bar" < ((v3 + v3) + ("hello" + "foo bar"))) {
                            v11 += 42
                            v20 := []int{-633, fi15, (fi13 + (fi15 % ((int(v2)) | 1)))}
                            v6 += -814
                            _ = v20
                        } else {
                            if v5 {
                                fmt.Println(v19, v5, fi15)
                                v5 = (0 > -1)
                            } else {
                                switch 127 {
                                case 127:
                                    fmt.Println(v19, v2, v16)
                                case 42:
                                    var v21 int32 = (v6 << ((18446744073709551615) & 63))
                                    _ = v21
                                case -732:
                                    v22 := []int{(fi13 ^ (fi13 * fi13)), fi15, (fi13 / ((fi15) | 1)), (-(fi13))}
                                    if (42 <= 20) {
                                        var v23 int8 = 1
                                        if v5 {
                                            v24 := v1[42]
                                            var v25 float64 = float64(v16)
                                            if (!v24) {
                                                fmt.Println(v5, v24, v19)
                                            } else {
                                                fmt.Println(v16)
                                            }
                                            _ = v24
                                            _ = v25
                                        } else {
                                            if v5 {
                                                if false {
                                                    _ = len(v1)
                                                    v4 = 128
                                                } else {
                                                    v6 = v6
                                                    _ = len(v1)
                                                    var v26 string = v3
                                                    _ = v26
                                                }
                                                if false {
                                                    switch "foo bar" {
                                                    case "hello":
                                                        var v27 float32 = float32(1.5)
                                                        _ = v27
                                                    default:
                                                        var v28 int16 = 1
                                                        _ = v28
                                                    }
                                                    if ((v16 == float32(1.5)) && (v4 > (v11 & 113))) {
                                                        switch (fi13 % ((fi13) | 1)) {
                                                        case 42:
                                                            v29 := GenS1{F0: int8(fi13), F1: (fi13 >> ((v2) & 63)), F2: 1}
                                                            _ = v29
                                                        default:
                                                            var v30 float64 = (3.14159 - ((-2.25 + 0.1) * (float64(v16) * (float64(0.1) + ((3.14159 / float64(v16)) - (((1e10 * 0.1) * (0.0 / 1.0)) * ((1.5 / (3.14159 * ((3.14159 + 0.1) - (((((0.0 + 1e10) - (-2.25 + (0.1 + (1e10 * 1e10)))) / 1e10) * 0.1) + ((float64(-2.25) / 1.5) + 0.1))))) - (-2.25 / (3.14159 + 0.0)))))))))
                                                            _ = v30
                                                        }
                                                    }
                                                }
                                                var v31 uint32 = 4294967295
                                                _ = v31
                                            }
                                        }
                                        var v32 float32 = v16
                                        _ = v23
                                        _ = v32
                                    } else {
                                        v33 := GenS1{F0: 1, F1: fi13, F2: 1}
                                        switch (fi13 & fi15) {
                                        case -496:
                                            v34 := GenS2{F0: 0, F1: 65535}
                                            _ = v34
                                        case -2147483648:
                                            if (-767 > 9223372036854775807) {
                                                v35 := GenS3{F0: v11, F1: -9223372036854775808}
                                                v36 := GenS1{F0: 42, F1: 796, F2: 0}
                                                var v37 uint64 = 140
                                                _ = v35
                                                _ = v36
                                                _ = v37
                                            } else {
                                                if false {
                                                    switch (-(fi15)) {
                                                    case 9223372036854775807:
                                                        if true {
                                                            v19 = (v19 % ((v19) | 1))
                                                            var v38 int64 = 0
                                                            _ = v38
                                                        } else {
                                                            v39 := v1[65535]
                                                            var v40 uint64 = v19
                                                            v41 := GenS3{F0: (v11 ^ uint8(v19)), F1: 0}
                                                            _ = v39
                                                            _ = v40
                                                            _ = v41
                                                        }
                                                    case 0:
                                                        v42 := GenS4{F0: (fi15 * 32768), F1: 2}
                                                        _ = v42
                                                    default:
                                                        v6 -= v6
                                                        delete(v1, 602)
                                                        v19 -= 2
                                                    }
                                                }
                                            }
                                        }
                                        v19 = (v19 - v19)
                                        _ = v33
                                    }
                                    _ = v22
                                }
                            }
                            v43 := GenS1{F0: 91, F1: (fi15 >> ((v2) & 63)), F2: -77}
                            v6 = v6
                            _ = v43
                        }
                        _ = v19
                    }
                    fmt.Println(v11)
                }
                switch "a" {
                case "":
                    switch v3 {
                    case "a":
                        if true {
                            fmt.Println(v2)
                            v4 -= v11
                            var v44 int8 = 117
                            v45 := []int{fi15, fi15, -2147483648, fi13}
                            _ = v44
                            _ = v45
                        }
                    default:
                        if v5 {
                            var v46 int16 = 42
                            if (!false) {
                                _ = len(v1)
                                if (7 != 127) {
                                    if v5 {
                                        if ((v46 - v46) <= (v46 * (v46 - 42))) {
                                            var v47 int = (fi15 - (^(fi13)))
                                            v16 += 0.1
                                            _ = v47
                                        } else {
                                            var v48 uint = (^(v2))
                                            v49 := []int{-118, 42, 128}
                                            _ = v48
                                            _ = v49
                                        }
                                        fmt.Println(v11, v16, v11)
                                    }
                                    var v50 uint32 = 2147483648
                                    var v51 uint64 = 0
                                    _ = v50
                                    _ = v51
                                }
                            }
                            v6 += v6
                            _ = v46
                        } else {
                            var v52 uint8 = v4
                            _ = v52
                        }
                    }
                default:
                    fmt.Println(v2, v3, v4)
                }
                v53 := GenS4{F0: fi13, F1: v4}
                _ = fi15
                _ = v16
                _ = v17
                _ = v53
            }
            _ = fi13
        }
        v5 = (-888 <= 42)
        v2 += (v2 + (v2 / (((v2 * (v2 & 785))) | 1)))
    } else {
        v3 = (v3 + v3)
        for fi54 := 0; fi54 < 4; fi54++ {
            var v55 int16 = int16(v4)
            _ = fi54
            _ = v55
        }
        if (42 != 0) {
            var v56 int = 4294967295
            if (981 != 1) {
                v5 = (-2.25 > -2.25)
            } else {
                switch v56 {
                case 1:
                    switch "" {
                    case "":
                        switch (v56 % (((^(v56))) | 1)) {
                        case 626:
                            var v57 bool = (false || v5)
                            _ = v57
                        case -9223372036854775808:
                            v56 = v56
                        default:
                            for fi58 := 0; fi58 < 15; fi58++ {
                                for fi59 := 0; fi59 < 15; fi59++ {
                                    var v60 float64 = 0.0
                                    v2 = (v2 << ((470) & 63))
                                    v4 = (v11 | v4)
                                    _ = fi59
                                    _ = v60
                                }
                                _ = fi58
                            }
                        }
                    case "\n\t":
                        var v61 int16 = 32767
                        _ = v61
                    case "foo bar":
                        v2 = 42
                    default:
                        if false {
                            v5 = ((v3 + ((v3 + "\n\t") + v3)) == ("\n\t" + v3))
                        }
                    }
                case -789:
                    var v62 bool = (v5 && (v2 < (v2 / (((v2 - 42)) | 1))))
                    _ = v62
                case -2147483648:
                    switch v56 {
                    case -559:
                        var v63 int8 = 42
                        _ = v63
                    default:
                        v5 = (!((42 > 387) || v5))
                    }
                }
                if true {
                    var v64 string = v3
                    _ = v64
                }
            }
            _ = v56
        } else {
            for fi65 := 0; fi65 < 12; fi65++ {
                v2 = (v2 + v2)
                _ = fi65
            }
            var v66 bool = (0.1 == float32(3.14159))
            fmt.Println(v4, v4, v2)
            _ = v66
        }
    }
    fmt.Println(v2)
    v67 := v1[2147483648]
    v68 := v1[56]
    v1[775] = (!(317 >= (v2 + 2)))
    if (!false) {
        if ((-1 > 141) && (v6 < v6)) {
            v69 := v1[347]
            delete(v1, 411)
            v70 := v1[2]
            _ = v69
            _ = v70
        }
        if true {
            switch 32768 {
            case -526:
                v71 := []int{-504, 4294967295}
                _ = v71
            case 523:
                v72 := GenS1{F0: 1, F1: 32767, F2: -32768}
                _ = v72
            case -2147483648:
                _ = len(v1)
            default:
                var v73 int64 = -1
                _ = v73
            }
            v74 := GenS4{F0: 4294967295, F1: v4}
            _ = v74
        } else {
            v6 -= -860
            v67 = (!(1 != 652))
        }
        fmt.Println(v68, v2, v11)
    } else {
        v67 = (true || (!v68))
        if v5 {
            var v75 int = 124
            var v76 uint = 255
            for fi77 := 0; fi77 < 13; fi77++ {
                for fi78 := 0; fi78 < 13; fi78++ {
                    switch (fi77 + (fi78 << ((4294967295) & 63))) {
                    case -702:
                        switch "a" {
                        case "hello":
                            v3 = v3
                        case "foo bar":
                            v75 = fi78
                        case "a":
                            v79 := GenS2{F0: 0, F1: 536}
                            v80 := GenS1{F0: 1, F1: fi78, F2: -32768}
                            fmt.Println(v67, fi78, v76)
                            _ = v79
                            _ = v80
                        default:
                            switch "\n\t" {
                            case "foo bar":
                                var v81 bool = (v67 && v67)
                                _ = v81
                            default:
                                _ = len(v1)
                            }
                        }
                    }
                    _ = fi78
                }
                v82 := []int{}
                var v83 bool = false
                _ = fi77
                _ = v82
                _ = v83
            }
            _ = v75
            _ = v76
        } else {
            switch ("foo bar" + v3) {
            case "hello":
                if v67 {
                    if (!true) {
                        if (v2 >= 42) {
                            if (100 > -1) {
                                v2 = (v2 + v2)
                                v6 -= -1
                                v5 = v5
                            } else {
                                v2 -= v2
                                fmt.Println(v68, v6)
                                v1[2147483648] = (!(true || (!true)))
                            }
                            v5 = (v3 == v3)
                            v84 := GenS2{F0: 1, F1: 367}
                            _ = v84
                        }
                        switch -9223372036854775808 {
                        case 4294967295:
                            if v67 {
                                for fi85 := 0; fi85 < 17; fi85++ {
                                    var v86 bool = v68
                                    _ = fi85
                                    _ = v86
                                }
                                var v87 uint32 = 0
                                var v88 int64 = 909
                                _ = v87
                                _ = v88
                            } else {
                                var v89 float64 = 0.1
                                if v5 {
                                    if v5 {
                                        v67 = (0 > 1000)
                                    }
                                    if (!(!(!v68))) {
                                        v90 := GenS4{F0: 9223372036854775807, F1: (^(v11))}
                                        switch "hello" {
                                        case "":
                                            switch "a" {
                                            case "\n\t":
                                                if (1 != -1) {
                                                    for fi91 := 0; fi91 < 1; fi91++ {
                                                        var v92 string = v3
                                                        if v5 {
                                                            if (!v5) {
                                                                var v93 uint64 = 18446744073709551615
                                                                v92 = (v3 + v3)
                                                                _ = v93
                                                            }
                                                            v94 := GenS3{F0: v11, F1: 42}
                                                            if v5 {
                                                                if true {
                                                                    if v5 {
                                                                        for fi95 := 0; fi95 < 10; fi95++ {
                                                                            v68 = (!(-9223372036854775808 != -1))
                                                                            switch v92 {
                                                                            case "foo bar":
                                                                                var v96 uint8 = v11
                                                                                _ = v96
                                                                            case "a":
                                                                                if (-2.25 < 3.14159) {
                                                                                    fmt.Println(v2, v5)
                                                                                    var v97 uint64 = 4294967295
                                                                                    switch (fi91 / ((fi95) | 1)) {
                                                                                    case 127:
                                                                                        if (v67 || v5) {
                                                                                            v98 := []int{fi95, int(v97), fi95}
                                                                                            var v99 int16 = 583
                                                                                            switch v92 {
                                                                                            case "foo bar":
                                                                                                v4 -= (^(v4))
                                                                                                v3 = (v92 + "hello")
                                                                                            case "\n\t":
                                                                                                if len(v98) >= 2 {
                                                                                                    v98 = v98[1:len(v98)]
                                                                                                }
                                                                                            case "hello":
                                                                                                var v100 int16 = v99
                                                                                                _ = v100
                                                                                            }
                                                                                            var v101 bool = v67
                                                                                            v102 := GenS1{F0: -76, F1: (fi95 - 1), F2: 1}
                                                                                            _ = v101
                                                                                            _ = v102
                                                                                            _ = v98
                                                                                            _ = v99
                                                                                        }
                                                                                    case -1:
                                                                                        var v103 float64 = (v89 + (0.1 * v89))
                                                                                        _ = v103
                                                                                    }
                                                                                    _ = v97
                                                                                } else {
                                                                                    v92 = "foo bar"
                                                                                    v92 = (((((("hello" + "\n\t") + ((("" + "a") + v3) + "")) + v92) + "foo bar") + (("foo bar" + ("hello" + v3)) + (v3 + ("hello" + ((v3 + "\n\t") + ""))))) + "a")
                                                                                    v68 = v67
                                                                                }
                                                                            case "\n\t":
                                                                                v89 = (-2.25 / v89)
                                                                            default:
                                                                                if v68 {
                                                                                    if (810 <= 32767) {
                                                                                        fmt.Println(v3)
                                                                                        var v104 uint64 = 1
                                                                                        _ = v104
                                                                                    }
                                                                                    v105 := GenS2{F0: 42, F1: 1}
                                                                                    _ = v105
                                                                                }
                                                                            }
                                                                            _ = fi95
                                                                        }
                                                                        v67 = v68
                                                                        if (127 <= -85) {
                                                                            v92 = ""
                                                                        }
                                                                    } else {
                                                                        if (v11 <= (^(v11))) {
                                                                            v89 += (v89 * ((3.14159 * ((float64(-2.25) * v89) - float64(0.0))) - -2.25))
                                                                        }
                                                                        var v106 uint64 = 1
                                                                        _ = v106
                                                                    }
                                                                }
                                                            }
                                                            _ = v94
                                                        } else {
                                                            fmt.Println(v5, v4)
                                                            var v107 uint32 = 811
                                                            for fi108 := 0; fi108 < 13; fi108++ {
                                                                if (((612 == 902) || v67) || v5) {
                                                                    v6 = v6
                                                                    fmt.Println(fi108, v67)
                                                                } else {
                                                                    fmt.Println(v2, v4, v67)
                                                                }
                                                                fmt.Println(v6)
                                                                v5 = v67
                                                                _ = fi108
                                                            }
                                                            _ = v107
                                                        }
                                                        _ = fi91
                                                        _ = v92
                                                    }
                                                    var v109 string = "foo bar"
                                                    _ = v109
                                                } else {
                                                    v110 := []int{-9223372036854775808, -2147483648, 42}
                                                    v111 := GenS1{F0: 42, F1: -265, F2: -32768}
                                                    for fi112 := 0; fi112 < 4; fi112++ {
                                                        v67 = true
                                                        _ = fi112
                                                    }
                                                    _ = v110
                                                    _ = v111
                                                }
                                            case "foo bar":
                                                if true {
                                                    for fi113 := 0; fi113 < 4; fi113++ {
                                                        var v114 int64 = -9223372036854775808
                                                        v115 := GenS5{F0: (v2 >> ((v2) & 63))}
                                                        _ = fi113
                                                        _ = v114
                                                        _ = v115
                                                    }
                                                    delete(v1, 2147483647)
                                                    fmt.Println(v2)
                                                }
                                            case "":
                                                v11 += v11
                                            default:
                                                for fi116 := 0; fi116 < 9; fi116++ {
                                                    var v117 string = v3
                                                    v118 := v1[359]
                                                    var v119 uint32 = 2
                                                    _ = fi116
                                                    _ = v117
                                                    _ = v118
                                                    _ = v119
                                                }
                                            }
                                        case "hello":
                                            fmt.Println(v5, v5)
                                        case "foo bar":
                                            var v120 uint32 = uint32(v2)
                                            _ = v120
                                        default:
                                            for fi121 := 0; fi121 < 5; fi121++ {
                                                fmt.Println(v3)
                                                v122 := GenS6{F0: 237, F1: 398}
                                                v3 = (("a" + v3) + (((v3 + "") + v3) + "hello"))
                                                _ = fi121
                                                _ = v122
                                            }
                                        }
                                        _ = v90
                                    }
                                } else {
                                    var v123 int32 = (v6 / ((v6) | 1))
                                    var v124 float32 = 0.0
                                    v125 := GenS1{F0: -128, F1: 431, F2: -32768}
                                    _ = v123
                                    _ = v124
                                    _ = v125
                                }
                                var v126 int = 1
                                _ = v126
                                _ = v89
                            }
                        case -86:
                            for fi127 := 0; fi127 < 15; fi127++ {
                                delete(v1, 2)
                                var v128 bool = v67
                                v1[4294967295] = (((((("foo bar" + v3) + v3) + (v3 + (v3 + v3))) + v3) + v3) == (v3 + v3))
                                _ = fi127
                                _ = v128
                            }
                        default:
                            fmt.Println(v3, v68, v3)
                        }
                    } else {
                        v11 = 64
                    }
                    var v129 float32 = 3.14159
                    _ = v129
                } else {
                    if false {
                        v68 = (!(!((1 < uint64(v11)) || (v3 <= "foo bar"))))
                        v67 = v5
                    }
                }
            case "a":
                for fi130 := 0; fi130 < 4; fi130++ {
                    switch fi130 {
                    case 1:
                        v131 := GenS4{F0: fi130, F1: v11}
                        _ = v131
                    case 128:
                        if v67 {
                            v3 = (v3 + v3)
                            for fi132 := 0; fi132 < 15; fi132++ {
                                v133 := GenS3{F0: v4, F1: 33}
                                switch (-(fi130)) {
                                case 32767:
                                    v5 = (fi132 == 2147483647)
                                case -843:
                                    delete(v1, 42)
                                case -1:
                                    if false {
                                        if v67 {
                                            var v134 int8 = -16
                                            _ = v134
                                        }
                                        v135 := GenS1{F0: -1, F1: fi130, F2: 1}
                                        _ = v135
                                    } else {
                                        v136 := []int{}
                                        _ = v136
                                    }
                                    v2 -= (^(v2))
                                default:
                                    if v68 {
                                        if ((1 >= 4294967295) && v5) {
                                            v137 := GenS2{F0: 1, F1: 133}
                                            _ = v137
                                        }
                                        if v5 {
                                            v138 := GenS7{F0: -89, F1: -35}
                                            var v139 int32 = v6
                                            _ = v138
                                            _ = v139
                                        } else {
                                            var v140 uint8 = v4
                                            var v141 int32 = v6
                                            _ = v140
                                            _ = v141
                                        }
                                    }
                                }
                                delete(v1, 1)
                                v142 := []int{fi132, fi132, fi130}
                                _ = fi132
                                _ = v133
                                _ = v142
                            }
                            var v143 float64 = 1.5
                            _ = v143
                        } else {
                            var v144 int64 = 9223372036854775807
                            for fi145 := 0; fi145 < 11; fi145++ {
                                var v146 int16 = -510
                                v147 := GenS3{F0: 128, F1: (v144 & (^(v144)))}
                                _ = fi145
                                _ = v146
                                _ = v147
                            }
                            var v148 uint16 = 6
                            _ = v144
                            _ = v148
                        }
                    case 32767:
                        var v149 uint8 = (v11 * 98)
                        _ = v149
                    }
                    for fi150 := 0; fi150 < 2; fi150++ {
                        v151 := v1[2147483647]
                        if v5 {
                            var v152 int = fi130
                            v67 = v5
                            _ = v152
                        }
                        _ = fi150
                        _ = v151
                    }
                    if true {
                        v5 = v5
                        fmt.Println(v2, v3, v68)
                    }
                    _ = fi130
                }
            default:
                for fi153 := 0; fi153 < 5; fi153++ {
                    v4 = uint8(v2)
                    _ = fi153
                }
            }
            var v154 uint32 = 237
            v155 := []int{}
            _ = v154
            _ = v155
        }
        fmt.Println(v2, v6)
    }
    for fi156 := 0; fi156 < 6; fi156++ {
        v157 := v1[65535]
        v11 = v11
        _ = fi156
        _ = v157
    }
    for fi158 := 0; fi158 < 11; fi158++ {
        if (42 <= 350) {
            v68 = (!(v4 != v4))
            fmt.Println(v11, fi158)
            var v159 int = fi158
            _ = v159
        }
        v2 = v2
        if v5 {
            var v160 bool = ((!(!((v67 || (42 >= 239)) && (true || (v67 || v68))))) || (0.0 < 3.14159))
            var v161 uint32 = 910
            if ((fi158 << ((v2) & 63)) != fi158) {
                var v162 float32 = 1.5
                var v163 int32 = v6
                _ = v162
                _ = v163
            } else {
                var v164 float32 = 3.14159
                fmt.Println(v67, v67)
                switch (("" + v3) + "foo bar") {
                case "\n\t":
                    v165 := GenS8{F0: v160, F1: v67, F2: ("hello" + v3)}
                    _ = v165
                case "foo bar":
                    var v166 float32 = v164
                    _ = v166
                default:
                    delete(v1, (^(v161)))
                }
                _ = v164
            }
            _ = v160
            _ = v161
        } else {
            fmt.Println(fi158, v4, v4)
            var v167 uint64 = 42
            _ = v167
        }
        _ = fi158
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v67
    _ = v68
}
