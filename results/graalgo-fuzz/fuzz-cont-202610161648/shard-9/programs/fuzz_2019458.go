package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int16; F1 string; F2 uint }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS6 struct { F0 uint8; F1 uint16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint; F1 uint; F2 string }
type GenS9 struct { F0 uint32; F1 int8; F2 uint }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint16; F1 uint64 }
type GenS13 struct { F0 uint; F1 uint8; F2 int16 }
func genFn0() bool {
    v6 := GenS1{F0: 701, F1: 784}
    if false {
        var v7 int8 = 127
        v8 := GenS1{F0: 1, F1: 18446744073709551615}
        var v9 int16 = 42
        _ = v7
        _ = v8
        _ = v9
    } else {
        var v10 uint8 = 145
        _ = v10
    }
    v11 := GenS2{F0: 1000, F1: "hello", F2: 433}
    for fi12 := 0; fi12 < 13; fi12++ {
        for fi13 := 0; fi13 < 9; fi13++ {
            var v14 int16 = -1
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    if true {
        if true {
            if true {
                var v15 uint64 = 644
                if (0.1 != 0.0) {
                    v16 := GenS1{F0: 255, F1: 42}
                    if true {
                        v15 = v15
                        if (42 == -9223372036854775808) {
                            var v17 int8 = 0
                            switch 980 {
                            case 32767:
                                var v18 uint32 = 42
                                _ = v18
                            }
                            _ = v17
                        }
                        v15 += v15
                    }
                    var v19 uint64 = v15
                    _ = v16
                    _ = v19
                }
                if (v15 != 86) {
                    var v20 int64 = -666
                    for fi21 := 0; fi21 < 5; fi21++ {
                        if false {
                            var v22 bool = (!true)
                            _ = v22
                        } else {
                            v23 := GenS2{F0: 1, F1: ("" + "hello"), F2: 42}
                            _ = v23
                        }
                        var v24 bool = true
                        v25 := []int{(fi21 >> ((255) & 63)), (fi21 / ((fi21) | 1)), fi21}
                        _ = fi21
                        _ = v24
                        _ = v25
                    }
                    v20 -= (^(v20))
                    _ = v20
                } else {
                    v26 := []int{9223372036854775807, 0, 42}
                    _ = v26
                }
                _ = v15
            }
        }
    }
    _ = v11
    _ = v6
    return false
}
func genFn1(p0 int16, p1 int16, p2 uint32) int {
    p0 = 469
    v27 := GenS2{F0: (p0 - p1), F1: "hello", F2: uint(p2)}
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    return 2147483647
}

func main() {
    v1 := []uint32{4294967295, 311, 0, 65535}
    var v2 int = 743
    var v3 int64 = 581
    var v4 int64 = v3
    v5 := make(map[string]bool)
    v4 = -1
    fmt.Println(v4)
    v2 -= 9223372036854775807
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v28 := GenS2{F0: -32768, F1: "a", F2: 896}
    v29 := GenS1{F0: 32767, F1: 396}
    var v30 uint = 4294967295
    if ((true || true) || (!true)) {
        fmt.Println(v30)
        var v31 float64 = float64(3.14159)
        _ = v31
    }
    switch (("hello" + "") + "hello") {
    case "hello":
        if (42 > 127) {
            if true {
                for fi32 := 0; fi32 < 14; fi32++ {
                    v33 := GenS3{F0: 100, F1: (v30 ^ 490)}
                    _ = fi32
                    _ = v33
                }
                fmt.Println(v30)
                if false {
                    v2 += v2
                } else {
                    var v34 uint64 = 1
                    fmt.Println(v34, v3)
                    switch genFn1(-104, 0, 526) {
                    case 1:
                        var v35 float64 = 1e10
                        _ = v35
                    default:
                        v34 = (v34 / (((v34 * v34)) | 1))
                    }
                    _ = v34
                }
            }
            for fi36 := 0; fi36 < 3; fi36++ {
                var v37 int64 = v4
                _ = fi36
                _ = v37
            }
            _ = len(v5)
        }
    case "":
        v38 := GenS2{F0: 1000, F1: (("foo bar" + "hello") + "\n\t"), F2: 735}
        _ = v38
    default:
        v4 -= (v4 * v3)
    }
    for fi39 := 0; fi39 < 11; fi39++ {
        v4 -= v4
        var v40 float64 = 3.14159
        _ = fi39
        _ = v40
    }
    v4 = (v4 + (v3 ^ v4))
    if true {
        v41 := GenS2{F0: 1000, F1: (("hello" + ("foo bar" + "foo bar")) + ""), F2: (v30 * v30)}
        var v42 bool = false
        _ = v41
        _ = v42
    } else {
        for fi43 := 0; fi43 < 15; fi43++ {
            for fi44 := 0; fi44 < 10; fi44++ {
                v45 := GenS4{F0: 1}
                v3 -= v4
                _ = fi44
                _ = v45
            }
            _ = fi43
        }
        var v46 uint16 = uint16(v30)
        if (137 <= 0) {
            switch "hello" {
            case "":
                for fi47 := 0; fi47 < 10; fi47++ {
                    var v48 int64 = (v3 & 1)
                    _ = fi47
                    _ = v48
                }
            case "hello":
                for fi49 := 0; fi49 < 3; fi49++ {
                    if false {
                        var v50 uint8 = 1
                        _ = v50
                    } else {
                        var v51 float64 = -2.25
                        if true {
                            var v52 uint = v30
                            for fi53 := 0; fi53 < 18; fi53++ {
                                fmt.Println(fi53)
                                _ = fi53
                            }
                            switch 32768 {
                            case -526:
                                v46 = (v46 & 1)
                            case 127:
                                v30 -= (v52 | (^(v30)))
                            default:
                                switch "\n\t" {
                                case "\n\t":
                                    if true {
                                        for fi54 := 0; fi54 < 1; fi54++ {
                                            v55 := GenS3{F0: -1, F1: v30}
                                            var v56 string = "foo bar"
                                            _ = fi54
                                            _ = v55
                                            _ = v56
                                        }
                                        if (878 == -602) {
                                            v57 := GenS3{F0: 1, F1: v52}
                                            if (float64((1.5 - 0.1)) > -2.25) {
                                                var v58 int64 = v4
                                                var v59 float32 = -2.25
                                                _ = v58
                                                _ = v59
                                            } else {
                                                v60 := GenS1{F0: (v46 | (v46 - 739)), F1: 0}
                                                _ = v60
                                            }
                                            var v61 bool = false
                                            _ = v57
                                            _ = v61
                                        } else {
                                            v62 := GenS5{F0: true, F1: 139, F2: (v46 ^ 0)}
                                            for fi63 := 0; fi63 < 8; fi63++ {
                                                switch (fi63 - 42) {
                                                case 0:
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        fmt.Println(fi63, v4)
                                                    case "foo bar":
                                                        var v64 int8 = 100
                                                        _ = v64
                                                    }
                                                case 2147483647:
                                                    _ = len(v5)
                                                case 16:
                                                    fmt.Println(fi63)
                                                }
                                                switch "a" {
                                                case "foo bar":
                                                    var v65 float64 = (1.5 - 0.1)
                                                    _ = v65
                                                default:
                                                    var v66 string = "hello"
                                                    _ = v66
                                                }
                                                if true {
                                                    var v67 int16 = 0
                                                    _ = len(v5)
                                                    _ = v67
                                                }
                                                _ = fi63
                                            }
                                            _ = v62
                                        }
                                        var v68 uint16 = (v46 / ((v46) | 1))
                                        _ = v68
                                    }
                                case "a":
                                    if (1 <= 1000) {
                                        switch "" {
                                        case "hello":
                                            fmt.Println(v51)
                                        case "foo bar":
                                            for fi69 := 0; fi69 < 5; fi69++ {
                                                fmt.Println(v30, v46)
                                                var v70 uint32 = 42
                                                _ = fi69
                                                _ = v70
                                            }
                                        case "\n\t":
                                            if len(v1) > 0 {
                                                v1[0] = 1
                                            }
                                        default:
                                            fmt.Println(v2)
                                        }
                                        v46 = (v46 % (((v46 >> ((562) & 63))) | 1))
                                        if (100 > 1) {
                                            v46 -= 42
                                            v71 := GenS4{F0: -1}
                                            _ = v71
                                        } else {
                                            switch (fi49 - (v2 | v2)) {
                                            case 4294967295:
                                                fmt.Println(v52, v46, v2)
                                            case -490:
                                                v4 -= (v3 / ((42) | 1))
                                            case -2147483648:
                                                var v72 string = ""
                                                _ = v72
                                            }
                                            v52 += v30
                                            if (true || (!true)) {
                                                if false {
                                                    if ((v3 * (^(v3))) >= 0) {
                                                        for fi73 := 0; fi73 < 6; fi73++ {
                                                            v51 = v51
                                                            var v74 int = 2147483647
                                                            _ = fi73
                                                            _ = v74
                                                        }
                                                    } else {
                                                        fmt.Println(v30, v46, v51)
                                                        var v75 uint32 = 2147483648
                                                        var v76 uint8 = 255
                                                        _ = v75
                                                        _ = v76
                                                    }
                                                } else {
                                                    fmt.Println(v46, v46)
                                                    fmt.Println(v30)
                                                }
                                            } else {
                                                switch fi49 {
                                                case 95:
                                                    v77 := GenS6{F0: 163, F1: (v46 & v46)}
                                                    _ = v77
                                                case -175:
                                                    if false {
                                                        for fi78 := 0; fi78 < 4; fi78++ {
                                                            fmt.Println(fi49, v46, v2)
                                                            _ = fi78
                                                        }
                                                    }
                                                case 128:
                                                    for fi79 := 0; fi79 < 7; fi79++ {
                                                        switch int(v30) {
                                                        case 0:
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        }
                                                        _ = fi79
                                                    }
                                                default:
                                                    if (-553 >= -32768) {
                                                        var v80 uint64 = 779
                                                        switch fi49 {
                                                        case -128:
                                                            if (!true) {
                                                                if (2 <= (v46 / ((v46) | 1))) {
                                                                    for fi81 := 0; fi81 < 2; fi81++ {
                                                                        var v82 int32 = -947
                                                                        fmt.Println(v80, v80, v2)
                                                                        var v83 string = "hello"
                                                                        _ = fi81
                                                                        _ = v82
                                                                        _ = v83
                                                                    }
                                                                    var v84 uint16 = 255
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    _ = v84
                                                                } else {
                                                                    if (false || false) {
                                                                        v85 := GenS4{F0: 1}
                                                                        v86 := GenS6{F0: 1, F1: (v46 + v46)}
                                                                        if false {
                                                                            switch "a" {
                                                                            case "a":
                                                                                if (float32(((-2.25 + float64(3.14159)) * 0.0)) == 1e10) {
                                                                                    v87 := GenS3{F0: 42, F1: (v30 + uint(v80))}
                                                                                    switch "" {
                                                                                    case "\n\t":
                                                                                        v46 = uint16(v80)
                                                                                    default:
                                                                                        for fi88 := 0; fi88 < 1; fi88++ {
                                                                                            v46 = v46
                                                                                            v46 -= v46
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = 42
                                                                                            }
                                                                                            _ = fi88
                                                                                        }
                                                                                    }
                                                                                    _ = v87
                                                                                }
                                                                            default:
                                                                                v46 += 32767
                                                                            }
                                                                        } else {
                                                                            var v89 uint32 = 1
                                                                            _ = v89
                                                                        }
                                                                        _ = v85
                                                                        _ = v86
                                                                    } else {
                                                                        _ = len(v1)
                                                                    }
                                                                    var v90 int16 = 524
                                                                    v30 += v30
                                                                    _ = v90
                                                                }
                                                                fmt.Println(v3)
                                                            }
                                                        default:
                                                            for fi91 := 0; fi91 < 14; fi91++ {
                                                                v92 := GenS3{F0: 1, F1: 18446744073709551615}
                                                                v30 = uint(v80)
                                                                _ = fi91
                                                                _ = v92
                                                            }
                                                        }
                                                        switch fi49 {
                                                        case 128:
                                                            v93 := GenS7{F0: (v3 & (v4 - int64(v46)))}
                                                            _ = v93
                                                        case 42:
                                                            var v94 int8 = -1
                                                            _ = v94
                                                        default:
                                                            var v95 int64 = -523
                                                            _ = v95
                                                        }
                                                        _ = v80
                                                    } else {
                                                        v96 := GenS6{F0: 200, F1: 65535}
                                                        var v97 uint64 = 867
                                                        _ = v96
                                                        _ = v97
                                                    }
                                                }
                                            }
                                        }
                                    } else {
                                        switch (v2 - (v2 % (((fi49 >> (((^(v52))) & 63))) | 1))) {
                                        case 128:
                                            v98 := GenS7{F0: (v3 ^ (v3 + 42))}
                                            _ = v98
                                        case 1:
                                            for fi99 := 0; fi99 < 20; fi99++ {
                                                v52 += v52
                                                v2 = genFn1(32767, -32768, 345)
                                                _ = fi99
                                            }
                                        case -1:
                                            switch "" {
                                            case "\n\t":
                                                v100 := GenS4{F0: -305}
                                                _ = v100
                                            case "a":
                                                var v101 uint8 = 92
                                                _ = v101
                                            }
                                        default:
                                            var v102 uint8 = 127
                                            _ = v102
                                        }
                                        for fi103 := 0; fi103 < 4; fi103++ {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v104 := GenS3{F0: -35, F1: (v52 | v30)}
                                            _ = fi103
                                            _ = v104
                                        }
                                    }
                                case "":
                                    fmt.Println(v2)
                                default:
                                    fmt.Println(v46, v4)
                                }
                            }
                            _ = v52
                        } else {
                            if (false || ((v4 / ((v4) | 1)) != -1)) {
                                switch genFn1(-1, 221, 65535) {
                                case 2147483647:
                                    if true {
                                        v105 := GenS8{F0: (v30 % ((42) | 1)), F1: v30, F2: "a"}
                                        _ = v105
                                    } else {
                                        v106 := GenS9{F0: 0, F1: 69, F2: v30}
                                        v107 := GenS10{F0: (^(v30))}
                                        _ = v106
                                        _ = v107
                                    }
                                default:
                                    _ = len(v1)
                                }
                            }
                        }
                        _ = v51
                    }
                    if false {
                        var v108 uint64 = 258
                        switch "foo bar" {
                        case "\n\t":
                            if (true || true) {
                                switch (^(v2)) {
                                case -1:
                                    var v109 uint = (v30 | (^(v30)))
                                    _ = v109
                                }
                            } else {
                                switch (^(v2)) {
                                case -2147483648:
                                    var v110 int64 = v3
                                    _ = v110
                                case -1:
                                    for fi111 := 0; fi111 < 7; fi111++ {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        if false {
                                            if (32768 <= fi49) {
                                                fmt.Println(v46, v4, fi49)
                                            } else {
                                                if (false || true) {
                                                    var v112 int = (fi49 / (((v2 & (fi111 | v2))) | 1))
                                                    if (true || ((^(v108)) >= v108)) {
                                                        var v113 string = "foo bar"
                                                        _ = v113
                                                    } else {
                                                        v108 += 458
                                                        fmt.Println(v108)
                                                    }
                                                    switch ("" + "\n\t") {
                                                    case "foo bar":
                                                        v114 := v5["hello"]
                                                        _ = v114
                                                    default:
                                                        v5["foo bar"] = true
                                                    }
                                                    _ = v112
                                                }
                                                var v115 float64 = 3.14159
                                                v46 += v46
                                                _ = v115
                                            }
                                            switch v2 {
                                            case 42:
                                                v116 := GenS11{F0: 0}
                                                v2 = int(v108)
                                                _ = v116
                                            }
                                            v108 -= v108
                                        } else {
                                            var v117 float32 = (3.14159 * (0.1 + 0.0))
                                            fmt.Println(v4)
                                            _ = v117
                                        }
                                        _ = fi111
                                    }
                                case 507:
                                    v108 += 0
                                default:
                                    v46 += 32767
                                }
                                fmt.Println(v4, v108)
                            }
                        case "hello":
                            v118 := GenS9{F0: 2, F1: -128, F2: (v30 & v30)}
                            _ = v118
                        case "a":
                            v119 := GenS12{F0: v46, F1: (^(v108))}
                            _ = v119
                        }
                        _ = v108
                    }
                    _ = fi49
                }
            case "a":
                var v120 string = ""
                _ = v120
            }
            if true {
                for fi121 := 0; fi121 < 20; fi121++ {
                    var v122 int64 = (-(v3))
                    v123 := GenS1{F0: 42, F1: 42}
                    if len(v1) > 0 {
                        v1[0] = 1
                    }
                    _ = fi121
                    _ = v122
                    _ = v123
                }
            }
            v4 += 42
        } else {
            for fi124 := 0; fi124 < 2; fi124++ {
                fmt.Println(v30)
                v1 = append(v1, 691)
                _ = fi124
            }
            delete(v5, (("a" + ("a" + "a")) + ("" + ("\n\t" + "hello"))))
        }
        _ = v46
    }
    for fi125 := 0; fi125 < 3; fi125++ {
        v126 := v5[""]
        _ = fi125
        _ = v126
    }
    fmt.Println(v30, v30, v2)
    var v127 int32 = -2147483648
    if ((82 > 42) && ((1 <= 255) && (((((0.1 + -2.25) + (float32(0.1) + 1.5)) - 0.0) != 1.5) || (1 != 1)))) {
        for fi128 := 0; fi128 < 11; fi128++ {
            for fi129 := 0; fi129 < 18; fi129++ {
                v2 -= (fi129 << ((v30) & 63))
                _ = fi129
            }
            switch 32767 {
            case 0:
                if false {
                    v130 := GenS5{F0: true, F1: 2, F2: 660}
                    _ = v130
                }
            }
            v131 := GenS5{F0: true, F1: 0, F2: 255}
            _ = fi128
            _ = v131
        }
        switch (("" + "") + "foo bar") {
        case "foo bar":
            if (32767 >= 961) {
                var v132 uint = (v30 * (v30 * v30))
                _ = v132
            } else {
                var v133 uint64 = 4294967295
                if ((v2 ^ v2) == -878) {
                    _ = len(v1)
                    v1 = append(v1, 0)
                    if false {
                        if ((-2147483648 >= v127) && true) {
                            if true {
                                switch "\n\t" {
                                case "hello":
                                    var v134 uint32 = 19
                                    _ = v134
                                }
                                _ = len(v5)
                            } else {
                                v3 -= v3
                                var v135 int = 575
                                var v136 int = int(v3)
                                _ = v135
                                _ = v136
                            }
                            switch genFn1(1000, 32767, 65535) {
                            case -2147483648:
                                var v137 int16 = 30
                                _ = v137
                            }
                        }
                        fmt.Println(v2, v4, v2)
                        fmt.Println(v127)
                    } else {
                        if false {
                            v138 := GenS12{F0: 65535, F1: 42}
                            v139 := GenS13{F0: 0, F1: 0, F2: 219}
                            if true {
                                v2 -= v2
                                v2 = v2
                            } else {
                                switch (v2 % (((v2 * (-(v2)))) | 1)) {
                                case 4294967295:
                                    if (((!((v2 | genFn1(1000, -32768, 4294967295)) <= 128)) || (1 <= 42)) || (!true)) {
                                        var v140 int64 = (v4 << (((v30 * 18446744073709551615)) & 63))
                                        _ = v140
                                    } else {
                                        v141 := GenS11{F0: 2147483647}
                                        switch -9223372036854775808 {
                                        case 42:
                                            if false {
                                                _ = len(v1)
                                            } else {
                                                if true {
                                                    var v142 float32 = 3.14159
                                                    _ = v142
                                                }
                                                var v143 string = "a"
                                                _ = len(v1)
                                                _ = v143
                                            }
                                        }
                                        _ = v141
                                    }
                                case -9223372036854775808:
                                    var v144 float64 = 0.0
                                    _ = v144
                                case 128:
                                    v133 = v133
                                }
                            }
                            _ = v138
                            _ = v139
                        } else {
                            var v145 int32 = v127
                            if true {
                                var v146 int = (v2 | 2147483647)
                                for fi147 := 0; fi147 < 14; fi147++ {
                                    if (1 != v133) {
                                        for fi148 := 0; fi148 < 3; fi148++ {
                                            v1 = append(v1, 0)
                                            _ = fi148
                                        }
                                    } else {
                                        v127 += v145
                                    }
                                    delete(v5, "")
                                    _ = fi147
                                }
                                var v149 string = "a"
                                _ = v146
                                _ = v149
                            }
                            v4 += v3
                            _ = v145
                        }
                        var v150 uint = v30
                        _ = v150
                    }
                } else {
                    if (v127 <= (v127 << ((v30) & 63))) {
                        v30 += v30
                        for fi151 := 0; fi151 < 20; fi151++ {
                            var v152 uint64 = 230
                            _ = fi151
                            _ = v152
                        }
                    } else {
                        v153 := GenS11{F0: 248}
                        _ = v153
                    }
                    fmt.Println(v4, v133)
                }
                _ = v133
            }
        case "a":
            for fi154 := 0; fi154 < 3; fi154++ {
                var v155 int16 = 32767
                _ = fi154
                _ = v155
            }
        case "hello":
            v3 = v3
        default:
            if (255 <= 947) {
                fmt.Println(v30)
                _ = len(v5)
            } else {
                if false {
                    v1 = append(v1, 42)
                    for fi156 := 0; fi156 < 20; fi156++ {
                        fmt.Println(v3, v4)
                        fmt.Println(fi156, v4)
                        v30 = 4294967295
                        _ = fi156
                    }
                    var v157 uint = 491
                    _ = v157
                }
            }
        }
    }
    v158 := GenS3{F0: -19, F1: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v127:", v127)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v127
    _ = v158
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
