package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 int8; F2 int64 }
type GenS3 struct { F0 uint; F1 int16; F2 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int; F1 string }
type GenS10 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint16; F1 int32 }
type GenS13 struct { F0 int }

func main() {
    v1 := []string{("" + ("\n\t" + "\n\t")), "hello", "a"}
    var v2 int = 0
    v3 := []string{}
    v4 := make(map[string]bool)
    v5 := []string{(("a" + "a") + "\n\t"), "hello", ((("hello" + "") + "foo bar") + (("a" + "") + "a"))}
    v2 += v2
    if (-2147483648 >= -1) {
        var v6 uint8 = 127
        for fi7 := 0; fi7 < 9; fi7++ {
            v2 = 0
            v8 := GenS1{F0: 154}
            if (false && true) {
                switch 1 {
                case 32767:
                    for fi9 := 0; fi9 < 6; fi9++ {
                        var v10 int16 = 32767
                        _ = fi9
                        _ = v10
                    }
                case 127:
                    v6 = (v6 | 0)
                }
                if len(v5) > 0 {
                    v5[0] = ""
                }
            }
            _ = fi7
            _ = v8
        }
        _ = v6
    } else {
        fmt.Println(v2)
    }
    var v11 float64 = ((3.14159 * 1e10) / 3.14159)
    v12 := GenS1{F0: 0}
    if (true || (!true)) {
        fmt.Println(v11, v11)
        switch (v2 % ((-844) | 1)) {
        case -9223372036854775808:
            v13 := v4[(("\n\t" + "\n\t") + "foo bar")]
            _ = v13
        case 42:
            if (-1 <= v2) {
                if ((-1 >= 42) && true) {
                    if true {
                        var v14 string = "a"
                        var v15 float32 = 1.5
                        fmt.Println(v15, v2)
                        _ = v14
                        _ = v15
                    }
                    v2 = (^(v2))
                }
                if (1e10 >= 0.0) {
                    var v16 uint8 = 42
                    var v17 int8 = -9
                    if false {
                        var v18 float64 = 0.0
                        if true {
                            switch "\n\t" {
                            case "\n\t":
                                var v19 uint = 0
                                _ = v19
                            case "hello":
                                fmt.Println(v17)
                            default:
                                for fi20 := 0; fi20 < 12; fi20++ {
                                    v11 += (float64(1e10) + (float64((1.5 * 0.0)) * float64(1e10)))
                                    var v21 float32 = 3.14159
                                    _ = fi20
                                    _ = v21
                                }
                            }
                            for fi22 := 0; fi22 < 5; fi22++ {
                                for fi23 := 0; fi23 < 8; fi23++ {
                                    v4["a"] = ((float64(-2.25) / 0.1) == (0.1 + 3.14159))
                                    _ = fi23
                                }
                                _ = fi22
                            }
                            var v24 int64 = -9223372036854775808
                            _ = v24
                        } else {
                            var v25 uint = 988
                            var v26 float32 = 1.5
                            v27 := GenS1{F0: 65535}
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                        _ = v18
                    } else {
                        var v28 int8 = (v17 | (v17 * v17))
                        if len(v1) > 0 {
                            v1[0] = "hello"
                        }
                        if (true || true) {
                            for fi29 := 0; fi29 < 14; fi29++ {
                                v30 := GenS1{F0: 0}
                                if len(v1) > 0 {
                                    v1[0] = ""
                                }
                                _ = fi29
                                _ = v30
                            }
                            if len(v3) > 0 {
                                v3[0] = "\n\t"
                            }
                        } else {
                            for fi31 := 0; fi31 < 6; fi31++ {
                                for fi32 := 0; fi32 < 8; fi32++ {
                                    v28 = v17
                                    v28 -= (v17 | v17)
                                    _ = fi32
                                }
                                _ = fi31
                            }
                            for fi33 := 0; fi33 < 17; fi33++ {
                                v2 -= v2
                                v3 = append(v3, "\n\t")
                                _ = fi33
                            }
                        }
                        _ = v28
                    }
                    _ = v16
                    _ = v17
                } else {
                    if true {
                        v11 = 0.1
                        switch ("hello" + "\n\t") {
                        case "a":
                            if (false && (0 == 42)) {
                                if ((2 < 323) || false) {
                                    var v34 uint32 = 2
                                    var v35 int32 = -2147483648
                                    fmt.Println(v11, v2, v2)
                                    _ = v34
                                    _ = v35
                                } else {
                                    fmt.Println(v2, v2)
                                    fmt.Println(v2)
                                }
                                var v36 int = v2
                                if ((-1 >= 1) || (198 <= 1)) {
                                    fmt.Println(v2, v2, v11)
                                } else {
                                    switch ("foo bar" + "foo bar") {
                                    case "foo bar":
                                        v36 -= 247
                                    default:
                                        var v37 float64 = (v11 - v11)
                                        _ = v37
                                    }
                                    fmt.Println(v11, v2)
                                }
                                _ = v36
                            }
                        default:
                            if false {
                                var v38 int64 = 42
                                if false {
                                    v38 -= (v38 * v38)
                                } else {
                                    for fi39 := 0; fi39 < 10; fi39++ {
                                        if false {
                                            v40 := GenS1{F0: 65535}
                                            var v41 uint16 = 89
                                            var v42 bool = (-60 >= 272)
                                            _ = v40
                                            _ = v41
                                            _ = v42
                                        } else {
                                            fmt.Println(fi39)
                                        }
                                        _ = fi39
                                    }
                                }
                                _ = v38
                            } else {
                                var v43 int = 2147483647
                                switch (v2 / ((v2) | 1)) {
                                case -1:
                                    var v44 int64 = 9223372036854775807
                                    _ = v44
                                case 772:
                                    v11 -= v11
                                default:
                                    v2 = (v43 * (v43 % (((-(v43))) | 1)))
                                }
                                _ = v43
                            }
                        }
                        v45 := GenS1{F0: 460}
                        _ = v45
                    }
                    v2 -= 297
                    switch ("a" + ("foo bar" + ("a" + ""))) {
                    case "":
                        v11 = v11
                    }
                }
                if ((^(v2)) > (^(v2))) {
                    switch "a" {
                    case "hello":
                        switch v2 {
                        case 1:
                            var v46 int = (v2 * (v2 | v2))
                            _ = v46
                        }
                    case "a":
                        v11 += v11
                    case "":
                        var v47 uint64 = 4294967295
                        _ = v47
                    }
                    var v48 int = v2
                    v49 := GenS2{F0: 813, F1: 1, F2: 0}
                    _ = v48
                    _ = v49
                } else {
                    for fi50 := 0; fi50 < 3; fi50++ {
                        for fi51 := 0; fi51 < 17; fi51++ {
                            v11 = ((v11 + (float64(1.5) - ((v11 - v11) * v11))) - v11)
                            var v52 int = (^(v2))
                            _ = fi51
                            _ = v52
                        }
                        _ = fi50
                    }
                    v2 = -2147483648
                }
            } else {
                v11 += v11
            }
        case 0:
            fmt.Println(v2, v2)
        }
    } else {
        for fi53 := 0; fi53 < 10; fi53++ {
            _ = len(v4)
            _ = fi53
        }
        v54 := GenS2{F0: 218, F1: 1, F2: -393}
        v2 = (v2 << ((42) & 63))
        _ = v54
    }
    var v55 string = (((("\n\t" + ("a" + (("a" + "a") + (("" + ("hello" + "a")) + "")))) + "hello") + "a") + (("foo bar" + ("\n\t" + "")) + "foo bar"))
    for fi56 := 0; fi56 < 18; fi56++ {
        v57 := GenS2{F0: 78, F1: -62, F2: 313}
        _ = fi56
        _ = v57
    }
    if true {
        _ = len(v4)
        if true {
            if false {
                switch ("hello" + v55) {
                case "":
                    switch (((v55 + (v55 + v55)) + "hello") + (v55 + "foo bar")) {
                    case "a":
                        v2 += (v2 + 1)
                    case "foo bar":
                        v58 := GenS3{F0: 4294967295, F1: 115, F2: 32767}
                        _ = v58
                    case "":
                        v11 = v11
                    }
                case "a":
                    if (!false) {
                        v59 := GenS2{F0: 304, F1: 1, F2: 42}
                        _ = len(v4)
                        _ = v59
                    }
                }
                var v60 float64 = ((v11 * float64(0.0)) / 1e10)
                for fi61 := 0; fi61 < 4; fi61++ {
                    v55 = v55
                    _ = fi61
                }
                _ = v60
            } else {
                var v62 float64 = v11
                for fi63 := 0; fi63 < 10; fi63++ {
                    if len(v3) > 0 {
                        v3[0] = v55
                    }
                    v4[v55] = true
                    _ = fi63
                }
                _ = v62
            }
            for fi64 := 0; fi64 < 9; fi64++ {
                v65 := GenS4{F0: -32768}
                v66 := GenS4{F0: 1}
                var v67 int16 = -869
                _ = fi64
                _ = v65
                _ = v66
                _ = v67
            }
            v68 := GenS4{F0: -32768}
            _ = v68
        } else {
            var v69 float64 = v11
            for fi70 := 0; fi70 < 13; fi70++ {
                var v71 string = v55
                if true {
                    if false {
                        if false {
                            if true {
                                var v72 string = v55
                                v73 := GenS1{F0: 1}
                                _ = v72
                                _ = v73
                            } else {
                                for fi74 := 0; fi74 < 1; fi74++ {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v75 int32 = -640
                                    var v76 uint16 = 2
                                    _ = fi74
                                    _ = v75
                                    _ = v76
                                }
                                v4["\n\t"] = ((true || (42 > 1)) || (32767 != 2))
                                switch (fi70 - fi70) {
                                case 42:
                                    var v77 uint8 = 255
                                    _ = v77
                                case 4294967295:
                                    for fi78 := 0; fi78 < 13; fi78++ {
                                        switch fi70 {
                                        case 2147483647:
                                            v79 := GenS5{F0: v71}
                                            _ = v79
                                        case 0:
                                            if (v11 <= (float64(0.1) * ((v69 / ((v11 + float64((0.0 - (1e10 / ((1e10 / 3.14159) / 3.14159))))) / 1.0)) * (v69 - (0.1 / (v11 / float64(float32((v69 + ((v69 + v69) * v69)))))))))) {
                                                if true {
                                                    fmt.Println(v2, v69)
                                                } else {
                                                    v69 = ((((float64(1.5) / v11) / 1e10) - v69) / v69)
                                                    v2 = (v2 & 128)
                                                }
                                            }
                                        default:
                                            if false {
                                                switch (v55 + (("foo bar" + (((v71 + "\n\t") + ("" + (("" + ("\n\t" + ("" + v55))) + v55))) + v71)) + (v55 + v71))) {
                                                case "foo bar":
                                                    if (1 >= 18446744073709551615) {
                                                        v71 = v71
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        fmt.Println(fi78)
                                                    }
                                                }
                                                v80 := GenS6{F0: 42, F1: 1, F2: 1}
                                                if len(v3) > 0 {
                                                    v3[0] = v55
                                                }
                                                fmt.Println(v55, v69)
                                                if false {
                                                    var v81 float32 = (0.1 * 1.5)
                                                    _ = v81
                                                } else {
                                                    _ = len(v4)
                                                    switch fi70 {
                                                    case -22:
                                                        var v82 uint8 = 70
                                                        _ = v82
                                                    }
                                                }
                                                fmt.Println(fi70, fi78)
                                                if ((((true || (!true)) && true) || true) && false) {
                                                    if true {
                                                        v69 = float64(float32(v69))
                                                        if false {
                                                            v83 := GenS4{F0: -32768}
                                                            v71 = v71
                                                            var v84 int = (fi70 * 32768)
                                                            _ = v83
                                                            _ = v84
                                                        } else {
                                                            v69 += float64(0.1)
                                                        }
                                                    } else {
                                                        if len(v3) > 0 {
                                                            v3[0] = (v71 + v55)
                                                        }
                                                        v85 := GenS2{F0: 255, F1: 100, F2: 0}
                                                        _ = len(v4)
                                                        if false {
                                                            v86 := GenS6{F0: -1, F1: 130, F2: 1}
                                                            var v87 int32 = 1
                                                            _ = v86
                                                            _ = v87
                                                        }
                                                        fmt.Println(v55, v11, v55)
                                                        _ = v85
                                                    }
                                                    var v88 uint32 = 2147483648
                                                    var v89 int16 = -32768
                                                    _ = v88
                                                    _ = v89
                                                }
                                                _ = v80
                                            } else {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                        }
                                        v2 -= (fi78 & -700)
                                        _ = fi78
                                    }
                                case -648:
                                    v11 += 1e10
                                }
                            }
                        } else {
                            switch v71 {
                            case "hello":
                                if ((0.0 / (1.5 / 1.0)) > 3.14159) {
                                    if true {
                                        v69 = (1.5 - (v11 - 1e10))
                                        var v90 uint32 = 585
                                        _ = v90
                                    } else {
                                        if false {
                                            var v91 bool = true
                                            _ = v91
                                        }
                                    }
                                    v55 = v71
                                } else {
                                    v92 := v4[v71]
                                    for fi93 := 0; fi93 < 10; fi93++ {
                                        if (!((-2147483648 > -8) && v92)) {
                                            fmt.Println(v11)
                                            fmt.Println(v71)
                                            v92 = false
                                        }
                                        _ = fi93
                                    }
                                    var v94 uint32 = 4294967295
                                    _ = v92
                                    _ = v94
                                }
                            case "\n\t":
                                delete(v4, "\n\t")
                            }
                            v95 := v4[v55]
                            _ = v95
                        }
                        v96 := GenS4{F0: -381}
                        if false {
                            switch v2 {
                            case 0:
                                for fi97 := 0; fi97 < 17; fi97++ {
                                    v71 = (((v55 + "\n\t") + "hello") + "hello")
                                    fmt.Println(fi70, v69, v11)
                                    var v98 uint32 = 125
                                    var v99 float32 = 0.0
                                    v2 -= 0
                                    var v100 uint32 = 2147483648
                                    switch v55 {
                                    case "\n\t":
                                        v69 += (((v11 - (float64((v99 / v99)) + (v69 + 3.14159))) * float64(v99)) + v11)
                                    case "foo bar":
                                        var v101 int8 = 117
                                        _ = v101
                                    case "":
                                        v2 = (v2 >> ((65535) & 63))
                                    }
                                    _ = fi97
                                    _ = v100
                                    _ = v98
                                    _ = v99
                                }
                            case 4294967295:
                                for fi102 := 0; fi102 < 15; fi102++ {
                                    v69 = float64(-2.25)
                                    _ = fi102
                                }
                            case -1:
                                var v103 int8 = 127
                                _ = v103
                            default:
                                if false {
                                    fmt.Println(fi70, v2)
                                    for fi104 := 0; fi104 < 3; fi104++ {
                                        v2 += (-(v2))
                                        if false {
                                            _ = len(v4)
                                            switch v71 {
                                            case "\n\t":
                                                if true {
                                                    var v105 uint64 = 42
                                                    v106 := GenS7{F0: 42}
                                                    var v107 uint64 = (^(v105))
                                                    _ = v105
                                                    _ = v106
                                                    _ = v107
                                                } else {
                                                    if true {
                                                        _ = len(v4)
                                                        v108 := GenS8{F0: fi70}
                                                        _ = v108
                                                    } else {
                                                        v55 = v55
                                                    }
                                                    var v109 bool = false
                                                    switch v55 {
                                                    case "a":
                                                        v2 -= (v2 / ((1) | 1))
                                                    default:
                                                        v11 = 0.1
                                                        v2 -= 128
                                                        if len(v1) > 0 {
                                                            v1[0] = v55
                                                        }
                                                        switch (fi70 ^ fi70) {
                                                        case 9223372036854775807:
                                                            var v110 float32 = 0.0
                                                            _ = v110
                                                        case 32767:
                                                            var v111 float64 = v69
                                                            _ = v111
                                                        case 42:
                                                            switch (v2 ^ 32768) {
                                                            case 127:
                                                                var v112 float64 = (((v69 * (float64(0.0) * v69)) - (v11 * v69)) + v11)
                                                                _ = v112
                                                            case 1:
                                                                v3 = append(v3, "a")
                                                            default:
                                                                v69 = 0.0
                                                            }
                                                        }
                                                    }
                                                    fmt.Println(v109, v109)
                                                    switch fi104 {
                                                    case 9223372036854775807:
                                                        fmt.Println(v55)
                                                    case 667:
                                                        v69 = (float64(1e10) / v11)
                                                    default:
                                                        var v113 uint8 = 42
                                                        _ = v113
                                                    }
                                                    _ = v109
                                                }
                                            case "foo bar":
                                                v71 = v71
                                            }
                                        } else {
                                            var v114 int = 2147483647
                                            fmt.Println(fi104)
                                            _ = v114
                                        }
                                        switch "foo bar" {
                                        case "a":
                                            v115 := GenS9{F0: 32767, F1: (v71 + v71)}
                                            _ = v115
                                        default:
                                            v11 += v69
                                        }
                                        _ = fi104
                                    }
                                } else {
                                    v69 += (v69 + 1e10)
                                    _ = len(v4)
                                }
                            }
                            v2 -= (fi70 % ((9223372036854775807) | 1))
                        }
                        _ = v96
                    }
                } else {
                    var v116 uint = 18446744073709551615
                    if true {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        var v117 int16 = 32767
                        var v118 float32 = 0.0
                        _ = v117
                        _ = v118
                    }
                    v119 := GenS10{F0: -1, F1: 0, F2: 232}
                    _ = v116
                    _ = v119
                }
                var v120 int16 = 0
                _ = fi70
                _ = v120
                _ = v71
            }
            _ = v69
        }
        v4[""] = (65535 >= 0)
    }
    var v121 int32 = 2147483647
    for fi122 := 0; fi122 < 5; fi122++ {
        _ = len(v3)
        if true {
            if true {
                var v123 int8 = 1
                _ = v123
            }
            var v124 float64 = v11
            _ = v124
        }
        for fi125 := 0; fi125 < 2; fi125++ {
            if (100 > 42) {
                if (int16(v2) == 1) {
                    if false {
                        v121 = -1
                        if len(v5) > 0 {
                            v5[0] = v55
                        }
                        var v126 int16 = -778
                        _ = v126
                    } else {
                        var v127 int = (fi125 & (fi122 ^ (v2 | v2)))
                        v127 += (v127 * fi125)
                        var v128 int8 = 0
                        _ = v127
                        _ = v128
                    }
                    if ((!(367 == 570)) && false) {
                        switch fi125 {
                        case 724:
                            v11 = v11
                        }
                    } else {
                        switch ("" + v55) {
                        case "\n\t":
                            switch v2 {
                            case 2147483647:
                                if (false || false) {
                                    v129 := GenS6{F0: 32767, F1: 12, F2: 2}
                                    _ = v129
                                }
                            case -731:
                                if (!(true || true)) {
                                    v121 += (v121 | -2147483648)
                                    if true {
                                        v130 := GenS6{F0: 32767, F1: 127, F2: 108}
                                        _ = v130
                                    }
                                    v11 = (v11 + 1e10)
                                } else {
                                    v121 = (v121 << ((0) & 63))
                                }
                            }
                        case "foo bar":
                            var v131 float32 = 0.0
                            _ = v131
                        case "hello":
                            v121 = -795
                        default:
                            v132 := GenS11{F0: -59}
                            _ = v132
                        }
                    }
                    var v133 uint64 = 291
                    _ = v133
                } else {
                    switch fi125 {
                    case 4294967295:
                        v55 = v55
                    case -6:
                        if (173 == 1) {
                            fmt.Println(fi122, v55)
                            v2 -= v2
                        }
                    }
                    v4[v55] = (255 == 128)
                }
                _ = len(v4)
                var v134 int16 = 32767
                var v135 uint8 = 255
                v136 := GenS12{F0: 2, F1: (-(v121))}
                _ = v134
                _ = v135
                _ = v136
            }
            if len(v3) > 0 {
                v3[0] = v55
            }
            var v137 uint32 = 2
            _ = fi125
            _ = v137
        }
        _ = fi122
    }
    v121 += v121
    if (65535 <= 4294967295) {
        for fi138 := 0; fi138 < 19; fi138++ {
            v4[(v55 + "hello")] = true
            var v139 uint8 = 0
            var v140 string = ""
            _ = fi138
            _ = v139
            _ = v140
        }
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
        switch (v55 + (("\n\t" + "hello") + v55)) {
        case "foo bar":
            v141 := GenS11{F0: 0}
            _ = v141
        case "\n\t":
            if (!((!false) && ((2147483648 > 313) || (0.0 >= ((3.14159 - 3.14159) * (1e10 - ((1.5 * (-2.25 / 0.1)) - 3.14159))))))) {
                _ = len(v1)
                v142 := GenS7{F0: 1}
                var v143 int64 = -450
                _ = v142
                _ = v143
            }
        case "":
            switch v55 {
            case "foo bar":
                if (550 != 2) {
                    if ((("\n\t" + (v55 + v55)) + v55) > "hello") {
                        for fi144 := 0; fi144 < 4; fi144++ {
                            var v145 uint16 = 42
                            var v146 uint32 = 4294967295
                            _ = fi144
                            _ = v145
                            _ = v146
                        }
                        v147 := v4[(v55 + v55)]
                        fmt.Println(v11, v121)
                        _ = v147
                    } else {
                        v121 += -479
                        switch v2 {
                        case 32768:
                            var v148 uint32 = 2
                            _ = v148
                        }
                        for fi149 := 0; fi149 < 4; fi149++ {
                            v121 = v121
                            _ = fi149
                        }
                    }
                } else {
                    v11 -= v11
                }
            case "":
                for fi150 := 0; fi150 < 19; fi150++ {
                    if true {
                        _ = len(v4)
                        v151 := v4[v55]
                        _ = v151
                    }
                    _ = fi150
                }
            case "hello":
                for fi152 := 0; fi152 < 20; fi152++ {
                    for fi153 := 0; fi153 < 16; fi153++ {
                        var v154 int32 = 77
                        var v155 uint32 = 2
                        _ = fi153
                        _ = v154
                        _ = v155
                    }
                    v156 := GenS4{F0: 0}
                    _ = fi152
                    _ = v156
                }
            }
        default:
            v157 := GenS13{F0: v2}
            _ = v157
        }
    } else {
        v158 := GenS2{F0: 255, F1: 42, F2: -1}
        fmt.Println(v55, v121)
        _ = v158
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v121:", v121)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v11
    _ = v12
    _ = v121
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
}
