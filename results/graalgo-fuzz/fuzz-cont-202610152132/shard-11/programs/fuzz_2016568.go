package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 int; F1 int; F2 int8 }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 int8; F1 int8; F2 uint }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 string; F1 bool; F2 uint8 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 bool; F1 int8; F2 uint64 }
type GenS10 struct { F0 uint64; F1 int8; F2 string }
type GenS11 struct { F0 int; F1 uint8 }
type GenS12 struct { F0 int64; F1 uint; F2 uint64 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS18 struct { F0 int32; F1 string; F2 uint64 }
type GenS19 struct { F0 uint64 }
type GenS20 struct { F0 uint; F1 uint16 }
type GenS21 struct { F0 uint32; F1 int64 }
func genFn0(p0 int16, p1 bool) bool {
    if (2 >= 129) {
        p1 = (!p1)
        v7 := []int{-388, 128}
        _ = v7
    } else {
        p0 = p0
    }
    p1 = (3.14159 != -2.25)
    fmt.Println(p1, p0)
    if (0 <= 60) {
        fmt.Println(p0, p1)
        for fi8 := 0; fi8 < 16; fi8++ {
            v9 := []int{(fi8 + (fi8 >> ((553) & 63))), (-(fi8)), (fi8 / ((-340) | 1))}
            p0 = p0
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 19; fi10++ {
            if (!(!(2 < 255))) {
                var v11 int16 = p0
                v11 = 681
                switch fi10 {
                case -2147483648:
                    v11 += v11
                case 0:
                    p0 = p0
                case -112:
                    var v12 int8 = 127
                    _ = v12
                default:
                    v13 := []int{(fi10 ^ -9223372036854775808), fi10, (fi10 / ((fi10) | 1))}
                    _ = v13
                }
                _ = v11
            }
            if p1 {
                for fi14 := 0; fi14 < 13; fi14++ {
                    p0 += p0
                    switch "a" {
                    case "":
                        if (-368 > 1) {
                            v15 := []int{(fi10 + (fi14 / ((32768) | 1))), (fi14 * (-(fi10))), fi10, fi14}
                            v15 = append(v15, fi10)
                            var v16 int = (fi10 | fi14)
                            var v17 string = ("hello" + "hello")
                            _ = v15
                            _ = v16
                            _ = v17
                        } else {
                            if p1 {
                                p0 = 845
                            } else {
                                if (1 < 113) {
                                    p0 -= p0
                                    var v18 int = (fi14 + fi10)
                                    var v19 int8 = 111
                                    _ = v18
                                    _ = v19
                                } else {
                                    v20 := []int{(fi14 ^ fi14)}
                                    v20 = append(v20, (^(fi14)))
                                    fmt.Println(fi10)
                                    _ = v20
                                }
                                switch ("" + ("" + "")) {
                                case "foo bar":
                                    if p1 {
                                        v21 := GenS1{F0: (1.5 < 0.1), F1: "foo bar"}
                                        _ = v21
                                    } else {
                                        p1 = ((-2.25 / 1.0) == ((float32(1.5) + 0.0) + 0.1))
                                        var v22 uint8 = 42
                                        _ = v22
                                    }
                                case "hello":
                                    var v23 string = ((((((("\n\t" + "\n\t") + (("\n\t" + ("a" + "a")) + (("hello" + "hello") + ("a" + "foo bar")))) + "\n\t") + "hello") + "\n\t") + ("foo bar" + "hello")) + "hello")
                                    _ = v23
                                case "a":
                                    if ((false || p1) && p1) {
                                        p0 = (p0 % ((p0) | 1))
                                        fmt.Println(p0, p0, fi10)
                                    } else {
                                        if p1 {
                                            p1 = p1
                                            v24 := GenS1{F0: p1, F1: "foo bar"}
                                            _ = v24
                                        }
                                        v25 := GenS1{F0: p1, F1: "foo bar"}
                                        v26 := []int{fi14, -1, fi10}
                                        _ = v25
                                        _ = v26
                                    }
                                }
                                if p1 {
                                    fmt.Println(p1)
                                    switch fi14 {
                                    case -2147483648:
                                        switch "\n\t" {
                                        case "a":
                                            switch -37 {
                                            case -970:
                                                p0 = 1000
                                            case 2147483647:
                                                var v27 uint16 = 587
                                                _ = v27
                                            }
                                            p0 = (p0 / ((32767) | 1))
                                        case "hello":
                                            var v28 int64 = 157
                                            _ = v28
                                        }
                                        if (((3.14159 / 0.1) * 3.14159) != -2.25) {
                                            var v29 string = "hello"
                                            _ = v29
                                        }
                                        var v30 bool = (!(-111 > 1))
                                        if p1 {
                                            v31 := []int{fi14, fi14}
                                            _ = v31
                                        }
                                        p0 -= 42
                                        _ = v30
                                    case -1:
                                        switch "" {
                                        case "foo bar":
                                            p0 = 0
                                        case "hello":
                                            p1 = p1
                                        case "a":
                                            v32 := []int{128, (fi14 + fi14), -868}
                                            _ = v32
                                        }
                                    }
                                    v33 := []int{9223372036854775807}
                                    _ = v33
                                } else {
                                    fmt.Println(p1, fi10)
                                    v34 := []int{}
                                    _ = v34
                                }
                            }
                            p1 = ((((1e10 - (float32((float64(((((1e10 * 0.0) * 1e10) + 3.14159) + float32(1.5))) * ((3.14159 + 1e10) / (3.14159 + 1.5)))) / float32((1.5 - float64(-2.25))))) / 0.1) + 0.1) <= 3.14159)
                        }
                    case "foo bar":
                        var v35 uint8 = 144
                        v35 = (v35 - 127)
                        _ = v35
                    case "a":
                        fmt.Println(fi10, fi14)
                    }
                    v36 := []int{(fi14 - (fi14 % ((-1) | 1))), fi14, (fi14 - fi10)}
                    var v37 float32 = (1.5 * 0.1)
                    fmt.Println(p1, fi10, p0)
                    _ = fi14
                    _ = v36
                    _ = v37
                }
                v38 := GenS1{F0: (p0 > p0), F1: "foo bar"}
                p1 = ((-(p0)) < p0)
                _ = v38
            }
            _ = fi10
        }
    } else {
        for fi39 := 0; fi39 < 2; fi39++ {
            p0 -= (p0 % (((p0 ^ 1)) | 1))
            if p1 {
                var v40 int = (fi39 & (fi39 / (((fi39 - (fi39 % ((fi39) | 1)))) | 1)))
                v40 += -2147483648
                switch ("a" + "foo bar") {
                case "":
                    for fi41 := 0; fi41 < 19; fi41++ {
                        switch fi41 {
                        case -694:
                            if (1 < 42) {
                                var v42 int8 = -108
                                switch (fi41 - v40) {
                                case -1:
                                    v43 := GenS1{F0: (p1 || ((v42 <= int8(p0)) && p1)), F1: "a"}
                                    _ = v43
                                }
                                _ = v42
                            } else {
                                var v44 int32 = -1
                                v44 += -1
                                _ = v44
                            }
                        case 0:
                            var v45 float64 = 1e10
                            _ = v45
                        default:
                            fmt.Println(p1)
                        }
                        switch (fi39 >> ((255) & 63)) {
                        case 127:
                            switch v40 {
                            case 4294967295:
                                switch "" {
                                case "":
                                    v46 := []int{}
                                    _ = v46
                                }
                            case -1:
                                if false {
                                    var v47 uint8 = 138
                                    fmt.Println(v40, fi39, fi41)
                                    switch fi41 {
                                    case 32768:
                                        if false {
                                            v40 = v40
                                            v48 := []int{}
                                            _ = v48
                                        } else {
                                            if true {
                                                v49 := []int{}
                                                var v50 uint64 = 1
                                                var v51 string = ("\n\t" + (("a" + "") + "foo bar"))
                                                _ = v49
                                                _ = v50
                                                _ = v51
                                            }
                                            var v52 int8 = 1
                                            _ = v52
                                        }
                                    default:
                                        if (42 != 42) {
                                            switch fi41 {
                                            case 32767:
                                                var v53 bool = (p0 == (^(p0)))
                                                _ = v53
                                            case 42:
                                                v54 := []int{v40, fi41, 655}
                                                _ = v54
                                            case 9223372036854775807:
                                                var v55 uint16 = 32767
                                                _ = v55
                                            default:
                                                if (p1 && (((0.1 - 1e10) * 1.5) == -2.25)) {
                                                    switch v40 {
                                                    case 4294967295:
                                                        v56 := []int{}
                                                        _ = v56
                                                    }
                                                    fmt.Println(p0, v47, p1)
                                                }
                                            }
                                            switch ((("a" + "foo bar") + "hello") + "") {
                                            case "hello":
                                                fmt.Println(p0, fi39)
                                            case "\n\t":
                                                p0 += p0
                                            case "foo bar":
                                                var v57 int = fi39
                                                _ = v57
                                            }
                                            fmt.Println(v40, p1, v40)
                                        } else {
                                            var v58 int8 = int8(p0)
                                            if (p1 || (2147483647 == 2)) {
                                                if true {
                                                    v59 := []int{(v40 | fi41), 32768, (fi39 & (v40 % ((9223372036854775807) | 1)))}
                                                    _ = v59
                                                } else {
                                                    p0 += p0
                                                    v60 := GenS1{F0: (("\n\t" < ("\n\t" + "")) || p1), F1: (("foo bar" + "hello") + "")}
                                                    var v61 uint32 = 65535
                                                    _ = v60
                                                    _ = v61
                                                }
                                                v62 := GenS2{F0: (v58 ^ (v58 * (v58 ^ (-(v58))))), F1: -776}
                                                fmt.Println(fi41, fi39, fi39)
                                                _ = v62
                                            }
                                            v63 := GenS3{F0: (fi41 - fi39), F1: fi39, F2: 0}
                                            if (!true) {
                                                v64 := GenS1{F0: true, F1: ("hello" + "a")}
                                                var v65 uint = 18446744073709551615
                                                var v66 bool = p1
                                                var v67 uint8 = 125
                                                if true {
                                                    fmt.Println(p0)
                                                    p0 -= 0
                                                    var v68 int64 = -196
                                                    _ = v68
                                                } else {
                                                    if p1 {
                                                        v67 -= v47
                                                        var v69 int16 = p0
                                                        _ = v69
                                                    }
                                                    v70 := GenS3{F0: int(v58), F1: v40, F2: (-(v58))}
                                                    v71 := GenS4{F0: -299, F1: fi41}
                                                    if v66 {
                                                        if ((4294967295 <= 221) && (true || (!false))) {
                                                            if p1 {
                                                                fmt.Println(fi41, p1)
                                                            } else {
                                                                v66 = v66
                                                                v58 += 42
                                                            }
                                                            var v72 uint16 = 65535
                                                            _ = v72
                                                        }
                                                        fmt.Println(v65, v47, v40)
                                                    }
                                                    v73 := []int{9223372036854775807, 32767}
                                                    _ = v70
                                                    _ = v71
                                                    _ = v73
                                                }
                                                var v74 float32 = 3.14159
                                                var v75 string = ""
                                                v76 := []int{}
                                                if false {
                                                    v77 := GenS5{F0: 95, F1: v58, F2: v65}
                                                    v78 := GenS1{F0: (!p1), F1: v75}
                                                    v67 = v67
                                                    _ = v77
                                                    _ = v78
                                                }
                                                _ = v64
                                                _ = v65
                                                _ = v66
                                                _ = v67
                                                _ = v74
                                                _ = v75
                                                _ = v76
                                            } else {
                                                if p1 {
                                                    fmt.Println(v47, v40)
                                                    v79 := GenS3{F0: (fi41 + fi41), F1: fi41, F2: v58}
                                                    _ = v79
                                                } else {
                                                    if p1 {
                                                        v80 := GenS6{F0: 4294967295}
                                                        v40 -= fi41
                                                        v47 = (^(v47))
                                                        v81 := GenS1{F0: p1, F1: ("" + "foo bar")}
                                                        if (int16(v47) < (p0 & p0)) {
                                                            v82 := []int{fi39}
                                                            v83 := GenS7{F0: ((("foo bar" + "a") + "foo bar") + ""), F1: (p0 <= -1), F2: v47}
                                                            v84 := GenS8{F0: (^(v47))}
                                                            v85 := GenS6{F0: 1}
                                                            _ = v82
                                                            _ = v83
                                                            _ = v84
                                                            _ = v85
                                                        } else {
                                                            p1 = (0 >= 461)
                                                        }
                                                        _ = v80
                                                        _ = v81
                                                    } else {
                                                        v86 := GenS5{F0: (^(v58)), F1: (-(v58)), F2: 18446744073709551615}
                                                        var v87 string = ("a" + "hello")
                                                        var v88 bool = (p0 == p0)
                                                        _ = v86
                                                        _ = v87
                                                        _ = v88
                                                    }
                                                }
                                                var v89 string = "\n\t"
                                                _ = v89
                                            }
                                            _ = v58
                                            _ = v63
                                        }
                                    }
                                    _ = v47
                                } else {
                                    if (((!((100 != -1) || (((1e10 - 3.14159) == (-2.25 + (0.0 / 3.14159))) || p1))) && p1) && (p1 && (!p1))) {
                                        var v90 float64 = -2.25
                                        var v91 float32 = 1e10
                                        var v92 int = (-(fi39))
                                        if true {
                                            v93 := GenS4{F0: -1, F1: v92}
                                            if true {
                                                if p1 {
                                                    var v94 string = "hello"
                                                    fmt.Println(p1)
                                                    _ = v94
                                                }
                                                v95 := []int{(fi39 * 1), -700, v92}
                                                if p1 {
                                                    var v96 uint32 = 0
                                                    var v97 uint32 = (v96 % ((1) | 1))
                                                    fmt.Println(v90)
                                                    _ = v96
                                                    _ = v97
                                                } else {
                                                    v98 := GenS6{F0: 103}
                                                    _ = v98
                                                }
                                                _ = v95
                                            } else {
                                                if true {
                                                    v92 -= fi41
                                                    v92 -= 1
                                                    v99 := []int{127}
                                                    _ = v99
                                                }
                                                fmt.Println(p1, fi41)
                                            }
                                            v100 := GenS1{F0: p1, F1: ("a" + ("hello" + "\n\t"))}
                                            _ = v100
                                            _ = v93
                                        } else {
                                            v91 = -2.25
                                            fmt.Println(v90, v91)
                                            fmt.Println(v92, p1)
                                        }
                                        v101 := GenS5{F0: 127, F1: 1, F2: 65535}
                                        v90 -= (v90 * 1e10)
                                        var v102 int64 = -316
                                        p0 -= (^(p0))
                                        v103 := []int{410}
                                        _ = v101
                                        _ = v102
                                        _ = v103
                                        _ = v90
                                        _ = v91
                                        _ = v92
                                    }
                                    v104 := GenS9{F0: (int8(p0) > 81), F1: 31, F2: 837}
                                    v105 := GenS6{F0: 2147483648}
                                    _ = v104
                                    _ = v105
                                }
                            case 127:
                                if true {
                                    if ((p0 >> ((18446744073709551615) & 63)) > (^(p0))) {
                                        p0 -= p0
                                        v40 = -1
                                    }
                                    v106 := GenS10{F0: 0, F1: 42, F2: "hello"}
                                    _ = v106
                                }
                            default:
                                if p1 {
                                    if false {
                                        if p1 {
                                            p0 -= p0
                                        }
                                        if p1 {
                                            var v107 int64 = 42
                                            if p1 {
                                                var v108 int64 = 1
                                                _ = v108
                                            }
                                            _ = v107
                                        }
                                    } else {
                                        if (false && p1) {
                                            v109 := GenS11{F0: fi41, F1: 128}
                                            _ = v109
                                        }
                                        var v110 int16 = (p0 ^ 32767)
                                        v111 := []int{(fi39 ^ fi39), -2147483648, fi39, fi39}
                                        p0 = 42
                                        _ = v110
                                        _ = v111
                                    }
                                    v112 := GenS1{F0: p1, F1: "hello"}
                                    _ = v112
                                }
                            }
                        }
                        if (3.14159 <= 0.0) {
                            v40 -= (fi41 & v40)
                            v113 := GenS12{F0: -133, F1: 44, F2: 2}
                            _ = v113
                        }
                        _ = fi41
                    }
                case "a":
                    if p1 {
                        v114 := GenS13{F0: 2}
                        var v115 float32 = 1.5
                        _ = v114
                        _ = v115
                    } else {
                        switch ("foo bar" + ("hello" + "a")) {
                        case "a":
                            v116 := GenS6{F0: 734}
                            _ = v116
                        default:
                            var v117 int32 = -2147483648
                            _ = v117
                        }
                        var v118 int16 = p0
                        _ = v118
                    }
                case "hello":
                    fmt.Println(v40, fi39)
                }
                _ = v40
            }
            v119 := GenS9{F0: true, F1: -109, F2: 4294967295}
            _ = fi39
            _ = v119
        }
    }
    _ = p0
    _ = p1
    return p1
}
func genFn1(p0 string, p1 int8, p2 uint) int32 {
    if (((!false) && false) && (1 != 801)) {
        var v120 uint16 = 0
        switch (p0 + "") {
        case "":
            if true {
                for fi121 := 0; fi121 < 17; fi121++ {
                    v122 := GenS10{F0: 279, F1: (^(p1)), F2: p0}
                    _ = fi121
                    _ = v122
                }
            } else {
                var v123 bool = false
                if true {
                    fmt.Println(v120)
                } else {
                    p1 += p1
                    v123 = false
                    v124 := []int{4294967295, 127}
                    _ = v124
                }
                var v125 uint8 = 255
                _ = v123
                _ = v125
            }
        case "a":
            for fi126 := 0; fi126 < 6; fi126++ {
                for fi127 := 0; fi127 < 7; fi127++ {
                    if false {
                        var v128 string = ""
                        v128 = (p0 + ("a" + ((("a" + "\n\t") + p0) + "")))
                        _ = v128
                    } else {
                        v129 := GenS8{F0: 128}
                        switch "a" {
                        case "\n\t":
                            if ((true || true) && false) {
                                if (p0 <= p0) {
                                    var v130 int16 = -56
                                    _ = v130
                                } else {
                                    switch "" {
                                    case "foo bar":
                                        v131 := GenS9{F0: ((-(p1)) <= (p1 + (p1 % ((-105) | 1)))), F1: p1, F2: 144}
                                        _ = v131
                                    case "a":
                                        var v132 int64 = -1
                                        v133 := GenS13{F0: 42}
                                        _ = v132
                                        _ = v133
                                    }
                                }
                                p1 = p1
                            }
                            v120 = (v120 >> ((4294967295) & 63))
                            v134 := GenS2{F0: (-(p1)), F1: -1}
                            _ = v134
                        }
                        v135 := []int{(fi126 + (-(fi127))), fi126, 32767}
                        _ = v129
                        _ = v135
                    }
                    switch fi126 {
                    case -922:
                        switch p0 {
                        case "hello":
                            if false {
                                v120 += (v120 >> (((p2 ^ (p2 ^ p2))) & 63))
                                p1 += (p1 + p1)
                                v136 := []int{(fi126 + 4294967295), int(v120), (fi126 + -19)}
                                _ = v136
                            } else {
                                var v137 uint64 = 4294967295
                                v138 := []int{}
                                switch 32767 {
                                case 9223372036854775807:
                                    p2 = (p2 + (^(p2)))
                                case 32767:
                                    var v139 uint = (p2 / (((p2 >> ((p2) & 63))) | 1))
                                    _ = v139
                                case 4294967295:
                                    v140 := GenS14{F0: p1}
                                    _ = v140
                                default:
                                    var v141 uint64 = v137
                                    _ = v141
                                }
                                _ = v137
                                _ = v138
                            }
                        default:
                            fmt.Println(p2, fi126)
                        }
                    case 441:
                        var v142 int64 = 756
                        _ = v142
                    }
                    var v143 uint = 1
                    v144 := []int{(fi126 * 506), fi127, 42}
                    _ = fi127
                    _ = v143
                    _ = v144
                }
                for fi145 := 0; fi145 < 16; fi145++ {
                    var v146 int32 = 42
                    _ = fi145
                    _ = v146
                }
                _ = fi126
            }
        case "hello":
            switch (p0 + p0) {
            case "foo bar":
                var v147 bool = (42 > 2)
                _ = v147
            case "":
                v148 := []int{127, 838}
                _ = v148
            default:
                if false {
                    if true {
                        v149 := []int{0, 32767}
                        _ = v149
                    } else {
                        v150 := []int{-1, -999, -640}
                        switch 9223372036854775807 {
                        case 203:
                            var v151 int32 = 2147483647
                            _ = v151
                        case 588:
                            v152 := GenS8{F0: 2}
                            _ = v152
                        case 85:
                            v153 := GenS11{F0: 954, F1: uint8(v120)}
                            _ = v153
                        }
                        if false {
                            for fi154 := 0; fi154 < 9; fi154++ {
                                v120 -= (v120 + (^(v120)))
                                v155 := GenS7{F0: "", F1: (4294967295 == 2), F2: 64}
                                switch (fi154 >> (((p2 % ((4294967295) | 1))) & 63)) {
                                case 127:
                                    for fi156 := 0; fi156 < 18; fi156++ {
                                        v157 := GenS15{F0: (^(p2)), F1: 2147483647, F2: 42}
                                        _ = fi156
                                        _ = v157
                                    }
                                case -9223372036854775808:
                                    var v158 float32 = 1.5
                                    _ = v158
                                default:
                                    var v159 uint16 = 258
                                    _ = v159
                                }
                                _ = fi154
                                _ = v155
                            }
                            v160 := GenS16{F0: 458}
                            if false {
                                v161 := GenS17{F0: 835, F1: -747, F2: (v120 / (((v120 % ((v120) | 1))) | 1))}
                                var v162 uint8 = 96
                                _ = v161
                                _ = v162
                            } else {
                                var v163 int16 = 33
                                v164 := GenS8{F0: 42}
                                _ = v163
                                _ = v164
                            }
                            _ = v160
                        }
                        _ = v150
                    }
                    for fi165 := 0; fi165 < 5; fi165++ {
                        fmt.Println(fi165, p1, fi165)
                        var v166 uint16 = 65535
                        _ = fi165
                        _ = v166
                    }
                } else {
                    v167 := GenS5{F0: p1, F1: (p1 << (((p2 / ((p2) | 1))) & 63)), F2: (p2 + (p2 - p2))}
                    var v168 int32 = 42
                    _ = v167
                    _ = v168
                }
            }
        default:
            var v169 int64 = 42
            _ = v169
        }
        switch 128 {
        case -73:
            switch p0 {
            case "foo bar":
                for fi170 := 0; fi170 < 2; fi170++ {
                    var v171 int32 = -2147483648
                    if true {
                        if true {
                            v172 := GenS18{F0: v171, F1: p0, F2: 196}
                            _ = v172
                        }
                        for fi173 := 0; fi173 < 6; fi173++ {
                            v174 := []int{(fi173 & (fi170 << (((p2 + p2)) & 63))), (fi173 / ((42) | 1)), (fi173 ^ fi170), (^(fi170))}
                            if (-70 == (-(p1))) {
                                var v175 float32 = 0.0
                                var v176 float32 = v175
                                v176 -= v176
                                fmt.Println(fi173)
                                if false {
                                    var v177 int32 = (v171 | -2147483648)
                                    fmt.Println(v175)
                                    _ = v177
                                } else {
                                    switch 42 {
                                    case -9223372036854775808:
                                        var v178 int16 = -32768
                                        _ = v178
                                    case -287:
                                        p0 = p0
                                        var v179 float64 = (1.5 - float64(((v176 - (0.1 - v175)) / ((0.0 + v175) + (float32(0.0) - 1.5)))))
                                        p2 = 2
                                        _ = v179
                                    case -442:
                                        if false {
                                            switch int(p1) {
                                            case 9223372036854775807:
                                                if (!(-1 >= 732)) {
                                                    switch "a" {
                                                    case "hello":
                                                        if true {
                                                            v175 = -2.25
                                                            var v180 bool = true
                                                            var v181 uint64 = 100
                                                            switch p0 {
                                                            case "a":
                                                                v182 := GenS8{F0: 1}
                                                                _ = v182
                                                            case "hello":
                                                                v174 = append(v174, fi173)
                                                            default:
                                                                fmt.Println(v120)
                                                            }
                                                            _ = v180
                                                            _ = v181
                                                        } else {
                                                            p2 = 65535
                                                            if true {
                                                                var v183 int8 = p1
                                                                _ = len(v174)
                                                                _ = v183
                                                            }
                                                        }
                                                    case "\n\t":
                                                        var v184 uint32 = 2147483647
                                                        _ = v184
                                                    }
                                                } else {
                                                    var v185 int16 = 476
                                                    var v186 uint = p2
                                                    _ = v185
                                                    _ = v186
                                                }
                                            }
                                            var v187 int = fi170
                                            v188 := GenS13{F0: 2147483648}
                                            _ = v187
                                            _ = v188
                                        }
                                    default:
                                        _ = len(v174)
                                    }
                                    var v189 uint = (p2 % ((p2) | 1))
                                    _ = v189
                                }
                                _ = v175
                                _ = v176
                            }
                            _ = fi173
                            _ = v174
                        }
                    }
                    _ = fi170
                    _ = v171
                }
            case "hello":
                v190 := GenS17{F0: v120, F1: 42, F2: v120}
                _ = v190
            }
        case -2147483648:
            fmt.Println(p0, p0)
        case -9223372036854775808:
            switch -250 {
            case -1:
                if false {
                    if (920 == uint64(v120)) {
                        for fi191 := 0; fi191 < 15; fi191++ {
                            if ((p1 <= p1) && (-1 <= 1)) {
                                if (((32767 == 42) || true) && (p2 != (^(p2)))) {
                                    for fi192 := 0; fi192 < 20; fi192++ {
                                        v193 := GenS14{F0: p1}
                                        var v194 int16 = 1
                                        _ = fi192
                                        _ = v193
                                        _ = v194
                                    }
                                    var v195 string = (p0 + p0)
                                    _ = v195
                                } else {
                                    fmt.Println(p1)
                                }
                            } else {
                                p1 -= p1
                                for fi196 := 0; fi196 < 15; fi196++ {
                                    var v197 uint16 = (v120 | (v120 + v120))
                                    _ = fi196
                                    _ = v197
                                }
                            }
                            p0 = "foo bar"
                            _ = fi191
                        }
                        var v198 uint16 = (v120 * (v120 >> ((p2) & 63)))
                        _ = v198
                    }
                    v120 = 2
                    var v199 uint16 = 65535
                    _ = v199
                } else {
                    var v200 int64 = -9223372036854775808
                    v201 := GenS19{F0: 18446744073709551615}
                    v202 := GenS14{F0: p1}
                    _ = v200
                    _ = v201
                    _ = v202
                }
            default:
                v120 -= (v120 << ((p2) & 63))
            }
        }
        _ = v120
    } else {
        var v203 float32 = (0.1 * 0.1)
        fmt.Println(p0, v203)
        _ = v203
    }
    switch p0 {
    case "":
        var v204 uint64 = uint64(p2)
        _ = v204
    }
    _ = p0
    _ = p1
    _ = p2
    return -364
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 255
    var v3 int32 = 2147483647
    v4 := make(map[uint32]bool)
    var v5 bool = false
    var v6 uint32 = 2
    if (!true) {
        if v5 {
            fmt.Println(v5)
            v205 := GenS20{F0: 4294967295, F1: v2}
            _ = v205
        }
        var v206 uint8 = 244
        v207 := GenS21{F0: v6, F1: -503}
        _ = v206
        _ = v207
    }
    switch "" {
    case "a":
        v3 += v3
    }
    for fi208 := 0; fi208 < 1; fi208++ {
        v209 := GenS5{F0: 42, F1: -126, F2: 144}
        fmt.Println(fi208, v6)
        v3 = genFn1(("hello" + ""), -1, 647)
        _ = fi208
        _ = v209
    }
    var v210 float32 = 1.5
    v5 = (v5 && (!true))
    switch -9223372036854775808 {
    case 9223372036854775807:
        var v211 bool = (-128 > 0)
        _ = v211
    }
    for fi212 := 0; fi212 < 2; fi212++ {
        v213 := GenS9{F0: true, F1: 1, F2: 2}
        _ = fi212
        _ = v213
    }
    switch "a" {
    case "a":
        switch int(v3) {
        case 9223372036854775807:
            v3 -= (^(v3))
        case 32768:
            v214 := GenS8{F0: 184}
            _ = v214
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v210:", v210)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v210
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
