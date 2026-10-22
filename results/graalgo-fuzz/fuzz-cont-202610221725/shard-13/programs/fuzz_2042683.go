package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 uint32; F2 int64 }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 int8; F1 int; F2 uint }
type GenS6 struct { F0 uint16; F1 uint16 }
type GenS7 struct { F0 string; F1 uint16 }
type GenS8 struct { F0 bool; F1 int32 }
type GenS9 struct { F0 bool; F1 int32 }
type GenS10 struct { F0 int; F1 uint32; F2 string }
type GenS11 struct { F0 uint64; F1 int }
type GenS12 struct { F0 uint64; F1 uint32 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int; F1 uint64 }
type GenS15 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS16 struct { F0 uint }
func genFn0() uint16 {
    if true {
        var v4 int8 = 1
        _ = v4
    }
    v5 := []int{-9223372036854775808, -1, 9223372036854775807, 369}
    var v6 uint = 42
    v7 := GenS1{F0: 1}
    for fi8 := 0; fi8 < 15; fi8++ {
        if true {
            if ((!true) || true) {
                var v9 uint16 = 0
                var v10 uint32 = 2
                _ = v10
                _ = v9
            } else {
                var v11 uint32 = 42
                if (false && true) {
                    v5 = append(v5, -242)
                    v6 = v6
                    v12 := GenS1{F0: 127}
                    _ = v12
                }
                _ = v11
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v6 = (v6 + v6)
        }
        _ = fi8
    }
    _ = v5
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 bool) int8 {
    for fi13 := 0; fi13 < 15; fi13++ {
        p0 = true
        fmt.Println(p0)
        _ = fi13
    }
    var v14 int8 = 1
    var v15 int8 = v14
    var v16 int32 = 2147483647
    _ = p0
    _ = v14
    _ = v15
    _ = v16
    return (-(v15))
}

func main() {
    var v1 uint8 = 147
    v2 := make(map[int]bool)
    v3 := make(map[int]bool)
    switch 9223372036854775807 {
    case 2147483647:
        v1 = (v1 << ((243) & 63))
    case 128:
        fmt.Println(v1)
    case 4294967295:
        v17 := []int{32767, -9223372036854775808}
        _ = v17
    default:
        if (!(18446744073709551615 <= 65535)) {
            for fi18 := 0; fi18 < 6; fi18++ {
                switch "\n\t" {
                case "":
                    fmt.Println(v1)
                case "\n\t":
                    if ("a" < "foo bar") {
                        for fi19 := 0; fi19 < 9; fi19++ {
                            var v20 uint = 327
                            switch "hello" {
                            case "hello":
                                if true {
                                    var v21 bool = ((((65535 <= 255) && (262 > 1)) || (true || false)) && (!(!false)))
                                    fmt.Println(v20)
                                    _ = v21
                                } else {
                                    switch "foo bar" {
                                    case "":
                                        var v22 string = "\n\t"
                                        _ = v22
                                    case "a":
                                        v1 = v1
                                    }
                                    var v23 string = "a"
                                    var v24 uint8 = 42
                                    _ = v23
                                    _ = v24
                                }
                            case "foo bar":
                                fmt.Println(v20, fi19)
                            }
                            var v25 int8 = 64
                            var v26 float32 = (3.14159 + float32(0.0))
                            switch "a" {
                            case "\n\t":
                                if (v1 != v1) {
                                    switch (("" + "hello") + ((("\n\t" + "a") + "") + "hello")) {
                                    case "":
                                        v27 := GenS1{F0: 1}
                                        _ = v27
                                    case "\n\t":
                                        var v28 uint8 = (^(v1))
                                        _ = v28
                                    }
                                    switch fi19 {
                                    case 218:
                                        fmt.Println(fi18, v20, fi18)
                                    }
                                } else {
                                    fmt.Println(v20)
                                    v29 := []int{(fi19 + fi19), int(v25), 185, (fi19 % ((fi19) | 1))}
                                    _ = v29
                                }
                            default:
                                var v30 uint64 = 4294967295
                                _ = v30
                            }
                            _ = fi19
                            _ = v20
                            _ = v25
                            _ = v26
                        }
                        if (("" + "foo bar") <= ("\n\t" + "a")) {
                            for fi31 := 0; fi31 < 19; fi31++ {
                                v2[(fi31 >> ((42) & 63))] = true
                                v1 -= (v1 & 127)
                                v3[fi31] = (2147483647 == 42)
                                _ = len(v2)
                                switch "" {
                                case "foo bar":
                                    if false {
                                        if (false && (false && (true || true))) {
                                            switch "foo bar" {
                                            case "\n\t":
                                                v1 = v1
                                            case "hello":
                                                v1 = (v1 & (v1 * v1))
                                            default:
                                                if (2 <= 65535) {
                                                    v3[(fi18 * (-(fi18)))] = false
                                                    if false {
                                                        if (false && (32767 > 0)) {
                                                            v32 := GenS1{F0: v1}
                                                            _ = v32
                                                        }
                                                        v1 -= v1
                                                    }
                                                } else {
                                                    if (false && false) {
                                                        var v33 int16 = -32768
                                                        fmt.Println(v33, v33, fi31)
                                                        _ = v33
                                                    } else {
                                                        v1 = (v1 / ((v1) | 1))
                                                    }
                                                    var v34 string = ("" + "a")
                                                    _ = v34
                                                }
                                            }
                                        }
                                    }
                                case "a":
                                    v1 -= v1
                                }
                                _ = fi31
                            }
                            if true {
                                switch (("\n\t" + "foo bar") + ("foo bar" + ("\n\t" + ("hello" + ("foo bar" + "\n\t"))))) {
                                case "":
                                    if (9223372036854775807 <= 9223372036854775807) {
                                        v1 = 127
                                        v1 = v1
                                        v1 = v1
                                    }
                                case "\n\t":
                                    if true {
                                        for fi35 := 0; fi35 < 2; fi35++ {
                                            var v36 string = ("" + "\n\t")
                                            _ = fi35
                                            _ = v36
                                        }
                                        v37 := GenS1{F0: v1}
                                        var v38 uint32 = 2147483648
                                        _ = v37
                                        _ = v38
                                    }
                                case "hello":
                                    v39 := GenS2{F0: -1}
                                    _ = v39
                                default:
                                    if (!(!false)) {
                                        v1 -= v1
                                        v1 += (v1 ^ 128)
                                    } else {
                                        if true {
                                            v2[fi18] = (!((3.14159 + (((((((3.14159 - 0.1) / (1e10 / ((1e10 / -2.25) / float64(-2.25)))) - (3.14159 * 1.5)) * -2.25) - -2.25) * 0.1) / (float64(3.14159) / (((3.14159 - -2.25) - -2.25) / 1.0)))) <= (0.1 * -2.25)))
                                            for fi40 := 0; fi40 < 1; fi40++ {
                                                v1 -= (v1 << ((65535) & 63))
                                                v1 -= v1
                                                if true {
                                                    if false {
                                                        v1 -= v1
                                                        v1 = (v1 | (v1 >> ((65535) & 63)))
                                                        v41 := GenS1{F0: (v1 ^ v1)}
                                                        _ = v41
                                                    }
                                                    v42 := GenS2{F0: -1}
                                                    v43 := GenS3{F0: fi40, F1: 0, F2: 710}
                                                    switch fi18 {
                                                    case 127:
                                                        v1 += 42
                                                    case 34:
                                                        v1 -= v1
                                                    case 9223372036854775807:
                                                        var v44 int16 = -1
                                                        _ = v44
                                                    default:
                                                        var v45 uint32 = 42
                                                        _ = v45
                                                    }
                                                    _ = v42
                                                    _ = v43
                                                }
                                                _ = fi40
                                            }
                                        } else {
                                            var v46 int8 = genFn1((uint16(v1) == 1))
                                            _ = v46
                                        }
                                        v1 -= v1
                                    }
                                }
                                for fi47 := 0; fi47 < 13; fi47++ {
                                    var v48 bool = (-32768 <= 42)
                                    v1 -= v1
                                    _ = fi47
                                    _ = v48
                                }
                            }
                        } else {
                            var v49 uint64 = 4294967295
                            var v50 int32 = 1
                            for fi51 := 0; fi51 < 14; fi51++ {
                                fmt.Println(fi18)
                                _ = fi51
                            }
                            _ = v49
                            _ = v50
                        }
                    } else {
                        var v52 uint32 = 0
                        fmt.Println(fi18)
                        _ = len(v3)
                        _ = v52
                    }
                default:
                    v53 := GenS3{F0: 60, F1: 1, F2: -370}
                    _ = v53
                }
                v54 := GenS2{F0: 0}
                v1 -= v1
                _ = fi18
                _ = v54
            }
            v1 = v1
        } else {
            var v55 float32 = 1e10
            if (681 >= 479) {
                var v56 int = 2147483647
                _ = v56
            }
            _ = len(v2)
            _ = v55
        }
    }
    if false {
        if false {
            for fi57 := 0; fi57 < 18; fi57++ {
                v1 = v1
                var v58 int = fi57
                _ = fi57
                _ = v58
            }
            switch 4294967295 {
            case 0:
                for fi59 := 0; fi59 < 19; fi59++ {
                    switch (-(fi59)) {
                    case -120:
                        for fi60 := 0; fi60 < 20; fi60++ {
                            var v61 float32 = 0.0
                            _ = fi60
                            _ = v61
                        }
                    }
                    _ = fi59
                }
            case -550:
                v62 := GenS2{F0: 42}
                _ = v62
            default:
                for fi63 := 0; fi63 < 13; fi63++ {
                    v64 := GenS4{F0: 42, F1: "\n\t"}
                    v1 = v1
                    if (!(-2.25 >= 1.5)) {
                        switch fi63 {
                        case 32768:
                            if false {
                                delete(v2, 0)
                                v65 := []int{(fi63 ^ 0), -846, (fi63 & fi63), -552}
                                fmt.Println(v1, fi63)
                                _ = v65
                            }
                        case -666:
                            var v66 int32 = -753
                            _ = v66
                        default:
                            if (42 > -1) {
                                v1 += 1
                                switch fi63 {
                                case 9223372036854775807:
                                    for fi67 := 0; fi67 < 15; fi67++ {
                                        if (1.5 == (0.1 / (3.14159 * 3.14159))) {
                                            v68 := GenS2{F0: -1}
                                            fmt.Println(fi67)
                                            if false {
                                                v1 -= v1
                                                v1 += (v1 / ((v1) | 1))
                                                v3[(^(fi67))] = (42 == 2)
                                            }
                                            _ = v68
                                        } else {
                                            switch (fi63 & fi67) {
                                            case 1:
                                                v69 := GenS5{F0: 98, F1: fi67, F2: 65535}
                                                delete(v2, fi67)
                                                _ = v69
                                            case 32767:
                                                fmt.Println(v1, fi63, v1)
                                            case 2147483647:
                                                v1 -= v1
                                            default:
                                                v70 := GenS4{F0: 175, F1: "a"}
                                                _ = v70
                                            }
                                            var v71 bool = false
                                            _ = v71
                                        }
                                        v1 = (v1 >> ((255) & 63))
                                        var v72 bool = false
                                        if v72 {
                                            v1 = v1
                                            var v73 string = "\n\t"
                                            v74 := GenS6{F0: 211, F1: genFn0()}
                                            _ = v73
                                            _ = v74
                                        }
                                        v72 = v72
                                        _ = fi67
                                        _ = v72
                                    }
                                }
                                var v75 uint = 42
                                _ = v75
                            }
                        }
                        switch 0 {
                        case 578:
                            for fi76 := 0; fi76 < 17; fi76++ {
                                switch ("a" + "hello") {
                                case "":
                                    var v77 uint64 = 42
                                    _ = v77
                                case "hello":
                                    v1 += v1
                                case "\n\t":
                                    v1 = v1
                                }
                                _ = fi76
                            }
                        }
                        v78 := GenS7{F0: "hello", F1: 983}
                        _ = v78
                    }
                    _ = fi63
                    _ = v64
                }
            }
        } else {
            fmt.Println(v1)
            v79 := GenS8{F0: false, F1: -1}
            _ = v79
        }
        v80 := []int{-1, 32768}
        if (!(255 != 0)) {
            var v81 uint32 = 42
            v1 -= v1
            _ = v81
        }
        _ = v80
    } else {
        var v82 uint8 = v1
        _ = v82
    }
    var v83 uint8 = v1
    var v84 string = "foo bar"
    switch 850 {
    case 42:
        fmt.Println(v1, v1, v1)
    case 0:
        v1 -= v83
    }
    var v85 float64 = float64(3.14159)
    switch v84 {
    case "":
        switch v84 {
        case "\n\t":
            switch 2147483647 {
            case 455:
                v83 -= (v1 % ((128) | 1))
            }
        case "hello":
            v1 -= v1
        default:
            _ = len(v2)
        }
    case "a":
        delete(v2, 32767)
    default:
        var v86 int16 = -32768
        _ = v86
    }
    for fi87 := 0; fi87 < 9; fi87++ {
        switch "hello" {
        case "\n\t":
            v1 += v83
        case "":
            if ((v83 <= (v83 >> ((0) & 63))) || false) {
                v85 = (((((v85 * float64((float32(v85) - 1e10))) / v85) / (0.0 * v85)) * 0.0) / v85)
            }
        case "foo bar":
            v88 := GenS9{F0: (65535 > 4294967295), F1: -753}
            _ = v88
        default:
            for fi89 := 0; fi89 < 13; fi89++ {
                if true {
                    var v90 uint = 2
                    _ = v90
                }
                if true {
                    switch (fi87 | 9223372036854775807) {
                    case -458:
                        var v91 int8 = -50
                        _ = v91
                    case -2147483648:
                        switch v84 {
                        case "foo bar":
                            var v92 uint64 = 18446744073709551615
                            _ = v92
                        case "\n\t":
                            if false {
                                var v93 string = "foo bar"
                                if false {
                                    v83 -= (^(v1))
                                } else {
                                    if true {
                                        var v94 uint8 = (^(v83))
                                        _ = v94
                                    } else {
                                        fmt.Println(v93)
                                    }
                                    v95 := []int{32768, (fi89 * (-(fi89))), (-(fi89))}
                                    if (!((!false) || (!((v84 + "a") != v84)))) {
                                        if (3.14159 == 0.0) {
                                            fmt.Println(v83, v83, v83)
                                            var v96 int8 = -128
                                            v2[fi89] = false
                                            _ = v96
                                        } else {
                                            var v97 uint64 = 710
                                            delete(v3, fi89)
                                            switch fi87 {
                                            case 128:
                                                v1 = 2
                                            case 102:
                                                if (!true) {
                                                    fmt.Println(v1)
                                                } else {
                                                    if true {
                                                        var v98 uint = 507
                                                        _ = v98
                                                    }
                                                    var v99 int8 = -69
                                                    if (0 >= 915) {
                                                        if (558 <= 4294967295) {
                                                            var v100 uint16 = 451
                                                            if len(v95) >= 2 {
                                                                v95 = v95[1:len(v95)]
                                                            }
                                                            _ = v100
                                                        } else {
                                                            v101 := GenS10{F0: 535, F1: 65535, F2: v84}
                                                            v1 = (v83 ^ (v83 % ((42) | 1)))
                                                            _ = v101
                                                        }
                                                        var v102 int = (fi87 - fi87)
                                                        _ = v102
                                                    }
                                                    fmt.Println(v93, v85, v85)
                                                    v84 = v84
                                                    _ = v99
                                                }
                                                if true {
                                                    switch (v84 + ("\n\t" + v93)) {
                                                    case "a":
                                                        var v103 float32 = (-2.25 - 0.1)
                                                        _ = v103
                                                    default:
                                                        fmt.Println(v1, fi87, v85)
                                                    }
                                                    var v104 int32 = 2147483647
                                                    v105 := GenS7{F0: "\n\t", F1: 0}
                                                    _ = v104
                                                    _ = v105
                                                }
                                                var v106 int64 = 42
                                                _ = v106
                                            case -2147483648:
                                                if true {
                                                    var v107 int8 = 42
                                                    _ = v107
                                                } else {
                                                    if len(v95) >= 2 {
                                                        v95 = v95[1:len(v95)]
                                                    }
                                                    v93 = "hello"
                                                    var v108 int32 = 1
                                                    if true {
                                                        var v109 int = (^(fi87))
                                                        if (true || (-1 > -128)) {
                                                            if (false || (124 != 0)) {
                                                                v1 = (v83 | 255)
                                                                v1 += (v1 & 44)
                                                                var v110 bool = true
                                                                _ = v110
                                                            }
                                                        } else {
                                                            if true {
                                                                fmt.Println(v83, v109)
                                                            }
                                                            fmt.Println(v84)
                                                            var v111 string = (v84 + v84)
                                                            switch 32767 {
                                                            case 2147483647:
                                                                var v112 int16 = -1
                                                                _ = v112
                                                            case 1:
                                                                switch fi87 {
                                                                case 1:
                                                                    if false {
                                                                        fmt.Println(v109, v93, v1)
                                                                        _ = len(v95)
                                                                        switch -9223372036854775808 {
                                                                        case -2147483648:
                                                                            var v113 bool = (v84 >= v84)
                                                                            v114 := GenS11{F0: (v97 - 167), F1: int(v108)}
                                                                            v95 = append(v95, fi87)
                                                                            _ = v113
                                                                            _ = v114
                                                                        case 2147483647:
                                                                            if true {
                                                                                var v115 int32 = (-(v108))
                                                                                var v116 uint16 = 1
                                                                                switch v84 {
                                                                                case "hello":
                                                                                    var v117 uint8 = 127
                                                                                    _ = v117
                                                                                }
                                                                                if len(v95) > 0 {
                                                                                    v95[0] = fi87
                                                                                }
                                                                                _ = v115
                                                                                _ = v116
                                                                            } else {
                                                                                v85 += (v85 - -2.25)
                                                                            }
                                                                        case 4294967295:
                                                                            v97 -= v97
                                                                        }
                                                                    } else {
                                                                        v111 = ((v111 + (((v84 + "foo bar") + v111) + "a")) + v111)
                                                                        if false {
                                                                            switch (fi87 * (fi87 >> ((18446744073709551615) & 63))) {
                                                                            case 127:
                                                                                var v118 uint = 255
                                                                                _ = v118
                                                                            default:
                                                                                v83 = v1
                                                                            }
                                                                            if (!(float32(0.0) != float32(v85))) {
                                                                                var v119 int32 = -1
                                                                                _ = v119
                                                                            }
                                                                            var v120 uint16 = 191
                                                                            _ = v120
                                                                        } else {
                                                                            _ = len(v95)
                                                                            if true {
                                                                                var v121 int8 = -10
                                                                                var v122 int = (v109 / ((fi87) | 1))
                                                                                v123 := GenS11{F0: (v97 * v97), F1: fi87}
                                                                                var v124 uint32 = uint32(v1)
                                                                                var v125 int32 = (^(v108))
                                                                                _ = v121
                                                                                _ = v122
                                                                                _ = v123
                                                                                _ = v124
                                                                                _ = v125
                                                                            } else {
                                                                                var v126 float32 = (0.1 - ((1e10 + (1.5 * float32((v85 + (v85 + (3.14159 + (v85 + float64(1e10)))))))) / 0.1))
                                                                                v83 += (v1 & (v1 % (((^(v83))) | 1)))
                                                                                _ = v126
                                                                            }
                                                                            _ = len(v95)
                                                                            switch (fi87 >> ((215) & 63)) {
                                                                            case 9223372036854775807:
                                                                                var v127 float64 = 0.0
                                                                                _ = v127
                                                                            default:
                                                                                v1 = (v83 ^ (v83 * 42))
                                                                            }
                                                                            var v128 int32 = -1
                                                                            _ = v128
                                                                        }
                                                                        v95 = append(v95, (v109 % (((-(v109))) | 1)))
                                                                        v1 = v1
                                                                        if len(v95) >= 2 {
                                                                            v95 = v95[1:len(v95)]
                                                                        }
                                                                    }
                                                                }
                                                            case -803:
                                                                var v129 float32 = (float32(0.1) / 1.5)
                                                                _ = v129
                                                            }
                                                            _ = v111
                                                        }
                                                        _ = v109
                                                    }
                                                    _ = v108
                                                }
                                            default:
                                                v1 = (v83 << ((255) & 63))
                                            }
                                            v93 = "a"
                                            _ = v97
                                        }
                                        var v130 uint32 = 42
                                        _ = v130
                                    }
                                    var v131 uint8 = v83
                                    var v132 uint32 = 63
                                    _ = v131
                                    _ = v132
                                    _ = v95
                                }
                                v133 := GenS7{F0: v93, F1: 1}
                                _ = v133
                                _ = v93
                            } else {
                                v85 += v85
                                v134 := GenS12{F0: 296, F1: 65535}
                                v84 = ""
                                v1 += v1
                                _ = v134
                            }
                        case "":
                            v135 := GenS13{F0: 811}
                            _ = v135
                        }
                    case -1:
                        if true {
                            fmt.Println(v83, fi89)
                        } else {
                            v84 = v84
                            var v136 int16 = 0
                            if true {
                                fmt.Println(v1)
                            }
                            var v137 int8 = -42
                            if false {
                                v83 += v83
                                if (v136 <= v136) {
                                    switch (fi87 + fi89) {
                                    case -291:
                                        switch "" {
                                        case "foo bar":
                                            var v138 int16 = (-(v136))
                                            _ = v138
                                        case "hello":
                                            v136 -= (v136 % (((v136 & 0)) | 1))
                                        case "\n\t":
                                            var v139 float32 = ((((float32(((1e10 - v85) - (((v85 / v85) - (1.5 + (v85 / v85))) / 1.0))) - 0.1) + (float32((3.14159 + v85)) - 1.5)) / float32((v85 / v85))) + 0.0)
                                            _ = v139
                                        }
                                    case -761:
                                        if (!false) {
                                            var v140 string = v84
                                            _ = v140
                                        } else {
                                            var v141 bool = ((v137 - v137) == v137)
                                            _ = v141
                                        }
                                    }
                                } else {
                                    var v142 int = (fi87 << ((255) & 63))
                                    delete(v3, v142)
                                    fmt.Println(v136, v137)
                                    _ = v142
                                }
                            } else {
                                if false {
                                    if ((v136 / ((int16(v137)) | 1)) <= (v136 | 1000)) {
                                        v143 := GenS6{F0: genFn0(), F1: 0}
                                        var v144 string = v84
                                        _ = v143
                                        _ = v144
                                    }
                                    var v145 float32 = float32(((0.1 / v85) - v85))
                                    _ = v145
                                }
                            }
                            if true {
                                if false {
                                    v146 := []int{}
                                    _ = v146
                                }
                            } else {
                                var v147 int8 = 38
                                if false {
                                    if false {
                                        var v148 bool = (genFn1((2 == 2)) == (v147 | -39))
                                        var v149 uint64 = 898
                                        v1 = (v1 * (v83 ^ 42))
                                        v150 := GenS14{F0: -9223372036854775808, F1: (v149 / ((v149) | 1))}
                                        _ = v148
                                        _ = v149
                                        _ = v150
                                    } else {
                                        v137 += genFn1(false)
                                    }
                                    fmt.Println(fi87)
                                    v2[(fi87 << ((4294967295) & 63))] = (false || (3.14159 <= ((1.5 - v85) * (0.1 - -2.25))))
                                } else {
                                    v151 := []int{}
                                    var v152 float32 = float32(v85)
                                    if true {
                                        v147 = v147
                                        v153 := GenS12{F0: 383, F1: 1}
                                        if len(v151) >= 2 {
                                            v151 = v151[1:len(v151)]
                                        }
                                        _ = v153
                                    }
                                    _ = v151
                                    _ = v152
                                }
                                fmt.Println(v83, v84, v1)
                                _ = v147
                            }
                            _ = v136
                            _ = v137
                        }
                    default:
                        switch v84 {
                        case "hello":
                            fmt.Println(v84, v84, fi87)
                        case "a":
                            v85 = v85
                        default:
                            if true {
                                v154 := GenS14{F0: (fi87 & 32768), F1: 18446744073709551615}
                                switch fi89 {
                                case 1:
                                    var v155 uint32 = 65535
                                    _ = v155
                                default:
                                    switch v84 {
                                    case "\n\t":
                                        v156 := GenS8{F0: false, F1: 0}
                                        _ = v156
                                    }
                                }
                                v157 := []int{1, (fi87 ^ (fi87 * (fi87 | fi87)))}
                                _ = v154
                                _ = v157
                            } else {
                                var v158 uint = 142
                                v158 = v158
                                v158 = (v158 * (v158 ^ 18446744073709551615))
                                _ = v158
                            }
                            v83 = v83
                            v83 = (v1 * 255)
                        }
                    }
                    v83 -= (v83 << ((0) & 63))
                } else {
                    var v159 uint32 = 65535
                    _ = v159
                }
                _ = fi89
            }
        }
        switch "" {
        case "\n\t":
            for fi160 := 0; fi160 < 9; fi160++ {
                if true {
                    v85 -= v85
                    var v161 int = (fi160 + fi87)
                    if true {
                        if false {
                            v162 := GenS15{F0: 362, F1: -443, F2: 1}
                            switch (^(fi87)) {
                            case 128:
                                v163 := []int{}
                                _ = v163
                            default:
                                v84 = (v84 + ((v84 + ((v84 + v84) + "hello")) + v84))
                            }
                            if false {
                                if ((!false) && (!true)) {
                                    _ = len(v3)
                                    switch 32767 {
                                    case 2147483647:
                                        v1 = (v83 + v1)
                                    case 32767:
                                        switch "hello" {
                                        case "":
                                            switch fi87 {
                                            case 128:
                                                var v164 int64 = -9223372036854775808
                                                _ = v164
                                            case 0:
                                                if ((!false) || (!(1 >= -999))) {
                                                    switch v84 {
                                                    case "":
                                                        var v165 int64 = 42
                                                        _ = v165
                                                    default:
                                                        v85 += v85
                                                    }
                                                    var v166 uint32 = 0
                                                    v167 := []int{}
                                                    _ = v166
                                                    _ = v167
                                                }
                                            default:
                                                fmt.Println(v84, fi87, fi160)
                                            }
                                        case "foo bar":
                                            v168 := GenS4{F0: 2, F1: "\n\t"}
                                            _ = v168
                                        case "hello":
                                            if false {
                                                v169 := GenS9{F0: false, F1: -2147483648}
                                                v83 = v1
                                                _ = v169
                                            }
                                        }
                                    case 1:
                                        switch (^(fi87)) {
                                        case 9223372036854775807:
                                            v1 += v1
                                        case 128:
                                            if false {
                                                fmt.Println(v161)
                                                var v170 float64 = v85
                                                switch (((("a" + ("a" + ("\n\t" + v84))) + (v84 + (v84 + v84))) + v84) + (("hello" + (("\n\t" + v84) + v84)) + "a")) {
                                                case "foo bar":
                                                    v161 = v161
                                                case "hello":
                                                    var v171 bool = true
                                                    _ = v171
                                                case "\n\t":
                                                    v172 := []int{}
                                                    _ = v172
                                                }
                                                v161 = fi160
                                                v173 := GenS16{F0: 0}
                                                _ = v170
                                                _ = v173
                                            }
                                        case 32768:
                                            if false {
                                                var v174 uint8 = v1
                                                v85 = v85
                                                _ = len(v2)
                                                _ = v174
                                            }
                                        }
                                    default:
                                        v1 = (v1 / ((0) | 1))
                                    }
                                    v161 = -9223372036854775808
                                }
                            } else {
                                if true {
                                    var v175 int32 = -468
                                    _ = v175
                                } else {
                                    v83 -= v83
                                    var v176 int64 = -916
                                    _ = v176
                                }
                                v177 := GenS3{F0: (fi160 >> ((18446744073709551615) & 63)), F1: 2, F2: -357}
                                var v178 uint16 = 0
                                _ = v177
                                _ = v178
                            }
                            v85 = (0.0 + (v85 * v85))
                            var v179 uint8 = (v1 / ((v1) | 1))
                            _ = v162
                            _ = v179
                        }
                    }
                    _ = v161
                } else {
                    v180 := []int{-996, (fi160 | fi160), fi160}
                    if true {
                        v181 := GenS10{F0: -1, F1: 272, F2: (v84 + v84)}
                        _ = v181
                    } else {
                        switch fi87 {
                        case -9223372036854775808:
                            fmt.Println(fi160, v85, fi160)
                        default:
                            if true {
                                fmt.Println(fi160)
                                v182 := GenS1{F0: (v1 >> ((4294967295) & 63))}
                                _ = v182
                            }
                        }
                        fmt.Println(v1, v84)
                        var v183 uint = 4294967295
                        _ = v183
                    }
                    _ = v180
                }
                v83 = 115
                _ = fi160
            }
        case "":
            v184 := GenS8{F0: (float64(0.0) < (((3.14159 * (v85 / 1.5)) * (v85 * ((float64(1.5) - -2.25) - float64(1.5)))) / (v85 * ((-2.25 / 3.14159) * v85)))), F1: 953}
            _ = v184
        case "a":
            var v185 uint16 = genFn0()
            _ = v185
        default:
            v186 := GenS7{F0: (v84 + ("" + "hello")), F1: 42}
            _ = v186
        }
        _ = fi87
    }
    switch 128 {
    case -1:
        v187 := []int{1, 9223372036854775807, -9223372036854775808}
        _ = v187
    }
    v188 := GenS16{F0: 42}
    if false {
        v83 += (^(v83))
    }
    v1 += v83
    var v189 uint64 = uint64(v83)
    fmt.Println("v1:", v1)
    fmt.Println("v189:", v189)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v188
    _ = v189
    _ = v2
    _ = v3
    _ = v83
    _ = v84
    _ = v85
}
