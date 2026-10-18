package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 int; F1 uint8; F2 int32 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint8; F1 int8 }
type GenS9 struct { F0 bool; F1 uint32 }
func genFn0() int {
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 12; fi8++ {
            var v9 uint16 = 2
            v9 -= (v9 * 32767)
            v9 += 42
            _ = fi8
            _ = v9
        }
        fmt.Println(fi7)
        _ = fi7
    }
    var v10 bool = true
    for fi11 := 0; fi11 < 14; fi11++ {
        for fi12 := 0; fi12 < 4; fi12++ {
            v10 = false
            switch fi11 {
            case 202:
                v13 := []int{-1, 1, fi11, fi12}
                _ = v13
            default:
                if (0 <= -65) {
                    if false {
                        v14 := GenS1{F0: 808, F1: 3, F2: v10}
                        if (v10 && (true || v10)) {
                            var v15 int16 = 0
                            fmt.Println(v10, fi11, v15)
                            if v10 {
                                v15 = (v15 ^ 1000)
                                fmt.Println(v15)
                            }
                            _ = v15
                        } else {
                            v16 := GenS1{F0: -615, F1: -426, F2: v10}
                            v17 := []int{215, (fi11 & (-(fi12))), -560}
                            _ = v16
                            _ = v17
                        }
                        var v18 uint = 1
                        _ = v14
                        _ = v18
                    } else {
                        v19 := GenS1{F0: 42, F1: -780, F2: (!false)}
                        v10 = (!(!v10))
                        _ = v19
                    }
                    if v10 {
                        v20 := GenS1{F0: 1, F1: 375, F2: v10}
                        v21 := GenS2{F0: 1}
                        v22 := []int{fi12}
                        if len(v22) > 0 {
                            v22[0] = fi12
                        }
                        var v23 float64 = 3.14159
                        _ = len(v22)
                        _ = v20
                        _ = v21
                        _ = v22
                        _ = v23
                    } else {
                        fmt.Println(fi11, fi11)
                    }
                    if v10 {
                        if false {
                            var v24 bool = (42 > 18446744073709551615)
                            _ = v24
                        } else {
                            v25 := []int{fi12, fi12, fi11}
                            switch fi11 {
                            case -9223372036854775808:
                                var v26 float64 = 3.14159
                                _ = v26
                            case 42:
                                if len(v25) > 0 {
                                    v25[0] = fi11
                                }
                            default:
                                v27 := GenS1{F0: 2147483647, F1: 1, F2: v10}
                                _ = v27
                            }
                            _ = v25
                        }
                    } else {
                        var v28 int = 32768
                        fmt.Println(fi11)
                        _ = v28
                    }
                    switch (fi11 + (fi11 * fi11)) {
                    case 2147483647:
                        v10 = v10
                    case -9223372036854775808:
                        if v10 {
                            v10 = (-2147483648 >= 339)
                            var v29 string = "hello"
                            v30 := GenS1{F0: 2147483647, F1: 0, F2: ((3.14159 - 1.5) <= ((3.14159 / float32((float64(3.14159) * ((((0.1 / (((1e10 - -2.25) / (3.14159 * (3.14159 * (((-2.25 / 1.5) * -2.25) - 1e10)))) + (3.14159 / float64(1.5)))) - 0.0) * -2.25) - 1e10)))) * (3.14159 * ((3.14159 + 1.5) * float32(0.1)))))}
                            _ = v29
                            _ = v30
                        }
                    }
                } else {
                    v10 = true
                    var v31 uint16 = 255
                    _ = v31
                }
            }
            v32 := GenS1{F0: -898, F1: 9223372036854775807, F2: v10}
            _ = fi12
            _ = v32
        }
        _ = fi11
    }
    for fi33 := 0; fi33 < 15; fi33++ {
        v10 = v10
        for fi34 := 0; fi34 < 3; fi34++ {
            fmt.Println(fi33)
            _ = fi34
        }
        _ = fi33
    }
    _ = v10
    return 1
}
func genFn1(p0 int32) uint16 {
    for fi35 := 0; fi35 < 2; fi35++ {
        p0 = (p0 << ((18446744073709551615) & 63))
        v36 := []int{(fi35 - -889), 32768, 9223372036854775807}
        _ = fi35
        _ = v36
    }
    v37 := GenS3{F0: 2}
    var v38 float64 = 3.14159
    fmt.Println(p0, v38)
    _ = p0
    _ = v37
    _ = v38
    return 255
}

