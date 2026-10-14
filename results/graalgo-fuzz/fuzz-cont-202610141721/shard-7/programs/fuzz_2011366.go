package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 int8; F1 uint64; F2 int8 }
type GenS5 struct { F0 int16; F1 uint64 }
type GenS6 struct { F0 uint32; F1 bool }
type GenS7 struct { F0 int32; F1 int64; F2 uint32 }
type GenS8 struct { F0 uint16; F1 uint32; F2 uint }
type GenS9 struct { F0 int8; F1 int64; F2 bool }
type GenS10 struct { F0 uint16; F1 int }
type GenS11 struct { F0 uint32; F1 int }
type GenS12 struct { F0 uint; F1 int; F2 uint8 }
type GenS13 struct { F0 int64; F1 uint }
type GenS14 struct { F0 uint; F1 uint64; F2 int8 }
type GenS15 struct { F0 string }
type GenS16 struct { F0 int32; F1 uint8 }
type GenS17 struct { F0 bool }
type GenS18 struct { F0 uint; F1 string; F2 int64 }
type GenS19 struct { F0 uint32; F1 int32 }
type GenS20 struct { F0 uint32 }
type GenS21 struct { F0 bool; F1 int32; F2 int8 }
func genFn0(p0 uint, p1 bool, p2 int32) int8 {
    if false {
        if ((3.14159 + (1.5 + (3.14159 - 1e10))) >= 1e10) {
            switch ("foo bar" + (("" + ("a" + "\n\t")) + "a")) {
            case "\n\t":
                v5 := GenS1{F0: 0, F1: 2}
                _ = v5
            case "a":
                var v6 int8 = -1
                _ = v6
            default:
                p1 = ("hello" == "hello")
            }
        }
        p0 += p0
    } else {
        var v7 uint8 = 255
        for fi8 := 0; fi8 < 7; fi8++ {
            v9 := []int{}
            for fi10 := 0; fi10 < 8; fi10++ {
                p1 = true
                v7 += v7
                var v11 uint32 = 245
                var v12 float32 = 3.14159
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = len(v9)
            _ = fi8
            _ = v9
        }
        p2 = (p2 >> (((p0 << (((p0 & p0)) & 63))) & 63))
        _ = v7
    }
    var v13 int16 = 396
    p0 = (p0 - (p0 - p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 42
}

func main() {
    var v1 string = ("\n\t" + "a")
    v2 := make(map[string]bool)
    v3 := make(map[int]int)
    var v4 int16 = -184
    for fi14 := 0; fi14 < 14; fi14++ {
        if (153 >= 300) {
            fmt.Println(v4, v4, v4)
            for fi15 := 0; fi15 < 6; fi15++ {
                v16 := GenS1{F0: 627, F1: 149}
                fmt.Println(v4, fi15)
                _ = fi15
                _ = v16
            }
        } else {
            v17 := []int{fi14, (fi14 % ((32768) | 1)), 1, fi14}
            if len(v17) > 0 {
                v17[0] = fi14
            }
            var v18 int = (fi14 % ((354) | 1))
            _ = v17
            _ = v18
        }
        switch fi14 {
        case -890:
            switch -9223372036854775808 {
            case 4294967295:
                switch (^(fi14)) {
                case -2147483648:
                    fmt.Println(v4, v1, v1)
                }
            default:
                var v19 int64 = 1
                _ = v19
            }
        case 670:
            var v20 int8 = -128
            _ = v20
        case -2147483648:
            if ((0.1 - 3.14159) <= (0.0 * -2.25)) {
                v21 := []int{}
                _ = v21
            }
        }
        v22 := []int{fi14, (fi14 >> ((265) & 63))}
        _ = fi14
        _ = v22
    }
    v1 = v1
    fmt.Println(v4, v1)
    fmt.Println(v1, v4)
    v23 := []int{-9223372036854775808, 32767, 822}
    for fi24 := 0; fi24 < 11; fi24++ {
        switch (fi24 - fi24) {
        case -93:
            fmt.Println(fi24, fi24, fi24)
        default:
            fmt.Println(v1, fi24)
        }
        if false {
            switch v1 {
            case "hello":
                var v25 int = fi24
                _ = v25
            case "foo bar":
                for fi26 := 0; fi26 < 4; fi26++ {
                    switch (fi24 >> ((2) & 63)) {
                    case 2147483647:
                        var v27 uint32 = 580
                        _ = v27
                    case 32768:
                        v28 := GenS1{F0: 493, F1: 396}
                        _ = v28
                    case -1:
                        v4 = (v4 + v4)
                    }
                    v29 := GenS1{F0: 1, F1: 2}
                    v4 = v4
                    _ = fi26
                    _ = v29
                }
            default:
                if len(v23) > 0 {
                    v23[0] = -2147483648
                }
            }
            v30 := GenS2{F0: 4294967295, F1: 65535}
            _ = v30
        } else {
            switch (fi24 + fi24) {
            case -97:
                if true {
                    if false {
                        switch "a" {
                        case "\n\t":
                            v1 = "\n\t"
                        case "a":
                            for fi31 := 0; fi31 < 15; fi31++ {
                                if false {
                                    v4 -= v4
                                    v4 += (v4 << ((65535) & 63))
                                    var v32 uint16 = 42
                                    _ = v32
                                } else {
                                    switch v1 {
                                    case "\n\t":
                                        v4 = v4
                                    case "":
                                        delete(v3, fi24)
                                    case "foo bar":
                                        var v33 uint16 = 0
                                        _ = v33
                                    }
                                    if true {
                                        switch (v1 + (v1 + v1)) {
                                        case "a":
                                            switch fi24 {
                                            case 42:
                                                v4 = v4
                                                var v34 string = v1
                                                var v35 int8 = 0
                                                _ = v34
                                                _ = v35
                                            default:
                                                switch "" {
                                                case "hello":
                                                    v1 = v1
                                                case "a":
                                                    var v36 uint32 = 572
                                                    _ = v36
                                                default:
                                                    v4 += (v4 - 32767)
                                                }
                                                var v37 float32 = 1e10
                                                _ = v37
                                            }
                                        case "foo bar":
                                            var v38 int8 = 1
                                            _ = v38
                                        case "\n\t":
                                            var v39 int8 = 1
                                            _ = v39
                                        }
                                        v40 := GenS3{F0: 255, F1: 100}
                                        v41 := GenS4{F0: -1, F1: 0, F2: genFn0(255, (0 > -1), -1)}
                                        _ = v40
                                        _ = v41
                                    }
                                }
                                switch (fi24 / ((225) | 1)) {
                                case 42:
                                    var v42 int8 = 127
                                    _ = v42
                                default:
                                    v4 = v4
                                }
                                if (true || true) {
                                    var v43 int64 = 9223372036854775807
                                    delete(v2, v1)
                                    v44 := GenS5{F0: -723, F1: 42}
                                    var v45 uint32 = 42
                                    v45 += v45
                                    _ = v43
                                    _ = v44
                                    _ = v45
                                } else {
                                    switch v1 {
                                    case "foo bar":
                                        v46 := GenS6{F0: 1, F1: (255 < 0)}
                                        _ = v46
                                    case "":
                                        var v47 int = 4294967295
                                        _ = v47
                                    case "\n\t":
                                        fmt.Println(fi31, v1)
                                    default:
                                        v4 += (v4 % ((-744) | 1))
                                    }
                                    if true {
                                        var v48 int16 = (v4 + (v4 ^ v4))
                                        if false {
                                            var v49 float64 = 1e10
                                            if (true || (18446744073709551615 == 4294967295)) {
                                                v49 += 1e10
                                                var v50 uint64 = 950
                                                var v51 uint = 223
                                                _ = v50
                                                _ = v51
                                            } else {
                                                switch fi31 {
                                                case 42:
                                                    if false {
                                                        var v52 bool = true
                                                        fmt.Println(v1, v4)
                                                        _ = v52
                                                    } else {
                                                        var v53 int64 = 1
                                                        switch fi31 {
                                                        case 414:
                                                            if (674 >= 4294967295) {
                                                                v1 = ((v1 + v1) + v1)
                                                                v54 := GenS7{F0: 1, F1: int64(v4), F2: 0}
                                                                v1 = v1
                                                                _ = v54
                                                            } else {
                                                                switch (v1 + v1) {
                                                                case "":
                                                                    if true {
                                                                        v3[(fi31 / (((fi31 % ((int(v53)) | 1))) | 1))] = 128
                                                                        var v55 string = v1
                                                                        if (-42 > (v4 / ((v48) | 1))) {
                                                                            if false {
                                                                                var v56 uint8 = 42
                                                                                _ = v56
                                                                            }
                                                                            var v57 float64 = v49
                                                                            if (v49 == v57) {
                                                                                if true {
                                                                                    v49 = -2.25
                                                                                    var v58 uint64 = 961
                                                                                    if false {
                                                                                        if (!false) {
                                                                                            switch (fi31 + (fi31 << ((255) & 63))) {
                                                                                            case 1:
                                                                                                if (!(2 >= 2147483648)) {
                                                                                                    if ((int32(v58) == 0) && (true && true)) {
                                                                                                        var v59 int16 = v48
                                                                                                        _ = v59
                                                                                                    }
                                                                                                    v60 := GenS2{F0: 2, F1: 566}
                                                                                                    switch v55 {
                                                                                                    case "\n\t":
                                                                                                        var v61 int32 = 0
                                                                                                        _ = v61
                                                                                                    case "foo bar":
                                                                                                        v57 -= ((((v49 - v49) * v57) + v49) / float64(-2.25))
                                                                                                    case "hello":
                                                                                                        delete(v3, fi24)
                                                                                                    }
                                                                                                    var v62 uint8 = 63
                                                                                                    _ = v60
                                                                                                    _ = v62
                                                                                                } else {
                                                                                                    fmt.Println(v55, v58, v53)
                                                                                                    v23 = append(v23, (fi24 << ((247) & 63)))
                                                                                                    v48 -= (v4 ^ -417)
                                                                                                    v58 = v58
                                                                                                }
                                                                                            case 32767:
                                                                                                var v63 bool = false
                                                                                                _ = v63
                                                                                            case -2147483648:
                                                                                                switch fi31 {
                                                                                                case 2147483647:
                                                                                                    if false {
                                                                                                        if false {
                                                                                                            var v64 uint8 = 154
                                                                                                            v49 -= v49
                                                                                                            _ = v64
                                                                                                        } else {
                                                                                                            if (false && (1 <= 255)) {
                                                                                                                fmt.Println(v48, v1, v53)
                                                                                                                switch (fi31 + (fi24 | 70)) {
                                                                                                                case 309:
                                                                                                                    var v65 uint32 = 4294967295
                                                                                                                    _ = v65
                                                                                                                case 42:
                                                                                                                    v49 = 0.1
                                                                                                                default:
                                                                                                                    _ = len(v23)
                                                                                                                }
                                                                                                            } else {
                                                                                                                var v66 int = (fi31 << ((uint(v58)) & 63))
                                                                                                                _ = v66
                                                                                                            }
                                                                                                            v23 = append(v23, fi24)
                                                                                                            switch v1 {
                                                                                                            case "\n\t":
                                                                                                                var v67 int32 = 42
                                                                                                                _ = v67
                                                                                                            case "hello":
                                                                                                                if false {
                                                                                                                    _ = len(v23)
                                                                                                                    v68 := GenS7{F0: 42, F1: v53, F2: 0}
                                                                                                                    _ = v68
                                                                                                                }
                                                                                                            case "a":
                                                                                                                var v69 uint64 = v58
                                                                                                                _ = v69
                                                                                                            default:
                                                                                                                v1 = v1
                                                                                                            }
                                                                                                        }
                                                                                                    } else {
                                                                                                        switch fi24 {
                                                                                                        case -472:
                                                                                                            switch v1 {
                                                                                                            case "a":
                                                                                                                var v70 uint = 106
                                                                                                                _ = v70
                                                                                                            case "hello":
                                                                                                                var v71 int8 = 100
                                                                                                                v58 = v58
                                                                                                                _ = v71
                                                                                                            case "foo bar":
                                                                                                                if (!(0.0 != 0.1)) {
                                                                                                                    v48 += v4
                                                                                                                    v1 = v1
                                                                                                                    var v72 int = (fi31 / (((fi24 / (((fi31 | fi24)) | 1))) | 1))
                                                                                                                    _ = v72
                                                                                                                } else {
                                                                                                                    fmt.Println(v55, v1, v48)
                                                                                                                }
                                                                                                            }
                                                                                                        case -459:
                                                                                                            var v73 uint8 = 219
                                                                                                            _ = v73
                                                                                                        }
                                                                                                    }
                                                                                                case 42:
                                                                                                    var v74 float64 = 0.1
                                                                                                    _ = v74
                                                                                                case 988:
                                                                                                    if (0.0 >= 0.1) {
                                                                                                        _ = len(v3)
                                                                                                    } else {
                                                                                                        v53 = (v53 >> ((587) & 63))
                                                                                                    }
                                                                                                default:
                                                                                                    if false {
                                                                                                        fmt.Println(v55, v4)
                                                                                                        fmt.Println(fi31, v57)
                                                                                                        v49 += float64(0.1)
                                                                                                    }
                                                                                                    v2[((("\n\t" + "\n\t") + v1) + "foo bar")] = ((false || ((!(((2 == 42) || (genFn0(65535, (1 > 0), -2147483648) == 100)) || (0 == 605))) || true)) || true)
                                                                                                    var v75 uint8 = 2
                                                                                                    _ = v75
                                                                                                }
                                                                                            default:
                                                                                                var v76 int64 = (v53 >> ((560) & 63))
                                                                                                _ = v76
                                                                                            }
                                                                                        }
                                                                                        v77 := GenS2{F0: (v58 & v58), F1: 2}
                                                                                        _ = v77
                                                                                    }
                                                                                    v78 := GenS6{F0: 730, F1: (!(fi31 > fi24))}
                                                                                    _ = v58
                                                                                    _ = v78
                                                                                } else {
                                                                                    v79 := GenS7{F0: -1, F1: (-(v53)), F2: 65535}
                                                                                    var v80 uint8 = 200
                                                                                    var v81 float32 = 1.5
                                                                                    _ = v79
                                                                                    _ = v80
                                                                                    _ = v81
                                                                                }
                                                                                v4 = v4
                                                                                fmt.Println(v1, v55)
                                                                            } else {
                                                                                var v82 uint16 = 68
                                                                                var v83 int = (-(fi24))
                                                                                _ = v82
                                                                                _ = v83
                                                                            }
                                                                            _ = v57
                                                                        } else {
                                                                            fmt.Println(v49, fi31)
                                                                        }
                                                                        _ = v55
                                                                    } else {
                                                                        v84 := v3[fi31]
                                                                        v49 -= 1e10
                                                                        v3[(v84 | v84)] = fi31
                                                                        _ = len(v23)
                                                                        _ = v84
                                                                    }
                                                                case "hello":
                                                                    v4 -= v48
                                                                case "\n\t":
                                                                    if len(v23) >= 2 {
                                                                        v23 = v23[1:len(v23)]
                                                                    }
                                                                }
                                                            }
                                                        case 1:
                                                            v85 := GenS8{F0: 42, F1: 65535, F2: 2}
                                                            _ = v85
                                                        default:
                                                            fmt.Println(v53)
                                                        }
                                                        delete(v2, (v1 + v1))
                                                        _ = v53
                                                    }
                                                case -569:
                                                    if true {
                                                        var v86 uint = 2
                                                        var v87 float32 = (((0.0 + -2.25) + 0.1) + -2.25)
                                                        _ = v86
                                                        _ = v87
                                                    }
                                                }
                                                if false {
                                                    _ = len(v23)
                                                    v88 := GenS9{F0: 0, F1: 1, F2: false}
                                                    _ = v88
                                                }
                                                switch "a" {
                                                case "hello":
                                                    if ((!(true && (112 >= 235))) || (160 != 182)) {
                                                        var v89 float64 = -2.25
                                                        fmt.Println(v4, v48)
                                                        _ = v89
                                                    }
                                                case "\n\t":
                                                    switch (("foo bar" + (v1 + "hello")) + v1) {
                                                    case "hello":
                                                        v4 = v48
                                                    default:
                                                        var v90 float64 = v49
                                                        fmt.Println(v90)
                                                        _ = v90
                                                    }
                                                default:
                                                    var v91 int32 = 108
                                                    _ = v91
                                                }
                                            }
                                            var v92 int8 = 0
                                            _ = v49
                                            _ = v92
                                        } else {
                                            var v93 uint = 65535
                                            v48 -= (v4 / (((v48 * v4)) | 1))
                                            fmt.Println(fi31)
                                            if (true && (!false)) {
                                                var v94 int16 = v48
                                                if ((1 != 4294967295) || (!false)) {
                                                    _ = len(v23)
                                                    switch (fi31 / ((-514) | 1)) {
                                                    case 708:
                                                        if (!(false && true)) {
                                                            fmt.Println(v1, fi31, v1)
                                                            switch 0 {
                                                            case 411:
                                                                fmt.Println(v93, fi24, v4)
                                                            case 9223372036854775807:
                                                                if false {
                                                                    fmt.Println(v48)
                                                                }
                                                            }
                                                        } else {
                                                            if (true || (120 > uint8(v93))) {
                                                                v95 := GenS10{F0: 57, F1: -586}
                                                                _ = v95
                                                            }
                                                            switch fi24 {
                                                            case -1:
                                                                if ((!(!(42 > 255))) && false) {
                                                                    fmt.Println(fi24)
                                                                    v4 = (v94 * v48)
                                                                    var v96 int8 = -128
                                                                    _ = v96
                                                                } else {
                                                                    if ((v93 % ((v93) | 1)) <= (v93 << (((v93 % (((^(v93))) | 1))) & 63))) {
                                                                        switch v1 {
                                                                        case "hello":
                                                                            if (true || (v1 < v1)) {
                                                                                v1 = (v1 + ((v1 + ("\n\t" + "a")) + ((v1 + v1) + "hello")))
                                                                            } else {
                                                                                var v97 uint8 = 136
                                                                                _ = v97
                                                                            }
                                                                            if (!true) {
                                                                                v4 = v94
                                                                            } else {
                                                                                v98 := GenS9{F0: 0, F1: -9223372036854775808, F2: false}
                                                                                var v99 uint64 = 18446744073709551615
                                                                                _ = v98
                                                                                _ = v99
                                                                            }
                                                                            var v100 bool = (-53 != 292)
                                                                            _ = v100
                                                                        default:
                                                                            if false {
                                                                                v23 = append(v23, (fi24 / ((2147483647) | 1)))
                                                                                var v101 int16 = (v4 ^ -592)
                                                                                fmt.Println(v93)
                                                                                _ = v101
                                                                            } else {
                                                                                var v102 int32 = 2147483647
                                                                                if true {
                                                                                    if (!(true || false)) {
                                                                                        switch 127 {
                                                                                        case 62:
                                                                                            switch fi24 {
                                                                                            case 883:
                                                                                                v93 -= (v93 >> (((v93 << ((v93) & 63))) & 63))
                                                                                            default:
                                                                                                if (true || false) {
                                                                                                    var v103 float64 = 3.14159
                                                                                                    v48 += v94
                                                                                                    switch v1 {
                                                                                                    case "hello":
                                                                                                        var v104 float32 = 3.14159
                                                                                                        _ = v104
                                                                                                    }
                                                                                                    _ = v103
                                                                                                }
                                                                                            }
                                                                                        case 0:
                                                                                            v105 := v3[1]
                                                                                            _ = v105
                                                                                        default:
                                                                                            v106 := GenS11{F0: 6, F1: (^(fi24))}
                                                                                            v48 = v4
                                                                                            _ = v106
                                                                                        }
                                                                                        var v107 float64 = -2.25
                                                                                        switch (v1 + v1) {
                                                                                        case "hello":
                                                                                            v48 = v48
                                                                                        case "\n\t":
                                                                                            v93 -= v93
                                                                                            v107 += v107
                                                                                        case "a":
                                                                                            var v108 uint64 = 2
                                                                                            _ = v108
                                                                                        default:
                                                                                            v109 := GenS12{F0: v93, F1: fi24, F2: 2}
                                                                                            switch v1 {
                                                                                            case "":
                                                                                                v110 := GenS11{F0: 247, F1: (fi31 ^ fi31)}
                                                                                                _ = v110
                                                                                            case "foo bar":
                                                                                                v102 = (^(v102))
                                                                                            case "hello":
                                                                                                var v111 int = 128
                                                                                                _ = v111
                                                                                            }
                                                                                            switch v1 {
                                                                                            case "\n\t":
                                                                                                v112 := v2["\n\t"]
                                                                                                _ = v112
                                                                                            case "a":
                                                                                                var v113 string = ("\n\t" + v1)
                                                                                                _ = v113
                                                                                            }
                                                                                            _ = v109
                                                                                        }
                                                                                        _ = v107
                                                                                    } else {
                                                                                        v114 := GenS6{F0: 2, F1: (v48 <= v94)}
                                                                                        _ = v114
                                                                                    }
                                                                                }
                                                                                v94 = v4
                                                                                _ = v102
                                                                            }
                                                                        }
                                                                        v48 = 1
                                                                        var v115 float32 = 0.0
                                                                        _ = v115
                                                                    } else {
                                                                        fmt.Println(v4, v94)
                                                                        v93 += v93
                                                                        fmt.Println(v48, fi24, v48)
                                                                    }
                                                                    v48 += v48
                                                                    v3[int(v4)] = fi31
                                                                }
                                                            default:
                                                                switch "" {
                                                                case "foo bar":
                                                                    v4 = 1
                                                                default:
                                                                    _ = len(v2)
                                                                }
                                                            }
                                                        }
                                                    case 42:
                                                        var v116 int = (fi24 % (((^(fi31))) | 1))
                                                        _ = v116
                                                    case -900:
                                                        var v117 uint64 = 759
                                                        _ = v117
                                                    default:
                                                        if true {
                                                            v48 += v4
                                                        }
                                                    }
                                                } else {
                                                    v118 := GenS11{F0: 2, F1: (fi31 % (((-(fi31))) | 1))}
                                                    v119 := GenS3{F0: 0, F1: 0}
                                                    _ = v118
                                                    _ = v119
                                                }
                                                _ = v94
                                            }
                                            v1 = v1
                                            var v120 uint64 = 2
                                            switch v1 {
                                            case "":
                                                switch (v1 + (v1 + v1)) {
                                                case "hello":
                                                    v1 = v1
                                                case "a":
                                                    v2[(v1 + v1)] = ((!true) || true)
                                                case "foo bar":
                                                    fmt.Println(v120, v4, fi24)
                                                }
                                            case "hello":
                                                v1 = ("a" + v1)
                                            case "a":
                                                v3[(fi31 + fi31)] = 32767
                                            default:
                                                v121 := GenS13{F0: 9223372036854775807, F1: (v93 << ((v93) & 63))}
                                                _ = v121
                                            }
                                            _ = v120
                                            _ = v93
                                        }
                                        fmt.Println(fi31, v4)
                                        _ = v48
                                    } else {
                                        if true {
                                            v4 -= 32767
                                            v122 := v3[(fi31 >> ((1) & 63))]
                                            if true {
                                                v123 := GenS14{F0: 1, F1: 2, F2: -56}
                                                _ = v123
                                            } else {
                                                var v124 int16 = v4
                                                v4 = v4
                                                _ = v124
                                            }
                                            switch fi31 {
                                            case 0:
                                                var v125 uint = 240
                                                _ = v125
                                            default:
                                                if false {
                                                    switch (fi31 & (v122 & v122)) {
                                                    case 4294967295:
                                                        var v126 uint = 255
                                                        _ = v126
                                                    }
                                                    var v127 int16 = v4
                                                    if false {
                                                        _ = len(v2)
                                                    }
                                                    _ = v127
                                                } else {
                                                    v122 -= v122
                                                    var v128 uint = 18446744073709551615
                                                    switch v1 {
                                                    case "":
                                                        v128 = (v128 / (((v128 << (((v128 >> ((v128) & 63))) & 63))) | 1))
                                                    }
                                                    _ = v128
                                                }
                                            }
                                            v129 := GenS14{F0: 255, F1: 18446744073709551615, F2: 127}
                                            v23 = append(v23, 4294967295)
                                            fmt.Println(v122, fi31, v4)
                                            _ = v122
                                            _ = v129
                                        } else {
                                            if false {
                                                fmt.Println(v4, fi31, v1)
                                            }
                                            v3[fi31] = (fi31 | fi31)
                                            var v130 int16 = v4
                                            _ = v130
                                        }
                                    }
                                    switch (fi31 ^ (fi24 >> ((65535) & 63))) {
                                    case 4294967295:
                                        v131 := GenS15{F0: (v1 + "")}
                                        _ = v131
                                    default:
                                        v132 := GenS4{F0: 0, F1: 819, F2: 42}
                                        _ = v132
                                    }
                                    v4 += v4
                                    if (!true) {
                                        switch v1 {
                                        case "foo bar":
                                            if true {
                                                var v133 int = (fi31 ^ 597)
                                                if true {
                                                    if false {
                                                        v1 = (v1 + (("hello" + "hello") + "a"))
                                                        v4 += 1
                                                    } else {
                                                        if (true && false) {
                                                            if (104 <= genFn0(1, false, 2147483647)) {
                                                                if true {
                                                                    v1 = v1
                                                                    var v134 uint32 = 4294967295
                                                                    _ = v134
                                                                } else {
                                                                    var v135 uint64 = 18446744073709551615
                                                                    _ = v135
                                                                }
                                                            }
                                                            switch (fi24 / ((fi31) | 1)) {
                                                            case 42:
                                                                v1 = (("\n\t" + "") + v1)
                                                            case 32768:
                                                                var v136 uint32 = 2
                                                                _ = v136
                                                            }
                                                        } else {
                                                            if true {
                                                                switch fi24 {
                                                                case 42:
                                                                    switch (v1 + "") {
                                                                    case "a":
                                                                        fmt.Println(v4)
                                                                    case "":
                                                                        v137 := GenS16{F0: 1, F1: 71}
                                                                        _ = v137
                                                                    case "\n\t":
                                                                        v138 := GenS7{F0: 1, F1: int64(fi24), F2: 42}
                                                                        var v139 int = fi31
                                                                        _ = v138
                                                                        _ = v139
                                                                    default:
                                                                        if (2147483647 <= (fi31 ^ fi31)) {
                                                                            fmt.Println(v133, v4, v133)
                                                                            v1 = (v1 + v1)
                                                                            var v140 uint16 = 65535
                                                                            _ = v140
                                                                        } else {
                                                                            if (1.5 > 1.5) {
                                                                                v3[fi24] = (-(fi31))
                                                                                v4 = -32768
                                                                                if false {
                                                                                    v133 = (fi24 & 549)
                                                                                } else {
                                                                                    switch (v1 + "a") {
                                                                                    case "\n\t":
                                                                                        switch v133 {
                                                                                        case 128:
                                                                                            var v141 float32 = float32(-2.25)
                                                                                            _ = v141
                                                                                        default:
                                                                                            switch 2147483647 {
                                                                                            case 32768:
                                                                                                v142 := GenS5{F0: (v4 << ((565) & 63)), F1: 1}
                                                                                                _ = v142
                                                                                            default:
                                                                                                fmt.Println(v1, v133)
                                                                                            }
                                                                                        }
                                                                                    case "foo bar":
                                                                                        var v143 uint = 2
                                                                                        _ = v143
                                                                                    }
                                                                                    v1 = v1
                                                                                    v133 -= (fi24 & (v133 * (fi24 & fi24)))
                                                                                }
                                                                            }
                                                                            v1 = (("a" + (v1 + (v1 + v1))) + "hello")
                                                                            v4 = v4
                                                                        }
                                                                    }
                                                                case -785:
                                                                    v144 := GenS14{F0: 4294967295, F1: 697, F2: 0}
                                                                    _ = v144
                                                                case -2147483648:
                                                                    v145 := GenS4{F0: 42, F1: 358, F2: 0}
                                                                    _ = v145
                                                                default:
                                                                    if len(v23) > 0 {
                                                                        v23[0] = fi31
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        v23 = append(v23, (-(fi31)))
                                                        var v146 uint16 = 65535
                                                        _ = v146
                                                    }
                                                    switch (("hello" + (v1 + v1)) + "") {
                                                    case "hello":
                                                        if (((!true) && false) && false) {
                                                            if (2 > 0) {
                                                                fmt.Println(v133, v133, fi24)
                                                                if true {
                                                                    v147 := GenS17{F0: ((30 > 843) || (1e10 >= float32(-2.25)))}
                                                                    if true {
                                                                        fmt.Println(v1, v1)
                                                                        switch v1 {
                                                                        case "foo bar":
                                                                            fmt.Println(fi24)
                                                                        case "\n\t":
                                                                            v148 := GenS5{F0: -569, F1: 42}
                                                                            _ = v148
                                                                        }
                                                                        v149 := v3[(-(fi31))]
                                                                        _ = v149
                                                                    }
                                                                    _ = v147
                                                                }
                                                                v133 -= 365
                                                            } else {
                                                                v150 := v3[(fi24 + (fi31 & -2147483648))]
                                                                v151 := GenS18{F0: 521, F1: v1, F2: 847}
                                                                v133 += (fi24 >> ((18446744073709551615) & 63))
                                                                _ = v150
                                                                _ = v151
                                                            }
                                                            v1 = v1
                                                            var v152 bool = ((!(18446744073709551615 >= 1)) || true)
                                                            _ = v152
                                                        } else {
                                                            switch v1 {
                                                            case "hello":
                                                                v133 -= fi24
                                                            }
                                                            v133 += 965
                                                            v4 = v4
                                                        }
                                                    }
                                                    fmt.Println(v133, v133, v133)
                                                    _ = len(v2)
                                                    v4 += 42
                                                    v153 := GenS19{F0: 2147483647, F1: -658}
                                                    var v154 string = "foo bar"
                                                    fmt.Println(v4, v133)
                                                    var v155 string = ((v154 + v154) + ((v154 + (v1 + (v154 + v154))) + v1))
                                                    v156 := GenS1{F0: 32767, F1: 42}
                                                    _ = v153
                                                    _ = v154
                                                    _ = v155
                                                    _ = v156
                                                } else {
                                                    v1 = v1
                                                    fmt.Println(v1)
                                                }
                                                v23 = append(v23, (-(fi24)))
                                                _ = v133
                                            } else {
                                                fmt.Println(fi24, fi31)
                                                var v157 float64 = float64(float32((1e10 / 1.0)))
                                                _ = v157
                                            }
                                        default:
                                            switch ((v1 + (v1 + v1)) + (v1 + ((v1 + ((((v1 + v1) + v1) + "hello") + (v1 + "hello"))) + v1))) {
                                            case "hello":
                                                if (65535 >= 0) {
                                                    v158 := GenS7{F0: 2147483647, F1: int64(fi31), F2: 0}
                                                    if (!((true || true) || ((0 != genFn0(4294967295, false, 2147483647)) || false))) {
                                                        v1 = v1
                                                        var v159 uint = 656
                                                        _ = len(v2)
                                                        _ = v159
                                                    } else {
                                                        if (255 != 204) {
                                                            fmt.Println(v1, fi24)
                                                            var v160 uint64 = 18446744073709551615
                                                            fmt.Println(v4, v1)
                                                            _ = v160
                                                        }
                                                    }
                                                    _ = v158
                                                } else {
                                                    if len(v23) > 0 {
                                                        v23[0] = fi24
                                                    }
                                                    v161 := GenS12{F0: 440, F1: fi24, F2: 1}
                                                    v4 -= v4
                                                    _ = v161
                                                }
                                            case "":
                                                if true {
                                                    var v162 uint = 65535
                                                    _ = v162
                                                }
                                            case "a":
                                                if false {
                                                    var v163 bool = (false || true)
                                                    var v164 int8 = -1
                                                    fmt.Println(v1, fi31)
                                                    _ = v163
                                                    _ = v164
                                                }
                                            }
                                        }
                                        v4 = v4
                                        v2[v1] = false
                                    }
                                    delete(v3, fi24)
                                    switch "hello" {
                                    case "\n\t":
                                        v4 -= v4
                                    case "hello":
                                        v165 := GenS4{F0: 0, F1: 309, F2: 81}
                                        _ = v165
                                    case "":
                                        if false {
                                            v4 -= v4
                                            var v166 bool = (!(2 == 2))
                                            if (1 != 1) {
                                                v167 := v3[(fi31 / (((fi31 / (((fi24 << ((0) & 63))) | 1))) | 1))]
                                                switch "a" {
                                                case "foo bar":
                                                    v166 = true
                                                default:
                                                    v166 = v166
                                                }
                                                _ = v167
                                            } else {
                                                v4 = (v4 / ((411) | 1))
                                            }
                                            _ = v166
                                        } else {
                                            var v168 uint8 = 128
                                            _ = v168
                                        }
                                    default:
                                        switch fi24 {
                                        case -9223372036854775808:
                                            if (v1 == v1) {
                                                if true {
                                                    if true {
                                                        var v169 int64 = -450
                                                        delete(v3, (fi31 ^ (fi31 * fi24)))
                                                        _ = v169
                                                    }
                                                    if false {
                                                        var v170 int8 = -49
                                                        _ = v170
                                                    } else {
                                                        switch fi24 {
                                                        case 9223372036854775807:
                                                            if false {
                                                                var v171 int64 = 688
                                                                _ = v171
                                                            }
                                                        }
                                                    }
                                                    var v172 float64 = 0.0
                                                    _ = v172
                                                }
                                                v4 -= v4
                                                var v173 uint64 = 42
                                                _ = v173
                                            }
                                        case 128:
                                            var v174 uint64 = 18446744073709551615
                                            _ = v174
                                        default:
                                            fmt.Println(v1, fi31, fi24)
                                        }
                                    }
                                }
                                _ = fi31
                            }
                        case "foo bar":
                            switch v1 {
                            case "\n\t":
                                var v175 string = "a"
                                _ = v175
                            case "foo bar":
                                switch fi24 {
                                case -132:
                                    var v176 int16 = (v4 * (v4 & 0))
                                    _ = v176
                                case 627:
                                    v1 = v1
                                }
                            }
                        default:
                            v4 -= -4
                        }
                    } else {
                        for fi177 := 0; fi177 < 18; fi177++ {
                            switch v1 {
                            case "a":
                                var v178 int = (fi24 << ((627) & 63))
                                _ = v178
                            default:
                                fmt.Println(v4, fi177)
                            }
                            v4 = (^(v4))
                            _ = fi177
                        }
                        v4 = -182
                        v179 := GenS4{F0: 29, F1: 2, F2: 0}
                        _ = v179
                    }
                } else {
                    fmt.Println(v4, v4)
                    var v180 uint32 = 2147483647
                    _ = v180
                }
            case -360:
                if (v1 >= "") {
                    var v181 int32 = 511
                    _ = v181
                }
            case -643:
                switch -9223372036854775808 {
                case -511:
                    if false {
                        v1 = "a"
                    } else {
                        var v182 string = v1
                        v183 := v2["hello"]
                        for fi184 := 0; fi184 < 14; fi184++ {
                            var v185 uint16 = 255
                            var v186 uint8 = 255
                            v186 = v186
                            var v187 float32 = (-2.25 / (((((1e10 + (1.5 / 0.1)) * ((0.1 * 3.14159) + 3.14159)) + 0.1) * 0.1) * 1e10))
                            var v188 float32 = float32(3.14159)
                            if false {
                                if ((true || (float32((-2.25 - ((float64(v187) + 0.1) / -2.25))) < ((v187 * 3.14159) - (float32(0.0) + v187)))) || (42 != 42)) {
                                    var v189 int8 = 0
                                    v1 = (((v182 + v1) + ("a" + ((v182 + v1) + (v1 + "hello")))) + v182)
                                    v190 := GenS20{F0: 65535}
                                    _ = v189
                                    _ = v190
                                }
                            }
                            var v191 uint16 = v185
                            _ = fi184
                            _ = v185
                            _ = v186
                            _ = v187
                            _ = v188
                            _ = v191
                        }
                        _ = v182
                        _ = v183
                    }
                case 4294967295:
                    var v192 int64 = 1
                    _ = v192
                }
            }
        }
        fmt.Println(v4)
        _ = fi24
    }
    var v193 float32 = float32((((float64(0.0) - -2.25) / 0.1) + (float64(float32(0.1)) + 0.1)))
    v194 := GenS6{F0: 2147483648, F1: true}
    var v195 int16 = (v4 + (v4 ^ (^(v4))))
    for fi196 := 0; fi196 < 4; fi196++ {
        v193 = 0.0
        _ = len(v2)
        _ = fi196
    }
    for fi197 := 0; fi197 < 5; fi197++ {
        var v198 int = fi197
        v199 := GenS14{F0: 321, F1: 18446744073709551615, F2: genFn0(4294967295, true, -2147483648)}
        _ = fi197
        _ = v198
        _ = v199
    }
    if (-1 < 1) {
        var v200 int32 = -285
        _ = v200
    }
    switch 32767 {
    case -935:
        v1 = "a"
    }
    if true {
        _ = len(v2)
        v201 := GenS21{F0: (true && true), F1: -2147483648, F2: genFn0(2, false, 2147483647)}
        _ = v201
    }
    fmt.Println(v4, v193, v1)
    var v202 string = v1
    fmt.Println(v202, v193)
    fmt.Println("v1:", v1)
    fmt.Println("v193:", v193)
    fmt.Println("v195:", v195)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v202:", v202)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v193
    _ = v194
    _ = v195
    _ = v2
    _ = v202
    _ = v23
    _ = v3
    _ = v4
}
