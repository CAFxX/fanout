package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 int; F1 bool; F2 uint16 }
type GenS4 struct { F0 int; F1 uint16; F2 int16 }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS7 struct { F0 int16; F1 int }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 bool; F1 int64; F2 uint }
type GenS10 struct { F0 int64; F1 int8 }
type GenS11 struct { F0 uint32; F1 uint16 }
type GenS12 struct { F0 int; F1 string; F2 string }
type GenS13 struct { F0 uint; F1 int16 }
type GenS14 struct { F0 int32; F1 int }
type GenS15 struct { F0 uint16; F1 uint64 }
type GenS16 struct { F0 uint64; F1 bool }
func genFn0(p0 int16) int64 {
    var v4 float64 = float64(0.1)
    var v5 bool = true
    if v5 {
        var v6 float64 = 1e10
        if true {
            v7 := GenS1{F0: -2147483648, F1: 0}
            v8 := GenS1{F0: -1, F1: 411}
            p0 = (p0 + 0)
            _ = v7
            _ = v8
        } else {
            for fi9 := 0; fi9 < 7; fi9++ {
                var v10 float32 = float32(v6)
                _ = fi9
                _ = v10
            }
            v5 = v5
        }
        _ = v6
    } else {
        if v5 {
            v11 := []int{128}
            _ = v11
        } else {
            v4 = -2.25
            var v12 int64 = 42
            _ = v12
        }
        var v13 float32 = 1.5
        if (v5 && v5) {
            v13 = 3.14159
        }
        _ = v13
    }
    if v5 {
        var v14 string = "\n\t"
        var v15 uint8 = 2
        _ = v14
        _ = v15
    }
    _ = p0
    _ = v4
    _ = v5
    return 0
}

