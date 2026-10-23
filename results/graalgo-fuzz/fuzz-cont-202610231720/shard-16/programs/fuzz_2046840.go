package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int16; F1 int16; F2 int64 }
type GenS3 struct { F0 int16; F1 int8; F2 bool }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int16; F1 uint8; F2 uint32 }
type GenS8 struct { F0 int8; F1 uint }
type GenS9 struct { F0 int16; F1 uint; F2 int64 }
type GenS10 struct { F0 uint8; F1 uint16; F2 int8 }
type GenS11 struct { F0 int8; F1 uint16; F2 int32 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 bool; F1 uint8; F2 int32 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 string }
type GenS16 struct { F0 int8; F1 string; F2 int }
type GenS17 struct { F0 int8; F1 uint32 }
type GenS18 struct { F0 string; F1 string }
type GenS19 struct { F0 int64; F1 int64 }
func genFn0(p0 int32, p1 string) uint16 {
    if (42 == -571) {
        v5 := []int{-9223372036854775808, 32768, -2147483648}
        var v6 string = "hello"
        var v7 float64 = 0.1
        _ = v5
        _ = v6
        _ = v7
    } else {
        for fi8 := 0; fi8 < 8; fi8++ {
            for fi9 := 0; fi9 < 15; fi9++ {
                fmt.Println(p0, fi8)
                if (128 <= 255) {
                    if (0 <= 2) {
                        p0 = (^(p0))
                        var v10 int16 = 0
                        switch p1 {
                        case "foo bar":
                            var v11 string = (p1 + p1)
                            _ = v11
                        case "hello":
                            p1 = p1
                        default:
                            v12 := GenS1{F0: -30, F1: 414, F2: 2}
                            _ = v12
                        }
                        _ = v10
                    }
                } else {
                    v13 := []int{(fi9 % ((127) | 1))}
                    if ("hello" == ((p1 + "\n\t") + "foo bar")) {
                        switch (p1 + "") {
                        case "foo bar":
                            v14 := GenS2{F0: -679, F1: 42, F2: 157}
                            _ = v14
                        case "a":
                            if false {
                                fmt.Println(p0)
                                if (-1 <= -55) {
                                    p0 = -2147483648
                                }
                            }
                        case "":
                            if (false && true) {
                                fmt.Println(fi9, fi8)
                            }
                        default:
                            switch p1 {
                            case "hello":
                                v15 := GenS3{F0: 1000, F1: 1, F2: true}
                                _ = v15
                            case "a":
                                p0 = 42
                            default:
                                if (!(true && false)) {
                                    p0 -= p0
                                    fmt.Println(p0, p1)
                                    fmt.Println(fi8, p1)
                                } else {
                                    v16 := GenS2{F0: -600, F1: 32767, F2: -1}
                                    var v17 uint = 498
                                    v13 = append(v13, fi8)
                                    _ = v16
                                    _ = v17
                                }
                                p0 = p0
                                v18 := GenS2{F0: 42, F1: 1000, F2: 0}
                                v19 := GenS3{F0: 42, F1: 0, F2: (!(true && false))}
                                if (858 < 2) {
                                    var v20 int = (fi8 - fi8)
                                    p0 = (p0 / (((p0 ^ p0)) | 1))
                                    if len(v13) >= 2 {
                                        v13 = v13[1:len(v13)]
                                    }
                                    _ = v20
                                } else {
                                    _ = len(v13)
                                    fmt.Println(p0, fi8, fi9)
                                    v13 = append(v13, (fi9 - fi9))
                                }
                                _ = v18
                                _ = v19
                            }
                        }
                    } else {
                        p1 = "\n\t"
                        if false {
                            v13 = append(v13, -1)
                            if false {
                                if false {
                                    fmt.Println(fi9, fi8, p1)
                                    if true {
                                        switch "a" {
                                        case "hello":
                                            var v21 float64 = (-2.25 + 0.1)
                                            _ = v21
                                        case "foo bar":
                                            var v22 int = (fi8 ^ fi8)
                                            _ = v22
                                        default:
                                            p0 = p0
                                            fmt.Println(fi9)
                                            var v23 uint64 = 2
                                            _ = v23
                                        }
                                        p1 = ("a" + p1)
                                        v24 := GenS1{F0: -73, F1: 931, F2: 2147483648}
                                        p1 = (((p1 + ("foo bar" + "foo bar")) + "") + "\n\t")
                                        if true {
                                            var v25 int64 = -588
                                            p1 = p1
                                            v25 = v25
                                            _ = v25
                                        } else {
                                            v26 := GenS4{F0: -2147483648, F1: 100}
                                            _ = v26
                                        }
                                        _ = v24
                                    } else {
                                        _ = len(v13)
                                        var v27 int16 = 976
                                        switch fi8 {
                                        case 360:
                                            v28 := GenS2{F0: (v27 % ((0) | 1)), F1: (v27 & 42), F2: 1}
                                            _ = v28
                                        case 917:
                                            var v29 bool = (false && true)
                                            _ = v29
                                        case 4294967295:
                                            if len(v13) > 0 {
                                                v13[0] = fi8
                                            }
                                        }
                                        _ = v27
                                    }
                                }
                                if (!(!false)) {
                                    if len(v13) > 0 {
                                        v13[0] = fi9
                                    }
                                    v30 := GenS1{F0: 42, F1: 42, F2: 4294967295}
                                    v31 := GenS2{F0: -1, F1: 1000, F2: -465}
                                    var v32 int64 = 9223372036854775807
                                    fmt.Println(fi8, fi9, p0)
                                    _ = v30
                                    _ = v31
                                    _ = v32
                                }
                                var v33 int8 = -68
                                _ = v33
                            }
                            var v34 uint64 = 0
                            _ = v34
                        }
                        var v35 int8 = 100
                        _ = v35
                    }
                    v13 = append(v13, fi9)
                    _ = v13
                }
                v36 := GenS5{F0: 200}
                _ = fi9
                _ = v36
            }
            if true {
                fmt.Println(p1, p0)
                if false {
                    p1 = p1
                }
            } else {
                for fi37 := 0; fi37 < 4; fi37++ {
                    fmt.Println(p1, p1, fi37)
                    fmt.Println(p0)
                    p0 -= p0
                    _ = fi37
                }
                p1 = p1
                switch 42 {
                case 32767:
                    for fi38 := 0; fi38 < 19; fi38++ {
                        v39 := []int{(fi38 & fi38), (fi8 ^ 31), (fi38 / ((fi38) | 1)), fi8}
                        p1 = (((p1 + (p1 + p1)) + "hello") + "a")
                        _ = fi38
                        _ = v39
                    }
                case -9223372036854775808:
                    var v40 int32 = (p0 * p0)
                    _ = v40
                case 71:
                    var v41 int32 = p0
                    _ = v41
                }
            }
            _ = fi8
        }
        v42 := []int{-365, 127}
        _ = v42
    }
    var v43 float64 = 0.0
    v44 := []int{-286}
    _ = p0
    _ = p1
    _ = v43
    _ = v44
    return 32767
}
func genFn1(p0 string, p1 int32) uint {
    v45 := GenS3{F0: -1, F1: 42, F2: ((true || false) || true)}
    var v46 int64 = int64(p1)
    v46 = (v46 + v46)
    v47 := GenS6{F0: 1}
    _ = p0
    _ = p1
    _ = v45
    _ = v46
    _ = v47
    return 440
}

