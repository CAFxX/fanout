package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 int; F2 int64 }
type GenS3 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int64; F1 int32 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int32; F1 int8 }
type GenS9 struct { F0 uint32; F1 string }
type GenS10 struct { F0 string; F1 int32 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint32; F1 bool }
type GenS13 struct { F0 uint64 }

func main() {
    var v1 int32 = 2147483647
    var v2 uint32 = 1
    var v3 bool = (115 <= 2)
    v4 := make(map[string]int)
    var v5 uint64 = 848
    v6 := []string{"\n\t", "\n\t"}
    v7 := GenS1{F0: (v5 * (v5 % ((v5) | 1)))}
    for fi8 := 0; fi8 < 11; fi8++ {
        v2 = 2147483647
        _ = fi8
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        if (54 == 128) {
            if v3 {
                v4[""] = (fi9 - -2147483648)
                var v10 int16 = -178
                if false {
                    for fi11 := 0; fi11 < 14; fi11++ {
                        if true {
                            if (float32(3.14159) == (1e10 / -2.25)) {
                                switch fi11 {
                                case 32767:
                                    if (-9223372036854775808 > -1) {
                                        var v12 uint8 = 128
                                        if (!v3) {
                                            var v13 float64 = 0.0
                                            var v14 uint16 = 42
                                            var v15 int32 = (v1 % ((42) | 1))
                                            v5 = (^(v5))
                                            var v16 bool = v3
                                            var v17 uint = 543
                                            _ = v13
                                            _ = v14
                                            _ = v15
                                            _ = v16
                                            _ = v17
                                        }
                                        fmt.Println(v10)
                                        _ = v12
                                    } else {
                                        v1 -= (^(v1))
                                        delete(v4, ("" + ""))
                                    }
                                case 2147483647:
                                    var v18 int32 = 42
                                    _ = v18
                                default:
                                    v5 = v5
                                }
                                v4["hello"] = (fi9 % ((fi9) | 1))
                                v5 = (v5 / ((v5) | 1))
                            }
                            var v19 float32 = ((1.5 - 0.0) + 1.5)
                            _ = v19
                        }
                        if len(v6) > 0 {
                            v6[0] = "foo bar"
                        }
                        _ = fi11
                    }
                    var v20 uint16 = 83
                    switch fi9 {
                    case 2147483647:
                        if v3 {
                            var v21 uint8 = 2
                            var v22 bool = v3
                            _ = v21
                            _ = v22
                        } else {
                            var v23 int32 = v1
                            for fi24 := 0; fi24 < 11; fi24++ {
                                var v25 uint32 = (v2 << ((42) & 63))
                                v26 := GenS2{F0: (!(65535 <= 65535)), F1: fi24, F2: -9223372036854775808}
                                v27 := GenS3{F0: (v1 / ((-2147483648) | 1)), F1: 128, F2: v2}
                                _ = fi24
                                _ = v25
                                _ = v26
                                _ = v27
                            }
                            switch ("a" + ("foo bar" + "hello")) {
                            case "foo bar":
                                var v28 string = "foo bar"
                                _ = v28
                            case "\n\t":
                                fmt.Println(v5)
                            case "hello":
                                if (-2147483648 == (-(v23))) {
                                    var v29 uint = 988
                                    fmt.Println(v5, v29, v10)
                                    _ = v29
                                } else {
                                    var v30 string = ""
                                    _ = v30
                                }
                            }
                            _ = v23
                        }
                    }
                    _ = v20
                } else {
                    var v31 uint64 = 0
                    fmt.Println(v5)
                    if (v3 || (true && false)) {
                        if v3 {
                            var v32 string = ("\n\t" + "")
                            _ = v32
                        } else {
                            if ((!((v3 && v3) && v3)) || v3) {
                                if true {
                                    switch fi9 {
                                    case 0:
                                        fmt.Println(v5, fi9)
                                    case -556:
                                        var v33 string = "a"
                                        _ = v33
                                    case 42:
                                        fmt.Println(fi9, v10, fi9)
                                    default:
                                        fmt.Println(fi9)
                                    }
                                    var v34 uint = 2
                                    _ = v34
                                }
                                if v3 {
                                    v35 := GenS4{F0: "a"}
                                    v36 := GenS5{F0: v3}
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    _ = v35
                                    _ = v36
                                }
                                if (230 <= 120) {
                                    v1 -= (v1 * v1)
                                }
                            }
                            v10 -= v10
                        }
                        if (123 != 65535) {
                            var v37 int8 = -5
                            v38 := GenS3{F0: (-(v1)), F1: 0, F2: v2}
                            _ = v37
                            _ = v38
                        } else {
                            var v39 float64 = 1.5
                            _ = v39
                        }
                    } else {
                        var v40 int64 = -1
                        var v41 float32 = 0.0
                        var v42 uint = 42
                        _ = v40
                        _ = v41
                        _ = v42
                    }
                    _ = v31
                }
                _ = v10
            }
            fmt.Println(fi9, v2)
        } else {
            var v43 float64 = 1e10
            v2 = v2
            var v44 bool = (!((fi9 & fi9) != fi9))
            _ = v43
            _ = v44
        }
        if true {
            var v45 uint32 = 2147483648
            v46 := GenS2{F0: (false || (true && v3)), F1: fi9, F2: 771}
            var v47 uint32 = 0
            _ = v45
            _ = v46
            _ = v47
        } else {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            if v3 {
                v3 = false
            }
        }
        _ = fi9
    }
    for fi48 := 0; fi48 < 13; fi48++ {
        v49 := GenS6{F0: 42, F1: (v1 * v1)}
        _ = fi48
        _ = v49
    }
    delete(v4, "a")
    v5 = v5
    if ((1e10 > 1e10) && (127 > 8)) {
        v4["a"] = 101
    }
    for fi50 := 0; fi50 < 16; fi50++ {
        if v3 {
            v1 = v1
            var v51 float64 = 1.5
            var v52 float64 = v51
            _ = v51
            _ = v52
        } else {
            v2 -= 494
            _ = len(v6)
        }
        v5 -= v5
        _ = fi50
    }
    var v53 int16 = -1
    switch (("foo bar" + "foo bar") + "\n\t") {
    case "":
        var v54 int = 128
        _ = v54
    default:
        v5 = v5
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if (v3 && (v3 && v3)) {
        if (v3 || ("foo bar" != ("a" + (("foo bar" + ("hello" + "hello")) + (("\n\t" + "\n\t") + "\n\t"))))) {
            v3 = v3
            v55 := GenS3{F0: (v1 ^ v1), F1: 42, F2: (v2 & 42)}
            if v3 {
                v56 := GenS3{F0: (v1 + (v1 >> ((0) & 63))), F1: 2, F2: 2147483648}
                for fi57 := 0; fi57 < 2; fi57++ {
                    switch fi57 {
                    case 128:
                        var v58 bool = ((3.14159 > float64(3.14159)) && v3)
                        _ = v58
                    case 0:
                        if len(v6) > 0 {
                            v6[0] = "\n\t"
                        }
                    }
                    fmt.Println(v2, v5)
                    fmt.Println(v2)
                    _ = fi57
                }
                _ = v56
            } else {
                var v59 int = 127
                _ = v59
            }
            _ = v55
        }
        v60 := GenS7{F0: 32767}
        switch "hello" {
        case "":
            if v3 {
                for fi61 := 0; fi61 < 15; fi61++ {
                    var v62 bool = (431 == 0)
                    for fi63 := 0; fi63 < 18; fi63++ {
                        var v64 int64 = 931
                        var v65 bool = true
                        _ = fi63
                        _ = v64
                        _ = v65
                    }
                    _ = fi61
                    _ = v62
                }
            }
        case "a":
            v1 -= v1
        case "\n\t":
            if (2 != 42) {
                v1 = (v1 | int32(v5))
                for fi66 := 0; fi66 < 13; fi66++ {
                    switch (fi66 >> ((255) & 63)) {
                    case 1:
                        v53 = 42
                    }
                    _ = fi66
                }
                v3 = (!v3)
            }
        }
        _ = v60
    } else {
        for fi67 := 0; fi67 < 4; fi67++ {
            var v68 bool = (true && ((1.5 - 0.1) > 1.5))
            v69 := GenS2{F0: (v2 == v2), F1: 127, F2: 9223372036854775807}
            for fi70 := 0; fi70 < 12; fi70++ {
                switch ("\n\t" + "") {
                case "\n\t":
                    v71 := GenS1{F0: (^(v5))}
                    _ = v71
                case "a":
                    var v72 int8 = -59
                    _ = v72
                case "":
                    _ = len(v6)
                default:
                    v73 := GenS8{F0: (v1 % (((v1 << ((112) & 63))) | 1)), F1: -116}
                    _ = v73
                }
                if v68 {
                    if v68 {
                        v2 = v2
                        var v74 int16 = v53
                        _ = v74
                    } else {
                        delete(v4, ((("hello" + "\n\t") + ("hello" + ("hello" + ("foo bar" + "hello")))) + "foo bar"))
                        _ = len(v6)
                        var v75 float64 = ((-2.25 - 0.0) - 0.1)
                        _ = v75
                    }
                    v3 = (true && v68)
                } else {
                    var v76 int64 = 9223372036854775807
                    _ = v76
                }
                _ = fi70
            }
            _ = fi67
            _ = v68
            _ = v69
        }
        if (84 != (^(v5))) {
            switch "" {
            case "a":
                var v77 uint64 = 0
                _ = v77
            case "":
                var v78 uint16 = 1
                _ = v78
            case "\n\t":
                for fi79 := 0; fi79 < 12; fi79++ {
                    var v80 string = ""
                    switch fi79 {
                    case 0:
                        switch (fi79 >> ((1) & 63)) {
                        case 32768:
                            var v81 uint32 = (v2 | v2)
                            _ = v81
                        case 2147483647:
                            var v82 int32 = (v1 >> ((1) & 63))
                            _ = v82
                        case -1:
                            _ = len(v4)
                        }
                    case -1:
                        for fi83 := 0; fi83 < 5; fi83++ {
                            var v84 uint16 = 32767
                            if ((true && (v80 < ("foo bar" + "foo bar"))) && v3) {
                                v5 += 2
                                var v85 uint = 429
                                var v86 uint32 = v2
                                _ = v85
                                _ = v86
                            }
                            fmt.Println(v3, fi79)
                            _ = fi83
                            _ = v84
                        }
                    }
                    v2 = v2
                    _ = fi79
                    _ = v80
                }
            default:
                switch (((("hello" + "foo bar") + "hello") + (("hello" + "") + "")) + ("" + "foo bar")) {
                case "":
                    if (42 >= 184) {
                        switch "hello" {
                        case "":
                            var v87 string = ((("\n\t" + "hello") + "a") + "")
                            _ = v87
                        case "a":
                            var v88 int = -40
                            _ = v88
                        default:
                            if v3 {
                                for fi89 := 0; fi89 < 6; fi89++ {
                                    for fi90 := 0; fi90 < 14; fi90++ {
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        if v3 {
                                            fmt.Println(v53)
                                            var v91 uint64 = (^(v5))
                                            var v92 float64 = 3.14159
                                            var v93 uint64 = (v91 - (v91 << ((2) & 63)))
                                            _ = len(v6)
                                            var v94 int8 = -39
                                            _ = v91
                                            _ = v92
                                            _ = v93
                                            _ = v94
                                        }
                                        v5 -= (v5 >> ((4294967295) & 63))
                                        _ = fi90
                                    }
                                    if false {
                                        var v95 uint32 = (v2 >> ((18446744073709551615) & 63))
                                        v96 := GenS9{F0: v2, F1: "a"}
                                        _ = v95
                                        _ = v96
                                    }
                                    _ = fi89
                                }
                                v3 = v3
                            } else {
                                var v97 bool = (!(763 >= 42))
                                _ = v97
                            }
                        }
                    }
                case "a":
                    v4["foo bar"] = 42
                case "foo bar":
                    fmt.Println(v1, v1, v2)
                default:
                    v3 = v3
                }
            }
        } else {
            for fi98 := 0; fi98 < 11; fi98++ {
                switch (-(fi98)) {
                case 127:
                    _ = len(v4)
                case 32767:
                    switch "a" {
                    case "\n\t":
                        for fi99 := 0; fi99 < 16; fi99++ {
                            var v100 string = ""
                            if false {
                                v1 -= (v1 + (v1 % (((v1 * v1)) | 1)))
                                switch fi99 {
                                case 746:
                                    switch int(v53) {
                                    case 4294967295:
                                        var v101 int64 = -1
                                        _ = v101
                                    }
                                default:
                                    if (v3 || (255 >= uint8(v2))) {
                                        v102 := v4[("a" + "\n\t")]
                                        _ = v102
                                    }
                                }
                                var v103 uint16 = 255
                                _ = v103
                            }
                            _ = fi99
                            _ = v100
                        }
                    case "":
                        for fi104 := 0; fi104 < 2; fi104++ {
                            v2 += uint32(v5)
                            switch "foo bar" {
                            case "\n\t":
                                switch (fi104 * (-(fi98))) {
                                case 127:
                                    var v105 float32 = 0.1
                                    _ = v105
                                case -2147483648:
                                    switch (fi104 ^ fi104) {
                                    case 2147483647:
                                        var v106 int64 = int64(v1)
                                        _ = v106
                                    }
                                default:
                                    if (v2 < (v2 >> ((65535) & 63))) {
                                        if false {
                                            var v107 int64 = 1
                                            v108 := v4["\n\t"]
                                            v108 -= fi98
                                            _ = v107
                                            _ = v108
                                        }
                                        var v109 string = "foo bar"
                                        _ = v109
                                    }
                                }
                            case "a":
                                fmt.Println(v3)
                                switch "hello" {
                                case "":
                                    if true {
                                        var v110 bool = (!v3)
                                        _ = v110
                                    } else {
                                        v2 = 2147483648
                                    }
                                }
                            }
                            if false {
                                if false {
                                    v1 = v1
                                } else {
                                    var v111 uint8 = 127
                                    switch (-(fi98)) {
                                    case -881:
                                        if len(v6) > 0 {
                                            v6[0] = "a"
                                        }
                                    default:
                                        var v112 int64 = int64(fi98)
                                        fmt.Println(v1)
                                        fmt.Println(v2)
                                        _ = v112
                                    }
                                    v111 = 128
                                    v113 := GenS2{F0: true, F1: (fi104 & fi104), F2: -420}
                                    _ = len(v6)
                                    _ = v111
                                    _ = v113
                                }
                            } else {
                                if true {
                                    if (true || (v3 || (0.0 == float32((1e10 * float64(-2.25)))))) {
                                        if false {
                                            var v114 int64 = 42
                                            if len(v6) > 0 {
                                                v6[0] = "\n\t"
                                            }
                                            fmt.Println(v114)
                                            _ = v114
                                        }
                                        _ = len(v4)
                                        switch ("hello" + "\n\t") {
                                        case "\n\t":
                                            fmt.Println(fi98, v1)
                                        default:
                                            fmt.Println(fi104, fi104, v53)
                                        }
                                        v115 := GenS8{F0: v1, F1: 42}
                                        _ = v115
                                    } else {
                                        var v116 uint16 = 2
                                        _ = v116
                                    }
                                    var v117 uint16 = 65535
                                    var v118 uint32 = (v2 << ((255) & 63))
                                    v5 -= 671
                                    var v119 uint64 = v5
                                    _ = v117
                                    _ = v118
                                    _ = v119
                                } else {
                                    var v120 int = (fi104 & fi98)
                                    var v121 uint = 706
                                    _ = v120
                                    _ = v121
                                }
                                v3 = (255 > 212)
                                if v3 {
                                    v122 := GenS6{F0: 9223372036854775807, F1: int32(v5)}
                                    _ = v122
                                }
                            }
                            if (!(55 != 190)) {
                                v123 := GenS9{F0: (v2 * (v2 / ((956) | 1))), F1: ""}
                                if false {
                                    v2 += 65535
                                    v4[""] = 32768
                                    switch "hello" {
                                    case "":
                                        v53 = v53
                                    case "\n\t":
                                        var v124 float32 = float32(1e10)
                                        _ = v124
                                    default:
                                        if (v53 < v53) {
                                            fmt.Println(v5)
                                            _ = len(v4)
                                            fmt.Println(v1, v2, fi104)
                                            if ((v2 > v2) && true) {
                                                switch -553 {
                                                case 9223372036854775807:
                                                    v53 += v53
                                                case 128:
                                                    switch "" {
                                                    case "hello":
                                                        if (!(-275 != (-(fi98)))) {
                                                            v6 = append(v6, ("a" + "\n\t"))
                                                        }
                                                    case "":
                                                        v1 -= (v1 >> ((1) & 63))
                                                        v125 := GenS10{F0: "", F1: 58}
                                                        _ = v125
                                                    }
                                                case -600:
                                                    var v126 uint32 = (v2 ^ (^(v2)))
                                                    _ = v126
                                                default:
                                                    switch (fi98 - fi104) {
                                                    case -1:
                                                        v3 = true
                                                    case 562:
                                                        if v3 {
                                                            v3 = (255 > 32767)
                                                        }
                                                    case 768:
                                                        fmt.Println(v2, fi104, fi98)
                                                    default:
                                                        v127 := GenS11{F0: (v5 % (((^(v5))) | 1))}
                                                        _ = v127
                                                    }
                                                }
                                                var v128 int16 = 967
                                                _ = v128
                                            } else {
                                                v2 -= 4294967295
                                                if ((^(v1)) <= v1) {
                                                    if v3 {
                                                        if v3 {
                                                            switch "foo bar" {
                                                            case "a":
                                                                var v129 bool = ((v1 - 1) > (v1 | v1))
                                                                var v130 string = ("a" + ("" + "hello"))
                                                                _ = v129
                                                                _ = v130
                                                            case "":
                                                                if len(v6) > 0 {
                                                                    v6[0] = "hello"
                                                                }
                                                            case "foo bar":
                                                                _ = len(v4)
                                                            default:
                                                                fmt.Println(v53)
                                                            }
                                                        }
                                                    }
                                                    if v3 {
                                                        _ = len(v6)
                                                    }
                                                } else {
                                                    var v131 int8 = 115
                                                    _ = v131
                                                }
                                                v132 := GenS12{F0: (v2 | (v2 + v2)), F1: v3}
                                                _ = v132
                                            }
                                        } else {
                                            var v133 uint32 = 4294967295
                                            _ = v133
                                        }
                                        var v134 float32 = 1.5
                                        switch fi104 {
                                        case 42:
                                            fmt.Println(v53, fi104, fi98)
                                        case 128:
                                            v135 := GenS7{F0: 32767}
                                            _ = v135
                                        case 127:
                                            fmt.Println(fi98, fi98)
                                            var v136 uint64 = (^(v5))
                                            if v3 {
                                                switch (("hello" + ("" + "foo bar")) + "\n\t") {
                                                case "foo bar":
                                                    fmt.Println(v53)
                                                case "hello":
                                                    v137 := GenS13{F0: (v5 & v136)}
                                                    _ = v137
                                                case "a":
                                                    fmt.Println(v53, v5)
                                                    v138 := GenS2{F0: (v3 && v3), F1: 42, F2: -792}
                                                    var v139 string = "\n\t"
                                                    fmt.Println(v5, v134)
                                                    _ = v138
                                                    _ = v139
                                                default:
                                                    if (v1 < v1) {
                                                        var v140 uint16 = 914
                                                        var v141 bool = (!((fi98 - 2147483647) > (^(fi98))))
                                                        fmt.Println(fi104, v2, v141)
                                                        _ = v140
                                                        _ = v141
                                                    }
                                                }
                                                if v3 {
                                                    if v3 {
                                                        v134 -= (float32(0.1) * 1e10)
                                                    }
                                                    if (true || v3) {
                                                        fmt.Println(v136, fi104, v53)
                                                    }
                                                    delete(v4, "\n\t")
                                                    delete(v4, "")
                                                }
                                            }
                                            _ = v136
                                        }
                                        if (3.14159 <= 0.0) {
                                            switch "\n\t" {
                                            case "":
                                                switch fi98 {
                                                case 665:
                                                    v142 := GenS3{F0: (v1 - (v1 ^ v1)), F1: 128, F2: 340}
                                                    _ = v142
                                                }
                                            default:
                                                var v143 string = "a"
                                                _ = v143
                                            }
                                            v2 += (^(v2))
                                            var v144 bool = (v3 && (13 < -6))
                                            _ = v144
                                        } else {
                                            v5 += 789
                                        }
                                        if v3 {
                                            if len(v6) > 0 {
                                                v6[0] = "hello"
                                            }
                                            v1 = v1
                                        } else {
                                            if true {
                                                var v145 int64 = 42
                                                fmt.Println(v5, v1)
                                                _ = v145
                                            } else {
                                                v2 += (v2 | (v2 % (((v2 - (v2 | 840))) | 1)))
                                            }
                                            var v146 float32 = (1e10 * 0.0)
                                            v147 := GenS10{F0: "hello", F1: (v1 / (((v1 ^ 0)) | 1))}
                                            _ = v146
                                            _ = v147
                                        }
                                        _ = v134
                                    }
                                }
                                _ = v123
                            }
                            _ = fi104
                        }
                    case "a":
                        v1 = 0
                    }
                }
                if (-128 >= 106) {
                    v53 += 1000
                    fmt.Println(v3, v2)
                    if (0.0 < 0.0) {
                        v3 = false
                        var v148 uint8 = 203
                        _ = v148
                    }
                }
                _ = fi98
            }
            v149 := GenS12{F0: 2147483647, F1: v3}
            _ = v149
        }
        v53 += int16(v2)
    }
    fmt.Println(v3)
    var v150 uint8 = 155
    fmt.Println("v1:", v1)
    fmt.Println("v150:", v150)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v150
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
    _ = v7
}
