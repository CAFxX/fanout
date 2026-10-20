package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 int32; F1 int8; F2 int8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint16; F1 int32; F2 bool }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int32; F1 bool }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 int }
type GenS16 struct { F0 uint; F1 int64 }
type GenS17 struct { F0 int64 }
type GenS18 struct { F0 int16; F1 int8; F2 uint16 }
type GenS19 struct { F0 bool; F1 uint16 }
type GenS20 struct { F0 int8; F1 uint }
type GenS21 struct { F0 int32; F1 bool }
type GenS22 struct { F0 string; F1 string }
func genFn0(p0 uint16) int32 {
    v5 := GenS1{F0: 42, F1: -2147483648}
    if false {
        p0 = 255
        v6 := []int{}
        _ = v6
    }
    if true {
        if (!true) {
            v7 := []int{-9223372036854775808, -659, 127, 128}
            p0 = 32767
            v8 := GenS2{F0: 68}
            _ = v7
            _ = v8
        } else {
            var v9 float32 = 0.0
            v10 := GenS3{F0: p0}
            var v11 int64 = 1
            _ = v10
            _ = v11
            _ = v9
        }
    } else {
        switch "hello" {
        case "":
            p0 += (^(p0))
        case "a":
            v12 := []int{1, 1, -284}
            _ = v12
        }
    }
    _ = p0
    _ = v5
    return 0
}

