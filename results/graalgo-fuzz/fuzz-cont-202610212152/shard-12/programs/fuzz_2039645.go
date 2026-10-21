package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 string; F1 int16; F2 uint16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 string; F1 uint16 }
type GenS9 struct { F0 uint32; F1 bool }
type GenS10 struct { F0 uint8; F1 uint16; F2 bool }
type GenS11 struct { F0 uint32; F1 uint }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 string }
type GenS15 struct { F0 bool; F1 int64; F2 uint8 }
type GenS16 struct { F0 string; F1 uint16 }
type GenS17 struct { F0 string; F1 int8; F2 uint8 }
type GenS18 struct { F0 int8; F1 string }
type GenS19 struct { F0 uint }
type GenS20 struct { F0 uint64; F1 int16; F2 bool }
type GenS21 struct { F0 int64 }
func genFn0(p0 uint8) uint16 {
    v5 := []int{-913, -2147483648, -566, 127}
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := GenS1{F0: (("foo bar" + (((((("" + (("" + "\n\t") + "\n\t")) + ("a" + ("a" + "foo bar"))) + "a") + ("a" + "foo bar")) + ("foo bar" + "\n\t")) + "a")) + "hello"), F1: 859}
        if true {
            if true {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                for fi8 := 0; fi8 < 5; fi8++ {
                    if (int32(fi8) < 2147483647) {
                        var v9 uint64 = 18446744073709551615
                        p0 = p0
                        var v10 uint = 65535
                        _ = v10
                        _ = v9
                    }
                    if len(v5) > 0 {
                        v5[0] = fi8
                    }
                    _ = fi8
                }
                if true {
                    var v11 int = fi6
                    _ = v11
                } else {
                    if true {
                        switch "" {
                        case "foo bar":
                            var v12 uint64 = 18446744073709551615
                            _ = v12
                        default:
                            p0 += (p0 << ((1) & 63))
                        }
                        var v13 int16 = -304
                        for fi14 := 0; fi14 < 9; fi14++ {
                            switch "\n\t" {
                            case "hello":
                                v15 := GenS1{F0: "foo bar", F1: 255}
                                _ = v15
                            case "\n\t":
                                fmt.Println(fi14, v13)
                            case "a":
                                var v16 float32 = 1e10
                                _ = v16
                            }
                            if false {
                                _ = len(v5)
                                v13 += -932
                                v13 = v13
                            } else {
                                var v17 float64 = ((0.1 + 3.14159) - 3.14159)
                                switch fi14 {
                                case 9223372036854775807:
                                    var v18 uint = 808
                                    fmt.Println(fi14)
                                    v19 := GenS2{F0: 0, F1: 42}
                                    switch (fi6 + fi14) {
                                    case -2147483648:
                                        var v20 int16 = v13
                                        _ = v20
                                    case 603:
                                        var v21 int = (^(fi6))
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        _ = v21
                                    case 128:
                                        switch (("hello" + ("a" + ("foo bar" + "foo bar"))) + (("\n\t" + "hello") + (("hello" + "hello") + "a"))) {
                                        case "hello":
                                            var v22 string = ("foo bar" + ("\n\t" + ("\n\t" + "hello")))
                                            _ = v22
                                        }
                                    }
                                    _ = v18
                                    _ = v19
                                case 4294967295:
                                    switch "\n\t" {
                                    case "a":
                                        _ = len(v5)
                                    case "hello":
                                        if false {
                                            var v23 bool = (42 == -466)
                                            v13 = -453
                                            if (!(v23 || v23)) {
                                                var v24 int64 = 0
                                                v5 = append(v5, fi14)
                                                _ = v24
                                            }
                                            switch "foo bar" {
                                            case "a":
                                                var v25 uint64 = 42
                                                _ = v25
                                            case "\n\t":
                                                var v26 uint32 = 2
                                                _ = v26
                                            case "foo bar":
                                                var v27 int = (^(fi14))
                                                _ = v27
                                            }
                                            var v28 float32 = 1.5
                                            if (p0 > p0) {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v29 string = ""
                                                    if v23 {
                                                        var v30 float64 = ((float64((float32(v17) * 3.14159)) / v17) / (((0.0 + v17) / v17) * v17))
                                                        v31 := GenS3{F0: 1, F1: 32767}
                                                        v32 := GenS3{F0: -162, F1: 255}
                                                        if v23 {
                                                            if len(v5) > 0 {
                                                                v5[0] = fi6
                                                            }
                                                            v17 += (v17 * -2.25)
                                                        }
                                                        fmt.Println(v30, fi14)
                                                        _ = v30
                                                        _ = v31
                                                        _ = v32
                                                    } else {
                                                        var v33 int16 = -606
                                                        if (1e10 != v28) {
                                                            _ = len(v5)
                                                        }
                                                        if v23 {
                                                            v34 := GenS2{F0: -122, F1: 127}
                                                            _ = v34
                                                        } else {
                                                            v35 := GenS4{F0: (((("hello" + "foo bar") + (v29 + ("hello" + (("" + "hello") + v29)))) + "") + v29), F1: v33, F2: 42}
                                                            if true {
                                                                var v36 uint16 = 32767
                                                                p0 -= p0
                                                                _ = v36
                                                            } else {
                                                                switch ("" + "a") {
                                                                case "a":
                                                                    var v37 int = (^(fi6))
                                                                    _ = v37
                                                                case "":
                                                                    _ = len(v5)
                                                                }
                                                            }
                                                            var v38 uint32 = 42
                                                            _ = v35
                                                            _ = v38
                                                        }
                                                        v5 = append(v5, (fi14 * 42))
                                                        switch v29 {
                                                        case "foo bar":
                                                            if false {
                                                                v17 -= v17
                                                                v33 += 1000
                                                            }
                                                        case "a":
                                                            if (218 < 2) {
                                                                v5 = append(v5, fi14)
                                                                if (-934 < (fi6 ^ (-(fi14)))) {
                                                                    v33 += 603
                                                                }
                                                            } else {
                                                                p0 += p0
                                                                v13 = 42
                                                                var v39 int64 = 1
                                                                if v23 {
                                                                    if ((^(fi14)) >= -331) {
                                                                        var v40 uint8 = p0
                                                                        _ = v40
                                                                    } else {
                                                                        v23 = (42 <= 4294967295)
                                                                        p0 -= (p0 / (((^(p0))) | 1))
                                                                    }
                                                                    fmt.Println(v29, v33, v23)
                                                                    _ = len(v5)
                                                                    var v41 int = (-(fi6))
                                                                    var v42 float64 = float64(v28)
                                                                    _ = v41
                                                                    _ = v42
                                                                }
                                                                _ = v39
                                                            }
                                                        }
                                                        _ = v33
                                                    }
                                                    _ = v29
                                                case "":
                                                    v43 := GenS5{F0: 406}
                                                    _ = v43
                                                default:
                                                    if (0 <= 65535) {
                                                        var v44 int32 = 0
                                                        _ = v44
                                                    }
                                                }
                                                v45 := GenS6{F0: 4294967295}
                                                _ = v45
                                            }
                                            _ = v23
                                            _ = v28
                                        } else {
                                            var v46 int8 = 72
                                            _ = v46
                                        }
                                    default:
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                case -295:
                                    var v47 int = fi14
                                    _ = v47
                                default:
                                    v48 := GenS7{F0: uint16(p0)}
                                    _ = v48
                                }
                                p0 += p0
                                _ = v17
                            }
                            v13 += (v13 ^ v13)
                            _ = fi14
                        }
                        _ = v13
                    } else {
                        for fi49 := 0; fi49 < 11; fi49++ {
                            if true {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                            _ = fi49
                        }
                        if (42 <= 1) {
                            var v50 int64 = 1
                            _ = v50
                        }
                    }
                    p0 += p0
                }
            }
            if true {
                var v51 string = ("a" + "foo bar")
                switch (v51 + "hello") {
                case "foo bar":
                    fmt.Println(v51, fi6, fi6)
                default:
                    var v52 int8 = 1
                    _ = v52
                }
                var v53 float64 = 1e10
                _ = v51
                _ = v53
            } else {
                var v54 int8 = 100
                fmt.Println(fi6)
                _ = v54
            }
        }
        _ = fi6
        _ = v7
    }
    var v55 float32 = -2.25
    _ = p0
    _ = v5
    _ = v55
    return 943
}

