package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS2 struct { F0 int; F1 uint32; F2 int64 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 uint64; F1 uint32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint32; F1 bool }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 bool; F1 int8; F2 bool }
type GenS10 struct { F0 int16; F1 int32; F2 int }
type GenS11 struct { F0 int8; F1 string; F2 uint64 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 int; F1 int16; F2 int32 }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS17 struct { F0 int16 }
type GenS18 struct { F0 int16 }
func genFn0(p0 int) int16 {
    v6 := GenS1{F0: 2, F1: -2147483648, F2: 94}
    if (false && (!(1 >= -65))) {
        p0 = p0
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 uint16 = 31
            for fi9 := 0; fi9 < 19; fi9++ {
                var v10 uint = 872
                _ = fi9
                _ = v10
            }
            switch "a" {
            case "foo bar":
                v11 := GenS2{F0: (fi7 + p0), F1: 2, F2: 9223372036854775807}
                _ = v11
            case "a":
                v12 := []int{(p0 << ((304) & 63)), p0, (-(fi7))}
                _ = v12
            case "":
                var v13 int = (p0 - fi7)
                _ = v13
            default:
                var v14 int = (^(p0))
                _ = v14
            }
            _ = fi7
            _ = v8
        }
        v15 := []int{p0}
        _ = v15
    }
    switch (p0 / (((^(p0))) | 1)) {
    case -984:
        p0 += (p0 & (p0 << ((42) & 63)))
    }
    _ = p0
    _ = v6
    return 453
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{0, 1, 1}
    v3 := make(map[string]uint32)
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]bool)
    for fi16 := 0; fi16 < 16; fi16++ {
        if (((("hello" + (("" + "foo bar") + "")) >= "\n\t") || false) && false) {
            fmt.Println(fi16)
            switch "\n\t" {
            case "\n\t":
                var v17 uint8 = 255
                _ = v17
            case "hello":
                switch (fi16 >> ((307) & 63)) {
                case 2147483647:
                    var v18 bool = true
                    _ = v18
                case 9223372036854775807:
                    fmt.Println(fi16)
                case 694:
                    v19 := v3["a"]
                    _ = v19
                }
            }
        }
        _ = fi16
    }
    var v20 int64 = 162
    v20 = v20
    fmt.Println(v20)
    v2 = append(v2, 2147483648)
    var v21 uint16 = 42
    _ = len(v2)
    switch 954 {
    case 32768:
        for fi22 := 0; fi22 < 3; fi22++ {
            var v23 uint32 = 4294967295
            _ = fi22
            _ = v23
        }
    }
    var v24 uint64 = 300
    var v25 int32 = -2147483648
    var v26 float32 = 0.0
    if true {
        v5[42] = (!((0.1 / 1.0) > 1.5))
        switch 351 {
        case 32768:
            var v27 int8 = 1
            _ = v27
        case 1:
            v21 -= (v21 - (v21 ^ v21))
        case -3:
            var v28 bool = true
            _ = v28
        }
    } else {
        var v29 uint32 = 1
        v30 := GenS2{F0: 128, F1: v29, F2: (v20 & -942)}
        _ = v29
        _ = v30
    }
    var v31 bool = false
    if v31 {
        v24 += v24
    } else {
        v26 = v26
        switch int(v21) {
        case 9223372036854775807:
            v31 = (v31 || v31)
        case -416:
            _ = len(v4)
        case 2147483647:
            v2 = append(v2, 2147483648)
        default:
            v32 := GenS3{F0: -809, F1: 103}
            _ = v32
        }
    }
    switch 904 {
    case -9223372036854775808:
        if (2 == 388) {
            fmt.Println(v31, v31)
            fmt.Println(v20)
        } else {
            v33 := GenS2{F0: 42, F1: 2147483648, F2: -9223372036854775808}
            for fi34 := 0; fi34 < 12; fi34++ {
                var v35 int = 42
                switch 2147483647 {
                case 42:
                    for fi36 := 0; fi36 < 20; fi36++ {
                        v26 += (((v26 + ((0.1 - (float32(0.0) * ((((v26 + 0.0) + -2.25) + (v26 / 1e10)) / ((((0.0 / v26) / 3.14159) - 0.1) / v26)))) - (v26 * v26))) - 3.14159) / 0.1)
                        switch ("a" + ("foo bar" + "hello")) {
                        case "a":
                            v24 = (v24 & v24)
                        case "hello":
                            var v37 float64 = float64(1e10)
                            switch fi36 {
                            case -2147483648:
                                if v31 {
                                    switch "" {
                                    case "\n\t":
                                        fmt.Println(v35, v25)
                                    case "hello":
                                        var v38 uint = 65535
                                        v37 = float64(v26)
                                        if false {
                                            v39 := GenS4{F0: "foo bar", F1: false}
                                            _ = v39
                                        } else {
                                            v2 = append(v2, 155)
                                            var v40 string = ("a" + "")
                                            v41 := GenS1{F0: 2, F1: (-(v25)), F2: v21}
                                            if v31 {
                                                switch v40 {
                                                case "hello":
                                                    var v42 int16 = 1000
                                                    _ = v42
                                                case "foo bar":
                                                    if true {
                                                        if true {
                                                            v43 := GenS1{F0: 255, F1: (v25 | (v25 % ((v25) | 1))), F2: v21}
                                                            v44 := GenS2{F0: (fi36 | (-(v35))), F1: 446, F2: v20}
                                                            _ = v43
                                                            _ = v44
                                                        } else {
                                                            switch (-(fi34)) {
                                                            case 1:
                                                                var v45 uint16 = v21
                                                                _ = v45
                                                            case 128:
                                                                v46 := GenS3{F0: -1, F1: fi34}
                                                                _ = v46
                                                            }
                                                        }
                                                    }
                                                case "\n\t":
                                                    var v47 uint32 = 2
                                                    _ = v47
                                                }
                                                var v48 int16 = -692
                                                fmt.Println(v21)
                                                switch v35 {
                                                case 127:
                                                    var v49 int64 = int64(v48)
                                                    v50 := v5[575]
                                                    fmt.Println(v37, v26, v24)
                                                    if (!v31) {
                                                        switch "a" {
                                                        case "\n\t":
                                                            fmt.Println(v37, v49, v20)
                                                        case "hello":
                                                            v51 := GenS4{F0: ("" + v40), F1: (genFn0(v35) <= -32768)}
                                                            _ = v51
                                                        default:
                                                            v52 := GenS5{F0: (^(v24)), F1: 4294967295}
                                                            _ = v52
                                                        }
                                                        switch (fi34 >> (((v38 >> (((v38 - 18446744073709551615)) & 63))) & 63)) {
                                                        case 2147483647:
                                                            v25 = -2147483648
                                                        }
                                                    } else {
                                                        fmt.Println(v48, v49, v38)
                                                        v26 += -2.25
                                                    }
                                                    switch -176 {
                                                    case 32768:
                                                        fmt.Println(v25)
                                                    case 42:
                                                        var v53 uint8 = 42
                                                        _ = v53
                                                    }
                                                    _ = v49
                                                    _ = v50
                                                default:
                                                    v40 = v40
                                                }
                                                v21 -= (^(v21))
                                                _ = v48
                                            }
                                            fmt.Println(v26)
                                            switch 335 {
                                            case 0:
                                                v35 -= int(v38)
                                                fmt.Println(v21, fi36)
                                                var v54 int = 128
                                                _ = v54
                                            case -2147483648:
                                                var v55 string = (v40 + (v40 + v40))
                                                _ = v55
                                            default:
                                                var v56 uint16 = 65535
                                                _ = v56
                                            }
                                            if v31 {
                                                if v31 {
                                                    var v57 float32 = (0.0 + float32(v37))
                                                    var v58 uint16 = v21
                                                    if ((^(v58)) != (v21 / (((v21 % (((v58 << (((v38 / (((v38 | (v38 / ((v38) | 1)))) | 1))) & 63))) | 1))) | 1))) {
                                                        fmt.Println(v24, v20, v57)
                                                        if ((!v31) && v31) {
                                                            if (!false) {
                                                                fmt.Println(v37, v25)
                                                                _ = len(v2)
                                                            } else {
                                                                v59 := GenS6{F0: v40}
                                                                _ = v59
                                                            }
                                                            v60 := GenS4{F0: v40, F1: v31}
                                                            _ = v60
                                                        }
                                                        fmt.Println(v31, v26, v40)
                                                    } else {
                                                        switch v40 {
                                                        case "\n\t":
                                                            var v61 int16 = 505
                                                            _ = v61
                                                        default:
                                                            _ = len(v2)
                                                        }
                                                        switch "hello" {
                                                        case "a":
                                                            v2 = append(v2, 42)
                                                        case "hello":
                                                            v62 := GenS1{F0: 255, F1: (^(v25)), F2: (v21 | (v58 ^ v58))}
                                                            var v63 int16 = 32767
                                                            fmt.Println(fi36, v63, v63)
                                                            _ = v62
                                                            _ = v63
                                                        case "foo bar":
                                                            var v64 int8 = 1
                                                            _ = v64
                                                        default:
                                                            v21 += 1
                                                        }
                                                    }
                                                    v65 := GenS1{F0: 127, F1: v25, F2: (v21 >> ((uint(v24)) & 63))}
                                                    fmt.Println(v21, v25)
                                                    _ = v57
                                                    _ = v58
                                                    _ = v65
                                                }
                                                var v66 int8 = 86
                                                if v31 {
                                                    var v67 int32 = v25
                                                    v20 -= v20
                                                    switch v40 {
                                                    case "hello":
                                                        v38 = v38
                                                    }
                                                    _ = v67
                                                } else {
                                                    v4[fi34] = 4294967295
                                                    v68 := GenS6{F0: (("foo bar" + v40) + v40)}
                                                    v20 = v20
                                                    _ = v68
                                                }
                                                _ = v66
                                            } else {
                                                var v69 int32 = 2147483647
                                                v3[v40] = 65535
                                                delete(v1, "foo bar")
                                                v2 = append(v2, 2147483648)
                                                _ = v69
                                            }
                                            var v70 int = 32767
                                            delete(v1, v40)
                                            v20 = -9223372036854775808
                                            v71 := v1[v40]
                                            v2 = append(v2, 157)
                                            _ = v40
                                            _ = v41
                                            _ = v70
                                            _ = v71
                                        }
                                        _ = v38
                                    case "foo bar":
                                        var v72 uint = 2
                                        _ = v72
                                    default:
                                        v35 = 32768
                                    }
                                    v21 = 712
                                    if (true || ((v25 + (v25 & 1)) >= (v25 % ((v25) | 1)))) {
                                        switch 948 {
                                        case 0:
                                            v73 := GenS7{F0: 232, F1: v31}
                                            _ = v73
                                        case 9223372036854775807:
                                            var v74 uint64 = v24
                                            _ = v74
                                        case -1:
                                            v4[32767] = 2147483647
                                        }
                                        var v75 int64 = 0
                                        _ = v75
                                    } else {
                                        v35 -= (fi34 | (^(fi36)))
                                        var v76 int8 = -58
                                        _ = v76
                                    }
                                }
                            default:
                                if v31 {
                                    v77 := v3["a"]
                                    delete(v4, 32768)
                                    v77 = v77
                                    v78 := v5[(v77 - (^(v77)))]
                                    _ = v77
                                    _ = v78
                                }
                            }
                            if ((false && v31) || v31) {
                                fmt.Println(fi34, v35)
                                var v79 uint32 = uint32(v21)
                                _ = v79
                            } else {
                                v80 := GenS8{F0: 123}
                                switch "" {
                                case "":
                                    var v81 int = 0
                                    _ = v81
                                case "\n\t":
                                    v37 -= (v37 + v37)
                                }
                                var v82 float64 = float64((v26 * float32(((3.14159 - float64(v26)) + -2.25))))
                                _ = v80
                                _ = v82
                            }
                            _ = v37
                        case "foo bar":
                            switch (v35 / (((fi36 + (fi36 | (fi34 % ((v35) | 1))))) | 1)) {
                            case -9223372036854775808:
                                var v83 float32 = -2.25
                                _ = v83
                            case 2147483647:
                                var v84 int32 = v25
                                _ = v84
                            default:
                                fmt.Println(fi34, v26)
                            }
                        }
                        var v85 int = (fi36 >> ((18446744073709551615) & 63))
                        switch "" {
                        case "\n\t":
                            v86 := GenS7{F0: 65535, F1: false}
                            var v87 uint64 = v24
                            if ("hello" >= "a") {
                                var v88 int16 = 84
                                v89 := v5[0]
                                if (!(!(1 <= 0))) {
                                    v90 := GenS8{F0: 255}
                                    if v31 {
                                        var v91 bool = ((-(fi36)) <= v85)
                                        _ = v91
                                    }
                                    fmt.Println(v24)
                                    _ = v90
                                } else {
                                    delete(v1, ("hello" + "hello"))
                                }
                                _ = v88
                                _ = v89
                            } else {
                                v92 := GenS9{F0: v31, F1: -1, F2: (151 <= 127)}
                                v93 := GenS10{F0: 32767, F1: (v25 - (v25 | 27)), F2: (^(v85))}
                                var v94 int16 = 42
                                v35 = v35
                                _ = v92
                                _ = v93
                                _ = v94
                            }
                            _ = v86
                            _ = v87
                        case "a":
                            v25 = v25
                        case "hello":
                            v95 := GenS7{F0: 4294967295, F1: (2147483648 < 4294967295)}
                            _ = v95
                        default:
                            var v96 uint = 0
                            _ = v96
                        }
                        _ = fi36
                        _ = v85
                    }
                case 0:
                    var v97 int64 = (v20 + v20)
                    _ = v97
                default:
                    var v98 uint16 = v21
                    _ = v98
                }
                v21 = (v21 + v21)
                _ = fi34
                _ = v35
            }
            for fi99 := 0; fi99 < 5; fi99++ {
                v21 = 1
                v2 = append(v2, 42)
                for fi100 := 0; fi100 < 3; fi100++ {
                    v26 = v26
                    var v101 string = "a"
                    _ = fi100
                    _ = v101
                }
                _ = fi99
            }
            _ = v33
        }
    case 949:
        v20 -= (v20 ^ v20)
    case 1:
        for fi102 := 0; fi102 < 15; fi102++ {
            var v103 string = ("" + "hello")
            v26 += 1.5
            switch (fi102 - 4294967295) {
            case -9223372036854775808:
                v24 += (^(v24))
            case 31:
                v104 := GenS3{F0: 0, F1: int(v24)}
                _ = v104
            default:
                if ((1 == 2147483648) || v31) {
                    v105 := GenS11{F0: 127, F1: "foo bar", F2: 42}
                    v103 = "\n\t"
                    _ = v105
                }
            }
            _ = fi102
            _ = v103
        }
    default:
        switch -857 {
        case 127:
            v106 := GenS6{F0: "foo bar"}
            _ = v106
        case 1:
            for fi107 := 0; fi107 < 8; fi107++ {
                v25 += v25
                if ((v26 + v26) != ((float32(0.0) * v26) + ((v26 - (0.0 / 1.0)) / v26))) {
                    var v108 uint16 = v21
                    _ = v108
                }
                for fi109 := 0; fi109 < 5; fi109++ {
                    v110 := GenS4{F0: "foo bar", F1: true}
                    v24 -= 92
                    fmt.Println(v26)
                    _ = fi109
                    _ = v110
                }
                _ = fi107
            }
        case -2147483648:
            v25 = (^(v25))
        }
    }
    delete(v1, "hello")
    switch 128 {
    case 127:
        for fi111 := 0; fi111 < 11; fi111++ {
            v31 = v31
            _ = fi111
        }
    case -428:
        var v112 uint8 = 195
        _ = v112
    default:
        for fi113 := 0; fi113 < 19; fi113++ {
            v26 += 1.5
            if v31 {
                if (65535 <= 757) {
                    var v114 uint64 = v24
                    if ("hello" == "") {
                        v115 := GenS12{F0: 127}
                        v2 = append(v2, 1)
                        if v31 {
                            v116 := GenS13{F0: (fi113 + (fi113 % ((9223372036854775807) | 1)))}
                            _ = v116
                        } else {
                            switch fi113 {
                            case 4294967295:
                                var v117 float32 = (v26 / 1.0)
                                _ = v117
                            case -9223372036854775808:
                                for fi118 := 0; fi118 < 3; fi118++ {
                                    var v119 bool = v31
                                    var v120 float32 = float32((-2.25 / (3.14159 + 0.0)))
                                    v119 = true
                                    _ = fi118
                                    _ = v119
                                    _ = v120
                                }
                            case -2147483648:
                                v31 = true
                            }
                            v24 += v24
                        }
                        _ = v115
                    } else {
                        switch (fi113 & (fi113 - fi113)) {
                        case 2147483647:
                            v21 = v21
                        case -1:
                            v4[fi113] = uint32(v114)
                        case 300:
                            if v31 {
                                v121 := GenS5{F0: 4294967295, F1: 648}
                                if v31 {
                                    if len(v2) > 0 {
                                        v2[0] = 42
                                    }
                                    v5[1] = true
                                }
                                v122 := GenS2{F0: fi113, F1: 42, F2: v20}
                                _ = v121
                                _ = v122
                            }
                        }
                        var v123 uint = 18446744073709551615
                        var v124 int16 = 78
                        _ = v123
                        _ = v124
                    }
                    var v125 int16 = 1000
                    _ = v114
                    _ = v125
                } else {
                    if len(v2) > 0 {
                        v2[0] = 864
                    }
                    switch 2147483647 {
                    case 127:
                        switch 127 {
                        case -536:
                            for fi126 := 0; fi126 < 18; fi126++ {
                                if false {
                                    v127 := GenS14{F0: (fi113 * (fi126 ^ 9223372036854775807)), F1: -219, F2: (v25 ^ 0)}
                                    var v128 uint64 = v24
                                    v129 := GenS14{F0: 32768, F1: -1, F2: v25}
                                    switch ("" + "") {
                                    case "hello":
                                        var v130 int32 = 527
                                        _ = v130
                                    case "a":
                                        v21 = v21
                                    case "\n\t":
                                        fmt.Println(v20, fi113)
                                        v24 += v24
                                        var v131 float64 = 1.5
                                        _ = v131
                                    }
                                    _ = v127
                                    _ = v128
                                    _ = v129
                                }
                                _ = fi126
                            }
                        case 2147483647:
                            if v31 {
                                fmt.Println(v21)
                                var v132 uint32 = 817
                                v25 += v25
                                _ = v132
                            } else {
                                v26 = v26
                                if len(v2) > 0 {
                                    v2[0] = 2
                                }
                            }
                        default:
                            v24 = v24
                        }
                    case -238:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    for fi133 := 0; fi133 < 5; fi133++ {
                        var v134 float32 = v26
                        v135 := GenS15{F0: v25}
                        _ = fi133
                        _ = v134
                        _ = v135
                    }
                }
                for fi136 := 0; fi136 < 18; fi136++ {
                    switch "\n\t" {
                    case "\n\t":
                        var v137 int8 = 23
                        _ = v137
                    case "a":
                        v26 = v26
                    case "hello":
                        v21 += v21
                    default:
                        v138 := GenS16{F0: 1, F1: 2, F2: 65535}
                        _ = v138
                    }
                    if v31 {
                        v139 := GenS17{F0: -172}
                        var v140 string = ((("a" + "foo bar") + "foo bar") + "\n\t")
                        _ = v139
                        _ = v140
                    }
                    v141 := GenS18{F0: -32768}
                    _ = fi136
                    _ = v141
                }
                v142 := GenS3{F0: -32768, F1: (fi113 | fi113)}
                _ = v142
            }
            var v143 int32 = 42
            _ = fi113
            _ = v143
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
