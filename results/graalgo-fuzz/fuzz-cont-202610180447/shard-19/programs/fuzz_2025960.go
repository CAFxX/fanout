package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 int64; F1 int32; F2 uint8 }
type GenS6 struct { F0 bool; F1 uint }
type GenS7 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS8 struct { F0 int; F1 uint64; F2 uint }
type GenS9 struct { F0 string; F1 int8 }
type GenS10 struct { F0 int64; F1 string }
type GenS11 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS12 struct { F0 int32; F1 uint32; F2 uint }
type GenS13 struct { F0 int8; F1 int8 }
type GenS14 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS15 struct { F0 int8; F1 int; F2 uint64 }
type GenS16 struct { F0 bool; F1 int32 }
type GenS17 struct { F0 int64; F1 int16 }
type GenS18 struct { F0 uint64; F1 string; F2 uint }
type GenS19 struct { F0 uint }
type GenS20 struct { F0 int8; F1 string; F2 int8 }

func main() {
    v1 := []uint32{2}
    v2 := make(map[uint32]bool)
    var v3 uint16 = 2
    v4 := []uint32{4294967295, 65535, 1}
    v5 := []int{32768, -683, 0, 0}
    if true {
        var v6 float32 = -2.25
        for fi7 := 0; fi7 < 6; fi7++ {
            switch -328 {
            case 9223372036854775807:
                v6 = (v6 / ((v6 / 0.1) + v6))
            case -9223372036854775808:
                v8 := GenS1{F0: fi7}
                _ = v8
            default:
                var v9 float64 = 3.14159
                _ = v9
            }
            fmt.Println(v6, v3)
            delete(v2, 914)
            _ = fi7
        }
        _ = v6
    }
    if false {
        switch 2147483647 {
        case 366:
            if true {
                v1 = append(v1, 42)
                for fi10 := 0; fi10 < 9; fi10++ {
                    if (("hello" < "foo bar") || (-1 != 70)) {
                        v3 += (v3 + v3)
                    } else {
                        if true {
                            var v11 int16 = -425
                            _ = v11
                        } else {
                            var v12 uint64 = 2
                            var v13 float32 = (1.5 + -2.25)
                            _ = v12
                            _ = v13
                        }
                        v3 = 42
                    }
                    var v14 uint8 = 2
                    _ = fi10
                    _ = v14
                }
            } else {
                v3 -= 0
                v3 -= (v3 - v3)
            }
        case 1:
            v15 := GenS1{F0: -401}
            _ = v15
        }
        v3 = v3
        fmt.Println(v3)
    }
    if len(v4) > 0 {
        v4[0] = 0
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        fmt.Println(fi16, fi16)
        v3 += v3
        _ = fi16
    }
    switch "" {
    case "":
        for fi17 := 0; fi17 < 19; fi17++ {
            switch "" {
            case "foo bar":
                fmt.Println(fi17, fi17)
            case "":
                v18 := GenS1{F0: (fi17 % ((fi17) | 1))}
                _ = v18
            default:
                for fi19 := 0; fi19 < 17; fi19++ {
                    if (128 >= 255) {
                        if (!true) {
                            if (2147483647 <= 0) {
                                v20 := GenS2{F0: 255, F1: false}
                                if (!(false && (42 >= 85))) {
                                    switch "hello" {
                                    case "a":
                                        v3 += (v3 + v3)
                                    }
                                    switch "a" {
                                    case "a":
                                        v21 := GenS3{F0: 42, F1: 1}
                                        _ = v21
                                    default:
                                        delete(v2, 958)
                                    }
                                }
                                _ = v20
                            } else {
                                if (v3 < (v3 | (v3 - 2))) {
                                    if true {
                                        var v22 int64 = 42
                                        if (!false) {
                                            switch ((("" + ("a" + "\n\t")) + "hello") + "") {
                                            case "a":
                                                if false {
                                                    var v23 float64 = (((0.0 / ((-2.25 + (1e10 - (-2.25 / 3.14159))) / 1.0)) + 1.5) + (((0.1 - 3.14159) * (float64(3.14159) + 0.0)) + -2.25))
                                                    _ = v23
                                                } else {
                                                    v24 := GenS4{F0: -128, F1: v3}
                                                    _ = v24
                                                }
                                            case "":
                                                switch ("a" + "hello") {
                                                case "foo bar":
                                                    fmt.Println(fi17)
                                                case "\n\t":
                                                    v25 := GenS5{F0: (v22 | (v22 ^ 975)), F1: 608, F2: 42}
                                                    _ = v25
                                                case "a":
                                                    v22 = (v22 - v22)
                                                }
                                            default:
                                                var v26 int16 = 1000
                                                _ = v26
                                            }
                                        } else {
                                            v22 = 42
                                            v27 := v2[0]
                                            v2[4294967295] = true
                                            var v28 bool = v27
                                            _ = v27
                                            _ = v28
                                        }
                                        if (v3 == v3) {
                                            if false {
                                                _ = len(v2)
                                                v29 := GenS3{F0: 721, F1: 127}
                                                var v30 bool = (0 >= 414)
                                                _ = len(v5)
                                                _ = v29
                                                _ = v30
                                            } else {
                                                fmt.Println(fi19)
                                                if false {
                                                    fmt.Println(fi17, fi17)
                                                } else {
                                                    var v31 int8 = 127
                                                    switch ("\n\t" + "hello") {
                                                    case "hello":
                                                        _ = len(v1)
                                                    case "foo bar":
                                                        var v32 bool = false
                                                        _ = v32
                                                    case "\n\t":
                                                        v22 += -1
                                                    }
                                                    v22 -= (^(v22))
                                                    _ = v31
                                                }
                                                v22 -= (v22 % ((-1) | 1))
                                            }
                                            if (!(int64(v3) <= (v22 & (v22 >> ((4294967295) & 63))))) {
                                                if false {
                                                    _ = len(v2)
                                                }
                                                v3 += v3
                                            }
                                            v33 := GenS2{F0: 2, F1: ((4294967295 <= 0) && (!false))}
                                            switch "foo bar" {
                                            case "\n\t":
                                                var v34 int16 = 42
                                                _ = v34
                                            case "foo bar":
                                                v3 = v3
                                            }
                                            _ = v33
                                        } else {
                                            var v35 float64 = 0.1
                                            _ = v35
                                        }
                                        _ = v22
                                    } else {
                                        if (42 != 42) {
                                            fmt.Println(fi19)
                                            var v36 string = ("hello" + (("foo bar" + "") + "a"))
                                            switch fi19 {
                                            case -901:
                                                v3 += v3
                                            case -746:
                                                v37 := GenS2{F0: 57, F1: true}
                                                _ = v37
                                            default:
                                                if (65535 < 595) {
                                                    var v38 uint = 18446744073709551615
                                                    if true {
                                                        v3 += (v3 << ((v38) & 63))
                                                        var v39 uint8 = 239
                                                        _ = v39
                                                    } else {
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            v40 := GenS6{F0: true, F1: 183}
                                                            _ = v40
                                                        case "hello":
                                                            if ((-2147483648 < -2147483648) && false) {
                                                                switch v36 {
                                                                case "hello":
                                                                    v41 := GenS4{F0: 42, F1: (^(v3))}
                                                                    _ = v41
                                                                case "":
                                                                    delete(v2, 2)
                                                                    var v42 float64 = 0.0
                                                                    v43 := GenS2{F0: 42, F1: true}
                                                                    var v44 int16 = 1
                                                                    _ = v42
                                                                    _ = v43
                                                                    _ = v44
                                                                case "foo bar":
                                                                    var v45 int8 = 42
                                                                    _ = v45
                                                                default:
                                                                    fmt.Println(fi19, v3, v38)
                                                                }
                                                                if (-9223372036854775808 != -1) {
                                                                    var v46 uint64 = 0
                                                                    var v47 uint16 = (v3 % ((2) | 1))
                                                                    v48 := GenS4{F0: -41, F1: v3}
                                                                    _ = v46
                                                                    _ = v47
                                                                    _ = v48
                                                                } else {
                                                                    var v49 bool = (127 < 128)
                                                                    v49 = (1 <= v38)
                                                                    v5 = append(v5, (fi17 ^ (fi17 ^ fi17)))
                                                                    _ = v49
                                                                }
                                                            }
                                                        default:
                                                            fmt.Println(v3, v3, v36)
                                                        }
                                                        v38 += 316
                                                        v36 = "hello"
                                                    }
                                                    var v50 uint = v38
                                                    if (false && (v36 >= v36)) {
                                                        v51 := GenS7{F0: (v38 * (v38 << (((v50 - (v50 >> (((v38 / ((v50) | 1))) & 63)))) & 63))), F1: (v3 | (v3 + (v3 / (((v3 | uint16(v50))) | 1)))), F2: v3}
                                                        _ = v51
                                                    } else {
                                                        v3 = 255
                                                        v52 := GenS8{F0: fi19, F1: 183, F2: (v50 - (v38 / ((4294967295) | 1)))}
                                                        var v53 int64 = 9223372036854775807
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            var v54 int = fi17
                                                            _ = v54
                                                        case "a":
                                                            v55 := GenS9{F0: v36, F1: 27}
                                                            _ = v55
                                                        case "\n\t":
                                                            var v56 uint8 = 42
                                                            _ = v56
                                                        }
                                                        v2[1] = (3.14159 == 3.14159)
                                                        if false {
                                                            var v57 int32 = 2147483647
                                                            _ = v57
                                                        }
                                                        if true {
                                                            var v58 uint8 = 1
                                                            v59 := GenS6{F0: true, F1: (v38 + 1)}
                                                            v36 = "foo bar"
                                                            _ = v58
                                                            _ = v59
                                                        }
                                                        switch -2147483648 {
                                                        case 4294967295:
                                                            var v60 int32 = -2147483648
                                                            _ = v60
                                                        case 42:
                                                            v61 := GenS4{F0: -54, F1: (v3 ^ v3)}
                                                            _ = v61
                                                        }
                                                        v62 := GenS6{F0: false, F1: 13}
                                                        var v63 string = v36
                                                        _ = v52
                                                        _ = v53
                                                        _ = v62
                                                        _ = v63
                                                    }
                                                    switch (fi19 + (-(fi19))) {
                                                    case 4294967295:
                                                        fmt.Println(v38, fi17, fi19)
                                                    case -359:
                                                        v50 -= v50
                                                    case -9223372036854775808:
                                                        v64 := GenS10{F0: -9223372036854775808, F1: v36}
                                                        _ = v64
                                                    default:
                                                        switch v36 {
                                                        case "a":
                                                            if (1 <= 0) {
                                                                v65 := GenS3{F0: 1, F1: 111}
                                                                _ = v65
                                                            }
                                                        case "hello":
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        }
                                                    }
                                                    v3 = v3
                                                    _ = v38
                                                    _ = v50
                                                } else {
                                                    v3 -= 65535
                                                    v2[229] = true
                                                }
                                            }
                                            _ = v36
                                        }
                                        v5 = append(v5, fi19)
                                        fmt.Println(fi17, v3, fi19)
                                    }
                                }
                                switch (fi17 ^ (fi19 - fi17)) {
                                case 32767:
                                    switch 127 {
                                    case -9223372036854775808:
                                        switch ((("hello" + (("hello" + ("\n\t" + (("a" + "") + "foo bar"))) + "hello")) + "hello") + "\n\t") {
                                        case "":
                                            if true {
                                                var v66 bool = (false && true)
                                                _ = v66
                                            } else {
                                                v3 += (^(v3))
                                                var v67 int16 = 1
                                                switch fi19 {
                                                case -9223372036854775808:
                                                    _ = len(v2)
                                                case 32767:
                                                    v68 := GenS11{F0: 235, F1: 127, F2: 4294967295}
                                                    _ = v68
                                                default:
                                                    v67 += 32767
                                                    v69 := GenS1{F0: fi19}
                                                    v67 = v67
                                                    _ = v69
                                                }
                                                _ = v67
                                            }
                                        }
                                    case 1:
                                        if true {
                                            var v70 int8 = -30
                                            v71 := GenS12{F0: 589, F1: 65535, F2: 255}
                                            fmt.Println(v3)
                                            _ = len(v2)
                                            v70 -= 127
                                            _ = v70
                                            _ = v71
                                        } else {
                                            var v72 float32 = 1.5
                                            _ = len(v2)
                                            var v73 uint64 = 49
                                            _ = v72
                                            _ = v73
                                        }
                                    case 9223372036854775807:
                                        if true {
                                            v3 = v3
                                        }
                                    default:
                                        var v74 uint8 = 42
                                        _ = v74
                                    }
                                case 128:
                                    if ("\n\t" >= "") {
                                        if len(v5) > 0 {
                                            v5[0] = (fi19 * fi19)
                                        }
                                    }
                                case 32768:
                                    v3 -= v3
                                default:
                                    var v75 string = ("a" + "a")
                                    _ = v75
                                }
                                if (1000 > -32768) {
                                    fmt.Println(fi19, fi19)
                                    var v76 float64 = 0.1
                                    var v77 float64 = ((0.1 * (((v76 * 0.1) - v76) * (1e10 + v76))) - ((v76 / ((1e10 + -2.25) - v76)) * v76))
                                    _ = v76
                                    _ = v77
                                }
                            }
                            var v78 uint64 = 1
                            switch (("foo bar" + "hello") + (("\n\t" + ("foo bar" + ("hello" + (("foo bar" + "a") + "\n\t")))) + "hello")) {
                            case "a":
                                fmt.Println(v3, fi19, v3)
                            case "hello":
                                v1 = append(v1, 2147483648)
                            }
                            _ = v78
                        }
                        v79 := GenS13{F0: -1, F1: -120}
                        fmt.Println(v3, v3, v3)
                        _ = v79
                    } else {
                        v3 -= (v3 >> ((0) & 63))
                        switch (fi17 / (((fi17 * (fi19 - 463))) | 1)) {
                        case -1:
                            var v80 int64 = 0
                            _ = v80
                        case 9223372036854775807:
                            fmt.Println(fi17, fi19)
                        case 127:
                            switch "\n\t" {
                            case "foo bar":
                                v81 := GenS5{F0: -135, F1: 42, F2: 19}
                                _ = v81
                            case "\n\t":
                                v3 -= 255
                            }
                        default:
                            _ = len(v4)
                        }
                        delete(v2, 0)
                    }
                    var v82 uint8 = 0
                    v4 = append(v4, 1)
                    _ = fi19
                    _ = v82
                }
            }
            _ = len(v4)
            _ = fi17
        }
    case "\n\t":
        var v83 uint16 = 32767
        _ = v83
    case "foo bar":
        for fi84 := 0; fi84 < 7; fi84++ {
            var v85 uint64 = 1
            var v86 int32 = -2147483648
            v85 = v85
            _ = fi84
            _ = v85
            _ = v86
        }
    }
    var v87 string = "foo bar"
    switch 32767 {
    case 42:
        switch 128 {
        case 0:
            switch v87 {
            case "":
                var v88 bool = false
                _ = v88
            case "\n\t":
                fmt.Println(v3, v87)
            case "foo bar":
                switch "hello" {
                case "foo bar":
                    var v89 string = ("\n\t" + ((v87 + ("\n\t" + ((v87 + ((("" + v87) + (("a" + v87) + ((v87 + ((v87 + v87) + v87)) + "foo bar"))) + v87)) + v87))) + (v87 + ("a" + v87))))
                    _ = v89
                case "\n\t":
                    if (!true) {
                        var v90 int64 = -94
                        v3 += 277
                        _ = v90
                    }
                }
            }
        case -9223372036854775808:
            v87 = "\n\t"
        default:
            v87 = v87
        }
    }
    switch 32768 {
    case 674:
        v3 = (v3 - 2)
    default:
        v87 = ""
    }
    for fi91 := 0; fi91 < 3; fi91++ {
        if ((v87 + v87) > "a") {
            var v92 uint = 255
            _ = v92
        }
        if (1 != 0) {
            switch fi91 {
            case 42:
                if true {
                    fmt.Println(v87)
                }
            case 9223372036854775807:
                v93 := GenS8{F0: fi91, F1: 1, F2: 65535}
                _ = v93
            default:
                for fi94 := 0; fi94 < 18; fi94++ {
                    v95 := GenS5{F0: 17, F1: 448, F2: 0}
                    if true {
                        var v96 float32 = 0.1
                        if false {
                            var v97 string = v87
                            var v98 uint64 = 18446744073709551615
                            _ = v97
                            _ = v98
                        } else {
                            v3 += (^(v3))
                            v96 -= v96
                        }
                        _ = v96
                    }
                    _ = fi94
                    _ = v95
                }
            }
            if (!(1 != 65535)) {
                for fi99 := 0; fi99 < 15; fi99++ {
                    v100 := GenS14{F0: 452, F1: 2, F2: -57}
                    v87 = v87
                    var v101 uint32 = 973
                    _ = fi99
                    _ = v100
                    _ = v101
                }
            } else {
                v3 += v3
                if (((^(v3)) < v3) || ((-(fi91)) <= 127)) {
                    for fi102 := 0; fi102 < 16; fi102++ {
                        var v103 uint8 = 205
                        if false {
                            if false {
                                var v104 int8 = 0
                                v105 := GenS6{F0: true, F1: 67}
                                _ = v104
                                _ = v105
                            }
                            if true {
                                v87 = v87
                                switch fi102 {
                                case 32768:
                                    var v106 uint = 648
                                    _ = v106
                                }
                            } else {
                                v107 := GenS15{F0: 100, F1: (fi102 % ((fi102) | 1)), F2: uint64(v3)}
                                var v108 bool = (v3 > v3)
                                if len(v5) > 0 {
                                    v5[0] = (^(fi102))
                                }
                                _ = v107
                                _ = v108
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            if (-2.25 <= -2.25) {
                                if (170 <= 255) {
                                    var v109 float32 = 3.14159
                                    fmt.Println(v103)
                                    delete(v2, 2147483648)
                                    _ = v109
                                }
                                switch fi91 {
                                case -760:
                                    _ = len(v4)
                                }
                                v110 := GenS16{F0: false, F1: -880}
                                _ = v110
                            }
                            switch ("a" + ((((v87 + v87) + (v87 + (v87 + v87))) + v87) + v87)) {
                            case "":
                                if false {
                                    fmt.Println(v3, fi91, v87)
                                    if false {
                                        v111 := GenS17{F0: 0, F1: -1}
                                        v3 += (^(v3))
                                        _ = v111
                                    } else {
                                        var v112 int8 = -52
                                        var v113 int64 = 485
                                        if len(v4) > 0 {
                                            v4[0] = 2
                                        }
                                        _ = v112
                                        _ = v113
                                    }
                                    v87 = "foo bar"
                                }
                            }
                            _ = len(v2)
                        }
                        v114 := GenS18{F0: 0, F1: v87, F2: uint(v103)}
                        v3 = (v3 >> ((2) & 63))
                        _ = fi102
                        _ = v103
                        _ = v114
                    }
                    var v115 int8 = -1
                    v115 += (-(v115))
                    _ = v115
                } else {
                    var v116 uint16 = 32767
                    _ = v116
                }
            }
            switch v87 {
            case "hello":
                v2[499] = ((((-2.25 / -2.25) / (3.14159 * ((1e10 / float32(-2.25)) + 1.5))) / (1.5 * ((1e10 * (((0.0 / 1.5) - 1e10) / 1.5)) / 3.14159))) < 3.14159)
            default:
                if len(v5) > 0 {
                    v5[0] = 576
                }
            }
        } else {
            for fi117 := 0; fi117 < 1; fi117++ {
                switch ((("\n\t" + v87) + ("a" + "\n\t")) + v87) {
                case "\n\t":
                    fmt.Println(v3)
                case "hello":
                    var v118 uint64 = 42
                    _ = v118
                default:
                    var v119 uint16 = (v3 | (v3 << ((65535) & 63)))
                    _ = v119
                }
                switch v87 {
                case "\n\t":
                    if true {
                        v3 += v3
                        switch v87 {
                        case "hello":
                            v120 := GenS19{F0: 1}
                            _ = v120
                        case "a":
                            v87 = "\n\t"
                        default:
                            var v121 int8 = 33
                            _ = v121
                        }
                    } else {
                        if true {
                            fmt.Println(fi117, fi91, v3)
                        }
                    }
                case "foo bar":
                    _ = len(v2)
                case "a":
                    var v122 int16 = -268
                    _ = v122
                default:
                    var v123 float32 = 1.5
                    _ = v123
                }
                v124 := GenS20{F0: 0, F1: (v87 + "foo bar"), F2: 127}
                delete(v2, 220)
                _ = fi117
                _ = v124
            }
            v4 = append(v4, 2)
            v87 = v87
        }
        _ = fi91
    }
    var v125 float64 = 3.14159
    v126 := GenS4{F0: -128, F1: v3}
    v87 = (("foo bar" + v87) + "a")
    var v127 uint = 42
    switch -197 {
    case 128:
        var v128 uint8 = 2
        _ = v128
    case 482:
        if true {
            v129 := GenS9{F0: ("foo bar" + v87), F1: 92}
            _ = v129
        } else {
            switch ("hello" + (v87 + "foo bar")) {
            case "\n\t":
                var v130 string = ("a" + "\n\t")
                _ = v130
            }
        }
    case -894:
        var v131 uint64 = 1
        _ = v131
    }
    for fi132 := 0; fi132 < 16; fi132++ {
        var v133 string = "a"
        v134 := GenS13{F0: -128, F1: 42}
        var v135 int8 = -87
        _ = fi132
        _ = v133
        _ = v134
        _ = v135
    }
    var v136 float64 = (((v125 * -2.25) + 0.0) + v125)
    v3 = (v3 % ((v3) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v125:", v125)
    fmt.Println("v127:", v127)
    fmt.Println("v136:", v136)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v125
    _ = v126
    _ = v127
    _ = v136
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v87
}
