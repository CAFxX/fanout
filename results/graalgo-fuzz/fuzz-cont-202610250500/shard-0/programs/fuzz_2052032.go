package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint64 }
type GenS2 struct { F0 uint16; F1 int8; F2 int8 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int8; F1 bool; F2 int32 }
type GenS9 struct { F0 int16; F1 uint }
type GenS10 struct { F0 uint16; F1 uint8 }
type GenS11 struct { F0 uint8; F1 int }
type GenS12 struct { F0 int8; F1 string; F2 bool }
type GenS13 struct { F0 int16; F1 int64; F2 uint32 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 int64; F1 int8; F2 uint }
type GenS16 struct { F0 string; F1 int16; F2 uint32 }
type GenS17 struct { F0 int16 }
type GenS18 struct { F0 uint16; F1 string; F2 string }
type GenS19 struct { F0 uint8; F1 bool; F2 int }
type GenS20 struct { F0 uint }
type GenS21 struct { F0 int64 }
type GenS22 struct { F0 bool; F1 bool }
type GenS23 struct { F0 uint16; F1 int }
func genFn0(p0 int, p1 string) uint32 {
    var v6 int = 1
    p0 = -2147483648
    var v7 float32 = (float32((0.0 / (-2.25 / 1.0))) + (float32(1.5) - -2.25))
    var v8 int32 = 1
    for fi9 := 0; fi9 < 15; fi9++ {
        if (false || (42 > 1)) {
            var v10 string = (("a" + (p1 + p1)) + "\n\t")
            for fi11 := 0; fi11 < 14; fi11++ {
                var v12 int8 = -67
                _ = fi11
                _ = v12
            }
            _ = v10
        } else {
            v13 := []int{(fi9 * (v6 - (v6 % ((fi9) | 1)))), (fi9 * 649), (p0 & fi9)}
            v6 -= p0
            switch "\n\t" {
            case "hello":
                if true {
                    v7 -= (((-2.25 + ((v7 * ((v7 * (v7 - 1.5)) / v7)) * 1.5)) + v7) / (v7 + v7))
                    fmt.Println(p0)
                }
            case "foo bar":
                for fi14 := 0; fi14 < 20; fi14++ {
                    v6 = fi14
                    v7 = v7
                    _ = fi14
                }
            case "\n\t":
                v8 = v8
            default:
                v15 := GenS1{F0: false, F1: "a", F2: 1}
                _ = v15
            }
            _ = v13
        }
        var v16 uint16 = 32767
        var v17 int32 = (v8 % ((v8) | 1))
        _ = fi9
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := []uint32{}
    var v2 int16 = 42
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 42
    var v5 bool = false
    for fi18 := 0; fi18 < 11; fi18++ {
        var v19 int = fi18
        _ = fi18
        _ = v19
    }
    var v20 uint32 = genFn0(0, "\n\t")
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v21 int8 = 100
    v2 -= v2
    v5 = (-2147483648 < 2147483647)
    for fi22 := 0; fi22 < 14; fi22++ {
        switch (fi22 + fi22) {
        case 42:
            for fi23 := 0; fi23 < 2; fi23++ {
                if ((!v5) && (v5 || false)) {
                    if (v5 || v5) {
                        if (!(-2.25 > (-2.25 + 1e10))) {
                            var v24 bool = v5
                            v25 := GenS2{F0: 32767, F1: (v21 ^ v21), F2: v21}
                            v20 = (^(v20))
                            _ = v24
                            _ = v25
                        }
                        var v26 float64 = (-2.25 * 3.14159)
                        _ = v26
                    }
                    var v27 bool = (2 > genFn0((fi23 + fi22), ("foo bar" + (("" + (("" + "") + "hello")) + ((("hello" + "a") + "\n\t") + "foo bar")))))
                    _ = v27
                } else {
                    var v28 int64 = 0
                    _ = v28
                }
                _ = fi23
            }
        case 127:
            switch "" {
            case "a":
                v29 := GenS1{F0: false, F1: "\n\t", F2: 47}
                _ = v29
            default:
                var v30 int64 = -9223372036854775808
                _ = v30
            }
        }
        switch fi22 {
        case 0:
            for fi31 := 0; fi31 < 11; fi31++ {
                v32 := GenS1{F0: true, F1: (("a" + "hello") + ("" + "hello")), F2: 49}
                var v33 int = fi22
                v21 += (v21 ^ (^(v21)))
                _ = fi31
                _ = v32
                _ = v33
            }
        case 9223372036854775807:
            for fi34 := 0; fi34 < 8; fi34++ {
                v35 := GenS2{F0: 1, F1: (-(v21)), F2: v21}
                v3[genFn0(fi34, "a")] = (^(v20))
                _ = fi34
                _ = v35
            }
        case 42:
            var v36 uint64 = 0
            _ = v36
        default:
            if ((v21 | (v21 >> ((146) & 63))) >= v21) {
                v5 = v5
                var v37 int8 = (-(v21))
                _ = v37
            } else {
                switch fi22 {
                case 42:
                    delete(v3, (v20 / ((v20) | 1)))
                }
            }
        }
        _ = fi22
    }
    if true {
        if (1 <= 1) {
            var v38 int = 32767
            var v39 uint32 = v20
            var v40 uint32 = v39
            _ = v38
            _ = v39
            _ = v40
        } else {
            v4 += 127
            v5 = (genFn0(-305, "foo bar") >= 664)
        }
        switch 887 {
        case -925:
            switch 884 {
            case 0:
                var v41 int32 = 461
                _ = v41
            case 9223372036854775807:
                switch 648 {
                case 42:
                    if false {
                        switch 1 {
                        case -377:
                            v42 := GenS1{F0: ((v20 >> ((65535) & 63)) == v20), F1: ("a" + "foo bar"), F2: 119}
                            _ = v42
                        case 128:
                            switch "foo bar" {
                            case "":
                                var v43 float32 = float32(1e10)
                                _ = v43
                            case "a":
                                fmt.Println(v21, v20, v20)
                            case "hello":
                                v2 -= v2
                            default:
                                for fi44 := 0; fi44 < 12; fi44++ {
                                    v45 := GenS1{F0: false, F1: "", F2: 987}
                                    v21 += (v21 % (((v21 + 98)) | 1))
                                    _ = fi44
                                    _ = v45
                                }
                            }
                        case -502:
                            _ = len(v3)
                        default:
                            if true {
                                v2 += (v2 >> ((65535) & 63))
                            } else {
                                v4 -= v4
                                _ = len(v1)
                            }
                        }
                        switch -174 {
                        case 42:
                            for fi46 := 0; fi46 < 4; fi46++ {
                                switch 468 {
                                case 4294967295:
                                    v2 = (v2 + (v2 & v2))
                                case 885:
                                    v1 = append(v1, (v20 + genFn0((-(fi46)), "foo bar")))
                                default:
                                    var v47 float32 = 1.5
                                    _ = v47
                                }
                                v48 := GenS3{F0: "hello", F1: (v2 % ((v2) | 1))}
                                var v49 uint32 = v20
                                _ = fi46
                                _ = v48
                                _ = v49
                            }
                        case -674:
                            v50 := GenS4{F0: 32768, F1: 2147483647}
                            _ = v50
                        case 306:
                            var v51 int16 = 32767
                            _ = v51
                        default:
                            if true {
                                fmt.Println(v5)
                                fmt.Println(v5, v4)
                                var v52 uint8 = 65
                                _ = v52
                            }
                        }
                        for fi53 := 0; fi53 < 14; fi53++ {
                            v4 -= v4
                            _ = fi53
                        }
                    } else {
                        switch 127 {
                        case 128:
                            for fi54 := 0; fi54 < 12; fi54++ {
                                switch 871 {
                                case 711:
                                    switch (fi54 * (fi54 / (((fi54 * (fi54 ^ fi54))) | 1))) {
                                    case 866:
                                        switch ("" + "a") {
                                        case "hello":
                                            var v55 bool = (!(!v5))
                                            _ = v55
                                        case "a":
                                            v4 = (v4 + 127)
                                        case "\n\t":
                                            for fi56 := 0; fi56 < 13; fi56++ {
                                                v57 := GenS5{F0: v5, F1: 18446744073709551615}
                                                var v58 float64 = 0.1
                                                v2 = 1000
                                                _ = fi56
                                                _ = v57
                                                _ = v58
                                            }
                                        default:
                                            for fi59 := 0; fi59 < 8; fi59++ {
                                                if v5 {
                                                    if (false || (-9223372036854775808 < 1)) {
                                                        _ = len(v1)
                                                        switch fi59 {
                                                        case -251:
                                                            if (0.0 <= (-2.25 * -2.25)) {
                                                                if false {
                                                                    v3[(^(v20))] = v20
                                                                    switch (fi54 ^ fi54) {
                                                                    case 0:
                                                                        fmt.Println(v2, v21)
                                                                    case 42:
                                                                        v5 = ((v2 ^ (v2 ^ 1000)) == (v2 | (v2 + v2)))
                                                                    default:
                                                                        switch fi54 {
                                                                        case 4294967295:
                                                                            if true {
                                                                                v60 := GenS1{F0: (0.1 == 0.1), F1: "foo bar", F2: 0}
                                                                                v61 := GenS1{F0: v5, F1: "\n\t", F2: 730}
                                                                                var v62 string = ((("" + ("" + "foo bar")) + (("a" + "hello") + "")) + "\n\t")
                                                                                _ = v60
                                                                                _ = v61
                                                                                _ = v62
                                                                            }
                                                                        }
                                                                    }
                                                                    delete(v3, (^(v20)))
                                                                } else {
                                                                    switch 128 {
                                                                    case 128:
                                                                        v2 -= 557
                                                                    case 9223372036854775807:
                                                                        var v63 bool = (1000 >= (v2 | (^(v2))))
                                                                        _ = v63
                                                                    case 953:
                                                                        v64 := GenS6{F0: int8(v4)}
                                                                        _ = v64
                                                                    default:
                                                                        var v65 int = fi54
                                                                        _ = v65
                                                                    }
                                                                    if ((genFn0(fi54, "\n\t") >= (v20 + v20)) || false) {
                                                                        fmt.Println(fi54, v4)
                                                                    } else {
                                                                        if (314 == 42) {
                                                                            var v66 float64 = float64((float32(1e10) - float32(0.1)))
                                                                            v67 := GenS3{F0: "a", F1: (v2 >> ((549) & 63))}
                                                                            _ = v66
                                                                            _ = v67
                                                                        }
                                                                        switch ("a" + "a") {
                                                                        case "\n\t":
                                                                            v68 := GenS7{F0: 89}
                                                                            _ = v68
                                                                        case "foo bar":
                                                                            v3[v20] = 42
                                                                        default:
                                                                            if (!(0 == 963)) {
                                                                                v69 := GenS5{F0: (255 == 0), F1: 640}
                                                                                _ = v69
                                                                            } else {
                                                                                fmt.Println(v21, v2)
                                                                                if true {
                                                                                    if false {
                                                                                        if false {
                                                                                            var v70 int16 = int16(v4)
                                                                                            fmt.Println(v4, fi59, v2)
                                                                                            _ = v70
                                                                                        } else {
                                                                                            _ = len(v1)
                                                                                        }
                                                                                    }
                                                                                    fmt.Println(v20)
                                                                                }
                                                                                var v71 uint32 = v20
                                                                                _ = v71
                                                                            }
                                                                        }
                                                                    }
                                                                    var v72 uint32 = (v20 + genFn0(32768, ("\n\t" + ("foo bar" + "hello"))))
                                                                    var v73 int16 = v2
                                                                    v73 -= (v2 % (((v2 >> ((734) & 63))) | 1))
                                                                    switch "" {
                                                                    case "a":
                                                                        v4 -= 128
                                                                    case "foo bar":
                                                                        v2 = (v73 >> ((uint(v4)) & 63))
                                                                    case "hello":
                                                                        v74 := GenS8{F0: v21, F1: true, F2: int32(v4)}
                                                                        _ = v74
                                                                    default:
                                                                        switch (fi54 * fi59) {
                                                                        case -695:
                                                                            if true {
                                                                                v75 := GenS9{F0: (v2 << ((695) & 63)), F1: 4294967295}
                                                                                switch ("hello" + "\n\t") {
                                                                                case "":
                                                                                    switch fi59 {
                                                                                    case 282:
                                                                                        if v5 {
                                                                                            var v76 uint = 714
                                                                                            fmt.Println(fi59, v72, v2)
                                                                                            _ = v76
                                                                                        } else {
                                                                                            if v5 {
                                                                                                v73 = v73
                                                                                            }
                                                                                            v3[2147483648] = v20
                                                                                            fmt.Println(fi59, fi54)
                                                                                            if (!false) {
                                                                                                fmt.Println(v20)
                                                                                                var v77 float64 = (0.0 + (1.5 + 1e10))
                                                                                                delete(v3, (v20 / ((v72) | 1)))
                                                                                                fmt.Println(fi59, v20)
                                                                                                var v78 int32 = int32(v4)
                                                                                                var v79 uint = 255
                                                                                                _ = v77
                                                                                                _ = v78
                                                                                                _ = v79
                                                                                            }
                                                                                        }
                                                                                    case -1:
                                                                                        switch "\n\t" {
                                                                                        case "":
                                                                                            var v80 uint8 = v4
                                                                                            _ = v80
                                                                                        case "foo bar":
                                                                                            v20 = v20
                                                                                        default:
                                                                                            v20 = genFn0(fi59, "foo bar")
                                                                                        }
                                                                                    case -648:
                                                                                        switch (fi59 | fi54) {
                                                                                        case 283:
                                                                                            v20 += v72
                                                                                            _ = len(v1)
                                                                                            fmt.Println(v4)
                                                                                        case 1:
                                                                                            var v81 int64 = -1
                                                                                            _ = v81
                                                                                        }
                                                                                        switch fi54 {
                                                                                        case -9223372036854775808:
                                                                                            v73 = v73
                                                                                        case 32768:
                                                                                            v20 = 42
                                                                                        case 369:
                                                                                            v82 := GenS2{F0: 539, F1: v21, F2: v21}
                                                                                            fmt.Println(v4)
                                                                                            if v5 {
                                                                                                if v5 {
                                                                                                    var v83 uint16 = 213
                                                                                                    if ((fi54 % ((fi54) | 1)) == int(v72)) {
                                                                                                        switch "foo bar" {
                                                                                                        case "\n\t":
                                                                                                            v20 = v20
                                                                                                            if v5 {
                                                                                                                switch (fi59 - (fi59 + 42)) {
                                                                                                                case -1:
                                                                                                                    fmt.Println(v73)
                                                                                                                case 5:
                                                                                                                    switch 127 {
                                                                                                                    case 42:
                                                                                                                        var v84 uint = 989
                                                                                                                        _ = v84
                                                                                                                    default:
                                                                                                                        var v85 int16 = 32767
                                                                                                                        _ = v85
                                                                                                                    }
                                                                                                                default:
                                                                                                                    fmt.Println(v20, v21)
                                                                                                                }
                                                                                                                v86 := GenS10{F0: v83, F1: (v4 / (((v4 & (v4 & 115))) | 1))}
                                                                                                                var v87 uint64 = 42
                                                                                                                _ = v86
                                                                                                                _ = v87
                                                                                                            }
                                                                                                            if v5 {
                                                                                                                if (89 > int(v83)) {
                                                                                                                    v88 := GenS11{F0: (v4 - (v4 - v4)), F1: (fi59 ^ fi59)}
                                                                                                                    fmt.Println(v83, v73, v5)
                                                                                                                    fmt.Println(v21, v72)
                                                                                                                    var v89 int = fi59
                                                                                                                    fmt.Println(v20, fi59)
                                                                                                                    _ = v88
                                                                                                                    _ = v89
                                                                                                                }
                                                                                                                switch (fi59 % (((fi59 + 1)) | 1)) {
                                                                                                                case 128:
                                                                                                                    v20 -= 767
                                                                                                                case -653:
                                                                                                                    v90 := v3[genFn0(2147483647, "foo bar")]
                                                                                                                    _ = v90
                                                                                                                }
                                                                                                            } else {
                                                                                                                delete(v3, (v20 / ((39) | 1)))
                                                                                                                fmt.Println(v73)
                                                                                                            }
                                                                                                        case "a":
                                                                                                            v20 = genFn0(fi54, "hello")
                                                                                                        }
                                                                                                        if v5 {
                                                                                                            delete(v3, 2147483647)
                                                                                                            v73 = v2
                                                                                                            v91 := GenS10{F0: uint16(v72), F1: (v4 / ((255) | 1))}
                                                                                                            _ = v91
                                                                                                        } else {
                                                                                                            v1 = append(v1, (v20 & v72))
                                                                                                            switch -124 {
                                                                                                            case 2147483647:
                                                                                                                if v5 {
                                                                                                                    v92 := GenS12{F0: (v21 & (v21 >> ((674) & 63))), F1: "hello", F2: (3.14159 == float64((1.5 - (0.1 + 1e10))))}
                                                                                                                    v93 := GenS13{F0: v2, F1: -248, F2: 2147483648}
                                                                                                                    if (1 != 42) {
                                                                                                                        v21 = -128
                                                                                                                        var v94 string = ("" + "a")
                                                                                                                        _ = len(v1)
                                                                                                                        _ = v94
                                                                                                                    } else {
                                                                                                                        if v5 {
                                                                                                                            v95 := GenS4{F0: 32768, F1: -1}
                                                                                                                            v4 = (^(v4))
                                                                                                                            _ = v95
                                                                                                                        } else {
                                                                                                                            var v96 int64 = 9223372036854775807
                                                                                                                            var v97 float64 = 3.14159
                                                                                                                            _ = v96
                                                                                                                            _ = v97
                                                                                                                        }
                                                                                                                    }
                                                                                                                    _ = v92
                                                                                                                    _ = v93
                                                                                                                }
                                                                                                            case 127:
                                                                                                                var v98 uint64 = 86
                                                                                                                v1 = append(v1, v20)
                                                                                                                _ = v98
                                                                                                            default:
                                                                                                                v99 := GenS9{F0: 483, F1: uint(v4)}
                                                                                                                _ = v99
                                                                                                            }
                                                                                                            v21 -= 0
                                                                                                        }
                                                                                                        v21 -= v21
                                                                                                    } else {
                                                                                                        switch (fi59 + -9223372036854775808) {
                                                                                                        case 0:
                                                                                                            v83 = (v83 & (^(v83)))
                                                                                                        }
                                                                                                        v5 = false
                                                                                                    }
                                                                                                    v100 := GenS14{F0: (v4 >> ((630) & 63))}
                                                                                                    _ = v100
                                                                                                    _ = v83
                                                                                                }
                                                                                                var v101 uint32 = genFn0((fi59 + 330), "a")
                                                                                                _ = v101
                                                                                            } else {
                                                                                                v2 = (v73 & 42)
                                                                                                v73 = (v73 / ((-1) | 1))
                                                                                            }
                                                                                            v1 = append(v1, (v20 << ((255) & 63)))
                                                                                            if v5 {
                                                                                                var v102 uint8 = (^(v4))
                                                                                                fmt.Println(v72)
                                                                                                _ = v102
                                                                                            } else {
                                                                                                var v103 bool = ("\n\t" != "hello")
                                                                                                var v104 int64 = -1
                                                                                                _ = v103
                                                                                                _ = v104
                                                                                            }
                                                                                            v105 := GenS15{F0: 400, F1: (-(v21)), F2: 1}
                                                                                            _ = v105
                                                                                            _ = v82
                                                                                        default:
                                                                                            var v106 string = "a"
                                                                                            _ = v106
                                                                                        }
                                                                                    }
                                                                                    var v107 bool = (v5 && (v21 != -1))
                                                                                    _ = v107
                                                                                case "a":
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = 1
                                                                                    }
                                                                                }
                                                                                switch "\n\t" {
                                                                                case "":
                                                                                    v73 = v2
                                                                                case "foo bar":
                                                                                    v108 := GenS16{F0: "\n\t", F1: v2, F2: (v20 >> ((1) & 63))}
                                                                                    _ = v108
                                                                                case "\n\t":
                                                                                    if (!(1 == 0)) {
                                                                                        v4 = (v4 / ((127) | 1))
                                                                                        v5 = (1000 != 1)
                                                                                        v109 := GenS17{F0: (v2 / (((v2 * -690)) | 1))}
                                                                                        _ = v109
                                                                                    }
                                                                                default:
                                                                                    v73 += (v73 + (v2 % (((v2 & v73)) | 1)))
                                                                                }
                                                                                _ = v75
                                                                            } else {
                                                                                switch "" {
                                                                                case "a":
                                                                                    var v110 int16 = 1000
                                                                                    _ = v110
                                                                                default:
                                                                                    if false {
                                                                                        switch "hello" {
                                                                                        case "foo bar":
                                                                                            if v5 {
                                                                                                v5 = v5
                                                                                                if len(v1) > 0 {
                                                                                                    v1[0] = (v20 * (v72 >> ((199) & 63)))
                                                                                                }
                                                                                            }
                                                                                        case "a":
                                                                                            var v111 uint8 = (v4 | uint8(v72))
                                                                                            var v112 uint64 = 4294967295
                                                                                            _ = v111
                                                                                            _ = v112
                                                                                        }
                                                                                    } else {
                                                                                        fmt.Println(v21)
                                                                                        if false {
                                                                                            var v113 int64 = 1
                                                                                            fmt.Println(v73)
                                                                                            _ = v113
                                                                                        } else {
                                                                                            var v114 string = ""
                                                                                            _ = v114
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                        case 9223372036854775807:
                                                                            v72 += v72
                                                                        default:
                                                                            v115 := GenS15{F0: 42, F1: v21, F2: 65535}
                                                                            _ = v115
                                                                        }
                                                                    }
                                                                    _ = v72
                                                                    _ = v73
                                                                }
                                                                v116 := GenS18{F0: 32767, F1: "hello", F2: ("" + ("foo bar" + ("foo bar" + "\n\t")))}
                                                                if (-2.25 < 0.0) {
                                                                    switch "foo bar" {
                                                                    case "foo bar":
                                                                        var v117 uint = 801
                                                                        _ = v117
                                                                    }
                                                                }
                                                                _ = v116
                                                            }
                                                        case -85:
                                                            v118 := GenS8{F0: (v21 ^ v21), F1: false, F2: 1}
                                                            _ = v118
                                                        case 127:
                                                            var v119 uint64 = 0
                                                            _ = v119
                                                        default:
                                                            switch (fi54 >> ((1) & 63)) {
                                                            case 704:
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            case 128:
                                                                v3[(v20 | 215)] = 2147483648
                                                                fmt.Println(v4)
                                                                if (!v5) {
                                                                    fmt.Println(v2, v2, fi59)
                                                                } else {
                                                                    fmt.Println(fi54)
                                                                }
                                                            }
                                                        }
                                                        var v120 int64 = 1
                                                        if (27 == (v21 + 1)) {
                                                            v120 -= (^(v120))
                                                            v20 += v20
                                                            v1 = append(v1, 308)
                                                            v1 = append(v1, genFn0((fi54 & 127), "hello"))
                                                            if len(v1) > 0 {
                                                                v1[0] = (v20 / (((v20 >> ((1) & 63))) | 1))
                                                            }
                                                        }
                                                        _ = v120
                                                    }
                                                    fmt.Println(v21, v2)
                                                    if v5 {
                                                        var v121 bool = v5
                                                        _ = v121
                                                    }
                                                    if v5 {
                                                        switch (-(fi54)) {
                                                        case 2147483647:
                                                            fmt.Println(fi54, v2)
                                                        case 1:
                                                            v3[v20] = genFn0((fi54 / ((fi54) | 1)), "foo bar")
                                                        case 42:
                                                            v20 = (^(v20))
                                                        }
                                                        var v122 int8 = -1
                                                        delete(v3, 42)
                                                        _ = v122
                                                    } else {
                                                        v123 := GenS17{F0: v2}
                                                        var v124 int16 = (v2 / (((v2 << ((18446744073709551615) & 63))) | 1))
                                                        var v125 string = "hello"
                                                        _ = len(v3)
                                                        _ = v123
                                                        _ = v124
                                                        _ = v125
                                                    }
                                                    var v126 int8 = v21
                                                    _ = v126
                                                }
                                                v21 = v21
                                                _ = fi59
                                            }
                                        }
                                    case -164:
                                        v127 := GenS17{F0: (v2 << ((255) & 63))}
                                        _ = v127
                                    case -144:
                                        v128 := GenS19{F0: v4, F1: (((((1e10 / 1e10) / -2.25) - ((3.14159 * (0.0 / (0.0 - 3.14159))) * (((-2.25 + (0.1 + 1e10)) - 1.5) / 1e10))) / -2.25) != 3.14159), F2: 9223372036854775807}
                                        _ = v128
                                    default:
                                        var v129 float32 = 1e10
                                        _ = v129
                                    }
                                default:
                                    if true {
                                        var v130 int8 = v21
                                        delete(v3, (^(v20)))
                                        _ = v130
                                    } else {
                                        switch fi54 {
                                        case -784:
                                            var v131 int64 = 903
                                            _ = v131
                                        default:
                                            if true {
                                                if (9223372036854775807 < 9223372036854775807) {
                                                    if v5 {
                                                        var v132 uint32 = v20
                                                        v133 := GenS19{F0: v4, F1: v5, F2: fi54}
                                                        var v134 uint = 255
                                                        _ = v132
                                                        _ = v133
                                                        _ = v134
                                                    } else {
                                                        var v135 uint64 = uint64(v4)
                                                        for fi136 := 0; fi136 < 15; fi136++ {
                                                            var v137 uint64 = 18446744073709551615
                                                            if (1 >= v4) {
                                                                v20 -= v20
                                                                var v138 float32 = 0.1
                                                                v20 -= (v20 % ((42) | 1))
                                                                _ = v138
                                                            } else {
                                                                v137 += uint64(v4)
                                                                var v139 int8 = (v21 / (((v21 | (v21 & (v21 + v21)))) | 1))
                                                                fmt.Println(fi54, fi136)
                                                                v21 -= 127
                                                                _ = v139
                                                            }
                                                            v140 := GenS20{F0: 0}
                                                            _ = fi136
                                                            _ = v137
                                                            _ = v140
                                                        }
                                                        _ = v135
                                                    }
                                                    v141 := GenS21{F0: 9223372036854775807}
                                                    v21 = v21
                                                    _ = v141
                                                } else {
                                                    if (!(("" + "foo bar") >= "a")) {
                                                        _ = len(v1)
                                                        var v142 bool = (342 == 58)
                                                        v143 := GenS7{F0: 255}
                                                        _ = v142
                                                        _ = v143
                                                    }
                                                    for fi144 := 0; fi144 < 12; fi144++ {
                                                        var v145 int32 = 373
                                                        switch "a" {
                                                        case "foo bar":
                                                            if v5 {
                                                                v20 -= v20
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            }
                                                        case "":
                                                            var v146 uint32 = genFn0((-(fi54)), "foo bar")
                                                            _ = v146
                                                        case "\n\t":
                                                            if (!v5) {
                                                                var v147 float64 = (0.1 + float64(float32(0.1)))
                                                                var v148 uint64 = 833
                                                                var v149 uint32 = 65535
                                                                _ = v147
                                                                _ = v148
                                                                _ = v149
                                                            } else {
                                                                v150 := GenS12{F0: (^(v21)), F1: (("hello" + "hello") + "a"), F2: v5}
                                                                var v151 int16 = -1
                                                                _ = v150
                                                                _ = v151
                                                            }
                                                        default:
                                                            v1 = append(v1, (v20 - (^(v20))))
                                                        }
                                                        if v5 {
                                                            v2 = (v2 | v2)
                                                        }
                                                        _ = fi144
                                                        _ = v145
                                                    }
                                                }
                                                fmt.Println(v5)
                                                v20 += genFn0((fi54 + fi54), "a")
                                            } else {
                                                switch ("\n\t" + "a") {
                                                case "a":
                                                    v2 = v2
                                                case "hello":
                                                    v3[genFn0(fi54, "hello")] = genFn0((fi54 & 1), ("" + (("" + "hello") + "hello")))
                                                default:
                                                    for fi152 := 0; fi152 < 16; fi152++ {
                                                        fmt.Println(v4)
                                                        _ = len(v1)
                                                        _ = fi152
                                                    }
                                                }
                                            }
                                        }
                                        var v153 int32 = 0
                                        v2 = v2
                                        _ = v153
                                    }
                                }
                                _ = fi54
                            }
                        case 2147483647:
                            v154 := GenS22{F0: v5, F1: ((!v5) && (v2 > (v2 + -384)))}
                            _ = v154
                        }
                    }
                case 128:
                    var v155 int8 = 6
                    _ = v155
                }
            }
        case -711:
            for fi156 := 0; fi156 < 13; fi156++ {
                var v157 bool = ((!((18446744073709551615 <= 988) && v5)) || true)
                _ = fi156
                _ = v157
            }
        case 2147483647:
            if (1 > 0) {
                if v5 {
                    var v158 int = 32767
                    _ = v158
                } else {
                    switch 4294967295 {
                    case 42:
                        fmt.Println(v20, v21, v2)
                    case -9223372036854775808:
                        v159 := GenS5{F0: v5, F1: 18446744073709551615}
                        _ = v159
                    default:
                        switch "foo bar" {
                        case "a":
                            v3[514] = (v20 ^ v20)
                        case "hello":
                            for fi160 := 0; fi160 < 20; fi160++ {
                                var v161 uint8 = (^(v4))
                                var v162 float32 = ((-2.25 - 0.1) - 1e10)
                                _ = fi160
                                _ = v161
                                _ = v162
                            }
                        }
                    }
                }
            }
        }
    }
    var v163 bool = (v5 || v5)
    v164 := GenS23{F0: 1, F1: 2147483647}
    for fi165 := 0; fi165 < 15; fi165++ {
        if (false || (2 >= 18446744073709551615)) {
            for fi166 := 0; fi166 < 10; fi166++ {
                switch fi165 {
                case 607:
                    var v167 uint16 = 32767
                    _ = v167
                default:
                    v21 = v21
                }
                v2 -= v2
                _ = fi166
            }
        } else {
            v4 -= 19
        }
        if v163 {
            fmt.Println(v2)
            v163 = (((^(fi165)) != (fi165 >> ((195) & 63))) && ((-(v21)) <= v21))
            switch ("a" + "") {
            case "hello":
                var v168 uint = 1
                _ = v168
            case "":
                var v169 uint = 843
                _ = v169
            case "foo bar":
                fmt.Println(fi165, v5)
            }
        }
        _ = fi165
    }
    var v170 float64 = 1.5
    if v5 {
        var v171 int32 = -1
        var v172 uint = 1
        _ = v171
        _ = v172
    } else {
        var v173 int = -841
        _ = v173
    }
    fmt.Println(v4)
    if (!v163) {
        v163 = v5
    }
    var v174 int16 = (v2 % ((v2) | 1))
    for fi175 := 0; fi175 < 7; fi175++ {
        v176 := GenS12{F0: v21, F1: "a", F2: (!(-32 > (fi175 % ((-180) | 1))))}
        if (!true) {
            var v177 uint = uint(v4)
            if v163 {
                v3[(v20 >> (((v177 & (v177 % ((354) | 1)))) & 63))] = 1
            }
            _ = v177
        } else {
            var v178 int = (fi175 * (-(fi175)))
            var v179 int = 2147483647
            var v180 string = "\n\t"
            _ = v178
            _ = v179
            _ = v180
        }
        v170 = (v170 * v170)
        _ = fi175
        _ = v176
    }
    if v163 {
        for fi181 := 0; fi181 < 1; fi181++ {
            if (v5 && true) {
                v3[(v20 / (((^(v20))) | 1))] = v20
                v20 += (v20 % ((v20) | 1))
            }
            var v182 bool = v5
            var v183 bool = false
            _ = fi181
            _ = v182
            _ = v183
        }
        v174 += v174
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v163:", v163)
    fmt.Println("v170:", v170)
    fmt.Println("v174:", v174)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v163
    _ = v164
    _ = v170
    _ = v174
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
