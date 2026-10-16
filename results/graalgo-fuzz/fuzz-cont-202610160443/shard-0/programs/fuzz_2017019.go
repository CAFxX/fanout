package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 int16; F1 int64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint32; F1 uint64; F2 uint }
type GenS9 struct { F0 string; F1 uint32 }
type GenS10 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 bool; F1 int8 }
type GenS13 struct { F0 uint16; F1 int8 }
type GenS14 struct { F0 bool; F1 int64 }
type GenS15 struct { F0 uint; F1 int32 }
type GenS16 struct { F0 int8; F1 uint16; F2 uint8 }
type GenS17 struct { F0 uint8; F1 bool; F2 int8 }
type GenS18 struct { F0 int16; F1 string }
func genFn0(p0 int, p1 bool) int8 {
    var v4 uint = 129
    switch (^(p0)) {
    case -97:
        var v5 int32 = 764
        _ = v5
    case 2147483647:
        v6 := GenS1{F0: 2147483648, F1: -1, F2: int(v4)}
        _ = v6
    default:
        switch "hello" {
        case "foo bar":
            p1 = (!(!true))
        case "":
            if false {
                p0 = (p0 - (^(p0)))
            } else {
                var v7 int64 = 17
                var v8 float32 = float32((((0.0 + float64(float32(-2.25))) - 0.0) / 0.1))
                for fi9 := 0; fi9 < 20; fi9++ {
                    for fi10 := 0; fi10 < 3; fi10++ {
                        v11 := GenS2{F0: (p1 || (false || ((0.0 != ((1.5 / 1e10) + -2.25)) || (!(!p1)))))}
                        _ = fi10
                        _ = v11
                    }
                    _ = fi9
                }
                _ = v7
                _ = v8
            }
        case "a":
            var v12 uint8 = 128
            _ = v12
        default:
            if p1 {
                for fi13 := 0; fi13 < 4; fi13++ {
                    var v14 uint16 = 42
                    _ = fi13
                    _ = v14
                }
                v15 := GenS3{F0: 128, F1: 9223372036854775807}
                _ = v15
            }
        }
    }
    var v16 int16 = 32767
    var v17 uint32 = 993
    for fi18 := 0; fi18 < 5; fi18++ {
        fmt.Println(p0, v17, fi18)
        _ = fi18
    }
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    _ = v4
    return 127
}

