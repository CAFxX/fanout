package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 int32; F2 int32 }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 string; F1 uint64; F2 string }
type GenS6 struct { F0 int64; F1 uint64 }
type GenS7 struct { F0 uint64; F1 int }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint64; F1 int32 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int8; F1 int }
type GenS12 struct { F0 uint8; F1 string }
type GenS13 struct { F0 uint8; F1 int16; F2 int }
type GenS14 struct { F0 uint16; F1 uint64; F2 uint }
type GenS15 struct { F0 uint; F1 int16; F2 bool }
type GenS16 struct { F0 int16; F1 uint }
type GenS17 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS18 struct { F0 uint }
type GenS19 struct { F0 bool; F1 int64 }

func main() {
    var v1 int16 = 42
    var v2 int32 = 42
    var v3 string = "a"
    switch 2147483647 {
    case 1:
        v4 := GenS1{F0: 0}
        _ = v4
    default:
        if (!(9223372036854775807 > 928)) {
            if true {
                v3 = ("a" + v3)
                fmt.Println(v3, v1)
                v5 := GenS2{F0: 2}
                _ = v5
            } else {
                v3 = v3
                if (!false) {
                    v3 = v3
                    switch 0 {
                    case 9223372036854775807:
                        var v6 float64 = 0.0
                        _ = v6
                    }
                    v1 = (v1 | 686)
                }
                if false {
                    for fi7 := 0; fi7 < 6; fi7++ {
                        v1 -= v1
                        var v8 int64 = 0
                        _ = fi7
                        _ = v8
                    }
                } else {
                    v9 := GenS3{F0: 2147483648, F1: v2, F2: -2147483648}
                    if ((false || true) && true) {
                        var v10 uint64 = 1
                        if (9223372036854775807 != 42) {
                            var v11 int32 = (v2 * v2)
                            if (!(!(true && ((false && true) && (0.1 == float64(((float32(0.0) / 0.1) / (1.5 / 1.0)))))))) {
                                var v12 float64 = 1.5
                                _ = v12
                            } else {
                                if false {
                                    var v13 string = v3
                                    for fi14 := 0; fi14 < 18; fi14++ {
                                        if true {
                                            var v15 int8 = 42
                                            v15 = v15
                                            _ = v15
                                        } else {
                                            fmt.Println(v13, v2)
                                        }
                                        for fi16 := 0; fi16 < 20; fi16++ {
                                            switch (fi14 >> ((599) & 63)) {
                                            case 42:
                                                fmt.Println(fi16, v1, fi16)
                                            }
                                            var v17 float64 = float64(float32(0.0))
                                            var v18 string = ((v3 + "a") + v3)
                                            _ = fi16
                                            _ = v17
                                            _ = v18
                                        }
                                        _ = fi14
                                    }
                                    _ = v13
                                }
                                v19 := []int{42, 128, -9223372036854775808, 4294967295}
                                _ = v19
                            }
                            _ = v11
                        } else {
                            for fi20 := 0; fi20 < 11; fi20++ {
                                var v21 bool = true
                                if (!(-124 >= -1)) {
                                    v22 := []int{(fi20 * fi20), fi20, fi20}
                                    _ = len(v22)
                                    for fi23 := 0; fi23 < 8; fi23++ {
                                        var v24 uint32 = 65535
                                        v24 -= (v24 - (v24 + (^(v24))))
                                        _ = fi23
                                        _ = v24
                                    }
                                    _ = v22
                                }
                                for fi25 := 0; fi25 < 1; fi25++ {
                                    if true {
                                        var v26 bool = v21
                                        _ = v26
                                    }
                                    switch v3 {
                                    case "":
                                        var v27 string = v3
                                        _ = v27
                                    case "\n\t":
                                        v28 := []int{fi20, fi20}
                                        _ = v28
                                    }
                                    var v29 int32 = v2
                                    v29 = int32(v1)
                                    if (9223372036854775807 <= 234) {
                                        v30 := []int{(fi25 | (fi25 - fi20))}
                                        if v21 {
                                            fmt.Println(v29, v1)
                                        }
                                        _ = v30
                                    }
                                    fmt.Println(v10)
                                    v1 -= 32767
                                    _ = fi25
                                    _ = v29
                                }
                                _ = fi20
                                _ = v21
                            }
                        }
                        _ = v10
                    }
                    v31 := GenS3{F0: 42, F1: v2, F2: v2}
                    _ = v31
                    _ = v9
                }
            }
            v32 := []int{}
            for fi33 := 0; fi33 < 16; fi33++ {
                v1 = -1
                if (42 == 4294967295) {
                    fmt.Println(fi33, v1)
                    if len(v32) > 0 {
                        v32[0] = fi33
                    }
                } else {
                    var v34 string = ((v3 + (v3 + "foo bar")) + v3)
                    if true {
                        switch (v34 + ("" + v3)) {
                        case "\n\t":
                            v32 = append(v32, (fi33 / ((42) | 1)))
                        case "hello":
                            v35 := GenS1{F0: 714}
                            _ = v35
                        }
                    } else {
                        v3 = ""
                    }
                    _ = v34
                }
                v2 -= v2
                _ = fi33
            }
            _ = v32
        } else {
            for fi36 := 0; fi36 < 6; fi36++ {
                switch 32768 {
                case 32767:
                    for fi37 := 0; fi37 < 16; fi37++ {
                        v2 = (-(v2))
                        _ = fi37
                    }
                case 785:
                    v1 -= v1
                case 128:
                    fmt.Println(v3)
                default:
                    switch fi36 {
                    case 127:
                        var v38 string = v3
                        _ = v38
                    case -2147483648:
                        if ((v2 ^ v2) > (v2 + (v2 >> ((67) & 63)))) {
                            var v39 int = (fi36 % ((fi36) | 1))
                            var v40 int8 = 100
                            _ = v39
                            _ = v40
                        } else {
                            v41 := GenS4{F0: 2, F1: 100}
                            switch v3 {
                            case "hello":
                                v3 = v3
                            case "foo bar":
                                v42 := []int{0, fi36, 127}
                                _ = v42
                            default:
                                v43 := GenS3{F0: 73, F1: -601, F2: v2}
                                _ = v43
                            }
                            _ = v41
                        }
                    }
                }
                v44 := GenS3{F0: 65535, F1: v2, F2: -2147483648}
                if (0.1 != -2.25) {
                    if (!false) {
                        fmt.Println(v1, fi36, v3)
                        for fi45 := 0; fi45 < 16; fi45++ {
                            var v46 int = (fi36 * -492)
                            if false {
                                if (35 < 132) {
                                    v47 := GenS5{F0: v3, F1: 762, F2: ""}
                                    v48 := GenS6{F0: 642, F1: 18446744073709551615}
                                    _ = v47
                                    _ = v48
                                }
                            } else {
                                var v49 float64 = float64(((((1e10 + 1.5) * float32((0.1 - 1e10))) - float32(1.5)) * 3.14159))
                                _ = v49
                            }
                            _ = fi45
                            _ = v46
                        }
                    }
                    fmt.Println(v1, v2)
                    v1 = v1
                }
                _ = fi36
                _ = v44
            }
            for fi50 := 0; fi50 < 4; fi50++ {
                for fi51 := 0; fi51 < 13; fi51++ {
                    if (!false) {
                        v52 := GenS7{F0: 4294967295, F1: fi50}
                        _ = v52
                    } else {
                        fmt.Println(v3, v2, fi51)
                        var v53 int16 = (^(v1))
                        _ = v53
                    }
                    fmt.Println(v1)
                    switch v3 {
                    case "a":
                        if false {
                            v1 = (v1 ^ (v1 >> ((2) & 63)))
                            v2 = (v2 * v2)
                        }
                    case "foo bar":
                        var v54 uint64 = 18446744073709551615
                        _ = v54
                    case "\n\t":
                        if ((255 > 0) && false) {
                            var v55 uint64 = 0
                            var v56 int32 = -74
                            if false {
                                if (1 < 554) {
                                    v57 := []int{fi50}
                                    _ = len(v57)
                                    var v58 uint16 = 920
                                    _ = v57
                                    _ = v58
                                }
                            }
                            switch fi50 {
                            case -1:
                                v2 = v56
                            case 574:
                                v59 := GenS3{F0: 0, F1: (v56 + v2), F2: -1}
                                _ = v59
                            }
                            var v60 uint16 = 32767
                            _ = v55
                            _ = v56
                            _ = v60
                        } else {
                            v61 := []int{1, fi50}
                            switch v3 {
                            case "foo bar":
                                switch (("a" + v3) + (v3 + v3)) {
                                case "":
                                    var v62 int32 = -326
                                    v3 = v3
                                    v63 := GenS2{F0: 2}
                                    _ = v62
                                    _ = v63
                                default:
                                    if (v2 != v2) {
                                        if (0 <= 42) {
                                            var v64 uint32 = 1
                                            var v65 uint8 = 127
                                            v66 := GenS8{F0: 503}
                                            _ = v64
                                            _ = v65
                                            _ = v66
                                        } else {
                                            v2 = (^(v2))
                                            switch (fi50 & -107) {
                                            case 178:
                                                v2 += 0
                                            case 42:
                                                var v67 int32 = 42
                                                if false {
                                                    var v68 int64 = -1
                                                    var v69 float64 = 3.14159
                                                    fmt.Println(fi51, v67)
                                                    if len(v61) > 0 {
                                                        v61[0] = (-(fi51))
                                                    }
                                                    _ = v68
                                                    _ = v69
                                                } else {
                                                    _ = len(v61)
                                                    _ = len(v61)
                                                    var v70 uint16 = 892
                                                    _ = v70
                                                }
                                                v71 := GenS6{F0: -9223372036854775808, F1: 42}
                                                fmt.Println(v2, v67)
                                                var v72 int8 = 127
                                                _ = v67
                                                _ = v71
                                                _ = v72
                                            case 128:
                                                if len(v61) > 0 {
                                                    v61[0] = (fi51 ^ 32767)
                                                }
                                                switch (v3 + "hello") {
                                                case "\n\t":
                                                    v1 = v1
                                                case "a":
                                                    var v73 int32 = v2
                                                    _ = v73
                                                }
                                                if len(v61) > 0 {
                                                    v61[0] = fi51
                                                }
                                            default:
                                                var v74 int32 = v2
                                                _ = v74
                                            }
                                        }
                                        v1 += (v1 | (v1 % ((v1) | 1)))
                                        v2 = -303
                                    } else {
                                        var v75 int8 = 0
                                        _ = v75
                                    }
                                }
                            default:
                                if false {
                                    fmt.Println(v1)
                                } else {
                                    switch -925 {
                                    case -9223372036854775808:
                                        switch (fi50 >> ((619) & 63)) {
                                        case 2147483647:
                                            v1 = (v1 << ((42) & 63))
                                        }
                                    case 4294967295:
                                        switch "foo bar" {
                                        case "hello":
                                            v3 = "\n\t"
                                        default:
                                            var v76 uint = 42
                                            _ = v76
                                        }
                                    case 2147483647:
                                        if ((0 == 713) && true) {
                                            v1 = v1
                                            var v77 int8 = int8(v1)
                                            v78 := GenS1{F0: 168}
                                            _ = v77
                                            _ = v78
                                        } else {
                                            v79 := GenS9{F0: 2, F1: v2}
                                            if true {
                                                v2 = (v2 >> ((18446744073709551615) & 63))
                                                switch fi51 {
                                                case -2147483648:
                                                    var v80 uint16 = 65535
                                                    _ = v80
                                                case -952:
                                                    var v81 int32 = -96
                                                    _ = v81
                                                case -258:
                                                    var v82 int32 = int32(v1)
                                                    _ = v82
                                                }
                                            } else {
                                                if true {
                                                    v61 = append(v61, -1)
                                                    v3 = ("hello" + v3)
                                                    fmt.Println(fi51, v1, v2)
                                                }
                                            }
                                            var v83 uint = 2
                                            _ = v79
                                            _ = v83
                                        }
                                    default:
                                        switch "foo bar" {
                                        case "hello":
                                            if true {
                                                if (true || false) {
                                                    var v84 int16 = (v1 / (((v1 * (v1 * (v1 >> ((1) & 63))))) | 1))
                                                    var v85 uint16 = 32767
                                                    _ = v84
                                                    _ = v85
                                                }
                                            }
                                        default:
                                            v2 += (v2 - v2)
                                        }
                                    }
                                }
                            }
                            _ = v61
                        }
                    default:
                        v2 -= v2
                    }
                    fmt.Println(v2, v3)
                    switch (((v3 + v3) + v3) + v3) {
                    case "\n\t":
                        v1 -= (^(v1))
                    }
                    _ = fi51
                }
                var v86 uint8 = 235
                _ = fi50
                _ = v86
            }
        }
    }
    v3 = (v3 + ("\n\t" + ((v3 + "a") + "")))
    v87 := []int{682, 2147483647, 0}
    v3 = "a"
    if false {
        fmt.Println(v3, v2, v1)
        switch 128 {
        case -431:
            if true {
                if true {
                    if len(v87) >= 2 {
                        v87 = v87[1:len(v87)]
                    }
                }
                if ((v1 - v1) == v1) {
                    var v88 int32 = (v2 + (v2 * v2))
                    for fi89 := 0; fi89 < 7; fi89++ {
                        switch v3 {
                        case "\n\t":
                            if true {
                                for fi90 := 0; fi90 < 12; fi90++ {
                                    v2 += v2
                                    _ = fi90
                                }
                            }
                        case "":
                            for fi91 := 0; fi91 < 10; fi91++ {
                                switch (^(fi89)) {
                                case -422:
                                    fmt.Println(v2)
                                case 389:
                                    if false {
                                        v2 = (v88 % ((362) | 1))
                                        var v92 int8 = -3
                                        _ = v92
                                    } else {
                                        var v93 uint64 = 0
                                        switch (fi89 ^ 4294967295) {
                                        case 42:
                                            fmt.Println(fi89)
                                        default:
                                            fmt.Println(v93)
                                        }
                                        v94 := GenS3{F0: 65535, F1: v2, F2: (v88 % ((v88) | 1))}
                                        _ = v93
                                        _ = v94
                                    }
                                case 4294967295:
                                    v88 -= (v2 + v88)
                                }
                                _ = fi91
                            }
                        }
                        switch (v3 + (v3 + v3)) {
                        case "foo bar":
                            if (1000 <= 42) {
                                fmt.Println(v1, v2, fi89)
                            }
                        case "\n\t":
                            for fi95 := 0; fi95 < 3; fi95++ {
                                _ = len(v87)
                                var v96 int64 = -1
                                _ = fi95
                                _ = v96
                            }
                        }
                        _ = fi89
                    }
                    switch (v3 + (v3 + v3)) {
                    case "hello":
                        for fi97 := 0; fi97 < 12; fi97++ {
                            var v98 uint64 = 42
                            for fi99 := 0; fi99 < 3; fi99++ {
                                v100 := GenS10{F0: fi97}
                                fmt.Println(fi99)
                                if len(v87) >= 2 {
                                    v87 = v87[1:len(v87)]
                                }
                                _ = fi99
                                _ = v100
                            }
                            if false {
                                v101 := GenS11{F0: 117, F1: fi97}
                                _ = v101
                            }
                            _ = fi97
                            _ = v98
                        }
                    case "\n\t":
                        var v102 int16 = v1
                        _ = v102
                    case "":
                        v103 := GenS7{F0: 2, F1: -1}
                        _ = v103
                    default:
                        v87 = append(v87, 2147483647)
                    }
                    _ = v88
                }
            }
        }
    }
    for fi104 := 0; fi104 < 18; fi104++ {
        v2 -= v2
        _ = fi104
    }
    v1 -= v1
    v105 := GenS12{F0: 2, F1: (("hello" + v3) + "")}
    if (3.14159 <= float64(((((float32((float64(1e10) - -2.25)) / 1.0) + 1e10) + (1.5 * 3.14159)) + 1e10))) {
        if true {
            switch (v3 + (v3 + "")) {
            case "foo bar":
                var v106 int32 = v2
                _ = v106
            case "\n\t":
                switch 405 {
                case -305:
                    for fi107 := 0; fi107 < 6; fi107++ {
                        if (((0.1 / (-2.25 / float64(-2.25))) != 0.0) || ((v2 | v2) == -1)) {
                            for fi108 := 0; fi108 < 13; fi108++ {
                                var v109 string = (((v3 + (v3 + "a")) + ("\n\t" + "foo bar")) + ("hello" + v3))
                                v109 = "\n\t"
                                v110 := GenS11{F0: 100, F1: fi107}
                                _ = fi108
                                _ = v109
                                _ = v110
                            }
                            if len(v87) > 0 {
                                v87[0] = (fi107 - -682)
                            }
                        }
                        _ = len(v87)
                        _ = fi107
                    }
                }
            case "":
                if true {
                    v111 := GenS13{F0: 214, F1: v1, F2: 4294967295}
                    _ = v111
                } else {
                    var v112 int32 = 1
                    if (128 >= 32767) {
                        if false {
                            v2 = 42
                            if len(v87) > 0 {
                                v87[0] = 2147483647
                            }
                            for fi113 := 0; fi113 < 15; fi113++ {
                                for fi114 := 0; fi114 < 18; fi114++ {
                                    v3 = ("a" + v3)
                                    var v115 bool = (-1 <= 42)
                                    _ = fi114
                                    _ = v115
                                }
                                _ = fi113
                            }
                        } else {
                            fmt.Println(v112, v3)
                            switch "foo bar" {
                            case "a":
                                if false {
                                    if (!false) {
                                        var v116 float64 = ((((((3.14159 * 1.5) - 1e10) - 0.1) / 1.0) / 1.0) * (1.5 * 0.1))
                                        var v117 float32 = (0.1 / 3.14159)
                                        _ = v116
                                        _ = v117
                                    } else {
                                        fmt.Println(v2)
                                        for fi118 := 0; fi118 < 19; fi118++ {
                                            fmt.Println(v1, fi118)
                                            switch (fi118 | fi118) {
                                            case -37:
                                                v2 = (v112 / (((v2 | v112)) | 1))
                                            case 0:
                                                v119 := GenS1{F0: 382}
                                                _ = v119
                                            default:
                                                if false {
                                                    fmt.Println(v112, v2)
                                                    for fi120 := 0; fi120 < 10; fi120++ {
                                                        var v121 int8 = -1
                                                        _ = fi120
                                                        _ = v121
                                                    }
                                                } else {
                                                    v3 = ((v3 + "") + "\n\t")
                                                    v3 = v3
                                                    if true {
                                                        if len(v87) >= 2 {
                                                            v87 = v87[1:len(v87)]
                                                        }
                                                    } else {
                                                        v2 = (-(v2))
                                                        v1 -= (^(v1))
                                                        if true {
                                                            for fi122 := 0; fi122 < 17; fi122++ {
                                                                if true {
                                                                    if (0 > 70) {
                                                                        _ = len(v87)
                                                                    } else {
                                                                        v87 = append(v87, 2147483647)
                                                                        if false {
                                                                            if true {
                                                                                if ((fi122 ^ (fi118 * fi118)) < (fi118 + (fi118 + (^(fi122))))) {
                                                                                    var v123 bool = ((-(fi118)) < (^(fi122)))
                                                                                    v124 := GenS14{F0: 2, F1: 18446744073709551615, F2: 0}
                                                                                    _ = v123
                                                                                    _ = v124
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v1)
                                                                                switch (fi122 / ((fi122) | 1)) {
                                                                                case -718:
                                                                                    var v125 int64 = -852
                                                                                    _ = v125
                                                                                case -802:
                                                                                    if (756 < 2) {
                                                                                        v112 += (v112 | v2)
                                                                                        var v126 string = v3
                                                                                        if (42 != v2) {
                                                                                            var v127 uint = 0
                                                                                            v1 -= (v1 * -32768)
                                                                                            var v128 uint64 = 18446744073709551615
                                                                                            _ = v127
                                                                                            _ = v128
                                                                                        } else {
                                                                                            switch (fi118 - fi118) {
                                                                                            case 42:
                                                                                                _ = len(v87)
                                                                                            case 0:
                                                                                                v3 = v126
                                                                                                v129 := GenS6{F0: -151, F1: 1}
                                                                                                fmt.Println(fi122)
                                                                                                _ = v129
                                                                                            case -9223372036854775808:
                                                                                                v2 = (v2 / ((v112) | 1))
                                                                                            default:
                                                                                                v126 = (v126 + (v126 + v3))
                                                                                            }
                                                                                        }
                                                                                        var v130 uint8 = 236
                                                                                        switch v3 {
                                                                                        case "\n\t":
                                                                                            var v131 uint = 684
                                                                                            _ = v131
                                                                                        default:
                                                                                            v132 := GenS15{F0: 376, F1: (v1 / ((v1) | 1)), F2: ((!(!true)) || false)}
                                                                                            _ = v132
                                                                                        }
                                                                                        _ = v126
                                                                                        _ = v130
                                                                                    } else {
                                                                                        if ((0.1 >= float32((float64(0.0) / 3.14159))) && true) {
                                                                                            v112 = v2
                                                                                            v3 = ""
                                                                                            v133 := GenS1{F0: 42}
                                                                                            if len(v87) >= 2 {
                                                                                                v87 = v87[1:len(v87)]
                                                                                            }
                                                                                            _ = v133
                                                                                        }
                                                                                        v134 := GenS5{F0: (v3 + v3), F1: 0, F2: "a"}
                                                                                        var v135 int32 = 42
                                                                                        _ = v134
                                                                                        _ = v135
                                                                                    }
                                                                                case 128:
                                                                                    if len(v87) > 0 {
                                                                                        v87[0] = (fi118 - fi122)
                                                                                    }
                                                                                default:
                                                                                    if false {
                                                                                        if len(v87) > 0 {
                                                                                            v87[0] = 42
                                                                                        }
                                                                                        if true {
                                                                                            var v136 uint = 1
                                                                                            if len(v87) > 0 {
                                                                                                v87[0] = (fi122 / (((-(fi118))) | 1))
                                                                                            }
                                                                                            v3 = v3
                                                                                            _ = len(v87)
                                                                                            _ = v136
                                                                                        }
                                                                                        fmt.Println(v1)
                                                                                    } else {
                                                                                        v2 += v112
                                                                                    }
                                                                                }
                                                                            }
                                                                            if false {
                                                                                v137 := GenS11{F0: 16, F1: (fi122 >> ((2) & 63))}
                                                                                _ = v137
                                                                            } else {
                                                                                if len(v87) >= 2 {
                                                                                    v87 = v87[1:len(v87)]
                                                                                }
                                                                                var v138 uint16 = 42
                                                                                var v139 uint8 = 1
                                                                                v140 := GenS16{F0: v1, F1: 65535}
                                                                                _ = v138
                                                                                _ = v139
                                                                                _ = v140
                                                                            }
                                                                            switch "a" {
                                                                            case "":
                                                                                var v141 int16 = v1
                                                                                _ = v141
                                                                            case "foo bar":
                                                                                fmt.Println(v2, v1, v3)
                                                                            }
                                                                        }
                                                                        if false {
                                                                            if (!(v2 <= -2147483648)) {
                                                                                v112 = v112
                                                                            } else {
                                                                                _ = len(v87)
                                                                            }
                                                                        }
                                                                    }
                                                                    var v142 int8 = -111
                                                                    _ = v142
                                                                }
                                                                if true {
                                                                    if false {
                                                                        v143 := GenS12{F0: 42, F1: (v3 + v3)}
                                                                        if len(v87) >= 2 {
                                                                            v87 = v87[1:len(v87)]
                                                                        }
                                                                        _ = v143
                                                                    }
                                                                    v3 = v3
                                                                    var v144 int = (fi118 / ((fi122) | 1))
                                                                    _ = v144
                                                                } else {
                                                                    v112 += (^(v112))
                                                                }
                                                                _ = fi122
                                                            }
                                                        } else {
                                                            var v145 int8 = 100
                                                            if false {
                                                                var v146 bool = false
                                                                var v147 float64 = float64((-2.25 + (3.14159 - 0.1)))
                                                                v147 += (v147 * (v147 * v147))
                                                                _ = v146
                                                                _ = v147
                                                            }
                                                            _ = v145
                                                        }
                                                    }
                                                }
                                            }
                                            _ = fi118
                                        }
                                    }
                                    for fi148 := 0; fi148 < 4; fi148++ {
                                        _ = len(v87)
                                        var v149 int32 = (v2 ^ v2)
                                        _ = fi148
                                        _ = v149
                                    }
                                }
                            case "foo bar":
                                v1 -= v1
                            case "":
                                switch (v3 + v3) {
                                case "a":
                                    var v150 float64 = ((-2.25 - ((0.1 / -2.25) + float64((1.5 * -2.25)))) - 0.0)
                                    _ = v150
                                default:
                                    for fi151 := 0; fi151 < 5; fi151++ {
                                        if true {
                                            v87 = append(v87, (fi151 >> ((0) & 63)))
                                            for fi152 := 0; fi152 < 20; fi152++ {
                                                var v153 float32 = (0.1 + -2.25)
                                                _ = fi152
                                                _ = v153
                                            }
                                            for fi154 := 0; fi154 < 18; fi154++ {
                                                v87 = append(v87, fi151)
                                                if len(v87) >= 2 {
                                                    v87 = v87[1:len(v87)]
                                                }
                                                var v155 uint8 = 0
                                                _ = fi154
                                                _ = v155
                                            }
                                        }
                                        _ = fi151
                                    }
                                }
                            }
                        }
                        fmt.Println(v2, v2)
                    }
                    _ = v112
                }
            }
        }
        for fi156 := 0; fi156 < 19; fi156++ {
            for fi157 := 0; fi157 < 9; fi157++ {
                v87 = append(v87, fi157)
                _ = fi157
            }
            _ = fi156
        }
    }
    v158 := GenS11{F0: 100, F1: -1}
    for fi159 := 0; fi159 < 15; fi159++ {
        var v160 float64 = (1e10 + float64(3.14159))
        _ = fi159
        _ = v160
    }
    switch 127 {
    case 1:
        v161 := GenS16{F0: v1, F1: 42}
        _ = v161
    case 4294967295:
        if ((v1 * v1) > 1000) {
            v3 = v3
            switch 9223372036854775807 {
            case 127:
                switch (v3 + v3) {
                case "\n\t":
                    for fi162 := 0; fi162 < 1; fi162++ {
                        for fi163 := 0; fi163 < 3; fi163++ {
                            var v164 bool = false
                            v164 = v164
                            if len(v87) >= 2 {
                                v87 = v87[1:len(v87)]
                            }
                            _ = fi163
                            _ = v164
                        }
                        switch (("" + "a") + v3) {
                        case "":
                            if len(v87) > 0 {
                                v87[0] = fi162
                            }
                        }
                        fmt.Println(v2)
                        _ = fi162
                    }
                case "foo bar":
                    for fi165 := 0; fi165 < 5; fi165++ {
                        var v166 uint16 = 657
                        if (0 != -128) {
                            if false {
                                var v167 uint64 = 4294967295
                                var v168 int32 = 42
                                _ = v167
                                _ = v168
                            }
                        }
                        _ = fi165
                        _ = v166
                    }
                default:
                    switch -1 {
                    case 1:
                        v169 := GenS10{F0: 570}
                        _ = v169
                    case -1:
                        v87 = append(v87, 9223372036854775807)
                    }
                }
            case 32767:
                var v170 int8 = 127
                _ = v170
            case -1:
                fmt.Println(v3, v3)
            default:
                for fi171 := 0; fi171 < 14; fi171++ {
                    v3 = v3
                    var v172 float32 = (3.14159 - (1.5 - float32(-2.25)))
                    var v173 uint8 = 169
                    _ = fi171
                    _ = v172
                    _ = v173
                }
            }
        } else {
            if len(v87) > 0 {
                v87[0] = -9223372036854775808
            }
            v174 := GenS2{F0: 32767}
            fmt.Println(v2, v1, v2)
            _ = v174
        }
    default:
        if ("a" != v3) {
            if ("foo bar" == "\n\t") {
                switch v3 {
                case "a":
                    v175 := GenS17{F0: 65535, F1: 4294967295, F2: 2}
                    _ = v175
                case "foo bar":
                    switch "" {
                    case "":
                        var v176 float32 = (1.5 * 1e10)
                        _ = v176
                    case "foo bar":
                        for fi177 := 0; fi177 < 2; fi177++ {
                            v1 -= -32768
                            if false {
                                if len(v87) >= 2 {
                                    v87 = v87[1:len(v87)]
                                }
                            }
                            v178 := GenS18{F0: 42}
                            _ = fi177
                            _ = v178
                        }
                    }
                default:
                    v3 = v3
                }
            } else {
                for fi179 := 0; fi179 < 8; fi179++ {
                    switch "\n\t" {
                    case "":
                        var v180 uint8 = 127
                        _ = v180
                    default:
                        if (fi179 < fi179) {
                            v2 = -1
                            var v181 float64 = -2.25
                            var v182 float32 = 0.0
                            _ = v181
                            _ = v182
                        }
                    }
                    _ = fi179
                }
                fmt.Println(v3, v2, v3)
            }
            v183 := GenS19{F0: (!(true && false)), F1: 216}
            _ = v183
        }
    }
    fmt.Println(v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v87:", len(v87), v87)
    _ = v1
    _ = v105
    _ = v158
    _ = v2
    _ = v3
    _ = v87
}
