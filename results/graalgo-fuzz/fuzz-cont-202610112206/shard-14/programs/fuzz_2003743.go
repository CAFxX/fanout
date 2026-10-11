package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int32; F1 bool }
type GenS10 struct { F0 int8; F1 int }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int32; F1 uint64 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 bool; F1 int16; F2 int8 }
type GenS17 struct { F0 bool; F1 uint64 }
type GenS18 struct { F0 uint64 }
type GenS19 struct { F0 string; F1 uint }
type GenS20 struct { F0 int32; F1 int64 }
type GenS21 struct { F0 int; F1 int32; F2 uint32 }
type GenS22 struct { F0 uint8; F1 uint32; F2 string }
type GenS23 struct { F0 string; F1 int16 }
type GenS24 struct { F0 int8 }
type GenS25 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS26 struct { F0 uint32 }
type GenS27 struct { F0 int8; F1 uint32 }
type GenS28 struct { F0 uint32; F1 uint64 }
type GenS29 struct { F0 int; F1 int }
type GenS30 struct { F0 uint8; F1 int16 }
func genFn0() int64 {
    var v7 float64 = -2.25
    v7 = (v7 * v7)
    v7 = (float64(1.5) + 3.14159)
    v8 := GenS1{F0: 2}
    v7 += v7
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 int8) bool {
    v9 := []int{}
    v10 := GenS2{F0: -2147483648, F1: -797}
    switch "\n\t" {
    case "":
        if (18446744073709551615 >= 42) {
            v11 := GenS3{F0: 232, F1: 627}
            p0 = (p0 ^ (^(p0)))
            _ = v11
        } else {
            for fi12 := 0; fi12 < 2; fi12++ {
                if len(v9) > 0 {
                    v9[0] = (^(fi12))
                }
                var v13 int16 = -32768
                _ = fi12
                _ = v13
            }
            _ = len(v9)
        }
    case "a":
        switch 32768 {
        case 42:
            switch 128 {
            case -1:
                var v14 uint = 376
                _ = v14
            case -2147483648:
                p0 -= p0
            default:
                for fi15 := 0; fi15 < 3; fi15++ {
                    v16 := GenS1{F0: 0}
                    if len(v9) > 0 {
                        v9[0] = 2147483647
                    }
                    _ = fi15
                    _ = v16
                }
            }
        case 127:
            p0 = (p0 | p0)
        case 4294967295:
            p0 = 127
        }
    case "\n\t":
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 uint16 = 1
            _ = fi17
            _ = v18
        }
    default:
        v19 := GenS1{F0: 2147483648}
        _ = v19
    }
    _ = p0
    _ = v10
    _ = v9
    return (!true)
}

