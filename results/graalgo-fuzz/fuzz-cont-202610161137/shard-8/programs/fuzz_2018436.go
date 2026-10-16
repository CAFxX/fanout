package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 uint8; F2 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 int8; F1 string; F2 int16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 string; F1 uint }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int32; F1 string }
type GenS10 struct { F0 string }
type GenS11 struct { F0 bool; F1 int8; F2 int32 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 uint; F1 int64; F2 int64 }

func main() {
    v1 := make(map[int]string)
    var v2 float32 = (0.1 * 0.1)
    var v3 float32 = (v2 * 0.0)
    var v4 int32 = -1
    v5 := make(map[int]uint32)
    var v6 uint8 = 1
    if true {
        v3 = 3.14159
    } else {
        for fi7 := 0; fi7 < 6; fi7++ {
            v8 := GenS1{F0: -32768}
            _ = len(v1)
            _ = fi7
            _ = v8
        }
        fmt.Println(v2)
    }
    v4 = (v4 + 2147483647)
    if false {
        v4 -= v4
        var v9 bool = true
        _ = v9
    }
    switch (("" + "\n\t") + "hello") {
    case "":
        for fi10 := 0; fi10 < 12; fi10++ {
            if (true && (!true)) {
                if true {
                    v11 := GenS2{F0: 9223372036854775807, F1: 128, F2: 1}
                    _ = v11
                } else {
                    v4 -= v4
                    for fi12 := 0; fi12 < 15; fi12++ {
                        if (v2 > v3) {
                            switch "\n\t" {
                            case "foo bar":
                                var v13 uint = 476
                                _ = v13
                            case "a":
                                var v14 int16 = 0
                                _ = v14
                            case "\n\t":
                                if (("a" + "a") != "hello") {
                                    var v15 uint8 = 0
                                    _ = v15
                                }
                            default:
                                v16 := GenS2{F0: fi10, F1: v6, F2: 42}
                                _ = v16
                            }
                        } else {
                            if (808 == 2) {
                                v17 := []int{(fi10 + (-(fi12))), 32767}
                                if true {
                                    if true {
                                        v18 := GenS3{F0: true}
                                        v19 := GenS4{F0: (fi10 - fi10), F1: -1}
                                        switch fi12 {
                                        case 42:
                                            var v20 uint8 = v6
                                            _ = v20
                                        default:
                                            switch "a" {
                                            case "a":
                                                var v21 int = (-(fi10))
                                                _ = v21
                                            default:
                                                v22 := GenS1{F0: 109}
                                                _ = v22
                                            }
                                        }
                                        _ = v18
                                        _ = v19
                                    }
                                    _ = len(v17)
                                }
                                v2 += (v3 + v2)
                                _ = v17
                            }
                        }
                        fmt.Println(v4, v6)
                        fmt.Println(v6)
                        v23 := []int{32767, (^(fi12)), fi12, (fi12 + (fi12 << ((255) & 63)))}
                        var v24 uint = 42
                        switch (fi10 << ((740) & 63)) {
                        case 678:
                            v2 -= 0.1
                        case 32768:
                            if (("foo bar" + "foo bar") == ("\n\t" + (("" + "a") + "hello"))) {
                                var v25 float32 = (((0.0 / (((float32(-2.25) * ((v3 - 0.0) - v2)) / 1.0) + 1.5)) + 3.14159) + float32(float64(v3)))
                                if (false && (-361 >= v4)) {
                                    if false {
                                        _ = len(v23)
                                    } else {
                                        v26 := GenS2{F0: fi12, F1: (v6 >> (((v24 - (^(v24)))) & 63)), F2: 885}
                                        var v27 uint = v24
                                        v28 := GenS4{F0: 354, F1: 1000}
                                        _ = v26
                                        _ = v27
                                        _ = v28
                                    }
                                }
                                v29 := GenS5{F0: 96, F1: (("hello" + "hello") + ""), F2: 840}
                                _ = v25
                                _ = v29
                            }
                        case -482:
                            v30 := GenS3{F0: ((!true) && false)}
                            _ = v30
                        default:
                            fmt.Println(v6, fi12)
                        }
                        var v31 uint16 = 255
                        var v32 int32 = (v4 ^ v4)
                        _ = fi12
                        _ = v23
                        _ = v24
                        _ = v31
                        _ = v32
                    }
                    v2 = v3
                }
                fmt.Println(v6, v4)
            } else {
                if ((v6 + v6) == v6) {
                    v33 := []int{fi10, 1}
                    _ = v33
                }
                for fi34 := 0; fi34 < 1; fi34++ {
                    fmt.Println(v6, fi34)
                    v1[fi34] = ("foo bar" + ("\n\t" + "\n\t"))
                    var v35 uint8 = (v6 | 2)
                    _ = fi34
                    _ = v35
                }
            }
            _ = fi10
        }
    case "a":
        switch "hello" {
        case "":
            fmt.Println(v2)
        case "a":
            var v36 int = 2147483647
            _ = v36
        case "\n\t":
            v4 += 42
        }
    default:
        for fi37 := 0; fi37 < 11; fi37++ {
            v38 := GenS4{F0: int(v4), F1: 42}
            var v39 uint8 = (v6 << ((65535) & 63))
            var v40 int64 = 187
            _ = fi37
            _ = v38
            _ = v39
            _ = v40
        }
    }
    var v41 int64 = 1
    switch -272 {
    case 32768:
        if false {
            switch (((((("foo bar" + (("foo bar" + "a") + "")) + "\n\t") + ("" + ("a" + ("foo bar" + "hello")))) + "") + ((((("" + "\n\t") + "foo bar") + "") + ("a" + (("" + "") + "a"))) + "")) + ((((("" + "hello") + ("" + "a")) + ("foo bar" + "a")) + "a") + "foo bar")) {
            case "":
                v4 += v4
            }
            var v42 int64 = v41
            switch "hello" {
            case "foo bar":
                fmt.Println(v6)
            default:
                v43 := GenS6{F0: v41}
                _ = v43
            }
            _ = v42
        }
    case 175:
        v44 := []int{-2147483648, 4294967295, 4294967295, 9223372036854775807}
        _ = v44
    }
    for fi45 := 0; fi45 < 1; fi45++ {
        if true {
            if (255 <= 838) {
                v41 -= 1
                v6 = 1
                for fi46 := 0; fi46 < 9; fi46++ {
                    var v47 uint64 = 1
                    _ = fi46
                    _ = v47
                }
            }
            var v48 int8 = 78
            if (!true) {
                var v49 uint64 = 830
                for fi50 := 0; fi50 < 6; fi50++ {
                    if (32767 >= 255) {
                        v51 := v1[0]
                        v41 += (v41 >> ((42) & 63))
                        _ = v51
                    } else {
                        if ((!(v4 == (-(v4)))) && true) {
                            v52 := GenS4{F0: fi45, F1: 303}
                            v53 := []int{fi50}
                            _ = v52
                            _ = v53
                        } else {
                            fmt.Println(v49, v49, v2)
                            v54 := []int{fi45, fi50, 42}
                            _ = v54
                        }
                        switch 0 {
                        case 1:
                            switch "" {
                            case "\n\t":
                                v6 += v6
                            case "a":
                                if false {
                                    var v55 float32 = 0.1
                                    v41 -= (v41 % (((^(v41))) | 1))
                                    if ("" >= "a") {
                                        fmt.Println(v4, fi45)
                                    } else {
                                        if true {
                                            fmt.Println(v3, v2)
                                        } else {
                                            if true {
                                                var v56 int64 = (v41 - v41)
                                                v57 := []int{fi45, fi50}
                                                if len(v57) > 0 {
                                                    v57[0] = (fi50 << ((4294967295) & 63))
                                                }
                                                _ = v56
                                                _ = v57
                                            }
                                            fmt.Println(v6, fi50, v41)
                                        }
                                        v58 := GenS5{F0: (v48 - v48), F1: ("a" + ("a" + ("a" + "foo bar"))), F2: 1000}
                                        _ = v58
                                    }
                                    _ = v55
                                }
                                fmt.Println(fi50, v49)
                                v59 := []int{fi45, (fi50 & -919), fi50}
                                _ = v59
                            case "":
                                v60 := v1[fi50]
                                _ = v60
                            default:
                                if true {
                                    var v61 int64 = -1
                                    _ = v61
                                }
                            }
                        case 32768:
                            _ = len(v5)
                        case -9223372036854775808:
                            if ((v49 & 2) <= 4294967295) {
                                v48 = (v48 | (v48 - v48))
                            } else {
                                var v62 float64 = 1.5
                                v63 := GenS7{F0: ("hello" + "\n\t"), F1: 4294967295}
                                v41 -= (-(v41))
                                var v64 float64 = v62
                                if (uint64(v6) == 903) {
                                    delete(v5, (fi50 * (fi45 ^ fi45)))
                                    if false {
                                        v65 := GenS3{F0: (v64 < v64)}
                                        var v66 uint16 = 562
                                        _ = v65
                                        _ = v66
                                    }
                                    var v67 uint8 = v6
                                    _ = v67
                                } else {
                                    if true {
                                        var v68 uint8 = (^(v6))
                                        var v69 float32 = 1.5
                                        _ = v68
                                        _ = v69
                                    } else {
                                        if false {
                                            switch (fi50 + fi45) {
                                            case 128:
                                                var v70 uint = uint(v49)
                                                _ = v70
                                            case 0:
                                                fmt.Println(v6)
                                            }
                                            switch fi50 {
                                            case -9223372036854775808:
                                                v71 := GenS8{F0: 0}
                                                _ = v71
                                            case -556:
                                                if false {
                                                    var v72 int = (^(fi45))
                                                    v73 := []int{(fi45 * 1), fi50}
                                                    v74 := GenS4{F0: fi45, F1: 1}
                                                    v62 = ((v64 + v64) / v64)
                                                    switch (^(fi50)) {
                                                    case 472:
                                                        if len(v73) > 0 {
                                                            v73[0] = (fi45 | 952)
                                                        }
                                                    case -1:
                                                        switch "a" {
                                                        case "hello":
                                                            v41 -= 1
                                                        default:
                                                            var v75 int16 = 130
                                                            _ = v75
                                                        }
                                                    default:
                                                        if (0 >= v41) {
                                                            var v76 int16 = 428
                                                            switch (fi50 - v72) {
                                                            case 32767:
                                                                _ = len(v1)
                                                            }
                                                            v41 -= 0
                                                            _ = v76
                                                        }
                                                    }
                                                    _ = v72
                                                    _ = v73
                                                    _ = v74
                                                } else {
                                                    _ = len(v1)
                                                    v77 := GenS8{F0: 2147483648}
                                                    v6 = (v6 + v6)
                                                    _ = v77
                                                }
                                            default:
                                                switch fi45 {
                                                case 4294967295:
                                                    if (!(true && (32767 != 32767))) {
                                                        v78 := GenS9{F0: int32(fi50), F1: "\n\t"}
                                                        fmt.Println(v6, v49)
                                                        _ = v78
                                                    } else {
                                                        switch fi45 {
                                                        case -1:
                                                            var v79 uint8 = uint8(v49)
                                                            switch (fi50 + 4294967295) {
                                                            case 42:
                                                                var v80 int16 = -1
                                                                _ = v80
                                                            }
                                                            v2 += v2
                                                            _ = v79
                                                        case -667:
                                                            if ((!false) || true) {
                                                                fmt.Println(v3, v49)
                                                                var v81 int64 = (v41 & v41)
                                                                _ = v81
                                                            }
                                                        case 128:
                                                            fmt.Println(v3, v64, fi45)
                                                            v6 -= (v6 + v6)
                                                            if true {
                                                                fmt.Println(v49)
                                                                v1[(fi45 >> ((4294967295) & 63))] = "hello"
                                                                var v82 uint32 = uint32(v49)
                                                                _ = v82
                                                            } else {
                                                                if (true && (false || (!true))) {
                                                                    v83 := []int{(-(fi50))}
                                                                    switch (("" + "\n\t") + "hello") {
                                                                    case "\n\t":
                                                                        var v84 uint8 = (^(v6))
                                                                        _ = v84
                                                                    case "foo bar":
                                                                        switch fi50 {
                                                                        case 737:
                                                                            fmt.Println(fi50, fi45)
                                                                        default:
                                                                            fmt.Println(v64, v41)
                                                                        }
                                                                    case "":
                                                                        v6 = 255
                                                                    default:
                                                                        v4 = v4
                                                                    }
                                                                    if ((v41 | v41) > int64(v6)) {
                                                                        if true {
                                                                            if ((fi50 & (^(fi45))) == fi45) {
                                                                                v85 := GenS4{F0: fi50, F1: -71}
                                                                                _ = v85
                                                                            } else {
                                                                                var v86 bool = true
                                                                                var v87 int16 = 1
                                                                                if true {
                                                                                    switch -411 {
                                                                                    case 4294967295:
                                                                                        v6 += (v6 | (v6 * (v6 + 0)))
                                                                                    }
                                                                                    v2 = float32(v62)
                                                                                    var v88 uint16 = 32767
                                                                                    switch ("a" + ((("a" + (("a" + "") + "hello")) + "foo bar") + "a")) {
                                                                                    case "":
                                                                                        var v89 int64 = 42
                                                                                        _ = v89
                                                                                    case "foo bar":
                                                                                        var v90 int64 = (v41 & v41)
                                                                                        _ = v90
                                                                                    case "a":
                                                                                        fmt.Println(v49, v48, v49)
                                                                                        v83 = append(v83, fi45)
                                                                                    default:
                                                                                        switch "\n\t" {
                                                                                        case "hello":
                                                                                            switch (^(fi45)) {
                                                                                            case 2147483647:
                                                                                                var v91 uint32 = 1
                                                                                                _ = v91
                                                                                            }
                                                                                        default:
                                                                                            v5[fi50] = 65535
                                                                                        }
                                                                                        v92 := GenS2{F0: 1, F1: v6, F2: v49}
                                                                                        v93 := v1[2147483647]
                                                                                        _ = v92
                                                                                        _ = v93
                                                                                    }
                                                                                    _ = v88
                                                                                } else {
                                                                                    v83 = append(v83, (fi45 << ((0) & 63)))
                                                                                    var v94 int64 = (v41 + v41)
                                                                                    _ = v94
                                                                                }
                                                                                _ = v86
                                                                                _ = v87
                                                                            }
                                                                            if (!true) {
                                                                                v48 = v48
                                                                                var v95 int16 = 141
                                                                                _ = v95
                                                                            }
                                                                        } else {
                                                                            if len(v83) > 0 {
                                                                                v83[0] = (fi45 - (fi45 * (fi50 / (((fi50 & (fi45 + fi45))) | 1))))
                                                                            }
                                                                        }
                                                                    } else {
                                                                        var v96 float32 = ((v2 * ((0.0 + -2.25) + (float32(v64) - (1e10 - ((v3 - v3) + 0.1))))) - v3)
                                                                        v62 = v62
                                                                        v2 = (float32(v64) * float32(((0.1 + v62) / v62)))
                                                                        _ = v96
                                                                    }
                                                                    _ = v83
                                                                }
                                                                if ((^(v49)) != (v49 ^ 1)) {
                                                                    var v97 uint16 = 27
                                                                    v48 -= v48
                                                                    v98 := GenS10{F0: ("a" + "a")}
                                                                    _ = v97
                                                                    _ = v98
                                                                } else {
                                                                    var v99 float64 = (float64(v2) / v64)
                                                                    _ = v99
                                                                }
                                                                var v100 float32 = v2
                                                                _ = v100
                                                            }
                                                        default:
                                                            v101 := []int{fi45, 127, (fi45 * fi45), int(v4)}
                                                            _ = v101
                                                        }
                                                        var v102 int8 = v48
                                                        _ = v102
                                                    }
                                                }
                                            }
                                            v103 := GenS5{F0: (v48 | (v48 >> ((695) & 63))), F1: "a", F2: 839}
                                            _ = v103
                                        }
                                    }
                                    var v104 int16 = 0
                                    v62 = (v64 - 1.5)
                                    _ = v104
                                }
                                _ = v62
                                _ = v63
                                _ = v64
                            }
                        default:
                            var v105 uint32 = 2147483648
                            _ = v105
                        }
                    }
                    _ = fi50
                }
                _ = v49
            } else {
                switch "hello" {
                case "a":
                    v3 -= v3
                case "hello":
                    v4 = (v4 & (v4 << ((4294967295) & 63)))
                case "\n\t":
                    for fi106 := 0; fi106 < 6; fi106++ {
                        if true {
                            switch fi45 {
                            case 327:
                                var v107 uint64 = 51
                                _ = v107
                            default:
                                var v108 int = (fi106 - (fi106 % (((fi45 / (((fi106 / ((fi106) | 1))) | 1))) | 1)))
                                _ = v108
                            }
                            switch fi45 {
                            case -885:
                                fmt.Println(v6, v2, v48)
                                v41 = (v41 >> ((4294967295) & 63))
                                v109 := []int{(fi106 - fi106)}
                                _ = v109
                            default:
                                v110 := []int{(fi45 * (fi106 % (((fi106 | int(v6))) | 1)))}
                                _ = v110
                            }
                        } else {
                            var v111 int8 = -1
                            _ = v111
                        }
                        v112 := []int{(fi45 | (fi106 % (((fi106 ^ (fi106 + 0))) | 1))), fi106}
                        var v113 int = -838
                        _ = fi106
                        _ = v112
                        _ = v113
                    }
                default:
                    for fi114 := 0; fi114 < 17; fi114++ {
                        var v115 int = (fi114 & fi114)
                        _ = fi114
                        _ = v115
                    }
                }
            }
            _ = v48
        }
        var v116 uint = 2
        _ = fi45
        _ = v116
    }
    if true {
        v117 := GenS11{F0: (1 >= 2), F1: 42, F2: (v4 & -2147483648)}
        _ = v117
    } else {
        v4 = v4
        var v118 uint = 0
        switch "" {
        case "foo bar":
            v2 = (3.14159 / (float32(float64((v3 / v2))) + v3))
        case "a":
            switch 4294967295 {
            case 9223372036854775807:
                var v119 bool = false
                _ = v119
            }
        case "\n\t":
            fmt.Println(v4)
        default:
            for fi120 := 0; fi120 < 1; fi120++ {
                if (-1 >= 42) {
                    var v121 uint64 = 42
                    var v122 uint16 = 848
                    _ = v121
                    _ = v122
                } else {
                    if ((v6 - v6) != (v6 >> (((v118 >> ((v118) & 63))) & 63))) {
                        var v123 int32 = v4
                        var v124 int16 = -697
                        for fi125 := 0; fi125 < 15; fi125++ {
                            v3 = -2.25
                            v126 := []int{128, (fi125 | (fi125 & (fi125 * (fi120 | (fi120 | -9223372036854775808))))), (fi125 - 514), fi125}
                            v127 := GenS10{F0: "a"}
                            if (-452 != v4) {
                                v128 := GenS8{F0: 1}
                                v129 := GenS8{F0: 4294967295}
                                _ = v128
                                _ = v129
                            } else {
                                switch ("\n\t" + ("foo bar" + "hello")) {
                                case "\n\t":
                                    var v130 int64 = v41
                                    _ = v130
                                case "":
                                    v6 += v6
                                }
                            }
                            _ = fi125
                            _ = v126
                            _ = v127
                        }
                        _ = v123
                        _ = v124
                    } else {
                        fmt.Println(v41, v2)
                        var v131 uint32 = 1
                        v132 := GenS9{F0: (v4 << (((v118 ^ v118)) & 63)), F1: ("a" + "")}
                        _ = v131
                        _ = v132
                    }
                    switch (fi120 ^ (fi120 % ((fi120) | 1))) {
                    case -2147483648:
                        v133 := GenS8{F0: uint32(v6)}
                        _ = v133
                    case 2147483647:
                        v134 := GenS12{F0: 2}
                        _ = v134
                    }
                    switch 967 {
                    case 128:
                        switch fi120 {
                        case 0:
                            if (v41 >= 1) {
                                if false {
                                    v2 = v3
                                    if (2 > uint32(v6)) {
                                        v41 -= (v41 + v41)
                                        var v135 uint64 = 1
                                        for fi136 := 0; fi136 < 1; fi136++ {
                                            var v137 uint32 = 0
                                            _ = fi136
                                            _ = v137
                                        }
                                        _ = v135
                                    }
                                }
                                if true {
                                    for fi138 := 0; fi138 < 4; fi138++ {
                                        if false {
                                            v139 := GenS13{F0: (v118 - 42)}
                                            _ = v139
                                        }
                                        _ = fi138
                                    }
                                }
                            }
                        case 127:
                            if (("" + "") != ("\n\t" + "")) {
                                for fi140 := 0; fi140 < 10; fi140++ {
                                    if true {
                                        v141 := GenS6{F0: (v41 << (((^(v118))) & 63))}
                                        _ = v141
                                    } else {
                                        fmt.Println(v41, v2)
                                        fmt.Println(fi140, v4)
                                        fmt.Println(v4, v2, v41)
                                    }
                                    fmt.Println(v2, v6)
                                    _ = fi140
                                }
                                for fi142 := 0; fi142 < 3; fi142++ {
                                    fmt.Println(v4)
                                    var v143 uint8 = v6
                                    _ = fi142
                                    _ = v143
                                }
                            } else {
                                if (!false) {
                                    v144 := GenS13{F0: v118}
                                    fmt.Println(v118, v3, v118)
                                    var v145 uint = v118
                                    _ = v144
                                    _ = v145
                                } else {
                                    switch fi120 {
                                    case -2147483648:
                                        v146 := v1[fi120]
                                        _ = v146
                                    case 727:
                                        for fi147 := 0; fi147 < 10; fi147++ {
                                            var v148 uint8 = v6
                                            if ("\n\t" >= (("\n\t" + "\n\t") + "a")) {
                                                v148 -= (v148 - (^(v148)))
                                                if (87 == -710) {
                                                    var v149 uint = 65535
                                                    v150 := []int{}
                                                    _ = v149
                                                    _ = v150
                                                } else {
                                                    v118 = v118
                                                }
                                            }
                                            if (!(856 != 4294967295)) {
                                                if true {
                                                    switch fi147 {
                                                    case -2147483648:
                                                        v1[fi120] = ((("" + ("foo bar" + "hello")) + "hello") + "foo bar")
                                                    }
                                                    fmt.Println(fi120, v4, v41)
                                                    v151 := []int{(fi120 ^ fi147), fi120, (fi120 << ((762) & 63))}
                                                    _ = v151
                                                } else {
                                                    var v152 bool = true
                                                    v6 += 127
                                                    v152 = (!((fi120 + -363) == fi120))
                                                    _ = v152
                                                }
                                            } else {
                                                if true {
                                                    v153 := GenS12{F0: 4294967295}
                                                    v154 := []int{(fi147 + (fi147 | fi147))}
                                                    var v155 string = "\n\t"
                                                    v156 := GenS14{F0: uint(v148), F1: 635, F2: 0}
                                                    _ = v153
                                                    _ = v154
                                                    _ = v155
                                                    _ = v156
                                                }
                                            }
                                            _ = fi147
                                            _ = v148
                                        }
                                    }
                                }
                            }
                        default:
                            fmt.Println(v2)
                        }
                    }
                }
                _ = fi120
            }
        }
        _ = v118
    }
    if (2 <= 214) {
        switch ("foo bar" + "\n\t") {
        case "\n\t":
            var v157 uint32 = 892
            _ = v157
        case "foo bar":
            v158 := []int{-423, 128, 127}
            _ = v158
        case "hello":
            v6 = 21
        default:
            var v159 float64 = (float64(v3) + (3.14159 + float64(v2)))
            _ = v159
        }
        for fi160 := 0; fi160 < 14; fi160++ {
            var v161 float32 = 3.14159
            _ = fi160
            _ = v161
        }
    }
    switch 127 {
    case 1:
        var v162 uint = 0
        _ = v162
    case 402:
        v41 += 1
    default:
        var v163 bool = true
        _ = v163
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
