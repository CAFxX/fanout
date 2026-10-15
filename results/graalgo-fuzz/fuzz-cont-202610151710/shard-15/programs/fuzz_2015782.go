package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint }
type GenS2 struct { F0 uint32; F1 uint; F2 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 uint8; F1 uint32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS10 struct { F0 int64; F1 uint; F2 int }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int64; F1 uint16; F2 int }
func genFn0(p0 bool) uint32 {
    v5 := []int{883}
    var v6 int16 = 1
    p0 = p0
    _ = len(v5)
    if (-2147483648 != -1) {
        v7 := GenS1{F0: -9223372036854775808, F1: 4294967295, F2: 42}
        _ = v7
    }
    _ = p0
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 int8, p1 uint32) uint {
    for fi8 := 0; fi8 < 1; fi8++ {
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 int8 = (p0 >> ((42) & 63))
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        switch "\n\t" {
        case "hello":
            if true {
                switch "hello" {
                case "foo bar":
                    var v12 uint8 = 7
                    _ = v12
                default:
                    v13 := []int{fi11, (fi11 - fi11), int(p0), fi11}
                    _ = v13
                }
                var v14 uint = 42
                for fi15 := 0; fi15 < 15; fi15++ {
                    v14 -= 540
                    if true {
                        if false {
                            var v16 uint32 = p1
                            v17 := GenS2{F0: v16, F1: 18446744073709551615, F2: 42}
                            var v18 int8 = (p0 >> ((uint(v16)) & 63))
                            v19 := []int{(fi11 / ((-881) | 1))}
                            _ = v16
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                        fmt.Println(p1, fi15, fi15)
                        var v20 int64 = 0
                        _ = v20
                    }
                    p0 = (^(p0))
                    v14 += (^(v14))
                    fmt.Println(p0, fi15, p1)
                    fmt.Println(fi11, v14)
                    _ = fi15
                }
                _ = v14
            }
        }
        for fi21 := 0; fi21 < 2; fi21++ {
            v22 := []int{89}
            p1 += p1
            _ = len(v22)
            if true {
                var v23 uint = 1
                switch (((("foo bar" + "foo bar") + ("\n\t" + (("\n\t" + ("foo bar" + ("" + (("a" + "hello") + "foo bar")))) + "hello"))) + ("hello" + "hello")) + "a") {
                case "foo bar":
                    fmt.Println(p0)
                }
                if (42 < 0) {
                    var v24 string = "foo bar"
                    var v25 float64 = ((-2.25 * (0.0 - ((0.1 / 1e10) - 1.5))) + ((3.14159 - 1e10) / 3.14159))
                    if len(v22) > 0 {
                        v22[0] = fi21
                    }
                    _ = v24
                    _ = v25
                }
                p0 += (-(p0))
                _ = v23
            } else {
                if (!(1 < 42)) {
                    if (true && true) {
                        if false {
                            var v26 string = (("hello" + "\n\t") + ("hello" + "foo bar"))
                            _ = v26
                        }
                        switch "" {
                        case "":
                            var v27 string = "foo bar"
                            p0 += p0
                            var v28 uint16 = 32767
                            var v29 int8 = (p0 ^ (p0 % (((p0 % ((p0) | 1))) | 1)))
                            _ = v27
                            _ = v28
                            _ = v29
                        case "hello":
                            if (42 == 2) {
                                p0 = p0
                                p0 += p0
                            } else {
                                _ = len(v22)
                                fmt.Println(fi21, p1, fi11)
                                v30 := GenS1{F0: -1, F1: (p1 / ((p1) | 1)), F2: 2}
                                p1 = 2
                                _ = v30
                            }
                            var v31 bool = (1 <= 1)
                            fmt.Println(fi21, p1)
                            _ = v31
                        case "\n\t":
                            p0 = (p0 % (((p0 | p0)) | 1))
                        }
                    } else {
                        fmt.Println(p1, fi11, p0)
                        if len(v22) >= 2 {
                            v22 = v22[1:len(v22)]
                        }
                    }
                } else {
                    v32 := GenS3{F0: 42}
                    if (!true) {
                        if false {
                            v33 := GenS4{F0: p1, F1: (p1 % ((2147483648) | 1))}
                            _ = v33
                        }
                        p0 -= p0
                        _ = len(v22)
                    }
                    v34 := GenS5{F0: 169, F1: (p1 * (p1 << ((136) & 63)))}
                    v35 := GenS2{F0: (p1 % (((p1 * 1)) | 1)), F1: 42, F2: 1000}
                    p1 += 65535
                    p0 = -75
                    _ = v32
                    _ = v34
                    _ = v35
                }
                if ((1000 > 1) && false) {
                    v36 := GenS1{F0: 42, F1: genFn0((65535 <= 2)), F2: 262}
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                    _ = v36
                } else {
                    if false {
                        if false {
                            p1 = 43
                            v37 := GenS5{F0: 127, F1: (^(p1))}
                            p0 -= (p0 + p0)
                            var v38 bool = true
                            _ = v37
                            _ = v38
                        }
                        _ = len(v22)
                    }
                    switch -637 {
                    case 0:
                        fmt.Println(fi21)
                    case -2147483648:
                        if len(v22) >= 2 {
                            v22 = v22[1:len(v22)]
                        }
                    default:
                        if (true && false) {
                            p1 -= (p1 & (p1 >> ((68) & 63)))
                            var v39 string = (("a" + "foo bar") + "")
                            _ = v39
                        }
                    }
                }
                switch ("foo bar" + "a") {
                case "foo bar":
                    if true {
                        v40 := GenS6{F0: -975}
                        if false {
                            if (!((p1 + p1) >= (p1 ^ genFn0(false)))) {
                                switch "foo bar" {
                                case "":
                                    var v41 int8 = p0
                                    if (-32768 >= 169) {
                                        if false {
                                            var v42 int32 = -157
                                            v41 = (p0 >> ((65535) & 63))
                                            _ = v42
                                        } else {
                                            v43 := GenS5{F0: 241, F1: p1}
                                            p0 += (v41 & int8(fi11))
                                            v41 = 127
                                            _ = v43
                                        }
                                    }
                                    _ = v41
                                case "hello":
                                    var v44 int = 32768
                                    _ = v44
                                default:
                                    p0 -= p0
                                }
                                p0 = (p0 ^ p0)
                            }
                        }
                        _ = len(v22)
                        _ = v40
                    } else {
                        if ((p0 >> ((1) & 63)) <= p0) {
                            p0 = p0
                        } else {
                            fmt.Println(p0, p0)
                            v22 = append(v22, fi21)
                        }
                        switch -9223372036854775808 {
                        case 9223372036854775807:
                            v45 := GenS2{F0: p1, F1: 255, F2: 42}
                            _ = v45
                        case 0:
                            p1 = (p1 << ((0) & 63))
                        default:
                            var v46 uint = 65535
                            _ = v46
                        }
                        switch fi21 {
                        case 4294967295:
                            switch fi21 {
                            case 42:
                                p1 -= (p1 % ((p1) | 1))
                            default:
                                switch ("hello" + "foo bar") {
                                case "foo bar":
                                    switch fi11 {
                                    case -526:
                                        var v47 uint8 = 1
                                        _ = v47
                                    }
                                case "hello":
                                    p1 -= p1
                                }
                            }
                        case -2147483648:
                            p1 = p1
                        case -1:
                            fmt.Println(p0)
                            var v48 int8 = p0
                            _ = v48
                        default:
                            p1 += (^(p1))
                        }
                        fmt.Println(fi11)
                        p1 = 2
                    }
                case "hello":
                    switch ("foo bar" + "foo bar") {
                    case "":
                        if true {
                            p1 += genFn0(false)
                            switch fi11 {
                            case -1:
                                if true {
                                    if false {
                                        if (!true) {
                                            var v49 int = fi11
                                            _ = v49
                                        }
                                        if len(v22) >= 2 {
                                            v22 = v22[1:len(v22)]
                                        }
                                        if true {
                                            _ = len(v22)
                                        } else {
                                            var v50 uint16 = 705
                                            _ = v50
                                        }
                                        if len(v22) > 0 {
                                            v22[0] = (fi21 ^ 1)
                                        }
                                    }
                                }
                            }
                            p1 -= 65535
                            fmt.Println(p1)
                            fmt.Println(p1)
                            if true {
                                var v51 int32 = 1
                                _ = v51
                            }
                            v52 := GenS1{F0: 1, F1: 79, F2: 546}
                            _ = len(v22)
                            _ = v52
                        } else {
                            v53 := GenS2{F0: 166, F1: 255, F2: 1000}
                            var v54 int = fi21
                            p1 += (^(p1))
                            var v55 float64 = ((0.1 / 1.0) * 0.0)
                            _ = v53
                            _ = v54
                            _ = v55
                        }
                    case "a":
                        var v56 int16 = 32767
                        _ = v56
                    }
                case "a":
                    fmt.Println(fi21, p1)
                }
                fmt.Println(fi21, fi11, fi21)
                v57 := GenS6{F0: 0}
                if true {
                    var v58 uint16 = 81
                    var v59 uint32 = 1
                    _ = v58
                    _ = v59
                }
                var v60 uint8 = 127
                _ = v57
                _ = v60
            }
            _ = fi21
            _ = v22
        }
        var v61 uint16 = 65535
        _ = fi11
        _ = v61
    }
    if ("a" >= "foo bar") {
        switch -9223372036854775808 {
        case 42:
            fmt.Println(p0)
        case -1:
            v62 := GenS7{F0: 1}
            _ = v62
        default:
            var v63 float64 = 3.14159
            _ = v63
        }
        p0 += (p0 + p0)
        p0 += -128
    } else {
        v64 := []int{-585, -135, 42, 2147483647}
        _ = v64
    }
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]string)
    v3 := make(map[int]uint32)
    var v4 uint64 = 1
    if true {
        for fi65 := 0; fi65 < 16; fi65++ {
            delete(v2, ("\n\t" + "foo bar"))
            if false {
                v4 = v4
                if false {
                    for fi66 := 0; fi66 < 12; fi66++ {
                        v4 -= v4
                        _ = fi66
                    }
                    for fi67 := 0; fi67 < 19; fi67++ {
                        fmt.Println(fi65, fi67)
                        v4 = v4
                        if ((!true) || false) {
                            delete(v1, "")
                            v68 := []int{(^(fi67)), -9223372036854775808, fi65}
                            _ = v68
                        }
                        switch (fi67 / ((fi65) | 1)) {
                        case 612:
                            v4 = v4
                        default:
                            v69 := GenS3{F0: 0}
                            var v70 uint = 649
                            _ = v69
                            _ = v70
                        }
                        _ = fi67
                    }
                    var v71 int = (^(fi65))
                    _ = v71
                }
                switch (-(fi65)) {
                case 54:
                    v4 -= v4
                default:
                    var v72 uint64 = (v4 & v4)
                    _ = v72
                }
            } else {
                if false {
                    switch (fi65 / ((fi65) | 1)) {
                    case 127:
                        if (true && (0.0 != float64(1.5))) {
                            v4 = v4
                            fmt.Println(fi65, v4)
                            if (-1 < fi65) {
                                if (!(!(fi65 == fi65))) {
                                    var v73 int32 = 2147483647
                                    for fi74 := 0; fi74 < 14; fi74++ {
                                        fmt.Println(v73)
                                        _ = fi74
                                    }
                                    _ = v73
                                }
                            }
                        }
                    case -423:
                        v4 -= (^(v4))
                    }
                    _ = len(v3)
                } else {
                    var v75 float64 = (-2.25 * 1.5)
                    _ = v75
                }
                for fi76 := 0; fi76 < 17; fi76++ {
                    var v77 uint8 = 127
                    v4 -= 18446744073709551615
                    _ = len(v3)
                    _ = fi76
                    _ = v77
                }
                for fi78 := 0; fi78 < 1; fi78++ {
                    v4 += v4
                    fmt.Println(fi78)
                    if true {
                        var v79 int16 = 489
                        var v80 float64 = (1e10 * 0.1)
                        fmt.Println(fi78, v4)
                        _ = v79
                        _ = v80
                    }
                    fmt.Println(fi78, fi65, v4)
                    v4 += 42
                    _ = fi78
                }
            }
            _ = fi65
        }
        var v81 uint64 = (v4 / ((v4) | 1))
        _ = v81
    } else {
        var v82 int = 567
        v83 := GenS2{F0: 42, F1: 466, F2: 1}
        _ = v82
        _ = v83
    }
    v4 = (^(v4))
    v4 += (v4 ^ (v4 * (v4 * v4)))
    for fi84 := 0; fi84 < 1; fi84++ {
        if (-1 <= 72) {
            if (1 != -469) {
                switch "hello" {
                case "hello":
                    v85 := GenS8{F0: 1}
                    _ = v85
                default:
                    if true {
                        var v86 int32 = 1
                        switch (fi84 % ((int(v4)) | 1)) {
                        case 128:
                            var v87 int8 = 122
                            _ = v87
                        default:
                            var v88 int32 = (^(v86))
                            _ = v88
                        }
                        var v89 uint16 = 991
                        _ = v86
                        _ = v89
                    } else {
                        var v90 uint8 = 0
                        switch "\n\t" {
                        case "a":
                            v91 := []int{fi84, 32768, -1}
                            _ = v91
                        }
                        v90 = 2
                        _ = v90
                    }
                }
                v4 -= (v4 + 1)
                switch -454 {
                case -1:
                    var v92 bool = true
                    _ = v92
                case 127:
                    switch "\n\t" {
                    case "\n\t":
                        switch "" {
                        case "a":
                            v4 = 1
                        case "":
                            for fi93 := 0; fi93 < 16; fi93++ {
                                if ((!true) && true) {
                                    if (!true) {
                                        fmt.Println(fi93, fi93, fi84)
                                        fmt.Println(v4, fi84)
                                        var v94 uint16 = 2
                                        var v95 int16 = 1
                                        v96 := GenS5{F0: 0, F1: genFn0(((1 <= 51) || ((v4 / (((v4 / (((v4 % ((4294967295) | 1))) | 1))) | 1)) >= v4)))}
                                        _ = v94
                                        _ = v95
                                        _ = v96
                                    } else {
                                        var v97 string = "foo bar"
                                        delete(v1, v97)
                                        var v98 string = (v97 + (v97 + ("" + v97)))
                                        v99 := []int{-789}
                                        _ = v97
                                        _ = v98
                                        _ = v99
                                    }
                                    v100 := []int{(fi93 % ((1) | 1)), (fi84 % ((fi84) | 1))}
                                    v4 += 483
                                    _ = v100
                                } else {
                                    var v101 float32 = 1.5
                                    _ = v101
                                }
                                fmt.Println(fi84, v4, v4)
                                if true {
                                    v102 := GenS8{F0: 2}
                                    _ = v102
                                }
                                _ = fi93
                            }
                        default:
                            if (67 < 154) {
                                var v103 uint = 18446744073709551615
                                _ = v103
                            }
                        }
                    default:
                        var v104 uint64 = (v4 | (v4 << ((1000) & 63)))
                        _ = v104
                    }
                case 678:
                    fmt.Println(v4)
                }
            }
            if true {
                if false {
                    var v105 uint = 255
                    if true {
                        var v106 uint16 = 42
                        v105 -= (v105 % (((^(v105))) | 1))
                        _ = v106
                    } else {
                        if (255 < uint8(v105)) {
                            var v107 uint32 = 2
                            v108 := GenS8{F0: 178}
                            _ = v107
                            _ = v108
                        } else {
                            var v109 string = "foo bar"
                            _ = v109
                        }
                        v110 := GenS1{F0: 0, F1: 2, F2: (v105 / ((genFn1(127, 2147483648)) | 1))}
                        _ = v110
                    }
                    _ = v105
                } else {
                    v111 := GenS2{F0: 65535, F1: 4294967295, F2: -1}
                    if false {
                        var v112 uint64 = (v4 + (v4 ^ v4))
                        var v113 uint = 49
                        v4 = (v112 >> ((v113) & 63))
                        _ = v112
                        _ = v113
                    } else {
                        fmt.Println(v4, fi84)
                        v3[fi84] = 170
                        v114 := GenS9{F0: (v4 / ((1) | 1)), F1: 42, F2: v4}
                        _ = v114
                    }
                    _ = v111
                }
                var v115 uint16 = 121
                v116 := []int{(fi84 * fi84), fi84, fi84}
                _ = v115
                _ = v116
            } else {
                for fi117 := 0; fi117 < 6; fi117++ {
                    fmt.Println(fi117)
                    var v118 uint8 = 17
                    _ = fi117
                    _ = v118
                }
            }
        } else {
            v4 = (v4 % ((v4) | 1))
        }
        for fi119 := 0; fi119 < 10; fi119++ {
            v120 := []int{fi84, (^(fi119))}
            var v121 string = "foo bar"
            _ = fi119
            _ = v120
            _ = v121
        }
        _ = fi84
    }
    for fi122 := 0; fi122 < 14; fi122++ {
        v123 := []int{(fi122 << ((187) & 63)), (^(fi122))}
        _ = fi122
        _ = v123
    }
    if false {
        v124 := GenS9{F0: (v4 | 2), F1: v4, F2: (^(v4))}
        if false {
            var v125 float32 = (1e10 - 1.5)
            _ = v125
        } else {
            v126 := GenS10{F0: -1, F1: 18446744073709551615, F2: -1}
            _ = v126
        }
        _ = v124
    } else {
        v127 := GenS11{F0: 127}
        switch -9223372036854775808 {
        case 4294967295:
            if true {
                var v128 bool = true
                v129 := GenS6{F0: 0}
                v130 := GenS12{F0: -9223372036854775808, F1: 65535, F2: 42}
                _ = v128
                _ = v129
                _ = v130
            } else {
                if (9223372036854775807 <= -1) {
                    var v131 int16 = 1000
                    _ = v131
                } else {
                    fmt.Println(v4)
                }
                for fi132 := 0; fi132 < 15; fi132++ {
                    v133 := []int{}
                    v134 := GenS6{F0: -1}
                    _ = fi132
                    _ = v133
                    _ = v134
                }
            }
        case 305:
            var v135 int = 128
            _ = v135
        case -1:
            for fi136 := 0; fi136 < 7; fi136++ {
                v4 = 6
                var v137 int32 = -615
                _ = fi136
                _ = v137
            }
        }
        if (1 >= 0) {
            v4 += v4
        } else {
            switch "" {
            case "foo bar":
                switch ("\n\t" + ("hello" + "")) {
                case "\n\t":
                    switch 2147483647 {
                    case -9223372036854775808:
                        v4 = 2
                    case -814:
                        switch 2147483647 {
                        case -1:
                            var v138 int8 = 42
                            _ = v138
                        case 127:
                            switch 1 {
                            case 9223372036854775807:
                                v4 -= 42
                            case 626:
                                for fi139 := 0; fi139 < 20; fi139++ {
                                    if true {
                                        var v140 uint8 = 255
                                        _ = v140
                                    }
                                    fmt.Println(v4, fi139)
                                    var v141 uint8 = 42
                                    _ = fi139
                                    _ = v141
                                }
                            case -239:
                                _ = len(v2)
                            default:
                                v4 += v4
                            }
                        }
                    }
                case "":
                    if (0 == -35) {
                        v4 += v4
                        var v142 bool = ((!((42 != 42) && (false || (false && true)))) && true)
                        _ = v142
                    } else {
                        switch "foo bar" {
                        case "hello":
                            var v143 uint32 = 768
                            _ = v143
                        case "a":
                            for fi144 := 0; fi144 < 4; fi144++ {
                                switch (fi144 | (fi144 >> ((42) & 63))) {
                                case 32768:
                                    v145 := GenS4{F0: 4294967295, F1: 65535}
                                    _ = v145
                                }
                                _ = fi144
                            }
                        default:
                            if true {
                                v146 := GenS10{F0: -173, F1: 18446744073709551615, F2: 0}
                                _ = v146
                            }
                        }
                        if true {
                            v4 = v4
                        } else {
                            if (18446744073709551615 < 430) {
                                v4 = (v4 & (v4 / ((v4) | 1)))
                            } else {
                                v147 := []int{0}
                                var v148 bool = (540 > 0)
                                fmt.Println(v4, v148)
                                _ = v147
                                _ = v148
                            }
                        }
                    }
                }
            case "a":
                var v149 uint8 = 0
                _ = v149
            }
            if false {
                v150 := v1[(((((("foo bar" + ((("a" + "\n\t") + ((("hello" + ("\n\t" + "hello")) + "a") + ("\n\t" + "\n\t"))) + "\n\t")) + "hello") + ("\n\t" + "\n\t")) + "") + ("a" + "foo bar")) + "\n\t")]
                _ = v150
            }
            if (!(false && false)) {
                var v151 int32 = -1
                var v152 int8 = 0
                _ = v151
                _ = v152
            }
        }
        _ = v127
    }
    for fi153 := 0; fi153 < 11; fi153++ {
        v154 := GenS7{F0: (fi153 / ((fi153) | 1))}
        switch -9223372036854775808 {
        case -663:
            v4 = 901
        case 471:
            var v155 uint16 = 65535
            _ = v155
        case 32768:
            v156 := []int{4294967295, 4294967295, (fi153 % (((fi153 << ((266) & 63))) | 1)), (fi153 / ((fi153) | 1))}
            _ = v156
        }
        v4 = (v4 * v4)
        _ = fi153
        _ = v154
    }
    var v157 uint8 = 13
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v157:", v157)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v157
    _ = v2
    _ = v3
    _ = v4
}
