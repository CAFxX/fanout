package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint; F1 int16; F2 int64 }
type GenS5 struct { F0 int; F1 uint16; F2 bool }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 bool; F1 int64; F2 uint16 }
type GenS8 struct { F0 int8; F1 bool; F2 int32 }
type GenS9 struct { F0 uint8; F1 int16 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int64; F1 uint32; F2 uint8 }
type GenS12 struct { F0 int8; F1 int64; F2 uint8 }
type GenS13 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS14 struct { F0 uint16; F1 int; F2 int8 }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 bool; F1 uint; F2 bool }
type GenS17 struct { F0 uint16; F1 string }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]uint32)
    v3 := []int{32768, 568, -1}
    var v4 int64 = 21
    var v5 uint32 = 4294967295
    v6 := GenS1{F0: -32768, F1: 2, F2: 0}
    v4 -= (v4 << ((92) & 63))
    switch 1 {
    case -2147483648:
        v4 -= (-(v4))
    case 2147483647:
        for fi7 := 0; fi7 < 19; fi7++ {
            v2[1] = (v5 | (v5 * (v5 - v5)))
            for fi8 := 0; fi8 < 20; fi8++ {
                if false {
                    if false {
                        switch fi7 {
                        case -741:
                            v3 = append(v3, (fi7 ^ (fi8 + (fi8 << ((0) & 63)))))
                        default:
                            switch -1 {
                            case 128:
                                switch (("\n\t" + "foo bar") + "\n\t") {
                                case "hello":
                                    v4 = v4
                                case "foo bar":
                                    var v9 int = (fi8 >> ((65535) & 63))
                                    _ = v9
                                case "":
                                    v4 += v4
                                default:
                                    var v10 uint = 0
                                    _ = v10
                                }
                                v5 -= (v5 / (((v5 & 2147483647)) | 1))
                            }
                            if true {
                                var v11 int = fi7
                                v11 = int(v5)
                                if false {
                                    var v12 float64 = 0.0
                                    v11 = v11
                                    v13 := GenS1{F0: 1, F1: 381, F2: v4}
                                    _ = v12
                                    _ = v13
                                } else {
                                    var v14 float64 = 0.0
                                    var v15 uint32 = v5
                                    v5 = v15
                                    var v16 int64 = 42
                                    v17 := GenS1{F0: 42, F1: 2, F2: v4}
                                    var v18 uint = 586
                                    _ = v14
                                    _ = v15
                                    _ = v16
                                    _ = v17
                                    _ = v18
                                }
                                if ((0.0 * 0.0) != (3.14159 - 0.1)) {
                                    var v19 int = 32767
                                    switch (("foo bar" + "hello") + (((((("\n\t" + ("a" + ("" + "hello"))) + "") + "a") + "") + "a") + "")) {
                                    case "":
                                        if true {
                                            switch v11 {
                                            case 822:
                                                if true {
                                                    v5 += (v5 << ((255) & 63))
                                                }
                                            case -209:
                                                v4 = v4
                                            case 2147483647:
                                                if (65535 != 1) {
                                                    if (-468 >= 1) {
                                                        v20 := GenS1{F0: -68, F1: 1, F2: 346}
                                                        var v21 bool = (true && true)
                                                        fmt.Println(v4, v4, v4)
                                                        _ = v20
                                                        _ = v21
                                                    }
                                                    var v22 float32 = 0.1
                                                    var v23 int = fi7
                                                    switch "hello" {
                                                    case "hello":
                                                        fmt.Println(v23, v22)
                                                    }
                                                    if (18446744073709551615 <= 42) {
                                                        var v24 string = "hello"
                                                        _ = v24
                                                    } else {
                                                        fmt.Println(v11, fi8)
                                                        switch "hello" {
                                                        case "foo bar":
                                                            v11 = 32768
                                                        case "":
                                                            delete(v2, (v5 & 2147483647))
                                                        case "hello":
                                                            delete(v2, 65535)
                                                        default:
                                                            v25 := GenS2{F0: (42 >= 2)}
                                                            _ = v25
                                                        }
                                                        var v26 bool = true
                                                        _ = v26
                                                    }
                                                    _ = v22
                                                    _ = v23
                                                }
                                            }
                                            v27 := GenS1{F0: 683, F1: 525, F2: (-(v4))}
                                            _ = v27
                                        }
                                    case "\n\t":
                                        v5 -= (v5 >> ((42) & 63))
                                        var v28 int64 = (v4 % ((int64(v5)) | 1))
                                        fmt.Println(fi8, fi8)
                                        _ = v28
                                    }
                                    _ = v19
                                } else {
                                    switch fi8 {
                                    case 0:
                                        switch (fi7 / ((v11) | 1)) {
                                        case -9223372036854775808:
                                            if false {
                                                _ = len(v2)
                                            } else {
                                                switch "hello" {
                                                case "foo bar":
                                                    var v29 int = 128
                                                    _ = v29
                                                case "hello":
                                                    var v30 uint16 = 42
                                                    v4 = 904
                                                    v31 := GenS2{F0: (0 != int16(v30))}
                                                    _ = v30
                                                    _ = v31
                                                case "\n\t":
                                                    var v32 uint16 = 1
                                                    switch -508 {
                                                    case 9223372036854775807:
                                                        v11 += (fi7 * fi7)
                                                    case 42:
                                                        var v33 int32 = 2147483647
                                                        _ = v33
                                                    }
                                                    v34 := GenS2{F0: (!("foo bar" <= ("\n\t" + "a")))}
                                                    var v35 uint16 = (v32 | v32)
                                                    switch (-(fi7)) {
                                                    case 0:
                                                        v35 = (v35 / ((v32) | 1))
                                                    }
                                                    _ = v32
                                                    _ = v34
                                                    _ = v35
                                                }
                                                if len(v3) > 0 {
                                                    v3[0] = (-(fi7))
                                                }
                                                var v36 uint = 4294967295
                                                _ = v36
                                            }
                                        case 0:
                                            switch "foo bar" {
                                            case "hello":
                                                v5 = (v5 >> ((1) & 63))
                                            default:
                                                v5 = (^(v5))
                                            }
                                        }
                                    case -865:
                                        var v37 int = (fi7 + v11)
                                        _ = v37
                                    case 9223372036854775807:
                                        var v38 int8 = 105
                                        _ = v38
                                    }
                                }
                                _ = v11
                            } else {
                                v5 = (v5 | (v5 - (v5 | (v5 ^ (v5 + v5)))))
                                var v39 int16 = 32767
                                _ = v39
                            }
                        }
                        v4 = v4
                        v4 += v4
                        if len(v3) > 0 {
                            v3[0] = 1
                        }
                    }
                } else {
                    if (!true) {
                        fmt.Println(v4)
                        fmt.Println(fi7)
                        fmt.Println(fi8, v4, fi8)
                        v40 := GenS2{F0: (!(false && (!true)))}
                        _ = v40
                    } else {
                        fmt.Println(v5, fi8, v5)
                        v3 = append(v3, (fi7 >> ((255) & 63)))
                        switch (("hello" + ("\n\t" + "\n\t")) + ("hello" + "a")) {
                        case "a":
                            var v41 uint32 = (v5 / ((v5) | 1))
                            _ = v41
                        case "hello":
                            if (((0.0 - 1.5) <= -2.25) || false) {
                                fmt.Println(fi8, v4)
                                _ = len(v3)
                            } else {
                                if ((598 != v5) || false) {
                                    if true {
                                        switch fi8 {
                                        case -440:
                                            if false {
                                                var v42 int64 = v4
                                                switch "" {
                                                case "foo bar":
                                                    _ = len(v3)
                                                }
                                                _ = v42
                                            } else {
                                                v4 += (^(v4))
                                            }
                                            if true {
                                                v2[v5] = (v5 & (^(v5)))
                                                v43 := GenS3{F0: 18446744073709551615}
                                                v4 = (v4 >> ((18446744073709551615) & 63))
                                                _ = v43
                                            } else {
                                                if false {
                                                    var v44 float64 = 0.1
                                                    _ = v44
                                                }
                                                switch (fi7 >> ((4294967295) & 63)) {
                                                case 0:
                                                    v5 = v5
                                                case 1:
                                                    v3 = append(v3, 0)
                                                case 587:
                                                    _ = len(v2)
                                                }
                                                v2[(v5 - v5)] = v5
                                                var v45 uint8 = 128
                                                _ = v45
                                            }
                                        case 2147483647:
                                            v46 := GenS2{F0: (v4 >= v4)}
                                            _ = v46
                                        case -791:
                                            v47 := GenS4{F0: 42, F1: 786, F2: v4}
                                            _ = v47
                                        }
                                    }
                                }
                            }
                        case "\n\t":
                            var v48 int32 = 0
                            _ = v48
                        default:
                            v49 := v2[7]
                            if true {
                                v49 -= (v5 | (^(v49)))
                                var v50 uint8 = 2
                                var v51 string = "foo bar"
                                _ = v50
                                _ = v51
                            }
                            switch fi7 {
                            case 128:
                                if (!true) {
                                    var v52 bool = (!true)
                                    _ = v52
                                } else {
                                    var v53 uint32 = 354
                                    fmt.Println(v53, v53)
                                    var v54 int64 = v4
                                    v55 := GenS2{F0: (42 > 723)}
                                    _ = v53
                                    _ = v54
                                    _ = v55
                                }
                            case 42:
                                if (false || false) {
                                    if (4294967295 <= 2) {
                                        if (1 <= -1) {
                                            if true {
                                                if true {
                                                    var v56 uint16 = 42
                                                    v5 -= 2147483647
                                                    v49 = (v5 % (((v5 >> ((4294967295) & 63))) | 1))
                                                    if (315 < 1) {
                                                        if ((false || false) || false) {
                                                            v4 -= (^(v4))
                                                            switch (fi8 & fi8) {
                                                            case 32768:
                                                                var v57 float32 = (-2.25 - 3.14159)
                                                                _ = v57
                                                            case 127:
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    var v58 int8 = 100
                                                                    _ = len(v2)
                                                                    _ = v58
                                                                case "hello":
                                                                    switch "foo bar" {
                                                                    case "\n\t":
                                                                        var v59 bool = true
                                                                        _ = v59
                                                                    case "":
                                                                        var v60 int = (fi8 << ((42) & 63))
                                                                        _ = v60
                                                                    default:
                                                                        if ((fi8 | fi7) != 128) {
                                                                            fmt.Println(fi7, v49, v5)
                                                                        } else {
                                                                            switch "foo bar" {
                                                                            case "\n\t":
                                                                                v49 += v5
                                                                            case "":
                                                                                if false {
                                                                                    var v61 float32 = (1.5 / -2.25)
                                                                                    v49 = v49
                                                                                    _ = v61
                                                                                } else {
                                                                                    var v62 int32 = -2147483648
                                                                                    v63 := GenS5{F0: 4294967295, F1: uint16(v5), F2: false}
                                                                                    v62 -= (v62 ^ (v62 | v62))
                                                                                    _ = v62
                                                                                    _ = v63
                                                                                }
                                                                            case "a":
                                                                                var v64 float64 = 1.5
                                                                                _ = v64
                                                                            default:
                                                                                v65 := GenS2{F0: (0 >= 42)}
                                                                                _ = v65
                                                                            }
                                                                            v1[fi8] = fi8
                                                                        }
                                                                    }
                                                                case "foo bar":
                                                                    if (!(!(!true))) {
                                                                        if (true || true) {
                                                                            v2[v49] = (v49 >> ((65535) & 63))
                                                                        } else {
                                                                            fmt.Println(v56, v5)
                                                                        }
                                                                        v5 = 2147483648
                                                                        if true {
                                                                            v49 = 2
                                                                            var v66 uint64 = 735
                                                                            switch "\n\t" {
                                                                            case "":
                                                                                v5 -= 0
                                                                            case "\n\t":
                                                                                if false {
                                                                                    var v67 string = ""
                                                                                    v49 = 2147483648
                                                                                    v68 := GenS6{F0: 127}
                                                                                    v69 := GenS5{F0: -9223372036854775808, F1: 42, F2: (0 < 2)}
                                                                                    v67 = "foo bar"
                                                                                    _ = v67
                                                                                    _ = v68
                                                                                    _ = v69
                                                                                } else {
                                                                                    fmt.Println(v49, v5)
                                                                                    var v70 float64 = 1e10
                                                                                    _ = v70
                                                                                }
                                                                            }
                                                                            _ = v66
                                                                        }
                                                                    }
                                                                default:
                                                                    if (!true) {
                                                                        v71 := v1[(fi7 + fi7)]
                                                                        if true {
                                                                            v1[fi8] = (fi7 >> ((4294967295) & 63))
                                                                        } else {
                                                                            var v72 int = fi8
                                                                            _ = v72
                                                                        }
                                                                        v56 = 65535
                                                                        _ = v71
                                                                    }
                                                                }
                                                            case 4294967295:
                                                                var v73 int32 = -1
                                                                _ = v73
                                                            default:
                                                                var v74 uint16 = (v56 >> ((33) & 63))
                                                                _ = v74
                                                            }
                                                        } else {
                                                            var v75 float32 = 1e10
                                                            fmt.Println(v5, v5)
                                                            _ = v75
                                                        }
                                                    } else {
                                                        _ = len(v2)
                                                    }
                                                    _ = v56
                                                } else {
                                                    var v76 uint8 = 1
                                                    var v77 string = ""
                                                    _ = v76
                                                    _ = v77
                                                }
                                                if false {
                                                    delete(v1, fi8)
                                                }
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        } else {
                                            var v78 uint8 = 42
                                            switch 282 {
                                            case 2147483647:
                                                var v79 bool = true
                                                _ = v79
                                            case 525:
                                                var v80 int32 = 0
                                                _ = v80
                                            case 32767:
                                                fmt.Println(v78, v5)
                                            default:
                                                v81 := v2[(v5 % ((v5) | 1))]
                                                _ = v81
                                            }
                                            v78 += (v78 & (v78 / ((v78) | 1)))
                                            var v82 uint32 = (v49 | v5)
                                            var v83 float64 = (1.5 - (0.0 / 3.14159))
                                            v4 = (v4 * (v4 - (v4 | v4)))
                                            if (v78 >= (v78 >> ((42) & 63))) {
                                                switch ((("foo bar" + ("" + "")) + "hello") + "hello") {
                                                case "hello":
                                                    if ("foo bar" < (("a" + "\n\t") + ("foo bar" + ("a" + "")))) {
                                                        switch ("a" + ("\n\t" + "a")) {
                                                        case "":
                                                            fmt.Println(v78)
                                                        case "a":
                                                            fmt.Println(v83, v4)
                                                        default:
                                                            if (v4 < (v4 >> ((1) & 63))) {
                                                                v78 = v78
                                                                fmt.Println(v83, v78)
                                                            } else {
                                                                v82 -= 2147483648
                                                            }
                                                        }
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        v3 = append(v3, (^(fi8)))
                                                    }
                                                }
                                                _ = len(v3)
                                                var v84 uint16 = 65535
                                                _ = v84
                                            } else {
                                                v83 -= -2.25
                                            }
                                            var v85 uint = 65535
                                            var v86 int8 = -1
                                            if (v85 <= v85) {
                                                v87 := GenS7{F0: ((!(1000 == 32767)) && (false || (!(((("" + "\n\t") + "hello") <= ("" + "foo bar")) && true)))), F1: -9223372036854775808, F2: 479}
                                                var v88 float64 = (v83 / (v83 + v83))
                                                _ = v87
                                                _ = v88
                                            } else {
                                                delete(v2, v82)
                                                v89 := GenS8{F0: v86, F1: ("hello" >= ("foo bar" + "foo bar")), F2: 42}
                                                var v90 uint = (v85 & (v85 >> ((0) & 63)))
                                                _ = v89
                                                _ = v90
                                            }
                                            switch (fi8 & 42) {
                                            case 938:
                                                v91 := GenS9{F0: (v78 / ((2) | 1)), F1: int16(v4)}
                                                _ = v91
                                            }
                                            _ = v78
                                            _ = v82
                                            _ = v83
                                            _ = v85
                                            _ = v86
                                        }
                                    }
                                    fmt.Println(v4, fi7)
                                    if false {
                                        var v92 float64 = ((0.0 / 0.1) + 1.5)
                                        v92 += (v92 * -2.25)
                                        _ = v92
                                    } else {
                                        var v93 float32 = (-2.25 - -2.25)
                                        var v94 uint16 = 2
                                        _ = v93
                                        _ = v94
                                    }
                                    if (false && false) {
                                        if true {
                                            v5 = v5
                                            if false {
                                                fmt.Println(v49)
                                                v4 += (v4 << ((282) & 63))
                                                v95 := GenS3{F0: 2}
                                                _ = v95
                                            }
                                            if true {
                                                fmt.Println(fi8, fi8, v5)
                                                var v96 uint64 = 42
                                                _ = v96
                                            } else {
                                                if (int(v4) <= (fi7 % (((fi7 ^ -734)) | 1))) {
                                                    switch fi7 {
                                                    case -658:
                                                        v4 = 301
                                                    case -2147483648:
                                                        v97 := GenS10{F0: -112}
                                                        v4 += (v4 ^ -9223372036854775808)
                                                        _ = v97
                                                    case -1:
                                                        var v98 float64 = -2.25
                                                        _ = v98
                                                    default:
                                                        var v99 uint32 = (v5 - 65535)
                                                        _ = v99
                                                    }
                                                    _ = len(v3)
                                                    fmt.Println(v4)
                                                }
                                            }
                                        } else {
                                            var v100 float64 = (1.5 + 1.5)
                                            v101 := v2[v49]
                                            _ = v100
                                            _ = v101
                                        }
                                    }
                                    v102 := GenS7{F0: (!(!true)), F1: (v4 / (((v4 * 1)) | 1)), F2: 600}
                                    _ = v102
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v49 = v49
                                v5 += 830
                                if true {
                                    v103 := v2[(^(v5))]
                                    var v104 int = (fi8 ^ fi7)
                                    _ = v103
                                    _ = v104
                                } else {
                                    var v105 int8 = -128
                                    _ = len(v1)
                                    var v106 int64 = (v4 >> ((42) & 63))
                                    _ = v105
                                    _ = v106
                                }
                            }
                            _ = v49
                        }
                    }
                    _ = len(v3)
                }
                var v107 uint64 = 18446744073709551615
                _ = fi8
                _ = v107
            }
            var v108 int32 = -2147483648
            _ = fi7
            _ = v108
        }
    case 665:
        fmt.Println(v4)
    default:
        if true {
            var v109 int = 42
            var v110 bool = false
            _ = len(v3)
            _ = v109
            _ = v110
        }
    }
    v111 := GenS11{F0: v4, F1: v5, F2: 128}
    v112 := GenS12{F0: 1, F1: -1, F2: 255}
    switch -768 {
    case 2147483647:
        if false {
            v113 := GenS13{F0: 0, F1: 0, F2: 18446744073709551615}
            var v114 uint = uint(v5)
            switch 0 {
            case -711:
                var v115 uint8 = 1
                _ = v115
            case -9223372036854775808:
                v116 := GenS14{F0: 0, F1: 996, F2: -65}
                _ = v116
            default:
                fmt.Println(v114, v114, v114)
            }
            _ = v113
            _ = v114
        }
    }
    v4 = -1
    v117 := GenS13{F0: 75, F1: 2, F2: 4294967295}
    var v118 int = 9223372036854775807
    v3 = append(v3, v118)
    v4 = v4
    v119 := v1[(v118 ^ v118)]
    fmt.Println(v5)
    _ = len(v3)
    var v120 float64 = 0.0
    var v121 int8 = 0
    switch 1 {
    case 42:
        switch v119 {
        case 325:
            switch "\n\t" {
            case "\n\t":
                v5 = (v5 & v5)
            case "":
                if (((0.1 - (3.14159 - (1.5 * 0.0))) / 1e10) < 1e10) {
                    v1[(v119 | (v119 * v119))] = v118
                    if (!true) {
                        fmt.Println(v118, v5, v118)
                        v122 := GenS13{F0: 1, F1: 255, F2: 2}
                        _ = v122
                    }
                }
            }
        default:
            fmt.Println(v119, v118, v120)
        }
    case 4294967295:
        if (674 >= 255) {
            for fi123 := 0; fi123 < 6; fi123++ {
                if true {
                    _ = len(v3)
                }
                if (!(65535 == 122)) {
                    v5 += v5
                    v124 := GenS15{F0: 42}
                    _ = v124
                } else {
                    v125 := GenS16{F0: (1 == 255), F1: 65535, F2: true}
                    _ = v125
                }
                if (448 <= 0) {
                    fmt.Println(v120, v5, v4)
                    switch "\n\t" {
                    case "\n\t":
                        v5 = 834
                    default:
                        if (float64(1e10) <= v120) {
                            _ = len(v3)
                            v126 := GenS5{F0: (^(v118)), F1: 42, F2: ("a" < "")}
                            _ = v126
                        }
                    }
                } else {
                    switch (v119 ^ (v119 % (((fi123 * 42)) | 1))) {
                    case 1:
                        for fi127 := 0; fi127 < 14; fi127++ {
                            v4 -= (v4 / (((v4 >> ((65535) & 63))) | 1))
                            _ = fi127
                        }
                    case 9223372036854775807:
                        switch "a" {
                        case "":
                            v128 := GenS17{F0: 731, F1: "hello"}
                            _ = v128
                        case "\n\t":
                            var v129 bool = false
                            _ = v129
                        case "hello":
                            fmt.Println(v118, v4, v5)
                        default:
                            var v130 float64 = v120
                            _ = v130
                        }
                    case 275:
                        v119 = v119
                    }
                }
                _ = fi123
            }
            v120 = v120
            var v131 string = ("a" + ("\n\t" + "\n\t"))
            _ = v131
        } else {
            if (!(((0.1 * 1.5) == 3.14159) && (508 != 2147483647))) {
                var v132 uint16 = 0
                _ = v132
            }
            fmt.Println(v119)
            fmt.Println(v120, v5)
        }
    case -551:
        if (!(321 != uint64(v5))) {
            for fi133 := 0; fi133 < 15; fi133++ {
                switch "a" {
                case "":
                    for fi134 := 0; fi134 < 8; fi134++ {
                        v135 := v1[(fi134 << ((114) & 63))]
                        v118 -= 0
                        var v136 int16 = -32768
                        v4 -= (v4 | int64(v118))
                        _ = fi134
                        _ = v135
                        _ = v136
                    }
                case "foo bar":
                    v5 = (v5 ^ v5)
                case "\n\t":
                    v4 -= (v4 - (v4 & (v4 % ((v4) | 1))))
                }
                _ = fi133
            }
        }
    default:
        var v137 uint = 0
        _ = v137
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v118:", v118)
    fmt.Println("v119:", v119)
    fmt.Println("v120:", v120)
    fmt.Println("v121:", v121)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v111
    _ = v112
    _ = v117
    _ = v118
    _ = v119
    _ = v120
    _ = v121
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
