package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int32 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 int16; F1 int16; F2 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16; F1 int32; F2 int16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int32; F1 uint16 }
type GenS9 struct { F0 string; F1 int64; F2 int }
type GenS10 struct { F0 int8; F1 uint }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 bool; F1 uint32 }

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 18446744073709551615
    v3 := []string{}
    v4 := make(map[int]bool)
    v5 := []uint32{341}
    if true {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v2 = v2
    } else {
        fmt.Println(v2)
        var v6 float32 = ((3.14159 * 1.5) + 3.14159)
        _ = v6
    }
    switch (("a" + "foo bar") + "a") {
    case "":
        v2 = (v2 | v2)
    case "a":
        switch -580 {
        case 4294967295:
            fmt.Println(v2)
        default:
            v2 -= 704
        }
    case "hello":
        if (false || ((v2 ^ (v2 + (v2 << ((915) & 63)))) != v2)) {
            fmt.Println(v2)
        } else {
            if false {
                v2 = v2
                if (("foo bar" < (("hello" + ("hello" + "a")) + ("" + "a"))) && false) {
                    fmt.Println(v2)
                    switch 462 {
                    case 309:
                        for fi7 := 0; fi7 < 1; fi7++ {
                            var v8 float64 = ((-2.25 * 1e10) + (0.0 + 1.5))
                            fmt.Println(v8)
                            _ = fi7
                            _ = v8
                        }
                    }
                    switch ("a" + "\n\t") {
                    case "":
                        v2 = (^(v2))
                    }
                } else {
                    for fi9 := 0; fi9 < 16; fi9++ {
                        v10 := GenS1{F0: 25, F1: 0, F2: -758}
                        v2 = 765
                        _ = fi9
                        _ = v10
                    }
                    for fi11 := 0; fi11 < 10; fi11++ {
                        v2 += v2
                        if false {
                            v2 -= v2
                        }
                        _ = fi11
                    }
                }
            }
            for fi12 := 0; fi12 < 4; fi12++ {
                v13 := GenS1{F0: -1, F1: 73, F2: -965}
                var v14 bool = (!false)
                v2 = v2
                _ = fi12
                _ = v13
                _ = v14
            }
        }
    default:
        var v15 int16 = 1
        _ = v15
    }
    v2 += v2
    var v16 int16 = 844
    switch 0 {
    case -247:
        switch 4294967295 {
        case 0:
            var v17 string = "a"
            _ = v17
        case -9223372036854775808:
            switch (("\n\t" + (("foo bar" + "") + "foo bar")) + "") {
            case "":
                switch ("a" + ("\n\t" + (("hello" + "hello") + ""))) {
                case "hello":
                    for fi18 := 0; fi18 < 6; fi18++ {
                        v19 := GenS1{F0: 58, F1: 504, F2: -70}
                        var v20 int16 = (v16 & v16)
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                case "foo bar":
                    switch "foo bar" {
                    case "a":
                        for fi21 := 0; fi21 < 14; fi21++ {
                            fmt.Println(v16, v2)
                            v16 = v16
                            _ = fi21
                        }
                    default:
                        var v22 uint16 = 65535
                        _ = v22
                    }
                default:
                    var v23 uint32 = 0
                    _ = v23
                }
            case "\n\t":
                v2 = (v2 | (v2 - (v2 | (v2 | v2))))
            case "hello":
                for fi24 := 0; fi24 < 8; fi24++ {
                    var v25 float32 = 1.5
                    _ = fi24
                    _ = v25
                }
            }
        case 9223372036854775807:
            var v26 bool = (133 != 0)
            _ = v26
        }
    case 753:
        if false {
            switch 0 {
            case -9223372036854775808:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        } else {
            if true {
                var v27 bool = (false && (4294967295 == -9223372036854775808))
                var v28 int8 = -113
                var v29 uint = 0
                _ = v27
                _ = v28
                _ = v29
            }
            fmt.Println(v2, v16)
        }
    default:
        for fi30 := 0; fi30 < 13; fi30++ {
            var v31 uint32 = 42
            if (-659 <= -1) {
                v3 = append(v3, (("a" + ((("foo bar" + ((((("\n\t" + "\n\t") + "hello") + ("\n\t" + "a")) + ("a" + "hello")) + "foo bar")) + "a") + ("a" + "\n\t"))) + "hello"))
                switch fi30 {
                case 127:
                    switch "\n\t" {
                    case "hello":
                        v31 -= (v31 - 2)
                    case "\n\t":
                        v32 := GenS2{F0: fi30, F1: v16}
                        _ = v32
                    case "":
                        v2 = 910
                    default:
                        if false {
                            _ = len(v4)
                            if true {
                                v31 = 4294967295
                                for fi33 := 0; fi33 < 14; fi33++ {
                                    v16 = -1
                                    v31 += 2147483647
                                    var v34 string = "a"
                                    _ = len(v3)
                                    _ = fi33
                                    _ = v34
                                }
                                var v35 float32 = 3.14159
                                _ = v35
                            } else {
                                var v36 int16 = (v16 / (((v16 + (v16 << ((42) & 63)))) | 1))
                                _ = v36
                            }
                            var v37 float32 = (0.0 * 1.5)
                            _ = v37
                        }
                    }
                case 333:
                    delete(v4, 42)
                case 32768:
                    var v38 uint = 0
                    _ = v38
                default:
                    var v39 uint32 = 563
                    _ = v39
                }
                for fi40 := 0; fi40 < 6; fi40++ {
                    v41 := GenS1{F0: 50, F1: 18446744073709551615, F2: 935}
                    if len(v5) > 0 {
                        v5[0] = v31
                    }
                    _ = fi40
                    _ = v41
                }
            } else {
                fmt.Println(fi30, v2)
                for fi42 := 0; fi42 < 2; fi42++ {
                    var v43 bool = true
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v2 = (^(v2))
                    _ = fi42
                    _ = v43
                }
                delete(v4, -1)
            }
            _ = fi30
            _ = v31
        }
    }
    if ((128 != 2) || true) {
        for fi44 := 0; fi44 < 5; fi44++ {
            v16 += v16
            var v45 int8 = 42
            var v46 int64 = 1
            _ = fi44
            _ = v45
            _ = v46
        }
        var v47 int8 = -77
        v48 := GenS2{F0: 32767, F1: v16}
        _ = v47
        _ = v48
    } else {
        v49 := GenS1{F0: 0, F1: 0, F2: -2147483648}
        _ = v49
    }
    if ((0.0 * 1.5) == float32(-2.25)) {
        if (447 <= 1) {
            var v50 uint64 = v2
            for fi51 := 0; fi51 < 20; fi51++ {
                for fi52 := 0; fi52 < 1; fi52++ {
                    v2 += v50
                    v53 := GenS1{F0: 0, F1: 4294967295, F2: -461}
                    if true {
                        _ = len(v3)
                        switch (fi51 - (fi52 ^ 4294967295)) {
                        case -1:
                            v50 += (v50 >> ((2) & 63))
                        case 128:
                            if false {
                                v2 += 4294967295
                            } else {
                                v54 := GenS3{F0: -1, F1: (v16 + v16), F2: (v16 << ((65535) & 63))}
                                v2 += v2
                                _ = v54
                            }
                        }
                    }
                    _ = fi52
                    _ = v53
                }
                switch (fi51 | (fi51 >> ((1) & 63))) {
                case 32768:
                    fmt.Println(fi51, v16)
                case 32767:
                    var v55 int32 = -2147483648
                    _ = v55
                }
                var v56 bool = ("foo bar" == "\n\t")
                _ = fi51
                _ = v56
            }
            _ = v50
        }
        v2 = (v2 ^ (v2 >> ((0) & 63)))
    } else {
        fmt.Println(v16, v16)
        v57 := v1[128]
        if true {
            var v58 uint64 = v2
            for fi59 := 0; fi59 < 14; fi59++ {
                var v60 int8 = int8(fi59)
                if (95 == 93) {
                    for fi61 := 0; fi61 < 15; fi61++ {
                        var v62 int16 = (v16 | 1000)
                        v63 := GenS2{F0: -215, F1: (v16 + (-(v16)))}
                        v60 = 67
                        _ = fi61
                        _ = v62
                        _ = v63
                    }
                    v2 = 13
                } else {
                    v64 := v4[128]
                    _ = v64
                }
                var v65 int64 = -9223372036854775808
                _ = fi59
                _ = v60
                _ = v65
            }
            for fi66 := 0; fi66 < 20; fi66++ {
                fmt.Println(v2, v57, v57)
                _ = fi66
            }
            _ = v58
        } else {
            v16 = v16
            if false {
                v67 := GenS4{F0: 411}
                v68 := GenS5{F0: v16, F1: 1, F2: (v16 | v16)}
                v4[-9223372036854775808] = (642 != 0)
                _ = v67
                _ = v68
            }
            v5 = append(v5, 42)
        }
        _ = v57
    }
    for fi69 := 0; fi69 < 7; fi69++ {
        for fi70 := 0; fi70 < 9; fi70++ {
            switch "foo bar" {
            case "a":
                v71 := GenS2{F0: (fi69 - int(v2)), F1: v16}
                _ = v71
            case "foo bar":
                var v72 float32 = (3.14159 + (((3.14159 * (float32(((1e10 + 1.5) + (1.5 / 0.1))) + 3.14159)) / 1.5) * float32((((0.1 * 0.0) / 0.1) / 1.0))))
                _ = v72
            case "":
                v16 += v16
            }
            _ = fi70
        }
        var v73 int = fi69
        _ = fi69
        _ = v73
    }
    for fi74 := 0; fi74 < 8; fi74++ {
        if true {
            v16 = int16(fi74)
            var v75 string = "hello"
            _ = v75
        } else {
            v76 := GenS6{F0: -1}
            switch (fi74 & (fi74 / ((fi74) | 1))) {
            case 32767:
                v1[(fi74 & (fi74 << ((95) & 63)))] = "\n\t"
            case -14:
                v16 = (v16 | (v16 % ((1000) | 1)))
            case 521:
                v2 = v2
            }
            _ = len(v4)
            _ = v76
        }
        v2 += 42
        _ = fi74
    }
    var v77 float64 = 0.1
    fmt.Println(v16, v16)
    switch ("a" + (("\n\t" + ("\n\t" + "a")) + "")) {
    case "a":
        for fi78 := 0; fi78 < 14; fi78++ {
            for fi79 := 0; fi79 < 19; fi79++ {
                fmt.Println(v16)
                v5 = append(v5, 22)
                _ = len(v1)
                _ = fi79
            }
            for fi80 := 0; fi80 < 5; fi80++ {
                if false {
                    var v81 uint = 4294967295
                    v81 -= (v81 << ((v81) & 63))
                    if false {
                        var v82 int32 = -1
                        v81 -= v81
                        switch "a" {
                        case "hello":
                            v16 += v16
                        }
                        v77 = 1.5
                        var v83 uint16 = 42
                        _ = v82
                        _ = v83
                    } else {
                        v84 := v4[(fi78 << ((v81) & 63))]
                        switch (("\n\t" + "hello") + "hello") {
                        case "\n\t":
                            var v85 int32 = -942
                            _ = v85
                        case "foo bar":
                            if (2 <= (v2 >> (((v81 / ((18446744073709551615) | 1))) & 63))) {
                                fmt.Println(v77)
                                if true {
                                    fmt.Println(v81, v81)
                                } else {
                                    v86 := GenS7{F0: ("foo bar" + "foo bar")}
                                    _ = v86
                                }
                                v84 = (v84 && true)
                            }
                        case "a":
                            var v87 uint64 = v2
                            _ = v87
                        }
                        v1[1] = "hello"
                        _ = v84
                    }
                    var v88 uint8 = 238
                    fmt.Println(fi78, v16)
                    fmt.Println(v81)
                    v89 := GenS3{F0: v16, F1: v16, F2: -780}
                    _ = len(v4)
                    v90 := v4[-596]
                    _ = v81
                    _ = v88
                    _ = v89
                    _ = v90
                } else {
                    switch ((("\n\t" + "") + ("\n\t" + "\n\t")) + "hello") {
                    case "foo bar":
                        v91 := GenS3{F0: 32767, F1: (-(v16)), F2: (v16 / ((v16) | 1))}
                        _ = v91
                    case "a":
                        switch 4294967295 {
                        case 0:
                            _ = len(v1)
                        case 128:
                            if ("" == "") {
                                var v92 uint16 = 65535
                                fmt.Println(fi78, v92, fi80)
                                _ = v92
                            }
                        case 1:
                            var v93 int8 = 1
                            _ = v93
                        default:
                            var v94 int = fi78
                            _ = v94
                        }
                    default:
                        var v95 uint64 = (v2 | v2)
                        _ = v95
                    }
                    var v96 uint64 = v2
                    v2 -= v2
                    _ = v96
                }
                switch fi80 {
                case 128:
                    fmt.Println(fi80)
                default:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                _ = fi80
            }
            _ = fi78
        }
    }
    switch "foo bar" {
    case "a":
        var v97 int8 = -99
        _ = v97
    case "foo bar":
        switch 128 {
        case 890:
            v98 := GenS6{F0: -40}
            _ = v98
        case -2147483648:
            if false {
                for fi99 := 0; fi99 < 3; fi99++ {
                    if ((3.14159 == 3.14159) || (true || (true && false))) {
                        var v100 bool = false
                        _ = v100
                    }
                    _ = fi99
                }
                switch 127 {
                case 788:
                    for fi101 := 0; fi101 < 1; fi101++ {
                        v77 += 1.5
                        var v102 uint = 4294967295
                        v16 = -32768
                        _ = fi101
                        _ = v102
                    }
                case 32768:
                    v103 := GenS1{F0: 42, F1: 2, F2: 1}
                    _ = v103
                case 2147483647:
                    var v104 uint8 = 179
                    _ = v104
                default:
                    v16 = v16
                }
                var v105 uint64 = (v2 - (^(v2)))
                _ = v105
            } else {
                if ((!((1 != -619) && ((false || ((0.1 * -2.25) > 0.1)) && false))) && false) {
                    v106 := GenS8{F0: 1, F1: 255}
                    for fi107 := 0; fi107 < 11; fi107++ {
                        v108 := GenS9{F0: "hello", F1: 924, F2: (fi107 - fi107)}
                        var v109 float32 = (-2.25 + 0.1)
                        switch fi107 {
                        case 0:
                            if (false || (384 == 0)) {
                                v3 = append(v3, "\n\t")
                                var v110 uint32 = 261
                                _ = v110
                            } else {
                                fmt.Println(v77, fi107, v109)
                                if ((!((-(v16)) <= -32768)) || false) {
                                    if (!false) {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    } else {
                                        switch 1 {
                                        case -79:
                                            for fi111 := 0; fi111 < 1; fi111++ {
                                                _ = len(v3)
                                                var v112 uint64 = (v2 >> ((78) & 63))
                                                var v113 int8 = 42
                                                _ = fi111
                                                _ = v112
                                                _ = v113
                                            }
                                        default:
                                            var v114 string = "\n\t"
                                            _ = v114
                                        }
                                        v2 -= v2
                                    }
                                    v115 := GenS9{F0: "a", F1: -1, F2: fi107}
                                    _ = v115
                                } else {
                                    for fi116 := 0; fi116 < 12; fi116++ {
                                        var v117 uint32 = 65535
                                        _ = fi116
                                        _ = v117
                                    }
                                    v16 = int16(fi107)
                                }
                                v118 := GenS3{F0: v16, F1: (v16 << ((18446744073709551615) & 63)), F2: (v16 | (v16 | 42))}
                                _ = v118
                            }
                        case 2147483647:
                            fmt.Println(v16)
                        default:
                            fmt.Println(v109, v2, v16)
                        }
                        _ = fi107
                        _ = v108
                        _ = v109
                    }
                    if (672 != 612) {
                        for fi119 := 0; fi119 < 7; fi119++ {
                            var v120 int = 32768
                            _ = fi119
                            _ = v120
                        }
                        switch 32767 {
                        case -9223372036854775808:
                            v121 := GenS2{F0: -819, F1: v16}
                            _ = v121
                        default:
                            switch -960 {
                            case -244:
                                for fi122 := 0; fi122 < 6; fi122++ {
                                    v123 := GenS3{F0: v16, F1: v16, F2: (v16 / ((v16) | 1))}
                                    _ = fi122
                                    _ = v123
                                }
                            case 127:
                                if false {
                                    fmt.Println(v77, v16)
                                    var v124 int64 = -1
                                    _ = v124
                                } else {
                                    for fi125 := 0; fi125 < 6; fi125++ {
                                        var v126 uint8 = 127
                                        _ = fi125
                                        _ = v126
                                    }
                                }
                            default:
                                var v127 uint = 2
                                _ = v127
                            }
                        }
                        _ = len(v5)
                    }
                    _ = v106
                } else {
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                }
            }
        case 9223372036854775807:
            fmt.Println(v16, v77, v77)
        default:
            if false {
                if false {
                    if true {
                        v128 := GenS10{F0: 1, F1: uint(v2)}
                        _ = v128
                    }
                } else {
                    delete(v1, -364)
                    _ = len(v1)
                }
                fmt.Println(v2, v2, v77)
            } else {
                v129 := GenS7{F0: "hello"}
                v16 -= v16
                if false {
                    if false {
                        var v130 uint16 = 65535
                        var v131 bool = false
                        if (v77 >= ((v77 / -2.25) - float64(0.0))) {
                            if ((v16 - v16) != v16) {
                                for fi132 := 0; fi132 < 4; fi132++ {
                                    fmt.Println(v2)
                                    v133 := GenS8{F0: 42, F1: v130}
                                    _ = fi132
                                    _ = v133
                                }
                            } else {
                                var v134 string = "foo bar"
                                fmt.Println(v77, v131, v134)
                                for fi135 := 0; fi135 < 14; fi135++ {
                                    v2 -= v2
                                    _ = fi135
                                }
                                _ = v134
                            }
                            if (475 <= 1) {
                                switch (("foo bar" + ("" + ((("hello" + "\n\t") + "foo bar") + "foo bar"))) + "hello") {
                                case "hello":
                                    fmt.Println(v130, v2, v2)
                                case "a":
                                    fmt.Println(v77, v130, v131)
                                }
                                var v136 uint = 18446744073709551615
                                var v137 uint64 = 2
                                _ = v136
                                _ = v137
                            } else {
                                v130 = (v130 / (((^(v130))) | 1))
                                v2 = (v2 >> ((18446744073709551615) & 63))
                            }
                            for fi138 := 0; fi138 < 7; fi138++ {
                                for fi139 := 0; fi139 < 19; fi139++ {
                                    switch "foo bar" {
                                    case "hello":
                                        v131 = v131
                                    case "a":
                                        var v140 int = 2147483647
                                        _ = v140
                                    case "\n\t":
                                        v2 += v2
                                    default:
                                        v131 = (3.14159 != (1.5 + float32(((v77 / (v77 + v77)) * (1.5 / v77)))))
                                    }
                                    _ = fi139
                                }
                                _ = fi138
                            }
                        } else {
                            v16 = (v16 << ((853) & 63))
                            var v141 int16 = v16
                            _ = v141
                        }
                        _ = v130
                        _ = v131
                    }
                    switch ("foo bar" + "a") {
                    case "foo bar":
                        var v142 int16 = 1000
                        _ = v142
                    case "hello":
                        var v143 int = -9223372036854775808
                        _ = v143
                    case "":
                        if (1 >= 156) {
                            v77 = (float64(3.14159) / 3.14159)
                        }
                    default:
                        v16 = (^(v16))
                    }
                }
                _ = v129
            }
        }
    default:
        if (true || true) {
            for fi144 := 0; fi144 < 8; fi144++ {
                v145 := GenS7{F0: ("hello" + "")}
                _ = len(v4)
                switch "\n\t" {
                case "\n\t":
                    switch (("hello" + "a") + "hello") {
                    case "a":
                        if (!true) {
                            v146 := GenS11{F0: 1}
                            _ = v146
                        }
                    case "\n\t":
                        var v147 int64 = -9223372036854775808
                        _ = v147
                    }
                case "foo bar":
                    switch -9223372036854775808 {
                    case 4294967295:
                        fmt.Println(fi144)
                    }
                case "":
                    switch "" {
                    case "a":
                        var v148 uint16 = 42
                        _ = v148
                    case "hello":
                        switch "\n\t" {
                        case "a":
                            v5 = append(v5, 934)
                        default:
                            for fi149 := 0; fi149 < 18; fi149++ {
                                v2 += 861
                                _ = fi149
                            }
                        }
                    case "foo bar":
                        v150 := v1[fi144]
                        _ = v150
                    }
                }
                _ = fi144
                _ = v145
            }
            _ = len(v5)
        }
    }
    v151 := GenS12{F0: "a"}
    v16 = 1
    switch "foo bar" {
    case "foo bar":
        for fi152 := 0; fi152 < 16; fi152++ {
            v153 := GenS8{F0: 544, F1: 817}
            v154 := GenS13{F0: (1 < 65535), F1: 65535}
            fmt.Println(v2)
            _ = fi152
            _ = v153
            _ = v154
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v151
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v77
}
