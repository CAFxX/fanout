package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 int64; F2 int16 }
type GenS4 struct { F0 uint16; F1 bool; F2 bool }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int32; F1 uint16 }
type GenS9 struct { F0 int; F1 int8 }
type GenS10 struct { F0 uint32; F1 string; F2 int }
type GenS11 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 bool; F1 int16 }
type GenS14 struct { F0 string; F1 uint64 }
type GenS15 struct { F0 int16; F1 string; F2 int64 }

func main() {
    var v1 int16 = -1
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]string)
    var v4 bool = (("" + "") == "a")
    v1 += (-(v1))
    v4 = false
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 int = (fi5 | fi5)
        if v4 {
            var v7 int64 = -183
            v8 := GenS1{F0: -1, F1: 86, F2: (v7 | v7)}
            _ = v7
            _ = v8
        }
        var v9 int64 = -9223372036854775808
        _ = fi5
        _ = v6
        _ = v9
    }
    if v4 {
        var v10 uint64 = 329
        switch "foo bar" {
        case "\n\t":
            v1 = v1
        case "hello":
            delete(v3, 871)
        }
        _ = v10
    } else {
        v1 += 0
        v11 := GenS2{F0: 368}
        _ = v11
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 uint32 = 0
        if true {
            switch "a" {
            case "foo bar":
                if v4 {
                    v1 = (v1 ^ 523)
                } else {
                    v2[fi12] = (^(v13))
                    fmt.Println(v4, v4)
                }
            case "hello":
                v14 := GenS2{F0: -1}
                _ = v14
            case "":
                fmt.Println(v1, v4)
            default:
                v15 := GenS1{F0: 0, F1: 100, F2: 22}
                _ = v15
            }
            switch fi12 {
            case 977:
                if false {
                    v16 := []int{-2147483648, fi12, (fi12 & 42), (-(fi12))}
                    _ = v16
                } else {
                    var v17 string = "hello"
                    v13 -= (v13 % (((v13 / (((v13 % ((0) | 1))) | 1))) | 1))
                    switch -9223372036854775808 {
                    case 0:
                        v18 := GenS3{F0: ((v17 + v17) + (((v17 + v17) + v17) + "a")), F1: 1, F2: (v1 % (((v1 / ((v1) | 1))) | 1))}
                        _ = v18
                    default:
                        var v19 string = ("a" + "a")
                        _ = v19
                    }
                    _ = v17
                }
            case 1:
                var v20 uint = 544
                _ = v20
            case -476:
                var v21 int = fi12
                _ = v21
            default:
                var v22 uint16 = 65535
                _ = v22
            }
        }
        for fi23 := 0; fi23 < 5; fi23++ {
            var v24 bool = (-2147483648 >= 714)
            if (v24 || true) {
                v1 += (v1 >> ((42) & 63))
                switch "" {
                case "a":
                    var v25 int = fi12
                    _ = v25
                default:
                    v26 := v3[(v13 - (v13 >> ((1) & 63)))]
                    _ = v26
                }
                v27 := GenS1{F0: 42, F1: 127, F2: -9223372036854775808}
                v1 = 0
                switch fi23 {
                case -1:
                    v28 := []int{fi12}
                    _ = v28
                case 127:
                    v29 := []int{}
                    _ = v29
                default:
                    v13 -= (v13 << ((288) & 63))
                }
                _ = v27
            }
            _ = fi23
            _ = v24
        }
        _ = fi12
        _ = v13
    }
    fmt.Println(v1, v4)
    switch 0 {
    case 4294967295:
        v30 := v2[127]
        _ = v30
    case -535:
        if (v4 || (98 > 0)) {
            fmt.Println(v4, v4)
        }
    case 2147483647:
        var v31 int8 = 0
        _ = v31
    }
    delete(v2, 2147483647)
    var v32 bool = ((!(!(9223372036854775807 == -2147483648))) || true)
    v1 = 366
    v33 := GenS4{F0: 42, F1: (18446744073709551615 <= 0), F2: (("\n\t" + ("a" + "\n\t")) < ("hello" + "a"))}
    if (!v32) {
        if v4 {
            var v34 uint32 = 2
            for fi35 := 0; fi35 < 9; fi35++ {
                for fi36 := 0; fi36 < 9; fi36++ {
                    switch fi36 {
                    case -472:
                        if true {
                            v37 := GenS5{F0: (fi36 << ((403) & 63)), F1: fi35}
                            v4 = ((^(v34)) == v34)
                            _ = v37
                        }
                    case 32767:
                        v1 = v1
                    case -9223372036854775808:
                        delete(v2, -120)
                    default:
                        v38 := []int{32767, (fi36 >> ((0) & 63)), 32767, 32768}
                        _ = v38
                    }
                    if ((1e10 - ((-2.25 * 3.14159) + (0.1 - (1.5 / 1.0)))) >= 1.5) {
                        fmt.Println(fi35, v34, v34)
                        var v39 string = ""
                        if v32 {
                            var v40 uint64 = 1
                            v40 -= (^(v40))
                            v40 -= (v40 | (v40 - 44))
                            _ = v40
                        } else {
                            fmt.Println(v32)
                            switch "" {
                            case "hello":
                                if (2 >= 42) {
                                    var v41 int = (fi36 % (((fi35 - 32768)) | 1))
                                    _ = v41
                                }
                                v42 := []int{0, fi36, (fi36 % ((424) | 1))}
                                _ = v42
                            default:
                                var v43 uint8 = 127
                                _ = v43
                            }
                            if true {
                                if false {
                                    v34 += (v34 % (((^(v34))) | 1))
                                } else {
                                    var v44 int64 = 9223372036854775807
                                    fmt.Println(v1, fi35, v44)
                                    _ = v44
                                }
                                if v4 {
                                    v34 = 1
                                    v2[fi36] = (^(v34))
                                    switch "" {
                                    case "\n\t":
                                        var v45 uint32 = v34
                                        _ = v45
                                    }
                                } else {
                                    if (v34 < 977) {
                                        var v46 int8 = 1
                                        _ = v46
                                    }
                                    var v47 uint32 = v34
                                    _ = v47
                                }
                                var v48 float32 = (3.14159 * (1.5 / 1.0))
                                _ = v48
                            } else {
                                var v49 float32 = 0.1
                                fmt.Println(v39, v1)
                                _ = v49
                            }
                            var v50 bool = true
                            if false {
                                if true {
                                    fmt.Println(v50, v34)
                                    v39 = v39
                                    if v50 {
                                        switch "a" {
                                        case "hello":
                                            v2[9223372036854775807] = (v34 - 65535)
                                        default:
                                            v50 = (2147483647 < -1)
                                        }
                                    } else {
                                        fmt.Println(fi35, v1, v34)
                                        v3[(^(v34))] = (v39 + v39)
                                        switch (v39 + v39) {
                                        case "\n\t":
                                            fmt.Println(fi35, fi36)
                                        case "foo bar":
                                            v51 := []int{}
                                            _ = v51
                                        case "hello":
                                            v52 := GenS5{F0: 32768, F1: -2147483648}
                                            _ = v52
                                        default:
                                            v53 := GenS3{F0: (((v39 + v39) + "foo bar") + v39), F1: 42, F2: 32767}
                                            _ = v53
                                        }
                                        var v54 int16 = 42
                                        v1 = 922
                                        v55 := []int{}
                                        _ = v54
                                        _ = v55
                                    }
                                } else {
                                    var v56 uint64 = 1
                                    if ((255 == 2) || v4) {
                                        v57 := GenS5{F0: (fi36 & (fi36 / (((fi36 >> ((1) & 63))) | 1))), F1: fi35}
                                        var v58 bool = (-2.25 != 0.0)
                                        _ = len(v2)
                                        _ = v57
                                        _ = v58
                                    }
                                    _ = v56
                                }
                                v39 = v39
                                v50 = v50
                                if (!true) {
                                    v59 := []int{fi35, fi35}
                                    var v60 bool = false
                                    _ = v59
                                    _ = v60
                                } else {
                                    var v61 float32 = 0.0
                                    _ = v61
                                }
                                if v4 {
                                    fmt.Println(v1, fi36)
                                    var v62 string = v39
                                    _ = v62
                                }
                                v4 = false
                            }
                            _ = v50
                        }
                        _ = v39
                    } else {
                        var v63 string = "\n\t"
                        v1 += v1
                        var v64 float64 = 0.1
                        _ = v63
                        _ = v64
                    }
                    _ = fi36
                }
                if false {
                    v1 = int16(fi35)
                    v32 = (4294967295 < 1)
                    if true {
                        if v4 {
                            switch ("\n\t" + "") {
                            case "foo bar":
                                if (0 != 2) {
                                    v65 := GenS2{F0: -1}
                                    switch (("" + "") + "foo bar") {
                                    case "a":
                                        for fi66 := 0; fi66 < 3; fi66++ {
                                            if v32 {
                                                v4 = true
                                                v67 := []int{9223372036854775807, fi66, (fi35 >> ((65535) & 63))}
                                                v32 = true
                                                _ = v67
                                            }
                                            v34 -= 551
                                            _ = fi66
                                        }
                                    default:
                                        if true {
                                            v32 = v32
                                            v34 = 273
                                        } else {
                                            var v68 uint64 = 0
                                            _ = v68
                                        }
                                    }
                                    _ = v65
                                }
                            case "\n\t":
                                v69 := GenS6{F0: (v34 % ((v34) | 1))}
                                _ = v69
                            case "a":
                                if (523 <= 0) {
                                    var v70 int8 = 117
                                    v71 := []int{(-(fi35))}
                                    fmt.Println(v4, v4, v34)
                                    _ = v70
                                    _ = v71
                                }
                            }
                            v72 := GenS2{F0: 42}
                            for fi73 := 0; fi73 < 16; fi73++ {
                                if v32 {
                                    if (255 > 2) {
                                        v32 = v32
                                        switch (fi35 + fi35) {
                                        case 127:
                                            var v74 int16 = v1
                                            _ = v74
                                        case 32767:
                                            delete(v2, 919)
                                        case 42:
                                            var v75 uint = 317
                                            _ = v75
                                        }
                                    }
                                    fmt.Println(v1)
                                } else {
                                    var v76 float64 = 3.14159
                                    switch "" {
                                    case "":
                                        if false {
                                            fmt.Println(v32, fi73)
                                        }
                                    case "\n\t":
                                        delete(v3, (^(v34)))
                                    }
                                    _ = v76
                                }
                                v32 = (v32 || (true && v32))
                                _ = fi73
                            }
                            _ = v72
                        } else {
                            v77 := []int{}
                            switch fi35 {
                            case 32768:
                                if false {
                                    if v4 {
                                        switch ("hello" + "a") {
                                        case "hello":
                                            if v4 {
                                                switch fi35 {
                                                case 32768:
                                                    if (v4 && v4) {
                                                        fmt.Println(v4)
                                                    } else {
                                                        switch "a" {
                                                        case "":
                                                            v34 = 0
                                                        case "a":
                                                            switch "hello" {
                                                            case "foo bar":
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    var v78 int32 = -1
                                                                    _ = v78
                                                                case "hello":
                                                                    var v79 string = "hello"
                                                                    _ = v79
                                                                }
                                                            case "\n\t":
                                                                v4 = (845 >= 8)
                                                            default:
                                                                _ = len(v2)
                                                            }
                                                        case "hello":
                                                            switch fi35 {
                                                            case 938:
                                                                _ = len(v77)
                                                            case 9223372036854775807:
                                                                if v4 {
                                                                    fmt.Println(v1, v1, v4)
                                                                    if v32 {
                                                                        delete(v3, v34)
                                                                    } else {
                                                                        for fi80 := 0; fi80 < 17; fi80++ {
                                                                            switch 32767 {
                                                                            case -2147483648:
                                                                                var v81 uint = 1
                                                                                _ = v81
                                                                            case 809:
                                                                                v82 := GenS7{F0: (v1 << ((18446744073709551615) & 63))}
                                                                                _ = v82
                                                                            }
                                                                            _ = fi80
                                                                        }
                                                                        v4 = false
                                                                    }
                                                                } else {
                                                                    if (-2147483648 < -72) {
                                                                        var v83 float32 = -2.25
                                                                        _ = v83
                                                                    }
                                                                    if (!true) {
                                                                        if v32 {
                                                                            switch fi35 {
                                                                            case 2147483647:
                                                                                v3[v34] = "foo bar"
                                                                            }
                                                                            v4 = v32
                                                                        }
                                                                    } else {
                                                                        if len(v77) >= 2 {
                                                                            v77 = v77[1:len(v77)]
                                                                        }
                                                                    }
                                                                }
                                                            case -2147483648:
                                                                if true {
                                                                    var v84 float32 = -2.25
                                                                    if v4 {
                                                                        v1 = 1000
                                                                        for fi85 := 0; fi85 < 6; fi85++ {
                                                                            v77 = append(v77, fi85)
                                                                            _ = fi85
                                                                        }
                                                                    } else {
                                                                        v1 = (v1 << ((65535) & 63))
                                                                    }
                                                                    _ = v84
                                                                } else {
                                                                    v2[fi35] = 0
                                                                    for fi86 := 0; fi86 < 15; fi86++ {
                                                                        v34 -= v34
                                                                        _ = fi86
                                                                    }
                                                                    v1 -= v1
                                                                }
                                                            default:
                                                                var v87 int = (fi35 ^ (-(fi35)))
                                                                _ = v87
                                                            }
                                                        default:
                                                            if len(v77) > 0 {
                                                                v77[0] = fi35
                                                            }
                                                        }
                                                        if ((-(v1)) < (v1 ^ (v1 ^ -1))) {
                                                            fmt.Println(fi35, v1, v4)
                                                            if len(v77) >= 2 {
                                                                v77 = v77[1:len(v77)]
                                                            }
                                                        }
                                                        fmt.Println(v34, v32)
                                                    }
                                                case -1:
                                                    v34 = v34
                                                }
                                                fmt.Println(v34)
                                                _ = len(v77)
                                            }
                                        default:
                                            v32 = (((((0.1 - (((0.1 - float32(1.5)) - (1.5 - 1e10)) - 1e10)) / (0.1 - 3.14159)) - (3.14159 + 3.14159)) + -2.25) >= 0.1)
                                        }
                                    }
                                    if (fi35 > (-(fi35))) {
                                        var v88 uint8 = 2
                                        switch "foo bar" {
                                        case "\n\t":
                                            switch fi35 {
                                            case 42:
                                                var v89 int32 = -1
                                                _ = v89
                                            case 325:
                                                v1 -= (v1 & (v1 & (^(v1))))
                                            case -659:
                                                v34 += (v34 >> ((42) & 63))
                                            default:
                                                for fi90 := 0; fi90 < 20; fi90++ {
                                                    if v4 {
                                                        v88 = v88
                                                    } else {
                                                        v91 := GenS3{F0: "hello", F1: -1, F2: (v1 << ((uint(v34)) & 63))}
                                                        v88 = (^(v88))
                                                        var v92 int8 = 1
                                                        _ = v91
                                                        _ = v92
                                                    }
                                                    if (2 < 0) {
                                                        _ = len(v77)
                                                        if (!v4) {
                                                            v88 = 7
                                                        }
                                                        v93 := v2[(fi90 >> ((47) & 63))]
                                                        var v94 float64 = 0.0
                                                        if v4 {
                                                            var v95 int64 = 1
                                                            fmt.Println(v93)
                                                            _ = v95
                                                        }
                                                        var v96 bool = (!(110 != 55))
                                                        _ = v93
                                                        _ = v94
                                                        _ = v96
                                                    }
                                                    switch (("a" + "hello") + (("hello" + "hello") + "hello")) {
                                                    case "a":
                                                        v97 := GenS8{F0: 42, F1: 286}
                                                        _ = v97
                                                    case "\n\t":
                                                        fmt.Println(fi90, fi35)
                                                    }
                                                    _ = fi90
                                                }
                                            }
                                        case "foo bar":
                                            for fi98 := 0; fi98 < 9; fi98++ {
                                                var v99 uint8 = 121
                                                _ = fi98
                                                _ = v99
                                            }
                                        case "a":
                                            var v100 int8 = 0
                                            _ = v100
                                        }
                                        var v101 float64 = 1.5
                                        _ = v101
                                        _ = v88
                                    }
                                    v34 += 2147483648
                                } else {
                                    switch 854 {
                                    case -9223372036854775808:
                                        if v4 {
                                            switch ("" + ("a" + ("" + "a"))) {
                                            case "":
                                                var v102 int16 = (v1 << ((4294967295) & 63))
                                                _ = v102
                                            case "a":
                                                v1 -= 0
                                            }
                                            var v103 float32 = ((0.0 / -2.25) / 3.14159)
                                            for fi104 := 0; fi104 < 8; fi104++ {
                                                v32 = (-2.25 == 0.1)
                                                if len(v77) >= 2 {
                                                    v77 = v77[1:len(v77)]
                                                }
                                                if v32 {
                                                    var v105 uint64 = 911
                                                    if (0 != -27) {
                                                        fmt.Println(v34, v32, v34)
                                                    } else {
                                                        var v106 int16 = (^(v1))
                                                        v107 := GenS6{F0: 871}
                                                        if (1 <= 65535) {
                                                            v108 := GenS7{F0: v1}
                                                            fmt.Println(v106, fi104)
                                                            _ = v108
                                                        } else {
                                                            if (v106 != v1) {
                                                                v106 += (v106 * (v106 % ((v1) | 1)))
                                                                fmt.Println(v103, fi104, v106)
                                                                var v109 int64 = -9223372036854775808
                                                                _ = v109
                                                            }
                                                            v4 = v32
                                                            v110 := GenS5{F0: (fi35 ^ fi35), F1: 42}
                                                            if (v32 || (2 < 101)) {
                                                                if (!v32) {
                                                                    v1 = -142
                                                                } else {
                                                                    var v111 float32 = 3.14159
                                                                    switch ("\n\t" + "") {
                                                                    case "":
                                                                        v77 = append(v77, (fi35 % ((32767) | 1)))
                                                                    case "hello":
                                                                        v77 = append(v77, -359)
                                                                    case "\n\t":
                                                                        var v112 float64 = 1.5
                                                                        _ = v112
                                                                    default:
                                                                        var v113 uint32 = (v34 & v34)
                                                                        _ = v113
                                                                    }
                                                                    v114 := GenS4{F0: 65535, F1: true, F2: v4}
                                                                    _ = v111
                                                                    _ = v114
                                                                }
                                                                if (-2147483648 < 0) {
                                                                    var v115 uint16 = 273
                                                                    _ = v115
                                                                }
                                                                if v4 {
                                                                    if true {
                                                                        var v116 uint32 = (v34 & v34)
                                                                        _ = len(v3)
                                                                        _ = v116
                                                                    }
                                                                    var v117 string = "foo bar"
                                                                    var v118 int8 = -72
                                                                    _ = v117
                                                                    _ = v118
                                                                }
                                                            } else {
                                                                v119 := GenS9{F0: (fi35 & fi35), F1: 100}
                                                                delete(v2, fi35)
                                                                if v4 {
                                                                    if v4 {
                                                                        var v120 uint = 180
                                                                        v77 = append(v77, 42)
                                                                        _ = v120
                                                                    }
                                                                }
                                                                _ = v119
                                                            }
                                                            v121 := GenS9{F0: fi35, F1: -128}
                                                            _ = v110
                                                            _ = v121
                                                        }
                                                        _ = v106
                                                        _ = v107
                                                    }
                                                    v122 := GenS6{F0: 2}
                                                    _ = v105
                                                    _ = v122
                                                } else {
                                                    if (v32 || (v32 && true)) {
                                                        v123 := v3[2147483648]
                                                        _ = len(v3)
                                                        var v124 uint8 = 255
                                                        _ = v123
                                                        _ = v124
                                                    } else {
                                                        if v32 {
                                                            var v125 int16 = int16(v34)
                                                            _ = len(v77)
                                                            _ = v125
                                                        } else {
                                                            v32 = false
                                                        }
                                                        if v32 {
                                                            delete(v2, (fi35 - (-(fi35))))
                                                            if false {
                                                                v32 = v32
                                                                v1 -= 1000
                                                            } else {
                                                                var v126 int64 = 0
                                                                var v127 float32 = v103
                                                                _ = v126
                                                                _ = v127
                                                            }
                                                        }
                                                        v4 = v4
                                                    }
                                                    delete(v3, 42)
                                                    var v128 int64 = 42
                                                    switch "a" {
                                                    case "hello":
                                                        v103 += (v103 + 1.5)
                                                    }
                                                    _ = v128
                                                }
                                                _ = fi104
                                            }
                                            _ = v103
                                        }
                                    case 284:
                                        if v32 {
                                            v32 = v32
                                            var v129 float64 = float64((-2.25 / (3.14159 + 3.14159)))
                                            _ = v129
                                        }
                                    }
                                    v130 := GenS10{F0: 0, F1: "\n\t", F2: fi35}
                                    if (v34 > (v34 / ((0) | 1))) {
                                        fmt.Println(v32, v32)
                                    }
                                    _ = v130
                                }
                            default:
                                var v131 string = "hello"
                                _ = v131
                            }
                            v1 -= 32767
                            _ = v77
                        }
                    }
                } else {
                    if true {
                        var v132 int64 = 535
                        for fi133 := 0; fi133 < 5; fi133++ {
                            v134 := GenS2{F0: v132}
                            switch fi35 {
                            case 4294967295:
                                if v4 {
                                    switch "foo bar" {
                                    case "":
                                        var v135 uint16 = 255
                                        _ = v135
                                    case "hello":
                                        v136 := []int{-9223372036854775808, -1}
                                        _ = v136
                                    default:
                                        fmt.Println(v132, v32, v32)
                                    }
                                    var v137 uint64 = 2
                                    _ = v137
                                }
                                v138 := v3[v34]
                                _ = v138
                            case -583:
                                v139 := GenS7{F0: -32768}
                                v140 := GenS11{F0: 128, F1: 42, F2: (v34 - (v34 << ((18446744073709551615) & 63)))}
                                v34 = v34
                                var v141 int16 = v1
                                if (47 <= 255) {
                                    v142 := GenS12{F0: 1}
                                    if (v32 || false) {
                                        v34 = 442
                                    }
                                    _ = v142
                                }
                                _ = v139
                                _ = v140
                                _ = v141
                            }
                            var v143 uint = 255
                            _ = fi133
                            _ = v134
                            _ = v143
                        }
                        _ = v132
                    } else {
                        if v32 {
                            if v32 {
                                v32 = (!false)
                                switch "" {
                                case "a":
                                    v4 = false
                                case "":
                                    var v144 bool = v4
                                    _ = v144
                                default:
                                    var v145 int64 = -186
                                    _ = v145
                                }
                            } else {
                                switch ((("hello" + "") + "a") + "") {
                                case "hello":
                                    var v146 int8 = -128
                                    _ = v146
                                case "":
                                    v4 = v4
                                case "foo bar":
                                    v147 := []int{(fi35 + 9223372036854775807), (fi35 | (fi35 << ((18446744073709551615) & 63))), (fi35 % ((fi35) | 1)), 2147483647}
                                    _ = v147
                                }
                                fmt.Println(v32, fi35, v4)
                            }
                            v1 = (v1 - v1)
                        }
                        switch fi35 {
                        case -1:
                            v32 = false
                        case 4294967295:
                            if (480 != 2) {
                                v148 := []int{(fi35 * fi35), (fi35 / ((fi35) | 1)), fi35}
                                switch (fi35 ^ (^(fi35))) {
                                case 32767:
                                    fmt.Println(v1, fi35, v32)
                                }
                                for fi149 := 0; fi149 < 5; fi149++ {
                                    var v150 int = fi35
                                    switch fi35 {
                                    case -9223372036854775808:
                                        var v151 int32 = 42
                                        _ = v151
                                    case 312:
                                        fmt.Println(v150, v1)
                                    case 128:
                                        switch "foo bar" {
                                        case "hello":
                                            var v152 float32 = (1e10 - 1e10)
                                            _ = v152
                                        }
                                        v34 = (v34 % (((v34 >> ((uint(v34)) & 63))) | 1))
                                        v4 = v4
                                    }
                                    _ = fi149
                                    _ = v150
                                }
                                _ = v148
                            } else {
                                for fi153 := 0; fi153 < 14; fi153++ {
                                    v154 := GenS13{F0: ((^(v1)) >= v1), F1: v1}
                                    _ = fi153
                                    _ = v154
                                }
                                switch fi35 {
                                case 0:
                                    v155 := GenS3{F0: "", F1: 0, F2: 485}
                                    _ = v155
                                case 483:
                                    if true {
                                        v32 = ((!false) || ((!((!v32) && true)) && (-128 < 127)))
                                        v156 := v2[fi35]
                                        v157 := []int{128, (fi35 % (((fi35 ^ (fi35 * fi35))) | 1)), 304}
                                        _ = v156
                                        _ = v157
                                    }
                                default:
                                    var v158 int16 = -1
                                    _ = v158
                                }
                                if v4 {
                                    if (0 != 610) {
                                        for fi159 := 0; fi159 < 14; fi159++ {
                                            var v160 bool = false
                                            v32 = false
                                            _ = fi159
                                            _ = v160
                                        }
                                        v34 = v34
                                        var v161 uint32 = (^(v34))
                                        _ = v161
                                    } else {
                                        fmt.Println(v1)
                                        var v162 uint64 = 2
                                        _ = v162
                                    }
                                    if (fi35 == 471) {
                                        fmt.Println(v34)
                                        v34 -= v34
                                        switch "\n\t" {
                                        case "foo bar":
                                            v4 = v32
                                        default:
                                            v163 := GenS9{F0: 0, F1: 100}
                                            _ = v163
                                        }
                                    }
                                }
                            }
                        }
                        v4 = v4
                    }
                }
                _ = fi35
            }
            _ = v34
        } else {
            var v164 uint8 = 255
            _ = v164
        }
        v4 = v32
        v165 := GenS14{F0: "hello", F1: 276}
        _ = v165
    } else {
        v1 += -127
        v166 := GenS15{F0: v1, F1: ("hello" + "a"), F2: -1}
        _ = v166
    }
    v167 := []int{}
    fmt.Println(v1)
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v167:", len(v167), v167)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v167
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
