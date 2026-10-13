package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 uint32; F1 bool; F2 int64 }
type GenS5 struct { F0 int; F1 int64; F2 int }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int8; F1 int8; F2 int }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint16; F1 uint64 }
type GenS13 struct { F0 uint8; F1 int }
func genFn0(p0 int8, p1 uint8, p2 uint32) int32 {
    p0 = p0
    if false {
        for fi6 := 0; fi6 < 6; fi6++ {
            if (!(uint16(p2) > 255)) {
                for fi7 := 0; fi7 < 7; fi7++ {
                    p1 = (p1 - 255)
                    _ = fi7
                }
                fmt.Println(p1)
            }
            v8 := []int{fi6, 0, fi6}
            if true {
                p1 += (p1 - (^(p1)))
                if len(v8) > 0 {
                    v8[0] = fi6
                }
                for fi9 := 0; fi9 < 20; fi9++ {
                    v10 := GenS1{F0: 4294967295}
                    var v11 int8 = (-(p0))
                    switch (fi9 & int(p1)) {
                    case 32767:
                        p0 -= (v11 | (v11 + (v11 << ((255) & 63))))
                    default:
                        v8 = append(v8, fi6)
                    }
                    v12 := GenS2{F0: ""}
                    switch "\n\t" {
                    case "foo bar":
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                    case "\n\t":
                        if ((!false) || false) {
                            fmt.Println(v11, fi9)
                        }
                    case "hello":
                        if true {
                            var v13 uint = 255
                            var v14 uint16 = 0
                            if len(v8) > 0 {
                                v8[0] = (^(fi9))
                            }
                            _ = v13
                            _ = v14
                        }
                    }
                    if len(v8) > 0 {
                        v8[0] = fi6
                    }
                    switch "" {
                    case "a":
                        _ = len(v8)
                    case "foo bar":
                        if (42 < 42) {
                            if (false && (32767 > 0)) {
                                var v15 float32 = (((-2.25 * 1e10) * (1.5 / 1e10)) / (-2.25 * -2.25))
                                _ = v15
                            }
                            var v16 uint32 = p2
                            _ = v16
                        } else {
                            p1 += p1
                            v8 = append(v8, 9223372036854775807)
                        }
                    }
                    var v17 int64 = 63
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                    _ = v17
                }
            }
            _ = fi6
            _ = v8
        }
    } else {
        p1 -= p1
        var v18 uint32 = 2147483648
        p2 = v18
        _ = v18
    }
    p1 += p1
    var v19 uint16 = 283
    v20 := GenS3{F0: 255, F1: (true || true)}
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    return 0
}

