package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 int8; F1 uint; F2 int32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint; F1 int64; F2 int16 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int8; F1 int16 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int64; F1 int16; F2 string }
type GenS12 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int16; F1 uint16 }
type GenS15 struct { F0 int8 }

func main() {
    v1 := make(map[int]int)
    var v2 int32 = -2147483648
    var v3 int16 = -356
    v4 := make(map[string]string)
    var v5 uint = 255
    for fi6 := 0; fi6 < 17; fi6++ {
        v5 = 703
        var v7 string = "\n\t"
        switch fi6 {
        case 32768:
            v8 := GenS1{F0: 52, F1: 534}
            _ = v8
        case -9223372036854775808:
            fmt.Println(fi6, v3)
        case -1:
            for fi9 := 0; fi9 < 5; fi9++ {
                if (float64((float32(((-2.25 + 1e10) - 3.14159)) + ((((3.14159 / (1e10 - 0.0)) + float32((3.14159 - (-2.25 * 0.1)))) * 0.0) - 3.14159))) > 1.5) {
                    v10 := GenS2{F0: "foo bar"}
                    _ = v10
                } else {
                    v11 := []int{4294967295}
                    var v12 uint8 = 127
                    _ = v11
                    _ = v12
                }
                v2 = v2
                _ = fi9
            }
        default:
            if (661 >= 105) {
                for fi13 := 0; fi13 < 6; fi13++ {
                    v4[(v7 + ((v7 + v7) + "hello"))] = v7
                    _ = fi13
                }
            } else {
                var v14 int32 = v2
                _ = v14
            }
        }
        _ = fi6
        _ = v7
    }
    if (9223372036854775807 > -772) {
        for fi15 := 0; fi15 < 4; fi15++ {
            var v16 int = 0
            _ = fi15
            _ = v16
        }
    } else {
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 int32 = (v2 * (v2 - v2))
            _ = fi17
            _ = v18
        }
        delete(v1, -646)
        v3 = v3
    }
    v5 += (v5 * (v5 * (v5 >> ((v5) & 63))))
    if true {
        v19 := GenS3{F0: 552, F1: 951}
        var v20 float32 = ((((3.14159 * 0.1) / -2.25) + 3.14159) * 1e10)
        _ = v19
        _ = v20
    } else {
        if (!true) {
            v21 := []int{}
            v2 -= v2
            _ = v21
        }
    }
    if false {
        var v22 int64 = 1
        fmt.Println(v22)
        switch "\n\t" {
        case "hello":
            for fi23 := 0; fi23 < 7; fi23++ {
                var v24 float64 = 0.1
                _ = fi23
                _ = v24
            }
        case "":
            v25 := GenS3{F0: 2, F1: 42}
            _ = v25
        default:
            if (0.0 != 1.5) {
                v26 := GenS3{F0: 184, F1: 952}
                fmt.Println(v22, v5, v2)
                _ = v26
            }
        }
        _ = v22
    } else {
        v27 := GenS3{F0: 18446744073709551615, F1: 65535}
        switch 32768 {
        case 9223372036854775807:
            delete(v1, 127)
        default:
            var v28 bool = (-86 < -128)
            _ = v28
        }
        var v29 int16 = v3
        _ = v27
        _ = v29
    }
    if false {
        v30 := []int{-1, 726, -2147483648, 128}
        _ = v30
    } else {
        v3 += (v3 * (^(v3)))
        switch "" {
        case "a":
            for fi31 := 0; fi31 < 4; fi31++ {
                var v32 uint = v5
                if false {
                    if false {
                        if false {
                            var v33 int16 = (v3 - 1)
                            _ = len(v1)
                            _ = v33
                        }
                        if (1.5 >= 0.1) {
                            if (float64(0.1) >= (0.1 + (3.14159 / ((float64(3.14159) * 3.14159) + 1e10)))) {
                                for fi34 := 0; fi34 < 8; fi34++ {
                                    v5 = 18446744073709551615
                                    _ = fi34
                                }
                            } else {
                                v35 := GenS4{F0: 42, F1: 1, F2: 42}
                                _ = v35
                            }
                            v36 := GenS2{F0: ""}
                            if true {
                                var v37 float32 = 0.0
                                var v38 uint = 173
                                _ = v37
                                _ = v38
                            }
                            _ = v36
                        } else {
                            v2 += v2
                            switch "a" {
                            case "foo bar":
                                if true {
                                    if (true || ((!true) || ((((1.5 + (3.14159 + 0.1)) - 0.1) == (-2.25 - float32(1e10))) || (v2 > -2147483648)))) {
                                        v32 -= 0
                                        if (1 != 4294967295) {
                                            v39 := []int{(fi31 & (fi31 & fi31)), (^(fi31))}
                                            for fi40 := 0; fi40 < 3; fi40++ {
                                                var v41 float64 = 1e10
                                                switch 128 {
                                                case 9223372036854775807:
                                                    switch "foo bar" {
                                                    case "":
                                                        v5 -= 18446744073709551615
                                                    }
                                                default:
                                                    v42 := GenS4{F0: 100, F1: v5, F2: (-(v2))}
                                                    _ = v42
                                                }
                                                v2 -= 0
                                                v32 += v32
                                                var v43 bool = false
                                                _ = fi40
                                                _ = v41
                                                _ = v43
                                            }
                                            if (!(92 != 179)) {
                                                v44 := v1[(fi31 | fi31)]
                                                v2 -= 271
                                                _ = v44
                                            }
                                            _ = v39
                                        }
                                    }
                                    fmt.Println(v3, v5)
                                    if (!(true || false)) {
                                        v45 := GenS2{F0: "a"}
                                        for fi46 := 0; fi46 < 3; fi46++ {
                                            var v47 uint16 = 866
                                            if (815 < (v47 / ((uint16(v5)) | 1))) {
                                                v48 := GenS2{F0: "foo bar"}
                                                _ = v48
                                            }
                                            v2 += (v2 + (^(v2)))
                                            v5 = v5
                                            _ = fi46
                                            _ = v47
                                        }
                                        v49 := []int{}
                                        _ = v45
                                        _ = v49
                                    } else {
                                        if (2 >= 18446744073709551615) {
                                            v50 := GenS4{F0: int8(v2), F1: v32, F2: v2}
                                            switch -1 {
                                            case 9223372036854775807:
                                                v51 := GenS2{F0: ("" + "a")}
                                                _ = v51
                                            case 1:
                                                for fi52 := 0; fi52 < 9; fi52++ {
                                                    v53 := GenS1{F0: 121, F1: 1}
                                                    switch 42 {
                                                    case 128:
                                                        v3 = v3
                                                    }
                                                    _ = fi52
                                                    _ = v53
                                                }
                                            case -9223372036854775808:
                                                if false {
                                                    v5 -= v5
                                                    v3 -= (^(v3))
                                                }
                                            }
                                            _ = v50
                                        }
                                        var v54 uint16 = 267
                                        v55 := GenS3{F0: 415, F1: 300}
                                        _ = v54
                                        _ = v55
                                    }
                                }
                            }
                            fmt.Println(fi31)
                        }
                        v56 := GenS5{F0: 9223372036854775807}
                        _ = v56
                    } else {
                        v5 += 871
                    }
                    v57 := GenS5{F0: -1}
                    if (!true) {
                        switch (-(fi31)) {
                        case 2147483647:
                            switch "" {
                            case "":
                                v58 := []int{128, fi31, fi31, (fi31 - fi31)}
                                _ = v58
                            }
                        case 32767:
                            v4["foo bar"] = ""
                        case 42:
                            if (float32(((((-2.25 * ((0.1 + 1e10) + ((0.0 + (1e10 + (1.5 / 3.14159))) + (3.14159 - 0.1)))) - float64(0.1)) * 1.5) + 3.14159)) >= (-2.25 / float32(-2.25))) {
                                switch "" {
                                case "hello":
                                    if true {
                                        var v59 string = "a"
                                        var v60 uint8 = 0
                                        for fi61 := 0; fi61 < 5; fi61++ {
                                            var v62 uint64 = 4294967295
                                            _ = fi61
                                            _ = v62
                                        }
                                        _ = v59
                                        _ = v60
                                    }
                                case "":
                                    v5 = v5
                                case "foo bar":
                                    for fi63 := 0; fi63 < 18; fi63++ {
                                        v5 -= v5
                                        _ = fi63
                                    }
                                }
                                v64 := []int{fi31}
                                _ = v64
                            }
                        }
                        var v65 uint = v32
                        _ = v65
                    }
                    _ = v57
                } else {
                    if false {
                        var v66 uint64 = 30
                        switch (^(fi31)) {
                        case -1:
                            var v67 uint8 = 2
                            _ = v67
                        case 0:
                            v2 -= (-(v2))
                        case 128:
                            fmt.Println(v66, v2)
                        }
                        v66 = 874
                        _ = v66
                    }
                    for fi68 := 0; fi68 < 17; fi68++ {
                        if false {
                            fmt.Println(fi31, v2, v5)
                            v2 += v2
                            fmt.Println(v5)
                        }
                        _ = fi68
                    }
                    if (!false) {
                        fmt.Println(v3)
                        switch (((("a" + "") + "hello") + (("foo bar" + "foo bar") + "")) + "a") {
                        case "":
                            _ = len(v1)
                        case "foo bar":
                            var v69 uint8 = 222
                            _ = v69
                        }
                        var v70 float32 = 1.5
                        _ = v70
                    }
                }
                _ = fi31
                _ = v32
            }
        case "\n\t":
            var v71 int16 = v3
            _ = v71
        case "":
            if (65535 <= 1) {
                if true {
                    if (("" == ((("a" + ((("\n\t" + ("hello" + "a")) + "") + "foo bar")) + "\n\t") + "\n\t")) && ((false && false) && false)) {
                        var v72 uint16 = 32767
                        v3 = v3
                        _ = v72
                    } else {
                        if (uint32(v5) < 0) {
                            for fi73 := 0; fi73 < 7; fi73++ {
                                if (false && false) {
                                    if true {
                                        fmt.Println(fi73, v5)
                                        for fi74 := 0; fi74 < 1; fi74++ {
                                            if (42 != 0) {
                                                if (false && true) {
                                                    var v75 float32 = -2.25
                                                    _ = v75
                                                } else {
                                                    v76 := []int{}
                                                    _ = v76
                                                }
                                                switch (("\n\t" + "\n\t") + "a") {
                                                case "foo bar":
                                                    _ = len(v1)
                                                case "hello":
                                                    v1[(fi73 * 127)] = (^(fi74))
                                                }
                                                var v77 int32 = (-(v2))
                                                v78 := GenS6{F0: (v5 - 983), F1: 869, F2: 32767}
                                                if false {
                                                    if (89 < 138) {
                                                        if true {
                                                            v5 = 4294967295
                                                        }
                                                    }
                                                    var v79 int = int(v5)
                                                    var v80 bool = (true && (!true))
                                                    var v81 int32 = v2
                                                    v77 -= (v81 - (v77 ^ v77))
                                                    if v80 {
                                                        var v82 bool = v80
                                                        _ = v82
                                                    } else {
                                                        v83 := GenS7{F0: 0}
                                                        if (42 < uint16(v5)) {
                                                            fmt.Println(v2, fi74, v2)
                                                            _ = len(v1)
                                                            v84 := []int{fi73, fi73, v79}
                                                            _ = v84
                                                        } else {
                                                            if true {
                                                                v85 := GenS1{F0: 65535, F1: 2}
                                                                switch fi73 {
                                                                case -2147483648:
                                                                    var v86 uint8 = 158
                                                                    _ = v86
                                                                case 0:
                                                                    if v80 {
                                                                        if v80 {
                                                                            var v87 string = "foo bar"
                                                                            _ = v87
                                                                        } else {
                                                                            v5 = (v5 << (((v5 >> ((v5) & 63))) & 63))
                                                                            v1[42] = v79
                                                                        }
                                                                        v5 = 42
                                                                    } else {
                                                                        v3 = -1
                                                                        if v80 {
                                                                            v88 := []int{(^(v79)), v79, (fi74 | 2147483647), (^(v79))}
                                                                            var v89 uint64 = 624
                                                                            _ = v88
                                                                            _ = v89
                                                                        } else {
                                                                            v81 = v2
                                                                            v90 := GenS4{F0: 42, F1: (v5 >> ((18446744073709551615) & 63)), F2: v77}
                                                                            var v91 uint8 = 2
                                                                            var v92 uint8 = 1
                                                                            delete(v4, "\n\t")
                                                                            v77 -= (-(v2))
                                                                            _ = v90
                                                                            _ = v91
                                                                            _ = v92
                                                                        }
                                                                    }
                                                                case -728:
                                                                    fmt.Println(v77, v80)
                                                                default:
                                                                    var v93 bool = (v80 && (((64 > -943) || (v3 > (v3 / (((v3 >> ((319) & 63))) | 1)))) || (32767 < 69)))
                                                                    _ = v93
                                                                }
                                                                _ = v85
                                                            }
                                                            v80 = (v80 || (3.14159 == (1.5 / float64(3.14159))))
                                                        }
                                                        switch 815 {
                                                        case 42:
                                                            if (v80 && v80) {
                                                                fmt.Println(v81, v3, v5)
                                                                var v94 uint16 = 32767
                                                                _ = v94
                                                            } else {
                                                                var v95 int32 = v2
                                                                v81 = v81
                                                                var v96 int8 = 127
                                                                _ = v95
                                                                _ = v96
                                                            }
                                                        default:
                                                            v97 := GenS6{F0: (v5 * v5), F1: -9223372036854775808, F2: (v3 - (v3 << (((v5 + (v5 % (((v5 - 689)) | 1)))) & 63)))}
                                                            _ = v97
                                                        }
                                                        _ = v83
                                                    }
                                                    if (!((v3 * (v3 | 0)) <= (v3 - 42))) {
                                                        var v98 uint64 = 260
                                                        var v99 bool = (0 >= 815)
                                                        v81 = (v81 + (v81 & 42))
                                                        _ = v98
                                                        _ = v99
                                                    }
                                                    var v100 uint16 = 255
                                                    v101 := []int{}
                                                    v3 = (v3 ^ (v3 / (((v3 + v3)) | 1)))
                                                    _ = v100
                                                    _ = v101
                                                    _ = v79
                                                    _ = v80
                                                    _ = v81
                                                }
                                                _ = v77
                                                _ = v78
                                            } else {
                                                if false {
                                                    if (!(true && ((float64(0.1) * 3.14159) == (0.1 + (float64(float32((3.14159 + 3.14159))) - 3.14159))))) {
                                                        switch (fi74 ^ 9223372036854775807) {
                                                        case -2147483648:
                                                            v4["a"] = "\n\t"
                                                        case 261:
                                                            if (-1 > int64(v5)) {
                                                                if true {
                                                                    fmt.Println(v3, fi73, v3)
                                                                    v102 := GenS8{F0: -1, F1: -32768}
                                                                    _ = v102
                                                                } else {
                                                                    fmt.Println(v3)
                                                                    v103 := []int{787, (fi73 << ((65535) & 63)), fi73, (fi74 * (fi73 ^ fi74))}
                                                                    switch 128 {
                                                                    case 2147483647:
                                                                        v5 += (v5 * (v5 - v5))
                                                                    case 128:
                                                                        if (1 <= 127) {
                                                                            v104 := GenS2{F0: "foo bar"}
                                                                            if false {
                                                                                v5 -= v5
                                                                            }
                                                                            v2 = (v2 + v2)
                                                                            _ = v104
                                                                        } else {
                                                                            var v105 int = fi73
                                                                            _ = len(v4)
                                                                            _ = v105
                                                                        }
                                                                    }
                                                                    _ = v103
                                                                }
                                                                delete(v4, ("\n\t" + (("a" + "hello") + "")))
                                                            }
                                                        case 4294967295:
                                                            if false {
                                                                var v106 int32 = (v2 / ((v2) | 1))
                                                                if (!((-938 <= -916) || ((v3 & v3) != v3))) {
                                                                    v3 = (v3 % (((-(v3))) | 1))
                                                                    v106 = v2
                                                                    v2 = -1
                                                                    v107 := v1[(fi74 ^ 32768)]
                                                                    fmt.Println(fi73)
                                                                    _ = v107
                                                                }
                                                                var v108 uint16 = 32767
                                                                var v109 uint64 = 1
                                                                _ = v106
                                                                _ = v108
                                                                _ = v109
                                                            }
                                                        }
                                                        v5 -= 1
                                                        var v110 uint16 = 42
                                                        _ = v110
                                                    }
                                                    v111 := []int{(fi73 / ((fi73) | 1)), 128, (^(fi73)), (^(fi74))}
                                                    var v112 uint = (v5 + (v5 << (((^(v5))) & 63)))
                                                    _ = v111
                                                    _ = v112
                                                }
                                            }
                                            fmt.Println(fi73, v5)
                                            fmt.Println(v3, v2, v2)
                                            var v113 int = fi74
                                            v5 += 18446744073709551615
                                            _ = fi74
                                            _ = v113
                                        }
                                        var v114 bool = (-9223372036854775808 >= 1)
                                        _ = v114
                                    }
                                    for fi115 := 0; fi115 < 9; fi115++ {
                                        v2 += v2
                                        var v116 int = fi73
                                        var v117 uint8 = 146
                                        _ = fi115
                                        _ = v116
                                        _ = v117
                                    }
                                    var v118 int64 = int64(v2)
                                    _ = v118
                                }
                                _ = fi73
                            }
                            for fi119 := 0; fi119 < 2; fi119++ {
                                delete(v4, (("\n\t" + "\n\t") + "foo bar"))
                                if (4294967295 != 131) {
                                    v120 := GenS7{F0: 65535}
                                    var v121 int8 = 18
                                    for fi122 := 0; fi122 < 7; fi122++ {
                                        v5 -= 2
                                        var v123 float64 = 1.5
                                        _ = fi122
                                        _ = v123
                                    }
                                    _ = v120
                                    _ = v121
                                } else {
                                    v124 := GenS9{F0: "hello"}
                                    switch (fi119 >> ((65535) & 63)) {
                                    case 0:
                                        v2 = v2
                                    }
                                    v5 = v5
                                    _ = v124
                                }
                                var v125 uint64 = uint64(v5)
                                _ = fi119
                                _ = v125
                            }
                        } else {
                            var v126 bool = ((!(v5 < (^(v5)))) && true)
                            _ = v126
                        }
                        fmt.Println(v3, v5, v2)
                        fmt.Println(v5, v3)
                    }
                    delete(v4, "hello")
                    if false {
                        v5 = v5
                        if false {
                            for fi127 := 0; fi127 < 3; fi127++ {
                                switch "foo bar" {
                                case "hello":
                                    fmt.Println(v2)
                                default:
                                    for fi128 := 0; fi128 < 16; fi128++ {
                                        var v129 float64 = 0.0
                                        v130 := GenS2{F0: "a"}
                                        v131 := v4["hello"]
                                        if (277 >= 822) {
                                            var v132 float32 = 3.14159
                                            var v133 uint = 4294967295
                                            _ = v132
                                            _ = v133
                                        } else {
                                            v134 := GenS5{F0: -1}
                                            fmt.Println(v129, fi128)
                                            _ = v134
                                        }
                                        _ = fi128
                                        _ = v129
                                        _ = v130
                                        _ = v131
                                    }
                                }
                                for fi135 := 0; fi135 < 3; fi135++ {
                                    fmt.Println(v3)
                                    _ = fi135
                                }
                                _ = fi127
                            }
                            for fi136 := 0; fi136 < 11; fi136++ {
                                v137 := GenS1{F0: 65535, F1: 2147483647}
                                _ = fi136
                                _ = v137
                            }
                        } else {
                            fmt.Println(v2)
                            v138 := v1[4294967295]
                            _ = v138
                        }
                    }
                }
                var v139 uint16 = 64
                if ((!true) && true) {
                    var v140 bool = (true || (true && ((int8(v139) == -57) || (v3 <= v3))))
                    if v140 {
                        switch "\n\t" {
                        case "hello":
                            for fi141 := 0; fi141 < 1; fi141++ {
                                fmt.Println(fi141, v3, fi141)
                                if (!(fi141 != (fi141 - fi141))) {
                                    fmt.Println(v2, v5)
                                    var v142 string = "hello"
                                    _ = v142
                                }
                                if v140 {
                                    v143 := GenS10{F0: "hello"}
                                    _ = v143
                                } else {
                                    if v140 {
                                        v139 = v139
                                        for fi144 := 0; fi144 < 12; fi144++ {
                                            v140 = ((-(v2)) >= (v2 * -848))
                                            v145 := []int{(fi141 ^ fi141), fi141, -605}
                                            _ = fi144
                                            _ = v145
                                        }
                                    }
                                    v146 := GenS4{F0: 1, F1: (^(v5)), F2: (v2 % ((v2) | 1))}
                                    var v147 uint32 = 2147483648
                                    _ = v146
                                    _ = v147
                                }
                                _ = fi141
                            }
                        case "foo bar":
                            var v148 int64 = 196
                            _ = v148
                        default:
                            v149 := GenS6{F0: (v5 - 752), F1: -9223372036854775808, F2: v3}
                            _ = v149
                        }
                    } else {
                        var v150 uint64 = 42
                        var v151 uint32 = 379
                        _ = v150
                        _ = v151
                    }
                    v152 := []int{}
                    _ = v140
                    _ = v152
                }
                _ = v139
            } else {
                v4["foo bar"] = "a"
                v153 := GenS5{F0: -631}
                fmt.Println(v2, v5)
                _ = v153
            }
        }
        for fi154 := 0; fi154 < 8; fi154++ {
            switch (^(fi154)) {
            case -1:
                var v155 int = fi154
                _ = v155
            case 42:
                switch 494 {
                case 2147483647:
                    var v156 int8 = -1
                    _ = v156
                case -288:
                    v3 += 0
                case 32767:
                    if false {
                        v2 += v2
                    }
                }
            case 127:
                var v157 int32 = (v2 % ((1) | 1))
                _ = v157
            default:
                v158 := GenS1{F0: 65535, F1: 2}
                _ = v158
            }
            _ = fi154
        }
    }
    if ((v5 & v5) > 1) {
        switch 32767 {
        case -626:
            fmt.Println(v3, v2)
        case 955:
            for fi159 := 0; fi159 < 19; fi159++ {
                for fi160 := 0; fi160 < 3; fi160++ {
                    switch "\n\t" {
                    case "":
                        v161 := GenS11{F0: 0, F1: -1, F2: "foo bar"}
                        _ = v161
                    case "foo bar":
                        v162 := GenS12{F0: (v5 % ((v5) | 1)), F1: 65535, F2: 58}
                        if (0 >= 4294967295) {
                            if false {
                                v5 += (v5 << ((v5) & 63))
                            }
                        }
                        var v163 int64 = 42
                        _ = v162
                        _ = v163
                    }
                    fmt.Println(v5, fi160, v5)
                    var v164 uint16 = 32767
                    _ = fi160
                    _ = v164
                }
                _ = fi159
            }
        case 0:
            v3 = v3
        }
        v5 = (^(v5))
        for fi165 := 0; fi165 < 7; fi165++ {
            for fi166 := 0; fi166 < 6; fi166++ {
                v167 := []int{}
                if false {
                    var v168 int8 = 80
                    v169 := GenS13{F0: 483}
                    v170 := GenS14{F0: (v3 << (((^(v5))) & 63)), F1: 2}
                    _ = v168
                    _ = v169
                    _ = v170
                }
                _ = fi166
                _ = v167
            }
            v171 := GenS15{F0: -66}
            _ = fi165
            _ = v171
        }
    } else {
        var v172 int16 = v3
        _ = v172
    }
    var v173 uint8 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v173:", v173)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v173
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
