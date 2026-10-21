package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint8; F1 uint }
type GenS7 struct { F0 string; F1 uint16; F2 uint8 }
type GenS8 struct { F0 uint32; F1 int64 }
type GenS9 struct { F0 bool; F1 int8; F2 uint32 }
type GenS10 struct { F0 string; F1 bool; F2 int64 }
type GenS11 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS12 struct { F0 bool; F1 uint64; F2 int8 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 uint16; F1 uint64 }
type GenS15 struct { F0 int; F1 bool }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 uint32 }
type GenS18 struct { F0 int8; F1 int64 }
type GenS19 struct { F0 int16 }
func genFn0(p0 bool) int16 {
    p0 = p0
    var v6 uint64 = 273
    var v7 float64 = 1.5
    p0 = p0
    _ = p0
    _ = v6
    _ = v7
    return 32767
}
func genFn1() int64 {
    v8 := GenS1{F0: 602, F1: 946}
    if (-128 == 0) {
        var v9 bool = true
        v10 := []int{42, 32767, 32767, 32767}
        _ = v10
        _ = v9
    } else {
        v11 := GenS1{F0: 2, F1: 2}
        var v12 uint32 = 1
        fmt.Println(v12)
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 float32 = (-2.25 - (((1.5 + 0.1) - (((float32(float64((1e10 - (0.1 * 1.5)))) * -2.25) - (float32((1e10 + 3.14159)) + (3.14159 + float32(0.0)))) * float32(0.0))) + 3.14159))
        _ = fi13
        _ = v14
    }
    if (1e10 < 1.5) {
        for fi15 := 0; fi15 < 17; fi15++ {
            var v16 float64 = (0.1 / float64(3.14159))
            _ = fi15
            _ = v16
        }
        var v17 int8 = 62
        v18 := GenS1{F0: 719, F1: 32767}
        _ = v17
        _ = v18
    } else {
        if (1 > 305) {
            if (!true) {
                v19 := []int{4294967295, 32768, 42, 999}
                _ = v19
            } else {
                if true {
                    if (473 <= 255) {
                        var v20 int32 = -1
                        _ = v20
                    }
                }
            }
        } else {
            v21 := GenS2{F0: 255}
            switch 32767 {
            case 9223372036854775807:
                if (-1 > -27) {
                    var v22 string = "hello"
                    if (65535 < 2) {
                        if (("foo bar" + (v22 + (v22 + (v22 + v22)))) <= "a") {
                            var v23 int16 = 0
                            v23 = v23
                            _ = v23
                        } else {
                            v22 = v22
                        }
                        var v24 int16 = 0
                        var v25 int32 = -994
                        _ = v24
                        _ = v25
                    }
                    _ = v22
                } else {
                    var v26 int16 = -1
                    _ = v26
                }
            case 2147483647:
                var v27 float64 = (0.1 + float64(((1.5 * -2.25) + float32((float64(1e10) * 0.0)))))
                _ = v27
            }
            _ = v21
        }
        for fi28 := 0; fi28 < 7; fi28++ {
            for fi29 := 0; fi29 < 2; fi29++ {
                switch (-(fi28)) {
                case 0:
                    var v30 uint32 = 2147483648
                    var v31 int8 = 1
                    fmt.Println(fi28, v30)
                    _ = v30
                    _ = v31
                }
                _ = fi29
            }
            switch "" {
            case "a":
                var v32 uint64 = 4294967295
                _ = v32
            default:
                if true {
                    var v33 uint8 = 118
                    v34 := GenS2{F0: v33}
                    v33 -= v33
                    _ = v33
                    _ = v34
                } else {
                    var v35 int = fi28
                    if (1000 >= genFn0(false)) {
                        v35 -= (fi28 >> ((255) & 63))
                        var v36 uint64 = 919
                        _ = v36
                    }
                    v35 -= v35
                    _ = v35
                }
            }
            if (((-2.25 * ((((1e10 + 0.0) / -2.25) / 1e10) / 1.0)) / 3.14159) > 0.0) {
                v37 := GenS1{F0: 65535, F1: 88}
                _ = v37
            } else {
                fmt.Println(fi28)
                var v38 bool = true
                if (!(0 >= 32767)) {
                    fmt.Println(v38, v38)
                    v39 := []int{fi28, 43, 32767}
                    for fi40 := 0; fi40 < 14; fi40++ {
                        var v41 int16 = 1
                        var v42 int = fi40
                        v38 = true
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                    _ = v39
                }
                _ = v38
            }
            _ = fi28
        }
        var v43 bool = true
        _ = v43
    }
    for fi44 := 0; fi44 < 1; fi44++ {
        v45 := []int{3}
        var v46 uint32 = 4294967295
        if len(v45) > 0 {
            v45[0] = (fi44 | 793)
        }
        _ = fi44
        _ = v45
        _ = v46
    }
    _ = v8
    return 56
}

func main() {
    v1 := []int{}
    var v2 string = ""
    v3 := []uint32{1, 2, 65535, 1}
    var v4 float64 = (1.5 - 3.14159)
    var v5 int = 128
    if (v5 == (v5 - v5)) {
        v47 := GenS2{F0: 2}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        switch v2 {
        case "foo bar":
            v48 := GenS1{F0: 32767, F1: 255}
            _ = v48
        case "\n\t":
            var v49 int = v5
            _ = v49
        default:
            v50 := GenS3{F0: 4294967295}
            _ = v50
        }
        _ = v47
    } else {
        v2 = (v2 + (v2 + "hello"))
        switch v5 {
        case -272:
            switch (v5 & (v5 << ((767) & 63))) {
            case 127:
                for fi51 := 0; fi51 < 15; fi51++ {
                    v52 := GenS3{F0: 977}
                    _ = fi51
                    _ = v52
                }
            case 351:
                if false {
                    _ = len(v1)
                    for fi53 := 0; fi53 < 7; fi53++ {
                        var v54 float64 = (-2.25 - (((v4 + 3.14159) + v4) / ((0.1 - 1.5) - (float64(-2.25) * (3.14159 - (0.0 * v4))))))
                        var v55 float64 = float64(0.1)
                        if true {
                            v54 += v4
                        } else {
                            _ = len(v1)
                            fmt.Println(v5, v54, v2)
                            v54 -= v4
                        }
                        _ = fi53
                        _ = v54
                        _ = v55
                    }
                    v2 = v2
                } else {
                    var v56 int8 = -120
                    fmt.Println(v4, v5)
                    var v57 uint32 = 2
                    _ = v56
                    _ = v57
                }
            case -865:
                for fi58 := 0; fi58 < 14; fi58++ {
                    for fi59 := 0; fi59 < 18; fi59++ {
                        var v60 float32 = ((-2.25 + -2.25) + 0.0)
                        var v61 int32 = 42
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi59
                        _ = v60
                        _ = v61
                    }
                    _ = fi58
                }
            }
        case -523:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case -1:
            if (77 > 1) {
                for fi62 := 0; fi62 < 20; fi62++ {
                    if true {
                        if (65535 > 4294967295) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v63 := GenS4{F0: 0, F1: 42, F2: 0}
                            switch (v5 - (v5 & (fi62 - (^(fi62))))) {
                            case 0:
                                for fi64 := 0; fi64 < 7; fi64++ {
                                    v5 += 9223372036854775807
                                    v4 -= (1e10 + v4)
                                    var v65 bool = ((true && false) && false)
                                    _ = fi64
                                    _ = v65
                                }
                            case -201:
                                fmt.Println(v4, v5, v5)
                            case 42:
                                v4 -= v4
                            }
                            _ = v63
                        } else {
                            v4 = -2.25
                            switch "" {
                            case "foo bar":
                                if false {
                                    var v66 float64 = 3.14159
                                    fmt.Println(v4, v66, v2)
                                    v67 := GenS5{F0: 1}
                                    _ = v66
                                    _ = v67
                                }
                            case "a":
                                if (((0.0 / 1.0) * 0.1) <= float32(v4)) {
                                    v68 := GenS6{F0: 128, F1: 0}
                                    v2 = v2
                                    _ = v68
                                } else {
                                    v1 = append(v1, (v5 & fi62))
                                    var v69 uint8 = 61
                                    for fi70 := 0; fi70 < 10; fi70++ {
                                        var v71 int32 = 1
                                        var v72 float32 = (-2.25 * 0.0)
                                        if false {
                                            if false {
                                                fmt.Println(v5, fi62)
                                                var v73 bool = (1 > 0)
                                                v73 = (!true)
                                                _ = v73
                                            } else {
                                                var v74 int16 = 347
                                                fmt.Println(v72, v5, fi62)
                                                _ = v74
                                            }
                                            v72 -= ((v72 * (float32((1e10 + 1.5)) + float32(v4))) * v72)
                                        } else {
                                            v4 += v4
                                            var v75 float64 = (0.0 / (1e10 / v4))
                                            fmt.Println(fi70, v71, v4)
                                            _ = v75
                                        }
                                        _ = fi70
                                        _ = v71
                                        _ = v72
                                    }
                                    _ = v69
                                }
                            }
                        }
                    }
                    _ = fi62
                }
            } else {
                var v76 bool = (!true)
                _ = v76
            }
        default:
            switch v5 {
            case -1:
                var v77 int16 = 387
                _ = v77
            case -845:
                v2 = ("\n\t" + v2)
            default:
                if ((v4 * (v4 + (v4 * v4))) <= float64((0.0 + (0.1 * (1e10 * (((1.5 / 3.14159) * 1.5) / 0.1)))))) {
                    if len(v1) > 0 {
                        v1[0] = (v5 / ((1) | 1))
                    }
                    fmt.Println(v2, v2, v4)
                } else {
                    var v78 string = v2
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if (!(752 < 0)) {
                        if true {
                            if true {
                                for fi79 := 0; fi79 < 6; fi79++ {
                                    if len(v3) > 0 {
                                        v3[0] = 42
                                    }
                                    fmt.Println(v2, v5, v4)
                                    var v80 int = (v5 * fi79)
                                    _ = fi79
                                    _ = v80
                                }
                                v5 -= -84
                                v5 = 128
                            }
                            var v81 float32 = 0.0
                            var v82 int32 = 2147483647
                            _ = v81
                            _ = v82
                        } else {
                            var v83 int16 = -32768
                            _ = v83
                        }
                        v2 = v78
                        v84 := GenS7{F0: v78, F1: 32767, F2: 2}
                        _ = v84
                    } else {
                        if (!true) {
                            var v85 bool = false
                            _ = len(v3)
                            _ = v85
                        }
                        var v86 int32 = -620
                        _ = v86
                    }
                    _ = v78
                }
            }
        }
    }
    v87 := GenS4{F0: 42, F1: -9, F2: 32767}
    for fi88 := 0; fi88 < 19; fi88++ {
        v2 = "foo bar"
        fmt.Println(fi88, fi88, fi88)
        _ = fi88
    }
    v1 = append(v1, -234)
    if true {
        if (32767 >= 538) {
            fmt.Println(v2)
        } else {
            if (!(v4 >= v4)) {
                for fi89 := 0; fi89 < 19; fi89++ {
                    v2 = "foo bar"
                    switch ((v2 + v2) + (v2 + v2)) {
                    case "":
                        if ((!false) && false) {
                            var v90 int = -881
                            v4 -= v4
                            for fi91 := 0; fi91 < 15; fi91++ {
                                var v92 float64 = 3.14159
                                var v93 int8 = 0
                                _ = fi91
                                _ = v92
                                _ = v93
                            }
                            _ = v90
                        }
                    case "foo bar":
                        v94 := GenS2{F0: 2}
                        _ = v94
                    case "hello":
                        fmt.Println(v4, fi89)
                    }
                    v4 -= ((-2.25 - v4) * ((1e10 * v4) - ((-2.25 * v4) * (((v4 - float64((0.0 - (-2.25 / (float32(1e10) * (-2.25 + (float32(1e10) - 0.1))))))) + (v4 * ((-2.25 - (v4 - (v4 - float64(((float32(v4) / 1e10) * ((0.0 - 1e10) / (float32(v4) * 0.0))))))) / v4))) * v4))))
                    _ = fi89
                }
            } else {
                var v95 float32 = 0.0
                if ((!true) && (-128 >= -128)) {
                    v4 -= (((0.1 + 1e10) / v4) - 0.1)
                }
                _ = v95
            }
        }
    } else {
        if true {
            for fi96 := 0; fi96 < 15; fi96++ {
                if false {
                    switch v5 {
                    case 2147483647:
                        switch fi96 {
                        case -9223372036854775808:
                            fmt.Println(v5, v4)
                        case -2147483648:
                            var v97 float64 = 1.5
                            _ = v97
                        }
                    default:
                        v5 = (^(fi96))
                    }
                    v2 = v2
                } else {
                    v98 := GenS2{F0: 255}
                    if (true || (!true)) {
                        var v99 float64 = (-2.25 * -2.25)
                        var v100 uint16 = 0
                        _ = v100
                        _ = v99
                    } else {
                        if (564 >= 432) {
                            for fi101 := 0; fi101 < 1; fi101++ {
                                v4 = ((1.5 - float64((1e10 - -2.25))) - v4)
                                v2 = (((v2 + v2) + v2) + "hello")
                                fmt.Println(v4, v5, v4)
                                _ = fi101
                            }
                            var v102 int8 = 126
                            _ = v102
                        }
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v98
                }
                _ = fi96
            }
            if (694 == 726) {
                _ = len(v1)
            } else {
                v5 += (v5 >> ((65535) & 63))
                if true {
                    fmt.Println(v2, v5, v5)
                    var v103 uint32 = 1
                    _ = v103
                }
            }
            if true {
                if (9223372036854775807 <= -173) {
                    for fi104 := 0; fi104 < 16; fi104++ {
                        v105 := GenS8{F0: 645, F1: 0}
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi104
                        _ = v105
                    }
                    var v106 uint16 = 401
                    v107 := GenS9{F0: ((!((((v4 + (v4 / v4)) * v4) * v4) < 1e10)) || false), F1: -58, F2: 42}
                    _ = v106
                    _ = v107
                } else {
                    if (false && true) {
                        v108 := GenS2{F0: 2}
                        for fi109 := 0; fi109 < 20; fi109++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi109
                        }
                        _ = v108
                    }
                    for fi110 := 0; fi110 < 2; fi110++ {
                        var v111 uint64 = 0
                        for fi112 := 0; fi112 < 12; fi112++ {
                            v5 += v5
                            v111 = v111
                            if len(v1) > 0 {
                                v1[0] = 0
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            fmt.Println(v5)
                            _ = fi112
                        }
                        _ = fi110
                        _ = v111
                    }
                }
                v113 := GenS10{F0: "hello", F1: true, F2: 9223372036854775807}
                switch 127 {
                case 9223372036854775807:
                    if false {
                        if true {
                            v1 = append(v1, v5)
                        }
                        for fi114 := 0; fi114 < 9; fi114++ {
                            for fi115 := 0; fi115 < 9; fi115++ {
                                var v116 int = v5
                                _ = fi115
                                _ = v116
                            }
                            _ = fi114
                        }
                    }
                case 32768:
                    switch ("\n\t" + "\n\t") {
                    case "":
                        v5 = v5
                    case "foo bar":
                        var v117 float64 = (((v4 - (((-2.25 * (v4 + float64(((1e10 + float32((1e10 * ((((float64(float32((float64((0.0 - 1.5)) * (v4 + float64(1.5))))) * (v4 * (v4 * (float64(3.14159) - float64(3.14159))))) - 0.1) + v4) / v4)))) + 1e10)))) + 0.0) / v4)) + 1e10) * -2.25)
                        _ = v117
                    case "hello":
                        for fi118 := 0; fi118 < 20; fi118++ {
                            var v119 uint8 = 127
                            if (-1 >= 42) {
                                v3 = append(v3, 2)
                                v4 += v4
                                if (false && (1 <= -2)) {
                                    var v120 uint64 = 2
                                    if false {
                                        v5 += (fi118 ^ fi118)
                                    }
                                    _ = v120
                                } else {
                                    for fi121 := 0; fi121 < 6; fi121++ {
                                        switch "hello" {
                                        case "\n\t":
                                            var v122 int8 = 42
                                            _ = v122
                                        case "hello":
                                            var v123 uint = 18446744073709551615
                                            _ = v123
                                        case "a":
                                            v119 += v119
                                        }
                                        _ = fi121
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = (^(fi118))
                                    }
                                    if true {
                                        v124 := GenS11{F0: 1, F1: v119, F2: 42}
                                        _ = v124
                                    }
                                }
                            } else {
                                if false {
                                    if false {
                                        switch v2 {
                                        case "\n\t":
                                            var v125 int16 = -32768
                                            _ = v125
                                        case "hello":
                                            var v126 int16 = 1000
                                            _ = v126
                                        case "foo bar":
                                            fmt.Println(fi118, v4)
                                        }
                                        switch v5 {
                                        case -846:
                                            for fi127 := 0; fi127 < 9; fi127++ {
                                                fmt.Println(v119)
                                                _ = fi127
                                            }
                                        }
                                    }
                                    var v128 uint = 42
                                    v119 = (v119 | (v119 ^ v119))
                                    _ = v128
                                }
                            }
                            fmt.Println(v4, v119)
                            _ = fi118
                            _ = v119
                        }
                    default:
                        var v129 int8 = 94
                        _ = v129
                    }
                default:
                    var v130 string = ("" + v2)
                    _ = v130
                }
                _ = v113
            } else {
                for fi131 := 0; fi131 < 15; fi131++ {
                    v4 = (v4 / v4)
                    _ = fi131
                }
            }
        } else {
            var v132 uint = 255
            _ = v132
        }
        switch v5 {
        case -726:
            switch v5 {
            case 293:
                switch v2 {
                case "":
                    var v133 uint16 = 255
                    _ = v133
                }
            case -1:
                fmt.Println(v4, v2)
            case 0:
                fmt.Println(v5, v2)
            default:
                switch 32768 {
                case 4294967295:
                    fmt.Println(v4)
                }
            }
        case 128:
            var v134 uint = 544
            _ = v134
        case 1:
            for fi135 := 0; fi135 < 2; fi135++ {
                v136 := GenS3{F0: 4294967295}
                _ = fi135
                _ = v136
            }
        }
    }
    fmt.Println(v2)
    if true {
        for fi137 := 0; fi137 < 5; fi137++ {
            switch (fi137 & (fi137 + (v5 >> ((42) & 63)))) {
            case 9223372036854775807:
                switch 971 {
                case 0:
                    if true {
                        v3 = append(v3, 2)
                    } else {
                        var v138 int = (v5 - 1)
                        var v139 int64 = -137
                        _ = v138
                        _ = v139
                    }
                default:
                    v5 = (-(v5))
                }
            }
            if ((((1e10 + 0.1) - (((float32(v4) - 0.1) / 3.14159) - 0.1)) - (0.1 + 3.14159)) != 0.1) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if (!(false && true)) {
                    if (true || true) {
                        if (!false) {
                            if false {
                                v3 = append(v3, 65535)
                                if (!true) {
                                    fmt.Println(v4)
                                    if true {
                                        var v140 int = (fi137 / ((v5) | 1))
                                        switch 295 {
                                        case 0:
                                            v141 := GenS11{F0: 230, F1: 0, F2: 1}
                                            _ = v141
                                        case 9223372036854775807:
                                            switch v2 {
                                            case "foo bar":
                                                v2 = v2
                                            }
                                        case -2147483648:
                                            v2 = "foo bar"
                                        default:
                                            var v142 uint64 = 42
                                            _ = v142
                                        }
                                        for fi143 := 0; fi143 < 11; fi143++ {
                                            if (-1 != 2147483647) {
                                                v1 = append(v1, fi143)
                                                switch fi143 {
                                                case 2147483647:
                                                    switch v140 {
                                                    case -886:
                                                        v5 -= 42
                                                        v3 = append(v3, 42)
                                                    default:
                                                        var v144 uint8 = 127
                                                        _ = v144
                                                    }
                                                    var v145 string = v2
                                                    if false {
                                                        v5 = (v5 * (v140 & 126))
                                                    } else {
                                                        var v146 string = v145
                                                        _ = v146
                                                    }
                                                    _ = v145
                                                case -9223372036854775808:
                                                    var v147 float32 = 1.5
                                                    _ = v147
                                                default:
                                                    v148 := GenS6{F0: 255, F1: 0}
                                                    _ = v148
                                                }
                                                v149 := GenS7{F0: v2, F1: 827, F2: 79}
                                                _ = len(v1)
                                                _ = v149
                                            }
                                            _ = fi143
                                        }
                                        _ = v140
                                    }
                                }
                                v150 := GenS12{F0: false, F1: 846, F2: 1}
                                _ = v150
                            }
                        }
                        var v151 uint = 65535
                        if (22 <= uint8(v151)) {
                            if (v5 < (v5 / ((v5) | 1))) {
                                var v152 string = (v2 + "a")
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                if true {
                                    switch v5 {
                                    case 128:
                                        v5 += fi137
                                    case 9223372036854775807:
                                        fmt.Println(v151, v5, v152)
                                    case 32768:
                                        var v153 int32 = 0
                                        _ = v153
                                    default:
                                        if (genFn1() > -1) {
                                            var v154 uint32 = uint32(v151)
                                            v155 := GenS2{F0: 209}
                                            v5 += 42
                                            _ = v154
                                            _ = v155
                                        } else {
                                            for fi156 := 0; fi156 < 10; fi156++ {
                                                fmt.Println(v152)
                                                v157 := GenS13{F0: 0}
                                                v158 := GenS14{F0: 65535, F1: 42}
                                                _ = fi156
                                                _ = v157
                                                _ = v158
                                            }
                                        }
                                    }
                                    v5 = v5
                                }
                                _ = v152
                            }
                            for fi159 := 0; fi159 < 8; fi159++ {
                                switch v2 {
                                case "foo bar":
                                    v160 := GenS15{F0: v5, F1: (1.5 >= (float64((((((((float32(v4) + float32(v4)) * 0.0) - ((-2.25 * float32((3.14159 - (v4 - ((v4 * 0.1) / 1e10))))) / float32((float64(1e10) * v4)))) * 1e10) / (((float32(((0.1 + v4) / -2.25)) * -2.25) - float32((float64(0.1) * (float64(0.0) - (v4 + v4))))) + ((3.14159 * (((-2.25 + float32((3.14159 / (1e10 + 0.1)))) + float32((1.5 + -2.25))) / (float32(v4) * 3.14159))) * -2.25))) * 1.5) + -2.25)) * (float64((3.14159 + float32(float64(0.1)))) - 1e10)))}
                                    _ = v160
                                case "a":
                                    var v161 uint = v151
                                    _ = v161
                                case "hello":
                                    fmt.Println(v5, fi137, v151)
                                default:
                                    v162 := GenS2{F0: 255}
                                    _ = v162
                                }
                                if true {
                                    var v163 uint32 = 65535
                                    _ = v163
                                }
                                v151 = 0
                                _ = fi159
                            }
                        } else {
                            if false {
                                if false {
                                    if false {
                                        for fi164 := 0; fi164 < 10; fi164++ {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v165 := GenS7{F0: v2, F1: 0, F2: 255}
                                            switch 0 {
                                            case 868:
                                                fmt.Println(v2, fi137)
                                            case 843:
                                                v4 -= v4
                                            case -2147483648:
                                                var v166 float32 = (((-2.25 - float32((3.14159 - (-2.25 + v4)))) + float32((1e10 / v4))) * float32(0.1))
                                                _ = v166
                                            default:
                                                if true {
                                                    v167 := GenS11{F0: 1, F1: 42, F2: 1}
                                                    fmt.Println(fi137, v151)
                                                    _ = v167
                                                } else {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    v4 = v4
                                                }
                                            }
                                            v151 = (^(v151))
                                            _ = fi164
                                            _ = v165
                                        }
                                    }
                                    var v168 int32 = -238
                                    v5 = (fi137 % (((^(fi137))) | 1))
                                    _ = v168
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                if (false || (!((true || false) && (225 != 78)))) {
                                    fmt.Println(v151)
                                }
                            }
                            if ((v151 - (v151 / ((v151) | 1))) != v151) {
                                for fi169 := 0; fi169 < 12; fi169++ {
                                    v4 += v4
                                    v5 = 4294967295
                                    _ = fi169
                                }
                            }
                        }
                        _ = v151
                    } else {
                        if (790 == 1) {
                            v4 = v4
                        } else {
                            if false {
                                var v170 int32 = 0
                                _ = v170
                            } else {
                                for fi171 := 0; fi171 < 9; fi171++ {
                                    var v172 int32 = -1
                                    fmt.Println(fi171, v5, v2)
                                    var v173 string = (v2 + v2)
                                    v172 -= v172
                                    _ = fi171
                                    _ = v172
                                    _ = v173
                                }
                            }
                            var v174 int16 = -32768
                            _ = v174
                        }
                        if (true || ((1.5 / v4) == v4)) {
                            v2 = (("\n\t" + "\n\t") + v2)
                            var v175 uint16 = 42
                            var v176 string = v2
                            _ = v175
                            _ = v176
                        }
                        v2 = v2
                    }
                    if (393 == 42) {
                        switch v2 {
                        case "":
                            switch 1 {
                            case 2147483647:
                                v5 += v5
                            case 159:
                                var v177 int16 = 1
                                _ = v177
                            case -254:
                                v178 := GenS16{F0: 1}
                                _ = v178
                            }
                        case "\n\t":
                            v2 = ("a" + v2)
                        case "a":
                            v179 := GenS17{F0: 65535}
                            _ = v179
                        }
                        v4 += v4
                        switch fi137 {
                        case 9223372036854775807:
                            for fi180 := 0; fi180 < 2; fi180++ {
                                var v181 uint8 = 128
                                v181 += (v181 ^ v181)
                                v2 = (("hello" + v2) + v2)
                                v2 = v2
                                v4 += 3.14159
                                _ = fi180
                                _ = v181
                            }
                        case 4294967295:
                            if len(v3) > 0 {
                                v3[0] = 2147483648
                            }
                        default:
                            switch v2 {
                            case "":
                                for fi182 := 0; fi182 < 4; fi182++ {
                                    v183 := GenS13{F0: 42}
                                    var v184 int16 = 42
                                    fmt.Println(v2, v184)
                                    if false {
                                        _ = len(v3)
                                        var v185 uint8 = 0
                                        v186 := GenS11{F0: 2147483648, F1: (v185 >> ((2) & 63)), F2: 490}
                                        _ = v185
                                        _ = v186
                                    } else {
                                        var v187 int32 = 42
                                        v5 = -1
                                        fmt.Println(fi137, fi182)
                                        _ = v187
                                    }
                                    switch (fi182 - (fi182 << ((18446744073709551615) & 63))) {
                                    case 225:
                                        fmt.Println(v4, fi182)
                                    case -1:
                                        switch v2 {
                                        case "\n\t":
                                            if true {
                                                fmt.Println(v4, v184)
                                                v188 := GenS11{F0: 2, F1: 127, F2: 255}
                                                _ = v188
                                            } else {
                                                var v189 string = "hello"
                                                var v190 float64 = 1.5
                                                _ = v189
                                                _ = v190
                                            }
                                        case "":
                                            v191 := GenS18{F0: 37, F1: 540}
                                            _ = v191
                                        case "a":
                                            var v192 bool = (1 <= 558)
                                            _ = v192
                                        default:
                                            var v193 bool = true
                                            _ = v193
                                        }
                                    case -2147483648:
                                        fmt.Println(v4)
                                    }
                                    _ = fi182
                                    _ = v183
                                    _ = v184
                                }
                            case "hello":
                                if len(v3) > 0 {
                                    v3[0] = 2
                                }
                            }
                        }
                    }
                } else {
                    fmt.Println(v4)
                }
            } else {
                v5 = (v5 + fi137)
            }
            _ = fi137
        }
        v4 -= v4
    } else {
        v4 = v4
        var v194 int32 = 1
        for fi195 := 0; fi195 < 20; fi195++ {
            if false {
                _ = len(v1)
            } else {
                var v196 uint16 = 255
                _ = v196
            }
            v197 := GenS17{F0: 0}
            _ = fi195
            _ = v197
        }
        _ = v194
    }
    if false {
        var v198 uint64 = 18446744073709551615
        _ = v198
    }
    fmt.Println(v2)
    fmt.Println(v4, v2)
    v199 := GenS19{F0: -1}
    v4 = float64(1e10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v199
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v87
}
