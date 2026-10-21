package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 uint64; F1 string; F2 int8 }
type GenS4 struct { F0 uint16; F1 uint }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int32; F1 uint8 }
type GenS7 struct { F0 int64; F1 int32 }
type GenS8 struct { F0 int32; F1 uint32; F2 int32 }
type GenS9 struct { F0 int64; F1 int32; F2 int64 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int64; F1 uint8 }
type GenS13 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS14 struct { F0 uint64; F1 uint8 }
type GenS15 struct { F0 int64; F1 string }
type GenS16 struct { F0 uint16; F1 uint16; F2 int }
type GenS17 struct { F0 uint8; F1 string }
type GenS18 struct { F0 int32; F1 string }
type GenS19 struct { F0 int8; F1 uint; F2 int8 }

func main() {
    var v1 float64 = -2.25
    var v2 uint64 = 182
    v3 := make(map[uint32]int)
    var v4 int8 = 100
    var v5 int8 = v4
    _ = len(v3)
    v4 = 100
    _ = len(v3)
    v6 := v3[2]
    v7 := GenS1{F0: 255, F1: 1, F2: 25}
    for fi8 := 0; fi8 < 8; fi8++ {
        fmt.Println(v5, v6)
        v6 += (-(fi8))
        if true {
            var v9 uint = 1
            if (770 <= 1) {
                for fi10 := 0; fi10 < 14; fi10++ {
                    v9 = v9
                    fmt.Println(v6, fi8)
                    _ = fi10
                }
                v11 := []int{(-(fi8)), 349, -350, (^(fi8))}
                var v12 int16 = 42
                _ = v11
                _ = v12
            } else {
                var v13 float32 = 0.1
                var v14 bool = false
                _ = v13
                _ = v14
            }
            _ = v9
        }
        _ = fi8
    }
    v2 = v2
    for fi15 := 0; fi15 < 19; fi15++ {
        for fi16 := 0; fi16 < 17; fi16++ {
            var v17 int8 = (v4 ^ v5)
            switch (fi16 * v6) {
            case 127:
                v4 += v4
            default:
                var v18 int = 2147483647
                _ = v18
            }
            if false {
                var v19 uint8 = 251
                switch (v6 | 2147483647) {
                case -9223372036854775808:
                    if (65535 == 255) {
                        v20 := v3[42]
                        v2 += (v2 + 4294967295)
                        if (!false) {
                            switch "foo bar" {
                            case "a":
                                fmt.Println(v5, v19)
                            case "\n\t":
                                v21 := GenS2{F0: false, F1: 1}
                                _ = v21
                            default:
                                fmt.Println(v5)
                            }
                        } else {
                            v1 += v1
                            var v22 uint = 1
                            _ = v22
                        }
                        _ = v20
                    }
                case 9223372036854775807:
                    if false {
                        v23 := GenS3{F0: (v2 % ((789) | 1)), F1: ("\n\t" + ""), F2: (v4 + v5)}
                        v6 = (fi16 >> ((637) & 63))
                        if false {
                            var v24 int16 = 1
                            fmt.Println(v6, v4)
                            if true {
                                fmt.Println(v2, v2, fi15)
                                fmt.Println(v2, v19, v2)
                            }
                            v25 := GenS1{F0: 4294967295, F1: 42, F2: 2}
                            _ = v24
                            _ = v25
                        }
                        _ = v23
                    }
                default:
                    if (1e10 == (((v1 + float64(1e10)) / 0.1) * ((1.5 / v1) - ((0.1 * float64(0.0)) - (v1 / (v1 * 1e10)))))) {
                        var v26 float32 = 3.14159
                        var v27 int = fi16
                        v2 = v2
                        _ = v26
                        _ = v27
                    } else {
                        switch (fi16 ^ v6) {
                        case 32767:
                            if (0.1 < -2.25) {
                                var v28 int = (fi15 / (((-(fi15))) | 1))
                                if false {
                                    v29 := GenS4{F0: uint16(v2), F1: 42}
                                    if true {
                                        switch (fi15 << ((1) & 63)) {
                                        case 919:
                                            var v30 uint32 = 2147483648
                                            _ = v30
                                        }
                                    } else {
                                        v6 = (fi16 & (fi15 | 2147483647))
                                    }
                                    v31 := []int{fi15, 128, (fi16 << ((1) & 63))}
                                    if true {
                                        if (false && (!true)) {
                                            var v32 uint16 = 0
                                            var v33 uint = 779
                                            fmt.Println(fi16)
                                            _ = v32
                                            _ = v33
                                        } else {
                                            var v34 uint8 = 0
                                            _ = v34
                                        }
                                        switch ("foo bar" + (("" + "a") + "a")) {
                                        case "foo bar":
                                            v6 = (-(fi16))
                                        case "a":
                                            fmt.Println(fi16, v4)
                                        }
                                        if true {
                                            var v35 int16 = 51
                                            fmt.Println(v1, v17)
                                            _ = v35
                                        } else {
                                            v3[0] = (v28 + (fi16 + (fi15 >> ((2) & 63))))
                                        }
                                    }
                                    v19 += (v19 % ((237) | 1))
                                    _ = v29
                                    _ = v31
                                } else {
                                    v6 = int(v5)
                                    if false {
                                        var v36 bool = (!((1e10 * 0.1) != (1.5 - 1.5)))
                                        v37 := GenS4{F0: 1, F1: 65535}
                                        if (v17 >= v4) {
                                            if ((!v36) && v36) {
                                                v6 += 4294967295
                                                v36 = ((v19 & v19) <= 0)
                                            }
                                        } else {
                                            var v38 int = -796
                                            var v39 bool = false
                                            var v40 uint32 = 0
                                            _ = v38
                                            _ = v39
                                            _ = v40
                                        }
                                        var v41 uint16 = 1
                                        _ = v36
                                        _ = v37
                                        _ = v41
                                    } else {
                                        switch "a" {
                                        case "":
                                            fmt.Println(fi15, v4, v19)
                                        }
                                        v42 := GenS5{F0: 18446744073709551615}
                                        v43 := v3[0]
                                        _ = v42
                                        _ = v43
                                    }
                                    v44 := []int{938}
                                    _ = v44
                                }
                                v5 += (v5 >> ((18446744073709551615) & 63))
                                _ = v28
                            } else {
                                var v45 uint8 = 255
                                switch (fi16 ^ 32767) {
                                case 203:
                                    var v46 int32 = -1
                                    _ = v46
                                case 0:
                                    fmt.Println(v1, fi15)
                                default:
                                    if (!true) {
                                        if (v17 >= (v4 / ((v4) | 1))) {
                                            v17 -= 0
                                        }
                                    } else {
                                        if ("a" > "foo bar") {
                                            v47 := GenS3{F0: v2, F1: "hello", F2: (v17 + v5)}
                                            _ = v47
                                        } else {
                                            v45 = 255
                                            delete(v3, 0)
                                            var v48 int16 = 1000
                                            v2 -= (^(v2))
                                            if (-9223372036854775808 < 9223372036854775807) {
                                                v49 := GenS6{F0: 1, F1: v45}
                                                v50 := GenS2{F0: (v19 > (v19 >> ((65535) & 63))), F1: 65535}
                                                _ = v49
                                                _ = v50
                                            } else {
                                                var v51 float64 = 0.1
                                                fmt.Println(v19, fi15, v51)
                                                fmt.Println(v45, v17)
                                                _ = v51
                                            }
                                            _ = v48
                                        }
                                        v52 := GenS7{F0: 0, F1: -2147483648}
                                        switch fi15 {
                                        case 32767:
                                            v53 := GenS4{F0: 381, F1: 255}
                                            _ = v53
                                        case -1:
                                            var v54 uint32 = 939
                                            _ = v54
                                        case -9223372036854775808:
                                            v4 -= 100
                                        }
                                        _ = v52
                                    }
                                }
                                _ = v45
                            }
                        case -362:
                            v1 = v1
                        }
                        v55 := v3[0]
                        v56 := GenS8{F0: -2147483648, F1: 1, F2: 166}
                        v57 := GenS6{F0: -1, F1: v19}
                        var v58 uint16 = 2
                        _ = v55
                        _ = v56
                        _ = v57
                        _ = v58
                    }
                }
                _ = v19
            } else {
                var v59 uint64 = 0
                var v60 int64 = 42
                v3[2147483648] = -9223372036854775808
                v61 := GenS8{F0: int32(v5), F1: 4294967295, F2: 1}
                _ = v59
                _ = v60
                _ = v61
            }
            var v62 float64 = v1
            if ("a" <= ("\n\t" + "\n\t")) {
                if false {
                    var v63 bool = false
                    _ = v63
                }
                switch "hello" {
                case "":
                    v64 := GenS9{F0: 346, F1: -2147483648, F2: -9223372036854775808}
                    _ = v64
                case "foo bar":
                    delete(v3, 65535)
                case "\n\t":
                    var v65 uint64 = (v2 * 42)
                    _ = v65
                }
                v4 -= (v4 + -58)
            }
            _ = fi16
            _ = v17
            _ = v62
        }
        var v66 int64 = 722
        _ = fi15
        _ = v66
    }
    fmt.Println(v6, v4)
    var v67 bool = true
    v68 := []int{(v6 / ((v6) | 1)), (v6 << ((255) & 63)), 0, (-(v6))}
    v69 := GenS3{F0: 4294967295, F1: "hello", F2: (v5 % ((v4) | 1))}
    if ((0 == 1) || (((0.0 + -2.25) + -2.25) != ((1e10 - (float32(v1) - float32(v1))) / (((3.14159 / 1.0) + 3.14159) + (0.0 / ((0.1 * 0.0) - 1.5)))))) {
        var v70 uint8 = 127
        for fi71 := 0; fi71 < 15; fi71++ {
            v6 += (v6 | (v6 | fi71))
            for fi72 := 0; fi72 < 5; fi72++ {
                v5 += v4
                switch (v6 ^ fi72) {
                case 2147483647:
                    switch "hello" {
                    case "":
                        switch "\n\t" {
                        case "foo bar":
                            if v67 {
                                if (int64(v70) <= int64(v70)) {
                                    switch (fi72 >> ((42) & 63)) {
                                    case -2147483648:
                                        v73 := GenS6{F0: 316, F1: 240}
                                        _ = v73
                                    default:
                                        if true {
                                            if len(v68) > 0 {
                                                v68[0] = (fi71 >> ((255) & 63))
                                            }
                                            v70 -= (^(v70))
                                            delete(v3, 324)
                                        } else {
                                            if v67 {
                                                if ((fi72 / ((1) | 1)) >= fi72) {
                                                    fmt.Println(v1, v1)
                                                } else {
                                                    var v74 bool = (86 >= 1)
                                                    switch (fi72 | fi72) {
                                                    case 9223372036854775807:
                                                        v75 := GenS10{F0: 732}
                                                        _ = v75
                                                    case 32768:
                                                        if false {
                                                            fmt.Println(fi72, v2, v6)
                                                            _ = len(v68)
                                                            if v67 {
                                                                v76 := GenS9{F0: 0, F1: 0, F2: -811}
                                                                _ = v76
                                                            }
                                                        } else {
                                                            if v67 {
                                                                if len(v68) > 0 {
                                                                    v68[0] = fi72
                                                                }
                                                            } else {
                                                                if (false || (v1 > float64((1.5 * ((((0.0 + (1e10 * ((0.1 * 3.14159) * 1e10))) / 0.1) / (3.14159 / float32((v1 + v1)))) + -2.25))))) {
                                                                    if (818 < 210) {
                                                                        switch "\n\t" {
                                                                        case "foo bar":
                                                                            delete(v3, 2147483647)
                                                                            v70 -= (v70 >> ((98) & 63))
                                                                            v70 -= v70
                                                                            v74 = (65535 <= 0)
                                                                            var v77 uint16 = 0
                                                                            _ = v77
                                                                        case "":
                                                                            var v78 uint16 = 255
                                                                            v5 += v4
                                                                            _ = v78
                                                                        default:
                                                                            var v79 int16 = 1
                                                                            _ = v79
                                                                        }
                                                                        _ = len(v3)
                                                                        var v80 float64 = 1e10
                                                                        _ = v80
                                                                    }
                                                                    v68 = append(v68, (v6 % ((fi71) | 1)))
                                                                } else {
                                                                    var v81 uint16 = 726
                                                                    fmt.Println(v5)
                                                                    _ = v81
                                                                }
                                                            }
                                                        }
                                                    case -76:
                                                        v82 := GenS2{F0: (1 >= 384), F1: 427}
                                                        _ = v82
                                                    default:
                                                        if len(v68) > 0 {
                                                            v68[0] = fi72
                                                        }
                                                    }
                                                    switch -486 {
                                                    case 32767:
                                                        if v67 {
                                                            switch v6 {
                                                            case 361:
                                                                v83 := v3[278]
                                                                _ = v83
                                                            case 128:
                                                                switch v6 {
                                                                case 4294967295:
                                                                    switch (fi71 << ((2) & 63)) {
                                                                    case 42:
                                                                        var v84 float64 = float64((float32((float64(0.0) * (float64(0.0) - float64(3.14159)))) + (float32(float64(-2.25)) + 0.1)))
                                                                        _ = v84
                                                                    }
                                                                case 42:
                                                                    if len(v68) >= 2 {
                                                                        v68 = v68[1:len(v68)]
                                                                    }
                                                                    v2 -= 242
                                                                case -889:
                                                                    v85 := GenS11{F0: v70}
                                                                    _ = v85
                                                                default:
                                                                    var v86 uint16 = 0
                                                                    v87 := GenS8{F0: 2147483647, F1: 280, F2: 42}
                                                                    v88 := GenS1{F0: 1, F1: 1, F2: 620}
                                                                    _ = v86
                                                                    _ = v87
                                                                    _ = v88
                                                                }
                                                            case 127:
                                                                fmt.Println(fi72, v5)
                                                            }
                                                        } else {
                                                            v68 = append(v68, fi71)
                                                        }
                                                    case 128:
                                                        v70 = (v70 * (v70 & v70))
                                                    default:
                                                        var v89 bool = (true && v67)
                                                        _ = v89
                                                    }
                                                    _ = v74
                                                }
                                                if false {
                                                    switch 142 {
                                                    case -1:
                                                        _ = len(v68)
                                                    case -720:
                                                        _ = len(v3)
                                                    case -397:
                                                        switch ("" + "hello") {
                                                        case "foo bar":
                                                            var v90 uint = 2
                                                            _ = v90
                                                        }
                                                    }
                                                    if v67 {
                                                        var v91 uint = 1
                                                        var v92 int64 = -1
                                                        switch "hello" {
                                                        case "hello":
                                                            var v93 string = "foo bar"
                                                            _ = v93
                                                        case "\n\t":
                                                            var v94 float64 = v1
                                                            _ = v94
                                                        case "foo bar":
                                                            v68 = append(v68, 222)
                                                        default:
                                                            v95 := GenS5{F0: v91}
                                                            _ = v95
                                                        }
                                                        _ = v91
                                                        _ = v92
                                                    }
                                                }
                                                v4 -= (^(v4))
                                            } else {
                                                v6 = fi72
                                                v5 += (v5 - v5)
                                                var v96 int64 = -1
                                                _ = v96
                                            }
                                            if v67 {
                                                v5 -= v4
                                                v6 += v6
                                                if len(v68) > 0 {
                                                    v68[0] = (v6 & -876)
                                                }
                                            } else {
                                                _ = len(v3)
                                                if (false || v67) {
                                                    if v67 {
                                                        _ = len(v3)
                                                        v70 -= (v70 * v70)
                                                        var v97 int32 = 1
                                                        _ = v97
                                                    } else {
                                                        _ = len(v3)
                                                        v2 = v2
                                                        v4 += v4
                                                    }
                                                    if ((-32768 == 94) || (!(false && (v67 || true)))) {
                                                        v98 := GenS3{F0: v2, F1: "\n\t", F2: v4}
                                                        _ = v98
                                                    } else {
                                                        switch ("a" + "foo bar") {
                                                        case "\n\t":
                                                            var v99 float32 = 3.14159
                                                            _ = v99
                                                        case "":
                                                            switch (v6 << ((1) & 63)) {
                                                            case -9223372036854775808:
                                                                var v100 bool = v67
                                                                _ = v100
                                                            default:
                                                                v5 += (v5 * v4)
                                                            }
                                                        case "hello":
                                                            fmt.Println(fi72)
                                                        default:
                                                            switch ("" + "") {
                                                            case "\n\t":
                                                                v101 := GenS3{F0: 1, F1: "", F2: v5}
                                                                _ = v101
                                                            case "a":
                                                                var v102 string = "hello"
                                                                _ = v102
                                                            }
                                                        }
                                                        v1 -= v1
                                                        if ((v70 + 1) == (v70 | (v70 % (((v70 | (v70 & 128))) | 1)))) {
                                                            v6 -= -2147483648
                                                            _ = len(v68)
                                                            v6 = v6
                                                        }
                                                        if v67 {
                                                            v103 := GenS11{F0: (^(v70))}
                                                            var v104 uint = 1
                                                            v104 -= v104
                                                            v105 := v3[65535]
                                                            _ = v103
                                                            _ = v104
                                                            _ = v105
                                                        } else {
                                                            var v106 int16 = 1
                                                            _ = v106
                                                        }
                                                        if len(v68) >= 2 {
                                                            v68 = v68[1:len(v68)]
                                                        }
                                                    }
                                                    fmt.Println(v4)
                                                }
                                                switch fi71 {
                                                case 2147483647:
                                                    v6 += fi71
                                                }
                                            }
                                        }
                                    }
                                    var v107 int32 = -769
                                    switch "" {
                                    case "hello":
                                        switch (fi72 >> ((42) & 63)) {
                                        case -311:
                                            fmt.Println(v70, v2)
                                        case 0:
                                            var v108 uint32 = 714
                                            _ = v108
                                        default:
                                            v109 := GenS11{F0: (v70 * v70)}
                                            _ = v109
                                        }
                                    case "\n\t":
                                        var v110 float32 = 1.5
                                        _ = v110
                                    case "foo bar":
                                        var v111 int16 = int16(v107)
                                        _ = v111
                                    default:
                                        fmt.Println(v4, v5, v4)
                                        fmt.Println(v70)
                                    }
                                    v70 = (v70 % ((v70) | 1))
                                    _ = v107
                                } else {
                                    if ((true && false) || ((((-2.25 - (1.5 / 1.0)) + 1.5) - float32((v1 / v1))) != 0.1)) {
                                        switch "foo bar" {
                                        case "a":
                                            if (v67 || (!((v2 + v2) >= (v2 & 2)))) {
                                                v112 := GenS12{F0: 1, F1: (v70 & (v70 * (v70 >> ((119) & 63))))}
                                                switch 0 {
                                                case 2147483647:
                                                    if (uint64(v70) != (v2 << ((2) & 63))) {
                                                        var v113 uint8 = (v70 << ((4294967295) & 63))
                                                        fmt.Println(v6, v67)
                                                        v5 = v5
                                                        _ = v113
                                                    } else {
                                                        v67 = v67
                                                    }
                                                case -1:
                                                    fmt.Println(v1, v2, v1)
                                                }
                                                _ = v112
                                            }
                                        default:
                                            v3[674] = v6
                                        }
                                    } else {
                                        if len(v68) > 0 {
                                            v68[0] = -55
                                        }
                                        if (v67 && true) {
                                            v114 := GenS9{F0: 42, F1: int32(v2), F2: 778}
                                            var v115 int = -9223372036854775808
                                            fmt.Println(v115, v6)
                                            _ = v114
                                            _ = v115
                                        } else {
                                            v116 := v3[42]
                                            fmt.Println(fi72)
                                            v116 += v6
                                            _ = v116
                                        }
                                        v4 = 1
                                    }
                                    var v117 uint8 = (v70 + (v70 ^ (v70 - v70)))
                                    _ = v117
                                }
                                if (255 <= 730) {
                                    v118 := GenS8{F0: 0, F1: 734, F2: 863}
                                    if false {
                                        v119 := GenS5{F0: 39}
                                        v2 = (v2 | 4)
                                        _ = v119
                                    } else {
                                        switch "a" {
                                        case "":
                                            _ = len(v3)
                                        default:
                                            fmt.Println(v2)
                                            delete(v3, 42)
                                            _ = len(v3)
                                            v6 = -2147483648
                                            if (!v67) {
                                                var v120 int16 = 1000
                                                _ = v120
                                            }
                                        }
                                        v121 := GenS9{F0: 42, F1: 42, F2: 0}
                                        _ = v121
                                    }
                                    _ = v118
                                } else {
                                    fmt.Println(v1, v2)
                                    switch "hello" {
                                    case "foo bar":
                                        var v122 float32 = 0.1
                                        _ = v122
                                    case "a":
                                        v123 := GenS13{F0: int16(v70), F1: 4294967295, F2: 42}
                                        _ = v123
                                    }
                                }
                                v124 := GenS3{F0: v2, F1: "a", F2: v5}
                                _ = v124
                            } else {
                                var v125 uint32 = 65535
                                v125 = (v125 - (v125 | v125))
                                if v67 {
                                    var v126 int8 = v5
                                    _ = v126
                                } else {
                                    v4 = (v5 | 42)
                                }
                                _ = v125
                            }
                        default:
                            v3[0] = fi72
                        }
                        fmt.Println(v70, v2, v5)
                        v3[345] = (v6 + (v6 / (((fi72 ^ v6)) | 1)))
                    }
                case 32768:
                    var v127 float64 = (0.1 - v1)
                    _ = v127
                }
                _ = fi72
            }
            _ = fi71
        }
        if v67 {
            fmt.Println(v1, v1)
        } else {
            switch (v6 ^ v6) {
            case -571:
                for fi128 := 0; fi128 < 2; fi128++ {
                    if len(v68) >= 2 {
                        v68 = v68[1:len(v68)]
                    }
                    var v129 int32 = 587
                    _ = fi128
                    _ = v129
                }
            }
            var v130 uint = 42
            _ = v130
        }
        _ = v70
    }
    for fi131 := 0; fi131 < 8; fi131++ {
        if (v1 > v1) {
            fmt.Println(v1, v1, fi131)
            v132 := GenS14{F0: v2, F1: 1}
            _ = v132
        }
        fmt.Println(fi131)
        if (783 == 0) {
            for fi133 := 0; fi133 < 9; fi133++ {
                v5 = v4
                _ = fi133
            }
            switch "a" {
            case "":
                if (77 != uint(v2)) {
                    v67 = (!v67)
                    if false {
                        if (!(int(v2) > (^(v6)))) {
                            v5 = (-(v4))
                            v134 := GenS15{F0: 42, F1: "a"}
                            _ = v134
                        }
                    } else {
                        for fi135 := 0; fi135 < 15; fi135++ {
                            var v136 uint64 = v2
                            if len(v68) >= 2 {
                                v68 = v68[1:len(v68)]
                            }
                            switch v6 {
                            case -1:
                                v67 = (v6 >= (fi135 >> ((65535) & 63)))
                            case -409:
                                fmt.Println(v2)
                            case -323:
                                if true {
                                    v137 := v3[0]
                                    _ = v137
                                } else {
                                    var v138 int16 = 1
                                    v139 := v3[0]
                                    v140 := GenS8{F0: 843, F1: 65535, F2: 714}
                                    _ = v138
                                    _ = v139
                                    _ = v140
                                }
                            }
                            _ = fi135
                            _ = v136
                        }
                    }
                } else {
                    v68 = append(v68, 9223372036854775807)
                    var v141 int64 = 42
                    _ = v141
                }
            case "a":
                v142 := GenS16{F0: 1, F1: 838, F2: (fi131 % (((fi131 + (fi131 ^ fi131))) | 1))}
                _ = v142
            }
        } else {
            switch "foo bar" {
            case "\n\t":
                var v143 string = "a"
                _ = v143
            case "hello":
                var v144 uint = 255
                _ = v144
            default:
                fmt.Println(v4, v1)
            }
            fmt.Println(v1, v2)
            v145 := v3[65535]
            _ = v145
        }
        _ = fi131
    }
    v146 := GenS1{F0: 2, F1: 1, F2: 42}
    v1 = (float64(-2.25) + v1)
    switch -383 {
    case 1:
        v68 = append(v68, v6)
    case 366:
        var v147 int = v6
        _ = v147
    case 4294967295:
        if (255 >= 128) {
            switch v6 {
            case 316:
                var v148 uint = 18446744073709551615
                _ = v148
            case 2147483647:
                v68 = append(v68, -198)
            case 32768:
                v149 := GenS17{F0: 128, F1: ((("foo bar" + ("hello" + "\n\t")) + ((("\n\t" + "a") + "\n\t") + "foo bar")) + "\n\t")}
                _ = v149
            }
            for fi150 := 0; fi150 < 11; fi150++ {
                v151 := GenS18{F0: 2147483647, F1: ""}
                for fi152 := 0; fi152 < 2; fi152++ {
                    v1 += (0.1 * 0.0)
                    _ = fi152
                }
                var v153 uint16 = 255
                _ = fi150
                _ = v151
                _ = v153
            }
        }
    default:
        v154 := GenS19{F0: 0, F1: 0, F2: (v4 >> ((255) & 63))}
        _ = v154
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v67:", v67)
    fmt.Println("slicelen v68:", len(v68), v68)
    _ = v1
    _ = v146
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v67
    _ = v68
    _ = v69
    _ = v7
}
