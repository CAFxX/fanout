package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint; F1 uint8; F2 bool }
type GenS7 struct { F0 int16; F1 int32 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint; F1 int16; F2 string }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int32; F1 int64; F2 int64 }
type GenS13 struct { F0 int64; F1 string; F2 uint16 }
type GenS14 struct { F0 uint; F1 int16; F2 uint16 }
type GenS15 struct { F0 int32; F1 int32; F2 uint32 }
type GenS16 struct { F0 uint64 }
type GenS17 struct { F0 int }
type GenS18 struct { F0 uint16 }
type GenS19 struct { F0 uint8; F1 int32; F2 int }
type GenS20 struct { F0 int64; F1 uint32; F2 int8 }
type GenS21 struct { F0 int8; F1 uint; F2 uint8 }
type GenS22 struct { F0 uint64 }
func genFn0() int {
    for fi5 := 0; fi5 < 13; fi5++ {
        switch fi5 {
        case 32768:
            fmt.Println(fi5)
        case 127:
            v6 := GenS1{F0: 2147483647, F1: 517}
            _ = v6
        default:
            if true {
                for fi7 := 0; fi7 < 20; fi7++ {
                    v8 := []int{fi7}
                    switch "hello" {
                    case "foo bar":
                        if len(v8) > 0 {
                            v8[0] = 0
                        }
                    }
                    _ = fi7
                    _ = v8
                }
                if false {
                    switch "foo bar" {
                    case "a":
                        var v9 uint16 = 42
                        _ = v9
                    case "foo bar":
                        var v10 float32 = 0.1
                        _ = v10
                    case "hello":
                        if false {
                            if (42 < 0) {
                                for fi11 := 0; fi11 < 6; fi11++ {
                                    var v12 int = (^(fi11))
                                    var v13 float64 = -2.25
                                    v14 := []int{30, (v12 | (fi11 << ((402) & 63))), fi5}
                                    v13 = 0.0
                                    _ = fi11
                                    _ = v12
                                    _ = v13
                                    _ = v14
                                }
                                var v15 uint64 = 42
                                if (true && true) {
                                    for fi16 := 0; fi16 < 17; fi16++ {
                                        if true {
                                            v17 := []int{(fi5 | (fi16 | fi5)), fi16, (fi5 << ((65535) & 63))}
                                            _ = len(v17)
                                            var v18 uint8 = 0
                                            _ = v17
                                            _ = v18
                                        } else {
                                            var v19 int = 127
                                            switch 896 {
                                            case 728:
                                                fmt.Println(fi5, v15)
                                            case 42:
                                                v20 := []int{fi5, v19, fi16}
                                                _ = v20
                                            }
                                            if (!false) {
                                                v21 := []int{v19, 32767, v19}
                                                v15 = (v15 - (v15 & v15))
                                                _ = v21
                                            } else {
                                                var v22 int32 = -1
                                                _ = v22
                                            }
                                            _ = v19
                                        }
                                        _ = fi16
                                    }
                                    v15 -= (^(v15))
                                    fmt.Println(v15)
                                } else {
                                    for fi23 := 0; fi23 < 13; fi23++ {
                                        v24 := []int{(fi5 - fi23), fi5}
                                        fmt.Println(v15, v15, fi23)
                                        v15 = 18446744073709551615
                                        _ = fi23
                                        _ = v24
                                    }
                                    switch ("foo bar" + "\n\t") {
                                    case "hello":
                                        fmt.Println(fi5, fi5)
                                    case "":
                                        fmt.Println(v15, fi5)
                                    case "foo bar":
                                        for fi25 := 0; fi25 < 16; fi25++ {
                                            switch 9223372036854775807 {
                                            case 508:
                                                v26 := GenS1{F0: 0, F1: 1}
                                                _ = v26
                                            default:
                                                v27 := []int{}
                                                _ = v27
                                            }
                                            v28 := []int{}
                                            if len(v28) >= 2 {
                                                v28 = v28[1:len(v28)]
                                            }
                                            v15 -= (^(v15))
                                            _ = fi25
                                            _ = v28
                                        }
                                    default:
                                        switch (fi5 & -2147483648) {
                                        case 1:
                                            var v29 bool = (!false)
                                            _ = v29
                                        case -1:
                                            v30 := GenS2{F0: 42}
                                            _ = v30
                                        case -987:
                                            var v31 int64 = 0
                                            _ = v31
                                        }
                                    }
                                }
                                _ = v15
                            } else {
                                v32 := GenS2{F0: 608}
                                var v33 uint8 = 128
                                switch ("a" + "\n\t") {
                                case "":
                                    v33 += (v33 >> ((1) & 63))
                                }
                                _ = v32
                                _ = v33
                            }
                            for fi34 := 0; fi34 < 2; fi34++ {
                                if true {
                                    if (true && true) {
                                        var v35 uint64 = 1
                                        v36 := GenS1{F0: -749, F1: 1}
                                        _ = v35
                                        _ = v36
                                    }
                                    var v37 uint32 = 65535
                                    var v38 float32 = (0.1 + ((3.14159 + ((1.5 / 3.14159) / ((float32(1e10) + (-2.25 - (0.1 - (float32(3.14159) - 0.0)))) + 0.0))) - -2.25))
                                    switch "hello" {
                                    case "\n\t":
                                        var v39 float32 = (((v38 / (float32(1.5) / (v38 - (float32(3.14159) / (3.14159 * v38))))) / v38) - v38)
                                        _ = v39
                                    case "":
                                        var v40 uint32 = v37
                                        _ = v40
                                    default:
                                        var v41 int8 = -110
                                        _ = v41
                                    }
                                    _ = v37
                                    _ = v38
                                } else {
                                    fmt.Println(fi34)
                                    var v42 int64 = -225
                                    _ = v42
                                }
                                if true {
                                    v43 := []int{(fi34 / ((fi5) | 1)), 128}
                                    var v44 int32 = 0
                                    if (2 < 127) {
                                        v44 = (-(v44))
                                    } else {
                                        v44 = v44
                                        fmt.Println(fi5, fi34, fi5)
                                        switch (fi5 - 32767) {
                                        case -603:
                                            switch "" {
                                            case "\n\t":
                                                if false {
                                                    switch (fi34 / (((fi34 | 127)) | 1)) {
                                                    case -353:
                                                        if len(v43) >= 2 {
                                                            v43 = v43[1:len(v43)]
                                                        }
                                                    case -2147483648:
                                                        switch 311 {
                                                        case -545:
                                                            v44 = 1
                                                        case -2147483648:
                                                            if len(v43) >= 2 {
                                                                v43 = v43[1:len(v43)]
                                                            }
                                                        }
                                                    default:
                                                        var v45 uint = 2
                                                        _ = v45
                                                    }
                                                    var v46 uint16 = 32767
                                                    v44 = (^(v44))
                                                    _ = v46
                                                } else {
                                                    if len(v43) > 0 {
                                                        v43[0] = fi34
                                                    }
                                                    switch 2147483647 {
                                                    case -9223372036854775808:
                                                        v44 += v44
                                                    case 32768:
                                                        if false {
                                                            _ = len(v43)
                                                        }
                                                    }
                                                    v47 := GenS2{F0: 2}
                                                    _ = v47
                                                }
                                            case "hello":
                                                var v48 uint64 = 42
                                                _ = v48
                                            case "a":
                                                switch "" {
                                                case "\n\t":
                                                    var v49 uint8 = 2
                                                    var v50 bool = false
                                                    v51 := GenS3{F0: (^(fi5)), F1: 277}
                                                    switch fi5 {
                                                    case 2147483647:
                                                        v52 := GenS4{F0: (fi34 | fi5)}
                                                        var v53 int8 = 1
                                                        _ = v52
                                                        _ = v53
                                                    case -438:
                                                        fmt.Println(v44, v49)
                                                    default:
                                                        switch ("foo bar" + "\n\t") {
                                                        case "":
                                                            switch (fi5 - (^(fi5))) {
                                                            case 9223372036854775807:
                                                                fmt.Println(v49, fi5, v49)
                                                            default:
                                                                v44 += v44
                                                            }
                                                        case "a":
                                                            v43 = append(v43, (-(fi5)))
                                                        }
                                                    }
                                                    switch fi5 {
                                                    case -2147483648:
                                                        v54 := GenS2{F0: 1}
                                                        _ = v54
                                                    case 42:
                                                        v49 = v49
                                                    case 4294967295:
                                                        var v55 bool = v50
                                                        _ = v55
                                                    default:
                                                        if ((-(v44)) > (v44 + 42)) {
                                                            v43 = append(v43, 32768)
                                                            switch (("" + ("" + "a")) + "foo bar") {
                                                            case "a":
                                                                var v56 int32 = 0
                                                                _ = v56
                                                            case "\n\t":
                                                                var v57 int32 = (v44 >> ((65535) & 63))
                                                                _ = v57
                                                            default:
                                                                if true {
                                                                    v50 = v50
                                                                }
                                                            }
                                                        } else {
                                                            var v58 int64 = 0
                                                            v59 := GenS3{F0: 128, F1: 2}
                                                            _ = v58
                                                            _ = v59
                                                        }
                                                        v49 -= v49
                                                    }
                                                    _ = v49
                                                    _ = v50
                                                    _ = v51
                                                }
                                            }
                                        case -9223372036854775808:
                                            v60 := GenS3{F0: (^(fi34)), F1: 245}
                                            _ = v60
                                        case 32767:
                                            fmt.Println(fi5, fi34)
                                        default:
                                            switch "foo bar" {
                                            case "a":
                                                v44 = v44
                                            default:
                                                fmt.Println(fi34, v44, fi5)
                                            }
                                        }
                                    }
                                    _ = v43
                                    _ = v44
                                } else {
                                    var v61 float32 = 1.5
                                    v62 := []int{(fi5 << ((168) & 63)), fi5}
                                    _ = v61
                                    _ = v62
                                }
                                _ = fi34
                            }
                            for fi63 := 0; fi63 < 13; fi63++ {
                                fmt.Println(fi5, fi63)
                                _ = fi63
                            }
                        } else {
                            var v64 uint16 = 32767
                            _ = v64
                        }
                    default:
                        var v65 int8 = 0
                        _ = v65
                    }
                    switch 0 {
                    case 364:
                        for fi66 := 0; fi66 < 2; fi66++ {
                            var v67 string = "a"
                            _ = fi66
                            _ = v67
                        }
                    case 32767:
                        var v68 uint8 = 5
                        _ = v68
                    case 4294967295:
                        var v69 bool = true
                        _ = v69
                    }
                    var v70 string = ("a" + "\n\t")
                    _ = v70
                } else {
                    v71 := GenS2{F0: 65535}
                    var v72 uint16 = 143
                    for fi73 := 0; fi73 < 17; fi73++ {
                        if (1000 > 231) {
                            v74 := GenS5{F0: 199}
                            _ = v74
                        } else {
                            v72 -= v72
                            if false {
                                var v75 int32 = -1
                                _ = v75
                            } else {
                                var v76 float32 = (0.0 + -2.25)
                                _ = v76
                            }
                            v72 = 461
                            v72 -= v72
                            v72 -= v72
                        }
                        _ = fi73
                    }
                    _ = v71
                    _ = v72
                }
                switch 32767 {
                case -9223372036854775808:
                    v77 := GenS1{F0: 42, F1: 255}
                    _ = v77
                }
            }
        }
        _ = fi5
    }
    v78 := []int{-1, 127}
    var v79 uint64 = 2
    _ = v78
    _ = v79
    return 117
}
func genFn1(p0 uint16) int16 {
    v80 := []int{42, 4294967295, -1, genFn0()}
    switch 4294967295 {
    case 0:
        for fi81 := 0; fi81 < 4; fi81++ {
            if true {
                for fi82 := 0; fi82 < 4; fi82++ {
                    fmt.Println(p0, fi81, p0)
                    var v83 float64 = (1e10 - 1e10)
                    _ = fi82
                    _ = v83
                }
                if (1e10 <= float64(3.14159)) {
                    if ((0 > 0) && (false && true)) {
                        switch "foo bar" {
                        case "a":
                            var v84 int = 1
                            _ = v84
                        case "\n\t":
                            var v85 uint = 255
                            _ = v85
                        }
                    }
                    for fi86 := 0; fi86 < 9; fi86++ {
                        var v87 uint64 = 18446744073709551615
                        v80 = append(v80, fi86)
                        if (p0 > (p0 + p0)) {
                            if ((fi86 | genFn0()) < genFn0()) {
                                v80 = append(v80, 9223372036854775807)
                                switch fi86 {
                                case -889:
                                    _ = len(v80)
                                case 32768:
                                    v87 -= v87
                                }
                                v88 := GenS2{F0: (p0 << ((255) & 63))}
                                _ = v88
                            }
                            if len(v80) > 0 {
                                v80[0] = fi86
                            }
                        }
                        _ = fi86
                        _ = v87
                    }
                    v80 = append(v80, 2147483647)
                } else {
                    for fi89 := 0; fi89 < 19; fi89++ {
                        switch "a" {
                        case "\n\t":
                            var v90 int = genFn0()
                            _ = v90
                        }
                        if len(v80) > 0 {
                            v80[0] = 32768
                        }
                        switch fi89 {
                        case 9223372036854775807:
                            p0 = (p0 << ((255) & 63))
                        case 761:
                            fmt.Println(fi81, fi89)
                        default:
                            v91 := GenS5{F0: p0}
                            _ = v91
                        }
                        _ = fi89
                    }
                }
            } else {
                switch 2147483647 {
                case 643:
                    v92 := GenS6{F0: 42, F1: 46, F2: (!(0 != 1))}
                    _ = v92
                case 127:
                    _ = len(v80)
                }
                _ = len(v80)
            }
            _ = fi81
        }
    case -756:
        var v93 int16 = 0
        _ = v93
    default:
        if (!false) {
            if (true && true) {
                v94 := GenS7{F0: 0, F1: 1}
                _ = v94
            }
            fmt.Println(p0)
            if ((p0 * p0) <= (p0 * p0)) {
                v95 := GenS8{F0: (!(!(1 <= 0)))}
                _ = len(v80)
                for fi96 := 0; fi96 < 20; fi96++ {
                    switch (fi96 / ((fi96) | 1)) {
                    case 9223372036854775807:
                        fmt.Println(p0, fi96)
                    default:
                        switch -2147483648 {
                        case 42:
                            if len(v80) > 0 {
                                v80[0] = fi96
                            }
                        case 0:
                            if true {
                                for fi97 := 0; fi97 < 6; fi97++ {
                                    if false {
                                        var v98 uint32 = uint32(p0)
                                        if len(v80) >= 2 {
                                            v80 = v80[1:len(v80)]
                                        }
                                        _ = v98
                                    }
                                    _ = fi97
                                }
                                fmt.Println(p0, p0)
                            } else {
                                switch (fi96 & (fi96 + -1)) {
                                case -366:
                                    v80 = append(v80, (fi96 >> ((726) & 63)))
                                case 127:
                                    fmt.Println(fi96, p0)
                                case 864:
                                    _ = len(v80)
                                default:
                                    p0 += (p0 * 0)
                                }
                            }
                        case 425:
                            var v99 float32 = 3.14159
                            _ = v99
                        }
                    }
                    p0 = (p0 & 2)
                    for fi100 := 0; fi100 < 13; fi100++ {
                        p0 = p0
                        _ = fi100
                    }
                    _ = fi96
                }
                _ = v95
            } else {
                fmt.Println(p0)
                for fi101 := 0; fi101 < 6; fi101++ {
                    if (1 == 42) {
                        if true {
                            if true {
                                v102 := GenS4{F0: (fi101 | fi101)}
                                switch genFn0() {
                                case 127:
                                    p0 -= 859
                                case 0:
                                    p0 = 32767
                                default:
                                    p0 += p0
                                }
                                var v103 uint8 = 127
                                _ = v102
                                _ = v103
                            } else {
                                p0 -= 32767
                                fmt.Println(p0, fi101)
                            }
                        } else {
                            switch ("a" + "") {
                            case "a":
                                if false {
                                    v104 := GenS9{F0: (p0 << ((1) & 63))}
                                    v105 := GenS10{F0: uint(p0), F1: 0, F2: ("" + (("" + (((("foo bar" + "a") + ((("\n\t" + "") + "") + (("hello" + ((("hello" + "a") + ("a" + "foo bar")) + "\n\t")) + "hello"))) + "a") + "")) + ("foo bar" + ("" + ("hello" + "foo bar")))))}
                                    _ = v104
                                    _ = v105
                                } else {
                                    var v106 string = ("hello" + (("a" + "hello") + "hello"))
                                    v106 = ("hello" + (v106 + v106))
                                    var v107 float64 = 1.5
                                    _ = v106
                                    _ = v107
                                }
                            case "\n\t":
                                fmt.Println(fi101)
                            case "foo bar":
                                p0 = 65535
                            }
                            if (3.14159 < (float32((float64(-2.25) - 0.0)) + 1.5)) {
                                var v108 uint16 = (p0 - p0)
                                _ = v108
                            } else {
                                var v109 uint8 = 1
                                _ = v109
                            }
                        }
                    } else {
                        switch (fi101 / ((fi101) | 1)) {
                        case 898:
                            switch 34 {
                            case 415:
                                if false {
                                    switch "\n\t" {
                                    case "\n\t":
                                        _ = len(v80)
                                    case "a":
                                        if false {
                                            v80 = append(v80, (fi101 * 1))
                                            var v110 bool = (42 == 2147483647)
                                            for fi111 := 0; fi111 < 15; fi111++ {
                                                v112 := GenS9{F0: 42}
                                                _ = fi111
                                                _ = v112
                                            }
                                            _ = v110
                                        }
                                    case "foo bar":
                                        p0 -= (p0 % (((p0 & 2)) | 1))
                                    default:
                                        if len(v80) > 0 {
                                            v80[0] = (fi101 - genFn0())
                                        }
                                    }
                                    if (1 >= 700) {
                                        p0 = (p0 % (((p0 << ((1) & 63))) | 1))
                                    } else {
                                        if len(v80) > 0 {
                                            v80[0] = (fi101 + (fi101 | (fi101 - (fi101 & (fi101 + fi101)))))
                                        }
                                    }
                                    p0 += p0
                                }
                            }
                        case 128:
                            switch "a" {
                            case "hello":
                                v113 := GenS11{F0: 941}
                                _ = v113
                            }
                        case 733:
                            p0 -= 32767
                        default:
                            if len(v80) > 0 {
                                v80[0] = genFn0()
                            }
                        }
                    }
                    if ((p0 / ((p0) | 1)) <= p0) {
                        for fi114 := 0; fi114 < 1; fi114++ {
                            var v115 int = (fi101 | (fi101 << ((471) & 63)))
                            v116 := GenS8{F0: ((false && (false || (1 >= 42))) && true)}
                            _ = fi114
                            _ = v115
                            _ = v116
                        }
                    } else {
                        for fi117 := 0; fi117 < 13; fi117++ {
                            p0 -= p0
                            var v118 uint = 677
                            _ = fi117
                            _ = v118
                        }
                        fmt.Println(fi101)
                        v119 := GenS12{F0: -2147483648, F1: 1, F2: -455}
                        _ = v119
                    }
                    _ = fi101
                }
            }
        } else {
            if false {
                switch -18 {
                case 2147483647:
                    fmt.Println(p0)
                case -2147483648:
                    if true {
                        var v120 int64 = 0
                        var v121 float64 = 0.1
                        for fi122 := 0; fi122 < 12; fi122++ {
                            v123 := GenS7{F0: int16(v120), F1: 42}
                            for fi124 := 0; fi124 < 11; fi124++ {
                                var v125 float64 = v121
                                _ = fi124
                                _ = v125
                            }
                            var v126 int = fi122
                            _ = fi122
                            _ = v123
                            _ = v126
                        }
                        _ = v120
                        _ = v121
                    } else {
                        for fi127 := 0; fi127 < 15; fi127++ {
                            switch "" {
                            case "":
                                var v128 uint64 = 18446744073709551615
                                _ = v128
                            case "a":
                                if len(v80) > 0 {
                                    v80[0] = genFn0()
                                }
                            case "hello":
                                switch 2147483647 {
                                case 0:
                                    for fi129 := 0; fi129 < 17; fi129++ {
                                        v130 := GenS6{F0: 1, F1: 0, F2: (fi129 == fi127)}
                                        _ = len(v80)
                                        var v131 string = "\n\t"
                                        _ = fi129
                                        _ = v130
                                        _ = v131
                                    }
                                case 42:
                                    fmt.Println(p0, p0)
                                case -806:
                                    if (float64(0.0) <= -2.25) {
                                        for fi132 := 0; fi132 < 15; fi132++ {
                                            fmt.Println(fi132, p0)
                                            if ((((p0 + 334) <= 32767) && true) || (379 > 406)) {
                                                _ = len(v80)
                                                var v133 uint64 = 1
                                                _ = v133
                                            }
                                            var v134 int16 = 671
                                            _ = fi132
                                            _ = v134
                                        }
                                        p0 = (^(p0))
                                    } else {
                                        var v135 int8 = -128
                                        _ = v135
                                    }
                                default:
                                    if (!(true && ((93 < fi127) && (448 >= -32768)))) {
                                        for fi136 := 0; fi136 < 1; fi136++ {
                                            if (!true) {
                                                switch (fi136 >> ((18446744073709551615) & 63)) {
                                                case 0:
                                                    var v137 int32 = 1
                                                    _ = v137
                                                default:
                                                    if (("foo bar" + "a") == ("hello" + "a")) {
                                                        var v138 uint16 = p0
                                                        fmt.Println(fi136, fi127)
                                                        _ = v138
                                                    }
                                                }
                                                if (-305 < -9223372036854775808) {
                                                    if (2 >= 42) {
                                                        var v139 int64 = 0
                                                        if len(v80) > 0 {
                                                            v80[0] = (fi136 << ((0) & 63))
                                                        }
                                                        switch genFn0() {
                                                        case 4294967295:
                                                            switch "\n\t" {
                                                            case "":
                                                                v140 := GenS1{F0: 446, F1: 65535}
                                                                _ = v140
                                                            case "a":
                                                                switch (("foo bar" + "foo bar") + "\n\t") {
                                                                case "foo bar":
                                                                    p0 += p0
                                                                case "hello":
                                                                    fmt.Println(fi136, v139)
                                                                case "":
                                                                    var v141 int64 = (^(v139))
                                                                    _ = v141
                                                                default:
                                                                    v142 := GenS8{F0: true}
                                                                    fmt.Println(fi127)
                                                                    var v143 uint16 = 65535
                                                                    _ = v142
                                                                    _ = v143
                                                                }
                                                                var v144 bool = false
                                                                _ = v144
                                                            case "foo bar":
                                                                if false {
                                                                    if (!false) {
                                                                        switch genFn0() {
                                                                        case 713:
                                                                            if (188 == 133) {
                                                                                var v145 uint = 372
                                                                                var v146 float32 = 0.1
                                                                                v145 += (v145 * uint(p0))
                                                                                _ = v145
                                                                                _ = v146
                                                                            } else {
                                                                                v139 += v139
                                                                                fmt.Println(fi136)
                                                                                var v147 uint16 = (p0 % (((p0 | p0)) | 1))
                                                                                fmt.Println(fi127, fi136, p0)
                                                                                _ = v147
                                                                            }
                                                                        case 0:
                                                                            v148 := GenS7{F0: 1000, F1: -105}
                                                                            _ = v148
                                                                        case 127:
                                                                            var v149 float64 = (float64(((float32(-2.25) + -2.25) * ((3.14159 + 0.0) / 1.5))) + ((-2.25 - -2.25) - 1.5))
                                                                            _ = v149
                                                                        default:
                                                                            var v150 uint16 = (^(p0))
                                                                            _ = v150
                                                                        }
                                                                        _ = len(v80)
                                                                    } else {
                                                                        var v151 uint64 = 1
                                                                        _ = v151
                                                                    }
                                                                    if (false || false) {
                                                                        v152 := GenS2{F0: p0}
                                                                        _ = v152
                                                                    } else {
                                                                        v139 -= (v139 | v139)
                                                                        var v153 uint = 122
                                                                        _ = v153
                                                                    }
                                                                    v154 := GenS9{F0: (^(p0))}
                                                                    _ = v154
                                                                }
                                                            }
                                                        default:
                                                            fmt.Println(p0)
                                                        }
                                                        _ = v139
                                                    }
                                                    p0 -= p0
                                                } else {
                                                    v80 = append(v80, (-(fi127)))
                                                }
                                                if len(v80) > 0 {
                                                    v80[0] = (fi136 * (fi136 ^ fi127))
                                                }
                                                switch (("a" + ("hello" + "foo bar")) + "\n\t") {
                                                case "hello":
                                                    var v155 int = (fi136 & -919)
                                                    _ = v155
                                                case "foo bar":
                                                    var v156 float32 = 3.14159
                                                    _ = v156
                                                default:
                                                    p0 += (p0 ^ 32767)
                                                }
                                                if true {
                                                    switch ("a" + "\n\t") {
                                                    case "a":
                                                        var v157 float32 = float32(-2.25)
                                                        _ = v157
                                                    case "hello":
                                                        if false {
                                                            var v158 uint8 = 203
                                                            _ = len(v80)
                                                            _ = v158
                                                        }
                                                    case "":
                                                        v159 := GenS13{F0: 0, F1: "", F2: p0}
                                                        var v160 float64 = -2.25
                                                        _ = v159
                                                        _ = v160
                                                    }
                                                    fmt.Println(fi127, p0)
                                                }
                                            } else {
                                                v80 = append(v80, (fi127 - 2147483647))
                                                v161 := GenS6{F0: 18446744073709551615, F1: 2, F2: true}
                                                _ = v161
                                            }
                                            if len(v80) >= 2 {
                                                v80 = v80[1:len(v80)]
                                            }
                                            _ = fi136
                                        }
                                        if len(v80) > 0 {
                                            v80[0] = fi127
                                        }
                                        fmt.Println(fi127)
                                    } else {
                                        fmt.Println(fi127)
                                        if (fi127 != (fi127 & (fi127 % ((32768) | 1)))) {
                                            fmt.Println(p0, p0)
                                        } else {
                                            var v162 uint64 = 4294967295
                                            if (v162 <= 0) {
                                                switch "hello" {
                                                case "\n\t":
                                                    p0 += p0
                                                case "foo bar":
                                                    var v163 int64 = 0
                                                    _ = v163
                                                case "a":
                                                    v162 += (^(v162))
                                                default:
                                                    for fi164 := 0; fi164 < 7; fi164++ {
                                                        v162 = v162
                                                        fmt.Println(fi164, fi164)
                                                        switch "foo bar" {
                                                        case "a":
                                                            var v165 int16 = -914
                                                            _ = v165
                                                        case "hello":
                                                            if (127 <= 127) {
                                                                if false {
                                                                    p0 = 65535
                                                                } else {
                                                                    if len(v80) >= 2 {
                                                                        v80 = v80[1:len(v80)]
                                                                    }
                                                                    if (!(false || false)) {
                                                                        p0 = p0
                                                                        var v166 uint32 = 42
                                                                        v167 := GenS14{F0: 2, F1: -298, F2: 407}
                                                                        _ = v166
                                                                        _ = v167
                                                                    } else {
                                                                        var v168 float32 = 3.14159
                                                                        fmt.Println(v162, p0)
                                                                        _ = v168
                                                                    }
                                                                }
                                                                var v169 int64 = -1
                                                                if (!true) {
                                                                    _ = len(v80)
                                                                }
                                                                if (-1 <= (fi127 ^ (fi127 & 9223372036854775807))) {
                                                                    var v170 uint32 = 2147483648
                                                                    _ = v170
                                                                } else {
                                                                    v171 := GenS6{F0: 0, F1: 2, F2: true}
                                                                    _ = v171
                                                                }
                                                                var v172 int64 = -990
                                                                _ = v169
                                                                _ = v172
                                                            } else {
                                                                p0 -= p0
                                                            }
                                                        case "\n\t":
                                                            v162 = (v162 << ((0) & 63))
                                                        default:
                                                            if true {
                                                                v80 = append(v80, (-(fi164)))
                                                                p0 = (p0 >> ((18446744073709551615) & 63))
                                                                var v173 int64 = 42
                                                                _ = v173
                                                            } else {
                                                                _ = len(v80)
                                                                fmt.Println(v162, fi164, v162)
                                                                p0 += p0
                                                            }
                                                        }
                                                        _ = fi164
                                                    }
                                                }
                                            }
                                            fmt.Println(p0, p0, fi127)
                                            _ = v162
                                        }
                                        var v174 bool = (true && (-122 != -128))
                                        _ = v174
                                    }
                                }
                            default:
                                p0 += p0
                            }
                            _ = fi127
                        }
                        p0 = 980
                    }
                case -1:
                    var v175 float32 = (float32(1.5) - 1e10)
                    _ = v175
                }
                v176 := GenS8{F0: (p0 <= 601)}
                for fi177 := 0; fi177 < 2; fi177++ {
                    switch ("hello" + "\n\t") {
                    case "":
                        var v178 uint = 18446744073709551615
                        _ = v178
                    default:
                        v179 := GenS1{F0: 42, F1: 18446744073709551615}
                        _ = v179
                    }
                    switch fi177 {
                    case 0:
                        switch fi177 {
                        case 1:
                            p0 = p0
                        }
                    case 128:
                        p0 -= p0
                    case -788:
                        var v180 uint16 = p0
                        _ = v180
                    }
                    var v181 int = fi177
                    _ = fi177
                    _ = v181
                }
                _ = v176
            }
            if len(v80) >= 2 {
                v80 = v80[1:len(v80)]
            }
            for fi182 := 0; fi182 < 6; fi182++ {
                for fi183 := 0; fi183 < 16; fi183++ {
                    var v184 float64 = ((float64(-2.25) + ((float64(float32((-2.25 - ((0.1 - 0.1) / 1.0)))) / 0.1) + (0.0 - 1.5))) * 1.5)
                    _ = fi183
                    _ = v184
                }
                for fi185 := 0; fi185 < 3; fi185++ {
                    v80 = append(v80, (fi182 >> ((255) & 63)))
                    if len(v80) > 0 {
                        v80[0] = genFn0()
                    }
                    _ = fi185
                }
                _ = fi182
            }
        }
    }
    for fi186 := 0; fi186 < 10; fi186++ {
        var v187 bool = false
        _ = fi186
        _ = v187
    }
    for fi188 := 0; fi188 < 16; fi188++ {
        if true {
            v189 := GenS8{F0: true}
            if true {
                if ((0.0 != 1.5) || true) {
                    for fi190 := 0; fi190 < 15; fi190++ {
                        v191 := GenS4{F0: 319}
                        fmt.Println(fi190, fi190, p0)
                        if ((^(fi190)) < fi190) {
                            v192 := GenS12{F0: 0, F1: 42, F2: -1}
                            fmt.Println(p0)
                            _ = v192
                        }
                        _ = fi190
                        _ = v191
                    }
                    for fi193 := 0; fi193 < 13; fi193++ {
                        if (-1 != 776) {
                            var v194 int = (^(fi193))
                            var v195 uint = 42
                            var v196 float64 = 0.1
                            _ = v194
                            _ = v195
                            _ = v196
                        }
                        var v197 int16 = 0
                        _ = fi193
                        _ = v197
                    }
                    fmt.Println(p0, fi188)
                }
                for fi198 := 0; fi198 < 11; fi198++ {
                    v199 := GenS15{F0: 7, F1: -1, F2: 2147483647}
                    var v200 int8 = 100
                    _ = fi198
                    _ = v199
                    _ = v200
                }
            }
            _ = v189
        }
        var v201 float32 = (((1.5 + 1.5) + 0.1) * -2.25)
        if false {
            v202 := GenS12{F0: -340, F1: 0, F2: -631}
            var v203 uint8 = 255
            _ = v202
            _ = v203
        }
        _ = fi188
        _ = v201
    }
    if len(v80) >= 2 {
        v80 = v80[1:len(v80)]
    }
    _ = p0
    _ = v80
    return -32768
}

