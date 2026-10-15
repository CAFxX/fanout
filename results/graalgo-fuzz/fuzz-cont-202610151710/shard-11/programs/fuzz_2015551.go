package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint32; F1 uint; F2 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint32; F1 uint }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int8; F1 uint64 }
type GenS9 struct { F0 int16; F1 int16 }
type GenS10 struct { F0 uint16; F1 int32 }
type GenS11 struct { F0 uint8; F1 uint; F2 bool }
func genFn0(p0 string, p1 int16, p2 uint8) bool {
    v4 := GenS1{F0: 2147483648, F1: 2}
    v5 := GenS2{F0: p0, F1: 2}
    v6 := GenS1{F0: 656, F1: 18446744073709551615}
    fmt.Println(p0)
    var v7 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return (9223372036854775807 >= -9223372036854775808)
}

func main() {
    v1 := []int{4294967295}
    var v2 string = "a"
    var v3 int32 = 0
    switch v2 {
    case "":
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    case "foo bar":
        var v8 uint8 = 42
        _ = v8
    default:
        var v9 float32 = 3.14159
        _ = v9
    }
    v2 = ("a" + (((v2 + v2) + v2) + v2))
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 int32 = v3
        v12 := GenS1{F0: 1, F1: 18446744073709551615}
        v2 = "hello"
        _ = fi10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        if (42 <= 2) {
            if (-566 > -135) {
                fmt.Println(fi13)
                v14 := GenS1{F0: 107, F1: 4294967295}
                _ = v14
            } else {
                v15 := GenS3{F0: 1, F1: 206, F2: 100}
                _ = v15
            }
            if len(v1) > 0 {
                v1[0] = (fi13 & fi13)
            }
            v16 := GenS2{F0: v2, F1: 343}
            _ = v16
        } else {
            var v17 uint16 = 65535
            var v18 int16 = 0
            _ = v17
            _ = v18
        }
        _ = fi13
    }
    var v19 uint64 = 2
    if false {
        _ = len(v1)
        v2 = v2
    }
    if (true && true) {
        if false {
            if false {
                var v20 uint64 = 18446744073709551615
                if (!((!(!false)) || true)) {
                    var v21 bool = true
                    v19 = v19
                    _ = v21
                } else {
                    if true {
                        var v22 uint32 = 845
                        if true {
                            if false {
                                v23 := GenS4{F0: 0}
                                _ = v23
                            } else {
                                if true {
                                    v2 = ("hello" + "\n\t")
                                    switch "hello" {
                                    case "":
                                        _ = len(v1)
                                    case "\n\t":
                                        v2 = "a"
                                    case "a":
                                        fmt.Println(v20, v22, v19)
                                    }
                                    for fi24 := 0; fi24 < 9; fi24++ {
                                        if false {
                                            var v25 int32 = -1
                                            v26 := GenS3{F0: (v22 | (v22 ^ v22)), F1: 255, F2: 42}
                                            for fi27 := 0; fi27 < 13; fi27++ {
                                                switch v2 {
                                                case "foo bar":
                                                    var v28 uint = 915
                                                    _ = v28
                                                case "\n\t":
                                                    if (-2147483648 > (fi27 & (fi24 | (^(fi27))))) {
                                                        fmt.Println(v25, fi27, v20)
                                                    } else {
                                                        v2 = v2
                                                        if (false || (-1 <= 0)) {
                                                            var v29 int8 = 1
                                                            if false {
                                                                switch fi27 {
                                                                case -960:
                                                                    if false {
                                                                        var v30 float64 = ((0.0 * 3.14159) + (0.1 + 1.5))
                                                                        var v31 int64 = -1
                                                                        _ = v30
                                                                        _ = v31
                                                                    }
                                                                    fmt.Println(v2, v29, v20)
                                                                    switch v2 {
                                                                    case "foo bar":
                                                                        v25 -= v25
                                                                        var v32 int32 = (^(v3))
                                                                        _ = v32
                                                                    case "a":
                                                                        var v33 int32 = v3
                                                                        _ = v33
                                                                    case "":
                                                                        var v34 uint8 = 255
                                                                        _ = v34
                                                                    default:
                                                                        if (((-(fi27)) > -1) || true) {
                                                                            v22 = 65535
                                                                            switch (fi24 / (((fi27 | 128)) | 1)) {
                                                                            case 42:
                                                                                switch v2 {
                                                                                case "":
                                                                                    var v35 float64 = ((-2.25 / 1.5) - (3.14159 + (float64(3.14159) / 3.14159)))
                                                                                    _ = v35
                                                                                }
                                                                            case 2147483647:
                                                                                _ = len(v1)
                                                                            default:
                                                                                v36 := GenS1{F0: v22, F1: 630}
                                                                                _ = v36
                                                                            }
                                                                            v37 := GenS5{F0: 42}
                                                                            _ = v37
                                                                        }
                                                                    }
                                                                case 1:
                                                                    v2 = (v2 + v2)
                                                                default:
                                                                    v3 = -2147483648
                                                                }
                                                                if true {
                                                                    switch (fi24 % (((fi24 / (((fi24 ^ 32768)) | 1))) | 1)) {
                                                                    case -9223372036854775808:
                                                                        if true {
                                                                            var v38 float64 = 1.5
                                                                            if len(v1) > 0 {
                                                                                v1[0] = 1
                                                                            }
                                                                            v19 -= 4294967295
                                                                            _ = v38
                                                                        } else {
                                                                            fmt.Println(v19)
                                                                        }
                                                                    }
                                                                    var v39 uint32 = 1
                                                                    v19 = (^(v20))
                                                                    v22 += 534
                                                                    _ = v39
                                                                }
                                                            } else {
                                                                v19 += v19
                                                            }
                                                            var v40 string = "foo bar"
                                                            _ = v29
                                                            _ = v40
                                                        } else {
                                                            if false {
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                if false {
                                                                    if (1 >= 0) {
                                                                        v1 = append(v1, fi24)
                                                                        v19 = (v19 | v19)
                                                                        v20 = v19
                                                                    }
                                                                }
                                                                v20 = (v20 ^ 4294967295)
                                                                switch v2 {
                                                                case "":
                                                                    var v41 uint32 = (v22 * v22)
                                                                    _ = v41
                                                                case "hello":
                                                                    v25 = v25
                                                                default:
                                                                    var v42 int64 = -1
                                                                    v19 = v19
                                                                    v43 := GenS6{F0: v22, F1: 2}
                                                                    _ = v42
                                                                    _ = v43
                                                                }
                                                            } else {
                                                                switch (-(fi24)) {
                                                                case 32768:
                                                                    fmt.Println(v20, v2, v20)
                                                                default:
                                                                    v44 := GenS3{F0: 65535, F1: 139, F2: -13}
                                                                    _ = v44
                                                                }
                                                                switch (fi24 + fi27) {
                                                                case 937:
                                                                    v20 = v19
                                                                case 2147483647:
                                                                    _ = len(v1)
                                                                default:
                                                                    if (!(true && false)) {
                                                                        v25 += (^(v25))
                                                                        var v45 int32 = v25
                                                                        var v46 float32 = 0.0
                                                                        _ = v45
                                                                        _ = v46
                                                                    }
                                                                }
                                                                fmt.Println(fi24)
                                                            }
                                                            switch "" {
                                                            case "":
                                                                var v47 uint16 = 867
                                                                _ = v47
                                                            default:
                                                                switch fi24 {
                                                                case 1:
                                                                    fmt.Println(v2, v20, v19)
                                                                default:
                                                                    var v48 uint64 = 18446744073709551615
                                                                    _ = v48
                                                                }
                                                            }
                                                            var v49 uint = 65535
                                                            _ = v49
                                                        }
                                                        v19 += (v19 + v20)
                                                    }
                                                case "hello":
                                                    if (127 < 20) {
                                                        var v50 uint32 = v22
                                                        v51 := GenS6{F0: 2, F1: 304}
                                                        _ = v50
                                                        _ = v51
                                                    }
                                                }
                                                var v52 bool = (0 >= 920)
                                                _ = fi27
                                                _ = v52
                                            }
                                            _ = v25
                                            _ = v26
                                        } else {
                                            var v53 int32 = 998
                                            _ = v53
                                        }
                                        var v54 int32 = (^(v3))
                                        var v55 uint32 = 0
                                        _ = fi24
                                        _ = v54
                                        _ = v55
                                    }
                                } else {
                                    for fi56 := 0; fi56 < 14; fi56++ {
                                        v19 = 18446744073709551615
                                        v2 = v2
                                        v57 := GenS1{F0: (v22 ^ (v22 >> ((4294967295) & 63))), F1: 255}
                                        _ = fi56
                                        _ = v57
                                    }
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                        } else {
                            var v58 int = -178
                            var v59 uint64 = (v19 ^ v20)
                            _ = v58
                            _ = v59
                        }
                        v3 = v3
                        _ = v22
                    }
                }
                _ = v20
            } else {
                if len(v1) > 0 {
                    v1[0] = 2147483647
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v60 uint8 = 207
                _ = v60
            }
        }
        v3 += 42
        fmt.Println(v2)
    } else {
        fmt.Println(v3)
        switch v2 {
        case "a":
            fmt.Println(v19)
        case "\n\t":
            fmt.Println(v19, v3)
        default:
            var v61 uint16 = 255
            _ = v61
        }
    }
    var v62 float32 = (float32(0.1) / 1.5)
    if ((1.5 / 1.0) <= float64((v62 + (0.1 + v62)))) {
        var v63 float32 = v62
        v64 := GenS6{F0: 2147483647, F1: 65535}
        v63 = 1.5
        _ = v63
        _ = v64
    } else {
        for fi65 := 0; fi65 < 20; fi65++ {
            v62 += 0.1
            fmt.Println(v2, v3)
            _ = fi65
        }
    }
    v19 += v19
    v19 += v19
    if (!(1 == 2)) {
        v66 := GenS3{F0: 206, F1: 0, F2: -41}
        if ((65535 >= 255) && true) {
            switch -295 {
            case 32767:
                switch -361 {
                case 9223372036854775807:
                    v19 = 0
                case 42:
                    for fi67 := 0; fi67 < 6; fi67++ {
                        v19 = v19
                        _ = fi67
                    }
                case -504:
                    var v68 uint8 = 242
                    _ = v68
                default:
                    v62 = 1e10
                }
            case 128:
                if true {
                    var v69 float64 = -2.25
                    v70 := GenS3{F0: 4294967295, F1: 4294967295, F2: 19}
                    _ = v69
                    _ = v70
                }
            }
            if (18 <= 252) {
                if (((^(v3)) > (v3 * 0)) && (0 == -1)) {
                    var v71 uint16 = 65535
                    _ = v71
                }
                for fi72 := 0; fi72 < 1; fi72++ {
                    v73 := GenS3{F0: 2, F1: 1, F2: 100}
                    _ = fi72
                    _ = v73
                }
            }
            if (!false) {
                for fi74 := 0; fi74 < 18; fi74++ {
                    v75 := GenS4{F0: 32767}
                    switch v2 {
                    case "hello":
                        var v76 string = ("foo bar" + "")
                        _ = v76
                    case "a":
                        if ((v2 + ("" + v2)) <= ((v2 + "hello") + v2)) {
                            var v77 uint16 = 1
                            var v78 uint32 = 185
                            _ = v77
                            _ = v78
                        } else {
                            v62 = ((0.1 + v62) / v62)
                            if true {
                                if true {
                                    switch v2 {
                                    case "hello":
                                        for fi79 := 0; fi79 < 4; fi79++ {
                                            var v80 int = -2147483648
                                            switch "" {
                                            case "\n\t":
                                                var v81 int16 = -32768
                                                _ = v81
                                            case "foo bar":
                                                v62 = (v62 / (0.0 * (float32(float64((v62 * v62))) / v62)))
                                            case "":
                                                _ = len(v1)
                                            }
                                            fmt.Println(v80, v2)
                                            _ = fi79
                                            _ = v80
                                        }
                                    }
                                    fmt.Println(v2, v19, v2)
                                    v82 := GenS3{F0: 702, F1: 211, F2: -128}
                                    _ = v82
                                } else {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    switch fi74 {
                                    case -2147483648:
                                        var v83 int16 = -1
                                        _ = v83
                                    case 128:
                                        v3 += 0
                                    }
                                    for fi84 := 0; fi84 < 11; fi84++ {
                                        switch v2 {
                                        case "a":
                                            switch v2 {
                                            case "":
                                                var v85 int16 = -688
                                                _ = v85
                                            case "foo bar":
                                                if (608 < 267) {
                                                    v3 = (^(v3))
                                                    var v86 int8 = -128
                                                    _ = v86
                                                }
                                            case "a":
                                                var v87 float32 = 1e10
                                                _ = v87
                                            default:
                                                switch fi84 {
                                                case -2147483648:
                                                    var v88 bool = false
                                                    _ = v88
                                                case 32767:
                                                    if true {
                                                        if true {
                                                            v2 = (v2 + v2)
                                                        }
                                                        v89 := GenS1{F0: 42, F1: 152}
                                                        v19 = 1
                                                        _ = v89
                                                    }
                                                }
                                            }
                                        case "":
                                            var v90 uint64 = 2
                                            _ = v90
                                        case "foo bar":
                                            var v91 int64 = 1
                                            _ = v91
                                        default:
                                            fmt.Println(fi84, v3, v3)
                                        }
                                        v92 := GenS7{F0: 1}
                                        switch fi84 {
                                        case -698:
                                            var v93 int64 = 9223372036854775807
                                            _ = v93
                                        case 4294967295:
                                            var v94 bool = true
                                            v95 := GenS5{F0: 0}
                                            var v96 float64 = 1.5
                                            _ = v94
                                            _ = v95
                                            _ = v96
                                        case 1:
                                            v62 -= v62
                                        }
                                        _ = fi84
                                        _ = v92
                                    }
                                }
                                var v97 float64 = 1.5
                                for fi98 := 0; fi98 < 11; fi98++ {
                                    var v99 uint64 = (v19 * 4294967295)
                                    _ = fi98
                                    _ = v99
                                }
                                _ = v97
                            }
                            switch (v2 + ((("foo bar" + ("hello" + v2)) + v2) + "foo bar")) {
                            case "foo bar":
                                for fi100 := 0; fi100 < 4; fi100++ {
                                    if len(v1) > 0 {
                                        v1[0] = 128
                                    }
                                    switch "" {
                                    case "hello":
                                        var v101 uint = 18446744073709551615
                                        _ = v101
                                    case "a":
                                        v19 -= v19
                                    }
                                    _ = fi100
                                }
                            default:
                                for fi102 := 0; fi102 < 1; fi102++ {
                                    v103 := GenS8{F0: 57, F1: (v19 << ((2) & 63))}
                                    _ = fi102
                                    _ = v103
                                }
                            }
                        }
                    case "foo bar":
                        fmt.Println(v62, v19)
                    }
                    var v104 float32 = -2.25
                    _ = fi74
                    _ = v104
                    _ = v75
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v1 = append(v1, 32767)
            }
        }
        for fi105 := 0; fi105 < 16; fi105++ {
            if false {
                v19 = (v19 % (((v19 * (^(v19)))) | 1))
            }
            var v106 string = "\n\t"
            _ = fi105
            _ = v106
        }
        _ = v66
    } else {
        var v107 uint32 = 42
        var v108 uint64 = (v19 << ((1) & 63))
        _ = v107
        _ = v108
    }
    switch 942 {
    case 4294967295:
        if (3.14159 < ((v62 - 0.0) + v62)) {
            v62 = v62
            var v109 float32 = (((v62 + (((v62 * 1e10) * v62) + ((v62 - 1e10) * (0.0 + (((0.1 + 0.1) * ((1.5 + v62) + 1e10)) - float32(1.5)))))) - 1e10) - v62)
            if false {
                var v110 uint16 = 65535
                v19 += v19
                _ = v110
            }
            _ = v109
        }
    case -545:
        if ((v19 % (((v19 / (((v19 % (((v19 << ((18446744073709551615) & 63))) | 1))) | 1))) | 1)) <= 1) {
            fmt.Println(v2, v19)
        }
    case 394:
        var v111 string = "\n\t"
        _ = v111
    }
    if true {
        v112 := GenS9{F0: -1, F1: 1}
        _ = v112
    } else {
        if true {
            var v113 int8 = 100
            _ = v113
        }
        switch (v2 + v2) {
        case "foo bar":
            var v114 string = v2
            _ = v114
        case "hello":
            for fi115 := 0; fi115 < 3; fi115++ {
                if false {
                    v19 = (v19 << ((603) & 63))
                    v116 := GenS1{F0: 2, F1: 2}
                    _ = v116
                }
                _ = fi115
            }
        case "":
            if false {
                v117 := GenS9{F0: 0, F1: -1}
                v19 -= (v19 & (v19 << ((224) & 63)))
                _ = v117
            }
        }
        if false {
            v2 = "foo bar"
            for fi118 := 0; fi118 < 18; fi118++ {
                if true {
                    switch "a" {
                    case "hello":
                        v3 = (v3 % (((v3 >> ((255) & 63))) | 1))
                    default:
                        var v119 int64 = -69
                        _ = v119
                    }
                    _ = len(v1)
                }
                _ = fi118
            }
        } else {
            if true {
                fmt.Println(v62, v2)
                for fi120 := 0; fi120 < 16; fi120++ {
                    for fi121 := 0; fi121 < 9; fi121++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi121
                    }
                    switch (^(fi120)) {
                    case 32768:
                        v122 := GenS4{F0: 1}
                        _ = v122
                    case -672:
                        var v123 uint32 = 65535
                        _ = v123
                    case 844:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    fmt.Println(v3, v19)
                    _ = fi120
                }
                if (false && (("\n\t" + ((((("foo bar" + v2) + (v2 + v2)) + "foo bar") + ("\n\t" + v2)) + v2)) < "a")) {
                    for fi124 := 0; fi124 < 10; fi124++ {
                        var v125 uint64 = (v19 * (v19 - v19))
                        for fi126 := 0; fi126 < 6; fi126++ {
                            _ = len(v1)
                            _ = fi126
                        }
                        _ = fi124
                        _ = v125
                    }
                    v127 := GenS10{F0: 65535, F1: 454}
                    var v128 bool = ((float64(v62) * 1.5) > -2.25)
                    _ = v127
                    _ = v128
                }
            } else {
                v2 = v2
            }
            switch 32768 {
            case -854:
                if true {
                    fmt.Println(v62, v62)
                    for fi129 := 0; fi129 < 1; fi129++ {
                        if (!true) {
                            var v130 float64 = float64(v62)
                            v131 := GenS4{F0: 154}
                            v132 := GenS9{F0: 448, F1: 1}
                            _ = v130
                            _ = v131
                            _ = v132
                        }
                        var v133 int32 = v3
                        for fi134 := 0; fi134 < 6; fi134++ {
                            if false {
                                v19 -= (v19 >> ((0) & 63))
                                var v135 uint16 = 0
                                _ = v135
                            }
                            _ = fi134
                        }
                        _ = fi129
                        _ = v133
                    }
                    for fi136 := 0; fi136 < 20; fi136++ {
                        var v137 int32 = (-(v3))
                        if (!(!false)) {
                            fmt.Println(fi136)
                        } else {
                            v2 = ""
                            if (((1 > 18446744073709551615) || false) || (!false)) {
                                for fi138 := 0; fi138 < 14; fi138++ {
                                    v1 = append(v1, (^(fi136)))
                                    _ = fi138
                                }
                                if false {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    fmt.Println(fi136)
                                    if (65535 > 2) {
                                        if len(v1) > 0 {
                                            v1[0] = (-(fi136))
                                        }
                                        v139 := GenS9{F0: 998, F1: 42}
                                        fmt.Println(fi136, v62, v137)
                                        _ = v139
                                    }
                                }
                            } else {
                                var v140 string = ("\n\t" + v2)
                                v141 := GenS5{F0: -1}
                                _ = v140
                                _ = v141
                            }
                            var v142 int = 42
                            _ = v142
                        }
                        _ = fi136
                        _ = v137
                    }
                }
            case 469:
                for fi143 := 0; fi143 < 10; fi143++ {
                    if false {
                        switch fi143 {
                        case 0:
                            fmt.Println(fi143, v3, v3)
                        case -9223372036854775808:
                            v62 = (v62 / float32(float64(float32(3.14159))))
                        }
                        for fi144 := 0; fi144 < 5; fi144++ {
                            switch fi143 {
                            case 9223372036854775807:
                                var v145 string = v2
                                _ = v145
                            case 2147483647:
                                var v146 int = 38
                                _ = v146
                            case -9223372036854775808:
                                if true {
                                    var v147 int32 = (v3 ^ (^(v3)))
                                    var v148 int32 = v147
                                    switch (fi144 + (^(fi143))) {
                                    case 2147483647:
                                        v19 = (v19 & v19)
                                    case 32767:
                                        if false {
                                            if (0.0 >= ((0.1 + 1.5) * 0.0)) {
                                                fmt.Println(v3, v2)
                                            }
                                            var v149 uint8 = 2
                                            var v150 uint32 = 0
                                            _ = v149
                                            _ = v150
                                        }
                                    default:
                                        if ((203 < 174) || true) {
                                            switch 715 {
                                            case 0:
                                                v19 -= 18446744073709551615
                                                var v151 uint32 = 2
                                                _ = v151
                                            case -118:
                                                v148 = (v148 | v148)
                                            case 479:
                                                switch (-(fi143)) {
                                                case 32768:
                                                    v152 := GenS7{F0: 125}
                                                    _ = v152
                                                case 128:
                                                    var v153 int16 = -32768
                                                    _ = v153
                                                case -1:
                                                    if (1 == 32767) {
                                                        v2 = v2
                                                        v19 -= v19
                                                        if false {
                                                            v154 := GenS7{F0: 1}
                                                            _ = v154
                                                        }
                                                    }
                                                }
                                            }
                                        } else {
                                            var v155 float64 = float64((v62 * v62))
                                            var v156 int64 = 742
                                            v3 += int32(fi143)
                                            v157 := GenS1{F0: 4294967295, F1: 18446744073709551615}
                                            _ = v155
                                            _ = v156
                                            _ = v157
                                        }
                                    }
                                    _ = v147
                                    _ = v148
                                }
                            default:
                                v158 := GenS11{F0: 42, F1: 255, F2: (true && false)}
                                v159 := GenS11{F0: 127, F1: 18446744073709551615, F2: (1 == 807)}
                                _ = v158
                                _ = v159
                            }
                            var v160 uint = 65535
                            _ = fi144
                            _ = v160
                        }
                    }
                    _ = fi143
                }
            }
            v161 := GenS9{F0: 1, F1: 90}
            _ = v161
        }
    }
    for fi162 := 0; fi162 < 4; fi162++ {
        v163 := GenS9{F0: 0, F1: 37}
        _ = len(v1)
        var v164 float32 = 3.14159
        _ = fi162
        _ = v163
        _ = v164
    }
    if true {
        for fi165 := 0; fi165 < 20; fi165++ {
            v166 := GenS10{F0: 36, F1: v3}
            v2 = v2
            _ = fi165
            _ = v166
        }
        v3 += -1
        v167 := GenS9{F0: 846, F1: 42}
        _ = v167
    } else {
        if false {
            fmt.Println(v3)
            fmt.Println(v2, v3, v19)
        }
        if (0 <= 2) {
            v168 := GenS4{F0: 414}
            if true {
                if true {
                    v62 -= v62
                } else {
                    if true {
                        var v169 uint8 = 142
                        _ = v169
                    } else {
                        var v170 uint16 = 507
                        _ = v170
                    }
                    var v171 uint = 2
                    _ = v171
                }
                _ = len(v1)
            } else {
                if false {
                    var v172 uint8 = 0
                    switch (("hello" + "hello") + v2) {
                    case "a":
                        var v173 float32 = v62
                        _ = v173
                    case "":
                        if len(v1) > 0 {
                            v1[0] = 128
                        }
                    }
                    _ = v172
                }
            }
            switch -1 {
            case -9223372036854775808:
                if (!false) {
                    var v174 float64 = -2.25
                    if ((42 != 42) && (127 < -62)) {
                        if true {
                            if false {
                                v3 += v3
                                if false {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v3 = v3
                                    for fi175 := 0; fi175 < 2; fi175++ {
                                        v3 += (^(v3))
                                        for fi176 := 0; fi176 < 18; fi176++ {
                                            if false {
                                                _ = len(v1)
                                                _ = len(v1)
                                                v1 = append(v1, fi176)
                                                var v177 float32 = v62
                                                _ = v177
                                            }
                                            var v178 int32 = (v3 % ((1) | 1))
                                            _ = fi176
                                            _ = v178
                                        }
                                        v179 := GenS8{F0: 42, F1: (v19 | 199)}
                                        _ = fi175
                                        _ = v179
                                    }
                                }
                            } else {
                                var v180 int64 = 42
                                var v181 uint32 = 0
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v180
                                _ = v181
                            }
                            if true {
                                v182 := GenS9{F0: 42, F1: 1000}
                                for fi183 := 0; fi183 < 5; fi183++ {
                                    for fi184 := 0; fi184 < 14; fi184++ {
                                        if false {
                                            switch fi183 {
                                            case -999:
                                                v1 = append(v1, -9)
                                            case 32768:
                                                fmt.Println(fi184, v62, v2)
                                            default:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                var v185 float64 = 0.1
                                                _ = v185
                                            }
                                        } else {
                                            v2 = ""
                                        }
                                        _ = fi184
                                    }
                                    _ = fi183
                                }
                                _ = v182
                            }
                        }
                        v19 = (^(v19))
                        v62 += (((v62 - (v62 * float32(((3.14159 + v174) * 3.14159)))) / v62) + (0.0 + v62))
                    }
                    var v186 int64 = -489
                    _ = v174
                    _ = v186
                }
            }
            _ = v168
        } else {
            if (!(733 > 4294967295)) {
                if (((0 >= 42) || true) && false) {
                    switch 0 {
                    case 128:
                        fmt.Println(v62, v62)
                    default:
                        for fi187 := 0; fi187 < 4; fi187++ {
                            v1 = append(v1, -9223372036854775808)
                            for fi188 := 0; fi188 < 1; fi188++ {
                                var v189 float32 = float32((0.1 / -2.25))
                                if (!((0 != 428) || (2 < v19))) {
                                    var v190 float32 = v189
                                    if (1.5 != (1.5 * float64(v190))) {
                                        fmt.Println(v3)
                                    } else {
                                        if true {
                                            if len(v1) > 0 {
                                                v1[0] = (fi188 * fi188)
                                            }
                                            if (!(26 < 966)) {
                                                if true {
                                                    var v191 uint16 = 0
                                                    _ = v191
                                                }
                                                switch 695 {
                                                case 0:
                                                    if false {
                                                        switch "" {
                                                        case "a":
                                                            v62 = v62
                                                        case "hello":
                                                            v2 = v2
                                                        case "foo bar":
                                                            var v192 uint32 = 0
                                                            _ = v192
                                                        default:
                                                            var v193 float64 = 0.1
                                                            _ = v193
                                                        }
                                                        var v194 int = fi188
                                                        _ = v194
                                                    }
                                                default:
                                                    fmt.Println(v62, v190)
                                                }
                                            } else {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                            var v195 uint16 = 0
                                            _ = v195
                                        }
                                    }
                                    _ = v190
                                } else {
                                    var v196 int16 = 468
                                    switch (-(fi187)) {
                                    case 9223372036854775807:
                                        if len(v1) > 0 {
                                            v1[0] = fi187
                                        }
                                    case 127:
                                        var v197 float64 = 0.1
                                        _ = v197
                                    default:
                                        var v198 uint16 = 0
                                        _ = v198
                                    }
                                    _ = v196
                                }
                                var v199 string = v2
                                _ = fi188
                                _ = v189
                                _ = v199
                            }
                            _ = fi187
                        }
                    }
                    v200 := GenS9{F0: 42, F1: 909}
                    _ = v200
                }
                v1 = append(v1, -9223372036854775808)
            }
            var v201 uint = 42
            var v202 int16 = 42
            _ = v201
            _ = v202
        }
        var v203 int64 = 1
        _ = v203
    }
    var v204 float64 = 0.0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v204:", v204)
    fmt.Println("v3:", v3)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v19
    _ = v2
    _ = v204
    _ = v3
    _ = v62
}
