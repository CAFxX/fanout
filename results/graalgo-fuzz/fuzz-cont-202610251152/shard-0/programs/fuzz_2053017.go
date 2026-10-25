package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int16; F1 int8; F2 int16 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8; F1 int; F2 uint8 }
type GenS6 struct { F0 int16; F1 uint }
type GenS7 struct { F0 int64; F1 int64; F2 uint16 }
type GenS8 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS9 struct { F0 bool; F1 int32; F2 int64 }
type GenS10 struct { F0 uint32; F1 bool; F2 uint32 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS13 struct { F0 int64; F1 uint }
type GenS14 struct { F0 uint8 }

func main() {
    v1 := []uint32{2147483648, 361}
    v2 := make(map[uint32]uint32)
    var v3 int32 = 115
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]int)
    var v6 int64 = -9223372036854775808
    var v7 int16 = 1000
    v3 += -2147483648
    var v8 bool = (!(0.1 < 1.5))
    v9 := GenS1{F0: "\n\t", F1: "hello"}
    if v8 {
        for fi10 := 0; fi10 < 4; fi10++ {
            fmt.Println(fi10, fi10)
            v11 := GenS2{F0: v7, F1: int8(v7), F2: v7}
            fmt.Println(fi10)
            _ = fi10
            _ = v11
        }
    } else {
        if (!(0 <= 255)) {
            v6 = -382
            if (v8 && (v8 || (!(0.0 == 0.1)))) {
                var v12 int = 0
                _ = v12
            } else {
                fmt.Println(v7)
            }
        }
    }
    switch 908 {
    case -368:
        v8 = (v8 && (65535 > 1))
    case 4294967295:
        v6 = (v6 * 0)
    default:
        if v8 {
            if v8 {
                var v13 int32 = v3
                for fi14 := 0; fi14 < 18; fi14++ {
                    var v15 int64 = v6
                    var v16 int64 = 1
                    for fi17 := 0; fi17 < 17; fi17++ {
                        var v18 float32 = float32(float64((3.14159 / 0.1)))
                        _ = fi17
                        _ = v18
                    }
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                fmt.Println(v3)
                _ = v13
            }
            switch ("" + (("foo bar" + "") + "\n\t")) {
            case "a":
                switch "a" {
                case "a":
                    fmt.Println(v3, v8)
                case "":
                    v7 = 0
                case "\n\t":
                    if (0 >= 271) {
                        v19 := GenS1{F0: ("a" + "a"), F1: "\n\t"}
                        for fi20 := 0; fi20 < 11; fi20++ {
                            fmt.Println(v8, v7)
                            if v8 {
                                v6 -= (-(v6))
                                v21 := GenS1{F0: "foo bar", F1: ("a" + "\n\t")}
                                switch "\n\t" {
                                case "":
                                    v3 -= v3
                                case "foo bar":
                                    v1 = append(v1, 4294967295)
                                default:
                                    if len(v1) > 0 {
                                        v1[0] = 42
                                    }
                                }
                                _ = v21
                            } else {
                                switch fi20 {
                                case -2147483648:
                                    if (v8 || (-128 == 0)) {
                                        switch fi20 {
                                        case -9223372036854775808:
                                            if v8 {
                                                for fi22 := 0; fi22 < 11; fi22++ {
                                                    switch (fi20 & (fi20 & fi22)) {
                                                    case 679:
                                                        v23 := GenS3{F0: (v7 * int16(v3)), F1: 638}
                                                        var v24 int64 = 0
                                                        var v25 bool = v8
                                                        _ = v23
                                                        _ = v24
                                                        _ = v25
                                                    default:
                                                        if v8 {
                                                            var v26 int = 2147483647
                                                            v27 := GenS4{F0: v3}
                                                            if false {
                                                                v28 := GenS4{F0: -2147483648}
                                                                v3 -= (v3 % (((v3 / ((v3) | 1))) | 1))
                                                                delete(v4, ("a" + ("foo bar" + "foo bar")))
                                                                _ = v28
                                                            }
                                                            v29 := GenS4{F0: v3}
                                                            _ = v26
                                                            _ = v27
                                                            _ = v29
                                                        } else {
                                                            v6 = -1
                                                            v30 := GenS1{F0: ("\n\t" + "foo bar"), F1: ("a" + (((("a" + ("hello" + "")) + ("\n\t" + "foo bar")) + "\n\t") + ("\n\t" + ("\n\t" + "hello"))))}
                                                            _ = v30
                                                        }
                                                    }
                                                    switch "a" {
                                                    case "hello":
                                                        if ("foo bar" < "hello") {
                                                            var v31 int16 = v7
                                                            var v32 uint32 = 0
                                                            fmt.Println(v6, v8)
                                                            _ = v31
                                                            _ = v32
                                                        } else {
                                                            v6 += v6
                                                        }
                                                    case "foo bar":
                                                        var v33 uint64 = 2
                                                        _ = v33
                                                    default:
                                                        v1 = append(v1, 86)
                                                    }
                                                    if (42 <= 65535) {
                                                        switch fi22 {
                                                        case -256:
                                                            if (true && (255 > 427)) {
                                                                var v34 float64 = (3.14159 - 0.0)
                                                                if v8 {
                                                                    switch (fi22 % (((fi20 + (fi20 / ((fi20) | 1)))) | 1)) {
                                                                    case 0:
                                                                        if v8 {
                                                                            switch fi20 {
                                                                            case 42:
                                                                                v34 += (1.5 * v34)
                                                                            case -2147483648:
                                                                                switch fi22 {
                                                                                case 128:
                                                                                    v34 -= 1e10
                                                                                default:
                                                                                    v1 = append(v1, 2147483647)
                                                                                }
                                                                            case 127:
                                                                                fmt.Println(fi22, v8)
                                                                            default:
                                                                                switch fi20 {
                                                                                case 2147483647:
                                                                                    v1 = append(v1, 42)
                                                                                default:
                                                                                    v34 = v34
                                                                                    v35 := GenS2{F0: v7, F1: 127, F2: (v7 >> ((42) & 63))}
                                                                                    _ = v35
                                                                                }
                                                                            }
                                                                            fmt.Println(fi20)
                                                                            switch -752 {
                                                                            case 9223372036854775807:
                                                                                v1 = append(v1, 1)
                                                                                v6 = v6
                                                                            case 4294967295:
                                                                                if (32767 < 0) {
                                                                                    if true {
                                                                                        v34 = (((float64(1e10) * v34) * v34) - float64(1e10))
                                                                                        if (75 > -93) {
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = 4294967295
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(v3)
                                                                                    delete(v4, "\n\t")
                                                                                }
                                                                            case 1:
                                                                                v7 -= int16(v6)
                                                                            }
                                                                            if (!(v7 < v7)) {
                                                                                v36 := GenS4{F0: v3}
                                                                                v37 := v4["foo bar"]
                                                                                var v38 int8 = 1
                                                                                v39 := GenS1{F0: (("hello" + (("\n\t" + "a") + ("\n\t" + "a"))) + "a"), F1: (("" + "a") + "a")}
                                                                                v40 := GenS3{F0: -574, F1: v37}
                                                                                v41 := GenS2{F0: v7, F1: (v38 / ((v38) | 1)), F2: v7}
                                                                                if (1 > 255) {
                                                                                    v42 := GenS4{F0: (v3 & (^(v3)))}
                                                                                    _ = v42
                                                                                }
                                                                                if false {
                                                                                    if (1 > 4294967295) {
                                                                                        var v43 bool = v8
                                                                                        _ = v43
                                                                                    } else {
                                                                                        fmt.Println(fi20, v37, v34)
                                                                                    }
                                                                                    _ = len(v4)
                                                                                    v34 -= v34
                                                                                } else {
                                                                                    v44 := GenS5{F0: (v38 ^ 42), F1: (-(fi20)), F2: 241}
                                                                                    _ = v44
                                                                                }
                                                                                _ = v36
                                                                                _ = v37
                                                                                _ = v38
                                                                                _ = v39
                                                                                _ = v40
                                                                                _ = v41
                                                                            } else {
                                                                                if (0 < (v6 << ((267) & 63))) {
                                                                                    v6 = -9223372036854775808
                                                                                } else {
                                                                                    switch 128 {
                                                                                    case 0:
                                                                                        if false {
                                                                                            switch "\n\t" {
                                                                                            case "hello":
                                                                                                var v45 uint8 = 128
                                                                                                _ = v45
                                                                                            case "":
                                                                                                var v46 bool = true
                                                                                                _ = v46
                                                                                            case "a":
                                                                                                if (v8 || ("\n\t" < "")) {
                                                                                                    fmt.Println(v8)
                                                                                                    v47 := GenS1{F0: "a", F1: ""}
                                                                                                    var v48 bool = (42 > (v7 >> ((42) & 63)))
                                                                                                    _ = v47
                                                                                                    _ = v48
                                                                                                } else {
                                                                                                    v49 := v4[("\n\t" + "hello")]
                                                                                                    var v50 int32 = int32(v49)
                                                                                                    if true {
                                                                                                        _ = len(v5)
                                                                                                    } else {
                                                                                                        v6 = v6
                                                                                                    }
                                                                                                    v8 = (-1 > -47)
                                                                                                    _ = v49
                                                                                                    _ = v50
                                                                                                }
                                                                                            default:
                                                                                                var v51 string = ("a" + "")
                                                                                                _ = v51
                                                                                            }
                                                                                            fmt.Println(v6, v8, fi22)
                                                                                        }
                                                                                    case -1:
                                                                                        v52 := v5[0]
                                                                                        _ = v52
                                                                                    case 2147483647:
                                                                                        v34 -= 0.0
                                                                                    }
                                                                                    var v53 float32 = 3.14159
                                                                                    v53 = 0.1
                                                                                    v54 := GenS4{F0: (v3 << ((1) & 63))}
                                                                                    _ = v53
                                                                                    _ = v54
                                                                                }
                                                                            }
                                                                        }
                                                                    default:
                                                                        if ((v3 * v3) == (v3 - 579)) {
                                                                            if (2 == 2) {
                                                                                if v8 {
                                                                                    if v8 {
                                                                                        v7 -= v7
                                                                                        fmt.Println(fi20, v7, fi20)
                                                                                        fmt.Println(v34, v6, fi20)
                                                                                    } else {
                                                                                        var v55 uint64 = 571
                                                                                        v1 = append(v1, 277)
                                                                                        v34 = (v34 + (3.14159 - (0.0 * ((1.5 * float64(float32(v34))) * -2.25))))
                                                                                        _ = v55
                                                                                    }
                                                                                } else {
                                                                                    v6 += 9223372036854775807
                                                                                    v56 := GenS6{F0: int16(v3), F1: 18446744073709551615}
                                                                                    _ = v56
                                                                                }
                                                                            } else {
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = 806
                                                                                }
                                                                            }
                                                                            if (v8 && true) {
                                                                                if true {
                                                                                    if v8 {
                                                                                        fmt.Println(v34)
                                                                                        v6 -= v6
                                                                                        v7 += v7
                                                                                    }
                                                                                } else {
                                                                                    v57 := GenS6{F0: 42, F1: 42}
                                                                                    switch ((("foo bar" + "hello") + (("\n\t" + "hello") + ((("\n\t" + ("foo bar" + "hello")) + "hello") + "hello"))) + ("foo bar" + "\n\t")) {
                                                                                    case "":
                                                                                        fmt.Println(v6, fi20, fi22)
                                                                                    default:
                                                                                        v1 = append(v1, 647)
                                                                                    }
                                                                                    var v58 uint64 = 42
                                                                                    _ = v57
                                                                                    _ = v58
                                                                                }
                                                                                if (65535 != 741) {
                                                                                    switch ("foo bar" + "a") {
                                                                                    case "a":
                                                                                        v59 := GenS4{F0: 42}
                                                                                        _ = v59
                                                                                    }
                                                                                    fmt.Println(fi20, fi20)
                                                                                } else {
                                                                                    fmt.Println(fi22)
                                                                                }
                                                                                v60 := GenS7{F0: (v6 * -9223372036854775808), F1: (-(v6)), F2: 42}
                                                                                _ = v60
                                                                            }
                                                                            if (v8 && (37 == 416)) {
                                                                                var v61 string = ("hello" + "hello")
                                                                                _ = v61
                                                                            }
                                                                        }
                                                                    }
                                                                    fmt.Println(v7, v34, v6)
                                                                } else {
                                                                    var v62 int64 = (v6 % (((-(v6))) | 1))
                                                                    v6 = v62
                                                                    if len(v1) > 0 {
                                                                        v1[0] = 4294967295
                                                                    }
                                                                    fmt.Println(v8)
                                                                    fmt.Println(v8, v62)
                                                                    var v63 string = "a"
                                                                    _ = v62
                                                                    _ = v63
                                                                }
                                                                _ = v34
                                                            } else {
                                                                v64 := GenS4{F0: 1}
                                                                v7 += v7
                                                                _ = v64
                                                            }
                                                        case -1:
                                                            if (((-2.25 / -2.25) * 1e10) < (1e10 - (3.14159 + 0.0))) {
                                                                v2[65535] = 2147483647
                                                            } else {
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            }
                                                        case 2147483647:
                                                            var v65 int32 = v3
                                                            v6 = (v6 << ((4294967295) & 63))
                                                            _ = v65
                                                        default:
                                                            var v66 int32 = (v3 | v3)
                                                            _ = v66
                                                        }
                                                    }
                                                    _ = fi22
                                                }
                                                for fi67 := 0; fi67 < 4; fi67++ {
                                                    var v68 uint32 = 42
                                                    var v69 int16 = v7
                                                    _ = fi67
                                                    _ = v68
                                                    _ = v69
                                                }
                                            }
                                        }
                                    } else {
                                        var v70 uint16 = 188
                                        _ = v70
                                    }
                                case 201:
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v71 float64 = ((-2.25 / (0.0 - (1.5 / 3.14159))) * 0.0)
                                        _ = v71
                                    case "hello":
                                        var v72 int8 = 0
                                        _ = v72
                                    default:
                                        v6 = 42
                                    }
                                }
                                v73 := GenS5{F0: -67, F1: (fi20 >> ((18446744073709551615) & 63)), F2: 2}
                                _ = v73
                            }
                            _ = fi20
                        }
                        for fi74 := 0; fi74 < 14; fi74++ {
                            if v8 {
                                var v75 string = ("foo bar" + "")
                                if true {
                                    v76 := GenS2{F0: (v7 / ((v7) | 1)), F1: 127, F2: v7}
                                    if v8 {
                                        delete(v4, "")
                                    } else {
                                        if false {
                                            if (-9223372036854775808 > v6) {
                                                var v77 string = v75
                                                v6 = (v6 - v6)
                                                var v78 int32 = 1
                                                _ = v77
                                                _ = v78
                                            }
                                            var v79 int64 = (v6 >> ((2) & 63))
                                            _ = v79
                                        } else {
                                            var v80 uint64 = 931
                                            for fi81 := 0; fi81 < 2; fi81++ {
                                                if false {
                                                    v80 -= (v80 / ((v80) | 1))
                                                    if v8 {
                                                        v4[("hello" + v75)] = 2
                                                        if (128 > 2) {
                                                            if true {
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                switch fi74 {
                                                                case 1:
                                                                    if false {
                                                                        v1 = append(v1, 65535)
                                                                        var v82 uint64 = (v80 | (v80 | (v80 | 245)))
                                                                        if v8 {
                                                                            var v83 int16 = (v7 % ((v7) | 1))
                                                                            _ = v83
                                                                        }
                                                                        v75 = "\n\t"
                                                                        switch v75 {
                                                                        case "":
                                                                            var v84 int8 = 52
                                                                            _ = v84
                                                                        case "hello":
                                                                            v6 -= 0
                                                                        case "a":
                                                                            v85 := GenS5{F0: -128, F1: (fi81 << ((42) & 63)), F2: 2}
                                                                            _ = v85
                                                                        default:
                                                                            if ((0.0 * (-2.25 + (1.5 + 0.1))) >= (0.0 - 0.0)) {
                                                                                var v86 uint16 = 853
                                                                                if (v8 || false) {
                                                                                    if (!(v8 || v8)) {
                                                                                        fmt.Println(v3)
                                                                                        v7 = v7
                                                                                    }
                                                                                    v87 := v4[v75]
                                                                                    switch (fi74 | (fi74 % (((fi74 & fi81)) | 1))) {
                                                                                    case -499:
                                                                                        switch fi74 {
                                                                                        case -795:
                                                                                            fmt.Println(v80)
                                                                                        default:
                                                                                            switch "hello" {
                                                                                            case "":
                                                                                                switch fi74 {
                                                                                                case 2147483647:
                                                                                                    v88 := GenS8{F0: (-(v3)), F1: 128, F2: 128}
                                                                                                    _ = v88
                                                                                                case 42:
                                                                                                    v75 = ("" + "")
                                                                                                case 127:
                                                                                                    var v89 int64 = (v6 + v6)
                                                                                                    _ = v89
                                                                                                }
                                                                                            default:
                                                                                                _ = len(v1)
                                                                                            }
                                                                                        }
                                                                                    case 9223372036854775807:
                                                                                        v8 = v8
                                                                                    }
                                                                                    _ = v87
                                                                                }
                                                                                var v90 int = (fi74 & (fi74 ^ fi81))
                                                                                _ = v86
                                                                                _ = v90
                                                                            } else {
                                                                                v8 = ((v8 && (!v8)) || (3.14159 < (3.14159 + 1e10)))
                                                                                var v91 string = "foo bar"
                                                                                v92 := GenS9{F0: ((!((fi81 / ((fi81) | 1)) == (fi81 ^ fi81))) && ((255 <= uint8(v82)) && true)), F1: v3, F2: v6}
                                                                                _ = v91
                                                                                _ = v92
                                                                            }
                                                                        }
                                                                        _ = v82
                                                                    }
                                                                default:
                                                                    switch "" {
                                                                    case "":
                                                                        if (0.0 <= -2.25) {
                                                                            _ = len(v2)
                                                                        }
                                                                    default:
                                                                        var v93 float32 = (0.0 - 0.0)
                                                                        var v94 int8 = int8(v3)
                                                                        _ = v93
                                                                        _ = v94
                                                                    }
                                                                }
                                                            }
                                                            var v95 uint8 = 1
                                                            v96 := GenS10{F0: 114, F1: v8, F2: 2147483647}
                                                            var v97 bool = false
                                                            _ = v95
                                                            _ = v96
                                                            _ = v97
                                                        }
                                                    }
                                                    var v98 bool = (1e10 > (1.5 * 0.0))
                                                    fmt.Println(fi81)
                                                    v1 = append(v1, 622)
                                                    delete(v5, 329)
                                                    _ = v98
                                                }
                                                var v99 uint8 = 127
                                                _ = fi81
                                                _ = v99
                                            }
                                            _ = v80
                                        }
                                        var v100 float32 = -2.25
                                        switch ((v75 + v75) + v75) {
                                        case "a":
                                            var v101 uint8 = 128
                                            _ = v101
                                        }
                                        _ = v100
                                    }
                                    _ = v76
                                } else {
                                    var v102 float32 = (-2.25 - 1e10)
                                    v103 := GenS1{F0: "hello", F1: v75}
                                    var v104 string = v75
                                    _ = v102
                                    _ = v103
                                    _ = v104
                                }
                                var v105 bool = v8
                                _ = v105
                                _ = v75
                            }
                            if v8 {
                                fmt.Println(v7, fi74, v8)
                            } else {
                                if v8 {
                                    var v106 uint64 = 391
                                    for fi107 := 0; fi107 < 1; fi107++ {
                                        switch "a" {
                                        case "foo bar":
                                            if (v6 == v6) {
                                                v108 := GenS9{F0: (false && v8), F1: v3, F2: (v6 & (v6 & (v6 ^ (-(v6)))))}
                                                var v109 int32 = (v3 - (v3 / (((v3 & -1)) | 1)))
                                                _ = v108
                                                _ = v109
                                            }
                                        case "a":
                                            _ = len(v4)
                                        case "\n\t":
                                            switch 12 {
                                            case -2147483648:
                                                var v110 uint64 = (v106 / (((v106 * 0)) | 1))
                                                var v111 uint16 = 781
                                                v3 -= (v3 << ((0) & 63))
                                                _ = v110
                                                _ = v111
                                            case 42:
                                                switch "a" {
                                                case "\n\t":
                                                    if ((v8 && (1.5 <= (0.1 - float32((1e10 / 1.5))))) || (-1 <= -108)) {
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            var v112 string = "hello"
                                                            _ = v112
                                                        default:
                                                            fmt.Println(v6)
                                                        }
                                                    } else {
                                                        delete(v2, 1)
                                                    }
                                                case "hello":
                                                    delete(v4, "foo bar")
                                                case "a":
                                                    v113 := GenS11{F0: -808}
                                                    _ = v113
                                                }
                                            }
                                        default:
                                            var v114 string = ("hello" + "foo bar")
                                            _ = v114
                                        }
                                        var v115 int64 = 42
                                        v1 = append(v1, 2147483647)
                                        _ = fi107
                                        _ = v115
                                    }
                                    var v116 uint64 = v106
                                    _ = v106
                                    _ = v116
                                }
                            }
                            if len(v1) > 0 {
                                v1[0] = 1
                            }
                            _ = fi74
                        }
                        _ = v19
                    }
                }
            case "hello":
                fmt.Println(v8, v3)
            default:
                var v117 int32 = int32(v6)
                _ = v117
            }
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
        } else {
            if (2147483647 <= (v3 << ((822) & 63))) {
                var v118 int32 = v3
                v119 := GenS6{F0: 0, F1: 364}
                _ = v118
                _ = v119
            } else {
                var v120 uint8 = 128
                v6 = (v6 + v6)
                _ = v120
            }
        }
    }
    for fi121 := 0; fi121 < 16; fi121++ {
        var v122 bool = (-323 <= v7)
        if len(v1) > 0 {
            v1[0] = 0
        }
        switch ((((("" + "\n\t") + ("" + ("" + ""))) + "foo bar") + "") + "a") {
        case "":
            if (1e10 <= 3.14159) {
                var v123 int = int(v6)
                _ = v123
            }
        case "a":
            switch "" {
            case "hello":
                if (v7 == 782) {
                    delete(v2, 2147483647)
                    var v124 int = 42
                    var v125 int32 = (^(v3))
                    _ = v124
                    _ = v125
                } else {
                    for fi126 := 0; fi126 < 8; fi126++ {
                        v7 = 0
                        v127 := GenS12{F0: 2, F1: 489, F2: 1}
                        _ = fi126
                        _ = v127
                    }
                    var v128 int16 = 550
                    _ = v128
                }
            }
        case "hello":
            for fi129 := 0; fi129 < 18; fi129++ {
                v7 = v7
                fmt.Println(fi129)
                var v130 int8 = 0
                _ = fi129
                _ = v130
            }
        default:
            v131 := GenS13{F0: (v6 ^ v6), F1: 196}
            _ = v131
        }
        _ = fi121
        _ = v122
    }
    var v132 uint32 = 2147483647
    if false {
        if v8 {
            for fi133 := 0; fi133 < 4; fi133++ {
                var v134 int32 = int32(v7)
                v8 = (!(1e10 >= (1e10 - float64((((((0.0 / 1e10) * (-2.25 + 1.5)) - float32(1e10)) / 1.0) / float32(((0.1 / ((0.1 * 0.1) * -2.25)) / 1e10)))))))
                _ = len(v5)
                _ = fi133
                _ = v134
            }
            var v135 uint8 = 240
            _ = v135
        } else {
            v136 := GenS14{F0: 255}
            v8 = (!(255 <= 1))
            for fi137 := 0; fi137 < 8; fi137++ {
                if false {
                    v3 -= (^(v3))
                }
                for fi138 := 0; fi138 < 3; fi138++ {
                    var v139 string = ""
                    v3 = (^(v3))
                    _ = fi138
                    _ = v139
                }
                _ = fi137
            }
            _ = v136
        }
    } else {
        v140 := GenS2{F0: -852, F1: 42, F2: (v7 << ((65535) & 63))}
        switch -149 {
        case -1:
            var v141 uint16 = 255
            _ = v141
        case 443:
            fmt.Println(v3, v132)
        case 4294967295:
            if len(v1) > 0 {
                v1[0] = (v132 % ((v132) | 1))
            }
        }
        fmt.Println(v7)
        _ = v140
    }
    switch "foo bar" {
    case "hello":
        for fi142 := 0; fi142 < 20; fi142++ {
            v132 = v132
            for fi143 := 0; fi143 < 18; fi143++ {
                if ("\n\t" < ("a" + (("\n\t" + "hello") + ""))) {
                    var v144 float32 = (1.5 * float32(1e10))
                    _ = v144
                } else {
                    v7 = v7
                    var v145 int64 = -1
                    _ = v145
                }
                var v146 uint = 42
                _ = fi143
                _ = v146
            }
            _ = fi142
        }
    case "a":
        switch int(v7) {
        case -9223372036854775808:
            var v147 int = 209
            _ = v147
        case 622:
            v8 = (v7 <= v7)
        case 1:
            var v148 uint32 = (v132 ^ (^(v132)))
            _ = v148
        }
    case "":
        var v149 float64 = 0.0
        _ = v149
    }
    for fi150 := 0; fi150 < 5; fi150++ {
        for fi151 := 0; fi151 < 9; fi151++ {
            if (!(false || (-2147483648 == (fi151 << ((255) & 63))))) {
                v6 += v6
                switch 32768 {
                case 4294967295:
                    var v152 float32 = 1.5
                    _ = v152
                case -1:
                    _ = len(v1)
                case -2147483648:
                    var v153 uint32 = v132
                    _ = v153
                default:
                    var v154 uint16 = 42
                    _ = v154
                }
            }
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
            _ = fi151
        }
        _ = fi150
    }
    for fi155 := 0; fi155 < 5; fi155++ {
        for fi156 := 0; fi156 < 6; fi156++ {
            fmt.Println(v6, v6, v7)
            fmt.Println(v6, v3, v7)
            var v157 bool = ((3.14159 == float32(1e10)) && v8)
            _ = fi156
            _ = v157
        }
        var v158 uint = 0
        v1 = append(v1, 1)
        _ = fi155
        _ = v158
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v132:", v132)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v132
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
