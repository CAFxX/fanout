package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 int16; F2 int16 }
type GenS3 struct { F0 uint32; F1 int16; F2 int32 }
type GenS4 struct { F0 uint; F1 int }
type GenS5 struct { F0 bool; F1 int8; F2 uint16 }
type GenS6 struct { F0 uint16; F1 uint32 }
type GenS7 struct { F0 int32; F1 uint8 }
type GenS8 struct { F0 uint; F1 int64 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint16; F1 int64; F2 int64 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint32 }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []int{34, 32768, 128}
    var v3 int8 = -128
    v4 := []uint32{2147483648, 2147483648}
    for fi5 := 0; fi5 < 8; fi5++ {
        var v6 int8 = v3
        var v7 int8 = (v3 - (v3 | (v3 ^ (v6 - -97))))
        var v8 uint64 = v1
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    if (false || (3.14159 >= 3.14159)) {
        switch 32767 {
        case -9223372036854775808:
            v3 = (-(v3))
        default:
            if true {
                if ((-(v3)) == v3) {
                    for fi9 := 0; fi9 < 17; fi9++ {
                        var v10 float32 = (1e10 * (((-2.25 - 3.14159) - 0.1) * 1e10))
                        if false {
                            v1 += (v1 - v1)
                        } else {
                            v4 = append(v4, 2147483648)
                        }
                        var v11 int64 = 1
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(v3, v1)
                    case "\n\t":
                        if false {
                            if true {
                                for fi12 := 0; fi12 < 15; fi12++ {
                                    var v13 int16 = 42
                                    v3 = v3
                                    v13 = v13
                                    _ = fi12
                                    _ = v13
                                }
                                v3 -= v3
                                if (42 == 1) {
                                    v4 = append(v4, 42)
                                }
                            } else {
                                var v14 uint8 = 54
                                switch -2147483648 {
                                case 0:
                                    var v15 int8 = 0
                                    _ = v15
                                case 367:
                                    switch "\n\t" {
                                    case "":
                                        fmt.Println(v1, v3, v14)
                                    }
                                case -2147483648:
                                    fmt.Println(v1, v1, v14)
                                }
                                _ = v14
                            }
                        }
                    default:
                        switch -1 {
                        case 32768:
                            switch 32768 {
                            case 9223372036854775807:
                                _ = len(v2)
                            case -2147483648:
                                switch 42 {
                                case 32768:
                                    for fi16 := 0; fi16 < 12; fi16++ {
                                        v3 = v3
                                        _ = len(v4)
                                        var v17 float64 = (0.0 - 1e10)
                                        _ = fi16
                                        _ = v17
                                    }
                                case -1:
                                    v1 += v1
                                case 42:
                                    var v18 uint32 = 42
                                    _ = v18
                                default:
                                    var v19 int = 9223372036854775807
                                    _ = v19
                                }
                            default:
                                var v20 uint64 = v1
                                _ = v20
                            }
                        default:
                            var v21 uint = 255
                            _ = v21
                        }
                    }
                } else {
                    var v22 int = 9223372036854775807
                    _ = v22
                }
            }
        }
        fmt.Println(v3, v3)
    } else {
        v23 := GenS1{F0: 65535}
        if (!false) {
            var v24 uint8 = 255
            _ = v24
        } else {
            v25 := GenS1{F0: 2147483647}
            _ = v25
        }
        _ = v23
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        switch "" {
        case "foo bar":
            _ = len(v2)
        case "\n\t":
            var v27 uint = 65535
            _ = v27
        case "":
            switch ("hello" + "") {
            case "foo bar":
                if true {
                    v3 = 0
                    for fi28 := 0; fi28 < 11; fi28++ {
                        var v29 float64 = 3.14159
                        _ = fi28
                        _ = v29
                    }
                    v30 := GenS1{F0: 1}
                    _ = v30
                } else {
                    var v31 int = fi26
                    switch "foo bar" {
                    case "foo bar":
                        if true {
                            v32 := GenS2{F0: (2 <= (^(v1))), F1: 1, F2: -1}
                            for fi33 := 0; fi33 < 13; fi33++ {
                                var v34 int32 = -2147483648
                                fmt.Println(fi26)
                                switch (fi26 - v31) {
                                case 1:
                                    var v35 float32 = 0.1
                                    _ = v35
                                case -9223372036854775808:
                                    switch "a" {
                                    case "a":
                                        if true {
                                            var v36 uint64 = v1
                                            if len(v4) > 0 {
                                                v4[0] = 4294967295
                                            }
                                            if len(v4) > 0 {
                                                v4[0] = 694
                                            }
                                            _ = v36
                                        } else {
                                            v37 := GenS3{F0: 136, F1: 354, F2: -1}
                                            var v38 uint32 = 2147483648
                                            _ = v37
                                            _ = v38
                                        }
                                    case "foo bar":
                                        v34 = 1
                                        v1 = v1
                                    case "":
                                        v1 -= v1
                                    default:
                                        switch "a" {
                                        case "\n\t":
                                            v3 += (v3 ^ -55)
                                        case "a":
                                            fmt.Println(fi33, v31)
                                        case "foo bar":
                                            if len(v2) > 0 {
                                                v2[0] = (fi33 * v31)
                                            }
                                        }
                                    }
                                case 4294967295:
                                    if false {
                                        v3 += v3
                                    } else {
                                        var v39 int64 = -978
                                        var v40 uint16 = 1
                                        _ = v39
                                        _ = v40
                                    }
                                default:
                                    switch "\n\t" {
                                    case "a":
                                        if false {
                                            v34 = 558
                                            var v41 int64 = 787
                                            v31 -= fi33
                                            _ = v41
                                        } else {
                                            switch v31 {
                                            case 127:
                                                if true {
                                                    _ = len(v2)
                                                    v1 += (v1 | (v1 * 18446744073709551615))
                                                    if true {
                                                        var v42 int32 = 42
                                                        switch "foo bar" {
                                                        case "hello":
                                                            if len(v2) > 0 {
                                                                v2[0] = v31
                                                            }
                                                        case "\n\t":
                                                            var v43 uint32 = 2147483648
                                                            _ = v43
                                                        case "":
                                                            if (((3.14159 - 3.14159) / 0.1) != 0.1) {
                                                                if (1 > -1) {
                                                                    var v44 uint = 775
                                                                    var v45 uint32 = 2
                                                                    _ = v44
                                                                    _ = v45
                                                                }
                                                                if true {
                                                                    if false {
                                                                        if (42 <= 0) {
                                                                            var v46 int8 = v3
                                                                            _ = v46
                                                                        }
                                                                        var v47 float64 = (0.1 * 0.0)
                                                                        switch ("\n\t" + ((("hello" + "") + "") + "")) {
                                                                        case "":
                                                                            v48 := GenS4{F0: 0, F1: fi26}
                                                                            _ = v48
                                                                        }
                                                                        _ = v47
                                                                    } else {
                                                                        fmt.Println(v1, v42, v34)
                                                                        v49 := GenS5{F0: false, F1: (v3 - (v3 + -88)), F2: 220}
                                                                        var v50 int16 = -529
                                                                        _ = v49
                                                                        _ = v50
                                                                    }
                                                                }
                                                                var v51 float64 = -2.25
                                                                _ = v51
                                                            }
                                                        }
                                                        switch "hello" {
                                                        case "foo bar":
                                                            v52 := GenS6{F0: 2, F1: 1}
                                                            v53 := GenS7{F0: (v42 & -1), F1: 227}
                                                            if ((v3 << ((0) & 63)) != (v3 % (((v3 ^ v3)) | 1))) {
                                                                fmt.Println(v42)
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                                fmt.Println(v34, v3)
                                                            } else {
                                                                switch "\n\t" {
                                                                case "a":
                                                                    switch ("" + (("hello" + "") + (("\n\t" + "foo bar") + "hello"))) {
                                                                    case "a":
                                                                        v54 := GenS5{F0: false, F1: v3, F2: 263}
                                                                        _ = v54
                                                                    default:
                                                                        v1 = 476
                                                                    }
                                                                case "\n\t":
                                                                    var v55 int = (fi33 >> ((65535) & 63))
                                                                    _ = v55
                                                                }
                                                                _ = len(v4)
                                                            }
                                                            _ = v52
                                                            _ = v53
                                                        }
                                                        _ = v42
                                                    } else {
                                                        v56 := GenS6{F0: 379, F1: 287}
                                                        if false {
                                                            v57 := GenS8{F0: 2, F1: -1}
                                                            switch "hello" {
                                                            case "foo bar":
                                                                var v58 int8 = v3
                                                                _ = v58
                                                            default:
                                                                if false {
                                                                    v59 := GenS9{F0: 32767}
                                                                    fmt.Println(v34, v3, v1)
                                                                    var v60 bool = true
                                                                    fmt.Println(v34)
                                                                    _ = v59
                                                                    _ = v60
                                                                } else {
                                                                    if (false && true) {
                                                                        switch "" {
                                                                        case "":
                                                                            v34 = (v34 - 2147483647)
                                                                        case "\n\t":
                                                                            var v61 uint8 = 96
                                                                            _ = v61
                                                                        }
                                                                        switch (fi33 * (^(fi26))) {
                                                                        case 32767:
                                                                            var v62 int8 = v3
                                                                            _ = v62
                                                                        case 42:
                                                                            var v63 uint64 = v1
                                                                            _ = v63
                                                                        default:
                                                                            v64 := GenS4{F0: 98, F1: 128}
                                                                            _ = v64
                                                                        }
                                                                    }
                                                                    var v65 int64 = 9223372036854775807
                                                                    switch "hello" {
                                                                    case "a":
                                                                        v3 = (-(v3))
                                                                    case "hello":
                                                                        v2 = append(v2, fi33)
                                                                    default:
                                                                        var v66 uint64 = v1
                                                                        _ = v66
                                                                    }
                                                                    _ = v65
                                                                }
                                                            }
                                                            switch fi26 {
                                                            case 4294967295:
                                                                if (!false) {
                                                                    v1 += v1
                                                                } else {
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                }
                                                            case 357:
                                                                v67 := GenS3{F0: 65535, F1: 784, F2: v34}
                                                                _ = v67
                                                            case 128:
                                                                if false {
                                                                    v31 -= fi33
                                                                    if (134 > 469) {
                                                                        if false {
                                                                            v68 := GenS1{F0: 2147483647}
                                                                            fmt.Println(v1)
                                                                            _ = v68
                                                                        } else {
                                                                            if false {
                                                                                switch "" {
                                                                                case "":
                                                                                    v34 -= 1
                                                                                case "a":
                                                                                    if true {
                                                                                        v1 -= 2
                                                                                        v31 = (^(v31))
                                                                                        v34 += (v34 | (v34 * 0))
                                                                                    } else {
                                                                                        fmt.Println(v31, v1, v34)
                                                                                        _ = len(v4)
                                                                                        if true {
                                                                                            fmt.Println(v1, v31)
                                                                                            v31 += 32768
                                                                                            v34 += 0
                                                                                            switch "hello" {
                                                                                            case "foo bar":
                                                                                                v3 = (^(v3))
                                                                                                if (v1 == 0) {
                                                                                                    var v69 bool = true
                                                                                                    v31 += -2147483648
                                                                                                    v69 = ("foo bar" != "\n\t")
                                                                                                    _ = v69
                                                                                                } else {
                                                                                                    v34 = 576
                                                                                                }
                                                                                                switch ("a" + "foo bar") {
                                                                                                case "\n\t":
                                                                                                    fmt.Println(v31, fi26, fi33)
                                                                                                    var v70 int64 = -1
                                                                                                    v1 = v1
                                                                                                    _ = v70
                                                                                                }
                                                                                            case "a":
                                                                                                if (false || (!true)) {
                                                                                                    var v71 int64 = -9223372036854775808
                                                                                                    _ = v71
                                                                                                }
                                                                                            case "\n\t":
                                                                                                if (true && true) {
                                                                                                    var v72 int64 = -342
                                                                                                    if len(v2) > 0 {
                                                                                                        v2[0] = (v31 / ((-456) | 1))
                                                                                                    }
                                                                                                    _ = v72
                                                                                                }
                                                                                            }
                                                                                            if false {
                                                                                                switch "hello" {
                                                                                                case "hello":
                                                                                                    if (!true) {
                                                                                                        var v73 int = 9223372036854775807
                                                                                                        if len(v4) > 0 {
                                                                                                            v4[0] = 258
                                                                                                        }
                                                                                                        _ = v73
                                                                                                    }
                                                                                                case "foo bar":
                                                                                                    switch (fi33 << ((uint(v1)) & 63)) {
                                                                                                    case 1:
                                                                                                        v3 = (v3 % ((v3) | 1))
                                                                                                    case 32768:
                                                                                                        if ("" != ("" + ("" + "\n\t"))) {
                                                                                                            var v74 bool = false
                                                                                                            v2 = append(v2, (fi26 & -2147483648))
                                                                                                            _ = v74
                                                                                                        }
                                                                                                    default:
                                                                                                        v3 = (v3 * (v3 * 127))
                                                                                                        var v75 bool = false
                                                                                                        var v76 int = v31
                                                                                                        _ = v75
                                                                                                        _ = v76
                                                                                                    }
                                                                                                default:
                                                                                                    _ = len(v2)
                                                                                                    v31 = fi33
                                                                                                    v77 := GenS4{F0: 18446744073709551615, F1: fi26}
                                                                                                    _ = v77
                                                                                                }
                                                                                                fmt.Println(fi26)
                                                                                            }
                                                                                        } else {
                                                                                            v3 -= v3
                                                                                            _ = len(v4)
                                                                                            var v78 int16 = 468
                                                                                            _ = v78
                                                                                        }
                                                                                    }
                                                                                case "foo bar":
                                                                                    switch fi33 {
                                                                                    case 4294967295:
                                                                                        fmt.Println(fi33, v34)
                                                                                    case -1:
                                                                                        v1 -= 18446744073709551615
                                                                                    }
                                                                                }
                                                                                fmt.Println(fi26, v3, v3)
                                                                                switch "foo bar" {
                                                                                case "hello":
                                                                                    fmt.Println(v3, fi26, v1)
                                                                                case "foo bar":
                                                                                    v3 += v3
                                                                                case "":
                                                                                    if false {
                                                                                        v3 += 57
                                                                                        v79 := GenS8{F0: 2, F1: -276}
                                                                                        _ = v79
                                                                                    }
                                                                                default:
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = 2
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v31, fi33, v31)
                                                                            }
                                                                            if len(v2) > 0 {
                                                                                v2[0] = fi33
                                                                            }
                                                                            if (v3 >= (v3 / (((v3 << ((0) & 63))) | 1))) {
                                                                                var v80 uint = 4294967295
                                                                                v81 := GenS9{F0: (fi26 / ((1) | 1))}
                                                                                _ = v80
                                                                                _ = v81
                                                                            }
                                                                            v31 -= (-(fi26))
                                                                        }
                                                                        switch v31 {
                                                                        case -183:
                                                                            if (1 < 1) {
                                                                                v3 = 111
                                                                                v1 = v1
                                                                                var v82 uint16 = 1
                                                                                _ = v82
                                                                            } else {
                                                                                switch v31 {
                                                                                case -9223372036854775808:
                                                                                    var v83 int32 = (v34 % ((v34) | 1))
                                                                                    _ = v83
                                                                                case -420:
                                                                                    v1 -= v1
                                                                                case 189:
                                                                                    switch ("a" + "\n\t") {
                                                                                    case "\n\t":
                                                                                        if false {
                                                                                            if len(v4) > 0 {
                                                                                                v4[0] = 2147483647
                                                                                            }
                                                                                            var v84 int32 = (v34 * v34)
                                                                                            v85 := GenS9{F0: v31}
                                                                                            _ = v84
                                                                                            _ = v85
                                                                                        }
                                                                                    case "a":
                                                                                        v4 = append(v4, 0)
                                                                                    }
                                                                                }
                                                                                switch "foo bar" {
                                                                                case "foo bar":
                                                                                    var v86 uint8 = 255
                                                                                    _ = v86
                                                                                }
                                                                            }
                                                                        }
                                                                    } else {
                                                                        v1 = (v1 / ((603) | 1))
                                                                        v87 := GenS10{F0: 0, F1: 1, F2: 42}
                                                                        if ((v3 % ((v3) | 1)) >= v3) {
                                                                            var v88 float64 = ((0.0 - -2.25) * -2.25)
                                                                            var v89 uint16 = 42
                                                                            _ = v88
                                                                            _ = v89
                                                                        }
                                                                        _ = len(v2)
                                                                        if ((v3 * (v3 % ((7) | 1))) >= v3) {
                                                                            v4 = append(v4, 2147483648)
                                                                        } else {
                                                                            _ = len(v4)
                                                                        }
                                                                        _ = v87
                                                                    }
                                                                    var v90 uint16 = 0
                                                                    _ = v90
                                                                } else {
                                                                    v31 = fi33
                                                                }
                                                            default:
                                                                fmt.Println(fi33)
                                                            }
                                                            _ = v57
                                                        }
                                                        v1 = (v1 / (((v1 >> ((188) & 63))) | 1))
                                                        _ = v56
                                                    }
                                                } else {
                                                    if (true || (!false)) {
                                                        v4 = append(v4, 42)
                                                        v31 = (fi33 + (fi33 - 128))
                                                    } else {
                                                        var v91 int8 = (v3 + (v3 & -128))
                                                        _ = v91
                                                    }
                                                    var v92 string = ""
                                                    _ = v92
                                                }
                                            }
                                            v34 += (v34 | -2147483648)
                                            var v93 float64 = ((3.14159 / 1.5) - (0.0 / 1.5))
                                            _ = v93
                                        }
                                    case "\n\t":
                                        fmt.Println(fi33, fi26, v3)
                                    }
                                }
                                _ = fi33
                                _ = v34
                            }
                            _ = v32
                        }
                    case "":
                        var v94 bool = ((-(v31)) < -859)
                        _ = v94
                    default:
                        var v95 int32 = -110
                        _ = v95
                    }
                    for fi96 := 0; fi96 < 6; fi96++ {
                        v97 := GenS11{F0: 42}
                        var v98 string = "foo bar"
                        _ = fi96
                        _ = v97
                        _ = v98
                    }
                    _ = v31
                }
            case "\n\t":
                fmt.Println(v3, v1, fi26)
            case "a":
                v99 := GenS12{F0: 4294967295}
                _ = v99
            default:
                for fi100 := 0; fi100 < 12; fi100++ {
                    if (1.5 > 1.5) {
                        var v101 float32 = -2.25
                        v3 = v3
                        v1 = 42
                        _ = v101
                    } else {
                        var v102 int16 = 42
                        _ = v102
                    }
                    v1 += 133
                    var v103 string = ((("\n\t" + "\n\t") + "") + ("" + "hello"))
                    _ = fi100
                    _ = v103
                }
            }
        }
        _ = fi26
    }
    v104 := GenS9{F0: 128}
    _ = len(v2)
    v1 += v1
    if (255 < 292) {
        if true {
            for fi105 := 0; fi105 < 2; fi105++ {
                switch "\n\t" {
                case "a":
                    v106 := GenS5{F0: (0 != -9223372036854775808), F1: v3, F2: 65535}
                    _ = v106
                default:
                    v1 += v1
                }
                _ = fi105
            }
            if (!false) {
                var v107 int64 = -201
                if false {
                    v108 := GenS8{F0: 1, F1: (v107 / ((v107) | 1))}
                    fmt.Println(v107, v3)
                    var v109 uint8 = 255
                    _ = v108
                    _ = v109
                }
                v3 = v3
                _ = v107
            }
        }
    }
    switch "a" {
    case "foo bar":
        v1 = v1
    }
    v1 -= 0
    v4 = append(v4, 65535)
    switch 42 {
    case 128:
        var v110 int64 = 0
        _ = v110
    case -22:
        if len(v2) > 0 {
            v2[0] = 32768
        }
    case -2147483648:
        var v111 uint32 = 954
        _ = v111
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v104
    _ = v2
    _ = v3
    _ = v4
}
