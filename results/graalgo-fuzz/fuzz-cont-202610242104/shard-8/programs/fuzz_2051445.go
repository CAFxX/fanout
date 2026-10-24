package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int; F1 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16; F1 int16; F2 int8 }
type GenS7 struct { F0 uint64; F1 uint }
type GenS8 struct { F0 uint64; F1 bool }
type GenS9 struct { F0 uint; F1 int; F2 string }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint8; F1 int32 }
type GenS13 struct { F0 int8; F1 int64 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 int8; F1 int }
type GenS17 struct { F0 int8 }
type GenS18 struct { F0 int; F1 uint }
type GenS19 struct { F0 uint; F1 bool }
type GenS20 struct { F0 int64; F1 uint32 }
type GenS21 struct { F0 int64; F1 uint32 }
type GenS22 struct { F0 int8 }
type GenS23 struct { F0 uint; F1 bool }
type GenS24 struct { F0 int8 }
type GenS25 struct { F0 uint64 }
func genFn0() uint {
    if (-2147483648 < -724) {
        var v7 string = "a"
        var v8 uint64 = 244
        var v9 int32 = 0
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 uint64 = 4294967295
    _ = v10
    return 18446744073709551615
}
func genFn1() uint8 {
    switch -1 {
    case -9223372036854775808:
        var v11 uint32 = 531
        _ = v11
    }
    switch "" {
    case "hello":
        v12 := GenS1{F0: 1}
        _ = v12
    case "foo bar":
        var v13 int32 = 0
        _ = v13
    case "\n\t":
        if (("foo bar" + ("\n\t" + ("foo bar" + (("a" + ("\n\t" + ("\n\t" + "\n\t"))) + "")))) <= "foo bar") {
            var v14 bool = false
            switch ("hello" + "\n\t") {
            case "hello":
                if (34 != -62) {
                    var v15 string = (("a" + "") + "\n\t")
                    _ = v15
                }
            default:
                v16 := GenS2{F0: 0}
                _ = v16
            }
            _ = v14
        }
    }
    var v17 uint16 = 152
    _ = v17
    return 128
}

func main() {
    v1 := []uint32{239}
    v2 := []int{2147483647, 32767, 774, 127}
    var v3 int8 = 100
    var v4 int64 = 42
    var v5 uint64 = 1
    v6 := make(map[int]bool)
    if false {
        v4 = (^(v4))
        var v18 bool = false
        _ = len(v1)
        _ = v18
    } else {
        var v19 float64 = float64(3.14159)
        var v20 int32 = -1
        _ = v19
        _ = v20
    }
    var v21 float32 = 1e10
    if true {
        switch 32768 {
        case 2147483647:
            v22 := GenS2{F0: 0}
            _ = v22
        case 421:
            if ("foo bar" < ("\n\t" + (("\n\t" + "hello") + "a"))) {
                var v23 uint64 = 715
                fmt.Println(v3, v21, v5)
                delete(v6, 128)
                _ = v23
            } else {
                if ((v5 >= (v5 % ((42) | 1))) && false) {
                    for fi24 := 0; fi24 < 11; fi24++ {
                        if ((v5 <= 323) || (false && true)) {
                            var v25 float32 = (v21 + (3.14159 - (v21 * v21)))
                            _ = v25
                        } else {
                            _ = len(v6)
                            v3 += (v3 << ((genFn0()) & 63))
                        }
                        _ = fi24
                    }
                    for fi26 := 0; fi26 < 15; fi26++ {
                        for fi27 := 0; fi27 < 15; fi27++ {
                            var v28 uint64 = (v5 - (v5 / ((2) | 1)))
                            v3 += v3
                            v3 = (v3 | v3)
                            v29 := GenS3{F0: (v3 / ((-1) | 1))}
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                        _ = fi26
                    }
                    v21 = v21
                } else {
                    var v30 uint = 2
                    var v31 int64 = v4
                    _ = v30
                    _ = v31
                }
                for fi32 := 0; fi32 < 18; fi32++ {
                    if (!true) {
                        v21 = (float32(float64((((v21 * (((1e10 - 1e10) * float32(-2.25)) + -2.25)) * 1.5) - (3.14159 + 1.5)))) * v21)
                        if (v3 == v3) {
                            v5 -= (v5 & 42)
                            _ = len(v6)
                        }
                    }
                    _ = fi32
                }
                v21 = float32(-2.25)
            }
        default:
            if ((^(v3)) != (v3 ^ int8(v5))) {
                var v33 int64 = v4
                _ = v33
            }
        }
        v34 := GenS2{F0: 408}
        fmt.Println(v3, v5, v4)
        _ = v34
    }
    v35 := GenS4{F0: 127, F1: v4}
    if (true || (1 != -2147483648)) {
        v36 := GenS2{F0: 65535}
        switch 32768 {
        case 128:
            if false {
                switch 127 {
                case 32767:
                    v3 -= (v3 << ((510) & 63))
                case 128:
                    v37 := GenS5{F0: 2147483647}
                    _ = v37
                case 115:
                    var v38 uint16 = 32767
                    _ = v38
                }
            }
        case 804:
            var v39 float32 = 1.5
            _ = v39
        case 32768:
            if true {
                var v40 bool = ((0 != 93) || true)
                if v40 {
                    if (2147483647 >= -2147483648) {
                        v5 = v5
                        if (false && ((-2.25 - 1e10) >= 0.0)) {
                            fmt.Println(v4, v4, v40)
                            for fi41 := 0; fi41 < 18; fi41++ {
                                for fi42 := 0; fi42 < 17; fi42++ {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = fi42
                                }
                                v2 = append(v2, (^(fi41)))
                                v40 = ((v4 % ((v4) | 1)) <= (v4 >> ((18446744073709551615) & 63)))
                                _ = fi41
                            }
                            switch ((("hello" + ((("a" + "a") + "foo bar") + ("foo bar" + "\n\t"))) + "a") + "") {
                            case "\n\t":
                                for fi43 := 0; fi43 < 6; fi43++ {
                                    var v44 int8 = v3
                                    _ = fi43
                                    _ = v44
                                }
                            case "a":
                                for fi45 := 0; fi45 < 4; fi45++ {
                                    if v40 {
                                        var v46 float64 = 1.5
                                        v3 -= (v3 / (((v3 ^ -1)) | 1))
                                        _ = v46
                                    } else {
                                        v47 := GenS2{F0: 1}
                                        var v48 uint = 4294967295
                                        v1 = append(v1, 65535)
                                        _ = v47
                                        _ = v48
                                    }
                                    _ = fi45
                                }
                            case "hello":
                                var v49 uint32 = 2
                                _ = v49
                            }
                        } else {
                            var v50 uint8 = 64
                            _ = v50
                        }
                        v3 += (v3 << ((360) & 63))
                    } else {
                        switch ("foo bar" + ("foo bar" + (("\n\t" + "hello") + "hello"))) {
                        case "":
                            if len(v2) > 0 {
                                v2[0] = -30
                            }
                        case "foo bar":
                            if v40 {
                                var v51 int = -9223372036854775808
                                for fi52 := 0; fi52 < 7; fi52++ {
                                    v3 += (v3 | v3)
                                    if v40 {
                                        v6[(v51 / (((fi52 | v51)) | 1))] = v40
                                        var v53 uint8 = 98
                                        if (!(!(52 <= 42))) {
                                            var v54 string = "foo bar"
                                            v51 = (^(fi52))
                                            var v55 uint8 = (v53 * v53)
                                            _ = v54
                                            _ = v55
                                        }
                                        _ = v53
                                    }
                                    switch (v51 * (fi52 ^ v51)) {
                                    case -9223372036854775808:
                                        for fi56 := 0; fi56 < 2; fi56++ {
                                            v57 := GenS3{F0: v3}
                                            v58 := GenS2{F0: 2147483647}
                                            v1 = append(v1, 42)
                                            fmt.Println(fi56)
                                            _ = fi56
                                            _ = v57
                                            _ = v58
                                        }
                                    case -536:
                                        if v40 {
                                            var v59 uint64 = v5
                                            _ = v59
                                        } else {
                                            v60 := GenS6{F0: 2, F1: 42, F2: v3}
                                            v4 += (v4 << ((0) & 63))
                                            _ = v60
                                        }
                                    case 2147483647:
                                        v6[fi52] = false
                                    default:
                                        for fi61 := 0; fi61 < 11; fi61++ {
                                            _ = len(v6)
                                            _ = fi61
                                        }
                                    }
                                    _ = fi52
                                }
                                _ = v51
                            } else {
                                v62 := GenS4{F0: 128, F1: (v4 + (-(v4)))}
                                _ = v62
                            }
                        }
                    }
                    switch "\n\t" {
                    case "\n\t":
                        if v40 {
                            var v63 float64 = -2.25
                            fmt.Println(v21)
                            _ = v63
                        }
                    }
                    var v64 int8 = (v3 | v3)
                    _ = v64
                } else {
                    if true {
                        for fi65 := 0; fi65 < 15; fi65++ {
                            switch fi65 {
                            case 1:
                                v66 := GenS7{F0: v5, F1: 1}
                                _ = v66
                            case 4294967295:
                                delete(v6, 2147483647)
                            case -9223372036854775808:
                                v5 += v5
                            default:
                                v21 = (-2.25 * (3.14159 * v21))
                            }
                            fmt.Println(v21)
                            for fi67 := 0; fi67 < 3; fi67++ {
                                v68 := GenS8{F0: (^(v5)), F1: true}
                                v3 += v3
                                _ = fi67
                                _ = v68
                            }
                            _ = fi65
                        }
                    }
                    v69 := GenS9{F0: 255, F1: 128, F2: (("hello" + "foo bar") + "foo bar")}
                    v21 = -2.25
                    _ = v69
                }
                if true {
                    switch "\n\t" {
                    case "\n\t":
                        v70 := GenS10{F0: 1}
                        _ = v70
                    case "a":
                        var v71 uint32 = 375
                        _ = v71
                    case "":
                        fmt.Println(v4, v4, v4)
                    default:
                        switch 614 {
                        case 32768:
                            switch "hello" {
                            case "a":
                                v4 = v4
                            case "foo bar":
                                for fi72 := 0; fi72 < 16; fi72++ {
                                    var v73 int = 42
                                    _ = fi72
                                    _ = v73
                                }
                            case "\n\t":
                                v74 := GenS1{F0: 0}
                                _ = v74
                            }
                        }
                    }
                    var v75 int64 = (v4 & (v4 ^ (v4 - v4)))
                    _ = v75
                } else {
                    for fi76 := 0; fi76 < 16; fi76++ {
                        var v77 int32 = 1
                        for fi78 := 0; fi78 < 11; fi78++ {
                            _ = len(v1)
                            switch 127 {
                            case 602:
                                v77 += v77
                            }
                            v21 = (0.1 - v21)
                            _ = fi78
                        }
                        switch fi76 {
                        case -577:
                            v79 := GenS7{F0: 531, F1: 0}
                            _ = v79
                        }
                        _ = fi76
                        _ = v77
                    }
                    var v80 int16 = 0
                    _ = v80
                }
                _ = v40
            } else {
                var v81 uint16 = 87
                var v82 uint32 = 2147483647
                fmt.Println(v21)
                _ = v81
                _ = v82
            }
        }
        v6[42] = false
        _ = v36
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v5)
    for fi83 := 0; fi83 < 14; fi83++ {
        for fi84 := 0; fi84 < 15; fi84++ {
            v4 -= 1
            if (18446744073709551615 > genFn0()) {
                switch "\n\t" {
                case "foo bar":
                    v21 = (1e10 * (float32(0.1) * 1.5))
                case "hello":
                    var v85 int8 = (v3 << ((2) & 63))
                    _ = v85
                default:
                    switch ("hello" + ("" + "hello")) {
                    case "a":
                        if false {
                            v86 := GenS11{F0: 142}
                            _ = v86
                        } else {
                            var v87 uint8 = 128
                            _ = v87
                        }
                    case "hello":
                        if len(v1) > 0 {
                            v1[0] = 65535
                        }
                    }
                }
            }
            if false {
                v21 = (v21 / v21)
                v5 = v5
                v3 = (v3 % (((^(v3))) | 1))
            }
            _ = fi84
        }
        for fi88 := 0; fi88 < 14; fi88++ {
            switch ("foo bar" + (("\n\t" + (("" + "foo bar") + "")) + ((("hello" + "a") + "foo bar") + ((("foo bar" + ("" + "hello")) + "") + ("\n\t" + "\n\t"))))) {
            case "":
                v89 := GenS12{F0: 1, F1: 2147483647}
                _ = v89
            case "\n\t":
                if (false && (1e10 <= (1e10 + 3.14159))) {
                    var v90 uint16 = 160
                    switch "hello" {
                    case "a":
                        var v91 uint16 = v90
                        _ = v91
                    case "foo bar":
                        v92 := GenS9{F0: 408, F1: (fi88 / ((fi83) | 1)), F2: "foo bar"}
                        _ = v92
                    case "\n\t":
                        v21 -= v21
                    }
                    _ = v90
                } else {
                    if (v3 < v3) {
                        if ((v4 & (v4 | -350)) == -1) {
                            switch "\n\t" {
                            case "foo bar":
                                v21 -= (-2.25 - v21)
                            case "hello":
                                fmt.Println(fi88)
                            case "a":
                                fmt.Println(fi83)
                                if (!(1 <= int32(v4))) {
                                    var v93 int = (fi88 | -2147483648)
                                    _ = v93
                                } else {
                                    switch "a" {
                                    case "foo bar":
                                        switch fi83 {
                                        case -9223372036854775808:
                                            var v94 uint8 = 198
                                            _ = v94
                                        }
                                    case "hello":
                                        _ = len(v6)
                                    default:
                                        v95 := v6[217]
                                        _ = v95
                                    }
                                }
                            }
                            v21 += (float32(float64(v21)) + (((3.14159 - v21) * (v21 - v21)) / 1e10))
                            v5 += (v5 >> ((645) & 63))
                        } else {
                            fmt.Println(v3)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v96 uint64 = 259
                            _ = v96
                        }
                    } else {
                        v21 += (v21 - (0.0 - 0.1))
                        fmt.Println(v5)
                    }
                    v97 := GenS13{F0: (v3 * (-(v3))), F1: v4}
                    if false {
                        switch 9223372036854775807 {
                        case 128:
                            v3 = v3
                        case 4294967295:
                            v4 -= -1
                        default:
                            delete(v6, fi88)
                        }
                        switch (^(fi88)) {
                        case 128:
                            v98 := GenS14{F0: uint16(v5)}
                            _ = v98
                        case 348:
                            var v99 int8 = v3
                            _ = v99
                        case -9223372036854775808:
                            if true {
                                v100 := GenS15{F0: 32767}
                                v5 = 291
                                var v101 int16 = -1
                                _ = v100
                                _ = v101
                            }
                        }
                        if true {
                            switch (^(fi83)) {
                            case 898:
                                if (false && (255 > genFn0())) {
                                    v6[(fi83 << ((4294967295) & 63))] = (((v5 ^ v5) > (v5 << ((682) & 63))) && false)
                                    var v102 uint64 = 2
                                    v102 += v102
                                    _ = v102
                                } else {
                                    var v103 float32 = float32(1e10)
                                    _ = v103
                                }
                            case 9223372036854775807:
                                var v104 int16 = -1
                                _ = v104
                            }
                            v4 = (^(v4))
                            if true {
                                v105 := GenS15{F0: 537}
                                _ = v105
                            }
                        } else {
                            _ = len(v6)
                        }
                    } else {
                        switch ("foo bar" + (("a" + "") + "foo bar")) {
                        case "\n\t":
                            switch -960 {
                            case -16:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            default:
                                v3 = 0
                            }
                        }
                        if true {
                            v106 := GenS16{F0: (-(v3)), F1: fi88}
                            _ = v106
                        } else {
                            _ = len(v2)
                        }
                        if (!false) {
                            v3 -= (v3 << ((uint(v5)) & 63))
                            v4 = (^(v4))
                            fmt.Println(fi88, fi83)
                        }
                        fmt.Println(v21, v21, v3)
                        if (-1 >= -202) {
                            switch "foo bar" {
                            case "":
                                switch "\n\t" {
                                case "":
                                    if (255 == 4294967295) {
                                        var v107 uint64 = (v5 / ((v5) | 1))
                                        _ = v107
                                    } else {
                                        if len(v1) > 0 {
                                            v1[0] = 2147483648
                                        }
                                        var v108 bool = false
                                        fmt.Println(v108)
                                        var v109 int = 1
                                        fmt.Println(v21, v109, v4)
                                        v108 = (-20 <= fi88)
                                        _ = v108
                                        _ = v109
                                    }
                                case "foo bar":
                                    v110 := GenS17{F0: -1}
                                    _ = v110
                                case "hello":
                                    if (0 <= 255) {
                                        var v111 int32 = 2147483647
                                        v3 += 42
                                        var v112 uint8 = 150
                                        _ = v111
                                        _ = v112
                                    }
                                default:
                                    var v113 bool = true
                                    _ = v113
                                }
                            default:
                                fmt.Println(fi88, v5)
                            }
                        }
                        fmt.Println(v3, v5, v4)
                        var v114 uint64 = 1
                        var v115 int8 = -35
                        v115 += (^(v3))
                        _ = v114
                        _ = v115
                    }
                    v5 = (v5 >> ((genFn0()) & 63))
                    _ = v97
                }
            case "foo bar":
                if (!(!false)) {
                    v3 -= v3
                }
            default:
                v116 := GenS18{F0: fi83, F1: 2}
                _ = v116
            }
            if (-32768 != 133) {
                delete(v6, fi88)
                _ = len(v6)
                fmt.Println(v4, v3, fi83)
            } else {
                if false {
                    var v117 int = (-(fi83))
                    if (!(1 != 4294967295)) {
                        if (false && (!true)) {
                            fmt.Println(fi88, v21)
                            if (42 > 134) {
                                v118 := GenS6{F0: 0, F1: -563, F2: v3}
                                if true {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v119 := GenS2{F0: 2147483647}
                                    _ = v119
                                }
                                switch (v117 + fi83) {
                                case 9223372036854775807:
                                    var v120 float32 = (((((v21 + 1e10) - ((0.1 - v21) + (v21 / 1.5))) / float32(-2.25)) - -2.25) * (((float32(3.14159) * (v21 * 3.14159)) / 0.1) - (v21 - (0.0 * (-2.25 * float32(3.14159))))))
                                    _ = v120
                                case -9223372036854775808:
                                    v121 := GenS19{F0: genFn0(), F1: true}
                                    _ = v121
                                case -2147483648:
                                    v122 := GenS20{F0: (v4 & v4), F1: 2}
                                    _ = v122
                                default:
                                    var v123 int32 = -216
                                    _ = v123
                                }
                                _ = v118
                            }
                            switch "a" {
                            case "hello":
                                var v124 int32 = 856
                                v3 = (v3 % ((114) | 1))
                                v124 = v124
                                _ = v124
                            default:
                                if (!(((v3 << ((551) & 63)) <= v3) && (!((-214 != -2147483648) || false)))) {
                                    fmt.Println(fi83)
                                    var v125 uint32 = 493
                                    _ = v125
                                }
                            }
                        }
                        fmt.Println(v3, v21, v21)
                    }
                    _ = len(v2)
                    _ = v117
                }
            }
            _ = len(v1)
            _ = fi88
        }
        _ = fi83
    }
    for fi126 := 0; fi126 < 8; fi126++ {
        fmt.Println(v21, v5)
        _ = fi126
    }
    switch 0 {
    case -947:
        switch "hello" {
        case "hello":
            for fi127 := 0; fi127 < 9; fi127++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi127
            }
        case "\n\t":
            var v128 uint8 = 42
            _ = v128
        case "a":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    case 137:
        var v129 int64 = (v4 * 42)
        _ = v129
    case 127:
        var v130 uint32 = 2147483647
        _ = v130
    default:
        delete(v6, -51)
    }
    for fi131 := 0; fi131 < 14; fi131++ {
        v132 := GenS21{F0: (v4 ^ 0), F1: 871}
        var v133 string = "hello"
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi131
        _ = v132
        _ = v133
    }
    if (true || false) {
        if (4294967295 <= -633) {
            if false {
                for fi134 := 0; fi134 < 18; fi134++ {
                    v135 := GenS22{F0: (v3 + (v3 - (^(v3))))}
                    _ = fi134
                    _ = v135
                }
                var v136 uint64 = (v5 << ((4294967295) & 63))
                v137 := GenS23{F0: 621, F1: (188 != 2)}
                _ = v136
                _ = v137
            } else {
                var v138 float32 = (v21 / v21)
                var v139 uint32 = 846
                _ = v138
                _ = v139
            }
        } else {
            fmt.Println(v5, v4)
            if (!false) {
                var v140 int64 = (v4 << ((255) & 63))
                if false {
                    if ((v5 | 906) >= v5) {
                        var v141 uint8 = genFn1()
                        if true {
                            var v142 int32 = 2147483647
                            if (0 == 1) {
                                var v143 int64 = (v4 >> ((0) & 63))
                                _ = v143
                            } else {
                                if len(v2) > 0 {
                                    v2[0] = 645
                                }
                                switch 0 {
                                case 4294967295:
                                    switch ("" + "\n\t") {
                                    case "":
                                        var v144 int32 = -853
                                        _ = v144
                                    default:
                                        delete(v6, 2147483647)
                                    }
                                case 127:
                                    if (0 <= 1) {
                                        v142 += -657
                                        v145 := GenS4{F0: 4294967295, F1: 0}
                                        _ = v145
                                    } else {
                                        var v146 float64 = 0.0
                                        v5 = v5
                                        var v147 bool = false
                                        _ = v146
                                        _ = v147
                                    }
                                case -1:
                                    v148 := GenS8{F0: v5, F1: false}
                                    _ = v148
                                default:
                                    if (!("foo bar" < "foo bar")) {
                                        var v149 string = ("hello" + "hello")
                                        switch 63 {
                                        case 128:
                                            var v150 int16 = -1
                                            _ = v150
                                        case -783:
                                            if ((3.14159 - ((1e10 + 0.1) / (0.0 + 0.1))) <= 0.1) {
                                                switch 32767 {
                                                case 32768:
                                                    v151 := GenS7{F0: (v5 * 42), F1: 1}
                                                    _ = v151
                                                }
                                            }
                                        default:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                        _ = v149
                                    } else {
                                        v4 += (v4 | v140)
                                        v152 := GenS24{F0: (v3 ^ (v3 << ((genFn0()) & 63)))}
                                        var v153 uint64 = 4294967295
                                        _ = v152
                                        _ = v153
                                    }
                                }
                            }
                            for fi154 := 0; fi154 < 7; fi154++ {
                                var v155 uint64 = v5
                                switch 2147483647 {
                                case 32768:
                                    v141 = 101
                                case -2147483648:
                                    var v156 uint32 = 2147483648
                                    _ = v156
                                case 2147483647:
                                    v140 = (^(v4))
                                }
                                v157 := GenS22{F0: (v3 - (-(v3)))}
                                _ = fi154
                                _ = v155
                                _ = v157
                            }
                            _ = v142
                        }
                        v3 = -102
                        _ = v141
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v158 uint32 = 1
                        switch 128 {
                        case 127:
                            v158 = (^(v158))
                        }
                        _ = v158
                    }
                    v159 := GenS25{F0: 168}
                    v3 = (v3 + v3)
                    _ = v159
                }
                switch -674 {
                case 2147483647:
                    v140 -= (^(v140))
                case 0:
                    for fi160 := 0; fi160 < 19; fi160++ {
                        v3 -= v3
                        _ = fi160
                    }
                }
                _ = v140
            }
        }
        for fi161 := 0; fi161 < 12; fi161++ {
            v3 = (v3 * (v3 ^ (v3 | int8(v4))))
            _ = fi161
        }
        var v162 int = 9223372036854775807
        _ = v162
    } else {
        for fi163 := 0; fi163 < 19; fi163++ {
            v164 := GenS9{F0: 197, F1: (fi163 >> ((65535) & 63)), F2: "a"}
            var v165 uint64 = 435
            delete(v6, -794)
            _ = fi163
            _ = v164
            _ = v165
        }
        fmt.Println(v4)
    }
    for fi166 := 0; fi166 < 14; fi166++ {
        v5 = (^(v5))
        v21 = (v21 - (((1e10 / (v21 * 0.0)) / v21) - v21))
        _ = fi166
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
