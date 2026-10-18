package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 uint16; F2 string }
type GenS4 struct { F0 uint64; F1 uint; F2 bool }
type GenS5 struct { F0 int32; F1 string }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int32; F1 uint8; F2 int16 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS12 struct { F0 string; F1 string; F2 int8 }
type GenS13 struct { F0 int64; F1 uint }
type GenS14 struct { F0 string; F1 uint; F2 uint32 }
type GenS15 struct { F0 string; F1 uint }
type GenS16 struct { F0 int64; F1 uint64; F2 uint16 }
type GenS17 struct { F0 int32; F1 uint16 }
func genFn0(p0 int8) uint64 {
    var v4 string = ("foo bar" + ("foo bar" + ("foo bar" + "")))
    switch v4 {
    case "foo bar":
        v5 := GenS1{F0: (p0 | (p0 % (((p0 % ((p0) | 1))) | 1)))}
        _ = v5
    case "":
        fmt.Println(v4, v4)
    default:
        switch 32767 {
        case -1:
            switch 2147483647 {
            case 72:
                switch 1 {
                case 2147483647:
                    v6 := GenS1{F0: (-(p0))}
                    _ = v6
                case 127:
                    if (false || false) {
                        var v7 uint64 = 820
                        var v8 int = -1
                        for fi9 := 0; fi9 < 14; fi9++ {
                            if false {
                                v10 := []int{v8, fi9}
                                _ = v10
                            }
                            v11 := GenS1{F0: -117}
                            _ = fi9
                            _ = v11
                        }
                        _ = v7
                        _ = v8
                    } else {
                        if (-32768 <= -1) {
                            v12 := GenS1{F0: (p0 & -1)}
                            _ = v12
                        } else {
                            var v13 bool = true
                            var v14 uint64 = 4294967295
                            _ = v13
                            _ = v14
                        }
                    }
                default:
                    if ((-331 >= -577) && true) {
                        fmt.Println(p0)
                        var v15 bool = ((("hello" + v4) > ("" + ("hello" + (v4 + (v4 + ((v4 + v4) + ("" + ((v4 + v4) + v4)))))))) && false)
                        _ = v15
                    } else {
                        v16 := GenS1{F0: p0}
                        for fi17 := 0; fi17 < 11; fi17++ {
                            v4 = "hello"
                            v4 = ""
                            _ = fi17
                        }
                        for fi18 := 0; fi18 < 12; fi18++ {
                            var v19 uint32 = 2147483648
                            _ = fi18
                            _ = v19
                        }
                        _ = v16
                    }
                }
            case -1:
                v20 := GenS2{F0: 9223372036854775807}
                _ = v20
            }
        case 127:
            if ((77 >= -1) || ((1 <= 4294967295) || true)) {
                var v21 int16 = 765
                _ = v21
            } else {
                for fi22 := 0; fi22 < 16; fi22++ {
                    v23 := []int{}
                    _ = fi22
                    _ = v23
                }
            }
        }
    }
    var v24 float64 = 0.0
    v24 += v24
    _ = p0
    _ = v24
    _ = v4
    return 4294967295
}
func genFn1(p0 uint64, p1 string, p2 uint8) uint64 {
    if true {
        for fi25 := 0; fi25 < 3; fi25++ {
            p0 = p0
            _ = fi25
        }
        p0 = (^(p0))
    } else {
        v26 := []int{}
        _ = v26
    }
    v27 := []int{-1}
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    return genFn0(1)
}

