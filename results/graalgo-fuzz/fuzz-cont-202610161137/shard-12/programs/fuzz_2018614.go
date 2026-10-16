package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int8; F1 bool; F2 bool }
type GenS5 struct { F0 bool; F1 uint8 }
type GenS6 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint; F1 uint8 }
type GenS9 struct { F0 int; F1 uint16; F2 uint8 }
type GenS10 struct { F0 int; F1 uint32 }
type GenS11 struct { F0 uint16; F1 int32; F2 string }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 uint8; F1 string }
type GenS14 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS15 struct { F0 uint16; F1 int16 }
type GenS16 struct { F0 int; F1 int32 }
type GenS17 struct { F0 bool; F1 int }
type GenS18 struct { F0 uint8; F1 int8 }
type GenS19 struct { F0 uint }
type GenS20 struct { F0 int }
type GenS21 struct { F0 int16 }
type GenS22 struct { F0 uint64; F1 bool }
type GenS23 struct { F0 int64; F1 int64 }
type GenS24 struct { F0 uint; F1 int; F2 int16 }
type GenS25 struct { F0 int8; F1 uint }
func genFn0() uint8 {
    if (0.1 < (3.14159 / -2.25)) {
        var v5 int64 = -26
        v5 = 42
        _ = v5
    }
    if (((!false) && true) || (1 != 752)) {
        v6 := GenS1{F0: 2147483647, F1: (100 != -101), F2: 0}
        var v7 float32 = (float32(1e10) + 1e10)
        _ = v6
        _ = v7
    } else {
        var v8 int32 = 2147483647
        _ = v8
    }
    switch ((("a" + "\n\t") + "") + ("" + "")) {
    case "":
        var v9 string = ("\n\t" + (("hello" + "") + ("foo bar" + "hello")))
        _ = v9
    case "\n\t":
        var v10 int8 = 1
        _ = v10
    }
    var v11 int64 = -9223372036854775808
    v11 = (v11 % (((v11 ^ -1)) | 1))
    _ = v11
    return 37
}

