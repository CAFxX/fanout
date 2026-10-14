package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 uint32; F2 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint; F1 uint32 }
type GenS6 struct { F0 bool; F1 uint16; F2 int }
type GenS7 struct { F0 int64; F1 uint }
type GenS8 struct { F0 int8; F1 int16 }
type GenS9 struct { F0 int; F1 uint32 }
type GenS10 struct { F0 int8; F1 bool }
type GenS11 struct { F0 int64; F1 int16; F2 uint32 }
type GenS12 struct { F0 int32; F1 int8 }
type GenS13 struct { F0 uint8; F1 uint16; F2 string }
type GenS14 struct { F0 int64; F1 string; F2 int8 }
type GenS15 struct { F0 string; F1 bool }
type GenS16 struct { F0 int8; F1 int32; F2 int64 }
type GenS17 struct { F0 int8; F1 uint32; F2 int32 }
type GenS18 struct { F0 int8; F1 bool }
type GenS19 struct { F0 uint16 }
func genFn0(p0 int, p1 uint64, p2 string) bool {
    var v4 string = ""
    if true {
        for fi5 := 0; fi5 < 2; fi5++ {
            v6 := GenS1{F0: 465}
            v7 := []int{564, fi5, (p0 | fi5), (p0 | (fi5 + 32767))}
            v7 = append(v7, (p0 - p0))
            _ = fi5
            _ = v6
            _ = v7
        }
        if ((true && ((false && (-1 != -1)) && (2147483647 > 580))) || (!(-782 == 2147483647))) {
            v8 := []int{}
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = v8
        } else {
            switch p2 {
            case "hello":
                switch p0 {
                case -2147483648:
                    v9 := GenS2{F0: true}
                    _ = v9
                case 4294967295:
                    if (758 != -221) {
                        fmt.Println(v4)
                    } else {
                        v10 := []int{p0}
                        v11 := GenS3{F0: (-1 > -32768), F1: 65535, F2: (false || false)}
                        _ = v10
                        _ = v11
                    }
                }
            case "":
                v12 := []int{p0}
                _ = v12
            default:
                switch v4 {
                case "\n\t":
                    var v13 bool = false
                    _ = v13
                case "a":
                    if ((589 < 9223372036854775807) || (float64((1e10 + -2.25)) == float64(1.5))) {
                        for fi14 := 0; fi14 < 11; fi14++ {
                            var v15 int16 = 42
                            var v16 int64 = -9223372036854775808
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                        for fi17 := 0; fi17 < 9; fi17++ {
                            for fi18 := 0; fi18 < 20; fi18++ {
                                p0 -= -705
                                _ = fi18
                            }
                            var v19 uint16 = 881
                            _ = fi17
                            _ = v19
                        }
                    } else {
                        v20 := []int{1, p0, p0, p0}
                        v21 := GenS1{F0: 42}
                        v22 := GenS1{F0: -1}
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                default:
                    if false {
                        switch (p0 & 4294967295) {
                        case 4294967295:
                            if (!true) {
                                var v23 float32 = 0.0
                                if true {
                                    v24 := GenS4{F0: 119}
                                    _ = v24
                                } else {
                                    fmt.Println(p0, p1, p2)
                                    fmt.Println(p1, v4)
                                    if (1000 == 1000) {
                                        p0 -= (-(p0))
                                        var v25 uint = 255
                                        _ = v25
                                    } else {
                                        p1 = p1
                                    }
                                }
                                switch 801 {
                                case -1:
                                    p2 = p2
                                case -17:
                                    if false {
                                        var v26 bool = (0.0 > (1.5 + (-2.25 / 1.5)))
                                        if (!(42 <= 42)) {
                                            var v27 uint32 = 4294967295
                                            _ = v27
                                        }
                                        _ = v26
                                    } else {
                                        v28 := []int{}
                                        for fi29 := 0; fi29 < 14; fi29++ {
                                            var v30 uint8 = 127
                                            _ = fi29
                                            _ = v30
                                        }
                                        p1 -= p1
                                        _ = v28
                                    }
                                case 2147483647:
                                    v4 = ((p2 + (v4 + ("foo bar" + v4))) + (v4 + p2))
                                default:
                                    var v31 int8 = 100
                                    _ = v31
                                }
                                _ = v23
                            }
                        case 32768:
                            p1 = (p1 ^ p1)
                        case 128:
                            var v32 int32 = -2147483648
                            _ = v32
                        }
                        var v33 int = p0
                        _ = v33
                    }
                }
            }
        }
        v34 := GenS5{F0: 181, F1: 657}
        _ = v34
    }
    p2 = p2
    v35 := GenS4{F0: uint32(p1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v35
    _ = v4
    return (!false)
}
func genFn1() int8 {
    if true {
        var v36 uint16 = 0
        v36 = (v36 + 364)
        v36 += v36
        _ = v36
    } else {
        v37 := GenS6{F0: (42 <= 756), F1: 65535, F2: 32768}
        var v38 uint64 = 2
        _ = v37
        _ = v38
    }
    if true {
        var v39 int8 = -29
        var v40 float64 = (float64(((3.14159 * 3.14159) + 0.0)) / float64(0.1))
        _ = v39
        _ = v40
    }
    var v41 uint = 65535
    var v42 uint8 = 2
    _ = v41
    _ = v42
    return -128
}

func main() {
    var v1 int = 698
    v2 := make(map[uint32]uint32)
    var v3 int64 = 106
    for fi43 := 0; fi43 < 12; fi43++ {
        v44 := []int{v1}
        v45 := GenS5{F0: 18446744073709551615, F1: 2147483647}
        for fi46 := 0; fi46 < 6; fi46++ {
            v3 += 42
            _ = fi46
        }
        _ = fi43
        _ = v44
        _ = v45
    }
    if (!false) {
        delete(v2, 2147483648)
        for fi47 := 0; fi47 < 3; fi47++ {
            v48 := GenS7{F0: v3, F1: 2}
            if ((1e10 + ((0.0 * ((float64((1.5 + 1.5)) + 0.0) / -2.25)) - 0.0)) < 0.1) {
                switch ("" + (("hello" + "\n\t") + "foo bar")) {
                case "a":
                    if true {
                        var v49 string = "hello"
                        switch (v1 + 2147483647) {
                        case 1:
                            v50 := GenS2{F0: true}
                            _ = v50
                        case -476:
                            if true {
                                if true {
                                    v51 := GenS5{F0: 2, F1: 294}
                                    v1 += fi47
                                    _ = v51
                                } else {
                                    for fi52 := 0; fi52 < 6; fi52++ {
                                        v53 := GenS8{F0: 127, F1: 0}
                                        _ = fi52
                                        _ = v53
                                    }
                                    if false {
                                        if false {
                                            switch v49 {
                                            case "\n\t":
                                                v3 = 1
                                            }
                                            if (!(3.14159 == 1.5)) {
                                                var v54 bool = true
                                                v55 := v2[42]
                                                for fi56 := 0; fi56 < 3; fi56++ {
                                                    v57 := GenS9{F0: (-(v1)), F1: (v55 / (((^(v55))) | 1))}
                                                    _ = fi56
                                                    _ = v57
                                                }
                                                _ = v54
                                                _ = v55
                                            }
                                        }
                                    } else {
                                        v58 := GenS10{F0: 1, F1: false}
                                        v59 := GenS5{F0: 310, F1: 65535}
                                        _ = v58
                                        _ = v59
                                    }
                                }
                                for fi60 := 0; fi60 < 18; fi60++ {
                                    if (false && (!(0 != 239))) {
                                        var v61 bool = true
                                        if (787 > 2147483648) {
                                            v62 := v2[2147483647]
                                            var v63 uint16 = 931
                                            var v64 int8 = -23
                                            _ = v62
                                            _ = v63
                                            _ = v64
                                        } else {
                                            switch ("\n\t" + v49) {
                                            case "foo bar":
                                                var v65 float32 = -2.25
                                                _ = v65
                                            }
                                            v49 = v49
                                            switch "" {
                                            case "foo bar":
                                                fmt.Println(fi47)
                                            case "\n\t":
                                                if v61 {
                                                    if (1 < 974) {
                                                        if ((v3 % (((v3 * v3)) | 1)) != (^(v3))) {
                                                            var v66 int32 = 2147483647
                                                            _ = v66
                                                        }
                                                        switch v49 {
                                                        case "":
                                                            fmt.Println(v49, v49)
                                                        case "a":
                                                            var v67 uint32 = 1
                                                            _ = v67
                                                        default:
                                                            fmt.Println(v3, v61)
                                                        }
                                                        var v68 uint32 = 4294967295
                                                        _ = v68
                                                    } else {
                                                        if true {
                                                            v69 := GenS4{F0: 2147483648}
                                                            v1 = (fi47 << ((876) & 63))
                                                            _ = v69
                                                        }
                                                    }
                                                }
                                            case "hello":
                                                if v61 {
                                                    v70 := GenS3{F0: (!(0.1 == 0.1)), F1: 65535, F2: ((1 != -904) && ((v1 | 42) < (^(fi47))))}
                                                    v71 := GenS3{F0: true, F1: 4294967295, F2: v61}
                                                    v72 := GenS10{F0: -1, F1: false}
                                                    _ = v70
                                                    _ = v71
                                                    _ = v72
                                                } else {
                                                    var v73 float64 = float64(1e10)
                                                    switch 32767 {
                                                    case -702:
                                                        v73 += float64(-2.25)
                                                    case 9223372036854775807:
                                                        v61 = v61
                                                        if v61 {
                                                            v49 = (v49 + "hello")
                                                            v74 := v2[786]
                                                            fmt.Println(v74, v73, fi47)
                                                            _ = v74
                                                        } else {
                                                            var v75 int64 = v3
                                                            v76 := []int{504, v1}
                                                            _ = v75
                                                            _ = v76
                                                        }
                                                        var v77 float32 = 1e10
                                                        _ = v77
                                                    case 656:
                                                        if (175 == 4294967295) {
                                                            v78 := GenS6{F0: (true && (!false)), F1: 255, F2: fi47}
                                                            v3 -= (v3 ^ (v3 * (^(v3))))
                                                            if (1 < 470) {
                                                                v79 := GenS11{F0: (v3 * (^(v3))), F1: 905, F2: 65535}
                                                                var v80 string = v49
                                                                var v81 float64 = (v73 - 0.0)
                                                                v82 := GenS12{F0: 0, F1: -81}
                                                                v80 = v80
                                                                v83 := v2[313]
                                                                _ = v79
                                                                _ = v80
                                                                _ = v81
                                                                _ = v82
                                                                _ = v83
                                                            }
                                                            _ = v78
                                                        }
                                                    }
                                                    v61 = (42 >= 470)
                                                    _ = v73
                                                }
                                            }
                                            var v84 float32 = (0.1 - (1.5 + (3.14159 - 0.0)))
                                            v61 = v61
                                            delete(v2, 2147483648)
                                            _ = len(v2)
                                            if (!(!v61)) {
                                                var v85 int64 = v3
                                                _ = v85
                                            }
                                            _ = v84
                                        }
                                        _ = v61
                                    }
                                    v3 -= 9223372036854775807
                                    _ = fi60
                                }
                            } else {
                                fmt.Println(v49, fi47)
                            }
                        case 32767:
                            if false {
                                v86 := GenS12{F0: 723, F1: -94}
                                if (true || (!true)) {
                                    for fi87 := 0; fi87 < 12; fi87++ {
                                        v49 = ((("" + "foo bar") + "\n\t") + ((v49 + (v49 + "hello")) + v49))
                                        var v88 string = v49
                                        _ = fi87
                                        _ = v88
                                    }
                                }
                                var v89 uint8 = 0
                                _ = v86
                                _ = v89
                            }
                        default:
                            var v90 string = v49
                            _ = v90
                        }
                        v91 := []int{713, 252, fi47, -543}
                        _ = v49
                        _ = v91
                    } else {
                        switch (-(fi47)) {
                        case 127:
                            v3 += -753
                        }
                    }
                case "":
                    if (4294967295 <= 433) {
                        var v92 uint64 = 266
                        _ = v92
                    }
                }
                v3 = (-(v3))
            }
            var v93 uint = 65535
            _ = fi47
            _ = v48
            _ = v93
        }
        fmt.Println(v3, v1)
    }
    v94 := []int{(v1 + (v1 >> ((364) & 63)))}
    for fi95 := 0; fi95 < 5; fi95++ {
        if false {
            v96 := GenS2{F0: false}
            if false {
                for fi97 := 0; fi97 < 9; fi97++ {
                    if (127 < 127) {
                        var v98 int = (fi97 * 42)
                        _ = v98
                    } else {
                        if ((!((!(!(!false))) || (-66 == genFn1()))) || false) {
                            v99 := GenS8{F0: -30, F1: -32768}
                            fmt.Println(fi97)
                            var v100 int8 = 42
                            _ = v100
                            _ = v99
                        }
                    }
                    _ = fi97
                }
                var v101 uint = 65535
                _ = v101
            } else {
                v3 += v3
                var v102 string = "\n\t"
                v3 = 470
                _ = v102
            }
            v103 := v2[215]
            _ = v103
            _ = v96
        } else {
            if true {
                fmt.Println(v3)
            }
            if (true && (-32768 < 32767)) {
                v3 -= (v3 - v3)
            }
            var v104 bool = true
            _ = v104
        }
        _ = fi95
    }
    v1 += (v1 % (((v1 + (v1 >> ((4294967295) & 63)))) | 1))
    if false {
        for fi105 := 0; fi105 < 6; fi105++ {
            for fi106 := 0; fi106 < 9; fi106++ {
                switch (v1 - (fi106 & (fi105 >> ((507) & 63)))) {
                case 1:
                    v1 -= -9223372036854775808
                case 7:
                    if (false || false) {
                        if false {
                            switch fi105 {
                            case -1:
                                v1 = fi105
                            case 9223372036854775807:
                                if (3.14159 != 3.14159) {
                                    var v107 float64 = float64(1.5)
                                    v3 += v3
                                    _ = v107
                                } else {
                                    delete(v2, 2147483648)
                                    v3 += (v3 << ((429) & 63))
                                }
                            case -257:
                                var v108 int64 = v3
                                _ = v108
                            }
                            v109 := GenS13{F0: 1, F1: 0, F2: "\n\t"}
                            _ = v109
                        } else {
                            v110 := GenS5{F0: 705, F1: 4294967295}
                            if (840 >= 18446744073709551615) {
                                v111 := GenS14{F0: 9223372036854775807, F1: "", F2: 100}
                                if (1 > genFn1()) {
                                    var v112 int32 = -882
                                    if (!(-1 >= -1)) {
                                        _ = len(v2)
                                    } else {
                                        var v113 uint32 = 281
                                        _ = v113
                                    }
                                    v112 -= -537
                                    _ = v112
                                }
                                v1 += (v1 * 475)
                                _ = v111
                            }
                            if false {
                                if ((("\n\t" + "hello") + "hello") < "\n\t") {
                                    v114 := GenS15{F0: "", F1: false}
                                    if (103 == -26) {
                                        v115 := GenS11{F0: v3, F1: 1, F2: 740}
                                        var v116 uint = 1
                                        v1 = fi106
                                        v116 = (v116 % ((18446744073709551615) | 1))
                                        v3 += v3
                                        _ = v115
                                        _ = v116
                                    } else {
                                        v1 += (v1 >> ((18446744073709551615) & 63))
                                        v3 -= (^(v3))
                                    }
                                    var v117 float64 = 3.14159
                                    _ = v114
                                    _ = v117
                                } else {
                                    switch (^(v1)) {
                                    case -72:
                                        switch fi105 {
                                        case -321:
                                            v118 := GenS6{F0: false, F1: 725, F2: (fi105 * (^(fi105)))}
                                            _ = v118
                                        default:
                                            v119 := GenS14{F0: (v3 | (v3 % (((v3 >> ((18446744073709551615) & 63))) | 1))), F1: ("foo bar" + ""), F2: -128}
                                            _ = v119
                                        }
                                    }
                                    v94 = append(v94, (v1 / ((v1) | 1)))
                                }
                                var v120 int16 = -224
                                _ = v120
                            }
                            switch v1 {
                            case 1:
                                v121 := GenS16{F0: 65, F1: 42, F2: 0}
                                _ = v121
                            default:
                                var v122 bool = ((33 != 8) || (-294 >= 2147483647))
                                _ = v122
                            }
                            _ = v110
                        }
                        v2[1] = 42
                        var v123 bool = ((((true && (true || true)) && (false || false)) && false) && true)
                        switch 128 {
                        case -1:
                            if (!v123) {
                                var v124 uint32 = 4294967295
                                fmt.Println(fi105, fi106)
                                _ = v124
                            }
                        case 128:
                            if (65535 == 931) {
                                fmt.Println(v3)
                                if len(v94) > 0 {
                                    v94[0] = 1
                                }
                                v1 -= 9223372036854775807
                            } else {
                                var v125 int = fi106
                                var v126 float64 = 3.14159
                                v3 += -1
                                _ = v125
                                _ = v126
                            }
                        default:
                            var v127 int16 = -697
                            _ = v127
                        }
                        _ = v123
                    } else {
                        v1 -= (-(fi106))
                    }
                    _ = len(v94)
                    var v128 string = "foo bar"
                    _ = v128
                default:
                    if true {
                        var v129 float32 = 0.0
                        var v130 uint8 = 127
                        v131 := GenS17{F0: -1, F1: 65535, F2: 42}
                        _ = v129
                        _ = v130
                        _ = v131
                    } else {
                        var v132 int16 = -1
                        var v133 int64 = (v3 >> ((2) & 63))
                        _ = v132
                        _ = v133
                    }
                }
                var v134 uint32 = 924
                fmt.Println(fi105)
                _ = fi106
                _ = v134
            }
            for fi135 := 0; fi135 < 10; fi135++ {
                v3 += 1
                switch v1 {
                case -173:
                    if false {
                        v136 := GenS11{F0: (v3 ^ v3), F1: 32767, F2: 2147483648}
                        var v137 uint = 18446744073709551615
                        v138 := GenS18{F0: -128, F1: ((fi105 & fi135) == fi135)}
                        _ = v136
                        _ = v137
                        _ = v138
                    }
                case -912:
                    var v139 int16 = -32768
                    if false {
                        if true {
                            switch 752 {
                            case 1:
                                if true {
                                    v140 := v2[42]
                                    var v141 int32 = -115
                                    v139 += v139
                                    var v142 uint32 = v140
                                    if (!(true && false)) {
                                        v140 = v142
                                    } else {
                                        var v143 bool = (498 != 42)
                                        v144 := v2[(v142 ^ (v142 - (v140 % ((v140) | 1))))]
                                        _ = v143
                                        _ = v144
                                    }
                                    switch (v1 | (fi135 ^ (fi135 - (fi105 / ((fi105) | 1))))) {
                                    case 127:
                                        var v145 uint16 = 2
                                        _ = v145
                                    case -2147483648:
                                        v146 := GenS19{F0: 1}
                                        _ = v146
                                    }
                                    _ = v140
                                    _ = v141
                                    _ = v142
                                } else {
                                    var v147 string = ("hello" + "\n\t")
                                    _ = v147
                                }
                            case 0:
                                fmt.Println(fi135, fi105)
                            case 826:
                                fmt.Println(v139, v139, fi105)
                            }
                        }
                    }
                    _ = v139
                default:
                    fmt.Println(v1, fi135)
                }
                _ = fi135
            }
            _ = fi105
        }
    }
    var v148 string = "\n\t"
    for fi149 := 0; fi149 < 13; fi149++ {
        v1 += (fi149 % ((fi149) | 1))
        _ = fi149
    }
    v148 = ""
    for fi150 := 0; fi150 < 7; fi150++ {
        v3 += 9223372036854775807
        _ = fi150
    }
    v151 := v2[42]
    var v152 uint32 = (v151 & (v151 + v151))
    v153 := GenS2{F0: false}
    fmt.Println("v1:", v1)
    fmt.Println("v148:", v148)
    fmt.Println("v151:", v151)
    fmt.Println("v152:", v152)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v94:", len(v94), v94)
    _ = v1
    _ = v148
    _ = v151
    _ = v152
    _ = v153
    _ = v2
    _ = v3
    _ = v94
}
