package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 uint32; F1 int64; F2 int16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint8; F1 uint8; F2 string }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint; F1 uint8 }
type GenS12 struct { F0 int; F1 uint }
type GenS13 struct { F0 int32; F1 int32; F2 int64 }
type GenS14 struct { F0 uint8; F1 uint32 }
type GenS15 struct { F0 uint32; F1 uint8 }
type GenS16 struct { F0 int16; F1 int; F2 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = -19
    var v3 bool = (0 == -9223372036854775808)
    v4 := make(map[uint32]string)
    v5 := make(map[int]uint32)
    v6 := []string{""}
    switch -2147483648 {
    case 128:
        if true {
            if false {
                v7 := GenS1{F0: 255, F1: ("\n\t" + "a"), F2: 42}
                v3 = false
                _ = v7
            }
        } else {
            v3 = (-789 >= 498)
        }
    default:
        v3 = ((!(2147483647 == 9223372036854775807)) && v3)
    }
    v8 := GenS1{F0: 127, F1: "\n\t", F2: 440}
    v9 := GenS1{F0: 2, F1: "\n\t", F2: 849}
    var v10 bool = (((!v3) || false) || false)
    v2 += 42
    v11 := GenS1{F0: 2, F1: "\n\t", F2: 376}
    if v3 {
        for fi12 := 0; fi12 < 6; fi12++ {
            switch (-(fi12)) {
            case 32768:
                v13 := GenS1{F0: 128, F1: (("a" + "hello") + "hello"), F2: 0}
                _ = v13
            case 9223372036854775807:
                v3 = v3
            }
            var v14 uint8 = 122
            if false {
                var v15 float32 = -2.25
                for fi16 := 0; fi16 < 8; fi16++ {
                    var v17 float64 = 1e10
                    fmt.Println(v14, v10)
                    v2 += 1
                    _ = fi16
                    _ = v17
                }
                _ = v15
            }
            _ = fi12
            _ = v14
        }
        fmt.Println(v10, v2, v3)
        var v18 int16 = 0
        _ = v18
    } else {
        var v19 int16 = -32768
        v3 = v3
        v2 = v2
        _ = v19
    }
    if true {
        if (374 >= 0) {
            v3 = false
            v4[65535] = (("" + "a") + ("hello" + "foo bar"))
        }
        if v3 {
            if false {
                _ = len(v6)
            } else {
                for fi20 := 0; fi20 < 3; fi20++ {
                    v3 = (v10 || (1 != 65535))
                    v2 += (v2 ^ (v2 + 42))
                    _ = fi20
                }
                var v21 uint8 = 128
                _ = v21
            }
            switch 2147483647 {
            case 127:
                if v3 {
                    v6 = append(v6, ("a" + ""))
                } else {
                    for fi22 := 0; fi22 < 18; fi22++ {
                        var v23 uint64 = 2
                        if (1 < (^(v23))) {
                            v23 = v23
                            switch (fi22 * (fi22 & fi22)) {
                            case 9223372036854775807:
                                v23 += (v23 ^ (^(v23)))
                            case 42:
                                for fi24 := 0; fi24 < 19; fi24++ {
                                    v25 := GenS1{F0: 128, F1: "hello", F2: 853}
                                    var v26 uint16 = 255
                                    v27 := GenS2{F0: (v23 % (((^(v23))) | 1))}
                                    _ = fi24
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                }
                            default:
                                var v28 int16 = 766
                                _ = v28
                            }
                        } else {
                            for fi29 := 0; fi29 < 14; fi29++ {
                                if v10 {
                                    switch fi29 {
                                    case 9223372036854775807:
                                        if true {
                                            var v30 uint32 = 4294967295
                                            if len(v6) > 0 {
                                                v6[0] = ("\n\t" + ("\n\t" + ((("a" + "hello") + "hello") + "")))
                                            }
                                            _ = v30
                                        }
                                    }
                                    var v31 int32 = -203
                                    var v32 int32 = (v31 % ((v31) | 1))
                                    if false {
                                        v10 = (v3 && false)
                                        if v10 {
                                            v32 += v32
                                        } else {
                                            if v10 {
                                                var v33 bool = true
                                                if (!(255 != 255)) {
                                                    v32 = (v31 / ((v32) | 1))
                                                    v33 = (!("\n\t" == "\n\t"))
                                                    if ("" > "foo bar") {
                                                        var v34 int64 = 1
                                                        _ = v34
                                                    }
                                                }
                                                _ = v33
                                            } else {
                                                v10 = (!v3)
                                                var v35 int32 = 941
                                                var v36 float32 = (1e10 / (3.14159 + 0.0))
                                                _ = v35
                                                _ = v36
                                            }
                                            var v37 uint32 = 2147483648
                                            if ((253 >= 65535) || false) {
                                                fmt.Println(fi29, v3)
                                                var v38 uint8 = 0
                                                v6 = append(v6, "hello")
                                                _ = v38
                                            } else {
                                                switch 4294967295 {
                                                case -9223372036854775808:
                                                    if (!(5 < fi29)) {
                                                        fmt.Println(v10)
                                                    }
                                                case -487:
                                                    v31 = v31
                                                case 127:
                                                    if v3 {
                                                        switch (fi22 * (-(fi22))) {
                                                        case 278:
                                                            if true {
                                                                v37 += v37
                                                            }
                                                        case 1:
                                                            var v39 uint16 = uint16(v37)
                                                            _ = v39
                                                        }
                                                        v23 = 43
                                                        if true {
                                                            if (((1e10 > -2.25) && false) && v10) {
                                                                v1[v37] = fi29
                                                                var v40 int64 = 0
                                                                v32 -= v31
                                                                _ = v40
                                                            }
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            var v41 uint8 = 0
                                                            if v10 {
                                                                v23 = v23
                                                                v3 = (!(v23 > (v23 - v23)))
                                                                var v42 bool = v10
                                                                _ = v42
                                                            } else {
                                                                v6 = append(v6, ("a" + "foo bar"))
                                                                var v43 uint64 = v23
                                                                _ = len(v5)
                                                                switch (("foo bar" + "hello") + "") {
                                                                case "":
                                                                    v44 := GenS2{F0: 4294967295}
                                                                    _ = v44
                                                                default:
                                                                    var v45 float64 = (float64(1e10) - ((1.5 + 3.14159) - 1.5))
                                                                    _ = v45
                                                                }
                                                                _ = v43
                                                            }
                                                            _ = v41
                                                        }
                                                        switch (fi22 ^ fi29) {
                                                        case -246:
                                                            if v3 {
                                                                var v46 uint = 2
                                                                var v47 uint16 = 2
                                                                var v48 uint16 = v47
                                                                _ = v46
                                                                _ = v47
                                                                _ = v48
                                                            } else {
                                                                v49 := GenS1{F0: 42, F1: "\n\t", F2: v37}
                                                                var v50 float32 = (1e10 * 0.1)
                                                                if ("\n\t" != ("" + "")) {
                                                                    _ = len(v1)
                                                                    delete(v5, fi22)
                                                                    var v51 int32 = v31
                                                                    _ = v51
                                                                }
                                                                _ = v49
                                                                _ = v50
                                                            }
                                                            var v52 int = 294
                                                            _ = v52
                                                        default:
                                                            _ = len(v5)
                                                        }
                                                        switch fi22 {
                                                        case 128:
                                                            fmt.Println(v31)
                                                        case -2147483648:
                                                            if false {
                                                                v53 := GenS1{F0: 2, F1: ("a" + ((((("a" + "") + (("\n\t" + "hello") + ("foo bar" + "hello"))) + ("a" + ("hello" + "hello"))) + ("a" + "\n\t")) + "\n\t")), F2: (v37 & (^(v37)))}
                                                                delete(v5, int(v23))
                                                                var v54 float64 = (float64(1.5) * 0.1)
                                                                _ = v53
                                                                _ = v54
                                                            } else {
                                                                if len(v6) > 0 {
                                                                    v6[0] = ("foo bar" + "")
                                                                }
                                                                if true {
                                                                    var v55 uint = 783
                                                                    var v56 string = "a"
                                                                    _ = v55
                                                                    _ = v56
                                                                }
                                                                var v57 uint = 35
                                                                _ = v57
                                                            }
                                                        case 4294967295:
                                                            fmt.Println(v31, v23)
                                                        default:
                                                            if true {
                                                                fmt.Println(v37, v37, v10)
                                                                v2 -= 50
                                                                v10 = ((v32 % ((v31) | 1)) <= v32)
                                                                var v58 float64 = (((0.0 / 1e10) + 1.5) / 1.0)
                                                                _ = v58
                                                            } else {
                                                                if true {
                                                                    switch (-(fi29)) {
                                                                    case 128:
                                                                        v37 -= v37
                                                                    }
                                                                    v59 := GenS2{F0: (^(v23))}
                                                                    _ = v59
                                                                }
                                                            }
                                                        }
                                                    }
                                                }
                                                switch "" {
                                                case "":
                                                    v60 := v1[1]
                                                    _ = v60
                                                case "\n\t":
                                                    if (-32768 < int16(v2)) {
                                                        var v61 bool = (2 == (^(v23)))
                                                        v32 = v32
                                                        if v61 {
                                                            var v62 int16 = int16(v23)
                                                            _ = v62
                                                        }
                                                        _ = v61
                                                    }
                                                }
                                            }
                                            _ = v37
                                        }
                                        switch fi29 {
                                        case 4294967295:
                                            switch ("hello" + "a") {
                                            case "hello":
                                                _ = len(v4)
                                            case "":
                                                var v63 float64 = 1e10
                                                _ = v63
                                            default:
                                                switch "a" {
                                                case "hello":
                                                    if (287 >= 752) {
                                                        v3 = false
                                                        if v10 {
                                                            delete(v1, 2)
                                                            if len(v6) > 0 {
                                                                v6[0] = ""
                                                            }
                                                        } else {
                                                            v64 := v5[420]
                                                            v2 += (^(v2))
                                                            _ = v64
                                                        }
                                                        if ((fi22 % (((fi22 * (-(fi29)))) | 1)) >= (fi29 % (((fi29 % ((fi22) | 1))) | 1))) {
                                                            if (1.5 != 1e10) {
                                                                if ((2147483648 < 2) && v3) {
                                                                    var v65 uint32 = 42
                                                                    v31 = v31
                                                                    _ = v65
                                                                }
                                                            } else {
                                                                v32 = v32
                                                            }
                                                        }
                                                        if ((v23 | (^(v23))) >= (^(v23))) {
                                                            var v66 bool = (false && (3.14159 <= -2.25))
                                                            v67 := GenS2{F0: 646}
                                                            v31 -= -2147483648
                                                            _ = v66
                                                            _ = v67
                                                        } else {
                                                            if (true || (1.5 < 3.14159)) {
                                                                var v68 uint16 = 2
                                                                _ = v68
                                                            }
                                                            v1[42] = (fi22 - 0)
                                                            v3 = v3
                                                        }
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                    }
                                                case "a":
                                                    var v69 float32 = 1.5
                                                    _ = v69
                                                }
                                            }
                                        case 42:
                                            fmt.Println(v10, v32)
                                        case -9223372036854775808:
                                            v3 = v3
                                        default:
                                            v23 -= (v23 + v23)
                                        }
                                    }
                                    _ = v31
                                    _ = v32
                                }
                                if true {
                                    v70 := GenS2{F0: v23}
                                    if v3 {
                                        var v71 uint16 = 723
                                        var v72 uint8 = 127
                                        _ = v71
                                        _ = v72
                                    } else {
                                        fmt.Println(fi29, v10)
                                        var v73 int32 = 2147483647
                                        v74 := GenS3{F0: 1000, F1: 65535}
                                        _ = v73
                                        _ = v74
                                    }
                                    var v75 uint32 = 4294967295
                                    _ = v70
                                    _ = v75
                                }
                                v3 = true
                                switch "" {
                                case "":
                                    var v76 int16 = 454
                                    _ = v76
                                case "\n\t":
                                    if v3 {
                                        v77 := GenS3{F0: 42, F1: 233}
                                        var v78 int32 = 1
                                        if false {
                                            v10 = true
                                            var v79 uint32 = 596
                                            var v80 string = "\n\t"
                                            _ = v79
                                            _ = v80
                                        }
                                        _ = v77
                                        _ = v78
                                    }
                                }
                                _ = fi29
                            }
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                        for fi81 := 0; fi81 < 15; fi81++ {
                            if ((0.1 / 3.14159) < 1.5) {
                                switch fi22 {
                                case 4294967295:
                                    if true {
                                        switch ("foo bar" + "a") {
                                        case "hello":
                                            fmt.Println(v3, v10, fi22)
                                        case "\n\t":
                                            fmt.Println(fi22, v2, fi22)
                                        case "foo bar":
                                            v2 -= int8(v23)
                                        default:
                                            var v82 bool = v10
                                            v23 = (v23 >> ((uint(v23)) & 63))
                                            _ = v82
                                        }
                                        v2 = (v2 * (v2 >> ((336) & 63)))
                                    } else {
                                        var v83 bool = v10
                                        _ = v83
                                    }
                                case 9223372036854775807:
                                    v84 := GenS2{F0: v23}
                                    _ = v84
                                case 127:
                                    var v85 uint = 18446744073709551615
                                    _ = v85
                                }
                                v2 = v2
                            }
                            v86 := GenS4{F0: fi81}
                            v3 = (!v10)
                            _ = fi81
                            _ = v86
                        }
                        _ = fi22
                        _ = v23
                    }
                }
            case 9223372036854775807:
                if (42 < 975) {
                    switch "\n\t" {
                    case "":
                        fmt.Println(v2, v2)
                    }
                } else {
                    v87 := GenS5{F0: 65535, F1: 1}
                    switch ("hello" + "") {
                    case "a":
                        switch (("\n\t" + (("a" + ("foo bar" + ("foo bar" + "a"))) + "")) + "a") {
                        case "":
                            if v3 {
                                for fi88 := 0; fi88 < 1; fi88++ {
                                    var v89 int64 = 294
                                    for fi90 := 0; fi90 < 7; fi90++ {
                                        v91 := GenS1{F0: 201, F1: ("foo bar" + ""), F2: 899}
                                        _ = fi90
                                        _ = v91
                                    }
                                    for fi92 := 0; fi92 < 7; fi92++ {
                                        fmt.Println(v2)
                                        _ = fi92
                                    }
                                    _ = fi88
                                    _ = v89
                                }
                                var v93 string = "foo bar"
                                _ = v93
                            }
                        case "hello":
                            var v94 int8 = v2
                            _ = v94
                        case "a":
                            var v95 int16 = 1000
                            _ = v95
                        }
                    case "":
                        v96 := GenS6{F0: 1, F1: -833, F2: 818}
                        _ = v96
                    case "\n\t":
                        for fi97 := 0; fi97 < 11; fi97++ {
                            switch "" {
                            case "a":
                                fmt.Println(v10)
                            case "foo bar":
                                var v98 int32 = 0
                                _ = v98
                            default:
                                switch 32768 {
                                case 0:
                                    switch "\n\t" {
                                    case "hello":
                                        v10 = false
                                    case "foo bar":
                                        v99 := GenS7{F0: 185}
                                        _ = v99
                                    case "a":
                                        if v3 {
                                            for fi100 := 0; fi100 < 10; fi100++ {
                                                v101 := GenS8{F0: 130, F1: 151, F2: "hello"}
                                                v102 := GenS6{F0: 65535, F1: 42, F2: 1000}
                                                var v103 uint64 = 2
                                                _ = fi100
                                                _ = v101
                                                _ = v102
                                                _ = v103
                                            }
                                        } else {
                                            if (1.5 == (0.0 * -2.25)) {
                                                if ("\n\t" == "") {
                                                    v104 := GenS9{F0: 9223372036854775807}
                                                    var v105 uint32 = 563
                                                    v2 = (v2 * (v2 | v2))
                                                    _ = v104
                                                    _ = v105
                                                } else {
                                                    if v10 {
                                                        if v10 {
                                                            switch fi97 {
                                                            case -2147483648:
                                                                var v106 int8 = (v2 % (((v2 | v2)) | 1))
                                                                _ = v106
                                                            case 0:
                                                                switch fi97 {
                                                                case 32767:
                                                                    fmt.Println(v10, fi97, fi97)
                                                                case 551:
                                                                    var v107 float64 = (1e10 - float64(0.0))
                                                                    _ = v107
                                                                }
                                                            default:
                                                                delete(v1, 1)
                                                            }
                                                            for fi108 := 0; fi108 < 4; fi108++ {
                                                                v109 := GenS10{F0: v3}
                                                                _ = fi108
                                                                _ = v109
                                                            }
                                                            switch (fi97 & fi97) {
                                                            case -2147483648:
                                                                var v110 int8 = (v2 ^ 0)
                                                                _ = v110
                                                            default:
                                                                if (false && true) {
                                                                    var v111 uint16 = 255
                                                                    if (3.14159 > float64(1.5)) {
                                                                        for fi112 := 0; fi112 < 8; fi112++ {
                                                                            v113 := GenS11{F0: 0, F1: 1}
                                                                            _ = fi112
                                                                            _ = v113
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v2, v111, v3)
                                                                        v2 = (v2 % ((v2) | 1))
                                                                    }
                                                                    for fi114 := 0; fi114 < 14; fi114++ {
                                                                        _ = len(v5)
                                                                        v115 := GenS2{F0: 1}
                                                                        _ = fi114
                                                                        _ = v115
                                                                    }
                                                                    _ = v111
                                                                }
                                                            }
                                                        }
                                                        var v116 int32 = -12
                                                        _ = v116
                                                    } else {
                                                        for fi117 := 0; fi117 < 1; fi117++ {
                                                            if v10 {
                                                                fmt.Println(v2, fi117)
                                                            } else {
                                                                v118 := GenS4{F0: (fi117 | (fi97 * fi117))}
                                                                fmt.Println(v2, v3)
                                                                v119 := GenS9{F0: -1}
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    _ = len(v5)
                                                                case "":
                                                                    if ((("\n\t" + "\n\t") + ("" + "\n\t")) < "foo bar") {
                                                                        v1[1] = fi117
                                                                    } else {
                                                                        _ = len(v6)
                                                                        if (!true) {
                                                                            v3 = (fi97 != (fi117 >> ((65535) & 63)))
                                                                            var v120 bool = (3.14159 != (0.0 / float64(0.1)))
                                                                            v121 := v5[769]
                                                                            if true {
                                                                                v3 = v10
                                                                                v3 = true
                                                                                v10 = false
                                                                                switch "foo bar" {
                                                                                case "\n\t":
                                                                                    if true {
                                                                                        v122 := GenS2{F0: 18446744073709551615}
                                                                                        v121 += 64
                                                                                        var v123 float32 = 0.1
                                                                                        _ = v122
                                                                                        _ = v123
                                                                                    } else {
                                                                                        v124 := GenS1{F0: 232, F1: "\n\t", F2: (v121 % ((v121) | 1))}
                                                                                        var v125 int64 = -9223372036854775808
                                                                                        _ = v124
                                                                                        _ = v125
                                                                                    }
                                                                                case "a":
                                                                                    var v126 int32 = -2147483648
                                                                                    _ = v126
                                                                                case "hello":
                                                                                    v10 = v120
                                                                                }
                                                                                fmt.Println(fi97)
                                                                                v6 = append(v6, "foo bar")
                                                                            }
                                                                            v127 := GenS12{F0: fi97, F1: 159}
                                                                            _ = v120
                                                                            _ = v121
                                                                            _ = v127
                                                                        }
                                                                        if (-2.25 < 0.0) {
                                                                            v10 = v3
                                                                            if len(v6) >= 2 {
                                                                                v6 = v6[1:len(v6)]
                                                                            }
                                                                        } else {
                                                                            v128 := GenS13{F0: -2147483648, F1: -1, F2: 0}
                                                                            _ = v128
                                                                        }
                                                                    }
                                                                default:
                                                                    v2 += v2
                                                                }
                                                                var v129 int16 = 0
                                                                if (128 < 0) {
                                                                    v4[2147483648] = "hello"
                                                                    var v130 float32 = (1.5 * 1.5)
                                                                    _ = v130
                                                                } else {
                                                                    fmt.Println(v129, v10)
                                                                }
                                                                if false {
                                                                    if (15 > 42) {
                                                                        delete(v1, 1)
                                                                        v5[fi97] = 42
                                                                        v10 = v3
                                                                    }
                                                                }
                                                                switch fi117 {
                                                                case -890:
                                                                    var v131 string = "hello"
                                                                    _ = v131
                                                                case 9223372036854775807:
                                                                    v6 = append(v6, "a")
                                                                case 128:
                                                                    v10 = v10
                                                                }
                                                                fmt.Println(v129, v3, v129)
                                                                _ = v118
                                                                _ = v119
                                                                _ = v129
                                                            }
                                                            var v132 uint64 = 18446744073709551615
                                                            _ = fi117
                                                            _ = v132
                                                        }
                                                    }
                                                }
                                            } else {
                                                v10 = true
                                                var v133 string = "a"
                                                v134 := GenS2{F0: 74}
                                                _ = v133
                                                _ = v134
                                            }
                                        }
                                    }
                                default:
                                    if true {
                                        for fi135 := 0; fi135 < 13; fi135++ {
                                            v136 := GenS14{F0: 25, F1: 118}
                                            var v137 uint8 = 42
                                            if false {
                                                var v138 int8 = int8(fi97)
                                                v139 := GenS15{F0: 65535, F1: (v137 * (v137 << ((42) & 63)))}
                                                var v140 float64 = 3.14159
                                                _ = v138
                                                _ = v139
                                                _ = v140
                                            } else {
                                                v10 = v3
                                            }
                                            _ = fi135
                                            _ = v136
                                            _ = v137
                                        }
                                        if true {
                                            for fi141 := 0; fi141 < 6; fi141++ {
                                                v2 = v2
                                                _ = fi141
                                            }
                                        } else {
                                            fmt.Println(v2)
                                            var v142 string = "\n\t"
                                            _ = v142
                                        }
                                    }
                                }
                            }
                            for fi143 := 0; fi143 < 8; fi143++ {
                                v144 := GenS16{F0: 32767, F1: -9223372036854775808, F2: "foo bar"}
                                _ = fi143
                                _ = v144
                            }
                            var v145 uint = 994
                            _ = fi97
                            _ = v145
                        }
                    }
                    _ = v87
                }
            case -748:
                var v146 uint16 = 2
                _ = v146
            }
        }
        switch 0 {
        case 884:
            v147 := GenS11{F0: 18446744073709551615, F1: 1}
            _ = v147
        case 0:
            var v148 float64 = float64(float32((0.1 / 1.0)))
            _ = v148
        default:
            for fi149 := 0; fi149 < 3; fi149++ {
                var v150 uint8 = 128
                var v151 int32 = 42
                _ = fi149
                _ = v150
                _ = v151
            }
        }
    }
    if (("a" < "") || (v3 || (!true))) {
        v3 = v3
    }
    var v152 uint8 = 1
    if (true || v3) {
        switch "hello" {
        case "foo bar":
            var v153 uint8 = (v152 - v152)
            _ = v153
        case "\n\t":
            v2 -= v2
        }
    }
    v10 = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v152:", v152)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v152
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