func main() {
    var v1 uint8 = 255
    var v2 int = -145
    v3 := []string{""}
    v4 := make(map[int]bool)
    v5 := []uint32{2147483648}
    v1 -= v1
    var v21 float32 = (0.0 + 3.14159)
    for fi22 := 0; fi22 < 1; fi22++ {
        fmt.Println(v1)
        v1 -= (v1 | 175)
        _ = fi22
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        for fi24 := 0; fi24 < 14; fi24++ {
            if ((0 == 42) || false) {
                if ((18446744073709551615 > 986) && (2 < 638)) {
                    var v25 uint8 = (v1 & (v1 << ((65535) & 63)))
                    var v26 string = ""
                    v27 := GenS2{F0: "foo bar"}
                    var v28 float32 = 3.14159
                    if true {
                        v29 := GenS3{F0: (v1 ^ (v1 * v1)), F1: (((!true) && false) && (!((42 <= 216) && false)))}
                        v2 = (fi24 * (^(fi23)))
                        _ = v29
                    }
                    _ = v25
                    _ = v26
                    _ = v27
                    _ = v28
                }
                if len(v5) > 0 {
                    v5[0] = 510
                }
                fmt.Println(fi24, fi24, v2)
            }
            _ = fi24
        }
        switch (fi23 - fi23) {
        case 1:
            v30 := GenS2{F0: ("foo bar" + (((("a" + "a") + "hello") + "\n\t") + ("\n\t" + "a")))}
            _ = v30
        }
        _ = fi23
    }
    v1 += 128
    for fi31 := 0; fi31 < 14; fi31++ {
        switch "\n\t" {
        case "":
            v2 += fi31
        case "\n\t":
            v4[-121] = (42 < 32767)
        case "hello":
            v2 -= 1
        default:
            var v32 int32 = genFn0(-1, v1, 810)
            _ = v32
        }
        fmt.Println(fi31, v21)
        var v33 string = ("a" + "a")
        _ = fi31
        _ = v33
    }
    switch -2147483648 {
    case 32767:
        fmt.Println(v2, v21, v21)
    case 87:
        var v34 int8 = -1
        _ = v34
    case 31:
        var v35 int8 = 44
        _ = v35
    default:
        v36 := GenS1{F0: 2}
        _ = v36
    }
    for fi37 := 0; fi37 < 15; fi37++ {
        v38 := GenS3{F0: (v1 >> ((0) & 63)), F1: (65535 > 1)}
        _ = fi37
        _ = v38
    }
    var v39 float32 = (v21 / -2.25)
    var v40 uint64 = 0
    if true {
        var v41 float64 = (((0.0 - 1e10) * 1.5) + -2.25)
        if ("" != "hello") {
            v39 += 3.14159
            v42 := GenS1{F0: (v40 & (v40 << ((42) & 63)))}
            _ = v42
        } else {
            for fi43 := 0; fi43 < 14; fi43++ {
                var v44 int64 = 9223372036854775807
                _ = fi43
                _ = v44
            }
            v4[(v2 << ((0) & 63))] = (0 <= 387)
            for fi45 := 0; fi45 < 4; fi45++ {
                for fi46 := 0; fi46 < 5; fi46++ {
                    fmt.Println(fi45)
                    _ = fi46
                }
                delete(v4, 32767)
                delete(v4, 968)
                _ = fi45
            }
        }
        switch "hello" {
        case "\n\t":
            var v47 int8 = 100
            _ = v47
        case "foo bar":
            for fi48 := 0; fi48 < 10; fi48++ {
                v2 -= 32767
                _ = fi48
            }
        case "a":
            fmt.Println(v1, v2, v39)
        }
        _ = v41
    }
    switch "\n\t" {
    case "hello":
        v49 := GenS4{F0: 4294967295, F1: false, F2: 42}
        _ = v49
    case "a":
        fmt.Println(v39, v1)
    case "":
        switch "" {
        case "hello":
            var v50 float64 = 1e10
            _ = v50
        default:
            for fi51 := 0; fi51 < 17; fi51++ {
                var v52 uint = 65535
                _ = fi51
                _ = v52
            }
        }
    }
    var v53 int64 = -1
    switch "hello" {
    case "":
        switch (v2 * v2) {
        case 128:
            var v54 uint = 1
            _ = v54
        default:
            if (!false) {
                var v55 float32 = ((0.1 - v21) * v21)
                if false {
                    var v56 int8 = -128
                    _ = v56
                } else {
                    if (v1 == 249) {
                        var v57 uint32 = 465
                        if (v40 == v40) {
                            if true {
                                if ((!true) && true) {
                                    v58 := v4[v2]
                                    _ = v58
                                }
                                for fi59 := 0; fi59 < 2; fi59++ {
                                    switch "\n\t" {
                                    case "hello":
                                        for fi60 := 0; fi60 < 7; fi60++ {
                                            var v61 int32 = -805
                                            v53 = v53
                                            _ = fi60
                                            _ = v61
                                        }
                                    case "a":
                                        for fi62 := 0; fi62 < 5; fi62++ {
                                            fmt.Println(v21, v2, v39)
                                            switch int(v57) {
                                            case 32767:
                                                if (255 != 593) {
                                                    switch (v2 | fi62) {
                                                    case 1:
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    case 4294967295:
                                                        if (0 < 570) {
                                                            if (true || (v1 >= v1)) {
                                                                v40 = 178
                                                                v2 = 9223372036854775807
                                                                v4[(^(fi59))] = false
                                                            } else {
                                                                _ = len(v3)
                                                            }
                                                        } else {
                                                            if (-1 <= -59) {
                                                                if (false && false) {
                                                                    switch ("hello" + "\n\t") {
                                                                    case "foo bar":
                                                                        var v63 int = -842
                                                                        _ = v63
                                                                    }
                                                                    var v64 uint64 = (v40 + 0)
                                                                    var v65 string = "a"
                                                                    v66 := GenS5{F0: int(v57), F1: v53, F2: (v2 % ((-719) | 1))}
                                                                    if true {
                                                                        fmt.Println(v65, fi62)
                                                                        if len(v5) > 0 {
                                                                            v5[0] = v57
                                                                        }
                                                                        if true {
                                                                            _ = len(v4)
                                                                            var v67 uint = 65535
                                                                            _ = v67
                                                                        } else {
                                                                            switch (fi59 & (fi62 * (fi59 % ((9223372036854775807) | 1)))) {
                                                                            case 2147483647:
                                                                                var v68 float64 = -2.25
                                                                                _ = v68
                                                                            }
                                                                            v2 -= fi62
                                                                        }
                                                                    }
                                                                    _ = v64
                                                                    _ = v65
                                                                    _ = v66
                                                                }
                                                            }
                                                            v39 = v39
                                                            fmt.Println(v1, v2)
                                                        }
                                                    default:
                                                        var v69 int8 = 42
                                                        _ = v69
                                                    }
                                                }
                                            case 821:
                                                fmt.Println(v2, v53, v2)
                                            default:
                                                var v70 string = (("\n\t" + ("\n\t" + "")) + "hello")
                                                _ = v70
                                            }
                                            _ = fi62
                                        }
                                    case "foo bar":
                                        var v71 uint32 = 2
                                        _ = v71
                                    default:
                                        for fi72 := 0; fi72 < 9; fi72++ {
                                            v73 := GenS6{F0: 373, F1: "foo bar"}
                                            fmt.Println(v53)
                                            if true {
                                                v1 = v1
                                                if (1e10 == 3.14159) {
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    if (false || true) {
                                                        v53 -= (v53 ^ (v53 << ((0) & 63)))
                                                    } else {
                                                        var v74 uint = 42
                                                        v55 -= 0.1
                                                        _ = v74
                                                    }
                                                }
                                                v4[(fi72 / ((v2) | 1))] = false
                                            } else {
                                                var v75 float64 = 0.0
                                                _ = v75
                                            }
                                            switch (fi72 >> ((18446744073709551615) & 63)) {
                                            case -2147483648:
                                                var v76 int = v2
                                                _ = v76
                                            default:
                                                v77 := GenS1{F0: 18446744073709551615}
                                                _ = v77
                                            }
                                            v78 := v4[(v2 - (v2 << ((4294967295) & 63)))]
                                            _ = fi72
                                            _ = v73
                                            _ = v78
                                        }
                                    }
                                    _ = fi59
                                }
                            } else {
                                switch v2 {
                                case 9223372036854775807:
                                    if (float64(v21) <= (1.5 * (((0.1 / 1e10) / (-2.25 - float64(0.0))) - float64(float32(((0.0 * 1.5) * 0.1)))))) {
                                        if ((195 == 745) || true) {
                                            var v79 int = v2
                                            _ = v79
                                        } else {
                                            var v80 float32 = float32(0.1)
                                            if true {
                                                var v81 uint32 = v57
                                                if (!true) {
                                                    switch -9223372036854775808 {
                                                    case -975:
                                                        fmt.Println(v1)
                                                    case 0:
                                                        v82 := GenS7{F0: "a"}
                                                        _ = v82
                                                    }
                                                } else {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    var v83 float32 = v21
                                                    for fi84 := 0; fi84 < 18; fi84++ {
                                                        var v85 string = "\n\t"
                                                        for fi86 := 0; fi86 < 8; fi86++ {
                                                            v87 := GenS3{F0: 147, F1: false}
                                                            _ = fi86
                                                            _ = v87
                                                        }
                                                        _ = fi84
                                                        _ = v85
                                                    }
                                                    _ = v83
                                                }
                                                _ = v81
                                            } else {
                                                v88 := GenS4{F0: uint32(v1), F1: (255 != 827), F2: v53}
                                                var v89 int8 = int8(v53)
                                                var v90 bool = (!(false || (65535 > 255)))
                                                _ = v88
                                                _ = v89
                                                _ = v90
                                            }
                                            switch "hello" {
                                            case "foo bar":
                                                v91 := GenS8{F0: 32767}
                                                _ = v91
                                            case "\n\t":
                                                v21 = (float32(1.5) - v80)
                                            default:
                                                if false {
                                                    var v92 int8 = -1
                                                    v80 -= 1e10
                                                    _ = v92
                                                } else {
                                                    if (!true) {
                                                        fmt.Println(v1, v40, v55)
                                                        _ = len(v3)
                                                    }
                                                    for fi93 := 0; fi93 < 13; fi93++ {
                                                        var v94 uint16 = 876
                                                        switch "" {
                                                        case "foo bar":
                                                            v95 := GenS6{F0: (v94 >> ((822) & 63)), F1: ((("a" + "a") + (("\n\t" + "\n\t") + ((("foo bar" + ("a" + "")) + "hello") + ""))) + ("hello" + "hello"))}
                                                            _ = v95
                                                        case "a":
                                                            v57 += v57
                                                        case "hello":
                                                            for fi96 := 0; fi96 < 12; fi96++ {
                                                                if true {
                                                                    if ((v40 << ((4294967295) & 63)) != (v40 & (v40 + (v40 ^ (v40 | 18446744073709551615))))) {
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        if false {
                                                                            fmt.Println(v57)
                                                                            v57 += (v57 + (v57 * v57))
                                                                        }
                                                                        v3 = append(v3, ((("hello" + "") + "hello") + (("\n\t" + "hello") + "foo bar")))
                                                                    } else {
                                                                        fmt.Println(fi93, v1)
                                                                        v53 -= (v53 << ((42) & 63))
                                                                    }
                                                                }
                                                                fmt.Println(v2, fi93, fi93)
                                                                var v97 uint32 = 208
                                                                _ = fi96
                                                                _ = v97
                                                            }
                                                        }
                                                        for fi98 := 0; fi98 < 1; fi98++ {
                                                            v99 := GenS1{F0: v40}
                                                            var v100 bool = (0 != 0)
                                                            _ = fi98
                                                            _ = v100
                                                            _ = v99
                                                        }
                                                        _ = fi93
                                                        _ = v94
                                                    }
                                                    var v101 string = "\n\t"
                                                    _ = v101
                                                }
                                            }
                                            _ = v80
                                        }
                                        v39 -= v55
                                        for fi102 := 0; fi102 < 5; fi102++ {
                                            var v103 int64 = v53
                                            var v104 float32 = 0.1
                                            _ = fi102
                                            _ = v103
                                            _ = v104
                                        }
                                    } else {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        if true {
                                            if len(v3) > 0 {
                                                v3[0] = (("foo bar" + "foo bar") + ("hello" + (("hello" + "hello") + ("" + ("a" + "foo bar")))))
                                            }
                                        }
                                        _ = len(v4)
                                    }
                                case 1:
                                    var v105 uint16 = 255
                                    _ = v105
                                }
                            }
                            v40 -= (v40 << ((uint(v1)) & 63))
                        } else {
                            v106 := GenS4{F0: 1, F1: (4294967295 == 255), F2: v53}
                            _ = v106
                        }
                        _ = v57
                    }
                }
                v39 = float32(1.5)
                _ = v55
            }
        }
    }
    for fi107 := 0; fi107 < 11; fi107++ {
        for fi108 := 0; fi108 < 8; fi108++ {
            v109 := GenS9{F0: false}
            delete(v4, 32767)
            var v110 float64 = (3.14159 / (-2.25 - ((float64((0.0 + (0.1 + (float32(((0.0 * (3.14159 - 1.5)) / 1e10)) / 1.5)))) * (float64((v39 * v39)) + -2.25)) - (0.1 * 3.14159))))
            _ = fi108
            _ = v109
            _ = v110
        }
        _ = fi107
    }
    for fi111 := 0; fi111 < 14; fi111++ {
        v112 := GenS10{F0: 1, F1: 1, F2: fi111}
        fmt.Println(v40, v2, v21)
        _ = fi111
        _ = v112
    }
    if (-113 == -1) {
        switch "a" {
        case "a":
            v1 += 255
        case "":
            for fi113 := 0; fi113 < 15; fi113++ {
                var v114 int8 = 1
                switch -481 {
                case 128:
                    v115 := GenS1{F0: 572}
                    _ = v115
                default:
                    switch (fi113 ^ v2) {
                    case 32767:
                        v116 := GenS11{F0: 0}
                        _ = v116
                    default:
                        fmt.Println(v1, v1)
                    }
                }
                var v117 bool = true
                _ = fi113
                _ = v114
                _ = v117
            }
        }
    } else {
        if false {
            if true {
                switch -1 {
                case 766:
                    switch "hello" {
                    case "a":
                        var v118 int8 = -28
                        _ = v118
                    case "\n\t":
                        v39 -= float32(float64((v21 / v39)))
                    case "foo bar":
                        v119 := GenS3{F0: 1, F1: (1.5 <= 3.14159)}
                        _ = v119
                    default:
                        if (genFn0(0, 2, 65535) != 42) {
                            v120 := GenS12{F0: 65535, F1: v40}
                            v53 += (v53 ^ 9223372036854775807)
                            _ = v120
                        } else {
                            switch 128 {
                            case -2147483648:
                                var v121 string = "hello"
                                _ = v121
                            case 32768:
                                var v122 int64 = v53
                                _ = v122
                            case -78:
                                v123 := GenS11{F0: 2147483648}
                                _ = v123
                            default:
                                v2 = (v2 & (v2 ^ 32767))
                            }
                        }
                    }
                default:
                    if (65535 <= 154) {
                        if (0.0 < 1.5) {
                            v21 = 0.0
                            for fi124 := 0; fi124 < 11; fi124++ {
                                for fi125 := 0; fi125 < 5; fi125++ {
                                    v126 := GenS5{F0: v2, F1: v53, F2: 127}
                                    v21 -= -2.25
                                    v40 += v40
                                    _ = fi125
                                    _ = v126
                                }
                                v127 := GenS13{F0: (^(v1)), F1: (-(v2))}
                                _ = fi124
                                _ = v127
                            }
                            if false {
                                var v128 float64 = 0.0
                                var v129 uint16 = 682
                                _ = v128
                                _ = v129
                            } else {
                                v130 := GenS7{F0: "\n\t"}
                                delete(v4, (v2 ^ (v2 << ((255) & 63))))
                                _ = v130
                            }
                        }
                        for fi131 := 0; fi131 < 18; fi131++ {
                            if true {
                                fmt.Println(v21)
                                var v132 bool = (255 < 938)
                                var v133 bool = (("" + ("\n\t" + "a")) == "hello")
                                _ = v132
                                _ = v133
                            } else {
                                v4[v2] = (0 > 770)
                                _ = len(v4)
                                for fi134 := 0; fi134 < 12; fi134++ {
                                    var v135 uint8 = v1
                                    fmt.Println(v2, v2, v2)
                                    _ = fi134
                                    _ = v135
                                }
                            }
                            switch 0 {
                            case -2147483648:
                                v136 := v4[(fi131 | (fi131 ^ v2))]
                                _ = v136
                            case -746:
                                delete(v4, fi131)
                            default:
                                fmt.Println(v53, v53, v1)
                            }
                            _ = fi131
                        }
                        fmt.Println(v40, v53, v2)
                    }
                }
                for fi137 := 0; fi137 < 18; fi137++ {
                    v138 := GenS12{F0: 145, F1: v40}
                    var v139 float64 = -2.25
                    _ = fi137
                    _ = v138
                    _ = v139
                }
                v5 = append(v5, 904)
            } else {
                v2 = v2
                var v140 uint16 = 2
                if len(v5) > 0 {
                    v5[0] = 541
                }
                _ = v140
            }
        } else {
            v141 := GenS10{F0: 127, F1: 127, F2: (v2 / ((4294967295) | 1))}
            delete(v4, v2)
            _ = v141
        }
        if true {
            for fi142 := 0; fi142 < 11; fi142++ {
                switch "" {
                case "foo bar":
                    fmt.Println(v39, v1)
                }
                switch "foo bar" {
                case "a":
                    switch ("foo bar" + "hello") {
                    case "foo bar":
                        fmt.Println(v21)
                    }
                }
                _ = fi142
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v53
}
