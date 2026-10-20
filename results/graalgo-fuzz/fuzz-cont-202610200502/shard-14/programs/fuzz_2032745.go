package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint; F1 uint; F2 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 int; F2 int16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int8; F1 bool; F2 int32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint32; F1 uint32 }
type GenS9 struct { F0 int32; F1 int }
type GenS10 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS11 struct { F0 uint32; F1 int64; F2 uint }
type GenS12 struct { F0 uint; F1 uint }
type GenS13 struct { F0 int64; F1 uint }
func genFn0(p0 uint64, p1 int32) int16 {
    var v7 int64 = -9223372036854775808
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 int32 = p1
        var v10 uint32 = 42
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := []int{221, 32768}
    v12 := GenS1{F0: 241, F1: 1}
    p1 = (^(p1))
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v7
    return 1000
}
func genFn1(p0 bool, p1 uint64, p2 uint64) uint8 {
    v13 := GenS1{F0: 128, F1: 1}
    p1 = (p1 % ((2) | 1))
    fmt.Println(p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 243
}

func main() {
    v1 := []int{-9223372036854775808, -23}
    v2 := []int{32767}
    v3 := make(map[string]bool)
    var v4 uint8 = 0
    var v5 float64 = float64((0.1 / (float32(-2.25) - 1.5)))
    var v6 int32 = 0
    v14 := GenS1{F0: v4, F1: 42}
    if false {
        var v15 bool = ((^(v4)) > 255)
        _ = v15
    }
    switch "hello" {
    case "foo bar":
        v5 += (0.0 * v5)
    case "":
        v16 := GenS1{F0: v4, F1: genFn0(18446744073709551615, (v6 ^ (v6 / ((1) | 1))))}
        _ = v16
    default:
        v6 -= v6
    }
    if true {
        for fi17 := 0; fi17 < 6; fi17++ {
            if false {
                v18 := GenS1{F0: (v4 >> ((65535) & 63)), F1: -91}
                for fi19 := 0; fi19 < 20; fi19++ {
                    v20 := GenS1{F0: (^(v4)), F1: genFn0(4294967295, (v6 >> ((165) & 63)))}
                    _ = fi19
                    _ = v20
                }
                for fi21 := 0; fi21 < 17; fi21++ {
                    v4 -= v4
                    var v22 uint64 = 2
                    if false {
                        var v23 int16 = genFn0(0, v6)
                        v4 = v4
                        var v24 int8 = -94
                        _ = v23
                        _ = v24
                    }
                    v22 -= (^(v22))
                    _ = fi21
                    _ = v22
                }
                _ = v18
            }
            v25 := GenS1{F0: (v4 ^ 255), F1: -32768}
            _ = fi17
            _ = v25
        }
        v6 -= (v6 / (((v6 | (-(v6)))) | 1))
    }
    v5 = float64(float32(v5))
    if true {
        for fi26 := 0; fi26 < 7; fi26++ {
            var v27 int = fi26
            switch fi26 {
            case 746:
                switch (-(fi26)) {
                case 2147483647:
                    v6 += (-(v6))
                case 327:
                    var v28 uint8 = v4
                    _ = v28
                case 32768:
                    var v29 uint64 = 257
                    _ = v29
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            case -222:
                v30 := GenS1{F0: v4, F1: -32768}
                _ = v30
            default:
                if false {
                    switch (fi26 % ((fi26) | 1)) {
                    case 223:
                        switch "a" {
                        case "\n\t":
                            v31 := GenS1{F0: genFn1(((!false) || (!true)), 1, 0), F1: 0}
                            _ = v31
                        case "hello":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case "a":
                            var v32 string = ((("hello" + "hello") + ((("\n\t" + "\n\t") + "") + "a")) + (("a" + "hello") + ""))
                            _ = v32
                        }
                    case 1:
                        v6 += (v6 ^ 2147483647)
                    default:
                        fmt.Println(v27)
                    }
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(fi26)
                    }
                }
            }
            if false {
                fmt.Println(v6, v5, fi26)
            } else {
                v5 -= (v5 - (v5 - ((v5 - v5) + -2.25)))
                if false {
                    v6 = 0
                    var v33 uint16 = 12
                    var v34 int64 = -1
                    _ = v33
                    _ = v34
                } else {
                    for fi35 := 0; fi35 < 12; fi35++ {
                        var v36 bool = (1e10 > 0.0)
                        _ = fi35
                        _ = v36
                    }
                    v6 = (v6 * v6)
                }
            }
            _ = fi26
            _ = v27
        }
    } else {
        v4 = (v4 & (v4 - (v4 | 128)))
        if (0 >= 42) {
            if false {
                fmt.Println(v5, v4, v5)
            }
            if ((v5 * v5) > -2.25) {
                if (false && (42 <= 32768)) {
                    switch -1 {
                    case 4294967295:
                        v37 := GenS2{F0: 77, F1: 833, F2: -128}
                        _ = v37
                    }
                    switch (("a" + "") + "a") {
                    case "a":
                        v5 = v5
                    case "":
                        for fi38 := 0; fi38 < 11; fi38++ {
                            delete(v3, (("foo bar" + "foo bar") + "a"))
                            var v39 uint16 = 1
                            v39 = (v39 % (((v39 % (((v39 / ((uint16(v4)) | 1))) | 1))) | 1))
                            _ = fi38
                            _ = v39
                        }
                    }
                }
                var v40 int32 = 42
                var v41 int32 = (v6 % ((v40) | 1))
                _ = v40
                _ = v41
            }
        } else {
            var v42 uint = 255
            for fi43 := 0; fi43 < 7; fi43++ {
                for fi44 := 0; fi44 < 13; fi44++ {
                    if false {
                        v45 := GenS2{F0: v42, F1: 1, F2: 42}
                        _ = v45
                    } else {
                        var v46 int = (-(fi43))
                        switch int(v4) {
                        case 2147483647:
                            v6 = (v6 << ((18446744073709551615) & 63))
                        case 127:
                            fmt.Println(fi44)
                        case -193:
                            var v47 uint64 = 2
                            _ = v47
                        default:
                            var v48 float64 = 0.0
                            var v49 string = "foo bar"
                            v3[(v49 + v49)] = (v49 == ((v49 + (("\n\t" + v49) + v49)) + (v49 + "a")))
                            if true {
                                switch (^(fi43)) {
                                case 42:
                                    var v50 float32 = ((0.1 / 1e10) / 1.5)
                                    _ = v50
                                case 349:
                                    var v51 uint16 = 302
                                    _ = v51
                                case 32767:
                                    v5 += (v48 / v5)
                                default:
                                    v3[v49] = true
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            var v52 uint32 = 2147483648
                            if true {
                                switch (fi44 + (fi44 / (((v46 - (v46 >> (((v42 - (v42 >> ((0) & 63)))) & 63)))) | 1))) {
                                case 128:
                                    v6 = v6
                                case -1:
                                    fmt.Println(v5, fi43)
                                    fmt.Println(v46)
                                default:
                                    v53 := GenS2{F0: 42, F1: (v42 - (v42 ^ (v42 ^ 2))), F2: 127}
                                    _ = v53
                                }
                            }
                            v5 -= 1e10
                            switch (v49 + v49) {
                            case "foo bar":
                                switch (^(v46)) {
                                case 4294967295:
                                    delete(v3, v49)
                                }
                                var v54 float32 = 3.14159
                                _ = v54
                            case "":
                                if ((v6 == -786) || true) {
                                    var v55 string = ("foo bar" + v49)
                                    v56 := GenS1{F0: uint8(v42), F1: 433}
                                    var v57 int8 = int8(v4)
                                    _ = v55
                                    _ = v56
                                    _ = v57
                                } else {
                                    v3[v49] = true
                                    var v58 uint16 = 893
                                    v49 = v49
                                    _ = v58
                                }
                            case "a":
                                if true {
                                    if true {
                                        var v59 int16 = 1000
                                        switch v49 {
                                        case "\n\t":
                                            switch 1 {
                                            case 252:
                                                if ((fi43 < fi43) || false) {
                                                    if false {
                                                        var v60 int8 = -1
                                                        var v61 float32 = 3.14159
                                                        _ = v60
                                                        _ = v61
                                                    } else {
                                                        v1 = append(v1, (v46 + (fi43 - fi44)))
                                                    }
                                                } else {
                                                    if true {
                                                        if (2147483647 != int32(v4)) {
                                                            var v62 uint16 = 32767
                                                            var v63 float64 = float64((0.0 + float32(-2.25)))
                                                            switch ("\n\t" + (v49 + "\n\t")) {
                                                            case "\n\t":
                                                                if (!true) {
                                                                    switch int(v42) {
                                                                    case -9223372036854775808:
                                                                        var v64 uint32 = (^(v52))
                                                                        _ = v64
                                                                    case 0:
                                                                        if true {
                                                                            switch (fi44 + fi43) {
                                                                            case 0:
                                                                                var v65 uint16 = v62
                                                                                _ = v65
                                                                            case 1:
                                                                                switch ("hello" + (v49 + "")) {
                                                                                case "\n\t":
                                                                                    var v66 string = v49
                                                                                    _ = v66
                                                                                case "a":
                                                                                    if (!((v52 ^ v52) >= 1)) {
                                                                                        v67 := GenS3{F0: (v6 % (((v6 - 2147483647)) | 1))}
                                                                                        _ = v67
                                                                                    }
                                                                                }
                                                                                if ((^(v4)) > v4) {
                                                                                    v6 = (v6 << ((v42) & 63))
                                                                                    v68 := GenS1{F0: 128, F1: -32768}
                                                                                    var v69 uint16 = (v62 / ((v62) | 1))
                                                                                    if len(v1) >= 2 {
                                                                                        v1 = v1[1:len(v1)]
                                                                                    }
                                                                                    _ = v68
                                                                                    _ = v69
                                                                                } else {
                                                                                    v70 := GenS1{F0: 190, F1: 0}
                                                                                    _ = v70
                                                                                }
                                                                                v46 += 42
                                                                            case 9223372036854775807:
                                                                                v6 = -1
                                                                            }
                                                                            v71 := GenS1{F0: genFn1(((v42 << (((^(v42))) & 63)) < v42), 755, 1), F1: v59}
                                                                            var v72 int8 = -11
                                                                            _ = v71
                                                                            _ = v72
                                                                        } else {
                                                                            v73 := GenS4{F0: (^(v42)), F1: -2147483648, F2: v59}
                                                                            _ = v73
                                                                        }
                                                                    default:
                                                                        v74 := v3[v49]
                                                                        _ = v74
                                                                    }
                                                                    v49 = "foo bar"
                                                                } else {
                                                                    v75 := GenS5{F0: 18446744073709551615}
                                                                    var v76 uint16 = 1
                                                                    v46 = v46
                                                                    _ = v75
                                                                    _ = v76
                                                                }
                                                            default:
                                                                switch (((((v49 + v49) + (("a" + ((v49 + "") + v49)) + v49)) + (v49 + ((((v49 + v49) + ((("a" + v49) + (v49 + v49)) + v49)) + (v49 + (v49 + ((v49 + (((v49 + v49) + "\n\t") + (((("hello" + "") + "hello") + "a") + v49))) + v49)))) + (v49 + ("" + v49))))) + ((("a" + v49) + "foo bar") + (v49 + (v49 + (((v49 + v49) + v49) + v49))))) + v49) {
                                                                case "foo bar":
                                                                    v63 = ((3.14159 * v63) / 0.1)
                                                                default:
                                                                    v46 -= 0
                                                                }
                                                            }
                                                            fmt.Println(v5)
                                                            _ = v62
                                                            _ = v63
                                                        }
                                                        var v77 int32 = v6
                                                        _ = len(v3)
                                                        _ = v77
                                                    }
                                                    v48 = 0.1
                                                    var v78 uint64 = 2
                                                    _ = v78
                                                }
                                            default:
                                                _ = len(v3)
                                            }
                                        case "":
                                            v48 -= float64(1e10)
                                        case "hello":
                                            delete(v3, v49)
                                        default:
                                            var v79 int = 9223372036854775807
                                            _ = v79
                                        }
                                        v48 += (-2.25 - v48)
                                        var v80 string = (v49 + v49)
                                        fmt.Println(v5, fi43)
                                        _ = v59
                                        _ = v80
                                    } else {
                                        var v81 string = v49
                                        _ = v81
                                    }
                                    v82 := GenS3{F0: -21}
                                    _ = v82
                                } else {
                                    switch "a" {
                                    case "foo bar":
                                        var v83 uint64 = 626
                                        _ = v83
                                    case "\n\t":
                                        var v84 int32 = (v6 + v6)
                                        _ = v84
                                    }
                                    fmt.Println(v5, fi43, v48)
                                    v85 := GenS6{F0: 42, F1: true, F2: (v6 * 42)}
                                    _ = v85
                                }
                            }
                            switch fi44 {
                            case 32767:
                                var v86 float64 = float64(0.1)
                                _ = v86
                            case 1:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            case 424:
                                v87 := GenS6{F0: 0, F1: true, F2: v6}
                                _ = v87
                            }
                            _ = v48
                            _ = v49
                            _ = v52
                        }
                        if false {
                            var v88 float32 = 0.0
                            v89 := GenS6{F0: 0, F1: false, F2: 2147483647}
                            var v90 bool = ((0.1 / 1e10) >= float32(v5))
                            _ = len(v1)
                            if (!(256 < 2147483648)) {
                                v91 := v3[("\n\t" + "\n\t")]
                                _ = v91
                            }
                            _ = v88
                            _ = v89
                            _ = v90
                        }
                        v92 := GenS1{F0: v4, F1: -1}
                        _ = v46
                        _ = v92
                    }
                    _ = fi44
                }
                fmt.Println(v42, fi43)
                for fi93 := 0; fi93 < 8; fi93++ {
                    var v94 int64 = 42
                    switch int(v42) {
                    case 2147483647:
                        fmt.Println(v42, v5, v94)
                    }
                    var v95 uint32 = 42
                    _ = fi93
                    _ = v94
                    _ = v95
                }
                _ = fi43
            }
            var v96 float32 = -2.25
            _ = v42
            _ = v96
        }
    }
    switch 386 {
    case -802:
        fmt.Println(v4, v6, v4)
    case -9223372036854775808:
        if (1 > 4294967295) {
            if true {
                for fi97 := 0; fi97 < 20; fi97++ {
                    for fi98 := 0; fi98 < 17; fi98++ {
                        v6 += (v6 >> ((0) & 63))
                        _ = fi98
                    }
                    fmt.Println(v5, v4)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi97
                }
                if ((!false) && false) {
                    if false {
                        v4 -= (v4 * (^(v4)))
                    } else {
                        for fi99 := 0; fi99 < 3; fi99++ {
                            if len(v2) > 0 {
                                v2[0] = 127
                            }
                            var v100 uint64 = 1
                            if (!true) {
                                v101 := GenS7{F0: ((v5 * v5) < v5)}
                                var v102 string = ("foo bar" + "\n\t")
                                _ = v101
                                _ = v102
                            }
                            _ = fi99
                            _ = v100
                        }
                        fmt.Println(v6, v6)
                    }
                    if (!(v6 >= -1)) {
                        v103 := GenS6{F0: 1, F1: (440 < 4294967295), F2: (^(v6))}
                        if (!(643 >= 750)) {
                            v5 = v5
                            var v104 uint32 = 2147483648
                            for fi105 := 0; fi105 < 8; fi105++ {
                                for fi106 := 0; fi106 < 1; fi106++ {
                                    var v107 uint8 = 0
                                    fmt.Println(v6, v104, v5)
                                    _ = fi106
                                    _ = v107
                                }
                                _ = fi105
                            }
                            _ = v104
                        } else {
                            var v108 int = -1
                            if (0 < -1) {
                                var v109 bool = false
                                switch -1 {
                                case 1:
                                    v110 := GenS2{F0: 79, F1: uint(v4), F2: 97}
                                    _ = v110
                                case 4294967295:
                                    v111 := GenS7{F0: (!v109)}
                                    _ = v111
                                case 127:
                                    for fi112 := 0; fi112 < 6; fi112++ {
                                        fmt.Println(fi112, v109, fi112)
                                        _ = fi112
                                    }
                                default:
                                    v113 := GenS8{F0: 1, F1: 155}
                                    _ = v113
                                }
                                _ = v109
                            }
                            _ = v108
                        }
                        var v114 int64 = 42
                        _ = v103
                        _ = v114
                    }
                }
            }
            v115 := GenS8{F0: 42, F1: 932}
            if (0 == 2147483647) {
                var v116 uint8 = 1
                var v117 bool = false
                switch -9223372036854775808 {
                case -1:
                    v118 := v3[(("foo bar" + "\n\t") + "foo bar")]
                    _ = v118
                case 42:
                    fmt.Println(v5, v6)
                case 528:
                    var v119 int16 = -978
                    _ = v119
                }
                _ = v116
                _ = v117
            } else {
                v6 += 2147483647
            }
            _ = v115
        } else {
            if false {
                v1 = append(v1, 9223372036854775807)
                for fi120 := 0; fi120 < 8; fi120++ {
                    v4 -= (v4 << ((0) & 63))
                    for fi121 := 0; fi121 < 12; fi121++ {
                        if (genFn1(true, 1, 82) <= genFn1(true, 75, 0)) {
                            var v122 bool = (true && false)
                            switch fi120 {
                            case 128:
                                if false {
                                    if v122 {
                                        if len(v2) > 0 {
                                            v2[0] = (fi120 & 128)
                                        }
                                        if false {
                                            var v123 int64 = -72
                                            if false {
                                                v5 += 3.14159
                                                if ((v5 / v5) != (v5 * v5)) {
                                                    switch "a" {
                                                    case "\n\t":
                                                        if len(v1) > 0 {
                                                            v1[0] = (fi120 / ((fi121) | 1))
                                                        }
                                                    case "":
                                                        fmt.Println(v123)
                                                    case "a":
                                                        if v122 {
                                                            if true {
                                                                var v124 string = ((((("hello" + "a") + "hello") + "\n\t") + (("\n\t" + "a") + "foo bar")) + "foo bar")
                                                                v125 := GenS8{F0: 916, F1: 178}
                                                                var v126 uint32 = uint32(v4)
                                                                v5 -= float64((3.14159 - 1e10))
                                                                v127 := v3[v124]
                                                                var v128 uint16 = 65535
                                                                var v129 uint32 = (v126 & (v126 - (v126 ^ v126)))
                                                                var v130 int32 = v6
                                                                v123 += (v123 - (v123 | (^(v123))))
                                                                _ = v124
                                                                _ = v125
                                                                _ = v126
                                                                _ = v127
                                                                _ = v128
                                                                _ = v129
                                                                _ = v130
                                                            } else {
                                                                var v131 uint64 = 2
                                                                if false {
                                                                    var v132 bool = (1 <= 19)
                                                                    if v132 {
                                                                        fmt.Println(v132)
                                                                        var v133 int8 = 42
                                                                        var v134 uint32 = 2
                                                                        _ = v133
                                                                        _ = v134
                                                                    } else {
                                                                        v132 = (42 >= 2)
                                                                        if (v6 != (^(v6))) {
                                                                            if v122 {
                                                                                if (("foo bar" >= "hello") && v132) {
                                                                                    if v122 {
                                                                                        v5 = 1.5
                                                                                        v6 = (v6 >> ((397) & 63))
                                                                                    }
                                                                                    v122 = (v131 == 18446744073709551615)
                                                                                }
                                                                                var v135 uint32 = 4294967295
                                                                                v136 := GenS9{F0: v6, F1: fi121}
                                                                                _ = v135
                                                                                _ = v136
                                                                            } else {
                                                                                v137 := GenS10{F0: 229, F1: v131, F2: 2}
                                                                                _ = v137
                                                                            }
                                                                            if true {
                                                                                if v122 {
                                                                                    switch (-(fi121)) {
                                                                                    case 553:
                                                                                        v123 -= v123
                                                                                    case 127:
                                                                                        var v138 float32 = 1e10
                                                                                        _ = v138
                                                                                    case -115:
                                                                                        var v139 uint64 = (v131 - 42)
                                                                                        _ = v139
                                                                                    }
                                                                                    v140 := GenS8{F0: 2, F1: 1}
                                                                                    v6 = v6
                                                                                    _ = v140
                                                                                }
                                                                                var v141 int8 = 55
                                                                                _ = v141
                                                                            } else {
                                                                                var v142 uint = 0
                                                                                var v143 uint32 = 374
                                                                                _ = v142
                                                                                _ = v143
                                                                            }
                                                                            v132 = (390 >= 32767)
                                                                            fmt.Println(v132, v5, v132)
                                                                            var v144 float64 = -2.25
                                                                            _ = v144
                                                                        } else {
                                                                            delete(v3, "\n\t")
                                                                            var v145 float64 = v5
                                                                            _ = v145
                                                                        }
                                                                    }
                                                                    v146 := GenS11{F0: 0, F1: -9223372036854775808, F2: 171}
                                                                    var v147 uint32 = 42
                                                                    _ = v132
                                                                    _ = v146
                                                                    _ = v147
                                                                } else {
                                                                    if v122 {
                                                                        v148 := GenS2{F0: 42, F1: 896, F2: -82}
                                                                        v123 += v123
                                                                        v149 := v3["\n\t"]
                                                                        _ = len(v1)
                                                                        _ = v148
                                                                        _ = v149
                                                                    }
                                                                    v150 := GenS12{F0: 520, F1: 476}
                                                                    _ = v150
                                                                }
                                                                var v151 uint16 = 780
                                                                _ = v131
                                                                _ = v151
                                                            }
                                                            if ((v6 % ((v6) | 1)) > v6) {
                                                                var v152 string = ("a" + (("" + ((("foo bar" + "") + "foo bar") + "\n\t")) + "a"))
                                                                var v153 int8 = -125
                                                                _ = v152
                                                                _ = v153
                                                            }
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                var v154 int32 = v6
                                                                _ = v154
                                                            case "hello":
                                                                var v155 string = "foo bar"
                                                                var v156 uint64 = 231
                                                                if true {
                                                                    var v157 uint64 = 42
                                                                    _ = v157
                                                                }
                                                                _ = v155
                                                                _ = v156
                                                            case "a":
                                                                fmt.Println(v5, v6, v5)
                                                            }
                                                        }
                                                    default:
                                                        var v158 int32 = (v6 * 2147483647)
                                                        _ = v158
                                                    }
                                                    var v159 int8 = 0
                                                    var v160 uint32 = 65535
                                                    v123 = v123
                                                    if v122 {
                                                        if ((!false) || true) {
                                                            var v161 uint8 = v4
                                                            var v162 uint8 = (^(v161))
                                                            _ = v161
                                                            _ = v162
                                                        }
                                                        fmt.Println(fi120, v123)
                                                        var v163 uint8 = (v4 / (((v4 >> ((2) & 63))) | 1))
                                                        _ = v163
                                                    } else {
                                                        var v164 uint64 = 717
                                                        var v165 float32 = 0.0
                                                        var v166 uint64 = uint64(v160)
                                                        _ = v164
                                                        _ = v165
                                                        _ = v166
                                                    }
                                                    _ = v159
                                                    _ = v160
                                                } else {
                                                    fmt.Println(v123)
                                                    var v167 uint = 0
                                                    var v168 int = fi120
                                                    _ = v167
                                                    _ = v168
                                                }
                                            } else {
                                                v169 := v3["\n\t"]
                                                var v170 bool = (-818 >= 32767)
                                                _ = v169
                                                _ = v170
                                            }
                                            _ = v123
                                        } else {
                                            v6 = v6
                                            var v171 string = "foo bar"
                                            v6 = (v6 >> ((1) & 63))
                                            _ = v171
                                        }
                                    }
                                    v172 := GenS13{F0: int64(fi121), F1: 73}
                                    _ = v172
                                } else {
                                    var v173 uint16 = 2
                                    v4 = (v4 >> ((564) & 63))
                                    _ = v173
                                }
                            default:
                                delete(v3, "foo bar")
                            }
                            v6 -= (v6 ^ v6)
                            _ = v122
                        } else {
                            delete(v3, "hello")
                            if (42 >= 4294967295) {
                                var v174 int32 = 1
                                _ = v174
                            }
                        }
                        fmt.Println(v6, v5, fi121)
                        switch "hello" {
                        case "":
                            delete(v3, "a")
                        default:
                            switch -200 {
                            case 2147483647:
                                switch (fi121 - (fi120 / (((fi120 % ((fi120) | 1))) | 1))) {
                                case -741:
                                    switch "a" {
                                    case "foo bar":
                                        if ((v4 / ((v4) | 1)) < 128) {
                                            var v175 int16 = 821
                                            _ = v175
                                        }
                                    }
                                case 13:
                                    v4 = v4
                                case -377:
                                    v5 -= ((1e10 / 3.14159) + (v5 + 0.0))
                                default:
                                    v2 = append(v2, -1000)
                                    var v176 int = -2147483648
                                    if false {
                                        var v177 int = (fi120 << ((255) & 63))
                                        _ = v177
                                    } else {
                                        var v178 int = (v176 - fi120)
                                        _ = v178
                                    }
                                    v179 := GenS9{F0: (-(v6)), F1: fi121}
                                    _ = v176
                                    _ = v179
                                }
                            }
                        }
                        _ = fi121
                    }
                    _ = fi120
                }
            }
            v6 -= (v6 >> ((660) & 63))
            var v180 int8 = 90
            _ = v180
        }
    case -1:
        fmt.Println(v6, v6)
    }
    v5 -= 1.5
    var v181 int32 = (v6 << ((65535) & 63))
    var v182 uint8 = (v4 % (((v4 ^ v4)) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v181:", v181)
    fmt.Println("v182:", v182)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v181
    _ = v182
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
