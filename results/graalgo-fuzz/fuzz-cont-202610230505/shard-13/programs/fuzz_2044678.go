package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int; F1 int8; F2 uint8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int64; F1 int32 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int16; F1 int8 }
type GenS13 struct { F0 int16; F1 uint16 }
type GenS14 struct { F0 uint8; F1 uint8; F2 string }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 uint16; F1 uint16 }
type GenS17 struct { F0 uint8; F1 uint16 }
type GenS18 struct { F0 bool; F1 uint64; F2 int64 }
type GenS19 struct { F0 string }
type GenS20 struct { F0 uint64; F1 int16 }
type GenS21 struct { F0 bool; F1 int8; F2 uint8 }
type GenS22 struct { F0 int }
type GenS23 struct { F0 uint32 }
type GenS24 struct { F0 string; F1 uint32 }

func main() {
    var v1 int = 42
    var v2 uint32 = 0
    v3 := make(map[uint32]string)
    var v4 uint = 393
    var v5 uint = v4
    if true {
        _ = len(v3)
        v5 = v4
        var v6 uint8 = 185
        _ = v6
    } else {
        fmt.Println(v1)
        if false {
            v1 = (^(v1))
            var v7 uint64 = 1
            _ = v7
        } else {
            var v8 uint16 = 124
            _ = v8
        }
        v2 = v2
    }
    switch "foo bar" {
    case "foo bar":
        for fi9 := 0; fi9 < 5; fi9++ {
            fmt.Println(v1, fi9)
            _ = fi9
        }
    default:
        if (!(true && ((808 != 42) && true))) {
            for fi10 := 0; fi10 < 8; fi10++ {
                v11 := []int{fi10, -615, (v1 * (fi10 % (((v1 % ((v1) | 1))) | 1)))}
                _ = fi10
                _ = v11
            }
            v12 := GenS1{F0: 2, F1: v2, F2: 116}
            _ = v12
        }
    }
    if (!true) {
        if false {
            v2 = v2
        } else {
            if (false && (-9223372036854775808 <= -9223372036854775808)) {
                if (!false) {
                    v13 := GenS2{F0: -2147483648, F1: 0}
                    var v14 uint16 = 42
                    switch 0 {
                    case -991:
                        for fi15 := 0; fi15 < 8; fi15++ {
                            v2 = (v2 | v2)
                            _ = fi15
                        }
                    }
                    _ = v13
                    _ = v14
                } else {
                    for fi16 := 0; fi16 < 1; fi16++ {
                        for fi17 := 0; fi17 < 4; fi17++ {
                            v4 = uint(v2)
                            v18 := []int{v1, (v1 ^ fi16), fi16, (v1 / ((-9223372036854775808) | 1))}
                            _ = fi17
                            _ = v18
                        }
                        var v19 bool = false
                        fmt.Println(v1, v2)
                        _ = fi16
                        _ = v19
                    }
                    var v20 float64 = (1.5 * -2.25)
                    _ = v20
                }
                if true {
                    v21 := []int{(v1 / (((v1 % ((1) | 1))) | 1))}
                    var v22 uint64 = 0
                    var v23 uint16 = 411
                    _ = v21
                    _ = v22
                    _ = v23
                } else {
                    v1 += int(v2)
                }
            }
        }
        v1 = v1
    } else {
        if false {
            if (!(1.5 != 3.14159)) {
                var v24 uint32 = (v2 / (((^(v2))) | 1))
                _ = v24
            }
            if true {
                v2 = (^(v2))
            } else {
                for fi25 := 0; fi25 < 4; fi25++ {
                    var v26 bool = (18446744073709551615 <= 0)
                    var v27 int16 = -1
                    var v28 int64 = 9223372036854775807
                    _ = fi25
                    _ = v26
                    _ = v27
                    _ = v28
                }
                for fi29 := 0; fi29 < 6; fi29++ {
                    if true {
                        if (false && false) {
                            if false {
                                v30 := []int{}
                                v5 += 65535
                                _ = v30
                            } else {
                                var v31 uint16 = 27
                                _ = v31
                            }
                        } else {
                            if (100 <= 0) {
                                var v32 int32 = 42
                                v33 := GenS2{F0: (v32 >> ((v4) & 63)), F1: 913}
                                _ = v32
                                _ = v33
                            } else {
                                var v34 bool = (!(false && false))
                                v4 = (v4 / ((uint(v2)) | 1))
                                var v35 int64 = 42
                                _ = v34
                                _ = v35
                            }
                            var v36 int8 = -114
                            if false {
                                v1 -= v1
                                for fi37 := 0; fi37 < 12; fi37++ {
                                    if false {
                                        v38 := GenS3{F0: 65535, F1: 0}
                                        var v39 uint64 = 18446744073709551615
                                        v1 -= 127
                                        _ = v38
                                        _ = v39
                                    } else {
                                        v36 -= -44
                                        fmt.Println(fi29)
                                    }
                                    if (v5 > (v5 - (v5 * (v4 ^ 255)))) {
                                        var v40 uint64 = 2
                                        _ = v40
                                    }
                                    v36 -= (v36 << (((v4 % ((v4) | 1))) & 63))
                                    _ = fi37
                                }
                            }
                            _ = v36
                        }
                        fmt.Println(v4, v2, v1)
                    } else {
                        fmt.Println(v5, v1)
                        for fi41 := 0; fi41 < 20; fi41++ {
                            if (541 > v4) {
                                v4 = v5
                                v5 = (v4 + 18446744073709551615)
                                var v42 int16 = -76
                                _ = v42
                            } else {
                                v3[v2] = "a"
                                var v43 int16 = -803
                                v5 = 255
                                _ = v43
                            }
                            v44 := GenS4{F0: 1, F1: 0}
                            _ = fi41
                            _ = v44
                        }
                    }
                    for fi45 := 0; fi45 < 12; fi45++ {
                        if true {
                            v5 = uint(v2)
                            var v46 int64 = 1
                            _ = v46
                        }
                        v47 := GenS3{F0: 32767, F1: 42}
                        _ = fi45
                        _ = v47
                    }
                    if true {
                        var v48 float32 = 0.1
                        _ = v48
                    } else {
                        if (!false) {
                            switch "\n\t" {
                            case "\n\t":
                                v49 := []int{32767}
                                _ = v49
                            case "a":
                                v4 += v4
                            case "hello":
                                fmt.Println(v1)
                            default:
                                v50 := GenS4{F0: -32768, F1: 1}
                                _ = v50
                            }
                            switch "a" {
                            case "a":
                                v4 = 2
                            case "foo bar":
                                switch v1 {
                                case 373:
                                    v51 := GenS4{F0: -1, F1: 4294967295}
                                    _ = v51
                                case 127:
                                    if (2 < 18446744073709551615) {
                                        for fi52 := 0; fi52 < 1; fi52++ {
                                            v53 := []int{}
                                            var v54 int16 = -400
                                            var v55 int8 = 127
                                            _ = fi52
                                            _ = v53
                                            _ = v54
                                            _ = v55
                                        }
                                        if (!((fi29 >> (((v5 | uint(v2))) & 63)) <= (v1 | (v1 % (((fi29 << ((18446744073709551615) & 63))) | 1))))) {
                                            v56 := GenS5{F0: 32767}
                                            var v57 uint64 = 1
                                            _ = v56
                                            _ = v57
                                        }
                                        var v58 int32 = 1
                                        _ = v58
                                    }
                                default:
                                    v59 := GenS6{F0: v1, F1: 104, F2: 119}
                                    _ = v59
                                }
                            }
                        } else {
                            var v60 int8 = 100
                            for fi61 := 0; fi61 < 13; fi61++ {
                                v5 -= (^(v4))
                                v5 = (v4 % ((1) | 1))
                                _ = fi61
                            }
                            v1 = (fi29 | 2147483647)
                            _ = v60
                        }
                        delete(v3, (^(v2)))
                    }
                    _ = fi29
                }
            }
        } else {
            fmt.Println(v4, v4)
            if (("foo bar" + "hello") > "") {
                for fi62 := 0; fi62 < 9; fi62++ {
                    if false {
                        for fi63 := 0; fi63 < 6; fi63++ {
                            v64 := []int{(fi63 | fi63), fi63, 0, (-(fi62))}
                            var v65 uint64 = 18446744073709551615
                            _ = fi63
                            _ = v64
                            _ = v65
                        }
                        v5 = v4
                    } else {
                        if true {
                            if (-445 <= 42) {
                                var v66 float32 = float32(-2.25)
                                v4 = (v5 & 0)
                                switch 793 {
                                case 4294967295:
                                    for fi67 := 0; fi67 < 5; fi67++ {
                                        var v68 int = -505
                                        var v69 bool = true
                                        _ = fi67
                                        _ = v68
                                        _ = v69
                                    }
                                case 42:
                                    fmt.Println(v2)
                                }
                                _ = v66
                            } else {
                                if (!false) {
                                    for fi70 := 0; fi70 < 14; fi70++ {
                                        var v71 float32 = 1e10
                                        _ = fi70
                                        _ = v71
                                    }
                                    var v72 int32 = 1
                                    _ = v72
                                }
                                if (v2 != (v2 ^ (^(v2)))) {
                                    fmt.Println(fi62)
                                    for fi73 := 0; fi73 < 7; fi73++ {
                                        v74 := GenS7{F0: -768}
                                        v75 := GenS8{F0: 195, F1: false, F2: v2}
                                        _ = fi73
                                        _ = v74
                                        _ = v75
                                    }
                                    var v76 float64 = float64(0.0)
                                    _ = v76
                                } else {
                                    v4 += 42
                                    v4 -= (v4 % ((18446744073709551615) | 1))
                                    if false {
                                        v77 := GenS6{F0: fi62, F1: -128, F2: 148}
                                        for fi78 := 0; fi78 < 18; fi78++ {
                                            v79 := GenS2{F0: 0, F1: 2147483647}
                                            v4 = v4
                                            delete(v3, 991)
                                            v80 := GenS9{F0: -2147483648}
                                            v2 -= (^(v2))
                                            _ = fi78
                                            _ = v79
                                            _ = v80
                                        }
                                        _ = v77
                                    }
                                }
                                var v81 uint8 = 2
                                _ = v81
                            }
                        } else {
                            var v82 int64 = 9
                            var v83 int64 = (v82 + int64(v1))
                            if true {
                                fmt.Println(v83)
                            }
                            _ = v82
                            _ = v83
                        }
                        for fi84 := 0; fi84 < 17; fi84++ {
                            v3[42] = "a"
                            v4 -= v5
                            _ = fi84
                        }
                    }
                    switch 127 {
                    case 42:
                        var v85 int64 = 0
                        _ = v85
                    case 2147483647:
                        switch ("foo bar" + "\n\t") {
                        case "\n\t":
                            v2 = (v2 - 42)
                        case "a":
                            if (255 > 4294967295) {
                                var v86 uint32 = 422
                                switch -547 {
                                case 32768:
                                    v87 := GenS10{F0: 393, F1: 2147483647}
                                    _ = v87
                                default:
                                    v4 = (v5 >> ((4294967295) & 63))
                                }
                                v1 = fi62
                                _ = v86
                            }
                        case "":
                            var v88 int16 = -32768
                            _ = v88
                        default:
                            for fi89 := 0; fi89 < 4; fi89++ {
                                if (float32(-2.25) != (-2.25 + ((0.0 - 3.14159) - 0.1))) {
                                    v90 := []int{}
                                    switch 128 {
                                    case 128:
                                        if true {
                                            var v91 int8 = 127
                                            _ = v91
                                        } else {
                                            var v92 uint16 = 255
                                            if false {
                                                fmt.Println(v5)
                                                if ((v2 * v2) <= (^(v2))) {
                                                    var v93 uint8 = 2
                                                    var v94 string = "foo bar"
                                                    v5 = v4
                                                    _ = v93
                                                    _ = v94
                                                }
                                            } else {
                                                v95 := GenS11{F0: 318}
                                                if true {
                                                    switch (^(fi62)) {
                                                    case 32767:
                                                        fmt.Println(v5)
                                                    case 42:
                                                        v5 -= (^(v4))
                                                    case 2147483647:
                                                        v5 += 42
                                                    }
                                                }
                                                var v96 int16 = 608
                                                _ = v95
                                                _ = v96
                                            }
                                            _ = v92
                                        }
                                        if len(v90) > 0 {
                                            v90[0] = int(v4)
                                        }
                                        if false {
                                            v5 -= (v4 + v4)
                                            var v97 int8 = 125
                                            fmt.Println(fi89, v2, v97)
                                            _ = v97
                                        } else {
                                            delete(v3, v2)
                                            if (!((v4 / (((v4 ^ (v5 | 42))) | 1)) < v4)) {
                                                switch "" {
                                                case "a":
                                                    var v98 float32 = float32((float64(0.0) + 0.0))
                                                    var v99 uint64 = 553
                                                    if false {
                                                        fmt.Println(v2, v98)
                                                        var v100 uint8 = 255
                                                        fmt.Println(fi89)
                                                        switch ("\n\t" + "\n\t") {
                                                        case "foo bar":
                                                            if (fi62 >= (fi89 % ((fi62) | 1))) {
                                                                if (474 >= 656) {
                                                                    if true {
                                                                        var v101 uint8 = (v100 % ((v100) | 1))
                                                                        v90 = append(v90, (v1 << ((2) & 63)))
                                                                        _ = v101
                                                                    }
                                                                    if true {
                                                                        if false {
                                                                            v5 -= v4
                                                                            _ = len(v90)
                                                                        } else {
                                                                            fmt.Println(v4)
                                                                            if false {
                                                                                if ((v4 % (((v4 ^ 255)) | 1)) < v5) {
                                                                                    if (918 == 1000) {
                                                                                        var v102 string = "\n\t"
                                                                                        switch (("a" + v102) + ("\n\t" + v102)) {
                                                                                        case "hello":
                                                                                            if true {
                                                                                                fmt.Println(fi62, v4, v5)
                                                                                                fmt.Println(fi62, fi89, v5)
                                                                                            } else {
                                                                                                _ = len(v90)
                                                                                                v2 += (v2 >> (((v4 ^ 4294967295)) & 63))
                                                                                                fmt.Println(fi62, v1, v2)
                                                                                            }
                                                                                        case "":
                                                                                            switch v102 {
                                                                                            case "\n\t":
                                                                                                v103 := GenS12{F0: 32767, F1: 100}
                                                                                                _ = v103
                                                                                            case "a":
                                                                                                var v104 uint64 = (^(v99))
                                                                                                _ = v104
                                                                                            }
                                                                                        }
                                                                                        switch (fi89 / ((fi89) | 1)) {
                                                                                        case -1:
                                                                                            if (!true) {
                                                                                                var v105 uint32 = (v2 / ((v2) | 1))
                                                                                                v2 = (v2 & 4294967295)
                                                                                                _ = v105
                                                                                            } else {
                                                                                                if len(v90) >= 2 {
                                                                                                    v90 = v90[1:len(v90)]
                                                                                                }
                                                                                            }
                                                                                        case -9223372036854775808:
                                                                                            delete(v3, 540)
                                                                                        case 990:
                                                                                            _ = len(v3)
                                                                                        }
                                                                                        v2 = (v2 - v2)
                                                                                        _ = v102
                                                                                    } else {
                                                                                        switch fi62 {
                                                                                        case 127:
                                                                                            switch ("a" + "a") {
                                                                                            case "foo bar":
                                                                                                v99 += v99
                                                                                            case "\n\t":
                                                                                                v4 += v5
                                                                                            }
                                                                                        case 562:
                                                                                            switch (fi89 * -9223372036854775808) {
                                                                                            case -1:
                                                                                                var v106 int16 = 619
                                                                                                _ = v106
                                                                                            }
                                                                                        case 979:
                                                                                            var v107 int16 = -32768
                                                                                            _ = v107
                                                                                        }
                                                                                        switch "a" {
                                                                                        case "":
                                                                                            if true {
                                                                                                v1 -= (fi89 ^ (fi89 | fi89))
                                                                                                v4 = v4
                                                                                                v4 = uint(v100)
                                                                                            }
                                                                                        default:
                                                                                            v108 := GenS13{F0: 32767, F1: 65535}
                                                                                            v109 := GenS4{F0: -32768, F1: (^(v5))}
                                                                                            _ = v108
                                                                                            _ = v109
                                                                                        }
                                                                                    }
                                                                                    v100 = v100
                                                                                    _ = len(v3)
                                                                                    switch ("\n\t" + "hello") {
                                                                                    case "\n\t":
                                                                                        var v110 uint8 = v100
                                                                                        _ = v110
                                                                                    }
                                                                                    v111 := GenS14{F0: v100, F1: (v100 ^ uint8(v5)), F2: "\n\t"}
                                                                                    _ = v111
                                                                                } else {
                                                                                    v98 = v98
                                                                                    v100 = (^(v100))
                                                                                    if len(v90) >= 2 {
                                                                                        v90 = v90[1:len(v90)]
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        var v112 uint32 = 2147483647
                                                                        _ = v112
                                                                    }
                                                                    v113 := GenS15{F0: -32768}
                                                                    if false {
                                                                        switch (fi89 * fi89) {
                                                                        case -254:
                                                                            fmt.Println(v98, v1)
                                                                        case -1:
                                                                            var v114 uint32 = v2
                                                                            v100 = (^(v100))
                                                                            if true {
                                                                                var v115 uint = v4
                                                                                if ((188 == 42) && true) {
                                                                                    if false {
                                                                                        v100 += v100
                                                                                        v98 -= 0.0
                                                                                    }
                                                                                } else {
                                                                                    v116 := GenS9{F0: 42}
                                                                                    var v117 int64 = 1
                                                                                    _ = v116
                                                                                    _ = v117
                                                                                }
                                                                                var v118 string = "hello"
                                                                                var v119 uint8 = 2
                                                                                _ = v115
                                                                                _ = v118
                                                                                _ = v119
                                                                            }
                                                                            _ = v114
                                                                        default:
                                                                            var v120 uint8 = 255
                                                                            _ = v120
                                                                        }
                                                                        var v121 int16 = 1000
                                                                        var v122 int = (-(v1))
                                                                        _ = v121
                                                                        _ = v122
                                                                    } else {
                                                                        var v123 string = ("a" + ((("a" + "") + "a") + "hello"))
                                                                        if (false && (0 != 127)) {
                                                                            if false {
                                                                                v124 := GenS11{F0: 1}
                                                                                v98 += v98
                                                                                _ = v124
                                                                            } else {
                                                                                delete(v3, (^(v2)))
                                                                                var v125 uint32 = uint32(v100)
                                                                                _ = v125
                                                                            }
                                                                            fmt.Println(v100, v4)
                                                                            v126 := GenS16{F0: 2, F1: 882}
                                                                            _ = v126
                                                                        }
                                                                        if false {
                                                                            v127 := GenS13{F0: 1, F1: 255}
                                                                            switch v123 {
                                                                            case "\n\t":
                                                                                v128 := GenS17{F0: v100, F1: uint16(v100)}
                                                                                _ = v128
                                                                            case "":
                                                                                v100 += (v100 % ((v100) | 1))
                                                                            case "foo bar":
                                                                                v129 := v3[v2]
                                                                                _ = v129
                                                                            }
                                                                            _ = v127
                                                                        } else {
                                                                            v130 := v3[v2]
                                                                            fmt.Println(v1, v100)
                                                                            var v131 int = (v1 + fi89)
                                                                            v132 := GenS11{F0: 573}
                                                                            _ = v130
                                                                            _ = v131
                                                                            _ = v132
                                                                        }
                                                                        if true {
                                                                            v90 = append(v90, fi62)
                                                                        }
                                                                        var v133 string = v123
                                                                        fmt.Println(v4, v1, v4)
                                                                        _ = v123
                                                                        _ = v133
                                                                    }
                                                                    if true {
                                                                        v5 += (v4 | v5)
                                                                        if (v5 >= 1) {
                                                                            var v134 uint16 = 32767
                                                                            switch ("a" + "foo bar") {
                                                                            case "":
                                                                                fmt.Println(v134, v2, v4)
                                                                            case "a":
                                                                                v135 := v3[v2]
                                                                                _ = v135
                                                                            default:
                                                                                var v136 uint64 = (v99 ^ (^(v99)))
                                                                                _ = v136
                                                                            }
                                                                            _ = v134
                                                                        } else {
                                                                            if ((!true) || false) {
                                                                                if (42 <= 127) {
                                                                                    var v137 uint16 = 32767
                                                                                    var v138 bool = (int8(v137) != int8(v137))
                                                                                    _ = v137
                                                                                    _ = v138
                                                                                } else {
                                                                                    var v139 uint8 = v100
                                                                                    var v140 string = ""
                                                                                    var v141 float32 = 3.14159
                                                                                    _ = v139
                                                                                    _ = v140
                                                                                    _ = v141
                                                                                }
                                                                            } else {
                                                                                v4 = v4
                                                                                if (!((true || (false && ((fi89 + v1) < fi89))) && true)) {
                                                                                    var v142 int8 = 127
                                                                                    switch fi62 {
                                                                                    case 9223372036854775807:
                                                                                        var v143 uint16 = 937
                                                                                        _ = v143
                                                                                    default:
                                                                                        v144 := GenS14{F0: uint8(v99), F1: (v100 % (((^(v100))) | 1)), F2: "foo bar"}
                                                                                        _ = v144
                                                                                    }
                                                                                    v145 := GenS18{F0: false, F1: (^(v99)), F2: 187}
                                                                                    _ = v142
                                                                                    _ = v145
                                                                                } else {
                                                                                    if (((v98 / v98) * -2.25) != v98) {
                                                                                        switch "" {
                                                                                        case "hello":
                                                                                            if false {
                                                                                                switch "" {
                                                                                                case "hello":
                                                                                                    switch -576 {
                                                                                                    case 127:
                                                                                                        v4 += 0
                                                                                                        if (v98 >= (v98 / v98)) {
                                                                                                            v5 -= (v5 % (((v4 ^ v4)) | 1))
                                                                                                            v4 = v5
                                                                                                        }
                                                                                                    case -9223372036854775808:
                                                                                                        fmt.Println(v100, v2)
                                                                                                    case 42:
                                                                                                        if false {
                                                                                                            v98 += (v98 - 0.0)
                                                                                                        } else {
                                                                                                            v146 := GenS19{F0: "hello"}
                                                                                                            v4 = (v5 ^ (v5 + 1))
                                                                                                            _ = v146
                                                                                                        }
                                                                                                        fmt.Println(fi62, v1)
                                                                                                    default:
                                                                                                        if ((fi62 * fi89) >= v1) {
                                                                                                            switch "foo bar" {
                                                                                                            case "hello":
                                                                                                                v147 := GenS8{F0: (v99 % (((^(v99))) | 1)), F1: true, F2: 2147483647}
                                                                                                                _ = v147
                                                                                                            }
                                                                                                            if len(v90) > 0 {
                                                                                                                v90[0] = 643
                                                                                                            }
                                                                                                            v5 -= v4
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                _ = len(v3)
                                                                                                v148 := GenS2{F0: 600, F1: -542}
                                                                                                _ = v148
                                                                                            }
                                                                                        }
                                                                                        switch "\n\t" {
                                                                                        case "\n\t":
                                                                                            fmt.Println(v2, v2)
                                                                                        default:
                                                                                            switch ((("hello" + ("a" + (("" + (("hello" + "\n\t") + ("\n\t" + ""))) + (((("hello" + "\n\t") + "") + (("\n\t" + "\n\t") + "")) + ((("a" + "hello") + "") + (("\n\t" + ("\n\t" + ("" + "foo bar"))) + ((("\n\t" + "\n\t") + ("a" + ("hello" + ("" + ("hello" + "foo bar"))))) + "hello"))))))) + "") + "hello") {
                                                                                            case "\n\t":
                                                                                                v4 -= (v4 % ((v5) | 1))
                                                                                            }
                                                                                        }
                                                                                        switch "hello" {
                                                                                        case "a":
                                                                                            v1 -= fi62
                                                                                        default:
                                                                                            v90 = append(v90, v1)
                                                                                            v149 := GenS13{F0: 806, F1: 0}
                                                                                            if false {
                                                                                                var v150 int32 = 1
                                                                                                _ = v150
                                                                                            } else {
                                                                                                var v151 int16 = -32768
                                                                                                v1 = (fi62 % ((fi89) | 1))
                                                                                                v152 := GenS18{F0: (!(!true)), F1: (v99 + uint64(v100)), F2: -9223372036854775808}
                                                                                                fmt.Println(v5, v2)
                                                                                                switch int(v2) {
                                                                                                case 743:
                                                                                                    v153 := GenS20{F0: v99, F1: (v151 & int16(v5))}
                                                                                                    _ = v153
                                                                                                }
                                                                                                v154 := GenS18{F0: false, F1: (v99 << ((1) & 63)), F2: 1}
                                                                                                var v155 int = fi62
                                                                                                if false {
                                                                                                    switch "hello" {
                                                                                                    case "hello":
                                                                                                        v156 := GenS1{F0: 42, F1: v2, F2: v100}
                                                                                                        _ = v156
                                                                                                    case "a":
                                                                                                        v98 -= 0.1
                                                                                                    }
                                                                                                    fmt.Println(v99, fi89)
                                                                                                    var v157 bool = false
                                                                                                    _ = v157
                                                                                                } else {
                                                                                                    var v158 int8 = 1
                                                                                                    v98 = (float32(0.0) / ((float32((float64(((float32(0.0) / v98) - 1e10)) / 1.5)) / float32(1e10)) / 1.0))
                                                                                                    var v159 bool = (255 != 42)
                                                                                                    _ = v158
                                                                                                    _ = v159
                                                                                                }
                                                                                                var v160 int64 = 42
                                                                                                v4 = v5
                                                                                                var v161 uint64 = 416
                                                                                                fmt.Println(v100, v2)
                                                                                                if true {
                                                                                                    if len(v90) >= 2 {
                                                                                                        v90 = v90[1:len(v90)]
                                                                                                    }
                                                                                                    var v162 uint16 = 831
                                                                                                    _ = v162
                                                                                                }
                                                                                                _ = v151
                                                                                                _ = v152
                                                                                                _ = v154
                                                                                                _ = v155
                                                                                                _ = v160
                                                                                                _ = v161
                                                                                            }
                                                                                            _ = v149
                                                                                        }
                                                                                    } else {
                                                                                        v163 := GenS6{F0: fi62, F1: 127, F2: v100}
                                                                                        v164 := GenS2{F0: 2147483647, F1: 42}
                                                                                        _ = v163
                                                                                        _ = v164
                                                                                    }
                                                                                }
                                                                            }
                                                                            fmt.Println(fi62, fi89)
                                                                            v165 := GenS21{F0: true, F1: 1, F2: (^(v100))}
                                                                            _ = v165
                                                                        }
                                                                        v166 := GenS1{F0: 881, F1: (v2 ^ (v2 / (((v2 >> ((v4) & 63))) | 1))), F2: (^(v100))}
                                                                        _ = v166
                                                                    }
                                                                    _ = v113
                                                                }
                                                                var v167 uint64 = (^(v99))
                                                                fmt.Println(v98)
                                                                _ = v167
                                                            }
                                                        case "hello":
                                                            if (false || false) {
                                                                _ = len(v90)
                                                                delete(v3, (v2 ^ (v2 ^ 639)))
                                                            } else {
                                                                v99 -= (^(v99))
                                                            }
                                                        case "a":
                                                            v98 = ((v98 * 1.5) + ((float32((1.5 + 1.5)) + -2.25) + float32(3.14159)))
                                                        }
                                                        _ = v100
                                                    } else {
                                                        v168 := v3[v2]
                                                        _ = v168
                                                    }
                                                    _ = v98
                                                    _ = v99
                                                case "hello":
                                                    v169 := GenS22{F0: (-(fi89))}
                                                    _ = v169
                                                case "foo bar":
                                                    v170 := GenS23{F0: (v2 | (^(v2)))}
                                                    _ = v170
                                                default:
                                                    var v171 uint = 18446744073709551615
                                                    _ = v171
                                                }
                                                fmt.Println(v4, fi62)
                                                v5 -= 255
                                                delete(v3, 0)
                                            } else {
                                                fmt.Println(v4, v4, fi89)
                                            }
                                            var v172 uint8 = 0
                                            var v173 uint8 = (v172 - 42)
                                            fmt.Println(v4, fi89, v4)
                                            fmt.Println(v1)
                                            _ = v172
                                            _ = v173
                                        }
                                        v174 := GenS16{F0: 255, F1: 255}
                                        v175 := GenS6{F0: (-(v1)), F1: 127, F2: uint8(v2)}
                                        _ = v174
                                        _ = v175
                                    case 32767:
                                        if true {
                                            v176 := v3[v2]
                                            fmt.Println(v1, v4)
                                            if false {
                                                if true {
                                                    fmt.Println(v2)
                                                } else {
                                                    v3[850] = (v176 + "")
                                                    if (!(128 > 128)) {
                                                        fmt.Println(fi62)
                                                    } else {
                                                        if ((!(-1 > -983)) && (("\n\t" + "a") >= v176)) {
                                                            fmt.Println(fi62)
                                                            var v177 string = v176
                                                            fmt.Println(fi89, v4, fi89)
                                                            _ = v177
                                                        } else {
                                                            v178 := GenS11{F0: -9223372036854775808}
                                                            var v179 int8 = 100
                                                            switch ((v176 + ("\n\t" + v176)) + v176) {
                                                            case "foo bar":
                                                                if false {
                                                                    var v180 int8 = v179
                                                                    _ = v180
                                                                }
                                                            case "\n\t":
                                                                delete(v3, (v2 % ((61) | 1)))
                                                            }
                                                            _ = v178
                                                            _ = v179
                                                        }
                                                    }
                                                }
                                                if true {
                                                    var v181 int64 = 42
                                                    switch 4294967295 {
                                                    case 128:
                                                        if true {
                                                            var v182 int64 = (v181 << ((2) & 63))
                                                            _ = v182
                                                        }
                                                    case 61:
                                                        v1 += (fi62 | fi89)
                                                    case 2147483647:
                                                        v4 -= v4
                                                    default:
                                                        switch (^(fi89)) {
                                                        case 128:
                                                            var v183 bool = (((!(false || (18446744073709551615 >= (v5 >> (((v5 ^ v4)) & 63))))) || false) && true)
                                                            _ = v183
                                                        default:
                                                            v181 = (v181 << ((v5) & 63))
                                                        }
                                                    }
                                                    if (!false) {
                                                        if len(v90) >= 2 {
                                                            v90 = v90[1:len(v90)]
                                                        }
                                                    } else {
                                                        fmt.Println(v2)
                                                        _ = len(v3)
                                                        v181 = 1
                                                    }
                                                    _ = v181
                                                } else {
                                                    var v184 int = v1
                                                    v5 += (v5 >> (((v5 + v5)) & 63))
                                                    var v185 float32 = 1.5
                                                    var v186 bool = false
                                                    _ = v184
                                                    _ = v185
                                                    _ = v186
                                                }
                                                v187 := v3[(^(v2))]
                                                _ = v187
                                            }
                                            _ = v176
                                        } else {
                                            fmt.Println(v2, v4)
                                        }
                                    case 480:
                                        _ = len(v90)
                                    }
                                    var v188 uint32 = 2147483647
                                    _ = v188
                                    _ = v90
                                } else {
                                    var v189 int8 = 66
                                    var v190 int = 742
                                    _ = v189
                                    _ = v190
                                }
                                v1 += fi89
                                var v191 int32 = -765
                                var v192 float64 = (float64(float32(1e10)) + (1.5 * (3.14159 * 1e10)))
                                v192 = (float64(1e10) + 0.1)
                                _ = fi89
                                _ = v191
                                _ = v192
                            }
                        }
                    case -270:
                        switch fi62 {
                        case 4294967295:
                            fmt.Println(fi62, v5, v4)
                        case -1:
                            for fi193 := 0; fi193 < 16; fi193++ {
                                fmt.Println(v1)
                                var v194 int16 = 1000
                                _ = fi193
                                _ = v194
                            }
                        default:
                            v195 := []int{0, v1, (v1 / ((42) | 1))}
                            _ = v195
                        }
                    default:
                        var v196 uint32 = 42
                        _ = v196
                    }
                    var v197 bool = ((!false) || false)
                    _ = fi62
                    _ = v197
                }
                if (true || ((!true) && false)) {
                    v5 += v5
                }
                fmt.Println(v2, v4, v2)
            }
            var v198 float64 = ((-2.25 / float64((1e10 * (float32(-2.25) + -2.25)))) + 0.1)
            _ = v198
        }
    }
    var v199 uint8 = 128
    var v200 float32 = -2.25
    for fi201 := 0; fi201 < 19; fi201++ {
        v4 -= v5
        fmt.Println(v5, v1, fi201)
        v202 := []int{}
        _ = fi201
        _ = v202
    }
    v5 += (^(v4))
    fmt.Println(v4, v4)
    if false {
        for fi203 := 0; fi203 < 1; fi203++ {
            v204 := GenS21{F0: ((uint64(v5) < 18446744073709551615) || true), F1: -1, F2: 255}
            _ = fi203
            _ = v204
        }
    }
    for fi205 := 0; fi205 < 17; fi205++ {
        for fi206 := 0; fi206 < 8; fi206++ {
            fmt.Println(fi206, v2, v1)
            _ = fi206
        }
        switch ("a" + "hello") {
        case "\n\t":
            if (-911 >= 9223372036854775807) {
                if true {
                    switch "foo bar" {
                    case "\n\t":
                        var v207 uint = (v4 & (v5 / ((v5) | 1)))
                        _ = v207
                    case "":
                        v208 := GenS24{F0: "", F1: (v2 & v2)}
                        _ = v208
                    }
                    var v209 string = "hello"
                    v4 = v5
                    _ = v209
                }
                v200 -= 1.5
            }
        case "":
            var v210 int = fi205
            _ = v210
        }
        _ = fi205
    }
    fmt.Println("v1:", v1)
    fmt.Println("v199:", v199)
    fmt.Println("v2:", v2)
    fmt.Println("v200:", v200)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v199
    _ = v2
    _ = v200
    _ = v3
    _ = v4
    _ = v5
}
