package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int64; F1 int8; F2 int16 }
type GenS7 struct { F0 uint32; F1 uint64 }
type GenS8 struct { F0 uint32; F1 int64; F2 int16 }
type GenS9 struct { F0 uint; F1 int }
type GenS10 struct { F0 int64; F1 bool; F2 string }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 int64; F1 uint8 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 uint8; F1 int16 }
type GenS16 struct { F0 int16; F1 uint64 }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 uint16; F1 int; F2 int8 }
type GenS19 struct { F0 int64; F1 bool; F2 bool }
type GenS20 struct { F0 int8; F1 uint }
func genFn0(p0 int32, p1 int64) int8 {
    var v6 uint16 = 65535
    v7 := GenS1{F0: v6, F1: 296}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -128
}

func main() {
    var v1 uint16 = 944
    var v2 float64 = float64(1e10)
    var v3 int16 = 0
    var v4 uint64 = 498
    v5 := []uint32{1, 2, 1}
    var v8 string = ""
    switch 336 {
    case 4294967295:
        v1 -= (v1 % ((v1) | 1))
    case 32767:
        v9 := GenS1{F0: 255, F1: 18446744073709551615}
        _ = v9
    case 0:
        var v10 uint16 = 32767
        _ = v10
    default:
        switch "\n\t" {
        case "":
            var v11 int = -9223372036854775808
            _ = v11
        case "a":
            for fi12 := 0; fi12 < 14; fi12++ {
                if false {
                    v13 := GenS1{F0: v1, F1: uint(v1)}
                    _ = v13
                } else {
                    if len(v5) > 0 {
                        v5[0] = 508
                    }
                    v14 := GenS1{F0: (v1 & (v1 << ((773) & 63))), F1: 1}
                    v15 := GenS2{F0: (fi12 * fi12)}
                    _ = v14
                    _ = v15
                }
                _ = fi12
            }
        case "hello":
            var v16 int = -641
            _ = v16
        }
    }
    var v17 string = ""
    v18 := GenS1{F0: (v1 | v1), F1: 983}
    var v19 float32 = 0.0
    if false {
        v2 = (((v2 * v2) / float64(v19)) / ((v2 * (v2 + v2)) - v2))
        if true {
            if false {
                if true {
                    if true {
                        switch "" {
                        case "a":
                            for fi20 := 0; fi20 < 18; fi20++ {
                                var v21 uint64 = v4
                                if true {
                                    v22 := GenS1{F0: (^(v1)), F1: 65535}
                                    var v23 uint32 = 2147483647
                                    _ = v22
                                    _ = v23
                                } else {
                                    fmt.Println(v3)
                                    for fi24 := 0; fi24 < 16; fi24++ {
                                        if (255 <= 17) {
                                            v19 = 0.0
                                        } else {
                                            var v25 float64 = 1.5
                                            _ = v25
                                        }
                                        if (4294967295 == 152) {
                                            v26 := GenS3{F0: 0, F1: -215, F2: uint64(v1)}
                                            switch (v8 + v8) {
                                            case "":
                                                fmt.Println(v17, v8, v4)
                                            case "\n\t":
                                                var v27 uint8 = 252
                                                _ = v27
                                            case "a":
                                                if true {
                                                    _ = len(v5)
                                                } else {
                                                    switch v8 {
                                                    case "hello":
                                                        v21 = v21
                                                    }
                                                    fmt.Println(v1, v1, v1)
                                                }
                                            }
                                            if (0 < 127) {
                                                _ = len(v5)
                                            } else {
                                                v4 += 42
                                                var v28 int32 = -1
                                                _ = v28
                                            }
                                            v2 += float64(v19)
                                            v3 += v3
                                            fmt.Println(fi24)
                                            if false {
                                                var v29 string = "a"
                                                fmt.Println(v21, fi24)
                                                _ = v29
                                            } else {
                                                var v30 int16 = (-(v3))
                                                var v31 bool = true
                                                _ = v30
                                                _ = v31
                                            }
                                            _ = v26
                                        } else {
                                            v32 := GenS2{F0: (fi24 * fi24)}
                                            _ = v32
                                        }
                                        _ = fi24
                                    }
                                }
                                _ = fi20
                                _ = v21
                            }
                        case "":
                            fmt.Println(v8, v17)
                        case "\n\t":
                            v19 = v19
                        }
                    } else {
                        v5 = append(v5, 2147483647)
                        v2 -= v2
                    }
                }
                v3 = v3
            }
        } else {
            if (v19 != (v19 + ((float32((float64(v19) - ((0.1 / (v2 / ((3.14159 / (3.14159 / (v2 * v2))) - v2))) + (float64(0.1) + (float64((float32(float64(float32(v2))) + 0.0)) / 1.0))))) + v19) * v19))) {
                v8 = v17
                switch v17 {
                case "a":
                    if ((true && false) && (0 < 255)) {
                        v8 = "hello"
                        var v33 uint32 = 421
                        fmt.Println(v8, v17, v33)
                        _ = v33
                    }
                case "\n\t":
                    if true {
                        v5 = append(v5, 500)
                        var v34 string = (v17 + "")
                        _ = v34
                    } else {
                        v3 -= v3
                        if true {
                            v17 = ("a" + (v8 + v8))
                            fmt.Println(v4)
                            var v35 uint32 = 42
                            _ = v35
                        } else {
                            v19 -= float32(0.1)
                        }
                        if true {
                            for fi36 := 0; fi36 < 20; fi36++ {
                                var v37 uint = 1
                                if (!false) {
                                    v38 := GenS4{F0: 0, F1: int32(v3), F2: 255}
                                    if (true || (-2147483648 > 0)) {
                                        v39 := GenS4{F0: 69, F1: 584, F2: 255}
                                        _ = v39
                                    } else {
                                        _ = len(v5)
                                        if len(v5) > 0 {
                                            v5[0] = 1
                                        }
                                        switch v17 {
                                        case "":
                                            v19 = v19
                                        case "hello":
                                            fmt.Println(v3, fi36)
                                        case "\n\t":
                                            v40 := GenS5{F0: 0}
                                            _ = v40
                                        default:
                                            if (!(42 >= -128)) {
                                                v41 := GenS2{F0: 32768}
                                                _ = len(v5)
                                                v2 -= 1.5
                                                _ = v41
                                            } else {
                                                if false {
                                                    fmt.Println(v2, v1, v1)
                                                }
                                                v2 = (0.1 + 1.5)
                                            }
                                        }
                                    }
                                    _ = v38
                                } else {
                                    for fi42 := 0; fi42 < 5; fi42++ {
                                        if true {
                                            v43 := GenS4{F0: 255, F1: 0, F2: 0}
                                            _ = v43
                                        } else {
                                            v44 := GenS6{F0: int64(fi36), F1: 38, F2: v3}
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            v17 = (v17 + v17)
                                            v2 = ((1.5 * float64(0.1)) / 1.0)
                                            var v45 uint = v37
                                            var v46 float64 = v2
                                            var v47 int16 = (v3 ^ (v3 | (v3 ^ v3)))
                                            _ = v44
                                            _ = v45
                                            _ = v46
                                            _ = v47
                                        }
                                        if false {
                                            if (18446744073709551615 >= (v37 ^ v37)) {
                                                var v48 uint8 = 255
                                                v19 = v19
                                                var v49 int32 = -1
                                                if (!(true || (((!((v37 * v37) <= 866)) || false) && (!false)))) {
                                                    v1 -= (^(v1))
                                                    fmt.Println(v49, fi36, v48)
                                                    switch (v17 + "foo bar") {
                                                    case "a":
                                                        fmt.Println(v48, v4, v37)
                                                    case "\n\t":
                                                        if (408 <= (v37 % ((0) | 1))) {
                                                            v50 := GenS2{F0: (fi42 % ((4294967295) | 1))}
                                                            _ = len(v5)
                                                            if true {
                                                                var v51 string = (v8 + "foo bar")
                                                                v1 -= (v1 << ((1) & 63))
                                                                v1 = (v1 * (v1 / (((v1 << ((v37) & 63))) | 1)))
                                                                _ = v51
                                                            }
                                                            _ = v50
                                                        } else {
                                                            v37 = 65535
                                                            var v52 uint = 2
                                                            _ = v52
                                                        }
                                                    }
                                                }
                                                var v53 int8 = 100
                                                v54 := GenS6{F0: 42, F1: v53, F2: v3}
                                                fmt.Println(fi42, v8, v3)
                                                _ = v48
                                                _ = v49
                                                _ = v53
                                                _ = v54
                                            }
                                            if (((1 == (v4 / ((18446744073709551615) | 1))) || false) || (1 == 2147483647)) {
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                var v55 float64 = v2
                                                var v56 int64 = 0
                                                _ = v55
                                                _ = v56
                                            }
                                            switch ((((v8 + v17) + v17) + v17) + v17) {
                                            case "hello":
                                                _ = len(v5)
                                            case "a":
                                                var v57 string = "hello"
                                                _ = v57
                                            case "":
                                                v58 := GenS5{F0: 127}
                                                _ = v58
                                            }
                                            v3 += 886
                                        } else {
                                            fmt.Println(fi36)
                                            if (true && true) {
                                                v59 := GenS7{F0: 2, F1: v4}
                                                _ = v59
                                            }
                                            fmt.Println(fi36)
                                        }
                                        v60 := GenS8{F0: 0, F1: -808, F2: 32767}
                                        _ = fi42
                                        _ = v60
                                    }
                                }
                                _ = fi36
                                _ = v37
                            }
                        }
                    }
                }
                switch v17 {
                case "":
                    if true {
                        if (((0 < 65535) || (!(true && false))) || true) {
                            v17 = v8
                        }
                    } else {
                        var v61 uint64 = (v4 & (v4 % ((v4) | 1)))
                        _ = v61
                    }
                case "foo bar":
                    v62 := GenS6{F0: 9223372036854775807, F1: 0, F2: -32768}
                    _ = v62
                case "hello":
                    var v63 uint8 = 127
                    _ = v63
                }
            }
            switch ((v8 + "foo bar") + "foo bar") {
            case "a":
                v64 := GenS8{F0: 2, F1: -839, F2: (v3 | v3)}
                _ = v64
            default:
                var v65 uint16 = v1
                _ = v65
            }
            if false {
                var v66 int32 = 1
                fmt.Println(v4, v2)
                switch (v8 + (v8 + "hello")) {
                case "a":
                    fmt.Println(v66, v4, v19)
                case "\n\t":
                    var v67 float32 = v19
                    _ = v67
                case "foo bar":
                    for fi68 := 0; fi68 < 18; fi68++ {
                        v17 = ("a" + v8)
                        _ = fi68
                    }
                default:
                    switch 128 {
                    case -2147483648:
                        var v69 float64 = ((float64((0.0 - 0.0)) + (0.1 * (v2 * (v2 + 1e10)))) * v2)
                        _ = v69
                    default:
                        v5 = append(v5, 2147483648)
                    }
                }
                _ = v66
            } else {
                var v70 bool = (1 < -1)
                _ = v70
            }
        }
    } else {
        if (!(852 < 872)) {
            if false {
                for fi71 := 0; fi71 < 16; fi71++ {
                    var v72 uint = uint(v1)
                    _ = fi71
                    _ = v72
                }
                var v73 int8 = 40
                switch ((("a" + "a") + "\n\t") + ((v17 + "foo bar") + ((v17 + "foo bar") + v8))) {
                case "foo bar":
                    for fi74 := 0; fi74 < 12; fi74++ {
                        for fi75 := 0; fi75 < 9; fi75++ {
                            if len(v5) > 0 {
                                v5[0] = 2147483647
                            }
                            var v76 uint = 490
                            if true {
                                v77 := GenS9{F0: (v76 & (v76 / (((v76 / ((v76) | 1))) | 1))), F1: (fi74 % ((297) | 1))}
                                v78 := GenS2{F0: (fi75 & fi75)}
                                var v79 uint = (v76 >> (((^(v76))) & 63))
                                _ = v77
                                _ = v78
                                _ = v79
                            } else {
                                v80 := GenS2{F0: (fi74 >> (((v76 % (((v76 - v76)) | 1))) & 63))}
                                var v81 int = (fi75 + 42)
                                _ = v80
                                _ = v81
                            }
                            var v82 bool = true
                            _ = fi75
                            _ = v76
                            _ = v82
                        }
                        _ = fi74
                    }
                case "":
                    if ((v3 + v3) == (v3 + v3)) {
                        switch (v17 + (v8 + v8)) {
                        case "hello":
                            v3 = v3
                        case "":
                            v2 += v2
                        case "a":
                            v83 := GenS10{F0: 42, F1: (v1 > (v1 / (((v1 % ((0) | 1))) | 1))), F2: ((v17 + "a") + (("" + (v8 + v8)) + v17))}
                            _ = v83
                        default:
                            if (v1 > v1) {
                                switch 127 {
                                case -576:
                                    v2 = (v2 / v2)
                                }
                                var v84 int = -794
                                switch v84 {
                                case 2147483647:
                                    var v85 int16 = (v3 >> ((4294967295) & 63))
                                    _ = v85
                                case 128:
                                    switch "foo bar" {
                                    case "foo bar":
                                        var v86 int64 = 42
                                        _ = v86
                                    case "hello":
                                        switch "hello" {
                                        case "\n\t":
                                            switch 2147483647 {
                                            case 128:
                                                var v87 int32 = 182
                                                _ = v87
                                            case 890:
                                                v8 = v17
                                            }
                                        case "a":
                                            v88 := GenS2{F0: int(v4)}
                                            _ = v88
                                        case "foo bar":
                                            for fi89 := 0; fi89 < 2; fi89++ {
                                                v90 := GenS11{F0: 1}
                                                fmt.Println(v73)
                                                _ = fi89
                                                _ = v90
                                            }
                                        default:
                                            var v91 float32 = v19
                                            _ = v91
                                        }
                                    case "a":
                                        for fi92 := 0; fi92 < 13; fi92++ {
                                            var v93 uint8 = 127
                                            var v94 int8 = (v73 + (v73 / ((-1) | 1)))
                                            v17 = ("foo bar" + v17)
                                            _ = fi92
                                            _ = v93
                                            _ = v94
                                        }
                                    default:
                                        for fi95 := 0; fi95 < 5; fi95++ {
                                            switch v8 {
                                            case "\n\t":
                                                v96 := GenS10{F0: -255, F1: ((v4 % (((v4 ^ v4)) | 1)) <= (v4 % ((v4) | 1))), F2: (v17 + (((("hello" + v8) + v8) + "a") + "hello"))}
                                                _ = v96
                                            }
                                            var v97 float64 = -2.25
                                            _ = fi95
                                            _ = v97
                                        }
                                    }
                                case 127:
                                    var v98 int32 = 0
                                    _ = v98
                                default:
                                    if (4294967295 != 4294967295) {
                                        for fi99 := 0; fi99 < 3; fi99++ {
                                            var v100 uint32 = 4294967295
                                            switch ("hello" + (((v8 + v8) + v17) + "")) {
                                            case "":
                                                for fi101 := 0; fi101 < 6; fi101++ {
                                                    var v102 bool = true
                                                    var v103 uint = 1
                                                    _ = fi101
                                                    _ = v102
                                                    _ = v103
                                                }
                                            case "a":
                                                if false {
                                                    v1 = 255
                                                }
                                            case "hello":
                                                switch v84 {
                                                case -9223372036854775808:
                                                    for fi104 := 0; fi104 < 13; fi104++ {
                                                        _ = len(v5)
                                                        var v105 int32 = -2147483648
                                                        _ = fi104
                                                        _ = v105
                                                    }
                                                case -1:
                                                    _ = len(v5)
                                                case 2147483647:
                                                    v1 = v1
                                                }
                                            default:
                                                if (v17 != (v8 + ((v8 + v17) + "a"))) {
                                                    var v106 string = v17
                                                    _ = v106
                                                }
                                            }
                                            _ = fi99
                                            _ = v100
                                        }
                                        switch v84 {
                                        case -576:
                                            v1 = (v1 - v1)
                                        case 1:
                                            var v107 uint32 = 42
                                            _ = v107
                                        case 127:
                                            for fi108 := 0; fi108 < 9; fi108++ {
                                                var v109 bool = (false && false)
                                                if (v109 && true) {
                                                    if v109 {
                                                        var v110 int64 = 415
                                                        _ = v110
                                                    }
                                                    switch fi108 {
                                                    case 42:
                                                        v111 := GenS1{F0: v1, F1: 65535}
                                                        _ = v111
                                                    default:
                                                        for fi112 := 0; fi112 < 12; fi112++ {
                                                            var v113 uint64 = v4
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                            _ = fi112
                                                            _ = v113
                                                        }
                                                    }
                                                }
                                                _ = fi108
                                                _ = v109
                                            }
                                        }
                                        var v114 int = -689
                                        _ = v114
                                    }
                                }
                                _ = v84
                            } else {
                                switch -9223372036854775808 {
                                case 862:
                                    switch (v17 + v8) {
                                    case "\n\t":
                                        var v115 int64 = -9223372036854775808
                                        _ = v115
                                    default:
                                        if len(v5) > 0 {
                                            v5[0] = 2147483648
                                        }
                                    }
                                case 21:
                                    for fi116 := 0; fi116 < 7; fi116++ {
                                        if (4 != (v73 ^ (v73 - v73))) {
                                            v117 := GenS6{F0: 42, F1: 0, F2: 625}
                                            _ = v117
                                        } else {
                                            for fi118 := 0; fi118 < 15; fi118++ {
                                                var v119 float32 = v19
                                                _ = fi118
                                                _ = v119
                                            }
                                        }
                                        v120 := GenS12{F0: v3}
                                        if len(v5) > 0 {
                                            v5[0] = 4294967295
                                        }
                                        _ = fi116
                                        _ = v120
                                    }
                                default:
                                    var v121 int = 32767
                                    _ = v121
                                }
                                switch 0 {
                                case 42:
                                    var v122 uint32 = 1
                                    _ = v122
                                case 9223372036854775807:
                                    var v123 int32 = 0
                                    _ = v123
                                }
                            }
                        }
                        for fi124 := 0; fi124 < 13; fi124++ {
                            if (!false) {
                                switch v8 {
                                case "foo bar":
                                    var v125 int32 = -326
                                    _ = v125
                                case "\n\t":
                                    v73 += v73
                                case "":
                                    v126 := GenS13{F0: -1, F1: 0}
                                    _ = v126
                                default:
                                    v5 = append(v5, 1)
                                }
                            } else {
                                if (!(false || (!(float32((v2 * v2)) >= (v19 * v19))))) {
                                    var v127 bool = false
                                    for fi128 := 0; fi128 < 7; fi128++ {
                                        v17 = "hello"
                                        _ = len(v5)
                                        _ = fi128
                                    }
                                    _ = v127
                                } else {
                                    _ = len(v5)
                                    switch v17 {
                                    case "foo bar":
                                        if false {
                                            v8 = v17
                                            var v129 int32 = -441
                                            _ = v129
                                        }
                                    default:
                                        v73 = (v73 & v73)
                                    }
                                    var v130 float32 = (v19 / float32(v2))
                                    _ = v130
                                }
                            }
                            _ = fi124
                        }
                    }
                default:
                    var v131 int32 = 574
                    _ = v131
                }
                _ = v73
            }
            var v132 uint64 = 883
            v1 -= 255
            _ = v132
        }
        v19 = v19
    }
    switch 4294967295 {
    case 127:
        var v133 float64 = (v2 - v2)
        _ = v133
    case -2147483648:
        var v134 int32 = -2147483648
        _ = v134
    }
    for fi135 := 0; fi135 < 10; fi135++ {
        v1 = v1
        _ = fi135
    }
    v136 := GenS2{F0: -2147483648}
    v137 := GenS12{F0: (v3 + (v3 ^ (v3 << ((42) & 63))))}
    if true {
        switch 1 {
        case 32767:
            var v138 uint8 = 166
            _ = v138
        case 128:
            if false {
                v139 := GenS14{F0: 2147483648}
                var v140 uint16 = v1
                var v141 uint = 42
                _ = v139
                _ = v140
                _ = v141
            } else {
                for fi142 := 0; fi142 < 18; fi142++ {
                    if true {
                        v17 = (v8 + (v17 + "foo bar"))
                        v143 := GenS15{F0: 255, F1: (v3 / ((v3) | 1))}
                        _ = v143
                    }
                    _ = fi142
                }
                if true {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = len(v5)
                    if len(v5) > 0 {
                        v5[0] = 0
                    }
                } else {
                    v144 := GenS16{F0: v3, F1: (v4 + v4)}
                    v145 := GenS5{F0: 2}
                    for fi146 := 0; fi146 < 17; fi146++ {
                        switch "foo bar" {
                        case "hello":
                            v147 := GenS17{F0: v1}
                            _ = v147
                        case "\n\t":
                            for fi148 := 0; fi148 < 3; fi148++ {
                                switch -1 {
                                case 916:
                                    switch (("hello" + (v17 + v8)) + ((v8 + "hello") + v17)) {
                                    case "foo bar":
                                        var v149 int16 = v3
                                        _ = v149
                                    case "a":
                                        v150 := GenS12{F0: -455}
                                        _ = v150
                                    case "\n\t":
                                        switch fi148 {
                                        case 0:
                                            v2 = v2
                                        case 127:
                                            switch fi146 {
                                            case 4294967295:
                                                if (true && ((100 != 86) || true)) {
                                                    v5 = append(v5, 522)
                                                    var v151 string = ((("\n\t" + (((v8 + ("a" + "a")) + (v8 + v8)) + v17)) + ((((v8 + v8) + ("foo bar" + v17)) + "\n\t") + (v8 + ""))) + v8)
                                                    _ = v151
                                                } else {
                                                    fmt.Println(v17, v4, fi146)
                                                }
                                            case 361:
                                                v152 := GenS5{F0: 42}
                                                if (255 == 51) {
                                                    if (true || true) {
                                                        if true {
                                                            var v153 bool = false
                                                            switch ((((("\n\t" + v17) + v17) + "") + v17) + v8) {
                                                            case "\n\t":
                                                                switch v17 {
                                                                case "":
                                                                    v2 = v2
                                                                case "hello":
                                                                    if ((v3 + 42) > (v3 << ((42) & 63))) {
                                                                        v19 += v19
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                    } else {
                                                                        if (!v153) {
                                                                            v3 = -263
                                                                            switch (fi146 << ((42) & 63)) {
                                                                            case 127:
                                                                                v154 := GenS18{F0: v1, F1: fi146, F2: 52}
                                                                                _ = v154
                                                                            }
                                                                            v1 -= 65535
                                                                            switch ((("hello" + v17) + ((v17 + (v17 + v8)) + v8)) + (v17 + (v17 + (v17 + v8)))) {
                                                                            case "":
                                                                                var v155 uint64 = v4
                                                                                _ = v155
                                                                            }
                                                                            if (1 < -851) {
                                                                                fmt.Println(v153, fi146)
                                                                                fmt.Println(v3)
                                                                                if (!(!((v3 & -964) >= 1))) {
                                                                                    v3 += (-(v3))
                                                                                    fmt.Println(v3)
                                                                                } else {
                                                                                    _ = len(v5)
                                                                                    var v156 bool = (48 == 42)
                                                                                    if len(v5) >= 2 {
                                                                                        v5 = v5[1:len(v5)]
                                                                                    }
                                                                                    _ = v156
                                                                                }
                                                                            } else {
                                                                                v1 = (v1 * v1)
                                                                                v2 -= v2
                                                                            }
                                                                        }
                                                                        v8 = ("a" + "foo bar")
                                                                        fmt.Println(v17, v19)
                                                                    }
                                                                }
                                                            case "hello":
                                                                var v157 int64 = 1
                                                                _ = v157
                                                            default:
                                                                v5 = append(v5, 608)
                                                            }
                                                            _ = v153
                                                        }
                                                    } else {
                                                        v158 := GenS19{F0: 124, F1: true, F2: true}
                                                        _ = v158
                                                    }
                                                }
                                                if (666 <= -9223372036854775808) {
                                                    var v159 int32 = 966
                                                    _ = v159
                                                }
                                                _ = v152
                                            case 42:
                                                v1 -= v1
                                            }
                                        default:
                                            v160 := GenS6{F0: 9223372036854775807, F1: -105, F2: (v3 << ((42) & 63))}
                                            _ = v160
                                        }
                                    }
                                }
                                _ = fi148
                            }
                        case "":
                            var v161 uint8 = 42
                            _ = v161
                        }
                        var v162 uint32 = 42
                        v5 = append(v5, 792)
                        _ = fi146
                        _ = v162
                    }
                    _ = v144
                    _ = v145
                }
                for fi163 := 0; fi163 < 7; fi163++ {
                    if false {
                        _ = len(v5)
                        _ = len(v5)
                    } else {
                        v4 = v4
                        for fi164 := 0; fi164 < 1; fi164++ {
                            var v165 int64 = 9223372036854775807
                            _ = fi164
                            _ = v165
                        }
                        if false {
                            switch -2147483648 {
                            case -95:
                                v8 = (v17 + ("foo bar" + v8))
                            }
                            for fi166 := 0; fi166 < 7; fi166++ {
                                if (!false) {
                                    v167 := GenS20{F0: -1, F1: 0}
                                    _ = v167
                                } else {
                                    v19 += 1.5
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                                _ = fi166
                            }
                            fmt.Println(v8)
                        }
                    }
                    if (-1 == -1) {
                        var v168 float32 = 0.0
                        if true {
                            for fi169 := 0; fi169 < 18; fi169++ {
                                switch "foo bar" {
                                case "a":
                                    v170 := GenS2{F0: (^(fi169))}
                                    _ = v170
                                case "hello":
                                    if (((880 < v4) || ((2147483647 < 4294967295) && ((fi163 >= (^(fi169))) && true))) || true) {
                                        if true {
                                            fmt.Println(v168, fi169)
                                            v171 := GenS20{F0: genFn0(int32(v4), 0), F1: 65535}
                                            v17 = "a"
                                            _ = v171
                                        }
                                        var v172 bool = false
                                        _ = v172
                                    }
                                }
                                switch v17 {
                                case "hello":
                                    var v173 bool = true
                                    _ = v173
                                case "foo bar":
                                    fmt.Println(v168)
                                }
                                _ = fi169
                            }
                            var v174 int16 = v3
                            v175 := GenS4{F0: 176, F1: 463, F2: 0}
                            _ = v174
                            _ = v175
                        } else {
                            if ((^(v3)) >= -368) {
                                var v176 int64 = 42
                                _ = v176
                            }
                        }
                        _ = v168
                    } else {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        switch "a" {
                        case "":
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        case "foo bar":
                            for fi177 := 0; fi177 < 16; fi177++ {
                                v178 := GenS19{F0: -383, F1: false, F2: false}
                                _ = fi177
                                _ = v178
                            }
                        }
                        var v179 int = fi163
                        _ = v179
                    }
                    _ = fi163
                }
            }
        case 4294967295:
            if false {
                var v180 uint64 = v4
                var v181 uint64 = v4
                switch v17 {
                case "hello":
                    v5 = append(v5, 2)
                default:
                    v17 = (v17 + v17)
                }
                _ = v180
                _ = v181
            }
        }
        v4 = v4
    }
    switch (v17 + ("a" + v17)) {
    case "":
        v19 -= v19
    case "\n\t":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    v3 = 1000
    v3 += (v3 - (v3 + (^(v3))))
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v136
    _ = v137
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
