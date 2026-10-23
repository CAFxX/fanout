package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 string; F1 int; F2 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int; F1 uint64; F2 uint64 }
type GenS7 struct { F0 int16; F1 bool }
type GenS8 struct { F0 uint16; F1 uint8 }
type GenS9 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS10 struct { F0 uint16; F1 uint8 }
type GenS11 struct { F0 bool; F1 int16 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int }
func genFn0() uint8 {
    if false {
        var v6 uint32 = 21
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := GenS1{F0: 601}
            switch "hello" {
            case "":
                if true {
                    fmt.Println(v6)
                    v6 = v6
                }
            case "a":
                v6 = v6
            }
            _ = fi7
            _ = v8
        }
        _ = v6
    } else {
        var v9 uint64 = 18446744073709551615
        _ = v9
    }
    var v10 uint8 = 128
    switch 127 {
    case 127:
        fmt.Println(v10)
    case -1:
        if (2147483648 == 827) {
            if (((true && ("a" <= "a")) || true) || (true || false)) {
                v11 := []int{1, 160}
                _ = v11
            }
            v10 = v10
            v12 := GenS1{F0: 0}
            _ = v12
        }
    }
    _ = v10
    return 128
}

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = 3.14159
    v3 := make(map[int]uint32)
    var v4 int = 9223372036854775807
    v5 := []int{(v4 ^ -911)}
    if false {
        var v13 float64 = float64(((v2 + ((v2 / 1.0) * (float32(-2.25) / v2))) - 0.0))
        for fi14 := 0; fi14 < 1; fi14++ {
            if len(v5) > 0 {
                v5[0] = -2147483648
            }
            _ = fi14
        }
        if len(v5) > 0 {
            v5[0] = (-(v4))
        }
        _ = v13
    } else {
        if (float32(1.5) >= float32((1e10 * 3.14159))) {
            v5 = append(v5, (v4 & (v4 ^ (v4 ^ (v4 * v4)))))
            fmt.Println(v2, v2)
        } else {
            var v15 string = "hello"
            var v16 uint32 = 289
            _ = v15
            _ = v16
        }
    }
    if false {
        v1[42] = (true && (!(42 < 0)))
        fmt.Println(v4)
    } else {
        if (!false) {
            v4 = (v4 % ((v4) | 1))
        } else {
            switch (("\n\t" + "\n\t") + (("hello" + "foo bar") + "hello")) {
            case "":
                fmt.Println(v2, v4)
            case "foo bar":
                if (0 <= 32767) {
                    v4 += v4
                } else {
                    v17 := GenS1{F0: 256}
                    var v18 uint32 = 0
                    switch (("foo bar" + ("" + "\n\t")) + "a") {
                    case "":
                        if (!false) {
                            switch v4 {
                            case 295:
                                v19 := GenS1{F0: 32767}
                                _ = v19
                            case -1:
                                switch "a" {
                                case "":
                                    for fi20 := 0; fi20 < 19; fi20++ {
                                        var v21 int = fi20
                                        _ = fi20
                                        _ = v21
                                    }
                                }
                            case 2147483647:
                                var v22 int16 = -32768
                                _ = v22
                            }
                            fmt.Println(v2)
                        } else {
                            if (!false) {
                                v23 := GenS1{F0: 219}
                                v24 := GenS2{F0: "\n\t", F1: (v4 >> ((255) & 63))}
                                _ = v23
                                _ = v24
                            } else {
                                for fi25 := 0; fi25 < 12; fi25++ {
                                    v4 += fi25
                                    v26 := GenS1{F0: 42}
                                    fmt.Println(fi25, v4)
                                    _ = fi25
                                    _ = v26
                                }
                                if (!(-36 == 1)) {
                                    var v27 uint8 = 2
                                    if false {
                                        v2 += (3.14159 + 0.0)
                                    }
                                    _ = v27
                                } else {
                                    v28 := GenS3{F0: "foo bar", F1: (v4 % ((v4) | 1)), F2: v18}
                                    fmt.Println(v18, v2)
                                    if ((3.14159 * (-2.25 / float64(v2))) >= float64(0.0)) {
                                        _ = len(v5)
                                        v29 := GenS3{F0: ((("\n\t" + "foo bar") + "hello") + "foo bar"), F1: v4, F2: 65535}
                                        _ = v29
                                    } else {
                                        if len(v5) > 0 {
                                            v5[0] = (v4 / ((2147483647) | 1))
                                        }
                                        fmt.Println(v4)
                                        for fi30 := 0; fi30 < 11; fi30++ {
                                            if false {
                                                var v31 float32 = 1.5
                                                var v32 uint64 = uint64(v18)
                                                _ = v31
                                                _ = v32
                                            } else {
                                                switch "" {
                                                case "":
                                                    v2 = 1.5
                                                case "\n\t":
                                                    var v33 uint32 = v18
                                                    _ = v33
                                                case "hello":
                                                    switch ("foo bar" + "a") {
                                                    case "":
                                                        var v34 uint8 = 227
                                                        _ = v34
                                                    }
                                                default:
                                                    fmt.Println(v18, fi30)
                                                }
                                                for fi35 := 0; fi35 < 17; fi35++ {
                                                    switch "hello" {
                                                    case "a":
                                                        if true {
                                                            if (false && (!((1 <= -229) && (-2.25 > 0.0)))) {
                                                                var v36 uint32 = (v18 % ((v18) | 1))
                                                                v37 := GenS4{F0: v18}
                                                                fmt.Println(v2)
                                                                _ = v36
                                                                _ = v37
                                                            }
                                                            switch fi35 {
                                                            case -2147483648:
                                                                v18 = (v18 | v18)
                                                            case 4294967295:
                                                                switch 32767 {
                                                                case -9223372036854775808:
                                                                    if false {
                                                                        switch "a" {
                                                                        case "\n\t":
                                                                            v2 = v2
                                                                        }
                                                                    }
                                                                case 1:
                                                                    v1[(-(fi30))] = false
                                                                case 42:
                                                                    switch "" {
                                                                    case "\n\t":
                                                                        if true {
                                                                            var v38 float32 = float32((1.5 + (1.5 * (float64(float32(0.0)) / 1.0))))
                                                                            fmt.Println(fi35, v4, fi30)
                                                                            _ = len(v5)
                                                                            _ = v38
                                                                        } else {
                                                                            _ = len(v3)
                                                                            if (181 != genFn0()) {
                                                                                var v39 float64 = float64((0.1 / (3.14159 * (v2 - ((v2 * v2) / ((v2 / 1.0) + (1e10 - v2)))))))
                                                                                _ = v39
                                                                            }
                                                                        }
                                                                    case "hello":
                                                                        var v40 int = -1
                                                                        _ = v40
                                                                    case "":
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                    }
                                                                default:
                                                                    delete(v1, fi30)
                                                                }
                                                            }
                                                            var v41 bool = true
                                                            _ = v41
                                                        } else {
                                                            var v42 uint8 = 15
                                                            var v43 uint16 = 0
                                                            _ = v42
                                                            _ = v43
                                                        }
                                                    case "hello":
                                                        delete(v3, (fi35 << ((42) & 63)))
                                                    default:
                                                        v2 = -2.25
                                                    }
                                                    if true {
                                                        _ = len(v3)
                                                        var v44 int = (^(fi35))
                                                        _ = v44
                                                    } else {
                                                        v2 -= 1e10
                                                        v45 := GenS4{F0: (v18 << ((1) & 63))}
                                                        switch -9223372036854775808 {
                                                        case 32767:
                                                            switch ("" + ("" + "foo bar")) {
                                                            case "\n\t":
                                                                v46 := v3[fi35]
                                                                _ = v46
                                                            case "foo bar":
                                                                if ((1000 == 1000) && true) {
                                                                    v47 := GenS2{F0: "", F1: 0}
                                                                    if true {
                                                                        if false {
                                                                            v48 := GenS5{F0: -1}
                                                                            v18 -= (v18 * 1)
                                                                            _ = v48
                                                                        }
                                                                        if (!(941 == 18446744073709551615)) {
                                                                            if (true || false) {
                                                                                if (!(false || false)) {
                                                                                    fmt.Println(fi30, fi35, fi35)
                                                                                    v4 = (v4 ^ (fi30 + v4))
                                                                                    var v49 float32 = v2
                                                                                    _ = v49
                                                                                } else {
                                                                                    v50 := GenS2{F0: ("\n\t" + "\n\t"), F1: (v4 * 9223372036854775807)}
                                                                                    fmt.Println(v2)
                                                                                    _ = v50
                                                                                }
                                                                            }
                                                                            if len(v5) > 0 {
                                                                                v5[0] = (v4 - v4)
                                                                            }
                                                                            var v51 string = (((("hello" + "") + "foo bar") + "\n\t") + "a")
                                                                            var v52 int16 = 42
                                                                            _ = v51
                                                                            _ = v52
                                                                        } else {
                                                                            v18 = (v18 ^ (v18 % ((v18) | 1)))
                                                                            v18 = (v18 / ((v18) | 1))
                                                                            var v53 uint32 = (v18 * 767)
                                                                            _ = v53
                                                                        }
                                                                    }
                                                                    _ = v47
                                                                } else {
                                                                    var v54 float64 = (((0.0 / 1.0) + 3.14159) / (0.1 / 1.5))
                                                                    var v55 int = (fi35 * fi35)
                                                                    _ = v54
                                                                    _ = v55
                                                                }
                                                            default:
                                                                switch (fi35 << ((42) & 63)) {
                                                                case 511:
                                                                    var v56 int16 = 456
                                                                    _ = v56
                                                                case -395:
                                                                    var v57 uint = 4294967295
                                                                    _ = v57
                                                                case 32767:
                                                                    var v58 int8 = 70
                                                                    _ = v58
                                                                }
                                                            }
                                                        default:
                                                            fmt.Println(v2)
                                                        }
                                                        _ = v45
                                                    }
                                                    if (128 >= (fi35 * v4)) {
                                                        if (376 != 32767) {
                                                            v18 += 2147483648
                                                            switch "\n\t" {
                                                            case "":
                                                                var v59 int32 = -2147483648
                                                                _ = v59
                                                            }
                                                            if len(v5) > 0 {
                                                                v5[0] = fi30
                                                            }
                                                        }
                                                        var v60 int = v4
                                                        _ = v60
                                                    } else {
                                                        var v61 int32 = -544
                                                        v62 := GenS2{F0: "\n\t", F1: 2147483647}
                                                        _ = v61
                                                        _ = v62
                                                    }
                                                    _ = fi35
                                                }
                                                _ = len(v5)
                                            }
                                            for fi63 := 0; fi63 < 6; fi63++ {
                                                v4 = fi30
                                                _ = fi63
                                            }
                                            v64 := GenS2{F0: "foo bar", F1: v4}
                                            _ = fi30
                                            _ = v64
                                        }
                                    }
                                    _ = v28
                                }
                                for fi65 := 0; fi65 < 16; fi65++ {
                                    v66 := GenS3{F0: "\n\t", F1: (fi65 * (v4 ^ (v4 * fi65))), F2: v18}
                                    var v67 int64 = 449
                                    _ = fi65
                                    _ = v66
                                    _ = v67
                                }
                            }
                            fmt.Println(v4, v4)
                            var v68 uint = 1
                            _ = v68
                        }
                    default:
                        if false {
                            v18 += (v18 | v18)
                            v18 -= v18
                        } else {
                            for fi69 := 0; fi69 < 8; fi69++ {
                                v2 = 3.14159
                                _ = fi69
                            }
                        }
                    }
                    _ = v17
                    _ = v18
                }
            default:
                _ = len(v1)
            }
            switch "\n\t" {
            case "":
                fmt.Println(v4, v2)
            case "a":
                for fi70 := 0; fi70 < 1; fi70++ {
                    switch ("\n\t" + "\n\t") {
                    case "":
                        var v71 float32 = ((float32((3.14159 + 3.14159)) + -2.25) * (((v2 + float32(0.0)) - -2.25) + float32(1.5)))
                        _ = v71
                    case "hello":
                        for fi72 := 0; fi72 < 11; fi72++ {
                            v4 = fi70
                            if (18446744073709551615 <= 335) {
                                _ = len(v3)
                                if false {
                                    fmt.Println(fi72, v4)
                                    v73 := GenS6{F0: -2147483648, F1: 2, F2: 18446744073709551615}
                                    _ = v73
                                } else {
                                    fmt.Println(fi72)
                                    v74 := v1[(fi70 & v4)]
                                    _ = v74
                                }
                            }
                            _ = fi72
                        }
                    default:
                        for fi75 := 0; fi75 < 13; fi75++ {
                            var v76 int16 = 240
                            if true {
                                var v77 string = ("" + ("" + ""))
                                v76 -= (v76 << ((0) & 63))
                                _ = v77
                            }
                            delete(v3, v4)
                            var v78 int = fi70
                            if len(v5) > 0 {
                                v5[0] = v78
                            }
                            _ = fi75
                            _ = v76
                            _ = v78
                        }
                    }
                    _ = fi70
                }
            case "foo bar":
                fmt.Println(v4)
            default:
                for fi79 := 0; fi79 < 18; fi79++ {
                    if true {
                        for fi80 := 0; fi80 < 16; fi80++ {
                            v2 += v2
                            _ = fi80
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    } else {
                        for fi81 := 0; fi81 < 15; fi81++ {
                            fmt.Println(fi79)
                            _ = fi81
                        }
                        if (!(127 == 54)) {
                            for fi82 := 0; fi82 < 18; fi82++ {
                                v83 := GenS5{F0: -1}
                                if (42 < -61) {
                                    v5 = append(v5, 32768)
                                } else {
                                    switch "foo bar" {
                                    case "":
                                        v84 := GenS5{F0: 176}
                                        _ = v84
                                    case "hello":
                                        switch 4294967295 {
                                        case -287:
                                            if true {
                                                v85 := GenS7{F0: 213, F1: (-9223372036854775808 <= 0)}
                                                var v86 uint16 = 42
                                                if false {
                                                    fmt.Println(v2, v4)
                                                }
                                                _ = v85
                                                _ = v86
                                            }
                                        case 9223372036854775807:
                                            v87 := v1[(fi82 >> ((140) & 63))]
                                            _ = v87
                                        default:
                                            v4 = 1
                                        }
                                    case "foo bar":
                                        v5 = append(v5, v4)
                                    default:
                                        var v88 int = fi82
                                        _ = v88
                                    }
                                    var v89 bool = true
                                    _ = v89
                                }
                                fmt.Println(v4, v4)
                                _ = fi82
                                _ = v83
                            }
                        }
                        var v90 uint16 = 255
                        _ = v90
                    }
                    var v91 int32 = 42
                    _ = fi79
                    _ = v91
                }
            }
            for fi92 := 0; fi92 < 17; fi92++ {
                if true {
                    var v93 string = ""
                    switch v4 {
                    case 137:
                        for fi94 := 0; fi94 < 17; fi94++ {
                            if (!true) {
                                v95 := GenS4{F0: 412}
                                if (v93 >= (v93 + (v93 + v93))) {
                                    fmt.Println(fi94, fi94, v4)
                                }
                                _ = v95
                            } else {
                                var v96 uint64 = 0
                                v2 -= v2
                                _ = v96
                            }
                            switch "" {
                            case "\n\t":
                                var v97 int32 = -631
                                _ = v97
                            }
                            fmt.Println(v2, v93)
                            _ = fi94
                        }
                    }
                    v93 = "hello"
                    _ = v93
                } else {
                    var v98 int8 = -128
                    v99 := GenS8{F0: 57, F1: 16}
                    if true {
                        var v100 int64 = 9223372036854775807
                        delete(v1, (-(v4)))
                        delete(v1, fi92)
                        _ = v100
                    }
                    _ = v98
                    _ = v99
                }
                if true {
                    v101 := GenS8{F0: 42, F1: 0}
                    for fi102 := 0; fi102 < 1; fi102++ {
                        fmt.Println(fi102, v4, v4)
                        var v103 uint8 = genFn0()
                        if (0 <= 100) {
                            var v104 string = (("hello" + "foo bar") + "hello")
                            if (1 != 42) {
                                v104 = (((v104 + "") + v104) + ((v104 + v104) + "hello"))
                            } else {
                                var v105 uint = 2
                                v106 := GenS7{F0: 1000, F1: false}
                                _ = v105
                                _ = v106
                            }
                            _ = v104
                        }
                        _ = fi102
                        _ = v103
                    }
                    switch ("hello" + "foo bar") {
                    case "a":
                        if (fi92 < (fi92 ^ v4)) {
                            for fi107 := 0; fi107 < 5; fi107++ {
                                var v108 uint = 2
                                v4 = (fi107 << ((v108) & 63))
                                v109 := GenS1{F0: 0}
                                _ = fi107
                                _ = v108
                                _ = v109
                            }
                            if true {
                                v5 = append(v5, (fi92 >> ((498) & 63)))
                            } else {
                                v4 += v4
                                for fi110 := 0; fi110 < 16; fi110++ {
                                    if false {
                                        v2 -= (1.5 - v2)
                                        v4 += fi92
                                        var v111 float64 = 0.0
                                        _ = v111
                                    }
                                    var v112 float32 = -2.25
                                    _ = fi110
                                    _ = v112
                                }
                                switch "a" {
                                case "a":
                                    for fi113 := 0; fi113 < 16; fi113++ {
                                        if (2147483648 >= 0) {
                                            v2 += ((((float32(-2.25) - (v2 + -2.25)) / (float32(1e10) + 0.1)) + (v2 / v2)) / v2)
                                            if (39 != 42) {
                                                switch "foo bar" {
                                                case "":
                                                    var v114 int8 = -128
                                                    _ = v114
                                                case "hello":
                                                    switch "foo bar" {
                                                    case "a":
                                                        switch v4 {
                                                        case -580:
                                                            v2 += float32((1e10 * (((float64(v2) * 1.5) + (((((-2.25 - (0.1 - (0.1 / (1.5 + float64((((float32(-2.25) - 0.1) * (v2 * v2)) / v2)))))) / (1e10 - (0.1 + float64((1e10 - float32(0.0)))))) / 1.0) * ((float64(v2) * (((-2.25 / 1e10) + -2.25) / ((1e10 + 1.5) / 1e10))) - (1.5 + ((((1.5 / 1.5) - 3.14159) / 3.14159) - (0.0 / 1e10))))) / 1.0)) - 0.0)))
                                                        case 17:
                                                            if (0 > 1) {
                                                                v4 = 128
                                                            } else {
                                                                var v115 uint = 335
                                                                v115 = (^(v115))
                                                                _ = v115
                                                            }
                                                        default:
                                                            var v116 uint16 = 65535
                                                            _ = v116
                                                        }
                                                    case "":
                                                        v4 -= 737
                                                    case "foo bar":
                                                        if (!(!(!(!true)))) {
                                                            var v117 uint16 = 127
                                                            _ = v117
                                                        } else {
                                                            if (42 != genFn0()) {
                                                                var v118 float32 = 0.0
                                                                _ = len(v5)
                                                                _ = v118
                                                            } else {
                                                                if (9223372036854775807 < 42) {
                                                                    switch (^(v4)) {
                                                                    case 0:
                                                                        var v119 float64 = 0.1
                                                                        _ = v119
                                                                    case -227:
                                                                        v5 = append(v5, fi92)
                                                                    default:
                                                                        var v120 uint16 = 1
                                                                        _ = v120
                                                                    }
                                                                } else {
                                                                    v121 := GenS2{F0: "a", F1: (fi113 ^ fi92)}
                                                                    var v122 uint16 = 2
                                                                    _ = v121
                                                                    _ = v122
                                                                }
                                                                var v123 bool = (-9223372036854775808 <= 562)
                                                                v124 := GenS9{F0: 4294967295, F1: 0, F2: 2}
                                                                _ = v123
                                                                _ = v124
                                                            }
                                                        }
                                                    default:
                                                        var v125 int8 = -1
                                                        _ = v125
                                                    }
                                                }
                                            } else {
                                                v2 -= v2
                                                var v126 uint16 = 65535
                                                _ = v126
                                            }
                                            if (128 == 0) {
                                                v127 := GenS9{F0: 42, F1: 42, F2: 0}
                                                v4 = 630
                                                fmt.Println(fi113, v4, fi92)
                                                _ = v127
                                            } else {
                                                var v128 int32 = 2147483647
                                                _ = v128
                                            }
                                        }
                                        _ = fi113
                                    }
                                }
                            }
                        } else {
                            v2 -= 1e10
                        }
                    case "":
                        switch (fi92 << ((1) & 63)) {
                        case 32767:
                            switch "hello" {
                            case "hello":
                                v2 += v2
                            case "":
                                var v129 float64 = 1.5
                                _ = v129
                            }
                        case -605:
                            fmt.Println(fi92)
                        }
                    case "\n\t":
                        if (!((false && false) || (0 < -1))) {
                            switch ((("" + ("hello" + "hello")) + "foo bar") + "\n\t") {
                            case "a":
                                var v130 string = ""
                                _ = v130
                            case "foo bar":
                                _ = len(v5)
                            }
                        }
                    default:
                        var v131 int16 = 1
                        _ = v131
                    }
                    _ = v101
                } else {
                    fmt.Println(v4, fi92)
                    fmt.Println(v2)
                }
                _ = fi92
            }
        }
        var v132 int8 = -128
        _ = v132
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi133 := 0; fi133 < 15; fi133++ {
        if false {
            for fi134 := 0; fi134 < 7; fi134++ {
                v135 := GenS10{F0: 255, F1: 0}
                if (127 == 162) {
                    fmt.Println(v2, fi133)
                    switch (fi133 | fi134) {
                    case 57:
                        v2 += v2
                    case 4294967295:
                        var v136 uint16 = 32767
                        _ = v136
                    }
                }
                _ = fi134
                _ = v135
            }
            for fi137 := 0; fi137 < 10; fi137++ {
                v4 += (^(v4))
                var v138 int16 = 0
                v139 := GenS11{F0: (fi137 > v4), F1: (v138 % ((802) | 1))}
                fmt.Println(v138)
                switch "foo bar" {
                case "foo bar":
                    v140 := GenS2{F0: "a", F1: v4}
                    v4 -= -2147483648
                    _ = v140
                }
                var v141 float64 = 1e10
                var v142 int = fi133
                _ = fi137
                _ = v138
                _ = v139
                _ = v141
                _ = v142
            }
            fmt.Println(v2)
        }
        var v143 bool = true
        for fi144 := 0; fi144 < 8; fi144++ {
            v143 = (2 == 42)
            _ = fi144
        }
        _ = fi133
        _ = v143
    }
    if true {
        for fi145 := 0; fi145 < 16; fi145++ {
            var v146 float64 = float64(v2)
            v147 := v3[-9223372036854775808]
            v2 = -2.25
            _ = fi145
            _ = v146
            _ = v147
        }
        fmt.Println(v4, v2)
    }
    switch "" {
    case "a":
        v2 = v2
    }
    v2 = ((v2 - v2) + float32((1e10 * 0.1)))
    for fi148 := 0; fi148 < 13; fi148++ {
        v4 += (fi148 ^ v4)
        for fi149 := 0; fi149 < 11; fi149++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v2 = v2
            var v150 int16 = -5
            v151 := GenS12{F0: 0}
            _ = fi149
            _ = v150
            _ = v151
        }
        _ = fi148
    }
    v152 := GenS3{F0: "hello", F1: v4, F2: 42}
    v153 := GenS5{F0: 478}
    _ = len(v5)
    v2 += 1.5
    var v154 string = ""
    if true {
        var v155 uint64 = 1
        if (!((float64(v2) + ((1e10 / (0.1 * (((float64(v2) / (1.5 + 3.14159)) + ((0.1 + 1e10) - float64(v2))) * ((1e10 + 0.1) - 0.1)))) * 1.5)) < (1e10 - -2.25))) {
            v156 := GenS13{F0: (v4 + (v4 - (v4 * (v4 ^ -528))))}
            switch v154 {
            case "":
                _ = len(v1)
            case "\n\t":
                v155 -= (v155 + (v155 >> ((42) & 63)))
            case "foo bar":
                fmt.Println(v2, v2, v154)
            }
            var v157 uint16 = 2
            _ = v156
            _ = v157
        }
        _ = v155
    }
    switch v4 {
    case 2147483647:
        var v158 uint = 42
        _ = v158
    case 1:
        fmt.Println(v2, v154, v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v154:", v154)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v152
    _ = v153
    _ = v154
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