func main() {
    var v1 uint = 217
    var v2 uint8 = 128
    v3 := make(map[uint32]int)
    for fi28 := 0; fi28 < 1; fi28++ {
        v29 := GenS2{F0: 704}
        var v30 int16 = 699
        var v31 int16 = (v30 / ((v30) | 1))
        _ = fi28
        _ = v29
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 10; fi32++ {
        var v33 bool = (!true)
        _ = fi32
        _ = v33
    }
    v34 := []int{2147483647}
    if ((v2 == v2) || false) {
        v2 += (^(v2))
        fmt.Println(v2, v2)
        v1 += v1
    } else {
        var v35 string = ""
        var v36 uint8 = (v2 | 255)
        _ = v35
        _ = v36
    }
    var v37 uint16 = 0
    switch -541 {
    case -9223372036854775808:
        var v38 float64 = 3.14159
        _ = v38
    case 1:
        v34 = append(v34, -2147483648)
    }
    var v39 uint16 = 792
    switch -2147483648 {
    case 0:
        switch 127 {
        case 1:
            v40 := GenS3{F0: 438, F1: (v37 >> ((v1) & 63)), F2: "hello"}
            _ = v40
        case 19:
            if false {
                var v41 int16 = -136
                v41 = v41
                fmt.Println(v1, v39)
                _ = v41
            } else {
                if len(v34) >= 2 {
                    v34 = v34[1:len(v34)]
                }
                v42 := GenS4{F0: 266, F1: (v1 & v1), F2: (42 >= 127)}
                v43 := GenS4{F0: 104, F1: 1, F2: false}
                _ = v42
                _ = v43
            }
        }
    case 2147483647:
        v37 = 32767
    case 42:
        if false {
            v1 = v1
            if (v1 > 255) {
                v44 := GenS1{F0: 1}
                var v45 int32 = -2147483648
                _ = v44
                _ = v45
            } else {
                switch "a" {
                case "a":
                    fmt.Println(v2)
                case "\n\t":
                    var v46 int64 = 625
                    _ = v46
                case "hello":
                    var v47 uint8 = uint8(v1)
                    _ = v47
                }
                if ((v2 != 1) && true) {
                    v2 = v2
                    for fi48 := 0; fi48 < 12; fi48++ {
                        v49 := GenS2{F0: 9223372036854775807}
                        if true {
                            v1 += uint(v2)
                            for fi50 := 0; fi50 < 2; fi50++ {
                                var v51 float64 = (0.0 - ((-2.25 + 1e10) - 0.1))
                                v52 := GenS2{F0: 9223372036854775807}
                                _ = len(v3)
                                _ = fi50
                                _ = v51
                                _ = v52
                            }
                            var v53 uint = v1
                            _ = v53
                        }
                        switch fi48 {
                        case 572:
                            var v54 uint = v1
                            _ = v54
                        case 127:
                            var v55 int16 = 1
                            _ = v55
                        default:
                            if false {
                                if false {
                                    if ((((v37 + (^(v37))) < (v37 | (v37 ^ (v39 ^ 32767)))) && (fi48 != (-(fi48)))) && true) {
                                        if true {
                                            if (!(2147483648 > 0)) {
                                                v56 := GenS5{F0: -2147483648, F1: "hello"}
                                                v57 := GenS5{F0: 42, F1: "a"}
                                                var v58 float32 = (0.1 - float32(1.5))
                                                _ = v56
                                                _ = v57
                                                _ = v58
                                            }
                                        }
                                        if false {
                                            var v59 uint8 = v2
                                            var v60 uint16 = v39
                                            var v61 int16 = -195
                                            _ = v59
                                            _ = v60
                                            _ = v61
                                        } else {
                                            var v62 float32 = float32(((((3.14159 - ((0.1 - float64(((1.5 * 0.1) / 1e10))) * (1.5 + (((((0.1 + 0.1) * 1e10) - ((-2.25 - 3.14159) / 1e10)) / -2.25) * 0.0)))) * 0.1) - (0.0 + 0.1)) + -2.25))
                                            for fi63 := 0; fi63 < 7; fi63++ {
                                                v2 += v2
                                                _ = fi63
                                            }
                                            _ = v62
                                        }
                                        v39 = v39
                                    }
                                    var v64 int32 = int32(fi48)
                                    _ = v64
                                }
                                v2 += v2
                            }
                        }
                        _ = fi48
                        _ = v49
                    }
                }
                if (1.5 <= float32(1.5)) {
                    var v65 int8 = int8(v1)
                    fmt.Println(v65, v1, v37)
                    _ = v65
                } else {
                    if (float64((-2.25 / 1e10)) <= 1e10) {
                        v66 := GenS3{F0: 42, F1: v39, F2: (("a" + "") + (("a" + "") + "a"))}
                        switch 529 {
                        case 128:
                            fmt.Println(v1, v37, v2)
                        case -182:
                            if false {
                                switch -15 {
                                case 0:
                                    v67 := GenS6{F0: v1}
                                    _ = v67
                                }
                            }
                        case 42:
                            if (!true) {
                                if (((0.0 - 0.0) / (1.5 / 3.14159)) > 0.1) {
                                    var v68 int8 = 1
                                    var v69 float64 = 1.5
                                    _ = v68
                                    _ = v69
                                }
                                v37 -= (v37 * v37)
                            }
                        }
                        _ = v66
                    }
                    v37 = (v39 ^ 0)
                }
            }
            if true {
                var v70 int8 = 127
                var v71 string = "foo bar"
                var v72 uint64 = 2
                _ = v70
                _ = v71
                _ = v72
            } else {
                for fi73 := 0; fi73 < 10; fi73++ {
                    v74 := GenS7{F0: 1}
                    _ = fi73
                    _ = v74
                }
                switch "foo bar" {
                case "foo bar":
                    var v75 uint64 = 719
                    _ = v75
                case "\n\t":
                    fmt.Println(v39, v37, v2)
                case "a":
                    if (uint32(v1) >= 2) {
                        var v76 bool = true
                        v37 = (^(v37))
                        for fi77 := 0; fi77 < 3; fi77++ {
                            fmt.Println(v1)
                            _ = fi77
                        }
                        _ = v76
                    } else {
                        if ((1000 > 42) || (119 >= 127)) {
                            var v78 int = 9223372036854775807
                            v78 = v78
                            switch 128 {
                            case 9223372036854775807:
                                switch ("foo bar" + "hello") {
                                case "\n\t":
                                    if true {
                                        switch (^(v78)) {
                                        case -9223372036854775808:
                                            if false {
                                                var v79 uint64 = 513
                                                switch v78 {
                                                case 9223372036854775807:
                                                    v34 = append(v34, v78)
                                                case 4294967295:
                                                    if (-446 != 42) {
                                                        v2 -= (v2 * (v2 << (((v1 >> ((v1) & 63))) & 63)))
                                                    }
                                                }
                                                _ = v79
                                            } else {
                                                if (!false) {
                                                    v80 := GenS8{F0: uint64(v2)}
                                                    var v81 uint64 = 710
                                                    var v82 int = (v78 & (v78 << ((v1) & 63)))
                                                    _ = v80
                                                    _ = v81
                                                    _ = v82
                                                }
                                                v83 := GenS9{F0: -482, F1: (v2 + (v2 - v2)), F2: 42}
                                                _ = v83
                                            }
                                        case 127:
                                            v78 += v78
                                        default:
                                            for fi84 := 0; fi84 < 4; fi84++ {
                                                v85 := v3[4294967295]
                                                _ = fi84
                                                _ = v85
                                            }
                                        }
                                        for fi86 := 0; fi86 < 12; fi86++ {
                                            for fi87 := 0; fi87 < 9; fi87++ {
                                                v37 += v39
                                                _ = fi87
                                            }
                                            _ = fi86
                                        }
                                    } else {
                                        if true {
                                            v88 := GenS5{F0: -307, F1: (((("hello" + "a") + "a") + "foo bar") + "foo bar")}
                                            fmt.Println(v2, v39)
                                            v39 = (v37 & v39)
                                            _ = v88
                                        } else {
                                            v2 = uint8(v37)
                                        }
                                        _ = len(v34)
                                    }
                                case "a":
                                    if (!false) {
                                        fmt.Println(v39)
                                        var v89 uint16 = (v37 << (((^(v1))) & 63))
                                        switch (-(v78)) {
                                        case 42:
                                            v89 -= 1
                                        case -1:
                                            v90 := GenS5{F0: -422, F1: ""}
                                            _ = v90
                                        case 32768:
                                            var v91 float32 = 0.0
                                            _ = v91
                                        default:
                                            v1 = v1
                                        }
                                        _ = v89
                                    } else {
                                        if (-1 == int32(v2)) {
                                            var v92 bool = true
                                            v93 := GenS7{F0: 4294967295}
                                            _ = v92
                                            _ = v93
                                        }
                                    }
                                default:
                                    if len(v34) > 0 {
                                        v34[0] = v78
                                    }
                                }
                            case 128:
                                fmt.Println(v78, v2)
                            case 1:
                                if ((225 <= -2147483648) || false) {
                                    if false {
                                        v94 := GenS10{F0: 100}
                                        v95 := GenS3{F0: 2147483647, F1: (v37 + (v37 / ((v39) | 1))), F2: "foo bar"}
                                        _ = v94
                                        _ = v95
                                    } else {
                                        v39 -= (v37 & v37)
                                    }
                                }
                            default:
                                v1 = 2
                            }
                            _ = v78
                        } else {
                            if (-72 > 13) {
                                if ("" < ("a" + "foo bar")) {
                                    fmt.Println(v1, v2, v37)
                                } else {
                                    v96 := GenS10{F0: 100}
                                    for fi97 := 0; fi97 < 1; fi97++ {
                                        for fi98 := 0; fi98 < 14; fi98++ {
                                            var v99 float32 = float32((0.0 / 1.0))
                                            v100 := GenS4{F0: 18446744073709551615, F1: (v1 ^ v1), F2: (true && (!true))}
                                            _ = fi98
                                            _ = v100
                                            _ = v99
                                        }
                                        for fi101 := 0; fi101 < 6; fi101++ {
                                            v1 = v1
                                            _ = fi101
                                        }
                                        _ = fi97
                                    }
                                    _ = v96
                                }
                                _ = len(v34)
                                v102 := GenS5{F0: 2147483647, F1: (("foo bar" + "hello") + ((("a" + (((("" + "\n\t") + "hello") + "a") + ("\n\t" + ""))) + ("" + (("\n\t" + "\n\t") + "a"))) + ""))}
                                _ = v102
                            }
                            if true {
                                if len(v34) > 0 {
                                    v34[0] = -9223372036854775808
                                }
                                for fi103 := 0; fi103 < 17; fi103++ {
                                    v104 := GenS11{F0: 65535, F1: 4294967295, F2: 0}
                                    for fi105 := 0; fi105 < 2; fi105++ {
                                        var v106 uint16 = 65535
                                        v107 := GenS5{F0: 16, F1: ""}
                                        var v108 float32 = float32(((3.14159 / float64(0.1)) * 1e10))
                                        _ = fi105
                                        _ = v106
                                        _ = v107
                                        _ = v108
                                    }
                                    v2 -= (v2 + 197)
                                    _ = fi103
                                    _ = v104
                                }
                                v109 := GenS12{F0: ((("foo bar" + "") + (("a" + ("hello" + (("" + (("" + "") + ("\n\t" + "foo bar"))) + "a"))) + (("foo bar" + ("" + (("a" + "hello") + "\n\t"))) + "a"))) + "\n\t"), F1: "\n\t", F2: 1}
                                _ = v109
                            }
                            var v110 uint64 = 4294967295
                            _ = v110
                        }
                    }
                default:
                    v111 := GenS2{F0: -626}
                    _ = v111
                }
                for fi112 := 0; fi112 < 18; fi112++ {
                    var v113 int32 = -164
                    _ = fi112
                    _ = v113
                }
            }
        } else {
            v3[2147483648] = 530
            if (0.0 == 1.5) {
                v114 := GenS5{F0: 93, F1: "hello"}
                _ = v114
            } else {
                switch "" {
                case "hello":
                    var v115 int8 = 112
                    _ = v115
                default:
                    switch 128 {
                    case 128:
                        if true {
                            v1 = 1
                        } else {
                            v37 += uint16(v1)
                        }
                    case -975:
                        for fi116 := 0; fi116 < 5; fi116++ {
                            if false {
                                if (776 >= 1) {
                                    v117 := GenS12{F0: "", F1: "a", F2: 14}
                                    _ = v117
                                }
                                v118 := GenS13{F0: -972, F1: (v1 + v1)}
                                v1 -= 18446744073709551615
                                _ = v118
                            } else {
                                v1 += v1
                                for fi119 := 0; fi119 < 19; fi119++ {
                                    switch (fi116 / ((32768) | 1)) {
                                    case -214:
                                        v120 := GenS14{F0: "a", F1: 65535, F2: 486}
                                        _ = v120
                                    default:
                                        v2 += (^(v2))
                                    }
                                    fmt.Println(v2)
                                    _ = fi119
                                }
                            }
                            _ = fi116
                        }
                    case -588:
                        v121 := v3[2147483647]
                        _ = v121
                    default:
                        v1 = (v1 | 50)
                    }
                }
            }
            fmt.Println(v2, v37)
        }
    default:
        for fi122 := 0; fi122 < 18; fi122++ {
            for fi123 := 0; fi123 < 20; fi123++ {
                var v124 uint8 = (^(v2))
                var v125 float32 = 1e10
                _ = fi123
                _ = v124
                _ = v125
            }
            _ = fi122
        }
    }
    var v126 string = "a"
    var v127 int = 1
    for fi128 := 0; fi128 < 7; fi128++ {
        fmt.Println(v37)
        if false {
            v2 = uint8(v37)
            fmt.Println(v39, fi128)
            v3[2147483648] = -2147483648
        } else {
            var v129 int = v127
            _ = v129
        }
        _ = fi128
    }
    if false {
        if (-128 < -52) {
            v39 = v37
            v126 = ("foo bar" + ("hello" + v126))
            var v130 string = v126
            _ = v130
        }
        if false {
            for fi131 := 0; fi131 < 20; fi131++ {
                v132 := GenS15{F0: v126, F1: 273}
                v133 := v3[65535]
                _ = fi131
                _ = v132
                _ = v133
            }
            for fi134 := 0; fi134 < 10; fi134++ {
                v135 := v3[429]
                v2 = (v2 - 253)
                var v136 bool = false
                _ = fi134
                _ = v135
                _ = v136
            }
            if ((0.0 / -2.25) <= 1.5) {
                v2 = (^(v2))
                switch (v126 + v126) {
                case "foo bar":
                    var v137 uint8 = v2
                    _ = v137
                case "a":
                    v1 = 65535
                default:
                    v138 := GenS16{F0: int64(v1), F1: 703, F2: v37}
                    _ = v138
                }
                var v139 int8 = 127
                _ = v139
            }
        } else {
            fmt.Println(v39, v127, v127)
            fmt.Println(v39, v37, v37)
            if (483 != 0) {
                fmt.Println(v2, v127, v1)
                var v140 int64 = 9223372036854775807
                _ = v140
            }
        }
        if (-55 < -944) {
            if true {
                var v141 string = ("" + v126)
                _ = v141
            }
            var v142 int64 = 1
            v143 := v3[0]
            _ = v142
            _ = v143
        } else {
            switch v126 {
            case "a":
                fmt.Println(v1)
            case "":
                fmt.Println(v126, v2)
            default:
                var v144 bool = (false || true)
                _ = v144
            }
        }
    } else {
        if ((v127 | v127) < (v127 - -880)) {
            v126 = (v126 + v126)
            for fi145 := 0; fi145 < 3; fi145++ {
                var v146 bool = (true && (int64(v1) > 0))
                var v147 int8 = 14
                _ = fi145
                _ = v146
                _ = v147
            }
            v148 := GenS2{F0: -896}
            _ = v148
        }
        var v149 uint64 = 697
        for fi150 := 0; fi150 < 17; fi150++ {
            if false {
                for fi151 := 0; fi151 < 3; fi151++ {
                    v152 := GenS5{F0: 0, F1: ("\n\t" + (v126 + ("" + ("foo bar" + "\n\t"))))}
                    if true {
                        if false {
                            v153 := GenS4{F0: (v149 * genFn0(-120)), F1: v1, F2: false}
                            v1 -= 2
                            v37 += (^(v39))
                            _ = v153
                        } else {
                            if (true || true) {
                                v1 = v1
                                v1 = v1
                            } else {
                                v154 := GenS11{F0: (v1 << ((v1) & 63)), F1: 4294967295, F2: 0}
                                _ = v154
                            }
                        }
                        v149 = (^(v149))
                        v155 := GenS17{F0: 860, F1: (v39 & v39)}
                        _ = v155
                    }
                    _ = fi151
                    _ = v152
                }
            }
            _ = fi150
        }
        _ = v149
    }
    fmt.Println("v1:", v1)
    fmt.Println("v126:", v126)
    fmt.Println("v127:", v127)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v126
    _ = v127
    _ = v2
    _ = v3
    _ = v34
    _ = v37
    _ = v39
}
