package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS5 struct { F0 string; F1 int32; F2 int8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int64; F1 int32 }
type GenS8 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS9 struct { F0 uint64; F1 int8 }
type GenS10 struct { F0 uint8; F1 uint8; F2 int64 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint; F1 uint64; F2 bool }
type GenS13 struct { F0 uint; F1 int32 }

func main() {
    var v1 uint64 = 4294967295
    var v2 int64 = 9223372036854775807
    v3 := make(map[string]uint32)
    v4 := make(map[string]bool)
    v5 := make(map[uint32]int)
    switch "hello" {
    case "\n\t":
        delete(v5, 2)
    case "hello":
        var v6 string = "a"
        _ = v6
    case "":
        var v7 uint16 = 628
        _ = v7
    default:
        if ((42 >= 42) || true) {
            v2 += -1
        } else {
            switch 32768 {
            case -367:
                if (-380 > -9223372036854775808) {
                    for fi8 := 0; fi8 < 18; fi8++ {
                        var v9 int32 = 42
                        _ = fi8
                        _ = v9
                    }
                    v1 = (v1 - v1)
                    v10 := []int{32767, 9223372036854775807}
                    _ = v10
                }
            default:
                fmt.Println(v2, v1)
            }
            v2 = 9223372036854775807
            fmt.Println(v1)
        }
    }
    v11 := GenS1{F0: 42, F1: 2147483647}
    for fi12 := 0; fi12 < 8; fi12++ {
        if false {
            v1 = 946
            if (-523 >= 401) {
                fmt.Println(v2, v2, v2)
            } else {
                v1 += v1
            }
            _ = len(v5)
        } else {
            var v13 uint32 = 955
            var v14 float32 = (((float32(1e10) + (((0.1 * -2.25) / (float32((3.14159 - (3.14159 * (((3.14159 - -2.25) - 3.14159) + 1.5)))) + float32(1e10))) - (0.1 + float32(1e10)))) + 3.14159) - (1e10 + 1e10))
            _ = v13
            _ = v14
        }
        switch 128 {
        case 42:
            if (1e10 >= 1e10) {
                if (v1 != v1) {
                    v15 := GenS2{F0: fi12}
                    fmt.Println(fi12)
                    v16 := GenS2{F0: 554}
                    _ = v15
                    _ = v16
                } else {
                    var v17 uint64 = v1
                    v18 := v5[2147483648]
                    _ = v17
                    _ = v18
                }
                v19 := []int{4294967295, fi12, (fi12 | fi12), (fi12 & fi12)}
                _ = v19
            }
        case 2147483647:
            var v20 int32 = -1
            _ = v20
        case 32768:
            var v21 uint32 = 4294967295
            _ = v21
        }
        v1 -= v1
        _ = fi12
    }
    v22 := GenS1{F0: 2147483648, F1: -904}
    var v23 uint8 = 127
    var v24 string = "foo bar"
    v2 = (v2 ^ (v2 & (-(v2))))
    for fi25 := 0; fi25 < 19; fi25++ {
        fmt.Println(v23, v24)
        for fi26 := 0; fi26 < 11; fi26++ {
            if (!true) {
                v27 := GenS3{F0: 617}
                v28 := GenS3{F0: 65535}
                var v29 uint32 = 2147483648
                _ = v27
                _ = v28
                _ = v29
            } else {
                switch 32768 {
                case 9223372036854775807:
                    if (((v24 + "a") + (v24 + ((("hello" + "foo bar") + v24) + (v24 + "foo bar")))) != v24) {
                        if (-2147483648 >= 0) {
                            v2 += v2
                        } else {
                            var v30 uint8 = (v23 << ((13) & 63))
                            fmt.Println(fi26, v1)
                            if false {
                                switch v24 {
                                case "hello":
                                    v1 = (v1 << ((18446744073709551615) & 63))
                                    switch (fi26 % ((32767) | 1)) {
                                    case 0:
                                        v31 := []int{565, (fi26 * fi26)}
                                        _ = v31
                                    case 4294967295:
                                        var v32 int16 = -38
                                        _ = v32
                                    default:
                                        fmt.Println(v2)
                                    }
                                }
                                var v33 int16 = 743
                                if true {
                                    v34 := GenS2{F0: -538}
                                    _ = v34
                                }
                                _ = v33
                            } else {
                                fmt.Println(v24, fi26, v30)
                                switch "" {
                                case "":
                                    var v35 float32 = 0.0
                                    _ = v35
                                case "foo bar":
                                    v36 := GenS1{F0: 2, F1: 0}
                                    _ = v36
                                case "a":
                                    var v37 string = (v24 + "\n\t")
                                    _ = v37
                                }
                            }
                            v1 = v1
                            _ = v30
                        }
                    }
                case -141:
                    switch ("foo bar" + v24) {
                    case "\n\t":
                        switch fi26 {
                        case 0:
                            delete(v5, uint32(v1))
                        case 128:
                            delete(v3, v24)
                            var v38 bool = ((v1 + v1) >= v1)
                            _ = v38
                        case -166:
                            v39 := []int{32768, fi25, 128}
                            _ = v39
                        }
                    case "a":
                        var v40 string = v24
                        _ = v40
                    case "":
                        v41 := v3[v24]
                        _ = v41
                    }
                default:
                    var v42 uint = 0
                    v1 = (v1 % ((4294967295) | 1))
                    _ = v42
                }
                if true {
                    var v43 string = v24
                    var v44 float64 = (0.1 / (float64(3.14159) + float64((0.1 * 1.5))))
                    _ = v43
                    _ = v44
                }
                v45 := GenS3{F0: 4294967295}
                _ = v45
            }
            var v46 int = 42
            _ = fi26
            _ = v46
        }
        v24 = v24
        _ = fi25
    }
    switch v24 {
    case "a":
        if true {
            switch 32768 {
            case -649:
                if true {
                    var v47 uint32 = 4294967295
                    _ = v47
                } else {
                    if (654 == 4294967295) {
                        fmt.Println(v1)
                        var v48 int8 = -76
                        if true {
                            var v49 uint = 18446744073709551615
                            if (478 != (^(v1))) {
                                v23 = (v23 | 255)
                                if (v24 <= v24) {
                                    var v50 uint16 = 0
                                    v2 = int64(v49)
                                    _ = v50
                                } else {
                                    v1 -= v1
                                }
                            }
                            v48 = 100
                            _ = v49
                        }
                        _ = v48
                    }
                    switch v24 {
                    case "hello":
                        v51 := GenS4{F0: -695, F1: 2147483647, F2: (v23 % (((^(v23))) | 1))}
                        _ = v51
                    case "":
                        v2 -= v2
                    case "a":
                        v1 += 4294967295
                    }
                    switch -2147483648 {
                    case 9223372036854775807:
                        if true {
                            fmt.Println(v1, v23)
                        }
                    case -2147483648:
                        v52 := []int{-9223372036854775808, 127}
                        _ = v52
                    }
                }
            }
        } else {
            for fi53 := 0; fi53 < 17; fi53++ {
                var v54 int32 = 42
                v55 := []int{(fi53 >> ((1) & 63)), fi53, fi53}
                if (42 >= 32767) {
                    if ((fi53 | 42) <= 4294967295) {
                        v56 := GenS2{F0: (fi53 ^ (fi53 << ((119) & 63)))}
                        _ = v56
                    }
                } else {
                    _ = len(v3)
                    fmt.Println(v54, v23)
                    for fi57 := 0; fi57 < 15; fi57++ {
                        v58 := GenS1{F0: 85, F1: 2147483647}
                        _ = fi57
                        _ = v58
                    }
                }
                _ = fi53
                _ = v54
                _ = v55
            }
        }
    case "":
        if true {
            var v59 int16 = 1
            delete(v3, (v24 + v24))
            if (v24 == v24) {
                for fi60 := 0; fi60 < 14; fi60++ {
                    v59 += v59
                    _ = fi60
                }
                v24 = "foo bar"
            }
            _ = v59
        }
    case "hello":
        v3[v24] = 3
    }
    for fi61 := 0; fi61 < 6; fi61++ {
        switch 1 {
        case 2147483647:
            fmt.Println(v23, v2)
        case 4294967295:
            v62 := v4[v24]
            _ = v62
        default:
            v63 := v5[65535]
            _ = v63
        }
        fmt.Println(fi61, fi61)
        if true {
            delete(v4, (v24 + v24))
            for fi64 := 0; fi64 < 8; fi64++ {
                v65 := GenS3{F0: 1}
                v66 := []int{int(v1), (^(fi61))}
                v67 := GenS5{F0: "hello", F1: 0, F2: 0}
                _ = fi64
                _ = v65
                _ = v66
                _ = v67
            }
        } else {
            var v68 uint8 = v23
            _ = len(v3)
            _ = v68
        }
        _ = fi61
    }
    for fi69 := 0; fi69 < 18; fi69++ {
        v1 += 1
        _ = fi69
    }
    if (true && (0.1 != 0.1)) {
        var v70 int16 = 32767
        var v71 uint32 = 241
        if true {
            if (117 < 32767) {
                delete(v5, v71)
                var v72 uint32 = (^(v71))
                _ = v72
            }
            var v73 uint64 = 230
            v74 := GenS4{F0: 0, F1: 2147483647, F2: 184}
            _ = v73
            _ = v74
        } else {
            fmt.Println(v70, v23)
            for fi75 := 0; fi75 < 17; fi75++ {
                v71 += v71
                v76 := GenS4{F0: -538, F1: 2147483647, F2: 1}
                if (0 < 18446744073709551615) {
                    if false {
                        fmt.Println(v24, v70)
                    } else {
                        fmt.Println(v23)
                        v77 := GenS5{F0: "foo bar", F1: -905, F2: int8(v2)}
                        _ = v77
                    }
                    switch v24 {
                    case "a":
                        var v78 int = (fi75 / ((-727) | 1))
                        _ = v78
                    case "hello":
                        switch v24 {
                        case "\n\t":
                            for fi79 := 0; fi79 < 1; fi79++ {
                                var v80 float64 = 0.0
                                v81 := v3[((v24 + v24) + v24)]
                                if (true || false) {
                                    v70 = v70
                                    var v82 float64 = float64((0.0 / 3.14159))
                                    _ = v82
                                } else {
                                    if false {
                                        var v83 uint64 = v1
                                        v84 := []int{(fi79 - fi79), fi75, -1, fi75}
                                        var v85 uint64 = v1
                                        _ = v83
                                        _ = v84
                                        _ = v85
                                    }
                                    v86 := []int{128, (fi79 + (fi79 >> ((255) & 63))), (fi75 / (((^(fi79))) | 1)), -9223372036854775808}
                                    _ = v86
                                }
                                _ = fi79
                                _ = v80
                                _ = v81
                            }
                        case "a":
                            var v87 int32 = 2147483647
                            _ = v87
                        default:
                            var v88 int8 = 42
                            _ = v88
                        }
                    case "":
                        if false {
                            v5[v71] = -2147483648
                            var v89 uint32 = v71
                            if false {
                                _ = len(v3)
                                for fi90 := 0; fi90 < 6; fi90++ {
                                    v91 := GenS6{F0: 1}
                                    _ = fi90
                                    _ = v91
                                }
                            }
                            _ = v89
                        } else {
                            v70 += v70
                        }
                    default:
                        if false {
                            v92 := []int{9223372036854775807, (fi75 << ((uint(v23)) & 63)), (fi75 | fi75)}
                            _ = v92
                        }
                    }
                } else {
                    if (!(!(!(((v1 / ((v1) | 1)) <= (v1 | v1)) || (60 < -1))))) {
                        for fi93 := 0; fi93 < 11; fi93++ {
                            v2 = (v2 & v2)
                            if false {
                                switch (^(fi93)) {
                                case 0:
                                    var v94 uint8 = (^(v23))
                                    _ = v94
                                case -1:
                                    if false {
                                        if true {
                                            var v95 string = ("a" + ("\n\t" + (v24 + ("hello" + v24))))
                                            var v96 int16 = (-(v70))
                                            _ = v95
                                            _ = v96
                                        }
                                        switch (fi93 ^ (fi93 - fi93)) {
                                        case 9223372036854775807:
                                            if (v71 < (v71 - 842)) {
                                                switch fi93 {
                                                case 0:
                                                    switch 42 {
                                                    case 9223372036854775807:
                                                        var v97 int16 = (-(v70))
                                                        _ = v97
                                                    default:
                                                        fmt.Println(v2)
                                                    }
                                                default:
                                                    v1 -= v1
                                                }
                                                var v98 uint32 = v71
                                                _ = v98
                                            }
                                        case 4294967295:
                                            fmt.Println(v70, fi75, v1)
                                        }
                                    } else {
                                        if false {
                                            if true {
                                                v99 := GenS7{F0: (v2 * (v2 * 373)), F1: -148}
                                                if false {
                                                    fmt.Println(v2, fi93)
                                                    var v100 bool = false
                                                    v101 := []int{(fi93 | fi93), -9223372036854775808, 0, fi93}
                                                    _ = v100
                                                    _ = v101
                                                }
                                                _ = v99
                                            }
                                            v24 = v24
                                            v24 = v24
                                            switch v24 {
                                            case "\n\t":
                                                var v102 int16 = (v70 ^ (v70 & (v70 ^ (v70 * v70))))
                                                _ = v102
                                            default:
                                                switch (((v24 + ("" + v24)) + v24) + ("" + v24)) {
                                                case "\n\t":
                                                    if ((v24 + (v24 + (v24 + v24))) <= ((("foo bar" + v24) + v24) + (("hello" + "") + "a"))) {
                                                        v1 = (v1 - 1)
                                                        fmt.Println(v24, v71)
                                                        v2 -= (v2 & (-(v2)))
                                                    }
                                                case "":
                                                    v71 -= v71
                                                case "hello":
                                                    var v103 int = 1
                                                    _ = v103
                                                default:
                                                    fmt.Println(v24, v70)
                                                }
                                            }
                                            v104 := []int{(fi75 - fi75), (fi93 - 32767)}
                                            _ = v104
                                        } else {
                                            var v105 int32 = 2147483647
                                            if true {
                                                var v106 int64 = (v2 | (v2 << ((379) & 63)))
                                                var v107 uint8 = 2
                                                _ = v106
                                                _ = v107
                                            }
                                            switch v24 {
                                            case "foo bar":
                                                v108 := GenS6{F0: 204}
                                                _ = v108
                                            case "a":
                                                switch ((v24 + "hello") + "foo bar") {
                                                case "\n\t":
                                                    if ((v2 % ((9223372036854775807) | 1)) < 0) {
                                                        v105 = 2147483647
                                                        var v109 uint32 = v71
                                                        v2 -= v2
                                                        _ = v109
                                                    }
                                                case "hello":
                                                    v70 = -32768
                                                }
                                            case "":
                                                v110 := []int{32767}
                                                _ = v110
                                            }
                                            _ = v105
                                        }
                                        _ = len(v5)
                                        switch v24 {
                                        case "\n\t":
                                            delete(v3, "a")
                                            switch (fi93 | fi75) {
                                            case 32768:
                                                var v111 string = "hello"
                                                switch v24 {
                                                case "a":
                                                    switch -818 {
                                                    case -1:
                                                        v71 -= v71
                                                    case -2147483648:
                                                        v1 += v1
                                                    case 726:
                                                        v112 := GenS8{F0: 110, F1: (!false), F2: 274}
                                                        _ = v112
                                                    }
                                                case "foo bar":
                                                    v113 := GenS9{F0: (v1 - 748), F1: 42}
                                                    _ = v113
                                                default:
                                                    var v114 float32 = 1.5
                                                    _ = v114
                                                }
                                                if (1.5 != 1.5) {
                                                    var v115 uint8 = 64
                                                    if true {
                                                        if false {
                                                            var v116 uint64 = (v1 >> ((381) & 63))
                                                            _ = v116
                                                        }
                                                    } else {
                                                        if (0.1 > 0.1) {
                                                            if (381 == 569) {
                                                                var v117 uint16 = 32767
                                                                v118 := GenS1{F0: v71, F1: 0}
                                                                _ = v117
                                                                _ = v118
                                                            } else {
                                                                v71 = v71
                                                            }
                                                        } else {
                                                            switch fi93 {
                                                            case 42:
                                                                v5[v71] = fi75
                                                            default:
                                                                v1 = (v1 ^ (v1 ^ (v1 & (v1 ^ 70))))
                                                            }
                                                            if true {
                                                                v2 -= v2
                                                            } else {
                                                                if (-395 == v70) {
                                                                    fmt.Println(v71, v23, fi93)
                                                                    switch fi93 {
                                                                    case -1:
                                                                        switch (^(fi75)) {
                                                                        case -1:
                                                                            delete(v3, v24)
                                                                        }
                                                                        if ((1e10 >= float64(1.5)) || (261 < 18446744073709551615)) {
                                                                            var v119 uint8 = v23
                                                                            v120 := GenS3{F0: 42}
                                                                            v121 := GenS5{F0: ((v111 + v24) + v24), F1: 965, F2: 0}
                                                                            _ = v119
                                                                            _ = v120
                                                                            _ = v121
                                                                        }
                                                                        var v122 float64 = (3.14159 + 0.1)
                                                                        v123 := GenS10{F0: (v23 + 42), F1: v115, F2: 9223372036854775807}
                                                                        _ = v122
                                                                        _ = v123
                                                                    case 32768:
                                                                        v124 := GenS11{F0: (-(v2))}
                                                                        v125 := []int{127, (fi75 * int(v70))}
                                                                        var v126 int32 = -2147483648
                                                                        if (false && (!(float32(3.14159) >= 3.14159))) {
                                                                            var v127 int16 = 1
                                                                            v128 := GenS10{F0: 0, F1: v115, F2: v2}
                                                                            v127 += v70
                                                                            _ = v127
                                                                            _ = v128
                                                                        }
                                                                        var v129 int64 = (-(v2))
                                                                        fmt.Println(v71, fi93, fi93)
                                                                        switch (^(fi93)) {
                                                                        case -936:
                                                                            if ((-1 != (v70 - (v70 >> ((42) & 63)))) && false) {
                                                                                v1 = 455
                                                                                v111 = ((v24 + ("foo bar" + v24)) + v111)
                                                                                v130 := GenS9{F0: v1, F1: int8(v126)}
                                                                                switch (v24 + v24) {
                                                                                case "foo bar":
                                                                                    var v131 float64 = (1e10 - 1e10)
                                                                                    _ = v131
                                                                                default:
                                                                                    _ = len(v5)
                                                                                }
                                                                                _ = v130
                                                                            } else {
                                                                                switch (fi75 | 0) {
                                                                                case 127:
                                                                                    v129 = v2
                                                                                }
                                                                            }
                                                                            v24 = (v111 + "hello")
                                                                            if true {
                                                                                var v132 float32 = ((0.0 * (-2.25 + 0.1)) + (-2.25 / ((0.1 / 1e10) + 3.14159)))
                                                                                fmt.Println(v115)
                                                                                _ = v132
                                                                            } else {
                                                                                delete(v4, "foo bar")
                                                                                switch fi75 {
                                                                                case 9223372036854775807:
                                                                                    var v133 float32 = float32(0.1)
                                                                                    _ = v133
                                                                                case -529:
                                                                                    fmt.Println(v71, v71)
                                                                                case -462:
                                                                                    v1 = uint64(v71)
                                                                                default:
                                                                                    v126 -= 0
                                                                                }
                                                                                v134 := GenS5{F0: v111, F1: v126, F2: -86}
                                                                                _ = v134
                                                                            }
                                                                            if false {
                                                                                var v135 int16 = (^(v70))
                                                                                if (!(90 > 42)) {
                                                                                    var v136 float64 = ((float64(1.5) / -2.25) * 1.5)
                                                                                    fmt.Println(fi93)
                                                                                    _ = v136
                                                                                } else {
                                                                                    var v137 int32 = v126
                                                                                    v137 += v126
                                                                                    fmt.Println(fi93)
                                                                                    _ = v137
                                                                                }
                                                                                _ = v135
                                                                            } else {
                                                                                fmt.Println(fi93)
                                                                            }
                                                                            v126 = (-(v126))
                                                                        case 42:
                                                                            var v138 int32 = (v126 | v126)
                                                                            _ = v138
                                                                        case 673:
                                                                            v71 = 878
                                                                        }
                                                                        v71 = 0
                                                                        v129 -= 1
                                                                        _ = len(v3)
                                                                        if (false && false) {
                                                                            var v139 int = fi75
                                                                            _ = v139
                                                                        }
                                                                        v111 = "foo bar"
                                                                        v140 := GenS12{F0: 65535, F1: (v1 + 2), F2: ((v71 / (((v71 | 65535)) | 1)) >= v71)}
                                                                        _ = v124
                                                                        _ = v125
                                                                        _ = v126
                                                                        _ = v129
                                                                        _ = v140
                                                                    case 524:
                                                                        v141 := GenS13{F0: 0, F1: int32(v23)}
                                                                        _ = v141
                                                                    default:
                                                                        switch ("foo bar" + v24) {
                                                                        case "hello":
                                                                            var v142 int16 = (-(v70))
                                                                            _ = v142
                                                                        case "foo bar":
                                                                            v70 += v70
                                                                        case "\n\t":
                                                                            v111 = ((v111 + "foo bar") + "hello")
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            v24 = (v24 + (v24 + (v111 + v111)))
                                                        }
                                                        v143 := GenS9{F0: (v1 + v1), F1: 42}
                                                        _ = v143
                                                    }
                                                    _ = v115
                                                }
                                                _ = v111
                                            case 503:
                                                v23 += 128
                                                v144 := GenS13{F0: 399, F1: 634}
                                                _ = v144
                                            case -186:
                                                var v145 int16 = v70
                                                _ = v145
                                            }
                                        case "":
                                            delete(v5, (v71 - v71))
                                        case "a":
                                            var v146 int16 = v70
                                            _ = v146
                                        }
                                    }
                                case 42:
                                    v147 := GenS6{F0: 1}
                                    _ = v147
                                default:
                                    v23 = v23
                                }
                                v148 := v3[v24]
                                v149 := []int{fi93, -9223372036854775808}
                                _ = v148
                                _ = v149
                            } else {
                                v150 := GenS13{F0: 688, F1: 42}
                                _ = v150
                            }
                            v151 := GenS7{F0: v2, F1: 884}
                            _ = fi93
                            _ = v151
                        }
                        for fi152 := 0; fi152 < 6; fi152++ {
                            var v153 float32 = 0.1
                            var v154 float64 = -2.25
                            v155 := GenS9{F0: v1, F1: 127}
                            _ = fi152
                            _ = v153
                            _ = v154
                            _ = v155
                        }
                    }
                    var v156 float64 = 3.14159
                    if ((true && true) || (0 < (fi75 + (fi75 - (fi75 >> ((65535) & 63)))))) {
                        for fi157 := 0; fi157 < 9; fi157++ {
                            v158 := GenS5{F0: (v24 + v24), F1: -2147483648, F2: 100}
                            v1 -= (v1 + v1)
                            _ = fi157
                            _ = v158
                        }
                    }
                    _ = v156
                }
                _ = fi75
                _ = v76
            }
            v159 := GenS11{F0: v2}
            _ = v159
        }
        _ = v70
        _ = v71
    } else {
        if (false && false) {
            var v160 uint32 = 65535
            var v161 bool = (true || (1.5 == 1.5))
            _ = v160
            _ = v161
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
