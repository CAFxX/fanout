package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 uint; F2 uint }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 int; F1 int16; F2 int32 }
type GenS6 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int16; F1 int }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS12 struct { F0 uint16; F1 uint64 }
type GenS13 struct { F0 uint32; F1 uint8; F2 int64 }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 int32; F1 int; F2 string }
func genFn0(p0 uint64, p1 int32, p2 int32) int {
    if (0 >= 1) {
        for fi7 := 0; fi7 < 11; fi7++ {
            for fi8 := 0; fi8 < 13; fi8++ {
                p2 = -1
                fmt.Println(p2, fi8)
                p0 = p0
                p1 -= p2
                _ = fi8
            }
            for fi9 := 0; fi9 < 9; fi9++ {
                switch "foo bar" {
                case "\n\t":
                    p2 -= (p1 * p2)
                case "foo bar":
                    v10 := GenS1{F0: 4294967295, F1: 1}
                    _ = v10
                }
                p0 = 1
                v11 := []int{1}
                fmt.Println(fi9, fi7)
                _ = fi9
                _ = v11
            }
            v12 := []int{(fi7 ^ fi7), -889, (fi7 + fi7), 32767}
            _ = fi7
            _ = v12
        }
        var v13 uint64 = p0
        for fi14 := 0; fi14 < 5; fi14++ {
            p1 = p1
            _ = fi14
        }
        _ = v13
    }
    var v15 int = 32767
    switch (^(v15)) {
    case 9223372036854775807:
        if true {
            var v16 float64 = (((1e10 + -2.25) * 1e10) + (3.14159 * 0.0))
            v17 := GenS2{F0: 128}
            v18 := GenS3{F0: 9223372036854775807, F1: uint(p0), F2: 846}
            _ = v16
            _ = v17
            _ = v18
        }
    case 32767:
        fmt.Println(p1, v15, v15)
    case 547:
        if (p0 != (p0 * (^(p0)))) {
            fmt.Println(p0, p2)
            switch "" {
            case "\n\t":
                var v19 uint8 = 79
                _ = v19
            default:
                var v20 float32 = 3.14159
                _ = v20
            }
        } else {
            for fi21 := 0; fi21 < 2; fi21++ {
                v22 := []int{}
                p0 = (p0 + (p0 | p0))
                switch "foo bar" {
                case "":
                    if len(v22) > 0 {
                        v22[0] = fi21
                    }
                case "a":
                    var v23 float32 = 0.0
                    _ = v23
                }
                _ = fi21
                _ = v22
            }
        }
    default:
        p1 = p2
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 0
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[int]int)
    var v3 int32 = 1
    var v4 float32 = ((((((1e10 + ((-2.25 - 0.0) + ((0.1 * ((0.0 * 3.14159) + 0.0)) * ((0.0 + (((3.14159 - 1.5) - float32(0.1)) + 1.5)) - float32(0.1))))) + (float32(0.0) + float32(3.14159))) * 1e10) * (0.0 / -2.25)) - 1e10) / 3.14159)
    var v5 int32 = v3
    v6 := make(map[string]bool)
    for fi24 := 0; fi24 < 18; fi24++ {
        var v25 uint64 = 2
        _ = fi24
        _ = v25
    }
    if false {
        v26 := GenS3{F0: 709, F1: 430, F2: 42}
        v3 = -2147483648
        _ = v26
    }
    var v27 uint64 = 622
    fmt.Println(v3, v4, v4)
    v2[127] = 42
    v28 := []int{2147483647, 832, -60, 9223372036854775807}
    v29 := GenS1{F0: v1, F1: 65535}
    if true {
        if false {
            fmt.Println(v5, v27)
            v4 += v4
        }
        var v30 float32 = (((float32((1e10 + 3.14159)) - ((float32((((0.1 + 1e10) * -2.25) - 3.14159)) + v4) + 0.1)) - v4) / v4)
        var v31 bool = false
        _ = v30
        _ = v31
    } else {
        switch 128 {
        case -9223372036854775808:
            v32 := GenS1{F0: v1, F1: 1}
            _ = v32
        }
        switch "hello" {
        case "foo bar":
            switch 127 {
            case 32768:
                _ = len(v28)
            }
        case "":
            v28 = append(v28, int(v27))
        case "\n\t":
            for fi33 := 0; fi33 < 14; fi33++ {
                v3 = v5
                var v34 uint8 = uint8(v1)
                _ = fi33
                _ = v34
            }
        default:
            v3 = (v3 + 2147483647)
        }
        if false {
            if len(v28) > 0 {
                v28[0] = genFn0((v27 & 0), v5, (^(v5)))
            }
            var v35 bool = true
            v28 = append(v28, int(v1))
            _ = v35
        } else {
            fmt.Println(v5)
        }
    }
    fmt.Println(v27, v27)
    for fi36 := 0; fi36 < 10; fi36++ {
        switch genFn0((^(v27)), 2147483647, v3) {
        case 344:
            switch fi36 {
            case 4294967295:
                for fi37 := 0; fi37 < 20; fi37++ {
                    v4 = float32(0.0)
                    v3 += (v3 + v3)
                    _ = fi37
                }
            case 9223372036854775807:
                switch "hello" {
                case "foo bar":
                    if false {
                        var v38 uint16 = 1
                        v39 := GenS1{F0: uint32(v38), F1: 532}
                        _ = v38
                        _ = v39
                    }
                }
            case -776:
                var v40 int64 = 42
                _ = v40
            default:
                var v41 int = genFn0((v27 + v27), v3, v3)
                _ = v41
            }
        }
        for fi42 := 0; fi42 < 13; fi42++ {
            if false {
                switch ((("" + "") + "\n\t") + "a") {
                case "foo bar":
                    v3 += (v5 | v3)
                case "hello":
                    fmt.Println(fi42, fi36, v3)
                case "\n\t":
                    var v43 int = (fi42 / ((2147483647) | 1))
                    _ = v43
                default:
                    switch ("foo bar" + "hello") {
                    case "":
                        v44 := GenS1{F0: (^(v1)), F1: 2}
                        _ = v44
                    case "hello":
                        fmt.Println(fi42)
                    }
                    v3 -= v3
                    v1 = v1
                }
                var v45 int = 2147483647
                var v46 string = ((("a" + ("" + "hello")) + "hello") + "foo bar")
                _ = v45
                _ = v46
            } else {
                if (!true) {
                    if ("\n\t" < "") {
                        v1 -= 2147483648
                        fmt.Println(v3, v1)
                        if false {
                            if true {
                                var v47 float32 = 1.5
                                var v48 int16 = -32768
                                _ = v47
                                _ = v48
                            } else {
                                switch (fi36 >> ((767) & 63)) {
                                case 1:
                                    if ((!true) && false) {
                                        v3 = 192
                                        v3 = -2147483648
                                    }
                                case 42:
                                    fmt.Println(v3)
                                case 4294967295:
                                    delete(v2, (fi36 + 128))
                                    v49 := GenS3{F0: -413, F1: 42, F2: 18446744073709551615}
                                    _ = v49
                                }
                                if true {
                                    var v50 float64 = 0.1
                                    if false {
                                        v51 := GenS2{F0: 1}
                                        var v52 int = 1
                                        _ = v51
                                        _ = v52
                                    }
                                    _ = v50
                                } else {
                                    v27 += v27
                                    v3 = v5
                                    v27 += v27
                                }
                            }
                            switch ("\n\t" + "\n\t") {
                            case "\n\t":
                                v2[fi36] = (fi42 - fi42)
                            }
                            var v53 int8 = 42
                            _ = v53
                        }
                        fmt.Println(v27, v3)
                        v1 -= v1
                    } else {
                        var v54 bool = false
                        _ = v54
                    }
                    var v55 int32 = (v3 & v5)
                    v55 -= (v55 * 2147483647)
                    _ = v55
                } else {
                    v56 := v2[209]
                    if false {
                        var v57 float32 = v4
                        if true {
                            if false {
                                v58 := GenS1{F0: v1, F1: 18446744073709551615}
                                v27 -= 67
                                v59 := GenS3{F0: 42, F1: uint(v27), F2: 545}
                                _ = v58
                                _ = v59
                            } else {
                                v60 := GenS1{F0: 2, F1: 4294967295}
                                var v61 float32 = v57
                                v3 = v3
                                v62 := GenS1{F0: 415, F1: 4294967295}
                                _ = v60
                                _ = v61
                                _ = v62
                            }
                        }
                        fmt.Println(fi36, fi42, fi36)
                        _ = v57
                    } else {
                        v1 += v1
                        switch ("foo bar" + "") {
                        case "hello":
                            if true {
                                if len(v28) > 0 {
                                    v28[0] = fi36
                                }
                                v63 := GenS1{F0: (v1 - (^(v1))), F1: 307}
                                fmt.Println(fi42, v27, fi36)
                                _ = v63
                            } else {
                                v4 = float32((1.5 * ((1.5 * 0.0) + 1.5)))
                                v64 := GenS2{F0: 42}
                                v56 += genFn0(v27, v5, v5)
                                _ = v64
                            }
                            if true {
                                var v65 int = (-(v56))
                                if (false || false) {
                                    fmt.Println(v3, fi42)
                                } else {
                                    if false {
                                        if false {
                                            if false {
                                                _ = len(v28)
                                                v27 = v27
                                            }
                                            v66 := v2[genFn0(v27, v5, (v5 - (v3 | -623)))]
                                            v67 := GenS1{F0: 78, F1: 4294967295}
                                            _ = v66
                                            _ = v67
                                        } else {
                                            fmt.Println(v5, v3)
                                            switch genFn0((v27 + (^(v27))), -192, v5) {
                                            case 0:
                                                _ = len(v28)
                                            case -2147483648:
                                                v5 -= (^(v3))
                                                v68 := GenS3{F0: -1, F1: 2, F2: 4294967295}
                                                if (!((((("hello" + "\n\t") >= "\n\t") || false) || false) && false)) {
                                                    var v69 int64 = 134
                                                    v70 := GenS2{F0: 128}
                                                    var v71 int32 = 42
                                                    _ = v69
                                                    _ = v70
                                                    _ = v71
                                                } else {
                                                    v1 = (v1 >> ((2) & 63))
                                                    fmt.Println(v56)
                                                }
                                                v72 := GenS4{F0: -137, F1: (v3 >> ((255) & 63))}
                                                _ = v68
                                                _ = v72
                                            default:
                                                if true {
                                                    if true {
                                                        fmt.Println(v56)
                                                        if len(v28) >= 2 {
                                                            v28 = v28[1:len(v28)]
                                                        }
                                                        if (2 != uint16(v27)) {
                                                            if false {
                                                                switch (v65 - fi42) {
                                                                case 0:
                                                                    v5 = 42
                                                                case 32768:
                                                                    var v73 uint32 = (^(v1))
                                                                    _ = v73
                                                                case -9223372036854775808:
                                                                    switch (-(fi36)) {
                                                                    case 42:
                                                                        if true {
                                                                            if (0 <= 2) {
                                                                                v74 := GenS3{F0: 0, F1: 42, F2: 42}
                                                                                _ = v74
                                                                            } else {
                                                                                switch ("foo bar" + ("hello" + "a")) {
                                                                                case "\n\t":
                                                                                    v75 := v2[fi36]
                                                                                    _ = v75
                                                                                case "a":
                                                                                    var v76 int16 = 1000
                                                                                    _ = v76
                                                                                }
                                                                                if len(v28) > 0 {
                                                                                    v28[0] = (-(v65))
                                                                                }
                                                                            }
                                                                        } else {
                                                                            if false {
                                                                                v3 -= v5
                                                                                v65 = (fi42 * v65)
                                                                                v77 := GenS4{F0: -509, F1: (v5 << ((364) & 63))}
                                                                                var v78 float64 = float64(float32(1e10))
                                                                                switch "foo bar" {
                                                                                case "hello":
                                                                                    var v79 bool = false
                                                                                    _ = v79
                                                                                default:
                                                                                    var v80 bool = (!(!(int16(v56) == -1)))
                                                                                    _ = v80
                                                                                }
                                                                                _ = v77
                                                                                _ = v78
                                                                            } else {
                                                                                switch ("a" + "\n\t") {
                                                                                case "hello":
                                                                                    fmt.Println(fi42)
                                                                                case "":
                                                                                    fmt.Println(fi36, v27, v56)
                                                                                }
                                                                                if len(v28) > 0 {
                                                                                    v28[0] = v56
                                                                                }
                                                                            }
                                                                            v81 := GenS5{F0: int(v1), F1: 32767, F2: 42}
                                                                            var v82 float32 = (((v4 * v4) - (v4 + (0.0 + 1e10))) * (1e10 * (((float32(1.5) / v4) * float32(((0.0 - 1e10) / ((-2.25 / (float64(v4) * 3.14159)) - 0.0)))) + v4)))
                                                                            _ = v81
                                                                            _ = v82
                                                                        }
                                                                    default:
                                                                        var v83 int8 = -1
                                                                        _ = v83
                                                                    }
                                                                }
                                                            } else {
                                                                _ = len(v2)
                                                                var v84 int16 = 0
                                                                _ = v84
                                                            }
                                                            fmt.Println(v65, v5, v3)
                                                        }
                                                    }
                                                    if len(v28) > 0 {
                                                        v28[0] = fi36
                                                    }
                                                    var v85 uint = 255
                                                    _ = v85
                                                }
                                                v86 := GenS2{F0: 42}
                                                if (v4 > (-2.25 - (-2.25 - 1e10))) {
                                                    switch (fi42 << ((317) & 63)) {
                                                    case 127:
                                                        var v87 bool = (1.5 < (float64(v4) * 1.5))
                                                        _ = v87
                                                    case -665:
                                                        v3 -= v3
                                                    case 0:
                                                        var v88 uint32 = 142
                                                        _ = v88
                                                    }
                                                }
                                                _ = v86
                                            }
                                        }
                                        if (!((989 != 1) && true)) {
                                            var v89 int8 = 42
                                            _ = v89
                                        } else {
                                            v56 = fi42
                                            if (v1 != 598) {
                                                if true {
                                                    v56 = fi36
                                                } else {
                                                    v90 := GenS5{F0: (^(v65)), F1: 1000, F2: (v5 * v3)}
                                                    v1 -= v1
                                                    _ = v90
                                                }
                                                switch (fi42 << ((42) & 63)) {
                                                case -590:
                                                    var v91 uint32 = 801
                                                    _ = v91
                                                case 32767:
                                                    var v92 int64 = 9223372036854775807
                                                    _ = v92
                                                case 1:
                                                    v27 += (^(v27))
                                                default:
                                                    switch "" {
                                                    case "\n\t":
                                                        switch "hello" {
                                                        case "a":
                                                            v3 = 1
                                                        }
                                                    case "foo bar":
                                                        delete(v6, "\n\t")
                                                    case "hello":
                                                        var v93 int16 = 767
                                                        _ = v93
                                                    default:
                                                        v94 := GenS4{F0: 1, F1: (v3 & (v3 | v5))}
                                                        _ = v94
                                                    }
                                                }
                                            } else {
                                                var v95 int32 = v5
                                                if ((v4 == 3.14159) && (!true)) {
                                                    var v96 int8 = -128
                                                    delete(v2, fi36)
                                                    _ = v96
                                                }
                                                _ = v95
                                            }
                                            if len(v28) > 0 {
                                                v28[0] = (v65 << ((255) & 63))
                                            }
                                            switch "\n\t" {
                                            case "":
                                                v97 := GenS6{F0: -806, F1: 344, F2: 2}
                                                _ = v97
                                            case "a":
                                                if (32767 > 32767) {
                                                    v98 := GenS7{F0: -32768}
                                                    fmt.Println(v56, v27, v4)
                                                    var v99 int64 = 9223372036854775807
                                                    v100 := GenS6{F0: (v99 << ((0) & 63)), F1: (v27 - (v27 % ((v27) | 1))), F2: uint8(v1)}
                                                    _ = v100
                                                    _ = v98
                                                    _ = v99
                                                }
                                            case "\n\t":
                                                switch 4294967295 {
                                                case 128:
                                                    var v101 uint64 = v27
                                                    _ = v101
                                                case 1:
                                                    v65 = (v56 ^ (fi42 << ((2) & 63)))
                                                case -9223372036854775808:
                                                    if len(v28) > 0 {
                                                        v28[0] = fi36
                                                    }
                                                }
                                                fmt.Println(v5, v3)
                                                v65 = (fi42 / ((v56) | 1))
                                            }
                                            v3 = (v5 ^ -2147483648)
                                            switch 32768 {
                                            case 42:
                                                switch v65 {
                                                case 544:
                                                    switch 9223372036854775807 {
                                                    case -328:
                                                        var v102 uint32 = v1
                                                        _ = v102
                                                    case -2147483648:
                                                        if true {
                                                            if len(v28) > 0 {
                                                                v28[0] = fi36
                                                            }
                                                        } else {
                                                            v103 := GenS8{F0: v1}
                                                            if true {
                                                                switch (v56 >> ((86) & 63)) {
                                                                case 289:
                                                                    v5 -= 1
                                                                }
                                                                v104 := GenS5{F0: fi42, F1: 68, F2: v5}
                                                                v3 = v3
                                                                if ((v4 * v4) > -2.25) {
                                                                    v3 += 1
                                                                    fmt.Println(fi36)
                                                                    fmt.Println(v27, v56, v3)
                                                                } else {
                                                                    var v105 uint8 = 42
                                                                    var v106 int16 = 649
                                                                    v107 := GenS5{F0: (fi42 ^ v56), F1: 1, F2: 0}
                                                                    _ = v105
                                                                    _ = v106
                                                                    _ = v107
                                                                }
                                                                var v108 int = fi42
                                                                _ = v104
                                                                _ = v108
                                                            } else {
                                                                v56 = (-(v65))
                                                                var v109 int8 = -128
                                                                _ = v109
                                                            }
                                                            if (!(1 < 42)) {
                                                                v5 = (v5 >> ((2) & 63))
                                                                v3 += (v3 & v3)
                                                            } else {
                                                                var v110 uint = 1
                                                                v111 := GenS1{F0: 65535, F1: (v110 << ((881) & 63))}
                                                                switch -9223372036854775808 {
                                                                case -140:
                                                                    v110 = (v110 + (^(v110)))
                                                                case -164:
                                                                    var v112 int8 = 100
                                                                    _ = v112
                                                                default:
                                                                    switch (fi36 + (v56 % ((32768) | 1))) {
                                                                    case 953:
                                                                        switch v65 {
                                                                        case 0:
                                                                            fmt.Println(fi36)
                                                                        case 128:
                                                                            v113 := GenS4{F0: -977, F1: 2147483647}
                                                                            _ = v113
                                                                        case 9223372036854775807:
                                                                            var v114 float32 = v4
                                                                            _ = v114
                                                                        default:
                                                                            var v115 uint8 = 90
                                                                            _ = v115
                                                                        }
                                                                    case 0:
                                                                        fmt.Println(v110, v5)
                                                                    case 32768:
                                                                        v65 += -1
                                                                    }
                                                                }
                                                                if (920 >= 32767) {
                                                                    fmt.Println(v65, v27)
                                                                } else {
                                                                    v116 := GenS9{F0: 0, F1: (fi42 | fi36)}
                                                                    v117 := GenS10{F0: 240}
                                                                    _ = v116
                                                                    _ = v117
                                                                }
                                                                _ = v110
                                                                _ = v111
                                                            }
                                                            _ = v103
                                                        }
                                                    case 2147483647:
                                                        v118 := GenS11{F0: 2, F1: v1, F2: v3}
                                                        _ = v118
                                                    default:
                                                        var v119 int = (-(v56))
                                                        _ = v119
                                                    }
                                                default:
                                                    var v120 uint8 = 57
                                                    _ = v120
                                                }
                                            case 32768:
                                                var v121 uint64 = (v27 & (^(v27)))
                                                _ = v121
                                            default:
                                                switch 473 {
                                                case 128:
                                                    var v122 int16 = 1000
                                                    _ = v122
                                                case -52:
                                                    v123 := GenS6{F0: 573, F1: (v27 ^ v27), F2: 127}
                                                    _ = v123
                                                case 1:
                                                    if false {
                                                        v3 += 42
                                                        var v124 int8 = 0
                                                        _ = v124
                                                    } else {
                                                        if (!(!(v3 != int32(v1)))) {
                                                            switch "hello" {
                                                            case "\n\t":
                                                                v3 += v5
                                                            case "a":
                                                                v5 = (v3 | (v5 * v5))
                                                            case "hello":
                                                                var v125 uint64 = (v27 % ((v27) | 1))
                                                                if false {
                                                                    if (100 < int8(fi42)) {
                                                                        var v126 int32 = (v5 + (v5 ^ v3))
                                                                        v6["\n\t"] = (false || ((((!true) && false) && (!true)) || (!true)))
                                                                        _ = v126
                                                                    }
                                                                    if (uint8(v1) != 42) {
                                                                        v1 = (v1 << ((255) & 63))
                                                                        var v127 uint64 = (^(v125))
                                                                        _ = v127
                                                                    } else {
                                                                        var v128 uint8 = 0
                                                                        if (false && ((true || (-128 != 0)) || true)) {
                                                                            v129 := GenS4{F0: 42, F1: -2147483648}
                                                                            switch "hello" {
                                                                            case "foo bar":
                                                                                v130 := v2[v65]
                                                                                _ = v130
                                                                            case "a":
                                                                                v6["\n\t"] = true
                                                                            default:
                                                                                var v131 int16 = -683
                                                                                _ = v131
                                                                            }
                                                                            _ = v129
                                                                        } else {
                                                                            var v132 uint64 = (v27 ^ 55)
                                                                            if false {
                                                                                v3 += v3
                                                                                fmt.Println(v1)
                                                                            }
                                                                            fmt.Println(v1)
                                                                            _ = v132
                                                                        }
                                                                        _ = v128
                                                                    }
                                                                    v3 += (v5 - v5)
                                                                    fmt.Println(v56, v56, fi36)
                                                                }
                                                                v5 += v3
                                                                _ = v125
                                                            default:
                                                                v4 = 3.14159
                                                            }
                                                            var v133 float64 = -2.25
                                                            _ = v133
                                                        }
                                                        switch (fi36 >> ((708) & 63)) {
                                                        case 707:
                                                            var v134 uint8 = 248
                                                            _ = v134
                                                        default:
                                                            var v135 uint16 = 510
                                                            _ = v135
                                                        }
                                                        var v136 int32 = (v5 / ((1) | 1))
                                                        _ = v136
                                                    }
                                                }
                                            }
                                        }
                                        fmt.Println(v5, v4, v1)
                                    } else {
                                        if false {
                                            var v137 uint64 = v27
                                            if true {
                                                var v138 bool = false
                                                if (v1 == v1) {
                                                    switch (fi42 | 644) {
                                                    case -1:
                                                        if len(v28) > 0 {
                                                            v28[0] = fi36
                                                        }
                                                    case 915:
                                                        var v139 int16 = -346
                                                        _ = v139
                                                    }
                                                }
                                                _ = v138
                                            }
                                            var v140 int = 1
                                            _ = v137
                                            _ = v140
                                        } else {
                                            v28 = append(v28, v65)
                                            if false {
                                                v141 := GenS3{F0: int64(v3), F1: 4294967295, F2: 42}
                                                v1 = (v1 * (v1 - v1))
                                                if ((((0.1 / float32(3.14159)) - (v4 + v4)) / (v4 + (v4 * v4))) <= v4) {
                                                    fmt.Println(fi42)
                                                    var v142 int64 = 0
                                                    if ("hello" >= "foo bar") {
                                                        var v143 uint32 = v1
                                                        _ = v143
                                                    }
                                                    _ = v142
                                                } else {
                                                    v1 += (v1 / ((4294967295) | 1))
                                                    var v144 float64 = float64(0.1)
                                                    _ = v144
                                                }
                                                _ = v141
                                            }
                                        }
                                        _ = len(v28)
                                        if false {
                                            switch v65 {
                                            case -904:
                                                var v145 int32 = v3
                                                _ = v145
                                            default:
                                                var v146 int32 = (^(v3))
                                                _ = v146
                                            }
                                        }
                                    }
                                    v4 -= float32((0.1 * 1.5))
                                    if false {
                                        delete(v2, 128)
                                        if (2 >= 127) {
                                            v1 -= v1
                                            if true {
                                                v2[(fi36 % ((0) | 1))] = (-(fi36))
                                                v56 = genFn0(v27, v5, v3)
                                            } else {
                                                var v147 int64 = 1
                                                _ = v147
                                            }
                                            v148 := GenS12{F0: 1, F1: v27}
                                            var v149 int8 = 0
                                            _ = v148
                                            _ = v149
                                        }
                                        if false {
                                            fmt.Println(fi36)
                                            v56 += -561
                                            fmt.Println(v1, v65)
                                            fmt.Println(fi36)
                                        }
                                    } else {
                                        var v150 string = "hello"
                                        v56 = 991
                                        _ = v150
                                    }
                                    var v151 uint = uint(v1)
                                    _ = v151
                                }
                                fmt.Println(v3, v4, v4)
                                switch (^(v56)) {
                                case 256:
                                    var v152 uint16 = 255
                                    _ = v152
                                case 1:
                                    if false {
                                        fmt.Println(v27)
                                    } else {
                                        v153 := GenS13{F0: v1, F1: 2, F2: 0}
                                        _ = v153
                                    }
                                case 127:
                                    if true {
                                        if len(v28) >= 2 {
                                            v28 = v28[1:len(v28)]
                                        }
                                    }
                                default:
                                    if (0 <= -9223372036854775808) {
                                        v56 = -779
                                        fmt.Println(fi36, v56, fi36)
                                    } else {
                                        v154 := GenS5{F0: (v65 / ((2147483647) | 1)), F1: -32768, F2: 0}
                                        _ = v154
                                    }
                                }
                                v27 = v27
                                if true {
                                    v155 := GenS14{F0: -1}
                                    _ = v155
                                }
                                _ = v65
                            } else {
                                v1 = 0
                                var v156 string = ("hello" + ("a" + "a"))
                                switch v156 {
                                case "hello":
                                    switch (fi42 * fi36) {
                                    case 9223372036854775807:
                                        _ = len(v28)
                                        v157 := GenS15{F0: (v3 + -2147483648), F1: fi36, F2: "foo bar"}
                                        _ = v157
                                    case 739:
                                        delete(v2, fi36)
                                    default:
                                        var v158 bool = (!(18446744073709551615 > uint(v27)))
                                        _ = v158
                                    }
                                case "":
                                    v159 := GenS13{F0: v1, F1: 42, F2: 395}
                                    _ = v159
                                }
                                fmt.Println(fi36)
                                var v160 int32 = v5
                                _ = v156
                                _ = v160
                            }
                            var v161 float64 = (((0.0 + 3.14159) + (1.5 - 0.0)) * 0.1)
                            _ = v161
                        }
                    }
                    _ = v56
                }
            }
            _ = fi42
        }
        _ = fi36
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
