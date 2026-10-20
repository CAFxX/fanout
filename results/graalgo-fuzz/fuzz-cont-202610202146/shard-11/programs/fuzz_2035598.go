package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 string }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 int8; F2 uint8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint8; F1 bool }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int16; F1 int64; F2 uint16 }
type GenS9 struct { F0 int32; F1 int16; F2 int }
type GenS10 struct { F0 uint; F1 int16 }
type GenS11 struct { F0 int; F1 uint32; F2 uint16 }
type GenS12 struct { F0 int8 }

func main() {
    var v1 int32 = -171
    var v2 uint64 = 0
    var v3 uint = 859
    v4 := make(map[uint32]bool)
    v5 := []int{}
    for fi6 := 0; fi6 < 4; fi6++ {
        for fi7 := 0; fi7 < 6; fi7++ {
            if (v1 > (v1 * (-(v1)))) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v8 := GenS1{F0: 1, F1: (v3 / (((v3 & v3)) | 1)), F2: "hello"}
                v1 = v1
                var v9 uint = 994
                switch 127 {
                case 1:
                    switch (fi7 + (fi6 & (fi6 - (fi7 / ((4294967295) | 1))))) {
                    case -787:
                        if false {
                            v9 += v9
                            var v10 uint8 = 127
                            if (79 >= 16) {
                                var v11 int8 = 1
                                var v12 float64 = (3.14159 + (((3.14159 / (1.5 * (0.1 * (3.14159 - (((0.1 * (0.0 / -2.25)) * -2.25) - 1e10))))) / ((float64(1e10) + -2.25) - 0.1)) + 1e10))
                                _ = v11
                                _ = v12
                            }
                            _ = v10
                        }
                    default:
                        if (42 > -1) {
                            var v13 uint = 42
                            _ = v13
                        }
                    }
                case 0:
                    if (!((421 == 109) || (65535 <= 2147483648))) {
                        var v14 int16 = 32767
                        var v15 uint = (v9 >> (((^(v3))) & 63))
                        _ = v14
                        _ = v15
                    } else {
                        var v16 int64 = -1
                        if true {
                            v17 := GenS1{F0: 42, F1: (v3 - 131), F2: "a"}
                            _ = v17
                        }
                        _ = v16
                    }
                }
                _ = v8
                _ = v9
            } else {
                v18 := GenS1{F0: 4294967295, F1: 341, F2: "foo bar"}
                _ = len(v4)
                if ((!true) || ((310 != 42) || (65535 >= 1))) {
                    var v19 int8 = 127
                    if false {
                        v20 := GenS2{F0: uint32(v3), F1: 248}
                        fmt.Println(v19, v2, fi6)
                        _ = v20
                    } else {
                        v3 += (v3 * (v3 | (v3 >> ((uint(v2)) & 63))))
                        fmt.Println(fi6, v3)
                        switch ((("\n\t" + ("hello" + ("" + ("" + (("\n\t" + "\n\t") + ("hello" + "hello")))))) + "foo bar") + "\n\t") {
                        case "a":
                            switch ("hello" + "\n\t") {
                            case "hello":
                                _ = len(v4)
                            case "foo bar":
                                var v21 int64 = -162
                                _ = v21
                            case "":
                                if false {
                                    _ = len(v5)
                                    v3 += 42
                                    if (-2147483648 >= v1) {
                                        var v22 uint8 = 127
                                        _ = v22
                                    } else {
                                        v23 := GenS2{F0: 651, F1: -9223372036854775808}
                                        _ = v23
                                    }
                                }
                            default:
                                var v24 bool = false
                                _ = v24
                            }
                        case "foo bar":
                            v1 = v1
                        default:
                            var v25 int16 = 1
                            _ = v25
                        }
                    }
                    _ = v19
                } else {
                    v1 += int32(v3)
                    v1 += (v1 * 42)
                    switch 128 {
                    case -453:
                        v26 := GenS1{F0: 2, F1: (^(v3)), F2: ("foo bar" + (("" + "\n\t") + ""))}
                        _ = v26
                    case 4294967295:
                        switch (fi7 % ((42) | 1)) {
                        case -9223372036854775808:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                    case 2147483647:
                        if true {
                            if len(v5) > 0 {
                                v5[0] = fi7
                            }
                        } else {
                            if ((true || ((v3 >> ((v3) & 63)) == (v3 / (((v3 - v3)) | 1)))) || (-1 > fi7)) {
                                var v27 string = "hello"
                                switch fi6 {
                                case -2147483648:
                                    if false {
                                        if (!(1.5 != 3.14159)) {
                                            v27 = "foo bar"
                                        }
                                        var v28 float32 = 1.5
                                        if ((fi6 | fi6) <= (fi6 / ((32768) | 1))) {
                                            fmt.Println(v1, fi7, fi6)
                                            v28 = (((0.0 / v28) - float32(0.0)) - float32(((0.1 / ((0.1 + 1.5) * 1.5)) * float64((v28 / v28)))))
                                            if false {
                                                if (2147483647 >= 42) {
                                                    v29 := GenS3{F0: 0}
                                                    if ((0 != 2147483647) || (1e10 == 0.0)) {
                                                        v30 := GenS1{F0: 2147483648, F1: 18446744073709551615, F2: "a"}
                                                        v27 = v27
                                                        v5 = append(v5, fi7)
                                                        _ = v30
                                                    } else {
                                                        var v31 int32 = v1
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            var v32 int64 = -9223372036854775808
                                                            _ = v32
                                                        case "":
                                                            v2 -= (v2 << (((v3 | (v3 ^ (v3 + v3)))) & 63))
                                                        case "hello":
                                                            var v33 int32 = (v31 >> ((v3) & 63))
                                                            _ = v33
                                                        default:
                                                            v2 = v2
                                                        }
                                                        v28 += -2.25
                                                        if false {
                                                            fmt.Println(fi6)
                                                            var v34 int64 = int64(v31)
                                                            if (true || false) {
                                                                var v35 int = (fi6 & 648)
                                                                var v36 int16 = 1000
                                                                v37 := GenS4{F0: -52, F1: -27, F2: 189}
                                                                _ = v35
                                                                _ = v36
                                                                _ = v37
                                                            }
                                                            _ = v34
                                                        } else {
                                                            v5 = append(v5, 518)
                                                            _ = len(v5)
                                                        }
                                                        _ = v31
                                                    }
                                                    v38 := GenS5{F0: -128}
                                                    v3 += (v3 << (((v3 - (^(v3)))) & 63))
                                                    _ = v29
                                                    _ = v38
                                                }
                                                v27 = (("\n\t" + "") + v27)
                                                v2 = 248
                                            }
                                            v3 += (^(v3))
                                            switch fi6 {
                                            case -317:
                                                if (false && ((v2 >> ((65535) & 63)) <= v2)) {
                                                    v39 := v4[65535]
                                                    fmt.Println(v27, v28, fi6)
                                                    _ = v39
                                                } else {
                                                    fmt.Println(v2, v1, v3)
                                                    if (!(true || false)) {
                                                        var v40 int16 = 42
                                                        _ = len(v5)
                                                        v41 := GenS2{F0: 4, F1: -527}
                                                        _ = v40
                                                        _ = v41
                                                    }
                                                    v2 -= v2
                                                }
                                            case 1:
                                                fmt.Println(v3, v27)
                                            default:
                                                v42 := GenS6{F0: 42, F1: true}
                                                _ = v42
                                            }
                                            if (int8(v1) != 1) {
                                                v43 := GenS7{F0: true}
                                                v28 = v28
                                                var v44 float32 = -2.25
                                                v1 = v1
                                                v3 -= (^(v3))
                                                _ = v43
                                                _ = v44
                                            } else {
                                                var v45 uint8 = 255
                                                _ = v45
                                            }
                                        } else {
                                            if (!(380 < 140)) {
                                                if ((566 <= 835) || true) {
                                                    if (!false) {
                                                        v4[42] = false
                                                    }
                                                    if true {
                                                        v5 = append(v5, (fi7 * (fi7 ^ 127)))
                                                        if true {
                                                            var v46 uint16 = 42
                                                            v47 := GenS8{F0: -884, F1: -1, F2: (^(v46))}
                                                            var v48 uint32 = 312
                                                            _ = v46
                                                            _ = v47
                                                            _ = v48
                                                        }
                                                        _ = len(v4)
                                                    }
                                                } else {
                                                    v1 -= (^(v1))
                                                    if ((^(fi7)) == (fi6 >> ((v3) & 63))) {
                                                        if true {
                                                            var v49 int64 = 0
                                                            v27 = v27
                                                            v3 -= v3
                                                            var v50 uint32 = 65535
                                                            var v51 uint32 = v50
                                                            _ = v49
                                                            _ = v50
                                                            _ = v51
                                                        } else {
                                                            switch (fi7 / ((0) | 1)) {
                                                            case 4294967295:
                                                                if ((v2 / (((v2 * 817)) | 1)) <= (v2 | (v2 - v2))) {
                                                                    v52 := GenS9{F0: 2147483647, F1: -32768, F2: 940}
                                                                    switch fi7 {
                                                                    case -1:
                                                                        var v53 float32 = 1.5
                                                                        _ = v53
                                                                    case -2147483648:
                                                                        fmt.Println(v27, v1, fi7)
                                                                    }
                                                                    fmt.Println(v1, v2, v1)
                                                                    _ = v52
                                                                }
                                                            case 32767:
                                                                v2 = v2
                                                            }
                                                            var v54 bool = false
                                                            var v55 int16 = -1
                                                            _ = v54
                                                            _ = v55
                                                        }
                                                        v56 := GenS2{F0: 2147483647, F1: 42}
                                                        _ = v56
                                                    }
                                                }
                                            } else {
                                                v28 = float32(3.14159)
                                                if true {
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        if (true || (fi7 <= fi7)) {
                                                            v1 += v1
                                                            if (fi7 < (fi6 - 4294967295)) {
                                                                v28 += v28
                                                            }
                                                        }
                                                    case "":
                                                        var v57 int16 = 0
                                                        _ = v57
                                                    case "\n\t":
                                                        var v58 uint16 = uint16(v2)
                                                        _ = v58
                                                    }
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                }
                                            }
                                        }
                                        _ = v28
                                    } else {
                                        v4[65535] = true
                                        if false {
                                            v5 = append(v5, 417)
                                            if (!false) {
                                                v59 := GenS7{F0: (!false)}
                                                v60 := GenS9{F0: v1, F1: -32768, F2: 32768}
                                                _ = v59
                                                _ = v60
                                            }
                                        }
                                        if (42 == 51) {
                                            var v61 bool = (fi6 != 128)
                                            switch v27 {
                                            case "hello":
                                                var v62 int8 = 42
                                                _ = v62
                                            }
                                            if false {
                                                v1 -= -158
                                                var v63 bool = (!(0.1 != 3.14159))
                                                var v64 uint16 = 66
                                                switch ((v27 + (v27 + (v27 + v27))) + "hello") {
                                                case "foo bar":
                                                    var v65 uint64 = v2
                                                    _ = v65
                                                }
                                                _ = v63
                                                _ = v64
                                            }
                                            delete(v4, 2)
                                            var v66 uint32 = 2
                                            _ = v61
                                            _ = v66
                                        } else {
                                            switch -9223372036854775808 {
                                            case -728:
                                                v67 := GenS2{F0: 460, F1: -786}
                                                _ = v67
                                            case 42:
                                                v1 = v1
                                            default:
                                                if (true || (v1 == v1)) {
                                                    v5 = append(v5, (fi7 - int(v2)))
                                                } else {
                                                    var v68 int64 = -165
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    v5 = append(v5, 42)
                                                    v2 = (^(v2))
                                                    if false {
                                                        var v69 float64 = -2.25
                                                        _ = v69
                                                    }
                                                    _ = v68
                                                }
                                            }
                                            var v70 uint16 = 32767
                                            _ = v70
                                        }
                                    }
                                case 127:
                                    var v71 uint = (v3 % (((v3 % (((v3 ^ (v3 ^ (^(v3))))) | 1))) | 1))
                                    _ = v71
                                case -9223372036854775808:
                                    var v72 uint = 1
                                    var v73 int32 = (v1 + int32(v72))
                                    fmt.Println(v1)
                                    var v74 int32 = (v1 - (^(v73)))
                                    _ = v72
                                    _ = v73
                                    _ = v74
                                }
                                _ = v27
                            }
                        }
                    }
                    if (false && false) {
                        v75 := GenS10{F0: 483, F1: 1}
                        if (493 == 4294967295) {
                            var v76 bool = ((-38 <= 110) || (3.14159 <= 1.5))
                            v4[4294967295] = (true && (v2 < (v2 % ((4294967295) | 1))))
                            v3 += 4294967295
                            _ = v76
                        }
                        _ = v75
                    }
                    v77 := GenS6{F0: 2, F1: true}
                    _ = v77
                }
                _ = v18
            }
            v78 := GenS6{F0: 1, F1: (fi6 >= fi6)}
            v3 += 65535
            fmt.Println(v1)
            _ = fi7
            _ = v78
        }
        for fi79 := 0; fi79 < 3; fi79++ {
            if (!(42 > 952)) {
                v2 = 4294967295
                var v80 float32 = 1e10
                if (!(392 < 2147483648)) {
                    v80 = ((0.1 * ((float32(3.14159) + float32((((0.0 + -2.25) * 1.5) * (1e10 - 0.1)))) + v80)) * 0.1)
                    var v81 uint16 = 32767
                    switch -775 {
                    case 128:
                        switch "" {
                        case "\n\t":
                            v3 = 323
                        case "foo bar":
                            var v82 int = (fi79 * (fi6 | 4294967295))
                            _ = v82
                        case "a":
                            v83 := GenS5{F0: 100}
                            _ = v83
                        default:
                            _ = len(v5)
                        }
                    case 2147483647:
                        if ((false || (true && false)) && ((v1 + (v1 / ((-1) | 1))) != 528)) {
                            switch 563 {
                            case 0:
                                v81 -= 245
                            }
                            var v84 string = ((("" + "hello") + ("hello" + "foo bar")) + "")
                            if true {
                                var v85 int = fi6
                                _ = v85
                            }
                            _ = v84
                        } else {
                            fmt.Println(v2, v3, v80)
                            v86 := GenS5{F0: 124}
                            fmt.Println(v80, fi6, fi79)
                            _ = v86
                        }
                    default:
                        v2 += (v2 + 0)
                    }
                    _ = v81
                } else {
                    if false {
                        v2 = (v2 + (v2 & (v2 * v2)))
                        var v87 bool = true
                        v88 := GenS10{F0: (v3 | v3), F1: 390}
                        _ = v87
                        _ = v88
                    } else {
                        v89 := GenS3{F0: -32768}
                        var v90 int8 = 127
                        _ = v89
                        _ = v90
                    }
                    if true {
                        switch ("foo bar" + "") {
                        case "foo bar":
                            v80 -= ((float32(((0.0 + (1e10 + (0.0 / 1.0))) / (1.5 / 0.1))) - v80) * 0.1)
                        case "\n\t":
                            var v91 float32 = (float32(1e10) + v80)
                            v4[2147483648] = (("a" + "a") == "\n\t")
                            delete(v4, 4294967295)
                            v91 -= (1.5 - v91)
                            _ = v91
                        case "a":
                            v3 = (v3 + v3)
                            v3 -= (v3 / ((65535) | 1))
                        }
                    } else {
                        var v92 string = ""
                        if (89 == 94) {
                            v5 = append(v5, 269)
                            var v93 int = (fi79 - (-(fi79)))
                            _ = v93
                        }
                        v80 = 1e10
                        if false {
                            v3 -= (v3 << ((65535) & 63))
                            var v94 float64 = 0.0
                            v2 += (v2 ^ v2)
                            _ = v94
                        }
                        var v95 uint32 = 65535
                        var v96 float64 = ((1e10 / 1.0) / 0.1)
                        _ = v92
                        _ = v95
                        _ = v96
                    }
                }
                _ = v80
            } else {
                v1 += v1
                v1 = v1
            }
            _ = fi79
        }
        _ = fi6
    }
    v3 += v3
    if len(v5) > 0 {
        v5[0] = 42
    }
    fmt.Println(v3, v3, v3)
    var v97 uint64 = (v2 & (v2 % ((v2) | 1)))
    if (0 < 36) {
        var v98 uint16 = 927
        _ = v98
    } else {
        var v99 int64 = -982
        if false {
            var v100 int64 = (-(v99))
            _ = v100
        } else {
            if ("a" < "a") {
                switch 127 {
                case -2147483648:
                    var v101 bool = (("foo bar" < (("foo bar" + ("hello" + "foo bar")) + (("\n\t" + "foo bar") + "a"))) || false)
                    _ = v101
                case 0:
                    var v102 uint = 16
                    _ = v102
                }
            }
            if true {
                v3 = 18446744073709551615
                v99 += (v99 * 9223372036854775807)
            } else {
                _ = len(v5)
                v97 = (v2 % ((628) | 1))
            }
        }
        v99 = (v99 * (v99 ^ (-(v99))))
        _ = v99
    }
    v103 := GenS8{F0: 0, F1: -1, F2: 223}
    fmt.Println(v3, v1)
    if (v1 <= 42) {
        var v104 string = "a"
        v105 := GenS11{F0: int(v97), F1: 198, F2: 65535}
        _ = v104
        _ = v105
    }
    if (2 <= 42) {
        if len(v5) > 0 {
            v5[0] = 127
        }
    } else {
        for fi106 := 0; fi106 < 2; fi106++ {
            v3 += 18446744073709551615
            _ = fi106
        }
        var v107 float32 = float32((1.5 / (1.5 / 1.0)))
        var v108 int16 = 42
        _ = v107
        _ = v108
    }
    for fi109 := 0; fi109 < 8; fi109++ {
        if true {
            v2 -= v97
        }
        v1 = v1
        _ = fi109
    }
    switch 128 {
    case 192:
        if (false || (!(2 <= 128))) {
            for fi110 := 0; fi110 < 7; fi110++ {
                for fi111 := 0; fi111 < 4; fi111++ {
                    var v112 int64 = 0
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(v1, fi110)
                    case "\n\t":
                        var v113 string = "a"
                        _ = v113
                    default:
                        if (!(!(-128 >= -128))) {
                            var v114 int = 4294967295
                            if ((-(v1)) >= (^(v1))) {
                                if len(v5) > 0 {
                                    v5[0] = 127
                                }
                            }
                            if true {
                                v115 := GenS12{F0: 67}
                                _ = v115
                            }
                            v1 += v1
                            var v116 string = ""
                            _ = v114
                            _ = v116
                        } else {
                            var v117 int64 = (^(v112))
                            fmt.Println(v97)
                            v5 = append(v5, -63)
                            _ = v117
                        }
                        if true {
                            if true {
                                var v118 int8 = 127
                                v2 -= (v97 - 42)
                                _ = v118
                            }
                        } else {
                            if false {
                                if false {
                                    var v119 string = "a"
                                    _ = v119
                                }
                                fmt.Println(v112, v112)
                            }
                        }
                    }
                    _ = fi111
                    _ = v112
                }
                v120 := GenS3{F0: 42}
                _ = fi110
                _ = v120
            }
            var v121 uint32 = 2147483648
            switch 127 {
            case 32768:
                var v122 string = "hello"
                _ = v122
            case 9223372036854775807:
                for fi123 := 0; fi123 < 16; fi123++ {
                    var v124 float32 = (-2.25 - 3.14159)
                    for fi125 := 0; fi125 < 16; fi125++ {
                        v126 := GenS11{F0: (fi123 >> (((v3 % (((v3 ^ v3)) | 1))) & 63)), F1: v121, F2: 65535}
                        _ = fi125
                        _ = v126
                    }
                    v1 = v1
                    _ = fi123
                    _ = v124
                }
            case -357:
                if (1e10 >= 0.0) {
                    fmt.Println(v2, v121)
                    v2 += (^(v97))
                } else {
                    fmt.Println(v1, v1)
                    v127 := GenS2{F0: v121, F1: -9223372036854775808}
                    _ = v127
                }
            default:
                v97 = v97
            }
            _ = v121
        } else {
            fmt.Println(v3, v3)
            v97 = 42
        }
    case 42:
        v1 += v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v103
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v97
}
