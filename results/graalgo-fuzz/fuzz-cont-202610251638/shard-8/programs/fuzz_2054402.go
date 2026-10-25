package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int32 }
type GenS2 struct { F0 uint8; F1 int; F2 uint8 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 int16; F1 int64; F2 int8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint16; F1 uint16; F2 bool }
type GenS7 struct { F0 uint64; F1 int32 }
type GenS8 struct { F0 uint32; F1 int }
type GenS9 struct { F0 uint64; F1 bool }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint64; F1 int8; F2 int16 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 uint32; F1 uint32 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 uint64; F1 int16 }
type GenS17 struct { F0 uint64; F1 int32 }
type GenS18 struct { F0 string; F1 uint32 }
func genFn0(p0 int, p1 uint8, p2 int32) int64 {
    v5 := GenS1{F0: 580, F1: 42, F2: p2}
    var v6 float64 = -2.25
    v6 = float64(1e10)
    for fi7 := 0; fi7 < 3; fi7++ {
        v6 += (v6 + 1.5)
        _ = fi7
    }
    v6 += -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -1
}
func genFn1(p0 uint8) int8 {
    p0 += (p0 + (p0 & p0))
    for fi8 := 0; fi8 < 1; fi8++ {
        if (false || ((false && true) || true)) {
            v9 := GenS2{F0: p0, F1: 2147483647, F2: p0}
            _ = v9
        }
        p0 = p0
        v10 := []int{fi8}
        _ = fi8
        _ = v10
    }
    if true {
        for fi11 := 0; fi11 < 14; fi11++ {
            p0 += (p0 + p0)
            var v12 int64 = 9223372036854775807
            v13 := []int{}
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    p0 -= p0
    p0 = (p0 - p0)
    _ = p0
    return -128
}

func main() {
    v1 := []int{}
    var v2 float32 = (((float32(0.0) - 0.0) - float32((1e10 * 1.5))) / (1e10 / 1e10))
    var v3 int8 = -18
    v4 := []int{2147483647}
    var v14 int8 = 89
    if true {
        if true {
            for fi15 := 0; fi15 < 8; fi15++ {
                v2 -= (v2 / (v2 * 0.1))
                _ = fi15
            }
            var v16 uint32 = 2
            _ = v16
        } else {
            if (true || ("a" == "foo bar")) {
                if (!true) {
                    var v17 uint16 = 0
                    _ = v17
                }
                for fi18 := 0; fi18 < 3; fi18++ {
                    for fi19 := 0; fi19 < 6; fi19++ {
                        if true {
                            switch (fi19 << ((950) & 63)) {
                            case 32768:
                                v20 := GenS2{F0: 255, F1: -1, F2: 21}
                                _ = v20
                            case 2147483647:
                                v21 := GenS2{F0: 76, F1: (-(fi18)), F2: 152}
                                _ = v21
                            }
                            if false {
                                if true {
                                    switch "foo bar" {
                                    case "":
                                        v22 := GenS3{F0: 2, F1: -2147483648}
                                        _ = v22
                                    case "hello":
                                        v23 := GenS4{F0: 0, F1: genFn0(9223372036854775807, 2, 2147483647), F2: (v14 - genFn1(2))}
                                        _ = v23
                                    }
                                    switch ("foo bar" + "") {
                                    case "a":
                                        v24 := GenS1{F0: 0, F1: genFn1(127), F2: 121}
                                        if (916 == 74) {
                                            v2 += 0.1
                                            v25 := GenS2{F0: 47, F1: fi18, F2: 1}
                                            switch (fi19 + (fi19 >> ((815) & 63))) {
                                            case 102:
                                                v3 -= v14
                                            default:
                                                v4 = append(v4, 0)
                                            }
                                            var v26 uint64 = 1
                                            _ = v25
                                            _ = v26
                                        } else {
                                            var v27 int16 = 42
                                            v14 = v14
                                            switch "\n\t" {
                                            case "":
                                                v28 := GenS4{F0: 117, F1: 908, F2: v14}
                                                _ = v28
                                            case "hello":
                                                switch fi18 {
                                                case 9223372036854775807:
                                                    v2 += float32(1e10)
                                                case 153:
                                                    v2 = v2
                                                case -735:
                                                    v29 := GenS1{F0: -1, F1: (v3 / ((genFn1(128)) | 1)), F2: 0}
                                                    _ = v29
                                                }
                                            case "a":
                                                v27 += 32767
                                            default:
                                                var v30 uint = 1
                                                _ = v30
                                            }
                                            _ = v27
                                        }
                                        _ = v24
                                    case "":
                                        var v31 float32 = (v2 / float32(3.14159))
                                        _ = v31
                                    case "\n\t":
                                        switch fi19 {
                                        case 32767:
                                            var v32 int16 = 1000
                                            _ = v32
                                        case -348:
                                            if len(v4) > 0 {
                                                v4[0] = fi18
                                            }
                                        default:
                                            switch (fi19 * (fi19 * (fi19 | fi19))) {
                                            case -964:
                                                var v33 uint16 = 54
                                                _ = v33
                                            case 9223372036854775807:
                                                var v34 uint16 = 265
                                                _ = v34
                                            case 0:
                                                if false {
                                                    v35 := GenS5{F0: 65535}
                                                    if true {
                                                        v14 += -23
                                                    } else {
                                                        if true {
                                                            switch (fi18 & 0) {
                                                            case 2147483647:
                                                                v1 = append(v1, fi19)
                                                            case -9223372036854775808:
                                                                if (false || true) {
                                                                    v3 = (v14 >> ((42) & 63))
                                                                } else {
                                                                    v2 = v2
                                                                    if true {
                                                                        fmt.Println(v14, fi19)
                                                                        if true {
                                                                            fmt.Println(fi18, fi18, v2)
                                                                            if true {
                                                                                var v36 int32 = -1
                                                                                fmt.Println(v14, v2, v3)
                                                                                v37 := GenS6{F0: 1, F1: 2, F2: (("foo bar" + "hello") == "a")}
                                                                                _ = v36
                                                                                _ = v37
                                                                            }
                                                                            var v38 float64 = (-2.25 * 0.1)
                                                                            v39 := GenS3{F0: 119, F1: 135}
                                                                            _ = v38
                                                                            _ = v39
                                                                        }
                                                                        var v40 int16 = 32767
                                                                        _ = v40
                                                                    }
                                                                }
                                                            case 9223372036854775807:
                                                                var v41 uint = 42
                                                                _ = v41
                                                            default:
                                                                v14 -= 127
                                                            }
                                                            var v42 uint16 = 65535
                                                            _ = v42
                                                        }
                                                        v43 := GenS7{F0: 18446744073709551615, F1: -2147483648}
                                                        if (!true) {
                                                            switch fi18 {
                                                            case 0:
                                                                v3 -= (-(v14))
                                                            case -9223372036854775808:
                                                                v14 -= (v14 / (((v3 ^ (v14 % ((v14) | 1)))) | 1))
                                                            case 639:
                                                                v2 += float32(-2.25)
                                                            }
                                                            switch -461 {
                                                            case -9223372036854775808:
                                                                if false {
                                                                    if (!(!(-1 == -138))) {
                                                                        v2 = 0.0
                                                                    } else {
                                                                        switch 0 {
                                                                        case 184:
                                                                            v14 = v3
                                                                        case 4294967295:
                                                                            v2 = (v2 / 1.5)
                                                                        }
                                                                        if false {
                                                                            if (false && (0.1 > 0.1)) {
                                                                                fmt.Println(v14, fi19)
                                                                                fmt.Println(v2)
                                                                            } else {
                                                                                v44 := GenS6{F0: 65535, F1: 65535, F2: (0.0 != 1e10)}
                                                                                switch "\n\t" {
                                                                                case "":
                                                                                    v2 = v2
                                                                                    v3 = 42
                                                                                    v14 += v3
                                                                                case "foo bar":
                                                                                    fmt.Println(v3, v14)
                                                                                default:
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = (-(fi19))
                                                                                    }
                                                                                    if len(v4) >= 2 {
                                                                                        v4 = v4[1:len(v4)]
                                                                                    }
                                                                                    var v45 float64 = (float64(v2) + (float64(float32((3.14159 - -2.25))) + 3.14159))
                                                                                    _ = v45
                                                                                }
                                                                                _ = v44
                                                                            }
                                                                            var v46 uint16 = 778
                                                                            v47 := GenS7{F0: 977, F1: -1}
                                                                            _ = v46
                                                                            _ = v47
                                                                        } else {
                                                                            var v48 uint8 = 255
                                                                            v3 = v3
                                                                            _ = v48
                                                                        }
                                                                        v49 := GenS8{F0: 0, F1: fi19}
                                                                        _ = v49
                                                                    }
                                                                    if (42 <= 2147483647) {
                                                                        v14 -= v3
                                                                        v14 += genFn1(170)
                                                                    }
                                                                    fmt.Println(v2)
                                                                } else {
                                                                    switch -523 {
                                                                    case 551:
                                                                        var v50 uint8 = 1
                                                                        _ = v50
                                                                    case 0:
                                                                        if ((4294967295 > -2147483648) || (42 <= 2147483648)) {
                                                                            if (!(1000 <= 311)) {
                                                                                _ = len(v1)
                                                                            }
                                                                            var v51 uint32 = 581
                                                                            _ = v51
                                                                        }
                                                                    case -9223372036854775808:
                                                                        fmt.Println(fi18, v2, v3)
                                                                    }
                                                                }
                                                            case 0:
                                                                var v52 int8 = (v3 ^ 127)
                                                                _ = v52
                                                            case 32768:
                                                                var v53 int16 = int16(fi18)
                                                                _ = v53
                                                            default:
                                                                if false {
                                                                    var v54 int8 = v14
                                                                    if true {
                                                                        if true {
                                                                            var v55 int8 = (v3 | (v3 * (^(v54))))
                                                                            v56 := GenS9{F0: 4294967295, F1: ((fi18 << ((1) & 63)) >= (fi19 & 2147483647))}
                                                                            _ = v55
                                                                            _ = v56
                                                                        }
                                                                        v57 := GenS10{F0: 543}
                                                                        switch ("foo bar" + "foo bar") {
                                                                        case "foo bar":
                                                                            var v58 uint32 = 0
                                                                            v58 -= v58
                                                                            if true {
                                                                                var v59 uint = 738
                                                                                switch "" {
                                                                                case "foo bar":
                                                                                    var v60 int32 = 2147483647
                                                                                    _ = v60
                                                                                case "\n\t":
                                                                                    if (true || (2 > 18446744073709551615)) {
                                                                                        if true {
                                                                                            fmt.Println(v14, v3, v2)
                                                                                            if (v3 < int8(fi19)) {
                                                                                                v59 -= (v59 >> ((v59) & 63))
                                                                                                v61 := GenS11{F0: 4294967295, F1: 127, F2: 42}
                                                                                                fmt.Println(v2, fi19)
                                                                                                _ = v61
                                                                                            }
                                                                                            v3 -= -128
                                                                                        }
                                                                                        if (42 > 499) {
                                                                                            fmt.Println(v3)
                                                                                            var v62 uint = (^(v59))
                                                                                            v2 = float32(0.1)
                                                                                            var v63 bool = false
                                                                                            var v64 bool = (9223372036854775807 < genFn0((fi19 - fi18), 227, 1))
                                                                                            _ = v62
                                                                                            _ = v63
                                                                                            _ = v64
                                                                                        } else {
                                                                                            fmt.Println(v58, fi18)
                                                                                            var v65 int64 = 42
                                                                                            v2 = v2
                                                                                            _ = v65
                                                                                        }
                                                                                    } else {
                                                                                        if len(v4) >= 2 {
                                                                                            v4 = v4[1:len(v4)]
                                                                                        }
                                                                                        v66 := GenS11{F0: 2, F1: (^(v3)), F2: -376}
                                                                                        var v67 float32 = v2
                                                                                        _ = v66
                                                                                        _ = v67
                                                                                    }
                                                                                    fmt.Println(v3)
                                                                                    fmt.Println(v58, v59)
                                                                                case "":
                                                                                    fmt.Println(v14)
                                                                                }
                                                                                var v68 float64 = ((0.1 + 3.14159) + float64(float32((-2.25 / float64(v2)))))
                                                                                _ = v59
                                                                                _ = v68
                                                                            }
                                                                            fmt.Println(fi19, v2, v54)
                                                                            _ = v58
                                                                        case "hello":
                                                                            fmt.Println(fi19)
                                                                        case "a":
                                                                            v3 += (-(v54))
                                                                            var v69 uint16 = 2
                                                                            fmt.Println(fi19, v54)
                                                                            _ = v69
                                                                        }
                                                                        v70 := GenS9{F0: 0, F1: false}
                                                                        fmt.Println(v3)
                                                                        _ = v57
                                                                        _ = v70
                                                                    } else {
                                                                        var v71 uint8 = 128
                                                                        fmt.Println(fi18)
                                                                        switch "" {
                                                                        case "hello":
                                                                            var v72 uint = 1
                                                                            _ = v72
                                                                        case "a":
                                                                            switch "foo bar" {
                                                                            case "":
                                                                                fmt.Println(v54)
                                                                            case "hello":
                                                                                var v73 int64 = 9223372036854775807
                                                                                _ = v73
                                                                            case "foo bar":
                                                                                v3 += (v54 << ((42) & 63))
                                                                                v74 := GenS11{F0: 1, F1: 42, F2: -32768}
                                                                                if (1 == 698) {
                                                                                    v75 := GenS12{F0: fi18}
                                                                                    _ = v75
                                                                                }
                                                                                _ = v74
                                                                            default:
                                                                                var v76 int16 = 42
                                                                                _ = v76
                                                                            }
                                                                        case "foo bar":
                                                                            var v77 int8 = (v54 / ((v14) | 1))
                                                                            _ = v77
                                                                        default:
                                                                            var v78 int16 = 616
                                                                            _ = v78
                                                                        }
                                                                        _ = v71
                                                                    }
                                                                    _ = v54
                                                                } else {
                                                                    v79 := GenS3{F0: 0, F1: 2147483647}
                                                                    v80 := GenS7{F0: 10, F1: 258}
                                                                    _ = v79
                                                                    _ = v80
                                                                }
                                                                _ = len(v4)
                                                            }
                                                        }
                                                        v81 := GenS5{F0: 87}
                                                        var v82 int16 = 1000
                                                        _ = v43
                                                        _ = v81
                                                        _ = v82
                                                    }
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        switch "" {
                                                        case "a":
                                                            v1 = append(v1, 128)
                                                        }
                                                    case "":
                                                        v3 += (^(v14))
                                                    }
                                                    _ = v35
                                                }
                                            default:
                                                switch ("a" + ("foo bar" + "hello")) {
                                                case "\n\t":
                                                    var v83 int8 = 42
                                                    _ = v83
                                                case "":
                                                    if len(v4) > 0 {
                                                        v4[0] = -769
                                                    }
                                                case "a":
                                                    fmt.Println(v14, v3)
                                                default:
                                                    v14 += (-(v14))
                                                }
                                            }
                                        }
                                    }
                                } else {
                                    v3 -= (v14 / ((-1) | 1))
                                    var v84 int16 = 0
                                    _ = v84
                                }
                                var v85 int32 = 740
                                switch "a" {
                                case "a":
                                    if true {
                                        fmt.Println(v2, v2, fi19)
                                    }
                                }
                                _ = v85
                            } else {
                                v2 += float32(1.5)
                                v14 = (v14 << ((971) & 63))
                            }
                        }
                        _ = fi19
                    }
                    for fi86 := 0; fi86 < 9; fi86++ {
                        v14 = (-(v3))
                        var v87 uint64 = 4294967295
                        _ = fi86
                        _ = v87
                    }
                    fmt.Println(v3, v2)
                    _ = fi18
                }
                if true {
                    switch ("" + "\n\t") {
                    case "hello":
                        for fi88 := 0; fi88 < 8; fi88++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v89 uint8 = 127
                            _ = fi88
                            _ = v89
                        }
                    case "":
                        v3 += v3
                    default:
                        if (((!(65535 <= 277)) || true) || ((v3 ^ v14) <= (v14 ^ (v14 | v14)))) {
                            var v90 uint8 = 128
                            _ = v90
                        }
                    }
                }
            }
        }
        if true {
            if true {
                if false {
                    v91 := GenS13{F0: 1, F1: 402}
                    _ = v91
                }
            } else {
                if true {
                    if (!(0 <= 239)) {
                        if (false && false) {
                            v14 = 119
                        }
                        _ = len(v4)
                        for fi92 := 0; fi92 < 18; fi92++ {
                            for fi93 := 0; fi93 < 11; fi93++ {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v4 = append(v4, 32767)
                                v94 := GenS14{F0: -1}
                                _ = fi93
                                _ = v94
                            }
                            var v95 float64 = 0.1
                            _ = fi92
                            _ = v95
                        }
                    }
                    v96 := GenS4{F0: 858, F1: -9223372036854775808, F2: (^(v14))}
                    _ = v96
                }
                v97 := GenS15{F0: 65535}
                _ = v97
            }
            v2 = (((v2 / v2) / (v2 * 0.1)) + v2)
            switch ("a" + ("hello" + "foo bar")) {
            case "a":
                v3 = (v3 % ((v3) | 1))
            case "hello":
                fmt.Println(v14, v3, v3)
            case "foo bar":
                v14 = (v3 * (v14 | v14))
            default:
                for fi98 := 0; fi98 < 14; fi98++ {
                    var v99 int8 = v3
                    for fi100 := 0; fi100 < 14; fi100++ {
                        v101 := GenS16{F0: 42, F1: 42}
                        var v102 float32 = v2
                        _ = fi100
                        _ = v101
                        _ = v102
                    }
                    v3 = (v3 ^ 0)
                    _ = fi98
                    _ = v99
                }
            }
        }
        var v103 int8 = 42
        _ = v103
    } else {
        v104 := GenS8{F0: 383, F1: 0}
        if false {
            var v105 uint32 = 2
            v14 = v3
            var v106 float64 = (-2.25 / ((((0.0 - (((-2.25 * (((float64((1.5 + float32(float64(v2)))) * ((1.5 * 0.1) + 0.1)) * 3.14159) - 1e10)) + 3.14159) - (float64(((3.14159 / v2) / ((v2 + (float32(0.1) - float32(-2.25))) / float32(1e10)))) - 1e10))) / -2.25) - 1e10) / 3.14159))
            _ = v105
            _ = v106
        } else {
            v3 += (v14 / ((-29) | 1))
            if false {
                var v107 float32 = v2
                _ = v107
            } else {
                v108 := GenS17{F0: 0, F1: 2147483647}
                v109 := GenS2{F0: 255, F1: 298, F2: 1}
                _ = v108
                _ = v109
            }
            switch "" {
            case "":
                fmt.Println(v2, v2, v3)
            case "foo bar":
                fmt.Println(v14, v2)
            default:
                v3 = 74
            }
        }
        var v110 float32 = v2
        _ = v104
        _ = v110
    }
    var v111 int64 = -1
    var v112 int32 = 1
    if len(v1) > 0 {
        v1[0] = -631
    }
    v111 -= genFn0(32768, 42, 1)
    v111 -= v111
    v113 := GenS10{F0: 0}
    fmt.Println(v14, v111, v2)
    var v114 int8 = 0
    v2 += (v2 - -2.25)
    for fi115 := 0; fi115 < 13; fi115++ {
        v111 -= -9223372036854775808
        for fi116 := 0; fi116 < 6; fi116++ {
            fmt.Println(v2)
            var v117 int8 = (-(v114))
            if len(v4) > 0 {
                v4[0] = (fi115 / ((32767) | 1))
            }
            _ = fi116
            _ = v117
        }
        v112 -= (v112 | int32(v14))
        _ = fi115
    }
    var v118 uint32 = 65535
    var v119 float32 = (v2 - (3.14159 + v2))
    v120 := GenS18{F0: "\n\t", F1: 23}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v111:", v111)
    fmt.Println("v112:", v112)
    fmt.Println("v114:", v114)
    fmt.Println("v118:", v118)
    fmt.Println("v119:", v119)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v111
    _ = v112
    _ = v113
    _ = v114
    _ = v118
    _ = v119
    _ = v120
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
