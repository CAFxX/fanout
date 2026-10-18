package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 uint8; F2 uint32 }
type GenS6 struct { F0 int32; F1 uint16; F2 int16 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 string; F1 bool }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint; F1 int32 }
type GenS12 struct { F0 bool; F1 uint }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 uint32; F1 int; F2 uint64 }
type GenS15 struct { F0 bool; F1 int32 }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 int32; F1 uint64 }
type GenS18 struct { F0 int64 }
type GenS19 struct { F0 int64 }
type GenS20 struct { F0 int16; F1 int64 }
type GenS21 struct { F0 bool; F1 int }
type GenS22 struct { F0 bool }
type GenS23 struct { F0 bool }
type GenS24 struct { F0 uint; F1 bool }
type GenS25 struct { F0 int; F1 uint64 }
type GenS26 struct { F0 bool }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 503
    var v3 uint32 = 4294967295
    v4 := make(map[uint32]bool)
    v5 := make(map[string]int)
    var v6 int64 = -1
    var v7 int32 = 191
    v3 = v3
    v8 := GenS1{F0: "hello"}
    v9 := GenS1{F0: "\n\t"}
    switch 581 {
    case -2147483648:
        if false {
            switch ("foo bar" + ("hello" + "hello")) {
            case "foo bar":
                var v10 int32 = v7
                _ = v10
            case "":
                v11 := GenS1{F0: "\n\t"}
                _ = v11
            case "\n\t":
                switch 952 {
                case 36:
                    for fi12 := 0; fi12 < 13; fi12++ {
                        v3 = v3
                        _ = fi12
                    }
                case -1:
                    if (127 > 128) {
                        v13 := GenS2{F0: 127, F1: -128}
                        _ = v13
                    } else {
                        v2 = 168
                        for fi14 := 0; fi14 < 19; fi14++ {
                            v15 := GenS2{F0: 255, F1: 1}
                            switch "a" {
                            case "a":
                                switch ("hello" + "\n\t") {
                                case "":
                                    switch "" {
                                    case "hello":
                                        fmt.Println(v3, fi14, v6)
                                    case "":
                                        var v16 uint = 255
                                        _ = v16
                                    default:
                                        v3 = (v3 << ((uint(v2)) & 63))
                                    }
                                case "hello":
                                    var v17 bool = (false && false)
                                    _ = v17
                                case "foo bar":
                                    v6 = (^(v6))
                                }
                            case "hello":
                                _ = len(v1)
                            default:
                                var v18 uint64 = 18446744073709551615
                                _ = v18
                            }
                            _ = fi14
                            _ = v15
                        }
                        v2 -= v2
                    }
                case 714:
                    v1[uint32(v2)] = v3
                default:
                    v3 -= 0
                }
            }
            fmt.Println(v3, v2, v3)
        }
    case 32768:
        v2 = 795
    case 42:
        switch "\n\t" {
        case "":
            for fi19 := 0; fi19 < 15; fi19++ {
                var v20 int32 = (^(v7))
                switch "hello" {
                case "a":
                    for fi21 := 0; fi21 < 14; fi21++ {
                        var v22 uint64 = 4294967295
                        _ = fi21
                        _ = v22
                    }
                case "hello":
                    var v23 uint16 = 0
                    _ = v23
                case "\n\t":
                    v20 -= v20
                default:
                    v1[833] = (^(v3))
                }
                v24 := GenS1{F0: ""}
                _ = fi19
                _ = v20
                _ = v24
            }
        case "hello":
            if true {
                for fi25 := 0; fi25 < 11; fi25++ {
                    for fi26 := 0; fi26 < 5; fi26++ {
                        v7 -= 2147483647
                        v7 -= 658
                        v27 := GenS2{F0: 255, F1: 1}
                        _ = fi26
                        _ = v27
                    }
                    switch ("a" + "hello") {
                    case "foo bar":
                        v3 += 978
                    default:
                        switch fi25 {
                        case 127:
                            fmt.Println(v3)
                        default:
                            if (!false) {
                                var v28 bool = ((1.5 - 0.0) >= 1.5)
                                if ((false && v28) || true) {
                                    v2 = v2
                                    switch (fi25 * fi25) {
                                    case 652:
                                        fmt.Println(v6)
                                    default:
                                        var v29 float64 = 0.0
                                        _ = v29
                                    }
                                    v28 = ((765 == 1) && true)
                                }
                                _ = v28
                            } else {
                                switch "foo bar" {
                                case "a":
                                    v30 := GenS1{F0: ("\n\t" + "a")}
                                    _ = v30
                                case "hello":
                                    v31 := GenS3{F0: 32767}
                                    _ = v31
                                }
                                fmt.Println(v7, v6, v2)
                            }
                        }
                    }
                    if (!true) {
                        switch -1 {
                        case 154:
                            v6 = (v6 + (v6 * v6))
                        case 32768:
                            if false {
                                switch fi25 {
                                case 0:
                                    if false {
                                        if true {
                                            for fi32 := 0; fi32 < 6; fi32++ {
                                                var v33 uint = 4294967295
                                                v3 = (v3 >> (((v33 & 0)) & 63))
                                                switch ("hello" + "\n\t") {
                                                case "hello":
                                                    v7 = (v7 << (((^(v33))) & 63))
                                                case "":
                                                    var v34 bool = true
                                                    v35 := []int{-9223372036854775808}
                                                    _ = v34
                                                    _ = v35
                                                case "a":
                                                    var v36 int16 = 837
                                                    _ = v36
                                                }
                                                var v37 float64 = 1.5
                                                _ = fi32
                                                _ = v33
                                                _ = v37
                                            }
                                            v38 := GenS3{F0: int16(v6)}
                                            _ = v38
                                        } else {
                                            for fi39 := 0; fi39 < 11; fi39++ {
                                                var v40 uint32 = 4294967295
                                                var v41 string = ""
                                                _ = fi39
                                                _ = v40
                                                _ = v41
                                            }
                                            v42 := GenS4{F0: (^(v3))}
                                            if false {
                                                var v43 uint64 = 2
                                                var v44 int64 = (v6 + (v6 % ((-9223372036854775808) | 1)))
                                                var v45 uint = uint(v2)
                                                _ = v43
                                                _ = v44
                                                _ = v45
                                            }
                                            _ = v42
                                        }
                                        for fi46 := 0; fi46 < 9; fi46++ {
                                            delete(v1, (v3 ^ (^(v3))))
                                            _ = fi46
                                        }
                                    } else {
                                        if true {
                                            var v47 uint32 = 65535
                                            var v48 float64 = ((((float64(((0.1 - 1.5) + float32((3.14159 - 0.0)))) / 0.1) - 3.14159) * float64(-2.25)) - (3.14159 * -2.25))
                                            switch int(v7) {
                                            case 0:
                                                v49 := []int{fi25, (fi25 / ((fi25) | 1)), (fi25 + fi25)}
                                                _ = v49
                                            case 9223372036854775807:
                                                var v50 uint8 = 128
                                                _ = v50
                                            case 679:
                                                fmt.Println(v7)
                                            default:
                                                v7 = (v7 >> ((543) & 63))
                                            }
                                            _ = v47
                                            _ = v48
                                        }
                                    }
                                case 4294967295:
                                    switch (fi25 % ((800) | 1)) {
                                    case 32768:
                                        v51 := GenS5{F0: 2147483647, F1: uint8(v3), F2: 65535}
                                        _ = v51
                                    case -1:
                                        switch "hello" {
                                        case "":
                                            fmt.Println(v6)
                                        default:
                                            switch (("hello" + "") + ("\n\t" + "")) {
                                            case "a":
                                                switch (^(fi25)) {
                                                case -536:
                                                    var v52 uint16 = v2
                                                    _ = v52
                                                case 32767:
                                                    v53 := GenS5{F0: fi25, F1: 1, F2: 977}
                                                    _ = v53
                                                case -248:
                                                    v54 := GenS6{F0: (^(v7)), F1: (v2 >> ((865) & 63)), F2: -1}
                                                    _ = v54
                                                }
                                            case "\n\t":
                                                v2 += 93
                                            case "foo bar":
                                                var v55 bool = false
                                                _ = v55
                                            default:
                                                var v56 uint8 = 209
                                                _ = v56
                                            }
                                        }
                                    default:
                                        v57 := GenS7{F0: 675}
                                        _ = v57
                                    }
                                case 127:
                                    if true {
                                        var v58 float32 = (3.14159 + 0.1)
                                        if ("hello" == ("\n\t" + "\n\t")) {
                                            switch "a" {
                                            case "\n\t":
                                                switch -931 {
                                                case -127:
                                                    for fi59 := 0; fi59 < 7; fi59++ {
                                                        var v60 string = "foo bar"
                                                        var v61 float32 = 0.1
                                                        _ = fi59
                                                        _ = v60
                                                        _ = v61
                                                    }
                                                case -2147483648:
                                                    if (!(!false)) {
                                                        v62 := GenS6{F0: v7, F1: 32767, F2: 572}
                                                        for fi63 := 0; fi63 < 9; fi63++ {
                                                            _ = len(v4)
                                                            v64 := GenS8{F0: fi63}
                                                            _ = fi63
                                                            _ = v64
                                                        }
                                                        _ = v62
                                                    } else {
                                                        v58 += ((v58 + v58) * 0.0)
                                                    }
                                                case 4294967295:
                                                    var v65 int32 = (^(v7))
                                                    _ = v65
                                                }
                                            }
                                        }
                                        _ = v58
                                    }
                                }
                                var v66 uint64 = 2
                                _ = v66
                            } else {
                                switch ("a" + "hello") {
                                case "":
                                    switch fi25 {
                                    case -469:
                                        fmt.Println(v6, v2, v6)
                                    case -1:
                                        if false {
                                            var v67 string = "a"
                                            if (1 != 0) {
                                                v68 := GenS6{F0: -2147483648, F1: (v2 - (v2 & (v2 / (((v2 / ((v2) | 1))) | 1)))), F2: 90}
                                                if (true || (222 >= 2147483647)) {
                                                    var v69 int = fi25
                                                    var v70 uint = 0
                                                    if false {
                                                        switch 2147483647 {
                                                        case 2147483647:
                                                            if false {
                                                                if false {
                                                                    switch int(v3) {
                                                                    case 42:
                                                                        var v71 uint32 = v3
                                                                        _ = v71
                                                                    case 9223372036854775807:
                                                                        v3 = v3
                                                                    case 32767:
                                                                        if true {
                                                                            for fi72 := 0; fi72 < 1; fi72++ {
                                                                                fmt.Println(v67, fi72)
                                                                                var v73 uint = (v70 << (((v70 & 18446744073709551615)) & 63))
                                                                                v74 := GenS9{F0: "\n\t", F1: false}
                                                                                _ = fi72
                                                                                _ = v73
                                                                                _ = v74
                                                                            }
                                                                            if (-1 != 127) {
                                                                                v70 -= (v70 & uint(v3))
                                                                                v75 := GenS7{F0: (^(v3))}
                                                                                _ = v75
                                                                            }
                                                                        } else {
                                                                            if ((v6 / (((v6 * int64(v2))) | 1)) < -935) {
                                                                                if false {
                                                                                    v76 := GenS1{F0: ("foo bar" + (v67 + "hello"))}
                                                                                    _ = v76
                                                                                }
                                                                                v77 := GenS2{F0: 42, F1: 42}
                                                                                _ = v77
                                                                            } else {
                                                                                if (true || (109 > 127)) {
                                                                                    var v78 int = fi25
                                                                                    var v79 int64 = v6
                                                                                    _ = v78
                                                                                    _ = v79
                                                                                } else {
                                                                                    v7 -= (v7 << ((42) & 63))
                                                                                }
                                                                                v3 += 760
                                                                                v80 := GenS5{F0: v69, F1: 42, F2: 2147483648}
                                                                                _ = v80
                                                                            }
                                                                            var v81 float32 = 1e10
                                                                            v82 := v1[v3]
                                                                            _ = v81
                                                                            _ = v82
                                                                        }
                                                                    default:
                                                                        for fi83 := 0; fi83 < 9; fi83++ {
                                                                            v70 += 0
                                                                            _ = fi83
                                                                        }
                                                                    }
                                                                    v2 = (v2 | (v2 - (^(v2))))
                                                                } else {
                                                                    var v84 int64 = (v6 + v6)
                                                                    _ = v84
                                                                }
                                                                var v85 uint8 = 2
                                                                fmt.Println(v2)
                                                                _ = v85
                                                            } else {
                                                                fmt.Println(v3, v67, v7)
                                                                v86 := GenS4{F0: v3}
                                                                v87 := []int{v69, -163, -2147483648, fi25}
                                                                _ = v86
                                                                _ = v87
                                                            }
                                                        default:
                                                            v88 := GenS6{F0: v7, F1: (v2 ^ 0), F2: 42}
                                                            _ = v88
                                                        }
                                                        if (!((v70 <= v70) && (((1e10 * (0.1 - 0.0)) / 1.0) < ((float32(0.1) - 1.5) / 3.14159)))) {
                                                            v89 := GenS2{F0: 1, F1: 0}
                                                            _ = v89
                                                        }
                                                    }
                                                    _ = v69
                                                    _ = v70
                                                } else {
                                                    switch "a" {
                                                    case "":
                                                        var v90 uint32 = v3
                                                        _ = v90
                                                    default:
                                                        v91 := GenS10{F0: 659}
                                                        _ = v91
                                                    }
                                                    v67 = ("hello" + (v67 + v67))
                                                }
                                                _ = v68
                                            }
                                            switch "hello" {
                                            case "":
                                                fmt.Println(v7, v67, v6)
                                            case "foo bar":
                                                for fi92 := 0; fi92 < 20; fi92++ {
                                                    if (v3 > 824) {
                                                        switch (fi92 >> ((0) & 63)) {
                                                        case 127:
                                                            var v93 uint = 1
                                                            _ = v93
                                                        }
                                                        v94 := GenS11{F0: 90, F1: (v7 / ((v7) | 1))}
                                                        var v95 uint8 = 2
                                                        switch (fi92 - fi25) {
                                                        case 4294967295:
                                                            if (false || (-9223372036854775808 < (v6 % (((v6 / ((v6) | 1))) | 1)))) {
                                                                var v96 bool = false
                                                                _ = v96
                                                            } else {
                                                                if ((^(v6)) < int64(v7)) {
                                                                    if true {
                                                                        v2 = (v2 / ((v2) | 1))
                                                                        if false {
                                                                            v95 -= v95
                                                                            if (false || (v3 <= 558)) {
                                                                                var v97 float32 = 1.5
                                                                                delete(v4, (v3 * v3))
                                                                                _ = v97
                                                                            }
                                                                            v98 := GenS7{F0: v3}
                                                                            if true {
                                                                                if (!false) {
                                                                                    v95 = v95
                                                                                    fmt.Println(v3, v7, fi25)
                                                                                }
                                                                            }
                                                                            _ = v98
                                                                        } else {
                                                                            if (-63 > 84) {
                                                                                fmt.Println(fi92, v67, v7)
                                                                            } else {
                                                                                var v99 int16 = -32768
                                                                                v100 := []int{}
                                                                                _ = len(v100)
                                                                                _ = v100
                                                                                _ = v99
                                                                            }
                                                                            v101 := GenS8{F0: fi25}
                                                                            _ = v101
                                                                        }
                                                                    }
                                                                    v95 = (v95 + (v95 * v95))
                                                                    if true {
                                                                        v3 += uint32(v95)
                                                                    }
                                                                } else {
                                                                    var v102 int = 132
                                                                    switch v67 {
                                                                    case "":
                                                                        var v103 uint64 = 333
                                                                        _ = v103
                                                                    case "a":
                                                                        v5["foo bar"] = (v102 ^ (v102 / ((fi25) | 1)))
                                                                    }
                                                                    var v104 float64 = 0.1
                                                                    _ = v102
                                                                    _ = v104
                                                                }
                                                                if true {
                                                                    v1[v3] = (^(v3))
                                                                    v3 -= v3
                                                                } else {
                                                                    var v105 uint64 = 1
                                                                    v106 := GenS1{F0: "\n\t"}
                                                                    v107 := v5[v67]
                                                                    v108 := GenS9{F0: v67, F1: ((-2.25 - 1e10) == 0.1)}
                                                                    _ = v105
                                                                    _ = v106
                                                                    _ = v107
                                                                    _ = v108
                                                                }
                                                                v2 = v2
                                                            }
                                                        }
                                                        var v109 int8 = 42
                                                        _ = v109
                                                        _ = v94
                                                        _ = v95
                                                    } else {
                                                        v6 = v6
                                                    }
                                                    v67 = v67
                                                    v3 = 42
                                                    _ = fi92
                                                }
                                            }
                                            _ = v67
                                        } else {
                                            for fi110 := 0; fi110 < 14; fi110++ {
                                                fmt.Println(v6)
                                                v111 := []int{(fi25 - (fi25 * (fi110 ^ fi110))), (fi25 + 32767), fi110, (fi25 >> ((460) & 63))}
                                                if (781 <= uint(v3)) {
                                                    v3 -= (v3 ^ v3)
                                                    if true {
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            var v112 bool = ((1.5 / 3.14159) != -2.25)
                                                            _ = v112
                                                        case "hello":
                                                            var v113 uint8 = 250
                                                            _ = v113
                                                        }
                                                        var v114 int32 = (v7 % (((^(v7))) | 1))
                                                        if len(v111) > 0 {
                                                            v111[0] = fi25
                                                        }
                                                        _ = v114
                                                    } else {
                                                        if true {
                                                            if (true && (("foo bar" + "hello") != ("foo bar" + ((("a" + ("foo bar" + ("foo bar" + "\n\t"))) + "") + ((("" + "a") + "") + "hello"))))) {
                                                                var v115 uint16 = v2
                                                                var v116 int32 = (v7 % (((v7 / (((v7 % (((v7 * (v7 / ((1) | 1)))) | 1))) | 1))) | 1))
                                                                _ = v115
                                                                _ = v116
                                                            } else {
                                                                var v117 string = ("a" + ("\n\t" + "a"))
                                                                if len(v111) > 0 {
                                                                    v111[0] = fi110
                                                                }
                                                                _ = v117
                                                            }
                                                        }
                                                        v118 := GenS12{F0: (false || (!false)), F1: 65535}
                                                        v119 := GenS13{F0: (v2 >> ((255) & 63))}
                                                        _ = v118
                                                        _ = v119
                                                    }
                                                    if ((!false) && false) {
                                                        v3 += (v3 * (^(v3)))
                                                        var v120 bool = (v6 > (v6 << ((255) & 63)))
                                                        _ = v120
                                                    } else {
                                                        v121 := GenS14{F0: v3, F1: (fi25 + fi110), F2: 180}
                                                        var v122 int16 = 824
                                                        _ = v121
                                                        _ = v122
                                                    }
                                                } else {
                                                    v6 -= (v6 >> ((717) & 63))
                                                    switch "" {
                                                    case "":
                                                        v2 -= (v2 - 143)
                                                    case "foo bar":
                                                        var v123 float64 = 0.0
                                                        _ = v123
                                                    case "\n\t":
                                                        var v124 uint64 = 2
                                                        _ = v124
                                                    default:
                                                        switch "" {
                                                        case "":
                                                            v125 := GenS15{F0: (false || false), F1: (v7 | v7)}
                                                            if false {
                                                                if false {
                                                                    v4[(v3 % ((v3) | 1))] = (false || (181 <= (v2 & (v2 << ((65535) & 63)))))
                                                                    if (!false) {
                                                                        v2 -= 32767
                                                                    } else {
                                                                        if (139 >= 18446744073709551615) {
                                                                            if true {
                                                                                v111 = append(v111, int(v6))
                                                                                fmt.Println(v7)
                                                                            }
                                                                        } else {
                                                                            switch 1 {
                                                                            case 128:
                                                                                _ = len(v5)
                                                                            }
                                                                        }
                                                                    }
                                                                    v126 := GenS13{F0: v2}
                                                                    fmt.Println(v3, fi110)
                                                                    _ = v126
                                                                }
                                                            } else {
                                                                v6 = (^(v6))
                                                                fmt.Println(v7, fi110, fi25)
                                                                switch 837 {
                                                                case 4294967295:
                                                                    v3 = (v3 % ((v3) | 1))
                                                                case 2147483647:
                                                                    v7 = v7
                                                                    v7 = 42
                                                                    v111 = append(v111, int(v2))
                                                                case -186:
                                                                    if true {
                                                                        v127 := v4[(v3 / (((v3 * v3)) | 1))]
                                                                        switch (fi25 * fi110) {
                                                                        case 4294967295:
                                                                            v128 := GenS1{F0: (("foo bar" + (("foo bar" + "a") + "a")) + "a")}
                                                                            _ = v128
                                                                        case -946:
                                                                            var v129 int = (fi25 * -488)
                                                                            _ = v129
                                                                        case 964:
                                                                            var v130 float32 = 3.14159
                                                                            _ = v130
                                                                        }
                                                                        var v131 uint32 = 965
                                                                        _ = v127
                                                                        _ = v131
                                                                    }
                                                                default:
                                                                    switch "foo bar" {
                                                                    case "foo bar":
                                                                        delete(v5, "")
                                                                    case "hello":
                                                                        delete(v4, (v3 - 309))
                                                                    case "":
                                                                        if len(v111) >= 2 {
                                                                            v111 = v111[1:len(v111)]
                                                                        }
                                                                    default:
                                                                        if len(v111) >= 2 {
                                                                            v111 = v111[1:len(v111)]
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            v7 += -2147483648
                                                            _ = v125
                                                        default:
                                                            var v132 float32 = 0.1
                                                            _ = v132
                                                        }
                                                    }
                                                }
                                                _ = fi110
                                                _ = v111
                                            }
                                            var v133 int8 = 0
                                            _ = v133
                                        }
                                    default:
                                        var v134 int = (^(fi25))
                                        _ = v134
                                    }
                                case "\n\t":
                                    var v135 float64 = -2.25
                                    _ = v135
                                }
                            }
                        case 2147483647:
                            for fi136 := 0; fi136 < 5; fi136++ {
                                delete(v1, v3)
                                _ = fi136
                            }
                        default:
                            var v137 uint = 18446744073709551615
                            _ = v137
                        }
                        for fi138 := 0; fi138 < 1; fi138++ {
                            fmt.Println(fi138, v6)
                            v139 := GenS16{F0: v6}
                            _ = fi138
                            _ = v139
                        }
                    } else {
                        _ = len(v1)
                        var v140 int64 = v6
                        _ = v140
                    }
                    _ = fi25
                }
                fmt.Println(v3)
                for fi141 := 0; fi141 < 10; fi141++ {
                    var v142 uint64 = 1
                    var v143 int8 = 1
                    if (-2.25 < ((1.5 * ((1.5 * (0.0 * (1e10 + 0.1))) - 0.1)) - ((1e10 * -2.25) - 3.14159))) {
                        if (((fi141 ^ fi141) >= fi141) || (!(1 < v3))) {
                            for fi144 := 0; fi144 < 1; fi144++ {
                                var v145 int16 = -32768
                                fmt.Println(fi141, v143, v145)
                                if true {
                                    v3 += v3
                                    if true {
                                        if (fi141 <= (fi144 ^ fi144)) {
                                            switch (fi144 * -2147483648) {
                                            case -991:
                                                if true {
                                                    v143 += v143
                                                } else {
                                                    var v146 uint32 = (v3 >> ((255) & 63))
                                                    var v147 int = 480
                                                    _ = v146
                                                    _ = v147
                                                }
                                            }
                                            var v148 uint32 = (v3 >> ((384) & 63))
                                            delete(v4, v3)
                                            _ = v148
                                        } else {
                                            v149 := GenS17{F0: (v7 & v7), F1: (v142 ^ (v142 << ((65535) & 63)))}
                                            _ = v149
                                        }
                                        v3 = 769
                                        v150 := GenS5{F0: (fi141 << ((725) & 63)), F1: 128, F2: uint32(v142)}
                                        if true {
                                            v7 = (v7 ^ v7)
                                            v143 += (v143 & v143)
                                            if ((v6 ^ 475) >= (v6 >> ((889) & 63))) {
                                                var v151 int64 = 42
                                                var v152 uint16 = v2
                                                switch -256 {
                                                case -1:
                                                    switch "a" {
                                                    case "a":
                                                        var v153 uint32 = v3
                                                        _ = v153
                                                    case "":
                                                        v154 := GenS18{F0: 48}
                                                        _ = v154
                                                    case "\n\t":
                                                        v155 := []int{(fi144 / (((fi144 | -2147483648)) | 1))}
                                                        _ = v155
                                                    default:
                                                        v156 := GenS10{F0: 138}
                                                        _ = v156
                                                    }
                                                case 9223372036854775807:
                                                    if ((float32(float64(0.0)) < (1e10 / 1.0)) || (true || (v7 == 2147483647))) {
                                                        delete(v1, v3)
                                                        v6 = v6
                                                    }
                                                case -15:
                                                    fmt.Println(v151)
                                                }
                                                if (v7 > v7) {
                                                    fmt.Println(v145, v7)
                                                    var v157 int8 = -47
                                                    _ = v157
                                                }
                                                _ = v151
                                                _ = v152
                                            }
                                        }
                                        _ = v150
                                    } else {
                                        switch ("foo bar" + "a") {
                                        case "":
                                            if (float64(float32(0.1)) <= 0.1) {
                                                var v158 int64 = v6
                                                _ = v158
                                            } else {
                                                v159 := GenS19{F0: -1}
                                                _ = v159
                                            }
                                        }
                                    }
                                }
                                if ((v2 % ((v2) | 1)) >= 0) {
                                    v160 := GenS20{F0: v145, F1: 42}
                                    _ = v160
                                } else {
                                    var v161 int32 = (v7 * v7)
                                    v2 += (^(v2))
                                    _ = v161
                                }
                                v7 = v7
                                _ = fi144
                                _ = v145
                            }
                        } else {
                            if true {
                                v162 := GenS19{F0: 42}
                                v143 += v143
                                fmt.Println(v143, v2)
                                _ = v162
                            }
                        }
                        for fi163 := 0; fi163 < 6; fi163++ {
                            fmt.Println(v3, v3, v142)
                            v164 := []int{fi141, 42, (fi163 % (((fi141 ^ 9223372036854775807)) | 1)), -1}
                            var v165 uint64 = v142
                            _ = fi163
                            _ = v164
                            _ = v165
                        }
                    } else {
                        var v166 float32 = 0.1
                        v7 -= int32(v142)
                        v167 := GenS10{F0: 321}
                        _ = v166
                        _ = v167
                    }
                    _ = fi141
                    _ = v142
                    _ = v143
                }
            } else {
                if true {
                    for fi168 := 0; fi168 < 5; fi168++ {
                        v169 := []int{-9223372036854775808}
                        delete(v1, v3)
                        var v170 uint = 923
                        _ = fi168
                        _ = v169
                        _ = v170
                    }
                } else {
                    switch "\n\t" {
                    case "a":
                        if false {
                            var v171 uint16 = (v2 + v2)
                            _ = v171
                        } else {
                            v2 -= v2
                            if false {
                                var v172 int32 = v7
                                switch -2147483648 {
                                case 0:
                                    v173 := GenS21{F0: (((true || (!false)) || true) && (58 == (v2 * 1))), F1: 0}
                                    _ = v173
                                case 127:
                                    var v174 uint32 = 174
                                    _ = v174
                                case -665:
                                    fmt.Println(v3)
                                default:
                                    if (true || true) {
                                        v175 := GenS19{F0: (^(v6))}
                                        if false {
                                            v176 := GenS22{F0: true}
                                            var v177 string = ""
                                            _ = len(v5)
                                            _ = v176
                                            _ = v177
                                        }
                                        _ = v175
                                    }
                                }
                                _ = v172
                            }
                            v178 := []int{4294967295, 2147483647, -1}
                            _ = v178
                        }
                    case "foo bar":
                        v179 := GenS5{F0: 32767, F1: 253, F2: v3}
                        _ = v179
                    case "":
                        if ("hello" != ("foo bar" + "foo bar")) {
                            for fi180 := 0; fi180 < 11; fi180++ {
                                if (!false) {
                                    for fi181 := 0; fi181 < 16; fi181++ {
                                        v3 -= v3
                                        _ = fi181
                                    }
                                } else {
                                    fmt.Println(v2, v6, v3)
                                }
                                v182 := []int{9223372036854775807}
                                _ = fi180
                                _ = v182
                            }
                        } else {
                            var v183 uint32 = v3
                            if false {
                                var v184 int32 = v7
                                _ = v184
                            } else {
                                for fi185 := 0; fi185 < 19; fi185++ {
                                    for fi186 := 0; fi186 < 16; fi186++ {
                                        v7 = (v7 & (v7 | int32(v6)))
                                        fmt.Println(v2, v2, v3)
                                        _ = fi186
                                    }
                                    var v187 int16 = 1
                                    _ = fi185
                                    _ = v187
                                }
                                v3 += (v3 >> ((743) & 63))
                            }
                            _ = v183
                        }
                    default:
                        v2 += 0
                    }
                    v188 := v1[80]
                    _ = v188
                }
            }
        case "foo bar":
            v3 -= v3
        }
    default:
        var v189 int16 = 32767
        _ = v189
    }
    switch "" {
    case "":
        v190 := GenS23{F0: false}
        _ = v190
    case "hello":
        var v191 float32 = 1e10
        _ = v191
    case "foo bar":
        for fi192 := 0; fi192 < 7; fi192++ {
            delete(v1, v3)
            _ = fi192
        }
    default:
        for fi193 := 0; fi193 < 13; fi193++ {
            for fi194 := 0; fi194 < 4; fi194++ {
                switch fi193 {
                case -9223372036854775808:
                    v1[v3] = v3
                case 42:
                    v195 := GenS24{F0: 42, F1: (-165 < 1000)}
                    _ = v195
                }
                switch "\n\t" {
                case "a":
                    var v196 int8 = -1
                    _ = v196
                case "\n\t":
                    delete(v5, "foo bar")
                    var v197 uint = 201
                    v6 = v6
                    var v198 bool = ((v6 >= 330) && (true && false))
                    _ = v197
                    _ = v198
                case "hello":
                    if ((true || (!true)) || ((484 > v6) || false)) {
                        var v199 float32 = 3.14159
                        _ = v199
                    }
                default:
                    v6 -= -604
                }
                _ = fi194
            }
            var v200 int = fi193
            _ = fi193
            _ = v200
        }
    }
    v2 -= v2
    v201 := v4[v3]
    v202 := GenS25{F0: -984, F1: 34}
    v203 := GenS24{F0: 2, F1: v201}
    v204 := GenS26{F0: ((!(v3 != v3)) && true)}
    switch "hello" {
    case "hello":
        if true {
            var v205 int = 607
            _ = v205
        }
    case "\n\t":
        v7 += 1
    case "foo bar":
        var v206 uint = 958
        _ = v206
    default:
        switch 2147483647 {
        case -2147483648:
            switch "hello" {
            case "foo bar":
                for fi207 := 0; fi207 < 15; fi207++ {
                    var v208 int8 = -114
                    var v209 int64 = v6
                    var v210 uint8 = 127
                    _ = fi207
                    _ = v208
                    _ = v209
                    _ = v210
                }
            case "":
                var v211 int16 = -32768
                _ = v211
            case "\n\t":
                for fi212 := 0; fi212 < 16; fi212++ {
                    v213 := GenS5{F0: (^(fi212)), F1: 2, F2: v3}
                    v214 := GenS24{F0: 0, F1: (true && ((v3 % ((v3) | 1)) < (v3 >> ((255) & 63))))}
                    _ = fi212
                    _ = v213
                    _ = v214
                }
            default:
                v3 -= 2147483648
            }
        case 4294967295:
            v215 := v1[1]
            _ = v215
        }
    }
    v216 := v4[v3]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v201:", v201)
    fmt.Println("v216:", v216)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v201
    _ = v202
    _ = v203
    _ = v204
    _ = v216
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
