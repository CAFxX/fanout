package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 int8; F1 uint32; F2 int }
type GenS7 struct { F0 int64; F1 uint16 }
type GenS8 struct { F0 string; F1 uint32 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint32; F1 int32; F2 int16 }
type GenS11 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 int8; F1 int }
type GenS14 struct { F0 int16; F1 uint64 }
type GenS15 struct { F0 int32; F1 int16; F2 uint32 }
type GenS16 struct { F0 uint }
type GenS17 struct { F0 uint64; F1 int8; F2 int64 }

func main() {
    var v1 float32 = 1.5
    v2 := make(map[int]bool)
    var v3 int64 = 118
    v4 := []uint32{535}
    v5 := GenS1{F0: 32767, F1: 2147483647}
    delete(v2, 9223372036854775807)
    if (980 >= 675) {
        v4 = append(v4, 0)
        fmt.Println(v1)
        if len(v4) > 0 {
            v4[0] = 1
        }
    } else {
        for fi6 := 0; fi6 < 10; fi6++ {
            switch (fi6 / ((fi6) | 1)) {
            case -2147483648:
                _ = len(v4)
            default:
                if (true && false) {
                    var v7 uint32 = 2
                    _ = len(v2)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v7
                }
            }
            _ = fi6
        }
    }
    if true {
        v1 -= (((v1 - (0.0 + 3.14159)) / 1.5) * v1)
    }
    if true {
        if (((v3 << ((444) & 63)) >= 1) || true) {
            for fi8 := 0; fi8 < 13; fi8++ {
                _ = len(v4)
                _ = fi8
            }
            if (false || ((true && false) || (42 >= 9223372036854775807))) {
                fmt.Println(v1, v3)
                for fi9 := 0; fi9 < 4; fi9++ {
                    if (true && ((!true) && (v3 != v3))) {
                        v10 := GenS1{F0: fi9, F1: fi9}
                        var v11 bool = false
                        if v11 {
                            v11 = ((1 <= 1) || v11)
                        } else {
                            var v12 int64 = v3
                            for fi13 := 0; fi13 < 7; fi13++ {
                                var v14 int16 = -641
                                v15 := GenS1{F0: 9223372036854775807, F1: fi9}
                                _ = fi13
                                _ = v14
                                _ = v15
                            }
                            for fi16 := 0; fi16 < 13; fi16++ {
                                if (v1 >= v1) {
                                    fmt.Println(v11)
                                } else {
                                    if (true || false) {
                                        var v17 uint32 = 632
                                        var v18 uint32 = v17
                                        _ = v17
                                        _ = v18
                                    } else {
                                        if (!v11) {
                                            v1 = float32(3.14159)
                                            fmt.Println(v3)
                                            var v19 uint16 = 811
                                            v3 = (-(v12))
                                            var v20 uint = 126
                                            _ = v19
                                            _ = v20
                                        } else {
                                            fmt.Println(v12, v3)
                                            switch fi9 {
                                            case -699:
                                                if len(v4) > 0 {
                                                    v4[0] = 544
                                                }
                                            case 0:
                                                _ = len(v4)
                                            }
                                            var v21 float32 = v1
                                            _ = v21
                                        }
                                        if true {
                                            v11 = v11
                                            if ((-338 != 0) || (float64(float32(float64(((((1.5 + v1) / 3.14159) - float32(1.5)) / v1)))) != 0.0)) {
                                                _ = len(v2)
                                                fmt.Println(v12)
                                                v22 := GenS1{F0: (fi9 | 127), F1: fi9}
                                                v11 = (!v11)
                                                _ = v22
                                            }
                                            var v23 string = "hello"
                                            _ = v23
                                        }
                                        var v24 bool = v11
                                        _ = v24
                                    }
                                    v25 := GenS1{F0: (-(fi16)), F1: (fi16 * (fi9 & fi9))}
                                    _ = v25
                                }
                                switch ("\n\t" + ((("foo bar" + "a") + "foo bar") + (("hello" + "a") + "\n\t"))) {
                                case "a":
                                    fmt.Println(v12, fi9, v1)
                                case "foo bar":
                                    var v26 uint64 = 18446744073709551615
                                    _ = v26
                                }
                                var v27 int8 = 1
                                _ = fi16
                                _ = v27
                            }
                            _ = v12
                        }
                        _ = v10
                        _ = v11
                    }
                    _ = fi9
                }
                if true {
                    v28 := GenS2{F0: -665, F1: 32767}
                    switch "hello" {
                    case "":
                        v1 += (0.1 - 0.0)
                    case "a":
                        var v29 int = 128
                        _ = v29
                    }
                    for fi30 := 0; fi30 < 11; fi30++ {
                        var v31 uint16 = 606
                        _ = fi30
                        _ = v31
                    }
                    _ = v28
                } else {
                    if ((true || true) || (true || (!true))) {
                        v1 += (-2.25 + v1)
                        v3 += v3
                        for fi32 := 0; fi32 < 10; fi32++ {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v3 += v3
                            v33 := v2[(fi32 / (((^(fi32))) | 1))]
                            _ = fi32
                            _ = v33
                        }
                    }
                    var v34 uint = 0
                    var v35 uint = v34
                    _ = v34
                    _ = v35
                }
            }
            var v36 int32 = 2147483647
            _ = v36
        }
        switch "foo bar" {
        case "hello":
            if (-449 == -2147483648) {
                if (-1 <= -128) {
                    for fi37 := 0; fi37 < 20; fi37++ {
                        v38 := GenS2{F0: fi37, F1: 2}
                        _ = fi37
                        _ = v38
                    }
                    for fi39 := 0; fi39 < 13; fi39++ {
                        if true {
                            switch "hello" {
                            case "a":
                                for fi40 := 0; fi40 < 2; fi40++ {
                                    v3 -= 21
                                    _ = fi40
                                }
                            case "foo bar":
                                var v41 uint64 = 0
                                _ = v41
                            }
                            v3 = v3
                            fmt.Println(v3, v1, v3)
                        }
                        v3 += v3
                        if len(v4) > 0 {
                            v4[0] = 1
                        }
                        _ = fi39
                    }
                    if (100 >= 100) {
                        v1 += (v1 / float32(3.14159))
                        v2[-1] = (357 == 42)
                        v3 = v3
                    } else {
                        v1 = 0.0
                        if false {
                            switch 181 {
                            case -393:
                                delete(v2, 128)
                            case 42:
                                v42 := GenS1{F0: -9223372036854775808, F1: 956}
                                _ = v42
                            case -988:
                                if true {
                                    _ = len(v4)
                                    var v43 uint8 = 127
                                    _ = v43
                                } else {
                                    var v44 bool = false
                                    _ = v44
                                }
                            }
                            v45 := GenS3{F0: "a", F1: 0}
                            v2[268] = ((0.1 / 1e10) >= 1e10)
                            _ = v45
                        } else {
                            delete(v2, 42)
                        }
                        if false {
                            v3 += v3
                            v46 := GenS1{F0: -1, F1: 251}
                            _ = v46
                        } else {
                            delete(v2, -429)
                            switch -2147483648 {
                            case 422:
                                for fi47 := 0; fi47 < 18; fi47++ {
                                    switch "\n\t" {
                                    case "\n\t":
                                        if len(v4) > 0 {
                                            v4[0] = 2147483647
                                        }
                                    case "":
                                        switch (fi47 + (^(fi47))) {
                                        case 9223372036854775807:
                                            var v48 uint16 = 2
                                            _ = v48
                                        case 693:
                                            for fi49 := 0; fi49 < 3; fi49++ {
                                                delete(v2, (fi49 % ((fi49) | 1)))
                                                _ = fi49
                                            }
                                        case 564:
                                            fmt.Println(v1)
                                        default:
                                            var v50 uint8 = 42
                                            _ = v50
                                        }
                                    case "a":
                                        v4 = append(v4, 42)
                                    default:
                                        if (!(!true)) {
                                            v51 := GenS4{F0: 749}
                                            _ = v51
                                        } else {
                                            fmt.Println(v3)
                                            v1 += ((1.5 + -2.25) - -2.25)
                                            for fi52 := 0; fi52 < 13; fi52++ {
                                                v53 := GenS5{F0: true, F1: 82}
                                                var v54 int64 = v3
                                                var v55 string = ("\n\t" + "a")
                                                _ = fi52
                                                _ = v53
                                                _ = v54
                                                _ = v55
                                            }
                                        }
                                    }
                                    if (true || (!(812 <= v3))) {
                                        var v56 uint32 = 442
                                        v57 := GenS6{F0: 100, F1: (v56 % ((v56) | 1)), F2: fi47}
                                        _ = v56
                                        _ = v57
                                    }
                                    _ = fi47
                                }
                            case 4294967295:
                                v1 -= (v1 + (-2.25 * 1.5))
                            case 907:
                                v1 += float32(-2.25)
                            }
                            var v58 float32 = -2.25
                            _ = v58
                        }
                    }
                } else {
                    var v59 int32 = 1
                    var v60 int8 = 127
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v59
                    _ = v60
                }
            }
        case "":
            if len(v4) > 0 {
                v4[0] = 42
            }
        }
    }
    v61 := GenS4{F0: 2}
    for fi62 := 0; fi62 < 10; fi62++ {
        for fi63 := 0; fi63 < 7; fi63++ {
            if (42 != -1) {
                if len(v4) > 0 {
                    v4[0] = 631
                }
            }
            _ = fi63
        }
        if false {
            v3 += v3
        } else {
            v64 := GenS2{F0: fi62, F1: 1}
            fmt.Println(v3)
            fmt.Println(v1, fi62, v3)
            _ = v64
        }
        _ = fi62
    }
    for fi65 := 0; fi65 < 17; fi65++ {
        v66 := GenS4{F0: 42}
        for fi67 := 0; fi67 < 4; fi67++ {
            v1 += 0.0
            switch fi67 {
            case 0:
                switch ("foo bar" + ("" + "foo bar")) {
                case "foo bar":
                    var v68 uint8 = 1
                    _ = v68
                case "":
                    v69 := GenS4{F0: 264}
                    _ = v69
                case "\n\t":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            case 32767:
                v3 = int64(fi65)
            case -616:
                var v70 uint8 = 0
                _ = v70
            }
            var v71 int64 = (v3 >> ((1) & 63))
            _ = fi67
            _ = v71
        }
        _ = fi65
        _ = v66
    }
    fmt.Println(v3, v1)
    if true {
        if (v3 < 971) {
            v72 := GenS7{F0: 42, F1: 2}
            for fi73 := 0; fi73 < 12; fi73++ {
                if true {
                    v1 = 0.0
                    var v74 uint16 = 32767
                    v3 += (v3 % ((v3) | 1))
                    _ = v74
                } else {
                    v1 = 0.0
                    var v75 int = (fi73 * 85)
                    v76 := GenS2{F0: 32768, F1: 651}
                    _ = v75
                    _ = v76
                }
                _ = len(v2)
                _ = fi73
            }
            _ = v72
        } else {
            var v77 int64 = (v3 << ((18446744073709551615) & 63))
            v78 := GenS4{F0: 65535}
            if true {
                fmt.Println(v3, v3, v1)
            }
            _ = v77
            _ = v78
        }
        if (4294967295 == 2) {
            v79 := GenS8{F0: "", F1: 34}
            switch (("\n\t" + (((("\n\t" + "a") + "hello") + "") + ("hello" + "\n\t"))) + ("a" + ("hello" + "\n\t"))) {
            case "\n\t":
                var v80 int8 = 98
                _ = v80
            default:
                var v81 int64 = (v3 | v3)
                _ = v81
            }
            if ((v3 % (((v3 / (((v3 / ((-713) | 1))) | 1))) | 1)) > -9223372036854775808) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v82 float32 = v1
                if true {
                    fmt.Println(v3, v1, v3)
                    if true {
                        fmt.Println(v82)
                        if ((72 == 255) && ((false && (0 < 0)) && (171 >= 65535))) {
                            for fi83 := 0; fi83 < 18; fi83++ {
                                v82 = 1.5
                                _ = len(v4)
                                if (!false) {
                                    v84 := GenS9{F0: 545}
                                    _ = v84
                                }
                                _ = fi83
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        } else {
                            for fi85 := 0; fi85 < 3; fi85++ {
                                if true {
                                    v86 := GenS4{F0: 712}
                                    v87 := GenS5{F0: false, F1: 647}
                                    v88 := GenS2{F0: int(v3), F1: 2}
                                    _ = v86
                                    _ = v87
                                    _ = v88
                                } else {
                                    v82 -= v82
                                    var v89 bool = (true && false)
                                    for fi90 := 0; fi90 < 3; fi90++ {
                                        v91 := GenS2{F0: fi90, F1: 42}
                                        v92 := GenS5{F0: v89, F1: 255}
                                        var v93 float64 = 1e10
                                        _ = fi90
                                        _ = v91
                                        _ = v92
                                        _ = v93
                                    }
                                    _ = v89
                                }
                                switch fi85 {
                                case 2147483647:
                                    var v94 int8 = 58
                                    _ = v94
                                }
                                _ = fi85
                            }
                        }
                    }
                }
                _ = v82
            }
            _ = v79
        } else {
            _ = len(v2)
        }
        if (42 < 127) {
            v1 = v1
            v1 += 1e10
            v3 -= -328
        } else {
            switch 0 {
            case -778:
                switch 485 {
                case 32768:
                    v1 += v1
                case 2147483647:
                    if false {
                        switch 42 {
                        case 2147483647:
                            switch "hello" {
                            case "\n\t":
                                fmt.Println(v3, v1)
                            default:
                                switch "foo bar" {
                                case "hello":
                                    if true {
                                        if (!(("a" + "foo bar") > "hello")) {
                                            if len(v4) > 0 {
                                                v4[0] = 0
                                            }
                                            v95 := GenS10{F0: 290, F1: -1, F2: 444}
                                            if false {
                                                v96 := GenS8{F0: ("a" + ""), F1: 2147483647}
                                                _ = v96
                                            }
                                            _ = v95
                                        }
                                    } else {
                                        var v97 bool = ((!true) || (!false))
                                        v3 = (-(v3))
                                        _ = v97
                                    }
                                case "a":
                                    fmt.Println(v1, v3)
                                }
                            }
                        }
                        v1 += (v1 - v1)
                        var v98 uint = 1
                        _ = v98
                    }
                default:
                    var v99 int64 = v3
                    _ = v99
                }
            case 4294967295:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case 128:
                if true {
                    _ = len(v2)
                    fmt.Println(v3, v3)
                    for fi100 := 0; fi100 < 19; fi100++ {
                        v101 := GenS11{F0: 4294967295, F1: v3, F2: 0}
                        if false {
                            for fi102 := 0; fi102 < 12; fi102++ {
                                _ = len(v2)
                                delete(v2, 715)
                                var v103 uint32 = 4294967295
                                _ = fi102
                                _ = v103
                            }
                            for fi104 := 0; fi104 < 10; fi104++ {
                                v3 -= -632
                                v1 -= ((0.0 + v1) + (3.14159 + 0.0))
                                v105 := GenS9{F0: 2}
                                _ = fi104
                                _ = v105
                            }
                            var v106 uint = 851
                            _ = v106
                        } else {
                            if false {
                                v107 := GenS12{F0: 32767}
                                if (!(0 != 18446744073709551615)) {
                                    var v108 int16 = int16(fi100)
                                    if (42 == 174) {
                                        v108 = v108
                                        v108 -= v108
                                    } else {
                                        v109 := GenS4{F0: 32767}
                                        v110 := GenS3{F0: "\n\t", F1: 0}
                                        fmt.Println(v1)
                                        _ = v109
                                        _ = v110
                                    }
                                    if (-2.25 > (1.5 * 1e10)) {
                                        fmt.Println(v1)
                                        if (true || (42 < 42)) {
                                            v3 += int64(v108)
                                        } else {
                                            fmt.Println(v1, v1)
                                            v111 := GenS13{F0: 1, F1: (fi100 & (fi100 / ((fi100) | 1)))}
                                            _ = v111
                                        }
                                        v108 = v108
                                    }
                                    _ = v108
                                }
                                _ = v107
                            } else {
                                v112 := GenS14{F0: 1000, F1: 0}
                                v4 = append(v4, 2147483647)
                                switch fi100 {
                                case 262:
                                    for fi113 := 0; fi113 < 18; fi113++ {
                                        fmt.Println(v1)
                                        fmt.Println(fi113, v1, fi100)
                                        var v114 uint32 = 2147483648
                                        _ = fi113
                                        _ = v114
                                    }
                                }
                                _ = v112
                            }
                            switch (("hello" + (("foo bar" + ("hello" + "hello")) + ("a" + ""))) + "hello") {
                            case "":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            default:
                                var v115 uint16 = 42
                                _ = v115
                            }
                        }
                        if (473 >= 18446744073709551615) {
                            var v116 int32 = 0
                            _ = v116
                        }
                        _ = fi100
                        _ = v101
                    }
                }
            default:
                v1 += 0.0
            }
            for fi117 := 0; fi117 < 4; fi117++ {
                fmt.Println(v1)
                v4 = append(v4, 226)
                _ = fi117
            }
            for fi118 := 0; fi118 < 8; fi118++ {
                _ = len(v2)
                _ = fi118
            }
        }
    }
    if true {
        var v119 int = -659
        switch v119 {
        case 417:
            if (("\n\t" == "") && true) {
                v119 = v119
                if true {
                    v1 += v1
                } else {
                    v2[96] = false
                    v3 = (v3 ^ -9223372036854775808)
                    if (472 < 2147483648) {
                        var v120 int32 = 1
                        _ = v120
                    } else {
                        for fi121 := 0; fi121 < 12; fi121++ {
                            switch "a" {
                            case "a":
                                v122 := v2[(fi121 % ((v119) | 1))]
                                _ = v122
                            }
                            fmt.Println(v119, v119, v3)
                            v3 = (-(v3))
                            _ = fi121
                        }
                    }
                }
            }
        case 4294967295:
            if ((65535 != 0) && true) {
                v123 := GenS7{F0: (v3 - v3), F1: 911}
                if true {
                    v119 += (v119 >> ((584) & 63))
                }
                _ = v123
            }
        case 2147483647:
            switch "\n\t" {
            case "a":
                if true {
                    if (50 > 729) {
                        for fi124 := 0; fi124 < 6; fi124++ {
                            var v125 int32 = 2147483647
                            _ = fi124
                            _ = v125
                        }
                    } else {
                        for fi126 := 0; fi126 < 15; fi126++ {
                            if len(v4) > 0 {
                                v4[0] = 981
                            }
                            var v127 bool = true
                            _ = fi126
                            _ = v127
                        }
                        for fi128 := 0; fi128 < 13; fi128++ {
                            fmt.Println(v119)
                            v119 -= (fi128 ^ 9223372036854775807)
                            switch 350 {
                            case 128:
                                fmt.Println(v3, v1)
                            default:
                                v129 := GenS11{F0: 2147483648, F1: v3, F2: 18446744073709551615}
                                _ = v129
                            }
                            _ = fi128
                        }
                    }
                } else {
                    fmt.Println(v3)
                    v1 = 1e10
                }
            case "foo bar":
                if true {
                    if len(v4) > 0 {
                        v4[0] = 161
                    }
                    if false {
                        var v130 bool = false
                        for fi131 := 0; fi131 < 17; fi131++ {
                            fmt.Println(v119, v119, v130)
                            var v132 int = -739
                            _ = fi131
                            _ = v132
                        }
                        _ = v130
                    } else {
                        v119 = (v119 / ((v119) | 1))
                    }
                    switch (v119 >> ((65535) & 63)) {
                    case 32768:
                        fmt.Println(v1, v1)
                    case -611:
                        if (!false) {
                            for fi133 := 0; fi133 < 16; fi133++ {
                                fmt.Println(fi133, v1, v119)
                                _ = fi133
                            }
                            if (160 >= 2) {
                                fmt.Println(v119, v1, v119)
                            }
                            if true {
                                v119 = (v119 * v119)
                                if (false || (62 != 4294967295)) {
                                    for fi134 := 0; fi134 < 2; fi134++ {
                                        v3 = v3
                                        v135 := GenS15{F0: -1, F1: 32767, F2: 0}
                                        _ = fi134
                                        _ = v135
                                    }
                                    v1 -= float32(float64((v1 - 1e10)))
                                } else {
                                    switch "a" {
                                    case "hello":
                                        if false {
                                            var v136 int8 = 42
                                            v136 = v136
                                            _ = v136
                                        } else {
                                            if len(v4) > 0 {
                                                v4[0] = 0
                                            }
                                            v137 := GenS15{F0: -1, F1: 42, F2: 4294967295}
                                            _ = v137
                                        }
                                    }
                                }
                                var v138 int32 = 49
                                _ = v138
                            } else {
                                v2[(v119 ^ v119)] = (v1 <= (v1 / ((v1 / ((float32((1.5 * 1e10)) + ((v1 - (v1 / -2.25)) / (float32(0.1) * 3.14159))) * v1)) / v1)))
                                if false {
                                    v139 := GenS14{F0: -930, F1: 2}
                                    v1 -= v1
                                    _ = v139
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        } else {
                            v140 := GenS16{F0: 0}
                            v1 = (float32((3.14159 + (float64((v1 * v1)) - 1.5))) - ((v1 - float32(1e10)) / (0.0 / v1)))
                            if false {
                                if false {
                                    for fi141 := 0; fi141 < 13; fi141++ {
                                        var v142 int = 1
                                        fmt.Println(v3, v119)
                                        v142 += -1
                                        _ = fi141
                                        _ = v142
                                    }
                                    var v143 float32 = (v1 * (v1 - v1))
                                    v144 := GenS17{F0: 4294967295, F1: -128, F2: v3}
                                    _ = v143
                                    _ = v144
                                }
                            } else {
                                fmt.Println(v3, v1, v119)
                                for fi145 := 0; fi145 < 3; fi145++ {
                                    v3 -= v3
                                    if (-2147483648 == 921) {
                                        fmt.Println(fi145)
                                    } else {
                                        v3 = -963
                                    }
                                    var v146 uint16 = 32767
                                    _ = fi145
                                    _ = v146
                                }
                                var v147 int8 = -96
                                _ = v147
                            }
                            _ = v140
                        }
                    }
                } else {
                    switch (^(v119)) {
                    case 32767:
                        _ = len(v4)
                    case 4294967295:
                        v119 -= (v119 % ((572) | 1))
                    }
                    v119 += 127
                }
            default:
                var v148 int64 = (v3 >> ((179) & 63))
                _ = v148
            }
        }
        for fi149 := 0; fi149 < 19; fi149++ {
            fmt.Println(v119)
            fmt.Println(v1)
            var v150 uint64 = 0
            _ = fi149
            _ = v150
        }
        _ = v119
    } else {
        var v151 string = "hello"
        if false {
            for fi152 := 0; fi152 < 5; fi152++ {
                v1 += ((3.14159 * 1.5) + -2.25)
                v1 = ((v1 / (v1 * v1)) / v1)
                for fi153 := 0; fi153 < 9; fi153++ {
                    v1 = v1
                    if (1 == 128) {
                        fmt.Println(v1, v151, fi153)
                        if (128 >= 127) {
                            var v154 bool = true
                            _ = v154
                        } else {
                            v155 := GenS4{F0: 0}
                            switch v151 {
                            case "":
                                var v156 int = 2147483647
                                _ = v156
                            case "hello":
                                var v157 uint8 = 42
                                _ = v157
                            case "\n\t":
                                if true {
                                    var v158 uint8 = 220
                                    _ = v158
                                }
                            }
                            switch fi153 {
                            case -949:
                                v3 = v3
                            case 0:
                                if len(v4) > 0 {
                                    v4[0] = 938
                                }
                            case -671:
                                var v159 int16 = -83
                                _ = v159
                            }
                            _ = v155
                        }
                        if (1e10 >= v1) {
                            v160 := GenS5{F0: (false && false), F1: 18446744073709551615}
                            if false {
                                fmt.Println(v151, v1, fi153)
                                v161 := v2[(-(fi152))]
                                v161 = false
                                _ = v161
                            }
                            _ = v160
                        } else {
                            switch 32767 {
                            case 9223372036854775807:
                                delete(v2, fi153)
                            case -414:
                                var v162 float32 = -2.25
                                _ = v162
                            case -2147483648:
                                var v163 int8 = 100
                                _ = v163
                            default:
                                v151 = (((v151 + "a") + v151) + ("hello" + "a"))
                            }
                        }
                    } else {
                        if (v1 >= ((((3.14159 / 1e10) - float32(0.0)) * (3.14159 - v1)) / 1e10)) {
                            v3 = v3
                        }
                    }
                    _ = fi153
                }
                _ = fi152
            }
            for fi164 := 0; fi164 < 4; fi164++ {
                var v165 int32 = -1
                var v166 string = v151
                _ = fi164
                _ = v165
                _ = v166
            }
        }
        _ = v151
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v4)
    var v167 string = "\n\t"
    v1 = float32(1.5)
    switch (((v167 + v167) + "a") + v167) {
    case "\n\t":
        var v168 int16 = 42
        _ = v168
    }
    var v169 string = v167
    fmt.Println("v1:", v1)
    fmt.Println("v167:", v167)
    fmt.Println("v169:", v169)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v167
    _ = v169
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v61
}