func main() {
    var v1 int8 = 1
    var v2 float32 = 0.1
    v3 := make(map[int]string)
    v4 := make(map[uint32]string)
    var v13 uint64 = 1
    if false {
        var v14 bool = (190 < 1)
        var v15 uint16 = 65535
        for fi16 := 0; fi16 < 16; fi16++ {
            v17 := GenS4{F0: "hello", F1: v13}
            if true {
                for fi18 := 0; fi18 < 4; fi18++ {
                    v19 := GenS5{F0: (^(v15)), F1: 2147483648}
                    _ = fi18
                    _ = v19
                }
                var v20 int16 = 597
                _ = v20
            } else {
                for fi21 := 0; fi21 < 20; fi21++ {
                    v22 := GenS1{F0: 1, F1: genFn0(255)}
                    if ((!((66 <= -131) && false)) || (1e10 < float64(0.1))) {
                        v23 := GenS6{F0: genFn0(1), F1: 1, F2: -128}
                        var v24 int16 = 1000
                        v25 := GenS7{F0: "hello"}
                        var v26 int32 = 496
                        _ = v23
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    v27 := GenS1{F0: 864, F1: 1}
                    _ = fi21
                    _ = v22
                    _ = v27
                }
                var v28 uint32 = 4294967295
                var v29 int = (-(fi16))
                _ = v28
                _ = v29
            }
            _ = fi16
            _ = v17
        }
        _ = v14
        _ = v15
    } else {
        var v30 int64 = 0
        var v31 string = "\n\t"
        if false {
            delete(v4, 0)
            if true {
                v32 := GenS2{F0: 42}
                v3[128] = v31
                for fi33 := 0; fi33 < 5; fi33++ {
                    _ = len(v3)
                    _ = fi33
                }
                _ = v32
            } else {
                var v34 int8 = (v1 * 1)
                var v35 int16 = 42
                fmt.Println(v2, v34, v13)
                _ = v34
                _ = v35
            }
            switch v31 {
            case "foo bar":
                if false {
                    if (65535 <= 72) {
                        if false {
                            if (527 != 720) {
                                v1 += (v1 ^ (v1 ^ 1))
                            }
                            for fi36 := 0; fi36 < 10; fi36++ {
                                for fi37 := 0; fi37 < 15; fi37++ {
                                    v1 = (v1 & v1)
                                    var v38 float64 = ((1e10 - -2.25) - -2.25)
                                    v31 = "\n\t"
                                    _ = fi37
                                    _ = v38
                                }
                                v39 := GenS8{F0: 966, F1: -377, F2: true}
                                _ = fi36
                                _ = v39
                            }
                        } else {
                            var v40 int = 4294967295
                            v41 := GenS1{F0: v30, F1: 0}
                            fmt.Println(v40)
                            _ = v40
                            _ = v41
                        }
                        switch "hello" {
                        case "foo bar":
                            if (!false) {
                                for fi42 := 0; fi42 < 12; fi42++ {
                                    v43 := []int{-9223372036854775808, (^(fi42)), fi42, (fi42 % ((544) | 1))}
                                    var v44 uint64 = v13
                                    _ = fi42
                                    _ = v43
                                    _ = v44
                                }
                                if (42 > 427) {
                                    delete(v4, 42)
                                    for fi45 := 0; fi45 < 4; fi45++ {
                                        if true {
                                            var v46 int64 = v30
                                            for fi47 := 0; fi47 < 20; fi47++ {
                                                fmt.Println(v46, v46)
                                                v48 := GenS9{F0: 794}
                                                var v49 int16 = 1000
                                                if false {
                                                    var v50 bool = true
                                                    fmt.Println(fi45, v50)
                                                    _ = v50
                                                } else {
                                                    var v51 uint32 = 33
                                                    v52 := GenS9{F0: 85}
                                                    v53 := GenS1{F0: v30, F1: 42}
                                                    _ = v51
                                                    _ = v52
                                                    _ = v53
                                                }
                                                v46 -= (-(v46))
                                                _ = fi47
                                                _ = v48
                                                _ = v49
                                            }
                                            switch v31 {
                                            case "":
                                                v13 = v13
                                            case "\n\t":
                                                v54 := []int{(fi45 / ((fi45) | 1)), fi45}
                                                _ = v54
                                            case "foo bar":
                                                if true {
                                                    v55 := GenS10{F0: 4294967295}
                                                    if (634 <= 0) {
                                                        if true {
                                                            v1 = (-(v1))
                                                            v2 += 3.14159
                                                        }
                                                    }
                                                    _ = v55
                                                }
                                            }
                                            _ = v46
                                        } else {
                                            v56 := GenS2{F0: 1}
                                            delete(v3, 9223372036854775807)
                                            v1 += v1
                                            _ = v56
                                        }
                                        _ = fi45
                                    }
                                } else {
                                    switch (v31 + (("\n\t" + "\n\t") + (v31 + "a"))) {
                                    case "":
                                        v57 := GenS7{F0: v31}
                                        _ = v57
                                    case "\n\t":
                                        var v58 string = v31
                                        _ = v58
                                    }
                                    switch v31 {
                                    case "a":
                                        fmt.Println(v31, v2)
                                    default:
                                        switch -729 {
                                        case 128:
                                            v2 = (3.14159 + float32((((float64(v2) / float64(v2)) * -2.25) * 1e10)))
                                        default:
                                            for fi59 := 0; fi59 < 12; fi59++ {
                                                var v60 float64 = 0.0
                                                v13 += (v13 >> ((42) & 63))
                                                _ = fi59
                                                _ = v60
                                            }
                                        }
                                    }
                                    delete(v4, 42)
                                }
                            }
                        case "":
                            v2 -= (v2 * (3.14159 - 1e10))
                        case "\n\t":
                            for fi61 := 0; fi61 < 20; fi61++ {
                                for fi62 := 0; fi62 < 5; fi62++ {
                                    var v63 string = "hello"
                                    v64 := GenS11{F0: 1}
                                    if false {
                                        switch (^(fi61)) {
                                        case 0:
                                            var v65 int = -2147483648
                                            _ = v65
                                        case 1:
                                            fmt.Println(v2, v13)
                                        case 4294967295:
                                            if true {
                                                v1 = (v1 + (^(v1)))
                                            } else {
                                                var v66 int16 = -32768
                                                _ = v66
                                            }
                                        default:
                                            if false {
                                                v67 := []int{(fi61 << ((964) & 63)), (fi62 | (fi61 ^ fi61))}
                                                var v68 uint = 255
                                                switch fi62 {
                                                case 872:
                                                    delete(v4, 334)
                                                    if false {
                                                        fmt.Println(v63)
                                                    }
                                                    v13 = v13
                                                case -709:
                                                    v69 := GenS12{F0: 605}
                                                    _ = v69
                                                case 1:
                                                    switch v63 {
                                                    case "foo bar":
                                                        v31 = v31
                                                    case "a":
                                                        var v70 int8 = v1
                                                        _ = v70
                                                    default:
                                                        var v71 float64 = (1e10 * (float64(v2) / (float64(float32(((0.1 * 1e10) - 3.14159))) * (((3.14159 - 1e10) * 1e10) + 1.5))))
                                                        _ = v71
                                                    }
                                                }
                                                _ = v67
                                                _ = v68
                                            }
                                        }
                                    } else {
                                        v63 = "foo bar"
                                    }
                                    v30 -= (v30 | (v30 - v30))
                                    _ = fi62
                                    _ = v63
                                    _ = v64
                                }
                                for fi72 := 0; fi72 < 10; fi72++ {
                                    if true {
                                        var v73 uint = 119
                                        v73 -= v73
                                        _ = v73
                                    }
                                    _ = len(v4)
                                    _ = fi72
                                }
                                if true {
                                    if false {
                                        switch 32767 {
                                        case -9223372036854775808:
                                            v74 := GenS4{F0: v31, F1: v13}
                                            _ = v74
                                        default:
                                            v2 = v2
                                        }
                                        fmt.Println(v1, v1, v2)
                                        if (-256 < genFn0(255)) {
                                            var v75 uint32 = 0
                                            _ = v75
                                        }
                                    } else {
                                        var v76 int32 = genFn0(65535)
                                        _ = v76
                                    }
                                } else {
                                    var v77 int32 = -2147483648
                                    v13 = (v13 + 751)
                                    _ = v77
                                }
                                _ = fi61
                            }
                        default:
                            if (2147483648 == 2147483647) {
                                for fi78 := 0; fi78 < 11; fi78++ {
                                    if ((!true) || (v13 > 1)) {
                                        v30 += (v30 / ((1) | 1))
                                        v3[fi78] = ("foo bar" + (v31 + "a"))
                                        if false {
                                            switch (fi78 - fi78) {
                                            case 9223372036854775807:
                                                switch "a" {
                                                case "\n\t":
                                                    if ((3.14159 != 0.1) || (-587 > 1000)) {
                                                        var v79 uint = 0
                                                        fmt.Println(v79)
                                                        for fi80 := 0; fi80 < 4; fi80++ {
                                                            if true {
                                                                switch v31 {
                                                                case "foo bar":
                                                                    switch int(v1) {
                                                                    case -160:
                                                                        v81 := []int{32767}
                                                                        _ = v81
                                                                    }
                                                                case "a":
                                                                    fmt.Println(v79, v2, v30)
                                                                }
                                                                var v82 uint16 = 899
                                                                _ = v82
                                                            } else {
                                                                v83 := GenS13{F0: -672, F1: (false || true)}
                                                                v84 := []int{fi78}
                                                                _ = len(v84)
                                                                _ = v83
                                                                _ = v84
                                                            }
                                                            if true {
                                                                v85 := []int{-704, -1, (fi80 + (^(fi78))), -156}
                                                                fmt.Println(fi78)
                                                                _ = v85
                                                            }
                                                            v30 = (v30 | (v30 / ((v30) | 1)))
                                                            if false {
                                                                v31 = v31
                                                            } else {
                                                                v30 = v30
                                                            }
                                                            fmt.Println(fi78, fi78, v30)
                                                            delete(v4, 42)
                                                            _ = fi80
                                                        }
                                                        _ = v79
                                                    }
                                                default:
                                                    for fi86 := 0; fi86 < 11; fi86++ {
                                                        v87 := GenS7{F0: v31}
                                                        if true {
                                                            var v88 uint32 = 2147483647
                                                            v88 -= v88
                                                            v89 := []int{(fi78 * fi78), (-(fi78)), (fi86 << ((529) & 63))}
                                                            if true {
                                                                fmt.Println(v88, v1)
                                                                v90 := v4[(v88 >> ((2) & 63))]
                                                                var v91 uint8 = 42
                                                                _ = v90
                                                                _ = v91
                                                            }
                                                            v13 -= v13
                                                            _ = v88
                                                            _ = v89
                                                        }
                                                        _ = fi86
                                                        _ = v87
                                                    }
                                                }
                                            case 32767:
                                                var v92 uint16 = 326
                                                _ = v92
                                            case 227:
                                                fmt.Println(v1)
                                            default:
                                                var v93 int8 = v1
                                                _ = v93
                                            }
                                            v94 := GenS14{F0: 1}
                                            var v95 uint64 = v13
                                            _ = v94
                                            _ = v95
                                        }
                                    }
                                    v1 += -128
                                    v30 = 42
                                    _ = fi78
                                }
                            }
                        }
                        _ = len(v4)
                    }
                    v96 := []int{9223372036854775807, 4294967295}
                    v30 = v30
                    _ = v96
                } else {
                    if (1 >= 883) {
                        if false {
                            for fi97 := 0; fi97 < 2; fi97++ {
                                var v98 string = v31
                                _ = fi97
                                _ = v98
                            }
                        } else {
                            var v99 uint8 = 2
                            if (-1 >= genFn0(uint16(v13))) {
                                v100 := GenS15{F0: -2147483648}
                                v101 := []int{}
                                _ = v100
                                _ = v101
                            }
                            v99 = 91
                            _ = v99
                        }
                        fmt.Println(v13, v2, v2)
                    } else {
                        for fi102 := 0; fi102 < 10; fi102++ {
                            v103 := GenS5{F0: 401, F1: 974}
                            v104 := []int{}
                            var v105 uint8 = 42
                            _ = fi102
                            _ = v103
                            _ = v104
                            _ = v105
                        }
                    }
                    if (!(!(4294967295 >= 18446744073709551615))) {
                        v106 := []int{42, 42}
                        for fi107 := 0; fi107 < 12; fi107++ {
                            v1 += v1
                            _ = fi107
                        }
                        _ = v106
                    } else {
                        switch v31 {
                        case "":
                            for fi108 := 0; fi108 < 16; fi108++ {
                                v109 := []int{584, 928}
                                if (100 < (v1 >> ((271) & 63))) {
                                    if false {
                                        var v110 uint64 = v13
                                        var v111 bool = ((2147483647 == genFn0(42)) && ((!("foo bar" >= "a")) || (895 != 1)))
                                        _ = v110
                                        _ = v111
                                    }
                                    var v112 int8 = -82
                                    _ = v112
                                }
                                switch v31 {
                                case "foo bar":
                                    v31 = "\n\t"
                                default:
                                    v113 := GenS16{F0: 0, F1: -9223372036854775808}
                                    _ = v113
                                }
                                _ = fi108
                                _ = v109
                            }
                        case "\n\t":
                            v114 := GenS17{F0: (v30 ^ v30)}
                            _ = v114
                        case "a":
                            var v115 int = 1
                            _ = v115
                        }
                        for fi116 := 0; fi116 < 2; fi116++ {
                            for fi117 := 0; fi117 < 15; fi117++ {
                                if true {
                                    if false {
                                        switch v31 {
                                        case "":
                                            fmt.Println(fi116, v1, fi116)
                                        case "\n\t":
                                            var v118 uint16 = 986
                                            _ = v118
                                        default:
                                            switch (fi116 << ((65535) & 63)) {
                                            case -338:
                                                v4[833] = v31
                                            }
                                        }
                                    }
                                }
                                var v119 int32 = 1
                                if (!true) {
                                    switch fi117 {
                                    case 876:
                                        var v120 uint = 1
                                        _ = v120
                                    case 42:
                                        var v121 float64 = 0.1
                                        _ = v121
                                    default:
                                        v1 = (v1 - (v1 + (v1 ^ v1)))
                                    }
                                    if true {
                                        v122 := v3[2147483647]
                                        _ = v122
                                    }
                                    v2 = v2
                                } else {
                                    switch fi117 {
                                    case 127:
                                        var v123 int64 = int64(v13)
                                        _ = v123
                                    }
                                    v124 := []int{}
                                    v125 := GenS11{F0: 644}
                                    _ = v124
                                    _ = v125
                                }
                                _ = fi117
                                _ = v119
                            }
                            if ((v13 % (((^(v13))) | 1)) < v13) {
                                if true {
                                    var v126 int64 = (v30 / ((-1) | 1))
                                    var v127 uint16 = 1
                                    switch 2147483647 {
                                    case 128:
                                        for fi128 := 0; fi128 < 2; fi128++ {
                                            v129 := GenS15{F0: (-(fi116))}
                                            _ = fi128
                                            _ = v129
                                        }
                                    case 4294967295:
                                        if false {
                                            var v130 float64 = (float64((v2 * (0.0 - (0.0 - ((v2 * ((v2 + 0.1) - v2)) - (v2 - ((((((v2 - 0.1) / (3.14159 / (v2 + v2))) + 0.1) / ((v2 + float32(1.5)) + ((float32(1.5) * 3.14159) - v2))) - (v2 + v2)) * (v2 * v2)))))))) - ((0.1 - -2.25) + 1e10))
                                            var v131 string = ""
                                            _ = v130
                                            _ = v131
                                        } else {
                                            var v132 int = fi116
                                            _ = v132
                                        }
                                    case 1:
                                        v126 += (v30 % (((v30 << ((365) & 63))) | 1))
                                    default:
                                        v133 := GenS16{F0: 699, F1: v126}
                                        _ = v133
                                    }
                                    _ = v126
                                    _ = v127
                                }
                                fmt.Println(v13, v31, v31)
                                v1 -= v1
                            } else {
                                delete(v4, 819)
                            }
                            _ = fi116
                        }
                        v1 -= 98
                    }
                    if (!(-2.25 <= (v2 * v2))) {
                        for fi134 := 0; fi134 < 10; fi134++ {
                            var v135 int32 = -332
                            _ = fi134
                            _ = v135
                        }
                    }
                }
            }
        }
        _ = v30
        _ = v31
    }
    switch "\n\t" {
    case "":
        v13 += (^(v13))
    case "\n\t":
        v1 += (v1 ^ (v1 * (v1 / (((^(v1))) | 1))))
    case "a":
        var v136 int8 = v1
        _ = v136
    default:
        if false {
            for fi137 := 0; fi137 < 14; fi137++ {
                v2 = v2
                if (!true) {
                    for fi138 := 0; fi138 < 15; fi138++ {
                        if (!true) {
                            if false {
                                if false {
                                    fmt.Println(v13, v13)
                                    var v139 string = ""
                                    _ = v139
                                } else {
                                    switch "" {
                                    case "a":
                                        v4[42] = (("" + (("\n\t" + ("" + ((("hello" + (("a" + "\n\t") + (((("" + ("a" + ("a" + "\n\t"))) + "") + ("foo bar" + "foo bar")) + ("foo bar" + "\n\t")))) + ("" + "hello")) + "hello"))) + "foo bar")) + ("\n\t" + "foo bar"))
                                    case "":
                                        v140 := v4[5]
                                        _ = v140
                                    case "foo bar":
                                        delete(v3, (^(fi138)))
                                    default:
                                        var v141 uint8 = 221
                                        _ = v141
                                    }
                                }
                                var v142 float32 = 1.5
                                _ = v142
                            } else {
                                var v143 int32 = -756
                                v2 = 0.1
                                switch "hello" {
                                case "\n\t":
                                    var v144 uint64 = v13
                                    _ = v144
                                }
                                if (((0.1 - 0.1) == (1e10 / float64(v2))) || false) {
                                    var v145 int8 = (v1 / (((v1 / ((v1) | 1))) | 1))
                                    _ = v145
                                }
                                v146 := GenS1{F0: -880, F1: (v143 << ((255) & 63))}
                                _ = v143
                                _ = v146
                            }
                            var v147 float64 = 1.5
                            _ = v147
                        } else {
                            v148 := GenS14{F0: 1}
                            _ = v148
                        }
                        var v149 string = "hello"
                        if (((!((^(v13)) == v13)) || (42 != 972)) && (-986 <= 0)) {
                            var v150 uint = 917
                            fmt.Println(v149, v149)
                            var v151 int16 = 1
                            _ = v150
                            _ = v151
                        }
                        v152 := []int{fi138, fi137}
                        _ = fi138
                        _ = v149
                        _ = v152
                    }
                } else {
                    var v153 int16 = 32767
                    if true {
                        var v154 float64 = float64(v2)
                        _ = v154
                    } else {
                        fmt.Println(fi137, fi137)
                        var v155 int32 = 569
                        _ = v155
                    }
                    _ = v153
                }
                if true {
                    var v156 bool = (42 > 42)
                    v156 = (!true)
                    _ = v156
                }
                _ = fi137
            }
            for fi157 := 0; fi157 < 4; fi157++ {
                if (!false) {
                    var v158 float64 = (1.5 - float64(((3.14159 / ((v2 * v2) / v2)) + (v2 - v2))))
                    fmt.Println(v1)
                    var v159 bool = ((^(v13)) > (v13 / ((v13) | 1)))
                    _ = v158
                    _ = v159
                }
                fmt.Println(v13, v2)
                for fi160 := 0; fi160 < 15; fi160++ {
                    v2 -= v2
                    _ = fi160
                }
                _ = fi157
            }
            switch int(v1) {
            case 127:
                var v161 int8 = (v1 & v1)
                _ = v161
            case 42:
                switch "a" {
                case "hello":
                    for fi162 := 0; fi162 < 15; fi162++ {
                        fmt.Println(v1)
                        switch int(v13) {
                        case 247:
                            if (!false) {
                                v163 := []int{(fi162 | (-(fi162)))}
                                v164 := GenS11{F0: 42}
                                _ = v163
                                _ = v164
                            } else {
                                var v165 int16 = -871
                                _ = v165
                            }
                        case 135:
                            fmt.Println(v13, v2)
                        default:
                            var v166 int32 = -2147483648
                            _ = v166
                        }
                        v2 -= -2.25
                        _ = fi162
                    }
                case "":
                    fmt.Println(v1, v2, v13)
                case "a":
                    var v167 bool = false
                    _ = v167
                }
            default:
                v168 := GenS18{F0: 0, F1: (v1 | v1), F2: 65535}
                _ = v168
            }
        } else {
            _ = len(v3)
            _ = len(v3)
        }
    }
    v169 := GenS19{F0: (388 < 2147483648), F1: 925}
    v170 := GenS4{F0: "hello", F1: (v13 / ((v13) | 1))}
    if false {
        delete(v3, -9223372036854775808)
        v171 := []int{}
        v172 := GenS6{F0: -1, F1: -128, F2: 1}
        _ = v171
        _ = v172
    } else {
        v173 := []int{}
        var v174 bool = (32767 != -263)
        if false {
            var v175 int = 32768
            _ = v175
        }
        _ = v173
        _ = v174
    }
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        var v176 int16 = -455
        _ = v176
    case "hello":
        switch "" {
        case "foo bar":
            for fi177 := 0; fi177 < 19; fi177++ {
                fmt.Println(v13, v1)
                for fi178 := 0; fi178 < 17; fi178++ {
                    v4[1] = "foo bar"
                    v179 := []int{(fi178 % ((fi178) | 1)), fi177}
                    _ = fi178
                    _ = v179
                }
                var v180 float64 = 3.14159
                _ = fi177
                _ = v180
            }
        case "hello":
            for fi181 := 0; fi181 < 15; fi181++ {
                var v182 int32 = 1
                _ = fi181
                _ = v182
            }
        case "":
            fmt.Println(v1)
        }
    case "a":
        if true {
            if (127 >= 2) {
                switch 4294967295 {
                case 475:
                    v1 = 43
                case 9223372036854775807:
                    fmt.Println(v13)
                }
                v183 := GenS4{F0: "a", F1: (v13 % ((v13) | 1))}
                var v184 int64 = 0
                _ = v183
                _ = v184
            } else {
                v2 = (v2 - v2)
                switch "a" {
                case "":
                    var v185 int16 = 1000
                    _ = v185
                case "foo bar":
                    var v186 float64 = 0.1
                    _ = v186
                case "a":
                    v187 := GenS20{F0: v1, F1: 644}
                    _ = v187
                }
                _ = len(v3)
            }
        }
    default:
        v188 := []int{32767, 13, 32767, 32767}
        _ = v188
    }
    v189 := []int{}
    for fi190 := 0; fi190 < 17; fi190++ {
        if true {
            if len(v189) >= 2 {
                v189 = v189[1:len(v189)]
            }
        }
        _ = fi190
    }
    v191 := GenS21{F0: -2147483648, F1: false}
    if false {
        switch "a" {
        case "\n\t":
            for fi192 := 0; fi192 < 15; fi192++ {
                fmt.Println(fi192, v2)
                _ = fi192
            }
        case "":
            v13 += 18446744073709551615
        default:
            var v193 uint8 = 1
            _ = v193
        }
        var v194 float32 = ((v2 / (0.0 + -2.25)) / (((v2 * v2) / v2) * v2))
        if len(v189) >= 2 {
            v189 = v189[1:len(v189)]
        }
        _ = v194
    }
    for fi195 := 0; fi195 < 20; fi195++ {
        for fi196 := 0; fi196 < 5; fi196++ {
            v197 := GenS22{F0: ("" + ("foo bar" + "\n\t")), F1: ("\n\t" + "")}
            _ = fi196
            _ = v197
        }
        if len(v189) > 0 {
            v189[0] = fi195
        }
        v13 += (v13 % (((v13 ^ (v13 >> ((599) & 63)))) | 1))
        _ = fi195
    }
    v2 = v2
    v3[861] = "foo bar"
    if (2147483647 != 1) {
        v2 -= (0.0 * 3.14159)
        var v198 int64 = 0
        switch -258 {
        case 4294967295:
            v13 = (v13 / (((^(v13))) | 1))
        default:
            v13 -= 0
        }
        _ = v198
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v189:", len(v189), v189)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v169
    _ = v170
    _ = v189
    _ = v191
    _ = v2
    _ = v3
    _ = v4
}
