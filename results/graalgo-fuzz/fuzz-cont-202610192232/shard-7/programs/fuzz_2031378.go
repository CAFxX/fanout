package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int16; F2 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint; F1 uint32; F2 uint }
type GenS8 struct { F0 int16; F1 int16 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint16; F1 int8; F2 bool }
type GenS11 struct { F0 int; F1 int32 }
type GenS12 struct { F0 int16; F1 int }
type GenS13 struct { F0 int16; F1 int8; F2 int16 }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 uint; F1 int32 }
type GenS16 struct { F0 uint }
type GenS17 struct { F0 int64; F1 string }
type GenS18 struct { F0 uint64; F1 int16; F2 int16 }
func genFn0() uint16 {
    if false {
        var v6 uint32 = 725
        v6 = (v6 | (v6 - 42))
        _ = v6
    }
    v7 := []int{-322, 42, 2147483647}
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 int64 = 31
        if (!(0 > -2147483648)) {
            if true {
                v9 -= v9
            } else {
                var v10 float32 = 0.1
                for fi11 := 0; fi11 < 16; fi11++ {
                    if len(v7) > 0 {
                        v7[0] = fi8
                    }
                    _ = fi11
                }
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = v10
            }
            var v12 float64 = -2.25
            switch ("hello" + (("hello" + "foo bar") + "foo bar")) {
            case "hello":
                switch 891 {
                case 103:
                    if false {
                        var v13 int16 = 1000
                        fmt.Println(v12)
                        v14 := GenS1{F0: v13}
                        _ = v13
                        _ = v14
                    } else {
                        if (!(0.1 <= (1e10 / (1.5 / 0.1)))) {
                            var v15 int8 = -1
                            if false {
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                                fmt.Println(fi8, v12)
                                switch "foo bar" {
                                case "foo bar":
                                    if (true || (42 >= -1)) {
                                        v7 = append(v7, fi8)
                                    }
                                case "hello":
                                    switch "" {
                                    case "\n\t":
                                        v16 := GenS2{F0: 4294967295}
                                        _ = v16
                                    case "hello":
                                        _ = len(v7)
                                    default:
                                        switch -474 {
                                        case 128:
                                            for fi17 := 0; fi17 < 15; fi17++ {
                                                _ = len(v7)
                                                if ((!((2 < 42) || true)) && (0 <= 18446744073709551615)) {
                                                    v18 := GenS3{F0: true, F1: 156, F2: 128}
                                                    if (!((false || (!((v15 - (v15 ^ -31)) == v15))) || (-1 <= (v15 | v15)))) {
                                                        if false {
                                                            if len(v7) > 0 {
                                                                v7[0] = 128
                                                            }
                                                            var v19 uint8 = 1
                                                            _ = v19
                                                        } else {
                                                            if true {
                                                                var v20 uint64 = 42
                                                                if false {
                                                                    if false {
                                                                        v21 := GenS3{F0: (678 < 2147483648), F1: 1, F2: 199}
                                                                        _ = v21
                                                                    }
                                                                }
                                                                _ = v20
                                                            }
                                                            v22 := GenS4{F0: "\n\t"}
                                                            _ = v22
                                                        }
                                                        if (v9 > v9) {
                                                            var v23 string = ("a" + "foo bar")
                                                            var v24 uint8 = 128
                                                            _ = v23
                                                            _ = v24
                                                        }
                                                    }
                                                    _ = v18
                                                }
                                                v25 := GenS3{F0: false, F1: 42, F2: 255}
                                                _ = fi17
                                                _ = v25
                                            }
                                        case 1:
                                            for fi26 := 0; fi26 < 20; fi26++ {
                                                v27 := GenS5{F0: 2147483647}
                                                fmt.Println(v9, fi26)
                                                var v28 int16 = 42
                                                _ = fi26
                                                _ = v27
                                                _ = v28
                                            }
                                        default:
                                            v15 = v15
                                        }
                                    }
                                case "":
                                    switch "a" {
                                    case "hello":
                                        var v29 float32 = 0.1
                                        _ = v29
                                    case "":
                                        if len(v7) >= 2 {
                                            v7 = v7[1:len(v7)]
                                        }
                                    case "\n\t":
                                        if len(v7) >= 2 {
                                            v7 = v7[1:len(v7)]
                                        }
                                    }
                                default:
                                    if false {
                                        if true {
                                            if len(v7) >= 2 {
                                                v7 = v7[1:len(v7)]
                                            }
                                            for fi30 := 0; fi30 < 18; fi30++ {
                                                switch "foo bar" {
                                                case "a":
                                                    fmt.Println(v9)
                                                case "":
                                                    var v31 string = "hello"
                                                    _ = v31
                                                }
                                                var v32 int32 = 0
                                                _ = fi30
                                                _ = v32
                                            }
                                            for fi33 := 0; fi33 < 4; fi33++ {
                                                v12 += 0.1
                                                v15 -= v15
                                                fmt.Println(fi33, v15, fi8)
                                                _ = fi33
                                            }
                                        }
                                        var v34 int64 = v9
                                        _ = v34
                                    }
                                }
                            } else {
                                var v35 uint64 = 636
                                v7 = append(v7, (fi8 << ((1) & 63)))
                                _ = v35
                            }
                            _ = v15
                        }
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    }
                case -2147483648:
                    if false {
                        for fi36 := 0; fi36 < 12; fi36++ {
                            fmt.Println(fi8, v9)
                            v7 = append(v7, fi36)
                            _ = fi36
                        }
                        v9 -= (v9 << ((4294967295) & 63))
                        var v37 int8 = -113
                        _ = v37
                    } else {
                        switch (fi8 % ((fi8) | 1)) {
                        case -1:
                            fmt.Println(v9, fi8)
                        case 32767:
                            v9 += v9
                        default:
                            if len(v7) > 0 {
                                v7[0] = (fi8 | (fi8 & (fi8 % (((fi8 + fi8)) | 1))))
                            }
                        }
                        var v38 float64 = v12
                        var v39 uint16 = 581
                        _ = v38
                        _ = v39
                    }
                case 9223372036854775807:
                    v40 := GenS3{F0: (true || true), F1: -1, F2: 156}
                    _ = v40
                default:
                    var v41 uint16 = 1
                    _ = v41
                }
            case "a":
                var v42 int16 = -1
                _ = v42
            case "":
                var v43 string = ("\n\t" + (((("a" + ("\n\t" + "hello")) + "foo bar") + "a") + ("hello" + ("hello" + ""))))
                _ = v43
            default:
                if (!false) {
                    v44 := GenS1{F0: 0}
                    switch (("\n\t" + "") + ("foo bar" + ("foo bar" + "hello"))) {
                    case "":
                        fmt.Println(v12, v12)
                    }
                    _ = v44
                }
            }
            _ = v12
        } else {
            var v45 float32 = (((0.1 * (0.0 - (float32(((-2.25 / 1.0) - (0.0 + 1.5))) + (float32((0.1 / float64(1.5))) - -2.25)))) / (((0.1 * 0.1) - 1e10) - 1e10)) / ((float32((1e10 + -2.25)) + (1.5 * (1.5 + 0.0))) + (-2.25 + 1e10)))
            v45 += ((0.0 / (v45 - v45)) * -2.25)
            v46 := GenS4{F0: (("a" + "a") + "\n\t")}
            _ = v45
            _ = v46
        }
        for fi47 := 0; fi47 < 3; fi47++ {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            fmt.Println(fi8, fi8)
            _ = fi47
        }
        _ = fi8
        _ = v9
    }
    var v48 uint64 = 786
    switch "a" {
    case "hello":
        v48 += (^(v48))
    case "foo bar":
        fmt.Println(v48)
    case "a":
        v49 := GenS2{F0: v48}
        _ = v49
    default:
        v7 = append(v7, 32768)
    }
    _ = v48
    _ = v7
    return 240
}
func genFn1(p0 uint8, p1 string) int8 {
    v50 := []int{128, -9223372036854775808, 42, 437}
    for fi51 := 0; fi51 < 17; fi51++ {
        p1 = "foo bar"
        if true {
            var v52 uint16 = 32767
            if len(v50) > 0 {
                v50[0] = (fi51 + 128)
            }
            var v53 bool = true
            _ = v52
            _ = v53
        }
        _ = fi51
    }
    _ = p0
    _ = p1
    _ = v50
    return -45
}

