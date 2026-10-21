package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 int8; F2 uint8 }
type GenS4 struct { F0 int8; F1 int; F2 int32 }
type GenS5 struct { F0 int; F1 int64; F2 int64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint16; F1 string; F2 int }
type GenS9 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint; F1 string; F2 uint64 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 bool; F1 int64 }
type GenS14 struct { F0 int8 }
func genFn0() bool {
    var v4 float64 = -2.25
    var v5 uint64 = 331
    for fi6 := 0; fi6 < 3; fi6++ {
        for fi7 := 0; fi7 < 18; fi7++ {
            v4 -= 3.14159
            _ = fi7
        }
        v8 := GenS1{F0: 42, F1: 32767}
        _ = fi6
        _ = v8
    }
    _ = v4
    _ = v5
    return false
}
func genFn1(p0 bool, p1 int64, p2 int16) int32 {
    if p0 {
        var v9 uint32 = 4294967295
        for fi10 := 0; fi10 < 3; fi10++ {
            p2 = (p2 | (p2 >> ((577) & 63)))
            _ = fi10
        }
        _ = v9
    } else {
        p2 = (p2 % ((p2) | 1))
        switch 42 {
        case 4294967295:
            v11 := []int{-546, -542, 504, 32768}
            _ = v11
        case 2147483647:
            if false {
                v12 := []int{2147483647}
                _ = v12
            }
        }
        if (65535 == 2) {
            switch "hello" {
            case "a":
                v13 := GenS1{F0: 1, F1: 42}
                _ = v13
            case "foo bar":
                var v14 float64 = 3.14159
                _ = v14
            case "\n\t":
                if p0 {
                    if (!(!false)) {
                        switch "a" {
                        case "foo bar":
                            v15 := GenS2{F0: p2}
                            _ = v15
                        case "hello":
                            var v16 float32 = 0.0
                            _ = v16
                        default:
                            for fi17 := 0; fi17 < 20; fi17++ {
                                var v18 uint16 = 765
                                v19 := GenS3{F0: ("" + ("a" + ("\n\t" + ("" + "\n\t")))), F1: 26, F2: 128}
                                fmt.Println(p0, v18)
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                        }
                        if (p0 && (32767 > 4294967295)) {
                            for fi20 := 0; fi20 < 13; fi20++ {
                                var v21 int8 = 100
                                for fi22 := 0; fi22 < 10; fi22++ {
                                    var v23 uint16 = 2
                                    var v24 uint = 668
                                    var v25 uint16 = 920
                                    v25 -= 458
                                    _ = fi22
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                }
                                _ = fi20
                                _ = v21
                            }
                            fmt.Println(p0)
                        } else {
                            switch "hello" {
                            case "foo bar":
                                if false {
                                    v26 := []int{-936}
                                    var v27 int16 = (^(p2))
                                    _ = v26
                                    _ = v27
                                }
                            default:
                                switch "\n\t" {
                                case "\n\t":
                                    fmt.Println(p0)
                                case "foo bar":
                                    p0 = (p0 || false)
                                case "a":
                                    var v28 uint64 = 0
                                    _ = v28
                                default:
                                    for fi29 := 0; fi29 < 3; fi29++ {
                                        v30 := GenS4{F0: -66, F1: fi29, F2: 743}
                                        _ = fi29
                                        _ = v30
                                    }
                                }
                            }
                            v31 := GenS2{F0: p2}
                            var v32 uint16 = 879
                            _ = v31
                            _ = v32
                        }
                        switch "foo bar" {
                        case "a":
                            switch 127 {
                            case 128:
                                var v33 uint = 65535
                                _ = v33
                            }
                        case "\n\t":
                            v34 := []int{-883, 128, int(p1), 0}
                            _ = v34
                        case "hello":
                            var v35 int = 1
                            _ = v35
                        }
                    } else {
                        var v36 int8 = 47
                        v37 := GenS5{F0: 32767, F1: (p1 * 0), F2: p1}
                        for fi38 := 0; fi38 < 11; fi38++ {
                            switch "" {
                            case "":
                                var v39 float32 = 0.1
                                _ = v39
                            case "\n\t":
                                var v40 int = (fi38 >> ((4294967295) & 63))
                                _ = v40
                            case "foo bar":
                                p0 = (4294967295 != 18446744073709551615)
                            default:
                                if p0 {
                                    p0 = (26 != 1)
                                    v41 := []int{fi38}
                                    p2 = 32767
                                    _ = v41
                                }
                            }
                            if ((-2.25 + 1e10) > 1e10) {
                                fmt.Println(p1, p1)
                                p1 -= (p1 & (p1 - p1))
                                var v42 int32 = 1
                                _ = v42
                            } else {
                                var v43 uint64 = 18446744073709551615
                                switch "" {
                                case "foo bar":
                                    fmt.Println(v43)
                                case "\n\t":
                                    fmt.Println(v36, p1, v36)
                                case "hello":
                                    v44 := []int{}
                                    _ = v44
                                default:
                                    var v45 int = (fi38 - int(v36))
                                    _ = v45
                                }
                                var v46 string = "hello"
                                _ = v43
                                _ = v46
                            }
                            v47 := []int{-1}
                            _ = fi38
                            _ = v47
                        }
                        _ = v36
                        _ = v37
                    }
                    fmt.Println(p2)
                    v48 := GenS3{F0: "", F1: 100, F2: 128}
                    _ = v48
                }
            }
            fmt.Println(p2, p0)
            fmt.Println(p1, p0, p1)
        } else {
            switch "a" {
            case "foo bar":
                v49 := GenS6{F0: 1}
                _ = v49
            case "\n\t":
                v50 := GenS2{F0: p2}
                _ = v50
            }
            if true {
                var v51 int16 = p2
                p1 -= p1
                _ = v51
            } else {
                p1 = (p1 + (p1 << ((1) & 63)))
                var v52 float64 = ((float64((0.0 + 0.1)) / (1e10 - 3.14159)) / -2.25)
                var v53 int16 = p2
                _ = v52
                _ = v53
            }
        }
    }
    var v54 int64 = p1
    v54 += (-(v54))
    fmt.Println(p2, v54)
    _ = p0
    _ = p1
    _ = p2
    _ = v54
    return 0
}

func main() {
    v1 := []int{32768, 127}
    var v2 int64 = 42
    var v3 uint = 65535
    if (627 <= 556) {
        if (114 != -1) {
            if true {
                var v55 uint16 = 32767
                v56 := GenS7{F0: 2147483647}
                _ = v55
                _ = v56
            } else {
                var v57 int64 = -1
                switch "a" {
                case "hello":
                    v58 := GenS8{F0: 351, F1: ("" + "hello"), F2: 450}
                    _ = v58
                }
                _ = v57
            }
            v3 -= (v3 / ((798) | 1))
        }
        switch "hello" {
        case "":
            _ = len(v1)
        case "a":
            if false {
                if false {
                    _ = len(v1)
                } else {
                    var v59 int64 = v2
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v59
                }
                v3 = 318
            } else {
                var v60 int8 = -1
                for fi61 := 0; fi61 < 1; fi61++ {
                    var v62 int16 = -32768
                    for fi63 := 0; fi63 < 13; fi63++ {
                        var v64 string = "a"
                        var v65 int = fi61
                        var v66 int8 = v60
                        _ = fi63
                        _ = v64
                        _ = v65
                        _ = v66
                    }
                    _ = fi61
                    _ = v62
                }
                _ = v60
            }
        }
        fmt.Println(v3, v2)
    } else {
        v67 := GenS9{F0: 2147483647, F1: ((!true) || false), F2: 42}
        if false {
            v68 := GenS5{F0: -622, F1: v2, F2: v2}
            v69 := GenS1{F0: 2, F1: 1}
            _ = v68
            _ = v69
        }
        _ = v67
    }
    v70 := GenS10{F0: 127}
    var v71 uint32 = 22
    var v72 uint8 = 42
    if false {
        fmt.Println(v2)
    } else {
        v2 = v2
        fmt.Println(v72, v2, v72)
        var v73 int32 = 1
        _ = v73
    }
    var v74 uint8 = (v72 * v72)
    v72 -= uint8(v3)
    if false {
        v1 = append(v1, 4294967295)
        if false {
            var v75 int8 = -1
            if len(v1) > 0 {
                v1[0] = 822
            }
            var v76 uint8 = 2
            _ = v75
            _ = v76
        } else {
            v74 += v74
            switch 9223372036854775807 {
            case -582:
                _ = len(v1)
            case -492:
                var v77 float32 = ((0.0 + 1.5) + 1.5)
                _ = v77
            case -1:
                v74 = v74
            default:
                v78 := GenS5{F0: 858, F1: v2, F2: 0}
                _ = v78
            }
        }
        for fi79 := 0; fi79 < 2; fi79++ {
            var v80 int8 = -1
            v72 -= (v72 / ((190) | 1))
            _ = fi79
            _ = v80
        }
    } else {
        v2 = (v2 + 0)
        if (true && (1 == 1)) {
            var v81 float64 = 1e10
            v72 = v72
            if (2 != 4294967295) {
                for fi82 := 0; fi82 < 5; fi82++ {
                    switch ("a" + "foo bar") {
                    case "a":
                        var v83 uint = 18446744073709551615
                        _ = v83
                    case "":
                        _ = len(v1)
                    case "hello":
                        fmt.Println(v81, v71, v71)
                    }
                    var v84 uint32 = (^(v71))
                    _ = fi82
                    _ = v84
                }
            }
            _ = v81
        }
        switch 42 {
        case 4294967295:
            for fi85 := 0; fi85 < 5; fi85++ {
                _ = len(v1)
                _ = fi85
            }
        case 9223372036854775807:
            if ((2147483648 < (v71 % ((v71) | 1))) || ((!true) && false)) {
                switch 128 {
                case 9223372036854775807:
                    v72 = 20
                default:
                    for fi86 := 0; fi86 < 13; fi86++ {
                        if (true && true) {
                            v3 -= v3
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if (true || false) {
                                v1 = append(v1, fi86)
                                if false {
                                    v87 := GenS9{F0: v71, F1: true, F2: 42}
                                    var v88 int8 = -76
                                    fmt.Println(fi86, v3)
                                    _ = v87
                                    _ = v88
                                } else {
                                    var v89 string = "\n\t"
                                    _ = v89
                                }
                                fmt.Println(v74, v3, fi86)
                            } else {
                                for fi90 := 0; fi90 < 16; fi90++ {
                                    if len(v1) > 0 {
                                        v1[0] = 90
                                    }
                                    if ((v3 + (^(v3))) == 797) {
                                        fmt.Println(v72)
                                        if (32767 <= 792) {
                                            v91 := GenS7{F0: -2147483648}
                                            var v92 int64 = v2
                                            var v93 string = "a"
                                            _ = v91
                                            _ = v92
                                            _ = v93
                                        } else {
                                            if true {
                                                if len(v1) > 0 {
                                                    v1[0] = fi86
                                                }
                                                v74 = (v72 / (((v72 >> (((v3 ^ 18446744073709551615)) & 63))) | 1))
                                                var v94 int = 128
                                                _ = v94
                                            }
                                            v3 += v3
                                        }
                                    }
                                    _ = fi90
                                }
                                fmt.Println(fi86)
                                fmt.Println(v71)
                            }
                        }
                        _ = fi86
                    }
                }
            } else {
                v72 = (v74 << ((v3) & 63))
                fmt.Println(v3, v71)
            }
        case -71:
            var v95 uint64 = 18446744073709551615
            _ = v95
        default:
            switch "hello" {
            case "a":
                v96 := GenS9{F0: 2, F1: (!((!(42 > 42)) && false)), F2: 18446744073709551615}
                _ = v96
            }
        }
    }
    v74 += 127
    if (!true) {
        v72 -= 0
        v1 = append(v1, -577)
        for fi97 := 0; fi97 < 1; fi97++ {
            fmt.Println(fi97, v74, v72)
            v98 := GenS1{F0: 2, F1: 312}
            _ = fi97
            _ = v98
        }
    }
    switch -9223372036854775808 {
    case -101:
        v99 := GenS9{F0: v71, F1: true, F2: 1}
        _ = v99
    }
    if (!true) {
        if ((65535 <= (v3 % ((255) | 1))) && false) {
            v2 += (^(v2))
            v74 = (v72 * 42)
            switch 32767 {
            case -88:
                switch "" {
                case "":
                    fmt.Println(v74, v71)
                case "hello":
                    fmt.Println(v71)
                case "a":
                    if len(v1) > 0 {
                        v1[0] = 4294967295
                    }
                }
            case -961:
                v3 -= v3
            default:
                if false {
                    var v100 float64 = 3.14159
                    fmt.Println(v71)
                    _ = v100
                } else {
                    for fi101 := 0; fi101 < 13; fi101++ {
                        switch 32768 {
                        case -640:
                            v74 += uint8(v71)
                        }
                        _ = fi101
                    }
                    v102 := GenS11{F0: v3, F1: "\n\t", F2: 1}
                    _ = v102
                }
            }
        } else {
            v72 += (v74 | (v74 / ((v72) | 1)))
            if false {
                switch "\n\t" {
                case "hello":
                    var v103 uint8 = v72
                    _ = v103
                }
                for fi104 := 0; fi104 < 13; fi104++ {
                    if len(v1) > 0 {
                        v1[0] = (fi104 + (fi104 / (((fi104 >> ((65535) & 63))) | 1)))
                    }
                    v74 = v72
                    if false {
                        var v105 int = fi104
                        v106 := GenS5{F0: 32768, F1: v2, F2: v2}
                        for fi107 := 0; fi107 < 12; fi107++ {
                            if (-128 == -128) {
                                if (v3 < (v3 % ((v3) | 1))) {
                                    if false {
                                        var v108 uint32 = v71
                                        var v109 int8 = 42
                                        v74 = v72
                                        _ = len(v1)
                                        _ = v108
                                        _ = v109
                                    }
                                    v71 -= v71
                                    if ((-15 > 84) && false) {
                                        fmt.Println(v72)
                                        v110 := GenS12{F0: 1}
                                        _ = v110
                                    } else {
                                        if false {
                                            var v111 bool = true
                                            v1 = append(v1, (fi107 + fi104))
                                            _ = v111
                                        }
                                        switch fi107 {
                                        case -2147483648:
                                            v1 = append(v1, (fi107 * (fi107 | 0)))
                                        default:
                                            var v112 int = fi107
                                            _ = v112
                                        }
                                    }
                                    fmt.Println(fi104, v71, fi107)
                                    fmt.Println(fi104, fi104, fi107)
                                }
                                if (0.0 == 0.0) {
                                    switch 42 {
                                    case -718:
                                        switch "foo bar" {
                                        case "":
                                            v3 = (v3 + (v3 << ((42) & 63)))
                                        case "\n\t":
                                            var v113 uint8 = 128
                                            _ = v113
                                        }
                                        v114 := GenS4{F0: -128, F1: v105, F2: 616}
                                        _ = v114
                                    case -225:
                                        var v115 uint16 = 129
                                        _ = v115
                                    default:
                                        v116 := GenS3{F0: (((("a" + "foo bar") + ("foo bar" + "hello")) + "") + "hello"), F1: -1, F2: (^(v74))}
                                        _ = v116
                                    }
                                }
                                fmt.Println(fi104, v105, v105)
                                v71 = 2147483648
                            } else {
                                v2 = v2
                                var v117 float32 = (float32(0.1) + 1.5)
                                v118 := GenS4{F0: -34, F1: -2147483648, F2: 1}
                                if len(v1) > 0 {
                                    v1[0] = (v105 << ((65535) & 63))
                                }
                                v3 = (v3 / (((^(v3))) | 1))
                                _ = v117
                                _ = v118
                            }
                            v74 += (v74 & (v72 << ((v3) & 63)))
                            _ = fi107
                        }
                        _ = v105
                        _ = v106
                    } else {
                        v119 := GenS3{F0: "", F1: 49, F2: 2}
                        _ = len(v1)
                        switch ("\n\t" + "hello") {
                        case "":
                            if len(v1) > 0 {
                                v1[0] = fi104
                            }
                        case "\n\t":
                            var v120 int16 = 1000
                            _ = v120
                        case "a":
                            if true {
                                var v121 uint16 = 32767
                                for fi122 := 0; fi122 < 5; fi122++ {
                                    var v123 string = (("\n\t" + "foo bar") + "\n\t")
                                    _ = fi122
                                    _ = v123
                                }
                                _ = v121
                            } else {
                                switch fi104 {
                                case 213:
                                    v71 += 2147483648
                                case 1:
                                    for fi124 := 0; fi124 < 4; fi124++ {
                                        switch ("" + (("" + ("foo bar" + "")) + "a")) {
                                        case "hello":
                                            var v125 uint16 = 910
                                            _ = v125
                                        }
                                        _ = fi124
                                    }
                                case -9223372036854775808:
                                    v71 += v71
                                default:
                                    if false {
                                        fmt.Println(v72, fi104, v74)
                                    }
                                }
                            }
                        }
                        _ = v119
                    }
                    _ = fi104
                }
            } else {
                v1 = append(v1, -2147483648)
                var v126 int32 = 42
                _ = len(v1)
                _ = v126
            }
            for fi127 := 0; fi127 < 5; fi127++ {
                if (((v72 / (((v74 % ((v74) | 1))) | 1)) < 39) || false) {
                    if (((101 != v71) && true) && false) {
                        fmt.Println(v2, v71, v71)
                        v74 = 1
                    }
                } else {
                    v3 += 65535
                    var v128 uint16 = 335
                    _ = len(v1)
                    _ = v128
                }
                for fi129 := 0; fi129 < 18; fi129++ {
                    var v130 int32 = 2147483647
                    switch fi127 {
                    case -9223372036854775808:
                        var v131 float32 = 0.1
                        switch (("" + "\n\t") + ("" + ("a" + "hello"))) {
                        case "hello":
                            v132 := GenS13{F0: true, F1: (v2 | v2)}
                            _ = v132
                        default:
                            fmt.Println(fi129)
                        }
                        v130 = v130
                        _ = v131
                    }
                    _ = fi129
                    _ = v130
                }
                if (!true) {
                    var v133 bool = (-1 < -21)
                    for fi134 := 0; fi134 < 2; fi134++ {
                        var v135 string = "foo bar"
                        if (true && v133) {
                            v72 += 42
                        }
                        _ = fi134
                        _ = v135
                    }
                    _ = v133
                }
                _ = fi127
            }
        }
    }
    v136 := GenS8{F0: 32767, F1: "foo bar", F2: -22}
    _ = len(v1)
    v137 := GenS14{F0: 17}
    var v138 float32 = (3.14159 / (-2.25 / -2.25))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v138:", v138)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v136
    _ = v137
    _ = v138
    _ = v2
    _ = v3
    _ = v70
    _ = v71
    _ = v72
    _ = v74
}