func main() {
    v1 := []uint32{2147483648, 2, 65535, 708}
    var v2 uint = 855
    v3 := make(map[string]string)
    var v4 int64 = 9223372036854775807
    v2 += 194
    if ((float32(1e10) / 0.1) <= 0.0) {
        var v56 uint32 = 2147483648
        var v57 string = ("foo bar" + ("hello" + "foo bar"))
        if (v57 > ("" + v57)) {
            v2 -= v2
            var v58 uint64 = 18446744073709551615
            fmt.Println(v56, v56)
            _ = v58
        } else {
            v2 = (v2 ^ v2)
            delete(v3, ((("foo bar" + v57) + ("hello" + (("\n\t" + v57) + (v57 + v57)))) + (v57 + "foo bar")))
            for fi59 := 0; fi59 < 4; fi59++ {
                for fi60 := 0; fi60 < 19; fi60++ {
                    v1 = append(v1, (v56 - (v56 >> ((v2) & 63))))
                    var v61 string = v57
                    v4 += v4
                    _ = fi60
                    _ = v61
                }
                _ = fi59
            }
        }
        _ = v56
        _ = v57
    }
    switch 127 {
    case 32768:
        _ = len(v1)
    default:
        _ = len(v3)
    }
    var v62 float32 = (((1e10 * 0.1) * 3.14159) + 1e10)
    v4 = (^(v4))
    fmt.Println(v2)
    v2 += 1
    v4 = (^(v4))
    v2 = (v2 & (v2 + 188))
    var v63 int8 = -128
    var v64 float64 = (-2.25 + (0.0 + 1.5))
    v62 += v62
    var v65 uint64 = 826
    if false {
        v66 := GenS8{F0: "hello", F1: 255}
        for fi67 := 0; fi67 < 4; fi67++ {
            if (v4 <= v4) {
                v62 = (v62 + v62)
                v62 += (v62 / 3.14159)
                fmt.Println(v4, v64, v63)
            } else {
                if ((fi67 % ((fi67) | 1)) >= -2147483648) {
                    var v68 int16 = -511
                    _ = v68
                } else {
                    var v69 uint64 = 2
                    v70 := GenS9{F0: 624, F1: (42 <= 1)}
                    var v71 uint16 = 2
                    _ = v69
                    _ = v70
                    _ = v71
                }
                var v72 uint = (v2 >> ((42) & 63))
                v63 = (^(v63))
                _ = v72
            }
            if ((-(v63)) == (v63 >> ((v2) & 63))) {
                v73 := GenS7{F0: 534}
                var v74 uint8 = 2
                switch "" {
                case "\n\t":
                    if true {
                        v2 += uint(v74)
                    } else {
                        if ((v64 + float64(v62)) < v64) {
                            for fi75 := 0; fi75 < 14; fi75++ {
                                v76 := v3[("" + "hello")]
                                v2 += (^(v2))
                                _ = fi75
                                _ = v76
                            }
                        } else {
                            fmt.Println(v4, v74)
                            v64 -= v64
                            if (fi67 != -2147483648) {
                                var v77 int = (fi67 % ((42) | 1))
                                v4 += int64(v63)
                                v77 += (fi67 + (fi67 * 2147483647))
                                _ = v77
                            }
                        }
                        fmt.Println(v2, v2)
                        switch ((((("" + "hello") + "") + ((("a" + (("hello" + ("hello" + "foo bar")) + "a")) + ("foo bar" + "")) + "hello")) + ("a" + "\n\t")) + (("" + "foo bar") + "\n\t")) {
                        case "hello":
                            _ = len(v1)
                        case "a":
                            v2 -= (v2 >> ((v2) & 63))
                        case "":
                            if true {
                                _ = len(v1)
                                for fi78 := 0; fi78 < 13; fi78++ {
                                    if ((v2 & 520) < (v2 ^ (^(v2)))) {
                                        if len(v1) > 0 {
                                            v1[0] = 104
                                        }
                                        v64 += v64
                                    }
                                    if true {
                                        var v79 int8 = 42
                                        fmt.Println(fi78, v62)
                                        var v80 uint8 = 255
                                        _ = v79
                                        _ = v80
                                    } else {
                                        if len(v1) > 0 {
                                            v1[0] = 2147483647
                                        }
                                        var v81 uint64 = v65
                                        _ = v81
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = 599
                                    }
                                    _ = fi78
                                }
                            }
                        }
                    }
                case "foo bar":
                    if false {
                        if false {
                            fmt.Println(fi67, v62)
                            var v82 string = ("" + "hello")
                            _ = v82
                        } else {
                            for fi83 := 0; fi83 < 3; fi83++ {
                                v3[""] = "a"
                                if true {
                                    fmt.Println(fi67, fi67)
                                    var v84 float32 = (v62 * (v62 + -2.25))
                                    var v85 int = (fi83 + fi67)
                                    v86 := GenS10{F0: 73, F1: genFn0((v74 / ((10) | 1))), F2: (!(!true))}
                                    _ = v84
                                    _ = v85
                                    _ = v86
                                } else {
                                    v87 := GenS11{F0: 0, F1: (v2 ^ (v2 * 4294967295))}
                                    switch fi83 {
                                    case 932:
                                        if (float64(v62) >= -2.25) {
                                            if (-1 < int16(v2)) {
                                                var v88 uint64 = v65
                                                switch ("a" + "hello") {
                                                case "\n\t":
                                                    if len(v1) > 0 {
                                                        v1[0] = 2147483647
                                                    }
                                                case "foo bar":
                                                    _ = len(v3)
                                                case "a":
                                                    if true {
                                                        if (!(-2147483648 < -436)) {
                                                            var v89 uint32 = 4294967295
                                                            _ = v89
                                                        } else {
                                                            v88 -= v88
                                                            v62 = ((v62 - v62) + v62)
                                                        }
                                                        if true {
                                                            v90 := GenS5{F0: v4}
                                                            var v91 uint8 = uint8(v2)
                                                            _ = v90
                                                            _ = v91
                                                        }
                                                        v64 = float64((3.14159 + float32(float64((1.5 - (v62 - -2.25))))))
                                                    } else {
                                                        var v92 float64 = v64
                                                        switch "hello" {
                                                        case "\n\t":
                                                            v88 = v88
                                                        case "hello":
                                                            if len(v1) > 0 {
                                                                v1[0] = 2147483647
                                                            }
                                                        case "":
                                                            if (42 >= v4) {
                                                                if true {
                                                                    v92 += (v64 * ((-2.25 + v64) / float64(v62)))
                                                                    var v93 uint = 42
                                                                    _ = v93
                                                                } else {
                                                                    if (false && (v4 > (v4 ^ v4))) {
                                                                        v63 = (v63 ^ -49)
                                                                        fmt.Println(v63, v2, v2)
                                                                    }
                                                                    v92 = -2.25
                                                                    var v94 int8 = (v63 | -16)
                                                                    _ = v94
                                                                }
                                                            }
                                                        }
                                                        v95 := v3[("hello" + (("foo bar" + "") + "foo bar"))]
                                                        if false {
                                                            switch v95 {
                                                            case "foo bar":
                                                                fmt.Println(v4, v92)
                                                            case "\n\t":
                                                                if (657 >= 2147483647) {
                                                                    var v96 uint16 = 65535
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        fmt.Println(fi67)
                                                                    default:
                                                                        v92 -= 3.14159
                                                                    }
                                                                    _ = v96
                                                                } else {
                                                                    v63 = 117
                                                                }
                                                            case "":
                                                                v88 = (^(v65))
                                                            }
                                                            switch (fi67 * fi83) {
                                                            case -108:
                                                                var v97 uint8 = (v74 % ((127) | 1))
                                                                _ = v97
                                                            case 0:
                                                                v98 := GenS10{F0: (v74 << ((813) & 63)), F1: 430, F2: ((v65 + (v88 >> ((uint(v74)) & 63))) != (v88 + (^(v65))))}
                                                                _ = v98
                                                            case -758:
                                                                v95 = ("a" + v95)
                                                            default:
                                                                switch int(v63) {
                                                                case -2147483648:
                                                                    _ = len(v1)
                                                                case -993:
                                                                    if true {
                                                                        fmt.Println(v64, v74, v88)
                                                                        v65 = v88
                                                                        v99 := GenS8{F0: v95, F1: uint16(v74)}
                                                                        _ = v99
                                                                    } else {
                                                                        fmt.Println(v74, v95, v95)
                                                                        v64 = v64
                                                                        v64 = v92
                                                                    }
                                                                case -204:
                                                                    var v100 uint8 = (v74 * uint8(v2))
                                                                    _ = v100
                                                                default:
                                                                    switch v95 {
                                                                    case "":
                                                                        if (false && (v88 != v65)) {
                                                                            var v101 int16 = 396
                                                                            v1 = append(v1, 42)
                                                                            _ = v101
                                                                        } else {
                                                                            var v102 float32 = (v62 - v62)
                                                                            v103 := GenS12{F0: 2}
                                                                            var v104 float64 = 1.5
                                                                            v64 -= v64
                                                                            fmt.Println(v74)
                                                                            _ = v102
                                                                            _ = v103
                                                                            _ = v104
                                                                        }
                                                                    case "hello":
                                                                        v2 -= v2
                                                                    }
                                                                }
                                                            }
                                                            var v105 uint8 = v74
                                                            _ = v105
                                                        }
                                                        _ = len(v3)
                                                        _ = v92
                                                        _ = v95
                                                    }
                                                }
                                                _ = v88
                                            } else {
                                                fmt.Println(v64)
                                            }
                                            if true {
                                                var v106 float32 = ((0.0 / v62) / v62)
                                                v65 = 16
                                                _ = v106
                                            } else {
                                                if len(v1) > 0 {
                                                    v1[0] = 2147483647
                                                }
                                                v63 += 127
                                            }
                                        } else {
                                            switch fi83 {
                                            case 32768:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            case -2147483648:
                                                switch "\n\t" {
                                                case "":
                                                    var v107 uint8 = (v74 % (((v74 >> (((^(v2))) & 63))) | 1))
                                                    _ = v107
                                                case "\n\t":
                                                    v65 -= v65
                                                case "hello":
                                                    v2 += uint(v65)
                                                }
                                                if ((-2.25 + float32(v64)) >= (1.5 / 1e10)) {
                                                    switch (fi67 * fi83) {
                                                    case 128:
                                                        if (fi83 != fi67) {
                                                            if (false || (42 >= 0)) {
                                                                var v108 string = ("hello" + ((("\n\t" + ("a" + ("\n\t" + "hello"))) + "foo bar") + ""))
                                                                _ = v108
                                                            } else {
                                                                switch (fi83 ^ fi67) {
                                                                case -918:
                                                                    switch 42 {
                                                                    case 1:
                                                                        v109 := GenS13{F0: (v2 + 2)}
                                                                        _ = v109
                                                                    case -871:
                                                                        switch fi67 {
                                                                        case 587:
                                                                            if false {
                                                                                if false {
                                                                                    fmt.Println(fi83)
                                                                                    v63 -= (v63 % ((v63) | 1))
                                                                                }
                                                                            }
                                                                        case -1:
                                                                            v110 := GenS14{F0: (("foo bar" + "") + "foo bar")}
                                                                            _ = v110
                                                                        default:
                                                                            switch "hello" {
                                                                            case "hello":
                                                                                var v111 uint = 0
                                                                                _ = v111
                                                                            case "a":
                                                                                v112 := GenS12{F0: 2147483647}
                                                                                _ = v112
                                                                            case "":
                                                                                if (false && true) {
                                                                                    var v113 uint8 = (^(v74))
                                                                                    _ = v113
                                                                                }
                                                                            default:
                                                                                if (4294967295 >= v65) {
                                                                                    v2 = v2
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                case 360:
                                                                    var v114 uint64 = (v65 + v65)
                                                                    _ = v114
                                                                case 32767:
                                                                    var v115 int8 = (v63 & (v63 / (((v63 ^ v63)) | 1)))
                                                                    _ = v115
                                                                default:
                                                                    v63 = (v63 | (v63 + 44))
                                                                }
                                                            }
                                                            v62 = v62
                                                        } else {
                                                            switch "" {
                                                            case "foo bar":
                                                                fmt.Println(v4, fi83, v62)
                                                            case "":
                                                                switch "" {
                                                                case "a":
                                                                    v1 = append(v1, 4294967295)
                                                                default:
                                                                    var v116 float32 = (v62 + v62)
                                                                    _ = v116
                                                                }
                                                            case "a":
                                                                delete(v3, "")
                                                            }
                                                        }
                                                    case 32768:
                                                        v4 += 1
                                                    case -1:
                                                        fmt.Println(v62, v4, v74)
                                                        if true {
                                                            var v117 uint8 = 2
                                                            _ = v117
                                                        }
                                                    default:
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    }
                                                } else {
                                                    var v118 uint64 = (v65 % (((v65 - (v65 | (v65 << ((110) & 63))))) | 1))
                                                    _ = v118
                                                }
                                            }
                                            if (v74 == v74) {
                                                var v119 float32 = (((v62 + v62) / (float32(v64) / (v62 / v62))) * (v62 / 0.1))
                                                v120 := GenS2{F0: v63, F1: v63}
                                                _ = v119
                                                _ = v120
                                            }
                                            v65 = 2
                                        }
                                    case 4294967295:
                                        var v121 float64 = 0.0
                                        _ = v121
                                    default:
                                        var v122 float32 = 3.14159
                                        _ = v122
                                    }
                                    if false {
                                        v123 := GenS9{F0: 2147483648, F1: (false || false)}
                                        var v124 int8 = (^(v63))
                                        _ = v123
                                        _ = v124
                                    } else {
                                        if (32767 >= 1000) {
                                            v125 := GenS15{F0: true, F1: v4, F2: 197}
                                            _ = v125
                                        }
                                        v126 := GenS16{F0: ("" + "a"), F1: 0}
                                        _ = v126
                                    }
                                    _ = v87
                                }
                                v127 := GenS17{F0: "foo bar", F1: v63, F2: v74}
                                _ = fi83
                                _ = v127
                            }
                            var v128 int64 = 42
                            v129 := GenS17{F0: ("foo bar" + "hello"), F1: v63, F2: v74}
                            _ = v128
                            _ = v129
                        }
                    } else {
                        if (((v74 >> ((v2) & 63)) < (^(v74))) && (fi67 <= (fi67 >> (((^(v2))) & 63)))) {
                            v62 -= (0.0 / float32(v64))
                            v130 := GenS14{F0: ("foo bar" + "")}
                            _ = v130
                        }
                        var v131 uint64 = v65
                        v4 = v4
                        _ = v131
                    }
                case "a":
                    v132 := v3[("hello" + "\n\t")]
                    _ = v132
                }
                _ = v73
                _ = v74
            } else {
                if ("" > "foo bar") {
                    for fi133 := 0; fi133 < 19; fi133++ {
                        var v134 int64 = v4
                        var v135 uint8 = 127
                        _ = fi133
                        _ = v134
                        _ = v135
                    }
                    if (1 > 1000) {
                        var v136 uint16 = 42
                        if false {
                            var v137 int32 = -1
                            _ = v137
                        } else {
                            v62 = 0.1
                        }
                        var v138 int64 = (v4 / (((v4 * -949)) | 1))
                        _ = v136
                        _ = v138
                    } else {
                        v65 = (^(v65))
                    }
                    for fi139 := 0; fi139 < 12; fi139++ {
                        fmt.Println(v64)
                        switch fi67 {
                        case 127:
                            v62 = v62
                        case 42:
                            switch "" {
                            case "foo bar":
                                fmt.Println(v64, v2)
                            case "a":
                                if ("hello" == ("\n\t" + "foo bar")) {
                                    v140 := GenS16{F0: ("hello" + ("foo bar" + "hello")), F1: 0}
                                    if true {
                                        switch 2147483647 {
                                        case -9223372036854775808:
                                            v141 := GenS13{F0: (v2 ^ v2)}
                                            _ = v141
                                        case 32768:
                                            v62 -= v62
                                        }
                                    } else {
                                        switch "" {
                                        case "a":
                                            switch ("foo bar" + "a") {
                                            case "":
                                                fmt.Println(fi67, fi139, v64)
                                            }
                                        case "foo bar":
                                            delete(v3, "\n\t")
                                            v142 := GenS10{F0: uint8(v65), F1: genFn0(205), F2: false}
                                            v65 -= v65
                                            _ = v142
                                        case "":
                                            var v143 float32 = v62
                                            _ = v143
                                        default:
                                            v64 = 1e10
                                        }
                                        v65 = (v65 + v65)
                                        if (v2 <= (v2 % (((v2 | 65535)) | 1))) {
                                            var v144 int = fi139
                                            var v145 uint8 = 1
                                            _ = v144
                                            _ = v145
                                        } else {
                                            v146 := GenS18{F0: 100, F1: "a"}
                                            v147 := GenS8{F0: (("a" + "") + ("foo bar" + (("\n\t" + "") + "a"))), F1: 900}
                                            var v148 int = -2147483648
                                            v64 = (v64 + ((v64 / 3.14159) / v64))
                                            if (v63 == -122) {
                                                v149 := GenS15{F0: false, F1: v4, F2: 141}
                                                var v150 int64 = 0
                                                _ = v149
                                                _ = v150
                                            }
                                            _ = v146
                                            _ = v147
                                            _ = v148
                                        }
                                    }
                                    _ = v140
                                }
                            case "hello":
                                switch (fi67 ^ 127) {
                                case -137:
                                    var v151 int = int(v2)
                                    _ = v151
                                case -1:
                                    var v152 uint8 = 45
                                    _ = v152
                                case 4294967295:
                                    v153 := GenS19{F0: v2}
                                    _ = v153
                                }
                            }
                        case 9223372036854775807:
                            v154 := GenS20{F0: (v65 | 153), F1: 32767, F2: (int32(v2) <= 1)}
                            _ = v154
                        default:
                            fmt.Println(v65, v65, v4)
                        }
                        var v155 float32 = v62
                        _ = fi139
                        _ = v155
                    }
                } else {
                    if false {
                        var v156 int32 = int32(v63)
                        switch "hello" {
                        case "hello":
                            switch (fi67 * fi67) {
                            case 127:
                                v4 = (-(v4))
                            case -1:
                                fmt.Println(v63, v62, v62)
                            default:
                                fmt.Println(v156, v63, v64)
                            }
                        default:
                            if true {
                                v157 := GenS5{F0: v4}
                                v158 := GenS2{F0: (v63 >> ((65535) & 63)), F1: v63}
                                if false {
                                    if (v4 >= (v4 | (v4 ^ v4))) {
                                        fmt.Println(v156, fi67)
                                    } else {
                                        for fi159 := 0; fi159 < 20; fi159++ {
                                            var v160 int16 = 32767
                                            _ = fi159
                                            _ = v160
                                        }
                                        fmt.Println(v65)
                                    }
                                }
                                _ = v157
                                _ = v158
                            } else {
                                v62 -= ((-2.25 - (v62 * (v62 - v62))) + v62)
                            }
                        }
                        _ = v156
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    if false {
                        v161 := GenS21{F0: (v4 ^ (v4 % ((-9223372036854775808) | 1)))}
                        for fi162 := 0; fi162 < 1; fi162++ {
                            switch "foo bar" {
                            case "a":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            case "":
                                v1 = append(v1, 160)
                            default:
                                v63 = v63
                                if (-9223372036854775808 == (^(v4))) {
                                    v4 += (v4 | (v4 + -1))
                                } else {
                                    v163 := GenS11{F0: 0, F1: (v2 ^ v2)}
                                    var v164 string = (("hello" + "") + "hello")
                                    var v165 int16 = 11
                                    _ = v163
                                    _ = v164
                                    _ = v165
                                }
                            }
                            _ = fi162
                        }
                        _ = v161
                    } else {
                        var v166 string = ("" + "foo bar")
                        v167 := GenS10{F0: 255, F1: 65535, F2: (255 < 65535)}
                        _ = v166
                        _ = v167
                    }
                    v65 = (v65 | v65)
                }
            }
            var v168 uint64 = (v65 << (((v2 - (v2 + 439))) & 63))
            _ = fi67
            _ = v168
        }
        _ = v66
    } else {
        for fi169 := 0; fi169 < 10; fi169++ {
            v2 += (v2 - v2)
            _ = fi169
        }
        v62 = (v62 - 1e10)
        switch -1 {
        case 0:
            fmt.Println(v64)
        case 866:
            var v170 string = "hello"
            _ = v170
        default:
            var v171 int16 = 1000
            _ = v171
        }
    }
    v172 := GenS1{F0: (("" + "a") + "a"), F1: (v2 << ((v2) & 63))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v172
    _ = v2
    _ = v3
    _ = v4
    _ = v62
    _ = v63
    _ = v64
    _ = v65
}
