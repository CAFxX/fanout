package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 uint64; F2 uint32 }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 int; F1 int16; F2 uint16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int32; F1 uint64 }
type GenS11 struct { F0 bool; F1 int32; F2 uint8 }
type GenS12 struct { F0 int16; F1 bool }
type GenS13 struct { F0 uint16; F1 bool; F2 int64 }
type GenS14 struct { F0 int; F1 string; F2 int32 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 32767
    var v3 uint16 = 0
    var v4 int16 = -1
    var v5 int64 = 1
    switch 32767 {
    case -2147483648:
        v6 := GenS1{F0: 42}
        _ = v6
    case -896:
        fmt.Println(v2, v4)
    default:
        var v7 float64 = 1.5
        _ = v7
    }
    var v8 uint8 = 2
    var v9 uint = 18446744073709551615
    v9 = v9
    switch ("hello" + "\n\t") {
    case "hello":
        var v10 string = ("foo bar" + ("\n\t" + "a"))
        _ = v10
    case "foo bar":
        v11 := []int{}
        _ = v11
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        switch 128 {
        case 127:
            var v13 int = fi12
            _ = v13
        case 116:
            for fi14 := 0; fi14 < 14; fi14++ {
                fmt.Println(v2, fi12)
                _ = fi14
            }
        case -9223372036854775808:
            if (!(!(float32((-2.25 / 1.0)) != ((1e10 / 1.0) + (float32(0.0) * (1.5 + (1.5 / float32(3.14159)))))))) {
                for fi15 := 0; fi15 < 12; fi15++ {
                    var v16 uint32 = uint32(v9)
                    v16 = (v16 - v16)
                    fmt.Println(fi12, v9, v4)
                    _ = fi15
                    _ = v16
                }
                v9 = v9
            } else {
                v17 := v1[""]
                for fi18 := 0; fi18 < 3; fi18++ {
                    var v19 bool = (!(fi18 > (^(fi18))))
                    _ = fi18
                    _ = v19
                }
                _ = v17
            }
        default:
            if true {
                if (130 > 0) {
                    var v20 int64 = (v5 & (v5 ^ v5))
                    v21 := []int{}
                    v4 += v4
                    _ = v20
                    _ = v21
                } else {
                    v22 := GenS1{F0: 18446744073709551615}
                    fmt.Println(v8, fi12, v8)
                    if false {
                        _ = len(v1)
                        if true {
                            if ((v3 / ((32767) | 1)) < 0) {
                                fmt.Println(v2, v5, v8)
                                fmt.Println(fi12, v8)
                            }
                            v9 += (v9 ^ v9)
                        }
                        var v23 int16 = v4
                        _ = v23
                    }
                    _ = v22
                }
                v3 = (v3 >> (((v9 % (((v9 % ((v9) | 1))) | 1))) & 63))
            }
        }
        for fi24 := 0; fi24 < 8; fi24++ {
            if (4294967295 > (v9 % (((v9 | v9)) | 1))) {
                if (v8 == v8) {
                    switch ("foo bar" + "\n\t") {
                    case "":
                        switch (fi12 % (((fi12 % (((-(fi24))) | 1))) | 1)) {
                        case 1:
                            fmt.Println(v4, v9, v9)
                        case 4294967295:
                            switch fi12 {
                            case -685:
                                v8 += v8
                            default:
                                if false {
                                    if (v4 != 42) {
                                        switch "foo bar" {
                                        case "foo bar":
                                            var v25 int32 = 165
                                            _ = v25
                                        case "\n\t":
                                            var v26 uint = v9
                                            _ = v26
                                        case "a":
                                            v27 := GenS1{F0: 806}
                                            _ = v27
                                        }
                                        fmt.Println(v2, v5, v3)
                                        v28 := []int{0, (fi24 | -1), fi24}
                                        _ = v28
                                    }
                                    switch (fi12 + fi24) {
                                    case 749:
                                        var v29 float64 = 0.0
                                        _ = v29
                                    case -348:
                                        v5 = (v5 - v5)
                                    }
                                    if (-2.25 >= 1e10) {
                                        v30 := GenS1{F0: 91}
                                        v4 = (v4 << (((v9 / ((965) | 1))) & 63))
                                        if false {
                                            var v31 uint32 = uint32(v8)
                                            if (!(int32(v3) == -2147483648)) {
                                                v32 := GenS1{F0: 1}
                                                fmt.Println(v4)
                                                var v33 int32 = 756
                                                v34 := GenS1{F0: 1}
                                                _ = v32
                                                _ = v33
                                                _ = v34
                                            } else {
                                                fmt.Println(v2)
                                                if true {
                                                    var v35 uint8 = (v8 + 255)
                                                    _ = v35
                                                } else {
                                                    v36 := GenS2{F0: (v9 + 18446744073709551615)}
                                                    v37 := []int{(fi24 * -9223372036854775808), (fi24 | fi24)}
                                                    if false {
                                                        var v38 int8 = 0
                                                        _ = len(v1)
                                                        fmt.Println(fi12)
                                                        _ = v38
                                                    } else {
                                                        v39 := GenS2{F0: 18446744073709551615}
                                                        switch (fi12 - -682) {
                                                        case 44:
                                                            switch fi24 {
                                                            case 127:
                                                                var v40 uint = uint(v3)
                                                                _ = v40
                                                            case 2147483647:
                                                                fmt.Println(v8, v31, v31)
                                                            case 32768:
                                                                v3 -= v2
                                                            default:
                                                                _ = len(v37)
                                                                v37 = append(v37, (fi12 << ((15) & 63)))
                                                            }
                                                        case -2147483648:
                                                            v41 := v1["\n\t"]
                                                            _ = v41
                                                        case 9223372036854775807:
                                                            var v42 int64 = v5
                                                            _ = v42
                                                        default:
                                                            var v43 int64 = -185
                                                            _ = v43
                                                        }
                                                        switch fi12 {
                                                        case -9223372036854775808:
                                                            v8 -= v8
                                                        case 128:
                                                            var v44 float64 = ((1.5 / 1.5) + 0.0)
                                                            _ = v44
                                                        case 32768:
                                                            v3 -= 32767
                                                        default:
                                                            if (1 >= 157) {
                                                                if true {
                                                                    if len(v37) > 0 {
                                                                        v37[0] = fi24
                                                                    }
                                                                } else {
                                                                    v45 := GenS3{F0: (v9 - uint(v31)), F1: 1, F2: uint32(v8)}
                                                                    v9 = v9
                                                                    _ = v45
                                                                }
                                                                _ = len(v37)
                                                            }
                                                        }
                                                        v46 := GenS4{F0: (v9 % ((v9) | 1)), F1: 452}
                                                        v47 := GenS1{F0: 4294967295}
                                                        _ = v39
                                                        _ = v46
                                                        _ = v47
                                                    }
                                                    _ = v36
                                                    _ = v37
                                                }
                                                fmt.Println(v3, fi12, v3)
                                            }
                                            var v48 uint8 = 31
                                            var v49 int16 = (v4 - v4)
                                            _ = v31
                                            _ = v48
                                            _ = v49
                                        } else {
                                            v1[""] = false
                                            fmt.Println(fi24)
                                            v2 += (^(v3))
                                        }
                                        _ = v30
                                    } else {
                                        var v50 string = ""
                                        _ = v50
                                    }
                                    fmt.Println(v3)
                                    switch "hello" {
                                    case "foo bar":
                                        var v51 uint8 = (^(v8))
                                        _ = v51
                                    case "a":
                                        v52 := GenS2{F0: v9}
                                        _ = v52
                                    default:
                                        if ("hello" == "\n\t") {
                                            switch -9223372036854775808 {
                                            case 72:
                                                v53 := GenS3{F0: v9, F1: 0, F2: 42}
                                                _ = v53
                                            case 0:
                                                v9 -= (v9 * v9)
                                            }
                                            var v54 float32 = 1e10
                                            v55 := GenS5{F0: (fi24 | fi12), F1: int16(v3), F2: v3}
                                            _ = v54
                                            _ = v55
                                        } else {
                                            v5 = (v5 / ((v5) | 1))
                                            if (!(uint(v8) == v9)) {
                                                v56 := []int{(fi12 + (fi24 / ((fi24) | 1))), 1, -40}
                                                if (((false || ((!(42 >= -2147483648)) && false)) && true) && false) {
                                                    if len(v56) >= 2 {
                                                        v56 = v56[1:len(v56)]
                                                    }
                                                    var v57 int64 = v5
                                                    var v58 int32 = -1
                                                    fmt.Println(v57, v5)
                                                    var v59 float32 = ((3.14159 / 1.0) - (3.14159 + 3.14159))
                                                    var v60 float32 = v59
                                                    _ = v57
                                                    _ = v58
                                                    _ = v59
                                                    _ = v60
                                                } else {
                                                    v9 = 18446744073709551615
                                                }
                                                if ((247 > 2147483648) && ((!true) || (780 > 944))) {
                                                    v61 := GenS3{F0: 1, F1: 471, F2: 1}
                                                    _ = v61
                                                } else {
                                                    var v62 uint64 = 255
                                                    v4 = (v4 << (((v9 & v9)) & 63))
                                                    v4 -= (v4 + (v4 % (((v4 - v4)) | 1)))
                                                    _ = v62
                                                }
                                                _ = v56
                                            } else {
                                                v9 = (v9 | 1)
                                                v2 -= uint16(v8)
                                                if (((^(v2)) < (v2 % ((1) | 1))) || true) {
                                                    if (0.0 >= (0.0 * -2.25)) {
                                                        fmt.Println(v8, v5, v5)
                                                    } else {
                                                        switch "foo bar" {
                                                        case "":
                                                            var v63 float32 = 0.0
                                                            _ = v63
                                                        default:
                                                            v9 -= v9
                                                        }
                                                    }
                                                    var v64 int = (fi24 >> ((302) & 63))
                                                    _ = v64
                                                } else {
                                                    v9 = (v9 % ((v9) | 1))
                                                }
                                                v4 += 55
                                                var v65 int32 = 742
                                                _ = v65
                                            }
                                            if true {
                                                switch ("" + ("foo bar" + ("foo bar" + "foo bar"))) {
                                                case "":
                                                    if ((2 >= 65535) || (42 >= -63)) {
                                                        v2 -= 255
                                                    } else {
                                                        delete(v1, "a")
                                                        if false {
                                                            var v66 int32 = -884
                                                            v66 += (v66 << ((909) & 63))
                                                            _ = v66
                                                        }
                                                        var v67 uint = 18446744073709551615
                                                        _ = v67
                                                    }
                                                    v5 = v5
                                                case "\n\t":
                                                    v68 := []int{(fi12 / ((fi24) | 1)), (-(fi24)), (fi24 - fi12)}
                                                    _ = v68
                                                default:
                                                    v8 = v8
                                                    v69 := GenS2{F0: (v9 | 65535)}
                                                    _ = v69
                                                }
                                            } else {
                                                var v70 float32 = 0.0
                                                if false {
                                                    v71 := GenS5{F0: 42, F1: v4, F2: (v3 * v2)}
                                                    var v72 int16 = 32767
                                                    v5 = v5
                                                    _ = v71
                                                    _ = v72
                                                }
                                                _ = v70
                                            }
                                        }
                                    }
                                    var v73 int32 = -511
                                    _ = v73
                                }
                            }
                        case -781:
                            if true {
                                switch "\n\t" {
                                case "hello":
                                    if false {
                                        var v74 float64 = 0.0
                                        delete(v1, "")
                                        _ = v74
                                    } else {
                                        if ((!("\n\t" >= "\n\t")) && ((fi12 - (fi24 / (((fi12 % ((fi24) | 1))) | 1))) != (fi12 % (((fi24 | (fi24 + 9223372036854775807))) | 1)))) {
                                            _ = len(v1)
                                        } else {
                                            if false {
                                                v75 := []int{}
                                                v8 += 0
                                                _ = v75
                                            } else {
                                                var v76 uint = v9
                                                _ = v76
                                            }
                                            var v77 int16 = v4
                                            _ = v77
                                        }
                                        var v78 int64 = (^(v5))
                                        _ = v78
                                    }
                                case "\n\t":
                                    fmt.Println(v2)
                                }
                                v5 -= (^(v5))
                                v79 := GenS3{F0: (v9 >> ((984) & 63)), F1: 790, F2: 0}
                                _ = v79
                            } else {
                                if false {
                                    v80 := GenS5{F0: (fi12 % ((fi12) | 1)), F1: (^(v4)), F2: v3}
                                    switch (fi12 | fi12) {
                                    case -40:
                                        if true {
                                            var v81 uint16 = v3
                                            _ = v81
                                        } else {
                                            fmt.Println(v3, v2)
                                            fmt.Println(fi12)
                                            fmt.Println(fi24)
                                        }
                                    }
                                    _ = v80
                                }
                                fmt.Println(v9)
                            }
                        }
                    default:
                        if (-2147483648 <= 0) {
                            fmt.Println(fi24)
                        }
                    }
                    var v82 uint8 = uint8(v2)
                    switch "" {
                    case "foo bar":
                        if false {
                            var v83 int64 = (v5 * (v5 * (v5 << ((65535) & 63))))
                            if true {
                                var v84 float64 = -2.25
                                v2 = v2
                                _ = v84
                            } else {
                                switch (-(fi12)) {
                                case -2147483648:
                                    if (0 != 597) {
                                        v8 -= v8
                                        v83 = v5
                                        v3 -= (v2 + 42)
                                    }
                                case 24:
                                    switch (fi12 * fi24) {
                                    case 42:
                                        var v85 int64 = v83
                                        _ = v85
                                    case 127:
                                        if true {
                                            if false {
                                                v86 := GenS4{F0: 634, F1: -2147483648}
                                                _ = v86
                                            } else {
                                                var v87 uint16 = uint16(v82)
                                                _ = v87
                                            }
                                        }
                                    default:
                                        v5 -= 9223372036854775807
                                    }
                                }
                                if (fi24 <= (fi24 + fi12)) {
                                    v88 := GenS6{F0: (^(v9))}
                                    _ = v88
                                }
                                v82 = v8
                            }
                            _ = v83
                        } else {
                            v89 := GenS3{F0: v9, F1: 803, F2: 560}
                            _ = v89
                        }
                        var v90 int = fi24
                        if (v9 != v9) {
                            if ((v2 & v3) == 65535) {
                                v91 := GenS7{F0: uint32(v3)}
                                v92 := v1[("" + (("" + ("" + "")) + (("\n\t" + "\n\t") + "hello")))]
                                var v93 uint64 = 674
                                var v94 uint64 = uint64(v82)
                                _ = v91
                                _ = v92
                                _ = v93
                                _ = v94
                            }
                            var v95 int32 = int32(v9)
                            var v96 uint64 = 776
                            _ = v95
                            _ = v96
                        }
                        if false {
                            v97 := GenS3{F0: 255, F1: 2, F2: 4294967295}
                            switch "" {
                            case "foo bar":
                                if false {
                                    var v98 string = "foo bar"
                                    _ = v98
                                }
                            case "\n\t":
                                var v99 float32 = 3.14159
                                _ = v99
                            case "hello":
                                v90 -= (fi12 << (((v9 + 4294967295)) & 63))
                            default:
                                if true {
                                    v90 = fi12
                                }
                            }
                            switch "" {
                            case "\n\t":
                                fmt.Println(v9)
                            }
                            v100 := GenS7{F0: 2}
                            if true {
                                fmt.Println(v5)
                            } else {
                                v4 = 1000
                                var v101 int = -9223372036854775808
                                _ = len(v1)
                                _ = v101
                            }
                            v102 := []int{fi24}
                            _ = v100
                            _ = v102
                            _ = v97
                        }
                        _ = v90
                    }
                    _ = v82
                }
                v103 := []int{(fi24 << (((v9 << ((v9) & 63))) & 63)), int(v4)}
                _ = v103
            } else {
                switch 670 {
                case -1:
                    v104 := []int{}
                    v8 = 2
                    fmt.Println(v2)
                    _ = v104
                }
            }
            v4 = v4
            _ = fi24
        }
        _ = fi12
    }
    if true {
        switch -422 {
        case 705:
            if false {
                var v105 float64 = 3.14159
                var v106 int8 = 31
                for fi107 := 0; fi107 < 1; fi107++ {
                    for fi108 := 0; fi108 < 5; fi108++ {
                        v106 += v106
                        var v109 uint16 = (v2 ^ (v3 % ((uint16(v8)) | 1)))
                        var v110 uint8 = v8
                        _ = fi108
                        _ = v109
                        _ = v110
                    }
                    if ((!false) || (((!false) && (false || (!(!true)))) || true)) {
                        if true {
                            if (v3 != 255) {
                                v111 := []int{(fi107 << (((v9 + v9)) & 63)), fi107, (fi107 - fi107)}
                                v112 := GenS8{F0: ((("hello" + "a") + "") + ("hello" + ("\n\t" + ("a" + "a"))))}
                                _ = v111
                                _ = v112
                            }
                            fmt.Println(v4, v5, v5)
                            switch (fi107 / (((fi107 % (((^(fi107))) | 1))) | 1)) {
                            case 4294967295:
                                v9 = (v9 % ((762) | 1))
                            }
                        }
                    } else {
                        v113 := GenS9{F0: -32768}
                        v114 := GenS10{F0: 42, F1: 18446744073709551615}
                        _ = v113
                        _ = v114
                    }
                    v5 = 9223372036854775807
                    _ = fi107
                }
                _ = v105
                _ = v106
            } else {
                if true {
                    var v115 int64 = (-(v5))
                    for fi116 := 0; fi116 < 9; fi116++ {
                        switch "\n\t" {
                        case "\n\t":
                            for fi117 := 0; fi117 < 14; fi117++ {
                                v4 = int16(v8)
                                fmt.Println(v3, fi117)
                                _ = fi117
                            }
                        }
                        for fi118 := 0; fi118 < 9; fi118++ {
                            var v119 int = (fi118 - fi118)
                            v3 -= 637
                            _ = fi118
                            _ = v119
                        }
                        for fi120 := 0; fi120 < 13; fi120++ {
                            fmt.Println(fi116)
                            if true {
                                v4 = v4
                                var v121 float64 = float64(1.5)
                                v3 -= 255
                                _ = v121
                            }
                            v115 = (v115 << ((v9) & 63))
                            v3 = v3
                            _ = fi120
                        }
                        _ = fi116
                    }
                    var v122 float32 = float32(1.5)
                    _ = v115
                    _ = v122
                }
            }
        case 977:
            v2 = (^(v3))
        case 771:
            var v123 int16 = (v4 - v4)
            _ = v123
        default:
            if false {
                var v124 uint32 = 42
                v125 := GenS5{F0: -9223372036854775808, F1: (v4 + 42), F2: v3}
                _ = v124
                _ = v125
            } else {
                if ((4294967295 == 921) && true) {
                    v126 := v1["foo bar"]
                    v127 := []int{}
                    for fi128 := 0; fi128 < 12; fi128++ {
                        var v129 int64 = (v5 * v5)
                        fmt.Println(v8)
                        if v126 {
                            v129 = 9223372036854775807
                        }
                        _ = fi128
                        _ = v129
                    }
                    _ = v126
                    _ = v127
                } else {
                    if ((true || ("foo bar" >= "")) || (true || true)) {
                        for fi130 := 0; fi130 < 3; fi130++ {
                            var v131 bool = (("a" + "foo bar") == (("foo bar" + "foo bar") + ("\n\t" + ("" + ("\n\t" + "")))))
                            switch ((("" + "hello") + ("\n\t" + ("foo bar" + (("" + "foo bar") + "foo bar")))) + "foo bar") {
                            case "a":
                                if v131 {
                                    v132 := GenS3{F0: 2, F1: 18446744073709551615, F2: 2147483647}
                                    switch (fi130 | (fi130 | (fi130 << ((v9) & 63)))) {
                                    case 32767:
                                        if false {
                                            switch fi130 {
                                            case 296:
                                                v133 := []int{(fi130 - (-(fi130))), 127}
                                                _ = v133
                                            case 42:
                                                v8 -= v8
                                            }
                                            fmt.Println(v4, v2, v8)
                                            v5 = v5
                                        }
                                    default:
                                        var v134 int32 = -2147483648
                                        _ = v134
                                    }
                                    _ = v132
                                }
                            }
                            var v135 uint64 = 18446744073709551615
                            _ = fi130
                            _ = v131
                            _ = v135
                        }
                        v136 := GenS11{F0: false, F1: int32(v8), F2: 0}
                        v137 := []int{-982, -9223372036854775808, 0}
                        _ = v136
                        _ = v137
                    }
                    var v138 string = ("hello" + "a")
                    _ = v138
                }
                v139 := GenS12{F0: -137, F1: true}
                v140 := []int{1, 32767}
                _ = v139
                _ = v140
            }
        }
        fmt.Println(v3)
        var v141 float32 = 0.1
        _ = v141
    } else {
        for fi142 := 0; fi142 < 2; fi142++ {
            for fi143 := 0; fi143 < 18; fi143++ {
                if false {
                    v144 := GenS13{F0: v2, F1: ((2 != 0) || (197 > (v5 - v5))), F2: (-(v5))}
                    v3 -= v3
                    if false {
                        if (-115 != 100) {
                            v9 += v9
                            fmt.Println(fi143, v2, v5)
                        } else {
                            delete(v1, "hello")
                            v8 += (^(v8))
                            v2 = (v2 + v3)
                        }
                        var v145 uint64 = 18446744073709551615
                        _ = v145
                    }
                    _ = v144
                } else {
                    v4 += (v4 | 405)
                }
                v1["foo bar"] = ((v9 | v9) >= v9)
                v146 := GenS14{F0: -1, F1: ("a" + "\n\t"), F2: -2147483648}
                _ = fi143
                _ = v146
            }
            fmt.Println(v2, v8)
            if true {
                for fi147 := 0; fi147 < 15; fi147++ {
                    var v148 uint8 = 128
                    fmt.Println(v3)
                    _ = fi147
                    _ = v148
                }
                v8 += v8
                v5 += v5
            }
            _ = fi142
        }
        v149 := GenS11{F0: (!false), F1: -2147483648, F2: v8}
        _ = v149
    }
    var v150 int = 505
    for fi151 := 0; fi151 < 15; fi151++ {
        v152 := GenS7{F0: 2147483648}
        switch ("foo bar" + ((("foo bar" + (("hello" + "hello") + ((("\n\t" + "\n\t") + ("\n\t" + "a")) + "a"))) + "\n\t") + (("" + "") + ("a" + "foo bar")))) {
        case "":
            fmt.Println(v2)
        case "\n\t":
            v153 := GenS8{F0: "foo bar"}
            _ = v153
        case "hello":
            switch 820 {
            case 1:
                var v154 float32 = 1e10
                _ = v154
            }
        }
        var v155 float64 = float64((float32(3.14159) + 0.0))
        _ = fi151
        _ = v152
        _ = v155
    }
    v156 := v1["hello"]
    v157 := GenS4{F0: 640, F1: -2147483648}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v150:", v150)
    fmt.Println("v156:", v156)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v150
    _ = v156
    _ = v157
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
