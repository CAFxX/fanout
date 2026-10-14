package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 bool; F2 bool }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 int; F1 uint8; F2 uint16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int64; F1 int8; F2 string }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int }
type GenS10 struct { F0 string; F1 int32; F2 uint8 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint8; F1 int32 }
type GenS13 struct { F0 int; F1 int32 }
type GenS14 struct { F0 uint; F1 int }
type GenS15 struct { F0 int; F1 uint8 }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 string }
type GenS18 struct { F0 int32; F1 int }
type GenS19 struct { F0 uint32 }
type GenS20 struct { F0 uint64; F1 int8 }
type GenS21 struct { F0 int64 }
func genFn0() int8 {
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 uint16 = 0
        var v8 uint8 = 42
        _ = fi6
        _ = v7
        _ = v8
    }
    switch 1 {
    case 32767:
        var v9 uint64 = 539
        _ = v9
    }
    if ((1.5 / float32(-2.25)) <= -2.25) {
        for fi10 := 0; fi10 < 2; fi10++ {
            v11 := GenS1{F0: -1, F1: 42, F2: 838}
            _ = fi10
            _ = v11
        }
        var v12 int32 = 2147483647
        if true {
            v13 := GenS2{F0: (490 != 61)}
            v14 := []int{2147483647, -1}
            if len(v14) > 0 {
                v14[0] = 705
            }
            _ = v13
            _ = v14
        } else {
            if true {
                for fi15 := 0; fi15 < 11; fi15++ {
                    var v16 bool = true
                    v16 = (-1 >= 1)
                    v17 := GenS3{F0: int(v12), F1: v16, F2: true}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                if (60 > 1) {
                    for fi18 := 0; fi18 < 16; fi18++ {
                        v19 := GenS4{F0: -32768, F1: -32768}
                        if (!true) {
                            v12 -= (v12 % ((v12) | 1))
                            for fi20 := 0; fi20 < 17; fi20++ {
                                var v21 int16 = -32768
                                _ = fi20
                                _ = v21
                            }
                        } else {
                            if ((-(v12)) != (v12 + v12)) {
                                if (1.5 != 1e10) {
                                    var v22 uint32 = 176
                                    v23 := GenS1{F0: 42, F1: 0, F2: 42}
                                    var v24 int8 = 5
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                } else {
                                    var v25 int64 = -9223372036854775808
                                    var v26 bool = true
                                    var v27 string = "\n\t"
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                }
                            } else {
                                v12 -= (v12 ^ v12)
                            }
                            v12 = v12
                            for fi28 := 0; fi28 < 17; fi28++ {
                                v29 := []int{(fi18 / ((fi28) | 1)), (fi18 / (((^(fi18))) | 1)), fi18, fi18}
                                var v30 float64 = 1.5
                                switch (fi18 * fi18) {
                                case 99:
                                    switch ("\n\t" + ("foo bar" + "a")) {
                                    case "\n\t":
                                        var v31 float32 = ((-2.25 - float32((v30 - -2.25))) * 1.5)
                                        _ = v31
                                    case "":
                                        var v32 int = fi28
                                        _ = v32
                                    }
                                case 1:
                                    var v33 float64 = v30
                                    v12 += (v12 - -292)
                                    v33 = (v30 * 1.5)
                                    if ((65 <= 82) && (v12 > 157)) {
                                        switch (-(fi28)) {
                                        case 4294967295:
                                            var v34 int32 = -2147483648
                                            _ = v34
                                        case 1:
                                            if true {
                                                v35 := GenS4{F0: 32767, F1: 0}
                                                v12 += v12
                                                v36 := GenS2{F0: (468 > 922)}
                                                _ = v35
                                                _ = v36
                                            }
                                        case 594:
                                            switch fi18 {
                                            case 32768:
                                                var v37 int64 = 1
                                                _ = v37
                                            case 32767:
                                                var v38 uint32 = 421
                                                _ = v38
                                            default:
                                                var v39 uint8 = 0
                                                _ = v39
                                            }
                                        default:
                                            v33 = v33
                                        }
                                        var v40 uint8 = 255
                                        if true {
                                            if true {
                                                if false {
                                                    switch fi28 {
                                                    case 1:
                                                        if true {
                                                            if (!true) {
                                                                v12 += (v12 ^ v12)
                                                            }
                                                        } else {
                                                            v41 := GenS4{F0: int16(v12), F1: 351}
                                                            var v42 string = "foo bar"
                                                            switch "hello" {
                                                            case "\n\t":
                                                                if true {
                                                                    v30 = -2.25
                                                                } else {
                                                                    v43 := GenS4{F0: -1, F1: 1}
                                                                    var v44 uint32 = 335
                                                                    _ = v43
                                                                    _ = v44
                                                                }
                                                            case "":
                                                                var v45 int32 = (v12 - (v12 << ((1) & 63)))
                                                                _ = v45
                                                            case "a":
                                                                if len(v29) >= 2 {
                                                                    v29 = v29[1:len(v29)]
                                                                }
                                                            default:
                                                                var v46 string = v42
                                                                _ = v46
                                                            }
                                                            _ = v41
                                                            _ = v42
                                                        }
                                                    case 4294967295:
                                                        switch fi18 {
                                                        case 42:
                                                            var v47 uint16 = 860
                                                            _ = v47
                                                        case 32767:
                                                            if (!false) {
                                                                v33 = (float64((1.5 - (-2.25 + -2.25))) * v30)
                                                                var v48 float64 = (((1e10 - v30) - (((((((float64(float32((v30 - v33))) + (float64(0.1) * -2.25)) + 1.5) - 1e10) - (v33 + (0.1 / v30))) / (v33 / (3.14159 + v30))) / v30) * v33)) - 0.0)
                                                                _ = v48
                                                            }
                                                        default:
                                                            if true {
                                                                if len(v29) >= 2 {
                                                                    v29 = v29[1:len(v29)]
                                                                }
                                                                if (false && true) {
                                                                    var v49 string = "foo bar"
                                                                    _ = v49
                                                                }
                                                            } else {
                                                                _ = len(v29)
                                                            }
                                                        }
                                                    default:
                                                        v50 := GenS5{F0: fi18, F1: 0, F2: 65535}
                                                        _ = v50
                                                    }
                                                    if false {
                                                        v33 -= float64(0.1)
                                                        v30 += -2.25
                                                    } else {
                                                        var v51 string = ((("a" + "") + ("" + "")) + "hello")
                                                        v12 = (v12 ^ -767)
                                                        _ = v51
                                                    }
                                                } else {
                                                    if false {
                                                        if (false || false) {
                                                            v52 := GenS6{F0: 2147483648}
                                                            _ = v52
                                                        }
                                                        v12 = v12
                                                    } else {
                                                        v12 = v12
                                                        var v53 float64 = 1.5
                                                        _ = v53
                                                    }
                                                    var v54 uint8 = (v40 & v40)
                                                    fmt.Println(v12, v33)
                                                    if (uint64(v54) < 1) {
                                                        var v55 uint64 = 18446744073709551615
                                                        var v56 int8 = 1
                                                        _ = v55
                                                        _ = v56
                                                    } else {
                                                        fmt.Println(v12)
                                                    }
                                                    v30 -= v30
                                                    _ = v54
                                                }
                                                v29 = append(v29, (-(fi18)))
                                                var v57 float32 = float32((0.0 - 1e10))
                                                _ = v57
                                            } else {
                                                var v58 int64 = -1
                                                switch ("\n\t" + "") {
                                                case "a":
                                                    v33 = ((((((((1.5 - 1e10) / 0.1) - -2.25) * -2.25) * 0.0) + float64((1e10 / 1.0))) + 3.14159) * v30)
                                                }
                                                _ = v58
                                            }
                                        } else {
                                            if true {
                                                v33 = 0.0
                                            }
                                            switch "" {
                                            case "foo bar":
                                                var v59 int = fi28
                                                _ = v59
                                            case "a":
                                                switch "hello" {
                                                case "hello":
                                                    v33 -= -2.25
                                                default:
                                                    if len(v29) > 0 {
                                                        v29[0] = 42
                                                    }
                                                }
                                            }
                                            v40 += 128
                                        }
                                        _ = v40
                                    } else {
                                        var v60 string = "foo bar"
                                        switch -1 {
                                        case 32768:
                                            fmt.Println(v12, fi28, fi18)
                                        case -758:
                                            var v61 uint8 = 255
                                            _ = v61
                                        case 32767:
                                            var v62 int8 = -1
                                            _ = v62
                                        }
                                        _ = v60
                                    }
                                    if (32768 != fi18) {
                                        if len(v29) >= 2 {
                                            v29 = v29[1:len(v29)]
                                        }
                                        if true {
                                            v63 := GenS7{F0: -1, F1: 42, F2: ""}
                                            _ = v63
                                        } else {
                                            v64 := GenS3{F0: fi28, F1: (545 >= 316), F2: false}
                                            v30 = (1.5 / (((1.5 + (v33 - float64(3.14159))) / 1.0) * v30))
                                            _ = v64
                                        }
                                    } else {
                                        fmt.Println(v30, v30, fi18)
                                        var v65 int32 = v12
                                        switch fi18 {
                                        case -2147483648:
                                            if len(v29) > 0 {
                                                v29[0] = 725
                                            }
                                        }
                                        if len(v29) >= 2 {
                                            v29 = v29[1:len(v29)]
                                        }
                                        var v66 int32 = v12
                                        _ = v65
                                        _ = v66
                                    }
                                    _ = v33
                                case 32767:
                                    var v67 float32 = (float32((v30 * v30)) * (0.1 - (0.1 * 0.0)))
                                    _ = v67
                                }
                                _ = fi28
                                _ = v29
                                _ = v30
                            }
                        }
                        _ = fi18
                        _ = v19
                    }
                    for fi68 := 0; fi68 < 1; fi68++ {
                        v69 := GenS8{F0: (true && (9223372036854775807 == -9223372036854775808))}
                        v12 += 2147483647
                        v70 := []int{fi68, (fi68 << ((18446744073709551615) & 63))}
                        _ = fi68
                        _ = v69
                        _ = v70
                    }
                } else {
                    var v71 int16 = -1
                    switch "" {
                    case "hello":
                        var v72 string = (("\n\t" + "foo bar") + "")
                        _ = v72
                    }
                    if false {
                        var v73 int32 = 519
                        var v74 string = ("hello" + "a")
                        switch ("a" + v74) {
                        case "a":
                            for fi75 := 0; fi75 < 16; fi75++ {
                                for fi76 := 0; fi76 < 3; fi76++ {
                                    var v77 int8 = 100
                                    v78 := GenS8{F0: true}
                                    switch (("foo bar" + (v74 + "hello")) + "a") {
                                    case "hello":
                                        switch v74 {
                                        case "foo bar":
                                            var v79 uint8 = 2
                                            _ = v79
                                        }
                                    }
                                    _ = fi76
                                    _ = v77
                                    _ = v78
                                }
                                v71 = v71
                                switch v74 {
                                case "a":
                                    v80 := []int{int(v71), fi75, 0}
                                    _ = v80
                                case "hello":
                                    v81 := GenS5{F0: (fi75 | (^(fi75))), F1: 202, F2: 37}
                                    _ = v81
                                case "foo bar":
                                    v71 -= (v71 ^ -32768)
                                }
                                _ = fi75
                            }
                        case "hello":
                            switch (v74 + v74) {
                            case "":
                                for fi82 := 0; fi82 < 4; fi82++ {
                                    for fi83 := 0; fi83 < 6; fi83++ {
                                        v84 := []int{-1}
                                        if true {
                                            var v85 uint64 = 947
                                            var v86 float32 = (1e10 / 0.1)
                                            _ = v85
                                            _ = v86
                                        }
                                        _ = fi83
                                        _ = v84
                                    }
                                    _ = fi82
                                }
                            case "\n\t":
                                v74 = v74
                            case "foo bar":
                                v12 = v12
                            }
                        }
                        _ = v73
                        _ = v74
                    } else {
                        switch "\n\t" {
                        case "a":
                            if (!(((((1e10 - float64(3.14159)) - -2.25) * 1.5) > 0.0) || false)) {
                                fmt.Println(v12, v71)
                            } else {
                                if false {
                                    switch "foo bar" {
                                    case "\n\t":
                                        switch 96 {
                                        case -210:
                                            switch "hello" {
                                            case "foo bar":
                                                var v87 int16 = (^(v71))
                                                _ = v87
                                            default:
                                                v88 := []int{}
                                                _ = v88
                                            }
                                        case -258:
                                            var v89 string = "\n\t"
                                            _ = v89
                                        case -469:
                                            v71 += v71
                                        }
                                    default:
                                        v90 := []int{-267}
                                        _ = v90
                                    }
                                }
                            }
                        case "":
                            v91 := GenS8{F0: (!false)}
                            _ = v91
                        case "\n\t":
                            switch 127 {
                            case -670:
                                v92 := GenS3{F0: -274, F1: true, F2: (-800 < (v71 >> ((305) & 63)))}
                                _ = v92
                            case -525:
                                if false {
                                    var v93 uint = 144
                                    if (311 > 255) {
                                        var v94 uint32 = 2
                                        if true {
                                            var v95 uint32 = 65535
                                            _ = v95
                                        } else {
                                            if false {
                                                for fi96 := 0; fi96 < 14; fi96++ {
                                                    v97 := []int{-1, int(v93), (fi96 << ((v93) & 63))}
                                                    _ = fi96
                                                    _ = v97
                                                }
                                            }
                                        }
                                        v12 += v12
                                        _ = v94
                                    } else {
                                        v98 := []int{}
                                        switch "hello" {
                                        case "\n\t":
                                            for fi99 := 0; fi99 < 18; fi99++ {
                                                for fi100 := 0; fi100 < 11; fi100++ {
                                                    var v101 int32 = -2147483648
                                                    _ = fi100
                                                    _ = v101
                                                }
                                                v71 = v71
                                                switch fi99 {
                                                case 32768:
                                                    for fi102 := 0; fi102 < 10; fi102++ {
                                                        v103 := GenS9{F0: -2147483648}
                                                        var v104 string = ""
                                                        if true {
                                                            if ((v93 % ((18446744073709551615) | 1)) <= v93) {
                                                                v105 := GenS10{F0: ("" + "foo bar"), F1: -1, F2: 42}
                                                                if (!(true || false)) {
                                                                    var v106 uint = (v93 - 44)
                                                                    if (915 <= 1) {
                                                                        _ = len(v98)
                                                                    }
                                                                    _ = v106
                                                                }
                                                                if false {
                                                                    if (!(true && false)) {
                                                                        v93 -= (v93 ^ (^(v93)))
                                                                        var v107 uint64 = 18446744073709551615
                                                                        if (966 > (v93 * (v93 - v93))) {
                                                                            v108 := GenS11{F0: v93}
                                                                            var v109 float64 = -2.25
                                                                            fmt.Println(v12)
                                                                            _ = v108
                                                                            _ = v109
                                                                        }
                                                                        _ = v107
                                                                    } else {
                                                                        var v110 uint8 = 1
                                                                        _ = v110
                                                                    }
                                                                    if (255 < 230) {
                                                                        if false {
                                                                            fmt.Println(fi99, v71, fi99)
                                                                        }
                                                                        fmt.Println(v104)
                                                                    }
                                                                }
                                                                _ = v105
                                                            }
                                                            v93 -= v93
                                                            v93 = 2
                                                        } else {
                                                            if (-2.25 > -2.25) {
                                                                var v111 string = (v104 + v104)
                                                                if len(v98) >= 2 {
                                                                    v98 = v98[1:len(v98)]
                                                                }
                                                                var v112 int8 = 1
                                                                _ = v111
                                                                _ = v112
                                                            } else {
                                                                v113 := GenS8{F0: false}
                                                                _ = v113
                                                            }
                                                            v104 = ("hello" + "\n\t")
                                                        }
                                                        _ = fi102
                                                        _ = v103
                                                        _ = v104
                                                    }
                                                case 127:
                                                    v12 -= (v12 | 2147483647)
                                                case 1:
                                                    if (18446744073709551615 == 1) {
                                                        var v114 uint8 = 2
                                                        for fi115 := 0; fi115 < 3; fi115++ {
                                                            v114 = (^(v114))
                                                            var v116 uint16 = 439
                                                            var v117 int8 = 1
                                                            _ = fi115
                                                            _ = v116
                                                            _ = v117
                                                        }
                                                        fmt.Println(v93, v93, v12)
                                                        _ = v114
                                                    }
                                                default:
                                                    v118 := GenS8{F0: true}
                                                    _ = v118
                                                }
                                                _ = fi99
                                            }
                                        case "foo bar":
                                            var v119 float32 = 1e10
                                            _ = v119
                                        case "a":
                                            v12 = (v12 << (((v93 + (^(v93)))) & 63))
                                        default:
                                            if (false || true) {
                                                var v120 int32 = (v12 >> (((v93 ^ (v93 << ((42) & 63)))) & 63))
                                                var v121 uint16 = 605
                                                v120 += (v120 ^ v120)
                                                _ = v120
                                                _ = v121
                                            } else {
                                                fmt.Println(v12)
                                            }
                                        }
                                        _ = v98
                                    }
                                    _ = v93
                                } else {
                                    if ((true || (42 > 255)) || (1e10 <= float64((0.0 * (0.1 * ((0.0 / float32(3.14159)) / 1e10)))))) {
                                        switch ("" + "") {
                                        case "hello":
                                            for fi122 := 0; fi122 < 2; fi122++ {
                                                if (false || (!(((((128 >= 255) && ((!false) || ((2147483648 >= 2147483647) || true))) && false) && false) || true))) {
                                                    if true {
                                                        for fi123 := 0; fi123 < 11; fi123++ {
                                                            v124 := []int{(fi122 >> ((192) & 63)), 32767, fi122, 32768}
                                                            if (90 > 42) {
                                                                var v125 uint64 = 42
                                                                var v126 int32 = (v12 & (^(v12)))
                                                                switch (fi123 << ((65535) & 63)) {
                                                                case 4294967295:
                                                                    var v127 uint64 = (v125 | 42)
                                                                    _ = v127
                                                                case 42:
                                                                    switch "hello" {
                                                                    case "":
                                                                        var v128 float32 = -2.25
                                                                        _ = v128
                                                                    case "a":
                                                                        _ = len(v124)
                                                                    case "foo bar":
                                                                        if true {
                                                                            var v129 uint32 = 2
                                                                            _ = v129
                                                                        }
                                                                        var v130 int64 = 0
                                                                        v131 := GenS12{F0: 255, F1: v126}
                                                                        if (-114 > 122) {
                                                                            switch (fi123 & fi122) {
                                                                            case 1:
                                                                                fmt.Println(v71, fi123)
                                                                            }
                                                                            var v132 uint64 = 1
                                                                            v124 = append(v124, fi122)
                                                                            v71 = v71
                                                                            _ = v132
                                                                        }
                                                                        _ = v130
                                                                        _ = v131
                                                                    default:
                                                                        v125 = v125
                                                                    }
                                                                default:
                                                                    if (-701 > 9223372036854775807) {
                                                                        v133 := GenS2{F0: (0.0 <= 1e10)}
                                                                        if len(v124) > 0 {
                                                                            v124[0] = 127
                                                                        }
                                                                        _ = v133
                                                                    }
                                                                }
                                                                if false {
                                                                    switch (fi123 - (fi123 >> ((255) & 63))) {
                                                                    case -2147483648:
                                                                        fmt.Println(fi122, v126)
                                                                    case 32767:
                                                                        if (255 < 846) {
                                                                            var v134 uint64 = 2
                                                                            _ = v134
                                                                        } else {
                                                                            fmt.Println(fi122, v126)
                                                                        }
                                                                    default:
                                                                        fmt.Println(fi123, fi123, fi122)
                                                                    }
                                                                    if false {
                                                                        v126 += int32(v125)
                                                                        if len(v124) > 0 {
                                                                            v124[0] = 2147483647
                                                                        }
                                                                    }
                                                                }
                                                                v126 -= v12
                                                                _ = v125
                                                                _ = v126
                                                            }
                                                            var v135 int32 = 82
                                                            if (false || (true || false)) {
                                                                v135 = 2147483647
                                                                if false {
                                                                    switch ("" + ("\n\t" + ("hello" + "a"))) {
                                                                    case "a":
                                                                        v136 := GenS5{F0: (fi122 << ((42) & 63)), F1: 255, F2: 283}
                                                                        _ = v136
                                                                    default:
                                                                        v137 := GenS13{F0: fi123, F1: (v12 + (v135 & v12))}
                                                                        _ = v137
                                                                    }
                                                                    switch "" {
                                                                    case "hello":
                                                                        if len(v124) > 0 {
                                                                            v124[0] = 127
                                                                        }
                                                                    case "a":
                                                                        var v138 int32 = -85
                                                                        _ = v138
                                                                    default:
                                                                        if len(v124) > 0 {
                                                                            v124[0] = (fi123 + (-(fi122)))
                                                                        }
                                                                    }
                                                                }
                                                                switch (fi123 | 9223372036854775807) {
                                                                case 127:
                                                                    v124 = append(v124, (fi123 - (^(fi123))))
                                                                    v71 = (v71 * (v71 >> ((18446744073709551615) & 63)))
                                                                    fmt.Println(fi123, fi122, v12)
                                                                case 42:
                                                                    var v139 int16 = v71
                                                                    _ = v139
                                                                default:
                                                                    fmt.Println(v12, v135, fi122)
                                                                }
                                                            } else {
                                                                v124 = append(v124, 2147483647)
                                                            }
                                                            _ = fi123
                                                            _ = v124
                                                            _ = v135
                                                        }
                                                        if false {
                                                            var v140 uint = 18446744073709551615
                                                            _ = v140
                                                        }
                                                    } else {
                                                        if true {
                                                            var v141 int8 = 0
                                                            _ = v141
                                                        }
                                                        v71 = (v71 - (-(v71)))
                                                        var v142 uint64 = 635
                                                        _ = v142
                                                    }
                                                } else {
                                                    if (float64(((float32(0.0) + -2.25) * 0.0)) <= 1.5) {
                                                        fmt.Println(v12, v71)
                                                        if true {
                                                            for fi143 := 0; fi143 < 4; fi143++ {
                                                                v12 = (v12 | v12)
                                                                v144 := GenS6{F0: 9}
                                                                _ = fi143
                                                                _ = v144
                                                            }
                                                            if ("" > ("\n\t" + "a")) {
                                                                v71 += (v71 >> ((674) & 63))
                                                                for fi145 := 0; fi145 < 1; fi145++ {
                                                                    v12 += (-(v12))
                                                                    _ = fi145
                                                                }
                                                            }
                                                            var v146 uint = 2
                                                            _ = v146
                                                        } else {
                                                            switch ("hello" + ("a" + "\n\t")) {
                                                            case "hello":
                                                                fmt.Println(v12, fi122)
                                                            case "foo bar":
                                                                if false {
                                                                    var v147 bool = ((100 < 42) && (726 > 572))
                                                                    _ = v147
                                                                }
                                                            }
                                                            if (true || true) {
                                                                switch ("" + "foo bar") {
                                                                case "":
                                                                    fmt.Println(fi122)
                                                                case "a":
                                                                    v148 := GenS1{F0: -1, F1: (v71 & -1), F2: -1}
                                                                    _ = v148
                                                                default:
                                                                    switch (("\n\t" + "foo bar") + "hello") {
                                                                    case "a":
                                                                        var v149 uint32 = 2
                                                                        _ = v149
                                                                    case "\n\t":
                                                                        fmt.Println(v12, fi122)
                                                                    default:
                                                                        v150 := GenS9{F0: fi122}
                                                                        _ = v150
                                                                    }
                                                                }
                                                                fmt.Println(fi122)
                                                                v151 := GenS8{F0: (!false)}
                                                                _ = v151
                                                            }
                                                        }
                                                    }
                                                }
                                                _ = fi122
                                            }
                                        }
                                    } else {
                                        v152 := []int{}
                                        v71 = v71
                                        _ = v152
                                    }
                                    v12 -= v12
                                }
                            case 127:
                                switch 4294967295 {
                                case 128:
                                    fmt.Println(v71)
                                case 42:
                                    var v153 float64 = (0.0 + (3.14159 - 1e10))
                                    _ = v153
                                default:
                                    v154 := GenS14{F0: 1, F1: 9223372036854775807}
                                    _ = v154
                                }
                            default:
                                var v155 int32 = (^(v12))
                                _ = v155
                            }
                        default:
                            v156 := []int{251, 128}
                            _ = v156
                        }
                        for fi157 := 0; fi157 < 8; fi157++ {
                            for fi158 := 0; fi158 < 4; fi158++ {
                                fmt.Println(fi157, v71)
                                fmt.Println(v71, v71, fi158)
                                var v159 int16 = (v71 / ((v71) | 1))
                                if false {
                                    v160 := GenS12{F0: 166, F1: 112}
                                    _ = v160
                                }
                                v161 := GenS15{F0: (fi157 << ((65535) & 63)), F1: 1}
                                _ = fi158
                                _ = v159
                                _ = v161
                            }
                            _ = fi157
                        }
                    }
                    _ = v71
                }
            }
            var v162 uint64 = 1
            if true {
                v163 := []int{-2147483648, 127, -703}
                _ = v163
            } else {
                v164 := []int{959, -9223372036854775808, -2147483648, 9223372036854775807}
                for fi165 := 0; fi165 < 11; fi165++ {
                    v12 = v12
                    var v166 bool = ((-32768 > 273) || ((("" == "") && (128 >= 1)) || false))
                    _ = fi165
                    _ = v166
                }
                _ = v164
            }
            _ = v162
        }
        _ = v12
    }
    var v167 int64 = -96
    var v168 int16 = 519
    _ = v167
    _ = v168
    return -4
}
func genFn1() bool {
    for fi169 := 0; fi169 < 10; fi169++ {
        switch (fi169 >> ((18446744073709551615) & 63)) {
        case -2147483648:
            fmt.Println(fi169)
        case 32768:
            for fi170 := 0; fi170 < 3; fi170++ {
                v171 := []int{fi169, fi170}
                _ = fi170
                _ = v171
            }
        case -9223372036854775808:
            var v172 float64 = 3.14159
            _ = v172
        }
        _ = fi169
    }
    for fi173 := 0; fi173 < 13; fi173++ {
        fmt.Println(fi173)
        switch 9223372036854775807 {
        case 127:
            v174 := GenS16{F0: -2147483648}
            _ = v174
        case 99:
            for fi175 := 0; fi175 < 19; fi175++ {
                v176 := GenS17{F0: "a"}
                var v177 string = "hello"
                var v178 int32 = -2147483648
                _ = fi175
                _ = v176
                _ = v177
                _ = v178
            }
        case 42:
            var v179 int16 = 32767
            _ = v179
        }
        _ = fi173
    }
    var v180 int8 = 1
    var v181 int = 195
    var v182 uint16 = 394
    _ = v180
    _ = v181
    _ = v182
    return true
}