func main() {
    v1 := []string{""}
    var v2 int8 = -83
    var v3 string = "\n\t"
    var v4 int8 = (v2 ^ v2)
    var v5 uint8 = 42
    var v6 uint32 = 0
    var v20 int8 = -32
    v4 -= v2
    fmt.Println(v20, v5)
    var v21 int16 = 32767
    if (int(v6) != 399) {
        if (!(false && true)) {
            var v22 float32 = 0.1
            var v23 uint16 = 440
            for fi24 := 0; fi24 < 4; fi24++ {
                v20 += v2
                var v25 int8 = (v2 & (v4 | 0))
                _ = fi24
                _ = v25
            }
            _ = v22
            _ = v23
        } else {
            for fi26 := 0; fi26 < 17; fi26++ {
                for fi27 := 0; fi27 < 4; fi27++ {
                    v3 = (("" + ("foo bar" + v3)) + ("a" + v3))
                    v1 = append(v1, "foo bar")
                    v2 = (v4 % (((v4 ^ 127)) | 1))
                    var v28 int32 = 784
                    _ = fi27
                    _ = v28
                }
                _ = fi26
            }
            v6 -= v6
        }
    } else {
        var v29 int64 = 42
        if true {
            v6 = v6
        } else {
            for fi30 := 0; fi30 < 10; fi30++ {
                var v31 uint8 = v5
                v29 = (^(v29))
                _ = fi30
                _ = v31
            }
            switch 2147483647 {
            case -1:
                for fi32 := 0; fi32 < 19; fi32++ {
                    if false {
                        v33 := GenS2{F0: fi32, F1: (v29 & (v29 | (v29 >> ((65535) & 63))))}
                        v4 = v4
                        _ = v33
                    }
                    _ = fi32
                }
            case 1:
                if false {
                    if true {
                        switch -1 {
                        case 128:
                            var v34 int16 = v21
                            _ = v34
                        default:
                            v6 -= (^(v6))
                        }
                        for fi35 := 0; fi35 < 16; fi35++ {
                            var v36 uint16 = 32767
                            for fi37 := 0; fi37 < 2; fi37++ {
                                if false {
                                    if false {
                                        v6 = v6
                                    } else {
                                        v38 := GenS4{F0: 1}
                                        var v39 int64 = v29
                                        _ = len(v1)
                                        _ = v38
                                        _ = v39
                                    }
                                    fmt.Println(v4, v2, fi37)
                                } else {
                                    v40 := GenS4{F0: v4}
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = v40
                                }
                                v4 += v20
                                fmt.Println(v21)
                                _ = fi37
                            }
                            if true {
                                v3 = "\n\t"
                                v41 := GenS4{F0: v4}
                                switch ((v3 + v3) + ((v3 + "\n\t") + ("foo bar" + "a"))) {
                                case "":
                                    var v42 string = v3
                                    _ = v42
                                case "hello":
                                    v6 += (v6 / ((2) | 1))
                                case "foo bar":
                                    switch v3 {
                                    case "":
                                        if len(v1) > 0 {
                                            v1[0] = v3
                                        }
                                    case "a":
                                        v43 := GenS1{F0: v6}
                                        _ = v43
                                    case "foo bar":
                                        v44 := GenS5{F0: 1, F1: 0}
                                        _ = v44
                                    default:
                                        v6 -= (v6 | v6)
                                    }
                                }
                                _ = v41
                            } else {
                                var v45 bool = (true || (-384 > genFn0()))
                                _ = v45
                            }
                            _ = fi35
                            _ = v36
                        }
                        var v46 int16 = -294
                        _ = v46
                    } else {
                        var v47 uint64 = 42
                        var v48 int64 = -9223372036854775808
                        var v49 uint32 = (v6 + v6)
                        _ = v47
                        _ = v48
                        _ = v49
                    }
                }
            case 32767:
                v3 = v3
            default:
                if true {
                    if false {
                        for fi50 := 0; fi50 < 7; fi50++ {
                            var v51 string = ""
                            if (((v51 + (v51 + ("a" + v51))) >= "a") && false) {
                                var v52 uint16 = 65535
                                v53 := GenS5{F0: uint64(v5), F1: 42}
                                _ = v52
                                _ = v53
                            } else {
                                v3 = v51
                                var v54 uint = 1
                                _ = v54
                            }
                            _ = fi50
                            _ = v51
                        }
                    }
                    for fi55 := 0; fi55 < 12; fi55++ {
                        for fi56 := 0; fi56 < 20; fi56++ {
                            if len(v1) > 0 {
                                v1[0] = (v3 + v3)
                            }
                            v6 -= v6
                            var v57 uint64 = 698
                            v4 = v4
                            v58 := GenS1{F0: v6}
                            _ = fi56
                            _ = v57
                            _ = v58
                        }
                        _ = fi55
                    }
                    if true {
                        if false {
                            if len(v1) > 0 {
                                v1[0] = v3
                            }
                            for fi59 := 0; fi59 < 4; fi59++ {
                                var v60 uint64 = 594
                                _ = fi59
                                _ = v60
                            }
                            v61 := GenS2{F0: -406, F1: v29}
                            _ = v61
                        } else {
                            for fi62 := 0; fi62 < 20; fi62++ {
                                v20 = (v20 * v20)
                                _ = fi62
                            }
                        }
                        v63 := GenS6{F0: 4294967295}
                        v64 := GenS5{F0: uint64(v6), F1: -2147483648}
                        _ = v63
                        _ = v64
                    } else {
                        var v65 float64 = (0.0 * ((3.14159 + (-2.25 + 1.5)) * float64(((1.5 - 0.1) - 0.1))))
                        v3 = v3
                        _ = v65
                    }
                }
            }
            if ((v21 | v21) < v21) {
                v66 := GenS3{F0: 4294967295, F1: 831}
                for fi67 := 0; fi67 < 13; fi67++ {
                    if ((v3 > (v3 + v3)) && (true && (!true))) {
                        for fi68 := 0; fi68 < 9; fi68++ {
                            v69 := GenS3{F0: v6, F1: 65535}
                            _ = fi68
                            _ = v69
                        }
                        for fi70 := 0; fi70 < 10; fi70++ {
                            if true {
                                v71 := GenS1{F0: 2}
                                v72 := GenS6{F0: 42}
                                var v73 uint32 = 65535
                                _ = v71
                                _ = v72
                                _ = v73
                            }
                            if (genFn0() != v29) {
                                var v74 int = 32768
                                _ = v74
                            }
                            v75 := GenS3{F0: v6, F1: 0}
                            _ = fi70
                            _ = v75
                        }
                    }
                    if ((-2.25 * (0.0 + (0.0 * ((-2.25 + (float64((((1.5 - float32(0.1)) * ((0.0 / -2.25) + -2.25)) - 1.5)) + (((1.5 + (0.0 + 0.1)) - (0.0 - (float64(0.1) * 0.0))) * 3.14159))) / 1.5)))) == 1e10) {
                        for fi76 := 0; fi76 < 9; fi76++ {
                            fmt.Println(v29, v2)
                            var v77 float64 = 0.1
                            var v78 int8 = v2
                            _ = fi76
                            _ = v77
                            _ = v78
                        }
                        if (772 <= 844) {
                            var v79 int32 = 2147483647
                            _ = v79
                        } else {
                            var v80 int32 = 42
                            if true {
                                for fi81 := 0; fi81 < 20; fi81++ {
                                    v82 := GenS7{F0: 2147483647}
                                    var v83 float64 = 0.1
                                    v20 = (-(v4))
                                    v20 = (v4 + (^(v2)))
                                    switch (-(fi81)) {
                                    case -61:
                                        if true {
                                            if false {
                                                var v84 uint16 = 65535
                                                _ = v84
                                            }
                                            switch (fi81 << ((562) & 63)) {
                                            case 32768:
                                                if ((v3 != "foo bar") || true) {
                                                    if false {
                                                        switch (v3 + "a") {
                                                        case "hello":
                                                            var v85 string = (("foo bar" + (("a" + "\n\t") + v3)) + "\n\t")
                                                            _ = v85
                                                        default:
                                                            var v86 uint64 = 18446744073709551615
                                                            _ = v86
                                                        }
                                                        var v87 uint8 = 0
                                                        _ = v87
                                                    } else {
                                                        switch ((v3 + ((v3 + v3) + "a")) + (v3 + ((v3 + "hello") + ("" + v3)))) {
                                                        case "":
                                                            if false {
                                                                if (255 != 0) {
                                                                    switch v3 {
                                                                    case "a":
                                                                        v88 := GenS6{F0: 956}
                                                                        _ = v88
                                                                    default:
                                                                        var v89 uint = 2
                                                                        _ = v89
                                                                    }
                                                                }
                                                            }
                                                        case "a":
                                                            v21 -= v21
                                                        case "\n\t":
                                                            if false {
                                                                v80 -= v80
                                                                switch ("a" + "\n\t") {
                                                                case "hello":
                                                                    v90 := GenS6{F0: 926}
                                                                    _ = v90
                                                                }
                                                            } else {
                                                                if len(v1) > 0 {
                                                                    v1[0] = (v3 + "")
                                                                }
                                                                if ((v80 << ((18446744073709551615) & 63)) >= (v80 | v80)) {
                                                                    if (uint16(v6) == 32767) {
                                                                        v91 := GenS3{F0: 666, F1: 65535}
                                                                        fmt.Println(v21, v4, v5)
                                                                        fmt.Println(v21)
                                                                        _ = v91
                                                                    }
                                                                    v80 += v80
                                                                } else {
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    v21 += v21
                                                                    var v92 uint16 = 255
                                                                    _ = v92
                                                                }
                                                                if false {
                                                                    _ = len(v1)
                                                                    v80 -= -423
                                                                    fmt.Println(v83, fi81)
                                                                    v21 -= (v21 / (((v21 + 42)) | 1))
                                                                    v93 := GenS1{F0: 2}
                                                                    if ((v80 * -599) != (v80 << ((891) & 63))) {
                                                                        switch v3 {
                                                                        case "hello":
                                                                            if (false && true) {
                                                                                v1 = append(v1, "a")
                                                                                if false {
                                                                                    var v94 int64 = v29
                                                                                    if (true && ("\n\t" > v3)) {
                                                                                        if ((v5 + (v5 >> ((255) & 63))) != v5) {
                                                                                            v80 = (v80 % ((v80) | 1))
                                                                                            v83 -= (v83 / (((((3.14159 * 0.1) - v83) - v83) / ((v83 * 3.14159) * (v83 - (v83 * v83)))) + 1e10))
                                                                                            _ = len(v1)
                                                                                        }
                                                                                        v95 := GenS2{F0: (fi67 / ((fi81) | 1)), F1: v94}
                                                                                        _ = v95
                                                                                    } else {
                                                                                        _ = len(v1)
                                                                                        if (false && (v80 != v80)) {
                                                                                            if true {
                                                                                                v20 -= 100
                                                                                            } else {
                                                                                                var v96 int = (fi67 / (((^(fi67))) | 1))
                                                                                                _ = v96
                                                                                            }
                                                                                            v97 := GenS8{F0: -143}
                                                                                            _ = v97
                                                                                        }
                                                                                        var v98 float32 = -2.25
                                                                                        _ = v98
                                                                                    }
                                                                                    _ = v94
                                                                                } else {
                                                                                    var v99 int64 = genFn0()
                                                                                    var v100 bool = ((0.0 - 0.1) > (0.1 - ((-2.25 + (0.0 * (3.14159 + -2.25))) * float32(1.5))))
                                                                                    var v101 uint = 0
                                                                                    v2 = (-(v20))
                                                                                    _ = v100
                                                                                    _ = v101
                                                                                    _ = v99
                                                                                }
                                                                                if true {
                                                                                    if true {
                                                                                        v20 = 25
                                                                                        v3 = ((v3 + v3) + v3)
                                                                                        var v102 uint64 = 4294967295
                                                                                        var v103 int = -9223372036854775808
                                                                                        var v104 uint64 = (v102 * 612)
                                                                                        _ = v102
                                                                                        _ = v103
                                                                                        _ = v104
                                                                                    }
                                                                                } else {
                                                                                    var v105 uint8 = v5
                                                                                    v106 := GenS9{F0: (v80 ^ (v80 | -970)), F1: true}
                                                                                    _ = v105
                                                                                    _ = v106
                                                                                }
                                                                            } else {
                                                                                v107 := GenS3{F0: 655, F1: 0}
                                                                                v1 = append(v1, v3)
                                                                                var v108 uint8 = v5
                                                                                var v109 float64 = float64((0.1 * ((1.5 * (0.1 + 0.0)) + (-2.25 / 1.0))))
                                                                                _ = v107
                                                                                _ = v108
                                                                                _ = v109
                                                                            }
                                                                        default:
                                                                            switch v3 {
                                                                            case "\n\t":
                                                                                var v110 uint32 = (v6 >> ((42) & 63))
                                                                                _ = v110
                                                                            case "":
                                                                                var v111 string = (((("foo bar" + v3) + "foo bar") + (("" + "a") + "\n\t")) + v3)
                                                                                _ = v111
                                                                            case "hello":
                                                                                if false {
                                                                                    if false {
                                                                                        fmt.Println(v83)
                                                                                    } else {
                                                                                        switch -2147483648 {
                                                                                        case -65:
                                                                                            var v112 int64 = -1
                                                                                            _ = v112
                                                                                        }
                                                                                        var v113 bool = false
                                                                                        v6 = (v6 / ((v6) | 1))
                                                                                        var v114 int64 = v29
                                                                                        switch ("hello" + v3) {
                                                                                        case "hello":
                                                                                            switch v3 {
                                                                                            case "hello":
                                                                                                if true {
                                                                                                    var v115 int16 = v21
                                                                                                    v116 := GenS7{F0: -2147483648}
                                                                                                    if (739 > 42) {
                                                                                                        v117 := GenS6{F0: 430}
                                                                                                        _ = v117
                                                                                                    } else {
                                                                                                        var v118 float64 = v83
                                                                                                        var v119 bool = (v6 > v6)
                                                                                                        v20 += (v20 & v4)
                                                                                                        var v120 bool = (!(255 >= 42))
                                                                                                        _ = v118
                                                                                                        _ = v119
                                                                                                        _ = v120
                                                                                                    }
                                                                                                    _ = v115
                                                                                                    _ = v116
                                                                                                }
                                                                                            default:
                                                                                                if true {
                                                                                                    v21 += -32768
                                                                                                    var v121 uint64 = 0
                                                                                                    if (495 != 255) {
                                                                                                        if (!v113) {
                                                                                                            var v122 bool = (uint32(v5) > uint32(v121))
                                                                                                            var v123 int16 = v21
                                                                                                            _ = v122
                                                                                                            _ = v123
                                                                                                        }
                                                                                                    }
                                                                                                    var v124 int8 = v20
                                                                                                    _ = v121
                                                                                                    _ = v124
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        _ = v113
                                                                                        _ = v114
                                                                                    }
                                                                                    var v125 int8 = (v2 - int8(v6))
                                                                                    v126 := GenS10{F0: v2, F1: (fi81 + fi81)}
                                                                                    _ = v125
                                                                                    _ = v126
                                                                                }
                                                                            default:
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = (v3 + v3)
                                                                                }
                                                                            }
                                                                        }
                                                                        fmt.Println(v3, v80, v80)
                                                                    } else {
                                                                        var v127 float64 = v83
                                                                        fmt.Println(v83, v2)
                                                                        if len(v1) > 0 {
                                                                            v1[0] = "a"
                                                                        }
                                                                        _ = v127
                                                                    }
                                                                    var v128 int16 = v21
                                                                    v21 += -1
                                                                    _ = v128
                                                                    _ = v93
                                                                }
                                                            }
                                                        default:
                                                            v6 = v6
                                                        }
                                                        var v129 float32 = (-2.25 - 1.5)
                                                        fmt.Println(fi81)
                                                        _ = v129
                                                    }
                                                } else {
                                                    fmt.Println(v4)
                                                    if (false && true) {
                                                        fmt.Println(fi67)
                                                        v21 += -1
                                                        v130 := GenS11{F0: (v80 + (v80 - v80))}
                                                        _ = v130
                                                    } else {
                                                        v131 := GenS12{F0: (v80 & 743), F1: 42}
                                                        fmt.Println(v21, v83)
                                                        fmt.Println(v83)
                                                        _ = v131
                                                    }
                                                }
                                            case 4294967295:
                                                var v132 int32 = 2147483647
                                                _ = v132
                                            case 32767:
                                                v133 := GenS13{F0: false}
                                                _ = v133
                                            }
                                            if false {
                                                switch "foo bar" {
                                                case "hello":
                                                    fmt.Println(v3)
                                                default:
                                                    switch 32767 {
                                                    case 0:
                                                        switch fi67 {
                                                        case 711:
                                                            switch (v3 + ((v3 + "") + v3)) {
                                                            case "foo bar":
                                                                v21 -= int16(v6)
                                                            default:
                                                                v134 := GenS8{F0: (v21 << ((898) & 63))}
                                                                _ = v134
                                                            }
                                                        case 128:
                                                            v5 -= 42
                                                        case 2147483647:
                                                            var v135 uint16 = 0
                                                            _ = v135
                                                        default:
                                                            if (!false) {
                                                                if ((v29 / (((v29 ^ 0)) | 1)) < -1) {
                                                                    var v136 bool = true
                                                                    switch fi67 {
                                                                    case 574:
                                                                        if (3.14159 != (0.0 / 1.0)) {
                                                                            if len(v1) > 0 {
                                                                                v1[0] = v3
                                                                            }
                                                                            v1 = append(v1, v3)
                                                                            switch fi81 {
                                                                            case 9223372036854775807:
                                                                                switch "a" {
                                                                                case "foo bar":
                                                                                    var v137 bool = (v136 && v136)
                                                                                    _ = v137
                                                                                case "hello":
                                                                                    v6 -= 328
                                                                                default:
                                                                                    v138 := GenS14{F0: v3}
                                                                                    _ = v138
                                                                                }
                                                                            case -1:
                                                                                fmt.Println(v136, v80)
                                                                            default:
                                                                                _ = len(v1)
                                                                            }
                                                                        } else {
                                                                            if true {
                                                                                if (735 == 410) {
                                                                                    var v139 int16 = 275
                                                                                    _ = v139
                                                                                }
                                                                            }
                                                                            var v140 uint64 = 0
                                                                            _ = v140
                                                                        }
                                                                    default:
                                                                        fmt.Println(v4, v5, v80)
                                                                    }
                                                                    _ = v136
                                                                } else {
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    v3 = v3
                                                                    v83 -= (0.1 * ((v83 * (0.1 + 0.1)) - 1.5))
                                                                }
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                var v141 string = ((v3 + "") + (v3 + (v3 + v3)))
                                                                _ = v141
                                                            }
                                                            var v142 int32 = (v80 & (v80 & v80))
                                                            _ = v142
                                                        }
                                                    }
                                                    v29 = (v29 & (v29 ^ v29))
                                                }
                                                var v143 uint = 2
                                                _ = v143
                                            } else {
                                                var v144 bool = (true && (!((v3 + (v3 + v3)) != v3)))
                                                fmt.Println(v80, v80, v6)
                                                if (v80 >= -2147483648) {
                                                    v4 = v4
                                                    var v145 uint8 = v5
                                                    var v146 float64 = 1.5
                                                    switch "\n\t" {
                                                    case "a":
                                                        v80 -= 1
                                                    case "":
                                                        var v147 bool = (v3 <= v3)
                                                        _ = v147
                                                    case "\n\t":
                                                        if v144 {
                                                            var v148 int32 = -2147483648
                                                            _ = v148
                                                        } else {
                                                            v149 := GenS6{F0: 0}
                                                            fmt.Println(v146, v3, v80)
                                                            if v144 {
                                                                var v150 int8 = v4
                                                                _ = v150
                                                            }
                                                            if (v146 < (v83 * v83)) {
                                                                if ((-(v29)) != (v29 % ((v29) | 1))) {
                                                                    v145 = v145
                                                                    var v151 int8 = v20
                                                                    _ = v151
                                                                } else {
                                                                    var v152 int16 = v21
                                                                    switch "a" {
                                                                    case "a":
                                                                        if v144 {
                                                                            if v144 {
                                                                                var v153 int8 = v4
                                                                                if true {
                                                                                    _ = len(v1)
                                                                                    v5 -= (v145 + (v5 << ((4294967295) & 63)))
                                                                                } else {
                                                                                    var v154 uint64 = 907
                                                                                    switch ("" + (v3 + v3)) {
                                                                                    case "\n\t":
                                                                                        v155 := GenS4{F0: v2}
                                                                                        _ = v155
                                                                                    }
                                                                                    switch ((((v3 + v3) + "") + "\n\t") + (v3 + (("hello" + "") + v3))) {
                                                                                    case "foo bar":
                                                                                        v5 = (v5 | v5)
                                                                                    default:
                                                                                        v156 := GenS15{F0: (v80 | v80)}
                                                                                        _ = v156
                                                                                    }
                                                                                    _ = v154
                                                                                }
                                                                                _ = v153
                                                                            }
                                                                            if false {
                                                                                fmt.Println(v5, v6, v6)
                                                                                var v157 bool = ((v4 - v2) != (v2 >> ((2) & 63)))
                                                                                if ((^(v6)) > (v6 * 2)) {
                                                                                    v21 = int16(v2)
                                                                                    if len(v1) >= 2 {
                                                                                        v1 = v1[1:len(v1)]
                                                                                    }
                                                                                } else {
                                                                                    if true {
                                                                                        v6 = v6
                                                                                        v157 = (!(v3 < (v3 + v3)))
                                                                                        fmt.Println(v2)
                                                                                    } else {
                                                                                        var v158 int = (-(fi67))
                                                                                        _ = v158
                                                                                    }
                                                                                    var v159 int8 = v2
                                                                                    v29 = int64(v21)
                                                                                    _ = v159
                                                                                }
                                                                                _ = v157
                                                                            }
                                                                        } else {
                                                                            var v160 uint64 = 1
                                                                            if v144 {
                                                                                if true {
                                                                                    v160 = uint64(v6)
                                                                                    fmt.Println(v5)
                                                                                    v161 := GenS16{F0: v144, F1: v21, F2: v2}
                                                                                    if v144 {
                                                                                        fmt.Println(v29, v20)
                                                                                        v162 := GenS17{F0: false, F1: (v160 ^ 621)}
                                                                                        _ = v162
                                                                                    } else {
                                                                                        v163 := GenS17{F0: (((v3 + (((v3 + ((v3 + "\n\t") + "a")) + v3) + ("" + v3))) + ((v3 + v3) + v3)) >= (v3 + "hello")), F1: (v160 - v160)}
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = v3
                                                                                        }
                                                                                        _ = len(v1)
                                                                                        _ = v163
                                                                                    }
                                                                                    _ = v161
                                                                                } else {
                                                                                    if (-65 > -9223372036854775808) {
                                                                                        switch v3 {
                                                                                        case "hello":
                                                                                            var v164 int64 = v29
                                                                                            _ = v164
                                                                                        case "a":
                                                                                            var v165 int64 = int64(v160)
                                                                                            _ = v165
                                                                                        default:
                                                                                            var v166 int = (fi67 & 1)
                                                                                            _ = v166
                                                                                        }
                                                                                        fmt.Println(v83)
                                                                                        v5 -= v145
                                                                                    } else {
                                                                                        var v167 uint8 = 0
                                                                                        var v168 int8 = (-(v20))
                                                                                        fmt.Println(v83, v21)
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = ((v3 + ("" + v3)) + v3)
                                                                                        }
                                                                                        _ = v167
                                                                                        _ = v168
                                                                                    }
                                                                                    var v169 uint8 = 0
                                                                                    v170 := GenS2{F0: fi81, F1: v29}
                                                                                    _ = v169
                                                                                    _ = v170
                                                                                }
                                                                            }
                                                                            if v144 {
                                                                                var v171 uint = 255
                                                                                v172 := GenS7{F0: int32(v145)}
                                                                                v144 = (65535 < 227)
                                                                                var v173 uint64 = 541
                                                                                var v174 uint32 = v6
                                                                                _ = v171
                                                                                _ = v172
                                                                                _ = v173
                                                                                _ = v174
                                                                            }
                                                                            _ = v160
                                                                        }
                                                                    case "":
                                                                        v1 = append(v1, "foo bar")
                                                                    default:
                                                                        fmt.Println(fi67)
                                                                    }
                                                                    v144 = (!(fi81 <= fi81))
                                                                    _ = v152
                                                                }
                                                                v1 = append(v1, ("foo bar" + v3))
                                                                if v144 {
                                                                    var v175 int64 = int64(fi81)
                                                                    _ = v175
                                                                }
                                                                v176 := GenS12{F0: (v80 << ((18446744073709551615) & 63)), F1: 2}
                                                                v177 := GenS18{F0: 2}
                                                                fmt.Println(fi67)
                                                                _ = v176
                                                                _ = v177
                                                            }
                                                            _ = v149
                                                        }
                                                    default:
                                                        v178 := GenS7{F0: v80}
                                                        _ = v178
                                                    }
                                                    _ = v145
                                                    _ = v146
                                                }
                                                _ = v144
                                            }
                                        } else {
                                            v1 = append(v1, v3)
                                            var v179 float64 = (3.14159 / float64(float32(v83)))
                                            _ = v179
                                        }
                                    case 127:
                                        var v180 uint32 = (v6 / ((v6) | 1))
                                        _ = v180
                                    case 530:
                                        v181 := GenS19{F0: "foo bar", F1: 0}
                                        _ = v181
                                    default:
                                        var v182 bool = (!true)
                                        _ = v182
                                    }
                                    _ = fi81
                                    _ = v82
                                    _ = v83
                                }
                                fmt.Println(v80)
                                var v183 uint8 = (v5 | 2)
                                _ = v183
                            } else {
                                var v184 float32 = 1e10
                                _ = v184
                            }
                            v5 -= v5
                            _ = v80
                        }
                    }
                    _ = fi67
                }
                _ = v66
            } else {
                for fi185 := 0; fi185 < 8; fi185++ {
                    var v186 uint8 = (v5 % ((uint8(v6)) | 1))
                    if len(v1) > 0 {
                        v1[0] = v3
                    }
                    _ = fi185
                    _ = v186
                }
                v187 := GenS20{F0: 2147483647, F1: v29}
                for fi188 := 0; fi188 < 14; fi188++ {
                    if len(v1) > 0 {
                        v1[0] = (((v3 + v3) + "") + v3)
                    }
                    fmt.Println(v21, v2)
                    _ = fi188
                }
                _ = v187
            }
        }
        v189 := GenS10{F0: (v20 & v20), F1: 4294967295}
        _ = v189
        _ = v29
    }
    var v190 uint32 = v6
    for fi191 := 0; fi191 < 8; fi191++ {
        if (65535 == 32767) {
            v192 := GenS3{F0: (v6 % (((v6 % ((v6) | 1))) | 1)), F1: 288}
            v3 = (v3 + "")
            var v193 int32 = 693
            _ = v192
            _ = v193
        } else {
            var v194 int16 = (v21 & (v21 & v21))
            var v195 int16 = -674
            _ = v194
            _ = v195
        }
        var v196 uint = 1
        _ = fi191
        _ = v196
    }
    v197 := GenS11{F0: -697}
    fmt.Println(v2, v6, v2)
    if false {
        switch 245 {
        case -1:
            v198 := GenS21{F0: -129, F1: 1, F2: 794}
            _ = v198
        case 127:
            v6 += v6
        case 9223372036854775807:
            if (false && true) {
                _ = len(v1)
                v6 -= (v190 | 629)
            }
        default:
            if len(v1) > 0 {
                v1[0] = v3
            }
        }
    }
    if false {
        var v199 uint32 = 306
        if (243 != 0) {
            for fi200 := 0; fi200 < 16; fi200++ {
                v201 := GenS21{F0: (fi200 ^ 137), F1: 42, F2: 2147483648}
                _ = fi200
                _ = v201
            }
            v202 := GenS19{F0: ("\n\t" + v3), F1: 863}
            _ = v202
        }
        _ = v199
    }
    v203 := GenS7{F0: 263}
    for fi204 := 0; fi204 < 9; fi204++ {
        for fi205 := 0; fi205 < 10; fi205++ {
            var v206 bool = ((v3 < v3) && true)
            var v207 float32 = 3.14159
            if ((v206 && (((v5 << ((255) & 63)) == 128) || (42 <= 227))) && (v206 || (false || (730 == 42)))) {
                var v208 int = (fi204 / (((-(fi205))) | 1))
                if v206 {
                    if (0.0 == -2.25) {
                        if len(v1) > 0 {
                            v1[0] = ("foo bar" + (v3 + ((v3 + (v3 + ("\n\t" + v3))) + "hello")))
                        }
                        if len(v1) > 0 {
                            v1[0] = ("foo bar" + v3)
                        }
                        var v209 uint = 846
                        _ = v209
                    } else {
                        if true {
                            v210 := GenS18{F0: 493}
                            v5 -= v5
                            _ = v210
                        } else {
                            switch 4294967295 {
                            case -283:
                                fmt.Println(v208)
                            case 33:
                                var v211 uint64 = 514
                                _ = v211
                            default:
                                v212 := GenS22{F0: v5, F1: (v190 / ((1) | 1)), F2: v3}
                                _ = v212
                            }
                            fmt.Println(v4, v2, v3)
                            var v213 uint64 = 4294967295
                            _ = v213
                        }
                        var v214 int16 = (v21 % (((v21 | -32768)) | 1))
                        fmt.Println(v4, v20)
                        _ = v214
                    }
                    if ((v21 % (((v21 * (v21 + (v21 >> ((42) & 63))))) | 1)) < v21) {
                        switch "foo bar" {
                        case "foo bar":
                            v20 += int8(v21)
                        case "\n\t":
                            v2 = (v20 & v2)
                        case "hello":
                            switch (v3 + (("hello" + ("a" + v3)) + (v3 + v3))) {
                            case "\n\t":
                                switch "\n\t" {
                                case "foo bar":
                                    var v215 uint = uint(v6)
                                    _ = v215
                                case "hello":
                                    fmt.Println(v20)
                                case "":
                                    var v216 int8 = v2
                                    if v206 {
                                        v6 = v190
                                        v208 -= (fi204 >> ((422) & 63))
                                    } else {
                                        v3 = v3
                                        var v217 int32 = 2147483647
                                        switch "hello" {
                                        case "a":
                                            switch fi204 {
                                            case -8:
                                                switch (fi205 & fi204) {
                                                case -2147483648:
                                                    v2 = v216
                                                default:
                                                    var v218 int8 = v4
                                                    _ = v218
                                                }
                                            case 0:
                                                if true {
                                                    if (v206 || (!v206)) {
                                                        v217 += v217
                                                        var v219 int8 = (v2 & v4)
                                                        switch v3 {
                                                        case "foo bar":
                                                            var v220 int32 = (v217 + v217)
                                                            _ = v220
                                                        case "":
                                                            if false {
                                                                fmt.Println(v216, v216, fi205)
                                                                var v221 bool = v206
                                                                if true {
                                                                    var v222 float32 = (1.5 / ((v207 / (v207 + float32(1.5))) * v207))
                                                                    fmt.Println(v3, v206, v20)
                                                                    fmt.Println(fi205, v4)
                                                                    _ = v222
                                                                }
                                                                _ = v221
                                                            }
                                                        default:
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        }
                                                        _ = v219
                                                    } else {
                                                        v223 := GenS17{F0: v206, F1: 798}
                                                        var v224 bool = v206
                                                        v225 := GenS21{F0: fi205, F1: v217, F2: 1}
                                                        _ = v223
                                                        _ = v224
                                                        _ = v225
                                                    }
                                                    switch v3 {
                                                    case "":
                                                        v207 = v207
                                                    case "a":
                                                        if len(v1) > 0 {
                                                            v1[0] = (v3 + "foo bar")
                                                        }
                                                    case "hello":
                                                        v226 := GenS17{F0: v206, F1: 745}
                                                        _ = v226
                                                    }
                                                    v5 -= (v5 | (v5 + (^(v5))))
                                                    var v227 uint = 510
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = v227
                                                }
                                            case 127:
                                                switch fi204 {
                                                case -974:
                                                    switch (fi204 & (-(fi205))) {
                                                    case 1:
                                                        var v228 uint64 = 42
                                                        _ = v228
                                                    case 0:
                                                        var v229 int64 = 1
                                                        _ = v229
                                                    case -9223372036854775808:
                                                        v230 := GenS1{F0: 2147483647}
                                                        _ = v230
                                                    default:
                                                        fmt.Println(v208, v217, v20)
                                                        var v231 uint = 255
                                                        v208 = (fi204 << (((v231 << ((v231) & 63))) & 63))
                                                        _ = v231
                                                    }
                                                case -9223372036854775808:
                                                    if v206 {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        v232 := GenS23{F0: v3, F1: v21}
                                                        _ = v232
                                                    } else {
                                                        v207 -= v207
                                                    }
                                                case -498:
                                                    var v233 bool = true
                                                    _ = v233
                                                default:
                                                    if (!true) {
                                                        v234 := GenS9{F0: -2147483648, F1: (!v206)}
                                                        v217 -= -547
                                                        _ = v234
                                                    }
                                                }
                                            default:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                        case "":
                                            if v206 {
                                                fmt.Println(v5, v206, v216)
                                            }
                                        default:
                                            switch fi204 {
                                            case 4294967295:
                                                var v235 uint64 = 154
                                                _ = v235
                                            default:
                                                var v236 float64 = 1.5
                                                switch (v208 ^ fi204) {
                                                case 127:
                                                    switch (fi204 / ((fi204) | 1)) {
                                                    case -1:
                                                        var v237 uint8 = v5
                                                        _ = v237
                                                    case 298:
                                                        if v206 {
                                                            fmt.Println(v216, v216)
                                                            switch "" {
                                                            case "foo bar":
                                                                var v238 uint64 = uint64(v6)
                                                                _ = v238
                                                            case "":
                                                                v190 = (^(v190))
                                                            case "hello":
                                                                if (!(v208 != (fi205 << ((2) & 63)))) {
                                                                    v21 = v21
                                                                    var v239 uint32 = 44
                                                                    _ = v239
                                                                }
                                                            }
                                                        } else {
                                                            v1 = append(v1, "foo bar")
                                                            fmt.Println(v236, v236)
                                                            if len(v1) > 0 {
                                                                v1[0] = v3
                                                            }
                                                        }
                                                    case 127:
                                                        v4 = v20
                                                    }
                                                default:
                                                    v1 = append(v1, "a")
                                                }
                                                _ = v236
                                            }
                                        }
                                        _ = v217
                                    }
                                    _ = v216
                                default:
                                    var v240 float32 = v207
                                    _ = v240
                                }
                            default:
                                var v241 bool = true
                                _ = v241
                            }
                        default:
                            v190 = v6
                        }
                        if true {
                            v242 := GenS24{F0: 100}
                            if ((((1e10 * float32(0.1)) / (-2.25 / v207)) - 1.5) >= -2.25) {
                                v243 := GenS1{F0: (^(v190))}
                                if len(v1) > 0 {
                                    v1[0] = v3
                                }
                                var v244 int64 = -632
                                if (v6 != (v6 + v6)) {
                                    var v245 int = v208
                                    var v246 uint8 = (v5 << ((252) & 63))
                                    _ = v245
                                    _ = v246
                                } else {
                                    v247 := GenS4{F0: (v2 ^ 1)}
                                    var v248 int32 = -2147483648
                                    if len(v1) > 0 {
                                        v1[0] = v3
                                    }
                                    v249 := GenS20{F0: v248, F1: v244}
                                    if (true || ((genFn0() >= (v244 & (v244 / ((v244) | 1)))) || v206)) {
                                        var v250 int16 = (v21 ^ v21)
                                        v20 -= (v20 & (^(v20)))
                                        if ((v206 || false) || ((!((fi205 / ((32767) | 1)) == 9223372036854775807)) && v206)) {
                                            if len(v1) > 0 {
                                                v1[0] = v3
                                            }
                                        } else {
                                            v1 = append(v1, "foo bar")
                                            v1 = append(v1, v3)
                                            var v251 uint = 65535
                                            _ = v251
                                        }
                                        _ = v250
                                    }
                                    _ = v247
                                    _ = v248
                                    _ = v249
                                }
                                if true {
                                    if ((!(!((v207 < 1e10) || (fi204 == (fi204 % ((fi205) | 1)))))) || v206) {
                                        var v252 uint = 65535
                                        _ = v252
                                    } else {
                                        v253 := GenS25{F0: -863, F1: 0, F2: (v190 >> ((0) & 63))}
                                        fmt.Println(v4, v2, v2)
                                        _ = v253
                                    }
                                    var v254 bool = (v5 >= 255)
                                    _ = v254
                                } else {
                                    var v255 string = v3
                                    _ = len(v1)
                                    _ = v255
                                }
                                switch 0 {
                                case -9223372036854775808:
                                    v207 = (1e10 + 3.14159)
                                case 0:
                                    var v256 int8 = v20
                                    _ = v256
                                case 128:
                                    v257 := GenS10{F0: -41, F1: v208}
                                    _ = v257
                                }
                                switch 32768 {
                                case 523:
                                    var v258 int32 = 42
                                    _ = v258
                                case 32768:
                                    switch (v208 ^ fi204) {
                                    case 239:
                                        if ((-2.25 != (float32(0.0) / (v207 - v207))) || v206) {
                                            v3 = "hello"
                                        }
                                    default:
                                        switch -254 {
                                        case 9223372036854775807:
                                            var v259 uint32 = v6
                                            _ = v259
                                        case -9223372036854775808:
                                            v20 -= int8(v21)
                                        case 32768:
                                            fmt.Println(v20, fi204, v190)
                                        }
                                        v206 = v206
                                        var v260 uint64 = 18446744073709551615
                                        _ = v260
                                    }
                                default:
                                    v261 := GenS26{F0: (v6 >> ((183) & 63))}
                                    _ = v261
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v243
                                _ = v244
                            }
                            _ = v242
                        }
                    }
                }
                switch (-(fi205)) {
                case 0:
                    var v262 uint = 255
                    _ = v262
                case 127:
                    if v206 {
                        if ((v206 && (v20 < (v4 + v4))) || ((fi204 / ((fi204) | 1)) == int(v190))) {
                            var v263 uint = 427
                            var v264 uint8 = (^(v5))
                            if true {
                                var v265 int = -2147483648
                                var v266 int32 = 587
                                if len(v1) > 0 {
                                    v1[0] = v3
                                }
                                _ = v265
                                _ = v266
                            } else {
                                v206 = true
                                fmt.Println(fi204, v21)
                                v190 += v6
                            }
                            if (26 > 2147483648) {
                                fmt.Println(v2, v2)
                                v264 += v5
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            } else {
                                v267 := GenS27{F0: (v20 - v2), F1: v190}
                                if false {
                                    v3 = "hello"
                                    var v268 uint16 = 0
                                    _ = v268
                                } else {
                                    var v269 uint16 = 65535
                                    switch fi205 {
                                    case -1:
                                        if len(v1) > 0 {
                                            v1[0] = v3
                                        }
                                    case 32768:
                                        v270 := GenS28{F0: (v190 >> ((507) & 63)), F1: 0}
                                        _ = v270
                                    default:
                                        switch -15 {
                                        case 531:
                                            v271 := GenS29{F0: (v208 + (fi205 * fi204)), F1: fi205}
                                            _ = v271
                                        case 42:
                                            v206 = v206
                                        }
                                    }
                                    _ = v269
                                }
                                _ = v267
                            }
                            var v272 uint32 = (v6 >> (((v263 >> (((v263 & v263)) & 63))) & 63))
                            _ = v263
                            _ = v264
                            _ = v272
                        }
                    }
                default:
                    fmt.Println(v4, v6, v206)
                }
                _ = v208
            }
            _ = fi205
            _ = v206
            _ = v207
        }
        var v273 uint16 = 390
        var v274 bool = (!(42 < 18446744073709551615))
        _ = fi204
        _ = v273
        _ = v274
    }
    switch (v3 + v3) {
    case "a":
        if ((((float64((1e10 * 3.14159)) + (3.14159 + 1e10)) - 1e10) / -2.25) <= 0.1) {
            var v275 float32 = 1.5
            _ = v275
        }
    case "foo bar":
        _ = len(v1)
    case "hello":
        for fi276 := 0; fi276 < 14; fi276++ {
            if (fi276 == fi276) {
                v1 = append(v1, v3)
            } else {
                var v277 uint32 = (^(v190))
                for fi278 := 0; fi278 < 11; fi278++ {
                    var v279 int8 = v2
                    v280 := GenS30{F0: v5, F1: -1}
                    if (false && ((4294967295 != 938) || (!true))) {
                        _ = len(v1)
                        v3 = ""
                    } else {
                        var v281 uint16 = 0
                        if len(v1) > 0 {
                            v1[0] = (("" + (v3 + "\n\t")) + "\n\t")
                        }
                        _ = v281
                    }
                    _ = fi278
                    _ = v279
                    _ = v280
                }
                _ = v277
            }
            _ = fi276
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v190:", v190)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v190
    _ = v197
    _ = v2
    _ = v20
    _ = v203
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