func main() {
    v1 := []string{"a"}
    var v2 uint = 781
    var v3 uint8 = 10
    var v4 uint = v2
    var v5 int8 = 127
    var v54 float32 = ((-2.25 * 0.1) / -2.25)
    var v55 bool = (-564 > 0)
    for fi56 := 0; fi56 < 12; fi56++ {
        if v55 {
            v3 = 127
            if len(v1) > 0 {
                v1[0] = "hello"
            }
        } else {
            v55 = v55
        }
        for fi57 := 0; fi57 < 8; fi57++ {
            switch (fi56 / ((fi57) | 1)) {
            case -2147483648:
                v5 -= genFn1(v3, "a")
            case 0:
                if (42 < -351) {
                    if v55 {
                        var v58 string = ("foo bar" + ("" + (((("hello" + "") + "hello") + "foo bar") + ("hello" + "hello"))))
                        _ = v58
                    } else {
                        var v59 uint16 = 32767
                        fmt.Println(v59)
                        _ = v59
                    }
                    v55 = true
                    v2 = v4
                }
            }
            var v60 int8 = genFn1((v3 / ((50) | 1)), "hello")
            v5 -= (v60 * genFn1((v3 - v3), "a"))
            _ = fi57
            _ = v60
        }
        var v61 float32 = v54
        _ = fi56
        _ = v61
    }
    var v62 bool = (32767 != 408)
    if (!v62) {
        if v62 {
            v63 := GenS5{F0: 1}
            _ = v63
        } else {
            v3 = (^(v3))
            var v64 uint16 = 65535
            _ = v64
        }
    } else {
        v65 := GenS6{F0: 1}
        _ = v65
    }
    if len(v1) > 0 {
        v1[0] = ""
    }
    var v66 float64 = -2.25
    for fi67 := 0; fi67 < 16; fi67++ {
        var v68 uint64 = 661
        _ = fi67
        _ = v68
    }
    var v69 float64 = (0.1 / v66)
    switch ("hello" + "\n\t") {
    case "hello":
        if v55 {
            var v70 int32 = 42
            _ = v70
        }
    case "a":
        v71 := GenS7{F0: (v2 / ((v4) | 1)), F1: 4294967295, F2: (v2 >> (((v2 * v2)) & 63))}
        _ = v71
    case "":
        var v72 uint64 = 921
        _ = v72
    default:
        if v55 {
            switch "hello" {
            case "\n\t":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "a":
                for fi73 := 0; fi73 < 12; fi73++ {
                    switch "foo bar" {
                    case "foo bar":
                        v66 -= (0.0 - 1e10)
                    case "\n\t":
                        if (uint16(v3) == 519) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v2 = (v4 >> (((v2 % (((v4 << ((v4) & 63))) | 1))) & 63))
                        } else {
                            if (18446744073709551615 <= v4) {
                                v74 := GenS8{F0: -32768, F1: 1000}
                                _ = v74
                            }
                            v54 = (v54 * v54)
                            v75 := GenS9{F0: v4}
                            _ = v75
                        }
                    case "":
                        var v76 uint16 = 32767
                        _ = v76
                    default:
                        var v77 int16 = 296
                        _ = v77
                    }
                    if ("" <= "hello") {
                        switch (fi73 ^ fi73) {
                        case 818:
                            v3 += 0
                        default:
                            v78 := GenS7{F0: 18446744073709551615, F1: 4294967295, F2: (v2 << (((v4 / ((255) | 1))) & 63))}
                            _ = v78
                        }
                        v66 -= (v69 * v66)
                    }
                    _ = fi73
                }
            case "":
                v66 = v66
            default:
                switch 1 {
                case 4294967295:
                    var v79 uint8 = (v3 ^ v3)
                    _ = v79
                }
            }
        } else {
            if (123 <= 1000) {
                v4 = (v2 % ((v4) | 1))
                fmt.Println(v4, v69)
                for fi80 := 0; fi80 < 11; fi80++ {
                    _ = len(v1)
                    v81 := GenS10{F0: 32767, F1: v5, F2: (-95 != -652)}
                    for fi82 := 0; fi82 < 6; fi82++ {
                        _ = len(v1)
                        switch "hello" {
                        case "hello":
                            if len(v1) > 0 {
                                v1[0] = ("hello" + "a")
                            }
                        case "":
                            if (-9223372036854775808 < 1) {
                                switch ("a" + "\n\t") {
                                case "foo bar":
                                    v83 := GenS5{F0: 42}
                                    _ = v83
                                case "hello":
                                    switch "hello" {
                                    case "":
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    case "a":
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    default:
                                        var v84 int8 = v5
                                        _ = v84
                                    }
                                default:
                                    v5 -= v5
                                }
                                var v85 uint16 = 1
                                var v86 float64 = -2.25
                                _ = v85
                                _ = v86
                            } else {
                                if true {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                } else {
                                    v87 := GenS11{F0: (-(fi82)), F1: 2147483647}
                                    _ = v87
                                }
                                if true {
                                    v54 = 1.5
                                    if true {
                                        v88 := GenS12{F0: 32767, F1: (fi80 & -282)}
                                        _ = v88
                                    } else {
                                        v3 = (v3 - uint8(v2))
                                        if (!v55) {
                                            v89 := GenS5{F0: 678}
                                            v1 = append(v1, "hello")
                                            var v90 bool = v62
                                            _ = v89
                                            _ = v90
                                        } else {
                                            if len(v1) > 0 {
                                                v1[0] = "\n\t"
                                            }
                                            switch "\n\t" {
                                            case "":
                                                var v91 float32 = v54
                                                _ = v91
                                            case "hello":
                                                var v92 int64 = 42
                                                _ = v92
                                            default:
                                                v54 = (((float32((float64(((float32(float64(0.1)) / (v54 * ((v54 - v54) - (v54 * -2.25)))) - v54)) + v69)) / (1e10 - v54)) * v54) - 0.1)
                                                var v93 uint32 = 832
                                                v93 -= 302
                                                _ = v93
                                            }
                                            v69 -= (float64(float32(-2.25)) * v69)
                                            if len(v1) > 0 {
                                                v1[0] = (("foo bar" + "a") + "foo bar")
                                            }
                                            switch fi82 {
                                            case 9223372036854775807:
                                                var v94 int32 = int32(v2)
                                                _ = v94
                                            case -741:
                                                if v62 {
                                                    if true {
                                                        var v95 int64 = 0
                                                        _ = v95
                                                    }
                                                    if ((float32(0.1) - (v54 + (v54 / v54))) >= ((3.14159 - v54) / v54)) {
                                                        v1 = append(v1, "\n\t")
                                                        var v96 float32 = 0.0
                                                        var v97 string = (("a" + "\n\t") + "hello")
                                                        _ = v96
                                                        _ = v97
                                                    }
                                                    var v98 uint64 = 2
                                                    var v99 int32 = -2147483648
                                                    _ = v98
                                                    _ = v99
                                                }
                                            case 32767:
                                                if (263 >= 868) {
                                                    switch 1 {
                                                    case 4294967295:
                                                        switch fi80 {
                                                        case -1:
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        }
                                                    case 9223372036854775807:
                                                        v100 := GenS13{F0: -32768, F1: v5, F2: 922}
                                                        _ = v100
                                                    case -491:
                                                        switch (fi80 ^ (^(fi82))) {
                                                        case 32768:
                                                            fmt.Println(fi82, v4, v4)
                                                        case 128:
                                                            if ("foo bar" > "hello") {
                                                                fmt.Println(v62, v54, v62)
                                                                var v101 int = -688
                                                                v102 := GenS14{F0: 1}
                                                                _ = v101
                                                                _ = v102
                                                            } else {
                                                                var v103 int = int(v4)
                                                                _ = v103
                                                            }
                                                        case 490:
                                                            v104 := GenS8{F0: 0, F1: -924}
                                                            _ = v104
                                                        default:
                                                            if v55 {
                                                                if v62 {
                                                                    if len(v1) > 0 {
                                                                        v1[0] = ""
                                                                    }
                                                                    if ((^(v2)) <= v2) {
                                                                        var v105 uint32 = 391
                                                                        fmt.Println(v54, v66, v55)
                                                                        _ = v105
                                                                    } else {
                                                                        var v106 uint16 = 2
                                                                        _ = v106
                                                                    }
                                                                    if (v4 <= v2) {
                                                                        if v62 {
                                                                            var v107 int16 = 704
                                                                            var v108 int = (fi80 << (((^(v4))) & 63))
                                                                            _ = v107
                                                                            _ = v108
                                                                        }
                                                                    }
                                                                } else {
                                                                    v2 = v4
                                                                    if v62 {
                                                                        var v109 uint = v2
                                                                        if v62 {
                                                                            var v110 uint32 = 0
                                                                            _ = v110
                                                                        } else {
                                                                            fmt.Println(v62, v2, v4)
                                                                        }
                                                                        _ = v109
                                                                    }
                                                                }
                                                                v5 = genFn1((v3 % ((v3) | 1)), "")
                                                                v62 = v62
                                                                v111 := GenS7{F0: 0, F1: 42, F2: (^(v2))}
                                                                _ = v111
                                                            }
                                                            _ = len(v1)
                                                        }
                                                    default:
                                                        var v112 bool = (-9223372036854775808 <= 0)
                                                        _ = v112
                                                    }
                                                    var v113 uint16 = uint16(v3)
                                                    _ = v113
                                                } else {
                                                    v1 = append(v1, ("\n\t" + "hello"))
                                                }
                                                fmt.Println(v5, fi82, v62)
                                            }
                                            var v114 uint8 = uint8(v2)
                                            _ = v114
                                        }
                                        v5 = (v5 % ((genFn1(128, "\n\t")) | 1))
                                    }
                                }
                                var v115 string = "\n\t"
                                if len(v1) > 0 {
                                    v1[0] = v115
                                }
                                var v116 uint16 = 906
                                var v117 int = 1
                                switch ("hello" + "") {
                                case "foo bar":
                                    v115 = v115
                                    if (genFn1(v3, "a") == 53) {
                                        if (127 <= v5) {
                                            if len(v1) > 0 {
                                                v1[0] = "hello"
                                            }
                                            fmt.Println(fi82)
                                        }
                                        if v55 {
                                            if ((fi80 >> ((18446744073709551615) & 63)) > 32768) {
                                                switch ("foo bar" + "\n\t") {
                                                case "":
                                                    if false {
                                                        _ = len(v1)
                                                    }
                                                case "\n\t":
                                                    _ = len(v1)
                                                case "a":
                                                    v118 := GenS9{F0: (v4 + (v4 << ((v2) & 63)))}
                                                    _ = v118
                                                }
                                                v117 += (fi82 % ((v117) | 1))
                                                if true {
                                                    if v55 {
                                                        var v119 uint8 = 228
                                                        var v120 uint32 = 42
                                                        var v121 string = ((v115 + v115) + v115)
                                                        var v122 float32 = v54
                                                        v123 := GenS3{F0: (42 != 18446744073709551615), F1: 1, F2: v119}
                                                        _ = v119
                                                        _ = v120
                                                        _ = v121
                                                        _ = v122
                                                        _ = v123
                                                    }
                                                    fmt.Println(v66, v66, fi80)
                                                    var v124 uint = 695
                                                    _ = v124
                                                } else {
                                                    if (1 <= 1000) {
                                                        v125 := GenS9{F0: uint(v116)}
                                                        _ = v125
                                                    }
                                                }
                                            }
                                            var v126 uint = (v2 / ((65535) | 1))
                                            _ = v126
                                        } else {
                                            switch ((v115 + v115) + ("hello" + "a")) {
                                            case "hello":
                                                if true {
                                                    switch v115 {
                                                    case "a":
                                                        if len(v1) > 0 {
                                                            v1[0] = v115
                                                        }
                                                    default:
                                                        if len(v1) > 0 {
                                                            v1[0] = v115
                                                        }
                                                    }
                                                    v2 = uint(v116)
                                                } else {
                                                    var v127 uint64 = 0
                                                    var v128 int64 = 42
                                                    _ = v127
                                                    _ = v128
                                                }
                                            case "\n\t":
                                                v117 = fi82
                                            case "a":
                                                if (v54 >= (((3.14159 + v54) / ((float32(v66) / 1.5) * v54)) / 0.1)) {
                                                    v5 -= 0
                                                    var v129 uint32 = 1
                                                    _ = v129
                                                } else {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    var v130 bool = v62
                                                    fmt.Println(v2, v4)
                                                    _ = v130
                                                }
                                            default:
                                                var v131 string = "foo bar"
                                                _ = v131
                                            }
                                        }
                                        fmt.Println(v66, v55)
                                    } else {
                                        fmt.Println(fi82, v66)
                                        if (((^(v2)) == (v4 | (^(v4)))) && (v54 > float32(float64(v54)))) {
                                            v116 = (v116 * (v116 | (v116 | (v116 | v116))))
                                            var v132 int64 = 229
                                            if (!v62) {
                                                v69 -= v69
                                                v133 := GenS7{F0: uint(v3), F1: 42, F2: v2}
                                                _ = v133
                                            } else {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v4 = v4
                                                var v134 uint16 = (v116 & v116)
                                                var v135 bool = v55
                                                _ = v134
                                                _ = v135
                                            }
                                            _ = v132
                                        } else {
                                            var v136 float64 = v66
                                            v136 = v136
                                            _ = v136
                                        }
                                        var v137 int = 9223372036854775807
                                        _ = v137
                                    }
                                case "\n\t":
                                    if v55 {
                                        v1 = append(v1, "\n\t")
                                    } else {
                                        v2 = v4
                                        var v138 float32 = float32(v69)
                                        _ = v138
                                    }
                                }
                                _ = v115
                                _ = v116
                                _ = v117
                            }
                        case "foo bar":
                            var v139 int64 = 0
                            _ = v139
                        default:
                            var v140 uint16 = 65535
                            _ = v140
                        }
                        v54 += (((v54 - v54) * v54) - float32((3.14159 - float64(v54))))
                        _ = fi82
                    }
                    _ = fi80
                    _ = v81
                }
            } else {
                for fi141 := 0; fi141 < 14; fi141++ {
                    if v55 {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        fmt.Println(fi141, v5, v4)
                    } else {
                        v1 = append(v1, "hello")
                    }
                    if true {
                        var v142 int16 = -32768
                        _ = v142
                    }
                    _ = fi141
                }
                var v143 uint = 255
                for fi144 := 0; fi144 < 20; fi144++ {
                    var v145 string = "\n\t"
                    for fi146 := 0; fi146 < 5; fi146++ {
                        switch "hello" {
                        case "hello":
                            v147 := GenS11{F0: (fi144 ^ (^(fi144))), F1: -76}
                            _ = v147
                        case "\n\t":
                            if (42 <= 18446744073709551615) {
                                if v62 {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                var v148 uint32 = uint32(v2)
                                switch fi146 {
                                case 32768:
                                    v149 := GenS13{F0: -605, F1: 127, F2: 1000}
                                    _ = v149
                                }
                                _ = v148
                            } else {
                                v150 := GenS15{F0: v4, F1: -1}
                                v62 = (676 >= 42)
                                if (-664 < -472) {
                                    if (3.14159 == -2.25) {
                                        v151 := GenS16{F0: 2}
                                        _ = v151
                                    } else {
                                        if true {
                                            switch 908 {
                                            case -88:
                                                fmt.Println(fi146, v66, v62)
                                            case 2147483647:
                                                switch v145 {
                                                case "hello":
                                                    var v152 int32 = 42
                                                    _ = v152
                                                case "foo bar":
                                                    _ = len(v1)
                                                default:
                                                    fmt.Println(v54, v66)
                                                }
                                            case 4294967295:
                                                var v153 uint8 = 127
                                                _ = v153
                                            default:
                                                v2 = 65535
                                            }
                                            var v154 uint32 = 65535
                                            _ = v154
                                        } else {
                                            switch (fi144 % ((fi144) | 1)) {
                                            case 493:
                                                var v155 uint32 = 0
                                                _ = v155
                                            case 4294967295:
                                                v156 := GenS14{F0: 267}
                                                _ = v156
                                            case 127:
                                                var v157 uint = 42
                                                _ = v157
                                            default:
                                                v158 := GenS16{F0: (v4 ^ v2)}
                                                _ = v158
                                            }
                                        }
                                    }
                                    v69 = float64(3.14159)
                                    if (true || (((v5 + (v5 >> (((v2 + v2)) & 63))) < -128) && v55)) {
                                        var v159 int64 = -1
                                        if len(v1) > 0 {
                                            v1[0] = ""
                                        }
                                        _ = v159
                                    } else {
                                        var v160 uint = (v4 | 42)
                                        v5 = (v5 * (v5 | genFn1(uint8(v160), v145)))
                                        var v161 int64 = -1
                                        v162 := GenS9{F0: (v143 / ((v2) | 1))}
                                        v161 -= (v161 + (v161 - (v161 | v161)))
                                        _ = v160
                                        _ = v161
                                        _ = v162
                                    }
                                    if v62 {
                                        var v163 bool = true
                                        v164 := GenS17{F0: 0, F1: ((v145 + "foo bar") + "a")}
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v163
                                        _ = v164
                                    } else {
                                        v165 := GenS18{F0: 42, F1: 42, F2: 42}
                                        var v166 int64 = -9223372036854775808
                                        v167 := GenS17{F0: (v166 | v166), F1: v145}
                                        var v168 string = v145
                                        var v169 float64 = float64((((v54 - v54) + 1.5) * v54))
                                        _ = v165
                                        _ = v166
                                        _ = v167
                                        _ = v168
                                        _ = v169
                                    }
                                    v170 := GenS12{F0: 0, F1: (fi144 / (((^(fi144))) | 1))}
                                    _ = v170
                                } else {
                                    fmt.Println(v62)
                                    v145 = "foo bar"
                                    var v171 uint32 = 2
                                    _ = v171
                                }
                                _ = v150
                            }
                        default:
                            if true {
                                var v172 int = 32768
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v172
                            }
                        }
                        _ = fi146
                    }
                    _ = fi144
                    _ = v145
                }
                _ = v143
            }
            v1 = append(v1, "hello")
        }
    }
    switch ("foo bar" + "hello") {
    case "\n\t":
        switch 32768 {
        case -2147483648:
            fmt.Println(v3, v69, v66)
        }
    default:
        v173 := GenS12{F0: -941, F1: -9223372036854775808}
        _ = v173
    }
    v3 += (v3 << (((v4 + v4)) & 63))
    var v174 bool = (0 > 2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v174:", v174)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v62:", v62)
    fmt.Println("v66:", v66)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v174
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v62
    _ = v66
    _ = v69
}
