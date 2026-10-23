package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16; F1 uint8; F2 int32 }
type GenS5 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS6 struct { F0 int8; F1 uint64 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 string; F1 int64; F2 uint8 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint8; F1 uint8; F2 uint }
type GenS12 struct { F0 int8; F1 bool; F2 uint32 }
type GenS13 struct { F0 string; F1 int8; F2 int32 }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 uint8; F1 int8 }
func genFn0() int16 {
    var v7 bool = (false || false)
    fmt.Println(v7)
    v7 = (270 <= 2)
    switch "\n\t" {
    case "a":
        var v8 int8 = 100
        _ = v8
    case "":
        switch "hello" {
        case "hello":
            v9 := []int{9223372036854775807}
            _ = v9
        default:
            v10 := []int{128}
            _ = v10
        }
    case "hello":
        v7 = (true && (!v7))
    }
    v11 := []int{}
    _ = v11
    _ = v7
    return -1
}
func genFn1() string {
    v12 := []int{657}
    switch "hello" {
    case "hello":
        for fi13 := 0; fi13 < 4; fi13++ {
            var v14 uint16 = 0
            for fi15 := 0; fi15 < 19; fi15++ {
                _ = len(v12)
                v14 -= 1
                v12 = append(v12, fi13)
                _ = fi15
            }
            var v16 int8 = -1
            _ = fi13
            _ = v14
            _ = v16
        }
    case "a":
        var v17 uint = 376
        _ = v17
    default:
        if ((1.5 * 0.0) == (-2.25 * 1.5)) {
            var v18 int8 = -1
            for fi19 := 0; fi19 < 10; fi19++ {
                for fi20 := 0; fi20 < 5; fi20++ {
                    var v21 uint = 18446744073709551615
                    var v22 uint16 = 785
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                for fi23 := 0; fi23 < 10; fi23++ {
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                    var v24 int64 = 0
                    if len(v12) > 0 {
                        v12[0] = 0
                    }
                    if (498 != -32768) {
                        fmt.Println(fi23, fi19)
                        switch fi23 {
                        case 32768:
                            v18 -= v18
                        case 0:
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                            var v25 bool = (0 == 1)
                            var v26 int32 = 481
                            _ = v25
                            _ = v26
                        case 42:
                            if true {
                                var v27 uint8 = 251
                                _ = v27
                            } else {
                                switch (^(fi19)) {
                                case 128:
                                    fmt.Println(v24, v18)
                                case 0:
                                    v18 = v18
                                }
                            }
                        }
                    } else {
                        var v28 int32 = 2147483647
                        _ = v28
                    }
                    _ = fi23
                    _ = v24
                }
                _ = fi19
            }
            _ = v18
        } else {
            if false {
                var v29 int = 127
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                _ = v29
            }
            var v30 int8 = 0
            _ = v30
        }
    }
    v31 := GenS1{F0: 127, F1: 65535}
    _ = v12
    _ = v31
    return "a"
}

func main() {
    var v1 uint8 = 42
    var v2 int64 = -655
    v3 := []int{-189, 42}
    var v4 int32 = 0
    v5 := []uint32{2147483648}
    var v6 uint = 65535
    _ = len(v5)
    fmt.Println(v1)
    var v32 string = ("hello" + (("hello" + "hello") + ("hello" + ("hello" + ("\n\t" + ("a" + "a"))))))
    for fi33 := 0; fi33 < 13; fi33++ {
        var v34 float32 = (0.0 * 1.5)
        var v35 string = v32
        v6 -= v6
        _ = fi33
        _ = v34
        _ = v35
    }
    if false {
        var v36 uint8 = v1
        switch 42 {
        case 32767:
            v37 := GenS1{F0: int8(v6), F1: 844}
            _ = v37
        case -2147483648:
            if ((v4 | (v4 ^ (v4 + v4))) == 1) {
                v6 += v6
                fmt.Println(v2)
            } else {
                v1 = v36
                switch 127 {
                case 42:
                    v38 := GenS1{F0: -1, F1: 2}
                    _ = v38
                case 32768:
                    switch (v32 + v32) {
                    case "a":
                        v39 := GenS1{F0: 42, F1: 1}
                        _ = v39
                    case "foo bar":
                        v40 := GenS2{F0: 305, F1: -9223372036854775808}
                        _ = v40
                    case "":
                        v1 = 248
                    default:
                        v3 = append(v3, 9223372036854775807)
                    }
                case -445:
                    if false {
                        for fi41 := 0; fi41 < 6; fi41++ {
                            v42 := GenS1{F0: 127, F1: 255}
                            switch fi41 {
                            case 32768:
                                v43 := GenS3{F0: 108}
                                _ = v43
                            default:
                                switch v32 {
                                case "a":
                                    if true {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                case "\n\t":
                                    for fi44 := 0; fi44 < 10; fi44++ {
                                        if true {
                                            var v45 int = 665
                                            fmt.Println(v2)
                                            _ = v45
                                        } else {
                                            var v46 int64 = v2
                                            v46 += v46
                                            switch (v32 + v32) {
                                            case "hello":
                                                if false {
                                                    v46 = int64(v1)
                                                    switch "" {
                                                    case "foo bar":
                                                        if true {
                                                            var v47 int32 = v4
                                                            _ = v47
                                                        } else {
                                                            if false {
                                                                if (-128 >= -128) {
                                                                    if (705 != v4) {
                                                                        if len(v3) > 0 {
                                                                            v3[0] = (^(fi41))
                                                                        }
                                                                        v4 = (-(v4))
                                                                        var v48 uint16 = 65535
                                                                        _ = v48
                                                                    } else {
                                                                        switch int(v6) {
                                                                        case 57:
                                                                            if true {
                                                                                v49 := GenS4{F0: 1, F1: v36, F2: v4}
                                                                                if (32767 < 1) {
                                                                                    var v50 int16 = genFn0()
                                                                                    if (127 != v1) {
                                                                                        var v51 float32 = (1e10 * (1.5 + float32(0.0)))
                                                                                        switch (v32 + v32) {
                                                                                        case "foo bar":
                                                                                            v52 := GenS4{F0: int16(fi44), F1: 2, F2: v4}
                                                                                            _ = v52
                                                                                        case "\n\t":
                                                                                            if len(v5) > 0 {
                                                                                                v5[0] = 2
                                                                                            }
                                                                                        }
                                                                                        if (!false) {
                                                                                            fmt.Println(v36)
                                                                                            switch fi41 {
                                                                                            case 42:
                                                                                                if true {
                                                                                                    var v53 int8 = 0
                                                                                                    v2 = (v46 | v2)
                                                                                                    _ = v53
                                                                                                }
                                                                                            case -2147483648:
                                                                                                if true {
                                                                                                    if len(v5) > 0 {
                                                                                                        v5[0] = 4294967295
                                                                                                    }
                                                                                                    v32 = ("\n\t" + v32)
                                                                                                    v54 := GenS5{F0: 600, F1: 42, F2: 2}
                                                                                                    _ = v54
                                                                                                }
                                                                                            default:
                                                                                                fmt.Println(v2)
                                                                                            }
                                                                                        }
                                                                                        _ = v51
                                                                                    }
                                                                                    _ = v50
                                                                                }
                                                                                _ = v49
                                                                            } else {
                                                                                v32 = ""
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                                if (42 <= v36) {
                                                                                    var v55 int = 32767
                                                                                    var v56 uint16 = 32767
                                                                                    _ = v55
                                                                                    _ = v56
                                                                                }
                                                                            }
                                                                        case 42:
                                                                            switch "a" {
                                                                            case "a":
                                                                                switch (fi41 ^ fi41) {
                                                                                case 2147483647:
                                                                                    switch (fi44 & fi41) {
                                                                                    case 42:
                                                                                        var v57 float64 = float64(float32(1e10))
                                                                                        _ = v57
                                                                                    case 0:
                                                                                        fmt.Println(v2, v46, v36)
                                                                                    }
                                                                                case -2147483648:
                                                                                    var v58 int8 = 47
                                                                                    _ = v58
                                                                                }
                                                                            case "\n\t":
                                                                                v59 := GenS5{F0: 201, F1: 396, F2: 582}
                                                                                _ = v59
                                                                            case "foo bar":
                                                                                v4 = v4
                                                                            }
                                                                            if false {
                                                                                fmt.Println(v32)
                                                                            }
                                                                        default:
                                                                            v60 := GenS5{F0: 0, F1: 122, F2: 619}
                                                                            _ = v60
                                                                        }
                                                                        v32 = v32
                                                                    }
                                                                    var v61 float64 = (0.0 / 3.14159)
                                                                    v6 += v6
                                                                    _ = v61
                                                                }
                                                                if (false && false) {
                                                                    var v62 uint16 = 33
                                                                    _ = v62
                                                                }
                                                                v63 := GenS6{F0: 100, F1: 42}
                                                                var v64 uint64 = 4294967295
                                                                v65 := GenS7{F0: -128}
                                                                _ = v63
                                                                _ = v64
                                                                _ = v65
                                                            }
                                                            var v66 uint64 = 18446744073709551615
                                                            _ = v66
                                                        }
                                                    }
                                                    _ = len(v3)
                                                } else {
                                                    v2 += 42
                                                    if (!true) {
                                                        v4 = v4
                                                        fmt.Println(v32, v36, v6)
                                                        v36 = v1
                                                        v46 = (^(v46))
                                                        fmt.Println(fi41, v6, fi44)
                                                    } else {
                                                        v2 += (v2 * 1)
                                                    }
                                                    var v67 int64 = v46
                                                    v68 := GenS1{F0: 37, F1: 32767}
                                                    _ = len(v5)
                                                    if (true || true) {
                                                        if true {
                                                            var v69 float64 = 3.14159
                                                            v2 = 9223372036854775807
                                                            _ = v69
                                                        }
                                                        v5 = append(v5, 4294967295)
                                                        if false {
                                                            v70 := GenS8{F0: (v4 & (v4 ^ v4))}
                                                            var v71 uint32 = 287
                                                            _ = v70
                                                            _ = v71
                                                        }
                                                        v67 -= v46
                                                        v32 = (((v32 + "foo bar") + ((v32 + v32) + ((v32 + v32) + ""))) + v32)
                                                    } else {
                                                        fmt.Println(v2)
                                                        fmt.Println(v6)
                                                    }
                                                    v72 := GenS3{F0: 44}
                                                    _ = v67
                                                    _ = v68
                                                    _ = v72
                                                }
                                            case "\n\t":
                                                v73 := GenS5{F0: 917, F1: 572, F2: 461}
                                                _ = v73
                                            case "foo bar":
                                                var v74 int64 = v46
                                                _ = v74
                                            default:
                                                var v75 uint16 = 520
                                                _ = v75
                                            }
                                            _ = v46
                                        }
                                        v1 -= v36
                                        var v76 string = v32
                                        _ = fi44
                                        _ = v76
                                    }
                                case "":
                                    var v77 uint32 = 266
                                    _ = v77
                                default:
                                    if false {
                                        if false {
                                            v3 = append(v3, fi41)
                                            v78 := GenS9{F0: "\n\t", F1: 42, F2: 42}
                                            if (3.14159 != ((0.0 - 0.0) * 0.1)) {
                                                if true {
                                                    fmt.Println(v2, fi41, v6)
                                                } else {
                                                    if false {
                                                        v79 := GenS9{F0: (v32 + v32), F1: 330, F2: 0}
                                                        if len(v5) > 0 {
                                                            v5[0] = 1
                                                        }
                                                        _ = v79
                                                    }
                                                }
                                                v80 := GenS5{F0: 382, F1: 4294967295, F2: 853}
                                                _ = v80
                                            }
                                            _ = v78
                                        }
                                        var v81 int16 = 1000
                                        v4 = 42
                                        _ = v81
                                    }
                                }
                            }
                            _ = fi41
                            _ = v42
                        }
                        _ = len(v5)
                        var v82 int = -467
                        _ = v82
                    }
                }
                fmt.Println(v36, v36)
            }
        case 9223372036854775807:
            var v83 int64 = (v2 >> ((v6) & 63))
            _ = v83
        }
        _ = v36
    }
    if (v6 != (v6 << (((^(v6))) & 63))) {
        var v84 uint64 = 0
        v32 = v32
        _ = v84
    } else {
        for fi85 := 0; fi85 < 7; fi85++ {
            if ((v1 / ((v1) | 1)) < (v1 ^ v1)) {
                var v86 uint32 = 2147483648
                if (((0.0 + (3.14159 + (float64(3.14159) * (1e10 - 3.14159)))) <= 0.1) || false) {
                    var v87 int16 = 42
                    v6 = (v6 % (((v6 % (((v6 * v6)) | 1))) | 1))
                    _ = v87
                } else {
                    var v88 uint = 0
                    switch int(v4) {
                    case -1:
                        var v89 uint32 = (^(v86))
                        _ = v89
                    }
                    var v90 float64 = -2.25
                    _ = v88
                    _ = v90
                }
                v91 := GenS1{F0: 63, F1: 498}
                _ = v86
                _ = v91
            } else {
                fmt.Println(v4, v32)
                v92 := GenS7{F0: -6}
                if false {
                    switch (fi85 + 1) {
                    case 127:
                        if (v1 != v1) {
                            var v93 int8 = 0
                            for fi94 := 0; fi94 < 18; fi94++ {
                                var v95 int = (-(fi85))
                                _ = fi94
                                _ = v95
                            }
                            _ = v93
                        } else {
                            switch (^(fi85)) {
                            case 128:
                                var v96 int8 = 100
                                _ = v96
                            case 127:
                                v97 := GenS10{F0: ((v6 & 1) <= v6)}
                                _ = v97
                            case 9223372036854775807:
                                v4 += 2147483647
                            default:
                                var v98 float32 = (float32(float64(((3.14159 + -2.25) + 3.14159))) + ((((-2.25 - (1.5 + float32(1.5))) / ((((((1.5 - (1.5 + 3.14159)) * (float32(0.0) - (1.5 / 0.1))) / -2.25) * 1.5) * (3.14159 - float32(0.0))) + float32((3.14159 + (1.5 - (1e10 - 0.1)))))) + float32(1.5)) * -2.25))
                                _ = v98
                            }
                            for fi99 := 0; fi99 < 17; fi99++ {
                                var v100 int = (fi85 * 319)
                                var v101 int = (^(fi85))
                                if false {
                                    switch ("" + "foo bar") {
                                    case "hello":
                                        if ((v4 & 0) == v4) {
                                            fmt.Println(fi85)
                                        }
                                    case "\n\t":
                                        fmt.Println(v100)
                                    case "a":
                                        v1 = 97
                                    }
                                }
                                _ = fi99
                                _ = v100
                                _ = v101
                            }
                            if (0.0 <= 1e10) {
                                v5 = append(v5, 65535)
                            }
                        }
                    case 128:
                        switch v32 {
                        case "":
                            v32 = ((v32 + "hello") + (v32 + v32))
                        case "hello":
                            fmt.Println(v1)
                        default:
                            fmt.Println(v4, v4)
                        }
                    case 4294967295:
                        for fi102 := 0; fi102 < 20; fi102++ {
                            var v103 uint32 = 4294967295
                            _ = fi102
                            _ = v103
                        }
                    default:
                        v2 = (v2 - v2)
                    }
                    for fi104 := 0; fi104 < 3; fi104++ {
                        var v105 int8 = 127
                        v3 = append(v3, int(v4))
                        var v106 int16 = 860
                        _ = fi104
                        _ = v105
                        _ = v106
                    }
                } else {
                    v107 := GenS6{F0: 42, F1: 332}
                    _ = v107
                }
                _ = v92
            }
            _ = fi85
        }
    }
    if false {
        switch v32 {
        case "hello":
            v108 := GenS11{F0: 127, F1: v1, F2: v6}
            _ = v108
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi109 := 0; fi109 < 6; fi109++ {
            _ = len(v5)
            _ = fi109
        }
    } else {
        var v110 string = "a"
        if (!(true && (true && (0.0 != 0.0)))) {
            for fi111 := 0; fi111 < 20; fi111++ {
                if false {
                    var v112 int8 = 100
                    var v113 bool = (int64(fi111) > 9223372036854775807)
                    v4 = (v4 % ((int32(v112)) | 1))
                    _ = v112
                    _ = v113
                } else {
                    v6 = v6
                    if false {
                        v32 = v110
                        if true {
                            if true {
                                if ((0 <= 4294967295) || (v6 >= (v6 ^ (v6 * (^(v6)))))) {
                                    for fi114 := 0; fi114 < 20; fi114++ {
                                        switch fi114 {
                                        case -1:
                                            v2 = (^(v2))
                                        case -842:
                                            switch fi111 {
                                            case 42:
                                                v115 := GenS6{F0: 0, F1: 18446744073709551615}
                                                _ = v115
                                            default:
                                                v110 = (((((("" + "") + (v32 + v32)) + v110) + v32) + v32) + "")
                                            }
                                        }
                                        var v116 uint32 = 2147483647
                                        var v117 string = v32
                                        _ = fi114
                                        _ = v116
                                        _ = v117
                                    }
                                }
                                var v118 float32 = 1.5
                                _ = v118
                            }
                            if (127 >= -33) {
                                for fi119 := 0; fi119 < 10; fi119++ {
                                    fmt.Println(v4, fi111)
                                    fmt.Println(fi119, v2, v110)
                                    var v120 uint16 = 32767
                                    v121 := GenS12{F0: -1, F1: false, F2: 4294967295}
                                    _ = fi119
                                    _ = v120
                                    _ = v121
                                }
                                if (127 != -43) {
                                    for fi122 := 0; fi122 < 12; fi122++ {
                                        if len(v3) > 0 {
                                            v3[0] = fi122
                                        }
                                        var v123 int8 = 0
                                        if (249 != 65535) {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            var v124 uint32 = 2
                                            if false {
                                                _ = len(v5)
                                                if (4294967295 > v124) {
                                                    var v125 float32 = (float32(((((-2.25 * float64(-2.25)) + -2.25) / 1.0) / (1e10 * 1e10))) - float32(0.1))
                                                    switch "foo bar" {
                                                    case "":
                                                        v2 += 0
                                                    }
                                                    _ = v125
                                                } else {
                                                    var v126 float32 = (1e10 - 1e10)
                                                    var v127 int64 = v2
                                                    _ = v126
                                                    _ = v127
                                                }
                                                v3 = append(v3, fi111)
                                            } else {
                                                fmt.Println(v1)
                                                v3 = append(v3, fi122)
                                                var v128 uint = 2
                                                _ = v128
                                            }
                                            _ = v124
                                        }
                                        _ = fi122
                                        _ = v123
                                    }
                                    var v129 int = (fi111 ^ (fi111 + (fi111 | (fi111 << ((v6) & 63)))))
                                    _ = v129
                                } else {
                                    var v130 int64 = (v2 / ((v2) | 1))
                                    var v131 uint32 = 65535
                                    if false {
                                        v132 := GenS6{F0: 1, F1: 686}
                                        v2 = (v130 / (((v2 ^ v2)) | 1))
                                        _ = v132
                                    }
                                    _ = v130
                                    _ = v131
                                }
                            }
                        }
                    }
                }
                v6 -= (v6 >> ((v6) & 63))
                var v133 uint64 = 0
                _ = fi111
                _ = v133
            }
            if (!false) {
                var v134 int = -1
                if (true && (255 <= 0)) {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                _ = v134
            } else {
                for fi135 := 0; fi135 < 15; fi135++ {
                    v136 := GenS13{F0: "a", F1: 0, F2: (v4 << ((65535) & 63))}
                    _ = fi135
                    _ = v136
                }
                v32 = "a"
            }
            v32 = v110
        } else {
            v1 = v1
            var v137 uint64 = 4294967295
            _ = v137
        }
        var v138 uint64 = 0
        _ = v110
        _ = v138
    }
    _ = len(v3)
    for fi139 := 0; fi139 < 7; fi139++ {
        for fi140 := 0; fi140 < 19; fi140++ {
            v32 = (v32 + "hello")
            v141 := GenS14{F0: 231}
            fmt.Println(v1, v32)
            _ = fi140
            _ = v141
        }
        _ = fi139
    }
    for fi142 := 0; fi142 < 2; fi142++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi142
    }
    var v143 int16 = 42
    for fi144 := 0; fi144 < 6; fi144++ {
        if false {
            var v145 int16 = (v143 | (v143 | v143))
            switch fi144 {
            case 2147483647:
                var v146 float64 = float64(3.14159)
                _ = v146
            case 4294967295:
                var v147 uint = 948
                _ = v147
            }
            v148 := GenS15{F0: (v1 | 242), F1: 127}
            _ = v145
            _ = v148
        } else {
            for fi149 := 0; fi149 < 17; fi149++ {
                var v150 int8 = 115
                v3 = append(v3, 1)
                _ = fi149
                _ = v150
            }
            _ = len(v3)
            v1 = v1
        }
        v151 := GenS1{F0: -114, F1: 255}
        _ = fi144
        _ = v151
    }
    fmt.Println("v1:", v1)
    fmt.Println("v143:", v143)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v143
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
