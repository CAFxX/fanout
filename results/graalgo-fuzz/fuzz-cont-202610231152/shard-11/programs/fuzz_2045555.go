package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int16; F1 int8 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 int64 }
func genFn0() uint32 {
    v6 := GenS1{F0: 1}
    switch (("\n\t" + ("foo bar" + "")) + ("" + "hello")) {
    case "a":
        var v7 uint32 = 303
        _ = v7
    case "\n\t":
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 uint32 = 4294967295
            v9 -= (v9 >> ((4294967295) & 63))
            v10 := GenS1{F0: 2}
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    if false {
        var v11 bool = true
        _ = v11
    } else {
        v12 := GenS1{F0: 18446744073709551615}
        var v13 uint32 = 2147483647
        for fi14 := 0; fi14 < 6; fi14++ {
            switch "a" {
            case "":
                fmt.Println(fi14, fi14)
            case "foo bar":
                v13 = v13
            default:
                switch (("foo bar" + "foo bar") + "\n\t") {
                case "":
                    if (((^(v13)) <= v13) || (700 > 969)) {
                        for fi15 := 0; fi15 < 1; fi15++ {
                            if false {
                                v16 := GenS1{F0: 18446744073709551615}
                                var v17 uint = 18446744073709551615
                                _ = v16
                                _ = v17
                            } else {
                                var v18 int = (fi14 ^ fi14)
                                v13 = 164
                                if (!true) {
                                    fmt.Println(fi15)
                                }
                                _ = v18
                            }
                            v19 := []int{(^(fi14)), 1}
                            _ = fi15
                            _ = v19
                        }
                        switch (fi14 ^ 128) {
                        case -344:
                            v20 := GenS2{F0: fi14, F1: 0}
                            _ = v20
                        case 0:
                            for fi21 := 0; fi21 < 12; fi21++ {
                                fmt.Println(fi14, v13, fi14)
                                v22 := []int{(^(fi21)), fi21, (^(fi21))}
                                _ = fi21
                                _ = v22
                            }
                        case 4294967295:
                            v23 := GenS3{F0: 200}
                            _ = v23
                        }
                    }
                }
            }
            v24 := GenS4{F0: 128, F1: fi14}
            _ = fi14
            _ = v24
        }
        _ = v12
        _ = v13
    }
    _ = v6
    return 2
}
func genFn1(p0 int8, p1 uint16) uint64 {
    for fi25 := 0; fi25 < 9; fi25++ {
        if true {
            var v26 int8 = 0
            _ = v26
        } else {
            var v27 uint32 = 2147483648
            for fi28 := 0; fi28 < 19; fi28++ {
                fmt.Println(fi25)
                _ = fi28
            }
            var v29 uint32 = v27
            _ = v27
            _ = v29
        }
        fmt.Println(p0, p1)
        if ((1e10 * (1.5 - 3.14159)) >= (float64(0.1) * 0.0)) {
            var v30 int = -80
            _ = v30
        } else {
            fmt.Println(p0, fi25, fi25)
        }
        _ = fi25
    }
    var v31 uint = 571
    p0 -= 0
    _ = p0
    _ = p1
    _ = v31
    return 0
}

func main() {
    var v1 uint8 = 2
    var v2 uint64 = 12
    var v3 string = ("hello" + ("" + "foo bar"))
    v4 := []int{9223372036854775807, 42, 32767, 32768}
    v5 := make(map[string]string)
    if false {
        v32 := GenS5{F0: true, F1: 1}
        var v33 int32 = -2147483648
        _ = len(v5)
        _ = v32
        _ = v33
    } else {
        for fi34 := 0; fi34 < 10; fi34++ {
            var v35 int8 = 127
            var v36 int = fi34
            if (-9223372036854775808 == -1) {
                if len(v4) > 0 {
                    v4[0] = (fi34 / ((fi34) | 1))
                }
            } else {
                v3 = v3
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            _ = fi34
            _ = v35
            _ = v36
        }
        var v37 float64 = 1e10
        if (!false) {
            var v38 int = 4294967295
            v39 := GenS3{F0: v1}
            _ = v38
            _ = v39
        } else {
            v5[v3] = v3
            for fi40 := 0; fi40 < 13; fi40++ {
                var v41 uint8 = v1
                v41 = (v41 - 2)
                _ = fi40
                _ = v41
            }
        }
        _ = v37
    }
    var v42 uint32 = 4294967295
    for fi43 := 0; fi43 < 3; fi43++ {
        v4 = append(v4, (-(fi43)))
        if (401 <= 2) {
            switch ("" + "a") {
            case "foo bar":
                for fi44 := 0; fi44 < 6; fi44++ {
                    v42 = v42
                    v3 = v3
                    _ = fi44
                }
            case "":
                fmt.Println(v2, fi43)
            case "hello":
                var v45 int16 = 1
                _ = v45
            }
        }
        if (false && (-2.25 == (-2.25 + 0.0))) {
            var v46 int64 = 9223372036854775807
            _ = v46
        } else {
            var v47 uint8 = v1
            _ = v47
        }
        _ = fi43
    }
    for fi48 := 0; fi48 < 16; fi48++ {
        v5["foo bar"] = (v3 + v3)
        switch v3 {
        case "hello":
            if false {
                var v49 int64 = 0
                v50 := GenS4{F0: (fi48 | -9223372036854775808), F1: -332}
                v2 += (v2 | v2)
                _ = v49
                _ = v50
            }
        case "\n\t":
            v51 := GenS5{F0: (v3 == v3), F1: 870}
            _ = v51
        }
        if false {
            var v52 uint32 = 2147483647
            v1 = 42
            _ = v52
        }
        _ = fi48
    }
    var v53 int64 = 9223372036854775807
    v2 -= 4294967295
    var v54 float64 = (1.5 - (0.1 / float64(3.14159)))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch (v3 + v3) {
    case "foo bar":
        for fi55 := 0; fi55 < 4; fi55++ {
            for fi56 := 0; fi56 < 1; fi56++ {
                if (fi55 > fi56) {
                    switch (fi56 | fi56) {
                    case 371:
                        var v57 uint8 = (v1 | 2)
                        _ = v57
                    }
                    v53 = v53
                    var v58 bool = (4294967295 <= 18446744073709551615)
                    _ = v58
                } else {
                    v3 = v3
                    switch ("" + "\n\t") {
                    case "hello":
                        v3 = v3
                    case "\n\t":
                        var v59 float32 = (3.14159 - (0.1 * (1.5 / (0.1 * -2.25))))
                        _ = v59
                    case "":
                        if true {
                            if false {
                                fmt.Println(v2, v3, v1)
                                v53 += (v53 + v53)
                            } else {
                                switch "hello" {
                                case "a":
                                    switch (fi55 & (fi56 << ((18446744073709551615) & 63))) {
                                    case 529:
                                        v5[v3] = ("\n\t" + (((("" + v3) + ((((v3 + "\n\t") + v3) + "a") + v3)) + v3) + "a"))
                                    default:
                                        var v60 uint8 = v1
                                        _ = v60
                                    }
                                case "hello":
                                    var v61 string = ("hello" + v3)
                                    _ = len(v4)
                                    _ = v61
                                case "\n\t":
                                    var v62 uint64 = (v2 >> ((65535) & 63))
                                    _ = v62
                                default:
                                    v53 = 42
                                }
                                fmt.Println(v2, v1, v53)
                            }
                        } else {
                            var v63 int16 = -32768
                            _ = v63
                        }
                    default:
                        v5[((v3 + v3) + ((((v3 + v3) + v3) + v3) + ((v3 + "hello") + ((v3 + "\n\t") + (v3 + v3)))))] = (("\n\t" + v3) + v3)
                    }
                    if (((genFn0() != (v42 << ((1) & 63))) || (!(127 != 127))) && (!(false || true))) {
                        var v64 uint16 = 593
                        var v65 int64 = (v53 ^ (-(v53)))
                        _ = v64
                        _ = v65
                    } else {
                        var v66 int32 = 1
                        var v67 uint64 = v2
                        _ = v66
                        _ = v67
                    }
                }
                v3 = (v3 + v3)
                v53 -= -9223372036854775808
                switch (fi56 << ((960) & 63)) {
                case 9223372036854775807:
                    v68 := GenS6{F0: 32767}
                    _ = v68
                case -9223372036854775808:
                    switch fi55 {
                    case 128:
                        if len(v4) > 0 {
                            v4[0] = 9223372036854775807
                        }
                        v3 = v3
                    default:
                        v53 += (v53 ^ v53)
                        if true {
                            if false {
                                var v69 uint = 1
                                if (int16(v2) > int16(v42)) {
                                    if (true && ((((0 < 877) || false) && (-32768 <= 32767)) && false)) {
                                        var v70 bool = false
                                        v42 = v42
                                        var v71 int32 = 598
                                        _ = v70
                                        _ = v71
                                    }
                                } else {
                                    v1 = v1
                                    v69 -= (v69 % ((v69) | 1))
                                    var v72 uint32 = 2147483648
                                    var v73 uint32 = genFn0()
                                    if len(v4) > 0 {
                                        v4[0] = fi56
                                    }
                                    _ = v72
                                    _ = v73
                                }
                                if (!false) {
                                    v5["foo bar"] = "\n\t"
                                    if true {
                                        var v74 int32 = -2147483648
                                        v3 = (v3 + ((v3 + "\n\t") + "hello"))
                                        if false {
                                            v5["hello"] = (v3 + "")
                                            v54 = 1.5
                                        } else {
                                            v75 := GenS7{F0: 0, F1: int8(v1)}
                                            if true {
                                                fmt.Println(v69, v42, v3)
                                                switch (v3 + v3) {
                                                case "hello":
                                                    if ((-2.25 != 0.1) || ((v53 + 271) <= v53)) {
                                                        var v76 int32 = int32(fi55)
                                                        _ = v76
                                                    } else {
                                                        v77 := GenS4{F0: int(v74), F1: 0}
                                                        var v78 uint32 = genFn0()
                                                        _ = v77
                                                        _ = v78
                                                    }
                                                case "foo bar":
                                                    var v79 int32 = v74
                                                    _ = v79
                                                case "":
                                                    fmt.Println(v2)
                                                    v5["\n\t"] = (("foo bar" + v3) + v3)
                                                    fmt.Println(fi56, v1)
                                                    fmt.Println(v69, v74, v1)
                                                }
                                                if true {
                                                    var v80 int64 = v53
                                                    switch "hello" {
                                                    case "foo bar":
                                                        switch (("foo bar" + v3) + v3) {
                                                        case "hello":
                                                            if true {
                                                                var v81 int64 = (v80 | (v80 << (((v69 | v69)) & 63)))
                                                                var v82 uint8 = (^(v1))
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                                _ = v81
                                                                _ = v82
                                                            } else {
                                                                v1 += 37
                                                                var v83 uint8 = (v1 / ((v1) | 1))
                                                                if false {
                                                                    v4 = append(v4, fi55)
                                                                    v84 := GenS1{F0: (v69 << (((v69 << ((v69) & 63))) & 63))}
                                                                    var v85 int8 = 0
                                                                    _ = v84
                                                                    _ = v85
                                                                }
                                                                fmt.Println(fi56)
                                                                if (true || true) {
                                                                    var v86 int8 = 100
                                                                    if (!((-(fi55)) < -463)) {
                                                                        fmt.Println(v74, v83)
                                                                        v2 += genFn1(v86, 407)
                                                                        v80 = 434
                                                                    } else {
                                                                        var v87 uint64 = v2
                                                                        var v88 float32 = (0.1 - 0.0)
                                                                        v89 := GenS6{F0: 1}
                                                                        if false {
                                                                            var v90 uint = (v69 ^ (v69 << (((v69 / (((v69 / (((v69 ^ (v69 >> ((18446744073709551615) & 63)))) | 1))) | 1))) & 63)))
                                                                            v91 := GenS8{F0: uint16(v87)}
                                                                            if false {
                                                                                v74 = 1
                                                                            }
                                                                            _ = v90
                                                                            _ = v91
                                                                        } else {
                                                                            _ = len(v5)
                                                                        }
                                                                        _ = len(v4)
                                                                        _ = v87
                                                                        _ = v88
                                                                        _ = v89
                                                                    }
                                                                    switch (v3 + v3) {
                                                                    case "foo bar":
                                                                        v3 = (v3 + v3)
                                                                    case "\n\t":
                                                                        var v92 string = v3
                                                                        _ = v92
                                                                    case "hello":
                                                                        var v93 int32 = v74
                                                                        _ = v93
                                                                    default:
                                                                        v2 += v2
                                                                    }
                                                                    _ = v86
                                                                }
                                                                _ = v83
                                                            }
                                                        case "a":
                                                            _ = len(v4)
                                                        }
                                                    default:
                                                        var v94 int16 = int16(v74)
                                                        _ = v94
                                                    }
                                                    v2 = 400
                                                    v95 := GenS4{F0: int(v80), F1: 128}
                                                    var v96 int16 = 1
                                                    _ = v80
                                                    _ = v95
                                                    _ = v96
                                                } else {
                                                    switch ("foo bar" + v3) {
                                                    case "\n\t":
                                                        v42 += v42
                                                    case "foo bar":
                                                        v97 := GenS9{F0: 0}
                                                        _ = v97
                                                    case "a":
                                                        if true {
                                                            var v98 int16 = -982
                                                            v99 := GenS1{F0: (v69 ^ v69)}
                                                            _ = v98
                                                            _ = v99
                                                        } else {
                                                            var v100 uint = 699
                                                            if (1000 > -415) {
                                                                v100 -= (v69 - v69)
                                                                v69 -= v69
                                                            }
                                                            _ = v100
                                                        }
                                                    }
                                                }
                                            }
                                            if (!(false && false)) {
                                                var v101 int = fi55
                                                v102 := GenS5{F0: (32767 != 32767), F1: v69}
                                                _ = v101
                                                _ = v102
                                            }
                                            _ = v75
                                        }
                                        _ = v74
                                    } else {
                                        var v103 int8 = 125
                                        v104 := GenS1{F0: (v69 % ((255) | 1))}
                                        v4 = append(v4, (fi55 & -9223372036854775808))
                                        _ = v103
                                        _ = v104
                                    }
                                } else {
                                    var v105 int16 = 0
                                    _ = v105
                                }
                                _ = v69
                            } else {
                                v42 -= 0
                            }
                        }
                        v53 += v53
                    }
                }
                _ = fi56
            }
            for fi106 := 0; fi106 < 17; fi106++ {
                switch fi106 {
                case 2147483647:
                    switch v3 {
                    case "foo bar":
                        if true {
                            switch 32768 {
                            case -231:
                                var v107 int64 = 0
                                _ = v107
                            case 42:
                                if len(v4) > 0 {
                                    v4[0] = 4294967295
                                }
                            case -526:
                                v1 = (v1 << ((1) & 63))
                            default:
                                var v108 uint32 = (^(v42))
                                _ = v108
                            }
                            v109 := GenS1{F0: 1}
                            _ = v109
                        }
                    case "\n\t":
                        var v110 int64 = (^(v53))
                        _ = v110
                    default:
                        if true {
                            v42 -= (v42 & (v42 % ((v42) | 1)))
                            fmt.Println(v2, fi55)
                        } else {
                            v54 = (v54 - -2.25)
                            var v111 uint64 = uint64(v42)
                            var v112 int64 = (-(v53))
                            _ = v111
                            _ = v112
                        }
                    }
                case 128:
                    v4 = append(v4, (fi55 | (fi106 * 9223372036854775807)))
                }
                _ = fi106
            }
            _ = fi55
        }
    case "a":
        switch -690 {
        case 1:
            for fi113 := 0; fi113 < 16; fi113++ {
                if (v42 >= (v42 >> ((18446744073709551615) & 63))) {
                    v53 = v53
                    var v114 uint8 = v1
                    for fi115 := 0; fi115 < 17; fi115++ {
                        if true {
                            var v116 uint64 = (^(v2))
                            var v117 int32 = -494
                            _ = v116
                            _ = v117
                        }
                        _ = fi115
                    }
                    _ = v114
                }
                v118 := GenS4{F0: (fi113 - 42), F1: fi113}
                _ = fi113
                _ = v118
            }
        case 42:
            var v119 int = int(v53)
            _ = v119
        default:
            v120 := GenS10{F0: -2147483648}
            _ = v120
        }
    }
    if (0.1 == 1.5) {
        v42 -= 42
        switch -71 {
        case -904:
            var v121 int = -9223372036854775808
            _ = v121
        default:
            fmt.Println(v2, v1)
        }
        if false {
            fmt.Println(v53, v1)
        } else {
            fmt.Println(v54, v3, v1)
        }
    }
    v54 += v54
    v53 -= (v53 % ((v53) | 1))
    if true {
        for fi122 := 0; fi122 < 4; fi122++ {
            for fi123 := 0; fi123 < 6; fi123++ {
                if true {
                    v124 := GenS11{F0: fi122}
                    var v125 uint8 = (v1 ^ 2)
                    v126 := GenS2{F0: 9223372036854775807, F1: 1}
                    v127 := v5[""]
                    _ = v124
                    _ = v125
                    _ = v126
                    _ = v127
                }
                var v128 int16 = 1
                if (fi122 <= (fi123 >> ((4294967295) & 63))) {
                    var v129 uint64 = (v2 / ((421) | 1))
                    if true {
                        var v130 int16 = 42
                        var v131 int16 = (v130 ^ v128)
                        var v132 uint8 = v1
                        _ = v130
                        _ = v131
                        _ = v132
                    } else {
                        v2 += (^(v2))
                        v42 = (^(v42))
                        var v133 bool = (true && false)
                        _ = v133
                    }
                    v134 := GenS7{F0: 42, F1: -128}
                    _ = v129
                    _ = v134
                }
                _ = fi123
                _ = v128
            }
            v53 = (v53 * (-(v53)))
            v54 += v54
            _ = fi122
        }
        v1 -= 127
        var v135 int = -920
        _ = v135
    }
    for fi136 := 0; fi136 < 11; fi136++ {
        switch v3 {
        case "hello":
            switch fi136 {
            case 128:
                if false {
                    v2 = v2
                    for fi137 := 0; fi137 < 9; fi137++ {
                        if (0.1 != (v54 - (v54 / v54))) {
                            v138 := GenS7{F0: -32768, F1: 38}
                            _ = v138
                        } else {
                            switch v3 {
                            case "":
                                switch fi136 {
                                case -2147483648:
                                    var v139 string = v3
                                    _ = v139
                                case 128:
                                    if (!false) {
                                        fmt.Println(fi137, v2, v3)
                                        switch (fi137 * (fi136 << ((1) & 63))) {
                                        case 53:
                                            _ = len(v4)
                                        case 1:
                                            v53 = 1
                                        default:
                                            if false {
                                                fmt.Println(fi137, v3, fi136)
                                                var v140 float32 = float32(v54)
                                                _ = v140
                                            } else {
                                                if false {
                                                    v3 = "\n\t"
                                                    var v141 string = "\n\t"
                                                    _ = v141
                                                }
                                            }
                                        }
                                    }
                                case 127:
                                    fmt.Println(fi137, fi136)
                                }
                            default:
                                fmt.Println(fi137, v2, fi136)
                            }
                            v142 := GenS12{F0: -624}
                            v54 = v54
                            v54 -= v54
                            if true {
                                v143 := GenS6{F0: 42}
                                _ = v143
                            } else {
                                fmt.Println(v42, fi136, fi137)
                            }
                            if (genFn1(100, 255) == uint64(v42)) {
                                if true {
                                    fmt.Println(v42, v53, v54)
                                }
                                v53 = (v53 >> ((992) & 63))
                                var v144 int16 = 1
                                _ = v144
                            } else {
                                if false {
                                    var v145 uint64 = v2
                                    _ = v145
                                }
                                if true {
                                    if ((18446744073709551615 > v2) || true) {
                                        if false {
                                            v146 := v5[v3]
                                            var v147 uint64 = (^(v2))
                                            var v148 int32 = -1
                                            _ = v146
                                            _ = v147
                                            _ = v148
                                        } else {
                                            fmt.Println(v54, v2)
                                            var v149 float64 = ((v54 - v54) + v54)
                                            fmt.Println(v42, v3)
                                            _ = v149
                                        }
                                    }
                                    v5["foo bar"] = (v3 + (((v3 + v3) + (v3 + v3)) + v3))
                                    v5[v3] = "a"
                                    v4 = append(v4, fi137)
                                    var v150 int64 = v53
                                    _ = v150
                                }
                                if true {
                                    switch (fi136 | fi136) {
                                    case 0:
                                        var v151 int = (fi136 - fi136)
                                        _ = v151
                                    case 32767:
                                        if (133 > (v1 + (v1 << ((381) & 63)))) {
                                            switch (v3 + (v3 + v3)) {
                                            case "hello":
                                                var v152 float64 = 0.1
                                                _ = v152
                                            case "\n\t":
                                                switch ((v3 + ("" + v3)) + "\n\t") {
                                                case "foo bar":
                                                    switch v3 {
                                                    case "a":
                                                        switch int(v1) {
                                                        case 4294967295:
                                                            var v153 int8 = -1
                                                            _ = v153
                                                        case 593:
                                                            fmt.Println(v2)
                                                        }
                                                    case "foo bar":
                                                        if (true || (0.1 != 0.0)) {
                                                            v5[""] = "a"
                                                            v5["foo bar"] = v3
                                                            var v154 uint16 = 1
                                                            if true {
                                                                v155 := GenS13{F0: 15, F1: 1, F2: 42}
                                                                if (fi136 <= 128) {
                                                                    switch -1 {
                                                                    case 1:
                                                                        var v156 uint8 = 199
                                                                        _ = v156
                                                                    }
                                                                    v1 += 123
                                                                }
                                                                var v157 uint64 = 1
                                                                _ = v155
                                                                _ = v157
                                                            }
                                                            _ = v154
                                                        }
                                                    case "hello":
                                                        delete(v5, "\n\t")
                                                    }
                                                    v1 += (^(v1))
                                                    v158 := GenS14{F0: 65535}
                                                    _ = v158
                                                case "":
                                                    var v159 string = v3
                                                    _ = v159
                                                default:
                                                    if false {
                                                        switch fi136 {
                                                        case 934:
                                                            v3 = v3
                                                            fmt.Println(v1, v53)
                                                            fmt.Println(v3, v54, v54)
                                                        case 32767:
                                                            var v160 float64 = 1.5
                                                            _ = v160
                                                        default:
                                                            var v161 float32 = 1e10
                                                            _ = v161
                                                        }
                                                        v53 += v53
                                                        switch fi137 {
                                                        case 319:
                                                            v162 := GenS15{F0: (v1 - (v1 + 89))}
                                                            _ = v162
                                                        case 4294967295:
                                                            var v163 bool = (641 > 213)
                                                            _ = v163
                                                        case 32767:
                                                            switch "hello" {
                                                            case "hello":
                                                                v2 = v2
                                                            case "a":
                                                                v164 := v5["hello"]
                                                                v3 = ""
                                                                _ = v164
                                                            }
                                                        }
                                                    } else {
                                                        var v165 float64 = -2.25
                                                        _ = v165
                                                    }
                                                }
                                            case "a":
                                                var v166 int64 = v53
                                                _ = v166
                                            }
                                            v3 = "a"
                                        } else {
                                            var v167 int = fi137
                                            if len(v4) > 0 {
                                                v4[0] = (v167 ^ (-(fi137)))
                                            }
                                            _ = v167
                                        }
                                    case 128:
                                        v1 -= v1
                                    default:
                                        v2 += genFn1(1, 2)
                                    }
                                }
                                _ = len(v4)
                                if (42 != 828) {
                                    v1 = (^(v1))
                                } else {
                                    v54 = ((((v54 + (v54 - v54)) * (float64(((float32(float64((1.5 - 1e10))) - -2.25) + -2.25)) * (v54 - v54))) / v54) + v54)
                                }
                                v4 = append(v4, (fi137 * fi136))
                            }
                            _ = v142
                        }
                        if (42 < 127) {
                            var v168 uint16 = 362
                            v169 := GenS10{F0: 0}
                            fmt.Println(v42, v53, v2)
                            if (uint8(v2) != (^(v1))) {
                                var v170 int8 = 127
                                var v171 uint32 = (v42 / (((v42 ^ 2147483647)) | 1))
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v170
                                _ = v171
                            }
                            var v172 float32 = (3.14159 - 1.5)
                            v168 += (v168 >> ((18446744073709551615) & 63))
                            var v173 uint8 = (v1 & v1)
                            _ = v168
                            _ = v169
                            _ = v172
                            _ = v173
                        }
                        fmt.Println(v53, fi136, v53)
                        fmt.Println(v42)
                        _ = fi137
                    }
                }
            case 32768:
                v174 := GenS16{F0: -1}
                _ = v174
            }
        case "\n\t":
            _ = len(v4)
        case "":
            _ = len(v5)
        default:
            var v175 uint64 = 42
            _ = v175
        }
        _ = fi136
    }
    fmt.Println(v1, v1)
    v176 := GenS1{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v176
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v53
    _ = v54
}
