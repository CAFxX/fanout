package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint8 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 int8; F1 int16 }
type GenS7 struct { F0 bool; F1 uint16; F2 int }
type GenS8 struct { F0 bool; F1 uint64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint16; F1 uint8 }
type GenS12 struct { F0 int64; F1 int32; F2 int8 }
type GenS13 struct { F0 uint32; F1 int }

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{4294967295}
    var v3 int16 = 0
    var v4 int16 = (v3 / ((-32768) | 1))
    var v5 float64 = 1.5
    fmt.Println(v3, v5)
    if false {
        v4 += v3
        for fi6 := 0; fi6 < 13; fi6++ {
            if (1e10 != (0.1 / (0.1 / ((-2.25 / 1e10) * 1.5)))) {
                v4 -= (v3 & v3)
                var v7 uint32 = 65535
                _ = v7
            } else {
                v3 = (v4 & v4)
            }
            switch ((("\n\t" + "hello") + "") + "hello") {
            case "foo bar":
                switch fi6 {
                case -1:
                    var v8 int64 = 781
                    _ = v8
                case -438:
                    v4 -= v3
                default:
                    var v9 uint8 = 42
                    _ = v9
                }
            case "":
                v4 = 42
            default:
                switch fi6 {
                case -1:
                    v4 -= v4
                case 42:
                    var v10 bool = true
                    _ = v10
                case -712:
                    if false {
                        for fi11 := 0; fi11 < 12; fi11++ {
                            v12 := GenS1{F0: 0, F1: -128, F2: 127}
                            _ = fi11
                            _ = v12
                        }
                        v5 = float64(0.1)
                    }
                }
            }
            v5 = v5
            _ = fi6
        }
        for fi13 := 0; fi13 < 9; fi13++ {
            fmt.Println(fi13, v3)
            fmt.Println(fi13)
            _ = fi13
        }
    } else {
        fmt.Println(v5, v4, v3)
        _ = len(v1)
        if false {
            v4 += (^(v3))
            if false {
                delete(v1, 42)
                for fi14 := 0; fi14 < 13; fi14++ {
                    var v15 float32 = 1e10
                    _ = fi14
                    _ = v15
                }
            }
        } else {
            v16 := GenS1{F0: 42, F1: 100, F2: 17}
            var v17 int8 = -128
            _ = v16
            _ = v17
        }
    }
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    v4 += (v3 / ((v4) | 1))
    for fi18 := 0; fi18 < 10; fi18++ {
        if (2 < 928) {
            var v19 int = (fi18 ^ 0)
            switch "foo bar" {
            case "a":
                for fi20 := 0; fi20 < 16; fi20++ {
                    if len(v2) > 0 {
                        v2[0] = 351
                    }
                    var v21 bool = false
                    var v22 float32 = -2.25
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            case "\n\t":
                v3 = (v4 | v4)
            }
            v23 := GenS2{F0: 100, F1: false}
            _ = v19
            _ = v23
        } else {
            for fi24 := 0; fi24 < 2; fi24++ {
                v1[2147483648] = (^(fi24))
                var v25 uint64 = 18446744073709551615
                _ = fi24
                _ = v25
            }
            if true {
                for fi26 := 0; fi26 < 6; fi26++ {
                    var v27 bool = ((fi18 - 32768) < (fi18 >> ((42) & 63)))
                    fmt.Println(v3)
                    if v27 {
                        _ = len(v1)
                        if (!((v3 >> ((0) & 63)) > -135)) {
                            if (v27 || (true && (!true))) {
                                fmt.Println(fi18)
                                v4 = v4
                            }
                            v28 := GenS3{F0: ("foo bar" + "a"), F1: (v3 & (v3 % ((v4) | 1)))}
                            if len(v2) > 0 {
                                v2[0] = 0
                            }
                            v27 = false
                            var v29 float32 = ((((float32(v5) / 3.14159) / 1e10) / 3.14159) + 1.5)
                            _ = v28
                            _ = v29
                        }
                    }
                    _ = fi26
                    _ = v27
                }
                var v30 uint8 = 127
                if (!(!false)) {
                    var v31 uint32 = 2147483647
                    _ = v31
                }
                _ = v30
            } else {
                v3 -= v3
                fmt.Println(v3, v5)
            }
            for fi32 := 0; fi32 < 10; fi32++ {
                v4 = v3
                var v33 float32 = 1e10
                _ = fi32
                _ = v33
            }
        }
        v34 := GenS4{F0: -591, F1: 42}
        v5 += v5
        _ = fi18
        _ = v34
    }
    switch "foo bar" {
    case "hello":
        for fi35 := 0; fi35 < 12; fi35++ {
            var v36 int16 = v3
            fmt.Println(fi35, v36)
            var v37 float64 = v5
            _ = fi35
            _ = v36
            _ = v37
        }
    default:
        v4 -= (v4 >> ((4294967295) & 63))
    }
    v38 := v1[2147483647]
    switch 2147483647 {
    case 1:
        fmt.Println(v38)
    case -970:
        v38 = (v38 | v38)
    case -700:
        var v39 int64 = 9223372036854775807
        _ = v39
    default:
        switch "" {
        case "a":
            for fi40 := 0; fi40 < 7; fi40++ {
                for fi41 := 0; fi41 < 1; fi41++ {
                    v42 := GenS5{F0: 77, F1: 656}
                    _ = fi41
                    _ = v42
                }
                if len(v2) > 0 {
                    v2[0] = 2147483648
                }
                _ = fi40
            }
        case "\n\t":
            if ((0 >= 818) || (93 <= 42)) {
                if false {
                    if false {
                        v43 := GenS1{F0: 991, F1: 100, F2: 128}
                        if (1 < 42) {
                            switch "a" {
                            case "a":
                                var v44 uint16 = 255
                                _ = v44
                            case "foo bar":
                                if false {
                                    if (((false || true) && (104 <= 100)) || false) {
                                        v3 -= 1
                                        var v45 uint = 1
                                        _ = v45
                                    }
                                    var v46 uint32 = 4294967295
                                    _ = v46
                                }
                            case "\n\t":
                                var v47 bool = (!true)
                                _ = v47
                            default:
                                fmt.Println(v38)
                            }
                        }
                        _ = v43
                    }
                    for fi48 := 0; fi48 < 2; fi48++ {
                        switch ("hello" + "\n\t") {
                        case "foo bar":
                            var v49 uint8 = 236
                            _ = v49
                        default:
                            _ = len(v2)
                        }
                        switch "hello" {
                        case "hello":
                            v50 := GenS3{F0: "foo bar", F1: v3}
                            _ = v50
                        case "foo bar":
                            if false {
                                v2 = append(v2, 693)
                                fmt.Println(v38, v3)
                                if false {
                                    var v51 string = "hello"
                                    v51 = "\n\t"
                                    _ = v51
                                } else {
                                    var v52 int16 = -224
                                    v5 += v5
                                    _ = v52
                                }
                            }
                        case "":
                            for fi53 := 0; fi53 < 1; fi53++ {
                                v4 = -32768
                                var v54 float32 = -2.25
                                _ = fi53
                                _ = v54
                            }
                        default:
                            v4 += (v3 / (((^(v3))) | 1))
                        }
                        _ = fi48
                    }
                } else {
                    var v55 float32 = 1.5
                    _ = v55
                }
                for fi56 := 0; fi56 < 10; fi56++ {
                    var v57 uint64 = 0
                    _ = fi56
                    _ = v57
                }
                for fi58 := 0; fi58 < 4; fi58++ {
                    v59 := GenS3{F0: "a", F1: (v3 | 42)}
                    _ = fi58
                    _ = v59
                }
            } else {
                if (-712 <= 0) {
                    var v60 int8 = -13
                    _ = v60
                }
            }
        }
    }
    _ = len(v1)
    if (1e10 == 1e10) {
        if ((1e10 == 0.0) && true) {
            if ((v5 / (v5 * (((float64(0.0) * float64((1.5 - ((float32(v5) - 1.5) * -2.25)))) / 3.14159) / v5))) > v5) {
                if len(v2) > 0 {
                    v2[0] = 2147483647
                }
            } else {
                switch "a" {
                case "\n\t":
                    var v61 string = ((("foo bar" + "hello") + "hello") + "\n\t")
                    _ = v61
                case "a":
                    delete(v1, 2147483647)
                case "hello":
                    fmt.Println(v5)
                }
            }
            fmt.Println(v4, v4, v38)
        } else {
            for fi62 := 0; fi62 < 16; fi62++ {
                if ((0.0 >= 0.0) || true) {
                    v5 = (v5 * v5)
                    if (true && (true || true)) {
                        v3 += (v3 << ((865) & 63))
                        v63 := GenS6{F0: -107, F1: v4}
                        var v64 int32 = 0
                        _ = v63
                        _ = v64
                    }
                } else {
                    if ((v38 | v38) != -9223372036854775808) {
                        if true {
                            for fi65 := 0; fi65 < 18; fi65++ {
                                fmt.Println(fi62, v3)
                                _ = fi65
                            }
                            if (1 > 0) {
                                switch "" {
                                case "":
                                    var v66 uint32 = 65535
                                    _ = v66
                                default:
                                    v2 = append(v2, 65535)
                                }
                                if ("foo bar" <= ("\n\t" + "foo bar")) {
                                    for fi67 := 0; fi67 < 9; fi67++ {
                                        var v68 int = (^(fi62))
                                        v68 += v38
                                        v4 = (v4 & (^(v3)))
                                        _ = fi67
                                        _ = v68
                                    }
                                    v4 += 42
                                } else {
                                    for fi69 := 0; fi69 < 17; fi69++ {
                                        fmt.Println(v38, fi62)
                                        var v70 uint32 = 4294967295
                                        v3 += 32767
                                        _ = fi69
                                        _ = v70
                                    }
                                }
                            }
                        }
                        v38 = fi62
                    } else {
                        var v71 int16 = v4
                        v4 = 32767
                        _ = v71
                    }
                    if true {
                        fmt.Println(v38, v38, fi62)
                        for fi72 := 0; fi72 < 8; fi72++ {
                            if true {
                                fmt.Println(fi62, fi62)
                                fmt.Println(v38, v4)
                                var v73 uint64 = 0
                                _ = v73
                            }
                            var v74 uint32 = 865
                            var v75 uint64 = 42
                            var v76 uint64 = v75
                            var v77 uint32 = v74
                            _ = fi72
                            _ = v74
                            _ = v75
                            _ = v76
                            _ = v77
                        }
                    }
                }
                for fi78 := 0; fi78 < 1; fi78++ {
                    var v79 float64 = ((v5 - (3.14159 * (v5 / (((float64(float32(1.5)) - (0.0 + 1.5)) / float64((-2.25 * 1e10))) / v5)))) * 1e10)
                    v79 += (1.5 - ((v79 - float64(1.5)) / (v5 * 3.14159)))
                    var v80 uint32 = 2147483648
                    _ = fi78
                    _ = v79
                    _ = v80
                }
                if false {
                    var v81 bool = ("\n\t" != "foo bar")
                    _ = v81
                } else {
                    var v82 uint = 703
                    _ = v82
                }
                _ = fi62
            }
            fmt.Println(v4, v4, v3)
        }
        var v83 bool = (!(v5 > 1e10))
        switch v38 {
        case 890:
            v38 += v38
        case -2147483648:
            if v83 {
                if ("foo bar" >= "\n\t") {
                    v83 = (v83 && (!(v83 || (float64(0.1) <= -2.25))))
                    for fi84 := 0; fi84 < 3; fi84++ {
                        for fi85 := 0; fi85 < 2; fi85++ {
                            v86 := GenS7{F0: (42 >= 20), F1: 2, F2: 1}
                            v4 -= v3
                            _ = fi85
                            _ = v86
                        }
                        if (1 < -107) {
                            for fi87 := 0; fi87 < 19; fi87++ {
                                v4 = -733
                                if (!((!(1.5 > 1e10)) && false)) {
                                    v3 -= (v4 - -851)
                                } else {
                                    switch -229 {
                                    case -2147483648:
                                        if (!(!((1e10 / (0.0 + float32((v5 / ((v5 / v5) - ((3.14159 + 0.0) / float64(3.14159))))))) < (((float32(v5) + (1.5 / 1.0)) * float32(v5)) * (1.5 / (0.1 + 1.5)))))) {
                                            v2 = append(v2, 65535)
                                            switch ("\n\t" + "") {
                                            case "foo bar":
                                                v83 = v83
                                            }
                                        }
                                    case 32767:
                                        fmt.Println(v4, v5, v38)
                                    case 42:
                                        v88 := v1[445]
                                        _ = v88
                                    default:
                                        _ = len(v1)
                                    }
                                }
                                _ = fi87
                            }
                        }
                        fmt.Println(v3, fi84, v5)
                        _ = fi84
                    }
                } else {
                    v3 += v3
                    if len(v2) > 0 {
                        v2[0] = 2147483648
                    }
                }
                v89 := GenS8{F0: ((float64((float32(v5) / (3.14159 / ((float32(float64((1e10 - 0.0))) - 0.0) / 1.0)))) - v5) == (-2.25 + v5)), F1: 18446744073709551615}
                var v90 uint8 = 2
                _ = v89
                _ = v90
            }
        case 9223372036854775807:
            switch ("foo bar" + (("a" + ("" + "hello")) + "a")) {
            case "foo bar":
                v4 += -838
            case "a":
                v91 := GenS6{F0: -66, F1: (v3 / (((v3 % ((1) | 1))) | 1))}
                _ = v91
            case "hello":
                switch v38 {
                case 699:
                    switch 0 {
                    case -919:
                        if (!((v38 - (-(v38))) >= 4294967295)) {
                            switch "" {
                            case "\n\t":
                                if v83 {
                                    v92 := GenS6{F0: 0, F1: v3}
                                    _ = v92
                                }
                            case "":
                                v93 := GenS9{F0: ("hello" + "")}
                                _ = v93
                            case "hello":
                                fmt.Println(v83)
                            }
                            var v94 uint8 = 42
                            _ = v94
                        } else {
                            for fi95 := 0; fi95 < 18; fi95++ {
                                switch 681 {
                                case -9223372036854775808:
                                    if (float32((v5 * v5)) < (1.5 + (1.5 / 1e10))) {
                                        var v96 float32 = (-2.25 / (1.5 - (1e10 + -2.25)))
                                        if true {
                                            if (221 != 406) {
                                                for fi97 := 0; fi97 < 8; fi97++ {
                                                    var v98 float64 = (1e10 / (v5 * ((1e10 / ((0.1 + -2.25) - v5)) - float64((-2.25 / v96)))))
                                                    var v99 uint16 = 255
                                                    _ = fi97
                                                    _ = v98
                                                    _ = v99
                                                }
                                                if (!v83) {
                                                    if (false && v83) {
                                                        switch "a" {
                                                        case "":
                                                            fmt.Println(v4, v3, v4)
                                                        case "a":
                                                            switch v38 {
                                                            case 403:
                                                                v100 := GenS10{F0: (v3 + (v4 >> ((2) & 63)))}
                                                                _ = v100
                                                            }
                                                        default:
                                                            for fi101 := 0; fi101 < 16; fi101++ {
                                                                var v102 uint32 = 702
                                                                v103 := GenS4{F0: 0, F1: uint64(v102)}
                                                                fmt.Println(v38, v96)
                                                                _ = fi101
                                                                _ = v102
                                                                _ = v103
                                                            }
                                                        }
                                                        v104 := GenS5{F0: 1, F1: 4294967295}
                                                        switch fi95 {
                                                        case -720:
                                                            if ((-486 > -1) && (0 != 4294967295)) {
                                                                var v105 int16 = -776
                                                                var v106 uint32 = 971
                                                                _ = v105
                                                                _ = v106
                                                            } else {
                                                                for fi107 := 0; fi107 < 6; fi107++ {
                                                                    var v108 float64 = (v5 * (1.5 - v5))
                                                                    fmt.Println(v83, v108, v4)
                                                                    _ = len(v2)
                                                                    v3 = v4
                                                                    _ = fi107
                                                                    _ = v108
                                                                }
                                                                if (-1 < 2147483647) {
                                                                    v5 = v5
                                                                }
                                                            }
                                                        }
                                                        _ = v104
                                                    }
                                                    v109 := GenS11{F0: 592, F1: 255}
                                                    var v110 int16 = 0
                                                    _ = v109
                                                    _ = v110
                                                } else {
                                                    var v111 uint64 = 18446744073709551615
                                                    _ = v111
                                                }
                                            }
                                            var v112 string = ("\n\t" + "")
                                            _ = v112
                                        } else {
                                            if v83 {
                                                for fi113 := 0; fi113 < 15; fi113++ {
                                                    var v114 uint = 333
                                                    _ = fi113
                                                    _ = v114
                                                }
                                                if (((v4 % ((v3) | 1)) == 32767) && (65535 > 172)) {
                                                    var v115 int64 = -9223372036854775808
                                                    _ = v115
                                                }
                                                switch "hello" {
                                                case "":
                                                    if (fi95 < (fi95 | fi95)) {
                                                        for fi116 := 0; fi116 < 8; fi116++ {
                                                            fmt.Println(v38, fi95, v38)
                                                            var v117 int8 = 98
                                                            var v118 uint32 = 2147483648
                                                            _ = fi116
                                                            _ = v117
                                                            _ = v118
                                                        }
                                                    }
                                                case "\n\t":
                                                    switch 0 {
                                                    case 32768:
                                                        v5 = v5
                                                    case -2147483648:
                                                        fmt.Println(v96)
                                                    }
                                                default:
                                                    if false {
                                                        if v83 {
                                                            v5 = 1e10
                                                            var v119 float64 = (v5 + v5)
                                                            _ = v119
                                                        } else {
                                                            fmt.Println(v96)
                                                            fmt.Println(v3)
                                                            var v120 uint32 = 4294967295
                                                            _ = v120
                                                        }
                                                        var v121 uint = 255
                                                        _ = v121
                                                    }
                                                }
                                            }
                                            var v122 uint8 = 42
                                            _ = v122
                                        }
                                        for fi123 := 0; fi123 < 7; fi123++ {
                                            v124 := GenS4{F0: 351, F1: 2}
                                            v125 := GenS4{F0: -2147483648, F1: 959}
                                            _ = fi123
                                            _ = v124
                                            _ = v125
                                        }
                                        _ = v96
                                    }
                                case 32767:
                                    delete(v1, 2147483648)
                                case 127:
                                    v5 -= (((v5 * 1e10) + v5) / (v5 + v5))
                                }
                                v126 := GenS12{F0: 0, F1: 1, F2: 42}
                                _ = fi95
                                _ = v126
                            }
                            if (!true) {
                                v83 = (!((v4 & (v4 | (v4 & v3))) > v3))
                                v83 = v83
                            }
                            switch "hello" {
                            case "\n\t":
                                if v83 {
                                    var v127 int16 = v4
                                    v128 := GenS13{F0: 0, F1: -2147483648}
                                    _ = v127
                                    _ = v128
                                }
                            case "foo bar":
                                v2 = append(v2, 663)
                            }
                        }
                    case 9223372036854775807:
                        var v129 string = ""
                        _ = v129
                    case 32768:
                        for fi130 := 0; fi130 < 20; fi130++ {
                            if (-385 >= 42) {
                                fmt.Println(v3)
                            } else {
                                v2 = append(v2, 1)
                                v83 = ((false || (v5 < -2.25)) || (1 < 2))
                                for fi131 := 0; fi131 < 2; fi131++ {
                                    v132 := GenS3{F0: "\n\t", F1: v4}
                                    delete(v1, 528)
                                    var v133 float32 = 1e10
                                    _ = fi131
                                    _ = v132
                                    _ = v133
                                }
                            }
                            _ = fi130
                        }
                    default:
                        v3 -= (^(v4))
                    }
                case 539:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case 32768:
                    v3 -= v3
                }
            default:
                for fi134 := 0; fi134 < 8; fi134++ {
                    var v135 uint32 = 0
                    _ = fi134
                    _ = v135
                }
            }
        }
        _ = v83
    } else {
        v3 = (v4 / (((v3 / ((32767) | 1))) | 1))
        v136 := v1[42]
        v4 = (v3 % ((v3) | 1))
        _ = v136
    }
    v5 += v5
    if (1 <= 269) {
        var v137 float64 = float64((((-2.25 - (3.14159 + 0.1)) * 1e10) * 1.5))
        if true {
            fmt.Println(v3)
            var v138 uint16 = 65535
            _ = v138
        }
        v5 = (((0.0 / v137) / 1.0) * v137)
        _ = v137
    } else {
        for fi139 := 0; fi139 < 4; fi139++ {
            _ = len(v2)
            v4 -= (v4 << ((1) & 63))
            _ = fi139
        }
        if (-9223372036854775808 != 0) {
            v38 = v38
            if false {
                _ = len(v1)
                if (!(!(2 <= 0))) {
                    v3 = v3
                    var v140 int64 = 516
                    _ = v140
                } else {
                    v5 -= v5
                    v38 += v38
                }
                v141 := GenS13{F0: 515, F1: v38}
                _ = v141
            } else {
                var v142 int16 = -32768
                _ = v142
            }
        }
    }
    v143 := GenS11{F0: 496, F1: 255}
    for fi144 := 0; fi144 < 17; fi144++ {
        v145 := v1[457]
        fmt.Println(v5, v5)
        _ = fi144
        _ = v145
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v143
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
