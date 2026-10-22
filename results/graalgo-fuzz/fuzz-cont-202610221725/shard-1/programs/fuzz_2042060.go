package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int8 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 bool; F1 bool }
type GenS7 struct { F0 int64; F1 int }
type GenS8 struct { F0 int16; F1 uint64; F2 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int16 = -1
    var v3 int = 507
    var v4 string = "a"
    var v5 bool = false
    var v6 int64 = 42
    var v7 uint64 = 546
    switch int(v6) {
    case 42:
        var v8 uint = 639
        _ = v8
    }
    var v9 string = v4
    if true {
        if true {
            var v10 uint = 0
            if ((v10 | (^(v10))) != 42) {
                fmt.Println(v9, v6)
                v4 = ("\n\t" + v9)
            } else {
                if false {
                    if v5 {
                        if false {
                            var v11 uint = (v10 / (((v10 * 458)) | 1))
                            var v12 string = ((v9 + "a") + (("foo bar" + v4) + v4))
                            _ = v11
                            _ = v12
                        } else {
                            v13 := GenS1{F0: 42, F1: int16(v3), F2: -128}
                            fmt.Println(v3, v10, v10)
                            switch int(v10) {
                            case 4294967295:
                                v14 := GenS1{F0: 125, F1: v2, F2: -18}
                                _ = v14
                            }
                            _ = v13
                        }
                        for fi15 := 0; fi15 < 9; fi15++ {
                            for fi16 := 0; fi16 < 20; fi16++ {
                                switch fi15 {
                                case 624:
                                    var v17 int32 = 0
                                    _ = v17
                                default:
                                    v18 := v1[2]
                                    _ = v18
                                }
                                v19 := []int{}
                                _ = fi16
                                _ = v19
                            }
                            _ = fi15
                        }
                    } else {
                        var v20 uint16 = 255
                        for fi21 := 0; fi21 < 1; fi21++ {
                            v22 := GenS2{F0: -32768, F1: -128}
                            for fi23 := 0; fi23 < 5; fi23++ {
                                switch v4 {
                                case "":
                                    v4 = "foo bar"
                                }
                                var v24 uint = v10
                                fmt.Println(fi23)
                                _ = fi23
                                _ = v24
                            }
                            if v5 {
                                var v25 int64 = -9223372036854775808
                                _ = v25
                            } else {
                                fmt.Println(v5)
                                v10 = v10
                            }
                            _ = fi21
                            _ = v22
                        }
                        if (((v3 / ((v3) | 1)) < int(v6)) || (v6 == -36)) {
                            switch int(v7) {
                            case 127:
                                for fi26 := 0; fi26 < 15; fi26++ {
                                    switch v9 {
                                    case "a":
                                        var v27 float64 = 0.1
                                        _ = v27
                                    case "\n\t":
                                        v6 += (v6 / (((v6 / (((v6 ^ (v6 & v6))) | 1))) | 1))
                                    default:
                                        for fi28 := 0; fi28 < 11; fi28++ {
                                            v29 := []int{}
                                            v30 := GenS2{F0: v2, F1: 42}
                                            if (1.5 <= 1e10) {
                                                v31 := GenS2{F0: v2, F1: -98}
                                                _ = v31
                                            }
                                            _ = fi28
                                            _ = v29
                                            _ = v30
                                        }
                                    }
                                    _ = fi26
                                }
                            }
                        } else {
                            var v32 int64 = (v6 * v6)
                            v32 += (v32 - 0)
                            _ = v32
                        }
                        _ = v20
                    }
                }
            }
            v33 := GenS2{F0: v2, F1: 33}
            _ = v10
            _ = v33
        }
        delete(v1, 87)
        v6 = 0
    } else {
        if false {
            v34 := []int{v3, (v3 << ((4294967295) & 63)), (v3 & v3), (v3 % ((v3) | 1))}
            for fi35 := 0; fi35 < 17; fi35++ {
                switch 756 {
                case 128:
                    fmt.Println(v6, v4, v7)
                case 1:
                    var v36 uint32 = 65535
                    _ = v36
                }
                if len(v34) > 0 {
                    v34[0] = v3
                }
                v7 = (^(v7))
                _ = fi35
            }
            _ = v34
        }
    }
    for fi37 := 0; fi37 < 18; fi37++ {
        if v5 {
            v38 := []int{128, (fi37 >> ((4294967295) & 63)), fi37}
            v2 = v2
            _ = v38
        } else {
            fmt.Println(v4)
            v1[2147483647] = 0
        }
        _ = fi37
    }
    v4 = "\n\t"
    for fi39 := 0; fi39 < 2; fi39++ {
        v5 = v5
        _ = fi39
    }
    for fi40 := 0; fi40 < 3; fi40++ {
        fmt.Println(v6, v9, v5)
        v41 := GenS1{F0: 1, F1: v2, F2: -1}
        if true {
            for fi42 := 0; fi42 < 7; fi42++ {
                switch ((v9 + v4) + v4) {
                case "foo bar":
                    switch (v3 - (fi42 - fi42)) {
                    case -2147483648:
                        if (921 >= 869) {
                            v43 := GenS3{F0: 4294967295, F1: 998, F2: 2}
                            _ = v43
                        }
                    case 41:
                        switch (fi40 % (((fi40 | 196)) | 1)) {
                        case 0:
                            var v44 bool = v5
                            _ = v44
                        case -1:
                            var v45 uint16 = 42
                            var v46 string = v9
                            _ = v45
                            _ = v46
                        }
                    case 0:
                        if (-128 == -128) {
                            v47 := []int{v3, (fi40 - (v3 % (((-(fi42))) | 1))), 106}
                            var v48 uint8 = 0
                            v49 := GenS1{F0: -128, F1: (v2 ^ v2), F2: 100}
                            _ = v47
                            _ = v48
                            _ = v49
                        } else {
                            var v50 bool = true
                            v51 := GenS4{F0: v9}
                            if v5 {
                                var v52 bool = (v5 || true)
                                v5 = v5
                                switch -2147483648 {
                                case -592:
                                    fmt.Println(v3)
                                }
                                _ = v52
                            } else {
                                var v53 string = "a"
                                var v54 int32 = 2147483647
                                var v55 int = 221
                                if ((!false) || v5) {
                                    fmt.Println(v54, fi42)
                                    v56 := GenS5{F0: -128}
                                    v1[698] = 0
                                    _ = v56
                                }
                                switch (fi40 & (fi40 + (^(fi42)))) {
                                case 42:
                                    v6 -= v6
                                case -222:
                                    if (uint(v7) < 4294967295) {
                                        fmt.Println(v7)
                                        var v57 uint = 255
                                        if false {
                                            if (0 > 42) {
                                                var v58 uint8 = 1
                                                _ = v58
                                            }
                                            if v5 {
                                                v4 = "foo bar"
                                                v59 := GenS4{F0: ""}
                                                _ = v59
                                            }
                                        } else {
                                            fmt.Println(v7, fi42, v3)
                                        }
                                        _ = v57
                                    } else {
                                        var v60 float32 = ((1e10 - ((1.5 / 1e10) + 3.14159)) - 0.1)
                                        _ = v60
                                    }
                                case -786:
                                    v50 = ((v9 + "") < v53)
                                }
                                _ = v53
                                _ = v54
                                _ = v55
                            }
                            _ = v50
                            _ = v51
                        }
                    default:
                        v61 := []int{}
                        v1[2] = uint32(v7)
                        var v62 float64 = -2.25
                        _ = v61
                        _ = v62
                    }
                case "hello":
                    v6 -= (v6 << ((4294967295) & 63))
                case "a":
                    if false {
                        if true {
                            if ("\n\t" >= "\n\t") {
                                fmt.Println(fi40, v6)
                            }
                        }
                    } else {
                        v63 := []int{-1}
                        var v64 int64 = (v6 * v6)
                        var v65 int8 = 42
                        _ = v63
                        _ = v64
                        _ = v65
                    }
                    _ = len(v1)
                    if (v5 && (0 < 1)) {
                        v2 = (v2 ^ 1)
                        v6 -= v6
                        delete(v1, 2147483647)
                        fmt.Println(v7, v3, v6)
                        v2 = (v2 + (v2 & v2))
                        fmt.Println(v3)
                    }
                default:
                    if (1.5 < float32(1e10)) {
                        var v66 bool = v5
                        if v66 {
                            var v67 int = (^(fi42))
                            var v68 float32 = -2.25
                            var v69 float64 = 0.0
                            _ = v67
                            _ = v68
                            _ = v69
                        } else {
                            var v70 int = -764
                            _ = v70
                        }
                        _ = v66
                    } else {
                        if (1 >= -30) {
                            _ = len(v1)
                            switch v4 {
                            case "a":
                                var v71 uint = 18446744073709551615
                                _ = v71
                            case "foo bar":
                                var v72 uint = 215
                                _ = v72
                            }
                        } else {
                            fmt.Println(v7)
                            if v5 {
                                v9 = (v9 + v4)
                            } else {
                                fmt.Println(v3, v5)
                                _ = len(v1)
                                var v73 int = 127
                                fmt.Println(v2)
                                var v74 uint64 = v7
                                _ = v73
                                _ = v74
                            }
                            v75 := []int{802}
                            var v76 int16 = v2
                            _ = v75
                            _ = v76
                        }
                        switch (-(v3)) {
                        case 698:
                            switch (fi42 | 0) {
                            case -829:
                                v5 = (-689 == -2147483648)
                                v2 -= 1000
                            }
                        case 0:
                            _ = len(v1)
                        case 127:
                            var v77 bool = v5
                            v6 -= v6
                            v2 -= (v2 ^ v2)
                            _ = v77
                        default:
                            delete(v1, 566)
                        }
                    }
                }
                v6 += int64(v2)
                if true {
                    fmt.Println(v9, v5)
                    var v78 uint64 = 0
                    _ = v78
                } else {
                    v79 := GenS2{F0: 42, F1: 127}
                    var v80 bool = v5
                    v81 := []int{}
                    _ = v79
                    _ = v80
                    _ = v81
                }
                var v82 float32 = 0.0
                if false {
                    if (true || (0 < 246)) {
                        v83 := GenS1{F0: 42, F1: 0, F2: 11}
                        fmt.Println(v4)
                        var v84 float32 = v82
                        _ = v83
                        _ = v84
                    } else {
                        if (662 < 65535) {
                            var v85 uint = 25
                            if (fi40 > 32767) {
                                fmt.Println(v82, v4, v7)
                                if v5 {
                                    v86 := v1[460]
                                    if v5 {
                                        fmt.Println(v2, v4, fi40)
                                        v7 = v7
                                        if true {
                                            var v87 float64 = (0.0 * (1e10 - 1e10))
                                            var v88 int = v3
                                            _ = v87
                                            _ = v88
                                        } else {
                                            switch (^(v3)) {
                                            case 191:
                                                switch ("hello" + ("" + v9)) {
                                                case "a":
                                                    delete(v1, (v86 >> ((uint(v7)) & 63)))
                                                }
                                            case 2147483647:
                                                var v89 string = ("foo bar" + "hello")
                                                _ = v89
                                            case -1:
                                                if v5 {
                                                    fmt.Println(v2)
                                                    v90 := GenS1{F0: -45, F1: v2, F2: 42}
                                                    _ = v90
                                                }
                                                switch (fi40 % ((105) | 1)) {
                                                case 0:
                                                    v91 := []int{fi40, fi42, v3}
                                                    _ = v91
                                                case 4294967295:
                                                    var v92 float32 = v82
                                                    _ = v92
                                                case -1:
                                                    v93 := GenS1{F0: 1, F1: 32767, F2: 100}
                                                    v6 -= 738
                                                    _ = v93
                                                default:
                                                    v3 -= fi42
                                                }
                                            default:
                                                _ = len(v1)
                                            }
                                            fmt.Println(fi42, v82)
                                            if v5 {
                                                var v94 int64 = v6
                                                v82 = 1.5
                                                var v95 string = (v9 + v9)
                                                v96 := GenS6{F0: (v5 || v5), F1: (v95 > "foo bar")}
                                                if ((false && false) && v5) {
                                                    if (v5 && true) {
                                                        var v97 int32 = 1
                                                        v82 = (((v82 * 0.1) + (v82 + (1.5 * (v82 + v82)))) + 3.14159)
                                                        _ = v97
                                                    }
                                                    var v98 uint8 = 114
                                                    var v99 uint64 = v7
                                                    var v100 uint32 = (v86 * (v86 / (((^(v86))) | 1)))
                                                    var v101 uint16 = 65535
                                                    _ = v100
                                                    _ = v101
                                                    _ = v98
                                                    _ = v99
                                                } else {
                                                    var v102 int = fi40
                                                    v2 -= (v2 & (v2 >> ((v85) & 63)))
                                                    _ = v102
                                                }
                                                _ = v94
                                                _ = v95
                                                _ = v96
                                            }
                                        }
                                    }
                                    var v103 uint8 = 234
                                    _ = v103
                                    _ = v86
                                } else {
                                    fmt.Println(v3, v2)
                                    var v104 int8 = -128
                                    v105 := GenS3{F0: 482, F1: 2, F2: 613}
                                    v7 -= 638
                                    _ = v104
                                    _ = v105
                                }
                                var v106 int = (fi40 ^ (v3 - fi42))
                                delete(v1, 2)
                                v5 = ((v5 && (v5 && (v4 > v9))) || (1 < 18446744073709551615))
                                var v107 int64 = (v6 + (v6 ^ (^(v6))))
                                _ = v106
                                _ = v107
                            } else {
                                v108 := v1[489]
                                _ = v108
                            }
                            _ = v85
                        } else {
                            var v109 uint64 = 2
                            v110 := GenS6{F0: (!(0 <= -1)), F1: v5}
                            _ = v109
                            _ = v110
                        }
                        var v111 float32 = (v82 + 3.14159)
                        _ = v111
                    }
                    v7 = (^(v7))
                    v1[42] = 42
                }
                _ = fi42
                _ = v82
            }
            var v112 int = (v3 >> ((65535) & 63))
            v5 = false
            _ = v112
        }
        _ = fi40
        _ = v41
    }
    for fi113 := 0; fi113 < 13; fi113++ {
        for fi114 := 0; fi114 < 20; fi114++ {
            var v115 uint8 = 137
            v115 += (^(v115))
            _ = fi114
            _ = v115
        }
        var v116 int16 = (v2 - v2)
        if ((v9 + v4) != "a") {
            if true {
                v4 = ((v4 + (v9 + (v9 + "a"))) + v4)
                v9 = ("\n\t" + "\n\t")
                v5 = (false || (!true))
            } else {
                for fi117 := 0; fi117 < 2; fi117++ {
                    var v118 uint64 = v7
                    v5 = v5
                    var v119 float64 = (((float64(((0.0 + 1e10) / float32(-2.25))) - 1.5) + 1.5) * -2.25)
                    _ = fi117
                    _ = v118
                    _ = v119
                }
                if v5 {
                    if v5 {
                        for fi120 := 0; fi120 < 9; fi120++ {
                            var v121 int32 = -71
                            v122 := []int{(v3 % ((4294967295) | 1))}
                            v4 = v4
                            _ = fi120
                            _ = v121
                            _ = v122
                        }
                    } else {
                        var v123 uint8 = 42
                        v124 := v1[0]
                        _ = v123
                        _ = v124
                    }
                    v7 -= v7
                    switch ("\n\t" + v4) {
                    case "\n\t":
                        var v125 int8 = -65
                        _ = v125
                    default:
                        v9 = (((v9 + ("\n\t" + (v4 + ((("\n\t" + (((v9 + "") + v9) + v9)) + "a") + v9)))) + "") + "")
                    }
                }
                v126 := GenS7{F0: v6, F1: 0}
                _ = v126
            }
        }
        _ = fi113
        _ = v116
    }
    v7 = (v7 >> ((655) & 63))
    if (68 != 100) {
        v127 := GenS3{F0: 2147483647, F1: 255, F2: 682}
        v3 += (v3 * 42)
        _ = v127
    } else {
        fmt.Println(v9, v2)
        if (v2 <= v2) {
            if false {
                if (v5 || v5) {
                    fmt.Println(v7)
                    if (v4 <= v9) {
                        fmt.Println(v2, v2)
                        switch "" {
                        case "a":
                            v5 = true
                        default:
                            v9 = v4
                        }
                        var v128 int = (-(v3))
                        _ = v128
                    } else {
                        v129 := GenS3{F0: 754, F1: 32767, F2: 1}
                        fmt.Println(v2, v9)
                        _ = v129
                    }
                    v6 = -380
                } else {
                    v6 -= v6
                    for fi130 := 0; fi130 < 14; fi130++ {
                        v3 -= -9223372036854775808
                        _ = fi130
                    }
                    switch -2147483648 {
                    case 32768:
                        for fi131 := 0; fi131 < 8; fi131++ {
                            var v132 int32 = 42
                            fmt.Println(v4)
                            _ = fi131
                            _ = v132
                        }
                    case 128:
                        if (((193 == 65535) && v5) || ((3.14159 - 1e10) > (float32((3.14159 - 0.0)) * 0.1))) {
                            var v133 uint16 = 32767
                            delete(v1, 656)
                            _ = v133
                        } else {
                            v7 += v7
                        }
                    case -2147483648:
                        v2 = 1000
                    }
                }
                var v134 int = v3
                _ = v134
            } else {
                v135 := GenS2{F0: v2, F1: -128}
                for fi136 := 0; fi136 < 19; fi136++ {
                    if (false || (v3 < (fi136 + 42))) {
                        if v5 {
                            v6 = (v6 + (v6 >> ((2) & 63)))
                            v7 = v7
                        }
                    }
                    if v5 {
                        var v137 float32 = -2.25
                        var v138 float32 = (0.0 / v137)
                        _ = v137
                        _ = v138
                    } else {
                        v7 = (v7 % ((v7) | 1))
                    }
                    v139 := []int{v3, v3, v3, v3}
                    _ = fi136
                    _ = v139
                }
                _ = v135
            }
            v4 = v4
            for fi140 := 0; fi140 < 15; fi140++ {
                v141 := []int{749, (v3 ^ v3), (^(v3))}
                v6 -= 9223372036854775807
                var v142 uint = uint(v7)
                _ = fi140
                _ = v141
                _ = v142
            }
        } else {
            var v143 uint16 = 42
            if false {
                fmt.Println(v5)
            } else {
                v144 := GenS8{F0: -32768, F1: v7, F2: 127}
                v1[387] = 2147483648
                var v145 uint16 = v143
                _ = v144
                _ = v145
            }
            v146 := GenS5{F0: 40}
            _ = v143
            _ = v146
        }
        for fi147 := 0; fi147 < 1; fi147++ {
            var v148 float32 = 3.14159
            for fi149 := 0; fi149 < 5; fi149++ {
                v150 := GenS8{F0: v2, F1: 2, F2: -104}
                var v151 float32 = 3.14159
                _ = fi149
                _ = v150
                _ = v151
            }
            _ = fi147
            _ = v148
        }
    }
    v152 := GenS5{F0: 127}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v152
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
