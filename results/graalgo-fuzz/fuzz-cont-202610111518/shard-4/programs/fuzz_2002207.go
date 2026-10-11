package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 uint }
type GenS5 struct { F0 uint64; F1 int64 }
type GenS6 struct { F0 uint8; F1 uint8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint64; F1 int32 }
type GenS9 struct { F0 int; F1 uint8; F2 uint32 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 string }
type GenS12 struct { F0 string; F1 int64; F2 uint }
type GenS13 struct { F0 int16; F1 bool }
type GenS14 struct { F0 int64; F1 string }
type GenS15 struct { F0 int; F1 uint32; F2 uint32 }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 int32; F1 int16; F2 int32 }
type GenS18 struct { F0 bool; F1 uint32 }
type GenS19 struct { F0 uint }
type GenS20 struct { F0 uint64 }
type GenS21 struct { F0 int64; F1 string }
type GenS22 struct { F0 uint16; F1 uint64; F2 uint32 }
type GenS23 struct { F0 uint8; F1 int16; F2 int32 }

func main() {
    var v1 uint16 = 65535
    var v2 int32 = int32(v1)
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]int)
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 uint = 1
        _ = fi5
        _ = v6
    }
    var v7 bool = false
    v7 = ((v1 - 511) >= 65535)
    var v8 uint8 = 26
    v9 := []int{}
    var v10 uint8 = (v8 * (v8 << ((18446744073709551615) & 63)))
    v7 = true
    v2 -= int32(v1)
    if v7 {
        var v11 int32 = (v2 / ((int32(v10)) | 1))
        if v7 {
            fmt.Println(v1)
        }
        _ = v11
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        switch 190 {
        case 0:
            v13 := GenS1{F0: v7, F1: v7}
            _ = v13
        case -1:
            var v14 uint = 594
            _ = v14
        case 353:
            for fi15 := 0; fi15 < 11; fi15++ {
                var v16 uint64 = 18446744073709551615
                _ = fi15
                _ = v16
            }
        }
        v17 := GenS1{F0: (1 > 42), F1: v7}
        if false {
            fmt.Println(v10)
            v18 := GenS1{F0: false, F1: (1 >= -1)}
            v19 := GenS1{F0: true, F1: (v7 || true)}
            _ = v18
            _ = v19
        } else {
            if v7 {
                var v20 uint32 = 2147483647
                var v21 int64 = 1
                _ = v20
                _ = v21
            } else {
                fmt.Println(v8, v7)
                _ = len(v9)
                v22 := GenS1{F0: (((!((!(2147483648 <= 0)) || (255 > (v1 / ((v1) | 1))))) || false) || ((fi12 ^ (fi12 << ((4294967295) & 63))) == (^(fi12)))), F1: v7}
                _ = v22
            }
            switch "foo bar" {
            case "hello":
                v8 = (v8 - v10)
            case "":
                for fi23 := 0; fi23 < 15; fi23++ {
                    var v24 int32 = v2
                    v24 -= v24
                    var v25 int8 = -128
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            case "a":
                fmt.Println(v2, v10)
            default:
                fmt.Println(v1, v10)
            }
        }
        _ = fi12
        _ = v17
    }
    switch 4294967295 {
    case -9223372036854775808:
        if false {
            v4[625] = 128
        } else {
            for fi26 := 0; fi26 < 15; fi26++ {
                if (v7 && (false && true)) {
                    switch "\n\t" {
                    case "a":
                        if true {
                            v7 = v7
                        }
                    case "\n\t":
                        if (!(1e10 == (((3.14159 - 1e10) * float32(0.0)) / (0.1 + 0.1)))) {
                            v27 := v4[42]
                            _ = v27
                        } else {
                            v28 := GenS2{F0: 295}
                            v1 = v1
                            _ = v28
                        }
                    }
                } else {
                    v1 = v1
                    if (!((2147483648 < 997) || (v7 || (0 != 100)))) {
                        v4[2147483647] = fi26
                        if (v1 != v1) {
                            if (v10 <= 255) {
                                if len(v9) > 0 {
                                    v9[0] = fi26
                                }
                            }
                            if (v7 && (!(!v7))) {
                                if v7 {
                                    v29 := GenS3{F0: 255}
                                    _ = v29
                                } else {
                                    var v30 uint16 = (^(v1))
                                    if true {
                                        delete(v4, 42)
                                        if (2147483647 == 1) {
                                            if (9223372036854775807 >= 0) {
                                                v30 -= (v30 & 255)
                                                if (!(18446744073709551615 != 42)) {
                                                    v31 := GenS1{F0: ((v30 > (v30 / ((209) | 1))) && ((242 <= 50) || (v7 && (0.0 < (1e10 * float64((-2.25 / 1.0))))))), F1: v7}
                                                    var v32 int8 = 42
                                                    delete(v3, 65535)
                                                    _ = v31
                                                    _ = v32
                                                } else {
                                                    if (((3.14159 - 3.14159) == 0.0) || (uint32(v1) >= 494)) {
                                                        fmt.Println(v30)
                                                    }
                                                }
                                            } else {
                                                v33 := GenS2{F0: 65535}
                                                v34 := v3[uint32(v1)]
                                                v30 += (v1 + v30)
                                                _ = v33
                                                _ = v34
                                            }
                                            _ = len(v9)
                                        } else {
                                            var v35 uint8 = (v10 >> ((18446744073709551615) & 63))
                                            var v36 uint = 2
                                            _ = v35
                                            _ = v36
                                        }
                                        if ("hello" < "foo bar") {
                                            fmt.Println(fi26, v1, v8)
                                        }
                                    }
                                    if true {
                                        if true {
                                            var v37 int64 = 42
                                            _ = v37
                                        }
                                        if true {
                                            for fi38 := 0; fi38 < 11; fi38++ {
                                                var v39 bool = ((!(-2.25 <= -2.25)) || v7)
                                                switch 755 {
                                                case 974:
                                                    var v40 bool = v39
                                                    _ = v40
                                                case 4294967295:
                                                    var v41 int32 = 42
                                                    _ = v41
                                                case 2147483647:
                                                    if (v39 && (494 < (^(v2)))) {
                                                        v8 -= v8
                                                        v8 = v10
                                                        var v42 int = fi38
                                                        _ = v42
                                                    }
                                                }
                                                if (!true) {
                                                    v10 += v10
                                                    var v43 uint8 = (^(v8))
                                                    if (!(1 > v10)) {
                                                        var v44 int64 = -9223372036854775808
                                                        var v45 int8 = 2
                                                        fmt.Println(v30, v43, fi26)
                                                        _ = v44
                                                        _ = v45
                                                    }
                                                    _ = v43
                                                }
                                                _ = fi38
                                                _ = v39
                                            }
                                            v2 = v2
                                        } else {
                                            v46 := v3[2147483647]
                                            v47 := GenS4{F0: 255, F1: 462}
                                            _ = v46
                                            _ = v47
                                        }
                                        for fi48 := 0; fi48 < 12; fi48++ {
                                            if (v8 != v8) {
                                                switch fi26 {
                                                case 9223372036854775807:
                                                    delete(v4, 42)
                                                default:
                                                    switch (("a" + ((("" + ("\n\t" + "a")) + "\n\t") + ((("hello" + "foo bar") + ("" + "hello")) + ("\n\t" + ("" + "hello"))))) + "") {
                                                    case "hello":
                                                        v1 += v30
                                                    case "":
                                                        switch ("hello" + ("foo bar" + ((("foo bar" + "hello") + "") + "foo bar"))) {
                                                        case "\n\t":
                                                            v30 = (v30 / ((0) | 1))
                                                        case "hello":
                                                            if true {
                                                                v30 -= 65535
                                                                var v49 uint64 = 1
                                                                fmt.Println(v10, fi48, v2)
                                                                _ = v49
                                                            }
                                                        case "":
                                                            delete(v4, 2147483647)
                                                        default:
                                                            fmt.Println(v30, v1)
                                                            fmt.Println(v10, v2, v7)
                                                            var v50 bool = v7
                                                            if v7 {
                                                                switch (fi48 | fi26) {
                                                                case 114:
                                                                    v9 = append(v9, fi48)
                                                                default:
                                                                    v51 := GenS3{F0: 1}
                                                                    _ = v51
                                                                }
                                                            } else {
                                                                if len(v9) > 0 {
                                                                    v9[0] = fi26
                                                                }
                                                                fmt.Println(v2, v7)
                                                            }
                                                            var v52 uint = 318
                                                            v53 := GenS2{F0: (v1 << (((v52 / ((v52) | 1))) & 63))}
                                                            v54 := GenS1{F0: (v52 != (^(v52))), F1: false}
                                                            var v55 int8 = 0
                                                            var v56 int16 = 1000
                                                            _ = v50
                                                            _ = v52
                                                            _ = v53
                                                            _ = v54
                                                            _ = v55
                                                            _ = v56
                                                        }
                                                        var v57 float32 = (-2.25 * 0.0)
                                                        _ = v57
                                                    default:
                                                        v58 := GenS2{F0: (v30 ^ (v30 | v1))}
                                                        _ = v58
                                                    }
                                                }
                                                v8 = 255
                                            }
                                            _ = fi48
                                        }
                                    } else {
                                        delete(v3, 2)
                                    }
                                    _ = v30
                                }
                            }
                        }
                        var v59 int = (fi26 + fi26)
                        _ = v59
                    } else {
                        if true {
                            if len(v9) > 0 {
                                v9[0] = 32768
                            }
                        }
                        v60 := GenS5{F0: 687, F1: -9223372036854775808}
                        fmt.Println(v7, v8, v7)
                        _ = v60
                    }
                    for fi61 := 0; fi61 < 6; fi61++ {
                        if (!v7) {
                            var v62 int = (fi26 + 32767)
                            var v63 int32 = (v2 - 0)
                            v63 -= v63
                            _ = v62
                            _ = v63
                        } else {
                            v2 = 0
                            v64 := GenS3{F0: uint(v8)}
                            switch 1 {
                            case 146:
                                v65 := GenS6{F0: (v10 - (v8 % (((v8 * 255)) | 1))), F1: v10}
                                _ = v65
                            case 127:
                                switch (-(fi61)) {
                                case 1:
                                    v9 = append(v9, (fi26 ^ fi61))
                                case -1:
                                    v2 -= -575
                                }
                            default:
                                if (!(!v7)) {
                                    var v66 uint64 = 849
                                    switch (("\n\t" + "foo bar") + "foo bar") {
                                    case "foo bar":
                                        fmt.Println(v66, fi26, fi61)
                                    case "a":
                                        v7 = (!v7)
                                    case "hello":
                                        v66 -= (^(v66))
                                    default:
                                        var v67 uint32 = 2147483647
                                        _ = v67
                                    }
                                    _ = v66
                                }
                            }
                            switch fi61 {
                            case 4294967295:
                                v9 = append(v9, fi26)
                            case -2147483648:
                                if (0 != 1) {
                                    switch "foo bar" {
                                    case "a":
                                        switch "a" {
                                        case "foo bar":
                                            switch (fi61 / (((fi61 % (((fi61 % ((fi61) | 1))) | 1))) | 1)) {
                                            case 127:
                                                v8 = (v8 / ((1) | 1))
                                            case 9223372036854775807:
                                                if len(v9) > 0 {
                                                    v9[0] = 22
                                                }
                                            }
                                        case "\n\t":
                                            if (!(0 < -788)) {
                                                var v68 float32 = float32(3.14159)
                                                fmt.Println(v2, v7, v8)
                                                v69 := v3[0]
                                                _ = v68
                                                _ = v69
                                            } else {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    v7 = true
                                                case "foo bar":
                                                    var v70 uint64 = 4294967295
                                                    _ = v70
                                                case "hello":
                                                    var v71 int64 = 0
                                                    _ = v71
                                                default:
                                                    if false {
                                                        if false {
                                                            delete(v3, 2)
                                                            v72 := GenS5{F0: 1, F1: -1}
                                                            if v7 {
                                                                v8 = v8
                                                                v10 = (v10 % ((v8) | 1))
                                                            }
                                                            _ = v72
                                                        } else {
                                                            if (!v7) {
                                                                var v73 uint32 = 2147483648
                                                                v10 += v8
                                                                _ = v73
                                                            }
                                                            switch fi26 {
                                                            case 128:
                                                                var v74 uint = 65535
                                                                _ = v74
                                                            case -298:
                                                                v75 := GenS1{F0: (!v7), F1: false}
                                                                _ = v75
                                                            case 42:
                                                                v1 += 352
                                                            }
                                                            switch "hello" {
                                                            case "hello":
                                                                v76 := v3[4294967295]
                                                                v77 := GenS4{F0: v1, F1: 2}
                                                                _ = v76
                                                                _ = v77
                                                            case "\n\t":
                                                                v7 = ((((3.14159 + ((0.1 * 1.5) - float64(-2.25))) - 1.5) / 1.5) >= 1.5)
                                                                var v78 uint8 = (^(v8))
                                                                _ = v78
                                                            case "foo bar":
                                                                _ = len(v9)
                                                            }
                                                        }
                                                        var v79 int64 = 0
                                                        var v80 int = fi61
                                                        _ = v79
                                                        _ = v80
                                                    } else {
                                                        v8 = 1
                                                        if ((!v7) && v7) {
                                                            v8 = 2
                                                            v8 = (^(v8))
                                                        }
                                                        v81 := GenS7{F0: v7}
                                                        _ = v81
                                                    }
                                                }
                                                v2 = 2147483647
                                                if len(v9) >= 2 {
                                                    v9 = v9[1:len(v9)]
                                                }
                                            }
                                        default:
                                            if v7 {
                                                var v82 int32 = (v2 * v2)
                                                _ = v82
                                            }
                                        }
                                    case "\n\t":
                                        fmt.Println(v1, v7, v8)
                                        v7 = v7
                                    case "":
                                        if v7 {
                                            fmt.Println(v8, fi26, fi61)
                                            switch (^(fi61)) {
                                            case 32768:
                                                if false {
                                                    v10 = (^(v8))
                                                    _ = len(v4)
                                                    var v83 string = ""
                                                    _ = v83
                                                } else {
                                                    v8 = (v10 + v10)
                                                    v1 = v1
                                                    _ = len(v4)
                                                }
                                                if ((fi26 & (-(fi26))) >= (fi26 << ((795) & 63))) {
                                                    if ((v1 - (v1 + 32767)) < 92) {
                                                        var v84 uint64 = 2
                                                        fmt.Println(v2, v1, v8)
                                                        v9 = append(v9, fi61)
                                                        _ = v84
                                                    }
                                                    var v85 int = 127
                                                    switch "hello" {
                                                    case "":
                                                        if (484 > 526) {
                                                            v8 += (v8 / ((v10) | 1))
                                                            v86 := GenS8{F0: 256, F1: (v2 >> ((4294967295) & 63))}
                                                            var v87 int8 = -128
                                                            _ = v86
                                                            _ = v87
                                                        }
                                                    case "hello":
                                                        fmt.Println(v7)
                                                        v85 -= fi61
                                                        v85 = v85
                                                    default:
                                                        v8 += (v10 - v8)
                                                    }
                                                    _ = v85
                                                }
                                            case -492:
                                                v2 = (v2 << ((0) & 63))
                                            }
                                        }
                                    }
                                    if false {
                                        delete(v3, 2147483648)
                                        var v88 uint = 169
                                        _ = v88
                                    }
                                    switch (fi61 | (^(fi26))) {
                                    case 1:
                                        var v89 int = fi61
                                        switch fi26 {
                                        case -2147483648:
                                            switch ("hello" + ("" + ("foo bar" + "hello"))) {
                                            case "\n\t":
                                                var v90 int16 = 32767
                                                v90 = (v90 % ((v90) | 1))
                                                _ = v90
                                            default:
                                                _ = len(v9)
                                            }
                                        case -571:
                                            switch fi26 {
                                            case -858:
                                                v2 = 408
                                            case 42:
                                                v9 = append(v9, v89)
                                            }
                                        default:
                                            if (1e10 > (1e10 * 0.1)) {
                                                v89 = fi61
                                                delete(v3, 855)
                                            } else {
                                                fmt.Println(fi61, fi61)
                                            }
                                        }
                                        var v91 uint64 = 2
                                        var v92 int = (fi61 & fi26)
                                        _ = v89
                                        _ = v91
                                        _ = v92
                                    case 32767:
                                        if false {
                                            var v93 int64 = -1
                                            var v94 uint8 = (v10 | (v8 & v10))
                                            _ = v93
                                            _ = v94
                                        }
                                    case 4294967295:
                                        fmt.Println(v10, v8)
                                    }
                                } else {
                                    switch ("\n\t" + (("foo bar" + "a") + ("" + "a"))) {
                                    case "hello":
                                        switch fi26 {
                                        case 72:
                                            v95 := GenS9{F0: (fi61 ^ fi26), F1: (v8 - v10), F2: 42}
                                            _ = v95
                                        case 657:
                                            v3[4294967295] = (^(fi61))
                                        default:
                                            v4[1] = fi61
                                        }
                                    }
                                }
                            case -707:
                                fmt.Println(v8)
                            default:
                                var v96 uint = 42
                                _ = v96
                            }
                            v1 = (^(v1))
                            _ = v64
                        }
                        v8 -= (v10 * v8)
                        var v97 int16 = 42
                        _ = fi61
                        _ = v97
                    }
                }
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                if v7 {
                    var v98 int32 = int32(fi26)
                    switch fi26 {
                    case -9223372036854775808:
                        v98 -= v98
                    case 287:
                        fmt.Println(v98)
                    }
                    if (((-(fi26)) >= (fi26 | fi26)) && (2 < 0)) {
                        if (!((^(v1)) != 449)) {
                            v7 = v7
                            for fi99 := 0; fi99 < 7; fi99++ {
                                if false {
                                    v8 -= (v8 % (((v8 / ((v8) | 1))) | 1))
                                    _ = len(v9)
                                    v8 = (^(v10))
                                    switch fi26 {
                                    case 825:
                                        var v100 int32 = v2
                                        _ = v100
                                    case 9223372036854775807:
                                        if v7 {
                                            if v7 {
                                                var v101 uint8 = uint8(v1)
                                                v102 := GenS4{F0: (v1 - v1), F1: 42}
                                                _ = v101
                                                _ = v102
                                            }
                                            var v103 int64 = 9223372036854775807
                                            var v104 string = ("hello" + (("a" + "") + (("a" + "hello") + "hello")))
                                            _ = v103
                                            _ = v104
                                        } else {
                                            if (-77 == 0) {
                                                v105 := GenS10{F0: (fi99 ^ int(v10))}
                                                _ = v105
                                            }
                                            var v106 uint8 = v8
                                            _ = v106
                                        }
                                    case -166:
                                        if v7 {
                                            if v7 {
                                                var v107 int16 = -827
                                                var v108 string = "\n\t"
                                                if false {
                                                    var v109 uint = 18446744073709551615
                                                    _ = v109
                                                } else {
                                                    v107 = (v107 ^ (v107 ^ v107))
                                                    v110 := GenS2{F0: 65535}
                                                    _ = v110
                                                }
                                                _ = v107
                                                _ = v108
                                            } else {
                                                var v111 float32 = 0.1
                                                if v7 {
                                                    if len(v9) >= 2 {
                                                        v9 = v9[1:len(v9)]
                                                    }
                                                    v1 -= 0
                                                } else {
                                                    var v112 float32 = float32((((-2.25 + (0.0 - -2.25)) / 1.5) - 1e10))
                                                    v113 := GenS10{F0: -939}
                                                    v114 := GenS3{F0: 4294967295}
                                                    _ = v112
                                                    _ = v113
                                                    _ = v114
                                                }
                                                fmt.Println(fi26)
                                                _ = v111
                                            }
                                            var v115 int = fi26
                                            switch "foo bar" {
                                            case "\n\t":
                                                v116 := GenS4{F0: v1, F1: 393}
                                                _ = v116
                                            case "hello":
                                                var v117 uint32 = 4294967295
                                                _ = v117
                                            default:
                                                v118 := GenS5{F0: 18446744073709551615, F1: -49}
                                                _ = v118
                                            }
                                            _ = v115
                                        }
                                    default:
                                        var v119 float64 = (1e10 - 3.14159)
                                        _ = v119
                                    }
                                } else {
                                    if false {
                                        fmt.Println(v2)
                                    }
                                    switch "foo bar" {
                                    case "hello":
                                        v10 += v8
                                        if (!true) {
                                            if len(v9) >= 2 {
                                                v9 = v9[1:len(v9)]
                                            }
                                            var v120 uint16 = v1
                                            _ = v120
                                        } else {
                                            var v121 uint8 = 42
                                            if false {
                                                fmt.Println(fi26)
                                            } else {
                                                var v122 float32 = ((float32(-2.25) / 1.5) * 0.1)
                                                v2 -= v2
                                                v123 := GenS8{F0: 1, F1: -849}
                                                fmt.Println(v1)
                                                _ = v122
                                                _ = v123
                                            }
                                            v8 -= (^(v8))
                                            _ = v121
                                        }
                                    case "":
                                        v1 = uint16(v10)
                                    case "\n\t":
                                        var v124 int = fi99
                                        _ = v124
                                    }
                                    v125 := GenS11{F0: "\n\t"}
                                    _ = v125
                                }
                                _ = fi99
                            }
                            var v126 uint = 2
                            _ = v126
                        }
                        v2 = (-(v98))
                    } else {
                        v127 := GenS4{F0: v1, F1: 70}
                        if true {
                            var v128 uint16 = 177
                            _ = v128
                        } else {
                            fmt.Println(fi26)
                            var v129 uint8 = 202
                            _ = v129
                        }
                        _ = v127
                    }
                    _ = v98
                }
                _ = fi26
            }
            if ((v2 >> ((18446744073709551615) & 63)) >= 0) {
                v8 = (v8 / ((v10) | 1))
                for fi130 := 0; fi130 < 16; fi130++ {
                    v131 := GenS12{F0: "\n\t", F1: 42, F2: uint(v10)}
                    delete(v3, 42)
                    _ = fi130
                    _ = v131
                }
                var v132 string = ("\n\t" + "\n\t")
                _ = v132
            }
            switch ("" + "foo bar") {
            case "a":
                for fi133 := 0; fi133 < 20; fi133++ {
                    var v134 uint = 1
                    _ = fi133
                    _ = v134
                }
            }
        }
    case -140:
        var v135 int16 = 42
        _ = v135
    default:
        v1 = 225
    }
    if ((32767 <= -200) || v7) {
        v136 := v3[363]
        _ = v136
    } else {
        if v7 {
            for fi137 := 0; fi137 < 12; fi137++ {
                switch fi137 {
                case -553:
                    var v138 int16 = 903
                    _ = v138
                default:
                    v8 -= v8
                }
                if false {
                    fmt.Println(fi137, v7)
                }
                _ = fi137
            }
            if v7 {
                for fi139 := 0; fi139 < 19; fi139++ {
                    if (!false) {
                        var v140 float32 = 1.5
                        _ = v140
                    } else {
                        v141 := GenS13{F0: -1, F1: v7}
                        _ = v141
                    }
                    _ = len(v3)
                    _ = fi139
                }
                v142 := GenS14{F0: 591, F1: ("foo bar" + (("hello" + "\n\t") + "foo bar"))}
                _ = v142
            } else {
                _ = len(v3)
                switch 2147483647 {
                case 42:
                    switch "\n\t" {
                    case "":
                        var v143 float64 = -2.25
                        _ = v143
                    }
                }
                if v7 {
                    fmt.Println(v2)
                    fmt.Println(v8, v8)
                    if ((((-2.25 + (3.14159 * 0.1)) * 1e10) / -2.25) != (-2.25 + ((3.14159 - 0.1) * 1e10))) {
                        for fi144 := 0; fi144 < 19; fi144++ {
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                            for fi145 := 0; fi145 < 10; fi145++ {
                                if v7 {
                                    switch fi145 {
                                    case 0:
                                        fmt.Println(v8)
                                    case -9223372036854775808:
                                        v1 = (v1 + (v1 | (v1 + (^(v1)))))
                                    case 920:
                                        v146 := GenS15{F0: 1, F1: 2, F2: 2147483647}
                                        _ = v146
                                    }
                                    if true {
                                        if (v8 != 255) {
                                            var v147 bool = v7
                                            _ = v147
                                        } else {
                                            v148 := GenS16{F0: v7}
                                            fmt.Println(v1)
                                            fmt.Println(fi145, v1)
                                            _ = v148
                                        }
                                        var v149 uint16 = 65535
                                        _ = v149
                                    }
                                } else {
                                    switch fi145 {
                                    case -80:
                                        fmt.Println(v10, v8)
                                        if true {
                                            var v150 uint8 = v8
                                            _ = v150
                                        } else {
                                            if (!v7) {
                                                switch "a" {
                                                case "hello":
                                                    if (1000 < 1) {
                                                        switch "hello" {
                                                        case "hello":
                                                            if false {
                                                                if v7 {
                                                                    var v151 int16 = 1
                                                                    v9 = append(v9, fi145)
                                                                    switch ("" + "hello") {
                                                                    case "foo bar":
                                                                        switch 0 {
                                                                        case 4294967295:
                                                                            v9 = append(v9, (fi144 % ((fi144) | 1)))
                                                                        case 494:
                                                                            v152 := GenS7{F0: ((fi144 / ((2147483647) | 1)) > fi145)}
                                                                            var v153 string = "foo bar"
                                                                            v154 := GenS17{F0: v2, F1: v151, F2: (v2 % (((-(v2))) | 1))}
                                                                            _ = v152
                                                                            _ = v153
                                                                            _ = v154
                                                                        default:
                                                                            v4[65535] = (fi145 ^ int(v151))
                                                                        }
                                                                    default:
                                                                        v9 = append(v9, (fi145 % (((^(fi145))) | 1)))
                                                                    }
                                                                    _ = v151
                                                                }
                                                            }
                                                        case "foo bar":
                                                            v155 := GenS3{F0: 255}
                                                            _ = v155
                                                        }
                                                        var v156 int16 = 0
                                                        _ = v156
                                                    } else {
                                                        v10 += (v10 << ((164) & 63))
                                                        fmt.Println(fi145, v10)
                                                        var v157 float32 = (1.5 - ((float32(3.14159) / 3.14159) + ((1.5 - (float32(1e10) + ((-2.25 + 1.5) * float32(0.0)))) - (0.1 / 1e10))))
                                                        _ = v157
                                                    }
                                                case "":
                                                    var v158 bool = ((float32(1e10) == (((((((((0.0 - (3.14159 * 0.0)) / 1.5) * 0.0) / -2.25) / (-2.25 + 0.1)) * 0.1) - 0.1) - (1.5 / float32(float64((3.14159 * float32((0.1 + 0.1))))))) * float32(0.0))) || (!v7))
                                                    _ = v158
                                                }
                                                switch 1 {
                                                case 247:
                                                    var v159 bool = (255 >= v1)
                                                    if false {
                                                        var v160 int8 = 100
                                                        var v161 string = ""
                                                        _ = v160
                                                        _ = v161
                                                    } else {
                                                        if false {
                                                            if (!v7) {
                                                                v159 = true
                                                                var v162 int32 = (v2 * 0)
                                                                v2 = 2147483647
                                                                _ = v162
                                                            } else {
                                                                v163 := GenS16{F0: v159}
                                                                fmt.Println(v159, v7, v7)
                                                                var v164 uint8 = 0
                                                                _ = v163
                                                                _ = v164
                                                            }
                                                        }
                                                        var v165 int32 = 148
                                                        v166 := GenS6{F0: 11, F1: v8}
                                                        _ = v165
                                                        _ = v166
                                                    }
                                                    fmt.Println(fi144)
                                                    fmt.Println(fi144)
                                                    var v167 int64 = -369
                                                    _ = v159
                                                    _ = v167
                                                case -160:
                                                    if (v7 && (true || (!(!(!("a" == "\n\t")))))) {
                                                        if true {
                                                            fmt.Println(fi144, fi145)
                                                            v8 = (v8 | v8)
                                                            switch 32768 {
                                                            case 42:
                                                                delete(v4, 4294967295)
                                                            case 411:
                                                                if v7 {
                                                                    var v168 uint = 402
                                                                    var v169 int8 = 46
                                                                    _ = v168
                                                                    _ = v169
                                                                } else {
                                                                    v170 := GenS18{F0: false, F1: 854}
                                                                    _ = v170
                                                                }
                                                            case -9223372036854775808:
                                                                v2 -= 2147483647
                                                            }
                                                        } else {
                                                            v2 = -1
                                                        }
                                                        v7 = false
                                                    } else {
                                                        var v171 int = fi144
                                                        _ = v171
                                                    }
                                                case 1:
                                                    v172 := GenS19{F0: 42}
                                                    _ = v172
                                                }
                                                switch "foo bar" {
                                                case "hello":
                                                    fmt.Println(v2, v2)
                                                case "":
                                                    v10 -= (v10 % (((^(v8))) | 1))
                                                    _ = len(v3)
                                                    if len(v9) > 0 {
                                                        v9[0] = 4294967295
                                                    }
                                                    if false {
                                                        var v173 bool = (v2 <= (v2 % (((v2 % ((v2) | 1))) | 1)))
                                                        if v7 {
                                                            var v174 uint8 = v10
                                                            var v175 string = "hello"
                                                            _ = v174
                                                            _ = v175
                                                        } else {
                                                            if v173 {
                                                                switch 2147483647 {
                                                                case -2147483648:
                                                                    if len(v9) > 0 {
                                                                        v9[0] = (fi144 * fi144)
                                                                    }
                                                                case -284:
                                                                    switch (fi144 % (((fi144 % (((fi144 % ((fi145) | 1))) | 1))) | 1)) {
                                                                    case 288:
                                                                        if len(v9) >= 2 {
                                                                            v9 = v9[1:len(v9)]
                                                                        }
                                                                    case 127:
                                                                        if true {
                                                                            v176 := GenS7{F0: false}
                                                                            switch "foo bar" {
                                                                            case "":
                                                                                v177 := GenS20{F0: 4294967295}
                                                                                _ = v177
                                                                            case "a":
                                                                                v178 := GenS12{F0: (((("\n\t" + "foo bar") + "") + "hello") + ""), F1: -9223372036854775808, F2: 255}
                                                                                _ = v178
                                                                            default:
                                                                                switch ("foo bar" + "foo bar") {
                                                                                case "a":
                                                                                    var v179 float64 = 0.0
                                                                                    _ = v179
                                                                                case "\n\t":
                                                                                    switch ("foo bar" + "foo bar") {
                                                                                    case "hello":
                                                                                        var v180 float32 = (float32(1e10) + float32(1.5))
                                                                                        _ = v180
                                                                                    case "\n\t":
                                                                                        v181 := GenS21{F0: 42, F1: "a"}
                                                                                        _ = v181
                                                                                    default:
                                                                                        fmt.Println(v8, v10)
                                                                                    }
                                                                                default:
                                                                                    var v182 int = 32768
                                                                                    _ = v182
                                                                                }
                                                                            }
                                                                            switch (-(fi144)) {
                                                                            case 706:
                                                                                var v183 uint32 = 1
                                                                                _ = v183
                                                                            default:
                                                                                var v184 int32 = 2147483647
                                                                                _ = v184
                                                                            }
                                                                            _ = v176
                                                                        }
                                                                    case 4294967295:
                                                                        v185 := GenS22{F0: (^(v1)), F1: 551, F2: 42}
                                                                        _ = v185
                                                                    default:
                                                                        switch ("" + (("\n\t" + "\n\t") + "\n\t")) {
                                                                        case "a":
                                                                            v8 -= (v10 - v10)
                                                                        }
                                                                        var v186 uint32 = 65535
                                                                        v2 = (-(v2))
                                                                        _ = v186
                                                                    }
                                                                default:
                                                                    if v173 {
                                                                        fmt.Println(v7, v1, v173)
                                                                        fmt.Println(v1, v173, fi144)
                                                                        var v187 uint16 = (v1 & (v1 << ((18446744073709551615) & 63)))
                                                                        _ = v187
                                                                    }
                                                                }
                                                            }
                                                            var v188 float32 = 0.0
                                                            _ = v188
                                                        }
                                                        _ = v173
                                                    } else {
                                                        v189 := GenS7{F0: ((837 < 1) && v7)}
                                                        _ = v189
                                                    }
                                                    v8 = 128
                                                default:
                                                    if true {
                                                        v7 = (v7 && (v7 || false))
                                                    }
                                                }
                                            } else {
                                                if len(v9) > 0 {
                                                    v9[0] = fi145
                                                }
                                            }
                                            fmt.Println(v7, v10)
                                            var v190 int = -1
                                            v191 := GenS9{F0: fi145, F1: 1, F2: 1}
                                            switch "" {
                                            case "foo bar":
                                                fmt.Println(v10)
                                            case "a":
                                                _ = len(v3)
                                            default:
                                                v9 = append(v9, (v190 & (-(fi145))))
                                            }
                                            _ = v190
                                            _ = v191
                                        }
                                    case 4294967295:
                                        v2 = -2147483648
                                    default:
                                        var v192 int16 = -59
                                        _ = v192
                                    }
                                    v2 += v2
                                }
                                var v193 uint16 = v1
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                                _ = fi145
                                _ = v193
                            }
                            _ = fi144
                        }
                        if v7 {
                            var v194 int8 = 0
                            _ = v194
                        }
                    } else {
                        for fi195 := 0; fi195 < 8; fi195++ {
                            v10 -= v8
                            switch (fi195 ^ fi195) {
                            case 32767:
                                var v196 float32 = 1e10
                                _ = v196
                            }
                            _ = fi195
                        }
                    }
                }
            }
        } else {
            if (-113 > 922) {
                v2 += 2147483647
            } else {
                if true {
                    if (v7 || v7) {
                        fmt.Println(v1, v8)
                        fmt.Println(v10, v1)
                        if ((v8 / ((v10) | 1)) < (v10 << ((361) & 63))) {
                            fmt.Println(v7)
                            var v197 int32 = -2147483648
                            var v198 float64 = float64(3.14159)
                            _ = v197
                            _ = v198
                        }
                    }
                    var v199 float32 = 1e10
                    _ = v199
                }
                v8 = v8
            }
            var v200 uint64 = 18446744073709551615
            _ = v200
        }
        fmt.Println(v8, v7, v7)
    }
    if ((-9223372036854775808 >= 0) && false) {
        v201 := GenS14{F0: 199, F1: "foo bar"}
        if (v7 || (v1 <= v1)) {
            v202 := GenS23{F0: v10, F1: 42, F2: int32(v8)}
            var v203 int16 = 42
            _ = v202
            _ = v203
        } else {
            switch 128 {
            case 32768:
                if (((3.14159 * 0.0) + -2.25) != 0.0) {
                    v204 := GenS6{F0: (v8 + 1), F1: (^(v8))}
                    v8 += (v8 - (v8 << ((4294967295) & 63)))
                    _ = v204
                } else {
                    v8 += v10
                }
            default:
                v10 = v10
            }
            var v205 uint64 = 0
            _ = v205
        }
        _ = v201
    }
    switch "\n\t" {
    case "":
        _ = len(v3)
    default:
        var v206 float32 = 1.5
        _ = v206
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
