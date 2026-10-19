package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int; F1 uint; F2 int32 }
type GenS8 struct { F0 int64; F1 int8; F2 int }
type GenS9 struct { F0 int16; F1 uint64; F2 int64 }
type GenS10 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS11 struct { F0 int64; F1 int32; F2 int16 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 int; F1 string }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 int64; F1 bool }
type GenS16 struct { F0 int64; F1 bool }

func main() {
    var v1 float64 = (-2.25 * 1.5)
    var v2 uint = 65535
    var v3 int64 = -9223372036854775808
    var v4 bool = true
    var v5 uint64 = 18446744073709551615
    var v6 bool = v4
    fmt.Println(v2, v2, v1)
    v7 := GenS1{F0: -1}
    v8 := GenS2{F0: v3, F1: -736}
    if true {
        var v9 float32 = 1e10
        _ = v9
    }
    var v10 uint16 = 2
    var v11 float64 = float64(0.1)
    switch 128 {
    case -1:
        if v6 {
            v12 := []int{0, -469}
            v11 -= v11
            for fi13 := 0; fi13 < 10; fi13++ {
                v12 = append(v12, fi13)
                v10 -= 355
                _ = fi13
            }
            _ = v12
        } else {
            v14 := []int{-2147483648, 32767, 2147483647}
            _ = v14
        }
    case -9223372036854775808:
        for fi15 := 0; fi15 < 2; fi15++ {
            var v16 float32 = 0.0
            v17 := GenS3{F0: -128}
            if (v10 <= 2) {
                for fi18 := 0; fi18 < 1; fi18++ {
                    v19 := GenS4{F0: v5, F1: 0}
                    if (!v6) {
                        var v20 uint64 = v5
                        fmt.Println(v6, v10, fi18)
                        switch "" {
                        case "":
                            v21 := []int{}
                            _ = v21
                        case "a":
                            switch "hello" {
                            case "\n\t":
                                if v4 {
                                    var v22 int32 = -1
                                    v23 := []int{(fi18 ^ 2147483647), 4294967295}
                                    _ = v22
                                    _ = v23
                                }
                            case "foo bar":
                                switch ("" + "foo bar") {
                                case "":
                                    switch ("\n\t" + "foo bar") {
                                    case "\n\t":
                                        v24 := GenS3{F0: 49}
                                        _ = v24
                                    case "a":
                                        v2 += (^(v2))
                                    case "":
                                        if true {
                                            var v25 uint = (^(v2))
                                            var v26 int8 = 0
                                            _ = v25
                                            _ = v26
                                        }
                                    }
                                case "foo bar":
                                    var v27 uint32 = 1
                                    _ = v27
                                case "a":
                                    if false {
                                        switch ("" + "") {
                                        case "\n\t":
                                            if false {
                                                v3 += (-(v3))
                                            } else {
                                                v6 = true
                                                v28 := []int{(fi15 | -962), fi15, (fi18 ^ (fi18 / (((fi15 - (fi15 << ((270) & 63)))) | 1))), fi18}
                                                switch (("" + "hello") + "hello") {
                                                case "foo bar":
                                                    var v29 uint8 = 255
                                                    _ = v29
                                                default:
                                                    if len(v28) >= 2 {
                                                        v28 = v28[1:len(v28)]
                                                    }
                                                }
                                                _ = v28
                                            }
                                        }
                                    } else {
                                        if v6 {
                                            var v30 uint8 = 42
                                            _ = v30
                                        }
                                        v31 := GenS5{F0: -32768}
                                        if (42 == 2147483647) {
                                            var v32 uint32 = 203
                                            fmt.Println(v5, v1)
                                            _ = v32
                                        } else {
                                            switch ("foo bar" + ("" + "")) {
                                            case "hello":
                                                v33 := GenS4{F0: v5, F1: 128}
                                                _ = v33
                                            }
                                            var v34 uint16 = v10
                                            switch "a" {
                                            case "":
                                                v35 := GenS2{F0: v3, F1: 2147483647}
                                                _ = v35
                                            default:
                                                if false {
                                                    var v36 uint = (^(v2))
                                                    var v37 uint = (v36 ^ v36)
                                                    if v6 {
                                                        if (1 <= 1) {
                                                            fmt.Println(fi18)
                                                        } else {
                                                            switch "\n\t" {
                                                            case "a":
                                                                var v38 int8 = int8(v5)
                                                                fmt.Println(v4, v10)
                                                                if v6 {
                                                                    var v39 float64 = v11
                                                                    var v40 int64 = v3
                                                                    _ = v39
                                                                    _ = v40
                                                                }
                                                                v34 = (^(v34))
                                                                v41 := GenS1{F0: 32767}
                                                                var v42 int16 = 1
                                                                if ((-(fi18)) < (fi15 + (fi15 | -1))) {
                                                                    var v43 uint = (^(v37))
                                                                    var v44 string = "a"
                                                                    _ = v43
                                                                    _ = v44
                                                                }
                                                                switch fi15 {
                                                                case 9223372036854775807:
                                                                    var v45 int64 = (v3 << ((uint(v20)) & 63))
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        if v4 {
                                                                            v45 = (v45 >> ((v2) & 63))
                                                                        } else {
                                                                            v42 -= (v42 / (((v42 | (v42 & -766))) | 1))
                                                                        }
                                                                    default:
                                                                        if v4 {
                                                                            var v46 int16 = (v42 | v42)
                                                                            var v47 bool = v6
                                                                            _ = v46
                                                                            _ = v47
                                                                        }
                                                                    }
                                                                    var v48 uint64 = (v5 & v5)
                                                                    _ = v45
                                                                    _ = v48
                                                                case 1:
                                                                    v49 := GenS2{F0: v3, F1: 1}
                                                                    _ = v49
                                                                case 947:
                                                                    if ((^(v20)) != 18446744073709551615) {
                                                                        var v50 uint = v37
                                                                        if ((-2.25 == -2.25) || v6) {
                                                                            v51 := GenS6{F0: 127}
                                                                            var v52 int16 = 284
                                                                            if v6 {
                                                                                var v53 int8 = v38
                                                                                _ = v53
                                                                            }
                                                                            var v54 int64 = v3
                                                                            _ = v51
                                                                            _ = v52
                                                                            _ = v54
                                                                        } else {
                                                                            if (v5 != (v20 ^ 1)) {
                                                                                if (!(true || (!(fi15 <= int(v42))))) {
                                                                                    v42 = 0
                                                                                } else {
                                                                                    v55 := GenS1{F0: (v42 + v42)}
                                                                                    var v56 uint16 = v10
                                                                                    _ = v55
                                                                                    _ = v56
                                                                                }
                                                                                v6 = true
                                                                                v57 := []int{(fi18 << ((v37) & 63)), fi18}
                                                                                v58 := GenS7{F0: (fi18 * fi15), F1: 65535, F2: 2147483647}
                                                                                _ = v57
                                                                                _ = v58
                                                                            } else {
                                                                                switch "" {
                                                                                case "hello":
                                                                                    var v59 int32 = 42
                                                                                    _ = v59
                                                                                case "a":
                                                                                    switch (("a" + "\n\t") + "foo bar") {
                                                                                    case "":
                                                                                        v60 := GenS6{F0: uint8(v10)}
                                                                                        _ = v60
                                                                                    case "foo bar":
                                                                                        v61 := []int{(fi15 >> (((v37 | v2)) & 63)), (fi18 - 42), (fi18 & fi15), 9223372036854775807}
                                                                                        _ = v61
                                                                                    case "hello":
                                                                                        v62 := []int{}
                                                                                        _ = v62
                                                                                    }
                                                                                }
                                                                                switch (fi18 % ((32767) | 1)) {
                                                                                case -537:
                                                                                    var v63 int = fi15
                                                                                    _ = v63
                                                                                case 0:
                                                                                    switch "" {
                                                                                    case "\n\t":
                                                                                        if v6 {
                                                                                            fmt.Println(v34, v11)
                                                                                            v64 := GenS8{F0: v3, F1: 127, F2: 154}
                                                                                            _ = v64
                                                                                        }
                                                                                    }
                                                                                    v65 := GenS8{F0: 497, F1: (-(v38)), F2: (fi18 ^ fi18)}
                                                                                    _ = v65
                                                                                case 400:
                                                                                    var v66 int16 = (v42 >> ((0) & 63))
                                                                                    _ = v66
                                                                                default:
                                                                                    var v67 int32 = -1
                                                                                    _ = v67
                                                                                }
                                                                                v3 = (v3 - v3)
                                                                                v68 := GenS1{F0: v42}
                                                                                if v4 {
                                                                                    v69 := GenS8{F0: (v3 / ((int64(v42)) | 1)), F1: v38, F2: (fi18 + fi15)}
                                                                                    if (fi18 >= fi15) {
                                                                                        var v70 int64 = (v3 >> (((^(v2))) & 63))
                                                                                        _ = v70
                                                                                    } else {
                                                                                        v71 := []int{(fi18 | (fi18 >> (((v2 | v36)) & 63)))}
                                                                                        if (!(!(!(("hello" + "\n\t") == "")))) {
                                                                                            v71 = append(v71, (fi18 * (fi15 & (fi15 * fi15))))
                                                                                            if (969 <= 42) {
                                                                                                if true {
                                                                                                    v3 += v3
                                                                                                    v6 = v4
                                                                                                    v72 := GenS4{F0: v5, F1: 127}
                                                                                                    _ = v72
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            switch (fi15 << ((64) & 63)) {
                                                                                            case 492:
                                                                                                _ = len(v71)
                                                                                            default:
                                                                                                v71 = append(v71, (fi18 % (((fi18 & 4294967295)) | 1)))
                                                                                            }
                                                                                        }
                                                                                        if (-1 >= 407) {
                                                                                            v3 -= v3
                                                                                            v1 = v11
                                                                                        } else {
                                                                                            v3 = v3
                                                                                            if (!false) {
                                                                                                v73 := GenS7{F0: 2147483647, F1: (v36 & (v50 / ((4294967295) | 1))), F2: -1}
                                                                                                v71 = append(v71, fi18)
                                                                                                if len(v71) >= 2 {
                                                                                                    v71 = v71[1:len(v71)]
                                                                                                }
                                                                                                var v74 uint64 = (v5 ^ v5)
                                                                                                switch (fi18 * 356) {
                                                                                                case 32768:
                                                                                                    v10 = v10
                                                                                                case 39:
                                                                                                    var v75 int16 = (v42 >> ((v36) & 63))
                                                                                                    _ = v75
                                                                                                default:
                                                                                                    v71 = append(v71, (fi15 | (fi18 * (fi15 * (fi18 / ((fi15) | 1))))))
                                                                                                }
                                                                                                _ = v73
                                                                                                _ = v74
                                                                                            }
                                                                                            if v6 {
                                                                                                fmt.Println(v11, v20)
                                                                                                v11 += v1
                                                                                                if len(v71) >= 2 {
                                                                                                    v71 = v71[1:len(v71)]
                                                                                                }
                                                                                            } else {
                                                                                                v10 -= (v10 << ((v50) & 63))
                                                                                                switch (fi18 ^ (fi15 << ((255) & 63))) {
                                                                                                case -2147483648:
                                                                                                    switch (("" + ((("\n\t" + "hello") + "") + "foo bar")) + "foo bar") {
                                                                                                    case "foo bar":
                                                                                                        v76 := GenS9{F0: v42, F1: 4294967295, F2: (v3 % (((v3 ^ 42)) | 1))}
                                                                                                        _ = v76
                                                                                                    case "a":
                                                                                                        fmt.Println(v5)
                                                                                                    case "\n\t":
                                                                                                        v77 := GenS8{F0: (v3 % ((v3) | 1)), F1: 0, F2: fi15}
                                                                                                        _ = v77
                                                                                                    default:
                                                                                                        switch "\n\t" {
                                                                                                        case "hello":
                                                                                                            var v78 int8 = v38
                                                                                                            v3 -= (v3 << (((v36 / (((v50 * v36)) | 1))) & 63))
                                                                                                            _ = v78
                                                                                                        case "a":
                                                                                                            var v79 uint32 = 468
                                                                                                            if false {
                                                                                                                var v80 float64 = v1
                                                                                                                _ = v80
                                                                                                            }
                                                                                                            var v81 uint16 = (v10 | (^(v34)))
                                                                                                            var v82 uint64 = (v5 ^ (v5 >> ((v2) & 63)))
                                                                                                            _ = v79
                                                                                                            _ = v81
                                                                                                            _ = v82
                                                                                                        }
                                                                                                    }
                                                                                                case 127:
                                                                                                    var v83 bool = ((v38 % ((v38) | 1)) <= 14)
                                                                                                    _ = v83
                                                                                                default:
                                                                                                    var v84 uint16 = (v10 / (((^(v34))) | 1))
                                                                                                    _ = v84
                                                                                                }
                                                                                                v85 := GenS10{F0: 65535, F1: 2, F2: 127}
                                                                                                _ = v85
                                                                                            }
                                                                                            if false {
                                                                                                v86 := GenS7{F0: fi15, F1: (v50 | (v2 + v37)), F2: int32(v3)}
                                                                                                v1 += v1
                                                                                                _ = v86
                                                                                            }
                                                                                            var v87 uint16 = (^(v34))
                                                                                            _ = v87
                                                                                        }
                                                                                        fmt.Println(v16, v36)
                                                                                        var v88 int32 = 44
                                                                                        fmt.Println(v16, v1)
                                                                                        _ = v71
                                                                                        _ = v88
                                                                                    }
                                                                                    fmt.Println(v1, fi18, v20)
                                                                                    _ = v69
                                                                                }
                                                                                fmt.Println(v5, v3)
                                                                                _ = v68
                                                                            }
                                                                        }
                                                                        _ = v50
                                                                    }
                                                                }
                                                                v89 := []int{32767, int(v34)}
                                                                if v6 {
                                                                    var v90 uint64 = 990
                                                                    _ = v90
                                                                } else {
                                                                    v16 += v16
                                                                    var v91 uint32 = 697
                                                                    _ = v91
                                                                }
                                                                switch ("a" + "\n\t") {
                                                                case "a":
                                                                    var v92 float32 = v16
                                                                    _ = v92
                                                                }
                                                                if len(v89) > 0 {
                                                                    v89[0] = (fi15 ^ fi15)
                                                                }
                                                                _ = v38
                                                                _ = v41
                                                                _ = v42
                                                                _ = v89
                                                            }
                                                            var v93 int32 = 2147483647
                                                            _ = v93
                                                        }
                                                        v94 := []int{}
                                                        _ = v94
                                                    }
                                                    _ = v36
                                                    _ = v37
                                                } else {
                                                    v4 = v4
                                                }
                                            }
                                            _ = v34
                                        }
                                        switch "" {
                                        case "a":
                                            var v95 uint8 = 1
                                            _ = v95
                                        case "":
                                            var v96 float32 = v16
                                            _ = v96
                                        case "foo bar":
                                            var v97 bool = v4
                                            _ = v97
                                        }
                                        var v98 string = "hello"
                                        v99 := []int{4294967295, fi15, (^(fi15))}
                                        _ = v31
                                        _ = v98
                                        _ = v99
                                    }
                                }
                            case "":
                                switch "foo bar" {
                                case "":
                                    v100 := GenS11{F0: v3, F1: 1, F2: 32767}
                                    _ = v100
                                case "foo bar":
                                    var v101 uint16 = v10
                                    _ = v101
                                }
                            default:
                                v20 = v5
                            }
                        default:
                            switch "a" {
                            case "hello":
                                v102 := []int{fi18, fi18}
                                _ = v102
                            case "a":
                                v20 -= 0
                            case "foo bar":
                                var v103 uint16 = v10
                                _ = v103
                            default:
                                v6 = v6
                            }
                            if (!(1000 >= -579)) {
                                var v104 uint = (v2 & v2)
                                v10 = 809
                                v6 = true
                                _ = v104
                            }
                        }
                        _ = v20
                    } else {
                        v105 := []int{}
                        _ = v105
                    }
                    switch fi18 {
                    case 4294967295:
                        switch ("foo bar" + ("\n\t" + ((((("foo bar" + ("foo bar" + "hello")) + "") + ("" + "a")) + (("hello" + "a") + "")) + ""))) {
                        case "a":
                            var v106 int64 = v3
                            _ = v106
                        case "\n\t":
                            var v107 uint8 = 42
                            _ = v107
                        case "foo bar":
                            v6 = v4
                        default:
                            if ((v6 || v4) || ((v10 >> ((v2) & 63)) > (v10 * v10))) {
                                if (!v6) {
                                    fmt.Println(v1, v2)
                                    v108 := []int{(fi18 / ((fi18) | 1)), (fi15 & 482)}
                                    v109 := GenS12{F0: 42}
                                    _ = v108
                                    _ = v109
                                } else {
                                    v110 := GenS1{F0: 988}
                                    v16 -= v16
                                    _ = v110
                                }
                            }
                        }
                    case 681:
                        v6 = ((uint32(v5) != 353) && ((v1 - float64(v16)) < v1))
                        v111 := []int{-994, (fi15 >> (((v2 - (v2 - (v2 * (^(v2)))))) & 63))}
                        v112 := GenS13{F0: fi18, F1: "\n\t"}
                        _ = v111
                        _ = v112
                    case 753:
                        switch ("foo bar" + "") {
                        case "":
                            v3 = v3
                        case "\n\t":
                            v113 := []int{}
                            _ = v113
                        case "a":
                            var v114 uint32 = 2147483648
                            _ = v114
                        default:
                            if v6 {
                                v115 := GenS9{F0: 731, F1: v5, F2: v3}
                                fmt.Println(v6, v4)
                                var v116 int8 = -1
                                v117 := GenS14{F0: 42}
                                _ = v115
                                _ = v116
                                _ = v117
                            }
                        }
                    }
                    switch "foo bar" {
                    case "\n\t":
                        if ((((v16 / v16) / 0.1) < 0.0) && true) {
                            v118 := GenS15{F0: (v3 * 820), F1: (!v4)}
                            _ = v118
                        }
                    }
                    _ = fi18
                    _ = v19
                }
                if (v3 > 837) {
                    var v119 int32 = 147
                    var v120 int8 = -65
                    _ = v119
                    _ = v120
                }
            }
            _ = fi15
            _ = v16
            _ = v17
        }
    case -466:
        var v121 uint32 = 1
        _ = v121
    }
    if (uint16(v5) != (v10 / ((65535) | 1))) {
        switch 32767 {
        case -1:
            v5 += 687
        case 4294967295:
            if (!v4) {
                v122 := GenS11{F0: (v3 / (((v3 % ((-95) | 1))) | 1)), F1: 684, F2: 275}
                for fi123 := 0; fi123 < 8; fi123++ {
                    for fi124 := 0; fi124 < 17; fi124++ {
                        v5 = (v5 | v5)
                        _ = fi124
                    }
                    fmt.Println(v5, v2, v5)
                    switch "foo bar" {
                    case "":
                        for fi125 := 0; fi125 < 2; fi125++ {
                            v126 := []int{1, fi123, -2147483648}
                            v127 := GenS4{F0: 2, F1: 134}
                            v126 = append(v126, 9223372036854775807)
                            _ = fi125
                            _ = v126
                            _ = v127
                        }
                    }
                    _ = fi123
                }
                v3 = v3
                _ = v122
            } else {
                var v128 int32 = -1
                switch -549 {
                case -474:
                    if false {
                        var v129 int64 = v3
                        fmt.Println(v129, v6)
                        if (((v5 >> (((v2 & (v2 & 2))) & 63)) >= v5) || v4) {
                            v128 -= v128
                            v129 = v129
                            v130 := []int{}
                            _ = v130
                        } else {
                            fmt.Println(v128, v128)
                            switch "" {
                            case "a":
                                var v131 int32 = 1
                                _ = v131
                            case "hello":
                                if ("a" != "hello") {
                                    fmt.Println(v10)
                                    for fi132 := 0; fi132 < 18; fi132++ {
                                        v128 = (v128 * 1)
                                        _ = fi132
                                    }
                                } else {
                                    v6 = (!false)
                                    v133 := []int{-9223372036854775808, 127}
                                    _ = v133
                                }
                            case "foo bar":
                                v134 := []int{-9223372036854775808, 42}
                                _ = v134
                            default:
                                v135 := []int{}
                                _ = v135
                            }
                            fmt.Println(v1, v128, v10)
                        }
                        _ = v129
                    }
                default:
                    fmt.Println(v10)
                }
                v10 -= v10
                _ = v128
            }
        case 763:
            var v136 uint = uint(v5)
            _ = v136
        }
        v11 += 3.14159
        v3 = (-(v3))
    }
    for fi137 := 0; fi137 < 15; fi137++ {
        var v138 string = "hello"
        for fi139 := 0; fi139 < 14; fi139++ {
            if (v6 && (true || v4)) {
                v5 = v5
                if (false || v4) {
                    switch (fi137 >> (((v2 >> ((v2) & 63))) & 63)) {
                    case 32767:
                        v140 := []int{-548}
                        _ = v140
                    case 127:
                        v10 = v10
                    default:
                        v3 = 0
                    }
                    v141 := []int{fi137}
                    _ = v141
                }
                var v142 uint64 = 16
                switch fi139 {
                case 128:
                    var v143 bool = (v5 > v142)
                    _ = v143
                case 9223372036854775807:
                    var v144 float32 = 1.5
                    v10 += (v10 << ((v2) & 63))
                    var v145 string = ("foo bar" + "hello")
                    v5 = (v142 & (^(v142)))
                    _ = v144
                    _ = v145
                case 4294967295:
                    v146 := []int{(fi137 >> (((^(v2))) & 63)), (fi139 >> ((v2) & 63)), fi139, fi139}
                    _ = v146
                }
                var v147 float32 = ((3.14159 / 1e10) / 0.1)
                fmt.Println(v6, v3, fi139)
                _ = v142
                _ = v147
            } else {
                v148 := GenS16{F0: (v3 | (v3 - v3)), F1: true}
                var v149 uint = v2
                if false {
                    var v150 string = v138
                    var v151 uint64 = (v5 / ((uint64(v2)) | 1))
                    v152 := []int{128, (^(fi137)), 2147483647}
                    var v153 float64 = v1
                    _ = v150
                    _ = v151
                    _ = v152
                    _ = v153
                }
                v154 := []int{fi137, 32767, 4294967295, (fi137 << ((2) & 63))}
                var v155 float64 = float64((0.1 - (float32(((3.14159 / float64(float32((((float64(3.14159) - ((v11 - 0.0) - v1)) - (1e10 / v1)) + ((3.14159 * v11) - ((v11 * v1) - v11)))))) + v11)) + -2.25)))
                v11 -= 3.14159
                _ = v148
                _ = v149
                _ = v154
                _ = v155
            }
            fmt.Println(v5, fi137)
            v156 := GenS6{F0: 128}
            _ = fi139
            _ = v156
        }
        _ = fi137
        _ = v138
    }
    var v157 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v157:", v157)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v157
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
