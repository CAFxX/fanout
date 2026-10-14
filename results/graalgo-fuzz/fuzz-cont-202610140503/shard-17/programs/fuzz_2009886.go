package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 int64; F1 int8; F2 int8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint; F1 uint32; F2 int64 }
type GenS7 struct { F0 uint16; F1 int16 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint16; F1 int8; F2 int16 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int16; F1 bool; F2 uint }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int64; F1 uint8; F2 uint64 }

func main() {
    var v1 int64 = -893
    var v2 float32 = 1e10
    var v3 uint = 34
    var v4 float32 = (v2 / -2.25)
    v5 := make(map[int]bool)
    var v6 bool = false
    v4 = v2
    var v7 int = 32767
    v8 := []int{v7, (v7 / ((2147483647) | 1)), (v7 - (v7 + v7)), 42}
    delete(v5, 4294967295)
    var v9 int32 = 0
    delete(v5, 32767)
    var v10 uint64 = 649
    switch v7 {
    case 9223372036854775807:
        fmt.Println(v1)
    case 272:
        switch (-(v7)) {
        case -1:
            v11 := GenS1{F0: (^(v9)), F1: false}
            _ = v11
        case 127:
            fmt.Println(v2)
        case 393:
            var v12 string = "foo bar"
            _ = v12
        }
    }
    v2 -= 1.5
    for fi13 := 0; fi13 < 6; fi13++ {
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 int16 = 1
            v16 := v5[(fi14 & (fi13 / ((fi14) | 1)))]
            _ = fi14
            _ = v15
            _ = v16
        }
        v10 = (^(v10))
        _ = fi13
    }
    v7 = v7
    for fi17 := 0; fi17 < 18; fi17++ {
        var v18 int32 = (v9 | v9)
        _ = fi17
        _ = v18
    }
    switch v7 {
    case 1:
        var v19 uint8 = 255
        _ = v19
    case 42:
        for fi20 := 0; fi20 < 12; fi20++ {
            var v21 float32 = -2.25
            _ = fi20
            _ = v21
        }
    default:
        v22 := GenS1{F0: v9, F1: v6}
        _ = v22
    }
    if false {
        var v23 int = 289
        for fi24 := 0; fi24 < 17; fi24++ {
            var v25 uint8 = 128
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = fi24
            _ = v25
        }
        _ = v23
    } else {
        if ((!true) && v6) {
            switch 4294967295 {
            case -9223372036854775808:
                var v26 uint32 = 65535
                _ = v26
            case 0:
                for fi27 := 0; fi27 < 15; fi27++ {
                    var v28 int32 = 584
                    v29 := v5[(v7 & (v7 | int(v10)))]
                    if v6 {
                        v30 := GenS1{F0: -489, F1: (!(32767 > 0))}
                        _ = v30
                    } else {
                        for fi31 := 0; fi31 < 1; fi31++ {
                            fmt.Println(fi31)
                            _ = fi31
                        }
                        v10 = 448
                        var v32 bool = (!(((v28 << (((v3 % (((v3 - 65535)) | 1))) & 63)) < v28) || ((v3 - v3) == (v3 >> ((v3) & 63)))))
                        _ = v32
                    }
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            case -2147483648:
                fmt.Println(v3)
            }
            v6 = (42 >= -164)
            v33 := v5[v7]
            _ = v33
        }
        for fi34 := 0; fi34 < 2; fi34++ {
            for fi35 := 0; fi35 < 6; fi35++ {
                switch "a" {
                case "hello":
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                case "a":
                    if true {
                        v1 = v1
                        v36 := GenS1{F0: (^(v9)), F1: ("" <= "a")}
                        switch (fi34 | fi34) {
                        case 42:
                            var v37 string = "foo bar"
                            _ = v37
                        default:
                            v4 = v4
                        }
                        _ = v36
                    } else {
                        v38 := GenS1{F0: (v9 * v9), F1: false}
                        var v39 float32 = 1.5
                        var v40 uint = 207
                        v7 -= (v7 / ((fi34) | 1))
                        delete(v5, (fi35 | (fi34 / ((-1) | 1))))
                        v7 += -1
                        v40 = v3
                        _ = v38
                        _ = v39
                        _ = v40
                    }
                case "foo bar":
                    var v41 uint8 = uint8(v10)
                    _ = v41
                default:
                    var v42 float64 = 1e10
                    _ = v42
                }
                if true {
                    fmt.Println(fi35)
                    var v43 float64 = 0.1
                    _ = v43
                }
                _ = fi35
            }
            for fi44 := 0; fi44 < 19; fi44++ {
                delete(v5, v7)
                switch ("" + ("\n\t" + "a")) {
                case "hello":
                    v6 = (!false)
                    var v45 int32 = (-(v9))
                    var v46 uint16 = 2
                    var v47 float32 = v2
                    v4 = (0.1 + float32(-2.25))
                    _ = v45
                    _ = v46
                    _ = v47
                case "\n\t":
                    switch (fi44 - v7) {
                    case 0:
                        v4 += 1e10
                    }
                default:
                    fmt.Println(v6, v4)
                }
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = fi44
            }
            if v6 {
                switch (-(fi34)) {
                case 1:
                    switch v7 {
                    case 1:
                        if (((((float32((1e10 - float64(1.5))) - v2) / ((v4 - 3.14159) + (v4 - 1e10))) + ((-2.25 + v4) + ((((float32(1.5) / v4) / (v4 + v2)) + (1.5 - float32(1.5))) / v4))) + (3.14159 / v2)) > 1e10) {
                            if v6 {
                                switch "hello" {
                                case "foo bar":
                                    var v48 uint64 = 4294967295
                                    _ = v48
                                }
                            } else {
                                var v49 int8 = -52
                                var v50 bool = v6
                                _ = v49
                                _ = v50
                            }
                            if false {
                                delete(v5, (v7 * 505))
                                if v6 {
                                    fmt.Println(v4, v1, fi34)
                                    if (255 > 65535) {
                                        v51 := GenS2{F0: v7, F1: 233}
                                        v52 := GenS3{F0: (^(v3)), F1: 65535}
                                        switch "a" {
                                        case "\n\t":
                                            for fi53 := 0; fi53 < 6; fi53++ {
                                                if (!((42 == v3) && (-2.25 < 1e10))) {
                                                    switch fi34 {
                                                    case 0:
                                                        fmt.Println(v9)
                                                    case -464:
                                                        var v54 int64 = 0
                                                        _ = v54
                                                    default:
                                                        if v6 {
                                                            v6 = false
                                                            var v55 uint32 = 2147483648
                                                            var v56 string = (("hello" + "a") + "\n\t")
                                                            v1 = 9223372036854775807
                                                            _ = v55
                                                            _ = v56
                                                        } else {
                                                            if v6 {
                                                                switch v7 {
                                                                case 42:
                                                                    v6 = (((fi53 + fi34) != (-(v7))) || v6)
                                                                case -2147483648:
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        if false {
                                                                            v2 = v2
                                                                        }
                                                                    case "a":
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            if (0 <= 42) {
                                                                                v7 -= fi53
                                                                                v6 = ((uint8(v3) != 185) && v6)
                                                                            }
                                                                        }
                                                                    case "":
                                                                        v1 -= (v1 * (v1 << ((v3) & 63)))
                                                                    }
                                                                }
                                                                if ((!((v3 | v3) != (v3 ^ 0))) && (1 >= int16(fi34))) {
                                                                    var v57 float32 = 1.5
                                                                    _ = v57
                                                                }
                                                                var v58 uint16 = 65535
                                                                _ = v58
                                                            } else {
                                                                v4 = ((float32(0.0) - (((v2 * 3.14159) * (float32(-2.25) + v2)) - v2)) / v2)
                                                            }
                                                            if (true && true) {
                                                                var v59 uint8 = 0
                                                                _ = v59
                                                            } else {
                                                                _ = len(v5)
                                                                if (!((true && true) && (v3 < v3))) {
                                                                    v60 := GenS4{F0: v1, F1: 106, F2: -1}
                                                                    v6 = v6
                                                                    _ = len(v8)
                                                                    _ = v60
                                                                } else {
                                                                    var v61 bool = (v10 <= v10)
                                                                    switch -1 {
                                                                    case 32767:
                                                                        if ((1 >= 1) && (false || false)) {
                                                                            v3 = v3
                                                                            v10 -= (v10 >> ((v3) & 63))
                                                                        } else {
                                                                            var v62 uint8 = 255
                                                                            _ = v62
                                                                        }
                                                                    }
                                                                    _ = v61
                                                                }
                                                                var v63 bool = (v6 || (true || v6))
                                                                _ = v63
                                                            }
                                                        }
                                                        v6 = (((float64(v4) / 1e10) <= -2.25) || (!true))
                                                    }
                                                    v64 := GenS1{F0: v9, F1: (false && (42 != 65535))}
                                                    _ = v64
                                                }
                                                var v65 uint16 = 42
                                                switch 549 {
                                                case -1:
                                                    if len(v8) >= 2 {
                                                        v8 = v8[1:len(v8)]
                                                    }
                                                case 251:
                                                    switch (fi34 % ((fi53) | 1)) {
                                                    case 2147483647:
                                                        if ((v9 * int32(v65)) < 0) {
                                                            delete(v5, v7)
                                                        } else {
                                                            switch "a" {
                                                            case "foo bar":
                                                                switch (v7 & (fi34 + fi53)) {
                                                                case 127:
                                                                    v66 := GenS5{F0: 1}
                                                                    _ = v66
                                                                case 128:
                                                                    if false {
                                                                        v5[(v7 / ((v7) | 1))] = (0.1 <= 1.5)
                                                                    }
                                                                case 587:
                                                                    fmt.Println(v65)
                                                                default:
                                                                    v2 = v2
                                                                }
                                                            case "":
                                                                v65 = uint16(v3)
                                                            default:
                                                                v3 = (v3 << ((65535) & 63))
                                                            }
                                                            switch ("\n\t" + "") {
                                                            case "foo bar":
                                                                if (!false) {
                                                                    v67 := GenS6{F0: 1, F1: 210, F2: (v1 >> ((v3) & 63))}
                                                                    switch v7 {
                                                                    case 128:
                                                                        var v68 uint = (v3 & 575)
                                                                        _ = v68
                                                                    case 9223372036854775807:
                                                                        var v69 uint8 = 1
                                                                        _ = v69
                                                                    case -790:
                                                                        if false {
                                                                            fmt.Println(v9, v4)
                                                                        }
                                                                    }
                                                                    _ = v67
                                                                } else {
                                                                    v8 = append(v8, (v7 >> ((v3) & 63)))
                                                                }
                                                            }
                                                        }
                                                    case 42:
                                                        v65 = v65
                                                    }
                                                    if v6 {
                                                        if (true && (!false)) {
                                                            v65 += v65
                                                            if ((!v6) || ((v1 & -624) < (v1 & (v1 ^ (v1 / ((-9223372036854775808) | 1)))))) {
                                                                v5[(-(fi53))] = (127 > 255)
                                                                var v70 string = "\n\t"
                                                                _ = v70
                                                            } else {
                                                                var v71 bool = ((v1 == (v1 ^ (v1 ^ (v1 % (((^(v1))) | 1))))) && ((float32((((((3.14159 - -2.25) * 3.14159) + 0.0) * (((0.0 * 0.1) + (1.5 + (0.1 - -2.25))) - 1.5)) / ((1.5 / (1.5 + 1.5)) + -2.25))) - 1.5) == -2.25))
                                                                _ = v71
                                                            }
                                                            var v72 uint16 = (^(v65))
                                                            fmt.Println(v1, v9, v6)
                                                            _ = v72
                                                        } else {
                                                            v3 = v3
                                                            v73 := GenS7{F0: (v65 & v65), F1: -1}
                                                            _ = v73
                                                        }
                                                        v74 := v5[4294967295]
                                                        _ = v74
                                                    } else {
                                                        v3 = 65535
                                                        fmt.Println(v3, fi53, fi53)
                                                        fmt.Println(v3)
                                                    }
                                                default:
                                                    v75 := GenS8{F0: -128}
                                                    _ = v75
                                                }
                                                var v76 int16 = 1
                                                _ = fi53
                                                _ = v65
                                                _ = v76
                                            }
                                        }
                                        _ = v51
                                        _ = v52
                                    } else {
                                        switch "hello" {
                                        case "a":
                                            for fi77 := 0; fi77 < 9; fi77++ {
                                                switch -492 {
                                                case 32767:
                                                    v78 := GenS9{F0: 1, F1: -92, F2: 328}
                                                    v10 += 672
                                                    if ((-32768 > 1) || (float32(3.14159) > (3.14159 + 1e10))) {
                                                        var v79 float64 = 1e10
                                                        _ = v79
                                                    } else {
                                                        v80 := GenS8{F0: 100}
                                                        _ = v80
                                                    }
                                                    _ = v78
                                                case 829:
                                                    var v81 int = (-(v7))
                                                    _ = v81
                                                default:
                                                    var v82 int32 = (v9 - 42)
                                                    _ = v82
                                                }
                                                var v83 uint64 = 18446744073709551615
                                                _ = fi77
                                                _ = v83
                                            }
                                        case "hello":
                                            var v84 uint16 = 0
                                            _ = v84
                                        case "foo bar":
                                            fmt.Println(v4, fi34, v4)
                                        default:
                                            for fi85 := 0; fi85 < 10; fi85++ {
                                                if (true || v6) {
                                                    fmt.Println(v2, v9, v9)
                                                    var v86 int8 = 42
                                                    _ = v86
                                                }
                                                if ((v9 + (v9 | v9)) >= 834) {
                                                    switch 424 {
                                                    case 1:
                                                        var v87 int64 = -9223372036854775808
                                                        _ = v87
                                                    }
                                                }
                                                _ = fi85
                                            }
                                        }
                                    }
                                    var v88 float64 = (float64(float32(0.0)) - 1e10)
                                    _ = v88
                                }
                            }
                            v7 = v7
                        }
                    case 127:
                        for fi89 := 0; fi89 < 12; fi89++ {
                            switch "foo bar" {
                            case "hello":
                                if (v4 >= (v2 * 0.1)) {
                                    var v90 uint8 = 128
                                    var v91 int64 = v1
                                    _ = v90
                                    _ = v91
                                } else {
                                    fmt.Println(v6)
                                    if v6 {
                                        switch ("a" + "foo bar") {
                                        case "foo bar":
                                            fmt.Println(v4)
                                        case "a":
                                            delete(v5, fi34)
                                        }
                                    } else {
                                        if v6 {
                                            var v92 uint64 = (v10 / ((v10) | 1))
                                            switch "foo bar" {
                                            case "hello":
                                                v93 := v5[-2147483648]
                                                _ = v93
                                            default:
                                                var v94 uint64 = (v92 / ((v10) | 1))
                                                _ = v94
                                            }
                                            v92 += v10
                                            _ = v92
                                        } else {
                                            v10 = v10
                                            if len(v8) > 0 {
                                                v8[0] = (-(fi34))
                                            }
                                        }
                                    }
                                }
                            case "foo bar":
                                if v6 {
                                    var v95 int32 = -1
                                    var v96 float32 = ((v4 * v2) + (1.5 - v4))
                                    _ = v95
                                    _ = v96
                                } else {
                                    delete(v5, (v7 ^ v7))
                                }
                            case "":
                                fmt.Println(v7, v9, fi89)
                            default:
                                _ = len(v5)
                            }
                            _ = fi89
                        }
                    default:
                        switch 0 {
                        case 936:
                            v97 := GenS1{F0: (v9 / ((v9) | 1)), F1: true}
                            _ = v97
                        case 32768:
                            if (v1 > v1) {
                                for fi98 := 0; fi98 < 20; fi98++ {
                                    if v6 {
                                        switch ("a" + "hello") {
                                        case "":
                                            var v99 string = "\n\t"
                                            _ = v99
                                        default:
                                            v100 := GenS10{F0: (v9 & 1)}
                                            _ = v100
                                        }
                                        if ((2 == 45) || (!true)) {
                                            fmt.Println(v9, v1, fi98)
                                            v1 = 9223372036854775807
                                        } else {
                                            v101 := GenS3{F0: v3, F1: uint16(v10)}
                                            v102 := v5[-9223372036854775808]
                                            var v103 int32 = 42
                                            switch fi34 {
                                            case 42:
                                                if len(v8) > 0 {
                                                    v8[0] = (fi98 << ((v3) & 63))
                                                }
                                            case -790:
                                                if v6 {
                                                    var v104 float32 = float32(float64((0.1 - v4)))
                                                    fmt.Println(v10, v9)
                                                    switch (fi98 | 42) {
                                                    case 32768:
                                                        switch "\n\t" {
                                                        case "a":
                                                            v10 -= v10
                                                        }
                                                    case 136:
                                                        switch fi34 {
                                                        case -2147483648:
                                                            v105 := v5[fi98]
                                                            v106 := GenS11{F0: -1, F1: v105, F2: v3}
                                                            switch (fi98 + (fi98 << ((2) & 63))) {
                                                            case 42:
                                                                fmt.Println(v9)
                                                            case 32768:
                                                                var v107 bool = true
                                                                _ = v107
                                                            case 4294967295:
                                                                v2 += 1e10
                                                            default:
                                                                fmt.Println(fi34)
                                                            }
                                                            _ = v105
                                                            _ = v106
                                                        case 9223372036854775807:
                                                            if v6 {
                                                                v108 := GenS12{F0: ((("" + ("\n\t" + "")) + "\n\t") + "a")}
                                                                _ = v108
                                                            }
                                                        }
                                                    default:
                                                        if (!(((-2.25 + -2.25) / 1e10) >= (1e10 + 1e10))) {
                                                            var v109 int32 = (v103 & (v9 >> ((v3) & 63)))
                                                            v102 = false
                                                            _ = v109
                                                        }
                                                    }
                                                    _ = v104
                                                } else {
                                                    var v110 uint64 = (^(v10))
                                                    _ = v110
                                                }
                                            }
                                            _ = v101
                                            _ = v102
                                            _ = v103
                                        }
                                    }
                                    _ = fi98
                                }
                                var v111 int = 32767
                                _ = v111
                            } else {
                                v6 = v6
                                var v112 uint64 = v10
                                _ = v112
                            }
                        case 127:
                            if (("hello" + ("foo bar" + "a")) >= (("foo bar" + "") + "a")) {
                                var v113 float64 = (1e10 / 3.14159)
                                _ = v113
                            } else {
                                fmt.Println(fi34)
                            }
                        default:
                            var v114 int = (v7 & v7)
                            _ = v114
                        }
                    }
                case 32768:
                    v2 -= v2
                }
                delete(v5, (v7 * fi34))
                v3 = v3
            } else {
                if false {
                    var v115 int32 = (v9 << (((v3 + v3)) & 63))
                    _ = v115
                }
                v116 := GenS13{F0: v1, F1: 1, F2: (v10 / ((v10) | 1))}
                if len(v8) > 0 {
                    v8[0] = v7
                }
                _ = v116
            }
            _ = fi34
        }
    }
    v6 = (!v6)
    var v117 int32 = 1
    for fi118 := 0; fi118 < 13; fi118++ {
        _ = len(v8)
        _ = fi118
    }
    var v119 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v117:", v117)
    fmt.Println("v119:", v119)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v117
    _ = v119
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