func main() {
    var v1 int16 = 32767
    v2 := []uint32{0}
    v3 := make(map[uint32]uint32)
    switch "foo bar" {
    case "a":
        if false {
            v16 := GenS1{F0: 32767, F1: 1}
            _ = v16
        }
    case "foo bar":
        switch 128 {
        case 1:
            v17 := GenS1{F0: 437, F1: 718}
            _ = v17
        case -2147483648:
            v1 = (v1 << ((42) & 63))
        case 127:
            var v18 uint32 = 594
            _ = v18
        default:
            v1 += v1
        }
    case "":
        v19 := GenS1{F0: -699, F1: 965}
        _ = v19
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        if (!(18446744073709551615 > 42)) {
            var v21 string = "a"
            var v22 int16 = (v1 | 32767)
            _ = v21
            _ = v22
        } else {
            var v23 uint = 1
            _ = v23
        }
        var v24 uint = 303
        v24 = (v24 + v24)
        _ = fi20
        _ = v24
    }
    for fi25 := 0; fi25 < 1; fi25++ {
        fmt.Println(fi25, v1)
        var v26 int = fi25
        for fi27 := 0; fi27 < 3; fi27++ {
            v1 -= -182
            _ = fi27
        }
        _ = fi25
        _ = v26
    }
    var v28 uint16 = 42
    switch (("" + "a") + "\n\t") {
    case "hello":
        var v29 uint16 = (v28 | v28)
        _ = v29
    case "a":
        v28 -= 2
    case "\n\t":
        var v30 uint16 = v28
        _ = v30
    }
    if true {
        v31 := GenS2{F0: 42, F1: 427}
        v32 := GenS3{F0: -9223372036854775808, F1: true, F2: 515}
        if ((v1 / (((^(v1))) | 1)) >= v1) {
            if false {
                v33 := GenS4{F0: 128, F1: (v28 >> ((65535) & 63)), F2: (-(v1))}
                var v34 uint = 65535
                v35 := GenS2{F0: -1, F1: -1}
                _ = v33
                _ = v34
                _ = v35
            }
        } else {
            v1 += (v1 ^ (v1 * v1))
            v36 := GenS3{F0: 42, F1: (0 >= genFn0(-299)), F2: 32767}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v36
        }
        _ = v31
        _ = v32
    } else {
        var v37 uint64 = 304
        if (175 < 255) {
            fmt.Println(v1)
        } else {
            if false {
                v1 += 949
                for fi38 := 0; fi38 < 2; fi38++ {
                    for fi39 := 0; fi39 < 2; fi39++ {
                        switch fi39 {
                        case 42:
                            switch "a" {
                            case "foo bar":
                                var v40 uint16 = v28
                                _ = v40
                            default:
                                if (-78 <= 0) {
                                    fmt.Println(v1, v28)
                                }
                                var v41 float32 = 0.1
                                _ = v41
                            }
                        case 32768:
                            delete(v3, 31)
                            var v42 uint8 = 255
                            v37 = v37
                            switch "\n\t" {
                            case "hello":
                                v3[420] = 2147483647
                            case "a":
                                v42 = v42
                            case "":
                                if true {
                                    switch fi39 {
                                    case -2147483648:
                                        v28 = v28
                                        switch "" {
                                        case "\n\t":
                                            switch (-(fi39)) {
                                            case 32768:
                                                v43 := GenS5{F0: (v37 >> ((65535) & 63)), F1: -1}
                                                _ = v43
                                            case 445:
                                                v3[218] = 4294967295
                                            }
                                        case "hello":
                                            var v44 float32 = 1.5
                                            _ = v44
                                        case "a":
                                            v2 = append(v2, 374)
                                        }
                                    default:
                                        v42 += 212
                                    }
                                    v45 := GenS6{F0: (v28 << ((2) & 63)), F1: (^(v42)), F2: genFn0(521)}
                                    var v46 uint32 = 2147483647
                                    _ = v45
                                    _ = v46
                                } else {
                                    switch (fi38 << ((2) & 63)) {
                                    case 768:
                                        var v47 float64 = -2.25
                                        _ = v47
                                    case 1:
                                        v48 := v3[127]
                                        _ = v48
                                    case -9223372036854775808:
                                        _ = len(v2)
                                    }
                                    if false {
                                        var v49 bool = true
                                        _ = v49
                                    }
                                    _ = len(v3)
                                }
                            default:
                                if ((2 == 42) || false) {
                                    var v50 int8 = -128
                                    var v51 float32 = 1.5
                                    _ = len(v3)
                                    _ = v50
                                    _ = v51
                                }
                            }
                            fmt.Println(v28, v42, v28)
                            _ = v42
                        default:
                            _ = len(v2)
                        }
                        fmt.Println(v1)
                        if false {
                            var v52 float64 = float64((1e10 * ((1.5 / (0.1 + 1e10)) * 1e10)))
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v53 uint64 = 90
                            _ = v52
                            _ = v53
                        } else {
                            var v54 uint = 255
                            var v55 uint8 = uint8(v28)
                            _ = v54
                            _ = v55
                        }
                        if true {
                            fmt.Println(v1, v1, fi38)
                        } else {
                            if true {
                                v1 = v1
                                v37 = v37
                                v28 = (v28 - 1)
                            } else {
                                v56 := GenS7{F0: v1, F1: (fi38 / ((fi38) | 1))}
                                v37 += (v37 | v37)
                                _ = v56
                            }
                        }
                        if false {
                            v28 = 593
                        }
                        _ = fi39
                    }
                    _ = fi38
                }
                switch -808 {
                case 127:
                    switch 2147483647 {
                    case -577:
                        var v57 int64 = 42
                        _ = v57
                    case 2147483647:
                        for fi58 := 0; fi58 < 14; fi58++ {
                            var v59 uint16 = 65535
                            var v60 uint = 2
                            v3[822] = 42
                            _ = fi58
                            _ = v59
                            _ = v60
                        }
                    case -144:
                        var v61 uint = 2
                        _ = v61
                    }
                default:
                    var v62 int64 = 9223372036854775807
                    _ = v62
                }
            } else {
                v37 -= 840
                for fi63 := 0; fi63 < 18; fi63++ {
                    switch "a" {
                    case "a":
                        v28 = (v28 & v28)
                    case "foo bar":
                        for fi64 := 0; fi64 < 4; fi64++ {
                            _ = len(v2)
                            var v65 int = (fi64 | fi63)
                            if (0.0 >= 0.1) {
                                switch ("" + "\n\t") {
                                case "":
                                    if (true || true) {
                                        v3[391] = 2147483648
                                        v28 = v28
                                    }
                                case "a":
                                    var v66 int16 = (v1 / ((v1) | 1))
                                    _ = v66
                                case "\n\t":
                                    v65 = v65
                                default:
                                    var v67 uint = 2
                                    _ = v67
                                }
                                delete(v3, 227)
                                v68 := GenS8{F0: v28}
                                fmt.Println(v28)
                                var v69 uint8 = 193
                                _ = v68
                                _ = v69
                            } else {
                                fmt.Println(fi64)
                                v65 += fi64
                                var v70 bool = true
                                v71 := GenS6{F0: v28, F1: 1, F2: 9223372036854775807}
                                _ = v70
                                _ = v71
                            }
                            var v72 uint32 = 1
                            if (19 >= uint8(v72)) {
                                switch 32767 {
                                case 42:
                                    fmt.Println(v37, v65)
                                case 143:
                                    switch 127 {
                                    case 128:
                                        _ = len(v3)
                                        v73 := GenS9{F0: (!(18446744073709551615 == 0)), F1: 640, F2: 255}
                                        _ = v73
                                    }
                                case -818:
                                    if false {
                                        v1 += v1
                                        var v74 float32 = 3.14159
                                        v75 := GenS10{F0: genFn0((v1 / ((v1) | 1))), F1: 0}
                                        _ = v74
                                        _ = v75
                                    } else {
                                        switch fi63 {
                                        case 9223372036854775807:
                                            v76 := GenS1{F0: fi63, F1: 42}
                                            _ = v76
                                        case 32768:
                                            v77 := GenS3{F0: fi64, F1: false, F2: (v28 * v28)}
                                            _ = v77
                                        default:
                                            v65 = (^(fi64))
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                }
                            } else {
                                switch 2147483647 {
                                case 963:
                                    v78 := GenS5{F0: v37, F1: 519}
                                    _ = v78
                                }
                                if (!false) {
                                    if true {
                                        v79 := GenS4{F0: 42, F1: v28, F2: 1}
                                        var v80 int64 = -1
                                        _ = v79
                                        _ = v80
                                    }
                                    v65 = (fi63 >> ((18446744073709551615) & 63))
                                    v65 -= (fi64 >> ((uint(v72)) & 63))
                                }
                            }
                            v72 = 1
                            _ = fi64
                            _ = v65
                            _ = v72
                        }
                    default:
                        _ = len(v2)
                    }
                    v81 := v3[2147483648]
                    _ = fi63
                    _ = v81
                }
            }
            v1 = v1
        }
        _ = v37
    }
    _ = len(v3)
    v82 := GenS11{F0: 1, F1: (v28 % ((650) | 1))}
    switch "\n\t" {
    case "\n\t":
        v28 = v28
    case "hello":
        if true {
            if true {
                if false {
                    for fi83 := 0; fi83 < 4; fi83++ {
                        _ = len(v3)
                        if len(v2) > 0 {
                            v2[0] = 2147483648
                        }
                        _ = fi83
                    }
                    v1 += (v1 ^ v1)
                    v84 := GenS11{F0: 710, F1: v28}
                    _ = v84
                }
                var v85 uint64 = 0
                if true {
                    var v86 float32 = ((float32(1.5) / (3.14159 / 3.14159)) - (3.14159 / ((1.5 + 0.1) * 3.14159)))
                    _ = v86
                }
                _ = v85
            } else {
                if (v28 <= v28) {
                    var v87 int = 4294967295
                    fmt.Println(v1)
                    _ = v87
                }
                var v88 int16 = (v1 & (v1 ^ (v1 | v1)))
                if len(v2) > 0 {
                    v2[0] = 4294967295
                }
                _ = v88
            }
            v89 := v3[48]
            for fi90 := 0; fi90 < 1; fi90++ {
                v91 := GenS12{F0: (fi90 << ((763) & 63)), F1: "", F2: "foo bar"}
                if (1 > (v1 | v1)) {
                    if len(v2) > 0 {
                        v2[0] = (v89 - v89)
                    }
                    v89 -= 2147483648
                    var v92 float64 = 1e10
                    _ = v92
                } else {
                    delete(v3, v89)
                    var v93 uint16 = uint16(v89)
                    v2 = append(v2, 42)
                    _ = v93
                }
                if (true && false) {
                    switch (("a" + ("\n\t" + "hello")) + "") {
                    case "a":
                        for fi94 := 0; fi94 < 1; fi94++ {
                            fmt.Println(v28, fi94, v28)
                            fmt.Println(fi90, fi90)
                            var v95 float32 = (float32(float64((0.0 + (3.14159 * 0.0)))) * 0.0)
                            _ = fi94
                            _ = v95
                        }
                    case "\n\t":
                        v28 += (v28 >> ((42) & 63))
                    case "foo bar":
                        v1 += v1
                    default:
                        v1 += v1
                    }
                }
                _ = fi90
                _ = v91
            }
            _ = v89
        }
    }
    if true {
        var v96 uint8 = 127
        if true {
            if false {
                switch 2147483647 {
                case 635:
                    v96 += v96
                default:
                    for fi97 := 0; fi97 < 19; fi97++ {
                        for fi98 := 0; fi98 < 20; fi98++ {
                            if (1 > -9223372036854775808) {
                                v1 = v1
                                if false {
                                    var v99 int8 = 127
                                    var v100 int32 = 2147483647
                                    _ = v100
                                    _ = v99
                                } else {
                                    var v101 float64 = 1.5
                                    var v102 float64 = v101
                                    if true {
                                        v103 := GenS13{F0: 210, F1: v1}
                                        _ = v103
                                    }
                                    _ = v101
                                    _ = v102
                                }
                            }
                            v104 := GenS5{F0: 2, F1: -497}
                            var v105 uint32 = 42
                            _ = fi98
                            _ = v104
                            _ = v105
                        }
                        _ = fi97
                    }
                }
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            if false {
                for fi106 := 0; fi106 < 3; fi106++ {
                    v107 := GenS13{F0: 0, F1: (-(v1))}
                    _ = fi106
                    _ = v107
                }
            } else {
                fmt.Println(v28, v28, v28)
                var v108 string = "a"
                _ = v108
            }
            for fi109 := 0; fi109 < 12; fi109++ {
                delete(v3, 2147483648)
                if false {
                    for fi110 := 0; fi110 < 18; fi110++ {
                        v2 = append(v2, 2147483647)
                        fmt.Println(v96)
                        _ = fi110
                    }
                    switch "foo bar" {
                    case "":
                        var v111 int64 = 1
                        _ = v111
                    case "a":
                        v112 := GenS14{F0: -234, F1: 127}
                        _ = v112
                    case "foo bar":
                        v113 := GenS15{F0: 0, F1: 1}
                        _ = v113
                    default:
                        v114 := v3[2]
                        _ = v114
                    }
                } else {
                    for fi115 := 0; fi115 < 2; fi115++ {
                        var v116 float32 = 1e10
                        v117 := GenS4{F0: 925, F1: 0, F2: (-(v1))}
                        var v118 string = "foo bar"
                        _ = fi115
                        _ = v116
                        _ = v117
                        _ = v118
                    }
                    if ((!(uint8(v28) < (v96 << ((0) & 63)))) && true) {
                        switch fi109 {
                        case 1:
                            v119 := GenS7{F0: (v1 / ((191) | 1)), F1: (fi109 >> ((4294967295) & 63))}
                            _ = v119
                        case 967:
                            v3[2147483647] = 662
                        }
                        if ((((2147483647 > -594) || (int8(fi109) == 127)) && (true && false)) || true) {
                            v120 := GenS9{F0: (541 == 0), F1: -1, F2: 4294967295}
                            _ = v120
                        }
                        for fi121 := 0; fi121 < 4; fi121++ {
                            var v122 uint64 = 1
                            _ = fi121
                            _ = v122
                        }
                    }
                }
                if true {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                _ = fi109
            }
        }
        if false {
            var v123 string = "hello"
            _ = v123
        } else {
            switch 2147483647 {
            case -278:
                var v124 int16 = v1
                _ = v124
            }
            switch "a" {
            case "foo bar":
                if ((true || (true && false)) || true) {
                    v28 = (v28 % (((v28 & v28)) | 1))
                    if false {
                        for fi125 := 0; fi125 < 10; fi125++ {
                            var v126 uint8 = uint8(v28)
                            var v127 float64 = (1.5 + 0.0)
                            _ = fi125
                            _ = v126
                            _ = v127
                        }
                        var v128 int16 = v1
                        for fi129 := 0; fi129 < 17; fi129++ {
                            var v130 int64 = 1
                            v131 := v3[839]
                            _ = fi129
                            _ = v130
                            _ = v131
                        }
                        _ = v128
                    }
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v132 uint64 = 18446744073709551615
                    _ = v132
                }
            default:
                var v133 float64 = 0.0
                _ = v133
            }
            switch int(v28) {
            case 0:
                v134 := GenS3{F0: 32768, F1: (0 >= 0), F2: 245}
                _ = v134
            case -2147483648:
                v135 := v3[2147483648]
                _ = v135
            case 32768:
                v1 += (v1 - v1)
            default:
                for fi136 := 0; fi136 < 13; fi136++ {
                    v137 := GenS9{F0: false, F1: 0, F2: 65535}
                    var v138 uint32 = 2
                    _ = fi136
                    _ = v137
                    _ = v138
                }
            }
        }
        _ = v96
    }
    v28 += 529
    var v139 int64 = 859
    if ((false && ((v139 / ((-1) | 1)) != (v139 | (v139 * 1)))) || ((-2147483648 < 2147483647) && false)) {
        for fi140 := 0; fi140 < 1; fi140++ {
            for fi141 := 0; fi141 < 12; fi141++ {
                if (-2.25 != 3.14159) {
                    v142 := GenS16{F0: 181, F1: false}
                    _ = v142
                } else {
                    if (!true) {
                        v1 -= (-(v1))
                    }
                    _ = len(v3)
                }
                fmt.Println(v1, fi141)
                _ = fi141
            }
            fmt.Println(v1)
            switch fi140 {
            case 4294967295:
                var v143 int16 = v1
                _ = v143
            case -2147483648:
                var v144 uint16 = v28
                _ = v144
            default:
                var v145 int = fi140
                _ = v145
            }
            _ = fi140
        }
        for fi146 := 0; fi146 < 6; fi146++ {
            if ((!(!false)) || (842 <= (v28 * (v28 + (v28 * v28))))) {
                delete(v3, 1)
                var v147 int64 = genFn0(1000)
                v148 := GenS5{F0: 496, F1: 42}
                _ = v147
                _ = v148
            } else {
                if false {
                    var v149 int8 = 1
                    if ((!true) || (1.5 != 0.0)) {
                        var v150 uint16 = 42
                        var v151 bool = true
                        switch (fi146 % ((fi146) | 1)) {
                        case 42:
                            switch "" {
                            case "":
                                v152 := GenS7{F0: v1, F1: fi146}
                                _ = v152
                            case "foo bar":
                                for fi153 := 0; fi153 < 20; fi153++ {
                                    fmt.Println(v151, v139, v1)
                                    _ = fi153
                                }
                            }
                        case 32767:
                            for fi154 := 0; fi154 < 16; fi154++ {
                                var v155 int = fi146
                                _ = len(v3)
                                _ = fi154
                                _ = v155
                            }
                        case 875:
                            var v156 uint = 2
                            _ = v156
                        }
                        _ = v150
                        _ = v151
                    } else {
                        v1 = (v1 & v1)
                    }
                    _ = v149
                }
            }
            if false {
                v157 := GenS14{F0: -181, F1: int(v1)}
                var v158 uint16 = 204
                _ = v157
                _ = v158
            } else {
                fmt.Println(fi146, v1, fi146)
            }
            switch "hello" {
            case "":
                for fi159 := 0; fi159 < 1; fi159++ {
                    var v160 int = fi146
                    var v161 uint16 = 177
                    _ = fi159
                    _ = v160
                    _ = v161
                }
            case "\n\t":
                v162 := GenS4{F0: int(v1), F1: v28, F2: (v1 * 1000)}
                _ = v162
            }
            _ = fi146
        }
        var v163 int64 = v139
        _ = v163
    }
    fmt.Println(v28, v1)
    switch "\n\t" {
    case "a":
        fmt.Println(v1, v139)
    case "foo bar":
        v139 = (v139 + v139)
    }
    var v164 uint16 = v28
    var v165 uint16 = (^(v28))
    if (!true) {
        if (3.14159 != 1.5) {
            var v166 int32 = 42
            v1 = 42
            switch 0 {
            case -615:
                if ((v165 * v165) == (v28 >> ((255) & 63))) {
                    v164 = v164
                    var v167 uint16 = 489
                    _ = v167
                } else {
                    var v168 float32 = 1.5
                    switch "hello" {
                    case "\n\t":
                        v169 := GenS4{F0: 1, F1: (v164 % ((v164) | 1)), F2: 345}
                        _ = v169
                    }
                    var v170 uint64 = 527
                    _ = v168
                    _ = v170
                }
            case 9223372036854775807:
                v171 := GenS2{F0: -1, F1: 42}
                _ = v171
            default:
                delete(v3, 1)
            }
            _ = v166
        } else {
            if (-2.25 > (0.1 / float64(1e10))) {
                v28 = v164
                v28 = 643
                fmt.Println(v28, v164)
            }
            v1 = (v1 << ((uint(v164)) & 63))
            v172 := GenS11{F0: 1, F1: (^(v28))}
            _ = v172
        }
    } else {
        v1 = (^(v1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v139:", v139)
    fmt.Println("v164:", v164)
    fmt.Println("v165:", v165)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v139
    _ = v164
    _ = v165
    _ = v2
    _ = v28
    _ = v3
    _ = v82
}