func main() {
    v1 := make(map[int]bool)
    var v2 bool = (!(-672 != 360))
    var v3 int64 = 42
    var v19 int16 = 32767
    fmt.Println(v19, v3, v2)
    switch "\n\t" {
    case "foo bar":
        for fi20 := 0; fi20 < 10; fi20++ {
            for fi21 := 0; fi21 < 1; fi21++ {
                v22 := GenS1{F0: 916, F1: -128, F2: -292}
                _ = fi21
                _ = v22
            }
            v3 -= (v3 * v3)
            _ = fi20
        }
    case "\n\t":
        fmt.Println(v19)
    case "a":
        var v23 bool = v2
        _ = v23
    default:
        v24 := GenS2{F0: v2}
        _ = v24
    }
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        v2 = v2
    case "":
        fmt.Println(v2, v19)
    default:
        v25 := GenS4{F0: 2, F1: 2147483648}
        _ = v25
    }
    v26 := GenS5{F0: -32768, F1: v3}
    for fi27 := 0; fi27 < 3; fi27++ {
        switch "a" {
        case "\n\t":
            v28 := GenS4{F0: 127, F1: 2147483647}
            _ = v28
        case "hello":
            v29 := GenS6{F0: v2}
            _ = v29
        case "a":
            delete(v1, -2147483648)
        }
        _ = fi27
    }
    switch 127 {
    case -2147483648:
        for fi30 := 0; fi30 < 4; fi30++ {
            if ((v19 / ((1000) | 1)) == (v19 | (v19 / ((v19) | 1)))) {
                fmt.Println(fi30, v19, v2)
            } else {
                v31 := GenS4{F0: 26, F1: 65535}
                if true {
                    v3 -= v3
                    if ((fi30 & fi30) == fi30) {
                        for fi32 := 0; fi32 < 3; fi32++ {
                            fmt.Println(fi32, v3, fi30)
                            _ = fi32
                        }
                        if true {
                            for fi33 := 0; fi33 < 8; fi33++ {
                                v34 := GenS6{F0: v2}
                                var v35 int64 = v3
                                _ = len(v1)
                                _ = fi33
                                _ = v34
                                _ = v35
                            }
                            for fi36 := 0; fi36 < 14; fi36++ {
                                fmt.Println(v19)
                                switch 32767 {
                                case 4294967295:
                                    v3 -= v3
                                case -1:
                                    var v37 int64 = v3
                                    _ = v37
                                case 1:
                                    v38 := GenS1{F0: 50, F1: -9, F2: 180}
                                    _ = v38
                                default:
                                    var v39 bool = v2
                                    _ = v39
                                }
                                if (!(!(32767 >= 2))) {
                                    if false {
                                        var v40 int32 = -620
                                        if v2 {
                                            v3 = v3
                                            var v41 uint64 = 1
                                            v42 := GenS4{F0: 42, F1: 2147483648}
                                            _ = v41
                                            _ = v42
                                        } else {
                                            v2 = v2
                                        }
                                        var v43 float64 = -2.25
                                        _ = v40
                                        _ = v43
                                    } else {
                                        v1[327] = (0 >= -2147483648)
                                        fmt.Println(v2)
                                        var v44 string = "a"
                                        _ = v44
                                    }
                                    v19 += 32767
                                }
                                _ = fi36
                            }
                        }
                        v3 = (v3 * 0)
                    } else {
                        v2 = v2
                        v45 := GenS4{F0: 127, F1: 2}
                        _ = v45
                    }
                    v46 := GenS7{F0: 0}
                    _ = v46
                }
                var v47 int8 = genFn0(fi30, (!v2))
                _ = v31
                _ = v47
            }
            v3 = (v3 + 1)
            _ = fi30
        }
    case 61:
        if true {
            var v48 int32 = 1
            var v49 float64 = 0.0
            v50 := GenS1{F0: 65535, F1: 1, F2: 9223372036854775807}
            _ = v48
            _ = v49
            _ = v50
        }
    }
    v3 -= v3
    v51 := GenS8{F0: 42, F1: 1, F2: 4294967295}
    for fi52 := 0; fi52 < 15; fi52++ {
        v1[(fi52 >> ((0) & 63))] = (v19 >= v19)
        if false {
            v53 := []int{fi52, 206}
            v2 = (!false)
            _ = v53
        } else {
            var v54 string = "\n\t"
            _ = v54
        }
        _ = fi52
    }
    if ("foo bar" <= "a") {
        v55 := GenS6{F0: v2}
        _ = v55
    } else {
        v56 := GenS9{F0: "a", F1: 65535}
        switch -9223372036854775808 {
        case 4294967295:
            var v57 int32 = -2147483648
            _ = v57
        case 0:
            v19 += 32767
        case -809:
            var v58 uint = 2
            _ = v58
        default:
            var v59 uint64 = 4294967295
            _ = v59
        }
        v19 -= v19
        _ = v56
    }
    for fi60 := 0; fi60 < 11; fi60++ {
        var v61 int64 = v3
        v62 := GenS10{F0: 372, F1: v61, F2: 127}
        _ = fi60
        _ = v61
        _ = v62
    }
    for fi63 := 0; fi63 < 15; fi63++ {
        _ = len(v1)
        switch 32768 {
        case -1:
            if true {
                var v64 int64 = 42
                _ = v64
            }
        case 1:
            for fi65 := 0; fi65 < 10; fi65++ {
                v66 := v1[fi65]
                var v67 uint32 = 42
                v68 := GenS5{F0: v19, F1: v3}
                _ = fi65
                _ = v66
                _ = v67
                _ = v68
            }
        case 32767:
            v69 := GenS3{F0: 87, F1: -9223372036854775808}
            _ = v69
        }
        _ = fi63
    }
    switch (("foo bar" + ("foo bar" + "hello")) + "a") {
    case "\n\t":
        if (128 >= 32768) {
            if (166 > 2) {
                var v70 int16 = v19
                switch ("a" + "foo bar") {
                case "foo bar":
                    v71 := GenS6{F0: v2}
                    _ = v71
                case "\n\t":
                    for fi72 := 0; fi72 < 2; fi72++ {
                        v73 := []int{}
                        if (((!(true || v2)) || v2) || v2) {
                            v74 := GenS8{F0: 65535, F1: 1, F2: 2}
                            fmt.Println(v70, fi72)
                            _ = v74
                        }
                        if v2 {
                            var v75 int = fi72
                            _ = v75
                        }
                        _ = fi72
                        _ = v73
                    }
                case "a":
                    switch ("" + "\n\t") {
                    case "":
                        switch "a" {
                        case "hello":
                            for fi76 := 0; fi76 < 18; fi76++ {
                                delete(v1, fi76)
                                for fi77 := 0; fi77 < 15; fi77++ {
                                    var v78 int16 = 1000
                                    var v79 uint64 = 42
                                    v2 = v2
                                    _ = fi77
                                    _ = v78
                                    _ = v79
                                }
                                for fi80 := 0; fi80 < 13; fi80++ {
                                    fmt.Println(fi80, fi80)
                                    var v81 float64 = 0.0
                                    switch "" {
                                    case "\n\t":
                                        var v82 bool = v2
                                        _ = v82
                                    case "foo bar":
                                        v81 += 1e10
                                    default:
                                        delete(v1, 32767)
                                    }
                                    _ = fi80
                                    _ = v81
                                }
                                _ = fi76
                            }
                        case "\n\t":
                            var v83 int8 = -128
                            _ = v83
                        }
                    }
                }
                _ = v70
            } else {
                fmt.Println(v2)
                for fi84 := 0; fi84 < 19; fi84++ {
                    v85 := GenS11{F0: (fi84 | fi84)}
                    for fi86 := 0; fi86 < 11; fi86++ {
                        var v87 int64 = v3
                        _ = fi86
                        _ = v87
                    }
                    _ = fi84
                    _ = v85
                }
            }
            if v2 {
                for fi88 := 0; fi88 < 18; fi88++ {
                    if v2 {
                        var v89 int = fi88
                        for fi90 := 0; fi90 < 20; fi90++ {
                            var v91 int64 = (v3 | (v3 << ((2) & 63)))
                            fmt.Println(v89, v2, v91)
                            var v92 int16 = (v19 * -32768)
                            _ = fi90
                            _ = v91
                            _ = v92
                        }
                        _ = v89
                    }
                    if v2 {
                        v2 = (155 == 42)
                        var v93 float64 = ((0.1 / 0.1) / 1.0)
                        _ = v93
                    } else {
                        if true {
                            var v94 uint64 = 2
                            switch fi88 {
                            case 32767:
                                v95 := GenS11{F0: fi88}
                                _ = v95
                            case 2147483647:
                                v19 = 32767
                            default:
                                switch 32768 {
                                case -1:
                                    v96 := []int{fi88}
                                    _ = v96
                                }
                            }
                            _ = v94
                        } else {
                            v3 -= 9223372036854775807
                        }
                        v2 = (-2.25 == 0.0)
                    }
                    _ = fi88
                }
                v97 := GenS4{F0: 118, F1: 1}
                _ = v97
            } else {
                v2 = (((3.14159 - 0.1) - (0.0 / 1.0)) > 1.5)
            }
            fmt.Println(v3, v2, v19)
        } else {
            if v2 {
                v3 = 0
                for fi98 := 0; fi98 < 7; fi98++ {
                    v99 := GenS5{F0: (v19 ^ (v19 / ((1000) | 1))), F1: (v3 ^ v3)}
                    var v100 int32 = 629
                    _ = fi98
                    _ = v100
                    _ = v99
                }
            } else {
                var v101 uint8 = 0
                v102 := GenS12{F0: v2, F1: -1}
                if (754 == 2) {
                    var v103 int32 = 223
                    fmt.Println(v3, v3, v3)
                    _ = v103
                }
                _ = v101
                _ = v102
            }
            v104 := []int{-246, 1, 128, 851}
            _ = v104
        }
    case "a":
        fmt.Println(v3, v2)
    case "hello":
        switch "" {
        case "\n\t":
            v105 := []int{190}
            _ = v105
        }
    default:
        if v2 {
            v19 = v19
            for fi106 := 0; fi106 < 19; fi106++ {
                v107 := GenS13{F0: 255, F1: -121}
                _ = fi106
                _ = v107
            }
            if false {
                for fi108 := 0; fi108 < 1; fi108++ {
                    v109 := []int{(fi108 ^ (fi108 % ((-2147483648) | 1))), fi108, fi108}
                    var v110 int64 = (-(v3))
                    _ = fi108
                    _ = v109
                    _ = v110
                }
                var v111 bool = (0 < 996)
                _ = v111
            } else {
                if (-1 == 236) {
                    if v2 {
                        for fi112 := 0; fi112 < 12; fi112++ {
                            v113 := GenS8{F0: 2147483647, F1: 4294967295, F2: 1}
                            _ = fi112
                            _ = v113
                        }
                        var v114 uint32 = 136
                        _ = v114
                    }
                    var v115 float64 = (float64(0.1) * 1.5)
                    _ = v115
                } else {
                    var v116 uint16 = 65535
                    _ = v116
                }
                v1[-9223372036854775808] = (!(v2 || v2))
            }
        } else {
            v117 := []int{-9223372036854775808, 32767, 128}
            _ = v117
        }
    }
    switch 1 {
    case 343:
        v1[128] = (v2 || v2)
    case 128:
        v118 := GenS4{F0: 26, F1: 0}
        _ = v118
    }
    for fi119 := 0; fi119 < 14; fi119++ {
        for fi120 := 0; fi120 < 12; fi120++ {
            switch -1 {
            case 32768:
                switch fi120 {
                case 128:
                    switch fi120 {
                    case 32767:
                        v121 := []int{}
                        _ = v121
                    default:
                        if true {
                            v122 := []int{fi119, (fi119 + (fi119 * 994))}
                            _ = v122
                        } else {
                            fmt.Println(v3)
                            v2 = v2
                        }
                    }
                case 9223372036854775807:
                    v123 := []int{}
                    _ = v123
                }
                v124 := []int{fi120, 127, 32767, (^(fi120))}
                if ((true && v2) && v2) {
                    v125 := GenS2{F0: (-9223372036854775808 >= (^(v3)))}
                    fmt.Println(v2, v3, v2)
                    switch "" {
                    case "":
                        v126 := GenS14{F0: (v2 && (!false)), F1: v3}
                        _ = v126
                    case "hello":
                        if v2 {
                            fmt.Println(v2)
                            fmt.Println(v2)
                            fmt.Println(fi119, v3)
                            switch "a" {
                            case "":
                                if v2 {
                                    v127 := GenS15{F0: 1, F1: 1}
                                    var v128 bool = ((fi120 >= (fi120 % ((32767) | 1))) && v2)
                                    v129 := v1[(fi119 | fi119)]
                                    _ = v127
                                    _ = v128
                                    _ = v129
                                } else {
                                    v3 = v3
                                }
                            default:
                                if (!false) {
                                    v130 := GenS2{F0: (399 == 430)}
                                    _ = v130
                                }
                            }
                            if ((v3 ^ v3) < (v3 >> ((362) & 63))) {
                                var v131 uint = 65535
                                v132 := GenS2{F0: (1 < genFn0((fi120 - (fi120 / ((fi120) | 1))), v2))}
                                var v133 int64 = 42
                                var v134 float64 = (float64(1.5) + 1e10)
                                _ = v131
                                _ = v132
                                _ = v133
                                _ = v134
                            } else {
                                if v2 {
                                    v135 := GenS16{F0: 127, F1: 42, F2: 128}
                                    _ = v135
                                } else {
                                    v136 := GenS4{F0: 127, F1: 285}
                                    v137 := GenS17{F0: 2, F1: ((0 >= 98) && v2), F2: 42}
                                    var v138 bool = v2
                                    v1[fi119] = v138
                                    if true {
                                        if (1 != 2) {
                                            v3 = v3
                                            v19 = v19
                                            v3 -= (v3 & int64(fi120))
                                            if v2 {
                                                v19 -= -1
                                                v19 -= v19
                                                if (-1 > 48) {
                                                    v139 := GenS18{F0: (v19 * v19), F1: ("foo bar" + "a")}
                                                    _ = v139
                                                }
                                            }
                                        }
                                        switch (fi119 - (^(fi119))) {
                                        case 2147483647:
                                            v1[fi119] = v138
                                        case 128:
                                            var v140 int64 = 1
                                            _ = v140
                                        case 9223372036854775807:
                                            v138 = v138
                                        }
                                        switch fi119 {
                                        case 4294967295:
                                            if true {
                                                v141 := GenS1{F0: 2147483647, F1: 60, F2: (fi120 % ((fi119) | 1))}
                                                v142 := GenS8{F0: 65535, F1: 18446744073709551615, F2: 42}
                                                switch "a" {
                                                case "a":
                                                    _ = len(v124)
                                                case "\n\t":
                                                    v138 = false
                                                case "foo bar":
                                                    var v143 uint32 = 721
                                                    _ = v143
                                                default:
                                                    v124 = append(v124, fi120)
                                                }
                                                _ = v141
                                                _ = v142
                                            } else {
                                                var v144 uint16 = 281
                                                v145 := GenS17{F0: 1, F1: (v2 && (!v2)), F2: 127}
                                                _ = v144
                                                _ = v145
                                            }
                                        case 127:
                                            v146 := GenS5{F0: v19, F1: -842}
                                            _ = v146
                                        case 9223372036854775807:
                                            v3 = v3
                                        default:
                                            var v147 int32 = -549
                                            var v148 uint64 = 453
                                            _ = v147
                                            _ = v148
                                        }
                                    } else {
                                        var v149 uint = 4294967295
                                        v150 := GenS3{F0: 158, F1: 1}
                                        _ = v149
                                        _ = v150
                                    }
                                    _ = v136
                                    _ = v137
                                    _ = v138
                                }
                                if (!v2) {
                                    var v151 int = (fi120 + -612)
                                    _ = v151
                                }
                            }
                        }
                    }
                    _ = v125
                } else {
                    if ((fi120 % ((0) | 1)) >= fi119) {
                        fmt.Println(fi120, v3)
                        fmt.Println(v3)
                    }
                }
                fmt.Println(v3)
                v2 = (v3 != v3)
                var v152 uint16 = 481
                _ = v124
                _ = v152
            }
            _ = fi120
        }
        var v153 uint16 = 105
        _ = fi119
        _ = v153
    }
    var v154 int8 = 100
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v154:", v154)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v154
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v51
}
