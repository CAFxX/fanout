package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 uint8; F2 string }
type GenS6 struct { F0 uint64; F1 string }
type GenS7 struct { F0 uint16; F1 int16; F2 int32 }
type GenS8 struct { F0 uint16; F1 bool; F2 int }
type GenS9 struct { F0 uint8; F1 uint }
type GenS10 struct { F0 int16; F1 int8 }
type GenS11 struct { F0 int8; F1 uint64 }
type GenS12 struct { F0 int64; F1 uint16 }
type GenS13 struct { F0 bool; F1 int; F2 uint }
func genFn0(p0 bool, p1 bool, p2 uint16) int {
    if p0 {
        if true {
            if (("foo bar" + "\n\t") >= "a") {
                p0 = (976 < 230)
                p1 = (32767 != 1000)
            }
            if p0 {
                var v6 int32 = 0
                switch "" {
                case "":
                    var v7 int = 9223372036854775807
                    _ = v7
                case "\n\t":
                    for fi8 := 0; fi8 < 1; fi8++ {
                        v9 := []int{128, (fi8 - fi8), fi8}
                        v10 := GenS1{F0: 255}
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                default:
                    var v11 int8 = -43
                    _ = v11
                }
                p2 = (p2 % ((255) | 1))
                _ = v6
            }
            if p1 {
                v12 := []int{-75}
                p2 = p2
                v12 = append(v12, 4294967295)
                _ = v12
            } else {
                p1 = (18446744073709551615 != 2)
                for fi13 := 0; fi13 < 15; fi13++ {
                    switch (fi13 + (fi13 * int(p2))) {
                    case 32768:
                        for fi14 := 0; fi14 < 16; fi14++ {
                            v15 := GenS2{F0: 1, F1: (fi13 & fi14)}
                            var v16 float64 = (1.5 - -2.25)
                            if p1 {
                                switch "hello" {
                                case "foo bar":
                                    p1 = (!p0)
                                default:
                                    if p1 {
                                        v17 := GenS2{F0: 255, F1: (fi14 * fi14)}
                                        var v18 int = (fi14 + (fi13 / (((fi13 & fi14)) | 1)))
                                        _ = v17
                                        _ = v18
                                    } else {
                                        var v19 uint16 = p2
                                        v20 := GenS2{F0: 1, F1: (fi13 >> ((874) & 63))}
                                        var v21 uint64 = 1
                                        _ = v19
                                        _ = v20
                                        _ = v21
                                    }
                                }
                                switch ("foo bar" + "\n\t") {
                                case "hello":
                                    if p0 {
                                        p1 = (!(!((v16 + v16) < v16)))
                                        var v22 uint64 = 17
                                        _ = v22
                                    } else {
                                        var v23 int = -23
                                        p1 = false
                                        _ = v23
                                    }
                                default:
                                    var v24 float64 = v16
                                    _ = v24
                                }
                                switch fi13 {
                                case 32768:
                                    v25 := []int{127, fi13, -1}
                                    if ((p2 | 859) < (p2 % (((p2 + (p2 - (p2 >> ((1) & 63))))) | 1))) {
                                        switch "\n\t" {
                                        case "hello":
                                            fmt.Println(v16)
                                        }
                                        var v26 float64 = float64(float32(v16))
                                        var v27 int = fi14
                                        v28 := GenS3{F0: 65535}
                                        _ = v26
                                        _ = v27
                                        _ = v28
                                    }
                                    var v29 int32 = 2147483647
                                    _ = v25
                                    _ = v29
                                case 2147483647:
                                    if false {
                                        var v30 bool = p0
                                        v31 := []int{fi13, -984, fi14}
                                        if (!p1) {
                                            switch fi13 {
                                            case 128:
                                                p2 = (p2 >> ((780) & 63))
                                            case 1:
                                                var v32 float32 = 0.1
                                                _ = v32
                                            case -9223372036854775808:
                                                var v33 uint64 = 531
                                                _ = v33
                                            }
                                        }
                                        _ = v30
                                        _ = v31
                                    } else {
                                        fmt.Println(p2, p0)
                                    }
                                }
                            }
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    case -565:
                        v34 := []int{(^(fi13))}
                        _ = v34
                    case 128:
                        for fi35 := 0; fi35 < 10; fi35++ {
                            if (2 >= 224) {
                                v36 := GenS2{F0: 4294967295, F1: (fi13 / ((42) | 1))}
                                switch "hello" {
                                case "":
                                    p2 -= 299
                                case "hello":
                                    fmt.Println(p0)
                                case "\n\t":
                                    v37 := []int{457}
                                    _ = v37
                                default:
                                    if p1 {
                                        p1 = (!(3.14159 <= ((-2.25 + -2.25) * 1e10)))
                                    } else {
                                        var v38 int64 = 42
                                        _ = v38
                                    }
                                }
                                p2 = p2
                                _ = v36
                            } else {
                                var v39 uint = 2
                                if (127 == 127) {
                                    var v40 uint16 = 255
                                    switch "foo bar" {
                                    case "foo bar":
                                        v41 := []int{179, fi35, (fi13 ^ (fi35 << ((v39) & 63)))}
                                        _ = v41
                                    case "\n\t":
                                        var v42 uint64 = 0
                                        _ = v42
                                    case "a":
                                        var v43 int = fi13
                                        _ = v43
                                    }
                                    var v44 uint32 = 156
                                    _ = v40
                                    _ = v44
                                }
                                var v45 uint16 = p2
                                _ = v39
                                _ = v45
                            }
                            var v46 bool = ((((0.0 + -2.25) * 0.1) / 1.5) < float64((float32(3.14159) - 1e10)))
                            _ = fi35
                            _ = v46
                        }
                    default:
                        p0 = (1 <= 591)
                    }
                    _ = fi13
                }
            }
        } else {
            switch ("foo bar" + "a") {
            case "":
                p0 = ((p2 >= (p2 & p2)) && p1)
            default:
                switch ((("hello" + "foo bar") + "hello") + ("" + ("hello" + "hello"))) {
                case "":
                    if ((((0.0 + 0.1) + ((3.14159 - 0.0) - -2.25)) - (3.14159 - 1e10)) > ((((0.0 * -2.25) - ((3.14159 - (0.0 + float64(-2.25))) + (0.1 * 3.14159))) - 0.1) / -2.25)) {
                        fmt.Println(p1, p0)
                        var v47 int16 = -550
                        _ = v47
                    }
                case "\n\t":
                    fmt.Println(p2, p2, p0)
                default:
                    v48 := []int{}
                    _ = v48
                }
            }
            for fi49 := 0; fi49 < 5; fi49++ {
                for fi50 := 0; fi50 < 13; fi50++ {
                    v51 := []int{9223372036854775807, (-(fi49)), fi49}
                    var v52 uint16 = 32767
                    _ = fi50
                    _ = v51
                    _ = v52
                }
                fmt.Println(p0, fi49, p0)
                _ = fi49
            }
            if p0 {
                v53 := []int{332, 9223372036854775807}
                fmt.Println(p2)
                for fi54 := 0; fi54 < 20; fi54++ {
                    p1 = (float32(3.14159) != (float32((-2.25 * ((float64((-2.25 / 0.1)) + ((1.5 * 3.14159) + 0.0)) + float64((float32(3.14159) * (0.1 - 1.5)))))) - (((((1e10 / 3.14159) * 1e10) / ((3.14159 - (0.1 * 0.0)) + (0.0 + (-2.25 + (-2.25 / 1.5))))) + -2.25) * 0.0)))
                    var v55 string = "a"
                    switch ((v55 + v55) + v55) {
                    case "hello":
                        var v56 uint64 = uint64(p2)
                        _ = v56
                    default:
                        switch fi54 {
                        case 128:
                            for fi57 := 0; fi57 < 4; fi57++ {
                                v58 := GenS3{F0: 65535}
                                p2 = 0
                                p2 = p2
                                _ = fi57
                                _ = v58
                            }
                        }
                    }
                    _ = fi54
                    _ = v55
                }
                _ = v53
            } else {
                fmt.Println(p1)
                p0 = p1
                var v59 uint16 = 2
                _ = v59
            }
        }
    } else {
        var v60 uint8 = 1
        fmt.Println(p2, p0)
        _ = v60
    }
    switch 0 {
    case 32768:
        fmt.Println(p1, p2)
    }
    _ = p0
    _ = p1
    _ = p2
    return -467
}
func genFn1(p0 string, p1 uint64) int8 {
    if false {
        var v61 float32 = (0.0 * (float32(1.5) * ((-2.25 / (0.1 * 3.14159)) - 0.1)))
        p0 = ""
        if true {
            for fi62 := 0; fi62 < 17; fi62++ {
                v61 -= float32(0.0)
                v63 := GenS4{F0: 0}
                var v64 int32 = 1
                _ = fi62
                _ = v63
                _ = v64
            }
            v61 = v61
        }
        _ = v61
    } else {
        v65 := GenS1{F0: 637}
        var v66 int16 = -957
        v67 := []int{32768, 128, -269}
        _ = v65
        _ = v66
        _ = v67
    }
    for fi68 := 0; fi68 < 19; fi68++ {
        fmt.Println(p1)
        p1 = (p1 ^ p1)
        _ = fi68
    }
    for fi69 := 0; fi69 < 19; fi69++ {
        p1 -= p1
        fmt.Println(p1, p1)
        var v70 float32 = (0.1 + (float32(1e10) / 1.0))
        _ = fi69
        _ = v70
    }
    for fi71 := 0; fi71 < 7; fi71++ {
        switch p0 {
        case "foo bar":
            var v72 int8 = -93
            _ = v72
        case "":
            switch "foo bar" {
            case "hello":
                var v73 int8 = -78
                _ = v73
            }
        default:
            for fi74 := 0; fi74 < 8; fi74++ {
                var v75 int = genFn0((true || false), true, uint16(p1))
                v76 := []int{}
                fmt.Println(p1)
                _ = fi74
                _ = v75
                _ = v76
            }
        }
        _ = fi71
    }
    _ = p0
    _ = p1
    return -72
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := []uint32{205, 890, 1}
    v3 := make(map[uint32]int)
    var v4 int8 = -79
    var v5 uint8 = 168
    var v77 int16 = 1000
    switch -9223372036854775808 {
    case 46:
        v77 = v77
    case 4294967295:
        _ = len(v3)
    case 32767:
        var v78 int8 = (v4 | (^(v4)))
        _ = v78
    }
    v79 := GenS5{F0: (v5 * v5), F1: (v5 - v5), F2: ("" + "")}
    for fi80 := 0; fi80 < 3; fi80++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v5)
        fmt.Println(v4)
        _ = fi80
    }
    v81 := GenS3{F0: v1}
    var v82 int32 = 2147483647
    for fi83 := 0; fi83 < 13; fi83++ {
        if false {
            if (false && ((!false) && false)) {
                fmt.Println(v5, v4)
            }
        } else {
            v77 = (-(v77))
            for fi84 := 0; fi84 < 17; fi84++ {
                if false {
                    var v85 uint = 42
                    _ = v85
                }
                _ = fi84
            }
        }
        switch (fi83 / ((fi83) | 1)) {
        case 1:
            for fi86 := 0; fi86 < 3; fi86++ {
                if false {
                    if true {
                        var v87 int16 = 1000
                        switch (-(fi86)) {
                        case 350:
                            var v88 int16 = 277
                            _ = v88
                        default:
                            switch 128 {
                            case -2147483648:
                                if false {
                                    v77 -= -67
                                    if ((v4 >= v4) && (v1 < (v1 << ((42) & 63)))) {
                                        v1 = v1
                                        var v89 int64 = 0
                                        _ = v89
                                    } else {
                                        fmt.Println(v1, v5, v1)
                                    }
                                    var v90 uint32 = 78
                                    _ = v90
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if ((!("hello" > ("\n\t" + "a"))) || true) {
                                        v91 := GenS3{F0: 973}
                                        _ = v91
                                    } else {
                                        v87 += 1
                                        v77 = v87
                                        if false {
                                            if len(v2) > 0 {
                                                v2[0] = 0
                                            }
                                            var v92 uint64 = 18446744073709551615
                                            var v93 uint32 = 987
                                            _ = v92
                                            _ = v93
                                        } else {
                                            switch "a" {
                                            case "hello":
                                                if (31 == int64(v1)) {
                                                    v87 = v87
                                                    v77 -= v77
                                                } else {
                                                    var v94 float32 = ((1e10 - 3.14159) * 0.1)
                                                    _ = v94
                                                }
                                                v95 := GenS1{F0: (v1 - v1)}
                                                var v96 uint32 = 2147483648
                                                _ = v95
                                                _ = v96
                                            default:
                                                v97 := GenS6{F0: 18446744073709551615, F1: "a"}
                                                _ = v97
                                            }
                                        }
                                    }
                                    v3[804] = (fi83 - (fi86 ^ 42))
                                    var v98 string = "a"
                                    v1 = (^(v1))
                                    switch fi83 {
                                    case 9223372036854775807:
                                        var v99 float64 = 3.14159
                                        _ = v99
                                    case 128:
                                        var v100 bool = (!true)
                                        _ = v100
                                    case -2147483648:
                                        var v101 uint8 = v5
                                        _ = v101
                                    default:
                                        var v102 int = genFn0(true, true, 42)
                                        _ = v102
                                    }
                                    v103 := GenS2{F0: 869, F1: (-(fi83))}
                                    v3[206] = (fi86 * (fi83 % ((fi86) | 1)))
                                    var v104 int = genFn0(false, (!true), 354)
                                    _ = v103
                                    _ = v104
                                    _ = v98
                                }
                                if false {
                                    v77 = (v77 + v77)
                                    if true {
                                        if (v4 <= v4) {
                                            v2 = append(v2, 2)
                                            v3[2147483648] = (fi86 - fi83)
                                            v105 := GenS6{F0: 18446744073709551615, F1: ""}
                                            _ = v105
                                        } else {
                                            v106 := GenS7{F0: 255, F1: v77, F2: int32(v87)}
                                            v1 = (v1 % ((18446744073709551615) | 1))
                                            if false {
                                                switch (("hello" + "hello") + "a") {
                                                case "hello":
                                                    if true {
                                                        if (!false) {
                                                            if false {
                                                                var v107 uint32 = 84
                                                                v107 -= (v107 % ((2147483648) | 1))
                                                                var v108 int = (fi83 ^ (fi86 ^ (fi83 % ((-9223372036854775808) | 1))))
                                                                _ = len(v3)
                                                                v109 := GenS2{F0: v1, F1: 42}
                                                                _ = v107
                                                                _ = v108
                                                                _ = v109
                                                            } else {
                                                                v110 := GenS7{F0: 2, F1: v87, F2: -2147483648}
                                                                _ = v110
                                                            }
                                                            var v111 int8 = (v4 - 100)
                                                            fmt.Println(v77, v5, v77)
                                                            _ = v111
                                                        }
                                                        if (!(("a" + "hello") >= "a")) {
                                                            v112 := GenS4{F0: v1}
                                                            v1 = v1
                                                            _ = v112
                                                        } else {
                                                            var v113 int = (-(fi83))
                                                            v114 := GenS2{F0: 255, F1: v113}
                                                            var v115 int32 = (v82 * v82)
                                                            _ = v113
                                                            _ = v114
                                                            _ = v115
                                                        }
                                                    }
                                                default:
                                                    v77 -= (v77 % ((-32768) | 1))
                                                }
                                            } else {
                                                if len(v2) > 0 {
                                                    v2[0] = 416
                                                }
                                            }
                                            _ = v106
                                        }
                                    }
                                }
                                v5 = (v5 << (((v1 - 352)) & 63))
                            default:
                                v82 += (-(v82))
                            }
                        }
                        switch "hello" {
                        case "\n\t":
                            if len(v2) > 0 {
                                v2[0] = 521
                            }
                        case "hello":
                            if false {
                                v116 := GenS5{F0: v5, F1: v5, F2: ((("\n\t" + "hello") + "foo bar") + ("hello" + "a"))}
                                fmt.Println(v87, v1, v4)
                                v117 := GenS3{F0: v1}
                                _ = v116
                                _ = v117
                            } else {
                                v4 -= v4
                                if len(v2) > 0 {
                                    v2[0] = 346
                                }
                            }
                        default:
                            v118 := GenS6{F0: 0, F1: "a"}
                            _ = v118
                        }
                        _ = v87
                    }
                    v82 += v82
                    _ = len(v3)
                }
                switch (fi86 / (((^(fi86))) | 1)) {
                case 1:
                    fmt.Println(v1, fi83)
                case 339:
                    var v119 uint64 = 1
                    _ = v119
                }
                _ = fi86
            }
        case -228:
            switch (fi83 / ((fi83) | 1)) {
            case 32768:
                for fi120 := 0; fi120 < 16; fi120++ {
                    v121 := GenS8{F0: 1, F1: (!((true && true) || ((65535 <= 32767) && (!((0.1 * (3.14159 * (0.1 * 3.14159))) != 0.0))))), F2: fi120}
                    v122 := GenS1{F0: v1}
                    v5 = (v5 * 128)
                    var v123 uint32 = 2147483647
                    switch "" {
                    case "hello":
                        switch (("" + "a") + "") {
                        case "foo bar":
                            v124 := GenS9{F0: v5, F1: v1}
                            _ = v124
                        case "":
                            if false {
                                v4 += (v4 * genFn1("\n\t", 1))
                                v125 := GenS10{F0: 42, F1: v4}
                                v1 = 1
                                var v126 bool = true
                                var v127 int32 = v82
                                _ = v125
                                _ = v126
                                _ = v127
                            }
                        case "hello":
                            v128 := v3[1]
                            v129 := v3[v123]
                            _ = v128
                            _ = v129
                        }
                    case "":
                        var v130 uint16 = 32767
                        _ = v130
                    case "a":
                        var v131 string = "hello"
                        _ = v131
                    default:
                        fmt.Println(v5, fi120, v123)
                    }
                    _ = fi120
                    _ = v121
                    _ = v122
                    _ = v123
                }
            }
        default:
            var v132 uint = 613
            _ = v132
        }
        switch "a" {
        case "foo bar":
            if true {
                var v133 uint32 = 2147483647
                _ = v133
            }
        case "":
            v77 += (v77 - (v77 % ((1000) | 1)))
        default:
            fmt.Println(v5, v4, v82)
        }
        _ = fi83
    }
    for fi134 := 0; fi134 < 18; fi134++ {
        v135 := GenS11{F0: (v4 ^ (^(v4))), F1: 932}
        for fi136 := 0; fi136 < 8; fi136++ {
            v1 += v1
            switch genFn0(((v77 ^ (v77 & (v77 + int16(v1)))) < (^(v77))), false, 42) {
            case -1:
                v137 := GenS12{F0: 9223372036854775807, F1: 32767}
                _ = v137
            case 477:
                fmt.Println(fi134)
            case 32768:
                v138 := GenS5{F0: 0, F1: 128, F2: "a"}
                _ = v138
            default:
                var v139 float32 = 1e10
                _ = v139
            }
            v3[2] = (fi136 - (fi134 & fi136))
            if (!(18446744073709551615 >= 42)) {
                var v140 string = "a"
                _ = v140
            }
            _ = fi136
        }
        var v141 string = "foo bar"
        _ = fi134
        _ = v135
        _ = v141
    }
    v142 := GenS5{F0: (^(v5)), F1: 25, F2: "foo bar"}
    _ = len(v2)
    v82 += v82
    v143 := GenS12{F0: 0, F1: 2}
    if (32768 > 0) {
        var v144 uint32 = uint32(v5)
        _ = v144
    } else {
        fmt.Println(v1)
    }
    v5 = 255
    v145 := GenS13{F0: (!false), F1: int(v5), F2: v1}
    fmt.Println(v4, v82, v82)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v77:", v77)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v142
    _ = v143
    _ = v145
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v77
    _ = v79
    _ = v81
    _ = v82
}
