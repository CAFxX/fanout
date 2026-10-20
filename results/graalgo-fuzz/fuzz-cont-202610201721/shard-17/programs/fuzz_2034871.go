package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint; F1 int64; F2 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8; F1 int16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64; F1 uint; F2 bool }
type GenS8 struct { F0 int8; F1 bool }
type GenS9 struct { F0 string; F1 uint16; F2 int8 }
type GenS10 struct { F0 int; F1 bool }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS14 struct { F0 int }
type GenS15 struct { F0 uint64; F1 uint }
type GenS16 struct { F0 uint64; F1 string; F2 int8 }
type GenS17 struct { F0 uint32; F1 bool; F2 int }
type GenS18 struct { F0 int8 }

func main() {
    var v1 uint16 = 65535
    var v2 int8 = -1
    v3 := make(map[int]uint32)
    var v4 uint8 = 0
    var v5 int32 = 42
    v6 := make(map[int]int)
    if ((v1 & 255) >= 0) {
        for fi7 := 0; fi7 < 18; fi7++ {
            if (v5 >= int32(v2)) {
                var v8 uint32 = 2
                switch "\n\t" {
                case "foo bar":
                    var v9 string = ""
                    _ = v9
                case "hello":
                    v10 := GenS1{F0: 0, F1: (v8 << ((0) & 63)), F2: -9223372036854775808}
                    _ = v10
                case "":
                    var v11 uint32 = v8
                    _ = v11
                default:
                    var v12 uint16 = (v1 * (v1 % (((v1 % ((v1) | 1))) | 1)))
                    _ = v12
                }
                switch int(v2) {
                case -344:
                    for fi13 := 0; fi13 < 11; fi13++ {
                        v8 -= (v8 | 0)
                        v4 = (^(v4))
                        if (false || ((197 == -1) || (99 <= 1))) {
                            v14 := GenS1{F0: 9223372036854775807, F1: v8, F2: 42}
                            v4 = 1
                            switch -9223372036854775808 {
                            case 32768:
                                fmt.Println(fi13, v4)
                            case 128:
                                var v15 bool = (1.5 < 1e10)
                                _ = v15
                            default:
                                var v16 uint16 = (v1 & v1)
                                _ = v16
                            }
                            v17 := []int{}
                            v5 += (v5 << ((38) & 63))
                            _ = v14
                            _ = v17
                        } else {
                            switch (("foo bar" + "foo bar") + "") {
                            case "a":
                                if (0 != 4294967295) {
                                    v18 := GenS2{F0: 1, F1: false, F2: (v1 << ((0) & 63))}
                                    var v19 bool = false
                                    _ = v18
                                    _ = v19
                                }
                            case "foo bar":
                                v20 := GenS1{F0: 0, F1: 1, F2: 0}
                                _ = v20
                            case "hello":
                                if (v2 > (v2 & (v2 - -128))) {
                                    if ((789 != 446) && false) {
                                        v21 := []int{fi7, fi7, (fi13 / ((fi13) | 1))}
                                        if false {
                                            switch -699 {
                                            case 127:
                                                if len(v21) > 0 {
                                                    v21[0] = fi13
                                                }
                                                v22 := GenS1{F0: 0, F1: 1, F2: 0}
                                                if (v5 > v5) {
                                                    v5 += (-(v5))
                                                    v21 = append(v21, (-(fi7)))
                                                }
                                                _ = v22
                                            default:
                                                switch fi13 {
                                                case -534:
                                                    v23 := GenS1{F0: 0, F1: 42, F2: -1}
                                                    fmt.Println(v2, v1, v8)
                                                    var v24 int8 = v2
                                                    _ = v23
                                                    _ = v24
                                                }
                                            }
                                            v4 = v4
                                        }
                                        var v25 int = (fi13 / (((fi13 << ((65535) & 63))) | 1))
                                        _ = v21
                                        _ = v25
                                    }
                                } else {
                                    v26 := v6[-584]
                                    _ = v26
                                }
                            default:
                                v27 := GenS3{F0: 18446744073709551615, F1: 0, F2: (fi13 >= (fi13 << ((65535) & 63)))}
                                _ = v27
                            }
                        }
                        _ = fi13
                    }
                case 32767:
                    var v28 uint16 = 255
                    _ = v28
                case 4294967295:
                    var v29 int32 = (v5 >> ((65535) & 63))
                    _ = v29
                }
                _ = v8
            } else {
                var v30 uint16 = (v1 >> ((629) & 63))
                var v31 uint32 = 1
                _ = v30
                _ = v31
            }
            var v32 string = "a"
            _ = fi7
            _ = v32
        }
        if (32768 < 42) {
            if (int8(v4) < (^(v2))) {
                if (127 >= -1) {
                    var v33 uint = 0
                    if (!false) {
                        for fi34 := 0; fi34 < 4; fi34++ {
                            v35 := GenS3{F0: v33, F1: 0, F2: false}
                            _ = fi34
                            _ = v35
                        }
                    } else {
                        v36 := []int{9223372036854775807, 959, 943}
                        _ = v36
                    }
                    if (false && (4294967295 != 267)) {
                        v3[4294967295] = 1
                        fmt.Println(v5, v1, v4)
                    }
                    _ = v33
                }
                for fi37 := 0; fi37 < 4; fi37++ {
                    if false {
                        v5 = v5
                        delete(v3, 2147483647)
                    }
                    _ = fi37
                }
                var v38 float64 = float64(float32(3.14159))
                _ = v38
            } else {
                v39 := GenS1{F0: -9223372036854775808, F1: 137, F2: -557}
                var v40 int = -750
                for fi41 := 0; fi41 < 8; fi41++ {
                    var v42 uint64 = 18446744073709551615
                    v43 := GenS3{F0: 93, F1: 633, F2: (!(true && ((v2 % ((100) | 1)) != int8(v42))))}
                    if false {
                        v44 := GenS4{F0: -673}
                        v45 := []int{}
                        v46 := GenS5{F0: v2, F1: 1}
                        _ = v44
                        _ = v45
                        _ = v46
                    }
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                _ = v39
                _ = v40
            }
            if false {
                var v47 int64 = 441
                if (32767 > (^(v1))) {
                    v48 := GenS6{F0: v1}
                    _ = v48
                } else {
                    v49 := GenS3{F0: uint(v1), F1: 0, F2: ((v47 >> ((4294967295) & 63)) <= (v47 + int64(v2)))}
                    _ = v49
                }
                _ = v47
            }
            switch ("" + ("foo bar" + "\n\t")) {
            case "\n\t":
                v50 := GenS3{F0: 779, F1: 1, F2: (false && false)}
                _ = v50
            case "a":
                fmt.Println(v4)
            }
        }
    } else {
        if false {
            if (!((("" + ((("hello" + "hello") + "hello") + ((("a" + "a") + "a") + ((("a" + ("\n\t" + "a")) + "hello") + "")))) >= "foo bar") && (true && false))) {
                var v51 uint8 = (v4 % ((v4) | 1))
                _ = v51
            } else {
                switch 127 {
                case 0:
                    var v52 int8 = v2
                    _ = v52
                case 32768:
                    var v53 int8 = (-(v2))
                    _ = v53
                default:
                    if (!((1.5 < 3.14159) && false)) {
                        v54 := GenS7{F0: 976, F1: 275, F2: false}
                        switch ("" + "\n\t") {
                        case "\n\t":
                            if (-427 < 42) {
                                switch 42 {
                                case -1:
                                    if (!(-2.25 != 3.14159)) {
                                        v1 -= 795
                                        switch 127 {
                                        case 32767:
                                            v2 = v2
                                        case 0:
                                            fmt.Println(v5, v2)
                                        case -462:
                                            switch 42 {
                                            case 9223372036854775807:
                                                switch -2147483648 {
                                                case 32767:
                                                    if false {
                                                        v4 = (v4 % ((v4) | 1))
                                                        var v55 int32 = 1
                                                        v2 -= 42
                                                        _ = v55
                                                    } else {
                                                        v56 := GenS8{F0: -128, F1: false}
                                                        v57 := GenS2{F0: 2147483647, F1: false, F2: 98}
                                                        switch "" {
                                                        case "":
                                                            if true {
                                                                var v58 float32 = 0.0
                                                                v59 := GenS2{F0: 0, F1: false, F2: (v1 * v1)}
                                                                _ = v58
                                                                _ = v59
                                                            }
                                                        case "foo bar":
                                                            for fi60 := 0; fi60 < 11; fi60++ {
                                                                var v61 string = "a"
                                                                v5 = (v5 | v5)
                                                                var v62 uint = 18446744073709551615
                                                                _ = fi60
                                                                _ = v61
                                                                _ = v62
                                                            }
                                                        case "a":
                                                            var v63 int32 = v5
                                                            _ = v63
                                                        default:
                                                            if (int8(v4) <= v2) {
                                                                v2 = v2
                                                            }
                                                        }
                                                        _ = v56
                                                        _ = v57
                                                    }
                                                case -9223372036854775808:
                                                    var v64 int32 = (-(v5))
                                                    _ = v64
                                                case -330:
                                                    if (29 > 42) {
                                                        var v65 string = (("" + "\n\t") + "")
                                                        switch 4294967295 {
                                                        case 32767:
                                                            if (false || false) {
                                                                var v66 uint = 889
                                                                v65 = v65
                                                                _ = v66
                                                            } else {
                                                                fmt.Println(v5, v1)
                                                                fmt.Println(v5, v5, v1)
                                                                v67 := []int{9223372036854775807}
                                                                _ = v67
                                                            }
                                                        case -155:
                                                            var v68 float64 = -2.25
                                                            _ = v68
                                                        case -2147483648:
                                                            v5 = (v5 ^ (v5 ^ (v5 << ((2) & 63))))
                                                        default:
                                                            v1 = 1
                                                        }
                                                        switch 2147483647 {
                                                        case 2147483647:
                                                            v4 += v4
                                                        case 503:
                                                            v4 = v4
                                                        case -2147483648:
                                                            v69 := GenS1{F0: 1, F1: 65535, F2: 1}
                                                            _ = v69
                                                        default:
                                                            for fi70 := 0; fi70 < 4; fi70++ {
                                                                v71 := GenS3{F0: 65535, F1: 9223372036854775807, F2: (!true)}
                                                                delete(v3, -9223372036854775808)
                                                                if false {
                                                                    if (true || true) {
                                                                        fmt.Println(v1, v5)
                                                                        if true {
                                                                            if true {
                                                                                for fi72 := 0; fi72 < 8; fi72++ {
                                                                                    v73 := GenS9{F0: ("foo bar" + v65), F1: v1, F2: v2}
                                                                                    _ = fi72
                                                                                    _ = v73
                                                                                }
                                                                                for fi74 := 0; fi74 < 12; fi74++ {
                                                                                    v75 := GenS10{F0: -164, F1: ((true || true) && true)}
                                                                                    v76 := GenS9{F0: v65, F1: 2, F2: (^(v2))}
                                                                                    switch (((v65 + v65) + ((v65 + "foo bar") + (v65 + "a"))) + ("hello" + v65)) {
                                                                                    case "foo bar":
                                                                                        _ = len(v6)
                                                                                    case "hello":
                                                                                        fmt.Println(v5, v4)
                                                                                        v2 = (v2 ^ v2)
                                                                                        fmt.Println(fi70, v65)
                                                                                    case "\n\t":
                                                                                        v4 = 127
                                                                                    }
                                                                                    v4 = (v4 >> ((379) & 63))
                                                                                    v1 = 609
                                                                                    var v77 uint64 = 962
                                                                                    _ = fi74
                                                                                    _ = v75
                                                                                    _ = v76
                                                                                    _ = v77
                                                                                }
                                                                                var v78 uint = 1
                                                                                _ = v78
                                                                            }
                                                                            switch fi70 {
                                                                            case -9223372036854775808:
                                                                                switch fi70 {
                                                                                case 32767:
                                                                                    var v79 uint32 = 833
                                                                                    _ = v79
                                                                                default:
                                                                                    v65 = ("\n\t" + v65)
                                                                                }
                                                                            default:
                                                                                if true {
                                                                                    v80 := GenS11{F0: 42}
                                                                                    v81 := []int{(fi70 << ((42) & 63)), fi70, (fi70 & (fi70 & 547))}
                                                                                    _ = v80
                                                                                    _ = v81
                                                                                } else {
                                                                                    v4 -= v4
                                                                                }
                                                                            }
                                                                            fmt.Println(v65)
                                                                        }
                                                                    } else {
                                                                        var v82 bool = ((false && true) && true)
                                                                        v5 += (^(v5))
                                                                        _ = v82
                                                                    }
                                                                } else {
                                                                    switch v65 {
                                                                    case "\n\t":
                                                                        if (3.14159 != (1e10 + 1.5)) {
                                                                            if (!(!(207 > 432))) {
                                                                                switch v65 {
                                                                                case "hello":
                                                                                    v83 := GenS10{F0: (fi70 * (-(fi70))), F1: (!(355 == 42))}
                                                                                    _ = v83
                                                                                case "":
                                                                                    var v84 uint = uint(v1)
                                                                                    _ = v84
                                                                                case "\n\t":
                                                                                    for fi85 := 0; fi85 < 15; fi85++ {
                                                                                        v86 := GenS12{F0: fi70}
                                                                                        _ = fi85
                                                                                        _ = v86
                                                                                    }
                                                                                default:
                                                                                    var v87 uint64 = uint64(v1)
                                                                                    _ = v87
                                                                                }
                                                                                v88 := GenS9{F0: (("a" + "hello") + "hello"), F1: 1, F2: -43}
                                                                                for fi89 := 0; fi89 < 16; fi89++ {
                                                                                    var v90 int8 = -128
                                                                                    v65 = "\n\t"
                                                                                    fmt.Println(fi89, v4, v1)
                                                                                    _ = fi89
                                                                                    _ = v90
                                                                                }
                                                                                _ = v88
                                                                            }
                                                                            v3[(fi70 | (-(fi70)))] = 2147483647
                                                                        }
                                                                    default:
                                                                        fmt.Println(v5, v65, v4)
                                                                    }
                                                                    if (!(824 > 65535)) {
                                                                        if (!(!false)) {
                                                                            delete(v6, (fi70 - -9223372036854775808))
                                                                            v91 := []int{fi70}
                                                                            switch (fi70 << ((0) & 63)) {
                                                                            case 4294967295:
                                                                                fmt.Println(v1)
                                                                            case 2147483647:
                                                                                v4 = v4
                                                                            }
                                                                            _ = v91
                                                                        }
                                                                        var v92 uint8 = (v4 >> ((2) & 63))
                                                                        _ = v92
                                                                    } else {
                                                                        if false {
                                                                            if (fi70 == (^(fi70))) {
                                                                                v4 = uint8(v1)
                                                                                var v93 bool = (0 > 2147483647)
                                                                                _ = v93
                                                                            }
                                                                            var v94 bool = ((fi70 ^ fi70) <= (^(fi70)))
                                                                            _ = v94
                                                                        } else {
                                                                            for fi95 := 0; fi95 < 2; fi95++ {
                                                                                fmt.Println(v2, v65)
                                                                                _ = len(v3)
                                                                                _ = fi95
                                                                            }
                                                                            switch ((v65 + (((v65 + (((v65 + (("\n\t" + "a") + (("" + ("a" + v65)) + ((v65 + v65) + v65)))) + (("a" + v65) + v65)) + v65)) + "foo bar") + (((v65 + (v65 + "")) + "\n\t") + v65))) + v65) {
                                                                            case "\n\t":
                                                                                v96 := []int{}
                                                                                _ = v96
                                                                            case "foo bar":
                                                                                if (-1 == -1) {
                                                                                    switch v65 {
                                                                                    case "foo bar":
                                                                                        var v97 uint32 = 30
                                                                                        _ = v97
                                                                                    case "hello":
                                                                                        v98 := v6[fi70]
                                                                                        _ = v98
                                                                                    case "":
                                                                                        for fi99 := 0; fi99 < 19; fi99++ {
                                                                                            var v100 uint32 = 183
                                                                                            v101 := v3[fi70]
                                                                                            var v102 string = (v65 + "hello")
                                                                                            _ = fi99
                                                                                            _ = v100
                                                                                            _ = v101
                                                                                            _ = v102
                                                                                        }
                                                                                    default:
                                                                                        switch "\n\t" {
                                                                                        case "hello":
                                                                                            for fi103 := 0; fi103 < 17; fi103++ {
                                                                                                if false {
                                                                                                    var v104 int64 = int64(fi70)
                                                                                                    v4 += 222
                                                                                                    var v105 float32 = 1e10
                                                                                                    _ = v104
                                                                                                    _ = v105
                                                                                                } else {
                                                                                                    fmt.Println(v2)
                                                                                                    var v106 int = 747
                                                                                                    var v107 float64 = ((-2.25 - (1e10 * ((((-2.25 + float64(1.5)) * 3.14159) + (-2.25 - -2.25)) / (((float64(1.5) + (0.1 / 1.0)) * 3.14159) + 1e10)))) - -2.25)
                                                                                                    fmt.Println(v5, v4, v4)
                                                                                                    if true {
                                                                                                        var v108 uint64 = 719
                                                                                                        v109 := GenS1{F0: -429, F1: 65535, F2: 42}
                                                                                                        var v110 uint8 = (^(v4))
                                                                                                        _ = v108
                                                                                                        _ = v109
                                                                                                        _ = v110
                                                                                                    } else {
                                                                                                        switch v65 {
                                                                                                        case "hello":
                                                                                                            if false {
                                                                                                                v2 += 42
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v106
                                                                                                    _ = v107
                                                                                                }
                                                                                                _ = fi103
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    if ((!true) && true) {
                                                                                        if false {
                                                                                            var v111 uint32 = 1
                                                                                            v65 = v65
                                                                                            if (!(!true)) {
                                                                                                fmt.Println(v111, fi70, v5)
                                                                                                var v112 int64 = 0
                                                                                                _ = v112
                                                                                            } else {
                                                                                                fmt.Println(v111, v65, v4)
                                                                                            }
                                                                                            _ = v111
                                                                                        } else {
                                                                                            v1 = (v1 | v1)
                                                                                            v2 -= v2
                                                                                            v4 -= v4
                                                                                        }
                                                                                        v113 := []int{(fi70 + (^(fi70))), (fi70 << ((65535) & 63)), -9223372036854775808}
                                                                                        _ = v113
                                                                                    } else {
                                                                                        v1 -= v1
                                                                                        v65 = "a"
                                                                                    }
                                                                                    if (((-2.25 - 1.5) < 0.1) || (!((("\n\t" + v65) >= ("" + "foo bar")) && false))) {
                                                                                        var v114 int = 32768
                                                                                        v2 -= (v2 & -128)
                                                                                        var v115 uint = 255
                                                                                        _ = v114
                                                                                        _ = v115
                                                                                    } else {
                                                                                        v1 -= v1
                                                                                    }
                                                                                }
                                                                            case "hello":
                                                                                delete(v3, 79)
                                                                            }
                                                                            v116 := GenS6{F0: v1}
                                                                            _ = v116
                                                                        }
                                                                    }
                                                                }
                                                                _ = fi70
                                                                _ = v71
                                                            }
                                                        }
                                                        _ = v65
                                                    } else {
                                                        switch 0 {
                                                        case 67:
                                                            fmt.Println(v4, v1, v4)
                                                        }
                                                    }
                                                }
                                            case 4294967295:
                                                v4 = 187
                                            case 0:
                                                var v117 uint64 = 42
                                                _ = v117
                                            }
                                        }
                                        for fi118 := 0; fi118 < 4; fi118++ {
                                            v1 -= 1
                                            fmt.Println(v4)
                                            _ = fi118
                                        }
                                    }
                                }
                                v1 -= (v1 / ((v1) | 1))
                                fmt.Println(v2, v4, v5)
                            }
                        case "a":
                            fmt.Println(v5, v2, v2)
                        default:
                            switch 0 {
                            case -894:
                                for fi119 := 0; fi119 < 13; fi119++ {
                                    var v120 uint64 = 0
                                    var v121 bool = (1e10 == (0.1 + 1.5))
                                    _ = fi119
                                    _ = v120
                                    _ = v121
                                }
                            case 32767:
                                for fi122 := 0; fi122 < 6; fi122++ {
                                    v123 := []int{fi122}
                                    _ = fi122
                                    _ = v123
                                }
                            }
                        }
                        _ = v54
                    }
                }
            }
            var v124 int16 = 753
            var v125 int32 = 2147483647
            _ = v124
            _ = v125
        }
        for fi126 := 0; fi126 < 1; fi126++ {
            var v127 int64 = 746
            var v128 string = (("hello" + "") + "a")
            _ = fi126
            _ = v127
            _ = v128
        }
    }
    fmt.Println(v2, v1)
    if false {
        switch "hello" {
        case "\n\t":
            for fi129 := 0; fi129 < 11; fi129++ {
                v5 += -185
                v1 += (v1 & 953)
                if true {
                    if (-724 > 566) {
                        fmt.Println(fi129, v5)
                        v130 := GenS6{F0: 36}
                        fmt.Println(fi129, v1, fi129)
                        _ = v130
                    }
                    v4 = 1
                    v131 := []int{(-(fi129)), (fi129 / ((-747) | 1)), 127, (fi129 | 32768)}
                    _ = v131
                } else {
                    v132 := GenS13{F0: (v1 << ((0) & 63)), F1: v2, F2: 1}
                    switch "hello" {
                    case "hello":
                        var v133 uint = 4294967295
                        _ = v133
                    case "":
                        var v134 int16 = 987
                        _ = v134
                    case "\n\t":
                        v135 := []int{(fi129 ^ 128)}
                        _ = v135
                    }
                    _ = v132
                }
                _ = fi129
            }
        case "a":
            if (990 != 2147483647) {
                if true {
                    v136 := GenS1{F0: 9223372036854775807, F1: 4294967295, F2: -1}
                    if false {
                        v4 = (v4 * (v4 * v4))
                    } else {
                        var v137 string = "hello"
                        _ = v137
                    }
                    _ = v136
                } else {
                    for fi138 := 0; fi138 < 5; fi138++ {
                        v139 := []int{32768}
                        _ = fi138
                        _ = v139
                    }
                }
            } else {
                v140 := []int{-554, 32768, 32767, 128}
                for fi141 := 0; fi141 < 15; fi141++ {
                    switch ("" + "a") {
                    case "\n\t":
                        fmt.Println(v4, v2)
                    case "":
                        switch (^(fi141)) {
                        case 324:
                            for fi142 := 0; fi142 < 9; fi142++ {
                                v2 -= (v2 - v2)
                                v140 = append(v140, (fi141 & (-(fi142))))
                                _ = fi142
                            }
                        }
                    default:
                        if false {
                            v5 = v5
                            v1 = v1
                            if false {
                                fmt.Println(v1, v4)
                                for fi143 := 0; fi143 < 16; fi143++ {
                                    switch (fi143 % ((2147483647) | 1)) {
                                    case -1:
                                        fmt.Println(fi141, fi143, fi143)
                                    case -183:
                                        _ = len(v140)
                                    }
                                    v144 := GenS7{F0: 18446744073709551615, F1: 323, F2: true}
                                    var v145 uint64 = 780
                                    _ = fi143
                                    _ = v144
                                    _ = v145
                                }
                                var v146 int = (fi141 % ((fi141) | 1))
                                _ = v146
                            }
                        } else {
                            if (!false) {
                                _ = len(v140)
                                v1 = (^(v1))
                                if false {
                                    if len(v140) >= 2 {
                                        v140 = v140[1:len(v140)]
                                    }
                                    if true {
                                        var v147 int = fi141
                                        v4 -= v4
                                        var v148 float32 = (1e10 - (0.0 * 3.14159))
                                        _ = v147
                                        _ = v148
                                    }
                                    if true {
                                        switch ("foo bar" + (("a" + "hello") + "hello")) {
                                        case "foo bar":
                                            v6[546] = fi141
                                        case "":
                                            v149 := GenS14{F0: fi141}
                                            _ = v149
                                        case "\n\t":
                                            for fi150 := 0; fi150 < 14; fi150++ {
                                                v4 = v4
                                                _ = fi150
                                            }
                                        }
                                    }
                                }
                            }
                            v1 -= v1
                            _ = len(v3)
                        }
                    }
                    _ = fi141
                }
                _ = v140
            }
        case "":
            v151 := []int{127, 127, 757}
            _ = v151
        }
        if false {
            v1 = v1
        }
        v6[128] = -2147483648
    } else {
        v5 -= (^(v5))
    }
    var v152 int = int(v2)
    for fi153 := 0; fi153 < 2; fi153++ {
        for fi154 := 0; fi154 < 17; fi154++ {
            fmt.Println(fi153)
            v152 += fi153
            _ = fi154
        }
        fmt.Println(v4, v1, v5)
        if (1000 > -1) {
            var v155 string = (("hello" + "\n\t") + "")
            var v156 uint32 = uint32(v4)
            _ = v155
            _ = v156
        }
        _ = fi153
    }
    var v157 uint16 = v1
    v158 := GenS15{F0: 18446744073709551615, F1: 0}
    var v159 bool = (658 != 4294967295)
    for fi160 := 0; fi160 < 18; fi160++ {
        switch ("" + "") {
        case "hello":
            v1 += v157
        case "\n\t":
            v161 := []int{}
            _ = v161
        }
        var v162 uint32 = 383
        if (!false) {
            v163 := GenS16{F0: uint64(v157), F1: "foo bar", F2: (v2 ^ v2)}
            v2 = (v2 ^ v2)
            fmt.Println(v1)
            _ = v163
        } else {
            fmt.Println(fi160)
            v164 := GenS17{F0: v162, F1: false, F2: -1}
            for fi165 := 0; fi165 < 16; fi165++ {
                v2 = -1
                _ = fi165
            }
            _ = v164
        }
        _ = fi160
        _ = v162
    }
    v166 := GenS4{F0: 433}
    v167 := GenS18{F0: (v2 << ((628) & 63))}
    var v168 uint32 = 1
    var v169 int32 = v5
    v4 = v4
    var v170 float32 = 1.5
    var v171 int16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v152:", v152)
    fmt.Println("v157:", v157)
    fmt.Println("v159:", v159)
    fmt.Println("v168:", v168)
    fmt.Println("v169:", v169)
    fmt.Println("v170:", v170)
    fmt.Println("v171:", v171)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v152
    _ = v157
    _ = v158
    _ = v159
    _ = v166
    _ = v167
    _ = v168
    _ = v169
    _ = v170
    _ = v171
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
