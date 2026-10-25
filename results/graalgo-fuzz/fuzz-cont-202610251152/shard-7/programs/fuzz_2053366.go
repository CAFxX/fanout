package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint16 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 int; F2 int64 }
type GenS5 struct { F0 int; F1 int16; F2 uint32 }
type GenS6 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint64; F1 int8 }
type GenS9 struct { F0 int64; F1 int64; F2 int64 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int16; F1 int32 }
type GenS13 struct { F0 uint32; F1 uint16 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int; F1 int16; F2 string }
type GenS16 struct { F0 int8; F1 int32 }
type GenS17 struct { F0 int64; F1 int; F2 uint }
type GenS18 struct { F0 uint64; F1 int }
type GenS19 struct { F0 int16; F1 bool }
type GenS20 struct { F0 string; F1 uint64; F2 int16 }
type GenS21 struct { F0 int64 }
type GenS22 struct { F0 uint64; F1 uint }
func genFn0(p0 int64, p1 int64) uint64 {
    var v7 int16 = -379
    v8 := []int{-225, 127}
    if false {
        if (false || ((-(p1)) == p1)) {
            v7 = v7
            var v9 uint16 = 255
            _ = v9
        } else {
            v7 += v7
            p1 = (^(p0))
            fmt.Println(v7)
        }
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int64) int8 {
    switch 32767 {
    case 282:
        v10 := GenS1{F0: 1, F1: p0, F2: 44}
        _ = v10
    }
    p0 -= 9223372036854775807
    switch "\n\t" {
    case "":
        for fi11 := 0; fi11 < 20; fi11++ {
            var v12 int8 = 0
            _ = fi11
            _ = v12
        }
    default:
        var v13 int32 = 2147483647
        _ = v13
    }
    if (!true) {
        switch ("foo bar" + ((("a" + (("hello" + "") + ("\n\t" + (("a" + "a") + "")))) + "\n\t") + "\n\t")) {
        case "hello":
            for fi14 := 0; fi14 < 8; fi14++ {
                v15 := GenS1{F0: -93, F1: p0, F2: 259}
                _ = fi14
                _ = v15
            }
        case "a":
            switch -360 {
            case 9223372036854775807:
                p0 = (p0 / ((526) | 1))
            case 170:
                p0 -= (p0 - p0)
            default:
                v16 := []int{32767, -967}
                _ = v16
            }
        case "\n\t":
            var v17 uint64 = 359
            _ = v17
        default:
            var v18 uint16 = 656
            _ = v18
        }
        p0 = p0
    } else {
        fmt.Println(p0)
        if (1 < 124) {
            var v19 uint = 255
            fmt.Println(p0)
            v20 := []int{-69}
            _ = v19
            _ = v20
        }
    }
    if (1 != 1) {
        fmt.Println(p0)
        if (!(1 >= 183)) {
            p0 += (p0 >> ((0) & 63))
        }
        v21 := GenS2{F0: 0, F1: p0}
        _ = v21
    }
    _ = p0
    return 127
}

func main() {
    v1 := []uint32{65535, 65535}
    var v2 int64 = 389
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]string)
    v5 := make(map[string]string)
    var v6 uint32 = 2
    if (2 == 128) {
        v22 := GenS3{F0: 1}
        v1 = append(v1, (v6 / ((4294967295) | 1)))
        _ = v22
    }
    v2 += 38
    switch "\n\t" {
    case "\n\t":
        v6 += v6
    case "a":
        v23 := v5[""]
        _ = v23
    case "foo bar":
        var v24 int = 300
        _ = v24
    }
    _ = len(v5)
    if false {
        if true {
            for fi25 := 0; fi25 < 16; fi25++ {
                switch "foo bar" {
                case "a":
                    for fi26 := 0; fi26 < 5; fi26++ {
                        if true {
                            var v27 int8 = 100
                            var v28 int8 = (v27 - (v27 << ((uint(v6)) & 63)))
                            _ = v27
                            _ = v28
                        } else {
                            var v29 int8 = 42
                            if true {
                                v30 := GenS3{F0: 1}
                                fmt.Println(v6)
                                delete(v3, 2)
                                _ = v30
                            }
                            _ = v29
                        }
                        var v31 string = "foo bar"
                        switch ("a" + v31) {
                        case "":
                            v6 = 660
                        case "\n\t":
                            var v32 uint16 = 255
                            _ = v32
                        case "hello":
                            delete(v4, (("foo bar" + "hello") + ("" + v31)))
                        }
                        _ = fi26
                        _ = v31
                    }
                case "hello":
                    switch -1 {
                    case -514:
                        for fi33 := 0; fi33 < 13; fi33++ {
                            var v34 uint16 = 65535
                            if false {
                                var v35 int = fi33
                                var v36 uint16 = v34
                                v1 = append(v1, 65535)
                                v37 := GenS2{F0: 0, F1: v2}
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v35
                                _ = v36
                                _ = v37
                            }
                            fmt.Println(v2, v6, v2)
                            v38 := GenS3{F0: 32767}
                            if false {
                                fmt.Println(v6, fi25, v2)
                                v2 -= 885
                                var v39 int64 = v2
                                var v40 uint32 = v6
                                _ = v39
                                _ = v40
                            } else {
                                switch "hello" {
                                case "hello":
                                    v6 += v6
                                case "":
                                    if false {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v2 += -1
                                    } else {
                                        fmt.Println(fi25)
                                        switch fi33 {
                                        case 712:
                                            if true {
                                                if (false || ((!(42 < 681)) || (float32(1e10) == (0.0 * 0.1)))) {
                                                    if false {
                                                        var v41 uint64 = 328
                                                        _ = v41
                                                    } else {
                                                        if (v6 > uint32(v34)) {
                                                            v6 += v6
                                                        } else {
                                                            if false {
                                                                var v42 float32 = -2.25
                                                                var v43 bool = false
                                                                _ = v42
                                                                _ = v43
                                                            }
                                                            v44 := GenS4{F0: (1 > -655), F1: (fi25 << ((18446744073709551615) & 63)), F2: (v2 >> ((uint(v34)) & 63))}
                                                            if false {
                                                                switch (fi33 << ((742) & 63)) {
                                                                case -2147483648:
                                                                    var v45 uint64 = genFn0((v2 % (((v2 / ((v2) | 1))) | 1)), 42)
                                                                    _ = v45
                                                                case 0:
                                                                    if len(v1) > 0 {
                                                                        v1[0] = 1
                                                                    }
                                                                case -1:
                                                                    v46 := GenS1{F0: 32767, F1: v2, F2: (^(v34))}
                                                                    _ = v46
                                                                default:
                                                                    if false {
                                                                        v47 := GenS5{F0: (fi25 & (fi33 << ((147) & 63))), F1: 1, F2: v6}
                                                                        v48 := GenS5{F0: 128, F1: 1, F2: (v6 + (v6 & (v6 | (v6 * v6))))}
                                                                        _ = v47
                                                                        _ = v48
                                                                    } else {
                                                                        var v49 uint32 = 79
                                                                        _ = v49
                                                                    }
                                                                }
                                                                v50 := GenS6{F0: v6, F1: (v34 ^ v34), F2: -128}
                                                                v6 += (v6 * v6)
                                                                _ = v50
                                                            }
                                                            _ = v44
                                                        }
                                                        if ((!((!(!true)) && ((32767 != 169) || true))) || (1000 > 957)) {
                                                            var v51 float64 = -2.25
                                                            _ = v51
                                                        } else {
                                                            var v52 uint64 = uint64(v34)
                                                            v53 := GenS5{F0: fi25, F1: int16(v2), F2: v6}
                                                            _ = v52
                                                            _ = v53
                                                        }
                                                    }
                                                    switch (fi33 / ((844) | 1)) {
                                                    case 32767:
                                                        v34 -= 42
                                                    case 9223372036854775807:
                                                        var v54 int8 = genFn1((v2 & -1))
                                                        _ = v54
                                                    case -1:
                                                        if (18446744073709551615 <= 4294967295) {
                                                            var v55 int16 = 0
                                                            v2 += 42
                                                            v56 := GenS7{F0: 421}
                                                            _ = v55
                                                            _ = v56
                                                        }
                                                    default:
                                                        v57 := GenS3{F0: (^(v34))}
                                                        _ = v57
                                                    }
                                                    switch ("hello" + ("a" + "foo bar")) {
                                                    case "\n\t":
                                                        v34 = (^(v34))
                                                    }
                                                }
                                                v4["hello"] = "hello"
                                                switch "" {
                                                case "\n\t":
                                                    fmt.Println(v2)
                                                case "foo bar":
                                                    v58 := GenS8{F0: 2, F1: 100}
                                                    _ = v58
                                                case "hello":
                                                    var v59 bool = false
                                                    _ = v59
                                                default:
                                                    delete(v4, "hello")
                                                }
                                                fmt.Println(fi25)
                                                v60 := GenS9{F0: v2, F1: v2, F2: (-(v2))}
                                                v61 := GenS10{F0: (^(fi33))}
                                                if ((v2 >> ((118) & 63)) > (v2 | (v2 % ((v2) | 1)))) {
                                                    var v62 int32 = 1
                                                    _ = len(v5)
                                                    var v63 uint = 4294967295
                                                    _ = v62
                                                    _ = v63
                                                } else {
                                                    _ = len(v5)
                                                }
                                                _ = v60
                                                _ = v61
                                            } else {
                                                if (739 > v2) {
                                                    if true {
                                                        _ = len(v3)
                                                    }
                                                    if (("" + ((("foo bar" + "a") + "\n\t") + "a")) == "\n\t") {
                                                        fmt.Println(v34, v34)
                                                        v1 = append(v1, (v6 % ((v6) | 1)))
                                                        var v64 uint8 = 0
                                                        fmt.Println(v64, v6)
                                                        _ = v64
                                                    }
                                                    var v65 int64 = v2
                                                    var v66 int8 = genFn1(42)
                                                    if true {
                                                        var v67 int = 2147483647
                                                        var v68 uint16 = 0
                                                        var v69 int64 = 9223372036854775807
                                                        _ = v67
                                                        _ = v68
                                                        _ = v69
                                                    }
                                                    _ = v65
                                                    _ = v66
                                                } else {
                                                    if true {
                                                        if true {
                                                            if (!(!((fi33 == 398) || ((!(((v2 | v2) != (-(v2))) && (1e10 <= 0.0))) && (0 == -60))))) {
                                                                var v70 uint = 2
                                                                _ = v70
                                                            }
                                                        }
                                                        v71 := v3[v6]
                                                        _ = v71
                                                    } else {
                                                        v2 = v2
                                                    }
                                                }
                                                v72 := GenS11{F0: 100}
                                                _ = v72
                                            }
                                        case -166:
                                            var v73 int = (^(fi33))
                                            _ = v73
                                        case 0:
                                            v34 += 65535
                                        }
                                        var v74 uint32 = (^(v6))
                                        _ = v74
                                    }
                                default:
                                    v2 += v2
                                }
                                v75 := GenS12{F0: 1000, F1: 2147483647}
                                _ = v75
                            }
                            _ = fi33
                            _ = v34
                            _ = v38
                        }
                    }
                default:
                    v6 = (v6 % ((42) | 1))
                }
                if false {
                    if true {
                        v2 -= (v2 >> ((870) & 63))
                        var v76 uint16 = 32767
                        fmt.Println(v6, fi25)
                        _ = v76
                    }
                } else {
                    if (1 != 134) {
                        _ = len(v1)
                    } else {
                        v77 := GenS13{F0: (v6 ^ (v6 - (v6 % (((v6 | v6)) | 1)))), F1: 65535}
                        v78 := GenS8{F0: 2, F1: 42}
                        v1 = append(v1, (v6 & 2147483647))
                        _ = v77
                        _ = v78
                    }
                    v79 := GenS3{F0: 32767}
                    v80 := GenS8{F0: 1, F1: -64}
                    _ = v79
                    _ = v80
                }
                var v81 uint8 = 128
                _ = fi25
                _ = v81
            }
            if true {
                switch 927 {
                case -121:
                    v2 = v2
                case 4294967295:
                    if len(v1) > 0 {
                        v1[0] = v6
                    }
                default:
                    v6 -= 2147483647
                }
                v6 += v6
                v82 := GenS14{F0: 419}
                _ = v82
            }
        }
        if true {
            v83 := GenS6{F0: 2, F1: 191, F2: 100}
            _ = v83
        } else {
            if len(v1) > 0 {
                v1[0] = v6
            }
            v6 = (v6 / ((v6) | 1))
        }
        fmt.Println(v6)
    }
    v84 := v5["\n\t"]
    if (1.5 > ((1e10 / float64(0.1)) + 0.0)) {
        if false {
            v85 := GenS15{F0: -677, F1: 1, F2: v84}
            v86 := GenS11{F0: -1}
            _ = v85
            _ = v86
        } else {
            var v87 uint8 = 0
            for fi88 := 0; fi88 < 4; fi88++ {
                var v89 int16 = 185
                for fi90 := 0; fi90 < 19; fi90++ {
                    if true {
                        v91 := GenS16{F0: 127, F1: -2147483648}
                        v2 -= -1
                        if false {
                            switch int(v6) {
                            case 42:
                                v84 = v84
                            case -260:
                                switch fi88 {
                                case 128:
                                    fmt.Println(v6, v6)
                                case -768:
                                    if ((0.0 + float64(-2.25)) > 3.14159) {
                                        v2 += v2
                                    }
                                case -40:
                                    if (1 >= 790) {
                                        v92 := GenS17{F0: (^(v2)), F1: (-(fi90)), F2: 796}
                                        _ = v92
                                    }
                                default:
                                    fmt.Println(v89)
                                }
                            case -2147483648:
                                v4[((("a" + "") + v84) + v84)] = "a"
                            }
                        }
                        _ = v91
                    }
                    _ = len(v5)
                    switch fi88 {
                    case 2147483647:
                        switch v84 {
                        case "foo bar":
                            var v93 float32 = 1e10
                            _ = v93
                        case "":
                            v94 := GenS6{F0: 2147483648, F1: 42, F2: -82}
                            _ = v94
                        case "a":
                            v95 := GenS18{F0: uint64(v87), F1: fi90}
                            _ = v95
                        }
                    case 0:
                        switch (v84 + "") {
                        case "a":
                            var v96 int = (fi88 - fi90)
                            _ = v96
                        case "hello":
                            v6 = (v6 ^ (v6 * (v6 >> ((218) & 63))))
                        }
                    case 127:
                        v5[((v84 + v84) + (v84 + "\n\t"))] = v84
                    default:
                        fmt.Println(fi88, fi90)
                    }
                    _ = fi90
                }
                _ = fi88
                _ = v89
            }
            var v97 int = 1
            _ = v87
            _ = v97
        }
    } else {
        v98 := GenS4{F0: false, F1: -2147483648, F2: (^(v2))}
        _ = v98
    }
    if true {
        switch v84 {
        case "\n\t":
            v6 = (^(v6))
        case "hello":
            v99 := GenS19{F0: -32768, F1: false}
            _ = v99
        case "a":
            var v100 uint16 = 611
            _ = v100
        }
    } else {
        fmt.Println(v84, v2)
    }
    fmt.Println(v84)
    fmt.Println(v84)
    v84 = v84
    v84 = ((v84 + ((("hello" + v84) + v84) + (((((v84 + "\n\t") + v84) + "a") + v84) + v84))) + "")
    if ((v6 << ((255) & 63)) > v6) {
        var v101 float64 = -2.25
        v2 += (v2 ^ (-(v2)))
        if len(v1) > 0 {
            v1[0] = v6
        }
        _ = v101
    } else {
        v84 = (v84 + "")
        if ((784 > 718) || (v2 <= v2)) {
            fmt.Println(v2, v84, v84)
        } else {
            if true {
                for fi102 := 0; fi102 < 5; fi102++ {
                    var v103 int = fi102
                    if false {
                        v103 = (fi102 >> ((425) & 63))
                    }
                    switch fi102 {
                    case 377:
                        switch "\n\t" {
                        case "":
                            switch fi102 {
                            case 32767:
                                v104 := v3[(v6 * v6)]
                                _ = v104
                            case -9223372036854775808:
                                v105 := GenS20{F0: ("\n\t" + v84), F1: 2, F2: -32768}
                                _ = v105
                            case 262:
                                var v106 bool = true
                                _ = v106
                            }
                        }
                    case -726:
                        switch (("a" + "") + "hello") {
                        case "\n\t":
                            v103 -= fi102
                        }
                    }
                    _ = fi102
                    _ = v103
                }
                var v107 int16 = -325
                if ((^(v6)) >= v6) {
                    v6 += 797
                    var v108 float64 = ((0.0 * 0.1) - -2.25)
                    _ = v108
                }
                _ = v107
            } else {
                switch 801 {
                case -795:
                    if true {
                        if true {
                            var v109 uint = 1
                            _ = v109
                        }
                    }
                case -2147483648:
                    if (!(false || (1.5 >= 1e10))) {
                        for fi110 := 0; fi110 < 11; fi110++ {
                            if len(v1) > 0 {
                                v1[0] = 2147483647
                            }
                            _ = fi110
                        }
                    } else {
                        v2 -= (-(v2))
                        for fi111 := 0; fi111 < 4; fi111++ {
                            for fi112 := 0; fi112 < 17; fi112++ {
                                v2 = v2
                                _ = fi112
                            }
                            _ = fi111
                        }
                        v6 += (v6 >> ((65535) & 63))
                    }
                case 4294967295:
                    if (2147483647 == 127) {
                        var v113 float32 = (1.5 / 0.1)
                        _ = v113
                    } else {
                        var v114 float32 = (float32(0.0) - float32(0.1))
                        for fi115 := 0; fi115 < 12; fi115++ {
                            v116 := GenS15{F0: fi115, F1: 376, F2: v84}
                            _ = len(v1)
                            _ = fi115
                            _ = v116
                        }
                        _ = v114
                    }
                default:
                    v2 += (v2 / (((-(v2))) | 1))
                }
                for fi117 := 0; fi117 < 4; fi117++ {
                    v118 := GenS11{F0: -1}
                    v84 = "\n\t"
                    var v119 int16 = 42
                    _ = fi117
                    _ = v118
                    _ = v119
                }
                fmt.Println(v2)
            }
            for fi120 := 0; fi120 < 3; fi120++ {
                v121 := GenS21{F0: v2}
                var v122 uint8 = 255
                _ = fi120
                _ = v121
                _ = v122
            }
            switch -891 {
            case 0:
                for fi123 := 0; fi123 < 14; fi123++ {
                    v6 = 1
                    _ = fi123
                }
            case 9223372036854775807:
                v6 -= (v6 * v6)
            default:
                v124 := GenS22{F0: 18446744073709551615, F1: 4294967295}
                _ = v124
            }
        }
        fmt.Println(v2, v6)
    }
    if (0 > 0) {
        v4[v84] = v84
        switch -239 {
        case 127:
            for fi125 := 0; fi125 < 14; fi125++ {
                v2 = 124
                v1 = append(v1, v6)
                _ = fi125
            }
        }
    } else {
        v126 := v5[("" + v84)]
        _ = v126
    }
    if ((42 >= 1) || (v84 <= "")) {
        if true {
            v2 = (v2 & (v2 | (v2 * v2)))
            v84 = "\n\t"
            v2 -= (v2 - v2)
        } else {
            var v127 uint = 210
            var v128 float64 = ((float64(0.0) - float64(0.1)) - -2.25)
            v84 = "hello"
            _ = v127
            _ = v128
        }
    } else {
        var v129 float32 = (0.0 * 0.0)
        v6 = 394
        _ = v129
    }
    fmt.Println(v2, v2)
    for fi130 := 0; fi130 < 17; fi130++ {
        for fi131 := 0; fi131 < 7; fi131++ {
            fmt.Println(fi130, fi131)
            _ = fi131
        }
        switch fi130 {
        case 9223372036854775807:
            if len(v1) > 0 {
                v1[0] = v6
            }
        case 4294967295:
            var v132 int8 = 127
            _ = v132
        }
        for fi133 := 0; fi133 < 14; fi133++ {
            v6 = v6
            _ = fi133
        }
        _ = fi130
    }
    v3[(v6 / ((v6) | 1))] = (v6 / ((0) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v84
}