func main() {
    var v1 int = 765
    var v2 uint16 = 255
    v3 := make(map[int]uint32)
    v4 := []string{"\n\t"}
    v1 = (v1 - 0)
    if true {
        v1 = v1
    }
    v204 := GenS16{F0: 54}
    v205 := GenS3{F0: 9223372036854775807, F1: 4294967295}
    var v206 uint = 2
    var v207 int = genFn0()
    var v208 int = v207
    var v209 float64 = 3.14159
    for fi210 := 0; fi210 < 1; fi210++ {
        switch fi210 {
        case 1:
            var v211 float32 = 1e10
            _ = v211
        case 4294967295:
            switch (fi210 * fi210) {
            case 128:
                var v212 float32 = -2.25
                _ = v212
            case -1:
                if (243 <= uint8(v206)) {
                    for fi213 := 0; fi213 < 14; fi213++ {
                        if false {
                            if true {
                                if (!true) {
                                    _ = len(v3)
                                    v209 += -2.25
                                } else {
                                    var v214 uint32 = 65535
                                    v215 := GenS3{F0: v208, F1: 18446744073709551615}
                                    _ = v214
                                    _ = v215
                                }
                            } else {
                                var v216 bool = false
                                fmt.Println(v216)
                                v209 -= (v209 - v209)
                                _ = len(v4)
                                v217 := GenS1{F0: 1, F1: (v206 * 347)}
                                var v218 uint8 = 255
                                _ = v216
                                _ = v217
                                _ = v218
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        _ = fi213
                    }
                }
            }
        default:
            fmt.Println(v209, v206, v1)
        }
        if true {
            if (false && false) {
                v219 := GenS8{F0: false}
                v220 := GenS17{F0: 565}
                _ = v219
                _ = v220
            }
            var v221 int64 = 0
            _ = v221
        } else {
            for fi222 := 0; fi222 < 14; fi222++ {
                v223 := GenS1{F0: -739, F1: v206}
                switch "foo bar" {
                case "\n\t":
                    var v224 bool = (uint32(v2) >= 1)
                    _ = v224
                case "a":
                    if (!false) {
                        var v225 uint32 = 984
                        v209 += (v209 - v209)
                        if false {
                            v225 = uint32(v2)
                            var v226 int16 = -32768
                            if true {
                                var v227 bool = true
                                var v228 string = "a"
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v227
                                _ = v228
                            }
                            _ = v226
                        } else {
                            fmt.Println(v206, v209)
                            var v229 int64 = -6
                            fmt.Println(fi210)
                            _ = v229
                        }
                        _ = v225
                    } else {
                        v230 := GenS12{F0: 2147483647, F1: 0, F2: 724}
                        fmt.Println(fi222, v2)
                        v231 := GenS7{F0: -1, F1: 42}
                        v209 += (v209 * 0.1)
                        if (-1 <= -128) {
                            var v232 float32 = (((1e10 - -2.25) / ((-2.25 - (float32(0.0) + float32(-2.25))) / (3.14159 - float32(v209)))) + 1.5)
                            var v233 uint = v206
                            _ = v232
                            _ = v233
                        }
                        var v234 bool = false
                        _ = v230
                        _ = v231
                        _ = v234
                    }
                case "":
                    var v235 int16 = -1
                    switch 2147483647 {
                    case -474:
                        var v236 uint64 = 18446744073709551615
                        if false {
                            if false {
                                if (!false) {
                                    if false {
                                        fmt.Println(v206, v207)
                                    }
                                    var v237 int16 = (v235 << ((0) & 63))
                                    v238 := GenS7{F0: (v235 | 1), F1: -60}
                                    _ = v237
                                    _ = v238
                                } else {
                                    v239 := GenS18{F0: v2}
                                    v208 = 127
                                    v240 := GenS6{F0: v206, F1: 127, F2: (75 >= 1)}
                                    _ = v239
                                    _ = v240
                                }
                                if (!((v2 ^ (^(v2))) < 65535)) {
                                    fmt.Println(v235)
                                } else {
                                    fmt.Println(v236)
                                    if (1 <= 1) {
                                        var v241 uint16 = 0
                                        _ = v241
                                    }
                                }
                                switch -2147483648 {
                                case 127:
                                    v242 := GenS12{F0: -2147483648, F1: -9223372036854775808, F2: -627}
                                    _ = v242
                                case -2147483648:
                                    v243 := GenS4{F0: v208}
                                    _ = v243
                                }
                                v235 = v235
                                delete(v3, v1)
                            } else {
                                var v244 uint64 = (v236 / ((2) | 1))
                                _ = v244
                            }
                            v245 := GenS19{F0: 39, F1: 1, F2: (v208 | fi222)}
                            v246 := GenS20{F0: 42, F1: 885, F2: -1}
                            _ = v245
                            _ = v246
                        } else {
                            switch genFn0() {
                            case 890:
                                v247 := GenS21{F0: 100, F1: 4294967295, F2: uint8(v206)}
                                _ = v247
                            case 2147483647:
                                v209 -= ((float64((0.0 + 0.1)) * v209) - float64(float32(v209)))
                            case 42:
                                var v248 float32 = 1.5
                                _ = v248
                            }
                            var v249 uint8 = 99
                            _ = v249
                        }
                        _ = v236
                    case -438:
                        switch "hello" {
                        case "hello":
                            if len(v4) > 0 {
                                v4[0] = ""
                            }
                            if true {
                                v208 = (fi222 / ((127) | 1))
                            }
                        }
                    default:
                        if false {
                            v250 := GenS9{F0: (v2 * (v2 % (((^(v2))) | 1)))}
                            v251 := GenS22{F0: 807}
                            _ = len(v4)
                            _ = v250
                            _ = v251
                        }
                    }
                    var v252 float32 = 3.14159
                    _ = v235
                    _ = v252
                default:
                    v206 += v206
                }
                if (!(1 < 42)) {
                    v207 -= 2147483647
                    var v253 int8 = 18
                    var v254 int = int(v2)
                    _ = v253
                    _ = v254
                } else {
                    v207 = v208
                }
                _ = fi222
                _ = v223
            }
            var v255 int = genFn0()
            _ = v255
        }
        for fi256 := 0; fi256 < 15; fi256++ {
            var v257 bool = ("hello" < "hello")
            _ = fi256
            _ = v257
        }
        _ = fi210
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v206:", v206)
    fmt.Println("v207:", v207)
    fmt.Println("v208:", v208)
    fmt.Println("v209:", v209)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v204
    _ = v205
    _ = v206
    _ = v207
    _ = v208
    _ = v209
    _ = v3
    _ = v4
}
