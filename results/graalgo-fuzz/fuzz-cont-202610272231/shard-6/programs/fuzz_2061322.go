package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 uint64; F2 bool }
type GenS3 struct { F0 uint8; F1 int; F2 int }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int; F1 int64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 string; F1 string }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int16; F1 int8 }
type GenS11 struct { F0 string; F1 uint16 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint16; F1 uint16 }
type GenS14 struct { F0 int32; F1 uint8; F2 int }
func genFn0(p0 int, p1 uint16) int8 {
    for fi5 := 0; fi5 < 12; fi5++ {
        fmt.Println(p0, p1, p1)
        _ = fi5
    }
    for fi6 := 0; fi6 < 6; fi6++ {
        p0 = p0
        _ = fi6
    }
    var v7 uint = 255
    p0 = p0
    if true {
        for fi8 := 0; fi8 < 12; fi8++ {
            p0 -= fi8
            _ = fi8
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return -67
}
func genFn1() uint8 {
    var v9 uint8 = 64
    fmt.Println(v9)
    var v10 uint = uint(v9)
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v9)
    case "hello":
        for fi11 := 0; fi11 < 3; fi11++ {
            for fi12 := 0; fi12 < 6; fi12++ {
                v13 := []int{}
                v9 += v9
                _ = fi12
                _ = v13
            }
            for fi14 := 0; fi14 < 4; fi14++ {
                v10 -= v10
                if (1000 == -920) {
                    if true {
                        v10 += v10
                        if ((-2147483648 > 2147483647) || true) {
                            v15 := []int{}
                            if (!(1 < -2147483648)) {
                                if true {
                                    v9 += v9
                                    v9 -= (v9 / ((0) | 1))
                                    var v16 float32 = (0.1 * 3.14159)
                                    _ = v16
                                }
                                fmt.Println(v9, fi14)
                                if (-763 <= 9223372036854775807) {
                                    if (false && true) {
                                        if (!(65535 <= 65535)) {
                                            var v17 int = fi11
                                            v10 += v10
                                            _ = v17
                                        } else {
                                            v18 := GenS1{F0: 1}
                                            v19 := GenS1{F0: -665}
                                            _ = v18
                                            _ = v19
                                        }
                                    } else {
                                        fmt.Println(fi14)
                                        if (((0.1 * 3.14159) / (0.1 / 3.14159)) > -2.25) {
                                            var v20 int8 = -18
                                            v21 := GenS2{F0: -276, F1: 4294967295, F2: (1 == 1)}
                                            _ = v20
                                            _ = v21
                                        }
                                    }
                                }
                            }
                            _ = v15
                        } else {
                            var v22 int8 = 127
                            _ = v22
                        }
                    } else {
                        v23 := GenS1{F0: 1}
                        if false {
                            if (1000 != 42) {
                                v9 -= (^(v9))
                                var v24 uint16 = 1
                                fmt.Println(fi14)
                                _ = v24
                            }
                        }
                        _ = v23
                    }
                    switch fi14 {
                    case 2147483647:
                        var v25 float64 = 0.0
                        _ = v25
                    }
                    switch int(v9) {
                    case -1:
                        v26 := []int{fi11, fi11, fi11}
                        _ = v26
                    case 2147483647:
                        var v27 float32 = (-2.25 - (-2.25 + -2.25))
                        _ = v27
                    case 4294967295:
                        v9 = v9
                    }
                }
                _ = fi14
            }
            _ = fi11
        }
    case "\n\t":
        fmt.Println(v9, v10)
    }
    v28 := GenS2{F0: -9223372036854775808, F1: 42, F2: false}
    _ = v10
    _ = v28
    _ = v9
    return v9
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 int = -1
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 61
    switch v2 {
    case 42:
        var v29 int = v2
        _ = v29
    case 458:
        if true {
            fmt.Println(v1)
        } else {
            v30 := GenS3{F0: v4, F1: 743, F2: 128}
            v1 += 1
            _ = v30
        }
    }
    if false {
        var v31 bool = true
        v32 := GenS4{F0: 166, F1: v31}
        for fi33 := 0; fi33 < 13; fi33++ {
            for fi34 := 0; fi34 < 7; fi34++ {
                if ((-993 <= -9223372036854775808) || true) {
                    v35 := []int{v2, (fi33 * (fi34 ^ 227)), fi33, fi34}
                    _ = v35
                } else {
                    if (3.14159 <= ((-2.25 + 0.1) * 0.1)) {
                        v1 = v1
                        v36 := GenS4{F0: 878, F1: ((!((((v1 >> (((v1 << ((v1) & 63))) & 63)) >= (v1 >> ((65535) & 63))) && v31) && (!(!false)))) && v31)}
                        var v37 float64 = 3.14159
                        _ = v36
                        _ = v37
                    } else {
                        v38 := GenS5{F0: ("a" + "\n\t")}
                        if v31 {
                            v31 = v31
                            v4 += 1
                        }
                        v3[1] = 1
                        _ = v38
                    }
                }
                delete(v3, 42)
                var v39 uint8 = v4
                if (v31 || (-1 != -1)) {
                    v40 := GenS2{F0: 543, F1: 1, F2: (-1 <= 32767)}
                    if v31 {
                        v41 := []int{fi33}
                        var v42 int64 = 9223372036854775807
                        var v43 int32 = 2147483647
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                    var v44 int64 = 1
                    fmt.Println(v44)
                    v45 := []int{}
                    if false {
                        var v46 uint32 = uint32(v39)
                        _ = v46
                    }
                    fmt.Println(fi34)
                    _ = v40
                    _ = v44
                    _ = v45
                } else {
                    switch ("hello" + "\n\t") {
                    case "":
                        v47 := GenS5{F0: "hello"}
                        _ = v47
                    default:
                        if true {
                            v31 = (((1.5 / 3.14159) < 0.0) || false)
                            v4 -= v4
                            v48 := GenS1{F0: 42}
                            var v49 int8 = 0
                            _ = v48
                            _ = v49
                        } else {
                            switch "\n\t" {
                            case "a":
                                fmt.Println(v39, v39, v2)
                            case "\n\t":
                                var v50 int16 = -32768
                                _ = v50
                            default:
                                var v51 int64 = -1
                                _ = v51
                            }
                            if (42 != 622) {
                                var v52 int32 = 1
                                fmt.Println(v4)
                                _ = v52
                            } else {
                                fmt.Println(v4, fi34)
                            }
                            switch "foo bar" {
                            case "\n\t":
                                var v53 int8 = -97
                                _ = v53
                            }
                            var v54 uint64 = 1
                            _ = v54
                        }
                    }
                    v1 -= v1
                }
                _ = fi34
                _ = v39
            }
            _ = fi33
        }
        _ = v31
        _ = v32
    }
    for fi55 := 0; fi55 < 6; fi55++ {
        var v56 bool = false
        var v57 int32 = 145
        for fi58 := 0; fi58 < 12; fi58++ {
            v57 = 0
            _ = fi58
        }
        _ = fi55
        _ = v56
        _ = v57
    }
    if (!(true || (!true))) {
        v2 -= v2
        fmt.Println(v1, v1, v2)
    }
    switch "a" {
    case "\n\t":
        if true {
            if true {
                if (!true) {
                    v1 = 42
                    if (false || (2 == 305)) {
                        switch "" {
                        case "hello":
                            v3[174] = 4294967295
                        }
                        if false {
                            if true {
                                var v59 uint32 = 4294967295
                                switch "foo bar" {
                                case "":
                                    v60 := GenS3{F0: (^(v4)), F1: 2147483647, F2: v2}
                                    _ = v60
                                case "hello":
                                    for fi61 := 0; fi61 < 18; fi61++ {
                                        var v62 int8 = -1
                                        _ = fi61
                                        _ = v62
                                    }
                                }
                                _ = v59
                            } else {
                                v63 := GenS6{F0: (^(v2)), F1: 1}
                                _ = v63
                            }
                            var v64 int32 = 1
                            var v65 float64 = float64(-2.25)
                            _ = v64
                            _ = v65
                        }
                        v4 += 1
                    } else {
                        v1 -= v1
                    }
                }
            } else {
                fmt.Println(v1, v4, v2)
            }
        }
    case "foo bar":
        fmt.Println(v1)
    case "a":
        _ = len(v3)
    }
    switch "" {
    case "foo bar":
        switch ("a" + "foo bar") {
        case "hello":
            v66 := []int{(v2 + (v2 ^ v2)), (v2 | v2)}
            _ = v66
        case "\n\t":
            v67 := GenS7{F0: genFn1()}
            _ = v67
        case "a":
            var v68 uint = (v1 | 4294967295)
            _ = v68
        default:
            v69 := []int{}
            _ = v69
        }
    case "":
        if true {
            v4 -= genFn1()
            v1 -= 0
        } else {
            for fi70 := 0; fi70 < 14; fi70++ {
                switch ((("a" + "a") + ("\n\t" + "\n\t")) + "a") {
                case "\n\t":
                    fmt.Println(v1, v4, v4)
                case "a":
                    for fi71 := 0; fi71 < 14; fi71++ {
                        v2 -= (fi70 & int(v1))
                        var v72 int64 = 873
                        v2 = fi71
                        v73 := GenS7{F0: genFn1()}
                        if true {
                            var v74 uint16 = 989
                            v72 = -645
                            var v75 uint64 = 1
                            if (84 != -1) {
                                if (-128 < -32) {
                                    var v76 uint8 = (v4 * (v4 ^ uint8(v74)))
                                    _ = v76
                                } else {
                                    switch (-(fi71)) {
                                    case -817:
                                        v77 := GenS4{F0: (v74 ^ 65535), F1: (true || true)}
                                        v78 := GenS8{F0: "\n\t", F1: "hello"}
                                        v72 = v72
                                        var v79 int8 = genFn0((v2 + fi70), v74)
                                        _ = v77
                                        _ = v78
                                        _ = v79
                                    case 2147483647:
                                        v80 := []int{fi70, fi70}
                                        _ = v80
                                    case 128:
                                        delete(v3, 1)
                                    default:
                                        var v81 float64 = 1.5
                                        _ = v81
                                    }
                                    if true {
                                        v82 := []int{(^(v2)), fi71, fi71, fi70}
                                        v4 = (^(v4))
                                        fmt.Println(v72, v72)
                                        _ = v82
                                    }
                                }
                                switch "\n\t" {
                                case "":
                                    var v83 int64 = (v72 - 0)
                                    _ = v83
                                }
                                v84 := GenS7{F0: genFn1()}
                                switch "hello" {
                                case "a":
                                    v85 := []int{(fi70 ^ fi71), 32768}
                                    _ = v85
                                case "\n\t":
                                    if ((v4 + (v4 ^ (v4 + (v4 & (^(v4)))))) >= (v4 * 229)) {
                                        v74 = (v74 % ((32767) | 1))
                                        fmt.Println(v4, v74)
                                    } else {
                                        v86 := GenS9{F0: uint32(v75)}
                                        if false {
                                            v87 := []int{(fi71 + -2147483648), (-(fi70))}
                                            if ((v1 / ((v1) | 1)) != (v1 << (((v1 * (^(v1)))) & 63))) {
                                                _ = len(v3)
                                                if len(v87) > 0 {
                                                    v87[0] = (fi71 % (((fi71 + fi71)) | 1))
                                                }
                                                var v88 uint64 = v75
                                                _ = v88
                                            }
                                            v89 := GenS10{F0: -1, F1: -128}
                                            _ = v87
                                            _ = v89
                                        } else {
                                            if false {
                                                var v90 uint16 = v74
                                                _ = v90
                                            } else {
                                                var v91 uint32 = 1
                                                var v92 int64 = (v72 % ((v72) | 1))
                                                _ = v91
                                                _ = v92
                                            }
                                            var v93 uint64 = v75
                                            _ = v93
                                        }
                                        v74 = (v74 / ((v74) | 1))
                                        _ = v86
                                    }
                                case "hello":
                                    v75 += (^(v75))
                                }
                                v4 = 2
                                _ = v84
                            } else {
                                fmt.Println(v75, v72)
                                if true {
                                    var v94 int32 = int32(v74)
                                    fmt.Println(v2, v2)
                                    _ = v94
                                } else {
                                    if (!false) {
                                        var v95 uint32 = 815
                                        v75 += v75
                                        delete(v3, uint32(v4))
                                        _ = v95
                                    }
                                }
                                _ = len(v3)
                                var v96 int8 = -74
                                _ = v96
                            }
                            if ((fi70 | (^(fi70))) > -258) {
                                v97 := GenS11{F0: ("foo bar" + ("\n\t" + "foo bar")), F1: 630}
                                var v98 int64 = (v72 % (((v72 | (v72 & 483))) | 1))
                                var v99 bool = (("\n\t" != "a") && (24 > -1))
                                _ = v97
                                _ = v98
                                _ = v99
                            }
                            _ = v74
                            _ = v75
                        }
                        v100 := GenS10{F0: -1, F1: genFn0(fi70, 32767)}
                        if (!(202 >= 1)) {
                            v101 := GenS3{F0: (v4 >> ((v1) & 63)), F1: (fi70 ^ (fi71 | 32768)), F2: v2}
                            _ = v101
                        }
                        _ = fi71
                        _ = v100
                        _ = v72
                        _ = v73
                    }
                default:
                    if false {
                        if (!false) {
                            if true {
                                v102 := GenS1{F0: 577}
                                v103 := []int{}
                                switch "\n\t" {
                                case "\n\t":
                                    for fi104 := 0; fi104 < 7; fi104++ {
                                        var v105 int16 = 0
                                        _ = fi104
                                        _ = v105
                                    }
                                case "a":
                                    if (false || false) {
                                        for fi106 := 0; fi106 < 11; fi106++ {
                                            if false {
                                                var v107 uint32 = 4294967295
                                                delete(v3, (^(v107)))
                                                _ = v107
                                            }
                                            v108 := GenS8{F0: ("\n\t" + ((("a" + "\n\t") + "\n\t") + "\n\t")), F1: "foo bar"}
                                            v1 = v1
                                            if len(v103) >= 2 {
                                                v103 = v103[1:len(v103)]
                                            }
                                            _ = fi106
                                            _ = v108
                                        }
                                        fmt.Println(v4)
                                    } else {
                                        fmt.Println(v4, v4, v2)
                                        if true {
                                            if len(v103) > 0 {
                                                v103[0] = (^(v2))
                                            }
                                            v109 := GenS5{F0: ("a" + ("hello" + "a"))}
                                            v2 -= 127
                                            _ = v109
                                        } else {
                                            var v110 int8 = -1
                                            switch "\n\t" {
                                            case "foo bar":
                                                v110 = v110
                                            default:
                                                v2 += v2
                                            }
                                            v111 := GenS11{F0: "", F1: 255}
                                            _ = v110
                                            _ = v111
                                        }
                                    }
                                }
                                _ = v102
                                _ = v103
                            }
                            v112 := []int{fi70, fi70, 612, v2}
                            switch v2 {
                            case -9223372036854775808:
                                for fi113 := 0; fi113 < 7; fi113++ {
                                    delete(v3, 0)
                                    if len(v112) > 0 {
                                        v112[0] = (-(v2))
                                    }
                                    _ = len(v3)
                                    v2 += 32768
                                    _ = fi113
                                }
                            }
                            _ = v112
                        } else {
                            var v114 float32 = 1.5
                            _ = v114
                        }
                    }
                }
                _ = fi70
            }
            for fi115 := 0; fi115 < 20; fi115++ {
                if (!true) {
                    v3[2] = 4294967295
                } else {
                    v2 += (fi115 & (^(fi115)))
                    var v116 float32 = -2.25
                    _ = v116
                }
                _ = fi115
            }
            var v117 uint8 = (v4 << ((4294967295) & 63))
            _ = v117
        }
    default:
        v2 = (v2 - (v2 % ((v2) | 1)))
    }
    if true {
        var v118 float64 = -2.25
        var v119 uint = v1
        for fi120 := 0; fi120 < 19; fi120++ {
            v119 += v119
            switch fi120 {
            case -1:
                v121 := []int{fi120, (fi120 - 9223372036854775807), (fi120 >> ((v1) & 63)), v2}
                _ = v121
            case 9223372036854775807:
                var v122 int32 = 96
                _ = v122
            case -9223372036854775808:
                if (!true) {
                    v3[2147483647] = 1
                    v2 -= (v2 >> (((^(v119))) & 63))
                } else {
                    var v123 string = "hello"
                    var v124 uint = (v119 ^ v119)
                    switch "" {
                    case "":
                        if (42 == -128) {
                            for fi125 := 0; fi125 < 11; fi125++ {
                                v1 = v124
                                var v126 int8 = int8(v4)
                                _ = fi125
                                _ = v126
                            }
                        } else {
                            v127 := GenS12{F0: uint32(v4)}
                            var v128 int64 = -9223372036854775808
                            _ = v127
                            _ = v128
                        }
                    case "\n\t":
                        v118 += v118
                    case "a":
                        switch (("" + v123) + (("foo bar" + v123) + v123)) {
                        case "foo bar":
                            switch (fi120 * v2) {
                            case 2147483647:
                                fmt.Println(v124)
                            default:
                                v129 := GenS6{F0: (^(fi120)), F1: 519}
                                _ = v129
                            }
                        default:
                            var v130 uint64 = 4294967295
                            _ = v130
                        }
                    default:
                        v4 = v4
                    }
                    _ = v123
                    _ = v124
                }
            }
            _ = fi120
        }
        _ = v118
        _ = v119
    } else {
        switch "" {
        case "":
            var v131 uint16 = 42
            _ = v131
        case "a":
            if true {
                v132 := v3[65535]
                _ = v132
            } else {
                v4 += v4
            }
        case "foo bar":
            var v133 float32 = 1.5
            _ = v133
        }
        var v134 uint32 = 0
        var v135 int32 = -726
        _ = v134
        _ = v135
    }
    v2 -= (v2 % ((v2) | 1))
    v136 := []int{(^(v2)), (v2 % ((v2) | 1)), (v2 >> (((v1 | v1)) & 63)), (v2 / (((v2 + v2)) | 1))}
    v137 := v3[782]
    for fi138 := 0; fi138 < 1; fi138++ {
        v1 = v1
        fmt.Println(v2)
        _ = fi138
    }
    for fi139 := 0; fi139 < 19; fi139++ {
        switch "hello" {
        case "foo bar":
            v140 := GenS13{F0: 65535, F1: 32767}
            _ = v140
        case "":
            if (true && true) {
                _ = len(v136)
            } else {
                if len(v136) >= 2 {
                    v136 = v136[1:len(v136)]
                }
                v1 = v1
            }
        }
        switch -963 {
        case 1:
            v3[1] = (v137 >> (((v1 * v1)) & 63))
        case 358:
            for fi141 := 0; fi141 < 5; fi141++ {
                if true {
                    if ((v137 >> ((v1) & 63)) == v137) {
                        switch (fi141 ^ (-(v2))) {
                        case 4294967295:
                            if (v4 <= (v4 + 128)) {
                                var v142 int = 4294967295
                                v3[(v137 * v137)] = v137
                                _ = v142
                            } else {
                                fmt.Println(v2, fi139)
                            }
                        case 489:
                            fmt.Println(v1, v4)
                            fmt.Println(v1)
                            if false {
                                v3[(v137 * (v137 | v137))] = 1
                                if false {
                                    var v143 int16 = 1000
                                    if ((!true) && true) {
                                        v4 -= (v4 >> (((v1 & v1)) & 63))
                                        var v144 uint32 = v137
                                        _ = v144
                                    } else {
                                        switch fi141 {
                                        case 1:
                                            var v145 int32 = -2147483648
                                            _ = v145
                                        case 9223372036854775807:
                                            var v146 int = (fi139 + -9223372036854775808)
                                            _ = v146
                                        case -1:
                                            _ = len(v3)
                                        default:
                                            _ = len(v3)
                                        }
                                        var v147 uint16 = 219
                                        if true {
                                            var v148 bool = (-662 > -1)
                                            v2 = (^(v2))
                                            switch fi141 {
                                            case 1:
                                                v149 := GenS14{F0: 42, F1: v4, F2: v2}
                                                _ = v149
                                            case 127:
                                                _ = len(v136)
                                            }
                                            _ = v148
                                        }
                                        fmt.Println(v1)
                                        if true {
                                            var v150 int8 = 51
                                            v137 += v137
                                            switch v2 {
                                            case 4294967295:
                                                v151 := v3[v137]
                                                _ = v151
                                            case 1:
                                                v3[(v137 ^ v137)] = v137
                                            case 32768:
                                                var v152 int64 = 9223372036854775807
                                                _ = v152
                                            }
                                            _ = v150
                                        }
                                        _ = v147
                                    }
                                    v143 += v143
                                    if true {
                                        fmt.Println(v137, v137)
                                    } else {
                                        var v153 string = ""
                                        var v154 int32 = -1
                                        if false {
                                            v137 -= (v137 - 4294967295)
                                            var v155 int16 = v143
                                            if (!true) {
                                                var v156 uint8 = (v4 & 128)
                                                _ = v156
                                            } else {
                                                var v157 int = fi139
                                                fmt.Println(fi139, v157, v155)
                                                v158 := v3[(v137 % (((v137 ^ v137)) | 1))]
                                                v154 = v154
                                                switch (fi141 | int(v4)) {
                                                case 1:
                                                    var v159 int32 = (v154 + int32(v155))
                                                    _ = v159
                                                }
                                                v160 := GenS10{F0: v155, F1: 0}
                                                v161 := v3[4294967295]
                                                _ = v157
                                                _ = v158
                                                _ = v160
                                                _ = v161
                                            }
                                            var v162 float32 = -2.25
                                            fmt.Println(v162, fi139)
                                            _ = v155
                                            _ = v162
                                        }
                                        _ = v153
                                        _ = v154
                                    }
                                    _ = v143
                                }
                            }
                            switch "a" {
                            case "":
                                v163 := GenS14{F0: 726, F1: v4, F2: 0}
                                var v164 int8 = 127
                                v1 -= v1
                                var v165 bool = true
                                _ = v163
                                _ = v164
                                _ = v165
                            }
                        case -2147483648:
                            _ = len(v136)
                        }
                        v4 = v4
                        var v166 uint32 = v137
                        v2 = 2147483647
                        _ = v166
                    } else {
                        v2 = fi139
                    }
                    if (!true) {
                        var v167 int8 = genFn0((fi141 ^ (^(fi141))), 625)
                        _ = v167
                    } else {
                        switch "hello" {
                        case "":
                            switch "\n\t" {
                            case "\n\t":
                                var v168 int8 = 127
                                _ = v168
                            }
                        default:
                            fmt.Println(v1)
                        }
                        v4 = 153
                    }
                    if (false && false) {
                        if (false && false) {
                            v169 := GenS1{F0: 737}
                            switch -289 {
                            case -586:
                                v4 += v4
                            case 128:
                                var v170 bool = (!true)
                                _ = v170
                            case -117:
                                switch (fi139 + v2) {
                                case -211:
                                    v4 = 42
                                }
                            }
                            if true {
                                fmt.Println(v1, v4)
                            }
                            _ = v169
                        }
                    } else {
                        var v171 uint64 = 42
                        _ = v171
                    }
                    switch "foo bar" {
                    case "hello":
                        var v172 uint32 = v137
                        _ = v172
                    case "":
                        v2 = 32767
                        if (genFn1() <= (v4 - (v4 ^ (v4 + (v4 >> (((v1 / ((v1) | 1))) & 63)))))) {
                            v136 = append(v136, int(v4))
                            if false {
                                switch "a" {
                                case "a":
                                    v1 = 255
                                case "hello":
                                    v173 := GenS14{F0: -2147483648, F1: v4, F2: fi141}
                                    _ = v173
                                }
                                var v174 int8 = 0
                                _ = v174
                            }
                            v136 = append(v136, (v2 * v2))
                            v1 -= 65535
                            v3[v137] = (^(v137))
                        }
                    }
                } else {
                    var v175 int32 = int32(v4)
                    fmt.Println(v1, v1)
                    _ = v175
                }
                var v176 int16 = 1000
                _ = fi141
                _ = v176
            }
        case 9223372036854775807:
            fmt.Println(fi139, v1)
        }
        var v177 uint32 = v137
        _ = fi139
        _ = v177
    }
    fmt.Println(v2, v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v136:", len(v136), v136)
    fmt.Println("v137:", v137)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v136
    _ = v137
    _ = v2
    _ = v3
    _ = v4
}