func main() {
    var v1 uint32 = 582
    v2 := []int{0, 808, 32767}
    var v3 uint32 = v1
    var v4 uint = 4294967295
    if false {
        var v12 uint32 = v3
        for fi13 := 0; fi13 < 18; fi13++ {
            if false {
                v14 := GenS1{F0: (v12 & v1), F1: true, F2: -32768}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if true {
                    var v15 bool = false
                    var v16 uint16 = 0
                    v17 := GenS1{F0: (^(v3)), F1: v15, F2: -32768}
                    _ = v15
                    _ = v16
                    _ = v17
                } else {
                    var v18 string = ("a" + (((("foo bar" + "a") + "") + "hello") + "hello"))
                    _ = v18
                }
                _ = v14
            } else {
                if true {
                    if true {
                        if ((-2.25 * (-2.25 - ((3.14159 * 0.1) / 3.14159))) > (3.14159 * 1e10)) {
                            v19 := GenS1{F0: (v1 % ((v1) | 1)), F1: false, F2: 0}
                            _ = v19
                        }
                        v20 := GenS2{F0: 32767, F1: 312}
                        v1 = (v1 % ((0) | 1))
                        _ = v20
                    }
                    v12 += v3
                }
            }
            _ = fi13
        }
        _ = v12
    }
    v2 = append(v2, 69)
    for fi21 := 0; fi21 < 8; fi21++ {
        v22 := GenS3{F0: 964, F1: genFn0()}
        for fi23 := 0; fi23 < 19; fi23++ {
            v24 := GenS3{F0: (fi21 * (fi21 + (^(fi21)))), F1: genFn0()}
            if len(v2) > 0 {
                v2[0] = fi21
            }
            v25 := GenS1{F0: (v1 * 1), F1: false, F2: 0}
            var v26 float64 = (-2.25 / 0.1)
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
        _ = fi21
        _ = v22
    }
    v1 += 954
    switch "foo bar" {
    case "":
        switch "" {
        case "foo bar":
            v27 := GenS4{F0: 127, F1: (!(1000 < -32768)), F2: true}
            _ = v27
        }
    default:
        v4 = v4
    }
    if (!(!(-390 != 127))) {
        if (!((v4 << ((0) & 63)) != v4)) {
            fmt.Println(v3, v4)
            fmt.Println(v3, v4, v1)
            v3 = v1
        }
    } else {
        if false {
            var v28 int = 1
            if (42 != -74) {
                var v29 int64 = -264
                switch (v28 & -856) {
                case 9223372036854775807:
                    for fi30 := 0; fi30 < 6; fi30++ {
                        var v31 uint16 = 255
                        _ = fi30
                        _ = v31
                    }
                case 1:
                    fmt.Println(v3)
                }
                for fi32 := 0; fi32 < 9; fi32++ {
                    switch 841 {
                    case 128:
                        switch "\n\t" {
                        case "\n\t":
                            var v33 int64 = 409
                            _ = v33
                        case "":
                            if ("\n\t" != "hello") {
                                if (95 != -118) {
                                    _ = len(v2)
                                    if (4294967295 < v1) {
                                        fmt.Println(v28, v29, v4)
                                        fmt.Println(fi32, v4)
                                        if (false && true) {
                                            var v34 uint32 = (^(v3))
                                            v34 = v3
                                            _ = v34
                                        } else {
                                            var v35 int16 = 815
                                            v1 = (v3 / (((v1 ^ 42)) | 1))
                                            var v36 int16 = (-(v35))
                                            _ = v35
                                            _ = v36
                                        }
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = -971
                                    }
                                } else {
                                    var v37 int16 = -546
                                    _ = v37
                                }
                                v38 := GenS1{F0: (v3 ^ (v3 % (((v1 >> ((v4) & 63))) | 1))), F1: (v29 != v29), F2: 0}
                                _ = v38
                            }
                        }
                    }
                    _ = fi32
                }
                _ = v29
            } else {
                var v39 uint64 = 18446744073709551615
                v3 += 2147483648
                if false {
                    v3 -= v3
                }
                _ = v39
            }
            var v40 float64 = 3.14159
            _ = v28
            _ = v40
        } else {
            var v41 int = 42
            _ = len(v2)
            _ = v41
        }
        if (true || true) {
            var v42 float32 = 1e10
            var v43 uint64 = 569
            _ = v42
            _ = v43
        } else {
            if true {
                fmt.Println(v3, v3)
            } else {
                switch "foo bar" {
                case "foo bar":
                    if true {
                        fmt.Println(v1, v4, v1)
                        for fi44 := 0; fi44 < 1; fi44++ {
                            switch fi44 {
                            case -2147483648:
                                v45 := GenS2{F0: 65535, F1: 723}
                                _ = v45
                            case 4294967295:
                                v46 := GenS2{F0: 934, F1: 255}
                                _ = v46
                            case 0:
                                if (false || (((((1e10 + 3.14159) + float32(1.5)) >= ((0.0 * -2.25) + (3.14159 + 1e10))) && (!true)) || true)) {
                                    if len(v2) > 0 {
                                        v2[0] = fi44
                                    }
                                    v4 = v4
                                } else {
                                    fmt.Println(v1, v3)
                                }
                            default:
                                var v47 uint16 = 255
                                _ = v47
                            }
                            fmt.Println(v3, v1, fi44)
                            _ = fi44
                        }
                        for fi48 := 0; fi48 < 12; fi48++ {
                            fmt.Println(v1, fi48, v4)
                            _ = fi48
                        }
                    } else {
                        v49 := GenS1{F0: (v3 % (((v3 * v1)) | 1)), F1: false, F2: -1}
                        switch ("hello" + "hello") {
                        case "":
                            v50 := GenS3{F0: 42, F1: 2}
                            _ = v50
                        default:
                            for fi51 := 0; fi51 < 19; fi51++ {
                                if len(v2) > 0 {
                                    v2[0] = (fi51 + (fi51 << ((v4) & 63)))
                                }
                                if (!true) {
                                    for fi52 := 0; fi52 < 18; fi52++ {
                                        var v53 int8 = 16
                                        _ = fi52
                                        _ = v53
                                    }
                                    var v54 bool = (-163 > -1)
                                    v55 := GenS5{F0: (v54 || true), F1: 1}
                                    _ = v54
                                    _ = v55
                                } else {
                                    var v56 uint64 = 4294967295
                                    if false {
                                        if len(v2) > 0 {
                                            v2[0] = fi51
                                        }
                                        v56 = v56
                                    }
                                    v57 := GenS2{F0: 925, F1: 18446744073709551615}
                                    _ = v56
                                    _ = v57
                                }
                                _ = fi51
                            }
                        }
                        var v58 float32 = (((-2.25 / -2.25) + 3.14159) + 3.14159)
                        _ = v49
                        _ = v58
                    }
                case "":
                    if (18446744073709551615 > 745) {
                        switch "hello" {
                        case "foo bar":
                            if true {
                                v4 -= v4
                                fmt.Println(v4, v3)
                            }
                        case "\n\t":
                            if true {
                                switch 42 {
                                case -849:
                                    var v59 uint8 = 1
                                    _ = v59
                                case 333:
                                    if (!(0.0 > (((3.14159 * (1e10 * 3.14159)) + 0.1) / float32(1e10)))) {
                                        if len(v2) > 0 {
                                            v2[0] = 461
                                        }
                                        for fi60 := 0; fi60 < 1; fi60++ {
                                            var v61 float32 = 3.14159
                                            _ = fi60
                                            _ = v61
                                        }
                                    } else {
                                        var v62 float64 = float64(1.5)
                                        for fi63 := 0; fi63 < 6; fi63++ {
                                            for fi64 := 0; fi64 < 15; fi64++ {
                                                if (v4 > v4) {
                                                    v3 = 1
                                                    v2 = append(v2, -1)
                                                    var v65 uint32 = uint32(v4)
                                                    _ = v65
                                                }
                                                v66 := GenS6{F0: (^(v3)), F1: uint64(v4), F2: 127}
                                                _ = fi64
                                                _ = v66
                                            }
                                            _ = fi63
                                        }
                                        v4 += (v4 | 4294967295)
                                        _ = v62
                                    }
                                case -994:
                                    var v67 int16 = 0
                                    _ = v67
                                default:
                                    fmt.Println(v3)
                                }
                            }
                        default:
                            v68 := GenS7{F0: 127}
                            _ = v68
                        }
                    }
                case "a":
                    for fi69 := 0; fi69 < 8; fi69++ {
                        for fi70 := 0; fi70 < 8; fi70++ {
                            if (2 <= 65535) {
                                fmt.Println(fi69, v1, v4)
                                var v71 int16 = -32768
                                _ = v71
                            } else {
                                v2 = append(v2, (fi69 ^ fi69))
                                switch -1 {
                                case 128:
                                    v1 += v3
                                case -1:
                                    fmt.Println(v3, v4, v4)
                                default:
                                    _ = len(v2)
                                }
                                v1 = (v3 >> ((v4) & 63))
                            }
                            v72 := GenS4{F0: 1, F1: true, F2: ((-1 <= -40) && (128 == 50))}
                            _ = fi70
                            _ = v72
                        }
                        _ = fi69
                    }
                default:
                    for fi73 := 0; fi73 < 12; fi73++ {
                        for fi74 := 0; fi74 < 20; fi74++ {
                            if (!true) {
                                var v75 int = (fi74 - fi73)
                                var v76 uint = 437
                                _ = v75
                                _ = v76
                            }
                            if len(v2) > 0 {
                                v2[0] = int(v4)
                            }
                            _ = fi74
                        }
                        _ = fi73
                    }
                }
                if false {
                    var v77 uint16 = 550
                    _ = v77
                } else {
                    v78 := GenS8{F0: (v4 << ((255) & 63)), F1: 2}
                    var v79 float32 = -2.25
                    _ = v78
                    _ = v79
                }
            }
        }
    }
    var v80 int64 = 9223372036854775807
    v80 += -9223372036854775808
    for fi81 := 0; fi81 < 13; fi81++ {
        var v82 uint32 = 338
        switch 127 {
        case 0:
            if (2 < 65535) {
                v2 = append(v2, fi81)
                v83 := GenS8{F0: 255, F1: 0}
                _ = v83
            } else {
                for fi84 := 0; fi84 < 16; fi84++ {
                    fmt.Println(fi81, v3)
                    v1 = v1
                    fmt.Println(v3, fi84)
                    _ = fi84
                }
            }
        case 32767:
            switch "foo bar" {
            case "":
                var v85 int8 = -9
                _ = v85
            }
        case -1:
            v82 += v3
        }
        _ = fi81
        _ = v82
    }
    if (false && true) {
        if false {
            for fi86 := 0; fi86 < 12; fi86++ {
                for fi87 := 0; fi87 < 8; fi87++ {
                    if true {
                        if false {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v3 += (v3 << ((65535) & 63))
                        } else {
                            v88 := GenS9{F0: fi86, F1: 1, F2: 127}
                            if (!true) {
                                v89 := GenS10{F0: (fi87 * fi87), F1: 2}
                                _ = v89
                            }
                            var v90 int32 = 42
                            _ = v88
                            _ = v90
                        }
                        switch (fi87 / (((fi87 << (((v4 | 24)) & 63))) | 1)) {
                        case 9223372036854775807:
                            v91 := GenS2{F0: 68, F1: v4}
                            _ = v91
                        case -92:
                            v92 := GenS3{F0: fi87, F1: 127}
                            _ = v92
                        case 424:
                            v93 := GenS11{F0: 1, F1: 2147483647, F2: "\n\t"}
                            _ = v93
                        }
                        _ = len(v2)
                    } else {
                        v94 := GenS1{F0: v1, F1: ((!(int64(v1) > 389)) || (!false)), F2: -218}
                        _ = v94
                    }
                    _ = fi87
                }
                if len(v2) > 0 {
                    v2[0] = 32767
                }
                if ((-460 >= 1000) || (2 > 32767)) {
                    var v95 uint16 = 255
                    v95 = v95
                    if (v4 == v4) {
                        if ((v95 / ((v95) | 1)) != v95) {
                            for fi96 := 0; fi96 < 12; fi96++ {
                                switch ("" + "a") {
                                case "a":
                                    v97 := GenS2{F0: 42, F1: v4}
                                    _ = v97
                                case "hello":
                                    var v98 bool = (!(uint32(v4) < v1))
                                    _ = v98
                                case "foo bar":
                                    if (true || true) {
                                        var v99 int8 = 100
                                        _ = v99
                                    } else {
                                        var v100 float64 = (((3.14159 / (0.1 / -2.25)) + 3.14159) + 1.5)
                                        _ = v100
                                    }
                                }
                                _ = fi96
                            }
                            v80 = int64(fi86)
                            v80 -= (v80 ^ (v80 % ((v80) | 1)))
                        } else {
                            if ((0 <= 1) && (((1e10 == -2.25) && (!(42 <= 32767))) || (-105 < int16(v95)))) {
                                v4 = v4
                            }
                            var v101 float32 = 3.14159
                            _ = v101
                        }
                        fmt.Println(v3, v95)
                        if true {
                            switch 312 {
                            case 1:
                                v1 = 2147483647
                            case 4294967295:
                                var v102 uint = (v4 - v4)
                                _ = v102
                            case -559:
                                var v103 uint = 0
                                _ = v103
                            default:
                                switch (fi86 << ((v4) & 63)) {
                                case -911:
                                    fmt.Println(v1, v4, v1)
                                default:
                                    for fi104 := 0; fi104 < 16; fi104++ {
                                        fmt.Println(v3)
                                        _ = fi104
                                    }
                                }
                            }
                            fmt.Println(v80, v3)
                            fmt.Println(v80, fi86, fi86)
                        } else {
                            for fi105 := 0; fi105 < 20; fi105++ {
                                v3 = (v3 % ((v3) | 1))
                                v4 -= (^(v4))
                                v4 = v4
                                _ = fi105
                            }
                            for fi106 := 0; fi106 < 8; fi106++ {
                                v107 := GenS3{F0: (^(fi86)), F1: 2}
                                var v108 uint32 = 865
                                var v109 int16 = -716
                                if len(v2) > 0 {
                                    v2[0] = (fi106 >> ((v4) & 63))
                                }
                                _ = fi106
                                _ = v107
                                _ = v108
                                _ = v109
                            }
                            var v110 uint16 = (v95 ^ (v95 & (^(v95))))
                            _ = v110
                        }
                    } else {
                        if false {
                            var v111 uint32 = uint32(v95)
                            _ = v111
                        } else {
                            if (((v95 ^ v95) != 764) && (!true)) {
                                for fi112 := 0; fi112 < 1; fi112++ {
                                    if (!false) {
                                        v113 := GenS8{F0: v4, F1: 1}
                                        if (false && ((388 > -32768) || false)) {
                                            if false {
                                                _ = len(v2)
                                                if (false || (!(!(1 < -1)))) {
                                                    v114 := GenS10{F0: 32768, F1: uint32(v4)}
                                                    fmt.Println(v1, v80, v80)
                                                    var v115 int8 = -128
                                                    v116 := GenS12{F0: v115}
                                                    if true {
                                                        fmt.Println(v95, v115, fi112)
                                                        switch ("\n\t" + "") {
                                                        case "foo bar":
                                                            if ((v1 | (v3 / ((65535) | 1))) != v3) {
                                                                switch (-(fi112)) {
                                                                case 32767:
                                                                    var v117 uint8 = 255
                                                                    _ = v117
                                                                case 127:
                                                                    if len(v2) > 0 {
                                                                        v2[0] = 32768
                                                                    }
                                                                case 444:
                                                                    v95 += 2
                                                                default:
                                                                    v118 := GenS5{F0: (fi86 >= fi112), F1: 128}
                                                                    _ = v118
                                                                }
                                                                if (229 >= -32768) {
                                                                    v1 -= v3
                                                                }
                                                                switch (fi112 << ((255) & 63)) {
                                                                case 9223372036854775807:
                                                                    switch 66 {
                                                                    case 656:
                                                                        if (float32((1e10 / 0.1)) != float32(-2.25)) {
                                                                            if (int64(v115) > (v80 ^ -122)) {
                                                                                if false {
                                                                                    var v119 uint = 1
                                                                                    if true {
                                                                                        v120 := GenS13{F0: 33, F1: "a"}
                                                                                        _ = v120
                                                                                    }
                                                                                    _ = v119
                                                                                } else {
                                                                                    fmt.Println(fi112, v1)
                                                                                    v121 := GenS14{F0: 42, F1: 889, F2: 2}
                                                                                    _ = v121
                                                                                }
                                                                                v122 := GenS5{F0: (!(false || ((v115 % ((-1) | 1)) != v115))), F1: 45}
                                                                                v95 = 297
                                                                                _ = v122
                                                                            } else {
                                                                                v80 += (v80 + (v80 - v80))
                                                                                var v123 float64 = 1e10
                                                                                _ = v123
                                                                            }
                                                                            var v124 uint32 = v3
                                                                            fmt.Println(v95, fi112)
                                                                            _ = v124
                                                                        }
                                                                    default:
                                                                        v3 = v3
                                                                    }
                                                                default:
                                                                    var v125 uint32 = 336
                                                                    _ = v125
                                                                }
                                                            } else {
                                                                if false {
                                                                    if (true || true) {
                                                                        v80 = v80
                                                                    }
                                                                    var v126 int32 = 1
                                                                    var v127 int8 = v115
                                                                    _ = v126
                                                                    _ = v127
                                                                } else {
                                                                    v95 -= v95
                                                                    if len(v2) > 0 {
                                                                        v2[0] = -309
                                                                    }
                                                                    v128 := GenS15{F0: v95, F1: -32768}
                                                                    _ = v128
                                                                }
                                                                v80 -= v80
                                                                v129 := GenS7{F0: int8(v80)}
                                                                _ = v129
                                                            }
                                                        case "":
                                                            v80 = -1
                                                        case "hello":
                                                            v4 = 675
                                                        default:
                                                            if true {
                                                                fmt.Println(v1)
                                                                v95 = 0
                                                                fmt.Println(fi86)
                                                            }
                                                        }
                                                        fmt.Println(v115, v1, v115)
                                                        var v130 int32 = 652
                                                        _ = v130
                                                    } else {
                                                        var v131 uint = 21
                                                        v80 += -98
                                                        if (1 == (-(fi86))) {
                                                            v4 = v131
                                                        } else {
                                                            var v132 bool = false
                                                            _ = v132
                                                        }
                                                        _ = v131
                                                    }
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    _ = v114
                                                    _ = v115
                                                    _ = v116
                                                } else {
                                                    var v133 string = "hello"
                                                    var v134 int16 = -32768
                                                    _ = v133
                                                    _ = v134
                                                }
                                            } else {
                                                var v135 bool = true
                                                _ = v135
                                            }
                                            v95 += (v95 << (((v4 * (^(v4)))) & 63))
                                            fmt.Println(v95, v95)
                                        } else {
                                            v80 += int64(fi112)
                                        }
                                        switch (fi112 >> (((v4 | (v4 & v4))) & 63)) {
                                        case -2147483648:
                                            v1 -= (^(v1))
                                        default:
                                            if len(v2) > 0 {
                                                v2[0] = -867
                                            }
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        v136 := GenS7{F0: 1}
                                        switch "" {
                                        case "":
                                            if (false || (false || (127 < 42))) {
                                                if true {
                                                    var v137 uint32 = v1
                                                    if false {
                                                        var v138 int8 = -102
                                                        v137 = (v1 | (v137 * v1))
                                                        v139 := GenS9{F0: (^(fi86)), F1: 32767, F2: 232}
                                                        _ = v138
                                                        _ = v139
                                                    }
                                                    var v140 uint = (v4 | (v4 ^ (v4 ^ (v4 | (v4 * 2)))))
                                                    _ = v137
                                                    _ = v140
                                                }
                                            }
                                        case "hello":
                                            v141 := GenS16{F0: (fi86 & (fi86 << ((522) & 63))), F1: 1}
                                            _ = v141
                                        default:
                                            v142 := GenS17{F0: true, F1: (-(fi86))}
                                            _ = v142
                                        }
                                        if ((((0.1 * 3.14159) + (0.1 + 0.0)) * (1e10 * float32((float64((3.14159 - float32((-2.25 * 1.5)))) - (0.1 + 1e10))))) > float32(-2.25)) {
                                            if true {
                                                var v143 int8 = 42
                                                _ = v143
                                            }
                                        } else {
                                            if (-828 >= -385) {
                                                if len(v2) > 0 {
                                                    v2[0] = -489
                                                }
                                            }
                                            v144 := GenS18{F0: 28, F1: 121}
                                            if false {
                                                switch (^(fi112)) {
                                                case 9223372036854775807:
                                                    if true {
                                                        v3 = 2147483647
                                                        v145 := GenS8{F0: v4, F1: 171}
                                                        v80 = (v80 & v80)
                                                        _ = v145
                                                    } else {
                                                        var v146 int = -698
                                                        _ = v146
                                                    }
                                                default:
                                                    var v147 float64 = 1e10
                                                    _ = v147
                                                }
                                                var v148 int8 = 1
                                                _ = v148
                                            }
                                            _ = v144
                                        }
                                        _ = v113
                                        _ = v136
                                    } else {
                                        switch "a" {
                                        case "":
                                            switch "" {
                                            case "foo bar":
                                                if true {
                                                    v4 -= (v4 << (((v4 % ((v4) | 1))) & 63))
                                                    fmt.Println(v3, fi112, v3)
                                                    v149 := GenS3{F0: 2147483647, F1: 255}
                                                    _ = v149
                                                } else {
                                                    var v150 float32 = 1.5
                                                    _ = v150
                                                }
                                            case "a":
                                                if (!false) {
                                                    switch "hello" {
                                                    case "a":
                                                        if len(v2) > 0 {
                                                            v2[0] = int(v4)
                                                        }
                                                    case "hello":
                                                        v3 = v3
                                                    default:
                                                        switch "a" {
                                                        case "a":
                                                            var v151 bool = (true || (false || (false || false)))
                                                            _ = v151
                                                        case "\n\t":
                                                            var v152 float32 = (0.1 - float32(1.5))
                                                            _ = v152
                                                        case "hello":
                                                            var v153 int16 = 0
                                                            _ = v153
                                                        default:
                                                            var v154 float64 = 0.1
                                                            _ = v154
                                                        }
                                                    }
                                                    if (-2147483648 >= -1) {
                                                        fmt.Println(v1)
                                                        v1 = 0
                                                    } else {
                                                        switch fi86 {
                                                        case -588:
                                                            switch (fi112 - fi112) {
                                                            case -9223372036854775808:
                                                                v2 = append(v2, (fi112 & fi86))
                                                            case 2147483647:
                                                                fmt.Println(fi86, v95)
                                                            case 32768:
                                                                v95 = (v95 ^ (^(v95)))
                                                            }
                                                        default:
                                                            if (!((("a" + (("a" + "foo bar") + "\n\t")) == ("hello" + "\n\t")) || false)) {
                                                                fmt.Println(v95)
                                                                v1 = uint32(v95)
                                                                v155 := GenS14{F0: (^(v95)), F1: 4294967295, F2: 139}
                                                                _ = v155
                                                            }
                                                        }
                                                        var v156 float64 = 1.5
                                                        var v157 float32 = (float32(1.5) * 1e10)
                                                        _ = v156
                                                        _ = v157
                                                    }
                                                    if (1e10 > 3.14159) {
                                                        var v158 int32 = 0
                                                        if (32767 != -188) {
                                                            fmt.Println(v1)
                                                            v3 = v1
                                                            v159 := GenS19{F0: v4}
                                                            _ = v159
                                                        } else {
                                                            v95 = v95
                                                        }
                                                        _ = v158
                                                    } else {
                                                        v160 := GenS10{F0: (fi86 - 176), F1: 2147483647}
                                                        v80 -= (v80 | -1)
                                                        var v161 int32 = -2147483648
                                                        fmt.Println(v95, v161, v80)
                                                        _ = v160
                                                        _ = v161
                                                    }
                                                }
                                            case "hello":
                                                v162 := GenS4{F0: 42, F1: true, F2: (uint8(v4) <= 2)}
                                                _ = v162
                                            default:
                                                v3 -= 4294967295
                                            }
                                        case "foo bar":
                                            var v163 float64 = ((float64((((3.14159 / 1.0) * 1e10) + 0.1)) * -2.25) * 0.0)
                                            _ = v163
                                        case "a":
                                            switch (("a" + "") + "\n\t") {
                                            case "a":
                                                switch -9223372036854775808 {
                                                case 127:
                                                    if false {
                                                        var v164 uint32 = (v1 - 2147483648)
                                                        v164 = 2
                                                        var v165 int64 = 1
                                                        _ = v164
                                                        _ = v165
                                                    } else {
                                                        v1 = 42
                                                        v166 := GenS20{F0: fi112}
                                                        v167 := GenS12{F0: 116}
                                                        if false {
                                                            v4 += (^(v4))
                                                            v168 := GenS21{F0: -1}
                                                            var v169 uint8 = 127
                                                            _ = v168
                                                            _ = v169
                                                        }
                                                        _ = v166
                                                        _ = v167
                                                    }
                                                case 4294967295:
                                                    var v170 int64 = (v80 >> (((v4 ^ v4)) & 63))
                                                    _ = v170
                                                default:
                                                    if (42 != -95) {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    } else {
                                                        if (-1 < 1) {
                                                            var v171 float32 = (((1.5 * 1e10) * (1e10 - ((0.1 / (float32(1.5) * (-2.25 - 1e10))) + float32(float64(1.5))))) * 0.1)
                                                            v172 := GenS22{F0: 199, F1: false}
                                                            v173 := GenS23{F0: (v80 >> ((v4) & 63)), F1: 9223372036854775807}
                                                            if ((fi112 - (fi112 / (((-(fi112))) | 1))) == (fi86 | (fi112 ^ (fi86 + (^(fi86)))))) {
                                                                fmt.Println(fi112, fi112, v80)
                                                                if false {
                                                                    var v174 uint64 = 2
                                                                    v175 := GenS6{F0: (v1 << (((v4 | (v4 ^ (^(v4))))) & 63)), F1: v174, F2: -88}
                                                                    _ = len(v2)
                                                                    _ = v174
                                                                    _ = v175
                                                                } else {
                                                                    if true {
                                                                        _ = len(v2)
                                                                    } else {
                                                                        var v176 uint16 = 65535
                                                                        v95 = 342
                                                                        _ = v176
                                                                    }
                                                                    v4 += (v4 % (((v4 * (v4 | v4))) | 1))
                                                                    var v177 bool = true
                                                                    _ = v177
                                                                }
                                                            } else {
                                                                v171 = -2.25
                                                                v178 := GenS19{F0: v4}
                                                                _ = v178
                                                            }
                                                            _ = v171
                                                            _ = v172
                                                            _ = v173
                                                        } else {
                                                            if true {
                                                                v95 -= v95
                                                                v179 := GenS4{F0: -128, F1: true, F2: (!false)}
                                                                v180 := GenS10{F0: int(v95), F1: v3}
                                                                _ = v179
                                                                _ = v180
                                                            }
                                                            if true {
                                                                v80 = (v80 / ((v80) | 1))
                                                            }
                                                            switch fi86 {
                                                            case -1:
                                                                var v181 float32 = (1e10 * 0.0)
                                                                _ = v181
                                                            }
                                                            var v182 int = (fi112 + fi86)
                                                            _ = v182
                                                        }
                                                        fmt.Println(v1, fi86, v4)
                                                    }
                                                }
                                            default:
                                                v183 := GenS23{F0: 9223372036854775807, F1: (v80 | (v80 + v80))}
                                                _ = v183
                                            }
                                        }
                                        v184 := GenS21{F0: -1}
                                        var v185 int8 = 127
                                        _ = v184
                                        _ = v185
                                    }
                                    var v186 uint16 = (v95 << (((v4 >> ((510) & 63))) & 63))
                                    var v187 uint8 = 60
                                    _ = fi112
                                    _ = v186
                                    _ = v187
                                }
                                if false {
                                    for fi188 := 0; fi188 < 14; fi188++ {
                                        if (true || (1.5 < 1.5)) {
                                            var v189 bool = (1 > 0)
                                            if len(v2) > 0 {
                                                v2[0] = (fi86 - 827)
                                            }
                                            _ = v189
                                        }
                                        if (v95 > v95) {
                                            fmt.Println(v4, v95, fi86)
                                            v3 = (v1 | uint32(v4))
                                        }
                                        v3 = v1
                                        _ = fi188
                                    }
                                    for fi190 := 0; fi190 < 14; fi190++ {
                                        fmt.Println(v1)
                                        _ = fi190
                                    }
                                    v191 := GenS16{F0: (fi86 >> ((v4) & 63)), F1: -245}
                                    _ = v191
                                }
                                if (0.0 <= 0.0) {
                                    if ((v4 * v4) == v4) {
                                        v3 = v3
                                    } else {
                                        fmt.Println(v95, v4)
                                        v192 := GenS23{F0: v80, F1: (v80 >> ((v4) & 63))}
                                        _ = v192
                                    }
                                    v193 := GenS3{F0: (fi86 + fi86), F1: 153}
                                    var v194 uint32 = (v1 + (^(v1)))
                                    _ = v193
                                    _ = v194
                                }
                            } else {
                                v95 -= v95
                                for fi195 := 0; fi195 < 16; fi195++ {
                                    var v196 int64 = (v80 + (^(v80)))
                                    _ = fi195
                                    _ = v196
                                }
                                switch "foo bar" {
                                case "foo bar":
                                    v197 := GenS24{F0: v4, F1: fi86, F2: 1000}
                                    _ = v197
                                case "a":
                                    v198 := GenS15{F0: (v95 / (((v95 % ((v95) | 1))) | 1)), F1: 32767}
                                    _ = v198
                                case "hello":
                                    if len(v2) > 0 {
                                        v2[0] = 2147483647
                                    }
                                }
                            }
                        }
                        var v199 uint32 = v3
                        v95 = (v95 / ((v95) | 1))
                        _ = v199
                    }
                    _ = v95
                }
                _ = fi86
            }
        }
    } else {
        var v200 uint64 = 18446744073709551615
        v80 = v80
        for fi201 := 0; fi201 < 3; fi201++ {
            v202 := GenS11{F0: 2, F1: -515, F2: ("hello" + ((("" + "hello") + ("foo bar" + "a")) + "foo bar"))}
            _ = fi201
            _ = v202
        }
        _ = v200
    }
    if false {
        v2 = append(v2, 916)
        fmt.Println(v80, v1)
        v80 += (v80 % ((v80) | 1))
    } else {
        switch "\n\t" {
        case "\n\t":
            var v203 int8 = -1
            _ = v203
        }
    }
    var v204 int32 = 252
    for fi205 := 0; fi205 < 6; fi205++ {
        for fi206 := 0; fi206 < 16; fi206++ {
            switch (fi206 & (fi206 & fi205)) {
            case -9223372036854775808:
                switch fi205 {
                case 1:
                    v207 := GenS25{F0: -128, F1: 2}
                    _ = v207
                case 42:
                    v80 += v80
                    v2 = append(v2, -350)
                    v4 += v4
                default:
                    var v208 string = "a"
                    _ = v208
                }
            default:
                var v209 int32 = (v204 / (((v204 + -2147483648)) | 1))
                _ = v209
            }
            v1 -= v3
            _ = fi206
        }
        _ = fi205
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v1, v1, v80)
    switch (("foo bar" + "foo bar") + "") {
    case "hello":
        var v210 float32 = 0.1
        _ = v210
    case "\n\t":
        fmt.Println(v204, v4, v3)
    case "foo bar":
        for fi211 := 0; fi211 < 20; fi211++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v212 := GenS20{F0: (-(fi211))}
            _ = fi211
            _ = v212
        }
    default:
        v204 = v204
    }
    switch "foo bar" {
    case "a":
        v2 = append(v2, -1)
    case "hello":
        v213 := GenS11{F0: 32767, F1: v204, F2: "foo bar"}
        _ = v213
    default:
        if len(v2) > 0 {
            v2[0] = -716
        }
    }
    var v214 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v204:", v204)
    fmt.Println("v214:", v214)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v204
    _ = v214
    _ = v3
    _ = v4
    _ = v80
}
