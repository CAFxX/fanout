package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS5 struct { F0 int32; F1 int64; F2 int8 }
type GenS6 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int64; F1 uint }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS12 struct { F0 int; F1 int64 }
type GenS13 struct { F0 int16; F1 int16 }
type GenS14 struct { F0 int; F1 uint64; F2 int16 }
type GenS15 struct { F0 int; F1 uint8; F2 uint8 }
func genFn0(p0 int16, p1 uint, p2 uint64) int64 {
    v7 := GenS1{F0: p0, F1: 42, F2: 2147483648}
    p0 = p0
    p1 = 18446744073709551615
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 float64 = 0.0
        p0 = p0
        v10 := []int{}
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -809
}
func genFn1(p0 int8) uint {
    fmt.Println(p0)
    if true {
        var v11 int32 = -210
        _ = v11
    } else {
        if true {
            var v12 uint = 4294967295
            _ = v12
        } else {
            v13 := []int{9223372036854775807, -2147483648, 2147483647}
            v14 := GenS1{F0: 1, F1: 168, F2: 1}
            fmt.Println(p0)
            _ = v13
            _ = v14
        }
        p0 = 127
    }
    fmt.Println(p0)
    _ = p0
    return 255
}

func main() {
    var v1 uint32 = 65535
    var v2 string = "foo bar"
    var v3 int16 = 0
    v4 := []uint32{v1, (v1 - 0), (v1 >> ((uint(v1)) & 63))}
    var v5 int16 = v3
    var v6 uint64 = 827
    if (60 >= -1) {
        var v15 bool = (true && (-1 >= -872))
        _ = v15
    }
    switch 49 {
    case 74:
        if ((v1 / (((^(v1))) | 1)) >= (v1 % (((v1 & (v1 + v1))) | 1))) {
            v4 = append(v4, 786)
            var v16 uint16 = uint16(v6)
            _ = v16
        } else {
            v4 = append(v4, uint32(v6))
        }
    case 127:
        if true {
            switch 0 {
            case 9223372036854775807:
                v17 := GenS2{F0: -9223372036854775808, F1: -1}
                _ = v17
            case -196:
                v4 = append(v4, v1)
            }
        } else {
            _ = len(v4)
            v3 = (v3 << ((764) & 63))
            var v18 int16 = v5
            _ = v18
        }
    }
    var v19 string = "\n\t"
    for fi20 := 0; fi20 < 12; fi20++ {
        for fi21 := 0; fi21 < 7; fi21++ {
            v22 := GenS1{F0: -1, F1: 212, F2: 2147483648}
            _ = fi21
            _ = v22
        }
        switch ("hello" + v2) {
        case "":
            v3 += (^(v3))
        default:
            v23 := GenS1{F0: -307, F1: 216, F2: (v1 - (v1 ^ v1))}
            _ = v23
        }
        v24 := GenS3{F0: -1}
        _ = fi20
        _ = v24
    }
    v6 += (^(v6))
    v2 = ("\n\t" + ((v19 + v19) + "\n\t"))
    _ = len(v4)
    v25 := GenS3{F0: 42}
    for fi26 := 0; fi26 < 2; fi26++ {
        for fi27 := 0; fi27 < 9; fi27++ {
            var v28 int16 = (-(v5))
            _ = fi27
            _ = v28
        }
        switch v19 {
        case "a":
            if (((false || true) && ((v19 + "") >= (v19 + v19))) && true) {
                v4 = append(v4, (v1 << ((42) & 63)))
                for fi29 := 0; fi29 < 7; fi29++ {
                    var v30 int = fi29
                    if false {
                        var v31 float64 = 1.5
                        if ((42 <= 367) && (2 >= 0)) {
                            var v32 float32 = 1e10
                            switch "a" {
                            case "hello":
                                var v33 bool = true
                                _ = v33
                            case "foo bar":
                                if false {
                                    v3 += v3
                                    if false {
                                        v3 = v3
                                        switch "hello" {
                                        case "hello":
                                            switch v2 {
                                            case "foo bar":
                                                v34 := GenS2{F0: fi26, F1: 42}
                                                _ = v34
                                            case "hello":
                                                if (-128 <= 104) {
                                                    var v35 uint16 = 65535
                                                    if len(v4) > 0 {
                                                        v4[0] = v1
                                                    }
                                                    v5 = v3
                                                    _ = v35
                                                }
                                            case "":
                                                v3 += v5
                                            }
                                        default:
                                            if len(v4) > 0 {
                                                v4[0] = v1
                                            }
                                            var v36 uint64 = (^(v6))
                                            v37 := GenS4{F0: -2147483648, F1: (^(v1)), F2: 1}
                                            _ = v36
                                            _ = v37
                                        }
                                        var v38 int8 = -95
                                        _ = v38
                                    } else {
                                        switch v30 {
                                        case 32767:
                                            fmt.Println(v2, v1)
                                        case 32768:
                                            if false {
                                                var v39 int16 = (v3 | (v5 & v3))
                                                _ = v39
                                            }
                                        case 4294967295:
                                            switch "hello" {
                                            case "\n\t":
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            case "a":
                                                v31 = ((-2.25 + 3.14159) - 1.5)
                                            case "hello":
                                                fmt.Println(fi29)
                                            }
                                            switch "\n\t" {
                                            case "foo bar":
                                                v1 += 2147483648
                                            case "hello":
                                                var v40 int32 = -521
                                                _ = v40
                                            default:
                                                var v41 float64 = v31
                                                _ = v41
                                            }
                                        }
                                    }
                                    if true {
                                        var v42 uint32 = (v1 & (v1 - (v1 | (v1 | 2))))
                                        v43 := GenS2{F0: int(v6), F1: 9223372036854775807}
                                        _ = v42
                                        _ = v43
                                    }
                                    var v44 string = "hello"
                                    switch (v30 + v30) {
                                    case 127:
                                        switch fi26 {
                                        case 4:
                                            var v45 string = v2
                                            _ = v45
                                        case -537:
                                            v44 = v19
                                        case 32767:
                                            fmt.Println(v44, v19, v1)
                                        }
                                    case -9223372036854775808:
                                        v46 := GenS3{F0: 42}
                                        _ = v46
                                    case 32767:
                                        var v47 bool = false
                                        _ = v47
                                    default:
                                        v3 = -1
                                    }
                                    _ = v44
                                } else {
                                    var v48 int32 = 0
                                    if len(v4) > 0 {
                                        v4[0] = 4294967295
                                    }
                                    _ = v48
                                }
                            default:
                                v5 = (v3 ^ (v5 << ((genFn1(-128)) & 63)))
                            }
                            v6 -= 42
                            _ = v32
                        }
                        v31 -= (((v31 - ((v31 + v31) * 1.5)) + v31) + float64((float32(0.1) * 1.5)))
                        v31 = v31
                        _ = v31
                    }
                    switch ("a" + ("\n\t" + ((((v2 + (v2 + (v19 + ("\n\t" + (v2 + ""))))) + v19) + "\n\t") + "foo bar"))) {
                    case "\n\t":
                        switch v19 {
                        case "":
                            fmt.Println(fi26, fi29)
                        case "a":
                            if false {
                                v30 += fi29
                            } else {
                                switch (v30 ^ fi26) {
                                case 702:
                                    _ = len(v4)
                                case -2147483648:
                                    v49 := GenS4{F0: 42, F1: (v1 / ((v1) | 1)), F2: v6}
                                    _ = v49
                                }
                            }
                        default:
                            v50 := GenS1{F0: -156, F1: 35, F2: (v1 << ((960) & 63))}
                            _ = v50
                        }
                    case "a":
                        if (false || (!(1 != 255))) {
                            var v51 uint32 = v1
                            if len(v4) > 0 {
                                v4[0] = (v51 + 2147483648)
                            }
                            _ = v51
                        } else {
                            v52 := GenS5{F0: 1, F1: 486, F2: 42}
                            var v53 uint8 = 58
                            _ = v52
                            _ = v53
                        }
                    }
                    _ = fi29
                    _ = v30
                }
                v54 := GenS5{F0: -2147483648, F1: -2, F2: 106}
                _ = v54
            } else {
                for fi55 := 0; fi55 < 6; fi55++ {
                    var v56 int8 = 100
                    v1 = v1
                    _ = fi55
                    _ = v56
                }
                fmt.Println(v19, v5)
            }
        default:
            for fi57 := 0; fi57 < 1; fi57++ {
                var v58 uint64 = (v6 | 18446744073709551615)
                _ = fi57
                _ = v58
            }
        }
        var v59 uint32 = 0
        _ = fi26
        _ = v59
    }
    fmt.Println(v6, v3, v2)
    v60 := GenS1{F0: (v3 & v5), F1: 224, F2: (v1 & (v1 & (v1 - (v1 | (^(v1))))))}
    fmt.Println(v6, v6)
    v3 += -1
    v61 := GenS4{F0: -2147483648, F1: v1, F2: (v6 & v6)}
    for fi62 := 0; fi62 < 2; fi62++ {
        for fi63 := 0; fi63 < 11; fi63++ {
            v3 += v3
            if (!(!false)) {
                if ((v19 == v19) || (false && (!false))) {
                    var v64 string = ((v19 + "") + (v2 + v2))
                    var v65 float64 = 3.14159
                    var v66 uint64 = (v6 << ((65535) & 63))
                    var v67 int8 = 1
                    v68 := GenS2{F0: (fi62 & fi62), F1: 170}
                    var v69 float32 = 3.14159
                    _ = v64
                    _ = v65
                    _ = v66
                    _ = v67
                    _ = v68
                    _ = v69
                }
                switch -9223372036854775808 {
                case 785:
                    if false {
                        if (!((34 < 71) && (0.1 >= 3.14159))) {
                            v6 = 923
                            if (127 == 42) {
                                var v70 float32 = -2.25
                                _ = v70
                            } else {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                if true {
                                    var v71 int64 = 22
                                    if (2 > 42) {
                                        var v72 uint64 = (v6 / ((v6) | 1))
                                        _ = v72
                                    }
                                    switch (fi63 & fi63) {
                                    case 42:
                                        var v73 int8 = 42
                                        _ = v73
                                    case 32767:
                                        var v74 bool = ((true && true) || false)
                                        _ = v74
                                    case 2147483647:
                                        if (-2147483648 < -2147483648) {
                                            switch v19 {
                                            case "hello":
                                                v4 = append(v4, (v1 / (((v1 ^ v1)) | 1)))
                                            default:
                                                fmt.Println(v19)
                                            }
                                        }
                                        v75 := GenS6{F0: uint32(v6), F1: v1, F2: (v6 << ((uint(v6)) & 63))}
                                        fmt.Println(v5, v5, v6)
                                        _ = v75
                                    }
                                    var v76 bool = true
                                    v77 := GenS7{F0: (-(v3))}
                                    _ = v71
                                    _ = v76
                                    _ = v77
                                } else {
                                    fmt.Println(v5, v1, v1)
                                    if (127 <= 42) {
                                        if false {
                                            v78 := GenS8{F0: 9223372036854775807, F1: 42}
                                            var v79 uint16 = 2
                                            var v80 bool = (!false)
                                            if len(v4) > 0 {
                                                v4[0] = (v1 << ((1) & 63))
                                            }
                                            _ = v78
                                            _ = v79
                                            _ = v80
                                        } else {
                                            if true {
                                                _ = len(v4)
                                                var v81 uint64 = (^(v6))
                                                _ = v81
                                            }
                                            var v82 string = v19
                                            if (!true) {
                                                fmt.Println(v2, v3)
                                            }
                                            v82 = ((v19 + v82) + (v2 + (v19 + ("a" + v2))))
                                            switch fi62 {
                                            case 706:
                                                var v83 string = "foo bar"
                                                _ = v83
                                            }
                                            var v84 int8 = 100
                                            if true {
                                                switch ("foo bar" + (v2 + "hello")) {
                                                case "a":
                                                    v1 -= v1
                                                case "foo bar":
                                                    switch (fi62 / ((fi63) | 1)) {
                                                    case 364:
                                                        var v85 float64 = 0.1
                                                        _ = v85
                                                    }
                                                }
                                                var v86 int8 = 127
                                                fmt.Println(v84)
                                                v87 := GenS9{F0: 42}
                                                _ = v86
                                                _ = v87
                                            } else {
                                                if len(v4) > 0 {
                                                    v4[0] = (v1 << ((255) & 63))
                                                }
                                            }
                                            var v88 uint8 = 0
                                            _ = v82
                                            _ = v84
                                            _ = v88
                                        }
                                        v4 = append(v4, v1)
                                        var v89 int32 = -29
                                        fmt.Println(v19)
                                        _ = v89
                                    }
                                    var v90 int = (fi63 / (((fi63 / ((42) | 1))) | 1))
                                    _ = v90
                                }
                                var v91 float64 = (float64(float32(0.0)) - 1.5)
                                _ = v91
                            }
                        } else {
                            v92 := GenS1{F0: -32768, F1: 128, F2: v1}
                            if true {
                                v93 := GenS5{F0: 1, F1: 1, F2: 127}
                                if (true && true) {
                                    var v94 string = (((v2 + "\n\t") + v19) + "\n\t")
                                    if true {
                                        var v95 float32 = 1.5
                                        _ = v95
                                    }
                                    _ = len(v4)
                                    _ = v94
                                }
                                _ = v93
                            } else {
                                v1 -= (v1 & (^(v1)))
                            }
                            if true {
                                var v96 uint = 42
                                switch (("foo bar" + v19) + ("hello" + v19)) {
                                case "foo bar":
                                    v97 := GenS7{F0: v3}
                                    _ = v97
                                case "":
                                    v1 += (v1 ^ v1)
                                default:
                                    var v98 uint64 = (v6 ^ (v6 - v6))
                                    _ = v98
                                }
                                _ = v96
                            } else {
                                v99 := GenS7{F0: (v5 & 42)}
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v99
                            }
                            _ = v92
                        }
                        fmt.Println(fi62, v1)
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v5 = (^(v5))
                        if true {
                            if true {
                                switch 2147483647 {
                                case 127:
                                    var v100 float32 = 0.1
                                    _ = v100
                                case 4294967295:
                                    v19 = v2
                                case -9223372036854775808:
                                    if len(v4) > 0 {
                                        v4[0] = v1
                                    }
                                default:
                                    fmt.Println(v19)
                                }
                                v101 := GenS3{F0: -970}
                                v102 := GenS2{F0: -743, F1: 42}
                                _ = v101
                                _ = v102
                            } else {
                                v2 = "foo bar"
                            }
                            if false {
                                v3 = (v3 & (v5 / ((586) | 1)))
                            }
                            fmt.Println(v3, v6)
                            var v103 float64 = 0.1
                            _ = v103
                        }
                    }
                }
                var v104 uint64 = 2
                fmt.Println(v6)
                switch "\n\t" {
                case "\n\t":
                    switch v2 {
                    case "":
                        fmt.Println(fi62, fi62)
                    }
                case "hello":
                    fmt.Println(v19, v6)
                case "":
                    switch "\n\t" {
                    case "foo bar":
                        if false {
                            v104 = v104
                            var v105 string = ((v19 + v2) + ("" + (((("" + v19) + "a") + "hello") + v2)))
                            var v106 int64 = 9223372036854775807
                            _ = v105
                            _ = v106
                        } else {
                            if (9223372036854775807 > genFn0(1000, 1, (v6 & v104))) {
                                if true {
                                    if ((2 <= 255) && (true && true)) {
                                        if (42 == 712) {
                                            var v107 bool = true
                                            v108 := GenS3{F0: 0}
                                            _ = v107
                                            _ = v108
                                        }
                                        v104 += (v6 ^ (v104 | v104))
                                    } else {
                                        fmt.Println(v1)
                                        v19 = v2
                                        if false {
                                            v5 = -32768
                                            switch ("a" + "\n\t") {
                                            case "foo bar":
                                                if false {
                                                    if true {
                                                        v109 := GenS10{F0: 63}
                                                        _ = v109
                                                    } else {
                                                        v6 -= (v6 ^ (v104 >> ((18446744073709551615) & 63)))
                                                    }
                                                }
                                            default:
                                                v110 := GenS11{F0: v6, F1: 32767, F2: 0}
                                                _ = v110
                                            }
                                            v2 = (("" + ("a" + "hello")) + ("foo bar" + v2))
                                        }
                                    }
                                    fmt.Println(v19)
                                }
                            }
                            if true {
                                fmt.Println(fi63)
                                var v111 uint = 799
                                var v112 int8 = 100
                                _ = v111
                                _ = v112
                            } else {
                                if false {
                                    v6 -= v6
                                }
                            }
                            var v113 string = v19
                            _ = v113
                        }
                    }
                default:
                    switch ((v19 + v19) + v2) {
                    case "foo bar":
                        var v114 uint32 = v1
                        _ = v114
                    default:
                        var v115 bool = (100 >= 127)
                        _ = v115
                    }
                }
                if (false && (!(!(0.0 < (1e10 / (-2.25 - 0.0)))))) {
                    _ = len(v4)
                    v116 := GenS4{F0: -2147483648, F1: v1, F2: (^(v104))}
                    _ = v116
                }
                if ((v19 != ("\n\t" + "hello")) || (-1 <= -1)) {
                    v104 += v6
                    fmt.Println(v5, v2, fi62)
                    fmt.Println(v5)
                }
                if ((v1 & 2147483648) == (^(v1))) {
                    if true {
                        switch -875 {
                        case -188:
                            switch int(v1) {
                            case 0:
                                switch ("hello" + v2) {
                                case "hello":
                                    var v117 float64 = ((3.14159 * 1e10) / 1e10)
                                    var v118 int32 = 1
                                    switch "a" {
                                    case "\n\t":
                                        v6 = v6
                                    case "a":
                                        v6 = (v6 & v6)
                                    case "foo bar":
                                        if len(v4) > 0 {
                                            v4[0] = (^(v1))
                                        }
                                    }
                                    _ = v117
                                    _ = v118
                                }
                                switch (-(fi63)) {
                                case -9223372036854775808:
                                    var v119 string = v2
                                    _ = v119
                                case 32767:
                                    if false {
                                        v120 := GenS3{F0: genFn0(v3, 2, (v104 + 363))}
                                        v121 := GenS5{F0: 2147483647, F1: -1, F2: -1}
                                        v5 = v5
                                        _ = v120
                                        _ = v121
                                    }
                                case 128:
                                    if ((v1 - (v1 >> ((42) & 63))) == (v1 << ((42) & 63))) {
                                        v2 = ("a" + v19)
                                        v122 := GenS7{F0: 42}
                                        switch v19 {
                                        case "foo bar":
                                            var v123 int = fi63
                                            _ = v123
                                        case "\n\t":
                                            if true {
                                                fmt.Println(v19)
                                                var v124 int32 = 42
                                                switch fi63 {
                                                case 32768:
                                                    if (!true) {
                                                        fmt.Println(v2, fi62, fi62)
                                                        v4 = append(v4, v1)
                                                        v125 := GenS12{F0: (fi62 >> ((18446744073709551615) & 63)), F1: 980}
                                                        _ = v125
                                                    } else {
                                                        v126 := GenS12{F0: (^(fi62)), F1: 0}
                                                        v1 -= (v1 | (v1 % ((v1) | 1)))
                                                        _ = v126
                                                    }
                                                case -9223372036854775808:
                                                    var v127 float32 = 1.5
                                                    _ = v127
                                                case 162:
                                                    v4 = append(v4, (v1 - 8))
                                                default:
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                }
                                                _ = v124
                                            } else {
                                                v128 := GenS5{F0: 2147483647, F1: 1, F2: 42}
                                                if (genFn1(-117) == genFn1(57)) {
                                                    v4 = append(v4, 0)
                                                    var v129 uint16 = 255
                                                    if true {
                                                        v1 = v1
                                                        var v130 uint64 = 889
                                                        _ = v130
                                                    }
                                                    fmt.Println(v3, v129, v104)
                                                    v131 := GenS1{F0: v5, F1: 42, F2: (v1 * v1)}
                                                    _ = v129
                                                    _ = v131
                                                }
                                                _ = v128
                                            }
                                        case "":
                                            if ((-88 <= 127) && (!(!(1 < 582)))) {
                                                var v132 int32 = 1
                                                _ = len(v4)
                                                var v133 float32 = ((3.14159 * 0.0) + -2.25)
                                                _ = v132
                                                _ = v133
                                            }
                                            var v134 int8 = 1
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            _ = v134
                                        }
                                        _ = v122
                                    } else {
                                        switch -9223372036854775808 {
                                        case -51:
                                            fmt.Println(fi62, v104, fi62)
                                        case 127:
                                            v6 = 4294967295
                                        }
                                    }
                                default:
                                    var v135 uint8 = 128
                                    _ = v135
                                }
                            case 9223372036854775807:
                                v1 += v1
                            case 942:
                                var v136 uint32 = 1
                                _ = v136
                            default:
                                var v137 string = v19
                                _ = v137
                            }
                        default:
                            if len(v4) > 0 {
                                v4[0] = 2147483647
                            }
                        }
                        if (((!true) || (1000 > (-(v3)))) && false) {
                            v138 := GenS6{F0: 813, F1: 91, F2: (v104 - v6)}
                            v2 = "hello"
                            if (0 < -784) {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                if (fi62 >= fi62) {
                                    if true {
                                        v6 -= (v6 << ((42) & 63))
                                        if false {
                                            if (386 >= -9223372036854775808) {
                                                v139 := GenS1{F0: (v3 % ((v3) | 1)), F1: 2, F2: v1}
                                                var v140 int = fi62
                                                _ = len(v4)
                                                if false {
                                                    v5 = v5
                                                }
                                                _ = v139
                                                _ = v140
                                            }
                                            fmt.Println(v19, v2, fi62)
                                            switch (v19 + ("hello" + "\n\t")) {
                                            case "foo bar":
                                                switch (v2 + v2) {
                                                case "foo bar":
                                                    var v141 float64 = 0.0
                                                    _ = v141
                                                case "hello":
                                                    v4 = append(v4, v1)
                                                }
                                            }
                                            v142 := GenS13{F0: v5, F1: int16(fi63)}
                                            var v143 string = ("foo bar" + v2)
                                            _ = v142
                                            _ = v143
                                        } else {
                                            var v144 string = v19
                                            _ = v144
                                        }
                                    }
                                    fmt.Println(v1, v19, v6)
                                    if (false || (!((1 >= 0) || true))) {
                                        if true {
                                            fmt.Println(fi63, v1)
                                            var v145 int = (fi62 / ((-859) | 1))
                                            var v146 uint64 = v6
                                            _ = v145
                                            _ = v146
                                        }
                                        fmt.Println(v19, fi63, v1)
                                        var v147 bool = (float64(float32(1.5)) <= (0.0 * 3.14159))
                                        _ = v147
                                    }
                                } else {
                                    v148 := GenS14{F0: -2147483648, F1: 772, F2: (v3 >> ((2) & 63))}
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    fmt.Println(v1, v1)
                                    _ = v148
                                }
                                if len(v4) > 0 {
                                    v4[0] = (v1 >> ((760) & 63))
                                }
                                v149 := GenS9{F0: 32767}
                                _ = v149
                            }
                            _ = v138
                        } else {
                            v150 := GenS8{F0: -459, F1: genFn1(14)}
                            _ = v150
                        }
                        var v151 float32 = float32(0.1)
                        var v152 int16 = (v3 & (v5 + v5))
                        _ = v151
                        _ = v152
                    }
                }
                _ = v104
            }
            v1 -= (v1 - (v1 - v1))
            _ = fi63
        }
        _ = fi62
    }
    v153 := GenS2{F0: 32767, F1: 0}
    for fi154 := 0; fi154 < 9; fi154++ {
        switch v19 {
        case "foo bar":
            if false {
                var v155 int8 = 0
                var v156 uint16 = uint16(v6)
                var v157 int8 = v155
                _ = v155
                _ = v156
                _ = v157
            }
        case "":
            v158 := GenS15{F0: fi154, F1: 75, F2: 91}
            _ = v158
        case "a":
            if len(v4) > 0 {
                v4[0] = (v1 << ((18446744073709551615) & 63))
            }
        }
        if (65535 < 32767) {
            v2 = v2
        }
        _ = fi154
    }
    var v159 uint64 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v159:", v159)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v153
    _ = v159
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v61
}
