package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 bool; F1 int32; F2 int8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint16; F1 uint16 }
type GenS8 struct { F0 int; F1 int64 }
type GenS9 struct { F0 int64; F1 uint }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int16; F1 int16 }
type GenS12 struct { F0 string; F1 uint32; F2 int16 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 string; F1 uint16 }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 uint32 }
type GenS17 struct { F0 uint; F1 uint }
type GenS18 struct { F0 uint64 }
type GenS19 struct { F0 int; F1 int }
type GenS20 struct { F0 int8; F1 uint16 }
type GenS21 struct { F0 uint64; F1 uint }
type GenS22 struct { F0 int16; F1 int32; F2 bool }
type GenS23 struct { F0 uint; F1 uint }
type GenS24 struct { F0 int32; F1 uint; F2 int64 }
type GenS25 struct { F0 uint }
type GenS26 struct { F0 string; F1 uint64; F2 uint }
type GenS27 struct { F0 uint64; F1 uint64; F2 int64 }
type GenS28 struct { F0 int8; F1 int64 }
type GenS29 struct { F0 uint }
type GenS30 struct { F0 bool; F1 int16; F2 uint16 }
type GenS31 struct { F0 string; F1 int; F2 uint32 }
type GenS32 struct { F0 uint32 }
func genFn0(p0 int, p1 uint64, p2 bool) int64 {
    v6 := GenS1{F0: 42}
    v7 := []int{p0, (p0 >> ((2) & 63)), (p0 + (p0 & p0)), p0}
    for fi8 := 0; fi8 < 17; fi8++ {
        p1 -= p1
        if (-2.25 != float64((-2.25 - 0.0))) {
            var v9 uint = 1
            if true {
                fmt.Println(p1, p1, p0)
            }
            switch "\n\t" {
            case "foo bar":
                switch "" {
                case "\n\t":
                    v10 := GenS1{F0: 1}
                    _ = v10
                case "foo bar":
                    if len(v7) > 0 {
                        v7[0] = 1
                    }
                case "hello":
                    p1 = p1
                }
            case "a":
                var v11 int32 = 386
                _ = v11
            case "\n\t":
                switch "foo bar" {
                case "a":
                    v12 := GenS1{F0: 220}
                    _ = v12
                }
            default:
                p1 -= 42
            }
            _ = v9
        } else {
            v13 := GenS1{F0: 2}
            v7 = append(v7, (fi8 << ((154) & 63)))
            p2 = p2
            _ = v13
        }
        _ = fi8
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        if (!(("" == ("foo bar" + ("foo bar" + ""))) || (-1 <= 1000))) {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            if len(v7) > 0 {
                v7[0] = 42
            }
            var v15 uint8 = 1
            _ = v15
        } else {
            p2 = ((p0 + (-(fi14))) > -1)
        }
        var v16 float64 = 0.1
        if p2 {
            var v17 uint16 = 1
            _ = v17
        }
        _ = fi14
        _ = v16
    }
    p2 = true
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 240
}
func genFn1() uint {
    if false {
        v18 := GenS1{F0: 42}
        var v19 float64 = 1e10
        _ = v18
        _ = v19
    } else {
        switch (("a" + ("foo bar" + "foo bar")) + ("\n\t" + "hello")) {
        case "":
            switch ("foo bar" + ("hello" + "a")) {
            case "a":
                if true {
                    switch "foo bar" {
                    case "\n\t":
                        var v20 uint16 = 255
                        _ = v20
                    case "foo bar":
                        v21 := []int{128, 4294967295, -2147483648, 1}
                        _ = v21
                    case "hello":
                        var v22 uint8 = 42
                        _ = v22
                    default:
                        var v23 uint = 329
                        _ = v23
                    }
                    var v24 int32 = -156
                    _ = v24
                } else {
                    if true {
                        for fi25 := 0; fi25 < 10; fi25++ {
                            var v26 bool = (412 <= 2147483647)
                            _ = fi25
                            _ = v26
                        }
                        v27 := []int{9223372036854775807, 527, 0}
                        for fi28 := 0; fi28 < 1; fi28++ {
                            var v29 int8 = 127
                            v29 -= (v29 % ((3) | 1))
                            var v30 uint64 = 42
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                        _ = v27
                    }
                    for fi31 := 0; fi31 < 18; fi31++ {
                        var v32 int64 = 0
                        var v33 float64 = 1.5
                        var v34 int16 = -32768
                        _ = fi31
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                    v35 := GenS2{F0: ("\n\t" + "foo bar")}
                    _ = v35
                }
            case "hello":
                var v36 uint32 = 65535
                _ = v36
            }
        case "foo bar":
            var v37 bool = (false && false)
            _ = v37
        }
        v38 := GenS3{F0: 42}
        var v39 uint8 = 255
        _ = v38
        _ = v39
    }
    var v40 uint32 = 0
    for fi41 := 0; fi41 < 4; fi41++ {
        v42 := []int{42}
        _ = fi41
        _ = v42
    }
    fmt.Println(v40)
    _ = v40
    return 255
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{-1, -974, -83, -2147483648}
    v3 := make(map[int]string)
    v4 := make(map[int]bool)
    v5 := []uint32{944, 470, 1}
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        var v43 uint16 = 0
        _ = v43
    default:
        var v44 uint16 = 940
        _ = v44
    }
    var v45 uint64 = 42
    for fi46 := 0; fi46 < 2; fi46++ {
        for fi47 := 0; fi47 < 19; fi47++ {
            fmt.Println(fi47, fi46)
            v45 -= v45
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi47
        }
        var v48 uint32 = 2147483648
        if (((fi46 >> ((78) & 63)) >= fi46) || ((0.0 - 3.14159) != 3.14159)) {
            fmt.Println(v45, v48, fi46)
            v45 = 42
        }
        _ = fi46
        _ = v48
    }
    v49 := GenS3{F0: v45}
    v45 = (v45 / (((v45 << ((1) & 63))) | 1))
    v45 = 4294967295
    var v50 float64 = -2.25
    if false {
        switch ("foo bar" + "\n\t") {
        case "a":
            if len(v2) > 0 {
                v2[0] = -1
            }
        case "foo bar":
            v50 = v50
        case "":
            for fi51 := 0; fi51 < 19; fi51++ {
                var v52 float64 = (v50 - float64(-2.25))
                _ = fi51
                _ = v52
            }
        default:
            var v53 string = "a"
            _ = v53
        }
    } else {
        var v54 int = 0
        _ = v54
    }
    _ = len(v2)
    if false {
        if len(v5) > 0 {
            v5[0] = 65535
        }
    }
    var v55 uint16 = 2
    v45 -= v45
    if (!(100 == 0)) {
        for fi56 := 0; fi56 < 10; fi56++ {
            for fi57 := 0; fi57 < 3; fi57++ {
                v58 := GenS2{F0: "a"}
                var v59 uint32 = 1
                var v60 uint = 18446744073709551615
                _ = fi57
                _ = v58
                _ = v59
                _ = v60
            }
            _ = fi56
        }
        fmt.Println(v50, v50, v50)
    } else {
        if len(v2) > 0 {
            v2[0] = -674
        }
        v61 := GenS4{F0: 9223372036854775807, F1: -2147483648}
        v5 = append(v5, 2147483647)
        _ = v61
    }
    for fi62 := 0; fi62 < 12; fi62++ {
        v55 = v55
        switch ("a" + "") {
        case "a":
            if false {
                if (int32(v45) != 539) {
                    if len(v2) > 0 {
                        v2[0] = fi62
                    }
                    for fi63 := 0; fi63 < 5; fi63++ {
                        var v64 bool = false
                        switch "" {
                        case "hello":
                            _ = len(v5)
                        default:
                            if (4294967295 == 281) {
                                var v65 int32 = 42
                                v66 := v3[fi62]
                                v64 = ((v50 * (v50 + ((1.5 + v50) + 1.5))) <= v50)
                                _ = v65
                                _ = v66
                            } else {
                                var v67 int = 42
                                v50 += ((1e10 / v50) + v50)
                                if v64 {
                                    _ = len(v5)
                                    switch ("" + "a") {
                                    case "foo bar":
                                        v68 := GenS1{F0: 186}
                                        _ = v68
                                    case "":
                                        fmt.Println(v55, v67, v64)
                                    case "hello":
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    default:
                                        v45 += (v45 >> ((genFn1()) & 63))
                                    }
                                    var v69 string = "hello"
                                    _ = v69
                                } else {
                                    var v70 uint8 = 128
                                    v2 = append(v2, 32768)
                                    if false {
                                        var v71 float32 = (float32((((((v50 * float64(-2.25)) / float64((-2.25 - 1e10))) / v50) / -2.25) / (float64(((3.14159 * 0.0) * -2.25)) / v50))) - -2.25)
                                        v71 = v71
                                        if false {
                                            switch "foo bar" {
                                            case "":
                                                v72 := GenS5{F0: (!(!(0 < -30))), F1: -1, F2: 0}
                                                _ = v72
                                            case "\n\t":
                                                var v73 uint64 = v45
                                                _ = v73
                                            case "foo bar":
                                                var v74 uint = 65535
                                                _ = v74
                                            default:
                                                fmt.Println(v64)
                                            }
                                        } else {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            if false {
                                                var v75 uint32 = 1
                                                _ = v75
                                            } else {
                                                var v76 uint16 = 0
                                                v64 = v64
                                                _ = v76
                                            }
                                            v77 := GenS5{F0: (!v64), F1: 0, F2: 100}
                                            _ = v77
                                        }
                                        if v64 {
                                            _ = len(v2)
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = (^(v67))
                                        }
                                        _ = v71
                                    } else {
                                        var v78 uint8 = v70
                                        if (275 > int16(v70)) {
                                            switch v67 {
                                            case 42:
                                                v79 := GenS6{F0: v55}
                                                _ = v79
                                            case 9223372036854775807:
                                                if (v64 && v64) {
                                                    if true {
                                                        var v80 uint8 = (v78 / ((141) | 1))
                                                        switch (v67 % ((-9223372036854775808) | 1)) {
                                                        case -1:
                                                            v64 = (-122 <= -42)
                                                        case 1:
                                                            var v81 float32 = ((-2.25 / 0.1) / float32(v50))
                                                            _ = v81
                                                        default:
                                                            v50 = ((-2.25 + v50) - float64(3.14159))
                                                        }
                                                        if v64 {
                                                            var v82 int32 = -1
                                                            v2 = append(v2, -9223372036854775808)
                                                            fmt.Println(fi62)
                                                            _ = v82
                                                        } else {
                                                            delete(v3, fi63)
                                                        }
                                                        var v83 uint64 = 18446744073709551615
                                                        _ = v80
                                                        _ = v83
                                                    }
                                                    if (582 > 9223372036854775807) {
                                                        var v84 bool = (-2147483648 <= 0)
                                                        fmt.Println(v55, v45, fi62)
                                                        _ = v84
                                                    } else {
                                                        v50 = v50
                                                        if v64 {
                                                            _ = len(v5)
                                                        }
                                                    }
                                                    if v64 {
                                                        if v64 {
                                                            v78 -= v78
                                                            fmt.Println(fi63, v70)
                                                        }
                                                        var v85 uint16 = (^(v55))
                                                        if v64 {
                                                            fmt.Println(fi62, v67)
                                                            v3[int(v85)] = "a"
                                                            v67 = fi62
                                                        } else {
                                                            v86 := GenS7{F0: v85, F1: v55}
                                                            v45 = v45
                                                            v64 = v64
                                                            _ = v86
                                                        }
                                                        _ = v85
                                                    }
                                                } else {
                                                    switch "" {
                                                    case "foo bar":
                                                        v50 -= float64((1.5 / 1.5))
                                                    case "hello":
                                                        var v87 uint16 = (v55 % ((v55) | 1))
                                                        _ = v87
                                                    case "a":
                                                        var v88 int8 = 13
                                                        _ = v88
                                                    default:
                                                        _ = len(v1)
                                                    }
                                                    var v89 uint64 = (v45 % (((v45 | 2)) | 1))
                                                    var v90 int64 = 9223372036854775807
                                                    if (!(602 < 112)) {
                                                        var v91 uint8 = v70
                                                        if (42 == int32(v70)) {
                                                            if ((v90 - (v90 % ((42) | 1))) != 9223372036854775807) {
                                                                v50 = float64((1.5 + (((0.0 * (1.5 - (float32(3.14159) / (-2.25 + 1e10)))) - (((-2.25 * (1.5 - ((0.1 - -2.25) * -2.25))) + 0.1) - 0.1)) - float32((((((0.0 - ((1e10 / (v50 + 0.0)) * v50)) / float64((0.1 + 0.0))) / (0.1 * v50)) - v50) - float64(1e10))))))
                                                            }
                                                            v92 := GenS1{F0: v78}
                                                            _ = v92
                                                        } else {
                                                            v70 += v78
                                                        }
                                                        v93 := GenS6{F0: (v55 + v55)}
                                                        switch (v67 ^ -798) {
                                                        case 611:
                                                            v78 = (v78 >> ((2) & 63))
                                                        }
                                                        var v94 int = 181
                                                        _ = v91
                                                        _ = v93
                                                        _ = v94
                                                    } else {
                                                        var v95 int32 = 1
                                                        var v96 uint32 = 42
                                                        if ((v78 * v70) <= (v78 ^ (v70 >> ((1) & 63)))) {
                                                            var v97 uint32 = 447
                                                            _ = len(v3)
                                                            _ = v97
                                                        } else {
                                                            switch "foo bar" {
                                                            case "hello":
                                                                if (!v64) {
                                                                    v96 = v96
                                                                    _ = len(v1)
                                                                    switch ("" + "a") {
                                                                    case "hello":
                                                                        var v98 string = ("hello" + "")
                                                                        _ = v98
                                                                    case "foo bar":
                                                                        if true {
                                                                            v3[fi63] = ""
                                                                            if v64 {
                                                                                var v99 float32 = (0.0 / float32(v50))
                                                                                if v64 {
                                                                                    fmt.Println(v99, v50)
                                                                                }
                                                                                _ = v99
                                                                            } else {
                                                                                fmt.Println(v50)
                                                                                if ((v95 ^ v95) > -472) {
                                                                                    v45 += (v45 * v89)
                                                                                }
                                                                                switch ("" + "a") {
                                                                                case "\n\t":
                                                                                    if false {
                                                                                        v100 := GenS6{F0: (v55 ^ 65535)}
                                                                                        if len(v2) >= 2 {
                                                                                            v2 = v2[1:len(v2)]
                                                                                        }
                                                                                        var v101 float64 = v50
                                                                                        _ = v100
                                                                                        _ = v101
                                                                                    } else {
                                                                                        var v102 string = ("hello" + "hello")
                                                                                        if v64 {
                                                                                            v95 = int32(fi63)
                                                                                        }
                                                                                        v67 -= (fi63 - 128)
                                                                                        v89 += v45
                                                                                        var v103 uint = 248
                                                                                        v104 := GenS8{F0: (fi63 / ((127) | 1)), F1: int64(v96)}
                                                                                        v105 := v1[(^(v96))]
                                                                                        _ = v102
                                                                                        _ = v103
                                                                                        _ = v104
                                                                                        _ = v105
                                                                                    }
                                                                                default:
                                                                                    var v106 uint8 = uint8(v55)
                                                                                    _ = v106
                                                                                }
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v95)
                                                                            v64 = (v89 == (v45 >> ((42) & 63)))
                                                                            var v107 int32 = v95
                                                                            _ = v107
                                                                        }
                                                                    case "\n\t":
                                                                        v108 := GenS7{F0: uint16(v96), F1: (v55 / (((v55 - (^(v55)))) | 1))}
                                                                        _ = v108
                                                                    default:
                                                                        v109 := GenS4{F0: (v90 & (v90 & v90)), F1: (^(v67))}
                                                                        _ = v109
                                                                    }
                                                                } else {
                                                                    var v110 uint = 774
                                                                    var v111 string = "foo bar"
                                                                    _ = v110
                                                                    _ = v111
                                                                }
                                                            default:
                                                                v67 += 127
                                                            }
                                                            switch int(v70) {
                                                            case -1:
                                                                switch 2147483647 {
                                                                case 53:
                                                                    v112 := GenS2{F0: "foo bar"}
                                                                    var v113 uint32 = v96
                                                                    _ = v112
                                                                    _ = v113
                                                                case 128:
                                                                    var v114 int64 = v90
                                                                    _ = v114
                                                                }
                                                            case 2147483647:
                                                                v78 = 128
                                                            case 203:
                                                                v115 := GenS9{F0: (v90 * (v90 | (v90 | (^(v90))))), F1: 4294967295}
                                                                _ = v115
                                                            default:
                                                                var v116 int64 = (-(v90))
                                                                _ = v116
                                                            }
                                                            var v117 int = v67
                                                            _ = v117
                                                        }
                                                        switch (fi63 | 127) {
                                                        case -619:
                                                            _ = len(v5)
                                                        case 32768:
                                                            var v118 uint = genFn1()
                                                            _ = v118
                                                        default:
                                                            var v119 int8 = -128
                                                            _ = v119
                                                        }
                                                        _ = v95
                                                        _ = v96
                                                    }
                                                    v90 += 1
                                                    v120 := GenS8{F0: fi63, F1: genFn0((fi63 << ((0) & 63)), v89, v64)}
                                                    v67 -= (fi62 / ((fi63) | 1))
                                                    v121 := v3[fi63]
                                                    if (66 >= -128) {
                                                        v122 := GenS3{F0: (v89 * 76)}
                                                        var v123 uint16 = v55
                                                        _ = v122
                                                        _ = v123
                                                    }
                                                    _ = len(v3)
                                                    _ = v120
                                                    _ = v121
                                                    _ = v89
                                                    _ = v90
                                                }
                                            case 333:
                                                var v124 uint8 = (v70 | 129)
                                                _ = v124
                                            }
                                        }
                                        if v64 {
                                            v125 := GenS6{F0: 2}
                                            switch (v67 | (^(fi63))) {
                                            case -877:
                                                v45 = (v45 - (v45 >> ((genFn1()) & 63)))
                                            case -1:
                                                v70 -= (v70 ^ (v70 % (((^(v70))) | 1)))
                                            default:
                                                _ = len(v2)
                                            }
                                            _ = v125
                                        }
                                        switch "" {
                                        case "a":
                                            var v126 int64 = 0
                                            _ = v126
                                        case "foo bar":
                                            fmt.Println(fi63, fi63)
                                        }
                                        if v64 {
                                            if (-2.25 >= (1e10 + ((0.0 + (float32(v50) - 1.5)) + 0.1))) {
                                                v70 += (v78 % (((^(v78))) | 1))
                                            }
                                        } else {
                                            var v127 int = (fi63 * -1)
                                            v78 = v78
                                            _ = v127
                                        }
                                        v64 = v64
                                        _ = v78
                                    }
                                    _ = v70
                                }
                                fmt.Println(v67)
                                v67 += 0
                                _ = v67
                            }
                        }
                        _ = fi63
                        _ = v64
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    var v128 uint32 = 2147483648
                    var v129 float64 = float64(float32(float64(0.0)))
                    var v130 float64 = 0.0
                    _ = v128
                    _ = v129
                    _ = v130
                }
                fmt.Println(v55)
                v131 := GenS5{F0: true, F1: -1, F2: 1}
                _ = v131
            }
        case "":
            v55 = 1
        case "hello":
            var v132 uint8 = 1
            _ = v132
        default:
            v50 = (v50 - -2.25)
        }
        _ = fi62
    }
    _ = len(v4)
    for fi133 := 0; fi133 < 17; fi133++ {
        if (((!true) && true) && ((-926 != 1000) || true)) {
            v55 = (v55 | v55)
            if true {
                var v134 string = "hello"
                _ = v134
            }
        } else {
            var v135 uint32 = 2147483648
            for fi136 := 0; fi136 < 13; fi136++ {
                var v137 uint16 = 2
                _ = fi136
                _ = v137
            }
            if ((v45 - (v45 ^ v45)) != (v45 % ((18446744073709551615) | 1))) {
                if true {
                    switch (fi133 % ((fi133) | 1)) {
                    case 1:
                        var v138 float32 = (-2.25 + (-2.25 / -2.25))
                        _ = v138
                    }
                    switch (-(fi133)) {
                    case 805:
                        for fi139 := 0; fi139 < 14; fi139++ {
                            v45 += (v45 * uint64(v55))
                            _ = fi139
                        }
                    case 128:
                        switch (fi133 ^ fi133) {
                        case 1:
                            v55 += (^(v55))
                        case 32767:
                            v140 := GenS10{F0: (v55 ^ 32767)}
                            _ = v140
                        case 128:
                            for fi141 := 0; fi141 < 4; fi141++ {
                                fmt.Println(fi133, v135, v135)
                                _ = fi141
                            }
                        }
                    }
                } else {
                    if (true && ((-2.25 * 1e10) > 0.0)) {
                        switch -355 {
                        case -133:
                            v142 := v3[fi133]
                            _ = v142
                        case 451:
                            v135 = 677
                        case -341:
                            var v143 uint8 = 2
                            _ = v143
                        }
                        v2 = append(v2, (fi133 / ((fi133) | 1)))
                        fmt.Println(v135, v55)
                    } else {
                        if (88 <= 42) {
                            switch -2147483648 {
                            case -1:
                                var v144 int8 = -128
                                _ = v144
                            }
                        }
                        if false {
                            if (!true) {
                                switch ("\n\t" + "a") {
                                case "\n\t":
                                    if (-1 == -1) {
                                        var v145 uint64 = (v45 >> ((1) & 63))
                                        _ = v145
                                    }
                                case "hello":
                                    if false {
                                        _ = len(v4)
                                        for fi146 := 0; fi146 < 18; fi146++ {
                                            if false {
                                                fmt.Println(v55, fi146, v135)
                                                fmt.Println(v50, v45, v45)
                                                v135 += 166
                                            } else {
                                                v55 = 2
                                                v147 := GenS4{F0: 42, F1: (fi133 & 0)}
                                                _ = v147
                                            }
                                            _ = fi146
                                        }
                                    } else {
                                        switch "foo bar" {
                                        case "hello":
                                            if true {
                                                v135 = 42
                                                for fi148 := 0; fi148 < 2; fi148++ {
                                                    var v149 float32 = 0.1
                                                    v150 := GenS3{F0: (^(v45))}
                                                    v151 := GenS11{F0: 354, F1: -1}
                                                    _ = fi148
                                                    _ = v149
                                                    _ = v150
                                                    _ = v151
                                                }
                                                var v152 uint = 0
                                                _ = v152
                                            }
                                        default:
                                            if false {
                                                for fi153 := 0; fi153 < 17; fi153++ {
                                                    if (true && ((v45 >> ((uint(v135)) & 63)) >= (v45 + 536))) {
                                                        v154 := GenS10{F0: (v55 % ((v55) | 1))}
                                                        var v155 uint8 = 0
                                                        switch 1 {
                                                        case 4294967295:
                                                            v55 = (v55 * uint16(v45))
                                                        case -407:
                                                            if true {
                                                                var v156 uint16 = 845
                                                                _ = v156
                                                            } else {
                                                                switch "\n\t" {
                                                                case "":
                                                                    fmt.Println(v155, v45, fi153)
                                                                case "\n\t":
                                                                    fmt.Println(v135, fi133, fi153)
                                                                    if false {
                                                                        v55 += (v55 | v55)
                                                                        switch "" {
                                                                        case "hello":
                                                                            v135 -= (v135 % (((^(v135))) | 1))
                                                                        case "foo bar":
                                                                            fmt.Println(fi153, fi153, v155)
                                                                        case "a":
                                                                            var v157 float64 = v50
                                                                            _ = v157
                                                                        }
                                                                    } else {
                                                                        v158 := v1[192]
                                                                        switch v158 {
                                                                        case "":
                                                                            switch (fi153 + fi153) {
                                                                            case -417:
                                                                                v55 = (^(v55))
                                                                            case 127:
                                                                                v159 := GenS12{F0: v158, F1: (v135 + v135), F2: 936}
                                                                                if len(v5) > 0 {
                                                                                    v5[0] = (v135 ^ 2147483647)
                                                                                }
                                                                                if (533 < fi133) {
                                                                                    v160 := GenS7{F0: (v55 - (v55 >> ((276) & 63))), F1: 2}
                                                                                    if ((v155 ^ v155) < 42) {
                                                                                        switch ((("hello" + v158) + (v158 + "hello")) + v158) {
                                                                                        case "":
                                                                                            v161 := GenS13{F0: v135}
                                                                                            _ = v161
                                                                                        }
                                                                                        v3[fi153] = ""
                                                                                    } else {
                                                                                        fmt.Println(v45, v135, fi133)
                                                                                    }
                                                                                    v55 += (v55 / ((v55) | 1))
                                                                                    _ = v160
                                                                                }
                                                                                _ = v159
                                                                            case 356:
                                                                                var v162 int = (-(fi133))
                                                                                _ = v162
                                                                            }
                                                                        case "\n\t":
                                                                            if (!((v135 << ((255) & 63)) > v135)) {
                                                                                delete(v4, (fi153 ^ fi133))
                                                                                delete(v3, (fi133 & fi133))
                                                                                v163 := GenS14{F0: "hello", F1: (v55 - v55)}
                                                                                v164 := GenS15{F0: 127}
                                                                                var v165 string = ("\n\t" + ("" + v158))
                                                                                switch v158 {
                                                                                case "foo bar":
                                                                                    if true {
                                                                                        switch int(v45) {
                                                                                        case -867:
                                                                                            v166 := GenS6{F0: 255}
                                                                                            _ = v166
                                                                                        case -9223372036854775808:
                                                                                            var v167 int32 = 2147483647
                                                                                            _ = v167
                                                                                        case 1:
                                                                                            var v168 uint8 = (v155 - 0)
                                                                                            _ = v168
                                                                                        }
                                                                                        var v169 int8 = 100
                                                                                        var v170 string = (v165 + "hello")
                                                                                        _ = v169
                                                                                        _ = v170
                                                                                    } else {
                                                                                        v155 = (v155 * 128)
                                                                                        switch ("\n\t" + "") {
                                                                                        case "foo bar":
                                                                                            var v171 bool = (v155 != 1)
                                                                                            var v172 int64 = -1
                                                                                            v45 += (v45 ^ v45)
                                                                                            var v173 int = fi133
                                                                                            v172 += int64(fi133)
                                                                                            var v174 string = (v165 + "")
                                                                                            _ = v171
                                                                                            _ = v172
                                                                                            _ = v173
                                                                                            _ = v174
                                                                                        case "a":
                                                                                            switch int(v155) {
                                                                                            case 127:
                                                                                                v1[210] = (("" + v158) + "\n\t")
                                                                                            }
                                                                                        default:
                                                                                            var v175 int = 981
                                                                                            _ = v175
                                                                                        }
                                                                                        var v176 int64 = -9223372036854775808
                                                                                        _ = v176
                                                                                    }
                                                                                case "\n\t":
                                                                                    if true {
                                                                                        switch (fi153 / ((fi133) | 1)) {
                                                                                        case -1:
                                                                                            v158 = (v165 + ("hello" + (("a" + v165) + v158)))
                                                                                            v55 = v55
                                                                                            if true {
                                                                                                v177 := GenS7{F0: v55, F1: v55}
                                                                                                var v178 uint8 = (v155 ^ (v155 * v155))
                                                                                                _ = v177
                                                                                                _ = v178
                                                                                            }
                                                                                            var v179 bool = false
                                                                                            _ = v179
                                                                                        case 1:
                                                                                            var v180 int32 = 1
                                                                                            _ = v180
                                                                                        case 32768:
                                                                                            fmt.Println(v165)
                                                                                        default:
                                                                                            if ((!false) || (!((!false) && true))) {
                                                                                                fmt.Println(v135, v50, v50)
                                                                                                v181 := GenS6{F0: v55}
                                                                                                _ = v181
                                                                                            } else {
                                                                                                v182 := GenS16{F0: v135}
                                                                                                var v183 int64 = genFn0(fi133, 4294967295, (!false))
                                                                                                v184 := GenS17{F0: 2, F1: 42}
                                                                                                _ = v182
                                                                                                _ = v183
                                                                                                _ = v184
                                                                                            }
                                                                                        }
                                                                                        switch (v158 + v165) {
                                                                                        case "\n\t":
                                                                                            var v185 bool = false
                                                                                            _ = v185
                                                                                        case "":
                                                                                            fmt.Println(fi133, v165, v158)
                                                                                            fmt.Println(fi133, v45)
                                                                                            v158 = "a"
                                                                                        case "a":
                                                                                            var v186 uint8 = uint8(v45)
                                                                                            _ = v186
                                                                                        }
                                                                                        switch 0 {
                                                                                        case 4294967295:
                                                                                            v165 = v165
                                                                                        case 9223372036854775807:
                                                                                            delete(v3, (^(fi153)))
                                                                                        case 0:
                                                                                            v50 -= (1.5 + (((float64(float32((3.14159 / 1.0))) * -2.25) - float64(3.14159)) / float64((float32(-2.25) + 1e10))))
                                                                                        }
                                                                                        if true {
                                                                                            v2 = append(v2, fi133)
                                                                                            if true {
                                                                                                v135 = 2147483648
                                                                                                v187 := GenS8{F0: fi153, F1: -1}
                                                                                                _ = v187
                                                                                            } else {
                                                                                                var v188 bool = false
                                                                                                _ = v188
                                                                                            }
                                                                                            v189 := GenS17{F0: 841, F1: genFn1()}
                                                                                            _ = v189
                                                                                        } else {
                                                                                            v190 := GenS18{F0: uint64(v155)}
                                                                                            _ = v190
                                                                                        }
                                                                                        if len(v2) > 0 {
                                                                                            v2[0] = (^(fi133))
                                                                                        }
                                                                                    } else {
                                                                                        fmt.Println(v135, v155)
                                                                                        v191 := GenS1{F0: 250}
                                                                                        _ = v191
                                                                                    }
                                                                                default:
                                                                                    if true {
                                                                                        v165 = v158
                                                                                        v192 := GenS3{F0: v45}
                                                                                        _ = v192
                                                                                    } else {
                                                                                        switch (v158 + "\n\t") {
                                                                                        case "a":
                                                                                            var v193 uint = 42
                                                                                            _ = v193
                                                                                        case "hello":
                                                                                            fmt.Println(v135)
                                                                                        default:
                                                                                            switch (fi133 | 32768) {
                                                                                            case 4294967295:
                                                                                                fmt.Println(fi153, v158)
                                                                                            }
                                                                                            v55 = 1
                                                                                        }
                                                                                        if len(v5) > 0 {
                                                                                            v5[0] = (v135 % ((v135) | 1))
                                                                                        }
                                                                                        fmt.Println(v158, fi133)
                                                                                        if len(v2) > 0 {
                                                                                            v2[0] = -2147483648
                                                                                        }
                                                                                        if false {
                                                                                            v1[(v135 | v135)] = ("" + v158)
                                                                                            if false {
                                                                                                if (v158 == v158) {
                                                                                                    var v194 int16 = 1
                                                                                                    _ = v194
                                                                                                } else {
                                                                                                    if false {
                                                                                                        v55 = uint16(v155)
                                                                                                    } else {
                                                                                                        v195 := GenS19{F0: fi153, F1: (^(fi133))}
                                                                                                        if true {
                                                                                                            v158 = (v158 + v165)
                                                                                                            switch "\n\t" {
                                                                                                            case "foo bar":
                                                                                                                _ = len(v4)
                                                                                                            }
                                                                                                            if (!(v55 >= (v55 | 713))) {
                                                                                                                switch fi133 {
                                                                                                                case 128:
                                                                                                                    if true {
                                                                                                                        if false {
                                                                                                                            v55 += (v55 + v55)
                                                                                                                            v135 = uint32(v55)
                                                                                                                            v55 -= (v55 / ((v55) | 1))
                                                                                                                        } else {
                                                                                                                            v196 := GenS20{F0: -79, F1: (v55 - (^(v55)))}
                                                                                                                            _ = v196
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        if false {
                                                                                                                            if false {
                                                                                                                                switch (fi133 & (fi133 / (((fi153 << ((0) & 63))) | 1))) {
                                                                                                                                case -1:
                                                                                                                                    var v197 int16 = 0
                                                                                                                                    _ = v197
                                                                                                                                case 127:
                                                                                                                                    v55 += v55
                                                                                                                                case 4294967295:
                                                                                                                                    var v198 string = (v165 + ("a" + (((v158 + v158) + v165) + (v165 + v165))))
                                                                                                                                    _ = v198
                                                                                                                                default:
                                                                                                                                    var v199 int64 = -331
                                                                                                                                    _ = v199
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                v200 := v1[v135]
                                                                                                                                _ = v200
                                                                                                                            }
                                                                                                                            v165 = (v158 + v165)
                                                                                                                        } else {
                                                                                                                            fmt.Println(fi133, v155)
                                                                                                                        }
                                                                                                                        var v201 uint = uint(v135)
                                                                                                                        v202 := GenS16{F0: 439}
                                                                                                                        _ = v201
                                                                                                                        _ = v202
                                                                                                                    }
                                                                                                                default:
                                                                                                                    if false {
                                                                                                                        var v203 int64 = 871
                                                                                                                        fmt.Println(v55)
                                                                                                                        _ = v203
                                                                                                                    }
                                                                                                                }
                                                                                                                v55 = 277
                                                                                                                _ = len(v4)
                                                                                                            } else {
                                                                                                                v204 := GenS10{F0: (v55 % (((^(v55))) | 1))}
                                                                                                                if (!(100 > 100)) {
                                                                                                                    v165 = "hello"
                                                                                                                }
                                                                                                                _ = v204
                                                                                                            }
                                                                                                        }
                                                                                                        switch (v165 + ("" + "a")) {
                                                                                                        case "hello":
                                                                                                            if true {
                                                                                                                v205 := v1[(v135 - v135)]
                                                                                                                var v206 int64 = -9223372036854775808
                                                                                                                if false {
                                                                                                                    v207 := GenS21{F0: 448, F1: 65}
                                                                                                                    if false {
                                                                                                                        v206 -= (v206 ^ genFn0(fi133, v45, false))
                                                                                                                    }
                                                                                                                    _ = v207
                                                                                                                }
                                                                                                                v158 = v158
                                                                                                                _ = v205
                                                                                                                _ = v206
                                                                                                            }
                                                                                                        case "\n\t":
                                                                                                            v208 := GenS22{F0: -295, F1: -2147483648, F2: (float32(3.14159) != ((((float32((-2.25 / v50)) * (-2.25 - float32(v50))) / ((0.1 + (-2.25 * (0.1 / (1e10 - -2.25)))) / (0.1 * float32(1e10)))) / (1.5 * 1.5)) / 1.5))}
                                                                                                            _ = v208
                                                                                                        case "":
                                                                                                            var v209 uint64 = 4294967295
                                                                                                            _ = v209
                                                                                                        }
                                                                                                        _ = v195
                                                                                                    }
                                                                                                    if ((v135 % ((v135) | 1)) > (v135 + v135)) {
                                                                                                        fmt.Println(fi153, v50)
                                                                                                    }
                                                                                                }
                                                                                                if len(v2) >= 2 {
                                                                                                    v2 = v2[1:len(v2)]
                                                                                                }
                                                                                                if false {
                                                                                                    v165 = "hello"
                                                                                                    if true {
                                                                                                        var v210 float32 = 1e10
                                                                                                        _ = v210
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            var v211 int = 42
                                                                                            if false {
                                                                                                v212 := v1[65535]
                                                                                                var v213 uint8 = (v155 >> ((4294967295) & 63))
                                                                                                _ = v212
                                                                                                _ = v213
                                                                                            }
                                                                                            v45 = v45
                                                                                            var v214 int16 = 60
                                                                                            v5 = append(v5, (v135 / (((v135 - 65535)) | 1)))
                                                                                            _ = v211
                                                                                            _ = v214
                                                                                        }
                                                                                        switch fi153 {
                                                                                        case 128:
                                                                                            v135 = v135
                                                                                        case 0:
                                                                                            if ("hello" > v165) {
                                                                                                if true {
                                                                                                    v50 += v50
                                                                                                    var v215 int8 = -34
                                                                                                    _ = v215
                                                                                                } else {
                                                                                                    if false {
                                                                                                        switch ((("" + v158) + v158) + (v165 + "\n\t")) {
                                                                                                        case "\n\t":
                                                                                                            switch v165 {
                                                                                                            case "\n\t":
                                                                                                                v216 := GenS2{F0: v165}
                                                                                                                _ = v216
                                                                                                            case "":
                                                                                                                var v217 float32 = 0.1
                                                                                                                _ = v217
                                                                                                            case "hello":
                                                                                                                v55 += (v55 % ((255) | 1))
                                                                                                            }
                                                                                                        case "":
                                                                                                            switch v165 {
                                                                                                            case "hello":
                                                                                                                if false {
                                                                                                                    v135 += 2147483648
                                                                                                                    var v218 uint16 = (v55 & 255)
                                                                                                                    fmt.Println(v55, v165)
                                                                                                                    v50 -= 0.1
                                                                                                                    _ = v218
                                                                                                                } else {
                                                                                                                    var v219 uint8 = (v155 + (v155 / ((v155) | 1)))
                                                                                                                    _ = v219
                                                                                                                }
                                                                                                                if (!true) {
                                                                                                                    v220 := GenS23{F0: uint(v135), F1: 2}
                                                                                                                    fmt.Println(v55)
                                                                                                                    v50 = (v50 - v50)
                                                                                                                    _ = v220
                                                                                                                } else {
                                                                                                                    v4[1] = false
                                                                                                                }
                                                                                                            case "foo bar":
                                                                                                                v155 = (v155 << ((376) & 63))
                                                                                                            default:
                                                                                                                fmt.Println(v55, v165)
                                                                                                            }
                                                                                                            v2 = append(v2, -632)
                                                                                                            switch "foo bar" {
                                                                                                            case "\n\t":
                                                                                                                if true {
                                                                                                                    fmt.Println(v50)
                                                                                                                    v55 -= 950
                                                                                                                    var v221 uint64 = (v45 >> ((1) & 63))
                                                                                                                    _ = v221
                                                                                                                } else {
                                                                                                                    var v222 uint16 = v55
                                                                                                                    var v223 uint32 = 0
                                                                                                                    _ = v222
                                                                                                                    _ = v223
                                                                                                                }
                                                                                                            case "a":
                                                                                                                var v224 int32 = 363
                                                                                                                _ = v224
                                                                                                            default:
                                                                                                                var v225 uint32 = (v135 % (((v135 * (v135 - (v135 ^ v135)))) | 1))
                                                                                                                _ = v225
                                                                                                            }
                                                                                                        case "a":
                                                                                                            fmt.Println(v155, v55)
                                                                                                        }
                                                                                                    }
                                                                                                    var v226 string = v158
                                                                                                    _ = v226
                                                                                                }
                                                                                            } else {
                                                                                                _ = len(v3)
                                                                                            }
                                                                                        default:
                                                                                            var v227 float32 = 0.1
                                                                                            _ = v227
                                                                                        }
                                                                                    }
                                                                                }
                                                                                _ = v163
                                                                                _ = v164
                                                                                _ = v165
                                                                            } else {
                                                                                if (1 != -1) {
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = fi133
                                                                                    }
                                                                                    var v228 int16 = 1000
                                                                                    _ = v228
                                                                                } else {
                                                                                    fmt.Println(v50, fi133, v155)
                                                                                    switch (v158 + ("a" + (v158 + ""))) {
                                                                                    case "a":
                                                                                        v158 = v158
                                                                                    case "foo bar":
                                                                                        v229 := GenS5{F0: (0 < 1), F1: -487, F2: 76}
                                                                                        _ = v229
                                                                                    case "\n\t":
                                                                                        var v230 uint64 = (v45 / ((v45) | 1))
                                                                                        _ = v230
                                                                                    }
                                                                                    v155 -= 127
                                                                                }
                                                                                if true {
                                                                                    v231 := GenS24{F0: -269, F1: 18446744073709551615, F2: int64(v45)}
                                                                                    if true {
                                                                                        v135 = uint32(v45)
                                                                                        v55 -= uint16(v135)
                                                                                        switch fi153 {
                                                                                        case -733:
                                                                                            var v232 uint64 = (v45 >> ((18446744073709551615) & 63))
                                                                                            if true {
                                                                                                var v233 uint16 = (^(v55))
                                                                                                if (false || (42 > 1)) {
                                                                                                    v234 := GenS22{F0: -32768, F1: 1, F2: true}
                                                                                                    _ = v234
                                                                                                }
                                                                                                var v235 uint16 = (v233 | 255)
                                                                                                _ = v233
                                                                                                _ = v235
                                                                                            } else {
                                                                                                var v236 int16 = 32767
                                                                                                _ = v236
                                                                                            }
                                                                                            v45 = v232
                                                                                            _ = v232
                                                                                        }
                                                                                    }
                                                                                    var v237 uint = genFn1()
                                                                                    _ = v231
                                                                                    _ = v237
                                                                                } else {
                                                                                    var v238 uint32 = (v135 - v135)
                                                                                    _ = v238
                                                                                }
                                                                                var v239 int32 = -1
                                                                                _ = v239
                                                                            }
                                                                        case "foo bar":
                                                                            var v240 bool = true
                                                                            _ = v240
                                                                        default:
                                                                            fmt.Println(v158, v155)
                                                                        }
                                                                        _ = v158
                                                                    }
                                                                    if true {
                                                                        switch fi133 {
                                                                        case 0:
                                                                            var v241 float32 = 3.14159
                                                                            v242 := GenS6{F0: v55}
                                                                            var v243 int32 = 2147483647
                                                                            _ = v241
                                                                            _ = v242
                                                                            _ = v243
                                                                        default:
                                                                            var v244 string = "a"
                                                                            _ = v244
                                                                        }
                                                                        switch fi133 {
                                                                        case 9223372036854775807:
                                                                            switch "\n\t" {
                                                                            case "hello":
                                                                                v245 := GenS5{F0: true, F1: -1, F2: -43}
                                                                                _ = v245
                                                                            case "\n\t":
                                                                                v246 := GenS13{F0: (^(v135))}
                                                                                _ = v246
                                                                            case "foo bar":
                                                                                v55 = (v55 & v55)
                                                                            default:
                                                                                v247 := GenS25{F0: 18446744073709551615}
                                                                                _ = v247
                                                                            }
                                                                        }
                                                                        v55 = v55
                                                                    } else {
                                                                        fmt.Println(v45, v50, v45)
                                                                        v248 := GenS26{F0: "a", F1: (^(v45)), F2: 18446744073709551615}
                                                                        _ = v248
                                                                    }
                                                                }
                                                                v45 = (^(v45))
                                                            }
                                                        case -2147483648:
                                                            switch 4294967295 {
                                                            case 680:
                                                                fmt.Println(v55)
                                                            case 4294967295:
                                                                v50 = ((0.1 + (1e10 * float64(1e10))) - v50)
                                                            default:
                                                                if (-645 != 2147483647) {
                                                                    switch "hello" {
                                                                    case "":
                                                                        fmt.Println(fi153, v45)
                                                                    case "\n\t":
                                                                        v45 = 2
                                                                        if true {
                                                                            v249 := GenS27{F0: v45, F1: 1, F2: 887}
                                                                            _ = v249
                                                                        }
                                                                        if true {
                                                                            v50 = v50
                                                                            if (!true) {
                                                                                var v250 string = "foo bar"
                                                                                _ = v250
                                                                            }
                                                                        }
                                                                        var v251 uint8 = (^(v155))
                                                                        _ = v251
                                                                    case "hello":
                                                                        if (((1e10 / 3.14159) * 1e10) != 1.5) {
                                                                            var v252 int = (fi153 >> ((65535) & 63))
                                                                            fmt.Println(v155, v50, fi133)
                                                                            _ = v252
                                                                        }
                                                                    default:
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                    }
                                                                } else {
                                                                    var v253 float32 = 0.1
                                                                    v254 := GenS28{F0: 42, F1: 42}
                                                                    _ = v253
                                                                    _ = v254
                                                                }
                                                            }
                                                        default:
                                                            if true {
                                                                var v255 string = ""
                                                                var v256 int16 = 32767
                                                                if true {
                                                                    var v257 uint64 = 2
                                                                    _ = v257
                                                                } else {
                                                                    switch "" {
                                                                    case "a":
                                                                        _ = len(v5)
                                                                    case "foo bar":
                                                                        _ = len(v5)
                                                                    case "\n\t":
                                                                        v255 = "foo bar"
                                                                    default:
                                                                        v135 += (^(v135))
                                                                    }
                                                                    switch fi133 {
                                                                    case 4294967295:
                                                                        var v258 uint16 = v55
                                                                        _ = v258
                                                                    default:
                                                                        switch fi133 {
                                                                        case 622:
                                                                            v259 := GenS4{F0: -1, F1: -2147483648}
                                                                            _ = v259
                                                                        case -740:
                                                                            v260 := GenS21{F0: 966, F1: 740}
                                                                            _ = v260
                                                                        default:
                                                                            fmt.Println(v135)
                                                                        }
                                                                    }
                                                                    v261 := GenS28{F0: 17, F1: 566}
                                                                    var v262 float32 = -2.25
                                                                    _ = v261
                                                                    _ = v262
                                                                }
                                                                _ = v255
                                                                _ = v256
                                                            } else {
                                                                if ((-955 > 0) && true) {
                                                                    v135 -= v135
                                                                    if ((true && (!(!(599 == 1)))) && (-88 <= 1000)) {
                                                                        var v263 float64 = -2.25
                                                                        if true {
                                                                            var v264 uint8 = v155
                                                                            _ = v264
                                                                        }
                                                                        _ = v263
                                                                    } else {
                                                                        var v265 int64 = 42
                                                                        _ = v265
                                                                    }
                                                                    var v266 uint16 = (v55 - (^(v55)))
                                                                    _ = v266
                                                                } else {
                                                                    delete(v3, (fi133 >> ((2) & 63)))
                                                                    v135 = (^(v135))
                                                                    var v267 int16 = -32768
                                                                    _ = v267
                                                                }
                                                                if len(v5) > 0 {
                                                                    v5[0] = v135
                                                                }
                                                                v268 := GenS18{F0: (v45 / (((v45 + (v45 >> ((uint(v155)) & 63)))) | 1))}
                                                                switch "" {
                                                                case "":
                                                                    fmt.Println(fi133)
                                                                case "a":
                                                                    var v269 int16 = -1
                                                                    _ = v269
                                                                case "foo bar":
                                                                    _ = len(v3)
                                                                default:
                                                                    v270 := GenS11{F0: 1000, F1: -307}
                                                                    _ = v270
                                                                }
                                                                v271 := GenS29{F0: 4294967295}
                                                                _ = v268
                                                                _ = v271
                                                            }
                                                        }
                                                        var v272 bool = false
                                                        v50 = (v50 * 0.0)
                                                        _ = v154
                                                        _ = v155
                                                        _ = v272
                                                    }
                                                    v273 := GenS10{F0: 65535}
                                                    _ = fi153
                                                    _ = v273
                                                }
                                                if ("hello" > (("hello" + "foo bar") + "")) {
                                                    fmt.Println(fi133)
                                                    v5 = append(v5, v135)
                                                    var v274 uint32 = (v135 + (v135 & 42))
                                                    _ = v274
                                                }
                                            }
                                        }
                                        if true {
                                            var v275 uint16 = (v55 / (((v55 | (^(v55)))) | 1))
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            _ = v275
                                        } else {
                                            if false {
                                                v55 -= v55
                                            } else {
                                                var v276 int32 = 16
                                                for fi277 := 0; fi277 < 18; fi277++ {
                                                    if false {
                                                        if true {
                                                            delete(v4, (fi133 % ((fi277) | 1)))
                                                            _ = len(v4)
                                                            delete(v1, (^(v135)))
                                                            v50 += (v50 + v50)
                                                            v45 -= v45
                                                        } else {
                                                            v45 -= 18446744073709551615
                                                            if ("hello" > "") {
                                                                var v278 int16 = 79
                                                                var v279 uint64 = v45
                                                                var v280 int64 = 11
                                                                _ = v278
                                                                _ = v279
                                                                _ = v280
                                                            } else {
                                                                v5 = append(v5, (v135 ^ 2147483647))
                                                            }
                                                            v55 -= (v55 - v55)
                                                            switch ("a" + (("foo bar" + "") + "a")) {
                                                            case "foo bar":
                                                                _ = len(v4)
                                                            case "\n\t":
                                                                v55 = (v55 + v55)
                                                            }
                                                            var v281 string = "a"
                                                            _ = v281
                                                        }
                                                        v45 = 4294967295
                                                    }
                                                    v50 -= 0.0
                                                    _ = fi277
                                                }
                                                _ = v276
                                            }
                                            if (-957 != 933) {
                                                for fi282 := 0; fi282 < 8; fi282++ {
                                                    fmt.Println(v135, v50)
                                                    _ = fi282
                                                }
                                            } else {
                                                delete(v1, v135)
                                            }
                                        }
                                    }
                                }
                                for fi283 := 0; fi283 < 5; fi283++ {
                                    v45 = (^(v45))
                                    _ = fi283
                                }
                                v284 := GenS30{F0: true, F1: -72, F2: (v55 % ((v55) | 1))}
                                _ = v284
                            }
                            v135 = v135
                            if false {
                                if (true || (-533 < 42)) {
                                    v285 := GenS13{F0: v135}
                                    var v286 bool = false
                                    v50 += v50
                                    _ = v285
                                    _ = v286
                                } else {
                                    var v287 string = ("a" + ("a" + "\n\t"))
                                    for fi288 := 0; fi288 < 7; fi288++ {
                                        var v289 int = 1
                                        _ = fi288
                                        _ = v289
                                    }
                                    _ = v287
                                }
                            }
                        } else {
                            var v290 uint32 = (^(v135))
                            var v291 uint64 = 377
                            _ = v290
                            _ = v291
                        }
                        v55 = v55
                    }
                    var v292 int64 = 1
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v292
                }
                var v293 float64 = (v50 - 1e10)
                _ = v293
            } else {
                if true {
                    switch fi133 {
                    case -747:
                        if (fi133 >= 127) {
                            switch "a" {
                            case "\n\t":
                                if len(v5) > 0 {
                                    v5[0] = (v135 + (^(v135)))
                                }
                            default:
                                if (252 > 165) {
                                    var v294 int64 = genFn0((fi133 << ((991) & 63)), v45, false)
                                    var v295 uint8 = 0
                                    for fi296 := 0; fi296 < 16; fi296++ {
                                        var v297 uint = genFn1()
                                        v298 := v3[fi296]
                                        var v299 int16 = 0
                                        v300 := GenS18{F0: v45}
                                        var v301 int8 = 73
                                        _ = fi296
                                        _ = v297
                                        _ = v298
                                        _ = v299
                                        _ = v300
                                        _ = v301
                                    }
                                    _ = v294
                                    _ = v295
                                }
                            }
                        } else {
                            if true {
                                var v302 bool = false
                                if (-1 == -1) {
                                    if v302 {
                                        _ = len(v1)
                                        v303 := GenS12{F0: ("foo bar" + ("\n\t" + "a")), F1: (v135 / (((v135 >> ((genFn1()) & 63))) | 1)), F2: -485}
                                        var v304 uint16 = (v55 / (((^(v55))) | 1))
                                        _ = v303
                                        _ = v304
                                    }
                                }
                                _ = v302
                            } else {
                                v55 = v55
                            }
                            v2 = append(v2, 9223372036854775807)
                            v305 := GenS31{F0: (("foo bar" + "") + "\n\t"), F1: 4294967295, F2: (v135 % (((v135 % ((v135) | 1))) | 1))}
                            _ = v305
                        }
                    case 32768:
                        if len(v5) > 0 {
                            v5[0] = (^(v135))
                        }
                    }
                    if (42 < 597) {
                        if (!(genFn0((-(fi133)), (v45 % (((v45 + v45)) | 1)), (!true)) >= -9223372036854775808)) {
                            v306 := GenS32{F0: 253}
                            _ = v306
                        }
                        v45 -= (v45 - 18446744073709551615)
                    }
                    var v307 string = ("" + ("" + ("foo bar" + ("foo bar" + ""))))
                    _ = v307
                } else {
                    v2 = append(v2, fi133)
                }
            }
            _ = v135
        }
        v50 -= v50
        v50 = v50
        _ = fi133
    }
    var v308 uint = 625
    v308 += 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v308:", v308)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v308
    _ = v4
    _ = v45
    _ = v49
    _ = v5
    _ = v50
    _ = v55
}
