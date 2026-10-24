package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int8 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 int; F1 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16; F1 int64; F2 int16 }
type GenS7 struct { F0 int16; F1 uint16 }
type GenS8 struct { F0 bool; F1 int64; F2 string }
type GenS9 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 bool; F1 uint32 }
type GenS12 struct { F0 string; F1 string }
type GenS13 struct { F0 int64; F1 int64 }
type GenS14 struct { F0 int; F1 int16; F2 uint8 }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 uint; F1 uint }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS19 struct { F0 int32 }
type GenS20 struct { F0 bool; F1 string }
type GenS21 struct { F0 uint32; F1 uint16; F2 uint }
type GenS22 struct { F0 uint }
type GenS23 struct { F0 int64; F1 string; F2 bool }
type GenS24 struct { F0 int16 }

func main() {
    v1 := []uint32{4294967295, 65535}
    v2 := make(map[int]uint32)
    var v3 bool = (((!((42 >= 65535) && (!(!false)))) && false) || false)
    var v4 float32 = (float32(0.1) / 1e10)
    v2[127] = 2147483647
    var v5 uint64 = 1
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := GenS1{F0: v5, F1: (0 > -1), F2: -1}
        if true {
            var v8 int8 = 57
            fmt.Println(v4, v4, v4)
            _ = v8
        }
        v3 = ("hello" == ("hello" + ("" + "foo bar")))
        _ = fi6
        _ = v7
    }
    fmt.Println(v4, v5)
    for fi9 := 0; fi9 < 14; fi9++ {
        v5 += 920
        var v10 uint32 = 65535
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            if true {
                fmt.Println(fi12, v4, v4)
                _ = len(v2)
            }
            if (v3 || v3) {
                if (!(-2147483648 != int32(v5))) {
                    v13 := GenS1{F0: v5, F1: false, F2: 127}
                    _ = v13
                } else {
                    var v14 uint16 = 42
                    var v15 float64 = 1.5
                    v16 := GenS1{F0: v5, F1: (true && true), F2: -1}
                    _ = v14
                    _ = v15
                    _ = v16
                }
                if v3 {
                    switch "" {
                    case "hello":
                        if v3 {
                            var v17 uint64 = 18446744073709551615
                            _ = len(v2)
                            v3 = true
                            switch (fi11 - fi12) {
                            case 175:
                                if len(v1) > 0 {
                                    v1[0] = 1
                                }
                            case -9223372036854775808:
                                switch fi11 {
                                case -9223372036854775808:
                                    _ = len(v2)
                                case -239:
                                    var v18 float64 = (((float64(v4) / 1e10) + 3.14159) / (1.5 / float64(v4)))
                                    _ = v18
                                default:
                                    v19 := GenS1{F0: v5, F1: true, F2: -128}
                                    _ = v19
                                }
                            default:
                                v20 := GenS1{F0: v17, F1: ((860 == 65535) || false), F2: 0}
                                _ = v20
                            }
                            v21 := GenS2{F0: 117, F1: 9223372036854775807}
                            _ = v17
                            _ = v21
                        } else {
                            v4 = 1.5
                            var v22 int = (fi12 * -1)
                            v4 = v4
                            _ = v22
                        }
                    case "foo bar":
                        v2[(fi12 << ((42) & 63))] = 505
                    }
                    v3 = v3
                    fmt.Println(v4)
                } else {
                    v23 := v2[(fi12 << ((2) & 63))]
                    _ = v23
                }
                v4 += 1.5
            } else {
                v1 = append(v1, 545)
                v24 := GenS3{F0: 740, F1: 1}
                v3 = (v3 || v3)
                _ = v24
            }
            _ = fi12
        }
        v25 := v2[(fi11 + 32768)]
        if true {
            if v3 {
                for fi26 := 0; fi26 < 17; fi26++ {
                    switch fi26 {
                    case -274:
                        if (830 > -127) {
                            var v27 int64 = 0
                            switch "" {
                            case "foo bar":
                                var v28 uint16 = 0
                                _ = v28
                            default:
                                var v29 uint = 176
                                _ = v29
                            }
                            v30 := GenS3{F0: 255, F1: (^(v27))}
                            switch (("foo bar" + ("hello" + "foo bar")) + ("foo bar" + "foo bar")) {
                            case "":
                                var v31 string = "foo bar"
                                _ = v31
                            default:
                                v32 := GenS3{F0: 4294967295, F1: (v27 % (((v27 << ((220) & 63))) | 1))}
                                _ = v32
                            }
                            if (!v3) {
                                fmt.Println(v4, v4)
                                var v33 uint32 = (v25 / (((v25 | (^(v25)))) | 1))
                                v27 = (-(v27))
                                _ = v33
                            } else {
                                delete(v2, (fi26 - (fi11 * 332)))
                                if (170 <= 42) {
                                    if (!v3) {
                                        fmt.Println(v25, v3)
                                        v5 += (^(v5))
                                        var v34 int = (fi11 ^ 886)
                                        switch -2147483648 {
                                        case 481:
                                            var v35 int = fi26
                                            _ = v35
                                        case 42:
                                            var v36 float64 = (0.1 / 0.1)
                                            _ = v36
                                        case 51:
                                            v37 := v2[(fi26 ^ fi11)]
                                            _ = v37
                                        default:
                                            v27 = v27
                                            v25 -= 356
                                            switch "\n\t" {
                                            case "foo bar":
                                                fmt.Println(v25)
                                            default:
                                                if true {
                                                    if v3 {
                                                        var v38 float64 = 0.0
                                                        _ = v38
                                                    } else {
                                                        v3 = v3
                                                        if (0.1 <= ((0.0 + 3.14159) - 0.0)) {
                                                            delete(v2, 516)
                                                            var v39 bool = false
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            _ = v39
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                        _ = v34
                                    }
                                    switch fi11 {
                                    case -9223372036854775808:
                                        fmt.Println(v4, v27, v5)
                                    default:
                                        v5 -= v5
                                        var v40 int32 = 364
                                        delete(v2, (fi11 & 4294967295))
                                        _ = v40
                                    }
                                } else {
                                    switch ("\n\t" + ("hello" + ("a" + ("" + "hello")))) {
                                    case "":
                                        if true {
                                            v4 += float32(float64((v4 - ((1e10 - v4) * -2.25))))
                                            fmt.Println(v5)
                                            v41 := GenS4{F0: -1, F1: (v27 + v27)}
                                            _ = v41
                                        } else {
                                            if false {
                                                var v42 uint = 619
                                                switch "a" {
                                                case "hello":
                                                    var v43 string = "hello"
                                                    _ = v43
                                                case "\n\t":
                                                    switch (^(fi26)) {
                                                    case -1:
                                                        if false {
                                                            v3 = v3
                                                        }
                                                    case 32768:
                                                        v44 := GenS5{F0: (1000 >= 123)}
                                                        _ = v44
                                                    default:
                                                        switch -2147483648 {
                                                        case 127:
                                                            if (false && v3) {
                                                                v45 := GenS6{F0: 42, F1: (v27 ^ 1), F2: 1}
                                                                var v46 uint32 = v25
                                                                v3 = (true && false)
                                                                var v47 uint32 = 2147483648
                                                                _ = v45
                                                                _ = v46
                                                                _ = v47
                                                            } else {
                                                                _ = len(v2)
                                                                switch (fi11 % ((fi11) | 1)) {
                                                                case 42:
                                                                    v48 := GenS2{F0: -833, F1: (^(v27))}
                                                                    _ = v48
                                                                case 0:
                                                                    var v49 int16 = 32767
                                                                    _ = v49
                                                                case 652:
                                                                    v4 = (1.5 / ((v4 + v4) * (float32(float64(float32(float64(v4)))) + v4)))
                                                                }
                                                            }
                                                        case 42:
                                                            var v50 uint = 255
                                                            _ = v50
                                                        default:
                                                            v27 = (v27 * (v27 ^ v27))
                                                        }
                                                    }
                                                case "foo bar":
                                                    if (2147483647 <= 0) {
                                                        fmt.Println(fi11)
                                                        if (1e10 == -2.25) {
                                                            if ((18446744073709551615 != 170) && v3) {
                                                                var v51 int32 = 76
                                                                _ = v51
                                                            } else {
                                                                v42 += (v42 & v42)
                                                                var v52 float32 = v4
                                                                var v53 bool = v3
                                                                v54 := v2[(fi11 - (fi26 - fi11))]
                                                                _ = v52
                                                                _ = v53
                                                                _ = v54
                                                            }
                                                            var v55 uint = v42
                                                            var v56 uint8 = 204
                                                            _ = v55
                                                            _ = v56
                                                        } else {
                                                            var v57 uint16 = 2
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                _ = len(v2)
                                                            case "a":
                                                                v58 := GenS6{F0: -1, F1: 0, F2: -1}
                                                                _ = v58
                                                            case "hello":
                                                                var v59 bool = true
                                                                _ = v59
                                                            default:
                                                                var v60 string = ("" + "a")
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                var v61 float32 = (((v4 / (v4 / v4)) * ((v4 / v4) * 3.14159)) / v4)
                                                                fmt.Println(v5)
                                                                _ = v60
                                                                _ = v61
                                                            }
                                                            var v62 uint64 = 694
                                                            _ = v57
                                                            _ = v62
                                                        }
                                                        var v63 uint32 = uint32(v42)
                                                        _ = v63
                                                    } else {
                                                        switch "" {
                                                        case "foo bar":
                                                            if true {
                                                                v64 := GenS1{F0: v5, F1: v3, F2: 69}
                                                                switch (fi26 ^ int(v5)) {
                                                                case 127:
                                                                    v65 := GenS6{F0: -744, F1: v27, F2: 889}
                                                                    _ = v65
                                                                case 32767:
                                                                    _ = len(v1)
                                                                case 42:
                                                                    v27 += (-(v27))
                                                                }
                                                                _ = v64
                                                            }
                                                        default:
                                                            delete(v2, 2147483647)
                                                            v25 -= (v25 ^ (v25 + 797))
                                                            switch fi26 {
                                                            case 127:
                                                                var v66 int = 784
                                                                _ = v66
                                                            default:
                                                                var v67 string = ""
                                                                _ = v67
                                                            }
                                                            v68 := GenS1{F0: (v5 % ((v5) | 1)), F1: v3, F2: 56}
                                                            v3 = false
                                                            _ = v68
                                                        }
                                                    }
                                                }
                                                var v69 uint64 = v5
                                                _ = v42
                                                _ = v69
                                            }
                                        }
                                    case "\n\t":
                                        var v70 uint16 = 1
                                        _ = v70
                                    }
                                }
                                v2[fi11] = (v25 - (v25 << ((65535) & 63)))
                            }
                            var v71 int64 = v27
                            switch "\n\t" {
                            case "":
                                if (v3 || (false || v3)) {
                                    if v3 {
                                        v5 = (^(v5))
                                        switch "\n\t" {
                                        case "":
                                            v27 = (v27 << ((4294967295) & 63))
                                        case "a":
                                            var v72 bool = v3
                                            _ = v72
                                        default:
                                            switch ("a" + (("foo bar" + "hello") + ("foo bar" + "a"))) {
                                            case "a":
                                                if (!(1 == 1)) {
                                                    if false {
                                                        var v73 uint = 168
                                                        var v74 uint32 = 2147483648
                                                        switch "\n\t" {
                                                        case "a":
                                                            v2[(^(fi26))] = (^(v25))
                                                        case "foo bar":
                                                            if v3 {
                                                                var v75 int = 1
                                                                _ = v75
                                                            } else {
                                                                switch (fi11 & fi26) {
                                                                case -801:
                                                                    var v76 uint16 = 0
                                                                    _ = v76
                                                                }
                                                                fmt.Println(fi11, v73)
                                                                fmt.Println(v27, v27, v74)
                                                            }
                                                        case "":
                                                            v77 := GenS5{F0: v3}
                                                            _ = v77
                                                        }
                                                        _ = v73
                                                        _ = v74
                                                    }
                                                } else {
                                                    var v78 uint16 = 32767
                                                    _ = v78
                                                }
                                            case "":
                                                fmt.Println(v27)
                                            default:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                        }
                                        v5 -= v5
                                    }
                                    var v79 float32 = 0.0
                                    fmt.Println(fi26, v3, v3)
                                    _ = v79
                                }
                            case "hello":
                                if (false || ("hello" != ("hello" + ("\n\t" + "")))) {
                                    fmt.Println(v27, v3, v5)
                                } else {
                                    v4 = 3.14159
                                }
                            default:
                                v27 = -498
                                var v80 float64 = (0.0 * ((1.5 * 3.14159) - float64(float32(3.14159))))
                                v5 -= v5
                                _ = v80
                            }
                            if true {
                                fmt.Println(v3)
                            } else {
                                v71 -= v71
                                v3 = (!v3)
                                v27 = 179
                            }
                            _ = v27
                            _ = v30
                            _ = v71
                        } else {
                            v3 = v3
                        }
                    default:
                        var v81 uint16 = 640
                        _ = v81
                    }
                    _ = fi26
                }
                switch (("hello" + "\n\t") + "hello") {
                case "\n\t":
                    var v82 uint = 172
                    _ = v82
                case "foo bar":
                    if (42 <= 9223372036854775807) {
                        fmt.Println(fi11, v5, v5)
                        switch ("" + ("hello" + "")) {
                        case "a":
                            v25 = (v25 / (((^(v25))) | 1))
                        case "\n\t":
                            fmt.Println(v5, v3, v3)
                        default:
                            fmt.Println(v25)
                        }
                        v4 += 3.14159
                    }
                case "hello":
                    v25 -= v25
                }
            } else {
                v83 := GenS6{F0: 32767, F1: 0, F2: 42}
                v25 += 65535
                _ = v83
            }
            for fi84 := 0; fi84 < 2; fi84++ {
                var v85 uint16 = 2
                _ = fi84
                _ = v85
            }
            v2[9223372036854775807] = (v25 >> ((4294967295) & 63))
        } else {
            var v86 int8 = 0
            v86 += (v86 * (v86 | v86))
            _ = v86
        }
        _ = fi11
        _ = v25
    }
    switch "a" {
    case "foo bar":
        v87 := GenS7{F0: -494, F1: 32767}
        _ = v87
    case "hello":
        v5 = v5
    default:
        if v3 {
            var v88 uint16 = 32767
            v2[-871] = 2147483647
            _ = v88
        } else {
            v89 := GenS2{F0: -1, F1: -9223372036854775808}
            _ = v89
        }
    }
    v3 = true
    v3 = false
    if (!(!(128 < 42))) {
        if v3 {
            v5 = v5
            var v90 int8 = 100
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v90
        }
    } else {
        switch ((("" + ("\n\t" + "a")) + ("" + "hello")) + "\n\t") {
        case "a":
            var v91 float32 = (1e10 + 0.1)
            _ = v91
        }
        var v92 float64 = 0.0
        _ = v92
    }
    for fi93 := 0; fi93 < 18; fi93++ {
        var v94 uint32 = 549
        switch ("\n\t" + "") {
        case "hello":
            if (127 <= 255) {
                var v95 string = "hello"
                if true {
                    v96 := GenS3{F0: 18446744073709551615, F1: 585}
                    _ = v96
                }
                for fi97 := 0; fi97 < 15; fi97++ {
                    fmt.Println(v95, v3)
                    v98 := GenS8{F0: (0.1 <= v4), F1: 42, F2: "\n\t"}
                    if (!(11 >= 0)) {
                        fmt.Println(v4, v95)
                        var v99 int32 = 2147483647
                        _ = v99
                    } else {
                        v100 := v2[fi97]
                        _ = v100
                    }
                    _ = fi97
                    _ = v98
                }
                _ = v95
            }
        case "":
            for fi101 := 0; fi101 < 3; fi101++ {
                v102 := GenS9{F0: (v94 + (v94 >> ((1) & 63))), F1: 1, F2: 65535}
                _ = fi101
                _ = v102
            }
        case "foo bar":
            var v103 uint8 = 28
            _ = v103
        }
        for fi104 := 0; fi104 < 14; fi104++ {
            fmt.Println(fi93)
            switch (fi104 | fi104) {
            case -1:
                var v105 uint = 4294967295
                _ = v105
            case 0:
                _ = len(v2)
            }
            var v106 int32 = 1
            _ = fi104
            _ = v106
        }
        _ = fi93
        _ = v94
    }
    v107 := GenS10{F0: 1}
    v108 := GenS6{F0: 32767, F1: -979, F2: -1}
    if (127 >= 126) {
        _ = len(v1)
        for fi109 := 0; fi109 < 11; fi109++ {
            if v3 {
                if v3 {
                    if (true || ((3.14159 + -2.25) >= 1.5)) {
                        v110 := GenS5{F0: (!((!(!v3)) || false))}
                        var v111 int8 = -52
                        if true {
                            var v112 int8 = (v111 + (v111 << ((255) & 63)))
                            var v113 float64 = ((0.1 + 0.0) * 1.5)
                            _ = v112
                            _ = v113
                        }
                        _ = v110
                        _ = v111
                    }
                    var v114 int8 = 100
                    if v3 {
                        switch -352 {
                        case -979:
                            for fi115 := 0; fi115 < 18; fi115++ {
                                var v116 uint32 = 0
                                fmt.Println(v116)
                                _ = fi115
                                _ = v116
                            }
                        case 1:
                            switch (fi109 + fi109) {
                            case 552:
                                if (7 == 2) {
                                    v4 += v4
                                } else {
                                    v4 = v4
                                    if true {
                                        if v3 {
                                            fmt.Println(fi109, v3)
                                            if v3 {
                                                v4 += float32((1.5 * 0.0))
                                                var v117 uint = 416
                                                v117 = 255
                                                _ = v117
                                            } else {
                                                if v3 {
                                                    v118 := GenS11{F0: (!v3), F1: 352}
                                                    v119 := GenS4{F0: (fi109 << ((153) & 63)), F1: 42}
                                                    _ = v118
                                                    _ = v119
                                                } else {
                                                    var v120 int64 = 1
                                                    _ = v120
                                                }
                                                if v3 {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                } else {
                                                    v121 := GenS8{F0: (1.5 < 1.5), F1: 0, F2: ("a" + (("hello" + ("a" + "hello")) + "hello"))}
                                                    var v122 int8 = 0
                                                    var v123 float64 = float64(-2.25)
                                                    _ = v121
                                                    _ = v122
                                                    _ = v123
                                                }
                                            }
                                            var v124 string = "a"
                                            _ = v124
                                        }
                                        v114 += v114
                                        if (2147483647 >= 42) {
                                            switch 1 {
                                            case 816:
                                                switch "foo bar" {
                                                case "foo bar":
                                                    if true {
                                                        if true {
                                                            v114 -= (v114 + -39)
                                                            var v125 int32 = -969
                                                            var v126 float64 = (0.1 / float64(1e10))
                                                            _ = v125
                                                            _ = v126
                                                        } else {
                                                            v127 := GenS12{F0: ((("hello" + "") + "a") + "hello"), F1: "foo bar"}
                                                            v128 := GenS3{F0: 18446744073709551615, F1: int64(v5)}
                                                            v114 -= (v114 << ((65535) & 63))
                                                            _ = v127
                                                            _ = v128
                                                        }
                                                    }
                                                case "a":
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                case "":
                                                    var v129 int32 = 517
                                                    _ = v129
                                                }
                                            case 42:
                                                v5 = v5
                                            case 32767:
                                                switch -824 {
                                                case 32768:
                                                    switch ("\n\t" + ("foo bar" + "hello")) {
                                                    case "hello":
                                                        v130 := GenS7{F0: -190, F1: 65535}
                                                        _ = v130
                                                    case "\n\t":
                                                        v131 := GenS13{F0: 9223372036854775807, F1: -9223372036854775808}
                                                        _ = v131
                                                    case "":
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    default:
                                                        var v132 float32 = 3.14159
                                                        _ = v132
                                                    }
                                                default:
                                                    for fi133 := 0; fi133 < 17; fi133++ {
                                                        v1 = append(v1, 2147483647)
                                                        if (!v3) {
                                                            v134 := GenS4{F0: (fi133 << ((4294967295) & 63)), F1: 0}
                                                            if (v3 && (42 > (v114 << ((1) & 63)))) {
                                                                if false {
                                                                    if len(v1) > 0 {
                                                                        v1[0] = 266
                                                                    }
                                                                    var v135 float32 = v4
                                                                    v114 = v114
                                                                    _ = v135
                                                                }
                                                                var v136 int16 = 255
                                                                v137 := GenS8{F0: ((v4 * -2.25) < 0.0), F1: 727, F2: (("foo bar" + "a") + ("hello" + "foo bar"))}
                                                                _ = v136
                                                                _ = v137
                                                            }
                                                            _ = v134
                                                        }
                                                        if (!v3) {
                                                            v138 := GenS14{F0: 0, F1: -1, F2: 255}
                                                            v5 = v5
                                                            var v139 int64 = -9223372036854775808
                                                            _ = v138
                                                            _ = v139
                                                        }
                                                        _ = fi133
                                                    }
                                                }
                                            default:
                                                var v140 int8 = -106
                                                _ = v140
                                            }
                                        } else {
                                            for fi141 := 0; fi141 < 4; fi141++ {
                                                switch (fi141 - fi109) {
                                                case 128:
                                                    v142 := GenS15{F0: -419}
                                                    var v143 uint16 = 65535
                                                    if v3 {
                                                        v3 = (255 == 972)
                                                        var v144 int32 = 0
                                                        _ = v144
                                                    } else {
                                                        var v145 uint8 = 127
                                                        fmt.Println(v114, v114, v4)
                                                        if v3 {
                                                            v146 := GenS16{F0: 65535, F1: 0}
                                                            v5 -= (v5 | v5)
                                                            _ = v146
                                                        }
                                                        var v147 string = ""
                                                        fmt.Println(v145, v5, fi109)
                                                        _ = v145
                                                        _ = v147
                                                    }
                                                    v4 -= 1e10
                                                    _ = v142
                                                    _ = v143
                                                case 127:
                                                    var v148 uint16 = 255
                                                    _ = v148
                                                default:
                                                    v4 += (v4 - v4)
                                                }
                                                var v149 uint16 = 32767
                                                _ = fi141
                                                _ = v149
                                            }
                                        }
                                    }
                                }
                            default:
                                v3 = false
                            }
                        }
                        var v150 float64 = (0.0 * 0.0)
                        _ = v150
                    }
                    _ = v114
                }
                if v3 {
                    var v151 bool = true
                    if true {
                        switch (fi109 - (fi109 << ((65535) & 63))) {
                        case 127:
                            for fi152 := 0; fi152 < 19; fi152++ {
                                var v153 bool = v151
                                switch (fi109 % ((fi152) | 1)) {
                                case 395:
                                    var v154 uint = 431
                                    _ = v154
                                default:
                                    var v155 string = ""
                                    _ = v155
                                }
                                _ = fi152
                                _ = v153
                            }
                        default:
                            for fi156 := 0; fi156 < 2; fi156++ {
                                var v157 bool = (!v151)
                                if v151 {
                                    v151 = ((float32((3.14159 / (0.0 + 0.1))) + (v4 - ((float32((0.0 / (0.0 - float64(v4)))) + v4) / (0.0 - float32((1.5 + (1.5 * 0.1))))))) >= 0.1)
                                    var v158 float32 = v4
                                    var v159 uint16 = 442
                                    _ = v158
                                    _ = v159
                                }
                                if (true && v157) {
                                    if (!(-2.25 > (float64(1e10) * (-2.25 + 1.5)))) {
                                        _ = len(v1)
                                        v160 := v2[(fi109 ^ -13)]
                                        v160 = (v160 << ((2) & 63))
                                        _ = v160
                                    }
                                    if false {
                                        fmt.Println(fi109, v4)
                                    } else {
                                        switch "" {
                                        case "foo bar":
                                            v4 = (v4 / v4)
                                            var v161 uint8 = 255
                                            var v162 uint32 = 2147483647
                                            _ = v161
                                            _ = v162
                                        }
                                        if (1e10 != -2.25) {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                        if (v3 || (false && false)) {
                                            switch "foo bar" {
                                            case "\n\t":
                                                v5 = v5
                                            case "hello":
                                                var v163 float32 = ((v4 - (1.5 - v4)) / 1.0)
                                                var v164 uint16 = 618
                                                _ = v163
                                                _ = v164
                                            default:
                                                var v165 int = (fi109 | -889)
                                                _ = v165
                                            }
                                            if v151 {
                                                fmt.Println(fi109, v157)
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v5 = 18446744073709551615
                                            }
                                            switch "foo bar" {
                                            case "hello":
                                                v166 := GenS17{F0: 1}
                                                fmt.Println(v5, v157)
                                                _ = v166
                                            case "":
                                                fmt.Println(v151, v151, fi109)
                                            }
                                        } else {
                                            v157 = ("\n\t" > "\n\t")
                                        }
                                    }
                                    if false {
                                        if v151 {
                                            v5 -= v5
                                            var v167 float64 = 1.5
                                            _ = v167
                                        } else {
                                            var v168 float32 = (((v4 + float32(((float64(3.14159) + 3.14159) + float64(0.0)))) / v4) - 3.14159)
                                            switch "\n\t" {
                                            case "":
                                                if ((fi109 * (fi109 % (((fi156 | fi109)) | 1))) < (fi109 - int(v5))) {
                                                    if (!true) {
                                                        v169 := GenS18{F0: 65535, F1: -14, F2: 255}
                                                        _ = v169
                                                    } else {
                                                        var v170 string = "\n\t"
                                                        _ = v170
                                                    }
                                                } else {
                                                    switch 9223372036854775807 {
                                                    case 0:
                                                        v5 = (v5 - v5)
                                                    case 42:
                                                        v5 = (^(v5))
                                                    default:
                                                        v157 = ((0 < 1) && v157)
                                                    }
                                                    _ = len(v2)
                                                }
                                            default:
                                                v171 := GenS13{F0: -9223372036854775808, F1: -127}
                                                _ = v171
                                            }
                                            v172 := v2[(fi109 | fi156)]
                                            _ = v168
                                            _ = v172
                                        }
                                        v173 := GenS9{F0: 0, F1: -128, F2: 255}
                                        v174 := GenS3{F0: 2, F1: 9223372036854775807}
                                        _ = v173
                                        _ = v174
                                    } else {
                                        if v3 {
                                            if (0 != 13) {
                                                v175 := GenS19{F0: 1}
                                                fmt.Println(fi109, fi109, v5)
                                                _ = v175
                                            }
                                        } else {
                                            v2[fi109] = 766
                                            v176 := GenS20{F0: v157, F1: "\n\t"}
                                            var v177 uint16 = 65535
                                            _ = v176
                                            _ = v177
                                        }
                                    }
                                } else {
                                    v157 = v151
                                    _ = len(v2)
                                    v157 = (uint32(v5) == 2)
                                }
                                _ = fi156
                                _ = v157
                            }
                        }
                    }
                    _ = v151
                }
                if v3 {
                    var v178 uint32 = 4294967295
                    if v3 {
                        fmt.Println(v4, v178, v3)
                        v179 := GenS8{F0: false, F1: 9223372036854775807, F2: "a"}
                        _ = v179
                    }
                    _ = v178
                }
            } else {
                for fi180 := 0; fi180 < 6; fi180++ {
                    if (!true) {
                        _ = len(v1)
                    }
                    var v181 uint = 38
                    var v182 uint16 = 1
                    v1 = append(v1, 42)
                    _ = fi180
                    _ = v181
                    _ = v182
                }
            }
            _ = fi109
        }
    } else {
        if v3 {
            if (1 == 155) {
                if (!v3) {
                    v5 = 0
                    fmt.Println(v4, v4)
                    var v183 int32 = -397
                    _ = v183
                } else {
                    _ = len(v1)
                    v3 = v3
                    v184 := GenS21{F0: 2147483647, F1: 2, F2: 65535}
                    _ = v184
                }
            } else {
                v185 := GenS22{F0: 0}
                if (!(-32768 < -1)) {
                    var v186 uint8 = 2
                    var v187 uint32 = uint32(v186)
                    v187 = 269
                    _ = v186
                    _ = v187
                }
                for fi188 := 0; fi188 < 15; fi188++ {
                    var v189 uint16 = 910
                    var v190 int = (fi188 >> ((4294967295) & 63))
                    if (false || v3) {
                        if (193 == uint8(v5)) {
                            if false {
                                switch fi188 {
                                case 1:
                                    var v191 int16 = 675
                                    _ = v191
                                case 2147483647:
                                    for fi192 := 0; fi192 < 5; fi192++ {
                                        v193 := GenS23{F0: int64(v190), F1: ((("foo bar" + "") + ("a" + "")) + "\n\t"), F2: (4294967295 != 790)}
                                        _ = len(v1)
                                        _ = fi192
                                        _ = v193
                                    }
                                case 454:
                                    fmt.Println(fi188, v5, v190)
                                default:
                                    if ((v190 != fi188) && (32767 != 42)) {
                                        v194 := GenS1{F0: (v5 ^ 940), F1: false, F2: -128}
                                        _ = v194
                                    } else {
                                        v195 := GenS2{F0: 1, F1: 401}
                                        var v196 int = v190
                                        fmt.Println(v189)
                                        _ = v195
                                        _ = v196
                                    }
                                }
                            }
                            var v197 uint64 = v5
                            _ = v197
                        } else {
                            var v198 int8 = 100
                            var v199 int32 = 2147483647
                            _ = v198
                            _ = v199
                        }
                        var v200 string = (((("hello" + ("a" + "\n\t")) + "\n\t") + "") + "")
                        _ = v200
                    } else {
                        if false {
                            v5 = v5
                            v1 = append(v1, 75)
                            fmt.Println(v4)
                        }
                        fmt.Println(fi188)
                        if (!false) {
                            v190 = (^(fi188))
                            var v201 float32 = -2.25
                            v202 := GenS15{F0: -1}
                            _ = v201
                            _ = v202
                        }
                    }
                    _ = fi188
                    _ = v189
                    _ = v190
                }
                _ = v185
            }
        }
        var v203 int = -2147483648
        _ = v203
    }
    for fi204 := 0; fi204 < 11; fi204++ {
        v205 := GenS24{F0: -999}
        var v206 uint = 514
        var v207 int32 = int32(fi204)
        _ = fi204
        _ = v205
        _ = v206
        _ = v207
    }
    var v208 string = "foo bar"
    v1 = append(v1, 918)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v208:", v208)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v107
    _ = v108
    _ = v2
    _ = v208
    _ = v3
    _ = v4
    _ = v5
}