func main() {
    v1 := []int{-662, 0, 190}
    v2 := make(map[int]int)
    var v3 uint32 = 65535
    v4 := make(map[int]bool)
    v3 = v3
    for fi48 := 0; fi48 < 13; fi48++ {
        fmt.Println(fi48)
        if true {
            if (true && (-1 != -759)) {
                v49 := GenS7{F0: 32767, F1: 1, F2: v3}
                _ = v49
            }
        }
        v3 = (v3 / (((v3 >> ((genFn1("", -1)) & 63))) | 1))
        _ = fi48
    }
    var v50 uint8 = 128
    v1 = append(v1, 4294967295)
    if ((float64(0.1) - (((-2.25 / 1.0) * (0.1 + float64(0.1))) / 1.0)) < (((1e10 - 1.5) + 0.0) - (0.1 * (3.14159 * (-2.25 * 1.5))))) {
        var v51 uint64 = 453
        v52 := GenS8{F0: int8(v3), F1: 4294967295}
        v53 := GenS1{F0: 100, F1: v51, F2: 65535}
        _ = v51
        _ = v52
        _ = v53
    } else {
        switch 42 {
        case 384:
            for fi54 := 0; fi54 < 2; fi54++ {
                var v55 float32 = (0.1 / 1.5)
                _ = fi54
                _ = v55
            }
        case -990:
            switch "" {
            case "a":
                fmt.Println(v3, v3)
            case "":
                v56 := GenS9{F0: -32768, F1: 925, F2: 167}
                _ = v56
            }
        case 32767:
            var v57 float32 = -2.25
            _ = v57
        }
        v50 = (^(v50))
    }
    switch 421 {
    case 519:
        var v58 int = 42
        _ = v58
    case 4294967295:
        var v59 uint8 = 2
        _ = v59
    default:
        v60 := GenS10{F0: (v50 | (v50 - 1)), F1: 42, F2: 1}
        _ = v60
    }
    v1 = append(v1, 2147483647)
    fmt.Println(v3, v50)
    fmt.Println(v50, v50)
    v61 := GenS4{F0: -1, F1: -1}
    for fi62 := 0; fi62 < 9; fi62++ {
        for fi63 := 0; fi63 < 7; fi63++ {
            if false {
                if (0.1 > (((3.14159 - ((((0.1 + -2.25) + 3.14159) - -2.25) + 1.5)) * -2.25) * 3.14159)) {
                    var v64 uint64 = 273
                    switch (fi62 | 2147483647) {
                    case 42:
                        var v65 uint64 = 117
                        v66 := GenS11{F0: 100, F1: 1, F2: 2147483647}
                        v50 = v50
                        _ = v65
                        _ = v66
                    case 2147483647:
                        var v67 uint16 = 32767
                        _ = v67
                    default:
                        switch (-(fi63)) {
                        case -746:
                            v50 -= (^(v50))
                        }
                        if (!false) {
                            var v68 string = (("foo bar" + "foo bar") + "foo bar")
                            var v69 uint64 = 2
                            switch (("a" + v68) + (v68 + (v68 + (v68 + v68)))) {
                            case "":
                                var v70 int = (^(fi62))
                                _ = v70
                            }
                            _ = v68
                            _ = v69
                        } else {
                            v3 = (v3 / ((925) | 1))
                        }
                        v64 = 193
                        var v71 float64 = 3.14159
                        delete(v2, fi62)
                        _ = v71
                    }
                    _ = v64
                } else {
                    if (-1 >= 119) {
                        var v72 bool = true
                        _ = v72
                    }
                    if false {
                        var v73 uint = 1
                        v73 -= (v73 & 255)
                        _ = v73
                    } else {
                        fmt.Println(v50, v50, v3)
                        fmt.Println(v50, v3)
                    }
                    v3 = v3
                }
                switch fi63 {
                case 127:
                    if true {
                        var v74 int64 = 1
                        switch "foo bar" {
                        case "hello":
                            fmt.Println(fi63)
                        default:
                            var v75 int64 = (v74 >> ((562) & 63))
                            _ = v75
                        }
                        _ = v74
                    } else {
                        switch fi62 {
                        case 1:
                            var v76 uint8 = v50
                            _ = v76
                        }
                        var v77 uint32 = v3
                        _ = v77
                    }
                    fmt.Println(fi62)
                }
            }
            switch "foo bar" {
            case "hello":
                v3 = (v3 ^ 2147483647)
            case "\n\t":
                v78 := GenS11{F0: -122, F1: 1, F2: int32(fi62)}
                _ = v78
            }
            _ = fi63
        }
        for fi79 := 0; fi79 < 1; fi79++ {
            v50 += 150
            var v80 uint32 = v3
            if true {
                var v81 int8 = -99
                var v82 uint32 = v80
                fmt.Println(v3, v81, v80)
                var v83 int64 = 1
                switch "hello" {
                case "a":
                    var v84 float64 = float64(1e10)
                    var v85 bool = true
                    if (((v84 - v84) + (1e10 * float64(1.5))) >= -2.25) {
                        v81 = v81
                        switch "a" {
                        case "hello":
                            fmt.Println(v84, v80)
                        case "\n\t":
                            var v86 float32 = (float32(v84) * float32((-2.25 * (v84 - float64(((((1.5 - ((0.1 + float32(0.1)) - (((1e10 + 0.1) / 1.0) + 0.0))) * 1e10) * (1.5 + -2.25)) * 3.14159))))))
                            _ = v86
                        default:
                            v87 := v2[fi62]
                            _ = v87
                        }
                        v50 = v50
                        var v88 uint32 = (v82 ^ v3)
                        fmt.Println(v88)
                        v81 += (v81 % ((-6) | 1))
                        _ = v88
                    }
                    _ = v84
                    _ = v85
                }
                _ = v81
                _ = v82
                _ = v83
            }
            fmt.Println(v50, v80)
            fmt.Println(v50, v80, fi62)
            _ = fi79
            _ = v80
        }
        var v89 uint16 = genFn0(-2147483648, "\n\t")
        _ = fi62
        _ = v89
    }
    for fi90 := 0; fi90 < 1; fi90++ {
        if false {
            _ = len(v1)
        } else {
            fmt.Println(fi90, fi90, v50)
        }
        v91 := GenS12{F0: 502}
        var v92 bool = true
        _ = fi90
        _ = v91
        _ = v92
    }
    for fi93 := 0; fi93 < 12; fi93++ {
        var v94 float32 = float32(3.14159)
        for fi95 := 0; fi95 < 17; fi95++ {
            var v96 string = "foo bar"
            switch fi95 {
            case -681:
                v96 = v96
                if (4294967295 <= 4294967295) {
                    var v97 int64 = 42
                    _ = v97
                } else {
                    var v98 uint32 = v3
                    switch "a" {
                    case "hello":
                        var v99 int32 = -824
                        _ = v99
                    case "":
                        v50 = v50
                    }
                    var v100 float32 = (((1.5 + -2.25) + v94) - v94)
                    _ = v100
                    _ = v98
                }
                v101 := GenS10{F0: (v50 + v50), F1: 618, F2: -128}
                _ = v101
            }
            _ = fi95
            _ = v96
        }
        _ = fi93
        _ = v94
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v3, v50)
    for fi102 := 0; fi102 < 20; fi102++ {
        for fi103 := 0; fi103 < 15; fi103++ {
            v50 = 255
            fmt.Println(v50, v50)
            if true {
                v104 := GenS13{F0: false, F1: v50, F2: -316}
                var v105 int64 = 0
                if ((128 >= 42) && (!(42 <= 0))) {
                    var v106 uint64 = 638
                    switch (-(fi102)) {
                    case -9223372036854775808:
                        var v107 bool = false
                        _ = v107
                    case -2147483648:
                        if ((v106 & 4294967295) <= (v106 << ((42) & 63))) {
                            var v108 uint64 = (v106 % ((uint64(v50)) | 1))
                            v105 -= v105
                            if ((v105 >= v105) && false) {
                                v3 += v3
                                var v109 int16 = 895
                                var v110 float32 = (1e10 - 0.1)
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                switch (("" + "foo bar") + "") {
                                case "a":
                                    switch "a" {
                                    case "hello":
                                        if true {
                                            var v111 uint32 = v3
                                            v3 = v111
                                            _ = v111
                                        } else {
                                            fmt.Println(v109)
                                            v50 = 127
                                            v106 -= 0
                                        }
                                    }
                                case "foo bar":
                                    switch (fi103 << ((825) & 63)) {
                                    case 260:
                                        v105 -= (^(v105))
                                    default:
                                        var v112 float32 = (v110 - v110)
                                        _ = v112
                                    }
                                case "hello":
                                    fmt.Println(v3, fi103)
                                default:
                                    v110 = v110
                                }
                                _ = v109
                                _ = v110
                            }
                            _ = v108
                        } else {
                            delete(v4, fi103)
                            if false {
                                if true {
                                    if true {
                                        var v113 float32 = ((0.1 + -2.25) / 3.14159)
                                        v3 = (v3 - (v3 << ((223) & 63)))
                                        v1 = append(v1, 127)
                                        switch fi103 {
                                        case 2147483647:
                                            var v114 float32 = v113
                                            _ = v114
                                        case 128:
                                            if (-1 > 1) {
                                                var v115 uint8 = (v50 / ((18) | 1))
                                                _ = v115
                                            }
                                        }
                                        var v116 uint16 = 42
                                        fmt.Println(v50, v113, v113)
                                        var v117 float32 = float32(1e10)
                                        _ = v113
                                        _ = v116
                                        _ = v117
                                    } else {
                                        switch "" {
                                        case "\n\t":
                                            if (0 == -2147483648) {
                                                if true {
                                                    if len(v1) > 0 {
                                                        v1[0] = (fi103 * -973)
                                                    }
                                                }
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v3 = (v3 / ((1) | 1))
                                            } else {
                                                var v118 int = fi102
                                                _ = v118
                                            }
                                        case "":
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        case "foo bar":
                                            v119 := GenS14{F0: (v50 * (v50 * (v50 + (v50 | (v50 << ((861) & 63))))))}
                                            _ = v119
                                        default:
                                            fmt.Println(fi103, v3, v50)
                                        }
                                        fmt.Println(v3, fi103, v50)
                                        v106 = (^(v106))
                                    }
                                } else {
                                    if ((42 != 32767) || false) {
                                        fmt.Println(fi102, fi102, fi103)
                                        v106 -= 42
                                        if (v106 == (v106 & 540)) {
                                            var v120 float32 = (((float32(float64(float32(3.14159))) + (-2.25 * 0.1)) + 1.5) / float32(float64(((float32(1e10) * 0.1) + 1.5))))
                                            var v121 uint64 = 963
                                            switch (fi102 % ((-9223372036854775808) | 1)) {
                                            case 32768:
                                                var v122 int32 = 2147483647
                                                _ = v122
                                            default:
                                                v123 := GenS15{F0: "a"}
                                                _ = v123
                                            }
                                            _ = v120
                                            _ = v121
                                        } else {
                                            var v124 int8 = 0
                                            if (1.5 > 1.5) {
                                                _ = len(v4)
                                                var v125 int = (fi102 / (((fi103 % ((1) | 1))) | 1))
                                                v126 := GenS9{F0: int16(v50), F1: 77, F2: (v105 >> ((255) & 63))}
                                                _ = v125
                                                _ = v126
                                            }
                                            if false {
                                                _ = len(v4)
                                                v50 = (v50 >> ((652) & 63))
                                            } else {
                                                var v127 int = fi103
                                                if false {
                                                    v128 := GenS6{F0: -508}
                                                    _ = v128
                                                } else {
                                                    var v129 uint64 = (v106 / ((308) | 1))
                                                    switch "hello" {
                                                    case "hello":
                                                        if len(v1) > 0 {
                                                            v1[0] = fi103
                                                        }
                                                    }
                                                    _ = v129
                                                }
                                                _ = v127
                                            }
                                            _ = v124
                                        }
                                        v130 := v4[(-(fi102))]
                                        _ = v130
                                    }
                                }
                                v2[fi102] = (fi103 + fi103)
                            } else {
                                v131 := GenS12{F0: v3}
                                _ = len(v2)
                                _ = v131
                            }
                            var v132 uint64 = (^(v106))
                            _ = v132
                        }
                    }
                    if true {
                        v133 := GenS16{F0: -1, F1: (("\n\t" + "hello") + ("a" + "hello")), F2: (fi103 | (fi103 << ((42) & 63)))}
                        var v134 int = fi102
                        _ = v133
                        _ = v134
                    }
                    if false {
                        v3 += (v3 << ((255) & 63))
                        v135 := GenS7{F0: -750, F1: 42, F2: (v3 * (v3 + v3))}
                        _ = v135
                    }
                    if (!(212 != v50)) {
                        var v136 bool = false
                        _ = v136
                    }
                    _ = v106
                }
                _ = v104
                _ = v105
            } else {
                v137 := GenS5{F0: (v3 ^ (v3 / (((v3 >> ((4294967295) & 63))) | 1)))}
                if true {
                    var v138 int8 = 65
                    if ((true || (!(947 >= 1000))) && false) {
                        if true {
                            delete(v4, 9223372036854775807)
                            switch (fi103 & fi102) {
                            case 71:
                                v3 += (v3 ^ v3)
                            case -1:
                                var v139 int16 = 1000
                                _ = v139
                            case 0:
                                v4[(-(fi102))] = false
                            }
                        }
                        switch (fi102 - fi103) {
                        case -1:
                            v140 := GenS17{F0: v138, F1: (v3 << ((0) & 63))}
                            _ = v140
                        case -477:
                            _ = len(v1)
                        case 0:
                            var v141 uint = 1
                            _ = v141
                        default:
                            var v142 int64 = 42
                            _ = v142
                        }
                    }
                    fmt.Println(v50, fi102)
                    if true {
                        var v143 string = "a"
                        v50 -= v50
                        var v144 uint32 = 1
                        _ = v143
                        _ = v144
                    } else {
                        if (false || (-2.25 != 0.0)) {
                            var v145 string = "a"
                            v146 := GenS18{F0: ("foo bar" + "hello"), F1: (v145 + (v145 + v145))}
                            if (128 > 42) {
                                switch "\n\t" {
                                case "hello":
                                    v147 := GenS19{F0: -623, F1: -716}
                                    _ = v147
                                case "foo bar":
                                    switch "a" {
                                    case "foo bar":
                                        v145 = "foo bar"
                                    }
                                case "":
                                    fmt.Println(v3, v50)
                                }
                                v50 = (v50 ^ (v50 & (v50 << ((genFn1(((v145 + "a") + v145), 42)) & 63))))
                                var v148 uint8 = (v50 * (v50 * (v50 / (((v50 >> ((1) & 63))) | 1))))
                                v149 := GenS9{F0: -594, F1: 65535, F2: 119}
                                _ = v148
                                _ = v149
                            }
                            v50 = v50
                            v3 -= v3
                            _ = v145
                            _ = v146
                        } else {
                            var v150 int16 = -32768
                            var v151 uint = 2
                            var v152 uint32 = (^(v3))
                            _ = v150
                            _ = v151
                            _ = v152
                        }
                    }
                    _ = v138
                }
                _ = v137
            }
            _ = fi103
        }
        v1 = append(v1, -698)
        v4[(fi102 << ((genFn1("\n\t", 306)) & 63))] = (!(false || false))
        _ = fi102
    }
    fmt.Println(v50)
    var v153 int16 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v153:", v153)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v50:", v50)
    _ = v1
    _ = v153
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v61
}
