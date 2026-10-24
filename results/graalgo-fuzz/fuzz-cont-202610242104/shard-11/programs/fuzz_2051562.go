package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 int64; F1 int16; F2 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint32; F1 int64; F2 int8 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int; F1 int32 }
type GenS11 struct { F0 int64; F1 uint64; F2 int8 }
type GenS12 struct { F0 uint32 }
func genFn0() int32 {
    var v7 string = "foo bar"
    fmt.Println(v7)
    _ = v7
    return 2147483647
}
func genFn1(p0 int8) uint {
    v8 := GenS1{F0: 0}
    p0 = -1
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := []int{(fi9 & fi9), 127, 32768, (fi9 + (fi9 & fi9))}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v8
    return 65535
}

func main() {
    var v1 string = (("foo bar" + "hello") + "\n\t")
    var v2 int16 = -1
    var v3 uint32 = 83
    var v4 uint16 = 507
    v5 := make(map[uint32]int)
    var v6 uint32 = (v3 & 2147483647)
    var v11 uint16 = (v4 / (((v4 << ((0) & 63))) | 1))
    for fi12 := 0; fi12 < 20; fi12++ {
        if (!(false || true)) {
            v13 := []int{159, 2147483647}
            if true {
                if len(v13) > 0 {
                    v13[0] = (-(fi12))
                }
            } else {
                v14 := GenS2{F0: v1, F1: 2}
                for fi15 := 0; fi15 < 6; fi15++ {
                    switch fi15 {
                    case 32768:
                        switch "hello" {
                        case "":
                            fmt.Println(fi15, v4, v4)
                            v2 = 1000
                            v6 += (v3 ^ (v6 >> ((2) & 63)))
                        default:
                            v13 = append(v13, (^(fi12)))
                        }
                    case 399:
                        var v16 float64 = 0.0
                        _ = v16
                    default:
                        if true {
                            var v17 uint8 = 127
                            switch ("\n\t" + "foo bar") {
                            case "":
                                v18 := v5[v6]
                                _ = v18
                            }
                            var v19 string = v1
                            _ = len(v13)
                            _ = v17
                            _ = v19
                        }
                    }
                    var v20 uint = 0
                    if (!((-2.25 - (((-2.25 - 3.14159) - 3.14159) / (1e10 * float32(-2.25)))) > 1.5)) {
                        switch (("hello" + (("\n\t" + v1) + v1)) + v1) {
                        case "\n\t":
                            var v21 int16 = v2
                            _ = v21
                        case "":
                            var v22 uint = v20
                            _ = v22
                        case "hello":
                            var v23 int = fi12
                            _ = v23
                        default:
                            v24 := v5[v6]
                            _ = v24
                        }
                        fmt.Println(v11, v2)
                        v25 := GenS3{F0: (v4 ^ v4), F1: 0}
                        _ = v25
                    }
                    var v26 int64 = 1
                    _ = fi15
                    _ = v20
                    _ = v26
                }
                v5[uint32(v4)] = int(v4)
                _ = v14
            }
            _ = v13
        } else {
            for fi27 := 0; fi27 < 20; fi27++ {
                if false {
                    var v28 uint = 0
                    v29 := []int{(^(fi12)), fi12, -755}
                    v29 = append(v29, (fi12 * 0))
                    v30 := GenS3{F0: v11, F1: v4}
                    _ = v28
                    _ = v29
                    _ = v30
                } else {
                    var v31 uint32 = v6
                    v32 := GenS4{F0: 42, F1: (v2 - int16(v6)), F2: v1}
                    _ = v31
                    _ = v32
                }
                _ = fi27
            }
            v33 := v5[(v6 | 2147483648)]
            _ = v33
        }
        for fi34 := 0; fi34 < 9; fi34++ {
            v35 := []int{-1}
            switch fi12 {
            case 127:
                var v36 float32 = (3.14159 / -2.25)
                _ = v36
            case -556:
                switch ((v1 + v1) + v1) {
                case "hello":
                    if false {
                        v1 = v1
                        switch (v1 + v1) {
                        case "a":
                            if (559 != -9223372036854775808) {
                                v11 -= (v11 & 2)
                                var v37 uint8 = 200
                                v38 := GenS5{F0: true}
                                v5[499] = (fi34 + -666)
                                _ = v37
                                _ = v38
                            }
                        case "hello":
                            if false {
                                if false {
                                    var v39 uint32 = 413
                                    var v40 int64 = 0
                                    var v41 string = v1
                                    _ = v39
                                    _ = v40
                                    _ = v41
                                } else {
                                    var v42 float32 = 1.5
                                    if (v3 > 65535) {
                                        if (-2.25 <= ((1.5 + 1.5) - float64((1.5 * ((((float32(0.1) - v42) * -2.25) - v42) + v42))))) {
                                            switch ((v1 + v1) + v1) {
                                            case "\n\t":
                                                fmt.Println(fi12, v3, v3)
                                            case "":
                                                if ((-256 >= 9223372036854775807) && (9223372036854775807 == 0)) {
                                                    var v43 float64 = ((0.0 - float64(1e10)) / ((1.5 + (0.1 + (1.5 / 3.14159))) + 0.1))
                                                    var v44 uint = 42
                                                    if (!false) {
                                                        v43 = (0.0 + v43)
                                                    } else {
                                                        var v45 uint16 = uint16(v6)
                                                        v46 := v5[211]
                                                        _ = v45
                                                        _ = v46
                                                    }
                                                    _ = v43
                                                    _ = v44
                                                }
                                            case "foo bar":
                                                v6 -= uint32(v4)
                                            default:
                                                var v47 uint32 = (v3 << ((911) & 63))
                                                v48 := GenS5{F0: (!true)}
                                                var v49 int32 = -2147483648
                                                _ = v47
                                                _ = v48
                                                _ = v49
                                            }
                                        }
                                    } else {
                                        v50 := GenS5{F0: false}
                                        _ = v50
                                    }
                                    _ = v42
                                }
                                delete(v5, v6)
                            }
                        case "":
                            v51 := GenS6{F0: -128}
                            _ = v51
                        default:
                            var v52 float64 = 0.0
                            _ = v52
                        }
                        var v53 int64 = -701
                        v1 = "foo bar"
                        _ = v53
                    }
                case "a":
                    if (0 >= 9223372036854775807) {
                        var v54 uint16 = v4
                        if (((^(v2)) == v2) && (false && false)) {
                            var v55 int64 = -1
                            _ = v55
                        } else {
                            fmt.Println(v1)
                            v1 = (v1 + v1)
                            _ = len(v35)
                            v56 := v5[uint32(v11)]
                            var v57 bool = true
                            switch (^(v56)) {
                            case -412:
                                var v58 int32 = 2147483647
                                _ = v58
                            }
                            _ = v56
                            _ = v57
                        }
                        var v59 uint32 = v3
                        _ = v54
                        _ = v59
                    }
                case "":
                    v6 -= (v6 >> ((4294967295) & 63))
                default:
                    fmt.Println(fi12)
                }
            case 9223372036854775807:
                if false {
                    fmt.Println(v1, v1, v2)
                    v60 := v5[v3]
                    _ = v60
                } else {
                    v6 = (v3 | (v6 | (^(v3))))
                    v4 = (v4 % ((32767) | 1))
                }
            default:
                if false {
                    fmt.Println(fi34, fi34)
                    var v61 int16 = v2
                    _ = v61
                } else {
                    v62 := GenS7{F0: v1}
                    _ = v62
                }
            }
            _ = fi34
            _ = v35
        }
        v1 = v1
        _ = fi12
    }
    fmt.Println(v4, v2, v4)
    var v63 float64 = (-2.25 / float64(0.1))
    var v64 float64 = v63
    fmt.Println(v11, v3, v63)
    v65 := GenS1{F0: (v4 + (^(v4)))}
    v2 = v2
    for fi66 := 0; fi66 < 18; fi66++ {
        if false {
            _ = len(v5)
        } else {
            v67 := []int{fi66, 32768}
            _ = v67
        }
        _ = fi66
    }
    v68 := GenS2{F0: "a", F1: 2}
    if false {
        var v69 uint = 2
        v70 := []int{-966, 9223372036854775807, 32768}
        _ = v69
        _ = v70
    }
    switch v1 {
    case "":
        for fi71 := 0; fi71 < 15; fi71++ {
            v72 := GenS2{F0: "a", F1: 478}
            for fi73 := 0; fi73 < 10; fi73++ {
                fmt.Println(v63)
                _ = fi73
            }
            _ = fi71
            _ = v72
        }
    case "\n\t":
        var v74 uint8 = 128
        _ = v74
    default:
        if ((v4 < (v4 + 42)) || (v63 != float64(-2.25))) {
            var v75 float64 = v63
            switch "foo bar" {
            case "foo bar":
                if ((v63 + 1.5) < float64(((float32(v75) + (1.5 + (0.0 + -2.25))) - (1e10 * 1e10)))) {
                    var v76 float32 = (0.0 + 1e10)
                    switch -9223372036854775808 {
                    case 128:
                        if (float64((v76 / 0.1)) >= (1e10 + (-2.25 / v63))) {
                            v77 := GenS3{F0: v11, F1: v11}
                            v78 := GenS6{F0: 0}
                            if (v2 != 33) {
                                var v79 string = v1
                                _ = v79
                            } else {
                                var v80 uint64 = 42
                                if ((v80 != 42) && false) {
                                    v81 := GenS4{F0: 42, F1: 42, F2: ("hello" + v1)}
                                    var v82 float64 = float64(float32(float64(-2.25)))
                                    if (958 > 737) {
                                        v83 := GenS3{F0: (v4 * uint16(v6)), F1: v11}
                                        var v84 int32 = int32(v80)
                                        var v85 float64 = float64(1e10)
                                        _ = v83
                                        _ = v84
                                        _ = v85
                                    } else {
                                        if true {
                                            v86 := []int{42, 9223372036854775807}
                                            var v87 uint16 = (v4 / ((1) | 1))
                                            _ = v86
                                            _ = v87
                                        }
                                    }
                                    _ = v81
                                    _ = v82
                                } else {
                                    var v88 int64 = 614
                                    _ = v88
                                }
                                _ = v80
                            }
                            _ = v77
                            _ = v78
                        } else {
                            var v89 float32 = ((1.5 - ((0.1 + (((float32((v64 - v63)) + (((v76 + (((0.1 / (v76 + (-2.25 - float32((v64 / v64))))) * 1e10) - v76)) - (v76 / v76)) * (((v76 - (1e10 - 3.14159)) - v76) - (-2.25 / (0.0 - (float32(((v63 / v75) - (1e10 / v63))) - (v76 - float32(float64((v76 - v76)))))))))) + float32(3.14159)) - 0.1)) * v76)) * v76)
                            for fi90 := 0; fi90 < 2; fi90++ {
                                fmt.Println(v63)
                                switch "hello" {
                                case "\n\t":
                                    var v91 uint8 = 129
                                    _ = v91
                                case "":
                                    fmt.Println(v4, v4, v2)
                                case "a":
                                    for fi92 := 0; fi92 < 2; fi92++ {
                                        if (255 > genFn1(0)) {
                                            var v93 uint16 = 32767
                                            fmt.Println(v4)
                                            delete(v5, (v3 / (((v6 ^ 0)) | 1)))
                                            var v94 uint32 = (^(v3))
                                            var v95 float32 = -2.25
                                            var v96 uint32 = (^(v3))
                                            if (!(!true)) {
                                                fmt.Println(v3, v76)
                                                v97 := GenS4{F0: 42, F1: v2, F2: (v1 + v1)}
                                                fmt.Println(v11, v96)
                                                _ = v97
                                            } else {
                                                v98 := GenS2{F0: (("a" + v1) + "\n\t"), F1: 4294967295}
                                                _ = v98
                                            }
                                            _ = v93
                                            _ = v94
                                            _ = v95
                                            _ = v96
                                        } else {
                                            var v99 uint = 1
                                            _ = v99
                                        }
                                        if false {
                                            var v100 uint64 = 0
                                            switch "" {
                                            case "foo bar":
                                                v101 := GenS2{F0: (v1 + ""), F1: v100}
                                                _ = v101
                                            case "hello":
                                                v102 := GenS1{F0: v11}
                                                _ = v102
                                            default:
                                                v103 := GenS8{F0: 2147483648, F1: 42, F2: 127}
                                                _ = v103
                                            }
                                            _ = v100
                                        }
                                        var v104 uint16 = 536
                                        _ = fi92
                                        _ = v104
                                    }
                                }
                                if (!true) {
                                    if (!false) {
                                        v2 = (v2 * v2)
                                    } else {
                                        var v105 uint16 = v4
                                        _ = v105
                                    }
                                    if true {
                                        if false {
                                            if ((v6 / ((v3) | 1)) > 0) {
                                                fmt.Println(v89, v3, v76)
                                            } else {
                                                v75 -= (1e10 / float64(v76))
                                                v63 = -2.25
                                                v106 := []int{fi90, (fi90 >> ((genFn1(7)) & 63)), (^(fi90)), fi90}
                                                _ = v106
                                            }
                                        } else {
                                            v107 := GenS8{F0: (v3 << ((748) & 63)), F1: int64(v2), F2: 22}
                                            fmt.Println(v6)
                                            _ = v107
                                        }
                                        var v108 bool = false
                                        switch v1 {
                                        case "a":
                                            v109 := GenS9{F0: v6}
                                            _ = v109
                                        case "hello":
                                            switch "" {
                                            case "a":
                                                fmt.Println(v64, v1, v75)
                                            case "\n\t":
                                                switch (fi90 + 128) {
                                                case -1:
                                                    var v110 int32 = 0
                                                    _ = v110
                                                case 217:
                                                    fmt.Println(v64)
                                                case 0:
                                                    v6 += (v3 << ((735) & 63))
                                                default:
                                                    v111 := GenS10{F0: (fi90 / ((fi90) | 1)), F1: -1}
                                                    _ = v111
                                                }
                                            case "foo bar":
                                                var v112 int8 = -1
                                                _ = v112
                                            }
                                        case "foo bar":
                                            for fi113 := 0; fi113 < 4; fi113++ {
                                                var v114 float32 = -2.25
                                                _ = fi113
                                                _ = v114
                                            }
                                        }
                                        _ = v108
                                    } else {
                                        switch fi90 {
                                        case 2147483647:
                                            var v115 int = (fi90 / ((fi90) | 1))
                                            _ = v115
                                        case -321:
                                            v11 -= 32767
                                        case 801:
                                            for fi116 := 0; fi116 < 18; fi116++ {
                                                v4 = v4
                                                var v117 float64 = float64((v89 - v89))
                                                var v118 string = ("hello" + (v1 + v1))
                                                _ = fi116
                                                _ = v117
                                                _ = v118
                                            }
                                        default:
                                            v63 += 0.0
                                        }
                                        switch ((("" + "") + ("a" + v1)) + v1) {
                                        case "":
                                            if true {
                                                v119 := []int{fi90, -9223372036854775808, fi90}
                                                _ = v119
                                            }
                                        default:
                                            v75 += (v75 + 3.14159)
                                        }
                                        v120 := []int{}
                                        _ = v120
                                    }
                                }
                                _ = fi90
                            }
                            fmt.Println(v76, v75, v1)
                            _ = v89
                        }
                    }
                    v121 := GenS4{F0: -9223372036854775808, F1: (v2 | v2), F2: "hello"}
                    _ = v121
                    _ = v76
                } else {
                    if (v11 == (v11 ^ (^(v11)))) {
                        v4 -= 32767
                    } else {
                        for fi122 := 0; fi122 < 1; fi122++ {
                            v123 := []int{}
                            if (3.14159 >= (v63 + v63)) {
                                var v124 bool = false
                                var v125 uint8 = 128
                                _ = v124
                                _ = v125
                            } else {
                                v126 := GenS8{F0: v3, F1: 9223372036854775807, F2: 0}
                                _ = v126
                            }
                            v3 -= v3
                            _ = fi122
                            _ = v123
                        }
                        switch "a" {
                        case "":
                            v4 = 2
                        case "foo bar":
                            v64 = (0.1 * (v63 / v63))
                        case "\n\t":
                            if (genFn0() < -2147483648) {
                                v64 = (float64(float32(v75)) - 1.5)
                                v127 := []int{-123, 807}
                                var v128 int16 = 42
                                _ = v127
                                _ = v128
                            }
                        default:
                            var v129 float32 = 0.1
                            _ = v129
                        }
                    }
                }
            case "":
                if true {
                    v6 += v3
                    v3 = (v6 - v6)
                }
            case "a":
                if false {
                    switch -999 {
                    case -539:
                        v75 = -2.25
                    case 1:
                        if true {
                            for fi130 := 0; fi130 < 16; fi130++ {
                                v131 := []int{32768, 4294967295, int(v3), (fi130 - fi130)}
                                _ = fi130
                                _ = v131
                            }
                            for fi132 := 0; fi132 < 7; fi132++ {
                                for fi133 := 0; fi133 < 3; fi133++ {
                                    var v134 int8 = -1
                                    v135 := GenS9{F0: v6}
                                    _ = fi133
                                    _ = v134
                                    _ = v135
                                }
                                _ = fi132
                            }
                        }
                    case 128:
                        v64 = (v64 - (v75 * v64))
                    default:
                        switch 0 {
                        case -495:
                            if (false || (!(0 < 776))) {
                                v11 += 255
                                switch ("a" + "\n\t") {
                                case "":
                                    switch 9223372036854775807 {
                                    case 42:
                                        var v136 uint16 = v11
                                        _ = v136
                                    case 9223372036854775807:
                                        v3 = 688
                                    default:
                                        v137 := []int{9223372036854775807}
                                        _ = v137
                                    }
                                default:
                                    v64 = ((0.0 + (((((v75 - 0.0) * 3.14159) * (3.14159 * 1e10)) - (float64(float32((((v75 - v63) / (-2.25 - v75)) * v64))) * (v75 - (v64 + v64)))) + v75)) + -2.25)
                                }
                                for fi138 := 0; fi138 < 13; fi138++ {
                                    v1 = v1
                                    if (v6 == v6) {
                                        var v139 int = fi138
                                        var v140 uint64 = 4294967295
                                        _ = v139
                                        _ = v140
                                    }
                                    v141 := v5[(v6 - (v3 / ((v3) | 1)))]
                                    _ = fi138
                                    _ = v141
                                }
                            }
                        }
                    }
                    for fi142 := 0; fi142 < 9; fi142++ {
                        switch (fi142 - (fi142 + (fi142 + 1))) {
                        case 42:
                            var v143 bool = (!(-9223372036854775808 > -907))
                            _ = v143
                        default:
                            fmt.Println(v6)
                        }
                        _ = fi142
                    }
                    v144 := GenS3{F0: 0, F1: v11}
                    _ = v144
                } else {
                    switch "hello" {
                    case "a":
                        v145 := []int{32768}
                        _ = v145
                    }
                    if (-2.25 < 1e10) {
                        var v146 int64 = -120
                        for fi147 := 0; fi147 < 4; fi147++ {
                            for fi148 := 0; fi148 < 7; fi148++ {
                                v149 := GenS10{F0: (fi147 % ((fi148) | 1)), F1: 460}
                                v150 := GenS1{F0: 1}
                                _ = fi148
                                _ = v149
                                _ = v150
                            }
                            for fi151 := 0; fi151 < 14; fi151++ {
                                var v152 uint = uint(v4)
                                switch 1 {
                                case -379:
                                    switch (v1 + (v1 + ((((v1 + ("\n\t" + "")) + v1) + v1) + (("hello" + ("hello" + "foo bar")) + (v1 + "\n\t"))))) {
                                    case "":
                                        var v153 int64 = v146
                                        _ = v153
                                    case "hello":
                                        if true {
                                            if (!(v146 > v146)) {
                                                if false {
                                                    v4 = 551
                                                    var v154 uint32 = 2
                                                    switch fi147 {
                                                    case 9223372036854775807:
                                                        v155 := []int{fi147}
                                                        _ = v155
                                                    case 2147483647:
                                                        var v156 int = (fi147 >> (((^(v152))) & 63))
                                                        _ = v156
                                                    case 327:
                                                        v1 = v1
                                                    }
                                                    _ = v154
                                                } else {
                                                    fmt.Println(v64)
                                                    v157 := []int{fi147, fi151}
                                                    _ = v157
                                                }
                                                v152 -= 661
                                            }
                                        } else {
                                            var v158 uint8 = 128
                                            v159 := GenS4{F0: 1, F1: -1, F2: v1}
                                            _ = v158
                                            _ = v159
                                        }
                                    default:
                                        fmt.Println(v63, v1)
                                    }
                                default:
                                    v4 = uint16(v152)
                                }
                                _ = fi151
                                _ = v152
                            }
                            v11 += (^(v4))
                            _ = fi147
                        }
                        _ = v146
                    }
                    if false {
                        switch v1 {
                        case "\n\t":
                            var v160 float32 = 1.5
                            _ = v160
                        case "hello":
                            for fi161 := 0; fi161 < 2; fi161++ {
                                if false {
                                    for fi162 := 0; fi162 < 13; fi162++ {
                                        v63 = 1.5
                                        v1 = v1
                                        switch (v1 + ("\n\t" + v1)) {
                                        case "foo bar":
                                            v163 := v5[v6]
                                            _ = v163
                                        }
                                        _ = fi162
                                    }
                                    v11 -= v4
                                } else {
                                    fmt.Println(v75, v6)
                                    v164 := GenS10{F0: fi161, F1: -1}
                                    _ = v164
                                }
                                v165 := GenS2{F0: ((v1 + (v1 + v1)) + ((v1 + ((v1 + v1) + v1)) + "\n\t")), F1: 618}
                                switch v1 {
                                case "hello":
                                    var v166 float32 = ((3.14159 - float32(float64((-2.25 / 1.0)))) / 0.1)
                                    _ = v166
                                }
                                _ = fi161
                                _ = v165
                            }
                        default:
                            for fi167 := 0; fi167 < 7; fi167++ {
                                v3 = (v3 & v3)
                                _ = fi167
                            }
                        }
                    }
                }
            default:
                v168 := GenS11{F0: 0, F1: 2, F2: 1}
                _ = v168
            }
            fmt.Println(v1, v64)
            _ = v75
        }
    }
    fmt.Println(v3, v3)
    for fi169 := 0; fi169 < 17; fi169++ {
        var v170 uint = 2
        v171 := GenS12{F0: v6}
        _ = fi169
        _ = v170
        _ = v171
    }
    v4 = (v4 - 640)
    if ((v6 >> ((4294967295) & 63)) > v3) {
        fmt.Println(v3, v3, v1)
    } else {
        v172 := GenS1{F0: (^(v4))}
        _ = v172
    }
    v11 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v64
    _ = v65
    _ = v68
}