func main() {
    var v1 int32 = 42
    v2 := make(map[string]string)
    v3 := make(map[uint32]uint32)
    var v4 string = ("" + ("" + "hello"))
    v5 := []uint32{256, 856, 550}
    v6 := []uint32{2147483648, 42, 2147483648, 65535}
    v1 = (v1 % (((v1 - v1)) | 1))
    if true {
        for fi39 := 0; fi39 < 10; fi39++ {
            if false {
                if ((^(fi39)) == fi39) {
                    v1 = v1
                }
                var v40 string = v4
                switch fi39 {
                case 1:
                    if true {
                        for fi41 := 0; fi41 < 15; fi41++ {
                            var v42 int16 = 32767
                            v1 -= v1
                            _ = fi41
                            _ = v42
                        }
                        switch (fi39 >> ((1) & 63)) {
                        case 276:
                            switch v40 {
                            case "a":
                                v43 := GenS2{F0: 194}
                                _ = v43
                            case "hello":
                                v40 = (v40 + ("\n\t" + v4))
                            case "foo bar":
                                var v44 bool = (1.5 != (1e10 * 0.1))
                                _ = v44
                            default:
                                var v45 int = 42
                                _ = v45
                            }
                        case 128:
                            var v46 bool = true
                            _ = v46
                        default:
                            switch fi39 {
                            case 809:
                                v4 = (v40 + v4)
                            }
                        }
                    } else {
                        v47 := GenS3{F0: 0}
                        _ = v47
                    }
                }
                _ = v40
            }
            if (4294967295 < 849) {
                var v48 uint32 = 65535
                switch (fi39 - 4294967295) {
                case 42:
                    v49 := GenS4{F0: ("" + v4)}
                    _ = v49
                case -335:
                    if (true && ((float64(0.0) * 3.14159) > (float64(0.1) - 0.1))) {
                        v3[(v48 + (v48 % ((956) | 1)))] = 42
                        var v50 int8 = -1
                        _ = v50
                    }
                case 127:
                    if (false && false) {
                        if true {
                            v48 = (v48 + v48)
                        } else {
                            for fi51 := 0; fi51 < 5; fi51++ {
                                var v52 uint = 255
                                switch v4 {
                                case "\n\t":
                                    fmt.Println(v4, v48, fi51)
                                default:
                                    v48 = v48
                                }
                                _ = fi51
                                _ = v52
                            }
                            var v53 float32 = 0.0
                            _ = v53
                        }
                        v48 += v48
                        var v54 uint64 = 42
                        _ = v54
                    } else {
                        if (!((!true) && true)) {
                            if true {
                                v4 = v4
                            } else {
                                switch 688 {
                                case -888:
                                    var v55 uint = 2
                                    _ = v55
                                default:
                                    v1 -= (-(v1))
                                }
                            }
                            var v56 string = ("\n\t" + v4)
                            _ = v56
                        } else {
                            for fi57 := 0; fi57 < 19; fi57++ {
                                var v58 uint16 = 329
                                if true {
                                    if len(v5) > 0 {
                                        v5[0] = v48
                                    }
                                    fmt.Println(v58, v48, v4)
                                    fmt.Println(v58)
                                    var v59 int = (fi39 ^ 1)
                                    _ = v59
                                }
                                v48 += 42
                                _ = fi57
                                _ = v58
                            }
                            v1 = (v1 / ((v1) | 1))
                        }
                        v1 = -2147483648
                    }
                default:
                    var v60 int16 = 42
                    _ = v60
                }
                v61 := GenS5{F0: 524, F1: 0}
                _ = v48
                _ = v61
            } else {
                v4 = ("a" + ((v4 + v4) + v4))
            }
            var v62 uint = 498
            _ = fi39
            _ = v62
        }
        var v63 int32 = (v1 >> ((4294967295) & 63))
        _ = v63
    } else {
        var v64 uint = 268
        var v65 string = v4
        for fi66 := 0; fi66 < 17; fi66++ {
            delete(v2, (v65 + (("" + v65) + "\n\t")))
            v67 := GenS6{F0: genFn0(), F1: 255, F2: v1}
            v5 = append(v5, 2147483648)
            _ = fi66
            _ = v67
        }
        _ = v64
        _ = v65
    }
    for fi68 := 0; fi68 < 16; fi68++ {
        v4 = "foo bar"
        _ = fi68
    }
    v4 = (v4 + "")
    if false {
        var v69 float32 = (1e10 - 3.14159)
        v4 = "foo bar"
        _ = v69
    }
    switch v4 {
    case "":
        if true {
            v1 = v1
            v4 = v4
            var v70 uint32 = 2
            _ = v70
        } else {
            v1 = v1
            if (18446744073709551615 <= 35) {
                v71 := GenS3{F0: 416}
                _ = v71
            }
            var v72 float32 = 1.5
            _ = v72
        }
    case "hello":
        for fi73 := 0; fi73 < 14; fi73++ {
            if (false || (((3.14159 - 0.1) - -2.25) != 3.14159)) {
                if false {
                    var v74 int64 = 1
                    _ = v74
                } else {
                    var v75 bool = false
                    if (-143 <= 0) {
                        var v76 int16 = 42
                        switch fi73 {
                        case 128:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        default:
                            switch ("a" + (v4 + v4)) {
                            case "hello":
                                var v77 string = v4
                                _ = v77
                            case "\n\t":
                                fmt.Println(v1, v4, v4)
                            default:
                                for fi78 := 0; fi78 < 19; fi78++ {
                                    var v79 int32 = (v1 / ((v1) | 1))
                                    var v80 uint16 = 65535
                                    _ = fi78
                                    _ = v79
                                    _ = v80
                                }
                            }
                        }
                        _ = v76
                    }
                    var v81 int32 = 0
                    _ = v75
                    _ = v81
                }
            } else {
                v1 = 40
                if (!(0 == 2)) {
                    switch ("\n\t" + "foo bar") {
                    case "":
                        var v82 int = (fi73 & (fi73 | (-(fi73))))
                        _ = v82
                    case "\n\t":
                        for fi83 := 0; fi83 < 15; fi83++ {
                            switch v4 {
                            case "a":
                                if len(v5) > 0 {
                                    v5[0] = 186
                                }
                                v6 = append(v6, 42)
                            case "\n\t":
                                if false {
                                    switch 661 {
                                    case -55:
                                        v1 -= (v1 | v1)
                                    case -650:
                                        v84 := GenS1{F0: v1, F1: 42, F2: (!(true || (true && true)))}
                                        _ = v84
                                    default:
                                        if ((true || (42 < 688)) || false) {
                                            if (4294967295 == 731) {
                                                v85 := GenS7{F0: 42}
                                                v5 = append(v5, 2)
                                                if true {
                                                    var v86 int8 = 29
                                                    v87 := GenS2{F0: 65535}
                                                    v88 := GenS1{F0: (v1 * 355), F1: 9223372036854775807, F2: ((1e10 < (float64((1e10 - (1.5 + -2.25))) - -2.25)) || false)}
                                                    _ = v86
                                                    _ = v87
                                                    _ = v88
                                                }
                                                _ = v85
                                            } else {
                                                if (16 > 0) {
                                                    switch "hello" {
                                                    case "\n\t":
                                                        if false {
                                                            v89 := GenS8{F0: 42, F1: 7}
                                                            switch "" {
                                                            case "hello":
                                                                v90 := v3[43]
                                                                _ = v90
                                                            case "foo bar":
                                                                v4 = v4
                                                            case "a":
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    fmt.Println(fi73)
                                                                default:
                                                                    v91 := GenS7{F0: 128}
                                                                    _ = v91
                                                                }
                                                            }
                                                            v1 -= (v1 << ((255) & 63))
                                                            _ = v89
                                                        }
                                                    case "hello":
                                                        if (!(!false)) {
                                                            if len(v5) > 0 {
                                                                v5[0] = 989
                                                            }
                                                        }
                                                    case "a":
                                                        v4 = (v4 + v4)
                                                    }
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                    v1 = -556
                                                }
                                                if true {
                                                    v4 = v4
                                                    v92 := GenS5{F0: -664, F1: 105}
                                                    if len(v5) > 0 {
                                                        v5[0] = 480
                                                    }
                                                    _ = v92
                                                } else {
                                                    switch v4 {
                                                    case "hello":
                                                        if true {
                                                            if (1 == -796) {
                                                                var v93 uint = 579
                                                                switch v4 {
                                                                case "a":
                                                                    var v94 uint = v93
                                                                    _ = v94
                                                                }
                                                                switch "" {
                                                                case "\n\t":
                                                                    v93 -= (^(v93))
                                                                case "a":
                                                                    if len(v6) > 0 {
                                                                        v6[0] = 42
                                                                    }
                                                                case "foo bar":
                                                                    switch -1 {
                                                                    case 4294967295:
                                                                        var v95 bool = false
                                                                        _ = v95
                                                                    }
                                                                }
                                                                v4 = (("\n\t" + v4) + (v4 + v4))
                                                                if true {
                                                                    if true {
                                                                        v93 = 777
                                                                    } else {
                                                                        if (!true) {
                                                                            var v96 int64 = 42
                                                                            if len(v6) >= 2 {
                                                                                v6 = v6[1:len(v6)]
                                                                            }
                                                                            var v97 int32 = (v1 - v1)
                                                                            var v98 bool = false
                                                                            var v99 int = int(v1)
                                                                            var v100 uint8 = 0
                                                                            v101 := GenS9{F0: (0 > 18446744073709551615), F1: 65535}
                                                                            _ = v100
                                                                            _ = v101
                                                                            _ = v96
                                                                            _ = v97
                                                                            _ = v98
                                                                            _ = v99
                                                                        }
                                                                    }
                                                                    v93 -= v93
                                                                    v1 = v1
                                                                } else {
                                                                    if ((fi73 + (-(fi73))) >= (fi83 % (((-(fi73))) | 1))) {
                                                                        var v102 int8 = 52
                                                                        var v103 string = v4
                                                                        _ = v102
                                                                        _ = v103
                                                                    }
                                                                    v93 += v93
                                                                }
                                                                _ = v93
                                                            }
                                                            var v104 int16 = 42
                                                            _ = v104
                                                        }
                                                    case "a":
                                                        var v105 uint32 = 2147483648
                                                        var v106 float64 = float64((0.1 + (0.1 / 3.14159)))
                                                        v1 = v1
                                                        v105 = (v105 % ((0) | 1))
                                                        v1 = (v1 + (v1 - v1))
                                                        delete(v2, (((v4 + "\n\t") + ((v4 + "hello") + "a")) + ("" + v4)))
                                                        _ = v105
                                                        _ = v106
                                                    case "\n\t":
                                                        v1 = v1
                                                    }
                                                    if (255 < 127) {
                                                        v4 = v4
                                                        v1 -= (v1 + (v1 % ((v1) | 1)))
                                                        var v107 string = ("\n\t" + v4)
                                                        _ = v107
                                                    } else {
                                                        var v108 float64 = (1.5 + (-2.25 - float64(0.0)))
                                                        if true {
                                                            var v109 string = v4
                                                            _ = v109
                                                        } else {
                                                            var v110 uint32 = 381
                                                            fmt.Println(v4)
                                                            switch "" {
                                                            case "hello":
                                                                v111 := GenS3{F0: (^(v110))}
                                                                _ = v111
                                                            default:
                                                                var v112 int16 = -32768
                                                                _ = v112
                                                            }
                                                            _ = v110
                                                        }
                                                        _ = v108
                                                    }
                                                }
                                                v4 = ("foo bar" + ((v4 + ("\n\t" + (((v4 + "a") + (v4 + v4)) + (v4 + (("foo bar" + ((v4 + v4) + v4)) + v4))))) + v4))
                                            }
                                        } else {
                                            var v113 int16 = 858
                                            v114 := GenS8{F0: 225, F1: 42}
                                            _ = v113
                                            _ = v114
                                        }
                                    }
                                    v4 = v4
                                }
                            case "":
                                v1 = v1
                            }
                            fmt.Println(v1, fi83, fi83)
                            fmt.Println(fi73, v1, v1)
                            _ = fi83
                        }
                    case "a":
                        switch fi73 {
                        case 0:
                            v115 := GenS4{F0: ((v4 + v4) + ("" + v4))}
                            _ = v115
                        }
                    }
                }
                var v116 bool = (32767 < 213)
                _ = v116
            }
            if len(v5) > 0 {
                v5[0] = 42
            }
            _ = fi73
        }
    default:
        var v117 uint16 = 701
        _ = v117
    }
    var v118 bool = (2 == 81)
    if false {
        v119 := GenS9{F0: ((1e10 / ((3.14159 - (float64(float32((0.0 + -2.25))) * float64(((0.1 + 1e10) * (0.1 + (0.1 + (0.0 - 1e10))))))) - ((0.0 - 1.5) / (((3.14159 - -2.25) + 0.1) + 1.5)))) > ((1e10 / 1.0) - ((-2.25 - (-2.25 * (1.5 * 1e10))) / 0.1))), F1: 4294967295}
        _ = v119
    }
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v118:", v118)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v118
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
