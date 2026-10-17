package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32; F1 uint; F2 int }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int16; F1 uint32; F2 int16 }
type GenS8 struct { F0 int64; F1 string }
type GenS9 struct { F0 string; F1 int32; F2 uint16 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint16; F1 bool; F2 int }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 uint8; F1 int32 }
type GenS17 struct { F0 uint }
type GenS18 struct { F0 int8 }
type GenS19 struct { F0 uint32; F1 uint64 }
type GenS20 struct { F0 int; F1 int64; F2 string }
type GenS21 struct { F0 int64 }
type GenS22 struct { F0 int8 }
func genFn0(p0 uint8, p1 uint, p2 int32) int16 {
    p0 = (p0 & p0)
    p0 = p0
    p0 -= 1
    _ = p0
    _ = p1
    _ = p2
    return 111
}
func genFn1() int16 {
    v4 := GenS1{F0: 0, F1: 925, F2: 42}
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 int8 = 42
        v6 = (v6 / (((v6 * v6)) | 1))
        var v7 int = (fi5 ^ 128)
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := []int{fi9, fi9, (fi8 >> ((297) & 63)), (^(fi9))}
            _ = len(v10)
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    _ = v4
    return 32767
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{-74, 164, 127}
    var v3 uint8 = 80
    if true {
        for fi11 := 0; fi11 < 14; fi11++ {
            switch "foo bar" {
            case "a":
                for fi12 := 0; fi12 < 10; fi12++ {
                    v13 := GenS1{F0: 116, F1: -1, F2: int64(fi12)}
                    _ = fi12
                    _ = v13
                }
            case "hello":
                _ = len(v2)
            case "":
                var v14 int16 = 0
                _ = v14
            default:
                if ((fi11 > fi11) && ((!((((((0.0 - (((1e10 - (1e10 - (1e10 + 3.14159))) + 1.5) / 1e10)) + 3.14159) * ((-2.25 - 0.1) * (float32(0.1) * 0.1))) + (float32(1.5) / (-2.25 * 1e10))) + 0.1) > (float32(float64((1e10 + ((0.0 + (3.14159 - (0.0 / 0.1))) - (1.5 * 0.0))))) / -2.25))) || (32767 <= 42))) {
                    fmt.Println(v3, fi11)
                    for fi15 := 0; fi15 < 19; fi15++ {
                        if true {
                            if (true && false) {
                                fmt.Println(fi11, fi11)
                            }
                            switch (^(fi11)) {
                            case -2147483648:
                                v3 += (v3 | 255)
                            default:
                                v2 = append(v2, -81)
                            }
                            v16 := GenS1{F0: 1, F1: 0, F2: -9223372036854775808}
                            _ = v16
                        } else {
                            var v17 string = ("" + "hello")
                            if (!((v3 + 128) > (^(v3)))) {
                                switch ("" + ((v17 + (v17 + (v17 + (((v17 + (v17 + (v17 + v17))) + "foo bar") + "")))) + v17)) {
                                case "a":
                                    if false {
                                        if (!(3.14159 <= 1e10)) {
                                            switch (fi11 << ((0) & 63)) {
                                            case 989:
                                                var v18 float32 = -2.25
                                                _ = v18
                                            case 127:
                                                var v19 float32 = -2.25
                                                _ = v19
                                            case 0:
                                                var v20 bool = (false && true)
                                                _ = v20
                                            }
                                            v3 = (v3 + v3)
                                            v21 := GenS2{F0: true}
                                            _ = v21
                                        }
                                    } else {
                                        var v22 uint32 = 0
                                        _ = len(v1)
                                        if (!true) {
                                            var v23 int64 = 687
                                            fmt.Println(v22, v17, v23)
                                            _ = v23
                                        }
                                        _ = v22
                                    }
                                default:
                                    v3 = (v3 | v3)
                                }
                            } else {
                                if (true || false) {
                                    v24 := v1[(fi15 - 2147483647)]
                                    _ = v24
                                } else {
                                    v25 := GenS1{F0: 42, F1: 42, F2: -1}
                                    _ = v25
                                }
                            }
                            v3 += 42
                            _ = v17
                        }
                        v26 := GenS3{F0: 2147483648, F1: 0}
                        v3 = v3
                        _ = fi15
                        _ = v26
                    }
                    switch ("foo bar" + "a") {
                    case "a":
                        for fi27 := 0; fi27 < 14; fi27++ {
                            switch "hello" {
                            case "\n\t":
                                v3 = (v3 ^ (v3 / (((v3 | (v3 * (v3 ^ (v3 % ((v3) | 1)))))) | 1)))
                                v28 := GenS2{F0: true}
                                _ = v28
                            case "":
                                var v29 int16 = 1000
                                _ = v29
                            case "foo bar":
                                var v30 int32 = int32(v3)
                                _ = v30
                            }
                            _ = fi27
                        }
                    }
                }
            }
            if len(v2) > 0 {
                v2[0] = 9223372036854775807
            }
            v31 := GenS3{F0: 4294967295, F1: -762}
            _ = fi11
            _ = v31
        }
        for fi32 := 0; fi32 < 14; fi32++ {
            for fi33 := 0; fi33 < 6; fi33++ {
                v3 += (^(v3))
                if false {
                    if (!(!true)) {
                        if ((((fi32 % (((fi33 | 747)) | 1)) <= fi32) && (!(!(4294967295 == 0)))) && true) {
                            v3 += (^(v3))
                        }
                        var v34 int16 = int16(v3)
                        v35 := v1[313]
                        v1[fi32] = (v35 % ((v35) | 1))
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v34
                        _ = v35
                    } else {
                        v3 = (^(v3))
                        if (1 > 1) {
                            v36 := GenS3{F0: 65535, F1: -132}
                            var v37 uint8 = 128
                            _ = v36
                            _ = v37
                        } else {
                            switch fi33 {
                            case -24:
                                v3 = (^(v3))
                                if (0.1 < (0.1 * (float64(1e10) + (((0.1 / 1e10) - 3.14159) + float64((0.0 - 3.14159)))))) {
                                    var v38 int64 = -9223372036854775808
                                    _ = v38
                                }
                            default:
                                delete(v1, fi32)
                            }
                            var v39 uint16 = 288
                            fmt.Println(v39)
                            switch "hello" {
                            case "hello":
                                v2 = append(v2, fi32)
                            }
                            if false {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v39 = (^(v39))
                            } else {
                                var v40 float64 = 1e10
                                v39 = 2
                                v39 = (v39 << ((0) & 63))
                                _ = v40
                            }
                            v41 := GenS4{F0: -174}
                            if true {
                                v42 := GenS5{F0: 1, F1: 1, F2: fi33}
                                v43 := GenS1{F0: 1, F1: 42, F2: 1}
                                v3 -= 2
                                v44 := GenS3{F0: 1, F1: -348}
                                if true {
                                    var v45 int16 = int16(v39)
                                    if len(v2) > 0 {
                                        v2[0] = (-(fi33))
                                    }
                                    _ = v45
                                } else {
                                    v39 -= (v39 | v39)
                                }
                                var v46 int32 = 0
                                switch fi33 {
                                case 128:
                                    switch "\n\t" {
                                    case "foo bar":
                                        fmt.Println(fi32, v39)
                                    case "":
                                        var v47 int8 = -128
                                        _ = v47
                                    }
                                    _ = len(v2)
                                case 0:
                                    v46 -= v46
                                default:
                                    var v48 string = "hello"
                                    _ = v48
                                }
                                fmt.Println(v46, v46)
                                _ = v42
                                _ = v43
                                _ = v44
                                _ = v46
                            } else {
                                var v49 int8 = -21
                                v50 := v1[fi33]
                                var v51 string = "\n\t"
                                _ = v49
                                _ = v50
                                _ = v51
                            }
                            _ = v39
                            _ = v41
                        }
                        var v52 int16 = -1
                        _ = v52
                    }
                } else {
                    if true {
                        var v53 int32 = 0
                        _ = v53
                    } else {
                        v54 := GenS1{F0: 100, F1: -1, F2: 60}
                        _ = v54
                    }
                    v55 := GenS6{F0: (("\n\t" + "hello") + "a")}
                    switch "foo bar" {
                    case "\n\t":
                        v3 -= v3
                    case "a":
                        v3 = (v3 / ((v3) | 1))
                    case "":
                        var v56 uint = 65535
                        _ = v56
                    }
                    switch fi32 {
                    case -948:
                        if (!false) {
                            v2 = append(v2, (fi33 + (fi33 + fi33)))
                            v2 = append(v2, fi33)
                            if true {
                                v57 := GenS7{F0: 32767, F1: 65535, F2: -32768}
                                if (!true) {
                                    v58 := GenS8{F0: 1, F1: "a"}
                                    var v59 bool = (!true)
                                    v60 := GenS1{F0: -128, F1: 349, F2: -9223372036854775808}
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                }
                                _ = v57
                            }
                        }
                    case 42:
                        if (278 <= 729) {
                            var v61 bool = true
                            if v61 {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                if (!false) {
                                    _ = len(v1)
                                }
                                v3 = (^(v3))
                            } else {
                                switch fi32 {
                                case 128:
                                    var v62 int8 = -128
                                    v63 := GenS1{F0: v62, F1: 1000, F2: 42}
                                    _ = v62
                                    _ = v63
                                case 9223372036854775807:
                                    var v64 uint16 = 32767
                                    _ = v64
                                }
                                v2 = append(v2, fi32)
                            }
                            if true {
                                v65 := v1[(^(fi32))]
                                v3 -= v3
                                _ = v65
                            } else {
                                v61 = v61
                            }
                            var v66 uint = 808
                            var v67 float32 = (0.0 / 1e10)
                            _ = v61
                            _ = v66
                            _ = v67
                        } else {
                            switch fi32 {
                            case 0:
                                v2 = append(v2, -1)
                            default:
                                fmt.Println(fi33, fi33)
                            }
                        }
                    case 4294967295:
                        if (42 != -1) {
                            fmt.Println(fi33, fi33)
                        } else {
                            if false {
                                if (1 != 653) {
                                    var v68 string = "a"
                                    v69 := GenS9{F0: (((v68 + "hello") + v68) + "a"), F1: -1, F2: 1}
                                    _ = v68
                                    _ = v69
                                } else {
                                    var v70 uint = 509
                                    _ = v70
                                }
                            }
                        }
                    default:
                        var v71 int8 = 100
                        _ = v71
                    }
                    switch "" {
                    case "hello":
                        v72 := GenS10{F0: -1}
                        _ = v72
                    case "foo bar":
                        var v73 uint = 4294967295
                        _ = v73
                    case "":
                        if true {
                            var v74 uint8 = v3
                            _ = v74
                        }
                    }
                    var v75 int = fi33
                    _ = v55
                    _ = v75
                }
                if ((!((!(fi33 != -1)) && (!true))) && true) {
                    var v76 uint8 = v3
                    v1[(fi33 | -195)] = 65535
                    if false {
                        if true {
                            v76 += 234
                            var v77 uint16 = 42
                            v2 = append(v2, (fi32 << ((835) & 63)))
                            _ = v77
                        } else {
                            v76 = 255
                            v78 := GenS8{F0: 0, F1: "a"}
                            if true {
                                v79 := GenS2{F0: (!false)}
                                _ = v79
                            }
                            _ = v78
                        }
                        _ = len(v2)
                        var v80 uint8 = (v3 * 2)
                        _ = v80
                    }
                    _ = v76
                }
                _ = fi33
            }
            if (genFn0((v3 & (v3 - 91)), 2, 128) < 278) {
                var v81 string = "hello"
                _ = v81
            } else {
                switch "a" {
                case "":
                    fmt.Println(v3, fi32)
                case "\n\t":
                    for fi82 := 0; fi82 < 4; fi82++ {
                        _ = len(v1)
                        v3 += v3
                        var v83 int64 = 0
                        _ = fi82
                        _ = v83
                    }
                case "foo bar":
                    v3 = 42
                default:
                    _ = len(v2)
                }
                fmt.Println(fi32, fi32)
                var v84 int8 = 0
                _ = v84
            }
            _ = fi32
        }
    }
    if len(v2) > 0 {
        v2[0] = 127
    }
    v3 -= v3
    switch -713 {
    case 980:
        _ = len(v1)
    case -9223372036854775808:
        v85 := GenS11{F0: 904, F1: false, F2: -1}
        _ = v85
    case 2:
        var v86 uint = 210
        _ = v86
    }
    var v87 float64 = 1e10
    var v88 uint64 = 18446744073709551615
    if (true || (true || ((!false) || false))) {
        if (0 < 32767) {
            v88 = (v88 + 2)
        } else {
            if (-1 > 9223372036854775807) {
                var v89 int = -425
                for fi90 := 0; fi90 < 1; fi90++ {
                    var v91 uint = 527
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if (1.5 >= (v87 - -2.25)) {
                        v91 = (v91 | (v91 * (v91 % (((v91 >> ((2) & 63))) | 1))))
                    } else {
                        v1[899] = 696
                    }
                    _ = fi90
                    _ = v91
                }
                if false {
                    v92 := GenS4{F0: v89}
                    v93 := GenS12{F0: 279}
                    for fi94 := 0; fi94 < 20; fi94++ {
                        v95 := GenS9{F0: (("\n\t" + ("foo bar" + "hello")) + "hello"), F1: 2147483647, F2: 65535}
                        _ = fi94
                        _ = v95
                    }
                    _ = v92
                    _ = v93
                } else {
                    switch (^(v89)) {
                    case 32768:
                        v88 += (v88 + 75)
                    case -2147483648:
                        _ = len(v1)
                    case -947:
                        for fi96 := 0; fi96 < 3; fi96++ {
                            switch "" {
                            case "foo bar":
                                v97 := GenS10{F0: -516}
                                _ = v97
                            case "hello":
                                for fi98 := 0; fi98 < 14; fi98++ {
                                    if false {
                                        var v99 string = "hello"
                                        v88 += 0
                                        _ = v99
                                    } else {
                                        if true {
                                            switch ("" + "hello") {
                                            case "":
                                                v100 := v1[-354]
                                                _ = v100
                                            case "hello":
                                                fmt.Println(fi98, fi98, fi96)
                                            case "\n\t":
                                                if true {
                                                    v3 = v3
                                                    var v101 uint = 255
                                                    _ = v101
                                                } else {
                                                    if (1 < genFn0(v3, 152, -2147483648)) {
                                                        if (false && (((v89 % ((9223372036854775807) | 1)) == fi98) && (1 > 42))) {
                                                            v102 := GenS13{F0: 0}
                                                            v3 = v3
                                                            _ = v102
                                                        }
                                                        var v103 bool = false
                                                        _ = v103
                                                    }
                                                    var v104 int64 = 0
                                                    _ = v104
                                                }
                                            }
                                            if true {
                                                v3 += (v3 % ((36) | 1))
                                            }
                                            if false {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                switch ("hello" + ("foo bar" + "foo bar")) {
                                                case "hello":
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                case "a":
                                                    fmt.Println(v89, v87, v89)
                                                case "\n\t":
                                                    switch "\n\t" {
                                                    case "":
                                                        var v105 int = v89
                                                        _ = v105
                                                    default:
                                                        switch "hello" {
                                                        case "hello":
                                                            v88 = 2
                                                        case "\n\t":
                                                            if true {
                                                                fmt.Println(v89, v88)
                                                                v87 = (v87 / (v87 - float64((0.1 + (-2.25 * 3.14159)))))
                                                            }
                                                        default:
                                                            v3 = (v3 + v3)
                                                            var v106 float64 = 1.5
                                                            v107 := GenS13{F0: 0}
                                                            _ = v106
                                                            _ = v107
                                                        }
                                                    }
                                                default:
                                                    v108 := GenS14{F0: -687}
                                                    _ = v108
                                                }
                                            } else {
                                                if false {
                                                    v89 += (fi96 << ((65535) & 63))
                                                } else {
                                                    delete(v1, v89)
                                                    var v109 uint8 = 255
                                                    switch ("" + "hello") {
                                                    case "a":
                                                        var v110 int32 = 0
                                                        _ = v110
                                                    case "foo bar":
                                                        switch fi98 {
                                                        case 42:
                                                            if true {
                                                                var v111 uint16 = 65535
                                                                if true {
                                                                    v89 -= (v89 ^ 2147483647)
                                                                    if (v87 <= (1.5 - 0.1)) {
                                                                        v87 -= (v87 * v87)
                                                                    }
                                                                } else {
                                                                    var v112 float32 = float32((1.5 * float64(float32(3.14159))))
                                                                    if (0 > -698) {
                                                                        v88 = 519
                                                                        var v113 uint32 = 634
                                                                        var v114 uint8 = (v3 - (v109 / ((v3) | 1)))
                                                                        _ = v113
                                                                        _ = v114
                                                                    }
                                                                    v87 = v87
                                                                    if true {
                                                                        fmt.Println(v87, v88, v88)
                                                                        var v115 bool = true
                                                                        v111 += (^(v111))
                                                                        _ = v115
                                                                    }
                                                                    fmt.Println(v109, v87, v88)
                                                                    _ = v112
                                                                }
                                                                _ = v111
                                                            } else {
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    var v116 uint16 = 962
                                                                    _ = v116
                                                                default:
                                                                    delete(v1, fi98)
                                                                    _ = len(v2)
                                                                    var v117 bool = true
                                                                    _ = v117
                                                                }
                                                                delete(v1, (fi98 ^ fi96))
                                                                var v118 int32 = 185
                                                                _ = v118
                                                            }
                                                        case -597:
                                                            if (322 > 580) {
                                                                v3 = v3
                                                                if len(v2) > 0 {
                                                                    v2[0] = 549
                                                                }
                                                            } else {
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    v89 += fi96
                                                                case "":
                                                                    if true {
                                                                        if false {
                                                                            v109 = (v3 << ((1) & 63))
                                                                            delete(v1, fi96)
                                                                        } else {
                                                                            v88 = v88
                                                                            v109 -= 1
                                                                        }
                                                                        var v119 float32 = 1e10
                                                                        _ = v119
                                                                    }
                                                                case "foo bar":
                                                                    switch fi96 {
                                                                    case -253:
                                                                        v3 += (v3 >> ((293) & 63))
                                                                    case -414:
                                                                        if ((((0.1 - 1e10) * ((1e10 / (3.14159 + 1.5)) / 1e10)) - -2.25) < -2.25) {
                                                                            switch "" {
                                                                            case "hello":
                                                                                switch (v89 << ((766) & 63)) {
                                                                                case 32768:
                                                                                    switch (v89 * fi96) {
                                                                                    case 1:
                                                                                        v88 = v88
                                                                                    default:
                                                                                        if ((fi96 * (fi98 ^ 1)) > 9223372036854775807) {
                                                                                            if (false && ((!false) || true)) {
                                                                                                var v120 uint16 = 780
                                                                                                v3 += (^(v109))
                                                                                                var v121 uint32 = 42
                                                                                                _ = v120
                                                                                                _ = v121
                                                                                            }
                                                                                            v89 = (fi96 - 32767)
                                                                                        }
                                                                                    }
                                                                                case -178:
                                                                                    var v122 int = (fi96 & fi96)
                                                                                    switch "\n\t" {
                                                                                    case "hello":
                                                                                        v3 = 241
                                                                                    }
                                                                                    v1[32768] = 0
                                                                                    _ = v122
                                                                                case 128:
                                                                                    v88 = (^(v88))
                                                                                default:
                                                                                    v87 = v87
                                                                                }
                                                                            default:
                                                                                v123 := GenS15{F0: (v88 << ((42) & 63))}
                                                                                _ = v123
                                                                            }
                                                                            var v124 uint64 = (v88 / (((v88 | v88)) | 1))
                                                                            v89 = fi98
                                                                            _ = v124
                                                                        } else {
                                                                            var v125 float64 = 3.14159
                                                                            var v126 int = fi98
                                                                            _ = v125
                                                                            _ = v126
                                                                        }
                                                                    case -2147483648:
                                                                        var v127 uint32 = 65535
                                                                        _ = v127
                                                                    default:
                                                                        v1[fi98] = uint32(v88)
                                                                    }
                                                                }
                                                                v128 := GenS15{F0: v88}
                                                                var v129 int32 = -1
                                                                _ = len(v2)
                                                                _ = v128
                                                                _ = v129
                                                            }
                                                        case -1:
                                                            v130 := GenS15{F0: 42}
                                                            var v131 int64 = 9223372036854775807
                                                            fmt.Println(v88)
                                                            _ = v130
                                                            _ = v131
                                                        }
                                                    case "\n\t":
                                                        if (((854 > 18446744073709551615) || (true || (1 < 1))) || (42 < 0)) {
                                                            if ((-1 <= 0) || (42 > 127)) {
                                                                _ = len(v1)
                                                                v132 := GenS16{F0: (v3 ^ v3), F1: -1}
                                                                _ = v132
                                                            } else {
                                                                if true {
                                                                    v89 -= (fi98 >> ((2) & 63))
                                                                    v88 = (v88 % ((302) | 1))
                                                                }
                                                            }
                                                            var v133 uint = 816
                                                            _ = v133
                                                        }
                                                    }
                                                    _ = v109
                                                }
                                                fmt.Println(fi96)
                                                var v134 uint8 = 42
                                                _ = v134
                                            }
                                        }
                                        if false {
                                            var v135 int8 = 63
                                            _ = v135
                                        } else {
                                            if (-1 != 1) {
                                                v3 += (v3 | (v3 | 2))
                                                v87 += -2.25
                                                fmt.Println(v88, fi98, v88)
                                            }
                                            _ = len(v2)
                                        }
                                    }
                                    _ = fi98
                                }
                            case "a":
                                for fi136 := 0; fi136 < 19; fi136++ {
                                    _ = len(v2)
                                    var v137 float32 = (float32((((v87 * (v87 * (v87 - (v87 + v87)))) + 0.0) - v87)) / (0.1 - float32((v87 + (float64(0.0) + (-2.25 - 1.5))))))
                                    if false {
                                        v138 := GenS1{F0: 127, F1: -32768, F2: int64(v88)}
                                        if len(v2) > 0 {
                                            v2[0] = (v89 >> ((0) & 63))
                                        }
                                        _ = v138
                                    }
                                    _ = fi136
                                    _ = v137
                                }
                            }
                            if (1 != 100) {
                                v139 := GenS14{F0: 290}
                                var v140 uint32 = 42
                                v141 := GenS4{F0: fi96}
                                _ = v139
                                _ = v140
                                _ = v141
                            }
                            _ = fi96
                        }
                    default:
                        fmt.Println(v87, v3, v87)
                    }
                }
                _ = v89
            } else {
                fmt.Println(v88, v87, v88)
            }
            var v142 int = -675
            _ = v142
        }
        v3 -= v3
    } else {
        for fi143 := 0; fi143 < 16; fi143++ {
            var v144 uint8 = v3
            for fi145 := 0; fi145 < 9; fi145++ {
                if true {
                    if len(v2) > 0 {
                        v2[0] = -398
                    }
                    switch (fi143 % ((-1) | 1)) {
                    case 4294967295:
                        var v146 int = (fi145 << ((382) & 63))
                        _ = v146
                    case -2147483648:
                        v147 := GenS17{F0: 1}
                        _ = v147
                    case 1:
                        fmt.Println(v3)
                    default:
                        v2 = append(v2, (^(fi145)))
                    }
                } else {
                    _ = len(v1)
                    switch "" {
                    case "\n\t":
                        v88 -= (v88 * (v88 % (((v88 * (v88 | (v88 / ((v88) | 1))))) | 1)))
                    case "":
                        if (!true) {
                            fmt.Println(fi143, fi145, v144)
                            v144 = v144
                        } else {
                            if (65535 != 1) {
                                if len(v2) > 0 {
                                    v2[0] = fi145
                                }
                            }
                            _ = len(v2)
                        }
                    }
                    v144 -= v3
                }
                v3 -= (v144 % ((v144) | 1))
                v148 := GenS10{F0: 2147483647}
                v88 -= (v88 | 628)
                var v149 bool = (false || true)
                _ = fi145
                _ = v148
                _ = v149
            }
            _ = fi143
            _ = v144
        }
        v150 := GenS3{F0: 2147483648, F1: -504}
        _ = v150
    }
    v151 := GenS18{F0: int8(v88)}
    if (-2147483648 > 1) {
        var v152 bool = false
        _ = v152
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v153 string = ""
    switch 0 {
    case 0:
        switch 0 {
        case 866:
            var v154 string = ""
            _ = v154
        }
    case -534:
        if false {
            if ((255 >= uint(v3)) || (!true)) {
                for fi155 := 0; fi155 < 12; fi155++ {
                    for fi156 := 0; fi156 < 2; fi156++ {
                        if len(v2) > 0 {
                            v2[0] = fi156
                        }
                        _ = fi156
                    }
                    if true {
                        for fi157 := 0; fi157 < 7; fi157++ {
                            switch "foo bar" {
                            case "\n\t":
                                switch v153 {
                                case "a":
                                    var v158 uint = 961
                                    _ = v158
                                case "hello":
                                    fmt.Println(v3, fi157, v87)
                                }
                            case "":
                                v2 = append(v2, fi155)
                            case "a":
                                if (!(!true)) {
                                    v88 = (v88 << ((496) & 63))
                                    var v159 int64 = 205
                                    _ = v159
                                }
                            }
                            _ = fi157
                        }
                        if true {
                            for fi160 := 0; fi160 < 8; fi160++ {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v161 float64 = v87
                                _ = len(v1)
                                _ = fi160
                                _ = v161
                            }
                        } else {
                            v87 = -2.25
                            v162 := GenS19{F0: 863, F1: (v88 | (v88 / ((v88) | 1)))}
                            _ = v162
                        }
                    } else {
                        var v163 string = "hello"
                        for fi164 := 0; fi164 < 5; fi164++ {
                            v1[-9223372036854775808] = 2
                            _ = fi164
                        }
                        _ = v163
                    }
                    _ = fi155
                }
                for fi165 := 0; fi165 < 1; fi165++ {
                    v88 = 465
                    _ = fi165
                }
                switch v153 {
                case "hello":
                    v166 := GenS15{F0: (v88 ^ v88)}
                    _ = v166
                case "a":
                    var v167 int32 = -2147483648
                    _ = v167
                case "\n\t":
                    v2 = append(v2, -1)
                default:
                    if (!false) {
                        if (v88 > (v88 ^ (^(v88)))) {
                            for fi168 := 0; fi168 < 4; fi168++ {
                                for fi169 := 0; fi169 < 5; fi169++ {
                                    v3 -= 42
                                    _ = fi169
                                }
                                v170 := GenS20{F0: (fi168 / ((fi168) | 1)), F1: -1, F2: ("" + ((((v153 + v153) + "foo bar") + "foo bar") + v153))}
                                var v171 uint = 42
                                _ = fi168
                                _ = v170
                                _ = v171
                            }
                            for fi172 := 0; fi172 < 18; fi172++ {
                                switch (fi172 | fi172) {
                                case 42:
                                    v88 -= v88
                                case -2147483648:
                                    var v173 float32 = -2.25
                                    _ = v173
                                default:
                                    fmt.Println(fi172, v87, v87)
                                }
                                v174 := GenS7{F0: 0, F1: 249, F2: 42}
                                _ = fi172
                                _ = v174
                            }
                            if true {
                                var v175 uint64 = (v88 << ((418) & 63))
                                v88 -= (v175 / ((0) | 1))
                                _ = v175
                            }
                        }
                    }
                }
            } else {
                switch 9223372036854775807 {
                case -2147483648:
                    if false {
                        v176 := GenS16{F0: 246, F1: 394}
                        _ = v176
                    }
                case -9223372036854775808:
                    v177 := GenS4{F0: -1}
                    _ = v177
                default:
                    v87 = (0.1 / 1.5)
                }
            }
            v178 := GenS15{F0: (v88 ^ 1)}
            if true {
                v153 = "foo bar"
            } else {
                switch 32767 {
                case -163:
                    fmt.Println(v153, v3)
                default:
                    v3 -= (v3 * v3)
                }
            }
            _ = v178
        } else {
            if (708 == 0) {
                switch -217 {
                case 127:
                    if true {
                        for fi179 := 0; fi179 < 11; fi179++ {
                            if true {
                                var v180 int64 = -962
                                v153 = ""
                                _ = v180
                            } else {
                                v181 := GenS3{F0: 0, F1: 2147483647}
                                if (!true) {
                                    for fi182 := 0; fi182 < 1; fi182++ {
                                        fmt.Println(fi179, v153, fi182)
                                        v183 := GenS4{F0: fi179}
                                        v184 := GenS19{F0: 583, F1: v88}
                                        _ = len(v1)
                                        _ = fi182
                                        _ = v183
                                        _ = v184
                                    }
                                    v3 -= (v3 | v3)
                                }
                                _ = v181
                            }
                            var v185 int16 = 313
                            _ = fi179
                            _ = v185
                        }
                        v87 = v87
                        v87 = ((((float64((float32(v87) * 1e10)) * ((((v87 / 1e10) + v87) - 3.14159) / 3.14159)) - 3.14159) / (3.14159 / v87)) / v87)
                    } else {
                        _ = len(v2)
                    }
                case 9223372036854775807:
                    v3 = v3
                case 42:
                    v3 = 2
                default:
                    for fi186 := 0; fi186 < 15; fi186++ {
                        var v187 uint16 = 2
                        _ = fi186
                        _ = v187
                    }
                }
                v2 = append(v2, 127)
            }
            fmt.Println(v87, v153, v88)
        }
    case -2147483648:
        v153 = "\n\t"
    default:
        for fi188 := 0; fi188 < 15; fi188++ {
            v189 := GenS11{F0: 2, F1: (0 >= -520), F2: (fi188 | -187)}
            _ = fi188
            _ = v189
        }
    }
    delete(v1, 32768)
    var v190 uint16 = 357
    if ((((v153 + v153) + v153) != (v153 + ((v153 + "foo bar") + v153))) || (!true)) {
        for fi191 := 0; fi191 < 18; fi191++ {
            for fi192 := 0; fi192 < 7; fi192++ {
                v193 := GenS21{F0: 1}
                var v194 uint32 = 2
                _ = fi192
                _ = v193
                _ = v194
            }
            var v195 int = fi191
            _ = fi191
            _ = v195
        }
        v153 = "foo bar"
    }
    switch 1 {
    case 128:
        for fi196 := 0; fi196 < 9; fi196++ {
            fmt.Println(v190, v190, v88)
            fmt.Println(v3, v87, v88)
            _ = fi196
        }
    case -115:
        if ((false || false) || false) {
            var v197 int8 = 42
            _ = v197
        } else {
            v198 := GenS1{F0: -128, F1: -135, F2: 9223372036854775807}
            switch 960 {
            case 568:
                v87 = v87
            case -642:
                var v199 bool = true
                _ = v199
            default:
                switch (((v153 + v153) + (((v153 + v153) + (v153 + ("" + (("" + v153) + ("" + "a"))))) + "\n\t")) + "") {
                case "hello":
                    _ = len(v2)
                }
            }
            var v200 int64 = 295
            _ = v198
            _ = v200
        }
    case -9223372036854775808:
        v201 := GenS22{F0: 1}
        _ = v201
    default:
        switch -2147483648 {
        case 1:
            fmt.Println(v87, v87)
        case 42:
            for fi202 := 0; fi202 < 7; fi202++ {
                v203 := v1[fi202]
                for fi204 := 0; fi204 < 9; fi204++ {
                    v190 -= v190
                    _ = fi204
                }
                switch fi202 {
                case -9223372036854775808:
                    v88 += 42
                }
                _ = fi202
                _ = v203
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v153:", v153)
    fmt.Println("v190:", v190)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v151
    _ = v153
    _ = v190
    _ = v2
    _ = v3
    _ = v87
    _ = v88
}
