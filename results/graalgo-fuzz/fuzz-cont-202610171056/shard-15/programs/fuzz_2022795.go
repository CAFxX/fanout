package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint; F1 int32; F2 uint }
type GenS3 struct { F0 bool; F1 uint8; F2 int }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS6 struct { F0 string; F1 uint16; F2 int8 }
type GenS7 struct { F0 uint8; F1 int16; F2 int }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS10 struct { F0 int64; F1 uint8 }
type GenS11 struct { F0 uint32; F1 uint16; F2 int }
type GenS12 struct { F0 int64; F1 uint8 }
type GenS13 struct { F0 bool; F1 uint64 }
type GenS14 struct { F0 int8; F1 uint8 }
type GenS15 struct { F0 int; F1 uint8 }
type GenS16 struct { F0 int8; F1 int8; F2 uint16 }

func main() {
    var v1 float64 = -2.25
    var v2 uint64 = 4294967295
    var v3 uint16 = 0
    var v4 int32 = -411
    var v5 uint32 = 0
    if true {
        if (!(2147483647 >= (v4 | v4))) {
            v6 := []int{}
            for fi7 := 0; fi7 < 6; fi7++ {
                var v8 uint32 = 2147483647
                v9 := GenS1{F0: fi7, F1: false}
                fmt.Println(v8, fi7, fi7)
                _ = fi7
                _ = v8
                _ = v9
            }
            _ = v6
        } else {
            if (!(!false)) {
                for fi10 := 0; fi10 < 18; fi10++ {
                    v4 = (^(v4))
                    var v11 int8 = 0
                    _ = fi10
                    _ = v11
                }
                var v12 uint64 = (v2 / ((42) | 1))
                _ = v12
            }
        }
        for fi13 := 0; fi13 < 3; fi13++ {
            v4 -= 2147483647
            _ = fi13
        }
    }
    v14 := GenS1{F0: -9223372036854775808, F1: true}
    v15 := GenS2{F0: 1, F1: (v4 + (v4 / ((v4) | 1))), F2: 79}
    if (!(0 <= 823)) {
        var v16 uint16 = (^(v3))
        var v17 uint32 = 42
        _ = v16
        _ = v17
    }
    switch ("a" + ("\n\t" + ("hello" + (("foo bar" + "") + ("a" + "foo bar"))))) {
    case "hello":
        if true {
            v18 := GenS3{F0: true, F1: 2, F2: -602}
            fmt.Println(v1, v3, v4)
            _ = v18
        }
    default:
        if true {
            fmt.Println(v3, v1, v1)
            var v19 int32 = int32(v2)
            if false {
                if false {
                    v20 := GenS4{F0: 479}
                    v4 = v4
                    v19 -= v19
                    _ = v20
                } else {
                    for fi21 := 0; fi21 < 7; fi21++ {
                        v2 = v2
                        for fi22 := 0; fi22 < 13; fi22++ {
                            v1 = v1
                            v23 := GenS5{F0: (v5 | (v5 * (v5 - (v5 / (((v5 * (v5 << ((65535) & 63)))) | 1))))), F1: 768, F2: (v3 - v3)}
                            _ = fi22
                            _ = v23
                        }
                        _ = fi21
                    }
                }
                var v24 uint16 = v3
                v2 += (v2 * 4294967295)
                _ = v24
            }
            _ = v19
        } else {
            v25 := GenS2{F0: 4294967295, F1: -1, F2: 774}
            _ = v25
        }
    }
    if (60 < -74) {
        if (false || ((float64(-2.25) - (v1 + v1)) > 1.5)) {
            var v26 uint8 = 1
            _ = v26
        }
    } else {
        v27 := []int{}
        if len(v27) > 0 {
            v27[0] = 32767
        }
        _ = v27
    }
    for fi28 := 0; fi28 < 5; fi28++ {
        var v29 string = "foo bar"
        v30 := GenS6{F0: "foo bar", F1: v3, F2: -116}
        _ = fi28
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 2; fi31++ {
        var v32 uint32 = 0
        switch "\n\t" {
        case "":
            if true {
                var v33 uint = uint(v32)
                if (!false) {
                    var v34 int64 = -300
                    if (v4 != (v4 << (((v33 ^ v33)) & 63))) {
                        v35 := GenS3{F0: (v2 != (^(v2))), F1: 228, F2: (-(fi31))}
                        switch fi31 {
                        case 9223372036854775807:
                            v36 := GenS2{F0: (v33 % ((4294967295) | 1)), F1: (v4 + (^(v4))), F2: 65535}
                            _ = v36
                        case 0:
                            v37 := []int{fi31, 215}
                            _ = v37
                        }
                        for fi38 := 0; fi38 < 10; fi38++ {
                            v39 := []int{(fi38 + (fi38 % (((fi31 / ((589) | 1))) | 1))), (fi31 - fi38), (fi31 ^ fi31), 32767}
                            fmt.Println(v1, fi38)
                            var v40 int32 = v4
                            _ = fi38
                            _ = v39
                            _ = v40
                        }
                        _ = v35
                    } else {
                        fmt.Println(v2, v3)
                        for fi41 := 0; fi41 < 11; fi41++ {
                            if true {
                                switch fi31 {
                                case 627:
                                    switch (("foo bar" + "hello") + ("hello" + ("foo bar" + ("hello" + ("" + (("hello" + ("a" + "\n\t")) + (((("\n\t" + "hello") + "") + "") + (("a" + "") + "\n\t")))))))) {
                                    case "foo bar":
                                        var v42 uint = (v33 * (v33 >> ((v33) & 63)))
                                        switch "hello" {
                                        case "":
                                            var v43 uint32 = 252
                                            _ = v43
                                        }
                                        var v44 float32 = 0.0
                                        _ = v42
                                        _ = v44
                                    }
                                    v45 := []int{}
                                    if true {
                                        v45 = append(v45, (fi41 << (((v33 * v33)) & 63)))
                                    }
                                    _ = v45
                                }
                                fmt.Println(v33)
                            }
                            var v46 float32 = 0.0
                            if false {
                                v46 -= v46
                            }
                            _ = fi41
                            _ = v46
                        }
                    }
                    _ = v34
                }
                var v47 uint8 = uint8(v2)
                _ = v33
                _ = v47
            }
        case "foo bar":
            v48 := []int{}
            _ = v48
        case "\n\t":
            for fi49 := 0; fi49 < 16; fi49++ {
                v2 = (^(v2))
                if (703 == 182) {
                    v32 -= v32
                    v32 -= v5
                }
                v50 := []int{(-(fi49)), fi49, 42}
                _ = fi49
                _ = v50
            }
        default:
            v51 := GenS1{F0: 1, F1: false}
            _ = v51
        }
        v52 := []int{int(v3), fi31, fi31}
        _ = fi31
        _ = v32
        _ = v52
    }
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        if ((1 >= 1) && (false || true)) {
            v53 := GenS7{F0: 33, F1: 1, F2: 128}
            v54 := []int{-389, 32768}
            _ = v53
            _ = v54
        } else {
            v55 := GenS8{F0: 0}
            switch 435 {
            case -592:
                var v56 float64 = v1
                _ = v56
            case 1:
                var v57 float64 = float64((0.0 + float32(3.14159)))
                _ = v57
            case -545:
                v4 = v4
            }
            _ = v55
        }
    }
    for fi58 := 0; fi58 < 14; fi58++ {
        var v59 uint = uint(v5)
        _ = fi58
        _ = v59
    }
    var v60 int8 = 58
    if (v4 != v4) {
        for fi61 := 0; fi61 < 6; fi61++ {
            for fi62 := 0; fi62 < 20; fi62++ {
                if (0 > 276) {
                    v63 := GenS9{F0: 2, F1: v4, F2: (v5 ^ (v5 + 0))}
                    v1 += 3.14159
                    var v64 int8 = v60
                    _ = v63
                    _ = v64
                } else {
                    fmt.Println(v1, fi61, v1)
                    if true {
                        var v65 bool = false
                        v66 := GenS4{F0: -1}
                        switch (fi61 & (fi61 & (fi61 | (fi62 - (-(fi62)))))) {
                        case -520:
                            v4 = (v4 / ((v4) | 1))
                            switch int(v3) {
                            case 127:
                                switch "\n\t" {
                                case "a":
                                    switch "a" {
                                    case "foo bar":
                                        fmt.Println(v5, v60)
                                    case "a":
                                        var v67 int8 = 1
                                        fmt.Println(v67, v2)
                                        v68 := GenS9{F0: 127, F1: v4, F2: 65535}
                                        _ = v67
                                        _ = v68
                                    case "":
                                        v2 += v2
                                    default:
                                        var v69 bool = (!true)
                                        _ = v69
                                    }
                                case "\n\t":
                                    var v70 float32 = 1e10
                                    _ = v70
                                default:
                                    var v71 int32 = v4
                                    _ = v71
                                }
                            default:
                                var v72 string = ""
                                _ = v72
                            }
                            v73 := GenS10{F0: 1, F1: 127}
                            _ = v73
                        }
                        _ = v65
                        _ = v66
                    }
                    if true {
                        if false {
                            var v74 string = "a"
                            v75 := GenS1{F0: fi61, F1: ((!((v1 / v1) == v1)) || true)}
                            if ((false || ((fi61 > int(v60)) && (!true))) || (!((0 != 169) && false))) {
                                var v76 int = fi61
                                var v77 uint = 4294967295
                                _ = v76
                                _ = v77
                            } else {
                                v3 = v3
                                var v78 int = 9223372036854775807
                                _ = v78
                            }
                            _ = v74
                            _ = v75
                        }
                        v60 += (v60 % ((v60) | 1))
                        switch 32767 {
                        case 896:
                            switch ("" + ("" + ((("a" + ("foo bar" + "")) + (("a" + ("\n\t" + ("hello" + (("foo bar" + ((((("" + "\n\t") + "\n\t") + "foo bar") + ((("a" + "a") + "a") + "foo bar")) + "hello")) + "hello")))) + "")) + "\n\t"))) {
                            case "":
                                v79 := []int{(fi62 | fi61), 1, int(v5), 42}
                                _ = v79
                            case "a":
                                v80 := []int{fi62, 32768, fi62, (fi62 / ((fi61) | 1))}
                                _ = v80
                            case "foo bar":
                                var v81 int32 = v4
                                _ = v81
                            }
                        case -9223372036854775808:
                            v82 := GenS4{F0: 32767}
                            _ = v82
                        case -777:
                            switch fi62 {
                            case 0:
                                v83 := GenS2{F0: uint(v3), F1: int32(fi61), F2: 145}
                                _ = v83
                            case -2147483648:
                                var v84 uint8 = 227
                                _ = v84
                            case -929:
                                switch (fi61 * 127) {
                                case -2147483648:
                                    fmt.Println(v4, fi62, v1)
                                case 32768:
                                    v5 = 453
                                    switch (("a" + ("\n\t" + "hello")) + ("foo bar" + ("foo bar" + "a"))) {
                                    case "":
                                        var v85 int = 42
                                        _ = v85
                                    case "hello":
                                        v86 := []int{fi61, fi62}
                                        var v87 int64 = 162
                                        if (true && true) {
                                            fmt.Println(fi61, v5, v4)
                                        } else {
                                            v1 -= float64((1.5 * 1.5))
                                        }
                                        v86 = append(v86, (fi61 / ((-1) | 1)))
                                        if false {
                                            if len(v86) > 0 {
                                                v86[0] = -1
                                            }
                                            if ((v4 ^ (^(v4))) < int32(v2)) {
                                                switch ((("hello" + "") + "\n\t") + "foo bar") {
                                                case "":
                                                    if false {
                                                        if false {
                                                            if len(v86) > 0 {
                                                                v86[0] = (fi62 & fi61)
                                                            }
                                                        } else {
                                                            var v88 uint64 = 0
                                                            var v89 uint = 419
                                                            v2 = uint64(v89)
                                                            v2 = v88
                                                            _ = v88
                                                            _ = v89
                                                        }
                                                        v60 = (v60 + (v60 % (((v60 << ((4294967295) & 63))) | 1)))
                                                        if (!("" > "")) {
                                                            _ = len(v86)
                                                            var v90 string = ("a" + ("a" + ""))
                                                            _ = v90
                                                        }
                                                    }
                                                    switch fi62 {
                                                    case 9223372036854775807:
                                                        _ = len(v86)
                                                        v86 = append(v86, 929)
                                                    }
                                                    v87 += 777
                                                    var v91 int8 = -128
                                                    _ = v91
                                                }
                                            } else {
                                                v2 -= uint64(v5)
                                                var v92 int16 = 1
                                                var v93 uint8 = 2
                                                _ = v92
                                                _ = v93
                                            }
                                            v94 := GenS9{F0: 170, F1: (^(v4)), F2: 2147483647}
                                            _ = v94
                                        }
                                        _ = v86
                                        _ = v87
                                    }
                                case 2147483647:
                                    v95 := []int{int(v4), (-(fi61)), (fi62 * fi61), fi62}
                                    _ = v95
                                }
                            default:
                                v96 := []int{fi61, (fi62 & -90), fi62}
                                _ = v96
                            }
                        default:
                            switch ((("" + "foo bar") + "a") + ("\n\t" + "")) {
                            case "":
                                fmt.Println(v60)
                                if ((v2 / (((^(v2))) | 1)) > (v2 << ((866) & 63))) {
                                    var v97 int32 = v4
                                    switch 32768 {
                                    case 549:
                                        if true {
                                            var v98 int8 = int8(v5)
                                            v3 = 327
                                            _ = v98
                                        } else {
                                            fmt.Println(v60, v1)
                                            var v99 uint16 = (v3 ^ v3)
                                            switch "hello" {
                                            case "foo bar":
                                                switch "" {
                                                case "foo bar":
                                                    var v100 uint8 = 0
                                                    _ = v100
                                                case "\n\t":
                                                    v101 := GenS8{F0: (-(v97))}
                                                    _ = v101
                                                case "hello":
                                                    v97 -= int32(v2)
                                                }
                                            case "hello":
                                                if true {
                                                    if false {
                                                        v4 -= 42
                                                    }
                                                    switch fi61 {
                                                    case 42:
                                                        var v102 uint64 = v2
                                                        _ = v102
                                                    case 329:
                                                        v103 := []int{(fi61 + fi61), fi61, 127}
                                                        _ = v103
                                                    }
                                                    v60 = v60
                                                } else {
                                                    if (-9223372036854775808 <= 0) {
                                                        if false {
                                                            var v104 int16 = -1
                                                            if ("hello" >= "\n\t") {
                                                                if true {
                                                                    v105 := GenS11{F0: (v5 >> ((2) & 63)), F1: 42, F2: (fi62 | fi62)}
                                                                    _ = v105
                                                                }
                                                                var v106 string = ("a" + ("\n\t" + ""))
                                                                v5 -= v5
                                                                _ = v106
                                                            }
                                                            var v107 string = "a"
                                                            _ = v104
                                                            _ = v107
                                                        }
                                                        var v108 int64 = 42
                                                        var v109 uint64 = v2
                                                        fmt.Println(fi62, v60)
                                                        v110 := []int{(fi62 << ((211) & 63)), fi61, (-(fi61)), fi62}
                                                        _ = v108
                                                        _ = v109
                                                        _ = v110
                                                    }
                                                    v5 = v5
                                                }
                                            default:
                                                v111 := []int{(fi62 % (((-(fi62))) | 1))}
                                                _ = v111
                                            }
                                            v112 := GenS6{F0: "hello", F1: 32767, F2: 127}
                                            _ = v112
                                            _ = v99
                                        }
                                    default:
                                        switch (fi61 % ((fi61) | 1)) {
                                        case 128:
                                            if (624 == 0) {
                                                var v113 uint = 538
                                                switch ((("hello" + "hello") + (("foo bar" + "hello") + "a")) + ("\n\t" + "hello")) {
                                                case "a":
                                                    v60 = (v60 / (((v60 & v60)) | 1))
                                                default:
                                                    v114 := []int{-625, 128}
                                                    _ = v114
                                                }
                                                var v115 bool = false
                                                _ = v113
                                                _ = v115
                                            } else {
                                                switch "\n\t" {
                                                case "hello":
                                                    if true {
                                                        v60 -= (v60 << ((0) & 63))
                                                    }
                                                case "":
                                                    var v116 bool = (v1 != 0.1)
                                                    _ = v116
                                                case "a":
                                                    v117 := GenS11{F0: 691, F1: (v3 << ((255) & 63)), F2: (-(fi62))}
                                                    _ = v117
                                                default:
                                                    fmt.Println(v60, fi61)
                                                }
                                            }
                                        case 9223372036854775807:
                                            switch "hello" {
                                            case "":
                                                v118 := []int{(fi61 * (fi61 * fi62)), fi62}
                                                _ = v118
                                            default:
                                                switch "foo bar" {
                                                case "\n\t":
                                                    var v119 string = "hello"
                                                    v120 := GenS12{F0: 323, F1: 127}
                                                    _ = v119
                                                    _ = v120
                                                }
                                            }
                                        default:
                                            v121 := GenS1{F0: fi62, F1: (128 == 0)}
                                            _ = v121
                                        }
                                    }
                                    v122 := []int{(fi62 ^ -2147483648), fi62, int(v5), (fi61 % (((fi61 | fi62)) | 1))}
                                    _ = v122
                                    _ = v97
                                } else {
                                    v60 = (-(v60))
                                    if (65535 >= 255) {
                                        if true {
                                            v5 = (v5 << ((18446744073709551615) & 63))
                                            v3 = 255
                                            v123 := GenS13{F0: true, F1: v2}
                                            _ = v123
                                        } else {
                                            switch -999 {
                                            case 243:
                                                switch (fi62 * fi61) {
                                                case 4294967295:
                                                    fmt.Println(v3)
                                                case -9223372036854775808:
                                                    if (false || false) {
                                                        if ((!(v4 < (^(v4)))) && false) {
                                                            v124 := GenS12{F0: 1, F1: 255}
                                                            _ = v124
                                                        }
                                                        v125 := []int{32767, -1}
                                                        _ = v125
                                                    } else {
                                                        v3 = (^(v3))
                                                    }
                                                case 32767:
                                                    var v126 uint32 = (v5 - v5)
                                                    _ = v126
                                                }
                                            }
                                            v127 := []int{}
                                            var v128 float64 = (v1 / (0.1 * ((((((v1 * (float64((1.5 + 1.5)) + ((v1 - (3.14159 / (float64(-2.25) - v1))) - 1.5))) * 1.5) * ((v1 - v1) + (0.1 / v1))) / 1.0) - 1.5) - v1)))
                                            var v129 uint32 = (v5 * v5)
                                            _ = v127
                                            _ = v128
                                            _ = v129
                                        }
                                        v130 := GenS4{F0: 32767}
                                        v4 = int32(fi61)
                                        _ = v130
                                    }
                                    if (!false) {
                                        switch "foo bar" {
                                        case "a":
                                            if false {
                                                v131 := []int{int(v60)}
                                                v132 := GenS2{F0: 65535, F1: v4, F2: 255}
                                                if false {
                                                    var v133 string = "a"
                                                    v1 = v1
                                                    _ = v133
                                                } else {
                                                    var v134 float64 = v1
                                                    var v135 float32 = (1e10 + 1.5)
                                                    v136 := GenS4{F0: -1}
                                                    _ = v134
                                                    _ = v135
                                                    _ = v136
                                                }
                                                _ = len(v131)
                                                _ = v131
                                                _ = v132
                                            }
                                        default:
                                            fmt.Println(v4, v2, v3)
                                        }
                                    } else {
                                        if false {
                                            if false {
                                                var v137 int64 = 42
                                                _ = v137
                                            } else {
                                                fmt.Println(fi62, v60, fi61)
                                                v138 := []int{(fi61 >> ((uint(v2)) & 63)), (fi61 | fi61)}
                                                _ = v138
                                            }
                                            var v139 float32 = ((-2.25 / -2.25) + 0.0)
                                            _ = v139
                                        } else {
                                            var v140 int = 939
                                            v141 := []int{822, (fi61 | 0), (fi62 + 32768)}
                                            var v142 int32 = 42
                                            _ = v140
                                            _ = v141
                                            _ = v142
                                        }
                                    }
                                }
                            case "hello":
                                fmt.Println(fi62, v1)
                            case "foo bar":
                                var v143 uint32 = (v5 - 112)
                                _ = v143
                            }
                        }
                    }
                    fmt.Println(fi62, v60)
                    v144 := []int{}
                    _ = v144
                }
                _ = fi62
            }
            v3 -= (v3 + (^(v3)))
            _ = fi61
        }
    } else {
        for fi145 := 0; fi145 < 18; fi145++ {
            fmt.Println(v3, v3, v3)
            var v146 uint64 = 42
            _ = fi145
            _ = v146
        }
        var v147 uint16 = v3
        v148 := []int{-129}
        _ = v147
        _ = v148
    }
    for fi149 := 0; fi149 < 12; fi149++ {
        switch fi149 {
        case -1:
            v150 := GenS14{F0: int8(v5), F1: 1}
            _ = v150
        case 128:
            if true {
                var v151 int8 = int8(v4)
                v3 = v3
                _ = v151
            }
        case 4294967295:
            switch (fi149 | -475) {
            case 2147483647:
                if true {
                    var v152 int16 = 42
                    for fi153 := 0; fi153 < 18; fi153++ {
                        if false {
                            var v154 int16 = v152
                            v60 = -1
                            _ = v154
                        } else {
                            v3 += v3
                            if false {
                                v155 := []int{fi149, int(v5), fi149}
                                fmt.Println(v3, v60)
                                _ = v155
                            }
                            fmt.Println(v3)
                        }
                        switch 32767 {
                        case 4294967295:
                            if (v152 != (v152 >> ((308) & 63))) {
                                var v156 int = fi153
                                _ = v156
                            }
                        case -9223372036854775808:
                            if ((true && false) && true) {
                                if (4294967295 <= 255) {
                                    v157 := GenS15{F0: (fi153 & (fi153 % ((fi153) | 1))), F1: 57}
                                    var v158 int64 = -1
                                    var v159 uint8 = uint8(v2)
                                    _ = v157
                                    _ = v158
                                    _ = v159
                                }
                            }
                        case 749:
                            switch fi153 {
                            case -283:
                                switch int(v3) {
                                case 42:
                                    v160 := []int{}
                                    _ = v160
                                case 626:
                                    var v161 uint16 = (^(v3))
                                    _ = v161
                                case 32767:
                                    switch fi149 {
                                    case 128:
                                        v4 = int32(v5)
                                    default:
                                        v162 := []int{(fi153 * (fi153 % (((fi153 & 127)) | 1))), fi153, int(v152), fi149}
                                        _ = v162
                                    }
                                }
                            default:
                                fmt.Println(v4, v5)
                                v163 := GenS16{F0: v60, F1: v60, F2: uint16(v5)}
                                fmt.Println(fi153, v4, v3)
                                _ = v163
                            }
                        }
                        fmt.Println(v60)
                        _ = fi153
                    }
                    _ = v152
                }
            case -349:
                v2 = (v2 >> ((42) & 63))
            case 9223372036854775807:
                fmt.Println(v3, v4)
            default:
                for fi164 := 0; fi164 < 10; fi164++ {
                    if ((-2.25 + 1.5) > 3.14159) {
                        var v165 float32 = -2.25
                        if (0 == v3) {
                            fmt.Println(fi164, v3)
                            if (v5 <= v5) {
                                var v166 int16 = -32768
                                _ = v166
                            } else {
                                var v167 uint32 = (v5 - 168)
                                var v168 uint16 = 273
                                v4 = v4
                                _ = v167
                                _ = v168
                            }
                        }
                        fmt.Println(v60)
                        _ = v165
                    }
                    _ = fi164
                }
            }
        }
        _ = fi149
    }
    fmt.Println(v4, v3)
    v4 -= (v4 + -1)
    v169 := []int{}
    v3 = (v3 >> ((2) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v169:", len(v169), v169)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v14
    _ = v15
    _ = v169
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v60
}