func main() {
    var v1 uint32 = 2147483648
    v2 := []string{"foo bar", ("" + ("\n\t" + "\n\t")), "", "a"}
    v3 := make(map[int]int)
    v4 := make(map[int]uint32)
    var v5 int64 = 1
    _ = len(v2)
    v1 += v1
    fmt.Println(v1, v5)
    switch ("a" + "hello") {
    case "a":
        var v183 float64 = 0.0
        _ = v183
    case "hello":
        var v184 uint8 = 47
        _ = v184
    default:
        for fi185 := 0; fi185 < 2; fi185++ {
            _ = len(v2)
            if (!true) {
                v1 -= 167
            }
            v5 = -337
            _ = fi185
        }
    }
    if (false || (!("\n\t" > "a"))) {
        v5 = 42
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v186 int16 = 0
        _ = v186
    }
    var v187 int32 = -1
    if ((!(!(v5 < v5))) && (2 <= (v1 & (v1 % (((v1 ^ (v1 | v1))) | 1))))) {
        switch "hello" {
        case "":
            v188 := GenS7{F0: v5, F1: -58, F2: "\n\t"}
            _ = v188
        case "foo bar":
            if false {
                if false {
                    v189 := GenS18{F0: (v187 * (v187 * v187)), F1: 0}
                    v190 := v4[-1]
                    if false {
                        var v191 string = "hello"
                        v192 := GenS19{F0: v1}
                        _ = v191
                        _ = v192
                    }
                    _ = v189
                    _ = v190
                }
                switch 128 {
                case 90:
                    v193 := GenS20{F0: 4294967295, F1: -128}
                    _ = v193
                }
                switch ("hello" + "a") {
                case "foo bar":
                    v5 -= (^(v5))
                case "\n\t":
                    delete(v4, 1)
                }
            }
        }
        v5 = (v5 << ((2) & 63))
        _ = len(v2)
    } else {
        if (false && false) {
            fmt.Println(v5, v5)
        }
    }
    fmt.Println(v5, v187, v1)
    var v194 int8 = 16
    for fi195 := 0; fi195 < 10; fi195++ {
        var v196 float64 = 1.5
        v187 = (v187 - v187)
        _ = fi195
        _ = v196
    }
    v187 = (v187 % ((35) | 1))
    for fi197 := 0; fi197 < 18; fi197++ {
        for fi198 := 0; fi198 < 12; fi198++ {
            v194 = v194
            if (189 >= 236) {
                fmt.Println(fi197, fi198)
            } else {
                switch 32768 {
                case 0:
                    switch "\n\t" {
                    case "\n\t":
                        v199 := v3[(fi197 ^ int(v1))]
                        _ = v199
                    }
                case 881:
                    v194 = (v194 * genFn0())
                case 128:
                    v1 = (v1 >> ((1) & 63))
                }
            }
            _ = fi198
        }
        v200 := GenS21{F0: (v5 << ((259) & 63))}
        _ = fi197
        _ = v200
    }
    v5 = 0
    var v201 uint64 = 915
    fmt.Println("v1:", v1)
    fmt.Println("v187:", v187)
    fmt.Println("v194:", v194)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v201:", v201)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v187
    _ = v194
    _ = v2
    _ = v201
    _ = v3
    _ = v4
    _ = v5
}
