package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 bool; F1 string; F2 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 int32; F1 int16; F2 int32 }
type GenS6 struct { F0 int64; F1 int8; F2 bool }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint8; F1 uint }
type GenS9 struct { F0 bool; F1 uint }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int8; F1 uint64 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int }
type GenS15 struct { F0 uint; F1 string }
type GenS16 struct { F0 int64; F1 int16; F2 uint32 }
type GenS17 struct { F0 uint8 }
type GenS18 struct { F0 int16 }
type GenS19 struct { F0 int; F1 int8; F2 string }
type GenS20 struct { F0 string; F1 bool; F2 uint32 }
func genFn0() uint16 {
    var v5 int16 = 0
    var v6 int32 = 42
    var v7 int = 32768
    _ = v5
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 int64, p1 int32, p2 uint64) int32 {
    v8 := []int{0}
    _ = len(v8)
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 4; fi10++ {
            v11 := GenS1{F0: 127, F1: (fi9 & int(p1))}
            p2 = 1
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    v8 = append(v8, 205)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 136
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = 1
    v3 := make(map[string]int)
    var v4 bool = false
    switch (("a" + "hello") + (("" + "a") + (("" + "\n\t") + ""))) {
    case "a":
        v12 := GenS2{F0: (628 <= 2147483647), F1: (("" + "\n\t") + ("" + "a")), F2: (v2 - 0)}
        _ = v12
    case "hello":
        v13 := []int{2147483647, -9223372036854775808, 394, -129}
        _ = v13
    default:
        fmt.Println(v2)
    }
    if v4 {
        v2 += (v2 / (((v2 & 2147483647)) | 1))
    } else {
        var v14 int64 = 761
        _ = v14
    }
    if true {
        var v15 float32 = -2.25
        _ = v15
    } else {
        for fi16 := 0; fi16 < 14; fi16++ {
            v17 := []int{}
            v18 := GenS3{F0: 1}
            _ = fi16
            _ = v17
            _ = v18
        }
        switch "" {
        case "foo bar":
            var v19 float64 = 0.0
            _ = v19
        case "a":
            var v20 uint32 = 2
            _ = v20
        default:
            switch "foo bar" {
            case "\n\t":
                v21 := v1[338]
                _ = v21
            }
        }
        if false {
            var v22 int64 = -145
            _ = v22
        }
    }
    v2 -= (v2 ^ 42)
    for fi23 := 0; fi23 < 8; fi23++ {
        for fi24 := 0; fi24 < 20; fi24++ {
            v2 -= v2
            switch "a" {
            case "\n\t":
                if ((!v4) || v4) {
                    v25 := GenS4{F0: (("" + "hello") + ("a" + (("\n\t" + "foo bar") + ""))), F1: 2}
                    v2 = genFn1(0, (v2 * -981), 2)
                    _ = v25
                } else {
                    switch "foo bar" {
                    case "\n\t":
                        v26 := GenS4{F0: ("hello" + "\n\t"), F1: 1}
                        _ = v26
                    }
                    _ = len(v3)
                }
            case "hello":
                if (!(fi23 == 1)) {
                    v2 = genFn1(484, (^(v2)), 42)
                    fmt.Println(fi24, v2, v2)
                } else {
                    v4 = (fi24 < 32767)
                    var v27 uint64 = 355
                    _ = v27
                }
            }
            _ = fi24
        }
        for fi28 := 0; fi28 < 1; fi28++ {
            if (336 < 460) {
                if (!true) {
                    var v29 uint64 = 173
                    switch "" {
                    case "a":
                        var v30 float32 = 1.5
                        _ = v30
                    case "\n\t":
                        v31 := []int{(fi28 / (((^(fi23))) | 1)), fi23, 4294967295, (fi28 | (fi23 % (((fi23 * (fi23 ^ -1))) | 1)))}
                        _ = v31
                    default:
                        var v32 float64 = (((3.14159 + 0.1) + 3.14159) / -2.25)
                        _ = v32
                    }
                    v4 = v4
                    _ = v29
                } else {
                    v2 = v2
                    var v33 uint16 = 218
                    if true {
                        fmt.Println(v4, v4)
                    } else {
                        v34 := []int{4294967295, 127}
                        v34 = append(v34, 1)
                        var v35 int = (-(fi28))
                        _ = v34
                        _ = v35
                    }
                    _ = v33
                }
                var v36 float32 = -2.25
                switch fi23 {
                case 32768:
                    var v37 uint16 = 928
                    _ = v37
                case 32767:
                    v38 := GenS5{F0: (^(v2)), F1: 1, F2: v2}
                    _ = v38
                default:
                    if (0 > 100) {
                        if true {
                            v4 = (128 > fi28)
                            v3[("foo bar" + "\n\t")] = fi28
                            if true {
                                switch (fi28 % ((fi23) | 1)) {
                                case 42:
                                    fmt.Println(v2)
                                case 360:
                                    v4 = false
                                case -243:
                                    v39 := GenS6{F0: -861, F1: -79, F2: (!(0 > 2))}
                                    _ = v39
                                }
                                fmt.Println(fi23, fi28, fi23)
                                v36 += ((float32((0.1 / 1.0)) * 3.14159) + ((v36 - 3.14159) - 1.5))
                            }
                        } else {
                            if (2147483647 >= fi23) {
                                if (!v4) {
                                    switch "\n\t" {
                                    case "foo bar":
                                        var v40 uint = 620
                                        _ = v40
                                    case "\n\t":
                                        v2 -= (v2 | 0)
                                    default:
                                        var v41 bool = v4
                                        v42 := GenS2{F0: (!(v4 && (800 <= 4294967295))), F1: "\n\t", F2: v2}
                                        var v43 int32 = 618
                                        v44 := GenS7{F0: (v43 >> ((420) & 63))}
                                        _ = v41
                                        _ = v42
                                        _ = v43
                                        _ = v44
                                    }
                                }
                                v3["\n\t"] = (fi23 | 42)
                            }
                            v45 := v3["hello"]
                            _ = v45
                        }
                        v2 -= v2
                        fmt.Println(v2, fi23)
                    } else {
                        if (v2 >= (^(v2))) {
                            var v46 uint32 = 1
                            _ = v46
                        } else {
                            v47 := GenS8{F0: 42, F1: 65535}
                            _ = v47
                        }
                        switch "foo bar" {
                        case "a":
                            if v4 {
                                switch (fi23 >> ((2) & 63)) {
                                case 127:
                                    v36 = float32(-2.25)
                                }
                                _ = len(v1)
                            } else {
                                v2 += v2
                            }
                        case "\n\t":
                            var v48 string = "\n\t"
                            _ = v48
                        case "foo bar":
                            var v49 uint = 2
                            var v50 int = (fi23 - fi23)
                            if (-192 > v2) {
                                var v51 float64 = (3.14159 * 1.5)
                                var v52 bool = false
                                switch (v50 * 128) {
                                case 0:
                                    fmt.Println(v52, v2)
                                case -2147483648:
                                    fmt.Println(v4)
                                default:
                                    if v52 {
                                        var v53 float64 = (1.5 - ((v51 / ((v51 - v51) - ((float64(v36) - (v51 / v51)) / 1.5))) * v51))
                                        _ = v53
                                    }
                                }
                                v54 := []int{-607}
                                _ = v51
                                _ = v52
                                _ = v54
                            } else {
                                v3[(("\n\t" + (("" + "hello") + "foo bar")) + "a")] = fi28
                            }
                            _ = v49
                            _ = v50
                        default:
                            switch 127 {
                            case 752:
                                v55 := GenS9{F0: true, F1: 42}
                                _ = v55
                            case 173:
                                v56 := v1[152]
                                _ = v56
                            case 2147483647:
                                delete(v1, 1)
                            }
                        }
                        if (fi23 <= fi23) {
                            v4 = true
                        } else {
                            switch ("hello" + "foo bar") {
                            case "a":
                                v4 = true
                            }
                        }
                    }
                }
                _ = v36
            }
            v57 := GenS8{F0: 2, F1: 42}
            _ = fi28
            _ = v57
        }
        _ = fi23
    }
    v4 = true
    for fi58 := 0; fi58 < 5; fi58++ {
        for fi59 := 0; fi59 < 3; fi59++ {
            if (!v4) {
                v60 := GenS1{F0: 42, F1: (fi59 >> ((313) & 63))}
                if true {
                    switch "hello" {
                    case "a":
                        if false {
                            fmt.Println(v4, v2, fi58)
                            v2 += genFn1(42, v2, 18446744073709551615)
                            fmt.Println(v4)
                        } else {
                            if v4 {
                                v61 := GenS10{F0: (!(((1 != 377) || v4) || (368 < 823)))}
                                v62 := GenS11{F0: 2}
                                _ = v61
                                _ = v62
                            }
                            v63 := []int{fi59, 4294967295, (fi59 - (fi58 << ((255) & 63)))}
                            _ = v63
                        }
                    }
                    delete(v3, ("" + "hello"))
                } else {
                    fmt.Println(v2)
                    v64 := GenS4{F0: "a", F1: 1}
                    _ = v64
                }
                _ = v60
            }
            _ = fi59
        }
        var v65 uint32 = 1
        _ = fi58
        _ = v65
    }
    v2 = v2
    v66 := []int{4294967295}
    if true {
        _ = len(v1)
        fmt.Println(v4, v2)
        for fi67 := 0; fi67 < 12; fi67++ {
            if false {
                switch fi67 {
                case 4294967295:
                    var v68 float32 = (((1.5 / 0.1) + 1.5) + -2.25)
                    _ = v68
                case 2147483647:
                    var v69 uint32 = 2147483648
                    _ = v69
                }
                if v4 {
                    for fi70 := 0; fi70 < 17; fi70++ {
                        var v71 float64 = 1e10
                        _ = fi70
                        _ = v71
                    }
                    if false {
                        if v4 {
                            var v72 uint64 = 480
                            if len(v66) > 0 {
                                v66[0] = 2147483647
                            }
                            if len(v66) > 0 {
                                v66[0] = 2147483647
                            }
                            _ = v72
                        } else {
                            if v4 {
                                switch "\n\t" {
                                case "a":
                                    var v73 int64 = 384
                                    _ = v73
                                case "":
                                    fmt.Println(fi67)
                                }
                            } else {
                                v66 = append(v66, 42)
                                fmt.Println(v4, fi67)
                            }
                            var v74 int64 = -528
                            if true {
                                var v75 float32 = 0.1
                                fmt.Println(v75, v2)
                                v76 := GenS7{F0: (-(v2))}
                                _ = v75
                                _ = v76
                            }
                            _ = v74
                        }
                        for fi77 := 0; fi77 < 18; fi77++ {
                            fmt.Println(v4, v2)
                            v3["foo bar"] = 2147483647
                            var v78 int16 = 23
                            _ = fi77
                            _ = v78
                        }
                        var v79 int32 = v2
                        _ = v79
                    }
                } else {
                    v2 -= v2
                }
            }
            _ = fi67
        }
    }
    for fi80 := 0; fi80 < 20; fi80++ {
        if (!v4) {
            var v81 int64 = -1
            var v82 bool = v4
            var v83 uint64 = 333
            _ = v81
            _ = v82
            _ = v83
        }
        if (true || v4) {
            if (2 <= 946) {
                var v84 int = 0
                _ = v84
            } else {
                var v85 bool = v4
                var v86 uint16 = 0
                _ = v85
                _ = v86
            }
            switch fi80 {
            case -9223372036854775808:
                v2 -= v2
            case 32768:
                if (139 > -32768) {
                    v66 = append(v66, fi80)
                    var v87 bool = (934 != 1)
                    _ = v87
                }
            case 42:
                var v88 float64 = 1e10
                _ = v88
            default:
                var v89 string = "a"
                _ = v89
            }
        } else {
            v4 = (!(4294967295 != 0))
        }
        fmt.Println(v2)
        _ = fi80
    }
    var v90 float64 = 1.5
    v1[65535] = 2
    _ = len(v3)
    for fi91 := 0; fi91 < 11; fi91++ {
        switch fi91 {
        case 549:
            v3["hello"] = 128
        case 42:
            delete(v1, 0)
        default:
            v4 = true
        }
        v2 = -330
        _ = fi91
    }
    v2 += v2
    fmt.Println(v2, v2, v4)
    for fi92 := 0; fi92 < 13; fi92++ {
        for fi93 := 0; fi93 < 9; fi93++ {
            var v94 int = fi92
            switch ("a" + "a") {
            case "hello":
                var v95 bool = v4
                _ = v95
            }
            v94 += -954
            _ = fi93
            _ = v94
        }
        switch "a" {
        case "hello":
            for fi96 := 0; fi96 < 4; fi96++ {
                v4 = true
                v2 = v2
                _ = fi96
            }
        case "foo bar":
            for fi97 := 0; fi97 < 6; fi97++ {
                v98 := GenS5{F0: v2, F1: 542, F2: v2}
                var v99 string = "\n\t"
                switch fi92 {
                case -2147483648:
                    if v4 {
                        if v4 {
                            if v4 {
                                v66 = append(v66, fi97)
                                var v100 string = v99
                                switch (fi92 & fi97) {
                                case -9223372036854775808:
                                    var v101 float32 = (-2.25 - (0.1 - (-2.25 + 1.5)))
                                    _ = v101
                                case 128:
                                    var v102 int8 = -43
                                    _ = v102
                                case 127:
                                    v90 += v90
                                    v103 := GenS12{F0: 11, F1: 674}
                                    var v104 uint = 0
                                    _ = v103
                                    _ = v104
                                default:
                                    switch (fi92 / ((fi97) | 1)) {
                                    case 32768:
                                        switch fi92 {
                                        case 9223372036854775807:
                                            delete(v3, v100)
                                        case -2147483648:
                                            fmt.Println(fi92)
                                        default:
                                            v105 := GenS11{F0: 42}
                                            _ = v105
                                        }
                                    default:
                                        fmt.Println(fi97)
                                    }
                                }
                                _ = v100
                            } else {
                                if (!v4) {
                                    v99 = v99
                                    switch (fi97 % ((53) | 1)) {
                                    case -9223372036854775808:
                                        var v106 string = (v99 + (v99 + ("a" + ((v99 + v99) + ("a" + v99)))))
                                        _ = v106
                                    case 32768:
                                        if (((((fi92 == (fi92 * fi97)) || false) && v4) || v4) || (838 != 870)) {
                                            v107 := GenS13{F0: 143}
                                            var v108 int64 = int64(v2)
                                            if (v4 && v4) {
                                                v109 := GenS9{F0: (65535 <= 255), F1: 255}
                                                if (0 != int16(v108)) {
                                                    v110 := GenS13{F0: 436}
                                                    _ = v110
                                                }
                                                if (0.1 < (v90 - -2.25)) {
                                                    v108 -= v108
                                                }
                                                _ = v109
                                            } else {
                                                if false {
                                                    var v111 int32 = (^(v2))
                                                    fmt.Println(v99)
                                                    v112 := GenS6{F0: (v108 + v108), F1: -114, F2: v4}
                                                    _ = v111
                                                    _ = v112
                                                }
                                            }
                                            _ = v107
                                            _ = v108
                                        }
                                    default:
                                        if ((0.1 + -2.25) > 1.5) {
                                            var v113 float32 = 3.14159
                                            if (-9223372036854775808 >= 1) {
                                                v113 = v113
                                                var v114 uint8 = 128
                                                _ = v114
                                            } else {
                                                v90 = v90
                                            }
                                            _ = v113
                                        }
                                    }
                                    v2 -= v2
                                }
                                v4 = v4
                                switch "\n\t" {
                                case "hello":
                                    if ((float64((((0.0 * 0.1) * -2.25) * 0.0)) == v90) && v4) {
                                        var v115 int = fi92
                                        switch ((v99 + v99) + ((v99 + (v99 + ((v99 + ((v99 + "\n\t") + v99)) + "\n\t"))) + v99)) {
                                        case "":
                                            v116 := GenS10{F0: true}
                                            _ = v116
                                        case "a":
                                            switch (fi97 * fi92) {
                                            case 630:
                                                var v117 uint64 = 18446744073709551615
                                                _ = v117
                                            case 4294967295:
                                                v118 := GenS14{F0: fi92}
                                                _ = v118
                                            case 42:
                                                v2 = v2
                                            }
                                        default:
                                            v119 := GenS15{F0: 312, F1: "hello"}
                                            _ = v119
                                        }
                                        v120 := GenS12{F0: 33, F1: 4294967295}
                                        if len(v66) >= 2 {
                                            v66 = v66[1:len(v66)]
                                        }
                                        _ = v115
                                        _ = v120
                                    } else {
                                        v4 = (("hello" + (v99 + v99)) == ("" + (v99 + (("hello" + v99) + "a"))))
                                        if ((-79 > 1) || v4) {
                                            v4 = (((18446744073709551615 <= 0) && (fi92 == fi97)) || ((v2 >> ((4294967295) & 63)) >= v2))
                                            v2 = v2
                                            if (1 == -275) {
                                                v121 := GenS16{F0: 0, F1: -244, F2: 1}
                                                v122 := GenS5{F0: (v2 + 128), F1: 1000, F2: (v2 % (((-(v2))) | 1))}
                                                v123 := GenS16{F0: -386, F1: 1000, F2: 2147483648}
                                                _ = v121
                                                _ = v122
                                                _ = v123
                                            }
                                        } else {
                                            switch (fi97 % ((fi92) | 1)) {
                                            case -1:
                                                if (152 >= 255) {
                                                    var v124 int16 = 0
                                                    if ((1e10 + ((v90 * 1e10) + (-2.25 / v90))) != float64(-2.25)) {
                                                        v125 := v3[v99]
                                                        if v4 {
                                                            v126 := v1[65535]
                                                            v127 := GenS12{F0: 100, F1: 42}
                                                            _ = v126
                                                            _ = v127
                                                        } else {
                                                            fmt.Println(v124, v99)
                                                            var v128 float32 = 1.5
                                                            v4 = false
                                                            _ = v128
                                                        }
                                                        _ = v125
                                                    } else {
                                                        var v129 uint16 = 0
                                                        _ = v129
                                                    }
                                                    v4 = v4
                                                    _ = v124
                                                } else {
                                                    v1[2147483648] = 4294967295
                                                    var v130 int64 = 9223372036854775807
                                                    v131 := GenS17{F0: 128}
                                                    _ = v130
                                                    _ = v131
                                                }
                                            default:
                                                v66 = append(v66, -559)
                                            }
                                            if v4 {
                                                var v132 uint8 = 1
                                                var v133 bool = (!(!true))
                                                _ = v132
                                                _ = v133
                                            }
                                        }
                                    }
                                }
                            }
                        }
                        if v4 {
                            fmt.Println(fi97)
                        } else {
                            var v134 int32 = genFn1(-1, (^(v2)), 1)
                            _ = v134
                        }
                        v135 := GenS18{F0: 42}
                        _ = v135
                    } else {
                        if (false && v4) {
                            v136 := GenS19{F0: 632, F1: -1, F2: "hello"}
                            _ = v136
                        }
                        switch "foo bar" {
                        case "":
                            v99 = v99
                        case "hello":
                            var v137 bool = false
                            if false {
                                var v138 int32 = 931
                                var v139 int64 = 281
                                _ = v138
                                _ = v139
                            }
                            fmt.Println(fi97, fi97)
                            _ = v137
                        case "foo bar":
                            switch v99 {
                            case "":
                                var v140 int = 42
                                _ = v140
                            default:
                                v141 := GenS20{F0: ((("a" + "\n\t") + ((v99 + ("a" + (v99 + (v99 + v99)))) + "")) + ("foo bar" + (((v99 + "\n\t") + v99) + "hello"))), F1: v4, F2: 2147483647}
                                _ = v141
                            }
                        default:
                            var v142 int = fi97
                            v143 := v1[591]
                            _ = v142
                            _ = v143
                        }
                        var v144 bool = (42 >= -376)
                        var v145 uint = 65535
                        _ = v144
                        _ = v145
                    }
                default:
                    var v146 uint32 = 838
                    _ = v146
                }
                _ = fi97
                _ = v98
                _ = v99
            }
        case "\n\t":
            fmt.Println(v4, v4)
        }
        for fi147 := 0; fi147 < 4; fi147++ {
            var v148 uint32 = 2
            v149 := GenS20{F0: "a", F1: (false || false), F2: (v148 ^ v148)}
            if len(v66) >= 2 {
                v66 = v66[1:len(v66)]
            }
            _ = fi147
            _ = v148
            _ = v149
        }
        _ = fi92
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v66:", len(v66), v66)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v66
    _ = v90
}
