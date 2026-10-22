package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 string; F1 bool; F2 int8 }
type GenS5 struct { F0 uint8; F1 int8; F2 string }
type GenS6 struct { F0 uint16; F1 int8 }
type GenS7 struct { F0 uint8; F1 uint16 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int16; F1 int; F2 int }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int8; F1 int16 }
type GenS12 struct { F0 uint; F1 uint }
type GenS13 struct { F0 uint16; F1 string; F2 int16 }
type GenS14 struct { F0 uint16; F1 int }
type GenS15 struct { F0 int64; F1 int64 }
type GenS16 struct { F0 uint32; F1 int64 }
type GenS17 struct { F0 int; F1 int8; F2 int32 }
type GenS18 struct { F0 uint32; F1 bool }

func main() {
    v1 := make(map[int]string)
    var v2 int8 = 42
    var v3 uint = 999
    v4 := []uint32{471, 2147483647, 2, 4294967295}
    v5 := make(map[uint32]bool)
    fmt.Println(v3, v2)
    v2 -= v2
    _ = len(v4)
    v3 -= v3
    fmt.Println(v3, v3)
    if true {
        fmt.Println(v2)
    } else {
        if ("a" != (("" + ("\n\t" + "hello")) + "")) {
            v3 -= v3
            fmt.Println(v3, v3)
            v2 -= 100
        } else {
            if (0.0 < 1e10) {
                switch "foo bar" {
                case "foo bar":
                    v3 += v3
                case "\n\t":
                    v2 = v2
                default:
                    fmt.Println(v2)
                }
                v3 -= (v3 ^ (v3 & (v3 / ((v3) | 1))))
                if (false && (490 <= 2)) {
                    for fi6 := 0; fi6 < 17; fi6++ {
                        v7 := GenS1{F0: (("hello" + "hello") + ""), F1: v3}
                        fmt.Println(v3, v3)
                        _ = fi6
                        _ = v7
                    }
                    var v8 float32 = -2.25
                    _ = v8
                } else {
                    v3 = v3
                    if false {
                        v2 = (v2 / (((v2 % (((v2 - v2)) | 1))) | 1))
                        if false {
                            var v9 int = 896
                            _ = v9
                        } else {
                            v10 := GenS1{F0: "\n\t", F1: 497}
                            v2 += (v2 + v2)
                            var v11 int = -902
                            _ = v10
                            _ = v11
                        }
                    }
                    fmt.Println(v2, v2)
                }
            } else {
                for fi12 := 0; fi12 < 16; fi12++ {
                    switch "\n\t" {
                    case "":
                        for fi13 := 0; fi13 < 18; fi13++ {
                            if (0.0 > (((-2.25 - ((float64(float32(-2.25)) + ((-2.25 * 0.1) + 1e10)) + (3.14159 * (-2.25 / -2.25)))) - 1e10) - 1e10)) {
                                v2 = v2
                                switch fi12 {
                                case 0:
                                    switch "a" {
                                    case "foo bar":
                                        if true {
                                            var v14 uint = 255
                                            v2 += (v2 >> ((v14) & 63))
                                            v2 = (v2 / ((v2) | 1))
                                            _ = v14
                                        }
                                    }
                                default:
                                    _ = len(v1)
                                }
                            } else {
                                var v15 float32 = 3.14159
                                if (true && (!true)) {
                                    fmt.Println(fi13, fi12)
                                    if (882 == 136) {
                                        delete(v5, 42)
                                        fmt.Println(fi12, fi12, fi13)
                                        switch (("\n\t" + "\n\t") + "\n\t") {
                                        case "":
                                            v16 := GenS1{F0: "\n\t", F1: v3}
                                            if false {
                                                v17 := GenS2{F0: (v2 >> ((255) & 63))}
                                                if (-902 != -472) {
                                                    if false {
                                                        fmt.Println(fi12, v3, fi13)
                                                        var v18 uint8 = 128
                                                        _ = v18
                                                    } else {
                                                        v19 := GenS2{F0: (v2 >> ((v3) & 63))}
                                                        _ = v19
                                                    }
                                                    if true {
                                                        v3 = (v3 / (((^(v3))) | 1))
                                                        v3 += v3
                                                        fmt.Println(fi12)
                                                    } else {
                                                        var v20 uint8 = 24
                                                        fmt.Println(v20)
                                                        _ = v20
                                                    }
                                                    var v21 bool = (int16(fi12) >= 42)
                                                    _ = v21
                                                } else {
                                                    var v22 uint64 = 1
                                                    _ = v22
                                                }
                                                _ = v17
                                            } else {
                                                var v23 bool = false
                                                v2 -= (v2 - v2)
                                                _ = v23
                                            }
                                            if true {
                                                if true {
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        switch "" {
                                                        case "hello":
                                                            fmt.Println(fi12)
                                                            if (-2.25 >= (-2.25 / (float64((v15 + v15)) + (float64(v15) / 1.5)))) {
                                                                delete(v1, (-(fi13)))
                                                            }
                                                        case "\n\t":
                                                            v3 = 691
                                                        default:
                                                            if (42 < 9223372036854775807) {
                                                                if true {
                                                                    v2 = (v2 << ((245) & 63))
                                                                } else {
                                                                    if len(v4) > 0 {
                                                                        v4[0] = 2147483648
                                                                    }
                                                                    v15 += ((float32(-2.25) - v15) + (-2.25 + v15))
                                                                }
                                                                var v24 float32 = 0.1
                                                                if false {
                                                                    v15 += (v24 - (v15 + 0.0))
                                                                    var v25 uint16 = 65535
                                                                    _ = v25
                                                                }
                                                                v26 := GenS3{F0: (v3 ^ (v3 | (v3 ^ (v3 ^ v3)))), F1: -2147483648}
                                                                _ = v24
                                                                _ = v26
                                                            }
                                                        }
                                                    case "a":
                                                        var v27 float64 = (-2.25 / 1.0)
                                                        _ = v27
                                                    default:
                                                        var v28 uint32 = 682
                                                        _ = v28
                                                    }
                                                }
                                            } else {
                                                switch "" {
                                                case "":
                                                    v3 += 42
                                                default:
                                                    if (true || true) {
                                                        switch fi13 {
                                                        case 4294967295:
                                                            fmt.Println(v2)
                                                        }
                                                        if (136 <= 0) {
                                                            fmt.Println(v3, v15)
                                                        } else {
                                                            v29 := GenS2{F0: -54}
                                                            v3 = v3
                                                            if (((((3.14159 / -2.25) / 1.5) - 0.0) < (3.14159 + (1.5 * (3.14159 / (3.14159 + (3.14159 / 0.1)))))) && true) {
                                                                if (71 > 1) {
                                                                    var v30 uint32 = 2
                                                                    _ = v30
                                                                }
                                                                v31 := GenS4{F0: "hello", F1: false, F2: 100}
                                                                _ = v31
                                                            } else {
                                                                var v32 int = fi12
                                                                fmt.Println(v3, v2)
                                                                v33 := GenS5{F0: 107, F1: v2, F2: "hello"}
                                                                v34 := GenS6{F0: 42, F1: v2}
                                                                v3 -= (v3 & (v3 / ((0) | 1)))
                                                                _ = v32
                                                                _ = v33
                                                                _ = v34
                                                            }
                                                            _ = v29
                                                        }
                                                    } else {
                                                        v35 := GenS1{F0: ("\n\t" + "a"), F1: (v3 * v3)}
                                                        _ = v35
                                                    }
                                                }
                                            }
                                            _ = v16
                                        case "\n\t":
                                            v15 = (v15 - (0.1 * (v15 * v15)))
                                        default:
                                            v36 := GenS7{F0: 1, F1: 32767}
                                            _ = v36
                                        }
                                    }
                                } else {
                                    v3 -= v3
                                    v37 := GenS6{F0: 790, F1: (v2 % ((v2) | 1))}
                                    if ("\n\t" == "") {
                                        if ((!(!true)) && false) {
                                            _ = len(v1)
                                            v3 -= 4294967295
                                        }
                                        switch (fi13 >> ((255) & 63)) {
                                        case 2147483647:
                                            v38 := GenS5{F0: 42, F1: v2, F2: ((("\n\t" + "") + "foo bar") + ("\n\t" + "hello"))}
                                            _ = v38
                                        case 42:
                                            fmt.Println(fi13, fi12)
                                        case 477:
                                            v39 := GenS5{F0: 42, F1: (-(v2)), F2: "foo bar"}
                                            _ = v39
                                        default:
                                            if true {
                                                v40 := v1[(^(fi12))]
                                                switch (fi13 | (fi13 << (((v3 % ((65535) | 1))) & 63))) {
                                                case 32768:
                                                    if (!true) {
                                                        v15 = v15
                                                    } else {
                                                        if (v40 <= (v40 + v40)) {
                                                            fmt.Println(v15, v40, v2)
                                                            v41 := GenS8{F0: 18446744073709551615}
                                                            var v42 int16 = -1
                                                            _ = v41
                                                            _ = v42
                                                        } else {
                                                            fmt.Println(fi12, v3)
                                                        }
                                                        v4 = append(v4, 1)
                                                    }
                                                case 127:
                                                    v2 = (v2 ^ -53)
                                                default:
                                                    if true {
                                                        v3 = (v3 & v3)
                                                        fmt.Println(fi12, fi13)
                                                        v15 = (0.1 * ((-2.25 / v15) - float32(1e10)))
                                                    }
                                                }
                                                switch "" {
                                                case "a":
                                                    var v43 string = v40
                                                    _ = v43
                                                }
                                                _ = v40
                                            }
                                        }
                                    } else {
                                        v15 -= (0.1 / ((v15 / v15) - float32(1.5)))
                                        v15 = 1e10
                                    }
                                    _ = v37
                                }
                                _ = v15
                            }
                            _ = fi13
                        }
                    case "\n\t":
                        v4 = append(v4, 0)
                    default:
                        var v44 uint64 = 42
                        _ = v44
                    }
                    v45 := GenS9{F0: 1000, F1: fi12, F2: (-(fi12))}
                    _ = fi12
                    _ = v45
                }
            }
            var v46 int32 = int32(v3)
            if (!(270 < 159)) {
                if (225 == 42) {
                    v2 = (v2 + (v2 | (v2 ^ v2)))
                    switch ("foo bar" + "") {
                    case "a":
                        v2 = -1
                    case "hello":
                        v3 = (v3 / (((v3 + 2)) | 1))
                    case "":
                        v46 -= (^(v46))
                    default:
                        _ = len(v4)
                    }
                    var v47 int32 = (v46 * (v46 + v46))
                    _ = v47
                }
                var v48 uint = 42
                _ = v48
            } else {
                v49 := GenS7{F0: 42, F1: 65535}
                if false {
                    fmt.Println(v3, v3)
                    var v50 uint16 = 757
                    _ = v50
                } else {
                    v3 += (v3 - (v3 | v3))
                    v46 -= (v46 << (((^(v3))) & 63))
                }
                _ = v49
            }
            _ = v46
        }
        v4 = append(v4, 2147483647)
    }
    var v51 uint = (v3 + (v3 ^ v3))
    fmt.Println(v2, v51)
    if (false || (3.14159 > (((1e10 / 0.1) - ((3.14159 / (0.0 - (float32((1.5 * 1e10)) * 0.1))) + 1e10)) + 1.5))) {
        v2 = (v2 >> ((v51) & 63))
        var v52 int8 = (v2 >> ((v3) & 63))
        var v53 uint16 = 278
        _ = v52
        _ = v53
    }
    v54 := GenS1{F0: ("foo bar" + "hello"), F1: v51}
    if len(v4) > 0 {
        v4[0] = 641
    }
    v51 -= v3
    if false {
        for fi55 := 0; fi55 < 19; fi55++ {
            var v56 uint32 = uint32(v3)
            if (!false) {
                if (367 >= -1) {
                    v3 -= v3
                    for fi57 := 0; fi57 < 3; fi57++ {
                        v58 := GenS3{F0: (v3 * (^(v51))), F1: -1}
                        _ = fi57
                        _ = v58
                    }
                } else {
                    if ((v56 % ((v56) | 1)) != 2147483647) {
                        switch (fi55 - (fi55 * int(v56))) {
                        case -335:
                            var v59 uint32 = v56
                            _ = v59
                        }
                        var v60 uint32 = v56
                        _ = v60
                    }
                    for fi61 := 0; fi61 < 7; fi61++ {
                        if true {
                            v62 := GenS1{F0: "foo bar", F1: (v51 - v3)}
                            v63 := GenS2{F0: v2}
                            v64 := GenS4{F0: "a", F1: true, F2: v2}
                            if false {
                                var v65 string = ""
                                _ = v65
                            } else {
                                var v66 uint32 = v56
                                _ = v66
                            }
                            _ = v62
                            _ = v63
                            _ = v64
                        } else {
                            var v67 int8 = 127
                            _ = v67
                        }
                        fmt.Println(fi61)
                        v68 := GenS4{F0: "hello", F1: (((!true) && (65535 <= 525)) || ((!false) && ((v51 != (v51 >> (((v51 + (v51 & v3))) & 63))) || (true || true)))), F2: (v2 & -128)}
                        switch (fi61 - (fi55 + -40)) {
                        case -1:
                            if ((!(0 != 527)) || ("hello" <= ("a" + "hello"))) {
                                v2 = (v2 * 105)
                                v2 += v2
                                if false {
                                    v51 = (v51 ^ 18446744073709551615)
                                } else {
                                    fmt.Println(v56)
                                }
                            } else {
                                var v69 uint16 = 739
                                var v70 uint = v51
                                _ = v69
                                _ = v70
                            }
                        default:
                            v56 = 2
                        }
                        if true {
                            switch (^(fi61)) {
                            case 128:
                                if false {
                                    var v71 uint64 = 4294967295
                                    if false {
                                        v3 += 581
                                        fmt.Println(fi55)
                                        v3 = (v3 & (v51 | v3))
                                        v72 := v5[v56]
                                        _ = v72
                                    }
                                    _ = v71
                                } else {
                                    var v73 int = 128
                                    if true {
                                        var v74 int8 = v2
                                        fmt.Println(v73, fi61, v51)
                                        _ = v74
                                    }
                                    if (42 >= 323) {
                                        var v75 uint8 = 42
                                        switch ("hello" + "") {
                                        case "a":
                                            v75 += v75
                                        }
                                        _ = v75
                                    } else {
                                        v56 -= (v56 ^ v56)
                                        fmt.Println(fi61, v51, v56)
                                        if false {
                                            v76 := GenS10{F0: ""}
                                            switch ("hello" + ((("foo bar" + ((("a" + ("foo bar" + "\n\t")) + "") + "\n\t")) + ("hello" + "a")) + "a")) {
                                            case "":
                                                var v77 int32 = 42
                                                _ = v77
                                            case "a":
                                                var v78 int = (v73 << (((v3 + 42)) & 63))
                                                _ = v78
                                            }
                                            var v79 float64 = 0.1
                                            _ = v76
                                            _ = v79
                                        }
                                    }
                                    _ = v73
                                }
                            default:
                                var v80 uint32 = (v56 >> ((v3) & 63))
                                _ = v80
                            }
                            switch "foo bar" {
                            case "\n\t":
                                if true {
                                    v81 := GenS9{F0: -32768, F1: (fi61 & fi61), F2: (fi61 + fi55)}
                                    if (((true || false) && (false && true)) || false) {
                                        v56 += v56
                                        v56 -= 4294967295
                                        v82 := GenS3{F0: (v3 + 42), F1: 0}
                                        switch ("\n\t" + "a") {
                                        case "foo bar":
                                            v83 := GenS7{F0: 0, F1: 65535}
                                            _ = v83
                                        case "hello":
                                            switch 1 {
                                            case 1:
                                                fmt.Println(v3, v56)
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                        _ = v82
                                    }
                                    _ = v81
                                }
                            case "hello":
                                switch (fi61 ^ fi55) {
                                case 660:
                                    v4 = append(v4, v56)
                                case 602:
                                    if ((v51 & 4294967295) >= uint(v56)) {
                                        v56 = 4294967295
                                        if false {
                                            var v84 uint = v51
                                            _ = v84
                                        }
                                    } else {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        fmt.Println(v3, fi55)
                                    }
                                default:
                                    var v85 bool = (!false)
                                    _ = v85
                                }
                            case "foo bar":
                                v2 = v2
                            }
                            v51 = v51
                            v56 = (v56 * (v56 - 2147483647))
                            if (!(false || (!false))) {
                                v3 -= 255
                                switch (("a" + "") + "a") {
                                case "a":
                                    var v86 string = "\n\t"
                                    _ = v86
                                default:
                                    v87 := GenS11{F0: 42, F1: -892}
                                    v56 = v56
                                    var v88 int16 = -752
                                    _ = v87
                                    _ = v88
                                }
                            }
                            v4 = append(v4, (v56 / ((v56) | 1)))
                            v89 := v1[42]
                            _ = v89
                        } else {
                            v2 += v2
                        }
                        _ = fi61
                        _ = v68
                    }
                }
                fmt.Println(fi55, fi55, fi55)
                var v90 bool = true
                _ = v90
            } else {
                v56 = (^(v56))
            }
            _ = fi55
            _ = v56
        }
    } else {
        if (128 >= 42) {
            for fi91 := 0; fi91 < 15; fi91++ {
                switch (fi91 >> (((v51 * (v51 ^ v3))) & 63)) {
                case -397:
                    v92 := GenS12{F0: (v51 + (v51 >> ((v3) & 63))), F1: (v3 << ((v3) & 63))}
                    _ = v92
                case -463:
                    if (-2.25 != 0.0) {
                        v2 = (v2 % ((v2) | 1))
                        var v93 uint = (^(v51))
                        for fi94 := 0; fi94 < 20; fi94++ {
                            v95 := v1[32767]
                            _ = fi94
                            _ = v95
                        }
                        _ = v93
                    }
                case 0:
                    fmt.Println(fi91, v3)
                }
                v96 := GenS8{F0: 898}
                if (!false) {
                    v3 += v51
                } else {
                    for fi97 := 0; fi97 < 8; fi97++ {
                        v2 = (v2 + (v2 - 1))
                        fmt.Println(fi91)
                        fmt.Println(v3, fi97, fi91)
                        _ = fi97
                    }
                    var v98 int32 = -1
                    if (65535 > 4294967295) {
                        switch ("foo bar" + "") {
                        case "hello":
                            var v99 int64 = 282
                            _ = v99
                        case "":
                            if (71 != 47) {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                fmt.Println(v98, v51)
                                if ((^(v3)) > (v51 << ((579) & 63))) {
                                    var v100 uint16 = uint16(v3)
                                    for fi101 := 0; fi101 < 10; fi101++ {
                                        v102 := GenS13{F0: v100, F1: "foo bar", F2: -394}
                                        v103 := GenS14{F0: (v100 % ((65535) | 1)), F1: fi91}
                                        _ = fi101
                                        _ = v102
                                        _ = v103
                                    }
                                    v104 := GenS15{F0: 42, F1: -9223372036854775808}
                                    _ = v100
                                    _ = v104
                                } else {
                                    v98 += v98
                                    switch "hello" {
                                    case "a":
                                        if (true || true) {
                                            v105 := v1[32768]
                                            fmt.Println(v105)
                                            if true {
                                                v106 := v1[(fi91 >> ((v51) & 63))]
                                                delete(v5, 720)
                                                v107 := GenS16{F0: 65535, F1: 42}
                                                _ = v106
                                                _ = v107
                                            }
                                            _ = v105
                                        }
                                    case "":
                                        if (false && true) {
                                            v2 += -128
                                            var v108 uint = v51
                                            _ = v108
                                        }
                                    }
                                }
                            }
                        case "a":
                            var v109 float64 = -2.25
                            _ = v109
                        }
                    }
                    _ = v98
                }
                _ = fi91
                _ = v96
            }
            v110 := GenS12{F0: 255, F1: (v51 % ((v3) | 1))}
            _ = v110
        } else {
            v51 = v3
        }
    }
    for fi111 := 0; fi111 < 5; fi111++ {
        var v112 int8 = v2
        _ = fi111
        _ = v112
    }
    if (32768 < 32768) {
        var v113 float64 = float64(0.1)
        var v114 int8 = v2
        _ = v113
        _ = v114
    } else {
        v115 := GenS17{F0: 4294967295, F1: 127, F2: 2147483647}
        _ = v115
    }
    if (1 > -169) {
        if false {
            switch (("hello" + "") + "foo bar") {
            case "hello":
                fmt.Println(v51)
            case "a":
                for fi116 := 0; fi116 < 9; fi116++ {
                    var v117 int8 = v2
                    _ = fi116
                    _ = v117
                }
            default:
                v118 := GenS3{F0: v3, F1: 42}
                _ = v118
            }
            v1[9223372036854775807] = ("a" + "a")
        }
        for fi119 := 0; fi119 < 12; fi119++ {
            if (257 != 566) {
                v3 += (v3 / ((v3) | 1))
                switch (fi119 % ((fi119) | 1)) {
                case 9223372036854775807:
                    delete(v1, fi119)
                case 745:
                    if true {
                        for fi120 := 0; fi120 < 20; fi120++ {
                            if false {
                                if (fi120 <= (fi120 << (((v3 % ((v3) | 1))) & 63))) {
                                    fmt.Println(v3)
                                    switch 32768 {
                                    case -234:
                                        v51 = (v3 | v3)
                                    }
                                    v51 += 42
                                }
                            } else {
                                var v121 int8 = v2
                                _ = v121
                            }
                            if false {
                                _ = len(v4)
                                var v122 uint8 = 11
                                var v123 uint16 = 2
                                var v124 string = "a"
                                _ = v122
                                _ = v123
                                _ = v124
                            }
                            _ = fi120
                        }
                        v2 = v2
                    } else {
                        v125 := GenS18{F0: 4294967295, F1: ("\n\t" == "a")}
                        _ = v125
                    }
                default:
                    fmt.Println(v51, v51)
                }
            }
            _ = fi119
        }
        _ = len(v4)
    } else {
        v3 += v3
        for fi126 := 0; fi126 < 10; fi126++ {
            var v127 int64 = -9223372036854775808
            _ = fi126
            _ = v127
        }
        var v128 uint32 = 2147483647
        _ = v128
    }
    delete(v5, 2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v54
}
