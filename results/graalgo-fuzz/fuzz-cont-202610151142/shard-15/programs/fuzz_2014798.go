package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 uint; F2 int32 }
type GenS4 struct { F0 int32; F1 int32; F2 int16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint32; F1 uint64; F2 bool }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 string; F1 int16 }
type GenS9 struct { F0 uint16; F1 int64 }
type GenS10 struct { F0 uint64; F1 int }
type GenS11 struct { F0 int64; F1 uint; F2 uint8 }
type GenS12 struct { F0 string; F1 string; F2 uint }
type GenS13 struct { F0 int8; F1 uint; F2 int32 }
type GenS14 struct { F0 int64; F1 int64; F2 int32 }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 int16; F1 int8; F2 int }
type GenS17 struct { F0 bool; F1 string; F2 int16 }
func genFn0(p0 bool, p1 uint8, p2 int32) int16 {
    var v7 int32 = 42
    switch 9223372036854775807 {
    case 0:
        var v8 bool = ((65535 != 42) && p0)
        _ = v8
    case 32768:
        var v9 int64 = 458
        _ = v9
    case 127:
        p0 = true
    default:
        var v10 int16 = 460
        _ = v10
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        v12 := GenS1{F0: -9223372036854775808}
        if false {
            v13 := []int{fi11}
            if p0 {
                var v14 int8 = -1
                _ = v14
            }
            fmt.Println(fi11)
            _ = v13
        } else {
            for fi15 := 0; fi15 < 9; fi15++ {
                var v16 float32 = 0.1
                v17 := []int{fi11, (fi11 ^ fi11)}
                var v18 uint16 = 381
                v16 = 0.1
                var v19 uint16 = 1
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
                _ = v19
            }
            p1 = (p1 << ((18446744073709551615) & 63))
        }
        v20 := GenS2{F0: 0}
        _ = fi11
        _ = v12
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    var v1 float32 = 1.5
    var v2 uint8 = 42
    var v3 uint16 = 255
    var v4 bool = false
    var v5 int8 = -1
    v6 := make(map[uint32]uint32)
    for fi21 := 0; fi21 < 12; fi21++ {
        v22 := GenS3{F0: 255, F1: 4294967295, F2: -362}
        var v23 bool = ((-(v5)) > v5)
        v2 = v2
        _ = fi21
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 10; fi24++ {
        var v25 uint32 = 157
        v26 := []int{(fi24 >> ((255) & 63)), fi24, (fi24 - fi24), fi24}
        _ = fi24
        _ = v25
        _ = v26
    }
    v27 := GenS2{F0: -1}
    v28 := GenS1{F0: -1}
    for fi29 := 0; fi29 < 2; fi29++ {
        switch "hello" {
        case "":
            for fi30 := 0; fi30 < 17; fi30++ {
                var v31 float64 = (0.1 * 0.1)
                switch int(v2) {
                case -2147483648:
                    if (v4 && (0 > 0)) {
                        _ = len(v6)
                    } else {
                        switch -745 {
                        case 42:
                            v4 = v4
                        case 32768:
                            if false {
                                if (!v4) {
                                    v31 = ((v31 + v31) - float64(v1))
                                    v32 := []int{}
                                    v5 -= (^(v5))
                                    var v33 uint16 = 32767
                                    _ = v32
                                    _ = v33
                                }
                                fmt.Println(v31)
                                v3 -= 255
                            }
                            if (true && v4) {
                                v34 := GenS4{F0: 956, F1: -1, F2: -882}
                                switch ("" + "hello") {
                                case "\n\t":
                                    var v35 uint64 = uint64(v2)
                                    _ = v35
                                }
                                _ = v34
                            } else {
                                var v36 string = ("" + "")
                                switch "foo bar" {
                                case "foo bar":
                                    if (!v4) {
                                        fmt.Println(v5)
                                    } else {
                                        v3 = 175
                                    }
                                case "a":
                                    switch v36 {
                                    case "\n\t":
                                        v37 := GenS5{F0: v5}
                                        _ = v37
                                    case "a":
                                        v5 -= (v5 * (v5 % ((42) | 1)))
                                    default:
                                        v1 = float32(1.5)
                                    }
                                }
                                _ = v36
                            }
                            if v4 {
                                v2 = v2
                                var v38 float64 = (v31 + v31)
                                _ = v38
                            } else {
                                v39 := v6[0]
                                _ = v39
                            }
                        case -9223372036854775808:
                            var v40 int64 = -9223372036854775808
                            _ = v40
                        default:
                            v3 += (v3 >> ((255) & 63))
                        }
                        v5 = 42
                    }
                case -1:
                    fmt.Println(v4)
                }
                _ = fi30
                _ = v31
            }
        }
        _ = fi29
    }
    v3 += 591
    v41 := GenS6{F0: 0, F1: 955, F2: v4}
    v42 := GenS2{F0: 0}
    v4 = (false && (640 == -32768))
    fmt.Println(v3)
    var v43 int16 = 32767
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        switch -738 {
        case 127:
            v44 := GenS7{F0: (v5 * (-(v5)))}
            _ = v44
        case -9223372036854775808:
            v5 = (^(v5))
        default:
            var v45 string = "\n\t"
            _ = v45
        }
    case "":
        if (!v4) {
            var v46 uint32 = 18
            _ = v46
        } else {
            for fi47 := 0; fi47 < 18; fi47++ {
                fmt.Println(v1)
                v48 := GenS5{F0: v5}
                _ = fi47
                _ = v48
            }
            v1 = ((((v1 - ((3.14159 - v1) / v1)) - 3.14159) * float32(3.14159)) + v1)
            var v49 float64 = float64(-2.25)
            _ = v49
        }
    case "a":
        for fi50 := 0; fi50 < 2; fi50++ {
            var v51 float64 = (3.14159 - (-2.25 * 0.0))
            var v52 uint64 = 4294967295
            v5 += (v5 + (v5 + -1))
            _ = fi50
            _ = v51
            _ = v52
        }
    }
    switch ("a" + (("\n\t" + ("a" + ("a" + "\n\t"))) + ((("a" + "foo bar") + "a") + "hello"))) {
    case "foo bar":
        var v53 uint16 = (v3 / ((v3) | 1))
        _ = v53
    case "hello":
        switch "foo bar" {
        case "hello":
            switch 127 {
            case -269:
                v54 := GenS8{F0: ("foo bar" + ("foo bar" + "a")), F1: v43}
                _ = v54
            case -9223372036854775808:
                if v4 {
                    var v55 uint16 = 2
                    switch 42 {
                    case 928:
                        var v56 uint = 207
                        _ = v56
                    case 964:
                        if (v4 && (v1 > v1)) {
                            for fi57 := 0; fi57 < 13; fi57++ {
                                v5 = v5
                                v4 = v4
                                _ = fi57
                            }
                        } else {
                            fmt.Println(v5, v55)
                            if true {
                                switch 32768 {
                                case 32768:
                                    v58 := []int{42, -9223372036854775808, 0}
                                    _ = v58
                                case 0:
                                    switch "hello" {
                                    case "foo bar":
                                        v1 -= (1e10 - ((-2.25 - v1) - v1))
                                    case "\n\t":
                                        for fi59 := 0; fi59 < 17; fi59++ {
                                            v60 := []int{-1, (fi59 / ((32768) | 1)), fi59, 959}
                                            _ = fi59
                                            _ = v60
                                        }
                                    case "":
                                        v2 -= uint8(v3)
                                    }
                                }
                                for fi61 := 0; fi61 < 3; fi61++ {
                                    var v62 float32 = v1
                                    v2 -= v2
                                    v43 -= (v43 | (v43 | (v43 + genFn0(((v43 - (v43 % (((-(v43))) | 1))) >= v43), (v2 * v2), -2147483648))))
                                    _ = fi61
                                    _ = v62
                                }
                            } else {
                                v63 := GenS7{F0: 42}
                                for fi64 := 0; fi64 < 8; fi64++ {
                                    v65 := []int{32768, -101}
                                    _ = fi64
                                    _ = v65
                                }
                                for fi66 := 0; fi66 < 15; fi66++ {
                                    v43 = 167
                                    _ = fi66
                                }
                                _ = v63
                            }
                        }
                    case 2147483647:
                        v4 = false
                    }
                    delete(v6, 564)
                    _ = v55
                }
            case 9223372036854775807:
                v1 = v1
            default:
                switch 4294967295 {
                case -26:
                    v67 := []int{2147483647, 1, -176}
                    _ = v67
                }
            }
        case "\n\t":
            fmt.Println(v43, v43)
        default:
            if (18446744073709551615 >= 0) {
                v68 := GenS1{F0: -1}
                if v4 {
                    for fi69 := 0; fi69 < 1; fi69++ {
                        if (((v43 / (((^(v43))) | 1)) >= v43) || ((fi69 + fi69) < -1)) {
                            if (867 <= 9223372036854775807) {
                                for fi70 := 0; fi70 < 4; fi70++ {
                                    switch (fi69 << ((4294967295) & 63)) {
                                    case 0:
                                        v71 := []int{fi70, 1}
                                        _ = v71
                                    }
                                    v1 -= v1
                                    var v72 float64 = 3.14159
                                    v73 := []int{}
                                    var v74 int = (fi69 % (((fi69 << ((46) & 63))) | 1))
                                    _ = fi70
                                    _ = v72
                                    _ = v73
                                    _ = v74
                                }
                            }
                            for fi75 := 0; fi75 < 12; fi75++ {
                                var v76 int = (^(fi75))
                                var v77 float32 = v1
                                fmt.Println(v4, fi75, v1)
                                _ = fi75
                                _ = v76
                                _ = v77
                            }
                            var v78 float64 = 1.5
                            _ = v78
                        } else {
                            var v79 uint = 579
                            fmt.Println(v4)
                            for fi80 := 0; fi80 < 19; fi80++ {
                                switch fi69 {
                                case 4294967295:
                                    var v81 uint16 = v3
                                    _ = v81
                                case 618:
                                    v4 = (v79 >= v79)
                                case -301:
                                    var v82 uint = v79
                                    _ = v82
                                }
                                switch ("\n\t" + "hello") {
                                case "hello":
                                    var v83 string = "a"
                                    _ = v83
                                case "a":
                                    switch int(v43) {
                                    case -9223372036854775808:
                                        switch fi80 {
                                        case 102:
                                            if false {
                                                v5 = (v5 ^ 0)
                                                v2 += v2
                                                switch (fi80 - fi80) {
                                                case 1:
                                                    v84 := GenS2{F0: v43}
                                                    _ = v84
                                                case 0:
                                                    if ((0.1 < v1) && ((v5 >> (((v79 / ((0) | 1))) & 63)) != (v5 | (v5 / ((127) | 1))))) {
                                                        if (!v4) {
                                                            switch -9223372036854775808 {
                                                            case 9223372036854775807:
                                                                fmt.Println(v79)
                                                            case 32768:
                                                                v1 += 1e10
                                                            }
                                                            fmt.Println(v43)
                                                        }
                                                    } else {
                                                        v3 += 530
                                                    }
                                                }
                                                if v4 {
                                                    var v85 string = "hello"
                                                    var v86 uint16 = (v3 | (v3 >> (((v79 - (v79 / ((v79) | 1)))) & 63)))
                                                    switch (fi80 >> (((^(v79))) & 63)) {
                                                    case 128:
                                                        v3 -= 954
                                                    case 127:
                                                        v87 := []int{}
                                                        _ = v87
                                                    case -984:
                                                        var v88 uint16 = v86
                                                        _ = v88
                                                    default:
                                                        fmt.Println(fi80, v2, fi69)
                                                    }
                                                    _ = v85
                                                    _ = v86
                                                }
                                            } else {
                                                v89 := GenS2{F0: 1000}
                                                var v90 uint32 = 102
                                                v3 = (v3 >> (((v79 >> ((255) & 63))) & 63))
                                                _ = v89
                                                _ = v90
                                            }
                                        case 0:
                                            if ((!v4) && false) {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    var v91 bool = v4
                                                    _ = v91
                                                }
                                            } else {
                                                if (true || (v3 != v3)) {
                                                    v2 += 220
                                                }
                                                switch (fi69 << ((v79) & 63)) {
                                                case 463:
                                                    var v92 uint = v79
                                                    _ = v92
                                                default:
                                                    var v93 uint = 112
                                                    if (4294967295 != 42) {
                                                        delete(v6, 1)
                                                        var v94 float32 = (v1 / (v1 * v1))
                                                        _ = v94
                                                    } else {
                                                        v95 := GenS3{F0: v79, F1: (^(v79)), F2: -2147483648}
                                                        v96 := []int{(fi80 ^ fi69), fi69, (fi69 & fi69)}
                                                        _ = v95
                                                        _ = v96
                                                    }
                                                    v97 := []int{(fi69 << ((4294967295) & 63)), 0, 32768}
                                                    _ = v93
                                                    _ = v97
                                                }
                                            }
                                        case 32767:
                                            v98 := GenS9{F0: 311, F1: int64(v43)}
                                            _ = v98
                                        default:
                                            v99 := GenS10{F0: 911, F1: -9223372036854775808}
                                            _ = v99
                                        }
                                    case 9223372036854775807:
                                        v100 := []int{(fi80 - fi69), fi80, -36}
                                        switch (fi69 | fi80) {
                                        case 668:
                                            if true {
                                                v100 = append(v100, int(v43))
                                                var v101 int32 = 1
                                                _ = len(v100)
                                                _ = v101
                                            }
                                        default:
                                            if ((-(v43)) > (-(v43))) {
                                                switch "a" {
                                                case "a":
                                                    if len(v100) >= 2 {
                                                        v100 = v100[1:len(v100)]
                                                    }
                                                default:
                                                    if (v4 || v4) {
                                                        var v102 uint32 = 4294967295
                                                        if v4 {
                                                            var v103 uint8 = (v2 / ((uint8(v79)) | 1))
                                                            _ = v103
                                                        }
                                                        _ = v102
                                                    }
                                                }
                                                if false {
                                                    v5 = (-(v5))
                                                    var v104 uint64 = 2
                                                    v43 = (v43 / ((v43) | 1))
                                                    _ = v104
                                                } else {
                                                    v100 = append(v100, 9223372036854775807)
                                                    v5 = int8(v3)
                                                }
                                                fmt.Println(v2, fi80, v79)
                                                var v105 int8 = v5
                                                var v106 uint16 = (v3 * v3)
                                                _ = v105
                                                _ = v106
                                            }
                                        }
                                        if ((v2 % ((v2) | 1)) > v2) {
                                            var v107 int64 = 1
                                            switch fi80 {
                                            case 128:
                                                fmt.Println(v79)
                                            case 2147483647:
                                                v1 = float32(1e10)
                                            default:
                                                if false {
                                                    if v4 {
                                                        var v108 uint8 = (v2 & 231)
                                                        var v109 int = (-(fi69))
                                                        delete(v6, 42)
                                                        v110 := GenS11{F0: -1, F1: 255, F2: (^(v108))}
                                                        var v111 int = (fi80 / ((1) | 1))
                                                        _ = v108
                                                        _ = v109
                                                        _ = v110
                                                        _ = v111
                                                    }
                                                    v112 := GenS12{F0: ("a" + "foo bar"), F1: (("\n\t" + (("a" + "\n\t") + "\n\t")) + "foo bar"), F2: v79}
                                                    _ = v112
                                                } else {
                                                    if false {
                                                        v1 += v1
                                                        if v4 {
                                                            switch (-(fi69)) {
                                                            case -228:
                                                                _ = len(v6)
                                                            case -698:
                                                                v5 += (^(v5))
                                                                var v113 int64 = v107
                                                                switch fi80 {
                                                                case 32768:
                                                                    var v114 uint8 = (v2 | uint8(v3))
                                                                    _ = v114
                                                                default:
                                                                    var v115 float64 = float64((v1 - (v1 - (1.5 * (((0.0 / v1) + ((1e10 * 1.5) * v1)) / v1)))))
                                                                    _ = v115
                                                                }
                                                                _ = v113
                                                            case 2147483647:
                                                                fmt.Println(fi80)
                                                            default:
                                                                var v116 uint = (v79 * (v79 - v79))
                                                                _ = v116
                                                            }
                                                            v6[804] = 42
                                                        } else {
                                                            v3 -= v3
                                                            var v117 uint16 = uint16(v2)
                                                            _ = v117
                                                        }
                                                        _ = len(v6)
                                                        fmt.Println(v4, fi80, v4)
                                                    }
                                                    if (((-2.25 - float32(1e10)) == (v1 - v1)) && (false && (v4 || v4))) {
                                                        if (v4 || (!false)) {
                                                            var v118 uint8 = 152
                                                            v119 := GenS10{F0: 354, F1: (^(fi69))}
                                                            _ = v118
                                                            _ = v119
                                                        } else {
                                                            var v120 int64 = (v107 | 42)
                                                            var v121 uint = 2
                                                            if ((fi69 | fi80) >= -14) {
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    if false {
                                                                        switch "a" {
                                                                        case "foo bar":
                                                                            if (v43 != v43) {
                                                                                v122 := GenS1{F0: v120}
                                                                                v107 = (^(v107))
                                                                                _ = v122
                                                                            } else {
                                                                                if v4 {
                                                                                    if (527 <= 494) {
                                                                                        v123 := GenS9{F0: (v3 / ((uint16(v2)) | 1)), F1: (v107 % ((0) | 1))}
                                                                                        _ = v123
                                                                                    }
                                                                                    if true {
                                                                                        v124 := GenS5{F0: (-(v5))}
                                                                                        v4 = true
                                                                                        _ = v124
                                                                                    } else {
                                                                                        fmt.Println(v120)
                                                                                        var v125 float32 = float32(3.14159)
                                                                                        _ = v125
                                                                                    }
                                                                                }
                                                                            }
                                                                        case "\n\t":
                                                                            var v126 string = "hello"
                                                                            _ = v126
                                                                        case "":
                                                                            v120 += (v107 >> ((v79) & 63))
                                                                        }
                                                                    }
                                                                case "hello":
                                                                    fmt.Println(v2)
                                                                }
                                                            }
                                                            _ = v120
                                                            _ = v121
                                                        }
                                                        fmt.Println(v107, v1, fi80)
                                                        switch fi80 {
                                                        case -9223372036854775808:
                                                            v1 += v1
                                                        case 42:
                                                            v4 = (!(!v4))
                                                        case 0:
                                                            v2 = v2
                                                            if ((v2 ^ (v2 ^ v2)) == uint8(v79)) {
                                                                var v127 int = int(v5)
                                                                switch v127 {
                                                                case -2147483648:
                                                                    var v128 int64 = (v107 | -1)
                                                                    var v129 uint8 = (^(v2))
                                                                    _ = v128
                                                                    _ = v129
                                                                case 9223372036854775807:
                                                                    switch 188 {
                                                                    case -525:
                                                                        v100 = append(v100, fi69)
                                                                        switch 368 {
                                                                        case 281:
                                                                            v130 := GenS13{F0: v5, F1: (v79 & v79), F2: -1}
                                                                            _ = v130
                                                                        case 32768:
                                                                            v79 = v79
                                                                            var v131 int = fi80
                                                                            _ = v131
                                                                        case 127:
                                                                            v79 = (v79 | v79)
                                                                        }
                                                                        switch fi69 {
                                                                        case -1:
                                                                            if v4 {
                                                                                var v132 int16 = -74
                                                                                fmt.Println(v43, fi80)
                                                                                _ = v132
                                                                            } else {
                                                                                switch fi69 {
                                                                                case 4294967295:
                                                                                    switch (fi80 & fi80) {
                                                                                    case 2147483647:
                                                                                        var v133 int64 = (v107 << ((2) & 63))
                                                                                        _ = v133
                                                                                    case 4294967295:
                                                                                        v3 = v3
                                                                                    case 128:
                                                                                        _ = len(v6)
                                                                                    }
                                                                                default:
                                                                                    var v134 float32 = (v1 / 3.14159)
                                                                                    _ = v134
                                                                                }
                                                                                var v135 int8 = (-(v5))
                                                                                var v136 bool = ((0 < 4294967295) && v4)
                                                                                v136 = (0 >= 180)
                                                                                v127 = int(v2)
                                                                                v137 := v6[2147483647]
                                                                                _ = v135
                                                                                _ = v136
                                                                                _ = v137
                                                                            }
                                                                        }
                                                                    case 32767:
                                                                        delete(v6, 295)
                                                                    }
                                                                default:
                                                                    v138 := v6[2147483647]
                                                                    _ = v138
                                                                }
                                                                _ = v127
                                                            }
                                                            if true {
                                                                switch "" {
                                                                case "hello":
                                                                    switch (fi69 + fi69) {
                                                                    case 1:
                                                                        switch "hello" {
                                                                        case "a":
                                                                            var v139 float64 = 0.1
                                                                            _ = v139
                                                                        case "hello":
                                                                            if len(v100) > 0 {
                                                                                v100[0] = fi69
                                                                            }
                                                                        }
                                                                    case 325:
                                                                        v2 -= v2
                                                                    default:
                                                                        var v140 bool = v4
                                                                        _ = v140
                                                                    }
                                                                    if (!false) {
                                                                        v141 := GenS14{F0: (-(v107)), F1: (v107 << (((^(v79))) & 63)), F2: -2147483648}
                                                                        var v142 int = fi80
                                                                        v143 := GenS15{F0: 221}
                                                                        _ = v141
                                                                        _ = v142
                                                                        _ = v143
                                                                    }
                                                                case "a":
                                                                    var v144 float32 = ((0.1 / v1) * v1)
                                                                    _ = v144
                                                                case "foo bar":
                                                                    switch fi80 {
                                                                    case -655:
                                                                        var v145 uint32 = 2147483648
                                                                        _ = v145
                                                                    case 42:
                                                                        var v146 int = fi80
                                                                        _ = v146
                                                                    default:
                                                                        v79 = v79
                                                                    }
                                                                }
                                                                v147 := GenS16{F0: (v43 / ((int16(v79)) | 1)), F1: int8(fi80), F2: fi80}
                                                                _ = v147
                                                            }
                                                        }
                                                        v6[uint32(v3)] = 912
                                                    }
                                                    v4 = (v2 != (^(v2)))
                                                }
                                            }
                                            _ = v107
                                        } else {
                                            var v148 int32 = -1
                                            var v149 int16 = (v43 % ((-354) | 1))
                                            _ = v148
                                            _ = v149
                                        }
                                        _ = v100
                                    case 159:
                                        v150 := []int{fi69, (fi69 >> ((65535) & 63)), fi69}
                                        _ = v150
                                    }
                                case "\n\t":
                                    var v151 uint32 = 65535
                                    _ = v151
                                default:
                                    v152 := []int{}
                                    _ = v152
                                }
                                _ = fi80
                            }
                            _ = v79
                        }
                        var v153 int16 = (v43 * int16(v5))
                        switch fi69 {
                        case -2147483648:
                            v153 = (v153 << ((42) & 63))
                        case 32768:
                            switch 32768 {
                            case 4294967295:
                                fmt.Println(v5, v5, v153)
                            case 127:
                                fmt.Println(v5, v4, v153)
                            default:
                                v154 := GenS17{F0: v4, F1: "\n\t", F2: v43}
                                _ = v154
                            }
                        }
                        _ = fi69
                        _ = v153
                    }
                } else {
                    v155 := []int{46}
                    var v156 string = "a"
                    _ = v155
                    _ = v156
                }
                _ = v68
            } else {
                fmt.Println(v43)
                var v157 float64 = 1.5
                _ = v157
            }
        }
    case "":
        delete(v6, 2)
    }
    var v158 uint32 = 385
    fmt.Println("v1:", v1)
    fmt.Println("v158:", v158)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v158
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
